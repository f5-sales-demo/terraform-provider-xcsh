---
page_title: "xcsh_virtual_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_virtual_site reference."
---

# xcsh_virtual_site reference

<a id="canonical-011be626f96641414cc8e65665477f0a7768f907c5427772ed92f6dd2c667a74"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f49b55c14d9e3febd663b8fc719300608090be5f8b98aafd4b778e9e8edb9a6"></a>

## Property reference — Property reference / 42b1cccd4164 / 2

Breadcrumbs:

- [xcsh_virtual_site](../data-sources/virtual_site.md#canonical-4d9c6e801f11ee0a4dcb7c662b2cdc5dffca2ee94e25d78ffb25ff05789e2980)
- Property reference

<a id="canonical-0b7ddb08f6ea4c1d999c6ecf139346c20e5c81a2558f9ac4aed319d51ed3cae2"></a>

## Direct properties — Property reference / 42b1cccd4164 / 3

<a id="canonical-3f6217e5f02db027ec8dc31b8712c71e8ddad0cb5af19a172d86e766833c596e"></a>

<a id="canonical-146f6cc98b5665542ee6c7f8566588adc8c4096a3df7afe50f9a4ab4ca209a71"></a>

## annotations property — Property reference / 42b1cccd4164 / 4

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

<a id="canonical-085676c61c64d1e6b28c8fa9fdbdd0b650fbc564962049f250b6e3589a55ec41"></a>

<a id="canonical-e7fc2b9088329a9aba8854dfb85a5599963f533e3a020079b5c9512533c4a10f"></a>

## description property — Property reference / 42b1cccd4164 / 5

Type: `"string"`. Computed.

Description of the VirtualSite.

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

<a id="canonical-237876726eb161a1d509491fe2a76eb4bf20f1cfe17ac0fdb9241446fc562b75"></a>

<a id="canonical-0d0b39a1056b78055a62f9126e3c494d7c67804c628299a970a899c2622cca1b"></a>

## id property — Property reference / 42b1cccd4164 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-224ea99f85a2df475b8dfe4a66abf8de2fe2bbd82e2c8e283f555143576f3618"></a>

<a id="canonical-58f2ef2ae34d7b8d1ec484993ff212a6423e97047d738ea727cf96ff59ab0dff"></a>

## labels property — Property reference / 42b1cccd4164 / 7

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

<a id="canonical-215a7651731fca9b85c31c19999d9c5b2754f6d6d6b6b8830a739f63d7b45690"></a>

<a id="canonical-e2b4a5b7c3bcda8913688623599c7bb392afa0f7a87cfaa0575fc86dbe1d5504"></a>

## name property — Property reference / 42b1cccd4164 / 8

Type: `"string"`. Required.

Name of the VirtualSite.

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

<a id="canonical-68ee4c82ca41308cb32b26fe4781bdb2a915d12fcbeea25b90a89851ce93b0c4"></a>

<a id="canonical-692e3f7e351784c0e922ed50edec390ec4e0eeaab168e0c399d16c7cad4961be"></a>

## namespace property — Property reference / 42b1cccd4164 / 9

Type: `"string"`. Required.

Namespace where the VirtualSite exists.

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

- [site_selector](data-sources--virtual_site--reference--group-001.md#canonical-b14df6e4b7249e796fe0b5d80e90e25055f2a4beabfc0f960bea5e5e60a1c6c3): complete subsection reference.

<a id="canonical-9fe59744e7e192c572b177d48aa59b40e1ffa5345e650969a6de65e95862e654"></a>

<a id="canonical-86ab28b3ac4c4466d616cbb0491dfdcf9a2dbf09a23689ed181874abcd92f36d"></a>

## site_type property — Property reference / 42b1cccd4164 / 10

Type: `"string"`. Computed.

\[Enum: INVALID|REGIONAL\_EDGE|CUSTOMER\_EDGE|NGINX\_ONE\] Site Type which can either RE or CE
Invalid type of site Regional Edge site Customer Edge site. Possible values are \`INVALID\`,
\`REGIONAL\_EDGE\`, \`CUSTOMER\_EDGE\`, \`NGINX\_ONE\`.

Upstream description:

Site Type which can either RE or CE

Invalid type of site Regional Edge site Customer Edge site.

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

<a id="canonical-10fca225d085a17d70852c31e5d8c99da850c7d6eae85a243034edf9ec367c1e"></a>

## All schema paths — Property reference / 42b1cccd4164 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--virtual_site--reference--group-001.md#canonical-3f6217e5f02db027ec8dc31b8712c71e8ddad0cb5af19a172d86e766833c596e) |
| `description` | [description](data-sources--virtual_site--reference--group-001.md#canonical-085676c61c64d1e6b28c8fa9fdbdd0b650fbc564962049f250b6e3589a55ec41) |
| `id` | [id](data-sources--virtual_site--reference--group-001.md#canonical-237876726eb161a1d509491fe2a76eb4bf20f1cfe17ac0fdb9241446fc562b75) |
| `labels` | [labels](data-sources--virtual_site--reference--group-001.md#canonical-224ea99f85a2df475b8dfe4a66abf8de2fe2bbd82e2c8e283f555143576f3618) |
| `name` | [name](data-sources--virtual_site--reference--group-001.md#canonical-215a7651731fca9b85c31c19999d9c5b2754f6d6d6b6b8830a739f63d7b45690) |
| `namespace` | [namespace](data-sources--virtual_site--reference--group-001.md#canonical-68ee4c82ca41308cb32b26fe4781bdb2a915d12fcbeea25b90a89851ce93b0c4) |
| `site_selector` | [site_selector](data-sources--virtual_site--reference--group-001.md#canonical-b119cf6899e07e8a2b19b0e0ba04af555f17f9c242f426a0e3b6a276f6743e7d) |
| `site_selector.expressions` | [site_selector.expressions](data-sources--virtual_site--reference--group-001.md#canonical-55a4ed995c11750cf6b1e46b96d2847acceb8f0a52bb0eff9a42af6d547505a5) |
| `site_type` | [site_type](data-sources--virtual_site--reference--group-001.md#canonical-9fe59744e7e192c572b177d48aa59b40e1ffa5345e650969a6de65e95862e654) |

<a id="canonical-c789554a1ff29bc9fb085d887ffff8d26b45e63448f52aad35eeaef08fee4b3a"></a>

## Next pages — Property reference / 42b1cccd4164 / 12

- [site_selector](data-sources--virtual_site--reference--group-001.md#canonical-b14df6e4b7249e796fe0b5d80e90e25055f2a4beabfc0f960bea5e5e60a1c6c3)
- [xcsh_virtual_site](../data-sources/virtual_site.md#canonical-4d9c6e801f11ee0a4dcb7c662b2cdc5dffca2ee94e25d78ffb25ff05789e2980)

<a id="canonical-b14df6e4b7249e796fe0b5d80e90e25055f2a4beabfc0f960bea5e5e60a1c6c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3ca7babab72e05ca092ca69acd269400baded47d8a08f39f20b03f501a984d00"></a>

## site_selector — site_selector / 661f96d0d875 / 2

Breadcrumbs:

- [xcsh_virtual_site](../data-sources/virtual_site.md#canonical-4d9c6e801f11ee0a4dcb7c662b2cdc5dffca2ee94e25d78ffb25ff05789e2980)
- [Property reference](data-sources--virtual_site--reference--group-001.md#canonical-011be626f96641414cc8e65665477f0a7768f907c5427772ed92f6dd2c667a74)
- site_selector

<a id="canonical-b119cf6899e07e8a2b19b0e0ba04af555f17f9c242f426a0e3b6a276f6743e7d"></a>

Type: `"single"`. Computed.

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

<a id="canonical-71b57536829d80fc800902b782c6ba39042e8fa5004f73d5060a9d894d5dd0d2"></a>

## Direct properties — site_selector / 661f96d0d875 / 3

<a id="canonical-55a4ed995c11750cf6b1e46b96d2847acceb8f0a52bb0eff9a42af6d547505a5"></a>

<a id="canonical-e2cda497b5cfeed674d7c437e496a67e23fc1708593437ca812941a37acf9249"></a>

## expressions property — site_selector / 661f96d0d875 / 4

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

<a id="canonical-8da782a66320f133839b44c2b3b5c651b6143d68c415dba7bdbfa1d128666e51"></a>

## Next pages — site_selector / 661f96d0d875 / 5

- [Property reference](data-sources--virtual_site--reference--group-001.md#canonical-011be626f96641414cc8e65665477f0a7768f907c5427772ed92f6dd2c667a74)
- [xcsh_virtual_site](../data-sources/virtual_site.md#canonical-4d9c6e801f11ee0a4dcb7c662b2cdc5dffca2ee94e25d78ffb25ff05789e2980)
