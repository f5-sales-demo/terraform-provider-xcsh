---
page_title: "xcsh_container_registry reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_container_registry reference."
---

# xcsh_container_registry reference

<a id="canonical-e58992eaf4724a8cf0c7f0fa641dc1f519eedb53d24397f6ae8cf4129ef62c19"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-599c8c65cbd7da1b9c34a088e83c9ae85e91f0f02ea1d5c99f856d54cff40221"></a>

## Property reference — Property reference / cbb069cb9767 / 2

Breadcrumbs:

- [xcsh_container_registry](../data-sources/container_registry.md#canonical-17f6d44ebd4531a50d36aec8be1db646843481c3c17219cc6a239a366b16525f)
- Property reference

<a id="canonical-e460cfabbebfd5e18879975f45f54a4e8168755afed42f53013baa7625484bf8"></a>

## Direct properties — Property reference / cbb069cb9767 / 3

<a id="canonical-40f96c0bbe0a4aa605b94ca12038890f1c48748a1d30efc11ef399f7b6aa5251"></a>

<a id="canonical-40ba5e1b81120dd4dd1d9ced814a12a437d04e9f9e53b57230c5a5d8aba1c493"></a>

## annotations property — Property reference / cbb069cb9767 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

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
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-69df409b51bba1f9502348cdf3070a0304ea60d1f57433c6565fe6cfe34e0a48"></a>

<a id="canonical-b72cf35ea2022d3515a9f73ae11771a315e92d8424ed7dea3bfb5579a8792a84"></a>

## description property — Property reference / cbb069cb9767 / 5

Type: `"string"`. Computed.

Description of the ContainerRegistry.

Upstream description:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="canonical-5e6a7fd69a9aaf876d7c4d12350c89ee5fa545bd783bb80021694addfa15f5e3"></a>

<a id="canonical-b915437617dc18983ce8b4c08101537c2a29593da9518611eeb707a28e0a1081"></a>

## email property — Property reference / cbb069cb9767 / 6

Type: `"string"`. Computed.

Email. Email used for the registry.

Upstream description:

Email used for the registry.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "email",
    "formatDescription": "RFC 5322 email address, max 254 characters",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 3,
    "pattern": "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$",
    "validation": {
      "rfc": "RFC 5322"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.email": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.email": "true"
  }
}
```

<a id="canonical-31d037631aa4e694a6dd13eb1ff804b7d6c0b7ee3e55f4ad934c19d2353f82ff"></a>

<a id="canonical-b5c677202c8b522afbfd752b7636071fbbf948b06f85bebb2141dcb30f9f8216"></a>

## id property — Property reference / cbb069cb9767 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-c38831ab374f53025595ab6497a325fac617d06977c0fe365ecd05bf7bbe1152"></a>

<a id="canonical-cf40b7dcbc9d8fadbc244a867b0e005071282cbecd5fb8cc617d0601fc674fb4"></a>

## labels property — Property reference / cbb069cb9767 / 8

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-b273d1fecbd5175ed1d1aad25f5c5e1290934179f6685a8c02f6d89482a7df63"></a>

<a id="canonical-81fc9948ec75fe4604ca7ea7779876f9f1a9f2633adfbf93bcc9c4a6ea0b416c"></a>

## name property — Property reference / cbb069cb9767 / 9

Type: `"string"`. Required.

Name of the ContainerRegistry.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-ac718034b043d05a053aab7517622a9f6eb3f11ac73d8811b232a10d27924149"></a>

<a id="canonical-ed083c4e2bc71524bd228ff30020aed7518bdf23f9e8b186e56a8cea43a021a3"></a>

## namespace property — Property reference / cbb069cb9767 / 10

Type: `"string"`. Required.

Namespace where the ContainerRegistry exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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

- [password](data-sources--container_registry--reference--group-001.md#canonical-6601081072a6059170c0f5e562772abb4ea7e56ef45943ede4a78ff9b8ce9e6f): complete subsection reference.

<a id="canonical-69063cf5386b95022d5df1800356296a3aeb64c303de6e6f5b9062af5484a9f2"></a>

<a id="canonical-df548e82714e1fa6e982161a2be8e2e156d895d871c408b68247363b4eb50588"></a>

## registry property — Property reference / cbb069cb9767 / 11

Type: `"string"`. Computed.

Fully qualified name of the registry login server.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-a20f5b21a70ca5b8737bd495b1948213c57b547ce83c3905971b78e18990746e"></a>

<a id="canonical-37665a69f900a410297da47ede8a687bbadf21f8e297326deeccc700a189a604"></a>

## user_name property — Property reference / cbb069cb9767 / 12

Type: `"string"`. Computed.

User Name. Username used to access the registry.

Upstream description:

Username used to access the registry.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-b71d35d1218e9e7600e0fbb399c362ddef137fed654864c9dce8e2b7e3bd25ec"></a>

## All schema paths — Property reference / cbb069cb9767 / 13

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--container_registry--reference--group-001.md#canonical-40f96c0bbe0a4aa605b94ca12038890f1c48748a1d30efc11ef399f7b6aa5251) |
| `description` | [description](data-sources--container_registry--reference--group-001.md#canonical-69df409b51bba1f9502348cdf3070a0304ea60d1f57433c6565fe6cfe34e0a48) |
| `email` | [email](data-sources--container_registry--reference--group-001.md#canonical-5e6a7fd69a9aaf876d7c4d12350c89ee5fa545bd783bb80021694addfa15f5e3) |
| `id` | [id](data-sources--container_registry--reference--group-001.md#canonical-31d037631aa4e694a6dd13eb1ff804b7d6c0b7ee3e55f4ad934c19d2353f82ff) |
| `labels` | [labels](data-sources--container_registry--reference--group-001.md#canonical-c38831ab374f53025595ab6497a325fac617d06977c0fe365ecd05bf7bbe1152) |
| `name` | [name](data-sources--container_registry--reference--group-001.md#canonical-b273d1fecbd5175ed1d1aad25f5c5e1290934179f6685a8c02f6d89482a7df63) |
| `namespace` | [namespace](data-sources--container_registry--reference--group-001.md#canonical-ac718034b043d05a053aab7517622a9f6eb3f11ac73d8811b232a10d27924149) |
| `password` | [password](data-sources--container_registry--reference--group-001.md#canonical-2d7931f5f9d7b1f876b5d7eafa2d563b66a6784a640a249ff88ee6f3b9688116) |
| `password.blindfold_secret_info` | [password.blindfold_secret_info](data-sources--container_registry--reference--group-001.md#canonical-c3a23af35dabc86cb9ff448aa1c7c1b48aa7271c94b851f5e9f744dfbfaa5ed8) |
| `password.blindfold_secret_info.decryption_provider` | [password.blindfold_secret_info.decryption_provider](data-sources--container_registry--reference--group-001.md#canonical-ef74e182520be0d25acb999c4412d003d348ade565dd4581f747d46b36893766) |
| `password.blindfold_secret_info.location` | [password.blindfold_secret_info.location](data-sources--container_registry--reference--group-001.md#canonical-78d3178471a6c46070b0cd4ee1979c0c5abee734cbb68d7b0537dcd03c89f86c) |
| `password.blindfold_secret_info.store_provider` | [password.blindfold_secret_info.store_provider](data-sources--container_registry--reference--group-001.md#canonical-e17e5d5bc549627f8d3219f0f4b0424abfbe5557c606f72e4d3d8564f44ba3e6) |
| `password.clear_secret_info` | [password.clear_secret_info](data-sources--container_registry--reference--group-001.md#canonical-72c75d5cb6132f7b9b8afb6e01743886ba6d54f4efeeb2a92de2a6477e23fde1) |
| `password.clear_secret_info.provider_ref` | [password.clear_secret_info.provider_ref](data-sources--container_registry--reference--group-001.md#canonical-9ff803d37e0bae204a4c07a9f34cfbc071d2c39f7c0c2d5816346ccfd874f5b7) |
| `password.clear_secret_info.url` | [password.clear_secret_info.url](data-sources--container_registry--reference--group-001.md#canonical-f433efee42c661b5dd092e15dbd6c6d65c577b87ea5080d1af5502da3a61043c) |
| `registry` | [registry](data-sources--container_registry--reference--group-001.md#canonical-69063cf5386b95022d5df1800356296a3aeb64c303de6e6f5b9062af5484a9f2) |
| `user_name` | [user_name](data-sources--container_registry--reference--group-001.md#canonical-a20f5b21a70ca5b8737bd495b1948213c57b547ce83c3905971b78e18990746e) |

<a id="canonical-344276c60805949a5e12da8aa730561b94b604210d0f6bb16c7f660eec8cc8df"></a>

## Next pages — Property reference / cbb069cb9767 / 14

- [password](data-sources--container_registry--reference--group-001.md#canonical-6601081072a6059170c0f5e562772abb4ea7e56ef45943ede4a78ff9b8ce9e6f)
- [xcsh_container_registry](../data-sources/container_registry.md#canonical-17f6d44ebd4531a50d36aec8be1db646843481c3c17219cc6a239a366b16525f)

<a id="canonical-6601081072a6059170c0f5e562772abb4ea7e56ef45943ede4a78ff9b8ce9e6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b903e8abfddb7fa52b5eaa96d7bb2d3f091c00766bff39643a450d98adc6ce5"></a>

## password — password / 4bd278cafa5d / 2

Breadcrumbs:

- [xcsh_container_registry](../data-sources/container_registry.md#canonical-17f6d44ebd4531a50d36aec8be1db646843481c3c17219cc6a239a366b16525f)
- [Property reference](data-sources--container_registry--reference--group-001.md#canonical-e58992eaf4724a8cf0c7f0fa641dc1f519eedb53d24397f6ae8cf4129ef62c19)
- password

<a id="canonical-2d7931f5f9d7b1f876b5d7eafa2d563b66a6784a640a249ff88ee6f3b9688116"></a>

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

<a id="canonical-fbbb204565ab5224bfbbe812a105e6c84849bc568e2a193b2e56db3a9f04da01"></a>

## Direct properties — password / 4bd278cafa5d / 3

- [blindfold_secret_info](data-sources--container_registry--reference--group-001.md#canonical-96e22793db545a5261a56c7d478adb7d74f2c0a063659eb25e27f67c8276706e): complete subsection reference.

- [clear_secret_info](data-sources--container_registry--reference--group-001.md#canonical-ad36ab936a7e20288d27adb32768e6765234a60296d538d44ad04430261af143): complete subsection reference.

<a id="canonical-33df97c7d13079167e439de9d1bb5a79cefa3e06f2fa3b36a1df55d43c461a8b"></a>

## Next pages — password / 4bd278cafa5d / 4

- [password.blindfold_secret_info](data-sources--container_registry--reference--group-001.md#canonical-96e22793db545a5261a56c7d478adb7d74f2c0a063659eb25e27f67c8276706e)
- [password.clear_secret_info](data-sources--container_registry--reference--group-001.md#canonical-ad36ab936a7e20288d27adb32768e6765234a60296d538d44ad04430261af143)
- [Property reference](data-sources--container_registry--reference--group-001.md#canonical-e58992eaf4724a8cf0c7f0fa641dc1f519eedb53d24397f6ae8cf4129ef62c19)
- [xcsh_container_registry](../data-sources/container_registry.md#canonical-17f6d44ebd4531a50d36aec8be1db646843481c3c17219cc6a239a366b16525f)

<a id="canonical-96e22793db545a5261a56c7d478adb7d74f2c0a063659eb25e27f67c8276706e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d1088aca275afc43abec30fed92dbf8988e5e8037198fc3503175698772bb47d"></a>

## password.blindfold_secret_info — password.blindfold_secret_info / b5c4d954dcc9 / 2

Breadcrumbs:

- [xcsh_container_registry](../data-sources/container_registry.md#canonical-17f6d44ebd4531a50d36aec8be1db646843481c3c17219cc6a239a366b16525f)
- [Property reference](data-sources--container_registry--reference--group-001.md#canonical-e58992eaf4724a8cf0c7f0fa641dc1f519eedb53d24397f6ae8cf4129ef62c19)
- [password](data-sources--container_registry--reference--group-001.md#canonical-6601081072a6059170c0f5e562772abb4ea7e56ef45943ede4a78ff9b8ce9e6f)
- password.blindfold_secret_info

<a id="canonical-c3a23af35dabc86cb9ff448aa1c7c1b48aa7271c94b851f5e9f744dfbfaa5ed8"></a>

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

<a id="canonical-754ae555be247436ff0061f3107769cd381de290e60f03d7ee45855b5a1f933a"></a>

## Direct properties — password.blindfold_secret_info / b5c4d954dcc9 / 3

<a id="canonical-ef74e182520be0d25acb999c4412d003d348ade565dd4581f747d46b36893766"></a>

<a id="canonical-501ecfb1001e97731f47b028125dbacd53238145ddacacdd44b2ccc1ef68557e"></a>

## decryption_provider property — password.blindfold_secret_info / b5c4d954dcc9 / 4

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

<a id="canonical-78d3178471a6c46070b0cd4ee1979c0c5abee734cbb68d7b0537dcd03c89f86c"></a>

<a id="canonical-cd2c34c64d7aa56b107bc7e70f38374067b80a1ef9cc52da938e2acb514c7366"></a>

## location property — password.blindfold_secret_info / b5c4d954dcc9 / 5

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

<a id="canonical-e17e5d5bc549627f8d3219f0f4b0424abfbe5557c606f72e4d3d8564f44ba3e6"></a>

<a id="canonical-61c36b8df0627453dd5ea201c480cce35fb2724c5561c7c009d12635a3b0374c"></a>

## store_provider property — password.blindfold_secret_info / b5c4d954dcc9 / 6

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

<a id="canonical-1e0bb427c3a1319e2e118585fa3bf16b0d54ff90927b2e7f11ad765db99f104d"></a>

## Next pages — password.blindfold_secret_info / b5c4d954dcc9 / 7

- [password](data-sources--container_registry--reference--group-001.md#canonical-6601081072a6059170c0f5e562772abb4ea7e56ef45943ede4a78ff9b8ce9e6f)
- [xcsh_container_registry](../data-sources/container_registry.md#canonical-17f6d44ebd4531a50d36aec8be1db646843481c3c17219cc6a239a366b16525f)

<a id="canonical-ad36ab936a7e20288d27adb32768e6765234a60296d538d44ad04430261af143"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e6f844d7d9f32b39d1ee590524db9241e6fcc1902fea76dc5055413f1b18464"></a>

## password.clear_secret_info — password.clear_secret_info / 447387cf3a1e / 2

Breadcrumbs:

- [xcsh_container_registry](../data-sources/container_registry.md#canonical-17f6d44ebd4531a50d36aec8be1db646843481c3c17219cc6a239a366b16525f)
- [Property reference](data-sources--container_registry--reference--group-001.md#canonical-e58992eaf4724a8cf0c7f0fa641dc1f519eedb53d24397f6ae8cf4129ef62c19)
- [password](data-sources--container_registry--reference--group-001.md#canonical-6601081072a6059170c0f5e562772abb4ea7e56ef45943ede4a78ff9b8ce9e6f)
- password.clear_secret_info

<a id="canonical-72c75d5cb6132f7b9b8afb6e01743886ba6d54f4efeeb2a92de2a6477e23fde1"></a>

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

<a id="canonical-2f2ee4f993f2bd5c268cb2d7c970a0103400705cd0ea5cbc520c859318c356b6"></a>

## Direct properties — password.clear_secret_info / 447387cf3a1e / 3

<a id="canonical-9ff803d37e0bae204a4c07a9f34cfbc071d2c39f7c0c2d5816346ccfd874f5b7"></a>

<a id="canonical-239d7776dc8e5193b99e2d11e6876993aaf444b6199c1bd6259354bcfed83713"></a>

## provider_ref property — password.clear_secret_info / 447387cf3a1e / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-f433efee42c661b5dd092e15dbd6c6d65c577b87ea5080d1af5502da3a61043c"></a>

<a id="canonical-0c2d28a63b197e9d4d5b265fd8888dc9f51f63f8fae40c357fcf6f66a082ddf6"></a>

## url property — password.clear_secret_info / 447387cf3a1e / 5

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

<a id="canonical-8440c59ab44c88833262ab050bf9d51add8363fce0a495b10d557bb198ef0cfb"></a>

## Next pages — password.clear_secret_info / 447387cf3a1e / 6

- [password](data-sources--container_registry--reference--group-001.md#canonical-6601081072a6059170c0f5e562772abb4ea7e56ef45943ede4a78ff9b8ce9e6f)
- [xcsh_container_registry](../data-sources/container_registry.md#canonical-17f6d44ebd4531a50d36aec8be1db646843481c3c17219cc6a239a366b16525f)
