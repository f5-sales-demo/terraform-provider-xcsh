---
page_title: "xcsh_fleet reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fleet reference."
---

# xcsh_fleet reference

<a id="canonical-e58c6d4eae482ccf40d4de9274054ecec42f302bbc0d8009daa503dc449fdb4a"></a>

## Direct properties — storage_device_list.storage_devices.hpe_storage / 1aeb90c98c0c / 3

<a id="canonical-106428d6e2e564d3731ac6db4eedd7270634a3482e5005c9ee8654bed88a83c1"></a>

<a id="canonical-e20c04c4ea5b98f8f5fc2bba1fa3d5e66f78de1584909af82a75e13cf3c3c942"></a>

## api_server_port property — storage_device_list.storage_devices.hpe_storage / 1aeb90c98c0c / 4

Type: `"number"`. Optional.

Storage server Port. Enter Storage Server Port.

Upstream description:

Enter Storage Server Port.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [iscsi_chap_password](resources--fleet--reference--group-003.md#canonical-289e4d06d45a889912c9acc7e5649c25be39ad6133cb0ae7b5019e55ec97919e): complete subsection reference.

<a id="canonical-572f5fb055b30fddbba565541374f58ea46aac3a4bae0df5c44af57d8856ac33"></a>

<a id="canonical-71a45f05394d4cfb1004e2f577898994f9005500f119a57e9cbfdda3e2afa8e4"></a>

## iscsi_chap_user property — storage_device_list.storage_devices.hpe_storage / 1aeb90c98c0c / 5

Type: `"string"`. Optional.

Chap Username to connect to the HPE storage.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [password](resources--fleet--reference--group-003.md#canonical-66462941a04187147e9d4f39a7c458b955f3e740aaa463ec4459f98fc68c9792): complete subsection reference.

<a id="canonical-dc749a13b470bacf66a132ad567d25ccbe9fc9a537b5f588bacca2fd426893d8"></a>

<a id="canonical-ce6ccc145cae6d4194d5cc177e5108285e4e1ec46a16b3919656ca38fc7d265f"></a>

## storage_server_ip_address property — storage_device_list.storage_devices.hpe_storage / 1aeb90c98c0c / 6

Type: `"string"`. Optional.

Storage Server IP address. Enter storage server IP address.

Upstream description:

Enter storage server IP address.

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

<a id="canonical-4ea4847e68b58a9a5bd990814a7406c0c9d10474518d322788909af3f01f4b50"></a>

<a id="canonical-a54de1a65db9cba95174b5183a48ab6f3b0735656c060b8b3db1156c784f2f67"></a>

## storage_server_name property — storage_device_list.storage_devices.hpe_storage / 1aeb90c98c0c / 7

Type: `"string"`. Optional.

Storage Server Name. Enter storage server Name.

Upstream description:

Enter storage server Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-9980cc76d5991a62bb04ce97cfb1c66cfbe570a7a0849dd3f31c1ca668ccf955"></a>

<a id="canonical-7677452deb1abb29170016bfeecce2419bf8955c8974c0b0a42f793e7f95c780"></a>

## username property — storage_device_list.storage_devices.hpe_storage / 1aeb90c98c0c / 8

Type: `"string"`. Optional.

Username to connect to the HPE storage management IP.

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
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-14b3430932b38cdefc63cd8af88257478cf93a039389902558edfbe06187fac3"></a>

## Next pages — storage_device_list.storage_devices.hpe_storage / 1aeb90c98c0c / 9

- [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](resources--fleet--reference--group-003.md#canonical-289e4d06d45a889912c9acc7e5649c25be39ad6133cb0ae7b5019e55ec97919e)
- [storage_device_list.storage_devices.hpe_storage.password](resources--fleet--reference--group-003.md#canonical-66462941a04187147e9d4f39a7c458b955f3e740aaa463ec4459f98fc68c9792)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-289e4d06d45a889912c9acc7e5649c25be39ad6133cb0ae7b5019e55ec97919e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-effe63d3eafaadc5d64e2f6e2bd401a94ece13a96ba1570573108dc25ecf9269"></a>

## storage_device_list.storage_devices.hpe_storage.iscsi_chap_password — storage_device_list.storage_devices.hpe_storage.iscsi_chap_password / 37beb98b2a8e / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.hpe_storage](resources--fleet--reference--group-002.md#canonical-3d22894c75dd7fbacace9c29569e97b7e7f9e8ac159d7c74cd40e547598fd937)
- storage_device_list.storage_devices.hpe_storage.iscsi_chap_password

<a id="canonical-7f8000444a74c9fd256601cf411654d0b0649b2ea27f7ef2530af8c04ec88ab4"></a>

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
iscsi_chap_password {
  # Configure direct properties listed below.
}
```

<a id="canonical-969b3d2cefa306808dfb3152f820ca1cccb78a49e4896e6c22ecb7805c9ee80c"></a>

## Direct properties — storage_device_list.storage_devices.hpe_storage.iscsi_chap_password / 37beb98b2a8e / 3

- [blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-94aa965cfb00ada26aa76f3da8d0a0fd8b76ef033230f36bbe8cc7803c16eba3): complete subsection reference.

- [clear_secret_info](resources--fleet--reference--group-003.md#canonical-79489ba771706434124d8754a33cbbc518524732937f100c64474fc3d8eb3187): complete subsection reference.

<a id="canonical-7b64bb239daf7c65aa4d97ffc2d1ff0e29f468fff3cc65cb8e479fd0d093c6ab"></a>

## Next pages — storage_device_list.storage_devices.hpe_storage.iscsi_chap_password / 37beb98b2a8e / 4

- [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-94aa965cfb00ada26aa76f3da8d0a0fd8b76ef033230f36bbe8cc7803c16eba3)
- [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info](resources--fleet--reference--group-003.md#canonical-79489ba771706434124d8754a33cbbc518524732937f100c64474fc3d8eb3187)
- [storage_device_list.storage_devices.hpe_storage](resources--fleet--reference--group-002.md#canonical-3d22894c75dd7fbacace9c29569e97b7e7f9e8ac159d7c74cd40e547598fd937)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-94aa965cfb00ada26aa76f3da8d0a0fd8b76ef033230f36bbe8cc7803c16eba3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c35810d58f939684adaa4dda5eb0d50c99b686a7a68f4bc92379cfe49089c7c"></a>

## storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info — storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_se / 977aadbac432 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.hpe_storage](resources--fleet--reference--group-002.md#canonical-3d22894c75dd7fbacace9c29569e97b7e7f9e8ac159d7c74cd40e547598fd937)
- [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](resources--fleet--reference--group-003.md#canonical-289e4d06d45a889912c9acc7e5649c25be39ad6133cb0ae7b5019e55ec97919e)
- storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info

<a id="canonical-858e1617458c60adf918fab6f282400f11e29f0f74a10e16f74d7a24d948d3c8"></a>

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

<a id="canonical-aeb3bc003e928edb516bbfd33138ec13ae41ded1609d0863991883619df65650"></a>

## Direct properties — storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_se / 977aadbac432 / 3

<a id="canonical-34ad5f93daa7e500515886c9439ab467aa1a40f459620ce18dbc2ebec7f0e7f0"></a>

<a id="canonical-da47cdadc29fa98ef2384b3887cec3a21af196c0a7beb71775a1033da77a420c"></a>

## decryption_provider property — storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_se / 977aadbac432 / 4

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

<a id="canonical-29d338b7991418763b389b586dcf48b3009af7fa9fb3d327cc034bc69c810f6d"></a>

<a id="canonical-37cbdb91adc43bb7d571cf3cc4bc6c110d15fbe7d879a4c2051ae5ab0abc83c8"></a>

## location property — storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_se / 977aadbac432 / 5

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

<a id="canonical-f507a9b9bf08926afa606f5930810388d25c6788b2bef3fd1635c5eb8dc92603"></a>

<a id="canonical-daba958a37912d6c7862ccfee7782eda316cb92351908a08749738b8ffcade49"></a>

## store_provider property — storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_se / 977aadbac432 / 6

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

<a id="canonical-03eb473204d688d1c70b71a6851dc12b7e9f49912ef9b781ada105e79e025a84"></a>

## Next pages — storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_se / 977aadbac432 / 7

- [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](resources--fleet--reference--group-003.md#canonical-289e4d06d45a889912c9acc7e5649c25be39ad6133cb0ae7b5019e55ec97919e)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-79489ba771706434124d8754a33cbbc518524732937f100c64474fc3d8eb3187"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-47e94d25128f4a50abf4119ff8a547e794d4e1b81417158efc9edd035d29792f"></a>

## storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info — storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret / f3453567558c / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.hpe_storage](resources--fleet--reference--group-002.md#canonical-3d22894c75dd7fbacace9c29569e97b7e7f9e8ac159d7c74cd40e547598fd937)
- [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](resources--fleet--reference--group-003.md#canonical-289e4d06d45a889912c9acc7e5649c25be39ad6133cb0ae7b5019e55ec97919e)
- storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info

<a id="canonical-8dce94aa0c54d08cbf4605a380a0556b47b53e32a89b269d080ac7b324064cc9"></a>

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

<a id="canonical-2e4238113ed7adef5ff81ffb3521fb23433ea46525503e33d8ac294a476bbc0f"></a>

## Direct properties — storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret / f3453567558c / 3

<a id="canonical-7536415376ea51cd94688a7273f8a0b3aa69774e5be704fdb835ed0172cae5a3"></a>

<a id="canonical-4403c262c32e812dae525f9677251c6071d6d0df16f34c4e341ca2046fa4e8de"></a>

## provider_ref property — storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret / f3453567558c / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0b15c10824887434b2d91ffdc691af1dfc9406e0dd3c0fb77d947a0ba62f26d0"></a>

<a id="canonical-41c410cd1f63e7b94f6e07823960fe6c1424e6410ad6a6ab72568b747de6600f"></a>

## url property — storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret / f3453567558c / 5

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

<a id="canonical-7c0b4552082d23272fcf6844bc9868f60ffa90d9bf6e20ab1941e95fbf0a19fc"></a>

## Next pages — storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret / f3453567558c / 6

- [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](resources--fleet--reference--group-003.md#canonical-289e4d06d45a889912c9acc7e5649c25be39ad6133cb0ae7b5019e55ec97919e)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-66462941a04187147e9d4f39a7c458b955f3e740aaa463ec4459f98fc68c9792"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-09be659c58cbb302273ee3ebeea2f8068ff40df2d7439009b3ce7785ca74d92e"></a>

## storage_device_list.storage_devices.hpe_storage.password — storage_device_list.storage_devices.hpe_storage.password / 1af9d0e1746a / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.hpe_storage](resources--fleet--reference--group-002.md#canonical-3d22894c75dd7fbacace9c29569e97b7e7f9e8ac159d7c74cd40e547598fd937)
- storage_device_list.storage_devices.hpe_storage.password

<a id="canonical-9d0c81aa7622a68dcfcfa088a31326c52ce5be59cff6a4357cc0a5370c3c2d20"></a>

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

<a id="canonical-7db8ec4f4696ff3c314dacf4ed0f5398ed16021ce68f3103c8d950a6dac7a065"></a>

## Direct properties — storage_device_list.storage_devices.hpe_storage.password / 1af9d0e1746a / 3

- [blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-041c2fd2fedca9ab46a9323a70045f6993eba4ba1c13d90cd2754a8c775424a5): complete subsection reference.

- [clear_secret_info](resources--fleet--reference--group-003.md#canonical-5f03176b6cc05d6911b9748820e41a3b239f583cd2496fadaadd791fd2e90b1b): complete subsection reference.

<a id="canonical-b073821dde88620a06f5d84660767c4c6d3615913fb931cbfcc0f0760eda1e8d"></a>

## Next pages — storage_device_list.storage_devices.hpe_storage.password / 1af9d0e1746a / 4

- [storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-041c2fd2fedca9ab46a9323a70045f6993eba4ba1c13d90cd2754a8c775424a5)
- [storage_device_list.storage_devices.hpe_storage.password.clear_secret_info](resources--fleet--reference--group-003.md#canonical-5f03176b6cc05d6911b9748820e41a3b239f583cd2496fadaadd791fd2e90b1b)
- [storage_device_list.storage_devices.hpe_storage](resources--fleet--reference--group-002.md#canonical-3d22894c75dd7fbacace9c29569e97b7e7f9e8ac159d7c74cd40e547598fd937)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-041c2fd2fedca9ab46a9323a70045f6993eba4ba1c13d90cd2754a8c775424a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-539061ffbd160eca4fd9887e2200c0f2d5239a7afdaf4e9f4250c126d6d325c7"></a>

## storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info — storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info / e27109de0213 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.hpe_storage](resources--fleet--reference--group-002.md#canonical-3d22894c75dd7fbacace9c29569e97b7e7f9e8ac159d7c74cd40e547598fd937)
- [storage_device_list.storage_devices.hpe_storage.password](resources--fleet--reference--group-003.md#canonical-66462941a04187147e9d4f39a7c458b955f3e740aaa463ec4459f98fc68c9792)
- storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info

<a id="canonical-759557e0475f901b594ba22f622de4c225ead2edd6c35328c0e5165315d9e0b9"></a>

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

<a id="canonical-c1d854fb99583c685bfa2d37be5439c7d01ad4c0dc40f23e9ec2771a21c79372"></a>

## Direct properties — storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info / e27109de0213 / 3

<a id="canonical-6c8259e30b4e82b886655ab6e28e20bde7f8e5fcd6815bbd3b3ca5485ab903d3"></a>

<a id="canonical-129e6a3ab3c9d37d3bc5a2eedb0bdc94d0ca76ae68ac5630b67d237f04f9daf0"></a>

## decryption_provider property — storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info / e27109de0213 / 4

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

<a id="canonical-841fb9b5364038b2b7a455b4e970405afde6e040070baf18f5087e563e73e584"></a>

<a id="canonical-9e8fd79048727cf484a388d81e2387e84d6532f851b46cb5874706634cac8a96"></a>

## location property — storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info / e27109de0213 / 5

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

<a id="canonical-7df905360e30c54889bbf0db2f8d7acd7f3ae6e68857fc3408716b4f6c4b578f"></a>

<a id="canonical-39de95c64353a08034b5b12bedbc397b9c0b4a0c744f923c519ba551c6b900eb"></a>

## store_provider property — storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info / e27109de0213 / 6

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

<a id="canonical-6014eea1648af5387e6af20d2a87922533db4efd554cfef6829695ed6148e805"></a>

## Next pages — storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info / e27109de0213 / 7

- [storage_device_list.storage_devices.hpe_storage.password](resources--fleet--reference--group-003.md#canonical-66462941a04187147e9d4f39a7c458b955f3e740aaa463ec4459f98fc68c9792)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-5f03176b6cc05d6911b9748820e41a3b239f583cd2496fadaadd791fd2e90b1b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e49f33eb6b6c4610bad13953d3740a7c3a3e7dfd504a7d2357af098177e56219"></a>

## storage_device_list.storage_devices.hpe_storage.password.clear_secret_info — storage_device_list.storage_devices.hpe_storage.password.clear_secret_info / 7ef514518223 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.hpe_storage](resources--fleet--reference--group-002.md#canonical-3d22894c75dd7fbacace9c29569e97b7e7f9e8ac159d7c74cd40e547598fd937)
- [storage_device_list.storage_devices.hpe_storage.password](resources--fleet--reference--group-003.md#canonical-66462941a04187147e9d4f39a7c458b955f3e740aaa463ec4459f98fc68c9792)
- storage_device_list.storage_devices.hpe_storage.password.clear_secret_info

<a id="canonical-f6af0b2aa94bc783c877b1d21edc329b6e4fed8d031ea3b546af0e0a8dd98d1e"></a>

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

<a id="canonical-1ae3ae2fb481c793259e896e588aea727aeeb6a1d839d0651a4d36e2653e5ff1"></a>

## Direct properties — storage_device_list.storage_devices.hpe_storage.password.clear_secret_info / 7ef514518223 / 3

<a id="canonical-daa21dc01df2087932566aa231d998ed82f309538c17f4bfd16a9ad4559b01be"></a>

<a id="canonical-a1d8b733bdac2e3d5b9863f193fb6add692abd93a0c36ec52b776b6b25dfb05b"></a>

## provider_ref property — storage_device_list.storage_devices.hpe_storage.password.clear_secret_info / 7ef514518223 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-7b3098fc5fb55607c8dc6a3dfc69076173fba7eea8a02e8aace90144408b1b52"></a>

<a id="canonical-9468c71e165b6fef959f6a5a4c936c1f8360361f192c1e36ba83f7560f0371f8"></a>

## url property — storage_device_list.storage_devices.hpe_storage.password.clear_secret_info / 7ef514518223 / 5

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

<a id="canonical-ba410fa90e57d759641c79477debba5a60acff3684a7f1bd373d78061b2f2a59"></a>

## Next pages — storage_device_list.storage_devices.hpe_storage.password.clear_secret_info / 7ef514518223 / 6

- [storage_device_list.storage_devices.hpe_storage.password](resources--fleet--reference--group-003.md#canonical-66462941a04187147e9d4f39a7c458b955f3e740aaa463ec4459f98fc68c9792)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-adff9dd5bf85cbf2a1b04c2338f782268be35b62daf728089ebca21616ce2527"></a>

## storage_device_list.storage_devices.netapp_trident — storage_device_list.storage_devices.netapp_trident / 49484080be07 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- storage_device_list.storage_devices.netapp_trident

<a id="canonical-d1244af590bd3a828da9363c0e7307104eb60654f746a495b1d5ff0d41f243cf"></a>

Type: `"object"`. single nested block, Optional.

Device configuration for NetApp Trident Storage.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("netapp_backend_ontap_nas",
    "netapp_backend_ontap_san")}
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
  "x-ves-oneof-field-backend_choice": "[\"netapp_backend_ontap_nas\",\"netapp_backend_ontap_san\"]"
}
```

Terraform syntax:

```terraform
netapp_trident {
  # Configure direct properties listed below.
}
```

<a id="canonical-ed8a5c057a088ca07955b43d1f07a9bd5e402c4ec325735493ab048eaad8885f"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident / 49484080be07 / 3

- [netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-26cd2b551b9e5c0a8fec323e52bc1d0eed1c17d073fde97b216c79c62b5c9a5d): complete subsection reference.

- [netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-9aaa33d28bdc9675ddfeacd2a1549952db4fb3dee17435ce405d097f40edb5d1): complete subsection reference.

<a id="canonical-15563a52ade7d211c239cb02c1fa8e0e4290ba1e82a8e437b8f9b979c18cb1b9"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident / 49484080be07 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-26cd2b551b9e5c0a8fec323e52bc1d0eed1c17d073fde97b216c79c62b5c9a5d)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-9aaa33d28bdc9675ddfeacd2a1549952db4fb3dee17435ce405d097f40edb5d1)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-26cd2b551b9e5c0a8fec323e52bc1d0eed1c17d073fde97b216c79c62b5c9a5d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f0ee323e24060f9f68f361e37fd8fec54445f52087a8e942fdf1ff702d94bdfd"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / 63a29a3bc81b / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas

<a id="canonical-7017ee263703d14d32a314dc98d2d2a4d76256bfaf7a0b131f1b82d543d9ba5e"></a>

Type: `"object"`. single nested block, Optional.

Configuration of storage backend for NetApp ONTAP NAS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("storage_driver_name",
    "username"),
  validators.ConflictingObjectAttributes("data_lif_dns_name",
    "data_lif_ip"),
  validators.ConflictingObjectAttributes("management_lif_dns_name",
    "management_lif_ip")}
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
  "x-ves-oneof-field-data_lif": "[\"data_lif_dns_name\",\"data_lif_ip\"]",
  "x-ves-oneof-field-management_lif": "[\"management_lif_dns_name\",\"management_lif_ip\"]"
}
```

Terraform syntax:

```terraform
netapp_backend_ontap_nas {
  # Configure direct properties listed below.
}
```

<a id="canonical-f2959466dea2a9943ab577674f41949e4cee9659c227ab63a4826a65581f401b"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / 63a29a3bc81b / 3

- [auto_export_cidrs](resources--fleet--reference--group-003.md#canonical-05d438ad6b00d7d61c7862769883ee2cdaba89841adec93b53af39b210fe0f38): complete subsection reference.

<a id="canonical-6ddf9c8800245cdbbcf3de046b2e94651f9ed6e35d80dc38cfd0ded0168ec1ac"></a>

<a id="canonical-0a311c4cc9e1030f7f915c8d327dbeef7cf2109c13cf977f55c37b6121374d12"></a>

## auto_export_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / 63a29a3bc81b / 4

Type: `"bool"`. Optional.

Policy configuration for this feature.

Upstream description:

Enable automatic export policy creation and updating.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-6436e772e5bba19236d81a6573898bd628b05e27da56b9b3e41f4b208bd5182f"></a>

<a id="canonical-c379696662bb8a3d152de8ac0ccac9eacde027b6359bf1820dabf09dc9a5e1f3"></a>

## backend_name property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / 63a29a3bc81b / 5

Type: `"string"`. Optional.

Configuration of Backend Name. Driver is name + '\_' + dataLIF.

Upstream description:

Configuration of Backend Name. Driver is name + "\_" + dataLIF.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 50),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 50,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 50,
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
    "ves.io.schema.rules.string.max_len": "50",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "50",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-c1dbebef51111a88a875f6552806c709452d47a4bd9f6d4d0fc30a7500422028"></a>

<a id="canonical-dc86cd868ef24a474a4dd78df7a77868655a79d0b7358f3b3b1b3e369cd7898a"></a>

## client_certificate property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / 63a29a3bc81b / 6

Type: `"string"`. Optional.

Please Enter Base64-encoded value of client certificate. Used for certificate-based auth.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

- [client_private_key](resources--fleet--reference--group-003.md#canonical-896f93cd9da32fde87bab850656939f55bd75efe7fc25ed0fe91c03a54cdca29): complete subsection reference.

<a id="canonical-e4ab7fd866cbb9d3dfdcea2ef1e5d5b51d3351d1a199a41c318cd4a2044a80a6"></a>

<a id="canonical-b14530aff0e4b8f79444458daf3432083a1005384638c77df518c80eae6ac7a5"></a>

## data_lif_dns_name property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / 63a29a3bc81b / 7

Type: `"string"`. Optional.

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

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

<a id="canonical-05d05409462dcb4cb0e95415b5d29ae66ea3e7d3bdf4c49531de075bf2e09988"></a>

<a id="canonical-be5f9867b8a32804f9383b40f655daafeb9b179ecb0f60ac3e709fd83d383560"></a>

## data_lif_ip property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / 63a29a3bc81b / 8

Type: `"string"`. Optional.

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

Upstream description:

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

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

<a id="canonical-5de6d27c3902c0549bc36e7c72496da91dded1180870300d44d3a58f1e29b0ae"></a>

<a id="canonical-310635622b897f845e32bada9860e566080d86a93781f9cbdac463ec7b22401c"></a>

## labels property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / 63a29a3bc81b / 9

Type: `["map", "string"]`. Optional.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class selection.

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

<a id="canonical-66188cbd50447b8a2fec8fcbfa04d10dba4a6ec28c3ada14d12b4eb3f80a71be"></a>

<a id="canonical-6c0a9091956f91973dfea8b61b169a2d55cbb0265ec53f4478ecd01f2b1abf4b"></a>

## limit_aggregate_usage property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / 63a29a3bc81b / 10

Type: `"string"`. Optional.

Fail provisioning if usage is above this percentage. Not enforced by default.

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

<a id="canonical-86c25388d04ef44756d88b14ae88bb00fd736c294daad60c5afe25b575ce25b3"></a>

<a id="canonical-2745b94fb16e520abfbb666612d61d15fbcb98353320cda044a07af41782b76a"></a>

## limit_volume_size property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / 63a29a3bc81b / 11

Type: `"string"`. Optional.

Fail provisioning if requested volume size is above this value. Not enforced by default.

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

<a id="canonical-f18601be4ed99b05c45a4868424b24619a48252d47351a2945d4edd0ca43e1b3"></a>

<a id="canonical-abb8e98ff62a3e35838b2ad2c43439ee5856e5246233ce1b93f560a7db3a07b1"></a>

## management_lif_dns_name property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / 63a29a3bc81b / 12

Type: `"string"`. Optional.

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

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

<a id="canonical-1c5e928d2d81a91e8c51b4e6a9e09affcbc2b20e67a26f5ac844185114f15512"></a>

<a id="canonical-012515f0eba259313ba6ed39981990cd4600f56f9cb097afdbd33830fe9ea960"></a>

## management_lif_ip property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / 63a29a3bc81b / 13

Type: `"string"`. Optional.

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

Upstream description:

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

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

<a id="canonical-baf668d18e983eb9b3ff9a99e7115467d3a6805d74a7065eddf7bd6a40ed8934"></a>

<a id="canonical-81876369d5b8913b436c6ef2840048baa18207519503e943d5dfcc026af1c01c"></a>

## nfs_mount_options property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / 63a29a3bc81b / 14

Type: `"string"`. Optional.

Comma-separated list of NFS mount OPTIONS. Not enforced by default.

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

- [password](resources--fleet--reference--group-003.md#canonical-70e181e855d483fe94870fab7cb06991de72402434892bbcace142a82526b5cf): complete subsection reference.

<a id="canonical-fd9e975e8bb198e68303d66d5d010cda23facb7aa5dd383d46502a3600be435b"></a>

<a id="canonical-94fe3f6bfbbccf318b456139007042e82ad716ce4373c96cff66e6694c4ca3e6"></a>

## region property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / 63a29a3bc81b / 15

Type: `"string"`. Optional.

Backend Region. Virtual Pool Region.

Upstream description:

Virtual Pool Region.

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

- [storage](resources--fleet--reference--group-003.md#canonical-2e977c79263e7217d825d481adf03b7e8da6d705b2ac7838cceeebf0a23931b0): complete subsection reference.

<a id="canonical-cce9c1c995c7023b9fac27fbe17b6bfbe72d4c5c85e14b1e8df96f7475358ffc"></a>

<a id="canonical-1855d259850d6ed99926bae3213bb7dc9a783ac3edd28b2df69d11ac0d0cb30b"></a>

## storage_driver_name property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / 63a29a3bc81b / 16

Type: `"string"`. Optional.

\[Enum: ontap-nas|ontap-nas-economy|ontap-nas-flexgroup\] Storage Backend Driver. Configuration of
Backend Name. Possible values are \`ontap-nas\`, \`ontap-nas-economy\`, \`ontap-nas-flexgroup\`.

Upstream description:

Configuration of Backend Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ontap-nas",
    "ontap-nas-economy",
    "ontap-nas-flexgroup"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ontap-nas",
    "ontap-nas-economy",
    "ontap-nas-flexgroup"
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
    "ves.io.schema.rules.string.in": "[\\\"ontap-nas\\\",\\\"ontap-nas-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ontap-nas\\\",\\\"ontap-nas-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  }
}
```

