---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-38488fc347d92b92b6a194a06da20426d2997716a9fcd06aaf6a0d53cf00fe0a"></a>

## uid property — api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets / 28065d210491 / 8

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

<a id="canonical-d60ba2087aa274d332416c5e2fb490e83ec5389e0aed5fc92809c849e751f432"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets / 28065d210491 / 9

- [api_rate_limit.server_url_rules.client_matcher.asn_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1f2a6b92564143901e3bdc1506de6e5fbdeb0c0a0c5d682f88bce33750e23a1d)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-52d7d988e583882abeab7b8a7381cef921e62e661c9612fe1ebbdcb3949a0e68"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d7154cf324f7b07cd1cf20993085155a4989e4a550c58554a68732a6a304360"></a>

## api_rate_limit.server_url_rules.client_matcher.client_selector — api_rate_limit.server_url_rules.client_matcher.client_selector / c567410a25f6 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-66b75078bf3f57cca6e480dadc7b071b3ab3423f328cf97b20e0eea3e110564b)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-bd87a783e34a923db2aed4a40eacd8f9e02254e0479b8f9940ab67baee99bcc5)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1471bc110f4f5f95bf0b957e3c3a12fe6806e7b2ef2b00d1424629b61ddecb0a)
- api_rate_limit.server_url_rules.client_matcher.client_selector

<a id="canonical-763576fde9c5175aa20fa84779ca29e49cc92ebd7d734d8c7338aa656a521466"></a>

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

<a id="canonical-2fa2341cd0e1b2ca73539568c220a2fa2678b68a40bfa78ffcf866b3fa4201e1"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.client_selector / c567410a25f6 / 3

<a id="canonical-c8e6debf1b428a804ab41c6f907954d1a8df6ce1e4cf9994080aab36b06469cf"></a>

<a id="canonical-9008bb91fb678e5a685cb3813751a2a59a22e34da22ef1125864cf4825ff5011"></a>

## expressions property — api_rate_limit.server_url_rules.client_matcher.client_selector / c567410a25f6 / 4

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

<a id="canonical-f5fc2eeb0b7fabc39df1491729d33df83d858bbc9ff8e1da4cb5efe9a0410a8d"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.client_selector / c567410a25f6 / 5

- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1471bc110f4f5f95bf0b957e3c3a12fe6806e7b2ef2b00d1424629b61ddecb0a)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-aed1cf0e2c5d751af587bf615bd3a8f4da1de2a42334945233528fbc5d1957dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8a03f42a427198ffa9e0d6360ef0f8f87e1a412bd26698ad78465b9653982f03"></a>

## api_rate_limit.server_url_rules.client_matcher.ip_matcher — api_rate_limit.server_url_rules.client_matcher.ip_matcher / 818a1f1d3dc9 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-66b75078bf3f57cca6e480dadc7b071b3ab3423f328cf97b20e0eea3e110564b)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-bd87a783e34a923db2aed4a40eacd8f9e02254e0479b8f9940ab67baee99bcc5)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1471bc110f4f5f95bf0b957e3c3a12fe6806e7b2ef2b00d1424629b61ddecb0a)
- api_rate_limit.server_url_rules.client_matcher.ip_matcher

<a id="canonical-c0d02defc1f7314cc5e9c78b2c496c8745cda683bb4190e2eb88681b1878b67a"></a>

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

<a id="canonical-bcf1370a53e8fa2f5e948f66f954fc593d89b1ac7962ee480373b5ab2f12a8fe"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.ip_matcher / 818a1f1d3dc9 / 3

<a id="canonical-688052fb3e9052ff716b686ac293ed5fefabe6fcfe58701268794c47f2d74169"></a>

<a id="canonical-f7b9f7db1b99674940c2fc41a227af1d8d7a8a2b9dce7662ab965f1ceddab6ad"></a>

## invert_matcher property — api_rate_limit.server_url_rules.client_matcher.ip_matcher / 818a1f1d3dc9 / 4

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

- [prefix_sets](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3e12947a3a4170a693d04e932a2408b6c7a8fd4e168df1dc84d1858a045c5a79): complete subsection reference.

<a id="canonical-358d73aab071df1d38f62a5ee1643279bef852dfc70ebf8ee4793f581a4b25b5"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.ip_matcher / 818a1f1d3dc9 / 5

- [api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3e12947a3a4170a693d04e932a2408b6c7a8fd4e168df1dc84d1858a045c5a79)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1471bc110f4f5f95bf0b957e3c3a12fe6806e7b2ef2b00d1424629b61ddecb0a)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-3e12947a3a4170a693d04e932a2408b6c7a8fd4e168df1dc84d1858a045c5a79"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-75eb69a53cdf051545bffd9dc5326774654d5031e69fb96734c9184c3459f96e"></a>

## api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / c3ac7fb9f26d / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-66b75078bf3f57cca6e480dadc7b071b3ab3423f328cf97b20e0eea3e110564b)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-bd87a783e34a923db2aed4a40eacd8f9e02254e0479b8f9940ab67baee99bcc5)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1471bc110f4f5f95bf0b957e3c3a12fe6806e7b2ef2b00d1424629b61ddecb0a)
- [api_rate_limit.server_url_rules.client_matcher.ip_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-aed1cf0e2c5d751af587bf615bd3a8f4da1de2a42334945233528fbc5d1957dc)
- api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-79b8a544d03f20067134c0cf83f454f7a8782a2abce34fc3af5a3fb1d861aaf9"></a>

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

<a id="canonical-95d2d95f8d707d0576df41db06e4679e3baf3154ca9a81e4dcd094b9f671934c"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / c3ac7fb9f26d / 3

<a id="canonical-f788205135ce7ef19eeffef074e7a5ffe613920e0988ef9fb484ba2f560ee5e1"></a>

<a id="canonical-84b6bcaba34cd5e6417bd9bfa05aa5535b9d6f7a4615606f0182a1b72759b15c"></a>

## kind property — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / c3ac7fb9f26d / 4

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

<a id="canonical-a0fcafe12f6cbc900072cad1062fa37f6b89ca3bdb7a162a7a790bb921a8ed45"></a>

<a id="canonical-187c8b2c858e3843be7f7d45bef4b4666bfd7b207844dd2f2dc6483d60dbb639"></a>

## name property — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / c3ac7fb9f26d / 5

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

<a id="canonical-47be64eb54789a2701c952b7978ef54ce22792c6740aa752d1e73af11937abd6"></a>

<a id="canonical-1b7c7e5ea2b5a96aa635a70a5a2ac18c4e42c5a3ade4865b20d710edad17a7f6"></a>

## namespace property — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / c3ac7fb9f26d / 6

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

<a id="canonical-f6f700d8479eddcf8971a43d3e901b1fc71638699e4557f6b795989c1e11d50c"></a>

<a id="canonical-12507056dc2951756484b0d5ac436ea60d27ae3c9c9934043650b58ae68e4bc5"></a>

## tenant property — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / c3ac7fb9f26d / 7

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

<a id="canonical-867b8de7127dd2144831840c745b23c70009c1df1dd46d6ba46fa1670034a451"></a>

<a id="canonical-b02a6399c01374c735753cf68ad6bb40e980702659fadf95d7f0d3ffca8af4f8"></a>

## uid property — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / c3ac7fb9f26d / 8

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

<a id="canonical-5fe7effe35c2cbcc6e29979d0bcd8e63145b6861de6db7e3ba7d7ae9b7b81369"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / c3ac7fb9f26d / 9

- [api_rate_limit.server_url_rules.client_matcher.ip_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-aed1cf0e2c5d751af587bf615bd3a8f4da1de2a42334945233528fbc5d1957dc)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-55916b94df00836ad4e43f3c92a7cc8f587f1fd358f053aa7b41f9476367c13b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1312014fb31f0cf5f117e7140332c9dce35e3f33c618d60c9272cb93c601c5b0"></a>

## api_rate_limit.server_url_rules.client_matcher.ip_prefix_list — api_rate_limit.server_url_rules.client_matcher.ip_prefix_list / fe5212d35283 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-66b75078bf3f57cca6e480dadc7b071b3ab3423f328cf97b20e0eea3e110564b)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-bd87a783e34a923db2aed4a40eacd8f9e02254e0479b8f9940ab67baee99bcc5)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1471bc110f4f5f95bf0b957e3c3a12fe6806e7b2ef2b00d1424629b61ddecb0a)
- api_rate_limit.server_url_rules.client_matcher.ip_prefix_list

<a id="canonical-99635e948cc04bdf0d8fc6ce9ff1258a0045939e4db6daaaab697af254e894dc"></a>

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

<a id="canonical-7e118be9128367d48984ce01b7d5eaa2de20529241bbe7c26af18ee02eb98862"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.ip_prefix_list / fe5212d35283 / 3

<a id="canonical-04ebb23401ca0ea26f3dcc99eb7a5c241b0bfbb5ed0d78c0b4b3976d4027d078"></a>

<a id="canonical-00bc5f394ea2b470a59ab9c945aa2cd67b12f62048b300784ad585f92ea10d45"></a>

## invert_match property — api_rate_limit.server_url_rules.client_matcher.ip_prefix_list / fe5212d35283 / 4

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

<a id="canonical-aec9410af74e14b3e705d85e858256fb469dffa50130cc4b849cf94cf52ef931"></a>

<a id="canonical-03098cdeee5f41e2e93f0abbe260a3272c6a9d4ca117d9e77a8d7710a3b939cc"></a>

## ip_prefixes property — api_rate_limit.server_url_rules.client_matcher.ip_prefix_list / fe5212d35283 / 5

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

<a id="canonical-fce61150067d2dc059490b0fe3a9f079f782dee828fa59a9cda30f06261ba469"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.ip_prefix_list / fe5212d35283 / 6

- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1471bc110f4f5f95bf0b957e3c3a12fe6806e7b2ef2b00d1424629b61ddecb0a)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-da2a07f712f751bac99800ab808ff8227af6cf8fa9703626ce0b3be5b39e154e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-842a9ee8184d111cc77f113c1ce093261908e5be19b0e94b9d9d55ec027aa062"></a>

## api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list — api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list / 665ddb0293a3 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-66b75078bf3f57cca6e480dadc7b071b3ab3423f328cf97b20e0eea3e110564b)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-bd87a783e34a923db2aed4a40eacd8f9e02254e0479b8f9940ab67baee99bcc5)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1471bc110f4f5f95bf0b957e3c3a12fe6806e7b2ef2b00d1424629b61ddecb0a)
- api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list

<a id="canonical-4629bd60983cbccfc5ee4f41a3b48f294a4b8b76081c84e1b1d9d27bedaac003"></a>

Type: `"single"`. Computed.

IP Threat Category List Type. List of IP threat categories.

Upstream description:

List of IP threat categories.

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

<a id="canonical-b5979b2cf09e0e19e2059570960529cce9bcd5e2f66a49531c1e9c4725cd6afc"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list / 665ddb0293a3 / 3

<a id="canonical-388a2ef5827ec9b894fd43ac9c343fa96f3af81daee5646541fea4d62358eb50"></a>

<a id="canonical-b64fbfc244bf58380795cb8d4d9d5ea000fc51abad7f734307cf242610666b8b"></a>

## ip_threat_categories property — api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list / 665ddb0293a3 / 4

Type: `["list", "string"]`. Computed.

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

<a id="canonical-b60d11c9585be6c52a46e1fde46d23f958c2108401d01e08f0ccefb8376d7e49"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list / 665ddb0293a3 / 5

- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1471bc110f4f5f95bf0b957e3c3a12fe6806e7b2ef2b00d1424629b61ddecb0a)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-7dfabd925e5b79b04bae70aabf9d600fddac45a1d59ef7e522c96883897470ba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2ced53325cd7bf8fad1485b6421e38bcd3885a0389eb33a541a8e1ffa424a44"></a>

## api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher — api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher / 4981eaa5eb34 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-66b75078bf3f57cca6e480dadc7b071b3ab3423f328cf97b20e0eea3e110564b)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-bd87a783e34a923db2aed4a40eacd8f9e02254e0479b8f9940ab67baee99bcc5)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1471bc110f4f5f95bf0b957e3c3a12fe6806e7b2ef2b00d1424629b61ddecb0a)
- api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-c97abbc9bf947d040eb895e22ae1016101bcaf9ec90fecd90a3d37dcd95106af"></a>

