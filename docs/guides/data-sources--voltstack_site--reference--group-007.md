---
page_title: "xcsh_voltstack_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_voltstack_site reference."
---

# xcsh_voltstack_site reference

<a id="canonical-4ed6154f9f938cd3cf5c638d430f3d842880c31adee4e976dc211f40c1f68607"></a>

## labels property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 030aa9bcfc8c / 4

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

- [volume_defaults](data-sources--voltstack_site--reference--group-007.md#canonical-ace7cff9e08ba63cf23dc073eeac62a1bda989a1635c2aadafe237cd2528e91d): complete subsection reference.

<a id="canonical-8bea313c0fd7a06b95588059ad4ce19ec0869712c90cef8429f60353f5354ce3"></a>

<a id="canonical-099211375b611d3e391dbfe44c475810211631d6ced30f0439bb0bebc8d60bbe"></a>

## zone property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 030aa9bcfc8c / 5

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

<a id="canonical-903dc848d2a7a569e9dc9d15fdf45286ba07b3953150982b9cb55aac1c41cfdb"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 030aa9bcfc8c / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults](data-sources--voltstack_site--reference--group-007.md#canonical-ace7cff9e08ba63cf23dc073eeac62a1bda989a1635c2aadafe237cd2528e91d)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-6b08c8ff5901aca0ac1338b0d9bf20f39f31d253847573540b28426d4f9662d3)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-ace7cff9e08ba63cf23dc073eeac62a1bda989a1635c2aadafe237cd2528e91d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8c0d8f2dbacbf067886208a88572c3bbc76a6b243dc55e73da4fcbcd593ac576"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 1271d738446d / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-6b08c8ff5901aca0ac1338b0d9bf20f39f31d253847573540b28426d4f9662d3)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](data-sources--voltstack_site--reference--group-006.md#canonical-79cc1e76076c7cf9601cbbbf98de5cfe2f2338ac110641094f165aa4aa84d2ee)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults

<a id="canonical-d0c147b2cd610cdf0ea1dd72bc8662efba54d61f6e945fd0dd5cd4443dc42809"></a>

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

<a id="canonical-2b93080b552afd634eb879c06d290dc31ee2e3f6756fe09834357e0c6aac81d0"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 1271d738446d / 3

<a id="canonical-b600959f1c4457b1849c77e88bd22d93d1a9e00143c871d4f35c059a4fe0e856"></a>

<a id="canonical-31f741991f08c26e60a112efb9b3170f215108e294d86f1eeb24557bafac0452"></a>

## adaptive_qos_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 1271d738446d / 4

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

<a id="canonical-8017d3647e9aa0482ec938b65ccccd74f8e2cabf02ea6f8d3ba543e6871937c1"></a>

<a id="canonical-e9cad0bd661708f4f6618d79ca1df359dde303ff2f72e4055b241b82d8101dd0"></a>

## encryption property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 1271d738446d / 5

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

<a id="canonical-0a8483c30bc2c834db88e9a3cf155a0830962d914368e81fbc9b9e742e227b14"></a>

<a id="canonical-78f0a572379718ccda567418aa5c526e623994af7138177eb43902d90a9a63cd"></a>

## export_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 1271d738446d / 6

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

- [no_qos](data-sources--voltstack_site--reference--group-007.md#canonical-260c1feecaec40c658ff6d67d0f4c9f1df84236a8f596a00f4a3d554b446f398): complete subsection reference.

<a id="canonical-4471ebfa3df90bb9289c1207dcd802d66d87d1d297b817374dca8ac6b56bc3ba"></a>

<a id="canonical-c5179005826cee847c6b9835310650d67794eb6e292aa651a68c5d74bf902efc"></a>

## qos_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 1271d738446d / 7

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

<a id="canonical-933c617230cac338e4b0c6b548537f2a93ff277541f00104b08dfe25681ce488"></a>

<a id="canonical-276d73e5ed9dea784994bbb87629fbdddbe522934c265f7837f657007792c360"></a>

## security_style property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 1271d738446d / 8

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

<a id="canonical-3868fbc25a18fa23e9ce0071d41bc4606d40ea932714725a8a76310fbedecf5d"></a>

<a id="canonical-9fce0b13dd58bb9a5bc17e7ba92e3e5634b8ba07a0aabe5cbff7ec1276943711"></a>

## snapshot_dir property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 1271d738446d / 9

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

<a id="canonical-22372a5fe9d024edb7bbc130d3da76b41a46d198aacbf0778228dfb440bfadda"></a>

<a id="canonical-1603fd9f63bad27210fa59b46d2ef411ad7cf473c24229c99cc82981a0639e16"></a>

## snapshot_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 1271d738446d / 10

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

<a id="canonical-6abdf0e57ca51270337103c093370aba44ae557474f4a9f05ecef54946157c94"></a>

<a id="canonical-9b5209cf1fef14a21ea5a1fa8b2466798a0f7bff92af976fc826813e2e4b826f"></a>

## snapshot_reserve property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 1271d738446d / 11

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

<a id="canonical-56a91b1695c4c25e0889f689b71ab3f457fce0d8b05088cf7e2600d345ef0da4"></a>

<a id="canonical-bb2d7890d36cd3d5d78ec784e801c3657988bccc521d4616dd5ff54b56793024"></a>

## space_reserve property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 1271d738446d / 12

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

<a id="canonical-90d0612428813555b3727c33688510e8bff4c9ed5a533c6709cec6a023a961e8"></a>

<a id="canonical-ed66b9a40f3905952035ded3ccec4f5752f7f360e2f258209d28a2fbfab47b8a"></a>

## split_on_clone property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 1271d738446d / 13

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

<a id="canonical-bf782808e602ca53a34e2eb32a6e70e4ef759b3c0cf82e8878cba40949300593"></a>

<a id="canonical-ac5166c0e99d5910912fbf34bb285cd9e9897f465c2cf0333e0ae5415a3f1a5a"></a>

## tiering_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 1271d738446d / 14

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

<a id="canonical-1c3a9aae62bd831bec72ca7e11724bb5549cbcda5e45526a0cfbe519e4581445"></a>

<a id="canonical-d60a1d02ed8b063ae0a95fbe271e00d0c205d2f88b4c63859d3b25598acf3064"></a>

## unix_permissions property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 1271d738446d / 15

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

<a id="canonical-7758d50253b11d77340d401aa551cfd8244c04a0ece65866761b4ae4efad5509"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 1271d738446d / 16

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos](data-sources--voltstack_site--reference--group-007.md#canonical-260c1feecaec40c658ff6d67d0f4c9f1df84236a8f596a00f4a3d554b446f398)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](data-sources--voltstack_site--reference--group-006.md#canonical-79cc1e76076c7cf9601cbbbf98de5cfe2f2338ac110641094f165aa4aa84d2ee)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-260c1feecaec40c658ff6d67d0f4c9f1df84236a8f596a00f4a3d554b446f398"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-099ec5d58e5aebfaa0a3c61df0bc36efb599ea35dd818bb7fe09450dd2d0e67a"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0cfc20115c47 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-6b08c8ff5901aca0ac1338b0d9bf20f39f31d253847573540b28426d4f9662d3)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](data-sources--voltstack_site--reference--group-006.md#canonical-79cc1e76076c7cf9601cbbbf98de5cfe2f2338ac110641094f165aa4aa84d2ee)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults](data-sources--voltstack_site--reference--group-007.md#canonical-ace7cff9e08ba63cf23dc073eeac62a1bda989a1635c2aadafe237cd2528e91d)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos

<a id="canonical-974943154788b35147f6df443f5cc990f63b30c508cd3094e57036e60ca50434"></a>

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

<a id="canonical-3f2a9fa4cede2947ba2152e4f635f197d26f8d294953c39b7a826540bea69c97"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0cfc20115c47 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-aa29e933f725430b5710d4c1806601987af00f91d2b6ef4f6e3af9e28f99f568"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0cfc20115c47 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults](data-sources--voltstack_site--reference--group-007.md#canonical-ace7cff9e08ba63cf23dc073eeac62a1bda989a1635c2aadafe237cd2528e91d)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-ebe00d06a5ec80e287e74893606d51318c5c59d010b4482d51f10ac95faf3fe9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-462e3b599f99ac63edc71dee3761bc8d6948ee2bcb1890497fda012733230044"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2e433c1ca86e / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-6b08c8ff5901aca0ac1338b0d9bf20f39f31d253847573540b28426d4f9662d3)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap

<a id="canonical-4d58ceb601cf73bc74e9195b6182ce8711dbf2a887a043c8af58672f8211d21b"></a>

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

<a id="canonical-0bdc101b8aa25b96c71c90b6dcac87aab8748948c2702aa278e31dfbcbf20a22"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2e433c1ca86e / 3

- [chap_initiator_secret](data-sources--voltstack_site--reference--group-007.md#canonical-406d1ca781653a8de2ffa1e1db04d287276706e8920897a540148b780a886797): complete subsection reference.

- [chap_target_initiator_secret](data-sources--voltstack_site--reference--group-007.md#canonical-7c2e369307ed2d6a2d15c4e3890053dc52d6dc08ccb9705623f54453e7c72f88): complete subsection reference.

<a id="canonical-3674283bb31f2f8a83ab12c1ffce90ce6cc7b7c8416ce7e2aaa836e472a5c242"></a>

<a id="canonical-6922f10c3f42aec59422dd26ed45be0b2e4f59d84d171abcd6d0e2ff9f40cb7d"></a>

## chap_target_username property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2e433c1ca86e / 4

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

<a id="canonical-76148f139cfe9d0cda7424c3f3987c15a5a3792112197d183cb463030d1a347f"></a>

<a id="canonical-f2ed4a74f59dabcbc19847f6a60094fcd7bac7dfab83fb0b758e74433bc283a3"></a>

## chap_username property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2e433c1ca86e / 5

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

<a id="canonical-4dc68cf2741a64b0c2da0dac10e8a54392e02b5fd7843730cc2b363d8f30d02b"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2e433c1ca86e / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](data-sources--voltstack_site--reference--group-007.md#canonical-406d1ca781653a8de2ffa1e1db04d287276706e8920897a540148b780a886797)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](data-sources--voltstack_site--reference--group-007.md#canonical-7c2e369307ed2d6a2d15c4e3890053dc52d6dc08ccb9705623f54453e7c72f88)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-6b08c8ff5901aca0ac1338b0d9bf20f39f31d253847573540b28426d4f9662d3)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-406d1ca781653a8de2ffa1e1db04d287276706e8920897a540148b780a886797"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b383eba8f2079715e5fba12f426c57e366ce852d9a3eeab40bc93cace7731ff"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 1c986dfc14f1 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-6b08c8ff5901aca0ac1338b0d9bf20f39f31d253847573540b28426d4f9662d3)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--voltstack_site--reference--group-007.md#canonical-ebe00d06a5ec80e287e74893606d51318c5c59d010b4482d51f10ac95faf3fe9)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret

<a id="canonical-01e04928e7db2f00b921426b1f084b9514ee6d739f73abc50720383285071769"></a>

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

<a id="canonical-25574505b8cf70b86e18911eed77f9b3c1cb45e254f4230c6a53e2242f2a784e"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 1c986dfc14f1 / 3

- [blindfold_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-1ea4682e031d1e37990d2dcb8a2f6d26959d1188b8a0c218d3775e4d18348c52): complete subsection reference.

- [clear_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-6438b85192975a35777c52f73ee54f4475e20be98f34374cab3bc331b9e591c3): complete subsection reference.

<a id="canonical-9b52c15e411979d0b3f97c13770bc9119368c88f622bee806423b31c1bc7fd9a"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 1c986dfc14f1 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-1ea4682e031d1e37990d2dcb8a2f6d26959d1188b8a0c218d3775e4d18348c52)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-6438b85192975a35777c52f73ee54f4475e20be98f34374cab3bc331b9e591c3)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--voltstack_site--reference--group-007.md#canonical-ebe00d06a5ec80e287e74893606d51318c5c59d010b4482d51f10ac95faf3fe9)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-1ea4682e031d1e37990d2dcb8a2f6d26959d1188b8a0c218d3775e4d18348c52"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6487312aef212a6321428409cf9cb0215ee45a987cd114463916286f7fbf85e9"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 56ed6fe19d63 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-6b08c8ff5901aca0ac1338b0d9bf20f39f31d253847573540b28426d4f9662d3)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--voltstack_site--reference--group-007.md#canonical-ebe00d06a5ec80e287e74893606d51318c5c59d010b4482d51f10ac95faf3fe9)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](data-sources--voltstack_site--reference--group-007.md#canonical-406d1ca781653a8de2ffa1e1db04d287276706e8920897a540148b780a886797)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info

<a id="canonical-45642075e326e4d92dd6fbf912ce75513b35c2e7ed04aabe9575d40851639032"></a>

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

<a id="canonical-874d44ce218352fb8ee61cac0b521f7fc8881549ad5975a724342b15d66363ae"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 56ed6fe19d63 / 3

<a id="canonical-ebc4892cf5efd2175e9167cdee3303c3c3ad42a5028b6ba030066c7e6f38fd6a"></a>

<a id="canonical-8fa6c22f1fd2a2df1d1000b9938ad5d15f57323b56044fff95e992ba7955e80d"></a>

## decryption_provider property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 56ed6fe19d63 / 4

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

<a id="canonical-51859bd2a3d70a64e72c32a6633b3ee433cee49ced107c66b18af617f619184c"></a>

<a id="canonical-941aacbfd9bd74eda909749630c8a58c82f1fce64adc813caada71c8c7f75bc2"></a>

## location property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 56ed6fe19d63 / 5

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

<a id="canonical-78ca5e18b83aa4ba83ea668a325526ba9c6570612a157b13fb10a88ee9377b8e"></a>

<a id="canonical-0da43a405615abe49ab33ce1c7c7eaf3c3fa6d8bace8d3b333c356fb25616ef0"></a>

## store_provider property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 56ed6fe19d63 / 6

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

<a id="canonical-019ccad121b66d075ffb6a5aec73cdb30d7f153f9918cec3dc263c2c1066ad4e"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 56ed6fe19d63 / 7

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](data-sources--voltstack_site--reference--group-007.md#canonical-406d1ca781653a8de2ffa1e1db04d287276706e8920897a540148b780a886797)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-6438b85192975a35777c52f73ee54f4475e20be98f34374cab3bc331b9e591c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-64ad5e44c918a1c23c79634902ecb30e0fcc7d68a9fba3147cba5ec86b9e42d0"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / c8c4897edfcc / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-6b08c8ff5901aca0ac1338b0d9bf20f39f31d253847573540b28426d4f9662d3)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--voltstack_site--reference--group-007.md#canonical-ebe00d06a5ec80e287e74893606d51318c5c59d010b4482d51f10ac95faf3fe9)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](data-sources--voltstack_site--reference--group-007.md#canonical-406d1ca781653a8de2ffa1e1db04d287276706e8920897a540148b780a886797)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info

<a id="canonical-7b10916fb870ba4ac4653cfa6fb00d77ee3d3ebeb92df700a20ec945b5df0375"></a>

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

<a id="canonical-f4d5c8302152c7fcb5da923925c08cfd1c4c78ea84f32327e6806505fbb6d9df"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / c8c4897edfcc / 3

<a id="canonical-876f9d8dea53736f5da36f5402f071c53cd2bcb2fa4efc92f552fa59042c5c5c"></a>

<a id="canonical-e7bd0cf5642219d282f9cd2c296805d95701c5d4fd92289d033ddba9b4456f15"></a>

## provider_ref property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / c8c4897edfcc / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-7da6c8a04a7f908cafc91aa9e8534d5e8649c520d9eba3d43b17ba6d947b70d7"></a>

<a id="canonical-6a777345039bd3fb12db35f2830c71da1ea6ba5aca0c360a48148eef387ccfbe"></a>

## url property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / c8c4897edfcc / 5

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

<a id="canonical-0787588cd757ae1d1568022959af187774085f4b351f8a4cbf0e73ebbc5e17a9"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / c8c4897edfcc / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](data-sources--voltstack_site--reference--group-007.md#canonical-406d1ca781653a8de2ffa1e1db04d287276706e8920897a540148b780a886797)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-7c2e369307ed2d6a2d15c4e3890053dc52d6dc08ccb9705623f54453e7c72f88"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f640088d339477fffef1bc5187545bdadfc9966b8320b6a1d2beff59860bc523"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / f2cb8410a738 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-6b08c8ff5901aca0ac1338b0d9bf20f39f31d253847573540b28426d4f9662d3)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--voltstack_site--reference--group-007.md#canonical-ebe00d06a5ec80e287e74893606d51318c5c59d010b4482d51f10ac95faf3fe9)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret

<a id="canonical-a4f4a76145a98d0f138213c9ba15d5587acbef7a37c893a728194f54fcf075bc"></a>

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

<a id="canonical-2c6d42a007f47708ca00faaecccf46aa40b058370415c58b92934de518799511"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / f2cb8410a738 / 3

- [blindfold_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-d27a2a98fc3c2e85e3247aee43085efce58e44d9927cbdf80e178c6b837f6b0c): complete subsection reference.

- [clear_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-d9742b5bae0dc9c3f9c181fbf855976d77222e13c57dcc69217137d2f71fd10f): complete subsection reference.

<a id="canonical-85c72f32b5e12023250c34928434aa0e0e698716238aab641478592583ab208a"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / f2cb8410a738 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-d27a2a98fc3c2e85e3247aee43085efce58e44d9927cbdf80e178c6b837f6b0c)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-d9742b5bae0dc9c3f9c181fbf855976d77222e13c57dcc69217137d2f71fd10f)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--voltstack_site--reference--group-007.md#canonical-ebe00d06a5ec80e287e74893606d51318c5c59d010b4482d51f10ac95faf3fe9)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-d27a2a98fc3c2e85e3247aee43085efce58e44d9927cbdf80e178c6b837f6b0c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5757edf32a88bfde6eb0f7a5c5f8a7722762247feac0eec5ef9584768e2df545"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 9853b5b00a86 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-6b08c8ff5901aca0ac1338b0d9bf20f39f31d253847573540b28426d4f9662d3)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--voltstack_site--reference--group-007.md#canonical-ebe00d06a5ec80e287e74893606d51318c5c59d010b4482d51f10ac95faf3fe9)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](data-sources--voltstack_site--reference--group-007.md#canonical-7c2e369307ed2d6a2d15c4e3890053dc52d6dc08ccb9705623f54453e7c72f88)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info

