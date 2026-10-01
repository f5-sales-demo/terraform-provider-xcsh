---
page_title: "xcsh_voltstack_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_voltstack_site reference."
---

# xcsh_voltstack_site reference

<a id="canonical-2e7efa530cb935449f7abc04f2e586939a01970616f0f58d8aac2f75b16aea96"></a>

## performance_policy property — custom_storage_config.storage_class_list.storage_classes.hpe_storage / 7e6b1eddd947 / 13

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

The name of the performance policy to assign to the volume.

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

<a id="canonical-906a242c22e5a41e014ef4499b7992c891295f84d012735acda5771c7fd8157d"></a>

<a id="canonical-4f2180b127a74ff70c593bf034edf0bc7f35f7311498b840bc3a213044362651"></a>

## pool property — custom_storage_config.storage_class_list.storage_classes.hpe_storage / 7e6b1eddd947 / 14

Type: `"string"`. Computed.

The name of the pool in which to place the volume.

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

<a id="canonical-501862e2ed9947aba65b478c000b34b942ba5dad6cbb4aed565d1cc380632c14"></a>

<a id="canonical-663329d30b799b929682a23ed349000faf3eb1f9a1620004f155616dad78e866"></a>

## protection_template property — custom_storage_config.storage_class_list.storage_classes.hpe_storage / 7e6b1eddd947 / 15

Type: `"string"`. Computed.

The name of the performance policy to assign to the volume.

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

<a id="canonical-debe16c60d9b4d75ac88f7c03f2f8e481079feb83b0eff1f9bbab148dd9ae653"></a>

<a id="canonical-951736e781796035404d02570a026ea45574ced643b370bd0701088dccfa8b7b"></a>

## secret_name property — custom_storage_config.storage_class_list.storage_classes.hpe_storage / 7e6b1eddd947 / 16

Type: `"string"`. Computed.

The SecretName parameter is used to identify name of secret to identify backend storage's auth
information.

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

<a id="canonical-4dfef60f5c45c441970d3da234ac07185ea659b241def557f410ae40d3f7b96f"></a>

<a id="canonical-f89a85fc28cf61674c7e94682641903ce5786143250d51ce60417d54a718d1dd"></a>

## secret_namespace property — custom_storage_config.storage_class_list.storage_classes.hpe_storage / 7e6b1eddd947 / 17

Type: `"string"`. Computed.

The SecretNamespace parameter is used to identify name of namespace where secret resides.

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

<a id="canonical-fd998d278a687e53fdebe661392f36cf2241a2169a647a3753a27bd778ab9085"></a>

<a id="canonical-b103ed2dcd2bd835fdc5f9313878c55e93f711ba9322d11f9ec0311fb8689ab1"></a>

## sync_on_detach property — custom_storage_config.storage_class_list.storage_classes.hpe_storage / 7e6b1eddd947 / 18

Type: `"bool"`. Computed.

Indicates that a snapshot of the volume should be synced to the replication partner each time it is
detached from a node.

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

<a id="canonical-b8045483de16cb31e2d2b6ad186229d36648374fb16b39b3e11f66dbe72eb8ae"></a>

<a id="canonical-9c94b804134f7dc1b1b4c3081b67dc95c08754fdb4398eb5913f23ff3fdbfdb4"></a>

## thick property — custom_storage_config.storage_class_list.storage_classes.hpe_storage / 7e6b1eddd947 / 19

Type: `"bool"`. Computed.

Indicates that the volume should be thick provisioned.

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

<a id="canonical-adfb0ca29f2036787261a8324fa9d259598d162a7afe862636aa5ec2280d4976"></a>

## Next pages — custom_storage_config.storage_class_list.storage_classes.hpe_storage / 7e6b1eddd947 / 20

- [custom_storage_config.storage_class_list.storage_classes](data-sources--voltstack_site--reference--group-005.md#canonical-b4622fcf5268abec23d9ee2e18b0190462b934556d66a842db0dc62e617cc43e)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-de2d633db39f5d50197bddf30599beea958f21119c2c4e9f2f8050820a1295d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-42ebeea3b48489561e6a8adf35596fc92633889e2f2aeeff9b67fb051709749c"></a>

## custom_storage_config.storage_class_list.storage_classes.netapp_trident — custom_storage_config.storage_class_list.storage_classes.netapp_trident / 768b23b1dfdc / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_class_list](data-sources--voltstack_site--reference--group-005.md#canonical-37c84ca63cefb321ab7815b54f13b3b9a071543b9d0f4b454b3469e0bf16137f)
- [custom_storage_config.storage_class_list.storage_classes](data-sources--voltstack_site--reference--group-005.md#canonical-b4622fcf5268abec23d9ee2e18b0190462b934556d66a842db0dc62e617cc43e)
- custom_storage_config.storage_class_list.storage_classes.netapp_trident

<a id="canonical-9abb0853b809e154c204d698ace69cdacd2cdd1485d9ae714534b6a0d06a668b"></a>

Type: `"single"`. Computed.

Storage class Device configuration for NetApp Trident.

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

<a id="canonical-1d796e4043ceb08eeb76aab937d758c5c3e6af5a8b5062b7e9e629ac139943a1"></a>

## Direct properties — custom_storage_config.storage_class_list.storage_classes.netapp_trident / 768b23b1dfdc / 3

- [selector](data-sources--voltstack_site--reference--group-006.md#canonical-d954396a5874671f0c06f2ef0ea57fb18c5c95c5e44f0cff976010ab0d847a76): complete subsection reference.

<a id="canonical-f65aefe48624d61de8667f4e9ab921058e572a10bc37c0c2b4084d1c9a5605c4"></a>

<a id="canonical-87141055eb8297b0b39a1d7ec3233c357e31cf3e3e16526a9970d6dd3168dca4"></a>

## storage_pools property — custom_storage_config.storage_class_list.storage_classes.netapp_trident / 768b23b1dfdc / 4

Type: `"string"`. Computed.

The storagePools parameter is used to further restrict the set of pools that match any specified
attributes.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512"
  }
}
```

<a id="canonical-2003478294647ffee9c5a075757461f54c0911c9fadd607b80ca30ede52a58f0"></a>

## Next pages — custom_storage_config.storage_class_list.storage_classes.netapp_trident / 768b23b1dfdc / 5

- [custom_storage_config.storage_class_list.storage_classes.netapp_trident.selector](data-sources--voltstack_site--reference--group-006.md#canonical-d954396a5874671f0c06f2ef0ea57fb18c5c95c5e44f0cff976010ab0d847a76)
- [custom_storage_config.storage_class_list.storage_classes](data-sources--voltstack_site--reference--group-005.md#canonical-b4622fcf5268abec23d9ee2e18b0190462b934556d66a842db0dc62e617cc43e)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-d954396a5874671f0c06f2ef0ea57fb18c5c95c5e44f0cff976010ab0d847a76"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-715d4620d82859fb0de8173375d52f9a934efdb10670b9315efb6528f5550b98"></a>

## custom_storage_config.storage_class_list.storage_classes.netapp_trident.selector — custom_storage_config.storage_class_list.storage_classes.netapp_trident.selector / 7a7c7afe926a / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_class_list](data-sources--voltstack_site--reference--group-005.md#canonical-37c84ca63cefb321ab7815b54f13b3b9a071543b9d0f4b454b3469e0bf16137f)
- [custom_storage_config.storage_class_list.storage_classes](data-sources--voltstack_site--reference--group-005.md#canonical-b4622fcf5268abec23d9ee2e18b0190462b934556d66a842db0dc62e617cc43e)
- [custom_storage_config.storage_class_list.storage_classes.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-de2d633db39f5d50197bddf30599beea958f21119c2c4e9f2f8050820a1295d5)
- custom_storage_config.storage_class_list.storage_classes.netapp_trident.selector

<a id="canonical-48d7657ef7f0929e54cd15be836a2c5b3f21a3c0c61b41d0ce4daa15ae7d50f5"></a>

Type: `"single"`. Computed.

Using the Selector field, each StorageClass calls out which virtual pool(s) may be used to host a
volume. The volume will have the aspects defined in the chosen virtual pool.

Upstream description:

Using the Selector field, each StorageClass calls out which virtual pool(s) may be used to host a
volume. The volume will have the aspects defined in the chosen virtual pool.

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

<a id="canonical-ad1e4bedec7d71be7fc0130ff8e23854deb068ce3522eb29dc37221203e849b0"></a>

## Direct properties — custom_storage_config.storage_class_list.storage_classes.netapp_trident.selector / 7a7c7afe926a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-31a16499754449499dcce9813bb8de3dceef5ad0b30d3baf80b42dcfe69e15be"></a>

## Next pages — custom_storage_config.storage_class_list.storage_classes.netapp_trident.selector / 7a7c7afe926a / 4

- [custom_storage_config.storage_class_list.storage_classes.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-de2d633db39f5d50197bddf30599beea958f21119c2c4e9f2f8050820a1295d5)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-a7c4ad912806ed72d3e5b413bffbdcf56759b9723bf1ce4e3f89fc646b8aaf8d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-01314f2e4c5c93119a7fbc138f50b120374e580a63650182e25f191247ece888"></a>

## custom_storage_config.storage_class_list.storage_classes.pure_service_orchestrator — custom_storage_config.storage_class_list.storage_classes.pure_service_orchestrat / 4a9ae21993a9 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_class_list](data-sources--voltstack_site--reference--group-005.md#canonical-37c84ca63cefb321ab7815b54f13b3b9a071543b9d0f4b454b3469e0bf16137f)
- [custom_storage_config.storage_class_list.storage_classes](data-sources--voltstack_site--reference--group-005.md#canonical-b4622fcf5268abec23d9ee2e18b0190462b934556d66a842db0dc62e617cc43e)
- custom_storage_config.storage_class_list.storage_classes.pure_service_orchestrator

<a id="canonical-e782d91be8487ece3e46209fe50fd4b646a979dc6e94c3a8a14be2097a2c0cd2"></a>

Type: `"single"`. Computed.

Storage class Device configuration for Pure Service Orchestrator.

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

<a id="canonical-bb3525d7f06b76d390de53cdb4ac181d8913df56126ab0a1f8a2f88f4e2634b3"></a>

## Direct properties — custom_storage_config.storage_class_list.storage_classes.pure_service_orchestrat / 4a9ae21993a9 / 3

<a id="canonical-ef45539edbc9feced619b23568d0e6e3871c45646b9f6c2f33baac2845b7a098"></a>

<a id="canonical-1452f66e0fe3cbe61d4995a09546b933d673bd5e9d465d1b2c8bd07233ebfd75"></a>

## backend property — custom_storage_config.storage_class_list.storage_classes.pure_service_orchestrat / 4a9ae21993a9 / 4

Type: `"string"`. Computed.

\[Enum: block|file\] Defines type of Pure storage backend block or file. The volume will have the
aspects defined in the chosen virtual pool. Possible values are \`block\`, \`file\`.

Upstream description:

Defines type of Pure storage backend block or file. The volume will have the aspects defined in the
chosen virtual pool.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "block",
    "file"
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
    "ves.io.schema.rules.string.in": "[\\\"block\\\",\\\"file\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"block\\\",\\\"file\\\"]"
  }
}
```

<a id="canonical-b67b24365310884c54c5ff9334682dd11e1c118d3da210b963a94595c4511efb"></a>

<a id="canonical-72c33ae5876b7277025b2027c5bd666c3329baa79ac19055a129264cd937afac"></a>

## bandwidth_limit property — custom_storage_config.storage_class_list.storage_classes.pure_service_orchestrat / 4a9ae21993a9 / 5

Type: `"string"`. Computed.

It must be between 1 MB/s and 512 GB/s. Enter the size as a number (bytes must be multiple of 512)
or number with a single character unit symbol. Valid unit symbols are K, M, G, representing KiB,
MiB, and GiB.

Upstream description:

It must be between 1 MB/s and 512 GB/s. Enter the size as a number (bytes must be multiple of 512)
or number with a single character unit symbol. Valid unit symbols are K, M, G, representing KiB,
MiB, and GiB.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 12,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 12,
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
    "ves.io.schema.rules.string.max_len": "12"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "12"
  }
}
```

<a id="canonical-377fee08e04d87dba9f2f7e1b7dd7d7e2f5de0e844f9c8e92482eafa07463baf"></a>

<a id="canonical-cb2a418c58edf80f10f307014ad87728955419de73864cfe621d813ee1417d2a"></a>

## iops_limit property — custom_storage_config.storage_class_list.storage_classes.pure_service_orchestrat / 4a9ae21993a9 / 6

Type: `"number"`. Computed.

Enable IOPS limitation. It must be between 100 and 100 million. If value is 0, IOPS limit is not
defined.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100000000,
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
    "ves.io.schema.rules.uint32.ranges": "0,100-100000000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,100-100000000"
  }
}
```

