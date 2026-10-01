---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-a78e589611f696217053f3eaa82ff80fc00a2b2b8360350db7fd0396b7463373"></a>

## name property — origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets / 651dee6aa533 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-fd24cafa5ac96ce415f9cb8a94fedfc3991afb5e1c23b3d870376391d676afa4"></a>

<a id="canonical-19b594c126ce0151a80ff3622ea8285837f21c2f318513cee1a1b3ea0f3ccb4c"></a>

## namespace property — origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets / 651dee6aa533 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-a0393093476eedd215c1ed9f90c5e462eddbe52e544a88224a6668d53c2ccc7d"></a>

<a id="canonical-c80210924d93413a83c966972f996abe9b25beb2ea887bb5a861025a15cabd1c"></a>

## tenant property — origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets / 651dee6aa533 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-a53c2fa9c09dad5c20253238d644aa90abc3fb625b6a4297c34c3b662723872e"></a>

<a id="canonical-6bdcdbe7e13abb689ad03ffed90a926cf47fb6551a85f4fd4311dff50cfb316e"></a>

## uid property — origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets / 651dee6aa533 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-461e40c0d2e62a89aca9a877a61649277595eb6f75358ceb50d1783d7bf103bf"></a>

## Next pages — origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets / 651dee6aa533 / 9

