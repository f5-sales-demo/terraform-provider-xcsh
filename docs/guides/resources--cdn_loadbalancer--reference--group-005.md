---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-a919b34c5e2653a9a68c1836131d64c2ac035e96c246f78e8535a2e22a6b4885"></a>

## api_rate_limit.server_url_rules.client_matcher.any_ip — api_rate_limit.server_url_rules.client_matcher.any_ip / 62dd308ae0f0 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-b57675087227f85e31d2f1e413d8dfa0895ab0596b75681660d5755b97a89275)
- api_rate_limit.server_url_rules.client_matcher.any_ip

<a id="canonical-10ee3f7a8fad8da92493e4e1c3ccb85f803a4c4a161ca5204c490f9eb5e7c66f"></a>

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

<a id="canonical-7628c45dc530f63b5b7bba4b221e88ca2e670726efbc32d41d3a3ebcb7ccdbef"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.any_ip / 62dd308ae0f0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0afb802597a47cfc554b3adbe82f544248f8b1729c814c4d7126d31afe57e91b"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.any_ip / 62dd308ae0f0 / 4

- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-b57675087227f85e31d2f1e413d8dfa0895ab0596b75681660d5755b97a89275)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-c99d6513f352a985f593897b8f86f593471932f73c0fe005d016c01e19ba8c2b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da5a75997c2d65ec9c35acc9620efe2f79fa67724ca68f1ecd01bdcb158cd40f"></a>

## api_rate_limit.server_url_rules.client_matcher.asn_list — api_rate_limit.server_url_rules.client_matcher.asn_list / f2a091ad8592 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-b57675087227f85e31d2f1e413d8dfa0895ab0596b75681660d5755b97a89275)
- api_rate_limit.server_url_rules.client_matcher.asn_list

<a id="canonical-c3d00729e58af74deeaedb51953f757f94bd8b29a0cc129a009734bf0aae6301"></a>

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

<a id="canonical-367ca14498179d11dfa7d70d57335867c0a6c4d5cacd4c44a099fc9a3e5a7a13"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.asn_list / f2a091ad8592 / 3

<a id="canonical-e7402ddb366731593e9cbcaa035e1b7ee49af02dd916d32470e03af94a071aa6"></a>

<a id="canonical-9b85a7a0ef81823987deb2437621b4d03739b2a932ace753debe0b6e62c49be2"></a>

## as_numbers property — api_rate_limit.server_url_rules.client_matcher.asn_list / f2a091ad8592 / 4

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

<a id="canonical-dafd17b9a4902d142f70b2e87e6507bb23c46c16d10c54b50d631668431bb580"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.asn_list / f2a091ad8592 / 5

- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-b57675087227f85e31d2f1e413d8dfa0895ab0596b75681660d5755b97a89275)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-db9b2bfe02531f7d46d5a91231b1c2bde4624d19ffe377c35dd9d0a14b023cfe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e1d1628e8b4df5c1a7755f6f4f6fe601bea24503406fa52d1ec5319864863bb1"></a>

## api_rate_limit.server_url_rules.client_matcher.asn_matcher — api_rate_limit.server_url_rules.client_matcher.asn_matcher / 79c5d47379aa / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-b57675087227f85e31d2f1e413d8dfa0895ab0596b75681660d5755b97a89275)
- api_rate_limit.server_url_rules.client_matcher.asn_matcher

<a id="canonical-b9affadaa7e112eb10512cd7331220d6d04662381063abb8ea0642c0ab9b9b90"></a>

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

<a id="canonical-221cd369f690a268f4b5297c5915efcccb260cd98da881d58eb1500265e252db"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.asn_matcher / 79c5d47379aa / 3

