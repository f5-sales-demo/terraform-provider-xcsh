---
page_title: "xcsh_nfv_service reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nfv_service reference."
---

# xcsh_nfv_service reference

<a id="canonical-d5bd9d147bf7aba3c20606693408e3dc8e195add536b7a2b8a4d947ee56903a3"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_disabled / 63d7369f7dea / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0f2d3c699c4b58d737cfe634d6f692682021e723211aa21106fd4c8cbce50f7d"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_disabled / 63d7369f7dea / 4

- [https_management.advertise_on_slo_internet_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-a8cd513a1608b86e1f196b42d35787d058f4210ab856187e96cf80249c2cffd6)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-a69d6df5da51dc13c9779e3e9d6ad788951c1a687499a422603c3abc63c8cd73"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f24fd8841d5dcc6ed04e9dbb1520abfa2a0e180c3f46a88d68bb8e78a1e5888"></a>

## https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options — https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options / eaaccb5322d8 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-cd7a014a77a7464ee01bdc9918de6c755195ebcf97fd000b19523ae733884b9c)
- [https_management.advertise_on_slo_internet_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-a8cd513a1608b86e1f196b42d35787d058f4210ab856187e96cf80249c2cffd6)
- https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options

<a id="canonical-e696d1ac2cac47dc8c900cc9b0cf61631e6c4b670d5ce55a47356ab1242ca61f"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("xfcc_header_elements")}
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
xfcc_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-bcb5a16beae10b44ae97a01afde5e58946e4d9ff5cff97ea44a9fff48732221c"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options / eaaccb5322d8 / 3

<a id="canonical-94c244e360da27714c01cb855c0af874a43eb2637739e167c313f2bf4e491a45"></a>

<a id="canonical-1979cff6977d11ef259cb9dd358d484bd2588fb53bb351394de3e60f3670551e"></a>

## xfcc_header_elements property — https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options / eaaccb5322d8 / 4

Type: `["list", "string"]`. Optional.

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

<a id="canonical-229e6a2ca8328ede77523c1037bc4feafb0be4c1f46e83efb9bcec66eb8c2a70"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options / eaaccb5322d8 / 5

- [https_management.advertise_on_slo_internet_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-a8cd513a1608b86e1f196b42d35787d058f4210ab856187e96cf80249c2cffd6)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-30022b387cd576f0c782dfcba45fea9f6211f83d534f947e7512e992c2608a46"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9a318c7a007cfa8333a1dfa6b8ebdbe5dff46db8caa7751c43ccd54f916c3b08"></a>

## https_management.advertise_on_slo_sli — https_management.advertise_on_slo_sli / 67e42c19d2bd / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- https_management.advertise_on_slo_sli

<a id="canonical-6ad2c4c38ca51dbdc0f2f54da0cc8dba14054e49bf85b81b7eaf5854bc31192c"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for advertise on slo sli.

Upstream description:

Inline TLS parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("tls_certificates"),
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
advertise_on_slo_sli {
  # Configure direct properties listed below.
}
```

<a id="canonical-8fe9da0ca84dfeb23558b45f9235c4e1e0cc618feb900e42f591c04b09349b9c"></a>

## Direct properties — https_management.advertise_on_slo_sli / 67e42c19d2bd / 3

- [no_mtls](resources--nfv_service--reference--group-003.md#canonical-5cdfd66215ba80c5133f3e2b982639e5195ab84ed4f522eae69843723497a1c8): complete subsection reference.

- [tls_certificates](resources--nfv_service--reference--group-003.md#canonical-9558627558b2ed3fb58a624e041ddbf1bb7ad9ccbcf3f1faaa0c30e60b58d257): complete subsection reference.

- [tls_config](resources--nfv_service--reference--group-003.md#canonical-e8f3a8275dfae5058f691263dfcf487950fd2ec4a4bcbf49a9f81cfe3d94e84d): complete subsection reference.

- [use_mtls](resources--nfv_service--reference--group-003.md#canonical-1257604a48a6882e78dae67503e07dec0e7ef0ab20383466d80273b5c231bff4): complete subsection reference.

<a id="canonical-cafe815597827e5c3ccf412a8bdfe68cf27513951e9aee72da1965ba2e44955e"></a>

## Next pages — https_management.advertise_on_slo_sli / 67e42c19d2bd / 4

- [https_management.advertise_on_slo_sli.no_mtls](resources--nfv_service--reference--group-003.md#canonical-5cdfd66215ba80c5133f3e2b982639e5195ab84ed4f522eae69843723497a1c8)
- [https_management.advertise_on_slo_sli.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-9558627558b2ed3fb58a624e041ddbf1bb7ad9ccbcf3f1faaa0c30e60b58d257)
- [https_management.advertise_on_slo_sli.tls_config](resources--nfv_service--reference--group-003.md#canonical-e8f3a8275dfae5058f691263dfcf487950fd2ec4a4bcbf49a9f81cfe3d94e84d)
- [https_management.advertise_on_slo_sli.use_mtls](resources--nfv_service--reference--group-003.md#canonical-1257604a48a6882e78dae67503e07dec0e7ef0ab20383466d80273b5c231bff4)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-5cdfd66215ba80c5133f3e2b982639e5195ab84ed4f522eae69843723497a1c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dee7dbe85c61abce7eff0eaea2edba8cf8af8faec7b0fc2da40d49c16c795194"></a>

## https_management.advertise_on_slo_sli.no_mtls — https_management.advertise_on_slo_sli.no_mtls / a66f4d866813 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-30022b387cd576f0c782dfcba45fea9f6211f83d534f947e7512e992c2608a46)
- https_management.advertise_on_slo_sli.no_mtls

<a id="canonical-3a3c4d5c0ae03661d7148af7eeb43e7830b0c6e2a369632e3b7c02940a8e9ecb"></a>

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

<a id="canonical-ffa6c87d250361608b53486bb57e3527f01fd338644bceedda37bdc2aca5a91b"></a>

## Direct properties — https_management.advertise_on_slo_sli.no_mtls / a66f4d866813 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7bf4380b2dfc71e7189b1541e942b6d8f2e3c6662bb2ede3a5e5a6e4916cb009"></a>

## Next pages — https_management.advertise_on_slo_sli.no_mtls / a66f4d866813 / 4

- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-30022b387cd576f0c782dfcba45fea9f6211f83d534f947e7512e992c2608a46)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-9558627558b2ed3fb58a624e041ddbf1bb7ad9ccbcf3f1faaa0c30e60b58d257"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b3fc65806fbd99454d856ba8ee3109379bc7ee69e5f6ff9946d760350d184ab"></a>

## https_management.advertise_on_slo_sli.tls_certificates — https_management.advertise_on_slo_sli.tls_certificates / d29e17749082 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-30022b387cd576f0c782dfcba45fea9f6211f83d534f947e7512e992c2608a46)
- https_management.advertise_on_slo_sli.tls_certificates

<a id="canonical-aa796af23ed9aa6f385192234d04a302db2615a7c8c2b0d3034d9fa16b89d461"></a>

Type: `"object"`. list nested block, Optional.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Upstream description:

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("certificate_url"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "disable_ocsp_stapling"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "use_system_defaults"),
  validators.ConflictingListObjectAttributes("disable_ocsp_stapling",
    "use_system_defaults")}
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

Terraform syntax:

```terraform
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-9c5e90cdbf9bd89dcdeb0045d766bb771241c251d2222db47838b6b0d28e2c65"></a>

## Direct properties — https_management.advertise_on_slo_sli.tls_certificates / d29e17749082 / 3

<a id="canonical-149d06d940e98998393c8945a35d42e0c3e4283fc6f129bf406517393ca73050"></a>

<a id="canonical-998e0b10a3158f5dd2253be650158d76071b25b2ee78f0b3cde3a5eb51bb4a04"></a>

## certificate_url property — https_management.advertise_on_slo_sli.tls_certificates / d29e17749082 / 4

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