- [origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher](data-sources--http_loadbalancer--reference--group-020.md#canonical-a81b6b1361704bc9d254dd5c043e34fb2c5ed70d9576294e8799b122e80cbd57)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-936ef951543a3b228226196d0a14ffb9f9fa9daaff2e0522135c0ddfaecd4534"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-54273e06d03ce82644ebebc6307f4fd0c45fd52974939f7bc9929a8e26ddab3f"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.client_selector — origin_server_subset_rule_list.origin_server_subset_rules.client_selector / 4dd6313b819b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-020.md#canonical-5ae23167a0dfbe57af9a3a473287c8a34fb3bdec8278ec98801ae5b3dff108f6)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-020.md#canonical-5f3334b7d88c1e3db25eca89a8cc3a16ba756a89cea9332325dc4099bf448357)
- origin_server_subset_rule_list.origin_server_subset_rules.client_selector

<a id="canonical-07d08a51801ccea8b6df151ef37c2117979358f2a0585d8782ee52c3b12216a3"></a>

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

<a id="canonical-0fdf0555bf319d10236426cb7fa9bd37d80941d7c433ad7bc3210bd5f035ea09"></a>

## Direct properties — origin_server_subset_rule_list.origin_server_subset_rules.client_selector / 4dd6313b819b / 3

<a id="canonical-63463fb12d5dedec7855c952b50bc4cdc175ff0eb3542d222a0074c809193e19"></a>

<a id="canonical-eaec1612adf2ba5a510231e24fbe75802d3c866e08c01c68675d197600edf147"></a>

## expressions property — origin_server_subset_rule_list.origin_server_subset_rules.client_selector / 4dd6313b819b / 4

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

<a id="canonical-1df89b95e3b8136c3d5abfae1a7aaba131c9c7dcda3ac1976cabb38ea0a0cf81"></a>

## Next pages — origin_server_subset_rule_list.origin_server_subset_rules.client_selector / 4dd6313b819b / 5

- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-020.md#canonical-5f3334b7d88c1e3db25eca89a8cc3a16ba756a89cea9332325dc4099bf448357)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-7330e39d567711896537b62d56a366af7c8e1b671f46adb0203aa25b8d967ec2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ec52b230721f8bb63a33fec596973fbe5097e2807e2b8ed9bef1894fcf70fac9"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher — origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher / 371964dcfa5a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-020.md#canonical-5ae23167a0dfbe57af9a3a473287c8a34fb3bdec8278ec98801ae5b3dff108f6)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-020.md#canonical-5f3334b7d88c1e3db25eca89a8cc3a16ba756a89cea9332325dc4099bf448357)
- origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher

<a id="canonical-898e858ab1619bca047b35ed9d89d0d4565962501ecda4b1ddfa8b5b9e1cc72d"></a>

Type: `"single"`. Computed.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

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

<a id="canonical-e66c5aa1763d670185d4493d6785e2284d6d89aa6b3fcac04fd9fde1768d5489"></a>

## Direct properties — origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher / 371964dcfa5a / 3

<a id="canonical-faa0021386a9d2039aacfd29c3942710587436d0e2685d796b4da2c6aea2b31c"></a>

<a id="canonical-4cad0f2a924b4290c6bc9a3c0d30df83fbb1d866579a6ac159ad7634decbbb64"></a>

## invert_matcher property — origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher / 371964dcfa5a / 4

Type: `"bool"`. Computed.

Invert IP Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [prefix_sets](data-sources--http_loadbalancer--reference--group-021.md#canonical-1ff6c6ffa30260f3f26aa3d7dd0bae8f993c494269177115d2b22b6b1dd5d69d): complete subsection reference.

<a id="canonical-9388af003cf06f5f110bc9bbcc4a797dcc8ba13d35ac4a00b174968d35a560c2"></a>

## Next pages — origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher / 371964dcfa5a / 5

- [origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets](data-sources--http_loadbalancer--reference--group-021.md#canonical-1ff6c6ffa30260f3f26aa3d7dd0bae8f993c494269177115d2b22b6b1dd5d69d)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-020.md#canonical-5f3334b7d88c1e3db25eca89a8cc3a16ba756a89cea9332325dc4099bf448357)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1ff6c6ffa30260f3f26aa3d7dd0bae8f993c494269177115d2b22b6b1dd5d69d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6e1d43148bd1a9d73cc796c3cb4773c18af3a5a5bea7df242c9d31a105a1194"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets — origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets / b54c3e36d228 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-020.md#canonical-5ae23167a0dfbe57af9a3a473287c8a34fb3bdec8278ec98801ae5b3dff108f6)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-020.md#canonical-5f3334b7d88c1e3db25eca89a8cc3a16ba756a89cea9332325dc4099bf448357)
- [origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher](data-sources--http_loadbalancer--reference--group-021.md#canonical-7330e39d567711896537b62d56a366af7c8e1b671f46adb0203aa25b8d967ec2)
- origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets

<a id="canonical-837865c4a754a639274b8fa1f9bcecc047eb64ff4f09cba4e1b05d208a5e03e7"></a>

Type: `"list"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-47db23ac29d0b89a5d668d00324bfe0b8555d232ac50c10b59258cfd7daa9a0d"></a>

## Direct properties — origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets / b54c3e36d228 / 3

<a id="canonical-99c52e369e675385d624d62f1bcf6791eb16c0c04ce2e1d9bfe11aede5d12afe"></a>

<a id="canonical-92381b253d59d1ba1788eed1c953f974c8c53617fe207261efdf55cf7f057e5f"></a>

## kind property — origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets / b54c3e36d228 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-6fd4805ca40f1747fbe59a0a12004195f87d1939399fe8e719a051566f0a9bde"></a>

<a id="canonical-6d547f9f0b3104ba621348db7d9181286f2f091e1ae03d53cecd2633a9948427"></a>

## name property — origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets / b54c3e36d228 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-90cb97aab00a8aacd9c1709f535f566db6d91345ae6fa450878b7f45c63fa8c3"></a>

<a id="canonical-13219fcee94761c79f3f3310aecfeeb44265aee88765249fb3822613079643c1"></a>

## namespace property — origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets / b54c3e36d228 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-50157e470e5ec17223682c09112e35af24115912bef5d0954e95ceff39b46249"></a>

<a id="canonical-251de385c2c77771900e1b8d1a97f74d9ec3d29c2eee745fb7a382535f7b704b"></a>

## tenant property — origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets / b54c3e36d228 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-a3bd532621e1612072af0be576394aab405cf81bd49c13515adb6693536b23bb"></a>

<a id="canonical-4647b96dbbeff585ca5ae7e5f9447fc8649c4bdfa0c5a3d34d4e18746c2ab019"></a>

## uid property — origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets / b54c3e36d228 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2509e07afd49d2bc29207f16f8bb23f91b3f3894ae6d6bcd0f3a5c67f2062cee"></a>

## Next pages — origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets / b54c3e36d228 / 9

- [origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher](data-sources--http_loadbalancer--reference--group-021.md#canonical-7330e39d567711896537b62d56a366af7c8e1b671f46adb0203aa25b8d967ec2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f3d9d1b9f2fe4b2d9b758bc20508ef83fdcb1915353a1cd13b9a2d8cec6e6f8a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-520840b2b803286dbfb4a5735de2dc71b64a765b517823486b286abdebd180b5"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list — origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list / 5f6423d3b245 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-020.md#canonical-5ae23167a0dfbe57af9a3a473287c8a34fb3bdec8278ec98801ae5b3dff108f6)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-020.md#canonical-5f3334b7d88c1e3db25eca89a8cc3a16ba756a89cea9332325dc4099bf448357)
- origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list

<a id="canonical-6d61e5543287e1d4c5dc39ca11a091631704b9153754925030bb0a0455dde9f3"></a>

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

<a id="canonical-a5f72e97e278da97beb709f8a43e4b5418dd00daeaf466acc3cdf0c4d179317a"></a>

## Direct properties — origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list / 5f6423d3b245 / 3

<a id="canonical-4c9e2fcbcb81f3a8acc0daa55969c387a41c715668d431e551d0402622e63a80"></a>

<a id="canonical-6ed3dda96d0cfa3ccf021de6a0c18e2a6f0b8c193a7cbc435c0d7908f3b9ead9"></a>

## invert_match property — origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list / 5f6423d3b245 / 4

Type: `"bool"`. Computed.

Invert Match Result. Invert the match result.

Upstream description:

Invert the match result.

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

<a id="canonical-57bbc92af84ac5205eb883499c7946aa31147e541397422c763ac1a1dd936a10"></a>

<a id="canonical-4c8263b90743e4ceee4149ad5ed287098c8af82e383dbd99b3b912167f119266"></a>

## ip_prefixes property — origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list / 5f6423d3b245 / 5

Type: `["list", "string"]`. Computed.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

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

<a id="canonical-535e7a35617377bb74c71849fe67020a67cdd08b5a18ad93ff2bf69f02c8722c"></a>

## Next pages — origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list / 5f6423d3b245 / 6

- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-020.md#canonical-5f3334b7d88c1e3db25eca89a8cc3a16ba756a89cea9332325dc4099bf448357)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f4d67e26695f3e71f67d7cb5a3835eeb4f3bd4c7f6b0a7193453fe8f5eb2bb3d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f46ab14fc79ee300cf8e7fea1e6a89024269c978370afb00380a653681e32f3"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.metadata — origin_server_subset_rule_list.origin_server_subset_rules.metadata / adf535a13c6e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-020.md#canonical-5ae23167a0dfbe57af9a3a473287c8a34fb3bdec8278ec98801ae5b3dff108f6)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-020.md#canonical-5f3334b7d88c1e3db25eca89a8cc3a16ba756a89cea9332325dc4099bf448357)
- origin_server_subset_rule_list.origin_server_subset_rules.metadata

<a id="canonical-b0973274abd973f10ab5ac6f8bc142201af9d66d00957c4f040597ec0d405720"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-3770263e3d9f5a4f3167fbcce2e08374ac7bb5c42256df98bc43d52ee84c11da"></a>

## Direct properties — origin_server_subset_rule_list.origin_server_subset_rules.metadata / adf535a13c6e / 3

<a id="canonical-e3d1174c166ebde42f6a1e09c55f4b2740497aa5d6cab6c7d753a1ac735db1c7"></a>

<a id="canonical-29211eb04bb1fc55e2f7b14261d7fb6a3f3c93affdd386b9aa09b03cefa15b60"></a>

## description_spec property — origin_server_subset_rule_list.origin_server_subset_rules.metadata / adf535a13c6e / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-2916821004490c0bb8e1b39105718241f20e27e7bb26652d38d18efee08b2a17"></a>

<a id="canonical-4f0a6997e716cc059ed4e6ab1a182caff1cd7a080420f63e7c2f47c2e615b127"></a>

## name property — origin_server_subset_rule_list.origin_server_subset_rules.metadata / adf535a13c6e / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-d83e68c768d53edaa5c0239383945a57a7699daf4a6e67bab687c18f1e764480"></a>

## Next pages — origin_server_subset_rule_list.origin_server_subset_rules.metadata / adf535a13c6e / 6

- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-020.md#canonical-5f3334b7d88c1e3db25eca89a8cc3a16ba756a89cea9332325dc4099bf448357)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-15bdf6183b155c8aac2cda666c53c1baa5e6d06182e7f9c94988aa9b1aa510ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3180bdf6bf0a12c1c1982ce23ff2a778a1ea7040d214d94d69f1e3ca24483a74"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.none — origin_server_subset_rule_list.origin_server_subset_rules.none / 8a3f18b92d3f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-020.md#canonical-5ae23167a0dfbe57af9a3a473287c8a34fb3bdec8278ec98801ae5b3dff108f6)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-020.md#canonical-5f3334b7d88c1e3db25eca89a8cc3a16ba756a89cea9332325dc4099bf448357)
- origin_server_subset_rule_list.origin_server_subset_rules.none

<a id="canonical-e800270a474a7a6f087fd06ba996b707d339d46e92bab8d3b1a0baabbdb6283b"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-4f87d8dc19fa71dd6c02a5ade51695fa68415eae3c934d1ae398837f4dcc3e4f"></a>

## Direct properties — origin_server_subset_rule_list.origin_server_subset_rules.none / 8a3f18b92d3f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f997dedceb1d4f347da308e098baa00e6d2fe71c3f1c3e856e3f6e51268e430a"></a>

## Next pages — origin_server_subset_rule_list.origin_server_subset_rules.none / 8a3f18b92d3f / 4

- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-020.md#canonical-5f3334b7d88c1e3db25eca89a8cc3a16ba756a89cea9332325dc4099bf448357)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ad8cd3774e6de73cb680d0cd4b5ac1d39939eda81a8b5cb1b457c2477a7743d"></a>

## policy_based_challenge — policy_based_challenge / 323f7e321248 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- policy_based_challenge

<a id="canonical-00114bfff9e8f35fff6fff86c0aeca8550664aeea9d1c4d88e22e58a24abd481"></a>

Type: `"single"`. Computed.

Specifies the settings for policy rule based challenge.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-captcha_challenge_parameters_choice": "[\"captcha_challenge_parameters\",\"default_captcha_challenge_parameters\"]",
  "x-ves-oneof-field-challenge_choice": "[\"always_enable_captcha_challenge\",\"always_enable_js_challenge\",\"no_challenge\"]",
  "x-ves-oneof-field-js_challenge_parameters_choice": "[\"default_js_challenge_parameters\",\"js_challenge_parameters\"]",
  "x-ves-oneof-field-malicious_user_mitigation_choice": "[\"default_mitigation_settings\",\"malicious_user_mitigation\"]",
  "x-ves-oneof-field-temporary_blocking_parameters_choice": "[\"default_temporary_blocking_parameters\",\"temporary_user_blocking\"]"
}
```

<a id="canonical-c4ea32b2740f28732a4f156d3a9bd75f8ae699edaf05c9c4410004ba1aaa3467"></a>

## Direct properties — policy_based_challenge / 323f7e321248 / 3

- [always_enable_captcha_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-a09b6c1c03f0a704b5e55eb2940165d06748e336730a1bcec3be96697fd8c429): complete subsection reference.

- [always_enable_js_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-5fa8076c0b3a00d6553ecd75d743ff265dc6dd013dabb162aaf287d9f5eae4d4): complete subsection reference.

- [captcha_challenge_parameters](data-sources--http_loadbalancer--reference--group-021.md#canonical-0db9bb3525490cb01a8aab7b4509943d53025edf100cde9fde70933b12613756): complete subsection reference.

- [default_captcha_challenge_parameters](data-sources--http_loadbalancer--reference--group-021.md#canonical-cf69a7acc6de96ff9c1d30009ce1292fa2d1a8bb381383e9c4cb51b669778f22): complete subsection reference.

- [default_js_challenge_parameters](data-sources--http_loadbalancer--reference--group-021.md#canonical-e4beb0c55b2c96dc2149e4602eb0039073c1e62b93cf2df1d5be220c7fa59260): complete subsection reference.

- [default_mitigation_settings](data-sources--http_loadbalancer--reference--group-021.md#canonical-c150262148c59b62041ff3611b825629253c9bed0691af38e8d14a93868e4c85): complete subsection reference.

- [default_temporary_blocking_parameters](data-sources--http_loadbalancer--reference--group-021.md#canonical-b776d220a4d4556e1e288c75a32e233c647d46a1dde043b011b689fc794a6971): complete subsection reference.

- [js_challenge_parameters](data-sources--http_loadbalancer--reference--group-021.md#canonical-763954b8189fdcdffeb54b71dbe21259109972aeb6e7af585d73af8d29a01253): complete subsection reference.

- [malicious_user_mitigation](data-sources--http_loadbalancer--reference--group-021.md#canonical-31167c4d724a10b4434ffd95b80b9dbf678a0d8fe36ce05b77c024a326762f70): complete subsection reference.

- [no_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-58f6c45bc8230d2dd4e8b01d9c897081a26c20c210a222c3ac9ec440bb61f80e): complete subsection reference.

- [rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66): complete subsection reference.

- [temporary_user_blocking](data-sources--http_loadbalancer--reference--group-022.md#canonical-d72d6cbe3fcf27ef049a25fc77b8ca27f89051a311fc8bc01e902a9e43113657): complete subsection reference.

<a id="canonical-2681f0701a3b29cb35851881bd1d841c463a30b9baf2a72984f7c8ab61c435b3"></a>

## Next pages — policy_based_challenge / 323f7e321248 / 4

- [policy_based_challenge.always_enable_captcha_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-a09b6c1c03f0a704b5e55eb2940165d06748e336730a1bcec3be96697fd8c429)
- [policy_based_challenge.always_enable_js_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-5fa8076c0b3a00d6553ecd75d743ff265dc6dd013dabb162aaf287d9f5eae4d4)
- [policy_based_challenge.captcha_challenge_parameters](data-sources--http_loadbalancer--reference--group-021.md#canonical-0db9bb3525490cb01a8aab7b4509943d53025edf100cde9fde70933b12613756)
- [policy_based_challenge.default_captcha_challenge_parameters](data-sources--http_loadbalancer--reference--group-021.md#canonical-cf69a7acc6de96ff9c1d30009ce1292fa2d1a8bb381383e9c4cb51b669778f22)
- [policy_based_challenge.default_js_challenge_parameters](data-sources--http_loadbalancer--reference--group-021.md#canonical-e4beb0c55b2c96dc2149e4602eb0039073c1e62b93cf2df1d5be220c7fa59260)
- [policy_based_challenge.default_mitigation_settings](data-sources--http_loadbalancer--reference--group-021.md#canonical-c150262148c59b62041ff3611b825629253c9bed0691af38e8d14a93868e4c85)
- [policy_based_challenge.default_temporary_blocking_parameters](data-sources--http_loadbalancer--reference--group-021.md#canonical-b776d220a4d4556e1e288c75a32e233c647d46a1dde043b011b689fc794a6971)
- [policy_based_challenge.js_challenge_parameters](data-sources--http_loadbalancer--reference--group-021.md#canonical-763954b8189fdcdffeb54b71dbe21259109972aeb6e7af585d73af8d29a01253)
- [policy_based_challenge.malicious_user_mitigation](data-sources--http_loadbalancer--reference--group-021.md#canonical-31167c4d724a10b4434ffd95b80b9dbf678a0d8fe36ce05b77c024a326762f70)
- [policy_based_challenge.no_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-58f6c45bc8230d2dd4e8b01d9c897081a26c20c210a222c3ac9ec440bb61f80e)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.temporary_user_blocking](data-sources--http_loadbalancer--reference--group-022.md#canonical-d72d6cbe3fcf27ef049a25fc77b8ca27f89051a311fc8bc01e902a9e43113657)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a09b6c1c03f0a704b5e55eb2940165d06748e336730a1bcec3be96697fd8c429"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c63d4980297b75892cecd5e44b312c2bcb4d50b22c94f8eb65c086bdfe8d5087"></a>

## policy_based_challenge.always_enable_captcha_challenge — policy_based_challenge.always_enable_captcha_challenge / 74eeed679120 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- policy_based_challenge.always_enable_captcha_challenge

<a id="canonical-d91861e5fb9dfb0419e8d6055373132bea8ea48e962702d8879133a4f17b3a12"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for always enable captcha challenge.

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

<a id="canonical-847a7c3941c9fe3a2072ac95cd4faf384b79a1375b7b20736c534b04692ae4c2"></a>

## Direct properties — policy_based_challenge.always_enable_captcha_challenge / 74eeed679120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3378d8995e08895267bf21fd7f34b250ada5e74887af64f14b706545c15be292"></a>

## Next pages — policy_based_challenge.always_enable_captcha_challenge / 74eeed679120 / 4

- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5fa8076c0b3a00d6553ecd75d743ff265dc6dd013dabb162aaf287d9f5eae4d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9152385c59f71bba88f345bf91f3233b5f2fa86ed5d5ca15a1a5f06ef441c965"></a>

## policy_based_challenge.always_enable_js_challenge — policy_based_challenge.always_enable_js_challenge / 96292d928891 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- policy_based_challenge.always_enable_js_challenge

<a id="canonical-fb286dccaf6e9744d39b1128b5e3c252ce9ec42070d7494121a902a676df38d2"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for always enable js challenge.

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

<a id="canonical-3965da3b82636663fcc55a10acd86600f8315faca15ae9c616b9cf1ad493e6e3"></a>

## Direct properties — policy_based_challenge.always_enable_js_challenge / 96292d928891 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-219c1818e18b603ac6801fa8bf63a96562a43a07c791f7acd826be890106894c"></a>

## Next pages — policy_based_challenge.always_enable_js_challenge / 96292d928891 / 4

- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-0db9bb3525490cb01a8aab7b4509943d53025edf100cde9fde70933b12613756"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb3140d35e82d7a7acc89b3ff47299e3ce12540329242fa5ab4ab5a655cba303"></a>

## policy_based_challenge.captcha_challenge_parameters — policy_based_challenge.captcha_challenge_parameters / a47cbec9cd5e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- policy_based_challenge.captcha_challenge_parameters

<a id="canonical-1ecd9df7e81d130aa2d5a135ee45f8bfab82f5779d80e1f80691d65a086ff20c"></a>

Type: `"single"`. Computed.

Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google
Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed
to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will
redirect..

Upstream description:

Enables loadbalancer to perform captcha challenge

Captcha challenge will be based on Google Recaptcha.

With this feature enabled, only clients that pass the captcha challenge will be allowed to complete
the HTTP request.

When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client
will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will
tag response header with a cookie to avoid Captcha challenge for subsequent requests.

CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally,
computers or bots are not capable of solving a captcha.

You can enable either Javascript challenge or Captcha challenge on a virtual host.

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

<a id="canonical-c4a50f71ce2ca3be01efdb1ff4b69a4cfd15555caee2348a3813583991f3e11b"></a>

## Direct properties — policy_based_challenge.captcha_challenge_parameters / a47cbec9cd5e / 3

<a id="canonical-fd8c68191811af0f629e78e170a8f4f1ffbd642f1bdf14f267d64d8294f85f43"></a>

<a id="canonical-b400ebf0356cc1e76402cc0e41d2cb91de4bc75b95c2d91644992fbd8e3df8e7"></a>

## cookie_expiry property — policy_based_challenge.captcha_challenge_parameters / a47cbec9cd5e / 4

Type: `"number"`. Computed.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-19697616742e5630f708fe4d649746298ccab854a16bb8a24da65a677c4d80b0"></a>

<a id="canonical-bea0e9dab19c1bde941976e245366fbc85402f1c94ba1aabd1b6191bee542ef8"></a>

## custom_page property — policy_based_challenge.captcha_challenge_parameters / a47cbec9cd5e / 5

Type: `"string"`. Computed.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-a89eb43af80be2c44616410e4a7b013b4cfaf849f3463d57cdcdb384c5d30536"></a>

## Next pages — policy_based_challenge.captcha_challenge_parameters / a47cbec9cd5e / 6

- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-cf69a7acc6de96ff9c1d30009ce1292fa2d1a8bb381383e9c4cb51b669778f22"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d7fa23677f731f76ad3001c2ccc1500709575b08377a3842eff07d820042c3c8"></a>

## policy_based_challenge.default_captcha_challenge_parameters — policy_based_challenge.default_captcha_challenge_parameters / e00d1db8afd5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- policy_based_challenge.default_captcha_challenge_parameters

<a id="canonical-fa44d584c3da54bccd513283f3b8573f71dc5dec2929fc663bfcaebae780e3c7"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default captcha challenge parameters.

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

<a id="canonical-991d982dd4beb27188e28c63e6bf5539d30d1f0bdb7004233227c1b862bd8b3b"></a>

## Direct properties — policy_based_challenge.default_captcha_challenge_parameters / e00d1db8afd5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d86c711b838aa64b1e1802419e4d9ad0488e91f2a6adf167c6a08fbbc43a41c1"></a>

## Next pages — policy_based_challenge.default_captcha_challenge_parameters / e00d1db8afd5 / 4

- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e4beb0c55b2c96dc2149e4602eb0039073c1e62b93cf2df1d5be220c7fa59260"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc58aa663041faa7f326cd64bdc38a0effdc35f088a023c3a2d325c95e7e6981"></a>

## policy_based_challenge.default_js_challenge_parameters — policy_based_challenge.default_js_challenge_parameters / e68276b80ffa / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- policy_based_challenge.default_js_challenge_parameters

<a id="canonical-baf24352a07d5e39b81806422061c6e528e548b2270ec37c08e21b9844cf7a61"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default js challenge parameters.

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

<a id="canonical-e5ea747549c02ffb421d1c299c0b98343cd0e66e54479025251d7f383c6e299f"></a>

## Direct properties — policy_based_challenge.default_js_challenge_parameters / e68276b80ffa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110b6ac5f01eedb20ec2f8992a43452ecf04bd7093af40ea1931c8ac48e2eb6"></a>

## Next pages — policy_based_challenge.default_js_challenge_parameters / e68276b80ffa / 4

- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c150262148c59b62041ff3611b825629253c9bed0691af38e8d14a93868e4c85"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-977f34089505a099d9d42ecbef41c516f8c148dbf4017790bfd022de1c63f1d9"></a>

## policy_based_challenge.default_mitigation_settings — policy_based_challenge.default_mitigation_settings / 5b341b5216ef / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- policy_based_challenge.default_mitigation_settings

<a id="canonical-28b35c9a455da7642a159d19fdbd26643aee107b207df9c81dac906cdc677f2f"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-97116a52f5c61969ace8aa89e77153eae38ed1c0a198de86931c24681d993819"></a>

## Direct properties — policy_based_challenge.default_mitigation_settings / 5b341b5216ef / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3416d4d3ec9c5d95e954543242dc21523b6e698b807e3eb3e575923bcda9fd7e"></a>

## Next pages — policy_based_challenge.default_mitigation_settings / 5b341b5216ef / 4

- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-b776d220a4d4556e1e288c75a32e233c647d46a1dde043b011b689fc794a6971"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7984f09157466a60da9c5fa42c7c0951e1143abe29e60cc6b13abbc6d7383ee"></a>

## policy_based_challenge.default_temporary_blocking_parameters — policy_based_challenge.default_temporary_blocking_parameters / 08a5dded352f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- policy_based_challenge.default_temporary_blocking_parameters

<a id="canonical-2064bd2c55c60f48e4b153d33c4b39e9ef6e44557c1743a25df21f8b84b9f3f5"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-e34808aa72126c7d3052175b19d0e8e082886dfdfdc7344a38bbf23546972961"></a>

## Direct properties — policy_based_challenge.default_temporary_blocking_parameters / 08a5dded352f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-aee132974a14509787b01e0b6158260008160fc5b83b8700da86ec48eaf7df29"></a>

## Next pages — policy_based_challenge.default_temporary_blocking_parameters / 08a5dded352f / 4

- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-763954b8189fdcdffeb54b71dbe21259109972aeb6e7af585d73af8d29a01253"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1ce3dfd9e55349e466f7751bd76b038695e17c29953077ad7437a218c53593e"></a>

## policy_based_challenge.js_challenge_parameters — policy_based_challenge.js_challenge_parameters / dd1dd136bdbc / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- policy_based_challenge.js_challenge_parameters

<a id="canonical-e30ba17cd7fa6603ea060a360bf560c7ea8f4b70d125ed509834b1d20ba836bd"></a>

Type: `"single"`. Computed.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly
browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do..

Upstream description:

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
Javascript.

With this feature enabled, only clients that are capable of executing Javascript(mostly browsers)
will be allowed to complete the HTTP request.

When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer
chooses a set of random numbers for every new client and sends these numbers along with an encrypted
answer with the request such that it embed these numbers as input in the Javascript. Javascript will
run on the requester browser and perform a complex Math operation. Script will submit the answer to
loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the
decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the
upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie
to avoid Javascript challenge for subsequent requests.

Javascript challenge serves following purposes \* Validate that the request is coming via a browser
that is capable for running Javascript \* Force the browser to run a complex operation, f(X), that
requires it to spend a large number of CPU cycles. This is to slow down a potential DoS attacker by
making it difficult to launch a large request flood without having to spend even larger CPU cost at
their end.

You can enable either Javascript challenge or Captcha challenge on a virtual host.

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

<a id="canonical-72fe7a4ae6bf97cbc709b4ddf8ff41150c60a57cec6b4f588a6efa31ba5aeb67"></a>

## Direct properties — policy_based_challenge.js_challenge_parameters / dd1dd136bdbc / 3

<a id="canonical-037fdd800564d6e9309de5e7f3792bf743c8ebe82abf16d2506f8df878fb6d21"></a>

<a id="canonical-c57a4f7fe579dcc9f0d18be64189a5ede105f156cc15d681d68e178bf96ecada"></a>

## cookie_expiry property — policy_based_challenge.js_challenge_parameters / dd1dd136bdbc / 4

Type: `"number"`. Computed.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-63b4fd1fbdb6871bf49ad1fb8c69c982b5637c009a0bb6f800fd3b706dfd21f9"></a>

<a id="canonical-038c51f9b7b6ea178031ce2d3c0b1874b52c2d91b51ec2bd83815beead1af36c"></a>

## custom_page property — policy_based_challenge.js_challenge_parameters / dd1dd136bdbc / 5

Type: `"string"`. Computed.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-e408c1eadeabf2d9f5bf6e4923789f1927a5157bea3c142e26b8b7f639122621"></a>

<a id="canonical-d8442048a95e95498e11b85a8752d84969cb4b3f9932310790371147b3905ab2"></a>

## js_script_delay property — policy_based_challenge.js_challenge_parameters / dd1dd136bdbc / 6

Type: `"number"`. Computed.

Delay introduced by Javascript, in milliseconds.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-a881b0a5b4f9d5933dd7f2efcfd143d53336c76eda65e8e20a98c56e0a03da5b"></a>

## Next pages — policy_based_challenge.js_challenge_parameters / dd1dd136bdbc / 7

- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-31167c4d724a10b4434ffd95b80b9dbf678a0d8fe36ce05b77c024a326762f70"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9afacf5575c8555d19afcc45a7610a9f9e2179e1f28c6665c8f09828df7e9f58"></a>

## policy_based_challenge.malicious_user_mitigation — policy_based_challenge.malicious_user_mitigation / 271d4cd6f589 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- policy_based_challenge.malicious_user_mitigation

<a id="canonical-f268e3bb3ed532681e2d77e8c373e066d5a0396d540d4e4403a7e1f041579c0f"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-0899f12a6319e5c79bb59cf6c870e8f273d5cd2a2f4dc18eb601ea8bab1e9e51"></a>

## Direct properties — policy_based_challenge.malicious_user_mitigation / 271d4cd6f589 / 3

<a id="canonical-fc869eaf9fa84c2850835a7c37a8aeca61d174f3a0de5a46edaaa8c5689d05ce"></a>

<a id="canonical-f5830fe9dbfa80342470ed46b354e27dd2b3fa71a6dfe4c7d72502c33ee97745"></a>

## name property — policy_based_challenge.malicious_user_mitigation / 271d4cd6f589 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-fec9c9585946c8673887904b2565524bedbdbcf8bd2509acfb5b63dd1cab70d0"></a>

<a id="canonical-64dbaf3f02af070c1e63b294e2682cd214d42d962b021f1653d7f96e0a312ab0"></a>

## namespace property — policy_based_challenge.malicious_user_mitigation / 271d4cd6f589 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0233dffb80905a9098fa9250ce6788f0b3ea88f4f62560683666c79a8b48e794"></a>

<a id="canonical-fa39b487e4bc2629208cfe819bec97559e6117595eee7ecb4539ecadce9de153"></a>

## tenant property — policy_based_challenge.malicious_user_mitigation / 271d4cd6f589 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-5ae14bf2e9797cc6a69fd030fae498b113026e7b11adfc0cc026e011d180edd0"></a>

## Next pages — policy_based_challenge.malicious_user_mitigation / 271d4cd6f589 / 7

- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-58f6c45bc8230d2dd4e8b01d9c897081a26c20c210a222c3ac9ec440bb61f80e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-db52fd884abc5bc209c012c52210cd3165db56436b59ff87ef46ab94ced7b900"></a>

## policy_based_challenge.no_challenge — policy_based_challenge.no_challenge / bad7be598a95 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- policy_based_challenge.no_challenge

<a id="canonical-b871ced944a512ee11db3771edc2d43316058d3f58699831f04fada720aa6ede"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no challenge.

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

<a id="canonical-d26207d6f03593f91753f878e37704522b14be40c4e04f4c169f5be578d42b2d"></a>

## Direct properties — policy_based_challenge.no_challenge / bad7be598a95 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-89aff0e4a74b108e65aff1f23d0151fa1c35978471dfa0894bf474e769867ee6"></a>

## Next pages — policy_based_challenge.no_challenge / bad7be598a95 / 4

- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-52ac82d5cb899e0a439d96247d17cf22c28de61472379b53a966e284fdee5547"></a>

## policy_based_challenge.rule_list — policy_based_challenge.rule_list / 540357091572 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- policy_based_challenge.rule_list

<a id="canonical-163c8033dc70dc554dd221e574adfcfca64bb2e6685293361e6537cb9d6cf3c9"></a>

Type: `"single"`. Computed.

List of challenge rules to be used in policy based challenge.

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

<a id="canonical-24f332e003189023a97feb05437bb4805f8943d18d10ca63228c4735ca38ecf2"></a>

## Direct properties — policy_based_challenge.rule_list / 540357091572 / 3

- [rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143): complete subsection reference.

<a id="canonical-dc1f8b9d55bf72c59594821e36b38c6f00f8db03f97553e29c1a5e3da3116877"></a>

## Next pages — policy_based_challenge.rule_list / 540357091572 / 4

- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-52ce98a855a0d6aa6bf2cbe48844ff2126f2393db5acc715a3fa6e5321139b6d"></a>

## policy_based_challenge.rule_list.rules — policy_based_challenge.rule_list.rules / 055e1e9bb702 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- policy_based_challenge.rule_list.rules

<a id="canonical-02167598a32dfa9e963e667f37ffc00042fd4a3cb2dffa70ab64edb32a50a7f7"></a>

Type: `"list"`. Computed.

Rules that specify the match conditions and challenge type to be launched. When a challenge type is
selected to be always enabled, these rules can be used to disable challenge or launch a different
challenge for requests that match the specified conditions.

Upstream description:

Rules that specify the match conditions and challenge type to be launched. When a challenge type is
selected to be always enabled, these rules can be used to disable challenge or launch a different
challenge for requests that match the specified conditions.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-33c6d8ede1964c7392b37fbe09c4caa18e66a4068ce84862380ff422887d9e7e"></a>

## Direct properties — policy_based_challenge.rule_list.rules / 055e1e9bb702 / 3

- [metadata](data-sources--http_loadbalancer--reference--group-021.md#canonical-a88b93e71c136f5f82e8356c909fb5f8f97329fa477ba14b948dbc2eac68edd8): complete subsection reference.

- [spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048): complete subsection reference.

<a id="canonical-587330dde6c02c483c91ab186d50541d53baf460b78f3d94910ee0f556b97cd2"></a>

## Next pages — policy_based_challenge.rule_list.rules / 055e1e9bb702 / 4

- [policy_based_challenge.rule_list.rules.metadata](data-sources--http_loadbalancer--reference--group-021.md#canonical-a88b93e71c136f5f82e8356c909fb5f8f97329fa477ba14b948dbc2eac68edd8)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a88b93e71c136f5f82e8356c909fb5f8f97329fa477ba14b948dbc2eac68edd8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8c0f5a539e4b4c954812fc8bf4d0f5c2efd413d877569d5bccc1dc2b224b85ad"></a>

## policy_based_challenge.rule_list.rules.metadata — policy_based_challenge.rule_list.rules.metadata / 648c4ae7f457 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- policy_based_challenge.rule_list.rules.metadata

<a id="canonical-786a636c8683a17592cd78aa8ff3dcabee656b996fd2974d6b9025aeb6357c96"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-490f79863a16b2c078c82a4dd73d549634cd31f93ebada51d529ac28a1dcdba9"></a>

## Direct properties — policy_based_challenge.rule_list.rules.metadata / 648c4ae7f457 / 3

<a id="canonical-ee4bdec5dfab7e8138f147123db363243a9bc20bd04128e55f916f666dc5bea3"></a>

<a id="canonical-08fe98d4a6048b30e2f8a231e1fa0935b7ee8346bf3eca7bc6f3afc2dd487b20"></a>

## description_spec property — policy_based_challenge.rule_list.rules.metadata / 648c4ae7f457 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-23e97b0cbc35a5c62ee270ba85bf602add420aeab6daa24924c4ca271314c14e"></a>

<a id="canonical-62ad47d2e0cc0c23cd69e659f7762301e552253849368a3bb068c6fee004f2ed"></a>

## name property — policy_based_challenge.rule_list.rules.metadata / 648c4ae7f457 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-c95b25b1fe209c622489784bd5b0ec14f51a97e09ca8b5f8c4903175b6db652f"></a>

## Next pages — policy_based_challenge.rule_list.rules.metadata / 648c4ae7f457 / 6

- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f2c7ca7f5913329a2d2a5ecd03407d4d53e573114082076270e955a4fe3a4279"></a>

## policy_based_challenge.rule_list.rules.spec — policy_based_challenge.rule_list.rules.spec / a4158d45115b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- policy_based_challenge.rule_list.rules.spec

<a id="canonical-3b364fea6f6a7e54a0a398ea7e7990a5e04f99c54501ee7013748a5ca77c1029"></a>

Type: `"single"`. Computed.

Challenge Rule consists of an unordered list of predicates and an action. The predicates are
evaluated against a set of input fields that are extracted from or derived from an L7 request API. A
request API is considered to match the rule if all predicates in the rule evaluate to true for
that..

Upstream description:

A Challenge Rule consists of an unordered list of predicates and an action. The predicates are
evaluated against a set of input fields that are extracted from or derived from an L7 request API. A
request API is considered to match the rule if all predicates in the rule evaluate to true for that
request. Any predicates that are not specified in a rule are implicitly considered to be true. If a
request API matches a challenge rule, the configured challenge is enforced.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-asn_choice": "[\"any_asn\",\"asn_list\",\"asn_matcher\"]",
  "x-ves-oneof-field-challenge_action": "[\"disable_challenge\",\"enable_captcha_challenge\",\"enable_javascript_challenge\"]",
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_selector\"]",
  "x-ves-oneof-field-ip_choice": "[\"any_ip\",\"ip_matcher\",\"ip_prefix_list\"]",
  "x-ves-oneof-field-tls_fingerprint_choice": "[\"tls_fingerprint_matcher\"]"
}
```

<a id="canonical-e0d3709f04d4499bb6491f6028a3c4b1941b3964b32bf5b00d40274499a1afd7"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec / a4158d45115b / 3

- [any_asn](data-sources--http_loadbalancer--reference--group-021.md#canonical-54581841b38cf0239cb482b931035b81f3503053e7ff140def543f35af830ab6): complete subsection reference.

- [any_client](data-sources--http_loadbalancer--reference--group-021.md#canonical-9d3e6523437e8b4418a4592c844eeafba14b2dff569461b4891ed01f54eede81): complete subsection reference.

- [any_ip](data-sources--http_loadbalancer--reference--group-021.md#canonical-a9b50c0ba4233ac90de12fe29280c3b5592c20ee6ce1843fccd64ddc38d8a5b7): complete subsection reference.

- [arg_matchers](data-sources--http_loadbalancer--reference--group-021.md#canonical-143c1711b2006f094e2001ed44dc4679829cef1ab14645997bdf19783ae956f9): complete subsection reference.

- [asn_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-921f66a51c2994c217f6c09c51917f443337e6b545274087c76e81ebf350e7e1): complete subsection reference.

- [asn_matcher](data-sources--http_loadbalancer--reference--group-021.md#canonical-64fcbf7242e313d33bf68832dd18ab4a4c930894aaed6be19061b651ad617d29): complete subsection reference.

- [body_matcher](data-sources--http_loadbalancer--reference--group-021.md#canonical-9cab2477afc810b9fa13fd67539c09719fe1bf3eb292458ab97d2e5ebb841a86): complete subsection reference.

- [client_selector](data-sources--http_loadbalancer--reference--group-021.md#canonical-80eebafba74da3ae8e1b203e532cfd41f3d0069d94c11525c9c8e1d2ab85eeaa): complete subsection reference.

- [cookie_matchers](data-sources--http_loadbalancer--reference--group-021.md#canonical-24c1aa264f8eb4a87fc88358f42b3c04cb2afaffb4d457b436297a30fd5c0940): complete subsection reference.

- [disable_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-c87d9460aa1e6756d53b4a5ea7ce14797079b5a1ddfecdd668072b4ca250c352): complete subsection reference.

- [domain_matcher](data-sources--http_loadbalancer--reference--group-021.md#canonical-f5693976020e0fe3642cc5a18d3715bce89e0cc47b0c2c4a6c128e50d8e5f49e): complete subsection reference.

- [enable_captcha_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-ba606181305fde7d30fc9c0a866fe75fd052c9ac01e6f24d53c403f2bdc85192): complete subsection reference.

- [enable_javascript_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1ce2ed30433f39b44c880855cae0e8c654804b94382ba72112502693629174e3): complete subsection reference.

<a id="canonical-cf5b6fee3ab6b474080c470d8aaaedd38a9551c81ba7bec2a3640c4b1e40ccbe"></a>

<a id="canonical-77204c1a8d786b010c872f78e3b047c50f27b01e721541ca5738f907435b235c"></a>

## expiration_timestamp property — policy_based_challenge.rule_list.rules.spec / a4158d45115b / 4

Type: `"string"`. Computed.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [headers](data-sources--http_loadbalancer--reference--group-021.md#canonical-578f1ac411c848abb3f41924f604ef7400bb2e9d8819058431c6cd4a3724b811): complete subsection reference.

- [http_method](data-sources--http_loadbalancer--reference--group-021.md#canonical-c8d31f5a660171504b9e6f3dda4d7939826a3e3898668d617fa0c00ed87292e3): complete subsection reference.

- [ip_matcher](data-sources--http_loadbalancer--reference--group-021.md#canonical-ca6c31e0e3e39f3e19be14e0bd99fff1219d5beca5dc8e55559dea263cb80880): complete subsection reference.

- [ip_prefix_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-929811c07d5b3b04bed514c363d8cf6d54eb5be41f99307e66bcd8a56a22bb19): complete subsection reference.

- [path](data-sources--http_loadbalancer--reference--group-022.md#canonical-5c5a7d6aa6611cf9b6261f1644001a97fd2267bb79e06050fc298a96b453cc7f): complete subsection reference.

- [query_params](data-sources--http_loadbalancer--reference--group-022.md#canonical-a6dd4dd1d3204263ff38e0262d2b6b68d0e4ca2bb4f2d3300cad7ac24228173e): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--http_loadbalancer--reference--group-022.md#canonical-0a5c9cf79b02cbcd2d25551795dc18bc688986f470da1226073d5aba1c58f6cb): complete subsection reference.

<a id="canonical-657a5918e790b120c2e5e57f3dde5e7b847ced0900d3d08cd40e26170541b32c"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec / a4158d45115b / 5

- [policy_based_challenge.rule_list.rules.spec.any_asn](data-sources--http_loadbalancer--reference--group-021.md#canonical-54581841b38cf0239cb482b931035b81f3503053e7ff140def543f35af830ab6)
- [policy_based_challenge.rule_list.rules.spec.any_client](data-sources--http_loadbalancer--reference--group-021.md#canonical-9d3e6523437e8b4418a4592c844eeafba14b2dff569461b4891ed01f54eede81)
- [policy_based_challenge.rule_list.rules.spec.any_ip](data-sources--http_loadbalancer--reference--group-021.md#canonical-a9b50c0ba4233ac90de12fe29280c3b5592c20ee6ce1843fccd64ddc38d8a5b7)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](data-sources--http_loadbalancer--reference--group-021.md#canonical-143c1711b2006f094e2001ed44dc4679829cef1ab14645997bdf19783ae956f9)
- [policy_based_challenge.rule_list.rules.spec.asn_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-921f66a51c2994c217f6c09c51917f443337e6b545274087c76e81ebf350e7e1)
- [policy_based_challenge.rule_list.rules.spec.asn_matcher](data-sources--http_loadbalancer--reference--group-021.md#canonical-64fcbf7242e313d33bf68832dd18ab4a4c930894aaed6be19061b651ad617d29)
- [policy_based_challenge.rule_list.rules.spec.body_matcher](data-sources--http_loadbalancer--reference--group-021.md#canonical-9cab2477afc810b9fa13fd67539c09719fe1bf3eb292458ab97d2e5ebb841a86)
- [policy_based_challenge.rule_list.rules.spec.client_selector](data-sources--http_loadbalancer--reference--group-021.md#canonical-80eebafba74da3ae8e1b203e532cfd41f3d0069d94c11525c9c8e1d2ab85eeaa)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](data-sources--http_loadbalancer--reference--group-021.md#canonical-24c1aa264f8eb4a87fc88358f42b3c04cb2afaffb4d457b436297a30fd5c0940)
- [policy_based_challenge.rule_list.rules.spec.disable_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-c87d9460aa1e6756d53b4a5ea7ce14797079b5a1ddfecdd668072b4ca250c352)
- [policy_based_challenge.rule_list.rules.spec.domain_matcher](data-sources--http_loadbalancer--reference--group-021.md#canonical-f5693976020e0fe3642cc5a18d3715bce89e0cc47b0c2c4a6c128e50d8e5f49e)
- [policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-ba606181305fde7d30fc9c0a866fe75fd052c9ac01e6f24d53c403f2bdc85192)
- [policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1ce2ed30433f39b44c880855cae0e8c654804b94382ba72112502693629174e3)
- [policy_based_challenge.rule_list.rules.spec.headers](data-sources--http_loadbalancer--reference--group-021.md#canonical-578f1ac411c848abb3f41924f604ef7400bb2e9d8819058431c6cd4a3724b811)
- [policy_based_challenge.rule_list.rules.spec.http_method](data-sources--http_loadbalancer--reference--group-021.md#canonical-c8d31f5a660171504b9e6f3dda4d7939826a3e3898668d617fa0c00ed87292e3)
- [policy_based_challenge.rule_list.rules.spec.ip_matcher](data-sources--http_loadbalancer--reference--group-021.md#canonical-ca6c31e0e3e39f3e19be14e0bd99fff1219d5beca5dc8e55559dea263cb80880)
- [policy_based_challenge.rule_list.rules.spec.ip_prefix_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-929811c07d5b3b04bed514c363d8cf6d54eb5be41f99307e66bcd8a56a22bb19)
- [policy_based_challenge.rule_list.rules.spec.path](data-sources--http_loadbalancer--reference--group-022.md#canonical-5c5a7d6aa6611cf9b6261f1644001a97fd2267bb79e06050fc298a96b453cc7f)
- [policy_based_challenge.rule_list.rules.spec.query_params](data-sources--http_loadbalancer--reference--group-022.md#canonical-a6dd4dd1d3204263ff38e0262d2b6b68d0e4ca2bb4f2d3300cad7ac24228173e)
- [policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher](data-sources--http_loadbalancer--reference--group-022.md#canonical-0a5c9cf79b02cbcd2d25551795dc18bc688986f470da1226073d5aba1c58f6cb)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-54581841b38cf0239cb482b931035b81f3503053e7ff140def543f35af830ab6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e1bd1521aa57f664ec6918565ec079bbd19dfde054085ca2eec4be13d4d2b8ab"></a>

## policy_based_challenge.rule_list.rules.spec.any_asn — policy_based_challenge.rule_list.rules.spec.any_asn / d019d0142502 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- policy_based_challenge.rule_list.rules.spec.any_asn

<a id="canonical-a82d0cce2120b67a489fbd07fd969b342ce4f4c39feadcf082eb2a63c7322ec4"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-39e7792fdf87b3c163e282ffb56e078a57ba6b22a42820253f83dc6f56049ea2"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.any_asn / d019d0142502 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b467d2305ff53db016f23d1abc3d1a644238dac07368f1070fc86564a7f16ef0"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.any_asn / d019d0142502 / 4

- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9d3e6523437e8b4418a4592c844eeafba14b2dff569461b4891ed01f54eede81"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a08ae74a219652a255c05ce8e0d83ae1373c2587bd4e0adb9547d38dd8a26901"></a>

## policy_based_challenge.rule_list.rules.spec.any_client — policy_based_challenge.rule_list.rules.spec.any_client / 09ce956bff3f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- policy_based_challenge.rule_list.rules.spec.any_client

<a id="canonical-78f9fa055ac6d2f626296733e16d374ace0d03afd6b5ea0df28b1779ac5f92db"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-75011afc8c3bc7eaf03955062bfd65f895e8ecd39f3815538bb337c5fbafd679"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.any_client / 09ce956bff3f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4606d5ba9d8c76b951012e70fc2688b44506b9e504ad1e737a67f7b776e8c9b0"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.any_client / 09ce956bff3f / 4

- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a9b50c0ba4233ac90de12fe29280c3b5592c20ee6ce1843fccd64ddc38d8a5b7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ea0585d445ca3b52c62da275ad7122e1ea016263321b6775fcafad1f5a0fb9c6"></a>

## policy_based_challenge.rule_list.rules.spec.any_ip — policy_based_challenge.rule_list.rules.spec.any_ip / 7f9ad3365b20 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- policy_based_challenge.rule_list.rules.spec.any_ip

<a id="canonical-088f4c632af84f8c3fbad6ee9e999bb43216b9ad7383e91b569090c22d539909"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-ad239257aab856dc1751f1e913e76292a6842df565d28f23a9784fd186ae09da"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.any_ip / 7f9ad3365b20 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-58e5e06e8d64fa5e17bdf3cf43eeed0933bacfc5d7afefc97bc6fa095b259829"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.any_ip / 7f9ad3365b20 / 4

- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-143c1711b2006f094e2001ed44dc4679829cef1ab14645997bdf19783ae956f9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-33831c5999124e9c528a0ae568a6decea4a68e2b56e97295febdf325d5471aea"></a>

## policy_based_challenge.rule_list.rules.spec.arg_matchers — policy_based_challenge.rule_list.rules.spec.arg_matchers / b67750d90511 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- policy_based_challenge.rule_list.rules.spec.arg_matchers

<a id="canonical-73d1aa1574848d750a8a5cdf17e834d6c0a0f146cbd1cc7ffc91f6505647e0b0"></a>

Type: `"list"`. Computed.

List of predicates for all POST args that need to be matched. The criteria for matching each arg are
described in individual instances of ArgMatcherType. The actual arg values are extracted from the
request API as a list of strings for each arg selector name.

Upstream description:

A list of predicates for all POST args that need to be matched. The criteria for matching each arg
are described in individual instances of ArgMatcherType. The actual arg values are extracted from
the request API as a list of strings for each arg selector name. Note that all specified arg matcher
predicates must evaluate to true. A request body greater than 64KB will not be evaluated.

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

<a id="canonical-fae3dfd236da1e680adba10d488722634da9b27dfcb311ba7fee84914268a4c0"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.arg_matchers / b67750d90511 / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-021.md#canonical-dfffc33aa80049c7d4110d2cdb97ec9a58e788ab0e7330ee1f9c9fe6f7a20fa4): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-021.md#canonical-1db0c8f888ee14991dbeb5dc6872d3419ea5648b3346efff59e367b13e4d9cff): complete subsection reference.

<a id="canonical-703e1648d960ecec30861c25a852f299b287d274dda3c6b8d14f6e1efc6bdb4a"></a>

<a id="canonical-c815f17bbc4f45808a5619c077f97c08d7704511903f210fa2dd2fc55352441f"></a>

## invert_matcher property — policy_based_challenge.rule_list.rules.spec.arg_matchers / b67750d90511 / 4

Type: `"bool"`. Computed.

Invert Matcher. Invert Match of the expression defined.

Upstream description:

Invert Match of the expression defined.

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

- [item](data-sources--http_loadbalancer--reference--group-021.md#canonical-530ce18c9bb7166d33b6b5f68887b08dd23c9c6321933608adc54a1fa1218010): complete subsection reference.

<a id="canonical-8d525a0a726857441c864442f0d239f01ead63361456d92c045a4bad9ae79133"></a>

<a id="canonical-d73e9ecebf4fe91e862ceb74044deaa30e06f7bab3c925d732f6156cd6cf1fa9"></a>

## name property — policy_based_challenge.rule_list.rules.spec.arg_matchers / b67750d90511 / 5

Type: `"string"`. Computed.

Case-sensitive JSON path in the HTTP request body.

Upstream description:

A case-sensitive JSON path in the HTTP request body.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.json_path": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.json_path": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-f93510ebdf31f87e60ebcae4a71a9a30e18fd2f2e32bedf928287cb426b2749a"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.arg_matchers / b67750d90511 / 6

- [policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present](data-sources--http_loadbalancer--reference--group-021.md#canonical-dfffc33aa80049c7d4110d2cdb97ec9a58e788ab0e7330ee1f9c9fe6f7a20fa4)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present](data-sources--http_loadbalancer--reference--group-021.md#canonical-1db0c8f888ee14991dbeb5dc6872d3419ea5648b3346efff59e367b13e4d9cff)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers.item](data-sources--http_loadbalancer--reference--group-021.md#canonical-530ce18c9bb7166d33b6b5f68887b08dd23c9c6321933608adc54a1fa1218010)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-dfffc33aa80049c7d4110d2cdb97ec9a58e788ab0e7330ee1f9c9fe6f7a20fa4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-189c3579527dd671d91b8f70b661fed29b505855033541e27d59ccee5d7f62c0"></a>

## policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present — policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present / 9c1aecd89203 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](data-sources--http_loadbalancer--reference--group-021.md#canonical-143c1711b2006f094e2001ed44dc4679829cef1ab14645997bdf19783ae956f9)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present

<a id="canonical-de503755369965d89b400ace96a8670ce49870b42807e15526614c1e1a129f74"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

<a id="canonical-68c04a97e8f8f36c55f0ec7cfc7ae88319b5ba9986a99a33b79037597e4c82d3"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present / 9c1aecd89203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-373abcd46c5e06a2d6cd03d5ef5a81e4e8f4c5520225bcc0656784ff970b21d2"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present / 9c1aecd89203 / 4

- [policy_based_challenge.rule_list.rules.spec.arg_matchers](data-sources--http_loadbalancer--reference--group-021.md#canonical-143c1711b2006f094e2001ed44dc4679829cef1ab14645997bdf19783ae956f9)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1db0c8f888ee14991dbeb5dc6872d3419ea5648b3346efff59e367b13e4d9cff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-87588ee56e3b94afda50940090b7adef49f74d3bf5b12b54a1b6044044eba3f3"></a>

## policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present — policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present / f1a3be1dd359 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](data-sources--http_loadbalancer--reference--group-021.md#canonical-143c1711b2006f094e2001ed44dc4679829cef1ab14645997bdf19783ae956f9)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present

<a id="canonical-3d4024c094c145a45bf698bc8831262026f85f0030d7901f48d64eaeb46fb945"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

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

<a id="canonical-d50cf7f3b54df781838ec0f40e7d813b7964244fc10a902b2b3bc16adc4d766b"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present / f1a3be1dd359 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-604ec642f3f5804989182827fb116f5940af8644422b250abc0d773dfbd1fe36"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present / f1a3be1dd359 / 4

- [policy_based_challenge.rule_list.rules.spec.arg_matchers](data-sources--http_loadbalancer--reference--group-021.md#canonical-143c1711b2006f094e2001ed44dc4679829cef1ab14645997bdf19783ae956f9)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-530ce18c9bb7166d33b6b5f68887b08dd23c9c6321933608adc54a1fa1218010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd9379d16d27b8f8c77127c6b9537d946dccf4cf081090d5cdb8a434d59dc227"></a>

## policy_based_challenge.rule_list.rules.spec.arg_matchers.item — policy_based_challenge.rule_list.rules.spec.arg_matchers.item / 80fc4bb49641 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](data-sources--http_loadbalancer--reference--group-021.md#canonical-143c1711b2006f094e2001ed44dc4679829cef1ab14645997bdf19783ae956f9)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.item

<a id="canonical-ca21daad8dcb0b5c259fb7bc7b510a247d3ad2b60505a975c118aae23b9ae0c1"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

<a id="canonical-cc2a63d518aaca4727dbccd9446b741c4170b4acd223d8c12c20e4bbfb99733c"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.arg_matchers.item / 80fc4bb49641 / 3

<a id="canonical-f14d0c48e6d5842fd04fdfc6bce90a03dd4af29e786cb1c7232ddc5afc8902b8"></a>

<a id="canonical-ad5084a931954eeb7a9c614d5f18e7be3323d946cd847b91e96dc4a372b50148"></a>

## exact_values property — policy_based_challenge.rule_list.rules.spec.arg_matchers.item / 80fc4bb49641 / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

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

<a id="canonical-238cb1a92114e6b2098e85289b2444fc6f845e787491cd6ada88b4507688b4b2"></a>

<a id="canonical-2f3492ea4e0cf58fd2723f290e5ea881f492ff29f6b06c10297cd1cd9deca973"></a>

## regex_values property — policy_based_challenge.rule_list.rules.spec.arg_matchers.item / 80fc4bb49641 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

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

<a id="canonical-1d2552743ef35ea48c38955f573f47c1dd2361332959ef44fe3aed9dcaaf5bda"></a>

<a id="canonical-bd77a64cb9608de0e02f175e9922fa2cc3cd94443a5542307c330f6f009dea8b"></a>

## transformers property — policy_based_challenge.rule_list.rules.spec.arg_matchers.item / 80fc4bb49641 / 6

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0f86827738c75de041212baff5a1481292d9ecdb180ff6754a544dc438d26b39"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.arg_matchers.item / 80fc4bb49641 / 7

- [policy_based_challenge.rule_list.rules.spec.arg_matchers](data-sources--http_loadbalancer--reference--group-021.md#canonical-143c1711b2006f094e2001ed44dc4679829cef1ab14645997bdf19783ae956f9)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-921f66a51c2994c217f6c09c51917f443337e6b545274087c76e81ebf350e7e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-601745e206d904419be1b7dbf494caf4d08a92ce1ae7f33ae4f406462c2655c1"></a>

## policy_based_challenge.rule_list.rules.spec.asn_list — policy_based_challenge.rule_list.rules.spec.asn_list / 12b614f47103 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- policy_based_challenge.rule_list.rules.spec.asn_list

<a id="canonical-ee2801f6e0bc1201f86fa56a6d5f48f791fc91d2b52118415634ceff601ea7c4"></a>

Type: `"single"`. Computed.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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

<a id="canonical-f54b3a3be10c7c3d28919bc42747a8b015bd248a6d5b39a5090df6423df21821"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.asn_list / 12b614f47103 / 3

<a id="canonical-9f1e1506ffbd612388e0213b93591fd0e8b43cd5e338913c5edabdd37168ec7b"></a>

<a id="canonical-6f15b6ab7159dfbb647ded064f5b48fae940b4736b7f6ae868684f8b3ed72c2b"></a>

## as_numbers property — policy_based_challenge.rule_list.rules.spec.asn_list / 12b614f47103 / 4

Type: `["list", "number"]`. Computed.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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

<a id="canonical-810909121dde78ce80571142d9c18544bb7f1af540d6bbde97a54d9092a301b4"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.asn_list / 12b614f47103 / 5

- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-64fcbf7242e313d33bf68832dd18ab4a4c930894aaed6be19061b651ad617d29"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ffdd7fe8f64a3be0434e7a322f37c842d0f3b355645cfa4094a88cece07035f"></a>

## policy_based_challenge.rule_list.rules.spec.asn_matcher — policy_based_challenge.rule_list.rules.spec.asn_matcher / 730873aca886 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- policy_based_challenge.rule_list.rules.spec.asn_matcher

<a id="canonical-e807c7210e4ca0865bf6e44469058385486ebe93412cb87e0e7b2c15c2951bcf"></a>

Type: `"single"`. Computed.

Match any AS number contained in the list of bgp\_asn\_sets.

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

<a id="canonical-515bb02fcc49ba7423cf9c275a9e107b72490703ff54c2f827bd89f4b21642e1"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.asn_matcher / 730873aca886 / 3

- [asn_sets](data-sources--http_loadbalancer--reference--group-021.md#canonical-bc7e3fecf8d8515803002af1b789d5326c806df46875551de48be0040ff0bc5d): complete subsection reference.

<a id="canonical-0c51505d9c4cce78151be3f01421db74dece43b28d9ed347a9bea01fc30ecc7a"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.asn_matcher / 730873aca886 / 4

- [policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets](data-sources--http_loadbalancer--reference--group-021.md#canonical-bc7e3fecf8d8515803002af1b789d5326c806df46875551de48be0040ff0bc5d)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-bc7e3fecf8d8515803002af1b789d5326c806df46875551de48be0040ff0bc5d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95b9342cacf148ac0fb99937d93610e48604f9d1c9bfe39c936a21e6cb097bc7"></a>

## policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets — policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets / 81c665787e1d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [policy_based_challenge.rule_list.rules.spec.asn_matcher](data-sources--http_loadbalancer--reference--group-021.md#canonical-64fcbf7242e313d33bf68832dd18ab4a4c930894aaed6be19061b651ad617d29)
- policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets

<a id="canonical-285edb75ec2c59e40921defa6af1d7e8834306bc796dd2844a90e3bdeb6d1cf2"></a>

Type: `"list"`. Computed.

List of references to bgp\_asn\_set objects.

Upstream description:

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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-de669a6efe608c90366faaf35a4b7f4786738e0383c43e84679bba8ae2e14774"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets / 81c665787e1d / 3

<a id="canonical-c2eb5c5e32a8531877aa76d6cfc4c38d03cbc8a6e2bbf49416c7f3b8228b93a9"></a>

<a id="canonical-8f29b9cc9f7aeae3f2cd239bf002b37be91d3b1e909e5713d3276cec439c9ca8"></a>

## kind property — policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets / 81c665787e1d / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-888b17d5df6cb3b487adf2ac556ab9482686c6e5893e01b81719e80d121c7bf3"></a>

<a id="canonical-8f62db27ffc8b0ca4fb9e0c6a7ec27221f64be2b62e07908d3a97eff7fd79f1b"></a>

## name property — policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets / 81c665787e1d / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-e06ec5de125e723cf0b5b04a0586c7d0d1d428a7c4a12e872f83420847e4df21"></a>

<a id="canonical-60204f1e94f54ea1486cea0c63096fda6101262a9382dfa184a1a5a8693cde8d"></a>

## namespace property — policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets / 81c665787e1d / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-9fb9153022f3e44505523a09de1eedfae8ef3befc31e1698709394e2c6f42448"></a>

<a id="canonical-29341e563dd7498f4caa3c680fb89dcc0f50605d4abf8bca6a026d7715262852"></a>

## tenant property — policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets / 81c665787e1d / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-f5c54aeef6f030bc47ce36d82ceef4e4149b31097037b32a3faf962905a1e5f7"></a>

<a id="canonical-68f5cd2e2296f2938ede18700bd5a24ab1fe2b3b3a857d55aa2239b997bbe60f"></a>

## uid property — policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets / 81c665787e1d / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-8cd03362f3f3c549a970d4dac882ea318c6d13bcd60ce19c38d80f595b334a01"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets / 81c665787e1d / 9

- [policy_based_challenge.rule_list.rules.spec.asn_matcher](data-sources--http_loadbalancer--reference--group-021.md#canonical-64fcbf7242e313d33bf68832dd18ab4a4c930894aaed6be19061b651ad617d29)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9cab2477afc810b9fa13fd67539c09719fe1bf3eb292458ab97d2e5ebb841a86"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b6f07206117e1341a14b3e86e0654ea7733f9915a0c0a45b2be4f32f3231a6ed"></a>

## policy_based_challenge.rule_list.rules.spec.body_matcher — policy_based_challenge.rule_list.rules.spec.body_matcher / 869f105152d8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- policy_based_challenge.rule_list.rules.spec.body_matcher

<a id="canonical-b88d073a789cead0a9c76dcf7598ddc6a1dc1523d5dd4d368bc8b2d6beb81adf"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

<a id="canonical-50a667bf52e843ccb05510e5128039828facd1aa23e57c844ce141ba6810a801"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.body_matcher / 869f105152d8 / 3

<a id="canonical-a3ce60733e7b5f9ea38e01eb95c3b52830d3f30eb620d17d9ec572297f618d5d"></a>

<a id="canonical-d24ffbd15028466e7960057127796d7318b13d9761a237e57672d5f2a70784b7"></a>

## exact_values property — policy_based_challenge.rule_list.rules.spec.body_matcher / 869f105152d8 / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

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

<a id="canonical-0c396dd533a7bdc471513329b7f0316890c6803e5638ec1f8d360bff59e379f1"></a>

<a id="canonical-6e4813982594d1e2fa0c85d14405d5d0cf0a2bea75c661ac5c247ea6d3a322c4"></a>

## regex_values property — policy_based_challenge.rule_list.rules.spec.body_matcher / 869f105152d8 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

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

<a id="canonical-51c04195803c8588580576ea5f33f453fd6795b649e7e49c56e97c1d16499090"></a>

<a id="canonical-c8efde8795b863ef2b2ef488f89c8110e608b8a9d86dd5f9ff060f728805e885"></a>

## transformers property — policy_based_challenge.rule_list.rules.spec.body_matcher / 869f105152d8 / 6

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-48798ad821b05ea44ee096a56a662042f9a06ea9635b58d6ee5d887358b51c90"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.body_matcher / 869f105152d8 / 7

- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-80eebafba74da3ae8e1b203e532cfd41f3d0069d94c11525c9c8e1d2ab85eeaa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3eeb36056b4b0f7715b3b47ad6dc533f596ca79b1da66acba8c8866a969f535d"></a>

## policy_based_challenge.rule_list.rules.spec.client_selector — policy_based_challenge.rule_list.rules.spec.client_selector / 12f6023c0df6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- policy_based_challenge.rule_list.rules.spec.client_selector

<a id="canonical-9cc81520d781531ce8ca96936afca23c7885ace6d8159794d516d9dfb71c733f"></a>

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

<a id="canonical-f0671134607b9fa4eb494cf9d793ed86d3477be4cd0c412d6ebef6b80cc6c603"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.client_selector / 12f6023c0df6 / 3

<a id="canonical-ab29dbd84b5e49bf02d201344cb160dc84000a2ae71495c5e435bb7972b0822f"></a>

<a id="canonical-d52cb5ace67e7cb49646733689aa562e2bd3dfa7edea6591e0e9e3e56d86f14d"></a>

## expressions property — policy_based_challenge.rule_list.rules.spec.client_selector / 12f6023c0df6 / 4

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

<a id="canonical-b0746d8f3c825a5c8df04f5f8f918577c55b43806d05a7fc5334d45c203a485e"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.client_selector / 12f6023c0df6 / 5

- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-24c1aa264f8eb4a87fc88358f42b3c04cb2afaffb4d457b436297a30fd5c0940"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c87c90b4cdabf5798c8e1d1ed2a12214e4246d38d27aa78f64ae30ed498db74a"></a>

## policy_based_challenge.rule_list.rules.spec.cookie_matchers — policy_based_challenge.rule_list.rules.spec.cookie_matchers / c81a870aae61 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers

<a id="canonical-3ca282b6977923dabd73087930a9a71c03346b7b70e2df8edfc8f04e87dd98ee"></a>

Type: `"list"`. Computed.

List of predicates for all cookies that need to be matched. The criteria for matching each cookie is
described in individual instances of CookieMatcherType. The actual cookie values are extracted from
the request API as a list of strings for each cookie name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-be2ff2ee94e9a728ee21deb9930a4cd608d46db371f424161446f37941892a3b"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.cookie_matchers / c81a870aae61 / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-021.md#canonical-0daaa24c5bb3aa51b6aa5fb8b7df5c3915d306750aaeab02e565444f52ccc979): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-021.md#canonical-014a5ba0cf07c1aaba783dcd3bffad7f1a284185f049e75167a87b870e5aede5): complete subsection reference.

<a id="canonical-75d82e183afdc02fedc51910372d0f2cb442b59fe7bb650621265a2a2ef974c4"></a>

<a id="canonical-078b9459234949239eaa5fe067caefd594b44242a191104037d256e18450c4ac"></a>

## invert_matcher property — policy_based_challenge.rule_list.rules.spec.cookie_matchers / c81a870aae61 / 4

Type: `"bool"`. Computed.

Invert Matcher. Invert Match of the expression defined.

Upstream description:

Invert Match of the expression defined.

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

- [item](data-sources--http_loadbalancer--reference--group-021.md#canonical-e22ce807960168c8fae6db25a9c3b5b60774074b0f734a094238c8c00003804e): complete subsection reference.

<a id="canonical-06ea20ac6a623fcc7d9ce14ed6f8c1f978b7b52631e87d75b92db1bf70b5cdde"></a>

<a id="canonical-71e797121d33a12e928554afec28695249f82b410512a14031d6689908a8ded5"></a>

## name property — policy_based_challenge.rule_list.rules.spec.cookie_matchers / c81a870aae61 / 5

Type: `"string"`. Computed.

Cookie Name. A case-sensitive cookie name.

Upstream description:

A case-sensitive cookie name.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-c9d4e0990100f384a786f73e2850513e5b540beb071384da1eac00acee5d8f19"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.cookie_matchers / c81a870aae61 / 6

- [policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present](data-sources--http_loadbalancer--reference--group-021.md#canonical-0daaa24c5bb3aa51b6aa5fb8b7df5c3915d306750aaeab02e565444f52ccc979)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present](data-sources--http_loadbalancer--reference--group-021.md#canonical-014a5ba0cf07c1aaba783dcd3bffad7f1a284185f049e75167a87b870e5aede5)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers.item](data-sources--http_loadbalancer--reference--group-021.md#canonical-e22ce807960168c8fae6db25a9c3b5b60774074b0f734a094238c8c00003804e)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-0daaa24c5bb3aa51b6aa5fb8b7df5c3915d306750aaeab02e565444f52ccc979"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2bd5955f526248a018148211267774371b45c1addac3974bd9751245f707701"></a>

## policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present — policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present / cd771e84c4d4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](data-sources--http_loadbalancer--reference--group-021.md#canonical-24c1aa264f8eb4a87fc88358f42b3c04cb2afaffb4d457b436297a30fd5c0940)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present

<a id="canonical-02e9f78229dd19c7a66a3d01e77d9dfa192baa319f4d48dddc6c090b4a07281c"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

<a id="canonical-0e5a4643e64ba28163421b1a5ea8de7fa45dcd54c5f4ca9a41d0c8e8b2ea5e10"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present / cd771e84c4d4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d68f4be5a3e11d4cec76a20e758cd81aeaa1e30253a7cea002acb7281d8edaea"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present / cd771e84c4d4 / 4

- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](data-sources--http_loadbalancer--reference--group-021.md#canonical-24c1aa264f8eb4a87fc88358f42b3c04cb2afaffb4d457b436297a30fd5c0940)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-014a5ba0cf07c1aaba783dcd3bffad7f1a284185f049e75167a87b870e5aede5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d06f6b5d374f4689414c384f2d434e28c806cee6f0244b722eb495e0c7fa92b8"></a>

## policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present — policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present / 0bcbb985bdad / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](data-sources--http_loadbalancer--reference--group-021.md#canonical-24c1aa264f8eb4a87fc88358f42b3c04cb2afaffb4d457b436297a30fd5c0940)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present

<a id="canonical-84326a0a8f4e6f7bed8e739f14884f70bc8753d98131f1756a174c0c2f2b6e98"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

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

<a id="canonical-5bf4711c0da2a32e19aa9bcad012294d92c9d4a16c5dba50f4e3b1ce6eb0f3fb"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present / 0bcbb985bdad / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0f213bf5c8ee3778c134630ac2d4578e15745611f9fd47cca49eff36ae0d1069"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present / 0bcbb985bdad / 4

- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](data-sources--http_loadbalancer--reference--group-021.md#canonical-24c1aa264f8eb4a87fc88358f42b3c04cb2afaffb4d457b436297a30fd5c0940)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e22ce807960168c8fae6db25a9c3b5b60774074b0f734a094238c8c00003804e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5b74c7aca8126797b42a5450ecdb5d081cb4cd6b01f279477b3db51849b8600"></a>

## policy_based_challenge.rule_list.rules.spec.cookie_matchers.item — policy_based_challenge.rule_list.rules.spec.cookie_matchers.item / ada63fcb325a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](data-sources--http_loadbalancer--reference--group-021.md#canonical-24c1aa264f8eb4a87fc88358f42b3c04cb2afaffb4d457b436297a30fd5c0940)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.item

<a id="canonical-ef6f4990871842cad666dfad27f44507dfd6ccd8c564f95e4d9d8121757aab22"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

<a id="canonical-d6640af381d4718e90b70f1de021347657e2910511e554206d034dffccb52b07"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.cookie_matchers.item / ada63fcb325a / 3

<a id="canonical-8b721b9fbf686c6eab77f427a171c2e602187d207ecd52b2565fe6dd80d3df44"></a>

<a id="canonical-a7db580ed2657f239a242a5503b0540f5a3dfe3d7b0e2ff15e1ddf58b99c5c9e"></a>

## exact_values property — policy_based_challenge.rule_list.rules.spec.cookie_matchers.item / ada63fcb325a / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

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

<a id="canonical-9240bfb24e846d1252fc891f16309e81bb0fee7d10ef6f62a22b1808c2ca317c"></a>

<a id="canonical-e78b024b90c0fd6ec324da5f9085e197fdc6c114796f10b07379623678930552"></a>

## regex_values property — policy_based_challenge.rule_list.rules.spec.cookie_matchers.item / ada63fcb325a / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

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

<a id="canonical-4dfa3ffe9f43e4c3429b0820c22c92cfb3c7b4432678dd24d86d0e4faa2d7074"></a>

<a id="canonical-df8f39624835b925bffd320dbc472867b4fc5484af2c9c9cf68c1950041bf885"></a>

## transformers property — policy_based_challenge.rule_list.rules.spec.cookie_matchers.item / ada63fcb325a / 6

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-32f89f188f4f4be53598611c6116e24cffae3fbdac60a5cb8a052423f526c494"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.cookie_matchers.item / ada63fcb325a / 7

- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](data-sources--http_loadbalancer--reference--group-021.md#canonical-24c1aa264f8eb4a87fc88358f42b3c04cb2afaffb4d457b436297a30fd5c0940)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c87d9460aa1e6756d53b4a5ea7ce14797079b5a1ddfecdd668072b4ca250c352"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-251fa06269e87cf2acb1a99bc334a7ab70fa7f05bc46d6d459179ea18445c492"></a>

## policy_based_challenge.rule_list.rules.spec.disable_challenge — policy_based_challenge.rule_list.rules.spec.disable_challenge / 1362b811b48d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- policy_based_challenge.rule_list.rules.spec.disable_challenge

<a id="canonical-144d52523be848df9ede417cb920255c4e85ccf542cba90cbb1281c80634fba0"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable challenge.

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

<a id="canonical-8119765f2ca3d11740727a7c1fe50e59fa55f17499a9784b75009f3376ed3e8e"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.disable_challenge / 1362b811b48d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b8168d1b1f2c1f744cff8dd0ca598d788dddee93584b494dde407ce9e52fd6c0"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.disable_challenge / 1362b811b48d / 4

- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f5693976020e0fe3642cc5a18d3715bce89e0cc47b0c2c4a6c128e50d8e5f49e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-17f182c1f24427af7f114ff3d5963a17082f16a201c6875b1a7db22b71b46140"></a>

## policy_based_challenge.rule_list.rules.spec.domain_matcher — policy_based_challenge.rule_list.rules.spec.domain_matcher / cc9e117da38d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- policy_based_challenge.rule_list.rules.spec.domain_matcher

<a id="canonical-0a0f179886f340ad8383925921ec2044f3871830fe29ff4021384990ced5d8a3"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

<a id="canonical-e736b4a17e208b276264cc97dd0c262c27b6f048d7c275730faada6de89f4c4a"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.domain_matcher / cc9e117da38d / 3

<a id="canonical-fbdfe3ea5e58daacc7609ddb50b7d37a7ea995c78ef141e3e5ec0896955339b2"></a>

<a id="canonical-7c40d637966cb7712c1f0873ffb30944e7a8941cfef788a164a57fe3fe5f961d"></a>

## exact_values property — policy_based_challenge.rule_list.rules.spec.domain_matcher / cc9e117da38d / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

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

<a id="canonical-8a242b43ef168daf2e2b15eedf61d7f6fef84098f402a5623f7459f2e8119c3a"></a>

<a id="canonical-01efebfb7c1fbc9d87a6f4d4d88613f28fb52888b809bbd5de9989231cd8d852"></a>

## regex_values property — policy_based_challenge.rule_list.rules.spec.domain_matcher / cc9e117da38d / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

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

<a id="canonical-f973f2835c7def16a0492c52eab5884b8696748614fccf37476990fb2bbabe3b"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.domain_matcher / cc9e117da38d / 6

- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ba606181305fde7d30fc9c0a866fe75fd052c9ac01e6f24d53c403f2bdc85192"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4b313b5bcdd009981a75842df5751bb0d7c4bf89ad8f83ad589e106fb02acc0d"></a>

## policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge — policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge / 76e783afc162 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge

<a id="canonical-602773191b2a930badea0c2b4f5e2747d54210a7c65264c87779c14e430eb0ec"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable captcha challenge.

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

<a id="canonical-05a9bd067a5c148d07d3770817022f71a94601803dcc4f513c583af4b75742bc"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge / 76e783afc162 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-215ca40b0d2c1d80b42022251003e9cc14e30f6ec4c18777b7308017de65720a"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge / 76e783afc162 / 4

- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1ce2ed30433f39b44c880855cae0e8c654804b94382ba72112502693629174e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-29f8d09e5ef8070a63426413bf0bd5e6cf1046a29f99404a62381d058181db25"></a>

## policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge — policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge / 9f36e4b4b46e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge

<a id="canonical-0840df3c7e63121380fad2695fb7a8c8f04f1879b8c19073ac5cc46266fd49c9"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-5e2e40bb1b1ff849b6380af65eb8806df0621895767987b88ae079d29abf36f7"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge / 9f36e4b4b46e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4d4cee12304a2f7cdcf94fef1f6b1f18f541a77e8858a8bc8f332a355e7214d6"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge / 9f36e4b4b46e / 4

- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-578f1ac411c848abb3f41924f604ef7400bb2e9d8819058431c6cd4a3724b811"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-722e313c6235b676a14e4e0918e7de80cae20970ee15a5886804d56eb442c5c6"></a>

## policy_based_challenge.rule_list.rules.spec.headers — policy_based_challenge.rule_list.rules.spec.headers / f7402bb85b5d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- policy_based_challenge.rule_list.rules.spec.headers

<a id="canonical-751944538a48399fbd6ab28ccdff32cc80d660ef39a5f6c11abcc1b496725957"></a>

Type: `"list"`. Computed.

List of predicates for various HTTP headers that need to match. The criteria for matching each HTTP
header are described in individual HeaderMatcherType instances. The actual HTTP header values are
extracted from the request API as a list of strings for each HTTP header type.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-7ca921184683d1c23923f964201e74083adaa347a049bb2ba6087e0486a54a7f"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.headers / f7402bb85b5d / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-021.md#canonical-fbd45ab88c1fbe98c420aacb4f1acca29bb84df6b3e4509ccfc8e3e041d5a3a9): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-021.md#canonical-a4aff1719a4def2790eda3ba375d64a267c8b3819f42d9fea28d27387160e898): complete subsection reference.

<a id="canonical-90ff966fab578bc9c887676fb0284c8ca9f17ac54cd873dfe84bfd7b31a51a48"></a>

<a id="canonical-3dc46ba13161387dd78563f6336c7410c3bf25f736a8465ce4bbd751b8fe5faa"></a>

## invert_matcher property — policy_based_challenge.rule_list.rules.spec.headers / f7402bb85b5d / 4

Type: `"bool"`. Computed.

Invert Header Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [item](data-sources--http_loadbalancer--reference--group-021.md#canonical-06f2f3f44fe27478f42f724d0850d972647480fba308e1d00ba1bfbed94fe46f): complete subsection reference.

<a id="canonical-1bbcd949735f815933f5fd1db730b0f1ed07a1c395c3b74d303ff88b0ce7c18b"></a>

<a id="canonical-33d737679864e329305f59991e244d156e876404604561ce98804a53e2114af6"></a>

## name property — policy_based_challenge.rule_list.rules.spec.headers / f7402bb85b5d / 5

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

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

<a id="canonical-3ae1fc0147039bb0de979f2ea7a50dae41b9ca607f75796be7dea80f5685acdc"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.headers / f7402bb85b5d / 6

- [policy_based_challenge.rule_list.rules.spec.headers.check_not_present](data-sources--http_loadbalancer--reference--group-021.md#canonical-fbd45ab88c1fbe98c420aacb4f1acca29bb84df6b3e4509ccfc8e3e041d5a3a9)
- [policy_based_challenge.rule_list.rules.spec.headers.check_present](data-sources--http_loadbalancer--reference--group-021.md#canonical-a4aff1719a4def2790eda3ba375d64a267c8b3819f42d9fea28d27387160e898)
- [policy_based_challenge.rule_list.rules.spec.headers.item](data-sources--http_loadbalancer--reference--group-021.md#canonical-06f2f3f44fe27478f42f724d0850d972647480fba308e1d00ba1bfbed94fe46f)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-fbd45ab88c1fbe98c420aacb4f1acca29bb84df6b3e4509ccfc8e3e041d5a3a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-542c7c3075fc8d6fabea3f9850f0dd007ebaa5c16e353395c09f5e843f144f75"></a>

## policy_based_challenge.rule_list.rules.spec.headers.check_not_present — policy_based_challenge.rule_list.rules.spec.headers.check_not_present / 7e0859e3d558 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [policy_based_challenge.rule_list.rules.spec.headers](data-sources--http_loadbalancer--reference--group-021.md#canonical-578f1ac411c848abb3f41924f604ef7400bb2e9d8819058431c6cd4a3724b811)
- policy_based_challenge.rule_list.rules.spec.headers.check_not_present

<a id="canonical-90e48be415474a0d28c89b463c0e7707c8b4855a255bb9ea2aa5d091a1dd5bad"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

<a id="canonical-d2bbb633691a0334b129fe352b057bbceb17721d6734f69cf6f9dab7a7fca986"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.headers.check_not_present / 7e0859e3d558 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-83327bf0d8242d1a03daa0f0b51207f044c8c5cb56ab0b3faa23bccbc644b557"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.headers.check_not_present / 7e0859e3d558 / 4

- [policy_based_challenge.rule_list.rules.spec.headers](data-sources--http_loadbalancer--reference--group-021.md#canonical-578f1ac411c848abb3f41924f604ef7400bb2e9d8819058431c6cd4a3724b811)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a4aff1719a4def2790eda3ba375d64a267c8b3819f42d9fea28d27387160e898"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-018f70f7e3d9d10a68c7c178bb8a89dd6215a79c9c8757ac088f1c0763771790"></a>

## policy_based_challenge.rule_list.rules.spec.headers.check_present — policy_based_challenge.rule_list.rules.spec.headers.check_present / 51ea37d1b14b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [policy_based_challenge.rule_list.rules.spec.headers](data-sources--http_loadbalancer--reference--group-021.md#canonical-578f1ac411c848abb3f41924f604ef7400bb2e9d8819058431c6cd4a3724b811)
- policy_based_challenge.rule_list.rules.spec.headers.check_present

<a id="canonical-d858efc993c7d1f30e7eda0a1f4e2b8d77c4dee9c55a6c6588a7fdf4de74747a"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

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

<a id="canonical-5d527ceafc98141fa4b23ce212334faa6539d1c26c0af70fae96b95bdb35b588"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.headers.check_present / 51ea37d1b14b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-410e3073e32fa443482ef292a0a4b0ff35397565593a114e46dc411a3c594a08"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.headers.check_present / 51ea37d1b14b / 4

- [policy_based_challenge.rule_list.rules.spec.headers](data-sources--http_loadbalancer--reference--group-021.md#canonical-578f1ac411c848abb3f41924f604ef7400bb2e9d8819058431c6cd4a3724b811)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-06f2f3f44fe27478f42f724d0850d972647480fba308e1d00ba1bfbed94fe46f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-201cfb0f66fa39fe890f26e266be29a72d3dcf1092e87880d3b3842f703c1f00"></a>

## policy_based_challenge.rule_list.rules.spec.headers.item — policy_based_challenge.rule_list.rules.spec.headers.item / b43d9aeb4f29 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [policy_based_challenge.rule_list.rules.spec.headers](data-sources--http_loadbalancer--reference--group-021.md#canonical-578f1ac411c848abb3f41924f604ef7400bb2e9d8819058431c6cd4a3724b811)
- policy_based_challenge.rule_list.rules.spec.headers.item

<a id="canonical-b5578991d31f260ae565798bb54aaf19a608c48bc0bc6596933852945afa2219"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

<a id="canonical-f3e49d4c333ed1b2339af179a16277b94653b95d05f52d041a8d9515c614784c"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.headers.item / b43d9aeb4f29 / 3

<a id="canonical-b6c171cd89fdb4e5031474a4a4e850027e2e48e88b972a7febe8f6804b3b58aa"></a>

<a id="canonical-6e1d77e1af38a8d94facefbe61b2e7193017abad0872cd48c1ad35a0e6f30012"></a>

## exact_values property — policy_based_challenge.rule_list.rules.spec.headers.item / b43d9aeb4f29 / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

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

<a id="canonical-243a677f72b7ca393b9c14573d0200462375220f224db8e360aa543f3b4d3842"></a>

<a id="canonical-56e358315cdfb02f084981021f17259aacc7c1c2f78b47034b6a47e980ace41f"></a>

## regex_values property — policy_based_challenge.rule_list.rules.spec.headers.item / b43d9aeb4f29 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

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

<a id="canonical-8f2719b4df823ab70cd115dbc81b54a5bd3cc37b1deb701282c1c7bf7927f9d9"></a>

<a id="canonical-8bd2602d67ba54d5f8cd81d385b9a79737e60e2fa9d9bba4d93910a98a95023f"></a>

## transformers property — policy_based_challenge.rule_list.rules.spec.headers.item / b43d9aeb4f29 / 6

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-90c7a4028e8ca3fa003b3217e1a5ddc7dfaccc9cac98568d06ee346a3b74888b"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.headers.item / b43d9aeb4f29 / 7

- [policy_based_challenge.rule_list.rules.spec.headers](data-sources--http_loadbalancer--reference--group-021.md#canonical-578f1ac411c848abb3f41924f604ef7400bb2e9d8819058431c6cd4a3724b811)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c8d31f5a660171504b9e6f3dda4d7939826a3e3898668d617fa0c00ed87292e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6c80d1efee6a4cf8175af570aba17f8c4a43802d1febae3c4425eef38590d203"></a>

## policy_based_challenge.rule_list.rules.spec.http_method — policy_based_challenge.rule_list.rules.spec.http_method / d13242ac868b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- policy_based_challenge.rule_list.rules.spec.http_method

<a id="canonical-ae0b9293b647fd667cc8a06ece52255dc15f1363229929316cbada40db537c73"></a>

Type: `"single"`. Computed.

HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
considered successful if the input method is a member of the list. The result of the match based on
the method list is inverted if invert\_matcher is true.

Upstream description:

A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
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

<a id="canonical-b2d5fc40fe25e4833ab220f31c87962f037bcdc32c65304c055fc638f145a466"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.http_method / d13242ac868b / 3

<a id="canonical-392045e1784713127ed1a1276bac7529f7f22167ec2157316321a82e23e47cf8"></a>

<a id="canonical-9888c814f0a67228f03e31dfe18d0e89b58a6ada0bc50dda8f38c9170ec391aa"></a>

## invert_matcher property — policy_based_challenge.rule_list.rules.spec.http_method / d13242ac868b / 4

Type: `"bool"`. Computed.

Invert Method Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

<a id="canonical-b75d3cdc859a1cb492c7080b3ea4d3897c90f5a75bc37d49762933b7a4987c86"></a>

<a id="canonical-4874400dbd798eebb494cbdb034ef701e572ba5e9fbe24092d6da86b9533d871"></a>

## methods property — policy_based_challenge.rule_list.rules.spec.http_method / d13242ac868b / 5

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] List of methods values to
match against. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`,
\`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

List of methods values to match against.

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

<a id="canonical-c63409e174e9bfc10e1a0520a098c2b6948f3febf93176e2f631d8edc47b956e"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.http_method / d13242ac868b / 6

- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ca6c31e0e3e39f3e19be14e0bd99fff1219d5beca5dc8e55559dea263cb80880"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4aec81a59aff5dc1e0270ea78ae4739712ff5d8761b5b28c2a35fa8d08010dac"></a>

## policy_based_challenge.rule_list.rules.spec.ip_matcher — policy_based_challenge.rule_list.rules.spec.ip_matcher / c48a630e0f6a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- policy_based_challenge.rule_list.rules.spec.ip_matcher

<a id="canonical-5065331c0a1b37ef42c9ba25cb02e1447f1834b34a4b6437d9bf01b5936f5dc6"></a>

Type: `"single"`. Computed.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

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

<a id="canonical-a5d57bb9723eb57644693d2749cac23face2ecd6ec5b42c8988284e2c20d1f98"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.ip_matcher / c48a630e0f6a / 3

<a id="canonical-e48e53ec8c0eb887c75c76c6c617972550c76acd6438eddf3dce27f299c24fb7"></a>

<a id="canonical-3dbad6b2ed9f749b5acd0f199fc4e4f47cb9ff791c16aeefe283fb78ed44581f"></a>

## invert_matcher property — policy_based_challenge.rule_list.rules.spec.ip_matcher / c48a630e0f6a / 4

Type: `"bool"`. Computed.

Invert IP Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [prefix_sets](data-sources--http_loadbalancer--reference--group-021.md#canonical-f9f2b16097a5504b447b3fe7274a5bca276cf6f612e55813f4ae1ba385d70218): complete subsection reference.

<a id="canonical-5117d92069358a00012702ef999e5bd407ae7580f25dd6236443f5d9fac4dc93"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.ip_matcher / c48a630e0f6a / 5

- [policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets](data-sources--http_loadbalancer--reference--group-021.md#canonical-f9f2b16097a5504b447b3fe7274a5bca276cf6f612e55813f4ae1ba385d70218)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f9f2b16097a5504b447b3fe7274a5bca276cf6f612e55813f4ae1ba385d70218"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