<a id="canonical-e2f9b45904fcf3b5dea6db23b15c5a10222bd613a388fa88391e9187326fd0f6"></a>

## Next pages — custom_storage_config.storage_class_list.storage_classes.pure_service_orchestrat / 4a9ae21993a9 / 7

- [custom_storage_config.storage_class_list.storage_classes](data-sources--voltstack_site--reference--group-005.md#canonical-b4622fcf5268abec23d9ee2e18b0190462b934556d66a842db0dc62e617cc43e)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ca5b1515800dd66b82f4aadf4e0633fbe33e08dfac6bae9b1879ec6c95c9695"></a>

## custom_storage_config.storage_device_list — custom_storage_config.storage_device_list / 558b2f22a283 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- custom_storage_config.storage_device_list

<a id="canonical-946731d387e02e40aa2f29a72d728e6865c88e53c341d3a380708277812a4316"></a>

Type: `"single"`. Computed.

Add additional custom storage classes in Kubernetes for this fleet.

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

<a id="canonical-31193b4dfb0bcd377f113c612efb10619ea9cc91d434f016e6e813f68ae807a6"></a>

## Direct properties — custom_storage_config.storage_device_list / 558b2f22a283 / 3

- [storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078): complete subsection reference.

<a id="canonical-4f7dd8dc9e21fa2688508db22ba8d324477279b42f52ecb73ccab649ae3405a5"></a>

## Next pages — custom_storage_config.storage_device_list / 558b2f22a283 / 4

- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf036a9d4f6c0f1ee32627004f77a4beafa804f01434edbd13b8b59db1f406fc"></a>

## custom_storage_config.storage_device_list.storage_devices — custom_storage_config.storage_device_list.storage_devices / 087d60a2f8b3 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- custom_storage_config.storage_device_list.storage_devices

<a id="canonical-d7472a5e7e60ad3bba4d9fc5c012233a96e0e70333776374363542177d2c4ef0"></a>

Type: `"list"`. Computed.

List of Storage Devices. List of custom storage devices.

Upstream description:

List of custom storage devices.

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

<a id="canonical-b4eb2a9bbcc8aa8aed70789f3f9735385f53329a7bb5d8f18d564fab5c4c5602"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices / 087d60a2f8b3 / 3

<a id="canonical-d55f238c73cbccb82237a8625241df5a93eb3a053d1e1e1bc0115b967aec273f"></a>

<a id="canonical-042590f00a78328e78d6e302305cccb140c1e07df95c0ceb2b23b1890a916822"></a>

## advanced_advanced_parameters property — custom_storage_config.storage_device_list.storage_devices / 087d60a2f8b3 / 4

Type: `["map", "string"]`. Computed.

Advanced Parameters. Map of parameter name and string value.

Upstream description:

Map of parameter name and string value.

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
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [custom_storage](data-sources--voltstack_site--reference--group-006.md#canonical-93ef5ca6c6c59216f208006baac799f82fb79c76df4b9aed21f1b838f8c08ad1): complete subsection reference.

- [hpe_storage](data-sources--voltstack_site--reference--group-006.md#canonical-67db46fb3b026505815885fa5b5ca058dcefb83ebd8279d77606ead400c459f3): complete subsection reference.

- [netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27): complete subsection reference.

- [pure_service_orchestrator](data-sources--voltstack_site--reference--group-007.md#canonical-678801c40c6f67b9dfaf26614f3fc8e471a101a0988408bcdab7079e531d9a28): complete subsection reference.

<a id="canonical-879e330f09206fe6b97205101d40433441f8f0aa214f85645e8f0b55f34268ae"></a>

<a id="canonical-c689d4651aa214f06370663d51e577e115ae6a192f2686a8f488aff212380d32"></a>

## storage_device property — custom_storage_config.storage_device_list.storage_devices / 087d60a2f8b3 / 5

Type: `"string"`. Computed.

Storage Device. Storage device and device unit.

Upstream description:

Storage device and device unit.

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

<a id="canonical-2fd42d796f738daedaf1cf64b8b73a5b04e6c4f6e70ee8596d232d6e3e4d270a"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices / 087d60a2f8b3 / 6

- [custom_storage_config.storage_device_list.storage_devices.custom_storage](data-sources--voltstack_site--reference--group-006.md#canonical-93ef5ca6c6c59216f208006baac799f82fb79c76df4b9aed21f1b838f8c08ad1)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](data-sources--voltstack_site--reference--group-006.md#canonical-67db46fb3b026505815885fa5b5ca058dcefb83ebd8279d77606ead400c459f3)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](data-sources--voltstack_site--reference--group-007.md#canonical-678801c40c6f67b9dfaf26614f3fc8e471a101a0988408bcdab7079e531d9a28)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-93ef5ca6c6c59216f208006baac799f82fb79c76df4b9aed21f1b838f8c08ad1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c28ee09c3166a81d69e8f4f9d20bb7e9d5dc22bc09d53cc7263284e6abd48ef9"></a>

## custom_storage_config.storage_device_list.storage_devices.custom_storage — custom_storage_config.storage_device_list.storage_devices.custom_storage / 490b032b341a / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- custom_storage_config.storage_device_list.storage_devices.custom_storage

<a id="canonical-be34e26255f3e1b1e27d4897f54f1cfa0e0d2825ddf182dcea425463504a5fc6"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for custom storage.

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

<a id="canonical-1e2d9571ba397b83dff2dc7537fc608e989112943f585d4942f1113c9d878e40"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.custom_storage / 490b032b341a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c0bc10967e7fff80ea5cc4f3fdaf969c03f22f4112dee13a4a6c52fd5a08a76b"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.custom_storage / 490b032b341a / 4

- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-67db46fb3b026505815885fa5b5ca058dcefb83ebd8279d77606ead400c459f3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab99931f23db879a5d68620b1b38f88938ff028eb2c8973dd8ed994d16989aab"></a>

## custom_storage_config.storage_device_list.storage_devices.hpe_storage — custom_storage_config.storage_device_list.storage_devices.hpe_storage / e8400ba4fb35 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- custom_storage_config.storage_device_list.storage_devices.hpe_storage

<a id="canonical-4b60e712b2701178d8e57a50bdd53a25d33d4c13a6123ea5c6b15d5eab4ae215"></a>

Type: `"single"`. Computed.

Configuration parameter for hpe storage.

Upstream description:

Device configuration for HPE Storage.

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

<a id="canonical-a066a317a653f2501a8422611216fcf2aced9f68151eb900d684ceab2d8e2c0f"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.hpe_storage / e8400ba4fb35 / 3

<a id="canonical-cf16a389bcf5087e9890a2ff339d5a21d08c624c351903571b89e7e667bd9451"></a>

<a id="canonical-5d56c7b8192d7ff7bce29d85dcba8daa83a919790a60117b67f734c84c5232b3"></a>

## api_server_port property — custom_storage_config.storage_device_list.storage_devices.hpe_storage / e8400ba4fb35 / 4

Type: `"number"`. Computed.

Storage server Port. Enter Storage Server Port.

Upstream description:

Enter Storage Server Port.

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

- [iscsi_chap_password](data-sources--voltstack_site--reference--group-006.md#canonical-1ac3b2fd2343ef7755a71f0d06bb1000262771b11882c83861a2e86b35dbbf9f): complete subsection reference.

<a id="canonical-bbcc6fd26ee263d2b635fa2309f3d59c5737462c218894e21c94ea0ffaabdb18"></a>

<a id="canonical-b6e6f821e60ee5988a20a15c723d745bd6567519e4f5c99341061946d20eb56c"></a>

## iscsi_chap_user property — custom_storage_config.storage_device_list.storage_devices.hpe_storage / e8400ba4fb35 / 5

Type: `"string"`. Computed.

Chap Username to connect to the HPE storage.

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

- [password](data-sources--voltstack_site--reference--group-006.md#canonical-7a31e682adc92925997c4ede6355b5c0afd73173d5cd7a468095c33970cecfdc): complete subsection reference.

<a id="canonical-90ecb2a8a9cd8ca0c084082a9e2436ea6ab6b98d68988e47da26cf1b4f7eef09"></a>

<a id="canonical-3097ab40e3660826ead6eea46523838297b23f6e29bb89dc485bb6308c6e1771"></a>

## storage_server_ip_address property — custom_storage_config.storage_device_list.storage_devices.hpe_storage / e8400ba4fb35 / 6

Type: `"string"`. Computed.

Storage Server IP address. Enter storage server IP address.

Upstream description:

Enter storage server IP address.

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

<a id="canonical-e31dbc443f4e83752c7713a20e6064952fe5239f102d26aff534e0f97a02e7f2"></a>

<a id="canonical-4c049f4cb914732e5d5434fb18fb35c92a7cfd3b22c2b52ab9962c999261a4a7"></a>

## storage_server_name property — custom_storage_config.storage_device_list.storage_devices.hpe_storage / e8400ba4fb35 / 7

Type: `"string"`. Computed.

Storage Server Name. Enter storage server Name.

Upstream description:

Enter storage server Name.

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

<a id="canonical-f761c7026a2918085e3c5f57c04a1ec98789455433752b155ff118e4a6b73f0f"></a>

<a id="canonical-e93eff379dd93d501df4cc7777a8a5e12b29943052285cad68865385bdc9c177"></a>

## username property — custom_storage_config.storage_device_list.storage_devices.hpe_storage / e8400ba4fb35 / 8

Type: `"string"`. Computed.

Username to connect to the HPE storage management IP.

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

<a id="canonical-f6bf2ea5cc3ae42cb65cb7765e8fb399e801e27ce2606cb3b292dfaf6fe43620"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.hpe_storage / e8400ba4fb35 / 9

- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](data-sources--voltstack_site--reference--group-006.md#canonical-1ac3b2fd2343ef7755a71f0d06bb1000262771b11882c83861a2e86b35dbbf9f)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.password](data-sources--voltstack_site--reference--group-006.md#canonical-7a31e682adc92925997c4ede6355b5c0afd73173d5cd7a468095c33970cecfdc)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-1ac3b2fd2343ef7755a71f0d06bb1000262771b11882c83861a2e86b35dbbf9f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-588e06f21d62dc764d3dcf4a2af32896c33b24839a1092475a4f10641f1d93c7"></a>

## custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password — custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap / dff10ad1bb64 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](data-sources--voltstack_site--reference--group-006.md#canonical-67db46fb3b026505815885fa5b5ca058dcefb83ebd8279d77606ead400c459f3)
- custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password

<a id="canonical-5ab7d3b962dede280ddd427194b7cb2ea016a610a7a73f74edea3b519ac28d7d"></a>

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

<a id="canonical-d5c8998afb190f7212bdb2cc6962a2af96fce0f178dae1f9b2e818741d30d1cd"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap / dff10ad1bb64 / 3

- [blindfold_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-697aafc4852452dff57d63dc9334f68fa687b767ce75e38b7f4eca7340d9d474): complete subsection reference.

- [clear_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-5d328a830d3bf049d29b48b53952761864e60c92a4b9dc8607d943caacdfb11e): complete subsection reference.

<a id="canonical-e265c7d6a337626d4961ef55d56b12ef1669ab7b350bc3c3d07a93d69e385e75"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap / dff10ad1bb64 / 4

- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-697aafc4852452dff57d63dc9334f68fa687b767ce75e38b7f4eca7340d9d474)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-5d328a830d3bf049d29b48b53952761864e60c92a4b9dc8607d943caacdfb11e)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](data-sources--voltstack_site--reference--group-006.md#canonical-67db46fb3b026505815885fa5b5ca058dcefb83ebd8279d77606ead400c459f3)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-697aafc4852452dff57d63dc9334f68fa687b767ce75e38b7f4eca7340d9d474"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dcd680c10a98ae6ba2b02873a1530b535dfbe6974ed43335ebcf1a717eb1d6fd"></a>

## custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info — custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap / b7609d36b69d / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](data-sources--voltstack_site--reference--group-006.md#canonical-67db46fb3b026505815885fa5b5ca058dcefb83ebd8279d77606ead400c459f3)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](data-sources--voltstack_site--reference--group-006.md#canonical-1ac3b2fd2343ef7755a71f0d06bb1000262771b11882c83861a2e86b35dbbf9f)
- custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info

<a id="canonical-0f314bc9f50a9f58706907ff39face16b16621488a8717019d6e714a75257faa"></a>

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

<a id="canonical-9497aac61090b060495bf2f6a8f453b65a8c193676d929c5833069e25aa3d657"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap / b7609d36b69d / 3

<a id="canonical-bee29b958e00f9dfa3ebfd99db43d2665ea7b50b05e8fb71d6604b552cc1e2af"></a>

<a id="canonical-853dc63b3c8a7c209e75144d41790d4512ee57b30e261eb28a024d9e56d3c23d"></a>

## decryption_provider property — custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap / b7609d36b69d / 4

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

<a id="canonical-3e56618b716d019c40092be5ce4920c27c28424852c3a5fea155bc40da129e16"></a>

<a id="canonical-cf4491704b1223d9313ab25f1d9fb8dae792446c786913242af1eb0338bed2d4"></a>

## location property — custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap / b7609d36b69d / 5

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

<a id="canonical-5814273773e9f91e11743f44135acd4bf9af54e6d80615cc8f240f73ea63af19"></a>

<a id="canonical-cfc65fc47b10cbdfaa7816e81c3bc4215cc1288e663f2dfdfcb7d0fa40ee7127"></a>

## store_provider property — custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap / b7609d36b69d / 6

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

<a id="canonical-b07a3f12c930562f66b399b61fac2d4b8774b002f10cb3e463f27fbdfa359563"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap / b7609d36b69d / 7

- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](data-sources--voltstack_site--reference--group-006.md#canonical-1ac3b2fd2343ef7755a71f0d06bb1000262771b11882c83861a2e86b35dbbf9f)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-5d328a830d3bf049d29b48b53952761864e60c92a4b9dc8607d943caacdfb11e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-957f5808ff19b2dbd5ff15e1bb684f0dfee0cc4615a55a6aa342ab66f0620a85"></a>

## custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info — custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap / 2a941483a087 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](data-sources--voltstack_site--reference--group-006.md#canonical-67db46fb3b026505815885fa5b5ca058dcefb83ebd8279d77606ead400c459f3)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](data-sources--voltstack_site--reference--group-006.md#canonical-1ac3b2fd2343ef7755a71f0d06bb1000262771b11882c83861a2e86b35dbbf9f)
- custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info

<a id="canonical-89469e02cd0580fb6a27866e7a103e7566140a23fc3772ddfc8489bd3d19ce7d"></a>

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

<a id="canonical-991a51ae16cac2e4a20234dc644721e532df6e81bb54b63eea312e4bff2a965b"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap / 2a941483a087 / 3

<a id="canonical-5f923ff41d33bf0ef942c7baa7251d47117bcea8497faa05cc0b14aced356c81"></a>

<a id="canonical-f3590290ac4ddc125f956eaf6001c15561c141c96c675f2dfc9983df8cb2f605"></a>

## provider_ref property — custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap / 2a941483a087 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-95625d24c936a6ebf49563fd688230a72685987b6e638444cf5faa450461be0a"></a>

<a id="canonical-3fed6a66ad6d0b8986afa01f2ab58c4b05d0d3dea5024b2f36f3bee47e4aebbe"></a>

## url property — custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap / 2a941483a087 / 5

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

<a id="canonical-2d6e68403ef73a9e50918da60a52a87e059bdc79a055251fc41869c30a8538ea"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap / 2a941483a087 / 6

- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](data-sources--voltstack_site--reference--group-006.md#canonical-1ac3b2fd2343ef7755a71f0d06bb1000262771b11882c83861a2e86b35dbbf9f)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-7a31e682adc92925997c4ede6355b5c0afd73173d5cd7a468095c33970cecfdc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4811a389a2fa71c47af5b145a0d95fec9dc11eba1fce0ec59960b681d812e14c"></a>

## custom_storage_config.storage_device_list.storage_devices.hpe_storage.password — custom_storage_config.storage_device_list.storage_devices.hpe_storage.password / d03b961e7fcc / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](data-sources--voltstack_site--reference--group-006.md#canonical-67db46fb3b026505815885fa5b5ca058dcefb83ebd8279d77606ead400c459f3)
- custom_storage_config.storage_device_list.storage_devices.hpe_storage.password

<a id="canonical-081d0a31ce4930aff501c9dee61fcee1e68218c40ba0cf426c7daeb481aaaec8"></a>

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

<a id="canonical-373941544e019a54a08e913a60e3114c362476da5479d46b265b9fde84fc5e0f"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.hpe_storage.password / d03b961e7fcc / 3

- [blindfold_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-d600c080717d7862a5c37cacc13829cabadf8ce96dc317fe31f937a5fd733d9c): complete subsection reference.

- [clear_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-3a2c72ad3544ce9916d2a3ef9753244349f809b82e929dedb3693cdef5cba1bd): complete subsection reference.

<a id="canonical-cbfb2784dc4d2278825f32b83f3b249a5007c2c46d321d8e957f0102ee18552f"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.hpe_storage.password / d03b961e7fcc / 4

- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-d600c080717d7862a5c37cacc13829cabadf8ce96dc317fe31f937a5fd733d9c)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.clear_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-3a2c72ad3544ce9916d2a3ef9753244349f809b82e929dedb3693cdef5cba1bd)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](data-sources--voltstack_site--reference--group-006.md#canonical-67db46fb3b026505815885fa5b5ca058dcefb83ebd8279d77606ead400c459f3)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-d600c080717d7862a5c37cacc13829cabadf8ce96dc317fe31f937a5fd733d9c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3f514bb3915a1fb18c572f4a26dd7888203d0c56087ff328e2b31eb525d90663"></a>

## custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info — custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.b / b83a0bfbe66a / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](data-sources--voltstack_site--reference--group-006.md#canonical-67db46fb3b026505815885fa5b5ca058dcefb83ebd8279d77606ead400c459f3)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.password](data-sources--voltstack_site--reference--group-006.md#canonical-7a31e682adc92925997c4ede6355b5c0afd73173d5cd7a468095c33970cecfdc)
- custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info

<a id="canonical-c708d0d3bdbc4f28bbb2b97534982473889cff392268c13afc1cf6672b07621f"></a>

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

<a id="canonical-d90fb3b7ef85c282119a09c643929f8a7528eadaaebb5bb7e506d9c4faf7c79e"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.b / b83a0bfbe66a / 3

<a id="canonical-7015be6064fd11f560ac29584753b2fec88bdd3111a613c5741636b75874217e"></a>

<a id="canonical-4a2c0832e1f7f4afe4182af7230a738653b5e353126c809ac406b52f217d5639"></a>

## decryption_provider property — custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.b / b83a0bfbe66a / 4

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

<a id="canonical-acab2c5b755c87da2e69a93c2a7e17b8375138ff64e2dcc358916ff25f9bce07"></a>

<a id="canonical-d93dc0ccf77865189563ffb896eeb1f6a021dd7041721b8f0449552fb7bb67ad"></a>

## location property — custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.b / b83a0bfbe66a / 5

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

<a id="canonical-23fe739f2282aa901449c88a52dc726201ec155dd564c3a8079982af46c91f79"></a>

<a id="canonical-a01d986ca1cebbd3dc0c6303e6863024135f0feb5217ca6c21bc674209933132"></a>

## store_provider property — custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.b / b83a0bfbe66a / 6

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

<a id="canonical-b99a771bcaeba66d79bb3b930bb51398411c6858e1f67a51bd12dad94df7f037"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.b / b83a0bfbe66a / 7

- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.password](data-sources--voltstack_site--reference--group-006.md#canonical-7a31e682adc92925997c4ede6355b5c0afd73173d5cd7a468095c33970cecfdc)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-3a2c72ad3544ce9916d2a3ef9753244349f809b82e929dedb3693cdef5cba1bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af56b83286ae8adaf39d10b6b5e6b72510a74eb1195b0625ed7b170a41c1d9b5"></a>

## custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.clear_secret_info — custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.c / 6b2a02f763ff / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](data-sources--voltstack_site--reference--group-006.md#canonical-67db46fb3b026505815885fa5b5ca058dcefb83ebd8279d77606ead400c459f3)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.password](data-sources--voltstack_site--reference--group-006.md#canonical-7a31e682adc92925997c4ede6355b5c0afd73173d5cd7a468095c33970cecfdc)
- custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.clear_secret_info

<a id="canonical-4fb00a13117f9817e28fa52afac6131613690cfa5619f4b45c8fcecab4a6f078"></a>

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

<a id="canonical-8e4e2b6108408d62420c1caa21094243674f83cad49a32d221b529fea776174c"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.c / 6b2a02f763ff / 3

<a id="canonical-cac299793301bf6a94074069c930f2933ea0de253463734d15254b7c3524a1c0"></a>

<a id="canonical-7aa80e54e1153e4bed58ea65e4539ba2139d51cc5a5439233226cb0ff10a856d"></a>

## provider_ref property — custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.c / 6b2a02f763ff / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-36f9c3c25f3c42ed4e05421e6318c6332bf7c0d9edf230cef164bd18134328d0"></a>

<a id="canonical-cb013910b44d3d3435f264adb9488daf70a19c35a1ee9f0d72afcbf4ef1dcc44"></a>

## url property — custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.c / 6b2a02f763ff / 5

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

<a id="canonical-379c5bc148bd74f05e2d5e63a6b60d5f4937b68772963a24f82906a548a982df"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.c / 6b2a02f763ff / 6

- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.password](data-sources--voltstack_site--reference--group-006.md#canonical-7a31e682adc92925997c4ede6355b5c0afd73173d5cd7a468095c33970cecfdc)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-44d3c0efd4276a82a6b16ea7ad524c08d0e85602b46b7a983074f352ea7fc673"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident — custom_storage_config.storage_device_list.storage_devices.netapp_trident / 053a92b75551 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident

<a id="canonical-0a5efb154196b6d29305d2c154a93b7d5eadf1d524de243f09e113eaf32fb76c"></a>

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

<a id="canonical-76934a8adfb690746215f1ed3954aaf1cae2f84fe4054b92cd48afe0e0d63194"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident / 053a92b75551 / 3

- [netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-d8d909e17b75e41c3beb4b14f6b7c6990571ac7ea4e563a7afa484ef06ae2399): complete subsection reference.

- [netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-6b08c8ff5901aca0ac1338b0d9bf20f39f31d253847573540b28426d4f9662d3): complete subsection reference.

<a id="canonical-a2afb52776e927a226cf3b757111c0603af57ae0bc47ea9a76cc9c07b10ded9f"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident / 053a92b75551 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-d8d909e17b75e41c3beb4b14f6b7c6990571ac7ea4e563a7afa484ef06ae2399)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-6b08c8ff5901aca0ac1338b0d9bf20f39f31d253847573540b28426d4f9662d3)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-d8d909e17b75e41c3beb4b14f6b7c6990571ac7ea4e563a7afa484ef06ae2399"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b614076eaa5c6722a5c5a60ba508fd4e3db89f9cc7ddd7ffd2af73edd563727d"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0b66d29058c2 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas

<a id="canonical-58ba52b7dc37f30fd9b5695f73f055b278ee5af92142a1257bf3cc3a008f7835"></a>

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

<a id="canonical-edf29d3e19a1691bb4771fb8c20880ea414053e169f749a9680c5bcb7d47e617"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0b66d29058c2 / 3

- [auto_export_cidrs](data-sources--voltstack_site--reference--group-006.md#canonical-cb5127c141485b45f92535a7f05cfb4ac06db38d7944238a69093434d1c33372): complete subsection reference.

<a id="canonical-7d0b2e1a6a7bb05d8c68d3f8779d7dc6376a58230db37354f0c0889f3964be7f"></a>

<a id="canonical-b9d1562fc2c091923d213d8bf349bac9ddd94e7e0adac8f3844f09f422ed60b8"></a>

## auto_export_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0b66d29058c2 / 4

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

<a id="canonical-583595c9e4126249c73dc19d4a3231909fc440ddacab34a91ef244b08b70e0e0"></a>

<a id="canonical-a513213a9b78ca5cce200f3e9b33c0ffecb49268a185155db568190c4370f5a5"></a>

## backend_name property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0b66d29058c2 / 5

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

<a id="canonical-ce661c8d4aae789693f9c473944b6ed408c0be0c21fdf5c6dcf3a5b167802650"></a>

<a id="canonical-57a286095c7af646399d7894934e759b2eea1de36866883763b5b916f64b2a81"></a>

## client_certificate property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0b66d29058c2 / 6

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

- [client_private_key](data-sources--voltstack_site--reference--group-006.md#canonical-cfa20b08642e6ed042a59be8509557cd71db722022dbd590ca1cd4c774f94c50): complete subsection reference.

<a id="canonical-4443967d6b07a19a2d61c78d47af437286c535b95a633224d33543b27a97a18c"></a>

<a id="canonical-bce0b8d4089a48f4a49dda97bcac5106ff58481c01977a55f48b32bd95bbf7a5"></a>

## data_lif_dns_name property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0b66d29058c2 / 7

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

<a id="canonical-72b2fd4dfe4450d9c1d152f1f8fc71f7f1f9f53413547f6ac347fe64d63e985a"></a>

<a id="canonical-b638a0fac54103e6409e411750e91e34b2ae9e3943d9793ac837205928b90024"></a>

## data_lif_ip property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0b66d29058c2 / 8

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

<a id="canonical-2f35159059c8940b750c4f086f121d272f855b2f03b01190e556831ec104a937"></a>

<a id="canonical-a693172d76a604dcd473ae4dc0698d5fbb7372c3b24d1bff0de7c0bf7a7f0417"></a>

## labels property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0b66d29058c2 / 9

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

<a id="canonical-9dace6f5f1662adf425216be5557fd4e1abeb42789735e3d8e729ad5774f7cc9"></a>

<a id="canonical-7a6fd97076aa220172767488c26acd0977682f792f8299962cb97f4c45ad7a02"></a>

## limit_aggregate_usage property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0b66d29058c2 / 10

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

<a id="canonical-ba2b4ec73b6518a850997079e2e06c559ccc810346da74f173fa4675c193f664"></a>

<a id="canonical-8923e0139fa7021fad81d0626781de0bcf773a00f99f36527b3e92547b564ee1"></a>

## limit_volume_size property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0b66d29058c2 / 11

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

<a id="canonical-3cd78325faf12e56d2154d067b7a590098f078c93a7470fdec22979e812a9666"></a>

<a id="canonical-dc8621ca70ad9d8158f84597c0c24844e47ba5221b89cd8b21d85f6a9b3061f6"></a>

## management_lif_dns_name property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0b66d29058c2 / 12

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

<a id="canonical-8a54d9910364e8f37aee0c2b1e64bd91b7243ef5a498cb7c601ce8838084947d"></a>

<a id="canonical-e3f0ee2ed3b602e27f081b5ee91e09346de396c70f31245380597ebb4d22d832"></a>

## management_lif_ip property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0b66d29058c2 / 13

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

<a id="canonical-048b90fcef787996770eab19297ed5d95beef9c2b7954d8b4a1c40b96e1eadb7"></a>

<a id="canonical-0a4b1cd318b1937bd92823e48eef6b93000beb9cfb6c0a9d90ff5d9468d48aee"></a>

## nfs_mount_options property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0b66d29058c2 / 14

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

- [password](data-sources--voltstack_site--reference--group-006.md#canonical-42d0a505936e9724ff0b5bd866c856d531805011ff33317156322a0a39631864): complete subsection reference.

<a id="canonical-a50d15b478e08d6cdace068ef403b2938618fca0708c10f1462710559ddbcceb"></a>

<a id="canonical-d85421e7d8069a02619b17cf86f616b93fe64c35fb873b55b577cb139e5f0003"></a>

## region property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0b66d29058c2 / 15

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

- [storage](data-sources--voltstack_site--reference--group-006.md#canonical-b5ccdf1ccfa879653b3ebad38eec27e8d53ab8cc0aa8a4261c57b468a5c4d151): complete subsection reference.

<a id="canonical-e41b43f56043442a2cb8b65077893fd6d1fc57ddba7e5a5ee94b7a101e604288"></a>

<a id="canonical-7b57594716ba97b7eb0df418c7696dafcda73696a69601aba4477e92e2eebfda"></a>

## storage_driver_name property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0b66d29058c2 / 16

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

<a id="canonical-73bc77ad7bff3a8cf5aededec5b36dd57595163fcb7734c99a0e121ae4440541"></a>

<a id="canonical-f4e6f9fbbb8d6bd90665a31cc0ee3b25221b8206d35bc915c2f96d3e6a5abd0d"></a>

## storage_prefix property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0b66d29058c2 / 17

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

<a id="canonical-143a0ffb50b01973812eb90dc5fd60475e3d400eeff2d266aa806eef3422d331"></a>

<a id="canonical-f43ce1c1b3e771d36d85ba3a675e4786e61453bc8854389eb6aaf7818660213a"></a>

## svm property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0b66d29058c2 / 18

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

<a id="canonical-487be59526097b0ccbeaf7c9abe0758c4d1dc2d155757c219bb81cbd4eb8685c"></a>

<a id="canonical-36e229b3210b9861ab4779b355f9fec61daf57676ee47bd04f17990d88540549"></a>

## trusted_ca_certificate property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0b66d29058c2 / 19

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

<a id="canonical-8d2195538dc4d9541f6a96d5a59052bdc76c9993ac5e95dfacf7c32c37530f36"></a>

<a id="canonical-394a1d304d792c6c6755bc828646f4929826aa49231b108b218338b9c6d92b26"></a>

## username property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0b66d29058c2 / 20

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

- [volume_defaults](data-sources--voltstack_site--reference--group-006.md#canonical-f2d05aa6b0e3ffbf67166d4e263396982099aefd635e60e6a71286fb46a6c93b): complete subsection reference.

<a id="canonical-83ed3cbf8454204f593092469148a47932e3b928f6c9ae6298574dad68e6e89e"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0b66d29058c2 / 21

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs](data-sources--voltstack_site--reference--group-006.md#canonical-cb5127c141485b45f92535a7f05cfb4ac06db38d7944238a69093434d1c33372)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](data-sources--voltstack_site--reference--group-006.md#canonical-cfa20b08642e6ed042a59be8509557cd71db722022dbd590ca1cd4c774f94c50)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](data-sources--voltstack_site--reference--group-006.md#canonical-42d0a505936e9724ff0b5bd866c856d531805011ff33317156322a0a39631864)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](data-sources--voltstack_site--reference--group-006.md#canonical-b5ccdf1ccfa879653b3ebad38eec27e8d53ab8cc0aa8a4261c57b468a5c4d151)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults](data-sources--voltstack_site--reference--group-006.md#canonical-f2d05aa6b0e3ffbf67166d4e263396982099aefd635e60e6a71286fb46a6c93b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-cb5127c141485b45f92535a7f05cfb4ac06db38d7944238a69093434d1c33372"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7582fa80b1388b592cce9d8381701a7b9d0e2ce5712ef200febb22d366a75e7"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / f57ffe746c16 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-d8d909e17b75e41c3beb4b14f6b7c6990571ac7ea4e563a7afa484ef06ae2399)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs

<a id="canonical-fcc3bcce2d34086f6616e9db6b010e56ce87650345b84840bd8e615567184b89"></a>

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

<a id="canonical-182791316da196fab46364aacb6544d6669edac42bef08ae06df7332756ae09b"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / f57ffe746c16 / 3

<a id="canonical-f8278358e93e1de64fc3f980a2a9d54cc532186e7adcd7ad539c3ae08ed2268c"></a>

<a id="canonical-b5cc92a538acf1590eccb0f78d48969e9a15660eca7019cec48e887d20c8fd55"></a>

## prefixes property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / f57ffe746c16 / 4

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

<a id="canonical-d3380846e1297bc2e518ef4e3f7b3036af827477d4bdc8489c5d04fe25f66104"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / f57ffe746c16 / 5

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-d8d909e17b75e41c3beb4b14f6b7c6990571ac7ea4e563a7afa484ef06ae2399)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-cfa20b08642e6ed042a59be8509557cd71db722022dbd590ca1cd4c774f94c50"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae33131fb4fb9c8900aac785438e310c6c8df44fe237acefe87eb483e97d8ccf"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e161ef7f1d83 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-d8d909e17b75e41c3beb4b14f6b7c6990571ac7ea4e563a7afa484ef06ae2399)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key

<a id="canonical-560cc67d2b904ae0bc0a73f12fb7c0fb7b3bf30f1d4e6ad1dbc0aa880a0b0fae"></a>

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

<a id="canonical-6c7e734cc24b9dc32619ab228ce1ef1f1b4bfb8ad90fe1e8d6c8570afdeed165"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e161ef7f1d83 / 3

- [blindfold_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-d765de613dc814f5b5b05404c5c2d216c53c137461a3922dd85c5fa41c8aaadc): complete subsection reference.

- [clear_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-a6c7898854f7e3619a5fb88abf66d935984ca94daab3d4d01b63b474b853c3d8): complete subsection reference.

<a id="canonical-8d5dddeb0b060806b5a1368f699dd1c006ddd0b7372e1841342dad5056d6dc67"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e161ef7f1d83 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-d765de613dc814f5b5b05404c5c2d216c53c137461a3922dd85c5fa41c8aaadc)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-a6c7898854f7e3619a5fb88abf66d935984ca94daab3d4d01b63b474b853c3d8)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-d8d909e17b75e41c3beb4b14f6b7c6990571ac7ea4e563a7afa484ef06ae2399)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-d765de613dc814f5b5b05404c5c2d216c53c137461a3922dd85c5fa41c8aaadc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7c4c8a36e06429fbdf94749913982b427caf607b96108ecf4e99e0f8818baa3b"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 27447872a9b5 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-d8d909e17b75e41c3beb4b14f6b7c6990571ac7ea4e563a7afa484ef06ae2399)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](data-sources--voltstack_site--reference--group-006.md#canonical-cfa20b08642e6ed042a59be8509557cd71db722022dbd590ca1cd4c774f94c50)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info

<a id="canonical-f4674cf364cd5abc7a768a0bd579a11a7f0e3590e5a9f4619fde173a6108d4af"></a>

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

<a id="canonical-5ecfd015c7b142f8dffa45e419bc0a313c91141b2c43cc33ccb86bf4a3cd276a"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 27447872a9b5 / 3

<a id="canonical-7289614686be9fa05d4a3871ec413b6db53b5df15da120bd0293e4f0073d23cd"></a>

<a id="canonical-cc41844b235b6a169a65405b8cdd2bf787032f9ae25f6a55a24122e187797a72"></a>

## decryption_provider property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 27447872a9b5 / 4

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

<a id="canonical-37bbeafec0d4e6413884510afde733aeb05eeead2f95c9a3752d32c0eeb9f49c"></a>

<a id="canonical-222ba4a4933bde9410f7182019e9055fd32d72d5a6abd6576d10ebe1a3089d0a"></a>

## location property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 27447872a9b5 / 5

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

<a id="canonical-9a359106e4a6e77f6aab06c00798eeef1f3a42b8bad01a6dbef71ef087877a59"></a>

<a id="canonical-ea44b1322f92b60dc350849cf98a5b0a648bf7598de55a7d1febd7177086ea09"></a>

## store_provider property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 27447872a9b5 / 6

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

<a id="canonical-ab7c5fac59537e0c7c0ee9d126e62986999dc22611a9174e648bf8e5406b12e5"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 27447872a9b5 / 7

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](data-sources--voltstack_site--reference--group-006.md#canonical-cfa20b08642e6ed042a59be8509557cd71db722022dbd590ca1cd4c774f94c50)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-a6c7898854f7e3619a5fb88abf66d935984ca94daab3d4d01b63b474b853c3d8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd2aae6d1ec6445655b18c3fec430cc4531af53053b9ec9798b18edea5fae450"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 37a357d0c33c / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-d8d909e17b75e41c3beb4b14f6b7c6990571ac7ea4e563a7afa484ef06ae2399)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](data-sources--voltstack_site--reference--group-006.md#canonical-cfa20b08642e6ed042a59be8509557cd71db722022dbd590ca1cd4c774f94c50)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info

<a id="canonical-055947db6b90b210a622294f1e7ef0be3a0aa52506867a578ec9bd1a8da3afb2"></a>

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

<a id="canonical-bd08ef9fc676cfc8ccf05fc272b91848023e06c66c6fed92328db7307c675a05"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 37a357d0c33c / 3

<a id="canonical-34406f54e87b1267a36a60818fe8f67663c2103adbc4bd8b319b49310875567a"></a>

<a id="canonical-28b333559ced49a2ee97b84bd696161ee1ff9403be91765230091436aff0980c"></a>

## provider_ref property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 37a357d0c33c / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-d8cb9c95a3cf40f4b13d26781bae062a97e080ba0e8401cc743367fd733c3cc0"></a>

<a id="canonical-fd263bf4f3e9d3a63cc76d4c8a3ba16e0c7654c50175264457b4e180e54bb1f5"></a>

## url property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 37a357d0c33c / 5

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

<a id="canonical-8196c0d854f61e05eb04dd83fceb062e0aea28055a21c72f3e6a24b368d57059"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 37a357d0c33c / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](data-sources--voltstack_site--reference--group-006.md#canonical-cfa20b08642e6ed042a59be8509557cd71db722022dbd590ca1cd4c774f94c50)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-42d0a505936e9724ff0b5bd866c856d531805011ff33317156322a0a39631864"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0dc5d13357608cb21dc1ebcd566f9ddd98b24ffdcfe3ad59231e57e48d3a1fe4"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 7d4ef18b99cc / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-d8d909e17b75e41c3beb4b14f6b7c6990571ac7ea4e563a7afa484ef06ae2399)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password

<a id="canonical-8c7107bbccb3d4ffa3bb8dd4b679a1a4e9a63150d2c1a33cfe35bc614b93b3d7"></a>

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

<a id="canonical-d84e73685d962fa35b73b59ac65f4fbdcd19c9d71c5e118afaf54c910ee58db9"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 7d4ef18b99cc / 3

- [blindfold_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-c9ede6f944390b30fd1797f88288de9d0714846f8652efd686405c483b589f3f): complete subsection reference.

- [clear_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-bfae9e73c70e09ba5e0645a32d8fa5b577b1655e81ccae2a57004678d5c4dfba): complete subsection reference.

<a id="canonical-c8a457fbc714400cffc2d3df7f7e15a44fb718ae6d846c61170c560d4f881541"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 7d4ef18b99cc / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-c9ede6f944390b30fd1797f88288de9d0714846f8652efd686405c483b589f3f)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-bfae9e73c70e09ba5e0645a32d8fa5b577b1655e81ccae2a57004678d5c4dfba)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-d8d909e17b75e41c3beb4b14f6b7c6990571ac7ea4e563a7afa484ef06ae2399)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-c9ede6f944390b30fd1797f88288de9d0714846f8652efd686405c483b589f3f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-755d469551ba0e20c71f1bb73b0b96ca7077eadf7c8158bea19649f00804410d"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 1bc6f9290feb / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-d8d909e17b75e41c3beb4b14f6b7c6990571ac7ea4e563a7afa484ef06ae2399)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](data-sources--voltstack_site--reference--group-006.md#canonical-42d0a505936e9724ff0b5bd866c856d531805011ff33317156322a0a39631864)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info

