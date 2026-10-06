---
page_title: "xcsh_fast_acl_rule reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fast_acl_rule reference."
---

# xcsh_fast_acl_rule reference

<a id="canonical-0130033020112332-1320113013031122-3323102231201130-2302010233220010-0332222213110322-3330103313001312-0313202331310301-0031300023203130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)
- Property reference

<a id="canonical-2112202213021033-1033233213220221-1133303320010003-2223302231333310-3231130010220310-3100212021023213-1132200332232023-1103021121312121"></a>

### Direct properties for `xcsh_fast_acl_rule`

- [action](data-sources--fast_acl_rule--reference--group-001.md#canonical-2323322103100113-0330221003132312-0022232121231222-2001112100311031-3033102332230020-2200221013103021-0101203130002020-0333000311113233): complete subsection reference.

<a id="canonical-3231311312011220-3321113331121132-0323222121323322-0112001222130322-2331323321210111-1120313320001223-3112032012003213-0233300210120032"></a>

<a id="canonical-0021210132033321-1101033331303321-0231100312323102-0322203221010033-0120121322023322-1101120232331220-1132111002320031-2111131310032231"></a>

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

<a id="canonical-0012321002201302-2322132122213120-1101200021213131-2123002320133111-3113220010031033-2021131310233301-0312000103011320-1021031120311130"></a>

<a id="canonical-0012100210201120-0222333223210033-0033132230012023-0330101031321003-1231302233321013-1010200233223230-0131001100002232-1332202001321112"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the FastACLRule.

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

<a id="canonical-1002022102311210-0031222000222023-2332120222021112-0122110213232223-3332200300101330-0110021132021332-3013221311102322-0033131222333123"></a>

<a id="canonical-0221233231020012-3133023010021130-1302122221000300-1101020121322132-3003102331333113-3023222020213312-3222031221132033-3232333032030230"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ip_prefix_set](data-sources--fast_acl_rule--reference--group-001.md#canonical-1101103210103331-2133222331132211-1320311020203200-0013220233202103-3311203330111331-3102002130201021-0313302320013023-2233201133112203): complete subsection reference.

<a id="canonical-1010303010111033-0132003011011303-1101221233232210-3112223332223001-2030320031023001-2212030303313310-0031031003211131-2003103120221333"></a>

<a id="canonical-3021322330210032-0333232330230331-2011112101100323-2223200120310231-2322133100310000-0010032321121100-3100222131012031-2331101200012032"></a>

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

<a id="canonical-2123322331303200-1030003330023020-1022312210121102-3312331033120000-3021130331233030-0203130030301221-3131103121310210-3000212312023232"></a>

<a id="canonical-2003020101132323-2331031023011213-0331012330103330-0120332202010231-2333020300202133-1021010313002023-0230100122302003-3012222233122013"></a>

#### `name` property

Type: `"string"`. Required.

Name of the FastACLRule.

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

<a id="canonical-3312333312111322-2103133032223323-3232202202220032-3130133120210310-0113222310100200-2011011201311311-3321323012131212-3101120013120233"></a>

<a id="canonical-3101200332311310-2113112311112312-0212210212201202-2103031211002001-0131000213230322-0330233130032221-3232020202022021-3203202031210320"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the FastACLRule exists.

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

- [port](data-sources--fast_acl_rule--reference--group-001.md#canonical-1322120012333330-2001323120031212-1213301121230220-1010301132000011-0020301112212330-0130033133120321-3222021133213110-3013020312031213): complete subsection reference.

- [prefix](data-sources--fast_acl_rule--reference--group-001.md#canonical-2002033300333000-1223112020311322-1210001022020221-0123333022303330-0030013300310300-3210021023233013-1021201300130211-2203221310300103): complete subsection reference.

<a id="canonical-0112221130230023-2331101211310121-1210302311320022-1201103122320202-0222020112100000-0211012011113121-1222031120221230-3013110203012000"></a>

### All schema paths for `xcsh_fast_acl_rule`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `action` | [action](data-sources--fast_acl_rule--reference--group-001.md#canonical-3132100012300031-2111102333221022-2022322111111231-3031030232212331-1112233132200233-2122333101000230-0232011012311132-0030300110332030) |
| `action.policer_action` | [action.policer_action](data-sources--fast_acl_rule--reference--group-001.md#canonical-2210203220110122-3320001100202302-1323003330003030-0132010200311013-0100320302032322-1133010332021211-3102100130130200-1100221100103233) |
| `action.policer_action.ref` | [action.policer_action.ref](data-sources--fast_acl_rule--reference--group-001.md#canonical-1012022211033322-2203131101031032-0332213220001302-2032222211002030-0033133302311233-3323100213012031-2001320022103221-3311110311111232) |
| `action.policer_action.ref.kind` | [action.policer_action.ref.kind](data-sources--fast_acl_rule--reference--group-001.md#canonical-1233322020320230-1031323120212001-2222320132312200-2003303322010133-0120112033322132-0212010022211033-3113200132112120-0133301330201102) |
| `action.policer_action.ref.name` | [action.policer_action.ref.name](data-sources--fast_acl_rule--reference--group-001.md#canonical-3232121221030311-3023100210311001-0202103012321111-2303011021202030-1231313213311322-3132031103213022-3310222332332203-1030233010001132) |
| `action.policer_action.ref.namespace` | [action.policer_action.ref.namespace](data-sources--fast_acl_rule--reference--group-001.md#canonical-0132032233210231-2121101312321123-3322123121032012-1201003320101310-1321331231111013-1011012331222000-2020323033112132-2103321310201110) |
| `action.policer_action.ref.tenant` | [action.policer_action.ref.tenant](data-sources--fast_acl_rule--reference--group-001.md#canonical-3033012230130110-0031000212301220-2023211320220130-3022013120002003-2121212010102223-1100122033021321-1231323033102211-2003333023110110) |
| `action.policer_action.ref.uid` | [action.policer_action.ref.uid](data-sources--fast_acl_rule--reference--group-001.md#canonical-3232100031300010-2201020113210112-1012212221311321-3303011320301303-3312201312301133-2021302032300332-0213111122032112-1011013021202201) |
| `action.protocol_policer_action` | [action.protocol_policer_action](data-sources--fast_acl_rule--reference--group-001.md#canonical-3122200030021231-2133022320033303-3332210123313233-0212000333103122-0030233122223031-1112333112033310-0010211202322202-1211231021321030) |
| `action.protocol_policer_action.ref` | [action.protocol_policer_action.ref](data-sources--fast_acl_rule--reference--group-001.md#canonical-3121203020022020-2123120211311233-0030031031032312-0010021201222030-0310200312033331-3303002232220113-3212032002332210-0321023033110213) |
| `action.protocol_policer_action.ref.kind` | [action.protocol_policer_action.ref.kind](data-sources--fast_acl_rule--reference--group-001.md#canonical-0320320012222200-2302300113302213-1321233013101202-2130103223203213-2030330001022203-0213130311002322-3103133311222212-0222200211232220) |
| `action.protocol_policer_action.ref.name` | [action.protocol_policer_action.ref.name](data-sources--fast_acl_rule--reference--group-001.md#canonical-3201130331123320-2020213313221111-0111332022220231-3233100001232202-0102110211010212-0100202021312021-1012020032200320-1231102032320333) |
| `action.protocol_policer_action.ref.namespace` | [action.protocol_policer_action.ref.namespace](data-sources--fast_acl_rule--reference--group-001.md#canonical-3030222121212120-0222201102301203-3323111010111010-3210023232231231-2011232320101013-0133333222310220-1121101022320002-0221113100011331) |
| `action.protocol_policer_action.ref.tenant` | [action.protocol_policer_action.ref.tenant](data-sources--fast_acl_rule--reference--group-001.md#canonical-1323221030001230-3110122203101331-2211300222320133-0323131321111303-0211233211320233-0332012203332001-2103322123312010-3323233010302310) |
| `action.protocol_policer_action.ref.uid` | [action.protocol_policer_action.ref.uid](data-sources--fast_acl_rule--reference--group-001.md#canonical-2233211222000100-0330323020033113-2110223031113132-1022312231111320-2103302021100132-2101031330222010-2003112003231211-2203112233101131) |
| `action.simple_action` | [action.simple_action](data-sources--fast_acl_rule--reference--group-001.md#canonical-2211233301111033-1103123322303312-0122202221021112-3322111311232222-1331031212211301-1313131132020103-2221120113220330-3222113031302111) |
| `annotations` | [annotations](data-sources--fast_acl_rule--reference--group-001.md#canonical-3231311312011220-3321113331121132-0323222121323322-0112001222130322-2331323321210111-1120313320001223-3112032012003213-0233300210120032) |
| `description` | [description](data-sources--fast_acl_rule--reference--group-001.md#canonical-0012321002201302-2322132122213120-1101200021213131-2123002320133111-3113220010031033-2021131310233301-0312000103011320-1021031120311130) |
| `id` | [ID](data-sources--fast_acl_rule--reference--group-001.md#canonical-1002022102311210-0031222000222023-2332120222021112-0122110213232223-3332200300101330-0110021132021332-3013221311102322-0033131222333123) |
| `ip_prefix_set` | [ip_prefix_set](data-sources--fast_acl_rule--reference--group-001.md#canonical-2112330021232201-2130332311100001-3110222102110212-1311313021111123-0233301202131303-2012123011110101-1311033030111231-0113030122312132) |
| `ip_prefix_set.ref` | [ip_prefix_set.ref](data-sources--fast_acl_rule--reference--group-001.md#canonical-2011013333111123-2101123311320203-3211021313101013-2232331133220332-2100333010020133-1333313133130211-0023300101012331-0112113002132321) |
| `ip_prefix_set.ref.kind` | [ip_prefix_set.ref.kind](data-sources--fast_acl_rule--reference--group-001.md#canonical-2110332222203230-1221012233131301-2120021002220211-1100200133003200-2322303233231232-2030002232123303-3103100212000121-1110203023313312) |
| `ip_prefix_set.ref.name` | [ip_prefix_set.ref.name](data-sources--fast_acl_rule--reference--group-001.md#canonical-0010222200030033-2033303020133001-3222121011322003-2300330331112133-2120331310112102-1333111211202101-3323310011012331-0120201322001003) |
| `ip_prefix_set.ref.namespace` | [ip_prefix_set.ref.namespace](data-sources--fast_acl_rule--reference--group-001.md#canonical-2212230101100032-1230113330112210-0032321203123231-3323310302002320-0100122213111130-0011300312221030-0121223113223322-0331313010210223) |
| `ip_prefix_set.ref.tenant` | [ip_prefix_set.ref.tenant](data-sources--fast_acl_rule--reference--group-001.md#canonical-2103201303200033-0221113210131021-3233131103003032-1012132220020120-3232001103132133-3010002133000001-0120023310221130-0213313332202323) |
| `ip_prefix_set.ref.uid` | [ip_prefix_set.ref.uid](data-sources--fast_acl_rule--reference--group-001.md#canonical-3101330010320323-2230101321020133-3203132113130212-3130202013120032-0222131212300201-3222333032322033-2100310002202021-1330022221321020) |
| `labels` | [labels](data-sources--fast_acl_rule--reference--group-001.md#canonical-1010303010111033-0132003011011303-1101221233232210-3112223332223001-2030320031023001-2212030303313310-0031031003211131-2003103120221333) |
| `name` | [name](data-sources--fast_acl_rule--reference--group-001.md#canonical-2123322331303200-1030003330023020-1022312210121102-3312331033120000-3021130331233030-0203130030301221-3131103121310210-3000212312023232) |
| `namespace` | [namespace](data-sources--fast_acl_rule--reference--group-001.md#canonical-3312333312111322-2103133032223323-3232202202220032-3130133120210310-0113222310100200-2011011201311311-3321323012131212-3101120013120233) |
| `port` | [port](data-sources--fast_acl_rule--reference--group-001.md#canonical-3132200202102321-0133231233111120-3232011130032300-3320003113012310-1300312201012233-1030333200030321-2220123012203301-0032032231210221) |
| `port.all` | [port.all](data-sources--fast_acl_rule--reference--group-001.md#canonical-1210320211123233-1300220212211033-2230003313022033-2332011033213331-0010323230311020-3202020122332110-0222030332310122-0203321100012201) |
| `port.dns` | [port.dns](data-sources--fast_acl_rule--reference--group-001.md#canonical-1212302101332033-1000001020302210-2322021323222021-3001021210212120-1001233131311020-2132203303203212-3020211020121002-3002302003103323) |
| `port.user_defined` | [port.user_defined](data-sources--fast_acl_rule--reference--group-001.md#canonical-0220000221302021-1232102303200301-1010322331300300-0031323113121300-3103032133001301-1002302311322323-3123312010200021-2022110102322130) |
| `prefix` | [prefix](data-sources--fast_acl_rule--reference--group-001.md#canonical-1030231312313323-0102110311201310-3013101312122033-2220310232222003-2211000331012021-0100012132032231-1023202203301000-1310121100210210) |
| `prefix.prefix` | [prefix.prefix](data-sources--fast_acl_rule--reference--group-001.md#canonical-0031203103323113-1013133202301223-2100022302201201-2310133022202132-0123321320300323-2201001330013113-1102233231011211-2233110112330311) |

<a id="canonical-2323322103100113-0330221003132312-0022232121231222-2001112100311031-3033102332230020-2200221013103021-0101203130002020-0333000311113233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `action` properties

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-0130033020112332-1320113013031122-3323102231201130-2302010233220010-0332222213110322-3330103313001312-0313202331310301-0031300023203130)
- action

<a id="canonical-3132100012300031-2111102333221022-2022322111111231-3031030232212331-1112233132200233-2122333101000230-0232011012311132-0030300110332030"></a>

Type: `"single"`. Computed.

FastAclRuleAction specifies possible action to be applied on traffic, possible action include
dropping, forwarding or ratelimiting the traffic.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action": "[\"policer_action\",\"protocol_policer_action\",\"simple_action\"]"
}
```

<a id="canonical-3202031201322311-2232122230032021-3311123333022301-1101013312201323-3223332020333110-1031302032210322-1313230103010202-0302033320211122"></a>

### Direct properties for `action`

- [policer_action](data-sources--fast_acl_rule--reference--group-001.md#canonical-2021310101310002-1033112032132002-0200011021220002-3211012311301002-1233100310222003-1010323122122123-1211031000200313-2120233320201300): complete subsection reference.

- [protocol_policer_action](data-sources--fast_acl_rule--reference--group-001.md#canonical-1331132021312303-2001020311321111-3113220023310202-1300303201220003-1310301113012133-2310332200031303-2022022322311213-2322323010332332): complete subsection reference.

<a id="canonical-2211233301111033-1103123322303312-0122202221021112-3322111311232222-1331031212211301-1313131132020103-2221120113220330-3222113031302111"></a>

<a id="canonical-0131033333003203-2011302200220211-3301321031133110-1122111021331221-1101300203333000-2211022313332310-1312003200111030-3010000331331013"></a>

#### `action.simple_action` property

Type: `"string"`. Computed.

\[Enum: DENY|ALLOW\] FastAclRuleSimpleAction specifies simple action like PASS or DENY Drop the
traffic Forward the traffic. Possible values are \`DENY\`, \`ALLOW\`. Defaults to \`DENY\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "DENY",
  "enum": [
    "DENY",
    "ALLOW"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2021310101310002-1033112032132002-0200011021220002-3211012311301002-1233100310222003-1010323122122123-1211031000200313-2120233320201300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `action.policer_action` properties

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-0130033020112332-1320113013031122-3323102231201130-2302010233220010-0332222213110322-3330103313001312-0313202331310301-0031300023203130)
- [action](data-sources--fast_acl_rule--reference--group-001.md#canonical-2323322103100113-0330221003132312-0022232121231222-2001112100311031-3033102332230020-2200221013103021-0101203130002020-0333000311113233)
- action.policer_action

<a id="canonical-2210203220110122-3320001100202302-1323003330003030-0132010200311013-0100320302032322-1133010332021211-3102100130130200-1100221100103233"></a>

Type: `"single"`. Computed.

Policer Reference. Reference to policer object.

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

<a id="canonical-1300330220111111-0100102102213030-1011021100332000-0320303333233303-0111100330031212-1320132023032002-3000123010122311-1303230332012131"></a>

### Direct properties for `action.policer_action`

- [ref](data-sources--fast_acl_rule--reference--group-001.md#canonical-2101211300001002-0103202212000023-1203000223003113-3210022201001231-0012222301132103-3120002002322311-2330111200102120-0212032023230202): complete subsection reference.

<a id="canonical-2101211300001002-0103202212000023-1203000223003113-3210022201001231-0012222301132103-3120002002322311-2330111200102120-0212032023230202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `action.policer_action.ref` properties

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-0130033020112332-1320113013031122-3323102231201130-2302010233220010-0332222213110322-3330103313001312-0313202331310301-0031300023203130)
- [action](data-sources--fast_acl_rule--reference--group-001.md#canonical-2323322103100113-0330221003132312-0022232121231222-2001112100311031-3033102332230020-2200221013103021-0101203130002020-0333000311113233)
- [action.policer_action](data-sources--fast_acl_rule--reference--group-001.md#canonical-2021310101310002-1033112032132002-0200011021220002-3211012311301002-1233100310222003-1010323122122123-1211031000200313-2120233320201300)
- action.policer_action.ref

<a id="canonical-1012022211033322-2203131101031032-0332213220001302-2032222211002030-0033133302311233-3323100213012031-2001320022103221-3311110311111232"></a>

Type: `"list"`. Computed.

Reference. A policer direct reference.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-0302123233011303-3222131330313003-1022221113020313-1022233220021221-2223302333021230-2320230321011310-0132013211020313-1123123113130230"></a>

### Direct properties for `action.policer_action.ref`

<a id="canonical-1233322020320230-1031323120212001-2222320132312200-2003303322010133-0120112033322132-0212010022211033-3113200132112120-0133301330201102"></a>

#### `action.policer_action.ref.kind` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3232121221030311-3023100210311001-0202103012321111-2303011021202030-1231313213311322-3132031103213022-3310222332332203-1030233010001132"></a>

<a id="canonical-1331031203200203-3220333201103201-1002133311311311-1321123210130020-1203303201201100-3000023120012131-2222323203220131-3100303211201331"></a>

#### `action.policer_action.ref.name` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0132032233210231-2121101312321123-3322123121032012-1201003320101310-1321331231111013-1011012331222000-2020323033112132-2103321310201110"></a>

<a id="canonical-0001131023332133-1210020323030213-2210000333211321-2312111102232100-1132303032200020-1230202330220122-2122011233133022-2002201300102333"></a>

#### `action.policer_action.ref.namespace` property

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

<a id="canonical-3033012230130110-0031000212301220-2023211320220130-3022013120002003-2121212010102223-1100122033021321-1231323033102211-2003333023110110"></a>

<a id="canonical-0322220003210332-3220202232131232-1303121233202203-1222230312231202-3110102311012112-2320122320320312-1330212113032201-3201313110233221"></a>

#### `action.policer_action.ref.tenant` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3232100031300010-2201020113210112-1012212221311321-3303011320301303-3312201312301133-2021302032300332-0213111122032112-1011013021202201"></a>

<a id="canonical-0012113021031133-2030223100113011-2120021213123201-1130102322320032-2223030131121213-3132032233112111-1013130103130033-1003003202101021"></a>

#### `action.policer_action.ref.uid` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1331132021312303-2001020311321111-3113220023310202-1300303201220003-1310301113012133-2310332200031303-2022022322311213-2322323010332332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `action.protocol_policer_action` properties

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-0130033020112332-1320113013031122-3323102231201130-2302010233220010-0332222213110322-3330103313001312-0313202331310301-0031300023203130)
- [action](data-sources--fast_acl_rule--reference--group-001.md#canonical-2323322103100113-0330221003132312-0022232121231222-2001112100311031-3033102332230020-2200221013103021-0101203130002020-0333000311113233)
- action.protocol_policer_action

<a id="canonical-3122200030021231-2133022320033303-3332210123313233-0212000333103122-0030233122223031-1112333112033310-0010211202322202-1211231021321030"></a>

Type: `"single"`. Computed.

Protocol Policer Reference. Reference to policer object.

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

<a id="canonical-2103320301133031-2313202232312121-1233122211303031-0320222322212210-2313203103031003-2020333321301212-2022232232330313-1130132012213103"></a>

### Direct properties for `action.protocol_policer_action`

- [ref](data-sources--fast_acl_rule--reference--group-001.md#canonical-1012101013330133-0032220210000112-0201022201002212-1012011030030200-3000201132220220-1100122113203131-3301220223131311-0002321100310310): complete subsection reference.

<a id="canonical-1012101013330133-0032220210000112-0201022201002212-1012011030030200-3000201132220220-1100122113203131-3301220223131311-0002321100310310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `action.protocol_policer_action.ref` properties

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-0130033020112332-1320113013031122-3323102231201130-2302010233220010-0332222213110322-3330103313001312-0313202331310301-0031300023203130)
- [action](data-sources--fast_acl_rule--reference--group-001.md#canonical-2323322103100113-0330221003132312-0022232121231222-2001112100311031-3033102332230020-2200221013103021-0101203130002020-0333000311113233)
- [action.protocol_policer_action](data-sources--fast_acl_rule--reference--group-001.md#canonical-1331132021312303-2001020311321111-3113220023310202-1300303201220003-1310301113012133-2310332200031303-2022022322311213-2322323010332332)
- action.protocol_policer_action.ref

<a id="canonical-3121203020022020-2123120211311233-0030031031032312-0010021201222030-0310200312033331-3303002232220113-3212032002332210-0321023033110213"></a>

Type: `"list"`. Computed.

Protocol policer Reference. Reference to protocol policer object.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-0233031233323213-2100332031132320-0333200011313201-1220320213013021-2221022001113300-3110333323222223-1110211302233110-1020100013312133"></a>

### Direct properties for `action.protocol_policer_action.ref`

<a id="canonical-0320320012222200-2302300113302213-1321233013101202-2130103223203213-2030330001022203-0213130311002322-3103133311222212-0222200211232220"></a>

#### `action.protocol_policer_action.ref.kind` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3201130331123320-2020213313221111-0111332022220231-3233100001232202-0102110211010212-0100202021312021-1012020032200320-1231102032320333"></a>

<a id="canonical-3113200300211300-0230331033013233-3211022010131131-2120210020230103-1333302310020010-2120122120233102-2330130232310023-2232001031330310"></a>

#### `action.protocol_policer_action.ref.name` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3030222121212120-0222201102301203-3323111010111010-3210023232231231-2011232320101013-0133333222310220-1121101022320002-0221113100011331"></a>

<a id="canonical-3133110203330033-1121303030133210-1313121001312111-2003302210331000-3331102320202000-3011111312231132-0302023020203133-0032031302311031"></a>

#### `action.protocol_policer_action.ref.namespace` property

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

<a id="canonical-1323221030001230-3110122203101331-2211300222320133-0323131321111303-0211233211320233-0332012203332001-2103322123312010-3323233010302310"></a>

<a id="canonical-2333232001130031-1110021023301233-3211120033121013-2002221221231203-1232132030313311-3220230331021121-3102030220313020-0031000020121330"></a>

#### `action.protocol_policer_action.ref.tenant` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2233211222000100-0330323020033113-2110223031113132-1022312231111320-2103302021100132-2101031330222010-2003112003231211-2203112233101131"></a>

<a id="canonical-2033111100031033-2230122322220113-3011002032221330-0132323212102312-3313301300010231-3033213131322122-2133320013113000-2103113312303331"></a>

#### `action.protocol_policer_action.ref.uid` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1101103210103331-2133222331132211-1320311020203200-0013220233202103-3311203330111331-3102002130201021-0313302320013023-2233201133112203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ip_prefix_set` properties

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-0130033020112332-1320113013031122-3323102231201130-2302010233220010-0332222213110322-3330103313001312-0313202331310301-0031300023203130)
- ip_prefix_set

<a id="canonical-2112330021232201-2130332311100001-3110222102110212-1311313021111123-0233301202131303-2012123011110101-1311033030111231-0113030122312132"></a>

Type: `"single"`. Computed.

\[OneOf: ip\_prefix\_set, prefix\] List of references to ip\_prefix\_set objects.

Additional upstream details:

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

OneOf alternatives in this subsection:

- [ip_prefix_set](data-sources--fast_acl_rule--reference--group-001.md#canonical-2112330021232201-2130332311100001-3110222102110212-1311313021111123-0233301202131303-2012123011110101-1311033030111231-0113030122312132)
- [prefix](data-sources--fast_acl_rule--reference--group-001.md#canonical-1030231312313323-0102110311201310-3013101312122033-2220310232222003-2211000331012021-0100012132032231-1023202203301000-1310121100210210)

Select alternatives according to the provider validators above.

<a id="canonical-2103113110012320-2301121110232112-3210300130001213-2323032200321333-2130120312112111-1210023103323010-1002030031012132-3123030102033120"></a>

### Direct properties for `ip_prefix_set`

- [ref](data-sources--fast_acl_rule--reference--group-001.md#canonical-0002230311211332-0003201210210103-1021122213201333-0300033323223111-1003200213003133-0112020311113311-2330111022311110-3113310211201100): complete subsection reference.

<a id="canonical-0002230311211332-0003201210210103-1021122213201333-0300033323223111-1003200213003133-0112020311113311-2330111022311110-3113310211201100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ip_prefix_set.ref` properties

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-0130033020112332-1320113013031122-3323102231201130-2302010233220010-0332222213110322-3330103313001312-0313202331310301-0031300023203130)
- [ip_prefix_set](data-sources--fast_acl_rule--reference--group-001.md#canonical-1101103210103331-2133222331132211-1320311020203200-0013220233202103-3311203330111331-3102002130201021-0313302320013023-2233201133112203)
- ip_prefix_set.ref

<a id="canonical-2011013333111123-2101123311320203-3211021313101013-2232331133220332-2100333010020133-1333313133130211-0023300101012331-0112113002132321"></a>

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-0120233310302230-1020222313100201-3230312113010121-0101301031310111-0232203212220211-1322011222232112-1321101101030313-1112302300312133"></a>

### Direct properties for `ip_prefix_set.ref`

<a id="canonical-2110332222203230-1221012233131301-2120021002220211-1100200133003200-2322303233231232-2030002232123303-3103100212000121-1110203023313312"></a>

#### `ip_prefix_set.ref.kind` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0010222200030033-2033303020133001-3222121011322003-2300330331112133-2120331310112102-1333111211202101-3323310011012331-0120201322001003"></a>

<a id="canonical-1203132310132312-1201232220312013-2232303300301320-1233023111320011-2102112333000103-2312010232113130-0312001321132130-3211012232103211"></a>

#### `ip_prefix_set.ref.name` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2212230101100032-1230113330112210-0032321203123231-3323310302002320-0100122213111130-0011300312221030-0121223113223322-0331313010210223"></a>

<a id="canonical-1033300202011322-1111232203313012-0113112003030100-1210120120130021-1110010101200312-3210231323301210-1000102202012000-1311132013133033"></a>

#### `ip_prefix_set.ref.namespace` property

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

<a id="canonical-2103201303200033-0221113210131021-3233131103003032-1012132220020120-3232001103132133-3010002133000001-0120023310221130-0213313332202323"></a>

<a id="canonical-2203003011103311-0130023033330312-1102232233101203-0100012301112223-0031110202010021-2303333023323133-1122112311010133-0202033310000003"></a>

#### `ip_prefix_set.ref.tenant` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3101330010320323-2230101321020133-3203132113130212-3130202013120032-0222131212300201-3222333032322033-2100310002202021-1330022221321020"></a>

<a id="canonical-2222333322210322-1001332012302230-3010123001301031-1233110300123323-0033111130233002-2033331002201003-2011221323121323-1121211222220132"></a>

#### `ip_prefix_set.ref.uid` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1322120012333330-2001323120031212-1213301121230220-1010301132000011-0020301112212330-0130033133120321-3222021133213110-3013020312031213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `port` properties

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-0130033020112332-1320113013031122-3323102231201130-2302010233220010-0332222213110322-3330103313001312-0313202331310301-0031300023203130)
- port

<a id="canonical-3132200202102321-0133231233111120-3232011130032300-3320003113012310-1300312201012233-1030333200030321-2220123012203301-0032032231210221"></a>

Type: `"list"`. Computed.

Source Ports. L4 port numbers to match.

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
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-0323231232002021-1122200302232110-0303113113021212-1111220022220330-2322203220133101-1231310130320020-3123110223012223-1130131220112032"></a>

### Direct properties for `port`

- [all](data-sources--fast_acl_rule--reference--group-001.md#canonical-0301031232230303-3322031021121223-0312023300121031-2000212223302202-0211231110121113-1300133132210132-0320322330302331-3122333010310123): complete subsection reference.

- [DNS](data-sources--fast_acl_rule--reference--group-001.md#canonical-2223232332312230-0332113221132320-0012131123212223-1030331012201002-3031212032312222-2203330131020122-3331200210103323-3321012113103333): complete subsection reference.

<a id="canonical-0220000221302021-1232102303200301-1010322331300300-0031323113121300-3103032133001301-1002302311322323-3123312010200021-2022110102322130"></a>

<a id="canonical-2223033232221233-1003232013102132-1110120101212110-3313313012111201-2223312112323032-3300003131131302-3303210232012001-2300210333201221"></a>

#### `port.user_defined` property

Type: `"number"`. Computed.

Exclusive with \[all DNS\] Matches the user defined port.

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

<a id="canonical-0301031232230303-3322031021121223-0312023300121031-2000212223302202-0211231110121113-1300133132210132-0320322330302331-3122333010310123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `port.all` properties

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-0130033020112332-1320113013031122-3323102231201130-2302010233220010-0332222213110322-3330103313001312-0313202331310301-0031300023203130)
- [port](data-sources--fast_acl_rule--reference--group-001.md#canonical-1322120012333330-2001323120031212-1213301121230220-1010301132000011-0020301112212330-0130033133120321-3222021133213110-3013020312031213)
- port.all

<a id="canonical-1210320211123233-1300220212211033-2230003313022033-2332011033213331-0010323230311020-3202020122332110-0222030332310122-0203321100012201"></a>

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

<a id="canonical-2223232332312230-0332113221132320-0012131123212223-1030331012201002-3031212032312222-2203330131020122-3331200210103323-3321012113103333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `port.dns` properties

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-0130033020112332-1320113013031122-3323102231201130-2302010233220010-0332222213110322-3330103313001312-0313202331310301-0031300023203130)
- [port](data-sources--fast_acl_rule--reference--group-001.md#canonical-1322120012333330-2001323120031212-1213301121230220-1010301132000011-0020301112212330-0130033133120321-3222021133213110-3013020312031213)
- port.DNS

<a id="canonical-1212302101332033-1000001020302210-2322021323222021-3001021210212120-1001233131311020-2132203303203212-3020211020121002-3002302003103323"></a>

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

<a id="canonical-2002033300333000-1223112020311322-1210001022020221-0123333022303330-0030013300310300-3210021023233013-1021201300130211-2203221310300103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `prefix` properties

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-0130033020112332-1320113013031122-3323102231201130-2302010233220010-0332222213110322-3330103313001312-0313202331310301-0031300023203130)
- prefix

<a id="canonical-1030231312313323-0102110311201310-3013101312122033-2220310232222003-2211000331012021-0100012132032231-1023202203301000-1310121100210210"></a>

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

<a id="canonical-2111322223131201-2000211022201300-3321313210031001-2202312332312033-3233130223033200-3223322330321201-3300211303112101-1033312001033311"></a>

### Direct properties for `prefix`

<a id="canonical-0031203103323113-1013133202301223-2100022302201201-2310133022202132-0123321320300323-2201001330013113-1102233231011211-2233110112330311"></a>

#### `prefix.prefix` property

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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```
