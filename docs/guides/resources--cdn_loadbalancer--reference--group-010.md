---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-992081cfb9829a55250f78b266c3740cc3f24645939521747bd226760fa55e50"></a>

## ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher — ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher / fc53a7973488 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [ddos_mitigation_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-5746a5502dfc8e0895cb2a3ef2a31f60a71459706a1427f5bec63040fb069f07)
- [ddos_mitigation_rules.ddos_client_source](resources--cdn_loadbalancer--reference--group-009.md#canonical-84a504b5b14092795e6fa5369922fa6b16e4fd9a7fa887303f90b8fb04294ab8)
- ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher

<a id="canonical-52a591c5e80b5d17b42d5b844ba59eb16a65691ec4528d118acc7b94d5a045c0"></a>

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

<a id="canonical-a0fcb19256f0304a7c6875facfcae36ec8779d9b07dd8e0ff46fefd44e6a1cd7"></a>

## Direct properties — ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher / fc53a7973488 / 3

<a id="canonical-04111c5af66987d27f9050969fd3a7dc7085331ebb1cf1b4171cbf04462d915c"></a>

<a id="canonical-d7368406d69f867e30d30d1667eed0ea06f2e866bffc55684b1bfed1821cd091"></a>

## classes property — ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher / fc53a7973488 / 4

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

<a id="canonical-9626b46fcf725ed3b9f936e392565c42628a59355b75d8515765a6c32d5bd60f"></a>

<a id="canonical-e25285414d9defbcec3e18a568548d396926488988c69b3ba9315c2ec23747db"></a>

## exact_values property — ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher / fc53a7973488 / 5

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

<a id="canonical-26fefa5f076c030110e411b39c32daee33d9aa5a9dc09cde208b29863ffddc11"></a>

<a id="canonical-4e9ea66edbeef306c058e975a6c6e4b9209953c71a7ff54ce13a115b423c0eaa"></a>

## excluded_values property — ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher / fc53a7973488 / 6

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

<a id="canonical-eabcdf88d6529f91620898416630a807575db3dcad5918b434b8b596eeefbccd"></a>

## Next pages — ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher / fc53a7973488 / 7

- [ddos_mitigation_rules.ddos_client_source](resources--cdn_loadbalancer--reference--group-009.md#canonical-84a504b5b14092795e6fa5369922fa6b16e4fd9a7fa887303f90b8fb04294ab8)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-4e88f110db2623c28d213b72c36de4caca775a7f8faff5e4415068d1977e8db2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-50889c32ceeb737e516e389854853505afd33c77c0f9a08260c4d5036f5f5408"></a>

## ddos_mitigation_rules.ip_prefix_list — ddos_mitigation_rules.ip_prefix_list / 800b19e72157 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [ddos_mitigation_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-5746a5502dfc8e0895cb2a3ef2a31f60a71459706a1427f5bec63040fb069f07)
- ddos_mitigation_rules.ip_prefix_list

<a id="canonical-4f61661bccab56d9d4ff91907599086d72db3b9070a9bf4b13435e6bae548dad"></a>

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

<a id="canonical-7b9234b8ab8d09623b8fcdc2c082533cfb6d9eafa9fb03e8e5792fff1763f1a0"></a>

## Direct properties — ddos_mitigation_rules.ip_prefix_list / 800b19e72157 / 3

<a id="canonical-2b9e48d4230fbe547b921c0e7c45a36f023e5c8bd64aec706e72875b96a1647a"></a>

<a id="canonical-de876129d6095459808f467a90d21fffd4af6b5da6c8ec8ebc9984a50f2419d3"></a>

## invert_match property — ddos_mitigation_rules.ip_prefix_list / 800b19e72157 / 4

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

<a id="canonical-53973ebcc6f42d3147e05344926e76b93104dfa1ff9a257c194565bab56d1b3f"></a>

<a id="canonical-8d43d4b587e02dbc080e127ae37c534472b9686acd565c2262b364860fe5927a"></a>

## ip_prefixes property — ddos_mitigation_rules.ip_prefix_list / 800b19e72157 / 5

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

<a id="canonical-4db08fbe8c63c8268e1e9bdf25ba2f2e68628ade613867610924fa91a6c42301"></a>

## Next pages — ddos_mitigation_rules.ip_prefix_list / 800b19e72157 / 6

- [ddos_mitigation_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-5746a5502dfc8e0895cb2a3ef2a31f60a71459706a1427f5bec63040fb069f07)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-09837e2faeb7e137d6648a5128ae52925de69eb93c24b0f3acdeb1a88df11b86"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f2b46b13da93c1f43f03c2df10476fd32699fae1026940c3156fb86be885c4f9"></a>

## ddos_mitigation_rules.metadata — ddos_mitigation_rules.metadata / 774a8bd8fc5f / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [ddos_mitigation_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-5746a5502dfc8e0895cb2a3ef2a31f60a71459706a1427f5bec63040fb069f07)
- ddos_mitigation_rules.metadata

<a id="canonical-dcee75ddb1b1f8ecf3808597b86ebbcb21db27b4fb63d994c9ebccf70679b1b1"></a>

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

<a id="canonical-02adff485544cd92c1caf8dc56e395ffef96e4993e4de97013ae69687d5a819c"></a>

## Direct properties — ddos_mitigation_rules.metadata / 774a8bd8fc5f / 3

<a id="canonical-c5a8775c6d371765e580f360835b258b09f914192a12f9a182719de48296c6b8"></a>

<a id="canonical-e34a9d7763a2a2da42baaec737650aca28752392cd8795ac42cc780037b2bd1a"></a>

## description_spec property — ddos_mitigation_rules.metadata / 774a8bd8fc5f / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-9945875dac257a66a22865b20548bcb8405fdfbcc1a43a31f79bd212f8ee53f9"></a>

<a id="canonical-b17a561d14fb66e7f40f4b2ff5320f767e82599ca9adfcc1cc8a631157f046ef"></a>

## name property — ddos_mitigation_rules.metadata / 774a8bd8fc5f / 5

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

<a id="canonical-9aeb03dfac82865897fb2d0c37a3e676ec0d5e7578e47b4831b81b8c2672cc17"></a>

## Next pages — ddos_mitigation_rules.metadata / 774a8bd8fc5f / 6

- [ddos_mitigation_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-5746a5502dfc8e0895cb2a3ef2a31f60a71459706a1427f5bec63040fb069f07)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-b025aa98d3860c4c9152491bc9a93a336955de2ba2530bd0bc7f447722ac7db7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8caaf570bd9f286ff91c2120cf91d7ef14f2221c3631c7b304a48bbe50547193"></a>

## default_cache_action — default_cache_action / 230473874ab2 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- default_cache_action

<a id="canonical-8f9bf68eb65165416fda771c53d84ee44b4198e22df9b6b27306c50786b6c919"></a>

Type: `"object"`. single nested block, Optional.

Default Cache Behaviour. This defines a Default Cache Action.

Upstream description:

This defines a Default Cache Action.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("cache_disabled",
    "cache_ttl_default"),
  validators.ConflictingObjectAttributes("cache_disabled",
    "cache_ttl_override"),
  validators.ConflictingObjectAttributes("cache_ttl_default",
    "cache_ttl_override")}
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
  "x-ves-oneof-field-cache_actions": "[\"cache_disabled\",\"cache_ttl_default\",\"cache_ttl_override\"]"
}
```

Terraform syntax:

```terraform
default_cache_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-d0041841be22e0cd46fa96fb8e27ebc1ba2620b01621464116201c70291006f3"></a>

## Direct properties — default_cache_action / 230473874ab2 / 3