- [asn_sets](resources--cdn_loadbalancer--reference--group-005.md#canonical-0815f7019f4fcabbe053ba3cd7f8a242510d4f1a76037906bafe07e3caf9bad5): complete subsection reference.

<a id="canonical-29ea92cfffcee60150a1643dd112d66755518c4d6283f7afd1c01d4e88abdc9f"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.asn_matcher / 79c5d47379aa / 4

- [api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets](resources--cdn_loadbalancer--reference--group-005.md#canonical-0815f7019f4fcabbe053ba3cd7f8a242510d4f1a76037906bafe07e3caf9bad5)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-b57675087227f85e31d2f1e413d8dfa0895ab0596b75681660d5755b97a89275)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-0815f7019f4fcabbe053ba3cd7f8a242510d4f1a76037906bafe07e3caf9bad5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d5e4eac43fdcc18fa83958f6973f42ff95c5a32b4f9b062eec479d1fc3c2b870"></a>

## api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets — api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets / 111d272be117 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-b57675087227f85e31d2f1e413d8dfa0895ab0596b75681660d5755b97a89275)
- [api_rate_limit.server_url_rules.client_matcher.asn_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-db9b2bfe02531f7d46d5a91231b1c2bde4624d19ffe377c35dd9d0a14b023cfe)
- api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-8590f38e6b2a0bc4a51b7e7d3560e7d99adcbb3bc6b4dc3f03202d157c16f551"></a>

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

<a id="canonical-ad044c5089c5207b16763dea212cbfde3188d052421e4911b761ea73e9cbcdaf"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets / 111d272be117 / 3

<a id="canonical-540f9aae52dd67615cd6bb84f9e450cb0572e3cfb8702715a88c13ae5119c12e"></a>

<a id="canonical-90e5a79a13e1bf5b7e32276828c809dcfa3a3e2b5e81acd11f622ab1a397713c"></a>

## kind property — api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets / 111d272be117 / 4

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

<a id="canonical-4b95c2d50f747d672c42e20df8faa14535f663ada0aca0b6b6cc9f9589e97503"></a>

<a id="canonical-6014eb6e246baccac10c9da0ddcf1105a2d9c24c317c1109b58ee4931e3ae253"></a>

## name property — api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets / 111d272be117 / 5

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

<a id="canonical-53d4e9ab44e9106a981bb6b2380a1bcb51d7f3eb2a036423ee932526828f238d"></a>

<a id="canonical-1fc2f5a91cdca2ed4215451d49d50fe85d41c6e32e6d89b521ffa8def6abc7b6"></a>

## namespace property — api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets / 111d272be117 / 6

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

<a id="canonical-c49faa0783e6957e08f2a32cc73b4fced1a4b312f010e30a8f27fcfca5e4b048"></a>

<a id="canonical-c8198d80d32479f8d1e7f6a4678fb43bdcc679a532b70aa2465d014a3cdda3d1"></a>

## tenant property — api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets / 111d272be117 / 7

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

<a id="canonical-930daf7bfd40e66078e61fea3cadb7b350c209db0e88887d74d6cbc98d2cac3e"></a>

<a id="canonical-43c5bedba1c558a84e5d3a66031be0949bc1063f2c15231431746c0fe6ff2f6c"></a>

## uid property — api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets / 111d272be117 / 8

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

<a id="canonical-ddbb8da728397064507a9494e5f377f7b1a7a14a3e7177b69b7494759d42bbf3"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets / 111d272be117 / 9

- [api_rate_limit.server_url_rules.client_matcher.asn_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-db9b2bfe02531f7d46d5a91231b1c2bde4624d19ffe377c35dd9d0a14b023cfe)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-0535fca52232cdfb0faf772dc41efee44bad609f21ce23324b156358dff54a8b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1bb8b550bf35de8fd698e5ffe26c8dfee08e00b05a4225c4c6f4cb34cdbccf21"></a>

## api_rate_limit.server_url_rules.client_matcher.client_selector — api_rate_limit.server_url_rules.client_matcher.client_selector / c99ffa27b432 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-b57675087227f85e31d2f1e413d8dfa0895ab0596b75681660d5755b97a89275)
- api_rate_limit.server_url_rules.client_matcher.client_selector

<a id="canonical-bc897005766445bfaf7d00bd002d7684d9966cd13249630d3d452b18a3ac09d4"></a>

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

<a id="canonical-f6e877e68bcaf05e5ff351ad035496d4fbe6e6af3eaa3b09f04b1ded8bf51ef7"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.client_selector / c99ffa27b432 / 3

<a id="canonical-654852342658b8519a6f08d134c6a258a72f78f62fbc2789bc54c5a35e80b51e"></a>

<a id="canonical-b57b0b55b4f75e93a9e68df988f8ffa8f3658cbf687a6adb12c3f40cc682ef07"></a>

## expressions property — api_rate_limit.server_url_rules.client_matcher.client_selector / c99ffa27b432 / 4

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

<a id="canonical-d5a261626fb5e809fb807eb426ff6198b692a4af067d3069f1a794ed49a21164"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.client_selector / c99ffa27b432 / 5

- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-b57675087227f85e31d2f1e413d8dfa0895ab0596b75681660d5755b97a89275)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-9c55db87686e1d62df013b59ba5fff3b2253597dd72994c171adeea78bcf40ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb93da5f51f6b90d59a08b0950347c65d599e8439f84dc964cac4886b1541109"></a>

## api_rate_limit.server_url_rules.client_matcher.ip_matcher — api_rate_limit.server_url_rules.client_matcher.ip_matcher / 16f5fff8e214 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-b57675087227f85e31d2f1e413d8dfa0895ab0596b75681660d5755b97a89275)
- api_rate_limit.server_url_rules.client_matcher.ip_matcher

<a id="canonical-5f6b8044374de811ce6f6ba67ff7962ac230d041ba807107df3f6ac1d1326e51"></a>

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

<a id="canonical-3b845676d504dbb2ba53eaf2a7cf8f35c0b353ecb8ab9a31312f54c9563d5ddd"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.ip_matcher / 16f5fff8e214 / 3

<a id="canonical-0a85ed29f76e43fe8f2972f796c218f6a4e9f31644d0c69014d47bbe11fe1403"></a>

<a id="canonical-4cbf8f596bdf462d98ab405bbf84488c39b5102b628b7e86bf662ee7ba53b377"></a>

## invert_matcher property — api_rate_limit.server_url_rules.client_matcher.ip_matcher / 16f5fff8e214 / 4

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

- [prefix_sets](resources--cdn_loadbalancer--reference--group-005.md#canonical-f70a802a472f29233d15247e79a05b57616eaa0321e76cdac07e88bd29ffa15b): complete subsection reference.

<a id="canonical-cfee398a9416d3e92fd429e08a1d1bc79ef6a7be7fb07b3c43a71607803f3155"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.ip_matcher / 16f5fff8e214 / 5

- [api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets](resources--cdn_loadbalancer--reference--group-005.md#canonical-f70a802a472f29233d15247e79a05b57616eaa0321e76cdac07e88bd29ffa15b)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-b57675087227f85e31d2f1e413d8dfa0895ab0596b75681660d5755b97a89275)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-f70a802a472f29233d15247e79a05b57616eaa0321e76cdac07e88bd29ffa15b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f29201a3f4fdece9ccf913c1c86b2d93f45e0482a136a417f11ba90c457aac1"></a>

## api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / 1708fa18fb19 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-b57675087227f85e31d2f1e413d8dfa0895ab0596b75681660d5755b97a89275)
- [api_rate_limit.server_url_rules.client_matcher.ip_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-9c55db87686e1d62df013b59ba5fff3b2253597dd72994c171adeea78bcf40ab)
- api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-1134868b5494f90b54427a8818b3b4035413e70d53ba9e2a6ea7f49f5576604b"></a>

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

<a id="canonical-b3de888a33b1658624ced26432df7a1b1eac8076f93004fd185946a3f210048c"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / 1708fa18fb19 / 3

<a id="canonical-afc57c29906a839c56396f0fa70f6d06d42c6f3c529028afe20972bba598e748"></a>

<a id="canonical-b83cc9bdcd698e3b97ffea9521dd6e23f2ec728a47e8462d7b0e1767edbae0bb"></a>

## kind property — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / 1708fa18fb19 / 4

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

<a id="canonical-84d77d8dcc81233d3495bb4b80dfd5d0afd8396c86a6c66fefd69c7f28a92452"></a>

<a id="canonical-b5df7bb8aebe9f8bb8a1e9af81c050dbd6c176e80264cdb7dd7223df17598a71"></a>

## name property — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / 1708fa18fb19 / 5

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

<a id="canonical-4191cc10251d5526705d052f174cd057b6119be7da46565edacba5d14f6d4cb1"></a>

<a id="canonical-a03ca7bc473a960278a35ae64b3b0528eb3a321bc95ea669ee2ed53ec668739c"></a>

## namespace property — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / 1708fa18fb19 / 6

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

<a id="canonical-ace890935370fc432d9a994e968f64e8807b661de02253e637054f8e9a8247b8"></a>

<a id="canonical-49de5e22e5bd383516e8c4665d85630c72782e606209ab5ec06627ad7e254fab"></a>

## tenant property — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / 1708fa18fb19 / 7

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

<a id="canonical-50cb7ec6b3cdeb08070aae4e46ca378c82f153074f8f564acd45fbc83a6d40a0"></a>

<a id="canonical-ceaf0fa97f0b10fa12f2dbceb0c8e2a590ccd95a80c5eabd783e6b1d59dab209"></a>

## uid property — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / 1708fa18fb19 / 8

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

<a id="canonical-9ccdfc29d58ce64cf41d6a53e3a6b10ff78e6397397cdf8a1215268904fa7c3a"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / 1708fa18fb19 / 9

- [api_rate_limit.server_url_rules.client_matcher.ip_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-9c55db87686e1d62df013b59ba5fff3b2253597dd72994c171adeea78bcf40ab)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-59e6fca4a1bccd722f194595f12b3708fcb25fa15a1d709fd0257483726bebc3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5853f92a3e91be2cf7d57d4eb4f3e72859874991210e9abd246d80b5fb7f6579"></a>

## api_rate_limit.server_url_rules.client_matcher.ip_prefix_list — api_rate_limit.server_url_rules.client_matcher.ip_prefix_list / bb0b61e82e43 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-b57675087227f85e31d2f1e413d8dfa0895ab0596b75681660d5755b97a89275)
- api_rate_limit.server_url_rules.client_matcher.ip_prefix_list

<a id="canonical-bdb17f623a068bce50a5be35c907259d635cf82a869730dc8845e730351b6a2c"></a>

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

<a id="canonical-48a62ce86eeb7f785c501719f9ee36230798b23f84cfaad2f2693dec407e7222"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.ip_prefix_list / bb0b61e82e43 / 3

<a id="canonical-d7f77a11bdf0f2f8c44ce3457e0adca3c5f0ced2c3207b8d8dc1c23401711e11"></a>

<a id="canonical-2f5674380b2efa670a6314f86f4dd46b86dcebf60c662b6d40d476559dd69973"></a>

## invert_match property — api_rate_limit.server_url_rules.client_matcher.ip_prefix_list / bb0b61e82e43 / 4

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

<a id="canonical-ee5a292050381cf2aee0647ba6b4239b697a5908c7294e3ba4ef68e9dab2c3f5"></a>

<a id="canonical-bc782e32cbbe4be5ccc43a23a20ada7f863d80c6a4989cddff71c4aa6a6b991c"></a>

## ip_prefixes property — api_rate_limit.server_url_rules.client_matcher.ip_prefix_list / bb0b61e82e43 / 5

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

<a id="canonical-299c349aeb130930645c08728a5ce67ca3ea7633b54d325b3a131de120135b9e"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.ip_prefix_list / bb0b61e82e43 / 6

- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-b57675087227f85e31d2f1e413d8dfa0895ab0596b75681660d5755b97a89275)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-c041a8794d03af4f38420db7346f31412d98ac4b8c7710912d024310840cced7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e872a6be93e42cdfd25528e4e7e84f184b26c578217192ae11dad13472cfffaf"></a>

## api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list — api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list / 245a0b22f7ae / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-b57675087227f85e31d2f1e413d8dfa0895ab0596b75681660d5755b97a89275)
- api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list

<a id="canonical-d0f3ac207d49822d28f2ec4337c3070ba6813632c1271b69a6b145be34fd38c3"></a>

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

<a id="canonical-de092f060a6908979caeaaa42ed7b677fcb0328c90ba3374fe098a68cfcedc8c"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list / 245a0b22f7ae / 3

<a id="canonical-90896cfa4a22758a434f17b73dd1272da27d452bc1d14e4e9e17d96844cd3f2e"></a>

<a id="canonical-1c536e1c8f0208990c5268efaa440f24d26b1eb3e3f4275264916041443c74cb"></a>

## ip_threat_categories property — api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list / 245a0b22f7ae / 4

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

<a id="canonical-f2f7ee23e031ac1bc0b67dc7e3d14005d4d578029b3f9cd62bf6168102998c53"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list / 245a0b22f7ae / 5

- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-b57675087227f85e31d2f1e413d8dfa0895ab0596b75681660d5755b97a89275)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-1c0e7b8fa2c29802bd4ae73f2433ad98e3dfc9ca347233d65bf563e916bce200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a79a5d8608db28324cd61a0505b27c2051aab7d8ebbf9b9ec0d13f84529ab0ac"></a>

## api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher — api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher / 4ee95606ce52 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-b57675087227f85e31d2f1e413d8dfa0895ab0596b75681660d5755b97a89275)
- api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-74849db8178bb1c2dce7718cf196da7d83d7c61feaa978e7f9912272f11531c2"></a>

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

<a id="canonical-0328fe3597f694e19d1832469a7057a5d306d8687784f54938cdc14e10f1d918"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher / 4ee95606ce52 / 3

<a id="canonical-e3191851f59ea0abbea01204736ab61f8f39120e2b33476ed73342643b016d5b"></a>

<a id="canonical-3f7d490dce6c2fc4e589a78e2ddecc2091094fb00474ac3409b558df2808c491"></a>

## classes property — api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher / 4ee95606ce52 / 4

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

<a id="canonical-fb10eeb6a5f786730fd95c7fb4b036dc654e0889648a44bb7e44305535f0b7c5"></a>

<a id="canonical-cfd20373e283163b25a25f90557189c66454e69e380808e3c9508a43908876ba"></a>

## exact_values property — api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher / 4ee95606ce52 / 5

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

<a id="canonical-07640cc4e82602ca6cf4ea4521ab5a523932c73234873879031bb48bdf462c9e"></a>

<a id="canonical-bf71d2e016725e9c8afbc42bee8e0d86b9dd2c001fbf6a2f95aa1e1b72c2bb44"></a>

## excluded_values property — api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher / 4ee95606ce52 / 6

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

<a id="canonical-bf4d427de875618859152ffb5f76ad9dac8fc431d89599e59ea187b5e1d74cb3"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher / 4ee95606ce52 / 7

- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-b57675087227f85e31d2f1e413d8dfa0895ab0596b75681660d5755b97a89275)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-393655f0c8b97058a5d7e5b2d287e80ff14bd3369485e1e2c4a9165d13f6d549"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2ae3840a28d40e5c369ef49326824015eb32962259c511ea1fa4b9e4a8ebb501"></a>

## api_rate_limit.server_url_rules.inline_rate_limiter — api_rate_limit.server_url_rules.inline_rate_limiter / 1be4032fbaf7 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- api_rate_limit.server_url_rules.inline_rate_limiter

<a id="canonical-96058c2027527952ae667a11dc55a996c97df9951bedf398186a029939835396"></a>

Type: `"object"`. single nested block, Optional.

Inline rate-limiter settings for this domain, base-path, or endpoint rule. Select this field as the
required rate\_limiter\_choice when no stored rate-limiter object is used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("threshold"),
  validators.ConflictingObjectAttributes("ref_user_id",
    "use_http_lb_user_id")}
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
  "x-ves-oneof-field-count_by_choice": "[\"ref_user_id\",\"use_http_lb_user_id\"]"
}
```

Terraform syntax:

```terraform
inline_rate_limiter {
  # Configure direct properties listed below.
}
```

<a id="canonical-aa4d04ba93135b3994d6376d15d27337ee1cc8ab923922f46c9ae64b6fcd3b8b"></a>

## Direct properties — api_rate_limit.server_url_rules.inline_rate_limiter / 1be4032fbaf7 / 3

- [ref_user_id](resources--cdn_loadbalancer--reference--group-005.md#canonical-07d46f22199b05ae0ba5807c1200c358f9bd8b3483b03a9e487d6b37a8c2f657): complete subsection reference.

<a id="canonical-d15240eb2f94ecc907e57cf0629282c080cbddf2bef753bb839645c653b11114"></a>

<a id="canonical-34c211244a7230340974cb77ad1e4bbdcd2f83e10fdcbf83fa997f9645e38e89"></a>

## threshold property — api_rate_limit.server_url_rules.inline_rate_limiter / 1be4032fbaf7 / 4

Type: `"number"`. Optional.

The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified
period.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 8192),
}
```

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

<a id="canonical-5d3bf82b910352e9d9db0ba9f75c95c0bb321da717b64bcf2f8635c401c70007"></a>

<a id="canonical-c8fc32b594f62bcfe982147bf675daea71c2b0da8a4a865a5745992e1591bd22"></a>

