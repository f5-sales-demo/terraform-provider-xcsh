---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-bc2e50a94cdd63daf8c71b02e5e38fe84076a742d9987e6c5f8dd0c612f33f6a"></a>

## job.volumes.persistent_volume — job.volumes.persistent_volume / 5ddd58d87a88 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [job](resources--workload--reference--group-004.md#canonical-917eb75a2dcdc1639c8e0bf1be98f4f7e4c0d041764379f04763666c3858ea35)
- [job.volumes](resources--workload--reference--group-004.md#canonical-cc94e922e06d07fabcfb5821e9946f08981c925771ae93920c97f3f455668e1d)
- job.volumes.persistent_volume

<a id="canonical-d51618e3f0d1d58f2e1bee9d358ee9726227a06fb2709b4109e2e820b45d8c13"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
persistent_volume {
  # Configure direct properties listed below.
}
```

<a id="canonical-58f76c70f1454c2f28baf3ec4ec8fa27196f319594db3a756d85553958a45d51"></a>

## Direct properties — job.volumes.persistent_volume / 5ddd58d87a88 / 3

- [mount](resources--workload--reference--group-005.md#canonical-8bf4a18298b693c4c5ba6b2ec11cbe519248dd6d2d291a07e6ddf2892d586e0d): complete subsection reference.

- [storage](resources--workload--reference--group-005.md#canonical-0d419a338b5aeda5b62a366c8ad91cccb49266ee2c77f853e8140a506bff79e3): complete subsection reference.

<a id="canonical-193fff87b2ec366c70ea10efed1bdc0958e39ef99c16596007a65bd9fcd31588"></a>

## Next pages — job.volumes.persistent_volume / 5ddd58d87a88 / 4

- [job.volumes.persistent_volume.mount](resources--workload--reference--group-005.md#canonical-8bf4a18298b693c4c5ba6b2ec11cbe519248dd6d2d291a07e6ddf2892d586e0d)
- [job.volumes.persistent_volume.storage](resources--workload--reference--group-005.md#canonical-0d419a338b5aeda5b62a366c8ad91cccb49266ee2c77f853e8140a506bff79e3)
- [job.volumes](resources--workload--reference--group-004.md#canonical-cc94e922e06d07fabcfb5821e9946f08981c925771ae93920c97f3f455668e1d)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-8bf4a18298b693c4c5ba6b2ec11cbe519248dd6d2d291a07e6ddf2892d586e0d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8bdaf9e3ee89fda6092c69416d99fd17a43fb992a9c4107dfa38dffc69344ac0"></a>

## job.volumes.persistent_volume.mount — job.volumes.persistent_volume.mount / 7a42f6401e8c / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [job](resources--workload--reference--group-004.md#canonical-917eb75a2dcdc1639c8e0bf1be98f4f7e4c0d041764379f04763666c3858ea35)
- [job.volumes](resources--workload--reference--group-004.md#canonical-cc94e922e06d07fabcfb5821e9946f08981c925771ae93920c97f3f455668e1d)
- [job.volumes.persistent_volume](resources--workload--reference--group-004.md#canonical-e38069c1d5c3c16a9ea3b6a85b1d7b22744a2fed8ec6756c367678b4eba2e0b5)
- job.volumes.persistent_volume.mount

<a id="canonical-e666af0cbc13d331a3469ae147311579a22aeb37de12d1be8250ec875a45b682"></a>

Type: `"object"`. single nested block, Optional.

Volume mount describes how volume is mounted inside a workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("mount_path")}
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
mount {
  # Configure direct properties listed below.
}
```

<a id="canonical-e77bb827e38427122c2d3cebf8219b38ce85716e2afdd9f31292d1b11e8d7edb"></a>

## Direct properties — job.volumes.persistent_volume.mount / 7a42f6401e8c / 3

<a id="canonical-23e7d810e564b9f1eea9f922e68fa73f5bbbb448067b0d916b2e70b0636a5835"></a>

<a id="canonical-64b951eecb3a1a18b6e5717b7b0d2aa7dc4bb5891bc752b4f6448310719ca05d"></a>

## mode property — job.volumes.persistent_volume.mount / 7a42f6401e8c / 4

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"),
}
```

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

<a id="canonical-55a2977618a353b75c79be2dc1cf8fa50986ba1b5bfa26e9815523f6bbb1d229"></a>

<a id="canonical-3dcb8ca4bf48158036e4f7bdd5005713d8a51f1e90853fcadfab80008f53dbcc"></a>

## mount_path property — job.volumes.persistent_volume.mount / 7a42f6401e8c / 5

Type: `"string"`. Optional.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

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

<a id="canonical-17c0b08acdc3fe45ea7c55d5c3b3773b79828f115e4ffedb90776ee1024108d1"></a>

<a id="canonical-ced81f458658e1783e099e126ea7cd94b6695a32d902be9412c3cfb5fe5d834f"></a>

## sub_path property — job.volumes.persistent_volume.mount / 7a42f6401e8c / 6

Type: `"string"`. Optional.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Upstream description:

Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's
root).

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

<a id="canonical-1b003a55f16c71a6b0363d52c8364482f798eee8375ab8abcdc1d249e845e4ed"></a>

## Next pages — job.volumes.persistent_volume.mount / 7a42f6401e8c / 7

- [job.volumes.persistent_volume](resources--workload--reference--group-004.md#canonical-e38069c1d5c3c16a9ea3b6a85b1d7b22744a2fed8ec6756c367678b4eba2e0b5)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-0d419a338b5aeda5b62a366c8ad91cccb49266ee2c77f853e8140a506bff79e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8a557c42900ae76b2301bb823dfe5863c0657f44aae1d66e0cb11ec3b2fa3abb"></a>

## job.volumes.persistent_volume.storage — job.volumes.persistent_volume.storage / 7f7fe9fee649 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [job](resources--workload--reference--group-004.md#canonical-917eb75a2dcdc1639c8e0bf1be98f4f7e4c0d041764379f04763666c3858ea35)
- [job.volumes](resources--workload--reference--group-004.md#canonical-cc94e922e06d07fabcfb5821e9946f08981c925771ae93920c97f3f455668e1d)
- [job.volumes.persistent_volume](resources--workload--reference--group-004.md#canonical-e38069c1d5c3c16a9ea3b6a85b1d7b22744a2fed8ec6756c367678b4eba2e0b5)
- job.volumes.persistent_volume.storage

<a id="canonical-2fcf704ea679ee60f7bed1bbee0aa51d8074f656410ae7e832c55d06412f6680"></a>

Type: `"object"`. single nested block, Optional.

Persistent storage configuration is used to configure Persistent Volume Claim (PVC).

Upstream description:

Persistent storage configuration is used to configure Persistent Volume Claim (PVC)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("storage_size"),
  validators.ConflictingObjectAttributes("class_name",
    "default")}
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
  "x-ves-oneof-field-class_name_choice": "[\"class_name\",\"default\"]"
}
```

Terraform syntax:

```terraform
storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-26e94bbec2666ad31530170c64e1fdb1ab678bdfec544e763a8c688a0f909cde"></a>

## Direct properties — job.volumes.persistent_volume.storage / 7f7fe9fee649 / 3

<a id="canonical-3414e07cf611f324b1abfc5a377b792f6502b8ed6d154d8440fcc3ce4ba779ad"></a>

<a id="canonical-30bedcaca615dd427b4291a0271ce0e99d8ff0635359f31aaa819aef3062751e"></a>

## access_mode property — job.volumes.persistent_volume.storage / 7f7fe9fee649 / 4

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ACCESS_MODE_READ_WRITE_ONCE",
    "ACCESS_MODE_READ_WRITE_MANY",
    "ACCESS_MODE_READ_ONLY_MANY"),
}
```

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

<a id="canonical-44aa8e5ba046b8db2b824e48bece469c088fbe0530f7184d92219afba6391165"></a>

<a id="canonical-9302bb878ce66d00e7b90adbc21148e167a905d1ffd3b34adbdfe507b257621c"></a>

## class_name property — job.volumes.persistent_volume.storage / 7f7fe9fee649 / 5

Type: `"string"`. Optional.

Exclusive with \[default\] Use the specified class name.

Upstream description:

Exclusive with \[default\] Use the specified class name.

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

