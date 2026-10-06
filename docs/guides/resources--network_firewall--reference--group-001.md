---
page_title: "xcsh_network_firewall reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_network_firewall reference."
---

# xcsh_network_firewall reference

<a id="canonical-1131321320323100-0133323001330022-1211003101102003-2000211000122312-2031221212010020-2323201312033210-3331113002030021-0010133133330010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md#canonical-0002103131303200-3202220110210100-2001233131212311-3322200221303002-0011132331222322-2101102312302011-0113033313013211-2001132200001333)
- Property reference

<a id="canonical-0212122232330331-3112311330331230-3301102230112300-1331320313203303-1012031302101301-2331000222112302-0303103101221021-2322210303201333"></a>

### Direct properties for `xcsh_network_firewall`

- [active_enhanced_firewall_policies](resources--network_firewall--reference--group-001.md#canonical-2313013321102121-1013201323302331-2222111323201110-2011320113312130-1311100031233101-2112200023113230-2233122231303002-3003213211332212): complete subsection reference.

- [active_fast_acls](resources--network_firewall--reference--group-001.md#canonical-3112210130022303-2030213233130033-2121231033222300-2033311300122320-0021030220003103-2012033313212000-3303200320201022-1230312021203230): complete subsection reference.

- [active_forward_proxy_policies](resources--network_firewall--reference--group-001.md#canonical-0102101002112321-0202233211212032-1112101221133102-3312333102301221-0010220230112111-0223201103230311-0103311201122130-3010200022203100): complete subsection reference.

- [active_network_policies](resources--network_firewall--reference--group-001.md#canonical-2113302023113302-3020022101313123-0003310010231013-1230332313211203-3122211313002211-2212312210212203-2312100331130132-3113020002020033): complete subsection reference.

<a id="canonical-3020232131320131-1120033131112201-0231001331000002-2111202103111222-0330330310300320-2322022132211233-0113000133010302-0232231120030012"></a>

<a id="canonical-2331113220013013-0111301202303301-1322311201012330-3312230332001130-3232101321223330-0123032101010020-3220220020121120-3203112320121033"></a>

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

<a id="canonical-2120230023032003-3222310300122032-3313122000131332-2233301330133323-0330033120010302-1132212301312211-0002113233312310-3231103333303022"></a>

<a id="canonical-2031300120322131-3033021332310000-0323101303123001-3003011123122230-0112013221300133-2202111300111220-2030110012310121-3230112301310211"></a>

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

<a id="canonical-2031011311022100-1331132131333012-3301110111222231-1021311032030312-1312030310220310-1012123310212222-2330333330133102-0110220010300301"></a>

<a id="canonical-1200033110323122-2201123000320211-2122032212320331-3320110111333323-1332323123110020-0201112033100033-3030210120230301-3201011003111312"></a>

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

- [disable_fast_acl](resources--network_firewall--reference--group-001.md#canonical-1030000322023102-2003301200002102-0123111133031210-3301113222100120-1330210120023320-2333233223032033-0232123000221113-0213001130313013): complete subsection reference.

- [disable_forward_proxy_policy](resources--network_firewall--reference--group-001.md#canonical-0103132011303301-0020002110202103-2131012331110100-3000212003030130-0102313021312333-1022012021122221-3100310310331201-1331222212320112): complete subsection reference.

- [disable_network_policy](resources--network_firewall--reference--group-001.md#canonical-1333010313123013-1233301021033300-2301032023021121-3312311230002220-0031232132010223-3220203322122032-0020300220312301-3100211112201021): complete subsection reference.

<a id="canonical-1203222120203221-1320012033120233-3131112211220103-3120232221113132-3130032030132102-3112321313202310-2113132220130000-0133311023320001"></a>

<a id="canonical-1302111233303301-1123131022101211-3200130213000122-1111222231021000-2000121120221302-0032121110003033-1021301131323301-1330123223113121"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1212132101033202-2013020232031300-0333122032103303-3111330023331201-1100200332232220-2030102030002221-0202203331223330-3211222321021100"></a>

<a id="canonical-3021100011023323-0020313222101131-1120131130003002-3020133010230102-3202031230210022-2303000201102002-3130103000232113-3123211032230030"></a>

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

<a id="canonical-0031330023003332-3023111021113303-0013220333313303-2003332131202201-3220303032332023-2023130330221111-1222303001333032-2203122222223230"></a>

<a id="canonical-1311112111022022-0230200232313130-1213312010100313-3123202210300302-2132012022211130-2122000331021110-1012332303322130-0200112321332122"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Network Firewall. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1110222302130203-3011102013200102-2320101021221302-2330133223200313-0233310201003320-2011023011310123-3111312211001113-2100010020122113"></a>

<a id="canonical-2011201110221022-2133012102113201-1011230310033012-3313022311130210-2223021221100233-2000003111121231-2010321111102202-2010133001300210"></a>

#### `namespace` property

Type: `"string"`. Optional, Computed.

Namespace for the Network Firewall. The F5 XC API restricts this resource to the system namespace;
it defaults to that value and may be omitted.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
Validators: []validator.String{
  validators.NamespaceValidator(),
  stringvalidator.OneOf("system"),
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

- [timeouts](resources--network_firewall--reference--group-001.md#canonical-0103233233330031-1302333012031133-2220013312030030-2230210022321230-1211331311321201-0321110000133010-0101323112031033-2212201030332022): complete subsection reference.

<a id="canonical-1130221120022020-1011302203103333-2320022323232033-2320321133030212-0332101230200231-0133331310000101-1122213312301132-1232113223320112"></a>

### All schema paths for `xcsh_network_firewall`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `active_enhanced_firewall_policies` | [active_enhanced_firewall_policies](resources--network_firewall--reference--group-001.md#canonical-0311322111323322-1011132212112200-2013012000000320-2212132202232000-1333333001313212-3223221210102330-2121210212033311-0331203232233011) |
| `active_enhanced_firewall_policies.enhanced_firewall_policies` | [active_enhanced_firewall_policies.enhanced_firewall_policies](resources--network_firewall--reference--group-001.md#canonical-1100320100321202-1130230003203110-1300333103021130-2120031322010333-0103022223000121-3010300323113131-2122232323313203-1233222120000000) |
| `active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [active_enhanced_firewall_policies.enhanced_firewall_policies.name](resources--network_firewall--reference--group-001.md#canonical-0031102131212112-3002131330313030-1332130312132032-3113120010230101-0122320123303302-1222221301010200-0320310323103212-2212010302022330) |
| `active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](resources--network_firewall--reference--group-001.md#canonical-0230333211300001-2122312213123203-3133003122311113-3220210202210203-1330103211102031-1022100311203331-1112010221301332-0223222311231203) |
| `active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](resources--network_firewall--reference--group-001.md#canonical-2301333221200120-0311000230013322-1032022303213223-0233220302100113-1310001111310012-3013111020331111-3021112213333133-2210210010220200) |
| `active_fast_acls` | [active_fast_acls](resources--network_firewall--reference--group-001.md#canonical-0322111303320012-2330221021300313-3222233133233113-0033032312200213-2221221000330013-1302120021013301-0223012013021230-0300131010131120) |
| `active_fast_acls.fast_acls` | [active_fast_acls.fast_acls](resources--network_firewall--reference--group-001.md#canonical-0130100123123020-0230221130103311-0332022021021222-0103313200133312-0132211312210032-1032003111032113-1011233111012330-2201131021202202) |
| `active_fast_acls.fast_acls.name` | [active_fast_acls.fast_acls.name](resources--network_firewall--reference--group-001.md#canonical-2121332112202020-1002021200000200-2001323000132123-0123210022313010-1012102331013313-1303200010230121-1220102111233133-0330113101013321) |
| `active_fast_acls.fast_acls.namespace` | [active_fast_acls.fast_acls.namespace](resources--network_firewall--reference--group-001.md#canonical-1121030132222323-2000222313221001-2130111100230022-2122210300312123-2003012101212221-3233203123000330-0130202300233123-3221202100001321) |
| `active_fast_acls.fast_acls.tenant` | [active_fast_acls.fast_acls.tenant](resources--network_firewall--reference--group-001.md#canonical-2230300101233110-3211131331320020-3213310232223023-0211130022031323-0213101330010212-2221012122210031-3322231320122131-1011210332113021) |
| `active_forward_proxy_policies` | [active_forward_proxy_policies](resources--network_firewall--reference--group-001.md#canonical-0303232213132003-1323231312100200-3233133320200000-3221102330203201-0113121200010100-3121031223101030-3232002031101103-2030232010032000) |
| `active_forward_proxy_policies.forward_proxy_policies` | [active_forward_proxy_policies.forward_proxy_policies](resources--network_firewall--reference--group-001.md#canonical-1122022212001223-3123310001121322-3201021033103302-3320123002321230-0102022201103310-3230130300120302-0003213331203123-3012200232302300) |
| `active_forward_proxy_policies.forward_proxy_policies.name` | [active_forward_proxy_policies.forward_proxy_policies.name](resources--network_firewall--reference--group-001.md#canonical-2102023110000001-1321311202131021-3301130132331312-1011001312021000-3300123223323012-2233300103023302-0033212222311122-2102010201010322) |
| `active_forward_proxy_policies.forward_proxy_policies.namespace` | [active_forward_proxy_policies.forward_proxy_policies.namespace](resources--network_firewall--reference--group-001.md#canonical-0313103330003201-2203012323333102-3020113103022120-2221211023130033-1010312012232122-3323320032102132-1032011112012222-2320203122030113) |
| `active_forward_proxy_policies.forward_proxy_policies.tenant` | [active_forward_proxy_policies.forward_proxy_policies.tenant](resources--network_firewall--reference--group-001.md#canonical-0123013233313330-2112202133003113-1303320310113213-3213213112211103-1213030030121011-3121130302103011-1331003112111120-0131032020011021) |
| `active_network_policies` | [active_network_policies](resources--network_firewall--reference--group-001.md#canonical-3212211303222321-3013200010213021-0000213220203202-3111032030130203-3301312123303332-3233303330022111-0103022311110022-3323202002223122) |
| `active_network_policies.network_policies` | [active_network_policies.network_policies](resources--network_firewall--reference--group-001.md#canonical-2000313303222221-0021112000020200-0223203110230001-1100210003023220-0102021330000232-0110213223231210-0000001331300113-0303001131213311) |
| `active_network_policies.network_policies.name` | [active_network_policies.network_policies.name](resources--network_firewall--reference--group-001.md#canonical-2211022211121212-2232103102323323-2122203033130123-0101301310332112-0213303033211120-2122321210033113-3212210323323103-2022102133213231) |
| `active_network_policies.network_policies.namespace` | [active_network_policies.network_policies.namespace](resources--network_firewall--reference--group-001.md#canonical-1121230031121012-1211000120020021-1321312103102131-1123232300122220-0000033032121012-2211232202111313-2203301332111010-2223123232131300) |
| `active_network_policies.network_policies.tenant` | [active_network_policies.network_policies.tenant](resources--network_firewall--reference--group-001.md#canonical-2202000322313320-0221232020303212-0100011320121231-0322233333111201-0123111000123110-2132302210100320-3131112332130011-2002101021220013) |
| `annotations` | [annotations](resources--network_firewall--reference--group-001.md#canonical-3020232131320131-1120033131112201-0231001331000002-2111202103111222-0330330310300320-2322022132211233-0113000133010302-0232231120030012) |
| `description` | [description](resources--network_firewall--reference--group-001.md#canonical-2120230023032003-3222310300122032-3313122000131332-2233301330133323-0330033120010302-1132212301312211-0002113233312310-3231103333303022) |
| `disable` | [disable](resources--network_firewall--reference--group-001.md#canonical-2031011311022100-1331132131333012-3301110111222231-1021311032030312-1312030310220310-1012123310212222-2330333330133102-0110220010300301) |
| `disable_fast_acl` | [disable_fast_acl](resources--network_firewall--reference--group-001.md#canonical-0020322223033222-1033131300010010-2012212322122112-2312012011133220-3202031230013310-1002012332122123-1102301030132202-0022221003330211) |
| `disable_forward_proxy_policy` | [disable_forward_proxy_policy](resources--network_firewall--reference--group-001.md#canonical-0103320300101221-1332203231121100-3022223331012323-2030110200313300-3311023131223001-0213213211030001-2022310201320021-2302012232122013) |
| `disable_network_policy` | [disable_network_policy](resources--network_firewall--reference--group-001.md#canonical-3330312320103103-2021102031012310-2002310231303130-2323202213222231-2113221133220012-2103201002223012-1120120100310212-2212100030131323) |
| `id` | [ID](resources--network_firewall--reference--group-001.md#canonical-1203222120203221-1320012033120233-3131112211220103-3120232221113132-3130032030132102-3112321313202310-2113132220130000-0133311023320001) |
| `labels` | [labels](resources--network_firewall--reference--group-001.md#canonical-1212132101033202-2013020232031300-0333122032103303-3111330023331201-1100200332232220-2030102030002221-0202203331223330-3211222321021100) |
| `name` | [name](resources--network_firewall--reference--group-001.md#canonical-0031330023003332-3023111021113303-0013220333313303-2003332131202201-3220303032332023-2023130330221111-1222303001333032-2203122222223230) |
| `namespace` | [namespace](resources--network_firewall--reference--group-001.md#canonical-1110222302130203-3011102013200102-2320101021221302-2330133223200313-0233310201003320-2011023011310123-3111312211001113-2100010020122113) |
| `timeouts` | [timeouts](resources--network_firewall--reference--group-001.md#canonical-2211321110131132-0331213000130022-0312111011102222-1111000113013331-1001330330331120-2230131033323131-2303232223031310-1003130010010012) |
| `timeouts.create` | [timeouts.create](resources--network_firewall--reference--group-001.md#canonical-2020212100303230-3011210100113201-2231211311313002-0223022012302111-0300000312031323-0001332121112110-1112223101022330-2332320021132131) |
| `timeouts.delete` | [timeouts.delete](resources--network_firewall--reference--group-001.md#canonical-0221031031320120-0020003122122302-1111130033210121-2200012200310231-1232202113300011-0332110112111022-2221130311010010-3221023030321103) |
| `timeouts.read` | [timeouts.read](resources--network_firewall--reference--group-001.md#canonical-3333033332200120-3222312023122330-2223323120202212-3101010023120212-1303030322222012-1100100032003323-3311222212203011-3210003133303212) |
| `timeouts.update` | [timeouts.update](resources--network_firewall--reference--group-001.md#canonical-0220233113211202-1321323313231311-1331133100023323-2110000130322333-0220221002301001-1233310032311012-3131223130221300-1030013000313011) |

<a id="canonical-2313013321102121-1013201323302331-2222111323201110-2011320113312130-1311100031233101-2112200023113230-2233122231303002-3003213211332212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_enhanced_firewall_policies` properties

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md#canonical-0002103131303200-3202220110210100-2001233131212311-3322200221303002-0011132331222322-2101102312302011-0113033313013211-2001132200001333)
- [Property reference](resources--network_firewall--reference--group-001.md#canonical-1131321320323100-0133323001330022-1211003101102003-2000211000122312-2031221212010020-2323201312033210-3331113002030021-0010133133330010)
- active_enhanced_firewall_policies

<a id="canonical-0311322111323322-1011132212112200-2013012000000320-2212132202232000-1333333001313212-3223221210102330-2121210212033311-0331203232233011"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: active\_enhanced\_firewall\_policies, active\_network\_policies, disable\_network\_policy;
Default: disable\_network\_policy\] List of Enhanced Firewall Policies These policies use
session-based rules and provide all OPTIONS available under firewall policies with an additional
option for service insertion.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("enhanced_firewall_policies")}
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

OneOf alternatives in this subsection:

- [active_enhanced_firewall_policies](resources--network_firewall--reference--group-001.md#canonical-0311322111323322-1011132212112200-2013012000000320-2212132202232000-1333333001313212-3223221210102330-2121210212033311-0331203232233011)
- [active_network_policies](resources--network_firewall--reference--group-001.md#canonical-3212211303222321-3013200010213021-0000213220203202-3111032030130203-3301312123303332-3233303330022111-0103022311110022-3323202002223122)
- [disable_network_policy](resources--network_firewall--reference--group-001.md#canonical-3330312320103103-2021102031012310-2002310231303130-2323202213222231-2113221133220012-2103201002223012-1120120100310212-2212100030131323)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
active_enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-0100213122113121-3131130011000303-3110121011321123-0130101321321132-1230110200122220-3331322213133000-0022101331103010-2331030300002212"></a>

### Direct properties for `active_enhanced_firewall_policies`

- [enhanced_firewall_policies](resources--network_firewall--reference--group-001.md#canonical-1033101302130033-1322031110323001-2010003000003002-3331030321111101-0102211013322011-3112301203222313-1010332030132333-1331300322012021): complete subsection reference.

<a id="canonical-1033101302130033-1322031110323001-2010003000003002-3331030321111101-0102211013322011-3112301203222313-1010332030132333-1331300322012021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_enhanced_firewall_policies.enhanced_firewall_policies` properties

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md#canonical-0002103131303200-3202220110210100-2001233131212311-3322200221303002-0011132331222322-2101102312302011-0113033313013211-2001132200001333)
- [Property reference](resources--network_firewall--reference--group-001.md#canonical-1131321320323100-0133323001330022-1211003101102003-2000211000122312-2031221212010020-2323201312033210-3331113002030021-0010133133330010)
- [active_enhanced_firewall_policies](resources--network_firewall--reference--group-001.md#canonical-2313013321102121-1013201323302331-2222111323201110-2011320113312130-1311100031233101-2112200023113230-2233122231303002-3003213211332212)
- active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-1100320100321202-1130230003203110-1300333103021130-2120031322010333-0103022223000121-3010300323113131-2122232323313203-1233222120000000"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Enhanced Firewall Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "minItems": 1
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
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-3230233201021001-0001302102231133-2131313202331332-1303231123222220-2200231120303213-2112231022332232-1313233031221231-1300013120021002"></a>

### Direct properties for `active_enhanced_firewall_policies.enhanced_firewall_policies`

<a id="canonical-0031102131212112-3002131330313030-1332130312132032-3113120010230101-0122320123303302-1222221301010200-0320310323103212-2212010302022330"></a>

#### `active_enhanced_firewall_policies.enhanced_firewall_policies.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0230333211300001-2122312213123203-3133003122311113-3220210202210203-1330103211102031-1022100311203331-1112010221301332-0223222311231203"></a>

<a id="canonical-2212123123202132-2103203003033323-0313123132213121-1132230002130232-0112312121022120-1323030001121223-2323301320122102-0112023330301323"></a>

#### `active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2301333221200120-0311000230013322-1032022303213223-0233220302100113-1310001111310012-3013111020331111-3021112213333133-2210210010220200"></a>

<a id="canonical-1030301233320200-0321322331103012-0131002332233132-2111113110231033-1000011122103313-1211202022320301-3313101001021331-0133310231332102"></a>

#### `active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3112210130022303-2030213233130033-2121231033222300-2033311300122320-0021030220003103-2012033313212000-3303200320201022-1230312021203230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_fast_acls` properties

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md#canonical-0002103131303200-3202220110210100-2001233131212311-3322200221303002-0011132331222322-2101102312302011-0113033313013211-2001132200001333)
- [Property reference](resources--network_firewall--reference--group-001.md#canonical-1131321320323100-0133323001330022-1211003101102003-2000211000122312-2031221212010020-2323201312033210-3331113002030021-0010133133330010)
- active_fast_acls

<a id="canonical-0322111303320012-2330221021300313-3222233133233113-0033032312200213-2221221000330013-1302120021013301-0223012013021230-0300131010131120"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: active\_fast\_acls, disable\_fast\_acl; Default: disable\_fast\_acl\] Configuration
parameter for active fast acls.

Additional upstream details:

List of Fast ACL(s).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("fast_acls")}
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

OneOf alternatives in this subsection:

- [active_fast_acls](resources--network_firewall--reference--group-001.md#canonical-0322111303320012-2330221021300313-3222233133233113-0033032312200213-2221221000330013-1302120021013301-0223012013021230-0300131010131120)
- [disable_fast_acl](resources--network_firewall--reference--group-001.md#canonical-0020322223033222-1033131300010010-2012212322122112-2312012011133220-3202031230013310-1002012332122123-1102301030132202-0022221003330211)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
active_fast_acls {
  # Configure direct properties listed below.
}
```

<a id="canonical-0103323231110202-1122210330331220-0330333120101100-1113100012021110-0320132203313313-3033032223320133-1210210220122213-2103033331223211"></a>

### Direct properties for `active_fast_acls`

- [fast_acls](resources--network_firewall--reference--group-001.md#canonical-2213130101201213-3221033011233030-3323131313202030-3110101232132100-3010210012203302-1120232110212020-1031330130000212-1320213201331300): complete subsection reference.

<a id="canonical-2213130101201213-3221033011233030-3323131313202030-3110101232132100-3010210012203302-1120232110212020-1031330130000212-1320213201331300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_fast_acls.fast_acls` properties

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md#canonical-0002103131303200-3202220110210100-2001233131212311-3322200221303002-0011132331222322-2101102312302011-0113033313013211-2001132200001333)
- [Property reference](resources--network_firewall--reference--group-001.md#canonical-1131321320323100-0133323001330022-1211003101102003-2000211000122312-2031221212010020-2323201312033210-3331113002030021-0010133133330010)
- [active_fast_acls](resources--network_firewall--reference--group-001.md#canonical-3112210130022303-2030213233130033-2121231033222300-2033311300122320-0021030220003103-2012033313212000-3303200320201022-1230312021203230)
- active_fast_acls.fast_acls

<a id="canonical-0130100123123020-0230221130103311-0332022021021222-0103313200133312-0132211312210032-1032003111032113-1011233111012330-2201131021202202"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Fast ACL(s) active for this network firewall.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "minItems": 1
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
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
fast_acls {
  # Configure direct properties listed below.
}
```

<a id="canonical-2012112222233102-3301210200223330-0111221133112012-1123001203001330-3131020130222223-2302131132202330-1220202123101110-1022111212101203"></a>

### Direct properties for `active_fast_acls.fast_acls`

<a id="canonical-2121332112202020-1002021200000200-2001323000132123-0123210022313010-1012102331013313-1303200010230121-1220102111233133-0330113101013321"></a>

#### `active_fast_acls.fast_acls.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1121030132222323-2000222313221001-2130111100230022-2122210300312123-2003012101212221-3233203123000330-0130202300233123-3221202100001321"></a>

<a id="canonical-3013300230321130-2011102102323330-0322330033000321-0202133110001211-2101301121301102-2031310133313200-0220031221130200-0010133112232200"></a>

#### `active_fast_acls.fast_acls.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2230300101233110-3211131331320020-3213310232223023-0211130022031323-0213101330010212-2221012122210031-3322231320122131-1011210332113021"></a>

<a id="canonical-2301230320121321-1101123212223110-0332022102120131-1210001100321122-2323320021033332-2310233203123022-0300221211313001-2302320303033323"></a>

#### `active_fast_acls.fast_acls.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0102101002112321-0202233211212032-1112101221133102-3312333102301221-0010220230112111-0223201103230311-0103311201122130-3010200022203100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_forward_proxy_policies` properties

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md#canonical-0002103131303200-3202220110210100-2001233131212311-3322200221303002-0011132331222322-2101102312302011-0113033313013211-2001132200001333)
- [Property reference](resources--network_firewall--reference--group-001.md#canonical-1131321320323100-0133323001330022-1211003101102003-2000211000122312-2031221212010020-2323201312033210-3331113002030021-0010133133330010)
- active_forward_proxy_policies

<a id="canonical-0303232213132003-1323231312100200-3233133320200000-3221102330203201-0113121200010100-3121031223101030-3232002031101103-2030232010032000"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: active\_forward\_proxy\_policies, disable\_forward\_proxy\_policy; Default:
disable\_forward\_proxy\_policy\] Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("forward_proxy_policies")}
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

OneOf alternatives in this subsection:

- [active_forward_proxy_policies](resources--network_firewall--reference--group-001.md#canonical-0303232213132003-1323231312100200-3233133320200000-3221102330203201-0113121200010100-3121031223101030-3232002031101103-2030232010032000)
- [disable_forward_proxy_policy](resources--network_firewall--reference--group-001.md#canonical-0103320300101221-1332203231121100-3022223331012323-2030110200313300-3311023131223001-0213213211030001-2022310201320021-2302012232122013)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
active_forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-2333123020120001-0012231010232233-1123123033211102-0222010130121321-3201131010111000-2130303330223202-3323121033300320-2013331000201101"></a>

### Direct properties for `active_forward_proxy_policies`

- [forward_proxy_policies](resources--network_firewall--reference--group-001.md#canonical-2032201211002330-0313030231032323-1023010320212131-3100203120200022-0231103010232103-1320121100100220-2321200111232312-1212121232102233): complete subsection reference.

<a id="canonical-2032201211002330-0313030231032323-1023010320212131-3100203120200022-0231103010232103-1320121100100220-2321200111232312-1212121232102233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_forward_proxy_policies.forward_proxy_policies` properties

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md#canonical-0002103131303200-3202220110210100-2001233131212311-3322200221303002-0011132331222322-2101102312302011-0113033313013211-2001132200001333)
- [Property reference](resources--network_firewall--reference--group-001.md#canonical-1131321320323100-0133323001330022-1211003101102003-2000211000122312-2031221212010020-2323201312033210-3331113002030021-0010133133330010)
- [active_forward_proxy_policies](resources--network_firewall--reference--group-001.md#canonical-0102101002112321-0202233211212032-1112101221133102-3312333102301221-0010220230112111-0223201103230311-0103311201122130-3010200022203100)
- active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-1122022212001223-3123310001121322-3201021033103302-3320123002321230-0102022201103310-3230130300120302-0003213331203123-3012200232302300"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "minItems": 1
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
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-0220212131303111-2202031132323101-0322212100201022-2022230222231200-3311212032212001-1013203323322222-2213131122301302-2200031213233202"></a>

### Direct properties for `active_forward_proxy_policies.forward_proxy_policies`

<a id="canonical-2102023110000001-1321311202131021-3301130132331312-1011001312021000-3300123223323012-2233300103023302-0033212222311122-2102010201010322"></a>

#### `active_forward_proxy_policies.forward_proxy_policies.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0313103330003201-2203012323333102-3020113103022120-2221211023130033-1010312012232122-3323320032102132-1032011112012222-2320203122030113"></a>

<a id="canonical-2102021113023331-0020010013030303-1123222332223333-2130310231203331-0100231203303320-1110310112223331-3322022200310200-3303303310203323"></a>

#### `active_forward_proxy_policies.forward_proxy_policies.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0123013233313330-2112202133003113-1303320310113213-3213213112211103-1213030030121011-3121130302103011-1331003112111120-0131032020011021"></a>

<a id="canonical-0002113022122210-0121203202131221-1122333212102301-3211032212120330-1032331113130021-3131221201233123-0202102310123310-3010133232203001"></a>

#### `active_forward_proxy_policies.forward_proxy_policies.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2113302023113302-3020022101313123-0003310010231013-1230332313211203-3122211313002211-2212312210212203-2312100331130132-3113020002020033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_network_policies` properties

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md#canonical-0002103131303200-3202220110210100-2001233131212311-3322200221303002-0011132331222322-2101102312302011-0113033313013211-2001132200001333)
- [Property reference](resources--network_firewall--reference--group-001.md#canonical-1131321320323100-0133323001330022-1211003101102003-2000211000122312-2031221212010020-2323201312033210-3331113002030021-0010133133330010)
- active_network_policies

<a id="canonical-3212211303222321-3013200010213021-0000213220203202-3111032030130203-3301312123303332-3233303330022111-0103022311110022-3323202002223122"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for active network policies.

Additional upstream details:

List of firewall policy views.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("network_policies")}
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
active_network_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-2333012221200220-2202100002011003-3113003230012130-2033100103211231-3313201022002032-1212103310001031-2300021302121001-1103033310123312"></a>

### Direct properties for `active_network_policies`

- [network_policies](resources--network_firewall--reference--group-001.md#canonical-1303133033202032-1000123223302030-2102310103323000-3102201220323311-0211123222032322-1230111233333012-2021230232010133-1031313033200211): complete subsection reference.

<a id="canonical-1303133033202032-1000123223302030-2102310103323000-3102201220323311-0211123222032322-1230111233333012-2021230232010133-1031313033200211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_network_policies.network_policies` properties

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md#canonical-0002103131303200-3202220110210100-2001233131212311-3322200221303002-0011132331222322-2101102312302011-0113033313013211-2001132200001333)
- [Property reference](resources--network_firewall--reference--group-001.md#canonical-1131321320323100-0133323001330022-1211003101102003-2000211000122312-2031221212010020-2323201312033210-3331113002030021-0010133133330010)
- [active_network_policies](resources--network_firewall--reference--group-001.md#canonical-2113302023113302-3020022101313123-0003310010231013-1230332313211203-3122211313002211-2212312210212203-2312100331130132-3113020002020033)
- active_network_policies.network_policies

<a id="canonical-2000313303222221-0021112000020200-0223203110230001-1100210003023220-0102021330000232-0110213223231210-0000001331300113-0303001131213311"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Firewall Policies active for this network firewall.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "minItems": 1
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
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
network_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-3013023201231333-0113113333301030-3203013300002222-3032222010110302-1113223022020020-3312103211111111-2002111222123002-0302003000110212"></a>

### Direct properties for `active_network_policies.network_policies`

<a id="canonical-2211022211121212-2232103102323323-2122203033130123-0101301310332112-0213303033211120-2122321210033113-3212210323323103-2022102133213231"></a>

#### `active_network_policies.network_policies.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1121230031121012-1211000120020021-1321312103102131-1123232300122220-0000033032121012-2211232202111313-2203301332111010-2223123232131300"></a>

<a id="canonical-3133201302132222-1012133011132311-0323330001111333-1303132122230201-1101101310211330-1032130123112212-1102013121321002-0202313320003310"></a>

#### `active_network_policies.network_policies.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2202000322313320-0221232020303212-0100011320121231-0322233333111201-0123111000123110-2132302210100320-3131112332130011-2002101021220013"></a>

<a id="canonical-1223113110233003-0112122310320111-1313233220310321-1211112130322322-1330221010323331-1301123312300111-1122322000231120-0132010113213233"></a>

#### `active_network_policies.network_policies.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1030000322023102-2003301200002102-0123111133031210-3301113222100120-1330210120023320-2333233223032033-0232123000221113-0213001130313013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_fast_acl` properties

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md#canonical-0002103131303200-3202220110210100-2001233131212311-3322200221303002-0011132331222322-2101102312302011-0113033313013211-2001132200001333)
- [Property reference](resources--network_firewall--reference--group-001.md#canonical-1131321320323100-0133323001330022-1211003101102003-2000211000122312-2031221212010020-2323201312033210-3331113002030021-0010133133330010)
- disable_fast_acl

<a id="canonical-0020322223033222-1033131300010010-2012212322122112-2312012011133220-3202031230013310-1002012332122123-1102301030132202-0022221003330211"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for disable fast acl. Defaults to \`map\[\]\`. Server applies default when
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

Terraform syntax:

```terraform
disable_fast_acl = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0103132011303301-0020002110202103-2131012331110100-3000212003030130-0102313021312333-1022012021122221-3100310310331201-1331222212320112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_forward_proxy_policy` properties

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md#canonical-0002103131303200-3202220110210100-2001233131212311-3322200221303002-0011132331222322-2101102312302011-0113033313013211-2001132200001333)
- [Property reference](resources--network_firewall--reference--group-001.md#canonical-1131321320323100-0133323001330022-1211003101102003-2000211000122312-2031221212010020-2323201312033210-3331113002030021-0010133133330010)
- disable_forward_proxy_policy

<a id="canonical-0103320300101221-1332203231121100-3022223331012323-2030110200313300-3311023131223001-0213213211030001-2022310201320021-2302012232122013"></a>

Type: `["object", {}]`. Optional, Computed.

Policy configuration for this feature. Defaults to \`map\[\]\`. Server applies default when omitted.

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
disable_forward_proxy_policy = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1333010313123013-1233301021033300-2301032023021121-3312311230002220-0031232132010223-3220203322122032-0020300220312301-3100211112201021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_network_policy` properties

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md#canonical-0002103131303200-3202220110210100-2001233131212311-3322200221303002-0011132331222322-2101102312302011-0113033313013211-2001132200001333)
- [Property reference](resources--network_firewall--reference--group-001.md#canonical-1131321320323100-0133323001330022-1211003101102003-2000211000122312-2031221212010020-2323201312033210-3331113002030021-0010133133330010)
- disable_network_policy

<a id="canonical-3330312320103103-2021102031012310-2002310231303130-2323202213222231-2113221133220012-2103201002223012-1120120100310212-2212100030131323"></a>

Type: `["object", {}]`. Optional, Computed.

Policy configuration for this feature. Defaults to \`map\[\]\`. Server applies default when omitted.

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
disable_network_policy = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0103233233330031-1302333012031133-2220013312030030-2230210022321230-1211331311321201-0321110000133010-0101323112031033-2212201030332022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md#canonical-0002103131303200-3202220110210100-2001233131212311-3322200221303002-0011132331222322-2101102312302011-0113033313013211-2001132200001333)
- [Property reference](resources--network_firewall--reference--group-001.md#canonical-1131321320323100-0133323001330022-1211003101102003-2000211000122312-2031221212010020-2323201312033210-3331113002030021-0010133133330010)
- timeouts

<a id="canonical-2211321110131132-0331213000130022-0312111011102222-1111000113013331-1001330330331120-2230131033323131-2303232223031310-1003130010010012"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3123031032031321-1132202332102122-1332211033320233-1322022201111020-2122103112001330-1231230310111213-3010020022003300-1211201200132102"></a>

### Direct properties for `timeouts`

<a id="canonical-2020212100303230-3011210100113201-2231211311313002-0223022012302111-0300000312031323-0001332121112110-1112223101022330-2332320021132131"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0221031031320120-0020003122122302-1111130033210121-2200012200310231-1232202113300011-0332110112111022-2221130311010010-3221023030321103"></a>

<a id="canonical-1033220020032203-2313123332132111-2212230002131111-0100302131210201-0012322130213110-3220333000023202-1233310313120330-1113102010311131"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3333033332200120-3222312023122330-2223323120202212-3101010023120212-1303030322222012-1100100032003323-3311222212203011-3210003133303212"></a>

<a id="canonical-0123210102231133-1102333012303100-2201300001021123-0111310230013021-0333303033212301-0303310100001230-3320211030002131-0103121310333113"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0220233113211202-1321323313231311-1331133100023323-2110000130322333-0220221002301001-1233310032311012-3131223130221300-1030013000313011"></a>

<a id="canonical-3133110010112320-1231232110110310-2123233202200302-0120322230002002-1213122022120311-1130013231011211-1311310103332322-0212213322021132"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
