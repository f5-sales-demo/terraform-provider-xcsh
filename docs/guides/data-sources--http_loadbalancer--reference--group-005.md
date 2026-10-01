---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-e547ba3719cdd34e7170b971a39407047ceaa39f90edd13ae560f61349b3af6d"></a>

## namespace property — api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets / 1b5b3b37c8b7 / 6

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

<a id="canonical-e3b66b0702331f097bec3502ce7620ef589dd06b327bad4271d1d603b468e6f0"></a>

<a id="canonical-5dec90ced911f763a70a64e618ad33e3a897c3ee512cddf69b49036869e9ddb6"></a>

## tenant property — api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets / 1b5b3b37c8b7 / 7

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

<a id="canonical-fcc83b307b6bc1106776aa06d25504e67dc45b87f543486f92fcfa8a9db56d25"></a>

<a id="canonical-a3d5f9a2a165aae44fab41e38c84ce037c336fe288c8d0e27d3820a992ccc0b5"></a>

## uid property — api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets / 1b5b3b37c8b7 / 8

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

<a id="canonical-acb7245a6b085f20917531acdf6568cdef15bda5ea95aabbbae85b3f2ef16ab6"></a>

## Next pages — api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets / 1b5b3b37c8b7 / 9

- [api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher](data-sources--http_loadbalancer--reference--group-004.md#canonical-d355f416e9945ed1563720dbc1c6b96bbe8f9ea0b980d3407d1e1e52bd928b7d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-45fd06ac56edaa3b1971364fbb740239624c63723641ea68f37d551151a1ab53"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6526af5e494aa67917e7ea0f10e23eb71903dff5959e176813516e2af33df505"></a>

## api_protection_rules.api_endpoint_rules.client_matcher.client_selector — api_protection_rules.api_endpoint_rules.client_matcher.client_selector / fcfeba8083b6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-c96eba43d7d0081fcdd77cdf4790ebd9c61f34b5fb8a78312b73a803fce74c50)
- [api_protection_rules.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-004.md#canonical-45b8660c5d13812289122ba7fb092edb1f3107c8291a7452c4a60000ed9e10b7)
- api_protection_rules.api_endpoint_rules.client_matcher.client_selector

<a id="canonical-49e01948ce499d8c3271ccb846ba4e693aca71783312408165806ea6e3df0692"></a>

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

<a id="canonical-82dc1d506253bea6436af5bf23682d7c2e4faff39e73b65b44a22e5840bc8ccb"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.client_matcher.client_selector / fcfeba8083b6 / 3

<a id="canonical-573b4b54b9ffe13f7d0b9c24791732d7d6f20ba3dd7f7666f051da4d0f803ed4"></a>

<a id="canonical-59be1830fb6684309fb9c3ba42f59cab307e5564f0125f0fdb638d5a7b4e41e4"></a>

## expressions property — api_protection_rules.api_endpoint_rules.client_matcher.client_selector / fcfeba8083b6 / 4

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

<a id="canonical-115f1105e91bd88c435dc3153df8a09eef1ec000d16625192f9143e7d4e61f4a"></a>

## Next pages — api_protection_rules.api_endpoint_rules.client_matcher.client_selector / fcfeba8083b6 / 5

- [api_protection_rules.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-004.md#canonical-45b8660c5d13812289122ba7fb092edb1f3107c8291a7452c4a60000ed9e10b7)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-018d3521236feea6419780f2d6d07d7c412d3f607a4c1825c4ba94f5a0bd9365"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5d46dd60c605db405d855b296cd532dad2acacea5c1f4c3920c1b1b4783c141"></a>

## api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher — api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher / 01962a9d92db / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-c96eba43d7d0081fcdd77cdf4790ebd9c61f34b5fb8a78312b73a803fce74c50)
- [api_protection_rules.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-004.md#canonical-45b8660c5d13812289122ba7fb092edb1f3107c8291a7452c4a60000ed9e10b7)
- api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher

<a id="canonical-aee7f0679b50d35e3ff4cc58518dd9d1cb8cc023bdfb9b762b2a6d57660d94ab"></a>

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

<a id="canonical-f14f2ce20f00c624dce90c3bf7a593e4884575d57bddbe17fa9d9c084437f3ad"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher / 01962a9d92db / 3

<a id="canonical-05c8c40c75fb8aeded8b6b83c8a50a919176f532ecfc3152153544aec93f9419"></a>

<a id="canonical-f5cf20721557dc7c8a083d18d878de237631dba62b917969dec5db86f51e3e0c"></a>

## invert_matcher property — api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher / 01962a9d92db / 4

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

- [prefix_sets](data-sources--http_loadbalancer--reference--group-005.md#canonical-de91f6565f16dabd6abf28321e1c29c5d35a57107a65339b4c7f7e66f5c75eba): complete subsection reference.

<a id="canonical-c0201bcd5f8adabcad481ab158dc132af4610bedd3e2b4b4009e40296895d987"></a>

## Next pages — api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher / 01962a9d92db / 5

- [api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets](data-sources--http_loadbalancer--reference--group-005.md#canonical-de91f6565f16dabd6abf28321e1c29c5d35a57107a65339b4c7f7e66f5c75eba)
- [api_protection_rules.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-004.md#canonical-45b8660c5d13812289122ba7fb092edb1f3107c8291a7452c4a60000ed9e10b7)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-de91f6565f16dabd6abf28321e1c29c5d35a57107a65339b4c7f7e66f5c75eba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac26e7abfad58720143b705ba2860b43d08f04da0f85f15b287c62216cd34dec"></a>

## api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets — api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / de201042afcc / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-c96eba43d7d0081fcdd77cdf4790ebd9c61f34b5fb8a78312b73a803fce74c50)
- [api_protection_rules.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-004.md#canonical-45b8660c5d13812289122ba7fb092edb1f3107c8291a7452c4a60000ed9e10b7)
- [api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-018d3521236feea6419780f2d6d07d7c412d3f607a4c1825c4ba94f5a0bd9365)
- api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-7d71f6107c171913052d21317edd4a8941925bf57e2867084f970dfb4930c466"></a>

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

<a id="canonical-95d6c87c331602331d505e9890ab9ef442caa85a2fc0647939d371deb3c7d159"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / de201042afcc / 3

<a id="canonical-d364a3641a072d0c7454cf0b817013c04ee9543d526d8877c0097c10b7f7d8b2"></a>

<a id="canonical-fa3731b44630686b5a094bf8d3612bf2cbf220fd15eb7278182786e17ad13cd0"></a>

## kind property — api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / de201042afcc / 4

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

<a id="canonical-c698b89e5f427b6f51f4262344ffe7e36c575f13e0c50627fb0acf2ee1ee0e02"></a>

<a id="canonical-097bd6a10a2e5de10d1837f0948ed91cb0067edc5472f6af8728fe0018a98e16"></a>

## name property — api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / de201042afcc / 5

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

<a id="canonical-5402fd5ba16e7699c89c838b6efaf3f583c8958e68550f63d1753043ea90decc"></a>

<a id="canonical-9792adb8bc2dcbf6d787db79bcb0e8822be63bb312218ef1e989a2df02a31d6c"></a>

## namespace property — api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / de201042afcc / 6

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

<a id="canonical-019c5be3c2eab63b61bc7315c739cc2e4e49412e4cfa9e79ca012fd61df1c2d8"></a>

<a id="canonical-65da1227322a9392ffb8667696442fad72222136563cd834710942d01fdc253a"></a>

## tenant property — api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / de201042afcc / 7

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

<a id="canonical-a2afeb3fe00f73e5edf0822b9c9e4de6e3cb7c2dad903e3d53adb9b6f2aaa8fe"></a>

<a id="canonical-76a61e299808be069aef6b805309a16851997982b38b8518454fdbd8fe146630"></a>

## uid property — api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / de201042afcc / 8

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

<a id="canonical-942209140f97b6436597157e6d7b8e25f4cda331290bd2f381910bc9daeb1bd1"></a>

## Next pages — api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / de201042afcc / 9

- [api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-018d3521236feea6419780f2d6d07d7c412d3f607a4c1825c4ba94f5a0bd9365)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-0916edc27c8256e6a92511bc26ed832a85e30c8db92f199ef621db292ef85a0e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e45992e7174adb660f490134f1b59b68bb8e71dfc62ef4bca45bd29956558d2b"></a>

## api_protection_rules.api_endpoint_rules.client_matcher.ip_prefix_list — api_protection_rules.api_endpoint_rules.client_matcher.ip_prefix_list / 166e50524d34 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-c96eba43d7d0081fcdd77cdf4790ebd9c61f34b5fb8a78312b73a803fce74c50)
- [api_protection_rules.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-004.md#canonical-45b8660c5d13812289122ba7fb092edb1f3107c8291a7452c4a60000ed9e10b7)
- api_protection_rules.api_endpoint_rules.client_matcher.ip_prefix_list

<a id="canonical-a96487d1a3fc1046e4d67bef49f705149b0b7a7245e5fe6d180ad4154f5b88e6"></a>

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

<a id="canonical-55a345d1082957e842b92a1ef7070cec1d9abd682865f639d721680c7baf1b7f"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.client_matcher.ip_prefix_list / 166e50524d34 / 3

<a id="canonical-b20daeb42771276a77beba77676d7194fcc0aa3c5dc5dda50e31b11ac097bf5d"></a>

<a id="canonical-0dfe6d4930aa714eadd9cc38fa2981dac06f43cd82abba33adcdcc6da34337d3"></a>

## invert_match property — api_protection_rules.api_endpoint_rules.client_matcher.ip_prefix_list / 166e50524d34 / 4

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

<a id="canonical-3b09f5442c0d11e64a4adfc060346e1792f3734a18cd6736cd01422fb986d805"></a>

<a id="canonical-9e4e4080ba59a11b278544ca4e688a536ce12c6fa5b43f75b3bfa7e1637f9351"></a>

## ip_prefixes property — api_protection_rules.api_endpoint_rules.client_matcher.ip_prefix_list / 166e50524d34 / 5

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

<a id="canonical-5868df7a591f218013c7fb379a272fe013cc09f4c683243516fb4f81ba6b1c52"></a>

## Next pages — api_protection_rules.api_endpoint_rules.client_matcher.ip_prefix_list / 166e50524d34 / 6

- [api_protection_rules.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-004.md#canonical-45b8660c5d13812289122ba7fb092edb1f3107c8291a7452c4a60000ed9e10b7)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-357e7d053d774e85e7d98b29c1624483471a651e196d207fe4ac59ace7f1d8a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-35934892c54ceb4e70736355de2d6aff735ce47e5c0ae90227604404185ac1ef"></a>

## api_protection_rules.api_endpoint_rules.client_matcher.ip_threat_category_list — api_protection_rules.api_endpoint_rules.client_matcher.ip_threat_category_list / 173f70bbde47 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-c96eba43d7d0081fcdd77cdf4790ebd9c61f34b5fb8a78312b73a803fce74c50)
- [api_protection_rules.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-004.md#canonical-45b8660c5d13812289122ba7fb092edb1f3107c8291a7452c4a60000ed9e10b7)
- api_protection_rules.api_endpoint_rules.client_matcher.ip_threat_category_list

<a id="canonical-76299a970f93fd568659413c6b9a5f2e542749a30f8f36a14876dc4b1922184c"></a>

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

<a id="canonical-da469b72b525f96f438181e71f7530c9486ecbce197d406696c1566ff1e79783"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.client_matcher.ip_threat_category_list / 173f70bbde47 / 3

<a id="canonical-ed41ee804b9ede2a82c538dd57a1772a4225b04b0a980b10e0cd94f1727cfa4d"></a>

<a id="canonical-b4e7b9f9ee857cc4e6ed1cf683aca4799df87d6ae9a82a5b3c9f16f3d9054bb5"></a>

## ip_threat_categories property — api_protection_rules.api_endpoint_rules.client_matcher.ip_threat_category_list / 173f70bbde47 / 4

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

<a id="canonical-f43c09740df1285c3609b5deb6ebae0bbbd8a1452329ad69d9b8a4288f484ec5"></a>

## Next pages — api_protection_rules.api_endpoint_rules.client_matcher.ip_threat_category_list / 173f70bbde47 / 5

- [api_protection_rules.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-004.md#canonical-45b8660c5d13812289122ba7fb092edb1f3107c8291a7452c4a60000ed9e10b7)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d13744684bfac6ee40dcb02fe8e0426d1cfdc72193b2a1fbede582258d843941"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-97ecfbac3d69dc11859ba5cb91fcb0b54957708c3d4c435cfcdeb785b71f763c"></a>

## api_protection_rules.api_endpoint_rules.client_matcher.tls_fingerprint_matcher — api_protection_rules.api_endpoint_rules.client_matcher.tls_fingerprint_matcher / 328669a98803 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-c96eba43d7d0081fcdd77cdf4790ebd9c61f34b5fb8a78312b73a803fce74c50)
- [api_protection_rules.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-004.md#canonical-45b8660c5d13812289122ba7fb092edb1f3107c8291a7452c4a60000ed9e10b7)
- api_protection_rules.api_endpoint_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-15b72225262a5189af71368d0b9e3e1e73a4a8d319b6e65e61f8f618420556e6"></a>

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

<a id="canonical-4b7ea4d3916f3aec1d45197b3c684311d087656aac2d018f9e04ab1013afde47"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.client_matcher.tls_fingerprint_matcher / 328669a98803 / 3

<a id="canonical-cb180b4a5449437a8a647a4b1fd5b2244850a60d7dea8822610d15fa7433ff66"></a>

<a id="canonical-6e95e8959955fa15a943606b9ebad2da3875e9f4a1197bdc7f74e24fb2de037c"></a>

## classes property — api_protection_rules.api_endpoint_rules.client_matcher.tls_fingerprint_matcher / 328669a98803 / 4

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

<a id="canonical-44a88a861c973e7be693a5fbf316fdc4cadd88c1451262b6ed0f4f12f114bf23"></a>

<a id="canonical-dd128d674c8df96c1a679d31a2b0642203ff9b142ee5f07c7cd3348032587780"></a>

## exact_values property — api_protection_rules.api_endpoint_rules.client_matcher.tls_fingerprint_matcher / 328669a98803 / 5

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

<a id="canonical-b920f2e8866fb1e89dacb0a7952db1e84042ba33a22f16090554d62059a826ff"></a>

<a id="canonical-00cb60f9555a86b693fc55126180530594e45f76378a60deadf11704d1012ca1"></a>

## excluded_values property — api_protection_rules.api_endpoint_rules.client_matcher.tls_fingerprint_matcher / 328669a98803 / 6

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

<a id="canonical-00d34586ea82a58a9d98db60d4ab9b0946296b40aa633f7755388669aff31eaa"></a>

## Next pages — api_protection_rules.api_endpoint_rules.client_matcher.tls_fingerprint_matcher / 328669a98803 / 7

- [api_protection_rules.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-004.md#canonical-45b8660c5d13812289122ba7fb092edb1f3107c8291a7452c4a60000ed9e10b7)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-8edc44035fd6847a096aeeda5f3c2ce86e3ca4b2b2427045a9cc1f4a2ea23237"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c0efcbe3446506084967dfdcb3b492a19d3cef132c0f1ecde8419cd9b8992d9"></a>

## api_protection_rules.api_endpoint_rules.metadata — api_protection_rules.api_endpoint_rules.metadata / 46391aa72d7a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-c96eba43d7d0081fcdd77cdf4790ebd9c61f34b5fb8a78312b73a803fce74c50)
- api_protection_rules.api_endpoint_rules.metadata

<a id="canonical-7c7088678194939645eae63e798cb7485224b072634d136f7d4d02711193c926"></a>

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

<a id="canonical-2211fe57542ee1d36993eef2a44d6acee6a79309d4191f740a06ca7e8139ab3f"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.metadata / 46391aa72d7a / 3

<a id="canonical-cf9495a68a942cedf799667b93a84c7e2148d003e133d0f1094c36a496c760bb"></a>

<a id="canonical-33b6d8fcacf8f1291cda964f10131ecddaa7f1a699513e985d14e4f8e70668df"></a>

## description_spec property — api_protection_rules.api_endpoint_rules.metadata / 46391aa72d7a / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-6ee29126155229d228b716f53183c10053beba52086953587c11f1b7b83f101d"></a>

<a id="canonical-cd421d58ae02c8c682ab110bcd1f1f272108332d015f5ae99f0f58b661339b26"></a>

## name property — api_protection_rules.api_endpoint_rules.metadata / 46391aa72d7a / 5

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

<a id="canonical-93f8d8ffe12e0e0b64e6a62fe0deb99c46332201cb9652549af293f7c57e2c14"></a>

## Next pages — api_protection_rules.api_endpoint_rules.metadata / 46391aa72d7a / 6

- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-c96eba43d7d0081fcdd77cdf4790ebd9c61f34b5fb8a78312b73a803fce74c50)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-df8514748df012f14f52ecfc543198dd4ba544ee223760d470b27fecfa657fe0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a00501c8ffb08d1aacb048fbd5b783f6153e6de6eee5dcc07c5f3f1f7a735eed"></a>

## api_protection_rules.api_endpoint_rules.request_matcher — api_protection_rules.api_endpoint_rules.request_matcher / 4e3cdece1f76 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-c96eba43d7d0081fcdd77cdf4790ebd9c61f34b5fb8a78312b73a803fce74c50)
- api_protection_rules.api_endpoint_rules.request_matcher

<a id="canonical-55b4032cc203853f7e5ab83c24721d4ab9fe859142dec49a39dce735baaa5f63"></a>

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

<a id="canonical-7d405332049a4e889f5b2cf72c5973269739896aa045fee9632d1c7b29da3af3"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher / 4e3cdece1f76 / 3

- [cookie_matchers](data-sources--http_loadbalancer--reference--group-005.md#canonical-3a535f08fb5acf97f4c84afd2a63ecc6f44da2e1da4f988df8cdf10ace1cc98f): complete subsection reference.

- [headers](data-sources--http_loadbalancer--reference--group-005.md#canonical-ad2b4e2f2012ba00fae9a27d351c4662c76e78c7cd460583c8b8ee723e4df145): complete subsection reference.

- [jwt_claims](data-sources--http_loadbalancer--reference--group-005.md#canonical-8548783742487d9518e74997225bdaa4f9a37ba78cc36891a3f19ebaf7ebbcfa): complete subsection reference.

- [query_params](data-sources--http_loadbalancer--reference--group-005.md#canonical-f469e31263618ed59e847c6910f7420401c09464a305c5ea8ed7474414669599): complete subsection reference.

<a id="canonical-da3ff12943e2effe680dda35f0fbc58041437ca50aa9d6ede4ddc58fff4d9a6e"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher / 4e3cdece1f76 / 4

- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-005.md#canonical-3a535f08fb5acf97f4c84afd2a63ecc6f44da2e1da4f988df8cdf10ace1cc98f)
- [api_protection_rules.api_endpoint_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-005.md#canonical-ad2b4e2f2012ba00fae9a27d351c4662c76e78c7cd460583c8b8ee723e4df145)
- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-005.md#canonical-8548783742487d9518e74997225bdaa4f9a37ba78cc36891a3f19ebaf7ebbcfa)
- [api_protection_rules.api_endpoint_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-005.md#canonical-f469e31263618ed59e847c6910f7420401c09464a305c5ea8ed7474414669599)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-c96eba43d7d0081fcdd77cdf4790ebd9c61f34b5fb8a78312b73a803fce74c50)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-3a535f08fb5acf97f4c84afd2a63ecc6f44da2e1da4f988df8cdf10ace1cc98f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7586547a6f2bfeddaf9ff280dd0723d58904f274db164e72df0adb5fa14789e3"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers / 3b017e0d1268 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-c96eba43d7d0081fcdd77cdf4790ebd9c61f34b5fb8a78312b73a803fce74c50)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-df8514748df012f14f52ecfc543198dd4ba544ee223760d470b27fecfa657fe0)
- api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers

<a id="canonical-c9e6208741182b0a46e319a15ba78420c1976aef8985f10a2e369cf63ac49e87"></a>

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

<a id="canonical-46b3a69f856f021b0773fe8a864e2aabae78a8e8307a2bb5d887e4f53b810ab1"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers / 3b017e0d1268 / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-005.md#canonical-1934fd94149e35397abf2d18ab354e29e7453326e1946b5e118d26e8c1d5158e): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-005.md#canonical-46f4ccefbc96fca170cc7c972ac81bb49f293f0c7a9849e5eb01524f0dea7f6c): complete subsection reference.

<a id="canonical-6ec671ef5b6fe340beec958d5f1c2c4dcd68e1ac201eda9b012fe9ad5edbc249"></a>

<a id="canonical-eea13f0cab725f4651c8c03ca37b5bd651e1edff0b8177a2600d1674b54ee8f8"></a>

## invert_matcher property — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers / 3b017e0d1268 / 4

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

- [item](data-sources--http_loadbalancer--reference--group-005.md#canonical-452096ffe2c0b91322282cd1695dbe4d6acb5eb1172a1f4170ef47c85658ddfe): complete subsection reference.

<a id="canonical-fbabe9064cb72663131453ca733c993921955b003175b471ebca4a971e650b7a"></a>

<a id="canonical-566d2540b34ab12811fecff5bfffb793f6dcd1fae08bfd637df95c15d6bddd2a"></a>

## name property — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers / 3b017e0d1268 / 5

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

<a id="canonical-2d759c69dee044cf98ed539f889a10f8069fde72daf41f6e49a4f58e4cb9a016"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers / 3b017e0d1268 / 6

- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present](data-sources--http_loadbalancer--reference--group-005.md#canonical-1934fd94149e35397abf2d18ab354e29e7453326e1946b5e118d26e8c1d5158e)
- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_present](data-sources--http_loadbalancer--reference--group-005.md#canonical-46f4ccefbc96fca170cc7c972ac81bb49f293f0c7a9849e5eb01524f0dea7f6c)
- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item](data-sources--http_loadbalancer--reference--group-005.md#canonical-452096ffe2c0b91322282cd1695dbe4d6acb5eb1172a1f4170ef47c85658ddfe)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-df8514748df012f14f52ecfc543198dd4ba544ee223760d470b27fecfa657fe0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1934fd94149e35397abf2d18ab354e29e7453326e1946b5e118d26e8c1d5158e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4c05462f00051bc72ec26d0558afde18fe9d9c18a7ed74aba7777d4bc32b42c2"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_no / 304a8b28e46d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-c96eba43d7d0081fcdd77cdf4790ebd9c61f34b5fb8a78312b73a803fce74c50)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-df8514748df012f14f52ecfc543198dd4ba544ee223760d470b27fecfa657fe0)
- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-005.md#canonical-3a535f08fb5acf97f4c84afd2a63ecc6f44da2e1da4f988df8cdf10ace1cc98f)
- api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-3e24eeed34c9bfd589a9eb8dd7de88f26e5c4802aa5cc9b28884312bb4ab8772"></a>

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

<a id="canonical-8948edf64f338b6023b9c594a10ba88422125bf21ccc2cbc788b532afaa3a55b"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_no / 304a8b28e46d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-21cfc7ae66319c76915b76e85770afa1000d6d599febff78a78985300f3f848e"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_no / 304a8b28e46d / 4

- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-005.md#canonical-3a535f08fb5acf97f4c84afd2a63ecc6f44da2e1da4f988df8cdf10ace1cc98f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-46f4ccefbc96fca170cc7c972ac81bb49f293f0c7a9849e5eb01524f0dea7f6c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-80e027b187deee71fc410984c3a5d496325c722677f3295549f0ff3c6dfb1aa3"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_present — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_pr / 2d1256056b13 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-c96eba43d7d0081fcdd77cdf4790ebd9c61f34b5fb8a78312b73a803fce74c50)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-df8514748df012f14f52ecfc543198dd4ba544ee223760d470b27fecfa657fe0)
- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-005.md#canonical-3a535f08fb5acf97f4c84afd2a63ecc6f44da2e1da4f988df8cdf10ace1cc98f)
- api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-ddaf832ea2a15bb2ff2a67f8bc91f24c3e1874141f4a4ee9c246d1fec1fb05b1"></a>

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

<a id="canonical-b990e12098f87dc53adb67569788223a96e870eeda2bddf1271c9ee8a148273c"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_pr / 2d1256056b13 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bc07792d4829a95df46587d586d97f133d9f93c00a9b80d5a14107a4030328bd"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_pr / 2d1256056b13 / 4

- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-005.md#canonical-3a535f08fb5acf97f4c84afd2a63ecc6f44da2e1da4f988df8cdf10ace1cc98f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-452096ffe2c0b91322282cd1695dbe4d6acb5eb1172a1f4170ef47c85658ddfe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1b593ab8b6c6942db367f6377fa353afb936784d8cfaa9d6393614438662c3e9"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item / aa8a0ef0f0de / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-c96eba43d7d0081fcdd77cdf4790ebd9c61f34b5fb8a78312b73a803fce74c50)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-df8514748df012f14f52ecfc543198dd4ba544ee223760d470b27fecfa657fe0)
- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-005.md#canonical-3a535f08fb5acf97f4c84afd2a63ecc6f44da2e1da4f988df8cdf10ace1cc98f)
- api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item

<a id="canonical-403b3882e2f1bcbd9e4561f65cec2973f63f58f6f8c599b4b8a5ce3fdc59f352"></a>

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

<a id="canonical-a4c0575bbf7457c434a128fcaf5d29703c1e5e59335786dc4875c495f64d597c"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item / aa8a0ef0f0de / 3

<a id="canonical-03007c32b9b0e3590ff9919292c5e2935431e47ab5259b58f2a29ae4597b8607"></a>

<a id="canonical-92d198d00f2f2a12e13484206ff9d5d3669f946e1b824a559be1672149ff25b5"></a>

## exact_values property — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item / aa8a0ef0f0de / 4

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

<a id="canonical-509546f3c0e5046198de7461607f4006e42db010ace8926ee9f91d3319bae6ea"></a>

<a id="canonical-71737093fa41ba5affe8497174de6e02430a9685412c6d4cf0ad46b681bc86d9"></a>

## regex_values property — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item / aa8a0ef0f0de / 5

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

<a id="canonical-b8a89d87019bd25070946f29212c9e2258817e33e9ce1a2ff4df4c4db282f7ec"></a>

<a id="canonical-f31fef81b828afbf415333d6de613db6baa1b45d6598b40dde732f125ed438c5"></a>

## transformers property — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item / aa8a0ef0f0de / 6

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

<a id="canonical-331dcc5c6db102cf759c4a4c71dee3d8758237407d6c037ee680b4c4fd7965af"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item / aa8a0ef0f0de / 7

- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-005.md#canonical-3a535f08fb5acf97f4c84afd2a63ecc6f44da2e1da4f988df8cdf10ace1cc98f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ad2b4e2f2012ba00fae9a27d351c4662c76e78c7cd460583c8b8ee723e4df145"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2cc355cbd8c438a31974957c6cb64b8b62bd3f9163b5152a9d641dd0e86a758c"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.headers — api_protection_rules.api_endpoint_rules.request_matcher.headers / ef45ab0fc7d3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-c96eba43d7d0081fcdd77cdf4790ebd9c61f34b5fb8a78312b73a803fce74c50)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-df8514748df012f14f52ecfc543198dd4ba544ee223760d470b27fecfa657fe0)
- api_protection_rules.api_endpoint_rules.request_matcher.headers

<a id="canonical-6b345627b346895e79efda8fe4d866260ecd5108018f2937c4272dbd5ea57735"></a>

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

<a id="canonical-eeba499e9439800b490744888eeed2a89822017d0ef127138bd5bd8f43a6b987"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.headers / ef45ab0fc7d3 / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-005.md#canonical-6d5b4b019b78c82649558989f80b70287ba4f73be573c952a5653b27e71c3a85): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-005.md#canonical-98b4fad687112d1fe4fe542ddf5a0068edf6dfeb159d92e12944934c4922098f): complete subsection reference.

<a id="canonical-9e25867cb540070db8feb09b40ecccfc8490340bb7ea9a6796478724190f577e"></a>

<a id="canonical-502b00d0f96aa1abb9b944a54aa53520f31beb5019031c60eb8eba7af3153194"></a>

## invert_matcher property — api_protection_rules.api_endpoint_rules.request_matcher.headers / ef45ab0fc7d3 / 4

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

- [item](data-sources--http_loadbalancer--reference--group-005.md#canonical-5b3c2c072899c8dce46eb3c0a1d900c5c565879191ba6b4e67b12567e2df1d46): complete subsection reference.

<a id="canonical-daa64469461f8782a278ddb06dd17e2f4ede48f60d195a1218d667e8d0470d98"></a>

<a id="canonical-83856a6f40c02934f19edf803c0446b6bccbec87c2da8b935272e7f824a62b8a"></a>

## name property — api_protection_rules.api_endpoint_rules.request_matcher.headers / ef45ab0fc7d3 / 5

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

<a id="canonical-7299dcaa27c99e0187e92cd3ce406a80d4fccaf242b1daca3249fb20386af954"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.headers / ef45ab0fc7d3 / 6

- [api_protection_rules.api_endpoint_rules.request_matcher.headers.check_not_present](data-sources--http_loadbalancer--reference--group-005.md#canonical-6d5b4b019b78c82649558989f80b70287ba4f73be573c952a5653b27e71c3a85)
- [api_protection_rules.api_endpoint_rules.request_matcher.headers.check_present](data-sources--http_loadbalancer--reference--group-005.md#canonical-98b4fad687112d1fe4fe542ddf5a0068edf6dfeb159d92e12944934c4922098f)
- [api_protection_rules.api_endpoint_rules.request_matcher.headers.item](data-sources--http_loadbalancer--reference--group-005.md#canonical-5b3c2c072899c8dce46eb3c0a1d900c5c565879191ba6b4e67b12567e2df1d46)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-df8514748df012f14f52ecfc543198dd4ba544ee223760d470b27fecfa657fe0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6d5b4b019b78c82649558989f80b70287ba4f73be573c952a5653b27e71c3a85"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ce79e065966baebbab7864063b04f64654f515bd62cac42f972d9ff4fb60e0b"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.headers.check_not_present — api_protection_rules.api_endpoint_rules.request_matcher.headers.check_not_presen / b66004cc9cc1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-c96eba43d7d0081fcdd77cdf4790ebd9c61f34b5fb8a78312b73a803fce74c50)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-df8514748df012f14f52ecfc543198dd4ba544ee223760d470b27fecfa657fe0)
- [api_protection_rules.api_endpoint_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-005.md#canonical-ad2b4e2f2012ba00fae9a27d351c4662c76e78c7cd460583c8b8ee723e4df145)
- api_protection_rules.api_endpoint_rules.request_matcher.headers.check_not_present

<a id="canonical-181381798d7a2fe9bca20d632d041d168c4d8c0c0b1fbbb42b90bd37ce869c1a"></a>

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

<a id="canonical-62b09b3ce7e1f1ad80514fe75c3af52ddb3e53bf05b2c2ea53be5c88aa579863"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.headers.check_not_presen / b66004cc9cc1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fd0bede36f1588ceaf96898c901e8cd306ad93169ad72677175a1f890ddee098"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.headers.check_not_presen / b66004cc9cc1 / 4

- [api_protection_rules.api_endpoint_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-005.md#canonical-ad2b4e2f2012ba00fae9a27d351c4662c76e78c7cd460583c8b8ee723e4df145)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-98b4fad687112d1fe4fe542ddf5a0068edf6dfeb159d92e12944934c4922098f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1bdd8152f301bd63813954c5cc7f53aa04865632de1dce964d7843f10b96d36f"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.headers.check_present — api_protection_rules.api_endpoint_rules.request_matcher.headers.check_present / ff6924cecd62 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-c96eba43d7d0081fcdd77cdf4790ebd9c61f34b5fb8a78312b73a803fce74c50)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-df8514748df012f14f52ecfc543198dd4ba544ee223760d470b27fecfa657fe0)
- [api_protection_rules.api_endpoint_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-005.md#canonical-ad2b4e2f2012ba00fae9a27d351c4662c76e78c7cd460583c8b8ee723e4df145)
- api_protection_rules.api_endpoint_rules.request_matcher.headers.check_present

<a id="canonical-f19ecb6d3dcebb9d65a9eff5c52e0f2e86df02b40270ea6295e3d0a6e58ca878"></a>

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

<a id="canonical-653961f102314a6e8e164404330ee19fe753e376220043a74f7571a350641f68"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.headers.check_present / ff6924cecd62 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3a1d45f42465f4dfe8f9a94d2cb79945fda7c77dcf9d3803f37153ed29c8af85"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.headers.check_present / ff6924cecd62 / 4

- [api_protection_rules.api_endpoint_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-005.md#canonical-ad2b4e2f2012ba00fae9a27d351c4662c76e78c7cd460583c8b8ee723e4df145)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5b3c2c072899c8dce46eb3c0a1d900c5c565879191ba6b4e67b12567e2df1d46"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-36b1d500b0b39802f7ed4a9198b3e1752c55d28114e4ae3cc3ed7a92b06399cd"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.headers.item — api_protection_rules.api_endpoint_rules.request_matcher.headers.item / 87e92ebb24ca / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-c96eba43d7d0081fcdd77cdf4790ebd9c61f34b5fb8a78312b73a803fce74c50)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-df8514748df012f14f52ecfc543198dd4ba544ee223760d470b27fecfa657fe0)
- [api_protection_rules.api_endpoint_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-005.md#canonical-ad2b4e2f2012ba00fae9a27d351c4662c76e78c7cd460583c8b8ee723e4df145)
- api_protection_rules.api_endpoint_rules.request_matcher.headers.item

<a id="canonical-4a0aed1d12b6b2696aa27a0e388018eb977bedb42518b9b542deee015a43dd95"></a>

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

<a id="canonical-0bd11c04fd8b86855b63e27e123d024d454e566b49dd34b37fe35f6007de7829"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.headers.item / 87e92ebb24ca / 3

<a id="canonical-fbdc38ebbd0fe557c18895198f78476f23b4f89c033ead099c97d46f60bc174e"></a>

<a id="canonical-4755f502c0413c22b5467dc8caa983679e804d0a3c8b02af3573da18de9a28b2"></a>

## exact_values property — api_protection_rules.api_endpoint_rules.request_matcher.headers.item / 87e92ebb24ca / 4

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

<a id="canonical-e269b09cbb1944ae52d4d74f45301e7ac8834bf892ae7dd8abd3631480e0450b"></a>

<a id="canonical-5e2c3817a95e1bfc7e586a8cc04b19af1455efcc8f887455af37ce6916347cda"></a>

## regex_values property — api_protection_rules.api_endpoint_rules.request_matcher.headers.item / 87e92ebb24ca / 5

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

<a id="canonical-82c82360919dc66e74e1e7e7b71a630e3d3f8f3e7f092bb528973d1dc9b3bcf6"></a>

<a id="canonical-eeb1898e5909714d168fd122819636c6278a5ea0f8a9557ed470d148a4aea1fd"></a>

## transformers property — api_protection_rules.api_endpoint_rules.request_matcher.headers.item / 87e92ebb24ca / 6

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

<a id="canonical-502632e83109830d49498f3796f7eb70ac11d34b9c495b6d4db5f4c65ca80396"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.headers.item / 87e92ebb24ca / 7

- [api_protection_rules.api_endpoint_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-005.md#canonical-ad2b4e2f2012ba00fae9a27d351c4662c76e78c7cd460583c8b8ee723e4df145)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-8548783742487d9518e74997225bdaa4f9a37ba78cc36891a3f19ebaf7ebbcfa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-398642d29e41bdf513060140f19159c5eb4b0a62127ae69d7c08f61ea8c22339"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims / 94521669c168 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-c96eba43d7d0081fcdd77cdf4790ebd9c61f34b5fb8a78312b73a803fce74c50)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-df8514748df012f14f52ecfc543198dd4ba544ee223760d470b27fecfa657fe0)
- api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims

<a id="canonical-caa8031aed0726ab27b1a8dc67294201c4f76d8f124d6ccafa062f56c356266a"></a>

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

<a id="canonical-1b9da2a8ac735a5f4edba576882da2da9421f9f18de3cbfcbb7b11b350b93732"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims / 94521669c168 / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-005.md#canonical-de94a01bd98117dc3a296c133a257f15fa38188d5acf2dbc6b85cd3bd1d7ae2c): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-005.md#canonical-1b125f082658ab89c71e7accf55657fe990c9ec04b38dc6adf75ed71745bee46): complete subsection reference.

<a id="canonical-16e2bcdfe2cfd42ab9764dacf859b04522d5ce97b8950af229c3822f069b3f70"></a>

<a id="canonical-7e279590eacf4f11cae077ebfbfbb33f2fcf3a4a70c43d2cfe843396608f7d95"></a>

## invert_matcher property — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims / 94521669c168 / 4

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

- [item](data-sources--http_loadbalancer--reference--group-005.md#canonical-53e144c39f82a74ebf131a73d57f23635900da13babc55261b7ec29301454b58): complete subsection reference.

<a id="canonical-b3b0bcd485c2ad476aea77a40ba6a8b69b799be0dbd54cd854915c281ba00c2e"></a>

<a id="canonical-baf6d7926d8170285c5695a6cac516fb6f3e9e4633a1a2342390fdc0489375ae"></a>

## name property — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims / 94521669c168 / 5

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

<a id="canonical-3a3bfdabc5632e6eb215bf31cb50d44ff5bc73cbbb2509c6eb31ccccb1b344c5"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims / 94521669c168 / 6

- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_not_present](data-sources--http_loadbalancer--reference--group-005.md#canonical-de94a01bd98117dc3a296c133a257f15fa38188d5acf2dbc6b85cd3bd1d7ae2c)
- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_present](data-sources--http_loadbalancer--reference--group-005.md#canonical-1b125f082658ab89c71e7accf55657fe990c9ec04b38dc6adf75ed71745bee46)
- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item](data-sources--http_loadbalancer--reference--group-005.md#canonical-53e144c39f82a74ebf131a73d57f23635900da13babc55261b7ec29301454b58)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-df8514748df012f14f52ecfc543198dd4ba544ee223760d470b27fecfa657fe0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-de94a01bd98117dc3a296c133a257f15fa38188d5acf2dbc6b85cd3bd1d7ae2c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac0318304b0ebe071ea3c43760264f791636294ff54e7a50c835dc142c385260"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_not_present — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_not_pre / 01ba681793f3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-c96eba43d7d0081fcdd77cdf4790ebd9c61f34b5fb8a78312b73a803fce74c50)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-df8514748df012f14f52ecfc543198dd4ba544ee223760d470b27fecfa657fe0)
- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-005.md#canonical-8548783742487d9518e74997225bdaa4f9a37ba78cc36891a3f19ebaf7ebbcfa)
- api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-1ed6383bfd7fc29479b1b51fd300db14fba4cd16e24e744d4934a450cc31a797"></a>

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

<a id="canonical-6ea1db536903e846f9d302c3a5b2e292890a825e67ed0fcf1b3f358759f2d6a7"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_not_pre / 01ba681793f3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-68b7377c8a449beb43e1f1b1ac53ca57e477601cdbb435904317999935366520"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_not_pre / 01ba681793f3 / 4

- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-005.md#canonical-8548783742487d9518e74997225bdaa4f9a37ba78cc36891a3f19ebaf7ebbcfa)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1b125f082658ab89c71e7accf55657fe990c9ec04b38dc6adf75ed71745bee46"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-87d4d59816d88db3751cd432ee70cac94a5c9d2e98f2b7416f4b9bfd71c37dcb"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_present — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_present / 10a9f16da888 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-c96eba43d7d0081fcdd77cdf4790ebd9c61f34b5fb8a78312b73a803fce74c50)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-df8514748df012f14f52ecfc543198dd4ba544ee223760d470b27fecfa657fe0)
- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-005.md#canonical-8548783742487d9518e74997225bdaa4f9a37ba78cc36891a3f19ebaf7ebbcfa)
- api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_present

<a id="canonical-7b975d952c10437b168235e91649de71f82d55bdf08c40802affd0371749a4b7"></a>

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

<a id="canonical-7408a515427a17b12bb497aca8a4c8acb23ae20f1bbcbf9db36a8b08c49ee3a3"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_present / 10a9f16da888 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b59bccca364f13243d8a1ba73ffc029fe45a3f2e2e4905be0999fb5f703af55d"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_present / 10a9f16da888 / 4

- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-005.md#canonical-8548783742487d9518e74997225bdaa4f9a37ba78cc36891a3f19ebaf7ebbcfa)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-53e144c39f82a74ebf131a73d57f23635900da13babc55261b7ec29301454b58"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4f8eac096024a3eb345e3014234c39731b8ae5bcc1dda0fde79ac0457e13d22f"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item / 1f92370cc33b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-c96eba43d7d0081fcdd77cdf4790ebd9c61f34b5fb8a78312b73a803fce74c50)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-df8514748df012f14f52ecfc543198dd4ba544ee223760d470b27fecfa657fe0)
- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-005.md#canonical-8548783742487d9518e74997225bdaa4f9a37ba78cc36891a3f19ebaf7ebbcfa)
- api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item

<a id="canonical-e91b8a35c52b75b92464c9ace20b334161033b152562b5996b48e78e45775f24"></a>

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

<a id="canonical-7dba67a78d1fa8adc5ea6ad65db63b89a90cf93c3bfa6ba3f7e74a64b8c055db"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item / 1f92370cc33b / 3

<a id="canonical-ecb39ce6523a56217f2ea476a363811357f57922b537fad042dcaf5decf2e8d9"></a>

<a id="canonical-a4faf466db3cd0a4f255b41b33488b94ea83189a7a36c4799dfeabc7d8d94e1e"></a>

## exact_values property — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item / 1f92370cc33b / 4

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

<a id="canonical-47bc0ab3227ed553a20a29010e72ab92cb7c64ed663776e1f3b8f26a5d052d02"></a>

<a id="canonical-6ac969c3c57f300a6092e097060f3a02378b408447b9c1d265aa33ce45c8f5d4"></a>

## regex_values property — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item / 1f92370cc33b / 5

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

<a id="canonical-b4ccbad58b08b21093bd94d865115a3f088e1017e2e66052f568e3cd86c1e76c"></a>

<a id="canonical-cac6fa950cd76bbd54cf1b05c542aa7e9556563ecd5cb5ce8a6baddaad8f8c7f"></a>

## transformers property — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item / 1f92370cc33b / 6

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

<a id="canonical-6020c6796ec6b66978bdebe5bfa6bcf4f34e0d4aa4042aa4f6865ede92323a15"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item / 1f92370cc33b / 7

- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-005.md#canonical-8548783742487d9518e74997225bdaa4f9a37ba78cc36891a3f19ebaf7ebbcfa)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f469e31263618ed59e847c6910f7420401c09464a305c5ea8ed7474414669599"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9830ae78a4aa538c56102a8faa7927e5639e64d86f6281001ed41f436c3306e"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.query_params — api_protection_rules.api_endpoint_rules.request_matcher.query_params / 86c2eb2ab22b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-c96eba43d7d0081fcdd77cdf4790ebd9c61f34b5fb8a78312b73a803fce74c50)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-df8514748df012f14f52ecfc543198dd4ba544ee223760d470b27fecfa657fe0)
- api_protection_rules.api_endpoint_rules.request_matcher.query_params

<a id="canonical-54cf773248542e658c6b4c3361d10857d32ac695b3efc5ea2dcce7973419379c"></a>

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

<a id="canonical-63474448859b5335bcadefed78112df090fcecd670de7944a5e6d27bb087e546"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.query_params / 86c2eb2ab22b / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-005.md#canonical-a8156f0a34c7e1cc90c213af8f94be1dcdad80df008e5dde53f38e5e8400e6d5): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-005.md#canonical-d88fb8c13af7b254f18e8a46bce073a34c41b2894747ae7c96f715f51b12e5d6): complete subsection reference.

<a id="canonical-fcccecc701360b51944602de177f1d72c1b4ee96873acb0746b1998d4e38fe7f"></a>

<a id="canonical-a9119f9c2f9cd5f421b4ff124be41d87a25f790dcc535efd8b79ceb2bb0c148b"></a>

## invert_matcher property — api_protection_rules.api_endpoint_rules.request_matcher.query_params / 86c2eb2ab22b / 4

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

- [item](data-sources--http_loadbalancer--reference--group-005.md#canonical-41310b60a4873e5327c95080fe73925a2c8419c057df5bd372ad8d2e930d8c7e): complete subsection reference.

<a id="canonical-4ff4a818c65c56d50f5b79fc1af6e03076f0ec8072d153126ac947a69ac2680f"></a>

<a id="canonical-9be830ed3da33b132f479eefab0a8f5ed4c3b6e9117ae252d8ab2fd6819890de"></a>

## key property — api_protection_rules.api_endpoint_rules.request_matcher.query_params / 86c2eb2ab22b / 5

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

<a id="canonical-e757c723aa740edbc76021edf1b11db791595bf7ceb0c03cf5cba0d92f00c620"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.query_params / 86c2eb2ab22b / 6

- [api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_not_present](data-sources--http_loadbalancer--reference--group-005.md#canonical-a8156f0a34c7e1cc90c213af8f94be1dcdad80df008e5dde53f38e5e8400e6d5)
- [api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_present](data-sources--http_loadbalancer--reference--group-005.md#canonical-d88fb8c13af7b254f18e8a46bce073a34c41b2894747ae7c96f715f51b12e5d6)
- [api_protection_rules.api_endpoint_rules.request_matcher.query_params.item](data-sources--http_loadbalancer--reference--group-005.md#canonical-41310b60a4873e5327c95080fe73925a2c8419c057df5bd372ad8d2e930d8c7e)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-df8514748df012f14f52ecfc543198dd4ba544ee223760d470b27fecfa657fe0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a8156f0a34c7e1cc90c213af8f94be1dcdad80df008e5dde53f38e5e8400e6d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ecef3d2b3a443c897ef1bbae63754b44899bb131e7f71a331e367e599096cfc2"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_not_present — api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_not_p / 1f4a1e6cb5f3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-c96eba43d7d0081fcdd77cdf4790ebd9c61f34b5fb8a78312b73a803fce74c50)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-df8514748df012f14f52ecfc543198dd4ba544ee223760d470b27fecfa657fe0)
- [api_protection_rules.api_endpoint_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-005.md#canonical-f469e31263618ed59e847c6910f7420401c09464a305c5ea8ed7474414669599)
- api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_not_present

<a id="canonical-ad2f37f86a1ba8763621c5ec7442421847a3f4031338178971512ba6715821ca"></a>

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

<a id="canonical-da1101a386bde877879d90fc1f155c99332c1924ac533fe4b861bb3f83d34b1c"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_not_p / 1f4a1e6cb5f3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-04e7d7b51c93d8770672727c6fd552ea6b28d1a85fb5c031803e3cc77626f239"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_not_p / 1f4a1e6cb5f3 / 4

- [api_protection_rules.api_endpoint_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-005.md#canonical-f469e31263618ed59e847c6910f7420401c09464a305c5ea8ed7474414669599)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d88fb8c13af7b254f18e8a46bce073a34c41b2894747ae7c96f715f51b12e5d6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a31fc686310fb74763eadb08c9730e66f86f9ba4883369cd848d2c52d952651f"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_present — api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_prese / 17598322089a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-c96eba43d7d0081fcdd77cdf4790ebd9c61f34b5fb8a78312b73a803fce74c50)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-df8514748df012f14f52ecfc543198dd4ba544ee223760d470b27fecfa657fe0)
- [api_protection_rules.api_endpoint_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-005.md#canonical-f469e31263618ed59e847c6910f7420401c09464a305c5ea8ed7474414669599)
- api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_present

<a id="canonical-cb28832305ceb7d1224e6576c88021e71ac0325cadf35a2444c3b2d3334e11d9"></a>

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

<a id="canonical-e2f8d255ba91892aec3d772b524840eb955ae46274f52b0357b41e913de50275"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_prese / 17598322089a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-74d074403dde2d785cab6dba4695a993706fc15f285d44d3d7c05647118e999b"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_prese / 17598322089a / 4

- [api_protection_rules.api_endpoint_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-005.md#canonical-f469e31263618ed59e847c6910f7420401c09464a305c5ea8ed7474414669599)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-41310b60a4873e5327c95080fe73925a2c8419c057df5bd372ad8d2e930d8c7e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c3e1bc6ed2591de61363ef82b50d0e4df55d721bacca3cbdc4b23ac3b6824b4"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.query_params.item — api_protection_rules.api_endpoint_rules.request_matcher.query_params.item / 52f240c344ad / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-c96eba43d7d0081fcdd77cdf4790ebd9c61f34b5fb8a78312b73a803fce74c50)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-df8514748df012f14f52ecfc543198dd4ba544ee223760d470b27fecfa657fe0)
- [api_protection_rules.api_endpoint_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-005.md#canonical-f469e31263618ed59e847c6910f7420401c09464a305c5ea8ed7474414669599)
- api_protection_rules.api_endpoint_rules.request_matcher.query_params.item

<a id="canonical-658dc56f247ad5c05753d46aee842f00a6549243c7ae6a87a95ac00eb9703db0"></a>

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

<a id="canonical-a62eb4f36394b8bb78c44a4604efc749b31ec664db8c2bfbdf5277d68e4ac70e"></a>

## Direct properties — api_protection_rules.api_endpoint_rules.request_matcher.query_params.item / 52f240c344ad / 3

<a id="canonical-1554d4583c7db865c32504ad53f7f472d5c95d19b3fec9b6645471c2f7f63ab8"></a>

<a id="canonical-fe4bc5d5a7ff92c76aaf6021d7ddedd78114feda725a7c21d7e4a52dc3665fbc"></a>

## exact_values property — api_protection_rules.api_endpoint_rules.request_matcher.query_params.item / 52f240c344ad / 4

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

<a id="canonical-8e5213116acdba58ff6a70facef6eb947ef2331daa3620fc6f61ac96cdb60970"></a>

<a id="canonical-1207810e4346cf60233592190eb79fcfc646c82d21965bb40dbce1c0d2557042"></a>

## regex_values property — api_protection_rules.api_endpoint_rules.request_matcher.query_params.item / 52f240c344ad / 5

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

<a id="canonical-e4f61cd2bf05edd48c8ba1cd39e43b80d59b5136814640e2ee1fecbbd50ce775"></a>

<a id="canonical-e12bebdf9a7cf8d378b93b372e67d46077c6a50610c44b1a7af48faaa55f93b3"></a>

## transformers property — api_protection_rules.api_endpoint_rules.request_matcher.query_params.item / 52f240c344ad / 6

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

<a id="canonical-2795aec1b96cec2b6329029670ca2bfacf7597214b0277e2f6b34ca924411aef"></a>

## Next pages — api_protection_rules.api_endpoint_rules.request_matcher.query_params.item / 52f240c344ad / 7

- [api_protection_rules.api_endpoint_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-005.md#canonical-f469e31263618ed59e847c6910f7420401c09464a305c5ea8ed7474414669599)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-86159ba371e34dac257c4dacb1386acafde12abf10fea1b5f1e62a611755b842"></a>

## api_protection_rules.api_groups_rules — api_protection_rules.api_groups_rules / f9b294a45302 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- api_protection_rules.api_groups_rules

<a id="canonical-85956d0e2c9041746576eed6f78671eb0d6487bbbdaa514ebf3552e4ecd901e4"></a>

Type: `"list"`. Computed.

Category includes rules per API group or Server URL. For API groups, refer to API Definition which
includes API groups derived from uploaded swaggers.

Upstream description:

This category includes rules per API group or Server URL. For API groups, refer to API Definition
which includes API groups derived from uploaded swaggers.

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

<a id="canonical-33c5a05c613b304b8968446569b19e1b719ba7c94610a1bf1a5a73e3ac682015"></a>

## Direct properties — api_protection_rules.api_groups_rules / f9b294a45302 / 3

- [action](data-sources--http_loadbalancer--reference--group-005.md#canonical-c86ede9ac2831bab5624eb5cbb9df404ab10aaa5682a9152f649c64b8935093f): complete subsection reference.

- [any_domain](data-sources--http_loadbalancer--reference--group-005.md#canonical-6645ce780c54e522c879903ec8481be4f7205eafc370d3b6fca2b83a79061572): complete subsection reference.

<a id="canonical-5a8fb0bae4db20a3f4848a0c08617f7adc3fcb764bd68360273b83b16de17945"></a>

<a id="canonical-7208d7f1fd40f618065e6f2ae47f4ff3d7b7e095fdcff497fee58ac56bbd30e7"></a>

## api_group property — api_protection_rules.api_groups_rules / f9b294a45302 / 4

Type: `"string"`. Computed.

API groups derived from API Definition swaggers. For example oas-all-operations including all paths
and methods from the swaggers, oas-base-URLs covering all requests under base-paths from the
swaggers. Custom groups can be created if user tags paths or operations with 'x-F5 Distributed..

Upstream description:

API groups derived from API Definition swaggers. For example oas-all-operations including all paths
and methods from the swaggers, oas-base-URLs covering all requests under base-paths from the
swaggers. Custom groups can be created if user tags paths or operations with "x-F5 Distributed
Cloud-API-group" extensions inside swaggers.

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

<a id="canonical-dd3fd92c908c620f0f38cd20723084d498de68d9b98d636e1b90006c157315a8"></a>

<a id="canonical-8c60067c14ccc9e1dccb155ef752d1dddc6fc53471bded8961d504db1aa6c4be"></a>

## base_path property — api_protection_rules.api_groups_rules / f9b294a45302 / 5

Type: `"string"`. Computed.

Base Path. Prefix of the request path. For example: /v1.

Upstream description:

Prefix of the request path. For example: /v1.

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

- [client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-32a937c93dd23dd269c99657114ebddca1bb524608724bbebd46e3a69fad6e5a): complete subsection reference.

- [metadata](data-sources--http_loadbalancer--reference--group-005.md#canonical-8f9610f28d4a71e22bbc8a4bf122b9d4e075db843ed5e87344008d76f9808f3f): complete subsection reference.

- [request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1799983b63e67772d3b71c4156c2cd18db14ba86436ab2971e1a87048d4ba462): complete subsection reference.

<a id="canonical-a197440385e6a12ba3c61f2a656db55878731f0a55f069bcde36e802d948fc36"></a>

<a id="canonical-1805af36f641a2a2ed69d9a831ae1953ea1eb7ff2e46e78a8b9731f6233c679a"></a>

## specific_domain property — api_protection_rules.api_groups_rules / f9b294a45302 / 6

Type: `"string"`. Computed.

Exclusive with \[any\_domain\] The rule will apply for a specific domain. For

Upstream description:

Exclusive with \[any\_domain\] The rule will apply for a specific domain. For example:
api.example.com.

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

<a id="canonical-f5ee7149ff4e854bb80f02666342eae777034a6b0e2b8a5d7a045943f8d872c8"></a>

## Next pages — api_protection_rules.api_groups_rules / f9b294a45302 / 7

- [api_protection_rules.api_groups_rules.action](data-sources--http_loadbalancer--reference--group-005.md#canonical-c86ede9ac2831bab5624eb5cbb9df404ab10aaa5682a9152f649c64b8935093f)
- [api_protection_rules.api_groups_rules.any_domain](data-sources--http_loadbalancer--reference--group-005.md#canonical-6645ce780c54e522c879903ec8481be4f7205eafc370d3b6fca2b83a79061572)
- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-32a937c93dd23dd269c99657114ebddca1bb524608724bbebd46e3a69fad6e5a)
- [api_protection_rules.api_groups_rules.metadata](data-sources--http_loadbalancer--reference--group-005.md#canonical-8f9610f28d4a71e22bbc8a4bf122b9d4e075db843ed5e87344008d76f9808f3f)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1799983b63e67772d3b71c4156c2cd18db14ba86436ab2971e1a87048d4ba462)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c86ede9ac2831bab5624eb5cbb9df404ab10aaa5682a9152f649c64b8935093f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e098046044ebd920012a2fcd2a349ae97e164c45ae1b8362810b9278778caf40"></a>

## api_protection_rules.api_groups_rules.action — api_protection_rules.api_groups_rules.action / d2d36d1f48a5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- api_protection_rules.api_groups_rules.action

<a id="canonical-416580969bd065cb73939c560e73c7c4f077c7f8c8e1c48e52b9e81884b640fa"></a>

Type: `"single"`. Computed.

The action to take if the input request matches the rule.

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

<a id="canonical-86845c65dcb28e6e4d59ea770160cc5771077b369732c02cc23e2306cce14738"></a>

## Direct properties — api_protection_rules.api_groups_rules.action / d2d36d1f48a5 / 3

- [allow](data-sources--http_loadbalancer--reference--group-005.md#canonical-cf3f34354103c0689a7646a8d7f3102c39f2de305106f9c286f6f9432d65cfa0): complete subsection reference.

- [deny](data-sources--http_loadbalancer--reference--group-005.md#canonical-8b58d87ba8296d32d8132c8c4dd3eb35f8a81b3206456e303dac9380e5813bf2): complete subsection reference.

<a id="canonical-24548537b90a0f20e4448c0d05c2a5e7ae77a30fcaad264b0308b2bcbd815fb7"></a>

## Next pages — api_protection_rules.api_groups_rules.action / d2d36d1f48a5 / 4

- [api_protection_rules.api_groups_rules.action.allow](data-sources--http_loadbalancer--reference--group-005.md#canonical-cf3f34354103c0689a7646a8d7f3102c39f2de305106f9c286f6f9432d65cfa0)
- [api_protection_rules.api_groups_rules.action.deny](data-sources--http_loadbalancer--reference--group-005.md#canonical-8b58d87ba8296d32d8132c8c4dd3eb35f8a81b3206456e303dac9380e5813bf2)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-cf3f34354103c0689a7646a8d7f3102c39f2de305106f9c286f6f9432d65cfa0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d74b96b1bd4e8a8a3f9854bd66950c1acc17f75a52c90f670de1acc853ae761a"></a>

## api_protection_rules.api_groups_rules.action.allow — api_protection_rules.api_groups_rules.action.allow / 2dd50712d866 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [api_protection_rules.api_groups_rules.action](data-sources--http_loadbalancer--reference--group-005.md#canonical-c86ede9ac2831bab5624eb5cbb9df404ab10aaa5682a9152f649c64b8935093f)
- api_protection_rules.api_groups_rules.action.allow

<a id="canonical-b9daa8e9bf67fc65997b4a0dfdd014de4c86e4ee1d42789bc3bb2926979fd254"></a>

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

<a id="canonical-6ec902eca9d62f9789659adaab5fb61b508a2c45e2ff896d2171d52867dd23ef"></a>

## Direct properties — api_protection_rules.api_groups_rules.action.allow / 2dd50712d866 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fdb42cbc1f48afb1babe64a1c044cd998c6fd2582612d4a325fed0409b17254e"></a>

## Next pages — api_protection_rules.api_groups_rules.action.allow / 2dd50712d866 / 4

- [api_protection_rules.api_groups_rules.action](data-sources--http_loadbalancer--reference--group-005.md#canonical-c86ede9ac2831bab5624eb5cbb9df404ab10aaa5682a9152f649c64b8935093f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-8b58d87ba8296d32d8132c8c4dd3eb35f8a81b3206456e303dac9380e5813bf2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a0fa9fd3a4b7316e66b8e6962f9395827c21c476dfb9e13f9f41938a740aeb22"></a>

## api_protection_rules.api_groups_rules.action.deny — api_protection_rules.api_groups_rules.action.deny / 18070af1832c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [api_protection_rules.api_groups_rules.action](data-sources--http_loadbalancer--reference--group-005.md#canonical-c86ede9ac2831bab5624eb5cbb9df404ab10aaa5682a9152f649c64b8935093f)
- api_protection_rules.api_groups_rules.action.deny

<a id="canonical-4267233bfeb859e5c4fc9bc0e42021eded12e154252e41137204bad3cb59a133"></a>

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

<a id="canonical-20f98836d683330f91d4997daa62c2fa266b66db3c763ab9cb70f95ccf432705"></a>

## Direct properties — api_protection_rules.api_groups_rules.action.deny / 18070af1832c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-19633bd39c7d2e878a28b5fb6dbb3c4a9a7bc0de60bf0524bfc6afb35f013292"></a>

## Next pages — api_protection_rules.api_groups_rules.action.deny / 18070af1832c / 4

- [api_protection_rules.api_groups_rules.action](data-sources--http_loadbalancer--reference--group-005.md#canonical-c86ede9ac2831bab5624eb5cbb9df404ab10aaa5682a9152f649c64b8935093f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6645ce780c54e522c879903ec8481be4f7205eafc370d3b6fca2b83a79061572"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0edda97363a007f49cd285638d6da334c0a4102c427de4e7ef9f812b7a5a1262"></a>

## api_protection_rules.api_groups_rules.any_domain — api_protection_rules.api_groups_rules.any_domain / 2d3c709d22c4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- api_protection_rules.api_groups_rules.any_domain

<a id="canonical-8b75ddcdc6cffdfa4bbb98b923f3e2d95e0140550b0a77f3ddb6f8cf925e6c2a"></a>

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

<a id="canonical-a2f6d86c445e25f6da1707eb942de324818c26c21217bf3255a58a968569896d"></a>

## Direct properties — api_protection_rules.api_groups_rules.any_domain / 2d3c709d22c4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-327ccc7631077418ba3437a10162c08b5c8344a5203369afd05709d2139fb0b6"></a>

## Next pages — api_protection_rules.api_groups_rules.any_domain / 2d3c709d22c4 / 4

- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-32a937c93dd23dd269c99657114ebddca1bb524608724bbebd46e3a69fad6e5a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9e6c948dd22d0e1179d91021e4d23fbd1c210075ff02b78b7896f413218dfde"></a>

## api_protection_rules.api_groups_rules.client_matcher — api_protection_rules.api_groups_rules.client_matcher / 48f5643c6d28 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- api_protection_rules.api_groups_rules.client_matcher

<a id="canonical-3e549fac58790010046887a111405d3b2df679697a9e6f8489dc9e4902870408"></a>

Type: `"single"`. Computed.

Client Matcher. Client conditions for matching a rule.

Upstream description:

Client conditions for matching a rule.

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

<a id="canonical-9d85664b788fd8db5cf7befce70f3afb0fb52a0058e9d97b09aecf2454e1a2c7"></a>

## Direct properties — api_protection_rules.api_groups_rules.client_matcher / 48f5643c6d28 / 3

- [any_client](data-sources--http_loadbalancer--reference--group-005.md#canonical-b129a3086f6cc6ec0a3e8e667377985539617f42577893b02f03e47dc645a6be): complete subsection reference.

- [any_ip](data-sources--http_loadbalancer--reference--group-005.md#canonical-6d3d173a96d65490c3a124c7a614b10a7835343786c92e6202b624d22ab21a4d): complete subsection reference.

- [asn_list](data-sources--http_loadbalancer--reference--group-005.md#canonical-0c221f8a0e40ef5233078e92f0f317077b7ff3a5de5cbdd1c9e60d545932a105): complete subsection reference.

- [asn_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-38f3806f5c2e0f1caaa4872f9e814b9b33a474ff26f2cd8dc10979a6e300b5a5): complete subsection reference.

- [client_selector](data-sources--http_loadbalancer--reference--group-005.md#canonical-4585982c4e07f268f7f7ff46f8a4c3ae626282f4e2d46621ec18d2591a988147): complete subsection reference.

- [ip_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-ec4fcfcffa1906e04f09be62c93527d61193b7f6fc8a01677f13c3513f034208): complete subsection reference.

- [ip_prefix_list](data-sources--http_loadbalancer--reference--group-005.md#canonical-75b43b6c8316b18cca80bb31b4a60dccc612857bfd06ca221149d1e4b3a9ddb1): complete subsection reference.

- [ip_threat_category_list](data-sources--http_loadbalancer--reference--group-005.md#canonical-664ae829c9536b672e8511b35de40027ea8bbca1422baf0f7c56e4bb9578d5d4): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-9adab7096d71f54d4da624e2e6f6b91fe528e52dbd4ee26e4b0d3a454cda4978): complete subsection reference.

<a id="canonical-896431ac13a84d2e435cc390041017a4334a5c2fb959e40445ad0a6405ed09a9"></a>

## Next pages — api_protection_rules.api_groups_rules.client_matcher / 48f5643c6d28 / 4

- [api_protection_rules.api_groups_rules.client_matcher.any_client](data-sources--http_loadbalancer--reference--group-005.md#canonical-b129a3086f6cc6ec0a3e8e667377985539617f42577893b02f03e47dc645a6be)
- [api_protection_rules.api_groups_rules.client_matcher.any_ip](data-sources--http_loadbalancer--reference--group-005.md#canonical-6d3d173a96d65490c3a124c7a614b10a7835343786c92e6202b624d22ab21a4d)
- [api_protection_rules.api_groups_rules.client_matcher.asn_list](data-sources--http_loadbalancer--reference--group-005.md#canonical-0c221f8a0e40ef5233078e92f0f317077b7ff3a5de5cbdd1c9e60d545932a105)
- [api_protection_rules.api_groups_rules.client_matcher.asn_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-38f3806f5c2e0f1caaa4872f9e814b9b33a474ff26f2cd8dc10979a6e300b5a5)
- [api_protection_rules.api_groups_rules.client_matcher.client_selector](data-sources--http_loadbalancer--reference--group-005.md#canonical-4585982c4e07f268f7f7ff46f8a4c3ae626282f4e2d46621ec18d2591a988147)
- [api_protection_rules.api_groups_rules.client_matcher.ip_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-ec4fcfcffa1906e04f09be62c93527d61193b7f6fc8a01677f13c3513f034208)
- [api_protection_rules.api_groups_rules.client_matcher.ip_prefix_list](data-sources--http_loadbalancer--reference--group-005.md#canonical-75b43b6c8316b18cca80bb31b4a60dccc612857bfd06ca221149d1e4b3a9ddb1)
- [api_protection_rules.api_groups_rules.client_matcher.ip_threat_category_list](data-sources--http_loadbalancer--reference--group-005.md#canonical-664ae829c9536b672e8511b35de40027ea8bbca1422baf0f7c56e4bb9578d5d4)
- [api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-9adab7096d71f54d4da624e2e6f6b91fe528e52dbd4ee26e4b0d3a454cda4978)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-b129a3086f6cc6ec0a3e8e667377985539617f42577893b02f03e47dc645a6be"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4fdb52a2fb533f7720abc985408f26fbf340ef8d582c926527cb7ac18036ec61"></a>

## api_protection_rules.api_groups_rules.client_matcher.any_client — api_protection_rules.api_groups_rules.client_matcher.any_client / b582a32f9a41 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-32a937c93dd23dd269c99657114ebddca1bb524608724bbebd46e3a69fad6e5a)
- api_protection_rules.api_groups_rules.client_matcher.any_client

<a id="canonical-bf0a9bd60f7d8e098c143ff1e52e435ac2a17f85f11d972eb90713128e24ee49"></a>

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

<a id="canonical-8eeed360a07130f5d08711df18f272f1f8e15fd3329c1c5fab5c6037f975a8e6"></a>

## Direct properties — api_protection_rules.api_groups_rules.client_matcher.any_client / b582a32f9a41 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2228234a12e6f4e31d8faec42a8b3a86d864d650088ea8fb0b308b04e01077b7"></a>

## Next pages — api_protection_rules.api_groups_rules.client_matcher.any_client / b582a32f9a41 / 4

- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-32a937c93dd23dd269c99657114ebddca1bb524608724bbebd46e3a69fad6e5a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6d3d173a96d65490c3a124c7a614b10a7835343786c92e6202b624d22ab21a4d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-657df175858daaac23e2f737982b17b921fc19eac1b57274c101df6b1f87564f"></a>

## api_protection_rules.api_groups_rules.client_matcher.any_ip — api_protection_rules.api_groups_rules.client_matcher.any_ip / a915858024c1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-32a937c93dd23dd269c99657114ebddca1bb524608724bbebd46e3a69fad6e5a)
- api_protection_rules.api_groups_rules.client_matcher.any_ip

<a id="canonical-058306c82254d75d817e5c1b0a0fdba34b685952adb820abcbedcbbb3b0ca311"></a>

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

<a id="canonical-a29a42f19bab9a6bc9d0cb9cec6d9d7d3b640fb9a6f53be8e01bb49a3c2d4289"></a>

## Direct properties — api_protection_rules.api_groups_rules.client_matcher.any_ip / a915858024c1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1bddc75d204a64e4e90795f9bd3f41d66d376eb7f307029d2cce891b25f5b321"></a>

## Next pages — api_protection_rules.api_groups_rules.client_matcher.any_ip / a915858024c1 / 4

- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-32a937c93dd23dd269c99657114ebddca1bb524608724bbebd46e3a69fad6e5a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-0c221f8a0e40ef5233078e92f0f317077b7ff3a5de5cbdd1c9e60d545932a105"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-973c90e5d7206a4a3e452e53be8c09a4f3fe1ce97acf9a63e73fbb3d5705fb28"></a>

## api_protection_rules.api_groups_rules.client_matcher.asn_list — api_protection_rules.api_groups_rules.client_matcher.asn_list / 377ba1f1d657 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-32a937c93dd23dd269c99657114ebddca1bb524608724bbebd46e3a69fad6e5a)
- api_protection_rules.api_groups_rules.client_matcher.asn_list

<a id="canonical-7024bbdc4c748d9aa1f3f532b50e86c489ab72266211a735f72229d2bc597b50"></a>

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

<a id="canonical-4afd5e066ed6d33b7d74b1a9fa825f3247d4756eaf79d68aaf779a2f4c4b1632"></a>

## Direct properties — api_protection_rules.api_groups_rules.client_matcher.asn_list / 377ba1f1d657 / 3

<a id="canonical-c03270b12323c44a438a7b567cafdb8c25deb7bb9077282ccda834d68669d6f6"></a>

<a id="canonical-85bdb7f722a7954445bf30b4cdfa8315909f8d601e9f8dbbe92a23fc66cbe53d"></a>

## as_numbers property — api_protection_rules.api_groups_rules.client_matcher.asn_list / 377ba1f1d657 / 4

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

<a id="canonical-79c5e59dcea9c33545401daab268ef5be279864eb109b8d8a6ee6948c546d1ae"></a>

## Next pages — api_protection_rules.api_groups_rules.client_matcher.asn_list / 377ba1f1d657 / 5

- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-32a937c93dd23dd269c99657114ebddca1bb524608724bbebd46e3a69fad6e5a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-38f3806f5c2e0f1caaa4872f9e814b9b33a474ff26f2cd8dc10979a6e300b5a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad23a2c1913f30955d9d226dcbf046012f3780beca21370af7d9c77d07187266"></a>

## api_protection_rules.api_groups_rules.client_matcher.asn_matcher — api_protection_rules.api_groups_rules.client_matcher.asn_matcher / c85fe55c7145 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-32a937c93dd23dd269c99657114ebddca1bb524608724bbebd46e3a69fad6e5a)
- api_protection_rules.api_groups_rules.client_matcher.asn_matcher

<a id="canonical-0de14e4fbc7d18e22161bc0cdff864ec774088b5578dc116856e4361fa6ca03c"></a>

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

<a id="canonical-3a92625ee1b67cdaa46c025bc18ed1c499c6c2c728cedd5301a72672020868d2"></a>

## Direct properties — api_protection_rules.api_groups_rules.client_matcher.asn_matcher / c85fe55c7145 / 3

- [asn_sets](data-sources--http_loadbalancer--reference--group-005.md#canonical-d0b3122956415ef389eaee65b3d83ee6e51c7d66e55066ace16f0dbec9fde502): complete subsection reference.

<a id="canonical-f15ee0bc2774b8d126cb766c1d5fcf091c087834a4229d367f09ba68021d4715"></a>

## Next pages — api_protection_rules.api_groups_rules.client_matcher.asn_matcher / c85fe55c7145 / 4

- [api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets](data-sources--http_loadbalancer--reference--group-005.md#canonical-d0b3122956415ef389eaee65b3d83ee6e51c7d66e55066ace16f0dbec9fde502)
- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-32a937c93dd23dd269c99657114ebddca1bb524608724bbebd46e3a69fad6e5a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d0b3122956415ef389eaee65b3d83ee6e51c7d66e55066ace16f0dbec9fde502"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc76b1c90be259feff96644f962b02c023f719b018c73808719b627467235fcd"></a>

## api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets — api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets / b6ca8bc6c407 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-32a937c93dd23dd269c99657114ebddca1bb524608724bbebd46e3a69fad6e5a)
- [api_protection_rules.api_groups_rules.client_matcher.asn_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-38f3806f5c2e0f1caaa4872f9e814b9b33a474ff26f2cd8dc10979a6e300b5a5)
- api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-7629011f9c242b951b5f3636003e2d9aa977420c312efd5bbddfeee722c210ac"></a>

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

<a id="canonical-fde62d61901cf3af989cce4f80b5853ea50ee29949cf4e706476de59c00401fb"></a>

## Direct properties — api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets / b6ca8bc6c407 / 3

<a id="canonical-b2f28e89903d40945fe9912823c2b83fbd22c8596d43c626b7501b49f66354a0"></a>

<a id="canonical-3fe1752e9cbdc66bfd29ab267c07bd30f4c44b0b8689db03d5a24a23b89abf9e"></a>

## kind property — api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets / b6ca8bc6c407 / 4

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

<a id="canonical-fe61732f91b6b4f8fa2b492805a61e9b82c8b516d4fd0cffe49283a4bc9e4c1f"></a>

<a id="canonical-e83c62ec2ec4b0bb483f2586728c3dffd31f102640e1f9eb2fcc5154f805e212"></a>

## name property — api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets / b6ca8bc6c407 / 5

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

<a id="canonical-d09c1dbb18ecbb327f38a77e99fea706744ef3d25a2c0740903817d588fdac9d"></a>

<a id="canonical-edc22526f98a608d9cf8288b30e49e62aa66ac4aba73552b44eec52a4f3e6c7a"></a>

## namespace property — api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets / b6ca8bc6c407 / 6

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

<a id="canonical-2c3aba4e46a8120d90bd224e3bdc06ce3f260c733ef70d78522073da09ca76a8"></a>

<a id="canonical-ca3147e6f51c920b2e394c28ac94fb456beebe027d61e12682f13128de188194"></a>

## tenant property — api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets / b6ca8bc6c407 / 7

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

<a id="canonical-e2bd71282e2b5241b91ce9f4613edf367a8f1f5c5590595ed14382828641b984"></a>

<a id="canonical-8fdc530b6eb4128683b67117d7676eba3b9bf8897dfacf096a753b723686ce3a"></a>

## uid property — api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets / b6ca8bc6c407 / 8

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

<a id="canonical-03abe1ea6c5b6c81f29132906374937607ee3d8467336fab155c6f42bd448dff"></a>

## Next pages — api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets / b6ca8bc6c407 / 9

- [api_protection_rules.api_groups_rules.client_matcher.asn_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-38f3806f5c2e0f1caaa4872f9e814b9b33a474ff26f2cd8dc10979a6e300b5a5)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-4585982c4e07f268f7f7ff46f8a4c3ae626282f4e2d46621ec18d2591a988147"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4dc8697fa68a60fe836062a92820313a004ee79d7c1497b4991ab6a788deb03a"></a>

## api_protection_rules.api_groups_rules.client_matcher.client_selector — api_protection_rules.api_groups_rules.client_matcher.client_selector / 575a6e63d384 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-32a937c93dd23dd269c99657114ebddca1bb524608724bbebd46e3a69fad6e5a)
- api_protection_rules.api_groups_rules.client_matcher.client_selector

<a id="canonical-34e75488a4bd130a4ae9b015c71622b5f0ec75d58972f194d29f6f2de74e9c46"></a>

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

<a id="canonical-e0e0e0b86264c37a268f74a024fb1158973f26458fd65a85b18acb8427c8d54e"></a>

## Direct properties — api_protection_rules.api_groups_rules.client_matcher.client_selector / 575a6e63d384 / 3

<a id="canonical-7badb0d4113546ad71caf7fea84fcec84db2754e1e24198fca24a3f27db59189"></a>

<a id="canonical-6c2025ba33d595dc8e0745663abbc3b278b6cce00fcbf900ba22d7b57bdbbd69"></a>

## expressions property — api_protection_rules.api_groups_rules.client_matcher.client_selector / 575a6e63d384 / 4

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

<a id="canonical-620031c11c2317bba101cfeb3f4eec5a8eb6d58ff05896c923fbad5b97365338"></a>

## Next pages — api_protection_rules.api_groups_rules.client_matcher.client_selector / 575a6e63d384 / 5

- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-32a937c93dd23dd269c99657114ebddca1bb524608724bbebd46e3a69fad6e5a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ec4fcfcffa1906e04f09be62c93527d61193b7f6fc8a01677f13c3513f034208"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-54f8ac86d3de3632379b83a7244b6662ff7afe9fa2f9365bdc33c360ad4d98c7"></a>

## api_protection_rules.api_groups_rules.client_matcher.ip_matcher — api_protection_rules.api_groups_rules.client_matcher.ip_matcher / 430a886184db / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-32a937c93dd23dd269c99657114ebddca1bb524608724bbebd46e3a69fad6e5a)
- api_protection_rules.api_groups_rules.client_matcher.ip_matcher

<a id="canonical-77c06b8a1cbb55cb6a4af6d1cd93c9a1154bf187cf3aec98cbcde8c6faec1ced"></a>

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

<a id="canonical-f00926dc865a501e147aa97246cf29a95df8c56b2b6a5e13ab156620343239b1"></a>

## Direct properties — api_protection_rules.api_groups_rules.client_matcher.ip_matcher / 430a886184db / 3

<a id="canonical-5930950a850c6d28562e7ff2649dde719ef620a9eafa9e250b43c364c8a7559e"></a>

<a id="canonical-d3a0f091c34310f99230580f64f0a6650a1d6a574ae5310ad341e53543032e8f"></a>

## invert_matcher property — api_protection_rules.api_groups_rules.client_matcher.ip_matcher / 430a886184db / 4

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

- [prefix_sets](data-sources--http_loadbalancer--reference--group-005.md#canonical-68b9148ea76a6dd77d034d4f5e7b3559e83de674b1884a70cc8383ee97e2b0f1): complete subsection reference.

<a id="canonical-37fe10da5bc264c09107a2af0c0c693cbf28d0ee08d24ae3cf85faf86314a54a"></a>

## Next pages — api_protection_rules.api_groups_rules.client_matcher.ip_matcher / 430a886184db / 5

- [api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets](data-sources--http_loadbalancer--reference--group-005.md#canonical-68b9148ea76a6dd77d034d4f5e7b3559e83de674b1884a70cc8383ee97e2b0f1)
- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-32a937c93dd23dd269c99657114ebddca1bb524608724bbebd46e3a69fad6e5a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-68b9148ea76a6dd77d034d4f5e7b3559e83de674b1884a70cc8383ee97e2b0f1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1cc91a7b0f962b5446f80e99dee4ff18fc5c0904a65c9c12a7ee0a2885e61f0"></a>

## api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets — api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets / 55bcc44c1574 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-32a937c93dd23dd269c99657114ebddca1bb524608724bbebd46e3a69fad6e5a)
- [api_protection_rules.api_groups_rules.client_matcher.ip_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-ec4fcfcffa1906e04f09be62c93527d61193b7f6fc8a01677f13c3513f034208)
- api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-504d73e279106c3fe8ef1d087ad6ca406298570284d985baee6b81dcab749df6"></a>

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

<a id="canonical-e6116cf34cc5c239e0ae503720f411c080feb38bb0079d2efca4522a416b8b91"></a>

## Direct properties — api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets / 55bcc44c1574 / 3

<a id="canonical-90f491abd848ea57e6cbec7a9b4a3afa8cda527f8adb1b108499879c81c95b12"></a>

<a id="canonical-476f098b98c50819949c8858c9315b4c8195ef65f4eca90948a4b6d70f6155e0"></a>

## kind property — api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets / 55bcc44c1574 / 4

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

<a id="canonical-0c929f95cff95a92bfd40f2670e921434dffa99e0d4057c0d1d061a7cae4ec27"></a>

<a id="canonical-416b87edb1265d868deaddc895b8c1ef83c43433c029353d36bdfa0d8c153e19"></a>

## name property — api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets / 55bcc44c1574 / 5

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

<a id="canonical-eb7404cb5377ac7d5d7135314801bd8cf03193c0d69efab10556040ba30083a4"></a>

<a id="canonical-49a7376b5cdf93f1cda683a3a1f0025916f57f316fcac13e2e1cd2353868dc05"></a>

## namespace property — api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets / 55bcc44c1574 / 6

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

<a id="canonical-b1ad4f2448bab4183e55d8b934fa65ad50174e669f98eaacff25f46127e188df"></a>

<a id="canonical-c96799702830532499b82019bca96bbba3c649669f64a18e2f90adcfa4505551"></a>

## tenant property — api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets / 55bcc44c1574 / 7

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

<a id="canonical-4afa49055b9f957ed88644afd6ce900fc184d7cf86ff7eb30f0fd518217df4be"></a>

<a id="canonical-4ec1e66c18cedd6c40296e3543e818b773c6380b306a0f6afe72b93ac8127d80"></a>

## uid property — api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets / 55bcc44c1574 / 8

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

<a id="canonical-199da241e1b451c3c1f4186ce814f6e0200cfc0add4a9953f91d773aaedabce4"></a>

## Next pages — api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets / 55bcc44c1574 / 9

- [api_protection_rules.api_groups_rules.client_matcher.ip_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-ec4fcfcffa1906e04f09be62c93527d61193b7f6fc8a01677f13c3513f034208)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-75b43b6c8316b18cca80bb31b4a60dccc612857bfd06ca221149d1e4b3a9ddb1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fdcb4a064048cad5114040a85d2d284c653a9022abb08c2474c7104239ccef7b"></a>

## api_protection_rules.api_groups_rules.client_matcher.ip_prefix_list — api_protection_rules.api_groups_rules.client_matcher.ip_prefix_list / 8c89dbfb1b07 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-32a937c93dd23dd269c99657114ebddca1bb524608724bbebd46e3a69fad6e5a)
- api_protection_rules.api_groups_rules.client_matcher.ip_prefix_list

<a id="canonical-7d17e161837a77b0e0a4a04c5a965fbe1bd9dccd4b892088be3c2d90264e7607"></a>

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

<a id="canonical-918e5138813ed174a8bdc62a9f1daf73208f7ca23d00c670a40df1698ffaffa2"></a>

## Direct properties — api_protection_rules.api_groups_rules.client_matcher.ip_prefix_list / 8c89dbfb1b07 / 3

<a id="canonical-6354da444714eff874c8ff0a5e970cf4451f377787d3a3db72205e5396f10857"></a>

<a id="canonical-0874ff7b1c8efac217d9761cafea60236f62a63d7bd0fdd21c0d9ec3f4ee0187"></a>

## invert_match property — api_protection_rules.api_groups_rules.client_matcher.ip_prefix_list / 8c89dbfb1b07 / 4

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

<a id="canonical-248f6236e78c1120f243c329e945e3d8a9810cf746819fd84bba1d54dcf3d068"></a>

<a id="canonical-bee3ec7cd5f28150b105097f150fd9f4982885e620fbffe7d304320a708a7a69"></a>

## ip_prefixes property — api_protection_rules.api_groups_rules.client_matcher.ip_prefix_list / 8c89dbfb1b07 / 5

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

<a id="canonical-4697f5aef2e02509b404797979edd1f625c35c850c792047150dc9d237d988fa"></a>

## Next pages — api_protection_rules.api_groups_rules.client_matcher.ip_prefix_list / 8c89dbfb1b07 / 6

- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-32a937c93dd23dd269c99657114ebddca1bb524608724bbebd46e3a69fad6e5a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-664ae829c9536b672e8511b35de40027ea8bbca1422baf0f7c56e4bb9578d5d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a374b3ced598ff0d209ee5cd7b998c5211d3baf13dd596c46b20194c58d5f2f"></a>

## api_protection_rules.api_groups_rules.client_matcher.ip_threat_category_list — api_protection_rules.api_groups_rules.client_matcher.ip_threat_category_list / f1a7944d3d3e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-32a937c93dd23dd269c99657114ebddca1bb524608724bbebd46e3a69fad6e5a)
- api_protection_rules.api_groups_rules.client_matcher.ip_threat_category_list

<a id="canonical-26299d13343a8a6d6b26fc9b6b5100e47e7ef8d3f4abcc843462ee4e06d13553"></a>

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

<a id="canonical-20046967506e5297fa8e4c0827eb47b62bdf141bbffadb6dda97119717c6cfd9"></a>

## Direct properties — api_protection_rules.api_groups_rules.client_matcher.ip_threat_category_list / f1a7944d3d3e / 3

<a id="canonical-98d2651de19af051bec7ac8cea8b4fc0c9003f3c50fdb9660d7197e870be2488"></a>

<a id="canonical-54a1c0007d2f4a8c20dbf57d05050a5e3cbd5868869f2b8388e817371ea5fe48"></a>

## ip_threat_categories property — api_protection_rules.api_groups_rules.client_matcher.ip_threat_category_list / f1a7944d3d3e / 4

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

<a id="canonical-21a429b09a9234828a60725dfac087ee4d5e20d9c03a1b8853940fbe73bc424e"></a>

## Next pages — api_protection_rules.api_groups_rules.client_matcher.ip_threat_category_list / f1a7944d3d3e / 5

- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-32a937c93dd23dd269c99657114ebddca1bb524608724bbebd46e3a69fad6e5a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9adab7096d71f54d4da624e2e6f6b91fe528e52dbd4ee26e4b0d3a454cda4978"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c123423ee049af6ae89d132538c9d6c970c3b2051c7ab711faa0b3c6ecd6e0ed"></a>

## api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher — api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher / bb66ac0a985d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-32a937c93dd23dd269c99657114ebddca1bb524608724bbebd46e3a69fad6e5a)
- api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-4b6946bc4fe9a8bb1465ad1df64e844831948afa7e6f19035930e9608e2975e6"></a>

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

<a id="canonical-beae4520101950a208bfe6c4baab8426fcb56c14b35b8b3d53e1f55ffd3c0362"></a>

## Direct properties — api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher / bb66ac0a985d / 3

<a id="canonical-08bc2bb2a531017df37cb042d85207ae93a495b797f68826224de12b4c4ecefa"></a>

<a id="canonical-f7330749c1ee0d9cc1e3576a40f6c278579309e086bb336bfca3712a98e9aa37"></a>

## classes property — api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher / bb66ac0a985d / 4

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

<a id="canonical-0da2b4ef735fed09a77af5b6fd3cf1917389eee87f55b8cec2a8da3fad2318b3"></a>

<a id="canonical-b28ca72cd11cddd0bbded8bc849504399fd557d566e76f4a38758ad3ee3d96b0"></a>

## exact_values property — api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher / bb66ac0a985d / 5

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

<a id="canonical-f626ee963e7d3c664e12de7ac3631a8b01f66a469ba0f235a35b2ea09ac80943"></a>

<a id="canonical-6ad6a76bc66530f92d79ef1f87a8be73cea27b97c51b7a1eaff04c5b2ec7e431"></a>

## excluded_values property — api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher / bb66ac0a985d / 6

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

<a id="canonical-403727c6388800fc2f1009e7061efe58de8c47ed5f7ca32061699f207fdf22d2"></a>

## Next pages — api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher / bb66ac0a985d / 7

- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-32a937c93dd23dd269c99657114ebddca1bb524608724bbebd46e3a69fad6e5a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-8f9610f28d4a71e22bbc8a4bf122b9d4e075db843ed5e87344008d76f9808f3f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f824357b03df8d1384d5c496a007020cf36e1a2e5ed5c37123540d2227a77d6"></a>

## api_protection_rules.api_groups_rules.metadata — api_protection_rules.api_groups_rules.metadata / b35441c1af2b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- api_protection_rules.api_groups_rules.metadata

<a id="canonical-57bfea08e1de43ed44290f901bf3c86eb24a09281d13d7afe189d66dae633cb2"></a>

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

<a id="canonical-b281e7c9c610428ad2aff3d10f3386007a810cca2b438e00f7f053ce600a7537"></a>

## Direct properties — api_protection_rules.api_groups_rules.metadata / b35441c1af2b / 3

<a id="canonical-41a2645f607bdaf9ba3de599fa5bb7bb11bd2b912dd9f6d8c49e5060dbe7086b"></a>

<a id="canonical-9d007486e32902bd1566e719a8610ddfd9365d30de6d79a36542babf899568e1"></a>

## description_spec property — api_protection_rules.api_groups_rules.metadata / b35441c1af2b / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-f8fddd39414fbdc3c7a1e1792f1fe3e5c11fe8ca46b1ae19ba11f247ef64c3f4"></a>

<a id="canonical-62a7dcf53699632852df607a4010a47b376c4f635f0077ce72d65dc801e7e2d8"></a>

## name property — api_protection_rules.api_groups_rules.metadata / b35441c1af2b / 5

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

<a id="canonical-56ae221d81bf20c091b43a2fd5f17c9ba441b6f19cf8f8d664fa89d17bf8caa2"></a>

## Next pages — api_protection_rules.api_groups_rules.metadata / b35441c1af2b / 6

- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1799983b63e67772d3b71c4156c2cd18db14ba86436ab2971e1a87048d4ba462"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-36a06a54ac2b86fe6ce1237ec2d490e7718c80b07367ac22538f9856fe173715"></a>

## api_protection_rules.api_groups_rules.request_matcher — api_protection_rules.api_groups_rules.request_matcher / d2a8ab8750b1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- api_protection_rules.api_groups_rules.request_matcher

<a id="canonical-f9589c0e06c81876ed350465e8193cd1aee65eb42d4062cee3132a4907114285"></a>

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

<a id="canonical-5bc4c7c2b13d629e653cf90b1d71f640683d7b8893f7b43e41bfd36b29b0c4d5"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher / d2a8ab8750b1 / 3

- [cookie_matchers](data-sources--http_loadbalancer--reference--group-005.md#canonical-e35159d06bb46f012f8c93b49dcaac00cf5704f4556cc3e091ab22aeef5df1f2): complete subsection reference.

- [headers](data-sources--http_loadbalancer--reference--group-006.md#canonical-2e992deebbcd5dfb1b352c1da1094f5bc08dd4fc5e22fc8ce2faa2c5f01a50bd): complete subsection reference.

- [jwt_claims](data-sources--http_loadbalancer--reference--group-006.md#canonical-57e6f0503ee00e535beacca7e3fc592b9c63a2aea6f567472d36593f7ac3d5d0): complete subsection reference.

- [query_params](data-sources--http_loadbalancer--reference--group-006.md#canonical-56d69a7fed65a4971608d20e3406eea39bed0b6487892eb7ef89edd532e78f14): complete subsection reference.

<a id="canonical-0fa878f183c939b299c7cecca0df6e1b315f4b30a938b815f33a12dcb49d1e67"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher / d2a8ab8750b1 / 4

- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-005.md#canonical-e35159d06bb46f012f8c93b49dcaac00cf5704f4556cc3e091ab22aeef5df1f2)
- [api_protection_rules.api_groups_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-006.md#canonical-2e992deebbcd5dfb1b352c1da1094f5bc08dd4fc5e22fc8ce2faa2c5f01a50bd)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-006.md#canonical-57e6f0503ee00e535beacca7e3fc592b9c63a2aea6f567472d36593f7ac3d5d0)
- [api_protection_rules.api_groups_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-006.md#canonical-56d69a7fed65a4971608d20e3406eea39bed0b6487892eb7ef89edd532e78f14)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e35159d06bb46f012f8c93b49dcaac00cf5704f4556cc3e091ab22aeef5df1f2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eeb3a5333f3a9704f8e367850472e3cb2eec45dcfe0b9fb0d46c1a06a9a9be37"></a>

## api_protection_rules.api_groups_rules.request_matcher.cookie_matchers — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers / 479ec8fb738d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1799983b63e67772d3b71c4156c2cd18db14ba86436ab2971e1a87048d4ba462)
- api_protection_rules.api_groups_rules.request_matcher.cookie_matchers

<a id="canonical-1c18cd34072985ea73ad7ea14a63b9954e8cf0b5c824412dae7d6f6b47e01a2b"></a>

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

<a id="canonical-98eb88ccb026b4b16e92e6fa5570075fb2c05f41659c61fbf0333925c171068e"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers / 479ec8fb738d / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-005.md#canonical-52df23fa349bb2aeb0a2508d3ed20f4237feae972082876e5af0a5a4130557db): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-005.md#canonical-0ea0faef1fd74d16c69c5c9814575ee8b6198e5d9c899438767c3577363f22cb): complete subsection reference.

<a id="canonical-f736e4f3bce07066f00d2c7a32f12f110f293baa2fca1075a693d6f0b7e46b2b"></a>

<a id="canonical-75ceeb081f25e2d3999149156c0ef95c22a6e9a1faf00929ecd590e47f4d3968"></a>

## invert_matcher property — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers / 479ec8fb738d / 4

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

- [item](data-sources--http_loadbalancer--reference--group-006.md#canonical-b051e8a43c9083ef46f46183bcdd9f2a79a23f30ba4190210dc9e419a4febc81): complete subsection reference.

<a id="canonical-8e97bd803bc2cc882f84ea32c29139d1d96182766388c68e9cfe8fd0b8b18ff2"></a>

<a id="canonical-ad3f531bbe4825a619e363eff7a64edad120a987cca5cc63d1fc60d14697a868"></a>

## name property — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers / 479ec8fb738d / 5

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

<a id="canonical-f215a3e97a0d7368d188e24d92673e6f70441c6ae114d1f1d0d7486587d12631"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers / 479ec8fb738d / 6

- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_not_present](data-sources--http_loadbalancer--reference--group-005.md#canonical-52df23fa349bb2aeb0a2508d3ed20f4237feae972082876e5af0a5a4130557db)
- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_present](data-sources--http_loadbalancer--reference--group-005.md#canonical-0ea0faef1fd74d16c69c5c9814575ee8b6198e5d9c899438767c3577363f22cb)
- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item](data-sources--http_loadbalancer--reference--group-006.md#canonical-b051e8a43c9083ef46f46183bcdd9f2a79a23f30ba4190210dc9e419a4febc81)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1799983b63e67772d3b71c4156c2cd18db14ba86436ab2971e1a87048d4ba462)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-52df23fa349bb2aeb0a2508d3ed20f4237feae972082876e5af0a5a4130557db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e2d3beac04354d2bf3ad4b5d4edcc749aacc045d69cc540fa7c2f846ec138d3"></a>

## api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_not_present — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_not_ / c7441c453339 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1799983b63e67772d3b71c4156c2cd18db14ba86436ab2971e1a87048d4ba462)
- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-005.md#canonical-e35159d06bb46f012f8c93b49dcaac00cf5704f4556cc3e091ab22aeef5df1f2)
- api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-8f92a1c9eeca67400c3ec8a7ab315c1fccd7156b29239788febab98b91b90ac4"></a>

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

<a id="canonical-fd5fc176c81b7b22931011d002ae8f9da5b2a17b3c56f562b2a4312339e4c3f4"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_not_ / c7441c453339 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e27522ed132a31c86c069ac5514e2326a304298de30ef018f650cacf6e5a442f"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_not_ / c7441c453339 / 4

- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-005.md#canonical-e35159d06bb46f012f8c93b49dcaac00cf5704f4556cc3e091ab22aeef5df1f2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-0ea0faef1fd74d16c69c5c9814575ee8b6198e5d9c899438767c3577363f22cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac7f7dc13e57a8ee8362b8a933b84a8edc73b9b87e3310f5c8ef92c9331e5698"></a>

## api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_present — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_pres / 5895f2ed1fea / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1799983b63e67772d3b71c4156c2cd18db14ba86436ab2971e1a87048d4ba462)
- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-005.md#canonical-e35159d06bb46f012f8c93b49dcaac00cf5704f4556cc3e091ab22aeef5df1f2)
- api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-3d8ba15c26224ca9263b9c6e041c7ebc1a5ba3116979924c486f2c743e1906cd"></a>

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
