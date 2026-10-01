---
page_title: "xcsh_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy reference."
---

# xcsh_proxy reference

<a id="canonical-5463dd0b4f625a5cde33b77bb2bb6e364acdabae2fba8632a9c0b304af7c3f23"></a>

## Next pages — site_virtual_sites.advertise_where.virtual_site / 1b1f32d97505 / 5

- [site_virtual_sites.advertise_where.virtual_site.virtual_site](resources--proxy--reference--group-005.md#canonical-2446fec07bb39696ba8c35542910b9a8ada175babf7e63d3661ee4fbc8b9c7f9)
- [site_virtual_sites.advertise_where](resources--proxy--reference--group-004.md#canonical-09def8c1ed4434ecd861127657544cd2810ba7bcbcc7fdf746874bbdee930402)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-2446fec07bb39696ba8c35542910b9a8ada175babf7e63d3661ee4fbc8b9c7f9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8db0a8e2d61dcc4050fad7b80e26480e4e34f23af0ccc4ef632ea3377ae3899a"></a>

## site_virtual_sites.advertise_where.virtual_site.virtual_site — site_virtual_sites.advertise_where.virtual_site.virtual_site / afcacfa1f37f / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [site_virtual_sites](resources--proxy--reference--group-004.md#canonical-ad367c86467f35f5c26ee33c952dcf2fcaf2a7515f5835cae209a925754863b2)
- [site_virtual_sites.advertise_where](resources--proxy--reference--group-004.md#canonical-09def8c1ed4434ecd861127657544cd2810ba7bcbcc7fdf746874bbdee930402)
- [site_virtual_sites.advertise_where.virtual_site](resources--proxy--reference--group-004.md#canonical-56cb8cc6f75ca94245ca58189fe0b09826f50a728714c3ad33c01761194c46d4)
- site_virtual_sites.advertise_where.virtual_site.virtual_site

<a id="canonical-29a3dfa68cd44bd3d5ef3f0f73688758c68d3fb24942c36ab338dc780692e6c9"></a>

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-08589df452a2914561470ad7283dec20de2387760600882de3948e5979eba34d"></a>

## Direct properties — site_virtual_sites.advertise_where.virtual_site.virtual_site / afcacfa1f37f / 3

<a id="canonical-3f144b13249edbbbeed35551078f149badf9a482576df616f0b14d24f52d2eb9"></a>

<a id="canonical-82fe45fda1f27e0aacc77ec7e21c14c128100a8ebaa5a176f7106981c11ecbb6"></a>

## name property — site_virtual_sites.advertise_where.virtual_site.virtual_site / afcacfa1f37f / 4

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

<a id="canonical-c4d0dee162da1a1ed2159a1a40d3d09295a7b669e0daab7fec78357a6752e5f3"></a>

<a id="canonical-d4d70448406364e1e946547766599888d43db31bd1c6119f1f1aa3f128cb8db0"></a>

## namespace property — site_virtual_sites.advertise_where.virtual_site.virtual_site / afcacfa1f37f / 5

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

<a id="canonical-6ea200dc7d3c48d66f310a5a3169581f14c2d9c26320b0c5235686ddeb86fb2c"></a>

<a id="canonical-207fb0b46fe3d7ce968628a83a1ef2eba8ab27274625a9ea49a202db5f4e7e4d"></a>

## tenant property — site_virtual_sites.advertise_where.virtual_site.virtual_site / afcacfa1f37f / 6

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

<a id="canonical-b4672a936bd5f9a09b0ec5818271a55c4793723e2d9fe2ba2765f47aa483206f"></a>

## Next pages — site_virtual_sites.advertise_where.virtual_site.virtual_site / afcacfa1f37f / 7

- [site_virtual_sites.advertise_where.virtual_site](resources--proxy--reference--group-004.md#canonical-56cb8cc6f75ca94245ca58189fe0b09826f50a728714c3ad33c01761194c46d4)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-9751a1a4304e20a986e19abdac2b60f6e30919c2760a358e62b4e168f6bfa05e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f3458b47d274a1ec6125eaa2ba2ce6c0ba252a11c012649bd720a651f7dee9e"></a>

## timeouts — timeouts / 197a948200a0 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- timeouts

<a id="canonical-d0a5e5eb29379d81d6c14549fa8c88f4c612ebb92a69c245ec1b81dd332eb21a"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-f798087d5515ff8b73beaf165c375fd5960848045aa7b17f2bcdd5a34a78f726"></a>

## Direct properties — timeouts / 197a948200a0 / 3

<a id="canonical-53629e55c23c4bc2d492abc92dd464d25b3eb05bc151619089e1416130f394b0"></a>

<a id="canonical-0a21ce3351f2c02d3db4cdff4962ebfcc76301baaee7f76a00a188feb27f497b"></a>

## create property — timeouts / 197a948200a0 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-69cfb48854127933bda9572f87a34292bba31a0d0a42bb27e52910ed8e3e8c70"></a>

<a id="canonical-2790bdc51e54cd0461264128c80f9078b2f04deeb63c3021bfe242b121fd127e"></a>

## delete property — timeouts / 197a948200a0 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-b82a9b406d2613e14b267fb16cd95c013b3f621b4b8413fd5f2452708e50f994"></a>

<a id="canonical-6418763a2862a9bc30b62c34e8577a4d8cb46ca966b88fc1709a961738cdcc11"></a>

## read property — timeouts / 197a948200a0 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-aedb71e93535ead60c15582b53b2dc5de0d700df665b9537d9e8c123df063f2b"></a>

<a id="canonical-7c4d1ba013c04bc7312f40e8b2e9c2814057e74253edccfd9b0fa83fffb2a2cb"></a>

## update property — timeouts / 197a948200a0 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-b60d6b25ae0efe65f6deb3ba61568c31f990d151db44457d7bcbe63330140dfc"></a>

## Next pages — timeouts / 197a948200a0 / 8

- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-4639ee2404690fbd00e5ce07dfcd45d8207d50ed983459435b1b3df804692b4d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a4dc4135a10d6040bd7fea354d3115d1879701134d4b7371bb0900d605f445a"></a>

## tls_intercept — tls_intercept / 357d75725273 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- tls_intercept

<a id="canonical-b09947c5b3a5d27f69f31c0a2da06055276dbe6a003232d163a89b3739b4716c"></a>

Type: `"object"`. single nested block, Optional.

Configuration to enable TLS interception.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_certificate",
    "volterra_certificate"),
  validators.ConflictingObjectAttributes("enable_for_all_domains",
    "policy"),
  validators.ConflictingObjectAttributes("trusted_ca_url",
    "volterra_trusted_ca")}
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
  "x-ves-oneof-field-interception_policy_choice": "[\"enable_for_all_domains\",\"policy\"]",
  "x-ves-oneof-field-signing_cert_choice": "[\"custom_certificate\",\"volterra_certificate\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca_url\",\"volterra_trusted_ca\"]"
}
```

Terraform syntax:

```terraform
tls_intercept {
  # Configure direct properties listed below.
}
```

<a id="canonical-577f14a27a302c6214f6257c5c666b8a8647ef0ba7bd76546b9b0b2746739a98"></a>

## Direct properties — tls_intercept / 357d75725273 / 3

- [custom_certificate](resources--proxy--reference--group-005.md#canonical-aaee6da9f3557a4a4b7e56d1f1be1236163b7cea90779d39d051158d211b5eeb): complete subsection reference.

- [enable_for_all_domains](resources--proxy--reference--group-005.md#canonical-04da480c4e48c82947b4b5552704c071f296e48a84a79e457875c2106cd72292): complete subsection reference.

- [policy](resources--proxy--reference--group-005.md#canonical-97d5300e2bd77f1f1170e86f5d3b0925b166d8e616fe40b3eeff08370d847453): complete subsection reference.

<a id="canonical-b2d18133147a3e800126d410bc2fd31d0b35cc2ee02f8b94571ef7fb74d38d62"></a>

<a id="canonical-f026629898d2df20678657bab8d6e07aabebddcebdcd085aef002967ccdd4311"></a>

## trusted_ca_url property — tls_intercept / 357d75725273 / 4

Type: `"string"`. Optional.

Exclusive with \[volterra\_trusted\_ca\] Custom Root CA Certificate for validating upstream server
certificate.

Upstream description:

Exclusive with \[volterra\_trusted\_ca\] Custom Root CA Certificate for validating upstream server
certificate.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(131072),
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
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [volterra_certificate](resources--proxy--reference--group-005.md#canonical-3a866c24dab44d630da87b83d903f073888f999314fee1a280191336a46835c5): complete subsection reference.

- [volterra_trusted_ca](resources--proxy--reference--group-005.md#canonical-296065aeea4832e5be6edf994ae5cbf0debd9ea52458ec265063b0ecff6e5441): complete subsection reference.

<a id="canonical-499f6f9d977bd1403bb9363b133a711f857a4d676ca2d4070473da50365b7626"></a>

## Next pages — tls_intercept / 357d75725273 / 5

- [tls_intercept.custom_certificate](resources--proxy--reference--group-005.md#canonical-aaee6da9f3557a4a4b7e56d1f1be1236163b7cea90779d39d051158d211b5eeb)
- [tls_intercept.enable_for_all_domains](resources--proxy--reference--group-005.md#canonical-04da480c4e48c82947b4b5552704c071f296e48a84a79e457875c2106cd72292)
- [tls_intercept.policy](resources--proxy--reference--group-005.md#canonical-97d5300e2bd77f1f1170e86f5d3b0925b166d8e616fe40b3eeff08370d847453)
- [tls_intercept.volterra_certificate](resources--proxy--reference--group-005.md#canonical-3a866c24dab44d630da87b83d903f073888f999314fee1a280191336a46835c5)
- [tls_intercept.volterra_trusted_ca](resources--proxy--reference--group-005.md#canonical-296065aeea4832e5be6edf994ae5cbf0debd9ea52458ec265063b0ecff6e5441)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-aaee6da9f3557a4a4b7e56d1f1be1236163b7cea90779d39d051158d211b5eeb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e24085031a3439719a998ea7a7d180cca6ef0a8d9a18f0df0ddaddffe3885f8"></a>

## tls_intercept.custom_certificate — tls_intercept.custom_certificate / 06a69f263787 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-4639ee2404690fbd00e5ce07dfcd45d8207d50ed983459435b1b3df804692b4d)
- tls_intercept.custom_certificate

<a id="canonical-1c296d6302f43d9e4f247cef86a712431b0a20f2f8f1d9a2f4e67f5e3ff3088c"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for custom certificate.

Upstream description:

Handle to fetch certificate and key.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("certificate_url"),
  validators.ConflictingObjectAttributes("custom_hash_algorithms",
    "disable_ocsp_stapling"),
  validators.ConflictingObjectAttributes("custom_hash_algorithms",
    "use_system_defaults"),
  validators.ConflictingObjectAttributes("disable_ocsp_stapling",
    "use_system_defaults")}
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
  "x-ves-oneof-field-ocsp_stapling_choice": "[\"custom_hash_algorithms\",\"disable_ocsp_stapling\",\"use_system_defaults\"]"
}
```

Terraform syntax:

```terraform
custom_certificate {
  # Configure direct properties listed below.
}
```

<a id="canonical-8f055464df7c579868c61c6c6415e470492d2bed84cda72dae9ecae6a1ad8db5"></a>

## Direct properties — tls_intercept.custom_certificate / 06a69f263787 / 3

<a id="canonical-a1287c752743836842397822744a2087bca40702fb550b6a3198fbb36a3bafe2"></a>

<a id="canonical-c088b561e9b01bd0cd15af1dc3b3e7480ebde70a6240f37581c54eef6088a4c8"></a>

## certificate_url property — tls_intercept.custom_certificate / 06a69f263787 / 4

Type: `"string"`. Optional.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](resources--proxy--reference--group-005.md#canonical-ce1e528d8323afa6aed3789c6674b816b552dbfa27f99f5e9d8d9cc8fff3b763): complete subsection reference.

<a id="canonical-df705fa38c0927ebb693f1fcc097bb6b82558a29516e4096210740136e5f9cd0"></a>

<a id="canonical-6e7683a38907d5ac453a383be7f37925035260a20d84ec2a4d15d51af6dd8b72"></a>

## description_spec property — tls_intercept.custom_certificate / 06a69f263787 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--proxy--reference--group-005.md#canonical-f7adc0909fee90a7b778b94c15b5deb1cdcae3514cb7a8bf8ab4c78c61084b0d): complete subsection reference.

- [private_key](resources--proxy--reference--group-005.md#canonical-f9b21f551ee01ca808f0869f7492618daec3588df207af1328e9678fa68479a9): complete subsection reference.

- [use_system_defaults](resources--proxy--reference--group-005.md#canonical-239b61a00af6fe91f6b8652d8132b3f8e5ace7ca2a4a792e276e95c2a09c3f98): complete subsection reference.

<a id="canonical-2ed78faf5a082a65b46bffe0d0b3a5138622c7cfa31db70810e50a004a3618ee"></a>

## Next pages — tls_intercept.custom_certificate / 06a69f263787 / 6

- [tls_intercept.custom_certificate.custom_hash_algorithms](resources--proxy--reference--group-005.md#canonical-ce1e528d8323afa6aed3789c6674b816b552dbfa27f99f5e9d8d9cc8fff3b763)
- [tls_intercept.custom_certificate.disable_ocsp_stapling](resources--proxy--reference--group-005.md#canonical-f7adc0909fee90a7b778b94c15b5deb1cdcae3514cb7a8bf8ab4c78c61084b0d)
- [tls_intercept.custom_certificate.private_key](resources--proxy--reference--group-005.md#canonical-f9b21f551ee01ca808f0869f7492618daec3588df207af1328e9678fa68479a9)
- [tls_intercept.custom_certificate.use_system_defaults](resources--proxy--reference--group-005.md#canonical-239b61a00af6fe91f6b8652d8132b3f8e5ace7ca2a4a792e276e95c2a09c3f98)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-4639ee2404690fbd00e5ce07dfcd45d8207d50ed983459435b1b3df804692b4d)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-ce1e528d8323afa6aed3789c6674b816b552dbfa27f99f5e9d8d9cc8fff3b763"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bc12d35fc869a22ea7399bd3655f205913d6397935d1ae26ca34bf4d0e211715"></a>

## tls_intercept.custom_certificate.custom_hash_algorithms — tls_intercept.custom_certificate.custom_hash_algorithms / 5b82a459c02a / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-4639ee2404690fbd00e5ce07dfcd45d8207d50ed983459435b1b3df804692b4d)
- [tls_intercept.custom_certificate](resources--proxy--reference--group-005.md#canonical-aaee6da9f3557a4a4b7e56d1f1be1236163b7cea90779d39d051158d211b5eeb)
- tls_intercept.custom_certificate.custom_hash_algorithms

<a id="canonical-04a425e3786ff2559d5f4c9dab6750912533738667badccc3e1c5fa186a0ddc1"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("hash_algorithms")}
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
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-7135391ea6b0a62cb2a7a8e0143edca59580a26dc82836ac7a8468705f1c2afb"></a>

## Direct properties — tls_intercept.custom_certificate.custom_hash_algorithms / 5b82a459c02a / 3

<a id="canonical-5b920e6113704575c5887f60bc917bc5aea10ef0342d75a599188fa2adb6d8d6"></a>

<a id="canonical-38520ef3b15b0f626f767a32d6f896c87a09967473fd14b5e317040071ee21b8"></a>

## hash_algorithms property — tls_intercept.custom_certificate.custom_hash_algorithms / 5b82a459c02a / 4

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-95626d6dfdbfffd9691d783b8a3a317e0b1b1598cfefb9edbc3622b6e700df56"></a>

## Next pages — tls_intercept.custom_certificate.custom_hash_algorithms / 5b82a459c02a / 5

- [tls_intercept.custom_certificate](resources--proxy--reference--group-005.md#canonical-aaee6da9f3557a4a4b7e56d1f1be1236163b7cea90779d39d051158d211b5eeb)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-f7adc0909fee90a7b778b94c15b5deb1cdcae3514cb7a8bf8ab4c78c61084b0d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3add04c3690d79321b7baaec4f3ba2bb4cda772c0f707947f417b6bffaaa4b6d"></a>

## tls_intercept.custom_certificate.disable_ocsp_stapling — tls_intercept.custom_certificate.disable_ocsp_stapling / 9c43213c528d / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-4639ee2404690fbd00e5ce07dfcd45d8207d50ed983459435b1b3df804692b4d)
- [tls_intercept.custom_certificate](resources--proxy--reference--group-005.md#canonical-aaee6da9f3557a4a4b7e56d1f1be1236163b7cea90779d39d051158d211b5eeb)
- tls_intercept.custom_certificate.disable_ocsp_stapling

<a id="canonical-259f6c97fa7cb199208edcdffb53d5fdf28cf2ab47fd6e11806cf0635c80655a"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable ocsp stapling.

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
disable_ocsp_stapling = {}
```

<a id="canonical-bdff0831aca955e8bc2be756acbb40096138ad36e87bc5e94e23e316e627d3e6"></a>

## Direct properties — tls_intercept.custom_certificate.disable_ocsp_stapling / 9c43213c528d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-68f358df0c0cb34379acb14ea620ba8f4925c5c8c27b771b790eb15d0e90be26"></a>

## Next pages — tls_intercept.custom_certificate.disable_ocsp_stapling / 9c43213c528d / 4

- [tls_intercept.custom_certificate](resources--proxy--reference--group-005.md#canonical-aaee6da9f3557a4a4b7e56d1f1be1236163b7cea90779d39d051158d211b5eeb)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-f9b21f551ee01ca808f0869f7492618daec3588df207af1328e9678fa68479a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eea0e419990520cf8a4fd610a3cd7bb0bdb9466e95c91cdaad002d16f5421437"></a>

## tls_intercept.custom_certificate.private_key — tls_intercept.custom_certificate.private_key / e95ae3ea646d / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-4639ee2404690fbd00e5ce07dfcd45d8207d50ed983459435b1b3df804692b4d)
- [tls_intercept.custom_certificate](resources--proxy--reference--group-005.md#canonical-aaee6da9f3557a4a4b7e56d1f1be1236163b7cea90779d39d051158d211b5eeb)
- tls_intercept.custom_certificate.private_key

<a id="canonical-a0838f597f776706faefb2de23b7b3b9d5fbf0766c4b13bee2277ea03c898a0c"></a>

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
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-3c9e55128af7e162bc7a9f83d771a90048aa9e3f2b7cf00626f105b648ec9f5d"></a>

## Direct properties — tls_intercept.custom_certificate.private_key / e95ae3ea646d / 3

- [blindfold_secret_info](resources--proxy--reference--group-005.md#canonical-cba734536536b70b1efcacb98ba1fd36494e003fb40d0eea994da6462c488b25): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-005.md#canonical-e100cc363a1e8737076c90390ce5555aa6cb966fb289d345539b35926b207687): complete subsection reference.

<a id="canonical-2803eff2ff38a8eebc050c9b8bca621aaa86c1115bbec759be3f7f50fdff5972"></a>

## Next pages — tls_intercept.custom_certificate.private_key / e95ae3ea646d / 4

- [tls_intercept.custom_certificate.private_key.blindfold_secret_info](resources--proxy--reference--group-005.md#canonical-cba734536536b70b1efcacb98ba1fd36494e003fb40d0eea994da6462c488b25)
- [tls_intercept.custom_certificate.private_key.clear_secret_info](resources--proxy--reference--group-005.md#canonical-e100cc363a1e8737076c90390ce5555aa6cb966fb289d345539b35926b207687)
- [tls_intercept.custom_certificate](resources--proxy--reference--group-005.md#canonical-aaee6da9f3557a4a4b7e56d1f1be1236163b7cea90779d39d051158d211b5eeb)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-cba734536536b70b1efcacb98ba1fd36494e003fb40d0eea994da6462c488b25"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0959d3d4fa944c86d3583fd159e84286238607c84aedca034a20bc3a4afb25c6"></a>

## tls_intercept.custom_certificate.private_key.blindfold_secret_info — tls_intercept.custom_certificate.private_key.blindfold_secret_info / 2ddd0dcd5ff6 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-4639ee2404690fbd00e5ce07dfcd45d8207d50ed983459435b1b3df804692b4d)
- [tls_intercept.custom_certificate](resources--proxy--reference--group-005.md#canonical-aaee6da9f3557a4a4b7e56d1f1be1236163b7cea90779d39d051158d211b5eeb)
- [tls_intercept.custom_certificate.private_key](resources--proxy--reference--group-005.md#canonical-f9b21f551ee01ca808f0869f7492618daec3588df207af1328e9678fa68479a9)
- tls_intercept.custom_certificate.private_key.blindfold_secret_info

<a id="canonical-7302469c0f348762abb8a573112f4e741385284ba0449950f3c8ef76b097ce13"></a>

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

<a id="canonical-f46099554733138b78edfca76c7f0afc528bccf622bcec3a1c20bb04ef84ff45"></a>

## Direct properties — tls_intercept.custom_certificate.private_key.blindfold_secret_info / 2ddd0dcd5ff6 / 3

<a id="canonical-b5cadd47772a6ac5f8c0481a93bbc0342502ed3e28bafb7294d27075348d2507"></a>

<a id="canonical-ef68ad03e0f9adb341ba7990c5acfe1e14255902511922a1dfca4a285d87e25f"></a>

## decryption_provider property — tls_intercept.custom_certificate.private_key.blindfold_secret_info / 2ddd0dcd5ff6 / 4

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

<a id="canonical-b268037a6cb3778a520ea69182e1edb6824c6f1af40b65694dbd22d84867b78a"></a>

<a id="canonical-9575731e541e52c2aba2454e79d28c8b61c770af5c04ebd2100e4eee8369e84a"></a>

## location property — tls_intercept.custom_certificate.private_key.blindfold_secret_info / 2ddd0dcd5ff6 / 5

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

<a id="canonical-f4fd8e4a497a694ebd3047940c5dabced2d9f019559d2c1806cfc4a3bff57461"></a>

<a id="canonical-b89eb912145c20553657b209c586af0a28aec97046a99add59e7310e1b83e27c"></a>

## store_provider property — tls_intercept.custom_certificate.private_key.blindfold_secret_info / 2ddd0dcd5ff6 / 6

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

<a id="canonical-152349cc1f9cbe8c1a9ba984b7e2a73cd2a48bba96f20b55f5d44d6c6715339c"></a>

## Next pages — tls_intercept.custom_certificate.private_key.blindfold_secret_info / 2ddd0dcd5ff6 / 7

- [tls_intercept.custom_certificate.private_key](resources--proxy--reference--group-005.md#canonical-f9b21f551ee01ca808f0869f7492618daec3588df207af1328e9678fa68479a9)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-e100cc363a1e8737076c90390ce5555aa6cb966fb289d345539b35926b207687"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0706550d23b717ef80d4f592ff074a0d786a7a58f73dd176c63503e4f5d2c3a7"></a>

## tls_intercept.custom_certificate.private_key.clear_secret_info — tls_intercept.custom_certificate.private_key.clear_secret_info / 6079dfa4b003 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-4639ee2404690fbd00e5ce07dfcd45d8207d50ed983459435b1b3df804692b4d)
- [tls_intercept.custom_certificate](resources--proxy--reference--group-005.md#canonical-aaee6da9f3557a4a4b7e56d1f1be1236163b7cea90779d39d051158d211b5eeb)
- [tls_intercept.custom_certificate.private_key](resources--proxy--reference--group-005.md#canonical-f9b21f551ee01ca808f0869f7492618daec3588df207af1328e9678fa68479a9)
- tls_intercept.custom_certificate.private_key.clear_secret_info

<a id="canonical-e581771cbb99e76684adc0a764970febb0a06454b786613284c02e3fe91243a4"></a>

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

<a id="canonical-39cb58e147c4b6b0bee93ef11c0cf8a1a937d328d3335d318be1dff96563a1a2"></a>

## Direct properties — tls_intercept.custom_certificate.private_key.clear_secret_info / 6079dfa4b003 / 3

<a id="canonical-e8bc692703edf76a89430b9de1ef750d3c9f6f7517981bf6ef9e8392c4a8bda3"></a>

<a id="canonical-01e9660351e3a4f121da029f16b3b7dc51f14b08c498c41d3772ad056e9e4142"></a>

## provider_ref property — tls_intercept.custom_certificate.private_key.clear_secret_info / 6079dfa4b003 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-a72935985b94ea510f9a705849f912e48d0fb34468d13e013f02440a2d600140"></a>

<a id="canonical-334602156e8446e29f428f6cb69940d3f3605acc67da573b79c8cb9866e3098d"></a>

## url property — tls_intercept.custom_certificate.private_key.clear_secret_info / 6079dfa4b003 / 5

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

<a id="canonical-55bb832337955f8dede0c37e2a07e901140bc8d64c16501abff1af86b76dbd7a"></a>

## Next pages — tls_intercept.custom_certificate.private_key.clear_secret_info / 6079dfa4b003 / 6

- [tls_intercept.custom_certificate.private_key](resources--proxy--reference--group-005.md#canonical-f9b21f551ee01ca808f0869f7492618daec3588df207af1328e9678fa68479a9)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-239b61a00af6fe91f6b8652d8132b3f8e5ace7ca2a4a792e276e95c2a09c3f98"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-186d706706c1bc1032f438a9275037ab83b0e4f904710e98946cea67f5f20b2e"></a>

## tls_intercept.custom_certificate.use_system_defaults — tls_intercept.custom_certificate.use_system_defaults / 0e1015e185d6 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-4639ee2404690fbd00e5ce07dfcd45d8207d50ed983459435b1b3df804692b4d)
- [tls_intercept.custom_certificate](resources--proxy--reference--group-005.md#canonical-aaee6da9f3557a4a4b7e56d1f1be1236163b7cea90779d39d051158d211b5eeb)
- tls_intercept.custom_certificate.use_system_defaults

<a id="canonical-548fe943c83be548c70aaa3b4e8599f80a1532f7eda17cdaa52d91b58054e641"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use system defaults.

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
use_system_defaults = {}
```

<a id="canonical-01870d50df0ce5a5606f16b9cb39e074a6bf0b50070eceef589c64e29801b31b"></a>

## Direct properties — tls_intercept.custom_certificate.use_system_defaults / 0e1015e185d6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c49ad5d2de6ab1681ef6b22bae2439a3af50a3c6c782918b69374d140b33cf50"></a>

## Next pages — tls_intercept.custom_certificate.use_system_defaults / 0e1015e185d6 / 4

- [tls_intercept.custom_certificate](resources--proxy--reference--group-005.md#canonical-aaee6da9f3557a4a4b7e56d1f1be1236163b7cea90779d39d051158d211b5eeb)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-04da480c4e48c82947b4b5552704c071f296e48a84a79e457875c2106cd72292"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72ff1fb6fcbf6e5a935dae3f4e3cc27fc6a83e93d036a59666b79c37e08848e2"></a>

## tls_intercept.enable_for_all_domains — tls_intercept.enable_for_all_domains / 7a0c2a1ed4a4 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-4639ee2404690fbd00e5ce07dfcd45d8207d50ed983459435b1b3df804692b4d)
- tls_intercept.enable_for_all_domains

<a id="canonical-f33fd7f2323687d0d1972b7f1f614e7c96f6cbae82067d599aff54c1cda7de22"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable for all domains.

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
enable_for_all_domains = {}
```

<a id="canonical-9921ec308f584154d5702d166b6efcd20513fd26eb2791d5db06f2c4a4ce1c93"></a>

## Direct properties — tls_intercept.enable_for_all_domains / 7a0c2a1ed4a4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7fa7fff0774b2eea3386079e4ce483955ec067990c6eabe121e5891be32426b2"></a>

## Next pages — tls_intercept.enable_for_all_domains / 7a0c2a1ed4a4 / 4

- [tls_intercept](resources--proxy--reference--group-005.md#canonical-4639ee2404690fbd00e5ce07dfcd45d8207d50ed983459435b1b3df804692b4d)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-97d5300e2bd77f1f1170e86f5d3b0925b166d8e616fe40b3eeff08370d847453"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d0e1852d6ab3598667a4bf2c4345b3861ccaaa1840950070d29b0ac11bb71f8"></a>

## tls_intercept.policy — tls_intercept.policy / afeec8fd2d17 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-4639ee2404690fbd00e5ce07dfcd45d8207d50ed983459435b1b3df804692b4d)
- tls_intercept.policy

<a id="canonical-c540da6563564ab3e14444da469dcd0d5d4a31e560a8f7efdb21002ec1410e30"></a>

Type: `"object"`. single nested block, Optional.

Policy to enable or disable TLS interception.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("interception_rules")}
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
policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-556f2107e36e526b786031ca92bcb55204c4fcf68aaddde6c1cd4113f361da1c"></a>

## Direct properties — tls_intercept.policy / afeec8fd2d17 / 3

- [interception_rules](resources--proxy--reference--group-005.md#canonical-f58883ddeb996fb9687611157effae795585ee63f81f86ace502ef20b439878e): complete subsection reference.

<a id="canonical-5099f742ba22dd02e4bdf0568a0eb33e9816f7822efb1d781705a24ea42edc2d"></a>

## Next pages — tls_intercept.policy / afeec8fd2d17 / 4

- [tls_intercept.policy.interception_rules](resources--proxy--reference--group-005.md#canonical-f58883ddeb996fb9687611157effae795585ee63f81f86ace502ef20b439878e)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-4639ee2404690fbd00e5ce07dfcd45d8207d50ed983459435b1b3df804692b4d)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-f58883ddeb996fb9687611157effae795585ee63f81f86ace502ef20b439878e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cb4368ea209ba18ee32021fc94624c1efb02deaf2026a3fadaa75153524804e5"></a>

## tls_intercept.policy.interception_rules — tls_intercept.policy.interception_rules / db2d6610910f / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-4639ee2404690fbd00e5ce07dfcd45d8207d50ed983459435b1b3df804692b4d)
- [tls_intercept.policy](resources--proxy--reference--group-005.md#canonical-97d5300e2bd77f1f1170e86f5d3b0925b166d8e616fe40b3eeff08370d847453)
- tls_intercept.policy.interception_rules

<a id="canonical-f660b514493ec2f7c63b152d36fa33187223cb088752e03de2dc5ee62214ecbb"></a>

Type: `"object"`. list nested block, Optional.

List of ordered rules to enable or disable for TLS interception.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("disable_interception",
    "enable_interception")}
```

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
interception_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-4f2381d01485367f31811dda499dac4615352ed07da6a0d43b637cfe24836372"></a>

## Direct properties — tls_intercept.policy.interception_rules / db2d6610910f / 3

- [disable_interception](resources--proxy--reference--group-005.md#canonical-405bd497e6976dec32f4f169756dd4732583a70f7e628e04b7adb3f1365dd539): complete subsection reference.

- [domain_match](resources--proxy--reference--group-005.md#canonical-4b62ce6a4f9471f00b4c3c44904822303e92e0ea713a05f1eaa3b9fca9a509f8): complete subsection reference.

- [enable_interception](resources--proxy--reference--group-005.md#canonical-5bdaab16a543de7bcab7d3755d7a92d5b13477bfe1a22b75b4966929bc73decf): complete subsection reference.

<a id="canonical-f418ede8ea3e1a3a0c60edfb76b2af6a78c7788903e750a2f2c4eb8ad82704ae"></a>

## Next pages — tls_intercept.policy.interception_rules / db2d6610910f / 4

- [tls_intercept.policy.interception_rules.disable_interception](resources--proxy--reference--group-005.md#canonical-405bd497e6976dec32f4f169756dd4732583a70f7e628e04b7adb3f1365dd539)
- [tls_intercept.policy.interception_rules.domain_match](resources--proxy--reference--group-005.md#canonical-4b62ce6a4f9471f00b4c3c44904822303e92e0ea713a05f1eaa3b9fca9a509f8)
- [tls_intercept.policy.interception_rules.enable_interception](resources--proxy--reference--group-005.md#canonical-5bdaab16a543de7bcab7d3755d7a92d5b13477bfe1a22b75b4966929bc73decf)
- [tls_intercept.policy](resources--proxy--reference--group-005.md#canonical-97d5300e2bd77f1f1170e86f5d3b0925b166d8e616fe40b3eeff08370d847453)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-405bd497e6976dec32f4f169756dd4732583a70f7e628e04b7adb3f1365dd539"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2b575cfb8185745a64a0a1999b29e658b5f8b83856afe6f955a06bc82433768"></a>

## tls_intercept.policy.interception_rules.disable_interception — tls_intercept.policy.interception_rules.disable_interception / f437a1016037 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-4639ee2404690fbd00e5ce07dfcd45d8207d50ed983459435b1b3df804692b4d)
- [tls_intercept.policy](resources--proxy--reference--group-005.md#canonical-97d5300e2bd77f1f1170e86f5d3b0925b166d8e616fe40b3eeff08370d847453)
- [tls_intercept.policy.interception_rules](resources--proxy--reference--group-005.md#canonical-f58883ddeb996fb9687611157effae795585ee63f81f86ace502ef20b439878e)
- tls_intercept.policy.interception_rules.disable_interception

<a id="canonical-1da008e286394eb45cd978c1e280dac995581cf8cd9d94e6a6e1f8664a2253c5"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable interception.

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
disable_interception = {}
```

<a id="canonical-af22d05389efc47755f3ea27345235ab720fe2bf2c11b290eaaa14a4b8031463"></a>

## Direct properties — tls_intercept.policy.interception_rules.disable_interception / f437a1016037 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9366d68f9dd5bebffc449af308be9f34aae841c6dec7ef6f1dfb1503a4285041"></a>

## Next pages — tls_intercept.policy.interception_rules.disable_interception / f437a1016037 / 4

- [tls_intercept.policy.interception_rules](resources--proxy--reference--group-005.md#canonical-f58883ddeb996fb9687611157effae795585ee63f81f86ace502ef20b439878e)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-4b62ce6a4f9471f00b4c3c44904822303e92e0ea713a05f1eaa3b9fca9a509f8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a0343948d03a43d8d119303b1f058d0c50f979df399bd634beaf4eab3c305d1"></a>

## tls_intercept.policy.interception_rules.domain_match — tls_intercept.policy.interception_rules.domain_match / 7653c3b858cd / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-4639ee2404690fbd00e5ce07dfcd45d8207d50ed983459435b1b3df804692b4d)
- [tls_intercept.policy](resources--proxy--reference--group-005.md#canonical-97d5300e2bd77f1f1170e86f5d3b0925b166d8e616fe40b3eeff08370d847453)
- [tls_intercept.policy.interception_rules](resources--proxy--reference--group-005.md#canonical-f58883ddeb996fb9687611157effae795585ee63f81f86ace502ef20b439878e)
- tls_intercept.policy.interception_rules.domain_match

<a id="canonical-2ab861b946294eca9c22691deab90ed1143fefd16777a82dbdfd77955902cd82"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for domain match.

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
domain_match {
  # Configure direct properties listed below.
}
```

<a id="canonical-f31a32c5e162b5a35ec67bb5edb5cfaac12c689caccac7ca1b255fc873f296e1"></a>

## Direct properties — tls_intercept.policy.interception_rules.domain_match / 7653c3b858cd / 3

<a id="canonical-e1034094d564137d68a4d6717729b30f0802bc08cb3c29efccc56bdd1b3ebaa3"></a>

<a id="canonical-ba67bdd3a9e4a9ccc147bd6623e9fadb0dc6d29313d1e4b76a52573af666cb42"></a>

## exact_value property — tls_intercept.policy.interception_rules.domain_match / 7653c3b858cd / 4

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

<a id="canonical-6eeb230b9e9e875c0ecf37c1b510c8a071b25d7baa11e6534bb49ec78df9d34c"></a>

<a id="canonical-613e0368c4171aab848e02835582111fdd8ee6eb821aa6ea6dd9d4dc3627e93c"></a>

## regex_value property — tls_intercept.policy.interception_rules.domain_match / 7653c3b858cd / 5

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

<a id="canonical-2cf09a16582aa892704993fa2d0509cc55e19429d8a6fb2405499e69c26733af"></a>

<a id="canonical-d27d01d04cb1375941e244b4aac11cb4a602e35ea3043ccf4c2ed11c1b6cbab3"></a>

## suffix_value property — tls_intercept.policy.interception_rules.domain_match / 7653c3b858cd / 6

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

<a id="canonical-d3faa081162b363b5f4a5e8fd613f0f1e8174b4d39dbe86145d7f50ce466ee6b"></a>

## Next pages — tls_intercept.policy.interception_rules.domain_match / 7653c3b858cd / 7

- [tls_intercept.policy.interception_rules](resources--proxy--reference--group-005.md#canonical-f58883ddeb996fb9687611157effae795585ee63f81f86ace502ef20b439878e)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-5bdaab16a543de7bcab7d3755d7a92d5b13477bfe1a22b75b4966929bc73decf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d7fcf26a86fae1dc24e941619eef8dafafc913532ff918e4431b814390c81f5f"></a>

## tls_intercept.policy.interception_rules.enable_interception — tls_intercept.policy.interception_rules.enable_interception / ee79fe240999 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-4639ee2404690fbd00e5ce07dfcd45d8207d50ed983459435b1b3df804692b4d)
- [tls_intercept.policy](resources--proxy--reference--group-005.md#canonical-97d5300e2bd77f1f1170e86f5d3b0925b166d8e616fe40b3eeff08370d847453)
- [tls_intercept.policy.interception_rules](resources--proxy--reference--group-005.md#canonical-f58883ddeb996fb9687611157effae795585ee63f81f86ace502ef20b439878e)
- tls_intercept.policy.interception_rules.enable_interception

<a id="canonical-c7023f5cc3272a4896ffb84105d795d302a4130ab69ec9c43684f8b9edcb1602"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable interception.

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
enable_interception = {}
```

<a id="canonical-0a758fc142fe32c2c8058ae2c92ba11f4d447a22261176e4a3873b08e83d0dd5"></a>

## Direct properties — tls_intercept.policy.interception_rules.enable_interception / ee79fe240999 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-44eb9caa07423792890608743ab4161f10a0c258e6bfa71c89c14713b46b5b1a"></a>

## Next pages — tls_intercept.policy.interception_rules.enable_interception / ee79fe240999 / 4

- [tls_intercept.policy.interception_rules](resources--proxy--reference--group-005.md#canonical-f58883ddeb996fb9687611157effae795585ee63f81f86ace502ef20b439878e)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-3a866c24dab44d630da87b83d903f073888f999314fee1a280191336a46835c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d39f0030b87504b80c42900c8ac1ecebee025a3cf2272266b50b35e3913b6a1"></a>

## tls_intercept.volterra_certificate — tls_intercept.volterra_certificate / 13c8ff54f7c6 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-4639ee2404690fbd00e5ce07dfcd45d8207d50ed983459435b1b3df804692b4d)
- tls_intercept.volterra_certificate

<a id="canonical-99682d5181e7da867e7689103b8933908533ea81d6f5f2ea72604302354cc287"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for volterra certificate.

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
volterra_certificate = {}
```

<a id="canonical-974579f5388cadbf0e48095073aa7c961ca6e521fa312a4a765e285fc639deb7"></a>

## Direct properties — tls_intercept.volterra_certificate / 13c8ff54f7c6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3da1a4f16e8e579dfddec3e0b54035e5130a29ef9436dd68cf642cff1fdb173b"></a>

## Next pages — tls_intercept.volterra_certificate / 13c8ff54f7c6 / 4

- [tls_intercept](resources--proxy--reference--group-005.md#canonical-4639ee2404690fbd00e5ce07dfcd45d8207d50ed983459435b1b3df804692b4d)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-296065aeea4832e5be6edf994ae5cbf0debd9ea52458ec265063b0ecff6e5441"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-807888a7b1e7ad035158f7d1c770a0667b71fb3117aa8f84ed1a14d7e6692659"></a>

## tls_intercept.volterra_trusted_ca — tls_intercept.volterra_trusted_ca / dc0fb4cfb44e / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-4639ee2404690fbd00e5ce07dfcd45d8207d50ed983459435b1b3df804692b4d)
- tls_intercept.volterra_trusted_ca

<a id="canonical-4762ccab0f0ca1c0ab12ae73896fdfda5af2967b72695a47f24a318a0308ce97"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for volterra trusted ca.

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
volterra_trusted_ca = {}
```

<a id="canonical-ea105d376b0204cb2d159849b9042848a359d43071c7532900f37c73007aff2a"></a>

## Direct properties — tls_intercept.volterra_trusted_ca / dc0fb4cfb44e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a46d0711191f83ebef6ea3ce3129a6a9cb686e1e8f59297dfa99c69c3cdecf6e"></a>

## Next pages — tls_intercept.volterra_trusted_ca / dc0fb4cfb44e / 4

- [tls_intercept](resources--proxy--reference--group-005.md#canonical-4639ee2404690fbd00e5ce07dfcd45d8207d50ed983459435b1b3df804692b4d)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
