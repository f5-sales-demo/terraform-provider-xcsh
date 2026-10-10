---
page_title: "xcsh_enhanced_firewall_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_enhanced_firewall_policy reference."
---

# xcsh_enhanced_firewall_policy reference

<a id="canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- Property reference

<a id="canonical-1322300112012020-2211210312202331-3311012011100020-3323312113011330-1130331303023212-1221013030211122-1113311313132203-0121010023023311"></a>

### Direct properties for `xcsh_enhanced_firewall_policy`

- [allow_all](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3033303330013011-3303131020130023-0120301102013120-2331313303103320-0103123003302022-2020121003023013-0331012313330210-2223122222322000): complete subsection reference.

- [allowed_destinations](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1121223222222033-1321103111123211-2322303130112211-1320223300030321-3330303232231310-3031333003200220-2001132333000233-3120233221311201): complete subsection reference.

- [allowed_sources](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0310120130032220-3210102231321012-2322312302113320-0123301300100020-0013321230021233-1002110311102222-3211131333310303-3000130033031121): complete subsection reference.

<a id="canonical-0232303010003103-1212210123202020-1030313023333203-1333122301010321-0131022312001333-1120331303210321-3021020330023120-2301221131323233"></a>

<a id="canonical-1001021113210110-1020110133220302-0203132211203201-0110303011210311-3211213123101210-1010033100003100-0302232201020130-2312301300321332"></a>

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

- [denied_destinations](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0321100103231102-3020030330213210-2203000323112003-1130001202103022-2233110030031311-0222201211232332-2011313100032013-3022022111322323): complete subsection reference.

- [denied_sources](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3210321321221232-1330102210103130-1121232320102122-2210300100320211-3233302002001113-2031213013120212-1210221121011303-1102113300303202): complete subsection reference.

- [deny_all](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2221303231002233-3313021213301220-3211322103313302-1020330122012331-3323333321023320-2321002021313111-0211022201312220-0122131303132031): complete subsection reference.

<a id="canonical-2333212022122300-2223311030122320-1103011303031211-3101200132212300-3301203200302111-3330021231032222-2322131111202030-2311031100032031"></a>

<a id="canonical-3000231221303301-1002031133133132-2000300322330221-3030131022111011-2130313133130133-3322120331120022-1331232002022011-2032132033102100"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the EnhancedFirewallPolicy.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2012020210033331-1222233301323013-2022211321222231-3122002300103002-1230303023310111-3303120021302301-0233012023111202-3332330233300230"></a>

<a id="canonical-1233021323203212-0230321313233122-0032122200300132-0111100033311102-1202030021312222-1113230013301330-3003332323312320-1011023131033223"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2303332011002232-2101330202212130-2020212202022020-0012321231221233-1210133133001012-0212011133012302-0312113132332330-1331310002301233"></a>

<a id="canonical-0211311023012331-2303132012310101-1310110033030201-0331020302110110-0113120002223010-0323231232302303-3333221100322320-2130020010221230"></a>

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

<a id="canonical-0301011330130312-2101010033031222-1003223320020131-3303213031022200-0121213123110000-2322321112121003-2210221003010313-1220011230123213"></a>

<a id="canonical-3013313200102100-1111112002321332-0133033003012232-3130301012121322-0031002001020300-1113010312111022-1130003133123200-3121321020120100"></a>

#### `name` property

Type: `"string"`. Required.

Name of the EnhancedFirewallPolicy.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1300021221302002-3123023331030022-2223031223033033-0331132221132111-1100101212210203-3021233023323123-0101030220320120-0002133211321021"></a>

<a id="canonical-0030132332232333-0302210012103333-2210220202333221-1233323223011110-3033110103311320-3313322023122322-2212101001103111-1023211132101322"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the EnhancedFirewallPolicy exists.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012): complete subsection reference.

<a id="canonical-0310211013322333-0010301101032312-3022121323121211-1122031211311321-2031321003222301-1331220000321323-3202331203312330-1010033023210212"></a>

