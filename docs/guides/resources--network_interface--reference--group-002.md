---
page_title: "xcsh_network_interface reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_interface reference."
---

# xcsh_network_interface reference

<a id="canonical-0f01dfedbc3aadd2d4dda2a29f687badf24702291a172cbf669b1778c36218c4"></a>

## Direct properties — ethernet_interface.static_ip / b8d8b7423779 / 3

- [cluster_static_ip](resources--network_interface--reference--group-002.md#canonical-2b0f8b6bad6aa6e8f4384686ba9b5f330cf5bfc67d0195ddc9bf32a7cc47c2a2): complete subsection reference.

- [node_static_ip](resources--network_interface--reference--group-002.md#canonical-b3449958c4a858769705fe1531100dbc7862f9bd32865b8b73d104b8cb05731d): complete subsection reference.

<a id="canonical-e91d320724e27a229e451206b4454ecc3fe02ba365273b12d1a7d22257625bd7"></a>

## Next pages — ethernet_interface.static_ip / b8d8b7423779 / 4

- [ethernet_interface.static_ip.cluster_static_ip](resources--network_interface--reference--group-002.md#canonical-2b0f8b6bad6aa6e8f4384686ba9b5f330cf5bfc67d0195ddc9bf32a7cc47c2a2)
- [ethernet_interface.static_ip.node_static_ip](resources--network_interface--reference--group-002.md#canonical-b3449958c4a858769705fe1531100dbc7862f9bd32865b8b73d104b8cb05731d)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-2b0f8b6bad6aa6e8f4384686ba9b5f330cf5bfc67d0195ddc9bf32a7cc47c2a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2bda4bf4add39873160e600064989d9c8f0b34135bbf3b54133b297a8cb413a"></a>

## ethernet_interface.static_ip.cluster_static_ip — ethernet_interface.static_ip.cluster_static_ip / e0c56d4f87ed / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [ethernet_interface.static_ip](resources--network_interface--reference--group-001.md#canonical-cab9271e1730e6f9466f7224ac0cdb714ed64f8faa8cc1318c98043323baf448)
- ethernet_interface.static_ip.cluster_static_ip

<a id="canonical-6a939f7a7a5d9fb7934c25d39a50aa46819ab8593b205ac02d66b183b4b61c95"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for cluster.

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
cluster_static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-73714c2dbb7cde7b6751997ef88b7b23a9e51c46b0f0e1fcca83b9fc958bd569"></a>

## Direct properties — ethernet_interface.static_ip.cluster_static_ip / e0c56d4f87ed / 3

<a id="canonical-5737da7b60ec4c327c8026de727f9b79d8d47a924d95eaf2941b4d26179bbdc3"></a>

<a id="canonical-77cfcb9ea3713e717864f4698cdc071c46ca94682a9ce38bb2513f19c742119f"></a>

## interface_ip_map property — ethernet_interface.static_ip.cluster_static_ip / e0c56d4f87ed / 4

Type: `["map", "string"]`. Optional.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

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
    "ves.io.schema.rules.map.max_pairs": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  }
}
```

<a id="canonical-d383a661aa1ffa04367582372925074c1cc3f1714f1d743479eeb896bee82a56"></a>

## Next pages — ethernet_interface.static_ip.cluster_static_ip / e0c56d4f87ed / 5

- [ethernet_interface.static_ip](resources--network_interface--reference--group-001.md#canonical-cab9271e1730e6f9466f7224ac0cdb714ed64f8faa8cc1318c98043323baf448)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-b3449958c4a858769705fe1531100dbc7862f9bd32865b8b73d104b8cb05731d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-effef2c709dd0cde1861697d41a76b1b777e186ffb4273d4e03fe559511feb73"></a>

## ethernet_interface.static_ip.node_static_ip — ethernet_interface.static_ip.node_static_ip / cdb6d6c5cbdd / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [ethernet_interface.static_ip](resources--network_interface--reference--group-001.md#canonical-cab9271e1730e6f9466f7224ac0cdb714ed64f8faa8cc1318c98043323baf448)
- ethernet_interface.static_ip.node_static_ip

<a id="canonical-d70deaa7005b2adac45dd4b06964f96ce201636b4b09ecc1fd8b3a813393f78d"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for a node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_address")}
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
node_static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-64636b596f5100658b5b50e07daa1d8fafe7f5eaa18e172ea6457ffcd8feba39"></a>

## Direct properties — ethernet_interface.static_ip.node_static_ip / cdb6d6c5cbdd / 3

<a id="canonical-2dbaf847402ec4cfc27eef77c4bcc1416f495d74d428cea25f2dc3ee25f0b3ea"></a>

<a id="canonical-2c083d14f3214a11ef226d7f9094c3662027c4504845cca1a4d3ffd702d2879d"></a>

## default_gw property — ethernet_interface.static_ip.node_static_ip / cdb6d6c5cbdd / 4

Type: `"string"`. Optional.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

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

<a id="canonical-4f883f9bf03d433e5d41c98daf15eeaa39eb4e15b64725621a048a49728f4b5a"></a>

<a id="canonical-bd583191736832f62439d9416d4e8e4a0b7ddf900f95f7d9e36355761c94b1af"></a>

## dns_server property — ethernet_interface.static_ip.node_static_ip / cdb6d6c5cbdd / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-7bdce07a4740d26871beccccd5da616e98e84ad4e0e4b1370110e3cd7a63a568"></a>

<a id="canonical-aae2fb06579e700d48d90dd986a9b1c09398f3327bef5d10b80a351902ecce85"></a>

## ip_address property — ethernet_interface.static_ip.node_static_ip / cdb6d6c5cbdd / 6

Type: `"string"`. Optional.

IP address of the interface and prefix length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  validators.CIDRValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

<a id="canonical-28638e7cb8c37cc6935f1fc3cf1b3f4bef9186d93bcb6991d2b2695141c5036e"></a>

## Next pages — ethernet_interface.static_ip.node_static_ip / cdb6d6c5cbdd / 7

- [ethernet_interface.static_ip](resources--network_interface--reference--group-001.md#canonical-cab9271e1730e6f9466f7224ac0cdb714ed64f8faa8cc1318c98043323baf448)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-82065ed290786f0db2ef4d7e91ef35f55e1eeb3eeb5b3835fbcfd751f3f864a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ffacb40df9d46b1b843c222374b65b1a1ef02e50a9790ec1819bdeb14bb7e6b"></a>

## ethernet_interface.static_ipv6_address — ethernet_interface.static_ipv6_address / f2f8d38771cb / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- ethernet_interface.static_ipv6_address

<a id="canonical-301fb268d2e32610b857aa8c1512cd5c6b32b1a2a69fdc553ddede6695b27c37"></a>

Type: `"object"`. single nested block, Optional.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("cluster_static_ip",
    "node_static_ip")}
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
  "x-ves-oneof-field-network_prefix_choice": "[\"cluster_static_ip\",\"node_static_ip\"]"
}
```

Terraform syntax:

```terraform
static_ipv6_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-9f71f4c4df3b9e0deaa80ca45fdbd8086a29da3e75b5eed6a17a109dbe8e48ff"></a>

## Direct properties — ethernet_interface.static_ipv6_address / f2f8d38771cb / 3

- [cluster_static_ip](resources--network_interface--reference--group-002.md#canonical-b0e8d04d45e4c20e185c9fdfc3959bbe73a2dc4ef70745ed831fbc3241e6f731): complete subsection reference.

- [node_static_ip](resources--network_interface--reference--group-002.md#canonical-2853885475b7ac9b80c8c72356624365cc15160b82550a9cf35425fa603a79a6): complete subsection reference.

<a id="canonical-e1f0be49d80fee8a0395297462ab8bf88e62fe3303dc1b64d353d3dd2dd0f2d0"></a>

## Next pages — ethernet_interface.static_ipv6_address / f2f8d38771cb / 4

- [ethernet_interface.static_ipv6_address.cluster_static_ip](resources--network_interface--reference--group-002.md#canonical-b0e8d04d45e4c20e185c9fdfc3959bbe73a2dc4ef70745ed831fbc3241e6f731)
- [ethernet_interface.static_ipv6_address.node_static_ip](resources--network_interface--reference--group-002.md#canonical-2853885475b7ac9b80c8c72356624365cc15160b82550a9cf35425fa603a79a6)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-b0e8d04d45e4c20e185c9fdfc3959bbe73a2dc4ef70745ed831fbc3241e6f731"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0efeac662a85be9bfe95a19e1dc0f0500c181075a97414491caef3ec2fbaceaa"></a>

## ethernet_interface.static_ipv6_address.cluster_static_ip — ethernet_interface.static_ipv6_address.cluster_static_ip / 6a1b293274c9 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [ethernet_interface.static_ipv6_address](resources--network_interface--reference--group-002.md#canonical-82065ed290786f0db2ef4d7e91ef35f55e1eeb3eeb5b3835fbcfd751f3f864a6)
- ethernet_interface.static_ipv6_address.cluster_static_ip

<a id="canonical-d153016fcdd189f2161cec708be092a601d3d42d2364deec057154f6c45cd7fe"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for cluster.

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
cluster_static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-6212e1733137e1ae06083408f02c4c5dc41566f72d3ec6433e26f7f738c04461"></a>

## Direct properties — ethernet_interface.static_ipv6_address.cluster_static_ip / 6a1b293274c9 / 3

<a id="canonical-7061c448dfa6b665754045c4d2b9d339578b0bcbb9066131b9544e2ccd82b066"></a>

<a id="canonical-dd8afec3e2348dfe4f99efcc85abe2c52bf66348a8ab51c74a5d17277e778a34"></a>

## interface_ip_map property — ethernet_interface.static_ipv6_address.cluster_static_ip / 6a1b293274c9 / 4

Type: `["map", "string"]`. Optional.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

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
    "ves.io.schema.rules.map.max_pairs": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  }
}
```

<a id="canonical-05df768a195ca48d2f5f86f3371a643274757c70ce4d48ad38594d37cf70d74a"></a>

## Next pages — ethernet_interface.static_ipv6_address.cluster_static_ip / 6a1b293274c9 / 5

- [ethernet_interface.static_ipv6_address](resources--network_interface--reference--group-002.md#canonical-82065ed290786f0db2ef4d7e91ef35f55e1eeb3eeb5b3835fbcfd751f3f864a6)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-2853885475b7ac9b80c8c72356624365cc15160b82550a9cf35425fa603a79a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9372878b6c2fa2a3bfa6111cabcdb15119d48bcab7d854cb3dc0c1c3e28ad8e9"></a>

## ethernet_interface.static_ipv6_address.node_static_ip — ethernet_interface.static_ipv6_address.node_static_ip / 36549812505f / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [ethernet_interface.static_ipv6_address](resources--network_interface--reference--group-002.md#canonical-82065ed290786f0db2ef4d7e91ef35f55e1eeb3eeb5b3835fbcfd751f3f864a6)
- ethernet_interface.static_ipv6_address.node_static_ip

<a id="canonical-d20402919e10de0d950f5ea8f899f5d67291e20eb8688d64947f65d8bbe5c7b2"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for a node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_address")}
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
node_static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-bcd871be7d8d27e3c683281bca9c9244bade148523f3a07615177e2f1923a62b"></a>

## Direct properties — ethernet_interface.static_ipv6_address.node_static_ip / 36549812505f / 3

<a id="canonical-152f91ddc7eae9822fd45692fea0bf829269c2ca52cab80cd2fc5ef2d2c1be1a"></a>

<a id="canonical-40fbeb6c462880808cb2882684bb4425d7998202612901b2fe83933ee676e8ea"></a>

## default_gw property — ethernet_interface.static_ipv6_address.node_static_ip / 36549812505f / 4

Type: `"string"`. Optional.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

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

<a id="canonical-30d74b0aa093f79afeccdf7803b6088108362f5c334fcfdfe057a0e44ecae714"></a>

<a id="canonical-f1dd5fea2069fc81fb74bea310b151c6d9dc8cab0af90ab6eafd1d6e97890a4a"></a>

## dns_server property — ethernet_interface.static_ipv6_address.node_static_ip / 36549812505f / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-7971030a9ac520e6eedecc19034aa7caeac217fb34a13138901deda62dc7580f"></a>

<a id="canonical-c6113993b67da8e6bacefcca9c3082b75f5c61bd19f8090e54af3b055218c6ca"></a>

## ip_address property — ethernet_interface.static_ipv6_address.node_static_ip / 36549812505f / 6

Type: `"string"`. Optional.

IP address of the interface and prefix length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  validators.CIDRValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

<a id="canonical-861d6c31859b3a290192af62496ac63d65a182bd7208e59f5772de08000ad897"></a>

## Next pages — ethernet_interface.static_ipv6_address.node_static_ip / 36549812505f / 7

- [ethernet_interface.static_ipv6_address](resources--network_interface--reference--group-002.md#canonical-82065ed290786f0db2ef4d7e91ef35f55e1eeb3eeb5b3835fbcfd751f3f864a6)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-65a5224a94a01239bef6787e02312d7c041201fc3cded5a2a79a9d0992228f05"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa7a8aa57ca7d322b33642badbc9fcb9125f204cff93809e265891729af65461"></a>

## ethernet_interface.storage_network — ethernet_interface.storage_network / 30f6481f71f4 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- ethernet_interface.storage_network

<a id="canonical-7ff07e6be90117e3e8bb1aad55041c91828b69ab7b035addabeada48182ad13a"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for storage network.

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
storage_network = {}
```

<a id="canonical-032f0c28152c6f5f56600c95fc46a504e9916c138616b84fb1efec732ef80c03"></a>

## Direct properties — ethernet_interface.storage_network / 30f6481f71f4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8fd2f934136c01569dcf4385fa89f8c3f146292a81ffce1a92305b9ee4428969"></a>

## Next pages — ethernet_interface.storage_network / 30f6481f71f4 / 4

- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-87f18d894e38d753a597d57b34415cec0096cc21131754751028e636f41b66ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-088ef060c6bcaec582de92ff71cfd6b6c1052128b0d289eb37e3ec7a259dbd55"></a>

## ethernet_interface.untagged — ethernet_interface.untagged / 8a56fe3f4e68 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- ethernet_interface.untagged

<a id="canonical-a87a5701bb14fcfda551d90bedf6696e6fa4b1ab0aa87b530b53689cba07b030"></a>

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
untagged = {}
```

<a id="canonical-46f8bacb9d5bb632c6262482570d319f01f4d04d83bf46ca6324d5371b6b171d"></a>

## Direct properties — ethernet_interface.untagged / 8a56fe3f4e68 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-812a96ed90228854c065951151462bcdcfaf7d9fa1b697c501a6f012def8f583"></a>

## Next pages — ethernet_interface.untagged / 8a56fe3f4e68 / 4

- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-1c594b57dc38583cede4438254bed4c980392fdf365967edc1ad24fac341cbd0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-209eeeff870f5bec27f93697a1b7859bb53d2ce4da6cf5435a51598f780dbf7d"></a>

## layer2_interface — layer2_interface / dd525c0627f0 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- layer2_interface

<a id="canonical-e83fe3d490b7b1b9797f3df5de4ff22bc8deb68ea422ed8378cb6ac6f9fb18b9"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for layer2 interface.

Upstream description:

Layer2 Interface Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("l2sriov_interface",
    "l2vlan_interface"),
  validators.ConflictingObjectAttributes("l2sriov_interface",
    "l2vlan_slo_interface"),
  validators.ConflictingObjectAttributes("l2vlan_interface",
    "l2vlan_slo_interface")}
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
  "x-ves-oneof-field-layer2_interface_choice": "[\"l2sriov_interface\",\"l2vlan_interface\",\"l2vlan_slo_interface\"]"
}
```

Terraform syntax:

```terraform
layer2_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-967b05069c320e6a187c378b95d718ce48cf3c756e7adbfd4a069e3d26b3cd9b"></a>

## Direct properties — layer2_interface / dd525c0627f0 / 3

- [l2sriov_interface](resources--network_interface--reference--group-002.md#canonical-819fc57a5223512bd7352f15eff6216809dc8c6947029b9e75b0b3af2e405d96): complete subsection reference.

- [l2vlan_interface](resources--network_interface--reference--group-002.md#canonical-c8437ab1fa91519402bf312b915a51d1d4e19a88406ca315a15e3ecfe38af8d5): complete subsection reference.

- [l2vlan_slo_interface](resources--network_interface--reference--group-002.md#canonical-72ed369e635d58839f9540b86ea58b25c05d7fd831b0570dfb083ac6bfb69496): complete subsection reference.

<a id="canonical-2ca0e6b2165a0cfc9cac8fe89dbe66e4427ba586bb0e893750044110c7e0dfca"></a>

## Next pages — layer2_interface / dd525c0627f0 / 4

- [layer2_interface.l2sriov_interface](resources--network_interface--reference--group-002.md#canonical-819fc57a5223512bd7352f15eff6216809dc8c6947029b9e75b0b3af2e405d96)
- [layer2_interface.l2vlan_interface](resources--network_interface--reference--group-002.md#canonical-c8437ab1fa91519402bf312b915a51d1d4e19a88406ca315a15e3ecfe38af8d5)
- [layer2_interface.l2vlan_slo_interface](resources--network_interface--reference--group-002.md#canonical-72ed369e635d58839f9540b86ea58b25c05d7fd831b0570dfb083ac6bfb69496)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-819fc57a5223512bd7352f15eff6216809dc8c6947029b9e75b0b3af2e405d96"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-32b615bb36dc264144fd4364ce31579e742e8fda2016310c3afdbce25d213e3b"></a>

## layer2_interface.l2sriov_interface — layer2_interface.l2sriov_interface / 19db76c79f6a / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [layer2_interface](resources--network_interface--reference--group-002.md#canonical-1c594b57dc38583cede4438254bed4c980392fdf365967edc1ad24fac341cbd0)
- layer2_interface.l2sriov_interface

<a id="canonical-2ae9ea6c87a7a85bbc990c5b62b10690956df01954b7faad8f67f41745e24aa3"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for l2sriov interface.

Upstream description:

Layer2 SR-IOV Interface Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("device"),
  validators.ConflictingObjectAttributes("untagged",
    "vlan_id")}
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
  "x-ves-oneof-field-vlan_choice": "[\"untagged\",\"vlan_id\"]"
}
```

Terraform syntax:

```terraform
l2sriov_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-fed22b0580cf0c291bb7e93d27fb6f7f9f155273e8b5922d55512bf3c65a60f7"></a>

## Direct properties — layer2_interface.l2sriov_interface / 19db76c79f6a / 3

<a id="canonical-404c1e6ba76e08387c9d77df9342e24d92f6172d9a00cc70bb43cff11a480ef0"></a>

<a id="canonical-976de8f654822171f5e9378f929344bea4ac89e93cfda01b7c581013c7e7432e"></a>

## device property — layer2_interface.l2sriov_interface / 19db76c79f6a / 4

Type: `"string"`. Optional.

Ethernet Device. Physical ethernet interface.

Upstream description:

Physical ethernet interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

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

- [untagged](resources--network_interface--reference--group-002.md#canonical-a5790a9b20d2b2b5342ba64471e0e743f8954b7d5493147365084261fa0a0d82): complete subsection reference.

<a id="canonical-d00c578d39e778cfd60387e5bdf3f496eb83221cc045beb9f44b51e12630bbac"></a>

<a id="canonical-4c40585653a1912526afb573baf68ffd7b57158971d7414b75376f42534eab39"></a>

## vlan_id property — layer2_interface.l2sriov_interface / 19db76c79f6a / 5

Type: `"number"`. Optional.

Exclusive with \[untagged\] Configure a VLAN tagged interface.

Upstream description:

Exclusive with \[untagged\] Configure a VLAN tagged interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 4095),
}
```

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

<a id="canonical-94cd1c36e44c57b7f7941f43ad995914373d8da01e46f93e30629d36e5125612"></a>

## Next pages — layer2_interface.l2sriov_interface / 19db76c79f6a / 6

- [layer2_interface.l2sriov_interface.untagged](resources--network_interface--reference--group-002.md#canonical-a5790a9b20d2b2b5342ba64471e0e743f8954b7d5493147365084261fa0a0d82)
- [layer2_interface](resources--network_interface--reference--group-002.md#canonical-1c594b57dc38583cede4438254bed4c980392fdf365967edc1ad24fac341cbd0)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-a5790a9b20d2b2b5342ba64471e0e743f8954b7d5493147365084261fa0a0d82"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-211d0e16544715d714f73bf3e434b3957284f0b09e9e39e9c66a05d43683a6be"></a>

## layer2_interface.l2sriov_interface.untagged — layer2_interface.l2sriov_interface.untagged / add6f93773ac / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [layer2_interface](resources--network_interface--reference--group-002.md#canonical-1c594b57dc38583cede4438254bed4c980392fdf365967edc1ad24fac341cbd0)
- [layer2_interface.l2sriov_interface](resources--network_interface--reference--group-002.md#canonical-819fc57a5223512bd7352f15eff6216809dc8c6947029b9e75b0b3af2e405d96)
- layer2_interface.l2sriov_interface.untagged

<a id="canonical-2b6240834c2b7be8af41328df0235100df6f3a933aa9f74b97b0e4d48c42f46d"></a>

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
untagged = {}
```

<a id="canonical-a97f03a614a33c419965f1b38a0e9b50a83f203572a3a45afa3fd1de3f7d3c2f"></a>

## Direct properties — layer2_interface.l2sriov_interface.untagged / add6f93773ac / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5514445f93acce854c2c461bb0bc20e5caa5af0a150555636d246434d9d1a71f"></a>

## Next pages — layer2_interface.l2sriov_interface.untagged / add6f93773ac / 4

- [layer2_interface.l2sriov_interface](resources--network_interface--reference--group-002.md#canonical-819fc57a5223512bd7352f15eff6216809dc8c6947029b9e75b0b3af2e405d96)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-c8437ab1fa91519402bf312b915a51d1d4e19a88406ca315a15e3ecfe38af8d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5ec0eeb7fe44e5b36a62c7f2bf3e089189fd91bfab0234165dea33cbf35d34e"></a>

## layer2_interface.l2vlan_interface — layer2_interface.l2vlan_interface / b951310f8af5 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [layer2_interface](resources--network_interface--reference--group-002.md#canonical-1c594b57dc38583cede4438254bed4c980392fdf365967edc1ad24fac341cbd0)
- layer2_interface.l2vlan_interface

<a id="canonical-6b68b78583b4ab1bbf657e78184145acc9af43f231669ec56a2789d6d59dc06a"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for l2vlan interface.

Upstream description:

Layer2 VLAN Interface Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("device",
    "vlan_id")}
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
l2vlan_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-1f328a71c0968dac24ce9294a9f49f8a49cbbc49408ad418edb562ce2ef37ee9"></a>

## Direct properties — layer2_interface.l2vlan_interface / b951310f8af5 / 3

<a id="canonical-af303c0f4aeecc60838cc6dc0a54f8558208a0fb7b29e336b4ee6cac8fec1de0"></a>

<a id="canonical-fc4f5e3e5fc6fc599350d735465fa7bdd4da7d69d63db8293a56b7126bdf68fb"></a>

## device property — layer2_interface.l2vlan_interface / b951310f8af5 / 4

Type: `"string"`. Optional.

Ethernet Device. Physical ethernet interface.

Upstream description:

Physical ethernet interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

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

<a id="canonical-0bf9ab6fdb2ca90b571c1a8b08fafbbbba6a6578aa527dadc77c04e01f18cc85"></a>

<a id="canonical-1201e72899a9ccbb11421c9039d3f2e4e9cfe872c360830fe39f0b112983e2ea"></a>

## vlan_id property — layer2_interface.l2vlan_interface / b951310f8af5 / 5

Type: `"number"`. Optional.

VLAN ID. VLAN ID

Upstream description:

VLAN ID

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 4095),
}
```

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```

<a id="canonical-cc76ad50a4dd61c53939792539be4e8498c2c9ece5c5b6eb935f0a0cbfb02abc"></a>

## Next pages — layer2_interface.l2vlan_interface / b951310f8af5 / 6

- [layer2_interface](resources--network_interface--reference--group-002.md#canonical-1c594b57dc38583cede4438254bed4c980392fdf365967edc1ad24fac341cbd0)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-72ed369e635d58839f9540b86ea58b25c05d7fd831b0570dfb083ac6bfb69496"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0cceecd968295e0ee58d32f65c4417b38f7754a29471785ea9149b7f7ca26706"></a>

## layer2_interface.l2vlan_slo_interface — layer2_interface.l2vlan_slo_interface / 844e5dd86bce / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [layer2_interface](resources--network_interface--reference--group-002.md#canonical-1c594b57dc38583cede4438254bed4c980392fdf365967edc1ad24fac341cbd0)
- layer2_interface.l2vlan_slo_interface

<a id="canonical-1e603ca6b506635b704b91fb12737a2a1b5b373ed71bd78f86f38dcad234de0e"></a>

Type: `"object"`. single nested block, Optional.

Layer2 Site Local Outside VLAN Interface Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("vlan_id")}
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
l2vlan_slo_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-42ff2e209cfae4f0c0b60293313fe35ea5237fb8088203e20dc51728e6157d0b"></a>

## Direct properties — layer2_interface.l2vlan_slo_interface / 844e5dd86bce / 3

<a id="canonical-67387467502d314677eb485651447c1910e7336ad39b4fd5fe202be5cac28714"></a>

<a id="canonical-92bbecab37e305b12c87b4ed876d6af92677046e8c6c2d2268955628d014f758"></a>

## vlan_id property — layer2_interface.l2vlan_slo_interface / 844e5dd86bce / 4

Type: `"number"`. Optional.

VLAN ID. VLAN ID

Upstream description:

VLAN ID

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 4095),
}
```

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```

<a id="canonical-b6866eb763847e21bec424ff7c2f5c04a8cc182b7d3bf708cb2318d959128a6a"></a>

## Next pages — layer2_interface.l2vlan_slo_interface / 844e5dd86bce / 5

- [layer2_interface](resources--network_interface--reference--group-002.md#canonical-1c594b57dc38583cede4438254bed4c980392fdf365967edc1ad24fac341cbd0)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-737b5fe469f062275aac4bc20ccf512c05ebce12de508b1729aab672e41fe4b7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-52c3f44a1a4622c821c56e057ca30c08a91c63189514d08fc738e2c6898365d9"></a>

## timeouts — timeouts / 3ca9de5ee088 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- timeouts

<a id="canonical-0f1a477aeb1c77ebf2be4f90b99bfacf1b54171a03d304b0c69a4f6c23e114e7"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-33ec8f39d21568249ad7975bc2105685759ed11de64b466fa650419712adb4d5"></a>

## Direct properties — timeouts / 3ca9de5ee088 / 3

<a id="canonical-efbe76065569c878d92de43cd3da30813ab33955833b53dd17984ae6a1481d8a"></a>

<a id="canonical-fb9f72c7692eab9d3c45d76c3b6662b83998c3dba3eccfa0623c7d2c64730b3e"></a>

## create property — timeouts / 3ca9de5ee088 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-68f60b2a82d750e732d9236e23839690ba7b53dc64a401ab97f41808c6f2f8e2"></a>

<a id="canonical-929c8210b91ec762d3ae55ea05f543862d71532dfae4cf06ce1238e022c4fb3b"></a>

## delete property — timeouts / 3ca9de5ee088 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1820c17b3dc561b882b296756e309f8b4c84ce3636eb55d3da302df03cc55dca"></a>

<a id="canonical-41f5a58dcd44393ad4cdb571e71f483cdcdc62698b85733e4556efe430e05ad2"></a>

## read property — timeouts / 3ca9de5ee088 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-c1a896e6be4dfe1ecbdfcbfd9b32da7905b4ce1feec017af921e88d22f226a67"></a>

<a id="canonical-56a65e907d7c5add906d3e5e99d5965013a31d7d2fde3c3987fe80cf476f6de2"></a>

## update property — timeouts / 3ca9de5ee088 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-86bd2e02eb66474635e19da1cd5db51ac71560feb58e5a0ab68b7adc20a350a9"></a>

## Next pages — timeouts / 3ca9de5ee088 / 8

- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-0db6b411298b0a3c07d71c63f1bb4732bc353ecea9c44af1f2a38335e20aeda0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6377bbcd8bbb0481d2e66f4749c94a8f3426194e8bd076e22faf93f898c8c50b"></a>

## tunnel_interface — tunnel_interface / 4a05e7b950c0 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- tunnel_interface

<a id="canonical-5b16fd24e80b66597695a9d58656efff23e866192b280fee851b04a157e01f46"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tunnel interface.

Upstream description:

Tunnel Interface Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site_local_inside_network",
    "site_local_network")}
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
  "x-ves-oneof-field-network_choice": "[\"site_local_inside_network\",\"site_local_network\"]",
  "x-ves-oneof-field-node_choice": "[\"node\"]"
}
```

Terraform syntax:

```terraform
tunnel_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-b78f0a942cbdedc178d22fe0dbbbb231ad20622b2b23490990ee6c7faf2f1689"></a>

