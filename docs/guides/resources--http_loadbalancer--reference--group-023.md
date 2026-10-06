---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-0300213200120212-0123302231201203-3233312031002311-0211102003111120-1311030013032310-2202000211131023-0303111203212021-0300131330131301"></a>

## `policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets.tenant` property

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

<a id="canonical-0212030212030313-1320332202232022-1320023311300203-1131000132123221-1222323212112202-0001200211302323-0233231320133300-0300232321131022"></a>

<a id="canonical-2312012303230012-1311213303330030-0211111103303233-1220321213132132-0133023121122111-1211011010311111-1223022030120032-2302032203313023"></a>

## `policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets.uid` property

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

<a id="canonical-1322031002303310-1111010110113010-3010000120211223-0112031202323033-3213001011000120-0032032230230101-3220223203111233-0023030020301031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.body_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.body_matcher

<a id="canonical-1120202020002120-0033002221031111-1101231111313331-0021301303031310-2020011020322002-3003230313003111-2203330013030122-0031120312000232"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
body_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-2300031320131112-3002112103211213-1312300331303122-0333230213013133-3323201332002202-3023103313131333-2330122111012121-2131113320221102"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.body_matcher`

<a id="canonical-1222110331201200-0100012130110120-0311010032103212-2031202223133001-3231120223330201-0232332132131331-0103113100303210-3212301233203210"></a>

#### `policy_based_challenge.rule_list.rules.spec.body_matcher.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-2101100123321313-3131330133110131-1011331311203301-2211111320222320-3130322230311101-1210230132203313-3132331112012122-2220132303230232"></a>

<a id="canonical-1223231322033311-1030001011232201-2121132101000131-2333313100323020-3020002333233102-1302123222012323-0022323023101333-2002300303030131"></a>

#### `policy_based_challenge.rule_list.rules.spec.body_matcher.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-1212100332130320-1012101311113323-1300030002112003-2133222013231000-3333013213332003-3033311011131030-0211010201021202-3333323210232100"></a>

<a id="canonical-0312111021133203-2210301111312123-0101311333122201-0203010133211100-0320301222332031-2302301312223110-3222301232330101-0330212010103302"></a>

#### `policy_based_challenge.rule_list.rules.spec.body_matcher.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3301001033202231-2133133323213110-2321233202103301-1310201321100023-1302201121201002-3200232221023023-0110031303001223-0212313222223133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.client_selector` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.client_selector

<a id="canonical-0121031210120212-2200023300010320-1303230332213032-2021223200323330-3021322303321000-3022022012331003-0202202311020323-2011303233000013"></a>

Type: `"object"`. single nested block, Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
client_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-0003111203100113-2100031022032131-2303230131013220-2223113232220002-2321003020230210-3003101202312112-2122203331202031-2331101203001132"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.client_selector`

<a id="canonical-0302312133102010-0010310201312322-2202231331120321-0223332301122022-1230212203132110-0022023102313323-3201133310001231-1223002021111220"></a>

#### `policy_based_challenge.rule_list.rules.spec.client_selector.expressions` property

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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

<a id="canonical-1112223030321202-2310131032330100-2302103332230131-0001200330203221-2212303302110122-3033121110232100-1130300130302113-3232330311033030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.cookie_matchers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers

<a id="canonical-2311302301030200-3113330322112312-3320022100223123-0111322103012232-3103312312332003-3310211202231110-2203000200033132-2003313132233310"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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

Terraform syntax:

```terraform
cookie_matchers {
  # Configure direct properties listed below.
}
```