### All schema paths for `xcsh_enhanced_firewall_policy`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allow_all` | [allow_all](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3203330213000211-3210131100132321-1131131203203320-0102233130102030-0113322100032012-2012301300033323-1023211001331012-1202131311033122) |
| `allowed_destinations` | [allowed_destinations](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3110211310222302-2000032231222232-3203331300331000-0032203001121203-3220033312132132-2303001232301102-2220122220303121-1000013123103003) |
| `allowed_destinations.prefix` | [allowed_destinations.prefix](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3232330322023332-0103333103110102-3201320303200012-1301302332030221-1322021012110120-3221233101033003-3221013220212301-0230120112102113) |
| `allowed_sources` | [allowed_sources](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3320302111113333-2301131330230320-3320013221033223-3302230333230320-3331311102111131-1111213133122110-0012001213132313-3110121321213332) |
| `allowed_sources.prefix` | [allowed_sources.prefix](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1001322220321221-3011020221010210-0200003000233230-2000300130031310-2101310301033213-3313213230231211-3311233331112302-3233110222121021) |
| `annotations` | [annotations](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0232303010003103-1212210123202020-1030313023333203-1333122301010321-0131022312001333-1120331303210321-3021020330023120-2301221131323233) |
| `denied_destinations` | [denied_destinations](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3133010313333001-1232131200330232-3031020201321101-2010002121213321-0123011223023232-2330113131322023-0213021100110320-2000310231303131) |
| `denied_destinations.prefix` | [denied_destinations.prefix](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1012123331330231-1102123002331013-2020203323000113-1210320011132302-1002111223320132-1131213030323021-1302023030233121-0231201210121011) |
| `denied_sources` | [denied_sources](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2330221131222210-3313311020012330-3300301231010120-0121131233320232-3222223223221323-2033202011322112-0302321230033201-2011133010220213) |
| `denied_sources.prefix` | [denied_sources.prefix](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1202120011033221-0221102132013100-2033222112202130-2010321301010111-3323311121110321-3130100202110011-1333201012100302-3132213321102120) |
| `deny_all` | [deny_all](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0121203302210010-1111122012321022-3223230312132022-0001012122330133-3232213323013012-2312230013130122-0213202211331231-1112002201122330) |
| `description` | [description](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2333212022122300-2223311030122320-1103011303031211-3101200132212300-3301203200302111-3330021231032222-2322131111202030-2311031100032031) |
| `id` | [ID](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2012020210033331-1222233301323013-2022211321222231-3122002300103002-1230303023310111-3303120021302301-0233012023111202-3332330233300230) |
| `labels` | [labels](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2303332011002232-2101330202212130-2020212202022020-0012321231221233-1210133133001012-0212011133012302-0312113132332330-1331310002301233) |
| `name` | [name](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0301011330130312-2101010033031222-1003223320020131-3303213031022200-0121213123110000-2322321112121003-2210221003010313-1220011230123213) |
| `namespace` | [namespace](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1300021221302002-3123023331030022-2223031223033033-0331132221132111-1100101212210203-3021233023323123-0101030220320120-0002133211321021) |
| `rule_list` | [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1303000032231103-3331013011132122-3013230230112102-3330200300332031-3201301300222223-3001033332231132-2202110322320331-2311131322210131) |
| `rule_list.rules` | [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1001213110323100-0132132131200023-3323312322310012-0101332202112231-3321131303021010-3110110021113300-1102321010031312-3110200001323332) |
| `rule_list.rules.advanced_action` | [rule_list.rules.advanced_action](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2001123323322313-0002021030133223-3020112232100112-1111322230312033-0222320020230302-0033100300130232-2331313100311310-3222221013011221) |
| `rule_list.rules.advanced_action.action` | [rule_list.rules.advanced_action.action](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1213021013031201-3310320033101033-0031200312231300-1012132201120023-2133003033122200-2210202001012112-2033302130320111-3322223033201233) |
| `rule_list.rules.all_destinations` | [rule_list.rules.all_destinations](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0002313231122132-2321133102220332-2201222213120012-1130130321332331-2221321022301113-1003103132013331-3130312121111002-2100000003003000) |
| `rule_list.rules.all_sli_vips` | [rule_list.rules.all_sli_vips](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3331011332212010-2120300231131320-3101112123310310-3001231133330231-2221123030211301-2221131121312001-1203111033221031-1131233222011100) |
| `rule_list.rules.all_slo_vips` | [rule_list.rules.all_slo_vips](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2002300132001113-1220031012100312-2133221010130213-0103032003022020-3011211123331320-2333322132011012-1022131022033223-2131322010012333) |
| `rule_list.rules.all_sources` | [rule_list.rules.all_sources](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3121110231101121-1113212012231123-0220010203133303-3112122103313313-0101331310212211-2320333303332332-3010231022212111-2130201000311112) |
| `rule_list.rules.all_tcp_traffic` | [rule_list.rules.all_tcp_traffic](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3230221310231300-1331103323303100-3022311010301333-3310023121302033-3132232101322113-1120010322231003-3333302332131311-1011013003021210) |
| `rule_list.rules.all_traffic` | [rule_list.rules.all_traffic](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1102131123333311-2011221201111323-0321002213023202-2332222110100221-3002221120011102-3333213112202101-0102021330221212-1332131132302121) |
| `rule_list.rules.all_udp_traffic` | [rule_list.rules.all_udp_traffic](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1313223031121020-1033200102021213-0231231110132130-3101222233312022-1222233232320111-0202011003212231-2333222013221300-3232030023100231) |
| `rule_list.rules.allow` | [rule_list.rules.allow](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2020303301222332-2030132222221010-3131200223221032-2312310012300013-1033012023130223-2011123000010233-0300230010332201-1210022313303201) |
| `rule_list.rules.applications` | [rule_list.rules.applications](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2021130303201010-0203331203332032-0200110200032021-1020200121100003-1332211211112211-3000210100303001-1001231201002033-1110121313120313) |
| `rule_list.rules.applications.applications` | [rule_list.rules.applications.applications](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1130230333222220-3312130201002313-1010332103330203-2313310120233120-2021311102101212-0130303302200312-3001101212233300-2030003310133023) |
| `rule_list.rules.deny` | [rule_list.rules.deny](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2213233120232231-0331323212222201-0030112320210322-1202312203230033-3202210103101130-0102030203102013-0031032201231220-1201301132111203) |
| `rule_list.rules.destination_aws_vpc_ids` | [rule_list.rules.destination_aws_vpc_ids](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3020312001320011-2201202102210132-2231221312213212-2321102020221002-3100311023002333-3032322203031103-2200313213023330-1112010232122222) |
| `rule_list.rules.destination_aws_vpc_ids.vpc_id` | [rule_list.rules.destination_aws_vpc_ids.vpc_id](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3201003021012301-2313231130210223-1211123321003321-2201200020021121-0003121010331013-2201121220212132-3000123112123330-2222211300123101) |
| `rule_list.rules.destination_ip_prefix_set` | [rule_list.rules.destination_ip_prefix_set](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0302210110120210-1032013100322311-0213310133212311-2023011103112103-0231310033322002-2303120012112011-2312212211011233-0021310102311132) |
| `rule_list.rules.destination_ip_prefix_set.ref` | [rule_list.rules.destination_ip_prefix_set.ref](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0131000021323102-2332031012213201-1201101000210110-1313120322300321-2130122012311000-0021203133200122-1302101123220123-2122130103130333) |
| `rule_list.rules.destination_ip_prefix_set.ref.kind` | [rule_list.rules.destination_ip_prefix_set.ref.kind](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2111103100110103-1212032101312011-3322321100131111-0113013322303223-3300132323022102-1020333310202323-3233022112213332-0131001311121030) |
| `rule_list.rules.destination_ip_prefix_set.ref.name` | [rule_list.rules.destination_ip_prefix_set.ref.name](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0100033020001200-3223300031311323-1311133320010321-3002132003322303-1110120302113123-1230300232302301-0202132320233021-1311033213010210) |
| `rule_list.rules.destination_ip_prefix_set.ref.namespace` | [rule_list.rules.destination_ip_prefix_set.ref.namespace](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3022332023021023-0330021032233320-0320130010232002-1013120320323212-3303212223003111-1123103223202221-2101031010221311-0130211233300212) |
| `rule_list.rules.destination_ip_prefix_set.ref.tenant` | [rule_list.rules.destination_ip_prefix_set.ref.tenant](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0221313002213212-1112110012012300-3200022321213110-1023300122320231-3322013132200122-0300201222303031-3331210031103102-0032213111110233) |
| `rule_list.rules.destination_ip_prefix_set.ref.uid` | [rule_list.rules.destination_ip_prefix_set.ref.uid](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0311232323332223-3222221312133031-2110322211230222-0311213102110223-2210323221213113-1003133330023323-1110001011013323-1303313311322321) |
| `rule_list.rules.destination_label_selector` | [rule_list.rules.destination_label_selector](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0302322003030101-1322211002120330-3330111320122130-3322021222231130-3203100233310231-3032211333233330-0223023122021132-1013300320313212) |
| `rule_list.rules.destination_label_selector.expressions` | [rule_list.rules.destination_label_selector.expressions](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2022121101211033-2312302122013221-2131300233130121-1311132330120033-3231101230003121-0103031233100311-2212310311301213-0113213211112001) |
| `rule_list.rules.destination_prefix_list` | [rule_list.rules.destination_prefix_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2020020022101033-1000300333122032-2123321332201302-0310303002021021-3322232130303022-2130200230233132-0312310102312033-1310331220222300) |
| `rule_list.rules.destination_prefix_list.prefixes` | [rule_list.rules.destination_prefix_list.prefixes](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3132323103233102-2133011032331223-0231303023011232-0230111101110203-3302121220323311-1333100232212332-0012132111213313-2122120022013010) |
| `rule_list.rules.insert_service` | [rule_list.rules.insert_service](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2301222121233112-2320000333313132-1321130333300023-1011001000031003-1212201012031132-3212233023230321-1231110300212211-2123213232132220) |
| `rule_list.rules.insert_service.nfv_service` | [rule_list.rules.insert_service.nfv_service](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1130132210323022-2202102323100110-1110131102132000-1003022312002011-2003231111131002-0300202032230202-1102032123230310-3311132111123000) |
| `rule_list.rules.insert_service.nfv_service.name` | [rule_list.rules.insert_service.nfv_service.name](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3201222013210303-2313330012122211-2200232210330200-1001032123200302-2230030330302030-3111012003120102-0231321332211232-1032012110131212) |
| `rule_list.rules.insert_service.nfv_service.namespace` | [rule_list.rules.insert_service.nfv_service.namespace](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1322110131310132-3000102201222331-1213011132303120-0003200002223013-1220132030101130-3221010111333301-2121312001201311-2023321323200130) |
| `rule_list.rules.insert_service.nfv_service.tenant` | [rule_list.rules.insert_service.nfv_service.tenant](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0332223301210331-3123333220121300-2022213213112330-1011203300302112-0012031021130131-1131322131111122-1102000122233011-0001133033302322) |
| `rule_list.rules.inside_destinations` | [rule_list.rules.inside_destinations](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0322110200202030-3112021333113100-2110230130231032-0102113303202122-0003213201001012-0022333322132112-1323031030221100-1201232120132023) |
| `rule_list.rules.inside_sources` | [rule_list.rules.inside_sources](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0033000212020112-1210220112213101-3203221130313301-0021000030133031-1030010122323132-3333010333122012-0231300310130131-1303223211131021) |
| `rule_list.rules.label_matcher` | [rule_list.rules.label_matcher](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1222123310123123-1003322020230300-3322332022202332-1101331330230001-0221101132211211-0201212131021010-3122313231322323-2202220311213223) |
| `rule_list.rules.label_matcher.keys` | [rule_list.rules.label_matcher.keys](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2213223131320103-3323010302100222-3001030111202113-3000303020103013-2223010302120300-0313021003021132-0011233010201101-2100031222332333) |
| `rule_list.rules.metadata` | [rule_list.rules.metadata](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0111232203213322-1230222120032121-3220233233003100-2013213220301101-0222133101213000-2312201332011033-0121212023000320-3221201122013303) |
| `rule_list.rules.metadata.description_spec` | [rule_list.rules.metadata.description_spec](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2323022123110220-3332313230201002-0311113121100112-2221012220112022-3213010020020012-1133113012030000-2330001303302022-0321302323031111) |
| `rule_list.rules.metadata.name` | [rule_list.rules.metadata.name](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1001001230322021-1322013233223313-3103020132102323-3322232002322332-3030300210300312-3011301133303202-3233101220001300-1202102103300013) |
| `rule_list.rules.outside_destinations` | [rule_list.rules.outside_destinations](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1330222020023201-1101022233003121-1012000000000001-1032313310101110-0233102020222113-0002220332002222-3222323311323233-0222301012131022) |
| `rule_list.rules.outside_sources` | [rule_list.rules.outside_sources](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1013331211213330-0002000210011223-3311112210301022-1330011322103330-2223330302322130-0322323121313111-0330222310320301-2132012301333123) |
| `rule_list.rules.protocol_port_range` | [rule_list.rules.protocol_port_range](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0123321303223122-0032010033121221-3100331311022322-0030123213311330-1011230033300133-3001210112120313-3100030323320313-1231132133233113) |
| `rule_list.rules.protocol_port_range.port_ranges` | [rule_list.rules.protocol_port_range.port_ranges](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2313013333200231-3101133332001000-1032331323103103-2111200213220321-0023130112102311-2331111203323103-2321010200202030-1303300312322223) |
| `rule_list.rules.protocol_port_range.protocol` | [rule_list.rules.protocol_port_range.protocol](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1131130031201310-1123300101112302-1003232022031222-0222322110330331-2112122330012231-3103201330311203-2001310322220023-1200200001303311) |
| `rule_list.rules.source_aws_vpc_ids` | [rule_list.rules.source_aws_vpc_ids](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2212121000313110-2302310112011032-0113303031020021-0200211110101232-2013130113123213-0103333111032123-3221102002002222-3111012032132313) |
| `rule_list.rules.source_aws_vpc_ids.vpc_id` | [rule_list.rules.source_aws_vpc_ids.vpc_id](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1213030122202103-2331122001021113-2010123222202300-3323130101233003-0321200130203110-3330321321220011-0230320210322021-0231003303332101) |
| `rule_list.rules.source_ip_prefix_set` | [rule_list.rules.source_ip_prefix_set](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3311001210322232-2001002302210030-1122332011222221-2003003101023130-0311121010103020-1020201110000311-1331021133323012-2012301132001300) |
| `rule_list.rules.source_ip_prefix_set.ref` | [rule_list.rules.source_ip_prefix_set.ref](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1310312311112330-0300101303131320-3003111210121220-3300201221102103-0010322203000211-2000011230230302-3113322131312110-3123202131001022) |
| `rule_list.rules.source_ip_prefix_set.ref.kind` | [rule_list.rules.source_ip_prefix_set.ref.kind](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3221111010220000-1002102101310311-2131023121111031-2300120033112101-2132021100130201-3012323123110123-1221302020200223-0003023222130113) |
| `rule_list.rules.source_ip_prefix_set.ref.name` | [rule_list.rules.source_ip_prefix_set.ref.name](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2201200230303000-2120202313031122-3231323311100221-0113121313201122-1010302203222322-2333332220120232-3201133120333023-1300200230132302) |
| `rule_list.rules.source_ip_prefix_set.ref.namespace` | [rule_list.rules.source_ip_prefix_set.ref.namespace](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2123311322120110-2132011200122322-1313002321022123-0121003023023102-1103332303222330-3000133021231000-1203333020302003-3303012311201331) |
| `rule_list.rules.source_ip_prefix_set.ref.tenant` | [rule_list.rules.source_ip_prefix_set.ref.tenant](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3203230132130212-2102033300133202-0032021102323120-1012003313121322-2103332322012003-2210220212002023-0000000023121021-0202320311123120) |
| `rule_list.rules.source_ip_prefix_set.ref.uid` | [rule_list.rules.source_ip_prefix_set.ref.uid](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1300123303000201-3230300010020201-0033302213131000-2303311010022322-0122322001202232-3110321010022012-2321330123330212-1321231120312233) |
| `rule_list.rules.source_label_selector` | [rule_list.rules.source_label_selector](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3233103021113000-3300202332111230-3313210211220023-3221101002313000-2132110133010300-3033232220230311-2102012333133103-1102313220222203) |
| `rule_list.rules.source_label_selector.expressions` | [rule_list.rules.source_label_selector.expressions](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2200111101111213-1202003211103131-1003130131120223-2331231110020332-2120022000200020-1331103012310022-1332032121210113-1101312110021300) |
| `rule_list.rules.source_prefix_list` | [rule_list.rules.source_prefix_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2300230031233203-2133002220310102-0120032012310201-0302332023210002-3112313311232332-1112311112202213-1103310311122121-0202022120221331) |
| `rule_list.rules.source_prefix_list.prefixes` | [rule_list.rules.source_prefix_list.prefixes](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2020201110132331-2303333013001020-1012213010113201-3113113023230032-2200300212333210-3000100022310110-1131203132330101-3022112333033111) |