Type: `"single"`. Computed.

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

<a id="canonical-9f336e8b5c3f7aee91de6551afb533b85d7a45448897e3dbea2d52aa9ef0cf93"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher / 4981eaa5eb34 / 3

<a id="canonical-064bc548f4be1556be660f58413c189588cd0fb9bfd2fc526592a8fdc64f4953"></a>

<a id="canonical-7381a5b610d30a830362855584e7e0ccf87465f2347117e9b487194878a3b326"></a>

## classes property — api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher / 4981eaa5eb34 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Upstream description:

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

<a id="canonical-15272d77bd2aef5c2746ba51427dcca3a1ccf81509e0d24205933ba9d41f3c46"></a>

<a id="canonical-9b4869d6da26bfb8bdbbf1655dbaf16994f18f787da0c81884b45756d6a2c7fb"></a>

## exact_values property — api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher / 4981eaa5eb34 / 5

Type: `["list", "string"]`. Computed.

List of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Upstream description:

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

<a id="canonical-b730766c7a6f78bf48ac7961fa441afe299a5cd563e09060587e9c3bc8ac9950"></a>

<a id="canonical-a793911f53cec39b983ff272067c4e3ed822da1031011b2f23bc1df3f8d6f082"></a>

## excluded_values property — api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher / 4981eaa5eb34 / 6

Type: `["list", "string"]`. Computed.

List of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be
used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Upstream description:

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

<a id="canonical-613686c4c150e3990e4f8edb296b56b3b6b66ecc7387b4528694ed8872ec233b"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher / 4981eaa5eb34 / 7

- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1471bc110f4f5f95bf0b957e3c3a12fe6806e7b2ef2b00d1424629b61ddecb0a)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-c88dd859394c050900079c63686c12c8ec9d10533d44ba0cebf21da6eee6baf8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-055b972877a769b740dc427aae3a3e45245d398031fb48e9776fff3b93b6cc62"></a>

## api_rate_limit.server_url_rules.inline_rate_limiter — api_rate_limit.server_url_rules.inline_rate_limiter / e4ee609fe4c3 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-66b75078bf3f57cca6e480dadc7b071b3ab3423f328cf97b20e0eea3e110564b)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-bd87a783e34a923db2aed4a40eacd8f9e02254e0479b8f9940ab67baee99bcc5)
- api_rate_limit.server_url_rules.inline_rate_limiter

<a id="canonical-1431dca83799dfda6fd33c17d879acf2be4b79c39e04dc9a82228f5ae0d9300d"></a>

Type: `"single"`. Computed.

Inline rate-limiter settings for this domain, base-path, or endpoint rule. Select this field as the
required rate\_limiter\_choice when no stored rate-limiter object is used.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-count_by_choice": "[\"ref_user_id\",\"use_http_lb_user_id\"]"
}
```

<a id="canonical-5bc91c247b24b8539f979d56e8a97ebd9bb20d158d3f0178f8fbc5852c4e4b35"></a>

## Direct properties — api_rate_limit.server_url_rules.inline_rate_limiter / e4ee609fe4c3 / 3

- [ref_user_id](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-ed7f4018e53b28f97ef68739f76b20e7c6bdd7d1f91fddb193e56f79b4fdb35c): complete subsection reference.

<a id="canonical-2d623a3d7accf7aaa7a1197639782ace9363828c53af60fa4740b4fa70b27fd7"></a>

<a id="canonical-591e56f876c85e1a31714c753afc8f0f2ff1d7838e5d2a16421f06c343e80a73"></a>

## threshold property — api_rate_limit.server_url_rules.inline_rate_limiter / e4ee609fe4c3 / 4

Type: `"number"`. Computed.

The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified
period.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  }
}
```

<a id="canonical-7192550bb1c2933f27bdd87f7424f3a5fe40c2b07907265dbbee005ed1154dbf"></a>

<a id="canonical-5939691e0e1443e25d03bb144ff672805b5cf9dcc59529a5a94c67be6ae7e8e9"></a>

## unit property — api_rate_limit.server_url_rules.inline_rate_limiter / e4ee609fe4c3 / 5

Type: `"string"`. Computed.

\[Enum: SECOND|MINUTE|HOUR\] Unit for the period per which the rate limit is applied. - SECOND:
Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR:
Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days. Possible values are
\`SECOND\`, \`MINUTE\`, \`HOUR\`. Defaults to \`SECOND\`.

Upstream description:

Unit for the period per which the rate limit is applied.

&#8203;- SECOND: Second

Rate limit period unit is seconds &#8203;- MINUTE: Minute

Rate limit period unit is minutes &#8203;- HOUR: Hour

Rate limit period unit is hours &#8203;- DAY: Day

Rate limit period unit is days.

Receipt-pinned upstream constraints:

```json
{
  "default": "SECOND",
  "enum": [
    "SECOND",
    "MINUTE",
    "HOUR"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [use_http_lb_user_id](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-f30d4d6d70437a1e4f34ebfc67e98d23fdb71fe7d900566099a6ef8204b9a196): complete subsection reference.

<a id="canonical-33c23a0135d530f672f792c1ae246899bfad9df6b88a6ae25b2f5ae223776049"></a>

## Next pages — api_rate_limit.server_url_rules.inline_rate_limiter / e4ee609fe4c3 / 6

- [api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-ed7f4018e53b28f97ef68739f76b20e7c6bdd7d1f91fddb193e56f79b4fdb35c)
- [api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-f30d4d6d70437a1e4f34ebfc67e98d23fdb71fe7d900566099a6ef8204b9a196)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-bd87a783e34a923db2aed4a40eacd8f9e02254e0479b8f9940ab67baee99bcc5)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-ed7f4018e53b28f97ef68739f76b20e7c6bdd7d1f91fddb193e56f79b4fdb35c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9b4bd7d425a8bc490e0faf66f0c94aefdfe777f813eca20930441ec9936cb5e"></a>

## api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id — api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id / cdc43efe6306 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-66b75078bf3f57cca6e480dadc7b071b3ab3423f328cf97b20e0eea3e110564b)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-bd87a783e34a923db2aed4a40eacd8f9e02254e0479b8f9940ab67baee99bcc5)
- [api_rate_limit.server_url_rules.inline_rate_limiter](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-c88dd859394c050900079c63686c12c8ec9d10533d44ba0cebf21da6eee6baf8)
- api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id

<a id="canonical-7d23dd6788d1a242f64611d9ae71c2f9397e4e38d1ce21daefd57999eb169263"></a>

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

<a id="canonical-bbaab355ddff2b0d8dc5b2796f443facec651cb23e2a83a05b69e7275795f86a"></a>

## Direct properties — api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id / cdc43efe6306 / 3

<a id="canonical-c3d1b5ac2e8f07fd2e85d46fdf29582d239460a4ce4417003d28143bb64b509a"></a>

<a id="canonical-b1f55a855d7e8737c9c90b4fc1c850e08130f2f20e95ac0a3cfc5427c3859aa7"></a>

## name property — api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id / cdc43efe6306 / 4

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

<a id="canonical-ea69e567c0e6ee31f3c38131b9e13ca56232acdfd0e2cb8b59e02ae0dd694f41"></a>

<a id="canonical-da53573ab4a8a3fb87bd83ba8ed0905a3b195e83dca1a01180b5eb228d169134"></a>

## namespace property — api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id / cdc43efe6306 / 5

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

<a id="canonical-18adb7c13730f8bdd4c6ae267ec8e75c15c14efae5fe7ed1c7a59bc47a75101a"></a>

<a id="canonical-9bb1d13d2406c2d3a6ef893d0def24735eb8842239a09286dfd8b7e4d022c347"></a>

## tenant property — api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id / cdc43efe6306 / 6

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

<a id="canonical-a717db5583173042a74f17999cf053c35a4c8979b2211c8840781ff8a6ae6ba9"></a>

## Next pages — api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id / cdc43efe6306 / 7

- [api_rate_limit.server_url_rules.inline_rate_limiter](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-c88dd859394c050900079c63686c12c8ec9d10533d44ba0cebf21da6eee6baf8)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-f30d4d6d70437a1e4f34ebfc67e98d23fdb71fe7d900566099a6ef8204b9a196"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dabc19c8f6e4fbb3c3c36d1bde57c8258983e9dd6179921fa64244a76dd36788"></a>

## api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id — api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id / 677cd72c619f / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-66b75078bf3f57cca6e480dadc7b071b3ab3423f328cf97b20e0eea3e110564b)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-bd87a783e34a923db2aed4a40eacd8f9e02254e0479b8f9940ab67baee99bcc5)
- [api_rate_limit.server_url_rules.inline_rate_limiter](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-c88dd859394c050900079c63686c12c8ec9d10533d44ba0cebf21da6eee6baf8)
- api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id

<a id="canonical-3febdc27859f695cb92e55aec6d3b75b23587e4911fcb013455dfdc69f296659"></a>

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

<a id="canonical-a415481da3c78230fdc82de8e282e84173721043b564e86e589da5ecd106b702"></a>

## Direct properties — api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id / 677cd72c619f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3ba96faf3a5feb21a5557d59c169a85378385735e10a51c463cdc8af28eb088e"></a>

## Next pages — api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id / 677cd72c619f / 4

- [api_rate_limit.server_url_rules.inline_rate_limiter](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-c88dd859394c050900079c63686c12c8ec9d10533d44ba0cebf21da6eee6baf8)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-590dae434c5cba8158ec25da533f2dfe28160c939858c53a7899e3e18b857844"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad6e943822918228ba2d7dd68c1a266bb47a18459a3950f6148aa7b44ec7af1f"></a>

## api_rate_limit.server_url_rules.ref_rate_limiter — api_rate_limit.server_url_rules.ref_rate_limiter / 3ef949027f67 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-66b75078bf3f57cca6e480dadc7b071b3ab3423f328cf97b20e0eea3e110564b)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-bd87a783e34a923db2aed4a40eacd8f9e02254e0479b8f9940ab67baee99bcc5)
- api_rate_limit.server_url_rules.ref_rate_limiter

<a id="canonical-c6302d1c90665b1610b84fd1354d9b71e27c98cb32a7a25405707c6729cee136"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

Reference to a stored rate-limiter object for this scoped rule. Select exactly one of
ref\_rate\_limiter and inline\_rate\_limiter.

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

<a id="canonical-3363f2db8de4f5afd4e2becffc7dacd7762bee589aa6fa6a062cb0e261d080d6"></a>

## Direct properties — api_rate_limit.server_url_rules.ref_rate_limiter / 3ef949027f67 / 3

<a id="canonical-bf8eaf3bbe26bc6a4b8834fe422f6858d68664c1ec1e92009c988943bd428219"></a>

<a id="canonical-8e3d2f9c3a2c1fd3bce66a5f227b05a9d8d8dd9ab892d7a2c899d95e1a4ba50a"></a>

## name property — api_rate_limit.server_url_rules.ref_rate_limiter / 3ef949027f67 / 4

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

<a id="canonical-4725e1ede499d586110071d1b34f6b7c2a9c69c1608f9d911aaacbd824844c56"></a>

<a id="canonical-f4e08921e6817687bd3d156e02c9292d81830bd0d1abd528686f943e2d089971"></a>

## namespace property — api_rate_limit.server_url_rules.ref_rate_limiter / 3ef949027f67 / 5

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

<a id="canonical-b05745856449d89c870d24a9e2fd085620b57f66b25fd73cf6ffd91ea528373f"></a>

<a id="canonical-e40e12838ad76cddb45ab5d310178723eeeb9159675cf2acbed32f3ba9b9d2dd"></a>

## tenant property — api_rate_limit.server_url_rules.ref_rate_limiter / 3ef949027f67 / 6

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

<a id="canonical-434fe08e4e352063f3aad5c3d23345b92d6f43fe48a78176047952e5d471f651"></a>

## Next pages — api_rate_limit.server_url_rules.ref_rate_limiter / 3ef949027f67 / 7

- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-bd87a783e34a923db2aed4a40eacd8f9e02254e0479b8f9940ab67baee99bcc5)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-fd63eb81ac8d0d1e1d4ddeb9a3d4d03a0e0b413d6edc24fa4ae959a3db2d0433"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e26b43f7a10e48c4759211c9122b4418ce77fcff5db80d0f224a4280266494f"></a>

