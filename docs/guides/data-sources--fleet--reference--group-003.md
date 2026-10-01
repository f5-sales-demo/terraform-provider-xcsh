---
page_title: "xcsh_fleet reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fleet reference."
---

# xcsh_fleet reference

<a id="canonical-46f82b18354b8457ea35b09a03ea4c58fe13274925bee4171522ec2402cb690d"></a>

## location property — storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_se / 3e32abab8ea7 / 5

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

<a id="canonical-3ceaaaffc3a5589f5ab28b5fc288d48e17d17043b01c9f9e693f4d7a30a63f8c"></a>

<a id="canonical-fba940edff7d223430364660e0c81b6809b5868e6961882e2a1b913b912f7ce0"></a>

## store_provider property — storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_se / 3e32abab8ea7 / 6

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

<a id="canonical-d01256879402b3ca9d0fc1ac88f7ad616abcec472013ff1450a6062d6ea0a5bc"></a>

## Next pages — storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_se / 3e32abab8ea7 / 7

- [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](data-sources--fleet--reference--group-002.md#canonical-857e40efddf272bf71ca9c783dc884d55c1c9c752b4452a3bfd34560770933d5)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-1dc1068e9c604a5ce653846780279e4b2a54696e8d6d7c1f9e5f77562e8a7e7e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23051895b7bbff6cb8c51daeb3e37635a51e5ef8bb07824547c23069b4acca6f"></a>

## storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info — storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret / 709df1640e45 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.hpe_storage](data-sources--fleet--reference--group-002.md#canonical-e33782be2bfea4e0b3bb991e64b209327a62e2329455fb701348a32623c3ba15)
- [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](data-sources--fleet--reference--group-002.md#canonical-857e40efddf272bf71ca9c783dc884d55c1c9c752b4452a3bfd34560770933d5)
- storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info

<a id="canonical-e5f7b5fb7c4021273e957bf96bb32343623e9280b73d822288513f3bcd639949"></a>

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

<a id="canonical-a1f451dddb2d7c0636b4059589d91543f889d42b9a838b161eef120e0244aa1a"></a>

## Direct properties — storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret / 709df1640e45 / 3

<a id="canonical-24b4c39e2b8e10e2af9f7ab2aa2c5f764b3905e534bdbd4b6e9951040fcf3e63"></a>

<a id="canonical-bba99027e6f302e43fca31aa546c1629286f3e25087c3d887e3099e4c5e26ab2"></a>

## provider_ref property — storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret / 709df1640e45 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-b165b50d97945b80d5f239411bc706bc03d0659614544af3df82e7b28cba047b"></a>

<a id="canonical-bf06049ebd1e509a27738e1bd569113476bca4bf9c5751a862f14d6169d9c2f3"></a>

## url property — storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret / 709df1640e45 / 5

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

<a id="canonical-99a4d03943eeb4039c75a5673e2573382f1a77b9ce0034df3c4164cc0291586a"></a>

## Next pages — storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret / 709df1640e45 / 6

- [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](data-sources--fleet--reference--group-002.md#canonical-857e40efddf272bf71ca9c783dc884d55c1c9c752b4452a3bfd34560770933d5)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-59605ae9a0ef21ad6a1b0b5fd854f34a24cdce64b84fcec9d30a2b23dd1f289c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5608b5d57f141f39f821e83a98963510a5bfd5c8ececf8311751305b21db12e7"></a>

## storage_device_list.storage_devices.hpe_storage.password — storage_device_list.storage_devices.hpe_storage.password / 29ea37fe667b / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.hpe_storage](data-sources--fleet--reference--group-002.md#canonical-e33782be2bfea4e0b3bb991e64b209327a62e2329455fb701348a32623c3ba15)
- storage_device_list.storage_devices.hpe_storage.password

<a id="canonical-cad629c71596e54508f4cb8798ec94088984fb05c95d654089234ca3316e179f"></a>

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

<a id="canonical-97da3945bbd21f503371e358495d0d73927d919ed82105be0cfe7724e7cafcd3"></a>

## Direct properties — storage_device_list.storage_devices.hpe_storage.password / 29ea37fe667b / 3

- [blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-2e7296556b4e139afea6a7dbc65d63e0253cbd4125fbc48021922f45e8a89502): complete subsection reference.

- [clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-1c18fa14239730b377917ad36191428b4276e83eaf712104de71560215262bd7): complete subsection reference.

<a id="canonical-05faedb03eab3ebf184313803cf10e2df5405ccfb33ba626b93a3b52f27dce79"></a>

## Next pages — storage_device_list.storage_devices.hpe_storage.password / 29ea37fe667b / 4

- [storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-2e7296556b4e139afea6a7dbc65d63e0253cbd4125fbc48021922f45e8a89502)
- [storage_device_list.storage_devices.hpe_storage.password.clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-1c18fa14239730b377917ad36191428b4276e83eaf712104de71560215262bd7)
- [storage_device_list.storage_devices.hpe_storage](data-sources--fleet--reference--group-002.md#canonical-e33782be2bfea4e0b3bb991e64b209327a62e2329455fb701348a32623c3ba15)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-2e7296556b4e139afea6a7dbc65d63e0253cbd4125fbc48021922f45e8a89502"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-42124b8e7285f15f9dcd4f6904c729cdca3fc72e751801783208b4046cf3b911"></a>

## storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info — storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info / 82ee64ddcce6 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.hpe_storage](data-sources--fleet--reference--group-002.md#canonical-e33782be2bfea4e0b3bb991e64b209327a62e2329455fb701348a32623c3ba15)
- [storage_device_list.storage_devices.hpe_storage.password](data-sources--fleet--reference--group-003.md#canonical-59605ae9a0ef21ad6a1b0b5fd854f34a24cdce64b84fcec9d30a2b23dd1f289c)
- storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info

<a id="canonical-f820ed53e7beede9fc6e2f654db83187ce593f146d0abc90bb4ddaf26b2c5270"></a>

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

<a id="canonical-48ed841b89c12728af67306c2fe3dbfc13b12d440693a263c3964c4d886717c9"></a>

## Direct properties — storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info / 82ee64ddcce6 / 3

<a id="canonical-59985ff0fe95da5e98e0c2732ce4b2e0ec4747667d3ad5864b2b16cce07906ba"></a>

<a id="canonical-9f16e94c0049fe84aeca95fb5b8b0abbbc53d73407dbfcf2f0ce07ceb3467592"></a>

## decryption_provider property — storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info / 82ee64ddcce6 / 4

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

<a id="canonical-9a453404c56e5cd9a07b3bb44e6e59fde1cf60b8fc796fb801e314f967d4a520"></a>

<a id="canonical-2933ad61a257c57a33388f588da70626c6c0906a5577f5274352cceaeffed653"></a>

## location property — storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info / 82ee64ddcce6 / 5

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

<a id="canonical-f2143b2ebd2cbcd4e9eb5942307181786de7c41c638d7905df64684675fb2cdc"></a>

<a id="canonical-d9dd57d903cad1c02aa8e208dc8a530f4b127fabe610fc45999d77e181c6e43e"></a>

## store_provider property — storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info / 82ee64ddcce6 / 6

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

<a id="canonical-0da9f7b497dcfc034f79cfddf3c832f6579fa0c851d8240eaa4b53e1eeb58164"></a>

## Next pages — storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info / 82ee64ddcce6 / 7

- [storage_device_list.storage_devices.hpe_storage.password](data-sources--fleet--reference--group-003.md#canonical-59605ae9a0ef21ad6a1b0b5fd854f34a24cdce64b84fcec9d30a2b23dd1f289c)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-1c18fa14239730b377917ad36191428b4276e83eaf712104de71560215262bd7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e74cb7951ee1ca386782631ba1766249670c3b6facef6a244a6596f400b4374"></a>

## storage_device_list.storage_devices.hpe_storage.password.clear_secret_info — storage_device_list.storage_devices.hpe_storage.password.clear_secret_info / e32b8b592d1a / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.hpe_storage](data-sources--fleet--reference--group-002.md#canonical-e33782be2bfea4e0b3bb991e64b209327a62e2329455fb701348a32623c3ba15)
- [storage_device_list.storage_devices.hpe_storage.password](data-sources--fleet--reference--group-003.md#canonical-59605ae9a0ef21ad6a1b0b5fd854f34a24cdce64b84fcec9d30a2b23dd1f289c)
- storage_device_list.storage_devices.hpe_storage.password.clear_secret_info

<a id="canonical-c0d9938be472df5f0f9d3cf105edacd7c3246c008d9cb1ee385de63d159f70fb"></a>

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

<a id="canonical-cff426fb8d1266b5fee771f17d27defd38a0d32c7190008e18aef480a375e6ea"></a>

## Direct properties — storage_device_list.storage_devices.hpe_storage.password.clear_secret_info / e32b8b592d1a / 3

<a id="canonical-48c86ebfd987ebde01ed831dc28efefa90537652501276fccd14cb192d681518"></a>

<a id="canonical-7e29002a92cc66a71b315420acc3705533f8a0e8c23de89b2395b55bbbd97aed"></a>

## provider_ref property — storage_device_list.storage_devices.hpe_storage.password.clear_secret_info / e32b8b592d1a / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-debdbbad88e9110853d98163faf3ca227de7ae8641bb2e90a6d256fc6fb80445"></a>

<a id="canonical-bc4c12805158dceed5e6d0f9d625ddcbb66c1ea0634f45b5de79ef8486634e67"></a>

## url property — storage_device_list.storage_devices.hpe_storage.password.clear_secret_info / e32b8b592d1a / 5

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

<a id="canonical-e57bd07ed4bf5fb44a74c4b3676b846e24c5fbea10bbcb958420f248e9550c53"></a>

## Next pages — storage_device_list.storage_devices.hpe_storage.password.clear_secret_info / e32b8b592d1a / 6

- [storage_device_list.storage_devices.hpe_storage.password](data-sources--fleet--reference--group-003.md#canonical-59605ae9a0ef21ad6a1b0b5fd854f34a24cdce64b84fcec9d30a2b23dd1f289c)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7416712a38d79d27e5bf58e3deaa11ef55978c5d52ba4c71d60fd042295379e"></a>

## storage_device_list.storage_devices.netapp_trident — storage_device_list.storage_devices.netapp_trident / e0a8afd4a736 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- storage_device_list.storage_devices.netapp_trident

<a id="canonical-e6e8ddc2004f3efc215209c333d6fc8e8f2e810252541a93741245b250b7f870"></a>

Type: `"single"`. Computed.

Device configuration for NetApp Trident Storage.

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

<a id="canonical-76ed914dfa5f11200bb844dc4a4de80729c704508f1425d9b1d35ec75fc966b7"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident / e0a8afd4a736 / 3

- [netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-af820a13292121258a8b737b37a6be7b1772437a86c37c360a519b899dd0ecdf): complete subsection reference.

- [netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-2beb37295fa85b4d6c4c5916001717fb8e9f448241abcb42bdc6b5678df6ed3c): complete subsection reference.

<a id="canonical-a7595ea84a1e78bcafe0fdd0ee5a49d159a3f7dac70fde42233da7ae77720c6e"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident / e0a8afd4a736 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-af820a13292121258a8b737b37a6be7b1772437a86c37c360a519b899dd0ecdf)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-2beb37295fa85b4d6c4c5916001717fb8e9f448241abcb42bdc6b5678df6ed3c)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-af820a13292121258a8b737b37a6be7b1772437a86c37c360a519b899dd0ecdf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be33e961cb8c0d2cf047b22aa02b500fc9da86d932ccdae4aa1e627402cbf77c"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / fccd70410c0f / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas

<a id="canonical-ef381087c940b70f900fad61a3d17841cf31289ef9d4306bdcf6c03bb654a105"></a>

Type: `"single"`. Computed.

Configuration of storage backend for NetApp ONTAP NAS.

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

<a id="canonical-db668531cf34b3e53c196fc3d92f75f294aa4fc047ebcc5608390e0288efbdee"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / fccd70410c0f / 3

- [auto_export_cidrs](data-sources--fleet--reference--group-003.md#canonical-d760c2375e41aedc5cd96875ef3660e1b5d93f34418dff726bf91cfde219d14f): complete subsection reference.

<a id="canonical-f527f14b77ca36542d945a79a2e1820cfa67c53dc08dada4765b200071fc3dca"></a>

<a id="canonical-4755f1a4e71dc418a99f8cc8a06ba390d82fa2ac757b89963c807d4221ae9a67"></a>

## auto_export_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / fccd70410c0f / 4

Type: `"bool"`. Computed.

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

<a id="canonical-911fa38195bdd95a9aba30c8f45575b87b43e56eb05ad9d1a54e95326edc8627"></a>

<a id="canonical-fce72ae9acdcb51d551423a13f5a7f8fdf77af6171215194d3862154ab9cd2e8"></a>

## backend_name property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / fccd70410c0f / 5

Type: `"string"`. Computed.

Configuration of Backend Name. Driver is name + '\_' + dataLIF.

Upstream description:

Configuration of Backend Name. Driver is name + "\_" + dataLIF.

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

<a id="canonical-93bac4abfb33ab6069e2a5fa8b0caed91b318223aff59c94b4485d2527275d53"></a>

<a id="canonical-1fda23e96a910d7a6485ddf08b695268aba861e4f3a3106bedf5b9e43e7ff486"></a>

## client_certificate property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / fccd70410c0f / 6

Type: `"string"`. Computed.

Please Enter Base64-encoded value of client certificate. Used for certificate-based auth.

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

- [client_private_key](data-sources--fleet--reference--group-003.md#canonical-107363901ddf1c59ac43b2f18d28651f126b531bd9bf58f7d0088ce762107167): complete subsection reference.

<a id="canonical-e1d23caf74e577d7e8b2c7417b7d17a951ad83b3eeadee065fd458d88925d819"></a>

<a id="canonical-1b0284e43473ddb4f752992e2240e2a87f9e5f7fb1463798d7ebeebcd55633d3"></a>

## data_lif_dns_name property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / fccd70410c0f / 7

Type: `"string"`. Computed.

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

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

<a id="canonical-7036ea1b996295da10771061ef8e08ea9ec0c37a54466825960735a4e1550964"></a>

<a id="canonical-04282752a2ccf4cd6349ad5ca6068fabe631db4d4c3e93c5d89031a7c66e5c46"></a>

## data_lif_ip property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / fccd70410c0f / 8

Type: `"string"`. Computed.

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

Upstream description:

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

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

<a id="canonical-d6bc4c1df5b6412cf922ee98c241fbfc0471a36f9efcfe8823fd3fd114f0a580"></a>

<a id="canonical-49cbf058d86e49d941e11b8d17727afb43910a497f2b8336a88d24f5f23e9112"></a>

## labels property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / fccd70410c0f / 9

Type: `["map", "string"]`. Computed.

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

<a id="canonical-60d936e6875ed9434e95eb40abb902c974820549c5b0eb506bde201d715d051d"></a>

<a id="canonical-a014d28c358f621afe404fd6c8d59f06d565091b43e45d008a802839fc13d4e4"></a>

## limit_aggregate_usage property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / fccd70410c0f / 10

Type: `"string"`. Computed.

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

<a id="canonical-82e9531e364225903fc89115fe4f46d3b887286f0939d40ea6f7b5989910c7db"></a>

<a id="canonical-e102c9ec94c91d1e2722269a693883883f500e63041277c0ab96d1bbdd4d45c8"></a>

## limit_volume_size property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / fccd70410c0f / 11

Type: `"string"`. Computed.

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

<a id="canonical-7d37f2939cc5eeb50a16ad4b078f6acc5fcad138a6062e69603aaa0986afbf8b"></a>

<a id="canonical-e643c79fedcdb39c55b5f2bdbb15b15d63e26c56a8ff697e95653ac49c5f6c7c"></a>

## management_lif_dns_name property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / fccd70410c0f / 12

Type: `"string"`. Computed.

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

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

<a id="canonical-4584e498ff56a9ab973227f7f4d74985fb87f8d313e1f042c23630f94cfdd1af"></a>

<a id="canonical-1a2e3506115a89ba633ddb8f87f14dbc8c1af0573b876a437d521dffde41767b"></a>

## management_lif_ip property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / fccd70410c0f / 13

Type: `"string"`. Computed.

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

Upstream description:

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

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

<a id="canonical-2f17a5b660b55226f79cad24638da083f29eeff3fc7b6f3a0dd29db9eb84f735"></a>

<a id="canonical-9cc39f3bb4db1f51edb2cee76fbd423ed0e16e0d39eebbd7feb12d143fc464de"></a>

## nfs_mount_options property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / fccd70410c0f / 14

Type: `"string"`. Computed.

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

- [password](data-sources--fleet--reference--group-003.md#canonical-9529afed48affccfdfb074d51093bfa7422bf7868d763061ea0d871c96fba260): complete subsection reference.

<a id="canonical-339f390d266a67e561fa7921218edc14e3b751a492f37e8060720eb75c56bec5"></a>

<a id="canonical-d9eba76dddb150ab97f996a3082e71d98aa82387e14db4c556d5f503ae1eaa6c"></a>

## region property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / fccd70410c0f / 15

Type: `"string"`. Computed.

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

- [storage](data-sources--fleet--reference--group-003.md#canonical-946081da32a34d13431ea7587157be7b965b84f81d0465790682cced55667b1f): complete subsection reference.

<a id="canonical-cfc67397bced749e408ced4ebb888bd8448ea88a894db62690efdcfb0984591e"></a>

<a id="canonical-ccde251ea5e02e9b6359ff4e8e286cf9c50a111391718ffcdce0f021bcb59a94"></a>

## storage_driver_name property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / fccd70410c0f / 16

Type: `"string"`. Computed.

\[Enum: ontap-nas|ontap-nas-economy|ontap-nas-flexgroup\] Storage Backend Driver. Configuration of
Backend Name. Possible values are \`ontap-nas\`, \`ontap-nas-economy\`, \`ontap-nas-flexgroup\`.

Upstream description:

Configuration of Backend Name.

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

<a id="canonical-c85ea91b14827785ab95c1ba86691f34d7fde9c084b22b179e0a83b5f5695930"></a>

<a id="canonical-c8a6fccefc04366c5ebe5fca7e6a0668d9b1a62a57fa63faa70c8da67c5782e0"></a>

## storage_prefix property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / fccd70410c0f / 17

Type: `"string"`. Computed.

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

<a id="canonical-0750b6275ca4edfd98816c1f1fab015212bfedf5c5ce68549fb2adc25103ca79"></a>

<a id="canonical-9bb958bc82735a030b026e25b4d2825d76388051bc50483ac2f7683ced48c3c6"></a>

## svm property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / fccd70410c0f / 18

Type: `"string"`. Computed.

Storage virtual machine to use. Derived if an SVM managementLIF is specified.

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

<a id="canonical-8bcfe3e0ec88da8fe5827c189f1606478370335f67ee291695f089f28036bd49"></a>

<a id="canonical-b5bd0096d41bc7135cd1742ec761bb4ed6cdc3012bb7dece2bea518173e0baad"></a>

## trusted_ca_certificate property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / fccd70410c0f / 19

Type: `"string"`. Computed.

Please Enter Base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth.

Upstream description:

Please Enter Base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth..

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

<a id="canonical-f0772c0af24cf87eaacfa85bc420fa2a70fb96b5cc5848e38daab1e29e38bcd8"></a>

<a id="canonical-a54209e11051d63d7a793eea5cbda598d657596746f9094075a68155ce00dcaa"></a>

## username property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / fccd70410c0f / 20

Type: `"string"`. Computed.

Username. Username to connect to the cluster/SVM.

Upstream description:

Username to connect to the cluster/SVM.

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

- [volume_defaults](data-sources--fleet--reference--group-003.md#canonical-4d0b1624655fe1c12998c26c67ef5e911355c1315edbe91d5b4e84c25936e44b): complete subsection reference.

<a id="canonical-3f6bfc293b78cb191c025b38a4417faa5e5fb9bbcf67a173f0f1810d2636c9eb"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas / fccd70410c0f / 21

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs](data-sources--fleet--reference--group-003.md#canonical-d760c2375e41aedc5cd96875ef3660e1b5d93f34418dff726bf91cfde219d14f)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](data-sources--fleet--reference--group-003.md#canonical-107363901ddf1c59ac43b2f18d28651f126b531bd9bf58f7d0088ce762107167)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](data-sources--fleet--reference--group-003.md#canonical-9529afed48affccfdfb074d51093bfa7422bf7868d763061ea0d871c96fba260)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](data-sources--fleet--reference--group-003.md#canonical-946081da32a34d13431ea7587157be7b965b84f81d0465790682cced55667b1f)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults](data-sources--fleet--reference--group-003.md#canonical-4d0b1624655fe1c12998c26c67ef5e911355c1315edbe91d5b4e84c25936e44b)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-d760c2375e41aedc5cd96875ef3660e1b5d93f34418dff726bf91cfde219d14f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f748ba619a92f750aa8b51bdb96d02bce3151cef2b8a852c3107c7f114bd95fe"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto / 9b1e04184d61 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-af820a13292121258a8b737b37a6be7b1772437a86c37c360a519b899dd0ecdf)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs

<a id="canonical-0a25d9d71924c1de296affa4818381fd2665cc19fe2f63efa8e49123c71e4ad2"></a>

Type: `"single"`. Computed.

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

<a id="canonical-f84fe1a5aa34efaf6ee134728ce1dc78960e79eee47cbb774328c7df8dc4732c"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto / 9b1e04184d61 / 3

<a id="canonical-21ad16ea62fe8663eeccb1599a1cf5c3893843021f19d1e47628352e1fdd4414"></a>

<a id="canonical-184ad20bd83a382da89f6647f709d94a9341d0bc4e5449c40172cb2b3c346910"></a>

## prefixes property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto / 9b1e04184d61 / 4

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-c7b03223f6871d3e2e821cd303aa63b090d0c03cce6b971094bd6d95498d3e48"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto / 9b1e04184d61 / 5

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-af820a13292121258a8b737b37a6be7b1772437a86c37c360a519b899dd0ecdf)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-107363901ddf1c59ac43b2f18d28651f126b531bd9bf58f7d0088ce762107167"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-302534510876539429a39d7db670dd763537997b6c4645bde3252bb31678fb05"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.clie / 863d27dc0758 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-af820a13292121258a8b737b37a6be7b1772437a86c37c360a519b899dd0ecdf)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key

<a id="canonical-d13f1fd894cae72721e639fa5bdf622eb66187416776dbd8eaa5941723f8ac02"></a>

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

<a id="canonical-f019f7aba61afcfa5e56ddb057b6db560b506314167f8018976af22398957bdd"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.clie / 863d27dc0758 / 3

- [blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-f1bb2ccf6c5305e162003aae62face17a307dfcb6f9c1e47a7c924996ddd0932): complete subsection reference.

- [clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-ab6149877f2e917fb8f8cc2308947ef77ebc06b566586e877d8e7e19b609a2a8): complete subsection reference.

<a id="canonical-4e452d0569b6b76ce56b6086d563fdfa22c886ba560ece35af669bd543a1379b"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.clie / 863d27dc0758 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-f1bb2ccf6c5305e162003aae62face17a307dfcb6f9c1e47a7c924996ddd0932)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-ab6149877f2e917fb8f8cc2308947ef77ebc06b566586e877d8e7e19b609a2a8)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-af820a13292121258a8b737b37a6be7b1772437a86c37c360a519b899dd0ecdf)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-f1bb2ccf6c5305e162003aae62face17a307dfcb6f9c1e47a7c924996ddd0932"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3df09518e4a3bfbb039abab037c7ec7ccd9bda1e40eca1b0e739c32b39312819"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.clie / 3b860fb20d79 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-af820a13292121258a8b737b37a6be7b1772437a86c37c360a519b899dd0ecdf)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](data-sources--fleet--reference--group-003.md#canonical-107363901ddf1c59ac43b2f18d28651f126b531bd9bf58f7d0088ce762107167)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info

<a id="canonical-eeabe534359e2251315326bdcf44d24f7d45e9bc5a13a04cbe6eb6349905b416"></a>

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

<a id="canonical-d3cf45e3b20360754991b51e2f8c60ffecf45d76a87f0b34a68f027079321ec1"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.clie / 3b860fb20d79 / 3

<a id="canonical-5717eb03112786b73e15f11200f900b0c7d4b6b64577b058c942f858780d0e89"></a>

<a id="canonical-65319caff557125f483aa63879c260f3b4b5b91329f7a6e84574fbfda1b433c3"></a>

## decryption_provider property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.clie / 3b860fb20d79 / 4

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

<a id="canonical-310a9f14f214c38d9edd472293e1b66506f86a322956a6e0aa7739ddce30830a"></a>

<a id="canonical-8e56de198ba61ba556a7cb99a2cb977ee0121fd4c0302cabed542176a0b2ea40"></a>

## location property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.clie / 3b860fb20d79 / 5

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

<a id="canonical-d01ad4fbf8d961fc1ca68a6dff4ef65387c1b309ad185855a32ffbb8bde6f2c8"></a>

<a id="canonical-4fd5d3b4d04f49e89de0260d8fb8ccb217ad34bb8c796293f6c926ded623f1da"></a>

## store_provider property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.clie / 3b860fb20d79 / 6

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

<a id="canonical-0ae5a7ef337db82358afaaf1dc196e523ce1ee243bc1431263c01ed98e0f7f93"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.clie / 3b860fb20d79 / 7

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](data-sources--fleet--reference--group-003.md#canonical-107363901ddf1c59ac43b2f18d28651f126b531bd9bf58f7d0088ce762107167)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-ab6149877f2e917fb8f8cc2308947ef77ebc06b566586e877d8e7e19b609a2a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-51845d3aec698f66da9171b5f0e593e76b77390347f641c42d34ee8f52bb4983"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.clie / 34caf92b515e / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-af820a13292121258a8b737b37a6be7b1772437a86c37c360a519b899dd0ecdf)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](data-sources--fleet--reference--group-003.md#canonical-107363901ddf1c59ac43b2f18d28651f126b531bd9bf58f7d0088ce762107167)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info

<a id="canonical-14687681dc4f6116b4d5eaec68310dfb2bc15640595caa16705035d9897d5bff"></a>

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

<a id="canonical-9363f3d781f1fd2dfde4e02259a51309407f3ebb5d1731893b654a8ef89d87ed"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.clie / 34caf92b515e / 3

<a id="canonical-981b5a9608de85209a71b3299e811cc3b19a21e2a3ff47364d4943e6042e33b5"></a>

<a id="canonical-c0bc8784a4d8b80cfbbd4207f8dad3732d961a4bc67ee05ab68326d8e37474c2"></a>

## provider_ref property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.clie / 34caf92b515e / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-6218a7defd2f125e5cfbfbc9edf67ccb8a8744367602deb25f954e2a8f4eded5"></a>

<a id="canonical-db0481a361d60b1f80197d74b860ddd7e87a09f944d921f5f7f16d64757f18f5"></a>

## url property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.clie / 34caf92b515e / 5

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

<a id="canonical-35d8c37e696e4be2cc46f5b0f915b9ac05d32866a127a29251fc260c178fac6c"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.clie / 34caf92b515e / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](data-sources--fleet--reference--group-003.md#canonical-107363901ddf1c59ac43b2f18d28651f126b531bd9bf58f7d0088ce762107167)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-9529afed48affccfdfb074d51093bfa7422bf7868d763061ea0d871c96fba260"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bcd7e7337b1fb3e9771e33000c8435235dbd6a04f214100377703c78606ff42a"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.pass / 4e79897f8bd4 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-af820a13292121258a8b737b37a6be7b1772437a86c37c360a519b899dd0ecdf)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password

<a id="canonical-1e8bec6765a10855ec41055588703959a40e68394e862b3e2a872d3f4a851cbf"></a>

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

<a id="canonical-f13ae172a2e413d10451172047b30941874c233890c7536463bfe773f6897965"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.pass / 4e79897f8bd4 / 3

- [blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-1f3e124ffba13f679f100b78b9996ec16aa92b05c318ad8780d0ad129ae6ee59): complete subsection reference.

- [clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-0ea7e17f461956ed3e12f45d6098ceb79af559a48f1c67f1c3fc09514ec1999d): complete subsection reference.

<a id="canonical-7b6f2e6355f349c2773f89d45b85296d6ea3270bd64d9e6ed2d671188f62a4a7"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.pass / 4e79897f8bd4 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-1f3e124ffba13f679f100b78b9996ec16aa92b05c318ad8780d0ad129ae6ee59)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-0ea7e17f461956ed3e12f45d6098ceb79af559a48f1c67f1c3fc09514ec1999d)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-af820a13292121258a8b737b37a6be7b1772437a86c37c360a519b899dd0ecdf)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-1f3e124ffba13f679f100b78b9996ec16aa92b05c318ad8780d0ad129ae6ee59"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4a0d7af3d3f3e6b4cd969e7c30d5cdedcedd80009a25e65db0b74b51ab49ffc"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.pass / 564967e3aeaf / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-af820a13292121258a8b737b37a6be7b1772437a86c37c360a519b899dd0ecdf)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](data-sources--fleet--reference--group-003.md#canonical-9529afed48affccfdfb074d51093bfa7422bf7868d763061ea0d871c96fba260)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info

<a id="canonical-600a40f39c54548c26d9e47b85f44b9e11c194ab7deaac6540af0256c789486e"></a>

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

<a id="canonical-a7a03eefe4d4400a2c992f66e9c2baad79d10e497237286ecf36c44a55119db6"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.pass / 564967e3aeaf / 3

<a id="canonical-105f021b10507171d89f554adf5ee42745d42671e045c39cfb5c0774c61148f4"></a>

<a id="canonical-3a37e556e765a2b507d86438767c025c39b5ba8a2ecade1dafd657585d46b928"></a>

## decryption_provider property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.pass / 564967e3aeaf / 4

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

<a id="canonical-24337c788254d47c3b811667c08e6927445cba6231e736de09821a82f2cce9b3"></a>

<a id="canonical-8ffb32d367ad83268e25d4fbdd2a6db79e1313ace8fc1ffe7873fbe67ab78fa7"></a>

## location property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.pass / 564967e3aeaf / 5

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

<a id="canonical-ab504f20508b9662dc1bf7ffab6736ce8ec04d2d4ac19524fac51d81a1e47a32"></a>

<a id="canonical-d6627d749931f25223689ae36fbdae01bc96f9f9d11b8104c1d3ba51f63d0de7"></a>

## store_provider property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.pass / 564967e3aeaf / 6

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

<a id="canonical-4a3efd21691ae02126119cd7bc55e0c36b6f3ae6359a7bb34d214203e25e0066"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.pass / 564967e3aeaf / 7

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](data-sources--fleet--reference--group-003.md#canonical-9529afed48affccfdfb074d51093bfa7422bf7868d763061ea0d871c96fba260)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-0ea7e17f461956ed3e12f45d6098ceb79af559a48f1c67f1c3fc09514ec1999d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e72bf3ea6965f8545826c8580cd72f8a2a22fdf3d684c786a8526343f3ea4b2c"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.pass / 124240b4a112 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-af820a13292121258a8b737b37a6be7b1772437a86c37c360a519b899dd0ecdf)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](data-sources--fleet--reference--group-003.md#canonical-9529afed48affccfdfb074d51093bfa7422bf7868d763061ea0d871c96fba260)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info

<a id="canonical-15e7a0cb8345478648d6ae970bbdde9b4f0cf3d8aad8e37f087917c44c5b8e4f"></a>

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

<a id="canonical-30afbea5a8c662d11824db53be4f11d8852ccf36aeb45db4c590a5258cdecfad"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.pass / 124240b4a112 / 3

<a id="canonical-ce96d4fa5a7209bc404a5104ab069c08516a887a3df2388afddac098f0d4bca9"></a>

<a id="canonical-0f15f52886ae958fa9045bb63d22f26b3333479dda3d3b12155952c90f357ddc"></a>

## provider_ref property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.pass / 124240b4a112 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-a308fb2f9e3f7a31998b15ea9ebd57eaae55c61c8ce7415af931885139f7478a"></a>

<a id="canonical-bc612d2a0da2d4d3116d19b5295d1e95210bd979a493aca16f7b4d2e9d42acb1"></a>

## url property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.pass / 124240b4a112 / 5

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

<a id="canonical-698b55adac1555be2b5480b08328927e6e5af5e363e3792fa6c35bd07c5a7abf"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.pass / 124240b4a112 / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](data-sources--fleet--reference--group-003.md#canonical-9529afed48affccfdfb074d51093bfa7422bf7868d763061ea0d871c96fba260)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-946081da32a34d13431ea7587157be7b965b84f81d0465790682cced55667b1f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-62688bf3bac82be251faf945446a50165ee5740b3755b827d3e6093d9924f86e"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 920309abf708 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-af820a13292121258a8b737b37a6be7b1772437a86c37c360a519b899dd0ecdf)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage

<a id="canonical-4ad8d17956f34c2269cbec737aa7bf8fd88c8059d8e352e48da99b51c39ffcf7"></a>

Type: `"list"`. Computed.

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

<a id="canonical-af0d3fb907594cbe59f73ecd7193c89c999b479885ca9c04075a727cb4b27f5c"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 920309abf708 / 3

<a id="canonical-113a506c1c74514995bae3f0a6e8f4af39efd3b48725fa57b94a4709093db7cc"></a>

<a id="canonical-7b1b4a99307a4d1326b22e66bdbd95f79e1e6eddcc009b0b3da9b2080cc100d7"></a>

## labels property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 920309abf708 / 4

Type: `["map", "string"]`. Computed.

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

- [volume_defaults](data-sources--fleet--reference--group-003.md#canonical-653f98fa9e64b5b16dbaed5a94cf8a021ad04a71fb8dead2599ea04253586a15): complete subsection reference.

<a id="canonical-a1350b100919214eed81f0513fb757410709e44548c31dd818023e7409e0c402"></a>

<a id="canonical-eea8df589b131c95ef294235b17dd1339fff0a35af199aef8859e667d77611f8"></a>

## zone property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 920309abf708 / 5

Type: `"string"`. Computed.

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

<a id="canonical-8c6c959c9850c84e7aba5cb815f94e7dfee928189f8634f8f003340e16518fd4"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 920309abf708 / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](data-sources--fleet--reference--group-003.md#canonical-653f98fa9e64b5b16dbaed5a94cf8a021ad04a71fb8dead2599ea04253586a15)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-af820a13292121258a8b737b37a6be7b1772437a86c37c360a519b899dd0ecdf)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-653f98fa9e64b5b16dbaed5a94cf8a021ad04a71fb8dead2599ea04253586a15"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5321cc0c3243359a3dacaab36e5bc7d317cd216d508ca5c13e8da3e06d160b9a"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 61f66ba5fd73 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-af820a13292121258a8b737b37a6be7b1772437a86c37c360a519b899dd0ecdf)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](data-sources--fleet--reference--group-003.md#canonical-946081da32a34d13431ea7587157be7b965b84f81d0465790682cced55667b1f)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults

<a id="canonical-aa08f95aba9f659666dd1807733ff7d353c6f958d035d08fd0de553fb056d43d"></a>

Type: `"single"`. Computed.

It controls how each volume is provisioned by default using these OPTIONS in a special section of
the configuration.

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

<a id="canonical-2c8d8089ab463c59633dc80f35fe49450535d314601d0247d88b85d9761eb3f6"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 61f66ba5fd73 / 3

<a id="canonical-19cb805289c1898ceeba55c8cb604f5f9a5b990e24507b3cfb9a27bbac13460a"></a>

<a id="canonical-9056d7132bf36a2705d170180978aae7d4fde8a3b2a4bacd9881861d4af58048"></a>

## adaptive_qos_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 61f66ba5fd73 / 4

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Exclusive with \[no\_qos qos\_policy\] Enter Adaptive QoS Policy Name.

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

<a id="canonical-5e7f7e8ac1e27c45fe6af98cd40049eb9f33736df751b4edd922306be869508f"></a>

<a id="canonical-a6d1675cf8152a22a5c6ee142873682954612609322ac4e67a02e09d83fb6aca"></a>

## encryption property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 61f66ba5fd73 / 5

Type: `"bool"`. Computed.

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

<a id="canonical-d097c9d12f85acf11c2805f5f31b86810195c61c1bd6b072965fa46640c5f08f"></a>

<a id="canonical-fe8ffbde4c7001b6db9b6519ba94513f653518ca08b3ea7984f4fb2a0fbb33a8"></a>

## export_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 61f66ba5fd73 / 6

Type: `"string"`. Computed.

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

- [no_qos](data-sources--fleet--reference--group-003.md#canonical-d7657d34a7c07214d971738021c636a6624d40728d4c44b1208fa3decc83001e): complete subsection reference.

<a id="canonical-5e180fd7f97f946a626f0f08341be06240d0368b2540cd5744a437610722600b"></a>

<a id="canonical-db6839b93c91601f9cfc2f3c6788698dcc7a9bd6e220a3ff37592c9967d1cf6b"></a>

## qos_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 61f66ba5fd73 / 7

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Exclusive with \[adaptive\_qos\_policy no\_qos\] Enter QoS Policy Name.

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

<a id="canonical-88768ea87a1410abb74135ac5a711628de703806e1bba21454b55fce1e26b23e"></a>

<a id="canonical-71ea290a885dbe0ab5c62bca7d159535d9760e3b1973b3de8bb5a11c38384002"></a>

## security_style property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 61f66ba5fd73 / 8

Type: `"string"`. Computed.

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

<a id="canonical-bb52305686ed274431c897d416577621afad1bb1f98c05cd115afb724f2fcc31"></a>

<a id="canonical-10a291c3b38f817f005fd67248e0aded104969f7db3c0cd1e924c88838c2bc95"></a>

## snapshot_dir property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 61f66ba5fd73 / 9

Type: `"bool"`. Computed.

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

<a id="canonical-736bbc8fddfccbb9b70adfddf5741a316452d8e8711ac6e9ad02269a33382b3d"></a>

<a id="canonical-cc94c91ca6f0eb86a897d2f928a6b571025872db25e1ddd23fb6a0cdde174d99"></a>

## snapshot_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 61f66ba5fd73 / 10

Type: `"string"`. Computed.

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

<a id="canonical-7675d6ce88606f8bd2ff8c1738a4c21c8655651388dd91354dd0ad1532ae950b"></a>

<a id="canonical-03271fd5543471fa919b0d1af57066c969948b202f01ed42c4d1258cf0845e16"></a>

## snapshot_reserve property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 61f66ba5fd73 / 11

Type: `"string"`. Computed.

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

<a id="canonical-dd1f61b66983ae1893bcacaa3c860df73add2846fd89058f62b89c4244a4fbe5"></a>

<a id="canonical-e7af917eece02dfb312b203d378865da0f6b8c3798d95aa1787a77cc1f3a24a7"></a>

## space_reserve property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 61f66ba5fd73 / 12

Type: `"string"`. Computed.

\[Enum: none|thick\] Space reservation mode; “none” (thin) or “volume” (thick). Possible values are
\`none\`, \`thick\`.

Upstream description:

Space reservation mode; “none” (thin) or “volume” (thick)

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

<a id="canonical-aab4d628fb2c01a5b453e519d47233cad72fcd342c6464136bda5ef2dda32889"></a>

<a id="canonical-19fb5857ca26bd619089e4b36eaaa3dfff01d8dcf6c3b344b462f187534a38c2"></a>

## split_on_clone property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 61f66ba5fd73 / 13

Type: `"bool"`. Computed.

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

<a id="canonical-83ec3364fd51b1733fac0038db7f7c22d298191ebc4a8f0b243b05c1459ea88f"></a>

<a id="canonical-40c7774e2c9392a77641294a902ced3b46a33c59bfa7cec4d06d28f7bf5316d7"></a>

## tiering_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 61f66ba5fd73 / 14

Type: `"string"`. Computed.

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

<a id="canonical-bddd4bd67fd9f364be8cfb4d98785e5e1cb6ab69849093667f9a78eae4aadd23"></a>

<a id="canonical-42552717238c674e94519f353ff4302e0c6451d13859bd771e7d1e63f8b3d1ef"></a>

## unix_permissions property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 61f66ba5fd73 / 15

Type: `"number"`. Computed.

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

<a id="canonical-f2297f51ec54dc356ff725f5aad413802caaa7f3fb88b436ea5ca4b05bd4f6e9"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 61f66ba5fd73 / 16

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos](data-sources--fleet--reference--group-003.md#canonical-d7657d34a7c07214d971738021c636a6624d40728d4c44b1208fa3decc83001e)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](data-sources--fleet--reference--group-003.md#canonical-946081da32a34d13431ea7587157be7b965b84f81d0465790682cced55667b1f)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-d7657d34a7c07214d971738021c636a6624d40728d4c44b1208fa3decc83001e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-88feb71d190d072e95d649864f5b8aa11e546a3e3b735821f7dde290bc6463aa"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 5537a1a82dd9 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-af820a13292121258a8b737b37a6be7b1772437a86c37c360a519b899dd0ecdf)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](data-sources--fleet--reference--group-003.md#canonical-946081da32a34d13431ea7587157be7b965b84f81d0465790682cced55667b1f)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](data-sources--fleet--reference--group-003.md#canonical-653f98fa9e64b5b16dbaed5a94cf8a021ad04a71fb8dead2599ea04253586a15)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos

<a id="canonical-3efa8b4c3be828276d53655519bc93ab54c0e513f3b02e8e3cb0a6414682d27f"></a>

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

<a id="canonical-d76cfedcde66f2758a04667247c1509e896ebc2f11155893d817d4ecfcc37eae"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 5537a1a82dd9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e788c16853c2171f32a2d229b3bae03e193ed56e3c0afbc35de448e796a5866a"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.stor / 5537a1a82dd9 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](data-sources--fleet--reference--group-003.md#canonical-653f98fa9e64b5b16dbaed5a94cf8a021ad04a71fb8dead2599ea04253586a15)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-4d0b1624655fe1c12998c26c67ef5e911355c1315edbe91d5b4e84c25936e44b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ab0d80da4b2e176eb560d83c7c21305e0e5f1bba606755151dcdaaaab5d9490"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / 09f454076bc3 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-af820a13292121258a8b737b37a6be7b1772437a86c37c360a519b899dd0ecdf)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults

<a id="canonical-d8a1addaf905223e9e84b887410f9051bfcf8df7d297edf7cde44d6f9abc5ae4"></a>

Type: `"single"`. Computed.

It controls how each volume is provisioned by default using these OPTIONS in a special section of
the configuration.

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

<a id="canonical-51476613936d74efe7114309db020f54618aafbfde5c17abbd57bf875c4b5b43"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / 09f454076bc3 / 3

<a id="canonical-ab3440b2d1bfaf73f1018b9d94b6c3c15e705fa238077be194c2485fe0690783"></a>

<a id="canonical-ce43d448ad44800465ae0ccdd4c1fd571b894de4d932fb474854807673202a37"></a>

## adaptive_qos_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / 09f454076bc3 / 4

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Exclusive with \[no\_qos qos\_policy\] Enter Adaptive QoS Policy Name.

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

<a id="canonical-d49a1e703f063d918795d43435bd60b88616076b5b05642643856e034f6c8504"></a>

<a id="canonical-fad8a85a49fd86a2e3ed914533376f25ee0fcc1b2fc6245d17a95da4e8184701"></a>

## encryption property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / 09f454076bc3 / 5

Type: `"bool"`. Computed.

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

<a id="canonical-c0a5b8eaf0b867652f0a7276bf543aa7acadbb9a624055ce06c7125558750269"></a>

<a id="canonical-acd74ab35ac1f2d53c3528c0fdca4faa779ac6cd47035334af7255925d4334d8"></a>

## export_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / 09f454076bc3 / 6

Type: `"string"`. Computed.

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

- [no_qos](data-sources--fleet--reference--group-003.md#canonical-a1b29e97d09a1f779a0aa3d12efe1472abb94fec7bb786411f199285c5a5fcd0): complete subsection reference.

<a id="canonical-a7c15e2a83715be3200a9cdb918704ae8ad7fe6f91f5e478fc2b4d4b2f68ef33"></a>

<a id="canonical-33dd44c15b2b28c8cc23782757795ca20881a651aaf7b21dfae6be5de2822b80"></a>

## qos_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / 09f454076bc3 / 7

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Exclusive with \[adaptive\_qos\_policy no\_qos\] Enter QoS Policy Name.

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

<a id="canonical-7e1693a5d69b1b310e5d92d1d288ac96b75261ef1a9e93f4aaa2d1ee516fe1cb"></a>

<a id="canonical-065a89e8f53e6234a2b218942889d84414adcdf39e7d1832116a155ae43f6685"></a>

## security_style property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / 09f454076bc3 / 8

Type: `"string"`. Computed.

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

<a id="canonical-7fdcc5b01b03c7b0f729cd6f95feb0dc583fc27bb19bbe3aa2cf9c56e19c380f"></a>

<a id="canonical-21c26c63281ebfaac22a9ad217f592fbbd2ac25c02a7bf8d54c955185b362181"></a>

## snapshot_dir property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / 09f454076bc3 / 9

Type: `"bool"`. Computed.

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

<a id="canonical-86a4e989d58efd9bec3ebfcd819aab285dbf9af0581a2e240004ba4a6fdc8533"></a>

<a id="canonical-999918eaacbb116b85a95a90d2143e6f6bf1def04e9b144eef11cc4665b3d685"></a>

## snapshot_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / 09f454076bc3 / 10

Type: `"string"`. Computed.

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

<a id="canonical-e45d85dc720e3a1271eb72e063932e2c40f6aa5722aad25a195231d08cb2202a"></a>

<a id="canonical-d7b8188989c5367aea3b289189764e7889658b5459ee35c4245cc729042e4182"></a>

## snapshot_reserve property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / 09f454076bc3 / 11

Type: `"string"`. Computed.

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

<a id="canonical-0cf5abc1c6c4ccce2274edeaf7e95ee14d7ba774c67b2ed389b4cb89974bbefa"></a>

<a id="canonical-a6b6d2c4bf1e3878c99d8abd355d8dbdf664085bcbaea8346936768682afd410"></a>

## space_reserve property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / 09f454076bc3 / 12

Type: `"string"`. Computed.

\[Enum: none|thick\] Space reservation mode; “none” (thin) or “volume” (thick). Possible values are
\`none\`, \`thick\`.

Upstream description:

Space reservation mode; “none” (thin) or “volume” (thick)

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

<a id="canonical-7af7d8fecb51f970dd986445abebe678d633956af9c01151527058e7cd60784d"></a>

<a id="canonical-bfcfdf20d53d81d6d5fa6891353972743b17a9ca18816e2f99eb0c8f773e3f0c"></a>

## split_on_clone property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / 09f454076bc3 / 13

Type: `"bool"`. Computed.

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

<a id="canonical-3b154d1a7dc42a5bc85cc66903f1d9ba034d27cd3f2268c2bf858d7be8b255c2"></a>

<a id="canonical-c53e9f4472b4f3982f184776f566efe614cc64e43359cadda6205d84cdfa663d"></a>

## tiering_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / 09f454076bc3 / 14

Type: `"string"`. Computed.

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

<a id="canonical-a391ae65bea0d117eaeb410f49ea43c31f67e586cc567e58e64ce181f6b7647b"></a>

<a id="canonical-0862996db2a59f021186fcaf936f5e9debcc4418ee3568a54d89d521cf16e7d5"></a>

## unix_permissions property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / 09f454076bc3 / 15

Type: `"number"`. Computed.

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

<a id="canonical-ada5d60a33b5b7e408fbc3e61061b37da393b927efa41e662e500ef58c8d3a53"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / 09f454076bc3 / 16

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos](data-sources--fleet--reference--group-003.md#canonical-a1b29e97d09a1f779a0aa3d12efe1472abb94fec7bb786411f199285c5a5fcd0)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-af820a13292121258a8b737b37a6be7b1772437a86c37c360a519b899dd0ecdf)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-a1b29e97d09a1f779a0aa3d12efe1472abb94fec7bb786411f199285c5a5fcd0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eaa2873722c2fb6d2a35a4e518b022028ca84e4a74e1eb5cff9184a8951d88f8"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / 60c11ad52e43 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-af820a13292121258a8b737b37a6be7b1772437a86c37c360a519b899dd0ecdf)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults](data-sources--fleet--reference--group-003.md#canonical-4d0b1624655fe1c12998c26c67ef5e911355c1315edbe91d5b4e84c25936e44b)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos

<a id="canonical-8904ac0de240ab7430156ac0c7ccb462d87c58e86ed81e06511091386b0237e2"></a>

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

<a id="canonical-92f0045e9920f70cbdd9711c1cbe9221fe1f7b182cd6fcfb9253177971ab79bc"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / 60c11ad52e43 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-35dcc58cd21611da7c180c2f10939b169d2f7094e716fefcb7303fcacd7e730e"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volu / 60c11ad52e43 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults](data-sources--fleet--reference--group-003.md#canonical-4d0b1624655fe1c12998c26c67ef5e911355c1315edbe91d5b4e84c25936e44b)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-2beb37295fa85b4d6c4c5916001717fb8e9f448241abcb42bdc6b5678df6ed3c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d2a96690fdd478977c9cad88e9b8fb89896dec54b2be548cd468244cd586c91a"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / c8cfffdff9b1 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san

<a id="canonical-36582d3705b383c029ae82b01b7cf1043dd509a781f5786cdefd7082350dcc31"></a>

Type: `"single"`. Computed.

Configuration of storage backend for NetApp ONTAP SAN.

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

<a id="canonical-1fcf01f0adf81f50a472666fb672230cb66655a4b0ecdc67cc3b2c4e8df53967"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / c8cfffdff9b1 / 3

<a id="canonical-534bbdba5ee2585110861b0fc3181cf1592d2129875ce64c9a49262a75575e82"></a>

<a id="canonical-17fbb3edcff56079519e1e7684270954b0b12f568f259b8f8b530b2c87dd82ad"></a>

## client_certificate property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / c8cfffdff9b1 / 4

Type: `"string"`. Computed.

Please Enter Base64-encoded value of client certificate. Used for certificate-based auth.

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

- [client_private_key](data-sources--fleet--reference--group-003.md#canonical-722b4fbaa4c05dac3285d9908593768e30884b7cf755aa193f359904fac06c72): complete subsection reference.

<a id="canonical-6b16d38df887113866eea136d0495ce315734658671d99e1ded159238a9f1874"></a>

<a id="canonical-7a85ee787ac3a937de14046ce1e6d192057eff5c078644968fb35559bd8a12b3"></a>

## data_lif_dns_name property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / c8cfffdff9b1 / 5

Type: `"string"`. Computed.

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

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

<a id="canonical-67ca8ca4a71ef18e0fa297193c67ff7d32870961561479e2ccd4cedd16993832"></a>

<a id="canonical-9503577b9931e63062e196add0c6b542613c7f8763f8df5315b6407f2be3ba1d"></a>

## data_lif_ip property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / c8cfffdff9b1 / 6

Type: `"string"`. Computed.

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

Upstream description:

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

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

<a id="canonical-84edcb7f4bcb89159bf34777ff15b3cb5f7cd840b98f1c49932ad5247933d943"></a>

<a id="canonical-2aded277be9b435d570a1c20da218067bf32da06f7476fd95614871ac01a21e2"></a>

## igroup_name property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / c8cfffdff9b1 / 7

Type: `"string"`. Computed.

Name of the igroup for SAN volumes to use.

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

<a id="canonical-2a0d90b35e5ecbfa91b77c71879d9f3f2d0f05a33291828cc84e5e7a543be643"></a>

<a id="canonical-497470dba2a79974078efdb768561dff72de8571644c9153814cd7d660f7a1c7"></a>

## labels property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / c8cfffdff9b1 / 8

Type: `["map", "string"]`. Computed.

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

<a id="canonical-ea75a01e99b30cfc5a8d559cbdb89ffdda4315fb00c9d729aa0aa9704ace6cad"></a>

<a id="canonical-ef6653ca9492a176589ad3181ec534b45e5aaadd651c15004b3055173e0f45f2"></a>

## limit_aggregate_usage property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / c8cfffdff9b1 / 9

Type: `"number"`. Computed.

Fail provisioning if usage is above this percentage. Not enforced by default.

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

<a id="canonical-1f6e0c4618b156220c8272389e6449ae6c9255fa35fbb34777e127513654d49a"></a>

<a id="canonical-b0a097cddfeb17dc4a636c08d3d02e4d0dbffa3ce86cba9e2896a9293f0c1b1e"></a>

## limit_volume_size property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / c8cfffdff9b1 / 10

Type: `"number"`. Computed.

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

<a id="canonical-f4bc4e2f90ffadcbe84cbfee964dbdfc4b2b1a42af28225ae0444e8fc8841ecd"></a>

<a id="canonical-c5c53aa3122934bea3f23bdd71673d5523a55b6d3b378a5e41454b1341bc1afb"></a>

## management_lif_dns_name property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / c8cfffdff9b1 / 11

Type: `"string"`. Computed.

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

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

<a id="canonical-40a13a5939bcdd2ae5bb31f34838d79145f3212dce4cb379e2c4f52413c57d0e"></a>

<a id="canonical-52c08b6332999f6819cd70612a542a7adade1ab1a2f636d73d22af477208ed63"></a>

## management_lif_ip property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / c8cfffdff9b1 / 12

Type: `"string"`. Computed.

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

Upstream description:

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

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

- [no_chap](data-sources--fleet--reference--group-003.md#canonical-f787885fb70888147b178eddb5192b8561df7bf13b8c6d698e49100063774df4): complete subsection reference.

- [password](data-sources--fleet--reference--group-003.md#canonical-584804818ee7574d5486fc5eaa502cd543bd4cca2e8e5079ca2de0590bd46f67): complete subsection reference.

<a id="canonical-882844eb8971292bdfecd9b9c7f8ef8907b2390544c9339ea8df1f82c0091f77"></a>

<a id="canonical-429ec65d8d624676449bf57f48e03bd40ffa3a73eae51fdb427996a4c14cab1c"></a>

## region property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / c8cfffdff9b1 / 13

Type: `"string"`. Computed.

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

- [storage](data-sources--fleet--reference--group-003.md#canonical-23c6a9da3581feb8bcb034a8042769d9431330c7508961c456f622bc4646bdc0): complete subsection reference.

<a id="canonical-25ea3ca1b06728e1dfa5464c0f68dccc8d34daed42909959ce82961bc1d1bc51"></a>

<a id="canonical-ab28e05e37d2457b0b6a612b874bc77277777f7a01da2ebdd1eb46d64ee506b8"></a>

## storage_driver_name property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / c8cfffdff9b1 / 14

Type: `"string"`. Computed.

\[Enum: ontap-san|ontap-san-economy|ontap-nas-flexgroup\] Storage Backend Driver. Configuration of
Backend Name. Possible values are \`ontap-san\`, \`ontap-san-economy\`, \`ontap-nas-flexgroup\`.

Upstream description:

Configuration of Backend Name.

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

<a id="canonical-3a699bcfc1367320c2317487d4b23b142d02990c79ab9a37e10d697a697f1e9b"></a>

<a id="canonical-f855f19358764f2addcc16f97b79eb81e8d636047499341001639809fdd8854b"></a>

## storage_prefix property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / c8cfffdff9b1 / 15

Type: `"string"`. Computed.

Prefix used when provisioning new volumes in the SVM. Once set this cannot be updated.

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

<a id="canonical-1f819dacd7165fa952cc6144afa7e500b725243ca79061a5b7258ebd95794dfc"></a>

<a id="canonical-8e58d9f1dc973ce11ea4aedd6f7e2019b47d6bc62f205e7057f67d7db077d8ea"></a>

## svm property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / c8cfffdff9b1 / 16

Type: `"string"`. Computed.

Storage virtual machine to use. Derived if an SVM managementLIF is specified.

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

<a id="canonical-f65301a556358bda53f5e9095f15576e66505a7cf5be85f619b92738e076a1c4"></a>

<a id="canonical-23a267458c9eea658e354f21c4edd7c3c7a0760c3f6ccf88aa8a0b4b0a23a206"></a>

## trusted_ca_certificate property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / c8cfffdff9b1 / 17

Type: `"string"`. Computed.

Please Enter Base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth.

Upstream description:

Please Enter Base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth..

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

- [use_chap](data-sources--fleet--reference--group-003.md#canonical-ceb2d887e3260398c1a5ef49341011226291126e755c1394ad71a3ba4b8c46ab): complete subsection reference.

<a id="canonical-8f0f52ad909f908fc7a4f915f440318f81182121448377f816cf3acf721bf6c7"></a>

<a id="canonical-5428588545d3df384c182fc0b7a47d6f655865013261014cc51b796526d1275e"></a>

## username property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / c8cfffdff9b1 / 18

Type: `"string"`. Computed.

Username. Username to connect to the cluster/SVM.

Upstream description:

Username to connect to the cluster/SVM.

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

- [volume_defaults](data-sources--fleet--reference--group-003.md#canonical-36b7100b6badb73b70262a1760b03aec196f91b255873a3ac1ad65e398eec79c): complete subsection reference.

<a id="canonical-d94e055f57b5dfa71d683bbdc2bd249889cb0007467c6f0e1173f420a01ea675"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san / c8cfffdff9b1 / 19

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](data-sources--fleet--reference--group-003.md#canonical-722b4fbaa4c05dac3285d9908593768e30884b7cf755aa193f359904fac06c72)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap](data-sources--fleet--reference--group-003.md#canonical-f787885fb70888147b178eddb5192b8561df7bf13b8c6d698e49100063774df4)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](data-sources--fleet--reference--group-003.md#canonical-584804818ee7574d5486fc5eaa502cd543bd4cca2e8e5079ca2de0590bd46f67)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](data-sources--fleet--reference--group-003.md#canonical-23c6a9da3581feb8bcb034a8042769d9431330c7508961c456f622bc4646bdc0)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--fleet--reference--group-003.md#canonical-ceb2d887e3260398c1a5ef49341011226291126e755c1394ad71a3ba4b8c46ab)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults](data-sources--fleet--reference--group-003.md#canonical-36b7100b6badb73b70262a1760b03aec196f91b255873a3ac1ad65e398eec79c)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-722b4fbaa4c05dac3285d9908593768e30884b7cf755aa193f359904fac06c72"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-83f6613cfe0ceb8cc28826f9e0c32552f3cad8806ceba870d5c886c016d533ad"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.clie / bc648bb1cb3f / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-2beb37295fa85b4d6c4c5916001717fb8e9f448241abcb42bdc6b5678df6ed3c)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key

<a id="canonical-cd0f2624564cff86baf21360ae7943ab0f825455d0ae41bd2fac8abe53bb056d"></a>

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

<a id="canonical-2af558f6d69ce3c1720b9e6101fdd0aa9b494170a7939a9558358e8515e5b7e6"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.clie / bc648bb1cb3f / 3

- [blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-ea3100b505955c4860e296f64880866e3a006c9763300837491223a8aa31f304): complete subsection reference.

- [clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-e98fd87a81e00f558545290259f018a1c2807c2dd5f9277e0a1f040d86c89fb1): complete subsection reference.

<a id="canonical-d91c0134d88a17d8d02d34e86577fcc3076b560cf80f9a858721d75bb4df3b04"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.clie / bc648bb1cb3f / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-ea3100b505955c4860e296f64880866e3a006c9763300837491223a8aa31f304)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-e98fd87a81e00f558545290259f018a1c2807c2dd5f9277e0a1f040d86c89fb1)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-2beb37295fa85b4d6c4c5916001717fb8e9f448241abcb42bdc6b5678df6ed3c)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-ea3100b505955c4860e296f64880866e3a006c9763300837491223a8aa31f304"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e1928e9ab1fca5a2e469a9f343378a2ea0a7c15b9b4943e969d796ab40ea0a7b"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.clie / c3ce7355b359 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-2beb37295fa85b4d6c4c5916001717fb8e9f448241abcb42bdc6b5678df6ed3c)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](data-sources--fleet--reference--group-003.md#canonical-722b4fbaa4c05dac3285d9908593768e30884b7cf755aa193f359904fac06c72)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info

<a id="canonical-e0bef78a9578e784b38d7175d1a6f8079067a7e72433512882a9331957eda40d"></a>

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

<a id="canonical-ce6818ce578f2c508ffe74353004a04cff7512d344d8a27c07dd7eb65724b174"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.clie / c3ce7355b359 / 3

<a id="canonical-542d5d966d25b3835494cc096cb9c6ba7d783038a4769869c67ef0a8e38d912b"></a>

<a id="canonical-2ba10b01c197632899294eb5c21dba5aa2b0c2cd2939f3e33f76465a08684eb4"></a>

## decryption_provider property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.clie / c3ce7355b359 / 4

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

<a id="canonical-03e50b74846cf4aedefd784076afdc184fd77ea973609ac37049ea24d1061d19"></a>

<a id="canonical-552bbff90e29d29ab644119f5fb349431fd95e763022b2c971e0dc415e208d99"></a>

## location property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.clie / c3ce7355b359 / 5

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

<a id="canonical-610002603def2cb443245b30941bf1e0b057e5975e39506e2eb69e5c1bfcfc0d"></a>

<a id="canonical-9a7e4acd38d50bca992347638b552d36c8a04ed6d679be953062a286746cd3c9"></a>

## store_provider property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.clie / c3ce7355b359 / 6

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

<a id="canonical-803e0ebb761a2679387b5869ff60904c49332e4ff40eaf6db0c479c4c8a4c9ed"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.clie / c3ce7355b359 / 7

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](data-sources--fleet--reference--group-003.md#canonical-722b4fbaa4c05dac3285d9908593768e30884b7cf755aa193f359904fac06c72)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-e98fd87a81e00f558545290259f018a1c2807c2dd5f9277e0a1f040d86c89fb1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-601e118dbeec4a4b8d699484e390437d9815cd22923d89f932239685b054abb4"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.clie / 34d28ba85ee0 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-2beb37295fa85b4d6c4c5916001717fb8e9f448241abcb42bdc6b5678df6ed3c)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](data-sources--fleet--reference--group-003.md#canonical-722b4fbaa4c05dac3285d9908593768e30884b7cf755aa193f359904fac06c72)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info

<a id="canonical-7b0a6b69cd83c552613ae4501ecb025fe1d576c1bda8cde921374dbd5db2aea1"></a>

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

<a id="canonical-d1e9f226d0813c185618dd223eebf3042f825e7f18e182dab39b6e0249dc8876"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.clie / 34d28ba85ee0 / 3

<a id="canonical-93deb9b356a3e653c37cf32cccc9e0f8d9416d8d4fb07d3f1bc42296df7a5cd0"></a>

<a id="canonical-6c29e166cf04445bb447cdce5ef843bac6068275313ca6d2d2237fdd12d94c84"></a>

## provider_ref property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.clie / 34d28ba85ee0 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2e8be5c557bb37c46147289dd7358767dfde8eee787c82ada1b74e895926ab09"></a>

<a id="canonical-56413ec334f0387adf61a88c2070fb49dff722849b9e834167230af091332631"></a>

## url property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.clie / 34d28ba85ee0 / 5

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

<a id="canonical-dfcc8980a145b7340640599e2b93a93f6306546ee1244b77cc630f9ea335ec76"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.clie / 34d28ba85ee0 / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](data-sources--fleet--reference--group-003.md#canonical-722b4fbaa4c05dac3285d9908593768e30884b7cf755aa193f359904fac06c72)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-f787885fb70888147b178eddb5192b8561df7bf13b8c6d698e49100063774df4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e247eb7d1d661e4f73488ba0a0e7afbea716317067787074dc07995b4749312"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_c / a951dbc2d53c / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-2beb37295fa85b4d6c4c5916001717fb8e9f448241abcb42bdc6b5678df6ed3c)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap

<a id="canonical-34356a201b9f4dafe60d6d7e8c1c7a5d7209b338c919a7a4f17f3b3a471f3d1a"></a>

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

<a id="canonical-56d5e5baa9542439b8ed2ee6d5e98e23564005e48e2ac8ea2a99c61ef17a0a53"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_c / a951dbc2d53c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ba61579ae6330e93dad6092b86386dd468138daa5aadde41d5d866f185e8ceaa"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_c / a951dbc2d53c / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-2beb37295fa85b4d6c4c5916001717fb8e9f448241abcb42bdc6b5678df6ed3c)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-584804818ee7574d5486fc5eaa502cd543bd4cca2e8e5079ca2de0590bd46f67"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7f07c2effab0360cf008e7c892d9549a447a856ee9c733874f1d7b1f03ac941b"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.pass / e4b2c3d11828 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-2beb37295fa85b4d6c4c5916001717fb8e9f448241abcb42bdc6b5678df6ed3c)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password

<a id="canonical-34b8d03a88fe8fafd20299f7fc341adb078cd349428a3aa32903a856ab2a51e0"></a>

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

<a id="canonical-b76bb56795123f97f41c14358f08f270a88d959f95e669309d53e22c9f1860c0"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.pass / e4b2c3d11828 / 3

- [blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-df5b0401a094247b963b6c5e4e8dc12ad0382475262fc6b6c6a6d826139dccbb): complete subsection reference.

- [clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-d28163d69e4b5a430db4b230bfee231c1368d7efab05e22f194a7a18c00f240d): complete subsection reference.

<a id="canonical-414de902d1729340077cdafd9a445743b682b9efec18221ca7f225016bd53ff6"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.pass / e4b2c3d11828 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-df5b0401a094247b963b6c5e4e8dc12ad0382475262fc6b6c6a6d826139dccbb)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-d28163d69e4b5a430db4b230bfee231c1368d7efab05e22f194a7a18c00f240d)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-2beb37295fa85b4d6c4c5916001717fb8e9f448241abcb42bdc6b5678df6ed3c)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-df5b0401a094247b963b6c5e4e8dc12ad0382475262fc6b6c6a6d826139dccbb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b543d1f771e8aa321beeaaab99e9ec6f63c7530081ed21d93f0a5235dbff7fe"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.pass / e7b79732cdae / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-2beb37295fa85b4d6c4c5916001717fb8e9f448241abcb42bdc6b5678df6ed3c)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](data-sources--fleet--reference--group-003.md#canonical-584804818ee7574d5486fc5eaa502cd543bd4cca2e8e5079ca2de0590bd46f67)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info

<a id="canonical-45384795643df51965e2e8d4d18598b912b7065518d3a724a40023074ba98ab1"></a>

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

<a id="canonical-6983160b57d0bb479d566408131aee362c34d8e3619ed853c5cd758976571f4d"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.pass / e7b79732cdae / 3

<a id="canonical-c94e9e9d67a64be00c9ed5c0b51c0f999dd420c593944a28aad7451055d72c52"></a>

<a id="canonical-31e303085cd32afaf91e4ad307597759461dd3064e185ff52ec50bd3e226be7b"></a>

## decryption_provider property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.pass / e7b79732cdae / 4

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

<a id="canonical-2aab6dcfdfe900466d79d0a247274935c95b54a98ffc4d61581fbb41a7dbe25b"></a>

<a id="canonical-44f292d0ff2b79d5d6ae155f2bc6423a06f483b0be0d963baa1cd71c8904e108"></a>

## location property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.pass / e7b79732cdae / 5

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

<a id="canonical-7567847b8b721a836d5c73dda0605f6dbdf546b21028e6f6080fedc2209c337f"></a>

<a id="canonical-bdc5a8c95d8a371a8ee0e670667d4e66b759ff671dbeb487019a8cae6377e8f8"></a>

## store_provider property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.pass / e7b79732cdae / 6

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

<a id="canonical-34c1866851aa739c248f0dc1f0b2bfdf04296fa439dd55dd54240aa9025f1c92"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.pass / e7b79732cdae / 7

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](data-sources--fleet--reference--group-003.md#canonical-584804818ee7574d5486fc5eaa502cd543bd4cca2e8e5079ca2de0590bd46f67)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-d28163d69e4b5a430db4b230bfee231c1368d7efab05e22f194a7a18c00f240d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-efc7b677d82b6326e7b63dae159541ca339e9bb7cb8a7f02b52445520aedd9b7"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.pass / 44fe835cddd1 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-2beb37295fa85b4d6c4c5916001717fb8e9f448241abcb42bdc6b5678df6ed3c)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](data-sources--fleet--reference--group-003.md#canonical-584804818ee7574d5486fc5eaa502cd543bd4cca2e8e5079ca2de0590bd46f67)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info

<a id="canonical-ed109466bc8c79ee3f4fca86762e4b25a920fc861965da9647e85843b1976793"></a>

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

<a id="canonical-96494f9e0d6c57107f1cd458ffcb5f5c5f3cd1450b3acf905254754f10e52f1f"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.pass / 44fe835cddd1 / 3

<a id="canonical-f2d4159a0b89b3e6ed4a3b0d6e39279ae43282073b4ddcd15de3f7c46d678df9"></a>

<a id="canonical-e00590c40050aa485d8544c0277944068ba09631be06578ac85273dfab1d2e2f"></a>

## provider_ref property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.pass / 44fe835cddd1 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-42ed4f0373d3ee88762cd50e48b14af826d98d4da07ab9c698a0f4900b44d472"></a>

<a id="canonical-d77c51263e0cfedd23262b1676191645311115309d198bd2067778d5d4b8f95c"></a>

## url property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.pass / 44fe835cddd1 / 5

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

<a id="canonical-66e3b51738651e03a90fdbbbf8504df5d7a4647bc7006a6323f239b9395e07c2"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.pass / 44fe835cddd1 / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](data-sources--fleet--reference--group-003.md#canonical-584804818ee7574d5486fc5eaa502cd543bd4cca2e8e5079ca2de0590bd46f67)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-23c6a9da3581feb8bcb034a8042769d9431330c7508961c456f622bc4646bdc0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-53d49278173dddf06c63ba15d8c734f2fbef2162f4e29d6c7c2c7f77ae67857f"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / 82f5d5c7a775 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-2beb37295fa85b4d6c4c5916001717fb8e9f448241abcb42bdc6b5678df6ed3c)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage

<a id="canonical-34db7506229fb88cc65890a3b1d6237493a5e5010b7250952413de616090e5d4"></a>

Type: `"list"`. Computed.

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

<a id="canonical-021bc7c6b1d50065e5a9dc4a0cb75316952ad19bda824f126340938a0d668023"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / 82f5d5c7a775 / 3

<a id="canonical-1ea2aeb98e05db4366b2d7adf2b6806ff7eea0016574d782b3dd60ae9b2fb80d"></a>

<a id="canonical-60a1e96847c7cd1efe4366e175e18569c66b6665b2dff0bb44499aa532466aac"></a>

## labels property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / 82f5d5c7a775 / 4

Type: `["map", "string"]`. Computed.

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

- [volume_defaults](data-sources--fleet--reference--group-003.md#canonical-102aa16548dfc8a51ce3d8854ef6eb7469623b5b8027c534b2381c619ab9e2fd): complete subsection reference.

<a id="canonical-c46b6c57c8d36c22cf9f1441208acfd573ffbcac7e02ff49ea9eda2a702b67c4"></a>

<a id="canonical-3bed0e1fbe19656ce4a7ff140e6b193dce5aa8de77cada5de2ed29bd71b75c61"></a>

## zone property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / 82f5d5c7a775 / 5

Type: `"string"`. Computed.

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

<a id="canonical-68a4d6eccb111eb9b01b2e8ff52398a7a1f25eb61f70dbfa5811da502edbae68"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / 82f5d5c7a775 / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults](data-sources--fleet--reference--group-003.md#canonical-102aa16548dfc8a51ce3d8854ef6eb7469623b5b8027c534b2381c619ab9e2fd)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-2beb37295fa85b4d6c4c5916001717fb8e9f448241abcb42bdc6b5678df6ed3c)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-102aa16548dfc8a51ce3d8854ef6eb7469623b5b8027c534b2381c619ab9e2fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a11a6e64adb0e003c7afbed5c764ef70ee8b7bc7bad5ba7f8576bc3dba193a1"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / bb2fdb3bae97 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-2beb37295fa85b4d6c4c5916001717fb8e9f448241abcb42bdc6b5678df6ed3c)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](data-sources--fleet--reference--group-003.md#canonical-23c6a9da3581feb8bcb034a8042769d9431330c7508961c456f622bc4646bdc0)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults

<a id="canonical-a34ddf84aace1ffcfbfe51bf78874c16d87e56cd9a3f56688169775cf5ce4334"></a>

Type: `"single"`. Computed.

It controls how each volume is provisioned by default using these OPTIONS in a special section of
the configuration.

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

<a id="canonical-78912ce77b916dc9156dfe6a8cbc1c83aebeb39a60efa67c266401beaa96afdb"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / bb2fdb3bae97 / 3

<a id="canonical-87e12be1f055c85f05af80600490ed0ab90150a7ae2847f936f719225e2c433f"></a>

<a id="canonical-318271ebd695245ca1440c09b0a3d1268e4961edcae868b435655b33391fd80e"></a>

## adaptive_qos_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / bb2fdb3bae97 / 4

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Exclusive with \[no\_qos qos\_policy\] Enter Adaptive QoS Policy Name.

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

<a id="canonical-fc03284f58ec6c78f23f815de85cee0ab46c78cb92102cced736d6a78e68a523"></a>

<a id="canonical-dc9a13cf8744e565847fa545848df7c11273da68675093eb6ea5950213e922d3"></a>

## encryption property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / bb2fdb3bae97 / 5

Type: `"bool"`. Computed.

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

<a id="canonical-9122c6b2e07d29a185a11a08b0f2f9bba37b5108230a48b382a97d0e7c6d4dca"></a>

<a id="canonical-35d1e751a46da452eb7dd720c64bd4384b34ca2f8e2a574f286bb32e907a9357"></a>

## export_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / bb2fdb3bae97 / 6

Type: `"string"`. Computed.

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

- [no_qos](data-sources--fleet--reference--group-003.md#canonical-8f6462e8be2165113934d577e06d77cb90fc355a9fb4ed7ec7decf4dacf1c596): complete subsection reference.

<a id="canonical-c4b1025e2058853902f304c8b7af0eb685e033633e1b890dceab3d70f15a50dc"></a>

<a id="canonical-4d25887a62d337a1c011037299e05e0383430955ec026de1a1238a25c2d913ef"></a>

## qos_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / bb2fdb3bae97 / 7

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Exclusive with \[adaptive\_qos\_policy no\_qos\] Enter QoS Policy Name.

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

<a id="canonical-2902d2087e75c3b6bcd5d2286d4b827656d208f63882e771393e18ace9849a06"></a>

<a id="canonical-44c9495a1239684526979f7860a8565c46d9a22086b0a7f8e565026df6a0d48d"></a>

## security_style property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / bb2fdb3bae97 / 8

Type: `"string"`. Computed.

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

<a id="canonical-5bb85f9c0584ea4025aebad3588a71473d8f1116ab1492f9a72ee8ab7c97b070"></a>

<a id="canonical-e96e51986ce2b3d845936f659c4c59a2233d92c1205f52be301c1422d8038454"></a>

## snapshot_dir property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / bb2fdb3bae97 / 9

Type: `"bool"`. Computed.

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

<a id="canonical-529526fe56b0edb077be696d667dd8083f5d8c7b5a6aa78da423decb05dabae1"></a>

<a id="canonical-b1cd1c5c3159f78d3aed394bec27f005c316e9189f82029bf7ef49ac7ce01104"></a>

## snapshot_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / bb2fdb3bae97 / 10

Type: `"string"`. Computed.

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

<a id="canonical-82c0cfb9cfd375038247c5cdb5074396830e125d908e31903392815ba4943451"></a>

<a id="canonical-5d19e24f7a82998fe7aae69b93e46877b6a8f1be2b502ade7b4fe4db0734a6bb"></a>

## snapshot_reserve property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / bb2fdb3bae97 / 11

Type: `"string"`. Computed.

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

<a id="canonical-c22e3cada904e2f4a5aa2b4bc8390f2fcc0e1f0bd17322cba291cd74cc03e34a"></a>

<a id="canonical-184e964d56401731a8666d934210a65f640de52e8bec2eba9ccea3f28f03a12e"></a>

## space_reserve property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / bb2fdb3bae97 / 12

Type: `"string"`. Computed.

\[Enum: none|thick\] Space reservation mode; “none” (thin) or “volume” (thick). Possible values are
\`none\`, \`thick\`.

Upstream description:

Space reservation mode; “none” (thin) or “volume” (thick)

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

<a id="canonical-b8dfd772563349242c46f26ba544931fd13fd978853ddb96f2a58f21c1366617"></a>

<a id="canonical-9aa44b15012acf24d08294bb12ca4d565b4f02647597877e55abf38b0f844110"></a>

## split_on_clone property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / bb2fdb3bae97 / 13

Type: `"bool"`. Computed.

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

<a id="canonical-4ccf1e7332551c3b49a505028c54d4613af75ef88ddc5f40f28c9966b178eb42"></a>

<a id="canonical-5fe5a074ba281c08d3d2d39f3b1accaddf630a3511a2e05c661aa1584ae06c66"></a>

## tiering_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / bb2fdb3bae97 / 14

Type: `"string"`. Computed.

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

<a id="canonical-86c2dd40c983a1efb7d8f2e59c0df79e865b6164224adac35b9dcc3e6255c5c1"></a>

<a id="canonical-8e57329f30590af84c6ed8b9df7bfba68900a8d729884977bb6a9cf15ccbc80f"></a>

## unix_permissions property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / bb2fdb3bae97 / 15

Type: `"number"`. Computed.

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

<a id="canonical-c6b1c8ee88e59660e1385cff04a93de1ecdcaddae64fb1c9a337ba18b1f194bd"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / bb2fdb3bae97 / 16

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos](data-sources--fleet--reference--group-003.md#canonical-8f6462e8be2165113934d577e06d77cb90fc355a9fb4ed7ec7decf4dacf1c596)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](data-sources--fleet--reference--group-003.md#canonical-23c6a9da3581feb8bcb034a8042769d9431330c7508961c456f622bc4646bdc0)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-8f6462e8be2165113934d577e06d77cb90fc355a9fb4ed7ec7decf4dacf1c596"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2a2ad40cb860313479050b4d0307412f88fdf6ef63930df029fff2df88fa4bd2"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / a14e6d0f60a6 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-2beb37295fa85b4d6c4c5916001717fb8e9f448241abcb42bdc6b5678df6ed3c)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](data-sources--fleet--reference--group-003.md#canonical-23c6a9da3581feb8bcb034a8042769d9431330c7508961c456f622bc4646bdc0)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults](data-sources--fleet--reference--group-003.md#canonical-102aa16548dfc8a51ce3d8854ef6eb7469623b5b8027c534b2381c619ab9e2fd)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos

<a id="canonical-9ede7a09eebe38c255849beee0839423ff27675cde2c591c212deac82788f264"></a>

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

<a id="canonical-f602ab3b0bd94edac727e2e629a7e75910cf5b61040d05b65c98025ab14cc139"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / a14e6d0f60a6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bcce6e01154d5db1e76794b1e229defcef5db8a5bffc6b76154119d9a888113f"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.stor / a14e6d0f60a6 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults](data-sources--fleet--reference--group-003.md#canonical-102aa16548dfc8a51ce3d8854ef6eb7469623b5b8027c534b2381c619ab9e2fd)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-ceb2d887e3260398c1a5ef49341011226291126e755c1394ad71a3ba4b8c46ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b8f36f19caea5329d1348a62d801894332cea2efea86d75a7b0a2ebeae5b68bd"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 1058e9561673 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-2beb37295fa85b4d6c4c5916001717fb8e9f448241abcb42bdc6b5678df6ed3c)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap

<a id="canonical-ca3a2b8dbc39e102a86a9168bbd20eb8bba941a3f223ab49cb16bd55a3b44714"></a>

Type: `"single"`. Computed.

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

<a id="canonical-370c1da11e5f8cbfbf0f5a4036e747bf9534cb241045d7654f36a2c0debec394"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 1058e9561673 / 3

- [chap_initiator_secret](data-sources--fleet--reference--group-003.md#canonical-aaca66e74ec3dc5ae435b6b984db5fd57bf7b89d13e001c2b49ac21864d5f17a): complete subsection reference.

- [chap_target_initiator_secret](data-sources--fleet--reference--group-003.md#canonical-f4fa78d98ebd3ddd15c42fdcfb9b1a71f8b0cf708f9325e55b5d413f74f2da0c): complete subsection reference.

<a id="canonical-9deb5ac36903f1e0c73541bae7b39e94ed702d5200f70744fdaa34fc50903671"></a>

<a id="canonical-8aeea85dc58289a1449b08b7f49d5a89a5e2907c32a2282300d81844bdb7ceb8"></a>

## chap_target_username property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 1058e9561673 / 4

Type: `"string"`. Computed.

Target username. Required if useCHAP=true.

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

<a id="canonical-dfbceb276a530b231b94fdcb5889f461c94e27773784a7a0816a01b0350eeca1"></a>

<a id="canonical-c373a515e4508f1bef5380a881a3296ac36dea511c6cb5738990ecd3bc6fcf3f"></a>

## chap_username property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 1058e9561673 / 5

Type: `"string"`. Computed.

Inbound username. Required if useCHAP=true.

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

<a id="canonical-b3b9d35f6912fddde2016a4af8477027186dfdfc523f20772a685f3e41fac16b"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 1058e9561673 / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](data-sources--fleet--reference--group-003.md#canonical-aaca66e74ec3dc5ae435b6b984db5fd57bf7b89d13e001c2b49ac21864d5f17a)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](data-sources--fleet--reference--group-003.md#canonical-f4fa78d98ebd3ddd15c42fdcfb9b1a71f8b0cf708f9325e55b5d413f74f2da0c)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-2beb37295fa85b4d6c4c5916001717fb8e9f448241abcb42bdc6b5678df6ed3c)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-aaca66e74ec3dc5ae435b6b984db5fd57bf7b89d13e001c2b49ac21864d5f17a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-071f93891d12f051ac007e46fb4cbf2f41c42d5d158d5afd7db0d618d635aec9"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / c6832b23c032 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-2beb37295fa85b4d6c4c5916001717fb8e9f448241abcb42bdc6b5678df6ed3c)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--fleet--reference--group-003.md#canonical-ceb2d887e3260398c1a5ef49341011226291126e755c1394ad71a3ba4b8c46ab)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret

<a id="canonical-596b557f8442c9feb215ee25e818b91644d047841d50d8114851ba3d89731cd3"></a>

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

<a id="canonical-8d53c611ba27b15aada2990bebf9f92b6a22b9b8b91ef5cede0a4dcd74abbc37"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / c6832b23c032 / 3

- [blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-8894ce0e3a137cdc10a8cddb8e3c16c5674fc0fd78b10d661d71ea30494a31f1): complete subsection reference.

- [clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-16c42582a608fcb6d125ea8012e381d9b3bac01cc19531c98ed485f302897c3e): complete subsection reference.

<a id="canonical-10eeacea8e52f208ea279a713dde57d49ed1bfd2befd96969bbc07d735d2f07e"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / c6832b23c032 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-8894ce0e3a137cdc10a8cddb8e3c16c5674fc0fd78b10d661d71ea30494a31f1)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-16c42582a608fcb6d125ea8012e381d9b3bac01cc19531c98ed485f302897c3e)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--fleet--reference--group-003.md#canonical-ceb2d887e3260398c1a5ef49341011226291126e755c1394ad71a3ba4b8c46ab)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-8894ce0e3a137cdc10a8cddb8e3c16c5674fc0fd78b10d661d71ea30494a31f1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a39318f4b186702cae5594011c0104a137545041933399d64bfa1e3fe5072685"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / e7d4a69bc4f6 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-2beb37295fa85b4d6c4c5916001717fb8e9f448241abcb42bdc6b5678df6ed3c)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--fleet--reference--group-003.md#canonical-ceb2d887e3260398c1a5ef49341011226291126e755c1394ad71a3ba4b8c46ab)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](data-sources--fleet--reference--group-003.md#canonical-aaca66e74ec3dc5ae435b6b984db5fd57bf7b89d13e001c2b49ac21864d5f17a)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info

<a id="canonical-64d88a9a3548696f090cd92c1da293a357460e7996d9932616c1be891c9abe9b"></a>

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

<a id="canonical-bc9e09812ed8fbbb979445a8d5badd9182b0105ccdffb4a5d5d2217ba14849cb"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / e7d4a69bc4f6 / 3

<a id="canonical-58a1a3b2ceb4a361877cb24a08247787861f85b0f4b933bfb913c1e09d6eefaa"></a>

<a id="canonical-49cb18e805e04619fb5996664e82325d8fa95712ef49b3d60a11de6571859d8b"></a>

## decryption_provider property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / e7d4a69bc4f6 / 4

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

<a id="canonical-71ab2e324f4e1c4407380397bf78908a94d69c499b73aeb0ff79ef6e84642a2d"></a>

<a id="canonical-e686c5508ca33d4b94e87caded1c1d765f6c2a338a4eb02d724d0e7b5c675912"></a>

## location property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / e7d4a69bc4f6 / 5

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

<a id="canonical-df964a94735ceef8c964ba63a6c2b6fb3bd4bde768da4ac03893d1ccae81ca2c"></a>

<a id="canonical-0d2917868fc915f6afe4a375bb1e511e7f0a071cf64431da9e2de2c678466e56"></a>

## store_provider property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / e7d4a69bc4f6 / 6

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

<a id="canonical-24f1d4739a5a46a89e0708233ed0a9c83f39efb07c25422867d9a7141441818b"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / e7d4a69bc4f6 / 7

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](data-sources--fleet--reference--group-003.md#canonical-aaca66e74ec3dc5ae435b6b984db5fd57bf7b89d13e001c2b49ac21864d5f17a)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-16c42582a608fcb6d125ea8012e381d9b3bac01cc19531c98ed485f302897c3e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d3333ac5ad8b33d960eaca58d2b8671ffba715bfbc4733130f2800ef09ac2cd5"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 10ed8c8d1959 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-2beb37295fa85b4d6c4c5916001717fb8e9f448241abcb42bdc6b5678df6ed3c)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--fleet--reference--group-003.md#canonical-ceb2d887e3260398c1a5ef49341011226291126e755c1394ad71a3ba4b8c46ab)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](data-sources--fleet--reference--group-003.md#canonical-aaca66e74ec3dc5ae435b6b984db5fd57bf7b89d13e001c2b49ac21864d5f17a)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info

<a id="canonical-ab2262bddd3fc16fee31169e7f31e6890d7d830f8687150906694245c330ac8d"></a>

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

<a id="canonical-8aa4fe232855a8c8d7640e71c81bf36809f321f5abc51408d2d5426998ad8746"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 10ed8c8d1959 / 3

<a id="canonical-d70c595b3c18e08dd7c4c0e5ec73f88a0c6cc4889b3759afd6c26d511d0587ff"></a>

<a id="canonical-505de6bb2baa4a45d063c30da63c263f87102d21af5a9feeecac4b55941d18a6"></a>

## provider_ref property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 10ed8c8d1959 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-291e976e774bd6adba46199f38f4471cc2f1e3c5dc2b59ad6059b6f93499b499"></a>

<a id="canonical-7e543d94fd260b688eadb180420c0113cf28a59d525825c89586e268393d2e2f"></a>

## url property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 10ed8c8d1959 / 5

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

<a id="canonical-6e3b9b50e2317f665de4ada4959247da317413358f9f391c84bee8528172c6d4"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 10ed8c8d1959 / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](data-sources--fleet--reference--group-003.md#canonical-aaca66e74ec3dc5ae435b6b984db5fd57bf7b89d13e001c2b49ac21864d5f17a)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-f4fa78d98ebd3ddd15c42fdcfb9b1a71f8b0cf708f9325e55b5d413f74f2da0c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa12b1904bc52cc1703a22122e195bdd2e04466b1290beb7f9f87899e8e6f308"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / d0ef1153819d / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-2beb37295fa85b4d6c4c5916001717fb8e9f448241abcb42bdc6b5678df6ed3c)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--fleet--reference--group-003.md#canonical-ceb2d887e3260398c1a5ef49341011226291126e755c1394ad71a3ba4b8c46ab)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret

<a id="canonical-953a2c3e224fb40d2b23c87f4535683e3f762c631fc423c1d2601794dadacd8f"></a>

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

<a id="canonical-7ace8b680e2020eb8424018746509dfd1980c55f4779eca3715cc1d2c6ca17ba"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / d0ef1153819d / 3

- [blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-fae6cf48104142dd152d172ee7dab8f2a1697e04ac08ca90b6a4835045f0bb5e): complete subsection reference.

- [clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-7429800e6bd0c1d5ae7905ada76d62a98614d30de932e92f4cd99b7c24c384b1): complete subsection reference.

<a id="canonical-c2f6f04126dfcf14f5080916b85b7d7ea4e79b57228df9955978519af0e9df9c"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / d0ef1153819d / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-fae6cf48104142dd152d172ee7dab8f2a1697e04ac08ca90b6a4835045f0bb5e)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-7429800e6bd0c1d5ae7905ada76d62a98614d30de932e92f4cd99b7c24c384b1)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--fleet--reference--group-003.md#canonical-ceb2d887e3260398c1a5ef49341011226291126e755c1394ad71a3ba4b8c46ab)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-fae6cf48104142dd152d172ee7dab8f2a1697e04ac08ca90b6a4835045f0bb5e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-48ca0f5c6424b433e07ddcbcc3978b59aae5d45fc43cf617c7649af319b83fd1"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 81c5dedd1fd8 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-2beb37295fa85b4d6c4c5916001717fb8e9f448241abcb42bdc6b5678df6ed3c)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--fleet--reference--group-003.md#canonical-ceb2d887e3260398c1a5ef49341011226291126e755c1394ad71a3ba4b8c46ab)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](data-sources--fleet--reference--group-003.md#canonical-f4fa78d98ebd3ddd15c42fdcfb9b1a71f8b0cf708f9325e55b5d413f74f2da0c)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info

<a id="canonical-b5437f38b1fb602a1e1881a8f658e9303597f88f5b28838739e2cb1f1f908dcd"></a>

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

<a id="canonical-c673c2dabd348303ccae22524b14a68171a27fac7239d8bdc642ab601fdca2a7"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 81c5dedd1fd8 / 3

<a id="canonical-a5e1c5745790f991cc48ec151d723c2dd59323e9259d69ef1b052af8d78f3c34"></a>

<a id="canonical-79d2ebb54cb06ed97cd7b7d3f525a90c9ad41fe3751f8e19fcbd2103c881dc07"></a>

## decryption_provider property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 81c5dedd1fd8 / 4

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

<a id="canonical-46964e747974f14a30e97b6fe4f4f2c581a53f1c67290ba548e74f55ceb46de4"></a>

<a id="canonical-8e6264f1e6b30ad50116a2794d1c9158f9eec671eb77ffecece7ba95006754e0"></a>

## location property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 81c5dedd1fd8 / 5

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

<a id="canonical-f477c5217c9709720629ce5f859c758e83fe5d3b5af235ecea8312cc4729717e"></a>

<a id="canonical-2d788e951ead34f929bb98978e181caf69a26841843b1bae4b3dcf4980064f8b"></a>

## store_provider property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 81c5dedd1fd8 / 6

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

<a id="canonical-734b039e29e04cc715c9c0532f66a05ee828b6fdffc1fa13714e582062a9b379"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 81c5dedd1fd8 / 7

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](data-sources--fleet--reference--group-003.md#canonical-f4fa78d98ebd3ddd15c42fdcfb9b1a71f8b0cf708f9325e55b5d413f74f2da0c)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-7429800e6bd0c1d5ae7905ada76d62a98614d30de932e92f4cd99b7c24c384b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7924cb029e81afcc25591ae3ff43a267126741a489de38cfd5ea6e001061614a"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 4b38dca39105 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-2beb37295fa85b4d6c4c5916001717fb8e9f448241abcb42bdc6b5678df6ed3c)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--fleet--reference--group-003.md#canonical-ceb2d887e3260398c1a5ef49341011226291126e755c1394ad71a3ba4b8c46ab)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](data-sources--fleet--reference--group-003.md#canonical-f4fa78d98ebd3ddd15c42fdcfb9b1a71f8b0cf708f9325e55b5d413f74f2da0c)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info

<a id="canonical-4500af7aa2fedb643cdf601ef719a122ef9bb004eb7c563c7aabee907db04a36"></a>

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

<a id="canonical-3a1582dfd0a430ea13b6caf145cc3c2cf3461e1b1ebd623ea28d08638180355c"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 4b38dca39105 / 3

<a id="canonical-dc901a3d2f9dc55d0661eef73597e33850a0f19848cd7ba6f5bde040b8895392"></a>

<a id="canonical-07ff3acddaf455032b651caf4bece917dd487115d2b475a442bac21b57e36e71"></a>

## provider_ref property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 4b38dca39105 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-8752a92467c1160daf47536314d0db90753211af1bf1c13cf50f95d1ef51335d"></a>

<a id="canonical-9b64c2c7a53124bcbc5b324c29513298479351811c9f78689d82c5418b1ea399"></a>

## url property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 4b38dca39105 / 5

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

<a id="canonical-28721969c5d754ed104fc76a4e1f69fb74b5ddcee546ec1d59ab8b314b53de87"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_ / 4b38dca39105 / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](data-sources--fleet--reference--group-003.md#canonical-f4fa78d98ebd3ddd15c42fdcfb9b1a71f8b0cf708f9325e55b5d413f74f2da0c)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-36b7100b6badb73b70262a1760b03aec196f91b255873a3ac1ad65e398eec79c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b66564d49263c3823ded9c81509ea3ca7790fbb3530f0c1704d59b3959174d0"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 45ab5d0f3438 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-2beb37295fa85b4d6c4c5916001717fb8e9f448241abcb42bdc6b5678df6ed3c)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults

<a id="canonical-7e60ccd045312c0d783d450da947be77416982e25accf6a985e24594f1ca6bf8"></a>

Type: `"single"`. Computed.

It controls how each volume is provisioned by default using these OPTIONS in a special section of
the configuration.

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

<a id="canonical-38e355ccc94c4a81af1f58f8f3f9a4899d07e53c2118bd750a439657eae57f9e"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 45ab5d0f3438 / 3

<a id="canonical-2bdaca4dab9daacce16edbd2659641c00046dbcf5a5af4c9257933cbccf1b6dd"></a>

<a id="canonical-0fc2016e1cc4ff4a40196517c33d84d29fc4bf3f3a4ab01950f91cd8c9fe27db"></a>

## adaptive_qos_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 45ab5d0f3438 / 4

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Exclusive with \[no\_qos qos\_policy\] Enter Adaptive QoS Policy Name.

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

<a id="canonical-c5d08422c13ec4b9c47ed41c98d796a5810f9fde0c71fc4b8bd17c6c5bd5ce2d"></a>

<a id="canonical-defb2f6e92cc245b65009ea0ce98627417254c5782a92b9a3e27a75cf29cdac5"></a>

## encryption property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 45ab5d0f3438 / 5

Type: `"bool"`. Computed.

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

<a id="canonical-7ddc801abea721291801ceaf9612c29a2abd7d96fd6c1b904fa810181d08e1ab"></a>