- [cache_disabled](resources--cdn_loadbalancer--reference--group-010.md#canonical-56c2ecb075c232ad1af4768bfd8b9e472b8c9c33bc6f67cdbc1cda4f9080dd84): complete subsection reference.

<a id="canonical-174655aed55b3d72b5f2366bdf66153a50e0e2b7c2399350e1f2acebe57de1fd"></a>

<a id="canonical-704f4b7ac9e77014a3c28c71a0316690ba9840525f5c1706764c9d4779d72f36"></a>

## cache_ttl_default property — default_cache_action / 230473874ab2 / 4

Type: `"string"`. Optional.

Exclusive with \[cache\_disabled cache\_ttl\_override\] Use Cache TTL Provided by Origin, and set a
contigency TTL value in case one is not provided.

Upstream description:

Exclusive with \[cache\_disabled cache\_ttl\_override\] Use Cache TTL Provided by Origin, and set a
contigency TTL value in case one is not provided.

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
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-e16963df86811c336e7db5ea056678ba786c762be10c1bc2bfd9d0f84670e12c"></a>

<a id="canonical-aef7a46d337a3b4c2e49b5fc322de2c7381b61b4b8101a34752a403ef05e2698"></a>

## cache_ttl_override property — default_cache_action / 230473874ab2 / 5

Type: `"string"`. Optional.

Exclusive with \[cache\_disabled cache\_ttl\_default\] Always override the Cache TTL provided by
Origin.

Upstream description:

Exclusive with \[cache\_disabled cache\_ttl\_default\] Always override the Cache TTL provided by
Origin.

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
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-67344f9eb1b8757cc4a01c51908394637c27f55ca9a7e96a6b4e49851c3083b1"></a>

## Next pages — default_cache_action / 230473874ab2 / 6

- [default_cache_action.cache_disabled](resources--cdn_loadbalancer--reference--group-010.md#canonical-56c2ecb075c232ad1af4768bfd8b9e472b8c9c33bc6f67cdbc1cda4f9080dd84)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-56c2ecb075c232ad1af4768bfd8b9e472b8c9c33bc6f67cdbc1cda4f9080dd84"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a891ede46e27d3563bac458bb02f2308b52485e98468ba1b62260c296f886d75"></a>

## default_cache_action.cache_disabled — default_cache_action.cache_disabled / 2d7be509936e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [default_cache_action](resources--cdn_loadbalancer--reference--group-010.md#canonical-b025aa98d3860c4c9152491bc9a93a336955de2ba2530bd0bc7f447722ac7db7)
- default_cache_action.cache_disabled

<a id="canonical-e5bfb4c7e8c7e28abcf71fda86a18ebded2be02e9b1640079abaafa0c984a45f"></a>

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
cache_disabled = {}
```

<a id="canonical-27e9f5f2e8e9aa961b52d60542eae7a31f0a98aaac1b5795bf6e0d4f4176470e"></a>

## Direct properties — default_cache_action.cache_disabled / 2d7be509936e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3ee3879d29f69a2109565db305e736a589b11de7489f2df86f818f4d9a569c70"></a>

## Next pages — default_cache_action.cache_disabled / 2d7be509936e / 4

- [default_cache_action](resources--cdn_loadbalancer--reference--group-010.md#canonical-b025aa98d3860c4c9152491bc9a93a336955de2ba2530bd0bc7f447722ac7db7)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-b4c6982c2159d86d5d460cc99c33456ce32b20be8c4e006cef5f974d2fbaf7bf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f9d13840e73b11e95b36e7eba75e1760d6a39b31302b0f25fd33b436155c775"></a>

## default_sensitive_data_policy — default_sensitive_data_policy / d741ba3f8ff7 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- default_sensitive_data_policy

<a id="canonical-2e856871ee6a5aa4618060e0ad6456a8642354ee55115d232d2530be716289cf"></a>

Type: `["object", {}]`. Optional.

\[OneOf: default\_sensitive\_data\_policy, sensitive\_data\_policy; Default:
default\_sensitive\_data\_policy\] Policy configuration for this feature.

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

OneOf alternatives in this subsection:

- [default_sensitive_data_policy](resources--cdn_loadbalancer--reference--group-010.md#canonical-2e856871ee6a5aa4618060e0ad6456a8642354ee55115d232d2530be716289cf)
- [sensitive_data_policy](resources--cdn_loadbalancer--reference--group-014.md#canonical-70219ffc73a735b00bea7c5242cebe9a9c3627417eb99b52705567f78aa04aec)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_sensitive_data_policy = {}
```

<a id="canonical-f3291c574a9c40990fe50097c6822008cd4ff9663abcd29b13d63d44677684c7"></a>

## Direct properties — default_sensitive_data_policy / d741ba3f8ff7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-60ff376e2f1e30ece1a6256da5d6a96d2edabd0dc4dc94bfc344900ed3af8d3d"></a>

## Next pages — default_sensitive_data_policy / d741ba3f8ff7 / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-f6b52b1f02e5223d4d1e8c0b11a436aa90f3703cbdf600691fd6fc74df939284"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b6a76516f94a3c37b067bcfdd2e875fe48e44dd5f5a2fc8071970b409c1273eb"></a>

## disable_api_definition — disable_api_definition / b7a0d8e31e56 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- disable_api_definition

<a id="canonical-47b71f79e1e4f34933aff9319652ff32324186f11ef4ffaf366520726a328f2d"></a>

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
disable_api_definition = {}
```

<a id="canonical-c07d9a48ba4b36cda4659cd6667471ac9277821ed543c8855ebe4372a54877a0"></a>

## Direct properties — disable_api_definition / b7a0d8e31e56 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bdf0e0f8596e0a2e3acc8cbf96e108aaaf72b4cb46a3d61652a6cd53a5e8160b"></a>

## Next pages — disable_api_definition / b7a0d8e31e56 / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-9c91e5fa1b5de026eca2b449600d794d171a250675f3db407f4024485d04df43"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a4d7cc16e7159bb4c746b72b341c982595f140ba626358ef70cf3131db235c4c"></a>

## disable_api_discovery — disable_api_discovery / 6cac485ee758 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- disable_api_discovery

<a id="canonical-7a399c0d7980c2912cd75b966bfb977af8cc2382b055438b317120d2d9a65065"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_api\_discovery, enable\_api\_discovery; Default: disable\_api\_discovery\] Enable
this option

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

OneOf alternatives in this subsection:

- [disable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-7a399c0d7980c2912cd75b966bfb977af8cc2382b055438b317120d2d9a65065)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-01861ffa5f27ba03bf301a9d2b1f04579fa47aa7a88ab327b31a745104b0bc94)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_api_discovery = {}
```

<a id="canonical-08a02c65057501e3fe6c532754ac4fd9ef08b9335d510698fc010fa8abb33788"></a>

## Direct properties — disable_api_discovery / 6cac485ee758 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0e7a792e56fbd1f5d08298564bba2ceeafbb64dc0e35ea4e63967cba0d0825c3"></a>

## Next pages — disable_api_discovery / 6cac485ee758 / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-3ad968aeeec75f14307d18b7058a7031d761b4b11eacc418a22217bc9e7ca4d8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a17d64bbb8aeeaf6f61cc0a7298bbcad0364653d730c67038fe76d8da6213c2f"></a>

## disable_client_side_defense — disable_client_side_defense / 580ab50e87b1 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- disable_client_side_defense

<a id="canonical-bb2a3dbac6975406cdc1e357f4ac5be5184d342f5138c0ba692f130782f07fbf"></a>

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
disable_client_side_defense = {}
```

<a id="canonical-1d1076a9fe95ca8db47640997ccca85d8fe15a0c2aa8eef2d199b7d2da135005"></a>

## Direct properties — disable_client_side_defense / 580ab50e87b1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-303d70fdf6e4f68607cf9a5197e9ecd774750090a7b112893484216a1efc119d"></a>

## Next pages — disable_client_side_defense / 580ab50e87b1 / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-4d062999b89e2ec7b48c891eea34217d6db7453eb76eb5c1ca992f27c0e6c039"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-71e1586d7f3e65fd2043225b3fc8d1cac4c02e92ade4ff46389abcae8a8b6824"></a>

## disable_ip_reputation — disable_ip_reputation / 7c81b9fc75ab / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- disable_ip_reputation

<a id="canonical-2e5b549006f65a664158a5ef4ae779562b5c476ea9fa5c436b0c14dce52e0a5b"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_ip\_reputation, enable\_ip\_reputation; Default: disable\_ip\_reputation\] Enable
this option

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

OneOf alternatives in this subsection:

- [disable_ip_reputation](resources--cdn_loadbalancer--reference--group-010.md#canonical-2e5b549006f65a664158a5ef4ae779562b5c476ea9fa5c436b0c14dce52e0a5b)
- [enable_ip_reputation](resources--cdn_loadbalancer--reference--group-010.md#canonical-5bac413884f254b3c379fc55c7e54bdee80f028c9708c3c3e95b73a2f8d5c478)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_ip_reputation = {}
```

<a id="canonical-2b3fadf14ad8aec8177325d56d4973825a81e6ae895f1adf43cf97f78e662950"></a>

## Direct properties — disable_ip_reputation / 7c81b9fc75ab / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b2fd388badc6d3a4d99f5a1a07a78196f088840ad26ce1b3a96438c829c8e7fc"></a>

## Next pages — disable_ip_reputation / 7c81b9fc75ab / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-fbcfe12bdcce7e716f099f50148c97b40fcb7c7bb76a476ee7b227537a527741"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9b5dba35e3ebd5a77fed381f496019c5070efa2018e5b0e324d59849f7a6030d"></a>

## disable_malicious_user_detection — disable_malicious_user_detection / b712e003ba83 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- disable_malicious_user_detection

<a id="canonical-bf4d32a024a6348258be4d1f2e1bc5177dabcee30f322826689e20d5471ad3d6"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_malicious\_user\_detection, enable\_malicious\_user\_detection; Default:
disable\_malicious\_user\_detection\] Configuration parameter for disable malicious user detection.

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

OneOf alternatives in this subsection:

- [disable_malicious_user_detection](resources--cdn_loadbalancer--reference--group-010.md#canonical-bf4d32a024a6348258be4d1f2e1bc5177dabcee30f322826689e20d5471ad3d6)
- [enable_malicious_user_detection](resources--cdn_loadbalancer--reference--group-010.md#canonical-5bba34c4a9e1b9d67339de478ebf9cd26663448a3eb30a6488d51abfa6fbc9f0)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_malicious_user_detection = {}
```

<a id="canonical-2b69f377a7714df2d4020f63feb18f337a5c571e28193b3f208b8bc9baa771b9"></a>

## Direct properties — disable_malicious_user_detection / b712e003ba83 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-da0bd919d462cd7fb7155a0c37fda9d667639e4130a10b8b82de5b0791e34bff"></a>

## Next pages — disable_malicious_user_detection / b712e003ba83 / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-919699123b08df44a4a2d6060b273b4e4eccbe2b03bfe046d34d4af96e8f2dc9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-252fd2fb66f13d542df8069035c14da463b88d8bff333f275d14f5ec0ddae175"></a>

## disable_rate_limit — disable_rate_limit / d8d3a26ece0e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- disable_rate_limit

<a id="canonical-df94447a3bc42e7daf142bbb5b424ea57cedd9d3c5956d6fdb32cf756cd15229"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable rate limit.

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
disable_rate_limit = {}
```

<a id="canonical-3ca08b9479bb540661a5e42bc7bdba26874234f39e68df4fe49136cd4840c4ff"></a>

## Direct properties — disable_rate_limit / d8d3a26ece0e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3b8d14555b38b02ca40a29c7752785c82fbfd639a3650b096e81c1ca1e6324cf"></a>

## Next pages — disable_rate_limit / d8d3a26ece0e / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-30068124851c4c974db1422b8003b7e1bd2902bf0a902448ee113ea0058278ed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e77d6f1856fc5d39be93f5a639257a406846b27edf7b06ab6ed12a11dde1cd0b"></a>

## disable_threat_mesh — disable_threat_mesh / 6df03ebacdc7 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- disable_threat_mesh

<a id="canonical-e66499e2b6f04f7e7c251e7bc14e58622d864b8796fe6143bc085bf17b6b5b78"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_threat\_mesh, enable\_threat\_mesh; Default: disable\_threat\_mesh\] Enable this
option

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

OneOf alternatives in this subsection:

- [disable_threat_mesh](resources--cdn_loadbalancer--reference--group-010.md#canonical-e66499e2b6f04f7e7c251e7bc14e58622d864b8796fe6143bc085bf17b6b5b78)
- [enable_threat_mesh](resources--cdn_loadbalancer--reference--group-010.md#canonical-9b9efccc5c5d1dbf2693d27cd3bfee17646a62a79935097e4d112aa7a438dac7)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_threat_mesh = {}
```

<a id="canonical-6aae59c54e5892563e1614eab3533aef9f540a5c8c2dfe5b8bc8e33921f0a924"></a>

## Direct properties — disable_threat_mesh / 6df03ebacdc7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5c2d45bd156418a0df8338e76069abd7cfef990e14e461f75b2d6a6d1494719f"></a>

## Next pages — disable_threat_mesh / 6df03ebacdc7 / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-582f1a4704469fb691c6c973eb67b1da8f6e5db5a115a88613597e82c377b485"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3aed8e714c76962b6fd7106ee81c5a5ece02bf9d396102d55d2d216667a01259"></a>

## disable_waf — disable_waf / e4470838c10d / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- disable_waf

<a id="canonical-3ce5858609fe6b57ae202e3520ad5f49645e7482910e5af32be7e23444970025"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable waf.

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
disable_waf = {}
```

<a id="canonical-bd1cc135688296786826efb596742dee5df5c672a7e5e7bb58180229be4848ba"></a>

## Direct properties — disable_waf / e4470838c10d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3f8942b0177cf4ba91aae9ebc71327c7966c0d67215e89cbf3c9a191e497e8c1"></a>

## Next pages — disable_waf / e4470838c10d / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-b3410bb7f2b0c9f28d0bb252433589a6024286ed4d766c0627313c10a630f1bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e4343a7529583f701ec710a72593a324a1a0a02ef99e36def910f9b0c527228"></a>

## enable_api_discovery — enable_api_discovery / f2cce544d267 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- enable_api_discovery

<a id="canonical-01861ffa5f27ba03bf301a9d2b1f04579fa47aa7a88ab327b31a745104b0bc94"></a>

Type: `"object"`. single nested block, Optional.

Specifies the settings used for API discovery.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_api_auth_discovery",
    "default_api_auth_discovery"),
  validators.ConflictingObjectAttributes("disable_learn_from_redirect_traffic",
    "enable_learn_from_redirect_traffic")}
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
  "x-ves-oneof-field-api_discovery_settings_choice": "[\"custom_api_auth_discovery\",\"default_api_auth_discovery\"]",
  "x-ves-oneof-field-learn_from_redirect_traffic": "[\"disable_learn_from_redirect_traffic\",\"enable_learn_from_redirect_traffic\"]"
}
```

Terraform syntax:

```terraform
enable_api_discovery {
  # Configure direct properties listed below.
}
```

<a id="canonical-cb15117cf648142342e64d6128a50255c6ba2fe46f6ce96f3de1ce0a9ab96a19"></a>

## Direct properties — enable_api_discovery / f2cce544d267 / 3

- [api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-b613b864c1ceb8611b9ffe6ae6bf22677c7678a7356a5c93dff6e464b4d8e232): complete subsection reference.

- [api_discovery_from_code_scan](resources--cdn_loadbalancer--reference--group-010.md#canonical-464c19507ba69b98c6f722aec48f6bfe7ba5d30eae8a1d2e08c865a03c960510): complete subsection reference.

- [custom_api_auth_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-bd5916ec5779220cc4c062090c1a20782a9e4ceda46b01efe4e12d0e3ecebee6): complete subsection reference.

- [default_api_auth_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-47a65a744d602f9d2e05f032c073e76e7364afb6c60b67d55667cbf9afb5a066): complete subsection reference.

- [disable_learn_from_redirect_traffic](resources--cdn_loadbalancer--reference--group-010.md#canonical-bd62c4883648eef54829ea177646bc4f31502b5bf9ae77f18fdfca7c7b4b8071): complete subsection reference.

- [discovered_api_settings](resources--cdn_loadbalancer--reference--group-010.md#canonical-963064d5662f12a4ddccd7216eaf4bc849913b3194ae10fe563e4a85d0f7ca57): complete subsection reference.

- [enable_learn_from_redirect_traffic](resources--cdn_loadbalancer--reference--group-010.md#canonical-09bdbc5a470a4c826d15b53e117497d679cec0cece46d12a007c8fa3a10d1784): complete subsection reference.

<a id="canonical-dfef974ecc863e42344bfd0bdae2c3980f38f57fdb00c98ef217b7c757c6a074"></a>

## Next pages — enable_api_discovery / f2cce544d267 / 4

- [enable_api_discovery.api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-b613b864c1ceb8611b9ffe6ae6bf22677c7678a7356a5c93dff6e464b4d8e232)
- [enable_api_discovery.api_discovery_from_code_scan](resources--cdn_loadbalancer--reference--group-010.md#canonical-464c19507ba69b98c6f722aec48f6bfe7ba5d30eae8a1d2e08c865a03c960510)
- [enable_api_discovery.custom_api_auth_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-bd5916ec5779220cc4c062090c1a20782a9e4ceda46b01efe4e12d0e3ecebee6)
- [enable_api_discovery.default_api_auth_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-47a65a744d602f9d2e05f032c073e76e7364afb6c60b67d55667cbf9afb5a066)
- [enable_api_discovery.disable_learn_from_redirect_traffic](resources--cdn_loadbalancer--reference--group-010.md#canonical-bd62c4883648eef54829ea177646bc4f31502b5bf9ae77f18fdfca7c7b4b8071)
- [enable_api_discovery.discovered_api_settings](resources--cdn_loadbalancer--reference--group-010.md#canonical-963064d5662f12a4ddccd7216eaf4bc849913b3194ae10fe563e4a85d0f7ca57)
- [enable_api_discovery.enable_learn_from_redirect_traffic](resources--cdn_loadbalancer--reference--group-010.md#canonical-09bdbc5a470a4c826d15b53e117497d679cec0cece46d12a007c8fa3a10d1784)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-b613b864c1ceb8611b9ffe6ae6bf22677c7678a7356a5c93dff6e464b4d8e232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0d7a868d0a9d653d70913cbd3fb7e3614d18509b295d3bb48190deb0bcf6007f"></a>

## enable_api_discovery.api_crawler — enable_api_discovery.api_crawler / 2f00275762a2 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-b3410bb7f2b0c9f28d0bb252433589a6024286ed4d766c0627313c10a630f1bb)
- enable_api_discovery.api_crawler

<a id="canonical-c3484e13810b3e34cd4b05b633f27422e1e5cb301d74902c5adc2e8772666a2e"></a>

Type: `"object"`. single nested block, Optional.

API Crawling. API Crawler message.

Upstream description:

API Crawler message.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("api_crawler_config",
    "disable_api_crawler")}
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
  "x-ves-oneof-field-api_crawler": "[\"api_crawler_config\",\"disable_api_crawler\"]"
}
```

Terraform syntax:

```terraform
api_crawler {
  # Configure direct properties listed below.
}
```

<a id="canonical-89c45e050cc3d7f56ad00ee98a20ab6d9b43c6ae555b7fad84df7f573359237a"></a>

## Direct properties — enable_api_discovery.api_crawler / 2f00275762a2 / 3

- [api_crawler_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-3eeeb808c4fb03b1a4d917e646ab75e25440194f97dada95a1336e37218708c9): complete subsection reference.

- [disable_api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-2bc7ad0a9693becb44360349c20fcc14753a282e18646101bd79ed8a7b710526): complete subsection reference.

<a id="canonical-c7331a6a768dffe458d3087f72a919bdb073bd0dae50c3da0e5a4927b83d2cf2"></a>

## Next pages — enable_api_discovery.api_crawler / 2f00275762a2 / 4

- [enable_api_discovery.api_crawler.api_crawler_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-3eeeb808c4fb03b1a4d917e646ab75e25440194f97dada95a1336e37218708c9)
- [enable_api_discovery.api_crawler.disable_api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-2bc7ad0a9693becb44360349c20fcc14753a282e18646101bd79ed8a7b710526)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-b3410bb7f2b0c9f28d0bb252433589a6024286ed4d766c0627313c10a630f1bb)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-3eeeb808c4fb03b1a4d917e646ab75e25440194f97dada95a1336e37218708c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3d1f710b93cba53183dcb5503b7be7bd72f6fa821355dd0f0a110457cf2c6cee"></a>

## enable_api_discovery.api_crawler.api_crawler_config — enable_api_discovery.api_crawler.api_crawler_config / 833741cc4a8a / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-b3410bb7f2b0c9f28d0bb252433589a6024286ed4d766c0627313c10a630f1bb)
- [enable_api_discovery.api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-b613b864c1ceb8611b9ffe6ae6bf22677c7678a7356a5c93dff6e464b4d8e232)
- enable_api_discovery.api_crawler.api_crawler_config

<a id="canonical-9a883e0818ff0c9b3c1c6587da2f649e3378225979d9f83907ae638719ecf482"></a>

Type: `"object"`. single nested block, Optional.

Crawler Configure.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domains")}
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
api_crawler_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-aca9110cc25bc400be44eb190e11c2fad99e1485f0c6ea04bab9acfb70b9c3a4"></a>

## Direct properties — enable_api_discovery.api_crawler.api_crawler_config / 833741cc4a8a / 3