## api_rate_limit.server_url_rules.request_matcher — api_rate_limit.server_url_rules.request_matcher / ebf86d4ac1c3 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-66b75078bf3f57cca6e480dadc7b071b3ab3423f328cf97b20e0eea3e110564b)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-bd87a783e34a923db2aed4a40eacd8f9e02254e0479b8f9940ab67baee99bcc5)
- api_rate_limit.server_url_rules.request_matcher

<a id="canonical-03cdb1bcb3728920b51e8e9ffc63241cdf2a89696b14dad71bf2efacc88e2d13"></a>

Type: `"single"`. Computed.

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

<a id="canonical-c3bc490a61bf600b98256be3194e612ffc10fce62c0657f9a448906368212c7b"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher / ebf86d4ac1c3 / 3

- [cookie_matchers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6bf8f3ea12c2438a923ab3e0c3972465d42b2a67b67c17d2082a1e9cc0370fb5): complete subsection reference.

- [headers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6b2dc62cd9ab58db228a9d79fcadea2eaabfe4d873afb68b96a9bb9fd6666cd2): complete subsection reference.

- [jwt_claims](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-ef6c803ff6f47ffadbc7494fa2296d086e29a5dea67b11f78c72a13f285daf1e): complete subsection reference.

- [query_params](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-5b383596ed3f7654770c1d6a7ab92bbe03da3a6f81ed2cdefeb100c256061b84): complete subsection reference.

<a id="canonical-8d8d7788a86da3d4707d92fbdc09cf12c2b5489ba92505f76b8264eec484e0c7"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher / ebf86d4ac1c3 / 4

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6bf8f3ea12c2438a923ab3e0c3972465d42b2a67b67c17d2082a1e9cc0370fb5)
- [api_rate_limit.server_url_rules.request_matcher.headers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6b2dc62cd9ab58db228a9d79fcadea2eaabfe4d873afb68b96a9bb9fd6666cd2)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-ef6c803ff6f47ffadbc7494fa2296d086e29a5dea67b11f78c72a13f285daf1e)
- [api_rate_limit.server_url_rules.request_matcher.query_params](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-5b383596ed3f7654770c1d6a7ab92bbe03da3a6f81ed2cdefeb100c256061b84)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-bd87a783e34a923db2aed4a40eacd8f9e02254e0479b8f9940ab67baee99bcc5)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-6bf8f3ea12c2438a923ab3e0c3972465d42b2a67b67c17d2082a1e9cc0370fb5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-50f723d3986fe6550239e53f4727e142c69e5db550d043e2e2cd6d999c6df4cc"></a>

## api_rate_limit.server_url_rules.request_matcher.cookie_matchers — api_rate_limit.server_url_rules.request_matcher.cookie_matchers / 3659232e6e0f / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-66b75078bf3f57cca6e480dadc7b071b3ab3423f328cf97b20e0eea3e110564b)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-bd87a783e34a923db2aed4a40eacd8f9e02254e0479b8f9940ab67baee99bcc5)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-fd63eb81ac8d0d1e1d4ddeb9a3d4d03a0e0b413d6edc24fa4ae959a3db2d0433)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers

<a id="canonical-b9db8f5ec10f95bd0491e4e1337d15b382435b32b958ad1280509418db9e192a"></a>

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

<a id="canonical-512322ca3162ca7d2445b073b399c69a868d5ca857f8ef591983a9bb1a282bbc"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.cookie_matchers / 3659232e6e0f / 3

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-38e934c8b7deee8d00e4b1ca2b575d48da2b4538ec55adc1188dc9bf4e090fbe): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-8ad1467e418c748fcd27455d23bce77dfa96466a0b55c11b46190fec21f96a21): complete subsection reference.

<a id="canonical-dcdfcf28a1c398c3c20e11d983d2ed42e11015bd4f9718da6420d79af43503fe"></a>

<a id="canonical-94c54048c7ad0fb2d2bfa64c068dd79c1915b0f3b934cea2aabbd64f3b973f0e"></a>

## invert_matcher property — api_rate_limit.server_url_rules.request_matcher.cookie_matchers / 3659232e6e0f / 4

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

- [item](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-c40670da3cbdcff8917c3820b34beaf94ff87f27e8b3beb1fb2c92c3ed332f35): complete subsection reference.

<a id="canonical-ffc807dfad0d1425f582cacfd420d9420d0234591dd4579c7d1ffbd1ed602367"></a>

<a id="canonical-c4c179b550a0c7e85ec3d3e8bb0153397b97d5a6ad11c6043fd667725eb43530"></a>

## name property — api_rate_limit.server_url_rules.request_matcher.cookie_matchers / 3659232e6e0f / 5

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

<a id="canonical-4210d868a202d6f324eb6b1416ce524c1cb759c513750409f31a6e526952bc92"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.cookie_matchers / 3659232e6e0f / 6

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-38e934c8b7deee8d00e4b1ca2b575d48da2b4538ec55adc1188dc9bf4e090fbe)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-8ad1467e418c748fcd27455d23bce77dfa96466a0b55c11b46190fec21f96a21)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-c40670da3cbdcff8917c3820b34beaf94ff87f27e8b3beb1fb2c92c3ed332f35)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-fd63eb81ac8d0d1e1d4ddeb9a3d4d03a0e0b413d6edc24fa4ae959a3db2d0433)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-38e934c8b7deee8d00e4b1ca2b575d48da2b4538ec55adc1188dc9bf4e090fbe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-38703eba17739ade1748fcb2c52539784583202cf113fdf842e7324f5b4bb1eb"></a>

## api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_present — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_presen / 987caf0d46cc / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-66b75078bf3f57cca6e480dadc7b071b3ab3423f328cf97b20e0eea3e110564b)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-bd87a783e34a923db2aed4a40eacd8f9e02254e0479b8f9940ab67baee99bcc5)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-fd63eb81ac8d0d1e1d4ddeb9a3d4d03a0e0b413d6edc24fa4ae959a3db2d0433)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6bf8f3ea12c2438a923ab3e0c3972465d42b2a67b67c17d2082a1e9cc0370fb5)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-f1bf1ba04faadafac9f9b21677a86a7e26c14639fb91522c2e03c1a275cae2c0"></a>

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

<a id="canonical-bca4591f9b2a322a0c4e108da264b0da5ca30f85ae11351c286e19f325edf17e"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_presen / 987caf0d46cc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-53fc700cfd9b7589e0eea50ebe4bd0682caf30f7db91e1f6b119d90eaa9d4e2f"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_presen / 987caf0d46cc / 4

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6bf8f3ea12c2438a923ab3e0c3972465d42b2a67b67c17d2082a1e9cc0370fb5)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-8ad1467e418c748fcd27455d23bce77dfa96466a0b55c11b46190fec21f96a21"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-07b9db52f5e91e7d7d371138a8947c3b76e4d9959273c3e9664f8daeeda9445d"></a>

## api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present / 2f29e9452c63 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-66b75078bf3f57cca6e480dadc7b071b3ab3423f328cf97b20e0eea3e110564b)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-bd87a783e34a923db2aed4a40eacd8f9e02254e0479b8f9940ab67baee99bcc5)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-fd63eb81ac8d0d1e1d4ddeb9a3d4d03a0e0b413d6edc24fa4ae959a3db2d0433)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6bf8f3ea12c2438a923ab3e0c3972465d42b2a67b67c17d2082a1e9cc0370fb5)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-f5333f39abab1a8c6156b85102379b6709fde6d8a5d1f1375af11aa57b0a9411"></a>

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

<a id="canonical-eeb30fd45ad595ad815e36ead19a654f7120725c2da613075628f75188510234"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present / 2f29e9452c63 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-842e98b96f0447a57d3831793208d99417a246264ee71d5a6dbff080d6191c95"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present / 2f29e9452c63 / 4

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6bf8f3ea12c2438a923ab3e0c3972465d42b2a67b67c17d2082a1e9cc0370fb5)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-c40670da3cbdcff8917c3820b34beaf94ff87f27e8b3beb1fb2c92c3ed332f35"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d1318d0097c6b8d66080d55a73d0b1c8cc60319bdd34c5ba15fd313e2740b7d7"></a>

## api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item / e4ac6ac3a8d3 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-66b75078bf3f57cca6e480dadc7b071b3ab3423f328cf97b20e0eea3e110564b)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-bd87a783e34a923db2aed4a40eacd8f9e02254e0479b8f9940ab67baee99bcc5)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-fd63eb81ac8d0d1e1d4ddeb9a3d4d03a0e0b413d6edc24fa4ae959a3db2d0433)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6bf8f3ea12c2438a923ab3e0c3972465d42b2a67b67c17d2082a1e9cc0370fb5)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item

<a id="canonical-aa3ce0c88039f52efc6f369d9a4776221cc59d0de8fd619dc41b0052066b58fe"></a>

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

<a id="canonical-8bf1f363eba878921efcfe3a0339d660fc873384a7bbe83959dbe33e3bb9c816"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item / e4ac6ac3a8d3 / 3

<a id="canonical-75a2ca91c73b8750c57ec1ccce47557ee8774bdee7b54971d1d737ca30952874"></a>

<a id="canonical-e7ccd23bc29d57cd7242dd8a1a331b717571c4432009f84e61a0aab25338a7a3"></a>

## exact_values property — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item / e4ac6ac3a8d3 / 4

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

<a id="canonical-c8b3c89d4b75584d5b149e83586843a27eaa7556b578fbef965cc67af4ce2de7"></a>

<a id="canonical-29e79e52e396967c9b7b62092afc554909287701e2102d438546422ee01ce5b5"></a>

## regex_values property — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item / e4ac6ac3a8d3 / 5

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

<a id="canonical-c768d8bdcf5719169ac124bdb932cd84fc1b01b6a0e334cc0602b6d868a21fff"></a>

<a id="canonical-6041dd6ab088bc1563ddbf7436e5b8243cb870bf0ea3e940c3f8cd72eafd6df4"></a>

## transformers property — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item / e4ac6ac3a8d3 / 6

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

<a id="canonical-a10c7d712e09bd40842e12008499205e4cd0f01de38cf75ecc6bf7bdecaad830"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item / e4ac6ac3a8d3 / 7

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6bf8f3ea12c2438a923ab3e0c3972465d42b2a67b67c17d2082a1e9cc0370fb5)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-6b2dc62cd9ab58db228a9d79fcadea2eaabfe4d873afb68b96a9bb9fd6666cd2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f7fd56c8069b4455e5e7661018eec9356cd7744945f4be8015707d7a35751321"></a>

## api_rate_limit.server_url_rules.request_matcher.headers — api_rate_limit.server_url_rules.request_matcher.headers / e568c5446e07 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-66b75078bf3f57cca6e480dadc7b071b3ab3423f328cf97b20e0eea3e110564b)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-bd87a783e34a923db2aed4a40eacd8f9e02254e0479b8f9940ab67baee99bcc5)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-fd63eb81ac8d0d1e1d4ddeb9a3d4d03a0e0b413d6edc24fa4ae959a3db2d0433)
- api_rate_limit.server_url_rules.request_matcher.headers

<a id="canonical-02e1450d9cbac56a45310eb48dd98947a13509f52ba2ee1fdb7bc4b593348f0e"></a>

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

<a id="canonical-e714f9d5ce33da1e37c6ea133d8a29080ace356ac5b718474de666e154a13db7"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.headers / e568c5446e07 / 3

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-044e84cf1e45dfff75063ca449bb7cf9b38d4bbd6ce94aa20ec97492d6128340): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-8bb4a0005d5fef9e9dbfc1294701814892a1d3c9f82f4e1213d7f2622e3d3629): complete subsection reference.

<a id="canonical-22065114b1b78d7b3536c84b769d6e0bf0041a85c5b4913af8f5f41bb001f15b"></a>

<a id="canonical-727872b9e7d0e1fefab838d234f21ce67b395d278fef3c9c88540ab3192813ce"></a>

## invert_matcher property — api_rate_limit.server_url_rules.request_matcher.headers / e568c5446e07 / 4

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

- [item](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-cad11c267a1773352dfc8fef058604349be18d1367e2f1acec5821a0db4c4e2f): complete subsection reference.

<a id="canonical-013862c37c4c4012ebf7e9e2d0501e6f7229824bdf8c4195122b1ae30c995ed6"></a>

<a id="canonical-a4cd8158f52a7e9e71487580fb4883c6bf6fbc9e61678771588a61e576e9c59d"></a>

## name property — api_rate_limit.server_url_rules.request_matcher.headers / e568c5446e07 / 5

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