<a id="canonical-3033303330013011-3303131020130023-0120301102013120-2331313303103320-0103123003302022-2020121003023013-0331012313330210-2223122222322000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_all` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- allow_all

<a id="canonical-3203330213000211-3210131100132321-1131131203203320-0102233130102030-0113322100032012-2012301300033323-1023211001331012-1202131311033122"></a>

Type: `["object", {}]`. Computed.

\[OneOf: allow\_all, allowed\_destinations, allowed\_sources, denied\_destinations, denied\_sources,
deny\_all, rule\_list\] Enable this option. Defaults to \`map\[\]\`. Server applies default when
omitted.

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

- [allow_all](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3203330213000211-3210131100132321-1131131203203320-0102233130102030-0113322100032012-2012301300033323-1023211001331012-1202131311033122)
- [allowed_destinations](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3110211310222302-2000032231222232-3203331300331000-0032203001121203-3220033312132132-2303001232301102-2220122220303121-1000013123103003)
- [allowed_sources](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3320302111113333-2301131330230320-3320013221033223-3302230333230320-3331311102111131-1111213133122110-0012001213132313-3110121321213332)
- [denied_destinations](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3133010313333001-1232131200330232-3031020201321101-2010002121213321-0123011223023232-2330113131322023-0213021100110320-2000310231303131)
- [denied_sources](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2330221131222210-3313311020012330-3300301231010120-0121131233320232-3222223223221323-2033202011322112-0302321230033201-2011133010220213)
- [deny_all](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0121203302210010-1111122012321022-3223230312132022-0001012122330133-3232213323013012-2312230013130122-0213202211331231-1112002201122330)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1303000032231103-3331013011132122-3013230230112102-3330200300332031-3201301300222223-3001033332231132-2202110322320331-2311131322210131)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1121223222222033-1321103111123211-2322303130112211-1320223300030321-3330303232231310-3031333003200220-2001132333000233-3120233221311201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allowed_destinations` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- allowed_destinations

