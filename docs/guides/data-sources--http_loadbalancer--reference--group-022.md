---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-18ee52b281a23daeea37365076e30df02276dd1bc336f500da16b4b81af67fe3"></a>

## policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets — policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets / 23f48ad13e28 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [policy_based_challenge.rule_list.rules.spec.ip_matcher](data-sources--http_loadbalancer--reference--group-021.md#canonical-ca6c31e0e3e39f3e19be14e0bd99fff1219d5beca5dc8e55559dea263cb80880)
- policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets

<a id="canonical-fb8d9ac51d14e3a31a03f1f980adfd814e3ed23cab3e3efc6470e2baad714e90"></a>

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

<a id="canonical-79e00642cbeb7d0d58639544b3deb2bfacf63b4a3cd8dedd6c1941ad7b6342fe"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets / 23f48ad13e28 / 3

<a id="canonical-c39e929b1ee1eb6219d83145b626043b265d69cb877f93a429d25500832d706c"></a>

<a id="canonical-8596105c1918c6c36c800f38b7a925351eb9692dc3e1887287a6b204b9663f9f"></a>

## kind property — policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets / 23f48ad13e28 / 4

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

<a id="canonical-f6f5f865facd79380b54d7bd9b53622caba18ea65ccd33083c6390c2d9090c42"></a>

<a id="canonical-dc36b97d26e9f7e5afdd62d7f571f77b6dc34b3bf50c422f85b4f21d4374a241"></a>

## name property — policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets / 23f48ad13e28 / 5

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

<a id="canonical-2a2f53a1951d5ce301e808d6418cca1bc7a9c571688c2c2c842c04ddb44785a2"></a>

<a id="canonical-9c1aee4d9008f257bbce1b757d0c8e3a12c50d13cf9068d84618a094731d31d6"></a>

## namespace property — policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets / 23f48ad13e28 / 6

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

<a id="canonical-0cd82b0a83cb05bba8de8c594260c0f0ff540c74b3bf87a8f2e274a9baaf7999"></a>

<a id="canonical-69bcffc4893c0cf39293f70ada06142644e43f50fd4b25bae13fc97bf7cef2d9"></a>

## tenant property — policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets / 23f48ad13e28 / 7

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

<a id="canonical-893e5f6cb14efc95e95b9fb87fbbc598b37908b1f1d7bf706a50e0d1ad36a0d7"></a>

<a id="canonical-656ae07f1a4e78befaf292b0916032761625758b25e490755df57384295770ad"></a>

## uid property — policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets / 23f48ad13e28 / 8

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

<a id="canonical-e42c2f1fa5260b6aa5e25c019f91c12e003bd81d07e26c720ba00ce9a9f7a0c6"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets / 23f48ad13e28 / 9

