---
page_title: "xcsh_virtual_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_virtual_site reference."
---

# xcsh_virtual_site reference

<a id="canonical-c6e1c58e8ed5d2e1ac06a6cd0b224843ecf7d102b47216aca5fb83c491698b76"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e8d90413a297fce6474b94961503d27eb67a6c0bd6481b0dc25b46e584d21c13"></a>

## Property reference — Property reference / 5e820d19125c / 2

Breadcrumbs:

- [xcsh_virtual_site](../resources/virtual_site.md#canonical-09806f2f14447d14dc58b9c131307be132d72f5eda3348f98adff903bfddcc45)
- Property reference

<a id="canonical-aee7fb3de1ccb9d8b8984c924a09236a35bbea19bc70c82b3fa36061e9b418b1"></a>

## Direct properties — Property reference / 5e820d19125c / 3

<a id="canonical-13f238dfc6715103940c2897d576735103efb752f565a3fb7d89292bd4a73850"></a>

<a id="canonical-bf00b6b1c75f3683b10920d02b80f72af90f15f0f20b2a7dbf5d3548e3883fa2"></a>

## annotations property — Property reference / 5e820d19125c / 4

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

<a id="canonical-384876ee148a493db65760be5e69dccb87312e163c0bbbfa9bd184e7773ae1ba"></a>

<a id="canonical-2d2103cd3e8ed535c74952b887e8cdb553390a65e1d2d7aaed5bd8f07328fabb"></a>

## description property — Property reference / 5e820d19125c / 5

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

<a id="canonical-fb9b8b6e6c588f66146ce9049207e143c433c0c5e50253132056e2651b75054c"></a>

<a id="canonical-d61aff27599f62694b883373adf9e7c140105e988c99e681472e51d87f9af615"></a>

## disable property — Property reference / 5e820d19125c / 6

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

<a id="canonical-464e57daf78316bcba8de530e3503361173db48fce9a24219e55646dc6e40858"></a>

<a id="canonical-565834b88e5e33724237ad63a088ac538ad1a28f76a0ef83533bc45e321c018b"></a>

## id property — Property reference / 5e820d19125c / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1800e4d3921fb4c5be0789e7b3cf923acc090a3912be58e7fa3ee29ff3922ba5"></a>

<a id="canonical-57c2e8647db8f1f1a595a64bf64253dd057279fe6b51836b1732004bde3c9107"></a>

## labels property — Property reference / 5e820d19125c / 8

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

<a id="canonical-2f99f978bc18f7339d3f2408934f697b7dcd0d8c3c6e8eab86c1977738d6c29b"></a>

<a id="canonical-b088c8921144d63ef6c9e8204c033238e4af9b918a20d1fc94b54fec1709bd5c"></a>

## name property — Property reference / 5e820d19125c / 9

Type: `"string"`. Required.

Name of the Virtual Site. Must be unique within the namespace.

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

<a id="canonical-675568eefff2cec65d2aacb667f5d2eadc0a70cfb285349e1d13a606dcd4a42e"></a>

<a id="canonical-6e9a8e95c86f57dda54055b25aaa1dcb33c53149e35a5448701e6c4e084e8c0e"></a>

## namespace property — Property reference / 5e820d19125c / 10

Type: `"string"`. Required.

Namespace where the Virtual Site is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
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

- [site_selector](resources--virtual_site--reference--group-001.md#canonical-ab6c3726b2a72354522e8c17f04bb89ae26f46201585529ba8368e26e9cf8ca1): complete subsection reference.

<a id="canonical-9b43d583f9e625407d780edc68d665e4c88d2940986f8417be8d95b66edd1f55"></a>

<a id="canonical-a2a38c50b7be9186484d975f967d65c2a3f3cc4e135e3e23ebccb7ebdcf1fbf7"></a>

## site_type property — Property reference / 5e820d19125c / 11

Type: `"string"`. Optional, Computed.

\[Enum: INVALID|REGIONAL\_EDGE|CUSTOMER\_EDGE|NGINX\_ONE\] Site Type which can either RE or CE
Invalid type of site Regional Edge site Customer Edge site. Possible values are \`INVALID\`,
\`REGIONAL\_EDGE\`, \`CUSTOMER\_EDGE\`, \`NGINX\_ONE\`.

Upstream description:

Site Type which can either RE or CE

Invalid type of site Regional Edge site Customer Edge site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("INVALID",
    "REGIONAL_EDGE",
    "CUSTOMER_EDGE",
    "NGINX_ONE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "INVALID",
  "enum": [
    "INVALID",
    "REGIONAL_EDGE",
    "CUSTOMER_EDGE",
    "NGINX_ONE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [timeouts](resources--virtual_site--reference--group-001.md#canonical-9167ad801f3ee23e14a589894bf0170759e1f8fdb6df4aaac6ae9b5484fe64ff): complete subsection reference.

<a id="canonical-43f9abfeb346611561560e83d57ade2d1ca53c2b38ce6465aeb55421196232bd"></a>

## All schema paths — Property reference / 5e820d19125c / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--virtual_site--reference--group-001.md#canonical-13f238dfc6715103940c2897d576735103efb752f565a3fb7d89292bd4a73850) |
| `description` | [description](resources--virtual_site--reference--group-001.md#canonical-384876ee148a493db65760be5e69dccb87312e163c0bbbfa9bd184e7773ae1ba) |
| `disable` | [disable](resources--virtual_site--reference--group-001.md#canonical-fb9b8b6e6c588f66146ce9049207e143c433c0c5e50253132056e2651b75054c) |
| `id` | [id](resources--virtual_site--reference--group-001.md#canonical-464e57daf78316bcba8de530e3503361173db48fce9a24219e55646dc6e40858) |
| `labels` | [labels](resources--virtual_site--reference--group-001.md#canonical-1800e4d3921fb4c5be0789e7b3cf923acc090a3912be58e7fa3ee29ff3922ba5) |
| `name` | [name](resources--virtual_site--reference--group-001.md#canonical-2f99f978bc18f7339d3f2408934f697b7dcd0d8c3c6e8eab86c1977738d6c29b) |
| `namespace` | [namespace](resources--virtual_site--reference--group-001.md#canonical-675568eefff2cec65d2aacb667f5d2eadc0a70cfb285349e1d13a606dcd4a42e) |
| `site_selector` | [site_selector](resources--virtual_site--reference--group-001.md#canonical-8597583d7e6b6bf121f7a9fa4c01c4e7bf3b680fcaa8654266997f5b93b24e72) |
| `site_selector.expressions` | [site_selector.expressions](resources--virtual_site--reference--group-001.md#canonical-f63edff3b9b0a456504abd245b4afe006b255b585cf163590009cd6fdddb6436) |
| `site_type` | [site_type](resources--virtual_site--reference--group-001.md#canonical-9b43d583f9e625407d780edc68d665e4c88d2940986f8417be8d95b66edd1f55) |
| `timeouts` | [timeouts](resources--virtual_site--reference--group-001.md#canonical-a42ac1c22f129aabd0d8331d9ec7e0c0eb2667087648a63982e92675a9b5f213) |
| `timeouts.create` | [timeouts.create](resources--virtual_site--reference--group-001.md#canonical-c831f11e655e30986d1e9af1e736babbeda73beed9d9b38b1ce63fccf8f38d85) |
| `timeouts.delete` | [timeouts.delete](resources--virtual_site--reference--group-001.md#canonical-6e49a2ff805eaa2d1f7ad827dcc561e0038a21654608f9b8a294da8bb7920e1e) |
| `timeouts.read` | [timeouts.read](resources--virtual_site--reference--group-001.md#canonical-29da4fbcced0cfb293e31c611a8ef0eeac9769c5601c4c227e8898d955599017) |
| `timeouts.update` | [timeouts.update](resources--virtual_site--reference--group-001.md#canonical-b9a124d89f37e83f614dbb2d3ede80c546f10e6a6eb46ed7c21ef2a0253a540d) |

<a id="canonical-0c11c0498af60557585449a089ff4c4e14bde5c0a2281c732d3c975e791344d6"></a>

## Next pages — Property reference / 5e820d19125c / 13

- [site_selector](resources--virtual_site--reference--group-001.md#canonical-ab6c3726b2a72354522e8c17f04bb89ae26f46201585529ba8368e26e9cf8ca1)
- [timeouts](resources--virtual_site--reference--group-001.md#canonical-9167ad801f3ee23e14a589894bf0170759e1f8fdb6df4aaac6ae9b5484fe64ff)
- [xcsh_virtual_site](../resources/virtual_site.md#canonical-09806f2f14447d14dc58b9c131307be132d72f5eda3348f98adff903bfddcc45)

<a id="canonical-ab6c3726b2a72354522e8c17f04bb89ae26f46201585529ba8368e26e9cf8ca1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d5379d59f20ca05ff7bd9215ce9c769139bbad1d795cfaa7496139b995fb280d"></a>

## site_selector — site_selector / 6d44efe44779 / 2

Breadcrumbs:

- [xcsh_virtual_site](../resources/virtual_site.md#canonical-09806f2f14447d14dc58b9c131307be132d72f5eda3348f98adff903bfddcc45)
- [Property reference](resources--virtual_site--reference--group-001.md#canonical-c6e1c58e8ed5d2e1ac06a6cd0b224843ecf7d102b47216aca5fb83c491698b76)
- site_selector

<a id="canonical-8597583d7e6b6bf121f7a9fa4c01c4e7bf3b680fcaa8654266997f5b93b24e72"></a>

Type: `"object"`. single nested block, Optional.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

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

Terraform syntax:

```terraform
site_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-5b8dbf6aa1e482c33225c79aa136273ee1fa2a9bd3ab7e86add8b164459cf85c"></a>

## Direct properties — site_selector / 6d44efe44779 / 3

<a id="canonical-f63edff3b9b0a456504abd245b4afe006b255b585cf163590009cd6fdddb6436"></a>

<a id="canonical-51ba592178c787e780d440c59391beb81d9fe4664269ab56fb36c3efac81c46e"></a>

## expressions property — site_selector / 6d44efe44779 / 4

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

<a id="canonical-25956354c4364542c92d0498daa1565365f3f00ecd815592e2ade4d360bbe40a"></a>

## Next pages — site_selector / 6d44efe44779 / 5

- [Property reference](resources--virtual_site--reference--group-001.md#canonical-c6e1c58e8ed5d2e1ac06a6cd0b224843ecf7d102b47216aca5fb83c491698b76)
- [xcsh_virtual_site](../resources/virtual_site.md#canonical-09806f2f14447d14dc58b9c131307be132d72f5eda3348f98adff903bfddcc45)

<a id="canonical-9167ad801f3ee23e14a589894bf0170759e1f8fdb6df4aaac6ae9b5484fe64ff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b98a728c33fa635f50e42f84e4d279336b9aa7d6f38294678dddf9bd562a75f"></a>

## timeouts — timeouts / d493cf98fb7d / 2

Breadcrumbs:

- [xcsh_virtual_site](../resources/virtual_site.md#canonical-09806f2f14447d14dc58b9c131307be132d72f5eda3348f98adff903bfddcc45)
- [Property reference](resources--virtual_site--reference--group-001.md#canonical-c6e1c58e8ed5d2e1ac06a6cd0b224843ecf7d102b47216aca5fb83c491698b76)
- timeouts

<a id="canonical-a42ac1c22f129aabd0d8331d9ec7e0c0eb2667087648a63982e92675a9b5f213"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-f4e4ec3f7b1a5ea8e53990fd101bc99f2ae2ef665e7bdc9caaa7d1ba8df98ca2"></a>

## Direct properties — timeouts / d493cf98fb7d / 3

<a id="canonical-c831f11e655e30986d1e9af1e736babbeda73beed9d9b38b1ce63fccf8f38d85"></a>

<a id="canonical-1acbbcdc6f23577776c92ee6d75d03a75db00c1e613c41da6c7b647b77749f1b"></a>

## create property — timeouts / d493cf98fb7d / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-6e49a2ff805eaa2d1f7ad827dcc561e0038a21654608f9b8a294da8bb7920e1e"></a>

<a id="canonical-d7ebe10ccf737b2d9805f53689b092444e3b5a4cbe75dd4dbfec93939069b7e0"></a>

## delete property — timeouts / d493cf98fb7d / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-29da4fbcced0cfb293e31c611a8ef0eeac9769c5601c4c227e8898d955599017"></a>

<a id="canonical-9ae0aa502efd5adc4ea1b5adca09363ec89d6fded5a381b01e6cafc38a543f56"></a>

## read property — timeouts / d493cf98fb7d / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-b9a124d89f37e83f614dbb2d3ede80c546f10e6a6eb46ed7c21ef2a0253a540d"></a>

<a id="canonical-05aa826e4236110b70e8b2d6922703620aa1fe92f5493652bbcd9b8bf0a38d44"></a>

## update property — timeouts / d493cf98fb7d / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-01bf2d7779dd7b45af6bdf27a44221fc3dbcaf75b5a268ab68e709ff862f3201"></a>

## Next pages — timeouts / d493cf98fb7d / 8

- [Property reference](resources--virtual_site--reference--group-001.md#canonical-c6e1c58e8ed5d2e1ac06a6cd0b224843ecf7d102b47216aca5fb83c491698b76)
- [xcsh_virtual_site](../resources/virtual_site.md#canonical-09806f2f14447d14dc58b9c131307be132d72f5eda3348f98adff903bfddcc45)
