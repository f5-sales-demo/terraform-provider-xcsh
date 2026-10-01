---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-689fb1724f8f3c25cb5bb3dd00dd9a53eec3a15990aac407d541d366df1962e2"></a>

## enable_api_discovery.api_discovery_from_code_scan.code_base_integrations — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations / 9023923937c5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-c92592437695ce4e484e16cf1b000e2558346e492b9be3ef743d389861095739)
- [enable_api_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-017.md#canonical-8a958211739e545afd1a84be8278b9eb87cd211a15e499b08f6fd6775b543169)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations

<a id="canonical-6b884cbd2f4b03f8b0822e55ee1a133c077c1ec3bc4cb4c481d773144ddc41d3"></a>

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

<a id="canonical-7d25f57c9755a63776adb1a88d93b919faa3ea445e225cecd4ec757af9346e5a"></a>

## Direct properties — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations / 9023923937c5 / 3

- [all_repos](resources--http_loadbalancer--reference--group-018.md#canonical-86b463fbb763cc0cd7a01074b4da193f3c16bb91a687de424309c08d50e938d1): complete subsection reference.

- [code_base_integration](resources--http_loadbalancer--reference--group-018.md#canonical-c521f2e34810d948e4a30a494c0ce946f8e5e454a06d6174db1cc10056f6b172): complete subsection reference.

- [selected_repos](resources--http_loadbalancer--reference--group-018.md#canonical-7ca524a71c2c9ca664ee8ac02047836da5f37907fc3d269e45482e53c73e71fe): complete subsection reference.

<a id="canonical-2782497bd67a177a4408c4dce196e97f8278fa38c14c296a82491ef79f645379"></a>

## Next pages — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations / 9023923937c5 / 4

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos](resources--http_loadbalancer--reference--group-018.md#canonical-86b463fbb763cc0cd7a01074b4da193f3c16bb91a687de424309c08d50e938d1)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration](resources--http_loadbalancer--reference--group-018.md#canonical-c521f2e34810d948e4a30a494c0ce946f8e5e454a06d6174db1cc10056f6b172)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos](resources--http_loadbalancer--reference--group-018.md#canonical-7ca524a71c2c9ca664ee8ac02047836da5f37907fc3d269e45482e53c73e71fe)
- [enable_api_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-017.md#canonical-8a958211739e545afd1a84be8278b9eb87cd211a15e499b08f6fd6775b543169)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-86b463fbb763cc0cd7a01074b4da193f3c16bb91a687de424309c08d50e938d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3f8029dbf150a6a76e30d9d536b63bb1c86d0e4e0c79412702643fe69b1b3c66"></a>

## enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_rep / 2f793a7da054 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-c92592437695ce4e484e16cf1b000e2558346e492b9be3ef743d389861095739)
- [enable_api_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-017.md#canonical-8a958211739e545afd1a84be8278b9eb87cd211a15e499b08f6fd6775b543169)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-017.md#canonical-06644f1edd1899894148c2f317ecb380c5ef66abd2f2f0786f21206bf906d7b3)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos

<a id="canonical-452c48089daf638df663dd602b4c87af0857168d4daf38bddda72313c9a96bec"></a>

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

<a id="canonical-fb976fe114f7b4120a5217c721f4c8bca77b4ea9f106b675e88194ef9b97aa86"></a>

## Direct properties — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_rep / 2f793a7da054 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3f3fb432abf0ae14143aae46d38f70fea06242b98d5d1bbd6fb9bef7717c58b3"></a>

## Next pages — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_rep / 2f793a7da054 / 4

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-017.md#canonical-06644f1edd1899894148c2f317ecb380c5ef66abd2f2f0786f21206bf906d7b3)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c521f2e34810d948e4a30a494c0ce946f8e5e454a06d6174db1cc10056f6b172"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f34fb8fb5a755c28f71f0ca938a77471b896f1a94772209d976ea4631b1588cf"></a>

## enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_ba / 5efac13c53bf / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-c92592437695ce4e484e16cf1b000e2558346e492b9be3ef743d389861095739)
- [enable_api_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-017.md#canonical-8a958211739e545afd1a84be8278b9eb87cd211a15e499b08f6fd6775b543169)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-017.md#canonical-06644f1edd1899894148c2f317ecb380c5ef66abd2f2f0786f21206bf906d7b3)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration

<a id="canonical-a7049a10260ad83f3e7492ba7ef2066cf977c7eb4516d054ec9e170dd7708739"></a>

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

<a id="canonical-7e7c3120af9baafda8efe8e38e24770c7072ebbce27b56753270fc9e09ac5db8"></a>

## Direct properties — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_ba / 5efac13c53bf / 3

<a id="canonical-da4fd4e3b089531b83d6ea0426d84c306453e77b8df775fe2b82e352eeb0667b"></a>

<a id="canonical-f8c139aecaeefd4daa035e8637dde3833085531ff97c94b0c3c3b462630484b8"></a>

## name property — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_ba / 5efac13c53bf / 4

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

<a id="canonical-fb97daae3e94eed2360133870cc631ea3fb3452bb06a639e7eadbc10c4feeaac"></a>

<a id="canonical-20cc856bbf9bd6ebde21dbabf665ff157a3231bceb5ff0c50d61b56664273755"></a>

## namespace property — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_ba / 5efac13c53bf / 5

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

<a id="canonical-902d7adf16cfe74a202b91cde87b06703c7568642a4f47afa1c352ba7c1e98c4"></a>

<a id="canonical-4f0d240c3543d6e63c5e23ed219249b81b7eb07fd9f16e30b765a8792fbf792b"></a>

## tenant property — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_ba / 5efac13c53bf / 6

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

<a id="canonical-6b7c75681819f3b6b25b879b301c1849d3d7475980ab59202c148de70f05fa5f"></a>

## Next pages — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_ba / 5efac13c53bf / 7

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-017.md#canonical-06644f1edd1899894148c2f317ecb380c5ef66abd2f2f0786f21206bf906d7b3)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-7ca524a71c2c9ca664ee8ac02047836da5f37907fc3d269e45482e53c73e71fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e6889b8297a7eb13b1cd9c11fb824fc5bca4d92bfd93ad6922cfe68be3006d8"></a>

## enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selecte / b53105dc5c05 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-c92592437695ce4e484e16cf1b000e2558346e492b9be3ef743d389861095739)
- [enable_api_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-017.md#canonical-8a958211739e545afd1a84be8278b9eb87cd211a15e499b08f6fd6775b543169)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-017.md#canonical-06644f1edd1899894148c2f317ecb380c5ef66abd2f2f0786f21206bf906d7b3)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos

<a id="canonical-23fd690302c46a4d1f8af33069ab5898e8955bfb25795f47d6a50d25e52a1fbf"></a>

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

<a id="canonical-44ef53f34a882bcaca57ee1ef69f793fcc4a12c24854c1df5364eafe19cacf0f"></a>

## Direct properties — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selecte / b53105dc5c05 / 3

<a id="canonical-165e6d42136e88d0e609eae81286e69272d14b97b51d473d932c270476b8bd2c"></a>

<a id="canonical-3647bc41b24c8b84b20db04ae34aa1b3992ac57a08069a785be063a0ee84c542"></a>

## api_code_repo property — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selecte / b53105dc5c05 / 4

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

<a id="canonical-24320ae1b56e86fc30d30f278c25c128edd1dd184f57fa659cf2c256f0a0fb1a"></a>

## Next pages — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selecte / b53105dc5c05 / 5

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-017.md#canonical-06644f1edd1899894148c2f317ecb380c5ef66abd2f2f0786f21206bf906d7b3)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-001b95fe7fcb9807071346b2bc36a98014302211e6ed07d4435d791bff84e6e8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6adbd5f3d649471a157314d012aa63d14dce26ade46f2a45f141de58834283e6"></a>

## enable_api_discovery.custom_api_auth_discovery — enable_api_discovery.custom_api_auth_discovery / 49754ea4c296 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-c92592437695ce4e484e16cf1b000e2558346e492b9be3ef743d389861095739)
- enable_api_discovery.custom_api_auth_discovery

<a id="canonical-8ea0bb51fc1afba048d2aae8fb6a11d14110b57c21c3504bd0703232505846ac"></a>

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

<a id="canonical-cd3ca47303e295e1e73a053b7eadac516614d261984e0a3f4ef4488c26a62cc2"></a>

## Direct properties — enable_api_discovery.custom_api_auth_discovery / 49754ea4c296 / 3

- [api_discovery_ref](resources--http_loadbalancer--reference--group-018.md#canonical-39782de31e4cf6fb9497014d40b080ff7ead777a874a73f7a1dc78af63c42d8e): complete subsection reference.

<a id="canonical-79e86f990238cbb9b2aab2dab252512fc71b911e8dce6400c9bd61a5de1a8c54"></a>

## Next pages — enable_api_discovery.custom_api_auth_discovery / 49754ea4c296 / 4

- [enable_api_discovery.custom_api_auth_discovery.api_discovery_ref](resources--http_loadbalancer--reference--group-018.md#canonical-39782de31e4cf6fb9497014d40b080ff7ead777a874a73f7a1dc78af63c42d8e)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-c92592437695ce4e484e16cf1b000e2558346e492b9be3ef743d389861095739)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-39782de31e4cf6fb9497014d40b080ff7ead777a874a73f7a1dc78af63c42d8e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a3028d385d0741e05437c9476ed0cae165252298f1444b9c14cbece9b8e1656"></a>

## enable_api_discovery.custom_api_auth_discovery.api_discovery_ref — enable_api_discovery.custom_api_auth_discovery.api_discovery_ref / d021ccf8c010 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-c92592437695ce4e484e16cf1b000e2558346e492b9be3ef743d389861095739)
- [enable_api_discovery.custom_api_auth_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-001b95fe7fcb9807071346b2bc36a98014302211e6ed07d4435d791bff84e6e8)
- enable_api_discovery.custom_api_auth_discovery.api_discovery_ref

<a id="canonical-65049eaa3656c80c9b6f56b2460430607b94e9aab64384e4d3e1532b61362851"></a>

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

<a id="canonical-4d4ec267d0f384ca4063dcc1c40b1439fbc3b0dce96bfaaf478006ac594660ce"></a>

## Direct properties — enable_api_discovery.custom_api_auth_discovery.api_discovery_ref / d021ccf8c010 / 3

<a id="canonical-03d6e6d2d9672b520be7b762252e154064aa785f66f374309512a1a6ae87ece1"></a>

<a id="canonical-d3e4c546ff243ded71f018bd8ebe6c5133d3d9cd3828addfca86252a26588a15"></a>

## name property — enable_api_discovery.custom_api_auth_discovery.api_discovery_ref / d021ccf8c010 / 4

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

<a id="canonical-8b801d0721d416da088b27dc6cd157ae8d89add23dd1ad8d22ab6e73d86009d4"></a>

<a id="canonical-fe485ed543c0a59f45917395256ff0c232465f78c78ff300cb098d3a4eca569e"></a>

## namespace property — enable_api_discovery.custom_api_auth_discovery.api_discovery_ref / d021ccf8c010 / 5

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

<a id="canonical-bae587851ea6fb65ebe592ea1fdcb60075761788ab33bf9ef5f63b27bcdd6f69"></a>

<a id="canonical-62c59fde74b71c6e5b109bb9e34f651c6e11a8565211e309dad453599d34c203"></a>

## tenant property — enable_api_discovery.custom_api_auth_discovery.api_discovery_ref / d021ccf8c010 / 6

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

<a id="canonical-231f6b3664267ef1a2d6c141cc831c6b590a82809d3ac24f1b141e9c0fd637f4"></a>

## Next pages — enable_api_discovery.custom_api_auth_discovery.api_discovery_ref / d021ccf8c010 / 7

- [enable_api_discovery.custom_api_auth_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-001b95fe7fcb9807071346b2bc36a98014302211e6ed07d4435d791bff84e6e8)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-953ae11281fa4529ec902f19117d94823bf805fb4f08835530b9fa97f9129c29"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a62b5343578988ed2d557a9b345d246c75d91dc4d44ee6e8fa46bb884bde9c4e"></a>

## enable_api_discovery.default_api_auth_discovery — enable_api_discovery.default_api_auth_discovery / dfc330542400 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-c92592437695ce4e484e16cf1b000e2558346e492b9be3ef743d389861095739)
- enable_api_discovery.default_api_auth_discovery

<a id="canonical-01e97eb75af224d0df51b9aca6be350ef3fe15a7945a5a255d90489e2211d039"></a>

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

<a id="canonical-78b2bc16b99f1957ef6a8618882bfe4982bb3b596874d9cca995c82ee73cfdb5"></a>

## Direct properties — enable_api_discovery.default_api_auth_discovery / dfc330542400 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7e9418e0597180f89b56c3cf52b3bea7e3bdaf7637809718a1febca50119f4ed"></a>

## Next pages — enable_api_discovery.default_api_auth_discovery / dfc330542400 / 4

- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-c92592437695ce4e484e16cf1b000e2558346e492b9be3ef743d389861095739)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-f3d58fdbeec479613663bd12a57c81d9c2af56ed18a38ff598298ff0e1fdc477"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c6a21acbb40e3283007e3bc471400968ab009a5dddcd6d719962dfe420144f7"></a>

## enable_api_discovery.disable_learn_from_redirect_traffic — enable_api_discovery.disable_learn_from_redirect_traffic / 749b311ea7b3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-c92592437695ce4e484e16cf1b000e2558346e492b9be3ef743d389861095739)
- enable_api_discovery.disable_learn_from_redirect_traffic

<a id="canonical-1b69a5e5e443f2a5dce43a0e99665dc9989b9274da0cebab109a055e48fadb98"></a>

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

<a id="canonical-05132e8e044a0a9f06a592f82b09f7632f460caa0e20c4da4e9a4f0fa2737bfd"></a>

## Direct properties — enable_api_discovery.disable_learn_from_redirect_traffic / 749b311ea7b3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7aefd307bc896dc0c67c828389e7887900029038ec83bcb1ee0b62be0fef439c"></a>

## Next pages — enable_api_discovery.disable_learn_from_redirect_traffic / 749b311ea7b3 / 4

- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-c92592437695ce4e484e16cf1b000e2558346e492b9be3ef743d389861095739)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e5e6ccc33d87a4887505b909d80117ad702076966af90990dd849b764a7fb68b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb474a509e3aa60eaf030a85da0bc2a02471d722af627d72715af4d1d4c08a7a"></a>

## enable_api_discovery.discovered_api_settings — enable_api_discovery.discovered_api_settings / 6a2898718c9a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-c92592437695ce4e484e16cf1b000e2558346e492b9be3ef743d389861095739)
- enable_api_discovery.discovered_api_settings

<a id="canonical-ec4f81dc5b911d164d54fa604be9c7b9c409d7c8f9c16464655fb7a18f20aa17"></a>

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

<a id="canonical-9c31ecae4f6943c211c0790d5fa0a64edcdd96bdd29c74af6d145c2ddada74b4"></a>

## Direct properties — enable_api_discovery.discovered_api_settings / 6a2898718c9a / 3

<a id="canonical-122505647b4b66f21452b90d7f5dbb8eed08a090e29f5035acd12c59ac109376"></a>

<a id="canonical-96ec3b1288d02ab93520cafe731f90f7488d84f193abd07ada7d39a5da489c47"></a>

## purge_duration_for_inactive_discovered_apis property — enable_api_discovery.discovered_api_settings / 6a2898718c9a / 4

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

<a id="canonical-0d326d7dc6254849d8709d179ff8733385955280c41ab450ed0ba1f2d55b37b7"></a>

## Next pages — enable_api_discovery.discovered_api_settings / 6a2898718c9a / 5

- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-c92592437695ce4e484e16cf1b000e2558346e492b9be3ef743d389861095739)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-3c3fef9df001bf6ec2a41b05448ae5c42a59901077d5b170d4f4d85f3157184c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-93fdb46af70d0d0aea629fb70260b5b6c6681d419bf0f9eb6b8f4221c393c9dd"></a>

## enable_api_discovery.enable_learn_from_redirect_traffic — enable_api_discovery.enable_learn_from_redirect_traffic / 65f70ef66428 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-c92592437695ce4e484e16cf1b000e2558346e492b9be3ef743d389861095739)
- enable_api_discovery.enable_learn_from_redirect_traffic

<a id="canonical-ffbc5f8e89a9d668bb5d3597df5a2ddbcfc2cf6fb1fc069702819a7fc1725bf8"></a>

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

<a id="canonical-c1de0f874e3b496f7f33d5735c4699c1821fb632525719e05c4c4ec8f256e5d9"></a>

## Direct properties — enable_api_discovery.enable_learn_from_redirect_traffic / 65f70ef66428 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ddb976882cecac404cce3c312bfeacb70ffcb1fc0a290f4a0a92625dc7fb5810"></a>

## Next pages — enable_api_discovery.enable_learn_from_redirect_traffic / 65f70ef66428 / 4

- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-c92592437695ce4e484e16cf1b000e2558346e492b9be3ef743d389861095739)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-402b08ef1e68bc5d4cd755fa7d308044f7c196ea49abd373321f605d815aebee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9fd4ec4c60668231cc7a32e917390c0a37ea4bc70fd9f9c30991b96b3016dbea"></a>

## enable_challenge — enable_challenge / 61b25ee01bd2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- enable_challenge

<a id="canonical-3dfa8587960b4f35004d16184a7dd796a97d011b48fea0d273cfa032b9b6832d"></a>

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

<a id="canonical-d9845762a5f64bfb7e2d8ba081e5a63f522b461c0d5d5de8a7f76164d7eaea78"></a>

## Direct properties — enable_challenge / 61b25ee01bd2 / 3

- [captcha_challenge_parameters](resources--http_loadbalancer--reference--group-018.md#canonical-5ad304b2c59293bed6cbfe95704984ff03b91287936ff2f7f42d06a7065bfcde): complete subsection reference.

- [default_captcha_challenge_parameters](resources--http_loadbalancer--reference--group-018.md#canonical-6da20abe77bd5ccb77f4a299abc065d16d4691c50558b8f95efd63c0bc931dca): complete subsection reference.

- [default_js_challenge_parameters](resources--http_loadbalancer--reference--group-018.md#canonical-90812b165a534bbbc2650f3399b1a3c35678cd3d3adb2744812cb363bb30e6e4): complete subsection reference.

- [default_mitigation_settings](resources--http_loadbalancer--reference--group-018.md#canonical-cfb8b16606a5286fc19b611289a8b253b1f1bc72e4f5cf76bbf15ad4e6c2bb1e): complete subsection reference.

- [js_challenge_parameters](resources--http_loadbalancer--reference--group-018.md#canonical-3fafabb4df3ed10d97d4683bae959bcfde18ba5248ff64c8f7ba21cb327de351): complete subsection reference.

- [malicious_user_mitigation](resources--http_loadbalancer--reference--group-018.md#canonical-a030089015ae059767361c63ce0ce2473f1eb995b54cd0828ee1b1bc5523999e): complete subsection reference.

<a id="canonical-37563d719ca160762a34a56da0e1eb0d1696385508b819c1ba283266d7a762e6"></a>

## Next pages — enable_challenge / 61b25ee01bd2 / 4

- [enable_challenge.captcha_challenge_parameters](resources--http_loadbalancer--reference--group-018.md#canonical-5ad304b2c59293bed6cbfe95704984ff03b91287936ff2f7f42d06a7065bfcde)
- [enable_challenge.default_captcha_challenge_parameters](resources--http_loadbalancer--reference--group-018.md#canonical-6da20abe77bd5ccb77f4a299abc065d16d4691c50558b8f95efd63c0bc931dca)
- [enable_challenge.default_js_challenge_parameters](resources--http_loadbalancer--reference--group-018.md#canonical-90812b165a534bbbc2650f3399b1a3c35678cd3d3adb2744812cb363bb30e6e4)
- [enable_challenge.default_mitigation_settings](resources--http_loadbalancer--reference--group-018.md#canonical-cfb8b16606a5286fc19b611289a8b253b1f1bc72e4f5cf76bbf15ad4e6c2bb1e)
- [enable_challenge.js_challenge_parameters](resources--http_loadbalancer--reference--group-018.md#canonical-3fafabb4df3ed10d97d4683bae959bcfde18ba5248ff64c8f7ba21cb327de351)
- [enable_challenge.malicious_user_mitigation](resources--http_loadbalancer--reference--group-018.md#canonical-a030089015ae059767361c63ce0ce2473f1eb995b54cd0828ee1b1bc5523999e)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-5ad304b2c59293bed6cbfe95704984ff03b91287936ff2f7f42d06a7065bfcde"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63b5e28a3430791b3524516667175212c35f4a9adefeb342f3cc5abd5c9e9964"></a>

## enable_challenge.captcha_challenge_parameters — enable_challenge.captcha_challenge_parameters / 3e58635f1d19 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-402b08ef1e68bc5d4cd755fa7d308044f7c196ea49abd373321f605d815aebee)
- enable_challenge.captcha_challenge_parameters

<a id="canonical-a3b5e10a7b4975ee345ad613515f65716e2533dcbc1b2f526a654bdec1c4efd5"></a>

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

<a id="canonical-996b129bf1e28c7782080265e2e72b30693e8e6497651e7a5a93c71d73bf2636"></a>

## Direct properties — enable_challenge.captcha_challenge_parameters / 3e58635f1d19 / 3

<a id="canonical-fed9aca086c7f7acf394661ff4bb642f36aa99d343033b30de916fab945780e7"></a>

<a id="canonical-aa757e597b95d58f521499b8a619f4d30a1fda51d6f2c27db374b44885780515"></a>

## cookie_expiry property — enable_challenge.captcha_challenge_parameters / 3e58635f1d19 / 4

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

<a id="canonical-dc51a9f360539ad26d9a02b149cd48cb848c0b5c7d0e3ff470676a3c3491fcbd"></a>

<a id="canonical-e213fca532b04f7ebddb5fb2680b495cba10f072fec1cc37d2a35e9a963d1f24"></a>

## custom_page property — enable_challenge.captcha_challenge_parameters / 3e58635f1d19 / 5

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

<a id="canonical-4c53b4ab08e4aaf8de300627da93452289abbdbaf21177fb9c1f2a5fcf6f5c2d"></a>

## Next pages — enable_challenge.captcha_challenge_parameters / 3e58635f1d19 / 6

- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-402b08ef1e68bc5d4cd755fa7d308044f7c196ea49abd373321f605d815aebee)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-6da20abe77bd5ccb77f4a299abc065d16d4691c50558b8f95efd63c0bc931dca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5226474e2702434b35b8d2b5cdf2c06c262b1976c07c6f0afe50e345fc69ce30"></a>

## enable_challenge.default_captcha_challenge_parameters — enable_challenge.default_captcha_challenge_parameters / 160f268c40b0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-402b08ef1e68bc5d4cd755fa7d308044f7c196ea49abd373321f605d815aebee)
- enable_challenge.default_captcha_challenge_parameters

<a id="canonical-a8c8f3bd97fad16b76d416ffbee3949b30a1802570089a53b97109c0025a29a4"></a>

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

<a id="canonical-bc46d787fe762923879af27748c5746f31b264d28e66054b10aca27caca000d9"></a>

## Direct properties — enable_challenge.default_captcha_challenge_parameters / 160f268c40b0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e1d1a9e815f9e158c4f1998894d7f217243f302b60aaaa91adcd60735aaf1c42"></a>

## Next pages — enable_challenge.default_captcha_challenge_parameters / 160f268c40b0 / 4

- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-402b08ef1e68bc5d4cd755fa7d308044f7c196ea49abd373321f605d815aebee)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-90812b165a534bbbc2650f3399b1a3c35678cd3d3adb2744812cb363bb30e6e4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d60319d9ddc2abc792e8a4301fd7ba19e47ec4af046f712c1bb98b1ae20e805e"></a>

## enable_challenge.default_js_challenge_parameters — enable_challenge.default_js_challenge_parameters / f866024c7e93 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-402b08ef1e68bc5d4cd755fa7d308044f7c196ea49abd373321f605d815aebee)
- enable_challenge.default_js_challenge_parameters

<a id="canonical-c8c07059eda688a09b026d4f521774d146184e2d7b2def58d3f32abd1b5351c2"></a>

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

<a id="canonical-2fca7f4d1243638e88c0d600ec0bcb52aa725db6443eda342917204727e6e681"></a>

## Direct properties — enable_challenge.default_js_challenge_parameters / f866024c7e93 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2a48406d6a37eda4312001d8e3f5d96083c53eb96756619c1338ef156cc37156"></a>

## Next pages — enable_challenge.default_js_challenge_parameters / f866024c7e93 / 4

- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-402b08ef1e68bc5d4cd755fa7d308044f7c196ea49abd373321f605d815aebee)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-cfb8b16606a5286fc19b611289a8b253b1f1bc72e4f5cf76bbf15ad4e6c2bb1e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-81474f4e5e23238e69e3892b5044e2bd4753760faeada82cbf06064dc53edcc0"></a>

## enable_challenge.default_mitigation_settings — enable_challenge.default_mitigation_settings / e3951ec358ac / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-402b08ef1e68bc5d4cd755fa7d308044f7c196ea49abd373321f605d815aebee)
- enable_challenge.default_mitigation_settings

<a id="canonical-069d67b9f0f5942593b9bd3fad64ef52ecb50cb8aedaf46eb7dbb4567d25c55a"></a>

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

<a id="canonical-f05188d863f4cc25054ff2298ced0f31a410bf8cf569296e7b23733c2b8c2fa0"></a>

## Direct properties — enable_challenge.default_mitigation_settings / e3951ec358ac / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-03642bd57e16fc704af266f9f7113a3fa495de4ee6ee2d24a25b91005fbaa8ce"></a>

## Next pages — enable_challenge.default_mitigation_settings / e3951ec358ac / 4

- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-402b08ef1e68bc5d4cd755fa7d308044f7c196ea49abd373321f605d815aebee)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-3fafabb4df3ed10d97d4683bae959bcfde18ba5248ff64c8f7ba21cb327de351"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-923f9e7c7bc1bb8b98128598ab01de930a9c1c4fe9ab0afe79fba1821145f81e"></a>

## enable_challenge.js_challenge_parameters — enable_challenge.js_challenge_parameters / 9ceaf77d099a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-402b08ef1e68bc5d4cd755fa7d308044f7c196ea49abd373321f605d815aebee)
- enable_challenge.js_challenge_parameters

<a id="canonical-a7b4ec55bfa6e49faa3aa12f3099c6f96f31bbb4d1438b978b883b0820f4680f"></a>

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

<a id="canonical-47cb9a3d4371e57b962bc559b98ec93ca95bb8a96e61f5c88ff09105366b5ce4"></a>

## Direct properties — enable_challenge.js_challenge_parameters / 9ceaf77d099a / 3

<a id="canonical-a74a8c8f4362240ecf377d5401848487e65ae6b598d3ee8ab493be1d80d4a13d"></a>

<a id="canonical-753dc867f53a65744b766d38fc55d942ae455aea516b46cdc75a4fa17d8ea0d7"></a>

## cookie_expiry property — enable_challenge.js_challenge_parameters / 9ceaf77d099a / 4

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

<a id="canonical-281be581025a9186e0b82f90146cedf00af98874ef10314f69432c051b2e2ad1"></a>

<a id="canonical-a6b2636127f544883e496e0a6ad46769655619dc6e565cc4caf75ee31639161e"></a>

## custom_page property — enable_challenge.js_challenge_parameters / 9ceaf77d099a / 5

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

<a id="canonical-3ff3157057ee1c90ff5f9a8acdcf30b1eea63bbb03fc5762bc6dc3ed32dc3847"></a>

<a id="canonical-e997e53348f36a758851b0bb51cc7e243a740286a16433993ad742e07148668c"></a>

## js_script_delay property — enable_challenge.js_challenge_parameters / 9ceaf77d099a / 6

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

<a id="canonical-ee9b1ead3ab2a27d5b2a2f1fce9ad21b3b8f04c82fcd86a2c360ef590530b253"></a>

## Next pages — enable_challenge.js_challenge_parameters / 9ceaf77d099a / 7

- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-402b08ef1e68bc5d4cd755fa7d308044f7c196ea49abd373321f605d815aebee)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-a030089015ae059767361c63ce0ce2473f1eb995b54cd0828ee1b1bc5523999e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-73d35ea389b4e1e56e392093b08dca9481772fb4bc5d2afd582a4f7ca618949e"></a>

## enable_challenge.malicious_user_mitigation — enable_challenge.malicious_user_mitigation / c247e94741b3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-402b08ef1e68bc5d4cd755fa7d308044f7c196ea49abd373321f605d815aebee)
- enable_challenge.malicious_user_mitigation

<a id="canonical-16fd7f18a04a487dd0cd69400edd54f413f678b7c7461bf17b8b1f0363f03935"></a>

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

<a id="canonical-20397dca8e4dcf876d39be149917acb903cb6d51aeec6299da4e4b771a2ffedc"></a>

## Direct properties — enable_challenge.malicious_user_mitigation / c247e94741b3 / 3

<a id="canonical-989b5e8a335926a5a6b5c5ae0801688382d12bbc7e1a0cc0e8b8c8099f03923e"></a>

<a id="canonical-8c5531da7669a69a7d04da54b5be20d5f7866bdd30c55a8463d833ad3a5da9b6"></a>

## name property — enable_challenge.malicious_user_mitigation / c247e94741b3 / 4

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

<a id="canonical-8cd4784b564850edc4acfdf1b6a34df54d30cb9e691fa647b018d78dddd50eb0"></a>

<a id="canonical-b46be0e78b2dd1aa02e2eebf9448fe596b80458ff64807072de9552673403c9d"></a>

## namespace property — enable_challenge.malicious_user_mitigation / c247e94741b3 / 5

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

<a id="canonical-6d5d79feea4bbad0209ea64444c3674787caedb441058242fcedca14688d4a3b"></a>

<a id="canonical-3db61a6fe382b12f205cfbce7d69a11277ef7544491f6196fc417156752d3ee0"></a>

## tenant property — enable_challenge.malicious_user_mitigation / c247e94741b3 / 6

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

<a id="canonical-973364d4a6c546eccacb7ba4141a04ea21301e35e3b61fa5e4e282aaba276618"></a>

## Next pages — enable_challenge.malicious_user_mitigation / c247e94741b3 / 7

- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-402b08ef1e68bc5d4cd755fa7d308044f7c196ea49abd373321f605d815aebee)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-ad9abe471b9e2c134289301619d5e2e343addf7ac9c75b2ef21b2ab519e4a2bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a2eaf444458bb03baddeeba342ad06bf4e6e8bc6a7cd047bbc40d3e71642c596"></a>

## enable_ip_reputation — enable_ip_reputation / deb43a859ed9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- enable_ip_reputation

<a id="canonical-7d35a14fb275919af7cef8f9b7c0ef1bbd579d5a6905ae11cf7f28882c2e175e"></a>

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

<a id="canonical-128376b1da07789d337484cfeb6c5c5513c13b7918e0e80457f8777491cc5280"></a>

## Direct properties — enable_ip_reputation / deb43a859ed9 / 3

<a id="canonical-d92614e721e30e2dffee10fb994bbaf667936b33cf93d988f6b0e7b59fdd2b28"></a>

<a id="canonical-9b99d61c9c7b00f392c3592db6389e5dc1482b337e8b508437a5514670731b06"></a>

## ip_threat_categories property — enable_ip_reputation / deb43a859ed9 / 4

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

<a id="canonical-a9672cdab23eede214d8fff235fdfc5b092f569e3c29c4b118a24ca86ee47842"></a>

## Next pages — enable_ip_reputation / deb43a859ed9 / 5

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-92c71e787d6f8b942b2de2f570db9e12d25adfadafb90be4dab1428b72d246e8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-42b630bc569217d0e4b081a6dd9ca10fa3a468cea974e308243a3da3fbc06a8f"></a>

## enable_malicious_user_detection — enable_malicious_user_detection / 9299248ec4d2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- enable_malicious_user_detection

<a id="canonical-607b3508bdae831af0f1ad3c8f873a93b57362c88115ae2f8fcd196f375107a3"></a>

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

<a id="canonical-4f320f4c778c16687873d610aaff2152df9108c78a6e0726ab1e2b05c6064dfd"></a>

## Direct properties — enable_malicious_user_detection / 9299248ec4d2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f32fe2bd83ded74748ffa2a92fdf4f0f5dbcf4ae9d2a794cf35f9119a0062161"></a>

## Next pages — enable_malicious_user_detection / 9299248ec4d2 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d5e845d5425c1311e52b352f9fa4ea398545c1b66ec415406320486cbd25daef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3ef0416755517c96e5a2213a8155a3323acb85089b76ed1c2a1d2c80a9e9f684"></a>

## enable_threat_mesh — enable_threat_mesh / d9522605743d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- enable_threat_mesh

<a id="canonical-9028820520dd7f20c408a5367458bb5254abd65af3b8c3287c02692c0cce4864"></a>

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

<a id="canonical-57cf92f06f2048d2c0779af643c347b2110ec81969c63529083f71e2e4723895"></a>

## Direct properties — enable_threat_mesh / d9522605743d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6d92e9a34cd8f219b60e321c92f3a3d31e8366a5dddcde876b457a7661eaa7e3"></a>

## Next pages — enable_threat_mesh / d9522605743d / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-8b81ab46828f74c3866c9e65d4068c4c9adac39de02d860dc0b15ed239a98e8f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-36d390b88787cffea897e66bed0880e01df0a05c0896331d9590c19a1231ab0b"></a>

## enable_trust_client_ip_headers — enable_trust_client_ip_headers / b168ee63e095 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- enable_trust_client_ip_headers

<a id="canonical-fd61051873fceb8231a37e8be01c074bfaeea1bfd44309d8f68dff8951e1ac19"></a>

Type: `"object"`. single nested block, Optional.

Trust Client IP Headers List. List of Client IP Headers.

Upstream description:

List of Client IP Headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("client_ip_headers")}
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
enable_trust_client_ip_headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-501f41d1b69ccf44826c4c26c8d77b48dfebca1241b0866b03f828e94d618d05"></a>

## Direct properties — enable_trust_client_ip_headers / b168ee63e095 / 3

<a id="canonical-c8ebe6f8b3b87d77961679c6b68f1661bfa3966d8e3b5d16a47a4e2b157e64fd"></a>

<a id="canonical-241db1d2adc9af61ac57b8e3cc9ddfb9d19ce21567dded86344222bfe00d39e2"></a>

## client_ip_headers property — enable_trust_client_ip_headers / b168ee63e095 / 4

Type: `["list", "string"]`. Optional.

Define the list of one or more Client IP Headers. Headers will be used in order from top to bottom,
meaning if the first header is not present in the request, the system will proceed to check for the
second header, and so on, until one of the listed headers is found. If none of the defined..

Upstream description:

Define the list of one or more Client IP Headers. Headers will be used in order from top to bottom,
meaning if the first header is not present in the request, the system will proceed to check for the
second header, and so on, until one of the listed headers is found. If none of the defined headers
exist, or the value is not an IP address, then the system will use the source IP of the packet. If
multiple defined headers with different names are present in the request, the value of the first
header name in the configuration will be used. If multiple defined headers with the same name are
present in the request, values of all those headers will be combined. The system will read the
right-most IP address from header, if there are multiple IP addresses in the header value. For
X-Forwarded-For header, the system will read the IP address(rightmost - 1), as the client IP.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-07e295b3fdac00b150532c3bbf70c263f9c32a7158ee15889902c92cc31c2f43"></a>

## Next pages — enable_trust_client_ip_headers / b168ee63e095 / 5

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-13765b09fae8733d02c20565cbeec906e4a91070c52e2eac1b36668592698cd6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac065c4704200cdc2184db05fb10dfd4a06e241d7f8bf348d763e2c86bc5d8b2"></a>

## graphql_rules — graphql_rules / 34120e01396e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- graphql_rules

<a id="canonical-4ea960844d00b1868538e9c21d0b114545a17a0e95faa44bb15248de98431b94"></a>

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

<a id="canonical-bd5f433d293a0841fd5848fe14e9345d3ee107a1c1e6e1a0643e51931ddbc82d"></a>

## Direct properties — graphql_rules / 34120e01396e / 3

- [any_domain](resources--http_loadbalancer--reference--group-018.md#canonical-eea5c994285cd42ecca0097e90097a62fcc6c2c958b0bb39a0c1ac116cf70c11): complete subsection reference.

<a id="canonical-af4a8a59ba511288239646ed9e3792d9a21c48326f2ceea03bf4a07bb6cbb258"></a>

<a id="canonical-9fe123cd271934310efa3f53dae7ef3f0c6c719da0459fc3f85dd377ffeb0ae8"></a>

## exact_path property — graphql_rules / 34120e01396e / 4

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

<a id="canonical-179ba1665113c291e98915b862b5a38b2bf4ca084773d1cacf80150a11315e0d"></a>

<a id="canonical-51f2e23c2c67728e94a7cc8545bb72cdb54a0336f53977d2d822582d770f2abd"></a>

## exact_value property — graphql_rules / 34120e01396e / 5

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

- [graphql_settings](resources--http_loadbalancer--reference--group-018.md#canonical-c5d4336126d8b90595d3bbc5d12e3242cc0977fe8023a32d587dd2b6587811cd): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-018.md#canonical-75693b3a3413856eeb17186fa0d5da0f9c792dc0129f8369d1731ff442635620): complete subsection reference.

- [method_get](resources--http_loadbalancer--reference--group-018.md#canonical-5d543071366be612a0512cf1e420bfadcaa911d9e1ff0c08a4e71556808d5a3e): complete subsection reference.

- [method_post](resources--http_loadbalancer--reference--group-018.md#canonical-adc2b44d2f21e5a81f8ceb7fa9e2c942a21fb6e7a8ac1cb7b567e1d6f6e59737): complete subsection reference.

<a id="canonical-58881ba0c9cb25425276f34d61fa28efd197b47321892affbf56bdf3d3a73aa3"></a>

<a id="canonical-4cf3d9be5e53167daacd1447aabd7999d6481745063fc5442484b6be8bb42c80"></a>

## suffix_value property — graphql_rules / 34120e01396e / 6

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

<a id="canonical-9fa77d99cbf0a2994f9f2877c36bca74f7b7921a6990c55b0273edbad0f6d43a"></a>

## Next pages — graphql_rules / 34120e01396e / 7

- [graphql_rules.any_domain](resources--http_loadbalancer--reference--group-018.md#canonical-eea5c994285cd42ecca0097e90097a62fcc6c2c958b0bb39a0c1ac116cf70c11)
- [graphql_rules.graphql_settings](resources--http_loadbalancer--reference--group-018.md#canonical-c5d4336126d8b90595d3bbc5d12e3242cc0977fe8023a32d587dd2b6587811cd)
- [graphql_rules.metadata](resources--http_loadbalancer--reference--group-018.md#canonical-75693b3a3413856eeb17186fa0d5da0f9c792dc0129f8369d1731ff442635620)
- [graphql_rules.method_get](resources--http_loadbalancer--reference--group-018.md#canonical-5d543071366be612a0512cf1e420bfadcaa911d9e1ff0c08a4e71556808d5a3e)
- [graphql_rules.method_post](resources--http_loadbalancer--reference--group-018.md#canonical-adc2b44d2f21e5a81f8ceb7fa9e2c942a21fb6e7a8ac1cb7b567e1d6f6e59737)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-eea5c994285cd42ecca0097e90097a62fcc6c2c958b0bb39a0c1ac116cf70c11"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-20aad8b5e34a01ca06bd56906b95deca6fb43db18f11bbaf59445b586527e61b"></a>

## graphql_rules.any_domain — graphql_rules.any_domain / 80ab0b6c69c9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-13765b09fae8733d02c20565cbeec906e4a91070c52e2eac1b36668592698cd6)
- graphql_rules.any_domain

<a id="canonical-70af566f06a05a77d6486a58f843ffc21b3d3b4a43b4a939d338c57958791661"></a>

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

<a id="canonical-60a9719e098298b03e0ce7101c06fb54bc5d94e445e0c2113944cfbbd4dee632"></a>

## Direct properties — graphql_rules.any_domain / 80ab0b6c69c9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-66cd58a0cd7f8fcd16bb27caa4c86e30657d33897bb3c7e2756818339e39cc25"></a>

## Next pages — graphql_rules.any_domain / 80ab0b6c69c9 / 4

- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-13765b09fae8733d02c20565cbeec906e4a91070c52e2eac1b36668592698cd6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c5d4336126d8b90595d3bbc5d12e3242cc0977fe8023a32d587dd2b6587811cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-49bf92f4ef7272eeff0aa4157ae841bb7321dd06de45f18d37621c9942e871e8"></a>

## graphql_rules.graphql_settings — graphql_rules.graphql_settings / 4edc9c088c45 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-13765b09fae8733d02c20565cbeec906e4a91070c52e2eac1b36668592698cd6)
- graphql_rules.graphql_settings

<a id="canonical-a4e4a638a9d3cecbe02f46e5c26e42b3cd8e67b59ccb5f6305b6f1f6401df4b9"></a>

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

<a id="canonical-f55bafb804ba7cc6f0ab708c0df97152f99ce3de505cc10a5583cbcb71236d77"></a>

## Direct properties — graphql_rules.graphql_settings / 4edc9c088c45 / 3

- [disable_introspection](resources--http_loadbalancer--reference--group-018.md#canonical-a6565be9933802ca3ed8112aa4d243d953207c61b81b9a4c23cdd68a797dbfbb): complete subsection reference.

- [enable_introspection](resources--http_loadbalancer--reference--group-018.md#canonical-1b3d234774062ba92a9bdf24df05437d205d0ecd626f6d497158ce27f2c03509): complete subsection reference.

<a id="canonical-b6c293f13ed48bfa49db807dc676c48ac0dd2317b97d9ba951bcf126c9c2b9d5"></a>

<a id="canonical-55a5e2d5c1920cb148aa4d6481146df0a6a805456f9599b0833a790e1722b304"></a>

## max_batched_queries property — graphql_rules.graphql_settings / 4edc9c088c45 / 4

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

<a id="canonical-e29e4f99b64f3bf3cf34f3a6ee98bcb3446185e0cab7d4378d2edceef006c7b0"></a>

<a id="canonical-82bd4137303966766eca8c053683b99a8de2b312745f25f88a727cbcb62c0e60"></a>

## max_depth property — graphql_rules.graphql_settings / 4edc9c088c45 / 5

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

<a id="canonical-9a97b6e28bb4887f14d9282bd92e29ffdbee76c3f5a6c56e9d341839345adb2f"></a>

<a id="canonical-77a10f30a30bba272977f11b687db78371b658b04bfddf0e005ff20b678faead"></a>

## max_total_length property — graphql_rules.graphql_settings / 4edc9c088c45 / 6

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

<a id="canonical-19440abb3b4f386e9c23fcabae5ba54b31aefcdaebd6a96f79d456d933ab18b6"></a>

## Next pages — graphql_rules.graphql_settings / 4edc9c088c45 / 7

- [graphql_rules.graphql_settings.disable_introspection](resources--http_loadbalancer--reference--group-018.md#canonical-a6565be9933802ca3ed8112aa4d243d953207c61b81b9a4c23cdd68a797dbfbb)
- [graphql_rules.graphql_settings.enable_introspection](resources--http_loadbalancer--reference--group-018.md#canonical-1b3d234774062ba92a9bdf24df05437d205d0ecd626f6d497158ce27f2c03509)
- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-13765b09fae8733d02c20565cbeec906e4a91070c52e2eac1b36668592698cd6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-a6565be9933802ca3ed8112aa4d243d953207c61b81b9a4c23cdd68a797dbfbb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-69e689f505fca13af5d7e16d172569ae5182f23a8e94f82db3092caa17529f36"></a>

## graphql_rules.graphql_settings.disable_introspection — graphql_rules.graphql_settings.disable_introspection / 8498da4c72d5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-13765b09fae8733d02c20565cbeec906e4a91070c52e2eac1b36668592698cd6)
- [graphql_rules.graphql_settings](resources--http_loadbalancer--reference--group-018.md#canonical-c5d4336126d8b90595d3bbc5d12e3242cc0977fe8023a32d587dd2b6587811cd)
- graphql_rules.graphql_settings.disable_introspection

<a id="canonical-b6630ca619884e88996727dc5314a5b66cd60fa4f9911dc6625f3c51b4ca2d3e"></a>

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

<a id="canonical-0b045f57304a29dd26f7f58be4c316b86aa146a6dd4d88a441ae5d38509b7fed"></a>

## Direct properties — graphql_rules.graphql_settings.disable_introspection / 8498da4c72d5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-952785baa98def19b31822e8a7c065d3ef577c1d5037e0cc891481e95066ec33"></a>

## Next pages — graphql_rules.graphql_settings.disable_introspection / 8498da4c72d5 / 4

- [graphql_rules.graphql_settings](resources--http_loadbalancer--reference--group-018.md#canonical-c5d4336126d8b90595d3bbc5d12e3242cc0977fe8023a32d587dd2b6587811cd)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-1b3d234774062ba92a9bdf24df05437d205d0ecd626f6d497158ce27f2c03509"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-651b9bb31bf0189960899b6a2421e7601adede0207781b0085a6e2a929f7d334"></a>

## graphql_rules.graphql_settings.enable_introspection — graphql_rules.graphql_settings.enable_introspection / 383a10c3d08e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-13765b09fae8733d02c20565cbeec906e4a91070c52e2eac1b36668592698cd6)
- [graphql_rules.graphql_settings](resources--http_loadbalancer--reference--group-018.md#canonical-c5d4336126d8b90595d3bbc5d12e3242cc0977fe8023a32d587dd2b6587811cd)
- graphql_rules.graphql_settings.enable_introspection

<a id="canonical-99c41a753bb03f7f6e7635e646df22d92f7698ea926c6a1ba18b0e826d714e88"></a>

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

<a id="canonical-6bed4558790009a7fdffc6a718ebab11fc763bcfc9be55d6c2d2d18ec97ab92a"></a>

## Direct properties — graphql_rules.graphql_settings.enable_introspection / 383a10c3d08e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5e7347a953d502470424c8b28b2457c0f1f30f489d52c873b20bca2ffaece7cd"></a>

## Next pages — graphql_rules.graphql_settings.enable_introspection / 383a10c3d08e / 4

- [graphql_rules.graphql_settings](resources--http_loadbalancer--reference--group-018.md#canonical-c5d4336126d8b90595d3bbc5d12e3242cc0977fe8023a32d587dd2b6587811cd)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-75693b3a3413856eeb17186fa0d5da0f9c792dc0129f8369d1731ff442635620"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-632d592563bbad86f1fc5397a40b056656266398e974e7defc38009e417905c7"></a>

## graphql_rules.metadata — graphql_rules.metadata / fff07f9b331c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-13765b09fae8733d02c20565cbeec906e4a91070c52e2eac1b36668592698cd6)
- graphql_rules.metadata

<a id="canonical-76f21db70715de84a3f426875af165808a72832195c9b9d781397aab2f4d3e01"></a>

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

<a id="canonical-b2eefc0efff04e92ba51d66302ee8ebcc87b52976783bd7420d9f62fa6302744"></a>

## Direct properties — graphql_rules.metadata / fff07f9b331c / 3

<a id="canonical-b42bf594e71851c9ae93a3e70960575d4ecdaa8c77f53b2e85426bd106e3d57a"></a>

<a id="canonical-1b655035ab7b98c5295dd35b95498e54a11bada064de1068fad9379a72b3a2d7"></a>

## description_spec property — graphql_rules.metadata / fff07f9b331c / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-3493ca1e4e66404db807900302b68146fb269f847b966cbc93b99a715bf73f5b"></a>

<a id="canonical-bb88853141ea56384183a550d9e3863435c33dabeb9f62cd5d7be7d8f1041765"></a>

## name property — graphql_rules.metadata / fff07f9b331c / 5

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

<a id="canonical-24d067f33f02c6fa837504379bea2919dc55bfbac639ebba1e2ce39624dd888c"></a>

## Next pages — graphql_rules.metadata / fff07f9b331c / 6

- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-13765b09fae8733d02c20565cbeec906e4a91070c52e2eac1b36668592698cd6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-5d543071366be612a0512cf1e420bfadcaa911d9e1ff0c08a4e71556808d5a3e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a6a5ffa2d45e17819101dad10b86ad5dd1a6191e34d2fadd19194b3701dd73e0"></a>

## graphql_rules.method_get — graphql_rules.method_get / 35aa21d97cf4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-13765b09fae8733d02c20565cbeec906e4a91070c52e2eac1b36668592698cd6)
- graphql_rules.method_get

<a id="canonical-44b95199557724f4bda614a0e8f5c149f34496d39d07ee917f8b3632fae73123"></a>

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

<a id="canonical-9d98958851fa9f217fa5a089681bda7bc4f36be64c75a012ae8aa57e23973973"></a>

## Direct properties — graphql_rules.method_get / 35aa21d97cf4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-06210682b887086f2ec25fb03a77c2570bd35a398762927cfdfb9b463837f17c"></a>

## Next pages — graphql_rules.method_get / 35aa21d97cf4 / 4

- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-13765b09fae8733d02c20565cbeec906e4a91070c52e2eac1b36668592698cd6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-adc2b44d2f21e5a81f8ceb7fa9e2c942a21fb6e7a8ac1cb7b567e1d6f6e59737"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41a538c3b4cdf31f3593d47fd0ec5c6c64949153c11c155761039d1482d848ad"></a>

## graphql_rules.method_post — graphql_rules.method_post / e1430dc40105 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-13765b09fae8733d02c20565cbeec906e4a91070c52e2eac1b36668592698cd6)
- graphql_rules.method_post

<a id="canonical-7db113bb7579e3eb44f5fc855421c25f237c0b4e0cf3aaedb2ddeb667b483b39"></a>

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

<a id="canonical-b82ebf2dad35c57492fb09f8fb46289a4f4cf0ef13f4e5795497cbbf9a553e9e"></a>

## Direct properties — graphql_rules.method_post / e1430dc40105 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-015652096a525331b47d8ff837003e126bc6cc5087c804c9c155f9dc6c158b5f"></a>

## Next pages — graphql_rules.method_post / e1430dc40105 / 4

- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-13765b09fae8733d02c20565cbeec906e4a91070c52e2eac1b36668592698cd6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e1f14d7c9b8e8fb184d02bb2a0fe678d962de127c1841a842aff87500721227b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7da6a9c45522e5e3dfaa304c3fee07d4f9f053ca512a9e7c83f4a2a124ced1d1"></a>

## http — http / f96612ca7ef0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- http

<a id="canonical-6214288c716eba64297c1420d7ed3db856962ec79283222845ca61029db87952"></a>

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

- [http](resources--http_loadbalancer--reference--group-018.md#canonical-6214288c716eba64297c1420d7ed3db856962ec79283222845ca61029db87952)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-7096d1d341cdca7b00c42509a47e45447b434b9e1157c5b7d0537010584d4c27)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-6b6f058e2a7c8a41e1f4634b1f217cf058ab1bd258b9943ea8c4d3640bf10090)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
http {
  # Configure direct properties listed below.
}
```

<a id="canonical-770695e9d2391eabfc7c09f5250c375b001fb0de5c2824e27ea2f01a92e286a2"></a>

## Direct properties — http / f96612ca7ef0 / 3

<a id="canonical-9d692e4abcde514df33ec1999f5c522878ccf42936812dd608861b7fa195df2a"></a>

<a id="canonical-6a7e6428828db2dc0ac4241e7b01c51f7e58220f7f719106ed37d04d4eb4623a"></a>

## dns_volterra_managed property — http / f96612ca7ef0 / 4

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

<a id="canonical-181523ac433c1962bcda8fc71c04eb5d822d759815215f7080b56f2501480a65"></a>

<a id="canonical-ca33180b278411f0fd4e9dbccf5805ce09f5d50378f1e220279a6d02a2055dcb"></a>

## port property — http / f96612ca7ef0 / 5

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

<a id="canonical-449b59aceb3b0d2165043a47ef7a9a1d827f34ce03e7556b92998dad72eee8f5"></a>

<a id="canonical-eb0679dd8bb02fa45c7e1ba6ff2fafa222bb9a081ac463e425b41d9abf06df81"></a>

## port_ranges property — http / f96612ca7ef0 / 6

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

<a id="canonical-3e057c8c439d764af41ebff5344e1f44eceb90e8e44ca079f989a9849a555637"></a>

## Next pages — http / f96612ca7ef0 / 7

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1144921b5a4ffff244ddc90450cce2812c7dd1d6aa28f39c157f43db9537e8a"></a>

## https — https / 759b75fef7f0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- https

<a id="canonical-7096d1d341cdca7b00c42509a47e45447b434b9e1157c5b7d0537010584d4c27"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting HTTP proxy with bring your own certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("append_server_name",
    "default_header"),
  validators.ConflictingObjectAttributes("append_server_name",
    "pass_through"),
  validators.ConflictingObjectAttributes("append_server_name",
    "server_name"),
  validators.ConflictingObjectAttributes("default_header",
    "pass_through"),
  validators.ConflictingObjectAttributes("default_header",
    "server_name"),
  validators.ConflictingObjectAttributes("default_loadbalancer",
    "non_default_loadbalancer"),
  validators.ConflictingObjectAttributes("disable_path_normalize",
    "enable_path_normalize"),
  validators.ConflictingObjectAttributes("pass_through",
    "server_name"),
  validators.ConflictingObjectAttributes("port",
    "port_ranges"),
  validators.ConflictingObjectAttributes("tls_cert_params",
    "tls_parameters")}
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
  "x-ves-oneof-field-default_lb_choice": "[\"default_loadbalancer\",\"non_default_loadbalancer\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]",
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_parameters\"]"
}
```

Terraform syntax:

```terraform
https {
  # Configure direct properties listed below.
}
```

<a id="canonical-b8814a428690d26c5b9cc2346191d32cf06300c18b429ee8b2f58222885e2b80"></a>

## Direct properties — https / 759b75fef7f0 / 3

<a id="canonical-9059a53984cd3a23eec568fdc2a299a4655eb7d42ab202ad10e4f4fed0ce1579"></a>

<a id="canonical-644088e8911f4f542afeff005259860e6c0aef37aa6df822425cd23608411a24"></a>

## add_hsts property — https / 759b75fef7f0 / 4

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

<a id="canonical-2b1448c262debc05910ccd43eebfe11c88c5b6ed8a5f66ee12da8c9fc0f4f0ea"></a>

<a id="canonical-74f3fa713ea05a508d7d9046f600b506cf9451ccad1b6ef6c3e0112424fc238e"></a>

## append_server_name property — https / 759b75fef7f0 / 5

Type: `"string"`. Optional.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Upstream description:

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [coalescing_options](resources--http_loadbalancer--reference--group-018.md#canonical-42dd2ed8dbc4f300d34415dedf8909f67a41e6b53b64cbc61663ead8f53242dc): complete subsection reference.

<a id="canonical-fc7d8412e26053e05f3f8f51115c2740f960b142c38f8c7b037af6a976c04292"></a>

<a id="canonical-70c24c2d5bdb746127cbcd0ab6f385253c14fffd916e8659bf05ce08bee1c8d2"></a>

## connection_idle_timeout property — https / 759b75fef7f0 / 6

Type: `"number"`. Optional.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed.

Upstream description:

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(600000),
}
```

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

- [default_header](resources--http_loadbalancer--reference--group-018.md#canonical-01e920d7e561c7f929c6c5a82db3b3cf486a64182d8590c6aba092453607bd97): complete subsection reference.

- [default_loadbalancer](resources--http_loadbalancer--reference--group-018.md#canonical-62e72e0786e677bc48e9ffcb8f67e8b4a10c486f3c01a0c49b5d85c5fe29d5f4): complete subsection reference.

- [disable_path_normalize](resources--http_loadbalancer--reference--group-018.md#canonical-cdb727cfa28aec9a454752ce6a9149ab23a3099d6f5a212a391cd2f9efcd19b2): complete subsection reference.

- [enable_path_normalize](resources--http_loadbalancer--reference--group-018.md#canonical-f4d3c03806ffecfc0716ea39d1a085e84d3e78bf8a5478b47982f9e05efc853d): complete subsection reference.

- [http_protocol_options](resources--http_loadbalancer--reference--group-018.md#canonical-6dda791fdf251be6f1dd8108d706d8252bc863f7c0c5a980dde0ec7865caf32a): complete subsection reference.

<a id="canonical-5a9d2372c06f84c2a8e4c6a9f033e1a6a0d4420463dd9fb5372983e18a734af7"></a>

<a id="canonical-038f2834d67650bd8fcfb7ba5e67b73f90184abfc4b997e866a40457900af3fc"></a>

## http_redirect property — https / 759b75fef7f0 / 7

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

- [non_default_loadbalancer](resources--http_loadbalancer--reference--group-018.md#canonical-da65490816f7086893b9e68a91c94959c7bbf195f4f220138163abc38c176b33): complete subsection reference.

- [pass_through](resources--http_loadbalancer--reference--group-018.md#canonical-7827b7da9d7a3b052a4b08639f9759e03c95eb68f26d8562e112c4de19300d11): complete subsection reference.

<a id="canonical-723ce587faf33896f7c30d31c3db3e07f3370d64a00030d372c1894ba01606e6"></a>

<a id="canonical-a438623a8fc42b71a380fd6f79dbd109a865f60cad62eae60562c68d643aff83"></a>

## port property — https / 759b75fef7f0 / 8

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTPS port to Listen.

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

<a id="canonical-ed59f2a2f03c484f8890483362212a0af3dd4b852741cd319d4ff8ef62868024"></a>

<a id="canonical-7a9d9f7e716894710a40abe3807059ca61477bf23ab47fc7d6dd5ca2137ab641"></a>

## port_ranges property — https / 759b75fef7f0 / 9

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

<a id="canonical-e20f37630cc35de587ce99dd61446d5da474b8bc7e6f32f6dfc8fb9d8bf66478"></a>

<a id="canonical-20dd3bb19df549bac3afead61f7d3a8ccecdddd5221702708cca4070e8c3d22a"></a>

## server_name property — https / 759b75fef7f0 / 10

Type: `"string"`. Optional.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [tls_cert_params](resources--http_loadbalancer--reference--group-018.md#canonical-31ca63fa137c47bc6b3bc3049c6a66381c9558e74ba34ad4e12c7ace5ed5a91f): complete subsection reference.

- [tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-6d5c7ce1efa80c26a575c1f2ba94312938bd3c5c92b95992b99bf5d0dc8f22a6): complete subsection reference.

<a id="canonical-f9b39991cbff6c692458a4784b1f0fa6a21efb151e5b7ab0d9bd12b6137cb33f"></a>

## Next pages — https / 759b75fef7f0 / 11

- [https.coalescing_options](resources--http_loadbalancer--reference--group-018.md#canonical-42dd2ed8dbc4f300d34415dedf8909f67a41e6b53b64cbc61663ead8f53242dc)
- [https.default_header](resources--http_loadbalancer--reference--group-018.md#canonical-01e920d7e561c7f929c6c5a82db3b3cf486a64182d8590c6aba092453607bd97)
- [https.default_loadbalancer](resources--http_loadbalancer--reference--group-018.md#canonical-62e72e0786e677bc48e9ffcb8f67e8b4a10c486f3c01a0c49b5d85c5fe29d5f4)
- [https.disable_path_normalize](resources--http_loadbalancer--reference--group-018.md#canonical-cdb727cfa28aec9a454752ce6a9149ab23a3099d6f5a212a391cd2f9efcd19b2)
- [https.enable_path_normalize](resources--http_loadbalancer--reference--group-018.md#canonical-f4d3c03806ffecfc0716ea39d1a085e84d3e78bf8a5478b47982f9e05efc853d)
- [https.http_protocol_options](resources--http_loadbalancer--reference--group-018.md#canonical-6dda791fdf251be6f1dd8108d706d8252bc863f7c0c5a980dde0ec7865caf32a)
- [https.non_default_loadbalancer](resources--http_loadbalancer--reference--group-018.md#canonical-da65490816f7086893b9e68a91c94959c7bbf195f4f220138163abc38c176b33)
- [https.pass_through](resources--http_loadbalancer--reference--group-018.md#canonical-7827b7da9d7a3b052a4b08639f9759e03c95eb68f26d8562e112c4de19300d11)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-018.md#canonical-31ca63fa137c47bc6b3bc3049c6a66381c9558e74ba34ad4e12c7ace5ed5a91f)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-6d5c7ce1efa80c26a575c1f2ba94312938bd3c5c92b95992b99bf5d0dc8f22a6)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-42dd2ed8dbc4f300d34415dedf8909f67a41e6b53b64cbc61663ead8f53242dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-337345b431dbeb68eebd7527a41a4bfb042affbdd5bd65fb6442342aa5dad731"></a>

## https.coalescing_options — https.coalescing_options / fdc422d91885 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- https.coalescing_options

<a id="canonical-5a3c88b978f5fdaa8ad742dbaf510bce8939247535ae03bb8dc4b3c21cd95f3f"></a>

Type: `"object"`. single nested block, Optional.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_coalescing",
    "strict_coalescing")}
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
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

Terraform syntax:

```terraform
coalescing_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-0f79f47c9dd42c00226707be8305783db86deda0a17ada051db6a7e14abb8eba"></a>

## Direct properties — https.coalescing_options / fdc422d91885 / 3

- [default_coalescing](resources--http_loadbalancer--reference--group-018.md#canonical-f868b33073cf11fff0a35fd635321a37e68c698fd4f79b90d2b93a5b516b8b00): complete subsection reference.

- [strict_coalescing](resources--http_loadbalancer--reference--group-018.md#canonical-b36bd7a437bad28d38e4da8afa5b86d12a0b9693a2b1e81ea1c3cc18c00cb8a2): complete subsection reference.

<a id="canonical-ec6855a00b82784de850bb46576d861a35ede616d78d290b90f0c50a892a387d"></a>

## Next pages — https.coalescing_options / fdc422d91885 / 4

- [https.coalescing_options.default_coalescing](resources--http_loadbalancer--reference--group-018.md#canonical-f868b33073cf11fff0a35fd635321a37e68c698fd4f79b90d2b93a5b516b8b00)
- [https.coalescing_options.strict_coalescing](resources--http_loadbalancer--reference--group-018.md#canonical-b36bd7a437bad28d38e4da8afa5b86d12a0b9693a2b1e81ea1c3cc18c00cb8a2)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-f868b33073cf11fff0a35fd635321a37e68c698fd4f79b90d2b93a5b516b8b00"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4610d6ce1d2e74a7c01651ad414f3a7da32cb62174c29c611cd111fbd79c1a8"></a>

## https.coalescing_options.default_coalescing — https.coalescing_options.default_coalescing / 550af1e11148 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.coalescing_options](resources--http_loadbalancer--reference--group-018.md#canonical-42dd2ed8dbc4f300d34415dedf8909f67a41e6b53b64cbc61663ead8f53242dc)
- https.coalescing_options.default_coalescing

<a id="canonical-4c3b88b172264146f92f70a14412e18e6760be670e6657aca76e30799dbe4825"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default coalescing.

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
default_coalescing = {}
```

<a id="canonical-9f04dbdaa4b6808e1202a6c2b8c91514b1ac7af25fe580d9385f9fef423521db"></a>

## Direct properties — https.coalescing_options.default_coalescing / 550af1e11148 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-af23b2ce7a46f7f43848804f9b67cfbf212587d6da1f3464927a51dcda4ee75c"></a>

## Next pages — https.coalescing_options.default_coalescing / 550af1e11148 / 4

- [https.coalescing_options](resources--http_loadbalancer--reference--group-018.md#canonical-42dd2ed8dbc4f300d34415dedf8909f67a41e6b53b64cbc61663ead8f53242dc)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b36bd7a437bad28d38e4da8afa5b86d12a0b9693a2b1e81ea1c3cc18c00cb8a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-55b0e86c7f8c25e0af3e06e256cf982f5b4416b878a609b510b2e54357ccd6fd"></a>

## https.coalescing_options.strict_coalescing — https.coalescing_options.strict_coalescing / 4360b317e4f3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.coalescing_options](resources--http_loadbalancer--reference--group-018.md#canonical-42dd2ed8dbc4f300d34415dedf8909f67a41e6b53b64cbc61663ead8f53242dc)
- https.coalescing_options.strict_coalescing

<a id="canonical-2e73c1740c5f0ca365c9f4c020e83f92c093ae682df67bbf5cb1c19009f9359c"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for strict coalescing.

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
strict_coalescing = {}
```

<a id="canonical-1595f19f07b9a103b1eb0144bb0c6eacfec1bff5ea42ec7018a7483c1ce8f871"></a>

## Direct properties — https.coalescing_options.strict_coalescing / 4360b317e4f3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b9536793621b3c80edb37710cd1b059ccd0050bb8b14b77f1c8b0c92707f0290"></a>

## Next pages — https.coalescing_options.strict_coalescing / 4360b317e4f3 / 4

- [https.coalescing_options](resources--http_loadbalancer--reference--group-018.md#canonical-42dd2ed8dbc4f300d34415dedf8909f67a41e6b53b64cbc61663ead8f53242dc)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-01e920d7e561c7f929c6c5a82db3b3cf486a64182d8590c6aba092453607bd97"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f29c75640b2f4ab9c3e31f17429f7d1640bae66d8e9a924980b8bb44c980c8df"></a>

## https.default_header — https.default_header / d0592b043bd0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- https.default_header

<a id="canonical-ecedcdad12e64efa652899b8fe07e06988065799306e91764b71d8d5a2308fd0"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default header.

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
default_header = {}
```

<a id="canonical-bf903d3c25b660399d0257404e03e3ef89202205cf003d6b167c3f6c654eaf3e"></a>

## Direct properties — https.default_header / d0592b043bd0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ef9fd91ff3840a5f53e8d348a8cfa037b047c3559de237414b4fa8bad0e6bd49"></a>

## Next pages — https.default_header / d0592b043bd0 / 4

- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-62e72e0786e677bc48e9ffcb8f67e8b4a10c486f3c01a0c49b5d85c5fe29d5f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-37026ae91454d9f0879e5b3cc5fb2bcc9f53c2b2f596c87253726ce73b0f11f9"></a>

## https.default_loadbalancer — https.default_loadbalancer / 51d1e51479b9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- https.default_loadbalancer

<a id="canonical-6b8be6befdb9bdc4722d98543130d2e7439f4280aa47f4526ddda73bb9e0ef28"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default loadbalancer.

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
default_loadbalancer = {}
```

<a id="canonical-87a0ff8d34094f3ff0f63cb9ce543ef2bb7606352c66ff5219e7e344d789b5ca"></a>

## Direct properties — https.default_loadbalancer / 51d1e51479b9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6e7feb4742d5d8220b52fc4c41b45ac5079f0e61746e79483c4fa6a4ab1ec289"></a>

## Next pages — https.default_loadbalancer / 51d1e51479b9 / 4

- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-cdb727cfa28aec9a454752ce6a9149ab23a3099d6f5a212a391cd2f9efcd19b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f91938b2c2958823d228b1276f663944092ada512b16400c484233858ebb67c3"></a>

## https.disable_path_normalize — https.disable_path_normalize / 738fe526d72c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- https.disable_path_normalize

<a id="canonical-4db19cb955a64c7980f1ec40eb63d3396a061f5f8524c62c70daeb2c4a42be0b"></a>

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
disable_path_normalize = {}
```

<a id="canonical-78089b22f83e7798840039060076f744572455ee24574f0e2e43c25e0c354986"></a>

## Direct properties — https.disable_path_normalize / 738fe526d72c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7a0167c458faeebcf047d41b81db994ca9df9b6addb11bfe0b3e9af8d09298b1"></a>

## Next pages — https.disable_path_normalize / 738fe526d72c / 4

- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-f4d3c03806ffecfc0716ea39d1a085e84d3e78bf8a5478b47982f9e05efc853d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a2dd6be861a9b3a6069b9e464544c41d430a3a048bef829fdbdd9200e0226fd9"></a>

## https.enable_path_normalize — https.enable_path_normalize / a4d3f4451e15 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- https.enable_path_normalize

<a id="canonical-d068ce4f450aa33e929c984b4fb8a95862b4cf13448204faa78693bdf4722c8d"></a>

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
enable_path_normalize = {}
```

<a id="canonical-d11a0cdb885f6de14d194b3ad9accedd17c96331d82e5f67da774878d6f9fd81"></a>

## Direct properties — https.enable_path_normalize / a4d3f4451e15 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-052946101add552aca6c94e3073d5f4734eba3558b82406bd2abcf74243f1ae0"></a>

## Next pages — https.enable_path_normalize / a4d3f4451e15 / 4

- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-6dda791fdf251be6f1dd8108d706d8252bc863f7c0c5a980dde0ec7865caf32a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b5e8270eef0e4186684cd6f032b0c86d9fe6c0c1dc8be330ef2067ac0e211f9"></a>

## https.http_protocol_options — https.http_protocol_options / 07be0adf7ef4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- https.http_protocol_options

<a id="canonical-28dcab14cc35088824f7e14dd4e06337cf77721f17f22f788b7760ac6189df6e"></a>

Type: `"object"`. single nested block, Optional.

HTTP protocol configuration OPTIONS for downstream connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v1_v2"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v2_only"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_v2",
    "http_protocol_enable_v2_only")}
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
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

Terraform syntax:

```terraform
http_protocol_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-2e71dc45743b384405eaad4ff67dc8750770d2bb2ebd15dcef0d56cdb980f553"></a>

## Direct properties — https.http_protocol_options / 07be0adf7ef4 / 3

- [http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-018.md#canonical-9f443e73d19e56a4af78922ae1db921012ef20f3a711092726357b8d29c14b33): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--http_loadbalancer--reference--group-018.md#canonical-04dea4678ca8b618cf550fb6c09d3971429d1f8846fa624fbf8230df42b3f455): complete subsection reference.

- [http_protocol_enable_v2_only](resources--http_loadbalancer--reference--group-018.md#canonical-27e0020f98056e1bbcf841d1d1bb8ae42be8d924d6352cbd3a6590fc40e34b92): complete subsection reference.

<a id="canonical-78248e53b482dbf00c868f5049202ab9cdbbe1a11d769691c6013a7ab323943c"></a>

## Next pages — https.http_protocol_options / 07be0adf7ef4 / 4

- [https.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-018.md#canonical-9f443e73d19e56a4af78922ae1db921012ef20f3a711092726357b8d29c14b33)
- [https.http_protocol_options.http_protocol_enable_v1_v2](resources--http_loadbalancer--reference--group-018.md#canonical-04dea4678ca8b618cf550fb6c09d3971429d1f8846fa624fbf8230df42b3f455)
- [https.http_protocol_options.http_protocol_enable_v2_only](resources--http_loadbalancer--reference--group-018.md#canonical-27e0020f98056e1bbcf841d1d1bb8ae42be8d924d6352cbd3a6590fc40e34b92)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-9f443e73d19e56a4af78922ae1db921012ef20f3a711092726357b8d29c14b33"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ed06904346cad9042ace410273013eca5d68c18df1caf7c853b1fb58227a18d0"></a>

## https.http_protocol_options.http_protocol_enable_v1_only — https.http_protocol_options.http_protocol_enable_v1_only / d0990432f21f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.http_protocol_options](resources--http_loadbalancer--reference--group-018.md#canonical-6dda791fdf251be6f1dd8108d706d8252bc863f7c0c5a980dde0ec7865caf32a)
- https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-49693d69129b260acf5c2a005d537fcee0869a7fcf0309bc0b8dc899991a0ab8"></a>

Type: `"object"`. single nested block, Optional.

HTTP/1.1 Protocol OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
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
http_protocol_enable_v1_only {
  # Configure direct properties listed below.
}
```

<a id="canonical-055557911df4b700d487184928589422580df7facb144e6a1902d0417fe60b6f"></a>

## Direct properties — https.http_protocol_options.http_protocol_enable_v1_only / d0990432f21f / 3

- [header_transformation](resources--http_loadbalancer--reference--group-018.md#canonical-fbb1dd40d1fda7e1edae875c1cbd055920c2d80c38b6ffc35c300d76f96e6770): complete subsection reference.

<a id="canonical-a7d24a31460f21dabb49000ba7a5319648a96f2c172ef66072413314b40c2938"></a>

## Next pages — https.http_protocol_options.http_protocol_enable_v1_only / d0990432f21f / 4

- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-018.md#canonical-fbb1dd40d1fda7e1edae875c1cbd055920c2d80c38b6ffc35c300d76f96e6770)
- [https.http_protocol_options](resources--http_loadbalancer--reference--group-018.md#canonical-6dda791fdf251be6f1dd8108d706d8252bc863f7c0c5a980dde0ec7865caf32a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-fbb1dd40d1fda7e1edae875c1cbd055920c2d80c38b6ffc35c300d76f96e6770"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4c001b5d570029d878b11632b360635b922549ce0f6511d7ddf52f1aec87d57c"></a>

## https.http_protocol_options.http_protocol_enable_v1_only.header_transformation — https.http_protocol_options.http_protocol_enable_v1_only.header_transformation / 974fbe8f63a8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.http_protocol_options](resources--http_loadbalancer--reference--group-018.md#canonical-6dda791fdf251be6f1dd8108d706d8252bc863f7c0c5a980dde0ec7865caf32a)
- [https.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-018.md#canonical-9f443e73d19e56a4af78922ae1db921012ef20f3a711092726357b8d29c14b33)
- https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-c55589a81c8731c458bf3943b684787151d255bd139de99e5254bcfba01e8d80"></a>

Type: `"object"`. single nested block, Optional.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_header_transformation",
    "preserve_case_header_transformation"),
  validators.ConflictingObjectAttributes("default_header_transformation",
    "proper_case_header_transformation"),
  validators.ConflictingObjectAttributes("preserve_case_header_transformation",
    "proper_case_header_transformation")}
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
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

Terraform syntax:

```terraform
header_transformation {
  # Configure direct properties listed below.
}
```

<a id="canonical-45f8048b2a0400c6028843bd79b93109a9e78b4309d2d32b9754967fde24a305"></a>

## Direct properties — https.http_protocol_options.http_protocol_enable_v1_only.header_transformation / 974fbe8f63a8 / 3

- [default_header_transformation](resources--http_loadbalancer--reference--group-018.md#canonical-7bf2e50f256838817efac66e3b80c6ab164ede65af60881cd8247b943486cb01): complete subsection reference.

- [preserve_case_header_transformation](resources--http_loadbalancer--reference--group-018.md#canonical-6dbf4ae2efd857b6a4123a98f7e58da191066c6bceb5ac49a2632131eba20f2b): complete subsection reference.

- [proper_case_header_transformation](resources--http_loadbalancer--reference--group-018.md#canonical-0080c48d32415932bb5e5eb79fb1d0863630094f70829fa7f791f209e1e67d7d): complete subsection reference.

<a id="canonical-acc4836f95c5764cabccfbc2fad0a93e669bed92f68d375e07386d57b8f39b44"></a>

## Next pages — https.http_protocol_options.http_protocol_enable_v1_only.header_transformation / 974fbe8f63a8 / 4

- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--http_loadbalancer--reference--group-018.md#canonical-7bf2e50f256838817efac66e3b80c6ab164ede65af60881cd8247b943486cb01)
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--http_loadbalancer--reference--group-018.md#canonical-6dbf4ae2efd857b6a4123a98f7e58da191066c6bceb5ac49a2632131eba20f2b)
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--http_loadbalancer--reference--group-018.md#canonical-0080c48d32415932bb5e5eb79fb1d0863630094f70829fa7f791f209e1e67d7d)
- [https.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-018.md#canonical-9f443e73d19e56a4af78922ae1db921012ef20f3a711092726357b8d29c14b33)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-7bf2e50f256838817efac66e3b80c6ab164ede65af60881cd8247b943486cb01"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-203c603dd539dfa58500887ce5a65677915da01ab54f64952544596d44908c2b"></a>

## https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.d / 8d573096734b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.http_protocol_options](resources--http_loadbalancer--reference--group-018.md#canonical-6dda791fdf251be6f1dd8108d706d8252bc863f7c0c5a980dde0ec7865caf32a)
- [https.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-018.md#canonical-9f443e73d19e56a4af78922ae1db921012ef20f3a711092726357b8d29c14b33)
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-018.md#canonical-fbb1dd40d1fda7e1edae875c1cbd055920c2d80c38b6ffc35c300d76f96e6770)
- https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-ca4cee65118e38a766eba1cf35dd30bdfa42374e61ca30a2151bef58afbf23f6"></a>

Type: `["object", {}]`. Optional.

Use the platform's current default HTTP header transformation behavior.

Receipt-pinned upstream constraints:

```json
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
default_header_transformation = {}
```

<a id="canonical-0dfe23c1d6a3f5cad66ac407928c919d50d1e8fba56a852c96b18488e72e9d7f"></a>

## Direct properties — https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.d / 8d573096734b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c06c74e8c34cb74b2c9ecbea9863b43367a4d7012832db727af8cd67795d410d"></a>

## Next pages — https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.d / 8d573096734b / 4

- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-018.md#canonical-fbb1dd40d1fda7e1edae875c1cbd055920c2d80c38b6ffc35c300d76f96e6770)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-6dbf4ae2efd857b6a4123a98f7e58da191066c6bceb5ac49a2632131eba20f2b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8cd652f1494039b55ac1679dbec9358c42aa15a3d20018aced1ed10ce885e44f"></a>

## https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.p / b728c204a49a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.http_protocol_options](resources--http_loadbalancer--reference--group-018.md#canonical-6dda791fdf251be6f1dd8108d706d8252bc863f7c0c5a980dde0ec7865caf32a)
- [https.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-018.md#canonical-9f443e73d19e56a4af78922ae1db921012ef20f3a711092726357b8d29c14b33)
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-018.md#canonical-fbb1dd40d1fda7e1edae875c1cbd055920c2d80c38b6ffc35c300d76f96e6770)
- https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-8254234369391a9305d0fc313d66a7e7354076038bd738c1918c8d23428a48de"></a>

Type: `["object", {}]`. Optional.

Preserve HTTP header-name case when upstream case must remain unchanged.

Receipt-pinned upstream constraints:

```json
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
preserve_case_header_transformation = {}
```

<a id="canonical-594fcc32ca857f8bf298e7f95b8ccb378e5714de8661c76a0d87b806c27ef514"></a>

## Direct properties — https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.p / b728c204a49a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dc91538751769ed305aefae9bafc7a83f69aaa3196fbebc7f452950f1a9e1079"></a>

## Next pages — https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.p / b728c204a49a / 4

- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-018.md#canonical-fbb1dd40d1fda7e1edae875c1cbd055920c2d80c38b6ffc35c300d76f96e6770)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-0080c48d32415932bb5e5eb79fb1d0863630094f70829fa7f791f209e1e67d7d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b4b4b7d4451ed2f47df3c3405da866e9d0cef15f96fe98296550437ab0afcb3"></a>

## https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.p / b7b86bdba2e3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.http_protocol_options](resources--http_loadbalancer--reference--group-018.md#canonical-6dda791fdf251be6f1dd8108d706d8252bc863f7c0c5a980dde0ec7865caf32a)
- [https.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-018.md#canonical-9f443e73d19e56a4af78922ae1db921012ef20f3a711092726357b8d29c14b33)
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-018.md#canonical-fbb1dd40d1fda7e1edae875c1cbd055920c2d80c38b6ffc35c300d76f96e6770)
- https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-43fe13c932052eed069632c9c3510d28e3d73f87df034a9f565a719e9d334148"></a>

Type: `["object", {}]`. Optional.

Transform HTTP header names to proper case when explicit transformation is required.

Receipt-pinned upstream constraints:

```json
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
proper_case_header_transformation = {}
```

<a id="canonical-5f70644d2136e0dcedf34a79fd182ee9486756adca52788cc63672bc6af0e33c"></a>

## Direct properties — https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.p / b7b86bdba2e3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e6c3864f6e747be5018af1f60542e165f28c29b3102c68b4b60b5bebfb7b3f1b"></a>

## Next pages — https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.p / b7b86bdba2e3 / 4

- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-018.md#canonical-fbb1dd40d1fda7e1edae875c1cbd055920c2d80c38b6ffc35c300d76f96e6770)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-04dea4678ca8b618cf550fb6c09d3971429d1f8846fa624fbf8230df42b3f455"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-55d16b53b9040edd3e73616ebf3f2b438f635e7cea440232b13326fd80069538"></a>

## https.http_protocol_options.http_protocol_enable_v1_v2 — https.http_protocol_options.http_protocol_enable_v1_v2 / e1acdca3dcc8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.http_protocol_options](resources--http_loadbalancer--reference--group-018.md#canonical-6dda791fdf251be6f1dd8108d706d8252bc863f7c0c5a980dde0ec7865caf32a)
- https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-b86765850dd2e901d493462b3b4797d3153dc151f45bc57971aad5c83201ebdb"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for http protocol enable v1 v2.

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
http_protocol_enable_v1_v2 = {}
```

<a id="canonical-ad52148ebf0d202000ca9045db5a85efbb6c885afc0c07e8a441b8a6f2b09746"></a>

## Direct properties — https.http_protocol_options.http_protocol_enable_v1_v2 / e1acdca3dcc8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ba51da50176000588c20cc6870d305259293c8e6a890d61bf9135b9c50a71836"></a>

## Next pages — https.http_protocol_options.http_protocol_enable_v1_v2 / e1acdca3dcc8 / 4

- [https.http_protocol_options](resources--http_loadbalancer--reference--group-018.md#canonical-6dda791fdf251be6f1dd8108d706d8252bc863f7c0c5a980dde0ec7865caf32a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-27e0020f98056e1bbcf841d1d1bb8ae42be8d924d6352cbd3a6590fc40e34b92"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-125957aea3dea15489abcdc87ba7b27787cea64afadeecca8db440a5df99ae55"></a>

## https.http_protocol_options.http_protocol_enable_v2_only — https.http_protocol_options.http_protocol_enable_v2_only / 87577546dcf3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.http_protocol_options](resources--http_loadbalancer--reference--group-018.md#canonical-6dda791fdf251be6f1dd8108d706d8252bc863f7c0c5a980dde0ec7865caf32a)
- https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-b28ed7024a97cac5f64756ea9ef70eb6ccc30d154be7f9dec8738b04e28b6117"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for http protocol enable v2 only.

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
http_protocol_enable_v2_only = {}
```

<a id="canonical-7d6393e0593f0e80cb0585afe773d2a92dd0d8c87e6f2afe3644ab8841049546"></a>

## Direct properties — https.http_protocol_options.http_protocol_enable_v2_only / 87577546dcf3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-863ef8c5d6aaa69e3c6affaa4149fd76b0865b280bc3e5d37b82b49730725022"></a>

## Next pages — https.http_protocol_options.http_protocol_enable_v2_only / 87577546dcf3 / 4

- [https.http_protocol_options](resources--http_loadbalancer--reference--group-018.md#canonical-6dda791fdf251be6f1dd8108d706d8252bc863f7c0c5a980dde0ec7865caf32a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-da65490816f7086893b9e68a91c94959c7bbf195f4f220138163abc38c176b33"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0d16902299d0635b953f233474b7b602b34af95bc4e1072db0b67d3222bd1e72"></a>

## https.non_default_loadbalancer — https.non_default_loadbalancer / 2bbfb8cfcf0f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- https.non_default_loadbalancer

<a id="canonical-ba4cbb9833f8f126449cf7a7ee0f1719bab2d5aeee8b52fc3eeaff9fc45372d3"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for non default loadbalancer.

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
non_default_loadbalancer = {}
```

<a id="canonical-a2f2a843d9b2d16c322c6be53c129cb67aceb188206cde164b3d63213ddae11f"></a>

## Direct properties — https.non_default_loadbalancer / 2bbfb8cfcf0f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ff9cd0cebf1bd0c3506805eae8a45b67f6abc898ec6bf8584e8b05baaf10453f"></a>

## Next pages — https.non_default_loadbalancer / 2bbfb8cfcf0f / 4

- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-7827b7da9d7a3b052a4b08639f9759e03c95eb68f26d8562e112c4de19300d11"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dce5f1d11372309516ba76cf12695273af125cd709a232ee6ac3bdf9722f96d5"></a>

## https.pass_through — https.pass_through / b0524d3d70bb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- https.pass_through

<a id="canonical-6f7795968b86c1dbd5f5aab47865e0ba81cd33fa36b9839cf5b670fa0197852e"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for pass through.

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
pass_through = {}
```

<a id="canonical-efac8a0737408292aeffcd1876d0b935d6d24a28012995b03f911c649629759a"></a>

## Direct properties — https.pass_through / b0524d3d70bb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-435a44f28c3b2ba5b1eafbd51eff1646f3b522ce2a05be24a37a278fdb6b055e"></a>

## Next pages — https.pass_through / b0524d3d70bb / 4

- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-31ca63fa137c47bc6b3bc3049c6a66381c9558e74ba34ad4e12c7ace5ed5a91f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5126032b732540c213b79122c4ca852766813232e1af78c2dc79cd8955804ff4"></a>

## https.tls_cert_params — https.tls_cert_params / 4b91e0e71301 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- https.tls_cert_params

<a id="canonical-efc459d1403569d104fa4796b3e0779f38f79b90fdefdf4278a37d298f8f436b"></a>

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

<a id="canonical-fd630f27e4dde27175fc941015c44924e1f795ffd9288f3e8a49a52c7cc8dc1d"></a>

## Direct properties — https.tls_cert_params / 4b91e0e71301 / 3

- [certificates](resources--http_loadbalancer--reference--group-018.md#canonical-8b9964ccefd4042c438500b02fe054ff444d6121048cc2e1f3a4060130bf798d): complete subsection reference.

- [no_mtls](resources--http_loadbalancer--reference--group-018.md#canonical-9a561e4c15a412ab30ddc0b7a9fb5868f624cb059ed9e81b036fd0fde925bacb): complete subsection reference.

- [tls_config](resources--http_loadbalancer--reference--group-018.md#canonical-2d4005d2832138352c6d05785312ffac8aa2fecfc2b55fb2d03b3874baf8da58): complete subsection reference.

- [use_mtls](resources--http_loadbalancer--reference--group-018.md#canonical-8c0ad9d7be7da72c3cbec82d199436133f816b111c68e517b73bde71c4aa45ac): complete subsection reference.

<a id="canonical-e326d970d698584501be6c7fb90b63115af5869d3bd4d47643f9492d664e6df2"></a>

## Next pages — https.tls_cert_params / 4b91e0e71301 / 4

- [https.tls_cert_params.certificates](resources--http_loadbalancer--reference--group-018.md#canonical-8b9964ccefd4042c438500b02fe054ff444d6121048cc2e1f3a4060130bf798d)
- [https.tls_cert_params.no_mtls](resources--http_loadbalancer--reference--group-018.md#canonical-9a561e4c15a412ab30ddc0b7a9fb5868f624cb059ed9e81b036fd0fde925bacb)
- [https.tls_cert_params.tls_config](resources--http_loadbalancer--reference--group-018.md#canonical-2d4005d2832138352c6d05785312ffac8aa2fecfc2b55fb2d03b3874baf8da58)
- [https.tls_cert_params.use_mtls](resources--http_loadbalancer--reference--group-018.md#canonical-8c0ad9d7be7da72c3cbec82d199436133f816b111c68e517b73bde71c4aa45ac)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-8b9964ccefd4042c438500b02fe054ff444d6121048cc2e1f3a4060130bf798d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c0326224b73b3a930cb4a16535170d95a315ede8cc435c7971fa92592a1cb0c3"></a>

## https.tls_cert_params.certificates — https.tls_cert_params.certificates / ed389b032884 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-018.md#canonical-31ca63fa137c47bc6b3bc3049c6a66381c9558e74ba34ad4e12c7ace5ed5a91f)
- https.tls_cert_params.certificates

<a id="canonical-4ee6367b0ca72e2bb052e16c953e9038dae4e23e1421e394adce2c584ff9794f"></a>

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

<a id="canonical-ac12d1cb9b72a1af62fb3cf5e5db0507a804357531db3967c59d79874ac37883"></a>

## Direct properties — https.tls_cert_params.certificates / ed389b032884 / 3

<a id="canonical-c3937bae039b29c5786d67c37d8727e57d562b06971fb896395939ea0a3966e5"></a>

<a id="canonical-ea96824c25e1cd8d5195d61ed9c90d0e1873291bc8d0b296a371c6d294548cc9"></a>

## name property — https.tls_cert_params.certificates / ed389b032884 / 4

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

<a id="canonical-4049e862a63fe4c2f2b0b5440c92bba23996cebfb9b9c61cecd7ecd8208494b3"></a>

<a id="canonical-f7e5cf639b50786da235431c9b4141fe6b8386e52f0420e9b9888c81f741e58d"></a>

## namespace property — https.tls_cert_params.certificates / ed389b032884 / 5

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

<a id="canonical-d9a3d448cec70d7b43d8ccb184a1af19e67e5ec79b659d38cfc39267b250f40b"></a>

<a id="canonical-6ace2ed7e5c3c84e54cdf5008290b4aac6b326bd312725f6f1f2255b5d074217"></a>

## tenant property — https.tls_cert_params.certificates / ed389b032884 / 6

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

<a id="canonical-00a17f3b869c3a206e463154691224f81ab169c65bc025e75f2d26d7857d01f0"></a>

## Next pages — https.tls_cert_params.certificates / ed389b032884 / 7

- [https.tls_cert_params](resources--http_loadbalancer--reference--group-018.md#canonical-31ca63fa137c47bc6b3bc3049c6a66381c9558e74ba34ad4e12c7ace5ed5a91f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-9a561e4c15a412ab30ddc0b7a9fb5868f624cb059ed9e81b036fd0fde925bacb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d909c7d67196986b3ce134df64cf30c7591f3afe6a5e193050195d64690dd24"></a>

## https.tls_cert_params.no_mtls — https.tls_cert_params.no_mtls / ea8e3a039741 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-018.md#canonical-31ca63fa137c47bc6b3bc3049c6a66381c9558e74ba34ad4e12c7ace5ed5a91f)
- https.tls_cert_params.no_mtls

<a id="canonical-5e9ccd8faad227438c298727d739a1c282f698a96e3cb421d45339a1be48131c"></a>

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

<a id="canonical-9971d32010459d541b8805ca2deb0472772a343a82a5abf4cc1591ba4da9690b"></a>

## Direct properties — https.tls_cert_params.no_mtls / ea8e3a039741 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4cd18bbb35518c734be065d29ff999481618ef66a33c53b5334faf0128a17619"></a>

## Next pages — https.tls_cert_params.no_mtls / ea8e3a039741 / 4

- [https.tls_cert_params](resources--http_loadbalancer--reference--group-018.md#canonical-31ca63fa137c47bc6b3bc3049c6a66381c9558e74ba34ad4e12c7ace5ed5a91f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-2d4005d2832138352c6d05785312ffac8aa2fecfc2b55fb2d03b3874baf8da58"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3dd27e9cb068d097f263c505332503e6966f7a2eafbbd681ab4e9156609d781b"></a>

## https.tls_cert_params.tls_config — https.tls_cert_params.tls_config / 02fe00e87729 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-018.md#canonical-31ca63fa137c47bc6b3bc3049c6a66381c9558e74ba34ad4e12c7ace5ed5a91f)
- https.tls_cert_params.tls_config

<a id="canonical-aada7e5c932414123a3cb7360ff7e913d4f75dadbd9c23283bf6315b2141c72e"></a>

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

<a id="canonical-6a50c15165d14a403e606aa0828af2f8ac62df1b330450db6d7e3842329f0426"></a>

## Direct properties — https.tls_cert_params.tls_config / 02fe00e87729 / 3

- [custom_security](resources--http_loadbalancer--reference--group-018.md#canonical-3c3caf67b5d9aad4b662d388cc9e7c26b312383bbd84577a94fe553411f88fa8): complete subsection reference.

- [default_security](resources--http_loadbalancer--reference--group-018.md#canonical-27e8c0f4815c2a30c697ec951d8f3e6d30feb89026141f5ff3b0461c0b5676e0): complete subsection reference.

- [low_security](resources--http_loadbalancer--reference--group-018.md#canonical-01907a4545ec4b03a9f9035776ce60296b169903f28e153f2f12fa0e60f33b3a): complete subsection reference.

- [medium_security](resources--http_loadbalancer--reference--group-018.md#canonical-b91daccc4f4816450d234fc61cb49885463dd4be0357c49dca4245c17f1f58fa): complete subsection reference.

<a id="canonical-54549ca3a14c63a280ecae3e306698595b17fe615a322f7a7d340c346f711d61"></a>

## Next pages — https.tls_cert_params.tls_config / 02fe00e87729 / 4

- [https.tls_cert_params.tls_config.custom_security](resources--http_loadbalancer--reference--group-018.md#canonical-3c3caf67b5d9aad4b662d388cc9e7c26b312383bbd84577a94fe553411f88fa8)
- [https.tls_cert_params.tls_config.default_security](resources--http_loadbalancer--reference--group-018.md#canonical-27e8c0f4815c2a30c697ec951d8f3e6d30feb89026141f5ff3b0461c0b5676e0)
- [https.tls_cert_params.tls_config.low_security](resources--http_loadbalancer--reference--group-018.md#canonical-01907a4545ec4b03a9f9035776ce60296b169903f28e153f2f12fa0e60f33b3a)
- [https.tls_cert_params.tls_config.medium_security](resources--http_loadbalancer--reference--group-018.md#canonical-b91daccc4f4816450d234fc61cb49885463dd4be0357c49dca4245c17f1f58fa)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-018.md#canonical-31ca63fa137c47bc6b3bc3049c6a66381c9558e74ba34ad4e12c7ace5ed5a91f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-3c3caf67b5d9aad4b662d388cc9e7c26b312383bbd84577a94fe553411f88fa8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d44b664f78a19dc6e385547011d9268d23753401576fcea8c40c282aafd59ed"></a>

## https.tls_cert_params.tls_config.custom_security — https.tls_cert_params.tls_config.custom_security / b7cd80cde9c8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-018.md#canonical-31ca63fa137c47bc6b3bc3049c6a66381c9558e74ba34ad4e12c7ace5ed5a91f)
- [https.tls_cert_params.tls_config](resources--http_loadbalancer--reference--group-018.md#canonical-2d4005d2832138352c6d05785312ffac8aa2fecfc2b55fb2d03b3874baf8da58)
- https.tls_cert_params.tls_config.custom_security

<a id="canonical-6aa64a59572d21430735ad4c9d3973bb643556a20736d852286e25d1daa99d7c"></a>

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

<a id="canonical-1ee49b35f7d0cad06aecc00b4c28d76b83ecff502d7b4be990f10b0318593d5f"></a>

## Direct properties — https.tls_cert_params.tls_config.custom_security / b7cd80cde9c8 / 3

<a id="canonical-8ab773e5d7a56df1d0b516a6cda454e2848701ab8cc85cce021b5e1034d93264"></a>

<a id="canonical-19776748e44f5512265c774da9382ce59fe95c32c263d5dd366d4ac0cbf8d410"></a>

## cipher_suites property — https.tls_cert_params.tls_config.custom_security / b7cd80cde9c8 / 4

Type: `["list", "string"]`. Optional.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-efaccb457799cc24fdf7da8e67269c5bff65130a7ad428652242557a11121624"></a>

<a id="canonical-88f50fd7c0d296d481e42bd604122a28b141da8fbe5ef8b2dea3585f4fd9f826"></a>

## max_version property — https.tls_cert_params.tls_config.custom_security / b7cd80cde9c8 / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-d2c533de89bb4668c02d6b145b36f3d1b9cf7b26af0d839ad0d74859742697ff"></a>

<a id="canonical-5b6bf1f98d425ca8a65149993512ce78e12a7fb928687371d9296bfdf6a7d60d"></a>

## min_version property — https.tls_cert_params.tls_config.custom_security / b7cd80cde9c8 / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-42087e22c4e617898816951e35097e2292a1156861dca21f8a7cee3a426f5a2f"></a>

## Next pages — https.tls_cert_params.tls_config.custom_security / b7cd80cde9c8 / 7

- [https.tls_cert_params.tls_config](resources--http_loadbalancer--reference--group-018.md#canonical-2d4005d2832138352c6d05785312ffac8aa2fecfc2b55fb2d03b3874baf8da58)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-27e8c0f4815c2a30c697ec951d8f3e6d30feb89026141f5ff3b0461c0b5676e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce86a291087026001ee470945f8135605abb05f8b9ec5e857c663c3733b0ce1e"></a>

## https.tls_cert_params.tls_config.default_security — https.tls_cert_params.tls_config.default_security / a971c10a78bf / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-018.md#canonical-31ca63fa137c47bc6b3bc3049c6a66381c9558e74ba34ad4e12c7ace5ed5a91f)
- [https.tls_cert_params.tls_config](resources--http_loadbalancer--reference--group-018.md#canonical-2d4005d2832138352c6d05785312ffac8aa2fecfc2b55fb2d03b3874baf8da58)
- https.tls_cert_params.tls_config.default_security

<a id="canonical-c3b5a4b4fe7baa79e155a43aa6fb37a1dc4d4d191600a5d45ff8c1c4d35594c0"></a>

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
default_security = {}
```

<a id="canonical-fdf938427ead3757ff99139d68477b091172dab96896625d6d7d36522a6748a5"></a>

## Direct properties — https.tls_cert_params.tls_config.default_security / a971c10a78bf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-68d360b44f695d75ab68877cd44c16ef07f1611c69b78e06c40a6f3a12d56852"></a>

## Next pages — https.tls_cert_params.tls_config.default_security / a971c10a78bf / 4

- [https.tls_cert_params.tls_config](resources--http_loadbalancer--reference--group-018.md#canonical-2d4005d2832138352c6d05785312ffac8aa2fecfc2b55fb2d03b3874baf8da58)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-01907a4545ec4b03a9f9035776ce60296b169903f28e153f2f12fa0e60f33b3a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-293b5e4d567905aa037a5191964dac75e3e67ad475401ea9d7a827f9042c4df8"></a>

## https.tls_cert_params.tls_config.low_security — https.tls_cert_params.tls_config.low_security / 456414dbc5db / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-018.md#canonical-31ca63fa137c47bc6b3bc3049c6a66381c9558e74ba34ad4e12c7ace5ed5a91f)
- [https.tls_cert_params.tls_config](resources--http_loadbalancer--reference--group-018.md#canonical-2d4005d2832138352c6d05785312ffac8aa2fecfc2b55fb2d03b3874baf8da58)
- https.tls_cert_params.tls_config.low_security

<a id="canonical-263f79852d3d3670323159e45eb8cfa0246b381919266ff1e9f31095356199cb"></a>

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
low_security = {}
```

<a id="canonical-a30cf61bd447804dcdf7f48327db3a669cee2764ec5c31abfbd10dd2ba434fee"></a>

## Direct properties — https.tls_cert_params.tls_config.low_security / 456414dbc5db / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-16c2b394ee6afa7bde3561f2df969fdb37d141acc05b88b5e81a7c034d8d9093"></a>

## Next pages — https.tls_cert_params.tls_config.low_security / 456414dbc5db / 4

- [https.tls_cert_params.tls_config](resources--http_loadbalancer--reference--group-018.md#canonical-2d4005d2832138352c6d05785312ffac8aa2fecfc2b55fb2d03b3874baf8da58)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b91daccc4f4816450d234fc61cb49885463dd4be0357c49dca4245c17f1f58fa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e9575379a27ff24758c0f4302e91cc019d3cfaa35dd834b5053e414669753133"></a>

## https.tls_cert_params.tls_config.medium_security — https.tls_cert_params.tls_config.medium_security / 9019575f77c2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-018.md#canonical-31ca63fa137c47bc6b3bc3049c6a66381c9558e74ba34ad4e12c7ace5ed5a91f)
- [https.tls_cert_params.tls_config](resources--http_loadbalancer--reference--group-018.md#canonical-2d4005d2832138352c6d05785312ffac8aa2fecfc2b55fb2d03b3874baf8da58)
- https.tls_cert_params.tls_config.medium_security

<a id="canonical-64892e88bde11ab1ceb9e03821e139f87a440755a71888eb62c97b3d640b5160"></a>

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
medium_security = {}
```

<a id="canonical-3f2026c353ed80615ec7af104ae54d236fad4279df31e94eaa21e9044af67748"></a>

## Direct properties — https.tls_cert_params.tls_config.medium_security / 9019575f77c2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-85dd7243895e0c6ccf81eac1a867bbb6132e73703ffc473fb19f831b50b5ae10"></a>

## Next pages — https.tls_cert_params.tls_config.medium_security / 9019575f77c2 / 4

- [https.tls_cert_params.tls_config](resources--http_loadbalancer--reference--group-018.md#canonical-2d4005d2832138352c6d05785312ffac8aa2fecfc2b55fb2d03b3874baf8da58)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-8c0ad9d7be7da72c3cbec82d199436133f816b111c68e517b73bde71c4aa45ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-57c1457543d1d4ae2d799a8a5ce2440aa0cc3236117a4987acc2370c47799844"></a>

## https.tls_cert_params.use_mtls — https.tls_cert_params.use_mtls / 20954948bc90 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-018.md#canonical-31ca63fa137c47bc6b3bc3049c6a66381c9558e74ba34ad4e12c7ace5ed5a91f)
- https.tls_cert_params.use_mtls

<a id="canonical-18bdd9acc37471158443fad3022df16e7a059c426c57d46b5e994fc2168637f3"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("crl",
    "no_crl"),
  validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url"),
  validators.ConflictingObjectAttributes("xfcc_disabled",
    "xfcc_options")}
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
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-a872536abc2ed5637a91ee06276ddbdf77d1a55ece36843d9835f63888a91190"></a>

## Direct properties — https.tls_cert_params.use_mtls / 20954948bc90 / 3

<a id="canonical-07a0dd9d2eacdbf6655fb1ca31e602c03b5ba333f3cadebd81c8e0ed08033311"></a>

<a id="canonical-57fe58b7df655377035f2facf29744e7f2a39e6638f3f898fbd47c4aaf4ee4b3"></a>

## client_certificate_optional property — https.tls_cert_params.use_mtls / 20954948bc90 / 4

Type: `"bool"`. Optional.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated.

Upstream description:

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [crl](resources--http_loadbalancer--reference--group-018.md#canonical-27ac75a7ee45aa6532c31eaf0d61e515cd7f446ed5cf441a5bb6a61ce211ffe3): complete subsection reference.

- [no_crl](resources--http_loadbalancer--reference--group-018.md#canonical-615161c0f829591e8ad7bfe1a48fbbe969f5a034d2c1c9e799e6afdc414140a7): complete subsection reference.

- [trusted_ca](resources--http_loadbalancer--reference--group-019.md#canonical-b02d79626ed0d98f53e005718b6fa3a30bc6c7deea38f903da99d20c3db82d87): complete subsection reference.

<a id="canonical-e26a6f5891346f6cc872dacb5e1911e6bf8461fffc638aaf2161bf9dd59e2a44"></a>

<a id="canonical-220f223de542e972ffe5f2ed3149c2b302c60efaa2bf11de64061f1abd24ca3d"></a>

## trusted_ca_url property — https.tls_cert_params.use_mtls / 20954948bc90 / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

- [xfcc_disabled](resources--http_loadbalancer--reference--group-019.md#canonical-4b63d7f3146ff0f3ebc46e220a87da7eafb5412b0329dc547df8915e04206e7c): complete subsection reference.

- [xfcc_options](resources--http_loadbalancer--reference--group-019.md#canonical-8c3b0c552caaf6526988ea4b0ae3e98ff612b2909697f15ed31c82ce519e2023): complete subsection reference.

<a id="canonical-5b6e32bebe7c4bc3e9d6341ea61940c3c2bcf15fe14e210912681d87b7764338"></a>

## Next pages — https.tls_cert_params.use_mtls / 20954948bc90 / 6

- [https.tls_cert_params.use_mtls.crl](resources--http_loadbalancer--reference--group-018.md#canonical-27ac75a7ee45aa6532c31eaf0d61e515cd7f446ed5cf441a5bb6a61ce211ffe3)
- [https.tls_cert_params.use_mtls.no_crl](resources--http_loadbalancer--reference--group-018.md#canonical-615161c0f829591e8ad7bfe1a48fbbe969f5a034d2c1c9e799e6afdc414140a7)
- [https.tls_cert_params.use_mtls.trusted_ca](resources--http_loadbalancer--reference--group-019.md#canonical-b02d79626ed0d98f53e005718b6fa3a30bc6c7deea38f903da99d20c3db82d87)
- [https.tls_cert_params.use_mtls.xfcc_disabled](resources--http_loadbalancer--reference--group-019.md#canonical-4b63d7f3146ff0f3ebc46e220a87da7eafb5412b0329dc547df8915e04206e7c)
- [https.tls_cert_params.use_mtls.xfcc_options](resources--http_loadbalancer--reference--group-019.md#canonical-8c3b0c552caaf6526988ea4b0ae3e98ff612b2909697f15ed31c82ce519e2023)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-018.md#canonical-31ca63fa137c47bc6b3bc3049c6a66381c9558e74ba34ad4e12c7ace5ed5a91f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-27ac75a7ee45aa6532c31eaf0d61e515cd7f446ed5cf441a5bb6a61ce211ffe3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8518af15c7d774b28055e20b8c8f656f7c6983fb8795535cea98fb77916bc683"></a>

## https.tls_cert_params.use_mtls.crl — https.tls_cert_params.use_mtls.crl / 783dfc0e9b54 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-018.md#canonical-31ca63fa137c47bc6b3bc3049c6a66381c9558e74ba34ad4e12c7ace5ed5a91f)
- [https.tls_cert_params.use_mtls](resources--http_loadbalancer--reference--group-018.md#canonical-8c0ad9d7be7da72c3cbec82d199436133f816b111c68e517b73bde71c4aa45ac)
- https.tls_cert_params.use_mtls.crl

<a id="canonical-86ea31c4c4f6de58b0fedf3ee8b188b9afe658deb894902de33d807085d72d71"></a>

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
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-18deeacfd7af4c311ac76d61519c21975ece9831e0a4ff575241023fab3ccf58"></a>

## Direct properties — https.tls_cert_params.use_mtls.crl / 783dfc0e9b54 / 3

<a id="canonical-bb2e3fccf0e2357dd1f9807157f1a3d5737beade627d151b540d60d12f3c116e"></a>

<a id="canonical-cb0cc8c97cb2e319d04b535c3f8fea09669732a1068ecd4a73e1476c6273050b"></a>

## name property — https.tls_cert_params.use_mtls.crl / 783dfc0e9b54 / 4

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

<a id="canonical-7e2b514d86370276a0cd10c38d1a96218cea8d933d533c9075e489277a0ffcf4"></a>

<a id="canonical-246c11f3cec6749cde561ac21076aa9967c7b92e14efda0db3966475d9b635c2"></a>

## namespace property — https.tls_cert_params.use_mtls.crl / 783dfc0e9b54 / 5

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

<a id="canonical-1c15493f658c4dfa1bbf7278a696ab7fedb719b8707d6dfe3e1ed45a7ab7c051"></a>

<a id="canonical-49fa2fe8eefd8c4403732907f370a09c81cd4b9c7717207f6a18a9a925a20d0c"></a>

## tenant property — https.tls_cert_params.use_mtls.crl / 783dfc0e9b54 / 6

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

<a id="canonical-559648690f6efce312f961301bb7a95bbe11e02dab8959b80e034729ae87c118"></a>

## Next pages — https.tls_cert_params.use_mtls.crl / 783dfc0e9b54 / 7

- [https.tls_cert_params.use_mtls](resources--http_loadbalancer--reference--group-018.md#canonical-8c0ad9d7be7da72c3cbec82d199436133f816b111c68e517b73bde71c4aa45ac)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-615161c0f829591e8ad7bfe1a48fbbe969f5a034d2c1c9e799e6afdc414140a7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-03e65f138b1ca67bd2e2d99d3fe5d59b403886c57c62cd3b56f2164955828826"></a>

## https.tls_cert_params.use_mtls.no_crl — https.tls_cert_params.use_mtls.no_crl / 30afa9b64747 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-018.md#canonical-31ca63fa137c47bc6b3bc3049c6a66381c9558e74ba34ad4e12c7ace5ed5a91f)
- [https.tls_cert_params.use_mtls](resources--http_loadbalancer--reference--group-018.md#canonical-8c0ad9d7be7da72c3cbec82d199436133f816b111c68e517b73bde71c4aa45ac)
- https.tls_cert_params.use_mtls.no_crl

<a id="canonical-b53494e1aae2210bd8a7bb2b875dcf130389684c9428671f7fe99577288ad776"></a>

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
no_crl = {}
```

<a id="canonical-91aa6a9afacc834885fe6008b9cf9bdf24aebf9e3e9876d454f57cb4f495dbb0"></a>

## Direct properties — https.tls_cert_params.use_mtls.no_crl / 30afa9b64747 / 3

This is an empty object or choice marker. It has no direct properties.