- [policy_based_challenge.rule_list.rules.spec.ip_matcher](data-sources--http_loadbalancer--reference--group-021.md#canonical-ca6c31e0e3e39f3e19be14e0bd99fff1219d5beca5dc8e55559dea263cb80880)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-929811c07d5b3b04bed514c363d8cf6d54eb5be41f99307e66bcd8a56a22bb19"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2bb29435e80e19ad210683721b81f0a51276d6375a601b06d92dea364fbf5d15"></a>

## policy_based_challenge.rule_list.rules.spec.ip_prefix_list — policy_based_challenge.rule_list.rules.spec.ip_prefix_list / 260c9f6f559d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- policy_based_challenge.rule_list.rules.spec.ip_prefix_list

<a id="canonical-90e069d34a92b471d88ba3ed0a62b5bb459800825d042d8273962edb5bda6852"></a>

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

<a id="canonical-26cae72df9ce0a97acce58b539f32748c22147d99819552dab9d7b9ae1e6242c"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.ip_prefix_list / 260c9f6f559d / 3

<a id="canonical-83e50e5cd888ca256e40632dde87bc0eec1143f025bdc212930df393647a79f4"></a>

<a id="canonical-30b61adc0df3fab2a638ad12ef3dae2358b9e401213c6827e443d957ee292dce"></a>

## invert_match property — policy_based_challenge.rule_list.rules.spec.ip_prefix_list / 260c9f6f559d / 4

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

<a id="canonical-c1bde288725a379c83079383c54fa170de1dd049ec2f3449700eb333b1f63060"></a>

<a id="canonical-e82e0ac25084d4c3dcb05644080c9e1d7fa69af706b6b7eb1ddc42cfaa17cb57"></a>

## ip_prefixes property — policy_based_challenge.rule_list.rules.spec.ip_prefix_list / 260c9f6f559d / 5

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

<a id="canonical-053cbb536b47b1cc87fb769e45a10ad63da6a143dbc0727b320b6ca2e3ebbd2e"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.ip_prefix_list / 260c9f6f559d / 6

- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5c5a7d6aa6611cf9b6261f1644001a97fd2267bb79e06050fc298a96b453cc7f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-58abc73204e4b010972d4945704c3d2f3dad523e5f013898cafc27f55effddec"></a>

## policy_based_challenge.rule_list.rules.spec.path — policy_based_challenge.rule_list.rules.spec.path / 26e1a30d1dc8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- policy_based_challenge.rule_list.rules.spec.path

<a id="canonical-547f9426b7e484a84bb6900651af4944ca4d08be3e7845d95eb6b0e242ec6c92"></a>

Type: `"single"`. Computed.

Path matcher specifies multiple criteria for matching an HTTP path string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of path prefixes, a list of exact path values and a list of regular expressions.

Upstream description:

A path matcher specifies multiple criteria for matching an HTTP path string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of path prefixes, a list of exact path values and a list of regular expressions.

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

<a id="canonical-c2fa8f08b301d5156070fac0149840ac07cf6c50f57f10340bd491448df8c9db"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.path / 26e1a30d1dc8 / 3

<a id="canonical-b587110dddb1bc701c0b011a7543d685d42e3071821b403302917d9bb05eeb3a"></a>

<a id="canonical-7e2aa6fa1cb2fe2015427e7e1de20ee852f9faf8a6f2c72b64e657cf715f5599"></a>

## encoded_path_matcher property — policy_based_challenge.rule_list.rules.spec.path / 26e1a30d1dc8 / 4

Type: `"bool"`. Computed.

Match against the encoded, escaped path.

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

<a id="canonical-4c3c177f982c7ca62bbac0d49025d411dc3213c1a86259a6f9e6bcb3f3634020"></a>

<a id="canonical-daa13b40c53b49513f30a9078b26c9bcd471ce9688f7814fdcf52279e58b72d7"></a>

## exact_values property — policy_based_challenge.rule_list.rules.spec.path / 26e1a30d1dc8 / 5

Type: `["list", "string"]`. Computed.

List of exact path values to match the input HTTP path against.

Upstream description:

A list of exact path values to match the input HTTP path against.

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
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-8abbb0cbd883c68e1a2614493652e8b6d2d658d782f1e657bd049aee5ed2b104"></a>

<a id="canonical-1604b8c755fa577e9b6fe23fc8349fcccfc3f393540627f683918d20f3dccd49"></a>

## invert_matcher property — policy_based_challenge.rule_list.rules.spec.path / 26e1a30d1dc8 / 6

Type: `"bool"`. Computed.

Invert Path Matcher. Invert the match result.

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

<a id="canonical-1a4937f31a0df5e1d68cf3d0cad9c3a664bb2822d390adca6364a496eca46584"></a>

<a id="canonical-25dd4c119deaf8d2f85a0e11a72c8d1b6e78068a34eca07257096ee76b1f1e7c"></a>

## prefix_values property — policy_based_challenge.rule_list.rules.spec.path / 26e1a30d1dc8 / 7

Type: `["list", "string"]`. Computed.

List of path prefix values to match the input HTTP path against.

Upstream description:

A list of path prefix values to match the input HTTP path against.

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
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-73135abeb6fd8550271ab0265916434495ff0c79786102eee2e398c7a6589020"></a>

<a id="canonical-ddace94f102963a3979c89a9b32e142a18965e462a69ac64610930e7bb6d79f9"></a>

## regex_values property — policy_based_challenge.rule_list.rules.spec.path / 26e1a30d1dc8 / 8

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input HTTP path against.

Upstream description:

A list of regular expressions to match the input HTTP path against.

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

<a id="canonical-3a3f1766106686af88974a307b33c42a805969a8eb7eee20c9046b629ecbeb41"></a>

<a id="canonical-01a7c87aead8364d895cb37b647e200001827ab3c00fc89d0694c859b74facb6"></a>

## suffix_values property — policy_based_challenge.rule_list.rules.spec.path / 26e1a30d1dc8 / 9

Type: `["list", "string"]`. Computed.

List of path suffix values to match the input HTTP path against.

Upstream description:

A list of path suffix values to match the input HTTP path against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-959fcf1e29e99fcdde2fc322cf3997babd76ab6495e120478c235fefe1b29b52"></a>

<a id="canonical-847e6e5b7bc69a1f78d8d8460ffc1037120715f630a8a7641a3057fea02914ff"></a>

## transformers property — policy_based_challenge.rule_list.rules.spec.path / 26e1a30d1dc8 / 10

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

<a id="canonical-514ad51ac7bb819a36b3599b0ba2c155177858017696827c4406790461730f05"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.path / 26e1a30d1dc8 / 11

- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a6dd4dd1d3204263ff38e0262d2b6b68d0e4ca2bb4f2d3300cad7ac24228173e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a83fbbd63fc7ac5fea5831c11674f1ba709ce0aa106b586641027b237fe5cd47"></a>

## policy_based_challenge.rule_list.rules.spec.query_params — policy_based_challenge.rule_list.rules.spec.query_params / 3f765d1fbdd9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- policy_based_challenge.rule_list.rules.spec.query_params

<a id="canonical-66e8287a94f138f99fc832b39227e095240d460beb639707c75548af874037c3"></a>

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

<a id="canonical-953d9a69dbb88bb27c994722f12f6c93a72ec127a3129dbe5202fc1dc1dedf4f"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.query_params / 3f765d1fbdd9 / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-022.md#canonical-f1f1800c74e1f0cd7bebbf66876491da064eba1b6cc03f7f1df7853c80e0d7e9): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-022.md#canonical-915d8232f0c13ae1c2c9ea33e2953096d0023838f46c9dff30bacc40e8b6807e): complete subsection reference.

<a id="canonical-b048a0409dd2e0a3988c0a8b961440bd2757cb044a131c66e68bf3b97d46da9b"></a>

<a id="canonical-2197be962f38c4d5290f88b3834d4bcc08078ad6154e3e65cc312858c8b84a97"></a>

## invert_matcher property — policy_based_challenge.rule_list.rules.spec.query_params / 3f765d1fbdd9 / 4

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

- [item](data-sources--http_loadbalancer--reference--group-022.md#canonical-721acd1a86e9a8fd7ca1bb5d3591efab10d184dab53673e5b97cb96b934ec14b): complete subsection reference.

<a id="canonical-e2f5d10e6c14af7d451c9191f0485ce62f6c43a7d2a48d0ab703d0f953c0810a"></a>

<a id="canonical-2a91ca804e2e7023722b4ccb5f42e846d2aa1a6f29b327ae684019ebe4db623c"></a>

## key property — policy_based_challenge.rule_list.rules.spec.query_params / 3f765d1fbdd9 / 5

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

<a id="canonical-dbe290bbb990f971d26b5e8d5faedd175c8fb3a6c1ce6a293abbfb7394bf5a1c"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.query_params / 3f765d1fbdd9 / 6

- [policy_based_challenge.rule_list.rules.spec.query_params.check_not_present](data-sources--http_loadbalancer--reference--group-022.md#canonical-f1f1800c74e1f0cd7bebbf66876491da064eba1b6cc03f7f1df7853c80e0d7e9)
- [policy_based_challenge.rule_list.rules.spec.query_params.check_present](data-sources--http_loadbalancer--reference--group-022.md#canonical-915d8232f0c13ae1c2c9ea33e2953096d0023838f46c9dff30bacc40e8b6807e)
- [policy_based_challenge.rule_list.rules.spec.query_params.item](data-sources--http_loadbalancer--reference--group-022.md#canonical-721acd1a86e9a8fd7ca1bb5d3591efab10d184dab53673e5b97cb96b934ec14b)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f1f1800c74e1f0cd7bebbf66876491da064eba1b6cc03f7f1df7853c80e0d7e9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-14bd8c5b99dffb0b81cb3cca787967e542b0afd2ae52cd9aa192bd0c48a37a84"></a>

## policy_based_challenge.rule_list.rules.spec.query_params.check_not_present — policy_based_challenge.rule_list.rules.spec.query_params.check_not_present / e600952c7b63 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [policy_based_challenge.rule_list.rules.spec.query_params](data-sources--http_loadbalancer--reference--group-022.md#canonical-a6dd4dd1d3204263ff38e0262d2b6b68d0e4ca2bb4f2d3300cad7ac24228173e)
- policy_based_challenge.rule_list.rules.spec.query_params.check_not_present

<a id="canonical-2eb0a04bddc2ece02b08319d4e48a536abca6501a2bbfac6c92e6ce136229bc3"></a>

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

<a id="canonical-5534fc74cbe774cfffa1ff9c654e0146004916342e6a630954c3cd9c8e9e0fc4"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.query_params.check_not_present / e600952c7b63 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-18019381d7029d5c45cd47cc7f9432881ce87f441ea08282c6da71dfa3231aa6"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.query_params.check_not_present / e600952c7b63 / 4

- [policy_based_challenge.rule_list.rules.spec.query_params](data-sources--http_loadbalancer--reference--group-022.md#canonical-a6dd4dd1d3204263ff38e0262d2b6b68d0e4ca2bb4f2d3300cad7ac24228173e)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-915d8232f0c13ae1c2c9ea33e2953096d0023838f46c9dff30bacc40e8b6807e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-795631c2027621a95e4e5488510598f2e54cbdc5c188b1a1e2a8d699cbecfc9e"></a>

## policy_based_challenge.rule_list.rules.spec.query_params.check_present — policy_based_challenge.rule_list.rules.spec.query_params.check_present / 2280b9c0b662 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [policy_based_challenge.rule_list.rules.spec.query_params](data-sources--http_loadbalancer--reference--group-022.md#canonical-a6dd4dd1d3204263ff38e0262d2b6b68d0e4ca2bb4f2d3300cad7ac24228173e)
- policy_based_challenge.rule_list.rules.spec.query_params.check_present

<a id="canonical-3bbf03ab171baa4498cbf56f9f22836b045828639b4f4a8abbc1744b90cddbfb"></a>

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

<a id="canonical-cd0adc77aeae79c9f775a5e24bd5a67d20b4475ddafe15a4811c09bd55ed83f4"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.query_params.check_present / 2280b9c0b662 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0807698eba72c7bb15a71e5ae9c1dafbc6623bfb8fa8b28aec6b0999d102d0f1"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.query_params.check_present / 2280b9c0b662 / 4

- [policy_based_challenge.rule_list.rules.spec.query_params](data-sources--http_loadbalancer--reference--group-022.md#canonical-a6dd4dd1d3204263ff38e0262d2b6b68d0e4ca2bb4f2d3300cad7ac24228173e)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-721acd1a86e9a8fd7ca1bb5d3591efab10d184dab53673e5b97cb96b934ec14b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c9f0a4d55a1379a69fb0d968a1866bf6389ecd5deb6d649ea8ee49d71439a861"></a>

## policy_based_challenge.rule_list.rules.spec.query_params.item — policy_based_challenge.rule_list.rules.spec.query_params.item / 073cb5d194e2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [policy_based_challenge.rule_list.rules.spec.query_params](data-sources--http_loadbalancer--reference--group-022.md#canonical-a6dd4dd1d3204263ff38e0262d2b6b68d0e4ca2bb4f2d3300cad7ac24228173e)
- policy_based_challenge.rule_list.rules.spec.query_params.item

<a id="canonical-ae20db4522e2eafe08927566b28297bc38f4761ad32328e7bd807a56605bf955"></a>

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

<a id="canonical-4a5627e00100f8f4ff1a83c8a8a46dda1caca5bf3a2ba1461e1e95832dffc687"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.query_params.item / 073cb5d194e2 / 3

<a id="canonical-42c70884e9052c90356ff235df761e1da067bac2df7e47f55a15d20bf5b03823"></a>

<a id="canonical-356728d922b3af5bb8ee9074a4b9a2f778d6334a538c05aa24ce8ab76dc8b8c4"></a>

## exact_values property — policy_based_challenge.rule_list.rules.spec.query_params.item / 073cb5d194e2 / 4

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

<a id="canonical-c527b5fda961c1fd509c33151dc402fa0b38dca85664c0275aa44f54581d87e2"></a>

<a id="canonical-5e859a95a26da829404d8d5d16863c3cc48f06ef853dff3d591e464a46717e2c"></a>

## regex_values property — policy_based_challenge.rule_list.rules.spec.query_params.item / 073cb5d194e2 / 5

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

<a id="canonical-5bf472d8f6928e9db4a761b14d8942c49d6af63e9eca335a1d25c3c991b6332f"></a>

<a id="canonical-2d07d1d874c24722f1f3a73739dd2cb135bc3682fa2e920acd88b7868a914350"></a>

## transformers property — policy_based_challenge.rule_list.rules.spec.query_params.item / 073cb5d194e2 / 6

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

<a id="canonical-8e5844b97348fb3c28b6d49e99185c6747c3291626223c60619c9e08f53d2acf"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.query_params.item / 073cb5d194e2 / 7

- [policy_based_challenge.rule_list.rules.spec.query_params](data-sources--http_loadbalancer--reference--group-022.md#canonical-a6dd4dd1d3204263ff38e0262d2b6b68d0e4ca2bb4f2d3300cad7ac24228173e)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-0a5c9cf79b02cbcd2d25551795dc18bc688986f470da1226073d5aba1c58f6cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a085fa48b5b9c43ff9ef663acf9908909303cc372044f0197507c771e18d1f9d"></a>

## policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher — policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher / b2a6a46ab98f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-c4149e523f8bdd66090c4e5b778d361b1da1ea90d4138bf2a63d443a8cb14b66)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-9830ddff37f0fe8c0cfe1e9a3edf8f20e146a060f53c430aa123bd1b40ca9143)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher

<a id="canonical-8eee0b0dacadcb098499a621029da0b204db690ee26ac7bda111fa57febf50c1"></a>

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

<a id="canonical-2f557a8c20b6686a1bbd7e0e628b0ca6a5ad2dc4454eb1b00a6d961df8ae3a1c"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher / b2a6a46ab98f / 3

<a id="canonical-cb572e65378b3c35b25aa03e07725ee4ae3b9858472eb78735bb233371fcd9b1"></a>

<a id="canonical-cbf92dd9e52164a6c97ea51023895c519586a969d88c3db7ffb5c18e8d366d33"></a>

## classes property — policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher / b2a6a46ab98f / 4

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

<a id="canonical-6678376af7e7432c720062e70f11a6caeceb3c9cdad922ddf97c0b9f529c56d3"></a>

<a id="canonical-146ac60aaae93c026bfe42f2b2d585b545cfae65bb51c3506dd94f26afacbe45"></a>

## exact_values property — policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher / b2a6a46ab98f / 5

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

<a id="canonical-ed67533fe788d5c5182665dc85feea65627b02250730b76c7857ea634576fc33"></a>

<a id="canonical-c13049e32a9b7439a09ad6cc27f2e6745fbf83e96a1e5e193d036ef8b803a68d"></a>

## excluded_values property — policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher / b2a6a46ab98f / 6

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

<a id="canonical-c9493116afeda2985ff1deec26cbdb36711434053e7f86d878952cc10daa4d22"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher / b2a6a46ab98f / 7

- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-021.md#canonical-4f3b5ee7a93a2054e46a5363d9c888149fe1ee4a5f170210d1fa699d96006048)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d72d6cbe3fcf27ef049a25fc77b8ca27f89051a311fc8bc01e902a9e43113657"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-832d6a55cd97619c99e5511548b0debbd7c44257c53a9f87297b24e3b61ca6c3"></a>

## policy_based_challenge.temporary_user_blocking — policy_based_challenge.temporary_user_blocking / cb0f909bfdc0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- policy_based_challenge.temporary_user_blocking

<a id="canonical-1b1d96a0dea2a130937cec9fd90344adc2119286b578c7023322eac9dc8dc1e4"></a>

Type: `"single"`. Computed.

Specifies configuration for temporary user blocking resulting from user behavior analysis. When
Malicious User Mitigation is enabled from service policy rules, users' accessing the application
will be analyzed for malicious activity and the configured mitigation actions will be taken on..

Upstream description:

Specifies configuration for temporary user blocking resulting from user behavior analysis.

When Malicious User Mitigation is enabled from service policy rules, users' accessing the
application will be analyzed for malicious activity and the configured mitigation actions will be
taken on identified malicious users. These mitigation actions include setting up temporary blocking
on that user. This configuration specifies settings on how that blocking should be done by the
loadbalancer.

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

<a id="canonical-96653652dd3638aa1ee990aee4ae3d2acc00b9760380e6dded4fa3d4f8e54866"></a>

## Direct properties — policy_based_challenge.temporary_user_blocking / cb0f909bfdc0 / 3

<a id="canonical-82d1faf3a7fa83398c45393b151c1e7078b4baabd496a2e6fce118c1a2521711"></a>

<a id="canonical-4228b058e142cff16f389b0be620c35c35d039ffa57e9c8e929f01a151bda760"></a>

## custom_page property — policy_based_challenge.temporary_user_blocking / cb0f909bfdc0 / 4

Type: `"string"`. Computed.

Custom message is of type . Currently supported URL schemes is . For scheme, message needs to be
encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g.
'Blocked.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '&lt;p&gt;
Blocked..

Upstream description:

Custom message is of type \`uri\_ref\`. Currently supported URL schemes is \`string:///\`. For
\`string:///\` scheme, message needs to be encoded in Base64 format. You can specify this message as
base64 encoded plain text message e.g. "Blocked.." or it can be HTML paragraph or a body string
encoded as base64 string E.g. "&lt;p&gt; Blocked &lt;/p&gt;". Base64 encoded string for this HTML is
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

<a id="canonical-515603ea1de68581bd10976e51b76e70a0e5a5e07293d4ec7f423a0ff98c8e5b"></a>

## Next pages — policy_based_challenge.temporary_user_blocking / cb0f909bfdc0 / 5

- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-be4cbb05d0eb2a144c0092ef2c3da5c3e2fe5e7fb7bc9d7641e25f48fda71855"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0cba5719f716feedf1bfc64a84ecfefb2aa211fbcebf451d2752adcd4682cdba"></a>

## protected_cookies — protected_cookies / 16d873a8ff2d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- protected_cookies

<a id="canonical-883b6af6269ae35ffc6f6a16e3c3ea5982ed8e1509fdde40176c653a06f925b7"></a>

Type: `"list"`. Computed.

Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering
Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering
Protection, enabling a web app firewall (WAF) is a prerequisite.

Upstream description:

Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering
Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering
Protection, enabling a web app firewall (WAF) is a prerequisite. The configured mode of WAF
(monitoring or blocking) will be enforced on the request when cookie tampering is identified. Note:
We recommend enabling Secure and HttpOnly attributes along with cookie tampering protection.

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

<a id="canonical-b5721c5e404d9f30bd9793068f274d432e958581002c2f73239b0e98568bd921"></a>

## Direct properties — protected_cookies / 16d873a8ff2d / 3

- [add_httponly](data-sources--http_loadbalancer--reference--group-022.md#canonical-fedb85712921330ae74526d0e4bf7f3d74aa6b5b9cc7cbba801c22a45d5719c3): complete subsection reference.

- [add_secure](data-sources--http_loadbalancer--reference--group-022.md#canonical-d33b39c2420c64e12df40833d8f4b156cd641a530c2e971b5f7f36aa6224c03a): complete subsection reference.

- [disable_tampering_protection](data-sources--http_loadbalancer--reference--group-022.md#canonical-be55d04b595e6201845e49c904be7d06f4f3a3bd30c7496547bafb69eac69323): complete subsection reference.

- [enable_tampering_protection](data-sources--http_loadbalancer--reference--group-022.md#canonical-a916fc99824175f23e018c01a71af0ce5d7f1a5f1ed74f33af5ae1f1ce83b2e9): complete subsection reference.

- [ignore_httponly](data-sources--http_loadbalancer--reference--group-022.md#canonical-d15d1c2ecccaa3979d0c9e26a3524a38ed22f40046a4ee145ff75baa9445ed23): complete subsection reference.

- [ignore_max_age](data-sources--http_loadbalancer--reference--group-022.md#canonical-ca17c227e3ecd6d78b6c075da2233f3394fcfc836bab6d520e47795faf71736f): complete subsection reference.

- [ignore_samesite](data-sources--http_loadbalancer--reference--group-022.md#canonical-56ce70e8afe5cd415d9fc8151b4928d93fadf6b90a72f1d387621660c0a7463b): complete subsection reference.

- [ignore_secure](data-sources--http_loadbalancer--reference--group-022.md#canonical-8da38f45e702ca96b6fc8c96c6a9e0ba3b670aaf51a51b19ca12ee39af316325): complete subsection reference.

<a id="canonical-28544eab57c3d54264057f89d272a0b1dbdbdd6d5eaaa12149e9284341e999b4"></a>

<a id="canonical-3f81c637228c91ac962a7eedc695f0a0d448328df41b8b44ef448fa47d093fff"></a>

## max_age_value property — protected_cookies / 16d873a8ff2d / 4

Type: `"number"`. Computed.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Upstream description:

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 34560000,
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
    "ves.io.schema.rules.uint32.lte": "34560000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  }
}
```

<a id="canonical-3a62e6419328d0e53f34b0f6a71a0239982b53feb9dcfec882aa663c9a784511"></a>

<a id="canonical-7e49c64cde8249e1fbdb1c79fe15c064241d8cd54705c6350f371adb0719cf5b"></a>

## name property — protected_cookies / 16d873a8ff2d / 5

Type: `"string"`. Computed.

Cookie Name. Name of the Cookie.

Upstream description:

Name of the Cookie.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [samesite_lax](data-sources--http_loadbalancer--reference--group-022.md#canonical-5b9babe22fe5283c0295fec5ba0246673ad00369c771fb8aaec909524eaf20f0): complete subsection reference.

- [samesite_none](data-sources--http_loadbalancer--reference--group-022.md#canonical-98d0ce59fc2c84b60e550304ffc6b227adb9b7cc5353913bd73db0f6cc6acd28): complete subsection reference.

- [samesite_strict](data-sources--http_loadbalancer--reference--group-022.md#canonical-cb3cee66587da0ac2299c805432c671e454de7a4942cce052f3d54cafded9c6a): complete subsection reference.

<a id="canonical-49a4f9c804d73bc59b709fa5971daac18ae85a5f1745e1ffd4d0dd0671326a05"></a>

## Next pages — protected_cookies / 16d873a8ff2d / 6

- [protected_cookies.add_httponly](data-sources--http_loadbalancer--reference--group-022.md#canonical-fedb85712921330ae74526d0e4bf7f3d74aa6b5b9cc7cbba801c22a45d5719c3)
- [protected_cookies.add_secure](data-sources--http_loadbalancer--reference--group-022.md#canonical-d33b39c2420c64e12df40833d8f4b156cd641a530c2e971b5f7f36aa6224c03a)
- [protected_cookies.disable_tampering_protection](data-sources--http_loadbalancer--reference--group-022.md#canonical-be55d04b595e6201845e49c904be7d06f4f3a3bd30c7496547bafb69eac69323)
- [protected_cookies.enable_tampering_protection](data-sources--http_loadbalancer--reference--group-022.md#canonical-a916fc99824175f23e018c01a71af0ce5d7f1a5f1ed74f33af5ae1f1ce83b2e9)
- [protected_cookies.ignore_httponly](data-sources--http_loadbalancer--reference--group-022.md#canonical-d15d1c2ecccaa3979d0c9e26a3524a38ed22f40046a4ee145ff75baa9445ed23)
- [protected_cookies.ignore_max_age](data-sources--http_loadbalancer--reference--group-022.md#canonical-ca17c227e3ecd6d78b6c075da2233f3394fcfc836bab6d520e47795faf71736f)
- [protected_cookies.ignore_samesite](data-sources--http_loadbalancer--reference--group-022.md#canonical-56ce70e8afe5cd415d9fc8151b4928d93fadf6b90a72f1d387621660c0a7463b)
- [protected_cookies.ignore_secure](data-sources--http_loadbalancer--reference--group-022.md#canonical-8da38f45e702ca96b6fc8c96c6a9e0ba3b670aaf51a51b19ca12ee39af316325)
- [protected_cookies.samesite_lax](data-sources--http_loadbalancer--reference--group-022.md#canonical-5b9babe22fe5283c0295fec5ba0246673ad00369c771fb8aaec909524eaf20f0)
- [protected_cookies.samesite_none](data-sources--http_loadbalancer--reference--group-022.md#canonical-98d0ce59fc2c84b60e550304ffc6b227adb9b7cc5353913bd73db0f6cc6acd28)
- [protected_cookies.samesite_strict](data-sources--http_loadbalancer--reference--group-022.md#canonical-cb3cee66587da0ac2299c805432c671e454de7a4942cce052f3d54cafded9c6a)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-fedb85712921330ae74526d0e4bf7f3d74aa6b5b9cc7cbba801c22a45d5719c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-75d318cf83782d3de3ec550726503d9ac3bdaa922bdab3cb6ec822623755d106"></a>

## protected_cookies.add_httponly — protected_cookies.add_httponly / 23b69b8b1e7f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-be4cbb05d0eb2a144c0092ef2c3da5c3e2fe5e7fb7bc9d7641e25f48fda71855)
- protected_cookies.add_httponly

<a id="canonical-2c273236425166712d7df2f1e33c511087016560720917501d8ae104390f113c"></a>

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

<a id="canonical-5b5f0d1cf6b4cb58e4980343280c06f28b1be0cb0804f339ba8eb31b7d12b66f"></a>

## Direct properties — protected_cookies.add_httponly / 23b69b8b1e7f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f39cc2bd3c64331d2b5256493f709487902b16c3f0c9c40f8f551ce58e69e427"></a>

## Next pages — protected_cookies.add_httponly / 23b69b8b1e7f / 4

- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-be4cbb05d0eb2a144c0092ef2c3da5c3e2fe5e7fb7bc9d7641e25f48fda71855)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d33b39c2420c64e12df40833d8f4b156cd641a530c2e971b5f7f36aa6224c03a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc0a75bfa10a030809230ac1562a37fea57ca4d0baab44248dd1078aa3997652"></a>

## protected_cookies.add_secure — protected_cookies.add_secure / 85675a3ae592 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-be4cbb05d0eb2a144c0092ef2c3da5c3e2fe5e7fb7bc9d7641e25f48fda71855)
- protected_cookies.add_secure

<a id="canonical-9e6758e1838046a144e488027774290e12db655016cc81961261dc9ca07c6917"></a>

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

<a id="canonical-2bb901e5c9d2b4f719b727ef88eabaf6a8e0a976b3d16c72297a042a5a21ac1b"></a>

## Direct properties — protected_cookies.add_secure / 85675a3ae592 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-91ef90c28ef2b4d0283942afba47aceaa0c14e5e019d1a008e26713ea544118d"></a>

## Next pages — protected_cookies.add_secure / 85675a3ae592 / 4

- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-be4cbb05d0eb2a144c0092ef2c3da5c3e2fe5e7fb7bc9d7641e25f48fda71855)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-be55d04b595e6201845e49c904be7d06f4f3a3bd30c7496547bafb69eac69323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1568bcf86e87900a01bd77b0d353bd53c0bdafdf601cb05e3fdc351e6074f7b"></a>

## protected_cookies.disable_tampering_protection — protected_cookies.disable_tampering_protection / f28def3c38ac / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-be4cbb05d0eb2a144c0092ef2c3da5c3e2fe5e7fb7bc9d7641e25f48fda71855)
- protected_cookies.disable_tampering_protection

<a id="canonical-f472b01e540ad5f7679a7c78a31de5bbe782fbea9830a36b68b6d14477bb8265"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable tampering protection.

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

<a id="canonical-99f2423df277d228fe0758f67d4d40d91a45c183e19a4733f2d575d138b38c01"></a>

## Direct properties — protected_cookies.disable_tampering_protection / f28def3c38ac / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-76e6ae0d0cc612e81293544dcc98a646bc4b30bdcd4e0312e53d55617401e71d"></a>

## Next pages — protected_cookies.disable_tampering_protection / f28def3c38ac / 4

- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-be4cbb05d0eb2a144c0092ef2c3da5c3e2fe5e7fb7bc9d7641e25f48fda71855)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a916fc99824175f23e018c01a71af0ce5d7f1a5f1ed74f33af5ae1f1ce83b2e9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3de53b309251c88ede64c76e522a944a564baa74ce09e99b62501947175f60f5"></a>

## protected_cookies.enable_tampering_protection — protected_cookies.enable_tampering_protection / 6e823cf0a876 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-be4cbb05d0eb2a144c0092ef2c3da5c3e2fe5e7fb7bc9d7641e25f48fda71855)
- protected_cookies.enable_tampering_protection

<a id="canonical-3431a1ba505c42ffd4a124e697c540292b73e7bf383c8d8966a60e2c94906743"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable tampering protection.

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

<a id="canonical-74aed896b7ba10062f2e56859d73c2d1108f3fb39854602c98dfa451011edd07"></a>

## Direct properties — protected_cookies.enable_tampering_protection / 6e823cf0a876 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bd18035f8f3880c1776eba8869c059ae515488c506645118f1bfa4d8e1368d81"></a>

## Next pages — protected_cookies.enable_tampering_protection / 6e823cf0a876 / 4

- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-be4cbb05d0eb2a144c0092ef2c3da5c3e2fe5e7fb7bc9d7641e25f48fda71855)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d15d1c2ecccaa3979d0c9e26a3524a38ed22f40046a4ee145ff75baa9445ed23"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f3fa2c184a914bb175501d9a8099e34e353cb9ea3fb397871fc8bfb194fca82"></a>

## protected_cookies.ignore_httponly — protected_cookies.ignore_httponly / 06a05fa729d5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-be4cbb05d0eb2a144c0092ef2c3da5c3e2fe5e7fb7bc9d7641e25f48fda71855)
- protected_cookies.ignore_httponly

<a id="canonical-c23ac39ebd259800ce9711bc3b968bd2541412fe65bbfbbede3148ccc1e218cb"></a>

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

<a id="canonical-753076e9936ccea42d9628bbbfb12ccce8fb48fe2c4470588923cb1fc7ac6827"></a>

## Direct properties — protected_cookies.ignore_httponly / 06a05fa729d5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7eb109eba4d7c3e7f8b75d905d6f435bf150da0cbbe839aa2c90c74dad8fc061"></a>

## Next pages — protected_cookies.ignore_httponly / 06a05fa729d5 / 4

- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-be4cbb05d0eb2a144c0092ef2c3da5c3e2fe5e7fb7bc9d7641e25f48fda71855)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ca17c227e3ecd6d78b6c075da2233f3394fcfc836bab6d520e47795faf71736f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a667d37885886ff614e121caad67b831b9913c9802fd19c9fb3ddad6ec74c3c8"></a>

## protected_cookies.ignore_max_age — protected_cookies.ignore_max_age / 2672b11f764f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-be4cbb05d0eb2a144c0092ef2c3da5c3e2fe5e7fb7bc9d7641e25f48fda71855)
- protected_cookies.ignore_max_age

<a id="canonical-20f69ed48cb89032a6fd64d18ec8f72f1f8c65eb834505701867b13dddbe5f00"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore max age.

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

<a id="canonical-36ff8af5c21730c0625e763c259d17af71c66d283a083021d4a3cc44dd54e18e"></a>

## Direct properties — protected_cookies.ignore_max_age / 2672b11f764f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d5f1fd44c5c5d6c0c4bb9cbe9d09dadf506c1f7b69de6918a1be41f2e65c32d7"></a>

## Next pages — protected_cookies.ignore_max_age / 2672b11f764f / 4

- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-be4cbb05d0eb2a144c0092ef2c3da5c3e2fe5e7fb7bc9d7641e25f48fda71855)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-56ce70e8afe5cd415d9fc8151b4928d93fadf6b90a72f1d387621660c0a7463b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-42048102e582caa816e5cc197f6cc62e0ee6f627edfe0076136bbc435adb091d"></a>

## protected_cookies.ignore_samesite — protected_cookies.ignore_samesite / 50bfcc3f7294 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-be4cbb05d0eb2a144c0092ef2c3da5c3e2fe5e7fb7bc9d7641e25f48fda71855)
- protected_cookies.ignore_samesite

<a id="canonical-56ebba9d040d011693e74b6bd73f9bf5ddbbf719dd0e56dc689abfcd6151642e"></a>

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

<a id="canonical-40f625a3dab9336f7d0ebebcf6ca9fdd0bfb1952f6c4b5d13a3f14a9757a366d"></a>

## Direct properties — protected_cookies.ignore_samesite / 50bfcc3f7294 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c2c1498d26959882b8912859419918b46e636edf61d027ea2a4f19c041586876"></a>

## Next pages — protected_cookies.ignore_samesite / 50bfcc3f7294 / 4

- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-be4cbb05d0eb2a144c0092ef2c3da5c3e2fe5e7fb7bc9d7641e25f48fda71855)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-8da38f45e702ca96b6fc8c96c6a9e0ba3b670aaf51a51b19ca12ee39af316325"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d67a0fd10eb1bc75d018cd74e3358dbfaf2f801d4cf515541101249727a3a8c6"></a>

## protected_cookies.ignore_secure — protected_cookies.ignore_secure / d3d5e9a498c2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-be4cbb05d0eb2a144c0092ef2c3da5c3e2fe5e7fb7bc9d7641e25f48fda71855)
- protected_cookies.ignore_secure

<a id="canonical-4efcf23d28991ccf7978ed71bf0307b9b3445d6c44ab690144cb026d4dd329ce"></a>

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

<a id="canonical-0b637c83c18d543949d27df3ee3aa793702e1a35912c306976f0e23b4a1b8db2"></a>

## Direct properties — protected_cookies.ignore_secure / d3d5e9a498c2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-38c964ce5eb01dc08f8fb45367865843d2b463fe9a30606f47fcdd75240aae07"></a>

## Next pages — protected_cookies.ignore_secure / d3d5e9a498c2 / 4

- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-be4cbb05d0eb2a144c0092ef2c3da5c3e2fe5e7fb7bc9d7641e25f48fda71855)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5b9babe22fe5283c0295fec5ba0246673ad00369c771fb8aaec909524eaf20f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e51cf427a2a5b2dfb3b3aae85a9d3a96e30e8eb119130dc9bf41a52fbfd068ec"></a>

## protected_cookies.samesite_lax — protected_cookies.samesite_lax / a7da192a2ac6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-be4cbb05d0eb2a144c0092ef2c3da5c3e2fe5e7fb7bc9d7641e25f48fda71855)
- protected_cookies.samesite_lax

<a id="canonical-72204830e138132bdcd8a58aff4ef50d39511039cad48b4a540412e0c5ea362c"></a>

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

<a id="canonical-e2068e530a019f21b802ff49273d0d9e99b8847069d4b312e51527dc3a9731f1"></a>

## Direct properties — protected_cookies.samesite_lax / a7da192a2ac6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d47bbc8dc546e4222001acc3fafb2e660282b083910378b8c028126e5c5778ea"></a>

## Next pages — protected_cookies.samesite_lax / a7da192a2ac6 / 4

- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-be4cbb05d0eb2a144c0092ef2c3da5c3e2fe5e7fb7bc9d7641e25f48fda71855)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-98d0ce59fc2c84b60e550304ffc6b227adb9b7cc5353913bd73db0f6cc6acd28"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a117704d73a0579684951046ea86b4a864ffd16eabc93e93c2dcf02c1b696a84"></a>

## protected_cookies.samesite_none — protected_cookies.samesite_none / 7f93e57ab9e1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-be4cbb05d0eb2a144c0092ef2c3da5c3e2fe5e7fb7bc9d7641e25f48fda71855)
- protected_cookies.samesite_none

<a id="canonical-5341b6280f2fafe0891662ad44ea315cc26968fdeff2e5305f24fc0d83cc2448"></a>

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

<a id="canonical-2785aaacec750199313848a302994161dee20d0764c7646615c747bfc634d749"></a>

## Direct properties — protected_cookies.samesite_none / 7f93e57ab9e1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b6eaceb586dd65ae164feb5ddbb6e659a412e92271f9c1825a3528c219bfd1ce"></a>

## Next pages — protected_cookies.samesite_none / 7f93e57ab9e1 / 4

- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-be4cbb05d0eb2a144c0092ef2c3da5c3e2fe5e7fb7bc9d7641e25f48fda71855)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-cb3cee66587da0ac2299c805432c671e454de7a4942cce052f3d54cafded9c6a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-42ee24f618479bba4e26006b6fb053da471db1c7e6b174aad154e534f4d16371"></a>

## protected_cookies.samesite_strict — protected_cookies.samesite_strict / 2446c4aa09be / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-be4cbb05d0eb2a144c0092ef2c3da5c3e2fe5e7fb7bc9d7641e25f48fda71855)
- protected_cookies.samesite_strict

<a id="canonical-91db128a975a3eaf6a497c50b5b21e138087e50b1574888a6e6e21d9985da3bb"></a>

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

<a id="canonical-6561997ad9b9baeb634c3fb05769f19146eb1ee0a7c39cd36e4d6598fbb718a7"></a>

## Direct properties — protected_cookies.samesite_strict / 2446c4aa09be / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c38a4a95011120b7429619b2d3ff076b486901e89f98973697395b574092c91f"></a>

## Next pages — protected_cookies.samesite_strict / 2446c4aa09be / 4

- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-be4cbb05d0eb2a144c0092ef2c3da5c3e2fe5e7fb7bc9d7641e25f48fda71855)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e5bf189bc969780bd93b95faed2206bf49f5f1b3fc912f32403ed032f262b529"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d3de4b335a5a0a10130a3ac9c4fe6128d0fb6e92ae40a74865003cabc78463c4"></a>

## random — random / c6b33b7ea68b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- random

<a id="canonical-9b3e8bbea8a3216200bd8ea4bb178a27ed62696ccb56345712cde3bec3253770"></a>

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

<a id="canonical-da1e7b86a67150b532c51ddf59b2912bbaac8769c5555795d6034a69fa593c79"></a>

## Direct properties — random / c6b33b7ea68b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ac160965ab13058f71e14648a8ecb283f6eaa7c178ab136db904629783a0baa7"></a>

## Next pages — random / c6b33b7ea68b / 4

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a2773b65b59fb5112f862f175b75084921aa8340dfdf9a9677b28b502ef4e739"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-46750e8eb222255dc5c0394af3f63896ab45caaa7ff772e0dd2945f25168afc6"></a>

## rate_limit — rate_limit / 9b0ca36801f2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- rate_limit

<a id="canonical-155da133b1ab0a39650fe0d779ba2bed147e2b177752a6c36db1e3d419936057"></a>

Type: `"single"`. Computed.

Load-balancer-wide per-client rate limiting. The counter applies across every path; use
api\_rate\_limit rules when only selected paths such as /login should be limited.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ip_allowed_list_choice": "[\"custom_ip_allowed_list\",\"ip_allowed_list\",\"no_ip_allowed_list\"]",
  "x-ves-oneof-field-policy_choice": "[\"no_policies\",\"policies\"]"
}
```

<a id="canonical-1f1d05c0ecdc7ea5312409eb9ad1cb292565a1e73e0cd18f0f1d08fb6d7b6b0a"></a>

## Direct properties — rate_limit / 9b0ca36801f2 / 3

- [custom_ip_allowed_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-c61725cdafd1b910116fd6a60b06d9c91c1f62c2a083c0503e67032c7c1de03f): complete subsection reference.

- [ip_allowed_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-5b49cf28713c41dbaef12806007ece1d6bd2e9fe19b26322da4e1653bbcc3652): complete subsection reference.

- [no_ip_allowed_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-40067a12a34818e35109cf5a527bd34155e7309e93be224f74cc550d7965b774): complete subsection reference.

- [no_policies](data-sources--http_loadbalancer--reference--group-022.md#canonical-b59939c54b0d7bf93e315cabd4627b3f2de9c16054fbe9ef5855f10b5d026821): complete subsection reference.

- [policies](data-sources--http_loadbalancer--reference--group-022.md#canonical-bb632a59ca0c39bf59fc31b135c0788b25ec69be8001e06238d8804d90dc2786): complete subsection reference.

- [rate_limiter](data-sources--http_loadbalancer--reference--group-022.md#canonical-818bea6eff21a962afaeb3f76c87aeb4074b1b50cc1691f7c378afbd43ad3b1d): complete subsection reference.

<a id="canonical-e2ff3a5ccbf55ee16bac1f5e09b7768e80073d31a92abac7327251c20b25b704"></a>

## Next pages — rate_limit / 9b0ca36801f2 / 4

- [rate_limit.custom_ip_allowed_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-c61725cdafd1b910116fd6a60b06d9c91c1f62c2a083c0503e67032c7c1de03f)
- [rate_limit.ip_allowed_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-5b49cf28713c41dbaef12806007ece1d6bd2e9fe19b26322da4e1653bbcc3652)
- [rate_limit.no_ip_allowed_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-40067a12a34818e35109cf5a527bd34155e7309e93be224f74cc550d7965b774)
- [rate_limit.no_policies](data-sources--http_loadbalancer--reference--group-022.md#canonical-b59939c54b0d7bf93e315cabd4627b3f2de9c16054fbe9ef5855f10b5d026821)
- [rate_limit.policies](data-sources--http_loadbalancer--reference--group-022.md#canonical-bb632a59ca0c39bf59fc31b135c0788b25ec69be8001e06238d8804d90dc2786)
- [rate_limit.rate_limiter](data-sources--http_loadbalancer--reference--group-022.md#canonical-818bea6eff21a962afaeb3f76c87aeb4074b1b50cc1691f7c378afbd43ad3b1d)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c61725cdafd1b910116fd6a60b06d9c91c1f62c2a083c0503e67032c7c1de03f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ed7ec2573862d07925d6ac1a3180a58c583d1294e9f087ce287da06ab22da617"></a>

## rate_limit.custom_ip_allowed_list — rate_limit.custom_ip_allowed_list / 94a550ff33f0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [rate_limit](data-sources--http_loadbalancer--reference--group-022.md#canonical-a2773b65b59fb5112f862f175b75084921aa8340dfdf9a9677b28b502ef4e739)
- rate_limit.custom_ip_allowed_list

<a id="canonical-24b348a954c954bb1d6191d26a3235007a5cb064e8d0255473b44f2f111c67c2"></a>

Type: `"single"`. Computed.

IP Allowed list using existing ip\_prefix\_set objects.

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

<a id="canonical-6ef25cf6c364f62fe44e7f8a36700a3cbba308c261e9045b6614f101fe34d140"></a>

## Direct properties — rate_limit.custom_ip_allowed_list / 94a550ff33f0 / 3

- [rate_limiter_allowed_prefixes](data-sources--http_loadbalancer--reference--group-022.md#canonical-d597f34e415ec5d6d49b07c1be59953b678d6290e0364004d7fd677f9b7fefb1): complete subsection reference.

<a id="canonical-975defb27f3a03d5b7dfcf8ac7d9af5f3a0c1a5a24875b64594729460b98f307"></a>

## Next pages — rate_limit.custom_ip_allowed_list / 94a550ff33f0 / 4

- [rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes](data-sources--http_loadbalancer--reference--group-022.md#canonical-d597f34e415ec5d6d49b07c1be59953b678d6290e0364004d7fd677f9b7fefb1)
- [rate_limit](data-sources--http_loadbalancer--reference--group-022.md#canonical-a2773b65b59fb5112f862f175b75084921aa8340dfdf9a9677b28b502ef4e739)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d597f34e415ec5d6d49b07c1be59953b678d6290e0364004d7fd677f9b7fefb1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c8dc9e48138c6f4f6be57aeef3477c592f2f1be350a3a4560e67843bc193ccc2"></a>

## rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes — rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes / fcfad64c2998 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [rate_limit](data-sources--http_loadbalancer--reference--group-022.md#canonical-a2773b65b59fb5112f862f175b75084921aa8340dfdf9a9677b28b502ef4e739)
- [rate_limit.custom_ip_allowed_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-c61725cdafd1b910116fd6a60b06d9c91c1f62c2a083c0503e67032c7c1de03f)
- rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes

<a id="canonical-8a384bf2fa48cb144aea582dd95b5c84c297f33eb9709aa59457dedb969743bb"></a>

Type: `"list"`. Computed.

References to ip\_prefix\_set objects. Requests from source IP addresses that are covered by one of
the allowed IP Prefixes are not subjected to rate limiting.

Upstream description:

References to ip\_prefix\_set objects. Requests from source IP addresses that are covered by one of
the allowed IP Prefixes are not subjected to rate limiting.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-3f4cc60e4b769dd1f143ae9937425021bb0c90d5562b4f148ccebf7d9d132035"></a>

## Direct properties — rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes / fcfad64c2998 / 3

<a id="canonical-9bedad407ad82b7b52cf41ae7981c8d04b29329595945cfc798b69c8f02e196d"></a>

<a id="canonical-38837fded7c988ae0ef06c2ad5bde3aa7c188a307b9ff929b02bb3dd24674733"></a>

## name property — rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes / fcfad64c2998 / 4

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

<a id="canonical-c5c5b70b5c9050913a65a28569e81328c30a7f1a2505c60790fc4c1cdb5da902"></a>

<a id="canonical-3fe163449abe5ce045a4ef7ebe655f0c0b3ca6b3e0f7f7866d848ddac76672b7"></a>

## namespace property — rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes / fcfad64c2998 / 5

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

<a id="canonical-339bea2674a3c9415f500e0359e6870359abb937e591699acb137a39459277d6"></a>

<a id="canonical-20aab6c6e1c5c7f00f485b00e42fce72565d454872cbaebeb0610a3497df1e77"></a>

## tenant property — rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes / fcfad64c2998 / 6

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

<a id="canonical-5fdb41263cf0ad149b76b0b38ac1e16b217819372b77b114849ff10805787bd2"></a>

## Next pages — rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes / fcfad64c2998 / 7

- [rate_limit.custom_ip_allowed_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-c61725cdafd1b910116fd6a60b06d9c91c1f62c2a083c0503e67032c7c1de03f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5b49cf28713c41dbaef12806007ece1d6bd2e9fe19b26322da4e1653bbcc3652"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-97f3310ea0e3ed09c6330192ecc2de807556c4b5d175ac612059247dc6da64ed"></a>

## rate_limit.ip_allowed_list — rate_limit.ip_allowed_list / ce5709a84ee6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [rate_limit](data-sources--http_loadbalancer--reference--group-022.md#canonical-a2773b65b59fb5112f862f175b75084921aa8340dfdf9a9677b28b502ef4e739)
- rate_limit.ip_allowed_list

<a id="canonical-337d476ebab435a6f586130f1c0da5cbb8c28eb7337105c73c3a61f946db7646"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-4d4ccf63f51cbb82eb795fca43dc761c0b78f9d6f4dbf313d4d99ffbd8e07375"></a>

## Direct properties — rate_limit.ip_allowed_list / ce5709a84ee6 / 3

<a id="canonical-65adf7026f954b1c5ec8d0c61bdad268986ce175ef5337fb7ea6bbd51aea0b20"></a>

<a id="canonical-b04edeb67533ce0c8f4841d78fe6659636e4ae99dc9d38770e973a67fd050391"></a>

## prefixes property — rate_limit.ip_allowed_list / ce5709a84ee6 / 4

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-54f1d78f622b6645aa8038b9b21583f900b8d520b2e0c56f08774c27b0067f2b"></a>

## Next pages — rate_limit.ip_allowed_list / ce5709a84ee6 / 5

- [rate_limit](data-sources--http_loadbalancer--reference--group-022.md#canonical-a2773b65b59fb5112f862f175b75084921aa8340dfdf9a9677b28b502ef4e739)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-40067a12a34818e35109cf5a527bd34155e7309e93be224f74cc550d7965b774"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fce0dea43b8167039c12ccc7872039c105a9dc0358643ebf68ed4c2f2016f2ed"></a>

## rate_limit.no_ip_allowed_list — rate_limit.no_ip_allowed_list / 6fe3e5e5c22e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [rate_limit](data-sources--http_loadbalancer--reference--group-022.md#canonical-a2773b65b59fb5112f862f175b75084921aa8340dfdf9a9677b28b502ef4e739)
- rate_limit.no_ip_allowed_list

<a id="canonical-ec46ef9b242da16f3574832274085b95d6d2a2c7f5de4299cfdda9136465874e"></a>

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

<a id="canonical-8136b24b1082cbb019c3c243d113d0c02a304ef11a69fccd26a70f728b7263d2"></a>

## Direct properties — rate_limit.no_ip_allowed_list / 6fe3e5e5c22e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5eeedea6f4f0e18e0705c4f232127537c9ca4f427ee892d36fc1d2ab5a0801d7"></a>

## Next pages — rate_limit.no_ip_allowed_list / 6fe3e5e5c22e / 4

- [rate_limit](data-sources--http_loadbalancer--reference--group-022.md#canonical-a2773b65b59fb5112f862f175b75084921aa8340dfdf9a9677b28b502ef4e739)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-b59939c54b0d7bf93e315cabd4627b3f2de9c16054fbe9ef5855f10b5d026821"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f1934eff3c94a81fb184ed0b8a3f49c377632c6042de01b25f8dd870a0365d9"></a>

## rate_limit.no_policies — rate_limit.no_policies / 66a73b932610 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [rate_limit](data-sources--http_loadbalancer--reference--group-022.md#canonical-a2773b65b59fb5112f862f175b75084921aa8340dfdf9a9677b28b502ef4e739)
- rate_limit.no_policies

<a id="canonical-8583386457fbb2770ff14cfb6cc21e546524d7a4f9d2aa3d67c42544edca4b76"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no policies. Defaults to \`map\[\]\`. Server applies default when
omitted.

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

<a id="canonical-cde29a46025b9e7444b12c96991a072dc294434495b98ec35d3a31c5b811bf6c"></a>

## Direct properties — rate_limit.no_policies / 66a73b932610 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-23f25e1f299a0a7ae05e4979200137b0ac703ef16327bc0b652ccafaa488c9c0"></a>

## Next pages — rate_limit.no_policies / 66a73b932610 / 4

- [rate_limit](data-sources--http_loadbalancer--reference--group-022.md#canonical-a2773b65b59fb5112f862f175b75084921aa8340dfdf9a9677b28b502ef4e739)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-bb632a59ca0c39bf59fc31b135c0788b25ec69be8001e06238d8804d90dc2786"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2374cc3cd2f9ef066596be64ce2c390412ff1498d40e0160056915ab2e567d5d"></a>

## rate_limit.policies — rate_limit.policies / 18e24c956a0d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [rate_limit](data-sources--http_loadbalancer--reference--group-022.md#canonical-a2773b65b59fb5112f862f175b75084921aa8340dfdf9a9677b28b502ef4e739)
- rate_limit.policies

<a id="canonical-9ca76d58615a70d34b149417838efc08a0e47512b57d0209ab773f7d11d607ed"></a>

Type: `"single"`. Computed.

List of rate limiter policies to be applied.

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

<a id="canonical-66265de5b668de11a955267beda6de812deb1577d7d19f1b70e8d859f310fe4d"></a>

## Direct properties — rate_limit.policies / 18e24c956a0d / 3

- [policies](data-sources--http_loadbalancer--reference--group-022.md#canonical-471a1d54c1b70e511cabe9117528a0cfcf50d7e37403d296192dcfd999ab452d): complete subsection reference.

<a id="canonical-e5c1ab4490b02f6a545780af3c87ea26001c09525855a687efcd37b5dcd75a08"></a>

## Next pages — rate_limit.policies / 18e24c956a0d / 4

- [rate_limit.policies.policies](data-sources--http_loadbalancer--reference--group-022.md#canonical-471a1d54c1b70e511cabe9117528a0cfcf50d7e37403d296192dcfd999ab452d)
- [rate_limit](data-sources--http_loadbalancer--reference--group-022.md#canonical-a2773b65b59fb5112f862f175b75084921aa8340dfdf9a9677b28b502ef4e739)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-471a1d54c1b70e511cabe9117528a0cfcf50d7e37403d296192dcfd999ab452d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-88c3721e5aa085a191aa7bb8927f091f0ec5405f7debf138343ed67fc0605983"></a>

## rate_limit.policies.policies — rate_limit.policies.policies / 3a2cdd0f76e2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [rate_limit](data-sources--http_loadbalancer--reference--group-022.md#canonical-a2773b65b59fb5112f862f175b75084921aa8340dfdf9a9677b28b502ef4e739)
- [rate_limit.policies](data-sources--http_loadbalancer--reference--group-022.md#canonical-bb632a59ca0c39bf59fc31b135c0788b25ec69be8001e06238d8804d90dc2786)
- rate_limit.policies.policies

<a id="canonical-1ca40818bbc40f2acdfe9c00f6ec097718ef6d4a2d85d16c0eb9912e4a19d495"></a>

Type: `"list"`. Computed.

Rate Limiter Policies. Ordered list of rate limiter policies.

Upstream description:

Ordered list of rate limiter policies.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-d2e3b067d34399a13eb53f38bae0f55ad6951bb8b3ceda6c28274847d2db7d78"></a>

## Direct properties — rate_limit.policies.policies / 3a2cdd0f76e2 / 3

<a id="canonical-a636b0ac991942eccedd2b943570d6cb1e349a7747a2b9668e14bfa4f5c5e428"></a>

<a id="canonical-d3d7db25ac7ab0e1ebaa6fa379ae3677f9bd5404c26720a55d31af37baaad1a2"></a>

## name property — rate_limit.policies.policies / 3a2cdd0f76e2 / 4

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

<a id="canonical-867e939b247a9ff13ca52e658793182d4ec8b3726ace85b01f8571f671dde468"></a>

<a id="canonical-26d06f0bc6f5629a786ebd2db010bc525d0e95c8d01a724d514cfd3868811bef"></a>

## namespace property — rate_limit.policies.policies / 3a2cdd0f76e2 / 5

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

<a id="canonical-85cb36e6f394cec9cf75b47233e88e28ab81c9bccb2aad2475cd03c111f09ad5"></a>

<a id="canonical-045c6ae3009a87e11048a724c3ea083de30818077f98caa4128534cc3e81f5a9"></a>

## tenant property — rate_limit.policies.policies / 3a2cdd0f76e2 / 6

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

<a id="canonical-6879606311c51c176bf91507ba688661232d6b784591d594521a82b9e3eae787"></a>

## Next pages — rate_limit.policies.policies / 3a2cdd0f76e2 / 7

- [rate_limit.policies](data-sources--http_loadbalancer--reference--group-022.md#canonical-bb632a59ca0c39bf59fc31b135c0788b25ec69be8001e06238d8804d90dc2786)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-818bea6eff21a962afaeb3f76c87aeb4074b1b50cc1691f7c378afbd43ad3b1d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-de2f8dd5a5ce195f8fd4b49b23d36f3b7d432b35c357e5d8c093296162b8043f"></a>

## rate_limit.rate_limiter — rate_limit.rate_limiter / 0523d56398ac / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [rate_limit](data-sources--http_loadbalancer--reference--group-022.md#canonical-a2773b65b59fb5112f862f175b75084921aa8340dfdf9a9677b28b502ef4e739)
- rate_limit.rate_limiter

<a id="canonical-8cbddd46a02c1a23646b408a1ea669069e34b407ee2717f882974d1d5ee60188"></a>

Type: `"single"`. Computed.

Tuple consisting of a rate limit period unit and the total number of allowed requests for that
period.

Upstream description:

A tuple consisting of a rate limit period unit and the total number of allowed requests for that
period.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_choice": "[\"action_block\",\"disabled\"]",
  "x-ves-oneof-field-algorithm": "[\"leaky_bucket\",\"token_bucket\"]"
}
```

<a id="canonical-d3d6853b635bf94b3ff84be1f4d983b61c6b6df58278a22eaa5b7e83c8df752c"></a>

## Direct properties — rate_limit.rate_limiter / 0523d56398ac / 3

- [action_block](data-sources--http_loadbalancer--reference--group-022.md#canonical-6095b0c46f423572246f07c39ac0d44961ddff1eadaaa1edb588922b0bd73bf5): complete subsection reference.

<a id="canonical-ca2c3fb376fbacd047b5d9e754895518b47eede34e86adffcfd40ed1b781227c"></a>

<a id="canonical-e7c5fb5705b6cb1ab2a15c727ef8adc6ec8990e797f21ee7c18d7810ae6c2c2d"></a>

## burst_multiplier property — rate_limit.rate_limiter / 0523d56398ac / 4

Type: `"number"`. Computed.

The maximum burst of requests to accommodate, expressed as a multiple of the rate.

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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

- [disabled](data-sources--http_loadbalancer--reference--group-022.md#canonical-c0e06cce0483ebbb67262aacb249b8921537d20c0aa124a040a67408929147a4): complete subsection reference.

- [leaky_bucket](data-sources--http_loadbalancer--reference--group-022.md#canonical-d2cd54da778856df60c57142687de646031c2c5e4440e61ec2ee269ecc4aeedf): complete subsection reference.

<a id="canonical-84ecc379303d271537ef64bb5ab901c8c324b7bbaa8ef1eaeed708ad6d35a8c1"></a>

<a id="canonical-80b0a2ac4018994757c9f17797f0c762e0a3e0a360de8a0eba940d6d2b79c6de"></a>

## period_multiplier property — rate_limit.rate_limiter / 0523d56398ac / 5

Type: `"number"`. Computed.

Setting, combined with Per Period units, provides a duration. Server applies default when omitted.

Upstream description:

This setting, combined with Per Period units, provides a duration.

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
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0"
  }
}
```

- [token_bucket](data-sources--http_loadbalancer--reference--group-022.md#canonical-f3b6183890a3b6f1d6fa1ba0755d0904dd4894046f4e27cd4db20c0204d4d3ee): complete subsection reference.

<a id="canonical-31a07fc169be57600aa36ad2efdc8eba20171c86642e24b884bced882f96c254"></a>

<a id="canonical-bc220296c72393b3eaeaa0c758a27afc230300999883a6f1f08463a9991a7537"></a>

## total_number property — rate_limit.rate_limiter / 0523d56398ac / 6

Type: `"number"`. Computed.

The total number of allowed requests per rate-limiting period.

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
    "minimum": 1
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

<a id="canonical-963d1f3838e201b6201b330d811514c14ef510f12b9c74b7b1babbe8fdffad18"></a>

<a id="canonical-2746c05e8c9f7452bbd56a428cb35c36d90278bdcd24b435bc5df4d5bfb64b1f"></a>

## unit property — rate_limit.rate_limiter / 0523d56398ac / 7

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

<a id="canonical-f5e654c475d09b372421bb08e2d9d1357e4df488c87b2f05b329519bf1b0be73"></a>

## Next pages — rate_limit.rate_limiter / 0523d56398ac / 8

- [rate_limit.rate_limiter.action_block](data-sources--http_loadbalancer--reference--group-022.md#canonical-6095b0c46f423572246f07c39ac0d44961ddff1eadaaa1edb588922b0bd73bf5)
- [rate_limit.rate_limiter.disabled](data-sources--http_loadbalancer--reference--group-022.md#canonical-c0e06cce0483ebbb67262aacb249b8921537d20c0aa124a040a67408929147a4)
- [rate_limit.rate_limiter.leaky_bucket](data-sources--http_loadbalancer--reference--group-022.md#canonical-d2cd54da778856df60c57142687de646031c2c5e4440e61ec2ee269ecc4aeedf)
- [rate_limit.rate_limiter.token_bucket](data-sources--http_loadbalancer--reference--group-022.md#canonical-f3b6183890a3b6f1d6fa1ba0755d0904dd4894046f4e27cd4db20c0204d4d3ee)
- [rate_limit](data-sources--http_loadbalancer--reference--group-022.md#canonical-a2773b65b59fb5112f862f175b75084921aa8340dfdf9a9677b28b502ef4e739)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6095b0c46f423572246f07c39ac0d44961ddff1eadaaa1edb588922b0bd73bf5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-761fc1ffa436c8aa7b6662be047e3cd2f353aefe302a22e6802ed49adee6a7f1"></a>

## rate_limit.rate_limiter.action_block — rate_limit.rate_limiter.action_block / 18e70352900c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [rate_limit](data-sources--http_loadbalancer--reference--group-022.md#canonical-a2773b65b59fb5112f862f175b75084921aa8340dfdf9a9677b28b502ef4e739)
- [rate_limit.rate_limiter](data-sources--http_loadbalancer--reference--group-022.md#canonical-818bea6eff21a962afaeb3f76c87aeb4074b1b50cc1691f7c378afbd43ad3b1d)
- rate_limit.rate_limiter.action_block

<a id="canonical-716226abf1b2c1fda31bbcc3ad2951ecec11531c86a9567301006686bce29883"></a>

Type: `"single"`. Computed.

Action where a user is blocked from making further requests after exceeding rate limit threshold.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-block_duration_choice": "[\"hours\",\"minutes\",\"seconds\"]"
}
```

<a id="canonical-d57dd1529064821cda89e5456932026d1322b2917004e5aba9c3a5c76c2fdbe3"></a>

## Direct properties — rate_limit.rate_limiter.action_block / 18e70352900c / 3

- [hours](data-sources--http_loadbalancer--reference--group-022.md#canonical-ab712f5a7ec4c44f66859b70eb24a4094eaa376abe8fee879d132cb4f9ba3d87): complete subsection reference.

- [minutes](data-sources--http_loadbalancer--reference--group-022.md#canonical-23445a71231122705bdd0104aefbded5fd270f7e6b22dfb9652d5407f4f68388): complete subsection reference.

- [seconds](data-sources--http_loadbalancer--reference--group-022.md#canonical-e6319433f82d93d390387157a95a0168bcb8449a8f6c3a5417245bf3d740d73a): complete subsection reference.

<a id="canonical-8c3c16081a9a3808c49731acf37e4d6b6dad1803270829ec88cf93894ee3e68c"></a>

## Next pages — rate_limit.rate_limiter.action_block / 18e70352900c / 4

- [rate_limit.rate_limiter.action_block.hours](data-sources--http_loadbalancer--reference--group-022.md#canonical-ab712f5a7ec4c44f66859b70eb24a4094eaa376abe8fee879d132cb4f9ba3d87)
- [rate_limit.rate_limiter.action_block.minutes](data-sources--http_loadbalancer--reference--group-022.md#canonical-23445a71231122705bdd0104aefbded5fd270f7e6b22dfb9652d5407f4f68388)
- [rate_limit.rate_limiter.action_block.seconds](data-sources--http_loadbalancer--reference--group-022.md#canonical-e6319433f82d93d390387157a95a0168bcb8449a8f6c3a5417245bf3d740d73a)
- [rate_limit.rate_limiter](data-sources--http_loadbalancer--reference--group-022.md#canonical-818bea6eff21a962afaeb3f76c87aeb4074b1b50cc1691f7c378afbd43ad3b1d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ab712f5a7ec4c44f66859b70eb24a4094eaa376abe8fee879d132cb4f9ba3d87"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef6d63d86bfa2f260cc080bcf6f256571f7beaf29593ba751968953fda9fb5a5"></a>

## rate_limit.rate_limiter.action_block.hours — rate_limit.rate_limiter.action_block.hours / 9c789ab1b160 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [rate_limit](data-sources--http_loadbalancer--reference--group-022.md#canonical-a2773b65b59fb5112f862f175b75084921aa8340dfdf9a9677b28b502ef4e739)
- [rate_limit.rate_limiter](data-sources--http_loadbalancer--reference--group-022.md#canonical-818bea6eff21a962afaeb3f76c87aeb4074b1b50cc1691f7c378afbd43ad3b1d)
- [rate_limit.rate_limiter.action_block](data-sources--http_loadbalancer--reference--group-022.md#canonical-6095b0c46f423572246f07c39ac0d44961ddff1eadaaa1edb588922b0bd73bf5)
- rate_limit.rate_limiter.action_block.hours

<a id="canonical-403c27234932f0507b97a1b78861da4f9faa15cd842691892a78cffedf0cd0f9"></a>

Type: `"single"`. Computed.

Hours. Input Duration Hours.

Upstream description:

Input Duration Hours.

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

<a id="canonical-830e879a77bdbd9df534d9fb5c0cf77f3d92c6291ea4e011224ecd1d02f68d5e"></a>

## Direct properties — rate_limit.rate_limiter.action_block.hours / 9c789ab1b160 / 3

<a id="canonical-f0079a13336f474de2652143f563a30ffc9a9c885bce8960200739601ae85b59"></a>

<a id="canonical-5e40e5839913ab5f626c7ef5532b96c29dc67ea10414445b11fa8137791a2b0f"></a>

## duration property — rate_limit.rate_limiter.action_block.hours / 9c789ab1b160 / 4

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 48,
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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "48"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "48"
  }
}
```

<a id="canonical-bfd4ca989c8069ddab8f2c0de6c9753010506a6c848c1fae9e6c89eec02b0216"></a>

## Next pages — rate_limit.rate_limiter.action_block.hours / 9c789ab1b160 / 5

- [rate_limit.rate_limiter.action_block](data-sources--http_loadbalancer--reference--group-022.md#canonical-6095b0c46f423572246f07c39ac0d44961ddff1eadaaa1edb588922b0bd73bf5)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-23445a71231122705bdd0104aefbded5fd270f7e6b22dfb9652d5407f4f68388"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be7e3c3a306f396a916bc6c7f16b1e7f5203d8c6b17030cdf4f18db91b9064c2"></a>

## rate_limit.rate_limiter.action_block.minutes — rate_limit.rate_limiter.action_block.minutes / 2fc31d0bc3a5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [rate_limit](data-sources--http_loadbalancer--reference--group-022.md#canonical-a2773b65b59fb5112f862f175b75084921aa8340dfdf9a9677b28b502ef4e739)
- [rate_limit.rate_limiter](data-sources--http_loadbalancer--reference--group-022.md#canonical-818bea6eff21a962afaeb3f76c87aeb4074b1b50cc1691f7c378afbd43ad3b1d)
- [rate_limit.rate_limiter.action_block](data-sources--http_loadbalancer--reference--group-022.md#canonical-6095b0c46f423572246f07c39ac0d44961ddff1eadaaa1edb588922b0bd73bf5)
- rate_limit.rate_limiter.action_block.minutes

<a id="canonical-32d325f790314e388b9f8c73cc79cd3c88f8dcb7bcb3ef80dce38752a58f6b99"></a>

Type: `"single"`. Computed.

Minutes. Input Duration Minutes.

Upstream description:

Input Duration Minutes.

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

<a id="canonical-c7a43d2536efeef31b1d222b99add37a5948e986b8ce81bb97a91e17ba117dcd"></a>

## Direct properties — rate_limit.rate_limiter.action_block.minutes / 2fc31d0bc3a5 / 3

<a id="canonical-7ba084fd1ddf8ef95f434dbd3657b8d8c6cccafb19e1572219f02664431fad83"></a>

<a id="canonical-2c70f69a25b43b009043a1bcd6a00e16d02fdb8e407f1f0712d93afbfb84163a"></a>

## duration property — rate_limit.rate_limiter.action_block.minutes / 2fc31d0bc3a5 / 4

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60,
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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "60"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "60"
  }
}
```

<a id="canonical-d497e4ec43e89bf91d94f1b5cbf469e355941e92c5cf2fe7e763bafd479dc4c3"></a>

## Next pages — rate_limit.rate_limiter.action_block.minutes / 2fc31d0bc3a5 / 5

- [rate_limit.rate_limiter.action_block](data-sources--http_loadbalancer--reference--group-022.md#canonical-6095b0c46f423572246f07c39ac0d44961ddff1eadaaa1edb588922b0bd73bf5)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e6319433f82d93d390387157a95a0168bcb8449a8f6c3a5417245bf3d740d73a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c8d534177f42398b97b5827666d1962a5c9360cd57c702e17ebdab72c5554cec"></a>

## rate_limit.rate_limiter.action_block.seconds — rate_limit.rate_limiter.action_block.seconds / fc0d3e430b7d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [rate_limit](data-sources--http_loadbalancer--reference--group-022.md#canonical-a2773b65b59fb5112f862f175b75084921aa8340dfdf9a9677b28b502ef4e739)
- [rate_limit.rate_limiter](data-sources--http_loadbalancer--reference--group-022.md#canonical-818bea6eff21a962afaeb3f76c87aeb4074b1b50cc1691f7c378afbd43ad3b1d)
- [rate_limit.rate_limiter.action_block](data-sources--http_loadbalancer--reference--group-022.md#canonical-6095b0c46f423572246f07c39ac0d44961ddff1eadaaa1edb588922b0bd73bf5)
- rate_limit.rate_limiter.action_block.seconds

<a id="canonical-392c180e4f0ddc918db0f5b7dcbdf48a04fa19055ad07aee2d2ba85e72471aec"></a>

Type: `"single"`. Computed.

Seconds. Input Duration Seconds.

Upstream description:

Input Duration Seconds.

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

<a id="canonical-ff97eb3ab4f4cf21e7675a0a3df3a6edd85bb4e48218822e9c2e822584b30cf7"></a>

## Direct properties — rate_limit.rate_limiter.action_block.seconds / fc0d3e430b7d / 3

<a id="canonical-a2d678277a1de998fce6694f8df225d3204b1cc1a5aaaecbd49eb4e0aec81efe"></a>

<a id="canonical-82a510e99c6f9bf2b20c35208659a78b9b41dec99705e74f300a6bc319f9b33d"></a>

## duration property — rate_limit.rate_limiter.action_block.seconds / fc0d3e430b7d / 4

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300,
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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "300"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "300"
  }
}
```

<a id="canonical-f00dc0d679fc0f0978b2f102234a59d9660ba3b1509ad3800b4ae13172a1de85"></a>

## Next pages — rate_limit.rate_limiter.action_block.seconds / fc0d3e430b7d / 5

- [rate_limit.rate_limiter.action_block](data-sources--http_loadbalancer--reference--group-022.md#canonical-6095b0c46f423572246f07c39ac0d44961ddff1eadaaa1edb588922b0bd73bf5)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c0e06cce0483ebbb67262aacb249b8921537d20c0aa124a040a67408929147a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0ce0dedba11d8126919bab0f80677ff335f935326bbbdb01130b6b446699c4f1"></a>

## rate_limit.rate_limiter.disabled — rate_limit.rate_limiter.disabled / 142d9f36740f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [rate_limit](data-sources--http_loadbalancer--reference--group-022.md#canonical-a2773b65b59fb5112f862f175b75084921aa8340dfdf9a9677b28b502ef4e739)
- [rate_limit.rate_limiter](data-sources--http_loadbalancer--reference--group-022.md#canonical-818bea6eff21a962afaeb3f76c87aeb4074b1b50cc1691f7c378afbd43ad3b1d)
- rate_limit.rate_limiter.disabled

<a id="canonical-2192eee7e45b42a91c215e13dd4a923a2d9254384344ce2e6ebd5369cb0608b4"></a>

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

<a id="canonical-be47d4472e39a1fe99b2315a7f512f486ec05770dc189a831a114a4c1f58caef"></a>

## Direct properties — rate_limit.rate_limiter.disabled / 142d9f36740f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-600d982e377a8bc96f13d867bf898e840bc39e34c292f2ae39a157c741039d9d"></a>

## Next pages — rate_limit.rate_limiter.disabled / 142d9f36740f / 4

- [rate_limit.rate_limiter](data-sources--http_loadbalancer--reference--group-022.md#canonical-818bea6eff21a962afaeb3f76c87aeb4074b1b50cc1691f7c378afbd43ad3b1d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d2cd54da778856df60c57142687de646031c2c5e4440e61ec2ee269ecc4aeedf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2406f6d9bd6647c34c4d19adde9ce4810450dae2e99ccd1a811fc049cb6aa592"></a>

## rate_limit.rate_limiter.leaky_bucket — rate_limit.rate_limiter.leaky_bucket / a8e3446e678c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [rate_limit](data-sources--http_loadbalancer--reference--group-022.md#canonical-a2773b65b59fb5112f862f175b75084921aa8340dfdf9a9677b28b502ef4e739)
- [rate_limit.rate_limiter](data-sources--http_loadbalancer--reference--group-022.md#canonical-818bea6eff21a962afaeb3f76c87aeb4074b1b50cc1691f7c378afbd43ad3b1d)
- rate_limit.rate_limiter.leaky_bucket

<a id="canonical-143f2611fb30aeb7ed2d821b2df6a3010d1e13c3d728baebadbee76ca06ca039"></a>

Type: `["object", {}]`. Computed.

Leaky-Bucket is the default rate limiter algorithm for F5.

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

<a id="canonical-6cdb5a578eace23a26f94198689fcea6aa11c7d13026daaa484818b5f0e4e820"></a>

## Direct properties — rate_limit.rate_limiter.leaky_bucket / a8e3446e678c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-da4fdf3886fa1d9bdfb68737c373f32a68101ec56fc7748d453de322d50eff00"></a>

## Next pages — rate_limit.rate_limiter.leaky_bucket / a8e3446e678c / 4

- [rate_limit.rate_limiter](data-sources--http_loadbalancer--reference--group-022.md#canonical-818bea6eff21a962afaeb3f76c87aeb4074b1b50cc1691f7c378afbd43ad3b1d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f3b6183890a3b6f1d6fa1ba0755d0904dd4894046f4e27cd4db20c0204d4d3ee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a97625e4ddeb239e20a07afa302db9eea6d9a931778dd618e30a7d8670ad245"></a>

## rate_limit.rate_limiter.token_bucket — rate_limit.rate_limiter.token_bucket / 71922d5c9a5a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [rate_limit](data-sources--http_loadbalancer--reference--group-022.md#canonical-a2773b65b59fb5112f862f175b75084921aa8340dfdf9a9677b28b502ef4e739)
- [rate_limit.rate_limiter](data-sources--http_loadbalancer--reference--group-022.md#canonical-818bea6eff21a962afaeb3f76c87aeb4074b1b50cc1691f7c378afbd43ad3b1d)
- rate_limit.rate_limiter.token_bucket

<a id="canonical-c11dfb43f19f5651541fefbe3f3dd7dbeb6246082046256f7cd734d477fc1c1d"></a>

Type: `["object", {}]`. Computed.

Token-Bucket is a rate limiter algorithm that is stricter with enforcing limits.

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

<a id="canonical-eca6cf07e5f94ca8ff38ae557c608f06583b9776c5dbb7434890603e7c8a1639"></a>

## Direct properties — rate_limit.rate_limiter.token_bucket / 71922d5c9a5a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3f9c4e64fd6acdc7fb19e944648d5b662eb38af6f9c5b149484360c784fd3ec4"></a>

## Next pages — rate_limit.rate_limiter.token_bucket / 71922d5c9a5a / 4

- [rate_limit.rate_limiter](data-sources--http_loadbalancer--reference--group-022.md#canonical-818bea6eff21a962afaeb3f76c87aeb4074b1b50cc1691f7c378afbd43ad3b1d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-eb0f6502296a7d3ad32cbe5d7b650ad8904c4ad46d9151d0118a8a74b08e8f3b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-efba29c4b56945a9958b147d4a3680871e0af3ecb9b68c999296bbcae2e7c797"></a>

## ring_hash — ring_hash / 946cc344ccb8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- ring_hash

<a id="canonical-3be9344020d3d328f50af074648df746217617e5883ae4fc2786ed82f1e690aa"></a>

Type: `"single"`. Computed.

Hash Policy List. List of hash policy rules.

Upstream description:

List of hash policy rules.

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

<a id="canonical-7c9f3123216788c846859a069c7ae6ff137f659277570484352ba37622cd0bdd"></a>

## Direct properties — ring_hash / 946cc344ccb8 / 3

- [hash_policy](data-sources--http_loadbalancer--reference--group-022.md#canonical-d0ef41a439ca9bcdbf4a745a12118c19591b0a96b48709da4b3bc174fe1f65d8): complete subsection reference.

<a id="canonical-055d96680f35ae0c11e3b6c76cbe9c4251bdc1887e52a5b72b7ee307f5389f06"></a>

## Next pages — ring_hash / 946cc344ccb8 / 4

- [ring_hash.hash_policy](data-sources--http_loadbalancer--reference--group-022.md#canonical-d0ef41a439ca9bcdbf4a745a12118c19591b0a96b48709da4b3bc174fe1f65d8)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d0ef41a439ca9bcdbf4a745a12118c19591b0a96b48709da4b3bc174fe1f65d8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-17b8594c91a135ced11f6743f6bb8cdfd4b2c03ea0a84453b8f2927d0bdeecb5"></a>

## ring_hash.hash_policy — ring_hash.hash_policy / 284ac7a98d72 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [ring_hash](data-sources--http_loadbalancer--reference--group-022.md#canonical-eb0f6502296a7d3ad32cbe5d7b650ad8904c4ad46d9151d0118a8a74b08e8f3b)
- ring_hash.hash_policy

<a id="canonical-2344e1b9eb85bfc31cc7fec6c886bda60f14bf260f2c6dca970df1313bfcb828"></a>

Type: `"list"`. Computed.

Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated
individually and the combined result is used to route the request.

Upstream description:

Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated
individually and the combined result is used to route the request.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-d9583ca35e60adbafdbfebdc7cb96c4b29925cae70e4bb7061bdbb761b613fbf"></a>

## Direct properties — ring_hash.hash_policy / 284ac7a98d72 / 3

- [cookie](data-sources--http_loadbalancer--reference--group-022.md#canonical-db64a31608dfbe92824f3d5855b3daddea80badfb44a89232585032930a5a83e): complete subsection reference.

<a id="canonical-87da5b198fd719b33f0aa094d6941c28ff79b3189c20f82359b8139e86cbdfd1"></a>

<a id="canonical-ff6988d7b9c6c2d23151e5d8ae2d7ad9a81d48129ece399f3573c2976c12934e"></a>

## header_name property — ring_hash.hash_policy / 284ac7a98d72 / 4

Type: `"string"`. Computed.

Exclusive with \[cookie source\_ip\] The name or key of the request header that will be used to
obtain the hash key.

Upstream description:

Exclusive with \[cookie source\_ip\] The name or key of the request header that will be used to
obtain the hash key.

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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-9f2d4ce02586eb203837bf58407499a98978f683229270d394474aa24a720a8d"></a>

<a id="canonical-13ac6e31378abcd8bb10843bb38cc4ac7b2b464dc71d026467db072e39b313ca"></a>

## source_ip property — ring_hash.hash_policy / 284ac7a98d72 / 5

Type: `"bool"`. Computed.

Exclusive with \[cookie header\_name\] Hash based on source IP address.

Upstream description:

Exclusive with \[cookie header\_name\] Hash based on source IP address.

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

<a id="canonical-107fa3ec7d6d67d2d0a9a0a9e0e4bee8d42833b2e7cbf947921bc0f8b89ea00e"></a>

<a id="canonical-501c317bd0ad9a64d64c0ebc43b0e2c5de3965f9e82d9261c6944dc99cf07fb9"></a>

## terminal property — ring_hash.hash_policy / 284ac7a98d72 / 6

Type: `"bool"`. Computed.

Terminal. Specify if its a terminal policy.

Upstream description:

Specify if its a terminal policy.

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

<a id="canonical-32c1e072aabf7f4178db14a1a3a3f525b3792611ad6b603712be7aa4d011456f"></a>

## Next pages — ring_hash.hash_policy / 284ac7a98d72 / 7

- [ring_hash.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-022.md#canonical-db64a31608dfbe92824f3d5855b3daddea80badfb44a89232585032930a5a83e)
- [ring_hash](data-sources--http_loadbalancer--reference--group-022.md#canonical-eb0f6502296a7d3ad32cbe5d7b650ad8904c4ad46d9151d0118a8a74b08e8f3b)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-db64a31608dfbe92824f3d5855b3daddea80badfb44a89232585032930a5a83e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-661823566b8c3c030ab52b0e9981ff2719d423f008dc8e854feffc086d762d45"></a>

## ring_hash.hash_policy.cookie — ring_hash.hash_policy.cookie / 53d00f8fb47e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [ring_hash](data-sources--http_loadbalancer--reference--group-022.md#canonical-eb0f6502296a7d3ad32cbe5d7b650ad8904c4ad46d9151d0118a8a74b08e8f3b)
- [ring_hash.hash_policy](data-sources--http_loadbalancer--reference--group-022.md#canonical-d0ef41a439ca9bcdbf4a745a12118c19591b0a96b48709da4b3bc174fe1f65d8)
- ring_hash.hash_policy.cookie

<a id="canonical-f72e7d2c892033892354a00fdba920f7766374a4c8837c201fedeb7310a0f8b9"></a>

Type: `"single"`. Computed.

Two types of cookie affinity: 1. Passive. Takes a cookie that's present in the cookies header and
hashes on its value. 2. Generated. Generates and sets a cookie with an expiration (TTL) on the first
request from the client in its response to the client, based on the endpoint the request gets..

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

<a id="canonical-9be18c9ede0141d1100544bd497a573a4770e87c659ad3be059596e38c4a2a1d"></a>

## Direct properties — ring_hash.hash_policy.cookie / 53d00f8fb47e / 3

- [add_httponly](data-sources--http_loadbalancer--reference--group-022.md#canonical-86f4d9d30e7ba51957097a891073eb8fc8219878afbacfba56d2a96bc6256bab): complete subsection reference.

- [add_secure](data-sources--http_loadbalancer--reference--group-022.md#canonical-078eb63c3e84530387be8f5f2a3c12ec4b4a7a27206fd80da4652af740e1d0b1): complete subsection reference.

- [ignore_httponly](data-sources--http_loadbalancer--reference--group-022.md#canonical-1f2edda37e8e4d0983b8c768449e6fa8768a938cf5cdb669b00f2ecb460823ac): complete subsection reference.

- [ignore_samesite](data-sources--http_loadbalancer--reference--group-022.md#canonical-31423eb893f99e39b68e048e366bd6710aa168b5289a644f8229705bd777de05): complete subsection reference.

- [ignore_secure](data-sources--http_loadbalancer--reference--group-022.md#canonical-6f83252e6cc1cc2f82d4fe0970824f75d0bd8a196dcfa979984c2badfa7d4a24): complete subsection reference.

<a id="canonical-1eb3a99ac2c89465b0661d32fd50a558f3f6d6dd17b0287f53a0fc2178063cbd"></a>

<a id="canonical-e7438bc23719317de989cc37c720dec32e8c89728aad9a773f1cfb23d5447f66"></a>

## name property — ring_hash.hash_policy.cookie / 53d00f8fb47e / 4

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

<a id="canonical-5011227c1ed7f705d8015a81580b7dcc0398b866dab6a81fdb987d778232f3a3"></a>

<a id="canonical-34cd865dc25c086a4f5efcebc36fe3132876af5ab7eb5abec70061df467fc0bb"></a>

## path property — ring_hash.hash_policy.cookie / 53d00f8fb47e / 5

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

- [samesite_lax](data-sources--http_loadbalancer--reference--group-022.md#canonical-3984762df61fa585a118a636298aef1aba082acecea9a47d6a2f0198165ce845): complete subsection reference.

- [samesite_none](data-sources--http_loadbalancer--reference--group-022.md#canonical-8167a3c1dba873138ef5cd9750994048de2db06498b4b3e878e57d9fdb714816): complete subsection reference.

- [samesite_strict](data-sources--http_loadbalancer--reference--group-022.md#canonical-ce8b1f9f68ceb41b909efe956c4c655e47cc4d138dd1ffbe53a4f2b093f53222): complete subsection reference.

<a id="canonical-d0ccc1dbe0b55906b9a92e4d2f473b4b00eea2029974f3d2776d9c41b6bf6322"></a>

<a id="canonical-f798fe9a3a323f3f20196e3fab7174ca5f9af5c2dde0e4b3c445428bb876564a"></a>

## ttl property — ring_hash.hash_policy.cookie / 53d00f8fb47e / 6

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

<a id="canonical-94d8e4b4425697107cda653d8ba3d89a3926367d63b1c5d1fe1d2c569a3b0e2a"></a>

## Next pages — ring_hash.hash_policy.cookie / 53d00f8fb47e / 7

- [ring_hash.hash_policy.cookie.add_httponly](data-sources--http_loadbalancer--reference--group-022.md#canonical-86f4d9d30e7ba51957097a891073eb8fc8219878afbacfba56d2a96bc6256bab)
- [ring_hash.hash_policy.cookie.add_secure](data-sources--http_loadbalancer--reference--group-022.md#canonical-078eb63c3e84530387be8f5f2a3c12ec4b4a7a27206fd80da4652af740e1d0b1)
- [ring_hash.hash_policy.cookie.ignore_httponly](data-sources--http_loadbalancer--reference--group-022.md#canonical-1f2edda37e8e4d0983b8c768449e6fa8768a938cf5cdb669b00f2ecb460823ac)
- [ring_hash.hash_policy.cookie.ignore_samesite](data-sources--http_loadbalancer--reference--group-022.md#canonical-31423eb893f99e39b68e048e366bd6710aa168b5289a644f8229705bd777de05)
- [ring_hash.hash_policy.cookie.ignore_secure](data-sources--http_loadbalancer--reference--group-022.md#canonical-6f83252e6cc1cc2f82d4fe0970824f75d0bd8a196dcfa979984c2badfa7d4a24)
- [ring_hash.hash_policy.cookie.samesite_lax](data-sources--http_loadbalancer--reference--group-022.md#canonical-3984762df61fa585a118a636298aef1aba082acecea9a47d6a2f0198165ce845)
- [ring_hash.hash_policy.cookie.samesite_none](data-sources--http_loadbalancer--reference--group-022.md#canonical-8167a3c1dba873138ef5cd9750994048de2db06498b4b3e878e57d9fdb714816)
- [ring_hash.hash_policy.cookie.samesite_strict](data-sources--http_loadbalancer--reference--group-022.md#canonical-ce8b1f9f68ceb41b909efe956c4c655e47cc4d138dd1ffbe53a4f2b093f53222)
- [ring_hash.hash_policy](data-sources--http_loadbalancer--reference--group-022.md#canonical-d0ef41a439ca9bcdbf4a745a12118c19591b0a96b48709da4b3bc174fe1f65d8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-86f4d9d30e7ba51957097a891073eb8fc8219878afbacfba56d2a96bc6256bab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-872240258722e6bda525b046911a006b6cea64cd848288b40b39f7da473e2526"></a>

## ring_hash.hash_policy.cookie.add_httponly — ring_hash.hash_policy.cookie.add_httponly / 355214b7a57e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [ring_hash](data-sources--http_loadbalancer--reference--group-022.md#canonical-eb0f6502296a7d3ad32cbe5d7b650ad8904c4ad46d9151d0118a8a74b08e8f3b)
- [ring_hash.hash_policy](data-sources--http_loadbalancer--reference--group-022.md#canonical-d0ef41a439ca9bcdbf4a745a12118c19591b0a96b48709da4b3bc174fe1f65d8)
- [ring_hash.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-022.md#canonical-db64a31608dfbe92824f3d5855b3daddea80badfb44a89232585032930a5a83e)
- ring_hash.hash_policy.cookie.add_httponly

<a id="canonical-7270bce37ad3cf1075258283ecf4e311ca95c8f8eb0236bfd09dd573d2d041bc"></a>

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

<a id="canonical-d11573adad92c30b52db246b754454ff9f76ec1dbf2ed0a062dfb3cdf683d511"></a>

## Direct properties — ring_hash.hash_policy.cookie.add_httponly / 355214b7a57e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-33c0efbfa0b7c416bb8e66a6240d4b5b4966fd13fe1c6c0a73fe0f306cc8fe80"></a>

## Next pages — ring_hash.hash_policy.cookie.add_httponly / 355214b7a57e / 4

- [ring_hash.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-022.md#canonical-db64a31608dfbe92824f3d5855b3daddea80badfb44a89232585032930a5a83e)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-078eb63c3e84530387be8f5f2a3c12ec4b4a7a27206fd80da4652af740e1d0b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8aa9e04bb611f7bc7bbb0bf5327cc27121b8d6c274a01743ef030a0eec62b49a"></a>

## ring_hash.hash_policy.cookie.add_secure — ring_hash.hash_policy.cookie.add_secure / 5481880378c3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [ring_hash](data-sources--http_loadbalancer--reference--group-022.md#canonical-eb0f6502296a7d3ad32cbe5d7b650ad8904c4ad46d9151d0118a8a74b08e8f3b)
- [ring_hash.hash_policy](data-sources--http_loadbalancer--reference--group-022.md#canonical-d0ef41a439ca9bcdbf4a745a12118c19591b0a96b48709da4b3bc174fe1f65d8)
- [ring_hash.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-022.md#canonical-db64a31608dfbe92824f3d5855b3daddea80badfb44a89232585032930a5a83e)
- ring_hash.hash_policy.cookie.add_secure

<a id="canonical-faa3a8e7b5ce1a0e9d1f4a528a1daef8526af275d9214b260e615ccc548c4711"></a>

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

<a id="canonical-39cddd0644bacf8ad4536018ff58e0aa0093358bbc0d2439dc6bd87c68cd4045"></a>

## Direct properties — ring_hash.hash_policy.cookie.add_secure / 5481880378c3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-34cc936cd0352e1135f2b9e7f105d84b3a7dfe0fb2ff90615a18d939793b0d61"></a>

## Next pages — ring_hash.hash_policy.cookie.add_secure / 5481880378c3 / 4

- [ring_hash.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-022.md#canonical-db64a31608dfbe92824f3d5855b3daddea80badfb44a89232585032930a5a83e)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1f2edda37e8e4d0983b8c768449e6fa8768a938cf5cdb669b00f2ecb460823ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ed4f0e39aedd925c88dcdfb6c2cba10fbdf1bbe5ffed638aa65b01940629721"></a>

## ring_hash.hash_policy.cookie.ignore_httponly — ring_hash.hash_policy.cookie.ignore_httponly / 58dec5aa8a17 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [ring_hash](data-sources--http_loadbalancer--reference--group-022.md#canonical-eb0f6502296a7d3ad32cbe5d7b650ad8904c4ad46d9151d0118a8a74b08e8f3b)
- [ring_hash.hash_policy](data-sources--http_loadbalancer--reference--group-022.md#canonical-d0ef41a439ca9bcdbf4a745a12118c19591b0a96b48709da4b3bc174fe1f65d8)
- [ring_hash.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-022.md#canonical-db64a31608dfbe92824f3d5855b3daddea80badfb44a89232585032930a5a83e)
- ring_hash.hash_policy.cookie.ignore_httponly

<a id="canonical-b08288dc5f7bc2a38f15a00b740dfbc5b7ba33edda4bdfe12de572d6860f7cbc"></a>

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

<a id="canonical-e8f1053f4f8f591bdbfe1777103bb3140a15345b8bf87689437ad7f3180a7543"></a>

## Direct properties — ring_hash.hash_policy.cookie.ignore_httponly / 58dec5aa8a17 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f49e3ab0a6c8639bea44d65ddc7ad1321a5bea27a1d1cf9ded2e996f3e50d8c5"></a>

## Next pages — ring_hash.hash_policy.cookie.ignore_httponly / 58dec5aa8a17 / 4

- [ring_hash.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-022.md#canonical-db64a31608dfbe92824f3d5855b3daddea80badfb44a89232585032930a5a83e)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-31423eb893f99e39b68e048e366bd6710aa168b5289a644f8229705bd777de05"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc7438843b69449e59fafa8544c54d4f84a1987a796ded18013e6f9b6335a368"></a>

## ring_hash.hash_policy.cookie.ignore_samesite — ring_hash.hash_policy.cookie.ignore_samesite / 73b6ff99879e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [ring_hash](data-sources--http_loadbalancer--reference--group-022.md#canonical-eb0f6502296a7d3ad32cbe5d7b650ad8904c4ad46d9151d0118a8a74b08e8f3b)
- [ring_hash.hash_policy](data-sources--http_loadbalancer--reference--group-022.md#canonical-d0ef41a439ca9bcdbf4a745a12118c19591b0a96b48709da4b3bc174fe1f65d8)
- [ring_hash.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-022.md#canonical-db64a31608dfbe92824f3d5855b3daddea80badfb44a89232585032930a5a83e)
- ring_hash.hash_policy.cookie.ignore_samesite

<a id="canonical-ad8e3d5ab7c505e196a60020660554ca4def6d66217037a9c00920be5c7cfe37"></a>

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

<a id="canonical-060877090a3a10b783dec4203064132d33c47ff9f531fe9ebe91b511408aa9c3"></a>

## Direct properties — ring_hash.hash_policy.cookie.ignore_samesite / 73b6ff99879e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7ba63a2e92de790a9383993da4f72f665a2a7ad435f15ae7d37ac5869fe96d99"></a>

## Next pages — ring_hash.hash_policy.cookie.ignore_samesite / 73b6ff99879e / 4

- [ring_hash.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-022.md#canonical-db64a31608dfbe92824f3d5855b3daddea80badfb44a89232585032930a5a83e)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6f83252e6cc1cc2f82d4fe0970824f75d0bd8a196dcfa979984c2badfa7d4a24"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fe88b19b612de8711316d76f9211797a2cd21f4586b0c4f6c8d04e67a28d74ed"></a>

## ring_hash.hash_policy.cookie.ignore_secure — ring_hash.hash_policy.cookie.ignore_secure / aafedb3dade8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [ring_hash](data-sources--http_loadbalancer--reference--group-022.md#canonical-eb0f6502296a7d3ad32cbe5d7b650ad8904c4ad46d9151d0118a8a74b08e8f3b)
- [ring_hash.hash_policy](data-sources--http_loadbalancer--reference--group-022.md#canonical-d0ef41a439ca9bcdbf4a745a12118c19591b0a96b48709da4b3bc174fe1f65d8)
- [ring_hash.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-022.md#canonical-db64a31608dfbe92824f3d5855b3daddea80badfb44a89232585032930a5a83e)
- ring_hash.hash_policy.cookie.ignore_secure

<a id="canonical-04ac5ee6cf430deba326bba8eda3d17b4cf806aa391a6192b38eb0cae2fae46c"></a>

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

<a id="canonical-c5cfc4918b9eb2de6b478ccb6cdfaae39c79e28172ca9121f519635ceefee6f9"></a>

## Direct properties — ring_hash.hash_policy.cookie.ignore_secure / aafedb3dade8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ad76cb5975b7ba3dde952684ea3e4170bfab4b50957159ffb10834be17246e19"></a>

## Next pages — ring_hash.hash_policy.cookie.ignore_secure / aafedb3dade8 / 4

- [ring_hash.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-022.md#canonical-db64a31608dfbe92824f3d5855b3daddea80badfb44a89232585032930a5a83e)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-3984762df61fa585a118a636298aef1aba082acecea9a47d6a2f0198165ce845"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-731943586da2c2b22d2b796dedabc97d452f2007af3c10fc06ca1e0c7305a66e"></a>

## ring_hash.hash_policy.cookie.samesite_lax — ring_hash.hash_policy.cookie.samesite_lax / 16754c8d95b2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [ring_hash](data-sources--http_loadbalancer--reference--group-022.md#canonical-eb0f6502296a7d3ad32cbe5d7b650ad8904c4ad46d9151d0118a8a74b08e8f3b)
- [ring_hash.hash_policy](data-sources--http_loadbalancer--reference--group-022.md#canonical-d0ef41a439ca9bcdbf4a745a12118c19591b0a96b48709da4b3bc174fe1f65d8)
- [ring_hash.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-022.md#canonical-db64a31608dfbe92824f3d5855b3daddea80badfb44a89232585032930a5a83e)
- ring_hash.hash_policy.cookie.samesite_lax

<a id="canonical-ea9782edba66b50770cc611fbb02357b5805005383cea4d20e8607f35545f641"></a>

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

<a id="canonical-1f71fbd5dba6962e033b3b21fdd2a1cf4b1c21790f9a84f69cfda10b105d66fc"></a>

## Direct properties — ring_hash.hash_policy.cookie.samesite_lax / 16754c8d95b2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-49b101b501c9a4a73aa6761b221330503a547e38700ae5aece0b786740fce791"></a>

## Next pages — ring_hash.hash_policy.cookie.samesite_lax / 16754c8d95b2 / 4

- [ring_hash.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-022.md#canonical-db64a31608dfbe92824f3d5855b3daddea80badfb44a89232585032930a5a83e)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-8167a3c1dba873138ef5cd9750994048de2db06498b4b3e878e57d9fdb714816"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c8d132408a702ebfbe6df266046ee45dfc7c3375f7398af63232c8a90faed490"></a>

## ring_hash.hash_policy.cookie.samesite_none — ring_hash.hash_policy.cookie.samesite_none / 9f237886d8ac / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [ring_hash](data-sources--http_loadbalancer--reference--group-022.md#canonical-eb0f6502296a7d3ad32cbe5d7b650ad8904c4ad46d9151d0118a8a74b08e8f3b)
- [ring_hash.hash_policy](data-sources--http_loadbalancer--reference--group-022.md#canonical-d0ef41a439ca9bcdbf4a745a12118c19591b0a96b48709da4b3bc174fe1f65d8)
- [ring_hash.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-022.md#canonical-db64a31608dfbe92824f3d5855b3daddea80badfb44a89232585032930a5a83e)
- ring_hash.hash_policy.cookie.samesite_none

<a id="canonical-e3be03a248b229804f357899662b0f1bee0d803b2bb9e35431a9850c03c3132d"></a>

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

<a id="canonical-5a9d7342207174b9f78365597ba17cc3a0e398fba132738bc3cd2f39c2366156"></a>

## Direct properties — ring_hash.hash_policy.cookie.samesite_none / 9f237886d8ac / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-59310d5a5b86f4bdc169f2cfde5506631c1652dd910a477e092abf47bc66cfe1"></a>

## Next pages — ring_hash.hash_policy.cookie.samesite_none / 9f237886d8ac / 4

- [ring_hash.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-022.md#canonical-db64a31608dfbe92824f3d5855b3daddea80badfb44a89232585032930a5a83e)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ce8b1f9f68ceb41b909efe956c4c655e47cc4d138dd1ffbe53a4f2b093f53222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-692b3fbb3d6c1654d50177d973518d995e97d425fa3ceb692985a68cc41e8c14"></a>

## ring_hash.hash_policy.cookie.samesite_strict — ring_hash.hash_policy.cookie.samesite_strict / 309f2247247a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [ring_hash](data-sources--http_loadbalancer--reference--group-022.md#canonical-eb0f6502296a7d3ad32cbe5d7b650ad8904c4ad46d9151d0118a8a74b08e8f3b)
- [ring_hash.hash_policy](data-sources--http_loadbalancer--reference--group-022.md#canonical-d0ef41a439ca9bcdbf4a745a12118c19591b0a96b48709da4b3bc174fe1f65d8)
- [ring_hash.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-022.md#canonical-db64a31608dfbe92824f3d5855b3daddea80badfb44a89232585032930a5a83e)
- ring_hash.hash_policy.cookie.samesite_strict

<a id="canonical-14ce478908db898d8df50f78e2b27c7f1a8c497c9763b5fc9f4f103f98daae77"></a>

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

<a id="canonical-709c4283cd9706948e4bbcb28646429afd2c8d6c5200e4ce159f4a5bb742e237"></a>

## Direct properties — ring_hash.hash_policy.cookie.samesite_strict / 309f2247247a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-da094b3a45b2e1e935e572fa6b2911c98fdd6c4193037cf6402a26829d1fefd1"></a>

## Next pages — ring_hash.hash_policy.cookie.samesite_strict / 309f2247247a / 4

- [ring_hash.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-022.md#canonical-db64a31608dfbe92824f3d5855b3daddea80badfb44a89232585032930a5a83e)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9026a0dc1016de72d40836f2636cb220ee7076668244e68df27aabfdc22a32d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-011a34cc4aaac2f0c36093cec25f7f8d5ecb0d0dd89a1f79fc36df330f2e86b2"></a>

## round_robin — round_robin / 34c51669348f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- round_robin

<a id="canonical-17acf581cfd27d1efd14d30e5aa809447140203492f494ebda66ef71a5d0801f"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for round robin. Defaults to \`map\[\]\`. Server applies default when
omitted.

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

<a id="canonical-ac4863dcdda4e22bb1c9e1c16d5c4900df8f39353acbbfede762e1a848ea37ac"></a>

## Direct properties — round_robin / 34c51669348f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c2b9a23f94ab4d9b100f638c534c4a338c6aeaeb0895352fdbcbe819f109050c"></a>

## Next pages — round_robin / 34c51669348f / 4

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-15da56d30adc2d5dea210a7811a7b473f64ee52c8efcdc9b58c8a2dfc94487d0"></a>

## routes — routes / 7804ef03ed11 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- routes

<a id="canonical-7af16371430136d66c2ac8059db9acd34ec0fab762d54d55d68b7577a6cd01c3"></a>

Type: `"list"`. Computed.

Routes allow users to define match condition on a path and/or HTTP method to either forward matching
traffic to origin pool or redirect matching traffic to a different URL or respond directly to
matching traffic.

Upstream description:

Routes allow users to define match condition on a path and/or HTTP method to either forward matching
traffic to origin pool or redirect matching traffic to a different URL or respond directly to
matching traffic.

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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-18a3a0628d7681754b4c25b7bb886ae1a065424d78528af8c266382985dcb900"></a>

## Direct properties — routes / 7804ef03ed11 / 3

- [custom_route_object](data-sources--http_loadbalancer--reference--group-022.md#canonical-95f0db169032a06f0e7d96993e6ced2625e2ba369802d056161b8de0cad74f6c): complete subsection reference.

- [direct_response_route](data-sources--http_loadbalancer--reference--group-022.md#canonical-241694bc7b22e71a0e4b05da09ef2aaf3109f204b89ccd6b2d0d4b7eb4e2876e): complete subsection reference.

- [redirect_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-cd706f2aae1efbae824fd3e0a5070d8ff662967e7b86604b485aa5d3b39b0ffe): complete subsection reference.

- [route_state_disabled](data-sources--http_loadbalancer--reference--group-023.md#canonical-26a3458c018c5878797838910c1bbb3aac76cbcc3d77c81d2434c1334aa80760): complete subsection reference.

- [route_state_enabled](data-sources--http_loadbalancer--reference--group-023.md#canonical-c40cbd077ed664a4f6b634a5b335155e6e239298163d67a59031a515d5ba9eb3): complete subsection reference.

- [simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3): complete subsection reference.

<a id="canonical-3f8175dfd727b8b3d0687039585a2f682645744502de205e361e041b6c92820b"></a>

## Next pages — routes / 7804ef03ed11 / 4

- [routes.custom_route_object](data-sources--http_loadbalancer--reference--group-022.md#canonical-95f0db169032a06f0e7d96993e6ced2625e2ba369802d056161b8de0cad74f6c)
- [routes.direct_response_route](data-sources--http_loadbalancer--reference--group-022.md#canonical-241694bc7b22e71a0e4b05da09ef2aaf3109f204b89ccd6b2d0d4b7eb4e2876e)
- [routes.redirect_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-cd706f2aae1efbae824fd3e0a5070d8ff662967e7b86604b485aa5d3b39b0ffe)
- [routes.route_state_disabled](data-sources--http_loadbalancer--reference--group-023.md#canonical-26a3458c018c5878797838910c1bbb3aac76cbcc3d77c81d2434c1334aa80760)
- [routes.route_state_enabled](data-sources--http_loadbalancer--reference--group-023.md#canonical-c40cbd077ed664a4f6b634a5b335155e6e239298163d67a59031a515d5ba9eb3)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-95f0db169032a06f0e7d96993e6ced2625e2ba369802d056161b8de0cad74f6c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0929abcec1dea2581ca0da3cc2b285b564f8112ea714f8c476d70e6edf6d0d0d"></a>

## routes.custom_route_object — routes.custom_route_object / 0b62d2c08519 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- routes.custom_route_object

<a id="canonical-2dd32af66511207e7fead40890ef6bf4cdfc4530cfd0a061833e44d1dd06dbae"></a>

Type: `"single"`. Computed.

Custom route uses a route object created outside of this view.

Upstream description:

A custom route uses a route object created outside of this view.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-caching": "[\"caching_disable\",\"caching_inherit\"]"
}
```

<a id="canonical-ffb572ccdf8fc2223bf193f5cd67e0e9fd47755e9222fc9b51d977b5e5016514"></a>

## Direct properties — routes.custom_route_object / 0b62d2c08519 / 3

- [caching_disable](data-sources--http_loadbalancer--reference--group-022.md#canonical-6f3b2f1777a33c37aebf4534ac3e88abe326793c008e9c17fbebd15edb6fbe45): complete subsection reference.

- [caching_inherit](data-sources--http_loadbalancer--reference--group-022.md#canonical-a8dc16e2438ba4fa3c2338ae103a85c9d6499d3448be289038a81631901e0fa0): complete subsection reference.

- [route_ref](data-sources--http_loadbalancer--reference--group-022.md#canonical-15e5ba184d6dfef750ae41d25c551ee2aa8d4618f95dbf36021eb1dec7d1ad85): complete subsection reference.

<a id="canonical-1a752802e702f5da310ba6f85f05848a416ddf3d01d3bda06117cbd03f6b42be"></a>

## Next pages — routes.custom_route_object / 0b62d2c08519 / 4

- [routes.custom_route_object.caching_disable](data-sources--http_loadbalancer--reference--group-022.md#canonical-6f3b2f1777a33c37aebf4534ac3e88abe326793c008e9c17fbebd15edb6fbe45)
- [routes.custom_route_object.caching_inherit](data-sources--http_loadbalancer--reference--group-022.md#canonical-a8dc16e2438ba4fa3c2338ae103a85c9d6499d3448be289038a81631901e0fa0)
- [routes.custom_route_object.route_ref](data-sources--http_loadbalancer--reference--group-022.md#canonical-15e5ba184d6dfef750ae41d25c551ee2aa8d4618f95dbf36021eb1dec7d1ad85)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6f3b2f1777a33c37aebf4534ac3e88abe326793c008e9c17fbebd15edb6fbe45"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d136035c5847f65760da28bd342c7f1dc50c05796b94d386b377589b9dd2679f"></a>

## routes.custom_route_object.caching_disable — routes.custom_route_object.caching_disable / 0c158e30873c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.custom_route_object](data-sources--http_loadbalancer--reference--group-022.md#canonical-95f0db169032a06f0e7d96993e6ced2625e2ba369802d056161b8de0cad74f6c)
- routes.custom_route_object.caching_disable

<a id="canonical-8018be3b24d487fbd879586c39e5f91309b18687235482b6589f8fdd61dea499"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for caching disable.

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

<a id="canonical-f41dd947b7016cb3410d3bcbd3615938395bb018505297fe7a23ba3fd5a531cf"></a>

## Direct properties — routes.custom_route_object.caching_disable / 0c158e30873c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ad1659bca2b3af41d7c461d4e40407ad90423733a3ff657f3b38fc449fcbca07"></a>

## Next pages — routes.custom_route_object.caching_disable / 0c158e30873c / 4

- [routes.custom_route_object](data-sources--http_loadbalancer--reference--group-022.md#canonical-95f0db169032a06f0e7d96993e6ced2625e2ba369802d056161b8de0cad74f6c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a8dc16e2438ba4fa3c2338ae103a85c9d6499d3448be289038a81631901e0fa0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f0132a574084b0e55b97b213056d832c8e07bd50c7ab5ebcc5e72e0333e4eed"></a>

## routes.custom_route_object.caching_inherit — routes.custom_route_object.caching_inherit / 91c7513a270a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.custom_route_object](data-sources--http_loadbalancer--reference--group-022.md#canonical-95f0db169032a06f0e7d96993e6ced2625e2ba369802d056161b8de0cad74f6c)
- routes.custom_route_object.caching_inherit

<a id="canonical-069c99d48e973cef7732f0a32649ca7d155ea8a3dbc71e37cb48c58b8d2e09b2"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for caching inherit.

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

<a id="canonical-c3ded78c2f0270434612406bbc1b7647591701efb9a546e3ba048a17fa9b64c0"></a>

## Direct properties — routes.custom_route_object.caching_inherit / 91c7513a270a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-824f10482982c1cae1dca5d73a2114eb1371a47eeda6da7dafa1e83c07fa7e24"></a>

## Next pages — routes.custom_route_object.caching_inherit / 91c7513a270a / 4

- [routes.custom_route_object](data-sources--http_loadbalancer--reference--group-022.md#canonical-95f0db169032a06f0e7d96993e6ced2625e2ba369802d056161b8de0cad74f6c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-15e5ba184d6dfef750ae41d25c551ee2aa8d4618f95dbf36021eb1dec7d1ad85"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2be6d74539bab85b9948d73482ecd85254231f6ab9308346104c38dea9170cda"></a>

## routes.custom_route_object.route_ref — routes.custom_route_object.route_ref / 2ce96d3ed143 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.custom_route_object](data-sources--http_loadbalancer--reference--group-022.md#canonical-95f0db169032a06f0e7d96993e6ced2625e2ba369802d056161b8de0cad74f6c)
- routes.custom_route_object.route_ref

<a id="canonical-5968c8fc4f83086f7f9b9f87669bce070304d417ce46f93ced9be69ee102f5bb"></a>

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

<a id="canonical-0f2d453a5e038efc380b016b878e01eb456d0e30709790aca229c79295ac453d"></a>

## Direct properties — routes.custom_route_object.route_ref / 2ce96d3ed143 / 3

<a id="canonical-db6d670d2374118337858696395852d41d4f6a7c85baede5114f5da6b85e0206"></a>

<a id="canonical-b873addd0e2d427a9abc3f2dc19d1c821aaf266002b779cff786b861b94f928a"></a>

## name property — routes.custom_route_object.route_ref / 2ce96d3ed143 / 4

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

<a id="canonical-253c586592224f9a45441e18dd289692d6ef5897536afa85cf92ff03d66f9ad2"></a>

<a id="canonical-24f82480cde38e3176571c45965bbe0a5b34c010b207eb572e7aef3adfbb935c"></a>

## namespace property — routes.custom_route_object.route_ref / 2ce96d3ed143 / 5

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

<a id="canonical-c8fed863aabae6809734969cf8261091901d2d05823a4d8115fe961f6b6362a5"></a>

<a id="canonical-7200a379e1d30d51d88f204398d63abce089d7494fe052ad77bb4e156d82d433"></a>

## tenant property — routes.custom_route_object.route_ref / 2ce96d3ed143 / 6

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

<a id="canonical-c10e2f93d2d6da757f2a1b6156410cf7e861fd141a790e8b78a601224f85f07b"></a>

## Next pages — routes.custom_route_object.route_ref / 2ce96d3ed143 / 7

- [routes.custom_route_object](data-sources--http_loadbalancer--reference--group-022.md#canonical-95f0db169032a06f0e7d96993e6ced2625e2ba369802d056161b8de0cad74f6c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-241694bc7b22e71a0e4b05da09ef2aaf3109f204b89ccd6b2d0d4b7eb4e2876e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f9c9b8074b037de9e8f8fc4a0c8f3e4c5de0f94f67795010efc1618ba9945f4b"></a>

## routes.direct_response_route — routes.direct_response_route / 763d61eb7b5c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- routes.direct_response_route

<a id="canonical-79e40fd453d7d71f47b77c9e50c0d145d194f2085d4303beab21ca200a11ac0c"></a>

Type: `"single"`. Computed.

Direct response route matches on path, incoming header, incoming port and/or HTTP method and
responds directly to the matching traffic.

Upstream description:

A direct response route matches on path, incoming header, incoming port and/or HTTP method and
responds directly to the matching traffic.

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

<a id="canonical-a5baaf4f1d1764cf257f8c24240fd1ccba0eafe44b2bf53ab74d6044ff47dd36"></a>

## Direct properties — routes.direct_response_route / 763d61eb7b5c / 3

- [headers](data-sources--http_loadbalancer--reference--group-022.md#canonical-760ffd381d3c5da6fe8f5e0b740c813fcdae7161db3ed69eea4c122d1c9654cc): complete subsection reference.

<a id="canonical-7bf8016cd7ff6202ec44b9e90f095578b02dd2db9751db75e428560c016cc146"></a>

<a id="canonical-435d8b0ff9315b66b6215ff893a6738a4a5494b08a8be47fd1bc4d7beb34fc12"></a>

## http_method property — routes.direct_response_route / 763d61eb7b5c / 4

Type: `"string"`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [incoming_port](data-sources--http_loadbalancer--reference--group-022.md#canonical-5b994a44d4478081c4a6f627b3c96d54caa1db45d8d57058460cf3bb7daea7cf): complete subsection reference.

- [path](data-sources--http_loadbalancer--reference--group-023.md#canonical-21be880e7260de75735f117af48f72d37af63bb1cce4eeddc20d5c9f05e80180): complete subsection reference.

- [route_direct_response](data-sources--http_loadbalancer--reference--group-023.md#canonical-bdea64573060609049924cbdc60e279435e0ad3cfe1b6b6cfee5a37d53c29fc9): complete subsection reference.

<a id="canonical-166cd07c08ca086a51d72c0860974f09f2cddcf935900be34fa73795bb80f811"></a>

## Next pages — routes.direct_response_route / 763d61eb7b5c / 5

- [routes.direct_response_route.headers](data-sources--http_loadbalancer--reference--group-022.md#canonical-760ffd381d3c5da6fe8f5e0b740c813fcdae7161db3ed69eea4c122d1c9654cc)
- [routes.direct_response_route.incoming_port](data-sources--http_loadbalancer--reference--group-022.md#canonical-5b994a44d4478081c4a6f627b3c96d54caa1db45d8d57058460cf3bb7daea7cf)
- [routes.direct_response_route.path](data-sources--http_loadbalancer--reference--group-023.md#canonical-21be880e7260de75735f117af48f72d37af63bb1cce4eeddc20d5c9f05e80180)
- [routes.direct_response_route.route_direct_response](data-sources--http_loadbalancer--reference--group-023.md#canonical-bdea64573060609049924cbdc60e279435e0ad3cfe1b6b6cfee5a37d53c29fc9)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-760ffd381d3c5da6fe8f5e0b740c813fcdae7161db3ed69eea4c122d1c9654cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4bfc422e0a32cb69850f625eca7876731011743cfea35d8f330a0fdf7df3e662"></a>

## routes.direct_response_route.headers — routes.direct_response_route.headers / 44bc2f68c75d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.direct_response_route](data-sources--http_loadbalancer--reference--group-022.md#canonical-241694bc7b22e71a0e4b05da09ef2aaf3109f204b89ccd6b2d0d4b7eb4e2876e)
- routes.direct_response_route.headers

<a id="canonical-bff733cf79b5b9f440ceafc79c6d03ba7ca0e19fb010c42d7e2123d22d50c1d4"></a>

Type: `"list"`. Computed.

Headers. List of (key, value) headers.

Upstream description:

List of (key, value) headers.

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

<a id="canonical-5038eaa616d5db898d882b4c5a48fe179f96c3b1a2c3c4e588882fbb4fd75d04"></a>

## Direct properties — routes.direct_response_route.headers / 44bc2f68c75d / 3

<a id="canonical-eb5ebd5c74a313569811ef3e7801b51015c07bb680fa4461e024d52775516618"></a>

<a id="canonical-3b862dfc3403e81bd5e2a95e09bd80877895f6683962ae84227ef922109a9015"></a>

## exact property — routes.direct_response_route.headers / 44bc2f68c75d / 4

Type: `"string"`. Computed.

Exclusive with \[presence regex\] Header value to match exactly.

Upstream description:

Exclusive with \[presence regex\] Header value to match exactly.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-682073dbb07c14a5b868212b2784f41946ba63a9edb41371f20e529bafe7c497"></a>

<a id="canonical-f38be90f2ce57f1cf52d59e6d51d248cb895b16bb732a9b4f38048ef5a2cea66"></a>

## invert_match property — routes.direct_response_route.headers / 44bc2f68c75d / 5

Type: `"bool"`. Computed.

Invert the result of the match to detect missing header or non-matching value.

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

<a id="canonical-dda1e9432e94dd27da616d736d42a4d1d0c639ef82e191cf2bf8f6fc440b6ca5"></a>

<a id="canonical-4a316edad906c67a9a4ef88119749e059629f9761cf9afff45a98ccc1ceb02e4"></a>

## name property — routes.direct_response_route.headers / 44bc2f68c75d / 6

Type: `"string"`. Computed.

Name. Name of the header.

Upstream description:

Name of the header.

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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-cff809cdb9717434afa5a94b82d6b96f7cc34452b6c5eb46919285e1fa39a461"></a>

<a id="canonical-a0db32b077630cff3ca0b7ad6f08e54a134e205ff41d806d07348eb382f0081f"></a>

## presence property — routes.direct_response_route.headers / 44bc2f68c75d / 7

Type: `"bool"`. Computed.

Exclusive with \[exact regex\] If true, check for presence of header.

Upstream description:

Exclusive with \[exact regex\] If true, check for presence of header.

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

<a id="canonical-4a572169dbe354d3c94b2c67dfcc03fdaaf1c629a5a0ead5aff4e203ff44ebe0"></a>

<a id="canonical-02cb22bc7f114bc5e3419c7d13637fad79189c2651c1ea38335b92a86fafe4b3"></a>

## regex property — routes.direct_response_route.headers / 44bc2f68c75d / 8

Type: `"string"`. Computed.

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

Upstream description:

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

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
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-fa815357052bf0a85921bf564e6a02d1ac341464d2fe34104430029ef8844d1a"></a>

## Next pages — routes.direct_response_route.headers / 44bc2f68c75d / 9

- [routes.direct_response_route](data-sources--http_loadbalancer--reference--group-022.md#canonical-241694bc7b22e71a0e4b05da09ef2aaf3109f204b89ccd6b2d0d4b7eb4e2876e)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5b994a44d4478081c4a6f627b3c96d54caa1db45d8d57058460cf3bb7daea7cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-39693178444c201a464b50aff8e3c635b0dee493fcce453ffddd3e6aa024a455"></a>

## routes.direct_response_route.incoming_port — routes.direct_response_route.incoming_port / 4dc940d7bce2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.direct_response_route](data-sources--http_loadbalancer--reference--group-022.md#canonical-241694bc7b22e71a0e4b05da09ef2aaf3109f204b89ccd6b2d0d4b7eb4e2876e)
- routes.direct_response_route.incoming_port

<a id="canonical-60910ae27449661bf6012b086b0d3a1a54380ab611604cec3c51c6982abf85dc"></a>

Type: `"single"`. Computed.

Port match of the request can be a range or a specific port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

<a id="canonical-3ef4d490614286de31f79e2026ac5514c7b04668e17fa78ca296a847309d5c06"></a>

## Direct properties — routes.direct_response_route.incoming_port / 4dc940d7bce2 / 3

- [no_port_match](data-sources--http_loadbalancer--reference--group-023.md#canonical-3867f520a0380b2671da101048d3cf9916f3a6325803c683088abdec7a0de256): complete subsection reference.

<a id="canonical-7d5461ca6fad08363fab356a20239a5c1b3c103490d817c5032be307dac588a1"></a>
