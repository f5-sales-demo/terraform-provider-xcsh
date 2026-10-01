---
page_title: "xcsh_k8s_pod_security_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_pod_security_policy reference."
---

# xcsh_k8s_pod_security_policy reference

<a id="canonical-3ca2e681586afce0edce8cbc8836a556f7345a01b126e0b7a741a4a76865647f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a083de4585953bd3770844e8d21b43492e2d82e7c97afc493c1906059288c42d"></a>

## Property reference — Property reference / 26fc0728cbac / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)
- Property reference

<a id="canonical-c3cf48628520f38076b8d0645c0e3986ab11367375761789631d6f6b378947fd"></a>

## Direct properties — Property reference / 26fc0728cbac / 3

<a id="canonical-31aabfbac60bde026d01e71a0587d4977bbb63d3fb69ac6546fbfe3368346dbd"></a>

<a id="canonical-e7604cd1d45dd0f4e811a3b3fcb4c038444d4b5902ec77093dc0c5a3029c9bff"></a>

## annotations property — Property reference / 26fc0728cbac / 4

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

<a id="canonical-d9a64c8754f799c29448b426f3cdd00391b343d5205b0b1682260e0b602f3e69"></a>

<a id="canonical-e3ea50330f67d4a7876245305b1753a5955a5050fc761249058ebdddf2f03944"></a>

## description property — Property reference / 26fc0728cbac / 5

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

<a id="canonical-7ecb685622f5118f621cf6e5e16a82232038806fc97bbc1764f34f9ec769c2a4"></a>

<a id="canonical-7d5dc4983e3f69d0b0771b0321a5cc22153cdce1c26b0404dc39cdffa332fe36"></a>

## disable property — Property reference / 26fc0728cbac / 6

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

<a id="canonical-db9c9f2c0831ef4b5dd5cba6ed378fd5d7b34810c27b1c57bd82b1ef2a0d60bf"></a>

<a id="canonical-53d3ee17ad6d9648858fcae39e8371b958c8b15f4f14d7f8b1c8853bff01112b"></a>

## id property — Property reference / 26fc0728cbac / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-459c63ba0a87f3b5de84540bdea80fd62e5e329556c92fe69164202e2a43a035"></a>

<a id="canonical-1a6a73f1bd50ff333aa41460faaa427a7d02519efdb76834705564b8d2519579"></a>

## labels property — Property reference / 26fc0728cbac / 8

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

<a id="canonical-d376eab5ef9b20efd5ea80b4c70145f9350238380b1d299522e64c9f59e2a3f6"></a>

<a id="canonical-98867d7153cfb8042453d0ed88fb3f16f9cf565516141ef48608632ccac6993e"></a>

## name property — Property reference / 26fc0728cbac / 9

Type: `"string"`. Required.

Name of the K8S Pod Security Policy. Must be unique within the namespace.

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

<a id="canonical-7295f36599811701773f58d6426c7ee3d65988f122f4c36cd8e8a5f6c4bfb062"></a>

<a id="canonical-5eb0d90d9b61d676f292a9adab8ba023da878aa346ddd1a57bfdc99461d1955d"></a>

## namespace property — Property reference / 26fc0728cbac / 10

Type: `"string"`. Required.

Namespace where the K8S Pod Security Policy is created.

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

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e): complete subsection reference.

- [timeouts](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0d0bb5077a31feaba9ea04921fd608059998b40cae432fc7fddba2580ee3d459): complete subsection reference.

<a id="canonical-45e2d558a17ca3f521e12d74c4f200048047a12082d5171c4f6d8b9819028488"></a>

<a id="canonical-bed6dfe646872ca7aa142be37a8d697a3daa229773b9c2f29d44bc2150de20f4"></a>

## yaml property — Property reference / 26fc0728cbac / 11

Type: `"string"`. Optional, Computed.

Exclusive with \[psp\_spec\] K8s YAML for Pod Security Policy.

Upstream description:

Exclusive with \[psp\_spec\] K8s YAML for Pod Security Policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 4096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "Valid parseable YAML",
    "maxLength": 4096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "validation": {
      "customRule": "Must be valid YAML"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-7e1a2a34857b8a57c5c0ede7d21370d513befe12579babe1123377863e26da34"></a>

## All schema paths — Property reference / 26fc0728cbac / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--k8s_pod_security_policy--reference--group-001.md#canonical-31aabfbac60bde026d01e71a0587d4977bbb63d3fb69ac6546fbfe3368346dbd) |
| `description` | [description](resources--k8s_pod_security_policy--reference--group-001.md#canonical-d9a64c8754f799c29448b426f3cdd00391b343d5205b0b1682260e0b602f3e69) |
| `disable` | [disable](resources--k8s_pod_security_policy--reference--group-001.md#canonical-7ecb685622f5118f621cf6e5e16a82232038806fc97bbc1764f34f9ec769c2a4) |
| `id` | [id](resources--k8s_pod_security_policy--reference--group-001.md#canonical-db9c9f2c0831ef4b5dd5cba6ed378fd5d7b34810c27b1c57bd82b1ef2a0d60bf) |
| `labels` | [labels](resources--k8s_pod_security_policy--reference--group-001.md#canonical-459c63ba0a87f3b5de84540bdea80fd62e5e329556c92fe69164202e2a43a035) |
| `name` | [name](resources--k8s_pod_security_policy--reference--group-001.md#canonical-d376eab5ef9b20efd5ea80b4c70145f9350238380b1d299522e64c9f59e2a3f6) |
| `namespace` | [namespace](resources--k8s_pod_security_policy--reference--group-001.md#canonical-7295f36599811701773f58d6426c7ee3d65988f122f4c36cd8e8a5f6c4bfb062) |
| `psp_spec` | [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-afed82101e7c1055f998cb95ae0f0ef2a7c7d64ad8c6c16b242bf14977012767) |
| `psp_spec.allow_privilege_escalation` | [psp_spec.allow_privilege_escalation](resources--k8s_pod_security_policy--reference--group-001.md#canonical-cde2c03261ec3735e5a82748320f01b072e968262e0ef4267c003ad439d8702d) |
| `psp_spec.allowed_capabilities` | [psp_spec.allowed_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-bdfd1fce5b9ae5da698804a90985a5ca033048542e55159f09c31950981d032d) |
| `psp_spec.allowed_capabilities.capabilities` | [psp_spec.allowed_capabilities.capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-e67cdb0f96b06b8987a22a047e70b4a6f9809a152fb580a81ed2b42f210abbd7) |
| `psp_spec.allowed_csi_drivers` | [psp_spec.allowed_csi_drivers](resources--k8s_pod_security_policy--reference--group-001.md#canonical-f9728c9195613efd4317d83a077f2831d5d691e5a7873a9a53e749fb3eda3d5f) |
| `psp_spec.allowed_flex_volumes` | [psp_spec.allowed_flex_volumes](resources--k8s_pod_security_policy--reference--group-001.md#canonical-f9f783cbf919599610ce11abca1188218498e2d7ea0ac0f8dab2b4c2853b0716) |
| `psp_spec.allowed_host_paths` | [psp_spec.allowed_host_paths](resources--k8s_pod_security_policy--reference--group-001.md#canonical-da3395337800e9f270a2754b519eacc6422cb27023a72798e6f2ac0ba5084074) |
| `psp_spec.allowed_host_paths.path_prefix` | [psp_spec.allowed_host_paths.path_prefix](resources--k8s_pod_security_policy--reference--group-001.md#canonical-74e4bcc01c9903019da5fae7e7732c45a3dcf52fcfe43646fc25d3b962b729c4) |
| `psp_spec.allowed_host_paths.read_only` | [psp_spec.allowed_host_paths.read_only](resources--k8s_pod_security_policy--reference--group-001.md#canonical-447bc81c21d683a2b18249f8114f816c02e471d38f718ce88af493dc04b11beb) |
| `psp_spec.allowed_proc_mounts` | [psp_spec.allowed_proc_mounts](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2ff6c425fd2a85440e814f4b05ef2761061e0e5f6a5d68e6bf44dc4ede05a5ea) |
| `psp_spec.allowed_unsafe_sysctls` | [psp_spec.allowed_unsafe_sysctls](resources--k8s_pod_security_policy--reference--group-001.md#canonical-262d383bbd36bd176bf9eff061edd2fccaacdb87943cc89f856619bbe802d38c) |
| `psp_spec.default_allow_privilege_escalation` | [psp_spec.default_allow_privilege_escalation](resources--k8s_pod_security_policy--reference--group-001.md#canonical-a8a7fa76042d4baabc2ccecd08401f204f3d9e8d2b03850bf960ec720fa42cf3) |
| `psp_spec.default_capabilities` | [psp_spec.default_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-9f60462a9f4ac3764b992806ffea75a93e60e2560762b32c45b007bcf8c837c8) |
| `psp_spec.default_capabilities.capabilities` | [psp_spec.default_capabilities.capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-b89b98c824e8c1b3581b004c9b985222d0a52db47033bc2a50bcc23567d5041d) |
| `psp_spec.drop_capabilities` | [psp_spec.drop_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-8f5242195018aabb5f407307391ff58d3286fb913073b0dbe7a0c95957ee402c) |
| `psp_spec.drop_capabilities.capabilities` | [psp_spec.drop_capabilities.capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-a1f92fc17d0fe5169fdc963b49be47b54c457cddc6296a6b2560367408032476) |
| `psp_spec.forbidden_sysctls` | [psp_spec.forbidden_sysctls](resources--k8s_pod_security_policy--reference--group-001.md#canonical-e3e5b500f01cbe2611c33afa516e57a18313885a6817892d12328b88d6ef40ee) |
| `psp_spec.fs_group_strategy_options` | [psp_spec.fs_group_strategy_options](resources--k8s_pod_security_policy--reference--group-001.md#canonical-445e6cbbe92febaaa68dc07fed386831f466f97f7a385bdd9713add9a6b1fb66) |
| `psp_spec.fs_group_strategy_options.id_ranges` | [psp_spec.fs_group_strategy_options.id_ranges](resources--k8s_pod_security_policy--reference--group-001.md#canonical-eb06e44b2a61a48b9a5e2c33dbf975ecf7ed8eeb58da1f04038139d4d01ee29e) |
| `psp_spec.fs_group_strategy_options.id_ranges.max_id` | [psp_spec.fs_group_strategy_options.id_ranges.max_id](resources--k8s_pod_security_policy--reference--group-001.md#canonical-7d20c6bdd927e8b364f53055003513aad7beb1657557a1030e61caccfce4aea4) |
| `psp_spec.fs_group_strategy_options.id_ranges.min_id` | [psp_spec.fs_group_strategy_options.id_ranges.min_id](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3ebb1ade79fd3d434514de424d7cb60344c20c1974c0578c0dc676dc36f145fd) |
| `psp_spec.fs_group_strategy_options.rule` | [psp_spec.fs_group_strategy_options.rule](resources--k8s_pod_security_policy--reference--group-001.md#canonical-bd3fe1a8f19b4a27df99f616e44805a564ade7f9080482c4539b3053f1e381bf) |
| `psp_spec.host_ipc` | [psp_spec.host_ipc](resources--k8s_pod_security_policy--reference--group-001.md#canonical-7f7891711edfdd3330042dbb337712f07c95934c8ed9c8f405c9907d517b280e) |
| `psp_spec.host_network` | [psp_spec.host_network](resources--k8s_pod_security_policy--reference--group-001.md#canonical-49f66fa2e1bd613c80671c509d21e23e548ad285241beeaff376135a7cb0440f) |
| `psp_spec.host_pid` | [psp_spec.host_pid](resources--k8s_pod_security_policy--reference--group-001.md#canonical-4c4d8ec1fe5b7b1a997df77afa6236e6c3ae9a1e6d6fcac9652f49b7cfa03928) |
| `psp_spec.host_port_ranges` | [psp_spec.host_port_ranges](resources--k8s_pod_security_policy--reference--group-001.md#canonical-f48a697341e8515e0334ab77bcbf357f91562c5a071639858c62d53e818ec6aa) |
| `psp_spec.no_allowed_capabilities` | [psp_spec.no_allowed_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-ab77591fa945d9263d2347b60a5f3955845efe1ad10c4a37c9da3402934be3b8) |
| `psp_spec.no_default_capabilities` | [psp_spec.no_default_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2d2789d93284091ca488d2d9d6cbb238a703b9a9d2942918deebd9a038f533ef) |
| `psp_spec.no_drop_capabilities` | [psp_spec.no_drop_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-5b71f47c4eee6489b097b1ef9ff13c355bcd9f0ae19e0368a5bbe6a21ef994c1) |
| `psp_spec.no_fs_groups` | [psp_spec.no_fs_groups](resources--k8s_pod_security_policy--reference--group-001.md#canonical-8f3f82cd3e686c13846052f7c653dd9156ff4114dd9b5d20dd251d8cf486984e) |
| `psp_spec.no_run_as_group` | [psp_spec.no_run_as_group](resources--k8s_pod_security_policy--reference--group-001.md#canonical-88b3e4fa92b23d54b010a36d514aedbaa04d80db1db3cc532d8c113d6aa18cb6) |
| `psp_spec.no_run_as_user` | [psp_spec.no_run_as_user](resources--k8s_pod_security_policy--reference--group-001.md#canonical-64243e16e92de50c1266f91d9c7b066b854ef365b15801ef784e5a0f859bc68e) |
| `psp_spec.no_runtime_class` | [psp_spec.no_runtime_class](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2b7876d89f0282beabe4a2f154505c222d29a526060f0cc00e3f66f422419f0e) |
| `psp_spec.no_se_linux_options` | [psp_spec.no_se_linux_options](resources--k8s_pod_security_policy--reference--group-001.md#canonical-9f8f221400e983582445146bab6191f79c77e3a89922d0837595ee38ec274fcf) |
| `psp_spec.no_supplemental_groups` | [psp_spec.no_supplemental_groups](resources--k8s_pod_security_policy--reference--group-001.md#canonical-6b2bc15332b788837b448d8aac80c0c32386f32a9e838c54320e1f037b1780b9) |
| `psp_spec.privileged` | [psp_spec.privileged](resources--k8s_pod_security_policy--reference--group-001.md#canonical-88ad95491d8b28a335e197a48d96a6a81c63be8e77cc6c0b9689feea4d75a775) |
| `psp_spec.read_only_root_filesystem` | [psp_spec.read_only_root_filesystem](resources--k8s_pod_security_policy--reference--group-001.md#canonical-99fab6f615dda13b324692dfcd38d8419896902344982289b54fdca3048bad46) |
| `psp_spec.run_as_group` | [psp_spec.run_as_group](resources--k8s_pod_security_policy--reference--group-001.md#canonical-32c6f232d49481b9d72b308b0d9124f3b5580a503f7c2e68cd3bd4252c2a54fd) |
| `psp_spec.run_as_group.id_ranges` | [psp_spec.run_as_group.id_ranges](resources--k8s_pod_security_policy--reference--group-001.md#canonical-dffcb361d095ae0b308b6da8de38d914b07311ebebc56ae9dca4a6cf955c165e) |
| `psp_spec.run_as_group.id_ranges.max_id` | [psp_spec.run_as_group.id_ranges.max_id](resources--k8s_pod_security_policy--reference--group-001.md#canonical-4ef6a4fa19c1d65095153e120b385b27cf2f0866d89f0949faf7f64c2cdf70be) |
| `psp_spec.run_as_group.id_ranges.min_id` | [psp_spec.run_as_group.id_ranges.min_id](resources--k8s_pod_security_policy--reference--group-001.md#canonical-41ae3802bbc9270f3fefd2ad6e1c8f61ad65f92af8e95bd695442cab807bf81d) |
| `psp_spec.run_as_group.rule` | [psp_spec.run_as_group.rule](resources--k8s_pod_security_policy--reference--group-001.md#canonical-d8bd241b7f20a8245e7e6062ff89168b4a52b73b37eed4aca590835a1212d9c9) |
| `psp_spec.run_as_user` | [psp_spec.run_as_user](resources--k8s_pod_security_policy--reference--group-001.md#canonical-78c630ad662e9c638ac1143363898842a1841f171c3807fc745ad12eabb17c49) |
| `psp_spec.run_as_user.id_ranges` | [psp_spec.run_as_user.id_ranges](resources--k8s_pod_security_policy--reference--group-001.md#canonical-d6866f39b08fbd481d3752e120fc606a1f8e6795ddd4cefc230cb92a824e4c98) |
| `psp_spec.run_as_user.id_ranges.max_id` | [psp_spec.run_as_user.id_ranges.max_id](resources--k8s_pod_security_policy--reference--group-001.md#canonical-629dc7e22352ddd9c2661ff63baf49636d32d056add40a754aa03e4f2503b6c5) |
| `psp_spec.run_as_user.id_ranges.min_id` | [psp_spec.run_as_user.id_ranges.min_id](resources--k8s_pod_security_policy--reference--group-001.md#canonical-7632f6a2d975e97046849cb48315e3c91c7ba4db14df7d1fe06bde4c6d3c117d) |
| `psp_spec.run_as_user.rule` | [psp_spec.run_as_user.rule](resources--k8s_pod_security_policy--reference--group-001.md#canonical-983582a56b8cb35b700cb8d226a3b9811fbdb57d36569666188e1e11b32290be) |
| `psp_spec.supplemental_groups` | [psp_spec.supplemental_groups](resources--k8s_pod_security_policy--reference--group-001.md#canonical-66c2a86f8da53f24f3676bb2bead442cee666e1bf0b6f397818e888ac5fd2843) |
| `psp_spec.supplemental_groups.id_ranges` | [psp_spec.supplemental_groups.id_ranges](resources--k8s_pod_security_policy--reference--group-001.md#canonical-38acbba07d6d6ee392ac928b5aa94a5a2bc257077e5b27b59e212e00cd332c51) |
| `psp_spec.supplemental_groups.id_ranges.max_id` | [psp_spec.supplemental_groups.id_ranges.max_id](resources--k8s_pod_security_policy--reference--group-001.md#canonical-fdadfcaa67ef914614e8de12341fd96c8477864bd7e2250fc72989acbda72279) |
| `psp_spec.supplemental_groups.id_ranges.min_id` | [psp_spec.supplemental_groups.id_ranges.min_id](resources--k8s_pod_security_policy--reference--group-001.md#canonical-d95276fb0a5bc35b35cf80178da0a23cebc35baf9734cc726df888657f830aea) |
| `psp_spec.supplemental_groups.rule` | [psp_spec.supplemental_groups.rule](resources--k8s_pod_security_policy--reference--group-001.md#canonical-5be18dfbb8d1a8086686d2d502211236bfecefc2172823607b9d0c5738c6018d) |
| `psp_spec.volumes` | [psp_spec.volumes](resources--k8s_pod_security_policy--reference--group-001.md#canonical-58fdf20e5e672e1772537bf6624b60bb129487f19ea40315c39c461d3d35c309) |
| `timeouts` | [timeouts](resources--k8s_pod_security_policy--reference--group-001.md#canonical-09b1578bcadc482f5cb3d0d631a9e17250c8a839bf47af9df10907296c2e5aea) |
| `timeouts.create` | [timeouts.create](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2f141ffca4fe2a3afe67c772f7cf4e3b8a21e68576f9dfc92fc4bd517429bf8f) |
| `timeouts.delete` | [timeouts.delete](resources--k8s_pod_security_policy--reference--group-001.md#canonical-7c990a5a0c7a9e6ca94b0cf1b69a726e291299e54dbf530382b58a9c7249a99b) |
| `timeouts.read` | [timeouts.read](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1a519a7d960bf83d846500ed715f662c6fbc3ab43356f06f4d5c44e5a4e2adf2) |
| `timeouts.update` | [timeouts.update](resources--k8s_pod_security_policy--reference--group-001.md#canonical-73556f12d92d4ee4cc4353392cf070603d837d9894fd9fcc19c2d563ee9faf53) |
| `yaml` | [yaml](resources--k8s_pod_security_policy--reference--group-001.md#canonical-45e2d558a17ca3f521e12d74c4f200048047a12082d5171c4f6d8b9819028488) |

<a id="canonical-14ee1034de9a7c8b78057f24dfbfad7ab782cc183b2259f5b2c7c6569aa14b4e"></a>

## Next pages — Property reference / 26fc0728cbac / 13

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- [timeouts](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0d0bb5077a31feaba9ea04921fd608059998b40cae432fc7fddba2580ee3d459)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)

<a id="canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-27150c5d67d957d0e6ba30388484698ba91987a2961a2c9e6d327857388e8f3b"></a>

## psp_spec — psp_spec / c699886c4d9c / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3ca2e681586afce0edce8cbc8836a556f7345a01b126e0b7a741a4a76865647f)
- psp_spec

<a id="canonical-afed82101e7c1055f998cb95ae0f0ef2a7c7d64ad8c6c16b242bf14977012767"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: psp\_spec, yaml\] Pod Security Policy Specification. Form based pod security specification.

Upstream description:

Form based pod security specification.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("allowed_capabilities",
    "no_allowed_capabilities"),
  validators.ConflictingObjectAttributes("default_capabilities",
    "no_default_capabilities"),
  validators.ConflictingObjectAttributes("drop_capabilities",
    "no_drop_capabilities"),
  validators.ConflictingObjectAttributes("fs_group_strategy_options",
    "no_fs_groups"),
  validators.ConflictingObjectAttributes("no_run_as_group",
    "run_as_group"),
  validators.ConflictingObjectAttributes("no_run_as_user",
    "run_as_user"),
  validators.ConflictingObjectAttributes("no_supplemental_groups",
    "supplemental_groups")}
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
  "x-ves-oneof-field-allowed_capabilities_choice": "[\"allowed_capabilities\",\"no_allowed_capabilities\"]",
  "x-ves-oneof-field-default_capabilities_choice": "[\"default_capabilities\",\"no_default_capabilities\"]",
  "x-ves-oneof-field-drop_capabilities_choice": "[\"drop_capabilities\",\"no_drop_capabilities\"]",
  "x-ves-oneof-field-fs_group_choice": "[\"fs_group_strategy_options\",\"no_fs_groups\"]",
  "x-ves-oneof-field-group_choice": "[\"no_run_as_group\",\"run_as_group\"]",
  "x-ves-oneof-field-runtime_class_choice": "[\"no_runtime_class\"]",
  "x-ves-oneof-field-se_linux_choice": "[\"no_se_linux_options\"]",
  "x-ves-oneof-field-supplemental_group_choice": "[\"no_supplemental_groups\",\"supplemental_groups\"]",
  "x-ves-oneof-field-user_choice": "[\"no_run_as_user\",\"run_as_user\"]"
}
```

OneOf alternatives in this subsection:

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-afed82101e7c1055f998cb95ae0f0ef2a7c7d64ad8c6c16b242bf14977012767)
- [yaml](resources--k8s_pod_security_policy--reference--group-001.md#canonical-45e2d558a17ca3f521e12d74c4f200048047a12082d5171c4f6d8b9819028488)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
psp_spec {
  # Configure direct properties listed below.
}
```

<a id="canonical-6455ca65bef21be76c36cf29e91034759d711eae00a2b92f8b1cc115df6b89dc"></a>

## Direct properties — psp_spec / c699886c4d9c / 3

<a id="canonical-cde2c03261ec3735e5a82748320f01b072e968262e0ef4267c003ad439d8702d"></a>

<a id="canonical-560efaa7c45d15335b82f32a9ea9063c36ee00651604e938dd9291c8552a0241"></a>

## allow_privilege_escalation property — psp_spec / c699886c4d9c / 4

Type: `"bool"`. Optional.

Pod can request to privilege escalation.

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

- [allowed_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-84b9a80fdd498440ac1ae1ff0cdb8f452ad978f195755e6bfce4996fbcff44d4): complete subsection reference.

<a id="canonical-f9728c9195613efd4317d83a077f2831d5d691e5a7873a9a53e749fb3eda3d5f"></a>

<a id="canonical-d3bceb49eabc3128362b764d222bf418682581d74c62ebc744d5bac2cef603d3"></a>

## allowed_csi_drivers property — psp_spec / c699886c4d9c / 5

Type: `["list", "string"]`. Optional.

Restrict the available CSI drivers for POD, default all drivers are available.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(8),
}
```

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-f9f783cbf919599610ce11abca1188218498e2d7ea0ac0f8dab2b4c2853b0716"></a>

<a id="canonical-502a0ac47f9fa97fa5627dc756cfee7033be55d0efb5ba562a010d82169d9735"></a>

## allowed_flex_volumes property — psp_spec / c699886c4d9c / 6

Type: `["list", "string"]`. Optional.

Restrict list of Flex volumes, default all volumes are allowed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(8),
}
```

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [allowed_host_paths](resources--k8s_pod_security_policy--reference--group-001.md#canonical-bf013790bfc3a2afd2cedf8e7b7d670f43e0f55a7503a1a2b6189836d26b47e4): complete subsection reference.

<a id="canonical-2ff6c425fd2a85440e814f4b05ef2761061e0e5f6a5d68e6bf44dc4ede05a5ea"></a>

<a id="canonical-5830c5a38523a99d2b97e7abe5315711c453403818ad6e89573130b1b228cbc5"></a>

## allowed_proc_mounts property — psp_spec / c699886c4d9c / 7

Type: `["list", "string"]`. Optional.

Allowed list of proc mounts, empty list allows default proc mounts.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(8),
}
```

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

<a id="canonical-262d383bbd36bd176bf9eff061edd2fccaacdb87943cc89f856619bbe802d38c"></a>

<a id="canonical-206d0583bceeb799960fa3834486b71f14e5140370087b033e145a1755bdb85c"></a>

## allowed_unsafe_sysctls property — psp_spec / c699886c4d9c / 8

Type: `["list", "string"]`. Optional.

Allowed list of unsafe sysctls, empty list allows none. Supports prefix reg-ex.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-a8a7fa76042d4baabc2ccecd08401f204f3d9e8d2b03850bf960ec720fa42cf3"></a>

<a id="canonical-5d2345740d6104a3cc82a15fb8b6c4a524722badfe63a4836f4b179d90e84798"></a>

## default_allow_privilege_escalation property — psp_spec / c699886c4d9c / 9

Type: `"bool"`. Optional.

Pod has permission for privilege escalation by default.

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

- [default_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-386606cef0acb0cdf651ecba8556997d19b4471a9bf32297eb68717d2239fb8d): complete subsection reference.

- [drop_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-071b501f2c33a0c4a6ad37dfc84cad19f4a66f04f373e525c87f43015de5429e): complete subsection reference.

<a id="canonical-e3e5b500f01cbe2611c33afa516e57a18313885a6817892d12328b88d6ef40ee"></a>

<a id="canonical-21341c5938e3dfa9d36386828ca0474ae86140b30ce9db1cd3d404ef6ce166db"></a>

## forbidden_sysctls property — psp_spec / c699886c4d9c / 10

Type: `["list", "string"]`. Optional.

Forbidden list of sysctls, empty list forbids none. Supports prefix reg-ex.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [fs_group_strategy_options](resources--k8s_pod_security_policy--reference--group-001.md#canonical-a24eb4e73e589188c96d3062ba8d335ca6645573c52e0a00e1b848804047b931): complete subsection reference.

<a id="canonical-7f7891711edfdd3330042dbb337712f07c95934c8ed9c8f405c9907d517b280e"></a>

<a id="canonical-8add88b7c1d92b979ff1baa4e01f384e37c444610bcc00e3849c0a5e2cfaf1e7"></a>

## host_ipc property — psp_spec / c699886c4d9c / 11

Type: `"bool"`. Optional.

Host IPC determines if the policy allows the use of host IPC in the pod spec.

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

<a id="canonical-49f66fa2e1bd613c80671c509d21e23e548ad285241beeaff376135a7cb0440f"></a>

<a id="canonical-ba7276721e148f7805baa4d0d2eeea75b8ce2559261f4020924a2e42a6dc1b34"></a>

## host_network property — psp_spec / c699886c4d9c / 12

Type: `"bool"`. Optional.

Host Network determines if the policy allows the use of host network in the pod spec.

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

<a id="canonical-4c4d8ec1fe5b7b1a997df77afa6236e6c3ae9a1e6d6fcac9652f49b7cfa03928"></a>

<a id="canonical-7620ea82bd060e1c4dbd133740a5c5e0b3e24c20bbb5576037eed163c253df98"></a>

## host_pid property — psp_spec / c699886c4d9c / 13

Type: `"bool"`. Optional.

Host PID determines if the policy allows the use of host PID in the pod spec.

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

<a id="canonical-f48a697341e8515e0334ab77bcbf357f91562c5a071639858c62d53e818ec6aa"></a>

<a id="canonical-18ac6fb27fe2d14c9d2f35251e2bf0a49ae644dc2cc091abc45b9f2604cc9eaa"></a>

## host_port_ranges property — psp_spec / c699886c4d9c / 14

Type: `"string"`. Optional.

Host port ranges determines which ports ranges are allowed to be exposed.

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
    "ves.io.schema.rules.string.port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.port_range_list": "true"
  }
}
```

- [no_allowed_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-328031be3fafd76341385adc61cb72207c185a43b9c5d97373e5f9c53dc91e68): complete subsection reference.

- [no_default_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-87c6151a1a875658a2c65fdcf49e66d68f30e3fc896a58e71d8bd15253628090): complete subsection reference.

- [no_drop_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-fec4370e0f3375ceb77c9eb03c31cc19f003b2a6480192a39079a1e97facd328): complete subsection reference.

- [no_fs_groups](resources--k8s_pod_security_policy--reference--group-001.md#canonical-f3fd466d4e281a9d87410df86e15905861bf90b8942548bbc1ebb9e2ab9d0a2b): complete subsection reference.

- [no_run_as_group](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1f45bb594ac919c49dbd0674083d6e786f2f0dc1a4531210c7bcb6f4d6fc4ae4): complete subsection reference.

- [no_run_as_user](resources--k8s_pod_security_policy--reference--group-001.md#canonical-cd76dc5cf40c5133b2e093886e9da0b03962dfee921f37de0d2a3e01cf0a1f77): complete subsection reference.

- [no_runtime_class](resources--k8s_pod_security_policy--reference--group-001.md#canonical-71a87ca7e8529cef6805c1dc2175d53371e01370ee4f7a0d358d045fdb11a9cf): complete subsection reference.

- [no_se_linux_options](resources--k8s_pod_security_policy--reference--group-001.md#canonical-dcdd91741b12485fddca17d5dfae9dfb70994ec497bd3bb870153574507086e4): complete subsection reference.

- [no_supplemental_groups](resources--k8s_pod_security_policy--reference--group-001.md#canonical-a980446abbcd8237876898a3860659ca0f7e968d4a0b6f4957654b771e3fc358): complete subsection reference.

<a id="canonical-88ad95491d8b28a335e197a48d96a6a81c63be8e77cc6c0b9689feea4d75a775"></a>

<a id="canonical-0398e10fdea8d3065c6fbe2def9e45291217d6521cf1412023bb3aa54ce9a4c3"></a>

## privileged property — psp_spec / c699886c4d9c / 15

Type: `"bool"`. Optional.

Privileged determines if a pod can request to be run as privileged.

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

<a id="canonical-99fab6f615dda13b324692dfcd38d8419896902344982289b54fdca3048bad46"></a>

<a id="canonical-2bbdfad9ccc5d1dad7aeb7bf5a66b039b96ca88720837f33ec955165f9b86eda"></a>

## read_only_root_filesystem property — psp_spec / c699886c4d9c / 16

Type: `"bool"`. Optional.

Containers can only run with read only root filesystem.

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

- [run_as_group](resources--k8s_pod_security_policy--reference--group-001.md#canonical-33b4cae5a31cb4163617062dd933f8b84308df99e523ce0b1909d2b61f7ee6be): complete subsection reference.

- [run_as_user](resources--k8s_pod_security_policy--reference--group-001.md#canonical-fa37acdeafc119714f22e986c8cf4f88ee9c29f1419b471b6389f1f80a9adb29): complete subsection reference.

- [supplemental_groups](resources--k8s_pod_security_policy--reference--group-001.md#canonical-e6e3fdeeafd8d2d4abb440c28b4f3296b2b2e3c745165a4b44005ded92e53f45): complete subsection reference.

<a id="canonical-58fdf20e5e672e1772537bf6624b60bb129487f19ea40315c39c461d3d35c309"></a>

<a id="canonical-cb5a1a0df68bc8a00d0f2f5d52c47017cbaf64ce32db05495bbd6f2b45c15793"></a>

## volumes property — psp_spec / c699886c4d9c / 17

Type: `["list", "string"]`. Optional.

Allow List of volume plugins. Empty no volumes are allowed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(8),
}
```

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-e6f4404b31cfb86535f62ab49cde4bbc36459ba17a75069a657c7bbafc99bbf6"></a>

## Next pages — psp_spec / c699886c4d9c / 18

- [psp_spec.allowed_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-84b9a80fdd498440ac1ae1ff0cdb8f452ad978f195755e6bfce4996fbcff44d4)
- [psp_spec.allowed_host_paths](resources--k8s_pod_security_policy--reference--group-001.md#canonical-bf013790bfc3a2afd2cedf8e7b7d670f43e0f55a7503a1a2b6189836d26b47e4)
- [psp_spec.default_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-386606cef0acb0cdf651ecba8556997d19b4471a9bf32297eb68717d2239fb8d)
- [psp_spec.drop_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-071b501f2c33a0c4a6ad37dfc84cad19f4a66f04f373e525c87f43015de5429e)
- [psp_spec.fs_group_strategy_options](resources--k8s_pod_security_policy--reference--group-001.md#canonical-a24eb4e73e589188c96d3062ba8d335ca6645573c52e0a00e1b848804047b931)
- [psp_spec.no_allowed_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-328031be3fafd76341385adc61cb72207c185a43b9c5d97373e5f9c53dc91e68)
- [psp_spec.no_default_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-87c6151a1a875658a2c65fdcf49e66d68f30e3fc896a58e71d8bd15253628090)
- [psp_spec.no_drop_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-fec4370e0f3375ceb77c9eb03c31cc19f003b2a6480192a39079a1e97facd328)
- [psp_spec.no_fs_groups](resources--k8s_pod_security_policy--reference--group-001.md#canonical-f3fd466d4e281a9d87410df86e15905861bf90b8942548bbc1ebb9e2ab9d0a2b)
- [psp_spec.no_run_as_group](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1f45bb594ac919c49dbd0674083d6e786f2f0dc1a4531210c7bcb6f4d6fc4ae4)
- [psp_spec.no_run_as_user](resources--k8s_pod_security_policy--reference--group-001.md#canonical-cd76dc5cf40c5133b2e093886e9da0b03962dfee921f37de0d2a3e01cf0a1f77)
- [psp_spec.no_runtime_class](resources--k8s_pod_security_policy--reference--group-001.md#canonical-71a87ca7e8529cef6805c1dc2175d53371e01370ee4f7a0d358d045fdb11a9cf)
- [psp_spec.no_se_linux_options](resources--k8s_pod_security_policy--reference--group-001.md#canonical-dcdd91741b12485fddca17d5dfae9dfb70994ec497bd3bb870153574507086e4)
- [psp_spec.no_supplemental_groups](resources--k8s_pod_security_policy--reference--group-001.md#canonical-a980446abbcd8237876898a3860659ca0f7e968d4a0b6f4957654b771e3fc358)
- [psp_spec.run_as_group](resources--k8s_pod_security_policy--reference--group-001.md#canonical-33b4cae5a31cb4163617062dd933f8b84308df99e523ce0b1909d2b61f7ee6be)
- [psp_spec.run_as_user](resources--k8s_pod_security_policy--reference--group-001.md#canonical-fa37acdeafc119714f22e986c8cf4f88ee9c29f1419b471b6389f1f80a9adb29)
- [psp_spec.supplemental_groups](resources--k8s_pod_security_policy--reference--group-001.md#canonical-e6e3fdeeafd8d2d4abb440c28b4f3296b2b2e3c745165a4b44005ded92e53f45)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3ca2e681586afce0edce8cbc8836a556f7345a01b126e0b7a741a4a76865647f)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)

<a id="canonical-84b9a80fdd498440ac1ae1ff0cdb8f452ad978f195755e6bfce4996fbcff44d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c046997ab15619052f856a679c05a469b11bdbc6f6cde26012697f4afceea4d2"></a>

## psp_spec.allowed_capabilities — psp_spec.allowed_capabilities / f6e7121d9907 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3ca2e681586afce0edce8cbc8836a556f7345a01b126e0b7a741a4a76865647f)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- psp_spec.allowed_capabilities

<a id="canonical-bdfd1fce5b9ae5da698804a90985a5ca033048542e55159f09c31950981d032d"></a>

Type: `"object"`. single nested block, Optional.

List of capabilities that docker container has.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("capabilities")}
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
allowed_capabilities {
  # Configure direct properties listed below.
}
```

<a id="canonical-da5d45acd5adecd70c4d3a303ddc880c062b1ce78555acbcba7eea7343c8a7ce"></a>

## Direct properties — psp_spec.allowed_capabilities / f6e7121d9907 / 3

<a id="canonical-e67cdb0f96b06b8987a22a047e70b4a6f9809a152fb580a81ed2b42f210abbd7"></a>

<a id="canonical-4f57633a3eef8a5324874ea88e421466f57305ce72a2908d36e4803b65e93c93"></a>

## capabilities property — psp_spec.allowed_capabilities / f6e7121d9907 / 4

Type: `["list", "string"]`. Optional.

List of capabilities that docker container has.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-f4c294144687a698600f0d96ed40ff4f3c0d026d20ece75ac5ee596e34a93a80"></a>

## Next pages — psp_spec.allowed_capabilities / f6e7121d9907 / 5

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)

<a id="canonical-bf013790bfc3a2afd2cedf8e7b7d670f43e0f55a7503a1a2b6189836d26b47e4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5bcd247b833a7f78c5fbe946b69b4dce36bf536feea74b103563675cd0144e4d"></a>

## psp_spec.allowed_host_paths — psp_spec.allowed_host_paths / 7874f43bc377 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3ca2e681586afce0edce8cbc8836a556f7345a01b126e0b7a741a4a76865647f)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- psp_spec.allowed_host_paths

<a id="canonical-da3395337800e9f270a2754b519eacc6422cb27023a72798e6f2ac0ba5084074"></a>

Type: `"object"`. list nested block, Optional.

Restrict list of host paths, default all host paths are allowed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("path_prefix")}
```

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

Terraform syntax:

```terraform
allowed_host_paths {
  # Configure direct properties listed below.
}
```

<a id="canonical-93d47f2be2222609a44011980d5a69cfaa29b73d58b52d5c8695fa634e78236c"></a>

## Direct properties — psp_spec.allowed_host_paths / 7874f43bc377 / 3

<a id="canonical-74e4bcc01c9903019da5fae7e7732c45a3dcf52fcfe43646fc25d3b962b729c4"></a>

<a id="canonical-240a66b9c111175b59014a0e6906617249fafd787e21a2a5406bbb65a6a10b74"></a>

## path_prefix property — psp_spec.allowed_host_paths / 7874f43bc377 / 4

Type: `"string"`. Optional.

Host path prefix is the path prefix that the host volume must match. It does not support \*.

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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-447bc81c21d683a2b18249f8114f816c02e471d38f718ce88af493dc04b11beb"></a>

<a id="canonical-dfd67a17d28a07aae16baa8129b964230fcab7b773e40a25728b70a0550fd5db"></a>

## read_only property — psp_spec.allowed_host_paths / 7874f43bc377 / 5

Type: `"bool"`. Optional.

Volume will be allowed to mount read only.

Upstream description:

This volume will be allowed to mount read only.

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

<a id="canonical-7dae7637b5fde580ac1973bcb84e3cbc7a7875b4878c01801bd7103453179aa9"></a>

## Next pages — psp_spec.allowed_host_paths / 7874f43bc377 / 6

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)

<a id="canonical-386606cef0acb0cdf651ecba8556997d19b4471a9bf32297eb68717d2239fb8d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-623f2785fe4821a3ff09da5de40545f60e2a27cae94ab97de3b8cf1ca8de9f91"></a>

## psp_spec.default_capabilities — psp_spec.default_capabilities / 01a523e172c0 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3ca2e681586afce0edce8cbc8836a556f7345a01b126e0b7a741a4a76865647f)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- psp_spec.default_capabilities

<a id="canonical-9f60462a9f4ac3764b992806ffea75a93e60e2560762b32c45b007bcf8c837c8"></a>

Type: `"object"`. single nested block, Optional.

List of capabilities that docker container has.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("capabilities")}
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
default_capabilities {
  # Configure direct properties listed below.
}
```

<a id="canonical-a62c183c9f61e464d1f7980cfd78993223e3cf00b343f6e6e58dbd0af63083b4"></a>

## Direct properties — psp_spec.default_capabilities / 01a523e172c0 / 3

<a id="canonical-b89b98c824e8c1b3581b004c9b985222d0a52db47033bc2a50bcc23567d5041d"></a>

<a id="canonical-320770a5748f0692b7f80a05273bfe87e74675a71ebede8154e15431877115d9"></a>

## capabilities property — psp_spec.default_capabilities / 01a523e172c0 / 4

Type: `["list", "string"]`. Optional.

List of capabilities that docker container has.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-6d4570cd6ee36b21ee350a742afb445ba5150f7fd2a9e1b89363a5ef896aee86"></a>

## Next pages — psp_spec.default_capabilities / 01a523e172c0 / 5

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)

<a id="canonical-071b501f2c33a0c4a6ad37dfc84cad19f4a66f04f373e525c87f43015de5429e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c96e6f67319a1073044a3ce1731d79ad0fd5d2aaeff1d8aa5ab92fbfe331e85"></a>

## psp_spec.drop_capabilities — psp_spec.drop_capabilities / 85922eb6da50 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3ca2e681586afce0edce8cbc8836a556f7345a01b126e0b7a741a4a76865647f)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- psp_spec.drop_capabilities

<a id="canonical-8f5242195018aabb5f407307391ff58d3286fb913073b0dbe7a0c95957ee402c"></a>

Type: `"object"`. single nested block, Optional.

List of capabilities that docker container has.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("capabilities")}
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
drop_capabilities {
  # Configure direct properties listed below.
}
```

<a id="canonical-190fa95ca504e1891e729de494c48d25cdabea4ce127b71ce6b4545d399e8701"></a>

## Direct properties — psp_spec.drop_capabilities / 85922eb6da50 / 3

<a id="canonical-a1f92fc17d0fe5169fdc963b49be47b54c457cddc6296a6b2560367408032476"></a>

<a id="canonical-64452cb06d0e66e1a46f449d89909592bf45233177b2740724d37dbee014b4a0"></a>

## capabilities property — psp_spec.drop_capabilities / 85922eb6da50 / 4

Type: `["list", "string"]`. Optional.

List of capabilities that docker container has.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-4edcdaa73ff9404ab5ef379c5e5724a084a1e62fdb36554fd0aa12b79e7c045a"></a>

## Next pages — psp_spec.drop_capabilities / 85922eb6da50 / 5

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)

<a id="canonical-a24eb4e73e589188c96d3062ba8d335ca6645573c52e0a00e1b848804047b931"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-137f3973981f3e41be3e0fb845e1803823275d73e5c6a0e66389dd83e45e4afd"></a>

## psp_spec.fs_group_strategy_options — psp_spec.fs_group_strategy_options / 279eb222fd34 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3ca2e681586afce0edce8cbc8836a556f7345a01b126e0b7a741a4a76865647f)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- psp_spec.fs_group_strategy_options

<a id="canonical-445e6cbbe92febaaa68dc07fed386831f466f97f7a385bdd9713add9a6b1fb66"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for fs group strategy options.

Upstream description:

ID ranges and rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rule")}
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
fs_group_strategy_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-055077845e866b59a80bd75a3dd6b66ff14706f5e9e295e338712bed551e9fb8"></a>

## Direct properties — psp_spec.fs_group_strategy_options / 279eb222fd34 / 3

- [id_ranges](resources--k8s_pod_security_policy--reference--group-001.md#canonical-67aabd7d9d9f4f56595cda603f2d2eefffdadbada48ce50aceb00ff06c47f9c8): complete subsection reference.

<a id="canonical-bd3fe1a8f19b4a27df99f616e44805a564ade7f9080482c4539b3053f1e381bf"></a>

<a id="canonical-fb9f6d7cf52fde080a1733744b9b903f68ce53e14e766e871a2fb98428ae7546"></a>

## rule property — psp_spec.fs_group_strategy_options / 279eb222fd34 / 4

Type: `"string"`. Optional.

Rule indicated how the FS group ID range is used.

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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-20589fc20c5b827165717767d3715387b0e949a56e854c1f8a86b6406d1f0c3b"></a>

## Next pages — psp_spec.fs_group_strategy_options / 279eb222fd34 / 5

- [psp_spec.fs_group_strategy_options.id_ranges](resources--k8s_pod_security_policy--reference--group-001.md#canonical-67aabd7d9d9f4f56595cda603f2d2eefffdadbada48ce50aceb00ff06c47f9c8)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)

<a id="canonical-67aabd7d9d9f4f56595cda603f2d2eefffdadbada48ce50aceb00ff06c47f9c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-51eb2763a88b31fa0e4eae8c057f113bb122ef7c783c399827215472e9231757"></a>

## psp_spec.fs_group_strategy_options.id_ranges — psp_spec.fs_group_strategy_options.id_ranges / 5f4ece81a6ac / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3ca2e681586afce0edce8cbc8836a556f7345a01b126e0b7a741a4a76865647f)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- [psp_spec.fs_group_strategy_options](resources--k8s_pod_security_policy--reference--group-001.md#canonical-a24eb4e73e589188c96d3062ba8d335ca6645573c52e0a00e1b848804047b931)
- psp_spec.fs_group_strategy_options.id_ranges

<a id="canonical-eb06e44b2a61a48b9a5e2c33dbf975ecf7ed8eeb58da1f04038139d4d01ee29e"></a>

Type: `"object"`. list nested block, Optional.

ID Ranges. List of range of ID(s)

Upstream description:

List of range of ID(s)

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("max_id",
    "min_id")}
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

Terraform syntax:

```terraform
id_ranges {
  # Configure direct properties listed below.
}
```

<a id="canonical-cec88630ba006a7a14530ab4d42510bc845fc0c8b4c30e0527d00605587da30b"></a>

## Direct properties — psp_spec.fs_group_strategy_options.id_ranges / 5f4ece81a6ac / 3

<a id="canonical-7d20c6bdd927e8b364f53055003513aad7beb1657557a1030e61caccfce4aea4"></a>

<a id="canonical-3488b772fd3b2b359e7755f3b355bce0428b3e78fe572293b3335b9d8424ff3d"></a>

## max_id property — psp_spec.fs_group_strategy_options.id_ranges / 5f4ece81a6ac / 4

Type: `"number"`. Optional.

Ending ID. Ending(maximum) ID for for ID range.

Upstream description:

Ending(maximum) ID for for ID range.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3ebb1ade79fd3d434514de424d7cb60344c20c1974c0578c0dc676dc36f145fd"></a>

<a id="canonical-499736fa47a3d4ab89b5786a2eb83d02466fe15af8c2e486f3fe77d5d1554826"></a>

## min_id property — psp_spec.fs_group_strategy_options.id_ranges / 5f4ece81a6ac / 5

Type: `"number"`. Optional.

Starting ID. Starting(minimum) ID for for ID range.

Upstream description:

Starting(minimum) ID for for ID range.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-736ba799b8b0ca53a988ac2f716b1918c49142416181886f029803b7cb61ec4d"></a>

## Next pages — psp_spec.fs_group_strategy_options.id_ranges / 5f4ece81a6ac / 6

- [psp_spec.fs_group_strategy_options](resources--k8s_pod_security_policy--reference--group-001.md#canonical-a24eb4e73e589188c96d3062ba8d335ca6645573c52e0a00e1b848804047b931)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)

<a id="canonical-328031be3fafd76341385adc61cb72207c185a43b9c5d97373e5f9c53dc91e68"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8c6bbbe7f51fff3eeae6dd5ab69270add4f7660fa85a12b091e23f82fd10ed6c"></a>

## psp_spec.no_allowed_capabilities — psp_spec.no_allowed_capabilities / 39a97de18b25 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3ca2e681586afce0edce8cbc8836a556f7345a01b126e0b7a741a4a76865647f)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- psp_spec.no_allowed_capabilities

<a id="canonical-ab77591fa945d9263d2347b60a5f3955845efe1ad10c4a37c9da3402934be3b8"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no allowed capabilities.

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
no_allowed_capabilities = {}
```

<a id="canonical-4d3c70d79ae8f83c403428913fa0a2c44d9c854b04ae4a1b19b991b3533c217a"></a>

## Direct properties — psp_spec.no_allowed_capabilities / 39a97de18b25 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-44a82e259e3da776b2a0e08d4e4d29f1b8ee9c7fd2f29e5769522a4785a34ce3"></a>

## Next pages — psp_spec.no_allowed_capabilities / 39a97de18b25 / 4

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)

<a id="canonical-87c6151a1a875658a2c65fdcf49e66d68f30e3fc896a58e71d8bd15253628090"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-00861753785e1ade0f3a49c6feaba14b22c6810227f11b999b25affdec376ce9"></a>

## psp_spec.no_default_capabilities — psp_spec.no_default_capabilities / 1e2c94ef22ad / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3ca2e681586afce0edce8cbc8836a556f7345a01b126e0b7a741a4a76865647f)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- psp_spec.no_default_capabilities

<a id="canonical-2d2789d93284091ca488d2d9d6cbb238a703b9a9d2942918deebd9a038f533ef"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no default capabilities.

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
no_default_capabilities = {}
```

<a id="canonical-294f3ff7a2311f66258c57ee54eeda4e741d0ddaa944d9ba02292597a3c72f16"></a>

## Direct properties — psp_spec.no_default_capabilities / 1e2c94ef22ad / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d75890da80b70f2502d86a73685a1a1d1ad412e4b6c58b87f17167bd4a348f3c"></a>

## Next pages — psp_spec.no_default_capabilities / 1e2c94ef22ad / 4

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)

<a id="canonical-fec4370e0f3375ceb77c9eb03c31cc19f003b2a6480192a39079a1e97facd328"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2f3b084fa2bf9c7ab7e00dae319e5c18bbd08e035701c321b938a4236de62b60"></a>

## psp_spec.no_drop_capabilities — psp_spec.no_drop_capabilities / b87dd5faa657 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3ca2e681586afce0edce8cbc8836a556f7345a01b126e0b7a741a4a76865647f)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- psp_spec.no_drop_capabilities

<a id="canonical-5b71f47c4eee6489b097b1ef9ff13c355bcd9f0ae19e0368a5bbe6a21ef994c1"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no drop capabilities.

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
no_drop_capabilities = {}
```

<a id="canonical-f7f19dc461fdd479cbb9c92eb79424a715a15ad4a5e0ec6676dba035b72a9c27"></a>

## Direct properties — psp_spec.no_drop_capabilities / b87dd5faa657 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ed9a1565e47f72bfa0927f5fdabb51ed59b3cc4658ace70e8016d7536a5335dd"></a>

## Next pages — psp_spec.no_drop_capabilities / b87dd5faa657 / 4

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)

<a id="canonical-f3fd466d4e281a9d87410df86e15905861bf90b8942548bbc1ebb9e2ab9d0a2b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d4efe9002f0eb9f50f9ac435b143d8cbe7d6bfb4d09098e25575cb0dbe544d73"></a>

## psp_spec.no_fs_groups — psp_spec.no_fs_groups / 964a5dba3366 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3ca2e681586afce0edce8cbc8836a556f7345a01b126e0b7a741a4a76865647f)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- psp_spec.no_fs_groups

<a id="canonical-8f3f82cd3e686c13846052f7c653dd9156ff4114dd9b5d20dd251d8cf486984e"></a>

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
no_fs_groups = {}
```

<a id="canonical-63420c99edd7cc41b4bc136a406ca86091f92a026aaebd3e49c22995dcdad3a7"></a>

## Direct properties — psp_spec.no_fs_groups / 964a5dba3366 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7ae099b9a2cb8cc8c18434ab86dd4677c73d632a4a948f280ca76dd697c60c11"></a>

## Next pages — psp_spec.no_fs_groups / 964a5dba3366 / 4

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)

<a id="canonical-1f45bb594ac919c49dbd0674083d6e786f2f0dc1a4531210c7bcb6f4d6fc4ae4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9a12e90d37cbc5a5d7cbab613356df27a2c9a27e1ecba89934bbf00dc662b96e"></a>

## psp_spec.no_run_as_group — psp_spec.no_run_as_group / 5dcc73d609ba / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3ca2e681586afce0edce8cbc8836a556f7345a01b126e0b7a741a4a76865647f)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- psp_spec.no_run_as_group

<a id="canonical-88b3e4fa92b23d54b010a36d514aedbaa04d80db1db3cc532d8c113d6aa18cb6"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no run as group.

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
no_run_as_group = {}
```

<a id="canonical-9e6289675a59f24c358f222078c880b69b4e898d2c92a4c0b3a40c7e1240ef9b"></a>

## Direct properties — psp_spec.no_run_as_group / 5dcc73d609ba / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a05ca9897594dacbf0dce1887704fb0844f062d1b2136f640330da7e0ea884a7"></a>

## Next pages — psp_spec.no_run_as_group / 5dcc73d609ba / 4

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)

<a id="canonical-cd76dc5cf40c5133b2e093886e9da0b03962dfee921f37de0d2a3e01cf0a1f77"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-704cfad2ae8140cc317c16e5290d8fd07efdc33073ff2502e2c903bd07d5a0b7"></a>

## psp_spec.no_run_as_user — psp_spec.no_run_as_user / b7c664371348 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3ca2e681586afce0edce8cbc8836a556f7345a01b126e0b7a741a4a76865647f)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- psp_spec.no_run_as_user

<a id="canonical-64243e16e92de50c1266f91d9c7b066b854ef365b15801ef784e5a0f859bc68e"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no run as user.

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
no_run_as_user = {}
```

<a id="canonical-7a90c5b9876e9987c39b571961ced9e300e525c77385a8cbc8aa35bf68083d0f"></a>

## Direct properties — psp_spec.no_run_as_user / b7c664371348 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-92d98ee8059a9454a9c097c5771ea312c8c8fa954851bbd390fe94546224cf73"></a>

## Next pages — psp_spec.no_run_as_user / b7c664371348 / 4

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)

<a id="canonical-71a87ca7e8529cef6805c1dc2175d53371e01370ee4f7a0d358d045fdb11a9cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c91333e20ead7f2c9854d27666530349c1b4f72d08ad600c30c66792243acadc"></a>

## psp_spec.no_runtime_class — psp_spec.no_runtime_class / 03f5af05f3f2 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3ca2e681586afce0edce8cbc8836a556f7345a01b126e0b7a741a4a76865647f)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- psp_spec.no_runtime_class

<a id="canonical-2b7876d89f0282beabe4a2f154505c222d29a526060f0cc00e3f66f422419f0e"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for no runtime class.

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
no_runtime_class {}
```

<a id="canonical-319a75e59368e6cc719c4425031f0ed9fa04172ee868d2181f563783be8d32cd"></a>

## Direct properties — psp_spec.no_runtime_class / 03f5af05f3f2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-006e2193c0dafbab1969f3ce82d4472a2a784988e1106fd0a601a6737395f06c"></a>

## Next pages — psp_spec.no_runtime_class / 03f5af05f3f2 / 4

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)

<a id="canonical-dcdd91741b12485fddca17d5dfae9dfb70994ec497bd3bb870153574507086e4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-35adb7799ba1332cbcad0e5fe9b3f1ec9a97d3f0b4f5daeaf620a5d681f0d39a"></a>

## psp_spec.no_se_linux_options — psp_spec.no_se_linux_options / 8f22010f246e / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3ca2e681586afce0edce8cbc8836a556f7345a01b126e0b7a741a4a76865647f)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- psp_spec.no_se_linux_options

<a id="canonical-9f8f221400e983582445146bab6191f79c77e3a89922d0837595ee38ec274fcf"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for no se linux options.

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
no_se_linux_options {}
```

<a id="canonical-a56aa450e547519e3459288b74cee29d6b2acbd2638b92cbac952b34e0a0031f"></a>

## Direct properties — psp_spec.no_se_linux_options / 8f22010f246e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e9274b28bbe3e858453fa5fc49c24adc950c607b3409471841348eb09d85a595"></a>

## Next pages — psp_spec.no_se_linux_options / 8f22010f246e / 4

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)

<a id="canonical-a980446abbcd8237876898a3860659ca0f7e968d4a0b6f4957654b771e3fc358"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-557e633008a830d434a80c729bc2128f6898b6fa74c5e117dd973fd1f24ba165"></a>

## psp_spec.no_supplemental_groups — psp_spec.no_supplemental_groups / b29df29861d8 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3ca2e681586afce0edce8cbc8836a556f7345a01b126e0b7a741a4a76865647f)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- psp_spec.no_supplemental_groups

<a id="canonical-6b2bc15332b788837b448d8aac80c0c32386f32a9e838c54320e1f037b1780b9"></a>

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
no_supplemental_groups = {}
```

<a id="canonical-7de9e6701355321c143e706c8f862f7ea50876a40275212e8cd7963940692e5a"></a>

## Direct properties — psp_spec.no_supplemental_groups / b29df29861d8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-75e61de2fd570b9d53ea68b31916c9fd94ca48bc8de8459d758debf287232f17"></a>

## Next pages — psp_spec.no_supplemental_groups / b29df29861d8 / 4

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)

<a id="canonical-33b4cae5a31cb4163617062dd933f8b84308df99e523ce0b1909d2b61f7ee6be"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c00a30cdad9736941ab9a3786b1c6bf367f6971840fd9ff365b6bdd05486c3cf"></a>

## psp_spec.run_as_group — psp_spec.run_as_group / dd122d8f7369 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3ca2e681586afce0edce8cbc8836a556f7345a01b126e0b7a741a4a76865647f)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- psp_spec.run_as_group

<a id="canonical-32c6f232d49481b9d72b308b0d9124f3b5580a503f7c2e68cd3bd4252c2a54fd"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for run as group.

Upstream description:

ID ranges and rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rule")}
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
run_as_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-b322f085c289d57261b3f525a08f5fe3ceaea53ac0deccf8c99f8baeed12c50b"></a>

## Direct properties — psp_spec.run_as_group / dd122d8f7369 / 3

- [id_ranges](resources--k8s_pod_security_policy--reference--group-001.md#canonical-bfbfc651c26895ba432a1c068e5055deae8def70d5e57ff9d7fe53b4f28238dc): complete subsection reference.

<a id="canonical-d8bd241b7f20a8245e7e6062ff89168b4a52b73b37eed4aca590835a1212d9c9"></a>

<a id="canonical-fe025ad15295028e312f179f3b872740c7999901393dc8e122453847385bb67d"></a>

## rule property — psp_spec.run_as_group / dd122d8f7369 / 4

Type: `"string"`. Optional.

Rule indicated how the FS group ID range is used.

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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-755fb0c666931cab45900d6ed8426e4d3025dc4be192f7fd4e9c64e4ee35e9b5"></a>

## Next pages — psp_spec.run_as_group / dd122d8f7369 / 5

- [psp_spec.run_as_group.id_ranges](resources--k8s_pod_security_policy--reference--group-001.md#canonical-bfbfc651c26895ba432a1c068e5055deae8def70d5e57ff9d7fe53b4f28238dc)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)

<a id="canonical-bfbfc651c26895ba432a1c068e5055deae8def70d5e57ff9d7fe53b4f28238dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e9841956ebe039eeadb698eab56ffd6696e3fa00bc0a11fe61b3b2b62002e940"></a>

## psp_spec.run_as_group.id_ranges — psp_spec.run_as_group.id_ranges / 42aa7427a96f / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3ca2e681586afce0edce8cbc8836a556f7345a01b126e0b7a741a4a76865647f)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- [psp_spec.run_as_group](resources--k8s_pod_security_policy--reference--group-001.md#canonical-33b4cae5a31cb4163617062dd933f8b84308df99e523ce0b1909d2b61f7ee6be)
- psp_spec.run_as_group.id_ranges

<a id="canonical-dffcb361d095ae0b308b6da8de38d914b07311ebebc56ae9dca4a6cf955c165e"></a>

Type: `"object"`. list nested block, Optional.

ID Ranges. List of range of ID(s)

Upstream description:

List of range of ID(s)

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("max_id",
    "min_id")}
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

Terraform syntax:

```terraform
id_ranges {
  # Configure direct properties listed below.
}
```

<a id="canonical-733f0d1935fc8a1b465baabbf00f5c9679e37e409d8ba13dfdc2bd37708e114e"></a>

## Direct properties — psp_spec.run_as_group.id_ranges / 42aa7427a96f / 3

<a id="canonical-4ef6a4fa19c1d65095153e120b385b27cf2f0866d89f0949faf7f64c2cdf70be"></a>

<a id="canonical-1cbe8b0f7f4e53a4b78c3fe58649b2d84177a7d7182e037cf649df5b602f9f32"></a>

## max_id property — psp_spec.run_as_group.id_ranges / 42aa7427a96f / 4

Type: `"number"`. Optional.

Ending ID. Ending(maximum) ID for for ID range.

Upstream description:

Ending(maximum) ID for for ID range.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-41ae3802bbc9270f3fefd2ad6e1c8f61ad65f92af8e95bd695442cab807bf81d"></a>

<a id="canonical-80ce4bc737b3d5a1b78719134e71b9a152f15aca6496e1483b7d16052aab5128"></a>

## min_id property — psp_spec.run_as_group.id_ranges / 42aa7427a96f / 5

Type: `"number"`. Optional.

Starting ID. Starting(minimum) ID for for ID range.

Upstream description:

Starting(minimum) ID for for ID range.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-87653e4b759dd673cbb47150e6e65a7006e2df5a268c6b8c29730b6b8e55a7d2"></a>

## Next pages — psp_spec.run_as_group.id_ranges / 42aa7427a96f / 6

- [psp_spec.run_as_group](resources--k8s_pod_security_policy--reference--group-001.md#canonical-33b4cae5a31cb4163617062dd933f8b84308df99e523ce0b1909d2b61f7ee6be)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)

<a id="canonical-fa37acdeafc119714f22e986c8cf4f88ee9c29f1419b471b6389f1f80a9adb29"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2d88df6dd14d490ae31306dc48d427393ab9f851e0be600deb06ede7a4d73cf3"></a>

## psp_spec.run_as_user — psp_spec.run_as_user / 97061b390f76 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3ca2e681586afce0edce8cbc8836a556f7345a01b126e0b7a741a4a76865647f)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- psp_spec.run_as_user

<a id="canonical-78c630ad662e9c638ac1143363898842a1841f171c3807fc745ad12eabb17c49"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for run as user.

Upstream description:

ID ranges and rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rule")}
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
run_as_user {
  # Configure direct properties listed below.
}
```

<a id="canonical-bf7b25dbd411a46e75a45474fe5652ec820cc2463ee7e8c433f4bef8156fea3d"></a>

## Direct properties — psp_spec.run_as_user / 97061b390f76 / 3

- [id_ranges](resources--k8s_pod_security_policy--reference--group-001.md#canonical-6e1ee3e439de7cc6a5e069f13c5a6a06d23cd0797c0807600082719aa82c4c59): complete subsection reference.

<a id="canonical-983582a56b8cb35b700cb8d226a3b9811fbdb57d36569666188e1e11b32290be"></a>

<a id="canonical-6ad9d6e5b96c802e6cfe3b969ca2663b4b0be5a8898967c6026e739d2bf68c07"></a>

## rule property — psp_spec.run_as_user / 97061b390f76 / 4

Type: `"string"`. Optional.

Rule indicated how the FS group ID range is used.

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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-8dcc04e79935f01ab195946ed02ec5d166594089fda972a2ba06804bf0a3d102"></a>

## Next pages — psp_spec.run_as_user / 97061b390f76 / 5

- [psp_spec.run_as_user.id_ranges](resources--k8s_pod_security_policy--reference--group-001.md#canonical-6e1ee3e439de7cc6a5e069f13c5a6a06d23cd0797c0807600082719aa82c4c59)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)

<a id="canonical-6e1ee3e439de7cc6a5e069f13c5a6a06d23cd0797c0807600082719aa82c4c59"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-97586462a0d327a8d0a94b2aec0186f9969c39fbeb466cf0eb1e04da1a48a619"></a>

## psp_spec.run_as_user.id_ranges — psp_spec.run_as_user.id_ranges / 1e0eadc77ea1 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3ca2e681586afce0edce8cbc8836a556f7345a01b126e0b7a741a4a76865647f)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- [psp_spec.run_as_user](resources--k8s_pod_security_policy--reference--group-001.md#canonical-fa37acdeafc119714f22e986c8cf4f88ee9c29f1419b471b6389f1f80a9adb29)
- psp_spec.run_as_user.id_ranges

<a id="canonical-d6866f39b08fbd481d3752e120fc606a1f8e6795ddd4cefc230cb92a824e4c98"></a>

Type: `"object"`. list nested block, Optional.

ID Ranges. List of range of ID(s)

Upstream description:

List of range of ID(s)

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("max_id",
    "min_id")}
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

Terraform syntax:

```terraform
id_ranges {
  # Configure direct properties listed below.
}
```

<a id="canonical-7ad2721e7813b14d26293f0f700d98c00f87d63a52e7c0ae68b0a82bc31088b3"></a>

## Direct properties — psp_spec.run_as_user.id_ranges / 1e0eadc77ea1 / 3

<a id="canonical-629dc7e22352ddd9c2661ff63baf49636d32d056add40a754aa03e4f2503b6c5"></a>

<a id="canonical-6d25b8d0fb4559ee6faaa33afd11188d738c42d8493ff65b028dfa7ad3a556dd"></a>

## max_id property — psp_spec.run_as_user.id_ranges / 1e0eadc77ea1 / 4

Type: `"number"`. Optional.

Ending ID. Ending(maximum) ID for for ID range.

Upstream description:

Ending(maximum) ID for for ID range.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-7632f6a2d975e97046849cb48315e3c91c7ba4db14df7d1fe06bde4c6d3c117d"></a>

<a id="canonical-ef6059e0a9b47e95fec8a572a842bf01e1612a4c70a10b961c0168a6ac05fbbd"></a>

## min_id property — psp_spec.run_as_user.id_ranges / 1e0eadc77ea1 / 5

Type: `"number"`. Optional.

Starting ID. Starting(minimum) ID for for ID range.

Upstream description:

Starting(minimum) ID for for ID range.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-876a7bb705633459e9c17d1f17957d2bc166047a3eb449f08fe1c46b0f72b4b3"></a>

## Next pages — psp_spec.run_as_user.id_ranges / 1e0eadc77ea1 / 6

- [psp_spec.run_as_user](resources--k8s_pod_security_policy--reference--group-001.md#canonical-fa37acdeafc119714f22e986c8cf4f88ee9c29f1419b471b6389f1f80a9adb29)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)

<a id="canonical-e6e3fdeeafd8d2d4abb440c28b4f3296b2b2e3c745165a4b44005ded92e53f45"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f4a7dadca7f69e20d4062d6d57e8faf6ce1ca200d20286c52bb0c8bdefdc3dbe"></a>

## psp_spec.supplemental_groups — psp_spec.supplemental_groups / 81cbf25272b5 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3ca2e681586afce0edce8cbc8836a556f7345a01b126e0b7a741a4a76865647f)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- psp_spec.supplemental_groups

<a id="canonical-66c2a86f8da53f24f3676bb2bead442cee666e1bf0b6f397818e888ac5fd2843"></a>

Type: `"object"`. single nested block, Optional.

ID(User,Group,FSGroup) Strategy. ID ranges and rules.

Upstream description:

ID ranges and rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rule")}
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
supplemental_groups {
  # Configure direct properties listed below.
}
```

<a id="canonical-b22527cf171e8f8b0cd44eacb848aa03d72551b75c27ac99da8c56b3144af9aa"></a>

## Direct properties — psp_spec.supplemental_groups / 81cbf25272b5 / 3

- [id_ranges](resources--k8s_pod_security_policy--reference--group-001.md#canonical-e157d231398eaaaacc94d1191d2f3c8243ffce395da152cbb968c6190fb31b7e): complete subsection reference.

<a id="canonical-5be18dfbb8d1a8086686d2d502211236bfecefc2172823607b9d0c5738c6018d"></a>

<a id="canonical-2b246d6cb7eab25b9d0b8c69b58b908f38c81d9c2ba697cf00b7a543bf462195"></a>

## rule property — psp_spec.supplemental_groups / 81cbf25272b5 / 4

Type: `"string"`. Optional.

Rule indicated how the FS group ID range is used.

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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-afdf329904981f1408470d315ac33b2095b4f95c27499a32cf1b6b46af97c010"></a>

## Next pages — psp_spec.supplemental_groups / 81cbf25272b5 / 5

- [psp_spec.supplemental_groups.id_ranges](resources--k8s_pod_security_policy--reference--group-001.md#canonical-e157d231398eaaaacc94d1191d2f3c8243ffce395da152cbb968c6190fb31b7e)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)

<a id="canonical-e157d231398eaaaacc94d1191d2f3c8243ffce395da152cbb968c6190fb31b7e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab629fdc880a6899ba93678c8da39f1001870ea3db0e41b47cc9934f5b968fac"></a>

## psp_spec.supplemental_groups.id_ranges — psp_spec.supplemental_groups.id_ranges / 9ea62078ed38 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3ca2e681586afce0edce8cbc8836a556f7345a01b126e0b7a741a4a76865647f)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3550ce1e33410d4bfdcfed3ce773320067a4290df9b7d35a9277ddedfecc234e)
- [psp_spec.supplemental_groups](resources--k8s_pod_security_policy--reference--group-001.md#canonical-e6e3fdeeafd8d2d4abb440c28b4f3296b2b2e3c745165a4b44005ded92e53f45)
- psp_spec.supplemental_groups.id_ranges

<a id="canonical-38acbba07d6d6ee392ac928b5aa94a5a2bc257077e5b27b59e212e00cd332c51"></a>

Type: `"object"`. list nested block, Optional.

ID Ranges. List of range of ID(s)

Upstream description:

List of range of ID(s)

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("max_id",
    "min_id")}
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

Terraform syntax:

```terraform
id_ranges {
  # Configure direct properties listed below.
}
```

<a id="canonical-6680c3b6efd5e8a7a3969edbb61ab0d8749b2e3ee9b1b8829eccae3c575f5fb3"></a>

## Direct properties — psp_spec.supplemental_groups.id_ranges / 9ea62078ed38 / 3

<a id="canonical-fdadfcaa67ef914614e8de12341fd96c8477864bd7e2250fc72989acbda72279"></a>

<a id="canonical-e321fe50a5f76aa902ad4793158f0fa919b122004170e58060cc2bab0d78f600"></a>

## max_id property — psp_spec.supplemental_groups.id_ranges / 9ea62078ed38 / 4

Type: `"number"`. Optional.

Ending ID. Ending(maximum) ID for for ID range.

Upstream description:

Ending(maximum) ID for for ID range.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-d95276fb0a5bc35b35cf80178da0a23cebc35baf9734cc726df888657f830aea"></a>

<a id="canonical-3aa0e54a4b7dd9a9cb6cb9e5c218d80e8cc4a65898e9d0e660322a7730fca786"></a>

## min_id property — psp_spec.supplemental_groups.id_ranges / 9ea62078ed38 / 5

Type: `"number"`. Optional.

Starting ID. Starting(minimum) ID for for ID range.

Upstream description:

Starting(minimum) ID for for ID range.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-8241e3bfd71f0a4ba8b9de07847528e335a82fa09fe527858e4a042374ceb2fb"></a>

## Next pages — psp_spec.supplemental_groups.id_ranges / 9ea62078ed38 / 6

- [psp_spec.supplemental_groups](resources--k8s_pod_security_policy--reference--group-001.md#canonical-e6e3fdeeafd8d2d4abb440c28b4f3296b2b2e3c745165a4b44005ded92e53f45)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)

<a id="canonical-0d0bb5077a31feaba9ea04921fd608059998b40cae432fc7fddba2580ee3d459"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a5ee836120cd3a422dcd4472863bc0a241b5f0f509ef2ae370989de162813cbd"></a>

## timeouts — timeouts / c3f8bafb3feb / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3ca2e681586afce0edce8cbc8836a556f7345a01b126e0b7a741a4a76865647f)
- timeouts

<a id="canonical-09b1578bcadc482f5cb3d0d631a9e17250c8a839bf47af9df10907296c2e5aea"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-000af5d632be00c3aa8485674f53a439ba68f8cd54c0fa77f3c1fa0cc74f8444"></a>

## Direct properties — timeouts / c3f8bafb3feb / 3

<a id="canonical-2f141ffca4fe2a3afe67c772f7cf4e3b8a21e68576f9dfc92fc4bd517429bf8f"></a>

<a id="canonical-08973e57fa937812629f6f4b4b050d593735dbb5ee71ebd00cbe02c40f2d6bc1"></a>

## create property — timeouts / c3f8bafb3feb / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-7c990a5a0c7a9e6ca94b0cf1b69a726e291299e54dbf530382b58a9c7249a99b"></a>

<a id="canonical-b5b7ce33804a4764dade997293981efad81777fadcb80ff4468c5898f022bc7b"></a>

## delete property — timeouts / c3f8bafb3feb / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1a519a7d960bf83d846500ed715f662c6fbc3ab43356f06f4d5c44e5a4e2adf2"></a>

<a id="canonical-f67e1297b752c05dd49cbb1347b41f1651159eeec6c48fd2a80fb2e6db7f4e22"></a>

## read property — timeouts / c3f8bafb3feb / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-73556f12d92d4ee4cc4353392cf070603d837d9894fd9fcc19c2d563ee9faf53"></a>

<a id="canonical-008c353b684a59564cb3c515828af028dbc0da1cb640222974a4e504a6ab1a9f"></a>

## update property — timeouts / c3f8bafb3feb / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-39e88f406407e835f877d0c8d87169e3934fa319330dba0bccdd44ead735579a"></a>

## Next pages — timeouts / c3f8bafb3feb / 8

- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3ca2e681586afce0edce8cbc8836a556f7345a01b126e0b7a741a4a76865647f)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-d1cf0037a1c7ad187d001a831cd0b85ab39532f9e1564acd0d0a64fb31870e96)