<a id="canonical-c5cd4a533cdaf182d1168a5a825a70656078c2c236f249a16a1c4032d8d8240d"></a>

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

<a id="canonical-2e838cb13046608d2f08e78fd8ba7a85c7aaad0d4cfa27ccc9a78bef7ad08a79"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 1bc6f9290feb / 3

<a id="canonical-8cf9a76da5949c0454ec80481ca2b24fda1e396db4288b179cdc9f9b90e480c3"></a>

<a id="canonical-babb98d5af4c22cad2f45c895afbe836dd94b521fa1bdbe4aa2f99420ac73d24"></a>

## decryption_provider property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 1bc6f9290feb / 4

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

<a id="canonical-7f7e8bf0552e01617216629297f7f73b3d1950c451672c5d8ca0fd2a65126575"></a>

<a id="canonical-22b72ba35d4cf9c9483301edab483fe9208352fa54b84425c97a587a5ff10195"></a>

## location property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 1bc6f9290feb / 5

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

<a id="canonical-d359db14abe2e1dda681effc41afa3ef16fc0a9446856bb181c8ee81781f1a01"></a>

<a id="canonical-4de774881d7552b22944f0204e745cbd8edeb77ee03bd96ec25b173699a42780"></a>

## store_provider property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 1bc6f9290feb / 6

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