- [domains](resources--cdn_loadbalancer--reference--group-010.md#canonical-59a345324af4811472f508b0c899abaf98d3c56106b0abc3b6a835505b6c2bbf): complete subsection reference.

<a id="canonical-8e5d07dd0988656277f95b7339f29de5a32c772a23ddf85bb1e2f2423d09160a"></a>

## Next pages — enable_api_discovery.api_crawler.api_crawler_config / 833741cc4a8a / 4

- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--cdn_loadbalancer--reference--group-010.md#canonical-59a345324af4811472f508b0c899abaf98d3c56106b0abc3b6a835505b6c2bbf)
- [enable_api_discovery.api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-b613b864c1ceb8611b9ffe6ae6bf22677c7678a7356a5c93dff6e464b4d8e232)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-59a345324af4811472f508b0c899abaf98d3c56106b0abc3b6a835505b6c2bbf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6264d8c7c80db6e375037acf8383a8dc8b7aef137ccfdae23c218c01403183e7"></a>

## enable_api_discovery.api_crawler.api_crawler_config.domains — enable_api_discovery.api_crawler.api_crawler_config.domains / bac70b1de2aa / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-b3410bb7f2b0c9f28d0bb252433589a6024286ed4d766c0627313c10a630f1bb)
- [enable_api_discovery.api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-b613b864c1ceb8611b9ffe6ae6bf22677c7678a7356a5c93dff6e464b4d8e232)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-3eeeb808c4fb03b1a4d917e646ab75e25440194f97dada95a1336e37218708c9)
- enable_api_discovery.api_crawler.api_crawler_config.domains

<a id="canonical-4f75d9e67f5b1b0d18a37adb9fe463a205e32ef2d8b0a74fca53f5441352d588"></a>

Type: `"object"`. list nested block, Optional.

Enter domains and their credentials to allow authenticated API crawling. You can only include
domains you own that are associated with this Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("domain")}
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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

Terraform syntax:

```terraform
domains {
  # Configure direct properties listed below.
}
```

<a id="canonical-c8df547d61c4f3da2c44fa1a212fa4f05e7842873a8cc16ccae6fc39c1c03be9"></a>

## Direct properties — enable_api_discovery.api_crawler.api_crawler_config.domains / bac70b1de2aa / 3

<a id="canonical-2cbf77ee4acf8dc7fc76fa4ecd26fb17b6452f9769e6f79aab98ce67f035a16c"></a>

<a id="canonical-0a77af879ee555f812242e3d351f05938c1d9220a2c515d55fcffa79200b837d"></a>

## domain property — enable_api_discovery.api_crawler.api_crawler_config.domains / bac70b1de2aa / 4

Type: `"string"`. Optional.

Select the domain to execute API Crawling with given credentials.

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
    "format": "fqdn",
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

- [simple_login](resources--cdn_loadbalancer--reference--group-010.md#canonical-a94c92a27055e59ec99629b28a1093937f5e38736e36d956159313e895fb52d5): complete subsection reference.

<a id="canonical-b43f5a39dd37d88b0b0798a04feb0ab77013f92fa88b65f28a13794ba1733098"></a>

## Next pages — enable_api_discovery.api_crawler.api_crawler_config.domains / bac70b1de2aa / 5

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--cdn_loadbalancer--reference--group-010.md#canonical-a94c92a27055e59ec99629b28a1093937f5e38736e36d956159313e895fb52d5)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-3eeeb808c4fb03b1a4d917e646ab75e25440194f97dada95a1336e37218708c9)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-a94c92a27055e59ec99629b28a1093937f5e38736e36d956159313e895fb52d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf66859328ec77119e9a317d4e05ca87a61e286190a98f7784e8e52695160e17"></a>

## enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login / e45487bd9abe / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-b3410bb7f2b0c9f28d0bb252433589a6024286ed4d766c0627313c10a630f1bb)
- [enable_api_discovery.api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-b613b864c1ceb8611b9ffe6ae6bf22677c7678a7356a5c93dff6e464b4d8e232)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-3eeeb808c4fb03b1a4d917e646ab75e25440194f97dada95a1336e37218708c9)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--cdn_loadbalancer--reference--group-010.md#canonical-59a345324af4811472f508b0c899abaf98d3c56106b0abc3b6a835505b6c2bbf)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login

<a id="canonical-329f26ecaed00c6d7008093343e139ffa561d9ebad8535c70c751eb0c8fd2cc6"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for simple login.

Receipt-pinned upstream constraints:

```json
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
simple_login {
  # Configure direct properties listed below.
}
```

<a id="canonical-b2ae68698dad36c1cc54accd3134597c11688afede5353ccb46800afdc08a93c"></a>

## Direct properties — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login / e45487bd9abe / 3

- [password](resources--cdn_loadbalancer--reference--group-010.md#canonical-5a24def46ccd76d9ebf66f032dc62d41d759b2ed599ef850c4f3f95fab1a92e8): complete subsection reference.

<a id="canonical-7b1c11c70c130728fc3f8a1b89adf50eca14b7069ed5443dc562c3a5e8708d0f"></a>

<a id="canonical-fce54ea6652e18c81605214eb5e0f7bd161f307a6376859d889270282911c846"></a>

## user property — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login / e45487bd9abe / 4

Type: `"string"`. Optional.

Enter the username to assign credentials for the selected domain to crawl.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-80cad9ddd9d8b73aba8d0110beb647a15e11624b3bff8504ad977b7aee47bf34"></a>

## Next pages — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login / e45487bd9abe / 5

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--cdn_loadbalancer--reference--group-010.md#canonical-5a24def46ccd76d9ebf66f032dc62d41d759b2ed599ef850c4f3f95fab1a92e8)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--cdn_loadbalancer--reference--group-010.md#canonical-59a345324af4811472f508b0c899abaf98d3c56106b0abc3b6a835505b6c2bbf)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-5a24def46ccd76d9ebf66f032dc62d41d759b2ed599ef850c4f3f95fab1a92e8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0dd1f946d8c82bf69d35621ae79beb98953b956218b3ac1570ba0d592c692d8d"></a>

## enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / 03599ff5339a / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-b3410bb7f2b0c9f28d0bb252433589a6024286ed4d766c0627313c10a630f1bb)
- [enable_api_discovery.api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-b613b864c1ceb8611b9ffe6ae6bf22677c7678a7356a5c93dff6e464b4d8e232)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-3eeeb808c4fb03b1a4d917e646ab75e25440194f97dada95a1336e37218708c9)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--cdn_loadbalancer--reference--group-010.md#canonical-59a345324af4811472f508b0c899abaf98d3c56106b0abc3b6a835505b6c2bbf)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--cdn_loadbalancer--reference--group-010.md#canonical-a94c92a27055e59ec99629b28a1093937f5e38736e36d956159313e895fb52d5)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password

<a id="canonical-fe44b1daa6733e6803db94e6971e583ffc12eb0b80bfc4e9fe0eef8b89f0912a"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
password {
  # Configure direct properties listed below.
}
```

<a id="canonical-43d30ad250aa9640b91c03244fed0725454690329eef39a3be96a4b2ff99cdf7"></a>

## Direct properties — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / 03599ff5339a / 3

- [blindfold_secret_info](resources--cdn_loadbalancer--reference--group-010.md#canonical-2325e9c84bb0b37801a24bcadd86b8313e90a602b7127a6de14a93656b2141a8): complete subsection reference.

- [clear_secret_info](resources--cdn_loadbalancer--reference--group-010.md#canonical-c5feabdd17fd715bc64181725646e3e78e9ff0356b6d54439d14b1f053703413): complete subsection reference.

<a id="canonical-f63b7cb6f512c3005030fbde12a4232060ff9630d999c519428de0a770d6f6fd"></a>

## Next pages — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / 03599ff5339a / 4

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info](resources--cdn_loadbalancer--reference--group-010.md#canonical-2325e9c84bb0b37801a24bcadd86b8313e90a602b7127a6de14a93656b2141a8)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info](resources--cdn_loadbalancer--reference--group-010.md#canonical-c5feabdd17fd715bc64181725646e3e78e9ff0356b6d54439d14b1f053703413)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--cdn_loadbalancer--reference--group-010.md#canonical-a94c92a27055e59ec99629b28a1093937f5e38736e36d956159313e895fb52d5)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-2325e9c84bb0b37801a24bcadd86b8313e90a602b7127a6de14a93656b2141a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd67cc01036eeb79200b73e265c11344e889def44645ed812484b5e2c63924f3"></a>

## enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / db3f7de5d910 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-b3410bb7f2b0c9f28d0bb252433589a6024286ed4d766c0627313c10a630f1bb)
- [enable_api_discovery.api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-b613b864c1ceb8611b9ffe6ae6bf22677c7678a7356a5c93dff6e464b4d8e232)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-3eeeb808c4fb03b1a4d917e646ab75e25440194f97dada95a1336e37218708c9)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--cdn_loadbalancer--reference--group-010.md#canonical-59a345324af4811472f508b0c899abaf98d3c56106b0abc3b6a835505b6c2bbf)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--cdn_loadbalancer--reference--group-010.md#canonical-a94c92a27055e59ec99629b28a1093937f5e38736e36d956159313e895fb52d5)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--cdn_loadbalancer--reference--group-010.md#canonical-5a24def46ccd76d9ebf66f032dc62d41d759b2ed599ef850c4f3f95fab1a92e8)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info

<a id="canonical-b185dd805bc5f180fd960b2512fa0f1cef4ee89439975f47a7593acd464549ff"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-18268f650cd7387d993b1331830695d9230b5a6a346f479c491214f794a2b3da"></a>

## Direct properties — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / db3f7de5d910 / 3

<a id="canonical-b27ae18e1dc809d3cddbdb6b80f4ff8b9c9a19edaad1d53298e8bab706e2fdc7"></a>

<a id="canonical-1a624f31e22ac29a804c658ad1b70c4f38eccba990401958f79c19d0461eda30"></a>

## decryption_provider property — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / db3f7de5d910 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-a6f1305ea539711e22f69ad4cb0707ab878b6bbd05c8f337bbfc747611ac43ea"></a>

<a id="canonical-fac41e47e6ef578fa7711f635f1e11fefef10ee99ae35beb41c016c83ee9e715"></a>

## location property — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / db3f7de5d910 / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-71b7cb2109548735dd9a3d14777a5219bef2c573e03f7f15d2071828c4aa2503"></a>

<a id="canonical-232dc8e3d2128e861d833fe2654f99f20af4cba1d22775b1fe34841561110e89"></a>

## store_provider property — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / db3f7de5d910 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-e0ac039d805359600848ad2d7aa3b5af8acd34462ec282a98140574ecda83d7f"></a>

## Next pages — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / db3f7de5d910 / 7

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--cdn_loadbalancer--reference--group-010.md#canonical-5a24def46ccd76d9ebf66f032dc62d41d759b2ed599ef850c4f3f95fab1a92e8)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-c5feabdd17fd715bc64181725646e3e78e9ff0356b6d54439d14b1f053703413"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e0a558b607afbc9a0e32d3186c2d472f33d58a3ce5bc7a4c640f13f2eda1b811"></a>

## enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / 4b37c04fea51 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-b3410bb7f2b0c9f28d0bb252433589a6024286ed4d766c0627313c10a630f1bb)
- [enable_api_discovery.api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-b613b864c1ceb8611b9ffe6ae6bf22677c7678a7356a5c93dff6e464b4d8e232)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-3eeeb808c4fb03b1a4d917e646ab75e25440194f97dada95a1336e37218708c9)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--cdn_loadbalancer--reference--group-010.md#canonical-59a345324af4811472f508b0c899abaf98d3c56106b0abc3b6a835505b6c2bbf)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--cdn_loadbalancer--reference--group-010.md#canonical-a94c92a27055e59ec99629b28a1093937f5e38736e36d956159313e895fb52d5)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--cdn_loadbalancer--reference--group-010.md#canonical-5a24def46ccd76d9ebf66f032dc62d41d759b2ed599ef850c4f3f95fab1a92e8)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info

<a id="canonical-9b5f7a1434cb1837e65d992fe092ed0814106825703e02f8227a498e04f03d4b"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-b8d951f6e9b1b78793bd26550669a730a5d0a4fb531da80c8f39797b98ec38d4"></a>

## Direct properties — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / 4b37c04fea51 / 3

<a id="canonical-fa0c29ed2cd3f31261dd3c37b6636994f8a95a44ad73ddc2c316dcb8bff851b3"></a>

<a id="canonical-e00dc5a72039b304698a45c5ffbe5abe63dc9bc1de0cc8f082bd76d8be65f6f7"></a>

## provider_ref property — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / 4b37c04fea51 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-7d9201abf51ed117f2563543775e3e26e31856a4ba5b4091e6e994a8fcdebe3b"></a>

<a id="canonical-531a7f82d301c4f66918b6854184ae8c4d03e9caf73a6f16333bb8d5d0895e1c"></a>

## url property — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / 4b37c04fea51 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-ccff260288b07bb356a211753bf10726595eb9b5d8f8733c47a8128dcd1e57a0"></a>

## Next pages — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / 4b37c04fea51 / 6

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--cdn_loadbalancer--reference--group-010.md#canonical-5a24def46ccd76d9ebf66f032dc62d41d759b2ed599ef850c4f3f95fab1a92e8)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-2bc7ad0a9693becb44360349c20fcc14753a282e18646101bd79ed8a7b710526"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd9edda87b85db49b07943f0454fb67d9974180788542405128c5b3f2c50d87a"></a>

## enable_api_discovery.api_crawler.disable_api_crawler — enable_api_discovery.api_crawler.disable_api_crawler / 0d41c6df7c25 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-b3410bb7f2b0c9f28d0bb252433589a6024286ed4d766c0627313c10a630f1bb)
- [enable_api_discovery.api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-b613b864c1ceb8611b9ffe6ae6bf22677c7678a7356a5c93dff6e464b4d8e232)
- enable_api_discovery.api_crawler.disable_api_crawler

<a id="canonical-df742ca9eb5754fe180eb0eb084b7c1973f9bbff3614d15c09e0d104d6371bf2"></a>

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
disable_api_crawler = {}
```

<a id="canonical-ceeb676d4cbbbcdf300b36f3620d15f78673f6c27fe380827fbe06a5aa727748"></a>

## Direct properties — enable_api_discovery.api_crawler.disable_api_crawler / 0d41c6df7c25 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bad197b0b1e3926a61ed3df03a90135a845916ca05d933d6beb841498fed38b4"></a>

## Next pages — enable_api_discovery.api_crawler.disable_api_crawler / 0d41c6df7c25 / 4

- [enable_api_discovery.api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-b613b864c1ceb8611b9ffe6ae6bf22677c7678a7356a5c93dff6e464b4d8e232)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-464c19507ba69b98c6f722aec48f6bfe7ba5d30eae8a1d2e08c865a03c960510"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b26786943e002cce8176049ac4a54612321c23db0d5774fb3c07b657bf571eb4"></a>