<a id="canonical-2ee5f0854eaee5c88671c396edb8860d49151ddc7e8c59357fbc6cd0fdccc1e7"></a>

<a id="canonical-5c26c662f2f9fd2ff4c6149c62c5293e1274a2aeb051399c5318561db444afd4"></a>

## storage_prefix property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / 63a29a3bc81b / 17

Type: `"string"`. Optional.

Prefix used when provisioning new volumes in the SVM. Once set this cannot be updated.

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

<a id="canonical-90d3cceb48d1fcd337cf08487e65cbba5555ae2bf963257595b8c45682888d5f"></a>

<a id="canonical-22c5089e11c1b5c3e1d1916514f4a4613a3576045cd8e8df1051eb0095560d46"></a>

## svm property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / 63a29a3bc81b / 18

Type: `"string"`. Optional.

Storage virtual machine to use. Derived if an SVM managementLIF is specified.

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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-f861a288c4ecb5fc7c068b15fba975a2bc4ef234d3330ec5c6120deebc9f86ac"></a>

<a id="canonical-ffd69793ccf9b4dac8d613b52d9a68c85d1ebf2114e398b5cba1d2edc51162dd"></a>

## trusted_ca_certificate property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / 63a29a3bc81b / 19

Type: `"string"`. Optional.

