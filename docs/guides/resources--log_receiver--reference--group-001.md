---
page_title: "xcsh_log_receiver reference"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_log_receiver reference."
---

# xcsh_log_receiver reference

<a id="canonical-ef02783c49c50c98ba0aaaf5e5b83465c290bcf3f2f090952d79c6692f4536e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223f8eff5b6a3438404b418e74375ae0c26778d4251dc68e001f1fc63d0c4dd"></a>

## Property reference — Property reference / 2860a3b8db11 / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)
- Property reference

<a id="canonical-addbb16a5c1a45d54ab30e7a6cdd6bac96de7dc4adea4282e7eaf689a1ab8c61"></a>

## Direct properties — Property reference / 2860a3b8db11 / 3

<a id="canonical-a7854974e6e91b91b772f86d69a7b0107c6f9349e3f35ef24a276f339f090b02"></a>

<a id="canonical-192bbb568358d78fd63a62f597ee195ee7a5ba9681616770b3558adfdaef7865"></a>

## annotations property — Property reference / 2860a3b8db11 / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

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

<a id="canonical-89dba7208acf0d38c8e6a872e5ba95f2e366ecc890539958453dd64a5ebc2a35"></a>

<a id="canonical-295e35107a56d226622f8cbca9c863295bd47b71f1d5481bc1f01aeeb547e27c"></a>

## description property — Property reference / 2860a3b8db11 / 5

Type: `"string"`. Optional.

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

<a id="canonical-820905da4722252944b3e9ddc3d4ffc03ce4e885cb470fa983cdac69278c056a"></a>

<a id="canonical-179d0e5dc921044f185d5ae9c1358970d5dc78fd06425be0b58f4a499cc54935"></a>

## disable property — Property reference / 2860a3b8db11 / 6

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1a83b9f883cd742928eed494b1885e1fedf18d9dc61763f6ba1cca59f3d47e22"></a>

<a id="canonical-c790e204102b376775b3f6d436eab39b7a0c017691ee9051746c9c79c552851a"></a>

## id property — Property reference / 2860a3b8db11 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-dedca48af0b999cd3657628949cda9ae1d19839136f41be383d8be9b2e0c268b"></a>

<a id="canonical-a7d8ce438cb5fba7979c48d236179124d9c7e1c417c8e2a9e5add35ad2233bde"></a>

## labels property — Property reference / 2860a3b8db11 / 8

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

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

<a id="canonical-e3cedaa86fa410ba4751ff7304588542608c66a92bfec4d397e185860609b313"></a>

<a id="canonical-48520f67bec8c4d2162ac32d3cde78083dc1b7c34834a01e937b16bf888dca2b"></a>

## name property — Property reference / 2860a3b8db11 / 9

Type: `"string"`. Required.

Name of the Log Receiver. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
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

<a id="canonical-d04b7dfea3edb4eefea89e7fba09e712372c2c88e37ccf1ca1cd0b26bf1666c6"></a>

<a id="canonical-bd2e5426d9bce4363f7f58411b61f8056d290a27b00b8722ddb0ecaedde89c71"></a>

## namespace property — Property reference / 2860a3b8db11 / 10

Type: `"string"`. Required.

Namespace where the Log Receiver is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
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

