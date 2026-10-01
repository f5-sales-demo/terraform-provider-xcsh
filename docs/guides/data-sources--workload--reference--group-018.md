---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-b2257317c77ad9d7c8c0553d41b7dd359f8f0bf2b306a4dae1ba09c4a3911fb3"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 8a73b670f314 / 6

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl](data-sources--workload--reference--group-018.md#canonical-1d67abd8348bfcdd4c0024b590371e2f6014232daceb4c0fb36a1d1dc84747cd)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl](data-sources--workload--reference--group-018.md#canonical-7bb8d8864d1806147ba0eb63643fcc179e30f75a656fd059d86d6db6c67c1663)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca](data-sources--workload--reference--group-018.md#canonical-a8e0718b1f2c2483f240bd02ff224a46e4f78e75db80065e7c9ad601944ea7c5)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled](data-sources--workload--reference--group-018.md#canonical-3724c86831aa76f3960aa7df9df8c88832dba6d97564a1b378c8ae14a705e4e3)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options](data-sources--workload--reference--group-018.md#canonical-6a8b473d36dd8abdfe8b143b1fc1f75515626b43b7dc31fdf7221254ff06f9b6)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-017.md#canonical-84045b21c8578756bf25fa652058fae0c30afa065bb3e0c1132ec32e6269fc78)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-1d67abd8348bfcdd4c0024b590371e2f6014232daceb4c0fb36a1d1dc84747cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-50054965842522a8996f7e13c85fbede40efb1bf01b56c8fb1a7798d626818e9"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 88ea64b27ae6 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-017.md#canonical-84045b21c8578756bf25fa652058fae0c30afa065bb3e0c1132ec32e6269fc78)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-017.md#canonical-b7c1216887c2c56b699d1fcaef1d8773467395aa0f4bbb9adf802d0b0fe80ad3)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl

<a id="canonical-b8989b2d31be554297be531b1d6ca862a93c4d64575377c247b9fb150aa31e85"></a>

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

<a id="canonical-325c92e4ba98f14fb3814200f2c794ccaf652026d0f74c53145c48c270e6d44b"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 88ea64b27ae6 / 3

<a id="canonical-312a4bba980bf0778710d891947244208422c6474104bb78965f4f89e36b35f2"></a>

<a id="canonical-4fd40987a440955538ef743639fa3b0366f4a9aedef1452c4d64a09fc58e0ef6"></a>

## name property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 88ea64b27ae6 / 4

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

<a id="canonical-7935b606291d6df294e2ade62a522e1cded53287da9454dbe5af5e602fe5c202"></a>

<a id="canonical-5a1c5aecdd08681fcec88729848d5a690e84096e3d62ef2cd863a6b580b53fe8"></a>

## namespace property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 88ea64b27ae6 / 5

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

<a id="canonical-f22700b70dfbb8c0d02bb90b79d4f87f47e0a3a1b61e1d87ced886a40a2b48a6"></a>

<a id="canonical-baf8a835ce29d60077f9ac60c6baaa080ffbac20802f1825d2af6651715b4db8"></a>

## tenant property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 88ea64b27ae6 / 6

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

<a id="canonical-1cbdb6d453cf7aaa16853d34f225ff8a5838b5132da57c83ff59fba0409aab93"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 88ea64b27ae6 / 7

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-017.md#canonical-b7c1216887c2c56b699d1fcaef1d8773467395aa0f4bbb9adf802d0b0fe80ad3)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-7bb8d8864d1806147ba0eb63643fcc179e30f75a656fd059d86d6db6c67c1663"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e54427f3c1751bfb5f11ab74d7188e86decc26ebab7366d5f1c8af93d37aa536"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 223cab09f0a8 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-017.md#canonical-84045b21c8578756bf25fa652058fae0c30afa065bb3e0c1132ec32e6269fc78)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-017.md#canonical-b7c1216887c2c56b699d1fcaef1d8773467395aa0f4bbb9adf802d0b0fe80ad3)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl

<a id="canonical-cbdbcda2615e3e1519adb29a3547f35a140f2b29d8b84fa335cdd53918080ba7"></a>

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

<a id="canonical-78df6e3766425ec314879141052ac6f7d6e5093595d0017816610a324a90f253"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 223cab09f0a8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c93fe20671654618faf28508341de0f0f7649169f8016d1fef6b66b3867266d2"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 223cab09f0a8 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-017.md#canonical-b7c1216887c2c56b699d1fcaef1d8773467395aa0f4bbb9adf802d0b0fe80ad3)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-a8e0718b1f2c2483f240bd02ff224a46e4f78e75db80065e7c9ad601944ea7c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-caca5ea4d411d8792d267f55e58114c9f24360485f2e6ac850fce5da40598853"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 6ff8a83c7e4e / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-017.md#canonical-84045b21c8578756bf25fa652058fae0c30afa065bb3e0c1132ec32e6269fc78)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-017.md#canonical-b7c1216887c2c56b699d1fcaef1d8773467395aa0f4bbb9adf802d0b0fe80ad3)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-4b86807a3d9fe9d6d38769f230bcc9d76a17e4c5cd7d3f4484e4f94be991e019"></a>

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

<a id="canonical-09d68a6ab1ede18158a741868dc9763144e0b35d815370d11ab3b93f67bfe770"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 6ff8a83c7e4e / 3

<a id="canonical-ad2950be804effd9731f0ed833e7ee71863872b00b94a0a3b1b3b775a3475864"></a>

<a id="canonical-505144dfe1a44b0a5853588361c7fec84c5c766be18c954530d87b168c25fcdc"></a>

## name property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 6ff8a83c7e4e / 4

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

<a id="canonical-893a6a42044621ef514394637da321cf551dea047992ec9d2cfc499a1e2e17fc"></a>

<a id="canonical-1e2e8b0fe59081bea92acc0afdb5f1cf509b333006581c0f501db6e72e62159b"></a>

## namespace property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 6ff8a83c7e4e / 5

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

<a id="canonical-8accd4e91f883c74b47666745cb8122bf58f743d1e3a4ee1ef9c79b323b450d0"></a>

<a id="canonical-98263683ee08a82d2c2032606ef7def1712d0dd760605c6f5488b24147769130"></a>

## tenant property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 6ff8a83c7e4e / 6

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

<a id="canonical-1e08445b354e160b1ebff384f45b8aebead1a4a9ac039c5980a41aa5bc306855"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 6ff8a83c7e4e / 7

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-017.md#canonical-b7c1216887c2c56b699d1fcaef1d8773467395aa0f4bbb9adf802d0b0fe80ad3)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-3724c86831aa76f3960aa7df9df8c88832dba6d97564a1b378c8ae14a705e4e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0ee0e8448a346ddc681414ef30fa688b48cc82c5052b1c25432bd82e568e019a"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 4d8e08c961df / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-017.md#canonical-84045b21c8578756bf25fa652058fae0c30afa065bb3e0c1132ec32e6269fc78)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-017.md#canonical-b7c1216887c2c56b699d1fcaef1d8773467395aa0f4bbb9adf802d0b0fe80ad3)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-45ddc556999de06229e118c071100f3f8162994af46d4434b0efb72a70c1c27a"></a>

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

<a id="canonical-5c61976e4060c5db755f2973a76b21f83206ebbd47e0b2702e75a80d6a7a9d3d"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 4d8e08c961df / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-957a32a866f6e0be7922849dd5ed149911d25d131b0ccadca0583d38ca21e91f"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 4d8e08c961df / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-017.md#canonical-b7c1216887c2c56b699d1fcaef1d8773467395aa0f4bbb9adf802d0b0fe80ad3)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-6a8b473d36dd8abdfe8b143b1fc1f75515626b43b7dc31fdf7221254ff06f9b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-521010f4b785e8041ffaf452c9e9f602e2c57062f13c28d7135c6622f2ff8067"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 7ac1fa31e9a8 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-017.md#canonical-84045b21c8578756bf25fa652058fae0c30afa065bb3e0c1132ec32e6269fc78)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-017.md#canonical-b7c1216887c2c56b699d1fcaef1d8773467395aa0f4bbb9adf802d0b0fe80ad3)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-731d93e8ab5ee423d8d86ce56c8f638570f483b02c90f502b3948cf5f409573a"></a>

Type: `"single"`. Computed.

X-Forwarded-Client-Cert header elements to be added to requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-c79f54b33792762c0e336fc2d9b3aa8a7486431d8d0f4faef866b7ff1e75ddd9"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 7ac1fa31e9a8 / 3

<a id="canonical-7267ac27ed87de1f5a13a0de82e0fed0cf3a6a9241a341f082692b16d0e48985"></a>

<a id="canonical-ec72cfd4dca468b8ab545f07454bdee9bc1094b976b0e5ee2d46ffe0f66e32a2"></a>

## xfcc_header_elements property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 7ac1fa31e9a8 / 4

Type: `["list", "string"]`. Computed.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be added to requests.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-08458c097bb7d22bea161ff21e927c5e3ba2e7c884a54de2ddfde67389328fb3"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 7ac1fa31e9a8 / 5

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-017.md#canonical-b7c1216887c2c56b699d1fcaef1d8773467395aa0f4bbb9adf802d0b0fe80ad3)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-ecd89b3af405ecf7541968a7fceb933b59a697d8b94962ff85990a43286152b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0ad00430040b318db11ebf8e8cb9431484a280ce628f300309812b28b647f903"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 05771147960c / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters

<a id="canonical-fb7f16bd700effa725b332d02312a0c06031124ff0a7b8bfa9e37ae17a68f24c"></a>

Type: `"single"`. Computed.

Configuration parameter for tls parameters.

Upstream description:

Inline TLS parameters.

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

<a id="canonical-a6d5abf42b9e278bd95ba99621e4306731b40fd1b8f3d961f55f3e67ea2e0d57"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 05771147960c / 3

- [no_mtls](data-sources--workload--reference--group-018.md#canonical-bee22c1ec9310f5c5aefe2f4cdf54f16be7dfb03fef9f74217d126b028a9b391): complete subsection reference.

- [tls_certificates](data-sources--workload--reference--group-018.md#canonical-d2d0afcebbbd23bff82eb0a31e4303ec158f4f3de2762a3b633f63d68ffe5a90): complete subsection reference.

- [tls_config](data-sources--workload--reference--group-018.md#canonical-bac05e1b3b9c72c78658b5b74a11c1e8ce5989ec52946a1f8dea5591c3875544): complete subsection reference.

- [use_mtls](data-sources--workload--reference--group-018.md#canonical-c431c7272fd63a6532c467f52bccbb1eff240da38d1148190cbd6c4484a2e7c1): complete subsection reference.

<a id="canonical-a09e1e69a99b126f115cff30cc8af612edc174efd534313de34bc931a344d822"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 05771147960c / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.no_mtls](data-sources--workload--reference--group-018.md#canonical-bee22c1ec9310f5c5aefe2f4cdf54f16be7dfb03fef9f74217d126b028a9b391)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-018.md#canonical-d2d0afcebbbd23bff82eb0a31e4303ec158f4f3de2762a3b633f63d68ffe5a90)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-018.md#canonical-bac05e1b3b9c72c78658b5b74a11c1e8ce5989ec52946a1f8dea5591c3875544)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-018.md#canonical-c431c7272fd63a6532c467f52bccbb1eff240da38d1148190cbd6c4484a2e7c1)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-bee22c1ec9310f5c5aefe2f4cdf54f16be7dfb03fef9f74217d126b028a9b391"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3be8dc311b0d020965e974496527f4ffb33b37cd64440f441ffc51191019415e"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.no_mtls — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / f2db5d7d67ce / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-018.md#canonical-ecd89b3af405ecf7541968a7fceb933b59a697d8b94962ff85990a43286152b2)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.no_mtls

<a id="canonical-59332ae1a41cca2ea5bfae979c47b5c50450b0a692fde28aafbce5a85b661504"></a>

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

<a id="canonical-ae240168da9d31c39c8410a4c4c4e9e8b40a0d79e567188d8e9ab0f75c947ed9"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / f2db5d7d67ce / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ffd205e1ae9fe38aee3a6956e12c539c605c0c4ad0c3dd0b3f38bd203af4188c"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / f2db5d7d67ce / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-018.md#canonical-ecd89b3af405ecf7541968a7fceb933b59a697d8b94962ff85990a43286152b2)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-d2d0afcebbbd23bff82eb0a31e4303ec158f4f3de2762a3b633f63d68ffe5a90"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4733a0fe96ac218fd033554a0fd0def81e85c25dd662a31c2afe995fbc298a94"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 35d3d3f07aca / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-018.md#canonical-ecd89b3af405ecf7541968a7fceb933b59a697d8b94962ff85990a43286152b2)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates

<a id="canonical-4a9d1367e0efbb5a60e0a6c70f98c7cb9a3e760aee1d4052c026717559a6527a"></a>

Type: `"list"`. Computed.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Upstream description:

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-d27e4c09a84bd56e1bdc3875a6ebf31f1dd7b7f878c966c5a7070b80511c9aa6"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 35d3d3f07aca / 3

<a id="canonical-a62a336b9c5f2ba147441faf6ab76e41792eb404ac43b492ea45b7314cc395d3"></a>

<a id="canonical-2274263ba88eec1b4c3c5bc1c2de5596110c38a2de65dc46db55a48a390f9c5e"></a>

## certificate_url property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 35d3d3f07aca / 4

Type: `"string"`. Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

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

- [custom_hash_algorithms](data-sources--workload--reference--group-018.md#canonical-4d0ccef79811753a6aea3d1345a2dabe34dcea16e0c588559248902d3bfafbf6): complete subsection reference.

<a id="canonical-6de70b2295f74871b073c00bd6fa42fefa477a97526b376b925ff594c01d1876"></a>

<a id="canonical-09dcf7793eb5a912eb0bd726551eaf40873a357dce7bc60708c9aef6458abc80"></a>

## description_spec property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 35d3d3f07aca / 5

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--workload--reference--group-018.md#canonical-c1e6d6e3cf774e3481875974b67d545fd2d49cb2a3186797aed70559cafdb149): complete subsection reference.

- [private_key](data-sources--workload--reference--group-018.md#canonical-de0b84c7586875d5a8c76bb297d1e68e7149a61d3cc2e3272baa42887edeb1af): complete subsection reference.

- [use_system_defaults](data-sources--workload--reference--group-018.md#canonical-f41767c4669e5669ee8fb2a6b8849af3444026162c74f8a5dd6108dc499d5cd3): complete subsection reference.

<a id="canonical-0cf75238f01435146b08c9f3669dc1a3787ff45238caecb2dfcdabba44aa523c"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 35d3d3f07aca / 6

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms](data-sources--workload--reference--group-018.md#canonical-4d0ccef79811753a6aea3d1345a2dabe34dcea16e0c588559248902d3bfafbf6)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling](data-sources--workload--reference--group-018.md#canonical-c1e6d6e3cf774e3481875974b67d545fd2d49cb2a3186797aed70559cafdb149)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](data-sources--workload--reference--group-018.md#canonical-de0b84c7586875d5a8c76bb297d1e68e7149a61d3cc2e3272baa42887edeb1af)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults](data-sources--workload--reference--group-018.md#canonical-f41767c4669e5669ee8fb2a6b8849af3444026162c74f8a5dd6108dc499d5cd3)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-018.md#canonical-ecd89b3af405ecf7541968a7fceb933b59a697d8b94962ff85990a43286152b2)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-4d0ccef79811753a6aea3d1345a2dabe34dcea16e0c588559248902d3bfafbf6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7feff1a889adef4fe4542463dab612fa9970b4d0f50d0fbc0f3a0a7b4f1ab2b6"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / e8a69d656dbd / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-018.md#canonical-ecd89b3af405ecf7541968a7fceb933b59a697d8b94962ff85990a43286152b2)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-018.md#canonical-d2d0afcebbbd23bff82eb0a31e4303ec158f4f3de2762a3b633f63d68ffe5a90)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-e6243cbc20edbbbee9d6d9e1e5dcd175203a2d7cc242979619e0fa5e16d2ed08"></a>

Type: `"single"`. Computed.

Specifies the hash algorithms to be used.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-58bf151b5aa5690b584a442b89933f37052a054045c249196776e08c2f3c002b"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / e8a69d656dbd / 3

<a id="canonical-87f90b36ee68c948c2d8d8fe6b810b2b65b675aa0e4f398d88d038c62f9fe20f"></a>

<a id="canonical-1764699dffe8a01aade560f4db410a678e0ac42b35be4194e58bda317110430c"></a>

## hash_algorithms property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / e8a69d656dbd / 4

Type: `["list", "string"]`. Computed.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

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

<a id="canonical-344b815db7dfadf7dd90baf361b6a05d2e1fa25b0fb7474941513257b963c008"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / e8a69d656dbd / 5

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-018.md#canonical-d2d0afcebbbd23bff82eb0a31e4303ec158f4f3de2762a3b633f63d68ffe5a90)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-c1e6d6e3cf774e3481875974b67d545fd2d49cb2a3186797aed70559cafdb149"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-535f23c213fa2af8236308bc1060a649d659efdecad87677f63e316c8d6000bf"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / fad16f024457 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-018.md#canonical-ecd89b3af405ecf7541968a7fceb933b59a697d8b94962ff85990a43286152b2)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-018.md#canonical-d2d0afcebbbd23bff82eb0a31e4303ec158f4f3de2762a3b633f63d68ffe5a90)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-63b49c8344612e3c646c0443620704af3820d40a2ac4f6ccb6b874d9d59a3e6f"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0308db42f47b0cc0cfde2cc585454423fce51907a9db6e4501c590a1e2259918"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / fad16f024457 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9a86b4ce6d044730319397b939879d68bf4f5e9aa5ca92e3b84de0132e105a1e"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / fad16f024457 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-018.md#canonical-d2d0afcebbbd23bff82eb0a31e4303ec158f4f3de2762a3b633f63d68ffe5a90)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-de0b84c7586875d5a8c76bb297d1e68e7149a61d3cc2e3272baa42887edeb1af"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-86397a17e3144acd1df0bc82e9bc91af846d329fe2b4ed9179989835e6150cac"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / d899c561601b / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-018.md#canonical-ecd89b3af405ecf7541968a7fceb933b59a697d8b94962ff85990a43286152b2)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-018.md#canonical-d2d0afcebbbd23bff82eb0a31e4303ec158f4f3de2762a3b633f63d68ffe5a90)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key

<a id="canonical-d918cdd75a7840282177568a2a39f7302738e466bc6a309efe376d9f7e668eac"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

<a id="canonical-b04f8e16201f6f332712b415258f79c94a83be1b716becc012f36b90cb8b11e6"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / d899c561601b / 3

- [blindfold_secret_info](data-sources--workload--reference--group-018.md#canonical-93e4343a5b9bcc8a4f1ab958491c8841483e852c0984d83598503803edfbab3e): complete subsection reference.

- [clear_secret_info](data-sources--workload--reference--group-018.md#canonical-9a9031bf375f1791c75f50d3ca18672d01da2b00da3a0739522522add2748d9f): complete subsection reference.

<a id="canonical-37298440f6b43df4244bf3dcc990eba0a8e7ae7d3b678d7a7bd86c15e5d220ed"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / d899c561601b / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info](data-sources--workload--reference--group-018.md#canonical-93e4343a5b9bcc8a4f1ab958491c8841483e852c0984d83598503803edfbab3e)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info](data-sources--workload--reference--group-018.md#canonical-9a9031bf375f1791c75f50d3ca18672d01da2b00da3a0739522522add2748d9f)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-018.md#canonical-d2d0afcebbbd23bff82eb0a31e4303ec158f4f3de2762a3b633f63d68ffe5a90)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-93e4343a5b9bcc8a4f1ab958491c8841483e852c0984d83598503803edfbab3e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e4781b031c4de1bcfc00ccd01a82e2bb14848d659409f0c94694c83ef91d5cb"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 8c0ceff13953 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-018.md#canonical-ecd89b3af405ecf7541968a7fceb933b59a697d8b94962ff85990a43286152b2)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-018.md#canonical-d2d0afcebbbd23bff82eb0a31e4303ec158f4f3de2762a3b633f63d68ffe5a90)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](data-sources--workload--reference--group-018.md#canonical-de0b84c7586875d5a8c76bb297d1e68e7149a61d3cc2e3272baa42887edeb1af)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-b1917d99d281bee8c25dda085f45c54b78e53ce2a9c8dde3393d6baffb59ef17"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-d789c44741199d50155adb4f1ef413f55f3ffdfe7f3a5189ada9469c3f212ba8"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 8c0ceff13953 / 3

<a id="canonical-ed37588063310a037e12c86becb8f22b8cef3dab9f9f28cf28dfe808404b5aee"></a>

<a id="canonical-022c8e2ec09167a6817e10dc5ae4c6369d83f625c6c043cc681d29c6c83f4b33"></a>

## decryption_provider property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 8c0ceff13953 / 4

Type: `"string"`. Computed.

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

<a id="canonical-7d741b715809ecc6be614d3930b481603aa3529f60de08bf66c45e617b8e507c"></a>

<a id="canonical-6904e77aa1c1af7eff88e9dde81fdb1da982e325ab503ddaa69361f40a30438f"></a>

## location property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 8c0ceff13953 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

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

<a id="canonical-ab916e5976ab9d0b510ccf0bec5b7bd65532401ad5b4b6aaa0726723bab5265f"></a>

<a id="canonical-0642311a4d760dee4918e20148083d170cf8efc591c4f1aab35f72a25eb65baa"></a>

## store_provider property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 8c0ceff13953 / 6

Type: `"string"`. Computed.

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

<a id="canonical-1e1f7a3af2d25d940c37e34dda8f99a08d441e4ebd41dcb5dac950fde0fab2f2"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 8c0ceff13953 / 7

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](data-sources--workload--reference--group-018.md#canonical-de0b84c7586875d5a8c76bb297d1e68e7149a61d3cc2e3272baa42887edeb1af)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-9a9031bf375f1791c75f50d3ca18672d01da2b00da3a0739522522add2748d9f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5526d5d8f4c34acda5ae48adbe3d9e503c1a51758ac4854761c8078505ea5388"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 4dd01e088326 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-018.md#canonical-ecd89b3af405ecf7541968a7fceb933b59a697d8b94962ff85990a43286152b2)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-018.md#canonical-d2d0afcebbbd23bff82eb0a31e4303ec158f4f3de2762a3b633f63d68ffe5a90)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](data-sources--workload--reference--group-018.md#canonical-de0b84c7586875d5a8c76bb297d1e68e7149a61d3cc2e3272baa42887edeb1af)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-e134763f2db70e30229a98f6d70d71a99d311932af56cf73deee246e196a5d73"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-102479bec34e49929a1ad893c1e872827411bea7019e901820402c646c94ba30"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 4dd01e088326 / 3

<a id="canonical-9e9ae6f7e798f370a55402621a7393b4aab1a606abec4a856447dc25285bc36c"></a>

<a id="canonical-72a006820fa74089333d93730f27e2b8c71c2411a9a2dbc969a8464566ee718a"></a>

## provider_ref property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 4dd01e088326 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-fa0b1526443daf696816f08ece9fe022c3ff8be5f35b8adb662d633b850441a1"></a>

<a id="canonical-8c5cea7160eb8dc7abb9df7a8d8ad74f4bcc77e1b19b9e8d1b5e41062b15ee55"></a>

## url property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 4dd01e088326 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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

<a id="canonical-26e6ca5b4e3b8b37b9c252b302d5010aee04990d74ab834d79ce47d826137ee7"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 4dd01e088326 / 6

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](data-sources--workload--reference--group-018.md#canonical-de0b84c7586875d5a8c76bb297d1e68e7149a61d3cc2e3272baa42887edeb1af)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-f41767c4669e5669ee8fb2a6b8849af3444026162c74f8a5dd6108dc499d5cd3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-09eca3d26f1481c5e9c6ab66460d7eaa9f1253c06f0809e59dcba2cb18b244f6"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 92265bfcb895 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-018.md#canonical-ecd89b3af405ecf7541968a7fceb933b59a697d8b94962ff85990a43286152b2)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-018.md#canonical-d2d0afcebbbd23bff82eb0a31e4303ec158f4f3de2762a3b633f63d68ffe5a90)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-ef75030b6f3d4e4bb21a0b5d63146cebec6ab70dc5d5db056a737bd0a34797af"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-b81fbb365955303e07040f8c7099fd8a8895d3beb36dcfe8f0e6eab3d1a3adc0"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 92265bfcb895 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6cd8c2ec8dc80fc0f566edc73e992d87ab279889b6b34994e36480f7436bfdf3"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 92265bfcb895 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-018.md#canonical-d2d0afcebbbd23bff82eb0a31e4303ec158f4f3de2762a3b633f63d68ffe5a90)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-bac05e1b3b9c72c78658b5b74a11c1e8ce5989ec52946a1f8dea5591c3875544"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cdeda4d3dae85b3ac07b49854b15642c9ea6d8bcf389ed79329a0d3272fec2d7"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 980e5ddf8f2d / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-018.md#canonical-ecd89b3af405ecf7541968a7fceb933b59a697d8b94962ff85990a43286152b2)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config

<a id="canonical-e3e2c775fb101273df85d17abe207ac318fd55ed40be4d487bd50490f60fa0b8"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

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

<a id="canonical-94fce5ecd4b12ab1a605521557434d2df2b43fdddb7913b783b1baf7446bb89b"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 980e5ddf8f2d / 3

- [custom_security](data-sources--workload--reference--group-018.md#canonical-1f82140047d9550ddc4e3b922a267322135cf45e3324f755a0cae7be32ec861b): complete subsection reference.

- [default_security](data-sources--workload--reference--group-018.md#canonical-f448c44de5c716d697b6ada9440e2bad79b96d74fe1022624a7bb05956a57e9f): complete subsection reference.

- [low_security](data-sources--workload--reference--group-018.md#canonical-7930d56097d58544b696b8ed9d5a87edb5a7ec1d20cfdea4bb3cd983518ad92b): complete subsection reference.

- [medium_security](data-sources--workload--reference--group-018.md#canonical-5f8682403eedb08da43dbb205cf2ae805fa0aada45944861558440242339bc41): complete subsection reference.

<a id="canonical-0d3dd081beed9417a94a7a72ec2707b077ebad0ab6834346963b68edcc6519da"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 980e5ddf8f2d / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security](data-sources--workload--reference--group-018.md#canonical-1f82140047d9550ddc4e3b922a267322135cf45e3324f755a0cae7be32ec861b)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security](data-sources--workload--reference--group-018.md#canonical-f448c44de5c716d697b6ada9440e2bad79b96d74fe1022624a7bb05956a57e9f)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security](data-sources--workload--reference--group-018.md#canonical-7930d56097d58544b696b8ed9d5a87edb5a7ec1d20cfdea4bb3cd983518ad92b)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security](data-sources--workload--reference--group-018.md#canonical-5f8682403eedb08da43dbb205cf2ae805fa0aada45944861558440242339bc41)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-018.md#canonical-ecd89b3af405ecf7541968a7fceb933b59a697d8b94962ff85990a43286152b2)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-1f82140047d9550ddc4e3b922a267322135cf45e3324f755a0cae7be32ec861b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a0ebb9a97b09ec06248e4b86af2c9ab955862b2d3eb453c4b3b733d016b0a45"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / b620e7e80ec1 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-018.md#canonical-ecd89b3af405ecf7541968a7fceb933b59a697d8b94962ff85990a43286152b2)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-018.md#canonical-bac05e1b3b9c72c78658b5b74a11c1e8ce5989ec52946a1f8dea5591c3875544)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security

<a id="canonical-1ad900860a1fcbf5342f21cf28d10f5c40b3b9b49c38675576e7c3af42f615fd"></a>

Type: `"single"`. Computed.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-5222c553bebb77b9b135cd28fba915cc20d5597b334654c8326c1d40bd54abfa"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / b620e7e80ec1 / 3

<a id="canonical-0ae135c98276ba4f7951e0ddb6721bc05751dd4cf7b6c9287a4beab88a98a540"></a>

<a id="canonical-873331a39982350b3ac4d56973d8742ad3f9bbd68b64508cb92682a481dd4110"></a>

## cipher_suites property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / b620e7e80ec1 / 4

Type: `["list", "string"]`. Computed.

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

<a id="canonical-c5e55ab2a03c53487759ce1f63db02e1a37097568d917de104648e00fd844520"></a>

<a id="canonical-3881c5af3cfdd29b9866e0338a7972bfecaf5ac46011729d64c03aa32bc1f714"></a>

## max_version property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / b620e7e80ec1 / 5

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

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

<a id="canonical-322ac81ee26dc0d632d0b6a1c2861d46616c76015080b6cbc6c79c222a98dd9c"></a>

<a id="canonical-a486bc45e8bfed537abbe63ca6c262840af0a9cc8d5ce772198536325d2c24bc"></a>

## min_version property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / b620e7e80ec1 / 6

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

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

<a id="canonical-1ed550d6cd1771c42849c02a718f08e7f31dfb47e23b6b4fcc1197f75f1cc83b"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / b620e7e80ec1 / 7

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-018.md#canonical-bac05e1b3b9c72c78658b5b74a11c1e8ce5989ec52946a1f8dea5591c3875544)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-f448c44de5c716d697b6ada9440e2bad79b96d74fe1022624a7bb05956a57e9f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ded1c96a6ef9ce5e84546af540c010dd04891af8663a1a74d20e4de7562fe768"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / ea780b551b39 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-018.md#canonical-ecd89b3af405ecf7541968a7fceb933b59a697d8b94962ff85990a43286152b2)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-018.md#canonical-bac05e1b3b9c72c78658b5b74a11c1e8ce5989ec52946a1f8dea5591c3875544)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security

<a id="canonical-9cc2f5915f83f6019bf06a12ac330307daf15bdc2c10e8f3656462ea1eff9198"></a>

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

<a id="canonical-91ee4656fb6632741eafd91cbf8c6b7e3dbf91493db191b50134f09da89b5798"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / ea780b551b39 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9e3f72d8bb78a9e1adde917e1dd68dd5cfe090b94bc215a179edcdac802f9b43"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / ea780b551b39 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-018.md#canonical-bac05e1b3b9c72c78658b5b74a11c1e8ce5989ec52946a1f8dea5591c3875544)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-7930d56097d58544b696b8ed9d5a87edb5a7ec1d20cfdea4bb3cd983518ad92b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a27f035959237ac8aa86f7fe64bc85e0b08c28abde7079f577fabe84e445de42"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / b7461909e0de / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-018.md#canonical-ecd89b3af405ecf7541968a7fceb933b59a697d8b94962ff85990a43286152b2)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-018.md#canonical-bac05e1b3b9c72c78658b5b74a11c1e8ce5989ec52946a1f8dea5591c3875544)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security

<a id="canonical-4a22209ec8128e6dceba1a81ab620fb7f9dbab4a541666a1ea9e6aac1d2c4327"></a>

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

<a id="canonical-40547b94647bc56a27966d99ccb0359c4e3880173d23d114b13e70ed149a1a96"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / b7461909e0de / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-53c782e6a2726ac244b67a9195b0eb37f22170e643a97e3ceede1aab399487f8"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / b7461909e0de / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-018.md#canonical-bac05e1b3b9c72c78658b5b74a11c1e8ce5989ec52946a1f8dea5591c3875544)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-5f8682403eedb08da43dbb205cf2ae805fa0aada45944861558440242339bc41"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-93eb2353e3137b4ca634b62f132edfead94de54e7f039699d6ecdf841678369b"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / b066ee92cbbe / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-018.md#canonical-ecd89b3af405ecf7541968a7fceb933b59a697d8b94962ff85990a43286152b2)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-018.md#canonical-bac05e1b3b9c72c78658b5b74a11c1e8ce5989ec52946a1f8dea5591c3875544)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security

<a id="canonical-701f156d77c8df2b1c28439ef0d8897abbcabde77e822f35078dcea32cc097e1"></a>

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

<a id="canonical-cf0fcb1e104bc33f5705770d202981789e8cb8e466e9aae77be4394b6ece8ece"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / b066ee92cbbe / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-87869ba66149b0c9988e03dec60a8cade8f2914dc1285601e68d0b516b132320"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / b066ee92cbbe / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-018.md#canonical-bac05e1b3b9c72c78658b5b74a11c1e8ce5989ec52946a1f8dea5591c3875544)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-c431c7272fd63a6532c467f52bccbb1eff240da38d1148190cbd6c4484a2e7c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da2e866aa47801c778b63c7f2aca2bb548561026e31e0025f98c923e538be9a6"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 63b053682e6e / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-018.md#canonical-ecd89b3af405ecf7541968a7fceb933b59a697d8b94962ff85990a43286152b2)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls

<a id="canonical-d63a78b7e6e97b82a262c11630006d43b73c508d5664421da396508170cdcb7b"></a>

Type: `"single"`. Computed.

Validation context for downstream client TLS connections.

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

<a id="canonical-cc41656d9ee66602fb501e49745581fadf69339e0def6d903a797717b90658c3"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 63b053682e6e / 3

<a id="canonical-9ca654685b140ef15ed2a86389ddabd3d908f4d52266acc80da7c0db02bcb3f2"></a>

<a id="canonical-d596b10ce8ee5e11e81f8bb1fb5b29eb6ba2280cf853f9e2e8a7bb438e1455f5"></a>

## client_certificate_optional property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 63b053682e6e / 4

Type: `"bool"`. Computed.

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

- [crl](data-sources--workload--reference--group-018.md#canonical-d1f775006c7de19eb569318e47420866f0a4e328fffbe1d3a3a6b3a01c710fd7): complete subsection reference.

- [no_crl](data-sources--workload--reference--group-018.md#canonical-74319b28eafb7b7b36027d19fcef7912916a4f6f0885ef2cdd1894831a5ed076): complete subsection reference.

- [trusted_ca](data-sources--workload--reference--group-018.md#canonical-ac61e2411675dbfd07e4a39525b7277bed492795cbd22f1874cc0778e15df62b): complete subsection reference.

<a id="canonical-6b3521ad1fd8a8461851bde566c2ad8c515d8887bd3091379b6dc39cfda509df"></a>

<a id="canonical-cb04aeb8ed4821561e3ceb3e240d3f5916bac43a94c4ea1e4111febceeacb1d9"></a>

## trusted_ca_url property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 63b053682e6e / 5

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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

- [xfcc_disabled](data-sources--workload--reference--group-018.md#canonical-4606683834412143ea79c4e5e536cf93f628f5dd07a52ab226c89a77bd50a046): complete subsection reference.

- [xfcc_options](data-sources--workload--reference--group-018.md#canonical-2e992be2ae82904c1160714e55959b59f7f92e9b6fc54da49d4d0dff678c95c1): complete subsection reference.

<a id="canonical-e3528ba8b7d84584370ed7adfdf8523504b77e8174119e2d0d3f6055d3957711"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 63b053682e6e / 6

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl](data-sources--workload--reference--group-018.md#canonical-d1f775006c7de19eb569318e47420866f0a4e328fffbe1d3a3a6b3a01c710fd7)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl](data-sources--workload--reference--group-018.md#canonical-74319b28eafb7b7b36027d19fcef7912916a4f6f0885ef2cdd1894831a5ed076)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca](data-sources--workload--reference--group-018.md#canonical-ac61e2411675dbfd07e4a39525b7277bed492795cbd22f1874cc0778e15df62b)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled](data-sources--workload--reference--group-018.md#canonical-4606683834412143ea79c4e5e536cf93f628f5dd07a52ab226c89a77bd50a046)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options](data-sources--workload--reference--group-018.md#canonical-2e992be2ae82904c1160714e55959b59f7f92e9b6fc54da49d4d0dff678c95c1)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-018.md#canonical-ecd89b3af405ecf7541968a7fceb933b59a697d8b94962ff85990a43286152b2)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-d1f775006c7de19eb569318e47420866f0a4e328fffbe1d3a3a6b3a01c710fd7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ff9f98957e713984302f8b83fa4ef544887e1d95b49e688f2222fca63d3d9e2"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 4e1ad6be85db / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-018.md#canonical-ecd89b3af405ecf7541968a7fceb933b59a697d8b94962ff85990a43286152b2)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-018.md#canonical-c431c7272fd63a6532c467f52bccbb1eff240da38d1148190cbd6c4484a2e7c1)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl

<a id="canonical-8d75f2a1ba5695cdd69833f93ee7c2cf639123716e10451aecf9d50fb24802d1"></a>

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

<a id="canonical-ec03cfa937442382d867c3bc074fc39063607afd7b73cfceff87c4970ffe733a"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 4e1ad6be85db / 3

<a id="canonical-723a973c25a9c570b6dd8124ebc334eecb8fad6af4ef682feed79efa1ce70179"></a>

<a id="canonical-735ad1916fb14f5123f2007529a7c98ecd4b8591d61b5c83332bd7b338d44c5f"></a>

## name property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 4e1ad6be85db / 4

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

<a id="canonical-fa5ca3d313a6d48a5d94c04916ac928a7a32b27fbe19007912f0358e11365ca8"></a>

<a id="canonical-1feb08c6c6efedaa8a7bfc6a7d536901222c551f8f253e0c6f859964b11f2992"></a>

## namespace property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 4e1ad6be85db / 5

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

<a id="canonical-a4704f3a9983391603a7a966bf7905decff3d85dcd04dacd3d27d3e81c754cfe"></a>

<a id="canonical-01578baad0a752cc5268c27f55830ebeb6a076499fdabe6fcbd476c5dae8a380"></a>

## tenant property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 4e1ad6be85db / 6

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

<a id="canonical-ff36a3e7a6056c5ac35aeae760d6d8b6e59cb6f7134a8327f1cf8858fcc2595d"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 4e1ad6be85db / 7

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-018.md#canonical-c431c7272fd63a6532c467f52bccbb1eff240da38d1148190cbd6c4484a2e7c1)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-74319b28eafb7b7b36027d19fcef7912916a4f6f0885ef2cdd1894831a5ed076"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-47e0271a28834ae129f63a542c73b89642919623f7c991fe500a4b4e87793674"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / c14aff2ef522 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-018.md#canonical-ecd89b3af405ecf7541968a7fceb933b59a697d8b94962ff85990a43286152b2)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-018.md#canonical-c431c7272fd63a6532c467f52bccbb1eff240da38d1148190cbd6c4484a2e7c1)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl

<a id="canonical-953fc2604aa493455a445605d9664ac9b18e7924ab9cbd02de09dac50221e690"></a>

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

<a id="canonical-f1141f88274745b5bcc60ac3c23e4100216bb369862eeb5df8c6142a483746e1"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / c14aff2ef522 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-477b9e70c134be4d9492a3b9e2a77f0285bd2bce2f69e7473b861d595362464f"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / c14aff2ef522 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-018.md#canonical-c431c7272fd63a6532c467f52bccbb1eff240da38d1148190cbd6c4484a2e7c1)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-ac61e2411675dbfd07e4a39525b7277bed492795cbd22f1874cc0778e15df62b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-33b907ea9cf3784fa1ad33412f8e0c41b62aba474a747cc50bc8f5af9d281223"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 8bc5dc1b3977 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-018.md#canonical-ecd89b3af405ecf7541968a7fceb933b59a697d8b94962ff85990a43286152b2)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-018.md#canonical-c431c7272fd63a6532c467f52bccbb1eff240da38d1148190cbd6c4484a2e7c1)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-fac3407053493d5aa57820c5b5961273523d8149440311ccceb5b1495ed15ba3"></a>

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

<a id="canonical-3de4bddab2908f29a6c1593db4fee78bfc668b46427c6d77151ebc19a050047a"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 8bc5dc1b3977 / 3

<a id="canonical-1f02405e8fb2393b031711025cefcfab26238f6759d8f5e0599cf78091a364a7"></a>

<a id="canonical-0e260fcbab674bdcad3b903c16a963eea77851649cc98d343b66fbf4d581cf2a"></a>

## name property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 8bc5dc1b3977 / 4

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

<a id="canonical-f3cdb53a196683d72d211a97efb1316d155f5ba30ef5a33846c9fe3d3e3daffb"></a>

<a id="canonical-6009baf7d22df200217a0f71710a3365409e42fd202661df239052621749372e"></a>

## namespace property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 8bc5dc1b3977 / 5

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

<a id="canonical-1dfcd5460a2d2c90bb456bd08051f7878704d5b8c224eb2f5f9049fdd0fa471a"></a>

<a id="canonical-879f5ce31b1d694a7d44405ee7f4f50a2f949c616cd8a5abe13dc907e2817a7d"></a>

## tenant property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 8bc5dc1b3977 / 6

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

<a id="canonical-cbceb51c7dbdd4c41e81fc36a93c2d84b06123651fed9162f04c5b3ebb9d23b8"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 8bc5dc1b3977 / 7

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-018.md#canonical-c431c7272fd63a6532c467f52bccbb1eff240da38d1148190cbd6c4484a2e7c1)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-4606683834412143ea79c4e5e536cf93f628f5dd07a52ab226c89a77bd50a046"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a6197dee9fc2969aa9b22930e80d08277c55674142f9e77eab755b76c6c0b5dc"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 3fd0791b8de0 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-018.md#canonical-ecd89b3af405ecf7541968a7fceb933b59a697d8b94962ff85990a43286152b2)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-018.md#canonical-c431c7272fd63a6532c467f52bccbb1eff240da38d1148190cbd6c4484a2e7c1)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-108b7b0436a9ad441f957b577caa5b242ef4104eb6f94935961f9b6b3eb11157"></a>

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

<a id="canonical-04aa1f5dbb08f8db005ad643ac940d882098dcfc6f8d2e42c87adbd29a30ea83"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 3fd0791b8de0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fa9e540cb58fed39bd7be47397e1cacfda3a8c971eef899d80fb13b2a7955acd"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 3fd0791b8de0 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-018.md#canonical-c431c7272fd63a6532c467f52bccbb1eff240da38d1148190cbd6c4484a2e7c1)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-2e992be2ae82904c1160714e55959b59f7f92e9b6fc54da49d4d0dff678c95c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9f388089676b0f6e97fdc0733ba55f0c27631f4952448bdadc95aacd439b0ca1"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 3566c9265ea6 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-018.md#canonical-ecd89b3af405ecf7541968a7fceb933b59a697d8b94962ff85990a43286152b2)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-018.md#canonical-c431c7272fd63a6532c467f52bccbb1eff240da38d1148190cbd6c4484a2e7c1)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-abac0b721f513ac0fa2a45032f39f0fa8816564221f4ca584c5a19241a0b6f83"></a>

Type: `"single"`. Computed.

X-Forwarded-Client-Cert header elements to be added to requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-94ebdfc6e1c1a37137c535dd182c92866a842c68f586f4b7b23d4f69dc435918"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 3566c9265ea6 / 3

<a id="canonical-2d9449fecf6286b5ee54a941e5c724aa68da7f241b052d6432c648edc53d6c0f"></a>

<a id="canonical-38c0a92ad50bfe0be7eef49cf44f28257593b84ce81cd4476a411892d9cafcbf"></a>

## xfcc_header_elements property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 3566c9265ea6 / 4

Type: `["list", "string"]`. Computed.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be added to requests.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-6e0defc065d11def42e6b0a471862f56cb0594e22c04f02b2f0136cabb5cff3a"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 3566c9265ea6 / 5

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-018.md#canonical-c431c7272fd63a6532c467f52bccbb1eff240da38d1148190cbd6c4484a2e7c1)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-aad260f9f514def74f6d2d0b5de8b2a3348f266b26c8f24655dde803b35aec50"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af8e058dd538254690e75b919dfa5df9d65e1fc54dbd0bfd29447f808f566386"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 7ebbeef86b5d / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert

<a id="canonical-7b14a7b317b551de85c658ced144247805ae67683440c821338b844d8caa81b2"></a>

Type: `"single"`. Computed.

Choice for selecting HTTP proxy with bring your own certificates.

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
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]"
}
```

<a id="canonical-6383395e276657ca202101f955e4e5006e545f3a872632bbd93960a3aec14604"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 7ebbeef86b5d / 3

<a id="canonical-02dbd6420668aedba77cec7dd435fabbdd18d450175c42623222c360caef4b45"></a>

<a id="canonical-dd655244a820b639e970697fb30c645d8212bba7d788a4b4cdfffbf30a6499d5"></a>

## add_hsts property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 7ebbeef86b5d / 4

Type: `"bool"`. Computed.

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

<a id="canonical-83543f22f3878697bb831a048835568e02cb59fbef99a87b76c2a25ee9df5f26"></a>

<a id="canonical-10e0c32aa94c1cec156ef038d4b79e07848a85e18f143446de04b90474966d41"></a>

## append_server_name property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 7ebbeef86b5d / 5

Type: `"string"`. Computed.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Upstream description:

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

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

- [coalescing_options](data-sources--workload--reference--group-018.md#canonical-3a1bf29a107b683535cea719caac513a5331eef1b61c56a6554b439d1b798df1): complete subsection reference.

<a id="canonical-c9559486e60b5108d5e6cd55a00130f32087e9e36bde80dad816914463cc3afe"></a>

<a id="canonical-8ee523c383de91e9e5abfd05eb1e3e26301325594430f75c19c9b36c12eef628"></a>

## connection_idle_timeout property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 7ebbeef86b5d / 6

Type: `"number"`. Computed.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed.

Upstream description:

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

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

- [default_header](data-sources--workload--reference--group-018.md#canonical-a071e2ad826f053483f38087efcb921dd7f0a521b04fd44e500d2227c3aa9714): complete subsection reference.

- [default_loadbalancer](data-sources--workload--reference--group-018.md#canonical-4d87e69446ab8d1a64bb9be550943f0253dbb13826bb2b50d683db45d5e7dcb9): complete subsection reference.

- [disable_path_normalize](data-sources--workload--reference--group-018.md#canonical-3fb9e38cdf7ff729ff7190eff71f479570c513916a56c15320dd812a80213b80): complete subsection reference.

- [enable_path_normalize](data-sources--workload--reference--group-018.md#canonical-f437cff2b00e62f7c2edec25e2541c11acbce8d4b798d89248af96e0b377f083): complete subsection reference.

- [http_protocol_options](data-sources--workload--reference--group-018.md#canonical-0ad1c98e3b6bd151fe9b2e21d56202a35e2f8c11699f56394b86077ae7f7bc27): complete subsection reference.

<a id="canonical-e1b0222d7523dbd38d6f21e979142fe9eef3a55af0a670a2aff09bbe62db1968"></a>

<a id="canonical-b73bb6686a440037396e5200ac83a13fb47666ab4da04f67d3ba884328de473d"></a>

## http_redirect property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 7ebbeef86b5d / 7

Type: `"bool"`. Computed.

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

- [no_mtls](data-sources--workload--reference--group-019.md#canonical-7146c3d6c8f2915934b276f9dabb3dc937cea4e99df3df62a8e1430fdc6c1a2d): complete subsection reference.

- [non_default_loadbalancer](data-sources--workload--reference--group-019.md#canonical-1b995465e9ee0d70e8bbe26182d2abdcf9f1e091338d723e6348230d4526992c): complete subsection reference.

- [pass_through](data-sources--workload--reference--group-019.md#canonical-51d5767ad0edfb1d393c12b757306d939415d321c7bcd3c270c8ca3d7bc898c3): complete subsection reference.

<a id="canonical-8ee517cb138dc5b27a3991bd7611a84ca5e57929210b7d29b619f931ac6cd86e"></a>

<a id="canonical-7e81fcdef02a898d5557e32f7949138a5b8e7f477e93b292558e546477cad74f"></a>

## port property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 7ebbeef86b5d / 8

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTPS port to Listen.

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

<a id="canonical-36b8b6b6b813a4b2651e29cd275e7a380bbc4f23a912ff46c3ca4d70976e5526"></a>

<a id="canonical-a4889d8484a33393944087c767159f87627867482e8ff588067fc5bd9193e1d8"></a>

## port_ranges property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 7ebbeef86b5d / 9

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

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

<a id="canonical-b5fdc430a45c7c20059c11f4be4373bc64e13a1a0938ff1dce5c9e38e1b246ed"></a>

<a id="canonical-a36cd7f8a00e29b5b2f548324783498bca35b4cc877f0c24c75dc273feadc9ec"></a>

## server_name property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 7ebbeef86b5d / 10

Type: `"string"`. Computed.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

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

- [tls_config](data-sources--workload--reference--group-019.md#canonical-4c159d49a0e4d99d500e8eeee4c1f24b553baa1365aab503dc79ad5941fc8d5c): complete subsection reference.

- [use_mtls](data-sources--workload--reference--group-019.md#canonical-1f098cee67b1b95343678632fe5d8686c3f5294ed2424ef5e72a4212f27d3034): complete subsection reference.

<a id="canonical-226fc18745fef4ab0f284e4fa7d80f1800915bd1fa6b5b844072e159d1b51730"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 7ebbeef86b5d / 11

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-018.md#canonical-3a1bf29a107b683535cea719caac513a5331eef1b61c56a6554b439d1b798df1)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_header](data-sources--workload--reference--group-018.md#canonical-a071e2ad826f053483f38087efcb921dd7f0a521b04fd44e500d2227c3aa9714)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_loadbalancer](data-sources--workload--reference--group-018.md#canonical-4d87e69446ab8d1a64bb9be550943f0253dbb13826bb2b50d683db45d5e7dcb9)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.disable_path_normalize](data-sources--workload--reference--group-018.md#canonical-3fb9e38cdf7ff729ff7190eff71f479570c513916a56c15320dd812a80213b80)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.enable_path_normalize](data-sources--workload--reference--group-018.md#canonical-f437cff2b00e62f7c2edec25e2541c11acbce8d4b798d89248af96e0b377f083)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-018.md#canonical-0ad1c98e3b6bd151fe9b2e21d56202a35e2f8c11699f56394b86077ae7f7bc27)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.no_mtls](data-sources--workload--reference--group-019.md#canonical-7146c3d6c8f2915934b276f9dabb3dc937cea4e99df3df62a8e1430fdc6c1a2d)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer](data-sources--workload--reference--group-019.md#canonical-1b995465e9ee0d70e8bbe26182d2abdcf9f1e091338d723e6348230d4526992c)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.pass_through](data-sources--workload--reference--group-019.md#canonical-51d5767ad0edfb1d393c12b757306d939415d321c7bcd3c270c8ca3d7bc898c3)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-019.md#canonical-4c159d49a0e4d99d500e8eeee4c1f24b553baa1365aab503dc79ad5941fc8d5c)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-019.md#canonical-1f098cee67b1b95343678632fe5d8686c3f5294ed2424ef5e72a4212f27d3034)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-3a1bf29a107b683535cea719caac513a5331eef1b61c56a6554b439d1b798df1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-775deb5c9367bbd9b5ebfa23a27ffcb7daa53fd21ad0b174009127d7bf0b4007"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / e8ae99a3f761 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-018.md#canonical-aad260f9f514def74f6d2d0b5de8b2a3348f266b26c8f24655dde803b35aec50)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options

<a id="canonical-82d21b2d58aa6ded9fe5c771ad412f91f6f2bba3b2790f464653e1f1d745a53b"></a>

Type: `"single"`. Computed.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

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

<a id="canonical-111f6b201b1632454d0aa4395dec3f38fdc9d5a89269be858080e64d75578225"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / e8ae99a3f761 / 3

- [default_coalescing](data-sources--workload--reference--group-018.md#canonical-e389130e6162894393ad8303e151d7d764bc23d18769657d024dc8976410b81c): complete subsection reference.

- [strict_coalescing](data-sources--workload--reference--group-018.md#canonical-b0318670e56eec66f4377f4e2406389aad038505130ec0822f6ff3b9b864a4f4): complete subsection reference.

<a id="canonical-cb5bb2772a372fec12856981ae6754f7208f051f48c6bea38abb6cb0232c9d79"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / e8ae99a3f761 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing](data-sources--workload--reference--group-018.md#canonical-e389130e6162894393ad8303e151d7d764bc23d18769657d024dc8976410b81c)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing](data-sources--workload--reference--group-018.md#canonical-b0318670e56eec66f4377f4e2406389aad038505130ec0822f6ff3b9b864a4f4)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-018.md#canonical-aad260f9f514def74f6d2d0b5de8b2a3348f266b26c8f24655dde803b35aec50)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-e389130e6162894393ad8303e151d7d764bc23d18769657d024dc8976410b81c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-67234c30eb4f94714ddcd2a3d233bc79f54c7407d8548ad598bca8fcdf5e325f"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 038dad64f1ba / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-018.md#canonical-aad260f9f514def74f6d2d0b5de8b2a3348f266b26c8f24655dde803b35aec50)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-018.md#canonical-3a1bf29a107b683535cea719caac513a5331eef1b61c56a6554b439d1b798df1)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing

<a id="canonical-c3cf0187c011dd79a72ce45aa807dee3bd4a985603520837d0f2609aedae1e91"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-9742c6f0ff2c7ced71a37172272e946f79873eb8059251d8814379c18a500245"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 038dad64f1ba / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8173b14228220b2e63172fe445e6e532c13694f738eefa962317227a9f2dded1"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 038dad64f1ba / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-018.md#canonical-3a1bf29a107b683535cea719caac513a5331eef1b61c56a6554b439d1b798df1)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-b0318670e56eec66f4377f4e2406389aad038505130ec0822f6ff3b9b864a4f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e6950e4d5bea63c58b83713be84e504609ef0c96b9ab054c1088d03f0ca2a95c"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 02ca35d0f1f7 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-018.md#canonical-aad260f9f514def74f6d2d0b5de8b2a3348f266b26c8f24655dde803b35aec50)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-018.md#canonical-3a1bf29a107b683535cea719caac513a5331eef1b61c56a6554b439d1b798df1)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing

<a id="canonical-ed3b031e761f60add857bcd70e2d68bbdaa05ccb9c6514d2f3a8b45d33f2a379"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-e19f5583bec7df04f78f04bd205a8c252018500bcb8ded7dc3538fccb3ae9c8e"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 02ca35d0f1f7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a0230edc3ba66193b54019149f57a7bd38af8d7ac56368a15a3bffd113a4a18c"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 02ca35d0f1f7 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-018.md#canonical-3a1bf29a107b683535cea719caac513a5331eef1b61c56a6554b439d1b798df1)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-a071e2ad826f053483f38087efcb921dd7f0a521b04fd44e500d2227c3aa9714"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fe63e6cf5d126db5219d551cf9445e5f2c7ae347242563e74243c4a3bb826d23"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_header — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 851ef2eedef6 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-018.md#canonical-aad260f9f514def74f6d2d0b5de8b2a3348f266b26c8f24655dde803b35aec50)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_header

<a id="canonical-4e5799e5232198ac5a88c46cddc3ae4a108f516670a019e983b4873188ec69bf"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-c869309245925250590f4002b878480241b576a7afc4fa217a80c632be45b8cf"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 851ef2eedef6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-92bee01946a1d27572f34729e212d5ca90cec92a11c37549e43d60a77a64a0ae"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 851ef2eedef6 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-018.md#canonical-aad260f9f514def74f6d2d0b5de8b2a3348f266b26c8f24655dde803b35aec50)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-4d87e69446ab8d1a64bb9be550943f0253dbb13826bb2b50d683db45d5e7dcb9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-372562dc6f0d14ef313dc2c33b93f9f491fa92f1bd4a03d0bae0f26a64b1cfa8"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_loadbalancer — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / f5d7b3f12d6c / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-018.md#canonical-aad260f9f514def74f6d2d0b5de8b2a3348f266b26c8f24655dde803b35aec50)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_loadbalancer

<a id="canonical-bb606202c131e789ea2b65c1fb88e359c2e21827289c883853a0e13e63b0e931"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-5f0f09495e117408391cdf0aae7f0fdd370addf4e2d463012f7124a4960aef5e"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / f5d7b3f12d6c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-28d98fcdc2daa303c917b607735d809e8242e1899b3be3ef0818b1230fb67079"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / f5d7b3f12d6c / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-018.md#canonical-aad260f9f514def74f6d2d0b5de8b2a3348f266b26c8f24655dde803b35aec50)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-3fb9e38cdf7ff729ff7190eff71f479570c513916a56c15320dd812a80213b80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a7928ea8c3a9bcacc48937555cdbe1442450293c765f0ef568af1c6d5ccd52c"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.disable_path_normalize — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 4578a95cac62 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-018.md#canonical-aad260f9f514def74f6d2d0b5de8b2a3348f266b26c8f24655dde803b35aec50)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.disable_path_normalize

<a id="canonical-1f599db243fdc7c1e7fffe296e0693d7abd0c0184fbfa2c5736188c604289c0c"></a>

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

<a id="canonical-fac56d8c38d872c40638df338d72471cc0f7a940acf9a181972b109a93d73519"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 4578a95cac62 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-48e31fc444d6148cc8c4417a6f0dc12a0dbfb3f2c6757da5d653732453ba2c02"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 4578a95cac62 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-018.md#canonical-aad260f9f514def74f6d2d0b5de8b2a3348f266b26c8f24655dde803b35aec50)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-f437cff2b00e62f7c2edec25e2541c11acbce8d4b798d89248af96e0b377f083"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f00e41bee1f391e5bf63e2e3695d708745dca51366c1e48abd357c4851d24e4b"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.enable_path_normalize — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 2e19d608dcce / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-018.md#canonical-aad260f9f514def74f6d2d0b5de8b2a3348f266b26c8f24655dde803b35aec50)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.enable_path_normalize

<a id="canonical-1e5c7a4c55d330fcdd868b8d39e8f6907b9804ccb0d4a4bd103ed18a86cae176"></a>

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

<a id="canonical-dcc91738651b6159962f6f6b830ea29bd6ab72f71d64bcb9a9fa968ac313ab73"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 2e19d608dcce / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-12a0346d7340402ee1d7cde4beae44f58a317837a8f3923081a16c0e993ff8ed"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 2e19d608dcce / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-018.md#canonical-aad260f9f514def74f6d2d0b5de8b2a3348f266b26c8f24655dde803b35aec50)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-0ad1c98e3b6bd151fe9b2e21d56202a35e2f8c11699f56394b86077ae7f7bc27"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-61a0beb39e471e44a184758b6cc8ffaf2588c967a4be1b13b8914970794f7c30"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / e8a78b9d53cc / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-018.md#canonical-aad260f9f514def74f6d2d0b5de8b2a3348f266b26c8f24655dde803b35aec50)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options

<a id="canonical-4a1e4ecd2f357cf8b4dac0b93cbf6a71cb63754dd3cd3764be5ebcca21c76dcb"></a>

Type: `"single"`. Computed.

HTTP protocol configuration OPTIONS for downstream connections.

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

<a id="canonical-2091d31b2eca232f9e49ff0d0ecd7e5cfa66647c4cf53cf933c698ffd565e49e"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / e8a78b9d53cc / 3

- [http_protocol_enable_v1_only](data-sources--workload--reference--group-018.md#canonical-7d0550f9d0fe4e18e50153304cf483baa15202305f3051ca27edcf661e24c043): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--workload--reference--group-019.md#canonical-e4a2959f413d001b85491d2ce00efbd0dd7a4623fe445f8d29ce50f35ccb9871): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--workload--reference--group-019.md#canonical-87c1d84c5e2c5429c1b95adc5cca41ed56cc677d1a70cad2b0240f5c4720e30f): complete subsection reference.

<a id="canonical-31b090f21f62f29c350399f4e0a2a6a0b7740a66a23a4e7b8080c09b94aa630d"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / e8a78b9d53cc / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-018.md#canonical-7d0550f9d0fe4e18e50153304cf483baa15202305f3051ca27edcf661e24c043)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2](data-sources--workload--reference--group-019.md#canonical-e4a2959f413d001b85491d2ce00efbd0dd7a4623fe445f8d29ce50f35ccb9871)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only](data-sources--workload--reference--group-019.md#canonical-87c1d84c5e2c5429c1b95adc5cca41ed56cc677d1a70cad2b0240f5c4720e30f)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-018.md#canonical-aad260f9f514def74f6d2d0b5de8b2a3348f266b26c8f24655dde803b35aec50)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-7d0550f9d0fe4e18e50153304cf483baa15202305f3051ca27edcf661e24c043"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df28762000c429d852599feffca0daf714c7dad500a534d9640780aacc912cc4"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / e68afbddbba6 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-018.md#canonical-aad260f9f514def74f6d2d0b5de8b2a3348f266b26c8f24655dde803b35aec50)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-018.md#canonical-0ad1c98e3b6bd151fe9b2e21d56202a35e2f8c11699f56394b86077ae7f7bc27)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-12acbd7854c6650c5aec5897bb976d4409226a567962c770ec2e2194a32c96c4"></a>

Type: `"single"`. Computed.

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

<a id="canonical-86e90d854d23d593fd4fec9a2ad092af239ba6873f7e0f4d5299dbf670778392"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / e68afbddbba6 / 3

- [header_transformation](data-sources--workload--reference--group-019.md#canonical-b5df8fa031971f2df8ff6a335149ecfd1e0f59060050e8c6b542da441c819066): complete subsection reference.