<a id="canonical-fb223264e75cdbffed5874a120df1be5265f7d16c0455cf7c8836af1a0b1b6bc"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.headers / e568c5446e07 / 6

- [api_rate_limit.server_url_rules.request_matcher.headers.check_not_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-044e84cf1e45dfff75063ca449bb7cf9b38d4bbd6ce94aa20ec97492d6128340)
- [api_rate_limit.server_url_rules.request_matcher.headers.check_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-8bb4a0005d5fef9e9dbfc1294701814892a1d3c9f82f4e1213d7f2622e3d3629)
- [api_rate_limit.server_url_rules.request_matcher.headers.item](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-cad11c267a1773352dfc8fef058604349be18d1367e2f1acec5821a0db4c4e2f)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-fd63eb81ac8d0d1e1d4ddeb9a3d4d03a0e0b413d6edc24fa4ae959a3db2d0433)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-044e84cf1e45dfff75063ca449bb7cf9b38d4bbd6ce94aa20ec97492d6128340"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d7e4c1d9fa0b999fc4a70223b7104596b39cf95256471a12f76854e9f2ad837"></a>

## api_rate_limit.server_url_rules.request_matcher.headers.check_not_present — api_rate_limit.server_url_rules.request_matcher.headers.check_not_present / ce745724b7cc / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-66b75078bf3f57cca6e480dadc7b071b3ab3423f328cf97b20e0eea3e110564b)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-bd87a783e34a923db2aed4a40eacd8f9e02254e0479b8f9940ab67baee99bcc5)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-fd63eb81ac8d0d1e1d4ddeb9a3d4d03a0e0b413d6edc24fa4ae959a3db2d0433)
- [api_rate_limit.server_url_rules.request_matcher.headers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6b2dc62cd9ab58db228a9d79fcadea2eaabfe4d873afb68b96a9bb9fd6666cd2)
- api_rate_limit.server_url_rules.request_matcher.headers.check_not_present

<a id="canonical-ce431981c56d4c63b6842d52157051fda0362b04efc6a1e5c5606b762efa3c5e"></a>

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

<a id="canonical-60c99f6d89c58465585db538fc412d2da9656d3a3f9eca114d445e1b4d136405"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.headers.check_not_present / ce745724b7cc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6dabdc86c671e13633b606a7412db37d0bd119f3c4afdadefaaeb6ec253e0aa7"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.headers.check_not_present / ce745724b7cc / 4

- [api_rate_limit.server_url_rules.request_matcher.headers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6b2dc62cd9ab58db228a9d79fcadea2eaabfe4d873afb68b96a9bb9fd6666cd2)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-8bb4a0005d5fef9e9dbfc1294701814892a1d3c9f82f4e1213d7f2622e3d3629"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-17d730ac3b0e4db8c33fac0910414da6e49e81a883fe827c7edc69f9d574912e"></a>

## api_rate_limit.server_url_rules.request_matcher.headers.check_present — api_rate_limit.server_url_rules.request_matcher.headers.check_present / 2ee100c42a76 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-66b75078bf3f57cca6e480dadc7b071b3ab3423f328cf97b20e0eea3e110564b)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-bd87a783e34a923db2aed4a40eacd8f9e02254e0479b8f9940ab67baee99bcc5)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-fd63eb81ac8d0d1e1d4ddeb9a3d4d03a0e0b413d6edc24fa4ae959a3db2d0433)
- [api_rate_limit.server_url_rules.request_matcher.headers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6b2dc62cd9ab58db228a9d79fcadea2eaabfe4d873afb68b96a9bb9fd6666cd2)
- api_rate_limit.server_url_rules.request_matcher.headers.check_present

<a id="canonical-a6ba287d3c422c35b81c1b2a91289d392e368b7b68a7e37b092960a5821a6e7b"></a>

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

<a id="canonical-db5c3356fce7c0a63e648c8852563281828f9ec47f6df5ff84d31fbc5e54df56"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.headers.check_present / 2ee100c42a76 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b3d4418d17c0a561041f1f73ffcf12299edb3b27bac619be32d0fb49e880ee76"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.headers.check_present / 2ee100c42a76 / 4

- [api_rate_limit.server_url_rules.request_matcher.headers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6b2dc62cd9ab58db228a9d79fcadea2eaabfe4d873afb68b96a9bb9fd6666cd2)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-cad11c267a1773352dfc8fef058604349be18d1367e2f1acec5821a0db4c4e2f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c4f32dcb5f8a67e50d9ee669168e4964fa3d333c80e82f45f65a6ac9fb1a677"></a>

## api_rate_limit.server_url_rules.request_matcher.headers.item — api_rate_limit.server_url_rules.request_matcher.headers.item / df129bb44cb9 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-66b75078bf3f57cca6e480dadc7b071b3ab3423f328cf97b20e0eea3e110564b)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-bd87a783e34a923db2aed4a40eacd8f9e02254e0479b8f9940ab67baee99bcc5)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-fd63eb81ac8d0d1e1d4ddeb9a3d4d03a0e0b413d6edc24fa4ae959a3db2d0433)
- [api_rate_limit.server_url_rules.request_matcher.headers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6b2dc62cd9ab58db228a9d79fcadea2eaabfe4d873afb68b96a9bb9fd6666cd2)
- api_rate_limit.server_url_rules.request_matcher.headers.item

<a id="canonical-0ed8999c295c7e2fffd9dad2114addd8afea64031674b8c6f6e467874cbfee4d"></a>

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

<a id="canonical-e6152f8dd808b3e7a3390e61f296395878fcd2723a4ef8ceb32d3554c6826b64"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.headers.item / df129bb44cb9 / 3

<a id="canonical-c8c3376235aae6601027f4467cda5f655a79f46cbc813ca35b81592082ab9e88"></a>

<a id="canonical-4f6fe575dbe5b1734c19d1d40954b58bd5318563a2eeb109dfab11cd75f13870"></a>

## exact_values property — api_rate_limit.server_url_rules.request_matcher.headers.item / df129bb44cb9 / 4

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

<a id="canonical-bbd88549e431df3ed7eb8f1afe1bd020afaf6d1d95b8efb55778bb35a0c9351e"></a>

<a id="canonical-17141a40996e85ff6748badab7bfb3f123bdaa8e18a6c055908a6fa9c11bf83f"></a>

## regex_values property — api_rate_limit.server_url_rules.request_matcher.headers.item / df129bb44cb9 / 5

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

<a id="canonical-ad7a2f61fb1e4e541ed39a4ec0d1302cdeec87cab3600937edca3e0cba6fdb76"></a>

<a id="canonical-459a65d175411fba27f68cac06282f0982e2c3113d413dc76f857dc175e35a2a"></a>

## transformers property — api_rate_limit.server_url_rules.request_matcher.headers.item / df129bb44cb9 / 6

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

<a id="canonical-311fb74c708932b63a8edb0a1647bf823bb60cca9c464552060ef7a600718755"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.headers.item / df129bb44cb9 / 7

- [api_rate_limit.server_url_rules.request_matcher.headers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6b2dc62cd9ab58db228a9d79fcadea2eaabfe4d873afb68b96a9bb9fd6666cd2)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-ef6c803ff6f47ffadbc7494fa2296d086e29a5dea67b11f78c72a13f285daf1e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-22c98515197c9a373e7930c971a458c979ea2f928f45a4998e0466cd4d6d27b5"></a>

## api_rate_limit.server_url_rules.request_matcher.jwt_claims — api_rate_limit.server_url_rules.request_matcher.jwt_claims / eb65dcd5b4f7 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-66b75078bf3f57cca6e480dadc7b071b3ab3423f328cf97b20e0eea3e110564b)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-bd87a783e34a923db2aed4a40eacd8f9e02254e0479b8f9940ab67baee99bcc5)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-fd63eb81ac8d0d1e1d4ddeb9a3d4d03a0e0b413d6edc24fa4ae959a3db2d0433)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims

<a id="canonical-35b9d64ce5b70d026795f76ff242985a5dedda64e0cede64fe6fc4ef97bae14a"></a>

Type: `"list"`. Computed.

List of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings.

Upstream description:

A list of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates
must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled.

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

<a id="canonical-5265646cfccbf67581afef0cde1d3af6b095a0f74d574f5f688241268365c11d"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.jwt_claims / eb65dcd5b4f7 / 3

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3f47cfacbe31b86073fd719f8bb7a3106116d0ab3dec54aa66cd47bccce75530): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-20dda07a7758008b888b747ceb712ee463fa1d538d3485e4ad959c556717c32b): complete subsection reference.

<a id="canonical-67c6ab64f82b7e99ecaf9f38ae2aa39123309bd6e7470219ffd4cda2fbfceaa8"></a>

<a id="canonical-0c899d4ccef6e74265a61fb54286df22ca6c999c0a51b07452247f525158f16d"></a>

## invert_matcher property — api_rate_limit.server_url_rules.request_matcher.jwt_claims / eb65dcd5b4f7 / 4

Type: `"bool"`. Computed.

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

- [item](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-b9baac0b1ed4a9f1a278ac6a3a1747f8f1df593022c08510b0cc0afb114defbf): complete subsection reference.

<a id="canonical-b918fab96dd5700edeca41af380d249113a4e4f8fc9f5a67c4581fedff1a343f"></a>

<a id="canonical-4f4a49a54596ac4b5c007fed4790f12f623f9963ddbd935ae3b87464062fa5d7"></a>

## name property — api_rate_limit.server_url_rules.request_matcher.jwt_claims / eb65dcd5b4f7 / 5

Type: `"string"`. Computed.

JWT Claim Name. JWT claim name.

Upstream description:

JWT claim name.

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

<a id="canonical-a4ab64edad40de656251869c60eccb0078b1c413bc8e9fff050e0320d6a36177"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.jwt_claims / eb65dcd5b4f7 / 6

- [api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3f47cfacbe31b86073fd719f8bb7a3106116d0ab3dec54aa66cd47bccce75530)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-20dda07a7758008b888b747ceb712ee463fa1d538d3485e4ad959c556717c32b)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims.item](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-b9baac0b1ed4a9f1a278ac6a3a1747f8f1df593022c08510b0cc0afb114defbf)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-fd63eb81ac8d0d1e1d4ddeb9a3d4d03a0e0b413d6edc24fa4ae959a3db2d0433)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-3f47cfacbe31b86073fd719f8bb7a3106116d0ab3dec54aa66cd47bccce75530"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5b9bd3fd4c1522ecfae4af90997c990e63e7103af2cb7363e16f5b27242b8a6b"></a>

## api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present — api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present / 2d5f51372615 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-66b75078bf3f57cca6e480dadc7b071b3ab3423f328cf97b20e0eea3e110564b)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-bd87a783e34a923db2aed4a40eacd8f9e02254e0479b8f9940ab67baee99bcc5)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-fd63eb81ac8d0d1e1d4ddeb9a3d4d03a0e0b413d6edc24fa4ae959a3db2d0433)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-ef6c803ff6f47ffadbc7494fa2296d086e29a5dea67b11f78c72a13f285daf1e)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-fd11e1c93fe162c259f5b8f7687450a05df5b78908ba35207986c6429e39ea40"></a>

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

<a id="canonical-1d50edb7b7e44e1af3f2fa5c35adc65fcb2437d5d425b1711a174395bebc79b7"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present / 2d5f51372615 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7b6ee030cb8ae6b3859b3ae90c5ef3139af847ed41ef48faf84bb95f0b600f7b"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present / 2d5f51372615 / 4

- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-ef6c803ff6f47ffadbc7494fa2296d086e29a5dea67b11f78c72a13f285daf1e)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-20dda07a7758008b888b747ceb712ee463fa1d538d3485e4ad959c556717c32b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d8c4c05c39467222f371a58a6960ccec661027b98b12cc843047b3e4382058f7"></a>

## api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present — api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present / cef65da8eabb / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-66b75078bf3f57cca6e480dadc7b071b3ab3423f328cf97b20e0eea3e110564b)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-bd87a783e34a923db2aed4a40eacd8f9e02254e0479b8f9940ab67baee99bcc5)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-fd63eb81ac8d0d1e1d4ddeb9a3d4d03a0e0b413d6edc24fa4ae959a3db2d0433)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-ef6c803ff6f47ffadbc7494fa2296d086e29a5dea67b11f78c72a13f285daf1e)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present