## unit property — api_rate_limit.server_url_rules.inline_rate_limiter / 1be4032fbaf7 / 5

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SECOND",
    "MINUTE",
    "HOUR"),
}
```

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

- [use_http_lb_user_id](resources--cdn_loadbalancer--reference--group-005.md#canonical-ac98ed61addb2e08e6220e8903771335fb6706e2ca4ad17c4d338a0966f8c24d): complete subsection reference.

<a id="canonical-a80ea3086fe4f6012673b5aa7a6a57ddb34f0af43e6b327d09e8e73cd57d9c29"></a>

## Next pages — api_rate_limit.server_url_rules.inline_rate_limiter / 1be4032fbaf7 / 6

- [api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id](resources--cdn_loadbalancer--reference--group-005.md#canonical-07d46f22199b05ae0ba5807c1200c358f9bd8b3483b03a9e487d6b37a8c2f657)
- [api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id](resources--cdn_loadbalancer--reference--group-005.md#canonical-ac98ed61addb2e08e6220e8903771335fb6706e2ca4ad17c4d338a0966f8c24d)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-07d46f22199b05ae0ba5807c1200c358f9bd8b3483b03a9e487d6b37a8c2f657"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-66f957ccc953572a3d360cc4f83959f8c3928e5b775d87e4bf7e124f6a3233eb"></a>

## api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id — api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id / 8a248d5bd774 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [api_rate_limit.server_url_rules.inline_rate_limiter](resources--cdn_loadbalancer--reference--group-005.md#canonical-393655f0c8b97058a5d7e5b2d287e80ff14bd3369485e1e2c4a9165d13f6d549)
- api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id

<a id="canonical-a9438491bec990878d16fb1d24abc8f3d2493d4f84e48e91f5fba0a80846aa00"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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
ref_user_id {
  # Configure direct properties listed below.
}
```

<a id="canonical-5bfdb160e01715164f082fd3a19a9a37965a0eb812b6d1dbde4f7d1ab0f23e5a"></a>

## Direct properties — api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id / 8a248d5bd774 / 3

<a id="canonical-7dddd81ad8ae9b19cbcec1afef5fef7a98397e75f7bbd3682953cb5f2f745c64"></a>

<a id="canonical-f8edaa6db759a12d70a3e3c38a9a446bcc01cd3df7a064ed556dda7a5dc834ef"></a>

## name property — api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id / 8a248d5bd774 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-77988a7b303f88061b37b6df731f5e6d3d45c3db170a1e3b3f36011e723bba82"></a>

<a id="canonical-a72c60bd2a00c20725ea0da1edbb8df24f1d785733ad39f33534313deac0a7eb"></a>

## namespace property — api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id / 8a248d5bd774 / 5

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
}
```

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

<a id="canonical-2e7bcc2d153b8179801e6784ac40017b5e738556862f4b142a62db4876fc7586"></a>

<a id="canonical-cb1d7ac89ed22bbe285f4de47700946fb0dc676697c12c639ef6074101438162"></a>

## tenant property — api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id / 8a248d5bd774 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-b595aa10157783c1fa73c4b7a028704cca1df9d0131c8a7f2aac0fdff7b6cca1"></a>

## Next pages — api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id / 8a248d5bd774 / 7

- [api_rate_limit.server_url_rules.inline_rate_limiter](resources--cdn_loadbalancer--reference--group-005.md#canonical-393655f0c8b97058a5d7e5b2d287e80ff14bd3369485e1e2c4a9165d13f6d549)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-ac98ed61addb2e08e6220e8903771335fb6706e2ca4ad17c4d338a0966f8c24d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d0f70ac21c97fe4aa65c392d2960d06b34a683807ab931f0314814a50eb6b234"></a>

## api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id — api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id / ae664450ade4 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [api_rate_limit.server_url_rules.inline_rate_limiter](resources--cdn_loadbalancer--reference--group-005.md#canonical-393655f0c8b97058a5d7e5b2d287e80ff14bd3369485e1e2c4a9165d13f6d549)
- api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id

<a id="canonical-5b58fa422d0081af7907580c764bd65825ce9bd63414675e402b6135cdae113d"></a>

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
use_http_lb_user_id = {}
```

<a id="canonical-51affd6e83b179fa734fedc38e19496bf9256f880f15ee79a0cab3fc8e19bbda"></a>

## Direct properties — api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id / ae664450ade4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3e031b367fa32d02cc7db1a8852a4f60e546b8cd2f96919f74f8990af1ddb3b2"></a>

## Next pages — api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id / ae664450ade4 / 4