## enable_api_discovery.api_discovery_from_code_scan — enable_api_discovery.api_discovery_from_code_scan / bdf2a9ad7c56 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-b3410bb7f2b0c9f28d0bb252433589a6024286ed4d766c0627313c10a630f1bb)
- enable_api_discovery.api_discovery_from_code_scan

<a id="canonical-46c35b34016432be2b32a9fb4cc22cf570f91665db8d796da987069c7520c418"></a>

Type: `"object"`. single nested block, Optional.

Select Code Base and Repositories.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("code_base_integrations")}
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
api_discovery_from_code_scan {
  # Configure direct properties listed below.
}
```

<a id="canonical-82a8b4205748ada07fa0c20da5f877c3dbd0a85eb93bea49fbf8eae8f1fa4f97"></a>

## Direct properties — enable_api_discovery.api_discovery_from_code_scan / bdf2a9ad7c56 / 3

- [code_base_integrations](resources--cdn_loadbalancer--reference--group-010.md#canonical-4925c1feae3fe68b2c00b30968d6bd1e5f4b3653197f0493a03eedf639f5774d): complete subsection reference.

<a id="canonical-9ef6691f6c1343d105abd65cb4feb3b9e0c9260ecad8392d1f2b630aec223143"></a>

## Next pages — enable_api_discovery.api_discovery_from_code_scan / bdf2a9ad7c56 / 4

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--cdn_loadbalancer--reference--group-010.md#canonical-4925c1feae3fe68b2c00b30968d6bd1e5f4b3653197f0493a03eedf639f5774d)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-b3410bb7f2b0c9f28d0bb252433589a6024286ed4d766c0627313c10a630f1bb)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-4925c1feae3fe68b2c00b30968d6bd1e5f4b3653197f0493a03eedf639f5774d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7f1e6c329ee5a83f50dd4fe9eedd4491bc13f53cfb4498e342f1eb5be9d0ddd5"></a>

## enable_api_discovery.api_discovery_from_code_scan.code_base_integrations — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations / f174e1da2051 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-b3410bb7f2b0c9f28d0bb252433589a6024286ed4d766c0627313c10a630f1bb)
- [enable_api_discovery.api_discovery_from_code_scan](resources--cdn_loadbalancer--reference--group-010.md#canonical-464c19507ba69b98c6f722aec48f6bfe7ba5d30eae8a1d2e08c865a03c960510)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations

<a id="canonical-5f84d483dd05297c947e38a6b4cb9ba8cc163533976641f158593fb9cd2f61ab"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for code base integrations.

Upstream description:

Configuration parameter for code base integrations

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("all_repos",
    "selected_repos")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
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
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
code_base_integrations {
  # Configure direct properties listed below.
}
```

<a id="canonical-e7885c7ea7cb9d9c09e98d04ef204afa14d5a43cb18b55525ebb651048b4c3ad"></a>

## Direct properties — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations / f174e1da2051 / 3

