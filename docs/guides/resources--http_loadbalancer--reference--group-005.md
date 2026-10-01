---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-91f0771938c380ac54367a71791c2e1f66c8f261313bf6b662e336b61157595f"></a>

## as_numbers property — api_protection_rules.api_endpoint_rules.client_matcher.asn_list / abe1fe3b72bf / 4

Type: `["list", "number"]`. Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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

<a id="canonical-8a7bdd2ed5d20178704c556305a5b8b5ea0c6d2f4e7a4923ac2694a7a9feead7"></a>

## Next pages — api_protection_rules.api_endpoint_rules.client_matcher.asn_list / abe1fe3b72bf / 5

- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-b0510836a500354619afdae6bb3e5e3e45631fa4a9bb22642c7bc24bd6a94961)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-0dacfec7c47d335e81f43c27619b2bd4e10448e0bd896c4ae691e7cd09b74638"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-04c36eb80f4aec86d6ac1a8c8e38cfb6cd1b94e01838234ded22539005931c1f"></a>

## api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher — api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher / a208711b3801 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-7bc735c85b50beaf50f1b45bdbe15c413d15aa714c8844dc469693a74dd86737)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-b0510836a500354619afdae6bb3e5e3e45631fa4a9bb22642c7bc24bd6a94961)
- api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher

<a id="canonical-f1340ab2ad1f65e28d12541fb723ed01799e63b4325000bd2c1afba1138fd1a6"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("asn_sets")}
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
asn_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-51db71526982e2e72ef479610e27085a10a59384863b58d03d505739674ff54f"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher / a208711b3801 / 3