<a id="canonical-2212230230323221-2130110112023221-0120331300331232-1001103221301133-3010102222223120-3232211311030100-0222312002030101-0111220332200213"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.cookie_matchers`

- [check_not_present](resources--http_loadbalancer--reference--group-023.md#canonical-0103313330202032-0333323222300211-3231010121102231-1022300331232331-1311002322210132-3210201002132031-2200001313310231-1132120002030213): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-023.md#canonical-1011011121103030-1023321311310333-2210310200001020-2031223022321101-0211112103002331-0113300012011220-0213102312210012-3331322211331322): complete subsection reference.

<a id="canonical-2320031132013223-2200212003011211-1122112213300011-3222201101010102-1323020003210220-3100132203131102-2000300332231321-3032220311311221"></a>

<a id="canonical-3002210022200332-2222312020011021-0332020211113222-3212033201111133-3213101231230110-3222031333222202-3312333232032233-0022013233230133"></a>

#### `policy_based_challenge.rule_list.rules.spec.cookie_matchers.invert_matcher` property

Type: `"bool"`. Optional.

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

- [item](resources--http_loadbalancer--reference--group-023.md#canonical-2301120303000033-1120110232011011-0311033303003031-3013101323203110-3103223000001020-2012122203222232-3020031332212020-3132031221321231): complete subsection reference.

<a id="canonical-2320220323220232-0222321220003111-1032113000013020-1302313312110003-3012331330201301-3333220103101322-3232200232133112-1121332003000322"></a>

<a id="canonical-3230203020331111-3113323233103233-0311121130001332-3320010001133222-3331111331133323-0320012020130333-1100210211221310-2313101122303120"></a>

#### `policy_based_challenge.rule_list.rules.spec.cookie_matchers.name` property

Type: `"string"`. Optional.

Cookie Name. A case-sensitive cookie name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-0103313330202032-0333323222300211-3231010121102231-1022300331232331-1311002322210132-3210201002132031-2200001313310231-1132120002030213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--http_loadbalancer--reference--group-023.md#canonical-1112223030321202-2310131032330100-2302103332230131-0001200330203221-2212303302110122-3033121110232100-1130300130302113-3232330311033030)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present

<a id="canonical-1301310221031133-3232011221111032-2201210210121031-1210020202020322-1203132333021121-0032111003223002-1203222123220112-3111022203232100"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1011011121103030-1023321311310333-2210310200001020-2031223022321101-0211112103002331-0113300012011220-0213102312210012-3331322211331322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--http_loadbalancer--reference--group-023.md#canonical-1112223030321202-2310131032330100-2302103332230131-0001200330203221-2212303302110122-3033121110232100-1130300130302113-3232330311033030)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present

<a id="canonical-2122002112213001-1220221103323200-3311212230131022-3010330110220012-2300122233332000-1300013310123211-1000320321023032-1223110132111310"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2301120303000033-1120110232011011-0311033303003031-3013101323203110-3103223000001020-2012122203222232-3020031332212020-3132031221321231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.cookie_matchers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--http_loadbalancer--reference--group-023.md#canonical-1112223030321202-2310131032330100-2302103332230131-0001200330203221-2212303302110122-3033121110232100-1130300130302113-3232330311033030)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.item

<a id="canonical-0233310200121001-1010031310210020-1002330202312030-0112101200233223-2100130121231000-0033203200213032-0002230032123233-0002301313312010"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-0310230223303231-3123232210103201-2330121313033000-0133213022202002-1023022311100332-1313003031312333-2000231111203212-0203231223131111"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.cookie_matchers.item`

<a id="canonical-0211321132033130-1330031123321102-1102111320112213-3222023003021013-2333321132233100-3123320101220031-1100321100123210-3210203310010021"></a>

#### `policy_based_challenge.rule_list.rules.spec.cookie_matchers.item.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-3132130322300322-2211222320230323-0232220200002330-2132033130013202-3320003301322210-2122030300203032-1130233112103121-0133012122111020"></a>

<a id="canonical-1033301233101110-3223213110013002-2103113302301020-0212212223230303-1333202300223123-2130211133302030-2223313020102131-1210011020233233"></a>

#### `policy_based_challenge.rule_list.rules.spec.cookie_matchers.item.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-1033022011011132-0030003332130331-2231313220102023-1302201010311232-3121220103132032-2211210132311132-2332233333322033-0321003333232312"></a>

<a id="canonical-1022120132220033-2023302023023331-3011331233033112-2213102232021201-3312223211220020-2031330221131101-1130111201201020-0013021200310033"></a>

#### `policy_based_challenge.rule_list.rules.spec.cookie_matchers.item.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0202302323030031-0332332132022130-1301100023003233-2302201010011300-1300320313311232-3332303233323013-2021332220011213-0032313121021322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.disable_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.disable_challenge

<a id="canonical-2303302200301023-0230123210232312-3110113233311112-2211230302131132-3222201213233310-3211011001032212-2100113110000000-1021122133032231"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable challenge.

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
disable_challenge = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2231322321300130-2211003200213110-0230120220232310-0330031202332122-1202330333221203-2022232030123301-0132023201301310-1122331113003030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.domain_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.domain_matcher

<a id="canonical-1232323002333033-0112032300221013-1133320131331310-0103302201222000-3001320023012200-2112013212022221-0032101200212311-2212321012220203"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
domain_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0301200211320232-1120113123120033-3011002003133223-1311033112021233-0312010303020113-1002233031132023-3323002333313301-2001310032111021"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.domain_matcher`

<a id="canonical-0013211133032103-3232013113002302-3230113022221032-1132211331203001-3113023233023301-0103120223333213-1020212133232220-3010130102010320"></a>