<a id="canonical-3110211310222302-2000032231222232-3203331300331000-0032203001121203-3220033312132132-2303001232301102-2220122220303121-1000013123103003"></a>

Type: `"single"`. Computed.

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

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

<a id="canonical-1101133120022303-3112101001323100-1311310013232101-3302322131203101-3100201033230222-1110203031303111-3033333313122221-2222322212321013"></a>

### Direct properties for `allowed_destinations`

<a id="canonical-3232330322023332-0103333103110102-3201320303200012-1301302332030221-1322021012110120-3221233101033003-3221013220212301-0230120112102113"></a>

#### `allowed_destinations.prefix` property

Type: `["list", "string"]`. Computed.

IP Address prefix in string format. String must contain both prefix and prefix-length.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-0310120130032220-3210102231321012-2322312302113320-0123301300100020-0013321230021233-1002110311102222-3211131333310303-3000130033031121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allowed_sources` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- allowed_sources

<a id="canonical-3320302111113333-2301131330230320-3320013221033223-3302230333230320-3331311102111131-1111213133122110-0012001213132313-3110121321213332"></a>

Type: `"single"`. Computed.

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

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

<a id="canonical-3113113331210300-0101000312201000-0110313021132023-3332302122311333-1013100003103330-3023000123301333-3032333311112213-0010303120003020"></a>

### Direct properties for `allowed_sources`

<a id="canonical-1001322220321221-3011020221010210-0200003000233230-2000300130031310-2101310301033213-3313213230231211-3311233331112302-3233110222121021"></a>

#### `allowed_sources.prefix` property

Type: `["list", "string"]`. Computed.

IP Address prefix in string format. String must contain both prefix and prefix-length.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-0321100103231102-3020030330213210-2203000323112003-1130001202103022-2233110030031311-0222201211232332-2011313100032013-3022022111322323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `denied_destinations` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- denied_destinations

<a id="canonical-3133010313333001-1232131200330232-3031020201321101-2010002121213321-0123011223023232-2330113131322023-0213021100110320-2000310231303131"></a>

Type: `"single"`. Computed.

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

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

<a id="canonical-1210202112122221-2211313332321220-1110012221232010-3230100000102010-3110122023110123-3232210103311001-0331030223232132-1232202100122231"></a>

### Direct properties for `denied_destinations`

<a id="canonical-1012123331330231-1102123002331013-2020203323000113-1210320011132302-1002111223320132-1131213030323021-1302023030233121-0231201210121011"></a>

#### `denied_destinations.prefix` property

Type: `["list", "string"]`. Computed.

IP Address prefix in string format. String must contain both prefix and prefix-length.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-3210321321221232-1330102210103130-1121232320102122-2210300100320211-3233302002001113-2031213013120212-1210221121011303-1102113300303202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `denied_sources` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- denied_sources

<a id="canonical-2330221131222210-3313311020012330-3300301231010120-0121131233320232-3222223223221323-2033202011322112-0302321230033201-2011133010220213"></a>

Type: `"single"`. Computed.

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

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

<a id="canonical-0231333012033033-3131012030022023-3001300323023300-2333013313303301-0033101113030201-2132221000012002-0232123001122300-0002003311010031"></a>

### Direct properties for `denied_sources`

<a id="canonical-1202120011033221-0221102132013100-2033222112202130-2010321301010111-3323311121110321-3130100202110011-1333201012100302-3132213321102120"></a>

#### `denied_sources.prefix` property

Type: `["list", "string"]`. Computed.

IP Address prefix in string format. String must contain both prefix and prefix-length.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-2221303231002233-3313021213301220-3211322103313302-1020330122012331-3323333321023320-2321002021313111-0211022201312220-0122131303132031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `deny_all` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- deny_all

<a id="canonical-0121203302210010-1111122012321022-3223230312132022-0001012122330133-3232213323013012-2312230013130122-0213202211331231-1112002201122330"></a>

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

<a id="canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- rule_list

<a id="canonical-1303000032231103-3331013011132122-3013230230112102-3330200300332031-3201301300222223-3001033332231132-2202110322320331-2311131322210131"></a>

Type: `"single"`. Computed.

Custom Enhanced Firewall Policy Rules. Custom Enhanced Firewall Policy Rules.

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

<a id="canonical-1131122000331010-3030333332302132-3321131103002303-0320232321323302-0100003131102313-3031012231133212-0301121022220123-3213113230223022"></a>

### Direct properties for `rule_list`

- [rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331): complete subsection reference.

<a id="canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012)
- rule_list.rules

<a id="canonical-1001213110323100-0132132131200023-3323312322310012-0101332202112231-3321131303021010-3110110021113300-1102321010031312-3110200001323332"></a>

Type: `"list"`. Computed.

Ordered List of Enhanced Firewall Policy Rules.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-1220323002223313-3212023033231310-1313323012102310-3302323013201110-3012203012102232-3020113022030132-2202323013002121-2321310130221230"></a>

### Direct properties for `rule_list.rules`

- [advanced_action](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3210330312201022-1312211112203320-1330221132101013-3220113330032321-0132010310003322-2233113033331302-3323202001221312-0002200303213131): complete subsection reference.

- [all_destinations](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0130323122311203-1332200231221200-2121103231231310-2213133003111331-1011321323023211-0311103300033332-1310012203131130-3010033031331231): complete subsection reference.

- [all_sli_vips](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3230230310311020-3321301201303100-2131202032013130-2203330001213010-2030210222121031-3013111332212120-0222201301303113-0000122332203303): complete subsection reference.

- [all_slo_vips](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1121001100021132-2120323001101331-0102200001332102-1330113110300033-3010303200233230-1121331102133222-1330012000031003-3012033111000232): complete subsection reference.

- [all_sources](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0323230103331203-1033011130030310-2312203320001032-1111033110001113-3212310023130120-0111311200322022-1133223032222131-2100102103100122): complete subsection reference.

- [all_tcp_traffic](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0011012213010101-0013203211130303-1223312300302211-1111133101011302-3312232113011230-2020021032000200-2203133302220330-2123131120221132): complete subsection reference.

- [all_traffic](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0030031230021223-3111210320321330-0233012113003333-2231211002210322-0130223312231332-2000112031302332-2330210230232200-1302312013312023): complete subsection reference.

- [all_udp_traffic](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3332033332303020-1000103300121303-3232011312013322-2203130032313022-1201101101321320-0231112212010321-1021023302100031-0232011022200231): complete subsection reference.

- [allow](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1303301310011003-2020023100102102-0323030013223113-2121220101123020-1311123331010233-2233112212230131-3210110321313001-2111020312333303): complete subsection reference.

- [applications](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2112031303132012-2220221022221021-1013202111321103-2103302022201013-1212120103103302-1020303301102310-3312233120130002-3231130031122110): complete subsection reference.

- [deny](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2200220201000311-3321211213031231-1311102010230112-0013301331313300-2330332232212031-1120230131212310-1033332101232203-1002201133031332): complete subsection reference.

- [destination_aws_vpc_ids](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1113032202220003-0210302003301211-3002002020103310-2003311013012120-1103200201230130-0311030332023133-0230122212332112-0332133320122212): complete subsection reference.

- [destination_ip_prefix_set](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1300023031323000-3301102010113002-3010003233100222-3121002222130112-3122212003213011-2133330230012031-3203121012202323-1123120032032021): complete subsection reference.

- [destination_label_selector](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3003123312131023-2312203311031330-2332011130331322-3010320103102222-2121330013010111-2312201223212312-3033320122323032-3123233233013010): complete subsection reference.

- [destination_prefix_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0202010333003233-1200030133123221-1312322333120332-3330312222001230-1110131220211332-0012211013100311-3222031211200000-2210011212001330): complete subsection reference.

- [insert_service](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1310013102000131-0322022303222123-1311213313233023-3013022312020200-1003112333212101-1120120311101303-3132222211200221-3012201032111123): complete subsection reference.

- [inside_destinations](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1333103023001323-2131103330332012-1022223103321111-3211123110100010-0230330303021330-2001232133321201-1322321301302210-1023313220312211): complete subsection reference.

- [inside_sources](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1302112120133231-3103131110123330-3322131110123200-3012001030213001-0212233030322132-1301311123121332-1311330232311232-0233030201030030): complete subsection reference.

- [label_matcher](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1102313302122332-3103303303311030-1312320112033333-1322210030101100-2330310321010220-0200013021130231-1122231303302013-1232203200232000): complete subsection reference.

- [metadata](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0321021231333001-0022230203210300-1002232012330002-3022111031230302-3103032313102323-2323130032322131-0001313111001332-1003112213103303): complete subsection reference.

- [outside_destinations](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0233033010303013-3213003233131011-2023031031100121-1211222010230211-2120031331311133-2333012030322021-0310020130103012-0131111012332101): complete subsection reference.

- [outside_sources](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0331001112202220-0222020210122001-2200031112222223-1232013223331332-0301001312220303-2011302001200322-0312213200131101-2122320102311110): complete subsection reference.

- [protocol_port_range](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1010210030311031-2302010003102121-3030001023003210-0303201120131000-3120000013333020-1222300030311223-0112130221003021-0312213010010203): complete subsection reference.

- [source_aws_vpc_ids](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1303330131133222-2011312300131102-3103333300023021-3113302301021331-2022112113023002-1302120120322100-0312301222031211-3331302220131002): complete subsection reference.

- [source_ip_prefix_set](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0200023231312311-1323123232302021-1000112103130300-0000330202320113-0122021212200300-0332113310010120-0323322200311330-0300303210223333): complete subsection reference.

- [source_label_selector](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0011023231130030-2211332223100010-3323313001301320-2003202111312011-0202033201001323-0003021023122232-0301132123113023-3323313201122300): complete subsection reference.

- [source_prefix_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0232011221020021-0110002001020030-0300123202033112-2132210231313033-2301100230222303-2302213313102120-3320330233302223-0013032000233310): complete subsection reference.

<a id="canonical-3210330312201022-1312211112203320-1330221132101013-3220113330032321-0132010310003322-2233113033331302-3323202001221312-0002200303213131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.advanced_action` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331)
- rule_list.rules.advanced_action

