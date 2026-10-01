---
page_title: "xcsh_fleet reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fleet reference."
---

# xcsh_fleet reference

<a id="canonical-11aeca06972c4fcc8713f8a4c687e5ee168b3b1365978397d620fe7e7bd19257"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 63c5564f23fd / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-9aaa33d28bdc9675ddfeacd2a1549952db4fb3dee17435ce405d097f40edb5d1)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--fleet--reference--group-003.md#canonical-24acca475feb36d716873496a71007399f5587fcaf785096f7fe139ff5078747)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret

<a id="canonical-a39974bf30becc8ec8833c7855629f8259f1203af39f57135e0b5b8734cd4e31"></a>

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
chap_target_initiator_secret {
  # Configure direct properties listed below.
}
```

<a id="canonical-eff36d919b689a38091f44bac79bf5bfa54bc3f6b5ce6c17875e67071e3c4103"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 63c5564f23fd / 3

- [blindfold_secret_info](resources--fleet--reference--group-004.md#canonical-19c9e30b47473f1f93edd1c9e596d87ebe280b3e8441aeb7a45033ee7f495a59): complete subsection reference.

- [clear_secret_info](resources--fleet--reference--group-004.md#canonical-47a152eb820350332593082b06469bd98f530be59044f722f8aadfb6dec5e2a5): complete subsection reference.

<a id="canonical-b6d757b6b79c38442a50dabd3ed8f1301647ffcbe99a2ebd4eb4e7f9e98f31ce"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 63c5564f23fd / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info](resources--fleet--reference--group-004.md#canonical-19c9e30b47473f1f93edd1c9e596d87ebe280b3e8441aeb7a45033ee7f495a59)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info](resources--fleet--reference--group-004.md#canonical-47a152eb820350332593082b06469bd98f530be59044f722f8aadfb6dec5e2a5)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--fleet--reference--group-003.md#canonical-24acca475feb36d716873496a71007399f5587fcaf785096f7fe139ff5078747)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-19c9e30b47473f1f93edd1c9e596d87ebe280b3e8441aeb7a45033ee7f495a59"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3972dbcc79783d76a6a09de26caa48c96c37c74f85364686bbd836acac9394cd"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / ad95fee57e39 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-9aaa33d28bdc9675ddfeacd2a1549952db4fb3dee17435ce405d097f40edb5d1)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--fleet--reference--group-003.md#canonical-24acca475feb36d716873496a71007399f5587fcaf785096f7fe139ff5078747)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](resources--fleet--reference--group-003.md#canonical-32cee025bc0732a76806a9f726b842427d21551a8d82acf9b5358db06b14163d)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info

<a id="canonical-18ab81b288ba6c03c8edcfb65dadca79487d3a34146078da4d20eaaa508e1cef"></a>

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

<a id="canonical-17e5bfb1b8aed5a824ea3e52c840942570837a068fd81e08110db85b7ae98bc1"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / ad95fee57e39 / 3

<a id="canonical-81c0a113b15d2c8a631f53ad9206e40fec006bfd909034f19414f06dabd75286"></a>

<a id="canonical-5d1869c6d1c43cbc4506b41c4c62553e18b82d3956cc7e4af1e802918c7b51a7"></a>

## decryption_provider property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / ad95fee57e39 / 4

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

<a id="canonical-212084729941c95fa650aa548e6d44bdac9870a0371512fa40fb1e01fb2c294a"></a>

<a id="canonical-b8e78a5f26e21595e21d8d58179fd07980f0b962fe8f8368ad6fd8b405b40011"></a>

## location property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / ad95fee57e39 / 5

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

<a id="canonical-9fc00afd4c97525d240f2b58866e8bef87cdea7d0c62be96b3fdf2639ce1aa07"></a>

<a id="canonical-abe85127fd1d09908234ba7331b36084336f96055ea41b21e13f4aaa94eba2f4"></a>

## store_provider property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / ad95fee57e39 / 6

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

<a id="canonical-fb07ed9adf1916104040ecf31fed3015229d8247c36d0b77a2a1177209b39594"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / ad95fee57e39 / 7

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](resources--fleet--reference--group-003.md#canonical-32cee025bc0732a76806a9f726b842427d21551a8d82acf9b5358db06b14163d)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-47a152eb820350332593082b06469bd98f530be59044f722f8aadfb6dec5e2a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b40728f0fca561d26dcfa6b7913c5a6b5564b4f84c87cc2eea2afc5201cebecd"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / f441cc62cd84 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-9aaa33d28bdc9675ddfeacd2a1549952db4fb3dee17435ce405d097f40edb5d1)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--fleet--reference--group-003.md#canonical-24acca475feb36d716873496a71007399f5587fcaf785096f7fe139ff5078747)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](resources--fleet--reference--group-003.md#canonical-32cee025bc0732a76806a9f726b842427d21551a8d82acf9b5358db06b14163d)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info

<a id="canonical-29394c83a341b03a174b410ace8310ee5284ccbb760ad280bee0d5ec65dd6aa9"></a>

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

<a id="canonical-3264afa1debd25eba9a8dd802c5ed199a8be5c7e86c25dec1474b3fa60aed921"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / f441cc62cd84 / 3

<a id="canonical-6a2cbfbf5f171e24cb9de47e8c82ca8eaaeece82d7341b3bc1e1403d10575d97"></a>

<a id="canonical-cc59d6ba044f392bb65fa58b1b09b37e985ba762fd2dcec50eaa95329604bab8"></a>

## provider_ref property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / f441cc62cd84 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2dce8fb9e609e156930fcb37368c256ecc6fe81e75b80c02d63283cda035dd52"></a>

<a id="canonical-120eed68d841cacbd7fdbd22d74dae1cbe73bb85e6492220a09ec2723e634bd5"></a>

## url property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / f441cc62cd84 / 5

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

<a id="canonical-57486d74894ba71ab90b7976a106a841364f8159b6cab30b8cfcd32c412548da"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / f441cc62cd84 / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](resources--fleet--reference--group-003.md#canonical-32cee025bc0732a76806a9f726b842427d21551a8d82acf9b5358db06b14163d)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-4bc1400cfae06aa3f69593fc2f6c01a537183649f34d046d0a1fbf9411e01f29"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bcb748d066b8c67027f79cbcd2ab39cd2367754f6ca5d7b93c62ff7a71b33175"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 756601f1f61c / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-9aaa33d28bdc9675ddfeacd2a1549952db4fb3dee17435ce405d097f40edb5d1)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults

<a id="canonical-a0ae7a818775106128a430651d7b71169d3f0335e49478f4293e5ff64d22cf28"></a>

Type: `"object"`. single nested block, Optional.

It controls how each volume is provisioned by default using these OPTIONS in a special section of
the configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "no_qos"),
  validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "qos_policy"),
  validators.ConflictingObjectAttributes("no_qos",
    "qos_policy")}
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
  "x-ves-oneof-field-qos_policy_choice": "[\"adaptive_qos_policy\",\"no_qos\",\"qos_policy\"]"
}
```

Terraform syntax:

```terraform
volume_defaults {
  # Configure direct properties listed below.
}
```

<a id="canonical-97129dfd6e2d1161da9becfb35ed6adfa72632d9f426d43f43729392e412c8b8"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 756601f1f61c / 3

<a id="canonical-6f3e941133826af29e60126b6d82e0d67e52830a559be2a2f9aef3d282e2e825"></a>

<a id="canonical-89abe9dd0c3246909cbcc59b96aad667ee9bf95ee1b016f48be5b66c95465383"></a>

## adaptive_qos_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 756601f1f61c / 4

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Exclusive with \[no\_qos qos\_policy\] Enter Adaptive QoS Policy Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-f8e0658e0c5b0b0c96e476d1bfe60b43abab5e08a63ff22671eee845d8818c61"></a>

<a id="canonical-06ea71e45ce94b42ad0c53c2955561e420b4452291d07b24e152c14250f299f2"></a>

## encryption property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 756601f1f61c / 5

Type: `"bool"`. Optional.

Enable Encryption. Enable NetApp volume encryption.

Upstream description:

Enable NetApp volume encryption.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-5268e3f39f6d9fb6099f28887a231a207f6a0228e60f20c6481b2751470501e0"></a>

<a id="canonical-9b5a916471c71afa869aa0943704fd6904d5c78275afb26645d6a31c3f5e96e6"></a>

## export_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 756601f1f61c / 6

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Export policy to use.

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