<a id="canonical-a46370073fb6b0a53164f82cd8c1af0b46b7d973db619d98dbab562bf20984d5"></a>

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

<a id="canonical-796b0c4427389b93c2e12423dd06ad45055eb70b9d1766a6f455d7bcc85673b5"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present / cef65da8eabb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bf0fef51f1e822a2b0046647152e79f1d5450dcfa88e5597367314160a237c4d"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present / cef65da8eabb / 4

- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-ef6c803ff6f47ffadbc7494fa2296d086e29a5dea67b11f78c72a13f285daf1e)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-b9baac0b1ed4a9f1a278ac6a3a1747f8f1df593022c08510b0cc0afb114defbf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dfbc4bff16a2019536ed24f9faddfa4bcf55de209c67e15ef2e424f5e06d0f36"></a>

## api_rate_limit.server_url_rules.request_matcher.jwt_claims.item — api_rate_limit.server_url_rules.request_matcher.jwt_claims.item / 40f8eb1657ec / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-66b75078bf3f57cca6e480dadc7b071b3ab3423f328cf97b20e0eea3e110564b)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-bd87a783e34a923db2aed4a40eacd8f9e02254e0479b8f9940ab67baee99bcc5)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-fd63eb81ac8d0d1e1d4ddeb9a3d4d03a0e0b413d6edc24fa4ae959a3db2d0433)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-ef6c803ff6f47ffadbc7494fa2296d086e29a5dea67b11f78c72a13f285daf1e)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.item

<a id="canonical-0d64d2a9143a258f16f7e7b075f7b769976231b2b2339552508461e06c6824ae"></a>

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

<a id="canonical-01c6aa2d908bf4cfe5152327cca637e6cc40b099ee2f69724b3c9e95529cbac1"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.jwt_claims.item / 40f8eb1657ec / 3

<a id="canonical-f2fe148c95ebdc9499e10fde25c7cd8ffeb328ba57519c6f39338a0672a1d7b3"></a>

<a id="canonical-e1a37dd902a23a36363cee6b9bc10a6f4eb0eecbe4372559d0c14106e2ede67c"></a>

## exact_values property — api_rate_limit.server_url_rules.request_matcher.jwt_claims.item / 40f8eb1657ec / 4

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

<a id="canonical-4dee9ae73223daaad11a059b36b6ef992528d6a50dd45bf89eae3232bf48bbad"></a>

<a id="canonical-c9e7aae2d3283676320965dd0f4e28c7c5ebeb7392c305d5a51d5d2ee675c5ae"></a>

## regex_values property — api_rate_limit.server_url_rules.request_matcher.jwt_claims.item / 40f8eb1657ec / 5

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

<a id="canonical-66cf83d43d43f6118bbd798c884ccbce16e315c1e1f6633a3ca0e73dc8b4d1d7"></a>

<a id="canonical-9668731a06ed7c1e6501dcaca2946b78a5ac19b4174b55bb71207b31eeb6ae44"></a>

## transformers property — api_rate_limit.server_url_rules.request_matcher.jwt_claims.item / 40f8eb1657ec / 6

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

<a id="canonical-9804590bb0a93b2230ad3b827d32c52793a7c7090b4c4c932a98f15346f30aed"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.jwt_claims.item / 40f8eb1657ec / 7

- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-ef6c803ff6f47ffadbc7494fa2296d086e29a5dea67b11f78c72a13f285daf1e)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-5b383596ed3f7654770c1d6a7ab92bbe03da3a6f81ed2cdefeb100c256061b84"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac43663fbbe2ad9d0582c66c99beec58bda3be089a7991026250e46e84e30303"></a>

## api_rate_limit.server_url_rules.request_matcher.query_params — api_rate_limit.server_url_rules.request_matcher.query_params / 5c70945af007 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-66b75078bf3f57cca6e480dadc7b071b3ab3423f328cf97b20e0eea3e110564b)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-bd87a783e34a923db2aed4a40eacd8f9e02254e0479b8f9940ab67baee99bcc5)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-fd63eb81ac8d0d1e1d4ddeb9a3d4d03a0e0b413d6edc24fa4ae959a3db2d0433)
- api_rate_limit.server_url_rules.request_matcher.query_params

<a id="canonical-98f87ed901b7e34897e4bdfeb01ceb522ebc550cc8d42cd5f382725dd48fe781"></a>

Type: `"list"`. Computed.

List of predicates for all query parameters that need to be matched. The criteria for matching each
query parameter are described in individual instances of QueryParameterMatcherType. The actual query
parameter values are extracted from the request API as a list of strings for each query..

Upstream description:

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

<a id="canonical-14768431444b82c4cbd77b7c28c6cc1d55f91254eb332a4820577c4527f9bf0c"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.query_params / 5c70945af007 / 3

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-f9b4e7febe14e8e246f1663c1a6e275e192ac0cf3699cae7150b30663429dc3e): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-7568b77caa40f99437006f110eff88400bb1a495fc70cec1cb204e1cf6673f41): complete subsection reference.

<a id="canonical-1b34d39d9bf5a82c103578483bfd8c013e7c655b987a76e08d96078fe511b9d4"></a>

<a id="canonical-bc4a315e6a2e5614c5f7e841e5b49e8f162971b0815eb0b9b4ed9d8ca74ea234"></a>

## invert_matcher property — api_rate_limit.server_url_rules.request_matcher.query_params / 5c70945af007 / 4

Type: `"bool"`. Computed.

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

- [item](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-8e5ac0f8515b5e792d72ea10349cab3618293c904ee3737217fd8e9f356756e7): complete subsection reference.

<a id="canonical-cd295729f0401f009ec198c91eb21c6096fb5578bcd88373d42dbf747f2fb31a"></a>

<a id="canonical-8858a5104580c0c44769f1af34cfd1741266ca5b4c3472555a0c92de190468ab"></a>

## key property — api_rate_limit.server_url_rules.request_matcher.query_params / 5c70945af007 / 5

Type: `"string"`. Computed.

Case-sensitive HTTP query parameter name.

Upstream description:

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

<a id="canonical-5f2d38382f6576697b6b457df73e6354193a6dfd3244ec34e8b4c5fcd4a81f28"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.query_params / 5c70945af007 / 6

- [api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-f9b4e7febe14e8e246f1663c1a6e275e192ac0cf3699cae7150b30663429dc3e)
- [api_rate_limit.server_url_rules.request_matcher.query_params.check_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-7568b77caa40f99437006f110eff88400bb1a495fc70cec1cb204e1cf6673f41)
- [api_rate_limit.server_url_rules.request_matcher.query_params.item](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-8e5ac0f8515b5e792d72ea10349cab3618293c904ee3737217fd8e9f356756e7)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-fd63eb81ac8d0d1e1d4ddeb9a3d4d03a0e0b413d6edc24fa4ae959a3db2d0433)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-f9b4e7febe14e8e246f1663c1a6e275e192ac0cf3699cae7150b30663429dc3e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d3296dd245d3d00063e8dd136756afc9b95fc3e419c7090f7b8ca54c23c1b56"></a>

## api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present — api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present / a640a7ee0898 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-66b75078bf3f57cca6e480dadc7b071b3ab3423f328cf97b20e0eea3e110564b)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-bd87a783e34a923db2aed4a40eacd8f9e02254e0479b8f9940ab67baee99bcc5)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-fd63eb81ac8d0d1e1d4ddeb9a3d4d03a0e0b413d6edc24fa4ae959a3db2d0433)
- [api_rate_limit.server_url_rules.request_matcher.query_params](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-5b383596ed3f7654770c1d6a7ab92bbe03da3a6f81ed2cdefeb100c256061b84)
- api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present

<a id="canonical-b802e8d06b8ca8e276a46fff0b72f1910be415755af0961e29a845afed9004f2"></a>

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

<a id="canonical-f363b4dfd6d5977f4171be371c1673891cb0ba794ecd37bee35761460ecd8e5e"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present / a640a7ee0898 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f81be20b94cba1a82f97553c62b3e9ca9a1be783c14357032ad7d5b34d1ea735"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present / a640a7ee0898 / 4

- [api_rate_limit.server_url_rules.request_matcher.query_params](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-5b383596ed3f7654770c1d6a7ab92bbe03da3a6f81ed2cdefeb100c256061b84)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-7568b77caa40f99437006f110eff88400bb1a495fc70cec1cb204e1cf6673f41"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e475194f5d2d0fefb77486f7c513069273e39ead59e6f1cf82f994535690a531"></a>

## api_rate_limit.server_url_rules.request_matcher.query_params.check_present — api_rate_limit.server_url_rules.request_matcher.query_params.check_present / 17e2d1325657 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-66b75078bf3f57cca6e480dadc7b071b3ab3423f328cf97b20e0eea3e110564b)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-bd87a783e34a923db2aed4a40eacd8f9e02254e0479b8f9940ab67baee99bcc5)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-fd63eb81ac8d0d1e1d4ddeb9a3d4d03a0e0b413d6edc24fa4ae959a3db2d0433)
- [api_rate_limit.server_url_rules.request_matcher.query_params](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-5b383596ed3f7654770c1d6a7ab92bbe03da3a6f81ed2cdefeb100c256061b84)
- api_rate_limit.server_url_rules.request_matcher.query_params.check_present

<a id="canonical-a5029074911de6f68ac81384244df4cd82b54fb30d0c5d8510a04f675d5de051"></a>

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

<a id="canonical-621fe47981d0deef3bb1e4015c92f9111855707700a6750a33e988f47ebe784f"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.query_params.check_present / 17e2d1325657 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-aada19ae92ac800b8764bfa7aebcd71d89a69711da391e76cc50f4a9cd2187ea"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.query_params.check_present / 17e2d1325657 / 4

- [api_rate_limit.server_url_rules.request_matcher.query_params](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-5b383596ed3f7654770c1d6a7ab92bbe03da3a6f81ed2cdefeb100c256061b84)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-8e5ac0f8515b5e792d72ea10349cab3618293c904ee3737217fd8e9f356756e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4923239482086efbc8fa2b606293d687f1a38f4e3620c4bc8277ab824a3665a3"></a>

## api_rate_limit.server_url_rules.request_matcher.query_params.item — api_rate_limit.server_url_rules.request_matcher.query_params.item / c57a2c0eeefc / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-66b75078bf3f57cca6e480dadc7b071b3ab3423f328cf97b20e0eea3e110564b)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-bd87a783e34a923db2aed4a40eacd8f9e02254e0479b8f9940ab67baee99bcc5)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-fd63eb81ac8d0d1e1d4ddeb9a3d4d03a0e0b413d6edc24fa4ae959a3db2d0433)
- [api_rate_limit.server_url_rules.request_matcher.query_params](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-5b383596ed3f7654770c1d6a7ab92bbe03da3a6f81ed2cdefeb100c256061b84)
- api_rate_limit.server_url_rules.request_matcher.query_params.item

<a id="canonical-8e773d7d112e3601e70515cf78a0306825bcd5688146d054734c851234d1deac"></a>

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

<a id="canonical-42ef68087b90de87a470aa62694e4e43aa57d196749dad08e18a860321856c3a"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.query_params.item / c57a2c0eeefc / 3

<a id="canonical-a1816e6b20ab9cd1b558bc531c3cd3b86ffcc54d98799b4d0cf7645fe59df1c7"></a>

<a id="canonical-09ba65bb8c2d0bd236298921f1aad42d9e47b60c47bc56df3eca353ffb564074"></a>

## exact_values property — api_rate_limit.server_url_rules.request_matcher.query_params.item / c57a2c0eeefc / 4

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

<a id="canonical-3567704be8bb37ce632f141d94425e49a56049ec09396747c0507c98f3338751"></a>

<a id="canonical-ce63ee9ca6937c01b7c90990e7c51663dee3066c95e2d7f358a4e317a1afabbf"></a>

## regex_values property — api_rate_limit.server_url_rules.request_matcher.query_params.item / c57a2c0eeefc / 5

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

<a id="canonical-ef3422b8270698a285bf7da17a2374170fe4d4d2e14c8154cb76a09d100c57fc"></a>

<a id="canonical-ab0aa6fe3953ac408c7aaa4a583758785664ab6257f7b2746b57d6e71bdeb9a1"></a>

## transformers property — api_rate_limit.server_url_rules.request_matcher.query_params.item / c57a2c0eeefc / 6

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

