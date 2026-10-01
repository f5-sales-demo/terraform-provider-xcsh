---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-22ddb5f44f8d621da443ad88a70b88fc5b811d226e9111a671e38af11877b89b"></a>

## bot_defense.policy.js_insert_all_pages_except — bot_defense.policy.js_insert_all_pages_except / 258daf6a98be / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- bot_defense.policy.js_insert_all_pages_except

<a id="canonical-1c872157f0defa293689c5810fca994e545e56f325ec462301180634208cdf16"></a>

Type: `"object"`. single nested block, Optional.

Insert Bot Defense JavaScript in all pages with the exceptions.

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
js_insert_all_pages_except {
  # Configure direct properties listed below.
}
```

<a id="canonical-19f49223d6dc002394525aa22bd6d35088088202ecea53087ae3f13b4954ff66"></a>

## Direct properties — bot_defense.policy.js_insert_all_pages_except / 258daf6a98be / 3

- [exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-289048303ff86986cc8bb3c805204fbfca35667135ea9b579ead5bde76f8543c): complete subsection reference.

<a id="canonical-cb07191612f7999b5e119f339e5c4c68f7b755a3a094971c5cc501f5bba953fe"></a>

<a id="canonical-f60baab9033aa8640b65487d16b2a8fd2944f11082707bcd9527357ecfb649c8"></a>

## javascript_location property — bot_defense.policy.js_insert_all_pages_except / 258daf6a98be / 4

Type: `"string"`. Optional.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-ce5947bd628b0da57fd4eb5cf42aefd7a3aebe0afded6fb934687bb175f41795"></a>

## Next pages — bot_defense.policy.js_insert_all_pages_except / 258daf6a98be / 5

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-289048303ff86986cc8bb3c805204fbfca35667135ea9b579ead5bde76f8543c)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-289048303ff86986cc8bb3c805204fbfca35667135ea9b579ead5bde76f8543c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a18b34e70a38f78e124bf70a79464556024f497140baff837508b8479972c01"></a>

## bot_defense.policy.js_insert_all_pages_except.exclude_list — bot_defense.policy.js_insert_all_pages_except.exclude_list / 34a3858bc14f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-010.md#canonical-035b1a4ae3260f37d29db3939a2013a1b761d7484d73425eef373a4961cb85e5)
- bot_defense.policy.js_insert_all_pages_except.exclude_list

<a id="canonical-7afa33fc2dfac9b83ac090e727210536a07e3c18d563a9343cc89a734e2548c6"></a>

Type: `"object"`. list nested block, Optional.

Optional JavaScript insertions exclude list of domain and path matchers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-dc384e9e963a7cc47f898466811c489410c6fe89ec31a26f5fd4bbf64e141232"></a>

## Direct properties — bot_defense.policy.js_insert_all_pages_except.exclude_list / 34a3858bc14f / 3

- [any_domain](resources--http_loadbalancer--reference--group-011.md#canonical-07f24c13635ba9ee8397845e481c60e967b8bf4818f935cc4d6af67eac17194a): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-011.md#canonical-76b8177d591ce1e869869f3bb3445fd7a60817e2590bec9adb682ad59aacda75): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-011.md#canonical-b51d3340cf418337f0d87700418f212d1284ee5d3929cb99e993a791ff3cba57): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-011.md#canonical-2f49f52ecca4bd13e26f38f9487e1f1634825a735992c6666820e5b7276dc9f5): complete subsection reference.

<a id="canonical-5e0c4fcf7ee621227ce9f3f224cc70c07d559c7b3877838518558ec938a9929c"></a>

## Next pages — bot_defense.policy.js_insert_all_pages_except.exclude_list / 34a3858bc14f / 4

- [bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain](resources--http_loadbalancer--reference--group-011.md#canonical-07f24c13635ba9ee8397845e481c60e967b8bf4818f935cc4d6af67eac17194a)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list.domain](resources--http_loadbalancer--reference--group-011.md#canonical-76b8177d591ce1e869869f3bb3445fd7a60817e2590bec9adb682ad59aacda75)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata](resources--http_loadbalancer--reference--group-011.md#canonical-b51d3340cf418337f0d87700418f212d1284ee5d3929cb99e993a791ff3cba57)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list.path](resources--http_loadbalancer--reference--group-011.md#canonical-2f49f52ecca4bd13e26f38f9487e1f1634825a735992c6666820e5b7276dc9f5)
- [bot_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-010.md#canonical-035b1a4ae3260f37d29db3939a2013a1b761d7484d73425eef373a4961cb85e5)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-07f24c13635ba9ee8397845e481c60e967b8bf4818f935cc4d6af67eac17194a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b73c0063a37fa2c2e2d2245abac0a7b3d3f55c5cff12fe6e9584c184ae106c59"></a>

## bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain — bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain / 672f132a6c93 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-010.md#canonical-035b1a4ae3260f37d29db3939a2013a1b761d7484d73425eef373a4961cb85e5)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-289048303ff86986cc8bb3c805204fbfca35667135ea9b579ead5bde76f8543c)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain

<a id="canonical-018de9cee03f634a7d7bae6e6bbbf0e3cc01167f9ee3b495d3187f898835621c"></a>

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

<a id="canonical-99c4785c83c2f125e7acecaf9dd4305e25777e818ae133c4072f9021583c679c"></a>

## Direct properties — bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain / 672f132a6c93 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d0fda2fcc6776e81addc4200088a3b79cf40b98fb85626ca079e45ebf4ee36b0"></a>

## Next pages — bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain / 672f132a6c93 / 4

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-289048303ff86986cc8bb3c805204fbfca35667135ea9b579ead5bde76f8543c)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-76b8177d591ce1e869869f3bb3445fd7a60817e2590bec9adb682ad59aacda75"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e8c5e7810fc1b74132c4aa3d039a753bdb4d75a6046c6b02e772720518929c3c"></a>

## bot_defense.policy.js_insert_all_pages_except.exclude_list.domain — bot_defense.policy.js_insert_all_pages_except.exclude_list.domain / 01c80ba67a74 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-010.md#canonical-035b1a4ae3260f37d29db3939a2013a1b761d7484d73425eef373a4961cb85e5)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-289048303ff86986cc8bb3c805204fbfca35667135ea9b579ead5bde76f8543c)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.domain

<a id="canonical-25681aaca0373d48e120fad6873e7c1d1454331401bb5eca8d906d572e4ed4da"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-a911cc1a11c9728c73ff825d794c4c90679d621a508e440644ef355547a8d59b"></a>

## Direct properties — bot_defense.policy.js_insert_all_pages_except.exclude_list.domain / 01c80ba67a74 / 3

<a id="canonical-b2b7f5397728a2f90f5f1d96e36d025002422f3c7d26337acfd09137b09758d8"></a>

<a id="canonical-e60cbb08867ad6dd1040aed31d8ae0b7385c4d7a18bfe7d12030d9917ac2b38f"></a>

## exact_value property — bot_defense.policy.js_insert_all_pages_except.exclude_list.domain / 01c80ba67a74 / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-4798de9fa948242b012a1a9ecc14fa09efb716673d522a37eb155856a5f47dc8"></a>

<a id="canonical-2f38aceb9c98e525ebf0eea7883b6c0a1ccf6d19c603549a014c420a92c4cf84"></a>

## regex_value property — bot_defense.policy.js_insert_all_pages_except.exclude_list.domain / 01c80ba67a74 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-671662ba268b9e75401510760c2af1ba5484ea5256b81a48d612f13b3720b8cb"></a>

<a id="canonical-8f3c41eac20053d55d9c2062c82f7e1475f99bad2da98f5706633088ff1401b1"></a>

## suffix_value property — bot_defense.policy.js_insert_all_pages_except.exclude_list.domain / 01c80ba67a74 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-a4a2f35890a1717bb2dfd572b3d8d3a1e1b33b759b228a84eac6900b0f70c574"></a>

## Next pages — bot_defense.policy.js_insert_all_pages_except.exclude_list.domain / 01c80ba67a74 / 7

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-289048303ff86986cc8bb3c805204fbfca35667135ea9b579ead5bde76f8543c)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b51d3340cf418337f0d87700418f212d1284ee5d3929cb99e993a791ff3cba57"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2593ac44d493d7c5f52013f22bb1d5fe7c55d7da8c022472cc7439e26587e8c2"></a>

## bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata — bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata / ed9dc7452899 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-010.md#canonical-035b1a4ae3260f37d29db3939a2013a1b761d7484d73425eef373a4961cb85e5)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-289048303ff86986cc8bb3c805204fbfca35667135ea9b579ead5bde76f8543c)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata

<a id="canonical-bb72dc72505fd18c28ee4f4d57ff419e7b8c551e1394628d79f9f3a551f8fd26"></a>

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

<a id="canonical-a8892851c2683221b7737e0e008b17ea7af6849a39364a0e73272a9965802b5d"></a>

## Direct properties — bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata / ed9dc7452899 / 3

<a id="canonical-91392486adab6b74c5aa0b14d5eb029daacbd0836f0c30e91b36d004b8b4eff8"></a>

<a id="canonical-b4093197b73e27cfd72a81fcb8e4ee622fe59f54be72d2e4b59de949b5d97909"></a>

## description_spec property — bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata / ed9dc7452899 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-8928617dbab8525ea1b4ea7f27079d2c7501df61b7f4412fd8417c5bfa3d103a"></a>

<a id="canonical-9566d3f6f8426fed3dc5264dbda81d98909afada38345a50503ceab74f1bf4f2"></a>

## name property — bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata / ed9dc7452899 / 5

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

<a id="canonical-4abab9df9530ba0816ef0bf262f5c8410a85310035dd7e00a9e5216ff2cfcfc1"></a>

## Next pages — bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata / ed9dc7452899 / 6

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-289048303ff86986cc8bb3c805204fbfca35667135ea9b579ead5bde76f8543c)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-2f49f52ecca4bd13e26f38f9487e1f1634825a735992c6666820e5b7276dc9f5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f4bfd164c1be225554ff753a9f759d8068b9976297807266a4d2fb6fd3e16cd2"></a>

## bot_defense.policy.js_insert_all_pages_except.exclude_list.path — bot_defense.policy.js_insert_all_pages_except.exclude_list.path / 48ad6f0e276f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-010.md#canonical-035b1a4ae3260f37d29db3939a2013a1b761d7484d73425eef373a4961cb85e5)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-289048303ff86986cc8bb3c805204fbfca35667135ea9b579ead5bde76f8543c)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.path

<a id="canonical-0e3e2e422498d5627d3c33e57c2f2f849a12d440da19ce44bc871b53e4eed98f"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-a157a666a5910e4d3bf4e331249879ee7a13801e06d3e0ae44ecafcfa4e27dc2"></a>

## Direct properties — bot_defense.policy.js_insert_all_pages_except.exclude_list.path / 48ad6f0e276f / 3

<a id="canonical-17b54e1df9caae96b171fea3b4fde5ca74a46c955ebacaea3c8f663f95125bfa"></a>

<a id="canonical-ddaa184abc580cc4c06836efaa3e550882a25beaae0895da7bc08c931af94f18"></a>

## path property — bot_defense.policy.js_insert_all_pages_except.exclude_list.path / 48ad6f0e276f / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3d89c1cc40dcec01d4533919dc1b5d03dedad8db76d8c8e52e0262c439dddf99"></a>

<a id="canonical-48107ad0579ef373361af730042fdd2bb2715e00c40f34a3d9afef6efc4f9c30"></a>

## prefix property — bot_defense.policy.js_insert_all_pages_except.exclude_list.path / 48ad6f0e276f / 5

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-27f5e497c292ec0ff1c90493f86521967991d6234daab91e6fbe54d504876db8"></a>

<a id="canonical-07c2c31f7803c68e63f46c3140668fd857c4b1cca63f963ea23bcb7c811bbb32"></a>

## regex property — bot_defense.policy.js_insert_all_pages_except.exclude_list.path / 48ad6f0e276f / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-4d27ebe423263869c57b1339de77fe384396b2919a48a62208fd16d034c0edc1"></a>

## Next pages — bot_defense.policy.js_insert_all_pages_except.exclude_list.path / 48ad6f0e276f / 7

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-289048303ff86986cc8bb3c805204fbfca35667135ea9b579ead5bde76f8543c)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c4bd537840118a268474298cc46c412b1c1b0a4f99c30d98fef38546b031777b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f85f648e799b0f656ae9204dc3fd63a78633a9f079b3c681c660336b55c3a8a5"></a>

## bot_defense.policy.js_insertion_rules — bot_defense.policy.js_insertion_rules / e0887f181825 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- bot_defense.policy.js_insertion_rules

<a id="canonical-5396d9cfabc4d4de4df3f739f12305fe61fad6381f4e2a35fe2922528b72991d"></a>

Type: `"object"`. single nested block, Optional.

Defines custom JavaScript insertion rules for Bot Defense Policy.

Upstream description:

This defines custom JavaScript insertion rules for Bot Defense Policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rules")}
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
js_insertion_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-dae4942d1080621ac80201d11a820abd4ba110f870cf7f376c8815ea053de591"></a>

## Direct properties — bot_defense.policy.js_insertion_rules / e0887f181825 / 3

- [exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-449bfaa81ce4c67e7284ec08b639061bf415a64a293944707649df7bfd58dd5e): complete subsection reference.

- [rules](resources--http_loadbalancer--reference--group-011.md#canonical-5cc68a531600301510b3d727bed81315f83d6d7d0f4337ee89e2f1fd76a1be0d): complete subsection reference.

<a id="canonical-51a9fc9dc7edcd18037a44536ade4b66aaf5b171a7e78292aba9c51ef7e941b4"></a>

## Next pages — bot_defense.policy.js_insertion_rules / e0887f181825 / 4

- [bot_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-449bfaa81ce4c67e7284ec08b639061bf415a64a293944707649df7bfd58dd5e)
- [bot_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-011.md#canonical-5cc68a531600301510b3d727bed81315f83d6d7d0f4337ee89e2f1fd76a1be0d)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-449bfaa81ce4c67e7284ec08b639061bf415a64a293944707649df7bfd58dd5e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-99128dc500b047b588f712a106ff8c929689eb173214d144ce667aae5bb43798"></a>

## bot_defense.policy.js_insertion_rules.exclude_list — bot_defense.policy.js_insertion_rules.exclude_list / e7eb6894ff54 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-c4bd537840118a268474298cc46c412b1c1b0a4f99c30d98fef38546b031777b)
- bot_defense.policy.js_insertion_rules.exclude_list

<a id="canonical-47a61558c7a2ac6169c1507b1b3851bef88b7b5783ffb76ca8c67e1c52de7826"></a>

Type: `"object"`. list nested block, Optional.

Optional JavaScript insertions exclude list of domain and path matchers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-bd215e712ac55163e6680399f6f7e3f25611e9509fb7e17869e5a8fb7ce93bd9"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.exclude_list / e7eb6894ff54 / 3

- [any_domain](resources--http_loadbalancer--reference--group-011.md#canonical-69a2bfe1046b5535ad7598e24934d2dc63e02482c85a8709321c41ae9a9887e9): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-011.md#canonical-712493f3ec6b99a6e1ee1ab655784d250b87391cf89d81026a3b25b9a985f6de): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-011.md#canonical-2a6781affcd0281f7189038d1192421c65cefa271f77842c071005cc79032b0a): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-011.md#canonical-177bfbbced1c1d2e2fe5583d35a0381cafc7cb69220e5fa7305c7e8fc48d05a9): complete subsection reference.

<a id="canonical-12dac224623ebf445ff533128aa2480f68dadd664f95986f5356c8e0209c0b00"></a>

## Next pages — bot_defense.policy.js_insertion_rules.exclude_list / e7eb6894ff54 / 4

- [bot_defense.policy.js_insertion_rules.exclude_list.any_domain](resources--http_loadbalancer--reference--group-011.md#canonical-69a2bfe1046b5535ad7598e24934d2dc63e02482c85a8709321c41ae9a9887e9)
- [bot_defense.policy.js_insertion_rules.exclude_list.domain](resources--http_loadbalancer--reference--group-011.md#canonical-712493f3ec6b99a6e1ee1ab655784d250b87391cf89d81026a3b25b9a985f6de)
- [bot_defense.policy.js_insertion_rules.exclude_list.metadata](resources--http_loadbalancer--reference--group-011.md#canonical-2a6781affcd0281f7189038d1192421c65cefa271f77842c071005cc79032b0a)
- [bot_defense.policy.js_insertion_rules.exclude_list.path](resources--http_loadbalancer--reference--group-011.md#canonical-177bfbbced1c1d2e2fe5583d35a0381cafc7cb69220e5fa7305c7e8fc48d05a9)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-c4bd537840118a268474298cc46c412b1c1b0a4f99c30d98fef38546b031777b)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-69a2bfe1046b5535ad7598e24934d2dc63e02482c85a8709321c41ae9a9887e9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fdb8c1188aac4d3356d7b1f2d7fcb3e56f3d72f427688ba1bb3aea1413b3a49e"></a>

## bot_defense.policy.js_insertion_rules.exclude_list.any_domain — bot_defense.policy.js_insertion_rules.exclude_list.any_domain / b7c3b9b330f1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-c4bd537840118a268474298cc46c412b1c1b0a4f99c30d98fef38546b031777b)
- [bot_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-449bfaa81ce4c67e7284ec08b639061bf415a64a293944707649df7bfd58dd5e)
- bot_defense.policy.js_insertion_rules.exclude_list.any_domain

<a id="canonical-908d90e93412f9df8a2d8ca68b2f4a4e96475b5af29653feec7d1a46aa66deeb"></a>

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

<a id="canonical-09a73f0c70571dfc7befb99973a561ae139edcd8c6b758b5a963c506a8c314af"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.exclude_list.any_domain / b7c3b9b330f1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6240d9b669112eefd75e700ebf059d6c5f7eb03aac99afc2c6a2ff302d668f33"></a>

## Next pages — bot_defense.policy.js_insertion_rules.exclude_list.any_domain / b7c3b9b330f1 / 4

- [bot_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-449bfaa81ce4c67e7284ec08b639061bf415a64a293944707649df7bfd58dd5e)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-712493f3ec6b99a6e1ee1ab655784d250b87391cf89d81026a3b25b9a985f6de"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eaa584a1fb41456791d296d882d61bc7f06bf83a64d8da9efe5a714842443af1"></a>

## bot_defense.policy.js_insertion_rules.exclude_list.domain — bot_defense.policy.js_insertion_rules.exclude_list.domain / 70f7dbff3c6d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-c4bd537840118a268474298cc46c412b1c1b0a4f99c30d98fef38546b031777b)
- [bot_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-449bfaa81ce4c67e7284ec08b639061bf415a64a293944707649df7bfd58dd5e)
- bot_defense.policy.js_insertion_rules.exclude_list.domain

<a id="canonical-5cb22a937c33b84892a6b0151845f28b4342edc8a2c669f847dc4949863c4cfb"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-45dd1faa9cb9e27674c43bbc76dee6cfc1350628856b6706862135eb77feffde"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.exclude_list.domain / 70f7dbff3c6d / 3

<a id="canonical-87302e815fad312f58b3bc139b7c00d80ef9634451c1a3f767903146bcdeaa67"></a>

<a id="canonical-5909f111b01e9370dce6d4a2e9c3ca6624b82ab0e486afc9578dda2a2be0f282"></a>

## exact_value property — bot_defense.policy.js_insertion_rules.exclude_list.domain / 70f7dbff3c6d / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-aa07cee6b5543c3158f854f889e9b841ca57b329bce8449ffe4bca882510b84c"></a>

<a id="canonical-7c6d68135ce5daedd0236a30090cd04f10c6f9fa2d74920f7f6fc7f63881f55b"></a>

## regex_value property — bot_defense.policy.js_insertion_rules.exclude_list.domain / 70f7dbff3c6d / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3144926ccc924c5ec9e21348f93d49fba972569cf35e9b77f145db54d99f9442"></a>

<a id="canonical-4661df91106196f3c37c4c64fedf6985a36da29e6b26169375e3605a624bee34"></a>

## suffix_value property — bot_defense.policy.js_insertion_rules.exclude_list.domain / 70f7dbff3c6d / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-326990a5bf6040b5c2ae54f4a52aefe33a28c24f809e578017abdbf988cfcc25"></a>

## Next pages — bot_defense.policy.js_insertion_rules.exclude_list.domain / 70f7dbff3c6d / 7

- [bot_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-449bfaa81ce4c67e7284ec08b639061bf415a64a293944707649df7bfd58dd5e)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-2a6781affcd0281f7189038d1192421c65cefa271f77842c071005cc79032b0a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-79dc6e8cf432de9d2963f4bb5794bb55d4801b975c7689850ec00e7ecacaaa4d"></a>

## bot_defense.policy.js_insertion_rules.exclude_list.metadata — bot_defense.policy.js_insertion_rules.exclude_list.metadata / 504ce2d64b70 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-c4bd537840118a268474298cc46c412b1c1b0a4f99c30d98fef38546b031777b)
- [bot_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-449bfaa81ce4c67e7284ec08b639061bf415a64a293944707649df7bfd58dd5e)
- bot_defense.policy.js_insertion_rules.exclude_list.metadata

<a id="canonical-6a23f137aa4b16187da1d94b35761fc9acd4cf6fa3fcfc9742cbb6046f6589f8"></a>

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

<a id="canonical-0a069ff863624c13a66abd8863030c6e8eec89b7394fc9a4b843bb4507789eef"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.exclude_list.metadata / 504ce2d64b70 / 3

<a id="canonical-6e09f5c8f8acfed7081def942137c4e51eb5ca9ddb8d9b8576ca58687d757287"></a>

<a id="canonical-c9cb1cd0d293e5b780ef6676c17abed42299fca444cf6450c160f6a65aa10574"></a>

## description_spec property — bot_defense.policy.js_insertion_rules.exclude_list.metadata / 504ce2d64b70 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-fc9ad4851cee6df3bde5d01dd1565c74fc238f6656795d0b573d3be01895bc3a"></a>

<a id="canonical-46064f4a5593c68823f7c808fe350a271824f728066485695684a0dd88c12c60"></a>

## name property — bot_defense.policy.js_insertion_rules.exclude_list.metadata / 504ce2d64b70 / 5

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

<a id="canonical-e2fb1e1ebe242114534c6163e31693a9f2cd5ae624ea28da14b4f2ea58d246ad"></a>

## Next pages — bot_defense.policy.js_insertion_rules.exclude_list.metadata / 504ce2d64b70 / 6

- [bot_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-449bfaa81ce4c67e7284ec08b639061bf415a64a293944707649df7bfd58dd5e)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-177bfbbced1c1d2e2fe5583d35a0381cafc7cb69220e5fa7305c7e8fc48d05a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c32d7edae62a2f72950ed4bb92ee6b61cc9931dc93ca3eaa0246f88f70e7587"></a>

## bot_defense.policy.js_insertion_rules.exclude_list.path — bot_defense.policy.js_insertion_rules.exclude_list.path / 7c291fa80e59 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-c4bd537840118a268474298cc46c412b1c1b0a4f99c30d98fef38546b031777b)
- [bot_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-449bfaa81ce4c67e7284ec08b639061bf415a64a293944707649df7bfd58dd5e)
- bot_defense.policy.js_insertion_rules.exclude_list.path

<a id="canonical-d212cfac16a70e542e715630c1cba09b43777d486de1ce01353ab61d2c4aa9ac"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-61700938417cce814e36efc9f7c9e444fe04e347b997018dcc8d14f918fa7b5e"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.exclude_list.path / 7c291fa80e59 / 3

<a id="canonical-0e09d1d2462be12ab188af28e7f0dd741063afdefaf5a29041b87cedb0372d47"></a>

<a id="canonical-a654cf7316c572a05a55a9e9a8fa86f0b17bcaf27442833785b5e2af4616de48"></a>

## path property — bot_defense.policy.js_insertion_rules.exclude_list.path / 7c291fa80e59 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-26492f0c9305446276babc016e18fc7b9b8a1678b2354302fa14d8f3d6293de5"></a>

<a id="canonical-8ecc1ac1f4141092128e1c870e8db525569af7f2015317dbf4a8b55cad55a6be"></a>

## prefix property — bot_defense.policy.js_insertion_rules.exclude_list.path / 7c291fa80e59 / 5

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-b318795291eb2aea0eb324d5968c74358f6f3759583d6f393c582705ecbce105"></a>

<a id="canonical-a22a8a1afe0d1cfaa01ef3c67a65a2805284cc0f015a8465715deb199201ad03"></a>

## regex property — bot_defense.policy.js_insertion_rules.exclude_list.path / 7c291fa80e59 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-10c13a615056ac9e0ebc83de2f2db637654f2cbcd7f354e88c84b37b78ecfa5b"></a>

## Next pages — bot_defense.policy.js_insertion_rules.exclude_list.path / 7c291fa80e59 / 7

- [bot_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-449bfaa81ce4c67e7284ec08b639061bf415a64a293944707649df7bfd58dd5e)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-5cc68a531600301510b3d727bed81315f83d6d7d0f4337ee89e2f1fd76a1be0d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b3c0684bed99d62e5ef8ea3915652733b00fbe3c537d4de0c0af8aa8b8ad7a6f"></a>

## bot_defense.policy.js_insertion_rules.rules — bot_defense.policy.js_insertion_rules.rules / fc8699e19d97 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-c4bd537840118a268474298cc46c412b1c1b0a4f99c30d98fef38546b031777b)
- bot_defense.policy.js_insertion_rules.rules

<a id="canonical-9009d7dd4df9fdd93ccde17fb7515512a0a178a9fe7fae77cdb987028089abe9"></a>

Type: `"object"`. list nested block, Optional.

Required list of pages to insert Bot Defense client JavaScript.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-7afac3d64e3e5a0a657c0e9019c5014db7a90f77c8d106018c1593c7a8c45948"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.rules / fc8699e19d97 / 3

- [any_domain](resources--http_loadbalancer--reference--group-011.md#canonical-410bb086d4d87adc961fa3ad99cbcf60c216061ae05a4a6b2e70c4341025bf34): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-011.md#canonical-80e5a032d43ff21c551d2f6dc7f5656ed587e927dc01a2bfbd5791698c620a4d): complete subsection reference.

<a id="canonical-a33f3134b71819bac319b35f472f72992369ef13f40db2438a6b22a5f3443fa6"></a>

<a id="canonical-f1d26dd09368c3141ac7a4b314b05212e488179e740997b10a785780c4d7f901"></a>

## javascript_location property — bot_defense.policy.js_insertion_rules.rules / fc8699e19d97 / 4

Type: `"string"`. Optional.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [metadata](resources--http_loadbalancer--reference--group-011.md#canonical-bf7528c579e50c3cfbdbc7cff2dc7d3b9a41c8f2ca203d6745aaea7154cdb576): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-011.md#canonical-d5e74b6b8800054a840056d857ef8714c65b5efb71cbd4c803d441c62a0e4a09): complete subsection reference.

<a id="canonical-e999384e4cbd256686e391be8ea8937933e1927947b7d596d6b9a7a90e40b34c"></a>

## Next pages — bot_defense.policy.js_insertion_rules.rules / fc8699e19d97 / 5

- [bot_defense.policy.js_insertion_rules.rules.any_domain](resources--http_loadbalancer--reference--group-011.md#canonical-410bb086d4d87adc961fa3ad99cbcf60c216061ae05a4a6b2e70c4341025bf34)
- [bot_defense.policy.js_insertion_rules.rules.domain](resources--http_loadbalancer--reference--group-011.md#canonical-80e5a032d43ff21c551d2f6dc7f5656ed587e927dc01a2bfbd5791698c620a4d)
- [bot_defense.policy.js_insertion_rules.rules.metadata](resources--http_loadbalancer--reference--group-011.md#canonical-bf7528c579e50c3cfbdbc7cff2dc7d3b9a41c8f2ca203d6745aaea7154cdb576)
- [bot_defense.policy.js_insertion_rules.rules.path](resources--http_loadbalancer--reference--group-011.md#canonical-d5e74b6b8800054a840056d857ef8714c65b5efb71cbd4c803d441c62a0e4a09)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-c4bd537840118a268474298cc46c412b1c1b0a4f99c30d98fef38546b031777b)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-410bb086d4d87adc961fa3ad99cbcf60c216061ae05a4a6b2e70c4341025bf34"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fca253d05d6d27dbda174c70b5f4821c27ecb93f9dc840976639ec58c03730ba"></a>

## bot_defense.policy.js_insertion_rules.rules.any_domain — bot_defense.policy.js_insertion_rules.rules.any_domain / 062c4b2d4a3c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-c4bd537840118a268474298cc46c412b1c1b0a4f99c30d98fef38546b031777b)
- [bot_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-011.md#canonical-5cc68a531600301510b3d727bed81315f83d6d7d0f4337ee89e2f1fd76a1be0d)
- bot_defense.policy.js_insertion_rules.rules.any_domain

<a id="canonical-a3f196737b76333009ea9cccef628c98a447bc1889ee54bcd44ff468db653638"></a>

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

<a id="canonical-e8e52364a9b1f2998f101e61fd45ade568d5fb30a73f5a31b7809ce07634631f"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.rules.any_domain / 062c4b2d4a3c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8dd1e8e7d23a515a28806baec7a049de01b846108e516251c469f19f95757976"></a>

## Next pages — bot_defense.policy.js_insertion_rules.rules.any_domain / 062c4b2d4a3c / 4

- [bot_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-011.md#canonical-5cc68a531600301510b3d727bed81315f83d6d7d0f4337ee89e2f1fd76a1be0d)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-80e5a032d43ff21c551d2f6dc7f5656ed587e927dc01a2bfbd5791698c620a4d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a147769b5675972e8b196fab9ec561be7ee5dd93eb842e533d001d384ac623d1"></a>

## bot_defense.policy.js_insertion_rules.rules.domain — bot_defense.policy.js_insertion_rules.rules.domain / 793d0f9fdc3b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-c4bd537840118a268474298cc46c412b1c1b0a4f99c30d98fef38546b031777b)
- [bot_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-011.md#canonical-5cc68a531600301510b3d727bed81315f83d6d7d0f4337ee89e2f1fd76a1be0d)
- bot_defense.policy.js_insertion_rules.rules.domain

<a id="canonical-eb0877b592e03f14413da5b99664c0a51caf9edfa6753d53805e24bcd61c03b8"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-ff6583f450c332d086571c3a871654d2831ccec39b40694eadb9cff6bb8bb821"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.rules.domain / 793d0f9fdc3b / 3

<a id="canonical-a2fed637f64c0e0bdaee956a8201a5ab85c019624ce8c12595efdf4fce13a615"></a>

<a id="canonical-b9f2337a6858ef9ee958f5f24d5847f3acfac053e61d5f67883662a0af23db32"></a>

## exact_value property — bot_defense.policy.js_insertion_rules.rules.domain / 793d0f9fdc3b / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-4adc20ee51da9bc7d9b08a3452ff18704561bc1a0316724e6f9fdee06c7126cf"></a>

<a id="canonical-7fd02df4b66c5464320c1bdb3e4d4233287beac1d12b1dfcd689a773a6882d1e"></a>

## regex_value property — bot_defense.policy.js_insertion_rules.rules.domain / 793d0f9fdc3b / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-f2d5cb456ef9a3091b5aa90525e7a35744a0e802a503865c1909309eafd2d715"></a>

<a id="canonical-3c2877c9c109ae69d48f67207c836120d217b1a162237304045d542baa6d04f2"></a>

## suffix_value property — bot_defense.policy.js_insertion_rules.rules.domain / 793d0f9fdc3b / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0a1a4188f977583359532db0a1a99576edfca7ea4b4c88feda0f59a3d8537147"></a>

## Next pages — bot_defense.policy.js_insertion_rules.rules.domain / 793d0f9fdc3b / 7

- [bot_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-011.md#canonical-5cc68a531600301510b3d727bed81315f83d6d7d0f4337ee89e2f1fd76a1be0d)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-bf7528c579e50c3cfbdbc7cff2dc7d3b9a41c8f2ca203d6745aaea7154cdb576"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ff0796dbf90f05b0952f3ac2d974d1f012f0077d10d8c5b9459eb83e96eb285"></a>

## bot_defense.policy.js_insertion_rules.rules.metadata — bot_defense.policy.js_insertion_rules.rules.metadata / cc8aa0337c96 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-c4bd537840118a268474298cc46c412b1c1b0a4f99c30d98fef38546b031777b)
- [bot_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-011.md#canonical-5cc68a531600301510b3d727bed81315f83d6d7d0f4337ee89e2f1fd76a1be0d)
- bot_defense.policy.js_insertion_rules.rules.metadata

<a id="canonical-9bc00e09e93c9087226912f119619e43974a1c9256a0f66d25584074de484dcd"></a>

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

<a id="canonical-f41715c83a8eb8a5eb0f3b728d31ff5196d168811c8e2f77bfa8923c44fa9790"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.rules.metadata / cc8aa0337c96 / 3

<a id="canonical-32f639903d0d1f7e4050ac8e2b27dabb0c96cd4037ce2f1bfa260be096ba1e72"></a>

<a id="canonical-8ce3c4dabf182ef3aea7b1a3081b7ad0d257717668562cc68b88141f9c78ddf1"></a>

## description_spec property — bot_defense.policy.js_insertion_rules.rules.metadata / cc8aa0337c96 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-bdb269d688cee79fe30f95746d90a10d5f827659ed09d6480430e8fc8136d36a"></a>

<a id="canonical-9e68e608e9a8805bccfad2d419fe3434e6d372e60b2131a9da1990460d059644"></a>

## name property — bot_defense.policy.js_insertion_rules.rules.metadata / cc8aa0337c96 / 5

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

<a id="canonical-02abb754533bde795911dd067aada1668cb203c2c4010a545ae0869029d4d8eb"></a>

## Next pages — bot_defense.policy.js_insertion_rules.rules.metadata / cc8aa0337c96 / 6

- [bot_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-011.md#canonical-5cc68a531600301510b3d727bed81315f83d6d7d0f4337ee89e2f1fd76a1be0d)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d5e74b6b8800054a840056d857ef8714c65b5efb71cbd4c803d441c62a0e4a09"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6784ed74b803296e0d73b2469929f6882bab7dfeb25eec9a2b7c67be77c31bc1"></a>

## bot_defense.policy.js_insertion_rules.rules.path — bot_defense.policy.js_insertion_rules.rules.path / 05cd343a77ce / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-c4bd537840118a268474298cc46c412b1c1b0a4f99c30d98fef38546b031777b)
- [bot_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-011.md#canonical-5cc68a531600301510b3d727bed81315f83d6d7d0f4337ee89e2f1fd76a1be0d)
- bot_defense.policy.js_insertion_rules.rules.path

<a id="canonical-b45030faf64308e7a9538cc1a4000e05779533d85941e3967b8c6e60138574a3"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-ae610b50d26f93201670d61aad5ace3b79ceac1f55a57129244989ec0e976e70"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.rules.path / 05cd343a77ce / 3

<a id="canonical-37b47682ff8a570f957ea2c619bce72c9dac2e624b8182e442fe519f9335b5b8"></a>

<a id="canonical-47857d279b110a1153ae6cb63c7074a487e79c2622f187ed8f332c7959bada46"></a>

## path property — bot_defense.policy.js_insertion_rules.rules.path / 05cd343a77ce / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-a3f13c0acdf1c210566907f8449643f8b4c893710911b604a23aea8feee30c1b"></a>

<a id="canonical-2d58eecef85084d000a6e38d5d472601e246916bbeaec9439e94dfbd1046e96f"></a>

## prefix property — bot_defense.policy.js_insertion_rules.rules.path / 05cd343a77ce / 5

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-f5da406c12b64eb69eac7a70ad95638271bd0d7b795818e3e5b17e6cd3ebf292"></a>

<a id="canonical-21cc44ad2cbfcd6cc1f6b480441afd21f84df31c7b5386de09bff8ace27f5291"></a>

## regex property — bot_defense.policy.js_insertion_rules.rules.path / 05cd343a77ce / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-573753b5778693a7a21725fb0cae54f6ad510cd823327bf0b35e667846303f2d"></a>

## Next pages — bot_defense.policy.js_insertion_rules.rules.path / 05cd343a77ce / 7

- [bot_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-011.md#canonical-5cc68a531600301510b3d727bed81315f83d6d7d0f4337ee89e2f1fd76a1be0d)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-9b456501acaee5f5e09376a1cb9063698f3bb87e2f63196ff20cbe7ecd77188c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b37aff7a435849aa0d982c6de96b4646c333f4ab98bf015f2a4f6f407067d5d5"></a>

## bot_defense.policy.mobile_sdk_config — bot_defense.policy.mobile_sdk_config / 3d8cdac236a6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- bot_defense.policy.mobile_sdk_config

<a id="canonical-1ca766cf37fa4ed60fd548ab32da51fcc28abc23f9479285fca66570211d09c7"></a>

Type: `"object"`. single nested block, Optional.

Mobile SDK Configuration. Mobile SDK configuration.

Upstream description:

Mobile SDK configuration.

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
mobile_sdk_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-f9a0fa4846e3cad2a49d92cf3690d3410d04c992621eb8113d0cff9306a40fcd"></a>

## Direct properties — bot_defense.policy.mobile_sdk_config / 3d8cdac236a6 / 3

- [mobile_identifier](resources--http_loadbalancer--reference--group-011.md#canonical-8805bf2e7d3f84cc909ad72ca5226ecbadeffd055ca68aa250b4cb930829cea7): complete subsection reference.

<a id="canonical-562c0292aa5405ae32fc2cfc83be88fc6fc5d06364d41f3d7014b98b6cd996b1"></a>

## Next pages — bot_defense.policy.mobile_sdk_config / 3d8cdac236a6 / 4

- [bot_defense.policy.mobile_sdk_config.mobile_identifier](resources--http_loadbalancer--reference--group-011.md#canonical-8805bf2e7d3f84cc909ad72ca5226ecbadeffd055ca68aa250b4cb930829cea7)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-8805bf2e7d3f84cc909ad72ca5226ecbadeffd055ca68aa250b4cb930829cea7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e739aa937c21431e0b54a8e96624d253c3078885b322ae334b0e78175463d313"></a>

## bot_defense.policy.mobile_sdk_config.mobile_identifier — bot_defense.policy.mobile_sdk_config.mobile_identifier / 49ca28fb60df / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.mobile_sdk_config](resources--http_loadbalancer--reference--group-011.md#canonical-9b456501acaee5f5e09376a1cb9063698f3bb87e2f63196ff20cbe7ecd77188c)
- bot_defense.policy.mobile_sdk_config.mobile_identifier

<a id="canonical-4b3ee0d3f87ee2fece081aa8f095f56eb027b169d9a29c8e8626550bc57dabc7"></a>

Type: `"object"`. single nested block, Optional.

Mobile Traffic Identifier. Mobile traffic identifier type.

Upstream description:

Mobile traffic identifier type.

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
mobile_identifier {
  # Configure direct properties listed below.
}
```

