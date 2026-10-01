---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-74979d25485f9a0cd0ab633d8985e298f2ad5189f61ec52d9505778d9aeab20d"></a>

## name property — stateful_service.persistent_volumes / 74c496f115e2 / 4

Type: `"string"`. Computed.

Name. Name of the volume.

Upstream description:

Name of the volume.

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z0-9]([-a-z0-9]*[a-z0-9])?$",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.dns_1123_label": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.dns_1123_label": "true"
  }
}
```

- [persistent_volume](data-sources--workload--reference--group-028.md#canonical-377c7fadf081c31382c8a80c43243b425956d994aed7a563e23934141cff9a0e): complete subsection reference.

<a id="canonical-8181cc38fb6a7b8c5bdf6aa8f0c3831bd6d94ef5b4dd1a74de86d10c7feb7b03"></a>

## Next pages — stateful_service.persistent_volumes / 74c496f115e2 / 5

- [stateful_service.persistent_volumes.persistent_volume](data-sources--workload--reference--group-028.md#canonical-377c7fadf081c31382c8a80c43243b425956d994aed7a563e23934141cff9a0e)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-377c7fadf081c31382c8a80c43243b425956d994aed7a563e23934141cff9a0e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f755a4112e010422ea8c4415b680bcbc7ccfe152214af518f865886568b5068f"></a>

## stateful_service.persistent_volumes.persistent_volume — stateful_service.persistent_volumes.persistent_volume / 16646c4d6eb1 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.persistent_volumes](data-sources--workload--reference--group-027.md#canonical-5e482719221f7b71db6737ffd54c9648576e3dfafe22e13897fb0dde9dea91b3)
- stateful_service.persistent_volumes.persistent_volume

<a id="canonical-6beb046f61585c0e90fef91dae95dcdb61cc9f1ffb89c8dc40fce5c4f587bfdd"></a>

Type: `"single"`. Computed.

Volume containing the Persistent Storage for the workload.

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

<a id="canonical-298b09e4d4b7d55cc0af5eeda784756bfed1849b7a96c57c10be3f6c6c5edc37"></a>

## Direct properties — stateful_service.persistent_volumes.persistent_volume / 16646c4d6eb1 / 3

- [mount](data-sources--workload--reference--group-028.md#canonical-294e3e2bfbcebaa372ed127ee33fcb830cc9583017c9e826637e4699e49e2800): complete subsection reference.

- [storage](data-sources--workload--reference--group-028.md#canonical-e1a8c34b077ea27d19916b93c32b7ac4cbf7852043a65d4228584d5415cf55ec): complete subsection reference.

<a id="canonical-fa7e5cb4aec52d181c5430f46957d0bd9f428bf51017da3abaebb3e267a58f1c"></a>

## Next pages — stateful_service.persistent_volumes.persistent_volume / 16646c4d6eb1 / 4

- [stateful_service.persistent_volumes.persistent_volume.mount](data-sources--workload--reference--group-028.md#canonical-294e3e2bfbcebaa372ed127ee33fcb830cc9583017c9e826637e4699e49e2800)
- [stateful_service.persistent_volumes.persistent_volume.storage](data-sources--workload--reference--group-028.md#canonical-e1a8c34b077ea27d19916b93c32b7ac4cbf7852043a65d4228584d5415cf55ec)
- [stateful_service.persistent_volumes](data-sources--workload--reference--group-027.md#canonical-5e482719221f7b71db6737ffd54c9648576e3dfafe22e13897fb0dde9dea91b3)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-294e3e2bfbcebaa372ed127ee33fcb830cc9583017c9e826637e4699e49e2800"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d7ec2878d8083cd05cb90e92d37f6512a9dc0fca5ed6deb7b10cf16c56bf370"></a>

## stateful_service.persistent_volumes.persistent_volume.mount — stateful_service.persistent_volumes.persistent_volume.mount / b24f0f3a6eca / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.persistent_volumes](data-sources--workload--reference--group-027.md#canonical-5e482719221f7b71db6737ffd54c9648576e3dfafe22e13897fb0dde9dea91b3)
- [stateful_service.persistent_volumes.persistent_volume](data-sources--workload--reference--group-028.md#canonical-377c7fadf081c31382c8a80c43243b425956d994aed7a563e23934141cff9a0e)
- stateful_service.persistent_volumes.persistent_volume.mount

<a id="canonical-179a5f6250f2c3af0f2d9a76b36fa011d0e54b23328a8d1b7f8ba4afd1235482"></a>

Type: `"single"`. Computed.

Volume mount describes how volume is mounted inside a workload.

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

<a id="canonical-f566a15697685f3152e2e9ed89c8d9076e0483c89e43b9f06e50dbd83aea9387"></a>

## Direct properties — stateful_service.persistent_volumes.persistent_volume.mount / b24f0f3a6eca / 3

<a id="canonical-f2632e0429ef33b8eccd138a6ec2f00159534ee126c79adb6305cee3b7aa41ff"></a>

<a id="canonical-02407adc70a640467512e9d85fa6f87446bfe52cb273fffc1c04aed0d2c59cba"></a>

## mode property — stateful_service.persistent_volumes.persistent_volume.mount / b24f0f3a6eca / 4

Type: `"string"`. Computed.

\[Enum: VOLUME\_MOUNT\_READ\_ONLY|VOLUME\_MOUNT\_READ\_WRITE\] Mode in which the volume should be
mounted to the workload - VOLUME\_MOUNT\_READ\_ONLY: ReadOnly Mount the volume in read-only mode -
VOLUME\_MOUNT\_READ\_WRITE: Read Write Mount the volume in read-write mode. Possible values are
\`VOLUME\_MOUNT\_READ\_ONLY\`, \`VOLUME\_MOUNT\_READ\_WRITE\`. Defaults to
\`VOLUME\_MOUNT\_READ\_ONLY\`.

Upstream description:

Mode in which the volume should be mounted to the workload

&#8203;- VOLUME\_MOUNT\_READ\_ONLY: ReadOnly

Mount the volume in read-only mode &#8203;- VOLUME\_MOUNT\_READ\_WRITE: Read Write

Mount the volume in read-write mode.

Receipt-pinned upstream constraints:

```json
{
  "default": "VOLUME_MOUNT_READ_ONLY",
  "enum": [
    "VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1c35a925f1c5996e113c2b6f2950a9cbcfec061a6409c67b6745f509763d0d8b"></a>

<a id="canonical-e0bc613b4952c953f2c11e12a6cff8d222c6098f869866e2aeb6ce0eb00f2bb5"></a>

## mount_path property — stateful_service.persistent_volumes.persistent_volume.mount / b24f0f3a6eca / 5

Type: `"string"`. Computed.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

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
    },
    "pattern": "^[^:]*$"
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
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  }
}
```

<a id="canonical-f41bc60c4591cb4d28446e9969bb2edb0b62e4ee5b084757d8f72d8bb0e20b6a"></a>

<a id="canonical-bc7d527d906ddd91814e08dfdeb9c2d15f49820ce409d02696756217f1461044"></a>

## sub_path property — stateful_service.persistent_volumes.persistent_volume.mount / b24f0f3a6eca / 6

Type: `"string"`. Computed.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Upstream description:

Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's
root).

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

<a id="canonical-f93e7d7de0968d48489292f0f6e591ad95c55befb0de6849cced08da7ea91ca5"></a>

## Next pages — stateful_service.persistent_volumes.persistent_volume.mount / b24f0f3a6eca / 7

- [stateful_service.persistent_volumes.persistent_volume](data-sources--workload--reference--group-028.md#canonical-377c7fadf081c31382c8a80c43243b425956d994aed7a563e23934141cff9a0e)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-e1a8c34b077ea27d19916b93c32b7ac4cbf7852043a65d4228584d5415cf55ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d3d6360cd28b098fc31e6d8b4dcce6739fed19879a95be8da0a28e8af2ca587"></a>

## stateful_service.persistent_volumes.persistent_volume.storage — stateful_service.persistent_volumes.persistent_volume.storage / 3c976d9ec818 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.persistent_volumes](data-sources--workload--reference--group-027.md#canonical-5e482719221f7b71db6737ffd54c9648576e3dfafe22e13897fb0dde9dea91b3)
- [stateful_service.persistent_volumes.persistent_volume](data-sources--workload--reference--group-028.md#canonical-377c7fadf081c31382c8a80c43243b425956d994aed7a563e23934141cff9a0e)
- stateful_service.persistent_volumes.persistent_volume.storage

<a id="canonical-b7c5bfcdc18a6d957886a737947b0be2ee2bde79df5cfa2fa02b5f16be892e27"></a>

Type: `"single"`. Computed.

Persistent storage configuration is used to configure Persistent Volume Claim (PVC).

Upstream description:

Persistent storage configuration is used to configure Persistent Volume Claim (PVC)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-class_name_choice": "[\"class_name\",\"default\"]"
}
```