<a id="canonical-781466ba257018b62869378071bf42bc433264a6075ce2503fbdda2f4586a5c6"></a>

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

<a id="canonical-8b6c5f52dc0a17cb750b09e1cd69451bd804c2c24e6d9a7a1f13a1f9802fb63c"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 9853b5b00a86 / 3

<a id="canonical-2c92437183794667c67b9c44c8bbad3329f2a4770476a22e02a2a68c8fc6bc0a"></a>

<a id="canonical-b54e5d1f3b32c9b43890a3468af676286ab47191b0a7ab020604761295edff28"></a>

## decryption_provider property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 9853b5b00a86 / 4

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

<a id="canonical-26c60ec4318ace8d40474ac25f0a85e941707b691f3cf2e80542e69d161e868d"></a>

<a id="canonical-2badd71f5b0ca6a290b4941aafda47dcfaad2cf122f9a7aa644d3cacf9a6d6fe"></a>

## location property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 9853b5b00a86 / 5

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

<a id="canonical-b2a51952b4ae7e774636c0e1a7aa9f7f36c56bae4e7fd83e9dc7c7d82bd8b1cd"></a>

<a id="canonical-37aaa3477675b4b97b8b5849ffa859be54e8e5bd89821d29c3f2d898f3f34e37"></a>

## store_provider property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 9853b5b00a86 / 6

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

