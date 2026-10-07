---
page_title: "xcsh_k8s_cluster_role reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_k8s_cluster_role reference."
---

# xcsh_k8s_cluster_role reference

<a id="canonical-2001310313031232-3003303113031320-0000111202132223-2030000001211031-3123033320222313-3312102121311301-3321002322000112-3233122202001100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_k8s_cluster_role](../data-sources/k8s_cluster_role.md#canonical-2232121312000210-3132130123201003-2131123201312211-2021232000312101-1032012313201322-3103300020131231-0202213330300122-0302302323323122)
- Property reference

<a id="canonical-3031122230221132-3130023302220101-3231233330110300-3020023303110003-3023031012000211-1003320030202231-0001011303200330-3102001131313313"></a>

### Direct properties for `xcsh_k8s_cluster_role`

<a id="canonical-3332331233023303-0002333020301331-1023223300323120-3312133230033112-1013222000032313-3200201000120303-2230321102010310-2210233320001103"></a>

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

<a id="canonical-0332310022002301-1021202132111112-1100213300012330-0020112202303112-2001010323210203-3310313201122202-2012102323313303-0222133222131032"></a>

<a id="canonical-3203313202310002-2122023223230131-0210301120331320-0232002031201300-1103111300300100-2102020011320212-2103012120022213-2101221120131322"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the K8SClusterRole.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2112221120023001-3000311300000313-3223103021321000-3232103131320032-2031222230023330-1323313010020123-0221021323022022-2121123203333032"></a>