<a id="canonical-dfeb188d5fc3cb873d9a2e1f3cbf93e17de5fa7d0063309caf83eaaba0e92bda"></a>

## Direct properties — stateful_service.persistent_volumes.persistent_volume.storage / 3c976d9ec818 / 3

<a id="canonical-0b2290d09f2b19089e1f48b6524834f237be916084acb9a7de2c4eac95c003a6"></a>

<a id="canonical-cb7bd94089eb707202058ff897611295e6c5e66ae3bf4288366cb84448006235"></a>

## access_mode property — stateful_service.persistent_volumes.persistent_volume.storage / 3c976d9ec818 / 4

Type: `"string"`. Computed.

\[Enum:
ACCESS\_MODE\_READ\_WRITE\_ONCE|ACCESS\_MODE\_READ\_WRITE\_MANY|ACCESS\_MODE\_READ\_ONLY\_MANY\]
Persistence storage access mode is used to configure access mode for persistent storage -
ACCESS\_MODE\_READ\_WRITE\_ONCE: Read Write Once Read Write Once is used to mount persistent storage
in read/write mode to exactly 1 host - ACCESS\_MODE\_READ\_WRITE\_MANY: Read Write Many Read Write
Many is used.. Possible values are \`ACCESS\_MODE\_READ\_WRITE\_ONCE\`,
\`ACCESS\_MODE\_READ\_WRITE\_MANY\`, \`ACCESS\_MODE\_READ\_ONLY\_MANY\`. Defaults to
\`ACCESS\_MODE\_READ\_WRITE\_ONCE\`.

Upstream description:

Persistence storage access mode is used to configure access mode for persistent storage

&#8203;- ACCESS\_MODE\_READ\_WRITE\_ONCE: Read Write Once

Read Write Once is used to mount persistent storage in read/write mode to exactly 1 host &#8203;-
ACCESS\_MODE\_READ\_WRITE\_MANY: Read Write Many

Read Write Many is used to mount persistent storage in read/write mode to many hosts &#8203;-
ACCESS\_MODE\_READ\_ONLY\_MANY: Read Only Many

Read Only Many is used to mount persistent storage in read-only mode to many hosts.

Receipt-pinned upstream constraints:

```json
{
  "default": "ACCESS_MODE_READ_WRITE_ONCE",
  "enum": [
    "ACCESS_MODE_READ_WRITE_ONCE",
    "ACCESS_MODE_READ_WRITE_MANY",
    "ACCESS_MODE_READ_ONLY_MANY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-cab33ae8a32d0a07973e9f394e3a2fe9efa93394c18a5b4f17c62a11a6ed0b30"></a>

<a id="canonical-8e69ff2cdc44c581942eac8db2900f6ac125006856a6b004e8b5009a35adb89f"></a>

## class_name property — stateful_service.persistent_volumes.persistent_volume.storage / 3c976d9ec818 / 5

Type: `"string"`. Computed.

Exclusive with \[default\] Use the specified class name.

Upstream description:

Exclusive with \[default\] Use the specified class name.

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

- [default](data-sources--workload--reference--group-028.md#canonical-a015abadec4058d7a4c321413feaa47e77650ae3abf0d437ecbc98016a4c25fc): complete subsection reference.

<a id="canonical-9471a56a57be97787f48d2ec66acaf6fd3c7d704758bc2b38fd8ac68e2bd2d4b"></a>

<a id="canonical-a9cdf9c86438350004e0fad839e8bcd059996ed7b7ad2ce656064e85bf24398f"></a>

## storage_size property — stateful_service.persistent_volumes.persistent_volume.storage / 3c976d9ec818 / 6

Type: `"number"`. Computed.

Size (in GiB). Size in GiB of the persistent storage.

Upstream description:

Size in GiB of the persistent storage.

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
    "ves.io.schema.rules.double.gte": "0.004",
    "ves.io.schema.rules.double.lte": "256",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.double.gte": "0.004",
    "ves.io.schema.rules.double.lte": "256",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-0857bdf94dd9aa46c7854100209dd13ab95041571d2b647b9316ef442795ccfb"></a>

## Next pages — stateful_service.persistent_volumes.persistent_volume.storage / 3c976d9ec818 / 7

- [stateful_service.persistent_volumes.persistent_volume.storage.default](data-sources--workload--reference--group-028.md#canonical-a015abadec4058d7a4c321413feaa47e77650ae3abf0d437ecbc98016a4c25fc)
- [stateful_service.persistent_volumes.persistent_volume](data-sources--workload--reference--group-028.md#canonical-377c7fadf081c31382c8a80c43243b425956d994aed7a563e23934141cff9a0e)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-a015abadec4058d7a4c321413feaa47e77650ae3abf0d437ecbc98016a4c25fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f3a11f655d1632f8e80326ec34fde1bf675b0ad5c1951be64c6d64aa866f994"></a>

## stateful_service.persistent_volumes.persistent_volume.storage.default — stateful_service.persistent_volumes.persistent_volume.storage.default / 0d720c8e4164 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.persistent_volumes](data-sources--workload--reference--group-027.md#canonical-5e482719221f7b71db6737ffd54c9648576e3dfafe22e13897fb0dde9dea91b3)
- [stateful_service.persistent_volumes.persistent_volume](data-sources--workload--reference--group-028.md#canonical-377c7fadf081c31382c8a80c43243b425956d994aed7a563e23934141cff9a0e)
- [stateful_service.persistent_volumes.persistent_volume.storage](data-sources--workload--reference--group-028.md#canonical-e1a8c34b077ea27d19916b93c32b7ac4cbf7852043a65d4228584d5415cf55ec)
- stateful_service.persistent_volumes.persistent_volume.storage.default

<a id="canonical-e0454654ce9167d914592ba3716c8a368925f7390be0dd8674b22dcb7440f44e"></a>

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

<a id="canonical-ae5b37ca286b48f3f0804aeed545402decba09c506c98411631af70fddd7a475"></a>

## Direct properties — stateful_service.persistent_volumes.persistent_volume.storage.default / 0d720c8e4164 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9a2b31e165faf9229a8f2b38dbcebe7fa1ac3dc47ae0fb5696c9160198e7d1db"></a>

## Next pages — stateful_service.persistent_volumes.persistent_volume.storage.default / 0d720c8e4164 / 4

- [stateful_service.persistent_volumes.persistent_volume.storage](data-sources--workload--reference--group-028.md#canonical-e1a8c34b077ea27d19916b93c32b7ac4cbf7852043a65d4228584d5415cf55ec)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-9cd20cbdf15c10a52dc1c08e388065f28422844a29edcef0046dc6be6e2cab36"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6bdb99c8e93b4d861150471f10d5f0dc5076af50f02a3a8735be613596ede4c5"></a>

## stateful_service.scale_to_zero — stateful_service.scale_to_zero / 46346dbce7b5 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- stateful_service.scale_to_zero

<a id="canonical-ca9b7ba1297522e581fbbf9f6a6583869dfe8de606be80ffaf3d4eff942a8b17"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for scale to zero.

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

<a id="canonical-1d5536eaf3ac4cdb94fd9c467567f88d0534d4dc94592238ac5ca17490e7bf57"></a>

## Direct properties — stateful_service.scale_to_zero / 46346dbce7b5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ab191c39e6aae6a76d9bb18837bfeebbb305862fe82e054c892d04be6f4441eb"></a>

## Next pages — stateful_service.scale_to_zero / 46346dbce7b5 / 4

- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-85f200eeee91d3e87d86544517585e13e6a11ea910756373fdfea814e6e61e66"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eb9a7ab581f357f46f2a7bc58f3b338bbb7b40c40da805231a0b0145b9708716"></a>

## stateful_service.volumes — stateful_service.volumes / 9ce54940390e / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- stateful_service.volumes

<a id="canonical-ed4a00d57d136f1702a212dcd6025931e72c4ad9ba96aae009498408c02146e4"></a>

Type: `"list"`. Computed.

Ephemeral Volumes. Ephemeral volumes for the service.

Upstream description:

Ephemeral volumes for the service.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-98e3c2f94fd55269ac1cf1f6da63e3f88e52df3bcd137dd53a299463dd82c6b5"></a>

## Direct properties — stateful_service.volumes / 9ce54940390e / 3

- [empty_dir](data-sources--workload--reference--group-028.md#canonical-0f89eddb8b79c2f4fe214907f28a21307c3b8cdc8f6997eff144015f62a2d857): complete subsection reference.

- [host_path](data-sources--workload--reference--group-028.md#canonical-7f5051b4b1dc760937eb6aa4a5214de15e2c956d4eb51b720215cb5555c8ca96): complete subsection reference.

<a id="canonical-e7f6859764e9f7410a8ac1106e5af897a5e35e61430acfc2494cf21fe93204cf"></a>

<a id="canonical-1c8284a55447c921a2acf061724a61086b68565108af4ee8271bda700e40bfa8"></a>

## name property — stateful_service.volumes / 9ce54940390e / 4

Type: `"string"`. Computed.

Name. Name of the volume.

Upstream description:

Name of the volume.

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z0-9]([-a-z0-9]*[a-z0-9])?$",
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
    "ves.io.schema.rules.string.dns_1123_label": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.dns_1123_label": "true"
  }
}
```

<a id="canonical-d859fcbda733548aaff4c239e97aa7779a7eecd5316e74ec54c727d0a0bb599c"></a>

## Next pages — stateful_service.volumes / 9ce54940390e / 5

- [stateful_service.volumes.empty_dir](data-sources--workload--reference--group-028.md#canonical-0f89eddb8b79c2f4fe214907f28a21307c3b8cdc8f6997eff144015f62a2d857)
- [stateful_service.volumes.host_path](data-sources--workload--reference--group-028.md#canonical-7f5051b4b1dc760937eb6aa4a5214de15e2c956d4eb51b720215cb5555c8ca96)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-0f89eddb8b79c2f4fe214907f28a21307c3b8cdc8f6997eff144015f62a2d857"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e527725705372cbced8af8d9166712132936b3bcb626de5a0c21cd1cd9ffcc63"></a>

## stateful_service.volumes.empty_dir — stateful_service.volumes.empty_dir / b62554cdfea4 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.volumes](data-sources--workload--reference--group-028.md#canonical-85f200eeee91d3e87d86544517585e13e6a11ea910756373fdfea814e6e61e66)
- stateful_service.volumes.empty_dir

<a id="canonical-742af86a17014c76d4339bc6c6566e70397e4f2d0248b240bd6931c84a853422"></a>

Type: `"single"`. Computed.

Volume containing a temporary directory whose lifetime is the same as a replica of a workload.

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

<a id="canonical-e0d2d42645a4fade1ca815773e36a925838b3a908d58bce456132f0c6a6588b8"></a>

## Direct properties — stateful_service.volumes.empty_dir / b62554cdfea4 / 3

- [mount](data-sources--workload--reference--group-028.md#canonical-f56a16b699d56311cffac9592a60d4590711583c172c1bbde0d71d7369c0ba5c): complete subsection reference.

<a id="canonical-50c4f6aa47912ce6dedfb75cd33bbbc40679997255e763e936511761b24e8c09"></a>

<a id="canonical-05195a314d62a48f15f4830968bba2a0683ab1325a688b5e664fb71b65d7bf4f"></a>

## size_limit property — stateful_service.volumes.empty_dir / b62554cdfea4 / 4

Type: `"number"`. Computed.

Size Limit (in GiB). Configuration parameter for size limit

Upstream description:

Configuration parameter for size limit

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
    "ves.io.schema.rules.double.lte": "10",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.double.lte": "10",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-fa090f305bc2c20d9ead64a128b500dcfa051e93cc185aed7aa75a96adda12a7"></a>

## Next pages — stateful_service.volumes.empty_dir / b62554cdfea4 / 5

- [stateful_service.volumes.empty_dir.mount](data-sources--workload--reference--group-028.md#canonical-f56a16b699d56311cffac9592a60d4590711583c172c1bbde0d71d7369c0ba5c)
- [stateful_service.volumes](data-sources--workload--reference--group-028.md#canonical-85f200eeee91d3e87d86544517585e13e6a11ea910756373fdfea814e6e61e66)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-f56a16b699d56311cffac9592a60d4590711583c172c1bbde0d71d7369c0ba5c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5bf63fb2aea5b7e60d6e5f5a182014829d5c1a162c590801876dc13de391b68c"></a>

## stateful_service.volumes.empty_dir.mount — stateful_service.volumes.empty_dir.mount / bde41b05bb54 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.volumes](data-sources--workload--reference--group-028.md#canonical-85f200eeee91d3e87d86544517585e13e6a11ea910756373fdfea814e6e61e66)
- [stateful_service.volumes.empty_dir](data-sources--workload--reference--group-028.md#canonical-0f89eddb8b79c2f4fe214907f28a21307c3b8cdc8f6997eff144015f62a2d857)
- stateful_service.volumes.empty_dir.mount

<a id="canonical-e014f469f8bc3f6f9fbd7005951acdb08a6705ff2f963b45f4659dd6fe4b94db"></a>

Type: `"single"`. Computed.

Volume mount describes how volume is mounted inside a workload.

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

<a id="canonical-e92a14668f22325166893d428a946e74edaa6575c2dcb5068134872b08063244"></a>

## Direct properties — stateful_service.volumes.empty_dir.mount / bde41b05bb54 / 3

<a id="canonical-6e137de8fb7f4eb0812cd809d6847074797ba99c9f76725bc0d22deb648c83e6"></a>

<a id="canonical-108a08212d7f4ec71ad73dc660c22d36ca9c95c6757e838d2d9d1301390c15af"></a>

## mode property — stateful_service.volumes.empty_dir.mount / bde41b05bb54 / 4

Type: `"string"`. Computed.

\[Enum: VOLUME\_MOUNT\_READ\_ONLY|VOLUME\_MOUNT\_READ\_WRITE\] Mode in which the volume should be
mounted to the workload - VOLUME\_MOUNT\_READ\_ONLY: ReadOnly Mount the volume in read-only mode -
VOLUME\_MOUNT\_READ\_WRITE: Read Write Mount the volume in read-write mode. Possible values are
\`VOLUME\_MOUNT\_READ\_ONLY\`, \`VOLUME\_MOUNT\_READ\_WRITE\`. Defaults to
\`VOLUME\_MOUNT\_READ\_ONLY\`.

Upstream description:

Mode in which the volume should be mounted to the workload

&#8203;- VOLUME\_MOUNT\_READ\_ONLY: ReadOnly

Mount the volume in read-only mode &#8203;- VOLUME\_MOUNT\_READ\_WRITE: Read Write

Mount the volume in read-write mode.

Receipt-pinned upstream constraints:

```json
{
  "default": "VOLUME_MOUNT_READ_ONLY",
  "enum": [
    "VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-6c914b3cfb0fa3dd06ed9acd4c285e4bb5a61a9819389aff4f01c13e4b201e0c"></a>

<a id="canonical-f7279358a8d7b6889f87f4acfd61b68a547cf633d3b087146b150f48f74c85dc"></a>

## mount_path property — stateful_service.volumes.empty_dir.mount / bde41b05bb54 / 5

Type: `"string"`. Computed.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

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
    },
    "pattern": "^[^:]*$"
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
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  }
}
```

<a id="canonical-0f80722e324240e6cdf0a75cebc453388052df2b63a2f8d7dd58dd6c19412091"></a>

<a id="canonical-d1664b786676084fd2863a5e1cab54ed9ccd758f2d3d6196cadd00adc694a6de"></a>

## sub_path property — stateful_service.volumes.empty_dir.mount / bde41b05bb54 / 6

Type: `"string"`. Computed.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Upstream description:

Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's
root).

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

<a id="canonical-e28c6d690c13bba7e89796b3934d9c13cd1422beb3dd7126e2029177ecfbde96"></a>

## Next pages — stateful_service.volumes.empty_dir.mount / bde41b05bb54 / 7

- [stateful_service.volumes.empty_dir](data-sources--workload--reference--group-028.md#canonical-0f89eddb8b79c2f4fe214907f28a21307c3b8cdc8f6997eff144015f62a2d857)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-7f5051b4b1dc760937eb6aa4a5214de15e2c956d4eb51b720215cb5555c8ca96"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-085f5cb742b8743db921ebc3474a816426d50ba191b0d1bfd3d164ff352b14b7"></a>

## stateful_service.volumes.host_path — stateful_service.volumes.host_path / 7536798c1712 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.volumes](data-sources--workload--reference--group-028.md#canonical-85f200eeee91d3e87d86544517585e13e6a11ea910756373fdfea814e6e61e66)
- stateful_service.volumes.host_path

<a id="canonical-aa096c6a0e5198222af2aaf5fcee1b11b9a8b5463455a4ddcd7b590adc13e1e8"></a>

Type: `"single"`. Computed.

Volume containing a host mapped path into the workload.

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

<a id="canonical-d97929aa1f80ec6bb0e8c560c9677ed2ef56a2c5c65dbdd95d9d60883af96798"></a>

## Direct properties — stateful_service.volumes.host_path / 7536798c1712 / 3

- [mount](data-sources--workload--reference--group-028.md#canonical-4ccc09a00d6b1441885074aa2030fea056ab8fc27f7f3ca2317019b085c57433): complete subsection reference.

<a id="canonical-99798b2dcfe9afa4d57c168f53d7d4a2abd918e2ad822ddad95fe1d48ffe6b23"></a>

<a id="canonical-42dd52a4d24b2b8ecfd57c84828df25f9e551de5e990de9020d2d97bca97d014"></a>

## path property — stateful_service.volumes.host_path / 7536798c1712 / 4

Type: `"string"`. Computed.

Path. Path of the directory on the host.

Upstream description:

Path of the directory on the host.

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
    },
    "minLength": 1,
    "pattern": "[^\\\\0]+"
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
    "ves.io.schema.rules.string.pattern": "[^\\\\0]+"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "[^\\\\0]+"
  }
}
```

<a id="canonical-bc3cb8fea77ca01e2cb170c75165d94eba1fffd98d4d6b215fde4a8b305d0278"></a>

## Next pages — stateful_service.volumes.host_path / 7536798c1712 / 5

- [stateful_service.volumes.host_path.mount](data-sources--workload--reference--group-028.md#canonical-4ccc09a00d6b1441885074aa2030fea056ab8fc27f7f3ca2317019b085c57433)
- [stateful_service.volumes](data-sources--workload--reference--group-028.md#canonical-85f200eeee91d3e87d86544517585e13e6a11ea910756373fdfea814e6e61e66)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-4ccc09a00d6b1441885074aa2030fea056ab8fc27f7f3ca2317019b085c57433"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7108b6310ecc1b5ed997cc6c3443a283adf3d06004ba91692859340e4439f23d"></a>

## stateful_service.volumes.host_path.mount — stateful_service.volumes.host_path.mount / a0c126b43c24 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.volumes](data-sources--workload--reference--group-028.md#canonical-85f200eeee91d3e87d86544517585e13e6a11ea910756373fdfea814e6e61e66)
- [stateful_service.volumes.host_path](data-sources--workload--reference--group-028.md#canonical-7f5051b4b1dc760937eb6aa4a5214de15e2c956d4eb51b720215cb5555c8ca96)
- stateful_service.volumes.host_path.mount

<a id="canonical-000ee37b08ef2d4b15f2fe2de50fa5c92e572a936cb356f878b427160c8fdd78"></a>

Type: `"single"`. Computed.

Volume mount describes how volume is mounted inside a workload.

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

<a id="canonical-f1be3efd9d38d2eaa11602e607d0ce59db778b84c55c96267f8f877f626d870e"></a>

## Direct properties — stateful_service.volumes.host_path.mount / a0c126b43c24 / 3

<a id="canonical-0a4dd16fbdae206bdaa6376ade06ce8771ec4a2f1e5dad4744cc7e067f856a80"></a>

<a id="canonical-b6d94325263796005f6cae8c7de1f846d93a4af6e8558b1a24b6b064c59839bc"></a>

## mode property — stateful_service.volumes.host_path.mount / a0c126b43c24 / 4

Type: `"string"`. Computed.

\[Enum: VOLUME\_MOUNT\_READ\_ONLY|VOLUME\_MOUNT\_READ\_WRITE\] Mode in which the volume should be
mounted to the workload - VOLUME\_MOUNT\_READ\_ONLY: ReadOnly Mount the volume in read-only mode -
VOLUME\_MOUNT\_READ\_WRITE: Read Write Mount the volume in read-write mode. Possible values are
\`VOLUME\_MOUNT\_READ\_ONLY\`, \`VOLUME\_MOUNT\_READ\_WRITE\`. Defaults to
\`VOLUME\_MOUNT\_READ\_ONLY\`.

Upstream description:

Mode in which the volume should be mounted to the workload

&#8203;- VOLUME\_MOUNT\_READ\_ONLY: ReadOnly

Mount the volume in read-only mode &#8203;- VOLUME\_MOUNT\_READ\_WRITE: Read Write

Mount the volume in read-write mode.

Receipt-pinned upstream constraints:

```json
{
  "default": "VOLUME_MOUNT_READ_ONLY",
  "enum": [
    "VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-f0665626ee22cc23b51eda39a6e261f483f4edef4e0f7d9e958fae94985405d7"></a>

<a id="canonical-dc55bee251d265f2231d1111234a930d6d299847806192689f5a24897a9f074e"></a>

## mount_path property — stateful_service.volumes.host_path.mount / a0c126b43c24 / 5

Type: `"string"`. Computed.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

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
    },
    "pattern": "^[^:]*$"
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
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  }
}
```

<a id="canonical-94c6ef2f7da6c799f8e71f72d38311bc53b0f2a6e3f6611b06e5a37127e70a69"></a>

<a id="canonical-11e4550f2fb9c5a6e8e47f0940f55b6ce54a0abedb1880d7221fe28d4bdc537d"></a>

## sub_path property — stateful_service.volumes.host_path.mount / a0c126b43c24 / 6

Type: `"string"`. Computed.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Upstream description:

Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's
root).

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

<a id="canonical-ee1295ab3c51848a6a39cc0189749f5b1fbf3c1d1efd5366b4cbbec70b036026"></a>

## Next pages — stateful_service.volumes.host_path.mount / a0c126b43c24 / 7

- [stateful_service.volumes.host_path](data-sources--workload--reference--group-028.md#canonical-7f5051b4b1dc760937eb6aa4a5214de15e2c956d4eb51b720215cb5555c8ca96)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