<a id="canonical-12633d2d0052d5c0e73570bdcfd40d47e43cc6404dabbaf6d7af32dc314261b8"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.query_params.item / c57a2c0eeefc / 7

- [api_rate_limit.server_url_rules.request_matcher.query_params](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-5b383596ed3f7654770c1d6a7ab92bbe03da3a6f81ed2cdefeb100c256061b84)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-362f6b7defed1ca0331731f6c193bdce3bc10a6cb21aa4e2772f46d6b5078255"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0a139baa9a21b75622b3b256d03d63ee52d54b434f6d698b50fc491c287ab56e"></a>

## api_specification — api_specification / 1c728c26e572 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- api_specification

<a id="canonical-13f9b8af790bd0ce01dd9f0c14ddb666b675c192b25c232fa2f01ce2dd771c1e"></a>

Type: `"single"`. Computed.

\[OneOf: api\_specification, disable\_api\_definition; Default: disable\_api\_definition\] Settings
for API specification (API definition, OpenAPI validation, etc.).

Upstream description:

Settings for API specification (API definition, OpenAPI validation, etc.)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-validation_target_choice": "[\"validation_all_spec_endpoints\",\"validation_custom_list\",\"validation_disabled\"]"
}
```

OneOf alternatives in this subsection:

- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-13f9b8af790bd0ce01dd9f0c14ddb666b675c192b25c232fa2f01ce2dd771c1e)
- [disable_api_definition](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1b6e410e8050a4c9750180748a209fbcf3b07cc1723c1f135fed669c8a85c585)

Select alternatives according to the provider validators above.

<a id="canonical-e9e704c4084257fe30345c3f65dc25bea56ac96522efc0cb7ee05acbc370fbd4"></a>

## Direct properties — api_specification / 1c728c26e572 / 3

- [api_definition](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-299b5dbbd470205039dbab135ccd5f6e9832b0fa0522a6791728c7d838daf7f8): complete subsection reference.

- [validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6c8af97b6ce4b4a57cc78265f926bbb155aa9163a7190e4cbe7c1ddd5b6bb641): complete subsection reference.

- [validation_custom_list](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-c316743be0286db6ee920c53f048f5a633506fdca8a4f5da9863dc739a318e00): complete subsection reference.

- [validation_disabled](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-a0e90daf7dfad76aa10dccb936e805071d0d3c1213d53a2bb88692d598406da4): complete subsection reference.

<a id="canonical-5c3450957ca8771dfed71e943481139c32884adbb7892f340b166199d45bab69"></a>

## Next pages — api_specification / 1c728c26e572 / 4

- [api_specification.api_definition](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-299b5dbbd470205039dbab135ccd5f6e9832b0fa0522a6791728c7d838daf7f8)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6c8af97b6ce4b4a57cc78265f926bbb155aa9163a7190e4cbe7c1ddd5b6bb641)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-c316743be0286db6ee920c53f048f5a633506fdca8a4f5da9863dc739a318e00)
- [api_specification.validation_disabled](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-a0e90daf7dfad76aa10dccb936e805071d0d3c1213d53a2bb88692d598406da4)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-299b5dbbd470205039dbab135ccd5f6e9832b0fa0522a6791728c7d838daf7f8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ccb6be4aa20b1fe0fd1680ea0d28bf00ecd7612ec416b39a1f015b9d81bcb0b"></a>

## api_specification.api_definition — api_specification.api_definition / 58162d1b5a10 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-362f6b7defed1ca0331731f6c193bdce3bc10a6cb21aa4e2772f46d6b5078255)
- api_specification.api_definition

<a id="canonical-a4ca4d01665b781dc6f21ff810d14c6ccfc7ee2ee3f1e20415b642b5abfa6e14"></a>

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

<a id="canonical-e767ca97982d9be0ebe78e0f1dde9811a1b336edd05f38cc0d0e5dee364c9ef8"></a>

## Direct properties — api_specification.api_definition / 58162d1b5a10 / 3

<a id="canonical-bc0828147a2b7e84996d1a52da2139ac4fa4972baf7e56ce048f84c6c64bc1f7"></a>

<a id="canonical-fcc829a691d4bf944117505dd11afe31044a0250989e84878e32f2bbde6c6d90"></a>

## name property — api_specification.api_definition / 58162d1b5a10 / 4

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

<a id="canonical-afa90d1e2998b8820e36066a24558821b700fbc672b4093884c2ca60e4b7098b"></a>

<a id="canonical-6c0523b0c5a54b9443c3971eb84aaf5243d2cfb51055674ff559900fed807fe6"></a>

## namespace property — api_specification.api_definition / 58162d1b5a10 / 5

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

<a id="canonical-0ec30daf31272fdda8f12addc17a5986d7a0c327009547f6be5b925a25f82b4a"></a>

<a id="canonical-faff2d05331cf63e07bbe06523cfddc398a0e566002d59646be0f4f3c29caef3"></a>

## tenant property — api_specification.api_definition / 58162d1b5a10 / 6

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

<a id="canonical-7af48fae35a853194ea2a4a721d686bb73dd11aee12d63606594db9b80ad8799"></a>

## Next pages — api_specification.api_definition / 58162d1b5a10 / 7

- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-362f6b7defed1ca0331731f6c193bdce3bc10a6cb21aa4e2772f46d6b5078255)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-6c8af97b6ce4b4a57cc78265f926bbb155aa9163a7190e4cbe7c1ddd5b6bb641"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e1641600a6978cceee81b4084c7b02a5aaf3ce82e4124f66c84f9d50657048e"></a>

## api_specification.validation_all_spec_endpoints — api_specification.validation_all_spec_endpoints / 8685efed1436 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-362f6b7defed1ca0331731f6c193bdce3bc10a6cb21aa4e2772f46d6b5078255)
- api_specification.validation_all_spec_endpoints

<a id="canonical-c1a048f86dfa26081f09cf8c90655cb4e4d77adbd3ddcbe39e54741f3e6571fa"></a>

Type: `"single"`. Computed.

API Inventory. Settings for API Inventory validation.

Upstream description:

Settings for API Inventory validation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-oversized_body_choice": "[]"
}
```

<a id="canonical-9784859a393da26ffc88edd97d6f55fb32b1a88c74363476ad9140228d69485b"></a>

## Direct properties — api_specification.validation_all_spec_endpoints / 8685efed1436 / 3

- [fall_through_mode](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2b91ecd52d76847b9ec8e8d50130d1012019109a44090eb480e52a89f4d51358): complete subsection reference.

- [settings](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3733ab6e8c1436e9d57cc5c33574237c9b5d49021065b6662aac60393a02af56): complete subsection reference.

- [validation_mode](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-94f66d202c5d9defa92a96078a4d5ee29478ac46b33a660e6328a2a640b4ba69): complete subsection reference.

<a id="canonical-51a433bcd678da57d49694ec4ee794f273550d61c0f15dc313f59264fc94be5a"></a>

## Next pages — api_specification.validation_all_spec_endpoints / 8685efed1436 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2b91ecd52d76847b9ec8e8d50130d1012019109a44090eb480e52a89f4d51358)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3733ab6e8c1436e9d57cc5c33574237c9b5d49021065b6662aac60393a02af56)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-94f66d202c5d9defa92a96078a4d5ee29478ac46b33a660e6328a2a640b4ba69)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-362f6b7defed1ca0331731f6c193bdce3bc10a6cb21aa4e2772f46d6b5078255)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-2b91ecd52d76847b9ec8e8d50130d1012019109a44090eb480e52a89f4d51358"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d41323ee31269556ca1a6395fce97e50c64807a4df15d3bdfa83a160162ee365"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode — api_specification.validation_all_spec_endpoints.fall_through_mode / ffa445cd0c30 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-362f6b7defed1ca0331731f6c193bdce3bc10a6cb21aa4e2772f46d6b5078255)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6c8af97b6ce4b4a57cc78265f926bbb155aa9163a7190e4cbe7c1ddd5b6bb641)
- api_specification.validation_all_spec_endpoints.fall_through_mode

<a id="canonical-7167fa9cfca65b09a59d39008d34171eb187d78bd17ffc6d0842828db06a33f7"></a>

Type: `"single"`. Computed.

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules).