<a id="canonical-2001123323322313-0002021030133223-3020112232100112-1111322230312033-0222320020230302-0033100300130232-2331313100311310-3222221013011221"></a>

Type: `"single"`. Computed.

Network Policy Rule Advanced Action provides additional OPTIONS along with RuleAction and
PBRRuleAction.

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

<a id="canonical-3013200021311312-1111320111022031-3033320133301200-2123220103112110-0013312112021123-1333102313110322-3312233023323310-3121233013020003"></a>

### Direct properties for `rule_list.rules.advanced_action`

<a id="canonical-1213021013031201-3310320033101033-0031200312231300-1012132201120023-2133003033122200-2210202001012112-2033302130320111-3322223033201233"></a>

#### `rule_list.rules.advanced_action.action` property

Type: `"string"`. Computed.

\[Enum: NOLOG|LOG\] Choice to choose logging or no logging This works together with option selected
via NetworkPolicyRuleAction or any other action specified x-. Possible values are \`NOLOG\`,
\`LOG\`. Defaults to \`NOLOG\`.

Additional upstream details:

Choice to choose logging or no logging This works together with option selected via
NetworkPolicyRuleAction or any other action specified x-example: (No Selection in
NetworkPolicyRuleAction + AdvancedAction as LOG) = LOG Only, (ALLOW/DENY in NetworkPolicyRuleAction
&#8203;+ AdvancedAction as LOG) = Log and Allow/Deny, (ALLOW/DENY in NetworkPolicyRuleAction + NOLOG in
AdvancedAction) = Allow/Deny with no log

Don't sample the traffic hitting the rule Sample the traffic hitting the rule.

Receipt-pinned upstream constraints:

```json
{
  "default": "NOLOG",
  "enum": [
    "NOLOG",
    "LOG"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0130323122311203-1332200231221200-2121103231231310-2213133003111331-1011321323023211-0311103300033332-1310012203131130-3010033031331231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.all_destinations` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331)
- rule_list.rules.all_destinations

<a id="canonical-0002313231122132-2321133102220332-2201222213120012-1130130321332331-2221321022301113-1003103132013331-3130312121111002-2100000003003000"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all destinations.

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

<a id="canonical-3230230310311020-3321301201303100-2131202032013130-2203330001213010-2030210222121031-3013111332212120-0222201301303113-0000122332203303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.all_sli_vips` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331)
- rule_list.rules.all_sli_vips

<a id="canonical-3331011332212010-2120300231131320-3101112123310310-3001231133330231-2221123030211301-2221131121312001-1203111033221031-1131233222011100"></a>

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

<a id="canonical-1121001100021132-2120323001101331-0102200001332102-1330113110300033-3010303200233230-1121331102133222-1330012000031003-3012033111000232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.all_slo_vips` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331)
- rule_list.rules.all_slo_vips

<a id="canonical-2002300132001113-1220031012100312-2133221010130213-0103032003022020-3011211123331320-2333322132011012-1022131022033223-2131322010012333"></a>

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

<a id="canonical-0323230103331203-1033011130030310-2312203320001032-1111033110001113-3212310023130120-0111311200322022-1133223032222131-2100102103100122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.all_sources` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331)
- rule_list.rules.all_sources

<a id="canonical-3121110231101121-1113212012231123-0220010203133303-3112122103313313-0101331310212211-2320333303332332-3010231022212111-2130201000311112"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all sources.

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

<a id="canonical-0011012213010101-0013203211130303-1223312300302211-1111133101011302-3312232113011230-2020021032000200-2203133302220330-2123131120221132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.all_tcp_traffic` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331)
- rule_list.rules.all_tcp_traffic

<a id="canonical-3230221310231300-1331103323303100-3022311010301333-3310023121302033-3132232101322113-1120010322231003-3333302332131311-1011013003021210"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all tcp traffic.

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

<a id="canonical-0030031230021223-3111210320321330-0233012113003333-2231211002210322-0130223312231332-2000112031302332-2330210230232200-1302312013312023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.all_traffic` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331)
- rule_list.rules.all_traffic

<a id="canonical-1102131123333311-2011221201111323-0321002213023202-2332222110100221-3002221120011102-3333213112202101-0102021330221212-1332131132302121"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all traffic.

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

<a id="canonical-3332033332303020-1000103300121303-3232011312013322-2203130032313022-1201101101321320-0231112212010321-1021023302100031-0232011022200231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.all_udp_traffic` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331)
- rule_list.rules.all_udp_traffic

<a id="canonical-1313223031121020-1033200102021213-0231231110132130-3101222233312022-1222233232320111-0202011003212231-2333222013221300-3232030023100231"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all udp traffic.

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

<a id="canonical-1303301310011003-2020023100102102-0323030013223113-2121220101123020-1311123331010233-2233112212230131-3210110321313001-2111020312333303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.allow` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331)
- rule_list.rules.allow

<a id="canonical-2020303301222332-2030132222221010-3131200223221032-2312310012300013-1033012023130223-2011123000010233-0300230010332201-1210022313303201"></a>

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

<a id="canonical-2112031303132012-2220221022221021-1013202111321103-2103302022201013-1212120103103302-1020303301102310-3312233120130002-3231130031122110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.applications` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331)
- rule_list.rules.applications

<a id="canonical-2021130303201010-0203331203332032-0200110200032021-1020200121100003-1332211211112211-3000210100303001-1001231201002033-1110121313120313"></a>

Type: `"single"`. Computed.

Configuration parameter for applications.

Additional upstream details:

Application protocols like HTTP, SNMP.

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

<a id="canonical-0101101023133030-2002101200013103-0312222301232202-1003220220223103-2002202023112020-3332101110330301-1203022111331222-2213023100233001"></a>

### Direct properties for `rule_list.rules.applications`

<a id="canonical-1130230333222220-3312130201002313-1010332103330203-2313310120233120-2021311102101212-0130303302200312-3001101212233300-2030003310133023"></a>

#### `rule_list.rules.applications.applications` property

Type: `["list", "string"]`. Computed.

\[Enum: APPLICATION\_HTTP|APPLICATION\_HTTPS|APPLICATION\_SNMP|APPLICATION\_DNS\] Application
Protocols. Application protocols like HTTP, SNMP. Possible values are \`APPLICATION\_HTTP\`,
\`APPLICATION\_HTTPS\`, \`APPLICATION\_SNMP\`, \`APPLICATION\_DNS\`. Defaults to
\`APPLICATION\_HTTP\`.

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

<a id="canonical-2200220201000311-3321211213031231-1311102010230112-0013301331313300-2330332232212031-1120230131212310-1033332101232203-1002201133031332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.deny` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331)
- rule_list.rules.deny

