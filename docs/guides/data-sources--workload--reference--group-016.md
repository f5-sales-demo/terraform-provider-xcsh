---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-8d078afca4b9dfecbacf62b35f4d1693cc04c178b7b38fab1e6c220cf237e708"></a>

## service.volumes — service.volumes / 22291bd8c1fa / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- service.volumes

<a id="canonical-ac7d147e5eb234df4049ef23627de8914afdc333024df1db9eb8bce224d7a70e"></a>

Type: `"list"`. Computed.

Volumes. Volumes for the service.

Upstream description:

Volumes for the service.

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

<a id="canonical-e573daca049cfaea65f68b475cd72f9b03a4231a64eb48180eb5714c680bf897"></a>

## Direct properties — service.volumes / 22291bd8c1fa / 3

- [empty_dir](data-sources--workload--reference--group-016.md#canonical-024da00517ac8cecbeba98f0d762786c03e061834fb20feb30eee69dea5c3e3a): complete subsection reference.

- [host_path](data-sources--workload--reference--group-016.md#canonical-e8895f1c343930a41aca95caca6b9502656c3930d99fd8c9ecda6daea70b371d): complete subsection reference.

<a id="canonical-54f8697bf267649fc2e7bc190d8b7cb5132bf2d62ed823af673b561c90982efe"></a>

<a id="canonical-35d2f5b11cb80cf0614902209de7ea973e2ca5863c29054c115e02d9607dfc4c"></a>

## name property — service.volumes / 22291bd8c1fa / 4

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

- [persistent_volume](data-sources--workload--reference--group-016.md#canonical-3b452d1a91fa286b2238ce5c6a1014fb4e9111ecbd0eb9bfd3e8cdaff3d4343d): complete subsection reference.

<a id="canonical-4ef6439c72ca5d1c9d5d4a98f2f18da6c08fb2a5488b75ac60b2e3129116daf0"></a>

## Next pages — service.volumes / 22291bd8c1fa / 5

- [service.volumes.empty_dir](data-sources--workload--reference--group-016.md#canonical-024da00517ac8cecbeba98f0d762786c03e061834fb20feb30eee69dea5c3e3a)
- [service.volumes.host_path](data-sources--workload--reference--group-016.md#canonical-e8895f1c343930a41aca95caca6b9502656c3930d99fd8c9ecda6daea70b371d)
- [service.volumes.persistent_volume](data-sources--workload--reference--group-016.md#canonical-3b452d1a91fa286b2238ce5c6a1014fb4e9111ecbd0eb9bfd3e8cdaff3d4343d)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-024da00517ac8cecbeba98f0d762786c03e061834fb20feb30eee69dea5c3e3a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2cc2e9b382624985cee5ab247a37e10b72d883c22f9525cb647ef3c3336cd6bd"></a>

## service.volumes.empty_dir — service.volumes.empty_dir / b4713ccd6e43 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.volumes](data-sources--workload--reference--group-015.md#canonical-8650f561bb6addd9f137c4e4b67ba52ffb9efd5338ba2a38ed06a7df7fb0e7b4)
- service.volumes.empty_dir

<a id="canonical-37c2a130b4ab500a93e2f1c2d77759767b31551c64e0c2ed3a1c60ead252b389"></a>

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

<a id="canonical-7eb28e734c9a8d3f0c92ea7b8462b6594b7018925f2bd6d704c14124e2a8d8ff"></a>

## Direct properties — service.volumes.empty_dir / b4713ccd6e43 / 3

- [mount](data-sources--workload--reference--group-016.md#canonical-76ac39547613187b9c566257a3e632de4258051e93da9a0b35c311f91cb7834d): complete subsection reference.

<a id="canonical-a9d6e7cd4db70c8d305824bb6444ddc4580dad1d6960c83d1da642c2369dc482"></a>

<a id="canonical-cb9bfd820ce8522a720ee5da8ab354962c4dfcb72191ef26e747b3b30c02dba1"></a>

## size_limit property — service.volumes.empty_dir / b4713ccd6e43 / 4

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

<a id="canonical-431ad4f244f46941e96145cb53ba8813bd3180bd941ae8160d4d3c1a2d7e4915"></a>

## Next pages — service.volumes.empty_dir / b4713ccd6e43 / 5

- [service.volumes.empty_dir.mount](data-sources--workload--reference--group-016.md#canonical-76ac39547613187b9c566257a3e632de4258051e93da9a0b35c311f91cb7834d)
- [service.volumes](data-sources--workload--reference--group-015.md#canonical-8650f561bb6addd9f137c4e4b67ba52ffb9efd5338ba2a38ed06a7df7fb0e7b4)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-76ac39547613187b9c566257a3e632de4258051e93da9a0b35c311f91cb7834d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c11edbf303a24103c98ff6bc6c20db3891f03bf04360c63140a22752a06c054"></a>

## service.volumes.empty_dir.mount — service.volumes.empty_dir.mount / d165d1d7772a / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.volumes](data-sources--workload--reference--group-015.md#canonical-8650f561bb6addd9f137c4e4b67ba52ffb9efd5338ba2a38ed06a7df7fb0e7b4)
- [service.volumes.empty_dir](data-sources--workload--reference--group-016.md#canonical-024da00517ac8cecbeba98f0d762786c03e061834fb20feb30eee69dea5c3e3a)
- service.volumes.empty_dir.mount

<a id="canonical-cfa77a9ef092448ccb6aa84d898b08d9aaf1ecdb3b461381328ee837a1c9f498"></a>

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

<a id="canonical-4fdf99632cf59062db52f6b21602e53884f1a14deccbd7f7d1b19c809947484c"></a>

## Direct properties — service.volumes.empty_dir.mount / d165d1d7772a / 3

<a id="canonical-178a0e0c36c3ba063fff9c359a04244dceca9247c1e72b59dc2bcde3f2113e9b"></a>

<a id="canonical-512c45a349814825866ca96350eab4c5e015ea4c7c873f437df373648da0cccb"></a>

## mode property — service.volumes.empty_dir.mount / d165d1d7772a / 4

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

<a id="canonical-e4fad6416cd86fc5d9200c92d351b4358067142351c948bcc280745f018ed05d"></a>

<a id="canonical-ecd3a43fcc096286f3d51dde64bafcf12ebc552999ebac3d7a921f8d418c6a3e"></a>

## mount_path property — service.volumes.empty_dir.mount / d165d1d7772a / 5

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

<a id="canonical-db64db9d57541fb95f9958b2d4ae5bf5083ddccba4543d515a3abc8a3497e7c2"></a>

<a id="canonical-8823fc866fe76d33cdd7d6979485b2c633fa5d59dc0306a877c5dd36a8571adb"></a>

## sub_path property — service.volumes.empty_dir.mount / d165d1d7772a / 6

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

<a id="canonical-0cd7c1c2269f9987fc0553beac8c08f761b4f31edd72f5c4b532f10f7f2d6f46"></a>

## Next pages — service.volumes.empty_dir.mount / d165d1d7772a / 7

- [service.volumes.empty_dir](data-sources--workload--reference--group-016.md#canonical-024da00517ac8cecbeba98f0d762786c03e061834fb20feb30eee69dea5c3e3a)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-e8895f1c343930a41aca95caca6b9502656c3930d99fd8c9ecda6daea70b371d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af5b9568da01c0db134325e03562f080824ab665777387c30430fc8daa3c5d70"></a>

## service.volumes.host_path — service.volumes.host_path / 5a2be0454555 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.volumes](data-sources--workload--reference--group-015.md#canonical-8650f561bb6addd9f137c4e4b67ba52ffb9efd5338ba2a38ed06a7df7fb0e7b4)
- service.volumes.host_path

<a id="canonical-b2e5fc758886030c0e5af154c41b5e0ee763002df25a68b1cf36fa11b701ce9d"></a>

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

<a id="canonical-6ae4b5f9d615ad564cba43dd7642caf2cc3d9e3c6180da5c627013717d38afb7"></a>

## Direct properties — service.volumes.host_path / 5a2be0454555 / 3

- [mount](data-sources--workload--reference--group-016.md#canonical-3c6ce2aa94687fecf6167b9ab1c5ad0f743ff01a5883f8a135ff504ef0284182): complete subsection reference.

<a id="canonical-be84d6678983f1d04ecbedd5511485f9e909c8f898e7917645f554bd9763bbf3"></a>

<a id="canonical-2ba1d54840a3a8afa8655ad2f52858e97603fcec63a834fc5ec901a1d6b1474d"></a>

## path property — service.volumes.host_path / 5a2be0454555 / 4

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

<a id="canonical-59406b8017cb4c80f6a4ab683d2c2a547da69b2b96ba72cc6a225f1de83c72f3"></a>

## Next pages — service.volumes.host_path / 5a2be0454555 / 5

- [service.volumes.host_path.mount](data-sources--workload--reference--group-016.md#canonical-3c6ce2aa94687fecf6167b9ab1c5ad0f743ff01a5883f8a135ff504ef0284182)
- [service.volumes](data-sources--workload--reference--group-015.md#canonical-8650f561bb6addd9f137c4e4b67ba52ffb9efd5338ba2a38ed06a7df7fb0e7b4)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-3c6ce2aa94687fecf6167b9ab1c5ad0f743ff01a5883f8a135ff504ef0284182"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23576bb1c8884e73ecf26ddea3b3d05e56338d2715048db8d40f19ae1fc6b725"></a>

## service.volumes.host_path.mount — service.volumes.host_path.mount / 9470d3638a2a / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.volumes](data-sources--workload--reference--group-015.md#canonical-8650f561bb6addd9f137c4e4b67ba52ffb9efd5338ba2a38ed06a7df7fb0e7b4)
- [service.volumes.host_path](data-sources--workload--reference--group-016.md#canonical-e8895f1c343930a41aca95caca6b9502656c3930d99fd8c9ecda6daea70b371d)
- service.volumes.host_path.mount

<a id="canonical-3a3833ce5290ce03d0deebd2b35b7d0e47cd729cc9da9c9c8dc51ec0759aaeda"></a>

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

<a id="canonical-a56389572a3813b8a92ef42317a3f0d0968ca05db2ebbe24afa560059dedaf2f"></a>

## Direct properties — service.volumes.host_path.mount / 9470d3638a2a / 3

<a id="canonical-c0873bc854f128837e777351b2ab0b0bc61db839d56bf3f14190b4b67dd8666e"></a>

<a id="canonical-3e4548bcdbb333babfc02c50ea9cae5d53d399f23de8c153bc66ae977b572e88"></a>

## mode property — service.volumes.host_path.mount / 9470d3638a2a / 4

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

<a id="canonical-7e835d83d6f8b54337022e93786933eb5346b0481bbd971ef992c0488c90fcce"></a>

<a id="canonical-aad12a69b0fee7d5add923b280287248b38d294467e2d04ca7496ea2bc2d4fec"></a>

## mount_path property — service.volumes.host_path.mount / 9470d3638a2a / 5

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

<a id="canonical-62e604dc2b298a7351ddf84b3d96577de7ed73748c15d37cc8adf35f6a05b2f4"></a>

<a id="canonical-0a707b3627f547a7aea3c11bd898532dc871d259bfe6f53b5ce4b0ef521c9d02"></a>

## sub_path property — service.volumes.host_path.mount / 9470d3638a2a / 6

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

<a id="canonical-600996a380f304b77401ed3213a99e4f1a845b5c1311e683c73fb75afe82b874"></a>

## Next pages — service.volumes.host_path.mount / 9470d3638a2a / 7

- [service.volumes.host_path](data-sources--workload--reference--group-016.md#canonical-e8895f1c343930a41aca95caca6b9502656c3930d99fd8c9ecda6daea70b371d)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-3b452d1a91fa286b2238ce5c6a1014fb4e9111ecbd0eb9bfd3e8cdaff3d4343d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aeb65326c7653b5fc6ac304b0f1287a44c118b88fc1e79530633c1e164423eb6"></a>

## service.volumes.persistent_volume — service.volumes.persistent_volume / 15c079c0bbf8 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.volumes](data-sources--workload--reference--group-015.md#canonical-8650f561bb6addd9f137c4e4b67ba52ffb9efd5338ba2a38ed06a7df7fb0e7b4)
- service.volumes.persistent_volume

<a id="canonical-5da4c0f09ed99deb019235a175b3613865bdae06bcab2e97b2e93c4da71282dd"></a>

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

<a id="canonical-be4801dc61b61e1bd70494316707364bfdccf51a9f867ac8795e3d0e2151e6fd"></a>

## Direct properties — service.volumes.persistent_volume / 15c079c0bbf8 / 3

- [mount](data-sources--workload--reference--group-016.md#canonical-e466a070b13131b9c4aa3c4039c34c7205aae0926deec4545c63ef479a343b6a): complete subsection reference.

- [storage](data-sources--workload--reference--group-016.md#canonical-c4d2dad21a98e9bee5df349fc270920e2210686744ac28a2f98a67abe9c52335): complete subsection reference.

<a id="canonical-ffa4c3b2785ad7903aeab5746f8251c2c5f569ee58969adbd3c159d152614cbc"></a>

## Next pages — service.volumes.persistent_volume / 15c079c0bbf8 / 4

- [service.volumes.persistent_volume.mount](data-sources--workload--reference--group-016.md#canonical-e466a070b13131b9c4aa3c4039c34c7205aae0926deec4545c63ef479a343b6a)
- [service.volumes.persistent_volume.storage](data-sources--workload--reference--group-016.md#canonical-c4d2dad21a98e9bee5df349fc270920e2210686744ac28a2f98a67abe9c52335)
- [service.volumes](data-sources--workload--reference--group-015.md#canonical-8650f561bb6addd9f137c4e4b67ba52ffb9efd5338ba2a38ed06a7df7fb0e7b4)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-e466a070b13131b9c4aa3c4039c34c7205aae0926deec4545c63ef479a343b6a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-666b9dc98b0e6a8803c9a9afc7dcc38fc3902c1fa0594f60f738abfd9e057088"></a>

## service.volumes.persistent_volume.mount — service.volumes.persistent_volume.mount / ad7ae384c2d7 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.volumes](data-sources--workload--reference--group-015.md#canonical-8650f561bb6addd9f137c4e4b67ba52ffb9efd5338ba2a38ed06a7df7fb0e7b4)
- [service.volumes.persistent_volume](data-sources--workload--reference--group-016.md#canonical-3b452d1a91fa286b2238ce5c6a1014fb4e9111ecbd0eb9bfd3e8cdaff3d4343d)
- service.volumes.persistent_volume.mount

<a id="canonical-7527be50c442599371c85ef82b893d3fdafb026e6bd432b9ae13658e2e12f2af"></a>

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

<a id="canonical-2ad5b3038a10f010f37aab9ab3ddc9bb7c4c9e1d8ffcde5744c89cdf35599a28"></a>

## Direct properties — service.volumes.persistent_volume.mount / ad7ae384c2d7 / 3

<a id="canonical-c447968b14a83e1f8095d5d72081cf37a322dc498c428cc90fe3044559c3160c"></a>

<a id="canonical-6aba64a92e9752c8d031a05fddea4ee82e4d1c7283d3bf2864f806af2e7f069a"></a>

## mode property — service.volumes.persistent_volume.mount / ad7ae384c2d7 / 4

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

<a id="canonical-6d995b7ba7d6be22305214e9efd4276401312fb9903eff79c673eedd5345645b"></a>

<a id="canonical-42055fb6db6558a8d9bca07f658e4ad4188da3f3de8d5218e136e92595695d52"></a>

## mount_path property — service.volumes.persistent_volume.mount / ad7ae384c2d7 / 5

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

<a id="canonical-cbe7f3a8bf1e089137dd84b4b787d81f32788e491a478d23d1b54d0099cb0105"></a>

<a id="canonical-70f6f2c81bb12a48f09ecd80aad593fe2b36d5a75491bdacc7a331a4179c3120"></a>

## sub_path property — service.volumes.persistent_volume.mount / ad7ae384c2d7 / 6

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

<a id="canonical-dd558df0cc91bfcf6d41b64a6ae123675043d4a7cf02e384af563115c7e2869a"></a>

## Next pages — service.volumes.persistent_volume.mount / ad7ae384c2d7 / 7

- [service.volumes.persistent_volume](data-sources--workload--reference--group-016.md#canonical-3b452d1a91fa286b2238ce5c6a1014fb4e9111ecbd0eb9bfd3e8cdaff3d4343d)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-c4d2dad21a98e9bee5df349fc270920e2210686744ac28a2f98a67abe9c52335"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-60e13dcdc22e7e60d8de4472f81cf7f289d55b92878771dbc39ec4bf2b52827d"></a>

## service.volumes.persistent_volume.storage — service.volumes.persistent_volume.storage / 629560a92554 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.volumes](data-sources--workload--reference--group-015.md#canonical-8650f561bb6addd9f137c4e4b67ba52ffb9efd5338ba2a38ed06a7df7fb0e7b4)
- [service.volumes.persistent_volume](data-sources--workload--reference--group-016.md#canonical-3b452d1a91fa286b2238ce5c6a1014fb4e9111ecbd0eb9bfd3e8cdaff3d4343d)
- service.volumes.persistent_volume.storage

<a id="canonical-81b53f461da0101e64c4f2b1b1a11ba78ee08eacf3edc81d1850edd5d03bc127"></a>

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

<a id="canonical-3ea4aaaaf50a12f2f0defde19da16b07b14fb065bd1d18a9570eddeb9081c975"></a>

## Direct properties — service.volumes.persistent_volume.storage / 629560a92554 / 3

<a id="canonical-05386c5b9982ccddc1bf34393dee668f0c93211e4d6044882fbb4445550550b3"></a>

<a id="canonical-c67ce75ea762c51552a4283edce9b26def2523814ad469779a888cc12ef42ecd"></a>

## access_mode property — service.volumes.persistent_volume.storage / 629560a92554 / 4

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

<a id="canonical-650548cab158651223052e33ade286fbef89ba33c1db03e7a0d082562f582e8b"></a>

<a id="canonical-4e69ce367680191b3eac32a0eca231b3623caab90e0bc752767ea8da471528c8"></a>

## class_name property — service.volumes.persistent_volume.storage / 629560a92554 / 5

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

- [default](data-sources--workload--reference--group-016.md#canonical-5d236bbd0c7bb0a8ff9e02c5969c60f397b210896a29c2e1d0a5f88852d17f9c): complete subsection reference.

<a id="canonical-b4c3f45d6ad25bde00a042f4d9c60f783dfe863a178401bba76cacf4094f714b"></a>

<a id="canonical-e748347c74c5703692b1985855dbcaacebd44286159099f049a738ded5399270"></a>

## storage_size property — service.volumes.persistent_volume.storage / 629560a92554 / 6

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

<a id="canonical-e29f72306643b62fa212c35f2e363fab96f5dd17e5d0d68577198969127c3911"></a>

## Next pages — service.volumes.persistent_volume.storage / 629560a92554 / 7

- [service.volumes.persistent_volume.storage.default](data-sources--workload--reference--group-016.md#canonical-5d236bbd0c7bb0a8ff9e02c5969c60f397b210896a29c2e1d0a5f88852d17f9c)
- [service.volumes.persistent_volume](data-sources--workload--reference--group-016.md#canonical-3b452d1a91fa286b2238ce5c6a1014fb4e9111ecbd0eb9bfd3e8cdaff3d4343d)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-5d236bbd0c7bb0a8ff9e02c5969c60f397b210896a29c2e1d0a5f88852d17f9c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1cf4fbe96ec79bc5460ceb43ce5d5d1bf5cb8e548bd1dcd12cdc3b301f76770e"></a>

## service.volumes.persistent_volume.storage.default — service.volumes.persistent_volume.storage.default / db5f02edd5b3 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.volumes](data-sources--workload--reference--group-015.md#canonical-8650f561bb6addd9f137c4e4b67ba52ffb9efd5338ba2a38ed06a7df7fb0e7b4)
- [service.volumes.persistent_volume](data-sources--workload--reference--group-016.md#canonical-3b452d1a91fa286b2238ce5c6a1014fb4e9111ecbd0eb9bfd3e8cdaff3d4343d)
- [service.volumes.persistent_volume.storage](data-sources--workload--reference--group-016.md#canonical-c4d2dad21a98e9bee5df349fc270920e2210686744ac28a2f98a67abe9c52335)
- service.volumes.persistent_volume.storage.default

<a id="canonical-d23d51eb242f3ed4715dcc87a2c83821c7c22d525e32b5b3c077debc74402145"></a>

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

<a id="canonical-9005eae09d828d08bde1425c42e969c35949ab0bde8b6d2a5150e3d6355a173f"></a>

## Direct properties — service.volumes.persistent_volume.storage.default / db5f02edd5b3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d466068b3013bb328186c2579561a7a3c5e625df38d3c2bd8542b27bf3642cf9"></a>

## Next pages — service.volumes.persistent_volume.storage.default / db5f02edd5b3 / 4

- [service.volumes.persistent_volume.storage](data-sources--workload--reference--group-016.md#canonical-c4d2dad21a98e9bee5df349fc270920e2210686744ac28a2f98a67abe9c52335)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-71960fd031859dbf944439defbb2a20d47baca84e7e4237c03adaefc73c0b1b2"></a>

## simple_service — simple_service / 9eec67e7cd94 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- simple_service

<a id="canonical-d9c526e6e39b1aaa28a718b61c37d668d7dacf4f27e2c10a2696b36ae0055554"></a>

Type: `"single"`. Computed.

SimpleService is a service having one container and one replica that is deployed on all Regional
Edges and advertised on Internet via HTTP loadbalancer on default VIP.

Upstream description:

SimpleService is a service having one container and one replica that is deployed on all Regional
Edges and advertised on Internet via HTTP loadbalancer on default VIP.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-advertise_choice": "[\"do_not_advertise\",\"simple_advertise\"]",
  "x-ves-oneof-field-persistence_choice": "[\"disabled\",\"enabled\"]"
}
```

<a id="canonical-50938c60871b086d8f5a0ed0d96537506e6bace364e7b6a8b16718e269a60a40"></a>

## Direct properties — simple_service / 9eec67e7cd94 / 3

- [configuration](data-sources--workload--reference--group-016.md#canonical-a8e67973086bdf37258bf1a40831a2bc6a2794cc03f145c95e45dbe3c46b031a): complete subsection reference.

- [container](data-sources--workload--reference--group-016.md#canonical-d89cfacac18a20187ddde73c5218eb8e7e4eb6f9a61040eedbae8e8f4ea82b6b): complete subsection reference.

- [disabled](data-sources--workload--reference--group-016.md#canonical-964dfbbca6df028d9e306c830facf5ffb5c7b283146bb44d117200d3a3502bd7): complete subsection reference.

- [do_not_advertise](data-sources--workload--reference--group-016.md#canonical-64a37c27c7f70d1aae07e783b2b158932c18f9c1e7a59236f35371db51188328): complete subsection reference.

- [enabled](data-sources--workload--reference--group-016.md#canonical-c19f09398561cccba5070e877e29bf882e759cf21d0165215f0da663e688010c): complete subsection reference.

<a id="canonical-b13e9bc09c197e47c12cb75d62a92c10cae05c7da059cc4eae76db4ddae78687"></a>

<a id="canonical-acd06f2570a8bedfcbf804d5ab5c448264adc89ebbdb5a39866636ee909647d6"></a>

## scale_to_zero property — simple_service / 9eec67e7cd94 / 4

Type: `"bool"`. Computed.

Scale down replicas of the service to zero.

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

- [simple_advertise](data-sources--workload--reference--group-016.md#canonical-7e2c8a2c2dd86a775ee7ba82abccf291a8cac9413bf85ab6b62762694e25956c): complete subsection reference.

<a id="canonical-aff8b8caeb2d6ed6e09a013e51a2fe2b80e30b27a92dde8807d8ce918deeaf60"></a>

## Next pages — simple_service / 9eec67e7cd94 / 5

- [simple_service.configuration](data-sources--workload--reference--group-016.md#canonical-a8e67973086bdf37258bf1a40831a2bc6a2794cc03f145c95e45dbe3c46b031a)
- [simple_service.container](data-sources--workload--reference--group-016.md#canonical-d89cfacac18a20187ddde73c5218eb8e7e4eb6f9a61040eedbae8e8f4ea82b6b)
- [simple_service.disabled](data-sources--workload--reference--group-016.md#canonical-964dfbbca6df028d9e306c830facf5ffb5c7b283146bb44d117200d3a3502bd7)
- [simple_service.do_not_advertise](data-sources--workload--reference--group-016.md#canonical-64a37c27c7f70d1aae07e783b2b158932c18f9c1e7a59236f35371db51188328)
- [simple_service.enabled](data-sources--workload--reference--group-016.md#canonical-c19f09398561cccba5070e877e29bf882e759cf21d0165215f0da663e688010c)
- [simple_service.simple_advertise](data-sources--workload--reference--group-016.md#canonical-7e2c8a2c2dd86a775ee7ba82abccf291a8cac9413bf85ab6b62762694e25956c)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-a8e67973086bdf37258bf1a40831a2bc6a2794cc03f145c95e45dbe3c46b031a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b358ef9c809cb69c7f721197ecd93cf0bcd2a58d3391a38065920dbd2b38c54"></a>

## simple_service.configuration — simple_service.configuration / b92cc3c586f9 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- simple_service.configuration

<a id="canonical-11fe2d0b1465e23fd88c1fd435d48adfd6fbdb980812ea3dcc7a026489a2b945"></a>

Type: `"single"`. Computed.

Configuration parameters of the workload.

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

<a id="canonical-0723361be41d864581ed5d27754c4842f09ffe0bd5af20bed38b12204f08f3e4"></a>

## Direct properties — simple_service.configuration / b92cc3c586f9 / 3

- [parameters](data-sources--workload--reference--group-016.md#canonical-76c90aa9f4bdcb36c56b6879c86cef126bb8dd9295e57d79561708e2bc2cf6e5): complete subsection reference.

<a id="canonical-17264c90cf947f8df10c24aeaf15c88e8ad92a246b0325a1b76b3cdfcd8023b3"></a>

## Next pages — simple_service.configuration / b92cc3c586f9 / 4

- [simple_service.configuration.parameters](data-sources--workload--reference--group-016.md#canonical-76c90aa9f4bdcb36c56b6879c86cef126bb8dd9295e57d79561708e2bc2cf6e5)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-76c90aa9f4bdcb36c56b6879c86cef126bb8dd9295e57d79561708e2bc2cf6e5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1026f3872c5705624e14104d83d2758e342b8e656f6a51970ab18f96c11a8ca3"></a>

## simple_service.configuration.parameters — simple_service.configuration.parameters / a7f756c69e93 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- [simple_service.configuration](data-sources--workload--reference--group-016.md#canonical-a8e67973086bdf37258bf1a40831a2bc6a2794cc03f145c95e45dbe3c46b031a)
- simple_service.configuration.parameters

<a id="canonical-4befc21123d8e9031e0a05375b66356b9a0d9b685500441206629a070722c7a7"></a>

Type: `"list"`. Computed.

Parameters. Parameters for the workload.

Upstream description:

Parameters for the workload.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-9dc421c5cfad08725b211e6bb1124309dcff4a60db62ce6e5966f15388caec07"></a>

## Direct properties — simple_service.configuration.parameters / a7f756c69e93 / 3

- [env_var](data-sources--workload--reference--group-016.md#canonical-354ed87ca389e85ead3b05821d66e5a4ff03d2f78a2870340dfafc20d9732273): complete subsection reference.

- [file](data-sources--workload--reference--group-016.md#canonical-c22f4f4705e64aa6a86b14da2097f6df7f45a2460239165c0b3509904b67065e): complete subsection reference.

<a id="canonical-5d007716b9fa52a853c50744a840e94dc4e7b1e5d69559f6bfb7e278cdac5479"></a>

## Next pages — simple_service.configuration.parameters / a7f756c69e93 / 4

- [simple_service.configuration.parameters.env_var](data-sources--workload--reference--group-016.md#canonical-354ed87ca389e85ead3b05821d66e5a4ff03d2f78a2870340dfafc20d9732273)
- [simple_service.configuration.parameters.file](data-sources--workload--reference--group-016.md#canonical-c22f4f4705e64aa6a86b14da2097f6df7f45a2460239165c0b3509904b67065e)
- [simple_service.configuration](data-sources--workload--reference--group-016.md#canonical-a8e67973086bdf37258bf1a40831a2bc6a2794cc03f145c95e45dbe3c46b031a)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-354ed87ca389e85ead3b05821d66e5a4ff03d2f78a2870340dfafc20d9732273"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f895044ebf8540c1f8c7b2d0f18935b9fed4c36c4e05b241d8fc980ede4d302"></a>

## simple_service.configuration.parameters.env_var — simple_service.configuration.parameters.env_var / 3a5b176f12c3 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- [simple_service.configuration](data-sources--workload--reference--group-016.md#canonical-a8e67973086bdf37258bf1a40831a2bc6a2794cc03f145c95e45dbe3c46b031a)
- [simple_service.configuration.parameters](data-sources--workload--reference--group-016.md#canonical-76c90aa9f4bdcb36c56b6879c86cef126bb8dd9295e57d79561708e2bc2cf6e5)
- simple_service.configuration.parameters.env_var

<a id="canonical-bdb89e14c6b591ddacb1b2b3624e6778bee2850eadcd6b4f537d57a378b845a6"></a>

Type: `"single"`. Computed.

Environment Variable. Environment Variable.

Upstream description:

Environment Variable.

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

<a id="canonical-b254c8664ca45c48e0da30cb2c805c054ffc7d34b71a5ec42b393438f06cb1a2"></a>

## Direct properties — simple_service.configuration.parameters.env_var / 3a5b176f12c3 / 3

<a id="canonical-2855bf0553b50f420fc1366ff91ca3f40036cba30ef721783e4d20168dd74292"></a>

<a id="canonical-3f4b7e8ed51d4b15dbadd32b74a8d28250438c164ec166fe7dd6e2bf8951b85f"></a>

## name property — simple_service.configuration.parameters.env_var / 3a5b176f12c3 / 4

Type: `"string"`. Computed.

Name. Name of Environment Variable.

Upstream description:

Name of Environment Variable.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-40ab2ef0f9b5a1d69c41833fa54c9c393a36306c98b9b660a0e66a753d663cbd"></a>

<a id="canonical-f25091d3aeaa5993c34daab87ae1cffbafe0086b355003fe28caebee792d164f"></a>

## value property — simple_service.configuration.parameters.env_var / 3a5b176f12c3 / 5

Type: `"string"`. Computed.

Value. Value of Environment Variable.

Upstream description:

Value of Environment Variable.

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

<a id="canonical-72397e7b49a73b6d4bcb2694af7ad99e49722311e3d2eea377a68d6281d6a4f3"></a>

## Next pages — simple_service.configuration.parameters.env_var / 3a5b176f12c3 / 6

- [simple_service.configuration.parameters](data-sources--workload--reference--group-016.md#canonical-76c90aa9f4bdcb36c56b6879c86cef126bb8dd9295e57d79561708e2bc2cf6e5)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-c22f4f4705e64aa6a86b14da2097f6df7f45a2460239165c0b3509904b67065e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b8355f27785bb5791391803ff69e06011cf79b10dc4f2c9c3920061ab1f75a84"></a>

## simple_service.configuration.parameters.file — simple_service.configuration.parameters.file / ecf298d8a291 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- [simple_service.configuration](data-sources--workload--reference--group-016.md#canonical-a8e67973086bdf37258bf1a40831a2bc6a2794cc03f145c95e45dbe3c46b031a)
- [simple_service.configuration.parameters](data-sources--workload--reference--group-016.md#canonical-76c90aa9f4bdcb36c56b6879c86cef126bb8dd9295e57d79561708e2bc2cf6e5)
- simple_service.configuration.parameters.file

<a id="canonical-632b9e2f144060d3017dfb785dcd6d28302fde5d8b0efc1f52bc8527da596fa9"></a>

Type: `"single"`. Computed.

Configuration File. Configuration File for the workload.

Upstream description:

Configuration File for the workload.

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

<a id="canonical-b8c5dd51d93fe0d9e68e5015bddafbb8c20d92d1815d25ff682495174d0699d2"></a>

## Direct properties — simple_service.configuration.parameters.file / ecf298d8a291 / 3

<a id="canonical-d2bb2bb4953d4af30cff0dd238b24f0ef35ab68dfa39b63552df258ca4b8e913"></a>

<a id="canonical-25cc9d5cf9bcf7cd66e26c14ebc55f025e7f17eab6da5339cfdfd544492c1788"></a>

## data property — simple_service.configuration.parameters.file / ecf298d8a291 / 4

Type: `"string"`. Computed.

Data. File data

Upstream description:

File data

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 16384,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 16384,
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
    "ves.io.schema.rules.string.max_len": "16384",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "16384",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [mount](data-sources--workload--reference--group-016.md#canonical-648a15d8617efd8739760c3009d82828c6ff37b77b5efac3e6813a94b98e5819): complete subsection reference.

<a id="canonical-979cb13d637cfcf2705d707820f9189a1c295141f270a6561668519a6e3384b8"></a>

<a id="canonical-c87ba51cb67ab7d20d32840fd5debcdd2f1bde5bcba1f1a048a62479e1eb31e4"></a>

## name property — simple_service.configuration.parameters.file / ecf298d8a291 / 5

Type: `"string"`. Computed.

Name. Name of the file.

Upstream description:

Name of the file.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-c241b4e9c3f5fdd6b2f2363ed2132c2adaccdb9b2a09a535c8eca5b37d06cff6"></a>

<a id="canonical-7e2641d4e082d5f5c825bf0ee61dd872fd21f2fff75093689ca3ce6d3bce497c"></a>

## volume_name property — simple_service.configuration.parameters.file / ecf298d8a291 / 6

Type: `"string"`. Computed.

Volume Name. Name of the Volume.

Upstream description:

Name of the Volume.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-ff09b1c48d2077f4aa06d06d14fbfb73339389fefe7117b0fbef2b61f7937115"></a>

## Next pages — simple_service.configuration.parameters.file / ecf298d8a291 / 7

- [simple_service.configuration.parameters.file.mount](data-sources--workload--reference--group-016.md#canonical-648a15d8617efd8739760c3009d82828c6ff37b77b5efac3e6813a94b98e5819)
- [simple_service.configuration.parameters](data-sources--workload--reference--group-016.md#canonical-76c90aa9f4bdcb36c56b6879c86cef126bb8dd9295e57d79561708e2bc2cf6e5)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-648a15d8617efd8739760c3009d82828c6ff37b77b5efac3e6813a94b98e5819"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b69de51ddd4ad7501f8f08a890f609ba28eedf7252353ee68ea0794f14f4b3f6"></a>

## simple_service.configuration.parameters.file.mount — simple_service.configuration.parameters.file.mount / 780a3a9dfe79 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- [simple_service.configuration](data-sources--workload--reference--group-016.md#canonical-a8e67973086bdf37258bf1a40831a2bc6a2794cc03f145c95e45dbe3c46b031a)
- [simple_service.configuration.parameters](data-sources--workload--reference--group-016.md#canonical-76c90aa9f4bdcb36c56b6879c86cef126bb8dd9295e57d79561708e2bc2cf6e5)
- [simple_service.configuration.parameters.file](data-sources--workload--reference--group-016.md#canonical-c22f4f4705e64aa6a86b14da2097f6df7f45a2460239165c0b3509904b67065e)
- simple_service.configuration.parameters.file.mount

<a id="canonical-02928bd84e9c459d1c7d298372265ed160b7bd654db2a69dc76e1467d7317dc4"></a>

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

<a id="canonical-256f53d2239d1245cee8773e8e9f355d88d954c833879ce52395865e61c11db0"></a>

## Direct properties — simple_service.configuration.parameters.file.mount / 780a3a9dfe79 / 3

<a id="canonical-b6cbc3267e9e1169bc0e2731aa99b99307ec27aa5e5222654dfed297d861ba24"></a>

<a id="canonical-ee6c0f549da22818cd2c989b4227afaf5898a9c336f9837ab35bd74b09bf8435"></a>

## mode property — simple_service.configuration.parameters.file.mount / 780a3a9dfe79 / 4

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

<a id="canonical-ea5ee34ac7a7a5d23f9f0c7c71df9e2f2a78122dbcdb53c252754cc2e5b5f526"></a>

<a id="canonical-0acd1c3dd435e8dcc844ff673d11507c84e0f7cba894dafe51f25bb38b01de14"></a>

## mount_path property — simple_service.configuration.parameters.file.mount / 780a3a9dfe79 / 5

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

<a id="canonical-90600de96c8857cf2d98508d38618f46de01bd9cf89fd6bf640178a28c6fc243"></a>

<a id="canonical-bc70ffda1004ff53bb70b1c3962432b0d7639d54c8261505a9a2ff4925534c7c"></a>

## sub_path property — simple_service.configuration.parameters.file.mount / 780a3a9dfe79 / 6

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

<a id="canonical-011a87695c8b1d3719c918c5bc8a734cdf59a0526c9cdbba05b8220a6d3bdf63"></a>

## Next pages — simple_service.configuration.parameters.file.mount / 780a3a9dfe79 / 7

- [simple_service.configuration.parameters.file](data-sources--workload--reference--group-016.md#canonical-c22f4f4705e64aa6a86b14da2097f6df7f45a2460239165c0b3509904b67065e)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-d89cfacac18a20187ddde73c5218eb8e7e4eb6f9a61040eedbae8e8f4ea82b6b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d148707ec78461c4c11d5aa644214b7f5f0d2ba2ec650c4d945919df15ebf35e"></a>

## simple_service.container — simple_service.container / c37f50f80f2e / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- simple_service.container

<a id="canonical-b0ed494b9e4e790109e9073f2127e29a6b277e8f7f99786f66678d2d05223c3a"></a>

Type: `"single"`. Computed.

ContainerType configures the container information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-flavor_choice": "[\"custom_flavor\",\"default_flavor\",\"flavor\"]"
}
```

<a id="canonical-a3cedb82823986a6fd05a54b20a8c82ddd88eab6555c33ace366697544246e77"></a>

## Direct properties — simple_service.container / c37f50f80f2e / 3

<a id="canonical-0cf77946f7d8d2d7f22d0e36bab25417cc51fb300d28987d61a318a8342d3615"></a>

<a id="canonical-4672251f589d9099c35c99c52e155ad99192f00782337c0f881e63ee1edad8aa"></a>

## args property — simple_service.container / c37f50f80f2e / 4

Type: `["list", "string"]`. Computed.

Arguments to the entrypoint. Overrides the docker image's CMD.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-d015f6de4034320b4496355d1657d5bf88f14c06961932eba79bd3806cff4789"></a>

<a id="canonical-7a180f887a06e0d2d0d21218a7a17f464e21d361ef6dd0a58796133ac155c61f"></a>

## command property — simple_service.container / c37f50f80f2e / 5

Type: `["list", "string"]`. Computed.

Command to execute. Overrides the docker image's ENTRYPOINT.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

- [custom_flavor](data-sources--workload--reference--group-016.md#canonical-3a5dc266180e7aaba225b3ba95854be64c9b87428a668023c6db7169bc2ffe5e): complete subsection reference.

- [default_flavor](data-sources--workload--reference--group-016.md#canonical-e172b5fc0572e1788a3d450588e62359a7cf434c3e94fc525838bddefa305001): complete subsection reference.

<a id="canonical-ff9fc5aab754e9f00a5da8b22bcc1688612ebcc7eaa83380a2c74e54d61cb199"></a>

<a id="canonical-de001fc6e6f2ff1655f9cfa67422065a61450cf2a00b992e1f9d849b7e57199f"></a>

## flavor property — simple_service.container / c37f50f80f2e / 6

Type: `"string"`. Computed.

\[Enum:
CONTAINER\_FLAVOR\_TYPE\_TINY|CONTAINER\_FLAVOR\_TYPE\_MEDIUM|CONTAINER\_FLAVOR\_TYPE\_LARGE\]
Container Flavor type - CONTAINER\_FLAVOR\_TYPE\_TINY: Tiny Tiny containers have limit of 0.1 vCPU
and 256 MiB (mebibyte) memory - CONTAINER\_FLAVOR\_TYPE\_MEDIUM: Medium Medium containers have limit
of 0.25 vCPU and 512 MiB (mebibyte) memory - CONTAINER\_FLAVOR\_TYPE\_LARGE: Large Large containers
have.. Possible values are \`CONTAINER\_FLAVOR\_TYPE\_TINY\`, \`CONTAINER\_FLAVOR\_TYPE\_MEDIUM\`,
\`CONTAINER\_FLAVOR\_TYPE\_LARGE\`. Defaults to \`CONTAINER\_FLAVOR\_TYPE\_TINY\`.

Upstream description:

Container Flavor type

&#8203;- CONTAINER\_FLAVOR\_TYPE\_TINY: Tiny

Tiny containers have limit of 0.1 vCPU and 256 MiB (mebibyte) memory &#8203;-
CONTAINER\_FLAVOR\_TYPE\_MEDIUM: Medium

Medium containers have limit of 0.25 vCPU and 512 MiB (mebibyte) memory &#8203;-
CONTAINER\_FLAVOR\_TYPE\_LARGE: Large

Large containers have limit of 1 vCPU and 2048 MiB (mebibyte) memory.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTAINER_FLAVOR_TYPE_TINY",
  "enum": [
    "CONTAINER_FLAVOR_TYPE_TINY",
    "CONTAINER_FLAVOR_TYPE_MEDIUM",
    "CONTAINER_FLAVOR_TYPE_LARGE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [image](data-sources--workload--reference--group-016.md#canonical-38cc62aa121ef839ef942c9b3d2064494706e4b02f489a31ffe26f444d554c23): complete subsection reference.

<a id="canonical-a57fb7799bb72998c762d2593a5e3f7a25f77953b9685dba5896ed8dbfae1551"></a>

<a id="canonical-7dc09c8b5f7d66990335a852300f984200a420c45a3d0914f2ab0d3e43e4450d"></a>

## init_container property — simple_service.container / c37f50f80f2e / 7

Type: `"bool"`. Computed.

Specialized container that runs before application container and runs to completion.

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

- [liveness_check](data-sources--workload--reference--group-016.md#canonical-40078366e56bd2b4da247b8efc748dabea92272e2ae589d8520a460a98fcdb13): complete subsection reference.

<a id="canonical-93393a1692e193d63866bead4f5f8158407977affe8bcb18e3a45e5eb0db3b18"></a>

<a id="canonical-e765c46255f6f9c3e1c36bfdce2e719a8f8a406be4df490a10e3053a940efcfa"></a>

## name property — simple_service.container / c37f50f80f2e / 8

Type: `"string"`. Computed.

Name. Name of the container.

Upstream description:

Name of the container.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [readiness_check](data-sources--workload--reference--group-016.md#canonical-36e993b318b755c2eeeff2b737c6b884176375afe202c97a6b8d3fc3ab1ea8a3): complete subsection reference.

<a id="canonical-a13f518c81d9b38a4befe01f7f252cb6a6cfb11fe6fb8dbc681dcabf413cd0e4"></a>

## Next pages — simple_service.container / c37f50f80f2e / 9

- [simple_service.container.custom_flavor](data-sources--workload--reference--group-016.md#canonical-3a5dc266180e7aaba225b3ba95854be64c9b87428a668023c6db7169bc2ffe5e)
- [simple_service.container.default_flavor](data-sources--workload--reference--group-016.md#canonical-e172b5fc0572e1788a3d450588e62359a7cf434c3e94fc525838bddefa305001)
- [simple_service.container.image](data-sources--workload--reference--group-016.md#canonical-38cc62aa121ef839ef942c9b3d2064494706e4b02f489a31ffe26f444d554c23)
- [simple_service.container.liveness_check](data-sources--workload--reference--group-016.md#canonical-40078366e56bd2b4da247b8efc748dabea92272e2ae589d8520a460a98fcdb13)
- [simple_service.container.readiness_check](data-sources--workload--reference--group-016.md#canonical-36e993b318b755c2eeeff2b737c6b884176375afe202c97a6b8d3fc3ab1ea8a3)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-3a5dc266180e7aaba225b3ba95854be64c9b87428a668023c6db7169bc2ffe5e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b10ffcd9a919574d7fc3ceb57a76b3ac145212ff089f7f85abfff1eb968997a4"></a>

## simple_service.container.custom_flavor — simple_service.container.custom_flavor / 16fac0a8badb / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- [simple_service.container](data-sources--workload--reference--group-016.md#canonical-d89cfacac18a20187ddde73c5218eb8e7e4eb6f9a61040eedbae8e8f4ea82b6b)
- simple_service.container.custom_flavor

<a id="canonical-fb3659529166ca6b26359a7c1373ae3fd6d5ef6e5a1b3b9b10bda8bcd1d6f72a"></a>

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

<a id="canonical-114087b239f059005f1a9243b06b86915bc4146c2f332a3b07dbb32d1848c47c"></a>

## Direct properties — simple_service.container.custom_flavor / 16fac0a8badb / 3

<a id="canonical-9f2d8e0095a4929f62dffebcc337fa9e4772430a39bc5cccad269bcd3de4129a"></a>

<a id="canonical-a03ae77c02b923696b184b5f55455c73e3bdddeb334f31b05bfacce8b0bef4dc"></a>

## name property — simple_service.container.custom_flavor / 16fac0a8badb / 4

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

<a id="canonical-a89d39fb3eaf971d9a49d585b9e1045eef1f1410ba881d46095ed6d8de70dff5"></a>

<a id="canonical-45cbce78f7941bf9ee56fff85242a31b91f1ba4e013e5bbefceb4176685a10ea"></a>

## namespace property — simple_service.container.custom_flavor / 16fac0a8badb / 5

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

<a id="canonical-ec0463b9df79f0f1ece08184538caa82be4d806b33d4696dd7cadebd2fe2c37b"></a>

<a id="canonical-b72441d42f0fea50ca2a87880d8cdf20145801d69311421622128e07dae3847b"></a>

## tenant property — simple_service.container.custom_flavor / 16fac0a8badb / 6

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

<a id="canonical-45430133fdde4cde77b3225a97cb5f45e2fc7fdff4a8945e77b620ea39882395"></a>

## Next pages — simple_service.container.custom_flavor / 16fac0a8badb / 7

- [simple_service.container](data-sources--workload--reference--group-016.md#canonical-d89cfacac18a20187ddde73c5218eb8e7e4eb6f9a61040eedbae8e8f4ea82b6b)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-e172b5fc0572e1788a3d450588e62359a7cf434c3e94fc525838bddefa305001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e844db92d94750d9689cfbf95f64efb79320d8e1c725d2b5a07340e7e0679a6"></a>

## simple_service.container.default_flavor — simple_service.container.default_flavor / d7728427bbd8 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- [simple_service.container](data-sources--workload--reference--group-016.md#canonical-d89cfacac18a20187ddde73c5218eb8e7e4eb6f9a61040eedbae8e8f4ea82b6b)
- simple_service.container.default_flavor

<a id="canonical-8c1a80764386be892cbea0261a9bd1e2d6432fff9d6e2bc9b66ee86fd45d8fe4"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default flavor.

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

<a id="canonical-956babb92eef76c98da3b5989db8dcc878320bb7d3be7439c84c1c7ce1b79d0c"></a>

## Direct properties — simple_service.container.default_flavor / d7728427bbd8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-11121fdff6ea851e916bc62301e30d09c5df19908e54d1f6f17eef21fc77d610"></a>

## Next pages — simple_service.container.default_flavor / d7728427bbd8 / 4

- [simple_service.container](data-sources--workload--reference--group-016.md#canonical-d89cfacac18a20187ddde73c5218eb8e7e4eb6f9a61040eedbae8e8f4ea82b6b)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-38cc62aa121ef839ef942c9b3d2064494706e4b02f489a31ffe26f444d554c23"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e3b04716a7e7c5191f9036e5fdcc1bd7f4b7744ace12937a449acc7f650dc20"></a>

## simple_service.container.image — simple_service.container.image / 856b348c736f / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- [simple_service.container](data-sources--workload--reference--group-016.md#canonical-d89cfacac18a20187ddde73c5218eb8e7e4eb6f9a61040eedbae8e8f4ea82b6b)
- simple_service.container.image

<a id="canonical-8a3f6cad33ee5547d77ac1b7475826c9c7464e2568947839f1612f4d80f4595f"></a>

Type: `"single"`. Computed.

ImageType configures the image to use, how to pull the image, and the associated secrets to use if
any.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-registry_choice": "[\"container_registry\",\"public\"]"
}
```

<a id="canonical-9e2b761e6297482f58dad35d77b453f50ca4e70410240046f4ca0e14df0fa1d8"></a>

## Direct properties — simple_service.container.image / 856b348c736f / 3

- [container_registry](data-sources--workload--reference--group-016.md#canonical-d07f42bdd7ae798ca02f25093cee7d03a5c22ddfe5fe532e3537e8e5a27745bc): complete subsection reference.

<a id="canonical-cceae3b06d22f3720607abb0c079c73f4f8a9726f08e29cbc7ed5e1e493c8d03"></a>

<a id="canonical-e44725c96d770771b0780d33aee0d6ed99b756d5a453a70624d05e2423a01afb"></a>

## name property — simple_service.container.image / 856b348c736f / 4

Type: `"string"`. Computed.

Name is a container image which are usually given a name such as alpine, ubuntu, or
quay.I/O/etcd:0.13. The format is registry/image:tag or registry/image@image-digest. If registry is
not specified, the Docker public registry is assumed.

Upstream description:

Name is a container image which are usually given a name such as alpine, ubuntu, or
quay.I/O/etcd:0.13. The format is registry/image:tag or registry/image@image-digest. If registry is
not specified, the Docker public registry is assumed. If tag is not specified, latest is assumed.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [public](data-sources--workload--reference--group-016.md#canonical-0bf9bbfb8282deb7d893fa36ba998db8aed73b0b46818e08c9f4e577f41996e8): complete subsection reference.

<a id="canonical-3a180d40b6654777b5c642d9ecae45b6ddf6724adab1a1c2ac0ada7eae8b8fce"></a>

<a id="canonical-b3bde55c8027b75083d2e2e369d5d048a0a0efdd8d9d021bdb34934cfae08feb"></a>

## pull_policy property — simple_service.container.image / 856b348c736f / 5

Type: `"string"`. Computed.

\[Enum:
IMAGE\_PULL\_POLICY\_DEFAULT|IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT|IMAGE\_PULL\_POLICY\_ALWAYS|IMAGE\_PULL\_POLICY\_NEVER\]
Image pull policy type enumerates the policy choices to use for pulling the image prior to starting
the workload - IMAGE\_PULL\_POLICY\_DEFAULT: Default Default will always pull image if :latest tag
is specified in image name. If :latest tag is not specified in image name, it will pull image only..
Possible values are \`IMAGE\_PULL\_POLICY\_DEFAULT\`, \`IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT\`,
\`IMAGE\_PULL\_POLICY\_ALWAYS\`, \`IMAGE\_PULL\_POLICY\_NEVER\`. Defaults to
\`IMAGE\_PULL\_POLICY\_DEFAULT\`.

Upstream description:

Image pull policy type enumerates the policy choices to use for pulling the image prior to starting
the workload

&#8203;- IMAGE\_PULL\_POLICY\_DEFAULT: Default

Default will always pull image if :latest tag is specified in image name. If :latest tag is not
specified in image name, it will pull image only if it does not already exist on the node &#8203;-
IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT: IfNotPresent

Only pull the image if it does not already exist on the node &#8203;- IMAGE\_PULL\_POLICY\_ALWAYS:
Always

Always pull the image &#8203;- IMAGE\_PULL\_POLICY\_NEVER: Never

Never pull the image.

Receipt-pinned upstream constraints:

```json
{
  "default": "IMAGE_PULL_POLICY_DEFAULT",
  "enum": [
    "IMAGE_PULL_POLICY_DEFAULT",
    "IMAGE_PULL_POLICY_IF_NOT_PRESENT",
    "IMAGE_PULL_POLICY_ALWAYS",
    "IMAGE_PULL_POLICY_NEVER"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-f7606c444420dd9b7f17df0a4829aecfd9ed8b81bf0fff8a0c7ba772d53065ce"></a>

## Next pages — simple_service.container.image / 856b348c736f / 6

- [simple_service.container.image.container_registry](data-sources--workload--reference--group-016.md#canonical-d07f42bdd7ae798ca02f25093cee7d03a5c22ddfe5fe532e3537e8e5a27745bc)
- [simple_service.container.image.public](data-sources--workload--reference--group-016.md#canonical-0bf9bbfb8282deb7d893fa36ba998db8aed73b0b46818e08c9f4e577f41996e8)
- [simple_service.container](data-sources--workload--reference--group-016.md#canonical-d89cfacac18a20187ddde73c5218eb8e7e4eb6f9a61040eedbae8e8f4ea82b6b)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-d07f42bdd7ae798ca02f25093cee7d03a5c22ddfe5fe532e3537e8e5a27745bc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1deec0f0677515974327b5d7487fc3e2634d6d07ee7d63511d4a88c61de9ab7"></a>

## simple_service.container.image.container_registry — simple_service.container.image.container_registry / 88eda6bc434f / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- [simple_service.container](data-sources--workload--reference--group-016.md#canonical-d89cfacac18a20187ddde73c5218eb8e7e4eb6f9a61040eedbae8e8f4ea82b6b)
- [simple_service.container.image](data-sources--workload--reference--group-016.md#canonical-38cc62aa121ef839ef942c9b3d2064494706e4b02f489a31ffe26f444d554c23)
- simple_service.container.image.container_registry

<a id="canonical-a99136cbd71921f7c35b5f86b09c97fa1d7ac376c964c60deffef85f10b8da3d"></a>

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

<a id="canonical-bf68bc494dbf8d93c8578c9a86e477ca33f1e62e9c9eb604411c9c3bf1db3c70"></a>

## Direct properties — simple_service.container.image.container_registry / 88eda6bc434f / 3

<a id="canonical-fb61ad2a4c4b0464de2bee8fd648675441c368eba685b71fc002aa38b45b025e"></a>

<a id="canonical-ac6dfea191dce9bc7aa808973ea8cad9ae17ae398f22c2e4089d94854370381c"></a>

## name property — simple_service.container.image.container_registry / 88eda6bc434f / 4

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

<a id="canonical-7ae423f3129a87f7098ef86e9912a669db6ad27119f50dd2976ffb12528f2cbf"></a>

<a id="canonical-604a45ae6ddaee1dc878296f0d69742731ac0387e44abfe9afba69e5aef4377c"></a>

## namespace property — simple_service.container.image.container_registry / 88eda6bc434f / 5

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

<a id="canonical-7eafeed8a42fb60a284f352d5949b717d69a205c8c697440305ce585e32687f0"></a>

<a id="canonical-909a4178365878d232c8694da8a472ce06b5ee7eaa69903e207e8dacd73165e7"></a>

## tenant property — simple_service.container.image.container_registry / 88eda6bc434f / 6

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

<a id="canonical-92a13c32dbf3a0ef4cc66b1918f96a342c3460f48e454ca14f809315273661fe"></a>

## Next pages — simple_service.container.image.container_registry / 88eda6bc434f / 7

- [simple_service.container.image](data-sources--workload--reference--group-016.md#canonical-38cc62aa121ef839ef942c9b3d2064494706e4b02f489a31ffe26f444d554c23)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-0bf9bbfb8282deb7d893fa36ba998db8aed73b0b46818e08c9f4e577f41996e8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5fffbaf1cc158e44ad9103b0924abd7fc96829e75570c99b0062aeb37dcdf076"></a>

## simple_service.container.image.public — simple_service.container.image.public / 0d3b1dfe8939 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- [simple_service.container](data-sources--workload--reference--group-016.md#canonical-d89cfacac18a20187ddde73c5218eb8e7e4eb6f9a61040eedbae8e8f4ea82b6b)
- [simple_service.container.image](data-sources--workload--reference--group-016.md#canonical-38cc62aa121ef839ef942c9b3d2064494706e4b02f489a31ffe26f444d554c23)
- simple_service.container.image.public

<a id="canonical-f3d1634e67ae54a1df68a266e3741e1fd3e3ed8392c39350e3729ef24fb93158"></a>

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

<a id="canonical-069cc41bb8d765ed0e06ba9fd652ab1609b6fb13af6f435cd306c2e3c1987806"></a>

## Direct properties — simple_service.container.image.public / 0d3b1dfe8939 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a014fbeac58ed78e51278f3519aa7fb090bec6bb49057718af0dcaf4fd244417"></a>

## Next pages — simple_service.container.image.public / 0d3b1dfe8939 / 4

- [simple_service.container.image](data-sources--workload--reference--group-016.md#canonical-38cc62aa121ef839ef942c9b3d2064494706e4b02f489a31ffe26f444d554c23)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-40078366e56bd2b4da247b8efc748dabea92272e2ae589d8520a460a98fcdb13"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-725591950acd03690d37cb2bb198739681fd6fdbb650fa4fae03063254da1595"></a>

## simple_service.container.liveness_check — simple_service.container.liveness_check / efebd95fd7f0 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- [simple_service.container](data-sources--workload--reference--group-016.md#canonical-d89cfacac18a20187ddde73c5218eb8e7e4eb6f9a61040eedbae8e8f4ea82b6b)
- simple_service.container.liveness_check

<a id="canonical-99e04ed49ffa1fc9dd379f0f8ec87bd589e60c7f150f006c3692a5dc96cafeb1"></a>

Type: `"single"`. Computed.

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Upstream description:

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-health_check_choice": "[\"exec_health_check\",\"http_health_check\",\"tcp_health_check\"]"
}
```

<a id="canonical-be4cae47673a8b7867656e7ddc75c33b69a32753048d66011966c58d01d6ed82"></a>

## Direct properties — simple_service.container.liveness_check / efebd95fd7f0 / 3

- [exec_health_check](data-sources--workload--reference--group-016.md#canonical-9f96816bbc8d017048263d5726922f787218764f01e1505c9e88f2a513527fbf): complete subsection reference.

<a id="canonical-b01d50f228344b5feff92151afcfafc4755434de35bf926eb3757f3e7a604861"></a>

<a id="canonical-816f69dd0538dbeeec724c7c7804fd9cbbc8d348021683469bf182bbd0b51299"></a>

## healthy_threshold property — simple_service.container.liveness_check / efebd95fd7f0 / 4

Type: `"number"`. Computed.

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container..

Upstream description:

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container
healthy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
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
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

- [http_health_check](data-sources--workload--reference--group-016.md#canonical-ea4fa79334702fa410e95719d01a782a87b0121c079e11ef9735bc9889539456): complete subsection reference.

<a id="canonical-8199efe386ade8d49e5c4ebf2bcc30c3c6629a1e818b7bd346a038a3e006242e"></a>

<a id="canonical-66dabda157b7d943bbcd40d7e4249b430b6dc147135fecac9e944bc1a75d47e6"></a>

## initial_delay property — simple_service.container.liveness_check / efebd95fd7f0 / 5

Type: `"number"`. Computed.

Number of seconds after the container has started before health checks are initiated.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-935fb2226031da310645bd2a0f9efa74c8fcff474bffa38a8fd1f72d10915d48"></a>

<a id="canonical-c285a1fd84dcb49382a5d927dabf431a5167e4a6f4d6d03fac476f225938e749"></a>

## interval property — simple_service.container.liveness_check / efebd95fd7f0 / 6

Type: `"number"`. Computed.

Time interval in seconds between two health check requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

- [tcp_health_check](data-sources--workload--reference--group-016.md#canonical-8eb53384e88de3fdc4f89b890c8d9ee5b8f4142ea7674c6511121613662e7f18): complete subsection reference.

<a id="canonical-b5778c66c22aea122f3bef494c0c69bef77d4ced2fabce9de95cb7f1f4d489b6"></a>

<a id="canonical-29d141cdf98bc0c668d0ace4387e92551112606acb1005dacad6ca21265536fd"></a>

## timeout property — simple_service.container.liveness_check / efebd95fd7f0 / 7

Type: `"number"`. Computed.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Upstream description:

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-92c116eabf73622beff5570c595575fd17c09c2055284d9718972219ec9b6f00"></a>

<a id="canonical-1b3936760d0662e1415f2f9b368d3364b740ec530cdbe52ccd09751cab7fc0e8"></a>

## unhealthy_threshold property — simple_service.container.liveness_check / efebd95fd7f0 / 8

Type: `"number"`. Computed.

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Upstream description:

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
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
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-80d0d7f75ac34efddef3d261f6abfb330ce716e6a2598201eb77ecdffd1142d0"></a>

## Next pages — simple_service.container.liveness_check / efebd95fd7f0 / 9

- [simple_service.container.liveness_check.exec_health_check](data-sources--workload--reference--group-016.md#canonical-9f96816bbc8d017048263d5726922f787218764f01e1505c9e88f2a513527fbf)
- [simple_service.container.liveness_check.http_health_check](data-sources--workload--reference--group-016.md#canonical-ea4fa79334702fa410e95719d01a782a87b0121c079e11ef9735bc9889539456)
- [simple_service.container.liveness_check.tcp_health_check](data-sources--workload--reference--group-016.md#canonical-8eb53384e88de3fdc4f89b890c8d9ee5b8f4142ea7674c6511121613662e7f18)
- [simple_service.container](data-sources--workload--reference--group-016.md#canonical-d89cfacac18a20187ddde73c5218eb8e7e4eb6f9a61040eedbae8e8f4ea82b6b)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-9f96816bbc8d017048263d5726922f787218764f01e1505c9e88f2a513527fbf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b13f01487be864893449bc368c57645d66e84ebf6b4038e2a6fef23bc2610370"></a>

## simple_service.container.liveness_check.exec_health_check — simple_service.container.liveness_check.exec_health_check / 44e819dc4fef / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- [simple_service.container](data-sources--workload--reference--group-016.md#canonical-d89cfacac18a20187ddde73c5218eb8e7e4eb6f9a61040eedbae8e8f4ea82b6b)
- [simple_service.container.liveness_check](data-sources--workload--reference--group-016.md#canonical-40078366e56bd2b4da247b8efc748dabea92272e2ae589d8520a460a98fcdb13)
- simple_service.container.liveness_check.exec_health_check

<a id="canonical-be25d6195ecb34ed8a540b7c1d31521c93737f84b4a12e52818a835dc504d2ef"></a>

Type: `"single"`. Computed.

ExecHealthCheckType describes a health check based on 'run in container' action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Upstream description:

ExecHealthCheckType describes a health check based on "run in container" action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

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

<a id="canonical-a87e40d1c56b627219ea913495baa3cef84670b73c59754f08b1285c0367c943"></a>

## Direct properties — simple_service.container.liveness_check.exec_health_check / 44e819dc4fef / 3

<a id="canonical-b8ef06c0f764e213293a9f273040ea037516d538a1fcdd79fc2c6c8d336c1d2e"></a>

<a id="canonical-9fda9b6c1061b618433e8d15a549d8f2a68c57803a459646a4f79866ee056dbb"></a>

## command property — simple_service.container.liveness_check.exec_health_check / 44e819dc4fef / 4

Type: `["list", "string"]`. Computed.

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to..

Upstream description:

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to
explicitly call out to that shell.

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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3fca91319f7e2a640bc174b282fa8d2237c4ba0e307e65872a3a3f07411533d7"></a>

## Next pages — simple_service.container.liveness_check.exec_health_check / 44e819dc4fef / 5

- [simple_service.container.liveness_check](data-sources--workload--reference--group-016.md#canonical-40078366e56bd2b4da247b8efc748dabea92272e2ae589d8520a460a98fcdb13)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-ea4fa79334702fa410e95719d01a782a87b0121c079e11ef9735bc9889539456"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-21c35749a325eeab72590b1bfdc1f94f0374aba56b3bd6ac222ecd385788a06a"></a>

## simple_service.container.liveness_check.http_health_check — simple_service.container.liveness_check.http_health_check / b3858be48928 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- [simple_service.container](data-sources--workload--reference--group-016.md#canonical-d89cfacac18a20187ddde73c5218eb8e7e4eb6f9a61040eedbae8e8f4ea82b6b)
- [simple_service.container.liveness_check](data-sources--workload--reference--group-016.md#canonical-40078366e56bd2b4da247b8efc748dabea92272e2ae589d8520a460a98fcdb13)
- simple_service.container.liveness_check.http_health_check

<a id="canonical-465ffd9fe0947ea18097ae6d58b7881e83245baa31e4349e320dcd4a891cffa3"></a>

Type: `"single"`. Computed.

HTTPHealthCheckType describes a health check based on HTTP GET requests.

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

<a id="canonical-3e795a0178423e7375b96d0c4b7d61391d5d9df264d723a629f623d22540cac7"></a>

## Direct properties — simple_service.container.liveness_check.http_health_check / b3858be48928 / 3

<a id="canonical-e24d01dd1afe2a3a58552821f8a15b2ef547a8d30e6dd0db76e4ccc0ebacf590"></a>

<a id="canonical-e68348fca6e73d6d2f8950ed962f92a74ca4ea1091eb803c753e78d0c302f7b4"></a>

## headers property — simple_service.container.liveness_check.http_health_check / b3858be48928 / 4

Type: `["map", "string"]`. Computed.

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

Upstream description:

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

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
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-74c260c1513856c1f36c0cd22ef65dbc7cc3d4128d8c91deefe2b4581bc3f230"></a>

<a id="canonical-89ef15f094dd44b77f78f297d7113f0162f5e7370f009d77624dd2e5d6dfc64a"></a>

## host_header property — simple_service.container.liveness_check.http_health_check / b3858be48928 / 5

Type: `"string"`. Computed.

The value of the host header in the HTTP health check request.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 262,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 262,
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
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

<a id="canonical-5199cd0d94efd202c9ed505796844bfdc4dff3564ac6e818bc37ef2c7d8a0da6"></a>

<a id="canonical-96e7535cda18779c16566c008ca61991619bf0019854c45542237cab3c19e33b"></a>

## path property — simple_service.container.liveness_check.http_health_check / b3858be48928 / 6

Type: `"string"`. Computed.

Path. Path to access on the HTTP server.

Upstream description:

Path to access on the HTTP server.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

- [port](data-sources--workload--reference--group-016.md#canonical-5f3da74903321d6deeb6b1b68d0ed69ee4af95f32c8090404168d9a1b3d8a2f1): complete subsection reference.

<a id="canonical-39629a339aa2415400f8c2817106eb5540eb471c9fd393c632adb1e2b73c91ea"></a>

## Next pages — simple_service.container.liveness_check.http_health_check / b3858be48928 / 7

- [simple_service.container.liveness_check.http_health_check.port](data-sources--workload--reference--group-016.md#canonical-5f3da74903321d6deeb6b1b68d0ed69ee4af95f32c8090404168d9a1b3d8a2f1)
- [simple_service.container.liveness_check](data-sources--workload--reference--group-016.md#canonical-40078366e56bd2b4da247b8efc748dabea92272e2ae589d8520a460a98fcdb13)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-5f3da74903321d6deeb6b1b68d0ed69ee4af95f32c8090404168d9a1b3d8a2f1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5270e49f4598a5e469ebdddd04031999d0c92234530ce141e87f89fdda661747"></a>

## simple_service.container.liveness_check.http_health_check.port — simple_service.container.liveness_check.http_health_check.port / 3f085aa5d9c2 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- [simple_service.container](data-sources--workload--reference--group-016.md#canonical-d89cfacac18a20187ddde73c5218eb8e7e4eb6f9a61040eedbae8e8f4ea82b6b)
- [simple_service.container.liveness_check](data-sources--workload--reference--group-016.md#canonical-40078366e56bd2b4da247b8efc748dabea92272e2ae589d8520a460a98fcdb13)
- [simple_service.container.liveness_check.http_health_check](data-sources--workload--reference--group-016.md#canonical-ea4fa79334702fa410e95719d01a782a87b0121c079e11ef9735bc9889539456)
- simple_service.container.liveness_check.http_health_check.port

<a id="canonical-5cf04b65d3231e5df7d5c51be64883f0b4de66b16b9629a3d4c8137124a106b5"></a>

Type: `"single"`. Computed.

Port. Port

Upstream description:

Port

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

<a id="canonical-1921aaf11da82e4c60642591d0d152f8d3c326895e72dfb8e75cfdb834f136ad"></a>

## Direct properties — simple_service.container.liveness_check.http_health_check.port / 3f085aa5d9c2 / 3

<a id="canonical-c8cc3e0fcb3e8f6ad39afdbf3b0bd4eaad612199f2efb86f47878621736691e1"></a>

<a id="canonical-1476f6ea495414bec87a2e201a60acee308572f9e6a870a549ee9f3c89569a6e"></a>

## name property — simple_service.container.liveness_check.http_health_check.port / 3f085aa5d9c2 / 4

Type: `"string"`. Computed.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-100c1006615ae083b2f073f6d15388b3cd977b6000ca290990cc91e4fb16c6f0"></a>

<a id="canonical-2ed362290f177916ee663540ac05f4a00fbc0070c2601ad6e497e0c2714772f8"></a>

## num property — simple_service.container.liveness_check.http_health_check.port / 3f085aa5d9c2 / 5

Type: `"number"`. Computed.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

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

<a id="canonical-28b40f8095c155cc6cab332bbfdd2b75f8ff549fd3a78ff8ba0004d40a24786a"></a>

## Next pages — simple_service.container.liveness_check.http_health_check.port / 3f085aa5d9c2 / 6

- [simple_service.container.liveness_check.http_health_check](data-sources--workload--reference--group-016.md#canonical-ea4fa79334702fa410e95719d01a782a87b0121c079e11ef9735bc9889539456)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-8eb53384e88de3fdc4f89b890c8d9ee5b8f4142ea7674c6511121613662e7f18"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9453d26014996934df9b01548ae97a1fef1a00698d68244cad925076d84c798"></a>

## simple_service.container.liveness_check.tcp_health_check — simple_service.container.liveness_check.tcp_health_check / ad20d70559df / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- [simple_service.container](data-sources--workload--reference--group-016.md#canonical-d89cfacac18a20187ddde73c5218eb8e7e4eb6f9a61040eedbae8e8f4ea82b6b)
- [simple_service.container.liveness_check](data-sources--workload--reference--group-016.md#canonical-40078366e56bd2b4da247b8efc748dabea92272e2ae589d8520a460a98fcdb13)
- simple_service.container.liveness_check.tcp_health_check

<a id="canonical-b1fab425943cfdef346b94a6a59b4eeb24b6230b69aaacb3d272649f5f6dcf0c"></a>

Type: `"single"`. Computed.

TCPHealthCheckType describes a health check based on opening a TCP connection.

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

<a id="canonical-1ac264ee077e80436f18258f0d5223ed2bb4099f0a573ba5c8fc4f0b86aa933c"></a>

## Direct properties — simple_service.container.liveness_check.tcp_health_check / ad20d70559df / 3

- [port](data-sources--workload--reference--group-016.md#canonical-b47712d9b7257b32be0eb58320765c9b3f6518fe66e30848b64e15593f774658): complete subsection reference.

<a id="canonical-d48f637cad1d5fc930bf234bf1d4f9ae8cfee1546da7a116b42654bad5ebb6d2"></a>

## Next pages — simple_service.container.liveness_check.tcp_health_check / ad20d70559df / 4

- [simple_service.container.liveness_check.tcp_health_check.port](data-sources--workload--reference--group-016.md#canonical-b47712d9b7257b32be0eb58320765c9b3f6518fe66e30848b64e15593f774658)
- [simple_service.container.liveness_check](data-sources--workload--reference--group-016.md#canonical-40078366e56bd2b4da247b8efc748dabea92272e2ae589d8520a460a98fcdb13)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-b47712d9b7257b32be0eb58320765c9b3f6518fe66e30848b64e15593f774658"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ea5b447e2949962352f8f93cdf17412a7a7e45c0755f802f6c143aee130b1b3a"></a>

## simple_service.container.liveness_check.tcp_health_check.port — simple_service.container.liveness_check.tcp_health_check.port / 0be8473d6cae / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- [simple_service.container](data-sources--workload--reference--group-016.md#canonical-d89cfacac18a20187ddde73c5218eb8e7e4eb6f9a61040eedbae8e8f4ea82b6b)
- [simple_service.container.liveness_check](data-sources--workload--reference--group-016.md#canonical-40078366e56bd2b4da247b8efc748dabea92272e2ae589d8520a460a98fcdb13)
- [simple_service.container.liveness_check.tcp_health_check](data-sources--workload--reference--group-016.md#canonical-8eb53384e88de3fdc4f89b890c8d9ee5b8f4142ea7674c6511121613662e7f18)
- simple_service.container.liveness_check.tcp_health_check.port

<a id="canonical-a0601218cb3c6b20d0af871c399eb2ecd31c0c2ac8c42f26f6075f20ee3ae4f2"></a>

Type: `"single"`. Computed.

Port. Port

Upstream description:

Port

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

<a id="canonical-b12c6420568764a07cffd257d5e623052bfc344b6186d6ba86df22fb8554344b"></a>

## Direct properties — simple_service.container.liveness_check.tcp_health_check.port / 0be8473d6cae / 3

<a id="canonical-d198210039bacb18b319dd9946f95208778a546b7f4a04e6f3d57203e45fee24"></a>

<a id="canonical-58afb87fbeab4fdbcfafaa1fb804c0a1dd63f8b36e38c860e8d1a2422e10b7e9"></a>

## name property — simple_service.container.liveness_check.tcp_health_check.port / 0be8473d6cae / 4

Type: `"string"`. Computed.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-34cd994b1c1f2bad17a385cecda488feb6d850ba5edc399700fa695d5534d12e"></a>

<a id="canonical-46244b45da553e8ebd61a30d8b6ccaf4b271c9ab40fb17bcfc5f5845ac34e8cc"></a>

## num property — simple_service.container.liveness_check.tcp_health_check.port / 0be8473d6cae / 5

Type: `"number"`. Computed.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

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

<a id="canonical-6ab41ae35302e3001f443534990ed3186827ece7cac05018705d8a312f994f1b"></a>

## Next pages — simple_service.container.liveness_check.tcp_health_check.port / 0be8473d6cae / 6

- [simple_service.container.liveness_check.tcp_health_check](data-sources--workload--reference--group-016.md#canonical-8eb53384e88de3fdc4f89b890c8d9ee5b8f4142ea7674c6511121613662e7f18)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-36e993b318b755c2eeeff2b737c6b884176375afe202c97a6b8d3fc3ab1ea8a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fe88e13d616dbddf76bf835b48b9f130e28a03a75db3fef20beaef66f69ddce4"></a>

## simple_service.container.readiness_check — simple_service.container.readiness_check / 53ec5e3dc86e / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- [simple_service.container](data-sources--workload--reference--group-016.md#canonical-d89cfacac18a20187ddde73c5218eb8e7e4eb6f9a61040eedbae8e8f4ea82b6b)
- simple_service.container.readiness_check

<a id="canonical-971219aa75aaee291c6b21641853d38912f1f6cee8c8ebdb2bbd8d7592c9664e"></a>

Type: `"single"`. Computed.

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Upstream description:

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-health_check_choice": "[\"exec_health_check\",\"http_health_check\",\"tcp_health_check\"]"
}
```

<a id="canonical-623bd333f7f790437f9bc186d7ee793bf167320cb6956616e399afc535861947"></a>

## Direct properties — simple_service.container.readiness_check / 53ec5e3dc86e / 3

- [exec_health_check](data-sources--workload--reference--group-016.md#canonical-d8890130f68a4d4711a2210238894f08487e520b7b56f8532575dbbfe8b34869): complete subsection reference.

<a id="canonical-42be2a13d59575d2c8fa7c8cbf5a3f200a3147c2e485afdbee998e4f2f25acde"></a>

<a id="canonical-37f869d2aecc06c23cf0f0906b27cfe7bef924f47f48f0debffdcf08a9409110"></a>

## healthy_threshold property — simple_service.container.readiness_check / 53ec5e3dc86e / 4

Type: `"number"`. Computed.

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container..

Upstream description:

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container
healthy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
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
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

- [http_health_check](data-sources--workload--reference--group-016.md#canonical-88a110529fce4d7034514a9c10525aef8c0ba848b4b6b920a756c4a0c52fbb4d): complete subsection reference.

<a id="canonical-60d2b9ff99c4f8d851b153428d2286fb043fcd0c331d2956a29d62fd144b6772"></a>

<a id="canonical-a8e204753562c6584218b6b5a085104b69c8e3c7d6cd07fe84d7c2f3fee786b7"></a>

## initial_delay property — simple_service.container.readiness_check / 53ec5e3dc86e / 5

Type: `"number"`. Computed.

Number of seconds after the container has started before health checks are initiated.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-816ab23dcbdc9714303f6125e858c60779780722c7669bbf271162336c9e566b"></a>

<a id="canonical-89d03b54658cfb1157116be1de18b756315174906bfc9e9f6af3a0bd0fccd790"></a>

## interval property — simple_service.container.readiness_check / 53ec5e3dc86e / 6

Type: `"number"`. Computed.

Time interval in seconds between two health check requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

- [tcp_health_check](data-sources--workload--reference--group-016.md#canonical-839abd09fdc6c3df94c2e0ac62c5ce1eed31edff9bf9357fbe47412bb2695c9b): complete subsection reference.

<a id="canonical-b9605016973748b8ea965b8d7174c3620491f454ab795c9d5187dce160c72146"></a>

<a id="canonical-73de0df4ffe6df2cc4352c622e2e4d834e112cdbc4e3f2909974624023033713"></a>

## timeout property — simple_service.container.readiness_check / 53ec5e3dc86e / 7

Type: `"number"`. Computed.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Upstream description:

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-31f1cb27784b40cb79ef92ed375a48dfd3177f7dcced4446858074d0a2c5a513"></a>

<a id="canonical-e86d8202d26978165b3108e430fd21a135ee41d7f7090fdd80f3e9b2301919db"></a>

## unhealthy_threshold property — simple_service.container.readiness_check / 53ec5e3dc86e / 8

Type: `"number"`. Computed.

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Upstream description:

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
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
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-8dcdac5ce1c8e43217e7478d63e8cc6b7fac030a87c6ec22bd30f6957d3d19de"></a>

## Next pages — simple_service.container.readiness_check / 53ec5e3dc86e / 9

- [simple_service.container.readiness_check.exec_health_check](data-sources--workload--reference--group-016.md#canonical-d8890130f68a4d4711a2210238894f08487e520b7b56f8532575dbbfe8b34869)
- [simple_service.container.readiness_check.http_health_check](data-sources--workload--reference--group-016.md#canonical-88a110529fce4d7034514a9c10525aef8c0ba848b4b6b920a756c4a0c52fbb4d)
- [simple_service.container.readiness_check.tcp_health_check](data-sources--workload--reference--group-016.md#canonical-839abd09fdc6c3df94c2e0ac62c5ce1eed31edff9bf9357fbe47412bb2695c9b)
- [simple_service.container](data-sources--workload--reference--group-016.md#canonical-d89cfacac18a20187ddde73c5218eb8e7e4eb6f9a61040eedbae8e8f4ea82b6b)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-d8890130f68a4d4711a2210238894f08487e520b7b56f8532575dbbfe8b34869"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-80b813067d408a3a28d657af8ad1d6c08c7ecaa0494a063c5d3ee345444dcd21"></a>

## simple_service.container.readiness_check.exec_health_check — simple_service.container.readiness_check.exec_health_check / 17f1c242dbb7 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- [simple_service.container](data-sources--workload--reference--group-016.md#canonical-d89cfacac18a20187ddde73c5218eb8e7e4eb6f9a61040eedbae8e8f4ea82b6b)
- [simple_service.container.readiness_check](data-sources--workload--reference--group-016.md#canonical-36e993b318b755c2eeeff2b737c6b884176375afe202c97a6b8d3fc3ab1ea8a3)
- simple_service.container.readiness_check.exec_health_check

<a id="canonical-6c16b7e243a98b1ca804db26ee6e0d7d22351ebf53c48f595dda40be47111089"></a>

Type: `"single"`. Computed.

ExecHealthCheckType describes a health check based on 'run in container' action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Upstream description:

ExecHealthCheckType describes a health check based on "run in container" action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

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

<a id="canonical-245643d0266ab6e70a5c2bf4a2220e28536c54f9925a13d32a3d9ddf2ea5a39d"></a>

## Direct properties — simple_service.container.readiness_check.exec_health_check / 17f1c242dbb7 / 3

<a id="canonical-0c2b9c6b4472a7ea1e7420a605874da72886b7bab17a93a7e95e0bda8b8d7cb6"></a>

<a id="canonical-300bc1d88cb3dfe4b24cbd890985ca48f413ecd4e4ec1db4079694484c9b8b39"></a>

## command property — simple_service.container.readiness_check.exec_health_check / 17f1c242dbb7 / 4

Type: `["list", "string"]`. Computed.

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to..

Upstream description:

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to
explicitly call out to that shell.

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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-55b4b28e21637382136a6ec8b6b972be7fe1ae335f3783cb97d31c2c77f19c15"></a>

## Next pages — simple_service.container.readiness_check.exec_health_check / 17f1c242dbb7 / 5

- [simple_service.container.readiness_check](data-sources--workload--reference--group-016.md#canonical-36e993b318b755c2eeeff2b737c6b884176375afe202c97a6b8d3fc3ab1ea8a3)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-88a110529fce4d7034514a9c10525aef8c0ba848b4b6b920a756c4a0c52fbb4d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7818a60cbda487ad34452d6a493019949baa179615b3d3e054a7739f24d3538"></a>

## simple_service.container.readiness_check.http_health_check — simple_service.container.readiness_check.http_health_check / b525d595e1ac / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- [simple_service.container](data-sources--workload--reference--group-016.md#canonical-d89cfacac18a20187ddde73c5218eb8e7e4eb6f9a61040eedbae8e8f4ea82b6b)
- [simple_service.container.readiness_check](data-sources--workload--reference--group-016.md#canonical-36e993b318b755c2eeeff2b737c6b884176375afe202c97a6b8d3fc3ab1ea8a3)
- simple_service.container.readiness_check.http_health_check

<a id="canonical-05454894f580106486aa1c9bc8b78bdfdbe51903973a2464b6f918d6109ab9c8"></a>

Type: `"single"`. Computed.

HTTPHealthCheckType describes a health check based on HTTP GET requests.

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

<a id="canonical-1bc22a67e01119bffc2eab884ea3e7137e59c83cba78991aee4dd4ff4bf3a4eb"></a>

## Direct properties — simple_service.container.readiness_check.http_health_check / b525d595e1ac / 3

<a id="canonical-7a0a62ef7ec58ac84b73d39f5ba975e5c1f09c28a73e3b8786722d8b0396b9c8"></a>

<a id="canonical-9c864699707fbbf3627b5d78f13879fc09f76a0ed23677bfb75e8da2e3968448"></a>

## headers property — simple_service.container.readiness_check.http_health_check / b525d595e1ac / 4

Type: `["map", "string"]`. Computed.

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

Upstream description:

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

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
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-c578eb1de837881f0d295c9c4ab12f0df8015ae3d783a1aee5ca19dc6d700048"></a>

<a id="canonical-298144d555242554a320827ae06e257a114bc674c0a44355b4ea8b4c0843ee5e"></a>

## host_header property — simple_service.container.readiness_check.http_health_check / b525d595e1ac / 5

Type: `"string"`. Computed.

The value of the host header in the HTTP health check request.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 262,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 262,
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
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

<a id="canonical-cda30a6bc2b2168c8fc1460fd86f521106e664de3443ff497a479680cceac9de"></a>

<a id="canonical-588777f3450afaf1be517cc822ea0eb9ed72516522a8fdbfba2fefd536ba79c3"></a>

## path property — simple_service.container.readiness_check.http_health_check / b525d595e1ac / 6

Type: `"string"`. Computed.

Path. Path to access on the HTTP server.

Upstream description:

Path to access on the HTTP server.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

- [port](data-sources--workload--reference--group-016.md#canonical-c68c620fe61be0e345cee5e77535bc457facd5d20459f4ca6efd5881cb97558e): complete subsection reference.

<a id="canonical-734c128c001a33af8728ac608d3f1932f9e1f1ac852ad6a7deb29b16133cab95"></a>

## Next pages — simple_service.container.readiness_check.http_health_check / b525d595e1ac / 7

- [simple_service.container.readiness_check.http_health_check.port](data-sources--workload--reference--group-016.md#canonical-c68c620fe61be0e345cee5e77535bc457facd5d20459f4ca6efd5881cb97558e)
- [simple_service.container.readiness_check](data-sources--workload--reference--group-016.md#canonical-36e993b318b755c2eeeff2b737c6b884176375afe202c97a6b8d3fc3ab1ea8a3)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-c68c620fe61be0e345cee5e77535bc457facd5d20459f4ca6efd5881cb97558e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72fb4d0eeeb6b70f72b1069f8cb428a9091c3e79795ec093295f3e7e067e84d8"></a>

## simple_service.container.readiness_check.http_health_check.port — simple_service.container.readiness_check.http_health_check.port / faa3126eda27 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- [simple_service.container](data-sources--workload--reference--group-016.md#canonical-d89cfacac18a20187ddde73c5218eb8e7e4eb6f9a61040eedbae8e8f4ea82b6b)
- [simple_service.container.readiness_check](data-sources--workload--reference--group-016.md#canonical-36e993b318b755c2eeeff2b737c6b884176375afe202c97a6b8d3fc3ab1ea8a3)
- [simple_service.container.readiness_check.http_health_check](data-sources--workload--reference--group-016.md#canonical-88a110529fce4d7034514a9c10525aef8c0ba848b4b6b920a756c4a0c52fbb4d)
- simple_service.container.readiness_check.http_health_check.port

<a id="canonical-4c8b6d533f2eaf4c699df3bc8a242a0648355d0e5491299f1c8a5dfb16e90fb0"></a>

Type: `"single"`. Computed.

Port. Port

Upstream description:

Port

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

<a id="canonical-d842987043bb80d1452e5dab6fecff8451c543e9711640f317955ab027f53c75"></a>

## Direct properties — simple_service.container.readiness_check.http_health_check.port / faa3126eda27 / 3

<a id="canonical-46103a6da46c1c30010eb70cf61cff45e61ebef197b853e722e3a5e509fb77a5"></a>

<a id="canonical-f4e7b90a9c188d4e64eedcd570537bb017cc5b339dd89091fc947c25a181a506"></a>

## name property — simple_service.container.readiness_check.http_health_check.port / faa3126eda27 / 4

Type: `"string"`. Computed.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-5cea9d19d1b808a85c450755bad4958bbf11d4f7328d0e6d38cea756a140408b"></a>

<a id="canonical-58d773f192bfc751c36f4db4f1f33e85d61d8c9f4abca6a5cb6779626ca2b147"></a>

## num property — simple_service.container.readiness_check.http_health_check.port / faa3126eda27 / 5

Type: `"number"`. Computed.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

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

<a id="canonical-8047f217b9f96df31aaaa96447cd20977621ad89cef68c50f2f8cefbb68fc5ed"></a>

## Next pages — simple_service.container.readiness_check.http_health_check.port / faa3126eda27 / 6

- [simple_service.container.readiness_check.http_health_check](data-sources--workload--reference--group-016.md#canonical-88a110529fce4d7034514a9c10525aef8c0ba848b4b6b920a756c4a0c52fbb4d)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-839abd09fdc6c3df94c2e0ac62c5ce1eed31edff9bf9357fbe47412bb2695c9b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-87021132390fc6ecb2f379cf7c557d210dd77a90db4f9c29af04a8568259887a"></a>

## simple_service.container.readiness_check.tcp_health_check — simple_service.container.readiness_check.tcp_health_check / 45619f37db13 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- [simple_service.container](data-sources--workload--reference--group-016.md#canonical-d89cfacac18a20187ddde73c5218eb8e7e4eb6f9a61040eedbae8e8f4ea82b6b)
- [simple_service.container.readiness_check](data-sources--workload--reference--group-016.md#canonical-36e993b318b755c2eeeff2b737c6b884176375afe202c97a6b8d3fc3ab1ea8a3)
- simple_service.container.readiness_check.tcp_health_check

<a id="canonical-e7b71f5f0f65b36ccaf659756e3884ef872923a5cbf2af44372e5a4bb5abfa77"></a>

Type: `"single"`. Computed.

TCPHealthCheckType describes a health check based on opening a TCP connection.

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

<a id="canonical-16763f714722d9547b688c216865a2b67afb400d03a6f112606c81a98d366d87"></a>

## Direct properties — simple_service.container.readiness_check.tcp_health_check / 45619f37db13 / 3

- [port](data-sources--workload--reference--group-016.md#canonical-e33919a89cff1d786d0ed925c571fcbd54c303036fa3b7be825ac7d46f400cff): complete subsection reference.

<a id="canonical-7ef42f0f7f8cc1f8914364548eca168f67a5df2ff819c2c5380eca4d3a98784d"></a>

## Next pages — simple_service.container.readiness_check.tcp_health_check / 45619f37db13 / 4

- [simple_service.container.readiness_check.tcp_health_check.port](data-sources--workload--reference--group-016.md#canonical-e33919a89cff1d786d0ed925c571fcbd54c303036fa3b7be825ac7d46f400cff)
- [simple_service.container.readiness_check](data-sources--workload--reference--group-016.md#canonical-36e993b318b755c2eeeff2b737c6b884176375afe202c97a6b8d3fc3ab1ea8a3)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-e33919a89cff1d786d0ed925c571fcbd54c303036fa3b7be825ac7d46f400cff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c9c89186773959f70603c834867bb283617b046f9be4f7d8ab76deec9ce659b3"></a>

## simple_service.container.readiness_check.tcp_health_check.port — simple_service.container.readiness_check.tcp_health_check.port / 636a57c32240 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- [simple_service.container](data-sources--workload--reference--group-016.md#canonical-d89cfacac18a20187ddde73c5218eb8e7e4eb6f9a61040eedbae8e8f4ea82b6b)
- [simple_service.container.readiness_check](data-sources--workload--reference--group-016.md#canonical-36e993b318b755c2eeeff2b737c6b884176375afe202c97a6b8d3fc3ab1ea8a3)
- [simple_service.container.readiness_check.tcp_health_check](data-sources--workload--reference--group-016.md#canonical-839abd09fdc6c3df94c2e0ac62c5ce1eed31edff9bf9357fbe47412bb2695c9b)
- simple_service.container.readiness_check.tcp_health_check.port

<a id="canonical-06fa2f34022e494f8df3afc1704ba61d596c291889b5e08fbe93f6e27b15a6c8"></a>

Type: `"single"`. Computed.

Port. Port

Upstream description:

Port

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

<a id="canonical-41f201d44590e3ce80ba9ac42f432fb707947458081055816feff217ed02c2df"></a>

## Direct properties — simple_service.container.readiness_check.tcp_health_check.port / 636a57c32240 / 3

<a id="canonical-6b2b68a2378cae82532f3fe6bfcdfeecb6d2423ca01918747576d3a027868918"></a>

<a id="canonical-fb806fa8d41749df94c7746417224951d13c1784a97f985a0be93a8aaf5bf713"></a>

## name property — simple_service.container.readiness_check.tcp_health_check.port / 636a57c32240 / 4

Type: `"string"`. Computed.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-ee324ebcd4ec7a9c8cf66ec8cd5cbd21a576f5ce1a6f26d7281e5610c6e66f58"></a>

<a id="canonical-1cd228428b3a3d2b4e57739b219bc272f357b0874aa861716bacb6584d9f34c5"></a>

## num property — simple_service.container.readiness_check.tcp_health_check.port / 636a57c32240 / 5

Type: `"number"`. Computed.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

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

<a id="canonical-314b36420473c422c1b00bc0a622f0f76c24e3a8e884ac2aa29e3572bdb9cdff"></a>

## Next pages — simple_service.container.readiness_check.tcp_health_check.port / 636a57c32240 / 6

- [simple_service.container.readiness_check.tcp_health_check](data-sources--workload--reference--group-016.md#canonical-839abd09fdc6c3df94c2e0ac62c5ce1eed31edff9bf9357fbe47412bb2695c9b)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-964dfbbca6df028d9e306c830facf5ffb5c7b283146bb44d117200d3a3502bd7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3dd4c1847b81ebe87d227c2eff5a7b91fa5cfb0297fcb579079d635fc73f7145"></a>

## simple_service.disabled — simple_service.disabled / f90b6be4e3bb / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- simple_service.disabled

<a id="canonical-6cadd0a5cd4f1bdd34b484a19dde44a4f3e62f1568ae49f7ae0e18a634c5c6f0"></a>

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

<a id="canonical-140916ca805fa03bcffc1657c78eb6666364c366b99fe40a22990ad4ffbe91e5"></a>

## Direct properties — simple_service.disabled / f90b6be4e3bb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-207d9cb763d4d6b97c78b1eb7456e6717f85d4c25c445a960522817e747cf3ba"></a>

## Next pages — simple_service.disabled / f90b6be4e3bb / 4

- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-64a37c27c7f70d1aae07e783b2b158932c18f9c1e7a59236f35371db51188328"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-db188f91ef2b8a3156b8301525cbf754b8cb41be003b3803d5ca8ae090dded2c"></a>

## simple_service.do_not_advertise — simple_service.do_not_advertise / 64be99ef9691 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- simple_service.do_not_advertise

<a id="canonical-7c0f4a182cd855e0eb75feaecfd0c3954e3047f6d12686b09c234c78097b2baf"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for do not advertise.

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

<a id="canonical-31a586344bbe940d2514cbe12ff525321174ca449ee6049cb1db667851bf9359"></a>

## Direct properties — simple_service.do_not_advertise / 64be99ef9691 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-429e5c8bcd938efa7ef3be19cd7d4588db55e284b44cfa9f16d23b866cf9d0d5"></a>

## Next pages — simple_service.do_not_advertise / 64be99ef9691 / 4

- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-c19f09398561cccba5070e877e29bf882e759cf21d0165215f0da663e688010c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6c4fa3ce9a238cdd49d237f91f747986cf9901b9cfa15fd2dbab09b3b3b1bfa"></a>

## simple_service.enabled — simple_service.enabled / 64a0b13c1a45 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- simple_service.enabled

<a id="canonical-7a5075e62707f09e7c94638b3afcae7690f1a49c621554e3536bd6e64fc800a4"></a>

Type: `"single"`. Computed.

Persistent storage volume configuration for the workload.

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

<a id="canonical-47b913227c402f8552a43a2103dcfdf64e3d4d5447ca2e96bfc0988f28b70c8f"></a>

## Direct properties — simple_service.enabled / 64a0b13c1a45 / 3

<a id="canonical-a280113c3edc411991684b5f25fa272df3e920bde418e3e50da8724f1db8a895"></a>

<a id="canonical-873dc7562448ce71d3b36773a899d10cd1b267f5f2144a54a7e859cb7b3560d8"></a>

## name property — simple_service.enabled / 64a0b13c1a45 / 4

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

- [persistent_volume](data-sources--workload--reference--group-016.md#canonical-9b579a402ac565efcf999975aa5d1fc5d5abd09449f94bcb5713cdbeae496e7d): complete subsection reference.

<a id="canonical-1f35da44b6dd489b971aa4645f304fdb37e5979e4e6cc63fb235200d8b096cd8"></a>

## Next pages — simple_service.enabled / 64a0b13c1a45 / 5

- [simple_service.enabled.persistent_volume](data-sources--workload--reference--group-016.md#canonical-9b579a402ac565efcf999975aa5d1fc5d5abd09449f94bcb5713cdbeae496e7d)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-9b579a402ac565efcf999975aa5d1fc5d5abd09449f94bcb5713cdbeae496e7d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2bd6e6d7662ba2ea21ac8efd0307509eb9e2adfe4a0207a13abbe7cdd39a65ba"></a>

## simple_service.enabled.persistent_volume — simple_service.enabled.persistent_volume / b6860da19c4e / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- [simple_service.enabled](data-sources--workload--reference--group-016.md#canonical-c19f09398561cccba5070e877e29bf882e759cf21d0165215f0da663e688010c)
- simple_service.enabled.persistent_volume

<a id="canonical-12073088abc9825a6bfadb0e238aa011b8d58eaa7ae1d24ce97464dfe4746680"></a>

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

<a id="canonical-943b70f1ce0fe64ecb605412b45531737ed06b6665b31f90c5d83656acc6722a"></a>

## Direct properties — simple_service.enabled.persistent_volume / b6860da19c4e / 3

- [mount](data-sources--workload--reference--group-016.md#canonical-59d0a3a5fa4236cc9c3a8a6e43ccea82af4477428e7e01370fa4f0334eebe47f): complete subsection reference.

- [storage](data-sources--workload--reference--group-016.md#canonical-ea59ad8c1077d994f543ce2ba96b2f8ea9ffbeda24c4e57701bdc6d3d502eeaf): complete subsection reference.

<a id="canonical-4b7b1bc8c3fd4b292ede421528f2296eeaea23614f4a80cb9ab43309adb935b8"></a>

## Next pages — simple_service.enabled.persistent_volume / b6860da19c4e / 4

- [simple_service.enabled.persistent_volume.mount](data-sources--workload--reference--group-016.md#canonical-59d0a3a5fa4236cc9c3a8a6e43ccea82af4477428e7e01370fa4f0334eebe47f)
- [simple_service.enabled.persistent_volume.storage](data-sources--workload--reference--group-016.md#canonical-ea59ad8c1077d994f543ce2ba96b2f8ea9ffbeda24c4e57701bdc6d3d502eeaf)
- [simple_service.enabled](data-sources--workload--reference--group-016.md#canonical-c19f09398561cccba5070e877e29bf882e759cf21d0165215f0da663e688010c)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-59d0a3a5fa4236cc9c3a8a6e43ccea82af4477428e7e01370fa4f0334eebe47f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7c1b7791b2aa8565485259026dcdaec91710325b15bad54a8a8742483c866892"></a>

## simple_service.enabled.persistent_volume.mount — simple_service.enabled.persistent_volume.mount / 9f7db65e6cda / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- [simple_service.enabled](data-sources--workload--reference--group-016.md#canonical-c19f09398561cccba5070e877e29bf882e759cf21d0165215f0da663e688010c)
- [simple_service.enabled.persistent_volume](data-sources--workload--reference--group-016.md#canonical-9b579a402ac565efcf999975aa5d1fc5d5abd09449f94bcb5713cdbeae496e7d)
- simple_service.enabled.persistent_volume.mount

<a id="canonical-5c848a19b7702db0fd1500c96a3289f0934f700b9b17851c4d096cc1913fc8e5"></a>

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

<a id="canonical-4c99127f6b9143daf356d164a0b0ded6b739afc3901724d58d3378d9a7a4468d"></a>

## Direct properties — simple_service.enabled.persistent_volume.mount / 9f7db65e6cda / 3

<a id="canonical-c19e950163665cc14162cea5079353652436bbe4b42a5270bd8985f0f8760e1a"></a>

<a id="canonical-854e2125f91fb728c5e3958f32a2c151782de45b14353310379d7bf053cad48f"></a>

## mode property — simple_service.enabled.persistent_volume.mount / 9f7db65e6cda / 4

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

<a id="canonical-b7d74e3d65cc5dcfaac22e76d880878d35f7dfb6d8dc196600dc046f37a5e902"></a>

<a id="canonical-9edb2100c587ca5191f5980e6e3ddd6f23e2a81d7714fc6109639d2718f51e81"></a>

## mount_path property — simple_service.enabled.persistent_volume.mount / 9f7db65e6cda / 5

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

<a id="canonical-921acf5984cda43a3c3b3418cdf0956296b90bb8c292e658aaff5e7aca5eb587"></a>

<a id="canonical-5f13133863c6135a5e469db218d68f547a64825850463a74dbba3e7ae3cb44b7"></a>

## sub_path property — simple_service.enabled.persistent_volume.mount / 9f7db65e6cda / 6

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

<a id="canonical-0952cafa80c078d6ea6a412a68f90eabd9b983c7ca5aabca2f5894817e509f6d"></a>

## Next pages — simple_service.enabled.persistent_volume.mount / 9f7db65e6cda / 7

- [simple_service.enabled.persistent_volume](data-sources--workload--reference--group-016.md#canonical-9b579a402ac565efcf999975aa5d1fc5d5abd09449f94bcb5713cdbeae496e7d)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-ea59ad8c1077d994f543ce2ba96b2f8ea9ffbeda24c4e57701bdc6d3d502eeaf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6fc9af2c534b20df39f2f89459a567cb0d1383a107a93e4953bfdccbddc95f58"></a>

## simple_service.enabled.persistent_volume.storage — simple_service.enabled.persistent_volume.storage / 2aa22751f18f / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- [simple_service.enabled](data-sources--workload--reference--group-016.md#canonical-c19f09398561cccba5070e877e29bf882e759cf21d0165215f0da663e688010c)
- [simple_service.enabled.persistent_volume](data-sources--workload--reference--group-016.md#canonical-9b579a402ac565efcf999975aa5d1fc5d5abd09449f94bcb5713cdbeae496e7d)
- simple_service.enabled.persistent_volume.storage

<a id="canonical-e07b08735c8f275212b785b4808726052b6ed4efcb7d8c8a3ed7c92dc7556049"></a>

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

<a id="canonical-51c4c1c4309191ebd912613c0faac9673975d74d4a06af51276d140f7405de24"></a>

## Direct properties — simple_service.enabled.persistent_volume.storage / 2aa22751f18f / 3

<a id="canonical-2ccecd4af20e863edfcdc7c8ebd27471edb8d138ef01a6a379d4ee698472ce90"></a>

<a id="canonical-51131ee1f67145d25ba864ff61fdc6c15a8ee69f49939fd484202f31750c206c"></a>

## access_mode property — simple_service.enabled.persistent_volume.storage / 2aa22751f18f / 4

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

<a id="canonical-7ea90c8ab8cdac1c34201edff3fa692ad0a7bc1141f055a802579676649d5a54"></a>

<a id="canonical-4dbee7854ec2003cfab687072f44baa4a5a1ddd5b2af1b84f59f7bf76fe83335"></a>

## class_name property — simple_service.enabled.persistent_volume.storage / 2aa22751f18f / 5

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

- [default](data-sources--workload--reference--group-016.md#canonical-8402a18e02dca705dbb72c262965089eaf3a767c5c3f92b4e020c8c2b84fdb29): complete subsection reference.

<a id="canonical-3b9dc92651d8618760f5fb0983605d357f8b9859c928a29d6ff7109e19748c6b"></a>

<a id="canonical-20fb2ceda646c7fd405a28966a1cd6939bd8f39cb8b2f7958fa1b2f62de04f95"></a>

## storage_size property — simple_service.enabled.persistent_volume.storage / 2aa22751f18f / 6

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

<a id="canonical-950492616c4bbc5423fc2131310837109fe60037686f2a8153354e5aca5b733f"></a>

## Next pages — simple_service.enabled.persistent_volume.storage / 2aa22751f18f / 7

- [simple_service.enabled.persistent_volume.storage.default](data-sources--workload--reference--group-016.md#canonical-8402a18e02dca705dbb72c262965089eaf3a767c5c3f92b4e020c8c2b84fdb29)
- [simple_service.enabled.persistent_volume](data-sources--workload--reference--group-016.md#canonical-9b579a402ac565efcf999975aa5d1fc5d5abd09449f94bcb5713cdbeae496e7d)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-8402a18e02dca705dbb72c262965089eaf3a767c5c3f92b4e020c8c2b84fdb29"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-245f5fc4ea48ad94cedab39858b91932875747245c781107cb23a4c9d6f55125"></a>

## simple_service.enabled.persistent_volume.storage.default — simple_service.enabled.persistent_volume.storage.default / 16cbf7bc2862 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- [simple_service.enabled](data-sources--workload--reference--group-016.md#canonical-c19f09398561cccba5070e877e29bf882e759cf21d0165215f0da663e688010c)
- [simple_service.enabled.persistent_volume](data-sources--workload--reference--group-016.md#canonical-9b579a402ac565efcf999975aa5d1fc5d5abd09449f94bcb5713cdbeae496e7d)
- [simple_service.enabled.persistent_volume.storage](data-sources--workload--reference--group-016.md#canonical-ea59ad8c1077d994f543ce2ba96b2f8ea9ffbeda24c4e57701bdc6d3d502eeaf)
- simple_service.enabled.persistent_volume.storage.default

<a id="canonical-21ac2126d44a89e32ab38d1792192bac5c7fa8485c7e6b30348be856778452bb"></a>

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

<a id="canonical-d3dd657189f987f62d781d0d758bd1ed8c29f13563b022876772f24edbf34cba"></a>

## Direct properties — simple_service.enabled.persistent_volume.storage.default / 16cbf7bc2862 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4885fb317d27f797e208d8568d1d5baecd67c7805a58794d2e01f4ae10921d12"></a>

## Next pages — simple_service.enabled.persistent_volume.storage.default / 16cbf7bc2862 / 4

- [simple_service.enabled.persistent_volume.storage](data-sources--workload--reference--group-016.md#canonical-ea59ad8c1077d994f543ce2ba96b2f8ea9ffbeda24c4e57701bdc6d3d502eeaf)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-7e2c8a2c2dd86a775ee7ba82abccf291a8cac9413bf85ab6b62762694e25956c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-209cdaadb1ee7a43f7c6de2f7fd910f739816c9777ba5c9de6e2941ae3b6795d"></a>

## simple_service.simple_advertise — simple_service.simple_advertise / 58020d9e7a38 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- simple_service.simple_advertise

<a id="canonical-bab95eb3a91fa55fa2fbc2c3b9717b22156187af63f849c1929f8182df1a75d9"></a>

Type: `"single"`. Computed.

Configuration parameter for simple advertise.

Upstream description:

Advertise OPTIONS for Simple Service.

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

<a id="canonical-2f3c96c6c338819836db056349809adb996e7e77d8d03bf7f51c99972c8770a5"></a>

## Direct properties — simple_service.simple_advertise / 58020d9e7a38 / 3

<a id="canonical-6aca7fa633b4c28298eaefd5780bd98e497a08c2a09a0e78803333eee1da820f"></a>

<a id="canonical-cca110b2312adf9745af04577be05f4020063d70abdf879b03afbe4cb3e9f6c5"></a>

## domains property — simple_service.simple_advertise / 58020d9e7a38 / 4

Type: `["list", "string"]`. Computed.

List of Domains (host/authority header) that will be matched to Load Balancer. Wildcard hosts are
supported in the suffix or prefix form Supported Domains and search order: 1. Exact Domain names:
www&#46;example.com. 2.

Upstream description:

A list of Domains (host/authority header) that will be matched to Load Balancer. Wildcard hosts are
supported in the suffix or prefix form

Supported Domains and search order: &#8203;1. Exact Domain names: www&#46;example.com. &#8203;2.
Domains starting with a Wildcard: \*.example.com.

Not supported Domains: &#8203;- Just a Wildcard: \* &#8203;- A Wildcard and TLD with no root Domain:
\*.com. &#8203;- A Wildcard not matching a whole DNS label. E.g. \*.example.com and
\*.bar.example.com are valid Wildcards however \*bar.example.com, \*-bar.example.com, and
bar\*.example.com are all invalid.

Additional notes: A Wildcard will not match empty string. E.g. \*.example.com will match
bar.example.com and baz-bar.example.com but not .example.com. The longest Wildcards match first.
Only a single virtual host in the entire route configuration can match on \*. Also a Domain must be
unique across all virtual hosts within an advertise policy.

Domains are also used for SNI matching if the Load Balancer type is HTTPS. Domains also indicate the
list of names for which DNS resolution will be automatically resolved to IP addresses by the system.

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

<a id="canonical-c68ef5033acb22ede304ef41e42eb2df609c251371a8bf789393a4202e3f9e79"></a>

<a id="canonical-902df89b388b7056a04c2195d90223d726f7c9a044d752d5c308d83427946f0b"></a>

## service_port property — simple_service.simple_advertise / 58020d9e7a38 / 5

Type: `"number"`. Computed.

Service port to advertise on Internet via HTTP loadbalancer using port 80.

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
    "minimum": 1024
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1024",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1024",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-b24c5be0687052737ae5515ebd831b1decb0b555c3e0d820a3af470ac00f6fd0"></a>

## Next pages — simple_service.simple_advertise / 58020d9e7a38 / 6

- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a66eb3fa2a746cbb6a21efc5fda405160cfd7d4a1f17fd62da4f43e44f414d8e"></a>

## stateful_service — stateful_service / 4219782e2e16 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- stateful_service

<a id="canonical-de4b99b64c1600a136a6d57fc0cf1002950e68fac7bdd38d216cec7ba16c2bd1"></a>

Type: `"single"`. Computed.

StatefulService maintains per replica state and each replica has its own persistent storage. Each
replica has a unique network identity and stable storage. Stateful service are used for distributed
stateful applications like cassandra, mongodb, redis, etc.

Upstream description:

StatefulService maintains per replica state and each replica has its own persistent storage. Each
replica has a unique network identity and stable storage. Stateful service are used for distributed
stateful applications like cassandra, mongodb, redis, etc.

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

<a id="canonical-02779f4fe0e2f1959d75c19da7961b945695bc07b3b70180490672ccb3f4a801"></a>

## Direct properties — stateful_service / 4219782e2e16 / 3

- [advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068): complete subsection reference.

- [configuration](data-sources--workload--reference--group-027.md#canonical-8e969e2aef3b664039569a5afd8572b60802bcec083e0e0aaa2c24b8a09aad10): complete subsection reference.

- [containers](data-sources--workload--reference--group-027.md#canonical-d1d5c3e6c49511babec1adc9b055b806eaaa1c37ba1053ac943617ded3771c83): complete subsection reference.

- [deploy_options](data-sources--workload--reference--group-027.md#canonical-47064dc0c11a532164eb7ca9de9d90380655c494fe75b1556c94b6635d1e271f): complete subsection reference.

<a id="canonical-258b26f502cb153b99ee9d0e5276962d0929db880f37098dece969947aba63ae"></a>

<a id="canonical-ed28e4a53f537a24abff7a5eb72583fc8befad8434c5c00e5bc862ea40fd3d58"></a>

## num_replicas property — stateful_service / 4219782e2e16 / 4

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

- [persistent_volumes](data-sources--workload--reference--group-027.md#canonical-5e482719221f7b71db6737ffd54c9648576e3dfafe22e13897fb0dde9dea91b3): complete subsection reference.

- [scale_to_zero](data-sources--workload--reference--group-028.md#canonical-9cd20cbdf15c10a52dc1c08e388065f28422844a29edcef0046dc6be6e2cab36): complete subsection reference.

- [volumes](data-sources--workload--reference--group-028.md#canonical-85f200eeee91d3e87d86544517585e13e6a11ea910756373fdfea814e6e61e66): complete subsection reference.

<a id="canonical-3a9c100bdba6922be3e2171026cc50852b53ed18365c4e4538509bcb198be461"></a>

## Next pages — stateful_service / 4219782e2e16 / 5

- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.configuration](data-sources--workload--reference--group-027.md#canonical-8e969e2aef3b664039569a5afd8572b60802bcec083e0e0aaa2c24b8a09aad10)
- [stateful_service.containers](data-sources--workload--reference--group-027.md#canonical-d1d5c3e6c49511babec1adc9b055b806eaaa1c37ba1053ac943617ded3771c83)
- [stateful_service.deploy_options](data-sources--workload--reference--group-027.md#canonical-47064dc0c11a532164eb7ca9de9d90380655c494fe75b1556c94b6635d1e271f)
- [stateful_service.persistent_volumes](data-sources--workload--reference--group-027.md#canonical-5e482719221f7b71db6737ffd54c9648576e3dfafe22e13897fb0dde9dea91b3)
- [stateful_service.scale_to_zero](data-sources--workload--reference--group-028.md#canonical-9cd20cbdf15c10a52dc1c08e388065f28422844a29edcef0046dc6be6e2cab36)
- [stateful_service.volumes](data-sources--workload--reference--group-028.md#canonical-85f200eeee91d3e87d86544517585e13e6a11ea910756373fdfea814e6e61e66)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23b4aceed99da3458432b6d4a2e4789718779c41820d4527d9901d4a2cf00f48"></a>

## stateful_service.advertise_options — stateful_service.advertise_options / 78039d379b9b / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- stateful_service.advertise_options

<a id="canonical-15a71fc67c95e99a92f991b16a6a77fa8159003f60a2bcfaf2997072b2ae9c86"></a>

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

<a id="canonical-1dcfcf929aa7f954368263cf1859f5c4209691d5e82e1bad6ad87e201cf6fffa"></a>

## Direct properties — stateful_service.advertise_options / 78039d379b9b / 3

- [advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b): complete subsection reference.

- [advertise_in_cluster](data-sources--workload--reference--group-020.md#canonical-801a3f0dc79f80f95c8e3246319b812187d867b7a15c5e349b114618cc80d69a): complete subsection reference.

- [advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28): complete subsection reference.

- [do_not_advertise](data-sources--workload--reference--group-027.md#canonical-9cc6a12f6b79dcea3a6499799f15f588f6af3da6e72ca6d34e268262837b7644): complete subsection reference.

<a id="canonical-96df95ac111204c8db56a31f1e99314ced9d17fc82fc958b264a0dbb7fdbe7bd"></a>

## Next pages — stateful_service.advertise_options / 78039d379b9b / 4

- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_in_cluster](data-sources--workload--reference--group-020.md#canonical-801a3f0dc79f80f95c8e3246319b812187d867b7a15c5e349b114618cc80d69a)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.do_not_advertise](data-sources--workload--reference--group-027.md#canonical-9cc6a12f6b79dcea3a6499799f15f588f6af3da6e72ca6d34e268262837b7644)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8aca3ada230fadd2aacfbd7ed40523a615805d1d6fa8bc65429526e9a6fc7c11"></a>

## stateful_service.advertise_options.advertise_custom — stateful_service.advertise_options.advertise_custom / d115c7a3a31c / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- stateful_service.advertise_options.advertise_custom

<a id="canonical-a2cf516c7b211cdcaf11c6ec273dd6c057bca7eaf320931fd95942261ad9cd59"></a>

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

<a id="canonical-528d5ef764371b93cbf32ccb3a5e38e8a72305d99781141551ad4bb17f660c60"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom / d115c7a3a31c / 3

- [advertise_where](data-sources--workload--reference--group-016.md#canonical-8b2881b89d6b7161b0018eb07e4064f729d67e6e32a1041032b0e7bc09d2631b): complete subsection reference.

- [ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929): complete subsection reference.

<a id="canonical-a7448ed6b5e9f636b71e43cda2ceaf65fca14736ea8100b6e9d28ab7e6c49873"></a>

## Next pages — stateful_service.advertise_options.advertise_custom / d115c7a3a31c / 4

- [stateful_service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-016.md#canonical-8b2881b89d6b7161b0018eb07e4064f729d67e6e32a1041032b0e7bc09d2631b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-8b2881b89d6b7161b0018eb07e4064f729d67e6e32a1041032b0e7bc09d2631b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-546a5c0aad021f39bc2d7765cfbf737f10c00eb4e7c2049370420c5d22e9a16f"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where — stateful_service.advertise_options.advertise_custom.advertise_where / b386543425d0 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- stateful_service.advertise_options.advertise_custom.advertise_where

<a id="canonical-a74e316fb9b63b8283a4c272507fc218e5c3d497f3374232b84d675e1a06413f"></a>

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

<a id="canonical-4c8eef8cf149fc8c31a7418610bfb01e3380aff9be04aab66fbe00443a36759a"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.advertise_where / b386543425d0 / 3

- [site](data-sources--workload--reference--group-016.md#canonical-f88a170e23a3fad88ce3c20929c3ef094c638500aa6afa7f1e046b641fec0c87): complete subsection reference.

- [virtual_site](data-sources--workload--reference--group-016.md#canonical-8e28a9c5bcdf92777ecf5f5af305e6166d9df51941c591590a2eeb5362ba7d3d): complete subsection reference.

- [vk8s_service](data-sources--workload--reference--group-017.md#canonical-c58489be85f9fe13706372bcc4b0121fba090f246f0ca6c549e0ce1d95574573): complete subsection reference.

<a id="canonical-40742da24c1e321fcb1147557297295804648fde0f72ec9d87ba1fb99bc4afb2"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.advertise_where / b386543425d0 / 4

- [stateful_service.advertise_options.advertise_custom.advertise_where.site](data-sources--workload--reference--group-016.md#canonical-f88a170e23a3fad88ce3c20929c3ef094c638500aa6afa7f1e046b641fec0c87)
- [stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site](data-sources--workload--reference--group-016.md#canonical-8e28a9c5bcdf92777ecf5f5af305e6166d9df51941c591590a2eeb5362ba7d3d)
- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service](data-sources--workload--reference--group-017.md#canonical-c58489be85f9fe13706372bcc4b0121fba090f246f0ca6c549e0ce1d95574573)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-f88a170e23a3fad88ce3c20929c3ef094c638500aa6afa7f1e046b641fec0c87"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e61400c0f4f6fd65cfe1e72703bb9442e3b0d452e9c891dacb5f68f21d166506"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where.site — stateful_service.advertise_options.advertise_custom.advertise_where.site / bc385c24a96c / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-016.md#canonical-8b2881b89d6b7161b0018eb07e4064f729d67e6e32a1041032b0e7bc09d2631b)
- stateful_service.advertise_options.advertise_custom.advertise_where.site

<a id="canonical-b962421404495b05fcd496bee5aa972c1b525ac945f3099b9137b2ff53227024"></a>

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

<a id="canonical-1dc42fe6da23f36fabc0e81ce187731579f8250d25f40286b4ed421a22e77c77"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.advertise_where.site / bc385c24a96c / 3

<a id="canonical-f4d479883060a395e86a63897b9b5048f8666f1d5ea357088ad9ec76c8047b58"></a>

<a id="canonical-57b8561656b75c14c6d11683f9258f7b11be7105d915810cf35a8e7ca3cc492e"></a>

## ip property — stateful_service.advertise_options.advertise_custom.advertise_where.site / bc385c24a96c / 4

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

<a id="canonical-d849d61008ff0c67e88798f30a575d3870d22abac94d98ae7178f5ba2ff6bec5"></a>

<a id="canonical-d97abf3a8a51c41219760673c77e6fecff521c2994c6b981c4c33dbc0276447a"></a>

## network property — stateful_service.advertise_options.advertise_custom.advertise_where.site / bc385c24a96c / 5

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

- [site](data-sources--workload--reference--group-016.md#canonical-35ed9e19591b14166f69e4eda2bb987c70373416fe308bf2582f645ecaa4d705): complete subsection reference.

<a id="canonical-229d4be2417fca68b275649b1cc0428365fd4889029fde1d65e36a98bf3a1ab1"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.advertise_where.site / bc385c24a96c / 6

- [stateful_service.advertise_options.advertise_custom.advertise_where.site.site](data-sources--workload--reference--group-016.md#canonical-35ed9e19591b14166f69e4eda2bb987c70373416fe308bf2582f645ecaa4d705)
- [stateful_service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-016.md#canonical-8b2881b89d6b7161b0018eb07e4064f729d67e6e32a1041032b0e7bc09d2631b)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-35ed9e19591b14166f69e4eda2bb987c70373416fe308bf2582f645ecaa4d705"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b67f59fcf476ecbe6405895472d702bb37f4c7621dd9e86fac9c78b2cb9bdcdd"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where.site.site — stateful_service.advertise_options.advertise_custom.advertise_where.site.site / 49a8229a2e66 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-016.md#canonical-8b2881b89d6b7161b0018eb07e4064f729d67e6e32a1041032b0e7bc09d2631b)
- [stateful_service.advertise_options.advertise_custom.advertise_where.site](data-sources--workload--reference--group-016.md#canonical-f88a170e23a3fad88ce3c20929c3ef094c638500aa6afa7f1e046b641fec0c87)
- stateful_service.advertise_options.advertise_custom.advertise_where.site.site

<a id="canonical-942dc81e668392baac7da8344f31dccdcbb4644caff72950d0115cf292204fd3"></a>

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

<a id="canonical-e0b29c2f28153d0906d3f6956857d540e1538e7e98babc0dae34b1f322e12b20"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.advertise_where.site.site / 49a8229a2e66 / 3

<a id="canonical-757304372943e1efd3f7d75fe645dcc0d544b4560d9c672223423698a69da6f2"></a>

<a id="canonical-4a310c396a0a8068fc3b0a65b450e64846078902dcffb622954bd539bc3c5d47"></a>

## name property — stateful_service.advertise_options.advertise_custom.advertise_where.site.site / 49a8229a2e66 / 4

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

<a id="canonical-1476c63169e54998a3166c14cfee97ae48a93d8c6220d6535a93fd68ac809bd0"></a>

<a id="canonical-b6846629036c9e4a5c3e4960bbbd3afe9e31332d5c0a055c857537d93a8c7737"></a>

## namespace property — stateful_service.advertise_options.advertise_custom.advertise_where.site.site / 49a8229a2e66 / 5

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

<a id="canonical-bba676109ab1d4261e03d84bd8f66ee38fb18595393106b1c1224fc449c56082"></a>

<a id="canonical-2c27fd0a491dcce9e53349ed503feea33d5f93001b65fb584a43d9ff25570108"></a>

## tenant property — stateful_service.advertise_options.advertise_custom.advertise_where.site.site / 49a8229a2e66 / 6

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

<a id="canonical-0935ca89458887a55ee9a8da29c9181b36d83c4be90f334cc4699324857bb29d"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.advertise_where.site.site / 49a8229a2e66 / 7

- [stateful_service.advertise_options.advertise_custom.advertise_where.site](data-sources--workload--reference--group-016.md#canonical-f88a170e23a3fad88ce3c20929c3ef094c638500aa6afa7f1e046b641fec0c87)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-8e28a9c5bcdf92777ecf5f5af305e6166d9df51941c591590a2eeb5362ba7d3d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
