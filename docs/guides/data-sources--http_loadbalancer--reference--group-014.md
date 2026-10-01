---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-d01008cfa430174287a812da0e8b3f22c12ac4e9f12d7b9137b6d61e6bde8c61"></a>

## client_side_defense.policy.js_insert_all_pages_except — client_side_defense.policy.js_insert_all_pages_except / 7e289f4f3479 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-013.md#canonical-367f873504986262d380bdfbd173acad4137578e1de683d131b3d0248b345e6f)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-187bc8e8a39a49b1be57c0106b3ea4eb53acefd0a6616f1a35708cf01ec7e22b)
- client_side_defense.policy.js_insert_all_pages_except

<a id="canonical-09822b8db8b3c7d0110b8c80ec5b9f6c51d672f3f27f7086fcbb93471585061d"></a>

Type: `"single"`. Computed.

Insert Client-Side Defense JavaScript in all pages with the exceptions.

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

<a id="canonical-1cf70c2eb459a77961da057b237e8aff663f4751db914f0b421e365416ac2153"></a>

## Direct properties — client_side_defense.policy.js_insert_all_pages_except / 7e289f4f3479 / 3

- [exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-c1749e904f25005d0e35c86c47c24f3b18c74e798dceded9f1ea751b6ba6294c): complete subsection reference.

<a id="canonical-d0b1e3216a37cf6014104612dfec2dda9d2c47abc63db799dca6f819b1b6cc09"></a>

