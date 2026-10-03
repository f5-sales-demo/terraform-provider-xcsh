---
page_title: "xcsh_fast_acl_rule reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fast_acl_rule reference."
---

# xcsh_fast_acl_rule reference

<a id="canonical-0130033020112332-1320113013031122-3323102231201130-2302010233220010-0332222213110322-3330103313001312-0313202331310301-0031300023203130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2112202213021033-1033233213220221-1133303320010003-2223302231333310-3231130010220310-3100212021023213-1132200332232023-1103021121312121"></a>

## Property reference — Property reference / 032131121302 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)
- Property reference

<a id="canonical-0021210132033321-1101033331303321-0231100312323102-0322203221010033-0120121322023322-1101120232331220-1132111002320031-2111131310032231"></a>

## Direct properties — Property reference / 032131121302 / 3

- [action](data-sources--fast_acl_rule--reference--group-001.md#canonical-2323322103100113-0330221003132312-0022232121231222-2001112100311031-3033102332230020-2200221013103021-0101203130002020-0333000311113233): complete subsection reference.

<a id="canonical-3231311312011220-3321113331121132-0323222121323322-0112001222130322-2331323321210111-1120313320001223-3112032012003213-0233300210120032"></a>

<a id="canonical-0012100210201120-0222333223210033-0033132230012023-0330101031321003-1231302233321013-1010200233223230-0131001100002232-1332202001321112"></a>

## annotations property — Property reference / 032131121302 / 4

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

<a id="canonical-0012321002201302-2322132122213120-1101200021213131-2123002320133111-3113220010031033-2021131310233301-0312000103011320-1021031120311130"></a>

<a id="canonical-0221233231020012-3133023010021130-1302122221000300-1101020121322132-3003102331333113-3023222020213312-3222031221132033-3232333032030230"></a>

## description property — Property reference / 032131121302 / 5

Type: `"string"`. Computed.

Description of the FastACLRule.

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

<a id="canonical-1002022102311210-0031222000222023-2332120222021112-0122110213232223-3332200300101330-0110021132021332-3013221311102322-0033131222333123"></a>

<a id="canonical-3021322330210032-0333232330230331-2011112101100323-2223200120310231-2322133100310000-0010032321121100-3100222131012031-2331101200012032"></a>

## ID property — Property reference / 032131121302 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ip_prefix_set](data-sources--fast_acl_rule--reference--group-001.md#canonical-1101103210103331-2133222331132211-1320311020203200-0013220233202103-3311203330111331-3102002130201021-0313302320013023-2233201133112203): complete subsection reference.

<a id="canonical-1010303010111033-0132003011011303-1101221233232210-3112223332223001-2030320031023001-2212030303313310-0031031003211131-2003103120221333"></a>

<a id="canonical-2003020101132323-2331031023011213-0331012330103330-0120332202010231-2333020300202133-1021010313002023-0230100122302003-3012222233122013"></a>

## labels property — Property reference / 032131121302 / 7

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

<a id="canonical-2123322331303200-1030003330023020-1022312210121102-3312331033120000-3021130331233030-0203130030301221-3131103121310210-3000212312023232"></a>

<a id="canonical-3101200332311310-2113112311112312-0212210212201202-2103031211002001-0131000213230322-0330233130032221-3232020202022021-3203202031210320"></a>

## name property — Property reference / 032131121302 / 8

Type: `"string"`. Required.

Name of the FastACLRule.

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

<a id="canonical-3312333312111322-2103133032223323-3232202202220032-3130133120210310-0113222310100200-2011011201311311-3321323012131212-3101120013120233"></a>

<a id="canonical-0112221130230023-2331101211310121-1210302311320022-1201103122320202-0222020112100000-0211012011113121-1222031120221230-3013110203012000"></a>

## namespace property — Property reference / 032131121302 / 9

Type: `"string"`. Required.

Namespace where the FastACLRule exists.

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

- [port](data-sources--fast_acl_rule--reference--group-001.md#canonical-1322120012333330-2001323120031212-1213301121230220-1010301132000011-0020301112212330-0130033133120321-3222021133213110-3013020312031213): complete subsection reference.

- [prefix](data-sources--fast_acl_rule--reference--group-001.md#canonical-2002033300333000-1223112020311322-1210001022020221-0123333022303330-0030013300310300-3210021023233013-1021201300130211-2203221310300103): complete subsection reference.

<a id="canonical-1233032222120311-1230220123210322-3330312333223131-2112100131223122-2223303123202130-1121320032203132-1321002200223131-3121032200310331"></a>

## All schema paths — Property reference / 032131121302 / 10

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

<a id="canonical-3222122313003102-0232223332220312-2003301011331301-1132222012012101-2231323013012320-0200202323001320-0322103023200023-0012300011010200"></a>

## Next pages — Property reference / 032131121302 / 11

- [action](data-sources--fast_acl_rule--reference--group-001.md#canonical-2323322103100113-0330221003132312-0022232121231222-2001112100311031-3033102332230020-2200221013103021-0101203130002020-0333000311113233)
- [ip_prefix_set](data-sources--fast_acl_rule--reference--group-001.md#canonical-1101103210103331-2133222331132211-1320311020203200-0013220233202103-3311203330111331-3102002130201021-0313302320013023-2233201133112203)
- [port](data-sources--fast_acl_rule--reference--group-001.md#canonical-1322120012333330-2001323120031212-1213301121230220-1010301132000011-0020301112212330-0130033133120321-3222021133213110-3013020312031213)
- [prefix](data-sources--fast_acl_rule--reference--group-001.md#canonical-2002033300333000-1223112020311322-1210001022020221-0123333022303330-0030013300310300-3210021023233013-1021201300130211-2203221310300103)
- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)

<a id="canonical-2323322103100113-0330221003132312-0022232121231222-2001112100311031-3033102332230020-2200221013103021-0101203130002020-0333000311113233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202031201322311-2232122230032021-3311123333022301-1101013312201323-3223332020333110-1031302032210322-1313230103010202-0302033320211122"></a>

## action — action / 330101112033 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-0130033020112332-1320113013031122-3323102231201130-2302010233220010-0332222213110322-3330103313001312-0313202331310301-0031300023203130)
- action

<a id="canonical-3132100012300031-2111102333221022-2022322111111231-3031030232212331-1112233132200233-2122333101000230-0232011012311132-0030300110332030"></a>

Type: `"single"`. Computed.

FastAclRuleAction specifies possible action to be applied on traffic, possible action include
dropping, forwarding or ratelimiting the traffic.

Upstream description:

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

<a id="canonical-0131033333003203-2011302200220211-3301321031133110-1122111021331221-1101300203333000-2211022313332310-1312003200111030-3010000331331013"></a>

## Direct properties — action / 330101112033 / 3

- [policer_action](data-sources--fast_acl_rule--reference--group-001.md#canonical-2021310101310002-1033112032132002-0200011021220002-3211012311301002-1233100310222003-1010323122122123-1211031000200313-2120233320201300): complete subsection reference.

- [protocol_policer_action](data-sources--fast_acl_rule--reference--group-001.md#canonical-1331132021312303-2001020311321111-3113220023310202-1300303201220003-1310301113012133-2310332200031303-2022022322311213-2322323010332332): complete subsection reference.

<a id="canonical-2211233301111033-1103123322303312-0122202221021112-3322111311232222-1331031212211301-1313131132020103-2221120113220330-3222113031302111"></a>

<a id="canonical-3313023302310322-2230012331323100-0221310332311303-0021311312320100-3311313103033003-0230231200311123-3330201331301212-3103132332311131"></a>

## simple_action property — action / 330101112033 / 4

Type: `"string"`. Computed.

\[Enum: DENY|ALLOW\] FastAclRuleSimpleAction specifies simple action like PASS or DENY Drop the
traffic Forward the traffic. Possible values are \`DENY\`, \`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

FastAclRuleSimpleAction specifies simple action like PASS or DENY

Drop the traffic Forward the traffic.

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

<a id="canonical-1210233121311322-1100301331313322-1221311201123013-1113332222322100-3022323212320033-0203223201202333-1221113022021303-3032203200101020"></a>

## Next pages — action / 330101112033 / 5

- [action.policer_action](data-sources--fast_acl_rule--reference--group-001.md#canonical-2021310101310002-1033112032132002-0200011021220002-3211012311301002-1233100310222003-1010323122122123-1211031000200313-2120233320201300)
- [action.protocol_policer_action](data-sources--fast_acl_rule--reference--group-001.md#canonical-1331132021312303-2001020311321111-3113220023310202-1300303201220003-1310301113012133-2310332200031303-2022022322311213-2322323010332332)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-0130033020112332-1320113013031122-3323102231201130-2302010233220010-0332222213110322-3330103313001312-0313202331310301-0031300023203130)
- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)

<a id="canonical-2021310101310002-1033112032132002-0200011021220002-3211012311301002-1233100310222003-1010323122122123-1211031000200313-2120233320201300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300330220111111-0100102102213030-1011021100332000-0320303333233303-0111100330031212-1320132023032002-3000123010122311-1303230332012131"></a>

## action.policer_action — policer_action / 330030230133 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-0130033020112332-1320113013031122-3323102231201130-2302010233220010-0332222213110322-3330103313001312-0313202331310301-0031300023203130)
- [action](data-sources--fast_acl_rule--reference--group-001.md#canonical-2323322103100113-0330221003132312-0022232121231222-2001112100311031-3033102332230020-2200221013103021-0101203130002020-0333000311113233)
- action.policer_action

<a id="canonical-2210203220110122-3320001100202302-1323003330003030-0132010200311013-0100320302032322-1133010332021211-3102100130130200-1100221100103233"></a>

Type: `"single"`. Computed.

Policer Reference. Reference to policer object.

Upstream description:

Reference to policer object.

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

<a id="canonical-0311030113101332-1321023313121022-3331013003123220-0323212021213023-1330000101212202-3202121013010303-2303203033221030-0300003221110033"></a>

## Direct properties — policer_action / 330030230133 / 3

- [ref](data-sources--fast_acl_rule--reference--group-001.md#canonical-2101211300001002-0103202212000023-1203000223003113-3210022201001231-0012222301132103-3120002002322311-2330111200102120-0212032023230202): complete subsection reference.

<a id="canonical-3031212230132021-1320322312332330-1100133223010320-1021000000120301-2122010320021020-1130230200120330-0202332333113320-0132012232110312"></a>

## Next pages — policer_action / 330030230133 / 4

- [action.policer_action.ref](data-sources--fast_acl_rule--reference--group-001.md#canonical-2101211300001002-0103202212000023-1203000223003113-3210022201001231-0012222301132103-3120002002322311-2330111200102120-0212032023230202)
- [action](data-sources--fast_acl_rule--reference--group-001.md#canonical-2323322103100113-0330221003132312-0022232121231222-2001112100311031-3033102332230020-2200221013103021-0101203130002020-0333000311113233)
- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)

<a id="canonical-2101211300001002-0103202212000023-1203000223003113-3210022201001231-0012222301132103-3120002002322311-2330111200102120-0212032023230202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302123233011303-3222131330313003-1022221113020313-1022233220021221-2223302333021230-2320230321011310-0132013211020313-1123123113130230"></a>

## action.policer_action.ref — ref / 033331201322 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-0130033020112332-1320113013031122-3323102231201130-2302010233220010-0332222213110322-3330103313001312-0313202331310301-0031300023203130)
- [action](data-sources--fast_acl_rule--reference--group-001.md#canonical-2323322103100113-0330221003132312-0022232121231222-2001112100311031-3033102332230020-2200221013103021-0101203130002020-0333000311113233)
- [action.policer_action](data-sources--fast_acl_rule--reference--group-001.md#canonical-2021310101310002-1033112032132002-0200011021220002-3211012311301002-1233100310222003-1010323122122123-1211031000200313-2120233320201300)
- action.policer_action.ref

<a id="canonical-1012022211033322-2203131101031032-0332213220001302-2032222211002030-0033133302311233-3323100213012031-2001320022103221-3311110311111232"></a>

Type: `"list"`. Computed.

Reference. A policer direct reference.

Upstream description:

A policer direct reference.

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

<a id="canonical-1331031203200203-3220333201103201-1002133311311311-1321123210130020-1203303201201100-3000023120012131-2222323203220131-3100303211201331"></a>

## Direct properties — ref / 033331201322 / 3

<a id="canonical-1233322020320230-1031323120212001-2222320132312200-2003303322010133-0120112033322132-0212010022211033-3113200132112120-0133301330201102"></a>

<a id="canonical-0001131023332133-1210020323030213-2210000333211321-2312111102232100-1132303032200020-1230202330220122-2122011233133022-2002201300102333"></a>

## kind property — ref / 033331201322 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-0322220003210332-3220202232131232-1303121233202203-1222230312231202-3110102311012112-2320122320320312-1330212113032201-3201313110233221"></a>

## name property — ref / 033331201322 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-0012113021031133-2030223100113011-2120021213123201-1130102322320032-2223030131121213-3132032233112111-1013130103130033-1003003202101021"></a>

## namespace property — ref / 033331201322 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-2120221023220101-3210212302103313-0232223201320010-1032123311011101-3003012023011330-2212211212303303-2223303011032302-1011303011321123"></a>

## tenant property — ref / 033331201322 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-1102030010211201-2310133010022200-1322022211333113-1132130232100111-2111302120003120-3302123200011021-3021123022213022-0011212032222232"></a>

## uid property — ref / 033331201322 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-1303103222031130-0100032200122313-3110132122320123-3232112211013011-0202120122012110-0132022333220233-1023302003222000-3220213313123330"></a>

## Next pages — ref / 033331201322 / 9

- [action.policer_action](data-sources--fast_acl_rule--reference--group-001.md#canonical-2021310101310002-1033112032132002-0200011021220002-3211012311301002-1233100310222003-1010323122122123-1211031000200313-2120233320201300)
- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)

<a id="canonical-1331132021312303-2001020311321111-3113220023310202-1300303201220003-1310301113012133-2310332200031303-2022022322311213-2322323010332332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103320301133031-2313202232312121-1233122211303031-0320222322212210-2313203103031003-2020333321301212-2022232232330313-1130132012213103"></a>

## action.protocol_policer_action — protocol_policer_action / 032022321103 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-0130033020112332-1320113013031122-3323102231201130-2302010233220010-0332222213110322-3330103313001312-0313202331310301-0031300023203130)
- [action](data-sources--fast_acl_rule--reference--group-001.md#canonical-2323322103100113-0330221003132312-0022232121231222-2001112100311031-3033102332230020-2200221013103021-0101203130002020-0333000311113233)
- action.protocol_policer_action

<a id="canonical-3122200030021231-2133022320033303-3332210123313233-0212000333103122-0030233122223031-1112333112033310-0010211202322202-1211231021321030"></a>

Type: `"single"`. Computed.

Protocol Policer Reference. Reference to policer object.

Upstream description:

Reference to policer object.

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

<a id="canonical-0012110313321100-0133120100030121-1221101211211020-2021303120311103-1001202232131211-0002310003233132-3002301012200111-0232032232303012"></a>

## Direct properties — protocol_policer_action / 032022321103 / 3

- [ref](data-sources--fast_acl_rule--reference--group-001.md#canonical-1012101013330133-0032220210000112-0201022201002212-1012011030030200-3000201132220220-1100122113203131-3301220223131311-0002321100310310): complete subsection reference.

<a id="canonical-0113333101022231-0300321310200230-0133202012132110-3100331121333212-3122213313032020-2032033302200330-2101001112332120-3132111201331021"></a>

## Next pages — protocol_policer_action / 032022321103 / 4

- [action.protocol_policer_action.ref](data-sources--fast_acl_rule--reference--group-001.md#canonical-1012101013330133-0032220210000112-0201022201002212-1012011030030200-3000201132220220-1100122113203131-3301220223131311-0002321100310310)
- [action](data-sources--fast_acl_rule--reference--group-001.md#canonical-2323322103100113-0330221003132312-0022232121231222-2001112100311031-3033102332230020-2200221013103021-0101203130002020-0333000311113233)
- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)

<a id="canonical-1012101013330133-0032220210000112-0201022201002212-1012011030030200-3000201132220220-1100122113203131-3301220223131311-0002321100310310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0233031233323213-2100332031132320-0333200011313201-1220320213013021-2221022001113300-3110333323222223-1110211302233110-1020100013312133"></a>

## action.protocol_policer_action.ref — ref / 001333123201 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-0130033020112332-1320113013031122-3323102231201130-2302010233220010-0332222213110322-3330103313001312-0313202331310301-0031300023203130)
- [action](data-sources--fast_acl_rule--reference--group-001.md#canonical-2323322103100113-0330221003132312-0022232121231222-2001112100311031-3033102332230020-2200221013103021-0101203130002020-0333000311113233)
- [action.protocol_policer_action](data-sources--fast_acl_rule--reference--group-001.md#canonical-1331132021312303-2001020311321111-3113220023310202-1300303201220003-1310301113012133-2310332200031303-2022022322311213-2322323010332332)
- action.protocol_policer_action.ref

<a id="canonical-3121203020022020-2123120211311233-0030031031032312-0010021201222030-0310200312033331-3303002232220113-3212032002332210-0321023033110213"></a>

Type: `"list"`. Computed.

Protocol policer Reference. Reference to protocol policer object.

Upstream description:

Reference to protocol policer object.

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

<a id="canonical-3113200300211300-0230331033013233-3211022010131131-2120210020230103-1333302310020010-2120122120233102-2330130232310023-2232001031330310"></a>

## Direct properties — ref / 001333123201 / 3

<a id="canonical-0320320012222200-2302300113302213-1321233013101202-2130103223203213-2030330001022203-0213130311002322-3103133311222212-0222200211232220"></a>

<a id="canonical-3133110203330033-1121303030133210-1313121001312111-2003302210331000-3331102320202000-3011111312231132-0302023020203133-0032031302311031"></a>

## kind property — ref / 001333123201 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-2333232001130031-1110021023301233-3211120033121013-2002221221231203-1232132030313311-3220230331021121-3102030220313020-0031000020121330"></a>

## name property — ref / 001333123201 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-2033111100031033-2230122322220113-3011002032221330-0132323212102312-3313301300010231-3033213131322122-2133320013113000-2103113312303331"></a>

## namespace property — ref / 001333123201 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-3200011112021322-3132113220231303-1322331012221302-1013311200222110-1102331020003301-0011220222233332-1102210000222323-0023013023312103"></a>

## tenant property — ref / 001333123201 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-1213301311313310-1001330010312023-0333220133111301-2002303133112020-1232110223330131-1301130321230222-0010333033233121-0310011330112033"></a>

## uid property — ref / 001333123201 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-1202323001313312-2001231313202120-3222103320312023-1010012203330211-2321312302212321-3122131031112233-3332130012200012-0302000131100232"></a>

## Next pages — ref / 001333123201 / 9

- [action.protocol_policer_action](data-sources--fast_acl_rule--reference--group-001.md#canonical-1331132021312303-2001020311321111-3113220023310202-1300303201220003-1310301113012133-2310332200031303-2022022322311213-2322323010332332)
- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)

<a id="canonical-1101103210103331-2133222331132211-1320311020203200-0013220233202103-3311203330111331-3102002130201021-0313302320013023-2233201133112203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103113110012320-2301121110232112-3210300130001213-2323032200321333-2130120312112111-1210023103323010-1002030031012132-3123030102033120"></a>

## ip_prefix_set — ip_prefix_set / 201331312100 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-0130033020112332-1320113013031122-3323102231201130-2302010233220010-0332222213110322-3330103313001312-0313202331310301-0031300023203130)
- ip_prefix_set

<a id="canonical-2112330021232201-2130332311100001-3110222102110212-1311313021111123-0233301202131303-2012123011110101-1311033030111231-0113030122312132"></a>

Type: `"single"`. Computed.

\[OneOf: ip\_prefix\_set, prefix\] List of references to ip\_prefix\_set objects.

Upstream description:

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

<a id="canonical-0032310300003221-1231122122223110-0022233200210311-2010012302220332-2220302213331310-3003201311220223-0130001213333333-3032033310133233"></a>

## Direct properties — ip_prefix_set / 201331312100 / 3

- [ref](data-sources--fast_acl_rule--reference--group-001.md#canonical-0002230311211332-0003201210210103-1021122213201333-0300033323223111-1003200213003133-0112020311113311-2330111022311110-3113310211201100): complete subsection reference.

<a id="canonical-0332120120132021-0030232133013202-2331202133331013-2002303011101311-0002010223101123-2100311310223123-1311013223102203-1323112033312210"></a>

## Next pages — ip_prefix_set / 201331312100 / 4

- [ip_prefix_set.ref](data-sources--fast_acl_rule--reference--group-001.md#canonical-0002230311211332-0003201210210103-1021122213201333-0300033323223111-1003200213003133-0112020311113311-2330111022311110-3113310211201100)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-0130033020112332-1320113013031122-3323102231201130-2302010233220010-0332222213110322-3330103313001312-0313202331310301-0031300023203130)
- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)

<a id="canonical-0002230311211332-0003201210210103-1021122213201333-0300033323223111-1003200213003133-0112020311113311-2330111022311110-3113310211201100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120233310302230-1020222313100201-3230312113010121-0101301031310111-0232203212220211-1322011222232112-1321101101030313-1112302300312133"></a>

## ip_prefix_set.ref — ref / 123212123302 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-0130033020112332-1320113013031122-3323102231201130-2302010233220010-0332222213110322-3330103313001312-0313202331310301-0031300023203130)
- [ip_prefix_set](data-sources--fast_acl_rule--reference--group-001.md#canonical-1101103210103331-2133222331132211-1320311020203200-0013220233202103-3311203330111331-3102002130201021-0313302320013023-2233201133112203)
- ip_prefix_set.ref

<a id="canonical-2011013333111123-2101123311320203-3211021313101013-2232331133220332-2100333010020133-1333313133130211-0023300101012331-0112113002132321"></a>

Type: `"list"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

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

<a id="canonical-1203132310132312-1201232220312013-2232303300301320-1233023111320011-2102112333000103-2312010232113130-0312001321132130-3211012232103211"></a>

## Direct properties — ref / 123212123302 / 3

<a id="canonical-2110332222203230-1221012233131301-2120021002220211-1100200133003200-2322303233231232-2030002232123303-3103100212000121-1110203023313312"></a>

<a id="canonical-1033300202011322-1111232203313012-0113112003030100-1210120120130021-1110010101200312-3210231323301210-1000102202012000-1311132013133033"></a>

## kind property — ref / 123212123302 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-2203003011103311-0130023033330312-1102232233101203-0100012301112223-0031110202010021-2303333023323133-1122112311010133-0202033310000003"></a>

## name property — ref / 123212123302 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-2222333322210322-1001332012302230-3010123001301031-1233110300123323-0033111130233002-2033331002201003-2011221323121323-1121211222220132"></a>

## namespace property — ref / 123212123302 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-2300200122030333-1020120202103123-0230122120202201-3202002223210012-0233110311123102-0003013200230002-2310230013222003-0213221211103333"></a>

## tenant property — ref / 123212123302 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-3302031321111331-3010032300211201-0011110203330232-0322122000330302-2112020202202123-2200320120321001-3321013013020101-3031130323122221"></a>

## uid property — ref / 123212123302 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-0033321032312113-2233303322221023-2203131320021221-2000211211021220-2001201220122003-2332123323121121-2102233012210322-0021001131121222"></a>

## Next pages — ref / 123212123302 / 9

- [ip_prefix_set](data-sources--fast_acl_rule--reference--group-001.md#canonical-1101103210103331-2133222331132211-1320311020203200-0013220233202103-3311203330111331-3102002130201021-0313302320013023-2233201133112203)
- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)

<a id="canonical-1322120012333330-2001323120031212-1213301121230220-1010301132000011-0020301112212330-0130033133120321-3222021133213110-3013020312031213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323231232002021-1122200302232110-0303113113021212-1111220022220330-2322203220133101-1231310130320020-3123110223012223-1130131220112032"></a>

## port — port / 100030330301 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-0130033020112332-1320113013031122-3323102231201130-2302010233220010-0332222213110322-3330103313001312-0313202331310301-0031300023203130)
- port

<a id="canonical-3132200202102321-0133231233111120-3232011130032300-3320003113012310-1300312201012233-1030333200030321-2220123012203301-0032032231210221"></a>

Type: `"list"`. Computed.

Source Ports. L4 port numbers to match.

Upstream description:

L4 port numbers to match.

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

<a id="canonical-2223033232221233-1003232013102132-1110120101212110-3313313012111201-2223312112323032-3300003131131302-3303210232012001-2300210333201221"></a>

## Direct properties — port / 100030330301 / 3

- [all](data-sources--fast_acl_rule--reference--group-001.md#canonical-0301031232230303-3322031021121223-0312023300121031-2000212223302202-0211231110121113-1300133132210132-0320322330302331-3122333010310123): complete subsection reference.

- [DNS](data-sources--fast_acl_rule--reference--group-001.md#canonical-2223232332312230-0332113221132320-0012131123212223-1030331012201002-3031212032312222-2203330131020122-3331200210103323-3321012113103333): complete subsection reference.

<a id="canonical-0220000221302021-1232102303200301-1010322331300300-0031323113121300-3103032133001301-1002302311322323-3123312010200021-2022110102322130"></a>

<a id="canonical-1310121331313000-2221111323130210-2113323132203031-0010210230101312-1013213211323310-3033003313211010-1003023223203300-0112002223123131"></a>

## user_defined property — port / 100030330301 / 4

Type: `"number"`. Computed.

Exclusive with \[all DNS\] Matches the user defined port.

Upstream description:

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

<a id="canonical-0023131312303301-3302322003313000-3231320021322022-2303101330132213-2111221000331223-3231132330111231-2133230210221303-2120032302213303"></a>

## Next pages — port / 100030330301 / 5

- [port.all](data-sources--fast_acl_rule--reference--group-001.md#canonical-0301031232230303-3322031021121223-0312023300121031-2000212223302202-0211231110121113-1300133132210132-0320322330302331-3122333010310123)
- [port.dns](data-sources--fast_acl_rule--reference--group-001.md#canonical-2223232332312230-0332113221132320-0012131123212223-1030331012201002-3031212032312222-2203330131020122-3331200210103323-3321012113103333)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-0130033020112332-1320113013031122-3323102231201130-2302010233220010-0332222213110322-3330103313001312-0313202331310301-0031300023203130)
- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)

<a id="canonical-0301031232230303-3322031021121223-0312023300121031-2000212223302202-0211231110121113-1300133132210132-0320322330302331-3122333010310123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220022233220322-3212003101212131-0201203321321013-0332122221010313-3333321112033002-2101300030133130-3310001333110100-3211111103133100"></a>

## port.all — all / 211313302202 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-0130033020112332-1320113013031122-3323102231201130-2302010233220010-0332222213110322-3330103313001312-0313202331310301-0031300023203130)
- [port](data-sources--fast_acl_rule--reference--group-001.md#canonical-1322120012333330-2001323120031212-1213301121230220-1010301132000011-0020301112212330-0130033133120321-3222021133213110-3013020312031213)
- port.all

<a id="canonical-1210320211123233-1300220212211033-2230003313022033-2332011033213331-0010323230311020-3202020122332110-0222030332310122-0203321100012201"></a>

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

<a id="canonical-0301132023001302-1331311211120222-0131030022220133-1201222032112120-3133020311231322-3100333101321113-0112300202103100-3023223232030202"></a>

## Direct properties — all / 211313302202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2233231112022031-0231213020011322-3233013232323220-3300102300321221-1022303212232201-1200032201213333-3011302120130033-2113223201010330"></a>

## Next pages — all / 211313302202 / 4

- [port](data-sources--fast_acl_rule--reference--group-001.md#canonical-1322120012333330-2001323120031212-1213301121230220-1010301132000011-0020301112212330-0130033133120321-3222021133213110-3013020312031213)
- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)

<a id="canonical-2223232332312230-0332113221132320-0012131123212223-1030331012201002-3031212032312222-2203330131020122-3331200210103323-3321012113103333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122032111001310-3003120201322221-1001002023330202-1320133001132213-1230212032210223-2233312110010031-0200210212130313-2213231301210001"></a>

## port.DNS — DNS / 230313100100 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-0130033020112332-1320113013031122-3323102231201130-2302010233220010-0332222213110322-3330103313001312-0313202331310301-0031300023203130)
- [port](data-sources--fast_acl_rule--reference--group-001.md#canonical-1322120012333330-2001323120031212-1213301121230220-1010301132000011-0020301112212330-0130033133120321-3222021133213110-3013020312031213)
- port.DNS

<a id="canonical-1212302101332033-1000001020302210-2322021323222021-3001021210212120-1001233131311020-2132203303203212-3020211020121002-3002302003103323"></a>

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

<a id="canonical-1212203010030030-2022021013320030-2112310012011123-1130133301031321-1210300303332203-0320002310030101-1320013021323020-1131302110313103"></a>

## Direct properties — DNS / 230313100100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3033212013110322-2330300132023313-0030233030222020-3001130302321222-2231302013331203-1002132110001300-3110330100221012-3131221320112232"></a>

## Next pages — DNS / 230313100100 / 4

- [port](data-sources--fast_acl_rule--reference--group-001.md#canonical-1322120012333330-2001323120031212-1213301121230220-1010301132000011-0020301112212330-0130033133120321-3222021133213110-3013020312031213)
- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)

<a id="canonical-2002033300333000-1223112020311322-1210001022020221-0123333022303330-0030013300310300-3210021023233013-1021201300130211-2203221310300103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111322223131201-2000211022201300-3321313210031001-2202312332312033-3233130223033200-3223322330321201-3300211303112101-1033312001033311"></a>

## prefix — prefix / 311332212112 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-0130033020112332-1320113013031122-3323102231201130-2302010233220010-0332222213110322-3330103313001312-0313202331310301-0031300023203130)
- prefix

<a id="canonical-1030231312313323-0102110311201310-3013101312122033-2220310232222003-2211000331012021-0100012132032231-1023202203301000-1310121100210210"></a>

Type: `"single"`. Computed.

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

Upstream description:

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

<a id="canonical-1030032223201200-0200113313111112-2031102021221030-3233011100323233-3120233131203111-2111323233311113-3132032311201202-1132232011132102"></a>

## Direct properties — prefix / 311332212112 / 3

<a id="canonical-0031203103323113-1013133202301223-2100022302201201-2310133022202132-0123321320300323-2201001330013113-1102233231011211-2233110112330311"></a>

<a id="canonical-3113000123102311-0211212232000212-3202220130231202-0133013213200202-0202211120113311-1102220222111330-3312221302211121-0300200233110113"></a>

## prefix property — prefix / 311332212112 / 4

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

<a id="canonical-2233211020020111-0013102130232300-0223202310221302-2111120312310301-0323202131020021-3120222120130211-1213003030331211-1230210012001102"></a>

## Next pages — prefix / 311332212112 / 5

- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-0130033020112332-1320113013031122-3323102231201130-2302010233220010-0332222213110322-3330103313001312-0313202331310301-0031300023203130)
- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)
