---
page_title: "xcsh_azure_vnet_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_azure_vnet_site reference."
---

# xcsh_azure_vnet_site reference

<a id="canonical-a515da2f61734d10fed07e7007359bcf3405dc332f39f56427a7e79057a6acb2"></a>

## Next pages — ingress_gw_ar.node.local_subnet / 6f38c925dfdc / 4

- [ingress_gw_ar.node.local_subnet.subnet](resources--azure_vnet_site--reference--group-008.md#canonical-589e6304a3b881158b62acb9f1ffbda5e070f464ca55ad8236f02c18a58964a8)
- [ingress_gw_ar.node.local_subnet.subnet_param](resources--azure_vnet_site--reference--group-008.md#canonical-5278ceccd573239486fadf316b56502fac6d8825a9f4cea180295e9fc41d4c75)
- [ingress_gw_ar.node](resources--azure_vnet_site--reference--group-007.md#canonical-87b70f7d188d23ff41722d94e8837588c63ec02ab72058e0fcb45f7d81e73a69)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-589e6304a3b881158b62acb9f1ffbda5e070f464ca55ad8236f02c18a58964a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-47487497443356d7e8ba5790d678b0ab94d2df50b5f054062b61a937efe38828"></a>

## ingress_gw_ar.node.local_subnet.subnet — ingress_gw_ar.node.local_subnet.subnet / 512f426231aa / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_gw_ar](resources--azure_vnet_site--reference--group-007.md#canonical-d0e215b74519793a40bfc7211040a95320d415ace2a874843eeb2ee665a518c4)
- [ingress_gw_ar.node](resources--azure_vnet_site--reference--group-007.md#canonical-87b70f7d188d23ff41722d94e8837588c63ec02ab72058e0fcb45f7d81e73a69)
- [ingress_gw_ar.node.local_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-c40a230559356b8bdb345774131fdd36f2125524a601a66842dee698de5a0873)
- ingress_gw_ar.node.local_subnet.subnet

<a id="canonical-be40192ae7756a115f3ecc56d63b49872f9b3339c0c809895f7a1bb4f0e87f38"></a>

Type: `"object"`. single nested block, Optional.

Subnet specification for network segmentation.

Upstream description:

Parameters for Azure subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subnet_name"),
  validators.ConflictingObjectAttributes("subnet_resource_grp",
    "vnet_resource_group")}
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
  "x-ves-oneof-field-resource_group_choice": "[\"subnet_resource_grp\",\"vnet_resource_group\"]"
}
```

Terraform syntax:

```terraform
subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-5c37c0e9f243ebe31eb4922c75b580c84d4758f5ad5179d32734a4cde96cc509"></a>

## Direct properties — ingress_gw_ar.node.local_subnet.subnet / 512f426231aa / 3

<a id="canonical-737a21aae476f9c9b5015d12f6160d2a1fad81419fb522ec6121745af9d044d6"></a>

<a id="canonical-b80c84d9f064680ecf5d7eb387473e9fc9e6cdcaa08c2eb5e57d16c7fdb57ec8"></a>

## subnet_name property — ingress_gw_ar.node.local_subnet.subnet / 512f426231aa / 4

Type: `"string"`. Optional.

Subnet Name. Name of existing subnet.

Upstream description:

Name of existing subnet.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-6595d7965f3801761a73b31207a2bc883cb066097b448d19fd6f3c3fbeee7272"></a>

<a id="canonical-bd6b04cb934c340c054446d008da390940ed1838ffe60851abf51b542545b5cc"></a>

## subnet_resource_grp property — ingress_gw_ar.node.local_subnet.subnet / 512f426231aa / 5

Type: `"string"`. Optional.

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Upstream description:

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [vnet_resource_group](resources--azure_vnet_site--reference--group-008.md#canonical-91edc28ed84f6b2688ef80b6218624950b072cb97206a20addd39a9784f73eb3): complete subsection reference.

<a id="canonical-622bf1949afc34126429ed8d69b15f8ab1b56d1e404b8fc2e486c8194954df4c"></a>

## Next pages — ingress_gw_ar.node.local_subnet.subnet / 512f426231aa / 6

- [ingress_gw_ar.node.local_subnet.subnet.vnet_resource_group](resources--azure_vnet_site--reference--group-008.md#canonical-91edc28ed84f6b2688ef80b6218624950b072cb97206a20addd39a9784f73eb3)
- [ingress_gw_ar.node.local_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-c40a230559356b8bdb345774131fdd36f2125524a601a66842dee698de5a0873)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-91edc28ed84f6b2688ef80b6218624950b072cb97206a20addd39a9784f73eb3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2f4cb0fc9fe917e42220ef1d1a86384842007839fbbcd5a6af35fe4a853cc277"></a>

## ingress_gw_ar.node.local_subnet.subnet.vnet_resource_group — ingress_gw_ar.node.local_subnet.subnet.vnet_resource_group / 6a8f3b914997 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_gw_ar](resources--azure_vnet_site--reference--group-007.md#canonical-d0e215b74519793a40bfc7211040a95320d415ace2a874843eeb2ee665a518c4)
- [ingress_gw_ar.node](resources--azure_vnet_site--reference--group-007.md#canonical-87b70f7d188d23ff41722d94e8837588c63ec02ab72058e0fcb45f7d81e73a69)
- [ingress_gw_ar.node.local_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-c40a230559356b8bdb345774131fdd36f2125524a601a66842dee698de5a0873)
- [ingress_gw_ar.node.local_subnet.subnet](resources--azure_vnet_site--reference--group-008.md#canonical-589e6304a3b881158b62acb9f1ffbda5e070f464ca55ad8236f02c18a58964a8)
- ingress_gw_ar.node.local_subnet.subnet.vnet_resource_group

<a id="canonical-39249ff2ef89d0d3e89070d0cefa34503603be5ace8d12ed0eb2de33eaabfd07"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for vnet resource group.

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
vnet_resource_group = {}
```

<a id="canonical-9b60ec87cf246f8bff47bde63bddc0f568a724a695da35c0217844d0ccd5c2fe"></a>

## Direct properties — ingress_gw_ar.node.local_subnet.subnet.vnet_resource_group / 6a8f3b914997 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a5632cbdae9e701c2bb6d23d7c5de19663375ed98f344eae8192690f526a5579"></a>

## Next pages — ingress_gw_ar.node.local_subnet.subnet.vnet_resource_group / 6a8f3b914997 / 4

- [ingress_gw_ar.node.local_subnet.subnet](resources--azure_vnet_site--reference--group-008.md#canonical-589e6304a3b881158b62acb9f1ffbda5e070f464ca55ad8236f02c18a58964a8)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-5278ceccd573239486fadf316b56502fac6d8825a9f4cea180295e9fc41d4c75"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4feb8fb9108459979f22a0ea002cf70f68562b515dfce1fe5f8c2f02828d198b"></a>

## ingress_gw_ar.node.local_subnet.subnet_param — ingress_gw_ar.node.local_subnet.subnet_param / 64e9d6d38573 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_gw_ar](resources--azure_vnet_site--reference--group-007.md#canonical-d0e215b74519793a40bfc7211040a95320d415ace2a874843eeb2ee665a518c4)
- [ingress_gw_ar.node](resources--azure_vnet_site--reference--group-007.md#canonical-87b70f7d188d23ff41722d94e8837588c63ec02ab72058e0fcb45f7d81e73a69)
- [ingress_gw_ar.node.local_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-c40a230559356b8bdb345774131fdd36f2125524a601a66842dee698de5a0873)
- ingress_gw_ar.node.local_subnet.subnet_param

<a id="canonical-8d9aaea6eb3d44744139bdf1e45b8d8bc9fca32abff858b53cf73c38ec0ea94d"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating a new cloud subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ipv4")}
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
subnet_param {
  # Configure direct properties listed below.
}
```

<a id="canonical-f23a20cc521aeba5b308a4623e67cbe1aa2e0168cb705d596f9341c58b12df82"></a>

## Direct properties — ingress_gw_ar.node.local_subnet.subnet_param / 64e9d6d38573 / 3

<a id="canonical-494a56c1e42ca7f6070fd3a7eeef06b73e33d018ee1e69c25a3cd4c66392de4b"></a>

<a id="canonical-96d3fe79dfe8b250b730608f9109b5dee11cda1400feece375c7ebde09950802"></a>

## ipv4 property — ingress_gw_ar.node.local_subnet.subnet_param / 64e9d6d38573 / 4

Type: `"string"`. Optional.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-54a8efdabc2b250c384df88c09cbe2e0a7ceadcd5452cf6b59c6d9c46e8a005b"></a>

## Next pages — ingress_gw_ar.node.local_subnet.subnet_param / 64e9d6d38573 / 5

- [ingress_gw_ar.node.local_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-c40a230559356b8bdb345774131fdd36f2125524a601a66842dee698de5a0873)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-3b7cf54e88a203c51b49b91964fd227cf387be5f726e62633846067d9131bbd2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-548e7e247f294b5e53e470a80e93ec45b2c360e9945cc5acec520fbfb2f34e0a"></a>

## ingress_gw_ar.performance_enhancement_mode — ingress_gw_ar.performance_enhancement_mode / b21df2575f66 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_gw_ar](resources--azure_vnet_site--reference--group-007.md#canonical-d0e215b74519793a40bfc7211040a95320d415ace2a874843eeb2ee665a518c4)
- ingress_gw_ar.performance_enhancement_mode

<a id="canonical-3e633eb2abcb351b0fdfc9df342f3dda728bd154aff9113a3efdeaae7986c52a"></a>

Type: `"object"`. single nested block, Optional.

Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("perf_mode_l3_enhanced",
    "perf_mode_l7_enhanced")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"perf_mode_l3_enhanced\",\"perf_mode_l7_enhanced\"]"
}
```

Terraform syntax:

```terraform
performance_enhancement_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-79ce5721683028513215bdff1b86e87d086ed09cf035f76e0bfc75e0c9eea6c0"></a>

## Direct properties — ingress_gw_ar.performance_enhancement_mode / b21df2575f66 / 3

- [perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-008.md#canonical-923930b3233972ffa841ce4d53f17ec0e451a9f4e6fae40f2cd30ceb1ee70d12): complete subsection reference.

- [perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-008.md#canonical-3c6a2b2b56c3b69b0f52970678d5afdbd2032d8d1929cde2915505286c75a8d7): complete subsection reference.

<a id="canonical-94039708c0fe1a17c55ce26302888cf2df7de411e1bbfc988ef562d8419cd817"></a>

## Next pages — ingress_gw_ar.performance_enhancement_mode / b21df2575f66 / 4

- [ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-008.md#canonical-923930b3233972ffa841ce4d53f17ec0e451a9f4e6fae40f2cd30ceb1ee70d12)
- [ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-008.md#canonical-3c6a2b2b56c3b69b0f52970678d5afdbd2032d8d1929cde2915505286c75a8d7)
- [ingress_gw_ar](resources--azure_vnet_site--reference--group-007.md#canonical-d0e215b74519793a40bfc7211040a95320d415ace2a874843eeb2ee665a518c4)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-923930b3233972ffa841ce4d53f17ec0e451a9f4e6fae40f2cd30ceb1ee70d12"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223ea846364e7008c1e1009e545323d4a4c708be4434c29c3f1a585ab539069"></a>

## ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced — ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced / 121aff688648 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_gw_ar](resources--azure_vnet_site--reference--group-007.md#canonical-d0e215b74519793a40bfc7211040a95320d415ace2a874843eeb2ee665a518c4)
- [ingress_gw_ar.performance_enhancement_mode](resources--azure_vnet_site--reference--group-008.md#canonical-3b7cf54e88a203c51b49b91964fd227cf387be5f726e62633846067d9131bbd2)
- ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-6aa06de546b2143523aaecb98d6c096bb7d6f0ce741abdab9694a8dab464d506"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for perf mode l3 enhanced.

Upstream description:

L3 enhanced performance mode OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("jumbo",
    "no_jumbo")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo\",\"no_jumbo\"]"
}
```

Terraform syntax:

```terraform
perf_mode_l3_enhanced {
  # Configure direct properties listed below.
}
```

<a id="canonical-6d2a1f0a71b415d4732b3d6916bb09b2906e49b929b7bc86df8f803c4d32b6b6"></a>

## Direct properties — ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced / 121aff688648 / 3

- [jumbo](resources--azure_vnet_site--reference--group-008.md#canonical-eb09277ed0a8eb684912b89f4e45ba44ef30e7ba53efd3de77a835523ff9f327): complete subsection reference.

- [no_jumbo](resources--azure_vnet_site--reference--group-008.md#canonical-67e72c5919229811af37fba336392ab5b97a4e20376fb4a934d4ce712e3b4a5b): complete subsection reference.

<a id="canonical-e3825295e70c311b126bf6460a244b0b6ef9781c421369c72b1ca50c260a3ac2"></a>

## Next pages — ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced / 121aff688648 / 4

- [ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--azure_vnet_site--reference--group-008.md#canonical-eb09277ed0a8eb684912b89f4e45ba44ef30e7ba53efd3de77a835523ff9f327)
- [ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--azure_vnet_site--reference--group-008.md#canonical-67e72c5919229811af37fba336392ab5b97a4e20376fb4a934d4ce712e3b4a5b)
- [ingress_gw_ar.performance_enhancement_mode](resources--azure_vnet_site--reference--group-008.md#canonical-3b7cf54e88a203c51b49b91964fd227cf387be5f726e62633846067d9131bbd2)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-eb09277ed0a8eb684912b89f4e45ba44ef30e7ba53efd3de77a835523ff9f327"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2382ae72a413f1dfe1ddb558c2ffc980aa6c18a561df928717f4d6dfe094a91a"></a>

## ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo — ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / 0cc72d6a21ff / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_gw_ar](resources--azure_vnet_site--reference--group-007.md#canonical-d0e215b74519793a40bfc7211040a95320d415ace2a874843eeb2ee665a518c4)
- [ingress_gw_ar.performance_enhancement_mode](resources--azure_vnet_site--reference--group-008.md#canonical-3b7cf54e88a203c51b49b91964fd227cf387be5f726e62633846067d9131bbd2)
- [ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-008.md#canonical-923930b3233972ffa841ce4d53f17ec0e451a9f4e6fae40f2cd30ceb1ee70d12)
- ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-1eee9f20bd4bea23f328226e84e0bbfc4848fea1c05b3fe9c59204bf14aa4ec3"></a>

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
jumbo = {}
```

<a id="canonical-7ecf405825879916b69760d4518a90de3317d4af95cca8cecd6aa496879942c7"></a>

## Direct properties — ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / 0cc72d6a21ff / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4f36e9a446adcdc529d37a62ff2d27f8c9066e28ed49d323dee6a344ceb04cc3"></a>

## Next pages — ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / 0cc72d6a21ff / 4

- [ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-008.md#canonical-923930b3233972ffa841ce4d53f17ec0e451a9f4e6fae40f2cd30ceb1ee70d12)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-67e72c5919229811af37fba336392ab5b97a4e20376fb4a934d4ce712e3b4a5b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd2d16f6df85eefac0af61ec470838f50e5ffdf507adac286c12227d63920233"></a>

## ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo — ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 8faabea0a807 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_gw_ar](resources--azure_vnet_site--reference--group-007.md#canonical-d0e215b74519793a40bfc7211040a95320d415ace2a874843eeb2ee665a518c4)
- [ingress_gw_ar.performance_enhancement_mode](resources--azure_vnet_site--reference--group-008.md#canonical-3b7cf54e88a203c51b49b91964fd227cf387be5f726e62633846067d9131bbd2)
- [ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-008.md#canonical-923930b3233972ffa841ce4d53f17ec0e451a9f4e6fae40f2cd30ceb1ee70d12)
- ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="canonical-5d7dc281bc441c09c1d053e60204a0818d0cddc32404bb290d6ad65a517b813b"></a>

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
no_jumbo = {}
```

<a id="canonical-5787d5065596da588113f2d26b0cc5d5a42caf5218419e9b072fc22f4d4e8078"></a>

## Direct properties — ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 8faabea0a807 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4c783ff922406900ce3e7ae5ab7192b5fd5e9334ac10daeb30163bbfe050586e"></a>

## Next pages — ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 8faabea0a807 / 4

- [ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-008.md#canonical-923930b3233972ffa841ce4d53f17ec0e451a9f4e6fae40f2cd30ceb1ee70d12)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-3c6a2b2b56c3b69b0f52970678d5afdbd2032d8d1929cde2915505286c75a8d7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-80f9699b4bde0153162b7f071eff6c4d7d1563e447b54b5ee8b664ca1a778ec9"></a>

## ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced — ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced / 3b6b19d52e5a / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_gw_ar](resources--azure_vnet_site--reference--group-007.md#canonical-d0e215b74519793a40bfc7211040a95320d415ace2a874843eeb2ee665a518c4)
- [ingress_gw_ar.performance_enhancement_mode](resources--azure_vnet_site--reference--group-008.md#canonical-3b7cf54e88a203c51b49b91964fd227cf387be5f726e62633846067d9131bbd2)
- ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced

<a id="canonical-282e0f4cab1078dc0f5d0e17a79c4370c56346c04130a319eb383361aa61b6ba"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for perf mode l7 enhanced.

Upstream description:

L7 enhanced performance mode OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("jumbo_disabled",
    "jumbo_enabled")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo_disabled\",\"jumbo_enabled\"]"
}
```

Terraform syntax:

```terraform
perf_mode_l7_enhanced {
  # Configure direct properties listed below.
}
```

<a id="canonical-b77da55d9c4d60f7e809c8ffc67df8b1d4f690e9d95fc8dc738ddf59731cad3b"></a>

## Direct properties — ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced / 3b6b19d52e5a / 3

- [jumbo_disabled](resources--azure_vnet_site--reference--group-008.md#canonical-b45f864d2dbb45faae69e8387d001f7a440cfb09edfabfe30f258f1f586efaa3): complete subsection reference.

- [jumbo_enabled](resources--azure_vnet_site--reference--group-008.md#canonical-a077d926679f68ec1cda521855dd26b2fad1ed77c8564bcdb80704c0027945cc): complete subsection reference.

<a id="canonical-2f2e60990822b396e6a98fa5eee3c2ecc3a381c5e3f78a4a57934bfe29714acc"></a>

## Next pages — ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced / 3b6b19d52e5a / 4

- [ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--azure_vnet_site--reference--group-008.md#canonical-b45f864d2dbb45faae69e8387d001f7a440cfb09edfabfe30f258f1f586efaa3)
- [ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--azure_vnet_site--reference--group-008.md#canonical-a077d926679f68ec1cda521855dd26b2fad1ed77c8564bcdb80704c0027945cc)
- [ingress_gw_ar.performance_enhancement_mode](resources--azure_vnet_site--reference--group-008.md#canonical-3b7cf54e88a203c51b49b91964fd227cf387be5f726e62633846067d9131bbd2)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-b45f864d2dbb45faae69e8387d001f7a440cfb09edfabfe30f258f1f586efaa3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a0ea82810d5be21a1358907d70979799001e289221733803c27c1d68e4f7ffff"></a>

## ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled — ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled / 0f1e50712809 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_gw_ar](resources--azure_vnet_site--reference--group-007.md#canonical-d0e215b74519793a40bfc7211040a95320d415ace2a874843eeb2ee665a518c4)
- [ingress_gw_ar.performance_enhancement_mode](resources--azure_vnet_site--reference--group-008.md#canonical-3b7cf54e88a203c51b49b91964fd227cf387be5f726e62633846067d9131bbd2)
- [ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-008.md#canonical-3c6a2b2b56c3b69b0f52970678d5afdbd2032d8d1929cde2915505286c75a8d7)
- ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled

<a id="canonical-fb668e94264a7a43603132efe2febf6c68456f2d799aff7fa9bfd12d837e42b5"></a>

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
jumbo_disabled = {}
```

<a id="canonical-124e258a7aed2bea7221f24b98dcaef9478f3e69cbac52e3504be94224b6af17"></a>

## Direct properties — ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled / 0f1e50712809 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7fa871b9840fd4ee612cefe81d57e340b65b848394e54fc02ecd6a2a21c1394f"></a>

## Next pages — ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled / 0f1e50712809 / 4

- [ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-008.md#canonical-3c6a2b2b56c3b69b0f52970678d5afdbd2032d8d1929cde2915505286c75a8d7)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-a077d926679f68ec1cda521855dd26b2fad1ed77c8564bcdb80704c0027945cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b3044d4a874892af18eaaddda02167bc10c031d856012e06454d49d35a64256"></a>

## ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled — ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled / a31dda0351f5 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_gw_ar](resources--azure_vnet_site--reference--group-007.md#canonical-d0e215b74519793a40bfc7211040a95320d415ace2a874843eeb2ee665a518c4)
- [ingress_gw_ar.performance_enhancement_mode](resources--azure_vnet_site--reference--group-008.md#canonical-3b7cf54e88a203c51b49b91964fd227cf387be5f726e62633846067d9131bbd2)
- [ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-008.md#canonical-3c6a2b2b56c3b69b0f52970678d5afdbd2032d8d1929cde2915505286c75a8d7)
- ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

<a id="canonical-ef425a95f92bd598524adb7a81ab2f9083a2f2e4b56aa5320137f328fb1ed876"></a>

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
jumbo_enabled = {}
```

<a id="canonical-5132a792ab7412fb4bdf35fb6fafd585b1336e401560acd452bac97013817cf8"></a>

## Direct properties — ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled / a31dda0351f5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b376ed7451aa80655e14512cd9b784d643756bce5ded7a29d3c9f502f1a50b0a"></a>

## Next pages — ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled / a31dda0351f5 / 4

- [ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-008.md#canonical-3c6a2b2b56c3b69b0f52970678d5afdbd2032d8d1929cde2915505286c75a8d7)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-b5c1af0f284a4e0425831c6c3358f56df9913afd1d87a0cde4ebdc2d992213f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d333335a709ad1478d14e81500e903a19d5a32aacaac665ef59aba72dfb08b2"></a>

## kubernetes_upgrade_drain — kubernetes_upgrade_drain / 2ecd35d7c57b / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- kubernetes_upgrade_drain

<a id="canonical-3b7a4fe6b6669f3eedeaf046ef82c9632f1c7b47b913c059c9105fa44a8d8095"></a>

Type: `"object"`. single nested block, Optional.

Specify how worker nodes within a site will be upgraded.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_upgrade_drain",
    "enable_upgrade_drain")}
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
  "x-ves-oneof-field-kubernetes_upgrade_drain_enable_choice": "[\"disable_upgrade_drain\",\"enable_upgrade_drain\"]"
}
```

Terraform syntax:

```terraform
kubernetes_upgrade_drain {
  # Configure direct properties listed below.
}
```

<a id="canonical-6c5246ebc6a94849dce9f10c5fc315abdaeb99a808d2a9713918f5829e2f0a31"></a>

## Direct properties — kubernetes_upgrade_drain / 2ecd35d7c57b / 3

- [disable_upgrade_drain](resources--azure_vnet_site--reference--group-008.md#canonical-02a34c1201382ab5f3b9eaaac21fd1e1d97e2557ad49c1965f7f9ea81181ad08): complete subsection reference.

- [enable_upgrade_drain](resources--azure_vnet_site--reference--group-008.md#canonical-7586154b636be221965b978d94ff361b752cc69fa951865b14fdf878c6b21256): complete subsection reference.

<a id="canonical-2c41b71067c7e6ec59d378fc7c8e961d973f875af165ec01ae4f6521991f1a6b"></a>

## Next pages — kubernetes_upgrade_drain / 2ecd35d7c57b / 4

- [kubernetes_upgrade_drain.disable_upgrade_drain](resources--azure_vnet_site--reference--group-008.md#canonical-02a34c1201382ab5f3b9eaaac21fd1e1d97e2557ad49c1965f7f9ea81181ad08)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--azure_vnet_site--reference--group-008.md#canonical-7586154b636be221965b978d94ff361b752cc69fa951865b14fdf878c6b21256)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-02a34c1201382ab5f3b9eaaac21fd1e1d97e2557ad49c1965f7f9ea81181ad08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-834029c99fd53b085f7524e9dce749b4b1e9361f7d64588de2de5252aa1f167b"></a>

## kubernetes_upgrade_drain.disable_upgrade_drain — kubernetes_upgrade_drain.disable_upgrade_drain / e6e87268f2ff / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [kubernetes_upgrade_drain](resources--azure_vnet_site--reference--group-008.md#canonical-b5c1af0f284a4e0425831c6c3358f56df9913afd1d87a0cde4ebdc2d992213f0)
- kubernetes_upgrade_drain.disable_upgrade_drain

<a id="canonical-c3cb5207fd67331f03520691a5b427d73e1b48ed1cfc54f5a590cf5cbfe0f66c"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable upgrade drain.

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
disable_upgrade_drain = {}
```

<a id="canonical-b23bbdb73c0408d41ee93662e4f451b4e0e44b2007342db1a70839903e60c4da"></a>

## Direct properties — kubernetes_upgrade_drain.disable_upgrade_drain / e6e87268f2ff / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-51688cadea7a07a4fb71522a9ca6aa03d30661b284fd88b307ce1709a12e5ff8"></a>

## Next pages — kubernetes_upgrade_drain.disable_upgrade_drain / e6e87268f2ff / 4

- [kubernetes_upgrade_drain](resources--azure_vnet_site--reference--group-008.md#canonical-b5c1af0f284a4e0425831c6c3358f56df9913afd1d87a0cde4ebdc2d992213f0)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-7586154b636be221965b978d94ff361b752cc69fa951865b14fdf878c6b21256"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b5dbc414be142f66a9874b3138b8d685d7d43d5364511a2e559d5a5633df8577"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain — kubernetes_upgrade_drain.enable_upgrade_drain / a174d56cf716 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [kubernetes_upgrade_drain](resources--azure_vnet_site--reference--group-008.md#canonical-b5c1af0f284a4e0425831c6c3358f56df9913afd1d87a0cde4ebdc2d992213f0)
- kubernetes_upgrade_drain.enable_upgrade_drain

<a id="canonical-e1cef55cea669b5cd798133258a3a2291309003469372bb8adceb8e25a45fcb0"></a>

Type: `"object"`. single nested block, Optional.

Specify batch upgrade settings for worker nodes within a site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("drain_node_timeout"),
  validators.ConflictingObjectAttributes("disable_vega_upgrade_mode",
    "enable_vega_upgrade_mode"),
  validators.ConflictingObjectAttributes("drain_max_unavailable_node_count",
    "drain_max_unavailable_node_percentage")}
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
  "x-ves-oneof-field-drain_max_unavailable_choice": "[\"drain_max_unavailable_node_count\", \"drain_max_unavailable_node_percentage\"]",
  "x-ves-oneof-field-vega_upgrade_mode_toggle_choice": "[\"disable_vega_upgrade_mode\",\"enable_vega_upgrade_mode\"]"
}
```

Terraform syntax:

```terraform
enable_upgrade_drain {
  # Configure direct properties listed below.
}
```

<a id="canonical-513e04798f3c25768e1406b1e8d7397ad0822f15d5a8e7d78e254d62d6dc522c"></a>

## Direct properties — kubernetes_upgrade_drain.enable_upgrade_drain / a174d56cf716 / 3

- [disable_vega_upgrade_mode](resources--azure_vnet_site--reference--group-008.md#canonical-f7f3cdd3187ad02b13272cfc811a861abad23eec052e02d7275e6eeb1365517d): complete subsection reference.

<a id="canonical-2b99f75607be86f732503d293864294ed34d18f605185d3579c4c3077bedfdd8"></a>

<a id="canonical-b1e63d5e6298fda7a468324b7a11d4aa49a76de848178b07278ee372e3a6bcfd"></a>

## drain_max_unavailable_node_count property — kubernetes_upgrade_drain.enable_upgrade_drain / a174d56cf716 / 4

Type: `"number"`. Optional.

Node Batch Size Count. Exclusive with \[\]

Upstream description:

Exclusive with \[\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 5000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5000,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5000"
  }
}
```

<a id="canonical-b3ad5f1b914da16d0a00117bbcdfcb7355c6e36a3774ae26d3653c0c3fc429c1"></a>

<a id="canonical-29b831dea70c963fa4443f19db6ab2d90ce8b359b48aa81a13a4c6fede13e073"></a>

## drain_max_unavailable_node_percentage property — kubernetes_upgrade_drain.enable_upgrade_drain / a174d56cf716 / 5

Type: `"number"`. Optional.

Maximum percentage of nodes unavailable during upgrade draining.

<a id="canonical-aebb56c2f32d0bc66f9a28a7a0134f5961406b39416552671f5a85ac86cd820c"></a>

<a id="canonical-50e8dc4dba5b3d799f8f176a86779eef5a4804679675e391f4510fa029207c3e"></a>

## drain_node_timeout property — kubernetes_upgrade_drain.enable_upgrade_drain / a174d56cf716 / 6

Type: `"number"`. Optional.

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It
is..

Upstream description:

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It is
recommended to use the default value).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 900),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 900,
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
    "ves.io.schema.rules.uint32.lte": "900"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "900"
  }
}
```

- [enable_vega_upgrade_mode](resources--azure_vnet_site--reference--group-008.md#canonical-9db5812bd60b468847bffecba4d9e0eeba116aa5393a9c9f77f9cccaffdbc8f5): complete subsection reference.

<a id="canonical-3645bb321affc2e5b8ea59da20ebf168f232218cea8616ddf8e06a55079e68ae"></a>

## Next pages — kubernetes_upgrade_drain.enable_upgrade_drain / a174d56cf716 / 7

- [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](resources--azure_vnet_site--reference--group-008.md#canonical-f7f3cdd3187ad02b13272cfc811a861abad23eec052e02d7275e6eeb1365517d)
- [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](resources--azure_vnet_site--reference--group-008.md#canonical-9db5812bd60b468847bffecba4d9e0eeba116aa5393a9c9f77f9cccaffdbc8f5)
- [kubernetes_upgrade_drain](resources--azure_vnet_site--reference--group-008.md#canonical-b5c1af0f284a4e0425831c6c3358f56df9913afd1d87a0cde4ebdc2d992213f0)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-f7f3cdd3187ad02b13272cfc811a861abad23eec052e02d7275e6eeb1365517d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-095cda919960131d13acbe386b1614a757737b2ff18df51cf02f7d890a84f9a0"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode — kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode / 1a655d64d8b1 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [kubernetes_upgrade_drain](resources--azure_vnet_site--reference--group-008.md#canonical-b5c1af0f284a4e0425831c6c3358f56df9913afd1d87a0cde4ebdc2d992213f0)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--azure_vnet_site--reference--group-008.md#canonical-7586154b636be221965b978d94ff361b752cc69fa951865b14fdf878c6b21256)
- kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode

<a id="canonical-d12030fe8dc106ac29990e718f128f1e1c9d622e752cc57fbe388366c606af57"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable vega upgrade mode.

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
disable_vega_upgrade_mode = {}
```

<a id="canonical-4c46beb8839a8d137f6625c31d01df3066202a95e410568d2994b15f3e501807"></a>

## Direct properties — kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode / 1a655d64d8b1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0f8d892ddb3c31caa46d4b08f776690484d9942a69fa4460ef3a282f12c8ecb3"></a>

## Next pages — kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode / 1a655d64d8b1 / 4

- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--azure_vnet_site--reference--group-008.md#canonical-7586154b636be221965b978d94ff361b752cc69fa951865b14fdf878c6b21256)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-9db5812bd60b468847bffecba4d9e0eeba116aa5393a9c9f77f9cccaffdbc8f5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-064fd052aa01ba140480b084907ced15980bc5360a2c401ef18613c71ddfe017"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode — kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode / 6f381819c664 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [kubernetes_upgrade_drain](resources--azure_vnet_site--reference--group-008.md#canonical-b5c1af0f284a4e0425831c6c3358f56df9913afd1d87a0cde4ebdc2d992213f0)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--azure_vnet_site--reference--group-008.md#canonical-7586154b636be221965b978d94ff361b752cc69fa951865b14fdf878c6b21256)
- kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode

<a id="canonical-6b6c2c70feba2f692b66d6af1cbded610ba896ce0d470498db8c16b90303b299"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable vega upgrade mode.

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
enable_vega_upgrade_mode = {}
```

<a id="canonical-4ee84bf16280a7519c1a7aa52b233038d0bb8c3e04d3a4b478ce8d15b1878475"></a>

## Direct properties — kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode / 6f381819c664 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1ca6deb4798f7087d2206ddfa21e2f951eef8b331c410b9f15b87aae063410c1"></a>

## Next pages — kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode / 6f381819c664 / 4

- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--azure_vnet_site--reference--group-008.md#canonical-7586154b636be221965b978d94ff361b752cc69fa951865b14fdf878c6b21256)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-56780c660093e40a35bccfb572d45434d392039269f7dfe6edd188bdde47f2aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-02710de725e1c2361e4e57bd729a31a20d06ce35172f80f7e22d8fae7ab416b9"></a>

## log_receiver — log_receiver / 5f9af07f3223 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- log_receiver

<a id="canonical-8b67b58809ea49fd5737d4c55904cf3d92c3d733e5cf30f1b5638ab47a07f778"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: log\_receiver, logs\_streaming\_disabled\] Type establishes a direct reference from one
object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.

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

OneOf alternatives in this subsection:

- [log_receiver](resources--azure_vnet_site--reference--group-008.md#canonical-8b67b58809ea49fd5737d4c55904cf3d92c3d733e5cf30f1b5638ab47a07f778)
- [logs_streaming_disabled](resources--azure_vnet_site--reference--group-008.md#canonical-1c606de63e0ddc324653afb9cba28bbf9e1af1dca9b8fbbedaedab81e2f2e2e9)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
log_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-60c377a488bffb7c8c73df3eebcaae9af35bb7e6ea44e0710493afcc6cae1848"></a>

## Direct properties — log_receiver / 5f9af07f3223 / 3

<a id="canonical-10fe2cdb1b59238f47c02c3a796d3f72b10fe48d5e3c8a65746519a877f9b4ec"></a>

<a id="canonical-2d12e72e98b0fa5ec29dd4de5717b4737aaa904df715295ba5080323f202f026"></a>

## name property — log_receiver / 5f9af07f3223 / 4

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

<a id="canonical-ef82e6661edec8b9e5e1dfc4d5be75a266615b732bafe830705aa93fb7fc06e5"></a>

<a id="canonical-292e4e8e7f478e5de90d4e8010ce69f526bb25f014071a1c082b33705f081496"></a>

## namespace property — log_receiver / 5f9af07f3223 / 5

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

<a id="canonical-fc79de369dcf18c9c80bae53c7da9988d7a7967a8c6bc660fcb7e40858177916"></a>

<a id="canonical-5ead8bfafc24443ff687ef060ce40565c6da77d062ac6fae4109663ebce2ef7c"></a>

## tenant property — log_receiver / 5f9af07f3223 / 6

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

<a id="canonical-ac02e3a588895d11f414bfdb5ee5c65e7b19512dd23c05440ec664ad467a392c"></a>

## Next pages — log_receiver / 5f9af07f3223 / 7

- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-25258cc231a9d4026b2cde1f1ec58b3aee1f285ea58854a696755c09feb1df65"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-01efcf47888e8a5d642ddcb6ad0ba2a391b01538bd0de56d5bf650174b98e7cb"></a>

## logs_streaming_disabled — logs_streaming_disabled / 2b933d374f9e / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- logs_streaming_disabled

<a id="canonical-1c606de63e0ddc324653afb9cba28bbf9e1af1dca9b8fbbedaedab81e2f2e2e9"></a>

Type: `["object", {}]`. Optional, Computed.

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

Terraform syntax:

```terraform
logs_streaming_disabled = {}
```

<a id="canonical-b47794f712007113d19af089f9f3040ca98c4023526654091b5b3a76268c8100"></a>

## Direct properties — logs_streaming_disabled / 2b933d374f9e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d8a128a757b5145f4916e25cde8164db1dd0528b34d49c31288960c286a25b65"></a>

## Next pages — logs_streaming_disabled / 2b933d374f9e / 4

- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-12091dabc447e067174432d59d031f10b3730bf83bc590517c664f1ebc67620e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ec2e4f54532c18f9d4674eec92804552478562b11f57a5e3b90692f0193d24f"></a>

## no_worker_nodes — no_worker_nodes / f167f71519e0 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- no_worker_nodes

<a id="canonical-901f5206c55f10226c0d3d72c42f8444a04c26056043c8fa8e178abf6a44dc0c"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: no\_worker\_nodes, nodes\_per\_az, total\_nodes; Default: no\_worker\_nodes\] Configuration
parameter for no worker nodes. Defaults to \`map\[\]\`. Server applies default when omitted.

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

- [no_worker_nodes](resources--azure_vnet_site--reference--group-008.md#canonical-901f5206c55f10226c0d3d72c42f8444a04c26056043c8fa8e178abf6a44dc0c)
- [nodes_per_az](resources--azure_vnet_site--reference--group-001.md#canonical-f7414ad3e0bf098c5c7e53b85428673559f45d751bb2fc3f22541b66596c634a)
- [total_nodes](resources--azure_vnet_site--reference--group-001.md#canonical-2ee62b44b3416fa22f983c0aa54c0693978eade311abfd14a43cbecfaff46541)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
no_worker_nodes = {}
```

<a id="canonical-4609d4ba7d7a4ac7109d63f09d3865eba019fbefc4da9461e86765682e8a75c0"></a>

## Direct properties — no_worker_nodes / f167f71519e0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-aaabc196703c8f3dc86ffa0b26e0e2f6d3ee6634495be22570990dc332392e64"></a>

## Next pages — no_worker_nodes / f167f71519e0 / 4

- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-6723ac4ac38dc4b1691aa0cf9dabdebad2791a8eb4fac00fded397135cac5838"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c20ab0f53e3dd67993e97663efea27de3e25e452452f6fb663a422c32a2933d4"></a>

## offline_survivability_mode — offline_survivability_mode / 14917c6144ea / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- offline_survivability_mode

<a id="canonical-b5477d9d237e4325d4cd6f1def89b907791578616d8758e44ca29a5790ab8ad1"></a>

Type: `"object"`. single nested block, Optional.

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7..

Upstream description:

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7
days, even when the site is offline. The certificates needed to keep the services running on this
site are signed using a local CA. Secrets would also be cached locally to handle the connectivity
loss. When the mode is toggled, services will restart and traffic disruption will be seen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("enable_offline_survivability_mode",
    "no_offline_survivability_mode")}
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
  "x-ves-oneof-field-offline_survivability_mode_choice": "[\"enable_offline_survivability_mode\",\"no_offline_survivability_mode\"]"
}
```

Terraform syntax:

```terraform
offline_survivability_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-ced2eb0cd99f54d318f30d4936735200adae633b368474fb75bd898710ce85cf"></a>

## Direct properties — offline_survivability_mode / 14917c6144ea / 3

- [enable_offline_survivability_mode](resources--azure_vnet_site--reference--group-008.md#canonical-e5ea75808f242f0e251269bda3967a77ecbc86b7bebb9ed42abca01ea77b7842): complete subsection reference.

- [no_offline_survivability_mode](resources--azure_vnet_site--reference--group-008.md#canonical-af0adacc26170f8a8606d5c9414835610dd9182251acd9a5880c5cd548752821): complete subsection reference.

<a id="canonical-e273258d1ddf70c6c8f681951c035f71aed62c4397abb9ff1ff9375c6be43900"></a>

## Next pages — offline_survivability_mode / 14917c6144ea / 4

- [offline_survivability_mode.enable_offline_survivability_mode](resources--azure_vnet_site--reference--group-008.md#canonical-e5ea75808f242f0e251269bda3967a77ecbc86b7bebb9ed42abca01ea77b7842)
- [offline_survivability_mode.no_offline_survivability_mode](resources--azure_vnet_site--reference--group-008.md#canonical-af0adacc26170f8a8606d5c9414835610dd9182251acd9a5880c5cd548752821)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-e5ea75808f242f0e251269bda3967a77ecbc86b7bebb9ed42abca01ea77b7842"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-73fec1e9173c40235331040b0805c11ff3f6275e61f0f0f36a662739d041f45d"></a>

## offline_survivability_mode.enable_offline_survivability_mode — offline_survivability_mode.enable_offline_survivability_mode / 2d39924dd231 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [offline_survivability_mode](resources--azure_vnet_site--reference--group-008.md#canonical-6723ac4ac38dc4b1691aa0cf9dabdebad2791a8eb4fac00fded397135cac5838)
- offline_survivability_mode.enable_offline_survivability_mode

<a id="canonical-e19fcd524baa6cb71db459238c99829a017dad515de0c0c747d85c3be8a2c819"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable offline survivability mode.

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
enable_offline_survivability_mode = {}
```

<a id="canonical-3d51bab976dfe8898c0f6eb43d83db9ac5c5e95469f508f618610b5addbe055e"></a>

## Direct properties — offline_survivability_mode.enable_offline_survivability_mode / 2d39924dd231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-86fe7173b61d9d47be17343805338aa255e37fc80dc0481bcd1fd7f7f0d3e2db"></a>

## Next pages — offline_survivability_mode.enable_offline_survivability_mode / 2d39924dd231 / 4

- [offline_survivability_mode](resources--azure_vnet_site--reference--group-008.md#canonical-6723ac4ac38dc4b1691aa0cf9dabdebad2791a8eb4fac00fded397135cac5838)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-af0adacc26170f8a8606d5c9414835610dd9182251acd9a5880c5cd548752821"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d28af3b03a50c79e13c2755d27b2b82bceced281f601f706aa938035a07957dd"></a>

## offline_survivability_mode.no_offline_survivability_mode — offline_survivability_mode.no_offline_survivability_mode / 164d3d87e418 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [offline_survivability_mode](resources--azure_vnet_site--reference--group-008.md#canonical-6723ac4ac38dc4b1691aa0cf9dabdebad2791a8eb4fac00fded397135cac5838)
- offline_survivability_mode.no_offline_survivability_mode

<a id="canonical-e59f7deaba71a742bf2a2eb6d36116d4bb16564061bfc713a501b8c0896a29a4"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no offline survivability mode.

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
no_offline_survivability_mode = {}
```

<a id="canonical-c0dfa5ea25703e3ebed1b06bb4cc2f391c184b69df39d48b7c0a7ea350fca56f"></a>

## Direct properties — offline_survivability_mode.no_offline_survivability_mode / 164d3d87e418 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-18fe7fdb129ddcce5e632ce90b5b36d68e816d56d6a4aa33d854272a826ab0c6"></a>

## Next pages — offline_survivability_mode.no_offline_survivability_mode / 164d3d87e418 / 4

- [offline_survivability_mode](resources--azure_vnet_site--reference--group-008.md#canonical-6723ac4ac38dc4b1691aa0cf9dabdebad2791a8eb4fac00fded397135cac5838)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-70dd020b208d2a6edaf4a3aae30f4ff67ffceee27ffc85af22c92dc13e8b70af"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0580abc401719f454bcf459acfc55b5aa2f7108db11707091e2a7a72c9ebbbb4"></a>

## os — os / 3ef8a9d92645 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- os

<a id="canonical-ec07e2c6748f8fc2a471c7d6cc0d504a3545ce5cd48cafa9d49ed6768cf874e4"></a>

Type: `"object"`. single nested block, Optional.

Select the F5XC Operating System Version for the site. By default, latest available OS Version will
be used. Refer to release notes to find required released OS versions.

Upstream description:

Select the F5XC Operating System Version for the site. By default, latest available OS Version will
be used. Refer to release notes to find required released OS versions.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_os_version",
    "operating_system_version")}
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
  "x-ves-oneof-field-operating_system_version_choice": "[\"default_os_version\",\"operating_system_version\"]"
}
```

Terraform syntax:

```terraform
os {
  # Configure direct properties listed below.
}
```

<a id="canonical-19f3de97fe5e10daa80c3524efbe44a259e9d2fb16c5b13b6f40180f5115cc44"></a>

## Direct properties — os / 3ef8a9d92645 / 3

- [default_os_version](resources--azure_vnet_site--reference--group-008.md#canonical-d9a4b1fea46c9de6da3e2132e2db0cb3d0ccf1482556b9d53759a06730ae53ab): complete subsection reference.

<a id="canonical-d98f47ec771868ba7e3268e0bedfca045c6bc51eb75ea84827e01362333df3d0"></a>

<a id="canonical-3b39b82435427af179ce5f0d8da38c1777258c56e5a165d7971e5ef4fb612c7b"></a>

## operating_system_version property — os / 3ef8a9d92645 / 4

Type: `"string"`. Optional.

Exclusive with \[default\_os\_version\] Specify a OS version to be used e.g. 9.2024.6.

Upstream description:

Exclusive with \[default\_os\_version\] Specify a OS version to be used e.g. 9.2024.6.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-4b8fe882ac1571bf3653a77480bd979590a8d5eed6336c8bed6fae9972cb609d"></a>

## Next pages — os / 3ef8a9d92645 / 5

- [os.default_os_version](resources--azure_vnet_site--reference--group-008.md#canonical-d9a4b1fea46c9de6da3e2132e2db0cb3d0ccf1482556b9d53759a06730ae53ab)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-d9a4b1fea46c9de6da3e2132e2db0cb3d0ccf1482556b9d53759a06730ae53ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-25a1959cd2801352247c31dc07999c982f6a1ae8421a688866f4b43cf68a7080"></a>

## os.default_os_version — os.default_os_version / 3e388daca4d6 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [os](resources--azure_vnet_site--reference--group-008.md#canonical-70dd020b208d2a6edaf4a3aae30f4ff67ffceee27ffc85af22c92dc13e8b70af)
- os.default_os_version

<a id="canonical-d06a4088d5ec806b8ff9219198fa36f5448fe66d0d4c51a95c43f0ce527f8c6a"></a>

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
default_os_version = {}
```

<a id="canonical-262e8b4bbf4729d3fd8ada2f3b7b924abde4ae9729a11e3922a829d9159bc0ce"></a>

## Direct properties — os.default_os_version / 3e388daca4d6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4237a0f1f4aceccc1a993c854d3081b65ad05b9d439bf247ef4796a25727909b"></a>

## Next pages — os.default_os_version / 3e388daca4d6 / 4

- [os](resources--azure_vnet_site--reference--group-008.md#canonical-70dd020b208d2a6edaf4a3aae30f4ff67ffceee27ffc85af22c92dc13e8b70af)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-44c595767a0eea838d43115b7d32a7c06c846d485a06c5ffc924da3ceec773d7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2816f2564e3c47a902da485a44f38e0881c88c991f3e4a83fb14b5403d08b338"></a>

## sw — sw / 96acab4f7876 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- sw

<a id="canonical-2893e8b106a8b5416e19c084a9c2168de8f6432be547282536cf8caf8864715c"></a>

Type: `"object"`. single nested block, Optional.

Select the F5XC Software Version for the site. By default, latest available F5XC Software Version
will be used. Refer to release notes to find required released SW versions.

Upstream description:

Select the F5XC Software Version for the site. By default, latest available F5XC Software Version
will be used. Refer to release notes to find required released SW versions.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_sw_version",
    "volterra_software_version")}
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
  "x-ves-oneof-field-volterra_sw_version_choice": "[\"default_sw_version\",\"volterra_software_version\"]"
}
```

Terraform syntax:

```terraform
sw {
  # Configure direct properties listed below.
}
```

<a id="canonical-fc6de07a64c5cd57b8232ececc456e9f2a92f17dd6116b97a2fa1a33bb3ce479"></a>

## Direct properties — sw / 96acab4f7876 / 3

- [default_sw_version](resources--azure_vnet_site--reference--group-008.md#canonical-7a1b3ef296447d4b2fa129266f73f8200fac4b841ad157616974775eadc214a2): complete subsection reference.

<a id="canonical-6b0b64a6b0d535c371d1ac7d9cad41ceef961b8a22e5e468d32b29ea44d2af01"></a>

<a id="canonical-37ee65d61b77491d77582230244f48bc5472a608917372b82fb877e862cb1c16"></a>

## volterra_software_version property — sw / 96acab4f7876 / 4

Type: `"string"`. Optional.

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

Upstream description:

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-961dcdd11fc6e96b214864af88f941048920466450aa93d9bde5ec343c920bf4"></a>

## Next pages — sw / 96acab4f7876 / 5

- [sw.default_sw_version](resources--azure_vnet_site--reference--group-008.md#canonical-7a1b3ef296447d4b2fa129266f73f8200fac4b841ad157616974775eadc214a2)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-7a1b3ef296447d4b2fa129266f73f8200fac4b841ad157616974775eadc214a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0a4d2a5c8a68b965322e6e2c9244d3698648ea14dc306941144236fb6b963407"></a>

## sw.default_sw_version — sw.default_sw_version / 10a392a3fded / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [sw](resources--azure_vnet_site--reference--group-008.md#canonical-44c595767a0eea838d43115b7d32a7c06c846d485a06c5ffc924da3ceec773d7)
- sw.default_sw_version

<a id="canonical-05bfa209d9d57202caa342ba37744a718c6499797e8784fba026127bbfbd0801"></a>

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
default_sw_version = {}
```

<a id="canonical-aac6c7e2c704ef9fd21fbf1adb0613359ae0298db467ba5069253791c2cad192"></a>

## Direct properties — sw.default_sw_version / 10a392a3fded / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5207987749a4c9a0f35c2715ceea02f71b722c518c0c569fba68a445cdea5137"></a>

## Next pages — sw.default_sw_version / 10a392a3fded / 4

- [sw](resources--azure_vnet_site--reference--group-008.md#canonical-44c595767a0eea838d43115b7d32a7c06c846d485a06c5ffc924da3ceec773d7)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-7aaf0ee72da98f9252c44a72ca2e351fbad6c9e97e7f0920851c965169a6da21"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d49c0655a5b4796e0b0642aff22f8e4fa12fdc76166bd1fedab6a2573af3097e"></a>

## timeouts — timeouts / ccc92bec6907 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- timeouts

<a id="canonical-b113bc5a804095a6de310b43deebf31812d44d26f28881f50e88bd87cf2e2cb2"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-cf69f5c5d94ebed59f82565823d94f967d756e661528d29c3784624459f02c9e"></a>

## Direct properties — timeouts / ccc92bec6907 / 3

<a id="canonical-eb9221f3bf75afd0e13039be7ab935241a0e3e65fdc2ccf8728f49f8eca46ef8"></a>

<a id="canonical-e16e14ad3d4a2a1502cb3a66e67166dd46e4e658bfa21e971954e947644df4a5"></a>

## create property — timeouts / ccc92bec6907 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3d9013fb875227685d14541e5aa97aebbfd21e1fd928e489f1db4cfa86a1cfba"></a>

<a id="canonical-3019a2168ff22db24c41dbe820ca2b2f01348f33b9f885395604af932897e8e5"></a>

## delete property — timeouts / ccc92bec6907 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-4b24225d238d0f96ff6f076a3f4b4e46a6adb072f07c909ee5d19e9913a6c71e"></a>

<a id="canonical-d162849fd85e85a236a7b58fcb43e11b8fcda1c360c104dc098c665e0077b1f6"></a>

## read property — timeouts / ccc92bec6907 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-afce0b0e9e3ad3addee7940f82f1e32bbfce3e3c7cd94ca46d1e398fe72f9f70"></a>

<a id="canonical-6056089c55fb164b65c17f1791bd75925e858b5601a8c1a4d1644411ae770172"></a>

## update property — timeouts / ccc92bec6907 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-a8fb30133d529b125d1e2fa19f59e80828e67182ad1846364131df9c7efeb2bb"></a>

## Next pages — timeouts / ccc92bec6907 / 8

- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-8c0e3483e56222f4feb0fd6a8c47c326590ec63eab2337454d8ded6295ce7d93"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3176ddc922087b19392cf173c17982f9c99318b29e4f6d9f51024381551d16e2"></a>

## vnet — vnet / 71b2deb25908 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- vnet

<a id="canonical-438fd3f40e4c60bfc68ab88a38b34430f6a6bb29780893808427a6558787960f"></a>

Type: `"object"`. single nested block, Optional.

Defines choice about Azure VNet for a view.

Upstream description:

This defines choice about Azure VNet for a view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_vnet",
    "new_vnet")}
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
  "x-ves-oneof-field-choice": "[\"existing_vnet\",\"new_vnet\"]"
}
```

Terraform syntax:

```terraform
vnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-94bc6905b44166565af0d206080bb685432a99f3409efcbca695a95116168350"></a>

## Direct properties — vnet / 71b2deb25908 / 3

- [existing_vnet](resources--azure_vnet_site--reference--group-008.md#canonical-4b76e83b6d3d61e39cc10afd26534209bc05d7e311c1d83324386a5434b92f4e): complete subsection reference.

- [new_vnet](resources--azure_vnet_site--reference--group-008.md#canonical-94a563c34282b9d43d0f79e66332fcfd82d81cf45fecf9872e32b7bb5c6f1ae4): complete subsection reference.

<a id="canonical-78ec581f3da4f1b62210507117d91dbdad3ed296351f412fba4e8988f413feb8"></a>

## Next pages — vnet / 71b2deb25908 / 4

- [vnet.existing_vnet](resources--azure_vnet_site--reference--group-008.md#canonical-4b76e83b6d3d61e39cc10afd26534209bc05d7e311c1d83324386a5434b92f4e)
- [vnet.new_vnet](resources--azure_vnet_site--reference--group-008.md#canonical-94a563c34282b9d43d0f79e66332fcfd82d81cf45fecf9872e32b7bb5c6f1ae4)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-4b76e83b6d3d61e39cc10afd26534209bc05d7e311c1d83324386a5434b92f4e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-62c0cdafb655d815ea25270bde2530df3e17e60cde690c7f3d85e7ca4515903e"></a>

## vnet.existing_vnet — vnet.existing_vnet / 8d1439a691b1 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [vnet](resources--azure_vnet_site--reference--group-008.md#canonical-8c0e3483e56222f4feb0fd6a8c47c326590ec63eab2337454d8ded6295ce7d93)
- vnet.existing_vnet

<a id="canonical-827c4fa2ce0d163a8c3583370877ddad1a336e7b8e8d5e8ee693530af4c64708"></a>

Type: `"object"`. single nested block, Optional.

Resource group and name of existing Azure VNet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("resource_group",
    "vnet_name"),
  validators.ConflictingObjectAttributes("f5_orchestrated_routing",
    "manual_routing")}
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
  "x-ves-oneof-field-routing_type": "[\"f5_orchestrated_routing\",\"manual_routing\"]"
}
```

Terraform syntax:

```terraform
existing_vnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-91c29001c446bb365f244801f727d0eca82f40ada57cbe4d28753ae658bb4072"></a>

## Direct properties — vnet.existing_vnet / 8d1439a691b1 / 3

- [f5_orchestrated_routing](resources--azure_vnet_site--reference--group-008.md#canonical-7bd7bf763e13f9c3a5cd6ef35bbb2e728ec3f370500d931d03b0516f50605ca1): complete subsection reference.

- [manual_routing](resources--azure_vnet_site--reference--group-008.md#canonical-28c1680035a532e99781ecb86832f2e7727947ccfba1d5d932fc8f509ce1ff80): complete subsection reference.

<a id="canonical-cfcfa2f47659cc2ec8fe16a5f51f4dac3a47d42d7e5490d032a08f623f1f20b7"></a>

<a id="canonical-34d3d7358b877475f090aa47b611aa3ae2f7d19f27b6e5103133c181a77e97d8"></a>

## resource_group property — vnet.existing_vnet / 8d1439a691b1 / 4

Type: `"string"`. Optional.

Existing VNet Resource Group. Resource group of existing VNet.

Upstream description:

Resource group of existing VNet.

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-05450d32aafd74c375cd7754dfd88afb716ef56de74350b5bbb19e2f1665c021"></a>

<a id="canonical-fe1c0409171d5132189c6abd20c43cf23df7d3d5df744233ac968c05a4cb54b9"></a>

## vnet_name property — vnet.existing_vnet / 8d1439a691b1 / 5

Type: `"string"`. Optional.

Existing VNet Name. Name of existing VNet.

Upstream description:

Name of existing VNet.

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-01c31fd471c0312c0582e102f0e0a776c50d4b7731985d6f85c00900c0d3728b"></a>

## Next pages — vnet.existing_vnet / 8d1439a691b1 / 6

- [vnet.existing_vnet.f5_orchestrated_routing](resources--azure_vnet_site--reference--group-008.md#canonical-7bd7bf763e13f9c3a5cd6ef35bbb2e728ec3f370500d931d03b0516f50605ca1)
- [vnet.existing_vnet.manual_routing](resources--azure_vnet_site--reference--group-008.md#canonical-28c1680035a532e99781ecb86832f2e7727947ccfba1d5d932fc8f509ce1ff80)
- [vnet](resources--azure_vnet_site--reference--group-008.md#canonical-8c0e3483e56222f4feb0fd6a8c47c326590ec63eab2337454d8ded6295ce7d93)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-7bd7bf763e13f9c3a5cd6ef35bbb2e728ec3f370500d931d03b0516f50605ca1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f36da7b7dd48a44e14d314697eb570c67fb3f31c9c4aaa1f94610a7f39b4ee60"></a>

## vnet.existing_vnet.f5_orchestrated_routing — vnet.existing_vnet.f5_orchestrated_routing / 3b2e2787bb32 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [vnet](resources--azure_vnet_site--reference--group-008.md#canonical-8c0e3483e56222f4feb0fd6a8c47c326590ec63eab2337454d8ded6295ce7d93)
- [vnet.existing_vnet](resources--azure_vnet_site--reference--group-008.md#canonical-4b76e83b6d3d61e39cc10afd26534209bc05d7e311c1d83324386a5434b92f4e)
- vnet.existing_vnet.f5_orchestrated_routing

<a id="canonical-8d183ffdf0f774e95c3c093550e918d045c1186260bee03e0fd6edc93446b293"></a>

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
f5_orchestrated_routing = {}
```

<a id="canonical-f05a5f16c5750b16ede27ed61eca36da65ba71aa4bf0277c7358d12f673ff092"></a>

## Direct properties — vnet.existing_vnet.f5_orchestrated_routing / 3b2e2787bb32 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0b1aad949f9d81666ba73f72359d4afd32787a0ebb6f90cca19dcf450a6e97b1"></a>

## Next pages — vnet.existing_vnet.f5_orchestrated_routing / 3b2e2787bb32 / 4

- [vnet.existing_vnet](resources--azure_vnet_site--reference--group-008.md#canonical-4b76e83b6d3d61e39cc10afd26534209bc05d7e311c1d83324386a5434b92f4e)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-28c1680035a532e99781ecb86832f2e7727947ccfba1d5d932fc8f509ce1ff80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7f6cf3b5932372af80ad09a991ad32cd9daea82f5fb6fe6d147f11907f7bee35"></a>

## vnet.existing_vnet.manual_routing — vnet.existing_vnet.manual_routing / 11a6d02c583f / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [vnet](resources--azure_vnet_site--reference--group-008.md#canonical-8c0e3483e56222f4feb0fd6a8c47c326590ec63eab2337454d8ded6295ce7d93)
- [vnet.existing_vnet](resources--azure_vnet_site--reference--group-008.md#canonical-4b76e83b6d3d61e39cc10afd26534209bc05d7e311c1d83324386a5434b92f4e)
- vnet.existing_vnet.manual_routing

<a id="canonical-85ded0f765fb98124ca2b6f8c669440f72af4e75bdc0327346cb4c3e04dc4556"></a>

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
manual_routing = {}
```

<a id="canonical-501d478148277662538c5788eeb6598f73aba0ea7547f900bb9f20e6f9c12bc2"></a>

## Direct properties — vnet.existing_vnet.manual_routing / 11a6d02c583f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c76bcdc7cbe1bc33d1841db57363e0d66cbb5a0ff02109f5e3a629d7274c6acf"></a>

## Next pages — vnet.existing_vnet.manual_routing / 11a6d02c583f / 4

- [vnet.existing_vnet](resources--azure_vnet_site--reference--group-008.md#canonical-4b76e83b6d3d61e39cc10afd26534209bc05d7e311c1d83324386a5434b92f4e)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-94a563c34282b9d43d0f79e66332fcfd82d81cf45fecf9872e32b7bb5c6f1ae4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2257b3999763a89071a5c7fc68fb536980bd68849ae52e5d5e30e1e2d03a8f3b"></a>

## vnet.new_vnet — vnet.new_vnet / 345e52b1e2bb / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [vnet](resources--azure_vnet_site--reference--group-008.md#canonical-8c0e3483e56222f4feb0fd6a8c47c326590ec63eab2337454d8ded6295ce7d93)
- vnet.new_vnet

<a id="canonical-1da9dca7b97de66ac28a6718a3e14c77e6035a379148c2523cd28eb6c789ec77"></a>

Type: `"object"`. single nested block, Optional.

Azure VNet Parameters. Parameters to create a new Azure VNet.

Upstream description:

Parameters to create a new Azure VNet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("primary_ipv4"),
  validators.ConflictingObjectAttributes("autogenerate",
    "name")}
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
  "x-ves-oneof-field-name_choice": "[\"autogenerate\",\"name\"]"
}
```

Terraform syntax:

```terraform
new_vnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-404dd37b64797d754e3cd29253206b4270e7183c7a0226212561004f3785c901"></a>

## Direct properties — vnet.new_vnet / 345e52b1e2bb / 3

- [autogenerate](resources--azure_vnet_site--reference--group-008.md#canonical-2f71e56cd3e8494bb88fc1f8ca07fa6047c4686e23ae90861ec5a102c4f2087e): complete subsection reference.

<a id="canonical-f7dc61aa380843b04d0be430d743b12aea41b75b6c18b09622407c2bc55ffe53"></a>

<a id="canonical-f46f890db5fbbfcfc2df7f094678e54370be1d0932c907b8bcce306d9f63c31a"></a>

## name property — vnet.new_vnet / 345e52b1e2bb / 4

Type: `"string"`. Optional.

Exclusive with \[autogenerate\] Specify the VNet Name.

Upstream description:

Exclusive with \[autogenerate\] Specify the VNet Name.

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
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-6463a2ac7db91c13d9270515ede2fe90a0a030871353d633d6d464951b1a973f"></a>

<a id="canonical-a995d13639efab783b60e797ed26921c74fbd41b32dd44a7e1dcaecd6e84781d"></a>

## primary_ipv4 property — vnet.new_vnet / 345e52b1e2bb / 5

Type: `"string"`. Optional.

IPv4 CIDR block for this VNet. It has to be private address space.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28",
    "ves.io.schema.rules.string.min_ip_prefix_length": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28",
    "ves.io.schema.rules.string.min_ip_prefix_length": "8"
  }
}
```

<a id="canonical-fabceee01afde91875afcc2910436942a0aa72bfd9fc24e1f7c5a0ee4b3e2ab7"></a>

## Next pages — vnet.new_vnet / 345e52b1e2bb / 6

- [vnet.new_vnet.autogenerate](resources--azure_vnet_site--reference--group-008.md#canonical-2f71e56cd3e8494bb88fc1f8ca07fa6047c4686e23ae90861ec5a102c4f2087e)
- [vnet](resources--azure_vnet_site--reference--group-008.md#canonical-8c0e3483e56222f4feb0fd6a8c47c326590ec63eab2337454d8ded6295ce7d93)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-2f71e56cd3e8494bb88fc1f8ca07fa6047c4686e23ae90861ec5a102c4f2087e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-551eb9b45d11bab0518e320ff4e02ecc556d9304638fe7030b3d1a8af0114e8e"></a>

## vnet.new_vnet.autogenerate — vnet.new_vnet.autogenerate / 397644797543 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [vnet](resources--azure_vnet_site--reference--group-008.md#canonical-8c0e3483e56222f4feb0fd6a8c47c326590ec63eab2337454d8ded6295ce7d93)
- [vnet.new_vnet](resources--azure_vnet_site--reference--group-008.md#canonical-94a563c34282b9d43d0f79e66332fcfd82d81cf45fecf9872e32b7bb5c6f1ae4)
- vnet.new_vnet.autogenerate

<a id="canonical-310690aba3efac87e537035f6ef1f89803592ab3cbc0c3d8ae6930a3c3ea6589"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for autogenerate.

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
autogenerate = {}
```

<a id="canonical-bf7749a3f9f5594724710aba6e2bb4692d700f56e67d398ec3f40dc0a8320b91"></a>

## Direct properties — vnet.new_vnet.autogenerate / 397644797543 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-53186b39fa62aaaaf0880bacc33c4c90e81222946a08944c94189d7601d40003"></a>

## Next pages — vnet.new_vnet.autogenerate / 397644797543 / 4

- [vnet.new_vnet](resources--azure_vnet_site--reference--group-008.md#canonical-94a563c34282b9d43d0f79e66332fcfd82d81cf45fecf9872e32b7bb5c6f1ae4)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4c64eff495ef9c9db76c102f798ac9f1fe7d20023825e5df9efb235a734e6b5"></a>

## voltstack_cluster — voltstack_cluster / 2b44c038cb45 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- voltstack_cluster

<a id="canonical-5e0e75d724420192b7836a38457cc26788a32f8d6a90b5db1a9ef84ebe09a914"></a>

Type: `"object"`. single nested block, Optional.

App Stack Cluster of single interface Azure nodes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("az_nodes",
    "azure_certified_hw"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "active_network_policies"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "forward_proxy_allow_all"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("active_network_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("dc_cluster_group",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("default_storage",
    "storage_class_list"),
  validators.ConflictingObjectAttributes("forward_proxy_allow_all",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("global_network_list",
    "no_global_network"),
  validators.ConflictingObjectAttributes("k8s_cluster",
    "no_k8s_cluster"),
  validators.ConflictingObjectAttributes("no_outside_static_routes",
    "outside_static_routes"),
  validators.ConflictingObjectAttributes("sm_connection_public_ip",
    "sm_connection_pvt_ip")}
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
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-k8s_cluster_choice": "[\"k8s_cluster\",\"no_k8s_cluster\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]",
  "x-ves-oneof-field-storage_class_choice": "[\"default_storage\",\"storage_class_list\"]"
}
```

Terraform syntax:

```terraform
voltstack_cluster {
  # Configure direct properties listed below.
}
```

<a id="canonical-6925f7e1a50155cb672222bc76f1a857bfb04375bde9e020ad731b93320ff314"></a>

## Direct properties — voltstack_cluster / 2b44c038cb45 / 3

- [accelerated_networking](resources--azure_vnet_site--reference--group-008.md#canonical-580b6e978a67fae17eebc01cf524e24c28a251ef267600b15fb016be469427ba): complete subsection reference.

- [active_enhanced_firewall_policies](resources--azure_vnet_site--reference--group-008.md#canonical-c7daa2b1eead26913cf943543104ade37afecb9636f69eb95b6b8b3b8d3ff556): complete subsection reference.

- [active_forward_proxy_policies](resources--azure_vnet_site--reference--group-008.md#canonical-33ec78eec3070b27dcac18ed362ee719d939506c91581ffdf5fe357b73351a31): complete subsection reference.

- [active_network_policies](resources--azure_vnet_site--reference--group-008.md#canonical-5f0ab925481ff546dbce389f846a2a1e83e98011584758d4497d6150c48cd228): complete subsection reference.

- [az_nodes](resources--azure_vnet_site--reference--group-008.md#canonical-543235bd47f7e2a85d5558a07785780e10c0b9a46c89b4babad14a7dbac537c9): complete subsection reference.

<a id="canonical-ad2371c20f2af1824179100ed50ebe6286a18cbf85271b890c2a416625e2a7f8"></a>

<a id="canonical-c698e3f16d605b95973215a5cfe7c6e5074bcd5879834272baaa140cb2c02e11"></a>

## azure_certified_hw property — voltstack_cluster / 2b44c038cb45 / 4

Type: `"string"`. Optional.

\[Enum: azure-byol-voltstack-combo\] Azure Certified Hardware. Name for Azure certified hardware.
The only possible value is \`azure-byol-voltstack-combo\`.

Upstream description:

Name for Azure certified hardware.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
  stringvalidator.OneOf("azure-byol-voltstack-combo"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "azure-byol-voltstack-combo"
  ],
  "maxLength": 64,
  "x-f5xc-constraints": {
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [dc_cluster_group](resources--azure_vnet_site--reference--group-008.md#canonical-9d2d12a1b34075b677dd04161498868b24e0857164f1b741844d65f8cccd3d0a): complete subsection reference.

- [default_storage](resources--azure_vnet_site--reference--group-008.md#canonical-ed31b36fb96ae3be15150a7719135aebdb54e2fdb1a7466cbe4c60a030e14425): complete subsection reference.

- [forward_proxy_allow_all](resources--azure_vnet_site--reference--group-008.md#canonical-c15dea8b1c56cbbe431e0b31debcee5d177ff009eac2f2f3d62ac6ee36c4344f): complete subsection reference.

- [global_network_list](resources--azure_vnet_site--reference--group-008.md#canonical-a8935381bd30aecbecc6a8cd035193d1f167660b2bf568f5541b0a720dd3966f): complete subsection reference.

- [k8s_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d30e0699286fdea85a05f142127c4e9412c34f7bfc1257205aae6153dfa5c193): complete subsection reference.

- [no_dc_cluster_group](resources--azure_vnet_site--reference--group-009.md#canonical-289e5681a9cf2e53242dcc84d84f97ca791fb0c3383962a8ba1d9a73adaba898): complete subsection reference.

- [no_forward_proxy](resources--azure_vnet_site--reference--group-009.md#canonical-da0a9f761b85bb12907956b73c9de07c013d2a221aed8b38eb299e527f86258f): complete subsection reference.

- [no_global_network](resources--azure_vnet_site--reference--group-009.md#canonical-a832f2f4936449e6da6a29d31b957acde0e04f79c09f32485b52ecabc67ab28c): complete subsection reference.

- [no_k8s_cluster](resources--azure_vnet_site--reference--group-009.md#canonical-eebc49664723328c5667c67d2c0b2ef51f3daaf7ae64eb1fd7f9bc93b7bdc76d): complete subsection reference.

- [no_network_policy](resources--azure_vnet_site--reference--group-009.md#canonical-07114fe9e1bfc024bf75e993864a9f4f594519b6e0a53400b7b83ec05ea2ce19): complete subsection reference.

- [no_outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-458d512fac829eb4805efac48006a46ceafc83a4a8805838e2191731127078a4): complete subsection reference.

- [outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-10b7b5c5ca2525688944c128802937ae34e397f8894fec2dde6dd487195ef2aa): complete subsection reference.

- [sm_connection_public_ip](resources--azure_vnet_site--reference--group-009.md#canonical-b97e72482667a76994786d259dab5022563f9130822b995a8826a686419910d5): complete subsection reference.

- [sm_connection_pvt_ip](resources--azure_vnet_site--reference--group-009.md#canonical-c7d7c49bef3b6826e08eb566bee76348916abaaa4a828a15db5482e7b5ef09e9): complete subsection reference.

- [storage_class_list](resources--azure_vnet_site--reference--group-009.md#canonical-f5eb60002d9390fbb074b58b83119e443d059edb53f00c5b35b3f02b327b480c): complete subsection reference.

<a id="canonical-510be12d12c044a36ffe9edc3facf031f75c82edb631d8a4cb7374a33718b982"></a>

## Next pages — voltstack_cluster / 2b44c038cb45 / 5

- [voltstack_cluster.accelerated_networking](resources--azure_vnet_site--reference--group-008.md#canonical-580b6e978a67fae17eebc01cf524e24c28a251ef267600b15fb016be469427ba)
- [voltstack_cluster.active_enhanced_firewall_policies](resources--azure_vnet_site--reference--group-008.md#canonical-c7daa2b1eead26913cf943543104ade37afecb9636f69eb95b6b8b3b8d3ff556)
- [voltstack_cluster.active_forward_proxy_policies](resources--azure_vnet_site--reference--group-008.md#canonical-33ec78eec3070b27dcac18ed362ee719d939506c91581ffdf5fe357b73351a31)
- [voltstack_cluster.active_network_policies](resources--azure_vnet_site--reference--group-008.md#canonical-5f0ab925481ff546dbce389f846a2a1e83e98011584758d4497d6150c48cd228)
- [voltstack_cluster.az_nodes](resources--azure_vnet_site--reference--group-008.md#canonical-543235bd47f7e2a85d5558a07785780e10c0b9a46c89b4babad14a7dbac537c9)
- [voltstack_cluster.dc_cluster_group](resources--azure_vnet_site--reference--group-008.md#canonical-9d2d12a1b34075b677dd04161498868b24e0857164f1b741844d65f8cccd3d0a)
- [voltstack_cluster.default_storage](resources--azure_vnet_site--reference--group-008.md#canonical-ed31b36fb96ae3be15150a7719135aebdb54e2fdb1a7466cbe4c60a030e14425)
- [voltstack_cluster.forward_proxy_allow_all](resources--azure_vnet_site--reference--group-008.md#canonical-c15dea8b1c56cbbe431e0b31debcee5d177ff009eac2f2f3d62ac6ee36c4344f)
- [voltstack_cluster.global_network_list](resources--azure_vnet_site--reference--group-008.md#canonical-a8935381bd30aecbecc6a8cd035193d1f167660b2bf568f5541b0a720dd3966f)
- [voltstack_cluster.k8s_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d30e0699286fdea85a05f142127c4e9412c34f7bfc1257205aae6153dfa5c193)
- [voltstack_cluster.no_dc_cluster_group](resources--azure_vnet_site--reference--group-009.md#canonical-289e5681a9cf2e53242dcc84d84f97ca791fb0c3383962a8ba1d9a73adaba898)
- [voltstack_cluster.no_forward_proxy](resources--azure_vnet_site--reference--group-009.md#canonical-da0a9f761b85bb12907956b73c9de07c013d2a221aed8b38eb299e527f86258f)
- [voltstack_cluster.no_global_network](resources--azure_vnet_site--reference--group-009.md#canonical-a832f2f4936449e6da6a29d31b957acde0e04f79c09f32485b52ecabc67ab28c)
- [voltstack_cluster.no_k8s_cluster](resources--azure_vnet_site--reference--group-009.md#canonical-eebc49664723328c5667c67d2c0b2ef51f3daaf7ae64eb1fd7f9bc93b7bdc76d)
- [voltstack_cluster.no_network_policy](resources--azure_vnet_site--reference--group-009.md#canonical-07114fe9e1bfc024bf75e993864a9f4f594519b6e0a53400b7b83ec05ea2ce19)
- [voltstack_cluster.no_outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-458d512fac829eb4805efac48006a46ceafc83a4a8805838e2191731127078a4)
- [voltstack_cluster.outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-10b7b5c5ca2525688944c128802937ae34e397f8894fec2dde6dd487195ef2aa)
- [voltstack_cluster.sm_connection_public_ip](resources--azure_vnet_site--reference--group-009.md#canonical-b97e72482667a76994786d259dab5022563f9130822b995a8826a686419910d5)
- [voltstack_cluster.sm_connection_pvt_ip](resources--azure_vnet_site--reference--group-009.md#canonical-c7d7c49bef3b6826e08eb566bee76348916abaaa4a828a15db5482e7b5ef09e9)
- [voltstack_cluster.storage_class_list](resources--azure_vnet_site--reference--group-009.md#canonical-f5eb60002d9390fbb074b58b83119e443d059edb53f00c5b35b3f02b327b480c)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-580b6e978a67fae17eebc01cf524e24c28a251ef267600b15fb016be469427ba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-71a94560f8b957b5877adb30208d3cb2e1be98e511f73207106c24d244763dfb"></a>

## voltstack_cluster.accelerated_networking — voltstack_cluster.accelerated_networking / 4c0aec71627d / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- voltstack_cluster.accelerated_networking

<a id="canonical-c0912438c4918c11ed247f0c9c8bc17a8c8e62d91b3ef4e0abcaf6623cdfa4e7"></a>

Type: `"object"`. single nested block, Optional.

Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.

Upstream description:

Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "enable")}
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
  "x-ves-oneof-field-accelerated_networking": "[\"disable\",\"enable\"]"
}
```

Terraform syntax:

```terraform
accelerated_networking {
  # Configure direct properties listed below.
}
```

<a id="canonical-23eb8af794af5f8c246d1ad15f8cb92e72263df993e5db744f5b1ea63cf65a0e"></a>

## Direct properties — voltstack_cluster.accelerated_networking / 4c0aec71627d / 3

- [disable_spec](resources--azure_vnet_site--reference--group-008.md#canonical-e26929193933411a82c1cfcc7c093ce35e28cae79dfe298fffa3a7524a9a785b): complete subsection reference.

- [enable](resources--azure_vnet_site--reference--group-008.md#canonical-a151b97071eda85972f2b9fe47c8590dcc3924eba66db356f92b0827afe17e40): complete subsection reference.

<a id="canonical-553a52b9f3fcbf0c0a76eb6e2822314b03c19ab97cf29b5a137dbeb1ac12778e"></a>

## Next pages — voltstack_cluster.accelerated_networking / 4c0aec71627d / 4

- [voltstack_cluster.accelerated_networking.disable_spec](resources--azure_vnet_site--reference--group-008.md#canonical-e26929193933411a82c1cfcc7c093ce35e28cae79dfe298fffa3a7524a9a785b)
- [voltstack_cluster.accelerated_networking.enable](resources--azure_vnet_site--reference--group-008.md#canonical-a151b97071eda85972f2b9fe47c8590dcc3924eba66db356f92b0827afe17e40)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-e26929193933411a82c1cfcc7c093ce35e28cae79dfe298fffa3a7524a9a785b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0a53f456e64ded16a5f9e1b1096bc3075e2ca54c068f007e9f131bd38220e669"></a>

## voltstack_cluster.accelerated_networking.disable_spec — voltstack_cluster.accelerated_networking.disable_spec / 10bfa456f947 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [voltstack_cluster.accelerated_networking](resources--azure_vnet_site--reference--group-008.md#canonical-580b6e978a67fae17eebc01cf524e24c28a251ef267600b15fb016be469427ba)
- voltstack_cluster.accelerated_networking.disable_spec

<a id="canonical-9dad896104f6ddcc7fd6e5dd36b20e5d0c02e7c5722f100194a92357575946b3"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

<a id="canonical-f104c826c655ca1ce9e037f9f46b43d7fa8a2697f6af854f7ea11b063fbbac65"></a>

## Direct properties — voltstack_cluster.accelerated_networking.disable_spec / 10bfa456f947 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6046f0a1c356795ca119d6960340f9db8abc4ef9de2ef148767deb255119c403"></a>

## Next pages — voltstack_cluster.accelerated_networking.disable_spec / 10bfa456f947 / 4

- [voltstack_cluster.accelerated_networking](resources--azure_vnet_site--reference--group-008.md#canonical-580b6e978a67fae17eebc01cf524e24c28a251ef267600b15fb016be469427ba)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-a151b97071eda85972f2b9fe47c8590dcc3924eba66db356f92b0827afe17e40"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-52a327788c67e9f745fb840973765234da59d5556900c5b0fd77627902ae4322"></a>

## voltstack_cluster.accelerated_networking.enable — voltstack_cluster.accelerated_networking.enable / 8e377b717aa5 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [voltstack_cluster.accelerated_networking](resources--azure_vnet_site--reference--group-008.md#canonical-580b6e978a67fae17eebc01cf524e24c28a251ef267600b15fb016be469427ba)
- voltstack_cluster.accelerated_networking.enable

<a id="canonical-d6550ba736994701544c5c4924302758464eda3e86c0d88c3732dd8052c6b9e1"></a>

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
enable = {}
```

<a id="canonical-ad32656525dbd49f5cefed89c260fb7288aa0e9733119505cb5f39a62cadf81a"></a>

## Direct properties — voltstack_cluster.accelerated_networking.enable / 8e377b717aa5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-db0aa733cd699fbdacc1d964c63e7c132cd1fef1b1e655f37134ffbf31307ba0"></a>

## Next pages — voltstack_cluster.accelerated_networking.enable / 8e377b717aa5 / 4

- [voltstack_cluster.accelerated_networking](resources--azure_vnet_site--reference--group-008.md#canonical-580b6e978a67fae17eebc01cf524e24c28a251ef267600b15fb016be469427ba)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-c7daa2b1eead26913cf943543104ade37afecb9636f69eb95b6b8b3b8d3ff556"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f418e63c809d25a7a4f50f4edc1e3a948babb55384e9b66ccd71c9517b5f1b65"></a>

## voltstack_cluster.active_enhanced_firewall_policies — voltstack_cluster.active_enhanced_firewall_policies / 95699ddd41db / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- voltstack_cluster.active_enhanced_firewall_policies

<a id="canonical-dfce3cc7050baa483d055eb8c0f2b4a2fe03d5e6a2114661c97c42e3f7c2f505"></a>

Type: `"object"`. single nested block, Optional.

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Upstream description:

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("enhanced_firewall_policies")}
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
active_enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-66be392321fc5ab9b528607c351cd202dc7f60fd1ca687482ea16c46217bf8c4"></a>

## Direct properties — voltstack_cluster.active_enhanced_firewall_policies / 95699ddd41db / 3

- [enhanced_firewall_policies](resources--azure_vnet_site--reference--group-008.md#canonical-6eceb5065a035f07af041a812a763cb480b9d7acfc65396cb0c2894889b3fc0b): complete subsection reference.

<a id="canonical-6359a11b2982fd53b83dfa892db3974ba092090f1b1b449a90effe1c018fec49"></a>

## Next pages — voltstack_cluster.active_enhanced_firewall_policies / 95699ddd41db / 4

- [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--azure_vnet_site--reference--group-008.md#canonical-6eceb5065a035f07af041a812a763cb480b9d7acfc65396cb0c2894889b3fc0b)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-6eceb5065a035f07af041a812a763cb480b9d7acfc65396cb0c2894889b3fc0b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0d0941939c7613c46827bcdd106c483cffaf60769309f571af434869edb56354"></a>

## voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies — voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies / bfbfba720b4b / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [voltstack_cluster.active_enhanced_firewall_policies](resources--azure_vnet_site--reference--group-008.md#canonical-c7daa2b1eead26913cf943543104ade37afecb9636f69eb95b6b8b3b8d3ff556)
- voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-f9245297e7c79c3227c65e207b2e31d2951741b4f39664754ed3f9e23916df5c"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Enhanced Firewall Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-6c1b6c116ea21fd7915430fb513b50e84314fe7c8cb38cde68158a28349e63fd"></a>

## Direct properties — voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies / bfbfba720b4b / 3

<a id="canonical-62e4d941604c3932703dae9962191554d72b38ebbbf2cb8e08bb859b917913cf"></a>

<a id="canonical-8c482ad7e95db3517f60f1c4d9ceb81218f1bb10bc56caf2c4bcbb105d80a2b1"></a>

## name property — voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies / bfbfba720b4b / 4

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

<a id="canonical-669fc373d2c1c3ae3911a344bf65566e552ba3745a164cf47623287f4c4d0fe0"></a>

<a id="canonical-291178121a7721f177e78b2e1dda211eef3edc14736b9b0a78cec5d4d1e31925"></a>

## namespace property — voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies / bfbfba720b4b / 5

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

<a id="canonical-24b9f8a54949d96107a7da8542c096360b45fb29b2af87024ab39875b9ec35dd"></a>

<a id="canonical-931d2981171602cbd943959dc6a0f59015a498c23d2a87a130ca20564c27f6d1"></a>

## tenant property — voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies / bfbfba720b4b / 6

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

<a id="canonical-c06f7bfb8f424ac6bc4ea6d6928bb280a6352330d2e3c4f82c528e34cc659040"></a>

## Next pages — voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies / bfbfba720b4b / 7

- [voltstack_cluster.active_enhanced_firewall_policies](resources--azure_vnet_site--reference--group-008.md#canonical-c7daa2b1eead26913cf943543104ade37afecb9636f69eb95b6b8b3b8d3ff556)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-33ec78eec3070b27dcac18ed362ee719d939506c91581ffdf5fe357b73351a31"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-62269fd268ae4388ff5d2417eea1704af72f391936450250150d8ddc9f378673"></a>

## voltstack_cluster.active_forward_proxy_policies — voltstack_cluster.active_forward_proxy_policies / d92cdc7efbed / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- voltstack_cluster.active_forward_proxy_policies

<a id="canonical-0a741f1e906a9070df6f630558c84973004dadffc55a2daef40d68093741b9c7"></a>

Type: `"object"`. single nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("forward_proxy_policies")}
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
active_forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-1c065d2a816c713c1034db7af2b1c16326c7148c803291e769a5f00c930a7959"></a>

## Direct properties — voltstack_cluster.active_forward_proxy_policies / d92cdc7efbed / 3

- [forward_proxy_policies](resources--azure_vnet_site--reference--group-008.md#canonical-d91da465e65e35409b8e7d441641a923fb41fb6a948cf457b8837609358f2393): complete subsection reference.

<a id="canonical-cc8196b81de54ec95f2a288c98f5c2616e184d297e1b0e28b2db236e7baf2d76"></a>

## Next pages — voltstack_cluster.active_forward_proxy_policies / d92cdc7efbed / 4

- [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies](resources--azure_vnet_site--reference--group-008.md#canonical-d91da465e65e35409b8e7d441641a923fb41fb6a948cf457b8837609358f2393)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-d91da465e65e35409b8e7d441641a923fb41fb6a948cf457b8837609358f2393"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7eb6545c1697b9a1afc14ee88682e737650446bd31b57855fb00e2927d933571"></a>

## voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies — voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies / a2ba65271e60 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [voltstack_cluster.active_forward_proxy_policies](resources--azure_vnet_site--reference--group-008.md#canonical-33ec78eec3070b27dcac18ed362ee719d939506c91581ffdf5fe357b73351a31)
- voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-0a932eebc1f05f1ceb4ff9bc00db6de65845148db4ec927505ba1126f9f78a1e"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-6c379a9c01c1796400526a30c1e9984e3cf0a211a8d9472b2e6fda95299327f0"></a>

## Direct properties — voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies / a2ba65271e60 / 3

<a id="canonical-3a563b978e248af6d27e6153a54bed87ae82987c47585cf6a065443252621b8d"></a>

<a id="canonical-86ff3b1e7107fbeb35d23727f53718ddcd5f7926b2ad8c814d0c94feac4c6fd5"></a>

## name property — voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies / a2ba65271e60 / 4

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

<a id="canonical-7de9826d9e6ea87149c6178b8f88918d687992ae278e41da62f6a6474c1bf466"></a>

<a id="canonical-29acae4feecb9ead7c9796afb6e57f95354b8338bc5e2a1a2fc9eb15a3d4a1a1"></a>

## namespace property — voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies / a2ba65271e60 / 5

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

<a id="canonical-149e60682c7830947c02b29b34b158acff4c54957f187fe0d8b9dd2eeb0b2714"></a>

<a id="canonical-de08454d52b372c51b63f44e95b379aecaca1bd12069ed7527c4c7d57dd150bf"></a>

## tenant property — voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies / a2ba65271e60 / 6

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

<a id="canonical-f625e6727dcac6da52d82d39c706ac12e02b3e1fe054255963763938e083a5b9"></a>

## Next pages — voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies / a2ba65271e60 / 7

- [voltstack_cluster.active_forward_proxy_policies](resources--azure_vnet_site--reference--group-008.md#canonical-33ec78eec3070b27dcac18ed362ee719d939506c91581ffdf5fe357b73351a31)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-5f0ab925481ff546dbce389f846a2a1e83e98011584758d4497d6150c48cd228"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b5b8ba4f0e9076ab3285b904bdafd5298bbeb1410ab88513a243f278571cbab3"></a>

## voltstack_cluster.active_network_policies — voltstack_cluster.active_network_policies / 6cd5c4404c61 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- voltstack_cluster.active_network_policies

<a id="canonical-8631e03cbd9c4bc9245efe87215c45602f4a848d6e95ba0b1867e443d8526d31"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("network_policies")}
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
active_network_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-2d5348b362e82746bf17f17932a251765219a16ffc003dcec2cef195459f2776"></a>

## Direct properties — voltstack_cluster.active_network_policies / 6cd5c4404c61 / 3

- [network_policies](resources--azure_vnet_site--reference--group-008.md#canonical-2c8f8e3fb0fe4142df1f19457030ae8fb3add7d5bb4a574c97cadc2a00694a77): complete subsection reference.

<a id="canonical-eca8c84fb161568e2ce81ec85d94fadbd587bb79eb32034cbafa0ddc62c6d35c"></a>

## Next pages — voltstack_cluster.active_network_policies / 6cd5c4404c61 / 4

- [voltstack_cluster.active_network_policies.network_policies](resources--azure_vnet_site--reference--group-008.md#canonical-2c8f8e3fb0fe4142df1f19457030ae8fb3add7d5bb4a574c97cadc2a00694a77)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-2c8f8e3fb0fe4142df1f19457030ae8fb3add7d5bb4a574c97cadc2a00694a77"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-00f3149c6a8d207c330ff0f1ff4090a35de0781d92126c9eca660fa8dc5b0074"></a>

## voltstack_cluster.active_network_policies.network_policies — voltstack_cluster.active_network_policies.network_policies / fa469f9827b6 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [voltstack_cluster.active_network_policies](resources--azure_vnet_site--reference--group-008.md#canonical-5f0ab925481ff546dbce389f846a2a1e83e98011584758d4497d6150c48cd228)
- voltstack_cluster.active_network_policies.network_policies

<a id="canonical-6dd886db0a41e1fff068759a28bc302c509d542397079d884e940bbfbe438678"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Firewall Policies active for this network firewall.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
network_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-6c7a10c34fe2160b5f3756fc24d86154d62b061449a630859ba2b435bdc55981"></a>

## Direct properties — voltstack_cluster.active_network_policies.network_policies / fa469f9827b6 / 3

<a id="canonical-a8e2f7f0283d52b01a3455e3ce6036c7af2b37a46577d07b46cf17cec29958d7"></a>

<a id="canonical-22cf8eba56fc5b22b7c2dfde1e9ad9d581e97cc2685d25ddd3d18407410b2769"></a>

## name property — voltstack_cluster.active_network_policies.network_policies / fa469f9827b6 / 4

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

<a id="canonical-ec9b80330ef86161b6ca596c20560ca32ea9d50a0a2b3ff11b23a2c44ed3b130"></a>

<a id="canonical-45ca461148ecb1b81a0a8f101fbc14ee09c52f93041090dd0183e00b4793b321"></a>

## namespace property — voltstack_cluster.active_network_policies.network_policies / fa469f9827b6 / 5

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

<a id="canonical-e6ae6d47b3ed90a6af1ff5fa854142c3dd3a0a479970eb655c0a9460d8bbb706"></a>

<a id="canonical-25b812d3dac29d4de1e9bbfc66e7656ea365fb7c6b2dec1f24da8aa43c52f12e"></a>

## tenant property — voltstack_cluster.active_network_policies.network_policies / fa469f9827b6 / 6

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

<a id="canonical-9c370eb0f6632b32b4b9782d6436dcdda22938b117f8bf4b755492342bff57f4"></a>

## Next pages — voltstack_cluster.active_network_policies.network_policies / fa469f9827b6 / 7

- [voltstack_cluster.active_network_policies](resources--azure_vnet_site--reference--group-008.md#canonical-5f0ab925481ff546dbce389f846a2a1e83e98011584758d4497d6150c48cd228)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-543235bd47f7e2a85d5558a07785780e10c0b9a46c89b4babad14a7dbac537c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc429a2ab2c2cad820a153ba48b1135237bc89edf151f6f4ae9d69a6053a640b"></a>

## voltstack_cluster.az_nodes — voltstack_cluster.az_nodes / 94d7cb342434 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- voltstack_cluster.az_nodes

<a id="canonical-018c0923bf253528b7f5b9a8928319e215905df3266a22820013d4f6cc5722a3"></a>

Type: `"object"`. list nested block, Optional.

Only Single AZ or Three AZ(s) nodes are supported currently.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("azure_az")}
```

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
    "ves.io.schema.rules.repeated.num_items": "1,3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1,3"
  }
}
```

Terraform syntax:

```terraform
az_nodes {
  # Configure direct properties listed below.
}
```

<a id="canonical-6e3bfee521bdb390f21d02332ee240a4ca92e96ef069496223af7a09e04e926f"></a>

## Direct properties — voltstack_cluster.az_nodes / 94d7cb342434 / 3

<a id="canonical-dfc14c513542c2f464d5ec6a7af49de1d5bc8db7f922d38535af338519b9d6d2"></a>

<a id="canonical-d6d1ac9c3adc54db0b6c035ba44d306436715863f652ac371b041e4358279684"></a>

## azure_az property — voltstack_cluster.az_nodes / 94d7cb342434 / 4

Type: `"string"`. Optional.

\[Enum: 1|2|3\] Zone depicting a grouping of datacenters within an Azure region. Expecting numeric
input. Possible values are \`1\`, \`2\`, \`3\`.

Upstream description:

A zone depicting a grouping of datacenters within an Azure region. Expecting numeric input.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("1",
    "2",
    "3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "1",
    "2",
    "3"
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"1\\\",\\\"2\\\",\\\"3\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"1\\\",\\\"2\\\",\\\"3\\\"]"
  }
}
```

- [local_subnet](resources--azure_vnet_site--reference--group-008.md#canonical-5e068dcb356af048b2b5e6233f8e7eff9bc32b8cc5111e1ce380f6b07f44c4c9): complete subsection reference.

<a id="canonical-b2b5966de4e757dc3c1285b3ff798738c761308fe0a8f999eb1e25fb80b4b55c"></a>

## Next pages — voltstack_cluster.az_nodes / 94d7cb342434 / 5

- [voltstack_cluster.az_nodes.local_subnet](resources--azure_vnet_site--reference--group-008.md#canonical-5e068dcb356af048b2b5e6233f8e7eff9bc32b8cc5111e1ce380f6b07f44c4c9)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-5e068dcb356af048b2b5e6233f8e7eff9bc32b8cc5111e1ce380f6b07f44c4c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-474e39817e16ad453e2a260ef0265a041a4e96928e60c477629255115bea634c"></a>

## voltstack_cluster.az_nodes.local_subnet — voltstack_cluster.az_nodes.local_subnet / b412e4e1e07d / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [voltstack_cluster.az_nodes](resources--azure_vnet_site--reference--group-008.md#canonical-543235bd47f7e2a85d5558a07785780e10c0b9a46c89b4babad14a7dbac537c9)
- voltstack_cluster.az_nodes.local_subnet

<a id="canonical-b4912f8a3b7424d5edcd1b8961a218331e038f46b87c054c7368ea5cf52f287f"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for local subnet.

Upstream description:

Parameters for Azure subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("subnet",
    "subnet_param")}
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
  "x-ves-oneof-field-choice": "[\"subnet\",\"subnet_param\"]"
}
```

Terraform syntax:

```terraform
local_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-73da251b196a67bf52e4973118fbfcc4570f6c17140e9840b7511a5e3fba7d21"></a>

## Direct properties — voltstack_cluster.az_nodes.local_subnet / b412e4e1e07d / 3

- [subnet](resources--azure_vnet_site--reference--group-008.md#canonical-b926010db376e38225d7e7ebceca7b5283fad4b6380ced347e24270b937efb08): complete subsection reference.

- [subnet_param](resources--azure_vnet_site--reference--group-008.md#canonical-8ab67dd0204d5e6753dee469e6b1cb4905e52125ce05241f90f4a9d0f3d44f7d): complete subsection reference.

<a id="canonical-45418f7ce935d1f324b2f9e5d80b0a8d7372cc262a1faa65ee77d00de540df9d"></a>

## Next pages — voltstack_cluster.az_nodes.local_subnet / b412e4e1e07d / 4

- [voltstack_cluster.az_nodes.local_subnet.subnet](resources--azure_vnet_site--reference--group-008.md#canonical-b926010db376e38225d7e7ebceca7b5283fad4b6380ced347e24270b937efb08)
- [voltstack_cluster.az_nodes.local_subnet.subnet_param](resources--azure_vnet_site--reference--group-008.md#canonical-8ab67dd0204d5e6753dee469e6b1cb4905e52125ce05241f90f4a9d0f3d44f7d)
- [voltstack_cluster.az_nodes](resources--azure_vnet_site--reference--group-008.md#canonical-543235bd47f7e2a85d5558a07785780e10c0b9a46c89b4babad14a7dbac537c9)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-b926010db376e38225d7e7ebceca7b5283fad4b6380ced347e24270b937efb08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3898da07c41bc405028eaa5ad7f4d83dedcc78cbaa2df30c77efe79813ace7e6"></a>

## voltstack_cluster.az_nodes.local_subnet.subnet — voltstack_cluster.az_nodes.local_subnet.subnet / 526dfcc2199a / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [voltstack_cluster.az_nodes](resources--azure_vnet_site--reference--group-008.md#canonical-543235bd47f7e2a85d5558a07785780e10c0b9a46c89b4babad14a7dbac537c9)
- [voltstack_cluster.az_nodes.local_subnet](resources--azure_vnet_site--reference--group-008.md#canonical-5e068dcb356af048b2b5e6233f8e7eff9bc32b8cc5111e1ce380f6b07f44c4c9)
- voltstack_cluster.az_nodes.local_subnet.subnet

<a id="canonical-478d05722190ea31211ca6753cefe4c6c610f86236cf54eb76b8334aeb6fa355"></a>

Type: `"object"`. single nested block, Optional.

Subnet specification for network segmentation.

Upstream description:

Parameters for Azure subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subnet_name"),
  validators.ConflictingObjectAttributes("subnet_resource_grp",
    "vnet_resource_group")}
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
  "x-ves-oneof-field-resource_group_choice": "[\"subnet_resource_grp\",\"vnet_resource_group\"]"
}
```

Terraform syntax:

```terraform
subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-c2b623edba6261b8b1d4477e16408288100b9fe2f098c275eb372a46808d99d5"></a>

## Direct properties — voltstack_cluster.az_nodes.local_subnet.subnet / 526dfcc2199a / 3

<a id="canonical-a004edc767d2207a4d8ceb6255d0b564a0a0a1155e40e531d47ebb41db2c0827"></a>

<a id="canonical-671c3ded510e35ebd5fd7de23e53e47df00cafa3935705d87d95a731e7e177ff"></a>

## subnet_name property — voltstack_cluster.az_nodes.local_subnet.subnet / 526dfcc2199a / 4

Type: `"string"`. Optional.

Subnet Name. Name of existing subnet.

Upstream description:

Name of existing subnet.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-2a19058e6c943dd69a44bba9fbcd2a9ad1b7c207620572765782b58f862a48aa"></a>

<a id="canonical-aad42adc83cbed61635cc9a54bf2823210a8cba2101b949ce71b203e8145d475"></a>

## subnet_resource_grp property — voltstack_cluster.az_nodes.local_subnet.subnet / 526dfcc2199a / 5

Type: `"string"`. Optional.

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Upstream description:

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [vnet_resource_group](resources--azure_vnet_site--reference--group-008.md#canonical-362af6321c6b88417ef0952c34395d40523cd7d2fa005cf20ef4b37c01eb6ef5): complete subsection reference.

<a id="canonical-2363435f02e382d415ae6eedd3978f721a5a8af3505e34ff1484ef2195a30098"></a>

## Next pages — voltstack_cluster.az_nodes.local_subnet.subnet / 526dfcc2199a / 6

- [voltstack_cluster.az_nodes.local_subnet.subnet.vnet_resource_group](resources--azure_vnet_site--reference--group-008.md#canonical-362af6321c6b88417ef0952c34395d40523cd7d2fa005cf20ef4b37c01eb6ef5)
- [voltstack_cluster.az_nodes.local_subnet](resources--azure_vnet_site--reference--group-008.md#canonical-5e068dcb356af048b2b5e6233f8e7eff9bc32b8cc5111e1ce380f6b07f44c4c9)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-362af6321c6b88417ef0952c34395d40523cd7d2fa005cf20ef4b37c01eb6ef5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4bbc2ac9f32cf8aa718b85895a87015a0f988e9a29c68a2015a6e6ae19ae1d4e"></a>

## voltstack_cluster.az_nodes.local_subnet.subnet.vnet_resource_group — voltstack_cluster.az_nodes.local_subnet.subnet.vnet_resource_group / 0ea08730e8b6 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [voltstack_cluster.az_nodes](resources--azure_vnet_site--reference--group-008.md#canonical-543235bd47f7e2a85d5558a07785780e10c0b9a46c89b4babad14a7dbac537c9)
- [voltstack_cluster.az_nodes.local_subnet](resources--azure_vnet_site--reference--group-008.md#canonical-5e068dcb356af048b2b5e6233f8e7eff9bc32b8cc5111e1ce380f6b07f44c4c9)
- [voltstack_cluster.az_nodes.local_subnet.subnet](resources--azure_vnet_site--reference--group-008.md#canonical-b926010db376e38225d7e7ebceca7b5283fad4b6380ced347e24270b937efb08)
- voltstack_cluster.az_nodes.local_subnet.subnet.vnet_resource_group

<a id="canonical-ef32c77701deec2cd527eb99a07d550a3e27178e9cae9e39f9889263cab25e17"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for vnet resource group.

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
vnet_resource_group = {}
```

<a id="canonical-6938b2e72f385c0bf7dea122411e29f599b17f9f640e9222186ecb16a5d30a45"></a>

## Direct properties — voltstack_cluster.az_nodes.local_subnet.subnet.vnet_resource_group / 0ea08730e8b6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9ecc730e578c736d4ed7430c4b9021942a14fac8ba18b7c985b8a2c908e13e7a"></a>

## Next pages — voltstack_cluster.az_nodes.local_subnet.subnet.vnet_resource_group / 0ea08730e8b6 / 4

- [voltstack_cluster.az_nodes.local_subnet.subnet](resources--azure_vnet_site--reference--group-008.md#canonical-b926010db376e38225d7e7ebceca7b5283fad4b6380ced347e24270b937efb08)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-8ab67dd0204d5e6753dee469e6b1cb4905e52125ce05241f90f4a9d0f3d44f7d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a4a1f8ccfe3ce10b4a764226ab249b2eef4d8c6464719c524d64aae06d838aa7"></a>

## voltstack_cluster.az_nodes.local_subnet.subnet_param — voltstack_cluster.az_nodes.local_subnet.subnet_param / 164cde1e982b / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [voltstack_cluster.az_nodes](resources--azure_vnet_site--reference--group-008.md#canonical-543235bd47f7e2a85d5558a07785780e10c0b9a46c89b4babad14a7dbac537c9)
- [voltstack_cluster.az_nodes.local_subnet](resources--azure_vnet_site--reference--group-008.md#canonical-5e068dcb356af048b2b5e6233f8e7eff9bc32b8cc5111e1ce380f6b07f44c4c9)
- voltstack_cluster.az_nodes.local_subnet.subnet_param

<a id="canonical-120bfe3049a1d05335c3c189dc45a7caecfa4b45aba2eb1e7e7f0f01fa34056e"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating a new cloud subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ipv4")}
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
subnet_param {
  # Configure direct properties listed below.
}
```

<a id="canonical-3e3a8e1b834f1024bb232879e239e00417a98c683913da6453880566b5f705ad"></a>

## Direct properties — voltstack_cluster.az_nodes.local_subnet.subnet_param / 164cde1e982b / 3

<a id="canonical-1957d3b6e2c91354a24d0f1890d4a4909cf16360dcd67487577d1572eb0dc345"></a>

<a id="canonical-754569f7f2992634b607e8012beb4280013c9386fb6986bda57c7a7cf96ff5f8"></a>

## ipv4 property — voltstack_cluster.az_nodes.local_subnet.subnet_param / 164cde1e982b / 4

Type: `"string"`. Optional.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-b36d10d0310c897c1588c000eafe577de3ada520f1d549ebac7859e5715815bc"></a>

## Next pages — voltstack_cluster.az_nodes.local_subnet.subnet_param / 164cde1e982b / 5

- [voltstack_cluster.az_nodes.local_subnet](resources--azure_vnet_site--reference--group-008.md#canonical-5e068dcb356af048b2b5e6233f8e7eff9bc32b8cc5111e1ce380f6b07f44c4c9)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-9d2d12a1b34075b677dd04161498868b24e0857164f1b741844d65f8cccd3d0a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-86880cd89ab92e9080a0471adfeabbdac2fcc5a3c889785babc07b00780058d6"></a>

## voltstack_cluster.dc_cluster_group — voltstack_cluster.dc_cluster_group / 07a8c8a45757 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- voltstack_cluster.dc_cluster_group

<a id="canonical-672ff5eb45af420ab9a04af84bc8711c408498dcac37e1e35cb0b5ec51e0ddb1"></a>

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
dc_cluster_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-1db668ce07314fc4021d2298b1e7584d3946c0315ee3cd4c317b466f68981dd7"></a>

## Direct properties — voltstack_cluster.dc_cluster_group / 07a8c8a45757 / 3

<a id="canonical-983d34187000c8616180826656e8f7f210be34896e9904683454d33cacfa1234"></a>

<a id="canonical-d87d174353491bf4744857df810593e81e09b0a2850557daeae657ade64a4bc9"></a>

## name property — voltstack_cluster.dc_cluster_group / 07a8c8a45757 / 4

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

<a id="canonical-7a5b0691f8225ed490656ff17d0ddaa966acc91fcb5ad719c4ebca27bd6a4a02"></a>

<a id="canonical-8adecbd7173dbbdf777d26cb415008b4f773bac59d8a226df59af6397db78eab"></a>

## namespace property — voltstack_cluster.dc_cluster_group / 07a8c8a45757 / 5

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

<a id="canonical-5d8667810900acec5e437cc8ab620fd154c61870ffb0f6c3cfbe99787c5b7470"></a>

<a id="canonical-c04356081145221811b54558287bf4c6ac69f1398546a02a401ec78332a58e3e"></a>

## tenant property — voltstack_cluster.dc_cluster_group / 07a8c8a45757 / 6

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

<a id="canonical-30c7a27032520473fc839eadb549855e7e6ef63d687b09e6af896ad7fb1d48f5"></a>

## Next pages — voltstack_cluster.dc_cluster_group / 07a8c8a45757 / 7

- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-ed31b36fb96ae3be15150a7719135aebdb54e2fdb1a7466cbe4c60a030e14425"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f7325437210496e1cf699b7bf78be168a2cf42a7e1e9ccfffbd887c5b8db4cd0"></a>

## voltstack_cluster.default_storage — voltstack_cluster.default_storage / 23167aa0d356 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- voltstack_cluster.default_storage

<a id="canonical-4e41219b6c24a4cfb4a8eb9c9f958480ec34646e09cd1a0c0239427e122f36a9"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default storage.

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
default_storage = {}
```

<a id="canonical-9d9d3330a06c3df26cc51f4a6467cd53cf22b5c7752fb991a4cf1faeeab7c101"></a>

## Direct properties — voltstack_cluster.default_storage / 23167aa0d356 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c277d8b92d76fcbbb2bb161fc0e1bdaeda808730b589fd57c2a97cded59cf61e"></a>

## Next pages — voltstack_cluster.default_storage / 23167aa0d356 / 4

- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-c15dea8b1c56cbbe431e0b31debcee5d177ff009eac2f2f3d62ac6ee36c4344f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-03bea655528d4143700e20801133d0c7e27081016addccc7a1a740cdd9f24a14"></a>

## voltstack_cluster.forward_proxy_allow_all — voltstack_cluster.forward_proxy_allow_all / e14d44caf1d7 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- voltstack_cluster.forward_proxy_allow_all

<a id="canonical-260cc8aa2e7bbb418144c5adc5ffcc8f34561601c7c7ef6aaa3bff01db233919"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for forward proxy allow all.

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
forward_proxy_allow_all = {}
```

<a id="canonical-c8e34771612d2e23771f7a5c709f20421bb5e13767e9da94093fbb0427c9968e"></a>

## Direct properties — voltstack_cluster.forward_proxy_allow_all / e14d44caf1d7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-066412e5e48b5ce4861924b5c7d1d3d03b4ce80528ccc738fc38ce0a31d2fadc"></a>

## Next pages — voltstack_cluster.forward_proxy_allow_all / e14d44caf1d7 / 4

- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-a8935381bd30aecbecc6a8cd035193d1f167660b2bf568f5541b0a720dd3966f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-393f0ff07c2d9e7f1d584b95b6ad76e8669e0dbce31a5bc4e3b39a8a7278be9f"></a>

## voltstack_cluster.global_network_list — voltstack_cluster.global_network_list / c267f62c9e31 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- voltstack_cluster.global_network_list

<a id="canonical-f068771a4de015179df61c453ddeb29e5ce9f75efa73fee717cc22d571c8c2e9"></a>

Type: `"object"`. single nested block, Optional.

Global Network Connection List. List of global network connections.

Upstream description:

List of global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("global_network_connections")}
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
global_network_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-fa2064984f6b214632558c5116158e2c7432149fff720c08e87593695d9b3d25"></a>

## Direct properties — voltstack_cluster.global_network_list / c267f62c9e31 / 3

- [global_network_connections](resources--azure_vnet_site--reference--group-008.md#canonical-4ff8031cf4922b4a592b1ca3ba34d1aa0e0c60d25f3a9ad96fc23e1842495b51): complete subsection reference.

<a id="canonical-14a12acca1d4bd8e15586b725f6cbc2682621da85eec0542bd1fca72f5f1f041"></a>

## Next pages — voltstack_cluster.global_network_list / c267f62c9e31 / 4

- [voltstack_cluster.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-008.md#canonical-4ff8031cf4922b4a592b1ca3ba34d1aa0e0c60d25f3a9ad96fc23e1842495b51)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-4ff8031cf4922b4a592b1ca3ba34d1aa0e0c60d25f3a9ad96fc23e1842495b51"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-850d8d9301f997154694ca6b62b7096bf9b9acc1b4b81390a92156d2fa052438"></a>

## voltstack_cluster.global_network_list.global_network_connections — voltstack_cluster.global_network_list.global_network_connections / 30d212ef0e13 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [voltstack_cluster.global_network_list](resources--azure_vnet_site--reference--group-008.md#canonical-a8935381bd30aecbecc6a8cd035193d1f167660b2bf568f5541b0a720dd3966f)
- voltstack_cluster.global_network_list.global_network_connections

<a id="canonical-52d7169bd68c50b94bd8996a9484d304c691acd64dd7d5b2a4948f5d0228f643"></a>

Type: `"object"`. list nested block, Optional.

Global Network Connections. Global network connections.

Upstream description:

Global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("sli_to_global_dr",
    "slo_to_global_dr")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
global_network_connections {
  # Configure direct properties listed below.
}
```

<a id="canonical-dc16d8b50b3737cd313ced2ab7150a16434c1329eb8583d44c1cf94f58fa0fc9"></a>

## Direct properties — voltstack_cluster.global_network_list.global_network_connections / 30d212ef0e13 / 3

- [sli_to_global_dr](resources--azure_vnet_site--reference--group-008.md#canonical-2b644667079b6ad86e305491b3705941401af9618ef4b1b923667c9b171d1574): complete subsection reference.

- [slo_to_global_dr](resources--azure_vnet_site--reference--group-008.md#canonical-3b662671275f7b205177e6e1937ee5400435f8e00316a9b29d3fc506ca0f58c3): complete subsection reference.

<a id="canonical-408860988740102aa27e8c07b5adf49913871657354a202e7b817100e9d8753b"></a>

## Next pages — voltstack_cluster.global_network_list.global_network_connections / 30d212ef0e13 / 4

- [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](resources--azure_vnet_site--reference--group-008.md#canonical-2b644667079b6ad86e305491b3705941401af9618ef4b1b923667c9b171d1574)
- [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](resources--azure_vnet_site--reference--group-008.md#canonical-3b662671275f7b205177e6e1937ee5400435f8e00316a9b29d3fc506ca0f58c3)
- [voltstack_cluster.global_network_list](resources--azure_vnet_site--reference--group-008.md#canonical-a8935381bd30aecbecc6a8cd035193d1f167660b2bf568f5541b0a720dd3966f)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-2b644667079b6ad86e305491b3705941401af9618ef4b1b923667c9b171d1574"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-985364cceeb08707603048e88811aa3e16afc59cd92c6b2944049b0196a08d3a"></a>

## voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr — voltstack_cluster.global_network_list.global_network_connections.sli_to_global_d / 33a4d5ed46a0 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [voltstack_cluster.global_network_list](resources--azure_vnet_site--reference--group-008.md#canonical-a8935381bd30aecbecc6a8cd035193d1f167660b2bf568f5541b0a720dd3966f)
- [voltstack_cluster.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-008.md#canonical-4ff8031cf4922b4a592b1ca3ba34d1aa0e0c60d25f3a9ad96fc23e1842495b51)
- voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr

<a id="canonical-e9936e8564c062480a2e28dc472b6d2600c8e7038cacf232554199348e85a56d"></a>

Type: `"object"`. single nested block, Optional.

Global network reference for direct connection.

Receipt-pinned upstream constraints:

```json
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
sli_to_global_dr {
  # Configure direct properties listed below.
}
```

<a id="canonical-e5e97316beb05606c1c94a5c4f78f50eeb1696f55882584e0c46b26a84013417"></a>

## Direct properties — voltstack_cluster.global_network_list.global_network_connections.sli_to_global_d / 33a4d5ed46a0 / 3

- [global_vn](resources--azure_vnet_site--reference--group-008.md#canonical-0cacfff6e5c151b140462797048b4ceee4a4098b8075da67a81024660f100e08): complete subsection reference.

<a id="canonical-ed7dc675ca55d5ab6a525fca990743e73c60ef1fe94c60046cf0857fe879bc61"></a>

## Next pages — voltstack_cluster.global_network_list.global_network_connections.sli_to_global_d / 33a4d5ed46a0 / 4

- [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--azure_vnet_site--reference--group-008.md#canonical-0cacfff6e5c151b140462797048b4ceee4a4098b8075da67a81024660f100e08)
- [voltstack_cluster.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-008.md#canonical-4ff8031cf4922b4a592b1ca3ba34d1aa0e0c60d25f3a9ad96fc23e1842495b51)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-0cacfff6e5c151b140462797048b4ceee4a4098b8075da67a81024660f100e08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d46f8f8e71ef606f4fc4a9939d86635a1acdf36ca8b5fdbf4ceced534c6d2dcf"></a>

## voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn — voltstack_cluster.global_network_list.global_network_connections.sli_to_global_d / e4f302f1f844 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [voltstack_cluster.global_network_list](resources--azure_vnet_site--reference--group-008.md#canonical-a8935381bd30aecbecc6a8cd035193d1f167660b2bf568f5541b0a720dd3966f)
- [voltstack_cluster.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-008.md#canonical-4ff8031cf4922b4a592b1ca3ba34d1aa0e0c60d25f3a9ad96fc23e1842495b51)
- [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](resources--azure_vnet_site--reference--group-008.md#canonical-2b644667079b6ad86e305491b3705941401af9618ef4b1b923667c9b171d1574)
- voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn

<a id="canonical-aa1bc7e242373ad8132dd098858a3369f038be358d1010c6013762f37ef33514"></a>

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
global_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-25c6152b3137c5d9b992528d2eb38898bca095725862be3a1cca8a78e8d89173"></a>

## Direct properties — voltstack_cluster.global_network_list.global_network_connections.sli_to_global_d / e4f302f1f844 / 3

<a id="canonical-6337f74a98203dc16e72d54b4438180518a0e0667b4c39b87ff7fafaa4d1a73e"></a>

<a id="canonical-4ac63fb6d92dc511521d9bcbd0fe6414a10dca80d9961f514651ae06849f0b4c"></a>

## name property — voltstack_cluster.global_network_list.global_network_connections.sli_to_global_d / e4f302f1f844 / 4

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

<a id="canonical-bc22162216eeb123c543571b970da09229e3b4c88435efe86c93d9db31cf6a06"></a>

<a id="canonical-c8f73fecd1d0d0609d3a70189d7b05b962dd5ae5693fdd76f1e0ca3c64820554"></a>

## namespace property — voltstack_cluster.global_network_list.global_network_connections.sli_to_global_d / e4f302f1f844 / 5

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

<a id="canonical-6d8241b79d106003e9e5044577d42f7869ea113adadfa4abfd8fbf85b7d8d2b7"></a>

<a id="canonical-360ffd40d44b953af3f4b0f019974ec4e0c209a45485fc186f71b56281297069"></a>

## tenant property — voltstack_cluster.global_network_list.global_network_connections.sli_to_global_d / e4f302f1f844 / 6

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

<a id="canonical-85f3fc18b0d2a5caf76743e8f06b6c584c859bcd84c5e88f54166d49d1ff0631"></a>

## Next pages — voltstack_cluster.global_network_list.global_network_connections.sli_to_global_d / e4f302f1f844 / 7

- [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](resources--azure_vnet_site--reference--group-008.md#canonical-2b644667079b6ad86e305491b3705941401af9618ef4b1b923667c9b171d1574)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-3b662671275f7b205177e6e1937ee5400435f8e00316a9b29d3fc506ca0f58c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8ef13561fb6f7239f0fba20ea532fa5f9a7df3a15721bc4295bcf776d4c7232"></a>

## voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr — voltstack_cluster.global_network_list.global_network_connections.slo_to_global_d / 80177bc5105b / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [voltstack_cluster.global_network_list](resources--azure_vnet_site--reference--group-008.md#canonical-a8935381bd30aecbecc6a8cd035193d1f167660b2bf568f5541b0a720dd3966f)
- [voltstack_cluster.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-008.md#canonical-4ff8031cf4922b4a592b1ca3ba34d1aa0e0c60d25f3a9ad96fc23e1842495b51)
- voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr

<a id="canonical-38957fc8b4c175f2a3c761b3fc5ca7944034a785149724b8c5638f73cc36df37"></a>

Type: `"object"`. single nested block, Optional.

Global network reference for direct connection.

Receipt-pinned upstream constraints:

```json
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
slo_to_global_dr {
  # Configure direct properties listed below.
}
```

<a id="canonical-5ed9173d21874623a34528c4dd1045274f0cdae5657ec6d52a1125d921db92e8"></a>

## Direct properties — voltstack_cluster.global_network_list.global_network_connections.slo_to_global_d / 80177bc5105b / 3

- [global_vn](resources--azure_vnet_site--reference--group-008.md#canonical-c6221cf660dad3d81b33a876a66a799c9352b2635d0caa6b35a2d51f2b7def7a): complete subsection reference.

<a id="canonical-433b411e0fa7f495ecd4f0b0cab5cf2480d61d46ae6296404d98d2016995e547"></a>

## Next pages — voltstack_cluster.global_network_list.global_network_connections.slo_to_global_d / 80177bc5105b / 4

- [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--azure_vnet_site--reference--group-008.md#canonical-c6221cf660dad3d81b33a876a66a799c9352b2635d0caa6b35a2d51f2b7def7a)
- [voltstack_cluster.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-008.md#canonical-4ff8031cf4922b4a592b1ca3ba34d1aa0e0c60d25f3a9ad96fc23e1842495b51)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-c6221cf660dad3d81b33a876a66a799c9352b2635d0caa6b35a2d51f2b7def7a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7ce369d42b938491321d48f1e5b404329b7cc4265a0e6d2b907f93b250cd188"></a>

## voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn — voltstack_cluster.global_network_list.global_network_connections.slo_to_global_d / e571cca46e96 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [voltstack_cluster.global_network_list](resources--azure_vnet_site--reference--group-008.md#canonical-a8935381bd30aecbecc6a8cd035193d1f167660b2bf568f5541b0a720dd3966f)
- [voltstack_cluster.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-008.md#canonical-4ff8031cf4922b4a592b1ca3ba34d1aa0e0c60d25f3a9ad96fc23e1842495b51)
- [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](resources--azure_vnet_site--reference--group-008.md#canonical-3b662671275f7b205177e6e1937ee5400435f8e00316a9b29d3fc506ca0f58c3)
- voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn

<a id="canonical-32ed6c86cf40a4a47871917a2077639c128452dbc2a377d16c81d5efdf017c50"></a>

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
global_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-26d3aa357511c3c7a0f1467a5e095ecc8b1a863bb6ff5167e50a83d47b074461"></a>

## Direct properties — voltstack_cluster.global_network_list.global_network_connections.slo_to_global_d / e571cca46e96 / 3

<a id="canonical-c5bbf8748503ee8d35249fa15642db87f8d4615bd59e5f52c97fb6233b8385ae"></a>

<a id="canonical-451159d0783cbf7a5dfb3a1e330533a0a802b94a438a9d24fbad847107044109"></a>

## name property — voltstack_cluster.global_network_list.global_network_connections.slo_to_global_d / e571cca46e96 / 4

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

<a id="canonical-a4b304ce51e360784340c251a036fa9c5caca8952846ed6382a9fb51b8f16b90"></a>

<a id="canonical-94677945c053127d3dc134100eebfec742e44f2e5bb05926a24dc3552864adb4"></a>

## namespace property — voltstack_cluster.global_network_list.global_network_connections.slo_to_global_d / e571cca46e96 / 5

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

<a id="canonical-9944adc031453762b608288e4b8acd82f080f17642e085b06e626ad1c3000c6b"></a>

<a id="canonical-13b21faa4baccc6fa454f518a4056191006b1733a889b16f0710886b9a0dabaf"></a>

## tenant property — voltstack_cluster.global_network_list.global_network_connections.slo_to_global_d / e571cca46e96 / 6

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

<a id="canonical-bb9aec1e275d3ad8dab858a16474eb645bf03ac3ba78e3c53bef58d6be904a69"></a>

## Next pages — voltstack_cluster.global_network_list.global_network_connections.slo_to_global_d / e571cca46e96 / 7

- [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](resources--azure_vnet_site--reference--group-008.md#canonical-3b662671275f7b205177e6e1937ee5400435f8e00316a9b29d3fc506ca0f58c3)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-d30e0699286fdea85a05f142127c4e9412c34f7bfc1257205aae6153dfa5c193"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1f1978d384a49ee71982bcdcad7c94423288726bad2e9794a9fc9f36c5abd55"></a>

## voltstack_cluster.k8s_cluster — voltstack_cluster.k8s_cluster / a267cbb72ae9 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- voltstack_cluster.k8s_cluster

<a id="canonical-520e534050b9dbeeb62ce579a73c966ec4a07c3371a5caf91971951ff60501ca"></a>

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
k8s_cluster {
  # Configure direct properties listed below.
}
```

<a id="canonical-8e4794d511e4137ebf09d6a9cce78b9778fedfe3fe83d2bcd45976fd1a706bda"></a>

## Direct properties — voltstack_cluster.k8s_cluster / a267cbb72ae9 / 3

<a id="canonical-84db2b6dffbeb9d357041cec53eae06869c5058f5fd97802a92def9914cd3d0c"></a>

<a id="canonical-caad6cb98a0a249b36d0419059acf7aa5dd2f364eb280cf1257e32969aef33f1"></a>

## name property — voltstack_cluster.k8s_cluster / a267cbb72ae9 / 4

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

<a id="canonical-304831524937ba2150c23e4e2595b0141c62624505e94af5076900c0995b8ea0"></a>