- [no_qos](resources--fleet--reference--group-004.md#canonical-041e032bf4fd0ba5affd962f3b55a8072738866777ef7135cdcc223ec13e2fba): complete subsection reference.

<a id="canonical-3a16b28a9ab6eb4a976857070372ecaedd1869bfd67125b44004e45565194399"></a>

<a id="canonical-e676c0fc933b61c5792bc15a1f8f37dd641fec398df091bdae5491b526eece8e"></a>

## qos_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 756601f1f61c / 7

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Exclusive with \[adaptive\_qos\_policy no\_qos\] Enter QoS Policy Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-b4fed030be70bde1648427dfffd91dac77ee0046c10ec256365dac68b6408ccc"></a>

<a id="canonical-ee949d7865aa019b0100a0c117ae049e70fef52e19fcf706deacbc34701a599b"></a>

## security_style property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 756601f1f61c / 8

Type: `"string"`. Optional.

Security Style. Security style for new volumes.

Upstream description:

Security style for new volumes.

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

<a id="canonical-79dd8006a746350204fcf269a962e3a2266c7bc11f3d7a632d9ce5a190b3665c"></a>

<a id="canonical-516cb285cb0a0b0000d041064c8b787d7d4b0b1fbd79ae21f693e09e24d3c01f"></a>

## snapshot_dir property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 756601f1f61c / 9

Type: `"bool"`. Optional.

Access to Snapshot Directory. Access to the .snapshot directory.

Upstream description:

Access to the .snapshot directory.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-d00234cd3d74a0fbcc9cb06b652ec7e02fa2f866ad3738460077b92acbbfccb3"></a>

<a id="canonical-ce766301f0f299de5086957917b1c3e566f71704fbd12ca8074b69d2632431ca"></a>

## snapshot_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 756601f1f61c / 10

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Snapshot policy to use.

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

<a id="canonical-47630ab2deb41e18e6911b72189ddb3bb06121b2638f29b27a2723abf1d9d1e2"></a>

<a id="canonical-0757961ec1d90992a4e987f8649af95adbb95dae25dda842698de4cdaa9ce58e"></a>

## snapshot_reserve property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 756601f1f61c / 11

Type: `"string"`. Optional.

Percentage of volume reserved for snapshots. '0' if snapshot policy is 'none', else ''.

Upstream description:

Percentage of volume reserved for snapshots. "0" if snapshot policy is "none", else ""

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

<a id="canonical-3f63d6571b89c6a295c887830d911ab39f717ea900086dc6958615ab312a01a8"></a>

<a id="canonical-d97794e1fadcfca7d5e2611cfea591f1a5517ae95c567982022c3bc0e97f2eb2"></a>

## space_reserve property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 756601f1f61c / 12

Type: `"string"`. Optional.

\[Enum: none|thick\] Space reservation mode; “none” (thin) or “volume” (thick). Possible values are
\`none\`, \`thick\`.

Upstream description:

Space reservation mode; “none” (thin) or “volume” (thick)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("none",
    "thick"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "none",
    "thick"
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
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  }
}
```

<a id="canonical-6d109a65cd5e36252a8a0d51e4f8c9dd45370aed4a7e80d2456f2e4e31d8006f"></a>

<a id="canonical-ef40960aeb85cee02b264bc2bd827773246be53235a9586612b30a55464eff5e"></a>

## split_on_clone property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 756601f1f61c / 13

Type: `"bool"`. Optional.

Split a clone from its parent upon creation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-e8c00df9ebd2cf6fe710d3bebdfa38f3d7c9e62da0591f0ed8a71a086330058d"></a>

<a id="canonical-bb74d64aa387cf63f69d95b06ae8252e9cd9de09b19f00baec709ac99972a2f5"></a>

## tiering_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 756601f1f61c / 14

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Tiering policy to use. "none" is default.

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

<a id="canonical-24930dec69be98f027ef93245f6609ec27b66216db6827f12aecdf545a4be826"></a>

<a id="canonical-54d54a3a40f2da5fa5fd31c09581778861ed4c249b7ddc62e19c60dd291605b0"></a>

## unix_permissions property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 756601f1f61c / 15

Type: `"number"`. Optional.

Unix permission mode for new volumes. All allowed 777.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-5b3afd89a0427afde96a8600893e22ea69a1adeb84bd2a86f9f8a3c461391756"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 756601f1f61c / 16

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos](resources--fleet--reference--group-004.md#canonical-041e032bf4fd0ba5affd962f3b55a8072738866777ef7135cdcc223ec13e2fba)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-9aaa33d28bdc9675ddfeacd2a1549952db4fb3dee17435ce405d097f40edb5d1)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-041e032bf4fd0ba5affd962f3b55a8072738866777ef7135cdcc223ec13e2fba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cb095bfcb109c7192ad037b6addffc9bff933e1a74b8cba39d2f4bbd936c022a"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / fd998a0fceec / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-9aaa33d28bdc9675ddfeacd2a1549952db4fb3dee17435ce405d097f40edb5d1)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults](resources--fleet--reference--group-004.md#canonical-4bc1400cfae06aa3f69593fc2f6c01a537183649f34d046d0a1fbf9411e01f29)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos

<a id="canonical-43a778e31d02fc7b7853163a83239f1532cdf0007b1a94428dba6f7d91298a56"></a>

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
no_qos = {}
```

<a id="canonical-ed8c6f9d2261510ef77e9105b8b89798ada028c8742e16f9649979cab131014f"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / fd998a0fceec / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6db63ff2c04f5560d940bc760de190863b19c5b28eba073acc90e37caca28797"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / fd998a0fceec / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults](resources--fleet--reference--group-004.md#canonical-4bc1400cfae06aa3f69593fc2f6c01a537183649f34d046d0a1fbf9411e01f29)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-02b188fbd08df4634f53310c47e1459dbcbcbf07675d4d48c3343bdad2ba9044"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1820a7a42aec3ee39a4f6189692a5e8fa8a02dd1e6b0e6092544a4d6e4abce44"></a>

## storage_device_list.storage_devices.pure_service_orchestrator — storage_device_list.storage_devices.pure_service_orchestrator / c54799de53b6 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- storage_device_list.storage_devices.pure_service_orchestrator

<a id="canonical-14de148481b302b8ff4bd9c3e2ea76e83cd523af856c343bdbdaa7a1122edd40"></a>

Type: `"object"`. single nested block, Optional.

Device configuration for Pure Storage Service Orchestrator.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cluster_id")}
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
pure_service_orchestrator {
  # Configure direct properties listed below.
}
```

<a id="canonical-6bcd8a8f22915407cd8473e8dbbfdbe4cbcb4365b8b444209cd8be67b19085ca"></a>

## Direct properties — storage_device_list.storage_devices.pure_service_orchestrator / c54799de53b6 / 3

- [arrays](resources--fleet--reference--group-004.md#canonical-558d5d8c784c45808aa7983b2b15b646d931426b9f4bb4de8b307ae1f82d078b): complete subsection reference.

<a id="canonical-7d50f7ebb09768647e049bf21aed47b8d11fa08641612c6f586fdfc7f254e477"></a>

<a id="canonical-27e26f13aedf05b74ba1d6bec784816f21375171cf35211989cfb19453501012"></a>

## cluster_id property — storage_device_list.storage_devices.pure_service_orchestrator / c54799de53b6 / 4

Type: `"string"`. Optional.

ClusterID is added as a prefix for all volumes created by this PSO installation. ClusterID is also
used to identify the volumes used by the datastore, pso-db. ClusterID MUST BE UNIQUE for multiple
K8s clusters running on top of the same storage arrays.

Upstream description:

ClusterID is added as a prefix for all volumes created by this PSO installation. ClusterID is also
used to identify the volumes used by the datastore, pso-db. ClusterID MUST BE UNIQUE for multiple
K8s clusters running on top of the same storage arrays. Characters allowed: alphanumeric and
underscores.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 22),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 22,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 22,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z0-9_]*$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "22",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9_]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "22",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9_]*$"
  }
}
```

<a id="canonical-559381884f22fe7b5d55da6ddd58ff7679ea9ffb8f26e97857ac9b3aef14e3cd"></a>

<a id="canonical-1447ca6358a4a8cc0e460eaee75e482ae903eda514d98407fa4334badc866458"></a>

## enable_storage_topology property — storage_device_list.storage_devices.pure_service_orchestrator / c54799de53b6 / 5

Type: `"bool"`. Optional.

Option is to enable/disable the csi topology feature for pso-csi.

Upstream description:

This option is to enable/disable the csi topology feature for pso-csi.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-57c44d71a25faa26758e8ae2266b8ad8223d0ee09516ed57362b63f3ffd4e8fe"></a>

<a id="canonical-bef59e1b6003042ab68a824e92113f510df926247cc22a11d9be26b226226393"></a>

## enable_strict_topology property — storage_device_list.storage_devices.pure_service_orchestrator / c54799de53b6 / 6

Type: `"bool"`. Optional.

Option is to enable/disable the strict csi topology feature for pso-csi.

Upstream description:

This option is to enable/disable the strict csi topology feature for pso-csi.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-9afaa92b37124454cf14c440d6849e6b7eaebd804ab4a81a6edde0bb5d8b9443"></a>

## Next pages — storage_device_list.storage_devices.pure_service_orchestrator / c54799de53b6 / 7

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--fleet--reference--group-004.md#canonical-558d5d8c784c45808aa7983b2b15b646d931426b9f4bb4de8b307ae1f82d078b)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-558d5d8c784c45808aa7983b2b15b646d931426b9f4bb4de8b307ae1f82d078b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43a646fa19640cdbc534dece7817704ca1f838f88517115ea172e10656a10a58"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays — storage_device_list.storage_devices.pure_service_orchestrator.arrays / 346f5f82dd7e / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.pure_service_orchestrator](resources--fleet--reference--group-004.md#canonical-02b188fbd08df4634f53310c47e1459dbcbcbf07675d4d48c3343bdad2ba9044)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays

<a id="canonical-6a87b82b39e53b9691486994ac9509b2d885d19f1a6f869d5837cb2baa77f482"></a>

Type: `"object"`. single nested block, Optional.

Arrays Configuration. Device configuration for PSO Arrays.

Upstream description:

Device configuration for PSO Arrays.

Receipt-pinned upstream constraints:

```json
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
arrays {
  # Configure direct properties listed below.
}
```

<a id="canonical-38a3577cc062171f77065bc8bc448869af6d677fbde24780b98fa2a4be521928"></a>

## Direct properties — storage_device_list.storage_devices.pure_service_orchestrator.arrays / 346f5f82dd7e / 3

- [flash_array](resources--fleet--reference--group-004.md#canonical-f499462f23484c13046d8a0cacebbfdd683bc47434362ec63390322486bd5e09): complete subsection reference.

- [flash_blade](resources--fleet--reference--group-004.md#canonical-34cb05c8510e2ca5ded4b76162fe66699380b39e32176756ed229c5362ee62b1): complete subsection reference.

<a id="canonical-ff7835c341c6b751a4104447a08868ed3651465384079b5a68418db9e21dc1a7"></a>

## Next pages — storage_device_list.storage_devices.pure_service_orchestrator.arrays / 346f5f82dd7e / 4

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](resources--fleet--reference--group-004.md#canonical-f499462f23484c13046d8a0cacebbfdd683bc47434362ec63390322486bd5e09)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](resources--fleet--reference--group-004.md#canonical-34cb05c8510e2ca5ded4b76162fe66699380b39e32176756ed229c5362ee62b1)
- [storage_device_list.storage_devices.pure_service_orchestrator](resources--fleet--reference--group-004.md#canonical-02b188fbd08df4634f53310c47e1459dbcbcbf07675d4d48c3343bdad2ba9044)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-f499462f23484c13046d8a0cacebbfdd683bc47434362ec63390322486bd5e09"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2935724ed2d735407f64844c2b7293f706245e921354aed226b8891f979e5b8"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 046c2aa5cdda / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.pure_service_orchestrator](resources--fleet--reference--group-004.md#canonical-02b188fbd08df4634f53310c47e1459dbcbcbf07675d4d48c3343bdad2ba9044)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--fleet--reference--group-004.md#canonical-558d5d8c784c45808aa7983b2b15b646d931426b9f4bb4de8b307ae1f82d078b)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array

<a id="canonical-4954b763738d6bf182a30ee330b1344a862fbaa33cabbe61d18ebad90473282c"></a>

Type: `"object"`. single nested block, Optional.

Specify what storage flash arrays should be managed the plugin.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("default_fs_type",
    "flash_arrays",
    "iscsi_login_timeout",
    "san_type")}
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
flash_array {
  # Configure direct properties listed below.
}
```

<a id="canonical-1eb9b6b5446225a014c12dfc69b909298e95380345b66f62d02514d57a00f456"></a>

## Direct properties — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 046c2aa5cdda / 3

<a id="canonical-49904b0bd4f8add33c08e33c61c7e633914dae1765638419fa1a0b28d38edd97"></a>

<a id="canonical-313f03333f7b9e8e7e11dc6ab92b354984642974d813e8ac2e8b23ff391b7c54"></a>

## default_fs_opt property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 046c2aa5cdda / 4

Type: `"string"`. Optional.

Block volume default mkfs OPTIONS. Not recommended to change!

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-c94642ce0c5acf73113766442063095192991ec052ce7c4cff1cf1330d8011f1"></a>

<a id="canonical-f06834225b59b33d58013ce70a169b03ffdef0b69c4af1504460f63d7af6bacc"></a>

## default_fs_type property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 046c2aa5cdda / 5

Type: `"string"`. Optional.

\[Enum: xfs|ext4\] Block volume default filesystem type. Not recommended to change!. Possible values
are \`xfs\`, \`ext4\`.

Upstream description:

Block volume default filesystem type. Not recommended to change!

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("xfs",
    "ext4"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "xfs",
    "ext4"
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
    "ves.io.schema.rules.string.in": "[\\\"xfs\\\",\\\"ext4\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"xfs\\\",\\\"ext4\\\"]"
  }
}
```

<a id="canonical-cd2fe82e141c4d02025cb196c99f2ec1738d259954f7458b45473df348957908"></a>

<a id="canonical-1d620fb835e7307b34d1d5d5cd1cca8012019a6785ad86036b46c092469fb9c4"></a>

## default_mount_opts property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 046c2aa5cdda / 6

Type: `["list", "string"]`. Optional.

Block volume default filesystem mount OPTIONS. Not recommended to change!

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
```

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-5ec703ee016eac4735ddeb45e45a5b4ea5696b59cab613e8157e0589987137bb"></a>

<a id="canonical-280e6206ec47c64cbbdcd4a3ad92fa8142558d4084c6f4330264d39174cfb337"></a>

## disable_preempt_attachments property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 046c2aa5cdda / 7

Type: `"bool"`. Optional.

Disable Preempt Attachments. Enable/Disable attachment preemption!

Upstream description:

Enable/Disable attachment preemption!

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [flash_arrays](resources--fleet--reference--group-004.md#canonical-da7bfa1a7036624714eb7b92025dda342ba14a006026764de38b8f3c59fb569f): complete subsection reference.

<a id="canonical-b9d936c399ab73eec9794988a75bc62e9660811612733ec614dbf64dd43d3193"></a>

<a id="canonical-fb430eb5aaca995ffc51394bc7757651716b17bd5801e3c9c1fe2de01498f259"></a>

## iscsi_login_timeout property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 046c2aa5cdda / 8

Type: `"number"`. Optional.

ISCSI login timeout in seconds. Not recommended to change!

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 100),
}
```

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-d7136f046bd0bf9ef154236d2a9bdaaaad1d81092c208fc7e8f1b8e8ee2a335f"></a>