- [all_repos](resources--cdn_loadbalancer--reference--group-010.md#canonical-5fab6befec0e2b1e3d22dc7347f09ac983fa88db7f9cfb799116f349710e77ae): complete subsection reference.

- [code_base_integration](resources--cdn_loadbalancer--reference--group-010.md#canonical-996e73cc117b12ea24e535cedcdde2d94be77d1d5a14d0699a80bd6943eae92b): complete subsection reference.

- [selected_repos](resources--cdn_loadbalancer--reference--group-010.md#canonical-fe0fc22f798d1d564ee2a5a6336f32b81cbef9a6bd6a4290866fb718b09bf7fd): complete subsection reference.

<a id="canonical-575de0347ac08e98b9568ff8efc2188537aef96b4891cee3ed9675fa18cb7ca0"></a>

## Next pages — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations / f174e1da2051 / 4

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos](resources--cdn_loadbalancer--reference--group-010.md#canonical-5fab6befec0e2b1e3d22dc7347f09ac983fa88db7f9cfb799116f349710e77ae)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration](resources--cdn_loadbalancer--reference--group-010.md#canonical-996e73cc117b12ea24e535cedcdde2d94be77d1d5a14d0699a80bd6943eae92b)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos](resources--cdn_loadbalancer--reference--group-010.md#canonical-fe0fc22f798d1d564ee2a5a6336f32b81cbef9a6bd6a4290866fb718b09bf7fd)
- [enable_api_discovery.api_discovery_from_code_scan](resources--cdn_loadbalancer--reference--group-010.md#canonical-464c19507ba69b98c6f722aec48f6bfe7ba5d30eae8a1d2e08c865a03c960510)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-5fab6befec0e2b1e3d22dc7347f09ac983fa88db7f9cfb799116f349710e77ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a8e4c2d31dc4d2d71813ce2d7884d0757b0e64fd8966705145cf7b1de2e4d92a"></a>

## enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_rep / 0fbe2aafab27 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-b3410bb7f2b0c9f28d0bb252433589a6024286ed4d766c0627313c10a630f1bb)
- [enable_api_discovery.api_discovery_from_code_scan](resources--cdn_loadbalancer--reference--group-010.md#canonical-464c19507ba69b98c6f722aec48f6bfe7ba5d30eae8a1d2e08c865a03c960510)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--cdn_loadbalancer--reference--group-010.md#canonical-4925c1feae3fe68b2c00b30968d6bd1e5f4b3653197f0493a03eedf639f5774d)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos

<a id="canonical-40f47f2ac05d35e501944c704f1d4adbe9622fc69e0b15cc4101ccfa38c18bd9"></a>

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
all_repos = {}
```

<a id="canonical-6055afd99ba484c6217312442406cd07bfb6dcdba8ba3a0b3d4e658b948f4da3"></a>

## Direct properties — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_rep / 0fbe2aafab27 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4a65b4343d7f83d19a65ee3df3852b723a99d356f84628e3850cd91d16a1e1e9"></a>

## Next pages — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_rep / 0fbe2aafab27 / 4

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--cdn_loadbalancer--reference--group-010.md#canonical-4925c1feae3fe68b2c00b30968d6bd1e5f4b3653197f0493a03eedf639f5774d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-996e73cc117b12ea24e535cedcdde2d94be77d1d5a14d0699a80bd6943eae92b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5527c27e083981f7b4b774687912487300e71d7ed2685fb30bc51d8411d7b25d"></a>

## enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_ba / 3efbd33dc618 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-b3410bb7f2b0c9f28d0bb252433589a6024286ed4d766c0627313c10a630f1bb)
- [enable_api_discovery.api_discovery_from_code_scan](resources--cdn_loadbalancer--reference--group-010.md#canonical-464c19507ba69b98c6f722aec48f6bfe7ba5d30eae8a1d2e08c865a03c960510)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--cdn_loadbalancer--reference--group-010.md#canonical-4925c1feae3fe68b2c00b30968d6bd1e5f4b3653197f0493a03eedf639f5774d)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration

<a id="canonical-e02b91c4a673ed76de6d838bed013c429988a4c2cc71fd02f42e2d053c5a4e4e"></a>

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
code_base_integration {
  # Configure direct properties listed below.
}
```

<a id="canonical-5525a7d870422c055bc3176602529d53cc8d5f64091d1bbf9bf2b49e97d6a967"></a>

## Direct properties — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_ba / 3efbd33dc618 / 3

<a id="canonical-a85b71f85e40edf2402c78ca9520405c4771f519d170f6e00834e7db67585a08"></a>

<a id="canonical-0ed1a216bc2f17e499fed5552ab6b5c2502436620772dc702fcbe2973e9a3d09"></a>

## name property — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_ba / 3efbd33dc618 / 4

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

<a id="canonical-8dd9e69b99327c17629d5838ec5b9e91a1340979480f9d9fefcd50fee43fa16a"></a>

<a id="canonical-a74061668b6e9dd0b48415b536ead7e506c32ebeae7298f93b3662ec6dc5b1fc"></a>

## namespace property — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_ba / 3efbd33dc618 / 5

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

<a id="canonical-706ac562aee9138d8636e0ade644bcc36d45dfb8b3c45b69351e91eeaa5343de"></a>

<a id="canonical-42be8dc723fa98086019b80dacb4f2ff78a9c0b25a4c818bfacd3deeb220c91c"></a>

## tenant property — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_ba / 3efbd33dc618 / 6

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

<a id="canonical-20bbbe2b0f8a00ac510c2123a24ce39949d30c30596de0e66c2878f28d5385e7"></a>

## Next pages — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_ba / 3efbd33dc618 / 7

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--cdn_loadbalancer--reference--group-010.md#canonical-4925c1feae3fe68b2c00b30968d6bd1e5f4b3653197f0493a03eedf639f5774d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-fe0fc22f798d1d564ee2a5a6336f32b81cbef9a6bd6a4290866fb718b09bf7fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c0707ea6ae31131e7f03be95510c9af63d8580fed5bffd93d924640e4694423"></a>

## enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selecte / 9217ab71efb5 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-b3410bb7f2b0c9f28d0bb252433589a6024286ed4d766c0627313c10a630f1bb)
- [enable_api_discovery.api_discovery_from_code_scan](resources--cdn_loadbalancer--reference--group-010.md#canonical-464c19507ba69b98c6f722aec48f6bfe7ba5d30eae8a1d2e08c865a03c960510)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--cdn_loadbalancer--reference--group-010.md#canonical-4925c1feae3fe68b2c00b30968d6bd1e5f4b3653197f0493a03eedf639f5774d)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos

<a id="canonical-8bb9f1d0a096e9c0a1168cdb787d335bd2599576a82d3d12d0df73a4511c421e"></a>

Type: `"object"`. single nested block, Optional.

Select which API repositories represent the LB applications.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("api_code_repo")}
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
selected_repos {
  # Configure direct properties listed below.
}
```

<a id="canonical-2de89c409facab088a213dac58a61f537022bfc9e339ab608e31e3af3993c89f"></a>

## Direct properties — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selecte / 9217ab71efb5 / 3

<a id="canonical-330c521aff9a803a9f7ca7ef697e4db40a53bf33ad853072bfd7de698a34e4a9"></a>

<a id="canonical-1cc04a0929db897e1270830a7a80074f0f9e4c46c50ac897a9308170fcf0357a"></a>

## api_code_repo property — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selecte / 9217ab71efb5 / 4

Type: `["list", "string"]`. Optional.

Code repository which contain API endpoints.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-5d372e8446231f636eade3ef3522e6f7aff893ed1b34d482649ead9373f4d494"></a>

## Next pages — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selecte / 9217ab71efb5 / 5

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--cdn_loadbalancer--reference--group-010.md#canonical-4925c1feae3fe68b2c00b30968d6bd1e5f4b3653197f0493a03eedf639f5774d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-bd5916ec5779220cc4c062090c1a20782a9e4ceda46b01efe4e12d0e3ecebee6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2da69c9325882279bb2547fbb32de0488a8bda5efa16f26665306cf89222d5a4"></a>

## enable_api_discovery.custom_api_auth_discovery — enable_api_discovery.custom_api_auth_discovery / 3d4ad58afdb5 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-b3410bb7f2b0c9f28d0bb252433589a6024286ed4d766c0627313c10a630f1bb)
- enable_api_discovery.custom_api_auth_discovery

<a id="canonical-80006aec25100bf7441fce37c30b3cc7c31aba27fecadc1fb1a2c5eb179e6a68"></a>

Type: `"object"`. single nested block, Optional.

API Discovery Advanced Settings. API Discovery Advanced settings.

Upstream description:

API Discovery Advanced settings.

Receipt-pinned upstream constraints:

```json
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
custom_api_auth_discovery {
  # Configure direct properties listed below.
}
```

<a id="canonical-12cbc9925281b6d8c03f36941474b932f3007248e745108cfc1c924df3c9e673"></a>

## Direct properties — enable_api_discovery.custom_api_auth_discovery / 3d4ad58afdb5 / 3

- [api_discovery_ref](resources--cdn_loadbalancer--reference--group-010.md#canonical-4eec21084af63db12b8c730243bdcde533dc7af22ae93cdb264c1b0159e16bbc): complete subsection reference.

<a id="canonical-4c84e6ce60dd942727694de7980c2a7c64e83585ec72d3ff4bcded45f9086968"></a>

## Next pages — enable_api_discovery.custom_api_auth_discovery / 3d4ad58afdb5 / 4

- [enable_api_discovery.custom_api_auth_discovery.api_discovery_ref](resources--cdn_loadbalancer--reference--group-010.md#canonical-4eec21084af63db12b8c730243bdcde533dc7af22ae93cdb264c1b0159e16bbc)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-b3410bb7f2b0c9f28d0bb252433589a6024286ed4d766c0627313c10a630f1bb)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-4eec21084af63db12b8c730243bdcde533dc7af22ae93cdb264c1b0159e16bbc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e646a3d497f12fbad1fa95f2c8799f777c84c65ca96c36ad332cd48f2765f448"></a>

## enable_api_discovery.custom_api_auth_discovery.api_discovery_ref — enable_api_discovery.custom_api_auth_discovery.api_discovery_ref / e08d83238b18 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-b3410bb7f2b0c9f28d0bb252433589a6024286ed4d766c0627313c10a630f1bb)
- [enable_api_discovery.custom_api_auth_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-bd5916ec5779220cc4c062090c1a20782a9e4ceda46b01efe4e12d0e3ecebee6)
- enable_api_discovery.custom_api_auth_discovery.api_discovery_ref

<a id="canonical-90ae550e0abafb52c207791f98733f54c0e04df1979b2e86146c001e076e25c3"></a>

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
api_discovery_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-798ed7e84403ecdada3c2cf89a7a6aa917d411468aa36510207829e2593a2c38"></a>

## Direct properties — enable_api_discovery.custom_api_auth_discovery.api_discovery_ref / e08d83238b18 / 3

<a id="canonical-621fe39775a534a237aa14142861669f644e9585d96671fde879bc984286e1f6"></a>

<a id="canonical-d352b601e8682fa0c0b2d8a7352b083785b507edd48ef025f741df4277366667"></a>

## name property — enable_api_discovery.custom_api_auth_discovery.api_discovery_ref / e08d83238b18 / 4

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

<a id="canonical-25dcbe07325ac0950311c9e76f0624ec3bca9eeaf678a14183f94c4990953d74"></a>

<a id="canonical-4f984350f36aae3a116ce8cd60722937302b9ac5025cc3290f6025ea5ffa26fb"></a>

## namespace property — enable_api_discovery.custom_api_auth_discovery.api_discovery_ref / e08d83238b18 / 5

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

<a id="canonical-6cbd076544a363ffb63b9bee7319ba413f2688dcff1efbb94a417ead12bf2b55"></a>

<a id="canonical-c598be8be99b9a9fee273379164de987f6aff861b91d8b3c3a3b5261e4100a1b"></a>

## tenant property — enable_api_discovery.custom_api_auth_discovery.api_discovery_ref / e08d83238b18 / 6

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

<a id="canonical-3a260ecdf6c31cdc9afbdb0d57b9b15bb5bd49347d4e94313422e1c3dba0d83e"></a>

## Next pages — enable_api_discovery.custom_api_auth_discovery.api_discovery_ref / e08d83238b18 / 7

- [enable_api_discovery.custom_api_auth_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-bd5916ec5779220cc4c062090c1a20782a9e4ceda46b01efe4e12d0e3ecebee6)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-47a65a744d602f9d2e05f032c073e76e7364afb6c60b67d55667cbf9afb5a066"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2a7c141fdd2a88961c93d4e5aa658c2c94411da4bc7d814f47a1581a03700ecf"></a>

## enable_api_discovery.default_api_auth_discovery — enable_api_discovery.default_api_auth_discovery / 4d65b3c0aec1 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-b3410bb7f2b0c9f28d0bb252433589a6024286ed4d766c0627313c10a630f1bb)
- enable_api_discovery.default_api_auth_discovery

<a id="canonical-fcd8f46c3eaa0bfe113b67190db2a71d0b80086c5363313378dea08e4865bddb"></a>

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
default_api_auth_discovery = {}
```

<a id="canonical-ba4cb0df76dde423928562194b76df4235150d78ded81e68b71f9a7fc93c0365"></a>

## Direct properties — enable_api_discovery.default_api_auth_discovery / 4d65b3c0aec1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0921362a3c7b140c6f9ec9425805e8747c0fb44cafdf9d7920c8ee8cb9d5d166"></a>

## Next pages — enable_api_discovery.default_api_auth_discovery / 4d65b3c0aec1 / 4

- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-b3410bb7f2b0c9f28d0bb252433589a6024286ed4d766c0627313c10a630f1bb)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-bd62c4883648eef54829ea177646bc4f31502b5bf9ae77f18fdfca7c7b4b8071"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-246ac313c9ae73f6401147b015edd15b866d3ace241c6a90860359f657c63d67"></a>

## enable_api_discovery.disable_learn_from_redirect_traffic — enable_api_discovery.disable_learn_from_redirect_traffic / 991eb6076497 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-b3410bb7f2b0c9f28d0bb252433589a6024286ed4d766c0627313c10a630f1bb)
- enable_api_discovery.disable_learn_from_redirect_traffic

<a id="canonical-47b4aedc1f282436a00be0c80ea84dcfe4c0b4ccb6e0d66e01af7c460fa1aa69"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable learn from redirect traffic.

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
disable_learn_from_redirect_traffic = {}
```

<a id="canonical-beb6887d79916cdc315bf6f38fb3b6795badba3cc34f924bbdbc51332fc5ffae"></a>

## Direct properties — enable_api_discovery.disable_learn_from_redirect_traffic / 991eb6076497 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-de8bd13f0c29f6a0cf3f6ad294aa6ad93477377a7bda694dafd72a2ad3bf57ca"></a>

## Next pages — enable_api_discovery.disable_learn_from_redirect_traffic / 991eb6076497 / 4

- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-b3410bb7f2b0c9f28d0bb252433589a6024286ed4d766c0627313c10a630f1bb)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-963064d5662f12a4ddccd7216eaf4bc849913b3194ae10fe563e4a85d0f7ca57"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e9bb259d30c2ce27a7d61f35ac38f354564bb7646afe812b8e2bec428759d58b"></a>

## enable_api_discovery.discovered_api_settings — enable_api_discovery.discovered_api_settings / 9ddcc487efc6 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-b3410bb7f2b0c9f28d0bb252433589a6024286ed4d766c0627313c10a630f1bb)
- enable_api_discovery.discovered_api_settings

<a id="canonical-e5c9b9efa74cbbb67ab21acdd77295fbe728688fd30a72881498bf2b5d7e38e6"></a>

Type: `"object"`. single nested block, Optional.

Discovered API Settings. Configure Discovered API Settings.

Upstream description:

Configure Discovered API Settings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("purge_duration_for_inactive_discovered_apis")}
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
discovered_api_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-44334e6b3c3ad617ccbd2c91b36f2c37365d4472fbf81e146321b5d2aefe9b07"></a>

## Direct properties — enable_api_discovery.discovered_api_settings / 9ddcc487efc6 / 3

<a id="canonical-001ed29a8c09f1f9f4a25d80a075e412c327469702687eef725d9a0d54a359d8"></a>

<a id="canonical-31214a8264863da7f0ddee60891813ba3c573a959d000ed29feabfe7d1486c75"></a>

## purge_duration_for_inactive_discovered_apis property — enable_api_discovery.discovered_api_settings / 9ddcc487efc6 / 4

Type: `"number"`. Optional.

Inactive discovered API will be deleted after configured duration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 7),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 7,
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
    "ves.io.schema.rules.uint32.lte": "7"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "7"
  }
}
```

<a id="canonical-e8fafcd91ab35a48b35097c2f9f1c92e46a51f4bb2862f695ad4b6bd7372d91a"></a>

## Next pages — enable_api_discovery.discovered_api_settings / 9ddcc487efc6 / 5

- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-b3410bb7f2b0c9f28d0bb252433589a6024286ed4d766c0627313c10a630f1bb)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-09bdbc5a470a4c826d15b53e117497d679cec0cece46d12a007c8fa3a10d1784"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-516f8cf7aed2b752ef1741d185f1e4b99cc584ac1c3d8733e023900df7fac6b9"></a>

## enable_api_discovery.enable_learn_from_redirect_traffic — enable_api_discovery.enable_learn_from_redirect_traffic / cbee7bbdc626 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-b3410bb7f2b0c9f28d0bb252433589a6024286ed4d766c0627313c10a630f1bb)
- enable_api_discovery.enable_learn_from_redirect_traffic

<a id="canonical-7c1ba928ad441281ef35b2a42c8871c5acfc50ade8ff236242e162f86de7516a"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable learn from redirect traffic.

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
enable_learn_from_redirect_traffic = {}
```

<a id="canonical-8e8269fb694951a017dce0c32d728e61ab62b138e7cc6017b670a76aed9c40bc"></a>

## Direct properties — enable_api_discovery.enable_learn_from_redirect_traffic / cbee7bbdc626 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a5a9a292f99652dc8329381434b7a4fd8c272f119e95c8e219f54825cdd3d58f"></a>

## Next pages — enable_api_discovery.enable_learn_from_redirect_traffic / cbee7bbdc626 / 4

- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-b3410bb7f2b0c9f28d0bb252433589a6024286ed4d766c0627313c10a630f1bb)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-be72efa8da41d7b196abe830ae2dc8afde3c2becf8166936a3708d45c8cf0a73"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b6e10d689954bac4aa6cbb40edb21fee80e470f6da67677fb54740face813897"></a>

## enable_challenge — enable_challenge / c0dd50298a49 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- enable_challenge

<a id="canonical-61db621bec5658e33b4d89b4822b6deea0e32dd5026aeca0af822d6d26be4067"></a>

Type: `"object"`. single nested block, Optional.

Configure auto mitigation i.e risk based challenges for malicious users.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("captcha_challenge_parameters",
    "default_captcha_challenge_parameters"),
  validators.ConflictingObjectAttributes("default_js_challenge_parameters",
    "js_challenge_parameters"),
  validators.ConflictingObjectAttributes("default_mitigation_settings",
    "malicious_user_mitigation")}
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
  "x-ves-oneof-field-captcha_challenge_parameters_choice": "[\"captcha_challenge_parameters\",\"default_captcha_challenge_parameters\"]",
  "x-ves-oneof-field-js_challenge_parameters_choice": "[\"default_js_challenge_parameters\",\"js_challenge_parameters\"]",
  "x-ves-oneof-field-malicious_user_mitigation_choice": "[\"default_mitigation_settings\",\"malicious_user_mitigation\"]"
}
```

Terraform syntax:

```terraform
enable_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-aa239be28121d5ef6b78d39fc606538cc909027dfa4a03a69b25cff99a43546d"></a>

## Direct properties — enable_challenge / c0dd50298a49 / 3

- [captcha_challenge_parameters](resources--cdn_loadbalancer--reference--group-010.md#canonical-465b3c54eaaca5593aef9fde2fb39c41c562c4c5796f03fd07dab23d12820b46): complete subsection reference.

- [default_captcha_challenge_parameters](resources--cdn_loadbalancer--reference--group-010.md#canonical-a66184f1ef3647c104fa08fccf06a627d67a0d977a42178a200b27fa3fcabf26): complete subsection reference.

- [default_js_challenge_parameters](resources--cdn_loadbalancer--reference--group-010.md#canonical-59537006f4e52820d74251bbc37b5d266ca5e185429567f10df055a3329f67d4): complete subsection reference.

- [default_mitigation_settings](resources--cdn_loadbalancer--reference--group-010.md#canonical-c36c46f0f141afb6c215ea08bf5415106641c28da79373a4190da13ffa88fdc3): complete subsection reference.

- [js_challenge_parameters](resources--cdn_loadbalancer--reference--group-010.md#canonical-36446b94011650fb4c11455734cbf98615e7a548c92b14a62eb0ce193fbad2da): complete subsection reference.

- [malicious_user_mitigation](resources--cdn_loadbalancer--reference--group-010.md#canonical-3bc129571f61df9ce982c8ad2e87f4ba5d0d632d27c3c5ea9368f14b576ac8a3): complete subsection reference.

<a id="canonical-86e1ef3275ad2838f4b49ea9f3d9a8053009fc1f9ab8bf97917ebcc52a1051d9"></a>

## Next pages — enable_challenge / c0dd50298a49 / 4

- [enable_challenge.captcha_challenge_parameters](resources--cdn_loadbalancer--reference--group-010.md#canonical-465b3c54eaaca5593aef9fde2fb39c41c562c4c5796f03fd07dab23d12820b46)
- [enable_challenge.default_captcha_challenge_parameters](resources--cdn_loadbalancer--reference--group-010.md#canonical-a66184f1ef3647c104fa08fccf06a627d67a0d977a42178a200b27fa3fcabf26)
- [enable_challenge.default_js_challenge_parameters](resources--cdn_loadbalancer--reference--group-010.md#canonical-59537006f4e52820d74251bbc37b5d266ca5e185429567f10df055a3329f67d4)
- [enable_challenge.default_mitigation_settings](resources--cdn_loadbalancer--reference--group-010.md#canonical-c36c46f0f141afb6c215ea08bf5415106641c28da79373a4190da13ffa88fdc3)
- [enable_challenge.js_challenge_parameters](resources--cdn_loadbalancer--reference--group-010.md#canonical-36446b94011650fb4c11455734cbf98615e7a548c92b14a62eb0ce193fbad2da)
- [enable_challenge.malicious_user_mitigation](resources--cdn_loadbalancer--reference--group-010.md#canonical-3bc129571f61df9ce982c8ad2e87f4ba5d0d632d27c3c5ea9368f14b576ac8a3)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-465b3c54eaaca5593aef9fde2fb39c41c562c4c5796f03fd07dab23d12820b46"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-caf7983d91fd5f93e5618915d721d55148dbef5c675db0e7da256af6d90b0506"></a>

## enable_challenge.captcha_challenge_parameters — enable_challenge.captcha_challenge_parameters / 775daef5814f / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-be72efa8da41d7b196abe830ae2dc8afde3c2becf8166936a3708d45c8cf0a73)
- enable_challenge.captcha_challenge_parameters

<a id="canonical-256f5b4803598d23a81dd51409fd87846e8af3275d83195cc18244563b70d233"></a>

Type: `"object"`. single nested block, Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry")}
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
captcha_challenge_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-96030270cfaeb849543877d1ab2c8f6e5dea5bc8a948334cfb7903782e065d47"></a>

## Direct properties — enable_challenge.captcha_challenge_parameters / 775daef5814f / 3

<a id="canonical-17bbdffdf1859544e4ad9acbd55d2ff4370a56220f708459b8465ed815af6bb1"></a>

<a id="canonical-ac20884274ddf48142e83ac19a905abf92acb023b1caf09560c8d2287dacbc15"></a>

## cookie_expiry property — enable_challenge.captcha_challenge_parameters / 775daef5814f / 4

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
}
```

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

<a id="canonical-f5748de79ebf8d22b44ab47dbeb583824791bd495f83877c6781fbb3051be042"></a>

<a id="canonical-876876103ef2a2099b7d1cf286e4f79ded9c837c09faa84e872fecf317ed2aa2"></a>

## custom_page property — enable_challenge.captcha_challenge_parameters / 775daef5814f / 5

Type: `"string"`. Optional.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

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

<a id="canonical-bd457eab03d832dd6ddb2d92de29ad30149c0ed9e20d4c43182bb9bb021681d7"></a>

## Next pages — enable_challenge.captcha_challenge_parameters / 775daef5814f / 6

- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-be72efa8da41d7b196abe830ae2dc8afde3c2becf8166936a3708d45c8cf0a73)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-a66184f1ef3647c104fa08fccf06a627d67a0d977a42178a200b27fa3fcabf26"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-de9a5388f0a0522f5a27fd039373ae8ac1e2d10a232689f9bdb9713cb2b233da"></a>

## enable_challenge.default_captcha_challenge_parameters — enable_challenge.default_captcha_challenge_parameters / b0b63921b325 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-be72efa8da41d7b196abe830ae2dc8afde3c2becf8166936a3708d45c8cf0a73)
- enable_challenge.default_captcha_challenge_parameters

<a id="canonical-ff5ce4ce4131f6af0079f5f5d8a158e11b950bf6f59976a8a12530a789fce8a8"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_captcha_challenge_parameters = {}
```

<a id="canonical-06c0a046648465c940bd06c5aab370391a753b7adf059ea3aea4b919ec3417a3"></a>

## Direct properties — enable_challenge.default_captcha_challenge_parameters / b0b63921b325 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-74776f6a8b8f26c408f4d0d6142762b2bef670e07160b8ca83d4f5e60de4c465"></a>

## Next pages — enable_challenge.default_captcha_challenge_parameters / b0b63921b325 / 4

- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-be72efa8da41d7b196abe830ae2dc8afde3c2becf8166936a3708d45c8cf0a73)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-59537006f4e52820d74251bbc37b5d266ca5e185429567f10df055a3329f67d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-81eadb263bbf2e952d57600ae3bdca6083e2026fbb22b08aea4090c667e7633d"></a>

## enable_challenge.default_js_challenge_parameters — enable_challenge.default_js_challenge_parameters / 9bf4788c5773 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-be72efa8da41d7b196abe830ae2dc8afde3c2becf8166936a3708d45c8cf0a73)
- enable_challenge.default_js_challenge_parameters

<a id="canonical-64f530af6c8a6afa2eb78e73c49a281cfef44a57455a93d0c444e390051ce8c1"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_js_challenge_parameters = {}
```

<a id="canonical-4a3cdf3859194d9e5a99af9570d34cca0b76a52292997943791a99f7a4dd597e"></a>

## Direct properties — enable_challenge.default_js_challenge_parameters / 9bf4788c5773 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e1b7b9e187cf5577ba41ba5ec7772ac688c998fbd90fa25add25f8b28bfb7265"></a>

## Next pages — enable_challenge.default_js_challenge_parameters / 9bf4788c5773 / 4

- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-be72efa8da41d7b196abe830ae2dc8afde3c2becf8166936a3708d45c8cf0a73)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-c36c46f0f141afb6c215ea08bf5415106641c28da79373a4190da13ffa88fdc3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-49b16671c777d06e8be0fdc15844ed85ccea96dcbf4baf9022fa125facd4a653"></a>

## enable_challenge.default_mitigation_settings — enable_challenge.default_mitigation_settings / 607ed145a838 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-be72efa8da41d7b196abe830ae2dc8afde3c2becf8166936a3708d45c8cf0a73)
- enable_challenge.default_mitigation_settings

<a id="canonical-30a6772b1bad90e13ad2ef992266c806b9fb07b82bc8dcc5abb04a6fc5f2a63c"></a>

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
default_mitigation_settings = {}
```

<a id="canonical-a794c43dee22ca8a958f92fab92bdae74afde72a411f06ee1136637b5939a666"></a>

## Direct properties — enable_challenge.default_mitigation_settings / 607ed145a838 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7e8c84612be9608267f6808be03b8211f9cd03ef64e91a5dac923f1fd70a9d21"></a>

## Next pages — enable_challenge.default_mitigation_settings / 607ed145a838 / 4

- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-be72efa8da41d7b196abe830ae2dc8afde3c2becf8166936a3708d45c8cf0a73)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-36446b94011650fb4c11455734cbf98615e7a548c92b14a62eb0ce193fbad2da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-624f62d1bc37950dbd00cd0b2fbf7e08d29fc0328182440ac19aa02c3c359a14"></a>

## enable_challenge.js_challenge_parameters — enable_challenge.js_challenge_parameters / b248fb0b3c4a / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-be72efa8da41d7b196abe830ae2dc8afde3c2becf8166936a3708d45c8cf0a73)
- enable_challenge.js_challenge_parameters

<a id="canonical-3b8f0c916aa2d09f73848e071ac47de26d81457eaf4f14c19b8d9886c4e94e8b"></a>

Type: `"object"`. single nested block, Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry",
    "js_script_delay")}
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
js_challenge_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-b1034fb8f78e61fac6f92491bc2253e2bab2af37b39c32c439eff8cc2811fd50"></a>

## Direct properties — enable_challenge.js_challenge_parameters / b248fb0b3c4a / 3

<a id="canonical-0f0cdedd8845747373d5c528bddce2a1f4f11ea947c944cca819305dc89d2917"></a>

<a id="canonical-74b18a7005b60819901aee2ed5216d92367a8a741168adfe1ea52ca0b13fd2a3"></a>

## cookie_expiry property — enable_challenge.js_challenge_parameters / b248fb0b3c4a / 4

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
}
```

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

<a id="canonical-bdd162d9d96f27cbb55177f290d6533b37d3a8ede8d78816f07591b85a898f22"></a>

<a id="canonical-eef50ec0b1b7fdbe45914f976501a40609f35422bf44bf2dee356882fa6e677e"></a>

## custom_page property — enable_challenge.js_challenge_parameters / b248fb0b3c4a / 5

Type: `"string"`. Optional.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

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

<a id="canonical-22796553ed6786fe07abfd4488447a332c8e362ed3993b076779aae7af9b5930"></a>

<a id="canonical-fce474489579d5821d1a120b36b33ba38aa28a41c7ec82ebdf0bf2ee63ad18d6"></a>

## js_script_delay property — enable_challenge.js_challenge_parameters / b248fb0b3c4a / 6

Type: `"number"`. Optional.

Delay introduced by Javascript, in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1000, 60000),
}
```

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

<a id="canonical-74a80a8d616e5cdc8765d79399be6d434041c3c6c17b0b309f5905af149250d1"></a>

## Next pages — enable_challenge.js_challenge_parameters / b248fb0b3c4a / 7

- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-be72efa8da41d7b196abe830ae2dc8afde3c2becf8166936a3708d45c8cf0a73)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-3bc129571f61df9ce982c8ad2e87f4ba5d0d632d27c3c5ea9368f14b576ac8a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5303c763b7569fc6e1fc8f82bd060edb5136f19678bcd6d4e055a3d26f9c1b3"></a>

## enable_challenge.malicious_user_mitigation — enable_challenge.malicious_user_mitigation / aac58bc7e4bf / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-be72efa8da41d7b196abe830ae2dc8afde3c2becf8166936a3708d45c8cf0a73)
- enable_challenge.malicious_user_mitigation

<a id="canonical-c5ca7eb2165459e3a54e91d552cad1a19346da53e5423d1d5a7e18d3c86cccdc"></a>

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
malicious_user_mitigation {
  # Configure direct properties listed below.
}
```