#### `policy_based_challenge.rule_list.rules.spec.domain_matcher.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-2202110300322313-2013321002200231-1311231313202100-0132003012202131-1122113020331011-0300000223013202-3213010220210303-3213212130020003"></a>

<a id="canonical-2130000102131321-2203131013303102-2330213000003203-2121011030102202-2312310211311121-0013021303123231-0120023113312203-2011203233211112"></a>

#### `policy_based_challenge.rule_list.rules.spec.domain_matcher.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-2331103200120201-0301001021002322-0132233020221032-2203102002320022-3001110313102332-0311012113002330-0202002031220211-1023110131120121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge

<a id="canonical-2132331203310320-1222200200202033-0011330110111031-0312132201102202-0013320113013122-3102310210001133-1120212312030103-0100322013013022"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable captcha challenge.

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
enable_captcha_challenge = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2221112113201102-3022103132211101-1320220303332222-2303330112030301-1020203320020311-2100330302222210-1012001232303023-3130312322023023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge

<a id="canonical-1130202110233103-2131102003313220-3030212130333030-2001123211330000-1012230010300023-2231123011102202-3222120333303323-1012211201202020"></a>

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
enable_javascript_challenge = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1313201031011033-0203021101211130-2233102100210202-2231200332022130-0212312030111003-0030330021121330-0011012123121011-2011131311103111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.headers

<a id="canonical-3213121202022121-0002321330203123-0233123032313100-2330323312031012-3330313002033230-0330302002020002-2201311111210003-0302110231212101"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-1332133313121102-0030220013123001-3301211213312110-0012131033113322-0101323000121030-1313332032100231-0301213222200002-0223130120112132"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.headers`