Please Enter Base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth.

Upstream description:

Please Enter Base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth..

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

<a id="canonical-c65c77bc7d93a2f58d1e870c5bd47b29e4dc6c7f489e80e1cc3aab99c02bf374"></a>

<a id="canonical-e6c8c940b3b8b26c9014e9a947a7a932b29b31db9f333a45bf0355b3a5c17b76"></a>

## username property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / 63a29a3bc81b / 20

Type: `"string"`. Optional.

Username. Username to connect to the cluster/SVM.

Upstream description:

Username to connect to the cluster/SVM.

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
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [volume_defaults](resources--fleet--reference--group-003.md#canonical-bd864a49908d14cbed12e5febd07a0acf564c0f324249e9ca457d51c4d0939fd): complete subsection reference.

<a id="canonical-37444614640ec39ab0674cac400cfe9192262b437ff36d289399405c5d3c1656"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / 63a29a3bc81b / 21

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs](resources--fleet--reference--group-003.md#canonical-05d438ad6b00d7d61c7862769883ee2cdaba89841adec93b53af39b210fe0f38)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](resources--fleet--reference--group-003.md#canonical-896f93cd9da32fde87bab850656939f55bd75efe7fc25ed0fe91c03a54cdca29)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](resources--fleet--reference--group-003.md#canonical-70e181e855d483fe94870fab7cb06991de72402434892bbcace142a82526b5cf)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](resources--fleet--reference--group-003.md#canonical-2e977c79263e7217d825d481adf03b7e8da6d705b2ac7838cceeebf0a23931b0)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults](resources--fleet--reference--group-003.md#canonical-bd864a49908d14cbed12e5febd07a0acf564c0f324249e9ca457d51c4d0939fd)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-05d438ad6b00d7d61c7862769883ee2cdaba89841adec93b53af39b210fe0f38"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a84d7be45514411732aaedd157eaecf61269acc70571fe5cd93ff804abf65dcb"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto / 773ce5e25d1d / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-26cd2b551b9e5c0a8fec323e52bc1d0eed1c17d073fde97b216c79c62b5c9a5d)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs

<a id="canonical-495339d4f23594f87edbb167a689e3a539392d92b5911f8d40bcea15dca87675"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

Receipt-pinned upstream constraints:

```json
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
auto_export_cidrs {
  # Configure direct properties listed below.
}
```

<a id="canonical-b5cdcf3a376ac4502f4d0dd3ed3f3a32fd7ff0dc23b4f98cd96b22cf5c39fa76"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto / 773ce5e25d1d / 3

<a id="canonical-1a772160af81b4f0b30f34e3565f9809fbfc59a44ce02ecbd3fd29372d62a853"></a>

<a id="canonical-0ed5e89d91261bceec40b3e91101c64863951b73d4a262db625cfff01bc2f953"></a>

## prefixes property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto / 773ce5e25d1d / 4

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-bda49e2de176898d2b95a4fab19e93914793d77893f9b6e00f25f71a9da9ba47"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto / 773ce5e25d1d / 5

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-26cd2b551b9e5c0a8fec323e52bc1d0eed1c17d073fde97b216c79c62b5c9a5d)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-896f93cd9da32fde87bab850656939f55bd75efe7fc25ed0fe91c03a54cdca29"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e1e41c2ebd69416e50d93dc694af96c2e7e2ab752e68771efab7d5239576badd"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.clie / fac51835556a / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-26cd2b551b9e5c0a8fec323e52bc1d0eed1c17d073fde97b216c79c62b5c9a5d)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key

<a id="canonical-6c5d6d2deb3975fd415ff6fc44eee0366fc4b3bbd2efc2f5b49b43754f5e1870"></a>

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
client_private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-14faadd25c7038b1af4411e2e7d33bd2864e801623f8fea27889373d1f6982da"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.clie / fac51835556a / 3

- [blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-3d2deb963de0b1563ff8e9c43029b9cb678b8ae2f3081a1703e5dbb4c0e0e2ce): complete subsection reference.

