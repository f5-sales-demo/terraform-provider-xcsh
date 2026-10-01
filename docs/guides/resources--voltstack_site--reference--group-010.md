---
page_title: "xcsh_voltstack_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_voltstack_site reference."
---

# xcsh_voltstack_site reference

<a id="canonical-53eba2182b11666927bc93f82cde19cb0a750a642fe07bf152ab1fb3c175f894"></a>

## offline_survivability_mode.no_offline_survivability_mode — offline_survivability_mode.no_offline_survivability_mode / ed18e34797de / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [offline_survivability_mode](resources--voltstack_site--reference--group-009.md#canonical-186e14d0d96fbfdebbd08044ee03b5772cde1a87e273a1f99fcc0fe6163d5171)
- offline_survivability_mode.no_offline_survivability_mode

<a id="canonical-e119e32b45d982802ef65249e293fdef60b7eca4cc9a47a688f5912aeac772e9"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no offline survivability mode.

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
no_offline_survivability_mode = {}
```

<a id="canonical-0cb485d72436753e095cc2f6d01e42a1248364e253632b393ad6ec4dd1163726"></a>

## Direct properties — offline_survivability_mode.no_offline_survivability_mode / ed18e34797de / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c045566c502c40acadad2e830d87c80debc967107fec8473968164f2a917d6da"></a>

## Next pages — offline_survivability_mode.no_offline_survivability_mode / ed18e34797de / 4

- [offline_survivability_mode](resources--voltstack_site--reference--group-009.md#canonical-186e14d0d96fbfdebbd08044ee03b5772cde1a87e273a1f99fcc0fe6163d5171)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-acafd9e483552322116dcf73798aa1ebffd5871a6d89ddb562916c57e59f1683"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a8c078b624674c5bf03d28760c98e544dc2005dfa118d330332f53866e23569a"></a>

## os — os / e4100b504e84 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- os

<a id="canonical-7c322f543af6e05ddb2a33bdb88436b9b0dc8d5a0efdfa776f3048361e391981"></a>

Type: `"object"`. single nested block, Optional.

Select the F5XC Operating System Version for the site. By default, latest available OS Version will
be used. Refer to release notes to find required released OS versions.

Upstream description:

Select the F5XC Operating System Version for the site. By default, latest available OS Version will
be used. Refer to release notes to find required released OS versions.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_os_version",
    "operating_system_version")}
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
  "x-ves-oneof-field-operating_system_version_choice": "[\"default_os_version\",\"operating_system_version\"]"
}
```

Terraform syntax:

```terraform
os {
  # Configure direct properties listed below.
}
```

<a id="canonical-76cd7f1ba31d542e74465d3f51313dbeb82325d9187c1982fec0802216986ca0"></a>

## Direct properties — os / e4100b504e84 / 3

