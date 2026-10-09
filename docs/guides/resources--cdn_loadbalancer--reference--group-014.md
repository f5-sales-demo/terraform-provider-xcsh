---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-0003030013130330-2331113132100113-3133013032111101-2330111202210322-2212222032222202-2223222120111202-1220132122303010-0212112303222130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [policy_based_challenge.rule_list.rules.spec.asn_matcher](resources--cdn_loadbalancer--reference--group-013.md#canonical-0003102132023010-3231100012022023-0021011213211001-3321012032332023-1221003122112110-0003331033100121-1032002020102300-1103330212100101)
- policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets

<a id="canonical-1121012211122000-0233332002121213-0002032011000131-0000200223132111-1320121001123300-1022121111103022-3122320310323001-2230201100313102"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
asn_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-1310022213331313-2011122223310302-2232122301322101-0103130111023223-3100030032011002-0221301110210303-1202222203302113-2312313033213123"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets`

<a id="canonical-0130033203003331-3111032011230332-3131330021223203-0033022001021333-3303133212212030-2210032220131022-2230000021121330-2330011121312003"></a>

#### `policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets.kind` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3111102122210211-2232221210102311-1000121120133211-0003331003211230-1112003133203131-0230320132132231-2222331123201020-0010212113301201"></a>

<a id="canonical-2232230123210323-1020122020203013-2321132200203112-3120031010113113-3020011213123133-0113102213031110-0120313301101011-2011321323312113"></a>

#### `policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0221022230212122-0212202203221012-0313203031013022-0030232101202313-2113020010113301-3111321100023320-2122203012200201-2233232213322023"></a>

<a id="canonical-1131312100111313-1113310200110331-3203023310122003-0101133031223013-2201021023000213-0203210210313023-2211331332033013-2220300310013203"></a>

#### `policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3303010112101120-1310121213122202-0112203103000330-0303221120223131-1323322103220303-1220121131123112-2110301101032213-3311311103001202"></a>

<a id="canonical-0111213013131020-2313030132312131-3123212232113000-0330223320111030-1230112013032311-2220100201011013-2101311313210020-2131303330111001"></a>

#### `policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0313111031022023-3133300101321001-1023213302101123-2111230011120001-2003302112210110-2300123113330210-1021110133130230-2220102203231202"></a>

<a id="canonical-0013030203111202-3100000030133300-1013332120100122-0020320322103000-1310223332133222-2223013030002201-2203220301011123-1102321101010130"></a>

#### `policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets.uid` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0020310012101102-3000230032112032-0132110110013312-1020032331321312-1203133212331130-0123320022003020-2120221312120320-3221212002033010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.body_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.body_matcher

<a id="canonical-1031030331011012-1102322013022301-0320330033223032-3222002223112121-1110032020123013-2002103323103102-1121113220132031-3212122032220032"></a>

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

<a id="canonical-3033012010220110-0102022132332000-1310301203021312-1210120302311233-1212202111032001-0331033203033023-1320213130030313-0022002310300233"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.body_matcher`

<a id="canonical-3032223032313111-0003202122300012-1032323312230310-2233032320332000-1013333012011120-2333201101032330-3311302320033122-0301230023200230"></a>

#### `policy_based_challenge.rule_list.rules.spec.body_matcher.exact_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0232300210032121-3313330332331310-2023232113023001-0020103210233310-0021120132021311-2323100233023300-0100001133202210-0011210121023033"></a>

<a id="canonical-2000103103200031-1300113131111032-3311111033012002-2103121320311201-2202303010313300-0132311102233020-0310322211323303-0000312122223212"></a>

#### `policy_based_challenge.rule_list.rules.spec.body_matcher.regex_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1020313011221122-3001121233212111-2123133013232020-2201033312102323-0021003330200133-1221200013132220-2131331333331120-3132322030100123"></a>

<a id="canonical-1132322002011312-2321133113033113-2320233000020000-1311232010323002-1223332200103200-2110011233111100-0231201112312033-3111010011221232"></a>

#### `policy_based_challenge.rule_list.rules.spec.body_matcher.transformers` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3230133003213133-3011200302131310-1110312033300333-3311222103011221-1231200100131103-1131223213322203-3302022113011000-2001203101232132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.client_selector` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.client_selector

<a id="canonical-2122013012232301-1120001333102030-1112313122033120-2032321322212210-1021113322221223-3311303323320002-2313032033022113-1233321132211320"></a>

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

<a id="canonical-2331302310233132-3233330030031010-1222331332301310-1121203221002200-1133130221321110-0020301332030311-0221203212121201-2211200130013130"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.client_selector`

<a id="canonical-0300021310102112-0233312223213331-0032223231303123-0030332203312132-1021231203021020-0133330113320333-1010033031230032-3232131303122312"></a>

#### `policy_based_challenge.rule_list.rules.spec.client_selector.expressions` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3313200200332201-1200210121121211-3220003300020012-2002110001123113-3233122133221310-0212323322101310-1333012331212122-2002300120333130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.cookie_matchers` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers

<a id="canonical-3122323331000232-0332310100303121-3202233023302312-1303320110303023-1131232311302033-3332313233323332-2030213320001300-2000102020032112"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0331122130032110-0000030232232002-2011201021123203-0000330331011311-2200022321103311-2012031213112333-1012012312033213-2100100310123113"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.cookie_matchers`

- [check_not_present](resources--cdn_loadbalancer--reference--group-014.md#canonical-0200311131313000-1330113300332313-0201331201321232-3321212100022002-0021011312320202-3303313033012313-2131223133333032-3010121011220301): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-014.md#canonical-0120122022030210-3120003232102010-1000000302001001-0102210101030320-3211221101103330-3010111213033233-0103032021303303-2112200021102030): complete subsection reference.

<a id="canonical-1323323221320221-2312213200230130-3332103333212320-2133300320000100-2121000322121322-2130133112231122-2103101021333221-1023303213123320"></a>

<a id="canonical-2230011003230322-3320220120102021-3023002200230202-3032113322230132-1313203223311011-1230200330132022-0131110122003313-0030231010302130"></a>

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

- [item](resources--cdn_loadbalancer--reference--group-014.md#canonical-0103003011323300-2200320200120210-1133213110133312-3102111212013223-0302030033013221-3011211222220211-1023030113323131-2133301300232302): complete subsection reference.

<a id="canonical-0121113003130233-0302303020330202-3311322320320211-1021332020200131-0132033110022301-0031031021301223-3010121221323300-2221203130311023"></a>

<a id="canonical-3020320111302323-1113312212302113-0132113321322033-1122111113003333-0021113313133133-1301131211110223-3113001223231221-3003000101202220"></a>

#### `policy_based_challenge.rule_list.rules.spec.cookie_matchers.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0200311131313000-1330113300332313-0201331201321232-3321212100022002-0021011312320202-3303313033012313-2131223133333032-3010121011220301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--cdn_loadbalancer--reference--group-014.md#canonical-3313200200332201-1200210121121211-3220003300020012-2002110001123113-3233122133221310-0212323322101310-1333012331212122-2002300120333130)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present

<a id="canonical-1320133103313133-1131123200213321-1301202313302323-3011130121123322-2203330102313202-2303121010333200-3000300130311331-1121213200021311"></a>

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

<a id="canonical-0120122022030210-3120003232102010-1000000302001001-0102210101030320-3211221101103330-3010111213033233-0103032021303303-2112200021102030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--cdn_loadbalancer--reference--group-014.md#canonical-3313200200332201-1200210121121211-3220003300020012-2002110001123113-3233122133221310-0212323322101310-1333012331212122-2002300120333130)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present

<a id="canonical-1201030003310031-1010232212330123-0201032133333021-1223302331213330-0223223023302132-1202022231212223-0130020211131130-1301132100322212"></a>

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

<a id="canonical-0103003011323300-2200320200120210-1133213110133312-3102111212013223-0302030033013221-3011211222220211-1023030113323131-2133301300232302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.cookie_matchers.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--cdn_loadbalancer--reference--group-014.md#canonical-3313200200332201-1200210121121211-3220003300020012-2002110001123113-3233122133221310-0212323322101310-1333012331212122-2002300120333130)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.item

<a id="canonical-0213100232112202-3011000033013312-3221010332113021-2313001103301111-0220300212002312-2220022131111312-2201311333031003-2111213210100220"></a>

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

<a id="canonical-1133121030121121-1233113000102130-1230111003321120-0003313122313323-1231300302301333-2312200303111022-1100303300023303-0002302222133121"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.cookie_matchers.item`

<a id="canonical-3301131023231003-1233013222312333-2302330232313233-0222003221221201-2312003322033012-2102232100033303-1012012110130330-1310231101210111"></a>

#### `policy_based_challenge.rule_list.rules.spec.cookie_matchers.item.exact_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0321230213121320-1120031333200120-0001012112321312-1131301310031100-3221221123123102-3301210312123033-1010332233003110-3211312022210221"></a>

<a id="canonical-0123013003203003-3311203003003212-3030122201312220-3231110120201332-1333113112312022-3132200222233333-3201110231220210-0232011323221110"></a>

#### `policy_based_challenge.rule_list.rules.spec.cookie_matchers.item.regex_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1313010213003200-1000022231313213-3122321210302221-3323301123122122-1300222331110003-2013222003133302-2320121000231113-2123220110030212"></a>

<a id="canonical-2210203331310200-1120122311001212-2103232311000311-0032232323120131-2330113013003321-3303031020230200-0311111311312311-1203310233100202"></a>

#### `policy_based_challenge.rule_list.rules.spec.cookie_matchers.item.transformers` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0032332002211123-1202012303111223-1300130212323212-3131233302210131-2031312311130100-1303223302131323-2102102310301121-2010003203213221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.disable_challenge` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.disable_challenge

<a id="canonical-0303000332313321-3320212331332232-0201122020300120-3211123012103021-2230320023003122-3012011010311201-1202213231130123-1111230210010233"></a>

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

<a id="canonical-1031330331223002-1133112231311120-0010033010221011-3200232312203013-1321321213331110-3211120301031200-0232002103001122-1233011032121333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.domain_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.domain_matcher

<a id="canonical-1101223320002201-0121100120130032-0022121223222023-0023113323133231-0003012013230310-0000000111121102-1003101103121320-3331122120213203"></a>

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

<a id="canonical-1023000330232230-0201313302211200-2212013010232033-1111101223323223-2132320000121222-1331010030211202-3010313322100131-3223023103121220"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.domain_matcher`

<a id="canonical-1111210100301113-3220313022122313-2120223223203132-3112330100312320-3330311203302110-0022332033121111-1030033331223332-0320202103132300"></a>

#### `policy_based_challenge.rule_list.rules.spec.domain_matcher.exact_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0323030030212232-3222300020331111-1301133310102021-2210121113202013-3031213322013100-2302011332210303-3121301103102231-3121213032311321"></a>

<a id="canonical-1002310121332002-2023222002232123-3001322232333001-1033010120033213-0312121003003033-2222320201222020-0212211131111113-1121110111202023"></a>

#### `policy_based_challenge.rule_list.rules.spec.domain_matcher.regex_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3103311303303301-3120122203132131-2113220010310230-3210210112101313-1321321213233223-2130133230221223-3223033330123030-1202010323003320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge

<a id="canonical-3232011221223112-1311003011100021-3023311130120120-2213322030301332-1212113021121301-2013230322112210-0010111131321010-2023112203220010"></a>

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

<a id="canonical-1101321300022020-2012121100121213-1300112332120133-3301000323233111-0133300213322003-1130203021233202-1031230021202010-0002003123331010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge

<a id="canonical-3130131113200201-3012320302303122-2200113101002110-3313132120231333-0021232020331303-1001222001331212-1132212103130213-0123220301123033"></a>

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

<a id="canonical-3330200010012232-0013311112232203-1032330210000200-2222122211223030-3322102023003230-0100300103120211-1113231232323333-0222103303312020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.headers` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.headers

<a id="canonical-0303223210112332-0020033010231131-2023130000030012-2131023021323100-3210100130123301-3101212200200310-1111032020212301-1000313320033001"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1023030213100012-0033330020330010-1021112131001313-1220330303002220-1103322331300131-3123313323231133-3211303311201032-0211302300320011"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.headers`

- [check_not_present](resources--cdn_loadbalancer--reference--group-014.md#canonical-2133232230322132-2013003023230313-2032312011333330-0323031212212313-1202223133002302-0130302300200130-1131302022302301-3222132301213310): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-014.md#canonical-2323233012323120-1101203011222330-0330221010331120-2003030011001211-2233323300223022-3110201200320233-0313322133232210-3320222232201200): complete subsection reference.

<a id="canonical-2123213032011220-3220320133331303-1310101330201121-1132310231132121-2003133133123112-3033013130021133-2002213233300002-1021332113001210"></a>

<a id="canonical-1201302020312300-1131100003313122-3222103201130203-2021223030103130-0012221121323023-0210220110011230-2231223230201202-3322232011330232"></a>

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

- [item](resources--cdn_loadbalancer--reference--group-014.md#canonical-0020010213313231-2230213012100103-3201102112031033-0120000120332332-2021003020323031-1222322303113232-1122212310101013-1031120222103022): complete subsection reference.

<a id="canonical-0121003112331220-1112103201302231-3010332210213111-3213100130103033-3230202230103332-2322113131302331-0330122313012030-3031132003310131"></a>

<a id="canonical-1012013332022101-3102113023112032-0301331333233331-2203100012002203-0232111111100110-0131212311311011-1002000113232332-3311331132022022"></a>

#### `policy_based_challenge.rule_list.rules.spec.headers.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2133232230322132-2013003023230313-2032312011333330-0323031212212313-1202223133002302-0130302300200130-1131302022302301-3222132301213310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.headers.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [policy_based_challenge.rule_list.rules.spec.headers](resources--cdn_loadbalancer--reference--group-014.md#canonical-3330200010012232-0013311112232203-1032330210000200-2222122211223030-3322102023003230-0100300103120211-1113231232323333-0222103303312020)
- policy_based_challenge.rule_list.rules.spec.headers.check_not_present

<a id="canonical-1203002130210320-0213000302220210-2010331232221233-1302222202222010-1021312133111030-3010123023110301-0001113203111132-0322032123213320"></a>

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

<a id="canonical-2323233012323120-1101203011222330-0330221010331120-2003030011001211-2233323300223022-3110201200320233-0313322133232210-3320222232201200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.headers.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [policy_based_challenge.rule_list.rules.spec.headers](resources--cdn_loadbalancer--reference--group-014.md#canonical-3330200010012232-0013311112232203-1032330210000200-2222122211223030-3322102023003230-0100300103120211-1113231232323333-0222103303312020)
- policy_based_challenge.rule_list.rules.spec.headers.check_present

<a id="canonical-3220201211002330-0011220302023230-2223222222122032-3112101223201233-3213323321232000-1113132030031200-0333332010002120-1312031011003230"></a>

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

<a id="canonical-0020010213313231-2230213012100103-3201102112031033-0120000120332332-2021003020323031-1222322303113232-1122212310101013-1031120222103022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.headers.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [policy_based_challenge.rule_list.rules.spec.headers](resources--cdn_loadbalancer--reference--group-014.md#canonical-3330200010012232-0013311112232203-1032330210000200-2222122211223030-3322102023003230-0100300103120211-1113231232323333-0222103303312020)
- policy_based_challenge.rule_list.rules.spec.headers.item

<a id="canonical-1001121230332002-2000120030103030-3112012121222110-2123030320333113-1010121333133320-0030303121301013-3312322022232023-1113320001011110"></a>

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

<a id="canonical-3313001122132311-1313221030102131-1331021130312111-1000130003201000-2112220112113101-3020133322102000-3312321202102003-1221022131010321"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.headers.item`

<a id="canonical-1130213221202233-0303321311221210-1010001122233012-1130121100320201-2203230323311320-0000203012311003-3222101133233020-3120013313012002"></a>

#### `policy_based_challenge.rule_list.rules.spec.headers.item.exact_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3000133123131333-0311133232321223-1322223102031322-2100103302030011-3030320010210132-2022032211212330-2121110322000112-1101210031301011"></a>

<a id="canonical-2132011333011110-0010113130331022-0321220021130010-1000100011010211-2203310312210132-1223230123320213-1123102013212323-0011323133323232"></a>

#### `policy_based_challenge.rule_list.rules.spec.headers.item.regex_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0230101033312311-2210010000103120-1100201110021122-3033123232303112-3210120111021200-1303210202220213-1312001002023330-0010203130100300"></a>

<a id="canonical-2120212130002311-2131031022023112-0211311223311201-1122301101101102-1030000220332320-2120112201332221-2131130332303100-0112120332133210"></a>

#### `policy_based_challenge.rule_list.rules.spec.headers.item.transformers` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2312131312221300-0230011100003021-0021131212211331-0123333322230033-1210310032303322-1320033320200030-2221332222001320-2330310311021003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.http_method` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.http_method

<a id="canonical-3320212121102313-2230031103220110-2211201003313212-1320210223011323-0001012002020232-2200111211312211-0132203132213100-0323223001122232"></a>

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

<a id="canonical-2000011311220310-3010010100122200-2213332310122103-3010031312312011-2100131103221231-1332113130102301-2020332223002021-0111321302112230"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.http_method`

<a id="canonical-2203232232111313-1221212000123010-3323200113023331-3200321032220321-2201110100201303-2222301123120121-0322202000303033-1213031011010112"></a>

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

<a id="canonical-0010102333021022-1312312211110233-2301231000333301-2033312002002101-2013303130302133-0113220022301202-1011222320210021-1031210211321100"></a>

<a id="canonical-2102313213231110-2313103231123113-3013011013323301-2321011032012130-3210300012101202-0221322112212332-3113102010021201-0221103322113300"></a>

#### `policy_based_challenge.rule_list.rules.spec.http_method.methods` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1221203300113022-2132333200200013-1003100021200202-1120013111223113-1232003120333102-0311332302111030-2000323030332330-3022032101132020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.ip_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.ip_matcher

<a id="canonical-2322333132221020-1232232220213220-1232221023030310-1211311021103233-0330010333012110-3211320002300323-0332002303231310-2132123021123110"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ip_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-1021102121333021-3330120312111330-2130202313113013-3311032032313032-0300231230332212-0211103313202313-2122101232301131-3201323100200121"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.ip_matcher`

<a id="canonical-2010112032010113-1330030033132030-1300312212033120-3232102120323323-1030232131012030-3232133302002001-3223122332103033-1021123300301313"></a>

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

- [prefix_sets](resources--cdn_loadbalancer--reference--group-014.md#canonical-2311012020023332-0201330100311030-3233112032312331-1222302001110020-1110031220112131-3112030003122113-1012032302112032-0312221130221132): complete subsection reference.

<a id="canonical-2311012020023332-0201330100311030-3233112032312331-1222302001110020-1110031220112131-3112030003122113-1012032302112032-0312221130221132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [policy_based_challenge.rule_list.rules.spec.ip_matcher](resources--cdn_loadbalancer--reference--group-014.md#canonical-1221203300113022-2132333200200013-1003100021200202-1120013111223113-1232003120333102-0311332302111030-2000323030332330-3022032101132020)
- policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets

<a id="canonical-2123201001123223-1201000120120331-1030021130323111-1100313200303112-3232210032301101-2300133322301003-0121222303211323-3110330220122200"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0231213013100021-3330023332021211-1203210100321020-2301231221223102-1300222021131000-0021032223202103-2222100233120220-2233021201321132"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets`

<a id="canonical-0232320111222030-1302203003330311-2331123202101230-3112132310130210-3232231202010102-1230122210230300-0333210333010111-2020101213212102"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3121132013100330-1220122031222003-2232213001333031-3213320213302331-3201203200212003-0102301321202333-3103310322133321-1101233112000301"></a>

<a id="canonical-3030233132133021-3021023220220203-3011232211222010-3323200132210133-2131322012101122-1131032023200101-2202020311232130-1111321133331001"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0011012202233302-3220331121000213-0120233003300210-1321200212322232-3011023221031233-1113012010233021-3101321232222030-2300331213123313"></a>

<a id="canonical-1010333022100323-2112310011033310-0310003013021231-2232001110001023-2022113213122011-3201203230003321-3220023213131300-2323312312220022"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0312000323213231-0111012032202000-3033211033210133-3013230131121200-3213010021120312-1222022132301313-3013202302313023-1002131033012000"></a>

<a id="canonical-1311200113013122-1023220122111011-2320322101211313-0230113102301313-3002202330223003-0031223222003210-0223232010101312-1101231213030212"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2300122200321331-3300303111213333-1112202332201133-3110022331232021-1203220200202301-1021203313300132-3102321003101113-3310002112332100"></a>

<a id="canonical-3011210021020201-3102221203112112-2301022223132323-2001130002011302-3302203331211011-0322131230100110-1021100133222121-3110321221201210"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1012033320010022-2222231020323103-2321000021201310-3031023311331022-0012002023131321-2230011300323100-3130312013222320-2213121223200210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.ip_prefix_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.ip_prefix_list

<a id="canonical-1123000020003201-1122230222010213-0301000031013301-3320012112200222-1230232102003203-2110011211121312-3232200333231210-0113303301210303"></a>

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

<a id="canonical-0022213012310232-3210021222221233-0330332132332321-2223100031113101-2223113110113123-3013102120021333-2111000233221222-3202033131021100"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.ip_prefix_list`

<a id="canonical-2312232003033113-2302032002221002-3032200012102023-3023122133002223-3011002031300010-2100101111132311-2333320033023322-1002323022212002"></a>

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

<a id="canonical-0330231110001002-0332222322223120-2100121122023210-3320332310302010-3233120113221123-0232010201020213-3210122311223300-0000110031131311"></a>

<a id="canonical-1233301313303023-0201011012032333-3021111331110333-2300020203011030-0201132032031113-2000001312002123-2310212121330213-3221020003323201"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_prefix_list.ip_prefixes` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1323031133100220-3003131300131310-1203112022000311-1230021132323020-2132121211131031-2202311300331212-3111230320020110-3011011131332111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.path` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.path

<a id="canonical-1000020111011030-3203310221322031-0320302111200221-3200133321222322-2311112012030313-1301310222031232-1100211231322230-3031222002130313"></a>

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

<a id="canonical-2113131132203211-1220321131033020-3230103032003001-3123031311130220-0101000203001023-3320310322131300-3111112233310233-3130132222222311"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.path`

<a id="canonical-3310030332322101-3102210022203201-1220320102220233-0013001232011113-1202232320121100-3203202211331303-1333133203321303-2001003203011000"></a>

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

<a id="canonical-1122223033203001-3000311122102201-1112321202310120-0311211011310301-2203113031002213-2011131220320102-3100013221113313-3000223103212201"></a>

<a id="canonical-2232233212103003-3233310321133303-3013301322221223-3002001113010110-1211332330201121-3113310021101132-0013003133120322-1321310020132102"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.exact_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1023110331303112-0102023321222113-0213003120013100-1301323300301220-1112023310321022-0030113121201213-0132123123211033-1231010312223202"></a>

<a id="canonical-0331210032231020-1200123320321223-0100310221032200-0322230031302103-1320311210002222-3223232202231232-2032001220123033-1002121100110323"></a>

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

<a id="canonical-0311320231223032-2120202323130222-2203102302001202-3220123301312230-0223002331212213-1232312201330221-3231103223100010-2131210333231203"></a>

<a id="canonical-2011101301212203-2011030010130321-1000032021130213-3000212120121301-3331230001332121-3212021311323031-3001133130202303-0233300101031313"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.prefix_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1333131200321220-0210021133020003-1310122003103122-3312320211030213-1130101133203003-3001321322211311-0033330023013123-0031220001122032"></a>

<a id="canonical-3113302200330222-3121110211310313-1031313330300131-3030132023110201-2022331130333023-2101331030333201-2221333130231332-2003003123010212"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.regex_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0302133323201220-0320231030111112-1220331012311220-2121311010302022-1212033000301113-3020112222212302-3200123011013011-2020200121131103"></a>

<a id="canonical-3023103333133333-1203022200002101-2032301121001101-0113112123131200-1003010200033111-2313021300003321-0211211231332113-0231102120013133"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.suffix_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3101323201032123-3130111021201301-2132330011201310-1001121301111111-3030112133003021-0200010300321033-0202010311201222-3220213200013312"></a>

<a id="canonical-2022221022310113-0011221221300231-3222121003330122-3313323033322012-1320321030032132-2200023022322213-3101313011123013-1021333200022213"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.transformers` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1101012003003233-2202222113013213-2002303112232202-1102102121311030-3310003312033111-2122121111000112-0123121121310202-3132330200111100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.query_params` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.query_params

<a id="canonical-0330013112223030-0223233103011231-2112031210313113-2332310022120312-1232033200100301-2011020202113021-0013031301001100-2302032002200200"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1132210231130311-1002030311322103-3230021211120112-3321020321331033-3132303223231231-0322330302023210-0333113213203310-0303033201012100"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.query_params`

- [check_not_present](resources--cdn_loadbalancer--reference--group-014.md#canonical-1312300330202202-0230330032101303-2331213220023113-3312230312301330-3331100020223113-1303131000020002-2031223202300310-1031130011210022): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-014.md#canonical-0011200213331102-3303001022301002-0122112221113011-1101032132033112-2322010100320111-1203021310120331-0020211133132211-2133230123112201): complete subsection reference.

<a id="canonical-3102010112201132-3222113311210212-1102022201021013-0212100133323203-2302133113021031-2021311301103220-0223311303200101-0312321030212100"></a>

<a id="canonical-2222012103022311-0030322220330112-1102120332311130-0021023322323202-3002021122122121-0332302132022332-2303133301303121-2331030231331102"></a>

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

- [item](resources--cdn_loadbalancer--reference--group-014.md#canonical-0300132120303011-2222013321322121-1103100003033103-0130000011023130-3102203201303332-3231303212012130-0201320300021112-0303323231120233): complete subsection reference.

<a id="canonical-0121103323123020-0322032200300210-1010000123320232-2213220011002320-0201100133022202-2313102110311320-1011212122101331-3220223033023222"></a>

<a id="canonical-0133030032331312-1233120133000302-3213300120100313-2123023110210031-1332300102200210-3123021122200312-2331021102302123-0300313120210132"></a>

#### `policy_based_challenge.rule_list.rules.spec.query_params.key` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1312300330202202-0230330032101303-2331213220023113-3312230312301330-3331100020223113-1303131000020002-2031223202300310-1031130011210022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.query_params.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [policy_based_challenge.rule_list.rules.spec.query_params](resources--cdn_loadbalancer--reference--group-014.md#canonical-1101012003003233-2202222113013213-2002303112232202-1102102121311030-3310003312033111-2122121111000112-0123121121310202-3132330200111100)
- policy_based_challenge.rule_list.rules.spec.query_params.check_not_present

<a id="canonical-3323303321102130-1313123201200100-3000011111221332-0333223210323220-3211233200332202-3033233301200002-1313312001233111-1320232330220113"></a>

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

<a id="canonical-0011200213331102-3303001022301002-0122112221113011-1101032132033112-2322010100320111-1203021310120331-0020211133132211-2133230123112201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.query_params.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [policy_based_challenge.rule_list.rules.spec.query_params](resources--cdn_loadbalancer--reference--group-014.md#canonical-1101012003003233-2202222113013213-2002303112232202-1102102121311030-3310003312033111-2122121111000112-0123121121310202-3132330200111100)
- policy_based_challenge.rule_list.rules.spec.query_params.check_present

<a id="canonical-1220320221120133-0033321113211020-1203312212003131-1313001232113100-3233202022203202-2313321203030320-3133133103002010-3333103002021312"></a>

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

<a id="canonical-0300132120303011-2222013321322121-1103100003033103-0130000011023130-3102203201303332-3231303212012130-0201320300021112-0303323231120233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.query_params.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [policy_based_challenge.rule_list.rules.spec.query_params](resources--cdn_loadbalancer--reference--group-014.md#canonical-1101012003003233-2202222113013213-2002303112232202-1102102121311030-3310003312033111-2122121111000112-0123121121310202-3132330200111100)
- policy_based_challenge.rule_list.rules.spec.query_params.item

<a id="canonical-0102033000211301-3202211131220331-2000333320311112-1232201230231123-3110102200032212-1133321033112230-3202312110233200-1323220303212130"></a>

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

<a id="canonical-0320302312110211-2213212022021231-0002001211221020-1213133231032100-3113200102120021-2112132033230013-1110222310221103-0210313323303010"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.query_params.item`

<a id="canonical-2302023302020011-3110113331121322-2323311121120322-3002132110300213-0301221333230233-3202122303131011-2031213033220032-2023112201313332"></a>

#### `policy_based_challenge.rule_list.rules.spec.query_params.item.exact_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0303000332223232-1232123101211223-1031132013231223-1313113202103030-2031230231323030-0201313012301232-0011021013103222-0022313031313211"></a>

<a id="canonical-2230113320132311-3112331203320201-2202013010133122-1331313012300010-1100330111030201-0113230033203030-3003320213102132-1132100133001000"></a>

#### `policy_based_challenge.rule_list.rules.spec.query_params.item.regex_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1303101121133330-0033031313313021-2302331021320122-3200200211103111-0112311321133021-3332111210220123-2120112213000020-0032123103120022"></a>

<a id="canonical-3020111202023211-1201032333023302-3312132103033210-1000332130003010-1111210023110000-1213010123330302-2120222123221121-1122203303232122"></a>

#### `policy_based_challenge.rule_list.rules.spec.query_params.item.transformers` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3232022003302213-0032131233110110-3231221110233032-1010300113202110-1012221322120230-1322201202221221-2233333330333200-3030322000303331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher

<a id="canonical-2123022203130322-0022112310210121-3033311120121203-1120020002120211-3320221302133201-3012120322302120-1222232201031222-1313313331203301"></a>

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

<a id="canonical-1223320300213331-1200031102303232-0230131000301003-1010132103310221-2303123111012030-3203032231221000-2313232221223113-1332031033323203"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher`

<a id="canonical-0332102033203002-1300033000302223-1320111233030302-2011330223122011-0202033001013312-3103112203002020-1113210220023011-1310023101331132"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0232122031310203-1130322311131121-1133103013313321-3121332003023112-1013321200030331-2202211230101021-0232311203201112-2203200311220213"></a>

<a id="canonical-2021320223331302-1012031110000302-1230311310202221-0032321113110212-2013021311003312-2010200121103002-0100131222101020-2201313210333011"></a>

#### `policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher.exact_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2101002303102000-2012230021120100-3120111331223100-1023203333233010-3323322133212220-1101133231221321-1113220103111033-3000333121121203"></a>

<a id="canonical-0311221132103132-3311012203132000-1231022131021130-2121201313312331-1021120010023320-3213121120101101-1222021223000233-0232311230020123"></a>

#### `policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher.excluded_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2020113203001122-0131003133101010-3112033112103122-1230132210230022-1110101121230003-3011322320122320-2001301222311323-1030230313310012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.temporary_user_blocking` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- policy_based_challenge.temporary_user_blocking

<a id="canonical-1201223313121033-1103213030333221-1331232332133120-3021210133032030-0032312022010311-2131030221021131-2322001320123223-3002233120103201"></a>

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

<a id="canonical-3110003331013331-3130133013301231-3010203210213123-2222111310330220-1203123301030033-3221112233210000-2113132010310313-0201133201120020"></a>

### Direct properties for `policy_based_challenge.temporary_user_blocking`

<a id="canonical-0010103303001310-0213230232201003-3123212123321232-2211122222002222-1330020101323220-1302331232121123-2113123000330012-2233321033223010"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- protected_cookies

<a id="canonical-2223231201320222-0020301102201100-2031133233311002-1000230332031133-1303122030112310-0231202302013302-3122033213202302-1312113133230002"></a>

Type: `"object"`. list nested block, Optional.

Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering
Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering
Protection, enabling a web app firewall (WAF) is a prerequisite. The configured mode of WAF
(monitoring or blocking) will be enforced on the request when cookie tampering is identified. Note:
We recommend enabling Secure and HttpOnly attributes along with cookie tampering protection.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