- [clear_secret_info](resources--fleet--reference--group-003.md#canonical-79458c18a6ed9f5ee19d730f917caf2214339d317fdf4f668563c5dd8c201283): complete subsection reference.

<a id="canonical-54849d5ff0b1346389f698aa3dc94be217d81382b6d91accc1efe199f7a0c6e6"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.clie / fac51835556a / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-3d2deb963de0b1563ff8e9c43029b9cb678b8ae2f3081a1703e5dbb4c0e0e2ce)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info](resources--fleet--reference--group-003.md#canonical-79458c18a6ed9f5ee19d730f917caf2214339d317fdf4f668563c5dd8c201283)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-26cd2b551b9e5c0a8fec323e52bc1d0eed1c17d073fde97b216c79c62b5c9a5d)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-3d2deb963de0b1563ff8e9c43029b9cb678b8ae2f3081a1703e5dbb4c0e0e2ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0cf54e50d759de4588c20308bead89988a8e454d3f7e72e5349460b5f483d09c"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.clie / 1fa89abe3584 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-26cd2b551b9e5c0a8fec323e52bc1d0eed1c17d073fde97b216c79c62b5c9a5d)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](resources--fleet--reference--group-003.md#canonical-896f93cd9da32fde87bab850656939f55bd75efe7fc25ed0fe91c03a54cdca29)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info

<a id="canonical-9478402b411b488e47b6e3ac621057fa48811cd12a95354adb11127d20e4053b"></a>

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

<a id="canonical-cd62de24bc995c7bd67217847955c592f91f625cefbc67a0f4c53f8e34225b72"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.clie / 1fa89abe3584 / 3

<a id="canonical-20ce298eb884a4feb435876497b722d66ac806cb512b376bc216b92a11d26fbd"></a>

<a id="canonical-f37de3bd4dbab6b6bd2e0ce2ff4cefe843e5573f8acf261167534ee3572b2410"></a>

## decryption_provider property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.clie / 1fa89abe3584 / 4

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

<a id="canonical-43a46cee6a1e9333885da725ba98209f5d542383326fde9dcf9bcb7c3a5c8770"></a>

<a id="canonical-9fd7762d939066391dad5d649c5c46887a8277d86d75a631baa9a16e7e6b3a4d"></a>

## location property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.clie / 1fa89abe3584 / 5

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

<a id="canonical-d49bb71bd9e35122614c8265bf67bdb09286b95793f4223c836e46777045858a"></a>

<a id="canonical-4a0c28c868563a588f102f1d995ec564190f31049e0e74f9842f0b065aa232a9"></a>

## store_provider property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.clie / 1fa89abe3584 / 6

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

<a id="canonical-09f8c2daaca3bc913661db2959fafdd25abe9fbec9a4fc7516332f538351ee5a"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.clie / 1fa89abe3584 / 7

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](resources--fleet--reference--group-003.md#canonical-896f93cd9da32fde87bab850656939f55bd75efe7fc25ed0fe91c03a54cdca29)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-79458c18a6ed9f5ee19d730f917caf2214339d317fdf4f668563c5dd8c201283"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e7cdb8e2cf68e1cda110980b15bedd7a0f2c7b9c7a21167d526dc3cca7153dd3"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.clie / 182be888cda5 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-26cd2b551b9e5c0a8fec323e52bc1d0eed1c17d073fde97b216c79c62b5c9a5d)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](resources--fleet--reference--group-003.md#canonical-896f93cd9da32fde87bab850656939f55bd75efe7fc25ed0fe91c03a54cdca29)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info

<a id="canonical-71fc60ac6566ea33417bc1e454d9c3d2950d49b0a7d369506473a1d0733a0fa8"></a>

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

<a id="canonical-c8c3449d1686c8de6b42d8d7067eba5cbc7d3de237312392bfabb086b672fdfb"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.clie / 182be888cda5 / 3

<a id="canonical-a6b1d09a8e954bda15df2fcb46920cf8d72db6d829a22fda1ef0920aa9533b88"></a>

<a id="canonical-a6b55a45a3cd5ae30b65a792e8f01bca95f22f75a254c7359d9416e8f87b109f"></a>

## provider_ref property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.clie / 182be888cda5 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-f324c25c32bc70eb882621ffaa334826f5fe61c78ec9fbef1e9acc3db400b2dd"></a>

<a id="canonical-d0358e678fc4246cf137ad1cbf4016601b63734ae54a966b21bfb071a788ac34"></a>

## url property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.clie / 182be888cda5 / 5

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

<a id="canonical-cf1bc3cba810afcba4ab4866b837413e87e90f55eb80f536710b4a5123368f63"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.clie / 182be888cda5 / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](resources--fleet--reference--group-003.md#canonical-896f93cd9da32fde87bab850656939f55bd75efe7fc25ed0fe91c03a54cdca29)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-70e181e855d483fe94870fab7cb06991de72402434892bbcace142a82526b5cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-66a899f909ff879b4fdbba6d65f669da39a678d781b79a7be8f389e42b9cde73"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.pass / b5dca49b180d / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-26cd2b551b9e5c0a8fec323e52bc1d0eed1c17d073fde97b216c79c62b5c9a5d)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password

<a id="canonical-b4871c70d0666f4fae3d08f57ccb40db2ebc9fc3a0f1b41ef06b70a690d4f66d"></a>

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

<a id="canonical-ad7d9158f3f4986ed0993579604b593e81d89ca5d041ab5b52019b35c73c8f69"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.pass / b5dca49b180d / 3

- [blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-37378a014a651e16d10a28f52c83b6a110d2b00996032adfedb791fb58e12b3c): complete subsection reference.

- [clear_secret_info](resources--fleet--reference--group-003.md#canonical-134c8f38a5f88e9315b91ee6b883730161be4f42fdaea1919f8ceb326e5c3dd5): complete subsection reference.

<a id="canonical-4559f0a5b45b2c005994a4098e02f90eeb7457e28226be16427796d9fd56c4e7"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.pass / b5dca49b180d / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-37378a014a651e16d10a28f52c83b6a110d2b00996032adfedb791fb58e12b3c)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info](resources--fleet--reference--group-003.md#canonical-134c8f38a5f88e9315b91ee6b883730161be4f42fdaea1919f8ceb326e5c3dd5)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-26cd2b551b9e5c0a8fec323e52bc1d0eed1c17d073fde97b216c79c62b5c9a5d)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-37378a014a651e16d10a28f52c83b6a110d2b00996032adfedb791fb58e12b3c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ec9aca8ab62291844944df1187525b272b52fca1be05118cec277ca6c4ae1a3f"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.pass / 343b10453cb6 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-26cd2b551b9e5c0a8fec323e52bc1d0eed1c17d073fde97b216c79c62b5c9a5d)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](resources--fleet--reference--group-003.md#canonical-70e181e855d483fe94870fab7cb06991de72402434892bbcace142a82526b5cf)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info

<a id="canonical-786e8083a3ba0b4ef135a9b2da0f0f93f3f6f52372acaa950cd46d44180be476"></a>

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

<a id="canonical-84b211c564d5d7648b2657b34290be0093ec07f14eb0022cf4e0eafbc7047985"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.pass / 343b10453cb6 / 3

<a id="canonical-06ab38934c2a3b8a4dc00f3dfa72d71c459a1506f5403f627bc54ff416994ec9"></a>

<a id="canonical-0e85f14aa48c6322d133786576b6e78acb934520e393e9345c68220722a43f3a"></a>

## decryption_provider property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.pass / 343b10453cb6 / 4

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

<a id="canonical-814f9e03af60ec81187d139dce23bcdada2671fc628ca20fd2418310a5ff5501"></a>

<a id="canonical-9259e2193a5ece153c183f0e75c28ee2b1757d083537bfbb1e0dd390c6e70f68"></a>

## location property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.pass / 343b10453cb6 / 5

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

<a id="canonical-79c79c7c086b7f7d4ed1c2eb14c4b4fb5f1f7a6ce6f481b075de0af3f38b0166"></a>

<a id="canonical-5e68179c30847836d9f4ef3f0bab3e01af10e6505fe73f68d99d6beccd10cd48"></a>

## store_provider property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.pass / 343b10453cb6 / 6

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

<a id="canonical-f02a8d029d529502493332bb42519e90cd8855bc9b552bf9fe34445fe80fdbeb"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.pass / 343b10453cb6 / 7

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](resources--fleet--reference--group-003.md#canonical-70e181e855d483fe94870fab7cb06991de72402434892bbcace142a82526b5cf)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-134c8f38a5f88e9315b91ee6b883730161be4f42fdaea1919f8ceb326e5c3dd5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c3018136affabfe7e2c8ec817bd470b8a368b4f701b9bffdfc655a05c475674"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.pass / 2fabb662c6ea / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-26cd2b551b9e5c0a8fec323e52bc1d0eed1c17d073fde97b216c79c62b5c9a5d)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](resources--fleet--reference--group-003.md#canonical-70e181e855d483fe94870fab7cb06991de72402434892bbcace142a82526b5cf)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info

<a id="canonical-ef11a2559c71c14679bb611dc2405a189f295385c98f156113514cf51cc97e5e"></a>

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

<a id="canonical-d69ede15581b8c95b799d2ab2fa36e336874708d51c8bc4b285d6f27c6350b94"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.pass / 2fabb662c6ea / 3

<a id="canonical-0bf129268d51881f8c86b42a2cf962184763aa2d4fa36e0b130208d15db3af8e"></a>

<a id="canonical-1a9032ce9d274cdc90d143cce6f6c5af1c2d3c1b8efe7e811d67f123763f6f03"></a>

## provider_ref property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.pass / 2fabb662c6ea / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1f8ae8099ab48390f9cbda79e5629771f1ebf930fe8e490e374876ae7e389eef"></a>

<a id="canonical-d6e54c6bd33648b315f511208ad85134b1d2ed3072a6d360757c270300255c2c"></a>

## url property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.pass / 2fabb662c6ea / 5

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

<a id="canonical-0fd0857203044c5e058544d984ceaef72a8955bb53a3411da6ced6c98d5b7689"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.pass / 2fabb662c6ea / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](resources--fleet--reference--group-003.md#canonical-70e181e855d483fe94870fab7cb06991de72402434892bbcace142a82526b5cf)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-2e977c79263e7217d825d481adf03b7e8da6d705b2ac7838cceeebf0a23931b0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da5e1a0952878f6dc6456ea0ea315037c118c3c76ae30a35348cf252df3f48ff"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / c9feedad15e0 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-26cd2b551b9e5c0a8fec323e52bc1d0eed1c17d073fde97b216c79c62b5c9a5d)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage

<a id="canonical-3bc3354cb89316f1cdf4087b9279db01682f095b0294b9dc68d5c348ecc25ff9"></a>

Type: `"object"`. list nested block, Optional.

List of Virtual Storage Pool definitions which are referred back by Storage Class label match
selection.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-29cceffffc1b4bad1ab34b0a1d5aefc9dee603847217ad383c84543785128c35"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / c9feedad15e0 / 3

<a id="canonical-d739d33ce5de26cedcf856d9db420c9c307feed679e1c654ca527399dfea4b9d"></a>

<a id="canonical-00a0644b1f7d6515ef5437fa436232300f618ad3a69f9c3800e1ed3242536c7a"></a>

## labels property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / c9feedad15e0 / 4

Type: `["map", "string"]`. Optional.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class label match
selection.

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

- [volume_defaults](resources--fleet--reference--group-003.md#canonical-97e1487191a2d087c22bc128893adbd8ebee1cbe32fca33b786a7be6863dce06): complete subsection reference.

<a id="canonical-55597840738278d18e3b1fc5757cc26ebb45f5c5ce2a18619572943bc84dacd3"></a>

<a id="canonical-2df62231c18914e07a625da500a66ef00ad4880b458aec405b6aeb2470b34a43"></a>

## zone property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / c9feedad15e0 / 5

Type: `"string"`. Optional.

Virtual Pool Zone. Virtual Storage Pool zone definition.

Upstream description:

Virtual Storage Pool zone definition.

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

<a id="canonical-46adbc2f85be65ddf885e6403d5ded3c693ac87b2496ff5b63b9234a2bcc468e"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / c9feedad15e0 / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](resources--fleet--reference--group-003.md#canonical-97e1487191a2d087c22bc128893adbd8ebee1cbe32fca33b786a7be6863dce06)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-26cd2b551b9e5c0a8fec323e52bc1d0eed1c17d073fde97b216c79c62b5c9a5d)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-97e1487191a2d087c22bc128893adbd8ebee1cbe32fca33b786a7be6863dce06"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1800755e78808627ab6c408513b445252691670244d9395de6093d476eae59d8"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 1aa45005cf1c / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-26cd2b551b9e5c0a8fec323e52bc1d0eed1c17d073fde97b216c79c62b5c9a5d)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](resources--fleet--reference--group-003.md#canonical-2e977c79263e7217d825d481adf03b7e8da6d705b2ac7838cceeebf0a23931b0)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults

<a id="canonical-0078a4bc2d30b9ada41644f0ee77f3a4d8cc7d62cb1b5dfd87c421180e3ee3e1"></a>

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

<a id="canonical-88de099a3ecb1501c69386897d9062818042d9527fcf6252ac2c3de7382a95ac"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 1aa45005cf1c / 3

<a id="canonical-5bdfa85ee14f6e2c96e4befec489622dac2f8ec5203add4a77b6bc8221efcc6d"></a>

<a id="canonical-cf5ebda67cae1a98a7d09c74a4425509494a0ed351a7129227d9d41cfc24723a"></a>

## adaptive_qos_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 1aa45005cf1c / 4

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

<a id="canonical-3c8afc757c851bdca8c443a0780a0d3db71bf3513e3573fff9836760a5d7009e"></a>

<a id="canonical-78f3b11fad8959d860b6dd2f0f9ee098b5a3b96c6be71c431018347cedd3ed11"></a>

## encryption property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 1aa45005cf1c / 5

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

<a id="canonical-97f3ab7e026153e22e1b8df7783281af045386254f8c098bc3362fb7d6f5029d"></a>

<a id="canonical-7cd5dc8e2f3a7de21d6fc67410db2de31ca4f694eb9093f0f409f3ce14d49105"></a>

## export_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 1aa45005cf1c / 6

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

- [no_qos](resources--fleet--reference--group-003.md#canonical-b79c35cc3589c3c646f2dce6b7a5095ff582b83ec71bc8564c8d7cdd84bf8c48): complete subsection reference.

<a id="canonical-d0d4f7067c985cb7208547ad687dc00edcfaca96f4c592df68c3eba3ad6a8e2e"></a>

<a id="canonical-508fb49e89334b01fd3df7235abcec879fe5da374fb5e847b8b334bbbcc90026"></a>

## qos_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 1aa45005cf1c / 7

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

<a id="canonical-c3d2be7c7eea712a4886b3c2ef0d3ac11822dbf5224790814822fa76726ff751"></a>

<a id="canonical-5ec96407fed59c1e00e1023bee81fa902744a5f632171121b368d79ed81d8112"></a>

## security_style property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 1aa45005cf1c / 8

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

<a id="canonical-78162d8797cb8db284a4173d0f6103c05cb75e40d2916b42863a629c2fd88ce7"></a>

<a id="canonical-c98ec5efdd68a58580f2e8519cdd157809fadeecf31ffd4c4922b60f8d170e2f"></a>

## snapshot_dir property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 1aa45005cf1c / 9

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

<a id="canonical-75e3ca6f19f526917fbeff30ac7aa3bcf82d7c78838514cc1a49ddfe6ea8ef1e"></a>

<a id="canonical-fd031780dfcb77335840fb018a9f643f4a3e8d975b1f20f9938ac56360baef27"></a>

## snapshot_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 1aa45005cf1c / 10

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

<a id="canonical-39479be8bfe2c67117c42f8a916553ca235058add4f33cdb8a69daad032b129c"></a>

<a id="canonical-7bd7c1993ac5d8ab585558c5f1d5fb7fb169c388f69928bc1ced05c4054e8a44"></a>

## snapshot_reserve property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 1aa45005cf1c / 11

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

<a id="canonical-49251069204ca897f242f90df21caeb1ddb7cedcd682f200828aee25e4b33e83"></a>

<a id="canonical-5000d23f7325cf94e77950a3518089eb5f82006c5a0cff1fcec2838f42854098"></a>

## space_reserve property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 1aa45005cf1c / 12

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

<a id="canonical-bb44bf1021846f3d4ab9a45e9f73b67eb9f0d5237cfb4940c649bba2672b3178"></a>

<a id="canonical-aff13959281e68e802d3ad3d335414790aa322d47a4042f816eea7cf815e2f77"></a>

## split_on_clone property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 1aa45005cf1c / 13

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

<a id="canonical-d0274ac74b5af632a48cc20acc8315806a137dcaff71f301f645cb45fec305c0"></a>

<a id="canonical-c1999c74ad1904c196ae001fc11ad49ebe343afb10a0bf7f62656a394886edc6"></a>

## tiering_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 1aa45005cf1c / 14

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

<a id="canonical-ba2141390f3bb86382ec4273bc36a61070f0fb08d299ffa932d05dbebd5ff723"></a>

<a id="canonical-d81d99774b9bef6fd1045257354f36e300fb6addd6caf9d2a6cbb225aa4e802d"></a>

## unix_permissions property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 1aa45005cf1c / 15

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

<a id="canonical-7021b1661e650d9e80fddf2a32902bd11a3c160d70dc2832ba8943fbcf148296"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 1aa45005cf1c / 16

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos](resources--fleet--reference--group-003.md#canonical-b79c35cc3589c3c646f2dce6b7a5095ff582b83ec71bc8564c8d7cdd84bf8c48)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](resources--fleet--reference--group-003.md#canonical-2e977c79263e7217d825d481adf03b7e8da6d705b2ac7838cceeebf0a23931b0)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-b79c35cc3589c3c646f2dce6b7a5095ff582b83ec71bc8564c8d7cdd84bf8c48"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-69ba75ab8670ae987e2003e0323fc9a85a6edde595bc6507c0304de47dab6358"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / a132e37f1438 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-26cd2b551b9e5c0a8fec323e52bc1d0eed1c17d073fde97b216c79c62b5c9a5d)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](resources--fleet--reference--group-003.md#canonical-2e977c79263e7217d825d481adf03b7e8da6d705b2ac7838cceeebf0a23931b0)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](resources--fleet--reference--group-003.md#canonical-97e1487191a2d087c22bc128893adbd8ebee1cbe32fca33b786a7be6863dce06)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos

<a id="canonical-55c7da2ff60eb2b87b98edbe293caecaa9499d59d0376daa6fa77071174e1d0b"></a>

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

<a id="canonical-c738b4744baa6f537c9ff6412313cc4690a9868cf077e7f61e0cfc3785f7c5c6"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / a132e37f1438 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dba24f931e1a036ed965aaa4470e5e76134336c3cc81e2812a1e3f6ecb68cd70"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / a132e37f1438 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](resources--fleet--reference--group-003.md#canonical-97e1487191a2d087c22bc128893adbd8ebee1cbe32fca33b786a7be6863dce06)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-bd864a49908d14cbed12e5febd07a0acf564c0f324249e9ca457d51c4d0939fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c3509d447b4972219a69cc10c2f750ee99ec40900f98332aedab52d55cc4b5d8"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / e888d5656b32 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-26cd2b551b9e5c0a8fec323e52bc1d0eed1c17d073fde97b216c79c62b5c9a5d)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults

<a id="canonical-2ed0708b574737bf6b846f5fb0423617980c50201854f3b3e708536803ef40c6"></a>

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

<a id="canonical-b18e42bcbebd7dfa7bdff0838a6c985fb7450a7deb548d7cc0c157e1ae615925"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / e888d5656b32 / 3

<a id="canonical-71d16e5b5345242bd08954b722dacf909d8b4e7daa18b7724d015660f072bd2f"></a>

<a id="canonical-f596d98c71ffe1051979bf6834a387849e41df152242f3ac37ea280377a9a52a"></a>

## adaptive_qos_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / e888d5656b32 / 4

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

<a id="canonical-e9a16501e4ef6c3ef2ff6621829ea98fe6796343ae08ffc4769e0174bd08353a"></a>

<a id="canonical-e08a088247064cc66cd7da8ba5e6f31fcd031ac79d87496f644e9640940da5c7"></a>

## encryption property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / e888d5656b32 / 5

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

<a id="canonical-f73b932346462d67c632f16f5ae56f145cd7dbf80cc05e1c01fda3e1b6c216f2"></a>

<a id="canonical-7bd96d088dc556d839c029d42f82d5b0dcbe9c24a5aa2c6761cfeb03032ee5cf"></a>

## export_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / e888d5656b32 / 6

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

- [no_qos](resources--fleet--reference--group-003.md#canonical-af47c74ab05288c56c60a70fb5c9fec98fb96b3bd8f596f21d5e5142085ed6ee): complete subsection reference.

<a id="canonical-7b96c18ae430eb01199dd29f3da8c376646c6b4b28e21b3b6cb8b507ed296298"></a>

<a id="canonical-28051fcd52951c0e7e494da77f4bdc9a9f7680c4f9a87513731596b32c0f8533"></a>

## qos_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / e888d5656b32 / 7

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

<a id="canonical-c0910bf8cce4982691e13cac7c45b8068325b05ba5541f021cc3df9c6b684e9a"></a>

<a id="canonical-5256e53d4233914323094183e8749c168fe0628137d25dc525844bff393ae676"></a>

## security_style property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / e888d5656b32 / 8

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

<a id="canonical-65c4049cad020625b3dc362858271b160e5307374df986ccde12c600be8f2eda"></a>

<a id="canonical-6eef033216beba5ddec5ae967d2c56f076aef286cae51c6a6d4bf673768c0151"></a>

## snapshot_dir property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / e888d5656b32 / 9

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

<a id="canonical-59d59982606260a385498353c3b2890acd7aea513cec45310dfd3f1d9b995eae"></a>

<a id="canonical-0ac150d95b8c36138c5ddd7f430633b4bcf5abe0f655bb7c7ac370e0020a8496"></a>

## snapshot_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / e888d5656b32 / 10

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

<a id="canonical-582e3534c9152938929d0e90fc92e1f4f659bb3ffede3a00e189fbdf23f89d84"></a>

<a id="canonical-c09807b15d95505312acd54261cd4ce862627e557eb1eabe725e08455813008a"></a>

## snapshot_reserve property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / e888d5656b32 / 11

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

<a id="canonical-3a81d20d82dca146f3eec85ba42f3b94c0b033cfb6bfdb324eb0dba8d1720df5"></a>

<a id="canonical-d13760f7d771fe1496c2e169c7a2defef71daac3458e27b6a4cdfd0191fe2054"></a>

## space_reserve property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / e888d5656b32 / 12

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

<a id="canonical-5ee7171993d7991c4e165fc74295bdaf3268e124a4de52e080b954fd356a66ee"></a>

<a id="canonical-7f82935e58819123a54ad3b0e4534e20ad1c15f85e33bd6379f326328ddace3b"></a>

## split_on_clone property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / e888d5656b32 / 13

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

<a id="canonical-719bf9ba926139cae109e75b987f21449e438bd1908bd5a5db15559aa69f3e25"></a>

<a id="canonical-ba0359712564efc2588cce05229fe334e55caf6e1b03d51fe45f238d701cebb2"></a>

## tiering_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / e888d5656b32 / 14

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

<a id="canonical-2b56104202699aff080896a936c0a1f7f134cfa5e6120e5ee070dc27c1e96cfc"></a>

<a id="canonical-f13cbeb271916cf203edd3e5a4c263a5ac60160d326afa7ed57c0a082f6f1e76"></a>

## unix_permissions property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / e888d5656b32 / 15

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

<a id="canonical-c1004c391f648788f61624f956ce9cc6e3a9c6adfe29352218bd121c67a5b29c"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / e888d5656b32 / 16

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos](resources--fleet--reference--group-003.md#canonical-af47c74ab05288c56c60a70fb5c9fec98fb96b3bd8f596f21d5e5142085ed6ee)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-26cd2b551b9e5c0a8fec323e52bc1d0eed1c17d073fde97b216c79c62b5c9a5d)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-af47c74ab05288c56c60a70fb5c9fec98fb96b3bd8f596f21d5e5142085ed6ee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-14add2e96557008fcc755ad260d18177c83b95396603e7090c4fb226c7f9ae01"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / 6131570544e4 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-26cd2b551b9e5c0a8fec323e52bc1d0eed1c17d073fde97b216c79c62b5c9a5d)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults](resources--fleet--reference--group-003.md#canonical-bd864a49908d14cbed12e5febd07a0acf564c0f324249e9ca457d51c4d0939fd)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos

<a id="canonical-c26f8c1e2bf49c28644cd67d9b37b9d50514a19a26d131c76fe3776f2e5959d3"></a>

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

<a id="canonical-b05ea58491955de329c904fadfd9712d961e65562541268b4f285633e955dcf4"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / 6131570544e4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-558ea74ed478caf4777dfbfdd4b5a7f443a6a7d76c3ad17925ebbac12d516874"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / 6131570544e4 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults](resources--fleet--reference--group-003.md#canonical-bd864a49908d14cbed12e5febd07a0acf564c0f324249e9ca457d51c4d0939fd)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-9aaa33d28bdc9675ddfeacd2a1549952db4fb3dee17435ce405d097f40edb5d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-137a38b36ea88da2e022f927da88555a60ae8c22280daad5c4615efc6c048a9f"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / a09bfdc4f255 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san

<a id="canonical-fb92cb2c840473dc7a79c5078151f9a867d5c7e6dce621d73b671f471bb0a0f3"></a>

Type: `"object"`. single nested block, Optional.

Configuration of storage backend for NetApp ONTAP SAN.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("storage_driver_name",
    "username"),
  validators.ConflictingObjectAttributes("data_lif_dns_name",
    "data_lif_ip"),
  validators.ConflictingObjectAttributes("management_lif_dns_name",
    "management_lif_ip"),
  validators.ConflictingObjectAttributes("no_chap",
    "use_chap")}
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
  "x-ves-oneof-field-chap_choice": "[\"no_chap\",\"use_chap\"]",
  "x-ves-oneof-field-data_lif": "[\"data_lif_dns_name\",\"data_lif_ip\"]",
  "x-ves-oneof-field-management_lif": "[\"management_lif_dns_name\",\"management_lif_ip\"]"
}
```

Terraform syntax:

```terraform
netapp_backend_ontap_san {
  # Configure direct properties listed below.
}
```

<a id="canonical-1bda9d21c5c825a19a01269667e39bdc355a81380c1d433ec4c3132ca024d635"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / a09bfdc4f255 / 3

<a id="canonical-7fa20fa352b4183980a6e68a2873244ba0b41c902e03655ebb7797849e7cf164"></a>

<a id="canonical-1bbef246eb930228abffde0beb45c21e8fcd6bb1f16d5f5dc7fecb6dbee48d36"></a>

## client_certificate property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / a09bfdc4f255 / 4

Type: `"string"`. Optional.

Please Enter Base64-encoded value of client certificate. Used for certificate-based auth.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

- [client_private_key](resources--fleet--reference--group-003.md#canonical-ae0a5fa7b77ee8ea895ff2975148f06917e2d6a92f458ca95b7fbe86115901f9): complete subsection reference.

<a id="canonical-8254d81ef56c539614e8c2750af9bec705f79769ad4836dd9ab26bc614166499"></a>

<a id="canonical-9973a13295e2ee2f114e0a37334010cfac52b7861c9d815da4313f2cdffb6410"></a>

## data_lif_dns_name property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / a09bfdc4f255 / 5

Type: `"string"`. Optional.

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

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

<a id="canonical-86547b7790f5008a789f6f6dea5fa26e94c20ad17ff58dbdf7b58d207a045e70"></a>

<a id="canonical-713f34bbc1c895696ceda5c2c8e2a8052696005260aea24c0c7f673d072b0b9a"></a>

## data_lif_ip property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / a09bfdc4f255 / 6

Type: `"string"`. Optional.

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

Upstream description:

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

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

<a id="canonical-8617a7295b347085807f5a519b54b8e97bf81bc2ab104396c9003d0095fc21ee"></a>

<a id="canonical-6017f0fe6fded40b8d49ba5d22d713914b5b6dc1947a206d1e20eb345bf6d29a"></a>

## igroup_name property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / a09bfdc4f255 / 7

Type: `"string"`. Optional.

Name of the igroup for SAN volumes to use.

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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-6f8cc426379bc7acad251cf1b49d66c50f7c9451b05a60db03616c9f11d6d76a"></a>

<a id="canonical-2fe119ed1ed914bbcf0c76b7befee1a126497e9e820d7295b87b72d790d0731a"></a>

## labels property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / a09bfdc4f255 / 8

Type: `["map", "string"]`. Optional.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class selection.

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

<a id="canonical-1de049e46800f0f2b68fe0c3343f809b40047020ec3ab60acfe7aa8747c164e0"></a>

<a id="canonical-649f660d7255681d93000973645edfa94520d27bd450a3d0d3d87a09526e803e"></a>

## limit_aggregate_usage property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / a09bfdc4f255 / 9

Type: `"number"`. Optional.

Fail provisioning if usage is above this percentage. Not enforced by default.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 100),
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
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-da961ac3c9ce86e43dc46d10d35241539fa2ade770f76a11b575213bbc0baa19"></a>

<a id="canonical-082ced53652be2edc08849759508f48b5b6c0e7306797c44c38fc3c90f0272e7"></a>

## limit_volume_size property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / a09bfdc4f255 / 10

Type: `"number"`. Optional.

Fail provisioning if requested volume size in GBi is above this value. Not enforced by default.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2748be39f761fb2b269d35b09315c8f88a3c641edae608b2a971aea83bcb7f02"></a>

<a id="canonical-e71add3bb13fba912e96990bc1d49ee1cdb4283f7359899bfd30a8637633db99"></a>

## management_lif_dns_name property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / a09bfdc4f255 / 11

Type: `"string"`. Optional.

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

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

<a id="canonical-f9cbe0ed882101add21c6839e03ef537a2a239c50de3ffcba20545f0a4aebd14"></a>

<a id="canonical-ff7a750165b5dd6d409f04dc8167b3cbe8f338d6e4f9f798a3f4fbbad96e48af"></a>

## management_lif_ip property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / a09bfdc4f255 / 12

Type: `"string"`. Optional.

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

Upstream description:

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

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

- [no_chap](resources--fleet--reference--group-003.md#canonical-8b9ad7812fbf313d7afbfb00cada1424078be68e417976be3709749daef276a6): complete subsection reference.

- [password](resources--fleet--reference--group-003.md#canonical-3e51df872867549b24ac4e0a906d38a11d8513cd29dba64d7b82ed15d8f64e1d): complete subsection reference.

<a id="canonical-d7271893588e1287d47792beb9115862286e2b1cdf4895eabfbc70fdf98a6cd4"></a>

<a id="canonical-a30480c72fa0da666250cd6e4d1523f902a646753a74e355410bf522937d3339"></a>

## region property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / a09bfdc4f255 / 13

Type: `"string"`. Optional.

Backend Region. Virtual Pool Region.

Upstream description:

Virtual Pool Region.

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

- [storage](resources--fleet--reference--group-003.md#canonical-1e6ce6cf703552081acf7928319efdd2c0f437a02e94ab57f77b03141bfdecb2): complete subsection reference.

<a id="canonical-72a044adb7d6c63616408f500d27be8f84509949b2233086fcef2a56709a8f8a"></a>

<a id="canonical-1f064f5c8713cd6c679ee9297aea67db309a9a70d2cb646b151e2b78b2c052f3"></a>

## storage_driver_name property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / a09bfdc4f255 / 14

Type: `"string"`. Optional.

\[Enum: ontap-san|ontap-san-economy|ontap-nas-flexgroup\] Storage Backend Driver. Configuration of
Backend Name. Possible values are \`ontap-san\`, \`ontap-san-economy\`, \`ontap-nas-flexgroup\`.

Upstream description:

Configuration of Backend Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ontap-san",
    "ontap-san-economy",
    "ontap-nas-flexgroup"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ontap-san",
    "ontap-san-economy",
    "ontap-nas-flexgroup"
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
    "ves.io.schema.rules.string.in": "[\\\"ontap-san\\\",\\\"ontap-san-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ontap-san\\\",\\\"ontap-san-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  }
}
```

<a id="canonical-a10bcf3f8241c0625e51f0d6f74a166942f834c7d507e72ff9b20547d0ce354d"></a>

<a id="canonical-a7380cd3b88e5425e42d241310153ccf9c5534062ce0a0597fb80d6811ef3bfd"></a>

## storage_prefix property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / a09bfdc4f255 / 15

Type: `"string"`. Optional.

Prefix used when provisioning new volumes in the SVM. Once set this cannot be updated.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 80),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 80,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 80,
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
    "ves.io.schema.rules.string.max_len": "80",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "80",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-569e8f93bd32f899e9d3a332a33f545f93a2b1adeabaa037ac8e50d68aa0833e"></a>