<a id="canonical-24eb1be3cea6893423e823176563ef732e57a24e90ff47133190db76993559dd"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 9853b5b00a86 / 7

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](data-sources--voltstack_site--reference--group-007.md#canonical-7c2e369307ed2d6a2d15c4e3890053dc52d6dc08ccb9705623f54453e7c72f88)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-d9742b5bae0dc9c3f9c181fbf855976d77222e13c57dcc69217137d2f71fd10f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3263e0529cc414b9fbe21fa06cd597e8917fd6bcf10a3f9f68e77db9123d4559"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / fc6c6f8c5b48 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-6b08c8ff5901aca0ac1338b0d9bf20f39f31d253847573540b28426d4f9662d3)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--voltstack_site--reference--group-007.md#canonical-ebe00d06a5ec80e287e74893606d51318c5c59d010b4482d51f10ac95faf3fe9)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](data-sources--voltstack_site--reference--group-007.md#canonical-7c2e369307ed2d6a2d15c4e3890053dc52d6dc08ccb9705623f54453e7c72f88)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info

<a id="canonical-b1e5b94747fba01ca3b98846a4534696bb2e3f97a0b416cb88a1f2448bf64025"></a>

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

<a id="canonical-4867f093fa46939a8cfd985dc1d32d798c7d7df56f4fcbd0f5dc8d5b6c327918"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / fc6c6f8c5b48 / 3

<a id="canonical-05261e31edb2377e108fdcca112e0b581ba4b984ecbb6d7791a8d91a58fabda4"></a>

<a id="canonical-45b63963e26662bddae3b56463f5310ceeea663854de6cf05f2477e2bfc619a6"></a>

## provider_ref property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / fc6c6f8c5b48 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-49a797e43a00cfe6225ded630ae8e87e2d1176dd4091115c1705077aa8b3d0ec"></a>

<a id="canonical-8fc0db27893004e04250610ba044bbabc67504436f314e715a2c40f9c83c750e"></a>

## url property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / fc6c6f8c5b48 / 5

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

<a id="canonical-77ddd3f0d3efdb4fa979e3d6af6744c98119a8fca93e464d9d90a3d172312cf6"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / fc6c6f8c5b48 / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](data-sources--voltstack_site--reference--group-007.md#canonical-7c2e369307ed2d6a2d15c4e3890053dc52d6dc08ccb9705623f54453e7c72f88)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-049d958c8b1328a10ee8c36da36fdbf3fd1cddf195cc903283f0846010b75345"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f7e9a9827b955d6f25794642125ee4090f265fa7c92e3cc4fac6a295ae486fad"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e90e0166b94d / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-6b08c8ff5901aca0ac1338b0d9bf20f39f31d253847573540b28426d4f9662d3)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults

<a id="canonical-cd47b55ed24e8d556af3682aa71f4f3b1541b30aab8d6d09f13b5a0c1c1a3df6"></a>

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

<a id="canonical-89e3d91dcfd69334e536b1090eaa3c15fcd4e84886f49c6c074b5b5b10d98dde"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e90e0166b94d / 3

<a id="canonical-b7754dbc7a654522d3a8a925262b0ea51d9f7fd74439d527523e68c0eb76e739"></a>

<a id="canonical-0d266f96d1f0a373838a442b4b15d77ae879c8e3544497ca71175869a61c95b4"></a>

## adaptive_qos_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e90e0166b94d / 4

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

<a id="canonical-68216d43370d9a1048c8ead23c0ca0ed480348db0cc1c5ed5f9e80b9eb4117b7"></a>

<a id="canonical-0899bfb28af1554147d2d7f2c144fa033ef33a60ef78cf565298256f4e23ee65"></a>

## encryption property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e90e0166b94d / 5

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

<a id="canonical-d54e69fa1a80271e8b0efc82b2c043fdbf8a89dc800c5f6d8dd988417182a7a5"></a>

<a id="canonical-e5a69ea3664165ed3412bac9c9b8ea9398f5fe65316922c95749a63a11204063"></a>

## export_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e90e0166b94d / 6

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

- [no_qos](data-sources--voltstack_site--reference--group-007.md#canonical-011cddf43bf71e95c32ff01eaacf28f0cfdcdce4971ade0206299f346cda8d40): complete subsection reference.

<a id="canonical-43670406436398b652ece626939bd1d74772e78271e30b50ea9dc57bc061bf16"></a>

<a id="canonical-62354bbeadc4b771ae187e4b807c70aa7f5222b97ead467f73e37d21aab5a0e4"></a>

## qos_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e90e0166b94d / 7

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

<a id="canonical-0280c93d68a5d73b3cedf97bcdd0e299e5df10d37355a070a989a86891331d43"></a>

<a id="canonical-19f9638b8f625aee057b80e437326cb3e809c7b5af24b48385ee2389bd707fba"></a>

## security_style property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e90e0166b94d / 8

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

<a id="canonical-cc71d1924ec4bba3a574179014d25005b405cb925555b6a92a6ae0c0635dd83f"></a>

<a id="canonical-8de7c4f0acfc5674e7a5783d4088c142cf164c6e32380020b47c26b0e12d47f1"></a>

## snapshot_dir property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e90e0166b94d / 9

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

<a id="canonical-163c931e2b47b5454f3d49ca4673c1ca90f0e0b4e7f50720f8b28e12be25fa83"></a>

<a id="canonical-1d54de76f13af744a02521e7db6fb251edd9a96a9fdc1d3cd787c26340d22fc5"></a>

## snapshot_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e90e0166b94d / 10

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

<a id="canonical-871e5017a43a83a3fd629fe7989cef93e146a7f0ae030f320d6077c841e3c252"></a>

<a id="canonical-3e2e5db038a7d7d7ab6debe13e42b93d18a6166ba8ae68348abb5305b3204c2d"></a>

## snapshot_reserve property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e90e0166b94d / 11

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

<a id="canonical-5d4c1d9431161ebab83632406368f2eaa8d3db053410ea8db117b2bc5ac4cd6f"></a>

<a id="canonical-285904915b278899a2a8c4352b0effffe3be7e2cbaa5dc3070258ba5f1f13a2a"></a>

## space_reserve property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e90e0166b94d / 12

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

<a id="canonical-3338adf35cfe4bf0fd95bf54c73bc7a47e233819355d2975d639501eb0fa3f6a"></a>

<a id="canonical-b6de7d0547c83874099399e94e5aa091028349817796c7e48d41a36929ac7fd2"></a>

## split_on_clone property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e90e0166b94d / 13

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

<a id="canonical-991309e992b9adfaa637d113c36644f5c5f9f9bf8f5ec72e9433021b3817043c"></a>

<a id="canonical-27bbca2a55b899095807e90742ae3f8da037d64d9c7d973bce364bbbaa7ff84e"></a>

## tiering_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e90e0166b94d / 14

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

<a id="canonical-46b75b645bd537c6a293ebd7efe5f41957c8f4ebf2d3ccc634c31d43ab1f062e"></a>

<a id="canonical-64107e1fc9c6cd4a7d36eaf8217c751b2beb02686e401a4f4cb1504804c2bd20"></a>

## unix_permissions property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e90e0166b94d / 15

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

<a id="canonical-9408e648813cb600c204c54aa31bbdc5acf4bf37a351a217e7bc89c015c2f70b"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / e90e0166b94d / 16

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos](data-sources--voltstack_site--reference--group-007.md#canonical-011cddf43bf71e95c32ff01eaacf28f0cfdcdce4971ade0206299f346cda8d40)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-6b08c8ff5901aca0ac1338b0d9bf20f39f31d253847573540b28426d4f9662d3)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-011cddf43bf71e95c32ff01eaacf28f0cfdcdce4971ade0206299f346cda8d40"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b6c937c6ca12bd36b19cb3e0a1c663df8de0a03ea02049af93f12293de23455f"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2666dd41c8b7 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-8c46ad16ba2e86b45acd966f7051f3862a5fffdefc12ccee3420bb970ab9cc27)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-6b08c8ff5901aca0ac1338b0d9bf20f39f31d253847573540b28426d4f9662d3)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults](data-sources--voltstack_site--reference--group-007.md#canonical-049d958c8b1328a10ee8c36da36fdbf3fd1cddf195cc903283f0846010b75345)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos

<a id="canonical-dc65b0ffa44bd1df484a65afe628fbac0b9b7c72bbae1d37a09a5f473e9c5178"></a>

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

<a id="canonical-bfd5b48180e9f6937ef4bb07ef086fbeb9af5aa076c9d30548247ecc0863d52d"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2666dd41c8b7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-927350e0d7e09447f034b4722443c8ab9f8948495cdc5bebd1cce77cb0ff491c"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2666dd41c8b7 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults](data-sources--voltstack_site--reference--group-007.md#canonical-049d958c8b1328a10ee8c36da36fdbf3fd1cddf195cc903283f0846010b75345)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-678801c40c6f67b9dfaf26614f3fc8e471a101a0988408bcdab7079e531d9a28"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-274523412e75d0c81629165af9a43d93b19e134ea6b729091bbcb484f65442a0"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / a5a32e7d855e / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator

<a id="canonical-420bba9e5167ac91059c012fb3567d478b0e327cbe524165567080b4fb78deba"></a>

Type: `"single"`. Computed.

Device configuration for Pure Storage Service Orchestrator.

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

<a id="canonical-b803e1174a594b31a7d3ef56dfdaa1e602d8d40f4a221e776f9f50076c70731f"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / a5a32e7d855e / 3

- [arrays](data-sources--voltstack_site--reference--group-007.md#canonical-dfd367f909005118fe66b891c33f0ecf305b417a73a712f9f57deec2ee2353b8): complete subsection reference.

<a id="canonical-4456417a4daebb8f4e77289ff68a443d3571489a340691c65a2f3000de91f814"></a>

<a id="canonical-53077672ca0b19527d765e103d78a56343838d35252c4b35211e2e6026335256"></a>

## cluster_id property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / a5a32e7d855e / 4

Type: `"string"`. Computed.

ClusterID is added as a prefix for all volumes created by this PSO installation. ClusterID is also
used to identify the volumes used by the datastore, pso-db. ClusterID MUST BE UNIQUE for multiple
K8s clusters running on top of the same storage arrays.

Upstream description:

ClusterID is added as a prefix for all volumes created by this PSO installation. ClusterID is also
used to identify the volumes used by the datastore, pso-db. ClusterID MUST BE UNIQUE for multiple
K8s clusters running on top of the same storage arrays. Characters allowed: alphanumeric and
underscores.

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

<a id="canonical-25572bbaf49ec85b3a7fb1998e2166dff9c0ae4a0f6ee22fc79ccbe440f52d7f"></a>

<a id="canonical-c82fa372ac5d15377660a04d5e41375b28a756ed55d3c3ecc5d903df840498ad"></a>

## enable_storage_topology property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / a5a32e7d855e / 5

Type: `"bool"`. Computed.

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

<a id="canonical-41dcb8ad97c5fe04644dbe4bfe6ea0d47cf1a4724f4ffa67726330bfb73e1a91"></a>

<a id="canonical-193ebdf5e7996d2faebc5c1d960a68067e12f369ef4914e68a55b6abf5d73bb9"></a>

## enable_strict_topology property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / a5a32e7d855e / 6

Type: `"bool"`. Computed.

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

<a id="canonical-7b9cf125deea1dbcd7dab408245d1dcbeb3bc083a4689c4fa8ec7a898a4632e6"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / a5a32e7d855e / 7

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--voltstack_site--reference--group-007.md#canonical-dfd367f909005118fe66b891c33f0ecf305b417a73a712f9f57deec2ee2353b8)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-dfd367f909005118fe66b891c33f0ecf305b417a73a712f9f57deec2ee2353b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e927633aada4328d99e6243dad5523a1c7fe3b67765b6e3ceb3bf86884cce43"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 16664714c9d1 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](data-sources--voltstack_site--reference--group-007.md#canonical-678801c40c6f67b9dfaf26614f3fc8e471a101a0988408bcdab7079e531d9a28)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays

<a id="canonical-3cb3c4426b75f1edf6b8a81db88610fdbd366e2ccb213d9761c3bfed4d37143c"></a>

Type: `"single"`. Computed.

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

<a id="canonical-ae5317e734f002b20272bfbf9906f0beec06fdab8dd4554d9ca1903eaa78cfa2"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 16664714c9d1 / 3

- [flash_array](data-sources--voltstack_site--reference--group-007.md#canonical-50dc081106afe105c82c6783368207f74e25b9263381f3fd546e6f75d724505f): complete subsection reference.

- [flash_blade](data-sources--voltstack_site--reference--group-007.md#canonical-0760631ea0372ec89b870269224d00e592f5f1eee73801bbb8a1405da724a6fc): complete subsection reference.

<a id="canonical-fd2d3fb58f424351304aac1a5f9e97904ebfc4d26a75f803d15b5eeb85d7f8cd"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 16664714c9d1 / 4

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](data-sources--voltstack_site--reference--group-007.md#canonical-50dc081106afe105c82c6783368207f74e25b9263381f3fd546e6f75d724505f)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--voltstack_site--reference--group-007.md#canonical-0760631ea0372ec89b870269224d00e592f5f1eee73801bbb8a1405da724a6fc)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](data-sources--voltstack_site--reference--group-007.md#canonical-678801c40c6f67b9dfaf26614f3fc8e471a101a0988408bcdab7079e531d9a28)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-50dc081106afe105c82c6783368207f74e25b9263381f3fd546e6f75d724505f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-339dbd1babad19381fcc4a06f204cd745455daa846d11d68904217bf862a7e79"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 8c8797530fd1 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](data-sources--voltstack_site--reference--group-007.md#canonical-678801c40c6f67b9dfaf26614f3fc8e471a101a0988408bcdab7079e531d9a28)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--voltstack_site--reference--group-007.md#canonical-dfd367f909005118fe66b891c33f0ecf305b417a73a712f9f57deec2ee2353b8)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array

<a id="canonical-b8e946c47673c19029f0f827524cb241d971889fd925be3db2126f88649bb652"></a>

Type: `"single"`. Computed.

Specify what storage flash arrays should be managed the plugin.

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

<a id="canonical-900bd317411d8d51121ef7f5bdb7435ef74810139777b2f492e64a37ef6144da"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 8c8797530fd1 / 3

<a id="canonical-cd70b2d079189988f09d6adf6277ceb83b5d80198df77a40bc5af29590caee44"></a>

<a id="canonical-c9a03f752779cffeee9bfcdd74eb0a45e30c1673bb622aa5ecf1deb53a439b93"></a>

## default_fs_opt property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 8c8797530fd1 / 4

Type: `"string"`. Computed.

Block volume default mkfs OPTIONS. Not recommended to change!

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

<a id="canonical-375b9360362033c2766a7bcace480bfbfc00b52ba0ee5d30403b18cdf93b9632"></a>

<a id="canonical-410ea55fabd0c0e17d08b949b0eeaa4a53b7346975122e0b4d7ebedf532073be"></a>

## default_fs_type property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 8c8797530fd1 / 5

Type: `"string"`. Computed.

\[Enum: xfs|ext4\] Block volume default filesystem type. Not recommended to change!. Possible values
are \`xfs\`, \`ext4\`.

Upstream description:

Block volume default filesystem type. Not recommended to change!

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

<a id="canonical-3e233722f4f2203ad2b1bcf0592103b696c7968a3aef7eae77a5d0d28ac714d5"></a>

<a id="canonical-2be7d917204335980d4ceb8a0c246eddb441443bed68380f12ab71b87b40bf5d"></a>

## default_mount_opts property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 8c8797530fd1 / 6

Type: `["list", "string"]`. Computed.

Block volume default filesystem mount OPTIONS. Not recommended to change!

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

<a id="canonical-a82e4a5dce2931858baf60bfe7a5a4daa6d69ed1a911c97c30285d32dd6fb27f"></a>

<a id="canonical-7624d59a28cf9bcea18799c8329146be4dd3fe4dca0e53b0098272af75546daa"></a>

## disable_preempt_attachments property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 8c8797530fd1 / 7

Type: `"bool"`. Computed.

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

- [flash_arrays](data-sources--voltstack_site--reference--group-007.md#canonical-4688a335030d30b7b2c5948ddd2b619f7cd5c4a236cfd6abc58990c7dcfffb07): complete subsection reference.

<a id="canonical-e78cf9ed904a612d48428603930d1c63fe9c1060d9ccce612d49295a90a40ba0"></a>

<a id="canonical-db15adcf65f516c193247d2a857933ee9f7e5ae6b84330307133e4513a40120a"></a>

## iscsi_login_timeout property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 8c8797530fd1 / 8

Type: `"number"`. Computed.

ISCSI login timeout in seconds. Not recommended to change!

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

<a id="canonical-380bc122c20b5d4f95a60556b3b9790a6d9428c56bbf410bd17b464270ec9a45"></a>

<a id="canonical-7dee0f3b928ee6721108dd192f33f9f60a0e3139333df5bea53beec22858577e"></a>

## san_type property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 8c8797530fd1 / 9

Type: `"string"`. Computed.

\[Enum: ISCSI|FC\] Block volume access protocol, either ISCSI or FC. Possible values are \`ISCSI\`,
\`FC\`.

Upstream description:

Block volume access protocol, either ISCSI or FC.

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

<a id="canonical-c1fe0a73a2ff74acd718b82f85453f3387cb9a81a76834fae7303b3ebfcd814c"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 8c8797530fd1 / 10

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](data-sources--voltstack_site--reference--group-007.md#canonical-4688a335030d30b7b2c5948ddd2b619f7cd5c4a236cfd6abc58990c7dcfffb07)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--voltstack_site--reference--group-007.md#canonical-dfd367f909005118fe66b891c33f0ecf305b417a73a712f9f57deec2ee2353b8)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-4688a335030d30b7b2c5948ddd2b619f7cd5c4a236cfd6abc58990c7dcfffb07"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-00ed2cc578815cde06ef296f8d320c8f0625051c694b80460180914fcc2a666b"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / e29490c124e3 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](data-sources--voltstack_site--reference--group-007.md#canonical-678801c40c6f67b9dfaf26614f3fc8e471a101a0988408bcdab7079e531d9a28)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--voltstack_site--reference--group-007.md#canonical-dfd367f909005118fe66b891c33f0ecf305b417a73a712f9f57deec2ee2353b8)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](data-sources--voltstack_site--reference--group-007.md#canonical-50dc081106afe105c82c6783368207f74e25b9263381f3fd546e6f75d724505f)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays

<a id="canonical-9147f370adf22ba4b393e6197dd7bf90ecff782b0f0b92e7ce33d19f89a87c0a"></a>

Type: `"list"`. Computed.

For FlashArrays you must set the 'mgmt\_endpoint' and 'api\_token'.

Upstream description:

For FlashArrays you must set the "mgmt\_endpoint" and "api\_token"

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

<a id="canonical-3381e19a5ccf8cc34f1d57127857780e01bf6734fe1d202a09321b3e4b30987e"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / e29490c124e3 / 3

- [api_token](data-sources--voltstack_site--reference--group-007.md#canonical-4543a2e6c962e277692d2aa98c4911ada37d9bbeb5bc1e49812de147b399b551): complete subsection reference.

<a id="canonical-4e9eaf945fb5e538bed3ab7741a5c912fd6c0a1c4091bd6f68f7004d5b4ec167"></a>

<a id="canonical-536b4447f21a3f9bb82a8b9fc170d737275ba768c380c47ef6dac2c8a8822fe6"></a>

## labels property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / e29490c124e3 / 4

Type: `["map", "string"]`. Computed.

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

<a id="canonical-59d19f431d0d7c44b4d538ba25a75564bd38e3abf195e631d2b2b088a49feda4"></a>

<a id="canonical-0ecc2bc46bd346bd1800df6bab893c661b65bca607702ca00cf2060232dd9e0d"></a>

## mgmt_dns_name property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / e29490c124e3 / 5

Type: `"string"`. Computed.

Exclusive with \[mgmt\_ip\] Management Endpoint's IP address is discovered using DNS name
resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[mgmt\_ip\] Management Endpoint's IP address is discovered using DNS name
resolution. The name given here is fully qualified domain name.

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

<a id="canonical-07c438b534f1e13dae456db1a426900ee5cc7112eca8b83f9faf306041a9d367"></a>

<a id="canonical-4fbf77da16ce4193836009df701da29569f7d6fc7033de84b6b6f356384733db"></a>

## mgmt_ip property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / e29490c124e3 / 6

Type: `"string"`. Computed.

Exclusive with \[mgmt\_dns\_name\] Management Endpoint is reachable at the given IP address.

Upstream description:

Exclusive with \[mgmt\_dns\_name\] Management Endpoint is reachable at the given IP address.

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

<a id="canonical-7bc54dc9e18dce13f2d2b8d1f94037ca1fc6843778ed8e4e834c5352e8152096"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / e29490c124e3 / 7

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](data-sources--voltstack_site--reference--group-007.md#canonical-4543a2e6c962e277692d2aa98c4911ada37d9bbeb5bc1e49812de147b399b551)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](data-sources--voltstack_site--reference--group-007.md#canonical-50dc081106afe105c82c6783368207f74e25b9263381f3fd546e6f75d724505f)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-4543a2e6c962e277692d2aa98c4911ada37d9bbeb5bc1e49812de147b399b551"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a892c22f53425d1f7d8ec288771581fa244db4e8770816467c4880584b2f777"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 361750b3061b / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](data-sources--voltstack_site--reference--group-007.md#canonical-678801c40c6f67b9dfaf26614f3fc8e471a101a0988408bcdab7079e531d9a28)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--voltstack_site--reference--group-007.md#canonical-dfd367f909005118fe66b891c33f0ecf305b417a73a712f9f57deec2ee2353b8)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](data-sources--voltstack_site--reference--group-007.md#canonical-50dc081106afe105c82c6783368207f74e25b9263381f3fd546e6f75d724505f)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](data-sources--voltstack_site--reference--group-007.md#canonical-4688a335030d30b7b2c5948ddd2b619f7cd5c4a236cfd6abc58990c7dcfffb07)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token

<a id="canonical-609bed96c41bc83692441aed8bb0683e52d4147c77a48de131cbea803dcc82e0"></a>

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

<a id="canonical-1c1f9019391d9e29baee5f158b8b9bae1a583e45d86fd0a24d856434271d836a"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 361750b3061b / 3

- [blindfold_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-b258621a8a7104c6c1d27fd7be195ff43e28d9f33aefa403b519f2bca9a91add): complete subsection reference.

- [clear_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-22a7ffc541ed678ed35b102bfd58fcd2cb2de6fad8e47025977e6381b98744a8): complete subsection reference.

<a id="canonical-7419c3f4eef64c94aabf293eef06907e3a133ad4b25c4203f9949c200f750bd0"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 361750b3061b / 4

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-b258621a8a7104c6c1d27fd7be195ff43e28d9f33aefa403b519f2bca9a91add)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-22a7ffc541ed678ed35b102bfd58fcd2cb2de6fad8e47025977e6381b98744a8)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](data-sources--voltstack_site--reference--group-007.md#canonical-4688a335030d30b7b2c5948ddd2b619f7cd5c4a236cfd6abc58990c7dcfffb07)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-b258621a8a7104c6c1d27fd7be195ff43e28d9f33aefa403b519f2bca9a91add"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd0b17c8bfcc77250632e511872bdf682b8b1a4bb8e1176fac902b2b928e367d"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / a312e0a843a6 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](data-sources--voltstack_site--reference--group-007.md#canonical-678801c40c6f67b9dfaf26614f3fc8e471a101a0988408bcdab7079e531d9a28)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--voltstack_site--reference--group-007.md#canonical-dfd367f909005118fe66b891c33f0ecf305b417a73a712f9f57deec2ee2353b8)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](data-sources--voltstack_site--reference--group-007.md#canonical-50dc081106afe105c82c6783368207f74e25b9263381f3fd546e6f75d724505f)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](data-sources--voltstack_site--reference--group-007.md#canonical-4688a335030d30b7b2c5948ddd2b619f7cd5c4a236cfd6abc58990c7dcfffb07)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](data-sources--voltstack_site--reference--group-007.md#canonical-4543a2e6c962e277692d2aa98c4911ada37d9bbeb5bc1e49812de147b399b551)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info

<a id="canonical-f26201563acae7551e353bc053e32da9b85ae74ec5b26992202d1378e03c778b"></a>

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

<a id="canonical-3bf9659a7f54ba4c792194a343c19f099e99278efa2faa20b842ce517c02f481"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / a312e0a843a6 / 3

<a id="canonical-53cf90bfdc1221e7537c530dfb4d6f605041d1d7f91e970c6de1d2b5c0ce2dde"></a>

<a id="canonical-c51e5d58bea3f97adf87151b348509cf29f56382dfb129b1128740e6d7d8b778"></a>

## decryption_provider property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / a312e0a843a6 / 4

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

<a id="canonical-3be91a64bae1e56b0af30a0f66f01c9d9260520d756c4f9e298be75f6e1fbf8c"></a>

<a id="canonical-7764fd8d5fe9c8d8b8a2389c02d59b285d278be9c54dbc26a7fd980400bf6423"></a>

## location property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / a312e0a843a6 / 5

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

<a id="canonical-0060370233d4f1e246724177f6df02e7bd394c8823113957485ee657ce4ccd68"></a>

<a id="canonical-8eadb6f9647a898d2c4b576f40896698f0c9e28d31f5ef1e1d6c2d98eb78fe70"></a>

## store_provider property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / a312e0a843a6 / 6

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

<a id="canonical-43b7bd89cc8d783d5bd0f368ce1de60077cda4d51a164af42ff76cad7f67afeb"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / a312e0a843a6 / 7

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](data-sources--voltstack_site--reference--group-007.md#canonical-4543a2e6c962e277692d2aa98c4911ada37d9bbeb5bc1e49812de147b399b551)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-22a7ffc541ed678ed35b102bfd58fcd2cb2de6fad8e47025977e6381b98744a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9739d1f7b1f0abeba8800c1fd1c5059181ee32feb53e8fa2348c48515c32d7d7"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / cf8d356651eb / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](data-sources--voltstack_site--reference--group-007.md#canonical-678801c40c6f67b9dfaf26614f3fc8e471a101a0988408bcdab7079e531d9a28)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--voltstack_site--reference--group-007.md#canonical-dfd367f909005118fe66b891c33f0ecf305b417a73a712f9f57deec2ee2353b8)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](data-sources--voltstack_site--reference--group-007.md#canonical-50dc081106afe105c82c6783368207f74e25b9263381f3fd546e6f75d724505f)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](data-sources--voltstack_site--reference--group-007.md#canonical-4688a335030d30b7b2c5948ddd2b619f7cd5c4a236cfd6abc58990c7dcfffb07)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](data-sources--voltstack_site--reference--group-007.md#canonical-4543a2e6c962e277692d2aa98c4911ada37d9bbeb5bc1e49812de147b399b551)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info

<a id="canonical-4d8db9eddc664b7137e0dc7c66c4ca09c52115acb2e9c00ba4c533dea2d8b64b"></a>

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

<a id="canonical-0ef0e55afad9a13c62920f1c68983fdf74643dd0276a9e4cb51e8df8c8cc985c"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / cf8d356651eb / 3

<a id="canonical-fda752095531bdfc0ab57133b53de03f58273ec91d5901dd54a788ceffc276a7"></a>

<a id="canonical-4ffe6ef8ab502131ef56f31fa4b4b74fcd45916bbf997f90ac928bab94a99d80"></a>

## provider_ref property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / cf8d356651eb / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-43df9b55f5e89d140143bfb69cb621b07764f967df505361fbd7b19884b431c9"></a>

<a id="canonical-41c9404a9fbb3bbc2f8ffba068ecbddf6d97c8da6d254c46dab37c00800f3e75"></a>

## url property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / cf8d356651eb / 5

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

<a id="canonical-e7e92a3c391a3c74284630c134c04d96a07587c66b5965630fd0e71290ab5918"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / cf8d356651eb / 6

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](data-sources--voltstack_site--reference--group-007.md#canonical-4543a2e6c962e277692d2aa98c4911ada37d9bbeb5bc1e49812de147b399b551)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-0760631ea0372ec89b870269224d00e592f5f1eee73801bbb8a1405da724a6fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d704696957010a573f3556808db739ab1dfc268fb7e664380dcfd437eaa8eab8"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / e7e239899d5e / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](data-sources--voltstack_site--reference--group-007.md#canonical-678801c40c6f67b9dfaf26614f3fc8e471a101a0988408bcdab7079e531d9a28)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--voltstack_site--reference--group-007.md#canonical-dfd367f909005118fe66b891c33f0ecf305b417a73a712f9f57deec2ee2353b8)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade

<a id="canonical-ea5fa5487608d2422c471c077ac8878214f382129425a19aba73689d7523a46c"></a>

Type: `"single"`. Computed.

Specify what storage flash blades should be managed the plugin.

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

<a id="canonical-cf8dc6c0b3dbcb9dbdae830c9b994a469dd09732e363c199051b06c99b101409"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / e7e239899d5e / 3

<a id="canonical-b4b92361bf34711e2c6a627d62dc2f986b0502b80c1610040ce84711f39baf93"></a>

<a id="canonical-642872a747eef1d64c8a75bbb6e9f88f2d67b7d6e898793edbc52cb1debbfff3"></a>

## enable_snapshot_directory property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / e7e239899d5e / 4

Type: `"bool"`. Computed.

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

<a id="canonical-58d3053ea2b13c440689c8c36f759819a1951ffe054fff29bc2079c22d108474"></a>

<a id="canonical-afefa6e60e570e1280c4bb205f8882993c36c3dd20a37b4789a9bc00b594493e"></a>

## export_rules property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / e7e239899d5e / 5

Type: `"string"`. Computed.

NFS Export Rules. NFS Export rules.

Upstream description:

NFS Export rules.

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

- [flash_blades](data-sources--voltstack_site--reference--group-007.md#canonical-ab14bae4eb66ac4dd325d120e59c0b2ca0e687ff68b14835bd165d92e41a1505): complete subsection reference.

<a id="canonical-31b9db5164a37fee6c21089b58b6518cdd976fd728c4499924fc0cc3eb093058"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / e7e239899d5e / 6

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](data-sources--voltstack_site--reference--group-007.md#canonical-ab14bae4eb66ac4dd325d120e59c0b2ca0e687ff68b14835bd165d92e41a1505)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--voltstack_site--reference--group-007.md#canonical-dfd367f909005118fe66b891c33f0ecf305b417a73a712f9f57deec2ee2353b8)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-ab14bae4eb66ac4dd325d120e59c0b2ca0e687ff68b14835bd165d92e41a1505"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-01ddc7e3f47a6400aa532b4e7d113d624fe25686efc810dc36a56663369c036f"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 7755466ba08b / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](data-sources--voltstack_site--reference--group-007.md#canonical-678801c40c6f67b9dfaf26614f3fc8e471a101a0988408bcdab7079e531d9a28)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--voltstack_site--reference--group-007.md#canonical-dfd367f909005118fe66b891c33f0ecf305b417a73a712f9f57deec2ee2353b8)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--voltstack_site--reference--group-007.md#canonical-0760631ea0372ec89b870269224d00e592f5f1eee73801bbb8a1405da724a6fc)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades

<a id="canonical-cbf52d977ecf608da890c1b9b8ad45ead4edfabd21c3c2cbdaf69760cfa2b7d2"></a>

Type: `"list"`. Computed.

For FlashBlades you must set the 'mgmt\_endpoint', 'api\_token' and nfs\_endpoint.

Upstream description:

For FlashBlades you must set the "mgmt\_endpoint", "api\_token" and nfs\_endpoint.

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

<a id="canonical-f4b396372f73655944899afa0b9279394a061c1846cb8598ad85d635449234d8"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 7755466ba08b / 3

- [api_token](data-sources--voltstack_site--reference--group-007.md#canonical-a825072dfd958d78837e48b21e5b16d23afc4de4c98cef7de43169d776b00a4d): complete subsection reference.

<a id="canonical-e753e4b77891015afcddd3f8d097d3bfaec27d7bc27d1e136684ef1222fbd1cf"></a>

<a id="canonical-f8689d12b5a77cf7d99e2360e8586424c8f43b96bc4bf07afc87ccec89af86ac"></a>

## labels property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 7755466ba08b / 4

Type: `["map", "string"]`. Computed.

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

<a id="canonical-8d77094f21fd3238a664b96f26a28e05b974ad5b8c0216f9ef44cc2e1310dd0d"></a>

<a id="canonical-b1658b2abe01c3ec267e2c64d30d00d574d42902e32a306e313db5bd9ea5d2ef"></a>

## mgmt_dns_name property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 7755466ba08b / 5

Type: `"string"`. Computed.

Exclusive with \[mgmt\_ip\] Management Endpoint's IP address is discovered using DNS name
resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[mgmt\_ip\] Management Endpoint's IP address is discovered using DNS name
resolution. The name given here is fully qualified domain name.

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

<a id="canonical-3e238fecf7e2eba0d74c98cf165f44c29dbcd040ec86195e5cb3b1c56d0996cb"></a>

<a id="canonical-4fb4b9afa1926da1ad312338248a24a21ed382db18647ef0ef2f86561f9aae53"></a>

## mgmt_ip property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 7755466ba08b / 6

Type: `"string"`. Computed.

Exclusive with \[mgmt\_dns\_name\] Management Endpoint is reachable at the given IP address.

Upstream description:

Exclusive with \[mgmt\_dns\_name\] Management Endpoint is reachable at the given IP address.

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

<a id="canonical-506646709cf96207ba5500b738395603ee3ad592b28a3c5067c2dc208d6defd6"></a>

<a id="canonical-2b97b9b19af5d4614e5f484d0a35bb224ea14632da8751a60dfb43e0450b5fe7"></a>

## nfs_endpoint_dns_name property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 7755466ba08b / 7

Type: `"string"`. Computed.

Exclusive with \[nfs\_endpoint\_ip\] Endpoint's IP address is discovered using DNS name resolution.
The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[nfs\_endpoint\_ip\] Endpoint's IP address is discovered using DNS name resolution.
The name given here is fully qualified domain name.

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

<a id="canonical-1ca56f2ca3c06e1160c3c30ad38d49479ce4c756744dd7afef69e533e2f60596"></a>

<a id="canonical-47adc0b9d9398639dfe2b93504c7affc1be0d4f83951a48f342e713483de3cc4"></a>

## nfs_endpoint_ip property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 7755466ba08b / 8

Type: `"string"`. Computed.

Exclusive with \[nfs\_endpoint\_dns\_name\] Endpoint is reachable at the given IP address.

Upstream description:

Exclusive with \[nfs\_endpoint\_dns\_name\] Endpoint is reachable at the given IP address.

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

<a id="canonical-35b47e266249cea4fa8d3d441a78e90dda7c63de983a90173debf846d52b6f38"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 7755466ba08b / 9

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](data-sources--voltstack_site--reference--group-007.md#canonical-a825072dfd958d78837e48b21e5b16d23afc4de4c98cef7de43169d776b00a4d)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--voltstack_site--reference--group-007.md#canonical-0760631ea0372ec89b870269224d00e592f5f1eee73801bbb8a1405da724a6fc)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-a825072dfd958d78837e48b21e5b16d23afc4de4c98cef7de43169d776b00a4d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-88a7e6e35cb40fac411ad40f541428c85bb38616b597a6089e56c839f9249aae"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / c0eff2cfac51 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](data-sources--voltstack_site--reference--group-007.md#canonical-678801c40c6f67b9dfaf26614f3fc8e471a101a0988408bcdab7079e531d9a28)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--voltstack_site--reference--group-007.md#canonical-dfd367f909005118fe66b891c33f0ecf305b417a73a712f9f57deec2ee2353b8)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--voltstack_site--reference--group-007.md#canonical-0760631ea0372ec89b870269224d00e592f5f1eee73801bbb8a1405da724a6fc)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](data-sources--voltstack_site--reference--group-007.md#canonical-ab14bae4eb66ac4dd325d120e59c0b2ca0e687ff68b14835bd165d92e41a1505)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token

<a id="canonical-2c07f962f16a33d669997d419b55137b6fc122f392bb66509a71b356cb18a897"></a>

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

<a id="canonical-67e3559f3e98e53d6a2bbca581b920cc1bdd8f0b16a0c40d764a03ae7336d13f"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / c0eff2cfac51 / 3

- [blindfold_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-eee8df65cef6474d77118301d10e79e1a8331fe635947fb347c670a6ed43bd1b): complete subsection reference.

- [clear_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-a99a1f107263e588c86a018ee9c8c94227ad6c9bc51ce404a7c811ec950273cd): complete subsection reference.

<a id="canonical-fe797c5cd229b38083a0fd2920216cb9673259411af3e7a57d578d580f35407b"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / c0eff2cfac51 / 4

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-eee8df65cef6474d77118301d10e79e1a8331fe635947fb347c670a6ed43bd1b)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-a99a1f107263e588c86a018ee9c8c94227ad6c9bc51ce404a7c811ec950273cd)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](data-sources--voltstack_site--reference--group-007.md#canonical-ab14bae4eb66ac4dd325d120e59c0b2ca0e687ff68b14835bd165d92e41a1505)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-eee8df65cef6474d77118301d10e79e1a8331fe635947fb347c670a6ed43bd1b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-67c8c35488dce255684d800756e254c03eb8da20e4446cfebd61c54abda18181"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 40dd13437882 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](data-sources--voltstack_site--reference--group-007.md#canonical-678801c40c6f67b9dfaf26614f3fc8e471a101a0988408bcdab7079e531d9a28)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--voltstack_site--reference--group-007.md#canonical-dfd367f909005118fe66b891c33f0ecf305b417a73a712f9f57deec2ee2353b8)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--voltstack_site--reference--group-007.md#canonical-0760631ea0372ec89b870269224d00e592f5f1eee73801bbb8a1405da724a6fc)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](data-sources--voltstack_site--reference--group-007.md#canonical-ab14bae4eb66ac4dd325d120e59c0b2ca0e687ff68b14835bd165d92e41a1505)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](data-sources--voltstack_site--reference--group-007.md#canonical-a825072dfd958d78837e48b21e5b16d23afc4de4c98cef7de43169d776b00a4d)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info

<a id="canonical-c2504516775fce931aefd9d6c6c410feff4aa46ad66c1af2e239b947db8436e1"></a>

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

<a id="canonical-32675611dcbabb1dcfca5042a446a6039814b9673732f6087b35f218ff7c11f1"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 40dd13437882 / 3

<a id="canonical-37b9aad84bfdc75dae3f70fd27993b1271aed5314a5c5c5a99e0328444dd0b07"></a>

<a id="canonical-e2323dc048481e64acce997a8dcd9b4544885b72db949f2fb4de20b17cb53f34"></a>

## decryption_provider property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 40dd13437882 / 4

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

<a id="canonical-94e8db79150681756024d13ba83111ef323b2953f935658a3e057706df9d1edf"></a>

<a id="canonical-862fc6cbbbe424f6b16c7313ffb9c1a3899729aad3a316c93c7e86d6783fe5ed"></a>

## location property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 40dd13437882 / 5

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

<a id="canonical-2e144c8edf93968425e5fbdc1a2d15ce80138cd58518e50501cad5b35c466721"></a>

<a id="canonical-d362051338782696187c00233cb12d17f7bde4f21d0df2abc3b91b14c4f603c2"></a>

## store_provider property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 40dd13437882 / 6

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

<a id="canonical-2095a290b1a544fc7a2ab430de4505971e63b7b3be914e634adb407e40414758"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / 40dd13437882 / 7

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](data-sources--voltstack_site--reference--group-007.md#canonical-a825072dfd958d78837e48b21e5b16d23afc4de4c98cef7de43169d776b00a4d)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-a99a1f107263e588c86a018ee9c8c94227ad6c9bc51ce404a7c811ec950273cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ebfdc7d5c1db5189a49b2370c2f30a305f84fe6e1c33a6cdef7ae45b60adad45"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / e4e1c5440b74 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-a2e6f81bea96936780b598737c778f5249e7284fed18998e5b2674f969c6ba71)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-cf0e2bc809ccd0cf3445f9e0e8418704cf88cec882f3e9d6eb9030937ee29078)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](data-sources--voltstack_site--reference--group-007.md#canonical-678801c40c6f67b9dfaf26614f3fc8e471a101a0988408bcdab7079e531d9a28)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--voltstack_site--reference--group-007.md#canonical-dfd367f909005118fe66b891c33f0ecf305b417a73a712f9f57deec2ee2353b8)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--voltstack_site--reference--group-007.md#canonical-0760631ea0372ec89b870269224d00e592f5f1eee73801bbb8a1405da724a6fc)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](data-sources--voltstack_site--reference--group-007.md#canonical-ab14bae4eb66ac4dd325d120e59c0b2ca0e687ff68b14835bd165d92e41a1505)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](data-sources--voltstack_site--reference--group-007.md#canonical-a825072dfd958d78837e48b21e5b16d23afc4de4c98cef7de43169d776b00a4d)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info

<a id="canonical-390abdaf31b159493a584856c13a38e3f30eea12be42543015bae57c61702a86"></a>

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

<a id="canonical-54b5ef9598c83f349de204c7f7e53ec3028ce65059f708503a4976375a159b84"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / e4e1c5440b74 / 3

<a id="canonical-e4826f03229cb0ce2a954b0364e728d409503d492a97fc7bc81f2aea839d6af8"></a>

<a id="canonical-6138be2875aead85bfda168372e47cbf0936652d9ccfb991a1a179f128413569"></a>

## provider_ref property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / e4e1c5440b74 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-f6b08021132b8c9bf301640edf972cfb694814da458b6ed8960113f2ba737ab2"></a>

<a id="canonical-08e8dadba6a6c8319b0b167edd4270468be005d0705f6ae957973c2efaa16040"></a>

## url property — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / e4e1c5440b74 / 5

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

<a id="canonical-1092bc9f558a1732dc00bdd5bfd2c368427a62c8ef1531418af3f71736951681"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.pure_service_orchestra / e4e1c5440b74 / 6

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](data-sources--voltstack_site--reference--group-007.md#canonical-a825072dfd958d78837e48b21e5b16d23afc4de4c98cef7de43169d776b00a4d)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-3bb13d64cc2c3a0201c344bec6e81817efa179d7afe31ee02ae6e9590a368179"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e7b4522dc06b434aa81961b64dfdcbedc899d9ee931234d84db6db7e109f127d"></a>

## custom_storage_config.storage_interface_list — custom_storage_config.storage_interface_list / 27366abf9fee / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- custom_storage_config.storage_interface_list

<a id="canonical-dfc59468a1c7222439eb363c96d59c31b00a72f23f3a787f37e0c7d0240a61f5"></a>

Type: `"single"`. Computed.

Configure storage interfaces for this App Stack site.

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

<a id="canonical-704c283e831bd6f86d01e44d0fe53139e2f554b59ee82456e3295faaf1dd5b84"></a>

## Direct properties — custom_storage_config.storage_interface_list / 27366abf9fee / 3

- [storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-d1ece71319de1833822bb11752465932960c0e29ccbb4294d3429541f376d15e): complete subsection reference.

<a id="canonical-18964cfe37ba028ec8e90408be528d68cee5ce09c32374cccf1eabbb1fc9f4df"></a>

## Next pages — custom_storage_config.storage_interface_list / 27366abf9fee / 4

- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-d1ece71319de1833822bb11752465932960c0e29ccbb4294d3429541f376d15e)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-d1ece71319de1833822bb11752465932960c0e29ccbb4294d3429541f376d15e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-15d0e0ae6dd76155d5acd49e18872b9154985652c28652258bfdac3b9ab9b7e4"></a>

## custom_storage_config.storage_interface_list.storage_interfaces — custom_storage_config.storage_interface_list.storage_interfaces / 57d548e22ba3 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-3bb13d64cc2c3a0201c344bec6e81817efa179d7afe31ee02ae6e9590a368179)
- custom_storage_config.storage_interface_list.storage_interfaces

<a id="canonical-ff931935102cb343f5afc51abc8b0d203681fc659bd822baf4682da9b5dcf2ab"></a>

Type: `"list"`. Computed.

Configure storage interfaces for this App Stack site.

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

<a id="canonical-7f45dd17a3f3682dbcbe332df67857fc7bf2dafdbe74eddfa2ed6425aa06a428"></a>

## Direct properties — custom_storage_config.storage_interface_list.storage_interfaces / 57d548e22ba3 / 3

<a id="canonical-23f1a2c7cfb53d83e99f0ddb4a2c6ac5bb47c19f303398db4723c468fda4cc90"></a>

<a id="canonical-8d12faf004312c356fd939ef10ed61b981a29ad4f1a805ca034140f5dfc63f53"></a>

## description_spec property — custom_storage_config.storage_interface_list.storage_interfaces / 57d548e22ba3 / 4

Type: `"string"`. Computed.

Interface Description. Description for this Interface.

- [labels](data-sources--voltstack_site--reference--group-007.md#canonical-4a178c472b8e817b8f5b10c5210a244765c46593280b03e1a4055ee3697267f8): complete subsection reference.

- [storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-02945ad86dafeb9323c105007d6add0eb03e0a25c20de0c990d6d8cd8cb3dbd6): complete subsection reference.

<a id="canonical-b155fa251328712edeb0297f5f79ba1b87c36bd53d7e9501abc2e1b0a7ef3eb0"></a>

## Next pages — custom_storage_config.storage_interface_list.storage_interfaces / 57d548e22ba3 / 5

- [custom_storage_config.storage_interface_list.storage_interfaces.labels](data-sources--voltstack_site--reference--group-007.md#canonical-4a178c472b8e817b8f5b10c5210a244765c46593280b03e1a4055ee3697267f8)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-02945ad86dafeb9323c105007d6add0eb03e0a25c20de0c990d6d8cd8cb3dbd6)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-3bb13d64cc2c3a0201c344bec6e81817efa179d7afe31ee02ae6e9590a368179)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-4a178c472b8e817b8f5b10c5210a244765c46593280b03e1a4055ee3697267f8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af2e3372415daf66b5c1c746f2ffb784ab4a233d082be2268c1cea58001d67a3"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.labels — custom_storage_config.storage_interface_list.storage_interfaces.labels / 7b7a790b5951 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-3bb13d64cc2c3a0201c344bec6e81817efa179d7afe31ee02ae6e9590a368179)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-d1ece71319de1833822bb11752465932960c0e29ccbb4294d3429541f376d15e)
- custom_storage_config.storage_interface_list.storage_interfaces.labels

<a id="canonical-54b21753bf9e906724fd05ab15676bdfc884b9e5504fdcc2dc591525f947557c"></a>

Type: `"single"`. Computed.

Add Labels for this Interface, these labels can be used in firewall policy.

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

<a id="canonical-7953a17b20e00f791fce8e376c004303d66622d299ef6cf8c67db7e2eb98188e"></a>

## Direct properties — custom_storage_config.storage_interface_list.storage_interfaces.labels / 7b7a790b5951 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-672133385c8efbf769071f4bf458ea96b41ca728fa9787e7b25d850410095530"></a>

## Next pages — custom_storage_config.storage_interface_list.storage_interfaces.labels / 7b7a790b5951 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-d1ece71319de1833822bb11752465932960c0e29ccbb4294d3429541f376d15e)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-02945ad86dafeb9323c105007d6add0eb03e0a25c20de0c990d6d8cd8cb3dbd6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b6ee55ececbeacd403eb7eb405800a87f33bdbcb12bb2fcd9297881736ee3f15"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / d56b25b6ac16 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-3bb13d64cc2c3a0201c344bec6e81817efa179d7afe31ee02ae6e9590a368179)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-d1ece71319de1833822bb11752465932960c0e29ccbb4294d3429541f376d15e)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface

<a id="canonical-8c77c54cde18bd527f3bf30a3a4137832204d644deef65dd4f589983d25edc4a"></a>

Type: `"single"`. Computed.

Configuration parameter for storage interface.

Upstream description:

Ethernet Interface Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-address_choice": "[\"dhcp_client\",\"dhcp_server\",\"static_ip\"]",
  "x-ves-oneof-field-ipv6_address_choice": "[\"ipv6_auto_config\",\"no_ipv6_address\",\"static_ipv6_address\"]",
  "x-ves-oneof-field-monitoring_choice": "[\"monitor\",\"monitor_disabled\"]",
  "x-ves-oneof-field-network_choice": "[\"segment_network\",\"site_local_inside_network\",\"site_local_network\",\"storage_network\"]",
  "x-ves-oneof-field-node_choice": "[\"cluster\",\"node\"]",
  "x-ves-oneof-field-primary_choice": "[\"is_primary\",\"not_primary\"]",
  "x-ves-oneof-field-vlan_choice": "[\"untagged\",\"vlan_id\"]"
}
```

<a id="canonical-3409c38aba709d2863d8052a0cd736a8d8dd350328432f34d2ca8aafc4a54c36"></a>

## Direct properties — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / d56b25b6ac16 / 3

- [cluster](data-sources--voltstack_site--reference--group-007.md#canonical-4124c3480dda8f7e337120b03cd41e1031bb7e3f12a0306c75da0da7dade79e6): complete subsection reference.

<a id="canonical-47e82772a2352eae4aaf61569ec2ea54ad0dd0cbb17a98944395151d993cac78"></a>

<a id="canonical-1466d813ecec7e42959647162b0037eae6b16ee4170e1426e0ee915d05620f65"></a>

## device property — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / d56b25b6ac16 / 4

Type: `"string"`. Computed.

Interface configuration for the ethernet device.

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

- [dhcp_client](data-sources--voltstack_site--reference--group-007.md#canonical-d85985d167679bc610716d026baa10fcfc217fb7c4fef65c978a73140fc2fb47): complete subsection reference.

- [dhcp_server](data-sources--voltstack_site--reference--group-007.md#canonical-40a7a698285063c8e3f0fa708f954d61a9e0752f6fe77f11e7c558daed7e2923): complete subsection reference.

- [ipv6_auto_config](data-sources--voltstack_site--reference--group-008.md#canonical-2e87518aa381373c6c8317ed2b58ceedea4620c66e8b9793889ffa638256929a): complete subsection reference.

- [is_primary](data-sources--voltstack_site--reference--group-008.md#canonical-6889e79b8104b7668a3cddc1491d98afeaa7b2518c3f6fda8d27801815cdc128): complete subsection reference.

- [monitor](data-sources--voltstack_site--reference--group-008.md#canonical-945b650b0cf856769fb32fcee20a1ecfab5e0ccc87e4db27f734a0f0f9901237): complete subsection reference.

- [monitor_disabled](data-sources--voltstack_site--reference--group-008.md#canonical-8c2f4045b8597a1f56fc3c634c8587b3e737617bed8e12921a43e1fb0d958251): complete subsection reference.

<a id="canonical-6dbb6f0eb8858f7a34af7559380be2d1f8cf3a0f0ef6b4c0017d791f5e08d6b2"></a>

<a id="canonical-c386ff85de3f9ed8918ee165eb4f1e9070993568df14e9f63ed7239457362b2c"></a>

## mtu property — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / d56b25b6ac16 / 5

Type: `"number"`. Computed.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 9000,
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
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  }
}
```

- [no_ipv6_address](data-sources--voltstack_site--reference--group-008.md#canonical-383b64e15302c5299854adf7ed6d5d4f91e5eb8655e9ccecea48de531bb953c5): complete subsection reference.

<a id="canonical-db79b15663759fed5c1400e0c76ee2455192c0813b963aa8a4a281214394c4fc"></a>

<a id="canonical-52eb79f771d41fe626050f84d660725c6500ef418c42e16294653c80057b0f91"></a>

## node property — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / d56b25b6ac16 / 6

Type: `"string"`. Computed.

Exclusive with \[cluster\] Configuration will apply to a device on the given node.

Upstream description:

Exclusive with \[cluster\] Configuration will apply to a device on the given node.

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

- [not_primary](data-sources--voltstack_site--reference--group-008.md#canonical-e26ebd99a1f6f920835c83fa5a8c016cc29a5e46986308bb00cea45b995ea585): complete subsection reference.

<a id="canonical-d434867b48d7ec5a34be44d0615f1d80bad03d1967cb956a507e3ec27976c2d4"></a>

<a id="canonical-99941fee84ce2c01d25d033a3e3ab0451ead24cc11427c4220b0c56c4f18e0c1"></a>

## priority property — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / d56b25b6ac16 / 7

Type: `"number"`. Computed.

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Upstream description:

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

- [site_local_inside_network](data-sources--voltstack_site--reference--group-008.md#canonical-5267ab323f2e8c276c626a25b2d8f10702c4279aefa8844e4b563789496acc67): complete subsection reference.

- [site_local_network](data-sources--voltstack_site--reference--group-008.md#canonical-c4cd284c267be1df886389550014330966e1eb5309ad797a24774da9ff38a2c3): complete subsection reference.

- [static_ip](data-sources--voltstack_site--reference--group-008.md#canonical-814a4505b10246614174d616a0f664425510b1a988c3636cc2d6b562971db69a): complete subsection reference.

- [static_ipv6_address](data-sources--voltstack_site--reference--group-008.md#canonical-21ceeef8ef8d21512e008c78bcadc32265565087dbf7f0ee74c3d321adb0001d): complete subsection reference.

- [storage_network](data-sources--voltstack_site--reference--group-008.md#canonical-e1467f0b1c38b64cb0730dd5738a56912c5137db985b767ba03cb6af20de76fb): complete subsection reference.

- [untagged](data-sources--voltstack_site--reference--group-008.md#canonical-1203cb8c87317dac688ddacd56b5fa37398d0a667a053b562ad5b2a370d36dde): complete subsection reference.

<a id="canonical-a0f223984a16e1c3ee344dba635502e77a0dac24c5feffe8dfdaa5764acf0354"></a>

<a id="canonical-4fb23e7ec3ff372d4acad08be477a6aae8c0c945e37e13c5af143ace44b066f9"></a>

## vlan_id property — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / d56b25b6ac16 / 8

Type: `"number"`. Computed.

Exclusive with \[untagged\] Configure a VLAN tagged ethernet interface.

Upstream description:

Exclusive with \[untagged\] Configure a VLAN tagged ethernet interface.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 4095,
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
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```

<a id="canonical-a57a6ff324e6813be44c353a14b14d9d39f317e2823f57fe20061ee8a169be1c"></a>

## Next pages — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / d56b25b6ac16 / 9

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.cluster](data-sources--voltstack_site--reference--group-007.md#canonical-4124c3480dda8f7e337120b03cd41e1031bb7e3f12a0306c75da0da7dade79e6)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_client](data-sources--voltstack_site--reference--group-007.md#canonical-d85985d167679bc610716d026baa10fcfc217fb7c4fef65c978a73140fc2fb47)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](data-sources--voltstack_site--reference--group-007.md#canonical-40a7a698285063c8e3f0fa708f954d61a9e0752f6fe77f11e7c558daed7e2923)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-008.md#canonical-2e87518aa381373c6c8317ed2b58ceedea4620c66e8b9793889ffa638256929a)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.is_primary](data-sources--voltstack_site--reference--group-008.md#canonical-6889e79b8104b7668a3cddc1491d98afeaa7b2518c3f6fda8d27801815cdc128)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.monitor](data-sources--voltstack_site--reference--group-008.md#canonical-945b650b0cf856769fb32fcee20a1ecfab5e0ccc87e4db27f734a0f0f9901237)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.monitor_disabled](data-sources--voltstack_site--reference--group-008.md#canonical-8c2f4045b8597a1f56fc3c634c8587b3e737617bed8e12921a43e1fb0d958251)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.no_ipv6_address](data-sources--voltstack_site--reference--group-008.md#canonical-383b64e15302c5299854adf7ed6d5d4f91e5eb8655e9ccecea48de531bb953c5)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.not_primary](data-sources--voltstack_site--reference--group-008.md#canonical-e26ebd99a1f6f920835c83fa5a8c016cc29a5e46986308bb00cea45b995ea585)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.site_local_inside_network](data-sources--voltstack_site--reference--group-008.md#canonical-5267ab323f2e8c276c626a25b2d8f10702c4279aefa8844e4b563789496acc67)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.site_local_network](data-sources--voltstack_site--reference--group-008.md#canonical-c4cd284c267be1df886389550014330966e1eb5309ad797a24774da9ff38a2c3)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip](data-sources--voltstack_site--reference--group-008.md#canonical-814a4505b10246614174d616a0f664425510b1a988c3636cc2d6b562971db69a)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ipv6_address](data-sources--voltstack_site--reference--group-008.md#canonical-21ceeef8ef8d21512e008c78bcadc32265565087dbf7f0ee74c3d321adb0001d)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.storage_network](data-sources--voltstack_site--reference--group-008.md#canonical-e1467f0b1c38b64cb0730dd5738a56912c5137db985b767ba03cb6af20de76fb)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.untagged](data-sources--voltstack_site--reference--group-008.md#canonical-1203cb8c87317dac688ddacd56b5fa37398d0a667a053b562ad5b2a370d36dde)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-d1ece71319de1833822bb11752465932960c0e29ccbb4294d3429541f376d15e)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-4124c3480dda8f7e337120b03cd41e1031bb7e3f12a0306c75da0da7dade79e6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b5bba64f0cfedd751b5e866f3b53965d35ae31f0865a01c19370482c951b1325"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.cluster — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / 6b9805b39610 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-3bb13d64cc2c3a0201c344bec6e81817efa179d7afe31ee02ae6e9590a368179)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-d1ece71319de1833822bb11752465932960c0e29ccbb4294d3429541f376d15e)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-02945ad86dafeb9323c105007d6add0eb03e0a25c20de0c990d6d8cd8cb3dbd6)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.cluster

<a id="canonical-99129caecfbc6639bc3523eb9c63d06071dbf78ad643eb5e88137f4b5b43405a"></a>

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

<a id="canonical-b95e1986a4db5dfd6da97d8390982716c7c2ab4b0650ffe0224c00b790455509"></a>

## Direct properties — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / 6b9805b39610 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-47f2be436d7b85e0875292ec3df4889bdeef642ec3796a067bb6c78cae90e269"></a>

## Next pages — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / 6b9805b39610 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-02945ad86dafeb9323c105007d6add0eb03e0a25c20de0c990d6d8cd8cb3dbd6)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-d85985d167679bc610716d026baa10fcfc217fb7c4fef65c978a73140fc2fb47"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8bbc54f5dc56fae39c86a90d035e4606e42f16cf8b8f793eda59a96dff997be1"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_client — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / 533a98d0c79a / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-3bb13d64cc2c3a0201c344bec6e81817efa179d7afe31ee02ae6e9590a368179)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-d1ece71319de1833822bb11752465932960c0e29ccbb4294d3429541f376d15e)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-02945ad86dafeb9323c105007d6add0eb03e0a25c20de0c990d6d8cd8cb3dbd6)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_client

<a id="canonical-21891277218cfdb8c1e9bcadf95a9bdd15133e1af0e684a92ed5f2a866eb5a63"></a>

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

<a id="canonical-4e3a792b1d6fb9bd3a08811e31082743c6cb7c23cd1617a73b84ee32d02c8156"></a>

## Direct properties — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / 533a98d0c79a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f47674b8e55530839d50c667cb8cd4f1a5261c7b4da3e9819b574fccaf6ab3e8"></a>

## Next pages — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / 533a98d0c79a / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-02945ad86dafeb9323c105007d6add0eb03e0a25c20de0c990d6d8cd8cb3dbd6)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-40a7a698285063c8e3f0fa708f954d61a9e0752f6fe77f11e7c558daed7e2923"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7cf39e730c57d6e47720554220a784cf2c4bd9d5172ad96e71933247e71163f6"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / 2a9ce8c54f85 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-3bb13d64cc2c3a0201c344bec6e81817efa179d7afe31ee02ae6e9590a368179)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-d1ece71319de1833822bb11752465932960c0e29ccbb4294d3429541f376d15e)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-02945ad86dafeb9323c105007d6add0eb03e0a25c20de0c990d6d8cd8cb3dbd6)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server

<a id="canonical-031840f1337e06ad637ee45b7f83b0753ac7819e0f4b1c4bae3c62c7ab0da38d"></a>

Type: `"single"`. Computed.

Configuration parameter for dhcp server.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

<a id="canonical-19c7319e835779bf1059e191316b5da141b9d120b8c32125a8da869c10e700cb"></a>

## Direct properties — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / 2a9ce8c54f85 / 3

- [automatic_from_end](data-sources--voltstack_site--reference--group-007.md#canonical-84c8dc85e900732a8aee3be20a2edd68bf89aa2e1690dd73ccf620c17ac5549c): complete subsection reference.

- [automatic_from_start](data-sources--voltstack_site--reference--group-007.md#canonical-bdaf715ded085fd8798b123eca3bf744f02122298aa4193c4199aa7fe951bab1): complete subsection reference.

- [dhcp_networks](data-sources--voltstack_site--reference--group-007.md#canonical-f1df7dbb32787004ecc47d5151cb0479720cf1880781bc551b75b15ed2a914bc): complete subsection reference.

<a id="canonical-401a6734bd46f48feab45e51c35c09ffebea847940e760721f2e045d816e0c5c"></a>

<a id="canonical-38d5d8c8c905f00206096fab210f60a4a68953780e4b43def68608684d619417"></a>

## dhcp_option82_tag property — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / 2a9ce8c54f85 / 4

Type: `"string"`. Computed.

DHCP option 82 tag.

<a id="canonical-aa96e8db7ebc22e01468af85ab23f1f0f0cfd2eb8b8cad414f0b761257330b60"></a>

<a id="canonical-738a7a57c1ace3c2d36489a247285d723bbb65a8949cef65744b411826c42bf8"></a>

## fixed_ip_map property — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / 2a9ce8c54f85 / 5

Type: `["map", "string"]`. Computed.

Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.

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
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  }
}
```

- [interface_ip_map](data-sources--voltstack_site--reference--group-008.md#canonical-3b46c26a70c99b7cf61241d9f3b2a98ef188ecae844d289afe9bd16a24e28075): complete subsection reference.

<a id="canonical-da3cf084c28accb177410db1109d4f7f5c779ffeae849727f80524e91a680d8f"></a>

## Next pages — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / 2a9ce8c54f85 / 6

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.automatic_from_end](data-sources--voltstack_site--reference--group-007.md#canonical-84c8dc85e900732a8aee3be20a2edd68bf89aa2e1690dd73ccf620c17ac5549c)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.automatic_from_start](data-sources--voltstack_site--reference--group-007.md#canonical-bdaf715ded085fd8798b123eca3bf744f02122298aa4193c4199aa7fe951bab1)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks](data-sources--voltstack_site--reference--group-007.md#canonical-f1df7dbb32787004ecc47d5151cb0479720cf1880781bc551b75b15ed2a914bc)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.interface_ip_map](data-sources--voltstack_site--reference--group-008.md#canonical-3b46c26a70c99b7cf61241d9f3b2a98ef188ecae844d289afe9bd16a24e28075)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-02945ad86dafeb9323c105007d6add0eb03e0a25c20de0c990d6d8cd8cb3dbd6)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-84c8dc85e900732a8aee3be20a2edd68bf89aa2e1690dd73ccf620c17ac5549c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b3f519367d43eca64ffe6e3d9342e74b0b5ab0656cb30746e8c25fe6d1e5391"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.automatic_from_end — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / 4432c5eb8c62 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-3bb13d64cc2c3a0201c344bec6e81817efa179d7afe31ee02ae6e9590a368179)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-d1ece71319de1833822bb11752465932960c0e29ccbb4294d3429541f376d15e)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-02945ad86dafeb9323c105007d6add0eb03e0a25c20de0c990d6d8cd8cb3dbd6)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](data-sources--voltstack_site--reference--group-007.md#canonical-40a7a698285063c8e3f0fa708f954d61a9e0752f6fe77f11e7c558daed7e2923)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.automatic_from_end

<a id="canonical-d5dc0ddbffd05d009d7f3e949f3a42a2a4beed2586b75134fccef8650a6e7b96"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for automatic from end.

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

<a id="canonical-57ff35e07ee0e701cc63378c0699a4ca0cf8b32e8bc0f5438c893b29dedce448"></a>

## Direct properties — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / 4432c5eb8c62 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e26966f2ffcc6a6bd1f7e54f05d8714d7898ebde4cd7c074a68807b8fb25e19f"></a>

## Next pages — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / 4432c5eb8c62 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](data-sources--voltstack_site--reference--group-007.md#canonical-40a7a698285063c8e3f0fa708f954d61a9e0752f6fe77f11e7c558daed7e2923)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-bdaf715ded085fd8798b123eca3bf744f02122298aa4193c4199aa7fe951bab1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5dfc781e54f7f130e483ddad04b9e8ff96a0aefddbcc3809529e4bddfc5929a8"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.automatic_from_start — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / c9940c890bc5 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-3bb13d64cc2c3a0201c344bec6e81817efa179d7afe31ee02ae6e9590a368179)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-d1ece71319de1833822bb11752465932960c0e29ccbb4294d3429541f376d15e)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-02945ad86dafeb9323c105007d6add0eb03e0a25c20de0c990d6d8cd8cb3dbd6)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](data-sources--voltstack_site--reference--group-007.md#canonical-40a7a698285063c8e3f0fa708f954d61a9e0752f6fe77f11e7c558daed7e2923)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.automatic_from_start

<a id="canonical-97ef56d82d536e80215c6db748e12ff88c476a8b8f819a9b11fa84e53ca9e182"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for automatic from start.

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

<a id="canonical-c412c9865b77c4285914844c31e10b489f5fabd233d4e6ddafe6d1ed71348bba"></a>

## Direct properties — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / c9940c890bc5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-42d77c793e79cf0606b32ac1dd6e045ecdefe37e40e1f80dd0a305853667bf48"></a>

## Next pages — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / c9940c890bc5 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](data-sources--voltstack_site--reference--group-007.md#canonical-40a7a698285063c8e3f0fa708f954d61a9e0752f6fe77f11e7c558daed7e2923)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-f1df7dbb32787004ecc47d5151cb0479720cf1880781bc551b75b15ed2a914bc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7c0cbee743ef68530d2e340e07e5a13e3758e63b8f836f6648512252ecf6e5bb"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / 241fb36e6f4a / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-3bb13d64cc2c3a0201c344bec6e81817efa179d7afe31ee02ae6e9590a368179)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-d1ece71319de1833822bb11752465932960c0e29ccbb4294d3429541f376d15e)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-02945ad86dafeb9323c105007d6add0eb03e0a25c20de0c990d6d8cd8cb3dbd6)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](data-sources--voltstack_site--reference--group-007.md#canonical-40a7a698285063c8e3f0fa708f954d61a9e0752f6fe77f11e7c558daed7e2923)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks

<a id="canonical-ddd688ea477a7d4084c1e919b90ebfd28fecc5079db44b302c93f5a378cd15bb"></a>

Type: `"list"`. Computed.

List of networks from which DHCP Server can allocate IPv4 Addresses.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-8fe56f1d3a88cf92cb93a3b4af7cd53660e2e3dedbf8014773bcd70061256e93"></a>

## Direct properties — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / 241fb36e6f4a / 3

<a id="canonical-459a31b796cd15cede219f1749ad83fa694229e9d0950f6a8734727e9c69b98f"></a>

<a id="canonical-dcb7b8c1a7324da75f683425aff72864ceb15e2506acdd41c53f255b2d76b88f"></a>

## dgw_address property — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / 241fb36e6f4a / 4

Type: `"string"`. Computed.

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Upstream description:

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

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

<a id="canonical-21cb661470f60a738082c326b49fca32de8603971658cfe06e4f0d2479e6d8ff"></a>

<a id="canonical-6b48cf7796cd5ab481c4348ec60a2c879d1a1929aa92c6b6968676df3dc71790"></a>

## dns_address property — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / 241fb36e6f4a / 5

Type: `"string"`. Computed.

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Upstream description:

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

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

- [first_address](data-sources--voltstack_site--reference--group-007.md#canonical-8c4a5b528204fcd3bd979cc2a391c20ce89069ac11024ec01ea68f3be6563fed): complete subsection reference.

- [last_address](data-sources--voltstack_site--reference--group-007.md#canonical-1032ee107161fea61dd6d722f4a376db9191e6d163c195f39870d23fbfccf9cb): complete subsection reference.

<a id="canonical-c9b73e31ec81e0612b114cfe24816b9e6bc9319c7fbefee703110e08ae5986d2"></a>

<a id="canonical-c0363ee96223446b06d778143d06f26e2b8930f0aebd72f716c0dbfe1cd2fc91"></a>

## network_prefix property — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / 241fb36e6f4a / 6

Type: `"string"`. Computed.

Exclusive with \[\] Set the network prefix for the site. Ex: 192.0.2.0/24.

Upstream description:

Exclusive with \[\] Set the network prefix for the site. Ex: 192.0.2.0/24.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-9feaa567541e0042994493d61add2bad7f7e5b39be25c1942271949beed6af63"></a>

<a id="canonical-0fe9dc227fab38713e3770130c400e574e04201ab029caf4d914d6bf21d05756"></a>

## pool_settings property — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / 241fb36e6f4a / 7

Type: `"string"`. Computed.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Upstream description:

Identifies the how to pick the network for Interface.

Address ranges in DHCP pool list are used for IP Address allocation Address ranges in DHCP pool list
are excluded from IP Address allocation.

Receipt-pinned upstream constraints:

```json
{
  "default": "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
  "enum": [
    "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [pools](data-sources--voltstack_site--reference--group-008.md#canonical-6bd6b40c80d631bc8a78232c15f3e8a03b156cd5aca46c938e60eb38e51717bf): complete subsection reference.

- [same_as_dgw](data-sources--voltstack_site--reference--group-008.md#canonical-67915ac7539beb0ff2b8d7f454a1d15698b2932084f7d5bed1d1767678b37601): complete subsection reference.

<a id="canonical-a5e8072544f26d5b502caa6047a97580e8a44ffbcc9d94a28d4f5a4524c2d405"></a>

## Next pages — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / 241fb36e6f4a / 8

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.first_address](data-sources--voltstack_site--reference--group-007.md#canonical-8c4a5b528204fcd3bd979cc2a391c20ce89069ac11024ec01ea68f3be6563fed)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.last_address](data-sources--voltstack_site--reference--group-007.md#canonical-1032ee107161fea61dd6d722f4a376db9191e6d163c195f39870d23fbfccf9cb)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.pools](data-sources--voltstack_site--reference--group-008.md#canonical-6bd6b40c80d631bc8a78232c15f3e8a03b156cd5aca46c938e60eb38e51717bf)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.same_as_dgw](data-sources--voltstack_site--reference--group-008.md#canonical-67915ac7539beb0ff2b8d7f454a1d15698b2932084f7d5bed1d1767678b37601)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](data-sources--voltstack_site--reference--group-007.md#canonical-40a7a698285063c8e3f0fa708f954d61a9e0752f6fe77f11e7c558daed7e2923)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-8c4a5b528204fcd3bd979cc2a391c20ce89069ac11024ec01ea68f3be6563fed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e57c75dc4dbd013a78c21c12fda153106f15c494299e95f16eb37d721d92faa3"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.first_address — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / f565f12c4960 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-3bb13d64cc2c3a0201c344bec6e81817efa179d7afe31ee02ae6e9590a368179)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-d1ece71319de1833822bb11752465932960c0e29ccbb4294d3429541f376d15e)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-02945ad86dafeb9323c105007d6add0eb03e0a25c20de0c990d6d8cd8cb3dbd6)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](data-sources--voltstack_site--reference--group-007.md#canonical-40a7a698285063c8e3f0fa708f954d61a9e0752f6fe77f11e7c558daed7e2923)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks](data-sources--voltstack_site--reference--group-007.md#canonical-f1df7dbb32787004ecc47d5151cb0479720cf1880781bc551b75b15ed2a914bc)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.first_address

<a id="canonical-c3d9642d06e42f5b5e9f5f3a2ded9118536e28bcd3acf810c0c8459efa6a1ed9"></a>

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

<a id="canonical-ee02b5ae0a6e4352876f1bed89b318ff844f860a1c558f3cf17bd45577133e7d"></a>

## Direct properties — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / f565f12c4960 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4b03210bf70058792da49766008281f04eccc052b541c8610703c0ccbea6c7da"></a>

## Next pages — custom_storage_config.storage_interface_list.storage_interfaces.storage_interfac / f565f12c4960 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks](data-sources--voltstack_site--reference--group-007.md#canonical-f1df7dbb32787004ecc47d5151cb0479720cf1880781bc551b75b15ed2a914bc)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-1032ee107161fea61dd6d722f4a376db9191e6d163c195f39870d23fbfccf9cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