<a id="canonical-ba5f8422ecc63c5f1e7dec902d3ac0aae50badf62ee0e80513f6da207469fbd2"></a>

## san_type property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 046c2aa5cdda / 9

Type: `"string"`. Optional.

\[Enum: ISCSI|FC\] Block volume access protocol, either ISCSI or FC. Possible values are \`ISCSI\`,
\`FC\`.

Upstream description:

Block volume access protocol, either ISCSI or FC.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ISCSI",
    "FC"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ISCSI",
    "FC"
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
    "ves.io.schema.rules.string.in": "[\\\"ISCSI\\\",\\\"FC\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ISCSI\\\",\\\"FC\\\"]"
  }
}
```

<a id="canonical-24dfb2dab21f9fb787274e7e7f4a0343ca0b65ffb2e557279b39b5433a81b89d"></a>

## Next pages — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 046c2aa5cdda / 10

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](resources--fleet--reference--group-004.md#canonical-da7bfa1a7036624714eb7b92025dda342ba14a006026764de38b8f3c59fb569f)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--fleet--reference--group-004.md#canonical-558d5d8c784c45808aa7983b2b15b646d931426b9f4bb4de8b307ae1f82d078b)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-da7bfa1a7036624714eb7b92025dda342ba14a006026764de38b8f3c59fb569f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2afa3b3960bf456329fd534013b2c50cf85e9212ddebb843b9b072915a86a5ad"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 35008a2b2d78 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.pure_service_orchestrator](resources--fleet--reference--group-004.md#canonical-02b188fbd08df4634f53310c47e1459dbcbcbf07675d4d48c3343bdad2ba9044)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--fleet--reference--group-004.md#canonical-558d5d8c784c45808aa7983b2b15b646d931426b9f4bb4de8b307ae1f82d078b)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](resources--fleet--reference--group-004.md#canonical-f499462f23484c13046d8a0cacebbfdd683bc47434362ec63390322486bd5e09)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays

<a id="canonical-2ae04a259b087ce3ffc5ab742cc03e4122ca15269f83e17b9fd4ace9d327a2ff"></a>

Type: `"object"`. list nested block, Optional.

For FlashArrays you must set the 'mgmt\_endpoint' and 'api\_token'.

Upstream description:

For FlashArrays you must set the "mgmt\_endpoint" and "api\_token"

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("mgmt_dns_name",
    "mgmt_ip")}
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
flash_arrays {
  # Configure direct properties listed below.
}
```

<a id="canonical-4a6da84535d930dfa95ba9a8fc98df14fe4f7b893798170590c19d58e0de3f35"></a>

## Direct properties — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 35008a2b2d78 / 3

- [api_token](resources--fleet--reference--group-004.md#canonical-2409e3bafaf0cb3dd21475006f6c02b1c30ff70df9d4fff949b3c97d081c6923): complete subsection reference.

<a id="canonical-1a1a5758c855c812678cec1f3d6931336b6b871f368a7347b87b8182cd963606"></a>

<a id="canonical-c4fd1cf59ac8e1f411456d4c4fecf1243c61799f3c702c44743ab9ee1f30eea9"></a>

## labels property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 35008a2b2d78 / 4

Type: `["map", "string"]`. Optional.

Specifies labels optional, and can be any key-value pair for use with the PSO 'fleet' provisioner.

Upstream description:

The labels are optional, and can be any key-value pair for use with the PSO "fleet" provisioner.

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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-4392a988121b3833a2dbec03fd8a956ec60cb07f16d3ea73e7c200e7e8ec790c"></a>

<a id="canonical-f3a3371ca17783cb107fdab1c95ced092cfda34697d4728c400c709914b01999"></a>

## mgmt_dns_name property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 35008a2b2d78 / 5

Type: `"string"`. Optional.

Exclusive with \[mgmt\_ip\] Management Endpoint's IP address is discovered using DNS name
resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[mgmt\_ip\] Management Endpoint's IP address is discovered using DNS name
resolution. The name given here is fully qualified domain name.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-bd524ada8fbb2e631b733663cf8571f37c4c57456e00b9a8aa6b7cf639a31cc3"></a>

<a id="canonical-5d60e7e057523965bdf74c1d9a3a65c2eb9ac327025500a9b9b02fd6e5aeac4a"></a>

## mgmt_ip property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 35008a2b2d78 / 6

Type: `"string"`. Optional.

Exclusive with \[mgmt\_dns\_name\] Management Endpoint is reachable at the given IP address.

Upstream description:

Exclusive with \[mgmt\_dns\_name\] Management Endpoint is reachable at the given IP address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-ac35f62a74659d80affe56a6c7c4a67ebdd3840857e6c6cb3d99ddf76194cd1c"></a>

## Next pages — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 35008a2b2d78 / 7

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](resources--fleet--reference--group-004.md#canonical-2409e3bafaf0cb3dd21475006f6c02b1c30ff70df9d4fff949b3c97d081c6923)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](resources--fleet--reference--group-004.md#canonical-f499462f23484c13046d8a0cacebbfdd683bc47434362ec63390322486bd5e09)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-2409e3bafaf0cb3dd21475006f6c02b1c30ff70df9d4fff949b3c97d081c6923"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-40bdda6959b2b8c9369493e2939e9894881a2f50f553047eec345e7bdf7cebc5"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 2e6270fa8787 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.pure_service_orchestrator](resources--fleet--reference--group-004.md#canonical-02b188fbd08df4634f53310c47e1459dbcbcbf07675d4d48c3343bdad2ba9044)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--fleet--reference--group-004.md#canonical-558d5d8c784c45808aa7983b2b15b646d931426b9f4bb4de8b307ae1f82d078b)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](resources--fleet--reference--group-004.md#canonical-f499462f23484c13046d8a0cacebbfdd683bc47434362ec63390322486bd5e09)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](resources--fleet--reference--group-004.md#canonical-da7bfa1a7036624714eb7b92025dda342ba14a006026764de38b8f3c59fb569f)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token

<a id="canonical-88d0198acab51deb0930663b32b7a8b520161eb24c830143b15b32c5aa21874d"></a>

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
api_token {
  # Configure direct properties listed below.
}
```

<a id="canonical-6751b1bdba466c94d0d4892e63d78ad5e536c7e2d582f16b8679434a90710a89"></a>

## Direct properties — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 2e6270fa8787 / 3