<a id="canonical-2213233120232231-0331323212222201-0030112320210322-1202312203230033-3202210103101130-0102030203102013-0031032201231220-1201301132111203"></a>

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

<a id="canonical-1113032202220003-0210302003301211-3002002020103310-2003311013012120-1103200201230130-0311030332023133-0230122212332112-0332133320122212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.destination_aws_vpc_ids` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331)
- rule_list.rules.destination_aws_vpc_ids

<a id="canonical-3020312001320011-2201202102210132-2231221312213212-2321102020221002-3100311023002333-3032322203031103-2200313213023330-1112010232122222"></a>

Type: `"single"`. Computed.

Configuration parameter for destination aws vpc IDs.

Additional upstream details:

List of VPC Identifiers in AWS.

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

<a id="canonical-2301023131102013-2100122130102321-3303321333101322-3133110100301310-3100030011203103-1021133331301011-2030203321201012-3102230311200320"></a>

### Direct properties for `rule_list.rules.destination_aws_vpc_ids`

<a id="canonical-3201003021012301-2313231130210223-1211123321003321-2201200020021121-0003121010331013-2201121220212132-3000123112123330-2222211300123101"></a>

#### `rule_list.rules.destination_aws_vpc_ids.vpc_id` property

Type: `["list", "string"]`. Computed.

AWS VPC List. List of VPC Identifiers in AWS.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1300023031323000-3301102010113002-3010003233100222-3121002222130112-3122212003213011-2133330230012031-3203121012202323-1123120032032021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.destination_ip_prefix_set` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331)
- rule_list.rules.destination_ip_prefix_set

<a id="canonical-0302210110120210-1032013100322311-0213310133212311-2023011103112103-0231310033322002-2303120012112011-2312212211011233-0021310102311132"></a>

Type: `"single"`. Computed.

A list of references to ip\_prefix\_set objects.

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

<a id="canonical-2033210120201020-2203213232022021-2123000310103320-2233210313331220-2110000201130200-1331330002323111-2313031130033220-3222211000323212"></a>

### Direct properties for `rule_list.rules.destination_ip_prefix_set`

- [ref](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1103002301111020-0133232311300230-1223331202302312-1203333232222020-2012331031030220-3202100012021102-3230232230101203-2211132300230123): complete subsection reference.

<a id="canonical-1103002301111020-0133232311300230-1223331202302312-1203333232222020-2012331031030220-3202100012021102-3230232230101203-2211132300230123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.destination_ip_prefix_set.ref` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331)
- [rule_list.rules.destination_ip_prefix_set](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1300023031323000-3301102010113002-3010003233100222-3121002222130112-3122212003213011-2133330230012031-3203121012202323-1123120032032021)
- rule_list.rules.destination_ip_prefix_set.ref

<a id="canonical-0131000021323102-2332031012213201-1201101000210110-1313120322300321-2130122012311000-0021203133200122-1302101123220123-2122130103130333"></a>

Type: `"list"`. Computed.

A list of references to ip\_prefix\_set objects.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0302333031101133-2112300020021230-0301022023230332-1333323000303110-0121012111033133-2212010220022110-3012132121320120-0012021213213122"></a>

### Direct properties for `rule_list.rules.destination_ip_prefix_set.ref`

<a id="canonical-2111103100110103-1212032101312011-3322321100131111-0113013322303223-3300132323022102-1020333310202323-3233022112213332-0131001311121030"></a>

#### `rule_list.rules.destination_ip_prefix_set.ref.kind` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0100033020001200-3223300031311323-1311133320010321-3002132003322303-1110120302113123-1230300232302301-0202132320233021-1311033213010210"></a>

<a id="canonical-3323302121220301-1031303333313201-1102301130223303-3311013200232220-2003010032002310-2201120310313330-2330031120333113-3313120103323200"></a>

#### `rule_list.rules.destination_ip_prefix_set.ref.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3022332023021023-0330021032233320-0320130010232002-1013120320323212-3303212223003111-1123103223202221-2101031010221311-0130211233300212"></a>

<a id="canonical-0313103012000100-1220213011210301-0203312011023203-1322130213330232-2330120111311100-1013200130230123-2031000310231301-3231020032210331"></a>

#### `rule_list.rules.destination_ip_prefix_set.ref.namespace` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0221313002213212-1112110012012300-3200022321213110-1023300122320231-3322013132200122-0300201222303031-3331210031103102-0032213111110233"></a>

<a id="canonical-3303122131301232-3332021313201013-1203333300302223-0303123033130132-2321213133230331-1321301202020303-3002112010300223-3203112000303233"></a>

#### `rule_list.rules.destination_ip_prefix_set.ref.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0311232323332223-3222221312133031-2110322211230222-0311213102110223-2210323221213113-1003133330023323-1110001011013323-1303313311322321"></a>

<a id="canonical-1023112103103011-3030302031132332-2332003312212221-1032221122113221-3010212123201011-3222230331001132-0133221012233322-2230001002112203"></a>

#### `rule_list.rules.destination_ip_prefix_set.ref.uid` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3003123312131023-2312203311031330-2332011130331322-3010320103102222-2121330013010111-2312201223212312-3033320122323032-3123233233013010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.destination_label_selector` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331)
- rule_list.rules.destination_label_selector

<a id="canonical-0302322003030101-1322211002120330-3330111320122130-3322021222231130-3203100233310231-3032211333233330-0223023122021132-1013300320313212"></a>

Type: `"single"`. Computed.

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

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

<a id="canonical-0322210102213310-2213030220123101-1130032322202200-2102022120122310-2311113223113302-0030112332133001-2330301132013001-0212333032101031"></a>

### Direct properties for `rule_list.rules.destination_label_selector`

<a id="canonical-2022121101211033-2312302122013221-2131300233130121-1311132330120033-3231101230003121-0103031233100311-2212310311301213-0113213211112001"></a>