<a id="canonical-e236694440d2c0686ae0f5125858bfa2b47c253b7e62e832273d0d47aef4d1a9"></a>

## Direct properties — enable_challenge.malicious_user_mitigation / aac58bc7e4bf / 3

<a id="canonical-ec940cb0b6c85bf16ffdb56e016947a12d7ecf3216ab11db526617ceeea3b4ff"></a>

<a id="canonical-c0e77f5bef7515028a429f1c90fd177bcb06b9fc3ee27430da137771589d5ee7"></a>

## name property — enable_challenge.malicious_user_mitigation / aac58bc7e4bf / 4

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

<a id="canonical-da1e371be82dea98513e817936fea601e17e8d8f699eae3e69fbd2791edbf665"></a>

<a id="canonical-c5e8364ea57754df87aafa788e6e3af5d4a2c01e316fe4b08916cedbe574d949"></a>

## namespace property — enable_challenge.malicious_user_mitigation / aac58bc7e4bf / 5

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

<a id="canonical-ba4dde9d1ad9c7f3307ea3044097f97099d3af29572e2a0ca26e6bae518a823c"></a>

<a id="canonical-df01fac16e689107587247040290847a7b221d271d211ae9b76a6073cca6904d"></a>

## tenant property — enable_challenge.malicious_user_mitigation / aac58bc7e4bf / 6

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

<a id="canonical-eaaded720d5eac3b4d9817ef4d953bb9ed1aee70785b9f611f43d88e17f36912"></a>

## Next pages — enable_challenge.malicious_user_mitigation / aac58bc7e4bf / 7

- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-be72efa8da41d7b196abe830ae2dc8afde3c2becf8166936a3708d45c8cf0a73)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-de2142fc296e6a30b369e4c9b13a4ed016e7ed01af30d672900c574d3983bcda"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ee5ab21f0da64aa9587fdf56096b3c233730d12c6839439f39496204beabcde"></a>

## enable_ip_reputation — enable_ip_reputation / cba1e017fd18 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- enable_ip_reputation

<a id="canonical-5bac413884f254b3c379fc55c7e54bdee80f028c9708c3c3e95b73a2f8d5c478"></a>

Type: `"object"`. single nested block, Optional.

IP Threat Category List. List of IP threat categories.

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
enable_ip_reputation {
  # Configure direct properties listed below.
}
```

<a id="canonical-a2c9f73bf2de0a2e4dbdf993eb23ecda52edc330644322fbc1edbfa927fe3415"></a>

## Direct properties — enable_ip_reputation / cba1e017fd18 / 3

<a id="canonical-aabe9c2bbfa2d1deaf0211a3a309c21994a3f6bcc7410519d90955e9d7e20978"></a>

<a id="canonical-265203b1a62874abc0d463cd6dcbf2e6076aaebf68e09cd71530203836e39b6b"></a>

## ip_threat_categories property — enable_ip_reputation / cba1e017fd18 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
If the source IP matches on atleast one of the enabled IP threat categories, the request will be
denied. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`, \`WEB\_ATTACKS\`, \`BOTNETS\`,
\`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`, \`MOBILE\_THREATS\`, \`TOR\_PROXY\`,
\`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to \`SPAM\_SOURCES\`.

Upstream description:

If the source IP matches on atleast one of the enabled IP threat categories, the request will be
denied.

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

<a id="canonical-41292151514ba85d48f5897407a8c31c862d9360bf0df8a434124e45fdba3631"></a>

## Next pages — enable_ip_reputation / cba1e017fd18 / 5

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-739cd48cdacabeb3c20e33071d28210102da0caafd6e525f2c7bc6a4f191a530"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2a7e23e4f72f36757894e6efc0fecf06d3922fe6f15553e14c27c6da7d78b10"></a>

## enable_malicious_user_detection — enable_malicious_user_detection / c0495ab4d597 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- enable_malicious_user_detection

<a id="canonical-5bba34c4a9e1b9d67339de478ebf9cd26663448a3eb30a6488d51abfa6fbc9f0"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable malicious user detection.

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
enable_malicious_user_detection = {}
```

<a id="canonical-15a9941c792330347b818af800a623f5c05e4140db4a977049089a9eeb17dca7"></a>

## Direct properties — enable_malicious_user_detection / c0495ab4d597 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-41ae9aec7805530a3dab7b78b51f8d3e05d81ab1ff3f5ac88c1315efec65301a"></a>

## Next pages — enable_malicious_user_detection / c0495ab4d597 / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-f1b1fa9bfde00138f0ebbee2539ef4ce9300d5d0a630096b677d0a200b99ba0c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2a4cd44d806bed6c9c7b47fa22fe247823c76a178cdc55089fbdc237eed3a71e"></a>

## enable_threat_mesh — enable_threat_mesh / fb0fb1c9106f / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- enable_threat_mesh

<a id="canonical-9b9efccc5c5d1dbf2693d27cd3bfee17646a62a79935097e4d112aa7a438dac7"></a>

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
enable_threat_mesh = {}
```

<a id="canonical-a3de92a99a2abf42fe16b5b8c5508ff5fba74128c651b63347701d5496f5d268"></a>

## Direct properties — enable_threat_mesh / fb0fb1c9106f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-da5e8e4769e018756e0a0d51ee0cbb75d8f0aa4cd6ac6a4405ca2d6fb6252fb3"></a>

## Next pages — enable_threat_mesh / fb0fb1c9106f / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-1dd230b542892f64d8052e5db7f27223d60233a84fe969e8ce6e411444895cab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ea784d3ff8bf11a9887cb37e126ea383a8de4b14928c6c6bcf49c69735cbf9f6"></a>

## graphql_rules — graphql_rules / 19983d10bf33 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- graphql_rules

<a id="canonical-cf5eafbddec6903ea6b2fb7ebcfe4a313a7417335de13bea7cc656fcae5ceb13"></a>

Type: `"object"`. list nested block, Optional.

GraphQL is a query language and server-side runtime for APIs which provides a complete and
understandable description of the data in API. GraphQL gives clients the power to ask for exactly
what they need, makes it easier to evolve APIs over time, and enables powerful developer tools.
Policy..

Upstream description:

GraphQL is a query language and server-side runtime for APIs which provides a complete and
understandable description of the data in API. GraphQL gives clients the power to ask for exactly
what they need, makes it easier to evolve APIs over time, and enables powerful developer tools.
Policy configuration to analyze GraphQL queries and prevent GraphQL tailored attacks.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("exact_path"),
  validators.ConflictingListObjectAttributes("any_domain",
    "exact_value"),
  validators.ConflictingListObjectAttributes("any_domain",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("method_get",
    "method_post")}
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

Terraform syntax:

```terraform
graphql_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-0eebabb559460ca5d8538ef7ef35969f529c8ef97aff639b09864097d2b3738c"></a>

## Direct properties — graphql_rules / 19983d10bf33 / 3

