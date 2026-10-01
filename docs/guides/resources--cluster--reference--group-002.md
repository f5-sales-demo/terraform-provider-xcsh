---
page_title: "xcsh_cluster reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cluster reference."
---

# xcsh_cluster reference

<a id="canonical-54fac20bda08ba8b7307d5f2dc07840c0e81343e942d7cca3a5241665682027b"></a>

## tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / ccaba1c8e209 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-e891a37ad95c6abd2732ec43824318ac4c40db0eaa681d108a9e72134d43fd0d)
- [tls_parameters.common_params](resources--cluster--reference--group-001.md#canonical-c0f92b4dae58a6cd257ee8960e68808fa66b0542b22cc0c62afe84cb381f66cd)
- [tls_parameters.common_params.validation_params](resources--cluster--reference--group-001.md#canonical-5df2cdb334c28e2a539dd2549ae97ddc8ccd52d91c6283b4ff928859fd0ccc93)
- [tls_parameters.common_params.validation_params.trusted_ca](resources--cluster--reference--group-001.md#canonical-9b219cc99447098cc906f098d54d9b188b41109365a80a24aed55db0b7aa516c)
- tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list

<a id="canonical-776a519ebb4c104526135872694f216de76fef866b5faef2133c4fecb549e86f"></a>

Type: `"object"`. list nested block, Optional.

Root CA Certificate Reference. Reference to Root CA Certificate.

Upstream description:

Reference to Root CA Certificate.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
trusted_ca_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-194fb87135944d981cd4bef9279239776e0f85843284e65e8221e30435b3cda6"></a>

## Direct properties — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / ccaba1c8e209 / 3

<a id="canonical-c64f97c45f7f33804dd32009753dea14692f433582bb4601f82fbddc91e8b2ae"></a>

<a id="canonical-cea7e3c90d57646632bce38d1d747e0375f28a76a3a13e5db29fa7ab4a16789b"></a>

## kind property — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / ccaba1c8e209 / 4

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

<a id="canonical-91fdb9da0b3d68ab472bc939127f2db621af3d9bc63a8e682e1049f4b28e4c04"></a>

<a id="canonical-6d41960aeb8f7d0b407714f68906ae6155f9dc10d29bfc9998c13ac1466c8831"></a>

## name property — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / ccaba1c8e209 / 5

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

<a id="canonical-404eeb2124262c10cc48f2073d86ab7ab56b6cfde2297c7c0561d8748e2d1974"></a>

<a id="canonical-3c8fdbccaeddd6f453d1c9395e300594303db8d208282d8fbf71bf026857c37f"></a>

## namespace property — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / ccaba1c8e209 / 6

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

<a id="canonical-1a4beef95ca51a95e6d8fdaa82cb9f18d1843c97cfb82d9478eb83361a936527"></a>

<a id="canonical-ef50d7ac0b7cfeb84b97ac0b11ed497dab4bd056f1d003d0c5342f2cf36b5aac"></a>

## tenant property — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / ccaba1c8e209 / 7

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

<a id="canonical-f37dca176e42eb6bcbb30d7ada0b4d8c74bb67e30d088771ea8bbbf43fd9419a"></a>

<a id="canonical-b8a5943ae2410945548e773d36f487d1bc6156af2710cd39b7a9d8ede553f6fc"></a>

## uid property — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / ccaba1c8e209 / 8

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

<a id="canonical-2a82684300855339164aaddd3cf8e9ac5ffc082477c7b91d39705c78065b48ac"></a>

## Next pages — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / ccaba1c8e209 / 9

- [tls_parameters.common_params.validation_params.trusted_ca](resources--cluster--reference--group-001.md#canonical-9b219cc99447098cc906f098d54d9b188b41109365a80a24aed55db0b7aa516c)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-d5938793cbc4f992c49a04eb7cd82d52a334f576743046a32b760b559c56a81f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-62b7bdc0f3688b2b4cecb8c2d1045c2802e60aa92adcbad2c12cc8dfa6ae2dc6"></a>

## tls_parameters.default_session_key_caching — tls_parameters.default_session_key_caching / 2e417f47b958 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-e891a37ad95c6abd2732ec43824318ac4c40db0eaa681d108a9e72134d43fd0d)
- tls_parameters.default_session_key_caching

<a id="canonical-332a750670cdb85df34e3b3b0a6f7e1d2490570aebb0a1e175e7bf2f449b66f2"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default session key caching.

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
default_session_key_caching = {}
```

<a id="canonical-4a06a6f4d89f4473cc8c099f53a9fbe05f2b72ee74f5a29a18d8ad1482aa3b04"></a>

## Direct properties — tls_parameters.default_session_key_caching / 2e417f47b958 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ef719aeab65a1f89cbdc6b736a85b0d866613ef4d2000dc8cf5bddc63158d1db"></a>

## Next pages — tls_parameters.default_session_key_caching / 2e417f47b958 / 4

- [tls_parameters](resources--cluster--reference--group-001.md#canonical-e891a37ad95c6abd2732ec43824318ac4c40db0eaa681d108a9e72134d43fd0d)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-92ec7b9b9a70b93cab15f9d7c462b9b4cfa2d511834f0888f5363db22418f5f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf452238cf5815f3a6b596fc67cdb2357e887751ac4a69b8ed700e97c80c236a"></a>

## tls_parameters.disable_session_key_caching — tls_parameters.disable_session_key_caching / 048b7d1c409a / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-e891a37ad95c6abd2732ec43824318ac4c40db0eaa681d108a9e72134d43fd0d)
- tls_parameters.disable_session_key_caching

<a id="canonical-b3c4b2e7f5fd767142ae13637e73d0cf3a5ddc7fc745da2eef570b922f6b6245"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable session key caching.

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
disable_session_key_caching = {}
```

<a id="canonical-9291ff9ddecc29e56195397612d3ff1fb784f0548d4711d7651679504dfb2dd2"></a>

## Direct properties — tls_parameters.disable_session_key_caching / 048b7d1c409a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4e3cf73582960b5a80cb2998b26fc6e1a8c9796bce22fb79ca27966e2729288d"></a>

## Next pages — tls_parameters.disable_session_key_caching / 048b7d1c409a / 4

- [tls_parameters](resources--cluster--reference--group-001.md#canonical-e891a37ad95c6abd2732ec43824318ac4c40db0eaa681d108a9e72134d43fd0d)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-6d84cd0eae7dcfacd28f41de2d59fbe2960aa31c3666b0f7722e3f2d3416ee9e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-84a084927a7a1bb382ad3abd81f7a7b0a597b39382d193eb9842fe3d34b0507c"></a>

## tls_parameters.disable_sni — tls_parameters.disable_sni / 44485fc6f51b / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-e891a37ad95c6abd2732ec43824318ac4c40db0eaa681d108a9e72134d43fd0d)
- tls_parameters.disable_sni

<a id="canonical-c5351cb89ee4e58ee13243ba03bd70661fb273440ddc816f0103f43283b09e27"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable sni.

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
disable_sni = {}
```

<a id="canonical-15fbfe45b9df4d0ef57a3ba15d8939aa5c45a398947a19ad67a5b720a85621e4"></a>

## Direct properties — tls_parameters.disable_sni / 44485fc6f51b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-65e61efca6e1000edadc7e8f6841a5a1f94340130a668fd4b85f871583b58f8b"></a>

## Next pages — tls_parameters.disable_sni / 44485fc6f51b / 4

- [tls_parameters](resources--cluster--reference--group-001.md#canonical-e891a37ad95c6abd2732ec43824318ac4c40db0eaa681d108a9e72134d43fd0d)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-f3317e3237accb9a6c0b6e3b1684f389cecf00eacc6e283dce291bbf3cf7e30c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cae7327331a746888530b39c5445c4e18159b98f30c5d24f570c101729a8d8c3"></a>

## tls_parameters.use_host_header_as_sni — tls_parameters.use_host_header_as_sni / 93c7a1bdb99f / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-e891a37ad95c6abd2732ec43824318ac4c40db0eaa681d108a9e72134d43fd0d)
- tls_parameters.use_host_header_as_sni

<a id="canonical-869a973133fb15ecd6f9aa17a836a35914a65f15d528f0ae86823e6f84734bfb"></a>

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
use_host_header_as_sni = {}
```

<a id="canonical-b65390b52d93f1a5c6c90b9502cb49331279f4f141f3dc35573d4b8e6477cfb2"></a>

## Direct properties — tls_parameters.use_host_header_as_sni / 93c7a1bdb99f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fbabe04e4c375a58d0b44747de0cebbbaa999fdfcd4e8fb31a93a006ca36b2ed"></a>

## Next pages — tls_parameters.use_host_header_as_sni / 93c7a1bdb99f / 4

- [tls_parameters](resources--cluster--reference--group-001.md#canonical-e891a37ad95c6abd2732ec43824318ac4c40db0eaa681d108a9e72134d43fd0d)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-9670852cd96a827e1246fe6fc573c2dedb956c77b7b1717fa85237418a3d9522"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8f41b12ad661a3f54ca7cb32fe1554bef2b4b32900d12141d26f96053727551"></a>

## upstream_conn_pool_reuse_type — upstream_conn_pool_reuse_type / d40f7ed865f5 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- upstream_conn_pool_reuse_type

<a id="canonical-9bfa6c40a0109eebbc9d72592925401c0144c4f71f455e533ac7d3d7190009a8"></a>

Type: `"object"`. single nested block, Optional.

Select upstream connection pool reuse state for every downstream connection. This configuration
choice is for HTTP(S) LB only.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_conn_pool_reuse",
    "enable_conn_pool_reuse")}
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
  "x-ves-oneof-field-map_downstream_to_upstream_conn_pool_type": "[\"disable_conn_pool_reuse\",\"enable_conn_pool_reuse\"]"
}
```

Terraform syntax:

```terraform
upstream_conn_pool_reuse_type {
  # Configure direct properties listed below.
}
```

<a id="canonical-f7038757d6f1916a7c0f7a29a8a8f7ceb2377b6f6b6becc5379555a85af7773d"></a>

## Direct properties — upstream_conn_pool_reuse_type / d40f7ed865f5 / 3

- [disable_conn_pool_reuse](resources--cluster--reference--group-002.md#canonical-fb2e809a71a04d2d4d1e0b1790d2d31b976506808f8c73af28aae5056ab54e4a): complete subsection reference.

- [enable_conn_pool_reuse](resources--cluster--reference--group-002.md#canonical-b3b0d79129d8e8adddfe2680399281f4862881b4a39625666bde4f961dc2f963): complete subsection reference.

<a id="canonical-cfa1c5fa031eb12476789ddf51ffe47b47e575c1bc7e8a2b69d0949423abc19a"></a>

## Next pages — upstream_conn_pool_reuse_type / d40f7ed865f5 / 4

- [upstream_conn_pool_reuse_type.disable_conn_pool_reuse](resources--cluster--reference--group-002.md#canonical-fb2e809a71a04d2d4d1e0b1790d2d31b976506808f8c73af28aae5056ab54e4a)
- [upstream_conn_pool_reuse_type.enable_conn_pool_reuse](resources--cluster--reference--group-002.md#canonical-b3b0d79129d8e8adddfe2680399281f4862881b4a39625666bde4f961dc2f963)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-fb2e809a71a04d2d4d1e0b1790d2d31b976506808f8c73af28aae5056ab54e4a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fde566fd7e33f9ee55ab9308634b0bb56cad5b100b0245131782d8776c28e72d"></a>

## upstream_conn_pool_reuse_type.disable_conn_pool_reuse — upstream_conn_pool_reuse_type.disable_conn_pool_reuse / 40c3afb800a6 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [upstream_conn_pool_reuse_type](resources--cluster--reference--group-002.md#canonical-9670852cd96a827e1246fe6fc573c2dedb956c77b7b1717fa85237418a3d9522)
- upstream_conn_pool_reuse_type.disable_conn_pool_reuse

<a id="canonical-37eef0fa3226f945d5b34dacd3c189172682ff9112adbe283ddc80b8d7bf353c"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable conn pool reuse.

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
disable_conn_pool_reuse = {}
```

<a id="canonical-e6866377ae237928fc3790db2f21da94318ea4eb58e36a003f9775b92f19015f"></a>

## Direct properties — upstream_conn_pool_reuse_type.disable_conn_pool_reuse / 40c3afb800a6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-93dc11d0001820b76f41b13283f13c6c1560f194c0515b131669dfc05cffcf55"></a>

## Next pages — upstream_conn_pool_reuse_type.disable_conn_pool_reuse / 40c3afb800a6 / 4

- [upstream_conn_pool_reuse_type](resources--cluster--reference--group-002.md#canonical-9670852cd96a827e1246fe6fc573c2dedb956c77b7b1717fa85237418a3d9522)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-b3b0d79129d8e8adddfe2680399281f4862881b4a39625666bde4f961dc2f963"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dea4a75504cf547a7719bfe152ae0a43f9048f4ef7ca676c350fa6f7f8a460af"></a>

## upstream_conn_pool_reuse_type.enable_conn_pool_reuse — upstream_conn_pool_reuse_type.enable_conn_pool_reuse / fdd81e20ffe9 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [upstream_conn_pool_reuse_type](resources--cluster--reference--group-002.md#canonical-9670852cd96a827e1246fe6fc573c2dedb956c77b7b1717fa85237418a3d9522)
- upstream_conn_pool_reuse_type.enable_conn_pool_reuse

<a id="canonical-99e4afe2af583441efae158d61e4138c88183f8ff71309dc9b02fff1370fc0df"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable conn pool reuse.

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
enable_conn_pool_reuse = {}
```

<a id="canonical-5a24f22becf94702bb500526d839ee8d6cebb417ac1cfec5f179dd0514ab4e88"></a>

## Direct properties — upstream_conn_pool_reuse_type.enable_conn_pool_reuse / fdd81e20ffe9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-405ac3446e188fb35e1d30612795460b8ceccb8b68cc16b6b87e34da0a929af4"></a>

## Next pages — upstream_conn_pool_reuse_type.enable_conn_pool_reuse / fdd81e20ffe9 / 4

- [upstream_conn_pool_reuse_type](resources--cluster--reference--group-002.md#canonical-9670852cd96a827e1246fe6fc573c2dedb956c77b7b1717fa85237418a3d9522)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