#### `rule_list.rules.destination_label_selector.expressions` property

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-0202010333003233-1200030133123221-1312322333120332-3330312222001230-1110131220211332-0012211013100311-3222031211200000-2210011212001330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.destination_prefix_list` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331)
- rule_list.rules.destination_prefix_list

<a id="canonical-2020020022101033-1000300333122032-2123321332201302-0310303002021021-3322232130303022-2130200230233132-0312310102312033-1310331220222300"></a>

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

<a id="canonical-1301010330201133-1130122011220222-1200031223122333-2013121213313111-0210110203303211-3003030220112100-1111131111201201-0010202023020111"></a>

### Direct properties for `rule_list.rules.destination_prefix_list`

<a id="canonical-3132323103233102-2133011032331223-0231303023011232-0230111101110203-3302121220323311-1333100232212332-0012132111213313-2122120022013010"></a>

#### `rule_list.rules.destination_prefix_list.prefixes` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1310013102000131-0322022303222123-1311213313233023-3013022312020200-1003112333212101-1120120311101303-3132222211200221-3012201032111123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.insert_service` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331)
- rule_list.rules.insert_service

<a id="canonical-2301222121233112-2320000333313132-1321130333300023-1011001000031003-1212201012031132-3212233023230321-1231110300212211-2123213232132220"></a>

Type: `"single"`. Computed.

Action to forward traffic to external service.

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

<a id="canonical-3322003230131012-1113013231301221-2303023312331332-3312033323033021-3133211323002021-3130020112313101-1302331320323321-1223312002112113"></a>

### Direct properties for `rule_list.rules.insert_service`

- [nfv_service](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1312331212331111-2320131012000211-1100333113212021-0200221110033323-0203313001033302-1303012210223231-0022001033021010-2010002302131102): complete subsection reference.

<a id="canonical-1312331212331111-2320131012000211-1100333113212021-0200221110033323-0203313001033302-1303012210223231-0022001033021010-2010002302131102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.insert_service.nfv_service` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331)
- [rule_list.rules.insert_service](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1310013102000131-0322022303222123-1311213313233023-3013022312020200-1003112333212101-1120120311101303-3132222211200221-3012201032111123)
- rule_list.rules.insert_service.nfv_service

<a id="canonical-1130132210323022-2202102323100110-1110131102132000-1003022312002011-2003231111131002-0300202032230202-1102032123230310-3311132111123000"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0100303122311231-0002131132032100-3001210222102331-0331131110003230-1211220101222032-1320122021231123-3011133023203030-0123132231230101"></a>

### Direct properties for `rule_list.rules.insert_service.nfv_service`

<a id="canonical-3201222013210303-2313330012122211-2200232210330200-1001032123200302-2230030330302030-3111012003120102-0231321332211232-1032012110131212"></a>

#### `rule_list.rules.insert_service.nfv_service.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1322110131310132-3000102201222331-1213011132303120-0003200002223013-1220132030101130-3221010111333301-2121312001201311-2023321323200130"></a>

<a id="canonical-3200030100012232-1223203102222113-2312321023300301-1311200111213210-0012020233010132-1103032221332313-0213101303130000-2003310301310333"></a>

#### `rule_list.rules.insert_service.nfv_service.namespace` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0332223301210331-3123333220121300-2022213213112330-1011203300302112-0012031021130131-1131322131111122-1102000122233011-0001133033302322"></a>

<a id="canonical-0010210212202332-2021313101023031-1133012310012022-2131021131112033-3221003300333120-3112223103020321-0222202331330330-0100200030122002"></a>

#### `rule_list.rules.insert_service.nfv_service.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1333103023001323-2131103330332012-1022223103321111-3211123110100010-0230330303021330-2001232133321201-1322321301302210-1023313220312211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.inside_destinations` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331)
- rule_list.rules.inside_destinations

<a id="canonical-0322110200202030-3112021333113100-2110230130231032-0102113303202122-0003213201001012-0022333322132112-1323031030221100-1201232120132023"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inside destinations.

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

<a id="canonical-1302112120133231-3103131110123330-3322131110123200-3012001030213001-0212233030322132-1301311123121332-1311330232311232-0233030201030030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.inside_sources` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331)
- rule_list.rules.inside_sources

<a id="canonical-0033000212020112-1210220112213101-3203221130313301-0021000030133031-1030010122323132-3333010333122012-0231300310130131-1303223211131021"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inside sources.

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

<a id="canonical-1102313302122332-3103303303311030-1312320112033333-1322210030101100-2330310321010220-0200013021130231-1122231303302013-1232203200232000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.label_matcher` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331)
- rule_list.rules.label_matcher

<a id="canonical-1222123310123123-1003322020230300-3322332022202332-1101331330230001-0221101132211211-0201212131021010-3122313231322323-2202220311213223"></a>

Type: `"single"`. Computed.

A label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

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

<a id="canonical-2203313101021200-2021031312220013-0330321210210321-2002202202313003-2022203000011233-0213232223230032-1031002122110011-1301212211102003"></a>

### Direct properties for `rule_list.rules.label_matcher`

<a id="canonical-2213223131320103-3323010302100222-3001030111202113-3000303020103013-2223010302120300-0313021003021132-0011233010201101-2100031222332333"></a>

#### `rule_list.rules.label_matcher.keys` property

Type: `["list", "string"]`. Computed.

The list of label key names that have to match.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0321021231333001-0022230203210300-1002232012330002-3022111031230302-3103032313102323-2323130032322131-0001313111001332-1003112213103303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.metadata` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331)
- rule_list.rules.metadata

<a id="canonical-0111232203213322-1230222120032121-3220233233003100-2013213220301101-0222133101213000-2312201332011033-0121212023000320-3221201122013303"></a>

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

<a id="canonical-2233031122212232-2332003332122331-1000200113133200-1122310013213003-3132111210130102-2213303323232111-1031301012330221-2320312300210003"></a>

### Direct properties for `rule_list.rules.metadata`

<a id="canonical-2323022123110220-3332313230201002-0311113121100112-2221012220112022-3213010020020012-1133113012030000-2330001303302022-0321302323031111"></a>

#### `rule_list.rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-1001001230322021-1322013233223313-3103020132102323-3322232002322332-3030300210300312-3011301133303202-3233101220001300-1202102103300013"></a>

<a id="canonical-0323233211021300-1303311232330210-3021301322121101-0113211021122320-0322023102323103-1031203013210030-3002002200300132-0322302232322101"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0233033010303013-3213003233131011-2023031031100121-1211222010230211-2120031331311133-2333012030322021-0310020130103012-0131111012332101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.outside_destinations` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331)
- rule_list.rules.outside_destinations

<a id="canonical-1330222020023201-1101022233003121-1012000000000001-1032313310101110-0233102020222113-0002220332002222-3222323311323233-0222301012131022"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for outside destinations.

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

<a id="canonical-0331001112202220-0222020210122001-2200031112222223-1232013223331332-0301001312220303-2011302001200322-0312213200131101-2122320102311110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.outside_sources` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331)
- rule_list.rules.outside_sources

<a id="canonical-1013331211213330-0002000210011223-3311112210301022-1330011322103330-2223330302322130-0322323121313111-0330222310320301-2132012301333123"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for outside sources.

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

<a id="canonical-1010210030311031-2302010003102121-3030001023003210-0303201120131000-3120000013333020-1222300030311223-0112130221003021-0312213010010203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.protocol_port_range` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331)
- rule_list.rules.protocol_port_range

<a id="canonical-0123321303223122-0032010033121221-3100331311022322-0030123213311330-1011230033300133-3001210112120313-3100030323320313-1231132133233113"></a>

Type: `"single"`. Computed.

Protocol and Port. Protocol and Port ranges.

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

<a id="canonical-1030233312332310-3213023210302011-1211320123333203-1113201311000321-2301133131003112-3113321303100013-1222000312332302-0220012232301323"></a>