- [any_domain](resources--cdn_loadbalancer--reference--group-010.md#canonical-34dbf830417d6871ad1a70cf65b5bd632f00af9506b86619503c2785ffa143ce): complete subsection reference.

<a id="canonical-961e0c14e174b7b01290c0230fea855d67d3d6d900c5ee5d6e3da15cc80ce11b"></a>

<a id="canonical-1226529990d0cd8ab2b7a615685adac6955134a96e9be66ab1dceaf052a778f8"></a>

## exact_path property — graphql_rules / 19983d10bf33 / 4

Type: `"string"`. Optional.

Specifies the exact path to GraphQL endpoint. Defaults to \`/graphql\`.

Upstream description:

Specifies the exact path to GraphQL endpoint. Default value is /graphql.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-50bafa52f23ef161aca16a13b170db64fa3700383029ba6d3fe09165e4587eb2"></a>

<a id="canonical-3fdca0778b3c6436e84ab65b8133b9da24f12b5f891866789993e6e0fcab6c25"></a>

## exact_value property — graphql_rules / 19983d10bf33 / 5

Type: `"string"`. Optional.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

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

- [graphql_settings](resources--cdn_loadbalancer--reference--group-010.md#canonical-33899e6867c8665ac41fd0693c8a579bfc9c6991cd0630a757078a072c86933d): complete subsection reference.

- [metadata](resources--cdn_loadbalancer--reference--group-010.md#canonical-954ce5e9b0f3d3b3c988ba467023424d415cd0b61a3658ff95588018a3303ea5): complete subsection reference.

- [method_get](resources--cdn_loadbalancer--reference--group-010.md#canonical-d6d7237487135201469ef0bde3dae4b8fea8b359972f819bc3ba36b5c8056bb9): complete subsection reference.

- [method_post](resources--cdn_loadbalancer--reference--group-010.md#canonical-d18487ea9bf8d8706f2ee1c7c4c6875b0015e0494d07f77cf349e98c61d4871f): complete subsection reference.

<a id="canonical-fea6b1db8fea977df69cde6547cf7dc1a41f2419de7da1da9350e30e39076c03"></a>

<a id="canonical-5cc9f1aa1f4f4fefa75bf27169f455728fc4d1054da6ac57df21c9b808a1eabb"></a>

## suffix_value property — graphql_rules / 19983d10bf33 / 6

Type: `"string"`. Optional.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
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

<a id="canonical-affe6ecf441be1739cc071a4114ca5f0d755f6bdfe3d829bcb4f65ebd34442ef"></a>

## Next pages — graphql_rules / 19983d10bf33 / 7

- [graphql_rules.any_domain](resources--cdn_loadbalancer--reference--group-010.md#canonical-34dbf830417d6871ad1a70cf65b5bd632f00af9506b86619503c2785ffa143ce)
- [graphql_rules.graphql_settings](resources--cdn_loadbalancer--reference--group-010.md#canonical-33899e6867c8665ac41fd0693c8a579bfc9c6991cd0630a757078a072c86933d)
- [graphql_rules.metadata](resources--cdn_loadbalancer--reference--group-010.md#canonical-954ce5e9b0f3d3b3c988ba467023424d415cd0b61a3658ff95588018a3303ea5)
- [graphql_rules.method_get](resources--cdn_loadbalancer--reference--group-010.md#canonical-d6d7237487135201469ef0bde3dae4b8fea8b359972f819bc3ba36b5c8056bb9)
- [graphql_rules.method_post](resources--cdn_loadbalancer--reference--group-010.md#canonical-d18487ea9bf8d8706f2ee1c7c4c6875b0015e0494d07f77cf349e98c61d4871f)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-34dbf830417d6871ad1a70cf65b5bd632f00af9506b86619503c2785ffa143ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9444a7edd431423eca32f35bb3568dadc58e2a8a9ec811b7376ed866a64aa2fd"></a>

## graphql_rules.any_domain — graphql_rules.any_domain / fb3f341c5417 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [graphql_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-1dd230b542892f64d8052e5db7f27223d60233a84fe969e8ce6e411444895cab)
- graphql_rules.any_domain

<a id="canonical-c16c3467b0ec3480dc2dc5a489bca14014c031eadcea33b22d5e83ac8971b05a"></a>

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

<a id="canonical-618619fa6c2ddc0443afb0ea78a1009599c88793a139e8079c37ff36222d4755"></a>

## Direct properties — graphql_rules.any_domain / fb3f341c5417 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f27fb14225a36e71660ae5b2f336bf491f4b410938c04188731002a2a1da58c8"></a>

## Next pages — graphql_rules.any_domain / fb3f341c5417 / 4

- [graphql_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-1dd230b542892f64d8052e5db7f27223d60233a84fe969e8ce6e411444895cab)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-33899e6867c8665ac41fd0693c8a579bfc9c6991cd0630a757078a072c86933d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a26c4782ece1169d89a25347ce861b472d52471d37a8fc329268ee2977f68900"></a>

## graphql_rules.graphql_settings — graphql_rules.graphql_settings / 4bb2d2bc3b0a / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [graphql_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-1dd230b542892f64d8052e5db7f27223d60233a84fe969e8ce6e411444895cab)
- graphql_rules.graphql_settings

<a id="canonical-055d7a3dff38ad92637edbdd1e71c22e8f150444776c706dd74cab2fdcd50d7d"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for graphql settings.

Upstream description:

GraphQL configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("max_batched_queries",
    "max_depth",
    "max_total_length"),
  validators.ConflictingObjectAttributes("disable_introspection",
    "enable_introspection")}
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
  "x-ves-oneof-field-allow_introspection_queries_choice": "[\"disable_introspection\",\"enable_introspection\"]"
}
```

Terraform syntax:

```terraform
graphql_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-c74771d5fe74f195e2ea0983a58ecf4ff675b5c21ac8d2d77ef2555583ec90bb"></a>

## Direct properties — graphql_rules.graphql_settings / 4bb2d2bc3b0a / 3

- [disable_introspection](resources--cdn_loadbalancer--reference--group-010.md#canonical-c58830ce01d137fd15d4ba2ce27f704d85ce33af73e135196e5ac5b1b1b26c6f): complete subsection reference.

- [enable_introspection](resources--cdn_loadbalancer--reference--group-010.md#canonical-fb2582bf4c0501994c58f5ca2ea842b89a77e782ff67f0fb622b28205aa183b4): complete subsection reference.

<a id="canonical-e5a1731f609dba6e31a379f4747aab243144d83312e63b2815e3865af6cf6678"></a>

<a id="canonical-822365152f6e84f18adca1077047c3e98ac2b8e023cad4e7931aa3549e6f00e2"></a>

## max_batched_queries property — graphql_rules.graphql_settings / 4bb2d2bc3b0a / 4

Type: `"number"`. Optional.

Specify maximum number of queries in a single batched request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="canonical-e79355ea8ad890e970ba75d199cff04838beed0a3645a8cbdd06a7db6894f415"></a>

<a id="canonical-57d204f4623b1d51043b12b4c813a3de1765816703363dab5044a6c95b0a5d25"></a>

## max_depth property — graphql_rules.graphql_settings / 4bb2d2bc3b0a / 5

Type: `"number"`. Optional.

Specify maximum depth for the GraphQL query.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="canonical-1cb39b573d17b8afb0e066d6d1a69b863a0cdcd53ce849662d136805daef7b45"></a>

<a id="canonical-0442c7105b6cbbbca91afeceb6de15b98c3cfd10aa5fbf718b0dc54133be46a9"></a>

## max_total_length property — graphql_rules.graphql_settings / 4bb2d2bc3b0a / 6

Type: `"number"`. Optional.

Specify maximum length in bytes for the GraphQL query.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 16386),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16386,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "16386"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "16386"
  }
}
```

<a id="canonical-aabfafe06091a7ad3e0bf70f9d070a4da58029c5960c3d0d54bc64aa641cd79e"></a>

## Next pages — graphql_rules.graphql_settings / 4bb2d2bc3b0a / 7

- [graphql_rules.graphql_settings.disable_introspection](resources--cdn_loadbalancer--reference--group-010.md#canonical-c58830ce01d137fd15d4ba2ce27f704d85ce33af73e135196e5ac5b1b1b26c6f)
- [graphql_rules.graphql_settings.enable_introspection](resources--cdn_loadbalancer--reference--group-010.md#canonical-fb2582bf4c0501994c58f5ca2ea842b89a77e782ff67f0fb622b28205aa183b4)
- [graphql_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-1dd230b542892f64d8052e5db7f27223d60233a84fe969e8ce6e411444895cab)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-c58830ce01d137fd15d4ba2ce27f704d85ce33af73e135196e5ac5b1b1b26c6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a3b2fca07cd890dcd37cfe391d5dd9a0e23d4540ef62e3bfd91619024d658222"></a>

## graphql_rules.graphql_settings.disable_introspection — graphql_rules.graphql_settings.disable_introspection / 7e984b7dc442 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [graphql_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-1dd230b542892f64d8052e5db7f27223d60233a84fe969e8ce6e411444895cab)
- [graphql_rules.graphql_settings](resources--cdn_loadbalancer--reference--group-010.md#canonical-33899e6867c8665ac41fd0693c8a579bfc9c6991cd0630a757078a072c86933d)
- graphql_rules.graphql_settings.disable_introspection

<a id="canonical-d10eb9dd322ff408fc037fef519e30b3e2d7b6c4e77da0563e000b0bacab3efb"></a>

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
disable_introspection = {}
```

<a id="canonical-b67413a2360855ce925d3b09ad47c7605df74c59ddafd2c08dddee7038e73d00"></a>

## Direct properties — graphql_rules.graphql_settings.disable_introspection / 7e984b7dc442 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c73448d94344c73c0f374dbbcab3f7edc2da8f5cb99ca439831093f518b464b9"></a>

## Next pages — graphql_rules.graphql_settings.disable_introspection / 7e984b7dc442 / 4

- [graphql_rules.graphql_settings](resources--cdn_loadbalancer--reference--group-010.md#canonical-33899e6867c8665ac41fd0693c8a579bfc9c6991cd0630a757078a072c86933d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-fb2582bf4c0501994c58f5ca2ea842b89a77e782ff67f0fb622b28205aa183b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e082b61dc919effc2548804b7bbe4a2c08c14287ddb45b880594ec0eaf344443"></a>

## graphql_rules.graphql_settings.enable_introspection — graphql_rules.graphql_settings.enable_introspection / f2db6021f7c9 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [graphql_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-1dd230b542892f64d8052e5db7f27223d60233a84fe969e8ce6e411444895cab)
- [graphql_rules.graphql_settings](resources--cdn_loadbalancer--reference--group-010.md#canonical-33899e6867c8665ac41fd0693c8a579bfc9c6991cd0630a757078a072c86933d)
- graphql_rules.graphql_settings.enable_introspection

<a id="canonical-166cc01ff9b786ddb3dd4683e69d48ea11f1849e3b493b80572f847f7837706e"></a>

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
enable_introspection = {}
```

<a id="canonical-3d20a3b6b830edc5fb70c160ed763f6bac98da8f504b37b6b771d3619ff56a3e"></a>

## Direct properties — graphql_rules.graphql_settings.enable_introspection / f2db6021f7c9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3f3b933cdeaf28819f09a5b1e5b10a7ee57c99c26ca894c95a2f420698599fef"></a>

## Next pages — graphql_rules.graphql_settings.enable_introspection / f2db6021f7c9 / 4

- [graphql_rules.graphql_settings](resources--cdn_loadbalancer--reference--group-010.md#canonical-33899e6867c8665ac41fd0693c8a579bfc9c6991cd0630a757078a072c86933d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-954ce5e9b0f3d3b3c988ba467023424d415cd0b61a3658ff95588018a3303ea5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-036ccb28c674895f0901c080f0bf2abad3a6f7730cf0bbdbb7e8516913f0b7c4"></a>

## graphql_rules.metadata — graphql_rules.metadata / 430c962499b4 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [graphql_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-1dd230b542892f64d8052e5db7f27223d60233a84fe969e8ce6e411444895cab)
- graphql_rules.metadata

<a id="canonical-bcd848fc4650733e64f7b205ad5a38a7376fb1e61c5354febbfb0e6b8a6170e7"></a>

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

<a id="canonical-6bab47eb19667f611ffacf20d01429c40351a1c9a931f79e6a54de8e8269ee9f"></a>

## Direct properties — graphql_rules.metadata / 430c962499b4 / 3

<a id="canonical-907fd3864e51ee05c463ef005913ec15004f4e77d535d84c5f9c7efb46cede79"></a>

<a id="canonical-ce4380fbc5d98194ecac61ffc7de21858864415789b2cc578098449d1510ce61"></a>

## description_spec property — graphql_rules.metadata / 430c962499b4 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-b258a4e2007d34b3dfa5e71410f66f80edf9a2397a7559764e7ca10952458573"></a>

<a id="canonical-d49e23ddfbea019cf7b7223f3bfd9d26d042332d13d16efad9dd7c748d469598"></a>

## name property — graphql_rules.metadata / 430c962499b4 / 5

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

<a id="canonical-681334332148ebe66caa8b16264743f0f1772b6f91789589b4e81ff07f2cb57b"></a>

## Next pages — graphql_rules.metadata / 430c962499b4 / 6

- [graphql_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-1dd230b542892f64d8052e5db7f27223d60233a84fe969e8ce6e411444895cab)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-d6d7237487135201469ef0bde3dae4b8fea8b359972f819bc3ba36b5c8056bb9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c7a12d94687e673fd8ce3483deb28a9cadf0cc6e80533b91ced2c7a7010303f2"></a>

## graphql_rules.method_get — graphql_rules.method_get / ebfc6c685c24 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [graphql_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-1dd230b542892f64d8052e5db7f27223d60233a84fe969e8ce6e411444895cab)
- graphql_rules.method_get

<a id="canonical-eabe94562279806161a96ed2624daa17fdaa0dfac35345ca9a101be539595333"></a>

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
method_get = {}
```

<a id="canonical-db4558b6ef9f61d6292141a164c71e5b7d991f1394163a9ed333480cf150668c"></a>

## Direct properties — graphql_rules.method_get / ebfc6c685c24 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-810d7e7a53fe97ccccf830d445f15660003f76cf8cf6a49a06ee41511bc683e0"></a>

## Next pages — graphql_rules.method_get / ebfc6c685c24 / 4

- [graphql_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-1dd230b542892f64d8052e5db7f27223d60233a84fe969e8ce6e411444895cab)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-d18487ea9bf8d8706f2ee1c7c4c6875b0015e0494d07f77cf349e98c61d4871f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb32d170812230fdf19faa2ab68eff376d70bc8af25b79e2d900a0301ec091a6"></a>

## graphql_rules.method_post — graphql_rules.method_post / c989097c1dda / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [graphql_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-1dd230b542892f64d8052e5db7f27223d60233a84fe969e8ce6e411444895cab)
- graphql_rules.method_post

<a id="canonical-07b22ed947f9907579c4b8546fc52da122937b7f88460511e3017186fb494e4a"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for method post.

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
method_post = {}
```

<a id="canonical-c0e8906fe22d122aaf54158964122f16edadf8d57e64ae4932d2d35d671192ad"></a>

## Direct properties — graphql_rules.method_post / c989097c1dda / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b77e58e7896af76dec0fc8ed00d2dbb519d26bc6e254d360863985b694516b00"></a>

## Next pages — graphql_rules.method_post / c989097c1dda / 4

- [graphql_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-1dd230b542892f64d8052e5db7f27223d60233a84fe969e8ce6e411444895cab)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-f81a9f81cb120d5f06895975b67987a4ceae629b66a48aeaa3ba2a2ea3bb414d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d48ce91a6c4c8d1bfaad4260cfbae529360c898454730e982c22a27cacf7b036"></a>

## http — http / faf9543ce75b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- http

<a id="canonical-4064fd63a477a3a8e6ec5749013165d110ceb921a6006b05dc1854475f9e923b"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: http, https, https\_auto\_cert; Default: https\_auto\_cert\] HTTP Choice. Choice for
selecting HTTP proxy.

Upstream description:

Choice for selecting HTTP proxy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("port",
    "port_ranges")}
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
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]"
}
```

OneOf alternatives in this subsection:

- [http](resources--cdn_loadbalancer--reference--group-010.md#canonical-4064fd63a477a3a8e6ec5749013165d110ceb921a6006b05dc1854475f9e923b)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-0c91bcac9ef4fc1454cda9750bdf030050a524aafd2b2d97358d25822e8613c6)
- [https_auto_cert](resources--cdn_loadbalancer--reference--group-011.md#canonical-4505e0388e96d2b6532ac846a2fb56512c8f4078a42d43348548aac21f6c9c37)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
http {
  # Configure direct properties listed below.
}
```

<a id="canonical-20c5dabda3301b1c5734c858aa5b10a5bfa0f181609477e76f0c5f0c3cadc105"></a>

## Direct properties — http / faf9543ce75b / 3

<a id="canonical-3310e275175db576720dc17a6c1f62895606fbdba4486ee7d9144d60a6a1d77b"></a>

<a id="canonical-a5deb9ed15ba937ac38e6bd4078cb089422ea5a3e582afbbcd890acae82db614"></a>

## dns_volterra_managed property — http / faf9543ce75b / 4

Type: `"bool"`. Optional.

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

Upstream description:

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-759d963d22508c41edeacba400409ed90dd49ebe1a120597b5e73b13be302a63"></a>

<a id="canonical-4df89bddad175383635e0c0da80969a5129d160b8ee70cda8fc0d481498a7448"></a>

## port property — http / faf9543ce75b / 5

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTP port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTP port to Listen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-846005e8657221cb212d7aabe052d26524f629ea4116106fe284dc00022e29b9"></a>

<a id="canonical-1301859e7f4a6fcb69821149c2a9497c32569347c67c7574fa6199ebaf744000"></a>

## port_ranges property — http / faf9543ce75b / 6

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-3746fe42c3e31e2fdeacf79e5e69fc3142a1cfe2991e9cb36dbdb8fd5f9b1913"></a>

## Next pages — http / faf9543ce75b / 7

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5daa8a1c03ed136144cb520ef80445f172d77afb77c530290046ca2473b98cda"></a>

## https — https / 800b920050b9 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- https

<a id="canonical-0c91bcac9ef4fc1454cda9750bdf030050a524aafd2b2d97358d25822e8613c6"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting CDN Distribution with bring your own certificates.

Receipt-pinned upstream constraints:

```json
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
https {
  # Configure direct properties listed below.
}
```

<a id="canonical-832fb678c8d4920794a42766c6b9f1d4cbaa4d57f62a28b566aa51420843f894"></a>

## Direct properties — https / 800b920050b9 / 3

<a id="canonical-1f35f99c793bb05caf077df489c67969c41d27352e6e673a606f3f0b99ad1730"></a>

<a id="canonical-779345ba08473713c6b28950a06b87b29f488deead7743affbb3029ed1199cec"></a>

## add_hsts property — https / 800b920050b9 / 4

Type: `"bool"`. Optional.

Add HTTP Strict-Transport-Security response header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-bda228a2e05251b81a9136c77864232e9cfacddc2421add47bf97a0566347909"></a>

<a id="canonical-168a782cad0975a7a02b350998e578beaca278fd2c4cb5a28472542b883ee83e"></a>

## http_redirect property — https / 800b920050b9 / 5

Type: `"bool"`. Optional.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

Upstream description:

Redirect HTTP traffic to HTTPS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e): complete subsection reference.

