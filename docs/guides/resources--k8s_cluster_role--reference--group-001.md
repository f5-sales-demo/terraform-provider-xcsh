---
page_title: "xcsh_k8s_cluster_role reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_k8s_cluster_role reference."
---

# xcsh_k8s_cluster_role reference

<a id="canonical-9f8cba3c529fa9bc1e4ba8c8f635e5cd17a64356b4e2c0920e65d8384a68de2c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b431ee8cb39ffadc5b8b8956c45afac939832c48509b65bf3f1a72b1fe5f975"></a>

## Property reference — Property reference / e837342496de / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-4e8d2df4747cdd8e8e9686cbca47d61ff08b2031a918d8da2f09af43bd55e66b)
- Property reference

<a id="canonical-0923356d8d35d25551fea99ff63fdee0c47cba14e70a428474d063e2a49f9599"></a>

## Direct properties — Property reference / e837342496de / 3

<a id="canonical-e1534ea43c51ebd7b454629abdb9315e32dccd425d11f6ffeecb76628857744f"></a>

<a id="canonical-4b4d3320c46846629dba4ad4d78285863912e2b41f2a3d7d5aca955d5be12d1c"></a>

## annotations property — Property reference / e837342496de / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
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

<a id="canonical-2b0ddb3da500bcf37547afd006e1830d01545e78c84507147a933c2603ae8e10"></a>

<a id="canonical-b8589ce9cfa2b436ce8a7c93bd3094e06cea25c62fb8cac10424ee77f2c03802"></a>

## description property — Property reference / e837342496de / 5

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

<a id="canonical-90faba22d30c6415303f944ed12c511661908847f6130c1a077817aafcdcfd2f"></a>

<a id="canonical-7c86a8cb8999958481d729da13c0b99db628ed9eef8f6708aa9923a3ef533e47"></a>

## disable property — Property reference / e837342496de / 6

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

<a id="canonical-e7946d451a3ee126082c60c0792402be8afac69c67092929b18b5777cb528554"></a>

<a id="canonical-dd4bd4a04ba0f49e02a553ce8fbdb3774f3ef64c9b80dd60f56786336cef7aa5"></a>

## id property — Property reference / e837342496de / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [k8s_cluster_role_selector](resources--k8s_cluster_role--reference--group-001.md#canonical-86716aeca906f17ac7b87f08c3cdc49a540e1fd3ee9f0061418d6f1076636b0d): complete subsection reference.

<a id="canonical-67eff2edc73d152d42ccb9247dc767bf7e76ebe5fd5eec4e9ba65e5267d35605"></a>

<a id="canonical-3ef8756ac251fca871a2f28783630fe285a8efd8dfa7b6b6bf9c422eec828d54"></a>

## labels property — Property reference / e837342496de / 8

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
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

<a id="canonical-2e4c44bae68c45d5d175c236bd6755b01cfc87283bf861436cd72549caac3d6a"></a>

<a id="canonical-77c6021713e0936e65df761df44df966438c91a0c402550349c20c56aa7a9dd8"></a>

## name property — Property reference / e837342496de / 9

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

<a id="canonical-77c169b77ae20ab3f6e20d4a151f8330f955f7bf12a8585c6990671a1b76e4ef"></a>

<a id="canonical-125a30b1c6677fc0835aacde3d05b631ec6859dadc901fcb33acf8984d88f2d1"></a>

## namespace property — Property reference / e837342496de / 10

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