### Direct properties for `rule_list.rules.protocol_port_range`

<a id="canonical-2313013333200231-3101133332001000-1032331323103103-2111200213220321-0023130112102311-2331111203323103-2321010200202030-1303300312322223"></a>

#### `rule_list.rules.protocol_port_range.port_ranges` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-1131130031201310-1123300101112302-1003232022031222-0222322110330331-2112122330012231-3103201330311203-2001310322220023-1200200001303311"></a>

<a id="canonical-1113301111013223-0301000020003010-0103012321131130-1131333333132330-3131220320112011-2001300332003231-1230222122131213-3130132012223303"></a>

#### `rule_list.rules.protocol_port_range.protocol` property

Type: `"string"`. Computed.

\[Enum: ALL|TCP|UDP|ICMP\] Protocol in IP packet to be used as match criteria Values are TCP, UDP,
and icmp. Possible values are \`ALL\`, \`TCP\`, \`UDP\`, \`ICMP\`.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ALL",
    "TCP",
    "UDP",
    "ICMP"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  }
}
```

<a id="canonical-1303330131133222-2011312300131102-3103333300023021-3113302301021331-2022112113023002-1302120120322100-0312301222031211-3331302220131002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.source_aws_vpc_ids` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331)
- rule_list.rules.source_aws_vpc_ids

<a id="canonical-2212121000313110-2302310112011032-0113303031020021-0200211110101232-2013130113123213-0103333111032123-3221102002002222-3111012032132313"></a>

Type: `"single"`. Computed.

Configuration parameter for source aws vpc IDs.

Additional upstream details:

List of VPC Identifiers in AWS.

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

<a id="canonical-1332301211321002-0111320032222313-3311130001123011-3120122122201200-0032033323221231-0131012033332012-1323102332200103-1330113323123220"></a>

### Direct properties for `rule_list.rules.source_aws_vpc_ids`

<a id="canonical-1213030122202103-2331122001021113-2010123222202300-3323130101233003-0321200130203110-3330321321220011-0230320210322021-0231003303332101"></a>

#### `rule_list.rules.source_aws_vpc_ids.vpc_id` property

Type: `["list", "string"]`. Computed.

AWS VPC List. List of VPC Identifiers in AWS.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0200023231312311-1323123232302021-1000112103130300-0000330202320113-0122021212200300-0332113310010120-0323322200311330-0300303210223333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.source_ip_prefix_set` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331)
- rule_list.rules.source_ip_prefix_set

<a id="canonical-3311001210322232-2001002302210030-1122332011222221-2003003101023130-0311121010103020-1020201110000311-1331021133323012-2012301132001300"></a>

Type: `"single"`. Computed.

A list of references to ip\_prefix\_set objects.

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

<a id="canonical-1033211122010020-0211310101023333-1323032113233311-0321323232122203-2120211331123230-0021333201132310-0102200301003032-3113103202000202"></a>

### Direct properties for `rule_list.rules.source_ip_prefix_set`

- [ref](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3133021033113302-0212300221213121-1323033311132310-3301102011112123-0013120210100221-0113300132232022-0032033013101200-3000003232031322): complete subsection reference.

<a id="canonical-3133021033113302-0212300221213121-1323033311132310-3301102011112123-0013120210100221-0113300132232022-0032033013101200-3000003232031322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.source_ip_prefix_set.ref` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331)
- [rule_list.rules.source_ip_prefix_set](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0200023231312311-1323123232302021-1000112103130300-0000330202320113-0122021212200300-0332113310010120-0323322200311330-0300303210223333)
- rule_list.rules.source_ip_prefix_set.ref

<a id="canonical-1310312311112330-0300101303131320-3003111210121220-3300201221102103-0010322203000211-2000011230230302-3113322131312110-3123202131001022"></a>

Type: `"list"`. Computed.

A list of references to ip\_prefix\_set objects.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1202231121021313-0033222101030001-3101020110311232-0123133203300121-1201023200020113-3202000203320310-1310022203313103-0300220113300332"></a>

### Direct properties for `rule_list.rules.source_ip_prefix_set.ref`

<a id="canonical-3221111010220000-1002102101310311-2131023121111031-2300120033112101-2132021100130201-3012323123110123-1221302020200223-0003023222130113"></a>

#### `rule_list.rules.source_ip_prefix_set.ref.kind` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2201200230303000-2120202313031122-3231323311100221-0113121313201122-1010302203222322-2333332220120232-3201133120333023-1300200230132302"></a>

<a id="canonical-3203211130312211-3202221031301123-2332332011132210-2133101301033221-0032110021012011-0011323312100211-0001100122320110-3310202201112103"></a>

#### `rule_list.rules.source_ip_prefix_set.ref.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2123311322120110-2132011200122322-1313002321022123-0121003023023102-1103332303222330-3000133021231000-1203333020302003-3303012311201331"></a>

<a id="canonical-3201203311200032-1000331333101212-2000303003230033-2321031021232200-0121110121232233-1123310100030122-0203130102330231-0302212210302223"></a>

#### `rule_list.rules.source_ip_prefix_set.ref.namespace` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3203230132130212-2102033300133202-0032021102323120-1012003313121322-2103332322012003-2210220212002023-0000000023121021-0202320311123120"></a>

<a id="canonical-2211331123331312-0310310030211313-0103220300232132-2100203110021112-1203101013110333-1132301111210221-1233210200213120-2010132230212103"></a>

#### `rule_list.rules.source_ip_prefix_set.ref.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1300123303000201-3230300010020201-0033302213131000-2303311010022322-0122322001202232-3110321010022012-2321330123330212-1321231120312233"></a>

<a id="canonical-3310023203002213-2103302102110032-1030301202113300-1023320133132130-3313020011010110-1122230322212011-2030031001000020-1010103212211201"></a>

#### `rule_list.rules.source_ip_prefix_set.ref.uid` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0011023231130030-2211332223100010-3323313001301320-2003202111312011-0202033201001323-0003021023122232-0301132123113023-3323313201122300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.source_label_selector` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331)
- rule_list.rules.source_label_selector

<a id="canonical-3233103021113000-3300202332111230-3313210211220023-3221101002313000-2132110133010300-3033232220230311-2102012333133103-1102313220222203"></a>

Type: `"single"`. Computed.

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

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

<a id="canonical-2222030003221022-3222102322332011-0101111313121003-3010322303300321-2023222201101120-3111323332012313-1020020022223022-3202100003001303"></a>

### Direct properties for `rule_list.rules.source_label_selector`

<a id="canonical-2200111101111213-1202003211103131-1003130131120223-2331231110020332-2120022000200020-1331103012310022-1332032121210113-1101312110021300"></a>

#### `rule_list.rules.source_label_selector.expressions` property

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-0232011221020021-0110002001020030-0300123202033112-2132210231313033-2301100230222303-2302213313102120-3320330233302223-0013032000233310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.source_prefix_list` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2222221232113201-2332332231203200-0002112022102311-1000133032211001-1011223300133010-3300003012031231-1130121231201200-1013230320331012)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331)
- rule_list.rules.source_prefix_list

<a id="canonical-2300230031233203-2133002220310102-0120032012310201-0302332023210002-3112313311232332-1112311112202213-1103310311122121-0202022120221331"></a>

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

<a id="canonical-2303210231330130-0231230133223200-2321121221133101-3223003303323200-3303030031312102-1101321332033212-2323110020222033-2003032011222100"></a>

### Direct properties for `rule_list.rules.source_prefix_list`

<a id="canonical-2020201110132331-2303333013001020-1012213010113201-3113113023230032-2200300212333210-3000100022310110-1131203132330101-3022112333033111"></a>

#### `rule_list.rules.source_prefix_list.prefixes` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