Upstream description:

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-fall_through_mode_choice": "[\"fall_through_mode_allow\",\"fall_through_mode_custom\"]"
}
```

<a id="canonical-062713a2518d0be6726e4b44902b6d459da07eb9c904fb6c4832ed1f8911db8a"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode / ffa445cd0c30 / 3

- [fall_through_mode_allow](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-621613e5d9fdb0c8e1f8e881b9449e86f6555ed1820cafa362893e4d315efbea): complete subsection reference.

- [fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-cdf5bc6304326e0f5acb1f5682c84f1c560939d7f62245869c999dc61cf9e87e): complete subsection reference.

<a id="canonical-8a578efdb574d41108f76f3e614827a266edb346c50c9d0c06f964eabefd3e6b"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode / ffa445cd0c30 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-621613e5d9fdb0c8e1f8e881b9449e86f6555ed1820cafa362893e4d315efbea)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-cdf5bc6304326e0f5acb1f5682c84f1c560939d7f62245869c999dc61cf9e87e)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6c8af97b6ce4b4a57cc78265f926bbb155aa9163a7190e4cbe7c1ddd5b6bb641)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-621613e5d9fdb0c8e1f8e881b9449e86f6555ed1820cafa362893e4d315efbea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f826bee10dbc14c1cd6b13f3854680a0ca26e8e71ab5fb5af6732cef37514e8e"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 833c4bac7119 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-362f6b7defed1ca0331731f6c193bdce3bc10a6cb21aa4e2772f46d6b5078255)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6c8af97b6ce4b4a57cc78265f926bbb155aa9163a7190e4cbe7c1ddd5b6bb641)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2b91ecd52d76847b9ec8e8d50130d1012019109a44090eb480e52a89f4d51358)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow

<a id="canonical-0778a9caaabd390dbc7f5543afe26947ff2377e50f00f03c44306442e9ba055c"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for fall through mode allow.

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

<a id="canonical-e74ac9c15b4a81945bd91858882e294e594f93a283353db8f5c6967f250c35bb"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 833c4bac7119 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6702c1dbd64c774d6597007342d8863e099b26bf472639aadeb81534beb69bb7"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 833c4bac7119 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2b91ecd52d76847b9ec8e8d50130d1012019109a44090eb480e52a89f4d51358)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-cdf5bc6304326e0f5acb1f5682c84f1c560939d7f62245869c999dc61cf9e87e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd5ce5b91ea3e2513a701237a3af6c8cb0028e0cb08e9871170db45b67319bff"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / e6b22477096a / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-362f6b7defed1ca0331731f6c193bdce3bc10a6cb21aa4e2772f46d6b5078255)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6c8af97b6ce4b4a57cc78265f926bbb155aa9163a7190e4cbe7c1ddd5b6bb641)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2b91ecd52d76847b9ec8e8d50130d1012019109a44090eb480e52a89f4d51358)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom

<a id="canonical-7f882d0a0ff3de42e04097f3b02b080165fde981be0d85938014fcb91bd4b9d3"></a>

Type: `"single"`. Computed.

Configuration parameter for fall through mode custom.

Upstream description:

Define the fall through settings.

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

<a id="canonical-924d28fd3fd2ffce7700f6ce82fa0db1e6b01d55b8c4f32edfe36641b36b9e6d"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / e6b22477096a / 3

- [open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-e0093ac24c850241502b4741328cedf261cdb6903e952a24b6f5bd70d061d624): complete subsection reference.

<a id="canonical-fbf12e688561b9a939ac33172a5424ba6222b2256f1c6f9c73ca1d94f32d3d8e"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / e6b22477096a / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-e0093ac24c850241502b4741328cedf261cdb6903e952a24b6f5bd70d061d624)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2b91ecd52d76847b9ec8e8d50130d1012019109a44090eb480e52a89f4d51358)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-e0093ac24c850241502b4741328cedf261cdb6903e952a24b6f5bd70d061d624"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-251b66b67f7d6eb231e145146bc9accf96b0c435b8126bde570b843970125a3c"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / edad1c9adc86 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-362f6b7defed1ca0331731f6c193bdce3bc10a6cb21aa4e2772f46d6b5078255)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6c8af97b6ce4b4a57cc78265f926bbb155aa9163a7190e4cbe7c1ddd5b6bb641)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2b91ecd52d76847b9ec8e8d50130d1012019109a44090eb480e52a89f4d51358)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-cdf5bc6304326e0f5acb1f5682c84f1c560939d7f62245869c999dc61cf9e87e)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules

<a id="canonical-c03b72f8e0cbd084214ee4a5c4d6522a273346e05d3711e926da226bb229816c"></a>

Type: `"list"`. Computed.

Custom Fall Through Rule List. Rule or policy definition

Upstream description:

Rule or policy definition

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 15,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 15,
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
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-8aac6575332a1001583c72087a3fbe89f907c04d253ac3c43dc2daa9921fac43"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / edad1c9adc86 / 3

- [action_block](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-9dac8c7ac34093bc6447ab96977066e0a99c08eb33d54c7dca2bbd9fde4b1fbd): complete subsection reference.

- [action_report](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3530ab3394b2939e3e34083221cc5c6e4b030e8d340911f7ef1c1dd65a0e21d2): complete subsection reference.

- [action_skip](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-e37cf21c816486f8fe1830cca6587552d9c5f0f7b9beea48c6de959f4e1ca897): complete subsection reference.

- [api_endpoint](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-c3b7572f2b066b07265fdd8ec759d56c83746bf3226bf1c067cc64904b9164b1): complete subsection reference.

<a id="canonical-3f481867d9eee3fab92e837830d9e022a5bb34a4f3ff59cfddb2c33d91f352ec"></a>

<a id="canonical-8e4cf2f5cb1edf4fe9526dcf83f7c0111cce24536321ca655a39c4fffbde88fc"></a>

## api_group property — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / edad1c9adc86 / 4

Type: `"string"`. Computed.

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

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

<a id="canonical-4db77fb009b1d34390d8457917ad73f12d2f30ef8457e3adaea920dc08cb6210"></a>

<a id="canonical-233220e51c8a9f725e719dfe2c86fbc5f2cdcd303f4fc21b8a53bfe6820d71cf"></a>

## base_path property — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / edad1c9adc86 / 5

Type: `"string"`. Computed.

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [metadata](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-06a96f5033b756eaa9080c7ddf3b1dcaf3c090f73ee1feae35071113ed7c3ec1): complete subsection reference.

<a id="canonical-6b90d91592c65ba75092a91490ad07bdb64cdb9220243f748f73d5fe77bb21a2"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / edad1c9adc86 / 6

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-9dac8c7ac34093bc6447ab96977066e0a99c08eb33d54c7dca2bbd9fde4b1fbd)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3530ab3394b2939e3e34083221cc5c6e4b030e8d340911f7ef1c1dd65a0e21d2)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-e37cf21c816486f8fe1830cca6587552d9c5f0f7b9beea48c6de959f4e1ca897)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-c3b7572f2b066b07265fdd8ec759d56c83746bf3226bf1c067cc64904b9164b1)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-06a96f5033b756eaa9080c7ddf3b1dcaf3c090f73ee1feae35071113ed7c3ec1)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-cdf5bc6304326e0f5acb1f5682c84f1c560939d7f62245869c999dc61cf9e87e)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-9dac8c7ac34093bc6447ab96977066e0a99c08eb33d54c7dca2bbd9fde4b1fbd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ced334b1dd58a260a3f57573e57a8ddfbc4adf8c9cefc068aab2380e3fda82c9"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 3e33862df3b0 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-362f6b7defed1ca0331731f6c193bdce3bc10a6cb21aa4e2772f46d6b5078255)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6c8af97b6ce4b4a57cc78265f926bbb155aa9163a7190e4cbe7c1ddd5b6bb641)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2b91ecd52d76847b9ec8e8d50130d1012019109a44090eb480e52a89f4d51358)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-cdf5bc6304326e0f5acb1f5682c84f1c560939d7f62245869c999dc61cf9e87e)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-e0093ac24c850241502b4741328cedf261cdb6903e952a24b6f5bd70d061d624)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block

<a id="canonical-45c4f33d0cdf1db778970ce959fab4480082f7f2a4fb397a05c30d9489c04146"></a>

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

<a id="canonical-db8d99e74f25c44fca6b5f29c0dc8fdb88ff62b42c4b31ff8f79dd99e81bb952"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 3e33862df3b0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5b2a3c9402375c4ee396923c3105e841a58d303999be0e74a4c67d0cc87193b2"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 3e33862df3b0 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-e0093ac24c850241502b4741328cedf261cdb6903e952a24b6f5bd70d061d624)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-3530ab3394b2939e3e34083221cc5c6e4b030e8d340911f7ef1c1dd65a0e21d2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-862756ca16121dda03e09293415ef03fc754e39c3eb2438495c73e3fcfe1281c"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 365adef03dae / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-362f6b7defed1ca0331731f6c193bdce3bc10a6cb21aa4e2772f46d6b5078255)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6c8af97b6ce4b4a57cc78265f926bbb155aa9163a7190e4cbe7c1ddd5b6bb641)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2b91ecd52d76847b9ec8e8d50130d1012019109a44090eb480e52a89f4d51358)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-cdf5bc6304326e0f5acb1f5682c84f1c560939d7f62245869c999dc61cf9e87e)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-e0093ac24c850241502b4741328cedf261cdb6903e952a24b6f5bd70d061d624)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report

<a id="canonical-b8fa8c1f8a592476f945f7db92dec3f7237aa1216281068dc3299084d4c3880e"></a>

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

<a id="canonical-f52e18af2fd0e465ba9230d9f3f6c192cfa317352a61db2a5015bfcabc9dffb1"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 365adef03dae / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-89174eba527995158a57cfb3bfbfdb42ddc6975b2fd3f7f8e7e549ecefae256d"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 365adef03dae / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-e0093ac24c850241502b4741328cedf261cdb6903e952a24b6f5bd70d061d624)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-e37cf21c816486f8fe1830cca6587552d9c5f0f7b9beea48c6de959f4e1ca897"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-112252a1ed5fbe00dca2ba4552f7982c76a9e504447a7f5711304d7c979f1646"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 54b2219d8cbb / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-362f6b7defed1ca0331731f6c193bdce3bc10a6cb21aa4e2772f46d6b5078255)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6c8af97b6ce4b4a57cc78265f926bbb155aa9163a7190e4cbe7c1ddd5b6bb641)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2b91ecd52d76847b9ec8e8d50130d1012019109a44090eb480e52a89f4d51358)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-cdf5bc6304326e0f5acb1f5682c84f1c560939d7f62245869c999dc61cf9e87e)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-e0093ac24c850241502b4741328cedf261cdb6903e952a24b6f5bd70d061d624)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip

<a id="canonical-1629170283930bf4ad88c851534be0025e5fc78fedf0f6fb1957c38c2bde522b"></a>

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

<a id="canonical-a3428053f27ae8202274516881664d0ff798c7560971d791f56aa6001a2410ab"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 54b2219d8cbb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dd40927530d2b3240c43461a38da8ba08fcd32a417c8cd89caf5c9c89844611b"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 54b2219d8cbb / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-e0093ac24c850241502b4741328cedf261cdb6903e952a24b6f5bd70d061d624)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-c3b7572f2b066b07265fdd8ec759d56c83746bf3226bf1c067cc64904b9164b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-44767686ef14583b6e59c79355225f8a529d8939c8eecff02033a1855153fce2"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 7f310b82557f / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-362f6b7defed1ca0331731f6c193bdce3bc10a6cb21aa4e2772f46d6b5078255)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6c8af97b6ce4b4a57cc78265f926bbb155aa9163a7190e4cbe7c1ddd5b6bb641)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2b91ecd52d76847b9ec8e8d50130d1012019109a44090eb480e52a89f4d51358)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-cdf5bc6304326e0f5acb1f5682c84f1c560939d7f62245869c999dc61cf9e87e)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-e0093ac24c850241502b4741328cedf261cdb6903e952a24b6f5bd70d061d624)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint

<a id="canonical-ad11f674678bef28dddc61ba92f1d9eecf4ce1828960028ae868d07144ec4302"></a>

Type: `"single"`. Computed.

API Endpoint. This defines API endpoint.

Upstream description:

This defines API endpoint.

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

<a id="canonical-2ca5f76c51d0f450b602f6f363961ad3a84f417ebbe6ec22507e95642f0f62bf"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 7f310b82557f / 3

<a id="canonical-38cf315feb2b915e2832fd13fb432807c5362aab7c48aa584dca889ac51f8226"></a>

<a id="canonical-fa5d03e1d222e52c14ea352120b77cc300d10d5a87af2f3605aa71a48024d8de"></a>

## methods property — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 7f310b82557f / 4

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

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

<a id="canonical-101b4ed48f38ff899484a24481e0af2ae6acb307376a1938b9f6200d30a0d40e"></a>

<a id="canonical-f05dd34043b3476302d911a794cf5be86a406c79c789f105092ea270f0573438"></a>

## path property — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 7f310b82557f / 5

Type: `"string"`. Computed.

Path. Path to be matched.

Upstream description:

Path to be matched.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

<a id="canonical-192a53c6496c7435ce8bb57567fa8d64edc7f960a68dfc40c3c77e1e3ad01593"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 7f310b82557f / 6

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-e0093ac24c850241502b4741328cedf261cdb6903e952a24b6f5bd70d061d624)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-06a96f5033b756eaa9080c7ddf3b1dcaf3c090f73ee1feae35071113ed7c3ec1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2b0cf83ba8e47b138e1f15a64372652fb9f299a8ddcd3b2c2ae4cd19fc677d9"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 9427fa4d6b1f / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-362f6b7defed1ca0331731f6c193bdce3bc10a6cb21aa4e2772f46d6b5078255)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6c8af97b6ce4b4a57cc78265f926bbb155aa9163a7190e4cbe7c1ddd5b6bb641)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2b91ecd52d76847b9ec8e8d50130d1012019109a44090eb480e52a89f4d51358)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-cdf5bc6304326e0f5acb1f5682c84f1c560939d7f62245869c999dc61cf9e87e)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-e0093ac24c850241502b4741328cedf261cdb6903e952a24b6f5bd70d061d624)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata

<a id="canonical-64610fb552eac4433bf85039e540198ba22ebe4f87fdb56d68cc42237795f795"></a>

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

<a id="canonical-ceac66efdbb28feb978012e2216653e82779616d4806d77680a42eb53bdb7dbc"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 9427fa4d6b1f / 3

<a id="canonical-bb3fbc1de59caffcd8dcd4c46b285da5bed1beba6837f51e3234609dd2d353a5"></a>

<a id="canonical-486460051d91200ebfdda1d09f55cb61f097343deaf33f3256651d2fe75bd214"></a>

## description_spec property — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 9427fa4d6b1f / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-c655407438f9ac195b0dd000cbd862268516feebb1369c9affe5d05653169c11"></a>

<a id="canonical-a9066b5d237790f747b6977b69cd49292aa7ff9e344a5fd1ec6fbefd3b0da371"></a>

## name property — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 9427fa4d6b1f / 5

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

<a id="canonical-eb7bb3453671e88f5dc1497bce8a4ad46a966b3c3151119697a7889da47ad6b1"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 9427fa4d6b1f / 6

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-e0093ac24c850241502b4741328cedf261cdb6903e952a24b6f5bd70d061d624)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-3733ab6e8c1436e9d57cc5c33574237c9b5d49021065b6662aac60393a02af56"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c114542b785467a8f8ed59407d6eea4806b6b14e9103b8ebdaa32a9d7ac57206"></a>

## api_specification.validation_all_spec_endpoints.settings — api_specification.validation_all_spec_endpoints.settings / ddfbd573bd3b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-362f6b7defed1ca0331731f6c193bdce3bc10a6cb21aa4e2772f46d6b5078255)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6c8af97b6ce4b4a57cc78265f926bbb155aa9163a7190e4cbe7c1ddd5b6bb641)
- api_specification.validation_all_spec_endpoints.settings

<a id="canonical-9eb58bbbd2fa7c133e17bd771d8567448e6fff2195f11a6aed72b84294bb6a17"></a>

Type: `"single"`. Computed.

OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom
list' enforcement.

Upstream description:

OpenAPI specification validation settings relevant for "API Inventory" enforcement and for "Custom
list" enforcement.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-fail_configuration": "[]",
  "x-ves-oneof-field-oversized_body_choice": "[\"oversized_body_fail_validation\",\"oversized_body_skip_validation\"]",
  "x-ves-oneof-field-property_validation_settings_choice": "[\"property_validation_settings_custom\",\"property_validation_settings_default\"]"
}
```