<a id="canonical-0b8457d1df8ce03b0f7f8b06888cd4eeeca5ae0fa3c6d4c030decdee1a086c74"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 1bc6f9290feb / 7

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](data-sources--voltstack_site--reference--group-006.md#canonical-42d0a505936e9724ff0b5bd866c856d531805011ff33317156322a0a39631864)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-bfae9e73c70e09ba5e0645a32d8fa5b577b1655e81ccae2a57004678d5c4dfba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dfc7a96b33d606ec9732e2e4416b0f881798cd313be4cd8eaaa0bbd1c23eef79"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / c9d7cb700d70 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-d8d909e17b75e41c3beb4b14f6b7c6990571ac7ea4e563a7afa484ef06ae2399)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](data-sources--voltstack_site--reference--group-006.md#canonical-42d0a505936e9724ff0b5bd866c856d531805011ff33317156322a0a39631864)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info

<a id="canonical-220437724616cbf0a79dec320cbc7e7a5bcd4fa36774aceef2cb417522b7f79f"></a>

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

<a id="canonical-1e7baaef8d1306b1058b3bc7552361690463af57717e993dbcc33eb84a4bbb73"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / c9d7cb700d70 / 3

<a id="canonical-4cb4e0fac1b1b1f5c7c8e4fc40bd78592842683ea02918ec7cc0c4e79109e73b"></a>

<a id="canonical-cf99c7c19e144f2467792edc684a6008c7f5fc03235fae194a6c9bc090e5031a"></a>

## provider_ref property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / c9d7cb700d70 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-366fa79fc8622d8a91b99b40fcaaadbd463da7b00acaa0c5b91b66c698c8925f"></a>

<a id="canonical-882b30601c2a00f7f4ae34bc142488399c9ebb6a1f3b431d7d190d3b3dd065bc"></a>

## url property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / c9d7cb700d70 / 5

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

<a id="canonical-0c59870936aff92017be4054f1b0d1b8a2d75a3dd1832ba76761b95c438c8069"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / c9d7cb700d70 / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](data-sources--voltstack_site--reference--group-006.md#canonical-42d0a505936e9724ff0b5bd866c856d531805011ff33317156322a0a39631864)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-b5ccdf1ccfa879653b3ebad38eec27e8d53ab8cc0aa8a4261c57b468a5c4d151"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-879123b1f15028d0e13f0fd72a201e774c77011f222b8be54c260614074b8788"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / c9f7c0e9345d / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-d8d909e17b75e41c3beb4b14f6b7c6990571ac7ea4e563a7afa484ef06ae2399)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage

<a id="canonical-aeb544dcf4b9a958e06dbf8ce82aa77c62bca793fe566aee25315904ddc2fe59"></a>

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

<a id="canonical-6132d01631084498764865a7fc2b2feb1b71cfcc2f086fd8852d6ba8108d4199"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / c9f7c0e9345d / 3

<a id="canonical-c322c669653a583ceefbfc348d0a11a99266ef5265bd5890d463b72b70c59ba3"></a>

<a id="canonical-2b2234bcad38a38039ec6fb73addaadb06e9d0e6fa2b5be158bbc1a7d329308e"></a>

## labels property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / c9f7c0e9345d / 4

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

- [volume_defaults](data-sources--voltstack_site--reference--group-006.md#canonical-a75c821b2aa7dc986c7a69f1ae01b9ab2ad769a59112659b34c5240bb32dbdb4): complete subsection reference.

<a id="canonical-1ebd97a739537304528920ae66d629a4961c9ee5c5d9dce92c9d81274dcc6444"></a>

<a id="canonical-56b766242e1255bfc906c26f44d5bba690943e90189acfc374f9b09f528683fd"></a>

## zone property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / c9f7c0e9345d / 5

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

<a id="canonical-8ad61eda3fb2dc25b502fd50c9583f9dfdd24d45c8063482c372296f0e22cfb9"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / c9f7c0e9345d / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](data-sources--voltstack_site--reference--group-006.md#canonical-a75c821b2aa7dc986c7a69f1ae01b9ab2ad769a59112659b34c5240bb32dbdb4)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-d8d909e17b75e41c3beb4b14f6b7c6990571ac7ea4e563a7afa484ef06ae2399)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-a75c821b2aa7dc986c7a69f1ae01b9ab2ad769a59112659b34c5240bb32dbdb4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f89ccde14d009e46781a138de2dea6dd6f29b7c5d4c22443306d647d67f418f"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / ca820d5691f6 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-d8d909e17b75e41c3beb4b14f6b7c6990571ac7ea4e563a7afa484ef06ae2399)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](data-sources--voltstack_site--reference--group-006.md#canonical-b5ccdf1ccfa879653b3ebad38eec27e8d53ab8cc0aa8a4261c57b468a5c4d151)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults

<a id="canonical-a26bb940fb6f87ef593272b925fdad3b5aa008112e1f36247f2a1af0e57e2f64"></a>

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

<a id="canonical-baa3265b6c572d59c2383aa6125a48567bcb6f4acadff2c27ed6e2d237d8c87d"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / ca820d5691f6 / 3

<a id="canonical-6f88741711f6441ab1f67886e421663ce88939bc3824c3ccaa2b7c68233fa064"></a>

<a id="canonical-19cefd7826e56b01ccec8e1d1f188a58ef41232ac60a1c867c2683331cf5e4f8"></a>

## adaptive_qos_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / ca820d5691f6 / 4

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

<a id="canonical-6182d1e28194cf2392aa05ca9f8f91dcb1034f65a775f991dc3de9554370c26e"></a>

<a id="canonical-8ccc30f148ad8b850aa9c9f44f79630c722768fc5c81e2c4fc2e5e6145fa9be7"></a>

## encryption property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / ca820d5691f6 / 5

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

<a id="canonical-f065626626b9baee92134a212e98b1216bb63b99a933b9aa1225af16b4d649e3"></a>

<a id="canonical-e9728647dd06e1086478c48fd99fb2021ef656738562798315c891f3d6b4121f"></a>

## export_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / ca820d5691f6 / 6

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

- [no_qos](data-sources--voltstack_site--reference--group-006.md#canonical-05739f7fef5926c30665d3a6c712f12251f9e9f474a1fc5500416dabdee31366): complete subsection reference.

<a id="canonical-a42585c41a385959d7a6668f2a549c913cca4e8bfecdd89f8d4399b266f32ad6"></a>

<a id="canonical-7d0af93140a81c63499ea97a8d4e8e94d303d1f88f759050ecfc33aae6691447"></a>

## qos_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / ca820d5691f6 / 7

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

<a id="canonical-cf0c93722f164c650c5a401af3d593a1379839b0ec757b42a4dc0959bce6e95f"></a>

<a id="canonical-029f0dfc80b43ec178f9734be3f24c11f89c9913984ca4127384fe6317b63c0b"></a>

## security_style property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / ca820d5691f6 / 8

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

<a id="canonical-64bddaa89143265555af970cb01cee0e163ccb1e9e03bc5def77731024066f1a"></a>

<a id="canonical-9262a08a48b9b7a0dbe5acf5c17e8fe850140026d14dd5f9aa5177f0706e9ccd"></a>

## snapshot_dir property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / ca820d5691f6 / 9

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

<a id="canonical-3585cbb68c484c1c3d4b46cfae9a23108828e2391a24828e21024751fff2e5c4"></a>

<a id="canonical-ba123bd09113468b49793bc442a4e958afd996ef5d1f08976f070f87baf7befe"></a>

## snapshot_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / ca820d5691f6 / 10

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

<a id="canonical-3516149b6bfb339cb06d8bd0ca060244dfe0b50daeb974cb44aa9a93fab9a86d"></a>

<a id="canonical-89fae7de0972ccce1cbfb9b9a8e45a50feeeef6052338ac09c5bb84491cf27ce"></a>

## snapshot_reserve property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / ca820d5691f6 / 11

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

<a id="canonical-5c00e9b6c81a66d1164735e62a14e8aea6a6d070992dea5469cd929bf9eae125"></a>

<a id="canonical-0da7340727b31a9ab3f3f45fe2493ea2cb90208fe734f66e1342678279a98e50"></a>

## space_reserve property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / ca820d5691f6 / 12

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

<a id="canonical-646d1414c8b3e52d13481476017d01621dc17ccd5e0124f96d175fef2b21f584"></a>

<a id="canonical-1e031092299f1059ec2dd129211c86356403e4f76efceca467a63f92d6009fb8"></a>

## split_on_clone property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / ca820d5691f6 / 13

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

<a id="canonical-3590c587e795bc53486e6e01eb751110a33eca379c8a7d665ffe0bf583e2fb1d"></a>

<a id="canonical-cd20c819cd4289429f819ca6595e49124c9decc6d4e9c53ca153a95f5e6a4ed8"></a>

## tiering_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / ca820d5691f6 / 14

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

<a id="canonical-577ebe613dd8c9dd729d6d5530719814bea72529ced2b916d3b8ca98254aa067"></a>

<a id="canonical-1d548c913be0c2dd7e505f29e0ed5911fff689d0177e969160d1af45393af1af"></a>

## unix_permissions property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / ca820d5691f6 / 15

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

<a id="canonical-9ee8199d2b7365a0264c77e6ec05b53feb8b4a5c83da3d9f279e8599ea28a661"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / ca820d5691f6 / 16

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos](data-sources--voltstack_site--reference--group-006.md#canonical-05739f7fef5926c30665d3a6c712f12251f9e9f474a1fc5500416dabdee31366)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](data-sources--voltstack_site--reference--group-006.md#canonical-b5ccdf1ccfa879653b3ebad38eec27e8d53ab8cc0aa8a4261c57b468a5c4d151)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-05739f7fef5926c30665d3a6c712f12251f9e9f474a1fc5500416dabdee31366"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d7a3e1584c6aaa05389294ab64a379941d06fb3fcca8a434f7302908a1ba8ffd"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / fe376bb23c00 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-d8d909e17b75e41c3beb4b14f6b7c6990571ac7ea4e563a7afa484ef06ae2399)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](data-sources--voltstack_site--reference--group-006.md#canonical-b5ccdf1ccfa879653b3ebad38eec27e8d53ab8cc0aa8a4261c57b468a5c4d151)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](data-sources--voltstack_site--reference--group-006.md#canonical-a75c821b2aa7dc986c7a69f1ae01b9ab2ad769a59112659b34c5240bb32dbdb4)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos

<a id="canonical-5a0bcc6dfe30e6abc4c424f2982d653f3ecc1c0f38a18513cc8129ded2cf0b7a"></a>

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

<a id="canonical-0a62e28b6d1e55bf67c6a179db93b1e3da1768bef14e5ee255c543bd06fc9cc1"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / fe376bb23c00 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8124e8de91ac09ae925abab40e4114aa3f4d2a802fcfc81fb67aa897e2cbcb09"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / fe376bb23c00 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](data-sources--voltstack_site--reference--group-006.md#canonical-a75c821b2aa7dc986c7a69f1ae01b9ab2ad769a59112659b34c5240bb32dbdb4)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-f2d05aa6b0e3ffbf67166d4e263396982099aefd635e60e6a71286fb46a6c93b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68c55070897d5491a29d1e40b1af2259b62e8df89a697d5784f01229f7b30219"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 93cf1005f8ce / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-d8d909e17b75e41c3beb4b14f6b7c6990571ac7ea4e563a7afa484ef06ae2399)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults

<a id="canonical-870d833688e08d6d46cd9cf3c04e886a222746776f52a6d8e537869e15e62c37"></a>

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

<a id="canonical-f4e3a97962cca356380c3ce52937c236a08613786a679bfccfdd99766029f814"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 93cf1005f8ce / 3

<a id="canonical-3e1bb197fdf73c55b630303393291776258c74e9865e594c25f0271b11fac371"></a>

<a id="canonical-56730f275b09dfdde5f6bb4e6879ad23a070d190fd60353122cf2bcb61ff1654"></a>

## adaptive_qos_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 93cf1005f8ce / 4

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

<a id="canonical-bc1b4f653fa3fb5d27abd9acc5c56b8548df22e45dedfd210f2ce72bb9bfde66"></a>

<a id="canonical-c8ba255410fec360cbaa96d389756b4695bf4953724b433bf6554b5376735b1b"></a>

## encryption property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 93cf1005f8ce / 5

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

<a id="canonical-bc4f0abeda44f9fc722eca8e0355aef04757c43634eb301797319ca8188727eb"></a>

<a id="canonical-e8d3b0593fa8bb5b8151ad1724e6120952ef3398fdd0ba86d09e70ccf79a0302"></a>

## export_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 93cf1005f8ce / 6

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

- [no_qos](data-sources--voltstack_site--reference--group-006.md#canonical-7b3e6a4735f23c31440abb7ad15f26304c214344ddd6f03cafe2bb5fcf3ecfbe): complete subsection reference.

<a id="canonical-2da2a2061c244a3c353fafc5de09fd75413ce5847dcb6eb272840bb62e564b05"></a>

<a id="canonical-e701431522fc14d95acf78bb7c77b47ff1e32980471b9a55fa8b13a41e393f28"></a>

## qos_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 93cf1005f8ce / 7

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

<a id="canonical-12bdd348ff320ef2753e7e5e497e803abf4d62e3351714344e9e99869a62d792"></a>

<a id="canonical-cf2cd6459f53eca152e0325773a5978a237220e928bd48c7b63f14610f44799d"></a>

## security_style property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 93cf1005f8ce / 8

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

<a id="canonical-77188c2b22fc492ffd98497a0bfb95be383fd46dc95e8464601d041b7c918abe"></a>

<a id="canonical-98c6f8c61b5d2a2375e77265744032823bf4e0192bf241d02ab9ca8533805734"></a>

## snapshot_dir property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 93cf1005f8ce / 9

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

<a id="canonical-3812ed402441486f951cb0e82879cd399d0e28856365c31a235df9f845bae20a"></a>

<a id="canonical-e2be62f0e4fca1222518b81b1278177610f387df58dd8c2d19343e92eb10c299"></a>

## snapshot_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 93cf1005f8ce / 10

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

<a id="canonical-ba107895e7c857453bb4afe3c38b3a1ff39f06e0028107a0e7a06a50ccd42766"></a>

<a id="canonical-e8d4f16cbfcec7f0b97809147a6faed8e8aeb53aa85abf1bca225433a8a35db8"></a>

## snapshot_reserve property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 93cf1005f8ce / 11

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

<a id="canonical-d6a1b1917a2b893578e2be930f1b55926c8ef28fbcd7df4c1cda1879b94e34f4"></a>

<a id="canonical-c53411e946cdf8eebdab2f1b51b0e3b684f62f216cb9cb08cf6352d7a0a73159"></a>

## space_reserve property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 93cf1005f8ce / 12

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

<a id="canonical-0fe4bfba4e46a00152bcf0b20b322e0c87288fe212dda7622f2d500f98c2c2e4"></a>

<a id="canonical-09e5f508771526c5c703c30d06faa8f9129bcc1449e14750219b6a2941c59fc7"></a>

## split_on_clone property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 93cf1005f8ce / 13

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

<a id="canonical-c69cd0afb9792dfecf471f8787462088a40a95f5b12494d5fd43c6d82a481f79"></a>

<a id="canonical-d1f8a71548b3567308e3f096f10cfc3dd10ee8bb23dbd5ce0b37cb670e3428df"></a>

## tiering_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 93cf1005f8ce / 14

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

<a id="canonical-5f41337b9f071aa3fde331d59f2ed57b763c6c1c275d46fe09bdf9ea75eda24d"></a>

<a id="canonical-929143a5f5b05812c296cf4705fa5a22dfa0d831faf5b77c95443ac458787587"></a>

## unix_permissions property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 93cf1005f8ce / 15

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

<a id="canonical-8c024289635ff2817d753f763a8515e2992b1156aee8b788055906b53fb38506"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 93cf1005f8ce / 16

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos](data-sources--voltstack_site--reference--group-006.md#canonical-7b3e6a4735f23c31440abb7ad15f26304c214344ddd6f03cafe2bb5fcf3ecfbe)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-d8d909e17b75e41c3beb4b14f6b7c6990571ac7ea4e563a7afa484ef06ae2399)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-7b3e6a4735f23c31440abb7ad15f26304c214344ddd6f03cafe2bb5fcf3ecfbe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-74f7daa63392b6ebc7cdaf3735943930f3c2acf8d1f7bd17224921a1ca3273ae"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 8b9dc16bff87 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-d8d909e17b75e41c3beb4b14f6b7c6990571ac7ea4e563a7afa484ef06ae2399)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults](data-sources--voltstack_site--reference--group-006.md#canonical-f2d05aa6b0e3ffbf67166d4e263396982099aefd635e60e6a71286fb46a6c93b)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos

<a id="canonical-e71cab4cd33d3f8f0d8577b653e778c71d2707c504611be05eed8a867837f5e9"></a>

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

<a id="canonical-278f9f4a294281686acb5fa45601da6517461ae50bff577ce24f65baf1a4550c"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 8b9dc16bff87 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3133ff11de4f02e8f0c5d18566bef89182a1e6f0a91509ef4030c18c2bce59de"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 8b9dc16bff87 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults](data-sources--voltstack_site--reference--group-006.md#canonical-f2d05aa6b0e3ffbf67166d4e263396982099aefd635e60e6a71286fb46a6c93b)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-6b08c8ff5901aca0ac1338b0d9bf20f39f31d253847573540b28426d4f9662d3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b706762a7130464e5c2a5b92af5a31ca545f1ae526833b2467530b51573aead2"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 787eab04a350 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san

<a id="canonical-ca3f06e01e158d65665aa6f0233de44a0a5c7649500fad33ba59ca17f948ed8c"></a>

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

<a id="canonical-7801ee0690356e7a98f489ce8909624c96cca33a1d456c8d9e6e60af9f353186"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 787eab04a350 / 3

<a id="canonical-8d60623755b81e1fdb24953165a6a45579e086703e8d98cda754a5969dc4012b"></a>

<a id="canonical-a078ac2619ac6bdcce1ffe526bad0521f828c52d53282d37227c2d7b64dfa0dd"></a>

## client_certificate property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 787eab04a350 / 4

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

- [client_private_key](data-sources--voltstack_site--reference--group-006.md#canonical-0da7819a82be7651c383737da6af25d0b7c8ffda8eb6df472902d043942fd968): complete subsection reference.

<a id="canonical-78bcd8f0ac5b82e70d157226bfa22714014728f49b936c377e7d44ff1c64657b"></a>

<a id="canonical-3b828b80ae60ef3b5ec69be769f37a96bc9655825ff2c995b3258f6f7bdb1c93"></a>

## data_lif_dns_name property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 787eab04a350 / 5

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

<a id="canonical-07ddcaab720afa669e29beb84cc5e63313f3613f1a5c913b262d7d6bfe9fa9ab"></a>

<a id="canonical-3110db573c418294a5321fae09b507068c51104879b856f59d9ab68d5ea975b4"></a>

## data_lif_ip property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 787eab04a350 / 6

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

<a id="canonical-b80ff9856d65d793d694720dd1e5606a8ade0e41929c69d11c87a83aae10f163"></a>

<a id="canonical-a578b5cfa2ef8449ebbcd1e2782c5a17379e3d57fd8c9a668a353ab83d6e7f9b"></a>

## igroup_name property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 787eab04a350 / 7

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

<a id="canonical-4ba38c6f68bd2a8393e654d36c0283984da4b90cd8de1d290d69149636b3c8d8"></a>

<a id="canonical-a05c5336b28d312257f4032c6587037ffa6ce9151797f7927e5c1e450bfa3f73"></a>

## labels property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 787eab04a350 / 8

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

<a id="canonical-c6a4137dae3f6857523327ebf203b0190628baf44eb6978ec9389233ca094147"></a>

<a id="canonical-39dbd39ebcbe15b465c013fb51716c7b0cfc74cea993bddde0b80c54c7f859d1"></a>

## limit_aggregate_usage property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 787eab04a350 / 9

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

<a id="canonical-a4cb939bdc9b62e6a63b8a56dfc5bff56e17c7789b76a023e0a271dfbfd8535b"></a>

<a id="canonical-e77d9ac93346fefe3fd5078e79bbbe5954acbb555268f4e6153eb2f7cccbef05"></a>

## limit_volume_size property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 787eab04a350 / 10

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

<a id="canonical-86803c1eae5289e82798f27f5af1a09bfd5bb511ba78509662cdfc5bc3894b1d"></a>

<a id="canonical-ac472403b1e6fc5c52a9f267625586c90ee18faf9055a1e5c55d02b96008a736"></a>

## management_lif_dns_name property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 787eab04a350 / 11

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

<a id="canonical-6a5be9fc0742b20781f30c977cb32eac6aff6b35a3fdb35e2870d4cbe14351f1"></a>

<a id="canonical-0a8e0159214043fdbc00500f4e7d32c6b408523939d8a6b70333ac637f36bc30"></a>

## management_lif_ip property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 787eab04a350 / 12

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

- [no_chap](data-sources--voltstack_site--reference--group-006.md#canonical-9b20b360bfa0a6645821cdc637798da2afdad5ed9aafcd517616fb6f741dbffd): complete subsection reference.

- [password](data-sources--voltstack_site--reference--group-006.md#canonical-82b1f280c58cbeae84539b6007e73c78a0128cff1bf2f26d672f55fac1bb46f5): complete subsection reference.

<a id="canonical-4e0ff5c7e4b025df3dc1ebcbd38f659746d122dffca6a0d85d5a59dccb2948ba"></a>

<a id="canonical-8371d65186abbc42dd4b7448748dd6152eb58d598facc791a2781d22d8de2e34"></a>

## region property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 787eab04a350 / 13

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

- [storage](data-sources--voltstack_site--reference--group-006.md#canonical-79cc1e76076c7cf9601cbbbf98de5cfe2f2338ac110641094f165aa4aa84d2ee): complete subsection reference.

<a id="canonical-829c084032417885c5300c8716ecda6874f765c32f39ba9a49347cc22ea0eb05"></a>

<a id="canonical-78fb54418344d4604e0ec55c4496f268da24d22bb3edf288c9205700ae2f30d1"></a>

## storage_driver_name property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 787eab04a350 / 14

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

<a id="canonical-052c66f99a63d235b3404253391b5af99f70751aaf500adbbfc3f02a216a3498"></a>

<a id="canonical-b0af14bf33d26fcd640196a5db6470901aa2ea16c6f254ddf3f9d24453493448"></a>

## storage_prefix property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 787eab04a350 / 15

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

<a id="canonical-a938b7e2ae0e8d7b2381dc6f03cd6e0725d932dee8de5d02946581166fc66f1d"></a>

<a id="canonical-7b463a23479b4124ede3b1c0de494cede6102d5cc7aff61ca85e174113180efc"></a>

## svm property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 787eab04a350 / 16

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

<a id="canonical-82d41442cf432b44e64880b46d9a289923f0b6e57087b9224221a7802d541a5e"></a>

<a id="canonical-19c0d7fd073ed4d07fb7c55642f2740905abd7cdca72672c1418d3c213852fc5"></a>

## trusted_ca_certificate property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 787eab04a350 / 17

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

- [use_chap](data-sources--voltstack_site--reference--group-007.md#canonical-ebe00d06a5ec80e287e74893606d51318c5c59d010b4482d51f10ac95faf3fe9): complete subsection reference.

<a id="canonical-c338447796166f6df4f32a9e089594fff4cbdab65e96dfd192f8c484bc9e8c40"></a>

<a id="canonical-e774086f027c0c4b0e59d784d2bea019615e5cd0c0b14ecf8959547b4df6cca7"></a>

## username property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 787eab04a350 / 18

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

- [volume_defaults](data-sources--voltstack_site--reference--group-007.md#canonical-049d958c8b1328a10ee8c36da36fdbf3fd1cddf195cc903283f0846010b75345): complete subsection reference.

<a id="canonical-cde5449a600dfb051fa3564b827c1182dc679feeaf6af06b6c79974bff1f001d"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 787eab04a350 / 19

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](data-sources--voltstack_site--reference--group-006.md#canonical-0da7819a82be7651c383737da6af25d0b7c8ffda8eb6df472902d043942fd968)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap](data-sources--voltstack_site--reference--group-006.md#canonical-9b20b360bfa0a6645821cdc637798da2afdad5ed9aafcd517616fb6f741dbffd)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](data-sources--voltstack_site--reference--group-006.md#canonical-82b1f280c58cbeae84539b6007e73c78a0128cff1bf2f26d672f55fac1bb46f5)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](data-sources--voltstack_site--reference--group-006.md#canonical-79cc1e76076c7cf9601cbbbf98de5cfe2f2338ac110641094f165aa4aa84d2ee)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--voltstack_site--reference--group-007.md#canonical-ebe00d06a5ec80e287e74893606d51318c5c59d010b4482d51f10ac95faf3fe9)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults](data-sources--voltstack_site--reference--group-007.md#canonical-049d958c8b1328a10ee8c36da36fdbf3fd1cddf195cc903283f0846010b75345)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-0da7819a82be7651c383737da6af25d0b7c8ffda8eb6df472902d043942fd968"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f458edd98db8b38b4bfa25db9026563d37df6bbd66fff01beca587103f6898c"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 6fef44619f31 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-6b08c8ff5901aca0ac1338b0d9bf20f39f31d253847573540b28426d4f9662d3)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key

<a id="canonical-2ace13b41d7efd3c8dfcd6bfb18ec3de8786c5ded3bb454be812a28fe778bde4"></a>

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

<a id="canonical-e07ab6b7e62e8deaa40f0d00d9ad320037cbdff1691daa42e2eef967433fa893"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 6fef44619f31 / 3

- [blindfold_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-4bed3c46b74e6cbbab48afae17190ebf57c1a5b2c80e58e9d773fc8a20813360): complete subsection reference.

- [clear_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-2bf3070a066ff600450035b4d4948eefd5d62bcdf43299bfa689b39a84ebcdeb): complete subsection reference.

<a id="canonical-8d0ccaa9ea479f6e82876cca0f0442f304873b45641a309fd48a8dda0e64cd5e"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 6fef44619f31 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-4bed3c46b74e6cbbab48afae17190ebf57c1a5b2c80e58e9d773fc8a20813360)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-2bf3070a066ff600450035b4d4948eefd5d62bcdf43299bfa689b39a84ebcdeb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-6b08c8ff5901aca0ac1338b0d9bf20f39f31d253847573540b28426d4f9662d3)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-4bed3c46b74e6cbbab48afae17190ebf57c1a5b2c80e58e9d773fc8a20813360"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b06b9637f621033011fda141d0da3f9fd196bc10f0c0d34e23fc7eb71a8686e"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / b6e8ca72cd6a / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-6b08c8ff5901aca0ac1338b0d9bf20f39f31d253847573540b28426d4f9662d3)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](data-sources--voltstack_site--reference--group-006.md#canonical-0da7819a82be7651c383737da6af25d0b7c8ffda8eb6df472902d043942fd968)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info

<a id="canonical-5f31e3aea353699229cc804c513ed461fa3c2330e0f54b7592f6e83df49dd4ce"></a>

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

<a id="canonical-0051536cae1e7738dedbac10266b63df02861a83ced9c11ec4baec0e0a37e4b9"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / b6e8ca72cd6a / 3

<a id="canonical-e169eadb74f7258d46aa1ed5e9fe2be664a745b021f3e85640c2688ad44d7465"></a>

<a id="canonical-88a3e5ea048d2d0f9d541039699e46865fc4c24b9a8a4c81bfa9f9c922e8bea2"></a>

## decryption_provider property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / b6e8ca72cd6a / 4

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

<a id="canonical-69bd89e705c5ddb2d16f4693e5f389402d3d7ec86cf75dcb912f533df6012e81"></a>

<a id="canonical-109fe87efe21d27eadc5fb2363396b55e357a6b4162741a9e5160825c7eafbd8"></a>

## location property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / b6e8ca72cd6a / 5

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

<a id="canonical-4719da7da9f039538752be28c0bf3b01671f276d85463b67a8ed2b885994c919"></a>

<a id="canonical-cd1e941a90447a16a2e1c07b784db2fb17107f586e004632493ea434c75d3fde"></a>

## store_provider property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / b6e8ca72cd6a / 6

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

<a id="canonical-8af4a2e98a229c65f1fa41198f8b140cf198dfd0858a8e6773afb417b360e3a3"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / b6e8ca72cd6a / 7

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](data-sources--voltstack_site--reference--group-006.md#canonical-0da7819a82be7651c383737da6af25d0b7c8ffda8eb6df472902d043942fd968)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-2bf3070a066ff600450035b4d4948eefd5d62bcdf43299bfa689b39a84ebcdeb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e46618eac39f73dd126c780d253fbaca69c49dce76517fedf595140153df6ee7"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2e3cd43cb8c9 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-6b08c8ff5901aca0ac1338b0d9bf20f39f31d253847573540b28426d4f9662d3)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](data-sources--voltstack_site--reference--group-006.md#canonical-0da7819a82be7651c383737da6af25d0b7c8ffda8eb6df472902d043942fd968)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info

<a id="canonical-ff77af0bed918cbfc0e4497121b98da3592406abac9a811f4c8760f5a7aba101"></a>

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

<a id="canonical-015004e2023bb3af09ee33cb695b6e8fdb171c51d79b87db39233e434dd0e2d5"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2e3cd43cb8c9 / 3

<a id="canonical-38c149f422130cc4070d4661f66b81cb9ea3b68c936c9fbc83adf2c06b74971c"></a>

<a id="canonical-a74722ffafcdb9377e6567aef6c112973212f164956ccaae1792dc8af355c01b"></a>

## provider_ref property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2e3cd43cb8c9 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1adaeec010450a34b405122046d2806f99a8fb754577ed4361aa4c182375458d"></a>

<a id="canonical-a94dbc0c5ca6ccaa20a9c8c43d66ab375f733facb9745d66f50f72e349a7453a"></a>

## url property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2e3cd43cb8c9 / 5

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

<a id="canonical-1257b8e7671f9af22693598bf8effda34d391299a0e2775b944de500b181fe87"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2e3cd43cb8c9 / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](data-sources--voltstack_site--reference--group-006.md#canonical-0da7819a82be7651c383737da6af25d0b7c8ffda8eb6df472902d043942fd968)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-9b20b360bfa0a6645821cdc637798da2afdad5ed9aafcd517616fb6f741dbffd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad5a7307fa052c455f126476b9d854ec935f2324163a9d501e658c567a126d06"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 40b2c1525cd3 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-6b08c8ff5901aca0ac1338b0d9bf20f39f31d253847573540b28426d4f9662d3)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap

<a id="canonical-d03aa1fde18c3467c65082fffad5e31528c067e28209569600c0347a0c946fa8"></a>

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

<a id="canonical-ea9ca8a41b13a4a512f07f09ef4be244a5cebf8ce20691c55bcc779e1a2ce48a"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 40b2c1525cd3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dac8408b05d708bf9db6d054f71a6d533056826a3e11d3d1eeaf67cf234964a0"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 40b2c1525cd3 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-6b08c8ff5901aca0ac1338b0d9bf20f39f31d253847573540b28426d4f9662d3)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-82b1f280c58cbeae84539b6007e73c78a0128cff1bf2f26d672f55fac1bb46f5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d7fb73ddcab74752df3cc065c80aefc44e10745a42edf237e5bb18ccac8f1aab"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / cd78a5583a25 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-6b08c8ff5901aca0ac1338b0d9bf20f39f31d253847573540b28426d4f9662d3)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password

<a id="canonical-3acf8242f7a168ee81cb0b7395750aae67205aacadf9a9caf09ed3665e01309b"></a>

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

<a id="canonical-9f656543fa0492d2139bcf0a4f15d9f55d0eb39d5cb1e0db4c9309af84a8b7a1"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / cd78a5583a25 / 3

- [blindfold_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-fd5eeb0835e165cc81d446ca73b31d6f9ca24bfc62dc2c9e9bc5a1dee2c81314): complete subsection reference.

- [clear_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-c37ba5e252b0a4ce8a67cb9aa8dd1b37dcc67a44dcc9f59c8d6105c680b237ef): complete subsection reference.

<a id="canonical-c2a452128e804bf4969b0c8970fca1dbf195c6d33c4e934968b44b88f31a7685"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / cd78a5583a25 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-fd5eeb0835e165cc81d446ca73b31d6f9ca24bfc62dc2c9e9bc5a1dee2c81314)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-c37ba5e252b0a4ce8a67cb9aa8dd1b37dcc67a44dcc9f59c8d6105c680b237ef)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-6b08c8ff5901aca0ac1338b0d9bf20f39f31d253847573540b28426d4f9662d3)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-fd5eeb0835e165cc81d446ca73b31d6f9ca24bfc62dc2c9e9bc5a1dee2c81314"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ca22be338826b6850623526a4cf7141ae5bd63259525970ca59c8371f68c1da"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 90c7ec7ef2dc / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-6b08c8ff5901aca0ac1338b0d9bf20f39f31d253847573540b28426d4f9662d3)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](data-sources--voltstack_site--reference--group-006.md#canonical-82b1f280c58cbeae84539b6007e73c78a0128cff1bf2f26d672f55fac1bb46f5)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info

<a id="canonical-da05462e73cd77ef6ae205b1ffb28eed28c00ac031e998cc65881ffa5242b4e0"></a>

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

<a id="canonical-ad9ce03f0d376c39390aa8660176b16506822ca735f7e9dfb358543898ff7e22"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 90c7ec7ef2dc / 3

<a id="canonical-b4319433dfb2ccb94718e67eb6c7395dce75a7e8547ab80bd627e52c1ef08279"></a>

<a id="canonical-f1b2d4bbfb6b9544870c6f7b56723d3cc1a114adbbab5bf9ea3f0ca0fa2e8499"></a>

## decryption_provider property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 90c7ec7ef2dc / 4

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

<a id="canonical-1aa97329fb9493dd15a858eefa2201af43bd0c639ff93676a4c3b842f1a4721e"></a>

<a id="canonical-4a6efb6de77e26355d8c227e0c8f544ac2cc0801319be427dfcf5aec8184333f"></a>

## location property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 90c7ec7ef2dc / 5

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

<a id="canonical-983db3a06462c4d2a3dc4b5bce96ecac19d7614dd84c6fc8bdf89dbc6b658441"></a>

<a id="canonical-d18e16d1a49a00ba8e25a2dde18d3d49c5d60b1945b941e7936e48184123f466"></a>

## store_provider property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 90c7ec7ef2dc / 6

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

<a id="canonical-b0ad8d01ff38e878c92a8f2e84c951c9f3cdc1a56944ac4a130300800c317eb0"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 90c7ec7ef2dc / 7

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](data-sources--voltstack_site--reference--group-006.md#canonical-82b1f280c58cbeae84539b6007e73c78a0128cff1bf2f26d672f55fac1bb46f5)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-c37ba5e252b0a4ce8a67cb9aa8dd1b37dcc67a44dcc9f59c8d6105c680b237ef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c596d6de5f029650427367583676223e15e90953cb52e40c543abb0a7386d9d4"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / ddc2ad9ef56c / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-6b08c8ff5901aca0ac1338b0d9bf20f39f31d253847573540b28426d4f9662d3)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](data-sources--voltstack_site--reference--group-006.md#canonical-82b1f280c58cbeae84539b6007e73c78a0128cff1bf2f26d672f55fac1bb46f5)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info

<a id="canonical-d8d3399defe809d4c488fad746e55a9a325672c903e2a4c67c7e0c5e267028a6"></a>

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

<a id="canonical-1437230d265847bc2144457f5ed8bf42cdcac27305e3e05b5d5341ed9601af2f"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / ddc2ad9ef56c / 3

<a id="canonical-1cefe7a975cbc448ffc72331ba0d032ba2328abf396194208727bf5a520edb64"></a>

<a id="canonical-35d06a813f04c0be1ede85244b475a61b065354e2f48109439d99b8a61f16bc1"></a>

## provider_ref property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / ddc2ad9ef56c / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-31cb207c8458107b8533c0a68178d8a63a94820d819a1c2bad5e964b8fc124ac"></a>

<a id="canonical-a30e4b0582a3b9ecc380201958472199bb968173cac5fc3beb0f397bb6063084"></a>

## url property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / ddc2ad9ef56c / 5

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

<a id="canonical-e6758b87a456881138ad7de2c70d8f19b6e006f40d3e12dbf236333479273610"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / ddc2ad9ef56c / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](data-sources--voltstack_site--reference--group-006.md#canonical-82b1f280c58cbeae84539b6007e73c78a0128cff1bf2f26d672f55fac1bb46f5)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-79cc1e76076c7cf9601cbbbf98de5cfe2f2338ac110641094f165aa4aa84d2ee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b9b0af4d3f05fac4df6156f7245a4908cde06308ed405e8c9eadea621221575"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 030aa9bcfc8c / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-6b08c8ff5901aca0ac1338b0d9bf20f39f31d253847573540b28426d4f9662d3)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage

<a id="canonical-98f8273a4a70bfd89ee4ced73a6418d76bf67b191927d638eb6589d464d08e0d"></a>

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

<a id="canonical-d9ad1d9fdbaac89a73751a71df786e1a43492123c5a88354c3b4be5814cf747f"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 030aa9bcfc8c / 3

<a id="canonical-f800294b63e89d8cfcb59f59d6601602df4d13258a73b1c1601e14d6e51be781"></a>
