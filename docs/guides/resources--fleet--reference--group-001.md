---
page_title: "xcsh_fleet reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fleet reference."
---

# xcsh_fleet reference

<a id="canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ba9b166a3ba49f0e62b72b13d3194a64e41fc0aba80a7e4180c69f9a4137314d"></a>

## Property reference — Property reference / 49d5ced34f3e / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- Property reference

<a id="canonical-79d4857334b1490b9c22311ce8915bf3bc0e28967f46843bda48fff88a672390"></a>

## Direct properties — Property reference / 49d5ced34f3e / 3

- [allow_all_usb](resources--fleet--reference--group-001.md#canonical-2f22e2a90601571efd4c85045edeecb1f252e897cfa44bf672d9267183518fdb): complete subsection reference.

<a id="canonical-fd9de44980a7067c7a00a48f76e87cf00a50a4bef4cd2a141c6fd203b18ca3ef"></a>

<a id="canonical-c9d0a66d5e94d21b92815593bf8a119ecf3539a5d4830c8b57f1570d3137f7d9"></a>

## annotations property — Property reference / 49d5ced34f3e / 4

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

- [blocked_services](resources--fleet--reference--group-001.md#canonical-6a97758cec25e40f0955dfad6ad390b3ed4989ec29accf52af9593619343596f): complete subsection reference.

- [bond_device_list](resources--fleet--reference--group-002.md#canonical-51c90dd8a67fb1ce21ae1dec5a8ff48b84edb403be963bbab7912f7718909b3d): complete subsection reference.

- [dc_cluster_group](resources--fleet--reference--group-002.md#canonical-a8156c5491eb298ebda885fbff5c7c44a0913dddc39143494557eabceea08563): complete subsection reference.

- [dc_cluster_group_inside](resources--fleet--reference--group-002.md#canonical-4ae9b77ddff6d8fbe7167465347d3472c315520e1693b3c6ea03e3c8130f508f): complete subsection reference.

- [default_config](resources--fleet--reference--group-002.md#canonical-f408aab60f432ad21626ed326f86b47229253e65115eb2eac015d6e7593204f7): complete subsection reference.

- [default_sriov_interface](resources--fleet--reference--group-002.md#canonical-39e0721784456474c166b87766c385cdafa1285483adb504e75a5e9fe8aa9371): complete subsection reference.

- [default_storage_class](resources--fleet--reference--group-002.md#canonical-b519b74545b0c70bd2e86bb26ab51f7b6e4d56ec062b14ef4390106d26383010): complete subsection reference.

- [deny_all_usb](resources--fleet--reference--group-002.md#canonical-8643889543f2b9b1ce29e52d76b7057c7e41e19a74827b7ff74eab7619639738): complete subsection reference.

<a id="canonical-9e4905de3a8a7b0d01f43bd052bcbe3e554aa54beda8f46afc864b8ea06a7e14"></a>

<a id="canonical-574ebf55bb2939a371f41c537c1a803ef51d1854aeac8c1d4883c08facde87b1"></a>

## description property — Property reference / 49d5ced34f3e / 5

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

- [device_list](resources--fleet--reference--group-002.md#canonical-d8c812327bc50c891fb3f0faee847413de41c529225e52728413f5a01e42a033): complete subsection reference.

<a id="canonical-c9441eaa682033ba7b744c2fb4362706a23103df98fc34bb834f6acb5e954cdc"></a>

<a id="canonical-ed1b94793a005a65fd2bbf7eecff1171212a8a9a4fd82f66881a24b078465b27"></a>

## disable property — Property reference / 49d5ced34f3e / 6

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

- [disable_gpu](resources--fleet--reference--group-002.md#canonical-1b3a5bfd1dfadfe6e0885816410305d81a80cdad6116c269273c917bce7e8343): complete subsection reference.

- [disable_log_anonymization](resources--fleet--reference--group-002.md#canonical-0cdf79b581c68d230805be43993a64ee5dcfe96ac81609eb67217b2410e598a0): complete subsection reference.

- [disable_vm](resources--fleet--reference--group-002.md#canonical-3904874165a1ef727a4f8ff00d8b5a126ccfc4977500dc375a5c5c689ac1c621): complete subsection reference.

<a id="canonical-0013f1351a9d9d3051a484dc2abe3ac0c43796a2206445fdf8f10907930a5b25"></a>

<a id="canonical-189515cd889621777ba9503424bcbfcbf91fc826d5dce56aa65547277a3fb15f"></a>

## enable_default_fleet_config_download property — Property reference / 49d5ced34f3e / 7

Type: `"bool"`. Optional, Computed.

Enable default fleet config, It must be set for storage config and GPU config.

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

- [enable_gpu](resources--fleet--reference--group-002.md#canonical-0205c6ba7929d8e95a82a34b489cee522a02ce75dfbbbea61460c97cdeb47934): complete subsection reference.

- [enable_log_anonymization](resources--fleet--reference--group-002.md#canonical-70701e28574e291f7bfa0aaa45bad6d40165ac73a4121c875f4b6fea4ece932c): complete subsection reference.

- [enable_vgpu](resources--fleet--reference--group-002.md#canonical-0ba113417dd322922714de6fd64f4975669552d3e9b96ef7fe2a01f57fde4284): complete subsection reference.

- [enable_vm](resources--fleet--reference--group-002.md#canonical-3a7d7361e54b7a7f248f5fa61458fd7e181f92f3712c92302c4da7638c0af9f3): complete subsection reference.

<a id="canonical-9d8cba2b689ef9b42060b1668e1e7e1f5b49fd0cef12eec182a7d152726aba41"></a>

<a id="canonical-55bd4912e66e17a5dd4474f0359de7b1741627b8898da278c9f44400574d0669"></a>

## fleet_label property — Property reference / 49d5ced34f3e / 8

Type: `"string"`. Required.

Fleet\_label value is used to create known\_label 'F5 XC/fleet=&lt;fleet\_label&gt;' The
known\_label is created in the 'shared' namespace for the tenant. A virtual\_site object with name
&lt;fleet\_label&gt; is also created in 'shared' namespace for tenant. The virtual\_site object will
select all sites..

Upstream description:

Fleet\_label value is used to create known\_label "F5 XC/fleet=&lt;fleet\_label&gt;" The
known\_label is created in the "shared" namespace for the tenant.

A virtual\_site object with name &lt;fleet\_label&gt; is also created in "shared" namespace for
tenant. The virtual\_site object will select all sites configured with the known\_label above
fleet\_label with "sfo" will create a known\_label "F5 XC/fleet=sfo" in tenant for the fleet.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.k8s_label_value": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.k8s_label_value": "true"
  }
}
```

<a id="canonical-5138b34b5f9001efeccbb382e0187dabce63e7ae78429bf8ff79d9afc7880a2b"></a>

<a id="canonical-47fdfca9891b92656f2c6dacbc6977f75c3f7b900188acbaf69280eac3fc9ca7"></a>

## id property — Property reference / 49d5ced34f3e / 9

Type: `"string"`. Computed.

Unique identifier for the resource.

- [inside_virtual_network](resources--fleet--reference--group-002.md#canonical-ff836b8717083785fcc763cf98298a2a5b4213a5691c3bba50f99e9fdc6f502a): complete subsection reference.

- [interface_list](resources--fleet--reference--group-002.md#canonical-8fded61b07484d6b4cb7278257551fb9fbae015181ec40bd2bb38d864b381672): complete subsection reference.

- [kubernetes_upgrade_drain](resources--fleet--reference--group-002.md#canonical-b3fca42c2aeb5b7061fe19294dfc4fa35eebac01b348c2d9c5f4431c9390a0b0): complete subsection reference.

<a id="canonical-54b376b4d82c837106c396ec557366685ef7e237df2b5c4803c820a9074866d8"></a>

<a id="canonical-5ba7f4b9c7689cc9c2e9e1edfc19c26f674054f790f29d70349a3259dc70560b"></a>

## labels property — Property reference / 49d5ced34f3e / 10

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

- [log_receiver](resources--fleet--reference--group-002.md#canonical-3df1cf1ed4f9c70511a4f38579ec1cd5749aa8984001404520bc464ea9f2010f): complete subsection reference.

- [logs_streaming_disabled](resources--fleet--reference--group-002.md#canonical-e69db08b22b6264d05d3610c2282151e0a7e3259540950b8051276833640ac65): complete subsection reference.

<a id="canonical-197e820e37907192217ecf4c63d937b6192815aca5a326b021fb4944a90c48e9"></a>

<a id="canonical-2ddec1cba24df11e2b84aaa899a83e9efc43e5f3f7f8ef479ef07962f7660811"></a>

## name property — Property reference / 49d5ced34f3e / 11

Type: `"string"`. Required.

Name of the Fleet. Must be unique within the namespace.

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

<a id="canonical-58f43f4d997495dc361c23cd667e12d7ebd17f99598e37446fafcf6f9bc338da"></a>

<a id="canonical-72bbca775b7ade22a8250404539dd5c58b9393b3215833ab16c5ce6b847e096e"></a>

## namespace property — Property reference / 49d5ced34f3e / 12

Type: `"string"`. Required.

Namespace where the Fleet is created.

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

- [network_connectors](resources--fleet--reference--group-002.md#canonical-fb55531a73258d6b540e94ca88f87f2ddeb571685940897e2797a550f3524359): complete subsection reference.

- [network_firewall](resources--fleet--reference--group-002.md#canonical-917086161f81fec62e50a56057d44ce7779fed3b60c6158a50c95e18eadb9b67): complete subsection reference.

- [no_bond_devices](resources--fleet--reference--group-002.md#canonical-e3e1473959428b2588fe02d236e1647fd23b202356dcfb55d20c97fa2e6f17e1): complete subsection reference.

- [no_dc_cluster_group](resources--fleet--reference--group-002.md#canonical-c2b28909dab56b08f193341580a018bedf81f18d2ada71c1098f29e2a4878160): complete subsection reference.

- [no_storage_device](resources--fleet--reference--group-002.md#canonical-7bfedf647d6488cf41b8b8a93ba5e3859a1b5bf73fe41b2b6c85e198f8287d67): complete subsection reference.

- [no_storage_interfaces](resources--fleet--reference--group-002.md#canonical-c94745f0bc638d751358bc861580c67e9a3a85171a44abe6f03990a8127968ec): complete subsection reference.

- [no_storage_static_routes](resources--fleet--reference--group-002.md#canonical-ad162335f354a460814f7dfc264820728c5dda3f75efcbcb98990ba30d3bd05c): complete subsection reference.

<a id="canonical-ee138f06cfe4db2f11c4c11f7b2929f44696ab9fe96b75998ced3800279e94aa"></a>

<a id="canonical-95f6f1e917ba7d75a8ac361a1977308056af23f1d09a52e237572b608746a273"></a>

## operating_system_version property — Property reference / 49d5ced34f3e / 13

Type: `"string"`. Optional, Computed.

Desired Operating System version that is applied to all sites that are member of the fleet. Current
Operating System version can be overridden via site config.

Upstream description:

Desired Operating System version that is applied to all sites that are member of the fleet. Current
Operating System version can be overridden via site config.

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

- [outside_virtual_network](resources--fleet--reference--group-002.md#canonical-ebf5a068c5d6d88d88f119f0df953517979070c881eac65fe6c7d2a99066adb1): complete subsection reference.

- [performance_enhancement_mode](resources--fleet--reference--group-002.md#canonical-3839c5e410e10cfb7034e4a56e6c05dcf39dfd4936bba7e3e253c7f17c10fcff): complete subsection reference.

- [sriov_interfaces](resources--fleet--reference--group-002.md#canonical-041b0767cc899be9804f15aa4eade24e06041aa19412ec8a5029f8932a306989): complete subsection reference.

- [storage_class_list](resources--fleet--reference--group-002.md#canonical-459e775e0d90b3cce1de02b34a6b7b22245cb6442ce300f58bbc83579f34bbcb): complete subsection reference.

- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251): complete subsection reference.

- [storage_interface_list](resources--fleet--reference--group-004.md#canonical-a0ad9e77e25f172cd7ebf3e39a25f59ef2e0197cf555254fa1eae93c1269011e): complete subsection reference.

- [storage_static_routes](resources--fleet--reference--group-004.md#canonical-f44c11471fbbf1faacf11875c91e3b194ee4c8c8c365df7691d83e66869443bc): complete subsection reference.

- [timeouts](resources--fleet--reference--group-004.md#canonical-c7daec114bab3c6e240b23c9acc29c136c7b4876645a944004681d627803d664): complete subsection reference.

- [usb_policy](resources--fleet--reference--group-004.md#canonical-777cf2733ada0f5b1100fb3c706ce2c64aefbf9dc25e8aae32a1eaa74ade7df2): complete subsection reference.

<a id="canonical-56d6fef330e305019d1650aa97c37e935e7e3e68f2df88243fb4806062b617f2"></a>

<a id="canonical-095bebd324c2a565325a72d086688e2589278a6d66cbd4744af9a5a951342720"></a>

## volterra_software_version property — Property reference / 49d5ced34f3e / 14

Type: `"string"`. Optional, Computed.

F5XC software version is human readable string matching released set of version components. The
given software version is applied to all sites that are member of the fleet. Current software
installed can be overridden via site config.

Upstream description:

F5XC software version is human readable string matching released set of version components. The
given software version is applied to all sites that are member of the fleet. Current software
installed can be overridden via site config.

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

<a id="canonical-7656c288a90a73e167d80b5b44bc4af985bd2b0856eb47d7079547cd262788ed"></a>

## All schema paths — Property reference / 49d5ced34f3e / 15

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allow_all_usb` | [allow_all_usb](resources--fleet--reference--group-001.md#canonical-056a0bb93d92a263e5d04f7d3785064f244a05aea8aeb43741141609e8a664d1) |
| `annotations` | [annotations](resources--fleet--reference--group-001.md#canonical-fd9de44980a7067c7a00a48f76e87cf00a50a4bef4cd2a141c6fd203b18ca3ef) |
| `blocked_services` | [blocked_services](resources--fleet--reference--group-001.md#canonical-b09e1a0d117993efd9166ea5f36dd75f63f536f3bdae105599276761796da6de) |
| `blocked_services.dns` | [blocked_services.dns](resources--fleet--reference--group-002.md#canonical-c10f7e7efc999629ad076c971c91c44c6bd6934562eb3b98d3c8e2fa09bca5f8) |
| `blocked_services.network_type` | [blocked_services.network_type](resources--fleet--reference--group-001.md#canonical-a71102c27ff35832e087b8a02411093e414de3f83ddae52cd768d3f52c35c06e) |
| `blocked_services.ssh` | [blocked_services.ssh](resources--fleet--reference--group-002.md#canonical-32afd442276ed413c0330a7ef0df23e2268e8e5c7c4f0eb73946c6100252224c) |
| `blocked_services.web_user_interface` | [blocked_services.web_user_interface](resources--fleet--reference--group-002.md#canonical-5b376e36687067a41f384e9b91ec9abc8bbe0acbb56b90b30cd67a4d83cee24f) |
| `bond_device_list` | [bond_device_list](resources--fleet--reference--group-002.md#canonical-d6f4fcde44eebbc9015f9223080391343690de95a665759806a0264cbe00fe92) |
| `bond_device_list.bond_devices` | [bond_device_list.bond_devices](resources--fleet--reference--group-002.md#canonical-efb0005235c408ef9515f518ef4d7fc8c9caf6d482232ca4fd2ec794ac57b9e8) |
| `bond_device_list.bond_devices.active_backup` | [bond_device_list.bond_devices.active_backup](resources--fleet--reference--group-002.md#canonical-9e5eb926e877b7f81625a0977d7a5091a63b69dc6ccfbcd3de680f94ed82fa06) |
| `bond_device_list.bond_devices.devices` | [bond_device_list.bond_devices.devices](resources--fleet--reference--group-002.md#canonical-524703ff6ac8163dc14967d057a4912b453d237b8187e15bc9426d61c9ae2d9e) |
| `bond_device_list.bond_devices.lacp` | [bond_device_list.bond_devices.lacp](resources--fleet--reference--group-002.md#canonical-f7eec160d6a831ff7e1c522b917e67e71f313484edeb812d14748ad2e1497f8c) |
| `bond_device_list.bond_devices.lacp.rate` | [bond_device_list.bond_devices.lacp.rate](resources--fleet--reference--group-002.md#canonical-19149360635e75b2b530fd690183c032b222a06877ff861d9fa2d4f5f59f4b6d) |
| `bond_device_list.bond_devices.link_polling_interval` | [bond_device_list.bond_devices.link_polling_interval](resources--fleet--reference--group-002.md#canonical-e88713fad803eebf278e30f6ee80810fb3188d1db2d6fc1d001c4b76c7710dbb) |
| `bond_device_list.bond_devices.link_up_delay` | [bond_device_list.bond_devices.link_up_delay](resources--fleet--reference--group-002.md#canonical-96753f1b330510955259383557a2650ab39183d60d81e53c6a20c4aed233ddfb) |
| `bond_device_list.bond_devices.name` | [bond_device_list.bond_devices.name](resources--fleet--reference--group-002.md#canonical-9fbd33687ce415c853daf219106a006f53398180dc1b213af70fd3a590da3717) |
| `dc_cluster_group` | [dc_cluster_group](resources--fleet--reference--group-002.md#canonical-ea785c94ebbd4cca04543d3d6e66782c26123daca8d6bc5d47f5bf5e05bee390) |
| `dc_cluster_group.name` | [dc_cluster_group.name](resources--fleet--reference--group-002.md#canonical-117455ab6c7107fe689f0cb9b009302ad9332fb3b78b560ad62a0419b01487d8) |
| `dc_cluster_group.namespace` | [dc_cluster_group.namespace](resources--fleet--reference--group-002.md#canonical-9f10988f5f52d8bef5b00af613b1bec183ff8daa3de34ec73346a8417b49bcf7) |
| `dc_cluster_group.tenant` | [dc_cluster_group.tenant](resources--fleet--reference--group-002.md#canonical-8be976c2d0e3a835c1e4221cfe940820eb00c5bcf39c12bc690eb19229a5bbe2) |
| `dc_cluster_group_inside` | [dc_cluster_group_inside](resources--fleet--reference--group-002.md#canonical-d7388d424eee59bb3cf2c40996e79250b7f6a50d67768142871c7409e4e85c9d) |
| `dc_cluster_group_inside.name` | [dc_cluster_group_inside.name](resources--fleet--reference--group-002.md#canonical-85ef35987666fcef75d6c5d9292d6ad3b8d7fffd2325b655fa2c14bd52de334e) |
| `dc_cluster_group_inside.namespace` | [dc_cluster_group_inside.namespace](resources--fleet--reference--group-002.md#canonical-90a0c188394ab001125a9ad26a66ee3c484d55b66905962ac47c10ea1c0479c1) |
| `dc_cluster_group_inside.tenant` | [dc_cluster_group_inside.tenant](resources--fleet--reference--group-002.md#canonical-8633f8343febdb49c8fd2c6f7a5fffaf59d5958499eca02ccb8639c59784eda5) |
| `default_config` | [default_config](resources--fleet--reference--group-002.md#canonical-2af4ded82fb31ddcaac369739df1ad90cc313036f9c0e188699d7d60f2e051f9) |
| `default_sriov_interface` | [default_sriov_interface](resources--fleet--reference--group-002.md#canonical-d8b697cfdf5f08fef0980e66accb2b4fd4efb742ff0b390355312944e8c237ed) |
| `default_storage_class` | [default_storage_class](resources--fleet--reference--group-002.md#canonical-eed09e31ee054f5b9aeb6096c8cb62afe9e1c35b50c24142b2dcdb7cbefbe61c) |
| `deny_all_usb` | [deny_all_usb](resources--fleet--reference--group-002.md#canonical-cda8d71228efbcd20187173834fe8393887e2250e76378ab87f6892e4271db3c) |
| `description` | [description](resources--fleet--reference--group-001.md#canonical-9e4905de3a8a7b0d01f43bd052bcbe3e554aa54beda8f46afc864b8ea06a7e14) |
| `device_list` | [device_list](resources--fleet--reference--group-002.md#canonical-e0ea2efc3b7c31fd5a71ae27e992e44760ee8d422eed5cd594765164089eadb7) |
| `device_list.devices` | [device_list.devices](resources--fleet--reference--group-002.md#canonical-2d3521d60b56bc9189d09e8360da968e70ad6047bd94c388f081341f8f896127) |
| `device_list.devices.name` | [device_list.devices.name](resources--fleet--reference--group-002.md#canonical-50b626e48550558197f2b705a69eaa37e153aa3cc062d29f55e50771036e278d) |
| `device_list.devices.network_device` | [device_list.devices.network_device](resources--fleet--reference--group-002.md#canonical-6f6a686cdc27e3d7d9159d17e9409539c6d58bda101effcd1a6f0c93ea064fc6) |
| `device_list.devices.network_device.interface` | [device_list.devices.network_device.interface](resources--fleet--reference--group-002.md#canonical-d0603db303505c8ef8366b59a8903d1693906264d2efb37d1fdb2c3a42cd38d6) |
| `device_list.devices.network_device.interface.kind` | [device_list.devices.network_device.interface.kind](resources--fleet--reference--group-002.md#canonical-461064d4cc1182f940b67c9ffe8ceb271e1b51b4f8ba48bbd3012a0b478896fb) |
| `device_list.devices.network_device.interface.name` | [device_list.devices.network_device.interface.name](resources--fleet--reference--group-002.md#canonical-efff50170807798660b71e5f6a53447e4aa1585d4687cb06757aeffd8f874038) |
| `device_list.devices.network_device.interface.namespace` | [device_list.devices.network_device.interface.namespace](resources--fleet--reference--group-002.md#canonical-a727f666499e59121a26dbc7f4a18530dab9f7ac878380dd499133b622d92d09) |
| `device_list.devices.network_device.interface.tenant` | [device_list.devices.network_device.interface.tenant](resources--fleet--reference--group-002.md#canonical-2ce738f541cc0beff7395226ec9b2b7e2bed072dc0ffaee0c1316f87316e832b) |
| `device_list.devices.network_device.interface.uid` | [device_list.devices.network_device.interface.uid](resources--fleet--reference--group-002.md#canonical-2fa4fdf4c4a03627a6358d91288a3d323dc9557706052aceffe211f45be94c62) |
| `device_list.devices.network_device.use` | [device_list.devices.network_device.use](resources--fleet--reference--group-002.md#canonical-ad9661816ce9ae8b91ef66930b18862753e4f4d39c40b4c8dcdab09ad80c5f74) |
| `device_list.devices.owner` | [device_list.devices.owner](resources--fleet--reference--group-002.md#canonical-a59cfb324dbfa1a1cac6521b29734af0848f138ec8cfbd53af95c4065655026f) |
| `disable` | [disable](resources--fleet--reference--group-001.md#canonical-c9441eaa682033ba7b744c2fb4362706a23103df98fc34bb834f6acb5e954cdc) |
| `disable_gpu` | [disable_gpu](resources--fleet--reference--group-002.md#canonical-273c9ae58535ab800df5c4b52b43706a0e0bf08bea30ba01684ff4b5fbcac87c) |
| `disable_log_anonymization` | [disable_log_anonymization](resources--fleet--reference--group-002.md#canonical-99072fd1e5f4f967ef9fdb47b1930b9b0f50d46a3609fa9c1843baf13dc382c1) |
| `disable_vm` | [disable_vm](resources--fleet--reference--group-002.md#canonical-ecf67a92c47a58bb32c2a6707a2cda35fbea70b25a4dd9f779c987d9a5418e1f) |
| `enable_default_fleet_config_download` | [enable_default_fleet_config_download](resources--fleet--reference--group-001.md#canonical-0013f1351a9d9d3051a484dc2abe3ac0c43796a2206445fdf8f10907930a5b25) |
| `enable_gpu` | [enable_gpu](resources--fleet--reference--group-002.md#canonical-a0ce803d41a88699778aac5cfcfdf29bf2e0bcc821216f38b1c95313e879eda2) |
| `enable_log_anonymization` | [enable_log_anonymization](resources--fleet--reference--group-002.md#canonical-e0e20339cf1eac3e84b1ed29cedffbfa3edd781961abd3f3d88a6471c84dd443) |
| `enable_vgpu` | [enable_vgpu](resources--fleet--reference--group-002.md#canonical-8a49041fefa8f979e0be44ff5d479735988a3f1172c05139c337d41ddacd460c) |
| `enable_vgpu.feature_type` | [enable_vgpu.feature_type](resources--fleet--reference--group-002.md#canonical-8b917d7f3895a0044be46081e412b403792265781a3d85773e3c536a1223feb4) |
| `enable_vgpu.server_address` | [enable_vgpu.server_address](resources--fleet--reference--group-002.md#canonical-4b76a8a73da3b73f1d5605238f6d1055d949426b11398aed19044462fca30b66) |
| `enable_vgpu.server_port` | [enable_vgpu.server_port](resources--fleet--reference--group-002.md#canonical-14cbb9a06e7877ae057d17ad45568da18820d64e3814f42e24574a6e21296dfe) |
| `enable_vm` | [enable_vm](resources--fleet--reference--group-002.md#canonical-30d153fffb1b0bd92be3f9a98f966c1a386aae419ca40a2b5d9d98dc9321d2ae) |
| `fleet_label` | [fleet_label](resources--fleet--reference--group-001.md#canonical-9d8cba2b689ef9b42060b1668e1e7e1f5b49fd0cef12eec182a7d152726aba41) |
| `id` | [id](resources--fleet--reference--group-001.md#canonical-5138b34b5f9001efeccbb382e0187dabce63e7ae78429bf8ff79d9afc7880a2b) |
| `inside_virtual_network` | [inside_virtual_network](resources--fleet--reference--group-002.md#canonical-9ae9d1b7dcc436c96673bc95239adbd4039ed61c258ab78facea55c2c7f1399c) |
| `inside_virtual_network.kind` | [inside_virtual_network.kind](resources--fleet--reference--group-002.md#canonical-ac338635710b7f514f34d7f2dc149f85630caa0845b7dcca04996d6eac96aa08) |
| `inside_virtual_network.name` | [inside_virtual_network.name](resources--fleet--reference--group-002.md#canonical-452d250abf99bab662a16913ee85d94d1f776b3995096f5358311292d9036ac3) |
| `inside_virtual_network.namespace` | [inside_virtual_network.namespace](resources--fleet--reference--group-002.md#canonical-952a0a2e74b8b47ab07f188578d8467f8cdf7c18336a46010ac7e915db5ef98f) |
| `inside_virtual_network.tenant` | [inside_virtual_network.tenant](resources--fleet--reference--group-002.md#canonical-8ab272073f93d4b6b494faa7fcdc3bced20d4e68c530d3a3640ee64d9c3f6773) |
| `inside_virtual_network.uid` | [inside_virtual_network.uid](resources--fleet--reference--group-002.md#canonical-e9ab759ddde1e9a3d5ebcb1e0bb15fc7f65e4e86c9e7e1dc641fa8bf6705403a) |
| `interface_list` | [interface_list](resources--fleet--reference--group-002.md#canonical-4a4b698279ff0f98428b2935438a8cbb8bf67ae0282b09729c810070c9e7a49b) |
| `interface_list.interfaces` | [interface_list.interfaces](resources--fleet--reference--group-002.md#canonical-adfcbaeab029791d1386946c72ee21901c65af3f4110431b324ae7915393e8a6) |
| `interface_list.interfaces.name` | [interface_list.interfaces.name](resources--fleet--reference--group-002.md#canonical-060b2c4172341b0f1160a472e8371681c0f1a349960e7c0119c523965415d000) |
| `interface_list.interfaces.namespace` | [interface_list.interfaces.namespace](resources--fleet--reference--group-002.md#canonical-40a6cbebbc83bfbcafc2fbc82b15d9d25748270b0096bde8b6b647963c728b1f) |
| `interface_list.interfaces.tenant` | [interface_list.interfaces.tenant](resources--fleet--reference--group-002.md#canonical-1d0c0432b2aa48e559fbf76283ad1a47503e62f37f7ee4b103dbfd905d2c3829) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](resources--fleet--reference--group-002.md#canonical-8971871d1d85a395101679abfde54551f500f5fb7e2ad2ac287b841f29168b8b) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](resources--fleet--reference--group-002.md#canonical-6cc4b3f1ad1bba8b4abcdf7599c6d80a8eb454be535932e8d05329a6eba685e1) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](resources--fleet--reference--group-002.md#canonical-a00f954692d8e0dc3de8bed04fac73c41e53ad8399e108d6a8058f0e39876b74) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](resources--fleet--reference--group-002.md#canonical-f34e0b9358558c9291923df0593b4a5c6709e1acaf15234f01808a0ca6244345) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](resources--fleet--reference--group-002.md#canonical-36ce0a30946235e41ff09963ca5f8575a54242bc7a2c5ba6284cc49b8d91f211) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](resources--fleet--reference--group-002.md#canonical-7c2ad584dfacb5c77461e1413d8031ec8ead54fdc04930a8b23557497c04f778) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](resources--fleet--reference--group-002.md#canonical-fb333a6f902b4ced29c9dcaec51a7b9601908a1932ada889fe02b3d28b8dedda) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](resources--fleet--reference--group-002.md#canonical-017aede9719bc56d21cee86108ef29f624c7b54df33d740a9d975f12b4eb2bf0) |
| `labels` | [labels](resources--fleet--reference--group-001.md#canonical-54b376b4d82c837106c396ec557366685ef7e237df2b5c4803c820a9074866d8) |
| `log_receiver` | [log_receiver](resources--fleet--reference--group-002.md#canonical-cdde11d2c6386cf836477713fa0f15f368ede2ca31331edfaa3c9cc48a69c523) |
| `log_receiver.name` | [log_receiver.name](resources--fleet--reference--group-002.md#canonical-5c217248f8b9d89732195d0304c4edfd9bf698c7b494a724b39adc44f31978a2) |
| `log_receiver.namespace` | [log_receiver.namespace](resources--fleet--reference--group-002.md#canonical-e3028909862a366bedc1854cb5f7e3431b9df0955a84e46cc439326ddae0dd21) |
| `log_receiver.tenant` | [log_receiver.tenant](resources--fleet--reference--group-002.md#canonical-56ba50b3e7908d57ec4a2b8edddd6b9f7c11dcb0f7d8e98f327ad0263e54796a) |
| `logs_streaming_disabled` | [logs_streaming_disabled](resources--fleet--reference--group-002.md#canonical-9d2cc920b55903be230ee5dbd1edcc516d596a464f975c019d15bc5a2ebb13fb) |
| `name` | [name](resources--fleet--reference--group-001.md#canonical-197e820e37907192217ecf4c63d937b6192815aca5a326b021fb4944a90c48e9) |
| `namespace` | [namespace](resources--fleet--reference--group-001.md#canonical-58f43f4d997495dc361c23cd667e12d7ebd17f99598e37446fafcf6f9bc338da) |
| `network_connectors` | [network_connectors](resources--fleet--reference--group-002.md#canonical-af5794e2f807506b7281a6dd584b78fa4736c40260ad523014d4daa50cc18a94) |
| `network_connectors.kind` | [network_connectors.kind](resources--fleet--reference--group-002.md#canonical-a8be39c2444f316208cadb5c3ef98801cc063f14a1ff3a2abd19880b7fa7cb1b) |
| `network_connectors.name` | [network_connectors.name](resources--fleet--reference--group-002.md#canonical-6af5848aa73f4c847716fae45b1f68c143137825548114addc45c44cf089b43e) |
| `network_connectors.namespace` | [network_connectors.namespace](resources--fleet--reference--group-002.md#canonical-0fa5f44d612335468c6b6aff809126615555b71eb1ddca02b9f54dbfd6149268) |
| `network_connectors.tenant` | [network_connectors.tenant](resources--fleet--reference--group-002.md#canonical-ee43e22e2b96d0091eebc03a116f98be72353cc41138e7a4bbb5a659c900cd50) |
| `network_connectors.uid` | [network_connectors.uid](resources--fleet--reference--group-002.md#canonical-d2296be840e97b073edb7b9913b47060652e6037bbbb04c8958bff3a289bc4b8) |
| `network_firewall` | [network_firewall](resources--fleet--reference--group-002.md#canonical-b89363f2032712c62d7729bf8ac837db6a3a230bf154d751646bd385b85b7679) |
| `network_firewall.kind` | [network_firewall.kind](resources--fleet--reference--group-002.md#canonical-2b321283d83180a8a128aa132ba0f635199e6109c14234ac4c8033ddb9e30f38) |
| `network_firewall.name` | [network_firewall.name](resources--fleet--reference--group-002.md#canonical-49155735b25eb6dbd87efc3f75e29fc676be1f14e78b6d31fbe385d6ecf9089c) |
| `network_firewall.namespace` | [network_firewall.namespace](resources--fleet--reference--group-002.md#canonical-c6c5ab9ef49b485f0b6f0967202af552615de41f1a5ca1d5cb3900545f301d7d) |
| `network_firewall.tenant` | [network_firewall.tenant](resources--fleet--reference--group-002.md#canonical-b611583647f93ef74c4b2a9afbb00a5d6332c059103eed6b10e3d918d5769a80) |
| `network_firewall.uid` | [network_firewall.uid](resources--fleet--reference--group-002.md#canonical-00424eedf504520efa160409f8152d60e714ba132e5ed888accb7fa455e2c5d9) |
| `no_bond_devices` | [no_bond_devices](resources--fleet--reference--group-002.md#canonical-af89e5997094bf29a67e421689e9aeb639ac189ee8706aa5e45126b37ea852ec) |
| `no_dc_cluster_group` | [no_dc_cluster_group](resources--fleet--reference--group-002.md#canonical-305750f6b8495e14bfa7faa9b9ac0f801ea25a1729571c25cdd5d1c4c16e69df) |
| `no_storage_device` | [no_storage_device](resources--fleet--reference--group-002.md#canonical-28b1e2e906684aa694e1bb68c060bc08f7118725fa0d5e4d450d2b1c1eee1139) |
| `no_storage_interfaces` | [no_storage_interfaces](resources--fleet--reference--group-002.md#canonical-1e96319fe137aeb87ecaa02341cc28a13b5aa7d7f70f279d04b68c68baea807a) |
| `no_storage_static_routes` | [no_storage_static_routes](resources--fleet--reference--group-002.md#canonical-ff0fc0bb88dba7eac3a72f77e9da989250b2cd2aa24707a217aa5fdf49fdd67d) |
| `operating_system_version` | [operating_system_version](resources--fleet--reference--group-001.md#canonical-ee138f06cfe4db2f11c4c11f7b2929f44696ab9fe96b75998ced3800279e94aa) |
| `outside_virtual_network` | [outside_virtual_network](resources--fleet--reference--group-002.md#canonical-d860e456781b5398db5261f12a005cf0088d63ace1b2f5af138331076ff9fbce) |
| `outside_virtual_network.kind` | [outside_virtual_network.kind](resources--fleet--reference--group-002.md#canonical-13e7b8845642c9ed87abc41dd0c1cd779b51484707b3b5e04056d9df9d24e1d5) |
| `outside_virtual_network.name` | [outside_virtual_network.name](resources--fleet--reference--group-002.md#canonical-d53330cf4c54d008222d99292a92e3cc7c6c78948b5977343912741e60d27d4f) |
| `outside_virtual_network.namespace` | [outside_virtual_network.namespace](resources--fleet--reference--group-002.md#canonical-5a5d396383704c636696fd44c2bce72f0c4a178a8b2ad570c1bde3fddd96f314) |
| `outside_virtual_network.tenant` | [outside_virtual_network.tenant](resources--fleet--reference--group-002.md#canonical-2618273b457b2be388765cf48a853d91087d603cb0748b91082b3240a6d20fc4) |
| `outside_virtual_network.uid` | [outside_virtual_network.uid](resources--fleet--reference--group-002.md#canonical-460f246f6b3723cede68144eb61a6cc1746a692f73620919cfd7a2713c0d9eff) |
| `performance_enhancement_mode` | [performance_enhancement_mode](resources--fleet--reference--group-002.md#canonical-cc5d743c067f55f909e04676cac0bfe81cf7c386b2c03a22c9e66efac39d25d4) |
| `performance_enhancement_mode.perf_mode_l3_enhanced` | [performance_enhancement_mode.perf_mode_l3_enhanced](resources--fleet--reference--group-002.md#canonical-e392a62667b95a4c7cc7cb7bf8af98f5642daed8540a13377bf4dfaf2ad9ab69) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--fleet--reference--group-002.md#canonical-b969da8e75f76b8b687839c477f6fad1b110a4cef9616bfa8caecd6e9560c118) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--fleet--reference--group-002.md#canonical-297459f3ffca55cc3e4c77bd37977ec27656b62a7ccbfb1b6929e7059d20a28d) |
| `performance_enhancement_mode.perf_mode_l7_enhanced` | [performance_enhancement_mode.perf_mode_l7_enhanced](resources--fleet--reference--group-002.md#canonical-127d441a67d9d59de6de2c12c49c294756fde2d2b3498a5f7638edea73e3fa11) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--fleet--reference--group-002.md#canonical-bb3cd0e21ba80a3674212b84797ebddec2a42f32abd036f5a5510336efdfd614) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--fleet--reference--group-002.md#canonical-d0513dc3a8a00814d954b44b757155e0589b70d03b918e5480633ec0b34aea58) |
| `sriov_interfaces` | [sriov_interfaces](resources--fleet--reference--group-002.md#canonical-1aa5dbb2713fa5f114ae6165ef4b850f1922244dfcb047c9bae34345301df9d4) |
| `sriov_interfaces.sriov_interface` | [sriov_interfaces.sriov_interface](resources--fleet--reference--group-002.md#canonical-15ddc56b11fafc30a3485c1af305026961dfda73b840b7fcb0c1b9eb07340d55) |
| `sriov_interfaces.sriov_interface.interface_name` | [sriov_interfaces.sriov_interface.interface_name](resources--fleet--reference--group-002.md#canonical-9bf749d8a4b0e116b8a12840476886ff9fc66271f2c1d618bfe177ef850ef8d7) |
| `sriov_interfaces.sriov_interface.number_of_vfio_vfs` | [sriov_interfaces.sriov_interface.number_of_vfio_vfs](resources--fleet--reference--group-002.md#canonical-8b7dab7417a839c9377b05c4ea7f1c465bf2b32262a1c8729b77efc1e7b40514) |
| `sriov_interfaces.sriov_interface.number_of_vfs` | [sriov_interfaces.sriov_interface.number_of_vfs](resources--fleet--reference--group-002.md#canonical-6f239eb5b61082319879cae4b60ae050bda64ce9dc0f3b16fc7663764b360565) |
| `storage_class_list` | [storage_class_list](resources--fleet--reference--group-002.md#canonical-aac0b2f79acb7a15afc301b7f3a50755163d2e5e7953a62913abaf5050949e98) |
| `storage_class_list.storage_classes` | [storage_class_list.storage_classes](resources--fleet--reference--group-002.md#canonical-7c9f9a9b595d198f1cf6ff18bb977da4e620a8ef9276605d3ca171548bee5d97) |
| `storage_class_list.storage_classes.advanced_storage_parameters` | [storage_class_list.storage_classes.advanced_storage_parameters](resources--fleet--reference--group-002.md#canonical-a931bc2f14b1817e73885700885ac1c2e9a1c49012a7e14697908f35571812b9) |
| `storage_class_list.storage_classes.allow_volume_expansion` | [storage_class_list.storage_classes.allow_volume_expansion](resources--fleet--reference--group-002.md#canonical-a27d5bb56df4619b592e90ed92fb03bd5e904555ffe071c67c9043460b64359d) |
| `storage_class_list.storage_classes.custom_storage` | [storage_class_list.storage_classes.custom_storage](resources--fleet--reference--group-002.md#canonical-cfd7db69750447be25501242b25ad2d49ae983e72b4eda012b0fc8d6f4b9fb44) |
| `storage_class_list.storage_classes.custom_storage.yaml` | [storage_class_list.storage_classes.custom_storage.yaml](resources--fleet--reference--group-002.md#canonical-06ae72c98ab4bc17a4aafff2b08a7fd1527c08a8f39d8fbec958692b68c4bce5) |
| `storage_class_list.storage_classes.default_storage_class` | [storage_class_list.storage_classes.default_storage_class](resources--fleet--reference--group-002.md#canonical-34758fd3dd1cb18bfee403782ca133545af7dcefb895d079f2886cccd50e6456) |
| `storage_class_list.storage_classes.description_spec` | [storage_class_list.storage_classes.description_spec](resources--fleet--reference--group-002.md#canonical-5886d054f6c2d1f5536f4467cb59c2f64f0211289d510d78ef9e824c781b4da3) |
| `storage_class_list.storage_classes.hpe_storage` | [storage_class_list.storage_classes.hpe_storage](resources--fleet--reference--group-002.md#canonical-0ce20e695a3ae2e2443d077133baea3518615d2f75fd4be94e4c38c7f7ad078b) |
| `storage_class_list.storage_classes.hpe_storage.allow_mutations` | [storage_class_list.storage_classes.hpe_storage.allow_mutations](resources--fleet--reference--group-002.md#canonical-c9ab33f50e2c59871d14ed554694f56819526e28fde39d68951df5e9af3adb11) |
| `storage_class_list.storage_classes.hpe_storage.allow_overrides` | [storage_class_list.storage_classes.hpe_storage.allow_overrides](resources--fleet--reference--group-002.md#canonical-530c0f9a5c0ac3e2207140849f6dd715d210409f6ff8b6f6c47af0f63b903fd0) |
| `storage_class_list.storage_classes.hpe_storage.dedupe_enabled` | [storage_class_list.storage_classes.hpe_storage.dedupe_enabled](resources--fleet--reference--group-002.md#canonical-3c9e319c9d3ab2182e55e40447ff220120c0797720c1fc9e470a5ae24eabd76e) |
| `storage_class_list.storage_classes.hpe_storage.description_spec` | [storage_class_list.storage_classes.hpe_storage.description_spec](resources--fleet--reference--group-002.md#canonical-26eda3f34ce38df21d99cc3ceb6c2c18db059ce81b127c8bc7a9e9f9c88e0741) |
| `storage_class_list.storage_classes.hpe_storage.destroy_on_delete` | [storage_class_list.storage_classes.hpe_storage.destroy_on_delete](resources--fleet--reference--group-002.md#canonical-8ed6190891c7adf83c4c8675752492c0456fcdbe322bb2f58f3bb0eb48d08de3) |
| `storage_class_list.storage_classes.hpe_storage.encrypted` | [storage_class_list.storage_classes.hpe_storage.encrypted](resources--fleet--reference--group-002.md#canonical-e7bce3a2ea6094dc4801062723386a110b3ffc751449b61a5f9304a9deb6d676) |
| `storage_class_list.storage_classes.hpe_storage.folder` | [storage_class_list.storage_classes.hpe_storage.folder](resources--fleet--reference--group-002.md#canonical-10af22f76c9991d5c81f1f2ac0943d67ef55e4e9b447e2a64058a8c046e68577) |
| `storage_class_list.storage_classes.hpe_storage.limit_iops` | [storage_class_list.storage_classes.hpe_storage.limit_iops](resources--fleet--reference--group-002.md#canonical-3ddbfcece73898c1f3f6a0d965a53b5a03915bb363edc4e2bd768bfe886bdaa5) |
| `storage_class_list.storage_classes.hpe_storage.limit_mbps` | [storage_class_list.storage_classes.hpe_storage.limit_mbps](resources--fleet--reference--group-002.md#canonical-008a1457805c9475764d820c695bee7c481da381ab898d45115c254ffd286d05) |
| `storage_class_list.storage_classes.hpe_storage.performance_policy` | [storage_class_list.storage_classes.hpe_storage.performance_policy](resources--fleet--reference--group-002.md#canonical-b63b5de793c756a5c30abe2897e55a9b142de97d9b2a60de2e441bfe978b730a) |
| `storage_class_list.storage_classes.hpe_storage.pool` | [storage_class_list.storage_classes.hpe_storage.pool](resources--fleet--reference--group-002.md#canonical-5b63700ec49e787421d1bac29c8c11e397839ced4123e1162ea0600bfb00230e) |
| `storage_class_list.storage_classes.hpe_storage.protection_template` | [storage_class_list.storage_classes.hpe_storage.protection_template](resources--fleet--reference--group-002.md#canonical-5cef584e58123d437bf413f7a163c26b3667e517011170ebd2bbd12ddb65b1f0) |
| `storage_class_list.storage_classes.hpe_storage.secret_name` | [storage_class_list.storage_classes.hpe_storage.secret_name](resources--fleet--reference--group-002.md#canonical-336e90e0d05f3fead6d401bb2a9d3434614107c941124a59421bb44634511fff) |
| `storage_class_list.storage_classes.hpe_storage.secret_namespace` | [storage_class_list.storage_classes.hpe_storage.secret_namespace](resources--fleet--reference--group-002.md#canonical-8984cbd50286959391aab60c68dcb2e42033d3ec232cad359686599208f6e904) |
| `storage_class_list.storage_classes.hpe_storage.sync_on_detach` | [storage_class_list.storage_classes.hpe_storage.sync_on_detach](resources--fleet--reference--group-002.md#canonical-69429462027a8a7c2e51959ae67a2600b9d93694e41e069e3b59050c88350cb0) |
| `storage_class_list.storage_classes.hpe_storage.thick` | [storage_class_list.storage_classes.hpe_storage.thick](resources--fleet--reference--group-002.md#canonical-93727d6b3d8ddfcd866023e948ddf9fa4feec21b416ef6b50487a05f556b445b) |
| `storage_class_list.storage_classes.netapp_trident` | [storage_class_list.storage_classes.netapp_trident](resources--fleet--reference--group-002.md#canonical-e5ed56a08ab5fd2a957e6f30445ed47c09cde002d047ac168eba253eed48e5f4) |
| `storage_class_list.storage_classes.netapp_trident.selector` | [storage_class_list.storage_classes.netapp_trident.selector](resources--fleet--reference--group-002.md#canonical-6c07db8b59814b2b7893613aaa573a624c38feb6672b0b2d1b8d8b0adf7a2a95) |
| `storage_class_list.storage_classes.netapp_trident.storage_pools` | [storage_class_list.storage_classes.netapp_trident.storage_pools](resources--fleet--reference--group-002.md#canonical-7b82a2f0e26e166ee0c87fdd626b3e0937802e54776fe8d7536bf7ae60b7fe0f) |
| `storage_class_list.storage_classes.pure_service_orchestrator` | [storage_class_list.storage_classes.pure_service_orchestrator](resources--fleet--reference--group-002.md#canonical-92387332f70bc6a9e53d9c8734d6417b2084d6c4af42b8c27fc5f30f81483bc9) |
| `storage_class_list.storage_classes.pure_service_orchestrator.backend` | [storage_class_list.storage_classes.pure_service_orchestrator.backend](resources--fleet--reference--group-002.md#canonical-24129632c22e833517c09f6f7fed65faa46dce6c5ae22db4e6b13b17cec623b6) |
| `storage_class_list.storage_classes.pure_service_orchestrator.bandwidth_limit` | [storage_class_list.storage_classes.pure_service_orchestrator.bandwidth_limit](resources--fleet--reference--group-002.md#canonical-6c1edae025669ddf5a722ef50c378a95e0b8e1f58ad6e31b40b0422103ab7b98) |
| `storage_class_list.storage_classes.pure_service_orchestrator.iops_limit` | [storage_class_list.storage_classes.pure_service_orchestrator.iops_limit](resources--fleet--reference--group-002.md#canonical-162d3c96732092f4a5a636a4653b74cb1c95f8e5b21c8421d90ee74c351b4c9d) |
| `storage_class_list.storage_classes.reclaim_policy` | [storage_class_list.storage_classes.reclaim_policy](resources--fleet--reference--group-002.md#canonical-69f0252e7c9049f3da764911bb8ddedfa4ad6e5df28cf5da33fc5eeb7c565423) |
| `storage_class_list.storage_classes.storage_class_name` | [storage_class_list.storage_classes.storage_class_name](resources--fleet--reference--group-002.md#canonical-5ad8c5169bc0e2507276c4874489a75bcca51962a0603632b9e327937ac00b2d) |
| `storage_class_list.storage_classes.storage_device` | [storage_class_list.storage_classes.storage_device](resources--fleet--reference--group-002.md#canonical-45b2d2bb467fbc34f14f8afa0d8d069b8b5a891c043acbfa69eb1e671d017f9b) |
| `storage_device_list` | [storage_device_list](resources--fleet--reference--group-002.md#canonical-68f57ce684d63f235b10bdc4bfd45d3b879dc8cc2e2bc8db8f84c948e982b7dd) |
| `storage_device_list.storage_devices` | [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-b329a485c3723c769226035c47e51b84ca526d8270aae686d45f259e53a4f965) |
| `storage_device_list.storage_devices.advanced_advanced_parameters` | [storage_device_list.storage_devices.advanced_advanced_parameters](resources--fleet--reference--group-002.md#canonical-f4349d215c3678daaa4cc5a4296d24ca4d04d23779fe2311abacb28a1b09ddc1) |
| `storage_device_list.storage_devices.custom_storage` | [storage_device_list.storage_devices.custom_storage](resources--fleet--reference--group-002.md#canonical-10614eab16f3b1921c6a9a96803e2f93a5a0062c62e60e5668e8b4506f66d11c) |
| `storage_device_list.storage_devices.hpe_storage` | [storage_device_list.storage_devices.hpe_storage](resources--fleet--reference--group-002.md#canonical-706c886464eddd7fc46afe0a030faa669f8617478fe3d97a3ec0a21bbb19edca) |
| `storage_device_list.storage_devices.hpe_storage.api_server_port` | [storage_device_list.storage_devices.hpe_storage.api_server_port](resources--fleet--reference--group-003.md#canonical-106428d6e2e564d3731ac6db4eedd7270634a3482e5005c9ee8654bed88a83c1) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](resources--fleet--reference--group-003.md#canonical-7f8000444a74c9fd256601cf411654d0b0649b2ea27f7ef2530af8c04ec88ab4) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-858e1617458c60adf918fab6f282400f11e29f0f74a10e16f74d7a24d948d3c8) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.decryption_provider](resources--fleet--reference--group-003.md#canonical-34ad5f93daa7e500515886c9439ab467aa1a40f459620ce18dbc2ebec7f0e7f0) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.location` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.location](resources--fleet--reference--group-003.md#canonical-29d338b7991418763b389b586dcf48b3009af7fa9fb3d327cc034bc69c810f6d) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.store_provider](resources--fleet--reference--group-003.md#canonical-f507a9b9bf08926afa606f5930810388d25c6788b2bef3fd1635c5eb8dc92603) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info](resources--fleet--reference--group-003.md#canonical-8dce94aa0c54d08cbf4605a380a0556b47b53e32a89b269d080ac7b324064cc9) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info.provider_ref](resources--fleet--reference--group-003.md#canonical-7536415376ea51cd94688a7273f8a0b3aa69774e5be704fdb835ed0172cae5a3) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info.url` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info.url](resources--fleet--reference--group-003.md#canonical-0b15c10824887434b2d91ffdc691af1dfc9406e0dd3c0fb77d947a0ba62f26d0) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_user` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_user](resources--fleet--reference--group-003.md#canonical-572f5fb055b30fddbba565541374f58ea46aac3a4bae0df5c44af57d8856ac33) |
| `storage_device_list.storage_devices.hpe_storage.password` | [storage_device_list.storage_devices.hpe_storage.password](resources--fleet--reference--group-003.md#canonical-9d0c81aa7622a68dcfcfa088a31326c52ce5be59cff6a4357cc0a5370c3c2d20) |
| `storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info` | [storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-759557e0475f901b594ba22f622de4c225ead2edd6c35328c0e5165315d9e0b9) |
| `storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.decryption_provider](resources--fleet--reference--group-003.md#canonical-6c8259e30b4e82b886655ab6e28e20bde7f8e5fcd6815bbd3b3ca5485ab903d3) |
| `storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.location` | [storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.location](resources--fleet--reference--group-003.md#canonical-841fb9b5364038b2b7a455b4e970405afde6e040070baf18f5087e563e73e584) |
| `storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.store_provider](resources--fleet--reference--group-003.md#canonical-7df905360e30c54889bbf0db2f8d7acd7f3ae6e68857fc3408716b4f6c4b578f) |
| `storage_device_list.storage_devices.hpe_storage.password.clear_secret_info` | [storage_device_list.storage_devices.hpe_storage.password.clear_secret_info](resources--fleet--reference--group-003.md#canonical-f6af0b2aa94bc783c877b1d21edc329b6e4fed8d031ea3b546af0e0a8dd98d1e) |
| `storage_device_list.storage_devices.hpe_storage.password.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.hpe_storage.password.clear_secret_info.provider_ref](resources--fleet--reference--group-003.md#canonical-daa21dc01df2087932566aa231d998ed82f309538c17f4bfd16a9ad4559b01be) |
| `storage_device_list.storage_devices.hpe_storage.password.clear_secret_info.url` | [storage_device_list.storage_devices.hpe_storage.password.clear_secret_info.url](resources--fleet--reference--group-003.md#canonical-7b3098fc5fb55607c8dc6a3dfc69076173fba7eea8a02e8aace90144408b1b52) |
| `storage_device_list.storage_devices.hpe_storage.storage_server_ip_address` | [storage_device_list.storage_devices.hpe_storage.storage_server_ip_address](resources--fleet--reference--group-003.md#canonical-dc749a13b470bacf66a132ad567d25ccbe9fc9a537b5f588bacca2fd426893d8) |
| `storage_device_list.storage_devices.hpe_storage.storage_server_name` | [storage_device_list.storage_devices.hpe_storage.storage_server_name](resources--fleet--reference--group-003.md#canonical-4ea4847e68b58a9a5bd990814a7406c0c9d10474518d322788909af3f01f4b50) |
| `storage_device_list.storage_devices.hpe_storage.username` | [storage_device_list.storage_devices.hpe_storage.username](resources--fleet--reference--group-003.md#canonical-9980cc76d5991a62bb04ce97cfb1c66cfbe570a7a0849dd3f31c1ca668ccf955) |
| `storage_device_list.storage_devices.netapp_trident` | [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-003.md#canonical-d1244af590bd3a828da9363c0e7307104eb60654f746a495b1d5ff0d41f243cf) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-003.md#canonical-7017ee263703d14d32a314dc98d2d2a4d76256bfaf7a0b131f1b82d543d9ba5e) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs](resources--fleet--reference--group-003.md#canonical-495339d4f23594f87edbb167a689e3a539392d92b5911f8d40bcea15dca87675) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs.prefixes` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs.prefixes](resources--fleet--reference--group-003.md#canonical-1a772160af81b4f0b30f34e3565f9809fbfc59a44ce02ecbd3fd29372d62a853) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_policy](resources--fleet--reference--group-003.md#canonical-6ddf9c8800245cdbbcf3de046b2e94651f9ed6e35d80dc38cfd0ded0168ec1ac) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.backend_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.backend_name](resources--fleet--reference--group-003.md#canonical-6436e772e5bba19236d81a6573898bd628b05e27da56b9b3e41f4b208bd5182f) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_certificate` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_certificate](resources--fleet--reference--group-003.md#canonical-c1dbebef51111a88a875f6552806c709452d47a4bd9f6d4d0fc30a7500422028) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](resources--fleet--reference--group-003.md#canonical-6c5d6d2deb3975fd415ff6fc44eee0366fc4b3bbd2efc2f5b49b43754f5e1870) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-9478402b411b488e47b6e3ac621057fa48811cd12a95354adb11127d20e4053b) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.decryption_provider](resources--fleet--reference--group-003.md#canonical-20ce298eb884a4feb435876497b722d66ac806cb512b376bc216b92a11d26fbd) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.location` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.location](resources--fleet--reference--group-003.md#canonical-43a46cee6a1e9333885da725ba98209f5d542383326fde9dcf9bcb7c3a5c8770) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.store_provider](resources--fleet--reference--group-003.md#canonical-d49bb71bd9e35122614c8265bf67bdb09286b95793f4223c836e46777045858a) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info](resources--fleet--reference--group-003.md#canonical-71fc60ac6566ea33417bc1e454d9c3d2950d49b0a7d369506473a1d0733a0fa8) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info.provider_ref](resources--fleet--reference--group-003.md#canonical-a6b1d09a8e954bda15df2fcb46920cf8d72db6d829a22fda1ef0920aa9533b88) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info.url` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info.url](resources--fleet--reference--group-003.md#canonical-f324c25c32bc70eb882621ffaa334826f5fe61c78ec9fbef1e9acc3db400b2dd) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.data_lif_dns_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.data_lif_dns_name](resources--fleet--reference--group-003.md#canonical-e4ab7fd866cbb9d3dfdcea2ef1e5d5b51d3351d1a199a41c318cd4a2044a80a6) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.data_lif_ip` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.data_lif_ip](resources--fleet--reference--group-003.md#canonical-05d05409462dcb4cb0e95415b5d29ae66ea3e7d3bdf4c49531de075bf2e09988) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.labels` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.labels](resources--fleet--reference--group-003.md#canonical-5de6d27c3902c0549bc36e7c72496da91dded1180870300d44d3a58f1e29b0ae) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.limit_aggregate_usage` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.limit_aggregate_usage](resources--fleet--reference--group-003.md#canonical-66188cbd50447b8a2fec8fcbfa04d10dba4a6ec28c3ada14d12b4eb3f80a71be) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.limit_volume_size` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.limit_volume_size](resources--fleet--reference--group-003.md#canonical-86c25388d04ef44756d88b14ae88bb00fd736c294daad60c5afe25b575ce25b3) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.management_lif_dns_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.management_lif_dns_name](resources--fleet--reference--group-003.md#canonical-f18601be4ed99b05c45a4868424b24619a48252d47351a2945d4edd0ca43e1b3) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.management_lif_ip` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.management_lif_ip](resources--fleet--reference--group-003.md#canonical-1c5e928d2d81a91e8c51b4e6a9e09affcbc2b20e67a26f5ac844185114f15512) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.nfs_mount_options` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.nfs_mount_options](resources--fleet--reference--group-003.md#canonical-baf668d18e983eb9b3ff9a99e7115467d3a6805d74a7065eddf7bd6a40ed8934) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](resources--fleet--reference--group-003.md#canonical-b4871c70d0666f4fae3d08f57ccb40db2ebc9fc3a0f1b41ef06b70a690d4f66d) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-786e8083a3ba0b4ef135a9b2da0f0f93f3f6f52372acaa950cd46d44180be476) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.decryption_provider](resources--fleet--reference--group-003.md#canonical-06ab38934c2a3b8a4dc00f3dfa72d71c459a1506f5403f627bc54ff416994ec9) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.location` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.location](resources--fleet--reference--group-003.md#canonical-814f9e03af60ec81187d139dce23bcdada2671fc628ca20fd2418310a5ff5501) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.store_provider](resources--fleet--reference--group-003.md#canonical-79c79c7c086b7f7d4ed1c2eb14c4b4fb5f1f7a6ce6f481b075de0af3f38b0166) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info](resources--fleet--reference--group-003.md#canonical-ef11a2559c71c14679bb611dc2405a189f295385c98f156113514cf51cc97e5e) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info.provider_ref](resources--fleet--reference--group-003.md#canonical-0bf129268d51881f8c86b42a2cf962184763aa2d4fa36e0b130208d15db3af8e) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info.url` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info.url](resources--fleet--reference--group-003.md#canonical-1f8ae8099ab48390f9cbda79e5629771f1ebf930fe8e490e374876ae7e389eef) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.region` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.region](resources--fleet--reference--group-003.md#canonical-fd9e975e8bb198e68303d66d5d010cda23facb7aa5dd383d46502a3600be435b) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](resources--fleet--reference--group-003.md#canonical-3bc3354cb89316f1cdf4087b9279db01682f095b0294b9dc68d5c348ecc25ff9) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.labels` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.labels](resources--fleet--reference--group-003.md#canonical-d739d33ce5de26cedcf856d9db420c9c307feed679e1c654ca527399dfea4b9d) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](resources--fleet--reference--group-003.md#canonical-0078a4bc2d30b9ada41644f0ee77f3a4d8cc7d62cb1b5dfd87c421180e3ee3e1) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.adaptive_qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.adaptive_qos_policy](resources--fleet--reference--group-003.md#canonical-5bdfa85ee14f6e2c96e4befec489622dac2f8ec5203add4a77b6bc8221efcc6d) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.encryption` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.encryption](resources--fleet--reference--group-003.md#canonical-3c8afc757c851bdca8c443a0780a0d3db71bf3513e3573fff9836760a5d7009e) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.export_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.export_policy](resources--fleet--reference--group-003.md#canonical-97f3ab7e026153e22e1b8df7783281af045386254f8c098bc3362fb7d6f5029d) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos](resources--fleet--reference--group-003.md#canonical-55c7da2ff60eb2b87b98edbe293caecaa9499d59d0376daa6fa77071174e1d0b) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.qos_policy](resources--fleet--reference--group-003.md#canonical-d0d4f7067c985cb7208547ad687dc00edcfaca96f4c592df68c3eba3ad6a8e2e) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.security_style` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.security_style](resources--fleet--reference--group-003.md#canonical-c3d2be7c7eea712a4886b3c2ef0d3ac11822dbf5224790814822fa76726ff751) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_dir` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_dir](resources--fleet--reference--group-003.md#canonical-78162d8797cb8db284a4173d0f6103c05cb75e40d2916b42863a629c2fd88ce7) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_policy](resources--fleet--reference--group-003.md#canonical-75e3ca6f19f526917fbeff30ac7aa3bcf82d7c78838514cc1a49ddfe6ea8ef1e) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_reserve](resources--fleet--reference--group-003.md#canonical-39479be8bfe2c67117c42f8a916553ca235058add4f33cdb8a69daad032b129c) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.space_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.space_reserve](resources--fleet--reference--group-003.md#canonical-49251069204ca897f242f90df21caeb1ddb7cedcd682f200828aee25e4b33e83) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.split_on_clone` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.split_on_clone](resources--fleet--reference--group-003.md#canonical-bb44bf1021846f3d4ab9a45e9f73b67eb9f0d5237cfb4940c649bba2672b3178) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.tiering_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.tiering_policy](resources--fleet--reference--group-003.md#canonical-d0274ac74b5af632a48cc20acc8315806a137dcaff71f301f645cb45fec305c0) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.unix_permissions` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.unix_permissions](resources--fleet--reference--group-003.md#canonical-ba2141390f3bb86382ec4273bc36a61070f0fb08d299ffa932d05dbebd5ff723) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.zone` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.zone](resources--fleet--reference--group-003.md#canonical-55597840738278d18e3b1fc5757cc26ebb45f5c5ce2a18619572943bc84dacd3) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage_driver_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage_driver_name](resources--fleet--reference--group-003.md#canonical-cce9c1c995c7023b9fac27fbe17b6bfbe72d4c5c85e14b1e8df96f7475358ffc) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage_prefix` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage_prefix](resources--fleet--reference--group-003.md#canonical-2ee5f0854eaee5c88671c396edb8860d49151ddc7e8c59357fbc6cd0fdccc1e7) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.svm` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.svm](resources--fleet--reference--group-003.md#canonical-90d3cceb48d1fcd337cf08487e65cbba5555ae2bf963257595b8c45682888d5f) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.trusted_ca_certificate` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.trusted_ca_certificate](resources--fleet--reference--group-003.md#canonical-f861a288c4ecb5fc7c068b15fba975a2bc4ef234d3330ec5c6120deebc9f86ac) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.username` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.username](resources--fleet--reference--group-003.md#canonical-c65c77bc7d93a2f58d1e870c5bd47b29e4dc6c7f489e80e1cc3aab99c02bf374) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults](resources--fleet--reference--group-003.md#canonical-2ed0708b574737bf6b846f5fb0423617980c50201854f3b3e708536803ef40c6) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.adaptive_qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.adaptive_qos_policy](resources--fleet--reference--group-003.md#canonical-71d16e5b5345242bd08954b722dacf909d8b4e7daa18b7724d015660f072bd2f) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.encryption` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.encryption](resources--fleet--reference--group-003.md#canonical-e9a16501e4ef6c3ef2ff6621829ea98fe6796343ae08ffc4769e0174bd08353a) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.export_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.export_policy](resources--fleet--reference--group-003.md#canonical-f73b932346462d67c632f16f5ae56f145cd7dbf80cc05e1c01fda3e1b6c216f2) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos](resources--fleet--reference--group-003.md#canonical-c26f8c1e2bf49c28644cd67d9b37b9d50514a19a26d131c76fe3776f2e5959d3) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.qos_policy](resources--fleet--reference--group-003.md#canonical-7b96c18ae430eb01199dd29f3da8c376646c6b4b28e21b3b6cb8b507ed296298) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.security_style` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.security_style](resources--fleet--reference--group-003.md#canonical-c0910bf8cce4982691e13cac7c45b8068325b05ba5541f021cc3df9c6b684e9a) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_dir` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_dir](resources--fleet--reference--group-003.md#canonical-65c4049cad020625b3dc362858271b160e5307374df986ccde12c600be8f2eda) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_policy](resources--fleet--reference--group-003.md#canonical-59d59982606260a385498353c3b2890acd7aea513cec45310dfd3f1d9b995eae) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_reserve](resources--fleet--reference--group-003.md#canonical-582e3534c9152938929d0e90fc92e1f4f659bb3ffede3a00e189fbdf23f89d84) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.space_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.space_reserve](resources--fleet--reference--group-003.md#canonical-3a81d20d82dca146f3eec85ba42f3b94c0b033cfb6bfdb324eb0dba8d1720df5) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.split_on_clone` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.split_on_clone](resources--fleet--reference--group-003.md#canonical-5ee7171993d7991c4e165fc74295bdaf3268e124a4de52e080b954fd356a66ee) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.tiering_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.tiering_policy](resources--fleet--reference--group-003.md#canonical-719bf9ba926139cae109e75b987f21449e438bd1908bd5a5db15559aa69f3e25) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.unix_permissions` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.unix_permissions](resources--fleet--reference--group-003.md#canonical-2b56104202699aff080896a936c0a1f7f134cfa5e6120e5ee070dc27c1e96cfc) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-fb92cb2c840473dc7a79c5078151f9a867d5c7e6dce621d73b671f471bb0a0f3) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_certificate` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_certificate](resources--fleet--reference--group-003.md#canonical-7fa20fa352b4183980a6e68a2873244ba0b41c902e03655ebb7797849e7cf164) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](resources--fleet--reference--group-003.md#canonical-7f2e3bd7a8359ee86d663c0b1a1c4b83688be9dd53c0f7e052b121828f764703) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-4e465f8f3be19a1f0c0f6fa0b2b8f1ad65e369eb1afe82195812ddeaa56d54db) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.decryption_provider](resources--fleet--reference--group-003.md#canonical-02f0003f90ee1f59c2e518f7d449b83f08037cdce226fbaf780e50496052a933) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.location` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.location](resources--fleet--reference--group-003.md#canonical-8e9f4aff07d5bad0d5792ce303eb5e0e8ce5631078e1cee0bdd1937dfc1f8699) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.store_provider](resources--fleet--reference--group-003.md#canonical-c5688594c835dec8220106ad2f64043c1fd0befb9fbab8843d3137ab6d829f38) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info](resources--fleet--reference--group-003.md#canonical-9593d9c94c1a79871c91527b0567433933132fec10caa88fd2657b2d5002e4e4) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info.provider_ref](resources--fleet--reference--group-003.md#canonical-dd4e4f63a832ceba591372f086c85231c6af39f508f49242278a211b92c92c46) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info.url` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info.url](resources--fleet--reference--group-003.md#canonical-7ee4b5cab99258e737dad9a3acfdd80f73fc7e27d7202999b7b1c031729f7f65) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.data_lif_dns_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.data_lif_dns_name](resources--fleet--reference--group-003.md#canonical-8254d81ef56c539614e8c2750af9bec705f79769ad4836dd9ab26bc614166499) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.data_lif_ip` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.data_lif_ip](resources--fleet--reference--group-003.md#canonical-86547b7790f5008a789f6f6dea5fa26e94c20ad17ff58dbdf7b58d207a045e70) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.igroup_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.igroup_name](resources--fleet--reference--group-003.md#canonical-8617a7295b347085807f5a519b54b8e97bf81bc2ab104396c9003d0095fc21ee) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.labels` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.labels](resources--fleet--reference--group-003.md#canonical-6f8cc426379bc7acad251cf1b49d66c50f7c9451b05a60db03616c9f11d6d76a) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.limit_aggregate_usage` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.limit_aggregate_usage](resources--fleet--reference--group-003.md#canonical-1de049e46800f0f2b68fe0c3343f809b40047020ec3ab60acfe7aa8747c164e0) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.limit_volume_size` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.limit_volume_size](resources--fleet--reference--group-003.md#canonical-da961ac3c9ce86e43dc46d10d35241539fa2ade770f76a11b575213bbc0baa19) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.management_lif_dns_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.management_lif_dns_name](resources--fleet--reference--group-003.md#canonical-2748be39f761fb2b269d35b09315c8f88a3c641edae608b2a971aea83bcb7f02) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.management_lif_ip` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.management_lif_ip](resources--fleet--reference--group-003.md#canonical-f9cbe0ed882101add21c6839e03ef537a2a239c50de3ffcba20545f0a4aebd14) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap](resources--fleet--reference--group-003.md#canonical-b6b53c6ab31e2a2f864f68f532c2b9d67a55f3fe346452af7d49a86b18dc94f9) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](resources--fleet--reference--group-003.md#canonical-35a380af8d747fd76c81e10718a99d53e6df02e338286d91b8bc474b97e59f9b) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-03ea32f0c68b2cc345ea3993f4ed3eb317676c9dd7d1153c74011e2f93c0af69) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.decryption_provider](resources--fleet--reference--group-003.md#canonical-cce8b8d4466ec1b87b12b082f5c3176f34de2dba6958cf9b08d6254508e6d2f1) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.location` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.location](resources--fleet--reference--group-003.md#canonical-db0d8f6e8fc6b5563931f5cdad986e8cc147245b533e4b2386b3b6ec348ffb2b) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.store_provider](resources--fleet--reference--group-003.md#canonical-7075762e2a7a4336cb1b8f31d545cc114a899287e7394f69058fcad70cc666f6) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info](resources--fleet--reference--group-003.md#canonical-ed2f6d04154a908e994c1980bf71e60e6b04e76bbcb28be9c9004635833298d1) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info.provider_ref](resources--fleet--reference--group-003.md#canonical-1ca35c55e10544bbf5c2a5f2ba780ad8bc546d78ec0f958671eb7a67adfee2b3) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info.url` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info.url](resources--fleet--reference--group-003.md#canonical-01b13ae1d3de6bb7f3891fcd7ef7e0d8ad46ff31be4b1883d7ac7f8d3b0c5cb7) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.region` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.region](resources--fleet--reference--group-003.md#canonical-d7271893588e1287d47792beb9115862286e2b1cdf4895eabfbc70fdf98a6cd4) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](resources--fleet--reference--group-003.md#canonical-25394e5df60f715f6ff4cc7af635f7ed27b365140e19fce2b931bcd1bfb5d3b1) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.labels` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.labels](resources--fleet--reference--group-003.md#canonical-4deecf41515815c9c3711721e649fd54e5f47901ddb77a35c5ec0b208dbcb522) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults](resources--fleet--reference--group-003.md#canonical-b129aef4f91fbb18b3d0bac918e60bdc8163204d8417f315879b6d66c0dc02f1) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.adaptive_qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.adaptive_qos_policy](resources--fleet--reference--group-003.md#canonical-4a982acda9588c6473c125c048cda204d00d21844dbf1777543ca2839cdd1ce1) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.encryption` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.encryption](resources--fleet--reference--group-003.md#canonical-3cf0aea4f1440c8e9e2e43d12b54e3314cf27d91ae277bf8bc4892d4c9acb76a) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.export_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.export_policy](resources--fleet--reference--group-003.md#canonical-a725885683457b3d339d2826c2ff7a475b9ad8869a88b6acbd5842e564b174dd) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos](resources--fleet--reference--group-003.md#canonical-5f0e9aec333a5358cac577ba8cbf0408f367836d7bbf56051b2ed70c55b8d847) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.qos_policy](resources--fleet--reference--group-003.md#canonical-c7d2db1427c1f70a8517f88d1b1de7c66bfca65a64c3cbd13b5ca330b6177cf9) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.security_style` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.security_style](resources--fleet--reference--group-003.md#canonical-7446e92740aaa75e1f67b7330fbc567b247ca84ba1eb2ba2b3742465a641b8f2) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_dir` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_dir](resources--fleet--reference--group-003.md#canonical-7372b58488d6c3c6448329b55e76b5c4400b59432cffd2ef4493925df936925c) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_policy](resources--fleet--reference--group-003.md#canonical-dad9603d873884b6f16355693b6a70803d7ae69b83abca5dbfdf59ded4e5a2fb) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_reserve](resources--fleet--reference--group-003.md#canonical-54480bd16b1a91306ade296f4de31003e7f772b3da8778414b5b3222a942113e) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.space_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.space_reserve](resources--fleet--reference--group-003.md#canonical-d64ef82c379dc843a26e9c5dd9ac498de1bff1c39bae62e3155f63f146853dec) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.split_on_clone` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.split_on_clone](resources--fleet--reference--group-003.md#canonical-53c423e68aab45909dd1a9f14ee26b9e28a2ad6b1e596de59e27494563ab2dd1) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.tiering_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.tiering_policy](resources--fleet--reference--group-003.md#canonical-78a7bad3541bc51b4b71e0ffab4a6aad9012a0e17e9e3aa4622ce63e633cc32c) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.unix_permissions` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.unix_permissions](resources--fleet--reference--group-003.md#canonical-46240e6f79210a9881fdf8f7edfa921470de22481193c329bbd9059995f86111) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.zone` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.zone](resources--fleet--reference--group-003.md#canonical-abc1c05b48a5c1bca98d7d65943e57314ae6cc444e4311420c9d392168314b7a) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage_driver_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage_driver_name](resources--fleet--reference--group-003.md#canonical-72a044adb7d6c63616408f500d27be8f84509949b2233086fcef2a56709a8f8a) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage_prefix` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage_prefix](resources--fleet--reference--group-003.md#canonical-a10bcf3f8241c0625e51f0d6f74a166942f834c7d507e72ff9b20547d0ce354d) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.svm` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.svm](resources--fleet--reference--group-003.md#canonical-569e8f93bd32f899e9d3a332a33f545f93a2b1adeabaa037ac8e50d68aa0833e) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.trusted_ca_certificate` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.trusted_ca_certificate](resources--fleet--reference--group-003.md#canonical-4c6a69734f691276469f0d8b0109a8e2f06e2c5dd00c57c7f9ddd4348a1e7d4e) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--fleet--reference--group-003.md#canonical-5f3abbadb3a45112ee9f9ccf03e1c9596744c75fd5d2f41e2237d2b45eb18cf3) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](resources--fleet--reference--group-003.md#canonical-c392202bd46d57b924068fee5808a2b08dc2b54534973eecffd971bdf09182cc) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-97b9c2182b9303b5dea6a0b31ee86c201ded45426172a572e3fbeb0af813a60c) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.decryption_provider](resources--fleet--reference--group-003.md#canonical-ce943248688dc31e2d61a516cd1b230fad3d38386553ac2af8df905c94716b1e) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.location` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.location](resources--fleet--reference--group-003.md#canonical-3f36e2c8b0a7615df5d3abfcf76f08a068e988fae0e59ff8e616d4280c2850b0) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.store_provider](resources--fleet--reference--group-003.md#canonical-5646bd9b3540df70d0c31e1ab9de7da2c76f7066a932a429dfb9bbf2a002f5db) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info](resources--fleet--reference--group-003.md#canonical-99185e34e3cbae7af8ec088ae4eeb00af3d7f0da2c1fd5f5fba03527d33a3250) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info.provider_ref](resources--fleet--reference--group-003.md#canonical-d3129e5024257b6ab33e8da601bfdc09b2016dd0f8ef41a19781aa9432404109) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info.url` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info.url](resources--fleet--reference--group-003.md#canonical-4df249798204d772b3ec08f27eca44cc9991011422be846643a53995189787b7) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](resources--fleet--reference--group-004.md#canonical-a39974bf30becc8ec8833c7855629f8259f1203af39f57135e0b5b8734cd4e31) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info](resources--fleet--reference--group-004.md#canonical-18ab81b288ba6c03c8edcfb65dadca79487d3a34146078da4d20eaaa508e1cef) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.decryption_provider](resources--fleet--reference--group-004.md#canonical-81c0a113b15d2c8a631f53ad9206e40fec006bfd909034f19414f06dabd75286) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.location` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.location](resources--fleet--reference--group-004.md#canonical-212084729941c95fa650aa548e6d44bdac9870a0371512fa40fb1e01fb2c294a) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.store_provider](resources--fleet--reference--group-004.md#canonical-9fc00afd4c97525d240f2b58866e8bef87cdea7d0c62be96b3fdf2639ce1aa07) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info](resources--fleet--reference--group-004.md#canonical-29394c83a341b03a174b410ace8310ee5284ccbb760ad280bee0d5ec65dd6aa9) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info.provider_ref](resources--fleet--reference--group-004.md#canonical-6a2cbfbf5f171e24cb9de47e8c82ca8eaaeece82d7341b3bc1e1403d10575d97) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info.url` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info.url](resources--fleet--reference--group-004.md#canonical-2dce8fb9e609e156930fcb37368c256ecc6fe81e75b80c02d63283cda035dd52) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_username` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_username](resources--fleet--reference--group-003.md#canonical-931f1b9029367784dba11298a35ceea356dde33ad60fd80d263795f0d2e4d147) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_username` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_username](resources--fleet--reference--group-003.md#canonical-90e4b8d2ac29c597ca36ebfaf8873f1f812e647a178bbc57d9bce68ad4f20587) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.username` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.username](resources--fleet--reference--group-003.md#canonical-4ae56b9c3509cba47c6a812a7f5473eed81b36f8dde85d01c9f00a08caef8470) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults](resources--fleet--reference--group-004.md#canonical-a0ae7a818775106128a430651d7b71169d3f0335e49478f4293e5ff64d22cf28) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.adaptive_qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.adaptive_qos_policy](resources--fleet--reference--group-004.md#canonical-6f3e941133826af29e60126b6d82e0d67e52830a559be2a2f9aef3d282e2e825) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.encryption` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.encryption](resources--fleet--reference--group-004.md#canonical-f8e0658e0c5b0b0c96e476d1bfe60b43abab5e08a63ff22671eee845d8818c61) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.export_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.export_policy](resources--fleet--reference--group-004.md#canonical-5268e3f39f6d9fb6099f28887a231a207f6a0228e60f20c6481b2751470501e0) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos](resources--fleet--reference--group-004.md#canonical-43a778e31d02fc7b7853163a83239f1532cdf0007b1a94428dba6f7d91298a56) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.qos_policy](resources--fleet--reference--group-004.md#canonical-3a16b28a9ab6eb4a976857070372ecaedd1869bfd67125b44004e45565194399) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.security_style` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.security_style](resources--fleet--reference--group-004.md#canonical-b4fed030be70bde1648427dfffd91dac77ee0046c10ec256365dac68b6408ccc) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_dir` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_dir](resources--fleet--reference--group-004.md#canonical-79dd8006a746350204fcf269a962e3a2266c7bc11f3d7a632d9ce5a190b3665c) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_policy](resources--fleet--reference--group-004.md#canonical-d00234cd3d74a0fbcc9cb06b652ec7e02fa2f866ad3738460077b92acbbfccb3) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_reserve](resources--fleet--reference--group-004.md#canonical-47630ab2deb41e18e6911b72189ddb3bb06121b2638f29b27a2723abf1d9d1e2) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.space_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.space_reserve](resources--fleet--reference--group-004.md#canonical-3f63d6571b89c6a295c887830d911ab39f717ea900086dc6958615ab312a01a8) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.split_on_clone` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.split_on_clone](resources--fleet--reference--group-004.md#canonical-6d109a65cd5e36252a8a0d51e4f8c9dd45370aed4a7e80d2456f2e4e31d8006f) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.tiering_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.tiering_policy](resources--fleet--reference--group-004.md#canonical-e8c00df9ebd2cf6fe710d3bebdfa38f3d7c9e62da0591f0ed8a71a086330058d) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.unix_permissions` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.unix_permissions](resources--fleet--reference--group-004.md#canonical-24930dec69be98f027ef93245f6609ec27b66216db6827f12aecdf545a4be826) |
| `storage_device_list.storage_devices.pure_service_orchestrator` | [storage_device_list.storage_devices.pure_service_orchestrator](resources--fleet--reference--group-004.md#canonical-14de148481b302b8ff4bd9c3e2ea76e83cd523af856c343bdbdaa7a1122edd40) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--fleet--reference--group-004.md#canonical-6a87b82b39e53b9691486994ac9509b2d885d19f1a6f869d5837cb2baa77f482) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](resources--fleet--reference--group-004.md#canonical-4954b763738d6bf182a30ee330b1344a862fbaa33cabbe61d18ebad90473282c) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_fs_opt` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_fs_opt](resources--fleet--reference--group-004.md#canonical-49904b0bd4f8add33c08e33c61c7e633914dae1765638419fa1a0b28d38edd97) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_fs_type` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_fs_type](resources--fleet--reference--group-004.md#canonical-c94642ce0c5acf73113766442063095192991ec052ce7c4cff1cf1330d8011f1) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_mount_opts` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_mount_opts](resources--fleet--reference--group-004.md#canonical-cd2fe82e141c4d02025cb196c99f2ec1738d259954f7458b45473df348957908) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.disable_preempt_attachments` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.disable_preempt_attachments](resources--fleet--reference--group-004.md#canonical-5ec703ee016eac4735ddeb45e45a5b4ea5696b59cab613e8157e0589987137bb) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](resources--fleet--reference--group-004.md#canonical-2ae04a259b087ce3ffc5ab742cc03e4122ca15269f83e17b9fd4ace9d327a2ff) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](resources--fleet--reference--group-004.md#canonical-88d0198acab51deb0930663b32b7a8b520161eb24c830143b15b32c5aa21874d) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info](resources--fleet--reference--group-004.md#canonical-67ca5ec1303c0309b7e51292cb41a84d2713e45b07e07aa24930daefe140fbe3) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.decryption_provider](resources--fleet--reference--group-004.md#canonical-c46d2b01b5d9aa5efcef68b767fda4d53320df3607a45d31eec933ef93154a86) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.location` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.location](resources--fleet--reference--group-004.md#canonical-03ce3af0153bddcb55d7a1269438f9ed0cabf2b9e4a914a904b7b1c13f701663) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.store_provider](resources--fleet--reference--group-004.md#canonical-8a494978431ce8ac93b754bdf2f64788fdfb903c2af5c22923b33fbd576c31cb) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info](resources--fleet--reference--group-004.md#canonical-32c71a2d8272cf6c3eb3c8ec7595567f624708a72419c458c68b86c57a3c3ca6) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info.provider_ref](resources--fleet--reference--group-004.md#canonical-60d543adfe2d3d23a831cbd58ec63b0436f4d58d8c41fb24bd38850683364248) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info.url` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info.url](resources--fleet--reference--group-004.md#canonical-9ec89b0dd37254ed8f560f7aa0429725bfe072c569d4fa05be134d55da721050) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.labels` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.labels](resources--fleet--reference--group-004.md#canonical-1a1a5758c855c812678cec1f3d6931336b6b871f368a7347b87b8182cd963606) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.mgmt_dns_name` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.mgmt_dns_name](resources--fleet--reference--group-004.md#canonical-4392a988121b3833a2dbec03fd8a956ec60cb07f16d3ea73e7c200e7e8ec790c) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.mgmt_ip` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.mgmt_ip](resources--fleet--reference--group-004.md#canonical-bd524ada8fbb2e631b733663cf8571f37c4c57456e00b9a8aa6b7cf639a31cc3) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.iscsi_login_timeout` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.iscsi_login_timeout](resources--fleet--reference--group-004.md#canonical-b9d936c399ab73eec9794988a75bc62e9660811612733ec614dbf64dd43d3193) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.san_type` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.san_type](resources--fleet--reference--group-004.md#canonical-d7136f046bd0bf9ef154236d2a9bdaaaad1d81092c208fc7e8f1b8e8ee2a335f) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](resources--fleet--reference--group-004.md#canonical-0b6b81ac3c773a1bd3093f659fb7acb1d467be05be7a26c8cfc2b4227feb4031) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.enable_snapshot_directory` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.enable_snapshot_directory](resources--fleet--reference--group-004.md#canonical-07e1125b38da5a005eca81559939c4346313a46fad122cdd475f006ed1982625) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.export_rules` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.export_rules](resources--fleet--reference--group-004.md#canonical-88a54b36c8d0ba51d9b3541429721aa762d9ef78aae5b45c4ef9b2fca6811733) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](resources--fleet--reference--group-004.md#canonical-dc1ee5201b87c4b15b21a74a6bbdc1bfc1ecc67b3f92777ca85a1edbbd5596c1) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](resources--fleet--reference--group-004.md#canonical-43591d5f0a5b8fd2150212a905ce02ee4f8e3c0796bceea275433445ba57c7fa) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info](resources--fleet--reference--group-004.md#canonical-25c12c98320818f0e4a74d1f830ccd2b4df08bd164723de0d2510aa49fc2c143) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.decryption_provider](resources--fleet--reference--group-004.md#canonical-ff58e74acade22a85f3fba73affaebced132f2c6915c706a0916bdb8050ab93b) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.location` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.location](resources--fleet--reference--group-004.md#canonical-f844d7f75552348e93c8557243c25d0c5096a0f5abcd4a2607ebc804e9526d0d) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.store_provider](resources--fleet--reference--group-004.md#canonical-54cd2031cb3fe10681c956a99c8b116c139faa65b7fd9f66d7acfc1095dccffe) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info](resources--fleet--reference--group-004.md#canonical-87d9bfaa4310edd73355f580cb95d67b2569c005c026795ff866289684c99e5e) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info.provider_ref](resources--fleet--reference--group-004.md#canonical-c66d1d1e30d1679d1cd929683b47f392df0217ac11126e045d8bffb640fb330a) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info.url` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info.url](resources--fleet--reference--group-004.md#canonical-4cd956c174d3ef413a45a4f4830d9fb0bd92d0c76d536bcd749b69941507339a) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.labels` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.labels](resources--fleet--reference--group-004.md#canonical-b818b07ec925b2d3a284d42c32f91940cb9cac18c3e1f2a9dac008fb17ea72c3) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.mgmt_dns_name` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.mgmt_dns_name](resources--fleet--reference--group-004.md#canonical-71ff5d40182896e202b6f1f46ef85d8ba38e28a363dceed52deb3b5fb2eac2d2) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.mgmt_ip` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.mgmt_ip](resources--fleet--reference--group-004.md#canonical-5caaec6b44903965ec8dc87fbd2be3e5375f7f94955fa9c6c75135b8f8774e96) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.nfs_endpoint_dns_name` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.nfs_endpoint_dns_name](resources--fleet--reference--group-004.md#canonical-94d9d31771c06151cb6857c6c87e83a685cd6786434fab3b62cf33b54db872dc) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.nfs_endpoint_ip` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.nfs_endpoint_ip](resources--fleet--reference--group-004.md#canonical-e81417ce3b647aa783db0f6ef60d8781d37fc069828b6d6a1807b9eeef40cd4e) |
| `storage_device_list.storage_devices.pure_service_orchestrator.cluster_id` | [storage_device_list.storage_devices.pure_service_orchestrator.cluster_id](resources--fleet--reference--group-004.md#canonical-7d50f7ebb09768647e049bf21aed47b8d11fa08641612c6f586fdfc7f254e477) |
| `storage_device_list.storage_devices.pure_service_orchestrator.enable_storage_topology` | [storage_device_list.storage_devices.pure_service_orchestrator.enable_storage_topology](resources--fleet--reference--group-004.md#canonical-559381884f22fe7b5d55da6ddd58ff7679ea9ffb8f26e97857ac9b3aef14e3cd) |
| `storage_device_list.storage_devices.pure_service_orchestrator.enable_strict_topology` | [storage_device_list.storage_devices.pure_service_orchestrator.enable_strict_topology](resources--fleet--reference--group-004.md#canonical-57c44d71a25faa26758e8ae2266b8ad8223d0ee09516ed57362b63f3ffd4e8fe) |
| `storage_device_list.storage_devices.storage_device` | [storage_device_list.storage_devices.storage_device](resources--fleet--reference--group-002.md#canonical-730922ec66266d5fc6d8f38792a91cfc68c0d2f22772041f3853aab16e90f9f4) |
| `storage_interface_list` | [storage_interface_list](resources--fleet--reference--group-004.md#canonical-abde75c3848a86f8fac961cec1ea5eb4baf07f3d441f29cca9d305a37ab7c837) |
| `storage_interface_list.interfaces` | [storage_interface_list.interfaces](resources--fleet--reference--group-004.md#canonical-eacb7474009f44727f047a6e40e448e0acfcd807f81fadad0d309364ddbe3ddd) |
| `storage_interface_list.interfaces.name` | [storage_interface_list.interfaces.name](resources--fleet--reference--group-004.md#canonical-0113450160775584addd16ae3a9900fadbe0b8dbf12a8727cfeb5e22d2ba0767) |
| `storage_interface_list.interfaces.namespace` | [storage_interface_list.interfaces.namespace](resources--fleet--reference--group-004.md#canonical-fddbda06f3d00de103ea8c6f670aafe40b2cb2459ca0621efbd145f5c6576437) |
| `storage_interface_list.interfaces.tenant` | [storage_interface_list.interfaces.tenant](resources--fleet--reference--group-004.md#canonical-56de481721b6336735ed470575f266eb8b527007ea2d25ddcadaca9f03a11f11) |
| `storage_static_routes` | [storage_static_routes](resources--fleet--reference--group-004.md#canonical-f0a29d64ab6d2465e8f1351c8716c9d7320d650e2542ff5d009cfdd3842125a5) |
| `storage_static_routes.storage_routes` | [storage_static_routes.storage_routes](resources--fleet--reference--group-004.md#canonical-574e362c139e5cae91672eee5a566dedb5eb580d59c5c76f5c7f2e9ff5cad6a4) |
| `storage_static_routes.storage_routes.attrs` | [storage_static_routes.storage_routes.attrs](resources--fleet--reference--group-004.md#canonical-8f731df96ea3edd9f55a6370a5ecdf2b6024385df47f0af498853b78ca356d03) |
| `storage_static_routes.storage_routes.labels` | [storage_static_routes.storage_routes.labels](resources--fleet--reference--group-004.md#canonical-e083c5ef23133ab0639f783b0284383b92c9ce1de006ff08b3ecf6228e1c326e) |
| `storage_static_routes.storage_routes.nexthop` | [storage_static_routes.storage_routes.nexthop](resources--fleet--reference--group-004.md#canonical-5f8f4abe2132ceb6c4c32fa5cc239550409fdb6fe75dc72ea78824997025665c) |
| `storage_static_routes.storage_routes.nexthop.interface` | [storage_static_routes.storage_routes.nexthop.interface](resources--fleet--reference--group-004.md#canonical-9168ce2a452993aba94d26ea38e706b62555f6fd226f3486dfb7675d24fcc3a0) |
| `storage_static_routes.storage_routes.nexthop.interface.kind` | [storage_static_routes.storage_routes.nexthop.interface.kind](resources--fleet--reference--group-004.md#canonical-57f6d7abbc24c68d1d4586f36c029301014daf9f08cf5259c857dea60665ccb3) |
| `storage_static_routes.storage_routes.nexthop.interface.name` | [storage_static_routes.storage_routes.nexthop.interface.name](resources--fleet--reference--group-004.md#canonical-7fd9d1e5378c8680f14b86f11e3a123f2b97af10c22ce30ca06e53e7792de9eb) |
| `storage_static_routes.storage_routes.nexthop.interface.namespace` | [storage_static_routes.storage_routes.nexthop.interface.namespace](resources--fleet--reference--group-004.md#canonical-586530fc599c7ae47839e40a13749fbd0a282949aa3d281a8c634ca97a58bbbb) |
| `storage_static_routes.storage_routes.nexthop.interface.tenant` | [storage_static_routes.storage_routes.nexthop.interface.tenant](resources--fleet--reference--group-004.md#canonical-44a46b73d5930f28ec632e49fd156e273d857f2697d265b167ee54f274672208) |
| `storage_static_routes.storage_routes.nexthop.interface.uid` | [storage_static_routes.storage_routes.nexthop.interface.uid](resources--fleet--reference--group-004.md#canonical-47d4145268dd24a155c3fa722473475bad85f15d5246e7432994e3a6082c0db1) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address` | [storage_static_routes.storage_routes.nexthop.nexthop_address](resources--fleet--reference--group-004.md#canonical-48a2b8a83c7486503e7ea76197f1136b20c66d1e5c2f49dd56089c2350078ad5) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack` | [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack](resources--fleet--reference--group-004.md#canonical-f6680bd27239df405ecff16ebc8851a1f79effaa9892839d0631ce32077db43f) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4` | [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4](resources--fleet--reference--group-004.md#canonical-e84d27fd2b2216a52deab656aeff38659b13ded354ee378f5544ddd45622b659) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4.addr` | [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4.addr](resources--fleet--reference--group-004.md#canonical-9d1ea7d5004e23947618b49a459415f207059fdfd9f48d92c3ff2c55f806a004) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6` | [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6](resources--fleet--reference--group-004.md#canonical-b5ea068053ad5c215c8be716302392d062fd7b06e5af00aafa0375e54916ef76) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6.addr` | [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6.addr](resources--fleet--reference--group-004.md#canonical-7057fc6efb84b3d5f8e0b4bbe32e5518c04ca7e4161facefb6f805925b771c95) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4` | [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4](resources--fleet--reference--group-004.md#canonical-9a3ed42f51407959a0ae52c84e3e2e51372122d9ad5dc14963f8226a24ce94c2) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4.addr` | [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4.addr](resources--fleet--reference--group-004.md#canonical-46cf39ad419a749c978f1b9ec84a18f7e6c0b555636c395ca93d23b51c475d5d) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6` | [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6](resources--fleet--reference--group-004.md#canonical-b8fb6d028fb0d22719bc9f30e8c04fce5670093869de686e388d533b591bf92f) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6.addr` | [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6.addr](resources--fleet--reference--group-004.md#canonical-3d9de131dffe24a5c324c8982433cf80afca137b604602de1fe844261768c04e) |
| `storage_static_routes.storage_routes.nexthop.type` | [storage_static_routes.storage_routes.nexthop.type](resources--fleet--reference--group-004.md#canonical-f61e8a92807082d0682a5525000d37c0f93a00809ae7093b9624b9f736ca3dbe) |
| `storage_static_routes.storage_routes.subnets` | [storage_static_routes.storage_routes.subnets](resources--fleet--reference--group-004.md#canonical-947210d8f776f1203d39e5ff072ca03fde339a8899643dbacb20d54d4e032d86) |
| `storage_static_routes.storage_routes.subnets.ipv4` | [storage_static_routes.storage_routes.subnets.ipv4](resources--fleet--reference--group-004.md#canonical-c5cbb38776c787e42ebd0d0955523528ba8ef1dba1d08a788770d65caf85b0ff) |
| `storage_static_routes.storage_routes.subnets.ipv4.plen` | [storage_static_routes.storage_routes.subnets.ipv4.plen](resources--fleet--reference--group-004.md#canonical-4cc866de173968d0a81bbc97659da2c73e8bedf19f0c9cb2101dbef7411c51c5) |
| `storage_static_routes.storage_routes.subnets.ipv4.prefix` | [storage_static_routes.storage_routes.subnets.ipv4.prefix](resources--fleet--reference--group-004.md#canonical-e786b9d6e94cc1718628dcc125c8799356f804b63d43b429ba1bc16b544d72dc) |
| `storage_static_routes.storage_routes.subnets.ipv6` | [storage_static_routes.storage_routes.subnets.ipv6](resources--fleet--reference--group-004.md#canonical-aa3d29551011a9bbbd026029854e5978496cf48c0c2020eb508cc0fba107c33d) |
| `storage_static_routes.storage_routes.subnets.ipv6.plen` | [storage_static_routes.storage_routes.subnets.ipv6.plen](resources--fleet--reference--group-004.md#canonical-62e4870346b5dae7ce3c0aaa4a486f51468d3f32271fbf4e7da05d50d58d7fa9) |
| `storage_static_routes.storage_routes.subnets.ipv6.prefix` | [storage_static_routes.storage_routes.subnets.ipv6.prefix](resources--fleet--reference--group-004.md#canonical-d39bdacde055b165e79d0ccc8298d77cf3a102bcf4d3369e7d67decf55317347) |
| `timeouts` | [timeouts](resources--fleet--reference--group-004.md#canonical-11ff33128bee74e8ef1fac4bfee63378effcbb132cdf821d3eb28aa7f1a7ba8d) |
| `timeouts.create` | [timeouts.create](resources--fleet--reference--group-004.md#canonical-fadb5a7b2cbf436d4d3c795a0b9600dd1e44c59da22af3acff2b82ab0215471a) |
| `timeouts.delete` | [timeouts.delete](resources--fleet--reference--group-004.md#canonical-e37c04db6a43a74f4056725e9915fc988e58779748d24c93b858140abb176d4e) |
| `timeouts.read` | [timeouts.read](resources--fleet--reference--group-004.md#canonical-48b318d5034cc2374beb8c8383cd7268315bc9c27319b19723f4199455c828b2) |
| `timeouts.update` | [timeouts.update](resources--fleet--reference--group-004.md#canonical-a584ae4fbc1b7398328321881f4b0bed5758decca954d39ad1ed01d120f0158c) |
| `usb_policy` | [usb_policy](resources--fleet--reference--group-004.md#canonical-5a526f17d1768fe3878cc166b4741dfc79b3df781f6f88d652bf81ec9a8992ce) |
| `usb_policy.name` | [usb_policy.name](resources--fleet--reference--group-004.md#canonical-243776d44787bbe8f928baa09df93bbba3395712f73bff8c00f82d6f38fdf068) |
| `usb_policy.namespace` | [usb_policy.namespace](resources--fleet--reference--group-004.md#canonical-808d59665a84fec1f43fdf13f8f0a5de302df7551d8776bb9715381b42f486a6) |
| `usb_policy.tenant` | [usb_policy.tenant](resources--fleet--reference--group-004.md#canonical-7a1d5cf0e6c87838f35cdd5ea8976470dc44f95c84bad2d3ec576dbaab564e22) |
| `volterra_software_version` | [volterra_software_version](resources--fleet--reference--group-001.md#canonical-56d6fef330e305019d1650aa97c37e935e7e3e68f2df88243fb4806062b617f2) |

<a id="canonical-03a88d336bd2e72cfaa5f9310b8f94f6829c11a0de1d970be3f3704b5fa1335c"></a>

## Next pages — Property reference / 49d5ced34f3e / 16

- [allow_all_usb](resources--fleet--reference--group-001.md#canonical-2f22e2a90601571efd4c85045edeecb1f252e897cfa44bf672d9267183518fdb)
- [blocked_services](resources--fleet--reference--group-001.md#canonical-6a97758cec25e40f0955dfad6ad390b3ed4989ec29accf52af9593619343596f)
- [bond_device_list](resources--fleet--reference--group-002.md#canonical-51c90dd8a67fb1ce21ae1dec5a8ff48b84edb403be963bbab7912f7718909b3d)
- [dc_cluster_group](resources--fleet--reference--group-002.md#canonical-a8156c5491eb298ebda885fbff5c7c44a0913dddc39143494557eabceea08563)
- [dc_cluster_group_inside](resources--fleet--reference--group-002.md#canonical-4ae9b77ddff6d8fbe7167465347d3472c315520e1693b3c6ea03e3c8130f508f)
- [default_config](resources--fleet--reference--group-002.md#canonical-f408aab60f432ad21626ed326f86b47229253e65115eb2eac015d6e7593204f7)
- [default_sriov_interface](resources--fleet--reference--group-002.md#canonical-39e0721784456474c166b87766c385cdafa1285483adb504e75a5e9fe8aa9371)
- [default_storage_class](resources--fleet--reference--group-002.md#canonical-b519b74545b0c70bd2e86bb26ab51f7b6e4d56ec062b14ef4390106d26383010)
- [deny_all_usb](resources--fleet--reference--group-002.md#canonical-8643889543f2b9b1ce29e52d76b7057c7e41e19a74827b7ff74eab7619639738)
- [device_list](resources--fleet--reference--group-002.md#canonical-d8c812327bc50c891fb3f0faee847413de41c529225e52728413f5a01e42a033)
- [disable_gpu](resources--fleet--reference--group-002.md#canonical-1b3a5bfd1dfadfe6e0885816410305d81a80cdad6116c269273c917bce7e8343)
- [disable_log_anonymization](resources--fleet--reference--group-002.md#canonical-0cdf79b581c68d230805be43993a64ee5dcfe96ac81609eb67217b2410e598a0)
- [disable_vm](resources--fleet--reference--group-002.md#canonical-3904874165a1ef727a4f8ff00d8b5a126ccfc4977500dc375a5c5c689ac1c621)
- [enable_gpu](resources--fleet--reference--group-002.md#canonical-0205c6ba7929d8e95a82a34b489cee522a02ce75dfbbbea61460c97cdeb47934)
- [enable_log_anonymization](resources--fleet--reference--group-002.md#canonical-70701e28574e291f7bfa0aaa45bad6d40165ac73a4121c875f4b6fea4ece932c)
- [enable_vgpu](resources--fleet--reference--group-002.md#canonical-0ba113417dd322922714de6fd64f4975669552d3e9b96ef7fe2a01f57fde4284)
- [enable_vm](resources--fleet--reference--group-002.md#canonical-3a7d7361e54b7a7f248f5fa61458fd7e181f92f3712c92302c4da7638c0af9f3)
- [inside_virtual_network](resources--fleet--reference--group-002.md#canonical-ff836b8717083785fcc763cf98298a2a5b4213a5691c3bba50f99e9fdc6f502a)
- [interface_list](resources--fleet--reference--group-002.md#canonical-8fded61b07484d6b4cb7278257551fb9fbae015181ec40bd2bb38d864b381672)
- [kubernetes_upgrade_drain](resources--fleet--reference--group-002.md#canonical-b3fca42c2aeb5b7061fe19294dfc4fa35eebac01b348c2d9c5f4431c9390a0b0)
- [log_receiver](resources--fleet--reference--group-002.md#canonical-3df1cf1ed4f9c70511a4f38579ec1cd5749aa8984001404520bc464ea9f2010f)
- [logs_streaming_disabled](resources--fleet--reference--group-002.md#canonical-e69db08b22b6264d05d3610c2282151e0a7e3259540950b8051276833640ac65)
- [network_connectors](resources--fleet--reference--group-002.md#canonical-fb55531a73258d6b540e94ca88f87f2ddeb571685940897e2797a550f3524359)
- [network_firewall](resources--fleet--reference--group-002.md#canonical-917086161f81fec62e50a56057d44ce7779fed3b60c6158a50c95e18eadb9b67)
- [no_bond_devices](resources--fleet--reference--group-002.md#canonical-e3e1473959428b2588fe02d236e1647fd23b202356dcfb55d20c97fa2e6f17e1)
- [no_dc_cluster_group](resources--fleet--reference--group-002.md#canonical-c2b28909dab56b08f193341580a018bedf81f18d2ada71c1098f29e2a4878160)
- [no_storage_device](resources--fleet--reference--group-002.md#canonical-7bfedf647d6488cf41b8b8a93ba5e3859a1b5bf73fe41b2b6c85e198f8287d67)
- [no_storage_interfaces](resources--fleet--reference--group-002.md#canonical-c94745f0bc638d751358bc861580c67e9a3a85171a44abe6f03990a8127968ec)
- [no_storage_static_routes](resources--fleet--reference--group-002.md#canonical-ad162335f354a460814f7dfc264820728c5dda3f75efcbcb98990ba30d3bd05c)
- [outside_virtual_network](resources--fleet--reference--group-002.md#canonical-ebf5a068c5d6d88d88f119f0df953517979070c881eac65fe6c7d2a99066adb1)
- [performance_enhancement_mode](resources--fleet--reference--group-002.md#canonical-3839c5e410e10cfb7034e4a56e6c05dcf39dfd4936bba7e3e253c7f17c10fcff)
- [sriov_interfaces](resources--fleet--reference--group-002.md#canonical-041b0767cc899be9804f15aa4eade24e06041aa19412ec8a5029f8932a306989)
- [storage_class_list](resources--fleet--reference--group-002.md#canonical-459e775e0d90b3cce1de02b34a6b7b22245cb6442ce300f58bbc83579f34bbcb)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-989eca577f306456b1a808c31b1988677bbb594fcacce6a85f042e4951dab251)
- [storage_interface_list](resources--fleet--reference--group-004.md#canonical-a0ad9e77e25f172cd7ebf3e39a25f59ef2e0197cf555254fa1eae93c1269011e)
- [storage_static_routes](resources--fleet--reference--group-004.md#canonical-f44c11471fbbf1faacf11875c91e3b194ee4c8c8c365df7691d83e66869443bc)
- [timeouts](resources--fleet--reference--group-004.md#canonical-c7daec114bab3c6e240b23c9acc29c136c7b4876645a944004681d627803d664)
- [usb_policy](resources--fleet--reference--group-004.md#canonical-777cf2733ada0f5b1100fb3c706ce2c64aefbf9dc25e8aae32a1eaa74ade7df2)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-2f22e2a90601571efd4c85045edeecb1f252e897cfa44bf672d9267183518fdb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d3686d72da47fe47aeadca4fb2a4d8f5f1793cb3a9a3d00ad659634c7dc4aafc"></a>

## allow_all_usb — allow_all_usb / 6ec9b0b7995e / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- allow_all_usb

<a id="canonical-056a0bb93d92a263e5d04f7d3785064f244a05aea8aeb43741141609e8a664d1"></a>

Type: `["object", {}]`. Optional.

\[OneOf: allow\_all\_usb, deny\_all\_usb, usb\_policy\] Configuration parameter for allow all usb.

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

OneOf alternatives in this subsection:

- [allow_all_usb](resources--fleet--reference--group-001.md#canonical-056a0bb93d92a263e5d04f7d3785064f244a05aea8aeb43741141609e8a664d1)
- [deny_all_usb](resources--fleet--reference--group-002.md#canonical-cda8d71228efbcd20187173834fe8393887e2250e76378ab87f6892e4271db3c)
- [usb_policy](resources--fleet--reference--group-004.md#canonical-5a526f17d1768fe3878cc166b4741dfc79b3df781f6f88d652bf81ec9a8992ce)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
allow_all_usb = {}
```

<a id="canonical-83e726f20f02079eb23453260ea95780fed51242a1bc06da9fb17536add82a70"></a>

## Direct properties — allow_all_usb / 6ec9b0b7995e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c0193f87a8b9befe64f97d877439ccc33940ef27e3d705765f182ff7e67357da"></a>

## Next pages — allow_all_usb / 6ec9b0b7995e / 4

- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-6a97758cec25e40f0955dfad6ad390b3ed4989ec29accf52af9593619343596f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7fe5da6934270b3dd6d38629bc94b002c32d4011948f549f5c6255a57cc1d3e7"></a>

## blocked_services — blocked_services / 133cd21a4170 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Property reference](resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- blocked_services

<a id="canonical-b09e1a0d117993efd9166ea5f36dd75f63f536f3bdae105599276761796da6de"></a>

Type: `"object"`. list nested block, Optional.

Disable node local services on this site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("dns",
    "ssh"),
  validators.ConflictingListObjectAttributes("dns",
    "web_user_interface"),
  validators.ConflictingListObjectAttributes("ssh",
    "web_user_interface")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 6,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 6,
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
    "ves.io.schema.rules.repeated.max_items": "6"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "6"
  }
}
```

Terraform syntax:

```terraform
blocked_services {
  # Configure direct properties listed below.
}
```

<a id="canonical-0b7df6b49e0d95b11e5d2cc54bafc77ab7414bb8f491585b60665d9e4b7e17d9"></a>

## Direct properties — blocked_services / 133cd21a4170 / 3

- [dns](resources--fleet--reference--group-002.md#canonical-e9f97568c4c941cee7a36ec8405fa8918ab7fdea73ed7b643a7c56ec86def790): complete subsection reference.

<a id="canonical-a71102c27ff35832e087b8a02411093e414de3f83ddae52cd768d3f52c35c06e"></a>

<a id="canonical-85672d90d9447887b1bcf091eb407534436de2d674d040c487ffc80b98627d3e"></a>

## network_type property — blocked_services / 133cd21a4170 / 4

Type: `"string"`. Optional.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Upstream description:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created automatically and present on all sites
Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE is a private network inside site. It
is a secure network and is not connected to public network. Virtual-network of this type is local to
every site. Two virtual networks of this type on different sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created during provisioning of site User defined per-site
virtual network. Scope of this virtual network is limited to the site. This is not yet supported
Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC directly connects to the public internet.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on RE sites only It is an internally created by the system. They must not be created by
user Virtual Networks with global scope across different sites in F5XC domain. An example global
virtual-network called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [ssh](resources--fleet--reference--group-002.md#canonical-768f1513dc197b11aac930556b08489ca520a72802757283fbd4a13c6ea9ff14): complete subsection reference.

- [web_user_interface](resources--fleet--reference--group-002.md#canonical-bc32b24af35722067ab10e5ef7198410b1c4d59dbf11422dbca441e86d049a18): complete subsection reference.
