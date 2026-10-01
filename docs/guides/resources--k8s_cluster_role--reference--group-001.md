---
page_title: "xcsh_k8s_cluster_role reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_k8s_cluster_role reference."
---

# xcsh_k8s_cluster_role reference

<a id="canonical-2133203023220330-1102213322212330-0132102322203020-3312031132113031-0113221210031112-2310320230002102-0032121131200320-1022122031320230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023100301323220-3023032133332231-3011232023202111-1230101122332230-2103212003023010-2011002123121123-3303330122130223-0133321133211311"></a>

## Property reference — Property reference / 021021123132 / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-1032203102313310-1310133031312032-2032211220123023-3022101331120133-3300202302000301-2221012031203122-0233002122331003-2331111132121223)
- Property reference

<a id="canonical-0021020303111231-2031031131021111-1101333222212133-3312033331323200-3010133023220110-3213002210022010-1310310012033202-2210213321112121"></a>

## Direct properties — Property reference / 021021123132 / 3

<a id="canonical-3201110310322210-0330110132233113-2310111012022122-2331232103011132-0302313030311002-1131010133123333-3232302313121202-2020111313101033"></a>

<a id="canonical-1023103103030200-3010122010121202-2131232210223110-3113200220112012-0321010232022310-0133022203311331-1122302221111131-1123320102310130"></a>

## annotations property — Property reference / 021021123132 / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

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

<a id="canonical-0223003131230331-2211000023303303-1311101322333100-0012320120030031-0001111011321320-3020101100130110-1322210303300212-0003223220320100"></a>

<a id="canonical-2320112021303221-3033220223100312-3032202213302103-2331030021103200-1230322202113012-0233232030223001-0010021032321313-3302300003200002"></a>

## description property — Property reference / 021021123132 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2100332223220202-3103003012100111-0300033321101032-3101023011010112-1201210020201013-3312010300300122-0013132001132222-3330313033310233"></a>

<a id="canonical-1330201222203023-2021212121112010-2001311302213122-0103300023212131-2312022032312132-3233203312130020-2222212102032203-3233110303321013"></a>

## disable property — Property reference / 021021123132 / 6

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

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

<a id="canonical-3213211012311011-0122033232010212-0020023012003000-1321021000022332-2022332230122130-1213002102210221-2301202311131313-3023110220111110"></a>

<a id="canonical-3131102331102200-1023220033102132-0002221111033032-2033233123031313-1033033233121030-2123200031311200-3311121320120303-1230323313222211"></a>

## ID property — Property reference / 021021123132 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [k8s_cluster_role_selector](resources--k8s_cluster_role--reference--group-001.md#canonical-2012130112223230-2221001233011322-3013232013330020-3003303130102122-1110003201333103-3232213300001201-1001203112330100-1312120312230031): complete subsection reference.

<a id="canonical-1213323333023231-3013033101110231-1002303023210210-1331301312132333-1332131232233211-3331113232301032-2123221211321102-1213310311120011"></a>

<a id="canonical-0332332013111222-3002110133302220-1301220233022013-2003120300333202-2011222032333120-3133221323122312-2333213010020232-3230200220311110"></a>

## labels property — Property reference / 021021123132 / 8

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

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

<a id="canonical-0232103010102322-3212203010113111-3101131130020312-2331121311112300-0130333020130220-0323332012011003-1230311302111021-3022223003311222"></a>

<a id="canonical-1313301200020113-0103320021031232-1211313313120131-3310103133211212-1003203021012200-3010000211110003-1021300200301112-2222132221313120"></a>

## name property — Property reference / 021021123132 / 9

Type: `"string"`. Required.

Name of the K8S Cluster Role. Must be unique within the namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1313300112212313-1322320200222303-3312320200311022-0111013320030300-3321111133132333-0102222011201130-1221210012130122-0123131232103233"></a>

<a id="canonical-0102112203002301-3012121313333000-2003112222303132-0331001123120301-3230122011213122-3130210001333023-0303223033202120-1031202033023101"></a>

## namespace property — Property reference / 021021123132 / 10

Type: `"string"`. Optional, Computed.

