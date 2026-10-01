---
page_title: "xcsh_secret_management_access reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_secret_management_access reference."
---

# xcsh_secret_management_access reference

<a id="canonical-3ffc03f97d600f4a4fc71be74fa45e147e36be525ef32d45111184b16479662e"></a>

## Direct properties — access_info.vault_auth_info / 9d354a40a6e0 / 3

- [app_role_auth](resources--secret_management_access--reference--group-002.md#canonical-0c399bcedfa9c28fc8b3bf62e5a0c54f106d7a9abd4ef28580642bc37f31e3ad): complete subsection reference.

- [token](resources--secret_management_access--reference--group-002.md#canonical-ed3248dc3b3a53b6488a7625991f7bf44d4f1f8ebd78e95a0b73869048702e11): complete subsection reference.

<a id="canonical-1d877c777217449d1d00f28e6c6da53ddb598c3a87d1f695ae4995ba42980bb7"></a>

## Next pages — access_info.vault_auth_info / 9d354a40a6e0 / 4

- [access_info.vault_auth_info.app_role_auth](resources--secret_management_access--reference--group-002.md#canonical-0c399bcedfa9c28fc8b3bf62e5a0c54f106d7a9abd4ef28580642bc37f31e3ad)
- [access_info.vault_auth_info.token](resources--secret_management_access--reference--group-002.md#canonical-ed3248dc3b3a53b6488a7625991f7bf44d4f1f8ebd78e95a0b73869048702e11)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-0c399bcedfa9c28fc8b3bf62e5a0c54f106d7a9abd4ef28580642bc37f31e3ad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd60d9bca768102529070822cb482822e46601a1ecb4c08e84309f729d4ed5fe"></a>

## access_info.vault_auth_info.app_role_auth — access_info.vault_auth_info.app_role_auth / ec29f5a040c8 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.vault_auth_info](resources--secret_management_access--reference--group-001.md#canonical-b0918fbb2da2dea4ce096b94bf3e5be045867e7018c64c823efb2cddcebd7036)
- access_info.vault_auth_info.app_role_auth

<a id="canonical-fb3ead27a67aacb3b0c34726b9acfc595c1f56efd8df8c65f47474ddb8f959bc"></a>

Type: `"object"`. single nested block, Optional.

AppRoleAuthInfoType contains parameters for AppRole authentication in Hashicorp Vault.

Receipt-pinned upstream constraints:

```json
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
app_role_auth {
  # Configure direct properties listed below.
}
```

<a id="canonical-3e2915c75349af8e3c26aa7c3bf2e3469453a6d56f0462f1a4dcaca21e698b36"></a>

## Direct properties — access_info.vault_auth_info.app_role_auth / ec29f5a040c8 / 3

<a id="canonical-85b6bcb752bc7104ceee9e7458358ec887a8f617f99c42880b7213af01c84e28"></a>

<a id="canonical-ace458e492bde28e602c4c67ac424b91eb1938f1a3c8a710202997d6e8aaa42c"></a>

## role_id property — access_info.vault_auth_info.app_role_auth / ec29f5a040c8 / 4

Type: `"string"`. Optional.

Role ID. Role-ID to be used for authentication.

Upstream description:

Role-ID to be used for authentication.

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

- [secret_id](resources--secret_management_access--reference--group-002.md#canonical-a3605f516b489aa07f97aeec043d1408c9958aa21d6d7b9810b5757f50db59a6): complete subsection reference.

<a id="canonical-fa08b78686c7e8369af9b2ca906ae1a60f5ad67e2e60ba3a900c21f55d72337d"></a>

## Next pages — access_info.vault_auth_info.app_role_auth / ec29f5a040c8 / 5

- [access_info.vault_auth_info.app_role_auth.secret_id](resources--secret_management_access--reference--group-002.md#canonical-a3605f516b489aa07f97aeec043d1408c9958aa21d6d7b9810b5757f50db59a6)
- [access_info.vault_auth_info](resources--secret_management_access--reference--group-001.md#canonical-b0918fbb2da2dea4ce096b94bf3e5be045867e7018c64c823efb2cddcebd7036)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-a3605f516b489aa07f97aeec043d1408c9958aa21d6d7b9810b5757f50db59a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-03ea67da79390029576a2756fc484a20aa8ed25a0152a9c20e8d39d8d422171b"></a>

## access_info.vault_auth_info.app_role_auth.secret_id — access_info.vault_auth_info.app_role_auth.secret_id / 032e408cbc3e / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.vault_auth_info](resources--secret_management_access--reference--group-001.md#canonical-b0918fbb2da2dea4ce096b94bf3e5be045867e7018c64c823efb2cddcebd7036)
- [access_info.vault_auth_info.app_role_auth](resources--secret_management_access--reference--group-002.md#canonical-0c399bcedfa9c28fc8b3bf62e5a0c54f106d7a9abd4ef28580642bc37f31e3ad)
- access_info.vault_auth_info.app_role_auth.secret_id

<a id="canonical-e9c33f4f50d30f02575861665eae5ddd457063888e4b7a78f47b21840288e053"></a>

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
secret_id {
  # Configure direct properties listed below.
}
```

<a id="canonical-021571141c9ac0e4b81ce2c062262267936a44b6fcb508b86a30ef99dc3158e7"></a>

## Direct properties — access_info.vault_auth_info.app_role_auth.secret_id / 032e408cbc3e / 3

- [blindfold_secret_info](resources--secret_management_access--reference--group-002.md#canonical-c61132937135530f1dc4b7c7aa47a70d097a61029c6280fce5dc9b9c26eaa394): complete subsection reference.

- [clear_secret_info](resources--secret_management_access--reference--group-002.md#canonical-cbb865a8e60d0156db5e9810815e5ef0c889535297ed1941414866e18b813d53): complete subsection reference.

<a id="canonical-86df938abe495ba283b7a12873970a9009ef97f58b68cbe9200084e40b826ea1"></a>

## Next pages — access_info.vault_auth_info.app_role_auth.secret_id / 032e408cbc3e / 4

- [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info](resources--secret_management_access--reference--group-002.md#canonical-c61132937135530f1dc4b7c7aa47a70d097a61029c6280fce5dc9b9c26eaa394)
- [access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info](resources--secret_management_access--reference--group-002.md#canonical-cbb865a8e60d0156db5e9810815e5ef0c889535297ed1941414866e18b813d53)
- [access_info.vault_auth_info.app_role_auth](resources--secret_management_access--reference--group-002.md#canonical-0c399bcedfa9c28fc8b3bf62e5a0c54f106d7a9abd4ef28580642bc37f31e3ad)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-c61132937135530f1dc4b7c7aa47a70d097a61029c6280fce5dc9b9c26eaa394"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb3471a36f2b6eed74a1da770e10b6f51ce2d2055394333934cfa8c044919b8c"></a>

## access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info — access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info / 578990481773 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.vault_auth_info](resources--secret_management_access--reference--group-001.md#canonical-b0918fbb2da2dea4ce096b94bf3e5be045867e7018c64c823efb2cddcebd7036)
- [access_info.vault_auth_info.app_role_auth](resources--secret_management_access--reference--group-002.md#canonical-0c399bcedfa9c28fc8b3bf62e5a0c54f106d7a9abd4ef28580642bc37f31e3ad)
- [access_info.vault_auth_info.app_role_auth.secret_id](resources--secret_management_access--reference--group-002.md#canonical-a3605f516b489aa07f97aeec043d1408c9958aa21d6d7b9810b5757f50db59a6)
- access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info

<a id="canonical-aacfe2c02732ef1123fa273595f3ae58e4672a6e34830287a9cdf8d5b053f020"></a>

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

<a id="canonical-7b69d1ddb18dfc919f2a455795f53483a929901b4e7557414cd44632b530a74a"></a>

## Direct properties — access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info / 578990481773 / 3

<a id="canonical-538657ea8ae7497acb1a678931c0bfc69767046d5aeabd3fd1868e3dba4dc188"></a>

<a id="canonical-a653ed5af9f8564118038592d85f36ddbce8d60d5876ed5c432fb88c844f8cd6"></a>

## decryption_provider property — access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info / 578990481773 / 4

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

<a id="canonical-10ad5833b6b9f6d9153c2040872b362e24ff147f44d52e35654e0a721caef6af"></a>

<a id="canonical-e9c186c128b884b0606994ecd3cfec0a1043bf699e3f61d87822db3de13469c1"></a>

## location property — access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info / 578990481773 / 5

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

<a id="canonical-c04c03f4b68928ea71ad06b1398cfb48dbb1b2287e4320bb3144be30e8ea76a1"></a>

<a id="canonical-dda07182a1bdef5cdc9bb48044d05b33de51489ff223af859754789829e917b4"></a>

## store_provider property — access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info / 578990481773 / 6

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

<a id="canonical-2f607bdbb6a43693ca602f993b008d96ad4d974ffb4a3f320e5fc4c8b8d25701"></a>

## Next pages — access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info / 578990481773 / 7

- [access_info.vault_auth_info.app_role_auth.secret_id](resources--secret_management_access--reference--group-002.md#canonical-a3605f516b489aa07f97aeec043d1408c9958aa21d6d7b9810b5757f50db59a6)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-cbb865a8e60d0156db5e9810815e5ef0c889535297ed1941414866e18b813d53"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d860e6991f961e7e02bf05f894ada6b2f0cdb3505adc41d94a7051fab9bcaf1b"></a>

## access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info — access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info / ecd6f80e0528 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.vault_auth_info](resources--secret_management_access--reference--group-001.md#canonical-b0918fbb2da2dea4ce096b94bf3e5be045867e7018c64c823efb2cddcebd7036)
- [access_info.vault_auth_info.app_role_auth](resources--secret_management_access--reference--group-002.md#canonical-0c399bcedfa9c28fc8b3bf62e5a0c54f106d7a9abd4ef28580642bc37f31e3ad)
- [access_info.vault_auth_info.app_role_auth.secret_id](resources--secret_management_access--reference--group-002.md#canonical-a3605f516b489aa07f97aeec043d1408c9958aa21d6d7b9810b5757f50db59a6)
- access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info

<a id="canonical-f19bcb05cc33222b41cf1ee07f8440cfe69a51a14e0ebd6ebaa57bb2f9c70af4"></a>

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

<a id="canonical-ca28cee5751e02dcdd71281fc03e66ed043ca884ebf888272ac9fabc4dac5486"></a>

## Direct properties — access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info / ecd6f80e0528 / 3

<a id="canonical-8546c5faaaf62dd3b434b095352da960fab9e410e96b338578c4b66d8c2a721a"></a>

<a id="canonical-438de27617cec204289958d806adaa84e400324ce913241af432fe5f2a270786"></a>

## provider_ref property — access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info / ecd6f80e0528 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-6805217a6507363df82fcb5ab041655aa4b8189500d0252f1d8ef6aef018b884"></a>

<a id="canonical-688ef102274d262c62e065cfbfd3dcef61d0f358b87f3422d87728605a036af7"></a>

## url property — access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info / ecd6f80e0528 / 5

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

<a id="canonical-2c5ef055bb815c32de35cd3ab811dec2f032830c3be8b13ed2d0f5edb1183d71"></a>

## Next pages — access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info / ecd6f80e0528 / 6

- [access_info.vault_auth_info.app_role_auth.secret_id](resources--secret_management_access--reference--group-002.md#canonical-a3605f516b489aa07f97aeec043d1408c9958aa21d6d7b9810b5757f50db59a6)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-ed3248dc3b3a53b6488a7625991f7bf44d4f1f8ebd78e95a0b73869048702e11"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd10b4bc599b8b75c75b3fe30e0ba0039ff11f46b863c347c482ea3d1809049a"></a>

## access_info.vault_auth_info.token — access_info.vault_auth_info.token / b1f85395014f / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.vault_auth_info](resources--secret_management_access--reference--group-001.md#canonical-b0918fbb2da2dea4ce096b94bf3e5be045867e7018c64c823efb2cddcebd7036)
- access_info.vault_auth_info.token

<a id="canonical-d8e82bc326f4510bf3290e45954a4547faaaebb36a346c30d11e344867b5f48a"></a>

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
token {
  # Configure direct properties listed below.
}
```

<a id="canonical-43ba957451cc29dffe5fc7d21f576de914734f4f499a1ff52e8f74a5daeaf487"></a>

## Direct properties — access_info.vault_auth_info.token / b1f85395014f / 3

- [blindfold_secret_info](resources--secret_management_access--reference--group-002.md#canonical-7c8891d71b5918e918e219349b14b1122e0d8f467605c2f822fffdab99dbdeb2): complete subsection reference.

- [clear_secret_info](resources--secret_management_access--reference--group-002.md#canonical-e458f684ee4f266bb2068c7b5f5337689eb766b45a2e6c4075753ec49311a26d): complete subsection reference.

<a id="canonical-2db141f106ba6735b351a1f659f1abc9e1871015e3925152a83f576d84f5f178"></a>

## Next pages — access_info.vault_auth_info.token / b1f85395014f / 4

- [access_info.vault_auth_info.token.blindfold_secret_info](resources--secret_management_access--reference--group-002.md#canonical-7c8891d71b5918e918e219349b14b1122e0d8f467605c2f822fffdab99dbdeb2)
- [access_info.vault_auth_info.token.clear_secret_info](resources--secret_management_access--reference--group-002.md#canonical-e458f684ee4f266bb2068c7b5f5337689eb766b45a2e6c4075753ec49311a26d)
- [access_info.vault_auth_info](resources--secret_management_access--reference--group-001.md#canonical-b0918fbb2da2dea4ce096b94bf3e5be045867e7018c64c823efb2cddcebd7036)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-7c8891d71b5918e918e219349b14b1122e0d8f467605c2f822fffdab99dbdeb2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-beae6aacf256fe7a08eae2309db61c69439df2965e248fcc54c844b7c4f0fb16"></a>

## access_info.vault_auth_info.token.blindfold_secret_info — access_info.vault_auth_info.token.blindfold_secret_info / 471cc090e81d / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.vault_auth_info](resources--secret_management_access--reference--group-001.md#canonical-b0918fbb2da2dea4ce096b94bf3e5be045867e7018c64c823efb2cddcebd7036)
- [access_info.vault_auth_info.token](resources--secret_management_access--reference--group-002.md#canonical-ed3248dc3b3a53b6488a7625991f7bf44d4f1f8ebd78e95a0b73869048702e11)
- access_info.vault_auth_info.token.blindfold_secret_info

<a id="canonical-0a3efc6d4d122e1fdc7ccae9683f73ca8e441e0c986d45bf0296490c657dd93a"></a>

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

<a id="canonical-2ed39bc85dd754059768f95bfa77f4baa7948c138b3515d6717c4f409e1ae74c"></a>

## Direct properties — access_info.vault_auth_info.token.blindfold_secret_info / 471cc090e81d / 3

<a id="canonical-a00536a4e7c9f02917f1c021cdfa04b5834da3391d7a9962958757c131496858"></a>

<a id="canonical-df03b11f365a114ae6bda9d44faee005ba8d54b9a439185ec5054cd160b08ff5"></a>

## decryption_provider property — access_info.vault_auth_info.token.blindfold_secret_info / 471cc090e81d / 4

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

<a id="canonical-82a124d87f4b6522b4847ef4ee70c8334c8a1ff74738019c15b07d29d5604244"></a>

<a id="canonical-41555528e79276ca435f1d67e9e905f8fa83351395079596909fb896174f3815"></a>

## location property — access_info.vault_auth_info.token.blindfold_secret_info / 471cc090e81d / 5

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

<a id="canonical-8e362a17e552e114780c7295409684f0909a677645b9f5f55cdec4e6d0edf2d3"></a>

<a id="canonical-2c08c76a0a839d70e85608e2ac39e75ce5a3452c6e6e02f48c11a68ddc48cc12"></a>

## store_provider property — access_info.vault_auth_info.token.blindfold_secret_info / 471cc090e81d / 6

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

<a id="canonical-f508ea4cfabcfccad4d6e0faae4e5e7d302040e985d16bce839f8d23ed452c65"></a>

## Next pages — access_info.vault_auth_info.token.blindfold_secret_info / 471cc090e81d / 7

- [access_info.vault_auth_info.token](resources--secret_management_access--reference--group-002.md#canonical-ed3248dc3b3a53b6488a7625991f7bf44d4f1f8ebd78e95a0b73869048702e11)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-e458f684ee4f266bb2068c7b5f5337689eb766b45a2e6c4075753ec49311a26d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-67eddfd08c846cb8bdebe5e6b8296a810f7d5256a6587cf6c1b0a177aa2193fd"></a>

## access_info.vault_auth_info.token.clear_secret_info — access_info.vault_auth_info.token.clear_secret_info / ff6cd6cd0b04 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.vault_auth_info](resources--secret_management_access--reference--group-001.md#canonical-b0918fbb2da2dea4ce096b94bf3e5be045867e7018c64c823efb2cddcebd7036)
- [access_info.vault_auth_info.token](resources--secret_management_access--reference--group-002.md#canonical-ed3248dc3b3a53b6488a7625991f7bf44d4f1f8ebd78e95a0b73869048702e11)
- access_info.vault_auth_info.token.clear_secret_info

<a id="canonical-eae2c87414c22374b521d9f6defee758cdcb8e31a008e9d3a51c35de3c4ca447"></a>

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

<a id="canonical-3d105396c8de2585d31786f65857f25154ceb97848d1ed60f7902c1352c54cfa"></a>

## Direct properties — access_info.vault_auth_info.token.clear_secret_info / ff6cd6cd0b04 / 3

<a id="canonical-4d7ef84491e3ef2270afa80ba1371b2863a1184d168033a7353b076d1df93a03"></a>

<a id="canonical-fc29b95f5f0476fd547f5fff4b86761ce7b327d34e6cba6922971be875bf11cc"></a>

## provider_ref property — access_info.vault_auth_info.token.clear_secret_info / ff6cd6cd0b04 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-d2c4e1409d0add568c498d0889eeb248799a6e14c7f43eca4307264a77371d4a"></a>

<a id="canonical-f96cebe9b95825e2aace236fe667241cf2aa66939ca641b936ad1397fa5f9fb3"></a>

## url property — access_info.vault_auth_info.token.clear_secret_info / ff6cd6cd0b04 / 5

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

<a id="canonical-7eeebb74e061c6f5eb38d42cfb715df9b43700883164b6faf36989203f669301"></a>

## Next pages — access_info.vault_auth_info.token.clear_secret_info / ff6cd6cd0b04 / 6

- [access_info.vault_auth_info.token](resources--secret_management_access--reference--group-002.md#canonical-ed3248dc3b3a53b6488a7625991f7bf44d4f1f8ebd78e95a0b73869048702e11)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-fb1d71681608bb290a6f1b56353da27906bf45ff58b32d0cf25c35930489c6f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9552b0162f5dfceb702894a8c0fa1a8cd222a5307839d636d6fe6eb918db6a07"></a>

## timeouts — timeouts / 91f12adbab1e / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- timeouts

<a id="canonical-ef69aa41b68a9ca1441273379d94edb35f6ca400b9dd4678bfab5efaa9272e0e"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-9ecf512b620f6d15682710faa3646259a2f49aaca2724f0e74187868c25ba26c"></a>

## Direct properties — timeouts / 91f12adbab1e / 3

<a id="canonical-97998da037a0860825a4d16d8398dba4b9a0f490073824ab4c96e0085c894a4e"></a>

<a id="canonical-460df7e1aca3144918f15c58cf35ec0f788c4b979dc54e67a82d59501f6e86a9"></a>

## create property — timeouts / 91f12adbab1e / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-5b75b6826f6f6b4f250546b62c364cfd50a2d754c8c27b822d00e326eacdf755"></a>

<a id="canonical-3975431431b889a2a2ec725f990dc73d05be8d96715e68287095604bd4e7a1a2"></a>

## delete property — timeouts / 91f12adbab1e / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-421d70ecc20ac1a050cf9da3b212ea80aa12e29cd8368f89daff3dbb932349ad"></a>

<a id="canonical-42d278c61bc0851121fba59fcfbc075fa879852cb20f2d5ad135577a431e5a65"></a>

## read property — timeouts / 91f12adbab1e / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-4ed5ba7cfc689cb36676cc70132b19dd220d9f901c2b769aa1ce891d21b77f40"></a>

<a id="canonical-15507e9a9efc1d6396d091dd9d9b247d63140a728f4bc20f74cf1046f108668f"></a>

## update property — timeouts / 91f12adbab1e / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3bcc6d177098e34d4c90d58f81f83701b901727af8fa8319091b76255987e5d2"></a>

## Next pages — timeouts / 91f12adbab1e / 8

- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-c4087d0e67a27ee79be7055e4721f539676ce06f807dbea1aea339f3cc0006a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1855cd0d61fcfb1966e852fba7e541a6539ea14ab6e8fefa6aa432b58561d09b"></a>

## where — where / 6c9fb4526f1b / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- where

<a id="canonical-aa22c1ca5e2ea26d49e03389a142c04d8cf47423488e4259b50904237a3821e9"></a>

Type: `"object"`. single nested block, Optional.

NetworkSiteRefSelector defines a union of reference to site or reference to virtual\_network or
reference to virtual\_site It is used to determine virtual network using following rules \* Direct
reference to virtual\_network object \* Site local network when referring to site object \* All site
local..

Upstream description:

NetworkSiteRefSelector defines a union of reference to site or reference to virtual\_network or
reference to virtual\_site It is used to determine virtual network using following rules \* Direct
reference to virtual\_network object \* Site local network when referring to site object \* All site
local networks for sites selected by referring to virtual\_site object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_network"),
  validators.ConflictingObjectAttributes("site",
    "virtual_site"),
  validators.ConflictingObjectAttributes("virtual_network",
    "virtual_site")}
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
  "x-ves-oneof-field-ref_or_selector": "[\"site\",\"virtual_network\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
where {
  # Configure direct properties listed below.
}
```

<a id="canonical-0cf08c7ac0b94c77db826d55487510f32b09c9cae286f9bffce392ba98e75342"></a>

## Direct properties — where / 6c9fb4526f1b / 3

- [site](resources--secret_management_access--reference--group-002.md#canonical-cf7474552b182d794f8efe76684c67465a25d74024974c5857381e3847fadef3): complete subsection reference.

- [virtual_network](resources--secret_management_access--reference--group-002.md#canonical-26537b26dd4d680f415020d60d64d9ae473329585f3fb9d5e42d216f825d18dc): complete subsection reference.

- [virtual_site](resources--secret_management_access--reference--group-002.md#canonical-b3a44f1ec541d360e42ab4b0fc76e3d51f113b1a64fe42ef7e943b97fe75d27c): complete subsection reference.

<a id="canonical-8da2f5c0c8597ee1bea84b1c301696aa4f67ca540131eb4f3869536742253d75"></a>

## Next pages — where / 6c9fb4526f1b / 4

- [where.site](resources--secret_management_access--reference--group-002.md#canonical-cf7474552b182d794f8efe76684c67465a25d74024974c5857381e3847fadef3)
- [where.virtual_network](resources--secret_management_access--reference--group-002.md#canonical-26537b26dd4d680f415020d60d64d9ae473329585f3fb9d5e42d216f825d18dc)
- [where.virtual_site](resources--secret_management_access--reference--group-002.md#canonical-b3a44f1ec541d360e42ab4b0fc76e3d51f113b1a64fe42ef7e943b97fe75d27c)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-cf7474552b182d794f8efe76684c67465a25d74024974c5857381e3847fadef3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-05c674c60d73b097321ed11feacee71bdc41555e69b1fab4332cedff564fa0af"></a>

## where.site — where.site / b720136a015f / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [where](resources--secret_management_access--reference--group-002.md#canonical-c4087d0e67a27ee79be7055e4721f539676ce06f807dbea1aea339f3cc0006a5)
- where.site

<a id="canonical-a91c31a37e1617dbb3b8c3f85dc683c5355e51e011d7f46635ff64e730708d70"></a>

Type: `"object"`. single nested block, Optional.

Specifies a direct reference to a site configuration object.

Upstream description:

This specifies a direct reference to a site configuration object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ref"),
  validators.ConflictingObjectAttributes("disable_internet_vip",
    "enable_internet_vip")}
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
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

Terraform syntax:

```terraform
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-7e1817ef1e0065b3da64164d1dbc4e2fe055283cdb4657c65fadb3c3dffb7b80"></a>

## Direct properties — where.site / b720136a015f / 3

- [disable_internet_vip](resources--secret_management_access--reference--group-002.md#canonical-fa25b6587d47990311a7d5b73b7d9a3d464944ac450922891205a4d6764edcd0): complete subsection reference.

- [enable_internet_vip](resources--secret_management_access--reference--group-002.md#canonical-1c39014fcd98a29c027ceda599305e450c698172ca517723c012b66a2a6ba109): complete subsection reference.

<a id="canonical-9f1e7cb7b86518ae26b1fb8e2280d3752343d67f9bc7879a96c7b15654e53bd4"></a>

<a id="canonical-9868141892898ea5b357c9cc0b53af155190f8e7af52dd101fe5c286f34e4dc8"></a>

## network_type property — where.site / b720136a015f / 4

Type: `"string"`. Optional.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Upstream description:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created automatically and present on all sites
Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE is a private network inside site. It
is a secure network and is not connected to public network. Virtual-network of this type is local to
every site. Two virtual networks of this type on different sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created during provisioning of site User defined per-site
virtual network. Scope of this virtual network is limited to the site. This is not yet supported
Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC directly connects to the public internet.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on RE sites only It is an internally created by the system. They must not be created by
user Virtual Networks with global scope across different sites in F5XC domain. An example global
virtual-network called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [ref](resources--secret_management_access--reference--group-002.md#canonical-d1102b8971071c52ddb35c7f7966abab034eb7f192669342d8749e83138b9839): complete subsection reference.

<a id="canonical-33e8592dc34e410b69df7ceeec395d4cd07170032548f956bd57f01b5fa08982"></a>

## Next pages — where.site / b720136a015f / 5

- [where.site.disable_internet_vip](resources--secret_management_access--reference--group-002.md#canonical-fa25b6587d47990311a7d5b73b7d9a3d464944ac450922891205a4d6764edcd0)
- [where.site.enable_internet_vip](resources--secret_management_access--reference--group-002.md#canonical-1c39014fcd98a29c027ceda599305e450c698172ca517723c012b66a2a6ba109)
- [where.site.ref](resources--secret_management_access--reference--group-002.md#canonical-d1102b8971071c52ddb35c7f7966abab034eb7f192669342d8749e83138b9839)
- [where](resources--secret_management_access--reference--group-002.md#canonical-c4087d0e67a27ee79be7055e4721f539676ce06f807dbea1aea339f3cc0006a5)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-fa25b6587d47990311a7d5b73b7d9a3d464944ac450922891205a4d6764edcd0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf62e7a41843a5c49ad5193cd83706f9870ec30cd6a8ee5e2535399855751543"></a>

## where.site.disable_internet_vip — where.site.disable_internet_vip / 1cbfdbcc7d51 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [where](resources--secret_management_access--reference--group-002.md#canonical-c4087d0e67a27ee79be7055e4721f539676ce06f807dbea1aea339f3cc0006a5)
- [where.site](resources--secret_management_access--reference--group-002.md#canonical-cf7474552b182d794f8efe76684c67465a25d74024974c5857381e3847fadef3)
- where.site.disable_internet_vip

<a id="canonical-65ad50e31a8a5d63611bb82eb2a3a4d3bd6fbe30b6b31e8316c159016486b008"></a>

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
disable_internet_vip = {}
```

<a id="canonical-ee8bb7af5030bba1e4d109a429b9e96c81e7eb38556e98f0ee4d6cf574d93614"></a>

## Direct properties — where.site.disable_internet_vip / 1cbfdbcc7d51 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b09ecb47d4592265eb1c58fd72e4f434856df9abd57c07be43704ce781ffdc81"></a>

## Next pages — where.site.disable_internet_vip / 1cbfdbcc7d51 / 4

- [where.site](resources--secret_management_access--reference--group-002.md#canonical-cf7474552b182d794f8efe76684c67465a25d74024974c5857381e3847fadef3)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-1c39014fcd98a29c027ceda599305e450c698172ca517723c012b66a2a6ba109"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7aa80646224c32e9753c86a9adc3458fcce205b4728eda1e0b3e5d676ef396db"></a>

## where.site.enable_internet_vip — where.site.enable_internet_vip / b70715230228 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [where](resources--secret_management_access--reference--group-002.md#canonical-c4087d0e67a27ee79be7055e4721f539676ce06f807dbea1aea339f3cc0006a5)
- [where.site](resources--secret_management_access--reference--group-002.md#canonical-cf7474552b182d794f8efe76684c67465a25d74024974c5857381e3847fadef3)
- where.site.enable_internet_vip

<a id="canonical-90d4b9a452bacae9a408edc49d766f0692b955b6ff96feeaf40f778355ede91a"></a>

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
enable_internet_vip = {}
```

<a id="canonical-aca08f98856fbf0df9c0a9f787f3642512fc723842af28cc458d6115435e7c87"></a>

## Direct properties — where.site.enable_internet_vip / b70715230228 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e1b081ea62d5a809028da776a6c984f8cac686c3a9cd33b80dc1c82cc08c4f9e"></a>

## Next pages — where.site.enable_internet_vip / b70715230228 / 4

- [where.site](resources--secret_management_access--reference--group-002.md#canonical-cf7474552b182d794f8efe76684c67465a25d74024974c5857381e3847fadef3)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-d1102b8971071c52ddb35c7f7966abab034eb7f192669342d8749e83138b9839"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-408de787862ff34c2cc34fb884924b49ee18bcec24b99206cecaa04682085cb7"></a>

## where.site.ref — where.site.ref / 7f5d378a9b16 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [where](resources--secret_management_access--reference--group-002.md#canonical-c4087d0e67a27ee79be7055e4721f539676ce06f807dbea1aea339f3cc0006a5)
- [where.site](resources--secret_management_access--reference--group-002.md#canonical-cf7474552b182d794f8efe76684c67465a25d74024974c5857381e3847fadef3)
- where.site.ref

<a id="canonical-781de395688e6ea87d58da1430e2d24d6314196af79947f05ab54fbb7e8e3eab"></a>

Type: `"object"`. list nested block, Optional.

Reference. A site direct reference.

Upstream description:

A site direct reference.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-d7725820a64e0c2027893ee2060eb53691c2955a2e6f91de23dd7c10cd048859"></a>

## Direct properties — where.site.ref / 7f5d378a9b16 / 3

<a id="canonical-a5205831a5bb57a998304ef0346d9a987ee789eef80f2050b3b0731aa9d2ca86"></a>

<a id="canonical-540afd893f4d8711e12fd2c0f34a6693e30801524cf0652eee1fbb4188df7c59"></a>

## kind property — where.site.ref / 7f5d378a9b16 / 4

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

<a id="canonical-ba97ecbc4c4b3e5423302d54832a9cd3361bcd972626b3666002e298ba18cf04"></a>

<a id="canonical-702b8d235da05e0ae39b149dc0a976efe2a0ca635967e2374c2a561e0ea29df6"></a>

## name property — where.site.ref / 7f5d378a9b16 / 5

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

<a id="canonical-456cb3738c69680cbae04767549c5f9ec460a78c531cd0e8d744c6dc47aa0d11"></a>

<a id="canonical-fde7a50a8ca44c7384315e7ec007d19770ddb72b23e245feac4cf37042feb56c"></a>

## namespace property — where.site.ref / 7f5d378a9b16 / 6

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

<a id="canonical-055c56cda9fbaf752a0d8c93b4675ec2719b9141f47a400790a74330dda19d28"></a>

<a id="canonical-c64774109d115e244bd44a436ec63f14c1acf066108c95db0f5fb0ec5dabd442"></a>

## tenant property — where.site.ref / 7f5d378a9b16 / 7

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

<a id="canonical-ec508bed7c6581375fd763033992674d525d95ab12d0d3b79ed882aa61e0cdb6"></a>

<a id="canonical-eee390e5766629dcf08f8ea64744b660e0e4821ac90eb5df39fd3fbf15ac76a2"></a>

## uid property — where.site.ref / 7f5d378a9b16 / 8

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

<a id="canonical-c7be2ef5d5c715a4ca440132057d2afb0ffee103f98acf505347a72ab727f8d4"></a>

## Next pages — where.site.ref / 7f5d378a9b16 / 9

- [where.site](resources--secret_management_access--reference--group-002.md#canonical-cf7474552b182d794f8efe76684c67465a25d74024974c5857381e3847fadef3)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-26537b26dd4d680f415020d60d64d9ae473329585f3fb9d5e42d216f825d18dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a5952395fe49582d1029994a0aaf7b9e73a8861f3e186764f6b81df6acfd793f"></a>

## where.virtual_network — where.virtual_network / f4c3a6c0d342 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [where](resources--secret_management_access--reference--group-002.md#canonical-c4087d0e67a27ee79be7055e4721f539676ce06f807dbea1aea339f3cc0006a5)
- where.virtual_network

<a id="canonical-7cfeab1ec7b549a6591ff824c08cbdcb2e70c5ec7bbfdf328ffde63fed41d6a7"></a>

Type: `"object"`. single nested block, Optional.

Specifies a direct reference to a network configuration object.

Upstream description:

This specifies a direct reference to a network configuration object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ref")}
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
virtual_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-187869b2e0f89d1b46d697d58a6982245ea100bc3ca7cd34292bb533caff879c"></a>

## Direct properties — where.virtual_network / f4c3a6c0d342 / 3

- [ref](resources--secret_management_access--reference--group-002.md#canonical-7ae2eb39a940bdd63a13e6b09ed2eac17bcbe90c58e9c024e6c56d208aa0c839): complete subsection reference.

<a id="canonical-5d3dfa19352f3f7a46ec7776021a7fc7254ff37aceded76be12a4259c1f415a0"></a>

## Next pages — where.virtual_network / f4c3a6c0d342 / 4

- [where.virtual_network.ref](resources--secret_management_access--reference--group-002.md#canonical-7ae2eb39a940bdd63a13e6b09ed2eac17bcbe90c58e9c024e6c56d208aa0c839)
- [where](resources--secret_management_access--reference--group-002.md#canonical-c4087d0e67a27ee79be7055e4721f539676ce06f807dbea1aea339f3cc0006a5)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-7ae2eb39a940bdd63a13e6b09ed2eac17bcbe90c58e9c024e6c56d208aa0c839"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9a9256f030141217ff5b5bb687b0fa69ffedfebd0e2a121ff25c0cc5dd807a0"></a>

## where.virtual_network.ref — where.virtual_network.ref / 4c59148b71ad / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [where](resources--secret_management_access--reference--group-002.md#canonical-c4087d0e67a27ee79be7055e4721f539676ce06f807dbea1aea339f3cc0006a5)
- [where.virtual_network](resources--secret_management_access--reference--group-002.md#canonical-26537b26dd4d680f415020d60d64d9ae473329585f3fb9d5e42d216f825d18dc)
- where.virtual_network.ref

<a id="canonical-a1e04129c99788db43360efa4d0d0ee086aab7de20f5b6bee6204b93f2e254ee"></a>

Type: `"object"`. list nested block, Optional.

Reference. A virtual network direct reference.

Upstream description:

A virtual network direct reference.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-728820b4d5b3fcdcd585f7a4056f9be980e962fd813d870b793dbb3098637ba3"></a>

## Direct properties — where.virtual_network.ref / 4c59148b71ad / 3

<a id="canonical-37fdea1cb03d29401e0527e4a9b8a4292f7f635f7ac79eff0575aea492b37b27"></a>

<a id="canonical-7d5dc640efee171bd116c4c62cef3309a11b247a0f4e5f7806931e62e8b20c81"></a>

## kind property — where.virtual_network.ref / 4c59148b71ad / 4

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

<a id="canonical-e84aac6c4ef7d9e6ef9e0346b398b59b0db2a735687b63e0c055bacf90c12a58"></a>

<a id="canonical-cf55f4a85e592604920421c77a9c0cba723b050912d8fe6cae9ff6e8c481924f"></a>

## name property — where.virtual_network.ref / 4c59148b71ad / 5

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

<a id="canonical-8ed0f4e57cbdde7438800c364643bd467366d89223d0af7d17621d9c407afcf9"></a>

<a id="canonical-d34fe7e695d60d671a125b884769bc7519c8ae6d347b8c4fa88058b609b300eb"></a>

## namespace property — where.virtual_network.ref / 4c59148b71ad / 6

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

<a id="canonical-a7f9d0891a5f490cf7209ab9dd53710c9fb7cd2acdc2ff72e81ea7edf3620aea"></a>

<a id="canonical-8ade3b884ddcad2b2823309a8deb0cb8c5ec05168f9c1ac23b6ce818a80097fc"></a>

## tenant property — where.virtual_network.ref / 4c59148b71ad / 7

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

<a id="canonical-560f10c7ddbdfe437cbe7832ca59f4bd791974a5659f1f0094152df2d83ae3ce"></a>

<a id="canonical-1d08e63c4a87f1772724a14593db8d71250f6ddc4f0c4e7d92b7c9d097939a3b"></a>

## uid property — where.virtual_network.ref / 4c59148b71ad / 8

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

<a id="canonical-0d9888943454593e3c8036ce24182500413d2d362e915c5c1a034ff9e0966559"></a>

## Next pages — where.virtual_network.ref / 4c59148b71ad / 9

- [where.virtual_network](resources--secret_management_access--reference--group-002.md#canonical-26537b26dd4d680f415020d60d64d9ae473329585f3fb9d5e42d216f825d18dc)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-b3a44f1ec541d360e42ab4b0fc76e3d51f113b1a64fe42ef7e943b97fe75d27c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b3f339e17419bb5afcc23efc8726bd3850e94cfbf1b2a515e87979455a20a028"></a>

## where.virtual_site — where.virtual_site / 15bd7dc8a58a / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [where](resources--secret_management_access--reference--group-002.md#canonical-c4087d0e67a27ee79be7055e4721f539676ce06f807dbea1aea339f3cc0006a5)
- where.virtual_site

<a id="canonical-dc5c57afbd23d619fd55105573448b74aa46a5e6aab529b71fd3e4cd346c7419"></a>

Type: `"object"`. single nested block, Optional.

Virtual Site. A reference to virtual\_site object.

Upstream description:

A reference to virtual\_site object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ref"),
  validators.ConflictingObjectAttributes("disable_internet_vip",
    "enable_internet_vip")}
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
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

Terraform syntax:

```terraform
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-78e705468ce506254d44b6dfd168d74019395715866bbe4aa9ddb8a4242d6bda"></a>

## Direct properties — where.virtual_site / 15bd7dc8a58a / 3

- [disable_internet_vip](resources--secret_management_access--reference--group-002.md#canonical-1495da0c76a9ffd12d17b24841392aec7d88ca2abc80c08dc0fe6a9d4b6e4b42): complete subsection reference.

- [enable_internet_vip](resources--secret_management_access--reference--group-002.md#canonical-481df73eece0e3ab0b67d64f3cc42366dff5a8e23d52ba3c469b834ba9380afe): complete subsection reference.

<a id="canonical-49054d9f1ae56e265831169149b9619a2eb355024fd48d6e1b916871d06aeb6f"></a>

<a id="canonical-9d344feebc0bc6a07e9d94425958caf99b54735ea540e3736011ea78faa6bf7a"></a>

## network_type property — where.virtual_site / 15bd7dc8a58a / 4

Type: `"string"`. Optional.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Upstream description:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created automatically and present on all sites
Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE is a private network inside site. It
is a secure network and is not connected to public network. Virtual-network of this type is local to
every site. Two virtual networks of this type on different sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created during provisioning of site User defined per-site
virtual network. Scope of this virtual network is limited to the site. This is not yet supported
Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC directly connects to the public internet.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on RE sites only It is an internally created by the system. They must not be created by
user Virtual Networks with global scope across different sites in F5XC domain. An example global
virtual-network called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [ref](resources--secret_management_access--reference--group-002.md#canonical-8147ea3979e4900be6e640e63ac3ef46cdb07ae396fb944fb9f66d4914354e70): complete subsection reference.

<a id="canonical-f301839b4de4e10d09df070e6378f7e1498111675a120952d8530e011544b1e5"></a>

## Next pages — where.virtual_site / 15bd7dc8a58a / 5

- [where.virtual_site.disable_internet_vip](resources--secret_management_access--reference--group-002.md#canonical-1495da0c76a9ffd12d17b24841392aec7d88ca2abc80c08dc0fe6a9d4b6e4b42)
- [where.virtual_site.enable_internet_vip](resources--secret_management_access--reference--group-002.md#canonical-481df73eece0e3ab0b67d64f3cc42366dff5a8e23d52ba3c469b834ba9380afe)
- [where.virtual_site.ref](resources--secret_management_access--reference--group-002.md#canonical-8147ea3979e4900be6e640e63ac3ef46cdb07ae396fb944fb9f66d4914354e70)
- [where](resources--secret_management_access--reference--group-002.md#canonical-c4087d0e67a27ee79be7055e4721f539676ce06f807dbea1aea339f3cc0006a5)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-1495da0c76a9ffd12d17b24841392aec7d88ca2abc80c08dc0fe6a9d4b6e4b42"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-145232929e28d8fbfaa57afbc860f37c3a61637b7acdc3c8570df17c9e8677d0"></a>

## where.virtual_site.disable_internet_vip — where.virtual_site.disable_internet_vip / 678c490a4dac / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [where](resources--secret_management_access--reference--group-002.md#canonical-c4087d0e67a27ee79be7055e4721f539676ce06f807dbea1aea339f3cc0006a5)
- [where.virtual_site](resources--secret_management_access--reference--group-002.md#canonical-b3a44f1ec541d360e42ab4b0fc76e3d51f113b1a64fe42ef7e943b97fe75d27c)
- where.virtual_site.disable_internet_vip

<a id="canonical-0c43e402f2667359ac97fe6b56aa56da32f2f81372946a51f50a27871d3c94fa"></a>

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
disable_internet_vip = {}
```

<a id="canonical-18e46fd0bd1c1cab2dd949ba83ed7ec68da6b23c5e17d71b2224205ec72694b2"></a>

## Direct properties — where.virtual_site.disable_internet_vip / 678c490a4dac / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4fbbdf1c09724081d37866eeb6ae0f0f3a05eddabd3e05becd8088e771dcb7dc"></a>

## Next pages — where.virtual_site.disable_internet_vip / 678c490a4dac / 4

- [where.virtual_site](resources--secret_management_access--reference--group-002.md#canonical-b3a44f1ec541d360e42ab4b0fc76e3d51f113b1a64fe42ef7e943b97fe75d27c)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-481df73eece0e3ab0b67d64f3cc42366dff5a8e23d52ba3c469b834ba9380afe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b0532efc70a3183c3a19d00dd019209e12c2bd8a4bffbcfa7f3d8698a951055"></a>

## where.virtual_site.enable_internet_vip — where.virtual_site.enable_internet_vip / abd56735fe73 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [where](resources--secret_management_access--reference--group-002.md#canonical-c4087d0e67a27ee79be7055e4721f539676ce06f807dbea1aea339f3cc0006a5)
- [where.virtual_site](resources--secret_management_access--reference--group-002.md#canonical-b3a44f1ec541d360e42ab4b0fc76e3d51f113b1a64fe42ef7e943b97fe75d27c)
- where.virtual_site.enable_internet_vip

<a id="canonical-77a7c5e3800f0688a8707a1ec1d219a517d8dd6725ed1a365e99a0dc1b0b38f3"></a>

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
enable_internet_vip = {}
```

<a id="canonical-8470415725158155c181240374fdffe5f318fb686661193eb983b369a92c3087"></a>

## Direct properties — where.virtual_site.enable_internet_vip / abd56735fe73 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-743e4b0ae2f83b4db62811d5ae8d60ccb81c35a43cb44fed7023b8c9eb78dae7"></a>

## Next pages — where.virtual_site.enable_internet_vip / abd56735fe73 / 4

- [where.virtual_site](resources--secret_management_access--reference--group-002.md#canonical-b3a44f1ec541d360e42ab4b0fc76e3d51f113b1a64fe42ef7e943b97fe75d27c)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-8147ea3979e4900be6e640e63ac3ef46cdb07ae396fb944fb9f66d4914354e70"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3951dd7cf296930254d23130b45c8f96aed316ebdaf7784558247757d7e931fd"></a>

## where.virtual_site.ref — where.virtual_site.ref / 198a629c1218 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [where](resources--secret_management_access--reference--group-002.md#canonical-c4087d0e67a27ee79be7055e4721f539676ce06f807dbea1aea339f3cc0006a5)
- [where.virtual_site](resources--secret_management_access--reference--group-002.md#canonical-b3a44f1ec541d360e42ab4b0fc76e3d51f113b1a64fe42ef7e943b97fe75d27c)
- where.virtual_site.ref

<a id="canonical-2ba82298eb47ed20fd17458f3d1a54689dd8d1944eb81efa4729fd3d3a1eb9e3"></a>

Type: `"object"`. list nested block, Optional.

Reference. A virtual\_site direct reference.

Upstream description:

A virtual\_site direct reference.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-39710864acdeb65eba52e7b414e9f12ed051368e9a3a56727cdb03b2abc7b6ac"></a>

## Direct properties — where.virtual_site.ref / 198a629c1218 / 3

<a id="canonical-22e223960b110165d6cdadff9ad29405d4df000479d48942dbc036c692bd7b08"></a>

<a id="canonical-178c605c6b8cb7db4fbd24423125945d66e3509a95b7d14523de2e65647e70f7"></a>

## kind property — where.virtual_site.ref / 198a629c1218 / 4

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

<a id="canonical-cb1efbad658971c542efeb7a52dfe4dd141a7d58ecc5d13e5f3ac559423886a4"></a>

<a id="canonical-504168d9d7dd4954e6075fa8ddb43bff62e5fc6b71c31c7c87424cf15616d6ee"></a>

## name property — where.virtual_site.ref / 198a629c1218 / 5

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

<a id="canonical-ce099a80cac9ebe68a3f786648301f722775087f9c364abd5eea78dddc06c3a3"></a>

<a id="canonical-a99a0daf62609010c812533bf15831976ec014660aa2dbbe3d198de6f782bcce"></a>

## namespace property — where.virtual_site.ref / 198a629c1218 / 6

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

<a id="canonical-23c49f598afbca47a0a43aca08709aaeaa2ee07fbf35432606560eeb9a1b4155"></a>

<a id="canonical-cd1fbd4c774c4cdb710195b8340d034f1ff95911b2a2c638fff6743a8b46f9f0"></a>

## tenant property — where.virtual_site.ref / 198a629c1218 / 7

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

<a id="canonical-6e3627b4f479d2a8d458e3f34f353aba0b07538ea3e78f8ff07f4ce7bdff5cb3"></a>

<a id="canonical-9b8aa15e143d371080a00e59604454b4101bce93f4e733ce20f24ecbc96f5e04"></a>

## uid property — where.virtual_site.ref / 198a629c1218 / 8

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

<a id="canonical-3cd4bd0da4439829d9ea670bead30f815113ee1eda4053780ed5edcc44fa2971"></a>

## Next pages — where.virtual_site.ref / 198a629c1218 / 9

- [where.virtual_site](resources--secret_management_access--reference--group-002.md#canonical-b3a44f1ec541d360e42ab4b0fc76e3d51f113b1a64fe42ef7e943b97fe75d27c)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
