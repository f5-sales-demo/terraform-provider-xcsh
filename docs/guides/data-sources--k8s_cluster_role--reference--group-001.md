---
page_title: "xcsh_k8s_cluster_role reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_k8s_cluster_role reference."
---

# xcsh_k8s_cluster_role reference

<a id="canonical-81d3736ec3cd7378005627ab8c00194ddb3f8ab7f6499d71f90ba016ef6a2050"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd6aca5edc2f2a11edbfc530c82f3503cb34602543e0c8ad0117383cd205ddf7"></a>

## Property reference — Property reference / 4806aea87813 / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role](../data-sources/k8s_cluster_role.md#canonical-ae676024de71b8439d6e1da589b80d914e1b787ad3c0876d229fcc1a32cbbeda)
- Property reference

<a id="canonical-e3de2d029a2ebb1d24c58f782e08d87053570c1092205e26931982a791a5877a"></a>

## Direct properties — Property reference / 4806aea87813 / 3

<a id="canonical-fef6f2f302fc8c7d4baf0ed8f67ec3d647a803b7e0840633ace52134a4bf8053"></a>

<a id="canonical-7e4c4d68fa77c1ba3d139b6a52e59bb9498f82b2e91d483ccbf2b712a0479deb"></a>

## annotations property — Property reference / 4806aea87813 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

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

<a id="canonical-3ed0a0b14989e556509f01bc085a2cd68113b923f4de16a2864bbdf32a7ea74e"></a>

<a id="canonical-0edbfe79387449b90a46a4a1b05f837b3ab0513b9a94566d932944b4cdf7af4b"></a>

## description property — Property reference / 4806aea87813 / 5

Type: `"string"`. Computed.

Description of the K8SClusterRole.

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

<a id="canonical-96a582c1c0d70037eb4c9e40ee4dde0e8daac2fc7bdc421b2927b28a996e3fce"></a>

<a id="canonical-88297962557ebe85159d87ff19447954514bcf31a5a959797d78fc12e44ec68f"></a>

## id property — Property reference / 4806aea87813 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [k8s_cluster_role_selector](data-sources--k8s_cluster_role--reference--group-001.md#canonical-d529956175e4467bf9e09f948a135e06b43515b468ac304ad3a7360aab8b6274): complete subsection reference.

<a id="canonical-c85952a46cfa7b7838e274d826ae3bb02e76c1d9d87c3967e580a6415b570016"></a>

<a id="canonical-14991dc3539c838bab917583acdea261d08d1012a795cf290822230589acf985"></a>

## labels property — Property reference / 4806aea87813 / 7

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

<a id="canonical-7575efaba1a3266c023695d28887fbaa58a98f85b9189a7e44a95728a4c850f2"></a>

<a id="canonical-285dc678385d81d1851d38cf1de189c93a8de250cd004f092b265c8b70bda324"></a>

## name property — Property reference / 4806aea87813 / 8

Type: `"string"`. Required.

Name of the K8SClusterRole.

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

<a id="canonical-38623f424c9c433476ab5900b20384f83f55c9e9eb8daa7bec865c575587aae0"></a>

<a id="canonical-e326ace2ae46eab545947bcdfb9c987dee6dc7b30233a1c1df2bb3317acaa930"></a>

## namespace property — Property reference / 4806aea87813 / 9

Type: `"string"`. Optional, Computed.

Namespace where the K8SClusterRole exists.

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

