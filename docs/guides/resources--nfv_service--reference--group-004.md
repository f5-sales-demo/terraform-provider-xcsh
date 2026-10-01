---
page_title: "xcsh_nfv_service reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nfv_service reference."
---

# xcsh_nfv_service reference

<a id="canonical-8f924ab6dbbe68bd9352fc26f8458e1e0b2d1363735024ec7191e8184bc6f7b4"></a>

## Direct properties — palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key / 2792bbab69c8 / 3

- [blindfold_secret_info](resources--nfv_service--reference--group-004.md#canonical-2c2ada6499df9557ad3b0e833ad8a6dbeebb7c31317f14a47f3b456090f2c5ea): complete subsection reference.

- [clear_secret_info](resources--nfv_service--reference--group-004.md#canonical-c1e9fe933c9e00f96834177d2b5220cd4e0e173191bfa99038ed0c7197dc7734): complete subsection reference.

<a id="canonical-74b2de8f3b07c1a35ab56caa06a410e83b42a11fcefff1c3e9d96dd009f1f93d"></a>

## Next pages — palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key / 2792bbab69c8 / 4

- [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info](resources--nfv_service--reference--group-004.md#canonical-2c2ada6499df9557ad3b0e833ad8a6dbeebb7c31317f14a47f3b456090f2c5ea)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info](resources--nfv_service--reference--group-004.md#canonical-c1e9fe933c9e00f96834177d2b5220cd4e0e173191bfa99038ed0c7197dc7734)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys](resources--nfv_service--reference--group-003.md#canonical-0fe0c13a624ad3fd90ff98e5424a8a2a934d3c04471ee911e58bc378984387a1)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-2c2ada6499df9557ad3b0e833ad8a6dbeebb7c31317f14a47f3b456090f2c5ea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-007dc9bd73e0d3ea11951265d7dd4ae5ee7b1989a4525702c1103b14fcfd8680"></a>

## palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info — palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_inf / 2982f9e92063 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-683b0e043ad34e16d220e00d2c0ad7d7794cd7a996943065814c81f0d29e03cf)
- [palo_alto_fw_service.auto_setup](resources--nfv_service--reference--group-003.md#canonical-7bf638e8f5f6d4cd41bc1908f771ac84af047472fffcf4f680086848cf0d43db)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys](resources--nfv_service--reference--group-003.md#canonical-0fe0c13a624ad3fd90ff98e5424a8a2a934d3c04471ee911e58bc378984387a1)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key](resources--nfv_service--reference--group-003.md#canonical-045c0321536770858b03978ae68fe304d694d5aec0fc64627c72a66bcbbc5095)
- palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info

<a id="canonical-6ba4f9d338374a9feff11dd54b618f9278ce692bd0347b867eb3040309b2ad65"></a>

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

<a id="canonical-e95b5a73bfe5cc3f8b0c16e744781c4c8d5719fd0419583ebad82b79df619a36"></a>

## Direct properties — palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_inf / 2982f9e92063 / 3

<a id="canonical-47b1187f3dcda3b6fa162b1b6aab769fdfab3b49fa3c7e0ff7bd4422181233df"></a>

<a id="canonical-5c05679ef87b7a6ba37cf972e7351263b4eb1c5f1e0d6bc1533d4ceed05f3d30"></a>

## decryption_provider property — palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_inf / 2982f9e92063 / 4

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

<a id="canonical-7416049b2a6dddd35de75eb2fc9cb35c9b5afd41f80ad369b4d9701154f2143c"></a>

<a id="canonical-319304b94b38da4c1bcbfa063158b866731b60d4ec60dede3963e2ba39ea076f"></a>

## location property — palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_inf / 2982f9e92063 / 5

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

<a id="canonical-c3b1ce4b8c312640ac81f0627c765be7cac34718a919dbee76834304d22c743c"></a>

<a id="canonical-42d71528696d826ac94f282d2fd9eadd7f38c49b198507e163e741eb02bb160f"></a>

## store_provider property — palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_inf / 2982f9e92063 / 6

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

<a id="canonical-bf2822d79a34cb9d94aee001ba44fa529645e66b92af86fffd4a8832d8be4995"></a>

## Next pages — palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_inf / 2982f9e92063 / 7

- [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key](resources--nfv_service--reference--group-003.md#canonical-045c0321536770858b03978ae68fe304d694d5aec0fc64627c72a66bcbbc5095)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-c1e9fe933c9e00f96834177d2b5220cd4e0e173191bfa99038ed0c7197dc7734"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0c15a75c93779da8437aad38c9c36e4411d3d6c691cf94b273dc275a1a1aeda5"></a>

## palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info — palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info / a610d6c04090 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-683b0e043ad34e16d220e00d2c0ad7d7794cd7a996943065814c81f0d29e03cf)
- [palo_alto_fw_service.auto_setup](resources--nfv_service--reference--group-003.md#canonical-7bf638e8f5f6d4cd41bc1908f771ac84af047472fffcf4f680086848cf0d43db)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys](resources--nfv_service--reference--group-003.md#canonical-0fe0c13a624ad3fd90ff98e5424a8a2a934d3c04471ee911e58bc378984387a1)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key](resources--nfv_service--reference--group-003.md#canonical-045c0321536770858b03978ae68fe304d694d5aec0fc64627c72a66bcbbc5095)
- palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info

<a id="canonical-d62311ee0f6873838214453f8462545c1e2ca30eb4ebb1a54a35cfdc69abfb0f"></a>

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

<a id="canonical-44f98ac9764af2e13479c023a2af31e120d64023beae93c93ac10df546601667"></a>

## Direct properties — palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info / a610d6c04090 / 3

<a id="canonical-4aae2f48d1e4ea7240a685c3946b5bdd9398406a6dcfaf96f73bc050ff5ea8d2"></a>

<a id="canonical-6740cb473559413d89eeba5efe004dc848ed3f0d4def617f265514984b001df5"></a>

## provider_ref property — palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info / a610d6c04090 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-ce4b7b21a6a3a1ce365da0e667c983638ee14d9e42671a1142088fdbd8cd07eb"></a>

<a id="canonical-f4a978c92ff529e498be27a126a3d861dc693f795623f76b52ed4762ef492196"></a>

## url property — palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info / a610d6c04090 / 5

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

<a id="canonical-98e96eb0f90845c1418e40f138271750ee0264d1f1075b93c36ec242a4472131"></a>

## Next pages — palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info / a610d6c04090 / 6

- [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key](resources--nfv_service--reference--group-003.md#canonical-045c0321536770858b03978ae68fe304d694d5aec0fc64627c72a66bcbbc5095)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-7154c513294dcde83423cb49e7c2c9583005df7a729e04e7b041b85d0a6fad29"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2369421710c1dad2d99824bffec6a90d6294faba614c783afeaf0ee726b5809c"></a>

## palo_alto_fw_service.aws_tgw_site — palo_alto_fw_service.aws_tgw_site / 7688ea1f5523 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-683b0e043ad34e16d220e00d2c0ad7d7794cd7a996943065814c81f0d29e03cf)
- palo_alto_fw_service.aws_tgw_site

<a id="canonical-06fb95e7aa6dede5ce324762abc7457cb971e96dbb1a76fc9e050b85740c49aa"></a>

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
aws_tgw_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-b0d599b6a67a5879876639f8b3c92467642ea95df0a6fc313f4817d5065adde7"></a>

## Direct properties — palo_alto_fw_service.aws_tgw_site / 7688ea1f5523 / 3

<a id="canonical-6c7e21ed60de62cc861575ea2d4442dc195ebb2d7ac5ced1c80a0b7d5ecbe3f5"></a>

<a id="canonical-7698a8c06be48b34dde3220de5261a1a84140763b60eee0f929a41b595b8a3e2"></a>

## name property — palo_alto_fw_service.aws_tgw_site / 7688ea1f5523 / 4

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

<a id="canonical-dad0fcf09c249a25fd60dbc0ed2e6a60a49d44194e646513a941bb79110bc396"></a>

<a id="canonical-846f887e8531c2981cbedf1fff908e2cdf14faee6eedbb7c2ec1588b5fab6c3f"></a>

## namespace property — palo_alto_fw_service.aws_tgw_site / 7688ea1f5523 / 5

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

<a id="canonical-9ade89855a8cf733c131fd48ce9a7fd7bb176c84b468369acaddbc0fd15b7259"></a>

<a id="canonical-ad7c8600344c08c3a5b7e909e89aaf9f7e6a6e5c4dcebe09f368b2420a13386a"></a>

## tenant property — palo_alto_fw_service.aws_tgw_site / 7688ea1f5523 / 6

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

<a id="canonical-ec9f513f363cde5e7655805cec94bbd426026e6882b2b4cefcbdceca3aa1c969"></a>

## Next pages — palo_alto_fw_service.aws_tgw_site / 7688ea1f5523 / 7

- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-683b0e043ad34e16d220e00d2c0ad7d7794cd7a996943065814c81f0d29e03cf)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-5b3a157c00b3631985c671ca0652d9b5bbacb64e999edb8b5e787209de0a1142"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c3d949399f6ebbed3a21b8476b9ee7bc6156af24ab7176d66e9d83fd370ae660"></a>

## palo_alto_fw_service.disable_panaroma — palo_alto_fw_service.disable_panaroma / 1291c8f023dc / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-683b0e043ad34e16d220e00d2c0ad7d7794cd7a996943065814c81f0d29e03cf)
- palo_alto_fw_service.disable_panaroma

<a id="canonical-c21e39b36de4a2e3578d2cd7a8fd908cd26c30e1d23dde0232d2b64dc0ae6670"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_panaroma = {}
```

<a id="canonical-765b4f50f5efa063770da05febf9c0531f8c4a0132050dc621b3e31102829140"></a>

## Direct properties — palo_alto_fw_service.disable_panaroma / 1291c8f023dc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6798f63b8a72f16c9ab61156006b912b114d6f791f3d547d2b8326ad96d7405f"></a>

## Next pages — palo_alto_fw_service.disable_panaroma / 1291c8f023dc / 4

- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-683b0e043ad34e16d220e00d2c0ad7d7794cd7a996943065814c81f0d29e03cf)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-cff2ad597e13c8f342ca161809452d25cd1d6d3a68e3cf59e6f5316bfe455e4c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fe3d3c5aba2194969b867f4a365685b9315052032fcdcfe4cc8c3c8aa499a383"></a>

## palo_alto_fw_service.pan_ami_bundle1 — palo_alto_fw_service.pan_ami_bundle1 / ccbee1dfb686 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-683b0e043ad34e16d220e00d2c0ad7d7794cd7a996943065814c81f0d29e03cf)
- palo_alto_fw_service.pan_ami_bundle1

<a id="canonical-7a56bae860591078e7e5798b9521b4de14942370395a086b03aec160d73d22e7"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
pan_ami_bundle1 = {}
```

<a id="canonical-c98732ec505019ea36fe81a78ff0306f1f2329feb6b4457eada55fad9f82b1e1"></a>

## Direct properties — palo_alto_fw_service.pan_ami_bundle1 / ccbee1dfb686 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5c5d5fc37d125ee656efbb0bc20fce866b328bd32a04ffef06b04cba19f87caf"></a>

## Next pages — palo_alto_fw_service.pan_ami_bundle1 / ccbee1dfb686 / 4

- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-683b0e043ad34e16d220e00d2c0ad7d7794cd7a996943065814c81f0d29e03cf)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-65382323b951c01510ff5617d2dcc84b1d074913e191352eb24b3b43e3471040"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b9f9499756e2e6b0e2dc61fab5e86cd2110a4d31f3092739e0c63a69fcb7a61"></a>

## palo_alto_fw_service.pan_ami_bundle2 — palo_alto_fw_service.pan_ami_bundle2 / 81c6e774e8ca / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-683b0e043ad34e16d220e00d2c0ad7d7794cd7a996943065814c81f0d29e03cf)
- palo_alto_fw_service.pan_ami_bundle2

<a id="canonical-5a14a412eb31509e25439ce526615379bd5cbea702ce2ed8462aff74dbd96cb6"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
pan_ami_bundle2 = {}
```

<a id="canonical-2cea6fd5f24b017ef195051b7d8010b514e119f93a46eff9ca4679da2013063b"></a>

## Direct properties — palo_alto_fw_service.pan_ami_bundle2 / 81c6e774e8ca / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1539d55b40d1d809bfa6e88dab331ff086fb5bc89534a87cfce628cb5589ae45"></a>

## Next pages — palo_alto_fw_service.pan_ami_bundle2 / 81c6e774e8ca / 4

- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-683b0e043ad34e16d220e00d2c0ad7d7794cd7a996943065814c81f0d29e03cf)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-45ec14f98c4298e49ab80b0eab6e4d912220244b1166f017f7ae9536733a9395"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a3b29559c4ea4780847858ade3b00c228ae7ab4ef81610dd7ab3c7c12872673"></a>

## palo_alto_fw_service.panorama_server — palo_alto_fw_service.panorama_server / e132d903b8ea / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-683b0e043ad34e16d220e00d2c0ad7d7794cd7a996943065814c81f0d29e03cf)
- palo_alto_fw_service.panorama_server

<a id="canonical-d367e7f39b3d910eb1713b0fe454f2d55c32800fb9a045a277dce3704693f29a"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for panorama server.

Upstream description:

Panorama Server Type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("server")}
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
panorama_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-d56c283e99febc1165cbd6c32a9aa5a148296499f38a753bfff30f429fb9f186"></a>

## Direct properties — palo_alto_fw_service.panorama_server / e132d903b8ea / 3

- [authorization_key](resources--nfv_service--reference--group-004.md#canonical-eb18479257d1b0b1b1334ce2a9daa05da3175e89c1e2daca86909038a18753c2): complete subsection reference.

<a id="canonical-d30d7aa5f847d278617534c713f0099cb9bd788f9ca76d2482bdc9183639f072"></a>

<a id="canonical-207ed4394db2639476c4eebb2bb9d90fae8708733f28c3a29e94509afa8aa35d"></a>

## device_group_name property — palo_alto_fw_service.panorama_server / e132d903b8ea / 4

Type: `"string"`. Optional.

Device Group Name. Device Group Name.

Upstream description:

Device Group Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-6e913aebe09c93d7062936b5fa3ba9bc357a1bd1fa8ec9a426e18deddde65238"></a>

<a id="canonical-4d272b0bb95ef14b5e523318ebec873dbe763bee7cde49f56bfc196b2d3efd7a"></a>

## server property — palo_alto_fw_service.panorama_server / e132d903b8ea / 5

Type: `"string"`. Optional.

Panorama Server Address to which the firewall should connect to.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-0bd3cc822ce9dce4c2f4bae5dab7910249e40641ef3795165276944d8a7ccbdc"></a>

<a id="canonical-9234d4aa1bcfaac9818552335c6bb59ab43676e2405aa998c696da05ee2cdf6b"></a>

## template_stack_name property — palo_alto_fw_service.panorama_server / e132d903b8ea / 6

Type: `"string"`. Optional.

Template stack name. Template Stack Name.

Upstream description:

Template Stack Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-a5fdf6fe8cbe5d21735d3f0bc010db17acda389e6317523529a69a2f68ac1b5c"></a>

## Next pages — palo_alto_fw_service.panorama_server / e132d903b8ea / 7

- [palo_alto_fw_service.panorama_server.authorization_key](resources--nfv_service--reference--group-004.md#canonical-eb18479257d1b0b1b1334ce2a9daa05da3175e89c1e2daca86909038a18753c2)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-683b0e043ad34e16d220e00d2c0ad7d7794cd7a996943065814c81f0d29e03cf)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-eb18479257d1b0b1b1334ce2a9daa05da3175e89c1e2daca86909038a18753c2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f80f1f583adb475c07dd3c9b3802cc4a608994f532d2d5d452941f9ef414838b"></a>

## palo_alto_fw_service.panorama_server.authorization_key — palo_alto_fw_service.panorama_server.authorization_key / 7406adb2394b / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-683b0e043ad34e16d220e00d2c0ad7d7794cd7a996943065814c81f0d29e03cf)
- [palo_alto_fw_service.panorama_server](resources--nfv_service--reference--group-004.md#canonical-45ec14f98c4298e49ab80b0eab6e4d912220244b1166f017f7ae9536733a9395)
- palo_alto_fw_service.panorama_server.authorization_key

<a id="canonical-4c0d99a8630ee8137ec0fa6bb6e1cbdfc41159d590e4351a19562837b3383a62"></a>

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
authorization_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-7ed86811c9a5fc9258889b07e533f21ed6bd4221e8a83edaeabaa5423748858e"></a>

## Direct properties — palo_alto_fw_service.panorama_server.authorization_key / 7406adb2394b / 3

- [blindfold_secret_info](resources--nfv_service--reference--group-004.md#canonical-ff475424dcda782e193c0b01ca9112420b918c5329c79fc2d0e82c0fabd7fdb4): complete subsection reference.

- [clear_secret_info](resources--nfv_service--reference--group-004.md#canonical-6d1b5a04798f2f144f8735d09925d645a613ff5a9be09ad7f5d59a14bf04a729): complete subsection reference.

<a id="canonical-d8f4791e4632a49781423ab87cddb140d85d8349d26c8b080ec177cc5532dc59"></a>

## Next pages — palo_alto_fw_service.panorama_server.authorization_key / 7406adb2394b / 4

- [palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info](resources--nfv_service--reference--group-004.md#canonical-ff475424dcda782e193c0b01ca9112420b918c5329c79fc2d0e82c0fabd7fdb4)
- [palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info](resources--nfv_service--reference--group-004.md#canonical-6d1b5a04798f2f144f8735d09925d645a613ff5a9be09ad7f5d59a14bf04a729)
- [palo_alto_fw_service.panorama_server](resources--nfv_service--reference--group-004.md#canonical-45ec14f98c4298e49ab80b0eab6e4d912220244b1166f017f7ae9536733a9395)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-ff475424dcda782e193c0b01ca9112420b918c5329c79fc2d0e82c0fabd7fdb4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b8b7f2e822de3464a519dc5a8528dfb14ad7f07bd2ad25efcc764d899eeed73b"></a>

## palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info — palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info / 361b34fea3e2 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-683b0e043ad34e16d220e00d2c0ad7d7794cd7a996943065814c81f0d29e03cf)
- [palo_alto_fw_service.panorama_server](resources--nfv_service--reference--group-004.md#canonical-45ec14f98c4298e49ab80b0eab6e4d912220244b1166f017f7ae9536733a9395)
- [palo_alto_fw_service.panorama_server.authorization_key](resources--nfv_service--reference--group-004.md#canonical-eb18479257d1b0b1b1334ce2a9daa05da3175e89c1e2daca86909038a18753c2)
- palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info

<a id="canonical-6428b5cf3cfe6a0473b99878f749bc3ff0d7f43c2f85415123ee46d938522dc9"></a>

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

<a id="canonical-9749d98c3b786b69207879a1f5320e87a5b5673c3374932f7b855157e4bf33ae"></a>

## Direct properties — palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info / 361b34fea3e2 / 3

<a id="canonical-971413d99e57529afa4bc414f63e860e6df5bed0d302b69800e3b12f111bde69"></a>

<a id="canonical-9c0b5d26bf3ec77898021c6d8ab1fc191c35065eec164c02bd99769e8dfb0680"></a>

## decryption_provider property — palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info / 361b34fea3e2 / 4

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

<a id="canonical-8146d74bcd2858d1dab58432c331441f782c04ccf3243a0e636c96b87832c834"></a>

<a id="canonical-e7865fac4bcf5518df45b5471d261471dc2ce237dde10e2618967db7a3d20732"></a>

## location property — palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info / 361b34fea3e2 / 5

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

<a id="canonical-a44e7227699d65365176b6fb0a1bc18fa3990256244dbe9b981adb3925ba2ae4"></a>

<a id="canonical-f77e6f54cb3297f5d84927154a0beb2b4eedcd8b07c4f8c37eb7d298a63f45be"></a>

## store_provider property — palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info / 361b34fea3e2 / 6

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

<a id="canonical-1c012034a0d3ed54a149cf7e02adb86c7b6f8092fd0ad0cacea59c3690090767"></a>

## Next pages — palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info / 361b34fea3e2 / 7

- [palo_alto_fw_service.panorama_server.authorization_key](resources--nfv_service--reference--group-004.md#canonical-eb18479257d1b0b1b1334ce2a9daa05da3175e89c1e2daca86909038a18753c2)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-6d1b5a04798f2f144f8735d09925d645a613ff5a9be09ad7f5d59a14bf04a729"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1c1e9cad40e905d6a39b2ffe165bf5244ee4f68662cc89304b27b0eef60e219"></a>

## palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info — palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info / c92fca8b42d1 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-683b0e043ad34e16d220e00d2c0ad7d7794cd7a996943065814c81f0d29e03cf)
- [palo_alto_fw_service.panorama_server](resources--nfv_service--reference--group-004.md#canonical-45ec14f98c4298e49ab80b0eab6e4d912220244b1166f017f7ae9536733a9395)
- [palo_alto_fw_service.panorama_server.authorization_key](resources--nfv_service--reference--group-004.md#canonical-eb18479257d1b0b1b1334ce2a9daa05da3175e89c1e2daca86909038a18753c2)
- palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info

<a id="canonical-2e84232c05d795390c14466bf502302b73bf86afdd2a418c4b129e753c09a370"></a>

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

<a id="canonical-8e08e733b2605643c5b397d2fb22d01e21574386dc3eabaa8b002a66f0e5e5f2"></a>

## Direct properties — palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info / c92fca8b42d1 / 3

<a id="canonical-87bd3e37ad89547ef42f316e123ce66a8735288c3485f6143a57225845f18aef"></a>

<a id="canonical-0ea4aaf02d5321476f4573ea6bc16b5e2cecb974d648763f00a49303221df888"></a>

## provider_ref property — palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info / c92fca8b42d1 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-049ab609317fc0c222c0583699c67ee46f4999cd828db90fcb027f32a732f1d4"></a>

<a id="canonical-fd0ba0636be992f713602332a7b9e48a095305b69d38b1ae2fd4c1dc3b571308"></a>

## url property — palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info / c92fca8b42d1 / 5

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

<a id="canonical-edec3e50318ecdef2e70d87e41cb7e1aa71686df955fb3a4cde6ebc0bd14c2ea"></a>

## Next pages — palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info / c92fca8b42d1 / 6

- [palo_alto_fw_service.panorama_server.authorization_key](resources--nfv_service--reference--group-004.md#canonical-eb18479257d1b0b1b1334ce2a9daa05da3175e89c1e2daca86909038a18753c2)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-8d807ab1575a9ed710e3454886f81ddf66001c6a42213bf4fcf98c2b5a47020d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8c5b61f602a49cb831ad3d30e4b8c8bd088306084a28ecab23b22cd1faa5b4ad"></a>

## palo_alto_fw_service.service_nodes — palo_alto_fw_service.service_nodes / 23daba81614b / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-683b0e043ad34e16d220e00d2c0ad7d7794cd7a996943065814c81f0d29e03cf)
- palo_alto_fw_service.service_nodes

<a id="canonical-72e64a43fbb90dbf2f5723ec26c747305a7dc378e24b4a51b262380ce816370a"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for service nodes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("nodes")}
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
service_nodes {
  # Configure direct properties listed below.
}
```

<a id="canonical-2e221df2a55ca6b24c8174144113883df970aec44748cbfe26284c83c3850647"></a>

## Direct properties — palo_alto_fw_service.service_nodes / 23daba81614b / 3

- [nodes](resources--nfv_service--reference--group-004.md#canonical-26814021098bf67772517b5af60c3526dcd29d3f08dff4528749814985599096): complete subsection reference.

<a id="canonical-a80a63b20299760ec23a405225225f87e79cc6aa565dd30d9445e4a93b5ca974"></a>

## Next pages — palo_alto_fw_service.service_nodes / 23daba81614b / 4

- [palo_alto_fw_service.service_nodes.nodes](resources--nfv_service--reference--group-004.md#canonical-26814021098bf67772517b5af60c3526dcd29d3f08dff4528749814985599096)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-683b0e043ad34e16d220e00d2c0ad7d7794cd7a996943065814c81f0d29e03cf)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-26814021098bf67772517b5af60c3526dcd29d3f08dff4528749814985599096"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-03880f326034d24126124069c104ba3063dc94cf0ef6a62136d439cb15c9ef30"></a>

## palo_alto_fw_service.service_nodes.nodes — palo_alto_fw_service.service_nodes.nodes / 3da9731c8b21 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-683b0e043ad34e16d220e00d2c0ad7d7794cd7a996943065814c81f0d29e03cf)
- [palo_alto_fw_service.service_nodes](resources--nfv_service--reference--group-004.md#canonical-8d807ab1575a9ed710e3454886f81ddf66001c6a42213bf4fcf98c2b5a47020d)
- palo_alto_fw_service.service_nodes.nodes

<a id="canonical-a45394d945aae609dcab8272ea96f8a4dff274534e81808b0005d3af98c60f2f"></a>

Type: `"object"`. list nested block, Optional.

Palo Alto Networks AZ Nodes. Configuration parameter for nodes

Upstream description:

Configuration parameter for nodes

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("aws_az_name",
    "node_name"),
  validators.ConflictingListObjectAttributes("mgmt_subnet",
    "reserved_mgmt_subnet")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 2,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 2,
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
    "ves.io.schema.rules.repeated.max_items": "2",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "2",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
nodes {
  # Configure direct properties listed below.
}
```

<a id="canonical-7f201a00e338e203592c4c0824cdd94dbb0e76e2ef4ae33a648ac96c5c690f4c"></a>

## Direct properties — palo_alto_fw_service.service_nodes.nodes / 3da9731c8b21 / 3

<a id="canonical-34f512f9488f5da8b2fefd7b60d32e53cf54084e706d5a8b594ad9788b3ae5c9"></a>

<a id="canonical-042e5f522cc16ad66704e4828ea036e9de4271e96a08e682e691537512edbd1e"></a>

## aws_az_name property — palo_alto_fw_service.service_nodes.nodes / 3da9731c8b21 / 4

Type: `"string"`. Optional.

AWS availability zone, must be consistent with the selected AWS region. It is recommended that AZ is
one of the AZ for sites.

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
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^([a-z]{2})-([a-z0-9]{4,20})-([a-z0-9]{2})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([a-z]{2})-([a-z0-9]{4,20})-([a-z0-9]{2})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([a-z]{2})-([a-z0-9]{4,20})-([a-z0-9]{2})$"
  }
}
```

- [mgmt_subnet](resources--nfv_service--reference--group-004.md#canonical-2acc2836dcc99bddbe41c6701f2ccf1e26a85e7bcf9522d5c5082a43b4e6aed8): complete subsection reference.

<a id="canonical-b92f7dd2fb935770a643cd2e0557113696a919e75a70c826b55052aa9d704ee7"></a>

<a id="canonical-be93fc9025c243759391bbd30630554de6d8d11cde0acbc7d858d6d27082d98e"></a>

## node_name property — palo_alto_fw_service.service_nodes.nodes / 3da9731c8b21 / 5

Type: `"string"`. Optional.

Node Name will be used to assign as hostname to the service.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [reserved_mgmt_subnet](resources--nfv_service--reference--group-004.md#canonical-842013ddc39ce8eb39a67cd0af654a1911afb8f3ed01ed58d9c01b2ab33b1a49): complete subsection reference.

<a id="canonical-3c923e65fc76e4d5124dd330f9b6020e0006bcf4cd2559a174372184e981ddee"></a>

## Next pages — palo_alto_fw_service.service_nodes.nodes / 3da9731c8b21 / 6

- [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet](resources--nfv_service--reference--group-004.md#canonical-2acc2836dcc99bddbe41c6701f2ccf1e26a85e7bcf9522d5c5082a43b4e6aed8)
- [palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet](resources--nfv_service--reference--group-004.md#canonical-842013ddc39ce8eb39a67cd0af654a1911afb8f3ed01ed58d9c01b2ab33b1a49)
- [palo_alto_fw_service.service_nodes](resources--nfv_service--reference--group-004.md#canonical-8d807ab1575a9ed710e3454886f81ddf66001c6a42213bf4fcf98c2b5a47020d)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-2acc2836dcc99bddbe41c6701f2ccf1e26a85e7bcf9522d5c5082a43b4e6aed8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2017460c2f5b9de0a77b1b2130078997c515a3133c06a979e38f66f7df5f01d8"></a>

## palo_alto_fw_service.service_nodes.nodes.mgmt_subnet — palo_alto_fw_service.service_nodes.nodes.mgmt_subnet / dbaf18f5949d / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-683b0e043ad34e16d220e00d2c0ad7d7794cd7a996943065814c81f0d29e03cf)
- [palo_alto_fw_service.service_nodes](resources--nfv_service--reference--group-004.md#canonical-8d807ab1575a9ed710e3454886f81ddf66001c6a42213bf4fcf98c2b5a47020d)
- [palo_alto_fw_service.service_nodes.nodes](resources--nfv_service--reference--group-004.md#canonical-26814021098bf67772517b5af60c3526dcd29d3f08dff4528749814985599096)
- palo_alto_fw_service.service_nodes.nodes.mgmt_subnet

<a id="canonical-f0ce6950ebabbe86a9f32a48446eac00937acbc1b3a78418ebc136b88e689f70"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for mgmt subnet.

Upstream description:

Parameters for AWS subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_subnet_id",
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
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

Terraform syntax:

```terraform
mgmt_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-d250f268f7c10b6abdc7b708dcf9f0622417efb529fb7aab614404dd8386afc6"></a>

## Direct properties — palo_alto_fw_service.service_nodes.nodes.mgmt_subnet / dbaf18f5949d / 3

<a id="canonical-3cc89bcaa3bd0c69c9a536c4202b5a70590691e32d71526a4823628f378f2bdd"></a>

<a id="canonical-2de4c8c95a2abe18e18b7673fe26034cf632b71f475ecf795c0dcb1bfc74fa00"></a>

## existing_subnet_id property — palo_alto_fw_service.service_nodes.nodes.mgmt_subnet / dbaf18f5949d / 4

Type: `"string"`. Optional.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Upstream description:

Exclusive with \[subnet\_param\] Information about existing subnet ID.

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
    },
    "pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [subnet_param](resources--nfv_service--reference--group-004.md#canonical-c51c091fcce83a439d441b872671e9926bff142c522e436d1d6b2aa1fa1f3b0f): complete subsection reference.

<a id="canonical-50174783f0b9fd26de2f570e7b4a6f4d45b08aa56b8d183409adbc886e32e051"></a>

## Next pages — palo_alto_fw_service.service_nodes.nodes.mgmt_subnet / dbaf18f5949d / 5

- [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param](resources--nfv_service--reference--group-004.md#canonical-c51c091fcce83a439d441b872671e9926bff142c522e436d1d6b2aa1fa1f3b0f)
- [palo_alto_fw_service.service_nodes.nodes](resources--nfv_service--reference--group-004.md#canonical-26814021098bf67772517b5af60c3526dcd29d3f08dff4528749814985599096)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-c51c091fcce83a439d441b872671e9926bff142c522e436d1d6b2aa1fa1f3b0f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-15824649090d935bf91becf84d5e9afcfdf7e132d1b04149ec76b777cce645b1"></a>

## palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param — palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param / b060d22357a4 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-683b0e043ad34e16d220e00d2c0ad7d7794cd7a996943065814c81f0d29e03cf)
- [palo_alto_fw_service.service_nodes](resources--nfv_service--reference--group-004.md#canonical-8d807ab1575a9ed710e3454886f81ddf66001c6a42213bf4fcf98c2b5a47020d)
- [palo_alto_fw_service.service_nodes.nodes](resources--nfv_service--reference--group-004.md#canonical-26814021098bf67772517b5af60c3526dcd29d3f08dff4528749814985599096)
- [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet](resources--nfv_service--reference--group-004.md#canonical-2acc2836dcc99bddbe41c6701f2ccf1e26a85e7bcf9522d5c5082a43b4e6aed8)
- palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param

<a id="canonical-6dd099fe25611dc2860495a878613d10cc0efd48dadf9a6a1792bdc4a868f82f"></a>

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

<a id="canonical-88b09752aaaf2bb16366ef736c3a4be7918e254a36cf5966015a05ded6dcd75b"></a>

## Direct properties — palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param / b060d22357a4 / 3

<a id="canonical-adc624dfb24fbda6ac0e4fe0f40607985a4917e46363f578275421488ffc3d7a"></a>

<a id="canonical-529b292902aea0833c5667fb58cdd334df102b2e376af3af05642376c2b23a69"></a>

## ipv4 property — palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param / b060d22357a4 / 4

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

<a id="canonical-d85d31c93965a83407a4866d42fc51357a318cfc284f34939cbb333f1e913ba6"></a>

## Next pages — palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param / b060d22357a4 / 5

- [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet](resources--nfv_service--reference--group-004.md#canonical-2acc2836dcc99bddbe41c6701f2ccf1e26a85e7bcf9522d5c5082a43b4e6aed8)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-842013ddc39ce8eb39a67cd0af654a1911afb8f3ed01ed58d9c01b2ab33b1a49"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-12283a10408d772694ca72471b74053ee78f48ba06cc705e6c377b8e67dab45b"></a>

## palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet — palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet / 217929d66fb4 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-683b0e043ad34e16d220e00d2c0ad7d7794cd7a996943065814c81f0d29e03cf)
- [palo_alto_fw_service.service_nodes](resources--nfv_service--reference--group-004.md#canonical-8d807ab1575a9ed710e3454886f81ddf66001c6a42213bf4fcf98c2b5a47020d)
- [palo_alto_fw_service.service_nodes.nodes](resources--nfv_service--reference--group-004.md#canonical-26814021098bf67772517b5af60c3526dcd29d3f08dff4528749814985599096)
- palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet

<a id="canonical-e70eda96bf31c1ad2f7bf668627cc5ec520812d29867a54f8edd34d9f48ce5de"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for reserved mgmt subnet.

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
reserved_mgmt_subnet = {}
```

<a id="canonical-046ec00c164e8e3e342e48c8ff28103c90d5a10d4f35e70baee002f7861f8b97"></a>

## Direct properties — palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet / 217929d66fb4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fd2867cc36d6d13c8b68eda4da2b26893d2f7d3d02be1a43aee36682c6f89912"></a>

## Next pages — palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet / 217929d66fb4 / 4

- [palo_alto_fw_service.service_nodes.nodes](resources--nfv_service--reference--group-004.md#canonical-26814021098bf67772517b5af60c3526dcd29d3f08dff4528749814985599096)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-22b89a5a007d5d26888d6c80645783ec0135cd1ced1c3d4d29ec4c3cdc0ebce1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a662333120a444798ae72bdaa7f1a35d1444d445be2d6c1331232c04c1139466"></a>

## timeouts — timeouts / 2c672f092a40 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- timeouts

<a id="canonical-ae9e0c0d2d16c9868cadf460e69bb17c9e1e22f0e0c3c5c6ffe73466fe9ef49a"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-a2f365664ef61a70a229f6150261ea683eff8ed4b02ebbec0e6394221efc7ce2"></a>

## Direct properties — timeouts / 2c672f092a40 / 3

<a id="canonical-eaba720097bc356f1e0eff85431f614c806744ca7d3f7f4dbc238ac990ea9ac6"></a>

<a id="canonical-b3eea08909177489b125bf04d7f2341411403cefcef6037c41f8dee35f042b0e"></a>

## create property — timeouts / 2c672f092a40 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-801d2ae594eda82dfe6af70557aaa2c01812c8afbdf8f2b7acccb839dfde8539"></a>

<a id="canonical-1b9ab56a69013918ffe35162341416f5a07858e2c49556f1809467e28719aca3"></a>

## delete property — timeouts / 2c672f092a40 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-734f0daabe3ed3e505ef32ffec0b5cb3c153389f615edc930a6c8c619a2aebbb"></a>

<a id="canonical-8d709543a24f4f5c14e732a49c7cc10a8c28fc7bc731861b2e66f9ee899c9861"></a>

## read property — timeouts / 2c672f092a40 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-aeefb5859002af36a87292f211b5e4ffcad6ef54ac594361cb627c95cbb9a7d0"></a>

<a id="canonical-017113936ee557fd0b1a5cf41137a6de495dbdedc4a7ce500e6ad9042f67fa12"></a>

## update property — timeouts / 2c672f092a40 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-5fd3d3c6e05b244b111cf325ad17a784f7cf3e6899b775cb68a5e62c444b304a"></a>

## Next pages — timeouts / 2c672f092a40 / 8

- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
