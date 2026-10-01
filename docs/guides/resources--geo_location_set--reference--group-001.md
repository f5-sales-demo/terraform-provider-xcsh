---
page_title: "xcsh_geo_location_set reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_geo_location_set reference."
---

# xcsh_geo_location_set reference

<a id="canonical-5b59e1bdb82766035879753e3dec625b78e3fc0781f7694c35dffba49ee9e586"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6fd6cf433943074166ae0cc79d4d54c970ef18415806998f733f4e929f25557f"></a>

## Property reference — Property reference / 4e9ef71831e6 / 2

Breadcrumbs:

- [xcsh_geo_location_set](../resources/geo_location_set.md#canonical-d316af30af83a2e4202cd927edfa772e4725c3cf87f11306f130cc11f9ddcd50)
- Property reference

<a id="canonical-ebe92468a0035e7f1dd84db4a8d7e23591f7401e2f69c7eb07e4f1920e9f59d8"></a>

## Direct properties — Property reference / 4e9ef71831e6 / 3

<a id="canonical-f248e5be8c5a3d6b5aeb0c4e67dc55c3b00c8ab9d3e380e6f45c7643b971bec7"></a>

<a id="canonical-6d2b8daecaec73b368a8d86eb7c3c93fa99b93c4b541a87f1853ed7721fe9d9b"></a>

## annotations property — Property reference / 4e9ef71831e6 / 4

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

- [custom_geo_location_selector](resources--geo_location_set--reference--group-001.md#canonical-760a79261e7498780508c8f5d4eded5b3678ebe5b0258c356b925a11a28fa949): complete subsection reference.

<a id="canonical-27cd81f2384cfdab1d08a9c897a2c3cb0a8c7c58a3bf67abbd6a00924a96eb3f"></a>

<a id="canonical-5cffc242c1d59ad4690fc2ba3e87d55bf46dc1553d5e4d7a281d05254fc9446e"></a>

## description property — Property reference / 4e9ef71831e6 / 5

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

<a id="canonical-1994108a7fea7faf0e219556d404c033db9bc166c1d3ea109049935ea582da7e"></a>

<a id="canonical-86c2deefda61fee9d1f7dacb15468d79a6bab7888f07cad024a4a8065349f7d8"></a>

## disable property — Property reference / 4e9ef71831e6 / 6

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

- [global](resources--geo_location_set--reference--group-001.md#canonical-f98c54e2132bf6857adeb1275480cfc953b8114201fb444fdea7aabd9a38ee42): complete subsection reference.

<a id="canonical-4a8f38b8ae3d48d7a343fc853880ba02538c050f829ceef93afdb2fd4f2d57da"></a>

<a id="canonical-d43ff0da3da761737e8782d10f4506165073797ed3c21a7c38601f62121a8bee"></a>

## id property — Property reference / 4e9ef71831e6 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-d028252146f1fa4e10d852db30ebdfd34a0f095aba8149167d63437227516c9c"></a>

<a id="canonical-ff8e0522a21bb93fd8c4f1ea250aec753df93187390849e3efb0b5c9f79c897b"></a>

## labels property — Property reference / 4e9ef71831e6 / 8

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

<a id="canonical-15f94800bdb808531e5e8ca0162a56755abe3af7dadf4c7fbd9f6c8bf7fd76bb"></a>

<a id="canonical-f84a6f01af408142c24aec609d704970b8ed8f096c1d56607a039592351bd53d"></a>

## name property — Property reference / 4e9ef71831e6 / 9

Type: `"string"`. Required.

Name of the Geo Location Set. Must be unique within the namespace.

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

<a id="canonical-4b3073b67882886298a6ad3348f1d35979f3db2112bea8b252c71b5e41c87dac"></a>

<a id="canonical-998bf5572b9ad418b03d106af40423780d6cdd8a92129671b8b0b2a6d7ad827c"></a>

## namespace property — Property reference / 4e9ef71831e6 / 10

Type: `"string"`. Optional, Computed.

Namespace for the Geo Location Set. The F5 XC API restricts this resource to the system namespace;
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

- [timeouts](resources--geo_location_set--reference--group-001.md#canonical-dde12e8c69d27a892a740ed38f709415049961f8b2b5a5ed8fe65fc3c73869da): complete subsection reference.

<a id="canonical-cc5d016f28b249f3ea694c81c9ccde52c39e7253c0c111bfa1c1222957dc2f6b"></a>

## All schema paths — Property reference / 4e9ef71831e6 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--geo_location_set--reference--group-001.md#canonical-f248e5be8c5a3d6b5aeb0c4e67dc55c3b00c8ab9d3e380e6f45c7643b971bec7) |
| `custom_geo_location_selector` | [custom_geo_location_selector](resources--geo_location_set--reference--group-001.md#canonical-67ea44891e194d32270d130fec22fdcc3f78f978dc0f195e043c5f0102968dda) |
| `custom_geo_location_selector.expressions` | [custom_geo_location_selector.expressions](resources--geo_location_set--reference--group-001.md#canonical-f5e044e9c8a2a672fbe053fd51c7b81fd2bab574fb823fcf1cb5c4d1b1d3a912) |
| `description` | [description](resources--geo_location_set--reference--group-001.md#canonical-27cd81f2384cfdab1d08a9c897a2c3cb0a8c7c58a3bf67abbd6a00924a96eb3f) |
| `disable` | [disable](resources--geo_location_set--reference--group-001.md#canonical-1994108a7fea7faf0e219556d404c033db9bc166c1d3ea109049935ea582da7e) |
| `global` | [global](resources--geo_location_set--reference--group-001.md#canonical-d7ad3b462dadde93d111ad93f1b2bf2c7c60fd99170c82ffde76c9c03eba95a7) |
| `id` | [id](resources--geo_location_set--reference--group-001.md#canonical-4a8f38b8ae3d48d7a343fc853880ba02538c050f829ceef93afdb2fd4f2d57da) |
| `labels` | [labels](resources--geo_location_set--reference--group-001.md#canonical-d028252146f1fa4e10d852db30ebdfd34a0f095aba8149167d63437227516c9c) |
| `name` | [name](resources--geo_location_set--reference--group-001.md#canonical-15f94800bdb808531e5e8ca0162a56755abe3af7dadf4c7fbd9f6c8bf7fd76bb) |
| `namespace` | [namespace](resources--geo_location_set--reference--group-001.md#canonical-4b3073b67882886298a6ad3348f1d35979f3db2112bea8b252c71b5e41c87dac) |
| `timeouts` | [timeouts](resources--geo_location_set--reference--group-001.md#canonical-1d089c0c9d56808c40a7a6f850eee2e05a5f6c858c391c542e7fb928e73aedff) |
| `timeouts.create` | [timeouts.create](resources--geo_location_set--reference--group-001.md#canonical-a8999c6626989ada16eb00d357124b8a6d47d0e08d4593a214699434b4df694a) |
| `timeouts.delete` | [timeouts.delete](resources--geo_location_set--reference--group-001.md#canonical-b408f821cebee156ab1f6953302e879c64b91608b9789da8e6ce04dbbfb57d18) |
| `timeouts.read` | [timeouts.read](resources--geo_location_set--reference--group-001.md#canonical-5dc0b539e557f406d88271d436c783165a2bf6c0317a92f756f9c6d93a74f673) |
| `timeouts.update` | [timeouts.update](resources--geo_location_set--reference--group-001.md#canonical-0e30f9449b96dd54920917de380d452c62c82a27f167b293462eb66c3433d7ef) |

<a id="canonical-992be7e58a95d662f57dbfaa7aae0d32b49e7a2214260c41da1b2fd01ef3c0ed"></a>

## Next pages — Property reference / 4e9ef71831e6 / 12

- [custom_geo_location_selector](resources--geo_location_set--reference--group-001.md#canonical-760a79261e7498780508c8f5d4eded5b3678ebe5b0258c356b925a11a28fa949)
- [global](resources--geo_location_set--reference--group-001.md#canonical-f98c54e2132bf6857adeb1275480cfc953b8114201fb444fdea7aabd9a38ee42)
- [timeouts](resources--geo_location_set--reference--group-001.md#canonical-dde12e8c69d27a892a740ed38f709415049961f8b2b5a5ed8fe65fc3c73869da)
- [xcsh_geo_location_set](../resources/geo_location_set.md#canonical-d316af30af83a2e4202cd927edfa772e4725c3cf87f11306f130cc11f9ddcd50)

<a id="canonical-760a79261e7498780508c8f5d4eded5b3678ebe5b0258c356b925a11a28fa949"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-705b354343749b0044c70bb6a4b2b06374ec73d4811ecfaa8150f677cf7e33e7"></a>

## custom_geo_location_selector — custom_geo_location_selector / 203c765396be / 2

Breadcrumbs:

- [xcsh_geo_location_set](../resources/geo_location_set.md#canonical-d316af30af83a2e4202cd927edfa772e4725c3cf87f11306f130cc11f9ddcd50)
- [Property reference](resources--geo_location_set--reference--group-001.md#canonical-5b59e1bdb82766035879753e3dec625b78e3fc0781f7694c35dffba49ee9e586)
- custom_geo_location_selector

<a id="canonical-67ea44891e194d32270d130fec22fdcc3f78f978dc0f195e043c5f0102968dda"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: custom\_geo\_location\_selector, global\] Type can be used to establish a 'selector
reference' from one object(called selector) to a set of other objects(called selectees) based on the
value of expressions. A label selector is a label query over a set of resources. An empty label
selector matches all objects.

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

- [custom_geo_location_selector](resources--geo_location_set--reference--group-001.md#canonical-67ea44891e194d32270d130fec22fdcc3f78f978dc0f195e043c5f0102968dda)
- [global](resources--geo_location_set--reference--group-001.md#canonical-d7ad3b462dadde93d111ad93f1b2bf2c7c60fd99170c82ffde76c9c03eba95a7)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
custom_geo_location_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-0e12eeac07f2330c787ff044a2fa54b45557b4ab9b441d2419222d8b2c229b87"></a>

## Direct properties — custom_geo_location_selector / 203c765396be / 3

<a id="canonical-f5e044e9c8a2a672fbe053fd51c7b81fd2bab574fb823fcf1cb5c4d1b1d3a912"></a>

<a id="canonical-e8ae657227b441aa6b2e90ae39c65e6147a77bec09fc3fbe18ea8c0485a47e5b"></a>

## expressions property — custom_geo_location_selector / 203c765396be / 4

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

<a id="canonical-635c22ecdca5053e2d9accb31afc0dbc08fea74ac1312caf2395c9c4ac558f58"></a>

## Next pages — custom_geo_location_selector / 203c765396be / 5

- [Property reference](resources--geo_location_set--reference--group-001.md#canonical-5b59e1bdb82766035879753e3dec625b78e3fc0781f7694c35dffba49ee9e586)
- [xcsh_geo_location_set](../resources/geo_location_set.md#canonical-d316af30af83a2e4202cd927edfa772e4725c3cf87f11306f130cc11f9ddcd50)

<a id="canonical-f98c54e2132bf6857adeb1275480cfc953b8114201fb444fdea7aabd9a38ee42"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8a3eaafc9dd37fb973f2139d5b5144afd26c801be2eb8785c5fb0e466e302f97"></a>

## global — global / 2b8edfffd3f1 / 2

Breadcrumbs:

- [xcsh_geo_location_set](../resources/geo_location_set.md#canonical-d316af30af83a2e4202cd927edfa772e4725c3cf87f11306f130cc11f9ddcd50)
- [Property reference](resources--geo_location_set--reference--group-001.md#canonical-5b59e1bdb82766035879753e3dec625b78e3fc0781f7694c35dffba49ee9e586)
- global

<a id="canonical-d7ad3b462dadde93d111ad93f1b2bf2c7c60fd99170c82ffde76c9c03eba95a7"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
global = {}
```

<a id="canonical-1f9fafd694b7fa720074a52af00a48bb71cd8253ec3f13e7d9993e33e551a83a"></a>

## Direct properties — global / 2b8edfffd3f1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-22df3846e262bd8f24fe791018cd248793b219d420cc590378a171c2d47ed197"></a>

## Next pages — global / 2b8edfffd3f1 / 4

- [Property reference](resources--geo_location_set--reference--group-001.md#canonical-5b59e1bdb82766035879753e3dec625b78e3fc0781f7694c35dffba49ee9e586)
- [xcsh_geo_location_set](../resources/geo_location_set.md#canonical-d316af30af83a2e4202cd927edfa772e4725c3cf87f11306f130cc11f9ddcd50)

<a id="canonical-dde12e8c69d27a892a740ed38f709415049961f8b2b5a5ed8fe65fc3c73869da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cbfceabc3e86ee85b24892d0c088b2ffddd3a489f537ab219a8064a85ac18ecb"></a>

## timeouts — timeouts / bced82b94206 / 2

Breadcrumbs:

- [xcsh_geo_location_set](../resources/geo_location_set.md#canonical-d316af30af83a2e4202cd927edfa772e4725c3cf87f11306f130cc11f9ddcd50)
- [Property reference](resources--geo_location_set--reference--group-001.md#canonical-5b59e1bdb82766035879753e3dec625b78e3fc0781f7694c35dffba49ee9e586)
- timeouts

<a id="canonical-1d089c0c9d56808c40a7a6f850eee2e05a5f6c858c391c542e7fb928e73aedff"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2b3a47ca977e8f271ac7ea990a82296955837bb85643df05f4b519d86a0a7b87"></a>

## Direct properties — timeouts / bced82b94206 / 3

<a id="canonical-a8999c6626989ada16eb00d357124b8a6d47d0e08d4593a214699434b4df694a"></a>

<a id="canonical-a7c63c82f93393dea199e4f6a6ec027712d1a0fd2a5283ad257938980669a3d8"></a>

## create property — timeouts / bced82b94206 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-b408f821cebee156ab1f6953302e879c64b91608b9789da8e6ce04dbbfb57d18"></a>

<a id="canonical-b9e3da924698d4592c818b725a46813f9bbb5c8fb15f8d336478f9da0192bf62"></a>

## delete property — timeouts / bced82b94206 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-5dc0b539e557f406d88271d436c783165a2bf6c0317a92f756f9c6d93a74f673"></a>

<a id="canonical-5fb7d7c7032e7dff07cbf4efdba9cd15031c4b98fe714331c688e6ed5c7690f4"></a>

## read property — timeouts / bced82b94206 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0e30f9449b96dd54920917de380d452c62c82a27f167b293462eb66c3433d7ef"></a>

<a id="canonical-fa5cd8b3e986d4811fb3a38d7ff330d820dc614586f248da52227ca338e4e9c1"></a>

## update property — timeouts / bced82b94206 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-91a15843f38bda627ddd1a6ebffb170e126047329680d414325b066ba66c41af"></a>

## Next pages — timeouts / bced82b94206 / 8

- [Property reference](resources--geo_location_set--reference--group-001.md#canonical-5b59e1bdb82766035879753e3dec625b78e3fc0781f7694c35dffba49ee9e586)
- [xcsh_geo_location_set](../resources/geo_location_set.md#canonical-d316af30af83a2e4202cd927edfa772e4725c3cf87f11306f130cc11f9ddcd50)