<a id="canonical-5ec7357ee9dc235b08ecbb61102114ebfcf1667d4544c80b421c1422205797f7"></a>

## svm property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / a09bfdc4f255 / 16

Type: `"string"`. Optional.

Storage virtual machine to use. Derived if an SVM managementLIF is specified.

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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-4c6a69734f691276469f0d8b0109a8e2f06e2c5dd00c57c7f9ddd4348a1e7d4e"></a>

<a id="canonical-579a472375980af5e84d198207e86812a075221725966125a8216e63e1a33944"></a>

## trusted_ca_certificate property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / a09bfdc4f255 / 17

Type: `"string"`. Optional.

Please Enter Base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth.

Upstream description:

Please Enter Base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth..

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

- [use_chap](resources--fleet--reference--group-003.md#canonical-24acca475feb36d716873496a71007399f5587fcaf785096f7fe139ff5078747): complete subsection reference.

<a id="canonical-4ae56b9c3509cba47c6a812a7f5473eed81b36f8dde85d01c9f00a08caef8470"></a>

<a id="canonical-ed2348674507195e34af3b86d13211bcf5e863a9e20328c3306e52c1968e40c5"></a>

## username property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / a09bfdc4f255 / 18

Type: `"string"`. Optional.

Username. Username to connect to the cluster/SVM.

Upstream description:

Username to connect to the cluster/SVM.

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
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [volume_defaults](resources--fleet--reference--group-004.md#canonical-4bc1400cfae06aa3f69593fc2f6c01a537183649f34d046d0a1fbf9411e01f29): complete subsection reference.

<a id="canonical-c1a67addc1812abc70b7f4526d743ae29ae67000fba5fff5b780f41156fb236f"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / a09bfdc4f255 / 19

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](resources--fleet--reference--group-003.md#canonical-ae0a5fa7b77ee8ea895ff2975148f06917e2d6a92f458ca95b7fbe86115901f9)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap](resources--fleet--reference--group-003.md#canonical-8b9ad7812fbf313d7afbfb00cada1424078be68e417976be3709749daef276a6)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](resources--fleet--reference--group-003.md#canonical-3e51df872867549b24ac4e0a906d38a11d8513cd29dba64d7b82ed15d8f64e1d)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](resources--fleet--reference--group-003.md#canonical-1e6ce6cf703552081acf7928319efdd2c0f437a02e94ab57f77b03141bfdecb2)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--fleet--reference--group-003.md#canonical-24acca475feb36d716873496a71007399f5587fcaf785096f7fe139ff5078747)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults](resources--fleet--reference--group-004.md#canonical-4bc1400cfae06aa3f69593fc2f6c01a537183649f34d046d0a1fbf9411e01f29)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-ae0a5fa7b77ee8ea895ff2975148f06917e2d6a92f458ca95b7fbe86115901f9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5c6ea406fae49b5162c17f59af32694162fe42b50adf3fe78bf2aa3f818ff03"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.clie / e055cb8d9f8f / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-9aaa33d28bdc9675ddfeacd2a1549952db4fb3dee17435ce405d097f40edb5d1)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key

<a id="canonical-7f2e3bd7a8359ee86d663c0b1a1c4b83688be9dd53c0f7e052b121828f764703"></a>

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
client_private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-45a8424dddaa6b9584338de9dc0248f4bf50eab1be9bfa32905911c2de98eb45"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.clie / e055cb8d9f8f / 3

- [blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-68ca9d4ea311ba79e7a24fcac7e59dd0ece652bf9f6060b4c6bd7db8ddf3832f): complete subsection reference.

- [clear_secret_info](resources--fleet--reference--group-003.md#canonical-9f87b0ac361ae6f3d6139738664894d0303ed56c220ad782fe88b903ffdd189e): complete subsection reference.

<a id="canonical-fd619abbe5b29bf7d22abd7122790d943eddb4c976035320617c303a44ea6554"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.clie / e055cb8d9f8f / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-68ca9d4ea311ba79e7a24fcac7e59dd0ece652bf9f6060b4c6bd7db8ddf3832f)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info](resources--fleet--reference--group-003.md#canonical-9f87b0ac361ae6f3d6139738664894d0303ed56c220ad782fe88b903ffdd189e)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-9aaa33d28bdc9675ddfeacd2a1549952db4fb3dee17435ce405d097f40edb5d1)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-68ca9d4ea311ba79e7a24fcac7e59dd0ece652bf9f6060b4c6bd7db8ddf3832f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1c90d1a8bd7edd81d2a61e9d9ebe8ce538dcf482c224124182e2d06d3eeef76"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.clie / a62288a81645 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-9aaa33d28bdc9675ddfeacd2a1549952db4fb3dee17435ce405d097f40edb5d1)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](resources--fleet--reference--group-003.md#canonical-ae0a5fa7b77ee8ea895ff2975148f06917e2d6a92f458ca95b7fbe86115901f9)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info

<a id="canonical-4e465f8f3be19a1f0c0f6fa0b2b8f1ad65e369eb1afe82195812ddeaa56d54db"></a>

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

<a id="canonical-5d97c090b1353e1713404a2e2255b2e1022c836f626414b327dc3b5ff8022f60"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.clie / a62288a81645 / 3

<a id="canonical-02f0003f90ee1f59c2e518f7d449b83f08037cdce226fbaf780e50496052a933"></a>

<a id="canonical-5d9a4dabceb0e2ae0468c677629cf55e7e5726e06aadb84fbbac1bd0ae125eb6"></a>

## decryption_provider property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.clie / a62288a81645 / 4

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

<a id="canonical-8e9f4aff07d5bad0d5792ce303eb5e0e8ce5631078e1cee0bdd1937dfc1f8699"></a>

<a id="canonical-5042504589d4aade0d658f813939e594025f5bb89e4fbca19f581f6ff721dd38"></a>

## location property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.clie / a62288a81645 / 5

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

<a id="canonical-c5688594c835dec8220106ad2f64043c1fd0befb9fbab8843d3137ab6d829f38"></a>

<a id="canonical-04c80fc12e373939f90d7b52425f399c9547e495edac6f4835ac4e1d1f98d5c0"></a>

## store_provider property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.clie / a62288a81645 / 6

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

<a id="canonical-cf64152f8610869fffb08049f218436904f4d6d0cf543c5613afe4cc98509b70"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.clie / a62288a81645 / 7

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](resources--fleet--reference--group-003.md#canonical-ae0a5fa7b77ee8ea895ff2975148f06917e2d6a92f458ca95b7fbe86115901f9)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-9f87b0ac361ae6f3d6139738664894d0303ed56c220ad782fe88b903ffdd189e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2547e86a673b2386351adf431bb8334048ca8f8c6c0727ba377c4521d9a3bbc4"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.clie / 5ab4bcaeeda9 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-9aaa33d28bdc9675ddfeacd2a1549952db4fb3dee17435ce405d097f40edb5d1)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](resources--fleet--reference--group-003.md#canonical-ae0a5fa7b77ee8ea895ff2975148f06917e2d6a92f458ca95b7fbe86115901f9)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info

<a id="canonical-9593d9c94c1a79871c91527b0567433933132fec10caa88fd2657b2d5002e4e4"></a>

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

<a id="canonical-33c1c776026e25c4a29eb2a56326a50c8ed4e310ab5d8e935339d38e35e55de5"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.clie / 5ab4bcaeeda9 / 3

<a id="canonical-dd4e4f63a832ceba591372f086c85231c6af39f508f49242278a211b92c92c46"></a>

<a id="canonical-4a0b3bd8416febf34078bbf5e118cfecda0c9cf937051d147e34b0807ba10742"></a>

## provider_ref property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.clie / 5ab4bcaeeda9 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-7ee4b5cab99258e737dad9a3acfdd80f73fc7e27d7202999b7b1c031729f7f65"></a>

<a id="canonical-e2e2115ac5b0c27a0611fc16d9ba89ed7b6f7bcccefba45b2ea1ab7b9a2a6a01"></a>

## url property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.clie / 5ab4bcaeeda9 / 5

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

<a id="canonical-d567f25688bb5469734e16210309adfae8c00a10725cf43c26c8bf9d8302c5c6"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.clie / 5ab4bcaeeda9 / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](resources--fleet--reference--group-003.md#canonical-ae0a5fa7b77ee8ea895ff2975148f06917e2d6a92f458ca95b7fbe86115901f9)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-8b9ad7812fbf313d7afbfb00cada1424078be68e417976be3709749daef276a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5b7fd85bb811c7abca67e84b7a4385b5435d6aaf75afd86930299158100315b"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_c / 3bdf646f83c8 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-9aaa33d28bdc9675ddfeacd2a1549952db4fb3dee17435ce405d097f40edb5d1)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap

<a id="canonical-b6b53c6ab31e2a2f864f68f532c2b9d67a55f3fe346452af7d49a86b18dc94f9"></a>

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
no_chap = {}
```

<a id="canonical-396030d4e487d6fed60cc8b6408a8f6e6114def5ddbf87dc2969d5a1157ac580"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_c / 3bdf646f83c8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-da0a62a6297cf7d100818280d2e8c7d8b2f3881e0370bf4731d7854bb84f7632"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_c / 3bdf646f83c8 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-9aaa33d28bdc9675ddfeacd2a1549952db4fb3dee17435ce405d097f40edb5d1)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-3e51df872867549b24ac4e0a906d38a11d8513cd29dba64d7b82ed15d8f64e1d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d417565ae1c0ef7964501e59bb0ce3a421444bd58e8c64750ba22ad86b9e2348"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.pass / 0a8bc84a5117 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-9aaa33d28bdc9675ddfeacd2a1549952db4fb3dee17435ce405d097f40edb5d1)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password

<a id="canonical-35a380af8d747fd76c81e10718a99d53e6df02e338286d91b8bc474b97e59f9b"></a>

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

<a id="canonical-a3e7b3d54134cb7dceec64b6a84674ac24682940069070cd79d29db942d6a56f"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.pass / 0a8bc84a5117 / 3

- [blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-33e40ca3c0437434bc7b61eb52314b9d0047c177adb0d1112ed8cdd9520af5a8): complete subsection reference.

- [clear_secret_info](resources--fleet--reference--group-003.md#canonical-a2a2d923808d495a7db6af37db3b73adcacd999b595bc6a10a11d2e00ec07450): complete subsection reference.

<a id="canonical-7d62d88d727a52599dba5ce5a0db378d6024a5b38658daad3be6cc11123dfa44"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.pass / 0a8bc84a5117 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-33e40ca3c0437434bc7b61eb52314b9d0047c177adb0d1112ed8cdd9520af5a8)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info](resources--fleet--reference--group-003.md#canonical-a2a2d923808d495a7db6af37db3b73adcacd999b595bc6a10a11d2e00ec07450)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-9aaa33d28bdc9675ddfeacd2a1549952db4fb3dee17435ce405d097f40edb5d1)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-33e40ca3c0437434bc7b61eb52314b9d0047c177adb0d1112ed8cdd9520af5a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef458b8258791d1e8ce80257e76c6251e7173fc6da0cc597778cb9c7b61296f2"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.pass / 6b259025c90d / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-9aaa33d28bdc9675ddfeacd2a1549952db4fb3dee17435ce405d097f40edb5d1)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](resources--fleet--reference--group-003.md#canonical-3e51df872867549b24ac4e0a906d38a11d8513cd29dba64d7b82ed15d8f64e1d)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info

<a id="canonical-03ea32f0c68b2cc345ea3993f4ed3eb317676c9dd7d1153c74011e2f93c0af69"></a>

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

<a id="canonical-4c185a756fc0be74d0d28ccdd7b36050ee1c9fe51f41348d4b522b728c88663e"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.pass / 6b259025c90d / 3

<a id="canonical-cce8b8d4466ec1b87b12b082f5c3176f34de2dba6958cf9b08d6254508e6d2f1"></a>

<a id="canonical-b77ea24632f2919d02ee33c95f8f6ca01510292b90568f8bdfde0383dae087db"></a>

## decryption_provider property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.pass / 6b259025c90d / 4

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

<a id="canonical-db0d8f6e8fc6b5563931f5cdad986e8cc147245b533e4b2386b3b6ec348ffb2b"></a>

<a id="canonical-23efd2c90deda380d1408ba107317ed44027ed4502dbf6abab8412143658fe30"></a>

## location property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.pass / 6b259025c90d / 5

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

<a id="canonical-7075762e2a7a4336cb1b8f31d545cc114a899287e7394f69058fcad70cc666f6"></a>

<a id="canonical-f811954376316d4b7be44272dbe06667098f3776fc8dfbc45dbe09d0b0abd5c5"></a>

## store_provider property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.pass / 6b259025c90d / 6

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

<a id="canonical-5948606a6fee57c52707045ddbbf8a70f947d9846db1184f2c40b63fed8f9cd3"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.pass / 6b259025c90d / 7

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](resources--fleet--reference--group-003.md#canonical-3e51df872867549b24ac4e0a906d38a11d8513cd29dba64d7b82ed15d8f64e1d)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-a2a2d923808d495a7db6af37db3b73adcacd999b595bc6a10a11d2e00ec07450"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2396b57d2d1aa4c2b159de9db3acb8f92af5559b665cb3e4dcc147a193f6259"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.pass / c6a9599eccee / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-9aaa33d28bdc9675ddfeacd2a1549952db4fb3dee17435ce405d097f40edb5d1)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](resources--fleet--reference--group-003.md#canonical-3e51df872867549b24ac4e0a906d38a11d8513cd29dba64d7b82ed15d8f64e1d)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info

<a id="canonical-ed2f6d04154a908e994c1980bf71e60e6b04e76bbcb28be9c9004635833298d1"></a>

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

<a id="canonical-c341d323772122a20657481a421ab9e915480f21337e580cdf351098b461a5dd"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.pass / c6a9599eccee / 3

<a id="canonical-1ca35c55e10544bbf5c2a5f2ba780ad8bc546d78ec0f958671eb7a67adfee2b3"></a>

<a id="canonical-9eab3ab40b46670a2fe87bfde443c4904781d767db648d44b2470e7337d3b8fc"></a>

## provider_ref property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.pass / c6a9599eccee / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-01b13ae1d3de6bb7f3891fcd7ef7e0d8ad46ff31be4b1883d7ac7f8d3b0c5cb7"></a>

<a id="canonical-c836dea1c5764da2b73e4b7bd42a5505764f6520a1c925aaab76160c1bac5e00"></a>

## url property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.pass / c6a9599eccee / 5

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

<a id="canonical-9f163541ab43f0c115dee0c0ce9deaa936a1620c0bc8888cf69a4907a14a10c9"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.pass / c6a9599eccee / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](resources--fleet--reference--group-003.md#canonical-3e51df872867549b24ac4e0a906d38a11d8513cd29dba64d7b82ed15d8f64e1d)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-1e6ce6cf703552081acf7928319efdd2c0f437a02e94ab57f77b03141bfdecb2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae335d69d5e19e5d31a55863cbf1c1f019cbb57925fdb11bd31582d696a9ad6c"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / c6a8c7b780d0 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-9aaa33d28bdc9675ddfeacd2a1549952db4fb3dee17435ce405d097f40edb5d1)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage

<a id="canonical-25394e5df60f715f6ff4cc7af635f7ed27b365140e19fce2b931bcd1bfb5d3b1"></a>

Type: `"object"`. list nested block, Optional.

List of Virtual Storage Pool definitions which are referred back by Storage Class label match
selection.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-95ebab59743aba14ea51f97d7e1442287c0a8ceace8681a78b6330551f8cbcd8"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / c6a8c7b780d0 / 3

<a id="canonical-4deecf41515815c9c3711721e649fd54e5f47901ddb77a35c5ec0b208dbcb522"></a>

<a id="canonical-f149375b54c3f5c67ae7b3ee1e5db3bb858bcf925f140f81e370189962dce69b"></a>

## labels property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / c6a8c7b780d0 / 4

Type: `["map", "string"]`. Optional.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class label match
selection.

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

- [volume_defaults](resources--fleet--reference--group-003.md#canonical-7d8e13e8cf40ba078410b8eaa4f26637b48ab5968fb3c058e8a5902964629aab): complete subsection reference.

<a id="canonical-abc1c05b48a5c1bca98d7d65943e57314ae6cc444e4311420c9d392168314b7a"></a>

<a id="canonical-d94bdc3e74b1604b4257059b6b47a76a2e4aa2e8887895e1748f3641e397c944"></a>

## zone property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / c6a8c7b780d0 / 5

Type: `"string"`. Optional.

Virtual Pool Zone. Virtual Storage Pool zone definition.

Upstream description:

Virtual Storage Pool zone definition.

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

<a id="canonical-071d29205e518370be609060de42e8d269fe491f6d9e75f31c9ee6833c0196aa"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / c6a8c7b780d0 / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults](resources--fleet--reference--group-003.md#canonical-7d8e13e8cf40ba078410b8eaa4f26637b48ab5968fb3c058e8a5902964629aab)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-9aaa33d28bdc9675ddfeacd2a1549952db4fb3dee17435ce405d097f40edb5d1)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-7d8e13e8cf40ba078410b8eaa4f26637b48ab5968fb3c058e8a5902964629aab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7704a79a5764fc026853e6810f98f3de4fcea5e6288a43265202c52c0fc691a4"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / 5edec47145c5 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-9aaa33d28bdc9675ddfeacd2a1549952db4fb3dee17435ce405d097f40edb5d1)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](resources--fleet--reference--group-003.md#canonical-1e6ce6cf703552081acf7928319efdd2c0f437a02e94ab57f77b03141bfdecb2)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults

<a id="canonical-b129aef4f91fbb18b3d0bac918e60bdc8163204d8417f315879b6d66c0dc02f1"></a>

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

<a id="canonical-87fbb20fdb66651767534da98f8c35890bd713a6a36c75917344f2567d6d3070"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / 5edec47145c5 / 3

<a id="canonical-4a982acda9588c6473c125c048cda204d00d21844dbf1777543ca2839cdd1ce1"></a>

<a id="canonical-c9de632bd99e7cb6d0318a5c981aba9169dc37d1124cbd8bfe99d68cc414e270"></a>

## adaptive_qos_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / 5edec47145c5 / 4

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

<a id="canonical-3cf0aea4f1440c8e9e2e43d12b54e3314cf27d91ae277bf8bc4892d4c9acb76a"></a>

<a id="canonical-09e81ff90af4abe54aa4b19ae7e5fc9935e10e8d7ac650e6708dff02aed20a0e"></a>

## encryption property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / 5edec47145c5 / 5

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

<a id="canonical-a725885683457b3d339d2826c2ff7a475b9ad8869a88b6acbd5842e564b174dd"></a>

<a id="canonical-8f88c3266aa5a67776f6fa57903966349b5d577b8dec32cb1dd421c878f3b425"></a>

## export_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / 5edec47145c5 / 6

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

- [no_qos](resources--fleet--reference--group-003.md#canonical-2c783bc57281160ef2f8e7c6dece1bb26ba5b0c79127dda17e74210c86e88598): complete subsection reference.

<a id="canonical-c7d2db1427c1f70a8517f88d1b1de7c66bfca65a64c3cbd13b5ca330b6177cf9"></a>

<a id="canonical-56da9a4270abd843b6874754d20d25900b2511f4ca65d852dee11bde75ea6d3d"></a>

## qos_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / 5edec47145c5 / 7

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

<a id="canonical-7446e92740aaa75e1f67b7330fbc567b247ca84ba1eb2ba2b3742465a641b8f2"></a>

<a id="canonical-020a91f2f252daa554995fa4607c20f03589f0a6f1f283d37acb46a24cdb7d8a"></a>

## security_style property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / 5edec47145c5 / 8

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

<a id="canonical-7372b58488d6c3c6448329b55e76b5c4400b59432cffd2ef4493925df936925c"></a>

<a id="canonical-14ab45ca33d033d92c48c0cad83f6a5a2bc8e366e230bf2aa864a0046d23f0d3"></a>

## snapshot_dir property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / 5edec47145c5 / 9

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

<a id="canonical-dad9603d873884b6f16355693b6a70803d7ae69b83abca5dbfdf59ded4e5a2fb"></a>

<a id="canonical-ab3657f2f6e4427f3ac22c33681e828711a61595cea8658150fdc69d44a1eaf5"></a>

## snapshot_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / 5edec47145c5 / 10

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

<a id="canonical-54480bd16b1a91306ade296f4de31003e7f772b3da8778414b5b3222a942113e"></a>

<a id="canonical-96f2af307f92cde8bf4f5800c5b0583a6188628db5c073777c4215630f9f46b6"></a>

## snapshot_reserve property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / 5edec47145c5 / 11

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

<a id="canonical-d64ef82c379dc843a26e9c5dd9ac498de1bff1c39bae62e3155f63f146853dec"></a>

<a id="canonical-5108e20208d9f87be7118cafa55e987ec3736f83b4771051b38397155ae44136"></a>

## space_reserve property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / 5edec47145c5 / 12

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

<a id="canonical-53c423e68aab45909dd1a9f14ee26b9e28a2ad6b1e596de59e27494563ab2dd1"></a>

<a id="canonical-8ad19a14662b2908e1207d36e6b3904bd56ecd650929cacb2d1c5953018d2963"></a>

## split_on_clone property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / 5edec47145c5 / 13

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

<a id="canonical-78a7bad3541bc51b4b71e0ffab4a6aad9012a0e17e9e3aa4622ce63e633cc32c"></a>

<a id="canonical-c6aa1c566f04989433bf7fc6df9fb2b9e0b226fc948037e5e153527971b84e3e"></a>

## tiering_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / 5edec47145c5 / 14

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

<a id="canonical-46240e6f79210a9881fdf8f7edfa921470de22481193c329bbd9059995f86111"></a>

<a id="canonical-2aa7069498b805791fb1df7fd7801fb784836dcf6a043ffec3538830ea873991"></a>

## unix_permissions property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / 5edec47145c5 / 15

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

<a id="canonical-8ea9f9abc9f8474e2b0024ab5f1f597eb8f59d25b76ce527dc608d7f99f17373"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / 5edec47145c5 / 16

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos](resources--fleet--reference--group-003.md#canonical-2c783bc57281160ef2f8e7c6dece1bb26ba5b0c79127dda17e74210c86e88598)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](resources--fleet--reference--group-003.md#canonical-1e6ce6cf703552081acf7928319efdd2c0f437a02e94ab57f77b03141bfdecb2)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-2c783bc57281160ef2f8e7c6dece1bb26ba5b0c79127dda17e74210c86e88598"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d00c844f5d7f352e8929675a5e9de5eea5daf66568e4d94a3ce1da45ac43f0c"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / 6de61e6757f5 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-9aaa33d28bdc9675ddfeacd2a1549952db4fb3dee17435ce405d097f40edb5d1)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](resources--fleet--reference--group-003.md#canonical-1e6ce6cf703552081acf7928319efdd2c0f437a02e94ab57f77b03141bfdecb2)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults](resources--fleet--reference--group-003.md#canonical-7d8e13e8cf40ba078410b8eaa4f26637b48ab5968fb3c058e8a5902964629aab)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos

<a id="canonical-5f0e9aec333a5358cac577ba8cbf0408f367836d7bbf56051b2ed70c55b8d847"></a>

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

<a id="canonical-db25dc5e05b184a8f45ea9a1804ec7c745006fa5858a25821d4aedfce97c68c4"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / 6de61e6757f5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ae1d0098034e62fbcf85c1f5b489e368b7937725d273e9fa78d312cfd723c52e"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / 6de61e6757f5 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults](resources--fleet--reference--group-003.md#canonical-7d8e13e8cf40ba078410b8eaa4f26637b48ab5968fb3c058e8a5902964629aab)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-24acca475feb36d716873496a71007399f5587fcaf785096f7fe139ff5078747"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6ef372a5d4e7359783b1e932141c3d3356c003807ee6f923889190e731b36ff"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 973608a5ec9f / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-9aaa33d28bdc9675ddfeacd2a1549952db4fb3dee17435ce405d097f40edb5d1)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap

<a id="canonical-5f3abbadb3a45112ee9f9ccf03e1c9596744c75fd5d2f41e2237d2b45eb18cf3"></a>

Type: `"object"`. single nested block, Optional.

Device NetApp Backend ONTAP SAN CHAP configuration OPTIONS for enabled CHAP.

Receipt-pinned upstream constraints:

```json
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
use_chap {
  # Configure direct properties listed below.
}
```

<a id="canonical-32b728919d47a3563debc5c02aa7301bdccd51217d6a2c86847696454510f56a"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 973608a5ec9f / 3

- [chap_initiator_secret](resources--fleet--reference--group-003.md#canonical-0f3a32e9ecaf3607d8a0c80c0aca752905d1ff1569244a6c0945843902352f88): complete subsection reference.

- [chap_target_initiator_secret](resources--fleet--reference--group-003.md#canonical-32cee025bc0732a76806a9f726b842427d21551a8d82acf9b5358db06b14163d): complete subsection reference.

<a id="canonical-931f1b9029367784dba11298a35ceea356dde33ad60fd80d263795f0d2e4d147"></a>

<a id="canonical-9e36ab4c9b51488fb762f7cbba9e9d3fe7d17b96037a8902e1e319bc499ab426"></a>

## chap_target_username property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 973608a5ec9f / 4

Type: `"string"`. Optional.

Target username. Required if useCHAP=true.

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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-90e4b8d2ac29c597ca36ebfaf8873f1f812e647a178bbc57d9bce68ad4f20587"></a>

<a id="canonical-7afbad8694f1eb7e312a642c2f56d266cce83027db3a62fcf344643ce89184a3"></a>

## chap_username property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 973608a5ec9f / 5

Type: `"string"`. Optional.

Inbound username. Required if useCHAP=true.

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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-35e4e3cdb1cc1478880e78eac04be4583db540eaf5a3b4286b05e8d755310c92"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 973608a5ec9f / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](resources--fleet--reference--group-003.md#canonical-0f3a32e9ecaf3607d8a0c80c0aca752905d1ff1569244a6c0945843902352f88)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](resources--fleet--reference--group-003.md#canonical-32cee025bc0732a76806a9f726b842427d21551a8d82acf9b5358db06b14163d)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-9aaa33d28bdc9675ddfeacd2a1549952db4fb3dee17435ce405d097f40edb5d1)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-0f3a32e9ecaf3607d8a0c80c0aca752905d1ff1569244a6c0945843902352f88"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-60305879d1ac984432e39c70ba4c54fe20e7bd607c348c2d9bad62f501eba39b"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / d81af83aba1d / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-9aaa33d28bdc9675ddfeacd2a1549952db4fb3dee17435ce405d097f40edb5d1)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--fleet--reference--group-003.md#canonical-24acca475feb36d716873496a71007399f5587fcaf785096f7fe139ff5078747)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret

<a id="canonical-c392202bd46d57b924068fee5808a2b08dc2b54534973eecffd971bdf09182cc"></a>

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
chap_initiator_secret {
  # Configure direct properties listed below.
}
```

<a id="canonical-6338e029a541871348697004b4504f9891e3e46b4cae650ecf924e970728c868"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / d81af83aba1d / 3

- [blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-752b49e93f76de02003e891e934b070fc5f205e55f3ed30fc791905e61834178): complete subsection reference.

- [clear_secret_info](resources--fleet--reference--group-003.md#canonical-720d0afa971373f6176d2b2aaa9a3eced1d4b1fc27bab7204ed5900592b345d4): complete subsection reference.

<a id="canonical-b71197a52509bd6c670b78e9975a6ace6edf2549f0a08d432d780df5e874149b"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / d81af83aba1d / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-752b49e93f76de02003e891e934b070fc5f205e55f3ed30fc791905e61834178)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info](resources--fleet--reference--group-003.md#canonical-720d0afa971373f6176d2b2aaa9a3eced1d4b1fc27bab7204ed5900592b345d4)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--fleet--reference--group-003.md#canonical-24acca475feb36d716873496a71007399f5587fcaf785096f7fe139ff5078747)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-752b49e93f76de02003e891e934b070fc5f205e55f3ed30fc791905e61834178"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee8430dca12745ac7fd4fdaa648147c1e41ca4b94ecf2f3c75b2b8c114233421"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 9dc89453a243 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-9aaa33d28bdc9675ddfeacd2a1549952db4fb3dee17435ce405d097f40edb5d1)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--fleet--reference--group-003.md#canonical-24acca475feb36d716873496a71007399f5587fcaf785096f7fe139ff5078747)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](resources--fleet--reference--group-003.md#canonical-0f3a32e9ecaf3607d8a0c80c0aca752905d1ff1569244a6c0945843902352f88)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info

<a id="canonical-97b9c2182b9303b5dea6a0b31ee86c201ded45426172a572e3fbeb0af813a60c"></a>

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

<a id="canonical-d7d8530e2e06835e88bf0c0a67f15ee4510b2390912e19144f4fdea077146e17"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 9dc89453a243 / 3

<a id="canonical-ce943248688dc31e2d61a516cd1b230fad3d38386553ac2af8df905c94716b1e"></a>

<a id="canonical-0abfddf55be19486fc5b0488ffd0aaee2bf3450dbd5792be2dda4a0c11cef49b"></a>

## decryption_provider property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 9dc89453a243 / 4

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

<a id="canonical-3f36e2c8b0a7615df5d3abfcf76f08a068e988fae0e59ff8e616d4280c2850b0"></a>

<a id="canonical-5559521edbb73ee054820cd78a15db7aeffff134bfd0fefa0ef55406135bc867"></a>

## location property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 9dc89453a243 / 5

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

<a id="canonical-5646bd9b3540df70d0c31e1ab9de7da2c76f7066a932a429dfb9bbf2a002f5db"></a>

<a id="canonical-3b9f8d7c207904f19909112f0c24593fec9bdc748140a5b0d9583c1a8b125073"></a>

## store_provider property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 9dc89453a243 / 6

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

<a id="canonical-440a87448e2116c20806e317090bf5b91be698400da7d180f9326ef9b7fb3326"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 9dc89453a243 / 7

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](resources--fleet--reference--group-003.md#canonical-0f3a32e9ecaf3607d8a0c80c0aca752905d1ff1569244a6c0945843902352f88)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-720d0afa971373f6176d2b2aaa9a3eced1d4b1fc27bab7204ed5900592b345d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e09c898c83425fe6aa5130d60818649f3b95890cfd50df3a43cb2dcd1d50a90"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 3b3d80a6dd37 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-59d0c2e4b77dacefecf90993c426cbbaeb92f39be02e0a19d6a2e3d9c831d369)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-a9610e3c935176d058f02c806103b3c7ef7234cd2d1fc5d26351cb372eb305cb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-9aaa33d28bdc9675ddfeacd2a1549952db4fb3dee17435ce405d097f40edb5d1)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--fleet--reference--group-003.md#canonical-24acca475feb36d716873496a71007399f5587fcaf785096f7fe139ff5078747)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](resources--fleet--reference--group-003.md#canonical-0f3a32e9ecaf3607d8a0c80c0aca752905d1ff1569244a6c0945843902352f88)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info

<a id="canonical-99185e34e3cbae7af8ec088ae4eeb00af3d7f0da2c1fd5f5fba03527d33a3250"></a>

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

<a id="canonical-8aa87907f3da5ad7c87c208264a2285e7cb08e18539690ecc3443d8b40c66d84"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 3b3d80a6dd37 / 3

<a id="canonical-d3129e5024257b6ab33e8da601bfdc09b2016dd0f8ef41a19781aa9432404109"></a>

<a id="canonical-f8fbeaccc32a22a7c1bc3f451abe40c79ab022e3b014a6494ce222e782dca0e1"></a>

## provider_ref property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 3b3d80a6dd37 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-4df249798204d772b3ec08f27eca44cc9991011422be846643a53995189787b7"></a>

<a id="canonical-f5b9b1b4a31bee1d0583e928001def3543a04d18128269478efb134a0cdbcd91"></a>

## url property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 3b3d80a6dd37 / 5

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

<a id="canonical-113b746596e9297a1818e014aabd065176a0200a663ac1348ae9792075290376"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 3b3d80a6dd37 / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](resources--fleet--reference--group-003.md#canonical-0f3a32e9ecaf3607d8a0c80c0aca752905d1ff1569244a6c0945843902352f88)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-32cee025bc0732a76806a9f726b842427d21551a8d82acf9b5358db06b14163d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