## Direct properties — tunnel_interface / 4a05e7b950c0 / 3

<a id="canonical-00466f20e230b514fb5e765632a201547ae0631eecd690ce090f58a44e5fec93"></a>

<a id="canonical-5c38d807925d6bf0e38c11fa8d41e4ff0f8d41ac038ee045037a3b4d408b67c2"></a>

## mtu property — tunnel_interface / 4a05e7b950c0 / 4

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 512, Maximum: 9000},
  ),
}
```

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

<a id="canonical-b97e10cdb6f73a19ae05dd0dd9baacc1a9e6c7ffd0d4d3520371a9b4da0d4699"></a>

<a id="canonical-761d39bfc420cebac5072a2f2de8bb6ab42700f4cc231110a826e1bd88493ec5"></a>

## node property — tunnel_interface / 4a05e7b950c0 / 5

Type: `"string"`. Optional.

Exclusive with \[\] Configuration will apply to a given device on the given node.

Upstream description:

Exclusive with \[\] Configuration will apply to a given device on the given node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

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

<a id="canonical-8004221dfacd326c6a35ce6bbc434636bcfa2feee72258cedf9a70acd3557888"></a>

<a id="canonical-5edab0a5f205c5d6d5395726ca8b33da8578335e3c0c8c27033a67180a584655"></a>

## priority property — tunnel_interface / 4a05e7b950c0 / 6

Type: `"number"`. Optional.

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Upstream description:

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 255),
}
```

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