Namespace for the K8S Cluster Role. The F5 XC API restricts this resource to the system namespace;
it defaults to that value and may be omitted.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [policy_rule_list](resources--k8s_cluster_role--reference--group-001.md#canonical-1023113210123132-3332103033023120-3101111131023112-2032220111002331-2223303002022013-3220332111132122-3323013121032023-1032303313310031): complete subsection reference.

- [timeouts](resources--k8s_cluster_role--reference--group-001.md#canonical-3002332110020001-2223323210230310-2210320301033330-1013333122000300-3200102002202230-1010030211022032-1300133321301222-1010222230130303): complete subsection reference.

<a id="canonical-1220131223122200-0202032323323322-0110303000212022-1302001311121022-2222100030110222-0132322321000331-3030123030132012-0302103222112103"></a>

<a id="canonical-3111203310333102-2212233130201030-3123010313131222-2201110221312023-2213322212033222-3211221313213301-1300211330330300-0121122311303323"></a>

## yaml property — Property reference / 021021123132 / 11

Type: `"string"`. Optional, Computed.

Exclusive with \[k8s\_cluster\_role\_selector policy\_rule\_list\] K8s YAML for ClusterRole.

Upstream description:

Exclusive with \[k8s\_cluster\_role\_selector policy\_rule\_list\] K8s YAML for ClusterRole.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 4096),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1033001330111130-3201300003102100-1132011301233023-3033112232111131-0111312001101103-2011030122210311-3001212120001301-3232301023003002"></a>

## All schema paths — Property reference / 021021123132 / 12

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--k8s_cluster_role--reference--group-001.md#canonical-3201110310322210-0330110132233113-2310111012022122-2331232103011132-0302313030311002-1131010133123333-3232302313121202-2020111313101033) |
| `description` | [description](resources--k8s_cluster_role--reference--group-001.md#canonical-0223003131230331-2211000023303303-1311101322333100-0012320120030031-0001111011321320-3020101100130110-1322210303300212-0003223220320100) |
| `disable` | [disable](resources--k8s_cluster_role--reference--group-001.md#canonical-2100332223220202-3103003012100111-0300033321101032-3101023011010112-1201210020201013-3312010300300122-0013132001132222-3330313033310233) |
| `id` | [id](resources--k8s_cluster_role--reference--group-001.md#canonical-3213211012311011-0122033232010212-0020023012003000-1321021000022332-2022332230122130-1213002102210221-2301202311131313-3023110220111110) |
| `k8s_cluster_role_selector` | [k8s_cluster_role_selector](resources--k8s_cluster_role--reference--group-001.md#canonical-1123312210301210-2200033231122200-0030121001030213-0212033113202233-3021102333131222-3131203212331110-3323033203203021-0211003011022313) |
| `k8s_cluster_role_selector.expressions` | [k8s_cluster_role_selector.expressions](resources--k8s_cluster_role--reference--group-001.md#canonical-3322011102302012-3000333230303123-2102100213212031-0033132230303212-0312210201312300-2230030331301220-1030102112210131-3202122200111100) |
| `labels` | [labels](resources--k8s_cluster_role--reference--group-001.md#canonical-1213323333023231-3013033101110231-1002303023210210-1331301312132333-1332131232233211-3331113232301032-2123221211321102-1213310311120011) |
| `name` | [name](resources--k8s_cluster_role--reference--group-001.md#canonical-0232103010102322-3212203010113111-3101131130020312-2331121311112300-0130333020130220-0323332012011003-1230311302111021-3022223003311222) |
| `namespace` | [namespace](resources--k8s_cluster_role--reference--group-001.md#canonical-1313300112212313-1322320200222303-3312320200311022-0111013320030300-3321111133132333-0102222011201130-1221210012130122-0123131232103233) |
| `policy_rule_list` | [policy_rule_list](resources--k8s_cluster_role--reference--group-001.md#canonical-2230123002311302-2120320000111321-2111130313201313-3103220022310232-1113131333103022-1013201000100010-1233323332303000-3130212210212032) |
| `policy_rule_list.policy_rule` | [policy_rule_list.policy_rule](resources--k8s_cluster_role--reference--group-001.md#canonical-1320101310210331-3021310002301101-2333331321011110-0220203101333000-0201333121311321-0012221333111111-2120220031002221-0212320200202010) |
| `policy_rule_list.policy_rule.non_resource_url_list` | [policy_rule_list.policy_rule.non_resource_url_list](resources--k8s_cluster_role--reference--group-001.md#canonical-3311133102301131-3020322220111013-1320310310003310-1110022230002331-2321311220120122-2323220221020320-2110103203313130-3122133122102000) |
| `policy_rule_list.policy_rule.non_resource_url_list.urls` | [policy_rule_list.policy_rule.non_resource_url_list.urls](resources--k8s_cluster_role--reference--group-001.md#canonical-0332321231302123-2313310021112231-3213112202322110-2102103011112333-3103100200313320-1111111012211032-2131333311330220-2021022203223110) |
| `policy_rule_list.policy_rule.non_resource_url_list.verbs` | [policy_rule_list.policy_rule.non_resource_url_list.verbs](resources--k8s_cluster_role--reference--group-001.md#canonical-1030222111112211-1002002220210033-0030122010102310-0333213313211121-2011113001132033-2331231001120101-0023103121033232-1021112030220110) |
| `policy_rule_list.policy_rule.resource_list` | [policy_rule_list.policy_rule.resource_list](resources--k8s_cluster_role--reference--group-001.md#canonical-1321032132021321-0112330121130211-2000011323313300-1113210212021112-2231113133232202-0332031103113213-3321113023312312-1213100301321302) |
| `policy_rule_list.policy_rule.resource_list.api_groups` | [policy_rule_list.policy_rule.resource_list.api_groups](resources--k8s_cluster_role--reference--group-001.md#canonical-1000212022211131-1231301003000022-1221311100331032-2331032012133303-1130003102233302-1210333203210101-1130232101022000-1121133302012033) |
| `policy_rule_list.policy_rule.resource_list.resource_instances` | [policy_rule_list.policy_rule.resource_list.resource_instances](resources--k8s_cluster_role--reference--group-001.md#canonical-0003332131102300-2102020021220133-0212012000223323-2313310131332233-1010313313122311-0222101022203302-2330331012102013-0323030030112221) |
| `policy_rule_list.policy_rule.resource_list.resource_types` | [policy_rule_list.policy_rule.resource_list.resource_types](resources--k8s_cluster_role--reference--group-001.md#canonical-1333102100103322-0212021110232003-1303330321103103-3213310223012113-0211111023033012-2113021110031312-1003231312232130-0210221130021200) |
| `policy_rule_list.policy_rule.resource_list.verbs` | [policy_rule_list.policy_rule.resource_list.verbs](resources--k8s_cluster_role--reference--group-001.md#canonical-0303213203322132-0222100113023102-2210020002312220-3200313333131010-0220130331231021-2112332001330312-3221200221120003-1022100200000013) |
| `timeouts` | [timeouts](resources--k8s_cluster_role--reference--group-001.md#canonical-1301303230031021-0101112030010130-3303200300221002-2323230121100023-2120301222111232-3312230031323301-3010230322123211-0231111123222130) |
| `timeouts.create` | [timeouts.create](resources--k8s_cluster_role--reference--group-001.md#canonical-0232022013322223-1113200020002022-3220032330030211-2301010230332133-3101210021220230-2132031321033233-0023213122200322-1131100030331020) |
| `timeouts.delete` | [timeouts.delete](resources--k8s_cluster_role--reference--group-001.md#canonical-0031303312322102-3200020002130123-3222300330122300-2003110201023120-1323121313201323-2313001333333201-2301033233103213-2320103123212103) |
| `timeouts.read` | [timeouts.read](resources--k8s_cluster_role--reference--group-001.md#canonical-2032012312122120-2310002110013312-1303010002211100-2110131123120330-3203132112010332-2131320202111010-2322203212212102-2012223303321330) |
| `timeouts.update` | [timeouts.update](resources--k8s_cluster_role--reference--group-001.md#canonical-0012210212313331-3331013331331220-1023023323233100-1223012323311333-1032333032231310-3002222120110210-1222121201120313-0023311003110110) |
| `yaml` | [yaml](resources--k8s_cluster_role--reference--group-001.md#canonical-1220131223122200-0202032323323322-0110303000212022-1302001311121022-2222100030110222-0132322321000331-3030123030132012-0302103222112103) |

<a id="canonical-0033211300010332-2002232002110302-2013102312030133-1331113023310201-1301222001013221-1001203300220032-0032130312010202-2011020222320211"></a>

## Next pages — Property reference / 021021123132 / 13

- [k8s_cluster_role_selector](resources--k8s_cluster_role--reference--group-001.md#canonical-2012130112223230-2221001233011322-3013232013330020-3003303130102122-1110003201333103-3232213300001201-1001203112330100-1312120312230031)
- [policy_rule_list](resources--k8s_cluster_role--reference--group-001.md#canonical-1023113210123132-3332103033023120-3101111131023112-2032220111002331-2223303002022013-3220332111132122-3323013121032023-1032303313310031)
- [timeouts](resources--k8s_cluster_role--reference--group-001.md#canonical-3002332110020001-2223323210230310-2210320301033330-1013333122000300-3200102002202230-1010030211022032-1300133321301222-1010222230130303)
- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-1032203102313310-1310133031312032-2032211220123023-3022101331120133-3300202302000301-2221012031203122-0233002122331003-2331111132121223)

<a id="canonical-2012130112223230-2221001233011322-3013232013330020-3003303130102122-1110003201333103-3232213300001201-1001203112330100-1312120312230031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121310200102301-0011332213231313-0132203302100001-1300030223310233-0330323033222312-0021002112212001-1012101333113231-2320303002000210"></a>

## k8s_cluster_role_selector — k8s_cluster_role_selector / 003220131331 / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-1032203102313310-1310133031312032-2032211220123023-3022101331120133-3300202302000301-2221012031203122-0233002122331003-2331111132121223)
- [Property reference](resources--k8s_cluster_role--reference--group-001.md#canonical-2133203023220330-1102213322212330-0132102322203020-3312031132113031-0113221210031112-2310320230002102-0032121131200320-1022122031320230)
- k8s_cluster_role_selector

<a id="canonical-1123312210301210-2200033231122200-0030121001030213-0212033113202233-3021102333131222-3131203212331110-3323033203203021-0211003011022313"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: k8s\_cluster\_role\_selector, policy\_rule\_list, yaml\] Type can be used to establish a
'selector reference' from one object(called selector) to a set of other objects(called selectees)
based on the value of expressions. A label selector is a label query over a set of resources. An
empty label selector matches all objects.

Upstream description:

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

OneOf alternatives in this subsection:

- [k8s_cluster_role_selector](resources--k8s_cluster_role--reference--group-001.md#canonical-1123312210301210-2200033231122200-0030121001030213-0212033113202233-3021102333131222-3131203212331110-3323033203203021-0211003011022313)
- [policy_rule_list](resources--k8s_cluster_role--reference--group-001.md#canonical-2230123002311302-2120320000111321-2111130313201313-3103220022310232-1113131333103022-1013201000100010-1233323332303000-3130212210212032)
- [yaml](resources--k8s_cluster_role--reference--group-001.md#canonical-1220131223122200-0202032323323322-0110303000212022-1302001311121022-2222100030110222-0132322321000331-3030123030132012-0302103222112103)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
k8s_cluster_role_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-1323011000223002-3023031313233103-0321220001111310-3222313313322033-2202312031133221-0003100331211302-2220203122033331-1311021010330312"></a>

## Direct properties — k8s_cluster_role_selector / 003220131331 / 3

<a id="canonical-3322011102302012-3000333230303123-2102100213212031-0033132230303212-0312210201312300-2230030331301220-1030102112210131-3202122200111100"></a>

<a id="canonical-0113301310201313-2102320031203130-1313101333323032-1322300123032131-1232102101203123-1010300220010002-1131202232113013-2320011022221212"></a>

## expressions property — k8s_cluster_role_selector / 003220131331 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3002011033200330-3323210100311030-0323210131103211-2113103211201301-3200102200211000-3322331100013201-1210021233113302-1011131230033212"></a>

## Next pages — k8s_cluster_role_selector / 003220131331 / 5

- [Property reference](resources--k8s_cluster_role--reference--group-001.md#canonical-2133203023220330-1102213322212330-0132102322203020-3312031132113031-0113221210031112-2310320230002102-0032121131200320-1022122031320230)
- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-1032203102313310-1310133031312032-2032211220123023-3022101331120133-3300202302000301-2221012031203122-0233002122331003-2331111132121223)

<a id="canonical-1023113210123132-3332103033023120-3101111131023112-2032220111002331-2223303002022013-3220332111132122-3323013121032023-1032303313310031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0010211320212333-2102011102321322-1103331202130313-0220211032321022-2333230233113122-3101332202120300-0132300133231120-3301103002011100"></a>

## policy_rule_list — policy_rule_list / 121312330202 / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-1032203102313310-1310133031312032-2032211220123023-3022101331120133-3300202302000301-2221012031203122-0233002122331003-2331111132121223)
- [Property reference](resources--k8s_cluster_role--reference--group-001.md#canonical-2133203023220330-1102213322212330-0132102322203020-3312031132113031-0113221210031112-2310320230002102-0032121131200320-1022122031320230)
- policy_rule_list

<a id="canonical-2230123002311302-2120320000111321-2111130313201313-3103220022310232-1113131333103022-1013201000100010-1233323332303000-3130212210212032"></a>

Type: `"object"`. single nested block, Optional.

Policy Rule List. List of rules for role permissions.

Upstream description:

List of rules for role permissions.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("policy_rule")}
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
policy_rule_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0200331310202131-1032233012131023-3110300121210022-2000020200220121-2213022032032311-3220023213311320-3210312000302121-1200101233302201"></a>

## Direct properties — policy_rule_list / 121312330202 / 3

- [policy_rule](resources--k8s_cluster_role--reference--group-001.md#canonical-3031232101002032-1311120200232222-2332131333022232-2322030333100023-1022010033311002-1333100321120122-0301021330113311-2030031003323331): complete subsection reference.

<a id="canonical-1232211123210330-1213322233132130-2303113010021112-3220013122300203-1331031010302011-0021130132033311-2301212023313200-0011112120203120"></a>

## Next pages — policy_rule_list / 121312330202 / 4

- [policy_rule_list.policy_rule](resources--k8s_cluster_role--reference--group-001.md#canonical-3031232101002032-1311120200232222-2332131333022232-2322030333100023-1022010033311002-1333100321120122-0301021330113311-2030031003323331)
- [Property reference](resources--k8s_cluster_role--reference--group-001.md#canonical-2133203023220330-1102213322212330-0132102322203020-3312031132113031-0113221210031112-2310320230002102-0032121131200320-1022122031320230)
- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-1032203102313310-1310133031312032-2032211220123023-3022101331120133-3300202302000301-2221012031203122-0233002122331003-2331111132121223)

<a id="canonical-3031232101002032-1311120200232222-2332131333022232-2322030333100023-1022010033311002-1333100321120122-0301021330113311-2030031003323331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322213110303223-3200021222110012-2310322201123312-1002203111331221-0112313002112122-0002332120113221-1012112032221133-2220211133303031"></a>

## policy_rule_list.policy_rule — policy_rule / 322132130211 / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-1032203102313310-1310133031312032-2032211220123023-3022101331120133-3300202302000301-2221012031203122-0233002122331003-2331111132121223)
- [Property reference](resources--k8s_cluster_role--reference--group-001.md#canonical-2133203023220330-1102213322212330-0132102322203020-3312031132113031-0113221210031112-2310320230002102-0032121131200320-1022122031320230)
- [policy_rule_list](resources--k8s_cluster_role--reference--group-001.md#canonical-1023113210123132-3332103033023120-3101111131023112-2032220111002331-2223303002022013-3220332111132122-3323013121032023-1032303313310031)
- policy_rule_list.policy_rule

<a id="canonical-1320101310210331-3021310002301101-2333331321011110-0220203101333000-0201333121311321-0012221333111111-2120220031002221-0212320200202010"></a>

Type: `"object"`. list nested block, Optional.

Policy Rules. List of rules for role permissions.

Upstream description:

List of rules for role permissions.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("non_resource_url_list",
    "resource_list")}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Terraform syntax:

```terraform
policy_rule {
  # Configure direct properties listed below.
}
```

<a id="canonical-2133002021030330-0200133322333033-1322200301223213-2332220122133130-0002323323012101-3312203101331232-0313100303110101-1333001212221112"></a>

## Direct properties — policy_rule / 322132130211 / 3

- [non_resource_url_list](resources--k8s_cluster_role--reference--group-001.md#canonical-0221101002211010-2332121033033023-3021032130332130-1233020012001021-2321102100233310-3202223331333212-2031221200210213-2311022221322023): complete subsection reference.

- [resource_list](resources--k8s_cluster_role--reference--group-001.md#canonical-1213332002311130-3231122203013000-2322301033131231-0011022002223223-2330322021103232-0312300212021310-2300212230331232-1003210302103211): complete subsection reference.

<a id="canonical-3112323222331201-1101120220001022-2311200301200303-0330122122301202-2001130310200222-3000012131102321-2212123203300321-2203322213020311"></a>

## Next pages — policy_rule / 322132130211 / 4

- [policy_rule_list.policy_rule.non_resource_url_list](resources--k8s_cluster_role--reference--group-001.md#canonical-0221101002211010-2332121033033023-3021032130332130-1233020012001021-2321102100233310-3202223331333212-2031221200210213-2311022221322023)
- [policy_rule_list.policy_rule.resource_list](resources--k8s_cluster_role--reference--group-001.md#canonical-1213332002311130-3231122203013000-2322301033131231-0011022002223223-2330322021103232-0312300212021310-2300212230331232-1003210302103211)
- [policy_rule_list](resources--k8s_cluster_role--reference--group-001.md#canonical-1023113210123132-3332103033023120-3101111131023112-2032220111002331-2223303002022013-3220332111132122-3323013121032023-1032303313310031)
- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-1032203102313310-1310133031312032-2032211220123023-3022101331120133-3300202302000301-2221012031203122-0233002122331003-2331111132121223)

<a id="canonical-0221101002211010-2332121033033023-3021032130332130-1233020012001021-2321102100233310-3202223331333212-2031221200210213-2311022221322023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220233012123300-0023211100313220-0101323312201222-1221231200130202-0302223332022133-0311102212031011-1111201122013300-3012013332120103"></a>

## policy_rule_list.policy_rule.non_resource_url_list — non_resource_url_list / 230313010321 / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-1032203102313310-1310133031312032-2032211220123023-3022101331120133-3300202302000301-2221012031203122-0233002122331003-2331111132121223)
- [Property reference](resources--k8s_cluster_role--reference--group-001.md#canonical-2133203023220330-1102213322212330-0132102322203020-3312031132113031-0113221210031112-2310320230002102-0032121131200320-1022122031320230)
- [policy_rule_list](resources--k8s_cluster_role--reference--group-001.md#canonical-1023113210123132-3332103033023120-3101111131023112-2032220111002331-2223303002022013-3220332111132122-3323013121032023-1032303313310031)
- [policy_rule_list.policy_rule](resources--k8s_cluster_role--reference--group-001.md#canonical-3031232101002032-1311120200232222-2332131333022232-2322030333100023-1022010033311002-1333100321120122-0301021330113311-2030031003323331)
- policy_rule_list.policy_rule.non_resource_url_list

<a id="canonical-3311133102301131-3020322220111013-1320310310003310-1110022230002331-2321311220120122-2323220221020320-2110103203313130-3122133122102000"></a>

Type: `"object"`. single nested block, Optional.

Permissions for URL(s) that do not represent K8s resource.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("urls",
    "verbs")}
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
non_resource_url_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3203132203133212-3332233322023010-0003220211110213-3312300000231211-1012123002312203-0022100123030313-3321121311201030-2310132322321032"></a>

## Direct properties — non_resource_url_list / 230313010321 / 3

<a id="canonical-0332321231302123-2313310021112231-3213112202322110-2102103011112333-3103100200313320-1111111012211032-2131333311330220-2021022203223110"></a>

<a id="canonical-3003232312003002-3331001123131202-2112020131030321-2013133000222320-3102021331032222-2232033201312320-1200311001111200-2133221321333320"></a>

## urls property — non_resource_url_list / 230313010321 / 4

Type: `["list", "string"]`. Optional.

Allowed URL(s) that do not represent any K8s resource. URL can be suffix or regex.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1030222111112211-1002002220210033-0030122010102310-0333213313211121-2011113001132033-2331231001120101-0023103121033232-1021112030220110"></a>

<a id="canonical-0011230032012322-1313230130103121-3211223100330323-3013133220001223-3120202220330011-3101131110011100-0013202201103121-2231312201333101"></a>

## verbs property — non_resource_url_list / 230313010321 / 5

Type: `["list", "string"]`. Optional.

Allowed list of verbs(operations) on resources. Use VerbAll for all operations.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3230032233303030-2100230103322320-1132012112030222-2121002303331001-0003332131000200-3102312000300122-2233020132233010-1302223102222203"></a>

## Next pages — non_resource_url_list / 230313010321 / 6

- [policy_rule_list.policy_rule](resources--k8s_cluster_role--reference--group-001.md#canonical-3031232101002032-1311120200232222-2332131333022232-2322030333100023-1022010033311002-1333100321120122-0301021330113311-2030031003323331)
- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-1032203102313310-1310133031312032-2032211220123023-3022101331120133-3300202302000301-2221012031203122-0233002122331003-2331111132121223)

<a id="canonical-1213332002311130-3231122203013000-2322301033131231-0011022002223223-2330322021103232-0312300212021310-2300212230331232-1003210302103211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0210212201103022-3330123333201300-3211230132032133-2222122232031013-1002122101002020-2200310002233321-0321231002232303-1223113212311110"></a>

## policy_rule_list.policy_rule.resource_list — resource_list / 223023221221 / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-1032203102313310-1310133031312032-2032211220123023-3022101331120133-3300202302000301-2221012031203122-0233002122331003-2331111132121223)
- [Property reference](resources--k8s_cluster_role--reference--group-001.md#canonical-2133203023220330-1102213322212330-0132102322203020-3312031132113031-0113221210031112-2310320230002102-0032121131200320-1022122031320230)
- [policy_rule_list](resources--k8s_cluster_role--reference--group-001.md#canonical-1023113210123132-3332103033023120-3101111131023112-2032220111002331-2223303002022013-3220332111132122-3323013121032023-1032303313310031)
- [policy_rule_list.policy_rule](resources--k8s_cluster_role--reference--group-001.md#canonical-3031232101002032-1311120200232222-2332131333022232-2322030333100023-1022010033311002-1333100321120122-0301021330113311-2030031003323331)
- policy_rule_list.policy_rule.resource_list

<a id="canonical-1321032132021321-0112330121130211-2000011323313300-1113210212021112-2231113133232202-0332031103113213-3321113023312312-1213100301321302"></a>

Type: `"object"`. single nested block, Optional.

List of resources in terms of API groups/resource types/resource instances and verbs allowed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("api_groups",
    "resource_types",
    "verbs")}
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
resource_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0133230230101323-1211102033232210-2323213023320320-2023000333001131-2120112213201231-0022133221013122-3332102232221210-2021110102111330"></a>

## Direct properties — resource_list / 223023221221 / 3

<a id="canonical-1000212022211131-1231301003000022-1221311100331032-2331032012133303-1130003102233302-1210333203210101-1130232101022000-1121133302012033"></a>

<a id="canonical-1110321030333110-2313303211310223-2112100022320130-0022330201211311-0210031023203131-2202321322003021-0321320230123333-1111023233310303"></a>

## api_groups property — resource_list / 223023221221 / 4

Type: `["list", "string"]`. Optional.

Allowed list of API group that contains resources, all resources of a given API group.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0003332131102300-2102020021220133-0212012000223323-2313310131332233-1010313313122311-0222101022203302-2330331012102013-0323030030112221"></a>

<a id="canonical-1301021110021221-1003020130230103-2000111013332011-3131030122312103-0132112101300201-0133213102300110-3323332320210002-2023112121220223"></a>

## resource_instances property — resource_list / 223023221221 / 5

Type: `["list", "string"]`. Optional.

Allowed list of resource instances within the resource types.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1333102100103322-0212021110232003-1303330321103103-3213310223012113-0211111023033012-2113021110031312-1003231312232130-0210221130021200"></a>

<a id="canonical-0200230222133211-0213120113220202-3233312230322120-1122230001300011-3333303010323021-0122122022331031-1013121311100221-0301100132033012"></a>

## resource_types property — resource_list / 223023221221 / 6

Type: `["list", "string"]`. Optional.

Allowed list of resource types within the API groups.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0303213203322132-0222100113023102-2210020002312220-3200313333131010-0220130331231021-2112332001330312-3221200221120003-1022100200000013"></a>

<a id="canonical-1103302013013312-1101302112313213-1103110200322100-0022002203232011-2130033303101022-3130322101310022-0313100322321320-3203022022302121"></a>

## verbs property — resource_list / 223023221221 / 7

Type: `["list", "string"]`. Optional.

Allowed list of verbs(operations) on resources. Use \* for all operations.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2130323011212123-3303031101110323-1010202121211220-0213130323202333-3021001013030111-0320302211312231-0122321131320120-1232113100321202"></a>

## Next pages — resource_list / 223023221221 / 8

- [policy_rule_list.policy_rule](resources--k8s_cluster_role--reference--group-001.md#canonical-3031232101002032-1311120200232222-2332131333022232-2322030333100023-1022010033311002-1333100321120122-0301021330113311-2030031003323331)
- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-1032203102313310-1310133031312032-2032211220123023-3022101331120133-3300202302000301-2221012031203122-0233002122331003-2331111132121223)

<a id="canonical-3002332110020001-2223323210230310-2210320301033330-1013333122000300-3200102002202230-1010030211022032-1300133321301222-1010222230130303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321210021011210-3301000320001233-2111112203332103-2331221012031131-0332212200111200-1200012231212312-3022031022010222-1312033232100031"></a>

## timeouts — timeouts / 221122100320 / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-1032203102313310-1310133031312032-2032211220123023-3022101331120133-3300202302000301-2221012031203122-0233002122331003-2331111132121223)
- [Property reference](resources--k8s_cluster_role--reference--group-001.md#canonical-2133203023220330-1102213322212330-0132102322203020-3312031132113031-0113221210031112-2310320230002102-0032121131200320-1022122031320230)
- timeouts

<a id="canonical-1301303230031021-0101112030010130-3303200300221002-2323230121100023-2120301222111232-3312230031323301-3010230322123211-0231111123222130"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0131310200113120-3221011112110111-1302321021320110-0111202033330331-3213002330110033-1112001122210222-3133003333102022-1023221113000300"></a>

## Direct properties — timeouts / 221122100320 / 3

<a id="canonical-0232022013322223-1113200020002022-3220032330030211-2301010230332133-3101210021220230-2132031321033233-0023213122200322-1131100030331020"></a>

<a id="canonical-0000312111220322-0301010101111213-3321211120120233-0120132222313202-1310332322133012-3110123231100201-3310201322331200-2111212132130231"></a>

## create property — timeouts / 221122100320 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0031303312322102-3200020002130123-3222300330122300-2003110201023120-1323121313201323-2313001333333201-2301033233103213-2320103123212103"></a>

<a id="canonical-0011311300311020-0021313100211113-0011013023333033-1301101021122121-1333011313031113-0203203312132323-1200200323120012-1201230001003000"></a>

## delete property — timeouts / 221122100320 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2032012312122120-2310002110013312-1303010002211100-2110131123120330-3203132112010332-2131320202111010-2322203212212102-2012223303321330"></a>

<a id="canonical-1003020123133323-1301103223013221-1231213312120203-2110320102132212-1223113202220110-0330111330112210-2331110031210011-1321021201301100"></a>

## read property — timeouts / 221122100320 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0012210212313331-3331013331331220-1023023323233100-1223012323311333-1032333032231310-3002222120110210-1222121201120313-0023311003110110"></a>

<a id="canonical-0203333321332233-0320220321313110-0230101002131102-0220331011111203-3121322321212230-0123120302222332-1221332231033220-3120001322131230"></a>

## update property — timeouts / 221122100320 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0013220212032310-3121201313010003-0210030123321223-1123221322203232-2211212132330022-3110320213101311-1101301020312331-3221012211121012"></a>

## Next pages — timeouts / 221122100320 / 8

- [Property reference](resources--k8s_cluster_role--reference--group-001.md#canonical-2133203023220330-1102213322212330-0132102322203020-3312031132113031-0113221210031112-2310320230002102-0032121131200320-1022122031320230)
- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-1032203102313310-1310133031312032-2032211220123023-3022101331120133-3300202302000301-2221012031203122-0233002122331003-2331111132121223)