- [custom_hash_algorithms](resources--nfv_service--reference--group-003.md#canonical-b3e56f03abd4532fbee6535661d6dac7ea85f318ecb2c3be6fb3c524e71b7f5e): complete subsection reference.

<a id="canonical-448a52b525745f5b42d525f68152fac0212afa96f41d302c7f979da21456ef87"></a>

<a id="canonical-f6bec9a5603793c66e931c85745a045cb423dfb6509c9e5fc605477697dbb9f8"></a>

## description_spec property — https_management.advertise_on_slo_sli.tls_certificates / d29e17749082 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--nfv_service--reference--group-003.md#canonical-6c73f4f34d68de0d1881aa020acde9578eb9b758f2aab0fcd65359f5b1e4e154): complete subsection reference.

- [private_key](resources--nfv_service--reference--group-003.md#canonical-a865d8eee24076eea9f3c8bbdd41aa47e9b2e649df0d523a130e4ab5ec4452c0): complete subsection reference.

- [use_system_defaults](resources--nfv_service--reference--group-003.md#canonical-4e1d7a69bc87776186b92e588924fa1a0b0152671d5c4d976d5210f07f1a7a7e): complete subsection reference.

<a id="canonical-833aa89e27f732f5bf693f901151a19346f0b9ca731181bb7c106cadf6d57065"></a>

## Next pages — https_management.advertise_on_slo_sli.tls_certificates / d29e17749082 / 6

- [https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms](resources--nfv_service--reference--group-003.md#canonical-b3e56f03abd4532fbee6535661d6dac7ea85f318ecb2c3be6fb3c524e71b7f5e)
- [https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling](resources--nfv_service--reference--group-003.md#canonical-6c73f4f34d68de0d1881aa020acde9578eb9b758f2aab0fcd65359f5b1e4e154)
- [https_management.advertise_on_slo_sli.tls_certificates.private_key](resources--nfv_service--reference--group-003.md#canonical-a865d8eee24076eea9f3c8bbdd41aa47e9b2e649df0d523a130e4ab5ec4452c0)
- [https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults](resources--nfv_service--reference--group-003.md#canonical-4e1d7a69bc87776186b92e588924fa1a0b0152671d5c4d976d5210f07f1a7a7e)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-30022b387cd576f0c782dfcba45fea9f6211f83d534f947e7512e992c2608a46)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-b3e56f03abd4532fbee6535661d6dac7ea85f318ecb2c3be6fb3c524e71b7f5e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9afd0f261b20dbe26b1f74535f069ad990a5e9a48067c4a1438a6da56562220d"></a>

## https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms — https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms / 2a97475f2869 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-30022b387cd576f0c782dfcba45fea9f6211f83d534f947e7512e992c2608a46)
- [https_management.advertise_on_slo_sli.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-9558627558b2ed3fb58a624e041ddbf1bb7ad9ccbcf3f1faaa0c30e60b58d257)
- https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms

<a id="canonical-0b1d23a02f771aa560bb99087839d7c89ff931a460a208261afcc09b8c836504"></a>

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

<a id="canonical-bf6b4648dbdc4d45b7f230916b7b576248c006bc5e42b9e237470f63614a8d1e"></a>

## Direct properties — https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms / 2a97475f2869 / 3

<a id="canonical-4338113a94f076a297daa4c079a1ff5cc7d409bd486ada61170262f7ddc5ee6f"></a>

<a id="canonical-97a7b46fb8ca16b358e650d91fc71414f8c0d66a0eef03dfdabba5df03b78203"></a>

## hash_algorithms property — https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms / 2a97475f2869 / 4

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

<a id="canonical-4302b4606bb4092b58f38752cb839becaa637bf4d6e620ace553aceb5a102671"></a>

## Next pages — https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms / 2a97475f2869 / 5

- [https_management.advertise_on_slo_sli.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-9558627558b2ed3fb58a624e041ddbf1bb7ad9ccbcf3f1faaa0c30e60b58d257)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-6c73f4f34d68de0d1881aa020acde9578eb9b758f2aab0fcd65359f5b1e4e154"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-106f89bd1892069980483aa0413a5af36930cbf2fea46ff599fa5c2c2058ad55"></a>

## https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling — https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling / dac11f288220 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-30022b387cd576f0c782dfcba45fea9f6211f83d534f947e7512e992c2608a46)
- [https_management.advertise_on_slo_sli.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-9558627558b2ed3fb58a624e041ddbf1bb7ad9ccbcf3f1faaa0c30e60b58d257)
- https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling

<a id="canonical-c5bd0739c7a83bad2c4198fa544a5a780fea74a3bf904fe246b67829d94d97c3"></a>

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

<a id="canonical-e4f9e8bc82cd371ba91d2dc3e07970cadb997ee1d3a94945d21bcf34c13b533d"></a>

## Direct properties — https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling / dac11f288220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ad46cdc0f73875773ce443fc37ea3575224b8efd4ab323e724cc4eab79890d21"></a>

## Next pages — https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling / dac11f288220 / 4

- [https_management.advertise_on_slo_sli.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-9558627558b2ed3fb58a624e041ddbf1bb7ad9ccbcf3f1faaa0c30e60b58d257)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-a865d8eee24076eea9f3c8bbdd41aa47e9b2e649df0d523a130e4ab5ec4452c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c794a4888e6c1552b61dbd2e9c40ce44ec3f107c76c8676e653059a51057375"></a>

## https_management.advertise_on_slo_sli.tls_certificates.private_key — https_management.advertise_on_slo_sli.tls_certificates.private_key / 5b94477b0f5c / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-30022b387cd576f0c782dfcba45fea9f6211f83d534f947e7512e992c2608a46)
- [https_management.advertise_on_slo_sli.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-9558627558b2ed3fb58a624e041ddbf1bb7ad9ccbcf3f1faaa0c30e60b58d257)
- https_management.advertise_on_slo_sli.tls_certificates.private_key

<a id="canonical-25a6ee25874ce5fbd310c488a848abcee050c025d7d55b5c639f0ec462dbf5db"></a>

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

<a id="canonical-adacd7ed09a1b76c44566212a5452eb41bd09bd2a019ed7d029cbc81bf63f0bb"></a>

## Direct properties — https_management.advertise_on_slo_sli.tls_certificates.private_key / 5b94477b0f5c / 3

- [blindfold_secret_info](resources--nfv_service--reference--group-003.md#canonical-905718933c536138fcee371e46957a969d8e5359122dab992494449a947d78ec): complete subsection reference.

- [clear_secret_info](resources--nfv_service--reference--group-003.md#canonical-2d1017f8ec3ba34779f93ee718630b6f6942847d85884f026862bcf61ec1e52f): complete subsection reference.

<a id="canonical-6aa7477913979d06b3bee859ea39dbea1eb5936b4b394ad49f98040a0a4a446c"></a>

## Next pages — https_management.advertise_on_slo_sli.tls_certificates.private_key / 5b94477b0f5c / 4

- [https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info](resources--nfv_service--reference--group-003.md#canonical-905718933c536138fcee371e46957a969d8e5359122dab992494449a947d78ec)
- [https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info](resources--nfv_service--reference--group-003.md#canonical-2d1017f8ec3ba34779f93ee718630b6f6942847d85884f026862bcf61ec1e52f)
- [https_management.advertise_on_slo_sli.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-9558627558b2ed3fb58a624e041ddbf1bb7ad9ccbcf3f1faaa0c30e60b58d257)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-905718933c536138fcee371e46957a969d8e5359122dab992494449a947d78ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f769fa8e162d6e9ac7198419b9afe79b289d4d111f34dea912e5e2f5c07087d"></a>

## https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info — https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_sec / 001747e1f77a / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-30022b387cd576f0c782dfcba45fea9f6211f83d534f947e7512e992c2608a46)
- [https_management.advertise_on_slo_sli.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-9558627558b2ed3fb58a624e041ddbf1bb7ad9ccbcf3f1faaa0c30e60b58d257)
- [https_management.advertise_on_slo_sli.tls_certificates.private_key](resources--nfv_service--reference--group-003.md#canonical-a865d8eee24076eea9f3c8bbdd41aa47e9b2e649df0d523a130e4ab5ec4452c0)
- https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-1ac7c1c3eee7fa5873d1dc3b2f5cec5f07584c91ea5060753f209915e80fc6b0"></a>

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

<a id="canonical-b993792c02d985f7d2dea2ddbd10075fadb2657733263b3bcbf6a750d27d15e2"></a>

## Direct properties — https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_sec / 001747e1f77a / 3

<a id="canonical-ea0d6f004df46f3ada61d2873be8e63a52aa0b510f6dce8450fbaeb162838088"></a>

<a id="canonical-2c2387aa20fae4ba284873d19f5139e67f32997a6699d6d6cdfdf5cf098760b8"></a>

## decryption_provider property — https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_sec / 001747e1f77a / 4

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

<a id="canonical-aeda5bfdd05af566154c66dafe1ac0428804e892046b808c41c6b49b7abe5de4"></a>

<a id="canonical-8604eb0124a0364846b431c6304b8d151a25dc9c16e8b3dc289a184ac27ff887"></a>

## location property — https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_sec / 001747e1f77a / 5

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

<a id="canonical-4b413584033989401741aaf86a65e5e013a6ebca14e90952172807728b9845b6"></a>

<a id="canonical-0ade52842303363a8cad65dc263bd7466add668d8ddfc0606b95128c51e019f5"></a>

## store_provider property — https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_sec / 001747e1f77a / 6

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

<a id="canonical-78594d46feafb9b1bdbd1012c97948d83c18afc48aac1c07f0cdd45ae91b469e"></a>

## Next pages — https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_sec / 001747e1f77a / 7

- [https_management.advertise_on_slo_sli.tls_certificates.private_key](resources--nfv_service--reference--group-003.md#canonical-a865d8eee24076eea9f3c8bbdd41aa47e9b2e649df0d523a130e4ab5ec4452c0)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-2d1017f8ec3ba34779f93ee718630b6f6942847d85884f026862bcf61ec1e52f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-021815432ebcbcf48ff5d9985e29cf887ba3b47a0b09b91f93b3fbc5a6a4a5f7"></a>

## https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info — https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_ / e67d4f6b260e / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-30022b387cd576f0c782dfcba45fea9f6211f83d534f947e7512e992c2608a46)
- [https_management.advertise_on_slo_sli.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-9558627558b2ed3fb58a624e041ddbf1bb7ad9ccbcf3f1faaa0c30e60b58d257)
- [https_management.advertise_on_slo_sli.tls_certificates.private_key](resources--nfv_service--reference--group-003.md#canonical-a865d8eee24076eea9f3c8bbdd41aa47e9b2e649df0d523a130e4ab5ec4452c0)
- https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info

<a id="canonical-2b51d377d9bebe3874becc73a6c54a079885afcfbaee4288a8422cab5a81cb70"></a>

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

<a id="canonical-908aae84af02d7348f1bc74d347d5dc37c26e2e5ad4e167a4775dd29b5022a76"></a>

## Direct properties — https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_ / e67d4f6b260e / 3

<a id="canonical-dea13368c828078ca6f26613795c72fb7e0b869021cb085474731eed1b4703d9"></a>

<a id="canonical-f04aa5b0dc237ca2dcfedb1f46736e00347608f65b534578951610bd5ef39434"></a>

## provider_ref property — https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_ / e67d4f6b260e / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-262934a37915e127e1960f37ebcd004d7307acee46cf3088f9e6235d955c23ff"></a>

<a id="canonical-f0559a2522122da4db459064f7581d3f5f425ae8f7aec0158990f4e2543fc525"></a>

## url property — https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_ / e67d4f6b260e / 5

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

<a id="canonical-c3454c415165686a52bfd35ab417d109ac4f45e0ff8a42a63156857c50fc7f2e"></a>

## Next pages — https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_ / e67d4f6b260e / 6

- [https_management.advertise_on_slo_sli.tls_certificates.private_key](resources--nfv_service--reference--group-003.md#canonical-a865d8eee24076eea9f3c8bbdd41aa47e9b2e649df0d523a130e4ab5ec4452c0)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-4e1d7a69bc87776186b92e588924fa1a0b0152671d5c4d976d5210f07f1a7a7e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c05ab6e9c9d00c5a402caadcc554b143e2adc5307b2dbc0f3bb64bf4b6dccb0"></a>

## https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults — https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults / dec7a0c2f272 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-30022b387cd576f0c782dfcba45fea9f6211f83d534f947e7512e992c2608a46)
- [https_management.advertise_on_slo_sli.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-9558627558b2ed3fb58a624e041ddbf1bb7ad9ccbcf3f1faaa0c30e60b58d257)
- https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults

<a id="canonical-a0c6155fb728a1d2d75080009d39664baaae49f275922e875cebe3810fcf9adc"></a>

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

<a id="canonical-4a069689c2355ba4fe6b9b3eeab210fbf8725586e251b783592f6116f1ef1bf4"></a>

## Direct properties — https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults / dec7a0c2f272 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7e4ff7cbbe5e2b89c9214a97a45b287147f8947e59813c3940035d0ae9799a62"></a>

## Next pages — https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults / dec7a0c2f272 / 4

- [https_management.advertise_on_slo_sli.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-9558627558b2ed3fb58a624e041ddbf1bb7ad9ccbcf3f1faaa0c30e60b58d257)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-e8f3a8275dfae5058f691263dfcf487950fd2ec4a4bcbf49a9f81cfe3d94e84d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c0d1027898bc33fa10d909b8ae02bc173e8af547c8b001d95e8b3a046f60b309"></a>

## https_management.advertise_on_slo_sli.tls_config — https_management.advertise_on_slo_sli.tls_config / 13b065303c23 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-30022b387cd576f0c782dfcba45fea9f6211f83d534f947e7512e992c2608a46)
- https_management.advertise_on_slo_sli.tls_config

<a id="canonical-cd8544e8be3dd22fd07c05f02839352fffff955e215f9b0b2e31b77dd7fd852a"></a>

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

<a id="canonical-f6542af3fc0a0f1e62285c6a4581e0e44168a84dc9786a40595fdd24d1b21c13"></a>

## Direct properties — https_management.advertise_on_slo_sli.tls_config / 13b065303c23 / 3

- [custom_security](resources--nfv_service--reference--group-003.md#canonical-cef235df309c68b883b8809023f169eaa176cd6749c7e14ae257cbfa4335c3b2): complete subsection reference.

- [default_security](resources--nfv_service--reference--group-003.md#canonical-42269be2dc2741a874503b11c3ef89faf452aebcdc65961061dc2f9ddd8196c2): complete subsection reference.

- [low_security](resources--nfv_service--reference--group-003.md#canonical-7733f25f473cb26b0f732df8a6462d8f3512ea03171af029b23881b38fc9de72): complete subsection reference.

- [medium_security](resources--nfv_service--reference--group-003.md#canonical-96c9211b4fc95e8b5906278c399712677226310c95f25f9ea44a43d045ced32f): complete subsection reference.

<a id="canonical-c25af30b58c898ead28f84ad37b9e0d526ab6b546a966f6a3ff6e6e7e7e011e7"></a>

## Next pages — https_management.advertise_on_slo_sli.tls_config / 13b065303c23 / 4

- [https_management.advertise_on_slo_sli.tls_config.custom_security](resources--nfv_service--reference--group-003.md#canonical-cef235df309c68b883b8809023f169eaa176cd6749c7e14ae257cbfa4335c3b2)
- [https_management.advertise_on_slo_sli.tls_config.default_security](resources--nfv_service--reference--group-003.md#canonical-42269be2dc2741a874503b11c3ef89faf452aebcdc65961061dc2f9ddd8196c2)
- [https_management.advertise_on_slo_sli.tls_config.low_security](resources--nfv_service--reference--group-003.md#canonical-7733f25f473cb26b0f732df8a6462d8f3512ea03171af029b23881b38fc9de72)
- [https_management.advertise_on_slo_sli.tls_config.medium_security](resources--nfv_service--reference--group-003.md#canonical-96c9211b4fc95e8b5906278c399712677226310c95f25f9ea44a43d045ced32f)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-30022b387cd576f0c782dfcba45fea9f6211f83d534f947e7512e992c2608a46)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-cef235df309c68b883b8809023f169eaa176cd6749c7e14ae257cbfa4335c3b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-94af186de8524de6241702e71081760b893f5af81c01cc33a4cb48becb2f5a2c"></a>

## https_management.advertise_on_slo_sli.tls_config.custom_security — https_management.advertise_on_slo_sli.tls_config.custom_security / 9e70b4195c7c / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-30022b387cd576f0c782dfcba45fea9f6211f83d534f947e7512e992c2608a46)
- [https_management.advertise_on_slo_sli.tls_config](resources--nfv_service--reference--group-003.md#canonical-e8f3a8275dfae5058f691263dfcf487950fd2ec4a4bcbf49a9f81cfe3d94e84d)
- https_management.advertise_on_slo_sli.tls_config.custom_security

<a id="canonical-c062269bc28137a8e983da4de769bad521b674e5ab7346de82924c81b03e7d77"></a>

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

<a id="canonical-53680f61bf43bb00b63e865368bda75b768b189ea8f2f0c790c5258030fb5f46"></a>

## Direct properties — https_management.advertise_on_slo_sli.tls_config.custom_security / 9e70b4195c7c / 3

<a id="canonical-7224521419785ef923ecf7ff120cac61ad2695464c924c14d99dd1f438fa8eab"></a>

<a id="canonical-b771e3db5635548d6b9a4d529c05def1dd95bc2abaf9d65674010085913e60ad"></a>

## cipher_suites property — https_management.advertise_on_slo_sli.tls_config.custom_security / 9e70b4195c7c / 4

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

<a id="canonical-8b0230d32c7b0f75c11d6c26d435e31eab6d1c4de9a2bc24641f3d5ddd67d4de"></a>

<a id="canonical-c728ff5b59ab28940ca2df24b72e1ae2019e4bc81ef732cf8992d43b23eb9f67"></a>

## max_version property — https_management.advertise_on_slo_sli.tls_config.custom_security / 9e70b4195c7c / 5

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

<a id="canonical-bd701b2d019b75aa32054e906064391f777c65ef33dc47eadaa77560c0e0b779"></a>

<a id="canonical-6b4c9f7a7184076c017515e6c75a1cb3456ff8eb69ef59f47065740065d4267c"></a>

## min_version property — https_management.advertise_on_slo_sli.tls_config.custom_security / 9e70b4195c7c / 6

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

<a id="canonical-3b07c2e924a93b93c659a37a830915225d8f9eb07fec90f6daa4364ec447707a"></a>

## Next pages — https_management.advertise_on_slo_sli.tls_config.custom_security / 9e70b4195c7c / 7

- [https_management.advertise_on_slo_sli.tls_config](resources--nfv_service--reference--group-003.md#canonical-e8f3a8275dfae5058f691263dfcf487950fd2ec4a4bcbf49a9f81cfe3d94e84d)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-42269be2dc2741a874503b11c3ef89faf452aebcdc65961061dc2f9ddd8196c2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c6c470117aed7d3ea1cbfd61d42a1b9a863354eca85a7bf4dd22f3efc3411783"></a>

## https_management.advertise_on_slo_sli.tls_config.default_security — https_management.advertise_on_slo_sli.tls_config.default_security / 51134bd27efc / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-30022b387cd576f0c782dfcba45fea9f6211f83d534f947e7512e992c2608a46)
- [https_management.advertise_on_slo_sli.tls_config](resources--nfv_service--reference--group-003.md#canonical-e8f3a8275dfae5058f691263dfcf487950fd2ec4a4bcbf49a9f81cfe3d94e84d)
- https_management.advertise_on_slo_sli.tls_config.default_security

<a id="canonical-a5555bf93d1dee72f894c0c86b288023506100181b5a3a943e94f51c93a1b339"></a>

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

<a id="canonical-418bb084bb35dbf6af3f0a138500d1e91bf2b97f1cbf4f2f909265631a8e86d1"></a>

## Direct properties — https_management.advertise_on_slo_sli.tls_config.default_security / 51134bd27efc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a22adacf40aa9d0aa9f2659a31baad7be8e7f83f1962e12ff269a1bcb7d0e17b"></a>

## Next pages — https_management.advertise_on_slo_sli.tls_config.default_security / 51134bd27efc / 4

- [https_management.advertise_on_slo_sli.tls_config](resources--nfv_service--reference--group-003.md#canonical-e8f3a8275dfae5058f691263dfcf487950fd2ec4a4bcbf49a9f81cfe3d94e84d)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-7733f25f473cb26b0f732df8a6462d8f3512ea03171af029b23881b38fc9de72"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d521df5b832f5a20b45f12cf35185930972cc9b2ef72dc69d7e2c28a1bc7469"></a>

## https_management.advertise_on_slo_sli.tls_config.low_security — https_management.advertise_on_slo_sli.tls_config.low_security / 285bd5ab810c / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-30022b387cd576f0c782dfcba45fea9f6211f83d534f947e7512e992c2608a46)
- [https_management.advertise_on_slo_sli.tls_config](resources--nfv_service--reference--group-003.md#canonical-e8f3a8275dfae5058f691263dfcf487950fd2ec4a4bcbf49a9f81cfe3d94e84d)
- https_management.advertise_on_slo_sli.tls_config.low_security

<a id="canonical-c5d3ac53c0c074c5d1a683d3eb5a22cbfa5d77d025d96a321d9147264fd46f8e"></a>

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

<a id="canonical-f0199536093daf3788344dcbf24aa640665b2167d8d642d6b3675c064d6a2613"></a>

## Direct properties — https_management.advertise_on_slo_sli.tls_config.low_security / 285bd5ab810c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-77737f2a377908a3a6250e5c61d786305374a7374be511a3fdecc674181d9874"></a>

## Next pages — https_management.advertise_on_slo_sli.tls_config.low_security / 285bd5ab810c / 4

- [https_management.advertise_on_slo_sli.tls_config](resources--nfv_service--reference--group-003.md#canonical-e8f3a8275dfae5058f691263dfcf487950fd2ec4a4bcbf49a9f81cfe3d94e84d)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-96c9211b4fc95e8b5906278c399712677226310c95f25f9ea44a43d045ced32f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b3a354a19504ec999329098020dee621dc8fcf726c442dcca6cb091c38c16a07"></a>

## https_management.advertise_on_slo_sli.tls_config.medium_security — https_management.advertise_on_slo_sli.tls_config.medium_security / d537e71768e3 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-30022b387cd576f0c782dfcba45fea9f6211f83d534f947e7512e992c2608a46)
- [https_management.advertise_on_slo_sli.tls_config](resources--nfv_service--reference--group-003.md#canonical-e8f3a8275dfae5058f691263dfcf487950fd2ec4a4bcbf49a9f81cfe3d94e84d)
- https_management.advertise_on_slo_sli.tls_config.medium_security

<a id="canonical-8a436f12937504d5a232265eb31f1211754528c69a17456e4799d523fe17ec8b"></a>

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

<a id="canonical-1e0c787d89c51538db2e6d38048c6b58d8df836838e7b927c1dba605d355ae4e"></a>

## Direct properties — https_management.advertise_on_slo_sli.tls_config.medium_security / d537e71768e3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2621090767f606b5088d2d2861aefd7e84a6ffe09f2b897511de7fcff56800e9"></a>

## Next pages — https_management.advertise_on_slo_sli.tls_config.medium_security / d537e71768e3 / 4

- [https_management.advertise_on_slo_sli.tls_config](resources--nfv_service--reference--group-003.md#canonical-e8f3a8275dfae5058f691263dfcf487950fd2ec4a4bcbf49a9f81cfe3d94e84d)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-1257604a48a6882e78dae67503e07dec0e7ef0ab20383466d80273b5c231bff4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7245a5780210441584872181737ecef29d73379d49d718db5edc4bf51d8656f"></a>

## https_management.advertise_on_slo_sli.use_mtls — https_management.advertise_on_slo_sli.use_mtls / 0a1abe430194 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-30022b387cd576f0c782dfcba45fea9f6211f83d534f947e7512e992c2608a46)
- https_management.advertise_on_slo_sli.use_mtls

<a id="canonical-1108693ddc7d3e5e1a96b6c2dd0639d21bb95beb0ebc978f51f39bb6f9d63673"></a>

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

<a id="canonical-3a5bb23a5a8007808af00d3e5590c8dae1c347ef325a1a43c17c8da5efe6fe60"></a>

## Direct properties — https_management.advertise_on_slo_sli.use_mtls / 0a1abe430194 / 3

<a id="canonical-62902d3d547a034789e39abe70c6e6c00a28ab2635384d840371c6a3edac218e"></a>

<a id="canonical-5ebcedaba6e7b089a667e32acbb0b6de8af61ee8cad8fba88eba788641092358"></a>

## client_certificate_optional property — https_management.advertise_on_slo_sli.use_mtls / 0a1abe430194 / 4

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

- [crl](resources--nfv_service--reference--group-003.md#canonical-9aec8eb7d6a7a7bbec9f4f829ab6b6d067edb8829a90bfcce47b57c344c2df13): complete subsection reference.

- [no_crl](resources--nfv_service--reference--group-003.md#canonical-632c0003b2c6dd138b42564521cc83404a2e1ff9684ad4a0dda82a2f5feb04b1): complete subsection reference.

- [trusted_ca](resources--nfv_service--reference--group-003.md#canonical-30292c765d9f1d7f7716039c4fa5879f7d05f12dab086e54a069ad6b5dbfafe3): complete subsection reference.

<a id="canonical-30496d380617931f9371a4a5428be8b199e7cd9ffb9cc78729836698c4a504c7"></a>

<a id="canonical-44cf1fd8442e080be95bb3a16976b0dd526c2f1586e7237868dedf4d917db354"></a>

## trusted_ca_url property — https_management.advertise_on_slo_sli.use_mtls / 0a1abe430194 / 5

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

- [xfcc_disabled](resources--nfv_service--reference--group-003.md#canonical-f3e64b90cfeb44595d31d50e83e7076cf8ab455b7818a60526deb4c7dc6cf4a4): complete subsection reference.

- [xfcc_options](resources--nfv_service--reference--group-003.md#canonical-a264812f1489795ebec79ed809256af1e48a4b8def08e60f2a8a495f44af8fe1): complete subsection reference.

<a id="canonical-47d7ebdb322d57c14a94adf1cc63e5a18a1dc7594b3f0d9705b5fca8d2bda8ed"></a>

## Next pages — https_management.advertise_on_slo_sli.use_mtls / 0a1abe430194 / 6

- [https_management.advertise_on_slo_sli.use_mtls.crl](resources--nfv_service--reference--group-003.md#canonical-9aec8eb7d6a7a7bbec9f4f829ab6b6d067edb8829a90bfcce47b57c344c2df13)
- [https_management.advertise_on_slo_sli.use_mtls.no_crl](resources--nfv_service--reference--group-003.md#canonical-632c0003b2c6dd138b42564521cc83404a2e1ff9684ad4a0dda82a2f5feb04b1)
- [https_management.advertise_on_slo_sli.use_mtls.trusted_ca](resources--nfv_service--reference--group-003.md#canonical-30292c765d9f1d7f7716039c4fa5879f7d05f12dab086e54a069ad6b5dbfafe3)
- [https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled](resources--nfv_service--reference--group-003.md#canonical-f3e64b90cfeb44595d31d50e83e7076cf8ab455b7818a60526deb4c7dc6cf4a4)
- [https_management.advertise_on_slo_sli.use_mtls.xfcc_options](resources--nfv_service--reference--group-003.md#canonical-a264812f1489795ebec79ed809256af1e48a4b8def08e60f2a8a495f44af8fe1)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-30022b387cd576f0c782dfcba45fea9f6211f83d534f947e7512e992c2608a46)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-9aec8eb7d6a7a7bbec9f4f829ab6b6d067edb8829a90bfcce47b57c344c2df13"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-540727b4065178265c66293810a8592c0d5575673257f1832638022ae904a089"></a>

## https_management.advertise_on_slo_sli.use_mtls.crl — https_management.advertise_on_slo_sli.use_mtls.crl / 2928394cac42 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-30022b387cd576f0c782dfcba45fea9f6211f83d534f947e7512e992c2608a46)
- [https_management.advertise_on_slo_sli.use_mtls](resources--nfv_service--reference--group-003.md#canonical-1257604a48a6882e78dae67503e07dec0e7ef0ab20383466d80273b5c231bff4)
- https_management.advertise_on_slo_sli.use_mtls.crl

<a id="canonical-92af364fc3ef5bcb74c130f4b506fb1e3fb7bb9552566a445c3c6dfaaa0b52ad"></a>

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

<a id="canonical-b04ce7376227fd1698f3c56dc0b06fd12ec566f0f983906dd7f738da5fde384f"></a>

## Direct properties — https_management.advertise_on_slo_sli.use_mtls.crl / 2928394cac42 / 3

<a id="canonical-c52a3a9e5e55e3c55c2a1cebbddb28f0b72f5b6fb81ab8d87d8dc548b0772643"></a>

<a id="canonical-ca2848747e57ff2f0b5f6d4051316d79815d232fdb9d362c8d09d28e883d0d45"></a>

## name property — https_management.advertise_on_slo_sli.use_mtls.crl / 2928394cac42 / 4

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

<a id="canonical-ac3c1b8f983edaedfe8489ecfed42314c2d223b1fe43ee221fcaae23fc37ac53"></a>

<a id="canonical-7e696561b2a5ee5d137948f59de93b325a50327a7b1d0b48868a2ddb9ec10437"></a>

## namespace property — https_management.advertise_on_slo_sli.use_mtls.crl / 2928394cac42 / 5

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

<a id="canonical-3ecf2a033ea96842b475f129b24f2cabb1af01c3ef1da084a9d2639b93cb0335"></a>

<a id="canonical-3867fbc71d429e1e7de211c8d949c1dfe64f343a8fb5382edffdb086515bb484"></a>

## tenant property — https_management.advertise_on_slo_sli.use_mtls.crl / 2928394cac42 / 6

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

<a id="canonical-0ed4dc5aaadcda9153c55830602a8550190679b42291f1f2521f1b2e4634eefc"></a>

## Next pages — https_management.advertise_on_slo_sli.use_mtls.crl / 2928394cac42 / 7

- [https_management.advertise_on_slo_sli.use_mtls](resources--nfv_service--reference--group-003.md#canonical-1257604a48a6882e78dae67503e07dec0e7ef0ab20383466d80273b5c231bff4)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-632c0003b2c6dd138b42564521cc83404a2e1ff9684ad4a0dda82a2f5feb04b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-680bc0cf976afed3f20071fce278c2b17a04fab712d2074c8281c565ded1ab79"></a>

## https_management.advertise_on_slo_sli.use_mtls.no_crl — https_management.advertise_on_slo_sli.use_mtls.no_crl / b00f09b8a4fd / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-30022b387cd576f0c782dfcba45fea9f6211f83d534f947e7512e992c2608a46)
- [https_management.advertise_on_slo_sli.use_mtls](resources--nfv_service--reference--group-003.md#canonical-1257604a48a6882e78dae67503e07dec0e7ef0ab20383466d80273b5c231bff4)
- https_management.advertise_on_slo_sli.use_mtls.no_crl

<a id="canonical-0c1acabd82d5c3aa051a870f4cefd1682b017aeeadcdcea4e19d1ff6a879c3c5"></a>

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

<a id="canonical-8b2957e1c2ae4fa15dbbf7e503c6ccb949fc15d56ddd41468a9400f4ec1faa03"></a>

## Direct properties — https_management.advertise_on_slo_sli.use_mtls.no_crl / b00f09b8a4fd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7b17a5b7772863a1a2c1ab1178a56d76e2e7e2f5a7d32de9bfd11cf708163ec2"></a>

## Next pages — https_management.advertise_on_slo_sli.use_mtls.no_crl / b00f09b8a4fd / 4

- [https_management.advertise_on_slo_sli.use_mtls](resources--nfv_service--reference--group-003.md#canonical-1257604a48a6882e78dae67503e07dec0e7ef0ab20383466d80273b5c231bff4)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-30292c765d9f1d7f7716039c4fa5879f7d05f12dab086e54a069ad6b5dbfafe3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-08b569dc203ec116fd9c61277934ab0b1882c15bd9a825e0d39202015eea61e7"></a>

## https_management.advertise_on_slo_sli.use_mtls.trusted_ca — https_management.advertise_on_slo_sli.use_mtls.trusted_ca / 21ab3f71c44d / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-30022b387cd576f0c782dfcba45fea9f6211f83d534f947e7512e992c2608a46)
- [https_management.advertise_on_slo_sli.use_mtls](resources--nfv_service--reference--group-003.md#canonical-1257604a48a6882e78dae67503e07dec0e7ef0ab20383466d80273b5c231bff4)
- https_management.advertise_on_slo_sli.use_mtls.trusted_ca

<a id="canonical-dd2c4b022d0a3eb15e1b5ff3741362e638758a2ca99b0510c5b7b51a76acf994"></a>

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-9e5bf32172c499f7a9e6099c102416cd4a49cf907fdce1ef81aa0ec0f4abc056"></a>

## Direct properties — https_management.advertise_on_slo_sli.use_mtls.trusted_ca / 21ab3f71c44d / 3

<a id="canonical-71c4c07308ee22f677fd73ed94ed1274ca406e617a87355f6d306d5d1bc5532f"></a>

<a id="canonical-c9ec22bfbee357e0ecd5933a068e262534e66a80d0c748c3dd6a80255f28d765"></a>

## name property — https_management.advertise_on_slo_sli.use_mtls.trusted_ca / 21ab3f71c44d / 4

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

<a id="canonical-72738dff86dce06560a799359237d4f927cd41c763d1458de861ae3ef68177ae"></a>

<a id="canonical-3d10435c1b55447c5e5da6ea7a2932db49a933b59e34fcb0464b74dc9984c2c0"></a>

## namespace property — https_management.advertise_on_slo_sli.use_mtls.trusted_ca / 21ab3f71c44d / 5

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

<a id="canonical-299269b493fa7854ce84bab7cc613e889493085b076e3d0203ae3e8c1544bf42"></a>

<a id="canonical-b82393baf00394d21806ad4a6ed9b4d37bcb2d37c4b5967fadb457592644af1f"></a>

## tenant property — https_management.advertise_on_slo_sli.use_mtls.trusted_ca / 21ab3f71c44d / 6

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

<a id="canonical-452b25e23fe643df2ac2100b8aafc1d2ea1e1c4a2d761af14f4829c1cb075282"></a>

## Next pages — https_management.advertise_on_slo_sli.use_mtls.trusted_ca / 21ab3f71c44d / 7

- [https_management.advertise_on_slo_sli.use_mtls](resources--nfv_service--reference--group-003.md#canonical-1257604a48a6882e78dae67503e07dec0e7ef0ab20383466d80273b5c231bff4)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-f3e64b90cfeb44595d31d50e83e7076cf8ab455b7818a60526deb4c7dc6cf4a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5b33794270f1657a07ff12715bfdc69211487f3bb7dbf0a6568ee68fa53581c"></a>

## https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled — https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled / 6a28bdcf6c6f / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-30022b387cd576f0c782dfcba45fea9f6211f83d534f947e7512e992c2608a46)
- [https_management.advertise_on_slo_sli.use_mtls](resources--nfv_service--reference--group-003.md#canonical-1257604a48a6882e78dae67503e07dec0e7ef0ab20383466d80273b5c231bff4)
- https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled

<a id="canonical-2c24d043f536442190167d3a8eb47c46cfcf797f17515db04cdb1fd15e88896a"></a>

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
xfcc_disabled = {}
```

<a id="canonical-b0408e2610b6d9a2dfd1563e5913b6349ac681f192f5738ee57b79979591f4b5"></a>

## Direct properties — https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled / 6a28bdcf6c6f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b44fbf398958491d5a774629e92271def62d32cba66d2dccc3ede8137729ed4e"></a>

## Next pages — https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled / 6a28bdcf6c6f / 4

- [https_management.advertise_on_slo_sli.use_mtls](resources--nfv_service--reference--group-003.md#canonical-1257604a48a6882e78dae67503e07dec0e7ef0ab20383466d80273b5c231bff4)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-a264812f1489795ebec79ed809256af1e48a4b8def08e60f2a8a495f44af8fe1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b762eba7109d257923e315f484db73ee6807bdc56466ae4dfcbaa9b816bfe48"></a>

## https_management.advertise_on_slo_sli.use_mtls.xfcc_options — https_management.advertise_on_slo_sli.use_mtls.xfcc_options / 4c8531502ee0 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-30022b387cd576f0c782dfcba45fea9f6211f83d534f947e7512e992c2608a46)
- [https_management.advertise_on_slo_sli.use_mtls](resources--nfv_service--reference--group-003.md#canonical-1257604a48a6882e78dae67503e07dec0e7ef0ab20383466d80273b5c231bff4)
- https_management.advertise_on_slo_sli.use_mtls.xfcc_options

<a id="canonical-38ae173a7133fe2a98ddcc16969193103920277175584d5cfbf1a949959668ac"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("xfcc_header_elements")}
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
xfcc_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-ecda6f270b125ae27eafa9663475f30cdbd44008a3190fd0abc16c339850fa35"></a>

## Direct properties — https_management.advertise_on_slo_sli.use_mtls.xfcc_options / 4c8531502ee0 / 3

<a id="canonical-6e0b30a1e85e32ebda9d9f4dc9e36c83d976b1666ad29d8953809177613458c9"></a>

<a id="canonical-86dd7e249b79d83495a7666447724c11f39408b275805488235da26e0b844f17"></a>

## xfcc_header_elements property — https_management.advertise_on_slo_sli.use_mtls.xfcc_options / 4c8531502ee0 / 4

Type: `["list", "string"]`. Optional.

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

<a id="canonical-f099700ad26328d8fcebb12a708f870b4ac189dafede16d4bcfeb0568433dece"></a>

## Next pages — https_management.advertise_on_slo_sli.use_mtls.xfcc_options / 4c8531502ee0 / 5

- [https_management.advertise_on_slo_sli.use_mtls](resources--nfv_service--reference--group-003.md#canonical-1257604a48a6882e78dae67503e07dec0e7ef0ab20383466d80273b5c231bff4)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-dbbbdded43c0c30607bcb8e329159cdf82634f24371abc298c763310ecccf580"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6af5e3539db8c2b66ddc82b7932f4c21b81634722a38aa8229d3f0c8d5b30329"></a>

## https_management.advertise_on_slo_vip — https_management.advertise_on_slo_vip / 8a44a5fa317b / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- https_management.advertise_on_slo_vip

<a id="canonical-3531b5c834ecb32b9e38f4efc093de765d809a806e2a740169f7efaaf764206e"></a>

Type: `"object"`. single nested block, Optional.

Inline TLS Parameters. Inline TLS parameters.

Upstream description:

Inline TLS parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("tls_certificates"),
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
advertise_on_slo_vip {
  # Configure direct properties listed below.
}
```

<a id="canonical-68aa3e7dd7011b2dbdfdf90f543e19ae9ca9c1bdf20faf17f0730e2543b27bd2"></a>

## Direct properties — https_management.advertise_on_slo_vip / 8a44a5fa317b / 3

- [no_mtls](resources--nfv_service--reference--group-003.md#canonical-25111c824ffe1d9b33e695522f9bc2d4693f13b7dc2d0cff781bd2edc4e58531): complete subsection reference.

- [tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2140062389231f0e346eecef12f9cd5e6f803e6c56bb7c71e4339f11117bffcf): complete subsection reference.

- [tls_config](resources--nfv_service--reference--group-003.md#canonical-68031dcce46dc3f3000cda912caa0242a6f089c8e9e3b34d63b91ca9ec885fb7): complete subsection reference.

- [use_mtls](resources--nfv_service--reference--group-003.md#canonical-07045da532f13d2ebc1cc4c495b995595e8a0a2ff3d3a88294887a10abcbbfab): complete subsection reference.

<a id="canonical-6d7a5d0dbcc9ef92459cb9e7a79c2ab19fca34973047be2740fef75abd91929c"></a>

## Next pages — https_management.advertise_on_slo_vip / 8a44a5fa317b / 4

- [https_management.advertise_on_slo_vip.no_mtls](resources--nfv_service--reference--group-003.md#canonical-25111c824ffe1d9b33e695522f9bc2d4693f13b7dc2d0cff781bd2edc4e58531)
- [https_management.advertise_on_slo_vip.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2140062389231f0e346eecef12f9cd5e6f803e6c56bb7c71e4339f11117bffcf)
- [https_management.advertise_on_slo_vip.tls_config](resources--nfv_service--reference--group-003.md#canonical-68031dcce46dc3f3000cda912caa0242a6f089c8e9e3b34d63b91ca9ec885fb7)
- [https_management.advertise_on_slo_vip.use_mtls](resources--nfv_service--reference--group-003.md#canonical-07045da532f13d2ebc1cc4c495b995595e8a0a2ff3d3a88294887a10abcbbfab)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-25111c824ffe1d9b33e695522f9bc2d4693f13b7dc2d0cff781bd2edc4e58531"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c5c1a4bb901aaa8fc57d68456e7856e67e81e841ff79500efb58aa729d304b5"></a>

## https_management.advertise_on_slo_vip.no_mtls — https_management.advertise_on_slo_vip.no_mtls / cf6abea1182a / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-dbbbdded43c0c30607bcb8e329159cdf82634f24371abc298c763310ecccf580)
- https_management.advertise_on_slo_vip.no_mtls

<a id="canonical-2b4929fbd87561bd855b8b75491587792bf121657378171098209ef10f613d3c"></a>

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

<a id="canonical-fee337fc88c18dbd06ed2a54fca7db293a939dcd9d0b2f16556af2c9fb18e4c9"></a>

## Direct properties — https_management.advertise_on_slo_vip.no_mtls / cf6abea1182a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3b51ae201c312de1f49bf67e674c6fa663e4ef60b799c2f73fa0686b999dd6b5"></a>

## Next pages — https_management.advertise_on_slo_vip.no_mtls / cf6abea1182a / 4

- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-dbbbdded43c0c30607bcb8e329159cdf82634f24371abc298c763310ecccf580)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-2140062389231f0e346eecef12f9cd5e6f803e6c56bb7c71e4339f11117bffcf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3fb66db1e172ab894e049560083234d7b2998c35b2385122598f7f364fff4312"></a>

## https_management.advertise_on_slo_vip.tls_certificates — https_management.advertise_on_slo_vip.tls_certificates / 354bc601e4a1 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-dbbbdded43c0c30607bcb8e329159cdf82634f24371abc298c763310ecccf580)
- https_management.advertise_on_slo_vip.tls_certificates

<a id="canonical-8fd4c68ed449a0ec050a39af4e95903a9a75ffec20de722562f4a410f12e8bc0"></a>

Type: `"object"`. list nested block, Optional.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Upstream description:

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("certificate_url"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "disable_ocsp_stapling"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "use_system_defaults"),
  validators.ConflictingListObjectAttributes("disable_ocsp_stapling",
    "use_system_defaults")}
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

Terraform syntax:

```terraform
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-3bd33da0d39d8f91ecbc7f1ef5232a9fa458e3a5e7d43458129466ce9e2137f5"></a>

## Direct properties — https_management.advertise_on_slo_vip.tls_certificates / 354bc601e4a1 / 3

<a id="canonical-5838ffcce3448515b0655a8a279edfdfeaf68f30b3583c397425d876606e56ab"></a>

<a id="canonical-277241b805119c4237da2aaeec2784138ad415583b049c56931c4d8b876f9c91"></a>

## certificate_url property — https_management.advertise_on_slo_vip.tls_certificates / 354bc601e4a1 / 4

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

- [custom_hash_algorithms](resources--nfv_service--reference--group-003.md#canonical-fe2495f7343837bc2af646fffa0ccf482445c21779e4688551444fcf4e164985): complete subsection reference.

<a id="canonical-7152c6b4c0d6f37ed954d080c4b01c51518dde59275068ddbbef5a39cbeab228"></a>

<a id="canonical-5358d82782368d7a594f15ad4f6965204cfe0228451810bb7e773d2eeb44cd1f"></a>

## description_spec property — https_management.advertise_on_slo_vip.tls_certificates / 354bc601e4a1 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--nfv_service--reference--group-003.md#canonical-6bdae75cfb48dbeed3b2cbb70be0fe222402faa5c6124924a101272ba8b80081): complete subsection reference.

- [private_key](resources--nfv_service--reference--group-003.md#canonical-0a19bd40668ece487cb4993ee122f7c2bb43322e368f2bdbc0c36894ab14f787): complete subsection reference.

- [use_system_defaults](resources--nfv_service--reference--group-003.md#canonical-4546511ef6e8edbdd7871a4c1ad929702901912dd2fd40acb6ee706b550aef69): complete subsection reference.

<a id="canonical-43a494cfcbd9a6ac1bcdd0c492f2d90864bf922c8be5455382a27d623bbf2eb5"></a>

## Next pages — https_management.advertise_on_slo_vip.tls_certificates / 354bc601e4a1 / 6

- [https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms](resources--nfv_service--reference--group-003.md#canonical-fe2495f7343837bc2af646fffa0ccf482445c21779e4688551444fcf4e164985)
- [https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling](resources--nfv_service--reference--group-003.md#canonical-6bdae75cfb48dbeed3b2cbb70be0fe222402faa5c6124924a101272ba8b80081)
- [https_management.advertise_on_slo_vip.tls_certificates.private_key](resources--nfv_service--reference--group-003.md#canonical-0a19bd40668ece487cb4993ee122f7c2bb43322e368f2bdbc0c36894ab14f787)
- [https_management.advertise_on_slo_vip.tls_certificates.use_system_defaults](resources--nfv_service--reference--group-003.md#canonical-4546511ef6e8edbdd7871a4c1ad929702901912dd2fd40acb6ee706b550aef69)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-dbbbdded43c0c30607bcb8e329159cdf82634f24371abc298c763310ecccf580)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-fe2495f7343837bc2af646fffa0ccf482445c21779e4688551444fcf4e164985"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8212fd44f9bd8689bec6d13891be598fb11bc3dfc22a620b276882e336189fe4"></a>

## https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms — https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms / b221f86e50af / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-dbbbdded43c0c30607bcb8e329159cdf82634f24371abc298c763310ecccf580)
- [https_management.advertise_on_slo_vip.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2140062389231f0e346eecef12f9cd5e6f803e6c56bb7c71e4339f11117bffcf)
- https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms

<a id="canonical-5b5b17b671aab13747e56d3d9c7c1dbebea803b371b44c74d88e5f85ada90e74"></a>

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

<a id="canonical-a9cf074e1fac14f13fd066ad00f2bfe3b940adbd3b4f1041c0eeda3c7f74189f"></a>

## Direct properties — https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms / b221f86e50af / 3

<a id="canonical-e7b1aae1306ce55c4af9dab21175bd7d72764fe78df43d9e6812a9324e508ebf"></a>

<a id="canonical-1de8276636300406d7557544642207af2ea4ea77223f578d8b78eb20827ffacc"></a>

## hash_algorithms property — https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms / b221f86e50af / 4

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

<a id="canonical-b452058f6e6d2d2669adee41a069e28807ad9291b72ccb4cc91e6b379b48c613"></a>

## Next pages — https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms / b221f86e50af / 5

- [https_management.advertise_on_slo_vip.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2140062389231f0e346eecef12f9cd5e6f803e6c56bb7c71e4339f11117bffcf)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-6bdae75cfb48dbeed3b2cbb70be0fe222402faa5c6124924a101272ba8b80081"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-395248d1fa19820d4b00e921759e70da7155875d54ef6d157bf76301ae6cb897"></a>

## https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling — https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling / ea79e6657454 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-dbbbdded43c0c30607bcb8e329159cdf82634f24371abc298c763310ecccf580)
- [https_management.advertise_on_slo_vip.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2140062389231f0e346eecef12f9cd5e6f803e6c56bb7c71e4339f11117bffcf)
- https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling

<a id="canonical-2965c9c329c90b457ef2e38700e6f277a64d37e4479a5fe77c58656f3bc3cad6"></a>

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

<a id="canonical-318688ed371b05a83c0e8778ef727c233aa775717b64693fa95c1fd5639a6c73"></a>

## Direct properties — https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling / ea79e6657454 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b5c6711cb2c6a21cdb5247ba3dcf3279581dee77464bfd32c1af8fe976af5d49"></a>

## Next pages — https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling / ea79e6657454 / 4

- [https_management.advertise_on_slo_vip.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2140062389231f0e346eecef12f9cd5e6f803e6c56bb7c71e4339f11117bffcf)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-0a19bd40668ece487cb4993ee122f7c2bb43322e368f2bdbc0c36894ab14f787"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c46faa19003aa115a76acb7e5d3944e66de05932c81642e53bde0e1b0ffb12c2"></a>

## https_management.advertise_on_slo_vip.tls_certificates.private_key — https_management.advertise_on_slo_vip.tls_certificates.private_key / 7c0e754a23ab / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-dbbbdded43c0c30607bcb8e329159cdf82634f24371abc298c763310ecccf580)
- [https_management.advertise_on_slo_vip.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2140062389231f0e346eecef12f9cd5e6f803e6c56bb7c71e4339f11117bffcf)
- https_management.advertise_on_slo_vip.tls_certificates.private_key

<a id="canonical-9b1197291d728060c51a9aa4530085b51530df2b3e86277c75f71ba661a55f50"></a>

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

<a id="canonical-cb295e56bfda1ab9d2f0f03a8faab8f7cddd36e6d429e08d830e4c0b89c94112"></a>

## Direct properties — https_management.advertise_on_slo_vip.tls_certificates.private_key / 7c0e754a23ab / 3

- [blindfold_secret_info](resources--nfv_service--reference--group-003.md#canonical-33bf85a0e83dce20eb3ae9809e3a76fe46928220d296ac1726323611fd4a42a4): complete subsection reference.

- [clear_secret_info](resources--nfv_service--reference--group-003.md#canonical-9636412455598c3e048c3a2e4ff820a475a0d57417484097001304b0412fce13): complete subsection reference.

<a id="canonical-c6f443324cb6aa5b5311ef111dc1cdcfb8b889fa659abe25dd8ba836dc61ecf9"></a>

## Next pages — https_management.advertise_on_slo_vip.tls_certificates.private_key / 7c0e754a23ab / 4

- [https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info](resources--nfv_service--reference--group-003.md#canonical-33bf85a0e83dce20eb3ae9809e3a76fe46928220d296ac1726323611fd4a42a4)
- [https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info](resources--nfv_service--reference--group-003.md#canonical-9636412455598c3e048c3a2e4ff820a475a0d57417484097001304b0412fce13)
- [https_management.advertise_on_slo_vip.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2140062389231f0e346eecef12f9cd5e6f803e6c56bb7c71e4339f11117bffcf)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-33bf85a0e83dce20eb3ae9809e3a76fe46928220d296ac1726323611fd4a42a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f0ec72821ec001635f7a7badaa20429ed57acbc46f758bfdab7d5f5ea88a3c0"></a>

## https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info — https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_sec / 1a874ede9788 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-dbbbdded43c0c30607bcb8e329159cdf82634f24371abc298c763310ecccf580)
- [https_management.advertise_on_slo_vip.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2140062389231f0e346eecef12f9cd5e6f803e6c56bb7c71e4339f11117bffcf)
- [https_management.advertise_on_slo_vip.tls_certificates.private_key](resources--nfv_service--reference--group-003.md#canonical-0a19bd40668ece487cb4993ee122f7c2bb43322e368f2bdbc0c36894ab14f787)
- https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-7d579e8fe6766374ad2c9fbaff0d18887c7b273958e438c6ac2d5567e33d52bb"></a>

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

<a id="canonical-3866902250bacb57f3668300a7af94ed753b7762024c530b5757e74d32609caa"></a>

## Direct properties — https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_sec / 1a874ede9788 / 3

<a id="canonical-e53e392fc9413e65179f3be8832a2a651d6bdabd01a7387e6526f1dd3d4e7c1e"></a>

<a id="canonical-d631dbd2e9a0b82a8d8aeea9798a19af2b89d816d9b07baf0dbda099913d6aee"></a>

## decryption_provider property — https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_sec / 1a874ede9788 / 4

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

<a id="canonical-5e9fd1503101eca30bc77d8867c1386c3ce45c851e09a472b19ad7341b0624b3"></a>

<a id="canonical-8fa3070aaf2cd2ee2bea453abd1c26d85d8792f0ac758346a6d3ce12f9d994cb"></a>

## location property — https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_sec / 1a874ede9788 / 5

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

<a id="canonical-3ca515216ff9f3895f458e62a44a3a3c9b7a841b67f9e882870abf48ee409c4c"></a>

<a id="canonical-4a431bcb401364f7b676f2a2e2f1810781a912c62864cc93c8b833d6208eec11"></a>

## store_provider property — https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_sec / 1a874ede9788 / 6

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

<a id="canonical-ffcf6a8d76a7d7c1a87e9ee91be6f8f88b21bd4bbac0a05f813413cba59916b9"></a>

## Next pages — https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_sec / 1a874ede9788 / 7

- [https_management.advertise_on_slo_vip.tls_certificates.private_key](resources--nfv_service--reference--group-003.md#canonical-0a19bd40668ece487cb4993ee122f7c2bb43322e368f2bdbc0c36894ab14f787)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-9636412455598c3e048c3a2e4ff820a475a0d57417484097001304b0412fce13"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-58e8980bd1365f6ad5bdad4fd5aeb01627daa9d2bde9f99a8196a5944b56d8e5"></a>

## https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info — https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_ / c342c9f8b788 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-dbbbdded43c0c30607bcb8e329159cdf82634f24371abc298c763310ecccf580)
- [https_management.advertise_on_slo_vip.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2140062389231f0e346eecef12f9cd5e6f803e6c56bb7c71e4339f11117bffcf)
- [https_management.advertise_on_slo_vip.tls_certificates.private_key](resources--nfv_service--reference--group-003.md#canonical-0a19bd40668ece487cb4993ee122f7c2bb43322e368f2bdbc0c36894ab14f787)
- https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info

<a id="canonical-d68bf7388433afa5dcbe11ba69c8e461a25edf6e51577cfcf990b419f8b5102e"></a>

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

<a id="canonical-50ab832875bf0dea5fee891fdfa656c16375fba8c04ccee210497de2eb626a25"></a>

## Direct properties — https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_ / c342c9f8b788 / 3

<a id="canonical-bcd29b0c8f0c3c6de8b75b3169d72cfec4768ab9dbe72f258be28876fc185c31"></a>

<a id="canonical-ef16e40d56490e154c05a35db0f566c977d6860674cd8ae3f6e8c0d4fdf74d18"></a>

## provider_ref property — https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_ / c342c9f8b788 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-7c0170f46996e2a0fededd199faa5ceb68da829761b208ef8ac4b3ec75399999"></a>

<a id="canonical-92ae1366bbc465b88da01406c72378cf81a3c13e4899c51a2747b086ed7c6bb1"></a>

## url property — https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_ / c342c9f8b788 / 5

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

<a id="canonical-e27754b918ff9c7e0420ec82f332fe6695daf43682c7899ce2a126af3b705e33"></a>

## Next pages — https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_ / c342c9f8b788 / 6

- [https_management.advertise_on_slo_vip.tls_certificates.private_key](resources--nfv_service--reference--group-003.md#canonical-0a19bd40668ece487cb4993ee122f7c2bb43322e368f2bdbc0c36894ab14f787)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-4546511ef6e8edbdd7871a4c1ad929702901912dd2fd40acb6ee706b550aef69"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf53dbf82db6c48a9601e02e85cab4eaa6b71dc399bb64da0eeb772c25b3b3cb"></a>

## https_management.advertise_on_slo_vip.tls_certificates.use_system_defaults — https_management.advertise_on_slo_vip.tls_certificates.use_system_defaults / 8991f31ea02b / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-dbbbdded43c0c30607bcb8e329159cdf82634f24371abc298c763310ecccf580)
- [https_management.advertise_on_slo_vip.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2140062389231f0e346eecef12f9cd5e6f803e6c56bb7c71e4339f11117bffcf)
- https_management.advertise_on_slo_vip.tls_certificates.use_system_defaults

<a id="canonical-8d1ba086845ea53aa1a2b0229366bd7a04924f882490d0a549b0415309613852"></a>

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

<a id="canonical-c9f7ffe90af815b68f6ce59dda6ecfe65947781eece490f015f5c144fa482448"></a>

## Direct properties — https_management.advertise_on_slo_vip.tls_certificates.use_system_defaults / 8991f31ea02b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d062c4c2c1f0aca4f7e7aee7fcd41afa746c5f77784af2be757e05ec1b1e176b"></a>

## Next pages — https_management.advertise_on_slo_vip.tls_certificates.use_system_defaults / 8991f31ea02b / 4

- [https_management.advertise_on_slo_vip.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2140062389231f0e346eecef12f9cd5e6f803e6c56bb7c71e4339f11117bffcf)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-68031dcce46dc3f3000cda912caa0242a6f089c8e9e3b34d63b91ca9ec885fb7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3fc262c65d757a9ebbb25e6feac6485f3eecdd09e854b4b30e7ff80e0884acd0"></a>

## https_management.advertise_on_slo_vip.tls_config — https_management.advertise_on_slo_vip.tls_config / d5b2b4ab7cdc / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-dbbbdded43c0c30607bcb8e329159cdf82634f24371abc298c763310ecccf580)
- https_management.advertise_on_slo_vip.tls_config

<a id="canonical-778600701aca3cc945ab71d9ea9912e0a2d10ae8b8fa5e944cf3ebd800a0c00d"></a>

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

<a id="canonical-9b471961e68af0c9809aeced4e396265fb472704d2599913fa86274379f87ebe"></a>

## Direct properties — https_management.advertise_on_slo_vip.tls_config / d5b2b4ab7cdc / 3

- [custom_security](resources--nfv_service--reference--group-003.md#canonical-e137e0cbcfec5a6b12488027704cf4a5943b8b46736edc123e5d17dc64b83abf): complete subsection reference.

- [default_security](resources--nfv_service--reference--group-003.md#canonical-7c7459ee36563dda8c40bfb595247650a2f0525cff31f5fa55ba87e1198d7353): complete subsection reference.

- [low_security](resources--nfv_service--reference--group-003.md#canonical-29583a34ead7ff30f3af4c684e287840930e09a19ffff3bd2ede689f096eb0da): complete subsection reference.

- [medium_security](resources--nfv_service--reference--group-003.md#canonical-dda7264d310e199719084d71aff403d1beebf89c6bf484e5af3f7f0366b351a3): complete subsection reference.

<a id="canonical-3b8d3b15e03136d02c2acdfa3a16bbb63c425357e9d6ca6cacff9b1594e4c3b4"></a>

## Next pages — https_management.advertise_on_slo_vip.tls_config / d5b2b4ab7cdc / 4

- [https_management.advertise_on_slo_vip.tls_config.custom_security](resources--nfv_service--reference--group-003.md#canonical-e137e0cbcfec5a6b12488027704cf4a5943b8b46736edc123e5d17dc64b83abf)
- [https_management.advertise_on_slo_vip.tls_config.default_security](resources--nfv_service--reference--group-003.md#canonical-7c7459ee36563dda8c40bfb595247650a2f0525cff31f5fa55ba87e1198d7353)
- [https_management.advertise_on_slo_vip.tls_config.low_security](resources--nfv_service--reference--group-003.md#canonical-29583a34ead7ff30f3af4c684e287840930e09a19ffff3bd2ede689f096eb0da)
- [https_management.advertise_on_slo_vip.tls_config.medium_security](resources--nfv_service--reference--group-003.md#canonical-dda7264d310e199719084d71aff403d1beebf89c6bf484e5af3f7f0366b351a3)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-dbbbdded43c0c30607bcb8e329159cdf82634f24371abc298c763310ecccf580)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-e137e0cbcfec5a6b12488027704cf4a5943b8b46736edc123e5d17dc64b83abf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ba3190eda1db8c861fa489e2ad1e1feec216c56e6587d9eacbffbe83bd223520"></a>

## https_management.advertise_on_slo_vip.tls_config.custom_security — https_management.advertise_on_slo_vip.tls_config.custom_security / be69f1efaa32 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-dbbbdded43c0c30607bcb8e329159cdf82634f24371abc298c763310ecccf580)
- [https_management.advertise_on_slo_vip.tls_config](resources--nfv_service--reference--group-003.md#canonical-68031dcce46dc3f3000cda912caa0242a6f089c8e9e3b34d63b91ca9ec885fb7)
- https_management.advertise_on_slo_vip.tls_config.custom_security

<a id="canonical-51ee633e44aeec128ffbcf858f572750ef1b41cacdd8a94e09d703fb664f3dac"></a>

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

<a id="canonical-3e6ca6224088a80c2f854e18fa91362773fdfff29df0d753a6302af128258e00"></a>

## Direct properties — https_management.advertise_on_slo_vip.tls_config.custom_security / be69f1efaa32 / 3

<a id="canonical-3db08b8a600f81980f3415d86226cd6476c87b3bb8325af83653352c0428f99d"></a>

<a id="canonical-da5f92cbc0a34cbc89c52897a3ecce2deeea36441d1018a9108dfb2e57dbd46d"></a>

## cipher_suites property — https_management.advertise_on_slo_vip.tls_config.custom_security / be69f1efaa32 / 4

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

<a id="canonical-ea8191063cb2062f6c810021dc07f73cf89bc4c63c3e8f2f97517bbdb6406c4c"></a>

<a id="canonical-524639dfe2ba410e48272e62f795d88846a7276f212453223c42310a823c1ff4"></a>

## max_version property — https_management.advertise_on_slo_vip.tls_config.custom_security / be69f1efaa32 / 5

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

<a id="canonical-b907185baf9e64d4aa84029b626c54bcbc8439422e33b95a9a99f5357c48bf72"></a>

<a id="canonical-3e90108e8a2b448b0efa1182be83b633b26acb02a18edc7948e56f6d973e3739"></a>

## min_version property — https_management.advertise_on_slo_vip.tls_config.custom_security / be69f1efaa32 / 6

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

<a id="canonical-2e85e87e8551532fa813f1e131b17afa686153925d1983be64d3fca626637ead"></a>

## Next pages — https_management.advertise_on_slo_vip.tls_config.custom_security / be69f1efaa32 / 7

- [https_management.advertise_on_slo_vip.tls_config](resources--nfv_service--reference--group-003.md#canonical-68031dcce46dc3f3000cda912caa0242a6f089c8e9e3b34d63b91ca9ec885fb7)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-7c7459ee36563dda8c40bfb595247650a2f0525cff31f5fa55ba87e1198d7353"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b039dda9ead9e64c1d2c6cd242115f7258b32edc098a0a1d232503b91eea9422"></a>

## https_management.advertise_on_slo_vip.tls_config.default_security — https_management.advertise_on_slo_vip.tls_config.default_security / d764d71e9dbd / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-dbbbdded43c0c30607bcb8e329159cdf82634f24371abc298c763310ecccf580)
- [https_management.advertise_on_slo_vip.tls_config](resources--nfv_service--reference--group-003.md#canonical-68031dcce46dc3f3000cda912caa0242a6f089c8e9e3b34d63b91ca9ec885fb7)
- https_management.advertise_on_slo_vip.tls_config.default_security

<a id="canonical-fe30954c57479a8f21b23b89a7270dc5a565934dab0f096c46c2697c9af27b5d"></a>

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

<a id="canonical-0a459032ee58a5aa7ff401610be19e3120e2042c9bfad3e47aefc2d5541dfcf3"></a>

## Direct properties — https_management.advertise_on_slo_vip.tls_config.default_security / d764d71e9dbd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bb1af071252efee9a6ec1f9ee0fb515a333477fc8ef149c9bce8afdb66dc7456"></a>

## Next pages — https_management.advertise_on_slo_vip.tls_config.default_security / d764d71e9dbd / 4

- [https_management.advertise_on_slo_vip.tls_config](resources--nfv_service--reference--group-003.md#canonical-68031dcce46dc3f3000cda912caa0242a6f089c8e9e3b34d63b91ca9ec885fb7)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-29583a34ead7ff30f3af4c684e287840930e09a19ffff3bd2ede689f096eb0da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b00a1c039382943d5e1d9621c36b9cc964427c0daa7a544c9460806d01262891"></a>

## https_management.advertise_on_slo_vip.tls_config.low_security — https_management.advertise_on_slo_vip.tls_config.low_security / 73f3641390c4 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-dbbbdded43c0c30607bcb8e329159cdf82634f24371abc298c763310ecccf580)
- [https_management.advertise_on_slo_vip.tls_config](resources--nfv_service--reference--group-003.md#canonical-68031dcce46dc3f3000cda912caa0242a6f089c8e9e3b34d63b91ca9ec885fb7)
- https_management.advertise_on_slo_vip.tls_config.low_security

<a id="canonical-40bfab942dcc19c6c3923ede13181a5d4be2159816d37e770f806ded9e3768e5"></a>

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

<a id="canonical-31f93a8f30d97103bd994bfbba3546ec7301c7211f9b6b6890e60de1b74d9fbc"></a>

## Direct properties — https_management.advertise_on_slo_vip.tls_config.low_security / 73f3641390c4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-69d556bedd441d7d83619a29939716b5a1552f427e221f76867f1c3e3633f73c"></a>

## Next pages — https_management.advertise_on_slo_vip.tls_config.low_security / 73f3641390c4 / 4

- [https_management.advertise_on_slo_vip.tls_config](resources--nfv_service--reference--group-003.md#canonical-68031dcce46dc3f3000cda912caa0242a6f089c8e9e3b34d63b91ca9ec885fb7)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-dda7264d310e199719084d71aff403d1beebf89c6bf484e5af3f7f0366b351a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca0b6f0faf5c7ef376d7c7bc66b53367fb4b8208844fa0d44e0dc5f79a27e4d7"></a>

## https_management.advertise_on_slo_vip.tls_config.medium_security — https_management.advertise_on_slo_vip.tls_config.medium_security / 257a3556800e / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-dbbbdded43c0c30607bcb8e329159cdf82634f24371abc298c763310ecccf580)
- [https_management.advertise_on_slo_vip.tls_config](resources--nfv_service--reference--group-003.md#canonical-68031dcce46dc3f3000cda912caa0242a6f089c8e9e3b34d63b91ca9ec885fb7)
- https_management.advertise_on_slo_vip.tls_config.medium_security

<a id="canonical-0bcc7fe65b417deca064a8b3b8cfee50900ff66aed671df81ccd9805f05723fa"></a>

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

<a id="canonical-9afe2f4f41e2982d73325d49eb9c9add52491079ec9ecf81d018851d767d3c46"></a>

## Direct properties — https_management.advertise_on_slo_vip.tls_config.medium_security / 257a3556800e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d861a284a98f1651c9ace46854852d9a7150521629c2484ae315b0dd1cce951c"></a>

## Next pages — https_management.advertise_on_slo_vip.tls_config.medium_security / 257a3556800e / 4

- [https_management.advertise_on_slo_vip.tls_config](resources--nfv_service--reference--group-003.md#canonical-68031dcce46dc3f3000cda912caa0242a6f089c8e9e3b34d63b91ca9ec885fb7)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-07045da532f13d2ebc1cc4c495b995595e8a0a2ff3d3a88294887a10abcbbfab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da6cecb21d505e9fdedfb8aff9e7d1ba00ba010786482a6c2c7390e87dbb8572"></a>

## https_management.advertise_on_slo_vip.use_mtls — https_management.advertise_on_slo_vip.use_mtls / 8a6d41599aae / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-dbbbdded43c0c30607bcb8e329159cdf82634f24371abc298c763310ecccf580)
- https_management.advertise_on_slo_vip.use_mtls

<a id="canonical-a058339644255cc34b12f45e16622453c8a7610525973fdd0527448879bf2904"></a>

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

<a id="canonical-f6b4aa200c9d7969a4bd0ce7515fdd175390ce679ffccd9267ee0620d0f63393"></a>

## Direct properties — https_management.advertise_on_slo_vip.use_mtls / 8a6d41599aae / 3

<a id="canonical-b7f391c2e21c2f2aef5c7a518ff4230eec01e809a02dded5a938c872abda2022"></a>

<a id="canonical-160389257d0da14f338cf3e29a4cf5568d89c8aede99c2f01671c95eb4dd79ee"></a>

## client_certificate_optional property — https_management.advertise_on_slo_vip.use_mtls / 8a6d41599aae / 4

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

- [crl](resources--nfv_service--reference--group-003.md#canonical-659df38b7df4e922aeaadb166b52955a0d7ff8dae566e0d2351de05e8c70efd4): complete subsection reference.

- [no_crl](resources--nfv_service--reference--group-003.md#canonical-b5fbfb09fa721b29f61d05b1ad9a6859649ae57f6a4a08c3467e173bd414c4a5): complete subsection reference.

- [trusted_ca](resources--nfv_service--reference--group-003.md#canonical-9b104d1e59db3217518f940f2345abc07d42cd9e3ae531e8b4ef5085eab06dc0): complete subsection reference.

<a id="canonical-1f1843db987714a4d0d4520377eaa28a9a67413556759cd420e32f3563aeb5fa"></a>

<a id="canonical-3e63263f20fd272b6041ecb1f8c633c37d193a6cd15ac4e18330733704efe428"></a>

## trusted_ca_url property — https_management.advertise_on_slo_vip.use_mtls / 8a6d41599aae / 5

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

- [xfcc_disabled](resources--nfv_service--reference--group-003.md#canonical-496fb9e69442ca3186f6512fdfbc597be6303b7e39b204fccc156e3e2a49115d): complete subsection reference.

- [xfcc_options](resources--nfv_service--reference--group-003.md#canonical-eb5a41d3755ccf44d9e7e17392ee105de82acde8b6c0aa9b10165d0ead514203): complete subsection reference.

<a id="canonical-6a68c9c579864a364a8446bdb80c02a8167340b87cc7650839485e7e859eb1c1"></a>

## Next pages — https_management.advertise_on_slo_vip.use_mtls / 8a6d41599aae / 6

- [https_management.advertise_on_slo_vip.use_mtls.crl](resources--nfv_service--reference--group-003.md#canonical-659df38b7df4e922aeaadb166b52955a0d7ff8dae566e0d2351de05e8c70efd4)
- [https_management.advertise_on_slo_vip.use_mtls.no_crl](resources--nfv_service--reference--group-003.md#canonical-b5fbfb09fa721b29f61d05b1ad9a6859649ae57f6a4a08c3467e173bd414c4a5)
- [https_management.advertise_on_slo_vip.use_mtls.trusted_ca](resources--nfv_service--reference--group-003.md#canonical-9b104d1e59db3217518f940f2345abc07d42cd9e3ae531e8b4ef5085eab06dc0)
- [https_management.advertise_on_slo_vip.use_mtls.xfcc_disabled](resources--nfv_service--reference--group-003.md#canonical-496fb9e69442ca3186f6512fdfbc597be6303b7e39b204fccc156e3e2a49115d)
- [https_management.advertise_on_slo_vip.use_mtls.xfcc_options](resources--nfv_service--reference--group-003.md#canonical-eb5a41d3755ccf44d9e7e17392ee105de82acde8b6c0aa9b10165d0ead514203)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-dbbbdded43c0c30607bcb8e329159cdf82634f24371abc298c763310ecccf580)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-659df38b7df4e922aeaadb166b52955a0d7ff8dae566e0d2351de05e8c70efd4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d8ce98a8e5da819b4c41f1ac22279fd5d6f19c7ccf83a3f7f4e36f691e336140"></a>

## https_management.advertise_on_slo_vip.use_mtls.crl — https_management.advertise_on_slo_vip.use_mtls.crl / 95dd7712e5df / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-dbbbdded43c0c30607bcb8e329159cdf82634f24371abc298c763310ecccf580)
- [https_management.advertise_on_slo_vip.use_mtls](resources--nfv_service--reference--group-003.md#canonical-07045da532f13d2ebc1cc4c495b995595e8a0a2ff3d3a88294887a10abcbbfab)
- https_management.advertise_on_slo_vip.use_mtls.crl

<a id="canonical-9c0adfe477a9f53262cf89cd58aa413f97d717ca2b525cbe811b0fb44fb1f562"></a>

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

<a id="canonical-bdc4623f320ebd19c5fbb9d77e3fc7e590b26c4119801cace5b4f078d139e60f"></a>

## Direct properties — https_management.advertise_on_slo_vip.use_mtls.crl / 95dd7712e5df / 3

<a id="canonical-02e7bede44c7332db85c270748c682dfb0e20ef653b8a6bc585e9f46246317ac"></a>

<a id="canonical-742833b0c102b30df75eaa83ae2e4c271b832cab4dbb3842118debe35e84cfb3"></a>

## name property — https_management.advertise_on_slo_vip.use_mtls.crl / 95dd7712e5df / 4

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

<a id="canonical-f69a894ee16ca97c8fce7eb6c6045cafc39cb4f9cedf682bb2c5e3bc9ae64b71"></a>

<a id="canonical-07f577225aed3ec798cdd0164861f1734af5344c618b92f29a629a5f9847c7ba"></a>

## namespace property — https_management.advertise_on_slo_vip.use_mtls.crl / 95dd7712e5df / 5

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

<a id="canonical-d382ba1156d991af2656bb16a1ba0f48e5d6f3a310cc4804f7df83ab35cb8083"></a>

<a id="canonical-ff76f81d1915b37dbee23610683b67e1daeea2d75b4f253764cf09ea58250125"></a>

## tenant property — https_management.advertise_on_slo_vip.use_mtls.crl / 95dd7712e5df / 6

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

<a id="canonical-1f224768e62a3c41dac3a49f2c5c98078c9142471bd048289debf9b30d960a53"></a>

## Next pages — https_management.advertise_on_slo_vip.use_mtls.crl / 95dd7712e5df / 7

- [https_management.advertise_on_slo_vip.use_mtls](resources--nfv_service--reference--group-003.md#canonical-07045da532f13d2ebc1cc4c495b995595e8a0a2ff3d3a88294887a10abcbbfab)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-b5fbfb09fa721b29f61d05b1ad9a6859649ae57f6a4a08c3467e173bd414c4a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df84b96923328f1cc19d4b305520318603e6b54cf6b49e13ed9c70deb07d2222"></a>

## https_management.advertise_on_slo_vip.use_mtls.no_crl — https_management.advertise_on_slo_vip.use_mtls.no_crl / 2788850f60e1 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-dbbbdded43c0c30607bcb8e329159cdf82634f24371abc298c763310ecccf580)
- [https_management.advertise_on_slo_vip.use_mtls](resources--nfv_service--reference--group-003.md#canonical-07045da532f13d2ebc1cc4c495b995595e8a0a2ff3d3a88294887a10abcbbfab)
- https_management.advertise_on_slo_vip.use_mtls.no_crl

<a id="canonical-9ef83c13d7142da78d7a009232713dd5dd5548f485a74d9a89ac96462fb243e1"></a>

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

<a id="canonical-e8c38e9fecf06b9d49acaed03203a37956e63cd262e1c8c3e5ad13d073ca039e"></a>

## Direct properties — https_management.advertise_on_slo_vip.use_mtls.no_crl / 2788850f60e1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1744d90ad865025abd2e9d075b9739c9cb95da24280bc0fe4c39294805716029"></a>

## Next pages — https_management.advertise_on_slo_vip.use_mtls.no_crl / 2788850f60e1 / 4

- [https_management.advertise_on_slo_vip.use_mtls](resources--nfv_service--reference--group-003.md#canonical-07045da532f13d2ebc1cc4c495b995595e8a0a2ff3d3a88294887a10abcbbfab)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-9b104d1e59db3217518f940f2345abc07d42cd9e3ae531e8b4ef5085eab06dc0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9508f3b622039e84b479bb341889b5c476e7336842339b07f23be97d35e2441"></a>

## https_management.advertise_on_slo_vip.use_mtls.trusted_ca — https_management.advertise_on_slo_vip.use_mtls.trusted_ca / b69befd6353d / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-dbbbdded43c0c30607bcb8e329159cdf82634f24371abc298c763310ecccf580)
- [https_management.advertise_on_slo_vip.use_mtls](resources--nfv_service--reference--group-003.md#canonical-07045da532f13d2ebc1cc4c495b995595e8a0a2ff3d3a88294887a10abcbbfab)
- https_management.advertise_on_slo_vip.use_mtls.trusted_ca

<a id="canonical-cd3105831bd9cf983b42b83ecc2606973143c29280b1650fa6030a00d39402b8"></a>

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-0408ead178b76a851e0c01907b4d967936a7df1f2e63b5cb396b6ab053ee5de7"></a>

## Direct properties — https_management.advertise_on_slo_vip.use_mtls.trusted_ca / b69befd6353d / 3

<a id="canonical-734c5e7a3b9952bd0af1b942b7ac13b5c5072c9fbdff0d09a9dddd89ccc130bf"></a>

<a id="canonical-65f542a4d9fe441f364b449d643e5de75b96c074a58c1869d6b85cdb7c9cc3cb"></a>

## name property — https_management.advertise_on_slo_vip.use_mtls.trusted_ca / b69befd6353d / 4

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

<a id="canonical-e51c8cecc51ef90fb8cf82f667cede501d1d72709c0be588c5fdc88d43e54182"></a>

<a id="canonical-59f662c4369322c15382deb120a694327a9f4233f8cfc515a0e55efd28a7f686"></a>

## namespace property — https_management.advertise_on_slo_vip.use_mtls.trusted_ca / b69befd6353d / 5

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

<a id="canonical-6ad41ef4192875b507bc061990b2c0483faf16ef29df4f878b1871fc46abd6c9"></a>

<a id="canonical-0edbd0ba7d29cb30396a9d53cf847f6cb64bc97e600156b1bab213583f80bac9"></a>

## tenant property — https_management.advertise_on_slo_vip.use_mtls.trusted_ca / b69befd6353d / 6

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

<a id="canonical-c1d64ecd44db3a3041febbd90dce23305a69dc61a72a8bb5bda4ebb86eb24cc5"></a>

## Next pages — https_management.advertise_on_slo_vip.use_mtls.trusted_ca / b69befd6353d / 7

- [https_management.advertise_on_slo_vip.use_mtls](resources--nfv_service--reference--group-003.md#canonical-07045da532f13d2ebc1cc4c495b995595e8a0a2ff3d3a88294887a10abcbbfab)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-496fb9e69442ca3186f6512fdfbc597be6303b7e39b204fccc156e3e2a49115d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a4941b9b54d442e501d9a846bcc6fd3f7d3f60773c03518bd9d65a8599997ad"></a>

## https_management.advertise_on_slo_vip.use_mtls.xfcc_disabled — https_management.advertise_on_slo_vip.use_mtls.xfcc_disabled / 8b0ee1d1cc94 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-dbbbdded43c0c30607bcb8e329159cdf82634f24371abc298c763310ecccf580)
- [https_management.advertise_on_slo_vip.use_mtls](resources--nfv_service--reference--group-003.md#canonical-07045da532f13d2ebc1cc4c495b995595e8a0a2ff3d3a88294887a10abcbbfab)
- https_management.advertise_on_slo_vip.use_mtls.xfcc_disabled

<a id="canonical-8ee97d0f2be9cb7facd806fa0dd55ff6a6f7c52e04c5c2d2bdd0d38513a74972"></a>

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
xfcc_disabled = {}
```

<a id="canonical-40377fe476855dd7ce683603d5b94944a4f07ec9151955d2fcb1134e7065d01a"></a>

## Direct properties — https_management.advertise_on_slo_vip.use_mtls.xfcc_disabled / 8b0ee1d1cc94 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7db4f762a8c3827e660ce109072774313750efc1541ba39d962dddfe246cd84c"></a>

## Next pages — https_management.advertise_on_slo_vip.use_mtls.xfcc_disabled / 8b0ee1d1cc94 / 4

- [https_management.advertise_on_slo_vip.use_mtls](resources--nfv_service--reference--group-003.md#canonical-07045da532f13d2ebc1cc4c495b995595e8a0a2ff3d3a88294887a10abcbbfab)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-eb5a41d3755ccf44d9e7e17392ee105de82acde8b6c0aa9b10165d0ead514203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6896691d0266ccd2ccae6f6d4c2ac80847482e74f08e1cc8051868e650ec6e3d"></a>

## https_management.advertise_on_slo_vip.use_mtls.xfcc_options — https_management.advertise_on_slo_vip.use_mtls.xfcc_options / f31054b09e48 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-dbbbdded43c0c30607bcb8e329159cdf82634f24371abc298c763310ecccf580)
- [https_management.advertise_on_slo_vip.use_mtls](resources--nfv_service--reference--group-003.md#canonical-07045da532f13d2ebc1cc4c495b995595e8a0a2ff3d3a88294887a10abcbbfab)
- https_management.advertise_on_slo_vip.use_mtls.xfcc_options

<a id="canonical-bfb6defb9f2a4323c3836436f0937744a856bafde7af20b85b009e784453cd73"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("xfcc_header_elements")}
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
xfcc_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-87523492aed0b2b4a7a2ff8f140e05d3bff3b19ce65c60d3e6a23014f908243a"></a>

## Direct properties — https_management.advertise_on_slo_vip.use_mtls.xfcc_options / f31054b09e48 / 3

<a id="canonical-afb7a24ac3224edafdf5940c257bdc4223221fc12111192a545e238880100791"></a>

<a id="canonical-72f6287aa4d9065ba9d1244ff5dc8e59b0add1263ca51451da5ee93b5c2a9d90"></a>

## xfcc_header_elements property — https_management.advertise_on_slo_vip.use_mtls.xfcc_options / f31054b09e48 / 4

Type: `["list", "string"]`. Optional.

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

<a id="canonical-d8e28702817e3c7059a88630d014aa142259768d395f3a3cabd516be0ee7a351"></a>

## Next pages — https_management.advertise_on_slo_vip.use_mtls.xfcc_options / f31054b09e48 / 5

- [https_management.advertise_on_slo_vip.use_mtls](resources--nfv_service--reference--group-003.md#canonical-07045da532f13d2ebc1cc4c495b995595e8a0a2ff3d3a88294887a10abcbbfab)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-4d3e72934bfc108341684f9a988111dde4fcdd5edc9a8f34d8eb22a612cddd3e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6a86b37bf8de15528ac1b8395db5e5d3446b07fa5b6334927f8461bce595c60"></a>

## https_management.default_https_port — https_management.default_https_port / 5fae7971a135 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- https_management.default_https_port

<a id="canonical-3894cf08a900164b908fcb9ff5b3d8413a0e9be08a34feff162eb595cae8e202"></a>

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
default_https_port = {}
```

<a id="canonical-2797a7d7919cd0ba7133deb330179adcf7ea7f35a4ee0262ea77e9fa354754e6"></a>

## Direct properties — https_management.default_https_port / 5fae7971a135 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2ecff9e7a40f066de0011f07b6516173747f4d5df5b790cd9981793e037908fa"></a>

## Next pages — https_management.default_https_port / 5fae7971a135 / 4

- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-683b0e043ad34e16d220e00d2c0ad7d7794cd7a996943065814c81f0d29e03cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5847167dc8bee2c9c653ae46b587151ef6ef373d0ce06afbe9ccc29f6057843a"></a>

## palo_alto_fw_service — palo_alto_fw_service / d254996e1398 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- palo_alto_fw_service

<a id="canonical-409735399e3d2be8e2f34d7145e44073824af1c2148dd24576cf8c92863b1104"></a>

Type: `"object"`. single nested block, Optional.

Palo Alto Networks VM-Series next-generation firewall configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto_setup",
    "ssh_key"),
  validators.ConflictingObjectAttributes("disable_panaroma",
    "panorama_server"),
  validators.ConflictingObjectAttributes("pan_ami_bundle1",
    "pan_ami_bundle2")}
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
  "x-ves-oneof-field-ami_choice": "[\"pan_ami_bundle1\",\"pan_ami_bundle2\"]",
  "x-ves-oneof-field-panaroma_connection": "[\"disable_panaroma\",\"panorama_server\"]",
  "x-ves-oneof-field-setup_options": "[\"auto_setup\",\"ssh_key\"]"
}
```

Terraform syntax:

```terraform
palo_alto_fw_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-2b2528494ff18acb69c7feeacc1d6502b16630ec00f330a7f9b43e4d1a1d4b63"></a>

## Direct properties — palo_alto_fw_service / d254996e1398 / 3

- [auto_setup](resources--nfv_service--reference--group-003.md#canonical-7bf638e8f5f6d4cd41bc1908f771ac84af047472fffcf4f680086848cf0d43db): complete subsection reference.

- [aws_tgw_site](resources--nfv_service--reference--group-004.md#canonical-7154c513294dcde83423cb49e7c2c9583005df7a729e04e7b041b85d0a6fad29): complete subsection reference.

- [disable_panaroma](resources--nfv_service--reference--group-004.md#canonical-5b3a157c00b3631985c671ca0652d9b5bbacb64e999edb8b5e787209de0a1142): complete subsection reference.

<a id="canonical-a76b746ee398a519d7abb7127e51b340798ee98ec4d46b61dc3ae10543c8cadd"></a>

<a id="canonical-d2196f3bd3f760bf7c2691e29429e6f9e0695d7ae0aeebaab5004a3c779a4a73"></a>

## instance_type property — palo_alto_fw_service / d254996e1398 / 4

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("PALO_ALTO_FW_AWS_INSTANCE_TYPE_M4_XLARGE",
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
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_R5_2XLARGE"),
}
```

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

- [pan_ami_bundle1](resources--nfv_service--reference--group-004.md#canonical-cff2ad597e13c8f342ca161809452d25cd1d6d3a68e3cf59e6f5316bfe455e4c): complete subsection reference.

- [pan_ami_bundle2](resources--nfv_service--reference--group-004.md#canonical-65382323b951c01510ff5617d2dcc84b1d074913e191352eb24b3b43e3471040): complete subsection reference.

- [panorama_server](resources--nfv_service--reference--group-004.md#canonical-45ec14f98c4298e49ab80b0eab6e4d912220244b1166f017f7ae9536733a9395): complete subsection reference.

- [service_nodes](resources--nfv_service--reference--group-004.md#canonical-8d807ab1575a9ed710e3454886f81ddf66001c6a42213bf4fcf98c2b5a47020d): complete subsection reference.

<a id="canonical-06e9d71b182ee53300cb5ba69196cf3709d569ef154907f1f86828e51357db69"></a>

<a id="canonical-a97cdfe1cf0d2197b442d74d47228255b202e4a21ce6db5aab49118e161b07ad"></a>

## ssh_key property — palo_alto_fw_service / d254996e1398 / 5

Type: `"string"`. Optional.

Exclusive with \[auto\_setup\] Setup Authorized Public SSH key. User will be able to SSH to the
vmseries nodes using its corresponding SSH private key.

Upstream description:

Exclusive with \[auto\_setup\] Setup Authorized Public SSH key. User will be able to SSH to the
vmseries nodes using its corresponding SSH private key.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 8192),
}
```

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

<a id="canonical-aa82a4b6a78a00bcba62b64625124cfc7d637a63e73a5948c26dc0dd54e22f6e"></a>

<a id="canonical-c639f6fed55d640ec21907ce06b4fb9ccbb33093c30a4a40aa50652f6c25e05e"></a>

## tags property — palo_alto_fw_service / d254996e1398 / 6

Type: `["map", "string"]`. Optional.

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

<a id="canonical-df873e96e04ff6469b73b68430a8994f4a57501f64657cc266b3342cad01c574"></a>

<a id="canonical-6f5cfd0ab21a568a9216341a3effb279e0e1d943cf39636ad17c81565924616c"></a>

## version property — palo_alto_fw_service / d254996e1398 / 7

Type: `"string"`. Optional.

\[Enum: 11.0.0\] PAN VM-Series version. PAN-OS version. The only possible value is \`11.0.0\`.

Upstream description:

PAN-OS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("11.0.0"),
}
```

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

<a id="canonical-1161d604981f192d4bad311696e72e66e707a22033e0b9872445c7d310be1814"></a>

## Next pages — palo_alto_fw_service / d254996e1398 / 8

- [palo_alto_fw_service.auto_setup](resources--nfv_service--reference--group-003.md#canonical-7bf638e8f5f6d4cd41bc1908f771ac84af047472fffcf4f680086848cf0d43db)
- [palo_alto_fw_service.aws_tgw_site](resources--nfv_service--reference--group-004.md#canonical-7154c513294dcde83423cb49e7c2c9583005df7a729e04e7b041b85d0a6fad29)
- [palo_alto_fw_service.disable_panaroma](resources--nfv_service--reference--group-004.md#canonical-5b3a157c00b3631985c671ca0652d9b5bbacb64e999edb8b5e787209de0a1142)
- [palo_alto_fw_service.pan_ami_bundle1](resources--nfv_service--reference--group-004.md#canonical-cff2ad597e13c8f342ca161809452d25cd1d6d3a68e3cf59e6f5316bfe455e4c)
- [palo_alto_fw_service.pan_ami_bundle2](resources--nfv_service--reference--group-004.md#canonical-65382323b951c01510ff5617d2dcc84b1d074913e191352eb24b3b43e3471040)
- [palo_alto_fw_service.panorama_server](resources--nfv_service--reference--group-004.md#canonical-45ec14f98c4298e49ab80b0eab6e4d912220244b1166f017f7ae9536733a9395)
- [palo_alto_fw_service.service_nodes](resources--nfv_service--reference--group-004.md#canonical-8d807ab1575a9ed710e3454886f81ddf66001c6a42213bf4fcf98c2b5a47020d)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-7bf638e8f5f6d4cd41bc1908f771ac84af047472fffcf4f680086848cf0d43db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-14377f115b75f0c12367bd89c1f132e8db6f7631604ddb19a8340e4c0ca96f9e"></a>

## palo_alto_fw_service.auto_setup — palo_alto_fw_service.auto_setup / 927d7025fdee / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-683b0e043ad34e16d220e00d2c0ad7d7794cd7a996943065814c81f0d29e03cf)
- palo_alto_fw_service.auto_setup

<a id="canonical-f4c86c8e1e1632d1724101dc4457dd9dae42e8c25ddaf04bdee5d21e0a7bfc3a"></a>

Type: `"object"`. single nested block, Optional.

For auto-setup, SSH public and pvt keys are needed. Using the given config user, SSH and API access
will be configured.

Upstream description:

For auto-setup, SSH public and pvt keys are needed. Using the given config user, SSH and API access
will be configured.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("admin_username")}
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
  "x-ves-oneof-field-ssh_keys_choice": "[\"manual_ssh_keys\"]"
}
```

Terraform syntax:

```terraform
auto_setup {
  # Configure direct properties listed below.
}
```

<a id="canonical-08f799e3f1f5685517e03bd86e5f914a3e4a60775a03ca329059a22629c581d0"></a>

## Direct properties — palo_alto_fw_service.auto_setup / 927d7025fdee / 3

- [admin_password](resources--nfv_service--reference--group-003.md#canonical-891e267f3e7873cc1cb3dc120f9b9997c1ff664c338ef0d59b130a8a1d61e525): complete subsection reference.

<a id="canonical-804d3e51d1e12b4a3465ee2b89eec0ae71e8d900c7432ac28ed363cc842c2bb0"></a>

<a id="canonical-87c60956500ced780b5a2634be5e96e70939c2d113937d244363419f158990a9"></a>

## admin_username property — palo_alto_fw_service.auto_setup / 927d7025fdee / 4

Type: `"string"`. Optional.

Firewall Admin Username. Firewall Admin Username.

Upstream description:

Firewall Admin Username.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [manual_ssh_keys](resources--nfv_service--reference--group-003.md#canonical-0fe0c13a624ad3fd90ff98e5424a8a2a934d3c04471ee911e58bc378984387a1): complete subsection reference.

<a id="canonical-0cb3ada3bceb3087a79cd4c37577925d04b547acf0bffe2f1f93849a650b1c0f"></a>

## Next pages — palo_alto_fw_service.auto_setup / 927d7025fdee / 5

- [palo_alto_fw_service.auto_setup.admin_password](resources--nfv_service--reference--group-003.md#canonical-891e267f3e7873cc1cb3dc120f9b9997c1ff664c338ef0d59b130a8a1d61e525)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys](resources--nfv_service--reference--group-003.md#canonical-0fe0c13a624ad3fd90ff98e5424a8a2a934d3c04471ee911e58bc378984387a1)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-683b0e043ad34e16d220e00d2c0ad7d7794cd7a996943065814c81f0d29e03cf)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-891e267f3e7873cc1cb3dc120f9b9997c1ff664c338ef0d59b130a8a1d61e525"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce586e9bbd42661f34478ef6d8a525fb2dea8ad3ef550441d5090bf2f9888dab"></a>

## palo_alto_fw_service.auto_setup.admin_password — palo_alto_fw_service.auto_setup.admin_password / 71d06ec9dc3c / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-683b0e043ad34e16d220e00d2c0ad7d7794cd7a996943065814c81f0d29e03cf)
- [palo_alto_fw_service.auto_setup](resources--nfv_service--reference--group-003.md#canonical-7bf638e8f5f6d4cd41bc1908f771ac84af047472fffcf4f680086848cf0d43db)
- palo_alto_fw_service.auto_setup.admin_password

<a id="canonical-9ff11183e8534b6f1990b81542da2981c1f0b50fbcd6389be94307fc2d03df45"></a>

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
admin_password {
  # Configure direct properties listed below.
}
```

<a id="canonical-7feed1161e969a1d90845d876b41634dc1f4d590aa26a7cd1238518597eb177c"></a>

## Direct properties — palo_alto_fw_service.auto_setup.admin_password / 71d06ec9dc3c / 3

- [blindfold_secret_info](resources--nfv_service--reference--group-003.md#canonical-4c08bb00dd3e0a826cd126875cff4c852fe1f9b58e7cff679336409e60426241): complete subsection reference.

- [clear_secret_info](resources--nfv_service--reference--group-003.md#canonical-933940831149e24b2f0bf53ee0a1ff78334858bd2a271da51a3626efdf5ba8ac): complete subsection reference.

<a id="canonical-6e22f9017cc7ae996fbd89b52e5a3a902db4f108db94f8501454dec84e2dca8d"></a>

## Next pages — palo_alto_fw_service.auto_setup.admin_password / 71d06ec9dc3c / 4

- [palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info](resources--nfv_service--reference--group-003.md#canonical-4c08bb00dd3e0a826cd126875cff4c852fe1f9b58e7cff679336409e60426241)
- [palo_alto_fw_service.auto_setup.admin_password.clear_secret_info](resources--nfv_service--reference--group-003.md#canonical-933940831149e24b2f0bf53ee0a1ff78334858bd2a271da51a3626efdf5ba8ac)
- [palo_alto_fw_service.auto_setup](resources--nfv_service--reference--group-003.md#canonical-7bf638e8f5f6d4cd41bc1908f771ac84af047472fffcf4f680086848cf0d43db)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-4c08bb00dd3e0a826cd126875cff4c852fe1f9b58e7cff679336409e60426241"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16c8311b0809056a5c9bba4e7aef7564e5b95e7c4b9d1b3e2551bd6e587ecb8b"></a>

## palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info — palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info / 7b8ef6541a13 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-683b0e043ad34e16d220e00d2c0ad7d7794cd7a996943065814c81f0d29e03cf)
- [palo_alto_fw_service.auto_setup](resources--nfv_service--reference--group-003.md#canonical-7bf638e8f5f6d4cd41bc1908f771ac84af047472fffcf4f680086848cf0d43db)
- [palo_alto_fw_service.auto_setup.admin_password](resources--nfv_service--reference--group-003.md#canonical-891e267f3e7873cc1cb3dc120f9b9997c1ff664c338ef0d59b130a8a1d61e525)
- palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info

<a id="canonical-f6e42f77395b0a1d6a4308281d87c983b65e3f0ae8fd551fee485e9a01071b34"></a>

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

<a id="canonical-b45e5d50c3c297fb50a7cd398a508d4d19f61b3d6c2f8757baaa6654dab491ce"></a>

## Direct properties — palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info / 7b8ef6541a13 / 3

<a id="canonical-faef2b0ccbae06875480ff758ce62c6ef09fc7e072ba2dd1495a7b7320d34d7e"></a>

<a id="canonical-bef058b336c75b049f82e0f97c185f0e7844369605ee21ad652aec48a7b78d78"></a>

## decryption_provider property — palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info / 7b8ef6541a13 / 4

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

<a id="canonical-fcdbf12f1809220226f2a0f347a5cbae43a8f48e1e50a30776cd2293fad007c3"></a>

<a id="canonical-eb4e63877c285e4498f9022e5a33989fb0a785c5e19425e6c382402ec3e59297"></a>

## location property — palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info / 7b8ef6541a13 / 5

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

<a id="canonical-8b50636744e9c723c9a3917d98a8e8a5d57b9990af7acb7fa5b79950b5df2682"></a>

<a id="canonical-4789894806b5ef126c2e09cf74a4825d02b03624ce6535f6da58881c0644b8b1"></a>

## store_provider property — palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info / 7b8ef6541a13 / 6

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

<a id="canonical-77d3652ab3ab1ee56c414f16bee92de9b4e635d07a4732043ff6c511b25fa8b7"></a>

## Next pages — palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info / 7b8ef6541a13 / 7

- [palo_alto_fw_service.auto_setup.admin_password](resources--nfv_service--reference--group-003.md#canonical-891e267f3e7873cc1cb3dc120f9b9997c1ff664c338ef0d59b130a8a1d61e525)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-933940831149e24b2f0bf53ee0a1ff78334858bd2a271da51a3626efdf5ba8ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a2d4ef2db74c3f12c0cc2c22d0b66c94e19d99c6619da5014a6cc237b27e5e7"></a>

## palo_alto_fw_service.auto_setup.admin_password.clear_secret_info — palo_alto_fw_service.auto_setup.admin_password.clear_secret_info / 22f140092e40 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-683b0e043ad34e16d220e00d2c0ad7d7794cd7a996943065814c81f0d29e03cf)
- [palo_alto_fw_service.auto_setup](resources--nfv_service--reference--group-003.md#canonical-7bf638e8f5f6d4cd41bc1908f771ac84af047472fffcf4f680086848cf0d43db)
- [palo_alto_fw_service.auto_setup.admin_password](resources--nfv_service--reference--group-003.md#canonical-891e267f3e7873cc1cb3dc120f9b9997c1ff664c338ef0d59b130a8a1d61e525)
- palo_alto_fw_service.auto_setup.admin_password.clear_secret_info

<a id="canonical-ec221483893bc68fb5afe59fa02f92e3e1c2be08a15d72a437ad96858cc0decf"></a>

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

<a id="canonical-d30d506a4bc697ba1a98a01dad4fa150c05369d9c941553da0c9197f9503d3ba"></a>

## Direct properties — palo_alto_fw_service.auto_setup.admin_password.clear_secret_info / 22f140092e40 / 3

<a id="canonical-e51ded7440edba3af7b44f54fcd155f5abd33281a74b952a8646ddf16ab2575e"></a>

<a id="canonical-22285289bd704bbea641f0270d0ff54263e7104f516c6611e04b1b83a03818c4"></a>

## provider_ref property — palo_alto_fw_service.auto_setup.admin_password.clear_secret_info / 22f140092e40 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1ed5c2f73f14ebcda659153569e128cf9cc139a3833d14edd3f5a4552f658a8e"></a>

<a id="canonical-27d2b10156d1ba4a17b23552efda4719f0664c34dae6420a9394701d44b6f194"></a>

## url property — palo_alto_fw_service.auto_setup.admin_password.clear_secret_info / 22f140092e40 / 5

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

<a id="canonical-fd8b88f0887ef25dd0105474a0adad6a724f67e6a3e13b1369172321ea668dce"></a>

## Next pages — palo_alto_fw_service.auto_setup.admin_password.clear_secret_info / 22f140092e40 / 6

- [palo_alto_fw_service.auto_setup.admin_password](resources--nfv_service--reference--group-003.md#canonical-891e267f3e7873cc1cb3dc120f9b9997c1ff664c338ef0d59b130a8a1d61e525)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-0fe0c13a624ad3fd90ff98e5424a8a2a934d3c04471ee911e58bc378984387a1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b8f279180f0d491078720a18f673191c99861e21bc0dcf95dfefb4f34dddae2"></a>

## palo_alto_fw_service.auto_setup.manual_ssh_keys — palo_alto_fw_service.auto_setup.manual_ssh_keys / ef906543b8e9 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-683b0e043ad34e16d220e00d2c0ad7d7794cd7a996943065814c81f0d29e03cf)
- [palo_alto_fw_service.auto_setup](resources--nfv_service--reference--group-003.md#canonical-7bf638e8f5f6d4cd41bc1908f771ac84af047472fffcf4f680086848cf0d43db)
- palo_alto_fw_service.auto_setup.manual_ssh_keys

<a id="canonical-2f6a959a99036fbb247fe691ddea6f65fa4b170d7fde29be419061b84acf2c97"></a>

Type: `"object"`. single nested block, Optional.

SSH Key includes both public and private key.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("public_key")}
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
manual_ssh_keys {
  # Configure direct properties listed below.
}
```

<a id="canonical-3f123cac6e8ec342b28c3491b4567be1b4abe3579960becee09291f6c85c3c0e"></a>

## Direct properties — palo_alto_fw_service.auto_setup.manual_ssh_keys / ef906543b8e9 / 3

- [private_key](resources--nfv_service--reference--group-003.md#canonical-045c0321536770858b03978ae68fe304d694d5aec0fc64627c72a66bcbbc5095): complete subsection reference.

<a id="canonical-dd567610e9bf5188e3c958750c2069debbd5cb832d34499cf6ea3d297a3b17f9"></a>

<a id="canonical-483d1bcf804b88151016b8bdc33e7f4d95ef00823aa1e396247d505691a4846a"></a>

## public_key property — palo_alto_fw_service.auto_setup.manual_ssh_keys / ef906543b8e9 / 4

Type: `"string"`. Optional.

Authorized Public SSH key which will be programmed on the node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 8192),
}
```

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

<a id="canonical-34e32671aedf814c676a9090f6b7dc21ecfdece5ec1dc9c1c0eb93bc64736fed"></a>

## Next pages — palo_alto_fw_service.auto_setup.manual_ssh_keys / ef906543b8e9 / 5

- [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key](resources--nfv_service--reference--group-003.md#canonical-045c0321536770858b03978ae68fe304d694d5aec0fc64627c72a66bcbbc5095)
- [palo_alto_fw_service.auto_setup](resources--nfv_service--reference--group-003.md#canonical-7bf638e8f5f6d4cd41bc1908f771ac84af047472fffcf4f680086848cf0d43db)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-045c0321536770858b03978ae68fe304d694d5aec0fc64627c72a66bcbbc5095"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2704e68791a3e77613afc397fea5141071a3830a52ce9d3fe9083b79d35e556e"></a>

## palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key — palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key / 2792bbab69c8 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-683b0e043ad34e16d220e00d2c0ad7d7794cd7a996943065814c81f0d29e03cf)
- [palo_alto_fw_service.auto_setup](resources--nfv_service--reference--group-003.md#canonical-7bf638e8f5f6d4cd41bc1908f771ac84af047472fffcf4f680086848cf0d43db)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys](resources--nfv_service--reference--group-003.md#canonical-0fe0c13a624ad3fd90ff98e5424a8a2a934d3c04471ee911e58bc378984387a1)
- palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key

<a id="canonical-be9c868da4cc8fcd2dec68114743f65c9b28ac5ef6d6daf01a9cef6f2944d4a1"></a>

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
