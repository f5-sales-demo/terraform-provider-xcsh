---
page_title: "xcsh_service_policy reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_service_policy reference."
---

# xcsh_service_policy reference

<a id="canonical-1312131031200012-0323222232212311-0110213200221122-3131110311121211-1002113122132211-1032112220111030-3100011123222312-2232213111221012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.bot_action` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.bot_action

<a id="canonical-0133132310332033-1121101121102222-0003231011211121-1121212002010233-2100112013133303-1010033012022130-0202303032132322-1233303100300200"></a>

Type: `"single"`. Computed.

Modify Bot protection behavior for a matching request. The modification could be to entirely skip
Bot processing.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_type": "[\"bot_skip_processing\",\"none\"]"
}
```

<a id="canonical-3223103031101211-2232111012302221-1323121231330221-2321300313302011-1222300120222311-1121330311030011-2212002012210231-1111300303022200"></a>

### Direct properties for `rule_list.rules.spec.bot_action`

- [bot_skip_processing](data-sources--service_policy--reference--group-002.md#canonical-2313033311032322-3032020223203001-1303102012100101-3031022232230001-0111200110021010-1332233011200222-0330020012123323-0101021013313210): complete subsection reference.

- [none](data-sources--service_policy--reference--group-002.md#canonical-3101221123320200-0333003020031002-0010313033133301-2303033033102010-2013133111011223-2321233032033212-2213100121232301-1103211221131121): complete subsection reference.

<a id="canonical-2313033311032322-3032020223203001-1303102012100101-3031022232230001-0111200110021010-1332233011200222-0330020012123323-0101021013313210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.bot_action.bot_skip_processing` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.bot_action](data-sources--service_policy--reference--group-002.md#canonical-1312131031200012-0323222232212311-0110213200221122-3131110311121211-1002113122132211-1032112220111030-3100011123222312-2232213111221012)
- rule_list.rules.spec.bot_action.bot_skip_processing

<a id="canonical-2130033111300203-1221231230001123-3000103031321132-3101102230212103-0201220011111323-0213130321010030-0231030100131201-1000112123012113"></a>

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

<a id="canonical-3101221123320200-0333003020031002-0010313033133301-2303033033102010-2013133111011223-2321233032033212-2213100121232301-1103211221131121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.bot_action.none` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.bot_action](data-sources--service_policy--reference--group-002.md#canonical-1312131031200012-0323222232212311-0110213200221122-3131110311121211-1002113122132211-1032112220111030-3100011123222312-2232213111221012)
- rule_list.rules.spec.bot_action.none

<a id="canonical-2312322002032001-1100123303302213-1212233230320012-2303113313033131-0000233201010130-0133000200102313-0200022313231202-0323012002200221"></a>

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

<a id="canonical-0121131233020103-2231131133023313-0021123010201112-3322100301021233-0200303032300002-1100203321101230-3201303322100031-3221133200303303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.client_name_matcher` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.client_name_matcher

<a id="canonical-3230203233231333-0311300302111332-3232031330201101-3202221313210010-3002012001210022-1010332200312300-2111010311003001-0211022333113300"></a>

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

<a id="canonical-3100230211132130-2221202313230312-2132002303212221-1101302213210303-0300010030133220-0001221003202033-3021231123320012-1311201111203301"></a>

### Direct properties for `rule_list.rules.spec.client_name_matcher`

<a id="canonical-0110221212023020-2002330213032203-1122033020223320-0222133233310211-1302200233230200-3012131310101303-0323332133032303-0210002203010120"></a>

#### `rule_list.rules.spec.client_name_matcher.exact_values` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3333103100022013-0203233120032003-0111203000021333-3033032013120030-3201130033121101-3121210303011031-1030300222130213-0032330102223102"></a>

<a id="canonical-1101030012233020-1200223221320123-0010032313102202-3213220331322131-3100312331223132-1110102022201122-2221323112031330-0100013132130301"></a>

#### `rule_list.rules.spec.client_name_matcher.regex_values` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0100232010003320-1023310102010202-0133323030020233-2013031202310302-3201031130110312-1310332130200103-2212030330200122-0303001130201331"></a>

<a id="canonical-0310230022330323-1203301322033322-0103313303322231-3320330002312103-3122100011130331-2120012103021023-0001000001100130-1030103311213223"></a>

#### `rule_list.rules.spec.client_name_matcher.transformers` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1230321030333321-2133202001231022-0210120112103312-3130210331103021-0123133331212113-1310313101233211-2300330310202031-0032002312032330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.client_selector` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.client_selector

<a id="canonical-0021003033023030-1032122012333232-3012330332331020-3331110221001300-3301031012220111-0101233002221021-3122102211210330-3220021210122323"></a>

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

<a id="canonical-1212302221230023-3122330011233130-3211201112323232-3112332131111123-3120310100110330-0010003220333020-2210103213033020-2121103113212300"></a>

### Direct properties for `rule_list.rules.spec.client_selector`

<a id="canonical-1122131031331101-1132203301003312-1302221230232301-1310200030032231-2230122013120202-0313323033120221-0303123120111012-0131110103212103"></a>

#### `rule_list.rules.spec.client_selector.expressions` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2111300231202303-2030301101223223-3032333120002222-2000302311201303-3012020113230021-1223300021220201-3302002113123122-3213110013333130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.cookie_matchers` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.cookie_matchers

<a id="canonical-3022303210322011-1331331100001332-1300031010023032-0312201111112020-0230322303020113-0122212030021002-3222211120321133-1021000200001303"></a>

Type: `"list"`. Computed.

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2223000323321332-1311212303321320-0031201011120203-2233100301230312-3332223302210223-0213131321312100-2121311331232311-3233121131032331"></a>

### Direct properties for `rule_list.rules.spec.cookie_matchers`

- [check_not_present](data-sources--service_policy--reference--group-002.md#canonical-0313331203112103-2002332123131322-1011132200123113-0210201011321220-2033012202300203-1102233112332010-2010233011033201-1110323313100300): complete subsection reference.

- [check_present](data-sources--service_policy--reference--group-002.md#canonical-3323212031121321-0010000010113121-2021333333300023-1002310112312230-3102213020322331-3331001300020232-2112121303121223-0213001011001111): complete subsection reference.

<a id="canonical-1202323210300133-0102023012133310-0031212001003222-3233201003001222-2032232001302230-2321102213120033-0021021123101301-3031020301200113"></a>

<a id="canonical-2133330001020333-1102121113110100-2010320230002303-1201013320023201-1102330222103132-0001320203130223-0223001210300322-0003331212100231"></a>

#### `rule_list.rules.spec.cookie_matchers.invert_matcher` property

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

- [item](data-sources--service_policy--reference--group-002.md#canonical-0113022300130222-3321103312230102-1331111312213321-3131223210222023-1000202002123023-3201130120020323-3220320331221131-2030211330312102): complete subsection reference.

<a id="canonical-2032322003122032-2331110210312111-0221031112221332-2010101031300110-2110122330130200-2001102313133300-1000211121210001-2303122201320102"></a>

<a id="canonical-0022032202231201-2331212320020003-3312303113223002-0231221020101331-2230100210003110-0201133201111021-0131011220322021-3220133323113301"></a>

#### `rule_list.rules.spec.cookie_matchers.name` property

Type: `"string"`. Computed.

Cookie Name. A case-sensitive cookie name.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-0313331203112103-2002332123131322-1011132200123113-0210201011321220-2033012202300203-1102233112332010-2010233011033201-1110323313100300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.cookie_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.cookie_matchers](data-sources--service_policy--reference--group-002.md#canonical-2111300231202303-2030301101223223-3032333120002222-2000302311201303-3012020113230021-1223300021220201-3302002113123122-3213110013333130)
- rule_list.rules.spec.cookie_matchers.check_not_present

<a id="canonical-2001300121223031-0020010012022030-3130113233001003-1230221303013313-0333133000323012-3332133023212011-2111011103222023-0313001212330233"></a>

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

<a id="canonical-3323212031121321-0010000010113121-2021333333300023-1002310112312230-3102213020322331-3331001300020232-2112121303121223-0213001011001111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.cookie_matchers.check_present` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.cookie_matchers](data-sources--service_policy--reference--group-002.md#canonical-2111300231202303-2030301101223223-3032333120002222-2000302311201303-3012020113230021-1223300021220201-3302002113123122-3213110013333130)
- rule_list.rules.spec.cookie_matchers.check_present

<a id="canonical-2213201131210103-0002320003313033-1022202333132200-0322112311013111-0022021211210301-3232213123101102-3010001301001012-3011130310220011"></a>

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

<a id="canonical-0113022300130222-3321103312230102-1331111312213321-3131223210222023-1000202002123023-3201130120020323-3220320331221131-2030211330312102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.cookie_matchers.item` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.cookie_matchers](data-sources--service_policy--reference--group-002.md#canonical-2111300231202303-2030301101223223-3032333120002222-2000302311201303-3012020113230021-1223300021220201-3302002113123122-3213110013333130)
- rule_list.rules.spec.cookie_matchers.item

<a id="canonical-0210020002333210-2312000302103031-3121223003210113-2112323112000000-2022110130132001-2332133033313001-2300021121312033-3213220313031033"></a>

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

<a id="canonical-1312102110201011-0131132332013030-2212223001332001-2000120223023113-1110333311132113-2231013012002130-2333201321100133-3113002202122002"></a>

### Direct properties for `rule_list.rules.spec.cookie_matchers.item`

<a id="canonical-1303200000020232-3031323201322203-1013212113130000-1031313303232131-2323333311023002-0333011310322133-3110300332330213-1312132022030011"></a>

#### `rule_list.rules.spec.cookie_matchers.item.exact_values` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1032032103031000-0030202201111313-3222200213121032-3323103030023212-1123213110311033-0120232012131213-3120031001103201-2033333123033022"></a>

<a id="canonical-0222122120331000-3030131011223202-1230122332000302-3232001201320321-0322130102323101-1203133302212222-1300003023330302-2033000202010010"></a>

#### `rule_list.rules.spec.cookie_matchers.item.regex_values` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1313200201122022-3031222210230132-1031302211212112-1013320120222001-2101231332103003-0300231011132200-1213030200003311-0303232223221132"></a>

<a id="canonical-2200302011311000-2332123322133103-3011231201332120-3220123201233322-2012203221002300-1130320300330323-1232130201030233-3331030033303001"></a>

#### `rule_list.rules.spec.cookie_matchers.item.transformers` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3322203033331023-2233220302031220-2301231103021132-2020113132210223-2001020313023320-0220311122312313-0121320222003303-2131021003133213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.domain_matcher` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.domain_matcher

<a id="canonical-3121212333330100-1103320130023311-3111311030320303-1001113330321331-2221021021212000-2300312331100012-1212202110302300-0023223223031331"></a>

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

<a id="canonical-1220032301313300-0300231301332022-3333313321323222-0330220301211223-2010030223113012-2023020022132112-3032010023332232-3011032221331213"></a>

### Direct properties for `rule_list.rules.spec.domain_matcher`

<a id="canonical-0023101030131212-3002101010110131-0132020223110202-3210132033203131-2001032203110222-3203311031222231-1202312022113010-2100121331131331"></a>

#### `rule_list.rules.spec.domain_matcher.exact_values` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1031132202230330-3321001313001231-1210311301130200-1303022132032102-1322222300103330-2300113210232122-0330202311231131-3033200212020030"></a>

<a id="canonical-0121321033223223-3222111223031231-3102230200130303-1003031101013220-3110320101212031-1103121211221030-0132213121012023-1022233021020033"></a>

#### `rule_list.rules.spec.domain_matcher.regex_values` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1333211021320022-2003033321121121-1301012022313223-1210023023001022-3301120312130221-2021122032200103-3213102031103112-3202000000022013"></a>

<a id="canonical-3230200323033320-1130322120223221-1310323002002112-3200222313013202-1100220030003000-3231133103131100-1300111331211110-2111321312232120"></a>

#### `rule_list.rules.spec.domain_matcher.transformers` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2022122312102320-3021230232003303-1301003301023011-2212121031302321-3123311303200100-0020331201223301-3103232233311212-3033201132111122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.headers` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.headers

<a id="canonical-1001323031223130-3200130200303132-0011101103022322-3221232223200332-2231010031002222-2303021010213331-0222313303011331-3002333222231021"></a>

Type: `"list"`. Computed.

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-1300033100303231-1332033233133122-0001233303120001-0233003130201032-2003221122231321-0112332030313031-3110203021223102-3123022321203013"></a>

### Direct properties for `rule_list.rules.spec.headers`

- [check_not_present](data-sources--service_policy--reference--group-002.md#canonical-1300330332022120-3002303012201222-1303023230221212-3001331122011220-2321110133312210-3222320030023233-2003222000302000-3220001323330221): complete subsection reference.

- [check_present](data-sources--service_policy--reference--group-002.md#canonical-2331020202230311-3120200212012302-2223120000130132-2013233310311112-2033001013131231-2010100303211213-3230121102321333-3003323330213311): complete subsection reference.

<a id="canonical-2200020300203210-2102123211333203-1000003020231003-0032032011302302-0020100330011102-0103132101213121-0033032033222102-0323132120000132"></a>

<a id="canonical-0203300321032213-1112032100202000-0030222122030321-2212012213320302-3010231130223031-0110303303300011-1131302103203332-3220211013111300"></a>

#### `rule_list.rules.spec.headers.invert_matcher` property

Type: `"bool"`. Computed.

Invert Header Matcher. Invert the match result.

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

- [item](data-sources--service_policy--reference--group-002.md#canonical-0231301332320130-2211022322210311-1130333032312220-3120301131220333-1300322223322203-2321013101223230-1102303113003202-1011231001220212): complete subsection reference.

<a id="canonical-0031122131232323-0311201220123003-3112232130032000-1001113030223200-0222110102200130-3230320200223321-1012032030013303-0020033020131231"></a>

<a id="canonical-1333320221010010-2003133231020110-1112200221011002-3031322130020331-2131123113032233-2211203331233211-0030203103012122-1030023212221300"></a>

#### `rule_list.rules.spec.headers.name` property

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-1300330332022120-3002303012201222-1303023230221212-3001331122011220-2321110133312210-3222320030023233-2003222000302000-3220001323330221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.headers.check_not_present` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.headers](data-sources--service_policy--reference--group-002.md#canonical-2022122312102320-3021230232003303-1301003301023011-2212121031302321-3123311303200100-0020331201223301-3103232233311212-3033201132111122)
- rule_list.rules.spec.headers.check_not_present

<a id="canonical-0113220031331322-3121130212333302-2002012121001223-1233211111021323-3012133323312321-3323111301030203-1233332121120031-1301001123313102"></a>

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

<a id="canonical-2331020202230311-3120200212012302-2223120000130132-2013233310311112-2033001013131231-2010100303211213-3230121102321333-3003323330213311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.headers.check_present` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.headers](data-sources--service_policy--reference--group-002.md#canonical-2022122312102320-3021230232003303-1301003301023011-2212121031302321-3123311303200100-0020331201223301-3103232233311212-3033201132111122)
- rule_list.rules.spec.headers.check_present

<a id="canonical-1323023110013302-3123031303031301-0231221112121332-2000120330213130-3021011110313020-0320012032132231-0010120020330220-1331113111201131"></a>

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

<a id="canonical-0231301332320130-2211022322210311-1130333032312220-3120301131220333-1300322223322203-2321013101223230-1102303113003202-1011231001220212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.headers.item` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.headers](data-sources--service_policy--reference--group-002.md#canonical-2022122312102320-3021230232003303-1301003301023011-2212121031302321-3123311303200100-0020331201223301-3103232233311212-3033201132111122)
- rule_list.rules.spec.headers.item

<a id="canonical-0131300331331012-3203011121023013-3230331023200313-0323032320020111-1023002001332211-2210321021023010-3212032223222121-2010132201032032"></a>

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

<a id="canonical-1333010300102103-1313001110302321-0233020210033231-0113230213111300-2121313320221012-1213221312223301-2303001013301200-0033030122232233"></a>

### Direct properties for `rule_list.rules.spec.headers.item`

<a id="canonical-1201123121330233-0320331203133120-0221032002020023-1131330133213111-3101311031300320-1321112001122010-3112021012223303-0100000010302330"></a>

#### `rule_list.rules.spec.headers.item.exact_values` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3212322023122313-3203321001100221-1103013223133210-0312123311031011-2011121001220312-1110330302022100-3230003122332311-3330011112011120"></a>

<a id="canonical-1021230101323332-3330112203300132-1233230023003311-1230003233310323-2311012231333010-3203222012101303-3201321011002201-3312221301103101"></a>

#### `rule_list.rules.spec.headers.item.regex_values` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3222213231201121-0022131033222302-2110320132000233-3222020123002023-1313022332112212-0323230221020202-2012001211001001-2012221130200021"></a>

<a id="canonical-0100331230130201-3220230332321323-0100001001000220-1323010112221200-1322311212110200-2201203302201210-0030230201013013-0321322012223301"></a>

#### `rule_list.rules.spec.headers.item.transformers` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0223130311323131-0303100033332230-3203111320203213-3102331312020011-3013112003331321-2122323011003121-2313130233330233-0103130010303202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.http_method` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.http_method

<a id="canonical-2313321331200130-0221313021311130-2202301020320121-3033333012111120-3100213131030221-3011122303123000-1123002030313001-3011130013233322"></a>

Type: `"single"`. Computed.

An HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
considered successful if the input method is a member of the list. The result of the match based on
the method list is inverted if invert\_matcher is true.

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

<a id="canonical-3001300203121032-2313312323201033-0133333030303203-2113310000003110-3310003222110301-0011330030103302-0122311300233032-3021302332111303"></a>

### Direct properties for `rule_list.rules.spec.http_method`

<a id="canonical-3002231303110022-0312113102132001-1321121302002030-2132132300112132-2013301103223020-1020300333332330-2310000321321100-3003211030320101"></a>

#### `rule_list.rules.spec.http_method.invert_matcher` property

Type: `"bool"`. Computed.

Invert Method Matcher. Invert the match result.

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

<a id="canonical-1112232112123232-1033332002021200-0313002233110322-1123010013230230-3203010331222302-1020001133220012-3013102202320232-1012312320233133"></a>

<a id="canonical-3102112231212010-0210000200111033-0101010111020222-1010030210012130-1233131303120013-3321012131032222-2101232302220300-0013222333020001"></a>

#### `rule_list.rules.spec.http_method.methods` property

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] List of methods values to
match against. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`,
\`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3001311313101200-0111313102213231-0330101313233022-0210001223230123-2232331303200003-2213132112112331-2233022233020222-3112013102100001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.ip_matcher` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.ip_matcher

<a id="canonical-1331103032030102-2202002232213222-1121220301103201-1312101033120123-2222122002112213-3011010203232302-2120300023201203-1221230211311001"></a>

Type: `"single"`. Computed.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

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

<a id="canonical-0020302111110011-3232013132220332-0302213012002033-1003331011332200-0311232200001102-0020010100312131-2020211133023330-1120332203133201"></a>

### Direct properties for `rule_list.rules.spec.ip_matcher`

<a id="canonical-3310210102200031-3113013002201213-2030322331233002-2033200311013101-0000103132112123-0002301231302331-3331202331311030-0301223131302313"></a>

#### `rule_list.rules.spec.ip_matcher.invert_matcher` property

Type: `"bool"`. Computed.

Invert IP Matcher. Invert the match result.

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

- [prefix_sets](data-sources--service_policy--reference--group-002.md#canonical-3202231202123333-2110110232001130-1131033132311212-2020131100002112-3310111023003233-0203301303212202-3100123112303020-3333230020323313): complete subsection reference.

<a id="canonical-3202231202123333-2110110232001130-1131033132311212-2020131100002112-3310111023003233-0203301303212202-3100123112303020-3333230020323313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.ip_matcher.prefix_sets` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.ip_matcher](data-sources--service_policy--reference--group-002.md#canonical-3001311313101200-0111313102213231-0330101313233022-0210001223230123-2232331303200003-2213132112112331-2233022233020222-3112013102100001)
- rule_list.rules.spec.ip_matcher.prefix_sets

<a id="canonical-0330011200233103-1103220323320011-1311331221001120-0211020322101330-3032112322011203-2003120203331101-3020030130023113-3230221113002313"></a>

Type: `"list"`. Computed.

A list of references to ip\_prefix\_set objects.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2011101002111230-2113122332003321-2321233222103110-3013020312012102-1133020003100312-0120121121321013-1320123022133202-2302131122133322"></a>

### Direct properties for `rule_list.rules.spec.ip_matcher.prefix_sets`

<a id="canonical-2322013233111133-1011230113131320-2111220110130101-1002020122033333-3122113211323310-2203021121123311-0233021020313210-3313103223230132"></a>

#### `rule_list.rules.spec.ip_matcher.prefix_sets.kind` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0330301031313300-3211323220200013-0331221123012112-3001311010313330-0001010300100130-1322232313103100-3331321000223200-1023221322132102"></a>

<a id="canonical-1031102000210321-2333313320102031-2033313123322213-0210022320233033-2210230321001030-2313310221331002-2101001233312302-0011302233230321"></a>

#### `rule_list.rules.spec.ip_matcher.prefix_sets.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3001031222221102-2221000130030010-3320022311220221-2212010013103330-1103103331101111-1322020012303333-1230210213020203-2131210213330011"></a>

<a id="canonical-3330103220210003-1310233122333201-3220301230131310-0031103102323331-3033110121210311-2021100000010332-0310221103003301-0032002223322103"></a>

#### `rule_list.rules.spec.ip_matcher.prefix_sets.namespace` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2121201132222230-1333213120120023-1233233012012311-3013220323021210-1310301301333321-2312012331003210-2123022211332323-0123210312032101"></a>

<a id="canonical-2130312102121120-3031200111003203-2303303100102312-1001130100312113-0132301300310031-0013101122122003-0303321330031302-1023300133113200"></a>

#### `rule_list.rules.spec.ip_matcher.prefix_sets.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0122311311233032-1022010012320301-2033020203222131-3122321101330203-3301300201003220-0210103122311300-2331133000333332-1233220301211333"></a>

<a id="canonical-3133331203100330-2313322120103200-2312130321221201-2010010113301012-0222321231030303-2113233200323121-3113123100130121-0002322032223312"></a>

#### `rule_list.rules.spec.ip_matcher.prefix_sets.uid` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1133021322231212-2103223100112132-2212113321333203-0113233322020010-0003213320223311-3112100133213021-3323011012332312-1123001232301321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.ip_prefix_list` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.ip_prefix_list

<a id="canonical-0332320110133201-3103032311222213-2312021222023232-0330203300133120-1232221103000032-0021300000103213-2220330322332103-1101231122101103"></a>

Type: `"single"`. Computed.

List of IP Prefix strings to match against.

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

<a id="canonical-3231101303322131-0130121003333210-3132310001212303-3003210223020012-0012211322010301-1122201133020222-1030223031101001-3031230102020312"></a>

### Direct properties for `rule_list.rules.spec.ip_prefix_list`

<a id="canonical-3020222320221332-3211131100011233-0233021301301223-3120032112132122-2322120322302233-2032302202303103-2332021132210330-0103123213331131"></a>

#### `rule_list.rules.spec.ip_prefix_list.invert_match` property

Type: `"bool"`. Computed.

Invert Match Result. Invert the match result.

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

<a id="canonical-0333212333332103-3030131233230322-3332112111013110-0013032331201233-2132232001000230-2332021221012113-0313130323113332-0122133333332230"></a>

<a id="canonical-0233031313320122-0222333301213230-1311000222130002-3002133331302112-0102131020310330-0003013022020112-0312001322022002-2010313303303221"></a>

#### `rule_list.rules.spec.ip_prefix_list.ip_prefixes` property

Type: `["list", "string"]`. Computed.

IPv4 Prefix List. List of IPv4 prefix strings.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0213001221332012-0333230020030103-1200211310131211-0231323331103203-2310310332330330-3313312022312012-1302033120012331-0021212310222310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.ip_threat_category_list` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.ip_threat_category_list

<a id="canonical-0332320111030232-3310113121200022-1031330203210102-2212203332210312-3201012132323000-0200201001103023-0210310132323111-3010331010001100"></a>

Type: `"single"`. Computed.

IP Threat Category List Type. List of IP threat categories.

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

<a id="canonical-1022310001321202-3321010120203002-1111102311122203-2033013213200301-3121230221022300-3100120312103213-1202312133011301-3332003130322300"></a>

### Direct properties for `rule_list.rules.spec.ip_threat_category_list`

<a id="canonical-2102010000011313-1231113311132130-1000300113122132-2132213113111123-1310321100110133-1220232223110023-1310101201212000-0201222211103130"></a>

#### `rule_list.rules.spec.ip_threat_category_list.ip_threat_categories` property

Type: `["list", "string"]`. Computed.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`,
\`WEB\_ATTACKS\`, \`BOTNETS\`, \`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`,
\`MOBILE\_THREATS\`, \`TOR\_PROXY\`, \`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to
\`SPAM\_SOURCES\`.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0020023013210332-1201130101003103-1003112321233212-2120210311320132-0321330033001332-2103123112111112-1013321012220033-1032311101122303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.ja4_tls_fingerprint` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.ja4_tls_fingerprint

<a id="canonical-2133321213201332-0103033131311023-0133012001231300-2210133121211321-1011003323303322-2313311011311000-3310323131231202-0121121103231010"></a>

Type: `"single"`. Computed.

An extended version of JA3 that includes additional fields for more comprehensive fingerprinting of
SSL/TLS clients and potentially has a different structure and length.

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

<a id="canonical-3013003022302013-2000013331221331-3333120332221301-2000010132010110-0100323130011102-1312022110211322-3011313121123133-3233011002211311"></a>

### Direct properties for `rule_list.rules.spec.ja4_tls_fingerprint`

<a id="canonical-0101100031300100-1122131102312303-0302101013011302-3133300312100221-0330321102320221-0321120232331131-2220323311331113-0113313030320320"></a>

#### `rule_list.rules.spec.ja4_tls_fingerprint.exact_values` property

Type: `["list", "string"]`. Computed.

A list of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.items.string.len": "36",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "36",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0302332200230030-0110111001203102-3332033203213321-1013111323001230-3121103113222113-0233230330232331-3022210003302100-3123110320330031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.jwt_claims` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.jwt_claims

<a id="canonical-1333311322310330-3102031002331332-3211032023103233-1223113012113130-3112020132003021-0221131203221330-1123130331202202-3210211021330122"></a>

Type: `"list"`. Computed.

A list of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates
must evaluate to true.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3030311212322200-2321303003133033-1321300013330003-0103300312131111-3133010002011310-0320103202110100-0010102203021212-1231033032300200"></a>

### Direct properties for `rule_list.rules.spec.jwt_claims`

- [check_not_present](data-sources--service_policy--reference--group-002.md#canonical-0000311320320001-0012330333232113-3022302231303110-1121332312320213-3031231121021330-0301231011101111-0023133330211233-0113113021230220): complete subsection reference.

- [check_present](data-sources--service_policy--reference--group-002.md#canonical-2133322320021233-2122213300200131-0302313202131032-2332132221322000-1200300032200122-2023003120002302-0100100320101012-2333212001112110): complete subsection reference.

<a id="canonical-1110220012201301-1232232002032110-2011231102301332-0332002321321300-1313230010031102-3212031322203320-1121222302320213-1031033001030322"></a>

<a id="canonical-0130130002330310-1203113200232311-1032222001001200-1223313310102121-2301032031011333-1020212132103210-2013300000301322-1233012121120221"></a>

#### `rule_list.rules.spec.jwt_claims.invert_matcher` property

Type: `"bool"`. Computed.

Invert Matcher. Invert the match result.

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

- [item](data-sources--service_policy--reference--group-002.md#canonical-2133001033003311-3333310103200033-1133032032120001-3330003032003221-0020213021332101-2022303132001200-3013022132221303-0312310002033302): complete subsection reference.

<a id="canonical-2133232001211003-0333000330222003-2203013111010320-2233123120202301-3032211102121320-3132232001102102-1102313112110000-2223003010031333"></a>

<a id="canonical-3223010211233103-2310122301033033-2313112020200102-3133133020213113-3202312223231121-1303310011233212-1011303102011103-0320213133113111"></a>

#### `rule_list.rules.spec.jwt_claims.name` property

Type: `"string"`. Computed.

JWT Claim Name. JWT claim name.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-0000311320320001-0012330333232113-3022302231303110-1121332312320213-3031231121021330-0301231011101111-0023133330211233-0113113021230220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.jwt_claims.check_not_present` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.jwt_claims](data-sources--service_policy--reference--group-002.md#canonical-0302332200230030-0110111001203102-3332033203213321-1013111323001230-3121103113222113-0233230330232331-3022210003302100-3123110320330031)
- rule_list.rules.spec.jwt_claims.check_not_present

<a id="canonical-1132010000203301-0012121211121100-2112003032022101-3022332231302313-2130000000002101-2231001202321032-0313330101211332-0302231121023303"></a>

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

<a id="canonical-2133322320021233-2122213300200131-0302313202131032-2332132221322000-1200300032200122-2023003120002302-0100100320101012-2333212001112110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.jwt_claims.check_present` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.jwt_claims](data-sources--service_policy--reference--group-002.md#canonical-0302332200230030-0110111001203102-3332033203213321-1013111323001230-3121103113222113-0233230330232331-3022210003302100-3123110320330031)
- rule_list.rules.spec.jwt_claims.check_present

<a id="canonical-1021011111003020-0111321101203321-1322100003311000-2230201220001321-1220231301200001-3321332130230331-3032222232331321-2223302012223311"></a>

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

<a id="canonical-2133001033003311-3333310103200033-1133032032120001-3330003032003221-0020213021332101-2022303132001200-3013022132221303-0312310002033302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.jwt_claims.item` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.jwt_claims](data-sources--service_policy--reference--group-002.md#canonical-0302332200230030-0110111001203102-3332033203213321-1013111323001230-3121103113222113-0233230330232331-3022210003302100-3123110320330031)
- rule_list.rules.spec.jwt_claims.item

<a id="canonical-3303300231310012-2210221200011322-3033332203300223-1033233132303122-1132123103000012-1300123112233213-0301111020322310-2223031312010032"></a>

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

<a id="canonical-3133031113331020-3022131123203123-1312111210010003-1111132311313301-0323021030231203-2112030010013212-1223311102001030-3100201210221033"></a>

### Direct properties for `rule_list.rules.spec.jwt_claims.item`

<a id="canonical-2131000112113003-3102230203212011-2313111102011211-0320331130011022-1011023030313311-2112120022300022-1102233112231130-2321301031220020"></a>

#### `rule_list.rules.spec.jwt_claims.item.exact_values` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3331310001031232-0232122121213310-0321013012102312-1122023212311313-0102332100121203-0133130200333312-3332212022020330-0022002031023010"></a>

<a id="canonical-3032222133112000-0103230211203111-2101202231021320-1213230033101310-0032320231133223-3003023310020103-0000233313212312-2011313110332213"></a>

#### `rule_list.rules.spec.jwt_claims.item.regex_values` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0313111331103320-2121332301210233-0223223210320232-0200220232033030-2231330133232122-3103033212010201-2021313300111303-2113203321030200"></a>

<a id="canonical-1301221101023333-1113222102111113-0312132020120210-2120331030101233-2013111021112320-3310033232220012-1023000133033300-3022132132331122"></a>

#### `rule_list.rules.spec.jwt_claims.item.transformers` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1221302121333112-2003223202021213-1330010203000001-3101131322012130-0100030033121220-2120330230202021-3310121000131102-1311103021210331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.label_matcher` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.label_matcher

<a id="canonical-0320223011033230-1210020210111321-2201022210233031-3323031132200230-0303021300212020-2023330312000311-0033033201332310-3310210302123032"></a>

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

<a id="canonical-0302123113311323-2332010011231333-2300201312121023-0101133102023321-2303321233102012-3031013311031313-1123013121012103-2122020020321202"></a>

### Direct properties for `rule_list.rules.spec.label_matcher`

<a id="canonical-1030102212313222-3121203023100211-2231011320110222-2320320110030322-0020232223211230-1231312100213111-1111222000300322-2111331013331012"></a>

#### `rule_list.rules.spec.label_matcher.keys` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2122103222123111-0310123221120033-3122230001202021-3231333033332103-3101331023112023-0203202032222223-1123223012003012-1103223132011322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.mum_action` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.mum_action

<a id="canonical-3011112132131102-2033202212222122-3103323112320223-3231220213112223-2002132103222211-2031301130001220-3031203330010021-1003020332333233"></a>

Type: `"single"`. Computed.

Modify behavior for a matching request. The modification could be to entirely skip processing.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_type": "[\"default\",\"skip_processing\"]"
}
```

<a id="canonical-0121330301022103-2212212001113003-2121121302132130-3213201022123323-2332002010013211-3222331012021132-2020211103201202-1333232122022032"></a>

### Direct properties for `rule_list.rules.spec.mum_action`

- [default](data-sources--service_policy--reference--group-002.md#canonical-2333101121212311-0111311200101033-0230111002102322-0230322310102300-1102212010300102-1012120023031322-0220232301011101-3300030301311203): complete subsection reference.

- [skip_processing](data-sources--service_policy--reference--group-002.md#canonical-3130133131120300-3021223233020031-1033121030010323-1211321123133201-2311230310022211-2101332000012031-2021100030032222-1130111221322312): complete subsection reference.

<a id="canonical-2333101121212311-0111311200101033-0230111002102322-0230322310102300-1102212010300102-1012120023031322-0220232301011101-3300030301311203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.mum_action.default` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.mum_action](data-sources--service_policy--reference--group-002.md#canonical-2122103222123111-0310123221120033-3122230001202021-3231333033332103-3101331023112023-0203202032222223-1123223012003012-1103223132011322)
- rule_list.rules.spec.mum_action.default

<a id="canonical-1120210102212202-0010103020100022-2221301130030031-0111213201001111-0200203303032211-3203122111120132-2210130310110222-0300313032333022"></a>

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

<a id="canonical-3130133131120300-3021223233020031-1033121030010323-1211321123133201-2311230310022211-2101332000012031-2021100030032222-1130111221322312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.mum_action.skip_processing` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.mum_action](data-sources--service_policy--reference--group-002.md#canonical-2122103222123111-0310123221120033-3122230001202021-3231333033332103-3101331023112023-0203202032222223-1123223012003012-1103223132011322)
- rule_list.rules.spec.mum_action.skip_processing

<a id="canonical-0330211221313200-0320010013212020-3300331232303321-2310123123312332-2231303100302130-2231120122313032-1232102102120200-1022213221322320"></a>

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

<a id="canonical-0320210310200110-3200233233321020-3202312003012032-3221022312030021-3213113211200131-0321203201123201-1301333233103210-0313130122100133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.path` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.path

<a id="canonical-3112202323333330-1002011132230113-0301131232010302-3223310322220113-2211000203231201-3331203031030111-2211031100031002-2232301030003023"></a>

Type: `"single"`. Computed.

A path matcher specifies multiple criteria for matching an HTTP path string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of path prefixes, a list of exact path values and a list of regular expressions.

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

<a id="canonical-1020321102121230-3310120211203230-3212010200001030-2232331303122333-0021030022002101-2000032112022201-1201230230012032-1310100331032230"></a>

### Direct properties for `rule_list.rules.spec.path`

<a id="canonical-1330100122300331-1010313230320222-2032001201021311-1210133203332112-0000133202122221-2023110121133022-0230312023000201-0210233201223311"></a>

#### `rule_list.rules.spec.path.encoded_path_matcher` property

Type: `"bool"`. Computed.

Match against the encoded, escaped path.

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

<a id="canonical-0120312020113032-1030120030023321-3210021211312110-1120320012301200-2203130312012220-3011232321202021-0110101100132130-2003310023000023"></a>

<a id="canonical-0222321221213102-3210221111013032-1021333101013003-2011233023200123-3233120021213112-2303300313122001-3021213220231210-1310323320133222"></a>

#### `rule_list.rules.spec.path.exact_values` property

Type: `["list", "string"]`. Computed.

A list of exact path values to match the input HTTP path against.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0001322103131130-1110233012103113-1113020300221322-2001213320020330-0201322323121321-0010112211101020-0020022103310202-0233020120132330"></a>

<a id="canonical-0323003110102121-1313003332010313-1032012213220232-3201012030030022-0032023200232020-1132031122213213-1102113100222330-1112102211001200"></a>

#### `rule_list.rules.spec.path.invert_matcher` property

Type: `"bool"`. Computed.

Invert Path Matcher. Invert the match result.

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

<a id="canonical-1010312003323223-3201132023023210-3212301221100123-1230300002332211-2301312202110203-1223312121230331-3230222233300220-2301311203112011"></a>

<a id="canonical-2001333130103330-0320021110011113-3323122121330213-0023303020200311-2322022003003111-3122031030100122-0011222103113223-0302220321333002"></a>

#### `rule_list.rules.spec.path.prefix_values` property

Type: `["list", "string"]`. Computed.

A list of path prefix values to match the input HTTP path against.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1320103323133211-0212121220121210-1212100113200310-3113320110302231-0112032211022330-3033331003200220-0120113131332010-2331101223110123"></a>

<a id="canonical-1333213231211012-2121133233120222-2333221300201101-0001232231313221-3210002010222010-0323012203012200-3023101130323302-3321013230133331"></a>

#### `rule_list.rules.spec.path.regex_values` property

Type: `["list", "string"]`. Computed.

A list of regular expressions to match the input HTTP path against.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1011132201120030-2303313012003131-0203001003333022-3103122030222022-3120002133012021-0021000302033131-0302020230322221-0302113102010323"></a>

<a id="canonical-2331100022223230-1110213133021013-1323333211222321-3033131201223120-2132232202232113-0331000023313120-2010000132232023-1332031320010022"></a>

#### `rule_list.rules.spec.path.suffix_values` property

Type: `["list", "string"]`. Computed.

A list of path suffix values to match the input HTTP path against.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1223332033210322-1002000230132301-1321120211330120-1120110302101211-2001020312003303-0002122010121332-0232222213030002-2012103003111103"></a>

<a id="canonical-0223330002110203-3212212111020302-3103001222123310-2221031001111110-1132022313101231-1231113022213010-1132232230000210-2031221223220120"></a>

#### `rule_list.rules.spec.path.transformers` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0002112321213331-2301331003031102-2101130012331010-0010130231301212-0111003003223033-3131022232020221-3200101100331330-2033223020101132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.port_matcher` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.port_matcher

<a id="canonical-0102020310300131-3033033100311323-3102032333033332-0203020321021112-0113020213001120-3112302213320120-3120120130023133-0323220002322302"></a>

Type: `"single"`. Computed.

Port matcher specifies a list of port ranges as match criteria. The match is considered successful
if the input port falls within any of the port ranges. The result of the match is inverted if
invert\_matcher is true. Server applies default when omitted.

Additional upstream details:

A port matcher specifies a list of port ranges as match criteria.

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

<a id="canonical-3032012120023132-1303112202102120-0231003022213230-1000013333220030-2332230011023002-0022223223002121-1131021220001223-2111021021110212"></a>

### Direct properties for `rule_list.rules.spec.port_matcher`

<a id="canonical-1122010102022223-1313322021133300-3201123231331023-0222100120310123-3233012330020101-3123212231013000-1032300211133103-2202030233330312"></a>

#### `rule_list.rules.spec.port_matcher.invert_matcher` property

Type: `"bool"`. Computed.

Invert Port Matcher. Invert the match result.

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

<a id="canonical-0132232331202033-0111310020202112-1310331330133210-0223232110220211-2032012001123101-1210332232122010-3212113213223233-2022211232120231"></a>

<a id="canonical-0123330001212310-3003312002213131-0011132003013221-1222303111022220-2300222010231023-2002203202322313-0203022110332013-3210012001021022"></a>

#### `rule_list.rules.spec.port_matcher.ports` property

Type: `["list", "string"]`. Computed.

List of strings, each of which is a single port value or a tuple of start and end port values
separated by '-'. The start and end values are considered to be part of the range.

Additional upstream details:

A list of strings, each of which is a single port value or a tuple of start and end port values
separated by "-".

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2203212033320122-1121331010023001-1301233321020313-1100002302220320-3313310022002320-0220102300031101-0221111311221203-2002111321113003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.query_params` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.query_params

<a id="canonical-0021032210113322-0333013200203201-0130322031100030-0211101301022012-1220031203102000-0311031333001100-1031333021331331-3030232201313010"></a>

Type: `"list"`. Computed.

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1133032030231213-0130210230002021-1011011003310310-0130320113321223-3213213130231321-0223003321001030-1000200320200320-3232021332011313"></a>

### Direct properties for `rule_list.rules.spec.query_params`

- [check_not_present](data-sources--service_policy--reference--group-002.md#canonical-3323212223023100-3303031300300230-2013030022001023-0121132221303322-3010110302111031-3001233200321223-1132003002212021-1202023203200233): complete subsection reference.

- [check_present](data-sources--service_policy--reference--group-002.md#canonical-0022302201001000-0323103321121110-0113033312303212-1233031031102132-1211103203021230-0312121112032002-0132033113313012-3321111302332203): complete subsection reference.

<a id="canonical-0220111311323221-1001132313020012-1033021211020310-0300203230123122-0131220111200032-2212102210012111-3210312220230120-3023022012110233"></a>

<a id="canonical-2102231221322321-3233100031021000-2332303133013300-0123220132021321-1311210033233112-2133013321233020-2010101032231200-1200333300012213"></a>

#### `rule_list.rules.spec.query_params.invert_matcher` property

Type: `"bool"`. Computed.

Invert Query Parameter Matcher. Invert the match result.

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

- [item](data-sources--service_policy--reference--group-002.md#canonical-2122330001330032-1100012200103201-3123301101212130-1023331101333310-0001201302132021-1222211031031303-1131310020111020-3220312012112030): complete subsection reference.

<a id="canonical-2103011311212312-2032012112102200-0311323221101311-0010312312220323-0300032233313000-3122132111021110-1322111100110320-0323113310320331"></a>

<a id="canonical-3323030221120312-3223233113331113-2112013133103313-3310012230220010-3020001221100202-3030301221320232-0333222203111010-2300023032032102"></a>

#### `rule_list.rules.spec.query_params.key` property

Type: `"string"`. Computed.

A case-sensitive HTTP query parameter name.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3323212223023100-3303031300300230-2013030022001023-0121132221303322-3010110302111031-3001233200321223-1132003002212021-1202023203200233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.query_params.check_not_present` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.query_params](data-sources--service_policy--reference--group-002.md#canonical-2203212033320122-1121331010023001-1301233321020313-1100002302220320-3313310022002320-0220102300031101-0221111311221203-2002111321113003)
- rule_list.rules.spec.query_params.check_not_present

<a id="canonical-3031301021203101-3023103333000213-1030313120311310-0102322303232101-1032000211322023-2031321302210032-0232310321133300-1300222013113020"></a>

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

<a id="canonical-0022302201001000-0323103321121110-0113033312303212-1233031031102132-1211103203021230-0312121112032002-0132033113313012-3321111302332203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.query_params.check_present` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.query_params](data-sources--service_policy--reference--group-002.md#canonical-2203212033320122-1121331010023001-1301233321020313-1100002302220320-3313310022002320-0220102300031101-0221111311221203-2002111321113003)
- rule_list.rules.spec.query_params.check_present

<a id="canonical-1320133322002031-0220121202223031-1010313232312202-3021131301000330-3203310133301313-0003232022010121-2120002031030321-0022110001103210"></a>

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

<a id="canonical-2122330001330032-1100012200103201-3123301101212130-1023331101333310-0001201302132021-1222211031031303-1131310020111020-3220312012112030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.query_params.item` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.query_params](data-sources--service_policy--reference--group-002.md#canonical-2203212033320122-1121331010023001-1301233321020313-1100002302220320-3313310022002320-0220102300031101-0221111311221203-2002111321113003)
- rule_list.rules.spec.query_params.item

<a id="canonical-3032123310110312-1100332321020022-2031230101313200-3122001101120303-0222221202231203-3312312201103021-3010330132100200-0331022113103311"></a>

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

<a id="canonical-1331211213220212-0001120231220203-2100123103231132-2212133303332033-1133021102123333-1032330002002302-0321121110221023-0012030020133202"></a>

### Direct properties for `rule_list.rules.spec.query_params.item`

<a id="canonical-3031110320310201-3213010203020133-1310031020002210-0122202023110333-2211020032131211-0033133300022200-2302013123303130-3332300303002220"></a>

#### `rule_list.rules.spec.query_params.item.exact_values` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2033303300020230-2200230100301121-0321021232131000-2213031322111010-3000112101020022-0021303000110010-0032001311313020-0301012201202311"></a>

<a id="canonical-3313003032211102-3211221103332001-0022032311030330-0300201213302332-1330203223020122-0221003013200103-0211120221132211-1233103022221001"></a>

#### `rule_list.rules.spec.query_params.item.regex_values` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2012231103001332-3203110113113021-2312122100101111-3101210123200311-2010223300311002-0213030222130002-1122211202301003-1233033133111212"></a>

<a id="canonical-2213312003103221-1022012303210111-0213111121013113-1231022102222111-1103002112313231-2031013232312032-0012213200111333-3232100223003033"></a>

#### `rule_list.rules.spec.query_params.item.transformers` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.request_constraints` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.request_constraints

<a id="canonical-2112123331021132-0331130211222303-1220320200120130-2123332130032032-1232330002003032-1003011220032133-0200332210010013-3303233302321003"></a>

Type: `"single"`. Computed.

Configuration parameter for request constraints.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-max_cookie_count_choice": "[\"max_cookie_count_exceeds\",\"max_cookie_count_none\"]",
  "x-ves-oneof-field-max_cookie_key_size_choice": "[\"max_cookie_key_size_exceeds\",\"max_cookie_key_size_none\"]",
  "x-ves-oneof-field-max_cookie_value_size_choice": "[\"max_cookie_value_size_exceeds\",\"max_cookie_value_size_none\"]",
  "x-ves-oneof-field-max_header_count_choice": "[\"max_header_count_exceeds\",\"max_header_count_none\"]",
  "x-ves-oneof-field-max_header_key_size_choice": "[\"max_header_key_size_exceeds\",\"max_header_key_size_none\"]",
  "x-ves-oneof-field-max_header_value_size_choice": "[\"max_header_value_size_exceeds\",\"max_header_value_size_none\"]",
  "x-ves-oneof-field-max_parameter_count_choice": "[\"max_parameter_count_exceeds\",\"max_parameter_count_none\"]",
  "x-ves-oneof-field-max_parameter_name_size_choice": "[\"max_parameter_name_size_exceeds\",\"max_parameter_name_size_none\"]",
  "x-ves-oneof-field-max_parameter_value_size_choice": "[\"max_parameter_value_size_exceeds\",\"max_parameter_value_size_none\"]",
  "x-ves-oneof-field-max_query_size_choice": "[\"max_query_size_exceeds\",\"max_query_size_none\"]",
  "x-ves-oneof-field-max_request_line_size_choice": "[\"max_request_line_size_exceeds\",\"max_request_line_size_none\"]",
  "x-ves-oneof-field-max_request_size_choice": "[\"max_request_size_exceeds\",\"max_request_size_none\"]",
  "x-ves-oneof-field-max_url_size_choice": "[\"max_url_size_exceeds\",\"max_url_size_none\"]"
}
```

<a id="canonical-3023321122200311-0122113303023133-0111010223332121-1101031223201220-2113103320120003-0103111103332302-3113232111133011-2112003031102302"></a>

### Direct properties for `rule_list.rules.spec.request_constraints`

<a id="canonical-1131210031223213-1132313020110003-1011223120333302-2100330202123121-3223011130121221-3201303130330123-0203131231131233-1331023213220011"></a>

#### `rule_list.rules.spec.request_constraints.max_cookie_count_exceeds` property

Type: `"number"`. Computed.

Match on the Count for all Cookies that exceed this value. Exclusive with
\[max\_cookie\_count\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_cookie_count_none](data-sources--service_policy--reference--group-002.md#canonical-1233023322123212-3210231301202012-3232321110232302-3303331333320223-0000211210010011-3020010300232033-1020132130321000-1121233123310122): complete subsection reference.

<a id="canonical-3031223001320213-2021311110120300-3032323002000221-0031033113031331-2101111233203233-0012311311100212-0111322132031123-0220330111110331"></a>

<a id="canonical-3030310330131331-2310020110220302-1110322312100301-2123301231223002-2022023300033333-1002033233230232-1330202211200331-1022302031003201"></a>

#### `rule_list.rules.spec.request_constraints.max_cookie_key_size_exceeds` property

Type: `"number"`. Computed.

Exclusive with \[max\_cookie\_key\_size\_none\].

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_cookie_key_size_none](data-sources--service_policy--reference--group-002.md#canonical-0231212010201010-0223201300213000-2202331323101213-0310120010202330-3012303020202332-0002030121103011-0311200120002022-0130200302313333): complete subsection reference.

<a id="canonical-2121130112321031-1313331103233102-2020121322012102-0301033212030200-3301123312331201-3131123103231230-0001330132231032-3120133300102203"></a>

<a id="canonical-1320121030133213-3111201013121020-2332022322033002-3202312111122131-1121031222331313-0120210203033103-0022223323231112-2022300313330321"></a>

#### `rule_list.rules.spec.request_constraints.max_cookie_value_size_exceeds` property

Type: `"number"`. Computed.

Exclusive with \[max\_cookie\_value\_size\_none\].

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32768,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.lte": "32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32768"
  }
}
```

- [max_cookie_value_size_none](data-sources--service_policy--reference--group-002.md#canonical-2231100030011131-1130113322112212-2210332232213102-2023102200203201-1033302101013133-3113210120333302-0000232331321022-0013221232231122): complete subsection reference.

<a id="canonical-2230132211021102-3303102212231301-2122030321302023-3021323033330122-0131231202013311-2302202323232223-2021201232131232-1132223133210203"></a>

<a id="canonical-3302210312220302-3331023000332131-2322332032223100-0312030013300100-0333122022311211-0120201030111101-0101232203322313-1001222320321021"></a>

#### `rule_list.rules.spec.request_constraints.max_header_count_exceeds` property

Type: `"number"`. Computed.

Match on the Count for all Headers that exceed this value. Exclusive with
\[max\_header\_count\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.lte": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "40"
  }
}
```

- [max_header_count_none](data-sources--service_policy--reference--group-002.md#canonical-0223013130031300-2131021200313021-3200122220002011-1103222320211221-1330021221203101-0303311221020313-2033221331010032-3030233200013100): complete subsection reference.

<a id="canonical-3302231322211230-0022112201303001-3103031030301013-2032031303130020-1331313020302123-1321222133121203-0013010232032212-1020013221012012"></a>

<a id="canonical-1330320011113201-0031323031223220-2132213331013323-0131031012111323-3322000032103013-1211030011223213-2200232302322101-1303113031323102"></a>

#### `rule_list.rules.spec.request_constraints.max_header_key_size_exceeds` property

Type: `"number"`. Computed.

Exclusive with \[max\_header\_key\_size\_none\].

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_header_key_size_none](data-sources--service_policy--reference--group-002.md#canonical-1023233320313320-2030010120301220-0321110122003322-1022010102033132-1330230222320103-0112122133101001-1320222312022112-3023312322023121): complete subsection reference.

<a id="canonical-3331202323031333-1320202023223301-3001120321133101-3100313222110002-2233021033000331-0101331233102332-2001120112020010-1002022013112131"></a>

<a id="canonical-0210002212121120-1113010212320131-0101023122323310-2310303300120123-2210201110012002-0210300312230303-0030021203013022-1012001221303031"></a>

#### `rule_list.rules.spec.request_constraints.max_header_value_size_exceeds` property

Type: `"number"`. Computed.

Exclusive with \[max\_header\_value\_size\_none\].

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.lte": "64000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "64000"
  }
}
```

- [max_header_value_size_none](data-sources--service_policy--reference--group-002.md#canonical-1222121330131233-3031102030122221-1230031313133222-0031223303313321-1231023011033113-0301121023330213-1022031331323200-2133232122112020): complete subsection reference.

<a id="canonical-0101110133003132-2032222120310332-3002311033110331-1333330113322011-1300032121032031-0101322011033000-3320233021023122-1201022030031223"></a>

<a id="canonical-2311213223130003-3222221113310003-2303333013301102-3122333312002110-3222310320121000-1013123130020011-1202110312123313-0001302122101200"></a>

#### `rule_list.rules.spec.request_constraints.max_parameter_count_exceeds` property

Type: `"number"`. Computed.

Exclusive with \[max\_parameter\_count\_none\].

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_parameter_count_none](data-sources--service_policy--reference--group-002.md#canonical-2123322111023022-1112333301122232-0103312312133313-0100232123231020-0233102100310221-1300100012131300-0233003200130022-1221121211322313): complete subsection reference.

<a id="canonical-2032232301003310-0000322220332033-1020311210122010-1001023223331112-3121120321302001-2022200311210120-3133313102013032-3123033213120001"></a>

<a id="canonical-0331231300130323-0011320100020112-3030121021231302-0131132333310011-3302111200011023-1111133130123221-3331011011212201-2233101210132103"></a>

#### `rule_list.rules.spec.request_constraints.max_parameter_name_size_exceeds` property

Type: `"number"`. Computed.

Exclusive with \[max\_parameter\_name\_size\_none\].

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_parameter_name_size_none](data-sources--service_policy--reference--group-002.md#canonical-3200031212011103-0120221233031230-0321321321201211-3322311002313301-2222121133100311-1212132210011122-3320000223021100-1032333013122123): complete subsection reference.

<a id="canonical-2033122330120011-2122021120213110-2100110310323022-2231120103003212-1310000210331302-2223101111230132-0021100223110221-3303113300121132"></a>

<a id="canonical-1010210000122033-0210220012302021-3322022110121200-0113311121022011-3331003102330010-1130333311122210-2132110123130202-1330122312223012"></a>

#### `rule_list.rules.spec.request_constraints.max_parameter_value_size_exceeds` property

Type: `"number"`. Computed.

Exclusive with \[max\_parameter\_value\_size\_none\].

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1073741824,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.lte": "1073741824"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1073741824"
  }
}
```

- [max_parameter_value_size_none](data-sources--service_policy--reference--group-002.md#canonical-1013302000001312-2130301022132020-1111003112120101-3132221102122221-1230011211011133-1321123133323131-2233202312121330-1300323300210221): complete subsection reference.

<a id="canonical-0032010330202320-3002001301210231-3323121111223121-3033200010213011-2231113121002333-0210211102212200-2022212130022111-3312221233131001"></a>

<a id="canonical-3120313211331033-0113311101313120-1230231313122212-3013131323322213-3231103000003330-2312103030200010-2313222302202322-3130300003333232"></a>

#### `rule_list.rules.spec.request_constraints.max_query_size_exceeds` property

Type: `"number"`. Computed.

Match on the URL Query Size that exceed this value. Exclusive with \[max\_query\_size\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

- [max_query_size_none](data-sources--service_policy--reference--group-002.md#canonical-1322121220312330-2020203201111200-2311220030220133-1230113221000100-1312023332002112-3123221103121002-2311130111202010-3113212203311312): complete subsection reference.

<a id="canonical-1233202013023031-0202311120202333-3012313030031120-2102101110133111-3210123211101232-0231032013100200-2032313323221111-0012303302203311"></a>

<a id="canonical-0100223232022222-1120220230223121-2210121210102210-3021312332230000-0110310131120003-2230303210203012-1000032301123303-0300321232103322"></a>

#### `rule_list.rules.spec.request_constraints.max_request_line_size_exceeds` property

Type: `"number"`. Computed.

Exclusive with \[max\_request\_line\_size\_none\].

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65536,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.lte": "65536"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65536"
  }
}
```

- [max_request_line_size_none](data-sources--service_policy--reference--group-002.md#canonical-3303310121220321-1333112201233030-2100110022032131-2321020113020030-3330321231211201-3102301030120222-1121101330213131-3032002032220221): complete subsection reference.

<a id="canonical-0011232231110030-3323303223203001-2113221232332112-3121131310333022-0303130033201111-0021300320323302-3322322332212231-2220102220200113"></a>

<a id="canonical-3232011011300313-3100111021120200-3133221300003122-2312030121220101-2331233321011103-1320020003130123-0102110021220322-3202013301001323"></a>

#### `rule_list.rules.spec.request_constraints.max_request_size_exceeds` property

Type: `"number"`. Computed.

Match on the Request Size that exceed this value. Exclusive with \[max\_request\_size\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65536,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.lte": "65536"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65536"
  }
}
```

- [max_request_size_none](data-sources--service_policy--reference--group-002.md#canonical-0122120303322110-0301111123212331-3221302320132302-3102111331131132-0100122223121322-0112231121010203-0223222202320320-0100221021320223): complete subsection reference.

<a id="canonical-3102321322200020-0201221200020012-2332200321023110-3021313221230132-2200132210221103-3103222112233222-0212001222221231-3120102233223013"></a>

<a id="canonical-2031211113001121-3130133032310100-1001212030312010-3330313230132020-3133232121333120-2233122322002333-0203323202123332-2233010230200331"></a>

#### `rule_list.rules.spec.request_constraints.max_url_size_exceeds` property

Type: `"number"`. Computed.

Match on the URL Size that exceed this value. Exclusive with \[max\_url\_size\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.lte": "128000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "128000"
  }
}
```

- [max_url_size_none](data-sources--service_policy--reference--group-002.md#canonical-1101220230310312-1222213030110211-1110021021323212-3323213032310020-0003122203330310-0310031302103111-0230311231310300-3102333033111103): complete subsection reference.

<a id="canonical-1233023322123212-3210231301202012-3232321110232302-3303331333320223-0000211210010011-3020010300232033-1020132130321000-1121233123310122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.request_constraints.max_cookie_count_none` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- rule_list.rules.spec.request_constraints.max_cookie_count_none

<a id="canonical-1111313123200300-3212220133232203-1300023320321201-0122130333002003-3001130311302312-1000301222001313-1323233131120130-1203222310311223"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for max cookie count none.

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

<a id="canonical-0231212010201010-0223201300213000-2202331323101213-0310120010202330-3012303020202332-0002030121103011-0311200120002022-0130200302313333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.request_constraints.max_cookie_key_size_none` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- rule_list.rules.spec.request_constraints.max_cookie_key_size_none

<a id="canonical-2021221011301022-3011212022323322-3001223110300212-0311230002132131-3310122203230113-3332032132012031-2332021012123322-3201221303223213"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for max cookie key size none.

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

<a id="canonical-2231100030011131-1130113322112212-2210332232213102-2023102200203201-1033302101013133-3113210120333302-0000232331321022-0013221232231122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.request_constraints.max_cookie_value_size_none` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- rule_list.rules.spec.request_constraints.max_cookie_value_size_none

<a id="canonical-0121330310211302-2003222211210310-0212230230301210-3330023103312300-1220132113112132-3313133311130222-0333022021331003-1113123123000333"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for max cookie value size none.

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

<a id="canonical-0223013130031300-2131021200313021-3200122220002011-1103222320211221-1330021221203101-0303311221020313-2033221331010032-3030233200013100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.request_constraints.max_header_count_none` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- rule_list.rules.spec.request_constraints.max_header_count_none

<a id="canonical-1311313201013113-1212123033301013-0201322220033001-2011103102121023-3232133201022011-1303310131012321-1200120011010321-2033301222023130"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for max header count none.

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

<a id="canonical-1023233320313320-2030010120301220-0321110122003322-1022010102033132-1330230222320103-0112122133101001-1320222312022112-3023312322023121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.request_constraints.max_header_key_size_none` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- rule_list.rules.spec.request_constraints.max_header_key_size_none

<a id="canonical-2221002200001123-3212303123011102-1222130123011013-2010123222030023-2330230120132310-1102132303311221-0333210211333013-2220232220220200"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for max header key size none.

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

<a id="canonical-1222121330131233-3031102030122221-1230031313133222-0031223303313321-1231023011033113-0301121023330213-1022031331323200-2133232122112020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.request_constraints.max_header_value_size_none` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- rule_list.rules.spec.request_constraints.max_header_value_size_none

<a id="canonical-2033001132200123-2321030230232231-2212130213013221-2121321112321123-3012121310323302-3023133310122302-3100010120101123-0003230301103103"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for max header value size none.

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

<a id="canonical-2123322111023022-1112333301122232-0103312312133313-0100232123231020-0233102100310221-1300100012131300-0233003200130022-1221121211322313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.request_constraints.max_parameter_count_none` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- rule_list.rules.spec.request_constraints.max_parameter_count_none

<a id="canonical-0313322233111221-1210021213200123-3231221212033211-3312310330100321-0201032030100102-0022133123230211-3212130210111222-2333100110021230"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for max parameter count none.

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

<a id="canonical-3200031212011103-0120221233031230-0321321321201211-3322311002313301-2222121133100311-1212132210011122-3320000223021100-1032333013122123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.request_constraints.max_parameter_name_size_none` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- rule_list.rules.spec.request_constraints.max_parameter_name_size_none

<a id="canonical-3101312003323232-0032113110210331-2321113220033213-0023130003312201-2223212030323220-1301323121000213-3320321233033012-3212333220323302"></a>

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

<a id="canonical-1013302000001312-2130301022132020-1111003112120101-3132221102122221-1230011211011133-1321123133323131-2233202312121330-1300323300210221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.request_constraints.max_parameter_value_size_none` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- rule_list.rules.spec.request_constraints.max_parameter_value_size_none

<a id="canonical-0002210320222020-1300110232313333-3012300310132032-1011303323333311-0100331302223311-2003333101032303-2301210011100320-1231223103102123"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for max parameter value size none.

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

<a id="canonical-1322121220312330-2020203201111200-2311220030220133-1230113221000100-1312023332002112-3123221103121002-2311130111202010-3113212203311312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.request_constraints.max_query_size_none` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- rule_list.rules.spec.request_constraints.max_query_size_none

<a id="canonical-2302321312102200-0233313322120033-2310132332032032-1001111220102121-3030020302032323-3321112312323322-2331303333100020-2003311223202201"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for max query size none.

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

<a id="canonical-3303310121220321-1333112201233030-2100110022032131-2321020113020030-3330321231211201-3102301030120222-1121101330213131-3032002032220221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.request_constraints.max_request_line_size_none` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- rule_list.rules.spec.request_constraints.max_request_line_size_none

<a id="canonical-3131101223332303-3021032331023101-2232323000321112-0230333123331211-0111103111303200-1323112132023102-0111110012123303-2120320320121113"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for max request line size none.

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

<a id="canonical-0122120303322110-0301111123212331-3221302320132302-3102111331131132-0100122223121322-0112231121010203-0223222202320320-0100221021320223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.request_constraints.max_request_size_none` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- rule_list.rules.spec.request_constraints.max_request_size_none

<a id="canonical-2032300221313213-2323021323221310-2230003003033312-2101131033103331-0210020313003103-3320100033301110-1102301113012302-2211132101112103"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for max request size none.

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

<a id="canonical-1101220230310312-1222213030110211-1110021021323212-3323213032310020-0003122203330310-0310031302103111-0230311231310300-3102333033111103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.request_constraints.max_url_size_none` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- rule_list.rules.spec.request_constraints.max_url_size_none

<a id="canonical-1313023023203300-0032003021013100-1302123103200301-2121223010312110-2013102002212233-1002000013001322-0121231202230231-0213321321331003"></a>

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

<a id="canonical-0121322232330011-2001200302133213-0020123132212012-0113020311112312-0201321011303010-1302103121002333-1323212012233310-1210231031031002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.segment_policy` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.segment_policy

<a id="canonical-3323232232032330-1122211331303130-1023302210003331-0013122122013120-1002031300033001-2331202230203020-1031113221232130-1002212101031213"></a>

Type: `"single"`. Computed.

Configure source and destination segment for policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dst_segment_choice": "[\"dst_any\",\"dst_segments\",\"intra_segment\"]",
  "x-ves-oneof-field-src_segment_choice": "[\"src_any\",\"src_segments\"]"
}
```

<a id="canonical-1212122133300330-2012112022120131-3312010002330001-3323131003002210-3032033112011110-0131312300031022-2312221131031312-3023211301030302"></a>

### Direct properties for `rule_list.rules.spec.segment_policy`

- [dst_any](data-sources--service_policy--reference--group-002.md#canonical-0031122310033231-2010230020011010-1303321201322031-3320020223231331-2303031122211232-0331300200012303-1021230330130131-2303123000000020): complete subsection reference.

- [dst_segments](data-sources--service_policy--reference--group-002.md#canonical-0011131210322333-3303310012013300-3023111203210330-0020021112012101-0303233333320023-2333301023133001-0000020301032201-2310201111202123): complete subsection reference.

- [intra_segment](data-sources--service_policy--reference--group-002.md#canonical-0203323232233001-2120013110020322-0221121112113023-2320113023302210-2100223233232320-2310322302003122-3100131032123133-0222202310211010): complete subsection reference.

- [src_any](data-sources--service_policy--reference--group-002.md#canonical-1300223022331330-1313302231013010-3330313313133130-2202101212123102-3202302012303311-2032313311113131-3222330200001323-3020110200313330): complete subsection reference.

- [src_segments](data-sources--service_policy--reference--group-002.md#canonical-0321011030032012-0311231002031103-3002101310212323-0023110012220020-1300300231320033-2222023130213033-3113130312302023-2330031210112233): complete subsection reference.

<a id="canonical-0031122310033231-2010230020011010-1303321201322031-3320020223231331-2303031122211232-0331300200012303-1021230330130131-2303123000000020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.segment_policy.dst_any` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.segment_policy](data-sources--service_policy--reference--group-002.md#canonical-0121322232330011-2001200302133213-0020123132212012-0113020311112312-0201321011303010-1302103121002333-1323212012233310-1210231031031002)
- rule_list.rules.spec.segment_policy.dst_any

<a id="canonical-1321312333231211-2012322212003330-1002310223333030-1202100132013202-2030132101123031-1130131223203111-1023332120210333-2122311100310133"></a>

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

<a id="canonical-0011131210322333-3303310012013300-3023111203210330-0020021112012101-0303233333320023-2333301023133001-0000020301032201-2310201111202123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.segment_policy.dst_segments` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.segment_policy](data-sources--service_policy--reference--group-002.md#canonical-0121322232330011-2001200302133213-0020123132212012-0113020311112312-0201321011303010-1302103121002333-1323212012233310-1210231031031002)
- rule_list.rules.spec.segment_policy.dst_segments

<a id="canonical-0031213203022332-1022112230002233-2203321002120022-1133112331111202-3300311331010213-0022102322032123-1023020002302211-3001021021023002"></a>

Type: `"single"`. Computed.

Configuration parameter for dst segments.

Additional upstream details:

List of references to Segments.

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

<a id="canonical-3032120203313332-0030231301012021-2113002311220303-3313133130301231-3223033122133303-2310132210012213-0321001332320001-1210300220102320"></a>

### Direct properties for `rule_list.rules.spec.segment_policy.dst_segments`

- [segments](data-sources--service_policy--reference--group-002.md#canonical-3201210223131120-2023030303012223-3333023031202332-0122002233013211-3110031303331122-0330302011320032-1333221323213100-3101330210002213): complete subsection reference.

<a id="canonical-3201210223131120-2023030303012223-3333023031202332-0122002233013211-3110031303331122-0330302011320032-1333221323213100-3101330210002213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.segment_policy.dst_segments.segments` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.segment_policy](data-sources--service_policy--reference--group-002.md#canonical-0121322232330011-2001200302133213-0020123132212012-0113020311112312-0201321011303010-1302103121002333-1323212012233310-1210231031031002)
- [rule_list.rules.spec.segment_policy.dst_segments](data-sources--service_policy--reference--group-002.md#canonical-0011131210322333-3303310012013300-3023111203210330-0020021112012101-0303233333320023-2333301023133001-0000020301032201-2310201111202123)
- rule_list.rules.spec.segment_policy.dst_segments.segments

<a id="canonical-1002310132133310-0201232300331013-0301123313122333-1131113103202010-3220122322323201-2201223110003021-2303112112230112-1303201201010000"></a>

Type: `"list"`. Computed.

Segments. Select list of segments.

Receipt-pinned upstream constraints:

```json
{
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

<a id="canonical-0133201211002022-1220032231310303-0333001321030011-3232300000333330-1012133310131311-2203210220003333-0013203101210322-0033022231122312"></a>

### Direct properties for `rule_list.rules.spec.segment_policy.dst_segments.segments`

<a id="canonical-1221110230301310-0131201212321122-0113003003312103-0232010031031121-1110302312211233-2030202321211333-0203333032201300-0223232330003130"></a>

#### `rule_list.rules.spec.segment_policy.dst_segments.segments.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2120202321223210-2122112212030221-2221021330201121-1013023030301202-0300121210330333-0100311130313211-3013202223322221-1012320300222201"></a>

<a id="canonical-0313312202320121-3121120233312213-1101223010010122-1010101010033210-2313323032011031-0302203201220103-1321330213311032-0223202110300212"></a>

#### `rule_list.rules.spec.segment_policy.dst_segments.segments.namespace` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3011301032103032-1233300121303113-0031122021013133-2220333002203010-2233131313023311-1001333033100330-3032030220010002-3233032100103301"></a>

<a id="canonical-1210310032010030-2211121313103110-1023103020221220-3331330320110220-1123102212131231-2210021231202320-3001020212021213-1032323103111030"></a>

#### `rule_list.rules.spec.segment_policy.dst_segments.segments.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0203323232233001-2120013110020322-0221121112113023-2320113023302210-2100223233232320-2310322302003122-3100131032123133-0222202310211010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.segment_policy.intra_segment` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.segment_policy](data-sources--service_policy--reference--group-002.md#canonical-0121322232330011-2001200302133213-0020123132212012-0113020311112312-0201321011303010-1302103121002333-1323212012233310-1210231031031002)
- rule_list.rules.spec.segment_policy.intra_segment

<a id="canonical-1230110130222112-0033112022211020-0200023213002330-1223122322211222-2123221023313332-3013022001202203-0001312021201121-0033213231022131"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for intra segment.

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

<a id="canonical-1300223022331330-1313302231013010-3330313313133130-2202101212123102-3202302012303311-2032313311113131-3222330200001323-3020110200313330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.segment_policy.src_any` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.segment_policy](data-sources--service_policy--reference--group-002.md#canonical-0121322232330011-2001200302133213-0020123132212012-0113020311112312-0201321011303010-1302103121002333-1323212012233310-1210231031031002)
- rule_list.rules.spec.segment_policy.src_any

<a id="canonical-1101103031230102-0230020121302021-2001333033001031-3023002312220232-1302000131223222-3313132130030110-0023031003332310-3132002333101203"></a>

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

<a id="canonical-0321011030032012-0311231002031103-3002101310212323-0023110012220020-1300300231320033-2222023130213033-3113130312302023-2330031210112233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.segment_policy.src_segments` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.segment_policy](data-sources--service_policy--reference--group-002.md#canonical-0121322232330011-2001200302133213-0020123132212012-0113020311112312-0201321011303010-1302103121002333-1323212012233310-1210231031031002)
- rule_list.rules.spec.segment_policy.src_segments

<a id="canonical-2301133103201303-2001023333321101-0132010032100323-1030203123002021-1100021200021233-1320101333332010-2230331010212112-2210232020122310"></a>

Type: `"single"`. Computed.

Configuration parameter for src segments.

Additional upstream details:

List of references to Segments.

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

<a id="canonical-0012001022002231-3231221021230311-2200012011213300-3310002110023321-0112222001022303-3030002232302212-2311030323000321-2223202133232102"></a>

### Direct properties for `rule_list.rules.spec.segment_policy.src_segments`

- [segments](data-sources--service_policy--reference--group-002.md#canonical-1313121023332203-0102232233233030-0032313202313131-1121121011302313-1103332101020131-1030113133012303-3010003230212213-3030323032021001): complete subsection reference.

<a id="canonical-1313121023332203-0102232233233030-0032313202313131-1121121011302313-1103332101020131-1030113133012303-3010003230212213-3030323032021001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.segment_policy.src_segments.segments` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.segment_policy](data-sources--service_policy--reference--group-002.md#canonical-0121322232330011-2001200302133213-0020123132212012-0113020311112312-0201321011303010-1302103121002333-1323212012233310-1210231031031002)
- [rule_list.rules.spec.segment_policy.src_segments](data-sources--service_policy--reference--group-002.md#canonical-0321011030032012-0311231002031103-3002101310212323-0023110012220020-1300300231320033-2222023130213033-3113130312302023-2330031210112233)
- rule_list.rules.spec.segment_policy.src_segments.segments

<a id="canonical-3010233120011013-0220003003232100-1220130101303213-0232203100213013-1122133000103112-1033220213020101-0210122121030133-1030030213021023"></a>

Type: `"list"`. Computed.

Segments. Select list of segments.

Receipt-pinned upstream constraints:

```json
{
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

<a id="canonical-0122013102010233-2311332233011031-0100001130311022-0103232202333313-2022100013120320-0123301033210033-3203033101102130-2013201202133323"></a>

### Direct properties for `rule_list.rules.spec.segment_policy.src_segments.segments`

<a id="canonical-2320133321111300-3211202300330000-3123020210222102-0112120102130231-2330211230222122-1100212120320010-2221233321323330-1101320131012010"></a>

#### `rule_list.rules.spec.segment_policy.src_segments.segments.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1230301122130022-0203020120330023-2330000013103220-1310100333320033-2213100010322130-2002032220100312-0211030220313133-2322310232202113"></a>

<a id="canonical-1130001223202333-1200023213022103-0123313301033332-2103210203031021-0221202113013201-3100203110020132-2111320112122022-3331231300131310"></a>

#### `rule_list.rules.spec.segment_policy.src_segments.segments.namespace` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1321133222213232-3313200100232002-3231022123330231-0130302331000310-2133101112123201-1030323223131133-1331313020103033-2312303012113023"></a>

<a id="canonical-2033220101012103-1101223022030002-1232113032011222-1201230300301122-2233013200223011-0133322013031120-1002312311211311-1010131233000333"></a>

#### `rule_list.rules.spec.segment_policy.src_segments.segments.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1333003020033103-2303031332133313-1101333013023130-2313101003001221-0111232133300132-2011202312030000-1012231232313131-2333220102021302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.tls_fingerprint_matcher` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.tls_fingerprint_matcher

<a id="canonical-3311100032302102-2323202000133230-2023032001333021-1201210301202211-1223001232030332-3002102203101103-0312130010013122-1133130200001323"></a>

Type: `"single"`. Computed.

A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are satisfied
and the input fingerprint is not one of the excluded values.

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

<a id="canonical-3223030322130232-2301333120002212-0313121223323103-2211303110300200-1312023032102010-2201223211013232-2003312203321012-2122022111022212"></a>

### Direct properties for `rule_list.rules.spec.tls_fingerprint_matcher`

<a id="canonical-0333121223330310-3322313032201230-0013120222103300-3003203120231123-1332230003322001-0112022203101232-0013231101021231-2321333333332220"></a>

#### `rule_list.rules.spec.tls_fingerprint_matcher.classes` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2011231012212202-2021321320303201-0312332010222313-2222120231002131-2330133220203113-0021122300313021-3012121121312031-0223303013010030"></a>

<a id="canonical-3021332033023300-1100030030323021-3332330312003302-3300121203132222-2022001123120030-3313222011001212-2121330020012113-3310011313202110"></a>

#### `rule_list.rules.spec.tls_fingerprint_matcher.exact_values` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1200131312313331-0021301110001231-2020033301023203-2020100202230221-0300203113201130-1333021011032201-1312231123332333-0000113031133221"></a>

<a id="canonical-3032021033100131-0102302103220132-2001033222300213-2223103212310100-2112223300023313-2331133001233312-1011031233110312-0303210133230320"></a>

#### `rule_list.rules.spec.tls_fingerprint_matcher.excluded_values` property

Type: `["list", "string"]`. Computed.

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3302232320022303-1301323100000232-2222223202031122-2132301121130100-3212113133030300-0223231120230100-3101210030303312-3003333002303311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.user_identity_matcher` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.user_identity_matcher

<a id="canonical-0100212130233102-1310321312010333-2030333021002020-0223013210212033-3201132300013202-3220332031312301-2333203202012002-0211003212223012"></a>

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

<a id="canonical-0033201132100321-0013130130220213-0030310103320100-2111210110220103-1002100121300022-3332313130300113-1111000001223032-0331221030322022"></a>

### Direct properties for `rule_list.rules.spec.user_identity_matcher`

<a id="canonical-2233132333222131-0231231101030331-1223112100320122-1323202021302012-0122002132001213-2302112323131103-0232310133211212-3132022030202221"></a>

#### `rule_list.rules.spec.user_identity_matcher.exact_values` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1131312023312021-1230101210102111-2311331301310200-3201232002222003-2202210120300210-2231031302101120-3221231123211010-3233111032100223"></a>

<a id="canonical-2332121312220322-0000012112030110-2000310123230101-3231310101022030-0302213012331001-0223233120321010-0313232020310120-3012321320012030"></a>

#### `rule_list.rules.spec.user_identity_matcher.regex_values` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1321030020301232-0323133310020002-1031312201202122-2313331303112210-3101113223232332-0322221031012323-3201210323110133-2121002301220100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.waf_action` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.waf_action

<a id="canonical-2201213013321111-0021023302222322-3012030010120000-0313121021213130-0213311111011011-3132001123022010-0013200303120001-0010312221111010"></a>

Type: `"single"`. Computed.

Modify App Firewall behavior for a matching request. The modification could either be to entirely
skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall
Rule Control settings.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_type": "[\"app_firewall_detection_control\",\"none\",\"waf_skip_processing\"]"
}
```

<a id="canonical-2011021110231330-0233011111310331-2330303310202331-2121232012310221-1130120220320201-2121302221100023-1223333013223331-3223212020233013"></a>

### Direct properties for `rule_list.rules.spec.waf_action`

- [app_firewall_detection_control](data-sources--service_policy--reference--group-002.md#canonical-0312203122000130-1211201300112130-2033131032302133-2131001020021211-1223121201301131-3203110303111102-2022322010100302-0112200311232020): complete subsection reference.

- [none](data-sources--service_policy--reference--group-003.md#canonical-0310123023311333-0022102211121010-0020210033011011-2001230110222211-2210213330032131-3212220310101201-1223020211031333-3230113300311321): complete subsection reference.

- [waf_skip_processing](data-sources--service_policy--reference--group-003.md#canonical-0000112103021210-0222132122221021-1220123212112130-2313130013212221-2010010233130222-3122031033113002-0031013220303221-1020013000133310): complete subsection reference.

<a id="canonical-0312203122000130-1211201300112130-2033131032302133-2131001020021211-1223121201301131-3203110303111102-2022322010100302-0112200311232020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.waf_action.app_firewall_detection_control` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.waf_action](data-sources--service_policy--reference--group-002.md#canonical-1321030020301232-0323133310020002-1031312201202122-2313331303112210-3101113223232332-0322221031012323-3201210323110133-2121002301220100)
- rule_list.rules.spec.waf_action.app_firewall_detection_control

<a id="canonical-0020213223313321-2120133310320131-2011210000132312-3113301020021323-0302030302133130-0203011310312310-0112222313103012-1001301021303002"></a>

Type: `"single"`. Computed.

Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded
from triggering on the defined match criteria.

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

<a id="canonical-3222213010011303-2110200332322113-2000020131200211-3122203321311201-0130233323000303-3101233122100120-1032212112321120-1323032011003021"></a>

### Direct properties for `rule_list.rules.spec.waf_action.app_firewall_detection_control`

- [exclude_attack_type_contexts](data-sources--service_policy--reference--group-002.md#canonical-3323003110031030-2201330032221222-2000010213012021-2131130212001113-2012322012011313-1201333120303103-0000122123120001-1011102131322323): complete subsection reference.

- [exclude_bot_name_contexts](data-sources--service_policy--reference--group-002.md#canonical-3223230210221003-2032232113223131-0333230231131301-2132210130000201-2120221212200230-0200130213332300-2002232320022030-0013000222011030): complete subsection reference.

- [exclude_signature_contexts](data-sources--service_policy--reference--group-002.md#canonical-2031101212110222-0203330312200223-3230231021002311-3022111302033032-3230123020220100-3001321120131013-0033323300220123-3122012332222033): complete subsection reference.

- [exclude_violation_contexts](data-sources--service_policy--reference--group-003.md#canonical-0100312021032302-3122211211201132-3012220103230032-1212223321010131-0230222232110210-0131112003213120-0132312221300330-3223101002023033): complete subsection reference.

<a id="canonical-3323003110031030-2201330032221222-2000010213012021-2131130212001113-2012322012011313-1201333120303103-0000122123120001-1011102131322323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.waf_action](data-sources--service_policy--reference--group-002.md#canonical-1321030020301232-0323133310020002-1031312201202122-2313331303112210-3101113223232332-0322221031012323-3201210323110133-2121002301220100)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control](data-sources--service_policy--reference--group-002.md#canonical-0312203122000130-1211201300112130-2033131032302133-2131001020021211-1223121201301131-3203110303111102-2022322010100302-0112200311232020)
- rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts

<a id="canonical-0313223222013300-2223010213233203-3000112212111221-1300011232121130-1012212112013302-3001330322220022-1201310312333133-3330230221001133"></a>

Type: `"list"`. Computed.

Exclude an entire attack type only in the named context. For migrated per-parameter exceptions,
prefer this over signature-ID exclusions because one payload can trigger several signatures;
unrelated parameters and attack types remain protected.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3221221322302021-1111333311102223-3113100332012122-0110210133301303-2201213010310311-1122013132221021-1033310001002021-1032133323101230"></a>

### Direct properties for `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts`

<a id="canonical-3121103133132030-0130203123101231-3223300333120333-0011100233023020-2101231323232111-1112123213130202-3020010302003300-0321200300301321"></a>

#### `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context` property

Type: `"string"`. Computed.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Additional upstream details:

Exclusion scope. Use CONTEXT\_PARAMETER with context\_name for one parameter, CONTEXT\_COOKIE for
one cookie, or CONTEXT\_ANY only for an intentionally global scope.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3110231203333213-2100132103102023-2002232113203223-3011311300330011-2220312103212200-1031213011003321-2002323130010303-1322133302203303"></a>

<a id="canonical-1120232023302023-2133031311132100-0031100100202320-2320013201002321-1001301121021131-3330033020121033-0123332012132030-2231102331330121"></a>

#### `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context_name` property

Type: `"string"`. Computed.

Parameter, cookie, or header name selected by context. For a parameter-scoped WAF exception, set
context to CONTEXT\_PARAMETER and name only the intended parameter.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-2012101323113130-2322301023023333-1001022231231003-1131310120203112-2123010102302122-1332001302033030-2332303331031232-0120332211102301"></a>

<a id="canonical-3322301133220330-1203231312023120-3121012012002303-1210302103331123-2323131112323033-0101332000321031-2202232003111010-3032303202331013"></a>

#### `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type` property

Type: `"string"`. Computed.

\[Enum:
ATTACK\_TYPE\_NONE|ATTACK\_TYPE\_NON\_BROWSER\_CLIENT|ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS|ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE|ATTACK\_TYPE\_DETECTION\_EVASION|ATTACK\_TYPE\_VULNERABILITY\_SCAN|ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY|ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS|ATTACK\_TYPE\_BUFFER\_OVERFLOW|ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION|ATTACK\_TYPE\_INFORMATION\_LEAKAGE|ATTACK\_TYPE\_DIRECTORY\_INDEXING|ATTACK\_TYPE\_PATH\_TRAVERSAL|ATTACK\_TYPE\_XPATH\_INJECTION|ATTACK\_TYPE\_LDAP\_INJECTION|ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION|ATTACK\_TYPE\_COMMAND\_EXECUTION|ATTACK\_TYPE\_SQL\_INJECTION|ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING|ATTACK\_TYPE\_DENIAL\_OF\_SERVICE|ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK|ATTACK\_TYPE\_SESSION\_HIJACKING|ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING|ATTACK\_TYPE\_FORCEFUL\_BROWSING|ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE|ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD|ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\]
List of all Attack Types ATTACK\_TYPE\_NONE ATTACK\_TYPE\_NON\_BROWSER\_CLIENT
ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE
ATTACK\_TYPE\_DETECTION\_EVASION ATTACK\_TYPE\_VULNERABILITY\_SCAN
ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS..
Possible values are \`ATTACK\_TYPE\_NONE\`, \`ATTACK\_TYPE\_NON\_BROWSER\_CLIENT\`,
\`ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS\`, \`ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE\`,
\`ATTACK\_TYPE\_DETECTION\_EVASION\`, \`ATTACK\_TYPE\_VULNERABILITY\_SCAN\`,
\`ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY\`,
\`ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS\`, \`ATTACK\_TYPE\_BUFFER\_OVERFLOW\`,
\`ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION\`, \`ATTACK\_TYPE\_INFORMATION\_LEAKAGE\`,
\`ATTACK\_TYPE\_DIRECTORY\_INDEXING\`, \`ATTACK\_TYPE\_PATH\_TRAVERSAL\`,
\`ATTACK\_TYPE\_XPATH\_INJECTION\`, \`ATTACK\_TYPE\_LDAP\_INJECTION\`,
\`ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION\`, \`ATTACK\_TYPE\_COMMAND\_EXECUTION\`,
\`ATTACK\_TYPE\_SQL\_INJECTION\`, \`ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING\`,
\`ATTACK\_TYPE\_DENIAL\_OF\_SERVICE\`, \`ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK\`,
\`ATTACK\_TYPE\_SESSION\_HIJACKING\`, \`ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING\`,
\`ATTACK\_TYPE\_FORCEFUL\_BROWSING\`, \`ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE\`,
\`ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD\`, \`ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\`. Defaults to
\`ATTACK\_TYPE\_NONE\`.

Additional upstream details:

Attack-type enum excluded in this context, for example ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING. Other
attack types remain enforced.

Receipt-pinned upstream constraints:

```json
{
  "default": "ATTACK_TYPE_NONE",
  "enum": [
    "ATTACK_TYPE_NONE",
    "ATTACK_TYPE_NON_BROWSER_CLIENT",
    "ATTACK_TYPE_OTHER_APPLICATION_ATTACKS",
    "ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE",
    "ATTACK_TYPE_DETECTION_EVASION",
    "ATTACK_TYPE_VULNERABILITY_SCAN",
    "ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY",
    "ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS",
    "ATTACK_TYPE_BUFFER_OVERFLOW",
    "ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION",
    "ATTACK_TYPE_INFORMATION_LEAKAGE",
    "ATTACK_TYPE_DIRECTORY_INDEXING",
    "ATTACK_TYPE_PATH_TRAVERSAL",
    "ATTACK_TYPE_XPATH_INJECTION",
    "ATTACK_TYPE_LDAP_INJECTION",
    "ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION",
    "ATTACK_TYPE_COMMAND_EXECUTION",
    "ATTACK_TYPE_SQL_INJECTION",
    "ATTACK_TYPE_CROSS_SITE_SCRIPTING",
    "ATTACK_TYPE_DENIAL_OF_SERVICE",
    "ATTACK_TYPE_HTTP_PARSER_ATTACK",
    "ATTACK_TYPE_SESSION_HIJACKING",
    "ATTACK_TYPE_HTTP_RESPONSE_SPLITTING",
    "ATTACK_TYPE_FORCEFUL_BROWSING",
    "ATTACK_TYPE_REMOTE_FILE_INCLUDE",
    "ATTACK_TYPE_MALICIOUS_FILE_UPLOAD",
    "ATTACK_TYPE_GRAPHQL_PARSER_ATTACK"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3223230210221003-2032232113223131-0333230231131301-2132210130000201-2120221212200230-0200130213332300-2002232320022030-0013000222011030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.waf_action](data-sources--service_policy--reference--group-002.md#canonical-1321030020301232-0323133310020002-1031312201202122-2313331303112210-3101113223232332-0322221031012323-3201210323110133-2121002301220100)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control](data-sources--service_policy--reference--group-002.md#canonical-0312203122000130-1211201300112130-2033131032302133-2131001020021211-1223121201301131-3203110303111102-2022322010100302-0112200311232020)
- rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts

<a id="canonical-0100022320130212-0121002013023332-2121022130012331-1013120021320110-1103233103001031-2020212020100033-3013123002120033-3213112223212333"></a>

Type: `"list"`. Computed.

Bot Names to be excluded for the defined match criteria.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3302002303301211-2200232120012101-0233022022210213-2023002330230231-1212101030120030-2301002023202110-3112311023301011-3132300001333220"></a>

### Direct properties for `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts`

<a id="canonical-1223031212301030-3302120030211111-2311012131320230-3220013222002312-1013232001233222-1110102202001032-1012112103321022-3322203213131131"></a>

#### `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts.bot_name` property

Type: `"string"`. Computed.

Bot Name. Human-readable name for the resource

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2031101212110222-0203330312200223-3230231021002311-3022111302033032-3230123020220100-3001321120131013-0033323300220123-3122012332222033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.waf_action](data-sources--service_policy--reference--group-002.md#canonical-1321030020301232-0323133310020002-1031312201202122-2313331303112210-3101113223232332-0322221031012323-3201210323110133-2121002301220100)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control](data-sources--service_policy--reference--group-002.md#canonical-0312203122000130-1211201300112130-2033131032302133-2131001020021211-1223121201301131-3203110303111102-2022322010100302-0112200311232020)
- rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts

<a id="canonical-3232303210031312-3033010222203121-0320100002303131-2030300312300110-2123132031302323-3331330031132312-1222222311001223-1320310302201320"></a>

Type: `"list"`. Computed.

Signature IDs to be excluded for the defined match criteria.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0003103011233303-0220032110331003-0300132320032003-1101023111011113-3131303122320120-2111033020311213-0111032033203130-2331321312221202"></a>

### Direct properties for `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts`

<a id="canonical-1332010312113312-1222033200113132-3130310022333322-2022301122032021-3102112122101313-3310200131020332-1012201313030131-1202232330323011"></a>

#### `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.context` property

Type: `"string"`. Computed.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Additional upstream details:

The available contexts for Exclusion rules.

&#8203;- CONTEXT\_ANY: CONTEXT\_ANY

Detection will be excluded for all contexts. &#8203;- CONTEXT\_BODY: CONTEXT\_BODY

Detection will be excluded for the request body. &#8203;- CONTEXT\_REQUEST: CONTEXT\_REQUEST

Detection will be excluded for the request. &#8203;- CONTEXT\_RESPONSE: CONTEXT\_RESPONSE

&#8203;- CONTEXT\_PARAMETER: CONTEXT\_PARAMETER

Detection will be excluded for the parameters. The parameter name is required in the Context name
field. If the field is left empty, the detection will be excluded for all parameters. &#8203;-
CONTEXT\_HEADER: CONTEXT\_HEADER

Detection will be excluded for the headers. The header name is required in the Context name field.
If the field is left empty, the detection will be excluded for all headers. &#8203;-
CONTEXT\_COOKIE: CONTEXT\_COOKIE

Detection will be excluded for the cookies. The cookie name is required in the Context name field.
If the field is left empty, the detection will be excluded for all cookies. &#8203;- CONTEXT\_URL:
CONTEXT\_URL

Detection will be excluded for the request URL. &#8203;- CONTEXT\_URI: CONTEXT\_URI.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0020111011133001-3233222011223222-1131010313311201-3230000233201331-1123103231210002-1232303332322002-2301321003132221-3032102300223220"></a>