<a id="canonical-1332103010311220-3322131330012322-0331010321231222-1102321121232321-1021203320022302-3221013110200330-3023330223130102-2200101321313223"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [k8s_cluster_role_selector](data-sources--k8s_cluster_role--reference--group-001.md#canonical-3111022121111201-1311321010121323-3321320021332110-2022010311320012-2310031101112310-1220223003001022-3103221303120022-2223202312021310): complete subsection reference.

<a id="canonical-3020112111022210-1230332213231320-0320320213103120-0212223203232300-0232131230013121-3120133003211213-3211200022121001-1123111300000112"></a>

<a id="canonical-0032312333321321-0320131010212321-0022101222102201-2300113320031323-0322230011010323-2122211011121231-2103022110102310-3031331322331023"></a>

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

<a id="canonical-1311131132332223-2201220302121230-0002031221113102-2020201333232222-1120222120332011-2321012021221332-1010222111130220-2210302011003302"></a>

<a id="canonical-2020022113211202-1111133223322011-0111213120133333-0121101013211110-1101102330330301-2211222111211321-1331132033300102-3210103230122033"></a>

#### `name` property

Type: `"string"`. Required.

Name of the K8SClusterRole.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0320120203331002-1030213010030310-1312222311210000-2302000320103320-0333111130213221-3223203122221323-3230201211301113-1111201322223200"></a>

<a id="canonical-0110212101313003-1103213020032023-2223210113112003-2230313222021201-3100203101000102-2213211130330221-0020020202030011-2021223033212011"></a>

#### `namespace` property

Type: `"string"`. Optional, Computed.

Namespace where the K8SClusterRole exists.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [policy_rule_list](data-sources--k8s_cluster_role--reference--group-001.md#canonical-1221133310032133-1000323320102332-1303231100200023-3322123110022320-3120213002030302-0233133133300231-2300021312001011-3030331123221023): complete subsection reference.

<a id="canonical-2121220023022210-2101001033033322-1222221132322103-2310102331132301-0120100013013121-2331112133330110-3013210210020031-0311130223012001"></a>

<a id="canonical-0220113130121320-0320113120013101-2011013103203033-0131320120213021-0322203132021100-3031000010330021-0223021211302023-1300233122030210"></a>

#### `yaml` property

Type: `"string"`. Computed.

Exclusive with \[k8s\_cluster\_role\_selector policy\_rule\_list\] K8s YAML for ClusterRole.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "Valid parseable YAML",
    "maxLength": 4096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 1,
    "validation": {
      "customRule": "Must be valid YAML"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3203021222303202-2232101232222311-1011211013233031-3323213021201331-3232123130132303-0002030322013001-3133022323030301-1322302222210300"></a>

### All schema paths for `xcsh_k8s_cluster_role`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--k8s_cluster_role--reference--group-001.md#canonical-3332331233023303-0002333020301331-1023223300323120-3312133230033112-1013222000032313-3200201000120303-2230321102010310-2210233320001103) |
| `description` | [description](data-sources--k8s_cluster_role--reference--group-001.md#canonical-0332310022002301-1021202132111112-1100213300012330-0020112202303112-2001010323210203-3310313201122202-2012102323313303-0222133222131032) |
| `id` | [ID](data-sources--k8s_cluster_role--reference--group-001.md#canonical-2112221120023001-3000311300000313-3223103021321000-3232103131320032-2031222230023330-1323313010020123-0221021323022022-2121123203333032) |
| `k8s_cluster_role_selector` | [k8s_cluster_role_selector](data-sources--k8s_cluster_role--reference--group-001.md#canonical-1101010332132132-1303031211313123-0303022320231030-2101100033132122-1313222300202311-2313320131133110-3220312301133213-3212223201120122) |
| `k8s_cluster_role_selector.expressions` | [k8s_cluster_role_selector.expressions](data-sources--k8s_cluster_role--reference--group-001.md#canonical-1210313323111232-3003131131313111-0132001302013302-1011300022120332-1310211023300112-3330022211003120-3302101022202132-2101231301113210) |
| `labels` | [labels](data-sources--k8s_cluster_role--reference--group-001.md#canonical-3020112111022210-1230332213231320-0320320213103120-0212223203232300-0232131230013121-3120133003211213-3211200022121001-1123111300000112) |
| `name` | [name](data-sources--k8s_cluster_role--reference--group-001.md#canonical-1311131132332223-2201220302121230-0002031221113102-2020201333232222-1120222120332011-2321012021221332-1010222111130220-2210302011003302) |
| `namespace` | [namespace](data-sources--k8s_cluster_role--reference--group-001.md#canonical-0320120203331002-1030213010030310-1312222311210000-2302000320103320-0333111130213221-3223203122221323-3230201211301113-1111201322223200) |
| `policy_rule_list` | [policy_rule_list](data-sources--k8s_cluster_role--reference--group-001.md#canonical-2321032230220330-2022223303033312-2130022012203110-1121230012313202-3110333130002103-2221332103013323-1101223230003302-2202030132303211) |
| `policy_rule_list.policy_rule` | [policy_rule_list.policy_rule](data-sources--k8s_cluster_role--reference--group-001.md#canonical-0223210100212110-3132102202200233-3212121023020330-0220111211330202-0321131200231021-3200322031023013-2300112023300330-1212113010020312) |
| `policy_rule_list.policy_rule.non_resource_url_list` | [policy_rule_list.policy_rule.non_resource_url_list](data-sources--k8s_cluster_role--reference--group-001.md#canonical-2310202103203313-3311102112122112-3113121203302102-0321222311332031-1030032013202303-0011002130120000-0030131132333110-3233300312001222) |
| `policy_rule_list.policy_rule.non_resource_url_list.urls` | [policy_rule_list.policy_rule.non_resource_url_list.urls](data-sources--k8s_cluster_role--reference--group-001.md#canonical-1021212320323102-1000220032020110-1121121220331001-3330201203102021-1220322003323130-2323001233303333-2321001223000133-1103321311123230) |
| `policy_rule_list.policy_rule.non_resource_url_list.verbs` | [policy_rule_list.policy_rule.non_resource_url_list.verbs](data-sources--k8s_cluster_role--reference--group-001.md#canonical-3000031322200112-1331033212213003-3033102122103301-3330030012103310-3302011312201111-3100301022201130-2132112201001303-2322001223313001) |
| `policy_rule_list.policy_rule.resource_list` | [policy_rule_list.policy_rule.resource_list](data-sources--k8s_cluster_role--reference--group-001.md#canonical-0230132322203022-1300202003232320-2123223211110212-3202023200200112-1301122100223002-1031011012213032-1020031330231113-3002213202100201) |
| `policy_rule_list.policy_rule.resource_list.api_groups` | [policy_rule_list.policy_rule.resource_list.api_groups](data-sources--k8s_cluster_role--reference--group-001.md#canonical-3130121223023202-0121210001020200-0330313232222323-1333322211321023-3230211103333310-0233121132001000-2201101202203313-1232121222010322) |
| `policy_rule_list.policy_rule.resource_list.resource_instances` | [policy_rule_list.policy_rule.resource_list.resource_instances](data-sources--k8s_cluster_role--reference--group-001.md#canonical-1300012111211111-0130233031211102-1003302301111222-3021021122021200-2333022112022120-2230322011011012-1031101122213120-3320322213330211) |
| `policy_rule_list.policy_rule.resource_list.resource_types` | [policy_rule_list.policy_rule.resource_list.resource_types](data-sources--k8s_cluster_role--reference--group-001.md#canonical-2021033121121303-1121023233120311-3133321003322333-2302233202233213-0221213210022001-2021323022313302-3111231213220001-3221232123001300) |
| `policy_rule_list.policy_rule.resource_list.verbs` | [policy_rule_list.policy_rule.resource_list.verbs](data-sources--k8s_cluster_role--reference--group-001.md#canonical-2200010030311330-2121212132010130-1213000323303332-2332122333010220-2112200122111103-1300020011102110-0332033231030003-3210333311003121) |
| `yaml` | [YAML](data-sources--k8s_cluster_role--reference--group-001.md#canonical-2121220023022210-2101001033033322-1222221132322103-2310102331132301-0120100013013121-2331112133330110-3013210210020031-0311130223012001) |

<a id="canonical-3111022121111201-1311321010121323-3321320021332110-2022010311320012-2310031101112310-1220223003001022-3103221303120022-2223202312021310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `k8s_cluster_role_selector` properties

Breadcrumbs:

- [xcsh_k8s_cluster_role](../data-sources/k8s_cluster_role.md#canonical-2232121312000210-3132130123201003-2131123201312211-2021232000312101-1032012313201322-3103300020131231-0202213330300122-0302302323323122)
- [Property reference](data-sources--k8s_cluster_role--reference--group-001.md#canonical-2001310313031232-3003303113031320-0000111202132223-2030000001211031-3123033320222313-3312102121311301-3321002322000112-3233122202001100)
- k8s_cluster_role_selector

<a id="canonical-1101010332132132-1303031211313123-0303022320231030-2101100033132122-1313222300202311-2313320131133110-3220312301133213-3212223201120122"></a>

Type: `"single"`. Computed.

\[OneOf: k8s\_cluster\_role\_selector, policy\_rule\_list, YAML\] Type can be used to establish a
'selector reference' from one object(called selector) to a set of other objects(called selectees)
based on the value of expressions. A label selector is a label query over a set of resources. An
empty label selector matches all objects.

Additional upstream details:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A null label selector matches
no objects. Label selector is immutable. Expressions is a list of strings of label selection
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

OneOf alternatives in this subsection:

- [k8s_cluster_role_selector](data-sources--k8s_cluster_role--reference--group-001.md#canonical-1101010332132132-1303031211313123-0303022320231030-2101100033132122-1313222300202311-2313320131133110-3220312301133213-3212223201120122)
- [policy_rule_list](data-sources--k8s_cluster_role--reference--group-001.md#canonical-2321032230220330-2022223303033312-2130022012203110-1121230012313202-3110333130002103-2221332103013323-1101223230003302-2202030132303211)
- [YAML](data-sources--k8s_cluster_role--reference--group-001.md#canonical-2121220023022210-2101001033033322-1222221132322103-2310102331132301-0120100013013121-2331112133330110-3013210210020031-0311130223012001)

Select alternatives according to the provider validators above.

<a id="canonical-1113312303122020-0130331223320132-2232311312330312-1022102330203203-3111110303123333-1101033332302033-0030233102203010-3013023020013230"></a>

### Direct properties for `k8s_cluster_role_selector`

<a id="canonical-1210313323111232-3003131131313111-0132001302013302-1011300022120332-1310211023300112-3330022211003120-3302101022202132-2101231301113210"></a>

#### `k8s_cluster_role_selector.expressions` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1221133310032133-1000323320102332-1303231100200023-3322123110022320-3120213002030302-0233133133300231-2300021312001011-3030331123221023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_rule_list` properties

Breadcrumbs:

- [xcsh_k8s_cluster_role](../data-sources/k8s_cluster_role.md#canonical-2232121312000210-3132130123201003-2131123201312211-2021232000312101-1032012313201322-3103300020131231-0202213330300122-0302302323323122)
- [Property reference](data-sources--k8s_cluster_role--reference--group-001.md#canonical-2001310313031232-3003303113031320-0000111202132223-2030000001211031-3123033320222313-3312102121311301-3321002322000112-3233122202001100)
- policy_rule_list

<a id="canonical-2321032230220330-2022223303033312-2130022012203110-1121230012313202-3110333130002103-2221332103013323-1101223230003302-2202030132303211"></a>

Type: `"single"`. Computed.

Policy Rule List. List of rules for role permissions.

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

<a id="canonical-0101331310001133-1122032131211123-0302133032101112-3012101300232022-0111122332011313-1313233221032002-0232220232120012-1131311212320113"></a>

### Direct properties for `policy_rule_list`

- [policy_rule](data-sources--k8s_cluster_role--reference--group-001.md#canonical-1231123021022003-1023022010131212-1013311331321110-0120011230122130-2100122120100221-3200002302323131-0202203323303002-2102030123102011): complete subsection reference.

<a id="canonical-1231123021022003-1023022010131212-1013311331321110-0120011230122130-2100122120100221-3200002302323131-0202203323303002-2102030123102011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_rule_list.policy_rule` properties

Breadcrumbs:

- [xcsh_k8s_cluster_role](../data-sources/k8s_cluster_role.md#canonical-2232121312000210-3132130123201003-2131123201312211-2021232000312101-1032012313201322-3103300020131231-0202213330300122-0302302323323122)
- [Property reference](data-sources--k8s_cluster_role--reference--group-001.md#canonical-2001310313031232-3003303113031320-0000111202132223-2030000001211031-3123033320222313-3312102121311301-3321002322000112-3233122202001100)
- [policy_rule_list](data-sources--k8s_cluster_role--reference--group-001.md#canonical-1221133310032133-1000323320102332-1303231100200023-3322123110022320-3120213002030302-0233133133300231-2300021312001011-3030331123221023)
- policy_rule_list.policy_rule

<a id="canonical-0223210100212110-3132102202200233-3212121023020330-0220111211330202-0321131200231021-3200322031023013-2300112023300330-1212113010020312"></a>

Type: `"list"`. Computed.

Policy Rules. List of rules for role permissions.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1113000110200011-1022203102322023-1303032102302221-0013002332203113-2300312301033233-3112201003023001-1222131300312203-2000310001020310"></a>

### Direct properties for `policy_rule_list.policy_rule`

- [non_resource_url_list](data-sources--k8s_cluster_role--reference--group-001.md#canonical-2112331321312302-3211323002123211-1321233110310211-2022320001313322-2332232211301120-3033230113210203-3011311100302002-3002113230203210): complete subsection reference.

- [resource_list](data-sources--k8s_cluster_role--reference--group-001.md#canonical-3312111220300133-2211133102232101-0011222300333020-2211023000112002-2230223223320322-2233123201230121-3310312202322121-1231113002303221): complete subsection reference.

<a id="canonical-2112331321312302-3211323002123211-1321233110310211-2022320001313322-2332232211301120-3033230113210203-3011311100302002-3002113230203210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_rule_list.policy_rule.non_resource_url_list` properties

Breadcrumbs:

- [xcsh_k8s_cluster_role](../data-sources/k8s_cluster_role.md#canonical-2232121312000210-3132130123201003-2131123201312211-2021232000312101-1032012313201322-3103300020131231-0202213330300122-0302302323323122)
- [Property reference](data-sources--k8s_cluster_role--reference--group-001.md#canonical-2001310313031232-3003303113031320-0000111202132223-2030000001211031-3123033320222313-3312102121311301-3321002322000112-3233122202001100)
- [policy_rule_list](data-sources--k8s_cluster_role--reference--group-001.md#canonical-1221133310032133-1000323320102332-1303231100200023-3322123110022320-3120213002030302-0233133133300231-2300021312001011-3030331123221023)
- [policy_rule_list.policy_rule](data-sources--k8s_cluster_role--reference--group-001.md#canonical-1231123021022003-1023022010131212-1013311331321110-0120011230122130-2100122120100221-3200002302323131-0202203323303002-2102030123102011)
- policy_rule_list.policy_rule.non_resource_url_list

<a id="canonical-2310202103203313-3311102112122112-3113121203302102-0321222311332031-1030032013202303-0011002130120000-0030131132333110-3233300312001222"></a>

Type: `"single"`. Computed.

Permissions for URL(s) that do not represent K8s resource.

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

<a id="canonical-3220200020132021-2223331303202121-3210232010001011-0300123102120311-3203233321002013-0030203300333303-3122300103313301-0013030302323330"></a>

### Direct properties for `policy_rule_list.policy_rule.non_resource_url_list`

<a id="canonical-1021212320323102-1000220032020110-1121121220331001-3330201203102021-1220322003323130-2323001233303333-2321001223000133-1103321311123230"></a>

#### `policy_rule_list.policy_rule.non_resource_url_list.urls` property

Type: `["list", "string"]`. Computed.

Allowed URL(s) that do not represent any K8s resource. URL can be suffix or regular expression.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3000031322200112-1331033212213003-3033102122103301-3330030012103310-3302011312201111-3100301022201130-2132112201001303-2322001223313001"></a>

<a id="canonical-0211203333101132-1133303000220203-3120002203020333-0200211110222232-2202013200031122-3010102020013311-1303031000312133-1113113103002212"></a>

#### `policy_rule_list.policy_rule.non_resource_url_list.verbs` property

Type: `["list", "string"]`. Computed.

Allowed list of verbs(operations) on resources. Use VerbAll for all operations.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3312111220300133-2211133102232101-0011222300333020-2211023000112002-2230223223320322-2233123201230121-3310312202322121-1231113002303221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_rule_list.policy_rule.resource_list` properties

Breadcrumbs:

- [xcsh_k8s_cluster_role](../data-sources/k8s_cluster_role.md#canonical-2232121312000210-3132130123201003-2131123201312211-2021232000312101-1032012313201322-3103300020131231-0202213330300122-0302302323323122)
- [Property reference](data-sources--k8s_cluster_role--reference--group-001.md#canonical-2001310313031232-3003303113031320-0000111202132223-2030000001211031-3123033320222313-3312102121311301-3321002322000112-3233122202001100)
- [policy_rule_list](data-sources--k8s_cluster_role--reference--group-001.md#canonical-1221133310032133-1000323320102332-1303231100200023-3322123110022320-3120213002030302-0233133133300231-2300021312001011-3030331123221023)
- [policy_rule_list.policy_rule](data-sources--k8s_cluster_role--reference--group-001.md#canonical-1231123021022003-1023022010131212-1013311331321110-0120011230122130-2100122120100221-3200002302323131-0202203323303002-2102030123102011)
- policy_rule_list.policy_rule.resource_list

<a id="canonical-0230132322203022-1300202003232320-2123223211110212-3202023200200112-1301122100223002-1031011012213032-1020031330231113-3002213202100201"></a>

Type: `"single"`. Computed.

List of resources in terms of API groups/resource types/resource instances and verbs allowed.

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

<a id="canonical-3111313000022001-0001320103130200-0131333022230013-2002033132132030-2220333122031303-1020330133112331-2330223113330332-0132222303133213"></a>

### Direct properties for `policy_rule_list.policy_rule.resource_list`

<a id="canonical-3130121223023202-0121210001020200-0330313232222323-1333322211321023-3230211103333310-0233121132001000-2201101202203313-1232121222010322"></a>

#### `policy_rule_list.policy_rule.resource_list.api_groups` property

Type: `["list", "string"]`. Computed.

Allowed list of API group that contains resources, all resources of a given API group.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1300012111211111-0130233031211102-1003302301111222-3021021122021200-2333022112022120-2230322011011012-1031101122213120-3320322213330211"></a>

<a id="canonical-3312201311011131-3222013103323213-0321012201123130-3320000130023333-1032222310212212-0311010320231000-1123212103131212-0211123132002003"></a>

#### `policy_rule_list.policy_rule.resource_list.resource_instances` property

Type: `["list", "string"]`. Computed.

Allowed list of resource instances within the resource types.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2021033121121303-1121023233120311-3133321003322333-2302233202233213-0221213210022001-2021323022313302-3111231213220001-3221232123001300"></a>

<a id="canonical-3312333300100003-2013311213233022-2223321101323232-0003020122222101-1300011030302202-1200222033131102-3123110301002032-2233220030210200"></a>

#### `policy_rule_list.policy_rule.resource_list.resource_types` property

Type: `["list", "string"]`. Computed.

Allowed list of resource types within the API groups.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2200010030311330-2121212132010130-1213000323303332-2332122333010220-2112200122111103-1300020011102110-0332033231030003-3210333311003121"></a>

<a id="canonical-2323100001212120-3111120133313231-2032030322302021-2313132201211312-0033310022013030-2321100202312233-0313112200022102-3013303102102021"></a>

#### `policy_rule_list.policy_rule.resource_list.verbs` property

Type: `["list", "string"]`. Computed.

Allowed list of verbs(operations) on resources. Use \* for all operations.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```