<a id="canonical-710aa1e7eb141f2f1cd94ee62d2d94da5142151b887901da99a0a51c03b36c51"></a>

## Next pages — https / 800b920050b9 / 6

- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-02d64e797fb76fb5e406a323099d5b646e7075fbe81ecafcaaef4b7ce9d55657"></a>

## https.tls_cert_options — https.tls_cert_options / 3dacc187dffe / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- https.tls_cert_options

<a id="canonical-825764684fb7503b02c7c9bc4c5dc6d8e8a8d28b6cbef9484a30d4d023821936"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls cert options.

Upstream description:

TLS Certificate OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("tls_cert_params",
    "tls_inline_params")}
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
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_inline_params\"]"
}
```

Terraform syntax:

```terraform
tls_cert_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-6b60b11d6849ff83dd6ae61a2910227d7a1ba8600b3ed7b5e9fe06341f173b00"></a>

## Direct properties — https.tls_cert_options / 3dacc187dffe / 3

- [tls_cert_params](resources--cdn_loadbalancer--reference--group-010.md#canonical-fadff03b94e0c5d74340d2ed58c1f2d8a3e84bb577ce7b31f80c2165765711d1): complete subsection reference.

- [tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-057144dabb883c47ae08cfc64b0e66ed7985be130eacf860907c9e83b0725d50): complete subsection reference.

<a id="canonical-a642b94ad04d634330a4ab51e0be07b140252dcde585133c3300691033619c02"></a>

## Next pages — https.tls_cert_options / 3dacc187dffe / 4

- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-010.md#canonical-fadff03b94e0c5d74340d2ed58c1f2d8a3e84bb577ce7b31f80c2165765711d1)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-057144dabb883c47ae08cfc64b0e66ed7985be130eacf860907c9e83b0725d50)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-fadff03b94e0c5d74340d2ed58c1f2d8a3e84bb577ce7b31f80c2165765711d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-435cadfba43760c4f30090dd07fac528365e37d8bacd1df137aaf7b4956f2c3f"></a>

## https.tls_cert_options.tls_cert_params — https.tls_cert_options.tls_cert_params / 65d5b24172b8 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- https.tls_cert_options.tls_cert_params

<a id="canonical-f7f7d3ba13127a29fc672235948b9b34fc515f5d4cd2a658ba1cc033553b67cb"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls cert params.

Upstream description:

Select TLS Parameters and Certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("certificates"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls")}
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
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
tls_cert_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-9f4c9bf6cbc456554b5e8afc707f4ddd60c8e1ba98d21734e004dd1d74472d69"></a>

## Direct properties — https.tls_cert_options.tls_cert_params / 65d5b24172b8 / 3

- [certificates](resources--cdn_loadbalancer--reference--group-010.md#canonical-4827f325c580098be4361139249dab448aaf3c9fde908897ebe4c564ca3956ba): complete subsection reference.

- [no_mtls](resources--cdn_loadbalancer--reference--group-010.md#canonical-c9827d560e5c4db10cccc7c66f1a9f4b9ff355e58a289d2456ccb825a331cfaf): complete subsection reference.

- [tls_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-4df5a409dd57c563bffde36f9cf494186e776d781835c6e847c44c94ff489dd8): complete subsection reference.

- [use_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-8735be1ccb4030c11c11db4bd0916a0b2db0891a407773421507e28f12bbdeb8): complete subsection reference.

<a id="canonical-4aa9e40cb196ef2fbac7cc71cc60ccff2b01f78dadb4f6f88fc8d50e23a0f428"></a>

## Next pages — https.tls_cert_options.tls_cert_params / 65d5b24172b8 / 4

- [https.tls_cert_options.tls_cert_params.certificates](resources--cdn_loadbalancer--reference--group-010.md#canonical-4827f325c580098be4361139249dab448aaf3c9fde908897ebe4c564ca3956ba)
- [https.tls_cert_options.tls_cert_params.no_mtls](resources--cdn_loadbalancer--reference--group-010.md#canonical-c9827d560e5c4db10cccc7c66f1a9f4b9ff355e58a289d2456ccb825a331cfaf)
- [https.tls_cert_options.tls_cert_params.tls_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-4df5a409dd57c563bffde36f9cf494186e776d781835c6e847c44c94ff489dd8)
- [https.tls_cert_options.tls_cert_params.use_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-8735be1ccb4030c11c11db4bd0916a0b2db0891a407773421507e28f12bbdeb8)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-4827f325c580098be4361139249dab448aaf3c9fde908897ebe4c564ca3956ba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-523ad53566dc65dcbf6521ef706fcbc93e09b488b991ee82f6156a9e9bc8b8a3"></a>

## https.tls_cert_options.tls_cert_params.certificates — https.tls_cert_options.tls_cert_params.certificates / 3980e41ab10f / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-010.md#canonical-fadff03b94e0c5d74340d2ed58c1f2d8a3e84bb577ce7b31f80c2165765711d1)
- https.tls_cert_options.tls_cert_params.certificates

<a id="canonical-8042a6a4a2b8c761408e185ec3199d8f6a96b8d70dfad3acd7ed999d013306b4"></a>

Type: `"object"`. list nested block, Optional.

Select one or more certificates with any domain names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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

Terraform syntax:

```terraform
certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-6b840fef452fc1b0b2526432e947b65443f2393ef31989e91e069aecb2506b88"></a>

## Direct properties — https.tls_cert_options.tls_cert_params.certificates / 3980e41ab10f / 3

<a id="canonical-46d847d8a575442a0710b448f19482f5cfb428101d6f9d9edbce972bd61b62d9"></a>

<a id="canonical-48751a685718e2b19aa48601e9ed95d23c1d49d5759924fc73a3743bb91a13d3"></a>

## name property — https.tls_cert_options.tls_cert_params.certificates / 3980e41ab10f / 4

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

<a id="canonical-30cca7634c400c20e82e04f5056d1cf8066e033c06c92208b42c69dd792f06fe"></a>

<a id="canonical-2b128fca5276310952069e1d24c77914e9c1ebd4fcedc3d2402ea964a1b5a346"></a>

## namespace property — https.tls_cert_options.tls_cert_params.certificates / 3980e41ab10f / 5

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

<a id="canonical-553c4acaccffdb57478ad88ab8868ef428ad992f82c91df145315724057992a7"></a>

<a id="canonical-2d13f3cd05c8090ba4389d595ee23b0fd64569e33a8919f49a562199f67d21bc"></a>

## tenant property — https.tls_cert_options.tls_cert_params.certificates / 3980e41ab10f / 6

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

<a id="canonical-b71fb3b64898e06abaf4e5b509972ad2f7588e0c541d0cef814884a4b649237b"></a>

## Next pages — https.tls_cert_options.tls_cert_params.certificates / 3980e41ab10f / 7

- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-010.md#canonical-fadff03b94e0c5d74340d2ed58c1f2d8a3e84bb577ce7b31f80c2165765711d1)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-c9827d560e5c4db10cccc7c66f1a9f4b9ff355e58a289d2456ccb825a331cfaf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-351a40566c98ecfe404892e0e4c78093c88a287dafc1d4418b26e1332202dfdf"></a>

## https.tls_cert_options.tls_cert_params.no_mtls — https.tls_cert_options.tls_cert_params.no_mtls / 718e15b85e1a / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-010.md#canonical-fadff03b94e0c5d74340d2ed58c1f2d8a3e84bb577ce7b31f80c2165765711d1)
- https.tls_cert_options.tls_cert_params.no_mtls

<a id="canonical-c8aa3bb4868dd82c51de3a57862d57c5715d4ae66b94133b7fd0b0e23dab2d46"></a>

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
no_mtls = {}
```

<a id="canonical-df7b92b04432d60f39ce9c6163a383d029daf1609188d9c83fa359abc7e3fb0c"></a>

## Direct properties — https.tls_cert_options.tls_cert_params.no_mtls / 718e15b85e1a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-50cc986f32b89c980ce92831f9d7c1fc85bcdc4dc899bc900873422f45a65647"></a>

## Next pages — https.tls_cert_options.tls_cert_params.no_mtls / 718e15b85e1a / 4

- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-010.md#canonical-fadff03b94e0c5d74340d2ed58c1f2d8a3e84bb577ce7b31f80c2165765711d1)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-4df5a409dd57c563bffde36f9cf494186e776d781835c6e847c44c94ff489dd8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2212dff76542cbe1f8fb41b0e408846c0fdcd5719d5200dfdafe4b084eb6861"></a>

## https.tls_cert_options.tls_cert_params.tls_config — https.tls_cert_options.tls_cert_params.tls_config / 3c8199184a33 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-010.md#canonical-fadff03b94e0c5d74340d2ed58c1f2d8a3e84bb577ce7b31f80c2165765711d1)
- https.tls_cert_options.tls_cert_params.tls_config

<a id="canonical-101dc3841913f7389f7d5652ea152ce40008d10e2acd4cab19cc476be6b7ef70"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-3bcf2ed39cc86b440458052ed938beae596db797fc089ea443985350f793aa95"></a>

## Direct properties — https.tls_cert_options.tls_cert_params.tls_config / 3c8199184a33 / 3

- [custom_security](resources--cdn_loadbalancer--reference--group-010.md#canonical-3bebb67f73d689ae2b85f24a48146d88f6933962cea1a8bdb624b09557bbc04f): complete subsection reference.

- [default_security](resources--cdn_loadbalancer--reference--group-011.md#canonical-1805f14acbfc1f2d8a48b2de257b79cfcd83583a0c69ead1e2a7aba9e0609879): complete subsection reference.

- [low_security](resources--cdn_loadbalancer--reference--group-011.md#canonical-e17221937085b12d0e8470a728fa6bf2f65a2e8b42116ab37fefc46c1540d0be): complete subsection reference.

- [medium_security](resources--cdn_loadbalancer--reference--group-011.md#canonical-bc1838534f155e4cc22e7f4e2b2f415b409f8ef9e3edcddea0957f88e0f75594): complete subsection reference.

<a id="canonical-617610aec074f45dbf6d359c028db258dc043e5c45f23d2a65548243df2f567f"></a>

## Next pages — https.tls_cert_options.tls_cert_params.tls_config / 3c8199184a33 / 4

- [https.tls_cert_options.tls_cert_params.tls_config.custom_security](resources--cdn_loadbalancer--reference--group-010.md#canonical-3bebb67f73d689ae2b85f24a48146d88f6933962cea1a8bdb624b09557bbc04f)
- [https.tls_cert_options.tls_cert_params.tls_config.default_security](resources--cdn_loadbalancer--reference--group-011.md#canonical-1805f14acbfc1f2d8a48b2de257b79cfcd83583a0c69ead1e2a7aba9e0609879)
- [https.tls_cert_options.tls_cert_params.tls_config.low_security](resources--cdn_loadbalancer--reference--group-011.md#canonical-e17221937085b12d0e8470a728fa6bf2f65a2e8b42116ab37fefc46c1540d0be)
- [https.tls_cert_options.tls_cert_params.tls_config.medium_security](resources--cdn_loadbalancer--reference--group-011.md#canonical-bc1838534f155e4cc22e7f4e2b2f415b409f8ef9e3edcddea0957f88e0f75594)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-010.md#canonical-fadff03b94e0c5d74340d2ed58c1f2d8a3e84bb577ce7b31f80c2165765711d1)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-3bebb67f73d689ae2b85f24a48146d88f6933962cea1a8bdb624b09557bbc04f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7360b141bfe1456f661379cbae2f564c4ec90c9cf6b565202c02910dd74b61d6"></a>

## https.tls_cert_options.tls_cert_params.tls_config.custom_security — https.tls_cert_options.tls_cert_params.tls_config.custom_security / 56bc8aa45d97 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-010.md#canonical-fadff03b94e0c5d74340d2ed58c1f2d8a3e84bb577ce7b31f80c2165765711d1)
- [https.tls_cert_options.tls_cert_params.tls_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-4df5a409dd57c563bffde36f9cf494186e776d781835c6e847c44c94ff489dd8)
- https.tls_cert_options.tls_cert_params.tls_config.custom_security

<a id="canonical-d6adde514e099d6fb48bcc14dfce58d025f0fe5379fe01b8dee39b8143af162a"></a>

Type: `"object"`. single nested block, Optional.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
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
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-c5030710a6b817f43d8f9a69e7ccf6053ff068a10818042c818be91c0371e413"></a>

## Direct properties — https.tls_cert_options.tls_cert_params.tls_config.custom_security / 56bc8aa45d97 / 3

<a id="canonical-10f89e56bcce54586ec9b24ba6c7109aa657df6fb687057f09be480c8201d8e9"></a>