## Next pages — client_side_defense.policy.js_insert_all_pages_except / 7e289f4f3479 / 4

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-c1749e904f25005d0e35c86c47c24f3b18c74e798dceded9f1ea751b6ba6294c)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-187bc8e8a39a49b1be57c0106b3ea4eb53acefd0a6616f1a35708cf01ec7e22b)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c1749e904f25005d0e35c86c47c24f3b18c74e798dceded9f1ea751b6ba6294c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ed3289a09d0777c150d423e1a2ad9b2933347cfb70361c0fde7a3adc4716339"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list — client_side_defense.policy.js_insert_all_pages_except.exclude_list / 50b850bb2e1d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-013.md#canonical-367f873504986262d380bdfbd173acad4137578e1de683d131b3d0248b345e6f)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-187bc8e8a39a49b1be57c0106b3ea4eb53acefd0a6616f1a35708cf01ec7e22b)
- [client_side_defense.policy.js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-013.md#canonical-fec9e6f94146815bcb028735c7065d5e00d72fa1a44bd123c68f5e2e592221b6)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list

<a id="canonical-08abba9370e482157b006a299ca5a88e3c4b059e6468068b72f2d523e974601c"></a>

Type: `"list"`. Computed.

Optional JavaScript insertions exclude list of domain and path matchers.

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

<a id="canonical-585993ab790248def3f610b60b854aa7f2911499887b8a32a9232b7193c2ab66"></a>

## Direct properties — client_side_defense.policy.js_insert_all_pages_except.exclude_list / 50b850bb2e1d / 3

- [any_domain](data-sources--http_loadbalancer--reference--group-014.md#canonical-42ad6be01503f55ab419932c74d422a824f083e9181ceca24ff171b55a6b0d12): complete subsection reference.

- [domain](data-sources--http_loadbalancer--reference--group-014.md#canonical-82774a4805fd7ca7ba351370fe6ce992af5cbbfc24ac4e2fab21f9142b24c805): complete subsection reference.

- [metadata](data-sources--http_loadbalancer--reference--group-014.md#canonical-4e257480ce673f0f3e1fc670388d633d6306c57133e3cb75bd5457a5efc7ace9): complete subsection reference.

- [path](data-sources--http_loadbalancer--reference--group-014.md#canonical-02d6a34938ad66fb478877d871776d7880e012ad6c591dbb6c98fbc6563c00a2): complete subsection reference.

<a id="canonical-3bebbb16ed0f1a535f1b1276ad755bdf29b1eacd9f3a8c673b3f20602d08eaaf"></a>

## Next pages — client_side_defense.policy.js_insert_all_pages_except.exclude_list / 50b850bb2e1d / 4

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain](data-sources--http_loadbalancer--reference--group-014.md#canonical-42ad6be01503f55ab419932c74d422a824f083e9181ceca24ff171b55a6b0d12)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain](data-sources--http_loadbalancer--reference--group-014.md#canonical-82774a4805fd7ca7ba351370fe6ce992af5cbbfc24ac4e2fab21f9142b24c805)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata](data-sources--http_loadbalancer--reference--group-014.md#canonical-4e257480ce673f0f3e1fc670388d633d6306c57133e3cb75bd5457a5efc7ace9)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.path](data-sources--http_loadbalancer--reference--group-014.md#canonical-02d6a34938ad66fb478877d871776d7880e012ad6c591dbb6c98fbc6563c00a2)
- [client_side_defense.policy.js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-013.md#canonical-fec9e6f94146815bcb028735c7065d5e00d72fa1a44bd123c68f5e2e592221b6)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-42ad6be01503f55ab419932c74d422a824f083e9181ceca24ff171b55a6b0d12"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8aece84a05ab03359a7d1e0cdd7fa202ba5179b1e885adbe96e73fa0a76d24e1"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain — client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain / c1f7f43e30e7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-013.md#canonical-367f873504986262d380bdfbd173acad4137578e1de683d131b3d0248b345e6f)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-187bc8e8a39a49b1be57c0106b3ea4eb53acefd0a6616f1a35708cf01ec7e22b)
- [client_side_defense.policy.js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-013.md#canonical-fec9e6f94146815bcb028735c7065d5e00d72fa1a44bd123c68f5e2e592221b6)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-c1749e904f25005d0e35c86c47c24f3b18c74e798dceded9f1ea751b6ba6294c)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain

<a id="canonical-fbef500a294a699e45450cd76a86637843636bd04fe2fbebcaebfebb2648452b"></a>

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

<a id="canonical-2f13fa0d889edd273349413ac84387db06abfc5b90bddd2bbda51529aecdc561"></a>

## Direct properties — client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain / c1f7f43e30e7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-64a4794275a00b268d850f64a13e74ea16589dcc2687a7e4ecc7eb12a2035f28"></a>

## Next pages — client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain / c1f7f43e30e7 / 4

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-c1749e904f25005d0e35c86c47c24f3b18c74e798dceded9f1ea751b6ba6294c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-82774a4805fd7ca7ba351370fe6ce992af5cbbfc24ac4e2fab21f9142b24c805"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-526194e0e7dcdece186567d08444be7a79e104826392e89ce8f0d15479e3ddf4"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain — client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain / ee29f7d3ff25 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-013.md#canonical-367f873504986262d380bdfbd173acad4137578e1de683d131b3d0248b345e6f)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-187bc8e8a39a49b1be57c0106b3ea4eb53acefd0a6616f1a35708cf01ec7e22b)
- [client_side_defense.policy.js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-013.md#canonical-fec9e6f94146815bcb028735c7065d5e00d72fa1a44bd123c68f5e2e592221b6)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-c1749e904f25005d0e35c86c47c24f3b18c74e798dceded9f1ea751b6ba6294c)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain

<a id="canonical-77a300aa0fe2f18f4f15f546f285b66cb59e0f1daa07ec9d8edf6fd49a030de0"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Upstream description:

Domains names.

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

<a id="canonical-8544ab849801e60077b1c5995b03aa86c0906699b36e46e13588080607eace6d"></a>

## Direct properties — client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain / ee29f7d3ff25 / 3

<a id="canonical-d10c7fef9144fd9cd301a7d6aa5a26afbc90b37371e8a717cc83a5424784d95c"></a>

<a id="canonical-307c3a40836398d45abcc44ffc8b64537a971c428efb38cf1f261b89c99e0763"></a>

## exact_value property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain / ee29f7d3ff25 / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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

<a id="canonical-fe514cb8f3a451d69d04636a6a426cabadb97e2e9cdf381f1e19cb73e862e5b4"></a>

<a id="canonical-59ede09dc8f793d62922a82c18ecc360b26cd675e379680cfd094e17e4b26644"></a>

## regex_value property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain / ee29f7d3ff25 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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

<a id="canonical-9f1379d10b035266858b23361d6db2d57a9be992f35e79321c58491d7a320fd7"></a>

<a id="canonical-0eda9a18d2620ae9e54cbff9523a16facd82928cd74182f55d87ddd16dbe7aeb"></a>

## suffix_value property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain / ee29f7d3ff25 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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

<a id="canonical-4c9e43846093a620789b60da7b9db8c7399b3144588e4c34d27664c726af09e0"></a>

## Next pages — client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain / ee29f7d3ff25 / 7

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-c1749e904f25005d0e35c86c47c24f3b18c74e798dceded9f1ea751b6ba6294c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-4e257480ce673f0f3e1fc670388d633d6306c57133e3cb75bd5457a5efc7ace9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e5882091149f2e71275eee26193b93df7610ca90894097203ead5f961d0194a4"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata — client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata / cca68bf45c6c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-013.md#canonical-367f873504986262d380bdfbd173acad4137578e1de683d131b3d0248b345e6f)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-187bc8e8a39a49b1be57c0106b3ea4eb53acefd0a6616f1a35708cf01ec7e22b)
- [client_side_defense.policy.js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-013.md#canonical-fec9e6f94146815bcb028735c7065d5e00d72fa1a44bd123c68f5e2e592221b6)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-c1749e904f25005d0e35c86c47c24f3b18c74e798dceded9f1ea751b6ba6294c)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata

<a id="canonical-795d00c0830d1d9e94488598b0ca44fdbb449452d2d797b74c3001a1ad1bec4a"></a>

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

<a id="canonical-4d825fa256c3aa3152d2680e5dfd736a37fe5552ba9afd5e9e5cfad70ed9764c"></a>

## Direct properties — client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata / cca68bf45c6c / 3

<a id="canonical-753004b7ee7d9c55d375bc883e9dd925c80deb594c43071fec3ef28b2ee593f0"></a>

<a id="canonical-f0b085fc76147e6cb2ffc3e20b112773c5a1aff2ea7d42192863e74759572523"></a>

## description_spec property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata / cca68bf45c6c / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-eba2e57fac5e1e973173965f7cb2bb6c99340ef8d832088170aee06488406eb0"></a>

<a id="canonical-2e88bbee7dd6d38dc1557904879f394224d209472c665cf840596d1e28a65906"></a>

## name property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata / cca68bf45c6c / 5

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

<a id="canonical-64294f5d58b9114165e5733a5576a6b017e1513b45773512f19bd14d10e4ad62"></a>

## Next pages — client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata / cca68bf45c6c / 6

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-c1749e904f25005d0e35c86c47c24f3b18c74e798dceded9f1ea751b6ba6294c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-02d6a34938ad66fb478877d871776d7880e012ad6c591dbb6c98fbc6563c00a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-74e29b10f5453590950a2b770e3d06bfeec7e026d56f0e48c66de8ae48032d57"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list.path — client_side_defense.policy.js_insert_all_pages_except.exclude_list.path / 5c78ca59be19 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-013.md#canonical-367f873504986262d380bdfbd173acad4137578e1de683d131b3d0248b345e6f)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-187bc8e8a39a49b1be57c0106b3ea4eb53acefd0a6616f1a35708cf01ec7e22b)
- [client_side_defense.policy.js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-013.md#canonical-fec9e6f94146815bcb028735c7065d5e00d72fa1a44bd123c68f5e2e592221b6)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-c1749e904f25005d0e35c86c47c24f3b18c74e798dceded9f1ea751b6ba6294c)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.path

<a id="canonical-cc1cb462d8c40e6ef4ab5a1774c044f0c0926390f30071aef3fcaaae423bf4ee"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

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

<a id="canonical-e06f91d3728c889990703247da5abd5bbc625cadc2f629ec71a74bea69b7e69e"></a>

## Direct properties — client_side_defense.policy.js_insert_all_pages_except.exclude_list.path / 5c78ca59be19 / 3

<a id="canonical-af519ba44f6bde9b1ec1accda1a68503b3e40785c3fc52df8ce438351ba7bf09"></a>

<a id="canonical-18c4f3243fecc954d2a0299f84f9ad50fc3df09747253592d0c0c615877c5884"></a>

## path property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.path / 5c78ca59be19 / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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

<a id="canonical-f86d09c1998e37a8c316c6157109269907db3559dceb16a6057979ef6363178b"></a>

<a id="canonical-0e3fa1dface863af07332a69ce17485439387bf068455fe8119f61dcd8450654"></a>

## prefix property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.path / 5c78ca59be19 / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-c7e34fc2a2657d57b6bbc74492b139c8c1bdf0125313541fedafaf8cd220c241"></a>

<a id="canonical-5d0f3eeee4a50f3887966ec25eff569c2d72382efad06f385ce9bdbd5abc2533"></a>

## regex property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.path / 5c78ca59be19 / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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

<a id="canonical-c8595012552e7ba22d97a270b531050a08bfa2c15e3de8df58255d02ae9331f6"></a>

## Next pages — client_side_defense.policy.js_insert_all_pages_except.exclude_list.path / 5c78ca59be19 / 7

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-c1749e904f25005d0e35c86c47c24f3b18c74e798dceded9f1ea751b6ba6294c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-cddcc8b626cac1ca512cd921feaf1053e00e0ff45310b85b8e9292b433ea905e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-40ed97a0fa12775b203c000b5eb0e9c7506514dedd450b824478d9b3647e0b0b"></a>

## client_side_defense.policy.js_insertion_rules — client_side_defense.policy.js_insertion_rules / a8c7ec2044b3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-013.md#canonical-367f873504986262d380bdfbd173acad4137578e1de683d131b3d0248b345e6f)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-187bc8e8a39a49b1be57c0106b3ea4eb53acefd0a6616f1a35708cf01ec7e22b)
- client_side_defense.policy.js_insertion_rules

<a id="canonical-ce4753d9cf1adf5e3ddf02bb476ceef777619e1495a9977c9140661bff6f577e"></a>

Type: `"single"`. Computed.

Defines custom JavaScript insertion rules for Client-Side Defense Policy.

Upstream description:

This defines custom JavaScript insertion rules for Client-Side Defense Policy.

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

<a id="canonical-e9fa6c9cd7ee4d59439d5e775ff9664751bd2972bbb86a0f151a33949ec5d98e"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules / a8c7ec2044b3 / 3

- [exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-6975d51072bab1120ea34c36e147af6db09b91779c1cb13eedb7c8bbae9de1db): complete subsection reference.

- [rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-02713d62645a8166d0c9fca71158dd7209f07d13c2af4dfd6d1cedaa5170fe71): complete subsection reference.

<a id="canonical-885eb48dfe3e64e497956c318a4ca57da97d2d1c219c8d9c202f4e7224619a67"></a>

## Next pages — client_side_defense.policy.js_insertion_rules / a8c7ec2044b3 / 4

- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-6975d51072bab1120ea34c36e147af6db09b91779c1cb13eedb7c8bbae9de1db)
- [client_side_defense.policy.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-02713d62645a8166d0c9fca71158dd7209f07d13c2af4dfd6d1cedaa5170fe71)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-187bc8e8a39a49b1be57c0106b3ea4eb53acefd0a6616f1a35708cf01ec7e22b)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6975d51072bab1120ea34c36e147af6db09b91779c1cb13eedb7c8bbae9de1db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-848ff98397c18b0062a8c50d4b942c44b50aea56221520b1b05355f3621415cd"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list — client_side_defense.policy.js_insertion_rules.exclude_list / 71b48006ad68 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-013.md#canonical-367f873504986262d380bdfbd173acad4137578e1de683d131b3d0248b345e6f)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-187bc8e8a39a49b1be57c0106b3ea4eb53acefd0a6616f1a35708cf01ec7e22b)
- [client_side_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-cddcc8b626cac1ca512cd921feaf1053e00e0ff45310b85b8e9292b433ea905e)
- client_side_defense.policy.js_insertion_rules.exclude_list

<a id="canonical-70de2e201263b72905f468e1c88eccd77f56676efbec9afb9dcd349f0655440a"></a>

Type: `"list"`. Computed.

Optional JavaScript insertions exclude list of domain and path matchers.

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

<a id="canonical-a00eda6f116fe1d162b00e3a591d0cea21292ae185a7e48ec9a4723986612d26"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.exclude_list / 71b48006ad68 / 3

- [any_domain](data-sources--http_loadbalancer--reference--group-014.md#canonical-fe2aa5c3bd6f946495ab6c9fffd0ee0bfb2c3e184c68f8a92e15e98e46bdabb9): complete subsection reference.

- [domain](data-sources--http_loadbalancer--reference--group-014.md#canonical-fb514d22d5020b4e8322f5139d94a08eccd6003c39b871558da685b9c7827f85): complete subsection reference.

- [metadata](data-sources--http_loadbalancer--reference--group-014.md#canonical-f4c48b51d2ba5dcc370d93d85d46e3e9fd068722c90aa1542b13bd30877a83b9): complete subsection reference.

- [path](data-sources--http_loadbalancer--reference--group-014.md#canonical-5298aa4e5e8e2e9cf2aaf88026cd4a4165a16593606644ac9f7a6d1bb75fca32): complete subsection reference.

<a id="canonical-55864de0b56b0adc771e0f34734fb2b6de45f5a6b042e36fbc38447cb0f7b0d7"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.exclude_list / 71b48006ad68 / 4

- [client_side_defense.policy.js_insertion_rules.exclude_list.any_domain](data-sources--http_loadbalancer--reference--group-014.md#canonical-fe2aa5c3bd6f946495ab6c9fffd0ee0bfb2c3e184c68f8a92e15e98e46bdabb9)
- [client_side_defense.policy.js_insertion_rules.exclude_list.domain](data-sources--http_loadbalancer--reference--group-014.md#canonical-fb514d22d5020b4e8322f5139d94a08eccd6003c39b871558da685b9c7827f85)
- [client_side_defense.policy.js_insertion_rules.exclude_list.metadata](data-sources--http_loadbalancer--reference--group-014.md#canonical-f4c48b51d2ba5dcc370d93d85d46e3e9fd068722c90aa1542b13bd30877a83b9)
- [client_side_defense.policy.js_insertion_rules.exclude_list.path](data-sources--http_loadbalancer--reference--group-014.md#canonical-5298aa4e5e8e2e9cf2aaf88026cd4a4165a16593606644ac9f7a6d1bb75fca32)
- [client_side_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-cddcc8b626cac1ca512cd921feaf1053e00e0ff45310b85b8e9292b433ea905e)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-fe2aa5c3bd6f946495ab6c9fffd0ee0bfb2c3e184c68f8a92e15e98e46bdabb9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d13701b10d2d91e70f8400916ae8a0f38f89a614b8885b8aa07754cf7b819ca9"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list.any_domain — client_side_defense.policy.js_insertion_rules.exclude_list.any_domain / 128aec1124d3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-013.md#canonical-367f873504986262d380bdfbd173acad4137578e1de683d131b3d0248b345e6f)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-187bc8e8a39a49b1be57c0106b3ea4eb53acefd0a6616f1a35708cf01ec7e22b)
- [client_side_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-cddcc8b626cac1ca512cd921feaf1053e00e0ff45310b85b8e9292b433ea905e)
- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-6975d51072bab1120ea34c36e147af6db09b91779c1cb13eedb7c8bbae9de1db)
- client_side_defense.policy.js_insertion_rules.exclude_list.any_domain

<a id="canonical-f7c43b255e7f78522179f64c3810474c1e6ab3471e81da9b5424a7a7aec2fcdd"></a>

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

<a id="canonical-8c24d7c2abbc62bf39255219e291e285b9a6a087fd49bcc0a45da5775142ebbd"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.exclude_list.any_domain / 128aec1124d3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ecad41a030303379555b0ad355f065d4fcdf7e92c5225d5014fde2826f1731b0"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.exclude_list.any_domain / 128aec1124d3 / 4

- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-6975d51072bab1120ea34c36e147af6db09b91779c1cb13eedb7c8bbae9de1db)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-fb514d22d5020b4e8322f5139d94a08eccd6003c39b871558da685b9c7827f85"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-51f8f546c56af1fbc98d4d8238aeb6fc7f2f686dae60ef5b127af105728fb228"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list.domain — client_side_defense.policy.js_insertion_rules.exclude_list.domain / 24b79edd5d33 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-013.md#canonical-367f873504986262d380bdfbd173acad4137578e1de683d131b3d0248b345e6f)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-187bc8e8a39a49b1be57c0106b3ea4eb53acefd0a6616f1a35708cf01ec7e22b)
- [client_side_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-cddcc8b626cac1ca512cd921feaf1053e00e0ff45310b85b8e9292b433ea905e)
- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-6975d51072bab1120ea34c36e147af6db09b91779c1cb13eedb7c8bbae9de1db)
- client_side_defense.policy.js_insertion_rules.exclude_list.domain

<a id="canonical-2435fc1648e4862738b3aef1b5feef0b2789c20a3f40e169185a587e6a225607"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Upstream description:

Domains names.

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

<a id="canonical-b0feee2595ddc197c57f8228eeb9cc1cfacf34364013408c62e8ad9e96ed3865"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.exclude_list.domain / 24b79edd5d33 / 3

<a id="canonical-715a83098e89ceeb80e9f7eb31da63b160ad844a492d8629ccc254c19d5b7bf7"></a>

<a id="canonical-47a87178ee1846ae66d94134c0479317b21cd1cf3a7dd652e14e065559e1dc7d"></a>

## exact_value property — client_side_defense.policy.js_insertion_rules.exclude_list.domain / 24b79edd5d33 / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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

<a id="canonical-f8d7f2201944ea4455888c66390be5fe010f05b4bd9d7d6ed5d33a0283c6e6b7"></a>

<a id="canonical-ea0c23f5a416fb93a5972b60fdb3909c1fe275789f9a37461d3fcc0afa4fa0ae"></a>

## regex_value property — client_side_defense.policy.js_insertion_rules.exclude_list.domain / 24b79edd5d33 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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

<a id="canonical-89f1678d25d9790d8b64e73f073af089efb66adb04c72238103f05c298de945d"></a>

<a id="canonical-d55ebdc22427c09e7a191a1ee6954bfb02ae8a7129ca685d2134140edd87784a"></a>

## suffix_value property — client_side_defense.policy.js_insertion_rules.exclude_list.domain / 24b79edd5d33 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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

<a id="canonical-a7323979147d9cc86f836d8999b40f6d7d82e850d221b5df70f470666c4ae4b6"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.exclude_list.domain / 24b79edd5d33 / 7

- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-6975d51072bab1120ea34c36e147af6db09b91779c1cb13eedb7c8bbae9de1db)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f4c48b51d2ba5dcc370d93d85d46e3e9fd068722c90aa1542b13bd30877a83b9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f62b5dbce62bb8f6c787904d6a3a6eec147c8e2973d99912e9354d0287d742ee"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list.metadata — client_side_defense.policy.js_insertion_rules.exclude_list.metadata / 0765b265a17b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-013.md#canonical-367f873504986262d380bdfbd173acad4137578e1de683d131b3d0248b345e6f)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-187bc8e8a39a49b1be57c0106b3ea4eb53acefd0a6616f1a35708cf01ec7e22b)
- [client_side_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-cddcc8b626cac1ca512cd921feaf1053e00e0ff45310b85b8e9292b433ea905e)
- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-6975d51072bab1120ea34c36e147af6db09b91779c1cb13eedb7c8bbae9de1db)
- client_side_defense.policy.js_insertion_rules.exclude_list.metadata

<a id="canonical-cc4e1667be954540b8636b3d28a68f7231d03421a25bb18c8c6ccca2c766f6d9"></a>

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

<a id="canonical-8dbcbdaeca0e73a533f67471d1c83f23f473e5b44f4616894e07ee0b35814dbc"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.exclude_list.metadata / 0765b265a17b / 3

<a id="canonical-04cec56d2967112fc2e9e2e98d89a8c38c7485b6a15c18e62bd67ffe315da706"></a>

<a id="canonical-13b0db9a78b42f6cd17c6ee416aa2f28abbc40e0511df0b23654cf3f859040e2"></a>

## description_spec property — client_side_defense.policy.js_insertion_rules.exclude_list.metadata / 0765b265a17b / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-aea950d43c379c897265a003a166574de66c020967aaab0b45e1177438893c18"></a>

<a id="canonical-3a990a8cf48e0e3869e74fe8a018a213af95f14b88a6ec365e0ce3d667435dc6"></a>

## name property — client_side_defense.policy.js_insertion_rules.exclude_list.metadata / 0765b265a17b / 5

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

<a id="canonical-5944805848a940a10dc00914d6b22cea38dd856b07f5e7a1d92665aabbd5262a"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.exclude_list.metadata / 0765b265a17b / 6

- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-6975d51072bab1120ea34c36e147af6db09b91779c1cb13eedb7c8bbae9de1db)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5298aa4e5e8e2e9cf2aaf88026cd4a4165a16593606644ac9f7a6d1bb75fca32"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-79f101848db9c5a08d8ef25b67bf1db37337f55fb466919843a4c62704a54b84"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list.path — client_side_defense.policy.js_insertion_rules.exclude_list.path / 75447f254fda / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-013.md#canonical-367f873504986262d380bdfbd173acad4137578e1de683d131b3d0248b345e6f)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-187bc8e8a39a49b1be57c0106b3ea4eb53acefd0a6616f1a35708cf01ec7e22b)
- [client_side_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-cddcc8b626cac1ca512cd921feaf1053e00e0ff45310b85b8e9292b433ea905e)
- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-6975d51072bab1120ea34c36e147af6db09b91779c1cb13eedb7c8bbae9de1db)
- client_side_defense.policy.js_insertion_rules.exclude_list.path

<a id="canonical-e080e6cfcd452b6cd0ff5c3df9a22fc3f198d6ee2873d435bfda2eca1fb79df4"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

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

<a id="canonical-6d8d807fd7a036d6a7dcfa33f92f7d32e4b96a425b95696823eaf88ffe897523"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.exclude_list.path / 75447f254fda / 3

<a id="canonical-7c245cf0d169045638f7f431580c7489832d649b49597b7585bbf5a9275a6da2"></a>

<a id="canonical-00bc6c424db4391144c956b50063125a1a4c4b2d3323683944d661401a6b404c"></a>

## path property — client_side_defense.policy.js_insertion_rules.exclude_list.path / 75447f254fda / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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

<a id="canonical-a754accd2462cded432433edae4d5f486a71fc961eb920b29800890ed55edbae"></a>

<a id="canonical-7d0c71e13ab32ac1d77bcb58ed55157d01144adbf9097e709c189a336b722656"></a>

## prefix property — client_side_defense.policy.js_insertion_rules.exclude_list.path / 75447f254fda / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-aab7853aa9e1c17946214e683e7fa0e1ecdc809f8ce8c223df6c1966f22a8685"></a>

<a id="canonical-7ac783dd27b52c9f2fde561d127f580b3728cbb626d01a2203788e7094fff832"></a>

## regex property — client_side_defense.policy.js_insertion_rules.exclude_list.path / 75447f254fda / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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

<a id="canonical-0979bcf08a4ef4f16d1c4f9882c8d2ce8b68569d1ea76b4c06c80ecebaa26405"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.exclude_list.path / 75447f254fda / 7

- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-6975d51072bab1120ea34c36e147af6db09b91779c1cb13eedb7c8bbae9de1db)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-02713d62645a8166d0c9fca71158dd7209f07d13c2af4dfd6d1cedaa5170fe71"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2337283b1051c2d081d4c10872a2ea469aaa7c63f3c956e265743a1bf4cf5f19"></a>

## client_side_defense.policy.js_insertion_rules.rules — client_side_defense.policy.js_insertion_rules.rules / e00a25231aef / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-013.md#canonical-367f873504986262d380bdfbd173acad4137578e1de683d131b3d0248b345e6f)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-187bc8e8a39a49b1be57c0106b3ea4eb53acefd0a6616f1a35708cf01ec7e22b)
- [client_side_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-cddcc8b626cac1ca512cd921feaf1053e00e0ff45310b85b8e9292b433ea905e)
- client_side_defense.policy.js_insertion_rules.rules

<a id="canonical-38a9ad19fea1fe5466babe6ed8dc55eaa072ec09abdd05e5107fe231d4c986bf"></a>

Type: `"list"`. Computed.

Required list of pages to insert Client-Side Defense client JavaScript.

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

<a id="canonical-db18167e92fee0b823432ed3d527c2c60e37d48a4e21b3b70ea928208728c125"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.rules / e00a25231aef / 3

- [any_domain](data-sources--http_loadbalancer--reference--group-014.md#canonical-e95f6f95ea088a879ada5bab4cec7deff96322e1b5d9edffc3b3f948f99eb692): complete subsection reference.

- [domain](data-sources--http_loadbalancer--reference--group-014.md#canonical-ffa2288495e0111462911ab0820e637f4d1259d767a07c12ea21ffcc2f6ade12): complete subsection reference.

- [metadata](data-sources--http_loadbalancer--reference--group-014.md#canonical-6f5a3dc406a8f6f0ea2b7e7bde44b43e2e247458f088c73e7cb1f7f1527e9f57): complete subsection reference.

- [path](data-sources--http_loadbalancer--reference--group-014.md#canonical-a3e270ce4c8afbfcfb9064c628a40d95a59036580ba15e7b82a448ac77e21736): complete subsection reference.

<a id="canonical-e59bd956d2b354ad7834d941ed0aa36a9c673e4c0efe1edc0c6e06084f443a25"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.rules / e00a25231aef / 4

- [client_side_defense.policy.js_insertion_rules.rules.any_domain](data-sources--http_loadbalancer--reference--group-014.md#canonical-e95f6f95ea088a879ada5bab4cec7deff96322e1b5d9edffc3b3f948f99eb692)
- [client_side_defense.policy.js_insertion_rules.rules.domain](data-sources--http_loadbalancer--reference--group-014.md#canonical-ffa2288495e0111462911ab0820e637f4d1259d767a07c12ea21ffcc2f6ade12)
- [client_side_defense.policy.js_insertion_rules.rules.metadata](data-sources--http_loadbalancer--reference--group-014.md#canonical-6f5a3dc406a8f6f0ea2b7e7bde44b43e2e247458f088c73e7cb1f7f1527e9f57)
- [client_side_defense.policy.js_insertion_rules.rules.path](data-sources--http_loadbalancer--reference--group-014.md#canonical-a3e270ce4c8afbfcfb9064c628a40d95a59036580ba15e7b82a448ac77e21736)
- [client_side_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-cddcc8b626cac1ca512cd921feaf1053e00e0ff45310b85b8e9292b433ea905e)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e95f6f95ea088a879ada5bab4cec7deff96322e1b5d9edffc3b3f948f99eb692"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-384979d435448c39bd1da7e96fde131fe520814a02eab911c0a96802c0f1b355"></a>

## client_side_defense.policy.js_insertion_rules.rules.any_domain — client_side_defense.policy.js_insertion_rules.rules.any_domain / 1f780e237417 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-013.md#canonical-367f873504986262d380bdfbd173acad4137578e1de683d131b3d0248b345e6f)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-187bc8e8a39a49b1be57c0106b3ea4eb53acefd0a6616f1a35708cf01ec7e22b)
- [client_side_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-cddcc8b626cac1ca512cd921feaf1053e00e0ff45310b85b8e9292b433ea905e)
- [client_side_defense.policy.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-02713d62645a8166d0c9fca71158dd7209f07d13c2af4dfd6d1cedaa5170fe71)
- client_side_defense.policy.js_insertion_rules.rules.any_domain

<a id="canonical-aca3296d62827129427ab4e070ca11129ca937b6f09e2bc667f1b8d4a0b5b272"></a>

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

<a id="canonical-49500e1049e62749f79321cfd7fcdd5fa0ec2e83e23c7bbfd143155cfcfc3a8d"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.rules.any_domain / 1f780e237417 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-be45f4792ef561b517bdce2de2b766cd9858919b7d74152283f9926be41584dc"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.rules.any_domain / 1f780e237417 / 4

- [client_side_defense.policy.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-02713d62645a8166d0c9fca71158dd7209f07d13c2af4dfd6d1cedaa5170fe71)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ffa2288495e0111462911ab0820e637f4d1259d767a07c12ea21ffcc2f6ade12"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2b3b88fbdfc6906a0710204e5f89aa01549cd6131bd5e3551de4a7a10187d952"></a>

## client_side_defense.policy.js_insertion_rules.rules.domain — client_side_defense.policy.js_insertion_rules.rules.domain / 202fd0382fb0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-013.md#canonical-367f873504986262d380bdfbd173acad4137578e1de683d131b3d0248b345e6f)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-187bc8e8a39a49b1be57c0106b3ea4eb53acefd0a6616f1a35708cf01ec7e22b)
- [client_side_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-cddcc8b626cac1ca512cd921feaf1053e00e0ff45310b85b8e9292b433ea905e)
- [client_side_defense.policy.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-02713d62645a8166d0c9fca71158dd7209f07d13c2af4dfd6d1cedaa5170fe71)
- client_side_defense.policy.js_insertion_rules.rules.domain

<a id="canonical-858ce9dc76032d97962d8c1613025b72109f1069ebbdfec8d63a960cca073e99"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Upstream description:

Domains names.

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

<a id="canonical-809664e21698089af6f92e0078506cdf999e0461ccb39e8a592f7793f964446d"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.rules.domain / 202fd0382fb0 / 3

<a id="canonical-3bd2872a13fbdeb50b95b2334a45173d66ba6caf09aba0db126cf735c06a8014"></a>

<a id="canonical-1e292adfbe9a1fa4cf620503b77df1759ddb1a9dcd2c2dd62984b1aec6182c49"></a>

## exact_value property — client_side_defense.policy.js_insertion_rules.rules.domain / 202fd0382fb0 / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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

<a id="canonical-9d4e1195424d01689aad7a1e509552f35a3c1a5661b017a4b3f93fd0e52ea132"></a>

<a id="canonical-e5315e06ac9714d483ba32598a72a3d7d34e4f64d3952a6e8e2fcb85fbdc0a73"></a>

## regex_value property — client_side_defense.policy.js_insertion_rules.rules.domain / 202fd0382fb0 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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

<a id="canonical-0be9a8e9478d7a14b1e50482a93fa0d4bf003578d75f153c1b6246c641f6bce4"></a>

<a id="canonical-a6721a861e5bf30bd673caf7d20908ca576af2720336c070226c9b91b6c812a5"></a>

## suffix_value property — client_side_defense.policy.js_insertion_rules.rules.domain / 202fd0382fb0 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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

<a id="canonical-166aab92c0c58e303cf8c06e0a3d10f439946c736817364bbe9c927e23e9576a"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.rules.domain / 202fd0382fb0 / 7

- [client_side_defense.policy.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-02713d62645a8166d0c9fca71158dd7209f07d13c2af4dfd6d1cedaa5170fe71)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6f5a3dc406a8f6f0ea2b7e7bde44b43e2e247458f088c73e7cb1f7f1527e9f57"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b34c64e4a7dfe2f4321dac8eb236e5e2280bdb7976c0b79a76ddcd38c3e83b58"></a>

## client_side_defense.policy.js_insertion_rules.rules.metadata — client_side_defense.policy.js_insertion_rules.rules.metadata / 1613880acb2b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-013.md#canonical-367f873504986262d380bdfbd173acad4137578e1de683d131b3d0248b345e6f)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-187bc8e8a39a49b1be57c0106b3ea4eb53acefd0a6616f1a35708cf01ec7e22b)
- [client_side_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-cddcc8b626cac1ca512cd921feaf1053e00e0ff45310b85b8e9292b433ea905e)
- [client_side_defense.policy.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-02713d62645a8166d0c9fca71158dd7209f07d13c2af4dfd6d1cedaa5170fe71)
- client_side_defense.policy.js_insertion_rules.rules.metadata

<a id="canonical-2f53039c7d52e0fdf9e90864e3debf76cfe98fabfc6c45a307c5c2b25477360b"></a>

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

<a id="canonical-ec561062f223dabc42f91d3f4ff791926edae06c72680f5243aec56e8246110b"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.rules.metadata / 1613880acb2b / 3

<a id="canonical-16cedd24abf3ed298a7558e69f6c1871266f67f3eaf42acfcf8453de74472cd6"></a>

<a id="canonical-a233e3a1076fe5f9a9cbe2306463473351ad9cfed99aebc675d6c8a759fbdcc6"></a>

## description_spec property — client_side_defense.policy.js_insertion_rules.rules.metadata / 1613880acb2b / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-ed5dde95ad5d2f91f919160ecc1891315cbcd7c00d63f7c07c11fd1f4864171b"></a>

<a id="canonical-4b2398adb5194c9cde4caf659653a68f69c59f7f5ca41369bc9ed4d4f7178b28"></a>

## name property — client_side_defense.policy.js_insertion_rules.rules.metadata / 1613880acb2b / 5

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

<a id="canonical-2c73fc1bbe82f103db7c34f0ccf22f6a6b928b6078fb9dccde73a366bb97f33d"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.rules.metadata / 1613880acb2b / 6

- [client_side_defense.policy.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-02713d62645a8166d0c9fca71158dd7209f07d13c2af4dfd6d1cedaa5170fe71)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a3e270ce4c8afbfcfb9064c628a40d95a59036580ba15e7b82a448ac77e21736"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-81502b2b138900b922635f0eda7d0c49d11060a21df3539949e5031eea8e7277"></a>

## client_side_defense.policy.js_insertion_rules.rules.path — client_side_defense.policy.js_insertion_rules.rules.path / f5cdf6e97c28 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-013.md#canonical-367f873504986262d380bdfbd173acad4137578e1de683d131b3d0248b345e6f)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-187bc8e8a39a49b1be57c0106b3ea4eb53acefd0a6616f1a35708cf01ec7e22b)
- [client_side_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-cddcc8b626cac1ca512cd921feaf1053e00e0ff45310b85b8e9292b433ea905e)
- [client_side_defense.policy.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-02713d62645a8166d0c9fca71158dd7209f07d13c2af4dfd6d1cedaa5170fe71)
- client_side_defense.policy.js_insertion_rules.rules.path

<a id="canonical-993c6a2081b7105ed89c6248d4a646ab15ec587488018e2b57de7192d177789c"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

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

<a id="canonical-ea77302b16fc95d2652d36bbd185d34d04099d2cf691aacf13a9e45a26d569fd"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.rules.path / f5cdf6e97c28 / 3

<a id="canonical-c9e50317b15bf295b65f7b63b32f6e8e15e176145a0bb762ee1c461051176676"></a>

<a id="canonical-4b853a2ca84407e36ed219fbed1d8bd838f2e842bcf1bb24f89fbfba1956ef7f"></a>

## path property — client_side_defense.policy.js_insertion_rules.rules.path / f5cdf6e97c28 / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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

<a id="canonical-6494092da2c89fed0932f3777c27d97ef831ab3ce2f162cde2efe8fdda18c122"></a>

<a id="canonical-14dd9b484a13aba6d48fa474e26032ed0799f6fc6ef835e0160a2a524435a266"></a>

## prefix property — client_side_defense.policy.js_insertion_rules.rules.path / f5cdf6e97c28 / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-1b8172c40408b7495a5ef9c52c59672d3980d4af4c4985e19657b90b5b5b0142"></a>

<a id="canonical-c3369ff12cb7085968173222b46e3812c5c3add7e707247fd1663fff3e27975e"></a>

## regex property — client_side_defense.policy.js_insertion_rules.rules.path / f5cdf6e97c28 / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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

<a id="canonical-145b38103505e1c946fe06e0942198a50dac2c6430428167311f61baeeddf8ac"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.rules.path / f5cdf6e97c28 / 7

- [client_side_defense.policy.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-02713d62645a8166d0c9fca71158dd7209f07d13c2af4dfd6d1cedaa5170fe71)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-acac9167259cfc120ba2f9c96337e59fc39aa9b471956e6aec6e6cd51b54cebb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-66f83ab5a13ab666aad62c5e96d2a572ffd7bf344ad1e3f05bee99426f957c00"></a>

## cookie_stickiness — cookie_stickiness / 3387fc63a050 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- cookie_stickiness

<a id="canonical-935e37365221f6bec8ac5bef261352aafe3d5bba8d7a04a5d0f81b0b22cf2df9"></a>

Type: `"single"`. Computed.

\[OneOf: cookie\_stickiness, least\_active, random, ring\_hash, round\_robin,
source\_ip\_stickiness; Default: round\_robin\] Two types of cookie affinity: 1. Passive. Takes a
cookie that's present in the cookies header and hashes on its value. 2. Generated. Generates and
sets a cookie with an expiration (TTL) on the first request from the client in its response to the
client, based on the endpoint the request gets..

Upstream description:

Two types of cookie affinity:

&#8203;1. Passive. Takes a cookie that's present in the cookies header and hashes on its value.

&#8203;2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from
the client in its response to the client, based on the endpoint the request gets sent to. The client
then presents this on the next and all subsequent requests. The hash of this is sufficient to ensure
these requests GET sent to the same endpoint. The cookie is generated by hashing the source and
destination ports and addresses so that multiple independent HTTP2 streams on the same connection
will independently receive the same cookie, even if they arrive simultaneously.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-httponly": "[\"add_httponly\",\"ignore_httponly\"]",
  "x-ves-oneof-field-samesite": "[\"ignore_samesite\",\"samesite_lax\",\"samesite_none\",\"samesite_strict\"]",
  "x-ves-oneof-field-secure": "[\"add_secure\",\"ignore_secure\"]"
}
```

OneOf alternatives in this subsection:

- [cookie_stickiness](data-sources--http_loadbalancer--reference--group-014.md#canonical-935e37365221f6bec8ac5bef261352aafe3d5bba8d7a04a5d0f81b0b22cf2df9)
- [least_active](data-sources--http_loadbalancer--reference--group-019.md#canonical-38ec9dcde696b28de7b1d7aef42896ffc9b35a0c76615bd6f1b9f772b396b0d3)
- [random](data-sources--http_loadbalancer--reference--group-022.md#canonical-9b3e8bbea8a3216200bd8ea4bb178a27ed62696ccb56345712cde3bec3253770)
- [ring_hash](data-sources--http_loadbalancer--reference--group-022.md#canonical-3be9344020d3d328f50af074648df746217617e5883ae4fc2786ed82f1e690aa)
- [round_robin](data-sources--http_loadbalancer--reference--group-022.md#canonical-17acf581cfd27d1efd14d30e5aa809447140203492f494ebda66ef71a5d0801f)
- [source_ip_stickiness](data-sources--http_loadbalancer--reference--group-025.md#canonical-aec3f86d7a9328509fd797ba929ddf4c65b50b051dc9cfc35600921fd2f6b5e0)

Select alternatives according to the provider validators above.

<a id="canonical-ca92aac34c5b6e5a999e0c4eb84f9f213fe3d081044591457d2c43125c0839d7"></a>

## Direct properties — cookie_stickiness / 3387fc63a050 / 3

- [add_httponly](data-sources--http_loadbalancer--reference--group-014.md#canonical-55453ddfd846622dab2d6fb83517eb02d809d1839f0ebd8ada16af1861d757ab): complete subsection reference.

- [add_secure](data-sources--http_loadbalancer--reference--group-014.md#canonical-56f958cf49190eedc35d623944211551fe0c4528225ac857063ef323364bb0aa): complete subsection reference.

- [ignore_httponly](data-sources--http_loadbalancer--reference--group-014.md#canonical-21a2b7c9c6e4ef40f251a1d4fe0ff27eb33e70a0b581464ab96dc09322097ec0): complete subsection reference.

- [ignore_samesite](data-sources--http_loadbalancer--reference--group-014.md#canonical-a3cda3173ef2c3585cfbd3a81dae993620e91dcba6dc851ff9f5b18dba1b36dd): complete subsection reference.

- [ignore_secure](data-sources--http_loadbalancer--reference--group-014.md#canonical-9f5971cb72bf3bf31d4c2f72e646818c59b2e77d336196a39f812ce583dc779d): complete subsection reference.

<a id="canonical-50ed4d01e461666da907be349aa3f55781ef620f69c2e19242bfc6baa105ae7f"></a>

<a id="canonical-54c722d57bc7978f42b31b378fc28d9955f2c47bcb7b9540820ba7a1227affe5"></a>

## name property — cookie_stickiness / 3387fc63a050 / 4

Type: `"string"`. Computed.

The name of the cookie that will be used to obtain the hash key. If the cookie is not present and
TTL below is not set, no hash will be produced.

Upstream description:

The name of the cookie that will be used to obtain the hash key. If the cookie is not present and
TTL below is not set, no hash will be produced.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-b7e1e3af61c916cd2389fe9b2feee48d7ad444fc8ecd4f32e62115511a5be9fb"></a>

<a id="canonical-746751661599f97bb3e95a67bd6edcf4ff831b48bf28d9b0179dbe1328f65555"></a>

## path property — cookie_stickiness / 3387fc63a050 / 5

Type: `"string"`. Computed.

The name of the path for the cookie. If no path is specified here, no path will be set for the
cookie.

Upstream description:

The name of the path for the cookie. If no path is specified here, no path will be set for the
cookie.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
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
  }
}
```

- [samesite_lax](data-sources--http_loadbalancer--reference--group-014.md#canonical-2b9a7c3e83f112d01c97d2857f99b58de01c57c7d4d7dfba31e5c68d9f61dd8f): complete subsection reference.

- [samesite_none](data-sources--http_loadbalancer--reference--group-014.md#canonical-0c0b7ea3943a2ab0683b0b06b42c02b85d45a12779cf608d1f20d894cf51e560): complete subsection reference.

- [samesite_strict](data-sources--http_loadbalancer--reference--group-014.md#canonical-ac9d33ac221c090bbf52d1612a9ba71dc893933cb199890cc50b423dfaeb9ab6): complete subsection reference.

<a id="canonical-77a67d9aad8a572961eba6a21cbbf0323b6dbecb67f7a084d2f49c5b3ba5843f"></a>

<a id="canonical-ad7be0f0b5e7b8be9f78e3e70ff9d577385407beaa385fe291f3bd13cd03f615"></a>

## ttl property — cookie_stickiness / 3387fc63a050 / 6

Type: `"number"`. Computed.

If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is
present and zero, the generated cookie will be a session cookie. TTL value is in milliseconds.

Upstream description:

If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is
present and zero, the generated cookie will be a session cookie. TTL value is in milliseconds.

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

<a id="canonical-28d26b6a3db42416497ee40ebd777ecfc794b7ef20e6d4af78ff19c89835202b"></a>

## Next pages — cookie_stickiness / 3387fc63a050 / 7

- [cookie_stickiness.add_httponly](data-sources--http_loadbalancer--reference--group-014.md#canonical-55453ddfd846622dab2d6fb83517eb02d809d1839f0ebd8ada16af1861d757ab)
- [cookie_stickiness.add_secure](data-sources--http_loadbalancer--reference--group-014.md#canonical-56f958cf49190eedc35d623944211551fe0c4528225ac857063ef323364bb0aa)
- [cookie_stickiness.ignore_httponly](data-sources--http_loadbalancer--reference--group-014.md#canonical-21a2b7c9c6e4ef40f251a1d4fe0ff27eb33e70a0b581464ab96dc09322097ec0)
- [cookie_stickiness.ignore_samesite](data-sources--http_loadbalancer--reference--group-014.md#canonical-a3cda3173ef2c3585cfbd3a81dae993620e91dcba6dc851ff9f5b18dba1b36dd)
- [cookie_stickiness.ignore_secure](data-sources--http_loadbalancer--reference--group-014.md#canonical-9f5971cb72bf3bf31d4c2f72e646818c59b2e77d336196a39f812ce583dc779d)
- [cookie_stickiness.samesite_lax](data-sources--http_loadbalancer--reference--group-014.md#canonical-2b9a7c3e83f112d01c97d2857f99b58de01c57c7d4d7dfba31e5c68d9f61dd8f)
- [cookie_stickiness.samesite_none](data-sources--http_loadbalancer--reference--group-014.md#canonical-0c0b7ea3943a2ab0683b0b06b42c02b85d45a12779cf608d1f20d894cf51e560)
- [cookie_stickiness.samesite_strict](data-sources--http_loadbalancer--reference--group-014.md#canonical-ac9d33ac221c090bbf52d1612a9ba71dc893933cb199890cc50b423dfaeb9ab6)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-55453ddfd846622dab2d6fb83517eb02d809d1839f0ebd8ada16af1861d757ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f2686afa722fd2b2a44e7e56f592d90942c2a8d3afbd7507ad8850655d1f83c"></a>

## cookie_stickiness.add_httponly — cookie_stickiness.add_httponly / 39a91528e471 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [cookie_stickiness](data-sources--http_loadbalancer--reference--group-014.md#canonical-acac9167259cfc120ba2f9c96337e59fc39aa9b471956e6aec6e6cd51b54cebb)
- cookie_stickiness.add_httponly

<a id="canonical-389adce98f50f0f3ad7e75887c967e0a4d6c2d16a79e11158bfbbd37cc59989b"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for add httponly.

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

<a id="canonical-b515c7fee5a187906b36a107d4c7b57a298a50c3fc3761f29208ed01473e8706"></a>

## Direct properties — cookie_stickiness.add_httponly / 39a91528e471 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4792a7f84f307300acc76fdf5103862828fc27526319f771d278e7b3dbb24f4b"></a>

## Next pages — cookie_stickiness.add_httponly / 39a91528e471 / 4

- [cookie_stickiness](data-sources--http_loadbalancer--reference--group-014.md#canonical-acac9167259cfc120ba2f9c96337e59fc39aa9b471956e6aec6e6cd51b54cebb)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-56f958cf49190eedc35d623944211551fe0c4528225ac857063ef323364bb0aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a174fe7748fb610348c0e1e5fea7e977c2fa79b52573bb45765dc1cf805eade"></a>

## cookie_stickiness.add_secure — cookie_stickiness.add_secure / cc1da324c522 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [cookie_stickiness](data-sources--http_loadbalancer--reference--group-014.md#canonical-acac9167259cfc120ba2f9c96337e59fc39aa9b471956e6aec6e6cd51b54cebb)
- cookie_stickiness.add_secure

<a id="canonical-24e09f631add61cab6eb8ff9431a147e064870f2cce8fdf85431bcf49effd133"></a>

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

<a id="canonical-17f2aa6d6391214c58d9b4f84017044d00d4184422ab401d57ac797d9c716db6"></a>

## Direct properties — cookie_stickiness.add_secure / cc1da324c522 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9074d22778ac4df069c286277bf4266e8a399defbc220fa8ddb98e54ea49dc5f"></a>

## Next pages — cookie_stickiness.add_secure / cc1da324c522 / 4

- [cookie_stickiness](data-sources--http_loadbalancer--reference--group-014.md#canonical-acac9167259cfc120ba2f9c96337e59fc39aa9b471956e6aec6e6cd51b54cebb)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-21a2b7c9c6e4ef40f251a1d4fe0ff27eb33e70a0b581464ab96dc09322097ec0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-77e6498baf9ba1c2fe686f3ba759438125f30cc54f8fb57d1675534f8897de1e"></a>

## cookie_stickiness.ignore_httponly — cookie_stickiness.ignore_httponly / 27bcec557a47 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [cookie_stickiness](data-sources--http_loadbalancer--reference--group-014.md#canonical-acac9167259cfc120ba2f9c96337e59fc39aa9b471956e6aec6e6cd51b54cebb)
- cookie_stickiness.ignore_httponly

<a id="canonical-97d6c5e92aa807b9f89ae6ef1128e7f27387aa1501283a334a42f7e840728d8c"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore httponly.

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

<a id="canonical-24820d39ff164ffa719471db31fe653d94a5b133f490cea205f9ace8576deba7"></a>

## Direct properties — cookie_stickiness.ignore_httponly / 27bcec557a47 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fefb6afa8cb1634e87c39e50d06999439331a0e8395f0bd671b2d138360ffee1"></a>

## Next pages — cookie_stickiness.ignore_httponly / 27bcec557a47 / 4

- [cookie_stickiness](data-sources--http_loadbalancer--reference--group-014.md#canonical-acac9167259cfc120ba2f9c96337e59fc39aa9b471956e6aec6e6cd51b54cebb)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a3cda3173ef2c3585cfbd3a81dae993620e91dcba6dc851ff9f5b18dba1b36dd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a8849025363c87a525e2745239fd742c9abf4c6e1e2b9a974fd854a5cf57bde9"></a>

## cookie_stickiness.ignore_samesite — cookie_stickiness.ignore_samesite / b1e8f120a5da / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [cookie_stickiness](data-sources--http_loadbalancer--reference--group-014.md#canonical-acac9167259cfc120ba2f9c96337e59fc39aa9b471956e6aec6e6cd51b54cebb)
- cookie_stickiness.ignore_samesite

<a id="canonical-ac6b98615a73fc394c4ac39c51bd6749d4a0a85523c3815f8a69012c45d02018"></a>

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

<a id="canonical-2bcfe40b174ce541ef4d8fe48251bb95c621c56937d1060aaa277352a7182c2b"></a>

## Direct properties — cookie_stickiness.ignore_samesite / b1e8f120a5da / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f838ff2755b9081f421e43cbfa7d7ad20d6437fefb4d1cd67ef3b13f63d55d49"></a>

## Next pages — cookie_stickiness.ignore_samesite / b1e8f120a5da / 4

- [cookie_stickiness](data-sources--http_loadbalancer--reference--group-014.md#canonical-acac9167259cfc120ba2f9c96337e59fc39aa9b471956e6aec6e6cd51b54cebb)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9f5971cb72bf3bf31d4c2f72e646818c59b2e77d336196a39f812ce583dc779d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-986e1d9e12f9bda2e94400cde9f799a736a0521f52924c13d476190e23eb2cf9"></a>

## cookie_stickiness.ignore_secure — cookie_stickiness.ignore_secure / efb04b8d0d78 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [cookie_stickiness](data-sources--http_loadbalancer--reference--group-014.md#canonical-acac9167259cfc120ba2f9c96337e59fc39aa9b471956e6aec6e6cd51b54cebb)
- cookie_stickiness.ignore_secure

<a id="canonical-0eb0f3ce02e88124226a0620459833a2cbaa44794abf54de7b221b4825bbef6f"></a>

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

<a id="canonical-92054c3f9aeb778cf66b51c2f1e04415cb3ba1489638d80565e17dea154fc5e9"></a>

## Direct properties — cookie_stickiness.ignore_secure / efb04b8d0d78 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-12c4e5c6469c34f1ccb8bc6adc67ca0000ff503f4bc59b89cfc505a6ba0d34c2"></a>

## Next pages — cookie_stickiness.ignore_secure / efb04b8d0d78 / 4

- [cookie_stickiness](data-sources--http_loadbalancer--reference--group-014.md#canonical-acac9167259cfc120ba2f9c96337e59fc39aa9b471956e6aec6e6cd51b54cebb)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-2b9a7c3e83f112d01c97d2857f99b58de01c57c7d4d7dfba31e5c68d9f61dd8f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-10924ed95cef30a41ef26af29dc4c74e9ff8c1d79f9f484f5a0b6635dbe6f46a"></a>

## cookie_stickiness.samesite_lax — cookie_stickiness.samesite_lax / 8d7b4fb387aa / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [cookie_stickiness](data-sources--http_loadbalancer--reference--group-014.md#canonical-acac9167259cfc120ba2f9c96337e59fc39aa9b471956e6aec6e6cd51b54cebb)
- cookie_stickiness.samesite_lax

<a id="canonical-edc9eb57f774ec2c174ba74cc15f085d5298353e9500bc18aa44980a0b358a63"></a>

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

<a id="canonical-c5faa875b986f064470f0e98cc6a2391adaa0f0794d992059dfa1ed541ed8b66"></a>

## Direct properties — cookie_stickiness.samesite_lax / 8d7b4fb387aa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d863abc37735fa7d57b2c889b1db14873b18c3d08e3ab5d22aaff6f38cc9fbff"></a>

## Next pages — cookie_stickiness.samesite_lax / 8d7b4fb387aa / 4

- [cookie_stickiness](data-sources--http_loadbalancer--reference--group-014.md#canonical-acac9167259cfc120ba2f9c96337e59fc39aa9b471956e6aec6e6cd51b54cebb)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-0c0b7ea3943a2ab0683b0b06b42c02b85d45a12779cf608d1f20d894cf51e560"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6456ee073094399b4a9aaf5848d7821e842c19bb27926ee4e349fe9c396e3f69"></a>

## cookie_stickiness.samesite_none — cookie_stickiness.samesite_none / 74ba11531095 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [cookie_stickiness](data-sources--http_loadbalancer--reference--group-014.md#canonical-acac9167259cfc120ba2f9c96337e59fc39aa9b471956e6aec6e6cd51b54cebb)
- cookie_stickiness.samesite_none

<a id="canonical-3f7082f9bea3085b330ac9b43ff150b9eea535ac6579af3d4144b97490f12f12"></a>

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

<a id="canonical-b6288c555bd6ce2d772cd57fed40190365468f3fa93e12ec79baa8e7fdd032ed"></a>

## Direct properties — cookie_stickiness.samesite_none / 74ba11531095 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3d15d68fe151cfc3f4f736854f6bd4b7f47a2b41cb554d4fc6a01d2dc8413fca"></a>

## Next pages — cookie_stickiness.samesite_none / 74ba11531095 / 4

- [cookie_stickiness](data-sources--http_loadbalancer--reference--group-014.md#canonical-acac9167259cfc120ba2f9c96337e59fc39aa9b471956e6aec6e6cd51b54cebb)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ac9d33ac221c090bbf52d1612a9ba71dc893933cb199890cc50b423dfaeb9ab6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d75ea9ac8177fe0b12a17e398bba0144d7e164e81b4bb67f167f75b9fd52a362"></a>

## cookie_stickiness.samesite_strict — cookie_stickiness.samesite_strict / 74caed25cfda / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [cookie_stickiness](data-sources--http_loadbalancer--reference--group-014.md#canonical-acac9167259cfc120ba2f9c96337e59fc39aa9b471956e6aec6e6cd51b54cebb)
- cookie_stickiness.samesite_strict

<a id="canonical-614408f54990c54525bc2716ac92ab09bd4fbe82f4c5b2c63686305b33269979"></a>

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

<a id="canonical-d8303b106b20a070b5c6586d5e002e3351b54cb1c9baead44ef00e8920b62630"></a>

## Direct properties — cookie_stickiness.samesite_strict / 74caed25cfda / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-24e21e2bfc34becac26f3a5fc7fa1eecfaa0ee867c4c387843b3d4f7281150eb"></a>

## Next pages — cookie_stickiness.samesite_strict / 74caed25cfda / 4

- [cookie_stickiness](data-sources--http_loadbalancer--reference--group-014.md#canonical-acac9167259cfc120ba2f9c96337e59fc39aa9b471956e6aec6e6cd51b54cebb)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-4ffd9159bc5a97623e1ac711bca5b5865d2f43de09a0b5b7552e4670e7a446be"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a54ac51607ce9df1025b419c2a44ce662b6f4ee96db9a587de1605fdd03ee83b"></a>

## cors_policy — cors_policy / 061a1c85e987 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- cors_policy

<a id="canonical-49b10f33465a7f2ec72ab8f86abd2648a54ec7cc1f9c26cf9a4855350cb5f772"></a>

Type: `"single"`. Computed.

Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route
level configuration takes precedence. An example of an Cross origin HTTP request GET
/resources/public-data/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel MAC OS
X 10.5..

Upstream description:

Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route
level configuration takes precedence.

An example of an Cross origin HTTP request GET /resources/public-data/ HTTP/1.1 Host: bar.other
User-Agent: Mozilla/5.0 (Macintosh; U; Intel MAC OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130
Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,\*/\*;q=0.8
Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset:
ISO-8859-1,utf-8;q=0.7,\*;q=0.7 Connection: keep-alive Referrer:
http&#58;//foo.example/examples/access-control/simplexsinvocation.html Origin:
http&#58;//foo.example

HTTP/1.1 200 OK Date: Mon, 01 Dec 2008 00:23:53 GMT Server: Apache/2.0.61
Access-Control-Allow-Origin: \* Keep-Alive: timeout=2, max=100 Connection: Keep-Alive
Transfer-Encoding: chunked Content-Type: application/XML

An example for cross origin HTTP OPTIONS request with Access-Control-Request-\* header

OPTIONS /resources/POST-here/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel
MAC OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept:
text/HTML,application/xhtml+XML,application/XML;q=0.9,\*/\*;q=0.8 Accept-Language: en-us,en;q=0.5
Accept-Encoding: gzip,deflate Accept-Charset: ISO-8859-1,utf-8;q=0.7,\*;q=0.7 Connection: keep-alive
Origin: http&#58;//foo.example Access-Control-Request-Method: POST Access-Control-Request-Headers:
X-PINGOTHER, Content-Type

HTTP/1.1 204 No Content Date: Mon, 01 Dec 2008 01:15:39 GMT Server: Apache/2.0.61 (Unix)
Access-Control-Allow-Origin: http&#58;//foo.example Access-Control-Allow-Methods: POST, GET, OPTIONS
Access-Control-Allow-Headers: X-PINGOTHER, Content-Type Access-Control-Max-Age: 86400 Vary:
Accept-Encoding, Origin Keep-Alive: timeout=2, max=100 Connection: Keep-Alive.

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

<a id="canonical-a6c68fd3fc975667afbde45620c72daeeaa111113939fe1389ba1900b8dfac09"></a>

## Direct properties — cors_policy / 061a1c85e987 / 3

<a id="canonical-194e05f4c95becc5ebcd3610372c1627171b9b5447f3e520fba50741ab8132ae"></a>

<a id="canonical-ca9ce716979ae52565ea9e3f374a4b65b20c8908c513a9b9ed705fdc9179ea23"></a>

## allow_credentials property — cors_policy / 061a1c85e987 / 4

Type: `"bool"`. Computed.

Specifies whether the resource allows credentials.

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

<a id="canonical-8c0762c47e35869a5a7f59130d95d5824a78c13e2b2f8d777cc070ab69eb1410"></a>

<a id="canonical-8728406c54e7804cf539bb106b9515a1ac479d7baa36334f2360d9cd5d7a50ab"></a>

## allow_headers property — cors_policy / 061a1c85e987 / 5

Type: `"string"`. Computed.

Specifies the content for the access-control-allow-headers header.

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

<a id="canonical-34aef95bc8ff150dd121e8a14bc3971e8d50289817b788e22d1b7727497bbb72"></a>

<a id="canonical-1e6fb57be8ec36737671d2fc5d75386fa60cec7d49bee0af22843632b0638a8e"></a>

## allow_methods property — cors_policy / 061a1c85e987 / 6

Type: `"string"`. Computed.

Specifies the content for the access-control-allow-methods header.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_valid_methods": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_valid_methods": "true"
  }
}
```

<a id="canonical-64e1e6d82577b2d6cbef84ef913a40f8d919c47cf413c77ac516ccf08fdc5d6e"></a>

<a id="canonical-b7013760e0974a8689ce6b08bb9a7929534198bc2536887d47c6d2260354dade"></a>

## allow_origin property — cors_policy / 061a1c85e987 / 7

Type: `["list", "string"]`. Computed.

Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either
allow\_origin or allow\_origin\_regex match.

Upstream description:

Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either
allow\_origin or allow\_origin\_regex match.

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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2dcef50b7ef733b0b4714f2b191f00834217b0685019a1c938e8ea5cc8cc3fb8"></a>

<a id="canonical-e16e9ee96d38e6276cd55b33dffa34355603d6b65e29c595d5a3b0b41d822400"></a>

## allow_origin_regex property — cors_policy / 061a1c85e987 / 8

Type: `["list", "string"]`. Computed.

Specifies regex patterns that match allowed origins. An origin is allowed if either allow\_origin or
allow\_origin\_regex match.

Upstream description:

Specifies regex patterns that match allowed origins. An origin is allowed if either allow\_origin or
allow\_origin\_regex match.

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
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-8d5ce08b6c908b34e4d32df30b5e7f5b2509d4c7a3d3c2695be47c83960ad0b2"></a>

<a id="canonical-f56aacf23086494fcf2327ff4f9a476d59b47e36bca030fc9d839a8249c57134"></a>

## disabled property — cors_policy / 061a1c85e987 / 9

Type: `"bool"`. Computed.

Disable the CorsPolicy for a particular route. This is useful when virtual-host has CorsPolicy, but
we need to disable it on a specific route. The value of this field is ignored for virtual-host.

Upstream description:

Disable the CorsPolicy for a particular route. This is useful when virtual-host has CorsPolicy, but
we need to disable it on a specific route. The value of this field is ignored for virtual-host.

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

<a id="canonical-dcd9b4b4f293715e395bd9791f6918de86714e870f23b5a964b3b20ab98049af"></a>

<a id="canonical-cbe1f90176d380ac436a9bbdf2f1a52bfba3b0c61d5b304f391c90381bfdf599"></a>

## expose_headers property — cors_policy / 061a1c85e987 / 10

Type: `"string"`. Computed.

Specifies the content for the access-control-expose-headers header.

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

<a id="canonical-fcc6aa99c0d3d961f8c571f00359c1ec3c452d1775091665812da78dcd939aec"></a>

<a id="canonical-9d86b938b02b476006c76c7bf9093b407fbb7f6135ae2f8f756a7933e4f66273"></a>

## maximum_age property — cors_policy / 061a1c85e987 / 11

Type: `"number"`. Computed.

Specifies the content for the access-control-max-age header in seconds. This indicates the maximum
number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted
value is 86400 seconds (24 hours).

Upstream description:

Specifies the content for the access-control-max-age header in seconds. This indicates the maximum
number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted
value is 86400 seconds (24 hours)

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
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": -1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "-1",
    "ves.io.schema.rules.int32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "-1",
    "ves.io.schema.rules.int32.lte": "86400"
  }
}
```

<a id="canonical-b395b77a61a8355a94d9a0179b1fd4b01f67d6de3428da5ce48c52ffa8660862"></a>

## Next pages — cors_policy / 061a1c85e987 / 12

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f1e6471143974860830fd021444acd48e74e2b1b6f98ec945d6dd4b24d617417"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9fc325dfe0da83206f5bc666a91431fa6708fa253a8ea97427929551226bb432"></a>

## csrf_policy — csrf_policy / c0b3165bb9b7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- csrf_policy

<a id="canonical-8a0c0aa5262380b506d1ef31f3f2ed8b8d18c62b88a16120288e504af647a1a1"></a>

Type: `"single"`. Computed.

To mitigate CSRF attack , the policy checks where a request is coming from to determine if the
request's origin is the same as its destination.the policy relies on two pieces of information used
in determining if a request originated from the same host. 1. The origin that caused the user
agent..

Upstream description:

To mitigate CSRF attack , the policy checks where a request is coming from to determine if the
request's origin is the same as its destination.the policy relies on two pieces of information used
in determining if a request originated from the same host.

&#8203;1. The origin that caused the user agent to issue the request (source origin). &#8203;2. The
origin that the request is going to (target origin). When the policy evaluating a request, it
ensures both pieces of information are present and compare their values. If the source origin is
missing or origins do not match the request is rejected. The exception to this being if the
source-origin has been added to they policy as valid. Because CSRF attacks specifically target
state-changing requests, the policy only acts on the HTTP requests that have state-changing method
(PUT,POST, etc.).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-allowed_domains": "[\"all_load_balancer_domains\",\"custom_domain_list\",\"disabled\"]"
}
```

<a id="canonical-5171e44734a74e071cb2abe8851ce326f6c0425298c20bf3c85f4bbbfe65a166"></a>

## Direct properties — csrf_policy / c0b3165bb9b7 / 3

- [all_load_balancer_domains](data-sources--http_loadbalancer--reference--group-014.md#canonical-fa30226d15e52eae51ea64a28ad7672f0ee7424cd7c245f16fbc1d7e1152a714): complete subsection reference.

- [custom_domain_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-2e1041aaa99491d4c3772a836d7e3c55ff98583b7c54293dc8e78b61d19ab1ac): complete subsection reference.

- [disabled](data-sources--http_loadbalancer--reference--group-014.md#canonical-7614e339b1bdbe39d979bfc35dfe573113a2407c643d5fad0118bcc8fe93fb20): complete subsection reference.

<a id="canonical-09a84b17ccb438293f0cdd4afb0730c52b01c6b5fe3fe468c0637cf774a05e7c"></a>

## Next pages — csrf_policy / c0b3165bb9b7 / 4

- [csrf_policy.all_load_balancer_domains](data-sources--http_loadbalancer--reference--group-014.md#canonical-fa30226d15e52eae51ea64a28ad7672f0ee7424cd7c245f16fbc1d7e1152a714)
- [csrf_policy.custom_domain_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-2e1041aaa99491d4c3772a836d7e3c55ff98583b7c54293dc8e78b61d19ab1ac)
- [csrf_policy.disabled](data-sources--http_loadbalancer--reference--group-014.md#canonical-7614e339b1bdbe39d979bfc35dfe573113a2407c643d5fad0118bcc8fe93fb20)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-fa30226d15e52eae51ea64a28ad7672f0ee7424cd7c245f16fbc1d7e1152a714"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c59d2bbeb62490530d60f798c7fbfe29c0deecabad5998ecfe9c0f7671a9a38"></a>

## csrf_policy.all_load_balancer_domains — csrf_policy.all_load_balancer_domains / 58ae450672e4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [csrf_policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-f1e6471143974860830fd021444acd48e74e2b1b6f98ec945d6dd4b24d617417)
- csrf_policy.all_load_balancer_domains

<a id="canonical-efbe5034b3ada496db7ca721e39606653bc76fc5cbe30ed5269ff2dfc8527c63"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all load balancer domains.

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

<a id="canonical-b79b4456c3904a73a33cfa38d83f01702cd0426bfb3459c588a998d23bb12a4c"></a>

## Direct properties — csrf_policy.all_load_balancer_domains / 58ae450672e4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bbd52959644b27e1458f55a36bf40ef5c19d349b4b7f63722cad57c3a857fa15"></a>

## Next pages — csrf_policy.all_load_balancer_domains / 58ae450672e4 / 4

- [csrf_policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-f1e6471143974860830fd021444acd48e74e2b1b6f98ec945d6dd4b24d617417)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-2e1041aaa99491d4c3772a836d7e3c55ff98583b7c54293dc8e78b61d19ab1ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e6d886416f2138fa4cb6e0320a21dba6f0827563f7854c585d028dbd6684b9ac"></a>

## csrf_policy.custom_domain_list — csrf_policy.custom_domain_list / e161ac5ebf2c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [csrf_policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-f1e6471143974860830fd021444acd48e74e2b1b6f98ec945d6dd4b24d617417)
- csrf_policy.custom_domain_list

<a id="canonical-e74179403765b3caa4353263ce0d961ad0512650ef5ead3baae881034a612ba6"></a>

Type: `"single"`. Computed.

List of domain names used for Host header matching.

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

<a id="canonical-2ffcbaf9860ec8df32fe32ca6bdd95f6c5cb6500a4fd572b8fffecfb05be7993"></a>

## Direct properties — csrf_policy.custom_domain_list / e161ac5ebf2c / 3

<a id="canonical-71c1f40bacae46cb4df329f1b0b7dba6de1e5d537e9560e8a2e9e01da00fae24"></a>

<a id="canonical-53979ccd8774e4a7670da7b851e76563c1244857c5b21b3019a2cd9bbc29aa7c"></a>

## domains property — csrf_policy.custom_domain_list / e161ac5ebf2c / 4

Type: `["list", "string"]`. Computed.

List of domain names that will be matched to loadbalancer. These domains are not used for SNI match.
Wildcard names are supported in the suffix or prefix form.

Upstream description:

A list of domain names that will be matched to loadbalancer. These domains are not used for SNI
match. Wildcard names are supported in the suffix or prefix form.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-8f8c62cfef7eadbfd196ae5bb1e97b0def017116f31de33016649cc72eb837f3"></a>

## Next pages — csrf_policy.custom_domain_list / e161ac5ebf2c / 5

- [csrf_policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-f1e6471143974860830fd021444acd48e74e2b1b6f98ec945d6dd4b24d617417)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-7614e339b1bdbe39d979bfc35dfe573113a2407c643d5fad0118bcc8fe93fb20"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fbf690d4c5f1f6b255c30ccfa5010b548711299467f23c875f35162e7ac46485"></a>

## csrf_policy.disabled — csrf_policy.disabled / bc9163dee62b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [csrf_policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-f1e6471143974860830fd021444acd48e74e2b1b6f98ec945d6dd4b24d617417)
- csrf_policy.disabled

<a id="canonical-b3075ed6b1cd4e6fa3e5c96553ecdbdaaf19f94705b0caa44c245c5f02c5be92"></a>

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

<a id="canonical-b01de9fada06dd3738cc09887aa85d3508568f3e85f943f855ca7556721242f2"></a>

## Direct properties — csrf_policy.disabled / bc9163dee62b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d74a2914cdef6dc26690c2726f5e4db347e2d8fb6d6766c96f40fbde74591141"></a>

## Next pages — csrf_policy.disabled / bc9163dee62b / 4

- [csrf_policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-f1e6471143974860830fd021444acd48e74e2b1b6f98ec945d6dd4b24d617417)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d2ef521c5a4c109fc00a5a94ff27b673fdc44185c8a89cd3ede83061de1f3343"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b866a137c46375c0d9a66a1ec1f6abea42b1a49732822a06cea9cabf547abe2c"></a>

## data_guard_rules — data_guard_rules / b2b34813aca0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- data_guard_rules

<a id="canonical-282985b0349ae82bfd0c941662e61392305bb82ada3b4745ae606850b4948720"></a>

Type: `"list"`. Computed.

Data Guard prevents responses from exposing sensitive information by masking the data. The system
masks credit card numbers and social security numbers leaked from the application from within the
HTTP response with a string of asterisks (\*).

Upstream description:

Data Guard prevents responses from exposing sensitive information by masking the data. The system
masks credit card numbers and social security numbers leaked from the application from within the
HTTP response with a string of asterisks (\*). Note: App Firewall should be enabled, to use Data
Guard feature.

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
    }
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

<a id="canonical-b1678e3443c20a186208414675255485a3ac0362fdca6871a0bb0872841e31fc"></a>

## Direct properties — data_guard_rules / b2b34813aca0 / 3

- [any_domain](data-sources--http_loadbalancer--reference--group-014.md#canonical-30c91f38e581b0064d08346702f91ec61dad8824fa77d9233179a3a1d0d2c895): complete subsection reference.

- [apply_data_guard](data-sources--http_loadbalancer--reference--group-014.md#canonical-2ee4a39abef89690b88272247edc125f7c0bdf6275e0ef43368fd9c858c53f7d): complete subsection reference.

<a id="canonical-619470489d6cfd7547a9c23fe4dec303ab1a38f7f29996723efaabbf0b7fdbe8"></a>

<a id="canonical-4de9679853d8fd782998708c205f6dba4c4a95c89791c16f0fb4e9a1f8175b6d"></a>

## exact_value property — data_guard_rules / b2b34813aca0 / 4

Type: `"string"`. Computed.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

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

- [metadata](data-sources--http_loadbalancer--reference--group-014.md#canonical-1f0895d9cd9c1399d0fa8e9ff032c1e63d2866041633c1958a6934a9079ebc64): complete subsection reference.

- [path](data-sources--http_loadbalancer--reference--group-014.md#canonical-218174d8662e256c4a8b441c9dcd73ad00046b2e59f4cfe553a68c853eb0993d): complete subsection reference.

- [skip_data_guard](data-sources--http_loadbalancer--reference--group-014.md#canonical-9582945e523344e21b75f1c98005cbc1b9ecff92621c4bc902cd85227ad23266): complete subsection reference.

<a id="canonical-5d0a2be038da00bc2335ee6add70672c6f9cb9df9e2dba0f67ac3ab5299cb760"></a>

<a id="canonical-a8165ec996a7fcb4affa8bb89fc501c5beab72db13a684a4cd883f969c0c82a2"></a>

## suffix_value property — data_guard_rules / b2b34813aca0 / 5

Type: `"string"`. Computed.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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

<a id="canonical-24b703bb3021ae93ab0dbe9a2ef5f7f7ef17f4b1a20db83ce65dad135f8ee391"></a>

## Next pages — data_guard_rules / b2b34813aca0 / 6

- [data_guard_rules.any_domain](data-sources--http_loadbalancer--reference--group-014.md#canonical-30c91f38e581b0064d08346702f91ec61dad8824fa77d9233179a3a1d0d2c895)
- [data_guard_rules.apply_data_guard](data-sources--http_loadbalancer--reference--group-014.md#canonical-2ee4a39abef89690b88272247edc125f7c0bdf6275e0ef43368fd9c858c53f7d)
- [data_guard_rules.metadata](data-sources--http_loadbalancer--reference--group-014.md#canonical-1f0895d9cd9c1399d0fa8e9ff032c1e63d2866041633c1958a6934a9079ebc64)
- [data_guard_rules.path](data-sources--http_loadbalancer--reference--group-014.md#canonical-218174d8662e256c4a8b441c9dcd73ad00046b2e59f4cfe553a68c853eb0993d)
- [data_guard_rules.skip_data_guard](data-sources--http_loadbalancer--reference--group-014.md#canonical-9582945e523344e21b75f1c98005cbc1b9ecff92621c4bc902cd85227ad23266)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-30c91f38e581b0064d08346702f91ec61dad8824fa77d9233179a3a1d0d2c895"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ec6bf32005af78f6d237bca49456d1994382881ded938beaa506cc0819991ff"></a>

## data_guard_rules.any_domain — data_guard_rules.any_domain / da07fcbfe7f9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [data_guard_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-d2ef521c5a4c109fc00a5a94ff27b673fdc44185c8a89cd3ede83061de1f3343)
- data_guard_rules.any_domain

<a id="canonical-4f8ba5bdd26ca55a593a8c7bb7b1408caac072a2e8b3f145fbc4a56e1897d4ff"></a>

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

<a id="canonical-0e1da0e2e866fa65055a661b677e1a38bc2dbcb83353de782134cc8359d0df4b"></a>

## Direct properties — data_guard_rules.any_domain / da07fcbfe7f9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ce5b5b4a296cd3e97d2d6f8bf35b39e9f78a7c916a3570cc53fb9e7ae599ea10"></a>

## Next pages — data_guard_rules.any_domain / da07fcbfe7f9 / 4

- [data_guard_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-d2ef521c5a4c109fc00a5a94ff27b673fdc44185c8a89cd3ede83061de1f3343)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-2ee4a39abef89690b88272247edc125f7c0bdf6275e0ef43368fd9c858c53f7d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1e0cd7687929f42eae4f89b945bc7a202c95cb59548df34fa03d9aefca20ea2"></a>

## data_guard_rules.apply_data_guard — data_guard_rules.apply_data_guard / 570edf0805a5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [data_guard_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-d2ef521c5a4c109fc00a5a94ff27b673fdc44185c8a89cd3ede83061de1f3343)
- data_guard_rules.apply_data_guard

<a id="canonical-dbf1cc16a9441d865b2cb8b859f872e7bf97e430dae9c3f423fcdbb13f9ae321"></a>

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

<a id="canonical-49f54f399b7c255b12f5f7f205bab16148cb053cd64efc933d6a7e486ce22100"></a>

## Direct properties — data_guard_rules.apply_data_guard / 570edf0805a5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-df133647aac48de1a88d71df91716066e7ff4b6a8c74db565f15cfe3722137d2"></a>

## Next pages — data_guard_rules.apply_data_guard / 570edf0805a5 / 4

- [data_guard_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-d2ef521c5a4c109fc00a5a94ff27b673fdc44185c8a89cd3ede83061de1f3343)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1f0895d9cd9c1399d0fa8e9ff032c1e63d2866041633c1958a6934a9079ebc64"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7666442087e6e7520eb48d1103e698899ff629c58f761a4922c5dfee7321f59f"></a>

## data_guard_rules.metadata — data_guard_rules.metadata / fdca88ac19f4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [data_guard_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-d2ef521c5a4c109fc00a5a94ff27b673fdc44185c8a89cd3ede83061de1f3343)
- data_guard_rules.metadata

<a id="canonical-45113001fc2c87a71604adf03bb3be2e88d6ec65992eed507b437a65970f1d48"></a>

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

<a id="canonical-4e8aad6fe014945c61995effe802faa0d47e106de96165676d4a4d655da807d3"></a>

## Direct properties — data_guard_rules.metadata / fdca88ac19f4 / 3

<a id="canonical-81d6a6422f60e43f47903a983a490af9f8b56575d50da00fa92888a25f3d90c7"></a>

<a id="canonical-ff5c1a07ddeb48a2553d0e18dfc0e2b6e41d98af620db00b88a40cce04a6aac5"></a>

## description_spec property — data_guard_rules.metadata / fdca88ac19f4 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-1c34ef51d37b9eb3600c3ede67a40a85e7341b1e769d6d7979b1e31447ea603d"></a>

<a id="canonical-4d6143a8310b80c120b3bd005dc9b4733dd2e58a4fec42334f1f6902474b9ccb"></a>

## name property — data_guard_rules.metadata / fdca88ac19f4 / 5

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

<a id="canonical-456cf55a5690fb3bb595f53bc6d7dbb0a0675c63db640bc3b5563aa63da155fe"></a>

## Next pages — data_guard_rules.metadata / fdca88ac19f4 / 6

- [data_guard_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-d2ef521c5a4c109fc00a5a94ff27b673fdc44185c8a89cd3ede83061de1f3343)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-218174d8662e256c4a8b441c9dcd73ad00046b2e59f4cfe553a68c853eb0993d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-825ca754b9b49948be3bde3dc23a66e06a76f3fdf246a83345da037d7a8d5bde"></a>

## data_guard_rules.path — data_guard_rules.path / 7801bd22ae49 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [data_guard_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-d2ef521c5a4c109fc00a5a94ff27b673fdc44185c8a89cd3ede83061de1f3343)
- data_guard_rules.path

<a id="canonical-c08c0e990034d444961c23f31cb8c7fa62505d44adb771d586bafba4cc5d192c"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

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

<a id="canonical-b0ce6f7e79ccfe6f53a5f5cd0b91068e1e0ff1bd6a3da0d1f2d1ba2bc9511a64"></a>

## Direct properties — data_guard_rules.path / 7801bd22ae49 / 3

<a id="canonical-edfd96cdee744e0a2c4e4eda9343e87492e690491940bf1b87c06f3eb076c661"></a>

<a id="canonical-9dcbaa80c9064c65299a28e3a5d0c7318d84e7b5f030b1693dd5dd3efb519ddd"></a>

## path property — data_guard_rules.path / 7801bd22ae49 / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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

<a id="canonical-ab237dc7cd9b27a82144d0bfc41aef9716c92caeb8f0931a22f38121310782c2"></a>

<a id="canonical-7120d1070728fd44379d867ddaf526cc3bac6714b45caf999b593fa4dc3b6796"></a>

## prefix property — data_guard_rules.path / 7801bd22ae49 / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-65b05765645a2f3d1e7c5c07b1e0189b3bdf21138f0536eae8b0eb2558525f83"></a>

<a id="canonical-36bc21aeac25b2edb4e727c6c6d00610351f8d0541f3129f417e65233b8cbf0b"></a>

## regex property — data_guard_rules.path / 7801bd22ae49 / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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

<a id="canonical-1ec73bc321fdf045da055542781045df9ed1aa3cd36c55602dfa13ccf8ad15dc"></a>

## Next pages — data_guard_rules.path / 7801bd22ae49 / 7

- [data_guard_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-d2ef521c5a4c109fc00a5a94ff27b673fdc44185c8a89cd3ede83061de1f3343)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9582945e523344e21b75f1c98005cbc1b9ecff92621c4bc902cd85227ad23266"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a96a1a77206c232994c26b73d8e939031848c898e43fb69e7d5c6f7d80b6c24b"></a>

## data_guard_rules.skip_data_guard — data_guard_rules.skip_data_guard / 97669865adde / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [data_guard_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-d2ef521c5a4c109fc00a5a94ff27b673fdc44185c8a89cd3ede83061de1f3343)
- data_guard_rules.skip_data_guard

<a id="canonical-9629cb3111d4c0894ea5fd84ac3a093588d74dd76a76621b22b6207f9eeffed2"></a>

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

<a id="canonical-61025e392dc2d54aab135df6cb581a12c802d483df6deaca60892d49231dd13d"></a>

## Direct properties — data_guard_rules.skip_data_guard / 97669865adde / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9bce515a1d870a5ac4a1d768ae06bd4f6695583c8ee77c2ebde81bfeb1d4baa6"></a>

## Next pages — data_guard_rules.skip_data_guard / 97669865adde / 4

- [data_guard_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-d2ef521c5a4c109fc00a5a94ff27b673fdc44185c8a89cd3ede83061de1f3343)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-31832e4dcb109b76102f24b4362682d28eb3957fb05c9480436f8c3c2c1d0201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9a4b350075a1a9a8c44d2e999083235ca20a2cc01a911666b1b45b0c166fcf3a"></a>

## ddos_mitigation_rules — ddos_mitigation_rules / 58847ad5649c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- ddos_mitigation_rules

<a id="canonical-2af1895b058be80796c6351a54cfa37e53748edabb41d6d3e294fa5744cc82d6"></a>

Type: `"list"`. Computed.

Define manual mitigation rules to block L7 DDoS attacks.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-7b8c310a7ab25e2f1bb04d7780a1dbc885d8678d85ffff734092b8da42e2abf5"></a>

## Direct properties — ddos_mitigation_rules / 58847ad5649c / 3

- [block](data-sources--http_loadbalancer--reference--group-014.md#canonical-dbd52bcf9fa021836028edb58f7a374f327c3912257fae6d4b22a5fda0ca7132): complete subsection reference.

- [ddos_client_source](data-sources--http_loadbalancer--reference--group-014.md#canonical-847beb6b79870ebe431a0f0bbae5d49ebc3c3e049fc0cd526552ae24618040a3): complete subsection reference.

<a id="canonical-a341b78f5aaef8b88e585842fe6530697c8df9c6c5282a7aefc4b5ef056f20fc"></a>

<a id="canonical-40be7d2bdfc85fc7af4dc530f1b9b15b93caa0090166329b06131a0ae3619562"></a>

## expiration_timestamp property — ddos_mitigation_rules / 58847ad5649c / 4

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  }
}
```

- [ip_prefix_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-a5b4a66c45ee66419f7da273fe324887a4ef9662fb7e88345b475c4f52553274): complete subsection reference.

- [metadata](data-sources--http_loadbalancer--reference--group-014.md#canonical-e5aea5fb48307c033abf4b127fc58abcbcc307958cdd290bf1bd15ac05c10686): complete subsection reference.

<a id="canonical-ba1463b0c160952e070140a3ddf15c21ef138bf6749259890ffb81a10e101f67"></a>

## Next pages — ddos_mitigation_rules / 58847ad5649c / 5

- [ddos_mitigation_rules.block](data-sources--http_loadbalancer--reference--group-014.md#canonical-dbd52bcf9fa021836028edb58f7a374f327c3912257fae6d4b22a5fda0ca7132)
- [ddos_mitigation_rules.ddos_client_source](data-sources--http_loadbalancer--reference--group-014.md#canonical-847beb6b79870ebe431a0f0bbae5d49ebc3c3e049fc0cd526552ae24618040a3)
- [ddos_mitigation_rules.ip_prefix_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-a5b4a66c45ee66419f7da273fe324887a4ef9662fb7e88345b475c4f52553274)
- [ddos_mitigation_rules.metadata](data-sources--http_loadbalancer--reference--group-014.md#canonical-e5aea5fb48307c033abf4b127fc58abcbcc307958cdd290bf1bd15ac05c10686)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-dbd52bcf9fa021836028edb58f7a374f327c3912257fae6d4b22a5fda0ca7132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-651527996c35427e84ec0e4b359b00f209c266fd443663aea015725a2f1ae304"></a>

## ddos_mitigation_rules.block — ddos_mitigation_rules.block / 20de64ac8f02 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [ddos_mitigation_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-31832e4dcb109b76102f24b4362682d28eb3957fb05c9480436f8c3c2c1d0201)
- ddos_mitigation_rules.block

<a id="canonical-641b90a0ff57122ba64dee3e39716e0a36fb935b7c7ac6b77f421e638c020e7b"></a>

Type: `"single"`. Computed.

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

<a id="canonical-438998b833d7e0f48bbc9b3b075aaf385f1568aa5528319bc656a5c034b8544f"></a>

## Direct properties — ddos_mitigation_rules.block / 20de64ac8f02 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-de9b53657d2510cbef0131badd3a7777236fb0b13232f04f12041f12ef4c6977"></a>

## Next pages — ddos_mitigation_rules.block / 20de64ac8f02 / 4

- [ddos_mitigation_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-31832e4dcb109b76102f24b4362682d28eb3957fb05c9480436f8c3c2c1d0201)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-847beb6b79870ebe431a0f0bbae5d49ebc3c3e049fc0cd526552ae24618040a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c79d917ce6c7e75f81bf74a9f0a4d0dcb81d283c2b4db499719c6d56e75da0f5"></a>

## ddos_mitigation_rules.ddos_client_source — ddos_mitigation_rules.ddos_client_source / 33b934c7e50f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [ddos_mitigation_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-31832e4dcb109b76102f24b4362682d28eb3957fb05c9480436f8c3c2c1d0201)
- ddos_mitigation_rules.ddos_client_source

<a id="canonical-8e0ed1e84b938c83ad72cde1ede96f3e89f760c955f7ccd84b6ef1874fb5efbb"></a>

Type: `"single"`. Computed.

DDoS Client Source Choice. DDoS Mitigation sources to be blocked.

Upstream description:

DDoS Mitigation sources to be blocked.

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

<a id="canonical-2f9e2729961756a2453506acf824a99b868827792b998a26c977aa64af2bf6ce"></a>

## Direct properties — ddos_mitigation_rules.ddos_client_source / 33b934c7e50f / 3

- [asn_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-fc3bc1c0657df6849c69062c1012639dbe9357af11cac641673db5d0bed620ae): complete subsection reference.

<a id="canonical-1af167af247323a38e2c09f2b9996130fd79e2b953d075e719561218130f4179"></a>

<a id="canonical-864fbda82da792c584d9c8c6bd73a05310a3519ead914789860c3a041efd6a5d"></a>

## country_list property — ddos_mitigation_rules.ddos_client_source / 33b934c7e50f / 4

Type: `["list", "string"]`. Computed.

\[Enum:
COUNTRY\_NONE|COUNTRY\_AD|COUNTRY\_AE|COUNTRY\_AF|COUNTRY\_AG|COUNTRY\_AI|COUNTRY\_AL|COUNTRY\_AM|COUNTRY\_AN|COUNTRY\_AO|COUNTRY\_AQ|COUNTRY\_AR|COUNTRY\_AS|COUNTRY\_AT|COUNTRY\_AU|COUNTRY\_AW|COUNTRY\_AX|COUNTRY\_AZ|COUNTRY\_BA|COUNTRY\_BB|COUNTRY\_BD|COUNTRY\_BE|COUNTRY\_BF|COUNTRY\_BG|COUNTRY\_BH|COUNTRY\_BI|COUNTRY\_BJ|COUNTRY\_BL|COUNTRY\_BM|COUNTRY\_BN|COUNTRY\_BO|COUNTRY\_BQ|COUNTRY\_BR|COUNTRY\_BS|COUNTRY\_BT|COUNTRY\_BV|COUNTRY\_BW|COUNTRY\_BY|COUNTRY\_BZ|COUNTRY\_CA|COUNTRY\_CC|COUNTRY\_CD|COUNTRY\_CF|COUNTRY\_CG|COUNTRY\_CH|COUNTRY\_CI|COUNTRY\_CK|COUNTRY\_CL|COUNTRY\_CM|COUNTRY\_CN|COUNTRY\_CO|COUNTRY\_CR|COUNTRY\_CS|COUNTRY\_CU|COUNTRY\_CV|COUNTRY\_CW|COUNTRY\_CX|COUNTRY\_CY|COUNTRY\_CZ|COUNTRY\_DE|COUNTRY\_DJ|COUNTRY\_DK|COUNTRY\_DM|COUNTRY\_DO|COUNTRY\_DZ|COUNTRY\_EC|COUNTRY\_EE|COUNTRY\_EG|COUNTRY\_EH|COUNTRY\_ER|COUNTRY\_ES|COUNTRY\_ET|COUNTRY\_FI|COUNTRY\_FJ|COUNTRY\_FK|COUNTRY\_FM|COUNTRY\_FO|COUNTRY\_FR|COUNTRY\_GA|COUNTRY\_GB|COUNTRY\_GD|COUNTRY\_GE|COUNTRY\_GF|COUNTRY\_GG|COUNTRY\_GH|COUNTRY\_GI|COUNTRY\_GL|COUNTRY\_GM|COUNTRY\_GN|COUNTRY\_GP|COUNTRY\_GQ|COUNTRY\_GR|COUNTRY\_GS|COUNTRY\_GT|COUNTRY\_GU|COUNTRY\_GW|COUNTRY\_GY|COUNTRY\_HK|COUNTRY\_HM|COUNTRY\_HN|COUNTRY\_HR|COUNTRY\_HT|COUNTRY\_HU|COUNTRY\_ID|COUNTRY\_IE|COUNTRY\_IL|COUNTRY\_IM|COUNTRY\_IN|COUNTRY\_IO|COUNTRY\_IQ|COUNTRY\_IR|COUNTRY\_IS|COUNTRY\_IT|COUNTRY\_JE|COUNTRY\_JM|COUNTRY\_JO|COUNTRY\_JP|COUNTRY\_KE|COUNTRY\_KG|COUNTRY\_KH|COUNTRY\_KI|COUNTRY\_KM|COUNTRY\_KN|COUNTRY\_KP|COUNTRY\_KR|COUNTRY\_KW|COUNTRY\_KY|COUNTRY\_KZ|COUNTRY\_LA|COUNTRY\_LB|COUNTRY\_LC|COUNTRY\_LI|COUNTRY\_LK|COUNTRY\_LR|COUNTRY\_LS|COUNTRY\_LT|COUNTRY\_LU|COUNTRY\_LV|COUNTRY\_LY|COUNTRY\_MA|COUNTRY\_MC|COUNTRY\_MD|COUNTRY\_ME|COUNTRY\_MF|COUNTRY\_MG|COUNTRY\_MH|COUNTRY\_MK|COUNTRY\_ML|COUNTRY\_MM|COUNTRY\_MN|COUNTRY\_MO|COUNTRY\_MP|COUNTRY\_MQ|COUNTRY\_MR|COUNTRY\_MS|COUNTRY\_MT|COUNTRY\_MU|COUNTRY\_MV|COUNTRY\_MW|COUNTRY\_MX|COUNTRY\_MY|COUNTRY\_MZ|COUNTRY\_NA|COUNTRY\_NC|COUNTRY\_NE|COUNTRY\_NF|COUNTRY\_NG|COUNTRY\_NI|COUNTRY\_NL|COUNTRY\_NO|COUNTRY\_NP|COUNTRY\_NR|COUNTRY\_NU|COUNTRY\_NZ|COUNTRY\_OM|COUNTRY\_PA|COUNTRY\_PE|COUNTRY\_PF|COUNTRY\_PG|COUNTRY\_PH|COUNTRY\_PK|COUNTRY\_PL|COUNTRY\_PM|COUNTRY\_PN|COUNTRY\_PR|COUNTRY\_PS|COUNTRY\_PT|COUNTRY\_PW|COUNTRY\_PY|COUNTRY\_QA|COUNTRY\_RE|COUNTRY\_RO|COUNTRY\_RS|COUNTRY\_RU|COUNTRY\_RW|COUNTRY\_SA|COUNTRY\_SB|COUNTRY\_SC|COUNTRY\_SD|COUNTRY\_SE|COUNTRY\_SG|COUNTRY\_SH|COUNTRY\_SI|COUNTRY\_SJ|COUNTRY\_SK|COUNTRY\_SL|COUNTRY\_SM|COUNTRY\_SN|COUNTRY\_SO|COUNTRY\_SR|COUNTRY\_SS|COUNTRY\_ST|COUNTRY\_SV|COUNTRY\_SX|COUNTRY\_SY|COUNTRY\_SZ|COUNTRY\_TC|COUNTRY\_TD|COUNTRY\_TF|COUNTRY\_TG|COUNTRY\_TH|COUNTRY\_TJ|COUNTRY\_TK|COUNTRY\_TL|COUNTRY\_TM|COUNTRY\_TN|COUNTRY\_TO|COUNTRY\_TR|COUNTRY\_TT|COUNTRY\_TV|COUNTRY\_TW|COUNTRY\_TZ|COUNTRY\_UA|COUNTRY\_UG|COUNTRY\_UM|COUNTRY\_US|COUNTRY\_UY|COUNTRY\_UZ|COUNTRY\_VA|COUNTRY\_VC|COUNTRY\_VE|COUNTRY\_VG|COUNTRY\_VI|COUNTRY\_VN|COUNTRY\_VU|COUNTRY\_WF|COUNTRY\_WS|COUNTRY\_XK|COUNTRY\_XT|COUNTRY\_YE|COUNTRY\_YT|COUNTRY\_ZA|COUNTRY\_ZM|COUNTRY\_ZW\]
Sources that are located in one of the countries in the given list. Possible values are
\`COUNTRY\_NONE\`, \`COUNTRY\_AD\`, \`COUNTRY\_AE\`, \`COUNTRY\_AF\`, \`COUNTRY\_AG\`,
\`COUNTRY\_AI\`, \`COUNTRY\_AL\`, \`COUNTRY\_AM\`, \`COUNTRY\_AN\`, \`COUNTRY\_AO\`,
\`COUNTRY\_AQ\`, \`COUNTRY\_AR\`, \`COUNTRY\_AS\`, \`COUNTRY\_AT\`, \`COUNTRY\_AU\`,
\`COUNTRY\_AW\`, \`COUNTRY\_AX\`, \`COUNTRY\_AZ\`, \`COUNTRY\_BA\`, \`COUNTRY\_BB\`,
\`COUNTRY\_BD\`, \`COUNTRY\_BE\`, \`COUNTRY\_BF\`, \`COUNTRY\_BG\`, \`COUNTRY\_BH\`,
\`COUNTRY\_BI\`, \`COUNTRY\_BJ\`, \`COUNTRY\_BL\`, \`COUNTRY\_BM\`, \`COUNTRY\_BN\`,
\`COUNTRY\_BO\`, \`COUNTRY\_BQ\`, \`COUNTRY\_BR\`, \`COUNTRY\_BS\`, \`COUNTRY\_BT\`,
\`COUNTRY\_BV\`, \`COUNTRY\_BW\`, \`COUNTRY\_BY\`, \`COUNTRY\_BZ\`, \`COUNTRY\_CA\`,
\`COUNTRY\_CC\`, \`COUNTRY\_CD\`, \`COUNTRY\_CF\`, \`COUNTRY\_CG\`, \`COUNTRY\_CH\`,
\`COUNTRY\_CI\`, \`COUNTRY\_CK\`, \`COUNTRY\_CL\`, \`COUNTRY\_CM\`, \`COUNTRY\_CN\`,
\`COUNTRY\_CO\`, \`COUNTRY\_CR\`, \`COUNTRY\_CS\`, \`COUNTRY\_CU\`, \`COUNTRY\_CV\`,
\`COUNTRY\_CW\`, \`COUNTRY\_CX\`, \`COUNTRY\_CY\`, \`COUNTRY\_CZ\`, \`COUNTRY\_DE\`,
\`COUNTRY\_DJ\`, \`COUNTRY\_DK\`, \`COUNTRY\_DM\`, \`COUNTRY\_DO\`, \`COUNTRY\_DZ\`,
\`COUNTRY\_EC\`, \`COUNTRY\_EE\`, \`COUNTRY\_EG\`, \`COUNTRY\_EH\`, \`COUNTRY\_ER\`,
\`COUNTRY\_ES\`, \`COUNTRY\_ET\`, \`COUNTRY\_FI\`, \`COUNTRY\_FJ\`, \`COUNTRY\_FK\`,
\`COUNTRY\_FM\`, \`COUNTRY\_FO\`, \`COUNTRY\_FR\`, \`COUNTRY\_GA\`, \`COUNTRY\_GB\`,
\`COUNTRY\_GD\`, \`COUNTRY\_GE\`, \`COUNTRY\_GF\`, \`COUNTRY\_GG\`, \`COUNTRY\_GH\`,
\`COUNTRY\_GI\`, \`COUNTRY\_GL\`, \`COUNTRY\_GM\`, \`COUNTRY\_GN\`, \`COUNTRY\_GP\`,
\`COUNTRY\_GQ\`, \`COUNTRY\_GR\`, \`COUNTRY\_GS\`, \`COUNTRY\_GT\`, \`COUNTRY\_GU\`,
\`COUNTRY\_GW\`, \`COUNTRY\_GY\`, \`COUNTRY\_HK\`, \`COUNTRY\_HM\`, \`COUNTRY\_HN\`,
\`COUNTRY\_HR\`, \`COUNTRY\_HT\`, \`COUNTRY\_HU\`, \`COUNTRY\_ID\`, \`COUNTRY\_IE\`,
\`COUNTRY\_IL\`, \`COUNTRY\_IM\`, \`COUNTRY\_IN\`, \`COUNTRY\_IO\`, \`COUNTRY\_IQ\`,
\`COUNTRY\_IR\`, \`COUNTRY\_IS\`, \`COUNTRY\_IT\`, \`COUNTRY\_JE\`, \`COUNTRY\_JM\`,
\`COUNTRY\_JO\`, \`COUNTRY\_JP\`, \`COUNTRY\_KE\`, \`COUNTRY\_KG\`, \`COUNTRY\_KH\`,
\`COUNTRY\_KI\`, \`COUNTRY\_KM\`, \`COUNTRY\_KN\`, \`COUNTRY\_KP\`, \`COUNTRY\_KR\`,
\`COUNTRY\_KW\`, \`COUNTRY\_KY\`, \`COUNTRY\_KZ\`, \`COUNTRY\_LA\`, \`COUNTRY\_LB\`,
\`COUNTRY\_LC\`, \`COUNTRY\_LI\`, \`COUNTRY\_LK\`, \`COUNTRY\_LR\`, \`COUNTRY\_LS\`,
\`COUNTRY\_LT\`, \`COUNTRY\_LU\`, \`COUNTRY\_LV\`, \`COUNTRY\_LY\`, \`COUNTRY\_MA\`,
\`COUNTRY\_MC\`, \`COUNTRY\_MD\`, \`COUNTRY\_ME\`, \`COUNTRY\_MF\`, \`COUNTRY\_MG\`,
\`COUNTRY\_MH\`, \`COUNTRY\_MK\`, \`COUNTRY\_ML\`, \`COUNTRY\_MM\`, \`COUNTRY\_MN\`,
\`COUNTRY\_MO\`, \`COUNTRY\_MP\`, \`COUNTRY\_MQ\`, \`COUNTRY\_MR\`, \`COUNTRY\_MS\`,
\`COUNTRY\_MT\`, \`COUNTRY\_MU\`, \`COUNTRY\_MV\`, \`COUNTRY\_MW\`, \`COUNTRY\_MX\`,
\`COUNTRY\_MY\`, \`COUNTRY\_MZ\`, \`COUNTRY\_NA\`, \`COUNTRY\_NC\`, \`COUNTRY\_NE\`,
\`COUNTRY\_NF\`, \`COUNTRY\_NG\`, \`COUNTRY\_NI\`, \`COUNTRY\_NL\`, \`COUNTRY\_NO\`,
\`COUNTRY\_NP\`, \`COUNTRY\_NR\`, \`COUNTRY\_NU\`, \`COUNTRY\_NZ\`, \`COUNTRY\_OM\`,
\`COUNTRY\_PA\`, \`COUNTRY\_PE\`, \`COUNTRY\_PF\`, \`COUNTRY\_PG\`, \`COUNTRY\_PH\`,
\`COUNTRY\_PK\`, \`COUNTRY\_PL\`, \`COUNTRY\_PM\`, \`COUNTRY\_PN\`, \`COUNTRY\_PR\`,
\`COUNTRY\_PS\`, \`COUNTRY\_PT\`, \`COUNTRY\_PW\`, \`COUNTRY\_PY\`, \`COUNTRY\_QA\`,
\`COUNTRY\_RE\`, \`COUNTRY\_RO\`, \`COUNTRY\_RS\`, \`COUNTRY\_RU\`, \`COUNTRY\_RW\`,
\`COUNTRY\_SA\`, \`COUNTRY\_SB\`, \`COUNTRY\_SC\`, \`COUNTRY\_SD\`, \`COUNTRY\_SE\`,
\`COUNTRY\_SG\`, \`COUNTRY\_SH\`, \`COUNTRY\_SI\`, \`COUNTRY\_SJ\`, \`COUNTRY\_SK\`,
\`COUNTRY\_SL\`, \`COUNTRY\_SM\`, \`COUNTRY\_SN\`, \`COUNTRY\_SO\`, \`COUNTRY\_SR\`,
\`COUNTRY\_SS\`, \`COUNTRY\_ST\`, \`COUNTRY\_SV\`, \`COUNTRY\_SX\`, \`COUNTRY\_SY\`,
\`COUNTRY\_SZ\`, \`COUNTRY\_TC\`, \`COUNTRY\_TD\`, \`COUNTRY\_TF\`, \`COUNTRY\_TG\`,
\`COUNTRY\_TH\`, \`COUNTRY\_TJ\`, \`COUNTRY\_TK\`, \`COUNTRY\_TL\`, \`COUNTRY\_TM\`,
\`COUNTRY\_TN\`, \`COUNTRY\_TO\`, \`COUNTRY\_TR\`, \`COUNTRY\_TT\`, \`COUNTRY\_TV\`,
\`COUNTRY\_TW\`, \`COUNTRY\_TZ\`, \`COUNTRY\_UA\`, \`COUNTRY\_UG\`, \`COUNTRY\_UM\`,
\`COUNTRY\_US\`, \`COUNTRY\_UY\`, \`COUNTRY\_UZ\`, \`COUNTRY\_VA\`, \`COUNTRY\_VC\`,
\`COUNTRY\_VE\`, \`COUNTRY\_VG\`, \`COUNTRY\_VI\`, \`COUNTRY\_VN\`, \`COUNTRY\_VU\`,
\`COUNTRY\_WF\`, \`COUNTRY\_WS\`, \`COUNTRY\_XK\`, \`COUNTRY\_XT\`, \`COUNTRY\_YE\`,
\`COUNTRY\_YT\`, \`COUNTRY\_ZA\`, \`COUNTRY\_ZM\`, \`COUNTRY\_ZW\`. Defaults to \`COUNTRY\_NONE\`.

Upstream description:

Sources that are located in one of the countries in the given list.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [ja4_tls_fingerprint_matcher](data-sources--http_loadbalancer--reference--group-014.md#canonical-8a45cb85a08e92ccae5a43d8ad67e363384d4b66333c2222277b777e132f830e): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--http_loadbalancer--reference--group-014.md#canonical-e9a00ebd53388ccbf31f2e35d54146b6c9bfd8a7128deac6b88a11251ec96ebc): complete subsection reference.

<a id="canonical-94fda7f6f5a53207ed5e510abd987bf853b31240bdc626b98ef1d450b95a4b40"></a>

## Next pages — ddos_mitigation_rules.ddos_client_source / 33b934c7e50f / 5

- [ddos_mitigation_rules.ddos_client_source.asn_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-fc3bc1c0657df6849c69062c1012639dbe9357af11cac641673db5d0bed620ae)
- [ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher](data-sources--http_loadbalancer--reference--group-014.md#canonical-8a45cb85a08e92ccae5a43d8ad67e363384d4b66333c2222277b777e132f830e)
- [ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher](data-sources--http_loadbalancer--reference--group-014.md#canonical-e9a00ebd53388ccbf31f2e35d54146b6c9bfd8a7128deac6b88a11251ec96ebc)
- [ddos_mitigation_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-31832e4dcb109b76102f24b4362682d28eb3957fb05c9480436f8c3c2c1d0201)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-fc3bc1c0657df6849c69062c1012639dbe9357af11cac641673db5d0bed620ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5b4746eee8108f39e1a68d76633ada559aa4c6265b587adf3fa167d64c4a3dc7"></a>

## ddos_mitigation_rules.ddos_client_source.asn_list — ddos_mitigation_rules.ddos_client_source.asn_list / 74fe8ffdbd93 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [ddos_mitigation_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-31832e4dcb109b76102f24b4362682d28eb3957fb05c9480436f8c3c2c1d0201)
- [ddos_mitigation_rules.ddos_client_source](data-sources--http_loadbalancer--reference--group-014.md#canonical-847beb6b79870ebe431a0f0bbae5d49ebc3c3e049fc0cd526552ae24618040a3)
- ddos_mitigation_rules.ddos_client_source.asn_list

<a id="canonical-cc611d21ddec9a18c24970a10771ca2f1c9196c665dade4671959c40ae71597e"></a>

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

<a id="canonical-e8aac29baaaa184465548f5eee142c68c120e9eafc32ddaec485e08ee3eccf1c"></a>

## Direct properties — ddos_mitigation_rules.ddos_client_source.asn_list / 74fe8ffdbd93 / 3

<a id="canonical-592a18d6f06a1f5e7d0dc564e5b7ee6a368a8ffd5f83d215befb1ccbff9184a8"></a>

<a id="canonical-dc72c42990978df772862e5ed40096b44b0720c143ee0191ec7ef45a05661a52"></a>

## as_numbers property — ddos_mitigation_rules.ddos_client_source.asn_list / 74fe8ffdbd93 / 4

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

<a id="canonical-28f683e87bfb34ff95b35010ba1f26c79e1a8623ab3614b3d76bee62f82aa70c"></a>

## Next pages — ddos_mitigation_rules.ddos_client_source.asn_list / 74fe8ffdbd93 / 5

- [ddos_mitigation_rules.ddos_client_source](data-sources--http_loadbalancer--reference--group-014.md#canonical-847beb6b79870ebe431a0f0bbae5d49ebc3c3e049fc0cd526552ae24618040a3)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-8a45cb85a08e92ccae5a43d8ad67e363384d4b66333c2222277b777e132f830e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f1fa0752e22ef3f83126aebefb5b1eca6f1c9b9e0aa23613a89a41460cb641f"></a>

## ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher — ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher / 5706e4b7ccee / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [ddos_mitigation_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-31832e4dcb109b76102f24b4362682d28eb3957fb05c9480436f8c3c2c1d0201)
- [ddos_mitigation_rules.ddos_client_source](data-sources--http_loadbalancer--reference--group-014.md#canonical-847beb6b79870ebe431a0f0bbae5d49ebc3c3e049fc0cd526552ae24618040a3)
- ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher

<a id="canonical-433f998ab82875a8db3edfb684dd4f240edc7bef52d6e271c401dcd3d4d362f5"></a>

Type: `"single"`. Computed.

Extended version of JA3 that includes additional fields for more comprehensive fingerprinting of
SSL/TLS clients and potentially has a different structure and length.

Upstream description:

An extended version of JA3 that includes additional fields for more comprehensive fingerprinting of
SSL/TLS clients and potentially has a different structure and length.

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

<a id="canonical-57f43da7dbc354f618df5ad4bcf67ba6637f33af2935da5a5ce38fc01cbe6e19"></a>

## Direct properties — ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher / 5706e4b7ccee / 3

<a id="canonical-146c73b026236ae1359a5f684d4cba25c5b6aba7d5c507c6d91aa4e25a0073c8"></a>

<a id="canonical-51c46a456607604038aad1855f1c91b45b3cbfeee23bf7f6d5afa9ae670f177f"></a>

## exact_values property — ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher / 5706e4b7ccee / 4

Type: `["list", "string"]`. Computed.

List of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.

Upstream description:

A list of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.

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
    "ves.io.schema.rules.repeated.items.string.len": "36",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "36",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-22180dfa4f55c36d3c7230d25f542d9701a5552f9f766266c72eaf83d28b9115"></a>

## Next pages — ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher / 5706e4b7ccee / 5

- [ddos_mitigation_rules.ddos_client_source](data-sources--http_loadbalancer--reference--group-014.md#canonical-847beb6b79870ebe431a0f0bbae5d49ebc3c3e049fc0cd526552ae24618040a3)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e9a00ebd53388ccbf31f2e35d54146b6c9bfd8a7128deac6b88a11251ec96ebc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-88dbfc7f10ccbab401bb8ae91a12eda97f2772ecc6ccc74aeb4beb0c3a1f49fc"></a>

## ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher — ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher / ace115d7708b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [ddos_mitigation_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-31832e4dcb109b76102f24b4362682d28eb3957fb05c9480436f8c3c2c1d0201)
- [ddos_mitigation_rules.ddos_client_source](data-sources--http_loadbalancer--reference--group-014.md#canonical-847beb6b79870ebe431a0f0bbae5d49ebc3c3e049fc0cd526552ae24618040a3)
- ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher

<a id="canonical-a105bcb8a19c493ee80a2899bf883e1cc1b120d98bb704ed2797f9c8a3cfd804"></a>

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

<a id="canonical-32454a6356087870a6317e753caa28cc80731ab675fe9f33e9df0ab38e865bac"></a>

## Direct properties — ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher / ace115d7708b / 3

<a id="canonical-9b9ac68705b15edd9206462714b14b48b7289f463d1d8530269fb45d21d2d76a"></a>

<a id="canonical-27ac51bd241bed694e2ba13145c00c7da610e3278b7f0abe864b2829c3d4f9c1"></a>

## classes property — ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher / ace115d7708b / 4

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

<a id="canonical-57a61391d9773da5de496375e7333c094cf3e4bb07fd3d49f78ffb988d7189b5"></a>

<a id="canonical-263256bb4b44dd4bc40d7d2934d5df000d0242dfca5bef6b11abac9337f3b125"></a>

## exact_values property — ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher / ace115d7708b / 5

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

<a id="canonical-21c14d080963235f7b57d75a918cd4be35629aa69f23b83387c57034f962559c"></a>

<a id="canonical-39b4b8c4e142df3385f9ac442eb6d3ce42d4af4bdb248e5b98c0120ced18946d"></a>

## excluded_values property — ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher / ace115d7708b / 6

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

<a id="canonical-f9e4b167dee2bfb3fe2f33599f4bc0426ed2ed6ae0cb553aea5bef1a9a420f0a"></a>

## Next pages — ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher / ace115d7708b / 7

- [ddos_mitigation_rules.ddos_client_source](data-sources--http_loadbalancer--reference--group-014.md#canonical-847beb6b79870ebe431a0f0bbae5d49ebc3c3e049fc0cd526552ae24618040a3)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a5b4a66c45ee66419f7da273fe324887a4ef9662fb7e88345b475c4f52553274"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-53d0e48e36c8ee35a935f258359b4415d5e7744cd792777b9a5bbc939ba92398"></a>

## ddos_mitigation_rules.ip_prefix_list — ddos_mitigation_rules.ip_prefix_list / 02cd2b19db8d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [ddos_mitigation_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-31832e4dcb109b76102f24b4362682d28eb3957fb05c9480436f8c3c2c1d0201)
- ddos_mitigation_rules.ip_prefix_list

<a id="canonical-e193f00d59265647f26d0426c9b62adc93ca229dab8dacdbc5622b2420a50cad"></a>

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

<a id="canonical-63dd27047eae5bcbaa87c565c667eb9459ade074020ba391d70c251d9f238222"></a>

## Direct properties — ddos_mitigation_rules.ip_prefix_list / 02cd2b19db8d / 3

<a id="canonical-4c122ef2bb035846ce48a032f1e471518c00c204b7945b5c501e276b77a98a20"></a>

<a id="canonical-89eadf0628bb71c0b3d6ff861b42319292da9037d146a1d205705802b99ee503"></a>

## invert_match property — ddos_mitigation_rules.ip_prefix_list / 02cd2b19db8d / 4

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

<a id="canonical-8db5a4cf3ae3eb54e30e603f302f26e66d93b823e56243db0b11a23db6caab80"></a>

<a id="canonical-67daaf11bd2a3537f61922e5708205ed2a19aca4697e81c5ee380766c239ab77"></a>

## ip_prefixes property — ddos_mitigation_rules.ip_prefix_list / 02cd2b19db8d / 5

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

<a id="canonical-f8d5799cc8d2e64ae781312ad387ee9b0c6e9df9147a570b45b3277669456e76"></a>

## Next pages — ddos_mitigation_rules.ip_prefix_list / 02cd2b19db8d / 6

- [ddos_mitigation_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-31832e4dcb109b76102f24b4362682d28eb3957fb05c9480436f8c3c2c1d0201)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e5aea5fb48307c033abf4b127fc58abcbcc307958cdd290bf1bd15ac05c10686"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-235008307b725397802e346934b543520926aecf566ccdf94db25b6ffd17b530"></a>

## ddos_mitigation_rules.metadata — ddos_mitigation_rules.metadata / 837c3bfd75b5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [ddos_mitigation_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-31832e4dcb109b76102f24b4362682d28eb3957fb05c9480436f8c3c2c1d0201)
- ddos_mitigation_rules.metadata

<a id="canonical-f10d7dbc9293aebb327d67e1b3fd5a857255f53438866db5f3df2bd5ea3a446a"></a>

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

<a id="canonical-7d0e50aab8d451bcd6860f9a6efc731d792e50fad87b87af65ff69f66816b179"></a>

## Direct properties — ddos_mitigation_rules.metadata / 837c3bfd75b5 / 3

<a id="canonical-45bb0434cc2c5df57ed8cc0d7c522cfa40d63405bbd31a8b57d599f9bb8a4e7b"></a>

<a id="canonical-b21dc732906093361fe2645ea07d4e94fbf959e775a17faa244266825e79afa4"></a>

## description_spec property — ddos_mitigation_rules.metadata / 837c3bfd75b5 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-204a0a25d90fe5e5289afb13998d3bad475c56ce02296f3caf5cbb718dcc8be7"></a>

<a id="canonical-d13b53a815c8f9fac7eda6821df983a918a0c239c51c78a31417c37c55f5f640"></a>

## name property — ddos_mitigation_rules.metadata / 837c3bfd75b5 / 5

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

<a id="canonical-482d2effceb1ec151d953b0c5b101334091c55c5d2d48c433d340b4d71ea28c0"></a>

## Next pages — ddos_mitigation_rules.metadata / 837c3bfd75b5 / 6

- [ddos_mitigation_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-31832e4dcb109b76102f24b4362682d28eb3957fb05c9480436f8c3c2c1d0201)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f84b4219435dd61021fd1d706f4fd696c2e03cb49b8bceea0677abbe01d481fa"></a>

## default_pool — default_pool / e1756bf405d3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- default_pool

<a id="canonical-97680757bdcc6b58774b122b94f621964a43741e306363bdc14b80107f9283f8"></a>

Type: `"single"`. Computed.

\[OneOf: default\_pool, default\_pool\_list; Default: default\_pool\] Configuration parameter for
default pool.

Upstream description:

Shape of the origin pool specification.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-health_check_port_choice": "[\"health_check_port\",\"same_as_endpoint_port\"]",
  "x-ves-oneof-field-port_choice": "[\"automatic_port\",\"lb_port\",\"port\"]",
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

OneOf alternatives in this subsection:

- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-97680757bdcc6b58774b122b94f621964a43741e306363bdc14b80107f9283f8)
- [default_pool_list](data-sources--http_loadbalancer--reference--group-016.md#canonical-5c3dedffe9b0bcb499f2f8cf2008c3c1d46d36456cce13f7c55e6b7a5fdb63ff)

Select alternatives according to the provider validators above.

<a id="canonical-38d0a889e131c2a5f8484b6d93ff60f8f2885ff562bb0f0622120d7132ffc416"></a>

## Direct properties — default_pool / e1756bf405d3 / 3

- [advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47): complete subsection reference.

- [automatic_port](data-sources--http_loadbalancer--reference--group-015.md#canonical-ff857da8254cb22d6d28b25edcf09e8945fc7f1069c626765bdb72560ecc69c3): complete subsection reference.

<a id="canonical-2223032dc450d24249fd3ebf8df9932d46a86f001f0b9b30dd04ec9afec32cc6"></a>

<a id="canonical-ea67a4d8f280cd2e056f7123d65da8ff26d83f96cab365078c995f695f0347cd"></a>

## endpoint_selection property — default_pool / e1756bf405d3 / 4

Type: `"string"`. Computed.

\[Enum: DISTRIBUTED|LOCAL\_ONLY|LOCAL\_PREFERRED\] Policy for selection of endpoints from local
site/remote site/both Consider both remote and local endpoints for load balancing LOCAL\_ONLY:
Consider only local endpoints for load balancing Enable this policy to load balance ONLY among
locally discovered endpoints Prefer the local endpoints for.. Possible values are \`DISTRIBUTED\`,
\`LOCAL\_ONLY\`, \`LOCAL\_PREFERRED\`. Defaults to \`DISTRIBUTED\`. Server applies default when
omitted.

Upstream description:

Policy for selection of endpoints from local site/remote site/both

Consider both remote and local endpoints for load balancing LOCAL\_ONLY: Consider only local
endpoints for load balancing Enable this policy to load balance ONLY among locally discovered
endpoints Prefer the local endpoints for load balancing. If local endpoints are not present remote
endpoints will be considered.

Receipt-pinned upstream constraints:

```json
{
  "default": "DISTRIBUTED",
  "enum": [
    "DISTRIBUTED",
    "LOCAL_ONLY",
    "LOCAL_PREFERRED"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-5346245d84e52341638ddcd769791d3714cea642aac434e8a079ef1b9d42fc1f"></a>

<a id="canonical-dd5145426efb19d94971d1290986f9f5c84bdd231e8ebc28051a5d3b22528f36"></a>

## health_check_port property — default_pool / e1756bf405d3 / 5

Type: `"number"`. Computed.

Exclusive with \[same\_as\_endpoint\_port\] Port used for performing health check.

Upstream description:

Exclusive with \[same\_as\_endpoint\_port\] Port used for performing health check.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "networking",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "category": "networking",
      "confidence": 0.99,
      "note": "Asymmetry: port enforces [1,65535], health_check_port allows 0",
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [healthcheck](data-sources--http_loadbalancer--reference--group-015.md#canonical-62f7672bfb3c3340703a1c669f486ca846258513ad74e17527fa62228935afc8): complete subsection reference.

- [lb_port](data-sources--http_loadbalancer--reference--group-015.md#canonical-9ad8ab5b9aa41cd455cbf362882fa6bd081274d406011ea56066430878b87477): complete subsection reference.

<a id="canonical-e1631586d1e91921d3030c1ba1c3c4004d8d99aa9fb846268856f3037e23ed7a"></a>

<a id="canonical-dd237c80c812044a69cf398918deb65ab34deff3b7d09496a6ea5d39889c86c6"></a>

## loadbalancer_algorithm property — default_pool / e1756bf405d3 / 6

Type: `"string"`. Computed.

\[Enum: ROUND\_ROBIN|LEAST\_REQUEST|RING\_HASH|RANDOM|LB\_OVERRIDE\] Different load balancing
algorithms supported When a connection to a endpoint in an upstream cluster is required, the load
balancer uses loadbalancer\_algorithm to determine which host is selected. - ROUND\_ROBIN:
ROUND\_ROBIN Policy in which each healthy/available upstream endpoint is selected in.. Possible
values are \`ROUND\_ROBIN\`, \`LEAST\_REQUEST\`, \`RING\_HASH\`, \`RANDOM\`, \`LB\_OVERRIDE\`.
Defaults to \`ROUND\_ROBIN\`. Server applies default when omitted.

Upstream description:

Different load balancing algorithms supported When a connection to a endpoint in an upstream cluster
is required, the load balancer uses loadbalancer\_algorithm to determine which host is selected.

&#8203;- ROUND\_ROBIN: ROUND\_ROBIN

Policy in which each healthy/available upstream endpoint is selected in round robin order. &#8203;-
LEAST\_REQUEST: LEAST\_REQUEST

Policy in which loadbalancer picks the upstream endpoint which has the fewest active requests
&#8203;- RING\_HASH: RING\_HASH

Policy implements consistent hashing to upstream endpoints using ring hash of endpoint names Hash of
the incoming request is calculated using request hash policy. The ring/modulo hash load balancer
implements consistent hashing to upstream hosts. The algorithm is based on mapping all hosts onto a
circle such that the addition or removal of a host from the host set changes only affect 1/N
requests. This technique is also commonly known as “ketama” hashing. A consistent hashing load
balancer is only effective when protocol routing is used that specifies a value to hash on. The
minimum ring size governs the replication factor for each host in the ring. For example, if the
minimum ring size is 1024 and there are 16 hosts, each host will be replicated 64 times. &#8203;-
RANDOM: RANDOM

Policy in which each available upstream endpoint is selected in random order. The random load
balancer selects a random healthy host. The random load balancer generally performs better than
round robin if no health checking policy is configured. Random selection avoids bias towards the
host in the set that comes after a failed host. &#8203;- LB\_OVERRIDE: Load Balancer Override

Hash policy is taken from from the load balancer which is using this origin pool.

Receipt-pinned upstream constraints:

```json
{
  "default": "ROUND_ROBIN",
  "enum": [
    "ROUND_ROBIN",
    "LEAST_REQUEST",
    "RING_HASH",
    "RANDOM",
    "LB_OVERRIDE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [no_tls](data-sources--http_loadbalancer--reference--group-015.md#canonical-529a47404046c3510b6d8e6b1b3c029d2ae5619501b1270b94491c2b636a29de): complete subsection reference.

- [origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0): complete subsection reference.

<a id="canonical-8a4337b54f12907134be0baa020da006652a527ac2089c6c28cb4d4dc48429c8"></a>

<a id="canonical-428f666475441029bb80da2056f26037fb9ca195271081a66bdc678b33880ecd"></a>

## port property — default_pool / e1756bf405d3 / 7

Type: `"number"`. Computed.

Exclusive with \[automatic\_port lb\_port\] Endpoint service is available on this port. Recommended:
\`443\`.

Upstream description:

Exclusive with \[automatic\_port lb\_port\] Endpoint service is available on this port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [same_as_endpoint_port](data-sources--http_loadbalancer--reference--group-016.md#canonical-68b742eb8998eeee0ab6a94544e8f3a137e49bf7bf3db74dc69111084a1ab5f5): complete subsection reference.

- [upstream_conn_pool_reuse_type](data-sources--http_loadbalancer--reference--group-016.md#canonical-16cee54849090462c0081b0bf98320e268d1a39c839416736c74d9d455f82f06): complete subsection reference.

- [use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d): complete subsection reference.

- [view_internal](data-sources--http_loadbalancer--reference--group-016.md#canonical-baa3be394b34d6e133fe6427378cd11e2a3f0092dddbc18782efd9521777f1bf): complete subsection reference.

<a id="canonical-6c79e1f45bb5f7ddb19346d417719bba71bf4db14cad79937001d14d33017551"></a>

## Next pages — default_pool / e1756bf405d3 / 8

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- [default_pool.automatic_port](data-sources--http_loadbalancer--reference--group-015.md#canonical-ff857da8254cb22d6d28b25edcf09e8945fc7f1069c626765bdb72560ecc69c3)
- [default_pool.healthcheck](data-sources--http_loadbalancer--reference--group-015.md#canonical-62f7672bfb3c3340703a1c669f486ca846258513ad74e17527fa62228935afc8)
- [default_pool.lb_port](data-sources--http_loadbalancer--reference--group-015.md#canonical-9ad8ab5b9aa41cd455cbf362882fa6bd081274d406011ea56066430878b87477)
- [default_pool.no_tls](data-sources--http_loadbalancer--reference--group-015.md#canonical-529a47404046c3510b6d8e6b1b3c029d2ae5619501b1270b94491c2b636a29de)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.same_as_endpoint_port](data-sources--http_loadbalancer--reference--group-016.md#canonical-68b742eb8998eeee0ab6a94544e8f3a137e49bf7bf3db74dc69111084a1ab5f5)
- [default_pool.upstream_conn_pool_reuse_type](data-sources--http_loadbalancer--reference--group-016.md#canonical-16cee54849090462c0081b0bf98320e268d1a39c839416736c74d9d455f82f06)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- [default_pool.view_internal](data-sources--http_loadbalancer--reference--group-016.md#canonical-baa3be394b34d6e133fe6427378cd11e2a3f0092dddbc18782efd9521777f1bf)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0866de943039e5e182ac05c8a1fd9bffdcbe49985be6a177d1d28728fd2d0d8"></a>

## default_pool.advanced_options — default_pool.advanced_options / 6b5ccddf6170 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- default_pool.advanced_options

<a id="canonical-6717be9ec410c681cacc05f0a747833669ad34e8b5f3e51a323188e4343feb58"></a>

Type: `"single"`. Computed.

Configure Advanced OPTIONS for origin pool.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-circuit_breaker_choice": "[\"circuit_breaker\",\"default_circuit_breaker\",\"disable_circuit_breaker\"]",
  "x-ves-oneof-field-http_protocol_type": "[\"auto_http_config\",\"http1_config\",\"http2_options\"]",
  "x-ves-oneof-field-lb_source_ip_persistence_choice": "[\"disable_lb_source_ip_persistence\",\"enable_lb_source_ip_persistence\"]",
  "x-ves-oneof-field-max_requests_per_connection_choice": "[\"max_requests_per_connection\",\"no_request_limit_per_connection\"]",
  "x-ves-oneof-field-outlier_detection_choice": "[\"disable_outlier_detection\",\"outlier_detection\"]",
  "x-ves-oneof-field-panic_threshold_type": "[\"no_panic_threshold\",\"panic_threshold\"]",
  "x-ves-oneof-field-proxy_protocol_choice": "[\"disable_proxy_protocol\",\"proxy_protocol_v1\",\"proxy_protocol_v2\"]",
  "x-ves-oneof-field-subset_choice": "[\"disable_subsets\",\"enable_subsets\"]"
}
```

<a id="canonical-7c8c2751be4eefe1ec4dfbbb422f54916ffa9f881455b0d1c20518173d4b1d43"></a>

## Direct properties — default_pool.advanced_options / 6b5ccddf6170 / 3

- [auto_http_config](data-sources--http_loadbalancer--reference--group-014.md#canonical-5ca52d632da7902844472758da74dd654a70ab0d6345a2d9b48001285f3d4fc6): complete subsection reference.

- [circuit_breaker](data-sources--http_loadbalancer--reference--group-014.md#canonical-2e2390a96f79e2a00734b4abe4400d98111ca8830fe98cd2f30ab3e4e155fe22): complete subsection reference.

<a id="canonical-1ef2f42025d08fc1f2378019a52a6912cae7cc30f60265e80a0630ace5156655"></a>

<a id="canonical-1c45bea520656e0ec55cb0f5a2a84aab9c81965e22f257d64ac7cbcfe8fdb863"></a>

## connection_timeout property — default_pool.advanced_options / 6b5ccddf6170 / 4

Type: `"number"`. Computed.

The timeout for new network connections to endpoints in the cluster. This is specified in
milliseconds. The default value is 2 seconds. Server applies default when omitted. Recommended:
\`2000\`.

Upstream description:

The timeout for new network connections to endpoints in the cluster. This is specified in
milliseconds. The default value is 2 seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1800000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "1800000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1800000"
  }
}
```

- [default_circuit_breaker](data-sources--http_loadbalancer--reference--group-015.md#canonical-5e8dcb1009b2ac8b463ae599d6c4ad6f44cfcf5319123f179023e6e372d6bc83): complete subsection reference.

- [disable_circuit_breaker](data-sources--http_loadbalancer--reference--group-015.md#canonical-5369d22b2f383f736398706d1dab0723c9ce1bc365c39df59263882d2109c493): complete subsection reference.

- [disable_lb_source_ip_persistence](data-sources--http_loadbalancer--reference--group-015.md#canonical-c8311f1076a3c891150774c73d1cc155d6b02aa5569075dbe87888088d9aa6db): complete subsection reference.

- [disable_outlier_detection](data-sources--http_loadbalancer--reference--group-015.md#canonical-7e1b53d3792b5137119e9bae1e0d70a7c949e29228523daff1dd68e445a13d1c): complete subsection reference.

- [disable_proxy_protocol](data-sources--http_loadbalancer--reference--group-015.md#canonical-4cdfee2073154ae2f880711b670a1a723ae3292bc5effa46459f1e554c76c26c): complete subsection reference.

- [disable_subsets](data-sources--http_loadbalancer--reference--group-015.md#canonical-39fb095beca65af1fd5b74a0538b5d96d45255975d62c2b096a5c0d8f482526a): complete subsection reference.

- [enable_lb_source_ip_persistence](data-sources--http_loadbalancer--reference--group-015.md#canonical-c89ce3e50e5d3bca2b0c257fc95ee303b081911127a64045ce78329a94954aae): complete subsection reference.

- [enable_subsets](data-sources--http_loadbalancer--reference--group-015.md#canonical-a42b09a9e56fe1e56d2e00a366cd6ed7cffd5b50c918287e6a810b2d86253c83): complete subsection reference.

- [http1_config](data-sources--http_loadbalancer--reference--group-015.md#canonical-868058d67c9ebb526444dd67146132159600476ec72b691f7a06ed687c4b3c4b): complete subsection reference.

- [http2_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-324e48dbb083a33bf32131888d62b13bb782e82de3fe5e1f8f8e84454ba903c0): complete subsection reference.

<a id="canonical-80596a30f61dcaa7839c137889f762ae08398906878738cfa41c985aa564cc75"></a>

<a id="canonical-435b15ea6e5f9342a9628b0b480ce8be7ebc338f9bd97cc03340f9d5e487c612"></a>

## http_idle_timeout property — default_pool.advanced_options / 6b5ccddf6170 / 5

Type: `"number"`. Computed.

The idle timeout for upstream connection pool connections. The idle timeout is defined as the period
in which there are no active requests. When the idle timeout is reached the connection will be
closed. Server applies default when omitted. Recommended: \`300000\`.

Upstream description:

The idle timeout for upstream connection pool connections. The idle timeout is defined as the period
in which there are no active requests. When the idle timeout is reached the connection will be
closed. Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive.
This is specified in milliseconds. The default value is 5 minutes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-1690a354f6d895de8c08435de11c0b08cb2adaf64e7db1d506a057084fa5a7f5"></a>

<a id="canonical-314679afb656e66a5bf8171dee78935b6d4597bd706469fd34f81076b0dc6d13"></a>

## max_requests_per_connection property — default_pool.advanced_options / 6b5ccddf6170 / 6

Type: `"number"`. Computed.

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests allowed
per connection to the origin server. Enter a value &gt;=1 to define the request limit per
connection.

Upstream description:

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests allowed
per connection to the origin server. Enter a value &gt;=1 to define the request limit per
connection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [no_panic_threshold](data-sources--http_loadbalancer--reference--group-015.md#canonical-261f8654cf935740e1982b5a54cbe968169ccc9bb826f97314724c097b01556b): complete subsection reference.

- [no_request_limit_per_connection](data-sources--http_loadbalancer--reference--group-015.md#canonical-791916a592038a377dd3ab4e636f8dd575739e3348f6dd937ccd095d0e248230): complete subsection reference.

- [outlier_detection](data-sources--http_loadbalancer--reference--group-015.md#canonical-2f40095c75e6ce28684437106211d9abf1384088ed8542adc2abebb8a7933358): complete subsection reference.

<a id="canonical-3cfc4694494ffe694b840cc79608ba80c133954a672cc36dacbbd94b6b49be80"></a>

<a id="canonical-60eec60551536b8d42719ac806fed3b6d19b4d7cb6c8eecaaacb121089887f75"></a>

## panic_threshold property — default_pool.advanced_options / 6b5ccddf6170 / 7

Type: `"number"`. Computed.

Exclusive with \[no\_panic\_threshold\] Configure a threshold (percentage of unhealthy endpoints)
below which all endpoints will be considered for load balancing ignoring its health status.

Upstream description:

Exclusive with \[no\_panic\_threshold\]

Configure a threshold (percentage of unhealthy endpoints) below which all endpoints will be
considered for load balancing ignoring its health status.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

- [proxy_protocol_v1](data-sources--http_loadbalancer--reference--group-015.md#canonical-78d8a16b1644f794bace42e17864703c7de6fd05e8552ea0c9d86d3d31bce0c1): complete subsection reference.

- [proxy_protocol_v2](data-sources--http_loadbalancer--reference--group-015.md#canonical-254f7e4e21689359b5440db0ac52ce34e7a08e0b840191efb5144931fb63bcfd): complete subsection reference.

<a id="canonical-0d9aac843f6220c6fd36ded272d104944397e6e6d94d3a5e11921b6f88d98d96"></a>

## Next pages — default_pool.advanced_options / 6b5ccddf6170 / 8

- [default_pool.advanced_options.auto_http_config](data-sources--http_loadbalancer--reference--group-014.md#canonical-5ca52d632da7902844472758da74dd654a70ab0d6345a2d9b48001285f3d4fc6)
- [default_pool.advanced_options.circuit_breaker](data-sources--http_loadbalancer--reference--group-014.md#canonical-2e2390a96f79e2a00734b4abe4400d98111ca8830fe98cd2f30ab3e4e155fe22)
- [default_pool.advanced_options.default_circuit_breaker](data-sources--http_loadbalancer--reference--group-015.md#canonical-5e8dcb1009b2ac8b463ae599d6c4ad6f44cfcf5319123f179023e6e372d6bc83)
- [default_pool.advanced_options.disable_circuit_breaker](data-sources--http_loadbalancer--reference--group-015.md#canonical-5369d22b2f383f736398706d1dab0723c9ce1bc365c39df59263882d2109c493)
- [default_pool.advanced_options.disable_lb_source_ip_persistence](data-sources--http_loadbalancer--reference--group-015.md#canonical-c8311f1076a3c891150774c73d1cc155d6b02aa5569075dbe87888088d9aa6db)
- [default_pool.advanced_options.disable_outlier_detection](data-sources--http_loadbalancer--reference--group-015.md#canonical-7e1b53d3792b5137119e9bae1e0d70a7c949e29228523daff1dd68e445a13d1c)
- [default_pool.advanced_options.disable_proxy_protocol](data-sources--http_loadbalancer--reference--group-015.md#canonical-4cdfee2073154ae2f880711b670a1a723ae3292bc5effa46459f1e554c76c26c)
- [default_pool.advanced_options.disable_subsets](data-sources--http_loadbalancer--reference--group-015.md#canonical-39fb095beca65af1fd5b74a0538b5d96d45255975d62c2b096a5c0d8f482526a)
- [default_pool.advanced_options.enable_lb_source_ip_persistence](data-sources--http_loadbalancer--reference--group-015.md#canonical-c89ce3e50e5d3bca2b0c257fc95ee303b081911127a64045ce78329a94954aae)
- [default_pool.advanced_options.enable_subsets](data-sources--http_loadbalancer--reference--group-015.md#canonical-a42b09a9e56fe1e56d2e00a366cd6ed7cffd5b50c918287e6a810b2d86253c83)
- [default_pool.advanced_options.http1_config](data-sources--http_loadbalancer--reference--group-015.md#canonical-868058d67c9ebb526444dd67146132159600476ec72b691f7a06ed687c4b3c4b)
- [default_pool.advanced_options.http2_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-324e48dbb083a33bf32131888d62b13bb782e82de3fe5e1f8f8e84454ba903c0)
- [default_pool.advanced_options.no_panic_threshold](data-sources--http_loadbalancer--reference--group-015.md#canonical-261f8654cf935740e1982b5a54cbe968169ccc9bb826f97314724c097b01556b)
- [default_pool.advanced_options.no_request_limit_per_connection](data-sources--http_loadbalancer--reference--group-015.md#canonical-791916a592038a377dd3ab4e636f8dd575739e3348f6dd937ccd095d0e248230)
- [default_pool.advanced_options.outlier_detection](data-sources--http_loadbalancer--reference--group-015.md#canonical-2f40095c75e6ce28684437106211d9abf1384088ed8542adc2abebb8a7933358)
- [default_pool.advanced_options.proxy_protocol_v1](data-sources--http_loadbalancer--reference--group-015.md#canonical-78d8a16b1644f794bace42e17864703c7de6fd05e8552ea0c9d86d3d31bce0c1)
- [default_pool.advanced_options.proxy_protocol_v2](data-sources--http_loadbalancer--reference--group-015.md#canonical-254f7e4e21689359b5440db0ac52ce34e7a08e0b840191efb5144931fb63bcfd)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5ca52d632da7902844472758da74dd654a70ab0d6345a2d9b48001285f3d4fc6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be4c3593860d8a38b7d2161aa2e66d3ce4cddef59a41e7528ac7ec815328c0f3"></a>

## default_pool.advanced_options.auto_http_config — default_pool.advanced_options.auto_http_config / 77fbf3fa2d33 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- default_pool.advanced_options.auto_http_config

<a id="canonical-b37ccacd373866906d501d60000735c8aa2829d6eec173cf717447f6068b02dc"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-b405b6020e66b96c0c766dc300c60673d45551ece0420a5e688b424d85be593f"></a>

## Direct properties — default_pool.advanced_options.auto_http_config / 77fbf3fa2d33 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-518bb8b430de4412b5c2bf772550a7b55ec4e3de11965506d14e916e32a97b84"></a>

## Next pages — default_pool.advanced_options.auto_http_config / 77fbf3fa2d33 / 4

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-2e2390a96f79e2a00734b4abe4400d98111ca8830fe98cd2f30ab3e4e155fe22"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
