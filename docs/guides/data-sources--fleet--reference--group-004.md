---
page_title: "xcsh_fleet reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fleet reference."
---

# xcsh_fleet reference

<a id="canonical-a9730d09a0ece5a44bdecbf3b5e9768dc6908eefaabf25c7950eb249cdc7ab48"></a>

## export_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 45ab5d0f3438 / 6

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

- [no_qos](data-sources--fleet--reference--group-004.md#canonical-1613b44bbd0bc4ce6b74104c29c4b39c07395f2f6d6a59690e8c13d2b2071ce4): complete subsection reference.

<a id="canonical-415cb181dd03b47e86268fc60369bdfd339ff8988d8f705578a4e29434b594e6"></a>

<a id="canonical-8115e56447814a12d8dbdbddca78b54fb8139f65840916291d8fef23fe65e981"></a>

## qos_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 45ab5d0f3438 / 7

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

<a id="canonical-bbdd048deb2a6729fb6f68361010e3704988d7028f630cb0507e1991c5440f94"></a>

<a id="canonical-349a1bd5505d8d74b801afc084dd2aaa8a8efa144b2f2ebb51e773a035a71acb"></a>

## security_style property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 45ab5d0f3438 / 8

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

<a id="canonical-3384a803c595d4917e3e8c0cfdacd641a5d4370bfebda2eacc84f445b8708de4"></a>

<a id="canonical-b299bf62f2d86c4790f1432b7610d11ae416452f6865594bc4d6ba9fb6928fe1"></a>

## snapshot_dir property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 45ab5d0f3438 / 9

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

<a id="canonical-bc0b53607bc8e0d0e90ecf01edfee40527f2ae75cc500b642b143f27f920bcd3"></a>

<a id="canonical-87e75ae39bca6620dd5c47e85e88bc0f7dafbfebe109e54f3a62c024726eee63"></a>

## snapshot_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 45ab5d0f3438 / 10

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

<a id="canonical-048cbbda15c8ec27702c3d4807aa2d51b6eada8a63169becd83b050051b0035c"></a>

<a id="canonical-56bdf11814806c359d21c0d2ee29619de696f2c724b412f5d29a379012c36995"></a>

## snapshot_reserve property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 45ab5d0f3438 / 11

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

<a id="canonical-ceb422608f5537b1bdc1300ae21e9346f094e3c2b3c182abdd7cbd1501a5a7ae"></a>

<a id="canonical-2ebe9cc054317a74eb446358f107effe6f4a6d8414bc0c5b5e04cfcc3fe7fe7a"></a>

## space_reserve property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 45ab5d0f3438 / 12

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

<a id="canonical-1eb7996cc2ab180869a1beab1705dc6897d9a2b5a783804b2c1df28ae574f082"></a>

<a id="canonical-a3af8f7b665c46f98b839ec10ad0440ec66ebe8f7b64d332095277a5fa9522fb"></a>

## split_on_clone property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 45ab5d0f3438 / 13

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

<a id="canonical-824567af9e108a784eb1ae716dd84f207fa1d1b77423b9bec1fe4145db352881"></a>

<a id="canonical-1d33221f6a94c9e69cc5247ec0cb537d18b1b0004d3474b73789a549fbe6d6bc"></a>

## tiering_policy property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 45ab5d0f3438 / 14

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

<a id="canonical-ff3a2df2eed6ced21c0a768f42e462eeef4b9a986c7438f42857c18fcf35e29c"></a>

<a id="canonical-5db4339cffd1b05bf805aef2c96ee77273f47caf6ae83b43725d2c60357f6167"></a>

## unix_permissions property — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 45ab5d0f3438 / 15

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

<a id="canonical-baceb24a8c27cdbbcc870e4dd6c56603b118b79460d812285c9aa492174b28d1"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 45ab5d0f3438 / 16

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos](data-sources--fleet--reference--group-004.md#canonical-1613b44bbd0bc4ce6b74104c29c4b39c07395f2f6d6a59690e8c13d2b2071ce4)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-2beb37295fa85b4d6c4c5916001717fb8e9f448241abcb42bdc6b5678df6ed3c)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-1613b44bbd0bc4ce6b74104c29c4b39c07395f2f6d6a59690e8c13d2b2071ce4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-435036467b801c697bf4abeb98d3bc5273be6c6ac0a980746fb29fbcdea66d57"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 3e53745cadf4 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-c9d9219cf96c1684800a42ff589a198fe11cb47c256f095b0e63bc380d3eeb79)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-2beb37295fa85b4d6c4c5916001717fb8e9f448241abcb42bdc6b5678df6ed3c)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults](data-sources--fleet--reference--group-003.md#canonical-36b7100b6badb73b70262a1760b03aec196f91b255873a3ac1ad65e398eec79c)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos

<a id="canonical-5355cc8b078c656808ad8b942c5f86b834bb9ee9f22973415bba5df4d7e8a7e2"></a>

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

<a id="canonical-29b651b6cfc5717cea4b96c6218e989dc328221e1d6b802723ac5af8ab252d49"></a>

## Direct properties — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 3e53745cadf4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6973449b6c6f446cdc5a0b241dd95a5ce584df79e7a3d456b704817d974aa44e"></a>

## Next pages — storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volu / 3e53745cadf4 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults](data-sources--fleet--reference--group-003.md#canonical-36b7100b6badb73b70262a1760b03aec196f91b255873a3ac1ad65e398eec79c)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-2ef4f77b4600ea9dd3b67d64e289b0dbc3ff961652cae751bc762c477e93440a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-06eb55ec7b1aef6264badeea4784faeb1e12286afadfed38a4d9e49bd5481caa"></a>

## storage_device_list.storage_devices.pure_service_orchestrator — storage_device_list.storage_devices.pure_service_orchestrator / a4456b9ec5d3 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- storage_device_list.storage_devices.pure_service_orchestrator

<a id="canonical-d38095c997988289cf8de0b08cf443f354f53ceea912861181847059421dc4ef"></a>

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

<a id="canonical-eb871b80c7c0b530359e5e7cba459633a7285eca1c97a336b43278d0d673176a"></a>

## Direct properties — storage_device_list.storage_devices.pure_service_orchestrator / a4456b9ec5d3 / 3

- [arrays](data-sources--fleet--reference--group-004.md#canonical-8a2606df46d4ad1d8504ee2845965271cebb18069564a29dd24c103bba0437ab): complete subsection reference.

<a id="canonical-913621344da69e5a163d523a77c4e9536f76bffe22213e1622f981a0a88a3b7e"></a>

<a id="canonical-1cb04db2c0fac7b69961e2e439224d2907f78c726c1d067921bf5e22dee0e69a"></a>

## cluster_id property — storage_device_list.storage_devices.pure_service_orchestrator / a4456b9ec5d3 / 4

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

<a id="canonical-90d43d1541a8f6b2a5489ebf6b406e902fb06a54c2f8ff4bf954290cfc24746a"></a>

<a id="canonical-037623133226e7fe0df2d8ff0905699ef48a19373393677a856e1e68b7562870"></a>

## enable_storage_topology property — storage_device_list.storage_devices.pure_service_orchestrator / a4456b9ec5d3 / 5

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

<a id="canonical-7b1cd8fae06b517874d5fafdaed8ed765ac0ae68db5ce20d9a3555247327c477"></a>

<a id="canonical-4f212854e3d0f3930cbae7f0152ccf63a93119469956e42b47adad51930d63b5"></a>

## enable_strict_topology property — storage_device_list.storage_devices.pure_service_orchestrator / a4456b9ec5d3 / 6

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

<a id="canonical-b8fc1c2606d87f3f1a8b2a25fea200722ea663749a5d3b21a4979c0bc8cdc9ad"></a>

## Next pages — storage_device_list.storage_devices.pure_service_orchestrator / a4456b9ec5d3 / 7

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-004.md#canonical-8a2606df46d4ad1d8504ee2845965271cebb18069564a29dd24c103bba0437ab)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-8a2606df46d4ad1d8504ee2845965271cebb18069564a29dd24c103bba0437ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9cb885578dc57478705228b3c957674a44cc67c51495444cd285f5dab2d9c69f"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays — storage_device_list.storage_devices.pure_service_orchestrator.arrays / 7c5a614b8315 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-004.md#canonical-2ef4f77b4600ea9dd3b67d64e289b0dbc3ff961652cae751bc762c477e93440a)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays

<a id="canonical-687fec49a8cb71f4aeb5657dfd9c37ed02b1f7bb311ed2e1fbe2023302676fde"></a>

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

<a id="canonical-61f64f5c3879396bf22ea65ef8a3382f6ec66639bce23eb02f98510bc1cc0d3b"></a>

## Direct properties — storage_device_list.storage_devices.pure_service_orchestrator.arrays / 7c5a614b8315 / 3

- [flash_array](data-sources--fleet--reference--group-004.md#canonical-d0fec4b8a19306527ef7209d6f52d8cdc7ed13968bdea5445dbf70fb2996c837): complete subsection reference.

- [flash_blade](data-sources--fleet--reference--group-004.md#canonical-8eb50b9759fb090eb068ec614aaeab3e38281e373a909a6f5c0a0c9514fdfd73): complete subsection reference.

<a id="canonical-96788a1c2d1d2d057198e6e32301ecb263cfc256468d5637f24304bb830c67c3"></a>

## Next pages — storage_device_list.storage_devices.pure_service_orchestrator.arrays / 7c5a614b8315 / 4

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](data-sources--fleet--reference--group-004.md#canonical-d0fec4b8a19306527ef7209d6f52d8cdc7ed13968bdea5445dbf70fb2996c837)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--fleet--reference--group-004.md#canonical-8eb50b9759fb090eb068ec614aaeab3e38281e373a909a6f5c0a0c9514fdfd73)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-004.md#canonical-2ef4f77b4600ea9dd3b67d64e289b0dbc3ff961652cae751bc762c477e93440a)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-d0fec4b8a19306527ef7209d6f52d8cdc7ed13968bdea5445dbf70fb2996c837"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-38cac062faaefe504cf04d6a28de1b3acc6567b1e5b0146537ebcb1a483008ab"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 8013c91afc93 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-004.md#canonical-2ef4f77b4600ea9dd3b67d64e289b0dbc3ff961652cae751bc762c477e93440a)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-004.md#canonical-8a2606df46d4ad1d8504ee2845965271cebb18069564a29dd24c103bba0437ab)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array

<a id="canonical-3284ab3c06179a7c18d094c630ce5a4e02a611e0325f49b611efc888ab3a55fc"></a>

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

<a id="canonical-d1c33a99398fa0b0f5586ba5cb2e08a32796190b4b3a9cff97c98f8fcbad70d7"></a>

## Direct properties — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 8013c91afc93 / 3

<a id="canonical-9dac598a14ee04a2d24da3e768b5ca5400e2f5002e1db9c40da43266059511c5"></a>

<a id="canonical-8001d408e4876e2eec2a6954586e931c81855aeef88890944152edfddf81b118"></a>

## default_fs_opt property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 8013c91afc93 / 4

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

<a id="canonical-5c59c07ec5a30f88cd3fc19dd5a706d21561aa61c4b38ae1f4c8e550eb8eef35"></a>

<a id="canonical-f9e17217d1804fd628ce0dc7da7df4749231beab9dd92ef5023a1a79145bbd44"></a>

## default_fs_type property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 8013c91afc93 / 5

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

<a id="canonical-95fa16b45525950148b1d2fe15cc170779fa4144c526dff7bc0a3d2aaf395992"></a>

<a id="canonical-8fae544b5e818294a11ba7fe9c7ae73dffa52cbf75b5bdc3ef596b4419494c56"></a>

## default_mount_opts property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 8013c91afc93 / 6

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

<a id="canonical-7624e30f3ceb2c156d729efb63335d3574e25a322182d469f17e4b58c663d092"></a>

<a id="canonical-4202fb6e95b5522b4a63f869c8583f4d4227016a0b34310b4c6261579c2c3129"></a>

## disable_preempt_attachments property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 8013c91afc93 / 7

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

- [flash_arrays](data-sources--fleet--reference--group-004.md#canonical-f2b5bb5c607e803a586c0f6598258c3f36911279f5973b96c6b3f5a3f2878f92): complete subsection reference.

<a id="canonical-d6459b8775e1d1d267ee5936c3f03e5f4b44125f4ed99ebaa899b672b331e7f6"></a>

<a id="canonical-22583d78a1bfeeacd3c26a17963944cae32e494e61fea0cc540f5f171b3fc4ac"></a>

## iscsi_login_timeout property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 8013c91afc93 / 8

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

<a id="canonical-3269b3d030eaa6941d9ee22a10333d7090ec08a33144bc63d6e315e740c713e7"></a>

<a id="canonical-d2d85f98a7ed1a4915b8c86d26c0108a12d6c70698c6eb7bec95f00aa3b4adaa"></a>

## san_type property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 8013c91afc93 / 9

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

<a id="canonical-9e6c68b3218b65a044622249d636a32188bb0530a5580b9e187ffe9f2fce77d5"></a>

## Next pages — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 8013c91afc93 / 10

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](data-sources--fleet--reference--group-004.md#canonical-f2b5bb5c607e803a586c0f6598258c3f36911279f5973b96c6b3f5a3f2878f92)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-004.md#canonical-8a2606df46d4ad1d8504ee2845965271cebb18069564a29dd24c103bba0437ab)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-f2b5bb5c607e803a586c0f6598258c3f36911279f5973b96c6b3f5a3f2878f92"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-136135bed4fdaf5eabd096a201566c84086fb4c09721fcfb9fc4206f799c1448"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 420217a50e61 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-004.md#canonical-2ef4f77b4600ea9dd3b67d64e289b0dbc3ff961652cae751bc762c477e93440a)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-004.md#canonical-8a2606df46d4ad1d8504ee2845965271cebb18069564a29dd24c103bba0437ab)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](data-sources--fleet--reference--group-004.md#canonical-d0fec4b8a19306527ef7209d6f52d8cdc7ed13968bdea5445dbf70fb2996c837)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays

<a id="canonical-d35fb1c9876d40710e26f710093a151eb58387e6e0f1976f04928a316ad95199"></a>

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

<a id="canonical-3f6d9d6649f51500b6ba9e44492ef67a3dc829f87dd04145b5c09cdbdf7b4e6b"></a>

## Direct properties — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 420217a50e61 / 3

- [api_token](data-sources--fleet--reference--group-004.md#canonical-f92e1dfbf3a59e60a024ec45c7162795e8e4b24f8a07e1e4ad186af9eae3c4f9): complete subsection reference.

<a id="canonical-56ab454d8b7b176589494ed39e25ecec05e5e5a33ea80a7637c7d3f5cdff2e8a"></a>

<a id="canonical-6f76fc70287dd2acab1719a4474b84114f2f8f89587ba36cf6f67264031f1753"></a>

## labels property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 420217a50e61 / 4

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

<a id="canonical-2006b920c64d338941cbce93b45be6a52764a0631ddfbf2c4b776ed85c0b8e14"></a>

<a id="canonical-dbfba0fa32c612a925d331fe782237a848126a22634507a8de01a36232c87e12"></a>

## mgmt_dns_name property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 420217a50e61 / 5

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

<a id="canonical-42aa26a44a83d9dd3cf62f0d915dfb2ec26372214a3a33d9e4132d088670fcb9"></a>

<a id="canonical-c78a84414772b466f8872955e9fae42a340dd33c1bd3a21a12a6f2a23323bf50"></a>

## mgmt_ip property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 420217a50e61 / 6

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

<a id="canonical-5b39a1d90d33bcac4f2706b4fe5b7633fe10e46f1aaadbd7b8af005b35f7b30d"></a>

## Next pages — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 420217a50e61 / 7

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](data-sources--fleet--reference--group-004.md#canonical-f92e1dfbf3a59e60a024ec45c7162795e8e4b24f8a07e1e4ad186af9eae3c4f9)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](data-sources--fleet--reference--group-004.md#canonical-d0fec4b8a19306527ef7209d6f52d8cdc7ed13968bdea5445dbf70fb2996c837)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-f92e1dfbf3a59e60a024ec45c7162795e8e4b24f8a07e1e4ad186af9eae3c4f9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5c4ae315c2e2f94926153c1a08bad7369b67809c0b54cf9184d5eb77cfead61"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / ee9e2ab899e8 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-004.md#canonical-2ef4f77b4600ea9dd3b67d64e289b0dbc3ff961652cae751bc762c477e93440a)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-004.md#canonical-8a2606df46d4ad1d8504ee2845965271cebb18069564a29dd24c103bba0437ab)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](data-sources--fleet--reference--group-004.md#canonical-d0fec4b8a19306527ef7209d6f52d8cdc7ed13968bdea5445dbf70fb2996c837)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](data-sources--fleet--reference--group-004.md#canonical-f2b5bb5c607e803a586c0f6598258c3f36911279f5973b96c6b3f5a3f2878f92)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token

<a id="canonical-36a0310de2813554b7b691fa2b3d920b9ad429794001ae147d49c7a8f231f91a"></a>

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

<a id="canonical-642a4ca51cd73beaa319b544fa8c0bd76ed6ffbb02c3196441d4602d0c603013"></a>

## Direct properties — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / ee9e2ab899e8 / 3

- [blindfold_secret_info](data-sources--fleet--reference--group-004.md#canonical-11fe9274d5e677ba7e495f9180b5d241ec64ec911d6a5a36892ff859b4cb3338): complete subsection reference.

- [clear_secret_info](data-sources--fleet--reference--group-004.md#canonical-fca2a879493c42c24062121886d914db208afad67a635dabbe0f554717041d6a): complete subsection reference.

<a id="canonical-70963326010fec503d8f62cc294eb26f01f282156a984f24d1c44bce4a1b4905"></a>

## Next pages — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / ee9e2ab899e8 / 4

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info](data-sources--fleet--reference--group-004.md#canonical-11fe9274d5e677ba7e495f9180b5d241ec64ec911d6a5a36892ff859b4cb3338)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info](data-sources--fleet--reference--group-004.md#canonical-fca2a879493c42c24062121886d914db208afad67a635dabbe0f554717041d6a)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](data-sources--fleet--reference--group-004.md#canonical-f2b5bb5c607e803a586c0f6598258c3f36911279f5973b96c6b3f5a3f2878f92)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-11fe9274d5e677ba7e495f9180b5d241ec64ec911d6a5a36892ff859b4cb3338"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6fe340b184c1d8fa300fcaa39841ec4219afdbcdfcc88b3553134aa1adc71ecb"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 6889d8ae836e / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-004.md#canonical-2ef4f77b4600ea9dd3b67d64e289b0dbc3ff961652cae751bc762c477e93440a)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-004.md#canonical-8a2606df46d4ad1d8504ee2845965271cebb18069564a29dd24c103bba0437ab)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](data-sources--fleet--reference--group-004.md#canonical-d0fec4b8a19306527ef7209d6f52d8cdc7ed13968bdea5445dbf70fb2996c837)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](data-sources--fleet--reference--group-004.md#canonical-f2b5bb5c607e803a586c0f6598258c3f36911279f5973b96c6b3f5a3f2878f92)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](data-sources--fleet--reference--group-004.md#canonical-f92e1dfbf3a59e60a024ec45c7162795e8e4b24f8a07e1e4ad186af9eae3c4f9)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info

<a id="canonical-33fbd312e7b10c6c9273b2106b9ab258e79b42a41fe9335b6c9a4bee52f20b40"></a>

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

<a id="canonical-74a0219a8f493f8c6d5adb457e687e34b907fa26d349d2a1696a58b31276574c"></a>

## Direct properties — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 6889d8ae836e / 3

<a id="canonical-719e9285ada694e2c05a09c6ad02bc474998f2002f79a493d60dc3b04d8cd057"></a>

<a id="canonical-37cc6ecc364ffa082278e1c9aefcc97bbd61699a9f8aaa7e5ef861990bbce481"></a>

## decryption_provider property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 6889d8ae836e / 4

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

<a id="canonical-580ae659087fcd926f09b28baa406364ff1801a1b64faeba81047cf138379efa"></a>

<a id="canonical-79b0557958adcd6c50634ffa4358d5364616f6d01966c9e6f2fc7edbb6c24c28"></a>

## location property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 6889d8ae836e / 5

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

<a id="canonical-56b7ad511c7eccf862b5fa411dbcdb0dd6d0e5b13a85ae2ad6726c6b890684d0"></a>

<a id="canonical-ddbeceb6c28ee18ac0036831d5c11b4594a47868b21e288e34568453f99ebd3c"></a>

## store_provider property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 6889d8ae836e / 6

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

<a id="canonical-659390af2c67b5f477fe5d5ebc720fd739899d1043ec708e44e0b21967181995"></a>

## Next pages — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / 6889d8ae836e / 7

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](data-sources--fleet--reference--group-004.md#canonical-f92e1dfbf3a59e60a024ec45c7162795e8e4b24f8a07e1e4ad186af9eae3c4f9)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-fca2a879493c42c24062121886d914db208afad67a635dabbe0f554717041d6a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-67cc199a8f1efc0c5ca489ec249c507dbf59405f6677bf5d5a5c516d48e7f4c3"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / b2d6f4d5a19b / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-004.md#canonical-2ef4f77b4600ea9dd3b67d64e289b0dbc3ff961652cae751bc762c477e93440a)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-004.md#canonical-8a2606df46d4ad1d8504ee2845965271cebb18069564a29dd24c103bba0437ab)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](data-sources--fleet--reference--group-004.md#canonical-d0fec4b8a19306527ef7209d6f52d8cdc7ed13968bdea5445dbf70fb2996c837)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](data-sources--fleet--reference--group-004.md#canonical-f2b5bb5c607e803a586c0f6598258c3f36911279f5973b96c6b3f5a3f2878f92)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](data-sources--fleet--reference--group-004.md#canonical-f92e1dfbf3a59e60a024ec45c7162795e8e4b24f8a07e1e4ad186af9eae3c4f9)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info

<a id="canonical-091fa3736fa811a166c64c919e44f7f38cbce6bd626b069cc108aead5d038afb"></a>

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

<a id="canonical-f3abfbf4ee38cf55071478afaaf3e4279486a11b77217b50ba7a94b5ec965d14"></a>

## Direct properties — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / b2d6f4d5a19b / 3

<a id="canonical-2f25d7064329fde6bb9cc17d13a2959e91c502d230c232f406fab7728d873e44"></a>

<a id="canonical-526b5756ca8404c81766d6e2f9d7da402d322026568910c707ebd3fd322cf6f7"></a>

## provider_ref property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / b2d6f4d5a19b / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-ed2b77e6b4b70b81d38a8d30d41c3d3b5aeaec1e80322e96361df71c4a1f2600"></a>

<a id="canonical-aea0383540f02a532c4dfd57d0c5da6ebe650b11571d687381aa4e2716789bb0"></a>

## url property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / b2d6f4d5a19b / 5

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

<a id="canonical-c6937c49148ef7275495dddcc43b84bb89e70e03c828127f12e56350cd6f96a2"></a>

## Next pages — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array / b2d6f4d5a19b / 6

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](data-sources--fleet--reference--group-004.md#canonical-f92e1dfbf3a59e60a024ec45c7162795e8e4b24f8a07e1e4ad186af9eae3c4f9)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-8eb50b9759fb090eb068ec614aaeab3e38281e373a909a6f5c0a0c9514fdfd73"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d0f95dcdec3f9b9cf223a90bf430ef0698fa60b27daea1242c604445492871e3"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / 50d5d8852861 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-004.md#canonical-2ef4f77b4600ea9dd3b67d64e289b0dbc3ff961652cae751bc762c477e93440a)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-004.md#canonical-8a2606df46d4ad1d8504ee2845965271cebb18069564a29dd24c103bba0437ab)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade

<a id="canonical-f19b0240ca575bfeeaf4a88ac53bf3b1719f8fe6fd6bc074f5b8d3a9774193ba"></a>

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

<a id="canonical-fb1dd30c61f77c3610f908ff071106424e3595f696fccd32319305310406cb77"></a>

## Direct properties — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / 50d5d8852861 / 3

<a id="canonical-0b56a2189ce5f2004d6ceecbb2ca85d9924a37a162cb89bc48690b050c6dab70"></a>

<a id="canonical-fda891a5819f6d01aabb5e33931484576b13ba873a0160c6ff3e292cf2c8bd3d"></a>

## enable_snapshot_directory property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / 50d5d8852861 / 4

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

<a id="canonical-2f235e49ffbe68631487ad1f6ebdfb694775ecbf3f9f03c21182387b2d6d9144"></a>

<a id="canonical-44352f30ed2e311c505861734f367ae6fca5421c36a2174064876e85effe4346"></a>

## export_rules property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / 50d5d8852861 / 5

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

- [flash_blades](data-sources--fleet--reference--group-004.md#canonical-fc6fe9da35ff322390c5fcf99db84aaf502d700fd7d66e10524be1902a4d2b7c): complete subsection reference.

<a id="canonical-000bd73f17ed95640fba3658024f0302b136b225e8ec2da143dcbe4fa090ff42"></a>

## Next pages — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / 50d5d8852861 / 6

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](data-sources--fleet--reference--group-004.md#canonical-fc6fe9da35ff322390c5fcf99db84aaf502d700fd7d66e10524be1902a4d2b7c)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-004.md#canonical-8a2606df46d4ad1d8504ee2845965271cebb18069564a29dd24c103bba0437ab)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-fc6fe9da35ff322390c5fcf99db84aaf502d700fd7d66e10524be1902a4d2b7c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-491fde04d5ffe64756e1dacd9b1b833dce41b70dd814672c9a053e2a5244f48e"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / 6b850be63cda / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-004.md#canonical-2ef4f77b4600ea9dd3b67d64e289b0dbc3ff961652cae751bc762c477e93440a)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-004.md#canonical-8a2606df46d4ad1d8504ee2845965271cebb18069564a29dd24c103bba0437ab)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--fleet--reference--group-004.md#canonical-8eb50b9759fb090eb068ec614aaeab3e38281e373a909a6f5c0a0c9514fdfd73)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades

<a id="canonical-d7ea3d96b84d2b0ad13cbe369b67739ab86c9fc0a93d5510b3d2284ab9b3f649"></a>

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

<a id="canonical-e9ae2243fcd4dd36415354ec0ba64d749f222fd471c12c6319b391728889f309"></a>

## Direct properties — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / 6b850be63cda / 3

- [api_token](data-sources--fleet--reference--group-004.md#canonical-183768a467b2e3b2e97151095b41d515904483d034279a48061cc239884139ab): complete subsection reference.

<a id="canonical-bf22fa6b63b77bf10530266e79306b032677a3a2e365e9f0ff0b883a15f9af97"></a>

<a id="canonical-eb6eeedeaa228c41305692a96c81329ebc3aafe81bae303b43b117555ff497bf"></a>

## labels property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / 6b850be63cda / 4

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

<a id="canonical-1ab94586f0fa385bc4e01a0a9edbe38ae355e88f1724a091ac9dbc90d474b011"></a>

<a id="canonical-3833524d897d7eedaed13a0fa4c52d9d6f0a6e8f0bc25984c94672816a02c75f"></a>

## mgmt_dns_name property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / 6b850be63cda / 5

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

<a id="canonical-d17dbeb8d64681b610c778ea3766dbeb5e6cebdf9c36f8e418c6cfc2c72b15ce"></a>

<a id="canonical-7200a6c398cac22c3a4dbc567c820953d83e736add943e39a1d9ea8c5fb91684"></a>

## mgmt_ip property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / 6b850be63cda / 6

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

<a id="canonical-dd0c819538ae1d52b71b3065ba04fa63f5f8d7cbc4a64cc788a154d8a928888e"></a>

<a id="canonical-9f3e82ca5ba945dc474ef6a5abc97172fe443a0152bbf742576fe3ad68470509"></a>

## nfs_endpoint_dns_name property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / 6b850be63cda / 7

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

<a id="canonical-6e07797af8f7c3c89962ba6e1fd91e3e8f3b333f0bd511c9301b4f6d4b3e8ade"></a>

<a id="canonical-4210a955cb1f443f3617f4952ce3e3b2fd9d4616aa1c5a1974a1a78b9c19b101"></a>

## nfs_endpoint_ip property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / 6b850be63cda / 8

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

<a id="canonical-fc0ee1c7f04036299ebd29f14c04d7d86b3c0db1f5c5c20d5e12ea952bd182db"></a>

## Next pages — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / 6b850be63cda / 9

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](data-sources--fleet--reference--group-004.md#canonical-183768a467b2e3b2e97151095b41d515904483d034279a48061cc239884139ab)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--fleet--reference--group-004.md#canonical-8eb50b9759fb090eb068ec614aaeab3e38281e373a909a6f5c0a0c9514fdfd73)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-183768a467b2e3b2e97151095b41d515904483d034279a48061cc239884139ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-91d1595aae6515b63c7147032275c2d60fedb3c63d87cd03846339d306d77a44"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / b1a7796442ff / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-004.md#canonical-2ef4f77b4600ea9dd3b67d64e289b0dbc3ff961652cae751bc762c477e93440a)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-004.md#canonical-8a2606df46d4ad1d8504ee2845965271cebb18069564a29dd24c103bba0437ab)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--fleet--reference--group-004.md#canonical-8eb50b9759fb090eb068ec614aaeab3e38281e373a909a6f5c0a0c9514fdfd73)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](data-sources--fleet--reference--group-004.md#canonical-fc6fe9da35ff322390c5fcf99db84aaf502d700fd7d66e10524be1902a4d2b7c)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token

<a id="canonical-093f313fa650cf94494741e2dc1413296c44e3c884b4bac97dc0acf782896481"></a>

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

<a id="canonical-03547055acf18d8a95693253282906c2543bbf463717a3a52ae6ed28eaae9478"></a>

## Direct properties — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / b1a7796442ff / 3

- [blindfold_secret_info](data-sources--fleet--reference--group-004.md#canonical-bf4be4d35404768852f9b7e59f65e40c5088c17dfbc11db8c7176c9d4813567a): complete subsection reference.

- [clear_secret_info](data-sources--fleet--reference--group-004.md#canonical-3a1d94bcad1b7c2356d427a1f9104e276930b5a0091902fc1c2cd8e5baf0daf2): complete subsection reference.

<a id="canonical-3eeb0f23f84b57ea78a544a508c425e4956751f16eaf419f5cad0d12b8457bff"></a>

## Next pages — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / b1a7796442ff / 4

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info](data-sources--fleet--reference--group-004.md#canonical-bf4be4d35404768852f9b7e59f65e40c5088c17dfbc11db8c7176c9d4813567a)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info](data-sources--fleet--reference--group-004.md#canonical-3a1d94bcad1b7c2356d427a1f9104e276930b5a0091902fc1c2cd8e5baf0daf2)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](data-sources--fleet--reference--group-004.md#canonical-fc6fe9da35ff322390c5fcf99db84aaf502d700fd7d66e10524be1902a4d2b7c)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-bf4be4d35404768852f9b7e59f65e40c5088c17dfbc11db8c7176c9d4813567a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d619ddad5446f6c0c580ee62c4eadbb6aae513bb878e622b63c91ecde4090b7"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / cbe38ae1f64a / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-004.md#canonical-2ef4f77b4600ea9dd3b67d64e289b0dbc3ff961652cae751bc762c477e93440a)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-004.md#canonical-8a2606df46d4ad1d8504ee2845965271cebb18069564a29dd24c103bba0437ab)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--fleet--reference--group-004.md#canonical-8eb50b9759fb090eb068ec614aaeab3e38281e373a909a6f5c0a0c9514fdfd73)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](data-sources--fleet--reference--group-004.md#canonical-fc6fe9da35ff322390c5fcf99db84aaf502d700fd7d66e10524be1902a4d2b7c)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](data-sources--fleet--reference--group-004.md#canonical-183768a467b2e3b2e97151095b41d515904483d034279a48061cc239884139ab)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info

<a id="canonical-567c98ef18ef11d9b4e3b5e667e8286e6191ff3acd29bf16e0e6ec127a89c179"></a>

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

<a id="canonical-606a08818adb12f416d094f6c6cb48164e24180ac1d8b0f970cab78ad5b4a1f6"></a>

## Direct properties — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / cbe38ae1f64a / 3

<a id="canonical-e064ec906b3c94bc262c085e834ba6901909125b1975a92b80050754c8700cfc"></a>

<a id="canonical-769fef700ceb72659df7777ef6b759655e813d4e48b4b0cbc2acd607e63f09c6"></a>

## decryption_provider property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / cbe38ae1f64a / 4

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

<a id="canonical-8f2acd56fe05ecd6058694cd664ee2614ddeb97a2dd678e744529f27a77582bd"></a>

<a id="canonical-fb22752443cc5cd9ab55fc1e57018f2041bf3a9a5d1b16e63dabeb44f9ad7955"></a>

## location property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / cbe38ae1f64a / 5

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

<a id="canonical-ff8c17c267495d0772400beb93a78ab961e8573d5e3a8f678696042e0decf00e"></a>

<a id="canonical-a92c20506a6855e2eb4365469d2ed4be51c859a422af5309b072f3f710b6c836"></a>

## store_provider property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / cbe38ae1f64a / 6

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

<a id="canonical-9a8ccce3b04d851dab41455e127183b1e94019d47fa9a862d14037947d521980"></a>

## Next pages — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / cbe38ae1f64a / 7

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](data-sources--fleet--reference--group-004.md#canonical-183768a467b2e3b2e97151095b41d515904483d034279a48061cc239884139ab)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-3a1d94bcad1b7c2356d427a1f9104e276930b5a0091902fc1c2cd8e5baf0daf2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-60fa33e8a8cba1c6f8ac5d2d9c8186066e20162019b7e6e708e0bd66815f59c8"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / b68d07282c11 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-8eae93ee6a1cc89226eee6cbedf164f748b6f24c1fdb2844bc54e90195051725)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-004.md#canonical-2ef4f77b4600ea9dd3b67d64e289b0dbc3ff961652cae751bc762c477e93440a)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-004.md#canonical-8a2606df46d4ad1d8504ee2845965271cebb18069564a29dd24c103bba0437ab)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--fleet--reference--group-004.md#canonical-8eb50b9759fb090eb068ec614aaeab3e38281e373a909a6f5c0a0c9514fdfd73)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](data-sources--fleet--reference--group-004.md#canonical-fc6fe9da35ff322390c5fcf99db84aaf502d700fd7d66e10524be1902a4d2b7c)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](data-sources--fleet--reference--group-004.md#canonical-183768a467b2e3b2e97151095b41d515904483d034279a48061cc239884139ab)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info

<a id="canonical-c3a0a9353de9281ab907ff79b160451a09b47cb3ecec1dc6e486901ca1359e67"></a>

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

<a id="canonical-1e082f553ae8d5b6be04a1d0b52cbd8f34008ad6dfe3631ffcbb37f501a46e69"></a>

## Direct properties — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / b68d07282c11 / 3

<a id="canonical-19b22d9cab256c1f5238d9963ae97425413256a185e4affe83d8af52fc2551e6"></a>

<a id="canonical-c9bf8e525819eff97639e8090f34d6028f7926afeb52558571a00f47158fea88"></a>

## provider_ref property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / b68d07282c11 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2ca48308f55f7da51103df9feb3f399275cc275b95decfb21ed8eb57a34b9a52"></a>

<a id="canonical-aa7aa70daf7641a0aaa9bd6e8808aadd4bd3782580403b07100b706f3d5a918f"></a>

## url property — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / b68d07282c11 / 5

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

<a id="canonical-c057b2005959c1cb5e2a6390098a1222441344bce2254a58074eb839b1ce511f"></a>

## Next pages — storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade / b68d07282c11 / 6

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](data-sources--fleet--reference--group-004.md#canonical-183768a467b2e3b2e97151095b41d515904483d034279a48061cc239884139ab)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-7e877450b08eec03e31927e555048ebdc52fb1f21283e08c037fbbe29d5cd3e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c83084fb895c2e532c72e6d99c8da4b8239695d092c7166b6208bcdd8e3b40af"></a>

## storage_interface_list — storage_interface_list / c0b4f96e457a / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- storage_interface_list

<a id="canonical-99c7687783cc9ee0578f5e98464c24ec87a14e0328f5d92b391257b3b8c422fa"></a>

Type: `"single"`. Computed.

Add all interfaces belonging to this fleet.

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

<a id="canonical-6217aadacb8007535ede1fa28d2f4f910f06b61dd987d53bda6f4f702e57d203"></a>

## Direct properties — storage_interface_list / c0b4f96e457a / 3

- [interfaces](data-sources--fleet--reference--group-004.md#canonical-8fea05909c9f8474ac5657ce93f6b54b20ff622148ca633e2d6de314a36a2ffc): complete subsection reference.

<a id="canonical-e614344d6f87571844943baeec157760902ff848a5b74017ce5ad287591a24ed"></a>

## Next pages — storage_interface_list / c0b4f96e457a / 4

- [storage_interface_list.interfaces](data-sources--fleet--reference--group-004.md#canonical-8fea05909c9f8474ac5657ce93f6b54b20ff622148ca633e2d6de314a36a2ffc)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-8fea05909c9f8474ac5657ce93f6b54b20ff622148ca633e2d6de314a36a2ffc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e357792ba7fba4fda0ec855f38f959a6e43abbb1b29eb2eb2071df1569646a08"></a>

## storage_interface_list.interfaces — storage_interface_list.interfaces / 086cb1c10e4b / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_interface_list](data-sources--fleet--reference--group-004.md#canonical-7e877450b08eec03e31927e555048ebdc52fb1f21283e08c037fbbe29d5cd3e0)
- storage_interface_list.interfaces

<a id="canonical-5a8479f667e67e56bf7050430333044db6f8f077c3dceb0f62f0fc7595cb7e77"></a>

Type: `"list"`. Computed.

Add all interfaces belonging to this fleet.

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

<a id="canonical-c435e6901cd1c004c52ade7a7c68e1b6538d0dab1fe68e8b3dcf3e796c6ad955"></a>

## Direct properties — storage_interface_list.interfaces / 086cb1c10e4b / 3

<a id="canonical-26f92560f8380e6362fb595221fb4e019be9e3e235dc85d2adcb3a99656fc74d"></a>

<a id="canonical-dfe4c96f465b25365b2fe8b9f4a91ac4f46e3df1d364a9d343b23462024c97c2"></a>

## name property — storage_interface_list.interfaces / 086cb1c10e4b / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-f122121084a18cd4107fe5a5babf7fa1b43157308cd9d707a20915973241cfa2"></a>

<a id="canonical-47b2f58123d88c8d5a012a14f0f8395b4d6db9cd448275a17a83bf1274f0ed2d"></a>

## namespace property — storage_interface_list.interfaces / 086cb1c10e4b / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-d1deb8111ed638377a6df50aec6c3c23b69f8ac81b1e4ee4eff232d0eaaaca19"></a>

<a id="canonical-ed89ab3c4a0dc9342081c64cd6bd1338420654fe3b8555575af02c69e7f127b3"></a>

## tenant property — storage_interface_list.interfaces / 086cb1c10e4b / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-51e5a241b54edcc7fb00cdc71b61b450f603a6133195877b5c0bc669d93802b9"></a>

## Next pages — storage_interface_list.interfaces / 086cb1c10e4b / 7

- [storage_interface_list](data-sources--fleet--reference--group-004.md#canonical-7e877450b08eec03e31927e555048ebdc52fb1f21283e08c037fbbe29d5cd3e0)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-9a825351348828d7439fc399c86956cc54202bb679948c2a6cb2e089de057870"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-84c51bf2c5224af9c77bb040f7e5cce8e7748bae9ec6b9126221545ee945901d"></a>

## storage_static_routes — storage_static_routes / 9006fa59d7c1 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- storage_static_routes

<a id="canonical-c6f6fd551b57976de73b52a5c0b44295413498063b2a0254fffa4e58fc42fbe4"></a>

Type: `"single"`. Computed.

Configuration parameter for storage static routes.

Upstream description:

List of storage static routes.

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

<a id="canonical-9d695f482a8b1631c4c34ecde0f3f0d673969649ee89296480001a8b7a755be0"></a>

## Direct properties — storage_static_routes / 9006fa59d7c1 / 3

- [storage_routes](data-sources--fleet--reference--group-004.md#canonical-fce6b494ab5a291450db996ea6b0672148b6621f72f8e6a18c7681f634f0c7d4): complete subsection reference.

<a id="canonical-b0de08ea60dbf71f1962c2fd4563444dfcf0c508cbb49e40ae0c9a4247f2c8a8"></a>

## Next pages — storage_static_routes / 9006fa59d7c1 / 4

- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-fce6b494ab5a291450db996ea6b0672148b6621f72f8e6a18c7681f634f0c7d4)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-fce6b494ab5a291450db996ea6b0672148b6621f72f8e6a18c7681f634f0c7d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ca9e1af37ca038179f84afcaa40705688c509c81f4d818848b666323c83bf25"></a>

## storage_static_routes.storage_routes — storage_static_routes.storage_routes / 9ae8aee31932 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_static_routes](data-sources--fleet--reference--group-004.md#canonical-9a825351348828d7439fc399c86956cc54202bb679948c2a6cb2e089de057870)
- storage_static_routes.storage_routes

<a id="canonical-026769669d10bb75b33fe023235fd8fad766a3522511608e3c2de4355712c1b3"></a>

Type: `"list"`. Computed.

List of Static Routes. List of storage static routes.

Upstream description:

List of storage static routes.

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

<a id="canonical-513cee484a49f7ef43dea38583b1dee5a19049526b20a445afe000be0f5508f2"></a>

## Direct properties — storage_static_routes.storage_routes / 9ae8aee31932 / 3

<a id="canonical-f45e34b928f7726b645d7a89341657f3155379440edd44b9c41099e2bf561461"></a>

<a id="canonical-0786e80ea5a6fd8cb60c00285562f4a72b4e64420a746ab9559ed93b2238def4"></a>

## attrs property — storage_static_routes.storage_routes / 9ae8aee31932 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of route attributes associated with the static route.

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

- [labels](data-sources--fleet--reference--group-004.md#canonical-bfebcc86d5657c4e474eee962d280ab1b7632d1c8811c3b916c079487235e609): complete subsection reference.

- [nexthop](data-sources--fleet--reference--group-004.md#canonical-b0d0ab16363610d70d070ba68e632cc705ff758a3074578037ee31ed883132c0): complete subsection reference.

- [subnets](data-sources--fleet--reference--group-004.md#canonical-e91d12b30d203fcd4c18ee9281ebf7ba3b04e676d6ddd8fb288933b8acaa2533): complete subsection reference.

<a id="canonical-bc0798b4b8b39e340e25d6154d4902462e01a6d1d77247c6788ea995035df9c9"></a>

## Next pages — storage_static_routes.storage_routes / 9ae8aee31932 / 5

- [storage_static_routes.storage_routes.labels](data-sources--fleet--reference--group-004.md#canonical-bfebcc86d5657c4e474eee962d280ab1b7632d1c8811c3b916c079487235e609)
- [storage_static_routes.storage_routes.nexthop](data-sources--fleet--reference--group-004.md#canonical-b0d0ab16363610d70d070ba68e632cc705ff758a3074578037ee31ed883132c0)
- [storage_static_routes.storage_routes.subnets](data-sources--fleet--reference--group-004.md#canonical-e91d12b30d203fcd4c18ee9281ebf7ba3b04e676d6ddd8fb288933b8acaa2533)
- [storage_static_routes](data-sources--fleet--reference--group-004.md#canonical-9a825351348828d7439fc399c86956cc54202bb679948c2a6cb2e089de057870)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-bfebcc86d5657c4e474eee962d280ab1b7632d1c8811c3b916c079487235e609"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2577dfffd00a9a3fe4f0e190c27d480b98f88a34ce28995608d17fa207e126f8"></a>

## storage_static_routes.storage_routes.labels — storage_static_routes.storage_routes.labels / c94ded23feaf / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_static_routes](data-sources--fleet--reference--group-004.md#canonical-9a825351348828d7439fc399c86956cc54202bb679948c2a6cb2e089de057870)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-fce6b494ab5a291450db996ea6b0672148b6621f72f8e6a18c7681f634f0c7d4)
- storage_static_routes.storage_routes.labels

<a id="canonical-d32e86d67d4d43b15082e447e1ced503c7dbe312f2dbacc36e92975626ec447f"></a>

Type: `"single"`. Computed.

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

<a id="canonical-05820fd8e3102d91d7004e11a94b2f7dd15e0b0084e652f0c1a90bc047facfec"></a>

## Direct properties — storage_static_routes.storage_routes.labels / c94ded23feaf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0eeb1017f210271c544fc5a954cfc91f70e80631c3e44a737b61d29042225313"></a>

## Next pages — storage_static_routes.storage_routes.labels / c94ded23feaf / 4

- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-fce6b494ab5a291450db996ea6b0672148b6621f72f8e6a18c7681f634f0c7d4)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-b0d0ab16363610d70d070ba68e632cc705ff758a3074578037ee31ed883132c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ea8d241fb3be2e60c465d426d5274e189c672589be6f2500051fea9cd7161c5d"></a>

## storage_static_routes.storage_routes.nexthop — storage_static_routes.storage_routes.nexthop / 709ed34be5ec / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_static_routes](data-sources--fleet--reference--group-004.md#canonical-9a825351348828d7439fc399c86956cc54202bb679948c2a6cb2e089de057870)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-fce6b494ab5a291450db996ea6b0672148b6621f72f8e6a18c7681f634f0c7d4)
- storage_static_routes.storage_routes.nexthop

<a id="canonical-7e01f9894bb00873031eca7328ca27bbee2f7361be7847fbf16a4537d163cc1e"></a>

Type: `"single"`. Computed.

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

<a id="canonical-539cad7f7683d8b339a08c0c0183d506d80f84d7d57047ff0fddd9a5f71153ea"></a>

## Direct properties — storage_static_routes.storage_routes.nexthop / 709ed34be5ec / 3

- [interface](data-sources--fleet--reference--group-004.md#canonical-c02d921406e126e2bdb1c33d6ba05f929fc51bdcfea0a2f68ea83496db649fb2): complete subsection reference.

- [nexthop_address](data-sources--fleet--reference--group-004.md#canonical-1f5ebbd69fa2dceb7795aa67ea4c40f9e6a0d211fe49c6c71867ca527b7d32b6): complete subsection reference.

<a id="canonical-5c87ee0ea571887a7706c54967563abbccbeafe9e7c7809342c9983d95a6467a"></a>

<a id="canonical-57a5afd0ef8188f1974ed7834f253142503a2a268b8997ff0099e00eb205d033"></a>

## type property — storage_static_routes.storage_routes.nexthop / 709ed34be5ec / 4

Type: `"string"`. Computed.

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

<a id="canonical-c3178a5eb9ba732fac550f3b09d3d648e6b3ad58cb46b1cd19f9fbdb44bdb596"></a>

## Next pages — storage_static_routes.storage_routes.nexthop / 709ed34be5ec / 5

- [storage_static_routes.storage_routes.nexthop.interface](data-sources--fleet--reference--group-004.md#canonical-c02d921406e126e2bdb1c33d6ba05f929fc51bdcfea0a2f68ea83496db649fb2)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](data-sources--fleet--reference--group-004.md#canonical-1f5ebbd69fa2dceb7795aa67ea4c40f9e6a0d211fe49c6c71867ca527b7d32b6)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-fce6b494ab5a291450db996ea6b0672148b6621f72f8e6a18c7681f634f0c7d4)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-c02d921406e126e2bdb1c33d6ba05f929fc51bdcfea0a2f68ea83496db649fb2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-73b771635e49112b4e75e52ef8acd7487de4f899f0712bf34f1c56fd871127be"></a>

## storage_static_routes.storage_routes.nexthop.interface — storage_static_routes.storage_routes.nexthop.interface / 23aa2afb45fb / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_static_routes](data-sources--fleet--reference--group-004.md#canonical-9a825351348828d7439fc399c86956cc54202bb679948c2a6cb2e089de057870)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-fce6b494ab5a291450db996ea6b0672148b6621f72f8e6a18c7681f634f0c7d4)
- [storage_static_routes.storage_routes.nexthop](data-sources--fleet--reference--group-004.md#canonical-b0d0ab16363610d70d070ba68e632cc705ff758a3074578037ee31ed883132c0)
- storage_static_routes.storage_routes.nexthop.interface

<a id="canonical-51e0c0e6936e9fbf04081f776ec1315b6bb4a90f8432b05fe37fcc16560cce85"></a>

Type: `"list"`. Computed.

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

<a id="canonical-2290fecdf13cc4d2da4f940f68795acd09f60e3cd9a4884249f56e1021e1a196"></a>

## Direct properties — storage_static_routes.storage_routes.nexthop.interface / 23aa2afb45fb / 3

<a id="canonical-29f39c9ed37a27c29bca04986139e5f0acfc60bd4f3c483a099334fc8546022c"></a>

<a id="canonical-53932a28a5547b19357807187ac6e00787c593a13c09ff65b4cc15325c697bf1"></a>

## kind property — storage_static_routes.storage_routes.nexthop.interface / 23aa2afb45fb / 4

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

<a id="canonical-a62c66b8f9e5328f61113d953de77b1b46eb06bb325b940802714eec86299e17"></a>

<a id="canonical-665e4c0d41e5ed5ef793f7c8256806a5036b5475f0e76c40ab521604e322e484"></a>

## name property — storage_static_routes.storage_routes.nexthop.interface / 23aa2afb45fb / 5

Type: `"string"`. Computed.

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

<a id="canonical-4cd3cba1af042ab39348ad15755d2a7c0df2444eebc0c4a89fa1d51fa53bb6a0"></a>

<a id="canonical-46c4126ba305ad96fc5e65ed877e1dbbe708e82e6dc06248ab5aab810a0d5045"></a>

## namespace property — storage_static_routes.storage_routes.nexthop.interface / 23aa2afb45fb / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-493278edb190a819f80ad05f91990cc7215ee7380ba10ae85f0a703a8a7b3ca2"></a>

<a id="canonical-51701afd633025e477d50026d978b34e5f92e07ca2e4f27bdec1c35d0e2a7467"></a>

## tenant property — storage_static_routes.storage_routes.nexthop.interface / 23aa2afb45fb / 7

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

<a id="canonical-41d6d21d47076e3407907cdd01c1bbfdcb2d2f856b453aac4448817991ec5619"></a>

<a id="canonical-dc9a6d4a300e5f1eef53cec0d1756113759e83ee35fb1661bb5f87f3739676ce"></a>

## uid property — storage_static_routes.storage_routes.nexthop.interface / 23aa2afb45fb / 8

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

<a id="canonical-cfcc153b637d4791d81e0227f2b98ae0ae18e3d76316a2fc05f34b11b69cdd6d"></a>

## Next pages — storage_static_routes.storage_routes.nexthop.interface / 23aa2afb45fb / 9

- [storage_static_routes.storage_routes.nexthop](data-sources--fleet--reference--group-004.md#canonical-b0d0ab16363610d70d070ba68e632cc705ff758a3074578037ee31ed883132c0)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-1f5ebbd69fa2dceb7795aa67ea4c40f9e6a0d211fe49c6c71867ca527b7d32b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa40ff5fd93680bb9177eabee39cd1365319f334f0f61ccfb217eb0deb56b257"></a>

## storage_static_routes.storage_routes.nexthop.nexthop_address — storage_static_routes.storage_routes.nexthop.nexthop_address / f5023b236346 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_static_routes](data-sources--fleet--reference--group-004.md#canonical-9a825351348828d7439fc399c86956cc54202bb679948c2a6cb2e089de057870)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-fce6b494ab5a291450db996ea6b0672148b6621f72f8e6a18c7681f634f0c7d4)
- [storage_static_routes.storage_routes.nexthop](data-sources--fleet--reference--group-004.md#canonical-b0d0ab16363610d70d070ba68e632cc705ff758a3074578037ee31ed883132c0)
- storage_static_routes.storage_routes.nexthop.nexthop_address

<a id="canonical-3b99b7b056f6e19ee89fbb49c3354b312e568d4adb31f50496d866a2d94fd809"></a>

Type: `"single"`. Computed.

IP Address used to specify an IPv4 or IPv6 address.

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

<a id="canonical-7b37106b54037a44432fe7e58b677a5d575617601c63fe6f901ce4141b7cb0b8"></a>

## Direct properties — storage_static_routes.storage_routes.nexthop.nexthop_address / f5023b236346 / 3

- [dual_stack](data-sources--fleet--reference--group-004.md#canonical-b1adf5b5808a6b2262ebb2f524ac0d18eabfab6475fa5ea5b454b9b6d2969acf): complete subsection reference.

- [ipv4](data-sources--fleet--reference--group-004.md#canonical-63b3cbfc38f4b0ad0668cd50d6dfa62e635eb221cf5993f64d822f61b983c460): complete subsection reference.

- [ipv6](data-sources--fleet--reference--group-004.md#canonical-30fbaa041e0891a77b4750c96ea6abbeeb9bf690f6e4bd59fb86c7eba1f3ffb9): complete subsection reference.

<a id="canonical-b14a8a6b249f37ab306d5dc2088911bd981fc5e2274ec14f4dac709effa5161f"></a>

## Next pages — storage_static_routes.storage_routes.nexthop.nexthop_address / f5023b236346 / 4

- [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack](data-sources--fleet--reference--group-004.md#canonical-b1adf5b5808a6b2262ebb2f524ac0d18eabfab6475fa5ea5b454b9b6d2969acf)
- [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4](data-sources--fleet--reference--group-004.md#canonical-63b3cbfc38f4b0ad0668cd50d6dfa62e635eb221cf5993f64d822f61b983c460)
- [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6](data-sources--fleet--reference--group-004.md#canonical-30fbaa041e0891a77b4750c96ea6abbeeb9bf690f6e4bd59fb86c7eba1f3ffb9)
- [storage_static_routes.storage_routes.nexthop](data-sources--fleet--reference--group-004.md#canonical-b0d0ab16363610d70d070ba68e632cc705ff758a3074578037ee31ed883132c0)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-b1adf5b5808a6b2262ebb2f524ac0d18eabfab6475fa5ea5b454b9b6d2969acf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4e772269a0269dc8f243601089f9058b09a84145bedc99f285bb4d62df4db906"></a>

## storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack — storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack / ee3b12d56d6e / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_static_routes](data-sources--fleet--reference--group-004.md#canonical-9a825351348828d7439fc399c86956cc54202bb679948c2a6cb2e089de057870)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-fce6b494ab5a291450db996ea6b0672148b6621f72f8e6a18c7681f634f0c7d4)
- [storage_static_routes.storage_routes.nexthop](data-sources--fleet--reference--group-004.md#canonical-b0d0ab16363610d70d070ba68e632cc705ff758a3074578037ee31ed883132c0)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](data-sources--fleet--reference--group-004.md#canonical-1f5ebbd69fa2dceb7795aa67ea4c40f9e6a0d211fe49c6c71867ca527b7d32b6)
- storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack

<a id="canonical-5382f0c2697d7807a3384b2271f040d45cbd65b3f53f53496b2bd849a47f3bb6"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1a949eb944cbef4d648c1fb71be4535820bd459dd6407e8050c21ba5c4d936fe"></a>

## Direct properties — storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack / ee3b12d56d6e / 3

- [ipv4](data-sources--fleet--reference--group-004.md#canonical-652ed1a171e2d4007926b23bfe1ba67d726ba9caa6b1f4f83851f645ab9960b2): complete subsection reference.

- [ipv6](data-sources--fleet--reference--group-004.md#canonical-85255487cb51888f3dd89d685bfbb2cc68e8e1eea48c63b864b1fab41821dbec): complete subsection reference.

<a id="canonical-e65737701595ea12fac73d6e1be5b13c86d413a4e6c072666558c3ff568c6c7e"></a>

## Next pages — storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack / ee3b12d56d6e / 4

- [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4](data-sources--fleet--reference--group-004.md#canonical-652ed1a171e2d4007926b23bfe1ba67d726ba9caa6b1f4f83851f645ab9960b2)
- [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6](data-sources--fleet--reference--group-004.md#canonical-85255487cb51888f3dd89d685bfbb2cc68e8e1eea48c63b864b1fab41821dbec)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](data-sources--fleet--reference--group-004.md#canonical-1f5ebbd69fa2dceb7795aa67ea4c40f9e6a0d211fe49c6c71867ca527b7d32b6)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-652ed1a171e2d4007926b23bfe1ba67d726ba9caa6b1f4f83851f645ab9960b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cc3553e77eda680120921dec79816eaa1e4a5012490c5e7712b10134c8dbe6f4"></a>

## storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4 — storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4 / 2fb0f0008fe5 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_static_routes](data-sources--fleet--reference--group-004.md#canonical-9a825351348828d7439fc399c86956cc54202bb679948c2a6cb2e089de057870)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-fce6b494ab5a291450db996ea6b0672148b6621f72f8e6a18c7681f634f0c7d4)
- [storage_static_routes.storage_routes.nexthop](data-sources--fleet--reference--group-004.md#canonical-b0d0ab16363610d70d070ba68e632cc705ff758a3074578037ee31ed883132c0)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](data-sources--fleet--reference--group-004.md#canonical-1f5ebbd69fa2dceb7795aa67ea4c40f9e6a0d211fe49c6c71867ca527b7d32b6)
- [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack](data-sources--fleet--reference--group-004.md#canonical-b1adf5b5808a6b2262ebb2f524ac0d18eabfab6475fa5ea5b454b9b6d2969acf)
- storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4

<a id="canonical-c5d1ae39bd21cb2fbaf4784eb2d615decc78b5c44e0c49bbe4dd7f2832e43acb"></a>

Type: `"single"`. Computed.

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

<a id="canonical-068f6d8eb35fcdb5a68d4f9648217c3913c94e3d28415f05ef9c2148a6d12981"></a>

## Direct properties — storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4 / 2fb0f0008fe5 / 3

<a id="canonical-ccdaa0bea55bf44a3a6876a9d23a27861810178c20dfd6cd202efd51b278e2e4"></a>

<a id="canonical-9ed1973c8e8cd0de1f8446882bed00e708cc8e8cba4186fb53546050f2fdf59c"></a>

## addr property — storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4 / 2fb0f0008fe5 / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-cebe3725d9cdc4afec319040169b3f57ab209424b2bd56ed0144044880e459e2"></a>

## Next pages — storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4 / 2fb0f0008fe5 / 5

- [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack](data-sources--fleet--reference--group-004.md#canonical-b1adf5b5808a6b2262ebb2f524ac0d18eabfab6475fa5ea5b454b9b6d2969acf)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-85255487cb51888f3dd89d685bfbb2cc68e8e1eea48c63b864b1fab41821dbec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c0940fab2525cd06630898bef91d4902292ed09362bc4433e500691190b94883"></a>

## storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6 — storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6 / 037084f7b840 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_static_routes](data-sources--fleet--reference--group-004.md#canonical-9a825351348828d7439fc399c86956cc54202bb679948c2a6cb2e089de057870)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-fce6b494ab5a291450db996ea6b0672148b6621f72f8e6a18c7681f634f0c7d4)
- [storage_static_routes.storage_routes.nexthop](data-sources--fleet--reference--group-004.md#canonical-b0d0ab16363610d70d070ba68e632cc705ff758a3074578037ee31ed883132c0)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](data-sources--fleet--reference--group-004.md#canonical-1f5ebbd69fa2dceb7795aa67ea4c40f9e6a0d211fe49c6c71867ca527b7d32b6)
- [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack](data-sources--fleet--reference--group-004.md#canonical-b1adf5b5808a6b2262ebb2f524ac0d18eabfab6475fa5ea5b454b9b6d2969acf)
- storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6

<a id="canonical-9498970e6251a0128afbe7ef495c744f03d05a1430016d9c5d1a6aed84a3d917"></a>

Type: `"single"`. Computed.

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

<a id="canonical-da28699587fceda7366ba17e81cb6eb2f75eb60b6ca8ca86c11c8257792c1193"></a>

## Direct properties — storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6 / 037084f7b840 / 3

<a id="canonical-0df9ab1bb848e6a383dca33d58ec7f368b2e55fd5943e1a0c0b1094d45ca4a46"></a>

<a id="canonical-77f7503cdd73b4a484f42c6104de6be8624a0592a6ebb1a2d3937962a95544d1"></a>

## addr property — storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6 / 037084f7b840 / 4

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

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

<a id="canonical-2a3918b17a130bb064caedf9623369ec1a9eba3cb0e552c8a1f1151ad68b8a7e"></a>

## Next pages — storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6 / 037084f7b840 / 5

- [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack](data-sources--fleet--reference--group-004.md#canonical-b1adf5b5808a6b2262ebb2f524ac0d18eabfab6475fa5ea5b454b9b6d2969acf)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-63b3cbfc38f4b0ad0668cd50d6dfa62e635eb221cf5993f64d822f61b983c460"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c3824c1b3e2f05bd3fd3417b004ee55ec677de6b821c0bf8c7fed7dbf458217d"></a>

## storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4 — storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4 / e708968bbcf1 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_static_routes](data-sources--fleet--reference--group-004.md#canonical-9a825351348828d7439fc399c86956cc54202bb679948c2a6cb2e089de057870)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-fce6b494ab5a291450db996ea6b0672148b6621f72f8e6a18c7681f634f0c7d4)
- [storage_static_routes.storage_routes.nexthop](data-sources--fleet--reference--group-004.md#canonical-b0d0ab16363610d70d070ba68e632cc705ff758a3074578037ee31ed883132c0)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](data-sources--fleet--reference--group-004.md#canonical-1f5ebbd69fa2dceb7795aa67ea4c40f9e6a0d211fe49c6c71867ca527b7d32b6)
- storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4

<a id="canonical-14f790ac84fb675d9c7bb54a25211b767476a97f9d95a7ed4c11797355c80c9d"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2dfe2b83ae3c6b59cf17b6955f3db5a5542492ad4d69b68ac76f5a2aafe8a5b8"></a>

## Direct properties — storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4 / e708968bbcf1 / 3

<a id="canonical-c7bee82edd4f1914e033d27a24ef39fd218878946607d146f0115563ddadc637"></a>

<a id="canonical-8d25db45bf464dfbab0d1d8a38092ca8167acf74da5bfc8436d24a8fb36bee75"></a>

## addr property — storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4 / e708968bbcf1 / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-d283e39718e5a4bf982ed929a218c01de0e19d3392c4538ea3a4876b00568a9b"></a>

## Next pages — storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4 / e708968bbcf1 / 5

- [storage_static_routes.storage_routes.nexthop.nexthop_address](data-sources--fleet--reference--group-004.md#canonical-1f5ebbd69fa2dceb7795aa67ea4c40f9e6a0d211fe49c6c71867ca527b7d32b6)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-30fbaa041e0891a77b4750c96ea6abbeeb9bf690f6e4bd59fb86c7eba1f3ffb9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a88beba5be871c4b866c56e9fa47dcc77b0f8cfbc12ed5d68b1dce94ce915db8"></a>

## storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6 — storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6 / 153f7f1b44b5 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_static_routes](data-sources--fleet--reference--group-004.md#canonical-9a825351348828d7439fc399c86956cc54202bb679948c2a6cb2e089de057870)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-fce6b494ab5a291450db996ea6b0672148b6621f72f8e6a18c7681f634f0c7d4)
- [storage_static_routes.storage_routes.nexthop](data-sources--fleet--reference--group-004.md#canonical-b0d0ab16363610d70d070ba68e632cc705ff758a3074578037ee31ed883132c0)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](data-sources--fleet--reference--group-004.md#canonical-1f5ebbd69fa2dceb7795aa67ea4c40f9e6a0d211fe49c6c71867ca527b7d32b6)
- storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6

<a id="canonical-a8725eb8320870a87a27715813e4239549470b44728e2c48ea8f3da48dfdbb26"></a>

Type: `"single"`. Computed.

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

<a id="canonical-25726aa5a0761881896f00dcee542fb5aeaf57e56c97b2f949423c6b969223b4"></a>

## Direct properties — storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6 / 153f7f1b44b5 / 3

<a id="canonical-b07d1272621c8b441793939283df1b25b8954b28e2184ac1a83d5133ef74e3ae"></a>

<a id="canonical-c0998cca85a9d86a80f23cc3b1e44ba775776e0391b725ff7895688af0300e4b"></a>

## addr property — storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6 / 153f7f1b44b5 / 4

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

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

<a id="canonical-94c0816dc1f1dc6b3455201b6d24addc5bfd1594d514c4eb8df98417d7a0ee4f"></a>

## Next pages — storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6 / 153f7f1b44b5 / 5

- [storage_static_routes.storage_routes.nexthop.nexthop_address](data-sources--fleet--reference--group-004.md#canonical-1f5ebbd69fa2dceb7795aa67ea4c40f9e6a0d211fe49c6c71867ca527b7d32b6)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-e91d12b30d203fcd4c18ee9281ebf7ba3b04e676d6ddd8fb288933b8acaa2533"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aeb868a889db3b817dd9930f33bf6a84cb5079181772798d368908a248b4cc71"></a>

## storage_static_routes.storage_routes.subnets — storage_static_routes.storage_routes.subnets / 833b9a4da71d / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_static_routes](data-sources--fleet--reference--group-004.md#canonical-9a825351348828d7439fc399c86956cc54202bb679948c2a6cb2e089de057870)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-fce6b494ab5a291450db996ea6b0672148b6621f72f8e6a18c7681f634f0c7d4)
- storage_static_routes.storage_routes.subnets

<a id="canonical-19a2db8fa1618e9888b690dd56d93f654f2f1cfb036ced150387f9ef57ccfa6c"></a>

Type: `"list"`. Computed.

Subnets. List of route prefixes.

Upstream description:

List of route prefixes.

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

<a id="canonical-8b130e9008c8824c07eda1b7f342589df380e15cb5f2a973a45fbbb4d7b1ae2c"></a>

## Direct properties — storage_static_routes.storage_routes.subnets / 833b9a4da71d / 3

- [ipv4](data-sources--fleet--reference--group-004.md#canonical-5e970dedeeaa1be329108cb0fbed0cc6abe665795c64cfb340d5b7e6ff718786): complete subsection reference.

- [ipv6](data-sources--fleet--reference--group-004.md#canonical-750e68defc3cab2d10ec08fd0b7abf1c9d32ba566b18442da3137d4a8473f477): complete subsection reference.

<a id="canonical-77c6d886bc250676ff09c080bc1f87f0ad0e4a5369828c6358be9b849f8baef7"></a>

## Next pages — storage_static_routes.storage_routes.subnets / 833b9a4da71d / 4

- [storage_static_routes.storage_routes.subnets.ipv4](data-sources--fleet--reference--group-004.md#canonical-5e970dedeeaa1be329108cb0fbed0cc6abe665795c64cfb340d5b7e6ff718786)
- [storage_static_routes.storage_routes.subnets.ipv6](data-sources--fleet--reference--group-004.md#canonical-750e68defc3cab2d10ec08fd0b7abf1c9d32ba566b18442da3137d4a8473f477)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-fce6b494ab5a291450db996ea6b0672148b6621f72f8e6a18c7681f634f0c7d4)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-5e970dedeeaa1be329108cb0fbed0cc6abe665795c64cfb340d5b7e6ff718786"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-716bf03369e2c8ef4e65bfdb319fe4a46e9d17d16b272990166232eddf4f42a3"></a>

## storage_static_routes.storage_routes.subnets.ipv4 — storage_static_routes.storage_routes.subnets.ipv4 / 9b5efb26b139 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_static_routes](data-sources--fleet--reference--group-004.md#canonical-9a825351348828d7439fc399c86956cc54202bb679948c2a6cb2e089de057870)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-fce6b494ab5a291450db996ea6b0672148b6621f72f8e6a18c7681f634f0c7d4)
- [storage_static_routes.storage_routes.subnets](data-sources--fleet--reference--group-004.md#canonical-e91d12b30d203fcd4c18ee9281ebf7ba3b04e676d6ddd8fb288933b8acaa2533)
- storage_static_routes.storage_routes.subnets.ipv4

<a id="canonical-d525b1d8572106bf3de79f855c2c8da3c9dcd65c18bbb674bb429388e7696c6f"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3ae9dcd1a7223db51834881036f291a8f1da6971e98edde9b4d77e2e079313f3"></a>

## Direct properties — storage_static_routes.storage_routes.subnets.ipv4 / 9b5efb26b139 / 3

<a id="canonical-4486e4efe62a81e1426cc45a320684bb12eb4e7f4919d7efbcc518597121ac6a"></a>

<a id="canonical-767757b8b4c0ca04bc078461074f63adfdbca9930dc970e52fab6ec769fd977f"></a>

## plen property — storage_static_routes.storage_routes.subnets.ipv4 / 9b5efb26b139 / 4

Type: `"number"`. Computed.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

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

<a id="canonical-fc68b29fc989e4ec222272507c5775276abf6f05892fe6422e62c4ebea8d8561"></a>

<a id="canonical-be62807310d40023b8d9028a37cd2832ee6a3bce6096d304f207ea5602a6d5e5"></a>

## prefix property — storage_static_routes.storage_routes.subnets.ipv4 / 9b5efb26b139 / 5

Type: `"string"`. Computed.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

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

<a id="canonical-9febc6904d4e3b4b8b15720980a5402c304c29e5fa812d471a1df3bd5a716e5f"></a>

## Next pages — storage_static_routes.storage_routes.subnets.ipv4 / 9b5efb26b139 / 6

- [storage_static_routes.storage_routes.subnets](data-sources--fleet--reference--group-004.md#canonical-e91d12b30d203fcd4c18ee9281ebf7ba3b04e676d6ddd8fb288933b8acaa2533)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-750e68defc3cab2d10ec08fd0b7abf1c9d32ba566b18442da3137d4a8473f477"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5edbba9fefcf5ca6685692f42fc2d086e9933324c8f45f1f1cb8790a10a2a8fc"></a>

## storage_static_routes.storage_routes.subnets.ipv6 — storage_static_routes.storage_routes.subnets.ipv6 / 7e8667a930c7 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [storage_static_routes](data-sources--fleet--reference--group-004.md#canonical-9a825351348828d7439fc399c86956cc54202bb679948c2a6cb2e089de057870)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-fce6b494ab5a291450db996ea6b0672148b6621f72f8e6a18c7681f634f0c7d4)
- [storage_static_routes.storage_routes.subnets](data-sources--fleet--reference--group-004.md#canonical-e91d12b30d203fcd4c18ee9281ebf7ba3b04e676d6ddd8fb288933b8acaa2533)
- storage_static_routes.storage_routes.subnets.ipv6

<a id="canonical-7cfad6b7e82a1de6fe312e8fc95d18f7837384a938ea20e4e31d74c4ff3d5207"></a>

Type: `"single"`. Computed.

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

<a id="canonical-fac95c9558c44beff5c1030cf4f34c6768378598254748a1fdf1d214a9de9e2a"></a>

## Direct properties — storage_static_routes.storage_routes.subnets.ipv6 / 7e8667a930c7 / 3

<a id="canonical-ae4e4dad694f726a1b41d0706098ce3741b4a08d382a1a149d424841778ff100"></a>

<a id="canonical-d0660cc8086133fd7099f22551f354b405eedbcaa588b698f96043c0331246f4"></a>

## plen property — storage_static_routes.storage_routes.subnets.ipv6 / 7e8667a930c7 / 4

Type: `"number"`. Computed.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

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

<a id="canonical-56e36ca3092453fe0e1ba96e8698f039e23708530e0ac3864a82dfca89d05869"></a>

<a id="canonical-39a643df987738eedf29c54c988d5b745248213c3c3911e11bdeb85d9d6a041a"></a>

## prefix property — storage_static_routes.storage_routes.subnets.ipv6 / 7e8667a930c7 / 5

Type: `"string"`. Computed.

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. '2001:db8:0:0:0:2:0:0' The address can be compacted by
suppressing zeros e.g. '2001:db8::2::'.

Upstream description:

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. "2001:db8:0:0:0:2:0:0" The address can be compacted by
suppressing zeros e.g. "2001:db8::2::"

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

<a id="canonical-a94f592817e80658f14317a23116b033e59b894f86974e0a64712af8fc0117b1"></a>

## Next pages — storage_static_routes.storage_routes.subnets.ipv6 / 7e8667a930c7 / 6

- [storage_static_routes.storage_routes.subnets](data-sources--fleet--reference--group-004.md#canonical-e91d12b30d203fcd4c18ee9281ebf7ba3b04e676d6ddd8fb288933b8acaa2533)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-9ec18332905de0a2ace95545b134c28e3348331c9500cac4450e177bac5b0c17"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-00dc9bbc7b852fcd69cdd264f1f4dc652ef8a6b50b5153ff263f1ff94242a652"></a>

## usb_policy — usb_policy / 812ed026bab1 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- usb_policy

<a id="canonical-0d5e453c786da8ccd200294830086acbd0edd330738a8ba1b231459c978d547c"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-e478716e7671549e1f5f9be5b83a8d22b66b17d8e11886567266d04a91767a02"></a>

## Direct properties — usb_policy / 812ed026bab1 / 3

<a id="canonical-31a448a51e44954976abbb84a6b8195ae5f90b907b6241c5c8a8ec70de589841"></a>

<a id="canonical-9b90af4638956ef517321dbf195b2823a3ea3b4ac79398ea57e12d77c2782d59"></a>

## name property — usb_policy / 812ed026bab1 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-03a3a8cd97bc12ca874ac3730d1f42b91a3bf1254411d5593ffe3c10b810bc7a"></a>

<a id="canonical-1ec65f212304f7fbe1cafa60c552f442b9fc48292524d857a4c47238dc8544a0"></a>

## namespace property — usb_policy / 812ed026bab1 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-854f34c61d4f3a8b7de4f4840cb265af31481c91d72f819b3aaf9c3e6e8e02c9"></a>

<a id="canonical-a48eaa142f4c1266a1f93d8408712f6b87535c28784405ab26ca27f48b624914"></a>

## tenant property — usb_policy / 812ed026bab1 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-6fb00597971b5fc6182b8faf00c233a4e7c83581489cd544417b6315b41928f8"></a>

## Next pages — usb_policy / 812ed026bab1 / 7

- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