- [api_rate_limit.server_url_rules.inline_rate_limiter](resources--cdn_loadbalancer--reference--group-005.md#canonical-393655f0c8b97058a5d7e5b2d287e80ff14bd3369485e1e2c4a9165d13f6d549)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-6455b887a64665e97c723c6bbd85c3c03ecc56318d08f0dee33becd1a7bc652f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d1e9c8d574112d83240cf46621205b50087f1b913c879d862d8378f120505e35"></a>

## api_rate_limit.server_url_rules.ref_rate_limiter — api_rate_limit.server_url_rules.ref_rate_limiter / 96116dcc166a / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- api_rate_limit.server_url_rules.ref_rate_limiter

<a id="canonical-1ff693d8ef53c7a0e6daf7ef04c7ed61a23727cbf9ee9000f3e5f31cb66401a5"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

Reference to a stored rate-limiter object for this scoped rule. Select exactly one of
ref\_rate\_limiter and inline\_rate\_limiter.

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
ref_rate_limiter {
  # Configure direct properties listed below.
}
```

<a id="canonical-be4543e37062d5ac0d84d326673d4641908614ede2550f9b87aa0bb92cd52398"></a>

## Direct properties — api_rate_limit.server_url_rules.ref_rate_limiter / 96116dcc166a / 3

<a id="canonical-e8cc0d1c97c192cdc91447de610a45acbd3cc74afd4d491bd40bf0f87a5936df"></a>

<a id="canonical-f60657717315e8fd7bb4be73beb334ff6cbfa24faec4c9996229fab5c1397e8a"></a>

## name property — api_rate_limit.server_url_rules.ref_rate_limiter / 96116dcc166a / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-0a0319dfdde0f0051f4116f7ba29e39dabb114f248c342f7b432c9020011dee9"></a>

<a id="canonical-c6b7e6557bd2a03260bdcd0ad591f2921f914daf3ef12f89edd47bcf8050170b"></a>

## namespace property — api_rate_limit.server_url_rules.ref_rate_limiter / 96116dcc166a / 5

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
}
```

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

<a id="canonical-dd46f04cf9d1ef239b63792c88157afb9b425456e98e916ced298dc6725563cb"></a>

<a id="canonical-0b1d59143c34cfa196995ce6993d844c94e944c7f25ce78a7cec4947477f2f66"></a>

## tenant property — api_rate_limit.server_url_rules.ref_rate_limiter / 96116dcc166a / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-3b997bfa15b5499dde38ea8652a547c504694cc109d36d5a8d5b1ed632e3cdb0"></a>

## Next pages — api_rate_limit.server_url_rules.ref_rate_limiter / 96116dcc166a / 7

- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-f58e4a2bf6386ee36e45a3c28c6b34ee61b7613b7b852544427ffe9250b32422"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e610440b3edac9553e789de152f2269724bdebe3c54e8ae7edc66130dd703ab7"></a>

## api_rate_limit.server_url_rules.request_matcher — api_rate_limit.server_url_rules.request_matcher / b84a995b593a / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- api_rate_limit.server_url_rules.request_matcher

<a id="canonical-053df271b7ac2a9faa9537f0a19bdb42084c0e7469d4a56834c37607801b61d2"></a>

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

<a id="canonical-d88c60c6e5d0a43b3ebae3f8448034d44b1afa42024b497d96b3f3b8f15990de"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher / b84a995b593a / 3

- [cookie_matchers](resources--cdn_loadbalancer--reference--group-005.md#canonical-cf9560ecb236feabc592d96edc8d2dc99bdee9420fe8a4973bd000d5bcc5f34c): complete subsection reference.

- [headers](resources--cdn_loadbalancer--reference--group-005.md#canonical-0269399d886a72fc2f192977ab34ab75d488b4aa73e8b21ef0dd605e098454bb): complete subsection reference.

- [jwt_claims](resources--cdn_loadbalancer--reference--group-005.md#canonical-98086ef824e1db1a5188e3663959a96582764ecada27b1c1755f9233beb823b8): complete subsection reference.

- [query_params](resources--cdn_loadbalancer--reference--group-005.md#canonical-1adcb595b0692ace06159f871fdddcfdd17afb08a35d25dfc27219f2c13f96ed): complete subsection reference.

<a id="canonical-bec29f0946a1c4ba8b91eddeac6d724b172177b035ce6fed2ba007e5bd9c0fac"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher / b84a995b593a / 4

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-005.md#canonical-cf9560ecb236feabc592d96edc8d2dc99bdee9420fe8a4973bd000d5bcc5f34c)
- [api_rate_limit.server_url_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-005.md#canonical-0269399d886a72fc2f192977ab34ab75d488b4aa73e8b21ef0dd605e098454bb)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-005.md#canonical-98086ef824e1db1a5188e3663959a96582764ecada27b1c1755f9233beb823b8)
- [api_rate_limit.server_url_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-005.md#canonical-1adcb595b0692ace06159f871fdddcfdd17afb08a35d25dfc27219f2c13f96ed)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-cf9560ecb236feabc592d96edc8d2dc99bdee9420fe8a4973bd000d5bcc5f34c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ca82f5b69f3f81ab2f6e7a40e3a7a80a4efe2e3e5c10983b52725e43bce36d2"></a>

## api_rate_limit.server_url_rules.request_matcher.cookie_matchers — api_rate_limit.server_url_rules.request_matcher.cookie_matchers / 37e166551501 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-f58e4a2bf6386ee36e45a3c28c6b34ee61b7613b7b852544427ffe9250b32422)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers

<a id="canonical-1101cb3a85b8f94203027888b672c425a7c7c1fb9bc4bd9de2c7e4050f1080fc"></a>

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

<a id="canonical-6bead221502dec23796ba3ac0361de3f822a44b35ad88cb582c781b21b4b6b92"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.cookie_matchers / 37e166551501 / 3

- [check_not_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-0a2c74849b519e7406ef46ce5f0e4ef7821979075361c7474a2bad5137672682): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-58c3aede60737d1d028ec0d8c3e795feaaf5f747ecfdf7d71a3dbba58db4af7b): complete subsection reference.

<a id="canonical-56820d9154732b68aaeeea91b33b2a49ff3e469f8ae016a16464e69d97842a92"></a>

<a id="canonical-86e503daa7143ba7ad8d51ea4a73deaddf5883598dc0932a22ca2a842ca186b8"></a>

## invert_matcher property — api_rate_limit.server_url_rules.request_matcher.cookie_matchers / 37e166551501 / 4

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

- [item](resources--cdn_loadbalancer--reference--group-005.md#canonical-a00ca8cfee4177d1fccec35af4dde3b072ee803e901dc9e248f7f2b0e9d95c74): complete subsection reference.

<a id="canonical-ff15bb8f573556b78029ce783edf3e54320c2965d6a7d7a2946bc9336ca63292"></a>

<a id="canonical-26a01dd7977a5dbaf421e9061a4c66c61d4d0e30272ce7867d63b6b772418a10"></a>

## name property — api_rate_limit.server_url_rules.request_matcher.cookie_matchers / 37e166551501 / 5

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

<a id="canonical-fcd78179b2b27acd8046468636045bdd3ff94223d00b2f0cdfca4f2b69d2196d"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.cookie_matchers / 37e166551501 / 6

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-0a2c74849b519e7406ef46ce5f0e4ef7821979075361c7474a2bad5137672682)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-58c3aede60737d1d028ec0d8c3e795feaaf5f747ecfdf7d71a3dbba58db4af7b)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item](resources--cdn_loadbalancer--reference--group-005.md#canonical-a00ca8cfee4177d1fccec35af4dde3b072ee803e901dc9e248f7f2b0e9d95c74)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-f58e4a2bf6386ee36e45a3c28c6b34ee61b7613b7b852544427ffe9250b32422)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-0a2c74849b519e7406ef46ce5f0e4ef7821979075361c7474a2bad5137672682"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb03e5f32ad24be1415ecf9d1fe3f5d06bd3e31bb75d01e6595628ddd63355bf"></a>

## api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_present — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_presen / ef79d4f71bc9 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-f58e4a2bf6386ee36e45a3c28c6b34ee61b7613b7b852544427ffe9250b32422)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-005.md#canonical-cf9560ecb236feabc592d96edc8d2dc99bdee9420fe8a4973bd000d5bcc5f34c)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-a9928fb11e9f2ca93964e5583b6634f1a2dcaa180f8e94bf6426f57c51fbfc74"></a>

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

<a id="canonical-9f1345dd625bf37428f76faccb3a0030345e12cff834ca38a95a7cbb9e5c5e16"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_presen / ef79d4f71bc9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-53c051911961ef78b50ca4e6d2b4616a2859e22595579807eee2f4bad3a88396"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_presen / ef79d4f71bc9 / 4

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-005.md#canonical-cf9560ecb236feabc592d96edc8d2dc99bdee9420fe8a4973bd000d5bcc5f34c)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-58c3aede60737d1d028ec0d8c3e795feaaf5f747ecfdf7d71a3dbba58db4af7b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ba73cb71b9be69b11f6f011b54247623e9637149653d707c4324bb23c084b87f"></a>

## api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present / a47963890e11 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-f58e4a2bf6386ee36e45a3c28c6b34ee61b7613b7b852544427ffe9250b32422)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-005.md#canonical-cf9560ecb236feabc592d96edc8d2dc99bdee9420fe8a4973bd000d5bcc5f34c)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-66ede8bfd0b5ba04e97ea20257b4ef6b2604312df0c526374f461ce5fd5a4072"></a>

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

<a id="canonical-a95d1990b3478de05451d431359688d8a7a1cea64caf28647895f4882f4c881b"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present / a47963890e11 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b6df3d149ad54612b516ec3a046ee04bf70af9f019a3f205adad4d3cba844aec"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present / a47963890e11 / 4

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-005.md#canonical-cf9560ecb236feabc592d96edc8d2dc99bdee9420fe8a4973bd000d5bcc5f34c)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-a00ca8cfee4177d1fccec35af4dde3b072ee803e901dc9e248f7f2b0e9d95c74"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e83990616330dc064948986ec8ecd5a1b4ed0a11e2d37228e48440e2066a68b"></a>

## api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item / d81b6935c79f / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-f58e4a2bf6386ee36e45a3c28c6b34ee61b7613b7b852544427ffe9250b32422)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-005.md#canonical-cf9560ecb236feabc592d96edc8d2dc99bdee9420fe8a4973bd000d5bcc5f34c)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item

<a id="canonical-693ea96da0a253119d3130a06a46b7f714459b73f9d143b0596a5534c20a4d83"></a>

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

<a id="canonical-28f18bafd800f1e3365f8fe1d519c9eea9180673bba2de8e5408a67eb7309df8"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item / d81b6935c79f / 3

<a id="canonical-fee9b85a664d6be415ef23e32f2f2dfec72acf05732384e26bc94604b744eb5a"></a>

<a id="canonical-8856cd00f55a7e1c2fdd1aca9fd5aef1a67ce214b88df4380af1812303f5acc5"></a>

## exact_values property — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item / d81b6935c79f / 4

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

<a id="canonical-228d77231ee321862c9182c5e54c0089d56b3c5161f2f1c889e7d99c0e82915f"></a>

<a id="canonical-d5df3f8a2bc7f6e7c2e8e525c29da7d6540ee0c41bf4a168c3d1137083bd9525"></a>

## regex_values property — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item / d81b6935c79f / 5

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

<a id="canonical-e603d769aa0c425edc7e67c789521e4b13e2ad59aa9429d16f343c84e3d1e84d"></a>

<a id="canonical-c46c46df0ba08ab95c7fc6da1c9fe34b5d710db83d5974b5fb3e5ad70af482ee"></a>

## transformers property — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item / d81b6935c79f / 6

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

<a id="canonical-0800cc6c51a48a8fd084a09bd39f1c2c86aa64d9ca5b5c80459414766b4eacca"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item / d81b6935c79f / 7

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-005.md#canonical-cf9560ecb236feabc592d96edc8d2dc99bdee9420fe8a4973bd000d5bcc5f34c)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-0269399d886a72fc2f192977ab34ab75d488b4aa73e8b21ef0dd605e098454bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dce2d391d339a07d812df74c1db8ecf9769ecd9c9cf8be74ae7e2b722bd6a985"></a>

## api_rate_limit.server_url_rules.request_matcher.headers — api_rate_limit.server_url_rules.request_matcher.headers / a3eb68196755 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-f58e4a2bf6386ee36e45a3c28c6b34ee61b7613b7b852544427ffe9250b32422)
- api_rate_limit.server_url_rules.request_matcher.headers

<a id="canonical-7de9f2018ba83bdfbb20218da779a3ee863ba266aeb510fc32345c04ff65888f"></a>

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

<a id="canonical-ba62a43c3ca822c154f114bb3b238a4e53ee9bf667d97d8887f9725bd0d1bfa8"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.headers / a3eb68196755 / 3

- [check_not_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-f77cb51fedfa1ee12e7b6475ec70ed8166497a96e20977ca0ea87b329afd02d9): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-7d063a2d64bb9717022ffd2858d55923db6d8e78fc7df43b2fe0cd64582ad37e): complete subsection reference.

<a id="canonical-2bcade07ea4d268ca5d392ab4186828654ee832dc58714022b4c5d1216f645e1"></a>

<a id="canonical-1884682e5325e122a2b82e1832ea607f7103c18970512a3f5c9bce4612780028"></a>

## invert_matcher property — api_rate_limit.server_url_rules.request_matcher.headers / a3eb68196755 / 4

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

- [item](resources--cdn_loadbalancer--reference--group-005.md#canonical-29b6be73d8c21a88588058a4e5f7561c28f8fb40450c8b14357c30cdec191be4): complete subsection reference.

<a id="canonical-0e8f3b6b33664b6c5e17fa71026d578b7c82909e3b0a07b31d78ab975e84eb37"></a>

<a id="canonical-4f9351a45bb2428bf94ea3164f959aad5549a530d77790a53294a37612754bc5"></a>

## name property — api_rate_limit.server_url_rules.request_matcher.headers / a3eb68196755 / 5

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

<a id="canonical-ad82e50ec24153dc823f3f9faf0618eab831f69f48f2e77fb0c5ee9c58dd1312"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.headers / a3eb68196755 / 6

- [api_rate_limit.server_url_rules.request_matcher.headers.check_not_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-f77cb51fedfa1ee12e7b6475ec70ed8166497a96e20977ca0ea87b329afd02d9)
- [api_rate_limit.server_url_rules.request_matcher.headers.check_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-7d063a2d64bb9717022ffd2858d55923db6d8e78fc7df43b2fe0cd64582ad37e)
- [api_rate_limit.server_url_rules.request_matcher.headers.item](resources--cdn_loadbalancer--reference--group-005.md#canonical-29b6be73d8c21a88588058a4e5f7561c28f8fb40450c8b14357c30cdec191be4)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-f58e4a2bf6386ee36e45a3c28c6b34ee61b7613b7b852544427ffe9250b32422)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-f77cb51fedfa1ee12e7b6475ec70ed8166497a96e20977ca0ea87b329afd02d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-07561ec0371a2b0e38af64bc8da2e7fcc3e574292dfcf431648f393238272b21"></a>

## api_rate_limit.server_url_rules.request_matcher.headers.check_not_present — api_rate_limit.server_url_rules.request_matcher.headers.check_not_present / 68bbc0b3c90e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-f58e4a2bf6386ee36e45a3c28c6b34ee61b7613b7b852544427ffe9250b32422)
- [api_rate_limit.server_url_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-005.md#canonical-0269399d886a72fc2f192977ab34ab75d488b4aa73e8b21ef0dd605e098454bb)
- api_rate_limit.server_url_rules.request_matcher.headers.check_not_present

<a id="canonical-dc306b6773efec18b06bcc4284c2c8bc34ef5adb72e02a462e58a839335bbe39"></a>

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

<a id="canonical-a69fb701db10c7a39bead510f0a6113e09a250a8264c2e077c96b5c2e596942c"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.headers.check_not_present / 68bbc0b3c90e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-72613d0b64c16315a2d322a404071b0561ce0aae9e3bac471c323a124793aecb"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.headers.check_not_present / 68bbc0b3c90e / 4

- [api_rate_limit.server_url_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-005.md#canonical-0269399d886a72fc2f192977ab34ab75d488b4aa73e8b21ef0dd605e098454bb)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-7d063a2d64bb9717022ffd2858d55923db6d8e78fc7df43b2fe0cd64582ad37e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a21428fdb929d50b68b51ef3cad9e3c9d04272e87d3a72720c708767f5e8c16"></a>

## api_rate_limit.server_url_rules.request_matcher.headers.check_present — api_rate_limit.server_url_rules.request_matcher.headers.check_present / f4efd666f44b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-f58e4a2bf6386ee36e45a3c28c6b34ee61b7613b7b852544427ffe9250b32422)
- [api_rate_limit.server_url_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-005.md#canonical-0269399d886a72fc2f192977ab34ab75d488b4aa73e8b21ef0dd605e098454bb)
- api_rate_limit.server_url_rules.request_matcher.headers.check_present

<a id="canonical-7fdebbbfa80279c3ec23fe3e0ed98475bf5a253935f9d90776a4f5feb1323b52"></a>

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

<a id="canonical-9d2397a97a4242f042dfd5e4c22cf4af1d4300f36f947da4e69d59d18eac4d9a"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.headers.check_present / f4efd666f44b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1ec3b070bae40b996bd6bd1ce1aa778f23a7cfce76dbd30c7ebd99d28246020e"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.headers.check_present / f4efd666f44b / 4

- [api_rate_limit.server_url_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-005.md#canonical-0269399d886a72fc2f192977ab34ab75d488b4aa73e8b21ef0dd605e098454bb)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-29b6be73d8c21a88588058a4e5f7561c28f8fb40450c8b14357c30cdec191be4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da3d78a33a4037d928517c76eb8557e8c2ca692f8afdabeac1e6204974f70018"></a>

## api_rate_limit.server_url_rules.request_matcher.headers.item — api_rate_limit.server_url_rules.request_matcher.headers.item / ce925043e2b8 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-f58e4a2bf6386ee36e45a3c28c6b34ee61b7613b7b852544427ffe9250b32422)
- [api_rate_limit.server_url_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-005.md#canonical-0269399d886a72fc2f192977ab34ab75d488b4aa73e8b21ef0dd605e098454bb)
- api_rate_limit.server_url_rules.request_matcher.headers.item

<a id="canonical-d6091df9968ec73a65bb9261575e863d2f79f360a8cb2b80ca54e1ee77d0fa6a"></a>

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

<a id="canonical-f7b83d690747186096c85f61d916048c823a6421a4e33f055eb1936a1beba3bf"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.headers.item / ce925043e2b8 / 3

<a id="canonical-09633c55f9b07da9ff0bc378c056151979ddbfde446d68833cfaa565ba8b459b"></a>

<a id="canonical-b86270921b6871057fa0ef96de4e8cc3e26602b1694d22a9102e7a8fd72e49ff"></a>

## exact_values property — api_rate_limit.server_url_rules.request_matcher.headers.item / ce925043e2b8 / 4

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

<a id="canonical-aefad3b9fcaa1c73978b0b716a0cce504e11c0c15e584872890f111041dedb7c"></a>

<a id="canonical-46a5f11d94b155d6ee5367afb63c2bdf3360f3c17ab674d533317a350465f469"></a>

## regex_values property — api_rate_limit.server_url_rules.request_matcher.headers.item / ce925043e2b8 / 5

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

<a id="canonical-2313f1a178948d13dd19568f22c535cb7fde8d38aed8ff77da330c6a5e8c9efc"></a>

<a id="canonical-78bf79388691403d80bc7727a079386c208b54b6af6283112b60a60db2a9b0b6"></a>

## transformers property — api_rate_limit.server_url_rules.request_matcher.headers.item / ce925043e2b8 / 6

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

<a id="canonical-728aa40da918f55a194504555ba33aeac7ffc616136d9f4fd8a10296e9b51d6c"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.headers.item / ce925043e2b8 / 7

- [api_rate_limit.server_url_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-005.md#canonical-0269399d886a72fc2f192977ab34ab75d488b4aa73e8b21ef0dd605e098454bb)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-98086ef824e1db1a5188e3663959a96582764ecada27b1c1755f9233beb823b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c944c3ad8d4df91070fb99791e8ca7ec59e3cafe23ed49934df726b56f7e670d"></a>

## api_rate_limit.server_url_rules.request_matcher.jwt_claims — api_rate_limit.server_url_rules.request_matcher.jwt_claims / 0854f2a1c526 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-f58e4a2bf6386ee36e45a3c28c6b34ee61b7613b7b852544427ffe9250b32422)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims

<a id="canonical-e6d972cf8f6d4bbad7f2a6899fb17d9d875d0f5c9c78f1b595ebae9eee63ad35"></a>

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

<a id="canonical-8f99030ad2ab26cdd74a145a460abfad388934b76177fd8858c3dea433b7d2e8"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.jwt_claims / 0854f2a1c526 / 3

- [check_not_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-54ec0699f3c72db47880357bce1688d89d3e15d2a3de29775eec17da51f4684d): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-baf58219025f619e15880d45900a3b67e0d901b4ea1385d3f7ca1780f94ba95d): complete subsection reference.

<a id="canonical-0cd68e81821051168ca0880efab220037ef549dd3786c78e00c5488619116bb1"></a>

<a id="canonical-61e36c891cbb48310b54164c723c484f7358275507d5fb4ad677d080c6f19d5b"></a>

## invert_matcher property — api_rate_limit.server_url_rules.request_matcher.jwt_claims / 0854f2a1c526 / 4

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

- [item](resources--cdn_loadbalancer--reference--group-005.md#canonical-0329bfcf1b677739997aa5561c14f2b9ae60890029f4d3cc2f1f632185c28701): complete subsection reference.

<a id="canonical-c27e3d3ccb4b9496576d7bc7000c6ebd5438dfa133adf563ddb917222e65f06a"></a>

<a id="canonical-ec0ab65b5d37cd9397e57e08fd73180399bf2a9bc71a1ceab6c6244b5e30802f"></a>

## name property — api_rate_limit.server_url_rules.request_matcher.jwt_claims / 0854f2a1c526 / 5

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

<a id="canonical-ecbbb3baf919e69cb61dca47c89cebd1acc3991eaf735f4c5a569188ccae86bb"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.jwt_claims / 0854f2a1c526 / 6

- [api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-54ec0699f3c72db47880357bce1688d89d3e15d2a3de29775eec17da51f4684d)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-baf58219025f619e15880d45900a3b67e0d901b4ea1385d3f7ca1780f94ba95d)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims.item](resources--cdn_loadbalancer--reference--group-005.md#canonical-0329bfcf1b677739997aa5561c14f2b9ae60890029f4d3cc2f1f632185c28701)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-f58e4a2bf6386ee36e45a3c28c6b34ee61b7613b7b852544427ffe9250b32422)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-54ec0699f3c72db47880357bce1688d89d3e15d2a3de29775eec17da51f4684d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ce2c5882698613d934ab6e1b5767ce0d8c9c7b404c534f146db7a084f4304a9"></a>

## api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present — api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present / 2a25ef640fa6 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-f58e4a2bf6386ee36e45a3c28c6b34ee61b7613b7b852544427ffe9250b32422)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-005.md#canonical-98086ef824e1db1a5188e3663959a96582764ecada27b1c1755f9233beb823b8)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-6545c3f4773fe781e2f08704e445bb1e0fdc81824e156c6c346608525e4b9743"></a>

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

<a id="canonical-cb29bf20149bffbb7c128749cf3b20a3c911628d29e23526f4be69ca6cdec310"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present / 2a25ef640fa6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8cfbe78b99a8492e2ee4cae56a05364004a82dacf11ca3bf0b74a25226ff5754"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present / 2a25ef640fa6 / 4

- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-005.md#canonical-98086ef824e1db1a5188e3663959a96582764ecada27b1c1755f9233beb823b8)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-baf58219025f619e15880d45900a3b67e0d901b4ea1385d3f7ca1780f94ba95d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-650d48755be2777e49bad2cd51b2d84f76548fde6486358afb298a44024a9b71"></a>

## api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present — api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present / f1c2eb084176 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-f58e4a2bf6386ee36e45a3c28c6b34ee61b7613b7b852544427ffe9250b32422)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-005.md#canonical-98086ef824e1db1a5188e3663959a96582764ecada27b1c1755f9233beb823b8)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present

<a id="canonical-aece0afe6e6d6b646aefc481265fee20cecad75a0594c2ef950f223d5811e74a"></a>

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

<a id="canonical-8f58de587ef8c99e1f980610151fe4ec466afecfdd4a4f5ae3482af742d2cf06"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present / f1c2eb084176 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fabb41bcbdb854d122aecf7e0ded217cdbfce78d895b6817603f8ccd7b75e362"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present / f1c2eb084176 / 4

- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-005.md#canonical-98086ef824e1db1a5188e3663959a96582764ecada27b1c1755f9233beb823b8)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-0329bfcf1b677739997aa5561c14f2b9ae60890029f4d3cc2f1f632185c28701"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab17fe724c8da662d74951254c5ef8692a3f955a4544d224e09b278f97091880"></a>

## api_rate_limit.server_url_rules.request_matcher.jwt_claims.item — api_rate_limit.server_url_rules.request_matcher.jwt_claims.item / 9ca3d558cd4f / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-f58e4a2bf6386ee36e45a3c28c6b34ee61b7613b7b852544427ffe9250b32422)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-005.md#canonical-98086ef824e1db1a5188e3663959a96582764ecada27b1c1755f9233beb823b8)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.item

<a id="canonical-99b7991ea303312488c262bf426a5fef306ee13735d4c1e5d790b13ba45500f0"></a>

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

<a id="canonical-8600bd0343bbc266c7934d38f21653e5c7bbb68bd558f19f71c79933d48247ce"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.jwt_claims.item / 9ca3d558cd4f / 3

<a id="canonical-8fc5cda0beaa3bfd6576663b232478a9aab69400e21a0a88a44718334af7651b"></a>

<a id="canonical-b38c5c3f2e551f885abc8f4e497ee82a36bd8d1ea360d35ba5d5d1e0defffedf"></a>

## exact_values property — api_rate_limit.server_url_rules.request_matcher.jwt_claims.item / 9ca3d558cd4f / 4

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

<a id="canonical-ce59d56ec018b3c14806e4ab8511c57a8d22649d60a298726001ebd2f4b9e0bb"></a>

<a id="canonical-78c60e0884035a4cf8363bbfa4cc6908dd57c6d3533cf27e06137f0713baf6ed"></a>

## regex_values property — api_rate_limit.server_url_rules.request_matcher.jwt_claims.item / 9ca3d558cd4f / 5

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

<a id="canonical-9321c0f374229362d81370f8b8efd41b18f598c5ff1f6908688aaefb0d4bb3e7"></a>

<a id="canonical-e427bfac42303769600cf8890d263ccb61a5e3b96ff770f042258a523d52a508"></a>

## transformers property — api_rate_limit.server_url_rules.request_matcher.jwt_claims.item / 9ca3d558cd4f / 6

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

<a id="canonical-2874dca987037c349aa1b75d212c6c042f7b20c476d2cefd37d32ba8ad727954"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.jwt_claims.item / 9ca3d558cd4f / 7

- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-005.md#canonical-98086ef824e1db1a5188e3663959a96582764ecada27b1c1755f9233beb823b8)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-1adcb595b0692ace06159f871fdddcfdd17afb08a35d25dfc27219f2c13f96ed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4364c38dde25a2733b23e5d9bda8f473fe18d2d3d0c2beba16fe16e9aa67bc43"></a>

## api_rate_limit.server_url_rules.request_matcher.query_params — api_rate_limit.server_url_rules.request_matcher.query_params / 056bbaeedd36 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-f58e4a2bf6386ee36e45a3c28c6b34ee61b7613b7b852544427ffe9250b32422)
- api_rate_limit.server_url_rules.request_matcher.query_params

<a id="canonical-6de6aa940f5ebec244570a30b5e6d8fb054fc61383a030e21ef97e5b129bc6c3"></a>

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

<a id="canonical-b6c1fb3abf649c22ce6cb9dad910759609dc6294811a7dd9e0649ef46a1906c7"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.query_params / 056bbaeedd36 / 3

- [check_not_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-0d94cfc9d97f03a2aa3905e62aeb5a75e01d315a8e07dfef5973df949cc235d4): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-f63e5d37f3cce361379b520342c186f4d241b9d3a4eb01ec0b6b6ee28f3ca031): complete subsection reference.

<a id="canonical-c165d34c645372540146256689ce0c69e378f3931a0d65dcf481908337d24d5c"></a>

<a id="canonical-d3619791983dbd9f1ba857eb4736fa37705c3982f19577ae2f3c9a877caa2e5a"></a>

## invert_matcher property — api_rate_limit.server_url_rules.request_matcher.query_params / 056bbaeedd36 / 4

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

- [item](resources--cdn_loadbalancer--reference--group-005.md#canonical-cacc8aee55ef2ca5cb074b81717eb3b28da6844cc5d724c9e7f30b2f3246cbd9): complete subsection reference.

<a id="canonical-229643623f68c80857d762d560e25477a808fbafde0f38ed314b43c3f234c408"></a>

<a id="canonical-5919780c5c821e595cdef614d6601d33ebc2addbba771e665b63e5fd4fe45c29"></a>

## key property — api_rate_limit.server_url_rules.request_matcher.query_params / 056bbaeedd36 / 5

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

<a id="canonical-62061caf8de1c01c3ffe5cf2aac134af8fa3b4ced93fa3116680b7c7c0b3dce0"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.query_params / 056bbaeedd36 / 6

- [api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-0d94cfc9d97f03a2aa3905e62aeb5a75e01d315a8e07dfef5973df949cc235d4)
- [api_rate_limit.server_url_rules.request_matcher.query_params.check_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-f63e5d37f3cce361379b520342c186f4d241b9d3a4eb01ec0b6b6ee28f3ca031)
- [api_rate_limit.server_url_rules.request_matcher.query_params.item](resources--cdn_loadbalancer--reference--group-005.md#canonical-cacc8aee55ef2ca5cb074b81717eb3b28da6844cc5d724c9e7f30b2f3246cbd9)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-f58e4a2bf6386ee36e45a3c28c6b34ee61b7613b7b852544427ffe9250b32422)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-0d94cfc9d97f03a2aa3905e62aeb5a75e01d315a8e07dfef5973df949cc235d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d5e8552957aca942fe2383790125e46c23db9d14b17a8cec5570966561e91938"></a>

## api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present — api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present / 77097a3456dd / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-f58e4a2bf6386ee36e45a3c28c6b34ee61b7613b7b852544427ffe9250b32422)
- [api_rate_limit.server_url_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-005.md#canonical-1adcb595b0692ace06159f871fdddcfdd17afb08a35d25dfc27219f2c13f96ed)
- api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present

<a id="canonical-0665fa884c75ab956fc8ff09ecef92c9207b89d0256e1ac13c56bb93142e65ed"></a>

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

<a id="canonical-2d2eef25ee2e2c077aba528872913813a25a6bbc77e76491c4d20ab4f5f9b52a"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present / 77097a3456dd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b2dc911a9b5d129871dc1abcfccb5fc42d06b860326f2923f5d5619ec9d5c3f5"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present / 77097a3456dd / 4

- [api_rate_limit.server_url_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-005.md#canonical-1adcb595b0692ace06159f871fdddcfdd17afb08a35d25dfc27219f2c13f96ed)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-f63e5d37f3cce361379b520342c186f4d241b9d3a4eb01ec0b6b6ee28f3ca031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e75c5791875862913b9d79f7553f6155b98a377355c7c3c6714023ee16999e4"></a>

## api_rate_limit.server_url_rules.request_matcher.query_params.check_present — api_rate_limit.server_url_rules.request_matcher.query_params.check_present / c13a92def7da / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-f58e4a2bf6386ee36e45a3c28c6b34ee61b7613b7b852544427ffe9250b32422)
- [api_rate_limit.server_url_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-005.md#canonical-1adcb595b0692ace06159f871fdddcfdd17afb08a35d25dfc27219f2c13f96ed)
- api_rate_limit.server_url_rules.request_matcher.query_params.check_present

<a id="canonical-21835d9466ba2d792620df19e2130fdebe26be154abaa1d4ffba42c234f95932"></a>

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

<a id="canonical-7404bd0d4cf13822d725b2b1d013004cde21baa890f5db90ff09c65337b50eee"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.query_params.check_present / c13a92def7da / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-44c3d5710705d710236a21ae2f862c477bef598114ccb9c0a3247b55aa0f6546"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.query_params.check_present / c13a92def7da / 4

- [api_rate_limit.server_url_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-005.md#canonical-1adcb595b0692ace06159f871fdddcfdd17afb08a35d25dfc27219f2c13f96ed)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-cacc8aee55ef2ca5cb074b81717eb3b28da6844cc5d724c9e7f30b2f3246cbd9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce8f3216ced01a3aabd3351d9d3c5c158dfb53f8aa41b94fa53e12d2e4413b17"></a>

## api_rate_limit.server_url_rules.request_matcher.query_params.item — api_rate_limit.server_url_rules.request_matcher.query_params.item / b6de3281911e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-f58e4a2bf6386ee36e45a3c28c6b34ee61b7613b7b852544427ffe9250b32422)
- [api_rate_limit.server_url_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-005.md#canonical-1adcb595b0692ace06159f871fdddcfdd17afb08a35d25dfc27219f2c13f96ed)
- api_rate_limit.server_url_rules.request_matcher.query_params.item

<a id="canonical-65e69645d55ca5aaf9bd100c9d4efe727f8122ac23ed6d288d9066eea27baaae"></a>

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

<a id="canonical-cedd332127bfd4f2a0bf0e10c97b649ac354935a0ec13563bb34e74bde962716"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.query_params.item / b6de3281911e / 3

<a id="canonical-f16c5cd7c93044b4aa751464d0055dd0f12f370ceafa173e6b2887f8964c0bcf"></a>

<a id="canonical-73ceedbb1b3bad419f659878c2a233aac34a313449adc206596e91e9ef7142b8"></a>

## exact_values property — api_rate_limit.server_url_rules.request_matcher.query_params.item / b6de3281911e / 4

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

<a id="canonical-99a5728c51746f040b83c83ededb400b69a775c201d1e5c06654d04d185911a3"></a>

<a id="canonical-c98055c8a3d24fb5811d49c52b22cb89aceddfa4f781dd1a6187ca78b1938991"></a>

## regex_values property — api_rate_limit.server_url_rules.request_matcher.query_params.item / b6de3281911e / 5

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

<a id="canonical-a78e4ffb9c5f204026b61ba34b176f399c3ce089364beb251b74f15958cc04f4"></a>

<a id="canonical-12780c170ef2d2d93cb898457fcc2cbf540825cf3527ed031b7c7d6a697022a6"></a>

## transformers property — api_rate_limit.server_url_rules.request_matcher.query_params.item / b6de3281911e / 6

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

<a id="canonical-7a1981acd8112ddd671a6345d7ce51ca5a3a9ff2da2e597cdbb424b7f421bc99"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.query_params.item / b6de3281911e / 7

- [api_rate_limit.server_url_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-005.md#canonical-1adcb595b0692ace06159f871fdddcfdd17afb08a35d25dfc27219f2c13f96ed)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9001de24fc9afc6b13272225b4cb598cc0e198d825a1e89e6b5a355e527e557"></a>

## api_specification — api_specification / 92e8bc583eba / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- api_specification

<a id="canonical-1234963d25aae7784a0ba7042903fd38c6ee5c9427ef88132213bfaa3c07b408"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: api\_specification, disable\_api\_definition; Default: disable\_api\_definition\] Settings
for API specification (API definition, OpenAPI validation, etc.).

Upstream description:

Settings for API specification (API definition, OpenAPI validation, etc.)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("validation_all_spec_endpoints",
    "validation_custom_list"),
  validators.ConflictingObjectAttributes("validation_all_spec_endpoints",
    "validation_disabled"),
  validators.ConflictingObjectAttributes("validation_custom_list",
    "validation_disabled")}
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
  "x-ves-oneof-field-validation_target_choice": "[\"validation_all_spec_endpoints\",\"validation_custom_list\",\"validation_disabled\"]"
}
```

OneOf alternatives in this subsection:

- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-1234963d25aae7784a0ba7042903fd38c6ee5c9427ef88132213bfaa3c07b408)
- [disable_api_definition](resources--cdn_loadbalancer--reference--group-010.md#canonical-47b71f79e1e4f34933aff9319652ff32324186f11ef4ffaf366520726a328f2d)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
api_specification {
  # Configure direct properties listed below.
}
```

<a id="canonical-59fec6c4ee1b24ce54e35b42901087670507bc495d569b24555afbf64566dfd0"></a>

## Direct properties — api_specification / 92e8bc583eba / 3

- [api_definition](resources--cdn_loadbalancer--reference--group-005.md#canonical-fef6f5c57291b09a520b2b242ad78d692213ad482f56b54094c7f0aa371c9908): complete subsection reference.

- [validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf): complete subsection reference.

- [validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f): complete subsection reference.

- [validation_disabled](resources--cdn_loadbalancer--reference--group-006.md#canonical-e1722fb09b74ea9433bd14a1bc5b67040e41a34915ba93990c87f0a463655eee): complete subsection reference.

<a id="canonical-9aa8fe4a6f84d037a98baf0a29c7d5a2a9cc1fb9ead071b82d2aa957b716ed39"></a>

## Next pages — api_specification / 92e8bc583eba / 4

- [api_specification.api_definition](resources--cdn_loadbalancer--reference--group-005.md#canonical-fef6f5c57291b09a520b2b242ad78d692213ad482f56b54094c7f0aa371c9908)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- [api_specification.validation_disabled](resources--cdn_loadbalancer--reference--group-006.md#canonical-e1722fb09b74ea9433bd14a1bc5b67040e41a34915ba93990c87f0a463655eee)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-fef6f5c57291b09a520b2b242ad78d692213ad482f56b54094c7f0aa371c9908"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3fce60b2b98ed2855390617a4da22ddc3b34a100dff35672b77309f32d27166a"></a>

## api_specification.api_definition — api_specification.api_definition / 595d5a7f7854 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- api_specification.api_definition

<a id="canonical-1e2ef487f18ff2d44be4ec6ccf600ffe66a21d1f3cebfa2deba727646ce08df9"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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
api_definition {
  # Configure direct properties listed below.
}
```

<a id="canonical-196e5d200e52f277edb2392fae930654c63f636428bf6365d2fe4538dab11172"></a>

## Direct properties — api_specification.api_definition / 595d5a7f7854 / 3

<a id="canonical-f851047b7f523bde8b13b933fe90525e8b09198ae3e790b60f45a6e32ab7c4e1"></a>

<a id="canonical-b0fd508acd1866675316b8cbf2bbb06b06c4dad87544a5ccd7afce7cd997abfe"></a>

## name property — api_specification.api_definition / 595d5a7f7854 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-1617cda02f730f8dd0397426be343ccb5ff30f4fa442d080d89f4cda939f7573"></a>

<a id="canonical-38b4b67d4c31a9534aa2c47f6ea271579b57a17873c3ccfc36c53726fe708282"></a>

## namespace property — api_specification.api_definition / 595d5a7f7854 / 5

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
}
```

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

<a id="canonical-19f7d193df8b130ebe0a119e29d64d7c3122bb3596d223ba6a8ea0d496d84870"></a>

<a id="canonical-5a148aaefb056988ea4394e9d45b92cf6ee1830025f427c1a6f887a2323932ef"></a>

## tenant property — api_specification.api_definition / 595d5a7f7854 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-adae66ffea29020b49e276d506304f6b448cccccbbe95ba77ea618099c48ee5b"></a>

## Next pages — api_specification.api_definition / 595d5a7f7854 / 7

- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23bd0d5987d34ef58bfad14946b5824f1a35dd3edc83bfaa318fab7e7758665a"></a>

## api_specification.validation_all_spec_endpoints — api_specification.validation_all_spec_endpoints / 846776d7af0a / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- api_specification.validation_all_spec_endpoints

<a id="canonical-a2ac4ec8bae105db09720c3d16ee1d0bac810d41a9d672a56ac84b4e1c8a9319"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
validation_all_spec_endpoints {
  # Configure direct properties listed below.
}
```

<a id="canonical-13a47c0890007899181a65612124b9783719c1249fe30204579a955a9e581ba0"></a>

## Direct properties — api_specification.validation_all_spec_endpoints / 846776d7af0a / 3

- [fall_through_mode](resources--cdn_loadbalancer--reference--group-005.md#canonical-5e885c4b750e99832ba179430a913c3720db297bc793085f9d1cbd82e4a0303e): complete subsection reference.

- [settings](resources--cdn_loadbalancer--reference--group-005.md#canonical-a66a49e08574d71974d3dce003662351a6bf6d897f2b85ae140c3b35e63b68f4): complete subsection reference.

- [validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-0e02e88b49ceb8eeb5b6a7013153c2a07a7c8b4d1a06d15dde773cd3479ac01d): complete subsection reference.

<a id="canonical-b97c62ba41dfbe53cdf105e01e3abd27b71373cd6707be53c9206e051186cbbd"></a>

## Next pages — api_specification.validation_all_spec_endpoints / 846776d7af0a / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--reference--group-005.md#canonical-5e885c4b750e99832ba179430a913c3720db297bc793085f9d1cbd82e4a0303e)
- [api_specification.validation_all_spec_endpoints.settings](resources--cdn_loadbalancer--reference--group-005.md#canonical-a66a49e08574d71974d3dce003662351a6bf6d897f2b85ae140c3b35e63b68f4)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-0e02e88b49ceb8eeb5b6a7013153c2a07a7c8b4d1a06d15dde773cd3479ac01d)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-5e885c4b750e99832ba179430a913c3720db297bc793085f9d1cbd82e4a0303e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-91b85ebec1dc431aa20af19d562ed59b69b6f077500582353d53c5b9ed9c2ef6"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode — api_specification.validation_all_spec_endpoints.fall_through_mode / 79cca19ce599 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf)
- api_specification.validation_all_spec_endpoints.fall_through_mode

<a id="canonical-c4e52e0758e3a743cb55dcc47f69115be85a75d9ce75f2727fd0d2cd54be06ad"></a>

Type: `"object"`. single nested block, Optional.

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules).

Upstream description:

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("fall_through_mode_allow",
    "fall_through_mode_custom")}
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
  "x-ves-oneof-field-fall_through_mode_choice": "[\"fall_through_mode_allow\",\"fall_through_mode_custom\"]"
}
```

Terraform syntax:

```terraform
fall_through_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-537f8847035f669d020afbc74e8aa713b8f9b009fc4e6cc36878852c3fd093c5"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode / 79cca19ce599 / 3

- [fall_through_mode_allow](resources--cdn_loadbalancer--reference--group-005.md#canonical-56d6c43fb70bec920d070979ae97e12ee9a027488b4d968e20747c91afe3bc06): complete subsection reference.

- [fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-005.md#canonical-775a22d75037d8a5e282340877849026cff4427aae7341c7378531e2c1273b87): complete subsection reference.

<a id="canonical-6f6d85d00a41aac6458fe614f9800d24bcc077fcfbca1e42f9a8773c2ad04944"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode / 79cca19ce599 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow](resources--cdn_loadbalancer--reference--group-005.md#canonical-56d6c43fb70bec920d070979ae97e12ee9a027488b4d968e20747c91afe3bc06)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-005.md#canonical-775a22d75037d8a5e282340877849026cff4427aae7341c7378531e2c1273b87)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-56d6c43fb70bec920d070979ae97e12ee9a027488b4d968e20747c91afe3bc06"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2b1b5517383597f756a273fce877930c37049b8006ebde6720dd9996a3216706"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / b58e4ccb5ef9 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--reference--group-005.md#canonical-5e885c4b750e99832ba179430a913c3720db297bc793085f9d1cbd82e4a0303e)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow

<a id="canonical-d4f0a4e740a78b1229b987cf719e92464ec6bb2da0e4938f10cd282c5238a6f2"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
fall_through_mode_allow = {}
```

<a id="canonical-4797565056168d9a0895c981c90884d23f965709c23c34908d650daae33d238b"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / b58e4ccb5ef9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6507766fe73bc95d3f4be614db1ac5aad290536b9fbe42dfabf6b9c71711fe54"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / b58e4ccb5ef9 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--reference--group-005.md#canonical-5e885c4b750e99832ba179430a913c3720db297bc793085f9d1cbd82e4a0303e)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-775a22d75037d8a5e282340877849026cff4427aae7341c7378531e2c1273b87"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-07664e43a137f30308ea5bb4ffa0418611ad5e64c6e9ec3ddcba9f3cda550b99"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / a5563460a511 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--reference--group-005.md#canonical-5e885c4b750e99832ba179430a913c3720db297bc793085f9d1cbd82e4a0303e)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom

<a id="canonical-50012fa20bc6fba146067e0a8c4ffb722371ae0df4533271db644aa4067c7f95"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for fall through mode custom.

Upstream description:

Define the fall through settings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("open_api_validation_rules")}
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
fall_through_mode_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-31a1ebea7978551b17c30b3e90d4a576ec289d1b4f267379fe4843d8be41dfde"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / a5563460a511 / 3

- [open_api_validation_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-1bc0f06613b481991af205b757884d7737451d6b19e2c787d48b06defb9f7f90): complete subsection reference.

<a id="canonical-d1167288f0ae972113be481592c0af6e6d31fe1a1a26deec1b123d3c81b437e7"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / a5563460a511 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-1bc0f06613b481991af205b757884d7737451d6b19e2c787d48b06defb9f7f90)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--reference--group-005.md#canonical-5e885c4b750e99832ba179430a913c3720db297bc793085f9d1cbd82e4a0303e)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-1bc0f06613b481991af205b757884d7737451d6b19e2c787d48b06defb9f7f90"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2620c76418f9b5b3892c6c4fce3e0725b48af4df13367d93565c9ac021f89b1d"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 3fdafd648d9b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--reference--group-005.md#canonical-5e885c4b750e99832ba179430a913c3720db297bc793085f9d1cbd82e4a0303e)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-005.md#canonical-775a22d75037d8a5e282340877849026cff4427aae7341c7378531e2c1273b87)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules

<a id="canonical-1410b9301379c2a064f9a2daedac026c4973a8f4dd249143a2fb1afb315fc65b"></a>

Type: `"object"`. list nested block, Optional.

Custom Fall Through Rule List. Rule or policy definition

Upstream description:

Rule or policy definition

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("action_block",
    "action_report"),
  validators.ConflictingListObjectAttributes("action_block",
    "action_skip"),
  validators.ConflictingListObjectAttributes("action_report",
    "action_skip"),
  validators.ConflictingListObjectAttributes("api_endpoint",
    "api_group"),
  validators.ConflictingListObjectAttributes("api_endpoint",
    "base_path"),
  validators.ConflictingListObjectAttributes("api_group",
    "base_path")}
```

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

Terraform syntax:

```terraform
open_api_validation_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-1172c28ed443a252439e778c571b864dc3bbfb4de00a106020512ad37ba77422"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 3fdafd648d9b / 3

- [action_block](resources--cdn_loadbalancer--reference--group-005.md#canonical-ca2684c4d2b0c6a1229caf390a6d8ca93e01956510d31f1ebb76d8908bc3e1ab): complete subsection reference.

- [action_report](resources--cdn_loadbalancer--reference--group-005.md#canonical-9e17b6238779a2f41f163892452a5a59d9375a00348c1ef25c8583b5a04f1159): complete subsection reference.

- [action_skip](resources--cdn_loadbalancer--reference--group-005.md#canonical-1f38fe2ad575eaf33977ada20771db44874f0cc4b9e4737a63e5baa2a266886e): complete subsection reference.

- [api_endpoint](resources--cdn_loadbalancer--reference--group-005.md#canonical-5357567f17fdbccb7aac4634d61787beaad165b52f752ee639732a6e85c2dd7b): complete subsection reference.

<a id="canonical-b5aba83e33451694636334730488d349270489f40f0f199c9e446c8edc962414"></a>

<a id="canonical-ef10665e455ffa8dcae0bf3f56167ca7004f01627adac0f306e1834c3ec1c4f3"></a>

## api_group property — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 3fdafd648d9b / 4

Type: `"string"`. Optional.

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

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

<a id="canonical-0911435eb1dc4d13abaf890ca55e761289b5238c3ae7499c3e8ad26b529bc489"></a>

<a id="canonical-d3ffbccc297b9de34a7c00ef38676ea187a7c680974abcc8c8c945f4f6b582a7"></a>

## base_path property — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 3fdafd648d9b / 5

Type: `"string"`. Optional.

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [metadata](resources--cdn_loadbalancer--reference--group-005.md#canonical-da2aea31cad54d6d9ebca4664f266852b2de50ffb9fba575b73abc61d630d239): complete subsection reference.

<a id="canonical-91ab0471e8e42dc016561e762f48f7897db5565bf453991f0c707ae6691a6037"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 3fdafd648d9b / 6

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block](resources--cdn_loadbalancer--reference--group-005.md#canonical-ca2684c4d2b0c6a1229caf390a6d8ca93e01956510d31f1ebb76d8908bc3e1ab)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report](resources--cdn_loadbalancer--reference--group-005.md#canonical-9e17b6238779a2f41f163892452a5a59d9375a00348c1ef25c8583b5a04f1159)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip](resources--cdn_loadbalancer--reference--group-005.md#canonical-1f38fe2ad575eaf33977ada20771db44874f0cc4b9e4737a63e5baa2a266886e)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint](resources--cdn_loadbalancer--reference--group-005.md#canonical-5357567f17fdbccb7aac4634d61787beaad165b52f752ee639732a6e85c2dd7b)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata](resources--cdn_loadbalancer--reference--group-005.md#canonical-da2aea31cad54d6d9ebca4664f266852b2de50ffb9fba575b73abc61d630d239)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-005.md#canonical-775a22d75037d8a5e282340877849026cff4427aae7341c7378531e2c1273b87)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-ca2684c4d2b0c6a1229caf390a6d8ca93e01956510d31f1ebb76d8908bc3e1ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1eb406b29658267085e2b7dfa37665de9b29bb1e23b1d96260c35170591a2da0"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / d71fffe925cb / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--reference--group-005.md#canonical-5e885c4b750e99832ba179430a913c3720db297bc793085f9d1cbd82e4a0303e)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-005.md#canonical-775a22d75037d8a5e282340877849026cff4427aae7341c7378531e2c1273b87)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-1bc0f06613b481991af205b757884d7737451d6b19e2c787d48b06defb9f7f90)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block

<a id="canonical-4830775ac5dd9a330b5b8e6b40cbc1458d2847d542eed5ed187773a39a84870d"></a>

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
action_block = {}
```

<a id="canonical-db0c0ef60765a7cdfe0b7c4c3873fe66865a86c6d142562217157dd7959d2656"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / d71fffe925cb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a8dcc5dcebb4a0f73b3b1e6f99d04e524961edd452c70c9c55abe09b49587a1d"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / d71fffe925cb / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-1bc0f06613b481991af205b757884d7737451d6b19e2c787d48b06defb9f7f90)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-9e17b6238779a2f41f163892452a5a59d9375a00348c1ef25c8583b5a04f1159"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9e8ff1474851b4bb67fa345c063026c6b35ebf9e7557e37cabe9b1e322cb82fe"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / b78d7877b805 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--reference--group-005.md#canonical-5e885c4b750e99832ba179430a913c3720db297bc793085f9d1cbd82e4a0303e)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-005.md#canonical-775a22d75037d8a5e282340877849026cff4427aae7341c7378531e2c1273b87)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-1bc0f06613b481991af205b757884d7737451d6b19e2c787d48b06defb9f7f90)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report

<a id="canonical-c9cab6fcfb16872d65afcd183c52e04cd690ee438c95e62e8ea5e090de125d55"></a>

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
action_report = {}
```

<a id="canonical-c2a6acb660451e8717ef4392c84642a229ab928e4f2462a91f77895b6e63c710"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / b78d7877b805 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8f466c3c4f53a8486044307d8a76da5c02d771c903e2b75223d4618d3dd93296"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / b78d7877b805 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-1bc0f06613b481991af205b757884d7737451d6b19e2c787d48b06defb9f7f90)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-1f38fe2ad575eaf33977ada20771db44874f0cc4b9e4737a63e5baa2a266886e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a4a5e29e0d15a9310166a1a362b73f1b9f09bb3a1bb3db5f64181700bde198fe"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 20db67a60202 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--reference--group-005.md#canonical-5e885c4b750e99832ba179430a913c3720db297bc793085f9d1cbd82e4a0303e)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-005.md#canonical-775a22d75037d8a5e282340877849026cff4427aae7341c7378531e2c1273b87)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-1bc0f06613b481991af205b757884d7737451d6b19e2c787d48b06defb9f7f90)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip

<a id="canonical-07a80d2cc24b0a64026792510cd71cde1f101b6191bf32829b6a8bbd3f39325b"></a>

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
action_skip = {}
```

<a id="canonical-7c9d1743edea031ed02c443e54b21045923e446a70220a652d544fbe02aab797"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 20db67a60202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c9f7a53b29d8320aba16cbfe7a5e118111d7f35c8223cd97bb5222f9312af81a"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 20db67a60202 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-1bc0f06613b481991af205b757884d7737451d6b19e2c787d48b06defb9f7f90)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-5357567f17fdbccb7aac4634d61787beaad165b52f752ee639732a6e85c2dd7b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-06ebf81afb821507a61f3f176ddb5558b06eeef1a278be34d59deda62f524e76"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 36141962fc1e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--reference--group-005.md#canonical-5e885c4b750e99832ba179430a913c3720db297bc793085f9d1cbd82e4a0303e)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-005.md#canonical-775a22d75037d8a5e282340877849026cff4427aae7341c7378531e2c1273b87)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-1bc0f06613b481991af205b757884d7737451d6b19e2c787d48b06defb9f7f90)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint

<a id="canonical-4d915f5f20cd40149d0ac13ea1f4ea6d35bf35a81b36877035787c9063f96700"></a>

Type: `"object"`. single nested block, Optional.

API Endpoint. This defines API endpoint.

Upstream description:

This defines API endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
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
api_endpoint {
  # Configure direct properties listed below.
}
```

<a id="canonical-64ea6ef2f8f748375609dbfe4c7232494fa3266cfba1e6f34403a8112a17fc86"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 36141962fc1e / 3

<a id="canonical-7922c1c9eb748eafbf592768f15fd4c9cbce592b5ced9d2f1d34c4c557cd0ac5"></a>

<a id="canonical-d0ecf865240820a223937995d95103cfd513f20962a8f2f83ad2d00a64a9dfa8"></a>

## methods property — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 36141962fc1e / 4

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

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

<a id="canonical-0b26f3222577552905d7d85dc4432272e8e2e22016eec4cb325e5061e753a0c7"></a>

<a id="canonical-8486006e5f80e4fd308ea3f39bfdb168e384becdeaa53e280ed2d099f431ef8a"></a>

## path property — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 36141962fc1e / 5

Type: `"string"`. Optional.

Path. Path to be matched.

Upstream description:

Path to be matched.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

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

<a id="canonical-9f427fee2239b9855019c0571e25d3120f43dc1d769c820dd5bd4aab77004545"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 36141962fc1e / 6

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-1bc0f06613b481991af205b757884d7737451d6b19e2c787d48b06defb9f7f90)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-da2aea31cad54d6d9ebca4664f266852b2de50ffb9fba575b73abc61d630d239"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-69951eb4f78f8007ce3f6bc39b4901e2a156eaf1e891bdcb38aa7824271828d9"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / b63bd553ab4a / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--reference--group-005.md#canonical-5e885c4b750e99832ba179430a913c3720db297bc793085f9d1cbd82e4a0303e)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-005.md#canonical-775a22d75037d8a5e282340877849026cff4427aae7341c7378531e2c1273b87)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-1bc0f06613b481991af205b757884d7737451d6b19e2c787d48b06defb9f7f90)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata

<a id="canonical-ca4031467bbfe3d3b2abfcf03f975fe2065f482c59486af8027608f0cf9ce8d0"></a>

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

<a id="canonical-e82d463f5f2b9fae941e70ada6784ed3b961bfdf8014eb2d80676044b7d3e97f"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / b63bd553ab4a / 3

<a id="canonical-1e77b8e493451b392cf605d8c3cdb6b819a4d9570cef3300f08970f90fc31495"></a>

<a id="canonical-19785ac32cd212f8b5de8196509b7e331fdf9ac5094bd44ad45e37cc384e6676"></a>

## description_spec property — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / b63bd553ab4a / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-49318fb27a36664fd58db9bddbec3cd26d1fb750331aadbb2f2e3814e3bca2cb"></a>

<a id="canonical-263749672fe35770a1faa7b607950f58c0433d1276f28613525f49a246f0d378"></a>

## name property — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / b63bd553ab4a / 5

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

<a id="canonical-1926843e5ff74987884395e3d1f64d68895068d33c81304c26ecd0a9db94236c"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / b63bd553ab4a / 6

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-1bc0f06613b481991af205b757884d7737451d6b19e2c787d48b06defb9f7f90)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-a66a49e08574d71974d3dce003662351a6bf6d897f2b85ae140c3b35e63b68f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-44819fb0ce3c9ee86e43daf1970a76b81142f70142047dec8aed2051fb998a0c"></a>

## api_specification.validation_all_spec_endpoints.settings — api_specification.validation_all_spec_endpoints.settings / 37231f0f6814 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf)
- api_specification.validation_all_spec_endpoints.settings

<a id="canonical-78de3ef7539235ddd0d214f239d8dfcd32d665e700892261bf77c5846b40fbb5"></a>

Type: `"object"`. single nested block, Optional.

OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom
list' enforcement.

Upstream description:

OpenAPI specification validation settings relevant for "API Inventory" enforcement and for "Custom
list" enforcement.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("oversized_body_fail_validation",
    "oversized_body_skip_validation"),
  validators.ConflictingObjectAttributes("property_validation_settings_custom",
    "property_validation_settings_default")}
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
  "x-ves-oneof-field-fail_configuration": "[]",
  "x-ves-oneof-field-oversized_body_choice": "[\"oversized_body_fail_validation\",\"oversized_body_skip_validation\"]",
  "x-ves-oneof-field-property_validation_settings_choice": "[\"property_validation_settings_custom\",\"property_validation_settings_default\"]"
}
```

Terraform syntax:

```terraform
settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-c4540b20ee8c7a7a5aaac09359de62025425899408b086e41e9f231d0c42cf10"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings / 37231f0f6814 / 3

- [oversized_body_fail_validation](resources--cdn_loadbalancer--reference--group-005.md#canonical-d529f078eb8d0a179e62b9eb192b02e7fbd1235846b613a1801cc9f84f91c5c6): complete subsection reference.

- [oversized_body_skip_validation](resources--cdn_loadbalancer--reference--group-006.md#canonical-b14e74d4a3f596f591fd0637812bebaa2f2933dba3da6f2ad5048edf88b112bd): complete subsection reference.

- [property_validation_settings_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-f764aee52452acb31cf4503aae35affbf273da93db1d51f138929d61a7a98e5b): complete subsection reference.

- [property_validation_settings_default](resources--cdn_loadbalancer--reference--group-006.md#canonical-dece8263d52bfb4349340b93842859243cb81e5034097e8d8710701237d5061d): complete subsection reference.

<a id="canonical-015586a46e4f1bbaa652cebb285b8f82e817b34f1a69e4df7c63032e238734a3"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings / 37231f0f6814 / 4

- [api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation](resources--cdn_loadbalancer--reference--group-005.md#canonical-d529f078eb8d0a179e62b9eb192b02e7fbd1235846b613a1801cc9f84f91c5c6)
- [api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation](resources--cdn_loadbalancer--reference--group-006.md#canonical-b14e74d4a3f596f591fd0637812bebaa2f2933dba3da6f2ad5048edf88b112bd)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-f764aee52452acb31cf4503aae35affbf273da93db1d51f138929d61a7a98e5b)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_default](resources--cdn_loadbalancer--reference--group-006.md#canonical-dece8263d52bfb4349340b93842859243cb81e5034097e8d8710701237d5061d)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-d529f078eb8d0a179e62b9eb192b02e7fbd1235846b613a1801cc9f84f91c5c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-55d3b054abd4ffc7b37396ea6ce243148a7491be4af9e78656408a3621323875"></a>

## api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation — api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_val / 52914c4fbbba / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf)
- [api_specification.validation_all_spec_endpoints.settings](resources--cdn_loadbalancer--reference--group-005.md#canonical-a66a49e08574d71974d3dce003662351a6bf6d897f2b85ae140c3b35e63b68f4)
- api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation

<a id="canonical-80f6d68372daf48a237f8f10cfdd4936880d541de2e033adf4dd731ce9644677"></a>

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
oversized_body_fail_validation = {}
```

<a id="canonical-9834d9497322a3d46b8b7053254131ded6969f32cfe2f39ce519ea19933d38ef"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_val / 52914c4fbbba / 3

This is an empty object or choice marker. It has no direct properties.