- [policy_rule_list](resources--k8s_cluster_role--reference--group-001.md#canonical-4b5e46defe4cf2d8d155d2d68ea150bdabcc2287e8f9579afb1d938b4ecf7d0d): complete subsection reference.

- [timeouts](resources--k8s_cluster_role--reference--group-001.md#canonical-c2f94201abee4b34a4e313fc47fda030e04828ac4432528e707f9c6a44aac733): complete subsection reference.

<a id="canonical-6876b6a0223bbefa14cc098a7207564aaa40c52a1eeb903dcc6cc786324ea593"></a>

<a id="canonical-d58f4fd2a6bdc84cdb13776aa1529d8ba7ea63eae5a779f17097cf30196b5cfb"></a>

## yaml property — Property reference / e837342496de / 11

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

<a id="canonical-4f07c55ce1c034905e171bcbcf5ae55d15d814538531a935c1998071eec4b0c2"></a>

## All schema paths — Property reference / e837342496de / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--k8s_cluster_role--reference--group-001.md#canonical-e1534ea43c51ebd7b454629abdb9315e32dccd425d11f6ffeecb76628857744f) |
| `description` | [description](resources--k8s_cluster_role--reference--group-001.md#canonical-2b0ddb3da500bcf37547afd006e1830d01545e78c84507147a933c2603ae8e10) |
| `disable` | [disable](resources--k8s_cluster_role--reference--group-001.md#canonical-90faba22d30c6415303f944ed12c511661908847f6130c1a077817aafcdcfd2f) |
| `id` | [id](resources--k8s_cluster_role--reference--group-001.md#canonical-e7946d451a3ee126082c60c0792402be8afac69c67092929b18b5777cb528554) |
| `k8s_cluster_role_selector` | [k8s_cluster_role_selector](resources--k8s_cluster_role--reference--group-001.md#canonical-5bda4c64a03ed6a00c641327263d78afc94bf76add8e6f54fb3e38c9250c52b7) |
| `k8s_cluster_role_selector.expressions` | [k8s_cluster_role_selector.expressions](resources--k8s_cluster_role--reference--group-001.md#canonical-fa152c86c0feccdb9242798d0f7acce636921db0ac33dc684c49691de26a0550) |
| `labels` | [labels](resources--k8s_cluster_role--reference--group-001.md#canonical-67eff2edc73d152d42ccb9247dc767bf7e76ebe5fd5eec4e9ba65e5267d35605) |
| `name` | [name](resources--k8s_cluster_role--reference--group-001.md#canonical-2e4c44bae68c45d5d175c236bd6755b01cfc87283bf861436cd72549caac3d6a) |
| `namespace` | [namespace](resources--k8s_cluster_role--reference--group-001.md#canonical-77c169b77ae20ab3f6e20d4a151f8330f955f7bf12a8585c6990671a1b76e4ef) |
| `policy_rule_list` | [policy_rule_list](resources--k8s_cluster_role--reference--group-001.md#canonical-ac6c2d7298e0057995737877d3a0ad2e5777f4ca478404046fefecc0dc9a498e) |
| `policy_rule_list.policy_rule` | [policy_rule_list.policy_rule](resources--k8s_cluster_role--reference--group-001.md#canonical-7847493dc9d02c51bff79154288d1fc021fd9d7906a7f55598a0d0a926e20884) |
| `policy_rule_list.policy_rule.non_resource_url_list` | [policy_rule_list.policy_rule.non_resource_url_list](resources--k8s_cluster_role--reference--group-001.md#canonical-f57d2c5dc8ea854778d340f4542ac0bdb9d6861abba29238944e3ddcda7da480) |
| `policy_rule_list.policy_rule.non_resource_url_list.urls` | [policy_rule_list.policy_rule.non_resource_url_list.urls](resources--k8s_cluster_role--reference--group-001.md#canonical-3ee6dc9bb7d095ade75a2e94924c55bfd3420df85554694e9dff5f28892a3ad4) |
| `policy_rule_list.policy_rule.non_resource_url_list.verbs` | [policy_rule_list.policy_rule.non_resource_url_list.verbs](resources--k8s_cluster_role--reference--group-001.md#canonical-4ca955a5420a890f0c6844b43f9f7959855c178fbdb416110b4d93ee4958ca14) |
| `policy_rule_list.policy_rule.resource_list` | [policy_rule_list.policy_rule.resource_list](resources--k8s_cluster_role--reference--group-001.md#canonical-7939e27916f197258017bdf057926256ad5dfba23e3535e7f95cbdb667431e72) |
| `policy_rule_list.policy_rule.resource_list.api_groups` | [policy_rule_list.policy_rule.resource_list.api_groups](resources--k8s_cluster_role--reference--group-001.md#canonical-4098a95d6dc4300a69d50f4ebd3867f35c0d2bf264fe39115cb91280597f218f) |
| `policy_rule_list.policy_rule.resource_list.resource_instances` | [policy_rule_list.policy_rule.resource_list.resource_instances](resources--k8s_cluster_role--reference--group-001.md#canonical-03f9d4b092209a1f26180afbb7d1dfaf44df76b52a44a8f2bcf464873b30c5a9) |
| `policy_rule_list.policy_rule.resource_list.resource_types` | [policy_rule_list.policy_rule.resource_list.resource_types](resources--k8s_cluster_role--reference--group-001.md#canonical-7f4904fa26254b8373f394d3e7d2b1972554b3c69725437643b76b9c24a5c260) |
| `policy_rule_list.policy_rule.resource_list.verbs` | [policy_rule_list.policy_rule.resource_list.verbs](resources--k8s_cluster_role--reference--group-001.md#canonical-339e3e9e2a4172d2a4202da8e0dff7442873db4996f81f36e98296034a420007) |
| `timeouts` | [timeouts](resources--k8s_cluster_role--reference--group-001.md#canonical-71cec3491158c11cf3830a42bbb1940b98c6a56ef6b0def1c4b3a6e52d55ba9c) |
| `timeouts.create` | [timeouts.create](resources--k8s_cluster_role--reference--group-001.md#canonical-2e287eab5780808ae83bc325b112cf9fd1909a2c9e3793ef0b9da83a5d40cf48) |
| `timeouts.delete` | [timeouts.delete](resources--k8s_cluster_role--reference--group-001.md#canonical-0dcf6e92e020271beac3c6b0835212d87b67787bb707ffe1b13ef4e7b84db993) |
| `timeouts.read` | [timeouts.read](resources--k8s_cluster_role--reference--group-001.md#canonical-8e1b6698b40941f6731029509475b63ce379613e9de22544ba8e699286af3e7c) |
| `timeouts.update` | [timeouts.update](resources--k8s_cluster_role--reference--group-001.md#canonical-06926dfdfd1fdf684b2fbbd06b1bbd7f4efceb74c2a985246a6616370bd43514) |
| `yaml` | [yaml](resources--k8s_cluster_role--reference--group-001.md#canonical-6876b6a0223bbefa14cc098a7207564aaa40c52a1eeb903dcc6cc786324ea593) |

<a id="canonical-0f97013e82b82532874b631f7d5cbd2171a811e9418f0a0e0e7361228522ae25"></a>

## Next pages — Property reference / e837342496de / 13

- [k8s_cluster_role_selector](resources--k8s_cluster_role--reference--group-001.md#canonical-86716aeca906f17ac7b87f08c3cdc49a540e1fd3ee9f0061418d6f1076636b0d)
- [policy_rule_list](resources--k8s_cluster_role--reference--group-001.md#canonical-4b5e46defe4cf2d8d155d2d68ea150bdabcc2287e8f9579afb1d938b4ecf7d0d)
- [timeouts](resources--k8s_cluster_role--reference--group-001.md#canonical-c2f94201abee4b34a4e313fc47fda030e04828ac4432528e707f9c6a44aac733)
- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-4e8d2df4747cdd8e8e9686cbca47d61ff08b2031a918d8da2f09af43bd55e66b)

<a id="canonical-86716aeca906f17ac7b87f08c3cdc49a540e1fd3ee9f0061418d6f1076636b0d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d9d204b105fa7b771e8f24017032bd2f3cecfab6090969814647f5edb8cc2024"></a>

## k8s_cluster_role_selector — k8s_cluster_role_selector / 203be40e877d / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-4e8d2df4747cdd8e8e9686cbca47d61ff08b2031a918d8da2f09af43bd55e66b)
- [Property reference](resources--k8s_cluster_role--reference--group-001.md#canonical-9f8cba3c529fa9bc1e4ba8c8f635e5cd17a64356b4e2c0920e65d8384a68de2c)
- k8s_cluster_role_selector

<a id="canonical-5bda4c64a03ed6a00c641327263d78afc94bf76add8e6f54fb3e38c9250c52b7"></a>

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

- [k8s_cluster_role_selector](resources--k8s_cluster_role--reference--group-001.md#canonical-5bda4c64a03ed6a00c641327263d78afc94bf76add8e6f54fb3e38c9250c52b7)
- [policy_rule_list](resources--k8s_cluster_role--reference--group-001.md#canonical-ac6c2d7298e0057995737877d3a0ad2e5777f4ca478404046fefecc0dc9a498e)
- [yaml](resources--k8s_cluster_role--reference--group-001.md#canonical-6876b6a0223bbefa14cc098a7207564aaa40c52a1eeb903dcc6cc786324ea593)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
k8s_cluster_role_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-7b140ac2cb377bd339a01574eadf7e8fa2d8d7e90343d972a88da3fd75244f36"></a>

## Direct properties — k8s_cluster_role_selector / 203be40e877d / 3

<a id="canonical-fa152c86c0feccdb9242798d0f7acce636921db0ac33dc684c49691de26a0550"></a>

<a id="canonical-17c7487792e0d8dc7747fece7ac1b39d6e4918db44c281025d8ae5c7b814aa66"></a>

## expressions property — k8s_cluster_role_selector / 203be40e877d / 4

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

<a id="canonical-c214f83cfb910d4c3b91d4e5974e5871e04a0940faf501e16426f5f24576c3e6"></a>

## Next pages — k8s_cluster_role_selector / 203be40e877d / 5

- [Property reference](resources--k8s_cluster_role--reference--group-001.md#canonical-9f8cba3c529fa9bc1e4ba8c8f635e5cd17a64356b4e2c0920e65d8384a68de2c)
- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-4e8d2df4747cdd8e8e9686cbca47d61ff08b2031a918d8da2f09af43bd55e66b)

<a id="canonical-4b5e46defe4cf2d8d155d2d68ea150bdabcc2287e8f9579afb1d938b4ecf7d0d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-049789bf92152e7a53f627372894ee4abfb2f5dad1fa26301ec1fb58f14c2150"></a>

## policy_rule_list — policy_rule_list / 0d5975676f22 / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-4e8d2df4747cdd8e8e9686cbca47d61ff08b2031a918d8da2f09af43bd55e66b)
- [Property reference](resources--k8s_cluster_role--reference--group-001.md#canonical-9f8cba3c529fa9bc1e4ba8c8f635e5cd17a64356b4e2c0920e65d8384a68de2c)
- policy_rule_list

<a id="canonical-ac6c2d7298e0057995737877d3a0ad2e5777f4ca478404046fefecc0dc9a498e"></a>

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

<a id="canonical-20f7489d4ebc674bd4c1990a80220a19a728e3b5e82e7d78e4d80c996046fca1"></a>

## Direct properties — policy_rule_list / 0d5975676f22 / 3

- [policy_rule](resources--k8s_cluster_role--reference--group-001.md#canonical-cdb9108e75620baabe77f2aeba33f40b4a10fd427f43961a3127c5f58c343efd): complete subsection reference.

<a id="canonical-6e95b93c67eaf79cb35c4256e81dac237d344c850971e3f5b198bde0055988d8"></a>

## Next pages — policy_rule_list / 0d5975676f22 / 4

- [policy_rule_list.policy_rule](resources--k8s_cluster_role--reference--group-001.md#canonical-cdb9108e75620baabe77f2aeba33f40b4a10fd427f43961a3127c5f58c343efd)
- [Property reference](resources--k8s_cluster_role--reference--group-001.md#canonical-9f8cba3c529fa9bc1e4ba8c8f635e5cd17a64356b4e2c0920e65d8384a68de2c)
- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-4e8d2df4747cdd8e8e9686cbca47d61ff08b2031a918d8da2f09af43bd55e66b)

<a id="canonical-cdb9108e75620baabe77f2aeba33f40b4a10fd427f43961a3127c5f58c343efd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a9d4cebe026a506b4ea16f6428d5f6916dc259a02f985e94658ea5fa895fccd"></a>

## policy_rule_list.policy_rule — policy_rule_list.policy_rule / 986a25e9e725 / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-4e8d2df4747cdd8e8e9686cbca47d61ff08b2031a918d8da2f09af43bd55e66b)
- [Property reference](resources--k8s_cluster_role--reference--group-001.md#canonical-9f8cba3c529fa9bc1e4ba8c8f635e5cd17a64356b4e2c0920e65d8384a68de2c)
- [policy_rule_list](resources--k8s_cluster_role--reference--group-001.md#canonical-4b5e46defe4cf2d8d155d2d68ea150bdabcc2287e8f9579afb1d938b4ecf7d0d)
- policy_rule_list.policy_rule

<a id="canonical-7847493dc9d02c51bff79154288d1fc021fd9d7906a7f55598a0d0a926e20884"></a>

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

<a id="canonical-9f08933c207fafcf7a831ae7bea1a7dc02efb191f68d1f6e374335117f066a56"></a>

## Direct properties — policy_rule_list.policy_rule / 986a25e9e725 / 3

- [non_resource_url_list](resources--k8s_cluster_role--reference--group-001.md#canonical-29442944be64f3cbc939cf9c6f206049b9490bf4e2afdfe68da60927b52a9e8b): complete subsection reference.

- [resource_list](resources--k8s_cluster_role--reference--group-001.md#canonical-67f82d5ced6a31c0bac4f76d05282aebbce894ee36c26274b09acf6e439324e5): complete subsection reference.

<a id="canonical-d6eeaf615162804ab58318333c69ac628173482ac019d4b9a66e3c39a3ea7235"></a>

## Next pages — policy_rule_list.policy_rule / 986a25e9e725 / 4

- [policy_rule_list.policy_rule.non_resource_url_list](resources--k8s_cluster_role--reference--group-001.md#canonical-29442944be64f3cbc939cf9c6f206049b9490bf4e2afdfe68da60927b52a9e8b)
- [policy_rule_list.policy_rule.resource_list](resources--k8s_cluster_role--reference--group-001.md#canonical-67f82d5ced6a31c0bac4f76d05282aebbce894ee36c26274b09acf6e439324e5)
- [policy_rule_list](resources--k8s_cluster_role--reference--group-001.md#canonical-4b5e46defe4cf2d8d155d2d68ea150bdabcc2287e8f9579afb1d938b4ecf7d0d)
- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-4e8d2df4747cdd8e8e9686cbca47d61ff08b2031a918d8da2f09af43bd55e66b)

<a id="canonical-29442944be64f3cbc939cf9c6f206049b9490bf4e2afdfe68da60927b52a9e8b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-28bc66f00b950de811ef686a69b6072232afe29f354a63455585a1f0c61fe613"></a>

## policy_rule_list.policy_rule.non_resource_url_list — policy_rule_list.policy_rule.non_resource_url_list / a6697eb37139 / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-4e8d2df4747cdd8e8e9686cbca47d61ff08b2031a918d8da2f09af43bd55e66b)
- [Property reference](resources--k8s_cluster_role--reference--group-001.md#canonical-9f8cba3c529fa9bc1e4ba8c8f635e5cd17a64356b4e2c0920e65d8384a68de2c)
- [policy_rule_list](resources--k8s_cluster_role--reference--group-001.md#canonical-4b5e46defe4cf2d8d155d2d68ea150bdabcc2287e8f9579afb1d938b4ecf7d0d)
- [policy_rule_list.policy_rule](resources--k8s_cluster_role--reference--group-001.md#canonical-cdb9108e75620baabe77f2aeba33f40b4a10fd427f43961a3127c5f58c343efd)
- policy_rule_list.policy_rule.non_resource_url_list

<a id="canonical-f57d2c5dc8ea854778d340f4542ac0bdb9d6861abba29238944e3ddcda7da480"></a>

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

<a id="canonical-e37a37e6febfa2c403a25527f6c00b65466c2da30a41b337f967584cb47bae4e"></a>

## Direct properties — policy_rule_list.policy_rule.non_resource_url_list / a6697eb37139 / 3

<a id="canonical-3ee6dc9bb7d095ade75a2e94924c55bfd3420df85554694e9dff5f28892a3ad4"></a>

<a id="canonical-c3bb60c2fd05b7629621d339877c0ab8d227d3aaae3e1db860d415609fa79ff8"></a>

## urls property — policy_rule_list.policy_rule.non_resource_url_list / a6697eb37139 / 4

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

<a id="canonical-4ca955a5420a890f0c6844b43f9f7959855c178fbdb416110b4d93ee4958ca14"></a>

<a id="canonical-05b0e1ba77b1c4d9e5ad0f3bc77e806bd88a8f05d1754150078a14d9adda1fd1"></a>

## verbs property — policy_rule_list.policy_rule.non_resource_url_list / a6697eb37139 / 5

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

<a id="canonical-ec3afccc90b13eb85e19632a990b3f4103f9d020d2d80c1aaf21ebc472ad2aa3"></a>

## Next pages — policy_rule_list.policy_rule.non_resource_url_list / a6697eb37139 / 6

- [policy_rule_list.policy_rule](resources--k8s_cluster_role--reference--group-001.md#canonical-cdb9108e75620baabe77f2aeba33f40b4a10fd427f43961a3127c5f58c343efd)
- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-4e8d2df4747cdd8e8e9686cbca47d61ff08b2031a918d8da2f09af43bd55e66b)

<a id="canonical-67f82d5ced6a31c0bac4f76d05282aebbce894ee36c26274b09acf6e439324e5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-249a14cafc6ff870e5b1e39faa6ae34742691088a0d02bf939b42bb36b5e6d54"></a>

## policy_rule_list.policy_rule.resource_list — policy_rule_list.policy_rule.resource_list / 770464acba69 / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-4e8d2df4747cdd8e8e9686cbca47d61ff08b2031a918d8da2f09af43bd55e66b)
- [Property reference](resources--k8s_cluster_role--reference--group-001.md#canonical-9f8cba3c529fa9bc1e4ba8c8f635e5cd17a64356b4e2c0920e65d8384a68de2c)
- [policy_rule_list](resources--k8s_cluster_role--reference--group-001.md#canonical-4b5e46defe4cf2d8d155d2d68ea150bdabcc2287e8f9579afb1d938b4ecf7d0d)
- [policy_rule_list.policy_rule](resources--k8s_cluster_role--reference--group-001.md#canonical-cdb9108e75620baabe77f2aeba33f40b4a10fd427f43961a3127c5f58c343efd)
- policy_rule_list.policy_rule.resource_list

<a id="canonical-7939e27916f197258017bdf057926256ad5dfba23e3535e7f95cbdb667431e72"></a>

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

<a id="canonical-1fb2c47b6548fba4bb9cbe388b03f05d985a786d0a7e91dafe4aea648951257c"></a>

## Direct properties — policy_rule_list.policy_rule.resource_list / 770464acba69 / 3

<a id="canonical-4098a95d6dc4300a69d50f4ebd3867f35c0d2bf264fe39115cb91280597f218f"></a>

<a id="canonical-54e4cfd4b7ce5d2b9640ae1c0af219752434b8dda2e7a0c939e2c6ff552efd33"></a>

## api_groups property — policy_rule_list.policy_rule.resource_list / 770464acba69 / 4

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

<a id="canonical-03f9d4b092209a1f26180afbb7d1dfaf44df76b52a44a8f2bcf464873b30c5a9"></a>

<a id="canonical-712542694321cb1380547f85dd31ad931e591c211f9d2c14fbfb89028b599a2b"></a>

## resource_instances property — policy_rule_list.policy_rule.resource_list / 770464acba69 / 5

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

<a id="canonical-7f4904fa26254b8373f394d3e7d2b1972554b3c69725437643b76b9c24a5c260"></a>

<a id="canonical-20b2a7e527617a22efdace985ab01c05ffcc4ec91a68af4d476754293141e3c6"></a>

## resource_types property — policy_rule_list.policy_rule.resource_list / 770464acba69 / 6

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

<a id="canonical-339e3e9e2a4172d2a4202da8e0dff7442873db4996f81f36e98296034a420007"></a>

<a id="canonical-53c871f651c96de753520e900a0a3b859c3f344adce91d0a3743ae78e328ac99"></a>

## verbs property — policy_rule_list.policy_rule.resource_list / 770464acba69 / 7

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

<a id="canonical-9cec599bf335153b448999682773b8bfc904731538ca5dad1ae5de186e5d0e62"></a>

## Next pages — policy_rule_list.policy_rule.resource_list / 770464acba69 / 8

- [policy_rule_list.policy_rule](resources--k8s_cluster_role--reference--group-001.md#canonical-cdb9108e75620baabe77f2aeba33f40b4a10fd427f43961a3127c5f58c343efd)
- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-4e8d2df4747cdd8e8e9686cbca47d61ff08b2031a918d8da2f09af43bd55e66b)

<a id="canonical-c2f94201abee4b34a4e313fc47fda030e04828ac4432528e707f9c6a44aac733"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-39909164f103806f955a3f93bda4635d3e9a0560601ad9b6ca34a12a763ee40d"></a>

## timeouts — timeouts / 4421e6a5a438 / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-4e8d2df4747cdd8e8e9686cbca47d61ff08b2031a918d8da2f09af43bd55e66b)
- [Property reference](resources--k8s_cluster_role--reference--group-001.md#canonical-9f8cba3c529fa9bc1e4ba8c8f635e5cd17a64356b4e2c0920e65d8384a68de2c)
- timeouts

<a id="canonical-71cec3491158c11cf3830a42bbb1940b98c6a56ef6b0def1c4b3a6e52d55ba9c"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1dd205d8e915651572e49e141588ff3de70bc50f5605a92adf0ff48a4ba57030"></a>

## Direct properties — timeouts / 4421e6a5a438 / 3

<a id="canonical-2e287eab5780808ae83bc325b112cf9fd1909a2c9e3793ef0b9da83a5d40cf48"></a>

<a id="canonical-00d95a3a31111567f995862f187aade274fba7c6d46ed421f487af609599e72d"></a>

## create property — timeouts / 4421e6a5a438 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0dcf6e92e020271beac3c6b0835212d87b67787bb707ffe1b13ef4e7b84db993"></a>

<a id="canonical-05d70d4809dd0957051cbfcf714496997f177357238f67bb6083b60661b010c0"></a>

## delete property — timeouts / 4421e6a5a438 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-8e1b6698b40941f6731029509475b63ce379613e9de22544ba8e699286af3e7c"></a>

<a id="canonical-4321b7fb714eb1e96d9f662394e127a66b5e2a143c57c5a4bd50d90579261c50"></a>

## read property — timeouts / 4421e6a5a438 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-06926dfdfd1fdf684b2fbbd06b1bbd7f4efceb74c2a985246a6616370bd43514"></a>

<a id="canonical-23ff9faf38a39dd42c44275228f45563d9eb99ac1b632abe69fad3e8d807a76c"></a>

## update property — timeouts / 4421e6a5a438 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-07a263b4d98771032431be6b5ba7a8eea599ef0ad4e2747551c48dbde91a5646"></a>

## Next pages — timeouts / 4421e6a5a438 / 8

- [Property reference](resources--k8s_cluster_role--reference--group-001.md#canonical-9f8cba3c529fa9bc1e4ba8c8f635e5cd17a64356b4e2c0920e65d8384a68de2c)
- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-4e8d2df4747cdd8e8e9686cbca47d61ff08b2031a918d8da2f09af43bd55e66b)