<a id="canonical-03781de49d687432c132ac6513b8873e2b351e68c808421d091bc96410231e66"></a>

## Direct properties — bot_defense.policy.mobile_sdk_config.mobile_identifier / 49ca28fb60df / 3

- [headers](resources--http_loadbalancer--reference--group-011.md#canonical-1ca10a2cd96f15c8b2e4aa331b6de6af2827fe21fae744bc38adb0cc5380ac39): complete subsection reference.

<a id="canonical-6e894cd2fd8f2faefd965607963db3f501c86208daf0e8ac50f7a5ad099d1ffd"></a>

## Next pages — bot_defense.policy.mobile_sdk_config.mobile_identifier / 49ca28fb60df / 4

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](resources--http_loadbalancer--reference--group-011.md#canonical-1ca10a2cd96f15c8b2e4aa331b6de6af2827fe21fae744bc38adb0cc5380ac39)
- [bot_defense.policy.mobile_sdk_config](resources--http_loadbalancer--reference--group-011.md#canonical-9b456501acaee5f5e09376a1cb9063698f3bb87e2f63196ff20cbe7ecd77188c)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-1ca10a2cd96f15c8b2e4aa331b6de6af2827fe21fae744bc38adb0cc5380ac39"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7f83f8f0d5a0dfaebf96a586b2b8d35a7d8e1d62e553ef234c6b44e3bb3f0cc3"></a>

## bot_defense.policy.mobile_sdk_config.mobile_identifier.headers — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers / 3ac5fc8e5c7d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.mobile_sdk_config](resources--http_loadbalancer--reference--group-011.md#canonical-9b456501acaee5f5e09376a1cb9063698f3bb87e2f63196ff20cbe7ecd77188c)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](resources--http_loadbalancer--reference--group-011.md#canonical-8805bf2e7d3f84cc909ad72ca5226ecbadeffd055ca68aa250b4cb930829cea7)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers

<a id="canonical-1b98f64272dd514116231679ecbf17032c98ccf6f8d1170eb620140187d4b775"></a>

Type: `"object"`. list nested block, Optional.

Headers that can be used to identify mobile traffic.

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
    "minItems": 0,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-285ab27fa6b23d0cdff5aac3870b69ce445113040e9c9c1785b72d9467060ba7"></a>

## Direct properties — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers / 3ac5fc8e5c7d / 3

- [check_not_present](resources--http_loadbalancer--reference--group-011.md#canonical-054478adf82be92a29564ba120c9feebc010c71207807c667ec41af3fd489e29): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-011.md#canonical-64fc8c55ccf0561adf1ca0a3239ee0f4a1a0022a092441784e9be6b405d50b71): complete subsection reference.

- [item](resources--http_loadbalancer--reference--group-011.md#canonical-aae87dc83696ec38d36a94afac7f1e93206281428724f9d02f471ee9cd6d418e): complete subsection reference.

<a id="canonical-ebc7e5fc7927791c6bf9cef830f741475c6207476dd381da3e2abebb955b1946"></a>

<a id="canonical-ce57097e94a9d0cc99a2b484974bf07f33160c04593429fad06843bf83dca884"></a>

## name property — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers / 3ac5fc8e5c7d / 4

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

<a id="canonical-3a18916f8a634da56332dbdeb1464803c856daca88f14215f769ed5d9f046d56"></a>

## Next pages — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers / 3ac5fc8e5c7d / 5

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present](resources--http_loadbalancer--reference--group-011.md#canonical-054478adf82be92a29564ba120c9feebc010c71207807c667ec41af3fd489e29)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present](resources--http_loadbalancer--reference--group-011.md#canonical-64fc8c55ccf0561adf1ca0a3239ee0f4a1a0022a092441784e9be6b405d50b71)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item](resources--http_loadbalancer--reference--group-011.md#canonical-aae87dc83696ec38d36a94afac7f1e93206281428724f9d02f471ee9cd6d418e)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](resources--http_loadbalancer--reference--group-011.md#canonical-8805bf2e7d3f84cc909ad72ca5226ecbadeffd055ca68aa250b4cb930829cea7)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-054478adf82be92a29564ba120c9feebc010c71207807c667ec41af3fd489e29"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63fa675d4da576e2f285ac02676d0a07ddd172447f22ebed5938e296b7777d88"></a>

## bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present / efb8357c617f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.mobile_sdk_config](resources--http_loadbalancer--reference--group-011.md#canonical-9b456501acaee5f5e09376a1cb9063698f3bb87e2f63196ff20cbe7ecd77188c)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](resources--http_loadbalancer--reference--group-011.md#canonical-8805bf2e7d3f84cc909ad72ca5226ecbadeffd055ca68aa250b4cb930829cea7)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](resources--http_loadbalancer--reference--group-011.md#canonical-1ca10a2cd96f15c8b2e4aa331b6de6af2827fe21fae744bc38adb0cc5380ac39)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present

<a id="canonical-5f59aaf4894e9676b43101ff044afc4de17caa00a25f6d4a6830a6892390147b"></a>

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

<a id="canonical-814b1db1122ad5dd4f9ec47f12e2da4ce9c70bccb8c482c179a2e64ef899039c"></a>

## Direct properties — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present / efb8357c617f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2005611c7592b7e18049b3a6203d1089df2092dc6fdcc7675f0d75691396a1eb"></a>

## Next pages — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present / efb8357c617f / 4

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](resources--http_loadbalancer--reference--group-011.md#canonical-1ca10a2cd96f15c8b2e4aa331b6de6af2827fe21fae744bc38adb0cc5380ac39)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-64fc8c55ccf0561adf1ca0a3239ee0f4a1a0022a092441784e9be6b405d50b71"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-799db799ffcc0723198337bffb61285229695e8d73dabb9e85b6f69162eb4a32"></a>

## bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present / 8e031097fa11 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.mobile_sdk_config](resources--http_loadbalancer--reference--group-011.md#canonical-9b456501acaee5f5e09376a1cb9063698f3bb87e2f63196ff20cbe7ecd77188c)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](resources--http_loadbalancer--reference--group-011.md#canonical-8805bf2e7d3f84cc909ad72ca5226ecbadeffd055ca68aa250b4cb930829cea7)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](resources--http_loadbalancer--reference--group-011.md#canonical-1ca10a2cd96f15c8b2e4aa331b6de6af2827fe21fae744bc38adb0cc5380ac39)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present

<a id="canonical-d306ffb2b9013389feb1bc487c2a816e33dc09f9d5b4f816452de82c78f52ac7"></a>

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

<a id="canonical-588381e19a64f4304f52fc99f9d2e5bf7e4ede6cecf5f726d11ba7d7449e329a"></a>

## Direct properties — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present / 8e031097fa11 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7cc7d224646cb8dded375e0826ce5a2bc31b20f3ea304a08f3adca777bdbabe3"></a>

## Next pages — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present / 8e031097fa11 / 4

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](resources--http_loadbalancer--reference--group-011.md#canonical-1ca10a2cd96f15c8b2e4aa331b6de6af2827fe21fae744bc38adb0cc5380ac39)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-aae87dc83696ec38d36a94afac7f1e93206281428724f9d02f471ee9cd6d418e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e8c0f6df3c4524c37fa5a08009c21724e1e0649d3a31a7c3b1789af4aed21f6c"></a>

## bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item / e32acbd90448 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.mobile_sdk_config](resources--http_loadbalancer--reference--group-011.md#canonical-9b456501acaee5f5e09376a1cb9063698f3bb87e2f63196ff20cbe7ecd77188c)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](resources--http_loadbalancer--reference--group-011.md#canonical-8805bf2e7d3f84cc909ad72ca5226ecbadeffd055ca68aa250b4cb930829cea7)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](resources--http_loadbalancer--reference--group-011.md#canonical-1ca10a2cd96f15c8b2e4aa331b6de6af2827fe21fae744bc38adb0cc5380ac39)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item

<a id="canonical-226b0662ac87ed1b79dcfc638e17bb69f1795ba52712d892617611cb1b438535"></a>

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

<a id="canonical-ededf858ba0260be8cfd5413eb6e29f46561b1059e5f87e852e90529fe5f99ee"></a>

## Direct properties — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item / e32acbd90448 / 3

<a id="canonical-bd175d54db56a644bbf9a3b780cc70cf4e78deab609e64fca989f50a7030a036"></a>

<a id="canonical-8f40ab7a7b2d778c589533a8fa3db930df7aae707052360d5f53b886ee7d8ae5"></a>

## exact_values property — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item / e32acbd90448 / 4

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

<a id="canonical-a2c94e5c392cc56c8cf680e6538fc0097e108ffe9713e19544d3bd1c8566072c"></a>

<a id="canonical-80cc1c8e234a7f7527305ce3ce74ad7ce66574f79c8db29285cdb3efc85cf8ef"></a>

## regex_values property — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item / e32acbd90448 / 5

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

<a id="canonical-e76fd824650618f36f901b135ae5832ad3d16b9faea4d76a3201b325d45edf88"></a>

<a id="canonical-55f167ce2bad8fd28e7fff725bce727149c5fa120111f0b9b2e462e0851df3f3"></a>

## transformers property — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item / e32acbd90448 / 6

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

<a id="canonical-3ef8abd8644733dc52b523fc5cf6fc5ef6a38585076020407c1a4163ebecb43e"></a>

## Next pages — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item / e32acbd90448 / 7

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](resources--http_loadbalancer--reference--group-011.md#canonical-1ca10a2cd96f15c8b2e4aa331b6de6af2827fe21fae744bc38adb0cc5380ac39)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-faf4809a06a5d579ba688a35f93e0d296e018f0f96e7675818eaaa1345fd9d1e"></a>

## bot_defense.policy.protected_app_endpoints — bot_defense.policy.protected_app_endpoints / 50e864146857 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- bot_defense.policy.protected_app_endpoints

<a id="canonical-da1952a107dde37afb8738a5a4522aabd69d0cb8ce9100fc3a22ecd26c06cedf"></a>

Type: `"object"`. list nested block, Optional.

List of protected endpoints. Limit: Approx '128 endpoints per Load Balancer (LB)' upto 4 LBs, '32
endpoints per LB' after 4 LBs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("http_methods"),
  validators.ConflictingListObjectAttributes("allow_good_bots",
    "mitigate_good_bots"),
  validators.ConflictingListObjectAttributes("any_domain",
    "domain"),
  validators.ConflictingListObjectAttributes("flow_label",
    "undefined_flow_label"),
  validators.ConflictingListObjectAttributes("mobile",
    "web"),
  validators.ConflictingListObjectAttributes("mobile",
    "web_mobile"),
  validators.ConflictingListObjectAttributes("web",
    "web_mobile")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
protected_app_endpoints {
  # Configure direct properties listed below.
}
```

<a id="canonical-b1f951f23e163f4733a356c4f74022243009a81d5834e20eb5b3f0d266909b80"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints / 50e864146857 / 3

- [allow_good_bots](resources--http_loadbalancer--reference--group-011.md#canonical-48d98c9eca313ea7d0e94be79a0cfd0906de16cccfc234711b7916d1291dc503): complete subsection reference.

- [any_domain](resources--http_loadbalancer--reference--group-011.md#canonical-61e97d620259421ab75c98a90bbe8b5abdfe277a596b4aef47b5d5e360ea7c87): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-011.md#canonical-7672589ee6db1ea64feb27c38da23765338348d872f64c17487a9d1d068de9e3): complete subsection reference.

- [flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82): complete subsection reference.

- [headers](resources--http_loadbalancer--reference--group-012.md#canonical-c50cfd111a202642cd8a6ca40991957a620e38ffdcbc622e946d7c9252a332a3): complete subsection reference.

<a id="canonical-4f52573273631c6849d316716a0fe3d73a723116ccf53fe9c6a889c067fdd133"></a>

<a id="canonical-09ab7cc69827173c263e43ed94f94c952efd60d2cfa594c5b5ee1da90c7ff42c"></a>

## http_methods property — bot_defense.policy.protected_app_endpoints / 50e864146857 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
METHOD\_ANY|METHOD\_GET|METHOD\_POST|METHOD\_PUT|METHOD\_PATCH|METHOD\_DELETE|METHOD\_GET\_DOCUMENT\]
HTTP Methods. List of HTTP methods. Possible values are \`METHOD\_ANY\`, \`METHOD\_GET\`,
\`METHOD\_POST\`, \`METHOD\_PUT\`, \`METHOD\_PATCH\`, \`METHOD\_DELETE\`, \`METHOD\_GET\_DOCUMENT\`.
Defaults to \`METHOD\_ANY\`.

Upstream description:

List of HTTP methods.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 5),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[0,1,3,4,10]",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[0,1,3,4,10]",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [metadata](resources--http_loadbalancer--reference--group-012.md#canonical-8b9bb68551efc78974162af61d69c50411f1fd66c3bbe0f9ba14caa05e79e09a): complete subsection reference.

- [mitigate_good_bots](resources--http_loadbalancer--reference--group-012.md#canonical-e0f93575fd4f9e09f02b9583a890ec54d9190308d4e103d4a770ed9bb6b8afca): complete subsection reference.

- [mitigation](resources--http_loadbalancer--reference--group-012.md#canonical-ea61011b5b44901d7f05b8d65cfa06ec62f1c302f3b1e7c4e32d6f8191aceffd): complete subsection reference.

- [mobile](resources--http_loadbalancer--reference--group-012.md#canonical-20d0718577d0874b23e50c5f919b2484f8b798aecd77321c1833315668f5d538): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-012.md#canonical-de209870a7c605b348e30eacce83fff5f5442f9b4484b18a1aa1170b78643f4f): complete subsection reference.

<a id="canonical-ada644bf49f257e5d8855f8312033cf894d42584f9341398f66846e385d7b6c6"></a>

<a id="canonical-9b2eafdffe6f2536547f59fd10162667039d8f08b4627870621a0f1751261389"></a>

## protocol property — bot_defense.policy.protected_app_endpoints / 50e864146857 / 5

Type: `"string"`. Optional.

\[Enum: BOTH|HTTP|HTTPS\] SchemeType is used to indicate URL scheme. - BOTH: BOTH URL scheme for
HTTPS:// or HTTP://. - HTTP: HTTP URL scheme HTTP:// only. - HTTPS: HTTPS URL scheme HTTPS:// only.
Possible values are \`BOTH\`, \`HTTP\`, \`HTTPS\`. Defaults to \`BOTH\`.

Upstream description:

SchemeType is used to indicate URL scheme.

&#8203;- BOTH: BOTH

URL scheme for HTTPS:// or HTTP://. &#8203;- HTTP: HTTP

URL scheme HTTP:// only. &#8203;- HTTPS: HTTPS

URL scheme HTTPS:// only.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("BOTH",
    "HTTP",
    "HTTPS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "BOTH",
  "enum": [
    "BOTH",
    "HTTP",
    "HTTPS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [query_params](resources--http_loadbalancer--reference--group-012.md#canonical-47b8ba621183a1aa542dafd9635d726b2f9d6fe73de146e5d57618e8f7e16b5e): complete subsection reference.

- [undefined_flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-63939a980f6231f79d98a18e6fe71dfcd8d673eef2c80a60cbc851752e35c7ab): complete subsection reference.

- [web](resources--http_loadbalancer--reference--group-012.md#canonical-1c21b5e02cda152a7fc62c5cf2972c46ada5da9df9ee1c3f1142fd678808b4f8): complete subsection reference.

- [web_mobile](resources--http_loadbalancer--reference--group-012.md#canonical-d1dd9a62713870cb68c6af01a169c34eb791ba8dcbd8fba2919ee95722f29d72): complete subsection reference.

<a id="canonical-fa5252a4650fce48265cd3ff8ebe4d7e654c5cd1826281608e0f03d51eb617bf"></a>

## Next pages — bot_defense.policy.protected_app_endpoints / 50e864146857 / 6

- [bot_defense.policy.protected_app_endpoints.allow_good_bots](resources--http_loadbalancer--reference--group-011.md#canonical-48d98c9eca313ea7d0e94be79a0cfd0906de16cccfc234711b7916d1291dc503)
- [bot_defense.policy.protected_app_endpoints.any_domain](resources--http_loadbalancer--reference--group-011.md#canonical-61e97d620259421ab75c98a90bbe8b5abdfe277a596b4aef47b5d5e360ea7c87)
- [bot_defense.policy.protected_app_endpoints.domain](resources--http_loadbalancer--reference--group-011.md#canonical-7672589ee6db1ea64feb27c38da23765338348d872f64c17487a9d1d068de9e3)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.headers](resources--http_loadbalancer--reference--group-012.md#canonical-c50cfd111a202642cd8a6ca40991957a620e38ffdcbc622e946d7c9252a332a3)
- [bot_defense.policy.protected_app_endpoints.metadata](resources--http_loadbalancer--reference--group-012.md#canonical-8b9bb68551efc78974162af61d69c50411f1fd66c3bbe0f9ba14caa05e79e09a)
- [bot_defense.policy.protected_app_endpoints.mitigate_good_bots](resources--http_loadbalancer--reference--group-012.md#canonical-e0f93575fd4f9e09f02b9583a890ec54d9190308d4e103d4a770ed9bb6b8afca)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--http_loadbalancer--reference--group-012.md#canonical-ea61011b5b44901d7f05b8d65cfa06ec62f1c302f3b1e7c4e32d6f8191aceffd)
- [bot_defense.policy.protected_app_endpoints.mobile](resources--http_loadbalancer--reference--group-012.md#canonical-20d0718577d0874b23e50c5f919b2484f8b798aecd77321c1833315668f5d538)
- [bot_defense.policy.protected_app_endpoints.path](resources--http_loadbalancer--reference--group-012.md#canonical-de209870a7c605b348e30eacce83fff5f5442f9b4484b18a1aa1170b78643f4f)
- [bot_defense.policy.protected_app_endpoints.query_params](resources--http_loadbalancer--reference--group-012.md#canonical-47b8ba621183a1aa542dafd9635d726b2f9d6fe73de146e5d57618e8f7e16b5e)
- [bot_defense.policy.protected_app_endpoints.undefined_flow_label](resources--http_loadbalancer--reference--group-012.md#canonical-63939a980f6231f79d98a18e6fe71dfcd8d673eef2c80a60cbc851752e35c7ab)
- [bot_defense.policy.protected_app_endpoints.web](resources--http_loadbalancer--reference--group-012.md#canonical-1c21b5e02cda152a7fc62c5cf2972c46ada5da9df9ee1c3f1142fd678808b4f8)
- [bot_defense.policy.protected_app_endpoints.web_mobile](resources--http_loadbalancer--reference--group-012.md#canonical-d1dd9a62713870cb68c6af01a169c34eb791ba8dcbd8fba2919ee95722f29d72)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-48d98c9eca313ea7d0e94be79a0cfd0906de16cccfc234711b7916d1291dc503"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e278c5949cc4f1a1f293754639ebbbcdfbf394984941f3e727ff04e12013c8b0"></a>

## bot_defense.policy.protected_app_endpoints.allow_good_bots — bot_defense.policy.protected_app_endpoints.allow_good_bots / 7b1451ccabe4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- bot_defense.policy.protected_app_endpoints.allow_good_bots

<a id="canonical-7462bc6a9e5848e4002ba0c23465089b5585d8d8d76cbe2fb67196fa00f590ee"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for allow good bots.

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
allow_good_bots = {}
```

<a id="canonical-7ee6ccf34c20ecd37815a241a4deca1250c81ec1c639e9be5ef005e4c082c58e"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.allow_good_bots / 7b1451ccabe4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-de7eacc232f06b0f2bb62ff4fb494fc0b8cbf44d8df62df819440427f409885e"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.allow_good_bots / 7b1451ccabe4 / 4

- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-61e97d620259421ab75c98a90bbe8b5abdfe277a596b4aef47b5d5e360ea7c87"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a80611df5e5dab02ee34283507e369c0a65e30eb437fd6b71de27b8f4dda48e7"></a>

## bot_defense.policy.protected_app_endpoints.any_domain — bot_defense.policy.protected_app_endpoints.any_domain / baedb258262b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- bot_defense.policy.protected_app_endpoints.any_domain

<a id="canonical-4fda515d5a396004debfe11fcccc8dca68f63abe3ee1db36db341c218566ec87"></a>

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

<a id="canonical-6fe1228074403e320053853e95cb5cf445a8db47f80455850bee517df769e3bd"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.any_domain / baedb258262b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cdbe415f677495e710faa741474c23e8718eb7c24007d8da217db26c87cc741f"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.any_domain / baedb258262b / 4

- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-7672589ee6db1ea64feb27c38da23765338348d872f64c17487a9d1d068de9e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae4c4c6fcb4e84beb419622fdb0dd291a95f007926e83783b22dba26020c1c22"></a>

## bot_defense.policy.protected_app_endpoints.domain — bot_defense.policy.protected_app_endpoints.domain / 7c1040326a89 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- bot_defense.policy.protected_app_endpoints.domain

<a id="canonical-f6e0f62cbc3ec22b9992f46fb4d12d63a12824cc312846728086ebf85c5830d3"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-95eaed395af7ae6902f9cfa4cdc37791d31804891250c86d1db52e35c48ec934"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.domain / 7c1040326a89 / 3

<a id="canonical-5b2db7f2e2a22632d2a7874baeff332284d28f66cdd967f9a61629196200589c"></a>

<a id="canonical-0690ac1969bbdfcd09709136da4618cb2178464c3c20a196a228aced78d1a4a1"></a>

## exact_value property — bot_defense.policy.protected_app_endpoints.domain / 7c1040326a89 / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-b27414e663f2f5eaf71769843f9e0a40937254b316c43c677ccddde6766ae08d"></a>

<a id="canonical-169607b097a43855fcb8aa06e4d95aaebee672882f777522bf2cd29e0a7f54c8"></a>

## regex_value property — bot_defense.policy.protected_app_endpoints.domain / 7c1040326a89 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-7b0d6425bb86ad0bf9253e16accfd0d62928edbc2c8fb046c61b6d22c5516c62"></a>

<a id="canonical-366ecd0062408a03efa8a35ab912253b44b325fed07ab60cb0a936cca913920a"></a>

## suffix_value property — bot_defense.policy.protected_app_endpoints.domain / 7c1040326a89 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-eb8cccd5d65d8246c5e867caaca3f40ada7975eafe005b607c819967a5b97961"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.domain / 7c1040326a89 / 7

- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2dafbb55ef7035716fb2248a85ed7167a9020bae3670a8e0b504bb9a2f35ce81"></a>

## bot_defense.policy.protected_app_endpoints.flow_label — bot_defense.policy.protected_app_endpoints.flow_label / 81d056f86e08 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- bot_defense.policy.protected_app_endpoints.flow_label

<a id="canonical-3b3b65df8a3a31da1c4ab2f40a4c7801146be84b5d2a38cb8da97af4676035b0"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Category allows to associate traffic with selected category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("account_management",
    "authentication"),
  validators.ConflictingObjectAttributes("account_management",
    "financial_services"),
  validators.ConflictingObjectAttributes("account_management",
    "flight"),
  validators.ConflictingObjectAttributes("account_management",
    "profile_management"),
  validators.ConflictingObjectAttributes("account_management",
    "search"),
  validators.ConflictingObjectAttributes("account_management",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("authentication",
    "financial_services"),
  validators.ConflictingObjectAttributes("authentication",
    "flight"),
  validators.ConflictingObjectAttributes("authentication",
    "profile_management"),
  validators.ConflictingObjectAttributes("authentication",
    "search"),
  validators.ConflictingObjectAttributes("authentication",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("financial_services",
    "flight"),
  validators.ConflictingObjectAttributes("financial_services",
    "profile_management"),
  validators.ConflictingObjectAttributes("financial_services",
    "search"),
  validators.ConflictingObjectAttributes("financial_services",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("flight",
    "profile_management"),
  validators.ConflictingObjectAttributes("flight",
    "search"),
  validators.ConflictingObjectAttributes("flight",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("profile_management",
    "search"),
  validators.ConflictingObjectAttributes("profile_management",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("search",
    "shopping_gift_cards")}
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
  "x-ves-oneof-field-flow_label_choice": "[\"account_management\",\"authentication\",\"financial_services\",\"flight\",\"profile_management\",\"search\",\"shopping_gift_cards\"]"
}
```

Terraform syntax:

```terraform
flow_label {
  # Configure direct properties listed below.
}
```

<a id="canonical-e0281b52a30bede1d17428f066a9c2b988bcf714dcb47ae78236e76b79992f49"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label / 81d056f86e08 / 3

- [account_management](resources--http_loadbalancer--reference--group-011.md#canonical-3447a6c5b5b82efd85712c91b02fbb97c514a466b0b27151694abef61be2627e): complete subsection reference.

- [authentication](resources--http_loadbalancer--reference--group-011.md#canonical-d04125eeaa5a9529cdc402b8474fb1d0a7e57c3b20fa58fe1fedb66a068fdcf2): complete subsection reference.

- [financial_services](resources--http_loadbalancer--reference--group-011.md#canonical-b7f2a614aa37817fa4cfda14645805e918d31e14f88c9070fb7fe9bf0e244a56): complete subsection reference.

- [flight](resources--http_loadbalancer--reference--group-012.md#canonical-9d2291128668009bdee438ef7bdab823cf9e4a995b191a3ca18e9bd4e945aa5e): complete subsection reference.

- [profile_management](resources--http_loadbalancer--reference--group-012.md#canonical-3f7e8ce91a5a7a42ffcb08c818ed2ec58c3a5db7cd6d319ad9b2d00570648189): complete subsection reference.

- [search](resources--http_loadbalancer--reference--group-012.md#canonical-20f780669afe0a12646e944d9b2bddcdd5cd2a6c25288a422dcd320041ed2259): complete subsection reference.

- [shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-e44dccc6c0c5614f6c31032d1204ce89a5f066aa9c0c0a540d64dbfef2633d28): complete subsection reference.

<a id="canonical-c10a458b671faca33d9d72527ebf6031e32bfa5d25279d213d0c447db12a0a4f"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label / 81d056f86e08 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](resources--http_loadbalancer--reference--group-011.md#canonical-3447a6c5b5b82efd85712c91b02fbb97c514a466b0b27151694abef61be2627e)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-011.md#canonical-d04125eeaa5a9529cdc402b8474fb1d0a7e57c3b20fa58fe1fedb66a068fdcf2)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](resources--http_loadbalancer--reference--group-011.md#canonical-b7f2a614aa37817fa4cfda14645805e918d31e14f88c9070fb7fe9bf0e244a56)
- [bot_defense.policy.protected_app_endpoints.flow_label.flight](resources--http_loadbalancer--reference--group-012.md#canonical-9d2291128668009bdee438ef7bdab823cf9e4a995b191a3ca18e9bd4e945aa5e)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--http_loadbalancer--reference--group-012.md#canonical-3f7e8ce91a5a7a42ffcb08c818ed2ec58c3a5db7cd6d319ad9b2d00570648189)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--http_loadbalancer--reference--group-012.md#canonical-20f780669afe0a12646e944d9b2bddcdd5cd2a6c25288a422dcd320041ed2259)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-e44dccc6c0c5614f6c31032d1204ce89a5f066aa9c0c0a540d64dbfef2633d28)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-3447a6c5b5b82efd85712c91b02fbb97c514a466b0b27151694abef61be2627e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa2227dbe5a5ec0bd2d4bfc2863ced5604e4cf0cbc0de8464cd9a6c15a90db03"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.account_management — bot_defense.policy.protected_app_endpoints.flow_label.account_management / fc85449c0509 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management

<a id="canonical-41cfe89f7172e5fc7e6857ea089a3552d466959af676d0afddc9a7ef42c20cf4"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Account Management Category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("create",
    "password_reset")}
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
  "x-ves-oneof-field-label_choice": "[\"create\",\"password_reset\"]"
}
```

Terraform syntax:

```terraform
account_management {
  # Configure direct properties listed below.
}
```

<a id="canonical-2129f2d009d7f6cba5f7dd07f1ba707d57dc65609687b16ac89d4534f82e40b1"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.account_management / fc85449c0509 / 3

- [create](resources--http_loadbalancer--reference--group-011.md#canonical-b1205de3e69076546d3a0aaf13f2621f413195792b32716e128496d3fcb3b2ab): complete subsection reference.

- [password_reset](resources--http_loadbalancer--reference--group-011.md#canonical-1a0423cf286bd8134e34c30fd85dea2a40ee8ef0852cdc489bbbaabc11b68909): complete subsection reference.

<a id="canonical-a01bcf6ee9f69c5fc9aa8efd7cf76148a060646bb7f0a7a90a959222ccd34bce"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.account_management / fc85449c0509 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management.create](resources--http_loadbalancer--reference--group-011.md#canonical-b1205de3e69076546d3a0aaf13f2621f413195792b32716e128496d3fcb3b2ab)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset](resources--http_loadbalancer--reference--group-011.md#canonical-1a0423cf286bd8134e34c30fd85dea2a40ee8ef0852cdc489bbbaabc11b68909)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b1205de3e69076546d3a0aaf13f2621f413195792b32716e128496d3fcb3b2ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8c7d0e050e4b065cca4d35b293b12f5f92a504072957330f960d11326a0c505e"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.account_management.create — bot_defense.policy.protected_app_endpoints.flow_label.account_management.create / 679211881537 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](resources--http_loadbalancer--reference--group-011.md#canonical-3447a6c5b5b82efd85712c91b02fbb97c514a466b0b27151694abef61be2627e)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management.create

<a id="canonical-0cec21265e12191caf50898b283b22da3e75d27b858dbdb12b64f3b502b1d53d"></a>

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
create = {}
```

<a id="canonical-b8493cd4a4507927a0a6ac077982d4d8581353f214aae03f1008d3ab3f39ef60"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.account_management.create / 679211881537 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5ab69fc7ee81d6e1688ae874fd09abe40f8bf55c7dd5bdf3c1126318b261c250"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.account_management.create / 679211881537 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](resources--http_loadbalancer--reference--group-011.md#canonical-3447a6c5b5b82efd85712c91b02fbb97c514a466b0b27151694abef61be2627e)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-1a0423cf286bd8134e34c30fd85dea2a40ee8ef0852cdc489bbbaabc11b68909"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a66db1273f0febbce63c13dd397ae80f36dbf8e2e61966231ade4ad524db1a1"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset — bot_defense.policy.protected_app_endpoints.flow_label.account_management.passwor / 6fec7b8bec65 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](resources--http_loadbalancer--reference--group-011.md#canonical-3447a6c5b5b82efd85712c91b02fbb97c514a466b0b27151694abef61be2627e)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset

<a id="canonical-45380df17acde1d46ef9ffc7e87bac9ec77bebc59b5ea721fc70326766318709"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for password reset.

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
password_reset = {}
```

<a id="canonical-df33fdf19b344e38ca1e821f079d208835d163729c5b3a82e2dbad29ff00c155"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.account_management.passwor / 6fec7b8bec65 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d827a2d1eef9c4a2ae35d8975488e5e95b1d68faf63bfaf6cb864e5e897aa052"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.account_management.passwor / 6fec7b8bec65 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](resources--http_loadbalancer--reference--group-011.md#canonical-3447a6c5b5b82efd85712c91b02fbb97c514a466b0b27151694abef61be2627e)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d04125eeaa5a9529cdc402b8474fb1d0a7e57c3b20fa58fe1fedb66a068fdcf2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-444272b2f2b9cea528e2ada072a0e4cd6683ef40eea9505bf33ad155163609a8"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication — bot_defense.policy.protected_app_endpoints.flow_label.authentication / 2a2082dfe73e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication

<a id="canonical-5ff8dd7a75237314045b4e2abae193d78895cff00304b4becb80c534350d671f"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Authentication Category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("login",
    "login_mfa"),
  validators.ConflictingObjectAttributes("login",
    "login_partner"),
  validators.ConflictingObjectAttributes("login",
    "logout"),
  validators.ConflictingObjectAttributes("login",
    "token_refresh"),
  validators.ConflictingObjectAttributes("login_mfa",
    "login_partner"),
  validators.ConflictingObjectAttributes("login_mfa",
    "logout"),
  validators.ConflictingObjectAttributes("login_mfa",
    "token_refresh"),
  validators.ConflictingObjectAttributes("login_partner",
    "logout"),
  validators.ConflictingObjectAttributes("login_partner",
    "token_refresh"),
  validators.ConflictingObjectAttributes("logout",
    "token_refresh")}
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
  "x-ves-oneof-field-label_choice": "[\"login\",\"login_mfa\",\"login_partner\",\"logout\",\"token_refresh\"]"
}
```

Terraform syntax:

```terraform
authentication {
  # Configure direct properties listed below.
}
```

<a id="canonical-f9eea748d2a95d35a39d532b8c9dbc46a7e05f0ab016c195375f27bafe99e145"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.authentication / 2a2082dfe73e / 3

- [login](resources--http_loadbalancer--reference--group-011.md#canonical-1c69355295dc2911462334fedcf45b10f7faeb7dd2bfad7b2f25595d04f808df): complete subsection reference.

- [login_mfa](resources--http_loadbalancer--reference--group-011.md#canonical-b45c074ba05489519090e4a33ec3e63ac06b7e1b22acc8c958b63bb5c34b7650): complete subsection reference.

- [login_partner](resources--http_loadbalancer--reference--group-011.md#canonical-e8fda43dfc569db51ff1b3645156ad968558b4c847d2e0dc3b12a7b4935430a2): complete subsection reference.

- [logout](resources--http_loadbalancer--reference--group-011.md#canonical-49bf3acbd55dd68e9e904ea33899d9fde7329ed7a554af1b71c121bf1eefff79): complete subsection reference.

- [token_refresh](resources--http_loadbalancer--reference--group-011.md#canonical-a72bc253261fdd17a21242ef0838eb1a3d78e7f457e550cfa0e4e2b11776040f): complete subsection reference.

<a id="canonical-63d73a17f9880de10ecee85dc59cb2596c6ffa13c0affafa4972d838a125f9dc"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.authentication / 2a2082dfe73e / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](resources--http_loadbalancer--reference--group-011.md#canonical-1c69355295dc2911462334fedcf45b10f7faeb7dd2bfad7b2f25595d04f808df)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa](resources--http_loadbalancer--reference--group-011.md#canonical-b45c074ba05489519090e4a33ec3e63ac06b7e1b22acc8c958b63bb5c34b7650)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partner](resources--http_loadbalancer--reference--group-011.md#canonical-e8fda43dfc569db51ff1b3645156ad968558b4c847d2e0dc3b12a7b4935430a2)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.logout](resources--http_loadbalancer--reference--group-011.md#canonical-49bf3acbd55dd68e9e904ea33899d9fde7329ed7a554af1b71c121bf1eefff79)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refresh](resources--http_loadbalancer--reference--group-011.md#canonical-a72bc253261fdd17a21242ef0838eb1a3d78e7f457e550cfa0e4e2b11776040f)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-1c69355295dc2911462334fedcf45b10f7faeb7dd2bfad7b2f25595d04f808df"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-19092ffd3aa03657e4df40f5874ed2a5a727478421fa26b0f552560d13d71aa6"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login / 1b533e452cf9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-011.md#canonical-d04125eeaa5a9529cdc402b8474fb1d0a7e57c3b20fa58fe1fedb66a068fdcf2)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login

<a id="canonical-802dfb9a0d370b1e43fce590eb47fd354a39ba91b14bb8d6edd43ea85b776c98"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Transaction Result. Bot Defense Transaction Result.

Upstream description:

Bot Defense Transaction Result.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_transaction_result",
    "transaction_result")}
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
  "x-ves-oneof-field-transaction_result_choice": "[\"disable_transaction_result\",\"transaction_result\"]"
}
```

Terraform syntax:

```terraform
login {
  # Configure direct properties listed below.
}
```

<a id="canonical-1212f590a78f98de01e7d75638c9a425d499185d6e032e547b91429ef928ff2a"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login / 1b533e452cf9 / 3

- [disable_transaction_result](resources--http_loadbalancer--reference--group-011.md#canonical-82bd4ab0cb51dfe0c2bb57c23022525034dc1b79866fe5c8a1aff9d35b162961): complete subsection reference.

- [transaction_result](resources--http_loadbalancer--reference--group-011.md#canonical-85b69f32f6caac6b200b31ce1538f4dbcf27fd89b62b3d77bdc4bd7d3b167710): complete subsection reference.

<a id="canonical-0e6c1c1ad48dc7346b86111d1b0fe80597d3b8890d9877cf9bcc7a36e375462d"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login / 1b533e452cf9 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disable_transaction_result](resources--http_loadbalancer--reference--group-011.md#canonical-82bd4ab0cb51dfe0c2bb57c23022525034dc1b79866fe5c8a1aff9d35b162961)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](resources--http_loadbalancer--reference--group-011.md#canonical-85b69f32f6caac6b200b31ce1538f4dbcf27fd89b62b3d77bdc4bd7d3b167710)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-011.md#canonical-d04125eeaa5a9529cdc402b8474fb1d0a7e57c3b20fa58fe1fedb66a068fdcf2)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-82bd4ab0cb51dfe0c2bb57c23022525034dc1b79866fe5c8a1aff9d35b162961"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4c90114ec7b0d9a9905d97b9d6d82c19e919eddcbd1f90bee38f2a3f7e628a45"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disable_transaction_result — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disab / 94177bd0d3d3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-011.md#canonical-d04125eeaa5a9529cdc402b8474fb1d0a7e57c3b20fa58fe1fedb66a068fdcf2)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](resources--http_loadbalancer--reference--group-011.md#canonical-1c69355295dc2911462334fedcf45b10f7faeb7dd2bfad7b2f25595d04f808df)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disable_transaction_result

<a id="canonical-6c986f68e1018b97bb384cbbbd5848cd6f87f564bbfc6ab52b89a5a92cc853cb"></a>

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
disable_transaction_result = {}
```

<a id="canonical-fd90b013ee0725bdefef3f56f7126b55fb02f32e7daab95ec330c6fd6918c6a8"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disab / 94177bd0d3d3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-74b690d6a4ecd27e4e4016ca23b67786e4f032a03031704aab62ee498e60d408"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disab / 94177bd0d3d3 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](resources--http_loadbalancer--reference--group-011.md#canonical-1c69355295dc2911462334fedcf45b10f7faeb7dd2bfad7b2f25595d04f808df)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-85b69f32f6caac6b200b31ce1538f4dbcf27fd89b62b3d77bdc4bd7d3b167710"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-65b163d144893fd8cb80fa4dcdfb6940843e6df99f26f5d829a17e74aa3440e6"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / 5d7f2f5c9b1c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-011.md#canonical-d04125eeaa5a9529cdc402b8474fb1d0a7e57c3b20fa58fe1fedb66a068fdcf2)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](resources--http_loadbalancer--reference--group-011.md#canonical-1c69355295dc2911462334fedcf45b10f7faeb7dd2bfad7b2f25595d04f808df)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result

<a id="canonical-a78ce9cf0fb5ad2a90e550021265810f7bdddd9368fb1e644207b35be507a0c3"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Transaction Result Type. Bot Defense Transaction ResultType.

Upstream description:

Bot Defense Transaction ResultType.

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
transaction_result {
  # Configure direct properties listed below.
}
```

<a id="canonical-eaa0c56b494bb321b4328ab1890dc0fad5f02f12863b1e289b1a478d2cb67991"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / 5d7f2f5c9b1c / 3

- [failure_conditions](resources--http_loadbalancer--reference--group-011.md#canonical-ee94ee375c37a3cba91b7a4f2170b3f2f529eaf59e473f640fbf0be7ea4f73b1): complete subsection reference.

- [success_conditions](resources--http_loadbalancer--reference--group-011.md#canonical-99992df3b7114df69b969d802f67030f08ef08366e368e3403ecf9a3efbd3f4c): complete subsection reference.

<a id="canonical-12ed83c317a9a10e11632bdafc30a0f86a437286583ec42896f3c99a1feb0808"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / 5d7f2f5c9b1c / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions](resources--http_loadbalancer--reference--group-011.md#canonical-ee94ee375c37a3cba91b7a4f2170b3f2f529eaf59e473f640fbf0be7ea4f73b1)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions](resources--http_loadbalancer--reference--group-011.md#canonical-99992df3b7114df69b969d802f67030f08ef08366e368e3403ecf9a3efbd3f4c)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](resources--http_loadbalancer--reference--group-011.md#canonical-1c69355295dc2911462334fedcf45b10f7faeb7dd2bfad7b2f25595d04f808df)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-ee94ee375c37a3cba91b7a4f2170b3f2f529eaf59e473f640fbf0be7ea4f73b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23c55a3a611669fffde1fa9ce2e733b2241281536f9aeacd6ca86074fb35ef67"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / bf899a5c31ac / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-011.md#canonical-d04125eeaa5a9529cdc402b8474fb1d0a7e57c3b20fa58fe1fedb66a068fdcf2)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](resources--http_loadbalancer--reference--group-011.md#canonical-1c69355295dc2911462334fedcf45b10f7faeb7dd2bfad7b2f25595d04f808df)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](resources--http_loadbalancer--reference--group-011.md#canonical-85b69f32f6caac6b200b31ce1538f4dbcf27fd89b62b3d77bdc4bd7d3b167710)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions

<a id="canonical-680429980728f01ab6eef8b964c1647522532d82870d26014e0443218b905121"></a>

Type: `"object"`. list nested block, Optional.

Failure Conditions. Failure Conditions.

Upstream description:

Failure Conditions.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
failure_conditions {
  # Configure direct properties listed below.
}
```

<a id="canonical-980053e3f52d3cae8867dca5868799897f4f540b5c3d037f1d71d987ceb6a59e"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / bf899a5c31ac / 3

<a id="canonical-10a677c8036e9d3cc3a04b798ccde8339110a33d10527d59a7f949d4b4e1b99e"></a>

<a id="canonical-8916e34b954433b00657202bd75ad5ab6a0d073978e57dffda18f9d8f88ad79f"></a>

## name property — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / bf899a5c31ac / 4

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3d43c240d9b598a838914c7509da0c32c9127181c57ef2a6e1bd5f0de238e392"></a>

<a id="canonical-208d377596526f59ff6171c49feb4d94bc8f2330ee42ea429b68fe03b0b58394"></a>

## regex_values property — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / bf899a5c31ac / 5

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
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-966459a64a57a8108fd977a8c30617b400abda1d27eb3d2dc00046c86efcf4f8"></a>

<a id="canonical-8f448b47536ed2931e1d4a01ff680da9bacd54c56fcbe4c588f37c181a0aa3a2"></a>

## status property — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / bf899a5c31ac / 6

Type: `"string"`. Optional.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Upstream description:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-4ad4f1fe47cec9fd03d9cdc2743f9397bf54f37dd4265f0a58c425ca8b3cc44d"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / bf899a5c31ac / 7

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](resources--http_loadbalancer--reference--group-011.md#canonical-85b69f32f6caac6b200b31ce1538f4dbcf27fd89b62b3d77bdc4bd7d3b167710)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-99992df3b7114df69b969d802f67030f08ef08366e368e3403ecf9a3efbd3f4c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a0619441be60f3d4158c02c104851a612b311b21113784a78bd0de3f0d774b0"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / 4bf77c6c8ba9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-011.md#canonical-d04125eeaa5a9529cdc402b8474fb1d0a7e57c3b20fa58fe1fedb66a068fdcf2)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](resources--http_loadbalancer--reference--group-011.md#canonical-1c69355295dc2911462334fedcf45b10f7faeb7dd2bfad7b2f25595d04f808df)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](resources--http_loadbalancer--reference--group-011.md#canonical-85b69f32f6caac6b200b31ce1538f4dbcf27fd89b62b3d77bdc4bd7d3b167710)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions

<a id="canonical-1e6df7c32adbd02dc21325e97c7830976f7d068b66f12969c5a0ce40ec5c806a"></a>

Type: `"object"`. list nested block, Optional.

Success Conditions. Success Conditions.

Upstream description:

Success Conditions.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
success_conditions {
  # Configure direct properties listed below.
}
```

<a id="canonical-0af4e962b2a4f2d97a2956b60f8a39759ef1958ff585eca2540f3be192d2898a"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / 4bf77c6c8ba9 / 3

<a id="canonical-74982512c9d1ddb2460b20b1a4c000609e5764de14d7bf9a84fcf1cbe17d985b"></a>

<a id="canonical-c96e51ec27996caf4ae6a85e230ffe1e622b1c66a667fad84d390d43ed9a154e"></a>

## name property — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / 4bf77c6c8ba9 / 4

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-709952583e86fbd0fb8b447a8c0a8ac567e1fece5739486530f3dde85a394f6f"></a>

<a id="canonical-38daa207b04703e728e38b7d9749940a1a482e89eedd9328a3466e6094100b74"></a>

## regex_values property — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / 4bf77c6c8ba9 / 5

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
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-07eaf209d3c8b9843d84055a5adf13afe6ff5c35468f274e537d1906dc2fd972"></a>

<a id="canonical-d35ee1697170ea83807f3d2b207d6863958132923fe6d264aeb59e2f0eb4299e"></a>

## status property — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / 4bf77c6c8ba9 / 6

Type: `"string"`. Optional.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Upstream description:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-241b206f338b99c5a5606fc0b713e34fb2b55031913b9cb56a78d0a9b46a72bc"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / 4bf77c6c8ba9 / 7

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](resources--http_loadbalancer--reference--group-011.md#canonical-85b69f32f6caac6b200b31ce1538f4dbcf27fd89b62b3d77bdc4bd7d3b167710)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b45c074ba05489519090e4a33ec3e63ac06b7e1b22acc8c958b63bb5c34b7650"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ff9d87f78ba6b05f809df753b389561a9c68cb031d13253bffbd22f6aa8114e4"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa / bbfe67238087 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-011.md#canonical-d04125eeaa5a9529cdc402b8474fb1d0a7e57c3b20fa58fe1fedb66a068fdcf2)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa

<a id="canonical-edb1585d5ea729adc769f72a893d0041de6e595f2931309642f0d227c5bf9efe"></a>

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
login_mfa = {}
```

<a id="canonical-489ce854b8a02716c41c45a628e09f0adb7325e3e7669683b662d173dbf72559"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa / bbfe67238087 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-76a0c7cafdff91f4fa5954915148ca924ff2ff5dbef4d9f63ac19624e39377d7"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa / bbfe67238087 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-011.md#canonical-d04125eeaa5a9529cdc402b8474fb1d0a7e57c3b20fa58fe1fedb66a068fdcf2)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e8fda43dfc569db51ff1b3645156ad968558b4c847d2e0dc3b12a7b4935430a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8e3d02b111b007958db8632a69560936e67c2decde14f779230e91ab9eaad3c"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partner — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partn / 3e9afa49d3e1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-011.md#canonical-d04125eeaa5a9529cdc402b8474fb1d0a7e57c3b20fa58fe1fedb66a068fdcf2)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partner

<a id="canonical-7a52fcb5d25b110cdbdd46d745c2eb7d39765bec10b2506058f2e3e4084c49da"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for login partner.

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
login_partner = {}
```

<a id="canonical-0d97b77ee6981b8844bed6807c28b4bc4bfd59bfca2e7da98f07163be75c1677"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partn / 3e9afa49d3e1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-da9486defd74487a2bfa87faed27b1c3a3dd6e2995594302cb456363c47a4489"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partn / 3e9afa49d3e1 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-011.md#canonical-d04125eeaa5a9529cdc402b8474fb1d0a7e57c3b20fa58fe1fedb66a068fdcf2)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-49bf3acbd55dd68e9e904ea33899d9fde7329ed7a554af1b71c121bf1eefff79"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be3e082689b0df6c6d168d707d1b31a1f39deacb9d3c4d40eeb560846cbe8a5f"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.logout — bot_defense.policy.protected_app_endpoints.flow_label.authentication.logout / e91299bb32ec / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-011.md#canonical-d04125eeaa5a9529cdc402b8474fb1d0a7e57c3b20fa58fe1fedb66a068fdcf2)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.logout

<a id="canonical-565e293d351ec466a7174372fe4e1f6865e22cfa9c5b1f628fd3aa7e9f5d4386"></a>

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
logout = {}
```

<a id="canonical-f78b5db8c5b7290e484ea6d5286187a2272d23aeee0d43925064be6775199e4a"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.authentication.logout / e91299bb32ec / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ed2add9d3810bb65fb5640a2be9e4e693a09c747072ff0ede0d3d263c16c40c1"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.authentication.logout / e91299bb32ec / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-011.md#canonical-d04125eeaa5a9529cdc402b8474fb1d0a7e57c3b20fa58fe1fedb66a068fdcf2)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-a72bc253261fdd17a21242ef0838eb1a3d78e7f457e550cfa0e4e2b11776040f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-85e6b899d3f1549bb5da13eb88bfcecfe0055f4db5fec636baa869e86890f0c7"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refresh — bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refre / 2e2d3cb1b520 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-011.md#canonical-d04125eeaa5a9529cdc402b8474fb1d0a7e57c3b20fa58fe1fedb66a068fdcf2)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refresh

<a id="canonical-78ec08849942f0fac767a27386d6547193e957669b8424b349bc8bce75701c56"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for token refresh.

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
token_refresh = {}
```

<a id="canonical-64c6a6342f1bd6eeab41ece1775df9e7d4126ff6d6534f413d66c7c3344af6d8"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refre / 2e2d3cb1b520 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3ef68f1ecc59eb86112bc8b55aa235305148c5f5fcf6199174e87164110f5f5d"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refre / 2e2d3cb1b520 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-011.md#canonical-d04125eeaa5a9529cdc402b8474fb1d0a7e57c3b20fa58fe1fedb66a068fdcf2)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b7f2a614aa37817fa4cfda14645805e918d31e14f88c9070fb7fe9bf0e244a56"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