- [site_local](resources--log_receiver--reference--group-001.md#canonical-9bfdd8780bc26e228d2dc526f5e469b82286b25af21d5dd772989264232a6e4d): complete subsection reference.

- [syslog](resources--log_receiver--reference--group-001.md#canonical-238a4fb5c929043bd6111b3f52ebb4d83281b91492c688ac149bbfc23554692d): complete subsection reference.

- [timeouts](resources--log_receiver--reference--group-001.md#canonical-9c77ae414286862fafbb9218e1483a458a6c06d4bf7bf5ec8aaca80a4ee18cb9): complete subsection reference.

<a id="canonical-20e3c0f379107444dc85276e909616a62527d75d456da569f1eb176ada070d43"></a>

## All schema paths — Property reference / 2860a3b8db11 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--log_receiver--reference--group-001.md#canonical-a7854974e6e91b91b772f86d69a7b0107c6f9349e3f35ef24a276f339f090b02) |
| `description` | [description](resources--log_receiver--reference--group-001.md#canonical-89dba7208acf0d38c8e6a872e5ba95f2e366ecc890539958453dd64a5ebc2a35) |
| `disable` | [disable](resources--log_receiver--reference--group-001.md#canonical-820905da4722252944b3e9ddc3d4ffc03ce4e885cb470fa983cdac69278c056a) |
| `id` | [id](resources--log_receiver--reference--group-001.md#canonical-1a83b9f883cd742928eed494b1885e1fedf18d9dc61763f6ba1cca59f3d47e22) |
| `labels` | [labels](resources--log_receiver--reference--group-001.md#canonical-dedca48af0b999cd3657628949cda9ae1d19839136f41be383d8be9b2e0c268b) |
| `name` | [name](resources--log_receiver--reference--group-001.md#canonical-e3cedaa86fa410ba4751ff7304588542608c66a92bfec4d397e185860609b313) |
| `namespace` | [namespace](resources--log_receiver--reference--group-001.md#canonical-d04b7dfea3edb4eefea89e7fba09e712372c2c88e37ccf1ca1cd0b26bf1666c6) |
| `site_local` | [site_local](resources--log_receiver--reference--group-001.md#canonical-e3afd9e56fbc82cf400955cbf5e9649f808bf1fd035368f16416b954911e8e2b) |
| `syslog` | [syslog](resources--log_receiver--reference--group-001.md#canonical-db5e9ea7503adddea3048dca908d006175559295ff23adade808f4a6cb1fb7e1) |
| `syslog.syslog_rfc5424` | [syslog.syslog_rfc5424](resources--log_receiver--reference--group-001.md#canonical-5d19a9898ef241aab6654f9f369217cd1dea7dae6d8562354e9b4fac229baf13) |
| `syslog.tcp_server` | [syslog.tcp_server](resources--log_receiver--reference--group-001.md#canonical-65be102e91aa7944f2ad6a69d18f15b09532b319a77b8ba410a14d3271267f47) |
| `syslog.tcp_server.port` | [syslog.tcp_server.port](resources--log_receiver--reference--group-001.md#canonical-d4a266e4291fc04ed01d2aab76775c9882c86fca0f59144fce0315f5147d10f3) |
| `syslog.tcp_server.server_name` | [syslog.tcp_server.server_name](resources--log_receiver--reference--group-001.md#canonical-07b2b55c310407d050961706c1fda8fbcae6d60b175ac7ba24cb96608ab57088) |
| `syslog.tls_server` | [syslog.tls_server](resources--log_receiver--reference--group-001.md#canonical-082f2063afc7494afb601ccb275061342f3791c2018cb8eea6bda9ed017352f8) |
| `syslog.tls_server.default_https_port` | [syslog.tls_server.default_https_port](resources--log_receiver--reference--group-001.md#canonical-46d87970e939247743adedabdb7a860900149999b5806bedfe912b0bf767847b) |
| `syslog.tls_server.default_syslog_tls_port` | [syslog.tls_server.default_syslog_tls_port](resources--log_receiver--reference--group-001.md#canonical-eadc4cc0ce99e78827fc5cc58dbf31c500ade1585ae1691e127306ac9a139c37) |
| `syslog.tls_server.mtls_disabled` | [syslog.tls_server.mtls_disabled](resources--log_receiver--reference--group-001.md#canonical-fd71c0b248c6c2d1eec25279b9dc948b5c40490f16089b922583ffcbeb7e4b3d) |
| `syslog.tls_server.mtls_enable` | [syslog.tls_server.mtls_enable](resources--log_receiver--reference--group-001.md#canonical-3e2f31f1ca1538e82161cdc827168bab8671eb560ffce14c48d6ea86eeac94d3) |
| `syslog.tls_server.mtls_enable.certificate` | [syslog.tls_server.mtls_enable.certificate](resources--log_receiver--reference--group-001.md#canonical-d445c177864bdae92658faf55e6e2382b988dee5f7647c8a7eb16b131d96f5e2) |
| `syslog.tls_server.mtls_enable.key_url` | [syslog.tls_server.mtls_enable.key_url](resources--log_receiver--reference--group-001.md#canonical-896e5015d97e3b9337e616c2b43b78877f3e986b97bdd89e219a44d2b12fe143) |
| `syslog.tls_server.mtls_enable.key_url.blindfold_secret_info` | [syslog.tls_server.mtls_enable.key_url.blindfold_secret_info](resources--log_receiver--reference--group-001.md#canonical-cc68d6314df0795772f1265497f9ed5c84ac1bdad6380b2997cf03f4270fb895) |
| `syslog.tls_server.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [syslog.tls_server.mtls_enable.key_url.blindfold_secret_info.decryption_provider](resources--log_receiver--reference--group-001.md#canonical-53508740443ad1be1db735325cd2d524beb37884ca20b57d8d076924d099f901) |
| `syslog.tls_server.mtls_enable.key_url.blindfold_secret_info.location` | [syslog.tls_server.mtls_enable.key_url.blindfold_secret_info.location](resources--log_receiver--reference--group-001.md#canonical-9b6130d37b921ab933d9ba2f5b91ae06d40ccd6127680eb1723138b59ded6809) |
| `syslog.tls_server.mtls_enable.key_url.blindfold_secret_info.store_provider` | [syslog.tls_server.mtls_enable.key_url.blindfold_secret_info.store_provider](resources--log_receiver--reference--group-001.md#canonical-d89491e1f3943a07d4e21a712728196b70fb63a2e2a53d7e883ad4f544fb6f30) |
| `syslog.tls_server.mtls_enable.key_url.clear_secret_info` | [syslog.tls_server.mtls_enable.key_url.clear_secret_info](resources--log_receiver--reference--group-001.md#canonical-134f5d15abf06ef95f50688e83bf41b1811d439a7b764d32e13a2cbc0b207567) |
| `syslog.tls_server.mtls_enable.key_url.clear_secret_info.provider_ref` | [syslog.tls_server.mtls_enable.key_url.clear_secret_info.provider_ref](resources--log_receiver--reference--group-001.md#canonical-a131c324b0b59e9976656575d152dde47761b9fde7fa53aef2ea0c0c86c54f99) |
| `syslog.tls_server.mtls_enable.key_url.clear_secret_info.url` | [syslog.tls_server.mtls_enable.key_url.clear_secret_info.url](resources--log_receiver--reference--group-001.md#canonical-b8c936579dcbad63e2663818236d8cd3ecb3dcf6c5dd20d8060473c7f2ec37fc) |
| `syslog.tls_server.port` | [syslog.tls_server.port](resources--log_receiver--reference--group-001.md#canonical-69ce0298762f495a23aa41d050db5717d86d366976da1be8c9f81c316661aeac) |
| `syslog.tls_server.server_name` | [syslog.tls_server.server_name](resources--log_receiver--reference--group-001.md#canonical-875e6e04560840e4d5cbba51df59878959c4e151c693760f211ded63c30988f1) |
| `syslog.tls_server.trusted_ca_url` | [syslog.tls_server.trusted_ca_url](resources--log_receiver--reference--group-001.md#canonical-131f25b5211ac181009ef5c9e71e5a11c7e45ef482de74ecac724db600b82793) |
| `syslog.tls_server.volterra_ca` | [syslog.tls_server.volterra_ca](resources--log_receiver--reference--group-001.md#canonical-3c8a84303f93f7f43a191405328472c6d77a3468fdae55a0dff504750fc94a30) |
| `syslog.udp_server` | [syslog.udp_server](resources--log_receiver--reference--group-001.md#canonical-eab9d957b67babd671ef574024676442b9f5d6e61679e7d36e7a74ce4f493737) |
| `syslog.udp_server.port` | [syslog.udp_server.port](resources--log_receiver--reference--group-001.md#canonical-2f3658cd8e9c440a14438b363b75c2a9970fd4e5bd680f94e2afa6f28120be32) |
| `syslog.udp_server.server_name` | [syslog.udp_server.server_name](resources--log_receiver--reference--group-001.md#canonical-5c65539356b12f6069d38c47a905c2c1a44676b4dc53ab8dc61e696d71238dd9) |
| `timeouts` | [timeouts](resources--log_receiver--reference--group-001.md#canonical-82593a3b8255332a63e32f5f39afa82adfc2c4418387f48a92595cc0cffbdee3) |
| `timeouts.create` | [timeouts.create](resources--log_receiver--reference--group-001.md#canonical-4cd0a9e655e22cfd7e13bc1213ea13e307f2faad193edc4a2fc92e97ecb2b6b2) |
| `timeouts.delete` | [timeouts.delete](resources--log_receiver--reference--group-001.md#canonical-8cec0affb97d41c592d02d95a0b3f049a4a232b229be74df6a0033fbb115171b) |
| `timeouts.read` | [timeouts.read](resources--log_receiver--reference--group-001.md#canonical-9b466b7c8a0a2c26b86febad03bc1dacb40d812c56af7a91f8e218e3c3ecf902) |
| `timeouts.update` | [timeouts.update](resources--log_receiver--reference--group-001.md#canonical-3f76d9edbc1c568a3905009a82e3183dee1edf62bac76b3a17781b881256a423) |

<a id="canonical-57c9221f49c7711c11cbcb4747a6d3b6d261b14e63ab075ac8740886d964d4a2"></a>

## Next pages — Property reference / 2860a3b8db11 / 12

- [site_local](resources--log_receiver--reference--group-001.md#canonical-9bfdd8780bc26e228d2dc526f5e469b82286b25af21d5dd772989264232a6e4d)
- [syslog](resources--log_receiver--reference--group-001.md#canonical-238a4fb5c929043bd6111b3f52ebb4d83281b91492c688ac149bbfc23554692d)
- [timeouts](resources--log_receiver--reference--group-001.md#canonical-9c77ae414286862fafbb9218e1483a458a6c06d4bf7bf5ec8aaca80a4ee18cb9)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)

<a id="canonical-9bfdd8780bc26e228d2dc526f5e469b82286b25af21d5dd772989264232a6e4d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-180e0222735a5e1bf779b9993e6fe9d63f4c5d04dd3def1e58ffad78eb25e153"></a>

## site_local — site_local / ddfab18c0a29 / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)
- [Property reference](resources--log_receiver--reference--group-001.md#canonical-ef02783c49c50c98ba0aaaf5e5b83465c290bcf3f2f090952d79c6692f4536e1)
- site_local

<a id="canonical-e3afd9e56fbc82cf400955cbf5e9649f808bf1fd035368f16416b954911e8e2b"></a>

Type: `"object"`. single nested block, Optional.

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
site_local {}
```

<a id="canonical-0d0823e4fa96434a1c8d099d5315c52681e02c93d086b857f3004c1a4c0e8e6f"></a>

## Direct properties — site_local / ddfab18c0a29 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cb3b9b1def3ea380d778f6a563acb640c2f08f8afe094de55123637fd40903c7"></a>

## Next pages — site_local / ddfab18c0a29 / 4

- [Property reference](resources--log_receiver--reference--group-001.md#canonical-ef02783c49c50c98ba0aaaf5e5b83465c290bcf3f2f090952d79c6692f4536e1)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)

<a id="canonical-238a4fb5c929043bd6111b3f52ebb4d83281b91492c688ac149bbfc23554692d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a43e6b6582e20b0d54970d64a596e1ee9d363fa7e359f3c763ec88f2d80c021"></a>

## syslog — syslog / e536f3468ed6 / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)
- [Property reference](resources--log_receiver--reference--group-001.md#canonical-ef02783c49c50c98ba0aaaf5e5b83465c290bcf3f2f090952d79c6692f4536e1)
- syslog

<a id="canonical-db5e9ea7503adddea3048dca908d006175559295ff23adade808f4a6cb1fb7e1"></a>

Type: `"object"`. single nested block, Optional.

Syslog Server Configuration. Configuration for syslog server.

Upstream description:

Configuration for syslog server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("tcp_server",
    "tls_server"),
  validators.ConflictingObjectAttributes("tcp_server",
    "udp_server"),
  validators.ConflictingObjectAttributes("tls_server",
    "udp_server")}
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
  "x-ves-oneof-field-format_choice": "[\"syslog_rfc5424\"]",
  "x-ves-oneof-field-mode_choice": "[\"tcp_server\",\"tls_server\",\"udp_server\"]"
}
```

Terraform syntax:

```terraform
syslog {
  # Configure direct properties listed below.
}
```

<a id="canonical-f74020771e5746d061ac551ef10ab94dcf0f8c7aa07d38e69a856e0cfaf8340b"></a>

## Direct properties — syslog / e536f3468ed6 / 3

<a id="canonical-5d19a9898ef241aab6654f9f369217cd1dea7dae6d8562354e9b4fac229baf13"></a>

<a id="canonical-f449a7316240df1667a5fd5cdf00fa04b2fba37b1c23a4c0616b90b37009b500"></a>

## syslog_rfc5424 property — syslog / e536f3468ed6 / 4

Type: `"number"`. Optional.

Exclusive with \[\] Select RFC5424 syslog format and maximum message length.

Upstream description:

Exclusive with \[\] Select RFC5424 syslog format and maximum message length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(408, 268435456),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 268435456,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 408
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "408",
    "ves.io.schema.rules.uint32.lte": "268435456"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "408",
    "ves.io.schema.rules.uint32.lte": "268435456"
  }
}
```

- [tcp_server](resources--log_receiver--reference--group-001.md#canonical-11c698ef4cfce6841a17ddfa62d4f5dd677611998dcd878e735ad0c3200f68b6): complete subsection reference.

- [tls_server](resources--log_receiver--reference--group-001.md#canonical-e04115abcd19df912fb601a0f0928efbc4a740cb9dfe730db2e3b5eaa7b81f1e): complete subsection reference.

- [udp_server](resources--log_receiver--reference--group-001.md#canonical-0faaaea294d489614a46833d43b29836faa8c574d45e6c71e7f52dadc8160c78): complete subsection reference.

<a id="canonical-8c510e5b42f114f536d066272d28694ef4cbf3d5726ada3d1466b7b16561df94"></a>

## Next pages — syslog / e536f3468ed6 / 5

- [syslog.tcp_server](resources--log_receiver--reference--group-001.md#canonical-11c698ef4cfce6841a17ddfa62d4f5dd677611998dcd878e735ad0c3200f68b6)
- [syslog.tls_server](resources--log_receiver--reference--group-001.md#canonical-e04115abcd19df912fb601a0f0928efbc4a740cb9dfe730db2e3b5eaa7b81f1e)
- [syslog.udp_server](resources--log_receiver--reference--group-001.md#canonical-0faaaea294d489614a46833d43b29836faa8c574d45e6c71e7f52dadc8160c78)
- [Property reference](resources--log_receiver--reference--group-001.md#canonical-ef02783c49c50c98ba0aaaf5e5b83465c290bcf3f2f090952d79c6692f4536e1)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)

<a id="canonical-11c698ef4cfce6841a17ddfa62d4f5dd677611998dcd878e735ad0c3200f68b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1848db5c810e8b103b9f503cc7dfc1acd26783070c1a8698c6b2146ab9bbecdc"></a>

## syslog.tcp_server — syslog.tcp_server / 0c91cdb1742e / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)
- [Property reference](resources--log_receiver--reference--group-001.md#canonical-ef02783c49c50c98ba0aaaf5e5b83465c290bcf3f2f090952d79c6692f4536e1)
- [syslog](resources--log_receiver--reference--group-001.md#canonical-238a4fb5c929043bd6111b3f52ebb4d83281b91492c688ac149bbfc23554692d)
- syslog.tcp_server

<a id="canonical-65be102e91aa7944f2ad6a69d18f15b09532b319a77b8ba410a14d3271267f47"></a>

Type: `"object"`. single nested block, Optional.

TCP Server name and Port Number. Name and port number for a TCP server.

Upstream description:

Name and port number for a TCP server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("port",
    "server_name")}
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
tcp_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-28e5142728479cd510da7836c116167a80fe18f84f48ba0175316be3731df30b"></a>

## Direct properties — syslog.tcp_server / 0c91cdb1742e / 3

<a id="canonical-d4a266e4291fc04ed01d2aab76775c9882c86fca0f59144fce0315f5147d10f3"></a>

<a id="canonical-94b59adb4df7ee6f679c90dae42966c473f23b4017574bf2d9d693229d7a54fc"></a>

## port property — syslog.tcp_server / 0c91cdb1742e / 4

Type: `"number"`. Optional.

Port Number. Port number used for communication.

Upstream description:

Port number used for communication.

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
    "minimum": 1,
    "multipleOf": 1
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-07b2b55c310407d050961706c1fda8fbcae6d60b175ac7ba24cb96608ab57088"></a>

<a id="canonical-e450b17cd62798b5338bddb3fe58a073e0bddeb8d847b14cc83514fae62784c5"></a>

## server_name property — syslog.tcp_server / 0c91cdb1742e / 5

Type: `"string"`. Optional.

Server name is fully qualified domain name or IP address of the server.

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
    "ves.io.schema.rules.string.hostname_or_ip": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname_or_ip": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-29440b8400bed155f631cd581261fdc0cd6a5fa8e77c957a3a0e444f2b3e04ef"></a>

## Next pages — syslog.tcp_server / 0c91cdb1742e / 6

- [syslog](resources--log_receiver--reference--group-001.md#canonical-238a4fb5c929043bd6111b3f52ebb4d83281b91492c688ac149bbfc23554692d)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)

<a id="canonical-e04115abcd19df912fb601a0f0928efbc4a740cb9dfe730db2e3b5eaa7b81f1e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a5401f5f560b1c84fbcb06f121e05051da70b8ff481691e86805f3e22a691d83"></a>

## syslog.tls_server — syslog.tls_server / 0c776ac16175 / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)
- [Property reference](resources--log_receiver--reference--group-001.md#canonical-ef02783c49c50c98ba0aaaf5e5b83465c290bcf3f2f090952d79c6692f4536e1)
- [syslog](resources--log_receiver--reference--group-001.md#canonical-238a4fb5c929043bd6111b3f52ebb4d83281b91492c688ac149bbfc23554692d)
- syslog.tls_server

<a id="canonical-082f2063afc7494afb601ccb275061342f3791c2018cb8eea6bda9ed017352f8"></a>

Type: `"object"`. single nested block, Optional.

TLS config for client of discovery service.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("server_name"),
  validators.ConflictingObjectAttributes("default_https_port",
    "default_syslog_tls_port"),
  validators.ConflictingObjectAttributes("default_https_port",
    "port"),
  validators.ConflictingObjectAttributes("default_syslog_tls_port",
    "port"),
  validators.ConflictingObjectAttributes("mtls_disabled",
    "mtls_enable"),
  validators.ConflictingObjectAttributes("trusted_ca_url",
    "volterra_ca")}
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
  "x-ves-oneof-field-ca_choice": "[\"trusted_ca_url\",\"volterra_ca\"]",
  "x-ves-oneof-field-mtls_choice": "[\"mtls_disabled\",\"mtls_enable\"]",
  "x-ves-oneof-field-port_choice": "[\"default_https_port\",\"default_syslog_tls_port\",\"port\"]"
}
```

Terraform syntax:

```terraform
tls_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-e3c3e7bb997ff34b50f1865cfa66705fa003e39b9a3296d46450138d4b09f25e"></a>

## Direct properties — syslog.tls_server / 0c776ac16175 / 3

- [default_https_port](resources--log_receiver--reference--group-001.md#canonical-2ebacabed412ae5ce841fd4f8398e00d9785b0ed54c2a2b0d61ca005025e087b): complete subsection reference.

- [default_syslog_tls_port](resources--log_receiver--reference--group-001.md#canonical-979cc51d6cbf9db9dea81fac45534d6a6b617b56f76a8f35c772ece5e5cacacb): complete subsection reference.

- [mtls_disabled](resources--log_receiver--reference--group-001.md#canonical-9de1e8c0460c53c63d8185ddb70cc148b23336e5c88ae28e55ffa90e97fff9f4): complete subsection reference.

- [mtls_enable](resources--log_receiver--reference--group-001.md#canonical-e327963cd48210663b691b7632b7ae92705aeb4a777db0401b883d3ef6cfa496): complete subsection reference.

<a id="canonical-69ce0298762f495a23aa41d050db5717d86d366976da1be8c9f81c316661aeac"></a>

<a id="canonical-f1e10bd8618d8e19a0d8e14058fa32754c2e9e5b14b4312c0e89ebbe444e2c17"></a>

## port property — syslog.tls_server / 0c776ac16175 / 4

Type: `"number"`. Optional.

Exclusive with \[default\_https\_port default\_syslog\_tls\_port\] Custom port number used for
communication.

Upstream description:

Exclusive with \[default\_https\_port default\_syslog\_tls\_port\] Custom port number used for
communication.

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-875e6e04560840e4d5cbba51df59878959c4e151c693760f211ded63c30988f1"></a>

<a id="canonical-e6e2be3416fdb5d7cb6ace61e2ce57fdc67fd0f82bdda2711ff5f1678b02d15c"></a>

## server_name property — syslog.tls_server / 0c776ac16175 / 5

Type: `"string"`. Optional.

ServerName is passed to the server for SNI and is used in the client to check server certificates
against.

Upstream description:

ServerName is passed to the server for SNI and is used in the client to check server certificates
against.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-131f25b5211ac181009ef5c9e71e5a11c7e45ef482de74ecac724db600b82793"></a>

<a id="canonical-b7dda2adf082d97ac9872308b68a678d111515566fcaa42cf8d93f0f0877833f"></a>

## trusted_ca_url property — syslog.tls_server / 0c776ac16175 / 6

Type: `"string"`. Optional.

Exclusive with \[volterra\_ca\] The URL or value for trusted Server CA certificate or certificate
chain Certificates in PEM format including the PEM headers.

Upstream description:

Exclusive with \[volterra\_ca\] The URL or value for trusted Server CA certificate or certificate
chain Certificates in PEM format including the PEM headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(131072),
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
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [volterra_ca](resources--log_receiver--reference--group-001.md#canonical-9d22ae2a30e391644727ba133bac83bbe67a8d6c3357ba44814ad521b9c8cdd8): complete subsection reference.

<a id="canonical-9694bcadbb33f973a146c6295d28762c378e0bf7521013b46e05f753342665a2"></a>

## Next pages — syslog.tls_server / 0c776ac16175 / 7

- [syslog.tls_server.default_https_port](resources--log_receiver--reference--group-001.md#canonical-2ebacabed412ae5ce841fd4f8398e00d9785b0ed54c2a2b0d61ca005025e087b)
- [syslog.tls_server.default_syslog_tls_port](resources--log_receiver--reference--group-001.md#canonical-979cc51d6cbf9db9dea81fac45534d6a6b617b56f76a8f35c772ece5e5cacacb)
- [syslog.tls_server.mtls_disabled](resources--log_receiver--reference--group-001.md#canonical-9de1e8c0460c53c63d8185ddb70cc148b23336e5c88ae28e55ffa90e97fff9f4)
- [syslog.tls_server.mtls_enable](resources--log_receiver--reference--group-001.md#canonical-e327963cd48210663b691b7632b7ae92705aeb4a777db0401b883d3ef6cfa496)
- [syslog.tls_server.volterra_ca](resources--log_receiver--reference--group-001.md#canonical-9d22ae2a30e391644727ba133bac83bbe67a8d6c3357ba44814ad521b9c8cdd8)
- [syslog](resources--log_receiver--reference--group-001.md#canonical-238a4fb5c929043bd6111b3f52ebb4d83281b91492c688ac149bbfc23554692d)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)

<a id="canonical-2ebacabed412ae5ce841fd4f8398e00d9785b0ed54c2a2b0d61ca005025e087b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ddf08a1a14675b8dede83cb61ffe60fc8d128e3add277cd932e044d80f3a5e97"></a>

## syslog.tls_server.default_https_port — syslog.tls_server.default_https_port / 708805ac2133 / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)
- [Property reference](resources--log_receiver--reference--group-001.md#canonical-ef02783c49c50c98ba0aaaf5e5b83465c290bcf3f2f090952d79c6692f4536e1)
- [syslog](resources--log_receiver--reference--group-001.md#canonical-238a4fb5c929043bd6111b3f52ebb4d83281b91492c688ac149bbfc23554692d)
- [syslog.tls_server](resources--log_receiver--reference--group-001.md#canonical-e04115abcd19df912fb601a0f0928efbc4a740cb9dfe730db2e3b5eaa7b81f1e)
- syslog.tls_server.default_https_port

<a id="canonical-46d87970e939247743adedabdb7a860900149999b5806bedfe912b0bf767847b"></a>

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

<a id="canonical-0a0335db387a63c6f990bcc9d2825859cd8594a316af52055b6d8d50f12042ce"></a>

## Direct properties — syslog.tls_server.default_https_port / 708805ac2133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bdf1a394c0bb772209e68b04e0e9c5710a59416af12331efcdadf16f9f4ad186"></a>

## Next pages — syslog.tls_server.default_https_port / 708805ac2133 / 4

- [syslog.tls_server](resources--log_receiver--reference--group-001.md#canonical-e04115abcd19df912fb601a0f0928efbc4a740cb9dfe730db2e3b5eaa7b81f1e)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)

<a id="canonical-979cc51d6cbf9db9dea81fac45534d6a6b617b56f76a8f35c772ece5e5cacacb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-383673793980b0230700e7759ce11045b16ea73db959a8da60665d45ffb85bce"></a>

## syslog.tls_server.default_syslog_tls_port — syslog.tls_server.default_syslog_tls_port / 534c2dc0259a / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)
- [Property reference](resources--log_receiver--reference--group-001.md#canonical-ef02783c49c50c98ba0aaaf5e5b83465c290bcf3f2f090952d79c6692f4536e1)
- [syslog](resources--log_receiver--reference--group-001.md#canonical-238a4fb5c929043bd6111b3f52ebb4d83281b91492c688ac149bbfc23554692d)
- [syslog.tls_server](resources--log_receiver--reference--group-001.md#canonical-e04115abcd19df912fb601a0f0928efbc4a740cb9dfe730db2e3b5eaa7b81f1e)
- syslog.tls_server.default_syslog_tls_port

<a id="canonical-eadc4cc0ce99e78827fc5cc58dbf31c500ade1585ae1691e127306ac9a139c37"></a>

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
default_syslog_tls_port = {}
```

<a id="canonical-d90d62ef5a835b98ca0b33d60bc983cad41ae367417498adfcf192a4b80534a7"></a>

## Direct properties — syslog.tls_server.default_syslog_tls_port / 534c2dc0259a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-99d3704056f668a672eec9a551909583714c97df0cac359331642da038e9b5ce"></a>

## Next pages — syslog.tls_server.default_syslog_tls_port / 534c2dc0259a / 4

- [syslog.tls_server](resources--log_receiver--reference--group-001.md#canonical-e04115abcd19df912fb601a0f0928efbc4a740cb9dfe730db2e3b5eaa7b81f1e)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)

<a id="canonical-9de1e8c0460c53c63d8185ddb70cc148b23336e5c88ae28e55ffa90e97fff9f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-196b4eeb24435728774311f5deef241329ea3159bc4d54e0f6a3c6146d3b1fc5"></a>

## syslog.tls_server.mtls_disabled — syslog.tls_server.mtls_disabled / 5219aab16ef5 / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)
- [Property reference](resources--log_receiver--reference--group-001.md#canonical-ef02783c49c50c98ba0aaaf5e5b83465c290bcf3f2f090952d79c6692f4536e1)
- [syslog](resources--log_receiver--reference--group-001.md#canonical-238a4fb5c929043bd6111b3f52ebb4d83281b91492c688ac149bbfc23554692d)
- [syslog.tls_server](resources--log_receiver--reference--group-001.md#canonical-e04115abcd19df912fb601a0f0928efbc4a740cb9dfe730db2e3b5eaa7b81f1e)
- syslog.tls_server.mtls_disabled

<a id="canonical-fd71c0b248c6c2d1eec25279b9dc948b5c40490f16089b922583ffcbeb7e4b3d"></a>

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
mtls_disabled = {}
```

<a id="canonical-599dcda33202fa42c6502a1bc04af9663d3bdca840ea1d5f42029d4024f88934"></a>

## Direct properties — syslog.tls_server.mtls_disabled / 5219aab16ef5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-464ee4d00ab7be056068fe2016366773515f2f7680d3e074ef43091039cc9dbb"></a>

## Next pages — syslog.tls_server.mtls_disabled / 5219aab16ef5 / 4

- [syslog.tls_server](resources--log_receiver--reference--group-001.md#canonical-e04115abcd19df912fb601a0f0928efbc4a740cb9dfe730db2e3b5eaa7b81f1e)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)

<a id="canonical-e327963cd48210663b691b7632b7ae92705aeb4a777db0401b883d3ef6cfa496"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad7fe9a3845cf756a5fab77d5f1782d75a6b41b6942ae981c6558e999300e4e1"></a>

## syslog.tls_server.mtls_enable — syslog.tls_server.mtls_enable / d18020990906 / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)
- [Property reference](resources--log_receiver--reference--group-001.md#canonical-ef02783c49c50c98ba0aaaf5e5b83465c290bcf3f2f090952d79c6692f4536e1)
- [syslog](resources--log_receiver--reference--group-001.md#canonical-238a4fb5c929043bd6111b3f52ebb4d83281b91492c688ac149bbfc23554692d)
- [syslog.tls_server](resources--log_receiver--reference--group-001.md#canonical-e04115abcd19df912fb601a0f0928efbc4a740cb9dfe730db2e3b5eaa7b81f1e)
- syslog.tls_server.mtls_enable

<a id="canonical-3e2f31f1ca1538e82161cdc827168bab8671eb560ffce14c48d6ea86eeac94d3"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for mtls enable.

Upstream description:

TLS config for client.

Receipt-pinned upstream constraints:

```json
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
mtls_enable {
  # Configure direct properties listed below.
}
```

<a id="canonical-dd8860195dbc70e582652f0199f4d4dc23044402a51f1c32cbeb2548df0201f8"></a>

## Direct properties — syslog.tls_server.mtls_enable / d18020990906 / 3

<a id="canonical-d445c177864bdae92658faf55e6e2382b988dee5f7647c8a7eb16b131d96f5e2"></a>

<a id="canonical-606029091b253486cf88355d05ff50836fd61769d37f688ca9e6654f245a539e"></a>

## certificate property — syslog.tls_server.mtls_enable / d18020990906 / 4

Type: `"string"`. Optional.

Client certificate is PEM-encoded certificate or certificate-chain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(100, 131072),
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
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 100,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [key_url](resources--log_receiver--reference--group-001.md#canonical-fa9111f776a3ef2581494ff7a28957f172fb8f562da898a4b7ea15cb67c15974): complete subsection reference.

<a id="canonical-16e844f19c8bb4c07e21f775f2b6320f20d86a42a0209a1f6f4cfa92a1227fde"></a>

## Next pages — syslog.tls_server.mtls_enable / d18020990906 / 5

- [syslog.tls_server.mtls_enable.key_url](resources--log_receiver--reference--group-001.md#canonical-fa9111f776a3ef2581494ff7a28957f172fb8f562da898a4b7ea15cb67c15974)
- [syslog.tls_server](resources--log_receiver--reference--group-001.md#canonical-e04115abcd19df912fb601a0f0928efbc4a740cb9dfe730db2e3b5eaa7b81f1e)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)

<a id="canonical-fa9111f776a3ef2581494ff7a28957f172fb8f562da898a4b7ea15cb67c15974"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-db9e92c48c65cb64db6e53f36c66d52bbb6e4d720dff15127cf3b8dbf0ee9c90"></a>

## syslog.tls_server.mtls_enable.key_url — syslog.tls_server.mtls_enable.key_url / 3edc045afef6 / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)
- [Property reference](resources--log_receiver--reference--group-001.md#canonical-ef02783c49c50c98ba0aaaf5e5b83465c290bcf3f2f090952d79c6692f4536e1)
- [syslog](resources--log_receiver--reference--group-001.md#canonical-238a4fb5c929043bd6111b3f52ebb4d83281b91492c688ac149bbfc23554692d)
- [syslog.tls_server](resources--log_receiver--reference--group-001.md#canonical-e04115abcd19df912fb601a0f0928efbc4a740cb9dfe730db2e3b5eaa7b81f1e)
- [syslog.tls_server.mtls_enable](resources--log_receiver--reference--group-001.md#canonical-e327963cd48210663b691b7632b7ae92705aeb4a777db0401b883d3ef6cfa496)
- syslog.tls_server.mtls_enable.key_url

<a id="canonical-896e5015d97e3b9337e616c2b43b78877f3e986b97bdd89e219a44d2b12fe143"></a>

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
key_url {
  # Configure direct properties listed below.
}
```

<a id="canonical-28a285ef582471c77fb60ae667d094c4eeb5146d7d6439b6173f59cd9de50b04"></a>

## Direct properties — syslog.tls_server.mtls_enable.key_url / 3edc045afef6 / 3

- [blindfold_secret_info](resources--log_receiver--reference--group-001.md#canonical-70190e1b64feb69f528df23d7f12d32082d4fd6b3642cf31f71a4f94d0aab166): complete subsection reference.

- [clear_secret_info](resources--log_receiver--reference--group-001.md#canonical-e9fc446e871e0d739d8e3e2fddca17b9fcf9f2d46f82c4704f7e3b224492acd4): complete subsection reference.

<a id="canonical-1e1468448ceadd30bc722985ded90cd11b80d38a2474aac417309e5e72fda7b0"></a>

## Next pages — syslog.tls_server.mtls_enable.key_url / 3edc045afef6 / 4

- [syslog.tls_server.mtls_enable.key_url.blindfold_secret_info](resources--log_receiver--reference--group-001.md#canonical-70190e1b64feb69f528df23d7f12d32082d4fd6b3642cf31f71a4f94d0aab166)
- [syslog.tls_server.mtls_enable.key_url.clear_secret_info](resources--log_receiver--reference--group-001.md#canonical-e9fc446e871e0d739d8e3e2fddca17b9fcf9f2d46f82c4704f7e3b224492acd4)
- [syslog.tls_server.mtls_enable](resources--log_receiver--reference--group-001.md#canonical-e327963cd48210663b691b7632b7ae92705aeb4a777db0401b883d3ef6cfa496)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)

<a id="canonical-70190e1b64feb69f528df23d7f12d32082d4fd6b3642cf31f71a4f94d0aab166"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7033c677c1c1a48cb8229a8080cd36a83130e8e8afe987c0993dc87179ff0deb"></a>

## syslog.tls_server.mtls_enable.key_url.blindfold_secret_info — syslog.tls_server.mtls_enable.key_url.blindfold_secret_info / ef298d58a739 / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)
- [Property reference](resources--log_receiver--reference--group-001.md#canonical-ef02783c49c50c98ba0aaaf5e5b83465c290bcf3f2f090952d79c6692f4536e1)
- [syslog](resources--log_receiver--reference--group-001.md#canonical-238a4fb5c929043bd6111b3f52ebb4d83281b91492c688ac149bbfc23554692d)
- [syslog.tls_server](resources--log_receiver--reference--group-001.md#canonical-e04115abcd19df912fb601a0f0928efbc4a740cb9dfe730db2e3b5eaa7b81f1e)
- [syslog.tls_server.mtls_enable](resources--log_receiver--reference--group-001.md#canonical-e327963cd48210663b691b7632b7ae92705aeb4a777db0401b883d3ef6cfa496)
- [syslog.tls_server.mtls_enable.key_url](resources--log_receiver--reference--group-001.md#canonical-fa9111f776a3ef2581494ff7a28957f172fb8f562da898a4b7ea15cb67c15974)
- syslog.tls_server.mtls_enable.key_url.blindfold_secret_info

<a id="canonical-cc68d6314df0795772f1265497f9ed5c84ac1bdad6380b2997cf03f4270fb895"></a>

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

<a id="canonical-00bbabc057e661ac64146503983261d6c46ab8701e1f6b224d59172dab5510ae"></a>

## Direct properties — syslog.tls_server.mtls_enable.key_url.blindfold_secret_info / ef298d58a739 / 3

<a id="canonical-53508740443ad1be1db735325cd2d524beb37884ca20b57d8d076924d099f901"></a>

<a id="canonical-ccfe344dae77093592514eedaa7285430f5d5ae0680797a0ece49d9c0f9c547d"></a>

## decryption_provider property — syslog.tls_server.mtls_enable.key_url.blindfold_secret_info / ef298d58a739 / 4

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

<a id="canonical-9b6130d37b921ab933d9ba2f5b91ae06d40ccd6127680eb1723138b59ded6809"></a>

<a id="canonical-cfb466797c4032c0c1774ca132d3c72cf8b3ac10e49e0488237171eb24076c83"></a>

## location property — syslog.tls_server.mtls_enable.key_url.blindfold_secret_info / ef298d58a739 / 5

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

<a id="canonical-d89491e1f3943a07d4e21a712728196b70fb63a2e2a53d7e883ad4f544fb6f30"></a>

<a id="canonical-de626a02aac33f3d6f4efaceb6d602ffdf5428d973bb3eeed8d85023a4f39623"></a>

## store_provider property — syslog.tls_server.mtls_enable.key_url.blindfold_secret_info / ef298d58a739 / 6

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

<a id="canonical-0c314bc0e81a6d3e0204237d5d140d088fcb49cf32757120bf968a8fa1be74b4"></a>

## Next pages — syslog.tls_server.mtls_enable.key_url.blindfold_secret_info / ef298d58a739 / 7

- [syslog.tls_server.mtls_enable.key_url](resources--log_receiver--reference--group-001.md#canonical-fa9111f776a3ef2581494ff7a28957f172fb8f562da898a4b7ea15cb67c15974)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)

<a id="canonical-e9fc446e871e0d739d8e3e2fddca17b9fcf9f2d46f82c4704f7e3b224492acd4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3ec89740f028a72e11464655d5a44547917a9d5707ccb65a4c8ba02d6b25702a"></a>

## syslog.tls_server.mtls_enable.key_url.clear_secret_info — syslog.tls_server.mtls_enable.key_url.clear_secret_info / f54e030466da / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)
- [Property reference](resources--log_receiver--reference--group-001.md#canonical-ef02783c49c50c98ba0aaaf5e5b83465c290bcf3f2f090952d79c6692f4536e1)
- [syslog](resources--log_receiver--reference--group-001.md#canonical-238a4fb5c929043bd6111b3f52ebb4d83281b91492c688ac149bbfc23554692d)
- [syslog.tls_server](resources--log_receiver--reference--group-001.md#canonical-e04115abcd19df912fb601a0f0928efbc4a740cb9dfe730db2e3b5eaa7b81f1e)
- [syslog.tls_server.mtls_enable](resources--log_receiver--reference--group-001.md#canonical-e327963cd48210663b691b7632b7ae92705aeb4a777db0401b883d3ef6cfa496)
- [syslog.tls_server.mtls_enable.key_url](resources--log_receiver--reference--group-001.md#canonical-fa9111f776a3ef2581494ff7a28957f172fb8f562da898a4b7ea15cb67c15974)
- syslog.tls_server.mtls_enable.key_url.clear_secret_info

<a id="canonical-134f5d15abf06ef95f50688e83bf41b1811d439a7b764d32e13a2cbc0b207567"></a>

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

<a id="canonical-e12548ae3b8e956a36a94d92f185de1933736a15e7a40c16a93669495928edcc"></a>

## Direct properties — syslog.tls_server.mtls_enable.key_url.clear_secret_info / f54e030466da / 3

<a id="canonical-a131c324b0b59e9976656575d152dde47761b9fde7fa53aef2ea0c0c86c54f99"></a>

<a id="canonical-f40f2109c3869c3da5f9c556f5ceabf383437bc6b37de5f90b502c2fe254224b"></a>

## provider_ref property — syslog.tls_server.mtls_enable.key_url.clear_secret_info / f54e030466da / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-b8c936579dcbad63e2663818236d8cd3ecb3dcf6c5dd20d8060473c7f2ec37fc"></a>

<a id="canonical-9a868d93e40f0924e9c3c177c028d5050de7666b106d551fc20e6883d7091d65"></a>

## url property — syslog.tls_server.mtls_enable.key_url.clear_secret_info / f54e030466da / 5

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

<a id="canonical-fecb9efa4f0a56d94cf212b58f67f790613920c4ef3b4e85640bd9d0b268cce0"></a>

## Next pages — syslog.tls_server.mtls_enable.key_url.clear_secret_info / f54e030466da / 6

- [syslog.tls_server.mtls_enable.key_url](resources--log_receiver--reference--group-001.md#canonical-fa9111f776a3ef2581494ff7a28957f172fb8f562da898a4b7ea15cb67c15974)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)

<a id="canonical-9d22ae2a30e391644727ba133bac83bbe67a8d6c3357ba44814ad521b9c8cdd8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa7e800e33621026ccca7240fca42babb0ced0f1684392f5d1ed07cc34b58d74"></a>

## syslog.tls_server.volterra_ca — syslog.tls_server.volterra_ca / 997b45b9f9a4 / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)
- [Property reference](resources--log_receiver--reference--group-001.md#canonical-ef02783c49c50c98ba0aaaf5e5b83465c290bcf3f2f090952d79c6692f4536e1)
- [syslog](resources--log_receiver--reference--group-001.md#canonical-238a4fb5c929043bd6111b3f52ebb4d83281b91492c688ac149bbfc23554692d)
- [syslog.tls_server](resources--log_receiver--reference--group-001.md#canonical-e04115abcd19df912fb601a0f0928efbc4a740cb9dfe730db2e3b5eaa7b81f1e)
- syslog.tls_server.volterra_ca

<a id="canonical-3c8a84303f93f7f43a191405328472c6d77a3468fdae55a0dff504750fc94a30"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for volterra ca.

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
volterra_ca = {}
```

<a id="canonical-b87576e4e77b965f2691c04d53015bb02ab0ffa0760f364d604d84de280c62c3"></a>

## Direct properties — syslog.tls_server.volterra_ca / 997b45b9f9a4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4ffe54f0a2f1afa474e23a036cc1ac63b3564cb9eb41782738d9024f41cd57db"></a>

## Next pages — syslog.tls_server.volterra_ca / 997b45b9f9a4 / 4

- [syslog.tls_server](resources--log_receiver--reference--group-001.md#canonical-e04115abcd19df912fb601a0f0928efbc4a740cb9dfe730db2e3b5eaa7b81f1e)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)

<a id="canonical-0faaaea294d489614a46833d43b29836faa8c574d45e6c71e7f52dadc8160c78"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b544837c583e698af9736ab2c3039e2b2e43e8417edc0bb9906e61c63387def8"></a>

## syslog.udp_server — syslog.udp_server / 84b6c578745a / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)
- [Property reference](resources--log_receiver--reference--group-001.md#canonical-ef02783c49c50c98ba0aaaf5e5b83465c290bcf3f2f090952d79c6692f4536e1)
- [syslog](resources--log_receiver--reference--group-001.md#canonical-238a4fb5c929043bd6111b3f52ebb4d83281b91492c688ac149bbfc23554692d)
- syslog.udp_server

<a id="canonical-eab9d957b67babd671ef574024676442b9f5d6e61679e7d36e7a74ce4f493737"></a>

Type: `"object"`. single nested block, Optional.

UDP Server Name and Port Number. Name and port number for a UDP server.

Upstream description:

Name and port number for a UDP server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("port",
    "server_name")}
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
udp_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-1cc661e2ecfbb6b6e50534f24b87fcd577f4c5281981b278d884dccb9478eb26"></a>

## Direct properties — syslog.udp_server / 84b6c578745a / 3

<a id="canonical-2f3658cd8e9c440a14438b363b75c2a9970fd4e5bd680f94e2afa6f28120be32"></a>

<a id="canonical-2656422c925c7c0d36fa8d5bc496e8ac72c9845efcbe150a2db3f12aedfd6d41"></a>

## port property — syslog.udp_server / 84b6c578745a / 4

Type: `"number"`. Optional.

Port Number. Port number used for communication.

Upstream description:

Port number used for communication.

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
    "minimum": 1,
    "multipleOf": 1
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-5c65539356b12f6069d38c47a905c2c1a44676b4dc53ab8dc61e696d71238dd9"></a>

<a id="canonical-1bad9918745b8c399cc04fb5a5c6d8f90ea92b15a41b6d94d7ee11b4d12f81db"></a>

## server_name property — syslog.udp_server / 84b6c578745a / 5

Type: `"string"`. Optional.

Server name is fully qualified domain name or IP address of the server.

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
    "ves.io.schema.rules.string.hostname_or_ip": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname_or_ip": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-854d9bf426c9d4174cf5fe3b44346412338483c8a163810a3082747bbf167fc6"></a>

## Next pages — syslog.udp_server / 84b6c578745a / 6

- [syslog](resources--log_receiver--reference--group-001.md#canonical-238a4fb5c929043bd6111b3f52ebb4d83281b91492c688ac149bbfc23554692d)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)

<a id="canonical-9c77ae414286862fafbb9218e1483a458a6c06d4bf7bf5ec8aaca80a4ee18cb9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4124febd266b229c97d438e3cf1b6e3069de0b8772159b22dfda752a633221d0"></a>

## timeouts — timeouts / 0b623fc22e5a / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)
- [Property reference](resources--log_receiver--reference--group-001.md#canonical-ef02783c49c50c98ba0aaaf5e5b83465c290bcf3f2f090952d79c6692f4536e1)
- timeouts

<a id="canonical-82593a3b8255332a63e32f5f39afa82adfc2c4418387f48a92595cc0cffbdee3"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2f59a8100d6b71d923cad561cdbf3f8c34bee7a4c61854356244ce5eb7d00c2c"></a>

## Direct properties — timeouts / 0b623fc22e5a / 3

<a id="canonical-4cd0a9e655e22cfd7e13bc1213ea13e307f2faad193edc4a2fc92e97ecb2b6b2"></a>

<a id="canonical-1b8644682e2a64b2c526f97fde9b75edaf929f8baab1e656f4fadad0975d3f48"></a>

## create property — timeouts / 0b623fc22e5a / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-8cec0affb97d41c592d02d95a0b3f049a4a232b229be74df6a0033fbb115171b"></a>

<a id="canonical-2898f183f8b5cda23446e0945591eb2fd37ef21ca64958727948d7d9b377bf4e"></a>

## delete property — timeouts / 0b623fc22e5a / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-9b466b7c8a0a2c26b86febad03bc1dacb40d812c56af7a91f8e218e3c3ecf902"></a>

<a id="canonical-b0c633de5574251bb8a11228bce944dec1a0ec4c1762c2c6142f26b2e61ef877"></a>

## read property — timeouts / 0b623fc22e5a / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3f76d9edbc1c568a3905009a82e3183dee1edf62bac76b3a17781b881256a423"></a>

<a id="canonical-31b016fcdd14f0269f40d3fcce794d463fd0fbe8b7c1423efff468516e6f351d"></a>

## update property — timeouts / 0b623fc22e5a / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-6ed398517b0b20a350ea7cb3d808dffb8649425029e001f71cfad238990e0827"></a>

## Next pages — timeouts / 0b623fc22e5a / 8

- [Property reference](resources--log_receiver--reference--group-001.md#canonical-ef02783c49c50c98ba0aaaf5e5b83465c290bcf3f2f090952d79c6692f4536e1)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)