- [default_os_version](resources--voltstack_site--reference--group-010.md#canonical-103b3b0c5fc0161701811a01d57f225c06e5e3f97ea2683d4fbcd5555b555018): complete subsection reference.

<a id="canonical-7258c19268d428b3590fcf905b8d40efe34a42334906155aff49601d65c0fde9"></a>

<a id="canonical-0fdb0f0550a4abef15a866ac393627178f6d2ee69661b796c2a8db354e082051"></a>

## operating_system_version property — os / e4100b504e84 / 4

Type: `"string"`. Optional.

Exclusive with \[default\_os\_version\] Specify a OS version to be used e.g. 9.2024.6.

Upstream description:

Exclusive with \[default\_os\_version\] Specify a OS version to be used e.g. 9.2024.6.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-c79ca5e412ea6852c9d433fa7bfc44c00e92851fc59edf2f8c090502212ebd11"></a>

## Next pages — os / e4100b504e84 / 5

- [os.default_os_version](resources--voltstack_site--reference--group-010.md#canonical-103b3b0c5fc0161701811a01d57f225c06e5e3f97ea2683d4fbcd5555b555018)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-103b3b0c5fc0161701811a01d57f225c06e5e3f97ea2683d4fbcd5555b555018"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cc0abb15a41addf3380cf5d8e0419f532d0136b6e4c4e0f503220e75378f3013"></a>

## os.default_os_version — os.default_os_version / 5f67f45eb284 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [os](resources--voltstack_site--reference--group-010.md#canonical-acafd9e483552322116dcf73798aa1ebffd5871a6d89ddb562916c57e59f1683)
- os.default_os_version

<a id="canonical-14ea3863a4e37dfadb5a929fecdbdae6fb0fc04f9934349912fa619a5d1fefac"></a>

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
default_os_version = {}
```

<a id="canonical-815f51ed7f70708ed6af917910e6df63ed8f2f7a1aff3b7118bcc136a7d1856c"></a>

## Direct properties — os.default_os_version / 5f67f45eb284 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-584d27112fded4bf657b52602feff84faf3a053f7d9f9bf109e48c28aed4d6c5"></a>

## Next pages — os.default_os_version / 5f67f45eb284 / 4

- [os](resources--voltstack_site--reference--group-010.md#canonical-acafd9e483552322116dcf73798aa1ebffd5871a6d89ddb562916c57e59f1683)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-e288685eed32403a1b38a053cd958daafe67a3c468f5a4c75bbb371860409f6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-481ce2fe251bbf31bdd77e21f2f796b2c20d970a105d9e7874b8e16bc0667367"></a>

## sriov_interfaces — sriov_interfaces / 0491b0875e7d / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- sriov_interfaces

<a id="canonical-4360557592b1911e6d5f16e0ad08fee6ba2e8e2e9a844cfedc9c481bcee83120"></a>

Type: `"object"`. single nested block, Optional.

List of all custom SR-IOV interfaces configuration.

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
sriov_interfaces {
  # Configure direct properties listed below.
}
```

<a id="canonical-41bc7b4b1ef0ca3283840fedb654a699fc337cca6df5034eed2b1c32f554f867"></a>

## Direct properties — sriov_interfaces / 0491b0875e7d / 3

- [sriov_interface](resources--voltstack_site--reference--group-010.md#canonical-65991a6949693f19b45939f45b9e4988e34d11e6209f598b2bf617ff31fcb5af): complete subsection reference.

<a id="canonical-849079e3909a7c2c4cbb1907b52be87d8f09d52cd2d0fc1796a2dbcc6602833c"></a>

## Next pages — sriov_interfaces / 0491b0875e7d / 4

- [sriov_interfaces.sriov_interface](resources--voltstack_site--reference--group-010.md#canonical-65991a6949693f19b45939f45b9e4988e34d11e6209f598b2bf617ff31fcb5af)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-65991a6949693f19b45939f45b9e4988e34d11e6209f598b2bf617ff31fcb5af"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ced10d27d0f040540ac8d22c9ab9d27a2a8d4cb0f9704b9e27e230d70cc1ed54"></a>

## sriov_interfaces.sriov_interface — sriov_interfaces.sriov_interface / a92a85913417 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [sriov_interfaces](resources--voltstack_site--reference--group-010.md#canonical-e288685eed32403a1b38a053cd958daafe67a3c468f5a4c75bbb371860409f6f)
- sriov_interfaces.sriov_interface

<a id="canonical-ca38562d5123e691e5a471ac14ecb68cd083ea0b3a97fa80a43a630c28c4f49d"></a>

Type: `"object"`. list nested block, Optional.

Use custom SR-IOV interfaces Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("interface_name",
    "number_of_vfs")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
sriov_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-dccf2302f4454d650585013b3f76d85a8c4f2447a12c9996c85b42726db351a0"></a>

## Direct properties — sriov_interfaces.sriov_interface / a92a85913417 / 3

<a id="canonical-ed151f5c65cde9307d787eb3162981878903f754dace4bc7683b18ea11e43f10"></a>

<a id="canonical-70befea6dfc6f15d52326e20421e1f28cb262cd8581818ce5ee21967e211a20c"></a>

## interface_name property — sriov_interfaces.sriov_interface / a92a85913417 / 4

Type: `"string"`. Optional.

Name of physical interface. Name of SR-IOV physical interface.

Upstream description:

Name of SR-IOV physical interface.

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

<a id="canonical-8a64e5e9ede667858956ce1b75901ba11e37651605688708b7c4438125ad6c42"></a>

<a id="canonical-869c44042c276787538dadeba0a397a792cfc983f003ed5ac0cae3f7ca417ec1"></a>

## number_of_vfio_vfs property — sriov_interfaces.sriov_interface / a92a85913417 / 5

Type: `"number"`. Optional.

Number of virtual functions reserved for VNFs and DPDK-based CNFs.

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

<a id="canonical-760a63b58b3384d12d28fa0ce5a6bca5749078ed6cde62d464e0cf8b5f5028d3"></a>

<a id="canonical-52dd0ac57dcbd398cba2a6e92ce56070e5b5663ddb2508d1762e90d66391f8de"></a>

## number_of_vfs property — sriov_interfaces.sriov_interface / a92a85913417 / 6

Type: `"number"`. Optional.

Total number of virtual functions. Total number of virtual functions.

Upstream description:

Total number of virtual functions.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-1cc50413cc8b85a9cc0db0ea5d2f26a7e2ee6a57c15978b3d32ebd792a43b610"></a>

## Next pages — sriov_interfaces.sriov_interface / a92a85913417 / 7

- [sriov_interfaces](resources--voltstack_site--reference--group-010.md#canonical-e288685eed32403a1b38a053cd958daafe67a3c468f5a4c75bbb371860409f6f)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-745052ba3f940f79c0a1e2dd6e9d5edafe93562c98b4f624e7b3234156449238"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fe5fe7d81ea384724e358511cfdba0bfcecaf0a3527825120054112ffb124bec"></a>

## sw — sw / e7a232d71fde / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- sw

<a id="canonical-39d8fa2adfb4c8c286d864b958b3a9f86229743ab999d815af80f42b503185b3"></a>

Type: `"object"`. single nested block, Optional.

Select the F5XC Software Version for the site. By default, latest available F5XC Software Version
will be used. Refer to release notes to find required released SW versions.

Upstream description:

Select the F5XC Software Version for the site. By default, latest available F5XC Software Version
will be used. Refer to release notes to find required released SW versions.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_sw_version",
    "volterra_software_version")}
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
  "x-ves-oneof-field-volterra_sw_version_choice": "[\"default_sw_version\",\"volterra_software_version\"]"
}
```

Terraform syntax:

```terraform
sw {
  # Configure direct properties listed below.
}
```

<a id="canonical-b5526a81571dc3c2136037fe7fefae3eb1f108f76690a2025572c20868aaa34a"></a>

## Direct properties — sw / e7a232d71fde / 3

- [default_sw_version](resources--voltstack_site--reference--group-010.md#canonical-ced8c87d4940d363359cee9cc35a3e01573f0c1a194d46cf5fc98635854619c1): complete subsection reference.

<a id="canonical-fdd2a5e74b5f6142359bb06f12bdcd231ca6e580634180e1df434ad4fbee6b72"></a>

<a id="canonical-a50b79c5594955c2051aabf125545a2fa1d3c44fb4b92932c99b3a5ad7ff07f6"></a>

## volterra_software_version property — sw / e7a232d71fde / 4

Type: `"string"`. Optional.

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

Upstream description:

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-7aac8ee0f6e1bd297ab9946913d441544a858f870f6b9d7c50a364aee622c763"></a>

## Next pages — sw / e7a232d71fde / 5

- [sw.default_sw_version](resources--voltstack_site--reference--group-010.md#canonical-ced8c87d4940d363359cee9cc35a3e01573f0c1a194d46cf5fc98635854619c1)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-ced8c87d4940d363359cee9cc35a3e01573f0c1a194d46cf5fc98635854619c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8437d1de835d0c1bba7f15317fae86a7974c7c99428c79d00e304ddc6e36c000"></a>

## sw.default_sw_version — sw.default_sw_version / 9dd11a1a57e4 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [sw](resources--voltstack_site--reference--group-010.md#canonical-745052ba3f940f79c0a1e2dd6e9d5edafe93562c98b4f624e7b3234156449238)
- sw.default_sw_version

<a id="canonical-4e3e9e3a736583096ca80d2e15f0b4bd995c55f75544681c166033a9ac670e56"></a>

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
default_sw_version = {}
```

<a id="canonical-2b5afd8405eea5275fa9cde4834c5a96a221703a1e9396a5f384a861647dcb60"></a>

## Direct properties — sw.default_sw_version / 9dd11a1a57e4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f56eb8dfd5e8dbddbfc81a4917e03184371e98cd79ba5b3737acbd5edf2de228"></a>

## Next pages — sw.default_sw_version / 9dd11a1a57e4 / 4

- [sw](resources--voltstack_site--reference--group-010.md#canonical-745052ba3f940f79c0a1e2dd6e9d5edafe93562c98b4f624e7b3234156449238)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-5c6a3b09311bf2ef7cafbadfd87aa53d9413906a85e90eac3b032f350aa4daaf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1b032a568e387352c24e3a59255efec5d8c78d62e45ac2cdf14cc4efdb4e17f"></a>

## timeouts — timeouts / 66ac1efae2c4 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- timeouts

<a id="canonical-91650c03f3ef814dc58a839df6200c681eb9a46895288b161eb36d48e3979a97"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-97debd36164e9cd63579fca436468fc60afdacb073a3a74b3fcce75ccd9fd78f"></a>

## Direct properties — timeouts / 66ac1efae2c4 / 3

<a id="canonical-8aee2e66d9a077cab0b36839d34978e8936cffd0ec58e18646427f297afcb1a2"></a>

<a id="canonical-8fe1b36f143e80b653caf1b539e8210819be54d1b902d66f668235ab9038d495"></a>

## create property — timeouts / 66ac1efae2c4 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-716ab09d26a485f39db7920e25e3b15023229e819fca372d3d6a5d4ce1751b5d"></a>

<a id="canonical-5aeffe8c12ca2bac8f1e02d7bafb5acaa8f8e5fd70751847d42c52ed7199097f"></a>

## delete property — timeouts / 66ac1efae2c4 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-45def977491fb29ab62f1b28fe4db82bd9f9b7cf99348e5d6b7c965391ace94f"></a>

<a id="canonical-17a988454894195636b99145cfe84d827ab140a8c2e60ad226b962f023a2241f"></a>

## read property — timeouts / 66ac1efae2c4 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-844eff155aa2dafb1e723002427ab702a2fa9bf81b10f2d07c082c5df22863ab"></a>

<a id="canonical-6662647951b83d668ea74b0287509259fc5c42532f758557bd31efcb5165f8b9"></a>

## update property — timeouts / 66ac1efae2c4 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-41ba8c4c684d17b8a35f9101461f1922698c8df410a8509d678da612409388a8"></a>

## Next pages — timeouts / 66ac1efae2c4 / 8

- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-fed8e4ab7457c4c6764b430be7605274aed28a3dfd7652d7cb4f1ce547934f33"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ab1b8b9f2629d380740afdf35a1b630e23809feb140f7dbc237249e0b3a1ba7"></a>

## usb_policy — usb_policy / 5dfd07bf0df8 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- usb_policy

<a id="canonical-f52a24b8f7ea06acbf4be025298d17455a23564eafd7bcf53723ce2a136a7ab6"></a>

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
usb_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-c74e48e21b201a10e1eb7d588c8d5cafd0d6b05e650caf8f097b957f56b351c0"></a>

## Direct properties — usb_policy / 5dfd07bf0df8 / 3

<a id="canonical-6609fc6e4e5b08dcadafb70ca66f3811c0d72a763202b11a755902e79f0422e0"></a>

<a id="canonical-74702c22322a3c5719e8abb06768713ad6db12d0cd374a4113f8c4d0b1f24a7b"></a>

## name property — usb_policy / 5dfd07bf0df8 / 4

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

<a id="canonical-fd3d73dbd387706fbcb26592b22fd10905973dc8da921fc20ccb506447984574"></a>

<a id="canonical-de5d2d32bc3326e3c5309605c35b9fc307a3dd3347a7cd7ce30f0f3870fae7c6"></a>

## namespace property — usb_policy / 5dfd07bf0df8 / 5

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

<a id="canonical-74c9f7f265929fea87bcd7d2cc8577592c17286754c4de2edc83812619d3be68"></a>

<a id="canonical-3ed3400e3cd2d784b5ce337be95ed9cd8546049cf5ac9e41a8bf13e73a2539e6"></a>

## tenant property — usb_policy / 5dfd07bf0df8 / 6

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

<a id="canonical-e4617123f00453f3de2d30b3c7b940269d5b5ca211fbee7aa380c1816ba41010"></a>

## Next pages — usb_policy / 5dfd07bf0df8 / 7

- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-c19033b27a37f0dcddc37a2716822743c64449717b58b7d4d99a91468d93821e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d532c5bb92ff4134c7d361321871c8fdc44df7e131495d09d0fdb7899124315a"></a>

## waf_signatures — waf_signatures / 66bef01dad48 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- waf_signatures

<a id="canonical-96efaabab6418fe872f388f3cb0b2682e2a6e0f7cf99004b295dcee0efe3ccf9"></a>

Type: `"object"`. single nested block, Optional.

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Upstream description:

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("automatic",
    "manual")}
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
  "x-ves-oneof-field-signatures_update_mode_choice": "[\"automatic\",\"manual\"]"
}
```

Terraform syntax:

```terraform
waf_signatures {
  # Configure direct properties listed below.
}
```

<a id="canonical-ebecd32d1d31f4f0483ee917041c5bac3156f0c46ef992b6b96477310e59b4ac"></a>

## Direct properties — waf_signatures / 66bef01dad48 / 3

- [automatic](resources--voltstack_site--reference--group-010.md#canonical-9e5a625ce48cf402a5849d431868d5c5eb40139b702ee887b80d078ca02437bf): complete subsection reference.

- [manual](resources--voltstack_site--reference--group-010.md#canonical-44f980ef10522eaf15a2816f63e4837eb4e1bd6e325135e1fdaa689915750c76): complete subsection reference.

<a id="canonical-a53e6375ef791276a9b2dfbce8ebf5c3b0bf481ceefa9b324250bef663a3aee2"></a>

## Next pages — waf_signatures / 66bef01dad48 / 4

- [waf_signatures.automatic](resources--voltstack_site--reference--group-010.md#canonical-9e5a625ce48cf402a5849d431868d5c5eb40139b702ee887b80d078ca02437bf)
- [waf_signatures.manual](resources--voltstack_site--reference--group-010.md#canonical-44f980ef10522eaf15a2816f63e4837eb4e1bd6e325135e1fdaa689915750c76)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-9e5a625ce48cf402a5849d431868d5c5eb40139b702ee887b80d078ca02437bf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4cbcce136706f3599e95b9f0704ddbec37f031b9415a9842fae46fa171b6fca1"></a>

## waf_signatures.automatic — waf_signatures.automatic / fe8f8ca9e0a8 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [waf_signatures](resources--voltstack_site--reference--group-010.md#canonical-c19033b27a37f0dcddc37a2716822743c64449717b58b7d4d99a91468d93821e)
- waf_signatures.automatic

<a id="canonical-a2582cf57687e55395acae14fb6f7fe8542d0f4b1029e3222650a9c9d728b31a"></a>

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
automatic = {}
```

<a id="canonical-ae6363ab85921e0693d9f384e822e12d7668cd7f65f9fb2cc76ee064be98fe2b"></a>

## Direct properties — waf_signatures.automatic / fe8f8ca9e0a8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-557763f9e0314715c2f23353ce9b5ea8b3cfb93b3a2fd4a94cf55222b245933f"></a>

## Next pages — waf_signatures.automatic / fe8f8ca9e0a8 / 4

- [waf_signatures](resources--voltstack_site--reference--group-010.md#canonical-c19033b27a37f0dcddc37a2716822743c64449717b58b7d4d99a91468d93821e)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-44f980ef10522eaf15a2816f63e4837eb4e1bd6e325135e1fdaa689915750c76"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e49791b245d3ac6da40dbc5cdfe4960dc326ce2e84f5f9c5177cf0d63c505243"></a>

## waf_signatures.manual — waf_signatures.manual / e44e525ade98 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [waf_signatures](resources--voltstack_site--reference--group-010.md#canonical-c19033b27a37f0dcddc37a2716822743c64449717b58b7d4d99a91468d93821e)
- waf_signatures.manual

<a id="canonical-3a8b68d0f46c16e2de244ad94af1a49e00e446c43a76a8a4c053e15c3350da24"></a>

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
manual = {}
```

<a id="canonical-9af3ba4f2d0ba8e415eb43f50534d30b177921738736dc447588998330ef7172"></a>

## Direct properties — waf_signatures.manual / e44e525ade98 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-98c9df966831d69057b7389bb3f868cb2996bef3245aebda3ad8a6b3599b0555"></a>

## Next pages — waf_signatures.manual / e44e525ade98 / 4

- [waf_signatures](resources--voltstack_site--reference--group-010.md#canonical-c19033b27a37f0dcddc37a2716822743c64449717b58b7d4d99a91468d93821e)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