<a id="canonical-e9af0cf34095e22444eababf8963d236c983ed98161cff8675c73e5dbecd048a"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings / ddfbd573bd3b / 3

- [oversized_body_fail_validation](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-8e6e78f0c6e518edcde867bcb65accd14d314f8cdd7b62bb8a08f6648e5fd2ac): complete subsection reference.

- [oversized_body_skip_validation](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-7a6eded63210454ceab9f42c32795454ba4f5def868791fd092963280c1fae68): complete subsection reference.

- [property_validation_settings_custom](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-daadd821532ad2cee8a90a8c8e32f2ef6fc24f588c9a50dc839d819537d7c1d5): complete subsection reference.

- [property_validation_settings_default](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6ded16c16209ffefa7989207e87f292b051f18ac42d1d3d18cd1553155960d9b): complete subsection reference.

<a id="canonical-7c5244c069db58d71d1002ecda070227c2040ef3ab4ed648d1a483cda84bd108"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings / ddfbd573bd3b / 4

- [api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-8e6e78f0c6e518edcde867bcb65accd14d314f8cdd7b62bb8a08f6648e5fd2ac)
- [api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-7a6eded63210454ceab9f42c32795454ba4f5def868791fd092963280c1fae68)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-daadd821532ad2cee8a90a8c8e32f2ef6fc24f588c9a50dc839d819537d7c1d5)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_default](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6ded16c16209ffefa7989207e87f292b051f18ac42d1d3d18cd1553155960d9b)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6c8af97b6ce4b4a57cc78265f926bbb155aa9163a7190e4cbe7c1ddd5b6bb641)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-8e6e78f0c6e518edcde867bcb65accd14d314f8cdd7b62bb8a08f6648e5fd2ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2740cc63b7d9735453968711337bbe2d4836f6fe428b02f8f4879dfa94a6c2cd"></a>

## api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation — api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_val / 10833865bddf / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-362f6b7defed1ca0331731f6c193bdce3bc10a6cb21aa4e2772f46d6b5078255)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6c8af97b6ce4b4a57cc78265f926bbb155aa9163a7190e4cbe7c1ddd5b6bb641)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3733ab6e8c1436e9d57cc5c33574237c9b5d49021065b6662aac60393a02af56)
- api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation

<a id="canonical-5a53c4f39d77dba350689bd9b97794da879c72a8e5b61e5cb2054fece4f28a2b"></a>

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

<a id="canonical-ad313103ae33f72bb95539a23593501251bdf6ce0d5f6da4942ec14d4f373c6b"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_val / 10833865bddf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a7562ecee0fc8f9e99096d34c875206a1c2ef2598a55bd3a8a6eb2fb0bb93ff3"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_val / 10833865bddf / 4

- [api_specification.validation_all_spec_endpoints.settings](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3733ab6e8c1436e9d57cc5c33574237c9b5d49021065b6662aac60393a02af56)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-7a6eded63210454ceab9f42c32795454ba4f5def868791fd092963280c1fae68"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-725b5555bcfe9d1d8347431a5afa92dfca72812f56a169eeee59932254f630c7"></a>

## api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation — api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_val / e4a5bffd4d32 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-362f6b7defed1ca0331731f6c193bdce3bc10a6cb21aa4e2772f46d6b5078255)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6c8af97b6ce4b4a57cc78265f926bbb155aa9163a7190e4cbe7c1ddd5b6bb641)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3733ab6e8c1436e9d57cc5c33574237c9b5d49021065b6662aac60393a02af56)
- api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation

<a id="canonical-9cc454de7fb662d798988bf316498fc394a5ad1f112ed37c39cb8267f9162c02"></a>

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

<a id="canonical-a10565914d805c53253a465eea65b7a005532e2943a46390d04793c0da9df9b9"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_val / e4a5bffd4d32 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1a097a39345dade04788db3417f6b4a320c3368051a6139a3537bc4c99a48827"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_val / e4a5bffd4d32 / 4

- [api_specification.validation_all_spec_endpoints.settings](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3733ab6e8c1436e9d57cc5c33574237c9b5d49021065b6662aac60393a02af56)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-daadd821532ad2cee8a90a8c8e32f2ef6fc24f588c9a50dc839d819537d7c1d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7f7a4ea71254bf993c84e23140862fa9cb7c52a987a389055f2ab3a3d81d4578"></a>

## api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom — api_specification.validation_all_spec_endpoints.settings.property_validation_set / f95d9cdc7079 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-362f6b7defed1ca0331731f6c193bdce3bc10a6cb21aa4e2772f46d6b5078255)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6c8af97b6ce4b4a57cc78265f926bbb155aa9163a7190e4cbe7c1ddd5b6bb641)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3733ab6e8c1436e9d57cc5c33574237c9b5d49021065b6662aac60393a02af56)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom

<a id="canonical-a3dc35b0c6ee2eb14bdc9bb1598a31623a419359d008cd09b472e7698329dbdf"></a>

Type: `"single"`. Computed.

Configuration parameter for property validation settings custom.

Upstream description:

Custom property validation settings.

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

<a id="canonical-1d09f3bcb73044babf8a66f5c3b174b41d90aca3b6158bca9c7d5e4ea0c9cd16"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings.property_validation_set / f95d9cdc7079 / 3

- [query_parameters](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-06a27169441821e0fd524730ede635841733079d7b9e0e69aa9326276961b431): complete subsection reference.

<a id="canonical-bf1e031b586271a396db258589206e288bf42e293b57b1a8a9da2bbbd88920a6"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings.property_validation_set / f95d9cdc7079 / 4

- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-06a27169441821e0fd524730ede635841733079d7b9e0e69aa9326276961b431)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3733ab6e8c1436e9d57cc5c33574237c9b5d49021065b6662aac60393a02af56)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-06a27169441821e0fd524730ede635841733079d7b9e0e69aa9326276961b431"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dfb800ea96c70a9484547c570c0d3f542f9dd5d9b0f91fb09872df94ea1b0327"></a>

## api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters — api_specification.validation_all_spec_endpoints.settings.property_validation_set / f75782f1d60c / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-362f6b7defed1ca0331731f6c193bdce3bc10a6cb21aa4e2772f46d6b5078255)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6c8af97b6ce4b4a57cc78265f926bbb155aa9163a7190e4cbe7c1ddd5b6bb641)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3733ab6e8c1436e9d57cc5c33574237c9b5d49021065b6662aac60393a02af56)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-daadd821532ad2cee8a90a8c8e32f2ef6fc24f588c9a50dc839d819537d7c1d5)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters

<a id="canonical-41bc77d78c3dc8da9f1d1163fa23fd06a537e1c690b7aff80b7c4b007eab38b9"></a>

Type: `"single"`. Computed.

Custom settings for query parameters validation.

<a id="canonical-6fcaa643815ffafc6461ac303846b208f7527a0aab607f35371837216b7c9a1a"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings.property_validation_set / f75782f1d60c / 3

- [allow_additional_parameters](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-5a599cb7e6437ebfaa71e84136f0c35587981aa26c0d6faf4d42aa8d50637ec4): complete subsection reference.

- [disallow_additional_parameters](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-4bcc1581d6ec2249df7b3b16e0900c75e7997f00b3f909683f68d9417919101d): complete subsection reference.

<a id="canonical-bbe8a5d2af06186111865ea0ec788426bde5beb454db4c289c6c73fb14c425e5"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings.property_validation_set / f75782f1d60c / 4

- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-5a599cb7e6437ebfaa71e84136f0c35587981aa26c0d6faf4d42aa8d50637ec4)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-4bcc1581d6ec2249df7b3b16e0900c75e7997f00b3f909683f68d9417919101d)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-daadd821532ad2cee8a90a8c8e32f2ef6fc24f588c9a50dc839d819537d7c1d5)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-5a599cb7e6437ebfaa71e84136f0c35587981aa26c0d6faf4d42aa8d50637ec4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-09edd5d6f5c14f0f0493c7d4adc53d38c59ad75775bd5e6b07d8454b44d4c91e"></a>

## api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters — api_specification.validation_all_spec_endpoints.settings.property_validation_set / 44ff66a7b2ba / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-362f6b7defed1ca0331731f6c193bdce3bc10a6cb21aa4e2772f46d6b5078255)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6c8af97b6ce4b4a57cc78265f926bbb155aa9163a7190e4cbe7c1ddd5b6bb641)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3733ab6e8c1436e9d57cc5c33574237c9b5d49021065b6662aac60393a02af56)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-daadd821532ad2cee8a90a8c8e32f2ef6fc24f588c9a50dc839d819537d7c1d5)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-06a27169441821e0fd524730ede635841733079d7b9e0e69aa9326276961b431)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters

<a id="canonical-e95fdc7b8f3b5f2183d5ffd8486b413519482fea0ecefa4e546eee7b87e51302"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for allow additional parameters.

<a id="canonical-508e96b3337adac63060fcb68d26af7c091bdf19c42acf302f4ec08c8d91312c"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings.property_validation_set / 44ff66a7b2ba / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-15bc955660f1e5d4202239791c92318212d7787da76b38051c78047b89824bfa"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings.property_validation_set / 44ff66a7b2ba / 4

- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-06a27169441821e0fd524730ede635841733079d7b9e0e69aa9326276961b431)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-4bcc1581d6ec2249df7b3b16e0900c75e7997f00b3f909683f68d9417919101d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a1fe570bbe709893d809044a64a91e22265fb5baa107e64a38c96699c12cfa1f"></a>

## api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters — api_specification.validation_all_spec_endpoints.settings.property_validation_set / 8e1b33eb57ee / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-362f6b7defed1ca0331731f6c193bdce3bc10a6cb21aa4e2772f46d6b5078255)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6c8af97b6ce4b4a57cc78265f926bbb155aa9163a7190e4cbe7c1ddd5b6bb641)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3733ab6e8c1436e9d57cc5c33574237c9b5d49021065b6662aac60393a02af56)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-daadd821532ad2cee8a90a8c8e32f2ef6fc24f588c9a50dc839d819537d7c1d5)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-06a27169441821e0fd524730ede635841733079d7b9e0e69aa9326276961b431)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters

<a id="canonical-c71a2a7002da235707baed3dd4c4401fc11f3bf4ac9acbe6956a45faf6a9e5a3"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disallow additional parameters.

<a id="canonical-d0e47368b8000d05dccdbdea29992c3b493457f7fb6d2e556e9d9a06480cf429"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings.property_validation_set / 8e1b33eb57ee / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-97ca5f90857f239e0a95593913a6b00742737300924c445baafc5f37c6210b3b"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings.property_validation_set / 8e1b33eb57ee / 4

- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-06a27169441821e0fd524730ede635841733079d7b9e0e69aa9326276961b431)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-6ded16c16209ffefa7989207e87f292b051f18ac42d1d3d18cd1553155960d9b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ec5509f86274f07a3982da2827be6672fc49defc2abb6f594ad41de3bf8b95a1"></a>

## api_specification.validation_all_spec_endpoints.settings.property_validation_settings_default — api_specification.validation_all_spec_endpoints.settings.property_validation_set / 134aa338fb59 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-362f6b7defed1ca0331731f6c193bdce3bc10a6cb21aa4e2772f46d6b5078255)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-6c8af97b6ce4b4a57cc78265f926bbb155aa9163a7190e4cbe7c1ddd5b6bb641)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3733ab6e8c1436e9d57cc5c33574237c9b5d49021065b6662aac60393a02af56)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_default

<a id="canonical-011f5f5a02cc715a88cca7289e118153eee4b9ecc94501a91bbca1f71132acdb"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for property validation settings default.

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

<a id="canonical-07834fd8d7092ae6fb283ee53fcd0d04d683c43077c11f4683784b299155c342"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings.property_validation_set / 134aa338fb59 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-da102cbc77a99ea2d93b0e6dae3a038bd95446f87870e85ea7d0c7a3c8c68133"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings.property_validation_set / 134aa338fb59 / 4

- [api_specification.validation_all_spec_endpoints.settings](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3733ab6e8c1436e9d57cc5c33574237c9b5d49021065b6662aac60393a02af56)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-94f66d202c5d9defa92a96078a4d5ee29478ac46b33a660e6328a2a640b4ba69"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
