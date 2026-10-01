---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-57db2089c02cd2ac888d53b831bd186f2019e918baa5f1116d2ca5c69f3c41c8"></a>

## mount_path property — job.volumes.persistent_volume.mount / 0f4553f2ffb3 / 5

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

<a id="canonical-22d61b4b93f645b8d3c25dfcbb6962eb5edcf3160dcb0672e080f62f6cfa65a6"></a>

<a id="canonical-3e80800b379e020ad644d5d1939f7893d1ba6862b88a8dc6cac61bdb53ec53d1"></a>

## sub_path property — job.volumes.persistent_volume.mount / 0f4553f2ffb3 / 6

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

<a id="canonical-f87582d99bfb0d74c35c54a44a7e179d8a0fc8205d5ea4a717e92a2fe96ff467"></a>

## Next pages — job.volumes.persistent_volume.mount / 0f4553f2ffb3 / 7

- [job.volumes.persistent_volume](data-sources--workload--reference--group-004.md#canonical-c5bd53230f9628d4dc4165d5b01d8465447f6c533d8f96b09874ff4c5c548b1a)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-a78bd60bad515eeb71e31c3b93e62800979644e0687e2a4f9122b527e316b13d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1b56ec817b10026a621510f58bb687bb6c2925192ae336c4d0534cc674ee05c"></a>

## job.volumes.persistent_volume.storage — job.volumes.persistent_volume.storage / 9dd2c0a9959a / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [job](data-sources--workload--reference--group-004.md#canonical-7350f554fe4d86d88468d1c9f06e5a182c8c9bdbd89f5cca6ed4489c26853e5d)
- [job.volumes](data-sources--workload--reference--group-004.md#canonical-c9516f3d0a999112413ca9af1373412768f69eaa59a335349812c8eed3413e85)
- [job.volumes.persistent_volume](data-sources--workload--reference--group-004.md#canonical-c5bd53230f9628d4dc4165d5b01d8465447f6c533d8f96b09874ff4c5c548b1a)
- job.volumes.persistent_volume.storage

<a id="canonical-5fdd7ffd1726382e3b44b75eaaee275801ba3b730e371bf2bcf7e9ea1b7561b5"></a>

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

<a id="canonical-0918eae266167cc967d7974b814e502c03690d4cbaefcd1caf86f6d9216960dd"></a>

## Direct properties — job.volumes.persistent_volume.storage / 9dd2c0a9959a / 3

<a id="canonical-cdb2e224fe84100267377aca36a0d007cbe06e2a83f2a2b84d3deefb0590f574"></a>

<a id="canonical-2cfde6b4f261802749c14df4517edb9be98b3ceb290ee3738bb160a898b9c60e"></a>

## access_mode property — job.volumes.persistent_volume.storage / 9dd2c0a9959a / 4

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

<a id="canonical-4ce475ab06ea282de0c0ca4082322213c5109743d0e7863266fe6a0a1e2a98cb"></a>

<a id="canonical-76e7a6bc32f6dea1804ec46bd7cdb27ac2218923ef452f9dce70e8c2fdde7eca"></a>

## class_name property — job.volumes.persistent_volume.storage / 9dd2c0a9959a / 5

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

- [default](data-sources--workload--reference--group-005.md#canonical-cdffea8002c4555b8882b67332ab60a1858c9cc58534c5b156ccf78792b58006): complete subsection reference.

<a id="canonical-ce333fcc5d87faf43f7d17f5c1622307149608f3d294a5a256bab313cb2729b8"></a>

<a id="canonical-888c2a0544af645ad14dd8c31c0f28097c0e8b0d79913e06878f80db632f6cb0"></a>

## storage_size property — job.volumes.persistent_volume.storage / 9dd2c0a9959a / 6

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

<a id="canonical-dfcbf9c4e604d347418d4876157887b9a799e6a62b65b5285a4ed394fe6a657b"></a>

## Next pages — job.volumes.persistent_volume.storage / 9dd2c0a9959a / 7

- [job.volumes.persistent_volume.storage.default](data-sources--workload--reference--group-005.md#canonical-cdffea8002c4555b8882b67332ab60a1858c9cc58534c5b156ccf78792b58006)
- [job.volumes.persistent_volume](data-sources--workload--reference--group-004.md#canonical-c5bd53230f9628d4dc4165d5b01d8465447f6c533d8f96b09874ff4c5c548b1a)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-cdffea8002c4555b8882b67332ab60a1858c9cc58534c5b156ccf78792b58006"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a06e8c03df168da990dcd2061fb8cce5c6d9230395f6d51c9f4b9df6f9567546"></a>

## job.volumes.persistent_volume.storage.default — job.volumes.persistent_volume.storage.default / 73c1459a9a4a / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [job](data-sources--workload--reference--group-004.md#canonical-7350f554fe4d86d88468d1c9f06e5a182c8c9bdbd89f5cca6ed4489c26853e5d)
- [job.volumes](data-sources--workload--reference--group-004.md#canonical-c9516f3d0a999112413ca9af1373412768f69eaa59a335349812c8eed3413e85)
- [job.volumes.persistent_volume](data-sources--workload--reference--group-004.md#canonical-c5bd53230f9628d4dc4165d5b01d8465447f6c533d8f96b09874ff4c5c548b1a)
- [job.volumes.persistent_volume.storage](data-sources--workload--reference--group-005.md#canonical-a78bd60bad515eeb71e31c3b93e62800979644e0687e2a4f9122b527e316b13d)
- job.volumes.persistent_volume.storage.default

<a id="canonical-bdcf79a7cfd7a78fe7b1d65282e43a49cfbb9d535146e46cfe3f96f2fbc2629e"></a>

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

<a id="canonical-472ee7571c73116d9a658da63ce6b13c0c5a59941acc2c1a5364010e5bf9dfbf"></a>

## Direct properties — job.volumes.persistent_volume.storage.default / 73c1459a9a4a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-62804a55b3b073d5b72c8324e1d1615c7a1f23eca1ea1db6bbfce0571b7b5426"></a>

## Next pages — job.volumes.persistent_volume.storage.default / 73c1459a9a4a / 4

- [job.volumes.persistent_volume.storage](data-sources--workload--reference--group-005.md#canonical-a78bd60bad515eeb71e31c3b93e62800979644e0687e2a4f9122b527e316b13d)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e7e755ba1b33c2074f37ffd91afe182c440808f0de89b73b94d769484a382a66"></a>

## service — service / 69f641a8110c / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- service

<a id="canonical-cfdf8b82004eb8aabc075faf46e2cf7f1d3508d8f4c7a6c6d5547380d8755543"></a>

Type: `"single"`. Computed.

Service does not maintain per replica state, however it can be configured to use persistent storage
that is shared amongst all the replicas. Replicas of a service are fungible and do not have a stable
network identity or storage. Common examples of services are web servers, application servers..

Upstream description:

Service does not maintain per replica state, however it can be configured to use persistent storage
that is shared amongst all the replicas. Replicas of a service are fungible and do not have a stable
network identity or storage. Common examples of services are web servers, application servers,
traditional SQL databases, etc.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-scaling_choice": "[\"num_replicas\",\"scale_to_zero\"]"
}
```

<a id="canonical-3f5fa0184d91693c761f11781673d4d0a36d75bcd25a70bf602405c5ff8efd5c"></a>

## Direct properties — service / 69f641a8110c / 3

- [advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff): complete subsection reference.

- [configuration](data-sources--workload--reference--group-015.md#canonical-db082f6c21bd52187fa7fe03d4ad630583fc4dd7bf94198dff5ab8dd3a742d24): complete subsection reference.

- [containers](data-sources--workload--reference--group-015.md#canonical-f3e5d811f1c619d3bb1afde6648794fd06112ea88c6a7458c78327031b9486cf): complete subsection reference.

- [deploy_options](data-sources--workload--reference--group-015.md#canonical-1938114db94be122e9e6f32ed52bcf7ce88fba44423daeb9b9b3cb3545bbc9b8): complete subsection reference.

<a id="canonical-cfe8e97f421a14a4f50fe801ace09cfe6d47a625081b715276d8e750c5d9a286"></a>

<a id="canonical-d6490432f903482198bf0dbd53a786c6d85ce6647a464b4445476a0109b78fcb"></a>

## num_replicas property — service / 69f641a8110c / 4

Type: `"number"`. Computed.

Exclusive with \[scale\_to\_zero\] Number of replicas of service to spawn per site.

Upstream description:

Exclusive with \[scale\_to\_zero\] Number of replicas of service to spawn per site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gt": "0",
    "ves.io.schema.rules.int32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gt": "0",
    "ves.io.schema.rules.int32.lte": "5"
  }
}
```

- [scale_to_zero](data-sources--workload--reference--group-015.md#canonical-6ebab7e2e2f069d9b9109fa69dd5907985757d4d36cc6e56ff9fd00e40411269): complete subsection reference.

- [volumes](data-sources--workload--reference--group-015.md#canonical-8650f561bb6addd9f137c4e4b67ba52ffb9efd5338ba2a38ed06a7df7fb0e7b4): complete subsection reference.

<a id="canonical-26f1495ba1ca161c8fc794f18e62e146dcc0d89799ff43b1b886b9118590d60b"></a>

## Next pages — service / 69f641a8110c / 5

- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.configuration](data-sources--workload--reference--group-015.md#canonical-db082f6c21bd52187fa7fe03d4ad630583fc4dd7bf94198dff5ab8dd3a742d24)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-f3e5d811f1c619d3bb1afde6648794fd06112ea88c6a7458c78327031b9486cf)
- [service.deploy_options](data-sources--workload--reference--group-015.md#canonical-1938114db94be122e9e6f32ed52bcf7ce88fba44423daeb9b9b3cb3545bbc9b8)
- [service.scale_to_zero](data-sources--workload--reference--group-015.md#canonical-6ebab7e2e2f069d9b9109fa69dd5907985757d4d36cc6e56ff9fd00e40411269)
- [service.volumes](data-sources--workload--reference--group-015.md#canonical-8650f561bb6addd9f137c4e4b67ba52ffb9efd5338ba2a38ed06a7df7fb0e7b4)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ff28ead1caea5ee37c87c383674df896c265eb11ababd254434638612b9e5ea5"></a>

## service.advertise_options — service.advertise_options / 7c78f144a404 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- service.advertise_options

<a id="canonical-9559a6c4f69b3970c550db809a1fd64b320c9a7610e7411e9832b2d1bbe1d759"></a>

Type: `"single"`. Computed.

Advertise OPTIONS are used to configure how and where to advertise the workload using load
balancers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-advertise_choice": "[\"advertise_custom\",\"advertise_in_cluster\",\"advertise_on_public\",\"do_not_advertise\"]"
}
```

<a id="canonical-a9c4a9e60f8d67f297ec97bc2d08b205b60c3d5f49f684e8e6ee1e7f06ca8d5c"></a>

## Direct properties — service.advertise_options / 7c78f144a404 / 3

- [advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62): complete subsection reference.

- [advertise_in_cluster](data-sources--workload--reference--group-008.md#canonical-9d22595127cd5d0e2ed99473331e76f716d729f28680fb396db73663194f15db): complete subsection reference.

- [advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c): complete subsection reference.

- [do_not_advertise](data-sources--workload--reference--group-015.md#canonical-f3f8ccca96d163616e3cb5e7b42b4d261cfcc700a0e926fbe177a14cd8f2c3c0): complete subsection reference.

<a id="canonical-902bebe47ebd0b64d6d6a3499d33f482c7ba767a9efb7d249caf15474c75bc3e"></a>

## Next pages — service.advertise_options / 7c78f144a404 / 4

- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_in_cluster](data-sources--workload--reference--group-008.md#canonical-9d22595127cd5d0e2ed99473331e76f716d729f28680fb396db73663194f15db)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.do_not_advertise](data-sources--workload--reference--group-015.md#canonical-f3f8ccca96d163616e3cb5e7b42b4d261cfcc700a0e926fbe177a14cd8f2c3c0)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee05664f677082171d9ead2441c7b4ef8fda9116af86442a742bce1004c3077c"></a>

## service.advertise_options.advertise_custom — service.advertise_options.advertise_custom / bf74d9afe125 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- service.advertise_options.advertise_custom

<a id="canonical-380b3cef5ffcdca0d600279342deffc60c8ed241e6d442336ab5810caba7194f"></a>

Type: `"single"`. Computed.

Advertise this workload via loadbalancer on specific sites.

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

<a id="canonical-ebc5d1e095cca47f1da273c5faee23adbac1a5170cfacb69e7534108155ad722"></a>

## Direct properties — service.advertise_options.advertise_custom / bf74d9afe125 / 3

- [advertise_where](data-sources--workload--reference--group-005.md#canonical-c828ef9397b96106571b5a7e48d0b7869905751a1f38a48eaf6662fdd150fd18): complete subsection reference.

- [ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912): complete subsection reference.

<a id="canonical-1b262d95f1fa7305e9e7a6ff15ebc5163888296c9dbbb918a0f28c69e7d4b057"></a>

## Next pages — service.advertise_options.advertise_custom / bf74d9afe125 / 4

- [service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-005.md#canonical-c828ef9397b96106571b5a7e48d0b7869905751a1f38a48eaf6662fdd150fd18)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-c828ef9397b96106571b5a7e48d0b7869905751a1f38a48eaf6662fdd150fd18"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a0782e7d2816cd5366f947cd75f1bd6c019a6271484c780eece618927c1f1668"></a>

## service.advertise_options.advertise_custom.advertise_where — service.advertise_options.advertise_custom.advertise_where / c153723054d7 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- service.advertise_options.advertise_custom.advertise_where

<a id="canonical-1e349faef4ac13d3a131cfe872c67b52d2e64770c8adfbfbe6edcff8da1bcfd6"></a>

Type: `"list"`. Computed.

Where should this load balancer be available.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-d70fb863710273f6b0f5f1f49ba74d3945709ebe2c11f18626b8c3e9f9a37351"></a>

## Direct properties — service.advertise_options.advertise_custom.advertise_where / c153723054d7 / 3

- [site](data-sources--workload--reference--group-005.md#canonical-241cb3a4738269e617f04ac2d9e3ed069f941593b2ed290de51857611664eace): complete subsection reference.

- [virtual_site](data-sources--workload--reference--group-005.md#canonical-966d0540fd2f7f482faff0251e2f54efa1f1e1d466df34e64f1e06c0cf596885): complete subsection reference.

- [vk8s_service](data-sources--workload--reference--group-005.md#canonical-ae3bbda4789096d2bb6def8889bf556d43025f52cee8343c45f3387a7374a178): complete subsection reference.

<a id="canonical-bb2e8ab3840c93db13ffadb06605c5f0b546e9c7c4ee82e0e2241f9bbbe22a3f"></a>

## Next pages — service.advertise_options.advertise_custom.advertise_where / c153723054d7 / 4

- [service.advertise_options.advertise_custom.advertise_where.site](data-sources--workload--reference--group-005.md#canonical-241cb3a4738269e617f04ac2d9e3ed069f941593b2ed290de51857611664eace)
- [service.advertise_options.advertise_custom.advertise_where.virtual_site](data-sources--workload--reference--group-005.md#canonical-966d0540fd2f7f482faff0251e2f54efa1f1e1d466df34e64f1e06c0cf596885)
- [service.advertise_options.advertise_custom.advertise_where.vk8s_service](data-sources--workload--reference--group-005.md#canonical-ae3bbda4789096d2bb6def8889bf556d43025f52cee8343c45f3387a7374a178)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-241cb3a4738269e617f04ac2d9e3ed069f941593b2ed290de51857611664eace"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-251511f2d94319530a75a2350abae0a9bc8b1ee6ea1ced82eeb53ebb722e51dc"></a>

## service.advertise_options.advertise_custom.advertise_where.site — service.advertise_options.advertise_custom.advertise_where.site / 5180561cfe72 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-005.md#canonical-c828ef9397b96106571b5a7e48d0b7869905751a1f38a48eaf6662fdd150fd18)
- service.advertise_options.advertise_custom.advertise_where.site

<a id="canonical-873703e8aa9711cc2ac6d9446f440f7f294de6424774f5561633df4506a09a64"></a>

Type: `"single"`. Computed.

Defines a reference to a CE site along with network type and an optional IP address where a load
balancer could be advertised.

Upstream description:

This defines a reference to a CE site along with network type and an optional IP address where a
load balancer could be advertised.

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

<a id="canonical-246b3ff9ac3c1da52e5a6a1319306cd83584fd56b28ff717f0d66d1620236571"></a>

## Direct properties — service.advertise_options.advertise_custom.advertise_where.site / 5180561cfe72 / 3

<a id="canonical-d368a54b1c328e1f6fd0af664612d40d4765773130f03fe29ef982b885550150"></a>

<a id="canonical-1481815762e744a5cc3281e71d59b289184b0d7378cba675b8a1165860b76f6b"></a>

## ip property — service.advertise_options.advertise_custom.advertise_where.site / 5180561cfe72 / 4

Type: `"string"`. Computed.

Use given IP address as VIP on the site.

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

<a id="canonical-67c25349acdbeac9a4a1907e7a1a78c39d8aca012cbd21ae41a8c0dde1337086"></a>

<a id="canonical-9d0080f15958b2b3ff5a78bb7c4bd2cca392db3d06fade545e023dee20010dd8"></a>

## network property — service.advertise_options.advertise_custom.advertise_where.site / 5180561cfe72 / 5

Type: `"string"`. Computed.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Upstream description:

This defines network types to be used on site

All inside and outside networks. All inside and outside networks with internet VIP support. All
inside networks. All outside networks. All outside networks with internet VIP support. VK8s service
network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [site](data-sources--workload--reference--group-005.md#canonical-1ca45b563f22d2330fe47ae8d062a87db8f8975d42ac179aac18f1facccfb10e): complete subsection reference.

<a id="canonical-5f167331b8189d3b1b3ab51cfc4c4e91f0ab29a262474f0e594c6c62a5b25bbf"></a>

## Next pages — service.advertise_options.advertise_custom.advertise_where.site / 5180561cfe72 / 6

- [service.advertise_options.advertise_custom.advertise_where.site.site](data-sources--workload--reference--group-005.md#canonical-1ca45b563f22d2330fe47ae8d062a87db8f8975d42ac179aac18f1facccfb10e)
- [service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-005.md#canonical-c828ef9397b96106571b5a7e48d0b7869905751a1f38a48eaf6662fdd150fd18)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-1ca45b563f22d2330fe47ae8d062a87db8f8975d42ac179aac18f1facccfb10e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4dd13f78ca2bf5191e42735252d7257be090eebafa85d9b1de3f220f20d35cad"></a>

## service.advertise_options.advertise_custom.advertise_where.site.site — service.advertise_options.advertise_custom.advertise_where.site.site / 3a9443faf44f / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-005.md#canonical-c828ef9397b96106571b5a7e48d0b7869905751a1f38a48eaf6662fdd150fd18)
- [service.advertise_options.advertise_custom.advertise_where.site](data-sources--workload--reference--group-005.md#canonical-241cb3a4738269e617f04ac2d9e3ed069f941593b2ed290de51857611664eace)
- service.advertise_options.advertise_custom.advertise_where.site.site

<a id="canonical-7331425e2668e40fc2cca40626d6f36520c3e599b5f6fe4a8758315cd2cdf15b"></a>

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

<a id="canonical-5711dec67236a7cfdb80b63ec82c848c0e2ce82472db978bd47cea7f582f854d"></a>

## Direct properties — service.advertise_options.advertise_custom.advertise_where.site.site / 3a9443faf44f / 3

<a id="canonical-88485ca8e18fd9570a91d7d4135dd90fedc2bd5853eecaf339ba3f1d86614c97"></a>

<a id="canonical-050109e23228e0364e20ea779508241e78c9250172e3f55df71f4ebe104fa9d9"></a>

## name property — service.advertise_options.advertise_custom.advertise_where.site.site / 3a9443faf44f / 4

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

<a id="canonical-1d3a8291bf5fe713c68ea2092687e62f1393f36f4b5b1ead236cccceff807857"></a>

<a id="canonical-14bfc8365fc1ca0c5dec8fb24129dc7288fdafc4a13fddccc83336eaeee5c8df"></a>

## namespace property — service.advertise_options.advertise_custom.advertise_where.site.site / 3a9443faf44f / 5

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

<a id="canonical-44f56052f92c92d73bae1e57ec38b136e75003c05e0fab7c952fa83ddb7cd7d2"></a>

<a id="canonical-58c863ab0a4dad58db0718465aba87bfa6c7b3ef748cfd4dce6dc3eae0ef0f3b"></a>

## tenant property — service.advertise_options.advertise_custom.advertise_where.site.site / 3a9443faf44f / 6

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

<a id="canonical-1935203994a46485d96cf2dee391d35fde34ab7beac4715a0e8741c7358743ea"></a>

## Next pages — service.advertise_options.advertise_custom.advertise_where.site.site / 3a9443faf44f / 7

- [service.advertise_options.advertise_custom.advertise_where.site](data-sources--workload--reference--group-005.md#canonical-241cb3a4738269e617f04ac2d9e3ed069f941593b2ed290de51857611664eace)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-966d0540fd2f7f482faff0251e2f54efa1f1e1d466df34e64f1e06c0cf596885"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e17cc0a20495a4350d65018c400a61d71c6240ccc015543a7e0c813722fda8b3"></a>

## service.advertise_options.advertise_custom.advertise_where.virtual_site — service.advertise_options.advertise_custom.advertise_where.virtual_site / c5cd761bf857 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-005.md#canonical-c828ef9397b96106571b5a7e48d0b7869905751a1f38a48eaf6662fdd150fd18)
- service.advertise_options.advertise_custom.advertise_where.virtual_site

<a id="canonical-6eb523488dc40097d93f55e01b9da2ea12c7653f79fc60c69666ee381fb5dd03"></a>

Type: `"single"`. Computed.

Defines a reference to a customer site virtual site along with network type where a load balancer
could be advertised.

Upstream description:

This defines a reference to a customer site virtual site along with network type where a load
balancer could be advertised.

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

<a id="canonical-3404f8e79e33e2b6b9f2eac64c09c941a62f447fe7556bfce8f58612545d2f17"></a>

## Direct properties — service.advertise_options.advertise_custom.advertise_where.virtual_site / c5cd761bf857 / 3

<a id="canonical-57468e4ac7dfd5d66a95d9ec1e93fc53400e73a41110f339552ce26744c9d031"></a>

<a id="canonical-38b31868cb41c47d66a6a6cc37466cd4c8b8ca44ecb7ca136e91e170b6cf8790"></a>

## network property — service.advertise_options.advertise_custom.advertise_where.virtual_site / c5cd761bf857 / 4

Type: `"string"`. Computed.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Upstream description:

This defines network types to be used on site

All inside and outside networks. All inside and outside networks with internet VIP support. All
inside networks. All outside networks. All outside networks with internet VIP support. VK8s service
network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](data-sources--workload--reference--group-005.md#canonical-50361f235059672c0c3f3284fbc9de3db78ce0c3183e7354b06f397dfe486167): complete subsection reference.

<a id="canonical-c9810270349f6c341d7a1028cf7ed03053896639ce7fcc42407a0a4434861079"></a>

## Next pages — service.advertise_options.advertise_custom.advertise_where.virtual_site / c5cd761bf857 / 5

- [service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site](data-sources--workload--reference--group-005.md#canonical-50361f235059672c0c3f3284fbc9de3db78ce0c3183e7354b06f397dfe486167)
- [service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-005.md#canonical-c828ef9397b96106571b5a7e48d0b7869905751a1f38a48eaf6662fdd150fd18)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-50361f235059672c0c3f3284fbc9de3db78ce0c3183e7354b06f397dfe486167"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d89bf3033f448083961f128c29689bc1788a018a7fab30be8725e0c49ba5633"></a>

## service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site — service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_ / 446a852ee047 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-005.md#canonical-c828ef9397b96106571b5a7e48d0b7869905751a1f38a48eaf6662fdd150fd18)
- [service.advertise_options.advertise_custom.advertise_where.virtual_site](data-sources--workload--reference--group-005.md#canonical-966d0540fd2f7f482faff0251e2f54efa1f1e1d466df34e64f1e06c0cf596885)
- service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site

<a id="canonical-2da4c9e21225c4d21963caf6d3b50adc6a974a4399478d4b0e4302ad772e0880"></a>

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

<a id="canonical-b15b04a0e379764f40bc93621924950f5e233ee29171231b68cd392943a65be7"></a>

## Direct properties — service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_ / 446a852ee047 / 3

<a id="canonical-cb42dfeee1e96d854f6243cf25a070f565a3877a5408b422b3e89bcce543fd6a"></a>

<a id="canonical-7c4c78ab0cefcc1008a3f5ac39b0e9eb8d29940e0697b5cbe2fa448d2dbef70b"></a>

## name property — service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_ / 446a852ee047 / 4

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

<a id="canonical-c246d164271d08889b6fd179a9bcdc957b2275effa2271c99c95521e694332c9"></a>

<a id="canonical-4d59480fbb4f7e4735f4d263f5a2ab8e52e09ec9832c42d007f50217a437a17e"></a>

## namespace property — service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_ / 446a852ee047 / 5

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

<a id="canonical-4777aa79a6f1d3df03865477def8ae3a6171ee21eac16d69bb6e7de8f9e79ad0"></a>

<a id="canonical-597b47824135b51bbf1f55e3e4431bfdd28fc2e2d6f40501add610f75f63cc8a"></a>

## tenant property — service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_ / 446a852ee047 / 6

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

<a id="canonical-2542a878485a2aeefd90d1f7dee0461419bd1817715fc5d8d90b3684c54653fa"></a>

## Next pages — service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_ / 446a852ee047 / 7

- [service.advertise_options.advertise_custom.advertise_where.virtual_site](data-sources--workload--reference--group-005.md#canonical-966d0540fd2f7f482faff0251e2f54efa1f1e1d466df34e64f1e06c0cf596885)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-ae3bbda4789096d2bb6def8889bf556d43025f52cee8343c45f3387a7374a178"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-83cea313870a8a2ce10fbba5a333637b12e1e3f09b7fb2d14255532f20289b6d"></a>

## service.advertise_options.advertise_custom.advertise_where.vk8s_service — service.advertise_options.advertise_custom.advertise_where.vk8s_service / 0fa6115dc85a / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-005.md#canonical-c828ef9397b96106571b5a7e48d0b7869905751a1f38a48eaf6662fdd150fd18)
- service.advertise_options.advertise_custom.advertise_where.vk8s_service

<a id="canonical-a8e519b3d924b0cc0776da360e0c2f7d6c731fcd0ccb6cf1aa414fbf564018c2"></a>

Type: `"single"`. Computed.

Defines a reference to a RE site or virtual site where a load balancer could be advertised in the
vK8s service network.

Upstream description:

This defines a reference to a RE site or virtual site where a load balancer could be advertised in
the vK8s service network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

<a id="canonical-965024b4e0ef9c5c5d52c6eccdc9ac4b9a71ab7d3badf58f6d6ce653c97dea55"></a>

## Direct properties — service.advertise_options.advertise_custom.advertise_where.vk8s_service / 0fa6115dc85a / 3

- [site](data-sources--workload--reference--group-005.md#canonical-3bd0d261c014977a335f2c8ce7515d18508ba56bafcc9317c5836a4f73d37abd): complete subsection reference.

- [virtual_site](data-sources--workload--reference--group-005.md#canonical-7e775b9fa5d101efe00b2fc7f08394366636eb5b18aa2a2089fb3f9e483c4edc): complete subsection reference.

<a id="canonical-22dc8fc790aebbe0314dd05bbf9f3f0464bcdf7e83999d7d79f3544f871a1ad2"></a>

## Next pages — service.advertise_options.advertise_custom.advertise_where.vk8s_service / 0fa6115dc85a / 4

- [service.advertise_options.advertise_custom.advertise_where.vk8s_service.site](data-sources--workload--reference--group-005.md#canonical-3bd0d261c014977a335f2c8ce7515d18508ba56bafcc9317c5836a4f73d37abd)
- [service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site](data-sources--workload--reference--group-005.md#canonical-7e775b9fa5d101efe00b2fc7f08394366636eb5b18aa2a2089fb3f9e483c4edc)
- [service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-005.md#canonical-c828ef9397b96106571b5a7e48d0b7869905751a1f38a48eaf6662fdd150fd18)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-3bd0d261c014977a335f2c8ce7515d18508ba56bafcc9317c5836a4f73d37abd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a98fcc368bb397d850350a3c0a94f2f9eb57816a74999e190bad613b4f4803f6"></a>

## service.advertise_options.advertise_custom.advertise_where.vk8s_service.site — service.advertise_options.advertise_custom.advertise_where.vk8s_service.site / 137e638fc33c / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-005.md#canonical-c828ef9397b96106571b5a7e48d0b7869905751a1f38a48eaf6662fdd150fd18)
- [service.advertise_options.advertise_custom.advertise_where.vk8s_service](data-sources--workload--reference--group-005.md#canonical-ae3bbda4789096d2bb6def8889bf556d43025f52cee8343c45f3387a7374a178)
- service.advertise_options.advertise_custom.advertise_where.vk8s_service.site

<a id="canonical-802568fd0c8533b898d3596f3a17997059b7f8a6dc68a173822b6a526d54fb24"></a>

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

<a id="canonical-255aca1dc3284c7ea089ba148ccb4174baa832f3ffdd262e1e5b975214f81faf"></a>

## Direct properties — service.advertise_options.advertise_custom.advertise_where.vk8s_service.site / 137e638fc33c / 3

<a id="canonical-ae4ffea4045915b55c5dd5296811225c1b40bec62f123225efacf8d79ac04fea"></a>

<a id="canonical-21b718c07ccb02268bab4193f3dcb61ad830fe9e260571680f31231c09c3cd0f"></a>

## name property — service.advertise_options.advertise_custom.advertise_where.vk8s_service.site / 137e638fc33c / 4

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

<a id="canonical-aa9631ccefe7bb4e0e3716c4864d0dac7bd0ca1f5e477d3a684842d5118821fd"></a>

<a id="canonical-7589ada4817a173261ecdb428544eb318c54bb6685a38d9d3af8b5aa7478f7ba"></a>

## namespace property — service.advertise_options.advertise_custom.advertise_where.vk8s_service.site / 137e638fc33c / 5

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

<a id="canonical-c09cf48e8677cbb910e968a2210921ee06c3bc66475c6c94ab12a0748ff582db"></a>

<a id="canonical-d32918618e46e7a6f1046991822318c44fcb4e5466126eb34b63a859de861c16"></a>

## tenant property — service.advertise_options.advertise_custom.advertise_where.vk8s_service.site / 137e638fc33c / 6

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

<a id="canonical-26d6153e41010e88e16fb6de19eb3762fe98ad62df304ea7ebf3a1d6e5dd9885"></a>

## Next pages — service.advertise_options.advertise_custom.advertise_where.vk8s_service.site / 137e638fc33c / 7

- [service.advertise_options.advertise_custom.advertise_where.vk8s_service](data-sources--workload--reference--group-005.md#canonical-ae3bbda4789096d2bb6def8889bf556d43025f52cee8343c45f3387a7374a178)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-7e775b9fa5d101efe00b2fc7f08394366636eb5b18aa2a2089fb3f9e483c4edc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e5f8d12c21cc823f668b99a5e4cd8f79c01520f48458884e7b952d6176093962"></a>

## service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site — service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_ / ae512e548982 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-005.md#canonical-c828ef9397b96106571b5a7e48d0b7869905751a1f38a48eaf6662fdd150fd18)
- [service.advertise_options.advertise_custom.advertise_where.vk8s_service](data-sources--workload--reference--group-005.md#canonical-ae3bbda4789096d2bb6def8889bf556d43025f52cee8343c45f3387a7374a178)
- service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="canonical-28830d8cb023deeb3ba1d8f9b1b925664cacefdc14ed433e71e1d4663075b5a3"></a>

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

<a id="canonical-1a1f9a4b60402bf3488a0c2380e35f5093ebdd70d0f9e8fb9fa9aa9803c6d49e"></a>

## Direct properties — service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_ / ae512e548982 / 3

<a id="canonical-dbbe916a8788f6b7d663b7514ef7aee773abfef02bc0e7b81201a6a9ed2872cf"></a>

<a id="canonical-8e09ae38bbd8a4d1cedbcc980ec481f3a30b362efd7dedc320ee8a7cfc0894e3"></a>

## name property — service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_ / ae512e548982 / 4

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

<a id="canonical-a410e6204583f45985faa0a7af8760cc811d4817eb9c4959e05d17bffd4df7a9"></a>

<a id="canonical-7247290dde3e6600f79e3304a531965e8b3faebbbc34d759af39b8bd1d6d79e7"></a>

## namespace property — service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_ / ae512e548982 / 5

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

<a id="canonical-4ea58fd3908d4dc865edc7460a71e64e197ba83a58e336ed3e97f7a8556e8f2d"></a>

<a id="canonical-5fa609af802e36c15799e636a7eb5573b6cbb0424272851a5489adc87663493b"></a>

## tenant property — service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_ / ae512e548982 / 6

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

<a id="canonical-528843cbf6ddee27b1e813bf5bca07e0263ba7ed8c87cad6b9ee1a2468c1486d"></a>

## Next pages — service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_ / ae512e548982 / 7

- [service.advertise_options.advertise_custom.advertise_where.vk8s_service](data-sources--workload--reference--group-005.md#canonical-ae3bbda4789096d2bb6def8889bf556d43025f52cee8343c45f3387a7374a178)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f0b9f844b01e126e9349e3c9041cb47b9a272f6b9feb084bf78257c128395720"></a>

## service.advertise_options.advertise_custom.ports — service.advertise_options.advertise_custom.ports / c763f6269d82 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- service.advertise_options.advertise_custom.ports

<a id="canonical-265d34f959effd3ecf5673c5c7743984cc4ae7965f75834487293ae0659176cb"></a>

Type: `"list"`. Computed.

Ports. Ports to advertise.

Upstream description:

Ports to advertise.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-253298ebbb49df575939d8b56ed4c1285299107c91dd40f65c470e4cf682ef16"></a>

## Direct properties — service.advertise_options.advertise_custom.ports / c763f6269d82 / 3

- [http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c): complete subsection reference.

- [port](data-sources--workload--reference--group-008.md#canonical-7c613d4efb08afb76c10e5423143ef3447c3a5c0c851a6306c101db9750216aa): complete subsection reference.

- [tcp_loadbalancer](data-sources--workload--reference--group-008.md#canonical-671edfd9f0410a2c5d55e0e8ca6a82232079d713e8ffed70a89baf527d62d65d): complete subsection reference.

<a id="canonical-e2f6cebccd1d082f2fa401138a2ee429ff4ce0a3c1b6c2f5fd50a16308235723"></a>

## Next pages — service.advertise_options.advertise_custom.ports / c763f6269d82 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.port](data-sources--workload--reference--group-008.md#canonical-7c613d4efb08afb76c10e5423143ef3447c3a5c0c851a6306c101db9750216aa)
- [service.advertise_options.advertise_custom.ports.tcp_loadbalancer](data-sources--workload--reference--group-008.md#canonical-671edfd9f0410a2c5d55e0e8ca6a82232079d713e8ffed70a89baf527d62d65d)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d097e11632c6e1db44da9b57dc1bd29a526933cd68e183a9daa583b2b86769d4"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer — service.advertise_options.advertise_custom.ports.http_loadbalancer / c7f9fa983438 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- service.advertise_options.advertise_custom.ports.http_loadbalancer

<a id="canonical-8623cf716aee2869d18aa9752a0d23e9b52b62c2eb79325886be15d4a3d7a0eb"></a>

Type: `"single"`. Computed.

Configuration parameter for http loadbalancer.

Upstream description:

HTTP/HTTPS Load balancer.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-loadbalancer_type": "[\"http\",\"https\",\"https_auto_cert\"]",
  "x-ves-oneof-field-route_choice": "[\"default_route\",\"specific_routes\"]"
}
```

<a id="canonical-17afa3bb40aa2fb6e7c0c849306de1317e2439a310f89160ea5bc5471fbde04a"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer / c7f9fa983438 / 3

- [default_route](data-sources--workload--reference--group-005.md#canonical-50d38510b2656a8a5020f53866f535989d6be78dfa73be0c3eb5a13385071a26): complete subsection reference.

<a id="canonical-3cb23cd884024ed2b9e858fd873938fde373a5c1e52fee9cb22778137c210564"></a>

<a id="canonical-ad997d7ff104ee9d5cd6518f7a8af6ab75fa0baca310ee4511ed4a9084cdf523"></a>

## domains property — service.advertise_options.advertise_custom.ports.http_loadbalancer / c7f9fa983438 / 4

Type: `["list", "string"]`. Computed.

List of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form Domain search order: 1. Exact domain names: \`\` is invalid
Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the..

Upstream description:

A list of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form

Domain search order: &#8203;1. Exact domain names: \`\`www&#46;example.com\`\`. &#8203;2. Prefix
domain wildcards: \`\`\*.example.com\`\` or \`\`\*.bar.example.com\`\`. &#8203;3. Special wildcard
\`\`\*\`\` matching any domain.

Wildcard will not match empty string. E.g. \`\`\*.example.com\`\` will match \`\`bar.example.com\`\`
and \`\`baz-bar.example.com\`\` but not \`\`.example.com\`\`. The longest wildcards match first.
Wildcards must match a whole DNS label. E.g. \`\`\*.example.com\`\` and \*.bar.example.com are
valid, however \`\`\*bar.example.com\`\` or \`\`\*-bar.example.com\`\` is invalid

Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the
list of names for which DNS resolution will be done by VER.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [http](data-sources--workload--reference--group-005.md#canonical-9d3963bd66017de21a370a0459b5316841c4505d304daa4952e0efc947c06786): complete subsection reference.

- [https](data-sources--workload--reference--group-005.md#canonical-d1b08993403cab5df5c13d77dd4886b4bc149ebb413780deb8732e32538d3f11): complete subsection reference.

- [https_auto_cert](data-sources--workload--reference--group-006.md#canonical-2635aea08ab6ac057fe2eb8518ae90c68497a79c42f9e3648d113dc8f81ce7fc): complete subsection reference.

- [specific_routes](data-sources--workload--reference--group-007.md#canonical-834006c937eaad6c8f1af05c2ebf1bd39aa6ba171e7d68395eaafc99f2b23801): complete subsection reference.

<a id="canonical-4e849f3873192257cea4e3a1d5df0875b7ee09a169d82bf57170f56875fef515"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer / c7f9fa983438 / 5

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-005.md#canonical-50d38510b2656a8a5020f53866f535989d6be78dfa73be0c3eb5a13385071a26)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.http](data-sources--workload--reference--group-005.md#canonical-9d3963bd66017de21a370a0459b5316841c4505d304daa4952e0efc947c06786)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-d1b08993403cab5df5c13d77dd4886b4bc149ebb413780deb8732e32538d3f11)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-2635aea08ab6ac057fe2eb8518ae90c68497a79c42f9e3648d113dc8f81ce7fc)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-834006c937eaad6c8f1af05c2ebf1bd39aa6ba171e7d68395eaafc99f2b23801)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-50d38510b2656a8a5020f53866f535989d6be78dfa73be0c3eb5a13385071a26"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9128e9b7983e10270f97df4cdd6252de68448edd8c0ff7950a4f12c6aa538808"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route — service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route / 13b59c72965e / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route

<a id="canonical-dbda06e677553c4145dd658dc8e25c3b217c38623625ea20c0a29ea127a98ccb"></a>

Type: `"single"`. Computed.

Configuration parameter for default route.

Upstream description:

Default route matching all APIs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-host_rewrite_params": "[\"auto_host_rewrite\",\"disable_host_rewrite\",\"host_rewrite\"]"
}
```

<a id="canonical-622f4b9f9aa267ba2b52eb9c8b906207b5eb1a5ca1c471470589d282ea1a5a75"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route / 13b59c72965e / 3

- [auto_host_rewrite](data-sources--workload--reference--group-005.md#canonical-e914865654f44764f8639799b8e4e0ff5a489a0c840c403ea6357166e2e5c36e): complete subsection reference.

- [disable_host_rewrite](data-sources--workload--reference--group-005.md#canonical-ceaea744209cb0b633609536e5964326660d491b91dd90cf68a714bb91362a75): complete subsection reference.

<a id="canonical-fb771fdaf20f663e96235b76356feff9fddafe264ff9b4b59ee8f9227a472360"></a>

<a id="canonical-94a253e9f128564eef9bc8ee0e4137f861203650b76e7af710c1087cb157ac47"></a>

## host_rewrite property — service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route / 13b59c72965e / 4

Type: `"string"`. Computed.

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Upstream description:

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

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

<a id="canonical-948ad2a1a7dd0fc3134fb1831876c786724c7cb405c521490d43abe687c4cf09"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route / 13b59c72965e / 5

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.auto_host_rewrite](data-sources--workload--reference--group-005.md#canonical-e914865654f44764f8639799b8e4e0ff5a489a0c840c403ea6357166e2e5c36e)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.disable_host_rewrite](data-sources--workload--reference--group-005.md#canonical-ceaea744209cb0b633609536e5964326660d491b91dd90cf68a714bb91362a75)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-e914865654f44764f8639799b8e4e0ff5a489a0c840c403ea6357166e2e5c36e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-90c8739515d14b91ba2d5b4a5f976e09587947cd32b03f033fb1287006e5fcb1"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.auto_host_rewrite — service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route / f32f23e4adb7 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-005.md#canonical-50d38510b2656a8a5020f53866f535989d6be78dfa73be0c3eb5a13385071a26)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.auto_host_rewrite

<a id="canonical-0a2a47d0df13fa6e855db348e7c24801fc197b5e5c293ce0029d0435cbaf91d7"></a>

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

<a id="canonical-78c82bec090d8efa60ebb733c5f608dcfb92e41d33bbd6bdd4a888e41dcb4e86"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route / f32f23e4adb7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a39bdba5f0863022f62ad9aa67bada333062b0179a84ff49533fd378b92a4efd"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route / f32f23e4adb7 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-005.md#canonical-50d38510b2656a8a5020f53866f535989d6be78dfa73be0c3eb5a13385071a26)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-ceaea744209cb0b633609536e5964326660d491b91dd90cf68a714bb91362a75"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4aff79f110563e3083e7639459f4e2c0a8f167835b696dbb833f8cee886da1dc"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.disable_host_rewrite — service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route / adea06397319 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-005.md#canonical-50d38510b2656a8a5020f53866f535989d6be78dfa73be0c3eb5a13385071a26)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.disable_host_rewrite

<a id="canonical-414be6aa57c74a65f0f9fc921c36d3364745d8a04df4f05ce4116ccc97d19a38"></a>

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

<a id="canonical-aa095319e140173f4bb4f72793691cb3a045a02dcdba57881c4d6a64bd15d5b8"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route / adea06397319 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f6bffe400626ffc731668bdf1bad34cb84df7299ed8dbe5a03fc979ad4650fef"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route / adea06397319 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-005.md#canonical-50d38510b2656a8a5020f53866f535989d6be78dfa73be0c3eb5a13385071a26)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-9d3963bd66017de21a370a0459b5316841c4505d304daa4952e0efc947c06786"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2732fc3bdab5407acbe3d7977e438ed26f57a82092db545d31288f77056a67d2"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.http — service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 07ad5d33446d / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.http

<a id="canonical-378881f176a36337a9aa23bf6ea8e980598a0c0df3e012981b88a17bab93ffcc"></a>

Type: `"single"`. Computed.

HTTP Choice. Choice for selecting HTTP proxy.

Upstream description:

Choice for selecting HTTP proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]"
}
```

<a id="canonical-86f8d2a6905fc365bb56287852ecd77a29ba1969711c37280d04d1454a070a16"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 07ad5d33446d / 3

<a id="canonical-48cc17d39c091b2a6d4e8112c0d50ca4ef90f346397707def7c476f642a45ebf"></a>

<a id="canonical-d16fe9a95da173466df59e45c8eb6a008f0595618c1ca48067b378d79901680b"></a>

## dns_volterra_managed property — service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 07ad5d33446d / 4

Type: `"bool"`. Computed.

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

Upstream description:

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

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

<a id="canonical-ab9166696e7be184d4b129bd381abefcc349566f180fae6a94802cb0ccb5ac0a"></a>

<a id="canonical-db06e335ae868edbcf120a0fcc9970f3036a2b993fc68895d3f2cde12bbacb27"></a>

## port property — service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 07ad5d33446d / 5

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTP port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTP port to Listen.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-e03c72f0133e9ecfec0c5ced03509529c212ae2d45f2d18adb73bb43b62aa0d1"></a>

<a id="canonical-be987221dc033640f97ed95195bfdb27b95f3a23399cd6147afb6211653939fc"></a>

## port_ranges property — service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 07ad5d33446d / 6

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-88952b062949b96413ad02d284a44381f9cb56651fe4c9696516d5a20f271c05"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 07ad5d33446d / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-d1b08993403cab5df5c13d77dd4886b4bc149ebb413780deb8732e32538d3f11"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e1f3b09817a962137dfc66e6cd5ba5d6cde0392621ef87f76ba564f47556e18"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https — service.advertise_options.advertise_custom.ports.http_loadbalancer.https / cdc00ed01bc0 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https

<a id="canonical-9c3c35ef5b46ea01eafb915a0f508f9bc0648bcb6667aa41dcd8c96521bac7ec"></a>

Type: `"single"`. Computed.

Choice for selecting HTTP proxy with bring your own certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-default_lb_choice": "[\"default_loadbalancer\",\"non_default_loadbalancer\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]",
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_parameters\"]"
}
```

<a id="canonical-20a48794af282f4d6cd958c5d3e26f31344257b681fd927166c592166f089c91"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https / cdc00ed01bc0 / 3

<a id="canonical-00df11d8b25eabc971c86188218a3af4ca1e61a1ca6cd6dfeeb7aac91e3bffe4"></a>

<a id="canonical-abcbbfc85f216ebcb5c5938a29e1ecb4b4db13ac0b1e9a43a6f36d1854d2cdec"></a>

## add_hsts property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https / cdc00ed01bc0 / 4

Type: `"bool"`. Computed.

Add HTTP Strict-Transport-Security response header.

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

<a id="canonical-32e9b5fe6f161f560f20ce867134175e971c37541543d4e594374fe66ea7b0f6"></a>

<a id="canonical-749c54d13c13b45f654507ddab3145a34246cdfe57b946748fd0578908f504c9"></a>

## append_server_name property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https / cdc00ed01bc0 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Upstream description:

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [coalescing_options](data-sources--workload--reference--group-005.md#canonical-3aad321140474664f38f14aa4e0b2eed31dec599f89938a6379e5b9e2665e89d): complete subsection reference.

<a id="canonical-4b3915816249e0fe917a3db6e53a7762a8da0b0b1e3ee3e8497e013c9521f95e"></a>

<a id="canonical-60ae8d33e16a7a5840decd3a4fc7ba1ec7dc046ccfc5262e34c23b6538779f0a"></a>

## connection_idle_timeout property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https / cdc00ed01bc0 / 6

Type: `"number"`. Computed.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed.

Upstream description:

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [default_header](data-sources--workload--reference--group-005.md#canonical-3311714c5045eeead4f28b355b5cae554d2d3b76ed1f5dcf4fdb570b82101f05): complete subsection reference.

- [default_loadbalancer](data-sources--workload--reference--group-005.md#canonical-b3d1c24ae62085f6a922a8a632ea0bc3c59d94bf77ae4b339af36915746dab8b): complete subsection reference.

- [disable_path_normalize](data-sources--workload--reference--group-005.md#canonical-7b959c143396089fe957f8bbfcc5be7732031e12c88365798f7b9c73ad9ebd72): complete subsection reference.

- [enable_path_normalize](data-sources--workload--reference--group-005.md#canonical-66ffa44ec3955e75b0457621035c292e4ad5eef60024a182a4067e7101896286): complete subsection reference.

- [http_protocol_options](data-sources--workload--reference--group-005.md#canonical-e93a4484ae0ff3a088ee961ce53b6e26b6fe43173eb5e3b76d003330eee577b2): complete subsection reference.

<a id="canonical-5c2f8345cfdefe7f1fa4b30d3935e4b3fdfb4e2bf4e2979b000b8b635ad239b0"></a>

<a id="canonical-e9e4d31ed4be70fe5f2d4941fae9db26ac9a564937ab3f1d3a363f5fb3334416"></a>

## http_redirect property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https / cdc00ed01bc0 / 7

Type: `"bool"`. Computed.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

Upstream description:

Redirect HTTP traffic to HTTPS.

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

- [non_default_loadbalancer](data-sources--workload--reference--group-005.md#canonical-f57f61c397bb7032e3acb8d3efc7093f3cf3e9976abba49aec384be7e244760a): complete subsection reference.

- [pass_through](data-sources--workload--reference--group-005.md#canonical-e06926e1633305c31d13ace68ff2e8c8cebd468c6d4ff002248512eb512e33cf): complete subsection reference.

<a id="canonical-9ebd974bbe1ee79fcca42a04cd308ee044f0870130f4cd2325eb791cbf506d3f"></a>

<a id="canonical-e471e9c006f5a50b62ab774a1d55240434dbedd69ba8c7ee93a36bb11ecca416"></a>

## port property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https / cdc00ed01bc0 / 8

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTPS port to Listen.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-89ae1ffa967b9b1abff70bf2fb8aa85983b4c42ee76a86090bd8309924345e7d"></a>

<a id="canonical-4302d7a9a4bfef66d7547974f2a69323acb84cc563880c47c37e610457ded929"></a>

## port_ranges property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https / cdc00ed01bc0 / 9

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-d742ffa2b79fbe9cb32c210b982e78a9c5cd4466e255066db907cbb36a6d2c21"></a>

<a id="canonical-b5d33c4afc44a2bfe2c31d719edce643b8d051cb1b1635c5414383312495826b"></a>

## server_name property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https / cdc00ed01bc0 / 10

Type: `"string"`. Computed.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [tls_cert_params](data-sources--workload--reference--group-005.md#canonical-12b8761594eb6c09dcd0306020eb8588653df046eb14cf94f3fa3bdf773bdb0f): complete subsection reference.

- [tls_parameters](data-sources--workload--reference--group-006.md#canonical-ecd940b92d0b2a95772cbaa3c856c213e60b1cdd47c7e079af23d6c112f924ae): complete subsection reference.

<a id="canonical-50ffb99196515b4b128a68b3ac2c8b77a27343218ac1078ede42011fafc07bb9"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https / cdc00ed01bc0 / 11

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-005.md#canonical-3aad321140474664f38f14aa4e0b2eed31dec599f89938a6379e5b9e2665e89d)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_header](data-sources--workload--reference--group-005.md#canonical-3311714c5045eeead4f28b355b5cae554d2d3b76ed1f5dcf4fdb570b82101f05)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_loadbalancer](data-sources--workload--reference--group-005.md#canonical-b3d1c24ae62085f6a922a8a632ea0bc3c59d94bf77ae4b339af36915746dab8b)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.disable_path_normalize](data-sources--workload--reference--group-005.md#canonical-7b959c143396089fe957f8bbfcc5be7732031e12c88365798f7b9c73ad9ebd72)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.enable_path_normalize](data-sources--workload--reference--group-005.md#canonical-66ffa44ec3955e75b0457621035c292e4ad5eef60024a182a4067e7101896286)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-005.md#canonical-e93a4484ae0ff3a088ee961ce53b6e26b6fe43173eb5e3b76d003330eee577b2)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.non_default_loadbalancer](data-sources--workload--reference--group-005.md#canonical-f57f61c397bb7032e3acb8d3efc7093f3cf3e9976abba49aec384be7e244760a)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.pass_through](data-sources--workload--reference--group-005.md#canonical-e06926e1633305c31d13ace68ff2e8c8cebd468c6d4ff002248512eb512e33cf)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-005.md#canonical-12b8761594eb6c09dcd0306020eb8588653df046eb14cf94f3fa3bdf773bdb0f)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-ecd940b92d0b2a95772cbaa3c856c213e60b1cdd47c7e079af23d6c112f924ae)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-3aad321140474664f38f14aa4e0b2eed31dec599f89938a6379e5b9e2665e89d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f54bdc0e317751530423f524d7b80bf6b6b852b98440d645adc27ff580f2e835"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalesc / 8551fcfd4ebf / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-d1b08993403cab5df5c13d77dd4886b4bc149ebb413780deb8732e32538d3f11)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options

<a id="canonical-895d4d1689007b6accba76f9b18bd972c5b59c11271dc45b0ce3601b4463f045"></a>

Type: `"single"`. Computed.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

<a id="canonical-0d8261010daf330e331446570dfc26e66be590f91f0a08cce2c4be3bdb082b89"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalesc / 8551fcfd4ebf / 3

- [default_coalescing](data-sources--workload--reference--group-005.md#canonical-4c5967bc0a35f8a13d1ef44453687cbe274ffab90b910cc05bc90cb91af239a8): complete subsection reference.

- [strict_coalescing](data-sources--workload--reference--group-005.md#canonical-537fcae19b453cf7369c01a3bccbd05b4eb51869fa1403215c6037f67e6040b5): complete subsection reference.

<a id="canonical-586df37d5f1dd3bfcc93b501845cb04150aaa9f31737afa33797b3e90448cc0d"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalesc / 8551fcfd4ebf / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.default_coalescing](data-sources--workload--reference--group-005.md#canonical-4c5967bc0a35f8a13d1ef44453687cbe274ffab90b910cc05bc90cb91af239a8)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.strict_coalescing](data-sources--workload--reference--group-005.md#canonical-537fcae19b453cf7369c01a3bccbd05b4eb51869fa1403215c6037f67e6040b5)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-d1b08993403cab5df5c13d77dd4886b4bc149ebb413780deb8732e32538d3f11)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-4c5967bc0a35f8a13d1ef44453687cbe274ffab90b910cc05bc90cb91af239a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-550c2d5118ed3cc77aa15fd2cd16f99100ffa89054a67e6e96624c55221b6460"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.default_coalescing — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalesc / eab7a0f3d816 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-d1b08993403cab5df5c13d77dd4886b4bc149ebb413780deb8732e32538d3f11)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-005.md#canonical-3aad321140474664f38f14aa4e0b2eed31dec599f89938a6379e5b9e2665e89d)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.default_coalescing

<a id="canonical-f339a3072a6b7514b4f06a89fa842e77e5d283d9ef3f418222fffa89151f210d"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default coalescing.

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

<a id="canonical-f16a72bb441ee271ec0053608765dd389e0e354bce3de321d78922141304f439"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalesc / eab7a0f3d816 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8b3fca4c5f97e3b61981c2b03145c9982c88ae3c3e2ff11fe1835b90f30db017"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalesc / eab7a0f3d816 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-005.md#canonical-3aad321140474664f38f14aa4e0b2eed31dec599f89938a6379e5b9e2665e89d)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-537fcae19b453cf7369c01a3bccbd05b4eb51869fa1403215c6037f67e6040b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f93d21c1e9ee5f9ab1bbf396319c100def00d80778fbfec9fac55db55ac72bf"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.strict_coalescing — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalesc / ab0f9bf8ee95 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-d1b08993403cab5df5c13d77dd4886b4bc149ebb413780deb8732e32538d3f11)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-005.md#canonical-3aad321140474664f38f14aa4e0b2eed31dec599f89938a6379e5b9e2665e89d)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.strict_coalescing

<a id="canonical-5768d9bd92230e0b1a38072af49602a015effb859b832793b2754b2392f88568"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for strict coalescing.

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

<a id="canonical-ab6af2bd6aa05b5cf06537b4142f0a57d224d8200cf04ffc6d8f7e6080b8c26d"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalesc / ab0f9bf8ee95 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cdf2c87c8344eb6dc53f196cd8c38dbf09eba7f2f66e6f2aa9ee060fec39e357"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalesc / ab0f9bf8ee95 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-005.md#canonical-3aad321140474664f38f14aa4e0b2eed31dec599f89938a6379e5b9e2665e89d)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-3311714c5045eeead4f28b355b5cae554d2d3b76ed1f5dcf4fdb570b82101f05"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8945e70b40a699dd3a363b7f9a09625e443b50118f622e3e9572240af0d2c870"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_header — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default / f0499a206ec7 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-d1b08993403cab5df5c13d77dd4886b4bc149ebb413780deb8732e32538d3f11)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_header

<a id="canonical-ee9ad7ad3ddb632af0e0d765319e02545893587cff50227642adde91e83467e7"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default header.

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

<a id="canonical-92e7d0e1a1a4ea1ac4da1d5f8c1486a6242e0e5bfa4b819734371442289ba5c8"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default / f0499a206ec7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-00711a122f9e5fddd6f6ac2b0745a073ac783e48c6d1d52999705381b72e6ef9"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default / f0499a206ec7 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-d1b08993403cab5df5c13d77dd4886b4bc149ebb413780deb8732e32538d3f11)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-b3d1c24ae62085f6a922a8a632ea0bc3c59d94bf77ae4b339af36915746dab8b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9fdcbc335a46464f330ec9c829684f50f99498f2961a8c00b2782d70ed233f10"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_loadbalancer — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default / 1e44eae3cc2d / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-d1b08993403cab5df5c13d77dd4886b4bc149ebb413780deb8732e32538d3f11)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_loadbalancer

<a id="canonical-209f2df8bf61cd52750b75d9023ae664ab9bf62f2c139ba7df6ccd3ce2f5cbd8"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default loadbalancer.

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

<a id="canonical-f146ddcc8905387b288fa553a1575e5955a0450980d10d7eb84feba3337fd5fe"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default / 1e44eae3cc2d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-82eded320f67d7c2bf76162813ab2b36534aea8a8200f12490df65281378d1ec"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default / 1e44eae3cc2d / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-d1b08993403cab5df5c13d77dd4886b4bc149ebb413780deb8732e32538d3f11)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-7b959c143396089fe957f8bbfcc5be7732031e12c88365798f7b9c73ad9ebd72"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dcc1f89964f1749fd23dc2f71ba4f26be90c83bafafd996542bb5ec8620a3860"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.disable_path_normalize — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.disable / 3bfcc4b15a55 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-d1b08993403cab5df5c13d77dd4886b4bc149ebb413780deb8732e32538d3f11)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.disable_path_normalize

<a id="canonical-86dcb79fea17e8f0f493511397c5bf3100e1584dcd3041330b92d30755af4a55"></a>

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

<a id="canonical-d36434cde79f893fc06bad2909f655a42815b1c911bacb63b898ccfb0e9da3dd"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.disable / 3bfcc4b15a55 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-12c7578df6083d49157b4059338300eca0e7c68a8b66b1d46245693d047c07b1"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.disable / 3bfcc4b15a55 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-d1b08993403cab5df5c13d77dd4886b4bc149ebb413780deb8732e32538d3f11)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-66ffa44ec3955e75b0457621035c292e4ad5eef60024a182a4067e7101896286"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df5537ad660a15f436f1713c11e839edae984c15f856f0be7f6dd585a85734ed"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.enable_path_normalize — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.enable_ / 8ef11fa682ce / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-d1b08993403cab5df5c13d77dd4886b4bc149ebb413780deb8732e32538d3f11)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.enable_path_normalize

<a id="canonical-0759be5991056351a1346f59275984094082bc25a5ce692c8df95bdb67b9d1a2"></a>

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

<a id="canonical-d6e42c6b1e23565d052ba687c8cc7463bbb1ed6308ee58fc485470caf4387a17"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.enable_ / 8ef11fa682ce / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0c59588a03b52827fe4fb077fd5f41f7e3900a25d838fa1e4b1f1ff6d0a09887"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.enable_ / 8ef11fa682ce / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-d1b08993403cab5df5c13d77dd4886b4bc149ebb413780deb8732e32538d3f11)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-e93a4484ae0ff3a088ee961ce53b6e26b6fe43173eb5e3b76d003330eee577b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df64d21c0d2b7214618fe942d2cd306287a4dd930b61c186d0b50b17eef0c69c"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 91321b07680b / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-d1b08993403cab5df5c13d77dd4886b4bc149ebb413780deb8732e32538d3f11)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options

<a id="canonical-acb7c1b9a001d0788d14f80654ecc216b2755580d374409ac90596c42d5952cf"></a>

Type: `"single"`. Computed.

HTTP protocol configuration OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

<a id="canonical-7bfab4409cae0d9ab66054d605097563c529922299644a2d7303523dda4be8ac"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 91321b07680b / 3

- [http_protocol_enable_v1_only](data-sources--workload--reference--group-005.md#canonical-c84f985aa9a22d077e735e9765ba2c27ca595c47bc7c85afcfc5b6ad2d143af8): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--workload--reference--group-005.md#canonical-4a38b6c802cd47806d9d63f5751ecb16d014b6ec7ae64ec70e7e8d315cca53ac): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--workload--reference--group-005.md#canonical-09efa3a9dccbe6b69648c54d95e300c31d985062672a7071dc3bdf61b8074856): complete subsection reference.

<a id="canonical-8e7cee93b7e467a5eea26a798e2619933f3105ec61c2cf2dea2c3dad6a2ac511"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 91321b07680b / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-005.md#canonical-c84f985aa9a22d077e735e9765ba2c27ca595c47bc7c85afcfc5b6ad2d143af8)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2](data-sources--workload--reference--group-005.md#canonical-4a38b6c802cd47806d9d63f5751ecb16d014b6ec7ae64ec70e7e8d315cca53ac)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only](data-sources--workload--reference--group-005.md#canonical-09efa3a9dccbe6b69648c54d95e300c31d985062672a7071dc3bdf61b8074856)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-d1b08993403cab5df5c13d77dd4886b4bc149ebb413780deb8732e32538d3f11)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-c84f985aa9a22d077e735e9765ba2c27ca595c47bc7c85afcfc5b6ad2d143af8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f183a53142f7a5f8bfafd496ca55730faa08c27c28d59bfa7a14098c1444de9"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 7757f8241122 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-d1b08993403cab5df5c13d77dd4886b4bc149ebb413780deb8732e32538d3f11)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-005.md#canonical-e93a4484ae0ff3a088ee961ce53b6e26b6fe43173eb5e3b76d003330eee577b2)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-f2c73e308b089aa93a31b85e37df423d98f2d31cbbff525c6a6eb9c71599ece6"></a>

Type: `"single"`. Computed.

HTTP/1.1 Protocol OPTIONS for downstream connections.

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

<a id="canonical-02ffdba79a7c614a206438aee6735d4345812a10166039760123521654adcb18"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 7757f8241122 / 3

- [header_transformation](data-sources--workload--reference--group-005.md#canonical-63ff932f43849ea58268252bacbdc7aec7c3cc2ba0bfd36728c39006978f8b1f): complete subsection reference.

<a id="canonical-0f2e29cbe02dd9581088edb74f8e40b84fc8463d66504848011e4885110c9ae1"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 7757f8241122 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-005.md#canonical-63ff932f43849ea58268252bacbdc7aec7c3cc2ba0bfd36728c39006978f8b1f)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-005.md#canonical-e93a4484ae0ff3a088ee961ce53b6e26b6fe43173eb5e3b76d003330eee577b2)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-63ff932f43849ea58268252bacbdc7aec7c3cc2ba0bfd36728c39006978f8b1f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-844f8c53c4726624aa8d9440cdc4f74a63577318dc096d0b70e3524cc756b0be"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / f2e8dc467de7 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-d1b08993403cab5df5c13d77dd4886b4bc149ebb413780deb8732e32538d3f11)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-005.md#canonical-e93a4484ae0ff3a088ee961ce53b6e26b6fe43173eb5e3b76d003330eee577b2)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-005.md#canonical-c84f985aa9a22d077e735e9765ba2c27ca595c47bc7c85afcfc5b6ad2d143af8)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-21eeb67f3212c17698725f2ce3d322bf78b5b6ef513ecf0c89c7f3ac5e5f7401"></a>

Type: `"single"`. Computed.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

<a id="canonical-3b76c869c4c797d7f7c067d33be1f1d6e31212685e9ee1a5ccef00ec511fd5d3"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / f2e8dc467de7 / 3

- [default_header_transformation](data-sources--workload--reference--group-005.md#canonical-72285ad952940abe4df34074485bba37ae3190a48418235fa9cfb90efe71518e): complete subsection reference.

- [preserve_case_header_transformation](data-sources--workload--reference--group-005.md#canonical-7e24f8a2fdd0e5aad7503285622661cf58e740377d9340fccab912b19d7341c3): complete subsection reference.

- [proper_case_header_transformation](data-sources--workload--reference--group-005.md#canonical-958c09e2eaab2cb025cb095e4e88e2d998c1cef1f88db7094aab5320194538cf): complete subsection reference.

<a id="canonical-ae8aa4cf7eb61125d932cbfe26debf59ecf05e238f1587b5d8e3889604e86524"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / f2e8dc467de7 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--workload--reference--group-005.md#canonical-72285ad952940abe4df34074485bba37ae3190a48418235fa9cfb90efe71518e)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--workload--reference--group-005.md#canonical-7e24f8a2fdd0e5aad7503285622661cf58e740377d9340fccab912b19d7341c3)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--workload--reference--group-005.md#canonical-958c09e2eaab2cb025cb095e4e88e2d998c1cef1f88db7094aab5320194538cf)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-005.md#canonical-c84f985aa9a22d077e735e9765ba2c27ca595c47bc7c85afcfc5b6ad2d143af8)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-72285ad952940abe4df34074485bba37ae3190a48418235fa9cfb90efe71518e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-528b1a9167c63f20c2a21b3f202c606acd9a8e872119f013f433df50f32b5b95"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 2d4509151b29 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-d1b08993403cab5df5c13d77dd4886b4bc149ebb413780deb8732e32538d3f11)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-005.md#canonical-e93a4484ae0ff3a088ee961ce53b6e26b6fe43173eb5e3b76d003330eee577b2)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-005.md#canonical-c84f985aa9a22d077e735e9765ba2c27ca595c47bc7c85afcfc5b6ad2d143af8)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-005.md#canonical-63ff932f43849ea58268252bacbdc7aec7c3cc2ba0bfd36728c39006978f8b1f)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-a9f058108a32febc7ff7cce93ef786f318fdabecf7f3ad0e02e9f2f9c519aeef"></a>

Type: `["object", {}]`. Computed.

Use the platform's current default HTTP header transformation behavior.

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

<a id="canonical-c23a391fc1e0c291d2e6c67610880d7f85df0ff51385150f60332cc60248a989"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 2d4509151b29 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e6133ab67d46b6e3009d148566ec303a83c813e28741381d26a32fd7f5fa9d2a"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 2d4509151b29 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-005.md#canonical-63ff932f43849ea58268252bacbdc7aec7c3cc2ba0bfd36728c39006978f8b1f)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-7e24f8a2fdd0e5aad7503285622661cf58e740377d9340fccab912b19d7341c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b30684bf529a0f5a7ef29289e12657cea404d2974d31d94662f15fd98f465743"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 320c0072a44a / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-d1b08993403cab5df5c13d77dd4886b4bc149ebb413780deb8732e32538d3f11)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-005.md#canonical-e93a4484ae0ff3a088ee961ce53b6e26b6fe43173eb5e3b76d003330eee577b2)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-005.md#canonical-c84f985aa9a22d077e735e9765ba2c27ca595c47bc7c85afcfc5b6ad2d143af8)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-005.md#canonical-63ff932f43849ea58268252bacbdc7aec7c3cc2ba0bfd36728c39006978f8b1f)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-1f305fb5be7f20ee23fa6774c35ee064c81b5d41966a239d475646d1e1e5995e"></a>

Type: `["object", {}]`. Computed.

Preserve HTTP header-name case when upstream case must remain unchanged.

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

<a id="canonical-291b47ad172bc6f8ea7cff2a6b0d754c384ee4b29a3195d00961236763eadf94"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 320c0072a44a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f5f4127b869eac346be8e50b4d8fe9b959811115c98f2776eb24a390354caaaf"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 320c0072a44a / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-005.md#canonical-63ff932f43849ea58268252bacbdc7aec7c3cc2ba0bfd36728c39006978f8b1f)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-958c09e2eaab2cb025cb095e4e88e2d998c1cef1f88db7094aab5320194538cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-05dc2f9523e9feef072e1246760317c56008b57dd33388a4714f195aaafaf993"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / e2702c88b0e7 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-d1b08993403cab5df5c13d77dd4886b4bc149ebb413780deb8732e32538d3f11)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-005.md#canonical-e93a4484ae0ff3a088ee961ce53b6e26b6fe43173eb5e3b76d003330eee577b2)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-005.md#canonical-c84f985aa9a22d077e735e9765ba2c27ca595c47bc7c85afcfc5b6ad2d143af8)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-005.md#canonical-63ff932f43849ea58268252bacbdc7aec7c3cc2ba0bfd36728c39006978f8b1f)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-288ea6b9f92328e7bc74771b30240fb9e6636656cc7bf408ac96ea1827f3765d"></a>

Type: `["object", {}]`. Computed.

Transform HTTP header names to proper case when explicit transformation is required.

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

<a id="canonical-d975a596a5127e0a1c44ffc35b8e0c05d8752f0fa0e1f4d69a9480856df29c2c"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / e2702c88b0e7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-48e8789bf53c159ca9140593139a3ae58028d683ce2545143d74ecd6a64de9b2"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / e2702c88b0e7 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-005.md#canonical-63ff932f43849ea58268252bacbdc7aec7c3cc2ba0bfd36728c39006978f8b1f)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-4a38b6c802cd47806d9d63f5751ecb16d014b6ec7ae64ec70e7e8d315cca53ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fca35ea10716f5d30658e09233165f1266df2d461cf3cc9be0282bb6a7d7577a"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2 — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / bcd829312c8f / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-d1b08993403cab5df5c13d77dd4886b4bc149ebb413780deb8732e32538d3f11)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-005.md#canonical-e93a4484ae0ff3a088ee961ce53b6e26b6fe43173eb5e3b76d003330eee577b2)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-a6a5266c311b5d3e2d49eb7da10bd72670a3687d9055603ff6275cd54b389d3c"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v1 v2.

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

<a id="canonical-1ad04f16e9ea16640cc5c63680efde297b6f78c6a56bb1e0247492c574021dab"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / bcd829312c8f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1cb8787ecc42f405c0a29f88525dfe3edcb3c829cf595ab7b61e58ba2e220632"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / bcd829312c8f / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-005.md#canonical-e93a4484ae0ff3a088ee961ce53b6e26b6fe43173eb5e3b76d003330eee577b2)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-09efa3a9dccbe6b69648c54d95e300c31d985062672a7071dc3bdf61b8074856"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f2b6360ec392c0b7b97f52aa6870559f64e0a8903755ccfa7c9d49e07b6e3bd7"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 45929f16f594 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-d1b08993403cab5df5c13d77dd4886b4bc149ebb413780deb8732e32538d3f11)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-005.md#canonical-e93a4484ae0ff3a088ee961ce53b6e26b6fe43173eb5e3b76d003330eee577b2)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-7b0b183c8df7875d127ac8e1c5d248e4d1ce604654296fb59159e6f66ba6a82d"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v2 only.

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

<a id="canonical-4311285b631079960037a4e09ede7bfeeeac74d05a5a070152bafc3430216b85"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 45929f16f594 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e107412e65341b523f408609cb803c826f24d241a4d6f7570cc01838f30d7868"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 45929f16f594 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-005.md#canonical-e93a4484ae0ff3a088ee961ce53b6e26b6fe43173eb5e3b76d003330eee577b2)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-f57f61c397bb7032e3acb8d3efc7093f3cf3e9976abba49aec384be7e244760a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a2177e31537336d6f3a02644025558638fa7a0d5d5aecf9031efd803fe7aac15"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.non_default_loadbalancer — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.non_def / 124761695d74 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-d1b08993403cab5df5c13d77dd4886b4bc149ebb413780deb8732e32538d3f11)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.non_default_loadbalancer

<a id="canonical-777169bbdff8f463f82cc4295578e01b611432822702fdea1c2487818cdc9bdd"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for non default loadbalancer.

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

<a id="canonical-5cf2ff574e7f2da8f97644b36d1435505d3fb1720ed5822ebd8acd1381eb84e7"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.non_def / 124761695d74 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bc62771444c149052ce120e83276a67e9abfee28f1c249f1b8eb0b0edcde246b"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.non_def / 124761695d74 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-d1b08993403cab5df5c13d77dd4886b4bc149ebb413780deb8732e32538d3f11)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-e06926e1633305c31d13ace68ff2e8c8cebd468c6d4ff002248512eb512e33cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2f792338c5f735e5de9e45e4875ff47459ec69268ff3465e0a526d08c8c4e75e"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.pass_through — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.pass_th / f9292ad0dd45 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-d1b08993403cab5df5c13d77dd4886b4bc149ebb413780deb8732e32538d3f11)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.pass_through

<a id="canonical-17f39651b638f66db8ca9ebfa34ea28bf53a49eee3a63546283888eca4d1649d"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for pass through.

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

<a id="canonical-73fd80687ba2dcd70d050d935c356dee2a3921c0bbf3c123e9fddce96cd6bb9b"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.pass_th / f9292ad0dd45 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a9e424db744ed04276bf3ded1835a127a47173afd142e1cb7a34ec7c271f627f"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.pass_th / f9292ad0dd45 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-d1b08993403cab5df5c13d77dd4886b4bc149ebb413780deb8732e32538d3f11)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-12b8761594eb6c09dcd0306020eb8588653df046eb14cf94f3fa3bdf773bdb0f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-81c0cf92745aa4fcaff79abf514e14efeaf1a49d022a3850eac31a69178b5552"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 5b6393aa24d4 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-d1b08993403cab5df5c13d77dd4886b4bc149ebb413780deb8732e32538d3f11)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params

<a id="canonical-c03cd3656b42cf3cca0c3c1587b02739c3535cbe49e88f058c67530e3c0ed1b1"></a>

Type: `"single"`. Computed.

Configuration parameter for tls cert params.

Upstream description:

Select TLS Parameters and Certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

<a id="canonical-d86c5a3e7286faf56a5b27fe956976562772296338abd99b18a369761187e895"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 5b6393aa24d4 / 3

- [certificates](data-sources--workload--reference--group-005.md#canonical-6f4fab47de49d6180b0abf9403c22b09fc74ee99f4e1329afd30f884677bdcc5): complete subsection reference.

- [no_mtls](data-sources--workload--reference--group-005.md#canonical-9c9fc1e191266932d53214debffec945c6307402ad5223c6424a8f75e282b7de): complete subsection reference.

- [tls_config](data-sources--workload--reference--group-006.md#canonical-82c8d47a6e1da863cfe89250b2cd57e25b543d6f719aefdf9dfe1a8cbf492fc1): complete subsection reference.

- [use_mtls](data-sources--workload--reference--group-006.md#canonical-9b7aac1436f033df4af49aa607a42ea8a86091cb1594a22731c538f8c391cdb3): complete subsection reference.

<a id="canonical-91b2d643b37e0feed14f5e4df96575c0a90d0824f8c61ea7119d663a86dd1197"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 5b6393aa24d4 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates](data-sources--workload--reference--group-005.md#canonical-6f4fab47de49d6180b0abf9403c22b09fc74ee99f4e1329afd30f884677bdcc5)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.no_mtls](data-sources--workload--reference--group-005.md#canonical-9c9fc1e191266932d53214debffec945c6307402ad5223c6424a8f75e282b7de)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-006.md#canonical-82c8d47a6e1da863cfe89250b2cd57e25b543d6f719aefdf9dfe1a8cbf492fc1)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-006.md#canonical-9b7aac1436f033df4af49aa607a42ea8a86091cb1594a22731c538f8c391cdb3)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-d1b08993403cab5df5c13d77dd4886b4bc149ebb413780deb8732e32538d3f11)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-6f4fab47de49d6180b0abf9403c22b09fc74ee99f4e1329afd30f884677bdcc5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-973a49da2ea3eb9a850b2fe95012e7f718f4ed96834ab7d728583070d803add5"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / e181361062d2 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-fe3574b4f2af2fc6d8f8000c8ef6e8baa5bd8847f507f148a99a7c6940e27c62)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-d2f2640345dbdf8a6ca20b7b3a2ceb6c31007e44a4f25546827758c4a23df912)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-7aee7d3f01b8589d73b60b2b9ba278f9df11405cd3f693ab4300b3a9cfbf189c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-d1b08993403cab5df5c13d77dd4886b4bc149ebb413780deb8732e32538d3f11)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-005.md#canonical-12b8761594eb6c09dcd0306020eb8588653df046eb14cf94f3fa3bdf773bdb0f)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates

<a id="canonical-0f35236ea8403a2c2f75af4b80b0efad3988a36339a1abd62b510dfb09fa66e1"></a>

Type: `"list"`. Computed.

Select one or more certificates with any domain names.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-33a1c691a28bf5fcbd88e35a440b88f0222677c2c8e6555e477f68c305ea74fe"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / e181361062d2 / 3

<a id="canonical-95e4b55f21ef8d8ca4ecb50ed3e0cf372b808e9893fadc5f6c907241548179f8"></a>

<a id="canonical-a9f2850a811e85bbcd0ecf2be7a2c7a196f9e6c01b7d1bffe5a010137420b7ad"></a>

## name property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / e181361062d2 / 4

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

<a id="canonical-2b2a8c2c5399ac80d0b6ee31fe52509140e8a55f2c4b770b59020cbf386fecd9"></a>

<a id="canonical-e2bbde6ff67fb3b6aee26f6e72eae8caa25fb7e34edc0db9259a77191c4417c4"></a>

## namespace property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / e181361062d2 / 5

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

<a id="canonical-9a493557a7d15896758661db85d82b9bf100488ba657b35870ce59d2d6246b18"></a>

<a id="canonical-84646b0f88ab9d7b87fc77e6dee0ab7d56173bdedbfd77a988475c686ebe57c1"></a>

## tenant property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / e181361062d2 / 6

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

<a id="canonical-4be9355f3ea217a721676c6756d47979ff950b3c0f3a469aed6afd7beb976154"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / e181361062d2 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-005.md#canonical-12b8761594eb6c09dcd0306020eb8588653df046eb14cf94f3fa3bdf773bdb0f)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-9c9fc1e191266932d53214debffec945c6307402ad5223c6424a8f75e282b7de"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