- [asn_sets](resources--http_loadbalancer--reference--group-005.md#canonical-96d3156592fb92d107da88e7c344fea4bab931a0cba46f20395e1c618fb8e99a): complete subsection reference.

<a id="canonical-d2a23577bae0bc85a36578fd22f9f13d183eccaf3eefa5ef1a3e38d4083e18a1"></a>

## Next pages — api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher / a208711b3801 / 4

- [api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets](resources--http_loadbalancer--reference--group-005.md#canonical-96d3156592fb92d107da88e7c344fea4bab931a0cba46f20395e1c618fb8e99a)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-b0510836a500354619afdae6bb3e5e3e45631fa4a9bb22642c7bc24bd6a94961)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-96d3156592fb92d107da88e7c344fea4bab931a0cba46f20395e1c618fb8e99a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ac19749880d4b6578775612b5c1954d196344b266cb5f0225a2070e469ec13e"></a>

## api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets — api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets / 858913ae7641 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-7bc735c85b50beaf50f1b45bdbe15c413d15aa714c8844dc469693a74dd86737)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-b0510836a500354619afdae6bb3e5e3e45631fa4a9bb22642c7bc24bd6a94961)
- [api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0dacfec7c47d335e81f43c27619b2bd4e10448e0bd896c4ae691e7cd09b74638)
- api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-b723f419e5183caf9e83659f6e73c015afc50a52a219ea2d54feeb2383f9fcbd"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
asn_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-461ae9f68a518340c6673a34c4e3d97bb8d3be323738a1896ae9ec784325de3d"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets / 858913ae7641 / 3

<a id="canonical-042e0e399d45a1f7994ce59e97d2fa24fa7bd94fc0472ee124939e8b19eb8463"></a>

<a id="canonical-8682cdae5db9d9a15febd4d8442ce2e14ab2615b7e48fc54474396557c459b62"></a>

## kind property — api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets / 858913ae7641 / 4

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

<a id="canonical-a543936a6e86063f6c4d2430723762fd583e4b74fa3355294fa9c10382f5344e"></a>

<a id="canonical-5cec049d7c1c1e49da71b9b5974df2d3be29dd09cb066f23e8c358b8ed1f0480"></a>

## name property — api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets / 858913ae7641 / 5

Type: `"string"`. Optional.

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

<a id="canonical-97b74a43f1c92f5fd46a8286678bc1f7cc35ca47f82eeef5e679a1e471f47b06"></a>

<a id="canonical-5643fda2d78859d93b1a3b3217b634d47668022e9c837eedb2c4753db9640a26"></a>

## namespace property — api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets / 858913ae7641 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-dd1ce3f6f0cd4ddf0bd3f4602813f1fe9165b79eeb1e531d1f408da4c022cba4"></a>

<a id="canonical-74266b5053c46b7da05b1c9f59544e49c7bdbe33eb1b5c42054262db1bf4a08c"></a>

## tenant property — api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets / 858913ae7641 / 7

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

<a id="canonical-bcfd7bbffec2032c701214713ea32b099caa248c39283827671da36534ba237f"></a>

<a id="canonical-3f2477013ad2b2647bf8501cd71986542221752632208289205b772e9705a476"></a>

## uid property — api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets / 858913ae7641 / 8

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

<a id="canonical-9ed6ec0e381e1f1fadaca4835585226be17c7fbe9f670607cd5b8adbacbfca92"></a>

## Next pages — api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets / 858913ae7641 / 9

- [api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0dacfec7c47d335e81f43c27619b2bd4e10448e0bd896c4ae691e7cd09b74638)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-5dda8ddf06faaa4593084d5c3d93030e1ac77fb2952b2b9ad20bd08efec8fa39"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ba503ef78674086f6f6d98bc69b1f8ddbed489ecec99a280ade343b367fefdd"></a>

## api_protection_rules.api_endpoint_rules.client_matcher.client_selector — api_protection_rules.api_endpoint_rules.client_matcher.client_selector / 13c22c980825 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-7bc735c85b50beaf50f1b45bdbe15c413d15aa714c8844dc469693a74dd86737)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-b0510836a500354619afdae6bb3e5e3e45631fa4a9bb22642c7bc24bd6a94961)
- api_protection_rules.api_endpoint_rules.client_matcher.client_selector

<a id="canonical-caabb528e037ce2d0d962158f92b918afaa9f6473203dffe347bb8fd463482d5"></a>

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
client_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-664655e5fefc8a96103b9fd7aad7f99bad4c4808201c91d8f93bd2b2dc04db8a"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.client_matcher.client_selector / 13c22c980825 / 3

<a id="canonical-ac91905063723e2fabfb31f06b213dcc6115c5eddeb0607978082a4d84a51849"></a>

<a id="canonical-8a622aecd8baeef51c74d4b25f7dc3a58f261ab00da1a7b91f2e0d5dd1d0ebac"></a>

## expressions property — api_protection_rules.api_endpoint_rules.client_matcher.client_selector / 13c22c980825 / 4

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

<a id="canonical-93f367ca202146379829808e398e4ab9e7cf605efad909ffb09e9826109d44e9"></a>

## Next pages — api_protection_rules.api_endpoint_rules.client_matcher.client_selector / 13c22c980825 / 5

- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-b0510836a500354619afdae6bb3e5e3e45631fa4a9bb22642c7bc24bd6a94961)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e4394b89e8196774897d8c2df9b93b2a6427c7bfd90ab790aa5ae7376b4247c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d825bf7d4eee2cdf4af8ff41b42e2159a6b32cc1f9908629586b83a5dc8a7212"></a>

## api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher — api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher / 9e4b80d7a980 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-7bc735c85b50beaf50f1b45bdbe15c413d15aa714c8844dc469693a74dd86737)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-b0510836a500354619afdae6bb3e5e3e45631fa4a9bb22642c7bc24bd6a94961)
- api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher

<a id="canonical-8bad03f9613bcd0f060162244342616923cfaad4e1a5c51803ecd4c76421c3be"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

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

<a id="canonical-3fba3df3210ff9f4e642d31518da4f012c4999a0b58dd1751144a72db417aa91"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher / 9e4b80d7a980 / 3

<a id="canonical-1cc7751293cf54bcc5bba48f194aee127252698d703ab01368f2b443c099d63e"></a>

<a id="canonical-68384cbc22e8d052e14a3ef41afb5f4f801c2866790496ecb08fa8c9df096729"></a>

## invert_matcher property — api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher / 9e4b80d7a980 / 4

Type: `"bool"`. Optional.

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

- [prefix_sets](resources--http_loadbalancer--reference--group-005.md#canonical-a94cdc50a5eea47acddb8752cc3ee6b45e04e2f8e00f371c5c3916efaebf468c): complete subsection reference.

<a id="canonical-c6ce73fa3fef2eec4c47defbd96d180261c86257ed73171e8b6ac64397357196"></a>

## Next pages — api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher / 9e4b80d7a980 / 5

- [api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets](resources--http_loadbalancer--reference--group-005.md#canonical-a94cdc50a5eea47acddb8752cc3ee6b45e04e2f8e00f371c5c3916efaebf468c)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-b0510836a500354619afdae6bb3e5e3e45631fa4a9bb22642c7bc24bd6a94961)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-a94cdc50a5eea47acddb8752cc3ee6b45e04e2f8e00f371c5c3916efaebf468c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6cb240cc2fd97801e8c25e052c4ef98a08a2b4c876bc63c8f9f4d4371a77b87d"></a>

## api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets — api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / 7bd91f00fe0f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-7bc735c85b50beaf50f1b45bdbe15c413d15aa714c8844dc469693a74dd86737)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-b0510836a500354619afdae6bb3e5e3e45631fa4a9bb22642c7bc24bd6a94961)
- [api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-e4394b89e8196774897d8c2df9b93b2a6427c7bfd90ab790aa5ae7376b4247c0)
- api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-0b26d5129d76e74cb9a330654b57c79c4b61b589f7f41776a14e0c5242e33c49"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
prefix_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-7be50fbb75365415e8fc2c0d8f571134cda663b5c62f627d4c7d1f648ac5463d"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / 7bd91f00fe0f / 3

<a id="canonical-e4653b1e79a111f61a643f6f605470bc3de950d11ee569d6f636ac4bf181a214"></a>

<a id="canonical-08c4b25e22746296f810a154c1809975036882ed245b979258d69a8cef5bb858"></a>

## kind property — api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / 7bd91f00fe0f / 4

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

<a id="canonical-00db8f9f55a5f4a5a3ed07ab5a8abbef322e30b990c571d7f38de45a93a7617f"></a>

<a id="canonical-d159a08db852d3c001550c55975c49ae161b25a6662843094ed3a3d97118ba41"></a>

## name property — api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / 7bd91f00fe0f / 5

Type: `"string"`. Optional.

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

<a id="canonical-f446465ea829a2ba897364268dd8f2004e10c87b116c54bb0d5bc60411e0592a"></a>

<a id="canonical-f1ff97d63eb8dbca41228d8a643fcb39502357d11d727cd616b217e38b1fedb9"></a>

## namespace property — api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / 7bd91f00fe0f / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-ee6b296b6a469b528672257437f9eb0c0a62265e33d8421b759416804052f1a3"></a>

<a id="canonical-b43683d3217c1a10d5041b826b670233a8e77ffaa573d831fbd4554f745853b3"></a>

## tenant property — api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / 7bd91f00fe0f / 7

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

<a id="canonical-1ec247879df0d3198e51e75b4bad26e89f9efe3128aa57c0340e0120db91c739"></a>

<a id="canonical-f78628d93b9dad285832c6882115fea5e0525a8cac0b39fe198924d4370fb5fe"></a>

## uid property — api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / 7bd91f00fe0f / 8

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

<a id="canonical-6218fdea4c29fa6ee8084e8f6c6431ce88f58dd916c8ee529702fc3cbd9e65cf"></a>

## Next pages — api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / 7bd91f00fe0f / 9

- [api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-e4394b89e8196774897d8c2df9b93b2a6427c7bfd90ab790aa5ae7376b4247c0)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-5f6242d902d15f7c8844e3ea59f43a3377bb6dd414ad7a5a397975ab6a36d0b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8584549811d3ab117eadc6248d94e46ebe14f3fb72cdbb06f78c37f911201a17"></a>

## api_protection_rules.api_endpoint_rules.client_matcher.ip_prefix_list — api_protection_rules.api_endpoint_rules.client_matcher.ip_prefix_list / 88aec7f13f8e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-7bc735c85b50beaf50f1b45bdbe15c413d15aa714c8844dc469693a74dd86737)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-b0510836a500354619afdae6bb3e5e3e45631fa4a9bb22642c7bc24bd6a94961)
- api_protection_rules.api_endpoint_rules.client_matcher.ip_prefix_list

<a id="canonical-67de789596b68227d33965e198bfab1b0c254ea344584d2499abcbb9bceab175"></a>

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

<a id="canonical-834cb70b1104ff06d30c4091f0dd3e7da49afab2ff1d130d1a5a60abb886d9c1"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.client_matcher.ip_prefix_list / 88aec7f13f8e / 3

<a id="canonical-551556b56b01b145397664ea654c2457d55a3ddfc33e1153a2f6891df113b44a"></a>

<a id="canonical-489140a1c0d6b0b29c9397d11b50eaf644e95e741d27055d0208e8c6daf11e7b"></a>

## invert_match property — api_protection_rules.api_endpoint_rules.client_matcher.ip_prefix_list / 88aec7f13f8e / 4

Type: `"bool"`. Optional.

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

<a id="canonical-3bb862e4f8fdf0e99b55732d388a09638408c40393f1c343bd646ca12a195059"></a>

<a id="canonical-e62932174bb67ebce08a26ba174d5355a9ba03123fb84c71ddd4a97155ba242f"></a>

## ip_prefixes property — api_protection_rules.api_endpoint_rules.client_matcher.ip_prefix_list / 88aec7f13f8e / 5

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

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

<a id="canonical-b2e4b665abbf7c45fccd63edd50dbc410ea846ee847634229a948ed4b61fc1c8"></a>

## Next pages — api_protection_rules.api_endpoint_rules.client_matcher.ip_prefix_list / 88aec7f13f8e / 6

- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-b0510836a500354619afdae6bb3e5e3e45631fa4a9bb22642c7bc24bd6a94961)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-1b601e93bd6e3725175502d1c69e6004aa76b3db1b9b9cf4d17c55df58166c8c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e19660af796f534e79c50cbf5b49459378846d2f9e155fa6f9d9663abce4ae68"></a>

## api_protection_rules.api_endpoint_rules.client_matcher.ip_threat_category_list — api_protection_rules.api_endpoint_rules.client_matcher.ip_threat_category_list / c526e5c7f7c4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-7bc735c85b50beaf50f1b45bdbe15c413d15aa714c8844dc469693a74dd86737)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-b0510836a500354619afdae6bb3e5e3e45631fa4a9bb22642c7bc24bd6a94961)
- api_protection_rules.api_endpoint_rules.client_matcher.ip_threat_category_list

<a id="canonical-9206f31968a2e8c033bb6b0dfd3e960d568cb611b7cb2835e44c90e5fb892d07"></a>

Type: `"object"`. single nested block, Optional.

IP Threat Category List Type. List of IP threat categories.

Upstream description:

List of IP threat categories.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_threat_categories")}
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
ip_threat_category_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-d82b98f7477f0c0ebbfef7b26029fcb6ad139dd11fac1730d77f09226643a58f"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.client_matcher.ip_threat_category_list / c526e5c7f7c4 / 3

<a id="canonical-446cc00679f1acb345ee171f6c160e3cafd227485d36eb406a6f978c9dac074d"></a>

<a id="canonical-63a616f582670e4b9efa2cb77b1ce8c8caa78ad03bd5e28e5157de949335ecb3"></a>

## ip_threat_categories property — api_protection_rules.api_endpoint_rules.client_matcher.ip_threat_category_list / c526e5c7f7c4 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`,
\`WEB\_ATTACKS\`, \`BOTNETS\`, \`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`,
\`MOBILE\_THREATS\`, \`TOR\_PROXY\`, \`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to
\`SPAM\_SOURCES\`.

Upstream description:

The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-d6097fb970ae1150417b6a5b582b3730fca7d5628a39eec8130d67ef82860d9d"></a>

## Next pages — api_protection_rules.api_endpoint_rules.client_matcher.ip_threat_category_list / c526e5c7f7c4 / 5

- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-b0510836a500354619afdae6bb3e5e3e45631fa4a9bb22642c7bc24bd6a94961)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-1c4c8497e1e210f68f8a2189739adab9ad669f937d14e41ad4ed5c178342bdf9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4174e117e2c15f63818a4ee91ec37bd3f3b333b88045c27a13d0aa6f164a367a"></a>

## api_protection_rules.api_endpoint_rules.client_matcher.tls_fingerprint_matcher — api_protection_rules.api_endpoint_rules.client_matcher.tls_fingerprint_matcher / 4c319460fe42 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-7bc735c85b50beaf50f1b45bdbe15c413d15aa714c8844dc469693a74dd86737)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-b0510836a500354619afdae6bb3e5e3e45631fa4a9bb22642c7bc24bd6a94961)
- api_protection_rules.api_endpoint_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-c311335a1fcf45bfd4d94fc0bd9424dd3bbfe8356570dcab479c32b9ca9ee684"></a>

Type: `"object"`. single nested block, Optional.

TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are
satisfied..

Upstream description:

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

<a id="canonical-86552ae69d4863383652b7f1fb9062da4fb7121de69b4fda50fab146f2bd60f8"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.client_matcher.tls_fingerprint_matcher / 4c319460fe42 / 3

<a id="canonical-20d3fcec4ac37fdc39f64fc804797518da721dfe49eaf856a183aaa7a682f48a"></a>

<a id="canonical-a358a3ba4c2e95aa3b1925f197d40e9f45c854983868103af04588a165cd4d78"></a>

## classes property — api_protection_rules.api_endpoint_rules.client_matcher.tls_fingerprint_matcher / 4c319460fe42 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Upstream description:

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-def6fab3cdf7164e0f6bf8fc6c8af79e9c4e4b56733e8ed63053f02ea180c6f3"></a>

<a id="canonical-75fe0f5c04c456a5d4180283c1aa195af6c9075588e342d01568d8d470c45c44"></a>

## exact_values property — api_protection_rules.api_endpoint_rules.client_matcher.tls_fingerprint_matcher / 4c319460fe42 / 5

Type: `["list", "string"]`. Optional.

List of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Upstream description:

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

<a id="canonical-aacaddcd47bae7d0da00bd3be7794f9cec6b01527233c98eee37d4b5571078f7"></a>

<a id="canonical-4a7dcf95065cc70f4fee4157dd28e470de6c99cfd271637b734b3cb31c075321"></a>

## excluded_values property — api_protection_rules.api_endpoint_rules.client_matcher.tls_fingerprint_matcher / 4c319460fe42 / 6

Type: `["list", "string"]`. Optional.

List of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be
used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Upstream description:

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

<a id="canonical-237ed45125a561cae6d3ac02572c1d1b6345b3d807c531b5919bf0d895ab72b7"></a>

## Next pages — api_protection_rules.api_endpoint_rules.client_matcher.tls_fingerprint_matcher / 4c319460fe42 / 7

- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-b0510836a500354619afdae6bb3e5e3e45631fa4a9bb22642c7bc24bd6a94961)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-cc0543c0ddd3105157a2af770357ca9c57288d2c391b2792c07e2933bebd06b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-97d683df484b41727de2133b0f8ab169a825ca8144d9f206acf1371e9e54f567"></a>

## api_protection_rules.api_endpoint_rules.metadata — api_protection_rules.api_endpoint_rules.metadata / f10c512902fc / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-7bc735c85b50beaf50f1b45bdbe15c413d15aa714c8844dc469693a74dd86737)
- api_protection_rules.api_endpoint_rules.metadata

<a id="canonical-f3f9c2c2642ef2183e18ea2d1b2a70d7252d19b836c254a6603fe798a480c527"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-0d1cf90d01a9a8a3c55dabeb4df3977a29f3df7f76b75f2d39ccc9e7eb026456"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.metadata / f10c512902fc / 3

<a id="canonical-2567a5247cc9705798a144d9de43d64c7a19ba4bed5f43aaaeaf0dd986f2f418"></a>

<a id="canonical-cdc8a595dd734cdc048fe872fa9d852ebf0f56276459383de57675bf6d3deae3"></a>

## description_spec property — api_protection_rules.api_endpoint_rules.metadata / f10c512902fc / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-45ea6aeea89d8bad4fe20adb7173b9ef1387be2b6a0492617bcc3b04a600501d"></a>

<a id="canonical-14fc7b5b6c971a327b55853982ae7ae9ef116930aadd578d1b70ed1722ec5eed"></a>

## name property — api_protection_rules.api_endpoint_rules.metadata / f10c512902fc / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-72af576068abaac3e54421b46f0078be15a35c2e069ea9fc61599a4b65730a8b"></a>

## Next pages — api_protection_rules.api_endpoint_rules.metadata / f10c512902fc / 6

- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-7bc735c85b50beaf50f1b45bdbe15c413d15aa714c8844dc469693a74dd86737)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-163cdecf7c0181dad7aee3a5cdf10583f2a999d7bf80f9c55daeae1685c3ec06"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2e4e42a1e2ee1757ff4fc29e8cdb842eab3af0816c6f9bf650df77a716a5c45c"></a>

## api_protection_rules.api_endpoint_rules.request_matcher — api_protection_rules.api_endpoint_rules.request_matcher / 03576e017400 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-7bc735c85b50beaf50f1b45bdbe15c413d15aa714c8844dc469693a74dd86737)
- api_protection_rules.api_endpoint_rules.request_matcher

<a id="canonical-af03fcf031630d2a1e854a85a226d13bd9ed69e916b4df59d6f3da7906623580"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for request matcher.

Upstream description:

Request conditions for matching a rule.

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
request_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-d32a73d3397eb4e3fb7026fefcf1a7393223a8900d890b6a67bd573130516132"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher / 03576e017400 / 3

- [cookie_matchers](resources--http_loadbalancer--reference--group-005.md#canonical-c9159b2d7d68ff6637dd56907182f3b1290c57ee75e723e4be6049918162bc3e): complete subsection reference.

- [headers](resources--http_loadbalancer--reference--group-005.md#canonical-eeadd538b6db5f14e51b49a09ff32cf48ca57172b0491611882adfae77b756a6): complete subsection reference.

- [jwt_claims](resources--http_loadbalancer--reference--group-005.md#canonical-aa520488c883a3f4c35636bf5ece99613c3504d7d11466b1f29ab3366c9fcc4f): complete subsection reference.

- [query_params](resources--http_loadbalancer--reference--group-005.md#canonical-f6b1929542d549af151af82574d8b4b6575576f372854bb534f956cee94fa63e): complete subsection reference.

<a id="canonical-a0ae3967f2e38be1b715c66e57b07d785a59236e4856af810d3c6d57607e7507"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher / 03576e017400 / 4

- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-005.md#canonical-c9159b2d7d68ff6637dd56907182f3b1290c57ee75e723e4be6049918162bc3e)
- [api_protection_rules.api_endpoint_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-005.md#canonical-eeadd538b6db5f14e51b49a09ff32cf48ca57172b0491611882adfae77b756a6)
- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-005.md#canonical-aa520488c883a3f4c35636bf5ece99613c3504d7d11466b1f29ab3366c9fcc4f)
- [api_protection_rules.api_endpoint_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-005.md#canonical-f6b1929542d549af151af82574d8b4b6575576f372854bb534f956cee94fa63e)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-7bc735c85b50beaf50f1b45bdbe15c413d15aa714c8844dc469693a74dd86737)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c9159b2d7d68ff6637dd56907182f3b1290c57ee75e723e4be6049918162bc3e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-37d3e3d888d7f889fbb626593fdce1249941e1234474ff06e13c4df17ac1c9b9"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers / 9050dd5be26f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-7bc735c85b50beaf50f1b45bdbe15c413d15aa714c8844dc469693a74dd86737)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-163cdecf7c0181dad7aee3a5cdf10583f2a999d7bf80f9c55daeae1685c3ec06)
- api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers

<a id="canonical-885c4efca241ac921da71fe78f94bacb465b90d038b9b2a57bdf4992a5c037bf"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all cookies that need to be matched. The criteria for matching each cookie is
described in individual instances of CookieMatcherType. The actual cookie values are extracted from
the request API as a list of strings for each cookie name.

Upstream description:

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

Terraform syntax:

```terraform
cookie_matchers {
  # Configure direct properties listed below.
}
```

<a id="canonical-0b583369020c2f5ed9fa7871f94c4fa4bfc62d11184fd125bda94ee4960ea3f1"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers / 9050dd5be26f / 3

- [check_not_present](resources--http_loadbalancer--reference--group-005.md#canonical-e4ba5ef50b95226775a855d5aa7ae90a662c71266bbd6ecfee74b21f70dc4d53): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-005.md#canonical-e8ad91a74d00f0cc48f44c332b33554bdb75e8e4c64bcafc138d120a8b87dc8c): complete subsection reference.

<a id="canonical-0530967851721b57dce1247bcd82fcaae4cb4729707b88f382026f0a1c448ce5"></a>

<a id="canonical-f5268aa407cc6c225eedad987e6dfc8ae4b96488f66c77230752edf94cee9b71"></a>

## invert_matcher property — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers / 9050dd5be26f / 4

Type: `"bool"`. Optional.

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

- [item](resources--http_loadbalancer--reference--group-005.md#canonical-427ca64f2169c7f25b4d08f44c2b30a25bf306a7b85568ac97bddb1df11791ef): complete subsection reference.

<a id="canonical-378d0e057a93f731b4a3d5352c7a387b3e35b68f5abba0265426dab879679f5c"></a>

<a id="canonical-e2b9974d4abe7f642068ecc857959ac09a1e839eff809926839200e3b15a0082"></a>

## name property — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers / 9050dd5be26f / 5

Type: `"string"`. Optional.

Cookie Name. A case-sensitive cookie name.

Upstream description:

A case-sensitive cookie name.

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

<a id="canonical-0c134f9c48ed1ca94a65d303b2ba572af1c0bb61522d901e8e07aeb691957b4c"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers / 9050dd5be26f / 6

- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present](resources--http_loadbalancer--reference--group-005.md#canonical-e4ba5ef50b95226775a855d5aa7ae90a662c71266bbd6ecfee74b21f70dc4d53)
- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_present](resources--http_loadbalancer--reference--group-005.md#canonical-e8ad91a74d00f0cc48f44c332b33554bdb75e8e4c64bcafc138d120a8b87dc8c)
- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item](resources--http_loadbalancer--reference--group-005.md#canonical-427ca64f2169c7f25b4d08f44c2b30a25bf306a7b85568ac97bddb1df11791ef)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-163cdecf7c0181dad7aee3a5cdf10583f2a999d7bf80f9c55daeae1685c3ec06)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e4ba5ef50b95226775a855d5aa7ae90a662c71266bbd6ecfee74b21f70dc4d53"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4d7f4d88536c50f5266ab7cec06894fc60720b761db3ff8def0490be702b2ea3"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_no / 80cba010e33b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-7bc735c85b50beaf50f1b45bdbe15c413d15aa714c8844dc469693a74dd86737)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-163cdecf7c0181dad7aee3a5cdf10583f2a999d7bf80f9c55daeae1685c3ec06)
- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-005.md#canonical-c9159b2d7d68ff6637dd56907182f3b1290c57ee75e723e4be6049918162bc3e)
- api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-67bf16e1b53eb77542c3c5268211f1d0f75fa0630f45650204ef7fbc0216ef83"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_not_present = {}
```

<a id="canonical-8956fdbd193ad75beea7a14d962d6ceb4d33a5cc51775108f07fb6e096257d88"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_no / 80cba010e33b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6b60a3717b2968a5ec2f411d9c47842201044ab70b3e6b6278e5a8e9c1becf7c"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_no / 80cba010e33b / 4

- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-005.md#canonical-c9159b2d7d68ff6637dd56907182f3b1290c57ee75e723e4be6049918162bc3e)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e8ad91a74d00f0cc48f44c332b33554bdb75e8e4c64bcafc138d120a8b87dc8c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2feb24a3a7ef6205b46d9ba342116b4a0439a0f54f651ec9e1fea7e072e7cacc"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_present — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_pr / 1997455758a0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-7bc735c85b50beaf50f1b45bdbe15c413d15aa714c8844dc469693a74dd86737)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-163cdecf7c0181dad7aee3a5cdf10583f2a999d7bf80f9c55daeae1685c3ec06)
- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-005.md#canonical-c9159b2d7d68ff6637dd56907182f3b1290c57ee75e723e4be6049918162bc3e)
- api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-d84b1ecdade22376b07ba3ba39085e62f8bb0ef3cacb5909315fbf07dd343db7"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_present = {}
```

<a id="canonical-70b19273170f9b00ebe52f2831e9e62afdfc368156426dc7f3c7d5171e39c572"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_pr / 1997455758a0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5105cc80ed77dd0d8b5f305f35ebb0f8cd383f8dd124d0bc2637e73a7f6a5c2b"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_pr / 1997455758a0 / 4

- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-005.md#canonical-c9159b2d7d68ff6637dd56907182f3b1290c57ee75e723e4be6049918162bc3e)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-427ca64f2169c7f25b4d08f44c2b30a25bf306a7b85568ac97bddb1df11791ef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a958716d090a291eb0c14d32772bdd09aaa3be4c0994b31fe3e6515657f11667"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item / fee01615e6d1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-7bc735c85b50beaf50f1b45bdbe15c413d15aa714c8844dc469693a74dd86737)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-163cdecf7c0181dad7aee3a5cdf10583f2a999d7bf80f9c55daeae1685c3ec06)
- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-005.md#canonical-c9159b2d7d68ff6637dd56907182f3b1290c57ee75e723e4be6049918162bc3e)
- api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item

<a id="canonical-8fa52e244d9ab0913859dd865be96d8505451998a8c5188422985a336ceca8cc"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-c769053a0d50c45802abdbd383555cf8b22b04e83938608edc69311088cc78ee"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item / fee01615e6d1 / 3

<a id="canonical-f18007b10568b9b8be0f92aba12a65afd447df70f0cad1212ba8a879a90229d6"></a>

<a id="canonical-bb6ba6af917de531cca4fc1d87c27dbc3a02d07980ab902d5e2317174f060ef6"></a>

## exact_values property — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item / fee01615e6d1 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

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

<a id="canonical-2f9e6f3685248ef52a4f624a1d16aec1642d7f500b68fde88a2e4adf6198d7ef"></a>

<a id="canonical-02fd69e392d7b80e6220291b32d8d3f87bb30f8b950f707f39e54ec03be927e8"></a>

## regex_values property — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item / fee01615e6d1 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

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

<a id="canonical-6b314e4e70df0eb7ef41e15d2f1dfad9e485d066a2ee909321999dadf06197dc"></a>

<a id="canonical-207139f947fd2db3901f7a8fe61e3a03c00941a953f65bc4272d2d0d879a46f5"></a>

## transformers property — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item / fee01615e6d1 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

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

<a id="canonical-84f25128ede1398b423037eb68ada11f120ee2bee754dbb146899a483d6cbbd5"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item / fee01615e6d1 / 7

- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-005.md#canonical-c9159b2d7d68ff6637dd56907182f3b1290c57ee75e723e4be6049918162bc3e)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-eeadd538b6db5f14e51b49a09ff32cf48ca57172b0491611882adfae77b756a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-936af892d2629e04ad7d01bf52a4f48390897a756ef791eedc33f54adfe9e4fe"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.headers — api_protection_rules.api_endpoint_rules.request_matcher.headers / 62645357ff12 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-7bc735c85b50beaf50f1b45bdbe15c413d15aa714c8844dc469693a74dd86737)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-163cdecf7c0181dad7aee3a5cdf10583f2a999d7bf80f9c55daeae1685c3ec06)
- api_protection_rules.api_endpoint_rules.request_matcher.headers

<a id="canonical-2af321bc693b022d03853aed5bffcf0361988ffe9a6230f31bf6312bc977d117"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for various HTTP headers that need to match. The criteria for matching each HTTP
header are described in individual HeaderMatcherType instances. The actual HTTP header values are
extracted from the request API as a list of strings for each HTTP header type.

Upstream description:

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

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-a1508c6d90d7db73c1bde4d9f57eeb0230ee75ef980c9dacd82750817c37a759"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.headers / 62645357ff12 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-005.md#canonical-5d99a08df14d7df1db5c783de48d0b3fd1fafaeb0203a508b3725cf8d76da8a3): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-005.md#canonical-336ca36c24f788b7c563f564b575cc8d7af5d1db325225b1b5f499769203fb57): complete subsection reference.

<a id="canonical-1caa22c7659a50de228d15b7f4a462c8fcbff9cdf4354ca6728787531fa14b64"></a>

<a id="canonical-8b546ca66c25ba1dec0893941d7513fa3d8fd3394fb78cbcd6342a74143da644"></a>

## invert_matcher property — api_protection_rules.api_endpoint_rules.request_matcher.headers / 62645357ff12 / 4

Type: `"bool"`. Optional.

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

- [item](resources--http_loadbalancer--reference--group-005.md#canonical-2a8bfb5fdcbd069e81c1708d6a02c3590fc08d125285f4a881cbeb4d61306a14): complete subsection reference.

<a id="canonical-6e4a1cbd05193a13c85b0df51ef892d26590045ae0ed10a8266bf366d7d1bfc8"></a>

<a id="canonical-79367088c3f1d1c7fb1a579e7b4deba3c07b19ce63f780abcf2c323783ce4cec"></a>

## name property — api_protection_rules.api_endpoint_rules.request_matcher.headers / 62645357ff12 / 5

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

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

<a id="canonical-d2f5a8d6285cc44eec9b2e218c825a932ed92dac795997796c3e4650664fa2d4"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.headers / 62645357ff12 / 6

- [api_protection_rules.api_endpoint_rules.request_matcher.headers.check_not_present](resources--http_loadbalancer--reference--group-005.md#canonical-5d99a08df14d7df1db5c783de48d0b3fd1fafaeb0203a508b3725cf8d76da8a3)
- [api_protection_rules.api_endpoint_rules.request_matcher.headers.check_present](resources--http_loadbalancer--reference--group-005.md#canonical-336ca36c24f788b7c563f564b575cc8d7af5d1db325225b1b5f499769203fb57)
- [api_protection_rules.api_endpoint_rules.request_matcher.headers.item](resources--http_loadbalancer--reference--group-005.md#canonical-2a8bfb5fdcbd069e81c1708d6a02c3590fc08d125285f4a881cbeb4d61306a14)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-163cdecf7c0181dad7aee3a5cdf10583f2a999d7bf80f9c55daeae1685c3ec06)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-5d99a08df14d7df1db5c783de48d0b3fd1fafaeb0203a508b3725cf8d76da8a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2f1cec13540d4fef77a6104a10eec8e1d0f9f8969f9d030f7677621b42a77745"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.headers.check_not_present — api_protection_rules.api_endpoint_rules.request_matcher.headers.check_not_presen / f1b55cea8595 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-7bc735c85b50beaf50f1b45bdbe15c413d15aa714c8844dc469693a74dd86737)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-163cdecf7c0181dad7aee3a5cdf10583f2a999d7bf80f9c55daeae1685c3ec06)
- [api_protection_rules.api_endpoint_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-005.md#canonical-eeadd538b6db5f14e51b49a09ff32cf48ca57172b0491611882adfae77b756a6)
- api_protection_rules.api_endpoint_rules.request_matcher.headers.check_not_present

<a id="canonical-34dea2dffa6ac31dbcd9858ac974ccb6b1279f8e098ecd02071f5ec018b24d0e"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_not_present = {}
```

<a id="canonical-9f52946e3c7c22dedb9b3c95eb17a73d3b304fca6da0b3968ef5c95ed028aba8"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.headers.check_not_presen / f1b55cea8595 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-668fd3e363da4ea3effec37b6eb97faec525e3e63f3d21176173cdc6d0a0e9b8"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.headers.check_not_presen / f1b55cea8595 / 4

- [api_protection_rules.api_endpoint_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-005.md#canonical-eeadd538b6db5f14e51b49a09ff32cf48ca57172b0491611882adfae77b756a6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-336ca36c24f788b7c563f564b575cc8d7af5d1db325225b1b5f499769203fb57"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be17f9ef71aa8d0e65f5fb0bd6c849b7cbfa7ca38fed5d994a2b65773fe27ea4"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.headers.check_present — api_protection_rules.api_endpoint_rules.request_matcher.headers.check_present / 1014b84f7422 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-7bc735c85b50beaf50f1b45bdbe15c413d15aa714c8844dc469693a74dd86737)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-163cdecf7c0181dad7aee3a5cdf10583f2a999d7bf80f9c55daeae1685c3ec06)
- [api_protection_rules.api_endpoint_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-005.md#canonical-eeadd538b6db5f14e51b49a09ff32cf48ca57172b0491611882adfae77b756a6)
- api_protection_rules.api_endpoint_rules.request_matcher.headers.check_present

<a id="canonical-ea5dcdd7f92561ea50650813832840bc38634f74cb47067bb7645c8672ae1d10"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_present = {}
```

<a id="canonical-4b101885b41f1f8b2664a1659deb47035e8e0bc91c41374c8eda4bcace45f0d3"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.headers.check_present / 1014b84f7422 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-edd1800b6b4f8e3318c5d7019a518609b7a8174b68b4e63a1d89ce20b26a181b"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.headers.check_present / 1014b84f7422 / 4

- [api_protection_rules.api_endpoint_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-005.md#canonical-eeadd538b6db5f14e51b49a09ff32cf48ca57172b0491611882adfae77b756a6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-2a8bfb5fdcbd069e81c1708d6a02c3590fc08d125285f4a881cbeb4d61306a14"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee79ef754b0a610e09a67aab7a5a89f2e5e1b29057c7df36cc6f5fa6528c98c5"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.headers.item — api_protection_rules.api_endpoint_rules.request_matcher.headers.item / 33e0b7a8fc89 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-7bc735c85b50beaf50f1b45bdbe15c413d15aa714c8844dc469693a74dd86737)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-163cdecf7c0181dad7aee3a5cdf10583f2a999d7bf80f9c55daeae1685c3ec06)
- [api_protection_rules.api_endpoint_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-005.md#canonical-eeadd538b6db5f14e51b49a09ff32cf48ca57172b0491611882adfae77b756a6)
- api_protection_rules.api_endpoint_rules.request_matcher.headers.item

<a id="canonical-37d0c7a02172055b755cb61347ddbe8fdb89eaa2891a942e2ca15712072c221e"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-0fb60452c05fde7a655f2b2b610d0698af2f8267af68903c5ebe7c58c70a5470"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.headers.item / 33e0b7a8fc89 / 3

<a id="canonical-2aadbb6f46ac4603eb930b5f2548d26714af411ef2ce613651035d61d4af8c1f"></a>

<a id="canonical-3d2cdcbedc6873eab824515959bde18c4d7fbd2b247cd5004282c5eab846d2ab"></a>

## exact_values property — api_protection_rules.api_endpoint_rules.request_matcher.headers.item / 33e0b7a8fc89 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

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

<a id="canonical-8c15fb5698f7bd2aa1bf2296aae8912b9ae4486f8c1f94827dc52f7225cc40c4"></a>

<a id="canonical-b29d5d013c1869e0009da7a36b25a9b1a6d2fda66c8b4396410c2ddd5d594904"></a>

## regex_values property — api_protection_rules.api_endpoint_rules.request_matcher.headers.item / 33e0b7a8fc89 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

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

<a id="canonical-5a041b134483b9302671be8e73b37e4ab6ccf891cef1f10f0f8892bc17dc2233"></a>

<a id="canonical-9a6119561c5e4eec664ece283786430145bcc2b6046ada816c37abc923db7896"></a>

## transformers property — api_protection_rules.api_endpoint_rules.request_matcher.headers.item / 33e0b7a8fc89 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

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

<a id="canonical-b1d38e70752415087510fd5a0bc80c0a6a0541deaac7ccf5d7d3a6e11e89960b"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.headers.item / 33e0b7a8fc89 / 7

- [api_protection_rules.api_endpoint_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-005.md#canonical-eeadd538b6db5f14e51b49a09ff32cf48ca57172b0491611882adfae77b756a6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-aa520488c883a3f4c35636bf5ece99613c3504d7d11466b1f29ab3366c9fcc4f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9753acb2804a2a8a5e2ce3700049055a165974ea647e4ddd021979881f985272"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims / 071733a76cad / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-7bc735c85b50beaf50f1b45bdbe15c413d15aa714c8844dc469693a74dd86737)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-163cdecf7c0181dad7aee3a5cdf10583f2a999d7bf80f9c55daeae1685c3ec06)
- api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims

<a id="canonical-0942203bc7485a4aad920b12025cc083b254fc18fa3859d3f0a55952a1ddefcb"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings.

Upstream description:

A list of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates
must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled.

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

Terraform syntax:

```terraform
jwt_claims {
  # Configure direct properties listed below.
}
```

<a id="canonical-2da580d5fdb63a98c24cf07bedb1d77f0992520538b81bd874bcdc69106f310a"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims / 071733a76cad / 3

- [check_not_present](resources--http_loadbalancer--reference--group-005.md#canonical-0316124327109ea9d8b71902c9daf6a770ed76a0c45bc76abef569dfbc2c09ca): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-005.md#canonical-47b65cdb683028be74c3ca78375f0845a9f53fe7508559e136fee8b517fe3131): complete subsection reference.

<a id="canonical-277474a8cbd321f2e2c39e623ed137cff95371df14c6bdcb106b35e06216533b"></a>

<a id="canonical-44221156343ec91875d261690a5c42677ab5537cca7fe933a3da449c626122bd"></a>

## invert_matcher property — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims / 071733a76cad / 4

Type: `"bool"`. Optional.

Invert Matcher. Invert the match result.

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

- [item](resources--http_loadbalancer--reference--group-005.md#canonical-23eb1c52d2c6f47042dd0d0f994b154001e6707ea9990a9b9f4361f830e598f0): complete subsection reference.

<a id="canonical-62d14e1e043d6181739950d143ae977e5fab0973c36e3c055b66cb73e36d1f03"></a>

<a id="canonical-5bd6bf78d2407123964ee60ce6cbbf5bcf4e6e572d977f4dffa32543b06b6e33"></a>

## name property — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims / 071733a76cad / 5

Type: `"string"`. Optional.

JWT Claim Name. JWT claim name.

Upstream description:

JWT claim name.

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

<a id="canonical-23a7859ff16385d53b2d8b80ce99ceb736f67f7a754c54614b3f5a58321cfff9"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims / 071733a76cad / 6

- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_not_present](resources--http_loadbalancer--reference--group-005.md#canonical-0316124327109ea9d8b71902c9daf6a770ed76a0c45bc76abef569dfbc2c09ca)
- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_present](resources--http_loadbalancer--reference--group-005.md#canonical-47b65cdb683028be74c3ca78375f0845a9f53fe7508559e136fee8b517fe3131)
- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item](resources--http_loadbalancer--reference--group-005.md#canonical-23eb1c52d2c6f47042dd0d0f994b154001e6707ea9990a9b9f4361f830e598f0)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-163cdecf7c0181dad7aee3a5cdf10583f2a999d7bf80f9c55daeae1685c3ec06)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-0316124327109ea9d8b71902c9daf6a770ed76a0c45bc76abef569dfbc2c09ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-db2b227a43d490cb2667c1bee116887679305cb4d8a8e09015c0bed8760a2858"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_not_present — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_not_pre / 600912dbaa18 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-7bc735c85b50beaf50f1b45bdbe15c413d15aa714c8844dc469693a74dd86737)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-163cdecf7c0181dad7aee3a5cdf10583f2a999d7bf80f9c55daeae1685c3ec06)
- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-005.md#canonical-aa520488c883a3f4c35636bf5ece99613c3504d7d11466b1f29ab3366c9fcc4f)
- api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-e63b1a6020410b398260eb7fd06ad81908ed5660ccaa3a064e2955f5604ee6c8"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_not_present = {}
```

<a id="canonical-d9fff47a1e44d818a36dcbec935b6dfcc296e49712334c4797e6fee572ff59f0"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_not_pre / 600912dbaa18 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-16e19674b95c0792cffaabbd79cbb7db2c0451aeb87dad0803dd30ab4d901281"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_not_pre / 600912dbaa18 / 4

- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-005.md#canonical-aa520488c883a3f4c35636bf5ece99613c3504d7d11466b1f29ab3366c9fcc4f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-47b65cdb683028be74c3ca78375f0845a9f53fe7508559e136fee8b517fe3131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b12c7a787251af79c9401a501fa774073f1db615220546afa7efb6ad8937a64b"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_present — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_present / 3ce277090834 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-7bc735c85b50beaf50f1b45bdbe15c413d15aa714c8844dc469693a74dd86737)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-163cdecf7c0181dad7aee3a5cdf10583f2a999d7bf80f9c55daeae1685c3ec06)
- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-005.md#canonical-aa520488c883a3f4c35636bf5ece99613c3504d7d11466b1f29ab3366c9fcc4f)
- api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_present

<a id="canonical-b204c1363ba500677d5eab60d122368bc4f87417b5f737a412cd3d3687e79095"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_present = {}
```

<a id="canonical-1b85011efe097a2eac577495b51d567f2d3c9c0f29ed5b45019f89f73801aabb"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_present / 3ce277090834 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e9068a972923cd019164ed37e13713396129e013261c89ee7db430dd1201b40b"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_present / 3ce277090834 / 4

- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-005.md#canonical-aa520488c883a3f4c35636bf5ece99613c3504d7d11466b1f29ab3366c9fcc4f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-23eb1c52d2c6f47042dd0d0f994b154001e6707ea9990a9b9f4361f830e598f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f4b62d5f4b7b0c92830fe0df9ccad23a74f1ab9b8b6dde2c4aebd43f98a497a0"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item / 132665847e05 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-7bc735c85b50beaf50f1b45bdbe15c413d15aa714c8844dc469693a74dd86737)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-163cdecf7c0181dad7aee3a5cdf10583f2a999d7bf80f9c55daeae1685c3ec06)
- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-005.md#canonical-aa520488c883a3f4c35636bf5ece99613c3504d7d11466b1f29ab3366c9fcc4f)
- api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item

<a id="canonical-d6b7194e0f66d5a3184377a997c334e856b3346274c2652200e004a065e2ba0d"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-83e6516c6fea3dbdce5f36b16a859b1e06a6c878e6fec3a8b2ab2254341569a6"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item / 132665847e05 / 3

<a id="canonical-f2d6313bab1cbd70cdff4d0c4d91adb6345b1e3f68ad857bf926e62a622785c1"></a>

<a id="canonical-2a53527feb33672df425019336cfc3cf37d8826120cf6a62260899c7d13f9763"></a>

## exact_values property — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item / 132665847e05 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

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

<a id="canonical-357fcadcf20378ada45d7ae7f9077ef007ca517f67528b8fe95760445ed27bc8"></a>

<a id="canonical-cec27b912e43da4aa9c898ce060cc6757db9bcc7884ce8b9fd0a70065495aa19"></a>

## regex_values property — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item / 132665847e05 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

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

<a id="canonical-b1d66c2ab49b6a6f6838580ca3adef44ff9e47832e45dac4333f11202627a865"></a>

<a id="canonical-79dfeefae888c6aa2b8f366612ef0fe51232abdd53daa213d84e736c96d15123"></a>

## transformers property — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item / 132665847e05 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

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

<a id="canonical-b23a57160594378a4e418355ee280c9170d9c1df750c1a1213441a35aee675a9"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item / 132665847e05 / 7

- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-005.md#canonical-aa520488c883a3f4c35636bf5ece99613c3504d7d11466b1f29ab3366c9fcc4f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-f6b1929542d549af151af82574d8b4b6575576f372854bb534f956cee94fa63e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c660e33b249c88681b33a338f0cb5e6a33bc2fc8ea853ea4b380a0dc928744d7"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.query_params — api_protection_rules.api_endpoint_rules.request_matcher.query_params / 3b5dd3a3bf18 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-7bc735c85b50beaf50f1b45bdbe15c413d15aa714c8844dc469693a74dd86737)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-163cdecf7c0181dad7aee3a5cdf10583f2a999d7bf80f9c55daeae1685c3ec06)
- api_protection_rules.api_endpoint_rules.request_matcher.query_params

<a id="canonical-4b781f9a17b2704e02acb0b90e79bf2fca461788e9d2571b40e0f2e42a18179a"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all query parameters that need to be matched. The criteria for matching each
query parameter are described in individual instances of QueryParameterMatcherType. The actual query
parameter values are extracted from the request API as a list of strings for each query..

Upstream description:

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

Terraform syntax:

```terraform
query_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-53d99d9a00592bfab7ffcf9edf7729cf68e961382d255ab30691f4d9f1dc6e2f"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.query_params / 3b5dd3a3bf18 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-005.md#canonical-c5a1aaa8716b893da7c88d4d657c5fd8998c3642ca092c7a1245117134aadc36): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-005.md#canonical-8b64eff92fafec67fceb8b3194fc837efa6ab899dfec8d74f77583267d316f46): complete subsection reference.

<a id="canonical-89908c6811a7766e2283544dbe61f43759c498972ba0f286d6f90e87b2023947"></a>

<a id="canonical-5b498a3cef0f9fe046e51706003f5e213f9471593e18ff9931440d149ed19d9d"></a>

## invert_matcher property — api_protection_rules.api_endpoint_rules.request_matcher.query_params / 3b5dd3a3bf18 / 4

Type: `"bool"`. Optional.

Invert Query Parameter Matcher. Invert the match result.

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

- [item](resources--http_loadbalancer--reference--group-005.md#canonical-721cb5d1751a1b51e0ba070f0fda856fd6858c3a0466c24c02c37b2ea3d93a06): complete subsection reference.

<a id="canonical-681829cd892b83cbeb8e97a2b6db5f3635b0ef46f11de33d9854063dabeb1b1b"></a>

<a id="canonical-b4dd8d26e24a4d95a8978f02abe7e7fa1f23ecd66b231dba2e0d89a1d8969039"></a>

## key property — api_protection_rules.api_endpoint_rules.request_matcher.query_params / 3b5dd3a3bf18 / 5

Type: `"string"`. Optional.

Case-sensitive HTTP query parameter name.

Upstream description:

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-f9263f9caec23a449556f2815f835ee105c22985a50cf0e2e1f84c7b78e9e28e"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.query_params / 3b5dd3a3bf18 / 6

- [api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_not_present](resources--http_loadbalancer--reference--group-005.md#canonical-c5a1aaa8716b893da7c88d4d657c5fd8998c3642ca092c7a1245117134aadc36)
- [api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_present](resources--http_loadbalancer--reference--group-005.md#canonical-8b64eff92fafec67fceb8b3194fc837efa6ab899dfec8d74f77583267d316f46)
- [api_protection_rules.api_endpoint_rules.request_matcher.query_params.item](resources--http_loadbalancer--reference--group-005.md#canonical-721cb5d1751a1b51e0ba070f0fda856fd6858c3a0466c24c02c37b2ea3d93a06)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-163cdecf7c0181dad7aee3a5cdf10583f2a999d7bf80f9c55daeae1685c3ec06)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c5a1aaa8716b893da7c88d4d657c5fd8998c3642ca092c7a1245117134aadc36"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7628b166baee2fcc9b9cd76b9466cce8b7720709f097c4c95a4fedbf683af209"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_not_present — api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_not_p / 2b84ca8c7b29 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-7bc735c85b50beaf50f1b45bdbe15c413d15aa714c8844dc469693a74dd86737)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-163cdecf7c0181dad7aee3a5cdf10583f2a999d7bf80f9c55daeae1685c3ec06)
- [api_protection_rules.api_endpoint_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-005.md#canonical-f6b1929542d549af151af82574d8b4b6575576f372854bb534f956cee94fa63e)
- api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_not_present

<a id="canonical-7a8e0b50bacd898a455d5fa2787e3791acffbb6da319b36dfa228ac85d1f469b"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_not_present = {}
```

<a id="canonical-bcae7c401105fc9cf3f8e8832e8fbeee5d1e1c166b0bdd74149c668460f56858"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_not_p / 2b84ca8c7b29 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-020c981f694b8636a4ba634565a7833e6617f3d8502b34b9c47c3af5f108c8d1"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_not_p / 2b84ca8c7b29 / 4

- [api_protection_rules.api_endpoint_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-005.md#canonical-f6b1929542d549af151af82574d8b4b6575576f372854bb534f956cee94fa63e)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-8b64eff92fafec67fceb8b3194fc837efa6ab899dfec8d74f77583267d316f46"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2f08f9995f2151d9538af697d43f5b43d2399a89c86aeb551fc0e317f7e23e36"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_present — api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_prese / d67ff6479905 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-7bc735c85b50beaf50f1b45bdbe15c413d15aa714c8844dc469693a74dd86737)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-163cdecf7c0181dad7aee3a5cdf10583f2a999d7bf80f9c55daeae1685c3ec06)
- [api_protection_rules.api_endpoint_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-005.md#canonical-f6b1929542d549af151af82574d8b4b6575576f372854bb534f956cee94fa63e)
- api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_present

<a id="canonical-c3bfffa17344d779d6e04aae66dacd6ba953d1e82afb5737cfddc9915d4df6b1"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_present = {}
```

<a id="canonical-61e54d521e510e67ee8d5e8488a021275233245366eec0cee71d07f6ac9681ad"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_prese / d67ff6479905 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cfd43e85dfdcacd27a524c0cadd1c8f1e222076cbbc0a8b9c8f362f828b7873b"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_prese / d67ff6479905 / 4

- [api_protection_rules.api_endpoint_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-005.md#canonical-f6b1929542d549af151af82574d8b4b6575576f372854bb534f956cee94fa63e)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-721cb5d1751a1b51e0ba070f0fda856fd6858c3a0466c24c02c37b2ea3d93a06"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-718585ce5d8296faf34f95261cc303835c839e89eeb1de64ca0800a267299fbb"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.query_params.item — api_protection_rules.api_endpoint_rules.request_matcher.query_params.item / 729ddedec5e5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-7bc735c85b50beaf50f1b45bdbe15c413d15aa714c8844dc469693a74dd86737)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-163cdecf7c0181dad7aee3a5cdf10583f2a999d7bf80f9c55daeae1685c3ec06)
- [api_protection_rules.api_endpoint_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-005.md#canonical-f6b1929542d549af151af82574d8b4b6575576f372854bb534f956cee94fa63e)
- api_protection_rules.api_endpoint_rules.request_matcher.query_params.item

<a id="canonical-25aa2ef44c2d26dc0e0d130c2a2213539f568fe58b06bfc1804b48bd715cfaa0"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-b1a6421fd5442048522ef58dea98c88bc4beda42bb6899e766728e431b34d295"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.query_params.item / 729ddedec5e5 / 3

<a id="canonical-e50957a28ce660a15dc53b6cea5358f1465e42a00bf80ee00f806308e7d04a1d"></a>

<a id="canonical-0050ecd2763e1c63abcd79cf166dc6f8fd42c342e7b99befc1d8cad5e4e7b5cc"></a>

## exact_values property — api_protection_rules.api_endpoint_rules.request_matcher.query_params.item / 729ddedec5e5 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

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

<a id="canonical-51dd455f73de392da1bdb129960be36ef84eefeaa636e9289d963667e505fe8b"></a>

<a id="canonical-8df6698cca78fe9aa5dede946cc744a26b58772eebc4e5ad0cb0136ada0bc838"></a>

## regex_values property — api_protection_rules.api_endpoint_rules.request_matcher.query_params.item / 729ddedec5e5 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

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

<a id="canonical-2903fb4b9a46f552f6d1fde15548d43401f415602e58e6cab1ff86463e38b7e4"></a>

<a id="canonical-f9d96ee0d39e6d68437ebb2f0a16e3df1dea9cde0e8572f93a04b5de431207b4"></a>

## transformers property — api_protection_rules.api_endpoint_rules.request_matcher.query_params.item / 729ddedec5e5 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

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

<a id="canonical-373fb9a100d3cf9689bd8de6f2161fe87f31b0caa309748905b8eabb0b6e30db"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.query_params.item / 729ddedec5e5 / 7

- [api_protection_rules.api_endpoint_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-005.md#canonical-f6b1929542d549af151af82574d8b4b6575576f372854bb534f956cee94fa63e)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eb23a5cd8f93baaae45d0aa6d4de94190d981f184d653831acbba55eb341a1f8"></a>

## api_protection_rules.api_groups_rules — api_protection_rules.api_groups_rules / a15f4f141ac4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- api_protection_rules.api_groups_rules

<a id="canonical-316084d9a35a9783d576a37693f257ff97189c97af2809cee6dce6a21744af17"></a>

Type: `"object"`. list nested block, Optional.

Category includes rules per API group or Server URL. For API groups, refer to API Definition which
includes API groups derived from uploaded swaggers.

Upstream description:

This category includes rules per API group or Server URL. For API groups, refer to API Definition
which includes API groups derived from uploaded swaggers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("base_path"),
  validators.ConflictingListObjectAttributes("any_domain",
    "specific_domain")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 20,
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
    "ves.io.schema.rules.repeated.max_items": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "20"
  }
}
```

Terraform syntax:

```terraform
api_groups_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-fdc4f8bea1affc68800bf397953844680c69cd8fb706989a11edcdf7ea2e0ef0"></a>

## Direct properties — api_protection_rules.api_groups_rules / a15f4f141ac4 / 3

- [action](resources--http_loadbalancer--reference--group-005.md#canonical-1de69af0465eaf9c52019846882d556ab548a1c547677838bd90200c082b7575): complete subsection reference.

- [any_domain](resources--http_loadbalancer--reference--group-005.md#canonical-e72f567b5726a2e75d7ab309fe0102ae78f73f6e5729ff463221cbe0381669d8): complete subsection reference.

<a id="canonical-05b166ba7c5e6499377f963de54bca56640681ec580b8c934d0ca0a30c60fc16"></a>

<a id="canonical-71a33f7ee19afcefb47888aa829f7ec5b90e5ff4d0fffc3c7e3bb3a2b211138c"></a>

## api_group property — api_protection_rules.api_groups_rules / a15f4f141ac4 / 4

Type: `"string"`. Optional.

API groups derived from API Definition swaggers. For example oas-all-operations including all paths
and methods from the swaggers, oas-base-URLs covering all requests under base-paths from the
swaggers. Custom groups can be created if user tags paths or operations with 'x-F5 Distributed..

Upstream description:

API groups derived from API Definition swaggers. For example oas-all-operations including all paths
and methods from the swaggers, oas-base-URLs covering all requests under base-paths from the
swaggers. Custom groups can be created if user tags paths or operations with "x-F5 Distributed
Cloud-API-group" extensions inside swaggers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-672e67e0caee0060fb07341c25b80cce7046f236e0bd9dc795e332a595001b14"></a>

<a id="canonical-f6089e78713a6b070f08304051b8d403b730540f4e5783f9ef1e1a6c0988d007"></a>

## base_path property — api_protection_rules.api_groups_rules / a15f4f141ac4 / 5

Type: `"string"`. Optional.

Base Path. Prefix of the request path. For example: /v1.

Upstream description:

Prefix of the request path. For example: /v1.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-279540d1b6e4ec14980e15dedc82deacdfc801d2a7d6aa5fb34f6d87530d3849): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-005.md#canonical-da4775e72c6bbdc8a73bf5011a0d454b88dfe48c53393ab165beda163a01482f): complete subsection reference.

- [request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2829dee6801699124f354f49d1b1d6b69abd8bdd76c05d8cef9252c8466a5546): complete subsection reference.

<a id="canonical-104fe9933558acaa7b45911652b34d0eb3c59b07f7cb1a1f1a9f50053a1289c5"></a>

<a id="canonical-35ba5ff1f08b7405e5a85e711a1b7330df138d51b1b96448db3eab2921bf99a0"></a>

## specific_domain property — api_protection_rules.api_groups_rules / a15f4f141ac4 / 6

Type: `"string"`. Optional.

Exclusive with \[any\_domain\] The rule will apply for a specific domain. For

Upstream description:

Exclusive with \[any\_domain\] The rule will apply for a specific domain. For example:
api.example.com.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

<a id="canonical-caa434ecd2b9bd23b9a9d84ccc333f30ee397c433ec76ff919ac937b595e9c31"></a>

## Next pages — api_protection_rules.api_groups_rules / a15f4f141ac4 / 7

- [api_protection_rules.api_groups_rules.action](resources--http_loadbalancer--reference--group-005.md#canonical-1de69af0465eaf9c52019846882d556ab548a1c547677838bd90200c082b7575)
- [api_protection_rules.api_groups_rules.any_domain](resources--http_loadbalancer--reference--group-005.md#canonical-e72f567b5726a2e75d7ab309fe0102ae78f73f6e5729ff463221cbe0381669d8)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-279540d1b6e4ec14980e15dedc82deacdfc801d2a7d6aa5fb34f6d87530d3849)
- [api_protection_rules.api_groups_rules.metadata](resources--http_loadbalancer--reference--group-005.md#canonical-da4775e72c6bbdc8a73bf5011a0d454b88dfe48c53393ab165beda163a01482f)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2829dee6801699124f354f49d1b1d6b69abd8bdd76c05d8cef9252c8466a5546)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-1de69af0465eaf9c52019846882d556ab548a1c547677838bd90200c082b7575"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-74a7e6951ef6428a28efad9f1e6664a85a34030b85cd16eb903415259997a98f"></a>

## api_protection_rules.api_groups_rules.action — api_protection_rules.api_groups_rules.action / 932d03111fb5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- api_protection_rules.api_groups_rules.action

<a id="canonical-13f2f9c1d7f50db4a50cc67636384745f62fd7af08142a19333907fff400cbef"></a>

Type: `"object"`. single nested block, Optional.

The action to take if the input request matches the rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("allow",
    "deny")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action": "[\"allow\",\"deny\"]"
}
```

Terraform syntax:

```terraform
action {
  # Configure direct properties listed below.
}
```

<a id="canonical-f737bea9e0c1b86d6da9e339189ae7b630e40462d34ae656bb18e29026251d2c"></a>

## Direct properties — api_protection_rules.api_groups_rules.action / 932d03111fb5 / 3

- [allow](resources--http_loadbalancer--reference--group-005.md#canonical-b5853ce2ff477435049ddf721149f6965acf23de524d9cf21fbbfcea8805cc41): complete subsection reference.

- [deny](resources--http_loadbalancer--reference--group-005.md#canonical-454ddfb103a808bbc0cd1b70794f548c66b6b8d5dcc4ccdd38784bf1f64ce346): complete subsection reference.

<a id="canonical-ee8f8c9c2c94c599d1ba150df6ec00c8f78e73901fd756c24f81628653085965"></a>

## Next pages — api_protection_rules.api_groups_rules.action / 932d03111fb5 / 4

- [api_protection_rules.api_groups_rules.action.allow](resources--http_loadbalancer--reference--group-005.md#canonical-b5853ce2ff477435049ddf721149f6965acf23de524d9cf21fbbfcea8805cc41)
- [api_protection_rules.api_groups_rules.action.deny](resources--http_loadbalancer--reference--group-005.md#canonical-454ddfb103a808bbc0cd1b70794f548c66b6b8d5dcc4ccdd38784bf1f64ce346)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b5853ce2ff477435049ddf721149f6965acf23de524d9cf21fbbfcea8805cc41"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-daa27e4933cc4e37f0ba053963429705697c9a8f176649ee3711ccc607b57ad4"></a>

## api_protection_rules.api_groups_rules.action.allow — api_protection_rules.api_groups_rules.action.allow / 446eae8a4b98 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [api_protection_rules.api_groups_rules.action](resources--http_loadbalancer--reference--group-005.md#canonical-1de69af0465eaf9c52019846882d556ab548a1c547677838bd90200c082b7575)
- api_protection_rules.api_groups_rules.action.allow

<a id="canonical-857d28ae9b1448046953b065ab378fbfc16b39f7081470e7e8524075aa857f31"></a>

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
allow = {}
```

<a id="canonical-06321120211b1119c587051725c64c1c0e15558a7d11e4f5d8d93008c9f4dea7"></a>

## Direct properties — api_protection_rules.api_groups_rules.action.allow / 446eae8a4b98 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1bcc1a33dd69f4c04959a83f0a420f032aa2068da7b6add89073d9e9e2754207"></a>

## Next pages — api_protection_rules.api_groups_rules.action.allow / 446eae8a4b98 / 4

- [api_protection_rules.api_groups_rules.action](resources--http_loadbalancer--reference--group-005.md#canonical-1de69af0465eaf9c52019846882d556ab548a1c547677838bd90200c082b7575)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-454ddfb103a808bbc0cd1b70794f548c66b6b8d5dcc4ccdd38784bf1f64ce346"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a48815662a074bd5f268f9021c55b32897f53847cdc5987ea47e6487c8fca3f"></a>

## api_protection_rules.api_groups_rules.action.deny — api_protection_rules.api_groups_rules.action.deny / 941bc537232a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [api_protection_rules.api_groups_rules.action](resources--http_loadbalancer--reference--group-005.md#canonical-1de69af0465eaf9c52019846882d556ab548a1c547677838bd90200c082b7575)
- api_protection_rules.api_groups_rules.action.deny

<a id="canonical-96a082007ae8dcbc38476c70d67de6394641bbe2a56a2266f6d70890a20cf351"></a>

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
deny = {}
```

<a id="canonical-cd84ac1ff3cd39469d2f0e4544e00912d8058e269139118962118476f76e848f"></a>

## Direct properties — api_protection_rules.api_groups_rules.action.deny / 941bc537232a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7ab94b67a1715f649e9c90c1e53849adc7a44fdd4b1cac899b5594b8f0bbe5b3"></a>

## Next pages — api_protection_rules.api_groups_rules.action.deny / 941bc537232a / 4

- [api_protection_rules.api_groups_rules.action](resources--http_loadbalancer--reference--group-005.md#canonical-1de69af0465eaf9c52019846882d556ab548a1c547677838bd90200c082b7575)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e72f567b5726a2e75d7ab309fe0102ae78f73f6e5729ff463221cbe0381669d8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78e7239f8efe8689a5540a2279b0cd8dc5f77e5b2904e7ae4f0de3574729c427"></a>

## api_protection_rules.api_groups_rules.any_domain — api_protection_rules.api_groups_rules.any_domain / 8250af2f2ce7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- api_protection_rules.api_groups_rules.any_domain

<a id="canonical-882e32ca3456284517d6cb66c6cf605933730329b5a70a8db4604b7bc1c06fa1"></a>

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
any_domain = {}
```

<a id="canonical-3d7de942644808eb06509cc63e747a096d815139d72dfd4e7cbba340e8f2c341"></a>

## Direct properties — api_protection_rules.api_groups_rules.any_domain / 8250af2f2ce7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f6e2f48ed3cd112b3109481a20919c0b2a7e32e0e5df2b7ad97b30edb29dd1c5"></a>

## Next pages — api_protection_rules.api_groups_rules.any_domain / 8250af2f2ce7 / 4

- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-279540d1b6e4ec14980e15dedc82deacdfc801d2a7d6aa5fb34f6d87530d3849"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e771bb5e443e1253c1d16b9489ad12974706038d720d15f495be6cc828c8c88"></a>

## api_protection_rules.api_groups_rules.client_matcher — api_protection_rules.api_groups_rules.client_matcher / 47745ec91486 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- api_protection_rules.api_groups_rules.client_matcher

<a id="canonical-a53b1538cb1da7ae59672955e5190320d4e326ccc219384ebb2ed57b1295d065"></a>

Type: `"object"`. single nested block, Optional.

Client Matcher. Client conditions for matching a rule.

Upstream description:

Client conditions for matching a rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("any_client",
    "client_selector"),
  validators.ConflictingObjectAttributes("any_client",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "asn_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("asn_list",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("asn_list",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_matcher",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("asn_matcher",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("client_selector",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("ip_matcher",
    "ip_prefix_list")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_selector\",\"ip_threat_category_list\"]",
  "x-ves-oneof-field-ip_asn_choice": "[\"any_ip\",\"asn_list\",\"asn_matcher\",\"ip_matcher\",\"ip_prefix_list\"]"
}
```

Terraform syntax:

```terraform
client_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-f38fa0d97207b2162e7816aba53619942b37e98c8453cea34386cb447e25b377"></a>

## Direct properties — api_protection_rules.api_groups_rules.client_matcher / 47745ec91486 / 3

- [any_client](resources--http_loadbalancer--reference--group-005.md#canonical-701facd37d518cac12be6ed09eeef23dc20b35f27135060e2601df6628fe33af): complete subsection reference.

- [any_ip](resources--http_loadbalancer--reference--group-005.md#canonical-381634cef37c6f43497bc3d30dbf14b1ba93bcf6a1415b4b9993eaaa9662b691): complete subsection reference.

- [asn_list](resources--http_loadbalancer--reference--group-005.md#canonical-b966ee2b2e2d932205fdcec2416e310592900b906b1c2b8084f824fa56eba85f): complete subsection reference.

- [asn_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-b503832ebd97be2edd31956b33af00b426715f6b55b5d40f6981b165764f35f7): complete subsection reference.

- [client_selector](resources--http_loadbalancer--reference--group-005.md#canonical-7c8742457c442ec3ae016d81ae1ad2b51173fc16378e0c05a6ffa30eda755839): complete subsection reference.

- [ip_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-b0ccc999bb6e8e34e25464b10a14a0bfea83caeca736e3e118bacd7ad2b54ec4): complete subsection reference.

- [ip_prefix_list](resources--http_loadbalancer--reference--group-005.md#canonical-83177f0c62cafae406951b0325c3f3ccf87ce13c1c20ff0c2f2ecbe683a42724): complete subsection reference.

- [ip_threat_category_list](resources--http_loadbalancer--reference--group-005.md#canonical-1b82d0d9acc74e3a5e52cc40c8c47aa544dba2086fe5e5e2ed607aafb6d11532): complete subsection reference.

- [tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-087357bd58b256e1a3bfa99c82f6d0f03053bc05eddf305ca2e8683fe61e163e): complete subsection reference.

<a id="canonical-521d4f2126b892f5dabedc039136a0bf4686c831eee54d804fa940e06a96e27a"></a>

## Next pages — api_protection_rules.api_groups_rules.client_matcher / 47745ec91486 / 4

- [api_protection_rules.api_groups_rules.client_matcher.any_client](resources--http_loadbalancer--reference--group-005.md#canonical-701facd37d518cac12be6ed09eeef23dc20b35f27135060e2601df6628fe33af)
- [api_protection_rules.api_groups_rules.client_matcher.any_ip](resources--http_loadbalancer--reference--group-005.md#canonical-381634cef37c6f43497bc3d30dbf14b1ba93bcf6a1415b4b9993eaaa9662b691)
- [api_protection_rules.api_groups_rules.client_matcher.asn_list](resources--http_loadbalancer--reference--group-005.md#canonical-b966ee2b2e2d932205fdcec2416e310592900b906b1c2b8084f824fa56eba85f)
- [api_protection_rules.api_groups_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-b503832ebd97be2edd31956b33af00b426715f6b55b5d40f6981b165764f35f7)
- [api_protection_rules.api_groups_rules.client_matcher.client_selector](resources--http_loadbalancer--reference--group-005.md#canonical-7c8742457c442ec3ae016d81ae1ad2b51173fc16378e0c05a6ffa30eda755839)
- [api_protection_rules.api_groups_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-b0ccc999bb6e8e34e25464b10a14a0bfea83caeca736e3e118bacd7ad2b54ec4)
- [api_protection_rules.api_groups_rules.client_matcher.ip_prefix_list](resources--http_loadbalancer--reference--group-005.md#canonical-83177f0c62cafae406951b0325c3f3ccf87ce13c1c20ff0c2f2ecbe683a42724)
- [api_protection_rules.api_groups_rules.client_matcher.ip_threat_category_list](resources--http_loadbalancer--reference--group-005.md#canonical-1b82d0d9acc74e3a5e52cc40c8c47aa544dba2086fe5e5e2ed607aafb6d11532)
- [api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-087357bd58b256e1a3bfa99c82f6d0f03053bc05eddf305ca2e8683fe61e163e)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-701facd37d518cac12be6ed09eeef23dc20b35f27135060e2601df6628fe33af"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-518c9c43a0138104ea40b589bb0b84b4ffc394010b87ce5838c9a538463bef8b"></a>

## api_protection_rules.api_groups_rules.client_matcher.any_client — api_protection_rules.api_groups_rules.client_matcher.any_client / 0f4159ffa5c0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-279540d1b6e4ec14980e15dedc82deacdfc801d2a7d6aa5fb34f6d87530d3849)
- api_protection_rules.api_groups_rules.client_matcher.any_client

<a id="canonical-07ae0c284035d4d0c673ae6b594ca3271aa87a4a27c0f25566f2506b9aaac7d7"></a>

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
any_client = {}
```

<a id="canonical-1d771a33f76039f568994f6fd04ba92f70fbdc4e4ad1fbca92909282272ad6e7"></a>

## Direct properties — api_protection_rules.api_groups_rules.client_matcher.any_client / 0f4159ffa5c0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2bdecf870b230b2918593d8cdb851f8627d7366b6a57a18a409e059cf5ea5699"></a>

## Next pages — api_protection_rules.api_groups_rules.client_matcher.any_client / 0f4159ffa5c0 / 4

- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-279540d1b6e4ec14980e15dedc82deacdfc801d2a7d6aa5fb34f6d87530d3849)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-381634cef37c6f43497bc3d30dbf14b1ba93bcf6a1415b4b9993eaaa9662b691"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a0af6668620fa5d461283479d3449c9279c923de339725dfb30a31267275e05"></a>

## api_protection_rules.api_groups_rules.client_matcher.any_ip — api_protection_rules.api_groups_rules.client_matcher.any_ip / b8662c806d3d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-279540d1b6e4ec14980e15dedc82deacdfc801d2a7d6aa5fb34f6d87530d3849)
- api_protection_rules.api_groups_rules.client_matcher.any_ip

<a id="canonical-f0115a6fe920ae4e64a8260a29223e8eb5e81e1a3115cd733a4cca4f6586a824"></a>

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
any_ip = {}
```

<a id="canonical-8a2959159eab71928e535654b9d7055b4eaa175c1a19d4020ba1868cf5e7e4d5"></a>

## Direct properties — api_protection_rules.api_groups_rules.client_matcher.any_ip / b8662c806d3d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-10de434bc9929273648548d62882706a237880e945be5bb4e85aacc2daf18351"></a>

## Next pages — api_protection_rules.api_groups_rules.client_matcher.any_ip / b8662c806d3d / 4

- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-279540d1b6e4ec14980e15dedc82deacdfc801d2a7d6aa5fb34f6d87530d3849)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b966ee2b2e2d932205fdcec2416e310592900b906b1c2b8084f824fa56eba85f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-846a350195e628fc9f806a400d6554589b36828d71f856fc39e498acdbe6cc03"></a>

## api_protection_rules.api_groups_rules.client_matcher.asn_list — api_protection_rules.api_groups_rules.client_matcher.asn_list / 323da7864226 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-279540d1b6e4ec14980e15dedc82deacdfc801d2a7d6aa5fb34f6d87530d3849)
- api_protection_rules.api_groups_rules.client_matcher.asn_list

<a id="canonical-5616a6a392afb3d02eade4f2bb47416d66012512b54e7886e230650dbceeac16"></a>

Type: `"object"`. single nested block, Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("as_numbers")}
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
asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-752c96201f62511c7a50389e0c698f2b9dc53e77b457874838ed19d78fd8a50d"></a>

## Direct properties — api_protection_rules.api_groups_rules.client_matcher.asn_list / 323da7864226 / 3

<a id="canonical-9a4edae67a4274cceab0652216614a672b67d72a6f8fa6636c879f5605d704fe"></a>

<a id="canonical-d9e495c8d80157d4364b6b2e48c9c37800a573c5d41777e08c80193180fbe090"></a>

## as_numbers property — api_protection_rules.api_groups_rules.client_matcher.asn_list / 323da7864226 / 4

Type: `["list", "number"]`. Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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

<a id="canonical-c219f08468e474e60940f6ebfd2e7700179596b2ce85140cc8d48e89bcf7d077"></a>

## Next pages — api_protection_rules.api_groups_rules.client_matcher.asn_list / 323da7864226 / 5

- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-279540d1b6e4ec14980e15dedc82deacdfc801d2a7d6aa5fb34f6d87530d3849)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b503832ebd97be2edd31956b33af00b426715f6b55b5d40f6981b165764f35f7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa4de8b05281bc5e07c7841388da39fd2774e0b3ecb6ac9bd4fc8a7f628df6a9"></a>

## api_protection_rules.api_groups_rules.client_matcher.asn_matcher — api_protection_rules.api_groups_rules.client_matcher.asn_matcher / 5f3d14455efd / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-279540d1b6e4ec14980e15dedc82deacdfc801d2a7d6aa5fb34f6d87530d3849)
- api_protection_rules.api_groups_rules.client_matcher.asn_matcher

<a id="canonical-085997e7be659895f3ae669c791b360559c446323816a94e2264d52812398777"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("asn_sets")}
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
asn_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-9889bdd235d7d112d93b051567e1c6d50a7383a888d64617fdfd4fe0303b3a7f"></a>

## Direct properties — api_protection_rules.api_groups_rules.client_matcher.asn_matcher / 5f3d14455efd / 3

- [asn_sets](resources--http_loadbalancer--reference--group-005.md#canonical-6a85093509ec91ab43ccb979c6827249e8925af6af0f48fdc31370767ea980cc): complete subsection reference.

<a id="canonical-411bc21d767cfb3ada8bc77918476f1471e360f4957c9a8958c6557a2d2c9a8f"></a>

## Next pages — api_protection_rules.api_groups_rules.client_matcher.asn_matcher / 5f3d14455efd / 4

- [api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets](resources--http_loadbalancer--reference--group-005.md#canonical-6a85093509ec91ab43ccb979c6827249e8925af6af0f48fdc31370767ea980cc)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-279540d1b6e4ec14980e15dedc82deacdfc801d2a7d6aa5fb34f6d87530d3849)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-6a85093509ec91ab43ccb979c6827249e8925af6af0f48fdc31370767ea980cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad7b9121f966d630667ea8471e6eb5e2cdf0c4ea2eb724cf7c959e6acc839113"></a>

## api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets — api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets / aa8549d992dc / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-279540d1b6e4ec14980e15dedc82deacdfc801d2a7d6aa5fb34f6d87530d3849)
- [api_protection_rules.api_groups_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-b503832ebd97be2edd31956b33af00b426715f6b55b5d40f6981b165764f35f7)
- api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-64e62c30459bff4af3fab86b0a07eaf01d4eb5a124bb456649bd3d825964f344"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
asn_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-835a201a04bcf6beb29e672e2ead7b6174e1b5542d1cefd77a9a4b4a802b6565"></a>

## Direct properties — api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets / aa8549d992dc / 3

<a id="canonical-f08864bb999c1280167299e40080e4a3ba7b0b5fc0851a832695026e8b1f66ce"></a>

<a id="canonical-385c2052d281da7a6daa43c3158e0c79526aeae2ade640d71203e3b67eeac2ef"></a>

## kind property — api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets / aa8549d992dc / 4

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

<a id="canonical-bfecb0d29f818fae125b32d91ee796edef9390d7ecff70ed4e2d91d40f87948a"></a>

<a id="canonical-131b68d04c96963387188ecd078f8c92f6adef1dd8cc7b00c2e3ee1db297ad6c"></a>

## name property — api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets / aa8549d992dc / 5

Type: `"string"`. Optional.

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

<a id="canonical-f4c371038ffc34896569f1a4213d943f75426c4da6d1e6302a746fbdd396133c"></a>

<a id="canonical-dd8a0b033549ba08099ef9607ede9d876ddcfcb2d55e0b557fe073629eab6626"></a>

## namespace property — api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets / aa8549d992dc / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-7126797487f76e039be79dbccf0af604fe327c9fbe6991369d54340a69558262"></a>

<a id="canonical-360b8e19c30c7f2d730555471c05113c03daf4869dabaf23ec7b598eb37cd689"></a>

## tenant property — api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets / aa8549d992dc / 7

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

<a id="canonical-3bc710809d9adfa4f712eb91cc83645302cfc3ac1d21508142e829903d6bd593"></a>

<a id="canonical-be41a1e8da8325be4c9b061b9f256dec1988e778517359379f9670173e78ed76"></a>

## uid property — api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets / aa8549d992dc / 8

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

<a id="canonical-0965823822980bcccdeea9c5ab6a62028814bdbd2a1ff4f27c6742f3d826dde4"></a>

## Next pages — api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets / aa8549d992dc / 9

- [api_protection_rules.api_groups_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-b503832ebd97be2edd31956b33af00b426715f6b55b5d40f6981b165764f35f7)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-7c8742457c442ec3ae016d81ae1ad2b51173fc16378e0c05a6ffa30eda755839"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1ac0c15e4740ee1950ba5e9f3d86f2141cfdb5cc54fdb1aa53f59591d7da4cd"></a>

## api_protection_rules.api_groups_rules.client_matcher.client_selector — api_protection_rules.api_groups_rules.client_matcher.client_selector / d03df9453f49 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-279540d1b6e4ec14980e15dedc82deacdfc801d2a7d6aa5fb34f6d87530d3849)
- api_protection_rules.api_groups_rules.client_matcher.client_selector

<a id="canonical-8129af5bbaff5c63e0b5d4fed3cdf3e4503ae3e5c44bcdf4e02be5645e49211a"></a>

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
client_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-74aecc3845c16cca99559ee8cd8a4aef176046b28b7b96ebe288e5dbca926479"></a>

## Direct properties — api_protection_rules.api_groups_rules.client_matcher.client_selector / d03df9453f49 / 3

<a id="canonical-a98a84df98a13b31dad6afc5c9c8ec8e965d36992d98642734b8a7418779d993"></a>

<a id="canonical-4029bc5540e385423e166c20a9cb2f237a4c9fd7b9303150a54a1bce71b7475d"></a>

## expressions property — api_protection_rules.api_groups_rules.client_matcher.client_selector / d03df9453f49 / 4

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

<a id="canonical-bfad7cce40b3f837422c1d13ee8e81b282e98258048478ab2bc285bc13ba7011"></a>

## Next pages — api_protection_rules.api_groups_rules.client_matcher.client_selector / d03df9453f49 / 5

- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-279540d1b6e4ec14980e15dedc82deacdfc801d2a7d6aa5fb34f6d87530d3849)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b0ccc999bb6e8e34e25464b10a14a0bfea83caeca736e3e118bacd7ad2b54ec4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e9ed1af02c35db040b711c36dccb68c1a6b02e580a2d0622c62e672fb79f1495"></a>

## api_protection_rules.api_groups_rules.client_matcher.ip_matcher — api_protection_rules.api_groups_rules.client_matcher.ip_matcher / 45f6a3c27eba / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-279540d1b6e4ec14980e15dedc82deacdfc801d2a7d6aa5fb34f6d87530d3849)
- api_protection_rules.api_groups_rules.client_matcher.ip_matcher

<a id="canonical-e7e22aaf094a824b7a35930f98ab9a5dfe6b5869a6d6b2a390e83554deb44d35"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

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

<a id="canonical-a22375eb2bc31ed8cb69aeb7c184422ddacbd2053a78bcb82794a800e0f4dcbc"></a>

## Direct properties — api_protection_rules.api_groups_rules.client_matcher.ip_matcher / 45f6a3c27eba / 3

<a id="canonical-8194267c6dbb32ce50a2876058cac83747b1589035b1a85859f1337939ba208b"></a>

<a id="canonical-10d5a730fa5477fd55af36029008986c64932707990b1bbe0fe31a4a97e2b804"></a>

## invert_matcher property — api_protection_rules.api_groups_rules.client_matcher.ip_matcher / 45f6a3c27eba / 4

Type: `"bool"`. Optional.

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

- [prefix_sets](resources--http_loadbalancer--reference--group-005.md#canonical-7b8765c56f1564423932733c0301a67eb95c49f2b8819a25b7eb4fb017cec54d): complete subsection reference.

<a id="canonical-66edd1825891010565dfe0b350ccd36b16f7cb19efd1724225dcb3530d8e41a8"></a>

## Next pages — api_protection_rules.api_groups_rules.client_matcher.ip_matcher / 45f6a3c27eba / 5

- [api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets](resources--http_loadbalancer--reference--group-005.md#canonical-7b8765c56f1564423932733c0301a67eb95c49f2b8819a25b7eb4fb017cec54d)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-279540d1b6e4ec14980e15dedc82deacdfc801d2a7d6aa5fb34f6d87530d3849)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-7b8765c56f1564423932733c0301a67eb95c49f2b8819a25b7eb4fb017cec54d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d86c2edd82768203252979c50426123f04938c522302209649f6bf256a49e78b"></a>

## api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets — api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets / 03a91113beb2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-279540d1b6e4ec14980e15dedc82deacdfc801d2a7d6aa5fb34f6d87530d3849)
- [api_protection_rules.api_groups_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-b0ccc999bb6e8e34e25464b10a14a0bfea83caeca736e3e118bacd7ad2b54ec4)
- api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-9b3fad992506bc70b45f620385d4dda2eb3be10f8f07ff0a61bb8b33b66017c7"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
prefix_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-d3c875d8551c8e0a5f5e3b3b9583b7756154efcf06863e3c4909e5b7687721cd"></a>

## Direct properties — api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets / 03a91113beb2 / 3

<a id="canonical-719a2a85092bf68278d300dd7c627a087bbc49e7c2017f0534267ebd1d3e9680"></a>

<a id="canonical-8b279f2707620b3e755df443a297061d9d99e004cb6ef783f936afee82b35637"></a>

## kind property — api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets / 03a91113beb2 / 4

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

<a id="canonical-7f68b1d62f5685d7110ed2889cac2e3745118cdf6cd363f19febcb5687baa38e"></a>

<a id="canonical-4292f68be047b5743af070ba4b4b8f8a26355a41b0219b965a4f3b7cea9c5a59"></a>

## name property — api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets / 03a91113beb2 / 5

Type: `"string"`. Optional.

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

<a id="canonical-d55aec34e34171e610879bab1b2d03b0a43ce77b48bdc62c25593a610f44398d"></a>

<a id="canonical-57bccb3b61f19eeaf2e327fee24475ec725ff1ff8799b3557ee1a99721170206"></a>

## namespace property — api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets / 03a91113beb2 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-f3d398e198891f8e17c7fbde51f1b4d4b1a9f67912c7729e2909532e4f39d64a"></a>

<a id="canonical-f3bd5375bca1dad60b2794d40e90be718fc4b60ad50a62cdcc6654b4d1baa2f6"></a>

## tenant property — api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets / 03a91113beb2 / 7

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

<a id="canonical-c30cca163a4789bb270420d80c1c13ac86b62e49418a641b029e8b81f4b89cd6"></a>

<a id="canonical-48c0b1ecdea1f54249a02357aaa60f5968365e7b44da5558ef5fc7dc8005de9c"></a>

## uid property — api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets / 03a91113beb2 / 8

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

<a id="canonical-6c311a5d3386cb32b22be7bd6d9be4bda7d1b5d6dd6166c219f642ea6490c5e7"></a>

## Next pages — api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets / 03a91113beb2 / 9

- [api_protection_rules.api_groups_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-b0ccc999bb6e8e34e25464b10a14a0bfea83caeca736e3e118bacd7ad2b54ec4)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-83177f0c62cafae406951b0325c3f3ccf87ce13c1c20ff0c2f2ecbe683a42724"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e464190464a850298b7b27f3fd27a862bb24c18a18bf7966012805b811022282"></a>

## api_protection_rules.api_groups_rules.client_matcher.ip_prefix_list — api_protection_rules.api_groups_rules.client_matcher.ip_prefix_list / fa7b92934752 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-279540d1b6e4ec14980e15dedc82deacdfc801d2a7d6aa5fb34f6d87530d3849)
- api_protection_rules.api_groups_rules.client_matcher.ip_prefix_list

<a id="canonical-8932f0bcb8b4ba41d8cf9ab19b10ff45a23aa559b409e9ad4a1f17c73b7a6e86"></a>

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

<a id="canonical-2f37825fdc3d688730a5d39ca4e5e67d116ffe6828b3a8ece45eb1b0f33d111d"></a>

## Direct properties — api_protection_rules.api_groups_rules.client_matcher.ip_prefix_list / fa7b92934752 / 3

<a id="canonical-dcb8daf0448f1e3ac0c2d91724f33fa4ca5653e514c9625b565c27d11aefcfaf"></a>

<a id="canonical-6b06281d0e649256fc5291264e08a3f49d4a0177a0d9b7d89f6fa779b0a685b5"></a>

## invert_match property — api_protection_rules.api_groups_rules.client_matcher.ip_prefix_list / fa7b92934752 / 4

Type: `"bool"`. Optional.

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

<a id="canonical-f8549de4810a4942fb69ae798f23f3b4ccbb0f8ed1198782e86d0e06c599ce6b"></a>

<a id="canonical-1fb47971a40d8c28a40f174b6bd2fbe570aeff6500295282337c1766bae24b36"></a>

## ip_prefixes property — api_protection_rules.api_groups_rules.client_matcher.ip_prefix_list / fa7b92934752 / 5

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

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

<a id="canonical-6c473f3cd4629c89e5c7f597891569d21d6bb3818ad115b30ed568ccf59f20c9"></a>

## Next pages — api_protection_rules.api_groups_rules.client_matcher.ip_prefix_list / fa7b92934752 / 6

- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-279540d1b6e4ec14980e15dedc82deacdfc801d2a7d6aa5fb34f6d87530d3849)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-1b82d0d9acc74e3a5e52cc40c8c47aa544dba2086fe5e5e2ed607aafb6d11532"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63a184190087776575ee41aefbb3873fb5b2244e8239532387afe086822209f8"></a>

## api_protection_rules.api_groups_rules.client_matcher.ip_threat_category_list — api_protection_rules.api_groups_rules.client_matcher.ip_threat_category_list / d40786867654 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-279540d1b6e4ec14980e15dedc82deacdfc801d2a7d6aa5fb34f6d87530d3849)
- api_protection_rules.api_groups_rules.client_matcher.ip_threat_category_list

<a id="canonical-c94dacb1ebfbff3649adfef45ba76382983d61f710f7e0cc0ad4ee33fc789c7c"></a>

Type: `"object"`. single nested block, Optional.

IP Threat Category List Type. List of IP threat categories.

Upstream description:

List of IP threat categories.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_threat_categories")}
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
ip_threat_category_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-599b32ab66bdb7019b089a6e7693f34c7eb5bfefdecc0e5e954e265bb571881b"></a>

## Direct properties — api_protection_rules.api_groups_rules.client_matcher.ip_threat_category_list / d40786867654 / 3

<a id="canonical-4b47dca0477bd7fe9323938d7a42525b3ea08111348672837f63c47ab181e264"></a>

<a id="canonical-b32c39b53b698fedf2f445df67b0fcabb0764e66261a2745ad935d952a4255ba"></a>

## ip_threat_categories property — api_protection_rules.api_groups_rules.client_matcher.ip_threat_category_list / d40786867654 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`,
\`WEB\_ATTACKS\`, \`BOTNETS\`, \`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`,
\`MOBILE\_THREATS\`, \`TOR\_PROXY\`, \`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to
\`SPAM\_SOURCES\`.

Upstream description:

The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-bba9e81f8226b6dffc1e8bdfdf665677b762220f9bc1f0030f621c07012e9728"></a>

## Next pages — api_protection_rules.api_groups_rules.client_matcher.ip_threat_category_list / d40786867654 / 5

- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-279540d1b6e4ec14980e15dedc82deacdfc801d2a7d6aa5fb34f6d87530d3849)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-087357bd58b256e1a3bfa99c82f6d0f03053bc05eddf305ca2e8683fe61e163e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c200ddc68ea5f768f5b346e53ee6beb68e932d53d7affeaddaf360b6614d44aa"></a>

## api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher — api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher / 881743b95ebe / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-279540d1b6e4ec14980e15dedc82deacdfc801d2a7d6aa5fb34f6d87530d3849)
- api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-c26d9955d0d3c5715a123fc19f616b3a880e125d6fe5f1726a2a4d6e7c97ce6d"></a>

Type: `"object"`. single nested block, Optional.

TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are
satisfied..

Upstream description:

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

<a id="canonical-beb4f25efc8c0818b95833e0171bb75480565a98d2cb5b7e05a150eb5524265b"></a>

## Direct properties — api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher / 881743b95ebe / 3

<a id="canonical-b96ceb2199bdd65e5e21b475f1e3c34408f17c176444f2d5e09133cb640c0f8d"></a>

<a id="canonical-c675ce82e254016db60aecb4948ba173fe4df54d58876ba672cba770f7cf2d5f"></a>

## classes property — api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher / 881743b95ebe / 4

Type: `["list", "string"]`. Optional.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Upstream description:

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-699e25500d07c9d590fff4ddbe85749dd5231334e51f74c8c2b50e226b40848e"></a>

<a id="canonical-8b8a55fb09506111d719a5e59a8f42f5c593226084cc4cc23dea64265dca37b2"></a>

## exact_values property — api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher / 881743b95ebe / 5

Type: `["list", "string"]`. Optional.

List of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Upstream description:

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

<a id="canonical-1dd48bf8f114803f78aa3bdd50d4b35dd436b7652c2079bc160a241e56d5b1b9"></a>

<a id="canonical-02cc9704dc5f07f51725fee2ef234bc51ab6bc757754d2957ddc915167fff1a9"></a>

## excluded_values property — api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher / 881743b95ebe / 6

Type: `["list", "string"]`. Optional.

List of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be
used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Upstream description:

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

<a id="canonical-b049aa2427cce86a3ffaca95fe077dec3261b6384492c9fa69c36ca679606ba4"></a>

## Next pages — api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher / 881743b95ebe / 7

- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-279540d1b6e4ec14980e15dedc82deacdfc801d2a7d6aa5fb34f6d87530d3849)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-da4775e72c6bbdc8a73bf5011a0d454b88dfe48c53393ab165beda163a01482f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-09c12a083d31558d5307c701eea6b9039a146a2f4f50f7e2aa5d49bdbb36c24e"></a>

## api_protection_rules.api_groups_rules.metadata — api_protection_rules.api_groups_rules.metadata / 097b751d5a41 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- api_protection_rules.api_groups_rules.metadata

<a id="canonical-e960ede7f2c9ca27c2ee1054fcc13e64dee759c56e923f73f61b1d6b921ad106"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-19cb5abe5eb0d3cecd052ce2a039c4ddb71aa5b15317246b4fbe6bc1af37bc95"></a>

## Direct properties — api_protection_rules.api_groups_rules.metadata / 097b751d5a41 / 3

<a id="canonical-4db054ed1db67f282561d1f8cee2cac360c5f300f90e98c2ec2c54d0e2d8bc4e"></a>

<a id="canonical-815025a6f1bf3541f3f742c5c44e6be5d9675c0c29cf6e2414afaeac37384d85"></a>

## description_spec property — api_protection_rules.api_groups_rules.metadata / 097b751d5a41 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-2fcbff1f92113d315f43a34abc77e68787b18067ab65d3dc81924a1f52b7db9c"></a>

<a id="canonical-6eeb047a15988d196143a9d31ed47d79759f3e44bfff508d1cc450f1795a2dc4"></a>

## name property — api_protection_rules.api_groups_rules.metadata / 097b751d5a41 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-780b1ae7ba06432d71aa2c558b0d714eda1d6387b2487acc48a7fd054261e09d"></a>

## Next pages — api_protection_rules.api_groups_rules.metadata / 097b751d5a41 / 6

- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-2829dee6801699124f354f49d1b1d6b69abd8bdd76c05d8cef9252c8466a5546"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