- [default](resources--workload--reference--group-005.md#canonical-e6191f0868c284aeacd9463fa2f613694f5110706b369e10e39dac35f80fae96): complete subsection reference.

<a id="canonical-b50f9decba4477433995912566790a13a2306d82f7dc21db6ee7d98c1d105bf7"></a>

<a id="canonical-e9f61b93eed012f924f3490bf7bbe668a9486625634d8ff6a4182b6e9ba6e6e7"></a>

## storage_size property — job.volumes.persistent_volume.storage / 7f7fe9fee649 / 6

Type: `"number"`. Optional.

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

<a id="canonical-3c6a06592ebdf3c6dd4e73574b96a4a2e3c44cdf53372d35b3508733a109702d"></a>

## Next pages — job.volumes.persistent_volume.storage / 7f7fe9fee649 / 7

- [job.volumes.persistent_volume.storage.default](resources--workload--reference--group-005.md#canonical-e6191f0868c284aeacd9463fa2f613694f5110706b369e10e39dac35f80fae96)
- [job.volumes.persistent_volume](resources--workload--reference--group-004.md#canonical-e38069c1d5c3c16a9ea3b6a85b1d7b22744a2fed8ec6756c367678b4eba2e0b5)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-e6191f0868c284aeacd9463fa2f613694f5110706b369e10e39dac35f80fae96"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59b346bf8a6665b82fee924059065c0db8158a3e4505b6ade7a12aedc80c34cf"></a>

## job.volumes.persistent_volume.storage.default — job.volumes.persistent_volume.storage.default / b7e71f417139 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [job](resources--workload--reference--group-004.md#canonical-917eb75a2dcdc1639c8e0bf1be98f4f7e4c0d041764379f04763666c3858ea35)
- [job.volumes](resources--workload--reference--group-004.md#canonical-cc94e922e06d07fabcfb5821e9946f08981c925771ae93920c97f3f455668e1d)
- [job.volumes.persistent_volume](resources--workload--reference--group-004.md#canonical-e38069c1d5c3c16a9ea3b6a85b1d7b22744a2fed8ec6756c367678b4eba2e0b5)
- [job.volumes.persistent_volume.storage](resources--workload--reference--group-005.md#canonical-0d419a338b5aeda5b62a366c8ad91cccb49266ee2c77f853e8140a506bff79e3)
- job.volumes.persistent_volume.storage.default

<a id="canonical-b3747c63e15f7d5539ced2d4fb487843868c7d6a040a38abcfb635530c316762"></a>

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
default = {}
```

<a id="canonical-ab0e32f86eff325e93b37490898da0c9ec9f2fdbee863a53cc023d54478655b8"></a>

## Direct properties — job.volumes.persistent_volume.storage.default / b7e71f417139 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-446ceeb01dd095facf53216a0d3e14b08052ac64a400798bf29589943f20dd18"></a>

## Next pages — job.volumes.persistent_volume.storage.default / b7e71f417139 / 4

- [job.volumes.persistent_volume.storage](resources--workload--reference--group-005.md#canonical-0d419a338b5aeda5b62a366c8ad91cccb49266ee2c77f853e8140a506bff79e3)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc798efcd6650fd3541cce838e77c43d4049656df71a9669513891efdae052b9"></a>

## service — service / 0d4925c79124 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- service

<a id="canonical-7ed254f36602be22a99eb0ff599d775031ded49db4c47a9006a9d0ce05a2fc0e"></a>

Type: `"object"`. single nested block, Optional.

Service does not maintain per replica state, however it can be configured to use persistent storage
that is shared amongst all the replicas. Replicas of a service are fungible and do not have a stable
network identity or storage. Common examples of services are web servers, application servers..

Upstream description:

Service does not maintain per replica state, however it can be configured to use persistent storage
that is shared amongst all the replicas. Replicas of a service are fungible and do not have a stable
network identity or storage. Common examples of services are web servers, application servers,
traditional SQL databases, etc.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("containers"),
  validators.ConflictingObjectAttributes("num_replicas",
    "scale_to_zero")}
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
  "x-ves-oneof-field-scaling_choice": "[\"num_replicas\",\"scale_to_zero\"]"
}
```

Terraform syntax:

```terraform
service {
  # Configure direct properties listed below.
}
```

<a id="canonical-926acaf5eefc6e3c877eec5ea95b746ab6e27c6c16c1826653a5961df1b11be2"></a>

## Direct properties — service / 0d4925c79124 / 3

- [advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c): complete subsection reference.

- [configuration](resources--workload--reference--group-015.md#canonical-18c1d9d98cecc2c731f787c41b62d9d968a8d90541ec19c8272ab7e594301dbf): complete subsection reference.

- [containers](resources--workload--reference--group-015.md#canonical-3f88422b353a65363a1ce1b0d566b8c6aa769a87c87c23cb23b8978323a9c01a): complete subsection reference.

- [deploy_options](resources--workload--reference--group-016.md#canonical-05165ff47537505ed9ace4fa2b50ac0d2124ed1fc4414c6d288c2e658d5e2ec3): complete subsection reference.

<a id="canonical-3f9861549c583367351241b269affbfa21555eee0c5b802266211b679b851d15"></a>

<a id="canonical-f261b87be5df5c8df656593dfe6cdcaa9ac61b7913ffe12dd858ad4c122f578f"></a>

## num_replicas property — service / 0d4925c79124 / 4

Type: `"number"`. Optional.

Exclusive with \[scale\_to\_zero\] Number of replicas of service to spawn per site.

Upstream description:

Exclusive with \[scale\_to\_zero\] Number of replicas of service to spawn per site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 5),
}
```

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

- [scale_to_zero](resources--workload--reference--group-016.md#canonical-968fb38b4a3e5a07bd544d6620a48cbb9b651535a5806eb59a7a0bb1caaa9125): complete subsection reference.

- [volumes](resources--workload--reference--group-016.md#canonical-faacd6685afddb473ef64bdf75dbfc0881f5d7702ee057ad8724fc6f0339f0c3): complete subsection reference.

<a id="canonical-b5bd813f5ebe7d6972369eef0b981c58f264d4a27e1a615a44e6e2dfea408882"></a>

## Next pages — service / 0d4925c79124 / 5

- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.configuration](resources--workload--reference--group-015.md#canonical-18c1d9d98cecc2c731f787c41b62d9d968a8d90541ec19c8272ab7e594301dbf)
- [service.containers](resources--workload--reference--group-015.md#canonical-3f88422b353a65363a1ce1b0d566b8c6aa769a87c87c23cb23b8978323a9c01a)
- [service.deploy_options](resources--workload--reference--group-016.md#canonical-05165ff47537505ed9ace4fa2b50ac0d2124ed1fc4414c6d288c2e658d5e2ec3)
- [service.scale_to_zero](resources--workload--reference--group-016.md#canonical-968fb38b4a3e5a07bd544d6620a48cbb9b651535a5806eb59a7a0bb1caaa9125)
- [service.volumes](resources--workload--reference--group-016.md#canonical-faacd6685afddb473ef64bdf75dbfc0881f5d7702ee057ad8724fc6f0339f0c3)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-901dfa7db1516c11bb742a01fc852c76bbdce7ff0e0d3fbef9a3b3dff5d10d49"></a>

## service.advertise_options — service.advertise_options / 5cc7aa7aff53 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- service.advertise_options

<a id="canonical-7c5d5ff92252bd534f27de377cded6445170308c226af5fc8d38be36628a6e4d"></a>

Type: `"object"`. single nested block, Optional.

Advertise OPTIONS are used to configure how and where to advertise the workload using load
balancers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_in_cluster"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_on_public"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_in_cluster",
    "advertise_on_public"),
  validators.ConflictingObjectAttributes("advertise_in_cluster",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_on_public",
    "do_not_advertise")}
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
  "x-ves-oneof-field-advertise_choice": "[\"advertise_custom\",\"advertise_in_cluster\",\"advertise_on_public\",\"do_not_advertise\"]"
}
```

Terraform syntax:

```terraform
advertise_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-40d800e18391ee3b07f95beed10bd2e907550dafbd46bd4aa94ba30fac653f67"></a>

## Direct properties — service.advertise_options / 5cc7aa7aff53 / 3

- [advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508): complete subsection reference.

- [advertise_in_cluster](resources--workload--reference--group-008.md#canonical-3c2f17491c942f70c268ee3196b5f46ff346531af030c2f06723830827076967): complete subsection reference.

- [advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1): complete subsection reference.

- [do_not_advertise](resources--workload--reference--group-015.md#canonical-34c8bee64727675e6ab55d819968b7a37e89ce671f1a8c0fd53905e1cd97c2bf): complete subsection reference.

<a id="canonical-617ed77ea03a47db6fb5548e2163c803757bf024acbbd474f7d020eaac12ba94"></a>

## Next pages — service.advertise_options / 5cc7aa7aff53 / 4

- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_in_cluster](resources--workload--reference--group-008.md#canonical-3c2f17491c942f70c268ee3196b5f46ff346531af030c2f06723830827076967)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.do_not_advertise](resources--workload--reference--group-015.md#canonical-34c8bee64727675e6ab55d819968b7a37e89ce671f1a8c0fd53905e1cd97c2bf)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e99b70952e5ca106db027f2ff856c339467ea950b3a892d19c133dcad8d8af6"></a>

## service.advertise_options.advertise_custom — service.advertise_options.advertise_custom / 9693af2a8b1b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- service.advertise_options.advertise_custom

<a id="canonical-a53beed37f67e6d9ce4f9e83fd688957612e65023bbee2b1924559efd71cf783"></a>

Type: `"object"`. single nested block, Optional.

Advertise this workload via loadbalancer on specific sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("advertise_where",
    "ports")}
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
advertise_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-e6f67174d54d3aff8e301febdde047c274c28a79d7ff0217e217c189830a5d03"></a>

## Direct properties — service.advertise_options.advertise_custom / 9693af2a8b1b / 3

- [advertise_where](resources--workload--reference--group-005.md#canonical-00ca5d0c769bfa33bc8a0e613544a3347aa85ba1183a334fc2dc88818236428f): complete subsection reference.

- [ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920): complete subsection reference.

<a id="canonical-211aa509eb106407b5dc2d6635b29f83920b47328139bd0da6cac6b56b798c94"></a>

## Next pages — service.advertise_options.advertise_custom / 9693af2a8b1b / 4

- [service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-005.md#canonical-00ca5d0c769bfa33bc8a0e613544a3347aa85ba1183a334fc2dc88818236428f)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-00ca5d0c769bfa33bc8a0e613544a3347aa85ba1183a334fc2dc88818236428f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1974aa6c9e3cd65de989b484187f2f1d4cb5a78faed0d2eedf543a685e6ced71"></a>

## service.advertise_options.advertise_custom.advertise_where — service.advertise_options.advertise_custom.advertise_where / a1f36b603397 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- service.advertise_options.advertise_custom.advertise_where

<a id="canonical-0204cadd2145e10638770d38abd5218ed5af503f55ac6872255a1b3384832da8"></a>

Type: `"object"`. list nested block, Optional.

Where should this load balancer be available.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("site",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("site",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_site",
    "vk8s_service")}
```

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

Terraform syntax:

```terraform
advertise_where {
  # Configure direct properties listed below.
}
```

<a id="canonical-2909e317ae0118c019e1c468225151e17ff81670318212e4647944c8f2baef34"></a>

## Direct properties — service.advertise_options.advertise_custom.advertise_where / a1f36b603397 / 3

- [site](resources--workload--reference--group-005.md#canonical-0cfd0cc77f68f034ba74d106f8d4968d040993fd82688fa36e176db500e8c0c1): complete subsection reference.

- [virtual_site](resources--workload--reference--group-005.md#canonical-2f92e45155798b1ec632e6dc8db45e11f585717297f37ece32f4dee3d1cec29d): complete subsection reference.

- [vk8s_service](resources--workload--reference--group-005.md#canonical-e32ccd455bbcc70499af0e933d37f48c43eeeb859ac2ef52f6fb568260ee3681): complete subsection reference.

<a id="canonical-48663b6ff4b505ebeac8e3e698a434c542f7d1765237f52ee93d5dfb00384b71"></a>

## Next pages — service.advertise_options.advertise_custom.advertise_where / a1f36b603397 / 4

- [service.advertise_options.advertise_custom.advertise_where.site](resources--workload--reference--group-005.md#canonical-0cfd0cc77f68f034ba74d106f8d4968d040993fd82688fa36e176db500e8c0c1)
- [service.advertise_options.advertise_custom.advertise_where.virtual_site](resources--workload--reference--group-005.md#canonical-2f92e45155798b1ec632e6dc8db45e11f585717297f37ece32f4dee3d1cec29d)
- [service.advertise_options.advertise_custom.advertise_where.vk8s_service](resources--workload--reference--group-005.md#canonical-e32ccd455bbcc70499af0e933d37f48c43eeeb859ac2ef52f6fb568260ee3681)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-0cfd0cc77f68f034ba74d106f8d4968d040993fd82688fa36e176db500e8c0c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-52284aac863cd54806242fe456f6eb4a1e039d378b74ef122d5535e11d6b7f7e"></a>

## service.advertise_options.advertise_custom.advertise_where.site — service.advertise_options.advertise_custom.advertise_where.site / 781472edaced / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-005.md#canonical-00ca5d0c769bfa33bc8a0e613544a3347aa85ba1183a334fc2dc88818236428f)
- service.advertise_options.advertise_custom.advertise_where.site

<a id="canonical-01d9cad205ba7ac453ab9f197d48a0fc96d03df7ff9b2d194a265d6bc60b5326"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-ddf0aa6febb152598aa91288b5a6874dfcd7840935d61fa0d40e7eb3a3cc46eb"></a>

## Direct properties — service.advertise_options.advertise_custom.advertise_where.site / 781472edaced / 3

<a id="canonical-7a28acdbab2cf7d80014de673e9561802713a25f9eae149eea7955e062918604"></a>

<a id="canonical-d27e0a56b0f04ff6b3864f20e94c9ac77800c93aa86afda988d8e363c4503f64"></a>

## ip property — service.advertise_options.advertise_custom.advertise_where.site / 781472edaced / 4

Type: `"string"`. Optional.

Use given IP address as VIP on the site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

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

<a id="canonical-3c80a18564d9db0e234e55dc91abc4025b49e165b94d26ee00ac046d02ec632e"></a>

<a id="canonical-8b1f06177bd3b0db2f7f3decd710178f63f013b1dfbc15040a8c958edd33b636"></a>

## network property — service.advertise_options.advertise_custom.advertise_where.site / 781472edaced / 5

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"),
}
```

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

- [site](resources--workload--reference--group-005.md#canonical-f82b550ad33038ba5d1ae21fb8e548fe151e4e5a3add17bd9dbd5da75807eb9a): complete subsection reference.

<a id="canonical-4ec28929b6c6c1c43b463192dfc92dbe3108df7db6b114a790bf68174fd82f1c"></a>

## Next pages — service.advertise_options.advertise_custom.advertise_where.site / 781472edaced / 6

- [service.advertise_options.advertise_custom.advertise_where.site.site](resources--workload--reference--group-005.md#canonical-f82b550ad33038ba5d1ae21fb8e548fe151e4e5a3add17bd9dbd5da75807eb9a)
- [service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-005.md#canonical-00ca5d0c769bfa33bc8a0e613544a3347aa85ba1183a334fc2dc88818236428f)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-f82b550ad33038ba5d1ae21fb8e548fe151e4e5a3add17bd9dbd5da75807eb9a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-989d48be4c38922f5873004d1817adc0b9b68ff989ef12fbb4d4343ebdb4b594"></a>

## service.advertise_options.advertise_custom.advertise_where.site.site — service.advertise_options.advertise_custom.advertise_where.site.site / a287aa08dc0d / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-005.md#canonical-00ca5d0c769bfa33bc8a0e613544a3347aa85ba1183a334fc2dc88818236428f)
- [service.advertise_options.advertise_custom.advertise_where.site](resources--workload--reference--group-005.md#canonical-0cfd0cc77f68f034ba74d106f8d4968d040993fd82688fa36e176db500e8c0c1)
- service.advertise_options.advertise_custom.advertise_where.site.site

<a id="canonical-9e6757c39e21be1efe8dd5835d419ce01c3aaf4f10a8ef672b394858f9d9b2bb"></a>

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-67282107beac8c540b104e850849ce9be600dc1e3441d967c5b77e702994633e"></a>

## Direct properties — service.advertise_options.advertise_custom.advertise_where.site.site / a287aa08dc0d / 3

<a id="canonical-e0ed842947cd3a13669d313fe57eef17f5aec932124ac6def82e0042d3ad87cb"></a>

<a id="canonical-f848cb818052cc4cda824f48c35d4d4b488a2d0bb96800c6d6be694fed0cb06e"></a>

## name property — service.advertise_options.advertise_custom.advertise_where.site.site / a287aa08dc0d / 4

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

<a id="canonical-0aa33129f70ec986baaa5187b0f256e80fceaebd9d39730a114bf2d377d88c27"></a>

<a id="canonical-4c617b8e2d2c6c0879cf8e806528ae5cbad43ca3cf29223a38359e30591bbbe0"></a>

## namespace property — service.advertise_options.advertise_custom.advertise_where.site.site / a287aa08dc0d / 5

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

<a id="canonical-cbbefd409170d6eb4ed7f8d19654fb7cd1cc2827d34d8f04d3d343a603a9a1cc"></a>

<a id="canonical-7d730a1702ebb7018b94b92bb5e5ce7a43690b018c4ed01901a1b4bf28b54e67"></a>

## tenant property — service.advertise_options.advertise_custom.advertise_where.site.site / a287aa08dc0d / 6

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

<a id="canonical-02147977522859740101e97bf50050ad5878f2c8a7c43100e8360ec4386e83f7"></a>

## Next pages — service.advertise_options.advertise_custom.advertise_where.site.site / a287aa08dc0d / 7

- [service.advertise_options.advertise_custom.advertise_where.site](resources--workload--reference--group-005.md#canonical-0cfd0cc77f68f034ba74d106f8d4968d040993fd82688fa36e176db500e8c0c1)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-2f92e45155798b1ec632e6dc8db45e11f585717297f37ece32f4dee3d1cec29d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4ae84a2547d280476049921247861cfff191abf577b640aff4c131ed7603b2d"></a>

## service.advertise_options.advertise_custom.advertise_where.virtual_site — service.advertise_options.advertise_custom.advertise_where.virtual_site / 74e194a11019 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-005.md#canonical-00ca5d0c769bfa33bc8a0e613544a3347aa85ba1183a334fc2dc88818236428f)
- service.advertise_options.advertise_custom.advertise_where.virtual_site

<a id="canonical-ffdcb15550e624909db94ea1e125290bb192cc624b3ca06fba30e0ae3f5dd0f1"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-186a9955c28f28d901110f7dc8e4eb81a0e7ba5442e058e56858ad14045573a5"></a>

## Direct properties — service.advertise_options.advertise_custom.advertise_where.virtual_site / 74e194a11019 / 3

<a id="canonical-cc26769fa563f5476fed6298f18ced1f62fc38276d00207c8906f91b18fc7b03"></a>

<a id="canonical-d5d0d9b5734ba0d3f7a24ccd85df16e6bd5442dc67968715d772a68a35c196bd"></a>

## network property — service.advertise_options.advertise_custom.advertise_where.virtual_site / 74e194a11019 / 4

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"),
}
```

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

- [virtual_site](resources--workload--reference--group-005.md#canonical-d594cca7149d0c6e48bf38e9ce087cd0880dd2089bdc2e0a120e71b5160fb7b9): complete subsection reference.

<a id="canonical-bd38c7739595fb8352eee74c5e7df83b047f285c97cdc5afcee2e19f99a3d8ff"></a>

## Next pages — service.advertise_options.advertise_custom.advertise_where.virtual_site / 74e194a11019 / 5

- [service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site](resources--workload--reference--group-005.md#canonical-d594cca7149d0c6e48bf38e9ce087cd0880dd2089bdc2e0a120e71b5160fb7b9)
- [service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-005.md#canonical-00ca5d0c769bfa33bc8a0e613544a3347aa85ba1183a334fc2dc88818236428f)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-d594cca7149d0c6e48bf38e9ce087cd0880dd2089bdc2e0a120e71b5160fb7b9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-32ba0d3eb4815c445816fad899a536ed96f6f5d01adf83d0a0a4983e2d9f50d8"></a>

## service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site — service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_ / 6c1f449cf1f6 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-005.md#canonical-00ca5d0c769bfa33bc8a0e613544a3347aa85ba1183a334fc2dc88818236428f)
- [service.advertise_options.advertise_custom.advertise_where.virtual_site](resources--workload--reference--group-005.md#canonical-2f92e45155798b1ec632e6dc8db45e11f585717297f37ece32f4dee3d1cec29d)
- service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site

<a id="canonical-e1376f934f18a70dd1066c80e0f2b42cd951865cea58da2e9c4be821f143406d"></a>

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-a7f83e732e81d42625de8655bb7035f1d2230134f6d7dd43a3d8b191cada14df"></a>

## Direct properties — service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_ / 6c1f449cf1f6 / 3

<a id="canonical-16a9aff119b7c2b87c743d14f47e16fd447147f23b3c19fcdf5ad163c34c5966"></a>

<a id="canonical-7ef82c04649baacd8dfc6a785603a90a4a9ca991aea77544c36e71021b72b71d"></a>

## name property — service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_ / 6c1f449cf1f6 / 4

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

<a id="canonical-64defedcb4d418e114d80530a4717cac426b590f5116df80f41566fd5ca6ea4e"></a>

<a id="canonical-82c035ba620596fe22fda94f5385dd1e48b0939664311430930a5b282d580ecf"></a>

## namespace property — service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_ / 6c1f449cf1f6 / 5

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

<a id="canonical-92c575c7461eef0a42e6bca1d0f5d5a1a8c0e656fbe027d62a750b6c09b3e00f"></a>

<a id="canonical-4a3cb9a5b6b89a304eb2aac96f392f989392e21b411cf8ee80c60265dc521b90"></a>

## tenant property — service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_ / 6c1f449cf1f6 / 6

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

<a id="canonical-50342474bff1af2e6489aad171cd842010b9489772f8d872ba60529f5bdcbfb7"></a>

## Next pages — service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_ / 6c1f449cf1f6 / 7

- [service.advertise_options.advertise_custom.advertise_where.virtual_site](resources--workload--reference--group-005.md#canonical-2f92e45155798b1ec632e6dc8db45e11f585717297f37ece32f4dee3d1cec29d)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-e32ccd455bbcc70499af0e933d37f48c43eeeb859ac2ef52f6fb568260ee3681"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa54b6300a0120037000217e84ff8d11b0acb35363eb46149766ce3dac4c9b76"></a>

## service.advertise_options.advertise_custom.advertise_where.vk8s_service — service.advertise_options.advertise_custom.advertise_where.vk8s_service / 9e692ad5ce0b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-005.md#canonical-00ca5d0c769bfa33bc8a0e613544a3347aa85ba1183a334fc2dc88818236428f)
- service.advertise_options.advertise_custom.advertise_where.vk8s_service

<a id="canonical-2805dbba151cd502d98558c0a955dd3b3e2f41c85ba646d50e436aa6f181fdbe"></a>

Type: `"object"`. single nested block, Optional.

Defines a reference to a RE site or virtual site where a load balancer could be advertised in the
vK8s service network.

Upstream description:

This defines a reference to a RE site or virtual site where a load balancer could be advertised in
the vK8s service network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_site")}
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
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
vk8s_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-67a0f2dd3f0cda373a1e5be005f4c97c3e9fa1b6f5aca1b06a8232762709f872"></a>

## Direct properties — service.advertise_options.advertise_custom.advertise_where.vk8s_service / 9e692ad5ce0b / 3

- [site](resources--workload--reference--group-005.md#canonical-9ae4a2bee3ab5504abbafaf3e964c3cb6ff0bd7f4ace0d78bdde7894072c1aa6): complete subsection reference.

- [virtual_site](resources--workload--reference--group-005.md#canonical-5f2f94827814bd8c4bc4a9911517d4d6539886aef20bb2d80f55106313d2d501): complete subsection reference.

<a id="canonical-68c998f0ad66bfa82c157431ea6e4814859a5d18ebc2a8048565a8cf2d53c84c"></a>

## Next pages — service.advertise_options.advertise_custom.advertise_where.vk8s_service / 9e692ad5ce0b / 4

- [service.advertise_options.advertise_custom.advertise_where.vk8s_service.site](resources--workload--reference--group-005.md#canonical-9ae4a2bee3ab5504abbafaf3e964c3cb6ff0bd7f4ace0d78bdde7894072c1aa6)
- [service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site](resources--workload--reference--group-005.md#canonical-5f2f94827814bd8c4bc4a9911517d4d6539886aef20bb2d80f55106313d2d501)
- [service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-005.md#canonical-00ca5d0c769bfa33bc8a0e613544a3347aa85ba1183a334fc2dc88818236428f)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-9ae4a2bee3ab5504abbafaf3e964c3cb6ff0bd7f4ace0d78bdde7894072c1aa6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f820ae6418de7fc021798d26fb24728ad27643927e4f4ca8c2d98826bc5ce8d"></a>

## service.advertise_options.advertise_custom.advertise_where.vk8s_service.site — service.advertise_options.advertise_custom.advertise_where.vk8s_service.site / a30624b49be0 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-005.md#canonical-00ca5d0c769bfa33bc8a0e613544a3347aa85ba1183a334fc2dc88818236428f)
- [service.advertise_options.advertise_custom.advertise_where.vk8s_service](resources--workload--reference--group-005.md#canonical-e32ccd455bbcc70499af0e933d37f48c43eeeb859ac2ef52f6fb568260ee3681)
- service.advertise_options.advertise_custom.advertise_where.vk8s_service.site

<a id="canonical-1a0023230b41dddecd590c6621045965b617616546bf0af0efafc9e16e6c0e3d"></a>

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-b048e7719abdf9ebec586d4f21362c26d1e6a1d79f54f76edf1f0dd2ed47a9b2"></a>

## Direct properties — service.advertise_options.advertise_custom.advertise_where.vk8s_service.site / a30624b49be0 / 3

<a id="canonical-7453ce30efa5e35daa8482bd82ec023cfbd95f80bdbc635b26801224b71353d3"></a>

<a id="canonical-6234fb0e930b03a98320937f4dab5d91d2417a5b4ea7ffedb2071981f5b05a60"></a>

## name property — service.advertise_options.advertise_custom.advertise_where.vk8s_service.site / a30624b49be0 / 4

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

<a id="canonical-86afe8fca8d82bed8dd08751ce63767a69069712c82bb721bfe1f3d9aa71b9cf"></a>

<a id="canonical-3bcc3e77d31b695c9a0da0e16dc3a145fa1364dd3779e24f9221cbf04ddaa56e"></a>

## namespace property — service.advertise_options.advertise_custom.advertise_where.vk8s_service.site / a30624b49be0 / 5

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

<a id="canonical-716253fb313475f50b430bd15ea10b3625056dce7cb1ba051e3645b3b9cc8a5f"></a>

<a id="canonical-c9e07311e5b4727defa88cae2960ad9d52fe9462fd08ef2d93a43e58b9ba6137"></a>

## tenant property — service.advertise_options.advertise_custom.advertise_where.vk8s_service.site / a30624b49be0 / 6

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

<a id="canonical-32ac3f35fb70bfec81fefb6d175a23ecc7faa78fc7905950864cc703415541eb"></a>

## Next pages — service.advertise_options.advertise_custom.advertise_where.vk8s_service.site / a30624b49be0 / 7

- [service.advertise_options.advertise_custom.advertise_where.vk8s_service](resources--workload--reference--group-005.md#canonical-e32ccd455bbcc70499af0e933d37f48c43eeeb859ac2ef52f6fb568260ee3681)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-5f2f94827814bd8c4bc4a9911517d4d6539886aef20bb2d80f55106313d2d501"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-00567328cb36862580f7c557e700db8410450d9f40f76ec005274784e86cfe83"></a>

## service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site — service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_ / 721ee3e79eea / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-005.md#canonical-00ca5d0c769bfa33bc8a0e613544a3347aa85ba1183a334fc2dc88818236428f)
- [service.advertise_options.advertise_custom.advertise_where.vk8s_service](resources--workload--reference--group-005.md#canonical-e32ccd455bbcc70499af0e933d37f48c43eeeb859ac2ef52f6fb568260ee3681)
- service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="canonical-c62144682f06e530a9ae1ff8421a8e1cbc523de8be529a4753de1b4d3e60a646"></a>

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-d3591a5442c5de7b2b1178ca85e446094f2d0ed074ca118b15e0cd8513a3636b"></a>

## Direct properties — service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_ / 721ee3e79eea / 3

<a id="canonical-a7fe648471d1e110ecf5c7615ba9d5be7bd8230bbdd83d5123ec688562a3e40c"></a>

<a id="canonical-2e9f11e42dd06539c3825f300340bacefb8af6e653793563ddec47f633ba4cd1"></a>

## name property — service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_ / 721ee3e79eea / 4

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

<a id="canonical-d8592ccdfe4d63713fa94a51ed03ed60bc71a6e937284154cfc33fc04adcc2e0"></a>

<a id="canonical-ced45cda872be95f8de8be1fd4b70b3471c1d1202dbaf3b74b0e17811b725d5c"></a>

## namespace property — service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_ / 721ee3e79eea / 5

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

<a id="canonical-97914187274f9e8a54ead2a6dbfd86278b85c811b98abefe36b4cc88ddd8153c"></a>

<a id="canonical-44678c0ff27caabfed4129df0c67ac291232a8fee2c5f0f2994b0d60e1172e42"></a>

## tenant property — service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_ / 721ee3e79eea / 6

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

<a id="canonical-24c5c7cf9b4462f9486d5778f5b7923efedbd5bd2c2ef7a9fcbe81c3fa30c9f8"></a>

## Next pages — service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_ / 721ee3e79eea / 7

- [service.advertise_options.advertise_custom.advertise_where.vk8s_service](resources--workload--reference--group-005.md#canonical-e32ccd455bbcc70499af0e933d37f48c43eeeb859ac2ef52f6fb568260ee3681)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f97f0eab18c22b4b72e86083b12ec3b1e83cb5ec078d036740b714751749929"></a>

## service.advertise_options.advertise_custom.ports — service.advertise_options.advertise_custom.ports / 4f15d78d959b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- service.advertise_options.advertise_custom.ports

<a id="canonical-9b1506e08baa0cc194fdd794fff98200dc9ca985d75e1fd5f756f5d3585c4a6f"></a>

Type: `"object"`. list nested block, Optional.

Ports. Ports to advertise.

Upstream description:

Ports to advertise.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("http_loadbalancer",
    "tcp_loadbalancer")}
```

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

Terraform syntax:

```terraform
ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-2bea6fe34a93957f96002354f7ccc3d07b996b21698a767ebfa4044bf8431bce"></a>

## Direct properties — service.advertise_options.advertise_custom.ports / 4f15d78d959b / 3

- [http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192): complete subsection reference.

- [port](resources--workload--reference--group-008.md#canonical-85b6f4e25c7f1b69d65aa7051f4382aac22e4c6dfbe3dde9ecc33d5532af381b): complete subsection reference.

- [tcp_loadbalancer](resources--workload--reference--group-008.md#canonical-bf03b979f8a27555c05ed659a5a2b154f3304086fa331a9795acd884385be280): complete subsection reference.

<a id="canonical-91a988d0ac1e00e1233d3ec6cfaae96bd226301638b3cc1190edb13f5cc7c8a9"></a>

## Next pages — service.advertise_options.advertise_custom.ports / 4f15d78d959b / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.port](resources--workload--reference--group-008.md#canonical-85b6f4e25c7f1b69d65aa7051f4382aac22e4c6dfbe3dde9ecc33d5532af381b)
- [service.advertise_options.advertise_custom.ports.tcp_loadbalancer](resources--workload--reference--group-008.md#canonical-bf03b979f8a27555c05ed659a5a2b154f3304086fa331a9795acd884385be280)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-79eef294907d10436f97097a9820ffbc49565984406c9dcddb43d8dcdb59e7bb"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer — service.advertise_options.advertise_custom.ports.http_loadbalancer / 1dd7081407e2 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- service.advertise_options.advertise_custom.ports.http_loadbalancer

<a id="canonical-33b0f150cd772d96f47b6081f5b9b409456197f32c7392880a8373e5db99714e"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http loadbalancer.

Upstream description:

HTTP/HTTPS Load balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domains"),
  validators.ConflictingObjectAttributes("default_route",
    "specific_routes"),
  validators.ConflictingObjectAttributes("http",
    "https"),
  validators.ConflictingObjectAttributes("http",
    "https_auto_cert"),
  validators.ConflictingObjectAttributes("https",
    "https_auto_cert")}
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
  "x-ves-oneof-field-loadbalancer_type": "[\"http\",\"https\",\"https_auto_cert\"]",
  "x-ves-oneof-field-route_choice": "[\"default_route\",\"specific_routes\"]"
}
```

Terraform syntax:

```terraform
http_loadbalancer {
  # Configure direct properties listed below.
}
```

<a id="canonical-8245b961b4064b652de18505d23c5bc8db3e8d50ff322008f72131aa8d714ff5"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer / 1dd7081407e2 / 3

- [default_route](resources--workload--reference--group-005.md#canonical-58fbfdd83a4752f4560b8b0452371470c11a31f1a96d635043ff9ae07d4aea92): complete subsection reference.

<a id="canonical-67cc8f5dc9ad8edc84ec6c5c9c43c2fe6b629f291055afaa72008556b25e6154"></a>

<a id="canonical-a9545b7cfcd129fcf218cee5e50a7c51adefe77bb45bd968d0534fb69e321474"></a>

## domains property — service.advertise_options.advertise_custom.ports.http_loadbalancer / 1dd7081407e2 / 4

Type: `["list", "string"]`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
```

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

- [http](resources--workload--reference--group-005.md#canonical-9c3ebce39f55bdb8564389438b789c0d0f73f05949e6d9ae831e29f588360ef4): complete subsection reference.

- [https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45): complete subsection reference.

- [https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336): complete subsection reference.

- [specific_routes](resources--workload--reference--group-007.md#canonical-c5972512540fc706bc0ed09b7f9a33e566440963ffde68a75abbe8fe2853eeee): complete subsection reference.

<a id="canonical-be2faf81edf96aea16e31b1aa9843292f6eae38cf22135a0fcef0819943b4e32"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer / 1dd7081407e2 / 5

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](resources--workload--reference--group-005.md#canonical-58fbfdd83a4752f4560b8b0452371470c11a31f1a96d635043ff9ae07d4aea92)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.http](resources--workload--reference--group-005.md#canonical-9c3ebce39f55bdb8564389438b789c0d0f73f05949e6d9ae831e29f588360ef4)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-007.md#canonical-c5972512540fc706bc0ed09b7f9a33e566440963ffde68a75abbe8fe2853eeee)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-58fbfdd83a4752f4560b8b0452371470c11a31f1a96d635043ff9ae07d4aea92"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2fa9f7e4a6c2522272687716f2d0795593c3a7461a0f2c30a726bfb3c28ba167"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route — service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route / bca99bd181a9 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route

<a id="canonical-892e0e14d7f186a68be07c895875511ff8a305e13720302f92bad57d05295682"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for default route.

Upstream description:

Default route matching all APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto_host_rewrite",
    "disable_host_rewrite"),
  validators.ConflictingObjectAttributes("auto_host_rewrite",
    "host_rewrite"),
  validators.ConflictingObjectAttributes("disable_host_rewrite",
    "host_rewrite")}
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
  "x-ves-oneof-field-host_rewrite_params": "[\"auto_host_rewrite\",\"disable_host_rewrite\",\"host_rewrite\"]"
}
```

Terraform syntax:

```terraform
default_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-64671477ffcc33122121af96d58eae9bba111c196f1f7b178290118ee893128c"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route / bca99bd181a9 / 3

- [auto_host_rewrite](resources--workload--reference--group-005.md#canonical-bf913b0456bb6018aaf7a7bb1507ce4d81cc9d119d9c31be7172b7ebc47a5de9): complete subsection reference.

- [disable_host_rewrite](resources--workload--reference--group-005.md#canonical-51e7060aee134acbe7acc551f50aa5065901549d339c8fdc718353a1b045323e): complete subsection reference.

<a id="canonical-c8c7f012313ce1f56aded56cf63bbf376a55197818529a21a148f3b8032b8aea"></a>

<a id="canonical-e57edeef22ddea4fac148fb3ae5e2ec56f051d171b9ecd2bfbc97fee93297cf6"></a>

## host_rewrite property — service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route / bca99bd181a9 / 4

Type: `"string"`. Optional.

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Upstream description:

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

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

<a id="canonical-4eb257a9aa13e58533a55b5782c96744446130702a592e98704987b99f4b4707"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route / bca99bd181a9 / 5

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.auto_host_rewrite](resources--workload--reference--group-005.md#canonical-bf913b0456bb6018aaf7a7bb1507ce4d81cc9d119d9c31be7172b7ebc47a5de9)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.disable_host_rewrite](resources--workload--reference--group-005.md#canonical-51e7060aee134acbe7acc551f50aa5065901549d339c8fdc718353a1b045323e)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-bf913b0456bb6018aaf7a7bb1507ce4d81cc9d119d9c31be7172b7ebc47a5de9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fad42ec715a2ef50ea40261bda7a8e9db1043fe935102e8d57c39ca4a11074dc"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.auto_host_rewrite — service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route / 2811af7eab6c / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](resources--workload--reference--group-005.md#canonical-58fbfdd83a4752f4560b8b0452371470c11a31f1a96d635043ff9ae07d4aea92)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.auto_host_rewrite

<a id="canonical-41a74bd62cdfacbe44cc81fbd14d6cdbf000af39360815dd4f67304fcfe75fe6"></a>

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
auto_host_rewrite = {}
```

<a id="canonical-bc6b59f0a402e8185a94afcf65f111dd71d86662c614445d0124e6877c7ac631"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route / 2811af7eab6c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dacf299800a5d7d94d0359b2d9ed4a7c6de737bb404ba60d29055afda07f7111"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route / 2811af7eab6c / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](resources--workload--reference--group-005.md#canonical-58fbfdd83a4752f4560b8b0452371470c11a31f1a96d635043ff9ae07d4aea92)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-51e7060aee134acbe7acc551f50aa5065901549d339c8fdc718353a1b045323e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c8c61295aebb809895d3fcf370e9847dbdb825c3a005d8a78a5be6429cd1615b"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.disable_host_rewrite — service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route / aff9f4559807 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](resources--workload--reference--group-005.md#canonical-58fbfdd83a4752f4560b8b0452371470c11a31f1a96d635043ff9ae07d4aea92)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.disable_host_rewrite

<a id="canonical-482dac511c414c84db1ad68c985102deb1fc7bed54fbf84b642ebdb5830e3c7e"></a>

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
disable_host_rewrite = {}
```

<a id="canonical-2457890880976783101ed0e8cf7006440fb428d7eeb005238bf35858baaef84c"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route / aff9f4559807 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dfef4c625c95d50bfc3fafe525ed326be9a67dffc19a6e137d65d934213f79a7"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route / aff9f4559807 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](resources--workload--reference--group-005.md#canonical-58fbfdd83a4752f4560b8b0452371470c11a31f1a96d635043ff9ae07d4aea92)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-9c3ebce39f55bdb8564389438b789c0d0f73f05949e6d9ae831e29f588360ef4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e52575926d67a8ddabbb29edef23a7cc5840335331acd3e0295bea5257c0d95b"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.http — service.advertise_options.advertise_custom.ports.http_loadbalancer.http / c2efa1484e7c / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.http

<a id="canonical-2adf66cef66038409f7c55bbd7a880778998b200a9265bf76b7515fdc5ea3fa8"></a>

Type: `"object"`. single nested block, Optional.

HTTP Choice. Choice for selecting HTTP proxy.

Upstream description:

Choice for selecting HTTP proxy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("port",
    "port_ranges")}
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
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]"
}
```

Terraform syntax:

```terraform
http {
  # Configure direct properties listed below.
}
```

<a id="canonical-132f5faafdbb5d30d147e3d2be9efe1acad7f0aa082aef609a617124909ef2ac"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.http / c2efa1484e7c / 3

<a id="canonical-914012dfcfe46bc4ee18f17d61f83172dc239730a2209e3fe061f764c47e0168"></a>

<a id="canonical-1e7325c0d90c01d72d72c53d9c02e570f8ad9f9895aa9c835a0dc59caa16328f"></a>

## dns_volterra_managed property — service.advertise_options.advertise_custom.ports.http_loadbalancer.http / c2efa1484e7c / 4

Type: `"bool"`. Optional.

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

<a id="canonical-da3119ba98dd63f9d51a9d1c055ff432e72aab1d6b0a72178f78ba4b43c20d67"></a>

<a id="canonical-c9d1f2664961569b3d4a3ec69749bcbca80b96e0b8e425cd97f66d564e1f0b18"></a>

## port property — service.advertise_options.advertise_custom.ports.http_loadbalancer.http / c2efa1484e7c / 5

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTP port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTP port to Listen.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-b57fe0abb7b3f2f9cb61b7cb2f5182c0c4f4bf6f6e385dd2395e85afadde377b"></a>

<a id="canonical-50f78f7df3b4bcf36050d9d1c95abc1156c459262dad7e86ee4728e670b8e4ca"></a>

## port_ranges property — service.advertise_options.advertise_custom.ports.http_loadbalancer.http / c2efa1484e7c / 6

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

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

<a id="canonical-8f79cb2074395cecb8f41360e57cd9c079be34055f1ca6e99951e55b1b2be605"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.http / c2efa1484e7c / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a4ff3e57dc357963a557202cd20a89e226ef5dea09ff4297c9b3f60cbcf2a7ce"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https — service.advertise_options.advertise_custom.ports.http_loadbalancer.https / 89300d3f0cab / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https

<a id="canonical-08cf1beb2f233b39da103cdbdd44aa8285a106f1fcac31b6c798bd84bee0041d"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting HTTP proxy with bring your own certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("append_server_name",
    "default_header"),
  validators.ConflictingObjectAttributes("append_server_name",
    "pass_through"),
  validators.ConflictingObjectAttributes("append_server_name",
    "server_name"),
  validators.ConflictingObjectAttributes("default_header",
    "pass_through"),
  validators.ConflictingObjectAttributes("default_header",
    "server_name"),
  validators.ConflictingObjectAttributes("default_loadbalancer",
    "non_default_loadbalancer"),
  validators.ConflictingObjectAttributes("disable_path_normalize",
    "enable_path_normalize"),
  validators.ConflictingObjectAttributes("pass_through",
    "server_name"),
  validators.ConflictingObjectAttributes("port",
    "port_ranges"),
  validators.ConflictingObjectAttributes("tls_cert_params",
    "tls_parameters")}
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
  "x-ves-oneof-field-default_lb_choice": "[\"default_loadbalancer\",\"non_default_loadbalancer\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]",
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_parameters\"]"
}
```

Terraform syntax:

```terraform
https {
  # Configure direct properties listed below.
}
```

<a id="canonical-d544c1a47758a5a324926eb635f074ec98dfd1f0a23a6824a6e01f24a18ac2a5"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https / 89300d3f0cab / 3

<a id="canonical-31cd6e2a1d90f053a14a36ad470eeb04b2feb7579bc2004fbddc39aa7aad0d8d"></a>

<a id="canonical-c511875b484e902f7d86e7d09bb1084c01fe2240a68183d7179bbae6701872ac"></a>

## add_hsts property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https / 89300d3f0cab / 4

Type: `"bool"`. Optional.

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

<a id="canonical-31a7ab05f5f0d86a8670b86139756cde45dac73f070eac82aab1dde61c31bb94"></a>

<a id="canonical-e113d461d78e9a5eaf1f6e692c226c2b711aa775ba5de026affd3885611770da"></a>

## append_server_name property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https / 89300d3f0cab / 5

Type: `"string"`. Optional.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Upstream description:

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

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

- [coalescing_options](resources--workload--reference--group-005.md#canonical-0f43d201a9dd0468f4408bcaa56fbc012def6ffb0ee9aa6b2a960cafda941fa7): complete subsection reference.

<a id="canonical-c06fdb549b5d8e0f775fba5d3809b34fa4c011020ea757c07c2ad6680be26e1a"></a>

<a id="canonical-94b8748b99c7aec86ed2fd328942f664b41f79cf2a2eb4169be3d293aa588adc"></a>

## connection_idle_timeout property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https / 89300d3f0cab / 6

Type: `"number"`. Optional.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed.

Upstream description:

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(600000),
}
```

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

- [default_header](resources--workload--reference--group-005.md#canonical-aa7818b0ab415c8de992c06168db771801ba2343321fe90d5c6485aca9490bf9): complete subsection reference.

- [default_loadbalancer](resources--workload--reference--group-005.md#canonical-c7e0f816bc9537a4552728a143bceb2a3001e7a5217c357f46252a4c8a2934df): complete subsection reference.

- [disable_path_normalize](resources--workload--reference--group-005.md#canonical-9c1ca457ee12eee824d315946fcb7e9a74d7afb6a1962849ffe376d790ee638f): complete subsection reference.

- [enable_path_normalize](resources--workload--reference--group-005.md#canonical-d4177bd0f2337c0d6551bb365394cb996c57369d002629a33164e727431c3532): complete subsection reference.

- [http_protocol_options](resources--workload--reference--group-005.md#canonical-72130b36d5da1857877f059316bfbd4028bdcc9e523d2fde50abb5bc6ac9f600): complete subsection reference.

<a id="canonical-b9357229e812d7f712b86eacaa89b687bbf0b1b38a1d0fa8746ac22e38faaca3"></a>

<a id="canonical-945f1d9e68b6da1449daaa681656a88289a43141ce15c8a3d9672a703bb40978"></a>

## http_redirect property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https / 89300d3f0cab / 7

Type: `"bool"`. Optional.

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

- [non_default_loadbalancer](resources--workload--reference--group-005.md#canonical-ce8b23a73374281934de38c133af5f63d47adb01f0277daf1c53c4d59e88a460): complete subsection reference.

- [pass_through](resources--workload--reference--group-005.md#canonical-4874000927edeb74f2da2199509392db0bf501582b64d98b8d22b352286d827d): complete subsection reference.

<a id="canonical-84cfc9320492ec7329ccbcaff5d233dd9446245aa844615bee0b9f52d8da9f3d"></a>

<a id="canonical-a52232f8181f60ce5bf1b18c65b7f808e99634207721118a6052c24e7a094b80"></a>

## port property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https / 89300d3f0cab / 8

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTPS port to Listen.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2ef28ca35ad84c96a03992280c40b55ffb57f1e45a75e5ebe02217f9ae9f8b10"></a>

<a id="canonical-140950576ab3f614277d5bd786b71cfc0cc7cb955490d44088417cc5e3439049"></a>

## port_ranges property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https / 89300d3f0cab / 9

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

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

<a id="canonical-b9e9ebd61d1516e429b3ddef50b1c553f5579507021dd1c3dc24c968b5806b00"></a>

<a id="canonical-dbacea2c63eef1a63c768ff7080eba8bffb5cc01a5fa49b66ef1a82be9b36713"></a>

## server_name property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https / 89300d3f0cab / 10

Type: `"string"`. Optional.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

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

- [tls_cert_params](resources--workload--reference--group-005.md#canonical-137ff8a95d257e7b9a269ba147dbae38b7fe843cc507f89319608cfafe000444): complete subsection reference.

- [tls_parameters](resources--workload--reference--group-006.md#canonical-cbc48f3d69bdae68858971c3bd0f7ef81a2402af198b106a45dd7171932f55b8): complete subsection reference.

<a id="canonical-443089cd21d52de8482aa72a806bf8e9e2c02e38b7e2a58eb8340e86de5a0977"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https / 89300d3f0cab / 11

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-005.md#canonical-0f43d201a9dd0468f4408bcaa56fbc012def6ffb0ee9aa6b2a960cafda941fa7)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_header](resources--workload--reference--group-005.md#canonical-aa7818b0ab415c8de992c06168db771801ba2343321fe90d5c6485aca9490bf9)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_loadbalancer](resources--workload--reference--group-005.md#canonical-c7e0f816bc9537a4552728a143bceb2a3001e7a5217c357f46252a4c8a2934df)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.disable_path_normalize](resources--workload--reference--group-005.md#canonical-9c1ca457ee12eee824d315946fcb7e9a74d7afb6a1962849ffe376d790ee638f)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.enable_path_normalize](resources--workload--reference--group-005.md#canonical-d4177bd0f2337c0d6551bb365394cb996c57369d002629a33164e727431c3532)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-005.md#canonical-72130b36d5da1857877f059316bfbd4028bdcc9e523d2fde50abb5bc6ac9f600)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.non_default_loadbalancer](resources--workload--reference--group-005.md#canonical-ce8b23a73374281934de38c133af5f63d47adb01f0277daf1c53c4d59e88a460)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.pass_through](resources--workload--reference--group-005.md#canonical-4874000927edeb74f2da2199509392db0bf501582b64d98b8d22b352286d827d)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-005.md#canonical-137ff8a95d257e7b9a269ba147dbae38b7fe843cc507f89319608cfafe000444)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-cbc48f3d69bdae68858971c3bd0f7ef81a2402af198b106a45dd7171932f55b8)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-0f43d201a9dd0468f4408bcaa56fbc012def6ffb0ee9aa6b2a960cafda941fa7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-caa1735d10d755d752f9a8b351e88b10e7a19671e7410e695d54759f187545d3"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalesc / c0cbfcfeee5d / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options

<a id="canonical-2082f651b67222f9f7e78ae3ca05c4c1e7d9d98884f736a966cd158437670a5c"></a>

Type: `"object"`. single nested block, Optional.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_coalescing",
    "strict_coalescing")}
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
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

Terraform syntax:

```terraform
coalescing_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-f221b0ce6521b99bafde5031a5934ebcb6546525419cbfba5f3d652d8977d743"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalesc / c0cbfcfeee5d / 3

- [default_coalescing](resources--workload--reference--group-005.md#canonical-854e5c7bba8688010963af027a39aa7adc82d372302fac2276d4629d4a2e6eb6): complete subsection reference.

- [strict_coalescing](resources--workload--reference--group-005.md#canonical-bcbf4d0526970348fa676fa70fae3410e124c802e7d6a5ce4f75161bdfe8426e): complete subsection reference.

<a id="canonical-045414d5240c39171326248960773d1cf6c400e5f8c6c26f9ba880b2dbfb1adf"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalesc / c0cbfcfeee5d / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.default_coalescing](resources--workload--reference--group-005.md#canonical-854e5c7bba8688010963af027a39aa7adc82d372302fac2276d4629d4a2e6eb6)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.strict_coalescing](resources--workload--reference--group-005.md#canonical-bcbf4d0526970348fa676fa70fae3410e124c802e7d6a5ce4f75161bdfe8426e)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-854e5c7bba8688010963af027a39aa7adc82d372302fac2276d4629d4a2e6eb6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-73fa0836112a06a513aa3b78f90585b8c32322e22c1e5f38d683642d44a1c084"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.default_coalescing — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalesc / f6af4ba81da2 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-005.md#canonical-0f43d201a9dd0468f4408bcaa56fbc012def6ffb0ee9aa6b2a960cafda941fa7)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.default_coalescing

<a id="canonical-5ca30ea0a3fbda4b4a8c0a9e21df819718c1057d8897d9c2ee457210df2edc0d"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_coalescing = {}
```

<a id="canonical-5d39950631e08d5e60b9006acd01e900f968ed33caa2fc2dc64f62848068fb4d"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalesc / f6af4ba81da2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f48b894388317a25598f73d9b7b0b3233f7a887f2f0e871d9baec9b2df5b8952"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalesc / f6af4ba81da2 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-005.md#canonical-0f43d201a9dd0468f4408bcaa56fbc012def6ffb0ee9aa6b2a960cafda941fa7)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-bcbf4d0526970348fa676fa70fae3410e124c802e7d6a5ce4f75161bdfe8426e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e3aeb85519e4dcf5392e003454dad601abc8527a4e1519b63bc3ae4f36cc6c23"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.strict_coalescing — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalesc / 9022ebac5f70 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-005.md#canonical-0f43d201a9dd0468f4408bcaa56fbc012def6ffb0ee9aa6b2a960cafda941fa7)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.strict_coalescing

<a id="canonical-fa7aae82cfcb89f2704dcc4c7b37edf613a08b3bb68cb73eae703040d371611e"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
strict_coalescing = {}
```

<a id="canonical-58cc814bf632af7e824322d9c6667869d37ad50c03967fc800d6ef30b0ac9003"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalesc / 9022ebac5f70 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-11e8c93489f1b2f559cbdc70b4d25039022a3ba666677af5d6e657699c379406"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalesc / 9022ebac5f70 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-005.md#canonical-0f43d201a9dd0468f4408bcaa56fbc012def6ffb0ee9aa6b2a960cafda941fa7)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-aa7818b0ab415c8de992c06168db771801ba2343321fe90d5c6485aca9490bf9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-99fb9ffc676b883d2f3cdca6ee31defb089c5890a1f5565078c5994ccf52102f"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_header — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default / 7a4272a828c1 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_header

<a id="canonical-5a1f99e3956de2288fdd08e6dfafb9deeda4124eb00e47a821930ce85a45af49"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_header = {}
```

<a id="canonical-12e84ac3530fc4a882f57e20f62eff4ae8ca0de37369eacd7c3ab04386a572a8"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default / 7a4272a828c1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b1457bdc0b1ade3ca969e4df9257f4d3c1a78150d290bb2fb15a3019b6becf4c"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default / 7a4272a828c1 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-c7e0f816bc9537a4552728a143bceb2a3001e7a5217c357f46252a4c8a2934df"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4cc37ade9fdaf544b3a7c5637ac6a6f8c56fc7743277ff31268c3bbb740f589e"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_loadbalancer — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default / 2325ef9ea9f1 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_loadbalancer

<a id="canonical-2094d4feebaa5b1b8524ecae0eacfd8e93deb70fec86e880ca59b475b6f418df"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_loadbalancer = {}
```

<a id="canonical-e20dd89b2deb6f1ce53c0f9f7f4d196523368f826d649b36d1cd308c5a36e8d6"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default / 2325ef9ea9f1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7486338e3b98a75860ba9c51cbdc109fcd0709a78b979ce2be027e701bd83d8e"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default / 2325ef9ea9f1 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-9c1ca457ee12eee824d315946fcb7e9a74d7afb6a1962849ffe376d790ee638f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-42621fb709726e429bc2b9c073bbec3f23664b602bd178872582d9bca860a82f"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.disable_path_normalize — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.disable / 29566cdfa4d6 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.disable_path_normalize

<a id="canonical-2ebe404e93301dc1039d7aa903dea00f1b918fa28fc6ff06834f317d6b7f061b"></a>

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
disable_path_normalize = {}
```

<a id="canonical-4fe09cc1a86175fbdef02c16fbb743165aeca81cf7266f0109c61e877f726344"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.disable / 29566cdfa4d6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1234cae87f5489c82c46c6ca76848e422fdea23d9b35b79ebf7188a9076a9466"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.disable / 29566cdfa4d6 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-d4177bd0f2337c0d6551bb365394cb996c57369d002629a33164e727431c3532"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b533306001053af0017c491b2a54ae6a34bfea4d29508ea105387c8ac76e5a74"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.enable_path_normalize — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.enable_ / 7b04ec83c203 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.enable_path_normalize

<a id="canonical-37271d8904fbd631fc1155e4529162a947196f6f56c86c92572de189efd091e5"></a>

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
enable_path_normalize = {}
```

<a id="canonical-b5615b04244dfbfd48d83183e64982187fe54f61578ef1cb989eb629703d3d96"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.enable_ / 7b04ec83c203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0f496600bf72cff60370484df142732e25176440e3f168e1207320e2b7952ef3"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.enable_ / 7b04ec83c203 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-72130b36d5da1857877f059316bfbd4028bdcc9e523d2fde50abb5bc6ac9f600"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a404c6f08431c0483c6b7207563d2547c38f2bb09760c1631086d598d5de6f9c"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 4d5e929d6cf6 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options

<a id="canonical-fcd0541ebb61028165ace7e260d0084fd0ac71e37b7e76dd3cd231bfc56cdb07"></a>

Type: `"object"`. single nested block, Optional.

HTTP protocol configuration OPTIONS for downstream connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v1_v2"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v2_only"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_v2",
    "http_protocol_enable_v2_only")}
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
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

Terraform syntax:

```terraform
http_protocol_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-14fd3f568abae2f5f280b0c41fed238fdaf24bc460a6f5b4350455f706787a87"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 4d5e929d6cf6 / 3

- [http_protocol_enable_v1_only](resources--workload--reference--group-005.md#canonical-af8d642d567bbb9dcdb4f6f24f109105b83d1cd4e4af5372e827c400e1262700): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--workload--reference--group-005.md#canonical-6c478c40bd19b66e68e822745242dfc87f874b97394193e614a3f245ac422c1f): complete subsection reference.

- [http_protocol_enable_v2_only](resources--workload--reference--group-005.md#canonical-4a122caa5a6963a3e4d8932b223a89e50c57488bc4cdfcd8a804835705a46269): complete subsection reference.

<a id="canonical-53e845d780e5e0a544360b5805dad733755c8f50b2f83c684ac19bca4d6fb45b"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 4d5e929d6cf6 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-005.md#canonical-af8d642d567bbb9dcdb4f6f24f109105b83d1cd4e4af5372e827c400e1262700)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2](resources--workload--reference--group-005.md#canonical-6c478c40bd19b66e68e822745242dfc87f874b97394193e614a3f245ac422c1f)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only](resources--workload--reference--group-005.md#canonical-4a122caa5a6963a3e4d8932b223a89e50c57488bc4cdfcd8a804835705a46269)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-af8d642d567bbb9dcdb4f6f24f109105b83d1cd4e4af5372e827c400e1262700"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb71debbfe80ccd778bdbc1ee1035561671925969e487a54fcf9255320095881"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 3585720c8761 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-005.md#canonical-72130b36d5da1857877f059316bfbd4028bdcc9e523d2fde50abb5bc6ac9f600)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-11fe73096ca12bd988ce75cfd3ff1fecd54c99234d1ef0d25fd5bdf8078ead90"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
http_protocol_enable_v1_only {
  # Configure direct properties listed below.
}
```

<a id="canonical-3bcac3f48ed73f3fdf745f96ac7dfadb3ef13665318feb4024c060bbbaed14f4"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 3585720c8761 / 3

- [header_transformation](resources--workload--reference--group-005.md#canonical-3b2f37c3365abb0a4fd1801ca1c3067f5c17e95be7e27334be6d2ebefd69e91e): complete subsection reference.

<a id="canonical-c68a56409691725993f62202699a1e9ad7d559ebf640cf5630f4cf620de48a0f"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 3585720c8761 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-005.md#canonical-3b2f37c3365abb0a4fd1801ca1c3067f5c17e95be7e27334be6d2ebefd69e91e)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-005.md#canonical-72130b36d5da1857877f059316bfbd4028bdcc9e523d2fde50abb5bc6ac9f600)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-3b2f37c3365abb0a4fd1801ca1c3067f5c17e95be7e27334be6d2ebefd69e91e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-46e65ea6f013a8205110f1c76c34bf11e542ac1b3903d27f50c4a7542958c3a6"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / d074e7bc4578 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-005.md#canonical-72130b36d5da1857877f059316bfbd4028bdcc9e523d2fde50abb5bc6ac9f600)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-005.md#canonical-af8d642d567bbb9dcdb4f6f24f109105b83d1cd4e4af5372e827c400e1262700)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-3625c3855106b8fd78ab7d5b08067dcba3610762e69ef616a98666befaed53ee"></a>

Type: `"object"`. single nested block, Optional.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_header_transformation",
    "preserve_case_header_transformation"),
  validators.ConflictingObjectAttributes("default_header_transformation",
    "proper_case_header_transformation"),
  validators.ConflictingObjectAttributes("preserve_case_header_transformation",
    "proper_case_header_transformation")}
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
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

Terraform syntax:

```terraform
header_transformation {
  # Configure direct properties listed below.
}
```

<a id="canonical-e585df899c9006383286ea0b25e59144ffd307829c715221cddde1ee5c3d21ea"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / d074e7bc4578 / 3

- [default_header_transformation](resources--workload--reference--group-005.md#canonical-a86ad7a54b9126181052ae506aa1654038d1e41df04ba431a2cf459a29a23c46): complete subsection reference.

- [preserve_case_header_transformation](resources--workload--reference--group-005.md#canonical-e57e09a5f484c374c7d064dcb4840ce75a87921607dd46243350fcf3260b06e2): complete subsection reference.

- [proper_case_header_transformation](resources--workload--reference--group-005.md#canonical-355dc06796652df80c9458ddb2f3b62455bf3be98c5211934e94ffe78d82a925): complete subsection reference.

<a id="canonical-eb9cd4ba33375d571d2f6ebf961bc4b8dbc732ea9533242bdb85b1210eec0893"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / d074e7bc4578 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--workload--reference--group-005.md#canonical-a86ad7a54b9126181052ae506aa1654038d1e41df04ba431a2cf459a29a23c46)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--workload--reference--group-005.md#canonical-e57e09a5f484c374c7d064dcb4840ce75a87921607dd46243350fcf3260b06e2)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--workload--reference--group-005.md#canonical-355dc06796652df80c9458ddb2f3b62455bf3be98c5211934e94ffe78d82a925)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-005.md#canonical-af8d642d567bbb9dcdb4f6f24f109105b83d1cd4e4af5372e827c400e1262700)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-a86ad7a54b9126181052ae506aa1654038d1e41df04ba431a2cf459a29a23c46"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-74beb4b5c18e3ef404cff96cb31332f3ec4b5f044b61d43bca35e7dda65bb6b8"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 789ffd67d86b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-005.md#canonical-72130b36d5da1857877f059316bfbd4028bdcc9e523d2fde50abb5bc6ac9f600)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-005.md#canonical-af8d642d567bbb9dcdb4f6f24f109105b83d1cd4e4af5372e827c400e1262700)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-005.md#canonical-3b2f37c3365abb0a4fd1801ca1c3067f5c17e95be7e27334be6d2ebefd69e91e)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-bd9c43d4f8df91a88b04c36eaaba68b6d965c0bb55036986f06b14430589485e"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_header_transformation = {}
```

<a id="canonical-0f5695a0e5d2ebb641c00326f496c57dd6da3082de7a3fe4991a3e2bbf4bf543"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 789ffd67d86b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8de3396f6dd654a5fdb061e5ab692abc798fa300bd219eb8e811fa8dbf88a1f9"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 789ffd67d86b / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-005.md#canonical-3b2f37c3365abb0a4fd1801ca1c3067f5c17e95be7e27334be6d2ebefd69e91e)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-e57e09a5f484c374c7d064dcb4840ce75a87921607dd46243350fcf3260b06e2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-91cd444fc25231329daad541c891c49c43895cf07327884a748b600e8813075d"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 4582a5a2cd2d / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-005.md#canonical-72130b36d5da1857877f059316bfbd4028bdcc9e523d2fde50abb5bc6ac9f600)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-005.md#canonical-af8d642d567bbb9dcdb4f6f24f109105b83d1cd4e4af5372e827c400e1262700)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-005.md#canonical-3b2f37c3365abb0a4fd1801ca1c3067f5c17e95be7e27334be6d2ebefd69e91e)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-afe495afb36ff37a42b76bd6fcf6348548e5e4c00f1d417875cec5c09a921276"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
preserve_case_header_transformation = {}
```

<a id="canonical-8efa61df847a156c3082363d5f89c51bfc2f6d374475d9c4afb7d443cfd20bee"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 4582a5a2cd2d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1719a11636477bf02359ac0251abee9b74ce1c41b736fa237f91c386a3f40099"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 4582a5a2cd2d / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-005.md#canonical-3b2f37c3365abb0a4fd1801ca1c3067f5c17e95be7e27334be6d2ebefd69e91e)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-355dc06796652df80c9458ddb2f3b62455bf3be98c5211934e94ffe78d82a925"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-17fc9101b1d31509f08f40f91278af3e202a6cbfcba2d8e33d9c3aa58ef1a3aa"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 3c3a339204d8 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-005.md#canonical-72130b36d5da1857877f059316bfbd4028bdcc9e523d2fde50abb5bc6ac9f600)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-005.md#canonical-af8d642d567bbb9dcdb4f6f24f109105b83d1cd4e4af5372e827c400e1262700)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-005.md#canonical-3b2f37c3365abb0a4fd1801ca1c3067f5c17e95be7e27334be6d2ebefd69e91e)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-b122055b0075f32531add653fd9f28285c4696f964e0c5362d29dd708d4b104a"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
proper_case_header_transformation = {}
```

<a id="canonical-a6deafec2acbfb9326770b366c9bffc58eb0e53feb983dea1c77bf07d87eb446"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 3c3a339204d8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-299ca067914405d49dec06781e7d7cbc68e27fb2c7d151cd08a3de250a504b9e"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 3c3a339204d8 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-005.md#canonical-3b2f37c3365abb0a4fd1801ca1c3067f5c17e95be7e27334be6d2ebefd69e91e)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-6c478c40bd19b66e68e822745242dfc87f874b97394193e614a3f245ac422c1f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4adc5c62c4b4fdc2f67166df4d9592f7c3086d9e554c7312b00f875d4149531b"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2 — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 863117cc7a25 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-005.md#canonical-72130b36d5da1857877f059316bfbd4028bdcc9e523d2fde50abb5bc6ac9f600)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-0e6469a176e437292015b9b15962ae147928b674539da363dc68e47a76417891"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
http_protocol_enable_v1_v2 = {}
```

<a id="canonical-9861b241575a3bef006968dfad4ddaa149b3d7cb6590abbae76471e5622c8d75"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 863117cc7a25 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a172f2d21625942d21bf652d7f2db0326e9bf0ff2f9e9f20837a6827a1b5158d"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 863117cc7a25 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-005.md#canonical-72130b36d5da1857877f059316bfbd4028bdcc9e523d2fde50abb5bc6ac9f600)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-4a122caa5a6963a3e4d8932b223a89e50c57488bc4cdfcd8a804835705a46269"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0530a40e4166d29176bdb6e66ce5598162e49175a9747e782c518f7a258bd91b"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 9153165639aa / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-005.md#canonical-72130b36d5da1857877f059316bfbd4028bdcc9e523d2fde50abb5bc6ac9f600)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-503875615a53b39b85809c2abd9c299889d1749ae7a222fe8a87b77306a0f0bc"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
http_protocol_enable_v2_only = {}
```

<a id="canonical-4ca9a1e99d28c58569a6bcb8d740cc9be67e81257037245a4e07a575435795cb"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 9153165639aa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f4a38b0390d6d8b3b15874b9e130f2d8bccb1ccc9e0d2c2cfa0944b2af3c084b"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_pr / 9153165639aa / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-005.md#canonical-72130b36d5da1857877f059316bfbd4028bdcc9e523d2fde50abb5bc6ac9f600)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-ce8b23a73374281934de38c133af5f63d47adb01f0277daf1c53c4d59e88a460"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9f8d8cb15b9d59f78c5d9b938df589e737c8660229c02fb42f87f0b8e1d2e2f0"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.non_default_loadbalancer — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.non_def / 5806bf862471 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.non_default_loadbalancer

<a id="canonical-d33a13c8237b492f9144b894e0263fcc7e46c26bf17c14206cec021acce72b00"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
non_default_loadbalancer = {}
```

<a id="canonical-61a5c68849dab444eff81bed9cffb397c0a3fd7fe45254c0c4d42a7d537a67de"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.non_def / 5806bf862471 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-373676eb8125109c0060d4f581664aff7730368d70014b8a8f9656ed5aa5d3be"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.non_def / 5806bf862471 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-4874000927edeb74f2da2199509392db0bf501582b64d98b8d22b352286d827d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-47c9d436375b5c87ee2debc964e3d1d6d8137218935d6bb9d13d7ad7d20fb517"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.pass_through — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.pass_th / 6c8ebe78220d / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.pass_through

<a id="canonical-0601f3c1e2a6835d9c4e6b2e03db417210518b038397ef3a17a423f5cad3ad52"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
pass_through = {}
```

<a id="canonical-0aa4be984e0ecf02e45c53d52af6368f19971a6478460a10e56544928bd429aa"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.pass_th / 6c8ebe78220d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f659591de225641461930b025b2198fa0354e819ca113dd7101b1346d90e2159"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.pass_th / 6c8ebe78220d / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-137ff8a95d257e7b9a269ba147dbae38b7fe843cc507f89319608cfafe000444"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4d361bf811a51610c95b8a7de478458e05671643a4659ed945664034434bfda6"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 498b804e2a84 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params

<a id="canonical-a08b5dc7df41e421dc79fde316f075991a6ae5d51f8934481787e5e557d559a0"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls cert params.

Upstream description:

Select TLS Parameters and Certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("certificates"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls")}
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
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
tls_cert_params {
  # Configure direct properties listed below.
}
```