- [check_not_present](resources--http_loadbalancer--reference--group-023.md#canonical-2323020220021030-1221122111120010-1000002310111301-3103132212310231-3222000202211323-2223022323033322-3200012012111322-2010121000223113): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-023.md#canonical-1211303211131310-1223121301221223-0331003023331300-3013201222033021-3213232122201131-1201322000101313-3231203021010120-2301213312132013): complete subsection reference.

<a id="canonical-1022001022301210-0301201121032130-1210031220100332-0233100233200020-0212020022312231-2111301031301200-3110000333023133-0130122100020212"></a>

<a id="canonical-1000021210212322-2221202301011212-3323301220003110-3023311102021310-2210033303112100-0210202222132233-3200121201230330-0101030031012210"></a>

#### `policy_based_challenge.rule_list.rules.spec.headers.invert_matcher` property

Type: `"bool"`. Optional.

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

- [item](resources--http_loadbalancer--reference--group-023.md#canonical-3000212303030230-2323101000310303-2133302302323120-1302012010122311-3230232013302223-2221311321310131-0313233101301123-0123020121000020): complete subsection reference.

<a id="canonical-0223111223131203-3012133001321010-0203303210331200-0332022302323220-2010233130030112-3211000312013102-1331130102132110-2202310222100323"></a>

<a id="canonical-2111211221322210-0023111231130122-3032203032030000-1012231210030220-0001011003332301-1112202233312322-1212302021033100-0030311222130323"></a>

#### `policy_based_challenge.rule_list.rules.spec.headers.name` property

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-2323020220021030-1221122111120010-1000002310111301-3103132212310231-3222000202211323-2223022323033322-3200012012111322-2010121000223113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.headers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.headers](resources--http_loadbalancer--reference--group-023.md#canonical-1313201031011033-0203021101211130-2233102100210202-2231200332022130-0212312030111003-0030330021121330-0011012123121011-2011131311103111)
- policy_based_challenge.rule_list.rules.spec.headers.check_not_present

<a id="canonical-0002300313103112-3301102110122300-3301202122012112-1033233202110120-0000213133220003-3202030031123301-2123031230210031-3211322121300300"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211303211131310-1223121301221223-0331003023331300-3013201222033021-3213232122201131-1201322000101313-3231203021010120-2301213312132013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.headers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.headers](resources--http_loadbalancer--reference--group-023.md#canonical-1313201031011033-0203021101211130-2233102100210202-2231200332022130-0212312030111003-0030330021121330-0011012123121011-2011131311103111)
- policy_based_challenge.rule_list.rules.spec.headers.check_present

<a id="canonical-1321330332113011-2101112110100033-3111113123330110-1121210322220332-3120030323110122-3002022122200332-3323113222112222-0331332222211122"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3000212303030230-2323101000310303-2133302302323120-1302012010122311-3230232013302223-2221311321310131-0313233101301123-0123020121000020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.headers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.headers](resources--http_loadbalancer--reference--group-023.md#canonical-1313201031011033-0203021101211130-2233102100210202-2231200332022130-0212312030111003-0030330021121330-0011012123121011-2011131311103111)
- policy_based_challenge.rule_list.rules.spec.headers.item

<a id="canonical-0111123330021001-3132230101323211-0013032210303001-1321303020130103-3312300320201001-0213013021220112-1003300220322122-0231003310002013"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-2310101323121001-0211110331100322-1312012100130231-2010100321320303-0110122012120201-2230301111233330-0330121003121120-2223212011222032"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.headers.item`

<a id="canonical-0333223002223333-1030033220212320-0230212212132010-0113330322320013-3211112120311110-2300121330023322-2213332311201021-3221110110100001"></a>

#### `policy_based_challenge.rule_list.rules.spec.headers.item.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-1121113112100010-1230120030331313-1122012133110102-1033030322321112-2332310033300322-1230323031111201-2123301113133002-2333321220010311"></a>

<a id="canonical-0303111321112133-0211133322230330-0103333121333212-0130011231113221-0012032233212201-3320300220220032-0332002102113230-1033021332202233"></a>

#### `policy_based_challenge.rule_list.rules.spec.headers.item.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-1211321000113030-1230333002310302-1211233023011231-3322213033113020-0231021023202210-0010303132011122-2021323311133122-2130123112012231"></a>

<a id="canonical-0300100100011322-3323230113201313-2120133313021112-1110201303331312-3013213111332323-3330312010121031-2321113003001333-2022321330121331"></a>

#### `policy_based_challenge.rule_list.rules.spec.headers.item.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2330011102111203-3200210100131120-2112200113232323-3120021212132232-3203312330132100-2321330223303101-1023321200013301-2303010303011323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.http_method` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.http_method

<a id="canonical-0320200311032022-0313303221002101-1020100111313001-1023123212203203-2012111201110121-2320030130233021-2202110322203113-3321111323232030"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
http_method {
  # Configure direct properties listed below.
}
```

<a id="canonical-3302310112031223-3010121020010212-1100133110333102-1033122102232203-0331203313322030-3102331103212303-0323320313002300-2112213220132102"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.http_method`

<a id="canonical-2322231000110332-1013032003311300-2233231232011320-0112133100311110-3121110121213321-3120203021201033-3301000023001232-3123311011303113"></a>

#### `policy_based_challenge.rule_list.rules.spec.http_method.invert_matcher` property

Type: `"bool"`. Optional.

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

<a id="canonical-0233332131131232-1200032030013311-3130032222102322-1300103210133013-2200101000121211-0020033303302232-3203123111323322-3001030103020221"></a>

<a id="canonical-1003231001311033-1302003133220021-1302010012003120-1033033010210322-0012131211210001-1112111100302113-0311301313311000-3133023001323010"></a>

#### `policy_based_challenge.rule_list.rules.spec.http_method.methods` property

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] List of methods values to
match against. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`,
\`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-3131232312332313-2331122113110111-0131010211100102-0310213132121112-3111200310110021-1010121212033013-3232311000121330-1220231122112110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.ip_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.ip_matcher

<a id="canonical-0023210031221301-0320301010030320-2003132110122223-3200103210200311-1321320111103111-2020021032132221-2030003211130130-1331022000121320"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("prefix_sets")}
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
ip_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-2311232001321213-0102202033130230-3310323102031032-3003233211221223-0200323302030220-3313202302231302-3012232110230012-0331311130301020"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.ip_matcher`

<a id="canonical-3132120210121311-3032132212210022-0301222301010110-3322331131000322-0302330031022333-1232013331133202-3003211233330032-2222122322110333"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_matcher.invert_matcher` property

Type: `"bool"`. Optional.

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

- [prefix_sets](resources--http_loadbalancer--reference--group-023.md#canonical-0212332012223100-0032200302201123-2123130003231102-1023020003022121-1133332212200021-0132021010113201-0202231131023203-2100111123101220): complete subsection reference.

<a id="canonical-0212332012223100-0032200302201123-2123130003231102-1023020003022121-1133332212200021-0132021010113201-0202231131023203-2100111123101220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.ip_matcher](resources--http_loadbalancer--reference--group-023.md#canonical-3131232312332313-2331122113110111-0131010211100102-0310213132121112-3111200310110021-1010121212033013-3232311000121330-1220231122112110)
- policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets

<a id="canonical-0330020020230210-0021231320002110-0010301220333003-3311213313022022-2332322121022201-0002110000300321-1310322230132333-1111103201211230"></a>

Type: `"object"`. list nested block, Optional.

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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
prefix_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-0333112213200013-2030313001330212-2131201001101331-0003130210230232-1210331010303300-3331103113222202-1130302102133133-2333201003310213"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets`

<a id="canonical-3013310222301300-1303103300311211-0110131313201031-2120010231023233-3133310030321131-0021102231212322-2220021133001110-0213221300022131"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets.kind` property

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

<a id="canonical-0133303333302231-0011331323122221-1330200113023302-2003111302011220-2032230022132102-2332320102202322-2130133210111130-0103311111203010"></a>

<a id="canonical-3011000033323231-3203002033123133-2333333110102202-0100011120123203-2303302203311201-3320301010111030-0030022313211123-0211002210031030"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets.name` property

Type: `"string"`. Optional.

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

<a id="canonical-1202123301203013-3302310302113032-3031301202212123-0201123331033010-0312211221203321-3022113130100123-0303321313302113-0010010011133123"></a>

<a id="canonical-3202112310212033-1300000123333020-2203313020100120-3313030332221011-2102312320221222-0133320202101300-2032013010023101-0130033210101023"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2003201002111230-0301232123032323-0220212331012131-1330230313032132-0121223032031130-1201120231210300-2001102010222222-2100311113012312"></a>

<a id="canonical-1232212210323111-2212011212002321-0122201331000220-1113113212002302-0303230213303312-1221023113222230-3211130123101331-0220312202300221"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets.tenant` property

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

<a id="canonical-0332300213330031-3130122101212030-1013222320002312-0313110030003131-3033103122010133-0011311223023230-2113101223322201-1122313320133003"></a>

<a id="canonical-3121303220012122-1101003131010301-3101210121313001-2100003011222222-1213031033110101-3113133110112221-0003122131302002-0002100133133021"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets.uid` property

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

<a id="canonical-1102303013213221-2101121232233030-3032103211220121-2031323223113213-2210001011111012-1001212110123010-0120200232130210-3120021232133212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.ip_prefix_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.ip_prefix_list

<a id="canonical-3210033200001103-2012311001110030-3021022230112021-3121211323010123-1101313300033203-1101232313320203-1022013101101330-1100220210102312"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ip_prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1113211320232312-2133312123123223-3201101113230213-3031131321020002-0023001232100112-1110330130202322-0310323211323313-1210232202120311"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.ip_prefix_list`

<a id="canonical-0023300312023321-1222112232000303-3312020333323100-3010130101213223-2102133130320000-0111122002211030-2021230230212113-0211012302331033"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_prefix_list.invert_match` property

Type: `"bool"`. Optional.

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

<a id="canonical-0133032110130220-2310121223020020-0123310010210101-3003212322113100-0332332333110222-3011012202111312-0122310132302330-3303201322113222"></a>

<a id="canonical-1112202021011323-1110001102022230-3111011223220212-0321122321203031-3130121012301020-0111313322102010-3302233200213011-0303120130212311"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_prefix_list.ip_prefixes` property

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

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

<a id="canonical-2110130302132222-1333323202323202-2101132211132333-3120123331020002-2321313212221101-0010000002320121-1033331331113211-1113310312332201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.path

<a id="canonical-1013203123330210-1012013200102220-0003230000130221-0032221333233303-0232201031022121-1113220202200323-1130210333213332-3201203112101200"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-2133220012331120-2122122233102121-3023302110001311-2010200200333100-3010013210303211-1033211101100330-2112133230221113-0323113121313221"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.path`

<a id="canonical-2322213223333133-2230111223213113-1101030332232301-2312111231210202-1001000210222102-1102333303220203-2232023020012330-1302120230032331"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.encoded_path_matcher` property

Type: `"bool"`. Optional.

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

<a id="canonical-2032022121322200-1132302300113032-1330302213013232-1212021303311003-3221103100300013-2321302312322201-0023000111101321-3101321303201233"></a>

<a id="canonical-2000210211012110-2111230122302121-1132332111210312-3033210212030213-1231121323202202-0111202231121310-3312310232022322-2112011232022002"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact path values to match the input HTTP path against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-0103011303301111-0020211031023001-0230102012233333-2211311222232032-2133103012030033-1300302330111333-3022200121020112-0013111211022231"></a>

<a id="canonical-1200231313123222-3013023230022331-2001023133233103-3331031011132121-2221130302320213-0223213323312022-0312332222113303-2221203110133210"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.invert_matcher` property

Type: `"bool"`. Optional.

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

<a id="canonical-1303031331230310-2001233201022112-0202233230120120-3101303133200212-1032201333023321-2120002230031330-1130220023203132-0023203311310311"></a>

<a id="canonical-1211100300033300-3002132113221111-3303001011132323-2221331013122012-2222002332302033-2010101213123301-2021003012322121-1020110211230030"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.prefix_values` property

Type: `["list", "string"]`. Optional.

A list of path prefix values to match the input HTTP path against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-1130102211330003-2222230333300210-3302213111012320-1003323322312130-0222303132322103-0331231303320230-1100120303100111-2321113312020121"></a>

<a id="canonical-3303123032221013-2211030022200201-2200303012133233-0012021213001013-1000310122000311-0132021110212202-1311200002130121-2201302013311111"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input HTTP path against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-0221232012101112-1313011111020110-2231103332101331-0310313113013220-2000010303033011-3221201010321131-3300132103331132-3222212030302220"></a>

<a id="canonical-3300131311312331-3212331322233102-0312222203003021-2121121222203312-3212013033230102-3000310000203310-2132112112002121-3120220213030102"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.suffix_values` property

Type: `["list", "string"]`. Optional.

A list of path suffix values to match the input HTTP path against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-1312322031021220-0320122303112122-0320230103103300-3102031321330000-3110132213132230-0310331232221303-2132301001233230-0311302031231000"></a>

<a id="canonical-2032202003321212-3232011202101102-3001200023110003-2022212211313232-2230213300331013-1011333303202131-1031320133323102-0130112100113010"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3301321013233130-3120220022102302-2010030102010301-1203130111033000-2101212222020102-2320031102212001-0133011132121102-0210000220320303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.query_params` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.query_params

<a id="canonical-2203120210113211-2112000210103232-3102220020002112-2020330112232021-2101223003312230-0312303201312023-3031122000103101-3013000312102301"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("key"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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

Terraform syntax:

```terraform
query_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-1102113321203001-1331211222203230-3223232032320001-3120100331313010-2312011112310032-3003003231020022-3212331203001023-3001110023212320"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.query_params`

- [check_not_present](resources--http_loadbalancer--reference--group-023.md#canonical-3103300010123120-3200220011010132-3033202200030111-2303133000000131-0311023323212320-0013122300031122-0222001002323220-2332111330223102): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-023.md#canonical-2112003030120300-3231323312231200-0322203333323223-3322103103222131-3012202202123232-0312103310010201-0032132013131023-0001222211300312): complete subsection reference.

<a id="canonical-3021231120131001-1310320220111012-2023102213231133-1100132023023123-1332132102033100-1212122232202300-1123111031233320-3332030310102010"></a>

<a id="canonical-2010301201110232-3113120110021303-0302321233223032-0101020230212321-1122200111012032-1122300121102333-1103321331000102-2210310213003031"></a>

#### `policy_based_challenge.rule_list.rules.spec.query_params.invert_matcher` property

Type: `"bool"`. Optional.

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

- [item](resources--http_loadbalancer--reference--group-023.md#canonical-3221303202213312-3211202232332112-3101121302023021-3001021033200203-3301011202001122-1221212210100122-1300213130231002-3330311103023100): complete subsection reference.

<a id="canonical-2213321300002311-2112331211233132-3210113330113321-0101011031122232-3022013312103303-3103031011021011-2121230131303312-3223130213312202"></a>

<a id="canonical-3210231322322222-1003000110123001-2301202033312322-2120033133231212-3230231210230122-1123313001230332-0113220220211021-2001311023131023"></a>

#### `policy_based_challenge.rule_list.rules.spec.query_params.key` property

Type: `"string"`. Optional.

A case-sensitive HTTP query parameter name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3103300010123120-3200220011010132-3033202200030111-2303133000000131-0311023323212320-0013122300031122-0222001002323220-2332111330223102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.query_params.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.query_params](resources--http_loadbalancer--reference--group-023.md#canonical-3301321013233130-3120220022102302-2010030102010301-1203130111033000-2101212222020102-2320031102212001-0133011132121102-0210000220320303)
- policy_based_challenge.rule_list.rules.spec.query_params.check_not_present

<a id="canonical-1101303100132312-0332013330032313-0322313012033222-1203311103123331-3301300321311102-1310132011213220-3200210002113130-1333201120100332"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2112003030120300-3231323312231200-0322203333323223-3322103103222131-3012202202123232-0312103310010201-0032132013131023-0001222211300312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.query_params.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.query_params](resources--http_loadbalancer--reference--group-023.md#canonical-3301321013233130-3120220022102302-2010030102010301-1203130111033000-2101212222020102-2320031102212001-0133011132121102-0210000220320303)
- policy_based_challenge.rule_list.rules.spec.query_params.check_present

<a id="canonical-2012232230312011-2000002102012122-2312132210110121-2112113203321232-0200231201101010-2323330333301321-2201301222133133-0023212023130201"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3221303202213312-3211202232332112-3101121302023021-3001021033200203-3301011202001122-1221212210100122-1300213130231002-3330311103023100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.query_params.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.query_params](resources--http_loadbalancer--reference--group-023.md#canonical-3301321013233130-3120220022102302-2010030102010301-1203130111033000-2101212222020102-2320031102212001-0133011132121102-0210000220320303)
- policy_based_challenge.rule_list.rules.spec.query_params.item

<a id="canonical-2021111131031122-1233033121101103-0232020210133021-3231312303321331-0332101000121310-3203301010032330-3000011302111023-0310102312313100"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-0031103030220132-2331300003310232-1312013111133223-3202032033210111-0332212132212101-0031101013312301-0103000011203012-2310311001030032"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.query_params.item`

<a id="canonical-2000211123332030-2312331230233130-0231333310131100-0233231231311302-3300203021110132-0232132211223001-0021331100330232-0230102120110212"></a>

#### `policy_based_challenge.rule_list.rules.spec.query_params.item.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-2231211210223302-3221220321110013-2222131131220002-0002300123031321-2230211202033333-0012031032022033-3013221012101301-3220312010312122"></a>

<a id="canonical-1131201012300113-0133232213202002-0032112122133311-3121111212323221-1300310103303020-2221003202333120-1010132011321113-1100223233232103"></a>

#### `policy_based_challenge.rule_list.rules.spec.query_params.item.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-1130330003220001-2130120332033120-0302103320323222-2323001321301111-3211023222033022-2030212331202010-2222121312012321-0133233211000333"></a>

<a id="canonical-2101011333123332-2023222030321302-3332230321032220-3202211212221111-1333222111030301-0302323001102300-2002021202031201-1310111211201011"></a>

#### `policy_based_challenge.rule_list.rules.spec.query_params.item.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1131033020202020-3130333330012030-3002001123121030-1323223321202032-1221131132031001-3323021113022011-1203131303033130-3010031021310313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher

<a id="canonical-0023121113310101-3000100110211120-0321313013123100-2223023123312331-0021220321232200-0320221203123231-3231231132101211-1112110300100212"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
tls_fingerprint_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-3013031323323213-3313121012212012-0112303321001030-1100023031023311-0210001200203111-3132002132022111-2010033301222100-0102210102133120"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher`

<a id="canonical-1021111321203203-3301132133030002-2010322033311100-1303332302003321-3213112020223121-0220221202133210-0303002333032202-1302001231301203"></a>

#### `policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher.classes` property

Type: `["list", "string"]`. Optional.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Additional upstream details:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-2200130213322322-1202121203133021-2311320332102331-2320112203331011-1132223232220312-1301022313313100-0322311023001212-1332110222103030"></a>

<a id="canonical-2010021303323201-3021331310213001-0121022023311100-3000002131303021-3213113020311211-3210100120233332-1131202302322303-1301020103220021"></a>

#### `policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-1201130212030203-2312220320013322-0320311001231222-2232020111320120-1110313223210301-3210113301312033-2121201200222023-0112222120120130"></a>

<a id="canonical-0112323210101301-1331121030333023-3033301212222231-2103032032321022-0110210132003200-2301130123200220-0111301112033112-2133233202130310"></a>

#### `policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher.excluded_values` property

Type: `["list", "string"]`. Optional.

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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

<a id="canonical-0332322110110101-3123223013302023-0321200302320201-2312111021001220-0032301331120201-2310003212230313-2113130213200312-2022222232311001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.temporary_user_blocking` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- policy_based_challenge.temporary_user_blocking

<a id="canonical-1213330020020110-0201003231121130-0322032122200333-1222232301221110-1231122130323233-0213300101032113-3211212012313013-1313011303000120"></a>

Type: `"object"`. single nested block, Optional.

Specifies configuration for temporary user blocking resulting from user behavior analysis.

When Malicious User Mitigation is enabled from service policy rules, users' accessing the
application will be analyzed for malicious activity and the configured mitigation actions will be
taken on identified malicious users. These mitigation actions include setting up temporary blocking
on that user. This configuration specifies settings on how that blocking should be done by the
loadbalancer.

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
temporary_user_blocking {
  # Configure direct properties listed below.
}
```

<a id="canonical-1001233223103102-2321002300333321-1010002332023330-2130131120101022-3320211110300131-0000103020320331-1111331221322023-3112130013330213"></a>

### Direct properties for `policy_based_challenge.temporary_user_blocking`

<a id="canonical-3120200303022222-1320212320330112-2211020312113103-1012301200100332-0131032211233132-2100000001010121-3122131220320112-0210110121030023"></a>

#### `policy_based_challenge.temporary_user_blocking.custom_page` property

Type: `"string"`. Optional.

Custom message is of type . Currently supported URL schemes is . For scheme, message needs to be
encoded in base64 format. You can specify this message as base64 encoded plain text message e.g.
'Blocked.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '&lt;p&gt;
Blocked..

Additional upstream details:

Custom message is of type \`uri\_ref\`. Currently supported URL schemes is \`string:///\`. For
\`string:///\` scheme, message needs to be encoded in base64 format. You can specify this message as
base64 encoded plain text message e.g. "Blocked.." or it can be HTML paragraph or a body string
encoded as base64 string E.g. "&lt;p&gt; Blocked &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- protected_cookies

<a id="canonical-0012220312310111-2303132013122322-1110302013133310-3121322300132331-3211323303223113-1121220330010001-3331130300302213-1211121030222300"></a>

Type: `"object"`. list nested block, Optional.

Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering
Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering
Protection, enabling a web app firewall (WAF) is a prerequisite. The configured mode of WAF
(monitoring or blocking) will be enforced on the request when cookie tampering is identified. Note:
We recommend enabling Secure and HttpOnly attributes along with cookie tampering protection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("add_httponly",
    "ignore_httponly"),
  validators.ConflictingListObjectAttributes("add_secure",
    "ignore_secure"),
  validators.ConflictingListObjectAttributes("disable_tampering_protection",
    "enable_tampering_protection"),
  validators.ConflictingListObjectAttributes("ignore_max_age",
    "max_age_value"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_lax"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("samesite_none",
    "samesite_strict")}
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
protected_cookies {
  # Configure direct properties listed below.
}
```

<a id="canonical-2313133323322130-1111101311101100-1311201113313021-2211332110031211-0312113310011223-1213211301102330-0220030111133311-1002023200222010"></a>

### Direct properties for `protected_cookies`

- [add_httponly](resources--http_loadbalancer--reference--group-024.md#canonical-3300300313122023-0131101133121122-0111302321111221-1320231013003202-0133111311232132-2000031230020032-2303313311132310-0320021312132211): complete subsection reference.

- [add_secure](resources--http_loadbalancer--reference--group-024.md#canonical-2231033302001101-0132231313110213-2031112210223222-3333133100110021-1233300112010200-2023001101013012-2002112312201120-1102200221332013): complete subsection reference.

- [disable_tampering_protection](resources--http_loadbalancer--reference--group-024.md#canonical-3213201012220223-2211332120011102-3023133313022103-0320131213311121-2310000233223030-2021102301111110-1000311211111311-1220303110323223): complete subsection reference.

- [enable_tampering_protection](resources--http_loadbalancer--reference--group-024.md#canonical-2120321222300232-3201213120231321-3002320121130331-3033120022323313-2001001201032333-3213123201002332-1312021122123131-1103100231332211): complete subsection reference.

- [ignore_httponly](resources--http_loadbalancer--reference--group-024.md#canonical-0203212110203021-2200220231120210-2230311113031102-3221022331201212-2022331233333222-1023330101000213-2212231131222231-2200000112100101): complete subsection reference.

- [ignore_max_age](resources--http_loadbalancer--reference--group-024.md#canonical-0113232103232213-1220110302131233-2220201300012301-1223330101120121-3102023230122321-2021001312001331-3212303132300032-1121003002111111): complete subsection reference.

- [ignore_samesite](resources--http_loadbalancer--reference--group-024.md#canonical-1222213222130010-1030032131320230-1313212131202032-0230100032002121-2200132212221021-2032011033232331-1222130131101301-3033111203030322): complete subsection reference.

- [ignore_secure](resources--http_loadbalancer--reference--group-024.md#canonical-2003122021101102-2103020000110303-1022023202020222-2212321100200322-1113103131312223-1311113032323233-2222101320213033-3231231121303232): complete subsection reference.

<a id="canonical-1323120312300302-3320100011310320-0011202033330112-3301121130120012-0301301111231001-0102330220212010-3133000111121113-1020102133110003"></a>

<a id="canonical-3320233203100101-0123331113233032-2013002133222330-3323310301303231-1011212011011202-0333210110322202-3220212211231022-3013220001211032"></a>

#### `protected_cookies.max_age_value` property

Type: `"number"`. Optional.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(34560000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 34560000,
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
    "ves.io.schema.rules.uint32.lte": "34560000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  }
}
```

<a id="canonical-1321133322311221-1000131311233203-0132313321131303-0222222031301101-0233031223010201-0331020201130000-1203312001333210-3333222232301112"></a>