- [site_local_inside_network](resources--network_interface--reference--group-002.md#canonical-fe24a641579dba26ceef956a3bcd787a668b1fa6bc4806bd1757dfa62a35bea1): complete subsection reference.

- [site_local_network](resources--network_interface--reference--group-002.md#canonical-e65e284ca2fd8383cdcb85b0b5af04e9770a8efb8470c9e35db4df980cb397ce): complete subsection reference.

- [static_ip](resources--network_interface--reference--group-002.md#canonical-c0c6260098cafe3d4ead8ad59833429f3195ead158ba247d1863faac8f9fe7c1): complete subsection reference.

- [tunnel](resources--network_interface--reference--group-002.md#canonical-23e29a15e228b32b68180e88b6dd05a1cf270519330a85308fa907f776fd60db): complete subsection reference.

<a id="canonical-b07776af8dfb062876957ae22efc83be24d109ab50db3a3452811b88dbdb3f04"></a>

## Next pages — tunnel_interface / 4a05e7b950c0 / 7

- [tunnel_interface.site_local_inside_network](resources--network_interface--reference--group-002.md#canonical-fe24a641579dba26ceef956a3bcd787a668b1fa6bc4806bd1757dfa62a35bea1)
- [tunnel_interface.site_local_network](resources--network_interface--reference--group-002.md#canonical-e65e284ca2fd8383cdcb85b0b5af04e9770a8efb8470c9e35db4df980cb397ce)
- [tunnel_interface.static_ip](resources--network_interface--reference--group-002.md#canonical-c0c6260098cafe3d4ead8ad59833429f3195ead158ba247d1863faac8f9fe7c1)
- [tunnel_interface.tunnel](resources--network_interface--reference--group-002.md#canonical-23e29a15e228b32b68180e88b6dd05a1cf270519330a85308fa907f776fd60db)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-fe24a641579dba26ceef956a3bcd787a668b1fa6bc4806bd1757dfa62a35bea1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ea99fc04a70caa3fb17b088c2a0970143c184dbe7b1c7b259cec861f6212e76"></a>

## tunnel_interface.site_local_inside_network — tunnel_interface.site_local_inside_network / c845a78e707b / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [tunnel_interface](resources--network_interface--reference--group-002.md#canonical-0db6b411298b0a3c07d71c63f1bb4732bc353ecea9c44af1f2a38335e20aeda0)
- tunnel_interface.site_local_inside_network

<a id="canonical-c3a8d12380b8f911d1c554b5e7d05168e525a9aa92cd50c8cd19ab9f8fee4152"></a>

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
site_local_inside_network = {}
```

<a id="canonical-1f291d5a470ebdcc6181e4b65768a748ea392b55d1ad89891233c623e84567cd"></a>

## Direct properties — tunnel_interface.site_local_inside_network / c845a78e707b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-14f575743a4419b80dcf80ce9d9c44dcab48bbaeeead828ba9f0fd75f6905b51"></a>

## Next pages — tunnel_interface.site_local_inside_network / c845a78e707b / 4

- [tunnel_interface](resources--network_interface--reference--group-002.md#canonical-0db6b411298b0a3c07d71c63f1bb4732bc353ecea9c44af1f2a38335e20aeda0)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-e65e284ca2fd8383cdcb85b0b5af04e9770a8efb8470c9e35db4df980cb397ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e005538a3b1edbb7625f1ec9ffb899089c2220730b9ca249aa371ebe795ccd7"></a>

## tunnel_interface.site_local_network — tunnel_interface.site_local_network / 25b29eaf6401 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [tunnel_interface](resources--network_interface--reference--group-002.md#canonical-0db6b411298b0a3c07d71c63f1bb4732bc353ecea9c44af1f2a38335e20aeda0)
- tunnel_interface.site_local_network

<a id="canonical-7956099f3a05592b4aa819ed490ff9f8c57dd6c146544af693a9bde8fef480eb"></a>

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
site_local_network = {}
```

<a id="canonical-4041b394a34a4dff561233d24a50a6dad78e06c2b65fac3d01c050d7aaf2a549"></a>

## Direct properties — tunnel_interface.site_local_network / 25b29eaf6401 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9a73bf699d0303341f2c380a8df4ac39a484a0f585ce51ff8b4f47f3e1c49971"></a>

## Next pages — tunnel_interface.site_local_network / 25b29eaf6401 / 4

- [tunnel_interface](resources--network_interface--reference--group-002.md#canonical-0db6b411298b0a3c07d71c63f1bb4732bc353ecea9c44af1f2a38335e20aeda0)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-c0c6260098cafe3d4ead8ad59833429f3195ead158ba247d1863faac8f9fe7c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-25cab52ee2952808070c566e4bdafde6a78f374add5531456ebe65b418566c50"></a>

## tunnel_interface.static_ip — tunnel_interface.static_ip / d77276c89bfb / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [tunnel_interface](resources--network_interface--reference--group-002.md#canonical-0db6b411298b0a3c07d71c63f1bb4732bc353ecea9c44af1f2a38335e20aeda0)
- tunnel_interface.static_ip

<a id="canonical-c66df417706b2af56978e29f09c01f35fe11f353aaceca85d8f98614b9a23b14"></a>

Type: `"object"`. single nested block, Optional.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("cluster_static_ip",
    "node_static_ip")}
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
  "x-ves-oneof-field-network_prefix_choice": "[\"cluster_static_ip\",\"node_static_ip\"]"
}
```

Terraform syntax:

```terraform
static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-16689ce7d2f95a7f7b51c5e8dc1eab90ef733ff7134a887b8d33b548a727efdc"></a>

## Direct properties — tunnel_interface.static_ip / d77276c89bfb / 3

- [cluster_static_ip](resources--network_interface--reference--group-002.md#canonical-2e60ae8df722f192b60b11d8b9fac4fd88ad04b2d5bc0d8ac81dfef034475628): complete subsection reference.

- [node_static_ip](resources--network_interface--reference--group-002.md#canonical-7c59e099aec50b1c585c732adc6d44f73a79a6e8516081698d499442db1977fd): complete subsection reference.

<a id="canonical-dc629f2911c96fcab08eb6bf91b3f8a4af03e86b0dea5ed785ba8fcd5d39415c"></a>

## Next pages — tunnel_interface.static_ip / d77276c89bfb / 4

- [tunnel_interface.static_ip.cluster_static_ip](resources--network_interface--reference--group-002.md#canonical-2e60ae8df722f192b60b11d8b9fac4fd88ad04b2d5bc0d8ac81dfef034475628)
- [tunnel_interface.static_ip.node_static_ip](resources--network_interface--reference--group-002.md#canonical-7c59e099aec50b1c585c732adc6d44f73a79a6e8516081698d499442db1977fd)
- [tunnel_interface](resources--network_interface--reference--group-002.md#canonical-0db6b411298b0a3c07d71c63f1bb4732bc353ecea9c44af1f2a38335e20aeda0)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-2e60ae8df722f192b60b11d8b9fac4fd88ad04b2d5bc0d8ac81dfef034475628"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7e7ec0848d4d6317956390e87611fc994180650fd8cb1078702d62d9c4d76ff"></a>

## tunnel_interface.static_ip.cluster_static_ip — tunnel_interface.static_ip.cluster_static_ip / 8a2bb5d419d3 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [tunnel_interface](resources--network_interface--reference--group-002.md#canonical-0db6b411298b0a3c07d71c63f1bb4732bc353ecea9c44af1f2a38335e20aeda0)
- [tunnel_interface.static_ip](resources--network_interface--reference--group-002.md#canonical-c0c6260098cafe3d4ead8ad59833429f3195ead158ba247d1863faac8f9fe7c1)
- tunnel_interface.static_ip.cluster_static_ip

<a id="canonical-d32d6df37cd00022a747281774056c3c2448dc85d3dd169936f9edccbd0dc0a8"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for cluster.

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
cluster_static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-ed08276a510d11a0a9bd03a851832946a94fd04f3c64ec67d53b35465922cfba"></a>

## Direct properties — tunnel_interface.static_ip.cluster_static_ip / 8a2bb5d419d3 / 3

<a id="canonical-c9589b25857b35d35ec1255197c2a76caea57365ba94a268e8f6c31199fbb4e4"></a>

<a id="canonical-675552006c6e3d5891218117e2be2b33ac10c069a64f7ac8b2351c7d39815c0f"></a>

## interface_ip_map property — tunnel_interface.static_ip.cluster_static_ip / 8a2bb5d419d3 / 4

Type: `["map", "string"]`. Optional.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

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
    "ves.io.schema.rules.map.max_pairs": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  }
}
```

<a id="canonical-f97165dd6d43c5b5b0131547b75f33f65456b8d6df205b63374ace647feecc51"></a>

## Next pages — tunnel_interface.static_ip.cluster_static_ip / 8a2bb5d419d3 / 5

- [tunnel_interface.static_ip](resources--network_interface--reference--group-002.md#canonical-c0c6260098cafe3d4ead8ad59833429f3195ead158ba247d1863faac8f9fe7c1)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-7c59e099aec50b1c585c732adc6d44f73a79a6e8516081698d499442db1977fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-392f308f74169d5094917fbc75f7057c56b004b83c4982a92c5b0293544a99be"></a>

## tunnel_interface.static_ip.node_static_ip — tunnel_interface.static_ip.node_static_ip / a826968c8137 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [tunnel_interface](resources--network_interface--reference--group-002.md#canonical-0db6b411298b0a3c07d71c63f1bb4732bc353ecea9c44af1f2a38335e20aeda0)
- [tunnel_interface.static_ip](resources--network_interface--reference--group-002.md#canonical-c0c6260098cafe3d4ead8ad59833429f3195ead158ba247d1863faac8f9fe7c1)
- tunnel_interface.static_ip.node_static_ip

<a id="canonical-c7f4a5adbb982884dc3c87beb9e6027f8b8b1b2e2fcf5fd12dc0f2ecf74d3a1a"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for a node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_address")}
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
node_static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-72884968f5683975c80380c25fc25c234b82b8ceeffad6bd6ac4d6e3dcfa9fdf"></a>

## Direct properties — tunnel_interface.static_ip.node_static_ip / a826968c8137 / 3

<a id="canonical-593f543f94c55632fb933a6d7fb6819b32a92404ec944c944aede4065bbd5237"></a>

<a id="canonical-a2d0254405f0891dc19bf34c50a704b2cccd4b053e0be2ad70780caabec3dfa8"></a>

## default_gw property — tunnel_interface.static_ip.node_static_ip / a826968c8137 / 4

Type: `"string"`. Optional.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

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

<a id="canonical-957a383307f14b2b9837d2629ef9d09b9aa2e247b49dae28fe494f549522dfe3"></a>

<a id="canonical-43e55cdd007081903d7444fd40d379c48f8ef28dbedffba96f6c30c9429724d0"></a>

## dns_server property — tunnel_interface.static_ip.node_static_ip / a826968c8137 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-7b3bacbd7966d385fa96314f5958111887a3abebe8eef40d1eae6d56b59f43f2"></a>

<a id="canonical-923abb1c4b685e86cbf11bc9f77fe8e03f95cebf81c537e0257db898bcd673d4"></a>

## ip_address property — tunnel_interface.static_ip.node_static_ip / a826968c8137 / 6

Type: `"string"`. Optional.

IP address of the interface and prefix length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  validators.CIDRValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

<a id="canonical-f12e4d6e2375c88d046d7cd3d88ee7be4f7cd6d322a2140e4b8946bf6915b794"></a>

## Next pages — tunnel_interface.static_ip.node_static_ip / a826968c8137 / 7

- [tunnel_interface.static_ip](resources--network_interface--reference--group-002.md#canonical-c0c6260098cafe3d4ead8ad59833429f3195ead158ba247d1863faac8f9fe7c1)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-23e29a15e228b32b68180e88b6dd05a1cf270519330a85308fa907f776fd60db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2801e8476659442a4b70bfda30d67dcf6c9c5fa85048a84aceaffd1692925aa6"></a>

## tunnel_interface.tunnel — tunnel_interface.tunnel / 268f1e2f7eea / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [tunnel_interface](resources--network_interface--reference--group-002.md#canonical-0db6b411298b0a3c07d71c63f1bb4732bc353ecea9c44af1f2a38335e20aeda0)
- tunnel_interface.tunnel

<a id="canonical-e69e395b805d387e01859c146fbd0e0d20250fbc1c2a2ad34babd72602f38120"></a>

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
tunnel {
  # Configure direct properties listed below.
}
```

<a id="canonical-3574aa211451a792b7c6b0f48a4c51351e336b7d818864709f38033cab4701c0"></a>

## Direct properties — tunnel_interface.tunnel / 268f1e2f7eea / 3

<a id="canonical-e61f2ad9ebabdbf83f94b1596141a831b9e5cecbf4b5c938a01bfb91baaadc73"></a>

<a id="canonical-306cf5bc2e942c38dc3b2c7197476f14b39606025139199ca886d9e032a961e1"></a>

## name property — tunnel_interface.tunnel / 268f1e2f7eea / 4

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

<a id="canonical-60edc75cd60824678ec11b0ade93b552c21c20b6a2fd8c30c79fd4399a4776b2"></a>

<a id="canonical-868d79d97448b344b33a25ee62b2cc53ba664e66a943b2b2a9e51f14c0753774"></a>

## namespace property — tunnel_interface.tunnel / 268f1e2f7eea / 5

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

<a id="canonical-79f3b8880bcb9b273cacd00d3bff242f1a1b475c89a996bf494a1b77b9192551"></a>

<a id="canonical-c445d3069b9eebfb69dd1c7c625a2e90516a1d9ff600b9cc76acd8eda69bd2b3"></a>

## tenant property — tunnel_interface.tunnel / 268f1e2f7eea / 6

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

<a id="canonical-27ffab3d1387728586f57e8f91545cf75dc3e708b17fc827f8b5df9a0298c9c3"></a>

## Next pages — tunnel_interface.tunnel / 268f1e2f7eea / 7

- [tunnel_interface](resources--network_interface--reference--group-002.md#canonical-0db6b411298b0a3c07d71c63f1bb4732bc353ecea9c44af1f2a38335e20aeda0)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