- [blindfold_secret_info](resources--fleet--reference--group-004.md#canonical-fad008caea033a5cf6726eb184ba9f601b3b66c2618614587245bf8e2afb1b6b): complete subsection reference.

- [clear_secret_info](resources--fleet--reference--group-004.md#canonical-2a44167344c035596b5d38446eef5794999fc1f9be23d9d1943bf65400d30e7b): complete subsection reference.

<a id="canonical-a566ecd14f7378f5bc93ab3b9e804cb2be052fb536e5a4be17366a8ba732647b"></a>

## Next pages — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 2e6270fa8787 / 4

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info](resources--fleet--reference--group-004.md#canonical-fad008caea033a5cf6726eb184ba9f601b3b66c2618614587245bf8e2afb1b6b)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info](resources--fleet--reference--group-004.md#canonical-2a44167344c035596b5d38446eef5794999fc1f9be23d9d1943bf65400d30e7b)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](resources--fleet--reference--group-004.md#canonical-da7bfa1a7036624714eb7b92025dda342ba14a006026764de38b8f3c59fb569f)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-fad008caea033a5cf6726eb184ba9f601b3b66c2618614587245bf8e2afb1b6b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-84b58d2f0bd3e22c189832005a2b53955a71135e44844b7cbf122eb09c06843b"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 65cb11b895b3 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.pure_service_orchestrator](resources--fleet--reference--group-004.md#canonical-02b188fbd08df4634f53310c47e1459dbcbcbf07675d4d48c3343bdad2ba9044)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--fleet--reference--group-004.md#canonical-558d5d8c784c45808aa7983b2b15b646d931426b9f4bb4de8b307ae1f82d078b)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](resources--fleet--reference--group-004.md#canonical-f499462f23484c13046d8a0cacebbfdd683bc47434362ec63390322486bd5e09)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](resources--fleet--reference--group-004.md#canonical-da7bfa1a7036624714eb7b92025dda342ba14a006026764de38b8f3c59fb569f)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](resources--fleet--reference--group-004.md#canonical-2409e3bafaf0cb3dd21475006f6c02b1c30ff70df9d4fff949b3c97d081c6923)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info

<a id="canonical-67ca5ec1303c0309b7e51292cb41a84d2713e45b07e07aa24930daefe140fbe3"></a>

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

<a id="canonical-360324b8568d8e9bca5fe7158cb657313f594eb6b9647f63cb908be55ef3c9ca"></a>

## Direct properties — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 65cb11b895b3 / 3

<a id="canonical-c46d2b01b5d9aa5efcef68b767fda4d53320df3607a45d31eec933ef93154a86"></a>

<a id="canonical-400e8df82ce4540addf1aa809961a0cf11e71ad32652f741e22e3dbd52264ed6"></a>

## decryption_provider property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 65cb11b895b3 / 4

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

<a id="canonical-03ce3af0153bddcb55d7a1269438f9ed0cabf2b9e4a914a904b7b1c13f701663"></a>

<a id="canonical-fa43887952874a2b9cfd133018b689bec0971ea6006a574da71e4113f774f706"></a>

## location property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 65cb11b895b3 / 5

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

<a id="canonical-8a494978431ce8ac93b754bdf2f64788fdfb903c2af5c22923b33fbd576c31cb"></a>

<a id="canonical-fd7fdfcd270951ee7aaa6859c34b12977a022a20d3fd165c7a4c897540dfed91"></a>

## store_provider property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 65cb11b895b3 / 6

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

<a id="canonical-42d7ce80318745892339c8f82df8a662bd3a8f32a2cabf8f28cf5bd311b0718c"></a>

## Next pages — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 65cb11b895b3 / 7

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](resources--fleet--reference--group-004.md#canonical-2409e3bafaf0cb3dd21475006f6c02b1c30ff70df9d4fff949b3c97d081c6923)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-2a44167344c035596b5d38446eef5794999fc1f9be23d9d1943bf65400d30e7b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dbf39c03402ef435d49d6dcbb663365634a2e87777cb91f753a046027bafe45b"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / bcc2fcf37d8c / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.pure_service_orchestrator](resources--fleet--reference--group-004.md#canonical-02b188fbd08df4634f53310c47e1459dbcbcbf07675d4d48c3343bdad2ba9044)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--fleet--reference--group-004.md#canonical-558d5d8c784c45808aa7983b2b15b646d931426b9f4bb4de8b307ae1f82d078b)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](resources--fleet--reference--group-004.md#canonical-f499462f23484c13046d8a0cacebbfdd683bc47434362ec63390322486bd5e09)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](resources--fleet--reference--group-004.md#canonical-da7bfa1a7036624714eb7b92025dda342ba14a006026764de38b8f3c59fb569f)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](resources--fleet--reference--group-004.md#canonical-2409e3bafaf0cb3dd21475006f6c02b1c30ff70df9d4fff949b3c97d081c6923)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info

<a id="canonical-32c71a2d8272cf6c3eb3c8ec7595567f624708a72419c458c68b86c57a3c3ca6"></a>

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

<a id="canonical-641a419661c5308c3d3506e88859c932fb183eeb3dca27e2d2cd1bed17777d46"></a>

## Direct properties — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / bcc2fcf37d8c / 3

<a id="canonical-60d543adfe2d3d23a831cbd58ec63b0436f4d58d8c41fb24bd38850683364248"></a>

<a id="canonical-1daaf8287fb1880108a41bff5a9c50242d8474621e31b9fbd0f8f828c0c1d843"></a>

## provider_ref property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / bcc2fcf37d8c / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-9ec89b0dd37254ed8f560f7aa0429725bfe072c569d4fa05be134d55da721050"></a>

<a id="canonical-6709ff8f962967f56a73bdc3af42e1ff7ce4eba905df3a1bb7d82a285d06d301"></a>

## url property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / bcc2fcf37d8c / 5

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

<a id="canonical-c8ac9128f41c9d6c9b03a0816c20e9523a5720fc447b689e65f97cfcb34bca9a"></a>

## Next pages — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / bcc2fcf37d8c / 6

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](resources--fleet--reference--group-004.md#canonical-2409e3bafaf0cb3dd21475006f6c02b1c30ff70df9d4fff949b3c97d081c6923)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-34cb05c8510e2ca5ded4b76162fe66699380b39e32176756ed229c5362ee62b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e33aed7a217a895cebfc5e4a60d16bc904146c8d09e7e73c088eee0904ae9095"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / ad6f68dc9b81 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.pure_service_orchestrator](resources--fleet--reference--group-004.md#canonical-02b188fbd08df4634f53310c47e1459dbcbcbf07675d4d48c3343bdad2ba9044)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--fleet--reference--group-004.md#canonical-558d5d8c784c45808aa7983b2b15b646d931426b9f4bb4de8b307ae1f82d078b)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade

<a id="canonical-0b6b81ac3c773a1bd3093f659fb7acb1d467be05be7a26c8cfc2b4227feb4031"></a>

Type: `"object"`. single nested block, Optional.

Specify what storage flash blades should be managed the plugin.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("flash_blades")}
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
flash_blade {
  # Configure direct properties listed below.
}
```

<a id="canonical-7e790ac9466d915be4188b6ff4e2419bf9618870cb9d40f20b1d37dffa1a16ee"></a>

## Direct properties — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / ad6f68dc9b81 / 3

<a id="canonical-07e1125b38da5a005eca81559939c4346313a46fad122cdd475f006ed1982625"></a>

<a id="canonical-9650741002d2022f4059c0c630f0302a5cb9e0ab9d9389a930a94ea2c8e78187"></a>

## enable_snapshot_directory property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / ad6f68dc9b81 / 4

Type: `"bool"`. Optional.

Enable Snapshot Directory. Enable/Disable FlashBlade snapshots.

Upstream description:

Enable/Disable FlashBlade snapshots.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-88a54b36c8d0ba51d9b3541429721aa762d9ef78aae5b45c4ef9b2fca6811733"></a>

<a id="canonical-8f7960ea5bb4f8d1f5ba7b8f84282b658b4a2851a1d45d1f4ad6fff90fceb95a"></a>

## export_rules property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / ad6f68dc9b81 / 5

Type: `"string"`. Optional.

NFS Export Rules. NFS Export rules.

Upstream description:

NFS Export rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 250),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 250,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 250,
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
    "ves.io.schema.rules.string.max_len": "250",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "250",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [flash_blades](resources--fleet--reference--group-004.md#canonical-abbb2d4a491a28fd83653ff392e53a995101cb91ee6eed682fda822a29cbf3fb): complete subsection reference.

<a id="canonical-6ae1e3ba0aff4e87d241bfb25de1a63850e456da7931953799c57e3ea19d0998"></a>

## Next pages — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / ad6f68dc9b81 / 6

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](resources--fleet--reference--group-004.md#canonical-abbb2d4a491a28fd83653ff392e53a995101cb91ee6eed682fda822a29cbf3fb)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--fleet--reference--group-004.md#canonical-558d5d8c784c45808aa7983b2b15b646d931426b9f4bb4de8b307ae1f82d078b)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-abbb2d4a491a28fd83653ff392e53a995101cb91ee6eed682fda822a29cbf3fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f28bfc3bcd6e177276ef1d198558886c88f067a8cd231b3ef8041cb4fb67b03"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / 8754b7a3a078 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.pure_service_orchestrator](resources--fleet--reference--group-004.md#canonical-02b188fbd08df4634f53310c47e1459dbcbcbf07675d4d48c3343bdad2ba9044)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--fleet--reference--group-004.md#canonical-558d5d8c784c45808aa7983b2b15b646d931426b9f4bb4de8b307ae1f82d078b)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](resources--fleet--reference--group-004.md#canonical-34cb05c8510e2ca5ded4b76162fe66699380b39e32176756ed229c5362ee62b1)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades

<a id="canonical-dc1ee5201b87c4b15b21a74a6bbdc1bfc1ecc67b3f92777ca85a1edbbd5596c1"></a>

Type: `"object"`. list nested block, Optional.

For FlashBlades you must set the 'mgmt\_endpoint', 'api\_token' and nfs\_endpoint.

Upstream description:

For FlashBlades you must set the "mgmt\_endpoint", "api\_token" and nfs\_endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("mgmt_dns_name",
    "mgmt_ip"),
  validators.ConflictingListObjectAttributes("nfs_endpoint_dns_name",
    "nfs_endpoint_ip")}
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
flash_blades {
  # Configure direct properties listed below.
}
```

<a id="canonical-8b0049606c8d25f03c4f862d1ff3394c8f71cd7dc81412abf706e18ba3f27e79"></a>

## Direct properties — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / 8754b7a3a078 / 3

- [api_token](resources--fleet--reference--group-004.md#canonical-c3c512102f72d5c7a7265cc62810d88ec99306cef4f8a8dd643ce6faa43b9e10): complete subsection reference.

<a id="canonical-b818b07ec925b2d3a284d42c32f91940cb9cac18c3e1f2a9dac008fb17ea72c3"></a>

<a id="canonical-4075ee61b19558cb9c017287969c5de3859964e2a3683fce50d55a48c40d54cd"></a>

## labels property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / 8754b7a3a078 / 4

Type: `["map", "string"]`. Optional.

Specifies labels optional, and can be any key-value pair for use with the PSO 'fleet' provisioner.

Upstream description:

The labels are optional, and can be any key-value pair for use with the PSO "fleet" provisioner.

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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-71ff5d40182896e202b6f1f46ef85d8ba38e28a363dceed52deb3b5fb2eac2d2"></a>

<a id="canonical-3f5243bedf8563feba9a7a6b477e647cddb5fddf4f8a71c25babbca97bd84fe7"></a>

## mgmt_dns_name property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / 8754b7a3a078 / 5

Type: `"string"`. Optional.

Exclusive with \[mgmt\_ip\] Management Endpoint's IP address is discovered using DNS name
resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[mgmt\_ip\] Management Endpoint's IP address is discovered using DNS name
resolution. The name given here is fully qualified domain name.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-5caaec6b44903965ec8dc87fbd2be3e5375f7f94955fa9c6c75135b8f8774e96"></a>

<a id="canonical-1404a0c2d63574d627934d0f0d65312f721ac38fddadb726e8ed49482fa5763f"></a>

## mgmt_ip property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / 8754b7a3a078 / 6

Type: `"string"`. Optional.

Exclusive with \[mgmt\_dns\_name\] Management Endpoint is reachable at the given IP address.

Upstream description:

Exclusive with \[mgmt\_dns\_name\] Management Endpoint is reachable at the given IP address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-94d9d31771c06151cb6857c6c87e83a685cd6786434fab3b62cf33b54db872dc"></a>

<a id="canonical-7a9926d41fe360229351abb6b36900a47ea7ff63ec2bdf2a269b465dd15621f7"></a>

## nfs_endpoint_dns_name property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / 8754b7a3a078 / 7

Type: `"string"`. Optional.

Exclusive with \[nfs\_endpoint\_ip\] Endpoint's IP address is discovered using DNS name resolution.
The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[nfs\_endpoint\_ip\] Endpoint's IP address is discovered using DNS name resolution.
The name given here is fully qualified domain name.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-e81417ce3b647aa783db0f6ef60d8781d37fc069828b6d6a1807b9eeef40cd4e"></a>

<a id="canonical-525ce0a6bae9868a43d57b5408e75b98cabea1b79af5f157ec6dc44017fdd2de"></a>

## nfs_endpoint_ip property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / 8754b7a3a078 / 8

Type: `"string"`. Optional.

Exclusive with \[nfs\_endpoint\_dns\_name\] Endpoint is reachable at the given IP address.

Upstream description:

Exclusive with \[nfs\_endpoint\_dns\_name\] Endpoint is reachable at the given IP address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-94956526fa28c0680bb4a90d0b0aa3504f11012fbcea6e4047fe2194c9820952"></a>

## Next pages — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / 8754b7a3a078 / 9

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](resources--fleet--reference--group-004.md#canonical-c3c512102f72d5c7a7265cc62810d88ec99306cef4f8a8dd643ce6faa43b9e10)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](resources--fleet--reference--group-004.md#canonical-34cb05c8510e2ca5ded4b76162fe66699380b39e32176756ed229c5362ee62b1)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-c3c512102f72d5c7a7265cc62810d88ec99306cef4f8a8dd643ce6faa43b9e10"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e1992dc1130c3c281b62bb7be88562441115f869f2298049886ff76a34381b58"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / b24fc6692040 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.pure_service_orchestrator](resources--fleet--reference--group-004.md#canonical-02b188fbd08df4634f53310c47e1459dbcbcbf07675d4d48c3343bdad2ba9044)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--fleet--reference--group-004.md#canonical-558d5d8c784c45808aa7983b2b15b646d931426b9f4bb4de8b307ae1f82d078b)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](resources--fleet--reference--group-004.md#canonical-34cb05c8510e2ca5ded4b76162fe66699380b39e32176756ed229c5362ee62b1)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](resources--fleet--reference--group-004.md#canonical-abbb2d4a491a28fd83653ff392e53a995101cb91ee6eed682fda822a29cbf3fb)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token

<a id="canonical-43591d5f0a5b8fd2150212a905ce02ee4f8e3c0796bceea275433445ba57c7fa"></a>

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
api_token {
  # Configure direct properties listed below.
}
```

<a id="canonical-615e42c03178d09945a85f66a780c84a4cc6274b196cfc38c5204dccdf2eea40"></a>

## Direct properties — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / b24fc6692040 / 3

- [blindfold_secret_info](resources--fleet--reference--group-004.md#canonical-b2fd9e2fc07082c5b05dd8787ab6d07dc52b138406faf8c74ce1909c510a30f3): complete subsection reference.

- [clear_secret_info](resources--fleet--reference--group-004.md#canonical-841560326ea5e082f653cbebd9f63603c5f91d82121f27abc50c0577d63e918b): complete subsection reference.

<a id="canonical-722ee456c6aac3110f17f610268d85b531c08fb92da2357ae0ca4c3d43800a27"></a>

## Next pages — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / b24fc6692040 / 4

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info](resources--fleet--reference--group-004.md#canonical-b2fd9e2fc07082c5b05dd8787ab6d07dc52b138406faf8c74ce1909c510a30f3)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info](resources--fleet--reference--group-004.md#canonical-841560326ea5e082f653cbebd9f63603c5f91d82121f27abc50c0577d63e918b)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](resources--fleet--reference--group-004.md#canonical-abbb2d4a491a28fd83653ff392e53a995101cb91ee6eed682fda822a29cbf3fb)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-b2fd9e2fc07082c5b05dd8787ab6d07dc52b138406faf8c74ce1909c510a30f3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-84974554c1d3aeab85ef6aaa116bf8f310ab1b6f26252c993349ab0b48642bf7"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / 331378023702 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.pure_service_orchestrator](resources--fleet--reference--group-004.md#canonical-02b188fbd08df4634f53310c47e1459dbcbcbf07675d4d48c3343bdad2ba9044)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--fleet--reference--group-004.md#canonical-558d5d8c784c45808aa7983b2b15b646d931426b9f4bb4de8b307ae1f82d078b)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](resources--fleet--reference--group-004.md#canonical-34cb05c8510e2ca5ded4b76162fe66699380b39e32176756ed229c5362ee62b1)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](resources--fleet--reference--group-004.md#canonical-abbb2d4a491a28fd83653ff392e53a995101cb91ee6eed682fda822a29cbf3fb)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](resources--fleet--reference--group-004.md#canonical-c3c512102f72d5c7a7265cc62810d88ec99306cef4f8a8dd643ce6faa43b9e10)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info

<a id="canonical-25c12c98320818f0e4a74d1f830ccd2b4df08bd164723de0d2510aa49fc2c143"></a>

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

<a id="canonical-e269564fb379a45ac9ab9c5f91caf9de3182ddab3a5bf3f4c444620eb1f1d103"></a>

## Direct properties — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / 331378023702 / 3

<a id="canonical-ff58e74acade22a85f3fba73affaebced132f2c6915c706a0916bdb8050ab93b"></a>

<a id="canonical-f7296bb1b4aeed23ef329055e191696b7613bb8604c66303c3713e0e48c6fae9"></a>

## decryption_provider property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / 331378023702 / 4

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

<a id="canonical-f844d7f75552348e93c8557243c25d0c5096a0f5abcd4a2607ebc804e9526d0d"></a>

<a id="canonical-ba09b7eb5041ac76cc29ae11d8d0a9e7f5877045747d7eed2ded3c7c92a74c09"></a>

## location property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / 331378023702 / 5

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

<a id="canonical-54cd2031cb3fe10681c956a99c8b116c139faa65b7fd9f66d7acfc1095dccffe"></a>

<a id="canonical-e7e3fd5ee00c9f6414e116944b432f05ff26cb535110197fb8d9d8fddb540ef4"></a>

## store_provider property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / 331378023702 / 6

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

<a id="canonical-a21248cb9a49a498420210db9fe4af50ac7341bb797cc471718ecf7b49e14c21"></a>

## Next pages — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / 331378023702 / 7

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](resources--fleet--reference--group-004.md#canonical-c3c512102f72d5c7a7265cc62810d88ec99306cef4f8a8dd643ce6faa43b9e10)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-841560326ea5e082f653cbebd9f63603c5f91d82121f27abc50c0577d63e918b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9cb35f3d92c750bfcbf3b8eb6882587a04fdc1e013882e81414f1e5fa5e5b81"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / fdcad6637415 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.pure_service_orchestrator](resources--fleet--reference--group-004.md#canonical-02b188fbd08df4634f53310c47e1459dbcbcbf07675d4d48c3343bdad2ba9044)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--fleet--reference--group-004.md#canonical-558d5d8c784c45808aa7983b2b15b646d931426b9f4bb4de8b307ae1f82d078b)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](resources--fleet--reference--group-004.md#canonical-34cb05c8510e2ca5ded4b76162fe66699380b39e32176756ed229c5362ee62b1)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](resources--fleet--reference--group-004.md#canonical-abbb2d4a491a28fd83653ff392e53a995101cb91ee6eed682fda822a29cbf3fb)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](resources--fleet--reference--group-004.md#canonical-c3c512102f72d5c7a7265cc62810d88ec99306cef4f8a8dd643ce6faa43b9e10)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info

<a id="canonical-87d9bfaa4310edd73355f580cb95d67b2569c005c026795ff866289684c99e5e"></a>

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

<a id="canonical-b2226fad5da30042dddfba9ac78713c5be42dbb1a356d9ebaad75cf5ccff57fe"></a>

## Direct properties — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / fdcad6637415 / 3

<a id="canonical-c66d1d1e30d1679d1cd929683b47f392df0217ac11126e045d8bffb640fb330a"></a>

<a id="canonical-df857e80ca22f38ce85627871ad01549da57ae864beae91ec981753b39051234"></a>

## provider_ref property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / fdcad6637415 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-4cd956c174d3ef413a45a4f4830d9fb0bd92d0c76d536bcd749b69941507339a"></a>

<a id="canonical-3082036332eff4eeb830fd6df83d154d9f9dd2d26e5afcce0ffce8b5d606b088"></a>

## url property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / fdcad6637415 / 5

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

<a id="canonical-26bee9c7acc44fb62bd57f0df497cd074e012e23f8fa81a319f23eb22237e323"></a>

## Next pages — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / fdcad6637415 / 6

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](resources--fleet--reference--group-004.md#canonical-c3c512102f72d5c7a7265cc62810d88ec99306cef4f8a8dd643ce6faa43b9e10)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-a0ad9e77e25f172cd7ebf3e39a25f59ef2e0197cf555254fa1eae93c1269011e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-538ce0adbb838c109886946ce0d47d71ce5c3e261b322e49469b1c562fa4a217"></a>

## storage_interface_list — storage_interface_list / 56c8177c1440 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- storage_interface_list

<a id="canonical-abde75c3848a86f8fac961cec1ea5eb4baf07f3d441f29cca9d305a37ab7c837"></a>

Type: `"object"`. single nested block, Optional.

Add all interfaces belonging to this fleet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("interfaces")}
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
storage_interface_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-d3b2f8581361f1b2fdc0b719c4134aeb0d818d0a08edf7373a26e695f8f5967c"></a>

## Direct properties — storage_interface_list / 56c8177c1440 / 3

- [interfaces](resources--fleet--reference--group-004.md#canonical-24d1d677711fede012d8a7c6e54726a25e86e007a0ce27766f4cc05f263f7d9f): complete subsection reference.

<a id="canonical-e1323f5ded79fa7ad6cd1ce2636750b29de9d3c640c0df685ad6fcefef3bc4b5"></a>

## Next pages — storage_interface_list / 56c8177c1440 / 4

- [storage_interface_list.interfaces](resources--fleet--reference--group-004.md#canonical-24d1d677711fede012d8a7c6e54726a25e86e007a0ce27766f4cc05f263f7d9f)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-24d1d677711fede012d8a7c6e54726a25e86e007a0ce27766f4cc05f263f7d9f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d51ddb84a03b363e6a9945d2924c33e9725b653227334fa5cf0b1ec10481ea83"></a>

## storage_interface_list.interfaces — storage_interface_list.interfaces / 071fdaf11fad / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_interface_list](resources--fleet--reference--group-004.md#canonical-a0ad9e77e25f172cd7ebf3e39a25f59ef2e0197cf555254fa1eae93c1269011e)
- storage_interface_list.interfaces

<a id="canonical-eacb7474009f44727f047a6e40e448e0acfcd807f81fadad0d309364ddbe3ddd"></a>

Type: `"object"`. list nested block, Optional.

Add all interfaces belonging to this fleet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
interfaces {
  # Configure direct properties listed below.
}
```

<a id="canonical-83df24cefdd388e24e61cadc65bf12d5de8910a6d6b04a0a63d7b152ed00781f"></a>

## Direct properties — storage_interface_list.interfaces / 071fdaf11fad / 3

<a id="canonical-0113450160775584addd16ae3a9900fadbe0b8dbf12a8727cfeb5e22d2ba0767"></a>

<a id="canonical-2900d4c4ce598e8bd8332c368b651855c8887809feff5a2c47749f0a54e005e0"></a>

## name property — storage_interface_list.interfaces / 071fdaf11fad / 4

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

<a id="canonical-fddbda06f3d00de103ea8c6f670aafe40b2cb2459ca0621efbd145f5c6576437"></a>

<a id="canonical-5d67ebfc42e1752dfa95d5733b99821a8ad7214e72255ca8fa8e9e9dce0b93ae"></a>

## namespace property — storage_interface_list.interfaces / 071fdaf11fad / 5

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

<a id="canonical-56de481721b6336735ed470575f266eb8b527007ea2d25ddcadaca9f03a11f11"></a>

<a id="canonical-a6af77c1b92ab14b85167870fa2d9eb09c6423968a08130fad293557e0add63b"></a>

## tenant property — storage_interface_list.interfaces / 071fdaf11fad / 6

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

<a id="canonical-2167ef53f05bf63152b075770ad30af5d160f0ad82f69fcb0d9a7fc48ee05d21"></a>

## Next pages — storage_interface_list.interfaces / 071fdaf11fad / 7

- [storage_interface_list](resources--fleet--reference--group-004.md#canonical-a0ad9e77e25f172cd7ebf3e39a25f59ef2e0197cf555254fa1eae93c1269011e)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-f44c11471fbbf1faacf11875c91e3b194ee4c8c8c365df7691d83e66869443bc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-773b00fb6099fdb62668996e73c2ef95964e6282133fccc5b466642ddce83f28"></a>

## storage_static_routes — storage_static_routes / 97418c7d1067 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- storage_static_routes

<a id="canonical-f0a29d64ab6d2465e8f1351c8716c9d7320d650e2542ff5d009cfdd3842125a5"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for storage static routes.

Upstream description:

List of storage static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("storage_routes")}
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
storage_static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-3539f714af4e1512fb6cda58d2bedba233ecde14dd203d39f4d175727aef66fa"></a>

## Direct properties — storage_static_routes / 97418c7d1067 / 3

- [storage_routes](resources--fleet--reference--group-004.md#canonical-9351405b406acacaa2f89b680f143dd77af8608118f28e55c979043eac71a36d): complete subsection reference.

<a id="canonical-c5b376e106313f9d543506f323c5d34e3c3bc90b37a163c9d940c638e1bf9736"></a>

## Next pages — storage_static_routes / 97418c7d1067 / 4

- [storage_static_routes.storage_routes](resources--fleet--reference--group-004.md#canonical-9351405b406acacaa2f89b680f143dd77af8608118f28e55c979043eac71a36d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-9351405b406acacaa2f89b680f143dd77af8608118f28e55c979043eac71a36d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-28d6ee7591eebf2c531c38b0946a46070427cde1cb12b374d2e71f181b99a6ed"></a>

## storage_static_routes.storage_routes — storage_static_routes.storage_routes / d2fc5c3f7f96 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_static_routes](resources--fleet--reference--group-004.md#canonical-f44c11471fbbf1faacf11875c91e3b194ee4c8c8c365df7691d83e66869443bc)
- storage_static_routes.storage_routes

<a id="canonical-574e362c139e5cae91672eee5a566dedb5eb580d59c5c76f5c7f2e9ff5cad6a4"></a>

Type: `"object"`. list nested block, Optional.

List of Static Routes. List of storage static routes.

Upstream description:

List of storage static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("subnets")}
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
storage_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-880f3c7eea2b4e8418d35d5feedee39bda3dbe3f402ad54e0fc41b6cc4653277"></a>

## Direct properties — storage_static_routes.storage_routes / d2fc5c3f7f96 / 3

<a id="canonical-8f731df96ea3edd9f55a6370a5ecdf2b6024385df47f0af498853b78ca356d03"></a>

<a id="canonical-6f83a6d6a01a4e12f934866a493e4c2786f1765527dd60dc2a09ad038d4f9a12"></a>

## attrs property — storage_static_routes.storage_routes / d2fc5c3f7f96 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of route attributes associated with the static route.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
```

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

- [labels](resources--fleet--reference--group-004.md#canonical-da7256ba71b495ab4665d52326632c5339f31e8c7acd6e76ae510a1de6f196d4): complete subsection reference.

- [nexthop](resources--fleet--reference--group-004.md#canonical-bb48477d2a29bcdd01e184acc2b5c8693819b534959739fe91ab237679c75867): complete subsection reference.

- [subnets](resources--fleet--reference--group-004.md#canonical-0537b4bf9f1ce278cb306f5091283296f27e7095a5ba95b33fea735d5d45b2aa): complete subsection reference.

<a id="canonical-dd03985bdcff36154a623edf4beee0c7f219ba982c47feb274d33d50cd2932fb"></a>

## Next pages — storage_static_routes.storage_routes / d2fc5c3f7f96 / 5

- [storage_static_routes.storage_routes.labels](resources--fleet--reference--group-004.md#canonical-da7256ba71b495ab4665d52326632c5339f31e8c7acd6e76ae510a1de6f196d4)
- [storage_static_routes.storage_routes.nexthop](resources--fleet--reference--group-004.md#canonical-bb48477d2a29bcdd01e184acc2b5c8693819b534959739fe91ab237679c75867)
- [storage_static_routes.storage_routes.subnets](resources--fleet--reference--group-004.md#canonical-0537b4bf9f1ce278cb306f5091283296f27e7095a5ba95b33fea735d5d45b2aa)
- [storage_static_routes](resources--fleet--reference--group-004.md#canonical-f44c11471fbbf1faacf11875c91e3b194ee4c8c8c365df7691d83e66869443bc)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-da7256ba71b495ab4665d52326632c5339f31e8c7acd6e76ae510a1de6f196d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a4c7fe4f54525d7a3e24d210f75cf2c79a1d388d6ea1501691727682db6b2fcd"></a>

## storage_static_routes.storage_routes.labels — storage_static_routes.storage_routes.labels / 7d8e7cde2b42 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_static_routes](resources--fleet--reference--group-004.md#canonical-f44c11471fbbf1faacf11875c91e3b194ee4c8c8c365df7691d83e66869443bc)
- [storage_static_routes.storage_routes](resources--fleet--reference--group-004.md#canonical-9351405b406acacaa2f89b680f143dd77af8608118f28e55c979043eac71a36d)
- storage_static_routes.storage_routes.labels

<a id="canonical-e083c5ef23133ab0639f783b0284383b92c9ce1de006ff08b3ecf6228e1c326e"></a>

Type: `"object"`. single nested block, Optional.

Add Labels for this Static Route, these labels can be used in network policy.

Receipt-pinned upstream constraints:

```json
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
labels {}
```

<a id="canonical-ca642cd337d3d56f56830b3b33705ed20963a507baf3839e71e9b49ce230073e"></a>

## Direct properties — storage_static_routes.storage_routes.labels / 7d8e7cde2b42 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-06f50219a8988b72745a6fe126ad14a803703d19f91478f031b1bfd28f09b971"></a>

## Next pages — storage_static_routes.storage_routes.labels / 7d8e7cde2b42 / 4

- [storage_static_routes.storage_routes](resources--fleet--reference--group-004.md#canonical-9351405b406acacaa2f89b680f143dd77af8608118f28e55c979043eac71a36d)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-bb48477d2a29bcdd01e184acc2b5c8693819b534959739fe91ab237679c75867"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c7c0110aa924aada03eff7267dd724415837f5aac01a4c4be902e8594c1e8c79"></a>

## storage_static_routes.storage_routes.nexthop — storage_static_routes.storage_routes.nexthop / f5a043aec1e6 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_static_routes](resources--fleet--reference--group-004.md#canonical-f44c11471fbbf1faacf11875c91e3b194ee4c8c8c365df7691d83e66869443bc)
- [storage_static_routes.storage_routes](resources--fleet--reference--group-004.md#canonical-9351405b406acacaa2f89b680f143dd77af8608118f28e55c979043eac71a36d)
- storage_static_routes.storage_routes.nexthop

<a id="canonical-5f8f4abe2132ceb6c4c32fa5cc239550409fdb6fe75dc72ea78824997025665c"></a>

Type: `"object"`. single nested block, Optional.

Nexthop. Identifies the next-hop for a route.

Upstream description:

Identifies the next-hop for a route.

Receipt-pinned upstream constraints:

```json
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
nexthop {
  # Configure direct properties listed below.
}
```

<a id="canonical-7380c1f3d039d035c8826e5f95b3234bfb30b356159ac17d871258137adf8a1f"></a>

## Direct properties — storage_static_routes.storage_routes.nexthop / f5a043aec1e6 / 3

- [interface](resources--fleet--reference--group-004.md#canonical-0b63bc081cae29664c27801126a2cc379c5b8f6896b2ba76cd06c09f79392d38): complete subsection reference.

- [nexthop_address](resources--fleet--reference--group-004.md#canonical-17c1317c89a405ac988cfe020037bf8cfce127e4fd9890c7de58c1991496f28d): complete subsection reference.

<a id="canonical-f61e8a92807082d0682a5525000d37c0f93a00809ae7093b9624b9f736ca3dbe"></a>

<a id="canonical-353aad2e41f5a8c449dfc1806dd29940d94a043a7b47f278010d3f837782d094"></a>

## type property — storage_static_routes.storage_routes.nexthop / f5a043aec1e6 / 4

Type: `"string"`. Optional.

\[Enum: NEXT\_HOP\_DEFAULT\_GATEWAY|NEXT\_HOP\_USE\_CONFIGURED|NEXT\_HOP\_NETWORK\_INTERFACE\]
Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes
there is only one local interface on the virtual network. Use the specified address as nexthop Use
the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN..
Possible values are \`NEXT\_HOP\_DEFAULT\_GATEWAY\`, \`NEXT\_HOP\_USE\_CONFIGURED\`,
\`NEXT\_HOP\_NETWORK\_INTERFACE\`. Defaults to \`NEXT\_HOP\_DEFAULT\_GATEWAY\`.

Upstream description:

Defines types of next-hop

Use default gateway on the local interface as gateway for route. Assumes there is only one local
interface on the virtual network. Use the specified address as nexthop Use the network interface as
nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN private virtual network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NEXT_HOP_DEFAULT_GATEWAY",
  "enum": [
    "NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-822ff9f919e9c13c00be93dd4658cf9c5a651fe55111841528772ba4efa087ea"></a>

## Next pages — storage_static_routes.storage_routes.nexthop / f5a043aec1e6 / 5

- [storage_static_routes.storage_routes.nexthop.interface](resources--fleet--reference--group-004.md#canonical-0b63bc081cae29664c27801126a2cc379c5b8f6896b2ba76cd06c09f79392d38)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](resources--fleet--reference--group-004.md#canonical-17c1317c89a405ac988cfe020037bf8cfce127e4fd9890c7de58c1991496f28d)
- [storage_static_routes.storage_routes](resources--fleet--reference--group-004.md#canonical-9351405b406acacaa2f89b680f143dd77af8608118f28e55c979043eac71a36d)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-0b63bc081cae29664c27801126a2cc379c5b8f6896b2ba76cd06c09f79392d38"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-840b5f58c8c89faa15b7f490c634c8c4db8ce3ee52404c8cc06317d3b43fe159"></a>

## storage_static_routes.storage_routes.nexthop.interface — storage_static_routes.storage_routes.nexthop.interface / 10135e7a8096 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_static_routes](resources--fleet--reference--group-004.md#canonical-f44c11471fbbf1faacf11875c91e3b194ee4c8c8c365df7691d83e66869443bc)
- [storage_static_routes.storage_routes](resources--fleet--reference--group-004.md#canonical-9351405b406acacaa2f89b680f143dd77af8608118f28e55c979043eac71a36d)
- [storage_static_routes.storage_routes.nexthop](resources--fleet--reference--group-004.md#canonical-bb48477d2a29bcdd01e184acc2b5c8693819b534959739fe91ab237679c75867)
- storage_static_routes.storage_routes.nexthop.interface

<a id="canonical-9168ce2a452993aba94d26ea38e706b62555f6fd226f3486dfb7675d24fcc3a0"></a>

Type: `"object"`. list nested block, Optional.

Nexthop is network interface when type is 'Network-Interface'.

Upstream description:

Nexthop is network interface when type is "Network-Interface"

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
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-2b93fd2df0d25fe59167c865c4c6f6ddf5db6e596ac6d9b6c7e8862ecd10052b"></a>

## Direct properties — storage_static_routes.storage_routes.nexthop.interface / 10135e7a8096 / 3

<a id="canonical-57f6d7abbc24c68d1d4586f36c029301014daf9f08cf5259c857dea60665ccb3"></a>

<a id="canonical-258dc34f17fcd9fb9ec7726f3adb1b97561708ca08700192f0d9735ca25e4c25"></a>

## kind property — storage_static_routes.storage_routes.nexthop.interface / 10135e7a8096 / 4

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

<a id="canonical-7fd9d1e5378c8680f14b86f11e3a123f2b97af10c22ce30ca06e53e7792de9eb"></a>

<a id="canonical-eb58b05ac4c935aedcd320663b76714f1d49243a31832a56f8d59862ae443dfe"></a>

## name property — storage_static_routes.storage_routes.nexthop.interface / 10135e7a8096 / 5

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

<a id="canonical-586530fc599c7ae47839e40a13749fbd0a282949aa3d281a8c634ca97a58bbbb"></a>

<a id="canonical-80133b439982e18ffae4b490227fa02ac3558e26a4db338bb473e42bccc5f8c5"></a>

## namespace property — storage_static_routes.storage_routes.nexthop.interface / 10135e7a8096 / 6

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

<a id="canonical-44a46b73d5930f28ec632e49fd156e273d857f2697d265b167ee54f274672208"></a>

<a id="canonical-15de2c74f98a82900b1238a8993dc323ff583cc73c4c0bfc43977ca6993aafd0"></a>

## tenant property — storage_static_routes.storage_routes.nexthop.interface / 10135e7a8096 / 7

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

<a id="canonical-47d4145268dd24a155c3fa722473475bad85f15d5246e7432994e3a6082c0db1"></a>

<a id="canonical-98b72c07709e16fa6c3205cec600f173f60b53d84e6c83aef6e2682c4c4e16f3"></a>

## uid property — storage_static_routes.storage_routes.nexthop.interface / 10135e7a8096 / 8

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

<a id="canonical-3f3e47d9e92586d015afdf3236e5e1f19e96ea624e1df5175dcf89b337cc8fc2"></a>

## Next pages — storage_static_routes.storage_routes.nexthop.interface / 10135e7a8096 / 9

- [storage_static_routes.storage_routes.nexthop](resources--fleet--reference--group-004.md#canonical-bb48477d2a29bcdd01e184acc2b5c8693819b534959739fe91ab237679c75867)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-17c1317c89a405ac988cfe020037bf8cfce127e4fd9890c7de58c1991496f28d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-37a8916c1c6a36e8dd00d725db7b75ad6a9ba0f6de0a151977688a30b976ab14"></a>

## storage_static_routes.storage_routes.nexthop.nexthop_address — storage_static_routes.storage_routes.nexthop.nexthop_address / f671370ed257 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_static_routes](resources--fleet--reference--group-004.md#canonical-f44c11471fbbf1faacf11875c91e3b194ee4c8c8c365df7691d83e66869443bc)
- [storage_static_routes.storage_routes](resources--fleet--reference--group-004.md#canonical-9351405b406acacaa2f89b680f143dd77af8608118f28e55c979043eac71a36d)
- [storage_static_routes.storage_routes.nexthop](resources--fleet--reference--group-004.md#canonical-bb48477d2a29bcdd01e184acc2b5c8693819b534959739fe91ab237679c75867)
- storage_static_routes.storage_routes.nexthop.nexthop_address

<a id="canonical-48a2b8a83c7486503e7ea76197f1136b20c66d1e5c2f49dd56089c2350078ad5"></a>

Type: `"object"`. single nested block, Optional.

IP Address used to specify an IPv4 or IPv6 address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dual_stack",
    "ipv4"),
  validators.ConflictingObjectAttributes("dual_stack",
    "ipv6"),
  validators.ConflictingObjectAttributes("ipv4",
    "ipv6")}
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
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

Terraform syntax:

```terraform
nexthop_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-45792546a2168d7f5b23c64be6e3f3ecae7eddf3d0851d6f6798af454eaedc28"></a>

## Direct properties — storage_static_routes.storage_routes.nexthop.nexthop_address / f671370ed257 / 3

- [dual_stack](resources--fleet--reference--group-004.md#canonical-d2676a37173125aa765b6a9aaa949517ff2ccb2c25fe067dead1726d416731dc): complete subsection reference.

- [ipv4](resources--fleet--reference--group-004.md#canonical-554cbb9cc6cbb78e487093141e13d5d589b5d89ad2026b85630bbb96bfd403a5): complete subsection reference.

- [ipv6](resources--fleet--reference--group-004.md#canonical-da392652c26ed067d1c5a8a09f2786ebc0fc702bc8ac0d03bdd62143a0414612): complete subsection reference.

<a id="canonical-6b0f594156c33450b81c9fb4ed24ec088a8befd7360a83cc45060a6d67ad1e5f"></a>

## Next pages — storage_static_routes.storage_routes.nexthop.nexthop_address / f671370ed257 / 4

- [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack](resources--fleet--reference--group-004.md#canonical-d2676a37173125aa765b6a9aaa949517ff2ccb2c25fe067dead1726d416731dc)
- [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4](resources--fleet--reference--group-004.md#canonical-554cbb9cc6cbb78e487093141e13d5d589b5d89ad2026b85630bbb96bfd403a5)
- [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6](resources--fleet--reference--group-004.md#canonical-da392652c26ed067d1c5a8a09f2786ebc0fc702bc8ac0d03bdd62143a0414612)
- [storage_static_routes.storage_routes.nexthop](resources--fleet--reference--group-004.md#canonical-bb48477d2a29bcdd01e184acc2b5c8693819b534959739fe91ab237679c75867)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-d2676a37173125aa765b6a9aaa949517ff2ccb2c25fe067dead1726d416731dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8133cf92597add234c0833f39e6a8b719ba802dfc9797e8515428af7476c7e87"></a>

## storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack — storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack / 50b7cce1f260 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_static_routes](resources--fleet--reference--group-004.md#canonical-f44c11471fbbf1faacf11875c91e3b194ee4c8c8c365df7691d83e66869443bc)
- [storage_static_routes.storage_routes](resources--fleet--reference--group-004.md#canonical-9351405b406acacaa2f89b680f143dd77af8608118f28e55c979043eac71a36d)
- [storage_static_routes.storage_routes.nexthop](resources--fleet--reference--group-004.md#canonical-bb48477d2a29bcdd01e184acc2b5c8693819b534959739fe91ab237679c75867)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](resources--fleet--reference--group-004.md#canonical-17c1317c89a405ac988cfe020037bf8cfce127e4fd9890c7de58c1991496f28d)
- storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack

<a id="canonical-f6680bd27239df405ecff16ebc8851a1f79effaa9892839d0631ce32077db43f"></a>

Type: `"object"`. single nested block, Optional.

DualStackAddressType represents both IPv4 and IPv6 together.

Receipt-pinned upstream constraints:

```json
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
dual_stack {
  # Configure direct properties listed below.
}
```

<a id="canonical-56bba2d269727dd306fc16532d3137c81d0d65aac46857fe63076493a1398088"></a>

## Direct properties — storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack / 50b7cce1f260 / 3

- [ipv4](resources--fleet--reference--group-004.md#canonical-fcb07d2993497c8408ea6bb9fa7f6931cd079baf867d8345713632862fffb44d): complete subsection reference.

- [ipv6](resources--fleet--reference--group-004.md#canonical-b94a5e0c2b0a817e6c9f44895d340006c1f76780804f758cc818ae6e517f9b66): complete subsection reference.

<a id="canonical-203020b6a8083458dbee951c0f51971365f43b1478ed87dc6bd29420ae7a671a"></a>

## Next pages — storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack / 50b7cce1f260 / 4

- [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4](resources--fleet--reference--group-004.md#canonical-fcb07d2993497c8408ea6bb9fa7f6931cd079baf867d8345713632862fffb44d)
- [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6](resources--fleet--reference--group-004.md#canonical-b94a5e0c2b0a817e6c9f44895d340006c1f76780804f758cc818ae6e517f9b66)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](resources--fleet--reference--group-004.md#canonical-17c1317c89a405ac988cfe020037bf8cfce127e4fd9890c7de58c1991496f28d)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-fcb07d2993497c8408ea6bb9fa7f6931cd079baf867d8345713632862fffb44d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-646a2eb263a3a48755c9627c9890e1aa66c4de750a3c3d7b1b213b0660c9105f"></a>

## storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4 — storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4 / b1c38f0c3130 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_static_routes](resources--fleet--reference--group-004.md#canonical-f44c11471fbbf1faacf11875c91e3b194ee4c8c8c365df7691d83e66869443bc)
- [storage_static_routes.storage_routes](resources--fleet--reference--group-004.md#canonical-9351405b406acacaa2f89b680f143dd77af8608118f28e55c979043eac71a36d)
- [storage_static_routes.storage_routes.nexthop](resources--fleet--reference--group-004.md#canonical-bb48477d2a29bcdd01e184acc2b5c8693819b534959739fe91ab237679c75867)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](resources--fleet--reference--group-004.md#canonical-17c1317c89a405ac988cfe020037bf8cfce127e4fd9890c7de58c1991496f28d)
- [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack](resources--fleet--reference--group-004.md#canonical-d2676a37173125aa765b6a9aaa949517ff2ccb2c25fe067dead1726d416731dc)
- storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4

<a id="canonical-e84d27fd2b2216a52deab656aeff38659b13ded354ee378f5544ddd45622b659"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

Receipt-pinned upstream constraints:

```json
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
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-22c9104728a702cf1f634f1b7278e6af91781a881b37578ba7ce6bcb0ab545a2"></a>

## Direct properties — storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4 / b1c38f0c3130 / 3

<a id="canonical-9d1ea7d5004e23947618b49a459415f207059fdfd9f48d92c3ff2c55f806a004"></a>

<a id="canonical-968f1316a53cc0c60ae0350dd8cab7f897e394b3d06633aa94ca70299e299b4f"></a>

## addr property — storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4 / b1c38f0c3130 / 4

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-e2c736c12c8ae99102d1814efddb2748a75b79c09248d285d0cc0e0618906869"></a>

## Next pages — storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4 / b1c38f0c3130 / 5

- [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack](resources--fleet--reference--group-004.md#canonical-d2676a37173125aa765b6a9aaa949517ff2ccb2c25fe067dead1726d416731dc)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-b94a5e0c2b0a817e6c9f44895d340006c1f76780804f758cc818ae6e517f9b66"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a4cc6e68b63b8db0efcadb21335c8bc93a42f6a84222e82836abaa3ee01121b"></a>

## storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6 — storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6 / ee42b2426233 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_static_routes](resources--fleet--reference--group-004.md#canonical-f44c11471fbbf1faacf11875c91e3b194ee4c8c8c365df7691d83e66869443bc)
- [storage_static_routes.storage_routes](resources--fleet--reference--group-004.md#canonical-9351405b406acacaa2f89b680f143dd77af8608118f28e55c979043eac71a36d)
- [storage_static_routes.storage_routes.nexthop](resources--fleet--reference--group-004.md#canonical-bb48477d2a29bcdd01e184acc2b5c8693819b534959739fe91ab237679c75867)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](resources--fleet--reference--group-004.md#canonical-17c1317c89a405ac988cfe020037bf8cfce127e4fd9890c7de58c1991496f28d)
- [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack](resources--fleet--reference--group-004.md#canonical-d2676a37173125aa765b6a9aaa949517ff2ccb2c25fe067dead1726d416731dc)
- storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6

<a id="canonical-b5ea068053ad5c215c8be716302392d062fd7b06e5af00aafa0375e54916ef76"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

Receipt-pinned upstream constraints:

```json
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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-834461ce8a1ffa5dbee50429f5fa06260b03d5786c16468d61b7337389f9359f"></a>

## Direct properties — storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6 / ee42b2426233 / 3

<a id="canonical-7057fc6efb84b3d5f8e0b4bbe32e5518c04ca7e4161facefb6f805925b771c95"></a>

<a id="canonical-1043548e0b26e3170f8b0e80662b142d88349ffa2cf4055bfe3b5b453a6ba03c"></a>

## addr property — storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6 / ee42b2426233 / 4

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-795c15fbbfb28149e85ed8335d65d16ac4eef89def87b2c5b437720e65059c5c"></a>

## Next pages — storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6 / ee42b2426233 / 5

- [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack](resources--fleet--reference--group-004.md#canonical-d2676a37173125aa765b6a9aaa949517ff2ccb2c25fe067dead1726d416731dc)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-554cbb9cc6cbb78e487093141e13d5d589b5d89ad2026b85630bbb96bfd403a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-84f0d4e421850f38403ef3979f21ab6a45b5be6aa7ae2ddd74a902f80cf2eda3"></a>

## storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4 — storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4 / 8497a1804f15 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_static_routes](resources--fleet--reference--group-004.md#canonical-f44c11471fbbf1faacf11875c91e3b194ee4c8c8c365df7691d83e66869443bc)
- [storage_static_routes.storage_routes](resources--fleet--reference--group-004.md#canonical-9351405b406acacaa2f89b680f143dd77af8608118f28e55c979043eac71a36d)
- [storage_static_routes.storage_routes.nexthop](resources--fleet--reference--group-004.md#canonical-bb48477d2a29bcdd01e184acc2b5c8693819b534959739fe91ab237679c75867)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](resources--fleet--reference--group-004.md#canonical-17c1317c89a405ac988cfe020037bf8cfce127e4fd9890c7de58c1991496f28d)
- storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4

<a id="canonical-9a3ed42f51407959a0ae52c84e3e2e51372122d9ad5dc14963f8226a24ce94c2"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

Receipt-pinned upstream constraints:

```json
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
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-d7d490f0344308410c2ed4f4f92f8ebeb22bb283e0766392662ffe30eb0df3fb"></a>

## Direct properties — storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4 / 8497a1804f15 / 3

<a id="canonical-46cf39ad419a749c978f1b9ec84a18f7e6c0b555636c395ca93d23b51c475d5d"></a>

<a id="canonical-130989bcd9ff480630cd4a31acfeb60f7d35451d1a6a8ba284291e797d295422"></a>

## addr property — storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4 / 8497a1804f15 / 4

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-caf8fbfa95e99d8aa28f1aeb3f16dc6d511d15db9d05529e032f8b7389730191"></a>

## Next pages — storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4 / 8497a1804f15 / 5

- [storage_static_routes.storage_routes.nexthop.nexthop_address](resources--fleet--reference--group-004.md#canonical-17c1317c89a405ac988cfe020037bf8cfce127e4fd9890c7de58c1991496f28d)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-da392652c26ed067d1c5a8a09f2786ebc0fc702bc8ac0d03bdd62143a0414612"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd4fc8f99ae8465d5438fb1e8f898de99af7d9b0edacf736794c605cb6ad20ad"></a>

## storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6 — storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6 / 771a25b66248 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_static_routes](resources--fleet--reference--group-004.md#canonical-f44c11471fbbf1faacf11875c91e3b194ee4c8c8c365df7691d83e66869443bc)
- [storage_static_routes.storage_routes](resources--fleet--reference--group-004.md#canonical-9351405b406acacaa2f89b680f143dd77af8608118f28e55c979043eac71a36d)
- [storage_static_routes.storage_routes.nexthop](resources--fleet--reference--group-004.md#canonical-bb48477d2a29bcdd01e184acc2b5c8693819b534959739fe91ab237679c75867)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](resources--fleet--reference--group-004.md#canonical-17c1317c89a405ac988cfe020037bf8cfce127e4fd9890c7de58c1991496f28d)
- storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6

<a id="canonical-b8fb6d028fb0d22719bc9f30e8c04fce5670093869de686e388d533b591bf92f"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

Receipt-pinned upstream constraints:

```json
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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-b47b1d59854b7ae5ea5f7ee133e658011032425f2257f5ef7de8cb5a75dd2a8b"></a>

## Direct properties — storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6 / 771a25b66248 / 3

<a id="canonical-3d9de131dffe24a5c324c8982433cf80afca137b604602de1fe844261768c04e"></a>

<a id="canonical-7c615fb995b2e3a06185c2f485c5e1147f6a1fab4fc1912e0d1b0712e8b8a661"></a>

## addr property — storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6 / 771a25b66248 / 4

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-5cf7b7efd475c1264e2ba3069dfea16a3d9a8a648fe569362605a1ad2074aa5a"></a>

## Next pages — storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6 / 771a25b66248 / 5

- [storage_static_routes.storage_routes.nexthop.nexthop_address](resources--fleet--reference--group-004.md#canonical-17c1317c89a405ac988cfe020037bf8cfce127e4fd9890c7de58c1991496f28d)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-0537b4bf9f1ce278cb306f5091283296f27e7095a5ba95b33fea735d5d45b2aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6212d3d49026898c93c753cd0e9fb0c28c25fb76378df084a7fb615bec2a6cd"></a>

## storage_static_routes.storage_routes.subnets — storage_static_routes.storage_routes.subnets / dc93834d021e / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_static_routes](resources--fleet--reference--group-004.md#canonical-f44c11471fbbf1faacf11875c91e3b194ee4c8c8c365df7691d83e66869443bc)
- [storage_static_routes.storage_routes](resources--fleet--reference--group-004.md#canonical-9351405b406acacaa2f89b680f143dd77af8608118f28e55c979043eac71a36d)
- storage_static_routes.storage_routes.subnets

<a id="canonical-947210d8f776f1203d39e5ff072ca03fde339a8899643dbacb20d54d4e032d86"></a>

Type: `"object"`. list nested block, Optional.

Subnets. List of route prefixes.

Upstream description:

List of route prefixes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("ipv4",
    "ipv6")}
```

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
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

Terraform syntax:

```terraform
subnets {
  # Configure direct properties listed below.
}
```

<a id="canonical-084d16ad5b4d41dc0c21974458d0d08ef146419e26b8c399802ab485533684aa"></a>

## Direct properties — storage_static_routes.storage_routes.subnets / dc93834d021e / 3

- [ipv4](resources--fleet--reference--group-004.md#canonical-9e9a58ac6002fa5fd673a82d79088f0bd3392e41147cdce9fd942e520ade0221): complete subsection reference.

- [ipv6](resources--fleet--reference--group-004.md#canonical-276aa29fa0a87470ee85cf627e167dce9e35713e1dae09fa5f7f28b05efa0132): complete subsection reference.

<a id="canonical-ac9458a5100a816b543d1dca02f903215d665963369eefbdae3b71917030ea6d"></a>

## Next pages — storage_static_routes.storage_routes.subnets / dc93834d021e / 4

- [storage_static_routes.storage_routes.subnets.ipv4](resources--fleet--reference--group-004.md#canonical-9e9a58ac6002fa5fd673a82d79088f0bd3392e41147cdce9fd942e520ade0221)
- [storage_static_routes.storage_routes.subnets.ipv6](resources--fleet--reference--group-004.md#canonical-276aa29fa0a87470ee85cf627e167dce9e35713e1dae09fa5f7f28b05efa0132)
- [storage_static_routes.storage_routes](resources--fleet--reference--group-004.md#canonical-9351405b406acacaa2f89b680f143dd77af8608118f28e55c979043eac71a36d)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-9e9a58ac6002fa5fd673a82d79088f0bd3392e41147cdce9fd942e520ade0221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5cf87c0a6dd832b18253c180d998d44bb05aa25b80d090474207deae0f88c273"></a>

## storage_static_routes.storage_routes.subnets.ipv4 — storage_static_routes.storage_routes.subnets.ipv4 / 55c5d8b165e9 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_static_routes](resources--fleet--reference--group-004.md#canonical-f44c11471fbbf1faacf11875c91e3b194ee4c8c8c365df7691d83e66869443bc)
- [storage_static_routes.storage_routes](resources--fleet--reference--group-004.md#canonical-9351405b406acacaa2f89b680f143dd77af8608118f28e55c979043eac71a36d)
- [storage_static_routes.storage_routes.subnets](resources--fleet--reference--group-004.md#canonical-0537b4bf9f1ce278cb306f5091283296f27e7095a5ba95b33fea735d5d45b2aa)
- storage_static_routes.storage_routes.subnets.ipv4

<a id="canonical-c5cbb38776c787e42ebd0d0955523528ba8ef1dba1d08a788770d65caf85b0ff"></a>

Type: `"object"`. single nested block, Optional.

IPv4 subnets specified as prefix and prefix-length. Prefix length must be &lt;= 32.

Receipt-pinned upstream constraints:

```json
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
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-d7fad1ed724ac02ee59f18a94f4606ad6bfaa1c3a0f38fdd08df3cdf234773e5"></a>

## Direct properties — storage_static_routes.storage_routes.subnets.ipv4 / 55c5d8b165e9 / 3

<a id="canonical-4cc866de173968d0a81bbc97659da2c73e8bedf19f0c9cb2101dbef7411c51c5"></a>

<a id="canonical-4aa740a047369c6b83f7f5421c2b788077a3bd59979b71c7fa222d4de45acbe5"></a>

## plen property — storage_static_routes.storage_routes.subnets.ipv4 / 55c5d8b165e9 / 4

Type: `"number"`. Optional.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-e786b9d6e94cc1718628dcc125c8799356f804b63d43b429ba1bc16b544d72dc"></a>

<a id="canonical-d09fb222367e001b568c74d3423317b4bbafb2bed20e8bb9ba00020b3a89f847"></a>

## prefix property — storage_static_routes.storage_routes.subnets.ipv4 / 55c5d8b165e9 / 5

Type: `"string"`. Optional.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-772dcfa62e1f17833756a80d28741fcf6e73f347c824f7a49da027f0eeb2da9c"></a>

## Next pages — storage_static_routes.storage_routes.subnets.ipv4 / 55c5d8b165e9 / 6

- [storage_static_routes.storage_routes.subnets](resources--fleet--reference--group-004.md#canonical-0537b4bf9f1ce278cb306f5091283296f27e7095a5ba95b33fea735d5d45b2aa)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-276aa29fa0a87470ee85cf627e167dce9e35713e1dae09fa5f7f28b05efa0132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-44886120bef2cc476618b001858f44854841ed8485222715f99258d26fe547a4"></a>

## storage_static_routes.storage_routes.subnets.ipv6 — storage_static_routes.storage_routes.subnets.ipv6 / 7878306aa399 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_static_routes](resources--fleet--reference--group-004.md#canonical-f44c11471fbbf1faacf11875c91e3b194ee4c8c8c365df7691d83e66869443bc)
- [storage_static_routes.storage_routes](resources--fleet--reference--group-004.md#canonical-9351405b406acacaa2f89b680f143dd77af8608118f28e55c979043eac71a36d)
- [storage_static_routes.storage_routes.subnets](resources--fleet--reference--group-004.md#canonical-0537b4bf9f1ce278cb306f5091283296f27e7095a5ba95b33fea735d5d45b2aa)
- storage_static_routes.storage_routes.subnets.ipv6

<a id="canonical-aa3d29551011a9bbbd026029854e5978496cf48c0c2020eb508cc0fba107c33d"></a>

Type: `"object"`. single nested block, Optional.

IPv6 subnets specified as prefix and prefix-length. Prefix-legnth must be &lt;= 128.

Receipt-pinned upstream constraints:

```json
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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-b3a7d9ca826892642ae0e59b72c0201dddc74ba0a4027a43924d8afe3d5a1b1c"></a>

## Direct properties — storage_static_routes.storage_routes.subnets.ipv6 / 7878306aa399 / 3

<a id="canonical-62e4870346b5dae7ce3c0aaa4a486f51468d3f32271fbf4e7da05d50d58d7fa9"></a>

<a id="canonical-47948fe90df672259259e7c9408d85bd106b8db73a832d6ed80f9a8f18506a5e"></a>

## plen property — storage_static_routes.storage_routes.subnets.ipv6 / 7878306aa399 / 4

Type: `"number"`. Optional.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128,
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
    "ves.io.schema.rules.uint32.lte": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  }
}
```

<a id="canonical-d39bdacde055b165e79d0ccc8298d77cf3a102bcf4d3369e7d67decf55317347"></a>

<a id="canonical-631e503dd3e313b6e6ac014d96795bc5b32f8f93720ca8006bf2c4db87366a8f"></a>

## prefix property — storage_static_routes.storage_routes.subnets.ipv6 / 7878306aa399 / 5

Type: `"string"`. Optional.

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. '2001:db8:0:0:0:2:0:0' The address can be compacted by
suppressing zeros e.g. '2001:db8::2::'.

Upstream description:

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. "2001:db8:0:0:0:2:0:0" The address can be compacted by
suppressing zeros e.g. "2001:db8::2::"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-0bf916243c7ae239742c7b83980f904e95b3ec400d90d41ab98a0c5d5f7c4c9e"></a>

## Next pages — storage_static_routes.storage_routes.subnets.ipv6 / 7878306aa399 / 6

- [storage_static_routes.storage_routes.subnets](resources--fleet--reference--group-004.md#canonical-0537b4bf9f1ce278cb306f5091283296f27e7095a5ba95b33fea735d5d45b2aa)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-c7daec114bab3c6e240b23c9acc29c136c7b4876645a944004681d627803d664"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-46f221a5ca1f14c714f6d47bfc3d6f873a07169b76b75c23a29326028c5600bd"></a>

## timeouts — timeouts / f9519ccc33b8 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- timeouts

<a id="canonical-11ff33128bee74e8ef1fac4bfee63378effcbb132cdf821d3eb28aa7f1a7ba8d"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-a4960bf63a90642f6c83be386eae8654ec4c665b873a6f3af11e4837affc426d"></a>

## Direct properties — timeouts / f9519ccc33b8 / 3

<a id="canonical-fadb5a7b2cbf436d4d3c795a0b9600dd1e44c59da22af3acff2b82ab0215471a"></a>

<a id="canonical-30609c1de620649d754c49471fb6837cfa416a6a74f2834d874f04886c569ae8"></a>

## create property — timeouts / f9519ccc33b8 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-e37c04db6a43a74f4056725e9915fc988e58779748d24c93b858140abb176d4e"></a>

<a id="canonical-ff4a03a28640542dd3229a9ed4c5781df0d496e2e20dff9bf675a248b44c06ee"></a>

## delete property — timeouts / f9519ccc33b8 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-48b318d5034cc2374beb8c8383cd7268315bc9c27319b19723f4199455c828b2"></a>

<a id="canonical-f218abd414ccd0970006bd72269a68b25ad43a032e96098e1c596c05019b389b"></a>

## read property — timeouts / f9519ccc33b8 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-a584ae4fbc1b7398328321881f4b0bed5758decca954d39ad1ed01d120f0158c"></a>

<a id="canonical-2610043ec3e097e20611bdd8bd88e4c1c08f8e56f9f59298963a3a56e81005d2"></a>

## update property — timeouts / f9519ccc33b8 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-b15797d4dcb6b546a0b5490c5864a3bace5ff0f3c97d4160472593b9e9e985e3"></a>

## Next pages — timeouts / f9519ccc33b8 / 8

- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-777cf2733ada0f5b1100fb3c706ce2c64aefbf9dc25e8aae32a1eaa74ade7df2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-85dccb2e89fb4613ea00cf1e87a04c7845c1224baffd34d102b004cb1c12f9ee"></a>

## usb_policy — usb_policy / 25f9bfaa3e7e / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- usb_policy

<a id="canonical-5a526f17d1768fe3878cc166b4741dfc79b3df781f6f88d652bf81ec9a8992ce"></a>

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
usb_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-05c30501705f72d4881cd9d43fc8fe1d9de21e1727afd93db1c9e62bdf3c7c1c"></a>

## Direct properties — usb_policy / 25f9bfaa3e7e / 3

<a id="canonical-243776d44787bbe8f928baa09df93bbba3395712f73bff8c00f82d6f38fdf068"></a>

<a id="canonical-b9099dbd102fc99bb13c90ff888c3287922b87c118fa63c26afa3c30741f02be"></a>

## name property — usb_policy / 25f9bfaa3e7e / 4

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

<a id="canonical-808d59665a84fec1f43fdf13f8f0a5de302df7551d8776bb9715381b42f486a6"></a>

<a id="canonical-87e60166f3863d80c8c729b8bc5eb7a3d7ced8193c90151195692a7ed4b5b37b"></a>

## namespace property — usb_policy / 25f9bfaa3e7e / 5

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

<a id="canonical-7a1d5cf0e6c87838f35cdd5ea8976470dc44f95c84bad2d3ec576dbaab564e22"></a>

<a id="canonical-77afc3ad6eb3264ac668fd917659c4c6f2809d3225813946303f4ebe77e93795"></a>

## tenant property — usb_policy / 25f9bfaa3e7e / 6

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

<a id="canonical-422376e47c2af7c471717f6d290fea225900679a20e0fb2678e971dcba565fd2"></a>

## Next pages — usb_policy / 25f9bfaa3e7e / 7

- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
