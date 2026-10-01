---
page_title: "xcsh_network_firewall reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_network_firewall reference."
---

# xcsh_network_firewall reference

<a id="canonical-5de78ed01fec1f0a650d1483809406b68da66108bb8763e4fd5c2309047dff04"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-266aef3dd6d7cf6cf14ac5b07de378f346372471bd02a5b2334d1a49ba93387f"></a>

## Property reference — Property reference / 155b9f575c66 / 2

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md#canonical-024ddce0e2a1491081bdd9b5fa829cc2057bdaba914b6c85173f71e5817a007f)
- Property reference

<a id="canonical-bd5e81c715c62cf17ad611bcf6b3e05cee479afc1b391108e8a08658e35b864f"></a>

## Direct properties — Property reference / 155b9f575c66 / 3

- [active_enhanced_firewall_policies](resources--network_firewall--reference--group-001.md#canonical-b71f94994787bcbdaa57b85485e17d9c7540dbd19680b5ecaf6adcc2c39e5fa6): complete subsection reference.

- [active_fast_acls](resources--network_firewall--reference--group-001.md#canonical-d691c2b38c9ef70f99b4fab08fd706b8093280d3863f7980f383884a6cd898ec): complete subsection reference.

- [active_forward_proxy_policies](resources--network_firewall--reference--group-001.md#canonical-124425b922be598e564697d2f6fd2c6904a2c5952b853b3513d6169cc480a8d0): complete subsection reference.

- [active_network_policies](resources--network_firewall--reference--group-001.md#canonical-97c8b5f2c8291ddb03d04b476cfb7963da9770a5a6da49a3b643d71ed720220f): complete subsection reference.

<a id="canonical-c8b9de1d583dd5a12d07d0029589356a3cf34c38ba29e96f1701f1322eb58306"></a>

<a id="canonical-8dc18e9dcf27ed003b4736c1c315b6ac161e9c1fa25705688c506d19ec5b1d25"></a>

## annotations property — Property reference / 155b9f575c66 / 4

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

<a id="canonical-98b0b383ead3068ef768077eafc7c7fb3c3d81325e9b1da5025efdb4ed4ffcca"></a>

<a id="canonical-603d4edaa16c0e259a3a6e3df8515ffb7eedb5082158f40fcc918b31e1143576"></a>

## description property — Property reference / 155b9f575c66 / 5

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

<a id="canonical-8d1752907d79dfc6f1515aad49d4e33676334a34466f49aabcffc7d214a04c31"></a>

<a id="canonical-7256fcf15b74a465e072701a55aad24080658a720e6540cf49c5def17c6eb5d9"></a>

## disable property — Property reference / 155b9f575c66 / 6

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

- [disable_fast_acl](resources--network_firewall--reference--group-001.md#canonical-4c03a2d283c600921b55f364f15ea4187c9182f8bfbeb38f2e6c0a572705cdc7): complete subsection reference.

- [disable_forward_proxy_policy](resources--network_firewall--reference--group-001.md#canonical-13785cf1080948939d1bd510c098331c12dc9dbf4a1896a9d0d34f617daa6e16): complete subsection reference.

- [disable_network_policy](resources--network_firewall--reference--group-001.md#canonical-7f1376c76fc493f0b138b259f6d6c0a80db9e12be88fa68e08c28db1d0956849): complete subsection reference.

<a id="canonical-63a988e97818f62fdd5a5a13d8ba95dedc38c792d6e778b4977a87001fd4be01"></a>

<a id="canonical-c94052fb08dea45d5875c0c2c87c4b12e236c90ab3021482dc4c0b97db94eb0c"></a>

## id property — Property reference / 155b9f575c66 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-667913e28722e3703f68e4f3d5f0bf615083eba88c48c0a9228fdafce5ab9250"></a>

<a id="canonical-7559528a2c82eddc67d84437db8a4c329e18a95c9a03d25446fb3e9c205b9f9a"></a>

## labels property — Property reference / 155b9f575c66 / 8

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

<a id="canonical-0df0b0fecb5495f307a3fdf383f9d8a1e8ccef8b8b73ca556acc1fcea36aaaec"></a>

<a id="canonical-85854a4a9f1925e145b343c6f72b5724ab26942f800d566d84e554a2847c1c24"></a>

## name property — Property reference / 155b9f575c66 / 9

Type: `"string"`. Required.

Name of the Network Firewall. Must be unique within the namespace.

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

<a id="canonical-54ab2723c5487812b8449a72bc7eb8372fd210f8852c5d1bd5da505790108697"></a>

<a id="canonical-5ca5828845ca34ffb82bbb8fb8e5f3263e46c82d1ff740115a9f6c5e6e5ebe16"></a>

## namespace property — Property reference / 155b9f575c66 / 10

Type: `"string"`. Optional, Computed.

Namespace for the Network Firewall. The F5 XC API restricts this resource to the system namespace;
it defaults to that value and may be omitted.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
Validators: []validator.String{
  validators.NamespaceValidator(),
  stringvalidator.OneOf("system"),
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

- [timeouts](resources--network_firewall--reference--group-001.md#canonical-13beff0d72fc635fa81f630cac90ae6c65f75e61395007c411ed634fa684cf8a): complete subsection reference.

<a id="canonical-a434d3ed2a7a4a60c5947d5213fac29e4e148d37d346a4fc6c4b31895538f78a"></a>

## All schema paths — Property reference / 155b9f575c66 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `active_enhanced_firewall_policies` | [active_enhanced_firewall_policies](resources--network_firewall--reference--group-001.md#canonical-35e95efa457a65a087180038a67a2b807ffc1de6eba644bc999263f53d8eebc5) |
| `active_enhanced_firewall_policies.enhanced_firewall_policies` | [active_enhanced_firewall_policies.enhanced_firewall_policies](resources--network_firewall--reference--group-001.md#canonical-50e10e625cb038d470fd325c9837a13f132ab019c4c3b5dd9abbbde36fa98000) |
| `active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [active_enhanced_firewall_policies.enhanced_firewall_policies.name](resources--network_firewall--reference--group-001.md#canonical-0d49d996c277cdcc7e73678ed7604b111ae1bcf26aa7112038d3b4e6a61322bc) |
| `active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](resources--network_firewall--reference--group-001.md#canonical-2cfe5c019ada76e3df0dad57e89229237c4e548d4a4358fd56129c7e2bab5b63) |
| `active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](resources--network_firewall--reference--group-001.md#canonical-b1fe98183502c1fa4e2b39eb2fa3241774055d06c7548f55c95a7fdfa4904a20) |
| `active_fast_acls` | [active_fast_acls](resources--network_firewall--reference--group-001.md#canonical-3a573e06bca49c37eabdfbd70f3b6827a9a40f07726091f12b18726c30744758) |
| `active_fast_acls.fast_acls` | [active_fast_acls.fast_acls](resources--network_firewall--reference--group-001.md#canonical-1c41b6c82ca5c4f53e28926a13de07f61e97690e4e0d539745bd51bca17498a2) |
| `active_fast_acls.fast_acls.name` | [active_fast_acls.fast_acls.name](resources--network_firewall--reference--group-001.md#canonical-99f968884226002081ec079b1b90adc4464bd1f773804b1968495bdf3c5d11f9) |
| `active_fast_acls.fast_acls.namespace` | [active_fast_acls.fast_acls.namespace](resources--network_firewall--reference--group-001.md#canonical-5931eabb80ab7a419c550b0a9a930d9b831919a9ef8db03c1c8b0bdbe9890079) |
| `active_fast_acls.fast_acls.tenant` | [active_fast_acls.fast_acls.tenant](resources--network_firewall--reference--group-001.md#canonical-acc11bd4e577de08e7d2eacb2570a37b2747c126a919a90dfab7869d4593e5c9) |
| `active_forward_proxy_policies` | [active_forward_proxy_policies](resources--network_firewall--reference--group-001.md#canonical-33ba77837bb76420ef7f8800e94bc8e117660110d936b44cee08d4538cb84380) |
| `active_forward_proxy_policies.forward_proxy_policies` | [active_forward_proxy_policies.forward_proxy_policies](resources--network_firewall--reference--group-001.md#canonical-5a2a606bdbd0167ae124f4f2f86c2e6c122a14f4ec730632039fd8dbc682ecb0) |
| `active_forward_proxy_policies.forward_proxy_policies.name` | [active_forward_proxy_policies.forward_proxy_policies.name](resources--network_firewall--reference--group-001.md#canonical-922d400179d62749f171ef7645076240f06ebec6afc132f20f9aad5a9212113a) |
| `active_forward_proxy_policies.forward_proxy_policies.namespace` | [active_forward_proxy_policies.forward_proxy_policies.namespace](resources--network_firewall--reference--group-001.md#canonical-374fc0e1a31bbfd2c85d3298a994b70f44d86b9afbe0e49e4e1561aab88da317) |
| `active_forward_proxy_policies.forward_proxy_policies.tenant` | [active_forward_proxy_policies.forward_proxy_policies.tenant](resources--network_firewall--reference--group-001.md#canonical-1b1efdfc9689f0d773e345e7e79d69536730c645d97324c57d0d65581d388149) |
| `active_network_policies` | [active_network_policies](resources--network_firewall--reference--group-001.md#canonical-e6973ab9c78049c9009e88e2d538c723f1d9bcfeefcfc295132b550afb882ada) |
| `active_network_policies.network_policies` | [active_network_policies.network_policies](resources--network_firewall--reference--group-001.md#canonical-80df3aa9095802202b8d4b01509032e81227c02e149ebb640007dc173305d9f5) |
| `active_network_policies.network_policies.name` | [active_network_policies.network_policies.name](resources--network_firewall--reference--group-001.md#canonical-a52a5666ae4d2efb9a8cf71b11c74f9627ccf9589ae643d7e693bed38a49f9ed) |
| `active_network_policies.network_policies.namespace` | [active_network_policies.network_policies.namespace](resources--network_firewall--reference--group-001.md#canonical-59b0d6466501820979d9349d5bbb06a8003ce646a5ba2577a3c7e544ab6ee770) |
| `active_network_policies.network_policies.tenant` | [active_network_policies.network_policies.tenant](resources--network_firewall--reference--group-001.md#canonical-a203adf829b88ce61017866d3abff5611b5406d49eca4438dd5be70582449a07) |
| `annotations` | [annotations](resources--network_firewall--reference--group-001.md#canonical-c8b9de1d583dd5a12d07d0029589356a3cf34c38ba29e96f1701f1322eb58306) |
| `description` | [description](resources--network_firewall--reference--group-001.md#canonical-98b0b383ead3068ef768077eafc7c7fb3c3d81325e9b1da5025efdb4ed4ffcca) |
| `disable` | [disable](resources--network_firewall--reference--group-001.md#canonical-8d1752907d79dfc6f1515aad49d4e33676334a34466f49aabcffc7d214a04c31) |
| `disable_fast_acl` | [disable_fast_acl](resources--network_firewall--reference--group-001.md#canonical-08eab3ea4f770104869ba696b61857e8e236c1f4421be69b52c4c7a20aa43f25) |
| `disable_forward_proxy_policy` | [disable_forward_proxy_policy](resources--network_firewall--reference--group-001.md#canonical-13e304697e8ed650caafd1bb8c520df0f52ddac1279e53018ad21e09b21ae687) |
| `disable_network_policy` | [disable_network_policy](resources--network_firewall--reference--group-001.md#canonical-fcdb84d38948d1b482d2dcdcbb8a7aad97a5fa0693842ac658610d26a640c77b) |
| `id` | [id](resources--network_firewall--reference--group-001.md#canonical-63a988e97818f62fdd5a5a13d8ba95dedc38c792d6e778b4977a87001fd4be01) |
| `labels` | [labels](resources--network_firewall--reference--group-001.md#canonical-667913e28722e3703f68e4f3d5f0bf615083eba88c48c0a9228fdafce5ab9250) |
| `name` | [name](resources--network_firewall--reference--group-001.md#canonical-0df0b0fecb5495f307a3fdf383f9d8a1e8ccef8b8b73ca556acc1fcea36aaaec) |
| `namespace` | [namespace](resources--network_firewall--reference--group-001.md#canonical-54ab2723c5487812b8449a72bc7eb8372fd210f8852c5d1bd5da505790108697) |
| `timeouts` | [timeouts](resources--network_firewall--reference--group-001.md#canonical-a5e5475e3d9c070a365454aa550171fd41f3cf58ac74feddb3bab37443704106) |
| `timeouts.create` | [timeouts.create](resources--network_firewall--reference--group-001.md#canonical-88990cecc59105e1ad975dc22b286c953003637b01f9959456ad12bcbee0979d) |
| `timeouts.delete` | [timeouts.delete](resources--network_firewall--reference--group-001.md#canonical-2934de18080da6b25570f919a01a0d2d6e897c053e51654aa9735104e92cce53) |
| `timeouts.read` | [timeouts.read](resources--network_firewall--reference--group-001.md#canonical-ff3fe818ead8b6bcabed88a6d110b6267333aa865040e0fbf5aa68c5e40dfce6) |
| `timeouts.update` | [timeouts.update](resources--network_firewall--reference--group-001.md#canonical-28bd796279ef7b757d7d02fb9401cebf28a42c416fd0ed46ddadca704c1c0dc5) |

<a id="canonical-a96f3c02197e744c7a273b7ef38f6b480ae79b83f500739ca750b56f8cad71ef"></a>

## Next pages — Property reference / 155b9f575c66 / 12

- [active_enhanced_firewall_policies](resources--network_firewall--reference--group-001.md#canonical-b71f94994787bcbdaa57b85485e17d9c7540dbd19680b5ecaf6adcc2c39e5fa6)
- [active_fast_acls](resources--network_firewall--reference--group-001.md#canonical-d691c2b38c9ef70f99b4fab08fd706b8093280d3863f7980f383884a6cd898ec)
- [active_forward_proxy_policies](resources--network_firewall--reference--group-001.md#canonical-124425b922be598e564697d2f6fd2c6904a2c5952b853b3513d6169cc480a8d0)
- [active_network_policies](resources--network_firewall--reference--group-001.md#canonical-97c8b5f2c8291ddb03d04b476cfb7963da9770a5a6da49a3b643d71ed720220f)
- [disable_fast_acl](resources--network_firewall--reference--group-001.md#canonical-4c03a2d283c600921b55f364f15ea4187c9182f8bfbeb38f2e6c0a572705cdc7)
- [disable_forward_proxy_policy](resources--network_firewall--reference--group-001.md#canonical-13785cf1080948939d1bd510c098331c12dc9dbf4a1896a9d0d34f617daa6e16)
- [disable_network_policy](resources--network_firewall--reference--group-001.md#canonical-7f1376c76fc493f0b138b259f6d6c0a80db9e12be88fa68e08c28db1d0956849)
- [timeouts](resources--network_firewall--reference--group-001.md#canonical-13beff0d72fc635fa81f630cac90ae6c65f75e61395007c411ed634fa684cf8a)
- [xcsh_network_firewall](../resources/network_firewall.md#canonical-024ddce0e2a1491081bdd9b5fa829cc2057bdaba914b6c85173f71e5817a007f)

<a id="canonical-b71f94994787bcbdaa57b85485e17d9c7540dbd19680b5ecaf6adcc2c39e5fa6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-109da5d9dd705033d4645e5b1c479e5e6c5206a8fdea77c00a47d4c4bd3300a6"></a>

## active_enhanced_firewall_policies — active_enhanced_firewall_policies / 219e24c2cab9 / 2

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md#canonical-024ddce0e2a1491081bdd9b5fa829cc2057bdaba914b6c85173f71e5817a007f)
- [Property reference](resources--network_firewall--reference--group-001.md#canonical-5de78ed01fec1f0a650d1483809406b68da66108bb8763e4fd5c2309047dff04)
- active_enhanced_firewall_policies

<a id="canonical-35e95efa457a65a087180038a67a2b807ffc1de6eba644bc999263f53d8eebc5"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: active\_enhanced\_firewall\_policies, active\_network\_policies, disable\_network\_policy;
Default: disable\_network\_policy\] List of Enhanced Firewall Policies These policies use
session-based rules and provide all OPTIONS available under firewall policies with an additional
option for service insertion.

Upstream description:

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("enhanced_firewall_policies")}
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

OneOf alternatives in this subsection:

- [active_enhanced_firewall_policies](resources--network_firewall--reference--group-001.md#canonical-35e95efa457a65a087180038a67a2b807ffc1de6eba644bc999263f53d8eebc5)
- [active_network_policies](resources--network_firewall--reference--group-001.md#canonical-e6973ab9c78049c9009e88e2d538c723f1d9bcfeefcfc295132b550afb882ada)
- [disable_network_policy](resources--network_firewall--reference--group-001.md#canonical-fcdb84d38948d1b482d2dcdcbb8a7aad97a5fa0693842ac658610d26a640c77b)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
active_enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-d296dafdb177e6f2c86e3516fdf87621085b1f5b4cbf1fbce27c47ac9cf8e2dd"></a>

## Direct properties — active_enhanced_firewall_policies / 219e24c2cab9 / 3

- [enhanced_firewall_policies](resources--network_firewall--reference--group-001.md#canonical-4f47270f7a354ec1840c00c2fd33955112947e85d6c63ab744f8c7bf7dc3a189): complete subsection reference.

<a id="canonical-042932f8458baf1bd715caa7e112cb1ced5b67c90bc2a90be82426651d1e93ff"></a>

## Next pages — active_enhanced_firewall_policies / 219e24c2cab9 / 4

- [active_enhanced_firewall_policies.enhanced_firewall_policies](resources--network_firewall--reference--group-001.md#canonical-4f47270f7a354ec1840c00c2fd33955112947e85d6c63ab744f8c7bf7dc3a189)
- [Property reference](resources--network_firewall--reference--group-001.md#canonical-5de78ed01fec1f0a650d1483809406b68da66108bb8763e4fd5c2309047dff04)
- [xcsh_network_firewall](../resources/network_firewall.md#canonical-024ddce0e2a1491081bdd9b5fa829cc2057bdaba914b6c85173f71e5817a007f)

<a id="canonical-4f47270f7a354ec1840c00c2fd33955112947e85d6c63ab744f8c7bf7dc3a189"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ecbe124101c92b5f9dde2f7e73b5baa8a0b58ce796b4afae77bcda6d701d8242"></a>

## active_enhanced_firewall_policies.enhanced_firewall_policies — active_enhanced_firewall_policies.enhanced_firewall_policies / 11934b7683fc / 2

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md#canonical-024ddce0e2a1491081bdd9b5fa829cc2057bdaba914b6c85173f71e5817a007f)
- [Property reference](resources--network_firewall--reference--group-001.md#canonical-5de78ed01fec1f0a650d1483809406b68da66108bb8763e4fd5c2309047dff04)
- [active_enhanced_firewall_policies](resources--network_firewall--reference--group-001.md#canonical-b71f94994787bcbdaa57b85485e17d9c7540dbd19680b5ecaf6adcc2c39e5fa6)
- active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-50e10e625cb038d470fd325c9837a13f132ab019c4c3b5dd9abbbde36fa98000"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Enhanced Firewall Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-a66db89e938c33fb376de9d95eb0272e16d992987b30166bbbc78692162fcc7b"></a>

## Direct properties — active_enhanced_firewall_policies.enhanced_firewall_policies / 11934b7683fc / 3

<a id="canonical-0d49d996c277cdcc7e73678ed7604b111ae1bcf26aa7112038d3b4e6a61322bc"></a>

<a id="canonical-4cc6fe2039ebd4c61d0bebde955d4b4f4015a4f76588ae31f744127d1fd2df92"></a>

## name property — active_enhanced_firewall_policies.enhanced_firewall_policies / 11934b7683fc / 4

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

<a id="canonical-2cfe5c019ada76e3df0dad57e89229237c4e548d4a4358fd56129c7e2bab5b63"></a>

<a id="canonical-1c32da1a5736c3ba3261840e8ba367e7c5ed6071a84634b731c82888e4e620c3"></a>

## namespace property — active_enhanced_firewall_policies.enhanced_firewall_policies / 11934b7683fc / 5

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

<a id="canonical-b1fe98183502c1fa4e2b39eb2fa3241774055d06c7548f55c95a7fdfa4904a20"></a>

<a id="canonical-81b5576d7132b171b2d591df2e6a7d8083dd1bd82278af987189ceaf212a19ca"></a>

## tenant property — active_enhanced_firewall_policies.enhanced_firewall_policies / 11934b7683fc / 6

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

<a id="canonical-ad6fb93fd922e4000fab01a0509b58b7fff3e0cfbac26851e5a068aa2922d406"></a>

## Next pages — active_enhanced_firewall_policies.enhanced_firewall_policies / 11934b7683fc / 7

- [active_enhanced_firewall_policies](resources--network_firewall--reference--group-001.md#canonical-b71f94994787bcbdaa57b85485e17d9c7540dbd19680b5ecaf6adcc2c39e5fa6)
- [xcsh_network_firewall](../resources/network_firewall.md#canonical-024ddce0e2a1491081bdd9b5fa829cc2057bdaba914b6c85173f71e5817a007f)

<a id="canonical-d691c2b38c9ef70f99b4fab08fd706b8093280d3863f7980f383884a6cd898ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-13eed5225a93cf683cfd845057406254387a3df7cf3abe1f649286a7933fdae5"></a>

## active_fast_acls — active_fast_acls / 122932e05b41 / 2

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md#canonical-024ddce0e2a1491081bdd9b5fa829cc2057bdaba914b6c85173f71e5817a007f)
- [Property reference](resources--network_firewall--reference--group-001.md#canonical-5de78ed01fec1f0a650d1483809406b68da66108bb8763e4fd5c2309047dff04)
- active_fast_acls

<a id="canonical-3a573e06bca49c37eabdfbd70f3b6827a9a40f07726091f12b18726c30744758"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: active\_fast\_acls, disable\_fast\_acl; Default: disable\_fast\_acl\] Configuration
parameter for active fast acls.

Upstream description:

List of Fast ACL(s).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("fast_acls")}
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

OneOf alternatives in this subsection:

- [active_fast_acls](resources--network_firewall--reference--group-001.md#canonical-3a573e06bca49c37eabdfbd70f3b6827a9a40f07726091f12b18726c30744758)
- [disable_fast_acl](resources--network_firewall--reference--group-001.md#canonical-08eab3ea4f770104869ba696b61857e8e236c1f4421be69b52c4c7a20aa43f25)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
active_fast_acls {
  # Configure direct properties listed below.
}
```

<a id="canonical-68a0750d9ae222f00dee7e39f2ce807295804018504a4f1f1d9e5ef61284bebe"></a>

## Direct properties — active_fast_acls / 122932e05b41 / 3

- [fast_acls](resources--network_firewall--reference--group-001.md#canonical-a7711867e93c5bccfb77788cd446e790c49068f258b949884df1c026789e1f70): complete subsection reference.

<a id="canonical-d5909773b307f6942c13646ca8af4fe50ec109fe21922746a3ce91982a596918"></a>

## Next pages — active_fast_acls / 122932e05b41 / 4

- [active_fast_acls.fast_acls](resources--network_firewall--reference--group-001.md#canonical-a7711867e93c5bccfb77788cd446e790c49068f258b949884df1c026789e1f70)
- [Property reference](resources--network_firewall--reference--group-001.md#canonical-5de78ed01fec1f0a650d1483809406b68da66108bb8763e4fd5c2309047dff04)
- [xcsh_network_firewall](../resources/network_firewall.md#canonical-024ddce0e2a1491081bdd9b5fa829cc2057bdaba914b6c85173f71e5817a007f)

<a id="canonical-a7711867e93c5bccfb77788cd446e790c49068f258b949884df1c026789e1f70"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-865aabd2f1920afc15a5f5865b06307cdd21caabb275e8bc6889b4544a566463"></a>

## active_fast_acls.fast_acls — active_fast_acls.fast_acls / ad48ebb07fa9 / 2

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md#canonical-024ddce0e2a1491081bdd9b5fa829cc2057bdaba914b6c85173f71e5817a007f)
- [Property reference](resources--network_firewall--reference--group-001.md#canonical-5de78ed01fec1f0a650d1483809406b68da66108bb8763e4fd5c2309047dff04)
- [active_fast_acls](resources--network_firewall--reference--group-001.md#canonical-d691c2b38c9ef70f99b4fab08fd706b8093280d3863f7980f383884a6cd898ec)
- active_fast_acls.fast_acls

<a id="canonical-1c41b6c82ca5c4f53e28926a13de07f61e97690e4e0d539745bd51bca17498a2"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Fast ACL(s) active for this network firewall.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
fast_acls {
  # Configure direct properties listed below.
}
```

<a id="canonical-c7c2ce5c85492efc3af0f039227d406591c59c528dd1fde028369720047d6ba0"></a>

## Direct properties — active_fast_acls.fast_acls / ad48ebb07fa9 / 3

<a id="canonical-99f968884226002081ec079b1b90adc4464bd1f773804b1968495bdf3c5d11f9"></a>

<a id="canonical-b1b38679516e6ad43e29261d64050e5abbe093feb4be36ca30a65dc1b2e333fb"></a>

## name property — active_fast_acls.fast_acls / ad48ebb07fa9 / 4

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

<a id="canonical-5931eabb80ab7a419c550b0a9a930d9b831919a9ef8db03c1c8b0bdbe9890079"></a>

<a id="canonical-76a4805c1fffa1ff2ac9e547ae5e20cd2f8c06d6c864c5c3830b80f89fa72888"></a>

## namespace property — active_fast_acls.fast_acls / ad48ebb07fa9 / 5

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

<a id="canonical-acc11bd4e577de08e7d2eacb2570a37b2747c126a919a90dfab7869d4593e5c9"></a>

<a id="canonical-6ce66992370ac9a5263c96f9af2ddf863f61898779cab219c273535015b37a81"></a>

## tenant property — active_fast_acls.fast_acls / ad48ebb07fa9 / 6

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

<a id="canonical-c4f4ded4234fc152997888d0c29f1babd1c0142a531c6f69071cc19bb92992c5"></a>

## Next pages — active_fast_acls.fast_acls / ad48ebb07fa9 / 7

- [active_fast_acls](resources--network_firewall--reference--group-001.md#canonical-d691c2b38c9ef70f99b4fab08fd706b8093280d3863f7980f383884a6cd898ec)
- [xcsh_network_firewall](../resources/network_firewall.md#canonical-024ddce0e2a1491081bdd9b5fa829cc2057bdaba914b6c85173f71e5817a007f)

<a id="canonical-124425b922be598e564697d2f6fd2c6904a2c5952b853b3513d6169cc480a8d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf6c860106b44baf5b6cf9522a11c679e17445409ccfcae2fb64fc3887f40851"></a>

## active_forward_proxy_policies — active_forward_proxy_policies / 2a1811290213 / 2

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md#canonical-024ddce0e2a1491081bdd9b5fa829cc2057bdaba914b6c85173f71e5817a007f)
- [Property reference](resources--network_firewall--reference--group-001.md#canonical-5de78ed01fec1f0a650d1483809406b68da66108bb8763e4fd5c2309047dff04)
- active_forward_proxy_policies

<a id="canonical-33ba77837bb76420ef7f8800e94bc8e117660110d936b44cee08d4538cb84380"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: active\_forward\_proxy\_policies, disable\_forward\_proxy\_policy; Default:
disable\_forward\_proxy\_policy\] Ordered List of Forward Proxy Policies active.

Upstream description:

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("forward_proxy_policies")}
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

OneOf alternatives in this subsection:

- [active_forward_proxy_policies](resources--network_firewall--reference--group-001.md#canonical-33ba77837bb76420ef7f8800e94bc8e117660110d936b44cee08d4538cb84380)
- [disable_forward_proxy_policy](resources--network_firewall--reference--group-001.md#canonical-13e304697e8ed650caafd1bb8c520df0f52ddac1279e53018ad21e09b21ae687)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
active_forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-dd038d8892788199528a2acdd7b5c091df396bc7c974b4d1663f5b7819099063"></a>

## Direct properties — active_forward_proxy_policies / 2a1811290213 / 3

- [forward_proxy_policies](resources--network_firewall--reference--group-001.md#canonical-8e8650bc3732d3bb4b13899dd08d880a2d4c4b9378650428b9815bb66666e4af): complete subsection reference.

<a id="canonical-e899535401a502c4e916791f1ac0c867733dc9e6cb5c940d90739c4b1b9003ed"></a>

## Next pages — active_forward_proxy_policies / 2a1811290213 / 4

- [active_forward_proxy_policies.forward_proxy_policies](resources--network_firewall--reference--group-001.md#canonical-8e8650bc3732d3bb4b13899dd08d880a2d4c4b9378650428b9815bb66666e4af)
- [Property reference](resources--network_firewall--reference--group-001.md#canonical-5de78ed01fec1f0a650d1483809406b68da66108bb8763e4fd5c2309047dff04)
- [xcsh_network_firewall](../resources/network_firewall.md#canonical-024ddce0e2a1491081bdd9b5fa829cc2057bdaba914b6c85173f71e5817a007f)

<a id="canonical-8e8650bc3732d3bb4b13899dd08d880a2d4c4b9378650428b9815bb66666e4af"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2899dcd5a235eed13a99084a8ab2ab60f598e981478fbeaaa775ac72a0367be2"></a>

## active_forward_proxy_policies.forward_proxy_policies — active_forward_proxy_policies.forward_proxy_policies / 8a59bfe837af / 2

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md#canonical-024ddce0e2a1491081bdd9b5fa829cc2057bdaba914b6c85173f71e5817a007f)
- [Property reference](resources--network_firewall--reference--group-001.md#canonical-5de78ed01fec1f0a650d1483809406b68da66108bb8763e4fd5c2309047dff04)
- [active_forward_proxy_policies](resources--network_firewall--reference--group-001.md#canonical-124425b922be598e564697d2f6fd2c6904a2c5952b853b3513d6169cc480a8d0)
- active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-5a2a606bdbd0167ae124f4f2f86c2e6c122a14f4ec730632039fd8dbc682ecb0"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-922572fd081073335babeaff9cd2d8fd10b63cf854d16afdfa2a0d20f3cf48fb"></a>

## Direct properties — active_forward_proxy_policies.forward_proxy_policies / 8a59bfe837af / 3

<a id="canonical-922d400179d62749f171ef7645076240f06ebec6afc132f20f9aad5a9212113a"></a>

<a id="canonical-025ca6a4198e27695afe64b1e53a663c4ef57709dda61bdb224b46f4c47ee8c1"></a>

## name property — active_forward_proxy_policies.forward_proxy_policies / 8a59bfe837af / 4

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

<a id="canonical-374fc0e1a31bbfd2c85d3298a994b70f44d86b9afbe0e49e4e1561aab88da317"></a>

<a id="canonical-f1c08806d3cb0e04c926474953905df0b4e36594cbc3069e1a9651df70dbb64e"></a>

## namespace property — active_forward_proxy_policies.forward_proxy_policies / 8a59bfe837af / 5

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

<a id="canonical-1b1efdfc9689f0d773e345e7e79d69536730c645d97324c57d0d65581d388149"></a>

<a id="canonical-c0578b14941067a05c9b9841125a7451703324e4b6a82671d3d3ae1fcd1cac06"></a>

## tenant property — active_forward_proxy_policies.forward_proxy_policies / 8a59bfe837af / 6

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

<a id="canonical-9240bca0862af5284b96b03733bbb44a76788bac5de725d3b453352bee468f40"></a>

## Next pages — active_forward_proxy_policies.forward_proxy_policies / 8a59bfe837af / 7

- [active_forward_proxy_policies](resources--network_firewall--reference--group-001.md#canonical-124425b922be598e564697d2f6fd2c6904a2c5952b853b3513d6169cc480a8d0)
- [xcsh_network_firewall](../resources/network_firewall.md#canonical-024ddce0e2a1491081bdd9b5fa829cc2057bdaba914b6c85173f71e5817a007f)

<a id="canonical-97c8b5f2c8291ddb03d04b476cfb7963da9770a5a6da49a3b643d71ed720220f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf1a9828a2402143d70ec19c8f41396df784a08e664f404db0272641533f46f6"></a>

## active_network_policies — active_network_policies / 5b7ddc6fc5e3 / 2

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md#canonical-024ddce0e2a1491081bdd9b5fa829cc2057bdaba914b6c85173f71e5817a007f)
- [Property reference](resources--network_firewall--reference--group-001.md#canonical-5de78ed01fec1f0a650d1483809406b68da66108bb8763e4fd5c2309047dff04)
- active_network_policies

<a id="canonical-e6973ab9c78049c9009e88e2d538c723f1d9bcfeefcfc295132b550afb882ada"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("network_policies")}
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
active_network_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-6dc7b78bd31ce128ae6142c354daa5e3b6f431d7c173e8fb0a675f75992066d6"></a>

## Direct properties — active_network_policies / 5b7ddc6fc5e3 / 3

- [network_policies](resources--network_firewall--reference--group-001.md#canonical-737cf88e406ebc8c92d13ec0d2868ef5256ea3ba6c56ffc689b2e11f4ddcf825): complete subsection reference.

<a id="canonical-4ac78dccd94f65f74f2566f764197266e65c740a084a0da8a12909def014ef03"></a>

## Next pages — active_network_policies / 5b7ddc6fc5e3 / 4

- [active_network_policies.network_policies](resources--network_firewall--reference--group-001.md#canonical-737cf88e406ebc8c92d13ec0d2868ef5256ea3ba6c56ffc689b2e11f4ddcf825)
- [Property reference](resources--network_firewall--reference--group-001.md#canonical-5de78ed01fec1f0a650d1483809406b68da66108bb8763e4fd5c2309047dff04)
- [xcsh_network_firewall](../resources/network_firewall.md#canonical-024ddce0e2a1491081bdd9b5fa829cc2057bdaba914b6c85173f71e5817a007f)

<a id="canonical-737cf88e406ebc8c92d13ec0d2868ef5256ea3ba6c56ffc689b2e11f4ddcf825"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c72e1b7f175ffc4ce31f00aacea8453257aca208f64e55558256a6c2320c0526"></a>

## active_network_policies.network_policies — active_network_policies.network_policies / 0faced5a5984 / 2

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md#canonical-024ddce0e2a1491081bdd9b5fa829cc2057bdaba914b6c85173f71e5817a007f)
- [Property reference](resources--network_firewall--reference--group-001.md#canonical-5de78ed01fec1f0a650d1483809406b68da66108bb8763e4fd5c2309047dff04)
- [active_network_policies](resources--network_firewall--reference--group-001.md#canonical-97c8b5f2c8291ddb03d04b476cfb7963da9770a5a6da49a3b643d71ed720220f)
- active_network_policies.network_policies

<a id="canonical-80df3aa9095802202b8d4b01509032e81227c02e149ebb640007dc173305d9f5"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Firewall Policies active for this network firewall.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
network_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-df8727aa467c57b53bf0157f7379ab215147497c4e71b5a6521d9e4222df80f4"></a>

## Direct properties — active_network_policies.network_policies / 0faced5a5984 / 3

<a id="canonical-a52a5666ae4d2efb9a8cf71b11c74f9627ccf9589ae643d7e693bed38a49f9ed"></a>

<a id="canonical-6b5d4bc3166b4e1577be8d396559ceba7ca44efd716f6c155ae80b581e1179ef"></a>

## name property — active_network_policies.network_policies / 0faced5a5984 / 4

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

<a id="canonical-59b0d6466501820979d9349d5bbb06a8003ce646a5ba2577a3c7e544ab6ee770"></a>

<a id="canonical-b7b9cfa8aba876d96d7cde011fed7711f9fcd345d7eb4693457ec039bbdfbcfe"></a>

## namespace property — active_network_policies.network_policies / 0faced5a5984 / 5

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

<a id="canonical-a203adf829b88ce61017866d3abff5611b5406d49eca4438dd5be70582449a07"></a>

<a id="canonical-7f86ce54d4fdec39e91194e86eecff914856176cbf7ee569f3be5309dbd307ad"></a>

## tenant property — active_network_policies.network_policies / 0faced5a5984 / 6

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

<a id="canonical-210d523c751e1258fbb16f658eca544aaa4793ed2fbca12cde81598508cc8165"></a>

## Next pages — active_network_policies.network_policies / 0faced5a5984 / 7

- [active_network_policies](resources--network_firewall--reference--group-001.md#canonical-97c8b5f2c8291ddb03d04b476cfb7963da9770a5a6da49a3b643d71ed720220f)
- [xcsh_network_firewall](../resources/network_firewall.md#canonical-024ddce0e2a1491081bdd9b5fa829cc2057bdaba914b6c85173f71e5817a007f)

<a id="canonical-4c03a2d283c600921b55f364f15ea4187c9182f8bfbeb38f2e6c0a572705cdc7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b213ac2736c3d9c8016612cd49e06d7d172ddc08e14a9fd384d4cebec53eb193"></a>

## disable_fast_acl — disable_fast_acl / a65fd462dd62 / 2

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md#canonical-024ddce0e2a1491081bdd9b5fa829cc2057bdaba914b6c85173f71e5817a007f)
- [Property reference](resources--network_firewall--reference--group-001.md#canonical-5de78ed01fec1f0a650d1483809406b68da66108bb8763e4fd5c2309047dff04)
- disable_fast_acl

<a id="canonical-08eab3ea4f770104869ba696b61857e8e236c1f4421be69b52c4c7a20aa43f25"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for disable fast acl. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
disable_fast_acl = {}
```

<a id="canonical-a07fdbccd06ab0124b7fa8631d63f71d20ae0cb9733c8885e5334578d5f31e19"></a>

## Direct properties — disable_fast_acl / a65fd462dd62 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f401ce5859d137c6e2642d17f76148f181c4bf7baf7b21bbb17558ae4d1303c4"></a>

## Next pages — disable_fast_acl / a65fd462dd62 / 4

- [Property reference](resources--network_firewall--reference--group-001.md#canonical-5de78ed01fec1f0a650d1483809406b68da66108bb8763e4fd5c2309047dff04)
- [xcsh_network_firewall](../resources/network_firewall.md#canonical-024ddce0e2a1491081bdd9b5fa829cc2057bdaba914b6c85173f71e5817a007f)

<a id="canonical-13785cf1080948939d1bd510c098331c12dc9dbf4a1896a9d0d34f617daa6e16"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4e73c5f68bdbda533641b3ca6719853e14dba141ae4655428d143762f532b214"></a>

## disable_forward_proxy_policy — disable_forward_proxy_policy / b83cd56923f6 / 2

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md#canonical-024ddce0e2a1491081bdd9b5fa829cc2057bdaba914b6c85173f71e5817a007f)
- [Property reference](resources--network_firewall--reference--group-001.md#canonical-5de78ed01fec1f0a650d1483809406b68da66108bb8763e4fd5c2309047dff04)
- disable_forward_proxy_policy

<a id="canonical-13e304697e8ed650caafd1bb8c520df0f52ddac1279e53018ad21e09b21ae687"></a>

Type: `["object", {}]`. Optional, Computed.

Policy configuration for this feature. Defaults to \`map\[\]\`. Server applies default when omitted.

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
disable_forward_proxy_policy = {}
```

<a id="canonical-3e7f8155b26f9cd2ce01c4df2ec8e2b969ee107c07d2470b0f2a466483231c73"></a>

## Direct properties — disable_forward_proxy_policy / b83cd56923f6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ffc5dfc8ce810021d6f1d5b2308ee6cdbe9805d01e480c1d416f9a77caaf2c4f"></a>

## Next pages — disable_forward_proxy_policy / b83cd56923f6 / 4

- [Property reference](resources--network_firewall--reference--group-001.md#canonical-5de78ed01fec1f0a650d1483809406b68da66108bb8763e4fd5c2309047dff04)
- [xcsh_network_firewall](../resources/network_firewall.md#canonical-024ddce0e2a1491081bdd9b5fa829cc2057bdaba914b6c85173f71e5817a007f)

<a id="canonical-7f1376c76fc493f0b138b259f6d6c0a80db9e12be88fa68e08c28db1d0956849"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6bf7ac7a831356f0d56cf2baac8342aded334a322b126de619bcf68a3f4eb9f8"></a>

## disable_network_policy — disable_network_policy / ba1236d31c1f / 2

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md#canonical-024ddce0e2a1491081bdd9b5fa829cc2057bdaba914b6c85173f71e5817a007f)
- [Property reference](resources--network_firewall--reference--group-001.md#canonical-5de78ed01fec1f0a650d1483809406b68da66108bb8763e4fd5c2309047dff04)
- disable_network_policy

<a id="canonical-fcdb84d38948d1b482d2dcdcbb8a7aad97a5fa0693842ac658610d26a640c77b"></a>

Type: `["object", {}]`. Optional, Computed.

Policy configuration for this feature. Defaults to \`map\[\]\`. Server applies default when omitted.

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
disable_network_policy = {}
```

<a id="canonical-dd7a46528ed516c464f88793b882f251814ea6d575d384e99e41234c1a615c25"></a>

## Direct properties — disable_network_policy / ba1236d31c1f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-07bc0ea2dc24ed982b866a0a936911808ea17f165c89a5d6deeb1c668aa40274"></a>

## Next pages — disable_network_policy / ba1236d31c1f / 4

- [Property reference](resources--network_firewall--reference--group-001.md#canonical-5de78ed01fec1f0a650d1483809406b68da66108bb8763e4fd5c2309047dff04)
- [xcsh_network_firewall](../resources/network_firewall.md#canonical-024ddce0e2a1491081bdd9b5fa829cc2057bdaba914b6c85173f71e5817a007f)

<a id="canonical-13beff0d72fc635fa81f630cac90ae6c65f75e61395007c411ed634fa684cf8a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-db34e3795e8be49a7e94fe2f7a2a15489a4d607c6db34567c420a0f065860792"></a>

## timeouts — timeouts / 38672fbd2562 / 2

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md#canonical-024ddce0e2a1491081bdd9b5fa829cc2057bdaba914b6c85173f71e5817a007f)
- [Property reference](resources--network_firewall--reference--group-001.md#canonical-5de78ed01fec1f0a650d1483809406b68da66108bb8763e4fd5c2309047dff04)
- timeouts

<a id="canonical-a5e5475e3d9c070a365454aa550171fd41f3cf58ac74feddb3bab37443704106"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-4fa083a3b76fe795a6b0275510c9d92106e9c9d4e8fc02e26fd3763c57484d5d"></a>

## Direct properties — timeouts / 38672fbd2562 / 3

<a id="canonical-88990cecc59105e1ad975dc22b286c953003637b01f9959456ad12bcbee0979d"></a>

<a id="canonical-1b912b5f52fc6cd0a1c0125b15d2c1c93fccf9b133d1006cf894c09d13674fd7"></a>

## create property — timeouts / 38672fbd2562 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2934de18080da6b25570f919a01a0d2d6e897c053e51654aa9735104e92cce53"></a>

<a id="canonical-df5045b86db945349bbe283218eac0826768a6355c1ed16575d13fba269fa25e"></a>

## delete property — timeouts / 38672fbd2562 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-ff3fe818ead8b6bcabed88a6d110b6267333aa865040e0fbf5aa68c5e40dfce6"></a>

<a id="canonical-6ff2698291e13842b406f47ef19325aded81c98bcb631dea4dd8ac88299bd0d9"></a>

## read property — timeouts / 38672fbd2562 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-28bd796279ef7b757d7d02fb9401cebf28a42c416fd0ed46ddadca704c1c0dc5"></a>

<a id="canonical-a80e71ddde1465a8c48e272c6044f09b3ce18cba22f286b343e512d8c5a4d17b"></a>

## update property — timeouts / 38672fbd2562 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-7e75c544c71432c491be96cf739b174195cc49e5fa62c70c9ff840f4d80d983c"></a>

## Next pages — timeouts / 38672fbd2562 / 8

- [Property reference](resources--network_firewall--reference--group-001.md#canonical-5de78ed01fec1f0a650d1483809406b68da66108bb8763e4fd5c2309047dff04)
- [xcsh_network_firewall](../resources/network_firewall.md#canonical-024ddce0e2a1491081bdd9b5fa829cc2057bdaba914b6c85173f71e5817a007f)
