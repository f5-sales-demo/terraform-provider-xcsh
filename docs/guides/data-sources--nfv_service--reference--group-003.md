---
page_title: "xcsh_nfv_service reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nfv_service reference."
---

# xcsh_nfv_service reference

<a id="canonical-2effb91f431d4afa35813aa88d97909654a18fff4e805ef00df0aa4a97f7e92b"></a>

## certificate_url property — https_management.advertise_on_slo_sli.tls_certificates / 4eb18c169d46 / 4

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

- [custom_hash_algorithms](data-sources--nfv_service--reference--group-003.md#canonical-aa42abb4870cdf7e379cf2055fb2a9b3bfae14d7879b645fe5d357ca3da31aed): complete subsection reference.

<a id="canonical-eff64a6461c9734b85a36752055e5ec8060aa574c95f544a70c2d617d62629db"></a>

<a id="canonical-d4f305c66940a91e5dcd76bf9c19ce35aecd8c9a4b568fdb6b4a3e32bb253ecf"></a>

## description_spec property — https_management.advertise_on_slo_sli.tls_certificates / 4eb18c169d46 / 5

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--nfv_service--reference--group-003.md#canonical-663a5127705c498b65ed12e5ad6d899733011a16295cde9e3390b51cf35cd532): complete subsection reference.

- [private_key](data-sources--nfv_service--reference--group-003.md#canonical-be46b1c4c1902baae6b4e735cc4a4fa24885a02b5c76d6df7fad39bb4b160f69): complete subsection reference.

- [use_system_defaults](data-sources--nfv_service--reference--group-003.md#canonical-1f46448736c4a2029744c416889864ec08c2c3add10b061e7f2306b4c8c6e221): complete subsection reference.

<a id="canonical-f4852e0e956e0ba75a066af14673c3012b2e77df24d05e21b07bbbe9e07fd94f"></a>

## Next pages — https_management.advertise_on_slo_sli.tls_certificates / 4eb18c169d46 / 6

- [https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms](data-sources--nfv_service--reference--group-003.md#canonical-aa42abb4870cdf7e379cf2055fb2a9b3bfae14d7879b645fe5d357ca3da31aed)
- [https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling](data-sources--nfv_service--reference--group-003.md#canonical-663a5127705c498b65ed12e5ad6d899733011a16295cde9e3390b51cf35cd532)
- [https_management.advertise_on_slo_sli.tls_certificates.private_key](data-sources--nfv_service--reference--group-003.md#canonical-be46b1c4c1902baae6b4e735cc4a4fa24885a02b5c76d6df7fad39bb4b160f69)
- [https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults](data-sources--nfv_service--reference--group-003.md#canonical-1f46448736c4a2029744c416889864ec08c2c3add10b061e7f2306b4c8c6e221)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-002.md#canonical-086d88128f5e5786a03853766b2c2649c0af61792f065c9bcfa6cdbb298fd85f)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-aa42abb4870cdf7e379cf2055fb2a9b3bfae14d7879b645fe5d357ca3da31aed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e435c2627e7c5a140d10b3f6e2b52b6d6a8b1f95ed8c66dacec05bc9106305da"></a>

## https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms — https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms / d6564a570aa7 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-002.md#canonical-086d88128f5e5786a03853766b2c2649c0af61792f065c9bcfa6cdbb298fd85f)
- [https_management.advertise_on_slo_sli.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-5964b4cebf3482d0880b7c54d204993a7c0648f12baf3a8071e4184a0f93bbea)
- https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms

<a id="canonical-79c7b2be0c40275f2b3fa8001ad50d56cbbb9f93c123fdc5a508afb6e29f9326"></a>

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

<a id="canonical-eb1e9c34512bf603a3db98350aab9ccf8594a147dcbf17875c2fa9bf91439d05"></a>

## Direct properties — https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms / d6564a570aa7 / 3

<a id="canonical-59f315f59c9fb84c0283304aea52cc6f1ea4a04b8bc68cde140b3b91bf937834"></a>

<a id="canonical-5b2f15e0dba43a6adf77e5098e79cebff46b29915dd95076e8fb2cac1cb758d6"></a>

## hash_algorithms property — https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms / d6564a570aa7 / 4

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

<a id="canonical-c81266be41060d78e19fb48903e1697bcde1b8bd8ea32db32e15ee3b92614197"></a>

## Next pages — https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms / d6564a570aa7 / 5

- [https_management.advertise_on_slo_sli.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-5964b4cebf3482d0880b7c54d204993a7c0648f12baf3a8071e4184a0f93bbea)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-663a5127705c498b65ed12e5ad6d899733011a16295cde9e3390b51cf35cd532"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e0301f06ad38b5534e4ef15fd09a6080e85d1cccc43e9eb0d4b8a01681f2228d"></a>

## https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling — https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling / fcde8b83a5f1 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-002.md#canonical-086d88128f5e5786a03853766b2c2649c0af61792f065c9bcfa6cdbb298fd85f)
- [https_management.advertise_on_slo_sli.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-5964b4cebf3482d0880b7c54d204993a7c0648f12baf3a8071e4184a0f93bbea)
- https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling

<a id="canonical-09d64b53b40ba47f5c771ddc77ae8a3641e008197e7e855a64a7ec64020ebb0a"></a>

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

<a id="canonical-98ea5dda5e8bfa1c615877a34628e70299f27697c204f0d895d77bd2caab8057"></a>

## Direct properties — https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling / fcde8b83a5f1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-02e8e1bf540ea4dc956f796bff3a1f68893aed3e3930e7e6bc0b841448f83443"></a>

## Next pages — https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling / fcde8b83a5f1 / 4

- [https_management.advertise_on_slo_sli.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-5964b4cebf3482d0880b7c54d204993a7c0648f12baf3a8071e4184a0f93bbea)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-be46b1c4c1902baae6b4e735cc4a4fa24885a02b5c76d6df7fad39bb4b160f69"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fe8263848a14e274fea898239143bbbeb92e10527e82960697d5eaa966a5d10b"></a>

## https_management.advertise_on_slo_sli.tls_certificates.private_key — https_management.advertise_on_slo_sli.tls_certificates.private_key / 20d1d98f9ce6 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-002.md#canonical-086d88128f5e5786a03853766b2c2649c0af61792f065c9bcfa6cdbb298fd85f)
- [https_management.advertise_on_slo_sli.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-5964b4cebf3482d0880b7c54d204993a7c0648f12baf3a8071e4184a0f93bbea)
- https_management.advertise_on_slo_sli.tls_certificates.private_key

<a id="canonical-84b843fc9fafde9d423d3df95bb103f91aac130fa0887188fbe62722663366b2"></a>

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

<a id="canonical-d7a0b4ef4c2851efcf97d0aea41002dd8b0ee6aaf4d40c0fdab509d55f50f49a"></a>

## Direct properties — https_management.advertise_on_slo_sli.tls_certificates.private_key / 20d1d98f9ce6 / 3

- [blindfold_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-1220fb76e58645dbb3a72a316fbe8e62d37733a05adc8bcbdef25f308db329f9): complete subsection reference.

- [clear_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-6d32ebf466bdd55aaa7e30acf083622ecb8380d5c9b82dafe4ffb46b24deb58a): complete subsection reference.

<a id="canonical-38cfe92b24aa558928ed6c350906ec2fe18244ecb3fb9341605788e087dcd387"></a>

## Next pages — https_management.advertise_on_slo_sli.tls_certificates.private_key / 20d1d98f9ce6 / 4

- [https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-1220fb76e58645dbb3a72a316fbe8e62d37733a05adc8bcbdef25f308db329f9)
- [https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-6d32ebf466bdd55aaa7e30acf083622ecb8380d5c9b82dafe4ffb46b24deb58a)
- [https_management.advertise_on_slo_sli.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-5964b4cebf3482d0880b7c54d204993a7c0648f12baf3a8071e4184a0f93bbea)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-1220fb76e58645dbb3a72a316fbe8e62d37733a05adc8bcbdef25f308db329f9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-427b0c4d463ea40a91f48259ad402792f3068118b8b48884edfab3617f9c8e95"></a>

## https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info — https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_sec / b2dd9c6cd851 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-002.md#canonical-086d88128f5e5786a03853766b2c2649c0af61792f065c9bcfa6cdbb298fd85f)
- [https_management.advertise_on_slo_sli.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-5964b4cebf3482d0880b7c54d204993a7c0648f12baf3a8071e4184a0f93bbea)
- [https_management.advertise_on_slo_sli.tls_certificates.private_key](data-sources--nfv_service--reference--group-003.md#canonical-be46b1c4c1902baae6b4e735cc4a4fa24885a02b5c76d6df7fad39bb4b160f69)
- https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-c757c94474ed7b4408987fef3e9dc9146b6b69af92c10a390a0f49c8972a10c7"></a>

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

<a id="canonical-0282968287af9bb14d3577c521140a69570ff15dca503060fbfe71d6c6e534e4"></a>

## Direct properties — https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_sec / b2dd9c6cd851 / 3

<a id="canonical-d6e7dea4072385e58f4e87e27750467467acb584a4884bd4f12c9959e483c0d3"></a>

<a id="canonical-2b49742007b47189dcd03637a66ebced37c4172f01fbf592ffe4c678798dd5e8"></a>

## decryption_provider property — https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_sec / b2dd9c6cd851 / 4

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

<a id="canonical-5a99515a2735d1c4a103d354521001bbeaa903062079e489bd6e3205ef77f3dd"></a>

<a id="canonical-9639e4c4c6c47a7e5ff03900a9a15a0ff308fea4047bef2d7cdd69dbc75d7aab"></a>

## location property — https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_sec / b2dd9c6cd851 / 5

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

<a id="canonical-262e1c90f40693f002f28505b0b8d9d12c53db1ce27771d652cd0ce440f65fbb"></a>

<a id="canonical-d27f9220c2d1e443d7787bed3a77a312582c1a3c6dc5a95d103e0a5de1c143f9"></a>

## store_provider property — https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_sec / b2dd9c6cd851 / 6

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

<a id="canonical-44eca90be704fb7b88c47feca68ee6f6fff64e6ed99b226eb107a0498bd1088d"></a>

## Next pages — https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_sec / b2dd9c6cd851 / 7

- [https_management.advertise_on_slo_sli.tls_certificates.private_key](data-sources--nfv_service--reference--group-003.md#canonical-be46b1c4c1902baae6b4e735cc4a4fa24885a02b5c76d6df7fad39bb4b160f69)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-6d32ebf466bdd55aaa7e30acf083622ecb8380d5c9b82dafe4ffb46b24deb58a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e16baa5cca2492eac9c6dc5810079e5e166157c073df6c8c7744d0d5e97676e"></a>

## https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info — https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_ / 7e9c818bb93f / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-002.md#canonical-086d88128f5e5786a03853766b2c2649c0af61792f065c9bcfa6cdbb298fd85f)
- [https_management.advertise_on_slo_sli.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-5964b4cebf3482d0880b7c54d204993a7c0648f12baf3a8071e4184a0f93bbea)
- [https_management.advertise_on_slo_sli.tls_certificates.private_key](data-sources--nfv_service--reference--group-003.md#canonical-be46b1c4c1902baae6b4e735cc4a4fa24885a02b5c76d6df7fad39bb4b160f69)
- https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info

<a id="canonical-d28130f973ec7b7d5a1d4fc264c5b68ccfdd189e75c49c967128fd0b53c9386b"></a>

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

<a id="canonical-7a37cc51a6eb847e596c3b9d59d8345eaab578be8fdc360555a57812773dc03c"></a>

## Direct properties — https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_ / 7e9c818bb93f / 3

<a id="canonical-8a60ffc19bc66d11c256b3107a2cac7fa06ed8d0b8beb09b55502dec3e4b8708"></a>

<a id="canonical-2ef5b4550646a4c2ee27470966d6d4e1f597065396945bf3764f71ae01c92231"></a>

## provider_ref property — https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_ / 7e9c818bb93f / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-e9a20bd620feb5bda585962058e7e8ee9dbe2a60cd992033cbf4cef0ee5a0fe1"></a>

<a id="canonical-75340577db58ca89de9436069f0ebe27205c4323815df993455bfd10bf82eafb"></a>

## url property — https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_ / 7e9c818bb93f / 5

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

<a id="canonical-83187c6d20fc0166fddda05f2570411435c4842ad0037b062d1e26f57ba78d10"></a>

## Next pages — https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_ / 7e9c818bb93f / 6

- [https_management.advertise_on_slo_sli.tls_certificates.private_key](data-sources--nfv_service--reference--group-003.md#canonical-be46b1c4c1902baae6b4e735cc4a4fa24885a02b5c76d6df7fad39bb4b160f69)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-1f46448736c4a2029744c416889864ec08c2c3add10b061e7f2306b4c8c6e221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-09cb75875b53a67e846cabdb7af37009f50362c5b95d139c3ed4c723840cb132"></a>

## https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults — https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults / a76d03d5e7e1 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-002.md#canonical-086d88128f5e5786a03853766b2c2649c0af61792f065c9bcfa6cdbb298fd85f)
- [https_management.advertise_on_slo_sli.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-5964b4cebf3482d0880b7c54d204993a7c0648f12baf3a8071e4184a0f93bbea)
- https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults

<a id="canonical-45c1d4fd51aaeb2e561880a399ea08d07d4b6d2b385bc9a1ab5e8f716ec6f220"></a>

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

<a id="canonical-f574279c90503e96fbcb30837ab51e0d9ad684e2b9aeedb5baa7412dd920160c"></a>

## Direct properties — https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults / a76d03d5e7e1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-627a72bbef5cf360cd174a3c0d38673aaaffd2cf37633d8290d883cafc2aa73a"></a>

## Next pages — https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults / a76d03d5e7e1 / 4

- [https_management.advertise_on_slo_sli.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-5964b4cebf3482d0880b7c54d204993a7c0648f12baf3a8071e4184a0f93bbea)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-079f523bfa7f99351fd61f57cc767cb2eb1e33ea71ab5e84e2155d35d598ba92"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c51d5cc48f9b2f49ed86c4b476d50a6f8c86c748aec00d2476775eeb02cae52"></a>

## https_management.advertise_on_slo_sli.tls_config — https_management.advertise_on_slo_sli.tls_config / ff815c5f0aad / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-002.md#canonical-086d88128f5e5786a03853766b2c2649c0af61792f065c9bcfa6cdbb298fd85f)
- https_management.advertise_on_slo_sli.tls_config

<a id="canonical-23f6066444f57c9f6b51ebfed1dad952e36465a9a8c8414c2a98a5a1beff09eb"></a>

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

<a id="canonical-f31585b805b628d541f8844ac78f038cd4c502c373f868b745a87bc8a7445e3b"></a>

## Direct properties — https_management.advertise_on_slo_sli.tls_config / ff815c5f0aad / 3

- [custom_security](data-sources--nfv_service--reference--group-003.md#canonical-d0b72160c6fba9842164b0c51d45043e14a42f0d14881b4895217a746ea5d208): complete subsection reference.

- [default_security](data-sources--nfv_service--reference--group-003.md#canonical-ac6731c2543a0ed2124346163fe5f08ba96280737b4195a49b28c20857e1b2ab): complete subsection reference.

- [low_security](data-sources--nfv_service--reference--group-003.md#canonical-39796de3596fe787555c6c0f0430dfade904da1c39f0442dfa39ae47798aac2e): complete subsection reference.

- [medium_security](data-sources--nfv_service--reference--group-003.md#canonical-e46d350baf14f69aa9052d04a9055c29e5181d2e37511953d8d3ac90cdd1e6e5): complete subsection reference.

<a id="canonical-61edb3d4b96c4cae46299878c1807bfd5ecc9194934c59b9fbfe156fe490e205"></a>

## Next pages — https_management.advertise_on_slo_sli.tls_config / ff815c5f0aad / 4

- [https_management.advertise_on_slo_sli.tls_config.custom_security](data-sources--nfv_service--reference--group-003.md#canonical-d0b72160c6fba9842164b0c51d45043e14a42f0d14881b4895217a746ea5d208)
- [https_management.advertise_on_slo_sli.tls_config.default_security](data-sources--nfv_service--reference--group-003.md#canonical-ac6731c2543a0ed2124346163fe5f08ba96280737b4195a49b28c20857e1b2ab)
- [https_management.advertise_on_slo_sli.tls_config.low_security](data-sources--nfv_service--reference--group-003.md#canonical-39796de3596fe787555c6c0f0430dfade904da1c39f0442dfa39ae47798aac2e)
- [https_management.advertise_on_slo_sli.tls_config.medium_security](data-sources--nfv_service--reference--group-003.md#canonical-e46d350baf14f69aa9052d04a9055c29e5181d2e37511953d8d3ac90cdd1e6e5)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-002.md#canonical-086d88128f5e5786a03853766b2c2649c0af61792f065c9bcfa6cdbb298fd85f)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-d0b72160c6fba9842164b0c51d45043e14a42f0d14881b4895217a746ea5d208"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c8d8b6cbf528df0bee6c0aeffc2a97807346f3cceb3654de46368c9d931f16f6"></a>

## https_management.advertise_on_slo_sli.tls_config.custom_security — https_management.advertise_on_slo_sli.tls_config.custom_security / 79bd2b40013b / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-002.md#canonical-086d88128f5e5786a03853766b2c2649c0af61792f065c9bcfa6cdbb298fd85f)
- [https_management.advertise_on_slo_sli.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-079f523bfa7f99351fd61f57cc767cb2eb1e33ea71ab5e84e2155d35d598ba92)
- https_management.advertise_on_slo_sli.tls_config.custom_security

<a id="canonical-83d1bc8acd2b912d7f71cfd3ae0e288c741ddc3f60cc8ac407fb6144683e4105"></a>

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

<a id="canonical-4405dafbddf2f60f09207aeb6417f7c9040f2548aa34e458ac6016f9c91b7aa1"></a>

## Direct properties — https_management.advertise_on_slo_sli.tls_config.custom_security / 79bd2b40013b / 3

<a id="canonical-dd45c0227f88504ba703e6ecf9934e65185b9da24d1787e3e42e8e09c7757047"></a>

<a id="canonical-8de6d53f53f4827affe5636336bda81b68bbc4bbbf1dc608c20ac65f65e1a224"></a>

## cipher_suites property — https_management.advertise_on_slo_sli.tls_config.custom_security / 79bd2b40013b / 4

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

<a id="canonical-859d8746c3718b8831461a25f854aae36da394cf0192706fa2f46f978f688d6e"></a>

<a id="canonical-f6fc39726261d4946ab12c7b6e467353a747a6a69c85999818eb5e014bbe57dc"></a>

## max_version property — https_management.advertise_on_slo_sli.tls_config.custom_security / 79bd2b40013b / 5

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

<a id="canonical-e660ab03a0fb8ece217f2e0111bfa39791b4470be604ea079dd4a77a277517da"></a>

<a id="canonical-b2f8df4da75f0af1ded17deec8fdbb77425067f09e4566d779a3947bf2ad8fc7"></a>

## min_version property — https_management.advertise_on_slo_sli.tls_config.custom_security / 79bd2b40013b / 6

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

<a id="canonical-8bcb7c1afffa7348a84b12e20e31c190a8c542bcedcfbca4ca07cbfe964d298b"></a>

## Next pages — https_management.advertise_on_slo_sli.tls_config.custom_security / 79bd2b40013b / 7

- [https_management.advertise_on_slo_sli.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-079f523bfa7f99351fd61f57cc767cb2eb1e33ea71ab5e84e2155d35d598ba92)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-ac6731c2543a0ed2124346163fe5f08ba96280737b4195a49b28c20857e1b2ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d812c25750fb4022c797ec654ee8f87cf81c7de942795b6b55563ac23433faad"></a>

## https_management.advertise_on_slo_sli.tls_config.default_security — https_management.advertise_on_slo_sli.tls_config.default_security / 212b35a1bb15 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-002.md#canonical-086d88128f5e5786a03853766b2c2649c0af61792f065c9bcfa6cdbb298fd85f)
- [https_management.advertise_on_slo_sli.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-079f523bfa7f99351fd61f57cc767cb2eb1e33ea71ab5e84e2155d35d598ba92)
- https_management.advertise_on_slo_sli.tls_config.default_security

<a id="canonical-1925de80a8d1d12c265e54b4fa215125d3b92e55519630fe804c6af12b7ed06f"></a>

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

<a id="canonical-a9b7133e557b69d4a18a4c4d3014ec5af48fce45cecb76fc0897bd008e4d0aff"></a>

## Direct properties — https_management.advertise_on_slo_sli.tls_config.default_security / 212b35a1bb15 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cb88f9cad7902b18138f255f8a8d054f61d2702c525dde8d1b7a4827de24e28c"></a>

## Next pages — https_management.advertise_on_slo_sli.tls_config.default_security / 212b35a1bb15 / 4

- [https_management.advertise_on_slo_sli.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-079f523bfa7f99351fd61f57cc767cb2eb1e33ea71ab5e84e2155d35d598ba92)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-39796de3596fe787555c6c0f0430dfade904da1c39f0442dfa39ae47798aac2e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f0b665d9dd9d0f1062c4b79bbb53b3a714a60959ba5707e3ba13c0814eefb543"></a>

## https_management.advertise_on_slo_sli.tls_config.low_security — https_management.advertise_on_slo_sli.tls_config.low_security / db1c3358efc1 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-002.md#canonical-086d88128f5e5786a03853766b2c2649c0af61792f065c9bcfa6cdbb298fd85f)
- [https_management.advertise_on_slo_sli.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-079f523bfa7f99351fd61f57cc767cb2eb1e33ea71ab5e84e2155d35d598ba92)
- https_management.advertise_on_slo_sli.tls_config.low_security

<a id="canonical-e94fe121dc0e9a30082278c233d763868f436ceba8ff36b8f91a0f4c3c966426"></a>

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

<a id="canonical-10a6c1f45a8df07ebcb749ca010bbcd4ab34d18cf20256dbb97947769e6153ac"></a>

## Direct properties — https_management.advertise_on_slo_sli.tls_config.low_security / db1c3358efc1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-912a74a9b72d4fd59a3670b40c0e68b7690b510b48136e437e6c3a231418df2c"></a>

## Next pages — https_management.advertise_on_slo_sli.tls_config.low_security / db1c3358efc1 / 4

- [https_management.advertise_on_slo_sli.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-079f523bfa7f99351fd61f57cc767cb2eb1e33ea71ab5e84e2155d35d598ba92)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-e46d350baf14f69aa9052d04a9055c29e5181d2e37511953d8d3ac90cdd1e6e5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-091b707c5dd7b22d0807e7d43d44b00bc5e3d0f09dfb1f0e0d57419ce5b50d14"></a>

## https_management.advertise_on_slo_sli.tls_config.medium_security — https_management.advertise_on_slo_sli.tls_config.medium_security / 2a26a119b160 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-002.md#canonical-086d88128f5e5786a03853766b2c2649c0af61792f065c9bcfa6cdbb298fd85f)
- [https_management.advertise_on_slo_sli.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-079f523bfa7f99351fd61f57cc767cb2eb1e33ea71ab5e84e2155d35d598ba92)
- https_management.advertise_on_slo_sli.tls_config.medium_security

<a id="canonical-f0fb041fd4aa111e0cc1de20273e94445f900afab39f2f0709ebdb3da127bb32"></a>

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

<a id="canonical-a1948ec8c4c23297b63132940afe29872d82e4f016353097f40da0e3503348f4"></a>

## Direct properties — https_management.advertise_on_slo_sli.tls_config.medium_security / 2a26a119b160 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e7e18ba86664f7ab1a40494d3696e4edb7446f58740a010a03c162078af9de3f"></a>

## Next pages — https_management.advertise_on_slo_sli.tls_config.medium_security / 2a26a119b160 / 4

- [https_management.advertise_on_slo_sli.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-079f523bfa7f99351fd61f57cc767cb2eb1e33ea71ab5e84e2155d35d598ba92)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-a81b93841729eecca0b56bf195a182d25d620ef5fb5b7bd2eebf84ece5ac3601"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be4f0c3096ea3a021ad1c5427fdeeb209d46b72e2e8dfa0d127314cfad538362"></a>

## https_management.advertise_on_slo_sli.use_mtls — https_management.advertise_on_slo_sli.use_mtls / 4950d2b0342d / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-002.md#canonical-086d88128f5e5786a03853766b2c2649c0af61792f065c9bcfa6cdbb298fd85f)
- https_management.advertise_on_slo_sli.use_mtls

<a id="canonical-28c641ed8d7d9e9db9973ddbb982cfd99d4cd27dfe01a178bae50903d296bfe6"></a>

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

<a id="canonical-3b86cbfa589c801111d778a21929b105a2836c00db38306529877c0706f7059a"></a>

## Direct properties — https_management.advertise_on_slo_sli.use_mtls / 4950d2b0342d / 3

<a id="canonical-5b9284e702c6b87173e3e3f209dae9fc2f2934b6459570ee32cfb104ed381f45"></a>

<a id="canonical-ad456ed5fb77447b5bedf60c09bf9002d0b299daad9a4d9a7a3a3610bec00089"></a>

## client_certificate_optional property — https_management.advertise_on_slo_sli.use_mtls / 4950d2b0342d / 4

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

- [crl](data-sources--nfv_service--reference--group-003.md#canonical-ad18c3dad0650fa09512d7fa7bce100a842d86a19497fb2e5b51d7ef6f1a53e3): complete subsection reference.

- [no_crl](data-sources--nfv_service--reference--group-003.md#canonical-74009a598efc773714f1f960bea5eb793056bf74a84a9338135886ba08c0dbb8): complete subsection reference.

- [trusted_ca](data-sources--nfv_service--reference--group-003.md#canonical-621aa9cf5589edb1a3b7d334bae98fbc7c350b68a819ed3b8fe9040e7daa6efe): complete subsection reference.

<a id="canonical-fe465369f2861e04b09594c8f592b7c7c7b53677f9be7bed030447bd5af34131"></a>

<a id="canonical-66a6ab62d4662fad98cd0d5e427c2a46931caa4994dfebef10040f1c21cd4d19"></a>

## trusted_ca_url property — https_management.advertise_on_slo_sli.use_mtls / 4950d2b0342d / 5

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

- [xfcc_disabled](data-sources--nfv_service--reference--group-003.md#canonical-8161acafd718425d4b71dc8ebcfaf7684a41952903a25286a9627f529e715442): complete subsection reference.

- [xfcc_options](data-sources--nfv_service--reference--group-003.md#canonical-0f2755fff2ee5d24be9d758e80ae3f42f73483833ceba8e206736a3b111d69ca): complete subsection reference.

<a id="canonical-6855b0558fdcc4cbac782bd748c2df12f4c537091284c68dda6872d6d6df09fe"></a>

## Next pages — https_management.advertise_on_slo_sli.use_mtls / 4950d2b0342d / 6

- [https_management.advertise_on_slo_sli.use_mtls.crl](data-sources--nfv_service--reference--group-003.md#canonical-ad18c3dad0650fa09512d7fa7bce100a842d86a19497fb2e5b51d7ef6f1a53e3)
- [https_management.advertise_on_slo_sli.use_mtls.no_crl](data-sources--nfv_service--reference--group-003.md#canonical-74009a598efc773714f1f960bea5eb793056bf74a84a9338135886ba08c0dbb8)
- [https_management.advertise_on_slo_sli.use_mtls.trusted_ca](data-sources--nfv_service--reference--group-003.md#canonical-621aa9cf5589edb1a3b7d334bae98fbc7c350b68a819ed3b8fe9040e7daa6efe)
- [https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled](data-sources--nfv_service--reference--group-003.md#canonical-8161acafd718425d4b71dc8ebcfaf7684a41952903a25286a9627f529e715442)
- [https_management.advertise_on_slo_sli.use_mtls.xfcc_options](data-sources--nfv_service--reference--group-003.md#canonical-0f2755fff2ee5d24be9d758e80ae3f42f73483833ceba8e206736a3b111d69ca)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-002.md#canonical-086d88128f5e5786a03853766b2c2649c0af61792f065c9bcfa6cdbb298fd85f)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-ad18c3dad0650fa09512d7fa7bce100a842d86a19497fb2e5b51d7ef6f1a53e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-04ff8ae1b5b4f62aa1afd508a8bc74fb33562ff7b329a51a706af7119f815b0a"></a>

## https_management.advertise_on_slo_sli.use_mtls.crl — https_management.advertise_on_slo_sli.use_mtls.crl / 63014a5502e0 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-002.md#canonical-086d88128f5e5786a03853766b2c2649c0af61792f065c9bcfa6cdbb298fd85f)
- [https_management.advertise_on_slo_sli.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-a81b93841729eecca0b56bf195a182d25d620ef5fb5b7bd2eebf84ece5ac3601)
- https_management.advertise_on_slo_sli.use_mtls.crl

<a id="canonical-4c94d6bcf025780e615d9c073800462905d890de39316978a9ded8594a4cd89b"></a>

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

<a id="canonical-0406aef5f4a285614e176dd47cbf5148781eca7decad79412dad323eacbcf2ba"></a>

## Direct properties — https_management.advertise_on_slo_sli.use_mtls.crl / 63014a5502e0 / 3

<a id="canonical-f64fb6a2d19d8f2f443c8341bd33d59abfe16221d677f8250278dafa148e889d"></a>

<a id="canonical-1d841a1bdfdbe6d4cb5622228317d931977a0810f55c2351df27c4ba1a8d341b"></a>

## name property — https_management.advertise_on_slo_sli.use_mtls.crl / 63014a5502e0 / 4

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

<a id="canonical-97f5f93866e8a7ac7408b7b75f8baa0e53d4d766fe3511fb8c0c0c116c285685"></a>

<a id="canonical-46f32d0b5f744bf83414123a20c6992b8c932cf6e985d9dacc2f0a4b6804b68c"></a>

## namespace property — https_management.advertise_on_slo_sli.use_mtls.crl / 63014a5502e0 / 5

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

<a id="canonical-84cf096ddccb4e8d34ca5884302fbaae8eba9ec859e2279d63f237b385fdbf80"></a>

<a id="canonical-04f8af664d641b166df7b4991e930b71968f2d8533e4da1625796d5a2dfe5399"></a>

## tenant property — https_management.advertise_on_slo_sli.use_mtls.crl / 63014a5502e0 / 6

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

<a id="canonical-8b0a742209f718f1040a4d0355cc0c2a8d2f2fe81695e8372dd6ef058b00cfff"></a>

## Next pages — https_management.advertise_on_slo_sli.use_mtls.crl / 63014a5502e0 / 7

- [https_management.advertise_on_slo_sli.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-a81b93841729eecca0b56bf195a182d25d620ef5fb5b7bd2eebf84ece5ac3601)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-74009a598efc773714f1f960bea5eb793056bf74a84a9338135886ba08c0dbb8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f711afad9aa8b4bd14325d0a781fd1ce49eacbe12c62429c97dca003fc873e18"></a>

## https_management.advertise_on_slo_sli.use_mtls.no_crl — https_management.advertise_on_slo_sli.use_mtls.no_crl / 5258671a809b / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-002.md#canonical-086d88128f5e5786a03853766b2c2649c0af61792f065c9bcfa6cdbb298fd85f)
- [https_management.advertise_on_slo_sli.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-a81b93841729eecca0b56bf195a182d25d620ef5fb5b7bd2eebf84ece5ac3601)
- https_management.advertise_on_slo_sli.use_mtls.no_crl

<a id="canonical-532821a9450cbd0ad155a7152b1032c85de0ca23a70953c6ca3c3556e497c10c"></a>

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

<a id="canonical-b1bdf5e96e863fc03aaa7b225feb7e09ddf3435eaf659ac53dd0d876e761f753"></a>

## Direct properties — https_management.advertise_on_slo_sli.use_mtls.no_crl / 5258671a809b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-656602fdfc8404f74b6bd29ba24a1215974160a27b97eb4e6db32d9064e035ee"></a>

## Next pages — https_management.advertise_on_slo_sli.use_mtls.no_crl / 5258671a809b / 4

- [https_management.advertise_on_slo_sli.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-a81b93841729eecca0b56bf195a182d25d620ef5fb5b7bd2eebf84ece5ac3601)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-621aa9cf5589edb1a3b7d334bae98fbc7c350b68a819ed3b8fe9040e7daa6efe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-75c107db46abdc34a7f3313677f0f5204528ea4175ffe944a7f9bd6eab0964bc"></a>

## https_management.advertise_on_slo_sli.use_mtls.trusted_ca — https_management.advertise_on_slo_sli.use_mtls.trusted_ca / f2f99cb290a6 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-002.md#canonical-086d88128f5e5786a03853766b2c2649c0af61792f065c9bcfa6cdbb298fd85f)
- [https_management.advertise_on_slo_sli.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-a81b93841729eecca0b56bf195a182d25d620ef5fb5b7bd2eebf84ece5ac3601)
- https_management.advertise_on_slo_sli.use_mtls.trusted_ca

<a id="canonical-bbe167848ce1428b4b53739b4d1165fe0c165974223c5621e330ecd805b6e556"></a>

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

<a id="canonical-2d2657d5b27c5ee79d69fde8168415fc9f915f5de8c2cce9e4d74a736bfb16e4"></a>

## Direct properties — https_management.advertise_on_slo_sli.use_mtls.trusted_ca / f2f99cb290a6 / 3

<a id="canonical-4e4bbaa4497827f5bef69e37d6b27a252d95326d292a4b0021ac4b7ccb7abdca"></a>

<a id="canonical-00a70157d5b4e102dff4e4b25f9ad01ca33d49e943db3cf2a70e7924baff42ed"></a>

## name property — https_management.advertise_on_slo_sli.use_mtls.trusted_ca / f2f99cb290a6 / 4

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

<a id="canonical-3dd59b5dcd5f2764028c368bc4135247ee00c4179f610043092ec1c0a39d97c5"></a>

<a id="canonical-dfe9947c52893fe921e6b979e2256a966fb3efddeae897588cc94c0cbf42e2cc"></a>

## namespace property — https_management.advertise_on_slo_sli.use_mtls.trusted_ca / f2f99cb290a6 / 5

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

<a id="canonical-ee7ebd9dfc78dc0a4a9921621677c1f0872042f2d61a6b4a429c839a4fd55bd0"></a>

<a id="canonical-8821d6b769bf13005a238d508b4a1d559c6196155c9875954929253a7ab6fb67"></a>

## tenant property — https_management.advertise_on_slo_sli.use_mtls.trusted_ca / f2f99cb290a6 / 6

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

<a id="canonical-e874300b2b1eb49edcc01bfb856ffe5cb49a63695745d2e261f5677cb1776ee3"></a>

## Next pages — https_management.advertise_on_slo_sli.use_mtls.trusted_ca / f2f99cb290a6 / 7

- [https_management.advertise_on_slo_sli.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-a81b93841729eecca0b56bf195a182d25d620ef5fb5b7bd2eebf84ece5ac3601)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-8161acafd718425d4b71dc8ebcfaf7684a41952903a25286a9627f529e715442"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0ddf590cf631df7452a919954c4c2bd883d2497a5d88d06af9a7bf705af9cee9"></a>

## https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled — https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled / b7941cc5cd09 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-002.md#canonical-086d88128f5e5786a03853766b2c2649c0af61792f065c9bcfa6cdbb298fd85f)
- [https_management.advertise_on_slo_sli.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-a81b93841729eecca0b56bf195a182d25d620ef5fb5b7bd2eebf84ece5ac3601)
- https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled

<a id="canonical-d3d6e282b61c5fe4c3804375233a947d6f8429e7307206c79e316be8079abd1a"></a>

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

<a id="canonical-0fec0c8b70537918bac867f894858fb58ec3e2f93cbb7e8f126b88b7a1c41803"></a>

## Direct properties — https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled / b7941cc5cd09 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-826e4d996e704273b28759d643ce8f445021ee4991c994354892404fc131d481"></a>

## Next pages — https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled / b7941cc5cd09 / 4

- [https_management.advertise_on_slo_sli.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-a81b93841729eecca0b56bf195a182d25d620ef5fb5b7bd2eebf84ece5ac3601)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-0f2755fff2ee5d24be9d758e80ae3f42f73483833ceba8e206736a3b111d69ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b9f6dd967325e0a8e214cccf746eedb79762535f94b717c4ed804f1c5718ec5"></a>

## https_management.advertise_on_slo_sli.use_mtls.xfcc_options — https_management.advertise_on_slo_sli.use_mtls.xfcc_options / 2a9ca8f45d64 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-002.md#canonical-086d88128f5e5786a03853766b2c2649c0af61792f065c9bcfa6cdbb298fd85f)
- [https_management.advertise_on_slo_sli.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-a81b93841729eecca0b56bf195a182d25d620ef5fb5b7bd2eebf84ece5ac3601)
- https_management.advertise_on_slo_sli.use_mtls.xfcc_options

<a id="canonical-72ca314c00ac27ddeb6fd4f5e2af73ed5c73cdd9b5dcd4269569544041662b3b"></a>

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

<a id="canonical-cb6e993981e5b53133f75f954ee2b118c1b789e006010e005dee4c8f60514eb1"></a>

## Direct properties — https_management.advertise_on_slo_sli.use_mtls.xfcc_options / 2a9ca8f45d64 / 3

<a id="canonical-dbaed32b004001951e8b80b17af6b9dbd4a4f1879a476c1e87895015c66d18be"></a>

<a id="canonical-58d52a0061e56c5784e7560a8f2bd87daa5c444023f49877408bbde59baf4703"></a>

## xfcc_header_elements property — https_management.advertise_on_slo_sli.use_mtls.xfcc_options / 2a9ca8f45d64 / 4

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

<a id="canonical-6b157f876b44b1d6b04c6aae06e9feeb19f78299e96d72891393ee6d14e3091d"></a>

## Next pages — https_management.advertise_on_slo_sli.use_mtls.xfcc_options / 2a9ca8f45d64 / 5

- [https_management.advertise_on_slo_sli.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-a81b93841729eecca0b56bf195a182d25d620ef5fb5b7bd2eebf84ece5ac3601)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-9cb7d734443534e78f77f5035f885e992d26930de647b960b4a5a8542d04f2a7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-493b8ae4675825322ae7039c2e0257970c06eb6721371736e726f825e4cd70d8"></a>

## https_management.advertise_on_slo_vip — https_management.advertise_on_slo_vip / 117c9b69ce59 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- https_management.advertise_on_slo_vip

<a id="canonical-61c35ba4d8fdbfa689005914e6e26d93208ce25252f7ccb02c4d234ec1ad5ddc"></a>

Type: `"single"`. Computed.

Inline TLS Parameters. Inline TLS parameters.

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

<a id="canonical-2928cc577f316593dca4f545f201cb6693a47fe4e35da7d6db2d4e2decb07754"></a>

## Direct properties — https_management.advertise_on_slo_vip / 117c9b69ce59 / 3

- [no_mtls](data-sources--nfv_service--reference--group-003.md#canonical-e897a2c601b906b0ac7e451ad82236688dd2a8586039be0e701bf78204e07019): complete subsection reference.

- [tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-20052c654f2618ef4cb83f6f15a638621a8f9049ebbf5264fa805b3e4cff6d4d): complete subsection reference.

- [tls_config](data-sources--nfv_service--reference--group-003.md#canonical-dcdd49a2ef8dd8cd70a62742d664c720f5e0f1983fc540088b7d48b06c825cfa): complete subsection reference.

- [use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-137dadd6204fcc9ee8516d19f54c9986ee16cac241901172042c9f9f35d39930): complete subsection reference.

<a id="canonical-d27b0d67a9ff7255b27b6c386a40409de378429f40dfc9f7b3e20154584b5f18"></a>

## Next pages — https_management.advertise_on_slo_vip / 117c9b69ce59 / 4

- [https_management.advertise_on_slo_vip.no_mtls](data-sources--nfv_service--reference--group-003.md#canonical-e897a2c601b906b0ac7e451ad82236688dd2a8586039be0e701bf78204e07019)
- [https_management.advertise_on_slo_vip.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-20052c654f2618ef4cb83f6f15a638621a8f9049ebbf5264fa805b3e4cff6d4d)
- [https_management.advertise_on_slo_vip.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-dcdd49a2ef8dd8cd70a62742d664c720f5e0f1983fc540088b7d48b06c825cfa)
- [https_management.advertise_on_slo_vip.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-137dadd6204fcc9ee8516d19f54c9986ee16cac241901172042c9f9f35d39930)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-e897a2c601b906b0ac7e451ad82236688dd2a8586039be0e701bf78204e07019"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1df49aa6270ebcdfd022cd641165e3eb0f2d5190b72b79b2d6da411b7086ab9f"></a>

## https_management.advertise_on_slo_vip.no_mtls — https_management.advertise_on_slo_vip.no_mtls / 359e000f2d5c / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-9cb7d734443534e78f77f5035f885e992d26930de647b960b4a5a8542d04f2a7)
- https_management.advertise_on_slo_vip.no_mtls

<a id="canonical-6883bf7ece16d5e676022b1e45c5027f0fd7e2e21524dd675cd586663062b76b"></a>

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

<a id="canonical-1f627a364e7c7c07eb133b553d8f9fcaab1bf2a09112df17116d4a6efdcecf19"></a>

## Direct properties — https_management.advertise_on_slo_vip.no_mtls / 359e000f2d5c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-72944f35a35f8209b4578cdf5da31330cb4cf9897fcdd2ca2990ac115b316d89"></a>

## Next pages — https_management.advertise_on_slo_vip.no_mtls / 359e000f2d5c / 4

- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-9cb7d734443534e78f77f5035f885e992d26930de647b960b4a5a8542d04f2a7)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-20052c654f2618ef4cb83f6f15a638621a8f9049ebbf5264fa805b3e4cff6d4d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c35ccce3cb618204fffff2b886a375e8c99308aff1172360333c7db0f6e930b"></a>

## https_management.advertise_on_slo_vip.tls_certificates — https_management.advertise_on_slo_vip.tls_certificates / c7629a67df7b / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-9cb7d734443534e78f77f5035f885e992d26930de647b960b4a5a8542d04f2a7)
- https_management.advertise_on_slo_vip.tls_certificates

<a id="canonical-1969382a46156908d4321597462fca7de565b3a5ef949c629b20e9d38ad9d033"></a>

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

<a id="canonical-ebe48c53b33173752cc38b87e28339b08f9f9595277ba8bb0d25101020997974"></a>

## Direct properties — https_management.advertise_on_slo_vip.tls_certificates / c7629a67df7b / 3

<a id="canonical-3aedda1d02e16268eef4acb1d3fa4936b1e6b18d0481e0e51b9e094d83a598bd"></a>

<a id="canonical-4174521d494890493a5104371c8550138989e11331de464c204facaf78d3420f"></a>

## certificate_url property — https_management.advertise_on_slo_vip.tls_certificates / c7629a67df7b / 4

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

- [custom_hash_algorithms](data-sources--nfv_service--reference--group-003.md#canonical-2902689ce4dc4284ae5e46500508caf16ebdff674e7a2494b90a8b23d83aae28): complete subsection reference.

<a id="canonical-db80fdb39468faa9bdde6f87e9b53e8b0a8fd21d1ebb40dc2e8ca269f6b35d2f"></a>

<a id="canonical-352aaccc784493512dc7a73a16414a3fd0ddf99837ea311414354d71c617c069"></a>

## description_spec property — https_management.advertise_on_slo_vip.tls_certificates / c7629a67df7b / 5

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--nfv_service--reference--group-003.md#canonical-d33d0dc24183c7c0db277037eed9f611231f254c53f120a2e1e36779e94842ef): complete subsection reference.

- [private_key](data-sources--nfv_service--reference--group-003.md#canonical-4ab00bda07c3b2338db263ff96044900b5357f519db69c5eb3ff9e328fbe07e1): complete subsection reference.

- [use_system_defaults](data-sources--nfv_service--reference--group-003.md#canonical-d805f5dd595dea65a8afffa45ff29a02be7ca978d46fe4eb467b856e6ef9db55): complete subsection reference.

<a id="canonical-9ff6df9697fc6d2d855cf6ddd1fc3afb86703e45eace667948f4d3f193673019"></a>

## Next pages — https_management.advertise_on_slo_vip.tls_certificates / c7629a67df7b / 6

- [https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms](data-sources--nfv_service--reference--group-003.md#canonical-2902689ce4dc4284ae5e46500508caf16ebdff674e7a2494b90a8b23d83aae28)
- [https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling](data-sources--nfv_service--reference--group-003.md#canonical-d33d0dc24183c7c0db277037eed9f611231f254c53f120a2e1e36779e94842ef)
- [https_management.advertise_on_slo_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-003.md#canonical-4ab00bda07c3b2338db263ff96044900b5357f519db69c5eb3ff9e328fbe07e1)
- [https_management.advertise_on_slo_vip.tls_certificates.use_system_defaults](data-sources--nfv_service--reference--group-003.md#canonical-d805f5dd595dea65a8afffa45ff29a02be7ca978d46fe4eb467b856e6ef9db55)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-9cb7d734443534e78f77f5035f885e992d26930de647b960b4a5a8542d04f2a7)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-2902689ce4dc4284ae5e46500508caf16ebdff674e7a2494b90a8b23d83aae28"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0def88ffa213494b057df96ed23b9abdb7bf0aabb1a372108f55c788dfa0a73"></a>

## https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms — https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms / 564ab758c8f5 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-9cb7d734443534e78f77f5035f885e992d26930de647b960b4a5a8542d04f2a7)
- [https_management.advertise_on_slo_vip.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-20052c654f2618ef4cb83f6f15a638621a8f9049ebbf5264fa805b3e4cff6d4d)
- https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms

<a id="canonical-69925e5b27787fb406173f1baef2609ee6e067f1ca515e255d37fc0b24127ec4"></a>

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

<a id="canonical-74cba53c7bc1a1510c6d6a6939d3854559b5d5ea0b5fb9212d2f23de152aac52"></a>

## Direct properties — https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms / 564ab758c8f5 / 3

<a id="canonical-b29437ece9e8feddc3b33bed83bb8083850bf46c6683a176240e8cd1b9797dbf"></a>

<a id="canonical-0743775a45c2944d8428da73178ae33294c98e00c04ad9ccb63bc97558c78adb"></a>

## hash_algorithms property — https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms / 564ab758c8f5 / 4

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

<a id="canonical-f8b0861d6f79c7708e752d87c1cee608ccac475a6f68e14a4d5d26db78b0d9ae"></a>

## Next pages — https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms / 564ab758c8f5 / 5

- [https_management.advertise_on_slo_vip.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-20052c654f2618ef4cb83f6f15a638621a8f9049ebbf5264fa805b3e4cff6d4d)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-d33d0dc24183c7c0db277037eed9f611231f254c53f120a2e1e36779e94842ef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-76fb1f6e30da23f2158af9b89e35686adb5fb92fe4696bbcfe31be15b29fd2c3"></a>

## https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling — https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling / 2e3b818c01b1 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-9cb7d734443534e78f77f5035f885e992d26930de647b960b4a5a8542d04f2a7)
- [https_management.advertise_on_slo_vip.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-20052c654f2618ef4cb83f6f15a638621a8f9049ebbf5264fa805b3e4cff6d4d)
- https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling

<a id="canonical-8be4e97e13eeaf538cacb8152d72b481ec009380da95db2e15a283559a12f194"></a>

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

<a id="canonical-3b80ffb80ff2ad4ecb99232269502f613a805a7c651cec00dbab843e5869d8c0"></a>

## Direct properties — https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling / 2e3b818c01b1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f735ba8fc804ec4cf33fbd3ac50422016f20caa71f3997ad5eaae9f26854478f"></a>

## Next pages — https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling / 2e3b818c01b1 / 4

- [https_management.advertise_on_slo_vip.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-20052c654f2618ef4cb83f6f15a638621a8f9049ebbf5264fa805b3e4cff6d4d)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-4ab00bda07c3b2338db263ff96044900b5357f519db69c5eb3ff9e328fbe07e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7ae989589c0465c22e432b2b4b5f98e3ba0c7f8e9dacb7beb01fa238e0d2f3d"></a>

## https_management.advertise_on_slo_vip.tls_certificates.private_key — https_management.advertise_on_slo_vip.tls_certificates.private_key / 4f9e4c76cf70 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-9cb7d734443534e78f77f5035f885e992d26930de647b960b4a5a8542d04f2a7)
- [https_management.advertise_on_slo_vip.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-20052c654f2618ef4cb83f6f15a638621a8f9049ebbf5264fa805b3e4cff6d4d)
- https_management.advertise_on_slo_vip.tls_certificates.private_key

<a id="canonical-dbb7055fd40b13c2dbbf28743457850e82e070934e285b7a84c66fbde59b9684"></a>

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

<a id="canonical-704d6b6de191bf659ad47e62570faef38119c920baa57de7fc6514b58834c4cf"></a>

## Direct properties — https_management.advertise_on_slo_vip.tls_certificates.private_key / 4f9e4c76cf70 / 3

- [blindfold_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-0de604bbcf380924d8dd20b42c36aeeaaf97e8d9ce16f2e0bba1e2cc0ad3aaa4): complete subsection reference.

- [clear_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-582967961b22633d756b6c3cdbe10a5f78369a5f440e600836e8329e8371ad1d): complete subsection reference.

<a id="canonical-250d8efb02c078f9745375d1c294ebaeddb0a48209f4e48ad5858825b847da98"></a>

## Next pages — https_management.advertise_on_slo_vip.tls_certificates.private_key / 4f9e4c76cf70 / 4

- [https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-0de604bbcf380924d8dd20b42c36aeeaaf97e8d9ce16f2e0bba1e2cc0ad3aaa4)
- [https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-582967961b22633d756b6c3cdbe10a5f78369a5f440e600836e8329e8371ad1d)
- [https_management.advertise_on_slo_vip.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-20052c654f2618ef4cb83f6f15a638621a8f9049ebbf5264fa805b3e4cff6d4d)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-0de604bbcf380924d8dd20b42c36aeeaaf97e8d9ce16f2e0bba1e2cc0ad3aaa4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ff4f57b851d26f00cb80355b320247242737e082f8871c7e1879653d1facb5f6"></a>

## https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info — https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_sec / 5e718b797bfc / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-9cb7d734443534e78f77f5035f885e992d26930de647b960b4a5a8542d04f2a7)
- [https_management.advertise_on_slo_vip.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-20052c654f2618ef4cb83f6f15a638621a8f9049ebbf5264fa805b3e4cff6d4d)
- [https_management.advertise_on_slo_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-003.md#canonical-4ab00bda07c3b2338db263ff96044900b5357f519db69c5eb3ff9e328fbe07e1)
- https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-c5e35e996905f8653c4a7cfc6c72839ca9702f6955873aa1f2d0425f252c2f55"></a>

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

<a id="canonical-7bedd556b695fcca3c8ac8dc6b9c70b2f7069d6b4011b81acaab26a2750381c5"></a>

## Direct properties — https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_sec / 5e718b797bfc / 3

<a id="canonical-da9786d03f5632d68af467c44c8b114bacf1de4ee5132618d79c5b5c26501966"></a>

<a id="canonical-0cef9c52f343f38a5a2ec01c5f683f79e39d571dc837cac65987bb126651459b"></a>

## decryption_provider property — https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_sec / 5e718b797bfc / 4

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

<a id="canonical-a1cb096441ddff7f2aae058f8c7d1eeb09473dcee578a3a3bd2846dd4c574b85"></a>

<a id="canonical-9c47351764c803d722ea7e9b3d2ffbdc9aaf8a95ccc44c539c185f9367532c34"></a>

## location property — https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_sec / 5e718b797bfc / 5

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

<a id="canonical-7e822f258f5602a099fe9987e6e66a6d627367c9a43a18ef0ea66168607d6f64"></a>

<a id="canonical-6d7a9e26a541446a5594f41b6c0c7d9841187de781451e8698a7cd4594dbee75"></a>

## store_provider property — https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_sec / 5e718b797bfc / 6

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

<a id="canonical-bad2b4ae569ef254cea13c5528aabec57b59eabd687dda3d4b4a9920a0d0371c"></a>

## Next pages — https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_sec / 5e718b797bfc / 7

- [https_management.advertise_on_slo_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-003.md#canonical-4ab00bda07c3b2338db263ff96044900b5357f519db69c5eb3ff9e328fbe07e1)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-582967961b22633d756b6c3cdbe10a5f78369a5f440e600836e8329e8371ad1d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-64a94b78925a5b25e35b505748e44ad9e3e9bef14080b3cd93375666a8c59db2"></a>

## https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info — https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_ / 17ba88f2d963 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-9cb7d734443534e78f77f5035f885e992d26930de647b960b4a5a8542d04f2a7)
- [https_management.advertise_on_slo_vip.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-20052c654f2618ef4cb83f6f15a638621a8f9049ebbf5264fa805b3e4cff6d4d)
- [https_management.advertise_on_slo_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-003.md#canonical-4ab00bda07c3b2338db263ff96044900b5357f519db69c5eb3ff9e328fbe07e1)
- https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info

<a id="canonical-b95ab7147060d75d7cb9919afccf50201d16bad99c6b7c38ae58e4e37abf7cf6"></a>

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

<a id="canonical-8107a724d5c415783f8aebdcaebf838b1607f685cf600abac98433d6191442b5"></a>

## Direct properties — https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_ / 17ba88f2d963 / 3

<a id="canonical-e2e6e2255b3c35c97f03c773fa5967a2f371ae160cc1ec1a96ab390b15241947"></a>

<a id="canonical-9946c9528a8681d8216520db57893090dea2fe55fe27048871d6fc2f8d77012d"></a>

## provider_ref property — https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_ / 17ba88f2d963 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-19958dde59a9c703ebd5727fa94b60560fc170ce24295aff6655817a7d245798"></a>

<a id="canonical-7285c06a7826ac9213f8dc1bcb068597ef3816747f12c47cd9783017f6108338"></a>

## url property — https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_ / 17ba88f2d963 / 5

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

<a id="canonical-feafc1fe2ad1ae83f433edb9261a703f8b3e1c7eefab7e431ba302b71da158e1"></a>

## Next pages — https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_ / 17ba88f2d963 / 6

- [https_management.advertise_on_slo_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-003.md#canonical-4ab00bda07c3b2338db263ff96044900b5357f519db69c5eb3ff9e328fbe07e1)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-d805f5dd595dea65a8afffa45ff29a02be7ca978d46fe4eb467b856e6ef9db55"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-14f4666ef8de938615643870bbca4c710c4ad08b6e102912682d96b2f2184076"></a>

## https_management.advertise_on_slo_vip.tls_certificates.use_system_defaults — https_management.advertise_on_slo_vip.tls_certificates.use_system_defaults / 436b1db4b469 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-9cb7d734443534e78f77f5035f885e992d26930de647b960b4a5a8542d04f2a7)
- [https_management.advertise_on_slo_vip.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-20052c654f2618ef4cb83f6f15a638621a8f9049ebbf5264fa805b3e4cff6d4d)
- https_management.advertise_on_slo_vip.tls_certificates.use_system_defaults

<a id="canonical-9721119e7389eaac5213504c8fc073f4915627506eaef7ecb81e52bd1beb9013"></a>

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

<a id="canonical-2a29472280dd0b686cb534fe0aa568de9026a79c610f1ccb3c85ea1ee9e0023b"></a>

## Direct properties — https_management.advertise_on_slo_vip.tls_certificates.use_system_defaults / 436b1db4b469 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-10884c4a5732539f20c31deebe92cd7dbb628a6697346c9bb0849e67788acb88"></a>

## Next pages — https_management.advertise_on_slo_vip.tls_certificates.use_system_defaults / 436b1db4b469 / 4

- [https_management.advertise_on_slo_vip.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-20052c654f2618ef4cb83f6f15a638621a8f9049ebbf5264fa805b3e4cff6d4d)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-dcdd49a2ef8dd8cd70a62742d664c720f5e0f1983fc540088b7d48b06c825cfa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-421e84270683568e58e398e71c2860c6979a58b06894b2fcfc80ded3fe43391d"></a>

## https_management.advertise_on_slo_vip.tls_config — https_management.advertise_on_slo_vip.tls_config / a0d82419b3c5 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-9cb7d734443534e78f77f5035f885e992d26930de647b960b4a5a8542d04f2a7)
- https_management.advertise_on_slo_vip.tls_config

<a id="canonical-3dd7a31bdb7d56eb79206dfbdf702d476f6728852704415e85b63d95579761f4"></a>

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

<a id="canonical-8e30730a1bb73e0accaf1e1ea2f66b5b0e6f7f612adee73f271c2eba3061f6e3"></a>

## Direct properties — https_management.advertise_on_slo_vip.tls_config / a0d82419b3c5 / 3

- [custom_security](data-sources--nfv_service--reference--group-003.md#canonical-13df29bc2e8df9cd4320f668d2ad94cd1e619503190e8a8d743fe8c14fa98a3f): complete subsection reference.

- [default_security](data-sources--nfv_service--reference--group-003.md#canonical-2f1cbb9a82207b778a846cc32a4a6a6eb4779b301f3b5ab1000a04701b7bf525): complete subsection reference.

- [low_security](data-sources--nfv_service--reference--group-003.md#canonical-a0065a03a2857f7d48b0e64fc16c076f35b6577984a7eef47eeb514c830d9181): complete subsection reference.

- [medium_security](data-sources--nfv_service--reference--group-003.md#canonical-6644d8b4098fdab6c6c186d82bbb6b38d284ba4f2f3675490438451de589dcd9): complete subsection reference.

<a id="canonical-5cc13b8ea0077d0a55412676dcaaa687c3534a5938077cef830a991bcc0915e4"></a>

## Next pages — https_management.advertise_on_slo_vip.tls_config / a0d82419b3c5 / 4

- [https_management.advertise_on_slo_vip.tls_config.custom_security](data-sources--nfv_service--reference--group-003.md#canonical-13df29bc2e8df9cd4320f668d2ad94cd1e619503190e8a8d743fe8c14fa98a3f)
- [https_management.advertise_on_slo_vip.tls_config.default_security](data-sources--nfv_service--reference--group-003.md#canonical-2f1cbb9a82207b778a846cc32a4a6a6eb4779b301f3b5ab1000a04701b7bf525)
- [https_management.advertise_on_slo_vip.tls_config.low_security](data-sources--nfv_service--reference--group-003.md#canonical-a0065a03a2857f7d48b0e64fc16c076f35b6577984a7eef47eeb514c830d9181)
- [https_management.advertise_on_slo_vip.tls_config.medium_security](data-sources--nfv_service--reference--group-003.md#canonical-6644d8b4098fdab6c6c186d82bbb6b38d284ba4f2f3675490438451de589dcd9)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-9cb7d734443534e78f77f5035f885e992d26930de647b960b4a5a8542d04f2a7)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-13df29bc2e8df9cd4320f668d2ad94cd1e619503190e8a8d743fe8c14fa98a3f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4af19e686da7c995fc701ffee82d962744ef38f4129de883990685ac18bd8bb2"></a>

## https_management.advertise_on_slo_vip.tls_config.custom_security — https_management.advertise_on_slo_vip.tls_config.custom_security / 505689d70009 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-9cb7d734443534e78f77f5035f885e992d26930de647b960b4a5a8542d04f2a7)
- [https_management.advertise_on_slo_vip.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-dcdd49a2ef8dd8cd70a62742d664c720f5e0f1983fc540088b7d48b06c825cfa)
- https_management.advertise_on_slo_vip.tls_config.custom_security

<a id="canonical-f254988aee1cb018793db493c2e36e81c359cc89e2d2501d852be9b125ba5114"></a>

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

<a id="canonical-4aef64093b06400171bbfbf7cd6d4c8d503f51d52bae7306bcf8854b82c9bdfb"></a>

## Direct properties — https_management.advertise_on_slo_vip.tls_config.custom_security / 505689d70009 / 3

<a id="canonical-183785fc7836679673f5c643f02a5415807704b8d6c4a425d510f87e846f4f0b"></a>

<a id="canonical-b7b89d4c75e03428ffbcb3f0911cb0c124e5bc8020cab9e8ea6a338e155cfb30"></a>

## cipher_suites property — https_management.advertise_on_slo_vip.tls_config.custom_security / 505689d70009 / 4

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

<a id="canonical-b556b1faba25c47824033c992f167f7888a6f9c3dd54128ee4b57f691cda8d3f"></a>

<a id="canonical-c8c7c15d9914a7cd7bc8acabb1187e1d8af1afbe86872f3eca1232120393bd3f"></a>

## max_version property — https_management.advertise_on_slo_vip.tls_config.custom_security / 505689d70009 / 5

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

<a id="canonical-925f8b32a62af57168c15cba292df14b863dad907cddd914beeea15cd501c304"></a>

<a id="canonical-acc7bd9ac37a42b758bcd16ba212930be10e28252965f79a19c4c919c2f6bab3"></a>

## min_version property — https_management.advertise_on_slo_vip.tls_config.custom_security / 505689d70009 / 6

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

<a id="canonical-cac6d61ddba4cfa37c40c9e2de94f34bfba6c6b72834c3ce76f2b1afaeacd612"></a>

## Next pages — https_management.advertise_on_slo_vip.tls_config.custom_security / 505689d70009 / 7

- [https_management.advertise_on_slo_vip.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-dcdd49a2ef8dd8cd70a62742d664c720f5e0f1983fc540088b7d48b06c825cfa)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-2f1cbb9a82207b778a846cc32a4a6a6eb4779b301f3b5ab1000a04701b7bf525"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2ef8901d79045343b92a3b0383ad7f6a83009961e6ff24a3f43bce1126774394"></a>

## https_management.advertise_on_slo_vip.tls_config.default_security — https_management.advertise_on_slo_vip.tls_config.default_security / e2ef8e0fcd2c / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-9cb7d734443534e78f77f5035f885e992d26930de647b960b4a5a8542d04f2a7)
- [https_management.advertise_on_slo_vip.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-dcdd49a2ef8dd8cd70a62742d664c720f5e0f1983fc540088b7d48b06c825cfa)
- https_management.advertise_on_slo_vip.tls_config.default_security

<a id="canonical-a0122f48601f4dfab730272e19ebc92ab66ad8bb2fff8a18d84e8678e06bd72d"></a>

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

<a id="canonical-7553a145aad1cdc979a90ef3645d716a06d773532892b53898b599e8d7ac902e"></a>

## Direct properties — https_management.advertise_on_slo_vip.tls_config.default_security / e2ef8e0fcd2c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-18e57769a33550a7fabc7b32d093949af4b2900c2ef30fc508327648b0444c8e"></a>

## Next pages — https_management.advertise_on_slo_vip.tls_config.default_security / e2ef8e0fcd2c / 4

- [https_management.advertise_on_slo_vip.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-dcdd49a2ef8dd8cd70a62742d664c720f5e0f1983fc540088b7d48b06c825cfa)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-a0065a03a2857f7d48b0e64fc16c076f35b6577984a7eef47eeb514c830d9181"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c21a6b0c0e3578a51fb9d76ef783274aa59fbc2d5690d7ec39a22b418c694310"></a>

## https_management.advertise_on_slo_vip.tls_config.low_security — https_management.advertise_on_slo_vip.tls_config.low_security / 6cc441b45ae1 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-9cb7d734443534e78f77f5035f885e992d26930de647b960b4a5a8542d04f2a7)
- [https_management.advertise_on_slo_vip.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-dcdd49a2ef8dd8cd70a62742d664c720f5e0f1983fc540088b7d48b06c825cfa)
- https_management.advertise_on_slo_vip.tls_config.low_security

<a id="canonical-df48d07c5fb927c20654622535886ae2c0af08bdfa5a57ee756828a5dde978e5"></a>

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

<a id="canonical-b98bdd6d94ba153d59a419848fb37445b5f87077896519a81c226a885812fc41"></a>

## Direct properties — https_management.advertise_on_slo_vip.tls_config.low_security / 6cc441b45ae1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d2317635c1fdac5c3b7b68dd564dcdd95776bc7d52c8bcdca73630ffdfbc022c"></a>

## Next pages — https_management.advertise_on_slo_vip.tls_config.low_security / 6cc441b45ae1 / 4

- [https_management.advertise_on_slo_vip.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-dcdd49a2ef8dd8cd70a62742d664c720f5e0f1983fc540088b7d48b06c825cfa)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-6644d8b4098fdab6c6c186d82bbb6b38d284ba4f2f3675490438451de589dcd9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2b93a7a6ae69cc9e8dacbf602f7426cc65c8acd5c42ac73f44ec8f833353cbc6"></a>

## https_management.advertise_on_slo_vip.tls_config.medium_security — https_management.advertise_on_slo_vip.tls_config.medium_security / 8ecca4a96bb9 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-9cb7d734443534e78f77f5035f885e992d26930de647b960b4a5a8542d04f2a7)
- [https_management.advertise_on_slo_vip.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-dcdd49a2ef8dd8cd70a62742d664c720f5e0f1983fc540088b7d48b06c825cfa)
- https_management.advertise_on_slo_vip.tls_config.medium_security

<a id="canonical-d7b18ddd02eeb8b867d8a9a9474499eaec2b20a12cd1da661f34a01898b10df9"></a>

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

<a id="canonical-982f24ce771c71c0a467eb72aeb234044be1883e59d1397dac299d013034f880"></a>

## Direct properties — https_management.advertise_on_slo_vip.tls_config.medium_security / 8ecca4a96bb9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ed0f57db4765746f21628b627e13e7353eab7f3c2778b47ab751a90365f253b6"></a>

## Next pages — https_management.advertise_on_slo_vip.tls_config.medium_security / 8ecca4a96bb9 / 4

- [https_management.advertise_on_slo_vip.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-dcdd49a2ef8dd8cd70a62742d664c720f5e0f1983fc540088b7d48b06c825cfa)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-137dadd6204fcc9ee8516d19f54c9986ee16cac241901172042c9f9f35d39930"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-75fef4c6ed01aea7cfc0d0904271f288d5d95ca5fb48f1c700225923af9c05d1"></a>

## https_management.advertise_on_slo_vip.use_mtls — https_management.advertise_on_slo_vip.use_mtls / d70a28a18217 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-9cb7d734443534e78f77f5035f885e992d26930de647b960b4a5a8542d04f2a7)
- https_management.advertise_on_slo_vip.use_mtls

<a id="canonical-001171af2d1cddec42d24fe8a9da51be52e428376aad717462a3817c46e60c4b"></a>

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

<a id="canonical-21d9d685f9d250fb18e7bf9a35ef4cd46a5f1e3e0178f14d3d6843a00933cd43"></a>

## Direct properties — https_management.advertise_on_slo_vip.use_mtls / d70a28a18217 / 3

<a id="canonical-6745f143828dd9edb569211e63b50762f4bc02fd614c8bcd6dfa1ddabe2f20fd"></a>

<a id="canonical-64dac48423e7a898ab64b5ded340b61267157a9244c9b3a644428fe4b51dfc01"></a>

## client_certificate_optional property — https_management.advertise_on_slo_vip.use_mtls / d70a28a18217 / 4

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

- [crl](data-sources--nfv_service--reference--group-003.md#canonical-970b6d3d1b4b093f187ce371d54e574281e90f2f205b48324f764c5da3489bcc): complete subsection reference.

- [no_crl](data-sources--nfv_service--reference--group-003.md#canonical-52e3f02bbbe62dfbc221a84ba15e8e3e7b7112740a902494facbf3547bc0f1dc): complete subsection reference.

- [trusted_ca](data-sources--nfv_service--reference--group-003.md#canonical-d321af3db28a2742695f5c592fe66e62e3a99d307fac95a1db51cbb733d55622): complete subsection reference.

<a id="canonical-5ef247326fbb14b372cc4ffcebe7e73ece2c3c88fc14b90dead0fc92c37cd6d1"></a>

<a id="canonical-fdee4fefc43b9a14f0c2f549a92ab1866c56cf2c2f362a180df7745ac6652dc8"></a>

## trusted_ca_url property — https_management.advertise_on_slo_vip.use_mtls / d70a28a18217 / 5

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

- [xfcc_disabled](data-sources--nfv_service--reference--group-003.md#canonical-caefe37536839893f1848d12537e1ce2272ca13ff4d92cee50cd0dd3e955fff5): complete subsection reference.

- [xfcc_options](data-sources--nfv_service--reference--group-003.md#canonical-7dc6197a140ccdc36138491053b7b7798f919a7d1935e92992e656b987b38335): complete subsection reference.

<a id="canonical-cf3da0700909be7d4379cfacd633f3213dab4a33f0e46ac8936cf7d28a9acdb7"></a>

## Next pages — https_management.advertise_on_slo_vip.use_mtls / d70a28a18217 / 6

- [https_management.advertise_on_slo_vip.use_mtls.crl](data-sources--nfv_service--reference--group-003.md#canonical-970b6d3d1b4b093f187ce371d54e574281e90f2f205b48324f764c5da3489bcc)
- [https_management.advertise_on_slo_vip.use_mtls.no_crl](data-sources--nfv_service--reference--group-003.md#canonical-52e3f02bbbe62dfbc221a84ba15e8e3e7b7112740a902494facbf3547bc0f1dc)
- [https_management.advertise_on_slo_vip.use_mtls.trusted_ca](data-sources--nfv_service--reference--group-003.md#canonical-d321af3db28a2742695f5c592fe66e62e3a99d307fac95a1db51cbb733d55622)
- [https_management.advertise_on_slo_vip.use_mtls.xfcc_disabled](data-sources--nfv_service--reference--group-003.md#canonical-caefe37536839893f1848d12537e1ce2272ca13ff4d92cee50cd0dd3e955fff5)
- [https_management.advertise_on_slo_vip.use_mtls.xfcc_options](data-sources--nfv_service--reference--group-003.md#canonical-7dc6197a140ccdc36138491053b7b7798f919a7d1935e92992e656b987b38335)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-9cb7d734443534e78f77f5035f885e992d26930de647b960b4a5a8542d04f2a7)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-970b6d3d1b4b093f187ce371d54e574281e90f2f205b48324f764c5da3489bcc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c88cc98090a3f7f56c33b46eba6a0cbfb2874267776764b20bf6eb2be4b12a30"></a>

## https_management.advertise_on_slo_vip.use_mtls.crl — https_management.advertise_on_slo_vip.use_mtls.crl / 0eb5c90c6bbe / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-9cb7d734443534e78f77f5035f885e992d26930de647b960b4a5a8542d04f2a7)
- [https_management.advertise_on_slo_vip.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-137dadd6204fcc9ee8516d19f54c9986ee16cac241901172042c9f9f35d39930)
- https_management.advertise_on_slo_vip.use_mtls.crl

<a id="canonical-9945360391efe29d27a15861870fd576ec8e144b6df28c2fed6b2553e701b11d"></a>

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

<a id="canonical-0c10299bf9fbba56ce0ebfd59487d7aff0fe9da37ed7596a56f46e12599aaa50"></a>

## Direct properties — https_management.advertise_on_slo_vip.use_mtls.crl / 0eb5c90c6bbe / 3

<a id="canonical-38da4ff54d9e162628a8d2cddcd47fa778f6cd1d5cc551e6be888a67aa33a893"></a>

<a id="canonical-855c4b5525df280b1240b3db1a4b5ebace7fe5682aebade13b2cb5557432362c"></a>

## name property — https_management.advertise_on_slo_vip.use_mtls.crl / 0eb5c90c6bbe / 4

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

<a id="canonical-db2716585fd36806a5031c95dd35965bb22a38bae471016ce7da639197f81844"></a>

<a id="canonical-2daab8da84febe64a14e0721f821a1b3a6383da57ff6561752e99fefd23bba02"></a>

## namespace property — https_management.advertise_on_slo_vip.use_mtls.crl / 0eb5c90c6bbe / 5

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

<a id="canonical-c81e4a38c10f87f4df26d52f16c19bb5b7fb15f523f12058c6a944ac6af92b72"></a>

<a id="canonical-20796f23c1ad15d0140c2da5c6e1f4f6a58f3e70c68da7e0cd0576130962061c"></a>

## tenant property — https_management.advertise_on_slo_vip.use_mtls.crl / 0eb5c90c6bbe / 6

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

<a id="canonical-1a2bb9a61282471bd5cebc098ba7739e9c22ec3b11f5b1789e27b5da369849e3"></a>

## Next pages — https_management.advertise_on_slo_vip.use_mtls.crl / 0eb5c90c6bbe / 7

- [https_management.advertise_on_slo_vip.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-137dadd6204fcc9ee8516d19f54c9986ee16cac241901172042c9f9f35d39930)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-52e3f02bbbe62dfbc221a84ba15e8e3e7b7112740a902494facbf3547bc0f1dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-accead812f5c4525637a86b0495abb8a643c03b598e3ac18a01d221152e0cc69"></a>

## https_management.advertise_on_slo_vip.use_mtls.no_crl — https_management.advertise_on_slo_vip.use_mtls.no_crl / 4ed80b1f3027 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-9cb7d734443534e78f77f5035f885e992d26930de647b960b4a5a8542d04f2a7)
- [https_management.advertise_on_slo_vip.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-137dadd6204fcc9ee8516d19f54c9986ee16cac241901172042c9f9f35d39930)
- https_management.advertise_on_slo_vip.use_mtls.no_crl

<a id="canonical-85fdbe0771f84eead3becd612d202fa0552582e52665d9d6dd695d25c38d6755"></a>

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

<a id="canonical-473eb21d70a820a1457a8e2cbee04712748d7423bb33250b5067d81515c37fdd"></a>

## Direct properties — https_management.advertise_on_slo_vip.use_mtls.no_crl / 4ed80b1f3027 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8201829324d551a53bec40f87a1f779a97678d3bac1f9178e3e082269783dfe6"></a>

## Next pages — https_management.advertise_on_slo_vip.use_mtls.no_crl / 4ed80b1f3027 / 4

- [https_management.advertise_on_slo_vip.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-137dadd6204fcc9ee8516d19f54c9986ee16cac241901172042c9f9f35d39930)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-d321af3db28a2742695f5c592fe66e62e3a99d307fac95a1db51cbb733d55622"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-90d61b2dd82e8ea714649d94276bf42c889ef06ebd594466e233dfa43a328de7"></a>

## https_management.advertise_on_slo_vip.use_mtls.trusted_ca — https_management.advertise_on_slo_vip.use_mtls.trusted_ca / ef07e382c05d / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-9cb7d734443534e78f77f5035f885e992d26930de647b960b4a5a8542d04f2a7)
- [https_management.advertise_on_slo_vip.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-137dadd6204fcc9ee8516d19f54c9986ee16cac241901172042c9f9f35d39930)
- https_management.advertise_on_slo_vip.use_mtls.trusted_ca

<a id="canonical-da2ff59bbaf9028b80fd1dd08b9662fc2cffdb4027ee8c01a2cdd8a5e1b7c3fa"></a>

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

<a id="canonical-2918011f11c2f29f7724a31bf21d338998665cd4739c9b609ad06e95bb24fc1a"></a>

## Direct properties — https_management.advertise_on_slo_vip.use_mtls.trusted_ca / ef07e382c05d / 3

<a id="canonical-2984a992c1c57fa522827ad9afc0da4f7f13238ad7010fe19881de7eb12c46b5"></a>

<a id="canonical-8e763591df45f5354b0d7ac2a9d1843f9856410669b0d160306b2c8957d829dc"></a>

## name property — https_management.advertise_on_slo_vip.use_mtls.trusted_ca / ef07e382c05d / 4

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

<a id="canonical-8a38cec89c617696a58d8a09bbe5897784576e525c0184c05550f242c4fffeea"></a>

<a id="canonical-d906b27dfd75e310a1325822ca62477904c8d27cac9c42cf44b2ccff7a35475b"></a>

## namespace property — https_management.advertise_on_slo_vip.use_mtls.trusted_ca / ef07e382c05d / 5

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

<a id="canonical-9f18439f34fd54485855b1e66390b4e314dbeb525f2efd40378ab42933ef2df7"></a>

<a id="canonical-115ba926070edf0fa7634c058b616a6c0e782944b6406b162b2b9c0e9c307907"></a>

## tenant property — https_management.advertise_on_slo_vip.use_mtls.trusted_ca / ef07e382c05d / 6

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

<a id="canonical-508e902961c0c846f32198e10e0af8d5d9bb68c9fe909f3c97f7a16ace9c755e"></a>

## Next pages — https_management.advertise_on_slo_vip.use_mtls.trusted_ca / ef07e382c05d / 7

- [https_management.advertise_on_slo_vip.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-137dadd6204fcc9ee8516d19f54c9986ee16cac241901172042c9f9f35d39930)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-caefe37536839893f1848d12537e1ce2272ca13ff4d92cee50cd0dd3e955fff5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4b7a8ec9724251451da4032094f16b4845aebeb672aa35d6ef66fe0a16116ee"></a>

## https_management.advertise_on_slo_vip.use_mtls.xfcc_disabled — https_management.advertise_on_slo_vip.use_mtls.xfcc_disabled / b9044e75663b / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-9cb7d734443534e78f77f5035f885e992d26930de647b960b4a5a8542d04f2a7)
- [https_management.advertise_on_slo_vip.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-137dadd6204fcc9ee8516d19f54c9986ee16cac241901172042c9f9f35d39930)
- https_management.advertise_on_slo_vip.use_mtls.xfcc_disabled

<a id="canonical-b04d2b0603b4833e33682a14dbb0b443a5b580896fe4ff99f2b8abb5503aee72"></a>

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

<a id="canonical-f06cb049af5125407f048d794b5d830c5b7e4e8462e41c839ad27890736a3522"></a>

## Direct properties — https_management.advertise_on_slo_vip.use_mtls.xfcc_disabled / b9044e75663b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6a9af0aea35adafc68f2bbf6f1abf5d4e19f9d5b86c36de28eb554ee47fcdcbb"></a>

## Next pages — https_management.advertise_on_slo_vip.use_mtls.xfcc_disabled / b9044e75663b / 4

- [https_management.advertise_on_slo_vip.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-137dadd6204fcc9ee8516d19f54c9986ee16cac241901172042c9f9f35d39930)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-7dc6197a140ccdc36138491053b7b7798f919a7d1935e92992e656b987b38335"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c95dc8e6ba8ab9a08c7265eb66df5261809fea4a29be0bdb2e75b4e777053637"></a>

## https_management.advertise_on_slo_vip.use_mtls.xfcc_options — https_management.advertise_on_slo_vip.use_mtls.xfcc_options / 253b7a9f4aeb / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-9cb7d734443534e78f77f5035f885e992d26930de647b960b4a5a8542d04f2a7)
- [https_management.advertise_on_slo_vip.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-137dadd6204fcc9ee8516d19f54c9986ee16cac241901172042c9f9f35d39930)
- https_management.advertise_on_slo_vip.use_mtls.xfcc_options

<a id="canonical-ace47e59ac30efa008552a19eca6181b8decd56737a12f83b0f14e13343b1e6c"></a>

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

<a id="canonical-92138236673b990ea72c58126198d92325cb025b6ce30a7bbc1a074cbecf29b3"></a>

## Direct properties — https_management.advertise_on_slo_vip.use_mtls.xfcc_options / 253b7a9f4aeb / 3

<a id="canonical-f12018f1580a6326a62594baab14b5b918feba7f0a6efe8e91aebb68afb3908e"></a>

<a id="canonical-472a3e077d34ed5659f9bbbadfd4437f0114739ea966a8a0b09c4720d6bc9e54"></a>

## xfcc_header_elements property — https_management.advertise_on_slo_vip.use_mtls.xfcc_options / 253b7a9f4aeb / 4

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

<a id="canonical-478977d0164ad05f2d649bb2ea6e71dc746966c268297bc089157f9728f94c76"></a>

## Next pages — https_management.advertise_on_slo_vip.use_mtls.xfcc_options / 253b7a9f4aeb / 5

- [https_management.advertise_on_slo_vip.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-137dadd6204fcc9ee8516d19f54c9986ee16cac241901172042c9f9f35d39930)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-1ea6170c2e0093e64ef8c11cfeadbca923ba5e0edcba2b15fbd606e1aa0c37e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ebd983480b3678d7aee7af54330f12009869ec87fa143bd78e19491845305a75"></a>

## https_management.default_https_port — https_management.default_https_port / 11105ac5c43b / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- https_management.default_https_port

<a id="canonical-4db773ebbba7e348c67299c69f54f99ca9339fa846c77664a35ecb00d6611edd"></a>

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

<a id="canonical-ffecec11c503dcd72c2b4afc313614b699530336630a5f727823982c7bdd79e7"></a>

## Direct properties — https_management.default_https_port / 11105ac5c43b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2134ce9fdbdc06df29f883717de931a52feec373ca0a94807dca226f11541a2b"></a>

## Next pages — https_management.default_https_port / 11105ac5c43b / 4

- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-1ca81a9ddb9348c620d239771df702d0b475bf02453eb0a2a936c6227d30b60c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a433e50cf9f426cfc0371913f5f3c8c35dfe0d779457a99cba470cc2f4c5ad6"></a>

## palo_alto_fw_service — palo_alto_fw_service / 9c62bc0e19af / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- palo_alto_fw_service

<a id="canonical-a6ed4f07d1a8b4c1b3411155b197eb4eb87fa37ad09b24fab219c1fbf7327634"></a>

Type: `"single"`. Computed.

Palo Alto Networks VM-Series next-generation firewall configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ami_choice": "[\"pan_ami_bundle1\",\"pan_ami_bundle2\"]",
  "x-ves-oneof-field-panaroma_connection": "[\"disable_panaroma\",\"panorama_server\"]",
  "x-ves-oneof-field-setup_options": "[\"auto_setup\",\"ssh_key\"]"
}
```

<a id="canonical-fa371b8a200df7266b5d20ff3ac0122e8f5c2e8dbd4fc44bcdbf029064d17618"></a>

## Direct properties — palo_alto_fw_service / 9c62bc0e19af / 3

- [auto_setup](data-sources--nfv_service--reference--group-003.md#canonical-24b3b11dc5b28f9e76d04202bebd157c62efc08c6fb827859c34c47ec01e6335): complete subsection reference.

- [aws_tgw_site](data-sources--nfv_service--reference--group-003.md#canonical-cffcbe3b01213e701ec3aee0fe35dd2ebc67092e8070350cf96ccadc207e06b3): complete subsection reference.

- [disable_panaroma](data-sources--nfv_service--reference--group-003.md#canonical-63a882e23e2d08b2d6cbb8908d6145c88319b668b999ab68e15bf9ef58e53f90): complete subsection reference.

<a id="canonical-dc6036faf89765945cb9f15b8dca63ce0193de201461e237b0b79fcb47621f36"></a>

<a id="canonical-33ff90c3bee7482e8372a0f8803e40049e132a3fd6c3668cc707efbbcda99f71"></a>

## instance_type property — palo_alto_fw_service / 9c62bc0e19af / 4

Type: `"string"`. Computed.

\[Enum:
PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_2XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_4XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_LARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_2XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_4XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_12XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_LARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_2XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_4XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_LARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_2XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_4XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_8XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_LARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_2XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_4XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_9XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_18XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_LARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_2XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_4XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_9XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_18XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_R5\_2XLARGE\]
&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_XLARGE: m4.xlarge -
PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_2XLARGE: m4.2xlarge -
PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_4XLARGE: m4.4xlarge -
PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_LARGE: m5.large -
PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_XLARGE: m5.xlarge .. Possible values are
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_2XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_4XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_LARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_2XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_4XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_12XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_LARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_2XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_4XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_LARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_2XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_4XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_8XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_LARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_2XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_4XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_9XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_18XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_LARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_2XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_4XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_9XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_18XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_R5\_2XLARGE\`. Defaults to
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_XLARGE\`.

Upstream description:

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_XLARGE: m4.xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_2XLARGE: m4.2xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_4XLARGE: m4.4xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_LARGE: m5.large

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_XLARGE: m5.xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_2XLARGE: m5.2xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_4XLARGE: m5.4xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_12XLARGE: m5.12xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_LARGE: m5n.large

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_XLARGE: m5n.xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_2XLARGE: m5n.2xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_4XLARGE: m5n.4xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_LARGE: c4.large

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_XLARGE: c4.xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_2XLARGE: c4.2xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_4XLARGE: c4.4xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_8XLARGE: c4.8xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_LARGE: c5.large

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_XLARGE: c5.xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_2XLARGE: c5.2xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_4XLARGE: c5.4xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_9XLARGE: c5.9xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_18XLARGE: c5.18xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_LARGE: c5n.large

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_XLARGE: c5n.xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_2XLARGE: c5n.2xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_4XLARGE: c5n.4xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_9XLARGE: c5n.9xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_18XLARGE: c5n.18xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_R5\_2XLARGE: r5.2xlarge.

Receipt-pinned upstream constraints:

```json
{
  "default": "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M4_XLARGE",
  "enum": [
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M4_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M4_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M4_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_LARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_12XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5N_LARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5N_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5N_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5N_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C4_LARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C4_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C4_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C4_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C4_8XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_LARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_9XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_18XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_LARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_9XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_18XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_R5_2XLARGE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [pan_ami_bundle1](data-sources--nfv_service--reference--group-003.md#canonical-94ad902fe9e33d9a8053221f1d57d6cff31c0d6bf2621aa5a88efc94e75e92b8): complete subsection reference.

- [pan_ami_bundle2](data-sources--nfv_service--reference--group-003.md#canonical-1b2ca967d2e70c69a956481e81998cf7317f335b8cdb785724842cec32d83bf3): complete subsection reference.

- [panorama_server](data-sources--nfv_service--reference--group-004.md#canonical-e1a6b7b002ad6cbdb482c5e45ad67e792c6a1ccc39ae457d5fda0a4476aba660): complete subsection reference.

- [service_nodes](data-sources--nfv_service--reference--group-004.md#canonical-b844e50e5ec67e22eb6280a754515913fecc9f21e564c553eeac9822b82d6daa): complete subsection reference.

<a id="canonical-67ead7dd3d136ea024d893c92d324b9bef06b7ed53e699550d8d9940ef61b786"></a>

<a id="canonical-daae4ce75657373d8f02a77dd4118a867cebe66ad53efd17422e39b1fffef96c"></a>

## ssh_key property — palo_alto_fw_service / 9c62bc0e19af / 5

Type: `"string"`. Computed.

Exclusive with \[auto\_setup\] Setup Authorized Public SSH key. User will be able to SSH to the
vmseries nodes using its corresponding SSH private key.

Upstream description:

Exclusive with \[auto\_setup\] Setup Authorized Public SSH key. User will be able to SSH to the
vmseries nodes using its corresponding SSH private key.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-addb9d5c31a1cc4dd4756f48346a8f2d0379ce1b2e25ed4a4346d8ed9ff7e150"></a>

<a id="canonical-700f4e8349714fa1264afd664c9f6f2be31fe0610c318df626950041101c3a08"></a>

## tags property — palo_alto_fw_service / 9c62bc0e19af / 6

Type: `["map", "string"]`. Computed.

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

Upstream description:

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  }
}
```

<a id="canonical-ad22288ba823d38e7d0adb39ea2e9c88dc3e1558811b0bf2b45f76ac6101444e"></a>

<a id="canonical-b8682ad8b97de781009fd35fb81560efd6f4ca4a393449d5d2b7de49920a8fc0"></a>

## version property — palo_alto_fw_service / 9c62bc0e19af / 7

Type: `"string"`. Computed.

\[Enum: 11.0.0\] PAN VM-Series version. PAN-OS version. The only possible value is \`11.0.0\`.

Upstream description:

PAN-OS version.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "11.0.0"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"11.0.0\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"11.0.0\\\"]"
  }
}
```

<a id="canonical-eecc544ff37940a19087d7dffbf3b158d6220ca3b0ada3b8b0b9181a76f919e0"></a>

## Next pages — palo_alto_fw_service / 9c62bc0e19af / 8

- [palo_alto_fw_service.auto_setup](data-sources--nfv_service--reference--group-003.md#canonical-24b3b11dc5b28f9e76d04202bebd157c62efc08c6fb827859c34c47ec01e6335)
- [palo_alto_fw_service.aws_tgw_site](data-sources--nfv_service--reference--group-003.md#canonical-cffcbe3b01213e701ec3aee0fe35dd2ebc67092e8070350cf96ccadc207e06b3)
- [palo_alto_fw_service.disable_panaroma](data-sources--nfv_service--reference--group-003.md#canonical-63a882e23e2d08b2d6cbb8908d6145c88319b668b999ab68e15bf9ef58e53f90)
- [palo_alto_fw_service.pan_ami_bundle1](data-sources--nfv_service--reference--group-003.md#canonical-94ad902fe9e33d9a8053221f1d57d6cff31c0d6bf2621aa5a88efc94e75e92b8)
- [palo_alto_fw_service.pan_ami_bundle2](data-sources--nfv_service--reference--group-003.md#canonical-1b2ca967d2e70c69a956481e81998cf7317f335b8cdb785724842cec32d83bf3)
- [palo_alto_fw_service.panorama_server](data-sources--nfv_service--reference--group-004.md#canonical-e1a6b7b002ad6cbdb482c5e45ad67e792c6a1ccc39ae457d5fda0a4476aba660)
- [palo_alto_fw_service.service_nodes](data-sources--nfv_service--reference--group-004.md#canonical-b844e50e5ec67e22eb6280a754515913fecc9f21e564c553eeac9822b82d6daa)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-24b3b11dc5b28f9e76d04202bebd157c62efc08c6fb827859c34c47ec01e6335"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b61fc47b04f7d717e4d7e0c5c50f0fa21992771a95bbea198fe6bb3bf681522"></a>

## palo_alto_fw_service.auto_setup — palo_alto_fw_service.auto_setup / e0ed667f6573 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-1ca81a9ddb9348c620d239771df702d0b475bf02453eb0a2a936c6227d30b60c)
- palo_alto_fw_service.auto_setup

<a id="canonical-8514ad49c69ab7c221bd188e1f9ca8e7e02560cb6bac39051fab7bd7e6fd253f"></a>

Type: `"single"`. Computed.

For auto-setup, SSH public and pvt keys are needed. Using the given config user, SSH and API access
will be configured.

Upstream description:

For auto-setup, SSH public and pvt keys are needed. Using the given config user, SSH and API access
will be configured.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ssh_keys_choice": "[\"manual_ssh_keys\"]"
}
```

<a id="canonical-6b8854143e655f438f114cfa4bdfc5ba8613e3f112674ee32df196d42f61c5c5"></a>

## Direct properties — palo_alto_fw_service.auto_setup / e0ed667f6573 / 3

- [admin_password](data-sources--nfv_service--reference--group-003.md#canonical-8f7dfbff34ff7504b9ad5a00f03c4692ab708051341775f13cf022fc8ccefaab): complete subsection reference.

<a id="canonical-39c673559a1b792e4d438814cbad0c5304facda3422d19731e0e9da94ab8a39e"></a>

<a id="canonical-14ea2706db7c756f04ab6e3278cbb63e378603c68a78bbc4a5f55dec53f65f4d"></a>

## admin_username property — palo_alto_fw_service.auto_setup / e0ed667f6573 / 4

Type: `"string"`. Computed.

Firewall Admin Username. Firewall Admin Username.

Upstream description:

Firewall Admin Username.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [manual_ssh_keys](data-sources--nfv_service--reference--group-003.md#canonical-38afa562029bd3aac77b1d826f00df1f84a495b02cda75d14d0ba5a540cc8e59): complete subsection reference.

<a id="canonical-89563e798fdb3cb139f3f0dff5855f01d7f5f9f16f360065862c3809043ec700"></a>

## Next pages — palo_alto_fw_service.auto_setup / e0ed667f6573 / 5

- [palo_alto_fw_service.auto_setup.admin_password](data-sources--nfv_service--reference--group-003.md#canonical-8f7dfbff34ff7504b9ad5a00f03c4692ab708051341775f13cf022fc8ccefaab)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys](data-sources--nfv_service--reference--group-003.md#canonical-38afa562029bd3aac77b1d826f00df1f84a495b02cda75d14d0ba5a540cc8e59)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-1ca81a9ddb9348c620d239771df702d0b475bf02453eb0a2a936c6227d30b60c)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-8f7dfbff34ff7504b9ad5a00f03c4692ab708051341775f13cf022fc8ccefaab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f82d226aa11101adba8c7f7299e45b1afb5a001023db3da11b8b21fea2215d9d"></a>

## palo_alto_fw_service.auto_setup.admin_password — palo_alto_fw_service.auto_setup.admin_password / 5c8bc34e4512 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-1ca81a9ddb9348c620d239771df702d0b475bf02453eb0a2a936c6227d30b60c)
- [palo_alto_fw_service.auto_setup](data-sources--nfv_service--reference--group-003.md#canonical-24b3b11dc5b28f9e76d04202bebd157c62efc08c6fb827859c34c47ec01e6335)
- palo_alto_fw_service.auto_setup.admin_password

<a id="canonical-08bc387b290764c36e6fe7c3623b1a4a728a5a0574d23c40ffedfdce9a195b3c"></a>

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

<a id="canonical-264f386451fa1ce7bed72af6a0c4fe459dd97f038a83e723564f07e038bd689c"></a>

## Direct properties — palo_alto_fw_service.auto_setup.admin_password / 5c8bc34e4512 / 3

- [blindfold_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-f5bc5209e1f2716ebe1a11c468b98c9733b074a00e095081c3930ea5ad80fa2b): complete subsection reference.

- [clear_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-ec58beeac0294d0b75b93a1bef2649dc34c6c072d199a2c7163fdae2b2c873f2): complete subsection reference.

<a id="canonical-9af6ca385cc140276397a56bf7ca38b08688b0a9febbae48ab386633d90ad748"></a>

## Next pages — palo_alto_fw_service.auto_setup.admin_password / 5c8bc34e4512 / 4

- [palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-f5bc5209e1f2716ebe1a11c468b98c9733b074a00e095081c3930ea5ad80fa2b)
- [palo_alto_fw_service.auto_setup.admin_password.clear_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-ec58beeac0294d0b75b93a1bef2649dc34c6c072d199a2c7163fdae2b2c873f2)
- [palo_alto_fw_service.auto_setup](data-sources--nfv_service--reference--group-003.md#canonical-24b3b11dc5b28f9e76d04202bebd157c62efc08c6fb827859c34c47ec01e6335)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-f5bc5209e1f2716ebe1a11c468b98c9733b074a00e095081c3930ea5ad80fa2b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0348679cb94b527cf405c3700bd4c99d94e30267a7f807533ad5631e8e5aad61"></a>

## palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info — palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info / bd4c312bcfd5 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-1ca81a9ddb9348c620d239771df702d0b475bf02453eb0a2a936c6227d30b60c)
- [palo_alto_fw_service.auto_setup](data-sources--nfv_service--reference--group-003.md#canonical-24b3b11dc5b28f9e76d04202bebd157c62efc08c6fb827859c34c47ec01e6335)
- [palo_alto_fw_service.auto_setup.admin_password](data-sources--nfv_service--reference--group-003.md#canonical-8f7dfbff34ff7504b9ad5a00f03c4692ab708051341775f13cf022fc8ccefaab)
- palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info

<a id="canonical-4b6d3d34f44bc2e3bd4ccb2183dc9d7311dce62890145eb395b7efbf525d54c8"></a>

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

<a id="canonical-04768916461381e7b61024341e0a9e217fa054579d970023199b30ce19398e6b"></a>

## Direct properties — palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info / bd4c312bcfd5 / 3

<a id="canonical-6836682ab9541bb6530ee2f81f762168b92a4284e95323319c1bea16b96b307e"></a>

<a id="canonical-f0ff57183ff43dbe5e6f62dcda2227293d8bc023dd29a9878c1acfe2c09e5179"></a>

## decryption_provider property — palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info / bd4c312bcfd5 / 4

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

<a id="canonical-a4a1a5141b3978261e743d88bffca3a8d64ac072be672731135dc0e3f5c2291c"></a>

<a id="canonical-0cdb125d858086cd06ede3fdbdb65c4edd53b90ddbc52c9f059f86db1595214a"></a>

## location property — palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info / bd4c312bcfd5 / 5

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

<a id="canonical-953cf2ac97613d6c194151a5016cb90d3a5a9a291dd8ecd6ef428d1559f38edf"></a>

<a id="canonical-608b04a049279c7ab6e23df5107e8fefa978f8c40880ce2093f50deae246462c"></a>

## store_provider property — palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info / bd4c312bcfd5 / 6

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

<a id="canonical-1b07587dee011f95df565513b075be7160229a4401ac99f00801f6f4a8951167"></a>

## Next pages — palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info / bd4c312bcfd5 / 7

- [palo_alto_fw_service.auto_setup.admin_password](data-sources--nfv_service--reference--group-003.md#canonical-8f7dfbff34ff7504b9ad5a00f03c4692ab708051341775f13cf022fc8ccefaab)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-ec58beeac0294d0b75b93a1bef2649dc34c6c072d199a2c7163fdae2b2c873f2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9f903c9189c19265f16fb496bbebfcaa950876a13c82ca5d522ee9ca11b037e7"></a>

## palo_alto_fw_service.auto_setup.admin_password.clear_secret_info — palo_alto_fw_service.auto_setup.admin_password.clear_secret_info / fb0f4d719f1a / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-1ca81a9ddb9348c620d239771df702d0b475bf02453eb0a2a936c6227d30b60c)
- [palo_alto_fw_service.auto_setup](data-sources--nfv_service--reference--group-003.md#canonical-24b3b11dc5b28f9e76d04202bebd157c62efc08c6fb827859c34c47ec01e6335)
- [palo_alto_fw_service.auto_setup.admin_password](data-sources--nfv_service--reference--group-003.md#canonical-8f7dfbff34ff7504b9ad5a00f03c4692ab708051341775f13cf022fc8ccefaab)
- palo_alto_fw_service.auto_setup.admin_password.clear_secret_info

<a id="canonical-dd28916f5379f7ed2fb92cf6e7aeb107527fc2e7e97be7c83b6bcabda635f745"></a>

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

<a id="canonical-7b17b2dddb9c134a99dd811dc47745b72b68f45d5b67001a0da7f311d257ec4e"></a>

## Direct properties — palo_alto_fw_service.auto_setup.admin_password.clear_secret_info / fb0f4d719f1a / 3

<a id="canonical-6086646350fd9991ca9863ac867659adb593d04c9e0410dc1ec439966e6e2e65"></a>

<a id="canonical-6f8454261a5b5faaf7c6bb95049f9c59f7a7e51b3cf448a5b4c75c91241a5e26"></a>

## provider_ref property — palo_alto_fw_service.auto_setup.admin_password.clear_secret_info / fb0f4d719f1a / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-ee1fd1b02ea4f5fb3ebbac324eb9ae1724ee92b5fa002fbfb81536c44ad35584"></a>

<a id="canonical-21d971dd76452629ae81f0075eb70a5821189ceca58d1a49012fd84a6153bb0c"></a>

## url property — palo_alto_fw_service.auto_setup.admin_password.clear_secret_info / fb0f4d719f1a / 5

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

<a id="canonical-665fce9fc58a678df680a024bbe842155366cef55a0527da2c63d7163d3f4226"></a>

## Next pages — palo_alto_fw_service.auto_setup.admin_password.clear_secret_info / fb0f4d719f1a / 6

- [palo_alto_fw_service.auto_setup.admin_password](data-sources--nfv_service--reference--group-003.md#canonical-8f7dfbff34ff7504b9ad5a00f03c4692ab708051341775f13cf022fc8ccefaab)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-38afa562029bd3aac77b1d826f00df1f84a495b02cda75d14d0ba5a540cc8e59"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2e9737d1f300c3e5a2026507c6ee57ab473878f140c6565dc09ea7674476293f"></a>

## palo_alto_fw_service.auto_setup.manual_ssh_keys — palo_alto_fw_service.auto_setup.manual_ssh_keys / fe019dcfb81e / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-1ca81a9ddb9348c620d239771df702d0b475bf02453eb0a2a936c6227d30b60c)
- [palo_alto_fw_service.auto_setup](data-sources--nfv_service--reference--group-003.md#canonical-24b3b11dc5b28f9e76d04202bebd157c62efc08c6fb827859c34c47ec01e6335)
- palo_alto_fw_service.auto_setup.manual_ssh_keys

<a id="canonical-8d10c552b11c1f65387f1e00ea58de1807f842e41991be40d30aabad007167c7"></a>

Type: `"single"`. Computed.

SSH Key includes both public and private key.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-e4aedb2e964751a5930cbd2c11f465d2e011e4baa4b7dec76ef10afde625d793"></a>

## Direct properties — palo_alto_fw_service.auto_setup.manual_ssh_keys / fe019dcfb81e / 3

- [private_key](data-sources--nfv_service--reference--group-003.md#canonical-7d40e3227a5c45eceb1586b87b2cc99bd54a0f81269f484bdd9eb985e2d10edd): complete subsection reference.

<a id="canonical-f7ccc85cc22f7b67536bd9d9f07e371d715b8370b3a681a7838ee11930c938e1"></a>

<a id="canonical-9440f689857a48dd4156b2d9e07a80a962b9af162762df30deeb01899d8a90a9"></a>

## public_key property — palo_alto_fw_service.auto_setup.manual_ssh_keys / fe019dcfb81e / 4

Type: `"string"`. Computed.

Authorized Public SSH key which will be programmed on the node.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 5242880,
      "min": 100
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "pem",
    "maxLength": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^-----BEGIN (RSA |EC |)?(PRIVATE |PUBLIC )?KEY-----\\n.*\\n-----END (RSA |EC |)?(PRIVATE |PUBLIC )?KEY-----$",
    "validation": {
      "standard": "PEM"
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
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-8296ec93bfa68623fa3ad877c28afad7206930c9a852051fd8fc86ccd1437fcb"></a>

## Next pages — palo_alto_fw_service.auto_setup.manual_ssh_keys / fe019dcfb81e / 5

- [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key](data-sources--nfv_service--reference--group-003.md#canonical-7d40e3227a5c45eceb1586b87b2cc99bd54a0f81269f484bdd9eb985e2d10edd)
- [palo_alto_fw_service.auto_setup](data-sources--nfv_service--reference--group-003.md#canonical-24b3b11dc5b28f9e76d04202bebd157c62efc08c6fb827859c34c47ec01e6335)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-7d40e3227a5c45eceb1586b87b2cc99bd54a0f81269f484bdd9eb985e2d10edd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7dfde7a5b259714a6d72e2e9e60e7f0f23e3701bb7a8d1aefc09967b6e9fd5bd"></a>

## palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key — palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key / 260026167294 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-1ca81a9ddb9348c620d239771df702d0b475bf02453eb0a2a936c6227d30b60c)
- [palo_alto_fw_service.auto_setup](data-sources--nfv_service--reference--group-003.md#canonical-24b3b11dc5b28f9e76d04202bebd157c62efc08c6fb827859c34c47ec01e6335)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys](data-sources--nfv_service--reference--group-003.md#canonical-38afa562029bd3aac77b1d826f00df1f84a495b02cda75d14d0ba5a540cc8e59)
- palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key

<a id="canonical-9c5a492f241810769925bc588d3575a9c335198a28cabb1febfb4bde73b49900"></a>

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

<a id="canonical-459d7ab920fe0981625af1d1de47e95bfd8f4faac79f2e781ca71aabd7f9c36d"></a>

## Direct properties — palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key / 260026167294 / 3

- [blindfold_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-820e073aebe76e6ed72f4eca8674a6d50ad51851b63c6aeb67cae5b021e2f398): complete subsection reference.

- [clear_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-525220e858cdcac2491cdbd75e4bcebf0ef59c0ccc2050300fc7b4f26b5263e9): complete subsection reference.

<a id="canonical-05185136261c7cd3d3514f99872600e67c0e373c484cccea6a31f1b11e162e6f"></a>

## Next pages — palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key / 260026167294 / 4

- [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-820e073aebe76e6ed72f4eca8674a6d50ad51851b63c6aeb67cae5b021e2f398)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-525220e858cdcac2491cdbd75e4bcebf0ef59c0ccc2050300fc7b4f26b5263e9)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys](data-sources--nfv_service--reference--group-003.md#canonical-38afa562029bd3aac77b1d826f00df1f84a495b02cda75d14d0ba5a540cc8e59)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-820e073aebe76e6ed72f4eca8674a6d50ad51851b63c6aeb67cae5b021e2f398"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a7f241cc1eed4ae082e4911d2c75b1e277104256a0a876c8ed1c88b6aead276"></a>

## palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info — palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_inf / 83c4649a0060 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-1ca81a9ddb9348c620d239771df702d0b475bf02453eb0a2a936c6227d30b60c)
- [palo_alto_fw_service.auto_setup](data-sources--nfv_service--reference--group-003.md#canonical-24b3b11dc5b28f9e76d04202bebd157c62efc08c6fb827859c34c47ec01e6335)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys](data-sources--nfv_service--reference--group-003.md#canonical-38afa562029bd3aac77b1d826f00df1f84a495b02cda75d14d0ba5a540cc8e59)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key](data-sources--nfv_service--reference--group-003.md#canonical-7d40e3227a5c45eceb1586b87b2cc99bd54a0f81269f484bdd9eb985e2d10edd)
- palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info

<a id="canonical-1d468e3937ce025437880db6a737b105d58b87c3d02a40d3d4b2d6a4468e9369"></a>

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

<a id="canonical-f046511b6e182cb355edb78c01cd448ccebd1a72d07bef4f76ab71467c5f42ab"></a>

## Direct properties — palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_inf / 83c4649a0060 / 3

<a id="canonical-bfe7248dc6a8be979cfc4207b7e08484e93c84865dc731d92df58194fd806592"></a>

<a id="canonical-1ae44531f313a81022f6eabd94d0b7392e1347b860a1a18508c58712779b067f"></a>

## decryption_provider property — palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_inf / 83c4649a0060 / 4

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

<a id="canonical-77bdeccdbef7cf07fe2135838faf0690539a9e79f4ce13546119fb5a8c5e0e68"></a>

<a id="canonical-c50200da6c714793a3f19f241d4ca067abc8d2ed886464e96f526656a90eae9c"></a>

## location property — palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_inf / 83c4649a0060 / 5

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

<a id="canonical-af7034f6b05cb7264ba3b359650e17b8f85b7a87b51c14db727f28f479c6bb2f"></a>

<a id="canonical-31ec42e7645ffd16366acb77a67f5f38842d1bdf0471acbac9c0c23d3107407f"></a>

## store_provider property — palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_inf / 83c4649a0060 / 6

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

<a id="canonical-63c5bcf05c7de05baccc0f667c8591160a5fbab6624fac7a6d923d5cdfb2e2c9"></a>

## Next pages — palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_inf / 83c4649a0060 / 7

- [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key](data-sources--nfv_service--reference--group-003.md#canonical-7d40e3227a5c45eceb1586b87b2cc99bd54a0f81269f484bdd9eb985e2d10edd)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-525220e858cdcac2491cdbd75e4bcebf0ef59c0ccc2050300fc7b4f26b5263e9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-661b88068ecc4c1f04e3e74cd28cfc8a08fdef7e3c921b443dfbcb5ece7c1fc5"></a>

## palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info — palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info / 2f9caa5b78b5 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-1ca81a9ddb9348c620d239771df702d0b475bf02453eb0a2a936c6227d30b60c)
- [palo_alto_fw_service.auto_setup](data-sources--nfv_service--reference--group-003.md#canonical-24b3b11dc5b28f9e76d04202bebd157c62efc08c6fb827859c34c47ec01e6335)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys](data-sources--nfv_service--reference--group-003.md#canonical-38afa562029bd3aac77b1d826f00df1f84a495b02cda75d14d0ba5a540cc8e59)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key](data-sources--nfv_service--reference--group-003.md#canonical-7d40e3227a5c45eceb1586b87b2cc99bd54a0f81269f484bdd9eb985e2d10edd)
- palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info

<a id="canonical-735cf0fd70083c7ab03e8fe5c69e9c84e2e6d03ce9ac6d8d8ff409b1b05f9e14"></a>

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

<a id="canonical-c7b8b53c991cf946579ab0c5c2622f1c6e4b87ad5b793f6fc2c433bb12522309"></a>

## Direct properties — palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info / 2f9caa5b78b5 / 3

<a id="canonical-39749afed7998d06b092e362df07dbb32669a8110ee9a02fc27bfe4b9371c3b9"></a>

<a id="canonical-65e6ab11bebab61ea770631afae8a1063cc1c70ba4c97306a8b0c4e5cb10e585"></a>

## provider_ref property — palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info / 2f9caa5b78b5 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-8f0ffa6b8b3cc48cb7305e3ddb584a6ec45ff747d8a4de9a3c314596d8965daf"></a>

<a id="canonical-1dec8ef97193a993b2b860da6fa9bdaa8b5f91857ccff5402657273f96f1ef9a"></a>

## url property — palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info / 2f9caa5b78b5 / 5

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

<a id="canonical-dabb6759388d03e53b0e5d17418d4ffd5dd0ec392c5790eb979c452d9c3e4ab9"></a>

## Next pages — palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info / 2f9caa5b78b5 / 6

- [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key](data-sources--nfv_service--reference--group-003.md#canonical-7d40e3227a5c45eceb1586b87b2cc99bd54a0f81269f484bdd9eb985e2d10edd)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-cffcbe3b01213e701ec3aee0fe35dd2ebc67092e8070350cf96ccadc207e06b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa08120aabe22d21f0098107ba7ad58b41e95e5773c0f80e1d0456f3e39e2658"></a>

## palo_alto_fw_service.aws_tgw_site — palo_alto_fw_service.aws_tgw_site / cedaf1b6e1e0 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-1ca81a9ddb9348c620d239771df702d0b475bf02453eb0a2a936c6227d30b60c)
- palo_alto_fw_service.aws_tgw_site

<a id="canonical-42ada5d9d2f2437c8bccebd5d4087b692a06344f534861244fd447c2b8e3e009"></a>

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

<a id="canonical-b265be7ba5f4533dc827491cd4b9382c0c844262e2dcde827c415675a22721b0"></a>

## Direct properties — palo_alto_fw_service.aws_tgw_site / cedaf1b6e1e0 / 3

<a id="canonical-8581643397a5cc4e48374a17361b3142b72ddbcda3da87f391788d5cf43f90ec"></a>

<a id="canonical-583ee77781d896e29d0c3816a69aca50206a70434101edb763f482e003458263"></a>

## name property — palo_alto_fw_service.aws_tgw_site / cedaf1b6e1e0 / 4

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

<a id="canonical-ef87214a53c6a1e92671d3deace6f7b7d63dd493c9beade6fa8155dc3c886ba4"></a>

<a id="canonical-71bf9be62c5ed9495bf9f44b6501ee8424dd85ffd95328d3a7360f90fd9d45fe"></a>

## namespace property — palo_alto_fw_service.aws_tgw_site / cedaf1b6e1e0 / 5

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

<a id="canonical-7b106d42392342088fd0719c7266e48f6135a91a8af0f66e3ef3f3808c3e0c6a"></a>

<a id="canonical-32c6d310d632ab24ea8e67cf5ff4a3aef1d77036fb6a9d11f7cd7dde1f27e3b1"></a>

## tenant property — palo_alto_fw_service.aws_tgw_site / cedaf1b6e1e0 / 6

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

<a id="canonical-de7cecccb61c2bae46ba69074b42bd1ad3baad2fc1f2d35575135cf5e4c3e32b"></a>

## Next pages — palo_alto_fw_service.aws_tgw_site / cedaf1b6e1e0 / 7

- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-1ca81a9ddb9348c620d239771df702d0b475bf02453eb0a2a936c6227d30b60c)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-63a882e23e2d08b2d6cbb8908d6145c88319b668b999ab68e15bf9ef58e53f90"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3539ecb9b4131a3b0fd75e98ba0f5d1389628f77d14668f835482aa995fe893f"></a>

## palo_alto_fw_service.disable_panaroma — palo_alto_fw_service.disable_panaroma / b069e221ab70 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-1ca81a9ddb9348c620d239771df702d0b475bf02453eb0a2a936c6227d30b60c)
- palo_alto_fw_service.disable_panaroma

<a id="canonical-e92d000f62de44f6fa7fb62dc087ae08c990e5f264328b8607b7631b5cfb504a"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable panaroma.

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

<a id="canonical-31a9367d839e062968ff4cf0ce4cbca5e64885cdf4b7f7fc5d294b65096e6994"></a>

## Direct properties — palo_alto_fw_service.disable_panaroma / b069e221ab70 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4031403e85dd5f01f51334bae3c2bf0bfb1edbab8e34d063ef402c62cf657f58"></a>

## Next pages — palo_alto_fw_service.disable_panaroma / b069e221ab70 / 4

- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-1ca81a9ddb9348c620d239771df702d0b475bf02453eb0a2a936c6227d30b60c)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-94ad902fe9e33d9a8053221f1d57d6cff31c0d6bf2621aa5a88efc94e75e92b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-11a60b5e46d8635bdb51626be05a23ae9c94f085b1fbe841b95351041258a7ba"></a>

## palo_alto_fw_service.pan_ami_bundle1 — palo_alto_fw_service.pan_ami_bundle1 / 920527fe79b1 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-1ca81a9ddb9348c620d239771df702d0b475bf02453eb0a2a936c6227d30b60c)
- palo_alto_fw_service.pan_ami_bundle1

<a id="canonical-3c089318ac95fea53a2af857196c2a909380ba950b8400927ffe840fc6e5ab1f"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for pan ami bundle1.

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

<a id="canonical-1eafc1dd1faf4d5b26173469a0450c07199f2ca0484bf3ae577cfb2e25ebbf2a"></a>

## Direct properties — palo_alto_fw_service.pan_ami_bundle1 / 920527fe79b1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-58d375f2e5df4f503dd59a689e2c6ce0c0698b6be8d89412aa8cbad269d9cce4"></a>

## Next pages — palo_alto_fw_service.pan_ami_bundle1 / 920527fe79b1 / 4

- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-1ca81a9ddb9348c620d239771df702d0b475bf02453eb0a2a936c6227d30b60c)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-1b2ca967d2e70c69a956481e81998cf7317f335b8cdb785724842cec32d83bf3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0c62a07ba8a4e79ba9ddf27b00eb1ae293aae24347d5c22eb630d85caf0e4da"></a>

## palo_alto_fw_service.pan_ami_bundle2 — palo_alto_fw_service.pan_ami_bundle2 / 93eea51f1fe6 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-1ca81a9ddb9348c620d239771df702d0b475bf02453eb0a2a936c6227d30b60c)
- palo_alto_fw_service.pan_ami_bundle2

<a id="canonical-ab5a02ed414e960cb3086b94c5129b4fe732496078b8c4aa0190e55818b82862"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for pan ami bundle2.

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

<a id="canonical-56c225fc078bf6e77d8624d2735100e51dc15fee451423286188eec892ec5492"></a>

## Direct properties — palo_alto_fw_service.pan_ami_bundle2 / 93eea51f1fe6 / 3

This is an empty object or choice marker. It has no direct properties.