- [policy_rule_list](data-sources--k8s_cluster_role--reference--group-001.md#canonical-697f439f40ef84be73b5080bfa6d42b8d89c23322f7dfc2db0276045ccf5ba4b): complete subsection reference.

<a id="canonical-99a0b2a49104f3fa6aa5ee93b44bd7b1184071d9bd59ff14c792420d3572b181"></a>

<a id="canonical-d3cc09cf0dddc56b289e8f49d701eb97c2af8f6945142e54186e3dd96421706e"></a>

## yaml property — Property reference / 4806aea87813 / 10

Type: `"string"`. Computed.

Exclusive with \[k8s\_cluster\_role\_selector policy\_rule\_list\] K8s YAML for ClusterRole.

Upstream description:

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

<a id="canonical-3bbad609d872ea99938c334b902ac20d41cd6f71c9a485bf79a3ac775591eb2f"></a>

## All schema paths — Property reference / 4806aea87813 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--k8s_cluster_role--reference--group-001.md#canonical-fef6f2f302fc8c7d4baf0ed8f67ec3d647a803b7e0840633ace52134a4bf8053) |
| `description` | [description](data-sources--k8s_cluster_role--reference--group-001.md#canonical-3ed0a0b14989e556509f01bc085a2cd68113b923f4de16a2864bbdf32a7ea74e) |
| `id` | [id](data-sources--k8s_cluster_role--reference--group-001.md#canonical-96a582c1c0d70037eb4c9e40ee4dde0e8daac2fc7bdc421b2927b28a996e3fce) |
| `k8s_cluster_role_selector` | [k8s_cluster_role_selector](data-sources--k8s_cluster_role--reference--group-001.md#canonical-5113e79e73365ddb332b8b4c9140f79a77ab08b5b7e1d7d4e8db17e7e6ae161a) |
| `k8s_cluster_role_selector.expressions` | [k8s_cluster_role_selector.expressions](data-sources--k8s_cluster_role--reference--group-001.md#canonical-64dfb56ec375ddd51e0721f245c0a63e7494bc16fc2a50d8f244a89e91b715e4) |
| `labels` | [labels](data-sources--k8s_cluster_role--reference--group-001.md#canonical-c85952a46cfa7b7838e274d826ae3bb02e76c1d9d87c3967e580a6415b570016) |
| `name` | [name](data-sources--k8s_cluster_role--reference--group-001.md#canonical-7575efaba1a3266c023695d28887fbaa58a98f85b9189a7e44a95728a4c850f2) |
| `namespace` | [namespace](data-sources--k8s_cluster_role--reference--group-001.md#canonical-38623f424c9c433476ab5900b20384f83f55c9e9eb8daa7bec865c575587aae0) |
| `policy_rule_list` | [policy_rule_list](data-sources--k8s_cluster_role--reference--group-001.md#canonical-b93aca3c8aaf33f69c2868d459b06de2d4fdc093a9f931fb51aec0f2a231ece5) |
| `policy_rule_list.policy_rule` | [policy_rule_list.policy_rule](data-sources--k8s_cluster_role--reference--group-001.md#canonical-2b910994de4a282fe664b23c28565f2239760b49e0e8d2c7b058bc3c665c4236) |
| `policy_rule_list.policy_rule.non_resource_url_list` | [policy_rule_list.policy_rule.non_resource_url_list](data-sources--k8s_cluster_role--reference--group-001.md#canonical-b48938f7f5496696d7663c9239ab5f8d4c3878b30509c6000c75efd4efc3606a) |
| `policy_rule_list.policy_rule.non_resource_url_list.urls` | [policy_rule_list.policy_rule.non_resource_url_list.urls](data-sources--k8s_cluster_role--reference--group-001.md#canonical-499b8ed240a0e21459668f41fc86348968e83edcbb06fcffb906b01f53e756ec) |
| `policy_rule_list.policy_rule.non_resource_url_list.verbs` | [policy_rule_list.policy_rule.non_resource_url_list.verbs](data-sources--k8s_cluster_role--reference--group-001.md#canonical-c037a8167d3e69c3cf49a4f1fc3064f4f2176855d0c4a85c9e5a1073ba06bdc1) |
| `policy_rule_list.policy_rule.resource_list` | [policy_rule_list.policy_rule.resource_list](data-sources--k8s_cluster_role--reference--group-001.md#canonical-2c7ba8ca70883bb89bae5526e22e081671690ac24d1469ce4837cb57c29e2421) |
| `policy_rule_list.policy_rule.resource_list.api_groups` | [policy_rule_list.policy_rule.resource_list.api_groups](data-sources--k8s_cluster_role--reference--group-001.md#canonical-dc66b2e2199012203cdeeabb7fea5e4bec953ff42f65e040a14628f76e66a13a) |
| `policy_rule_list.policy_rule.resource_list.resource_instances` | [policy_rule_list.policy_rule.resource_list.resource_instances](data-sources--k8s_cluster_role--reference--group-001.md#canonical-701959551cbcd95243cb156ac925a260bf296298ace851464d45a9d8f8ea7f25) |
| `policy_rule_list.policy_rule.resource_list.resource_types` | [policy_rule_list.policy_rule.resource_list.resource_types](data-sources--k8s_cluster_role--reference--group-001.md#canonical-893d9673592ef635dfe43ebfb2be2be7299e428189ecadf2d5b67a01e9b9b070) |
| `policy_rule_list.policy_rule.resource_list.verbs` | [policy_rule_list.policy_rule.resource_list.verbs](data-sources--k8s_cluster_role--reference--group-001.md#canonical-a010cd7c9999e11c6703bcfebe6bf1289681a553702054943e3ed303e4ff50d9) |
| `yaml` | [yaml](data-sources--k8s_cluster_role--reference--group-001.md#canonical-99a0b2a49104f3fa6aa5ee93b44bd7b1184071d9bd59ff14c792420d3572b181) |

<a id="canonical-21d494bcffb499e8af970783f96db712fe2dc400cbdbdd5a7037536627b25586"></a>

## Next pages — Property reference / 4806aea87813 / 12

- [k8s_cluster_role_selector](data-sources--k8s_cluster_role--reference--group-001.md#canonical-d529956175e4467bf9e09f948a135e06b43515b468ac304ad3a7360aab8b6274)
- [policy_rule_list](data-sources--k8s_cluster_role--reference--group-001.md#canonical-697f439f40ef84be73b5080bfa6d42b8d89c23322f7dfc2db0276045ccf5ba4b)
- [xcsh_k8s_cluster_role](../data-sources/k8s_cluster_role.md#canonical-ae676024de71b8439d6e1da589b80d914e1b787ad3c0876d229fcc1a32cbbeda)

<a id="canonical-d529956175e4467bf9e09f948a135e06b43515b468ac304ad3a7360aab8b6274"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-57db36881cf6be1eaed76f364a4bc8e3d55336ff513fec8f0cbd28c4c72c81ec"></a>

## k8s_cluster_role_selector — k8s_cluster_role_selector / a427f6af7b7e / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role](../data-sources/k8s_cluster_role.md#canonical-ae676024de71b8439d6e1da589b80d914e1b787ad3c0876d229fcc1a32cbbeda)
- [Property reference](data-sources--k8s_cluster_role--reference--group-001.md#canonical-81d3736ec3cd7378005627ab8c00194ddb3f8ab7f6499d71f90ba016ef6a2050)
- k8s_cluster_role_selector

<a id="canonical-5113e79e73365ddb332b8b4c9140f79a77ab08b5b7e1d7d4e8db17e7e6ae161a"></a>

Type: `"single"`. Computed.

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

- [k8s_cluster_role_selector](data-sources--k8s_cluster_role--reference--group-001.md#canonical-5113e79e73365ddb332b8b4c9140f79a77ab08b5b7e1d7d4e8db17e7e6ae161a)
- [policy_rule_list](data-sources--k8s_cluster_role--reference--group-001.md#canonical-b93aca3c8aaf33f69c2868d459b06de2d4fdc093a9f931fb51aec0f2a231ece5)
- [yaml](data-sources--k8s_cluster_role--reference--group-001.md#canonical-99a0b2a49104f3fa6aa5ee93b44bd7b1184071d9bd59ff14c792420d3572b181)

Select alternatives according to the provider validators above.

<a id="canonical-0bdafc4b97cecb91ddf876c5cdac3982230704015d2801907abd6f1bd86b84dd"></a>

## Direct properties — k8s_cluster_role_selector / a427f6af7b7e / 3

<a id="canonical-64dfb56ec375ddd51e0721f245c0a63e7494bc16fc2a50d8f244a89e91b715e4"></a>

<a id="canonical-62ee637445f1e350fe82e4c61791642727288d81f5d5b76b55f04a9d1992e9a5"></a>

## expressions property — k8s_cluster_role_selector / a427f6af7b7e / 4

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

<a id="canonical-4a25bb789588837d5de07cd390f253ee81f6ed8648cf03f4c5669c4cda2bdaf9"></a>

## Next pages — k8s_cluster_role_selector / a427f6af7b7e / 5

- [Property reference](data-sources--k8s_cluster_role--reference--group-001.md#canonical-81d3736ec3cd7378005627ab8c00194ddb3f8ab7f6499d71f90ba016ef6a2050)
- [xcsh_k8s_cluster_role](../data-sources/k8s_cluster_role.md#canonical-ae676024de71b8439d6e1da589b80d914e1b787ad3c0876d229fcc1a32cbbeda)

<a id="canonical-697f439f40ef84be73b5080bfa6d42b8d89c23322f7dfc2db0276045ccf5ba4b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-11f7405f5a39d95b327ce456c6470b8a156be17777be93822ea2e6065dd66e17"></a>

## policy_rule_list — policy_rule_list / 67413c9f93b0 / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role](../data-sources/k8s_cluster_role.md#canonical-ae676024de71b8439d6e1da589b80d914e1b787ad3c0876d229fcc1a32cbbeda)
- [Property reference](data-sources--k8s_cluster_role--reference--group-001.md#canonical-81d3736ec3cd7378005627ab8c00194ddb3f8ab7f6499d71f90ba016ef6a2050)
- policy_rule_list

<a id="canonical-b93aca3c8aaf33f69c2868d459b06de2d4fdc093a9f931fb51aec0f2a231ece5"></a>

Type: `"single"`. Computed.

Policy Rule List. List of rules for role permissions.

Upstream description:

List of rules for role permissions.

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

<a id="canonical-940d2b8f3f7d8e250c6013527760e11fddf37079e78321b5573959a60d7086f4"></a>

## Direct properties — policy_rule_list / 67413c9f93b0 / 3

- [policy_rule](data-sources--k8s_cluster_role--reference--group-001.md#canonical-6d6c92834b28476647d7de541816c69c90698429e00b2edd228fbcc29231b485): complete subsection reference.

<a id="canonical-79ee5eb67f17b14964e028e722ceaf053e1ea5b11d21be001213c18f85ef85f5"></a>

## Next pages — policy_rule_list / 67413c9f93b0 / 4

- [policy_rule_list.policy_rule](data-sources--k8s_cluster_role--reference--group-001.md#canonical-6d6c92834b28476647d7de541816c69c90698429e00b2edd228fbcc29231b485)
- [Property reference](data-sources--k8s_cluster_role--reference--group-001.md#canonical-81d3736ec3cd7378005627ab8c00194ddb3f8ab7f6499d71f90ba016ef6a2050)
- [xcsh_k8s_cluster_role](../data-sources/k8s_cluster_role.md#canonical-ae676024de71b8439d6e1da589b80d914e1b787ad3c0876d229fcc1a32cbbeda)

<a id="canonical-6d6c92834b28476647d7de541816c69c90698429e00b2edd228fbcc29231b485"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-570148054a8d2e8b73392ca9070be8d7b0db13efd68432c16a770da380d01234"></a>

## policy_rule_list.policy_rule — policy_rule_list.policy_rule / ec97edef2fca / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role](../data-sources/k8s_cluster_role.md#canonical-ae676024de71b8439d6e1da589b80d914e1b787ad3c0876d229fcc1a32cbbeda)
- [Property reference](data-sources--k8s_cluster_role--reference--group-001.md#canonical-81d3736ec3cd7378005627ab8c00194ddb3f8ab7f6499d71f90ba016ef6a2050)
- [policy_rule_list](data-sources--k8s_cluster_role--reference--group-001.md#canonical-697f439f40ef84be73b5080bfa6d42b8d89c23322f7dfc2db0276045ccf5ba4b)
- policy_rule_list.policy_rule

<a id="canonical-2b910994de4a282fe664b23c28565f2239760b49e0e8d2c7b058bc3c665c4236"></a>

Type: `"list"`. Computed.

Policy Rules. List of rules for role permissions.

Upstream description:

List of rules for role permissions.

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

<a id="canonical-755f5efafde55344d6a14d20c7a3001e9d9e565b52defb763d85267ced4f4b57"></a>

## Direct properties — policy_rule_list.policy_rule / ec97edef2fca / 3

- [non_resource_url_list](data-sources--k8s_cluster_role--reference--group-001.md#canonical-96f79db2e5ec26e579bd4d258ae01dfabeba5c58cfb17923c5d50c82c25ec8e4): complete subsection reference.

- [resource_list](data-sources--k8s_cluster_role--reference--group-001.md#canonical-f6568c1fa57d2b9105ab0fc8a52c0582acaebe3aaf6e1b19f4da2e996d5c2ce9): complete subsection reference.

<a id="canonical-85619d11cf2d5495d85c33eb8f025578aa1fbb5d971719df17a12dd35932419f"></a>

## Next pages — policy_rule_list.policy_rule / ec97edef2fca / 4

- [policy_rule_list.policy_rule.non_resource_url_list](data-sources--k8s_cluster_role--reference--group-001.md#canonical-96f79db2e5ec26e579bd4d258ae01dfabeba5c58cfb17923c5d50c82c25ec8e4)
- [policy_rule_list.policy_rule.resource_list](data-sources--k8s_cluster_role--reference--group-001.md#canonical-f6568c1fa57d2b9105ab0fc8a52c0582acaebe3aaf6e1b19f4da2e996d5c2ce9)
- [policy_rule_list](data-sources--k8s_cluster_role--reference--group-001.md#canonical-697f439f40ef84be73b5080bfa6d42b8d89c23322f7dfc2db0276045ccf5ba4b)
- [xcsh_k8s_cluster_role](../data-sources/k8s_cluster_role.md#canonical-ae676024de71b8439d6e1da589b80d914e1b787ad3c0876d229fcc1a32cbbeda)

<a id="canonical-96f79db2e5ec26e579bd4d258ae01dfabeba5c58cfb17923c5d50c82c25ec8e4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e8808789abf73899e4b84045306d2635e3bf90870c8f0ff3dac13df107332efc"></a>

## policy_rule_list.policy_rule.non_resource_url_list — policy_rule_list.policy_rule.non_resource_url_list / a0e78f9835cc / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role](../data-sources/k8s_cluster_role.md#canonical-ae676024de71b8439d6e1da589b80d914e1b787ad3c0876d229fcc1a32cbbeda)
- [Property reference](data-sources--k8s_cluster_role--reference--group-001.md#canonical-81d3736ec3cd7378005627ab8c00194ddb3f8ab7f6499d71f90ba016ef6a2050)
- [policy_rule_list](data-sources--k8s_cluster_role--reference--group-001.md#canonical-697f439f40ef84be73b5080bfa6d42b8d89c23322f7dfc2db0276045ccf5ba4b)
- [policy_rule_list.policy_rule](data-sources--k8s_cluster_role--reference--group-001.md#canonical-6d6c92834b28476647d7de541816c69c90698429e00b2edd228fbcc29231b485)
- policy_rule_list.policy_rule.non_resource_url_list

<a id="canonical-b48938f7f5496696d7663c9239ab5f8d4c3878b30509c6000c75efd4efc3606a"></a>

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

<a id="canonical-258ff45e5fcc0a23d80a323f20954aaea21e035ac44881f573340d9f575d30a6"></a>

## Direct properties — policy_rule_list.policy_rule.non_resource_url_list / a0e78f9835cc / 3

<a id="canonical-499b8ed240a0e21459668f41fc86348968e83edcbb06fcffb906b01f53e756ec"></a>

<a id="canonical-2728431758e3318a85769c630c118200041a88b57db82bf36e09637683008a5a"></a>

## urls property — policy_rule_list.policy_rule.non_resource_url_list / a0e78f9835cc / 4

Type: `["list", "string"]`. Computed.

Allowed URL(s) that do not represent any K8s resource. URL can be suffix or regex.

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

<a id="canonical-c037a8167d3e69c3cf49a4f1fc3064f4f2176855d0c4a85c9e5a1073ba06bdc1"></a>

<a id="canonical-1b40d892365bac9fb1eb83a4783023a14a85d85b2fdf1282ab57714c74a368df"></a>

## verbs property — policy_rule_list.policy_rule.non_resource_url_list / a0e78f9835cc / 5

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

<a id="canonical-261747f05f3e0c32fb5f8890b74d37492e6459e3b23b84014f4fc7cdafd20f56"></a>

## Next pages — policy_rule_list.policy_rule.non_resource_url_list / a0e78f9835cc / 6

- [policy_rule_list.policy_rule](data-sources--k8s_cluster_role--reference--group-001.md#canonical-6d6c92834b28476647d7de541816c69c90698429e00b2edd228fbcc29231b485)
- [xcsh_k8s_cluster_role](../data-sources/k8s_cluster_role.md#canonical-ae676024de71b8439d6e1da589b80d914e1b787ad3c0876d229fcc1a32cbbeda)

<a id="canonical-f6568c1fa57d2b9105ab0fc8a52c0582acaebe3aaf6e1b19f4da2e996d5c2ce9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d5dc028101e137201dfcab07823de78ca8fda37348f1f5bdbcad7f3e1eab37e7"></a>

## policy_rule_list.policy_rule.resource_list — policy_rule_list.policy_rule.resource_list / 12fa5e4ee632 / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role](../data-sources/k8s_cluster_role.md#canonical-ae676024de71b8439d6e1da589b80d914e1b787ad3c0876d229fcc1a32cbbeda)
- [Property reference](data-sources--k8s_cluster_role--reference--group-001.md#canonical-81d3736ec3cd7378005627ab8c00194ddb3f8ab7f6499d71f90ba016ef6a2050)
- [policy_rule_list](data-sources--k8s_cluster_role--reference--group-001.md#canonical-697f439f40ef84be73b5080bfa6d42b8d89c23322f7dfc2db0276045ccf5ba4b)
- [policy_rule_list.policy_rule](data-sources--k8s_cluster_role--reference--group-001.md#canonical-6d6c92834b28476647d7de541816c69c90698429e00b2edd228fbcc29231b485)
- policy_rule_list.policy_rule.resource_list

<a id="canonical-2c7ba8ca70883bb89bae5526e22e081671690ac24d1469ce4837cb57c29e2421"></a>

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

<a id="canonical-f687515dea1d3ee7391a16dcf801c2ff4eab49a635138b405b993766256de083"></a>

## Direct properties — policy_rule_list.policy_rule.resource_list / 12fa5e4ee632 / 3

<a id="canonical-dc66b2e2199012203cdeeabb7fea5e4bec953ff42f65e040a14628f76e66a13a"></a>

<a id="canonical-f6ff040387d67bcaabe51eee0321aa917014cca260a8f752db53108eafa0c920"></a>

## api_groups property — policy_rule_list.policy_rule.resource_list / 12fa5e4ee632 / 4

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

<a id="canonical-701959551cbcd95243cb156ac925a260bf296298ace851464d45a9d8f8ea7f25"></a>

<a id="canonical-bb401998d561fded8e33ac89b77a19760fd0a1ccb9422daf375a0292c7cd2489"></a>

## resource_instances property — policy_rule_list.policy_rule.resource_list / 12fa5e4ee632 / 5

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

<a id="canonical-893d9673592ef635dfe43ebfb2be2be7299e428189ecadf2d5b67a01e9b9b070"></a>

<a id="canonical-f5c49b14587718f3cb9fbb13174e07d74e9cd1b423804e420bd4fcdd19b3921b"></a>

## resource_types property — policy_rule_list.policy_rule.resource_list / 12fa5e4ee632 / 6

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

<a id="canonical-a010cd7c9999e11c6703bcfebe6bf1289681a553702054943e3ed303e4ff50d9"></a>

<a id="canonical-3a0bec75f957cc028d4d2671c4285dba404d10925809fa0f3b4a810042b6b516"></a>

## verbs property — policy_rule_list.policy_rule.resource_list / 12fa5e4ee632 / 7

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

<a id="canonical-bd1cd482c4d243d4dd6665077b30b62de942eb73345f04886c5dc647e8aeca0f"></a>

## Next pages — policy_rule_list.policy_rule.resource_list / 12fa5e4ee632 / 8

- [policy_rule_list.policy_rule](data-sources--k8s_cluster_role--reference--group-001.md#canonical-6d6c92834b28476647d7de541816c69c90698429e00b2edd228fbcc29231b485)
- [xcsh_k8s_cluster_role](../data-sources/k8s_cluster_role.md#canonical-ae676024de71b8439d6e1da589b80d914e1b787ad3c0876d229fcc1a32cbbeda)
