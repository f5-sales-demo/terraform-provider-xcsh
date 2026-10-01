---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-c9f5dff160a200e75386929f7946ddd83d00a58e9f97b977b723ef32338c8f9f"></a>

## Next pages — azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw / 8fa629d9735c / 4

- [azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-004.md#canonical-5eeccc42ebe5b289772b1c4a2a05221390d1a11974e23bd859c4938448633471)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-68223a91e96b6407954ba5471ea29ff204035a723304032dfe132f3790111ae9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef4f5d27034fc3fb66fbd7b63ee412363e40be8d9151152967155d3bb09d9649"></a>

## azure.not_managed.node_list.interface_list.dhcp_server.interface_ip_map — azure.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / c111ede98ffc / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [azure](resources--securemesh_site_v2--reference--group-004.md#canonical-d6b794d7e7e25165fa55b5be6623d8b4c48cd27cb72b40d9005227c301983678)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0579a21e4252f6dae8e786094b3ef46ed712e25dc0e83d91b29344662b25ec19)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-3f3f4bbb5b7bd2586cd71c68e3c3d54a6cfa6d08a40994c2f866af5461db8d70)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- [azure.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-004.md#canonical-ffc8513eb887917c7114a15fc180644828712262a926b980d5cf086e5e80cdb8)
- azure.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-653be44c5be6957e4e8f8eaef9da580d47951dcfadda4a611c05e74964188d8b"></a>

Type: `"object"`. single nested block, Optional.

Interface IPv4 Assignments. Specify static IPv4 addresses per node.

Upstream description:

Specify static IPv4 addresses per node.

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
interface_ip_map {
  # Configure direct properties listed below.
}
```

<a id="canonical-9685039c2e48a2bda5658c08800814bf548b755978777d857fd8ecc764d0f312"></a>

## Direct properties — azure.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / c111ede98ffc / 3

<a id="canonical-cb31caf0bd2bcab1f57cc9af56b7c8d962f8a326bd26f75c6667154438138189"></a>

<a id="canonical-2507d56de90284dbdad1a6a4bc82a118f02429328adf003df4e0821502e6ab5e"></a>

## interface_ip_map property — azure.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / c111ede98ffc / 4

Type: `["map", "string"]`. Optional.

Specify static IPv4 addresses per site:node.

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
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  }
}
```

<a id="canonical-22415da8b9992485002f97e774fa9d4f89ca9759167c684262e30cc91609f144"></a>

## Next pages — azure.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / c111ede98ffc / 5

- [azure.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-004.md#canonical-ffc8513eb887917c7114a15fc180644828712262a926b980d5cf086e5e80cdb8)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-1aa961a9732d538b0bc838ac403fe86f975808a158927540e476ad48dcda332f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8bc80ebeecab44740c6f6a98d842db455ba1872dae20874462955f0980e772e"></a>

## azure.not_managed.node_list.interface_list.ethernet_interface — azure.not_managed.node_list.interface_list.ethernet_interface / 2a120fea32cd / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [azure](resources--securemesh_site_v2--reference--group-004.md#canonical-d6b794d7e7e25165fa55b5be6623d8b4c48cd27cb72b40d9005227c301983678)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0579a21e4252f6dae8e786094b3ef46ed712e25dc0e83d91b29344662b25ec19)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-3f3f4bbb5b7bd2586cd71c68e3c3d54a6cfa6d08a40994c2f866af5461db8d70)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- azure.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-292aa326ab6d4f0858c25c41780479dbe656c92a70b7046e18e15b7c7fc0d2a3"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ethernet interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("mac")}
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
ethernet_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-874e86f4d3c3598f3d3e6aee8a3cf00d9d6b81841d8f59b63aa59362e0975d25"></a>

## Direct properties — azure.not_managed.node_list.interface_list.ethernet_interface / 2a120fea32cd / 3

<a id="canonical-3711503c8e59b5aadab8d9d01cd1030b40cceeefc4b3459b93913f00d98cf3f1"></a>

<a id="canonical-6803d8957da62b7472725c5ee6e333d0715e73f6357943b0e144e7a55faf2a8b"></a>

## device property — azure.not_managed.node_list.interface_list.ethernet_interface / 2a120fea32cd / 4

Type: `"string"`. Optional.

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

Upstream description:

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

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
    "create": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "false",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "false",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-24026db355d90053d6232f49eba946a3e1115ff312496c73849a9e87d13bede4"></a>

<a id="canonical-2991957be36063ef71653d95fd73b0a5503cc29440a3e994c0b6f5e07c55d4e8"></a>

## mac property — azure.not_managed.node_list.interface_list.ethernet_interface / 2a120fea32cd / 5

Type: `"string"`. Optional.

MAC Address. Configuration parameter for mac

Upstream description:

Configuration parameter for mac

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.MACValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "mac-address",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.mac": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.mac": "true"
  }
}
```

<a id="canonical-d16e118629330fef8c9be70a09ab30b64ca51a31ad44aa7f53aea57cf86a12b2"></a>

## Next pages — azure.not_managed.node_list.interface_list.ethernet_interface / 2a120fea32cd / 6

- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-3e65690606802c0fe1497637bd9e45a97cfd10496916a30646149db6f6c1f31d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-13bd22e082eef118fc4885d2ca71362a8d571ecbab166ed18c1117f1263c0eb2"></a>

## azure.not_managed.node_list.interface_list.ipv6_auto_config — azure.not_managed.node_list.interface_list.ipv6_auto_config / b8f943b8ca63 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [azure](resources--securemesh_site_v2--reference--group-004.md#canonical-d6b794d7e7e25165fa55b5be6623d8b4c48cd27cb72b40d9005227c301983678)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0579a21e4252f6dae8e786094b3ef46ed712e25dc0e83d91b29344662b25ec19)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-3f3f4bbb5b7bd2586cd71c68e3c3d54a6cfa6d08a40994c2f866af5461db8d70)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- azure.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-a06eddcd8e8d03ae54c689bc6062b34758052871c44bb0a8be4fb400ba324aac"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("host",
    "router")}
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
  "x-ves-oneof-field-autoconfig_choice": "[\"host\",\"router\"]"
}
```

Terraform syntax:

```terraform
ipv6_auto_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-52539a31ab2477e3706a3f1b2375b116707543a4dbfde3db5ae001eb93c4a065"></a>

## Direct properties — azure.not_managed.node_list.interface_list.ipv6_auto_config / b8f943b8ca63 / 3

- [host](resources--securemesh_site_v2--reference--group-005.md#canonical-a23f8e0ba4aa9be99e45e2369a216b78414f91274217ba6cc2931e4d2e499632): complete subsection reference.

- [router](resources--securemesh_site_v2--reference--group-005.md#canonical-d9d8d177f0b7dd20508380b4cc4396190c509f55901ad0a24378d73584854b09): complete subsection reference.

<a id="canonical-8ce892b7cb099f0f1eadb47d3a2ba18712bae2df6502fd96f8821191b1c8c2eb"></a>

## Next pages — azure.not_managed.node_list.interface_list.ipv6_auto_config / b8f943b8ca63 / 4

- [azure.not_managed.node_list.interface_list.ipv6_auto_config.host](resources--securemesh_site_v2--reference--group-005.md#canonical-a23f8e0ba4aa9be99e45e2369a216b78414f91274217ba6cc2931e4d2e499632)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-005.md#canonical-d9d8d177f0b7dd20508380b4cc4396190c509f55901ad0a24378d73584854b09)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-a23f8e0ba4aa9be99e45e2369a216b78414f91274217ba6cc2931e4d2e499632"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d5fa4d44944ed15a4f63bef6e8e88a83a9c87ea33f6e871500925d65ef3df428"></a>

## azure.not_managed.node_list.interface_list.ipv6_auto_config.host — azure.not_managed.node_list.interface_list.ipv6_auto_config.host / 31d00c52ba53 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [azure](resources--securemesh_site_v2--reference--group-004.md#canonical-d6b794d7e7e25165fa55b5be6623d8b4c48cd27cb72b40d9005227c301983678)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0579a21e4252f6dae8e786094b3ef46ed712e25dc0e83d91b29344662b25ec19)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-3f3f4bbb5b7bd2586cd71c68e3c3d54a6cfa6d08a40994c2f866af5461db8d70)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-005.md#canonical-3e65690606802c0fe1497637bd9e45a97cfd10496916a30646149db6f6c1f31d)
- azure.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-4828da1738cbf99cb724d3d4f7883d9146b9609077f3ea8d4ba994685ffc1f4e"></a>

Type: `["object", {}]`. Optional.

Hostname or IP address of the target server.

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
host = {}
```

<a id="canonical-6e897319118b09ac838a1940c96d7c81b8e8b4fab902cc90689708c0be8dd7fe"></a>

## Direct properties — azure.not_managed.node_list.interface_list.ipv6_auto_config.host / 31d00c52ba53 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-664bc5cb894fee1eea26d7d8f25b5076c99a7d94c91ef5146de92d722a2a1dbc"></a>

## Next pages — azure.not_managed.node_list.interface_list.ipv6_auto_config.host / 31d00c52ba53 / 4

- [azure.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-005.md#canonical-3e65690606802c0fe1497637bd9e45a97cfd10496916a30646149db6f6c1f31d)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-d9d8d177f0b7dd20508380b4cc4396190c509f55901ad0a24378d73584854b09"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-87341cf2f7dbae01967f8f25f4255277e14f66a5ee5dab73bb42d5d1fd9210da"></a>

## azure.not_managed.node_list.interface_list.ipv6_auto_config.router — azure.not_managed.node_list.interface_list.ipv6_auto_config.router / f0f1d00a4705 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [azure](resources--securemesh_site_v2--reference--group-004.md#canonical-d6b794d7e7e25165fa55b5be6623d8b4c48cd27cb72b40d9005227c301983678)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0579a21e4252f6dae8e786094b3ef46ed712e25dc0e83d91b29344662b25ec19)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-3f3f4bbb5b7bd2586cd71c68e3c3d54a6cfa6d08a40994c2f866af5461db8d70)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-005.md#canonical-3e65690606802c0fe1497637bd9e45a97cfd10496916a30646149db6f6c1f31d)
- azure.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-ced71d8eee2485a52fe278f005b5fb23b8f20f09526c5904f98cf49ae5e82ea0"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigRouterType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("network_prefix",
    "stateful")}
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
  "x-ves-oneof-field-address_choice": "[\"network_prefix\",\"stateful\"]"
}
```

Terraform syntax:

```terraform
router {
  # Configure direct properties listed below.
}
```

<a id="canonical-446394f03d238ae213cd15f72d1233d501f24986e112655d50ee7669865ab018"></a>

## Direct properties — azure.not_managed.node_list.interface_list.ipv6_auto_config.router / f0f1d00a4705 / 3

- [dns_config](resources--securemesh_site_v2--reference--group-005.md#canonical-5b609900e751bb4098290597017b601b32cd7cd14d0e22c1900cb89f2442299d): complete subsection reference.

<a id="canonical-649505a9fa985456bcc0a68d5b44dd6e41aba474d77bf1b36aba22ee51fa6cf7"></a>

<a id="canonical-95d4e8f4bc097d82e5b744873b43ee71233c39eecec8abb1d9a3bb82c5cce556"></a>

## network_prefix property — azure.not_managed.node_list.interface_list.ipv6_auto_config.router / f0f1d00a4705 / 4

Type: `"string"`. Optional.

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Upstream description:

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

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
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": ".*::/64$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true",
    "ves.io.schema.rules.string.pattern": ".*::/64$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true",
    "ves.io.schema.rules.string.pattern": ".*::/64$"
  }
}
```

- [stateful](resources--securemesh_site_v2--reference--group-005.md#canonical-0f9cf03213bb34e70ea118fe43171363612f1ff1b2f74f79b202fc9ee79a8150): complete subsection reference.

<a id="canonical-53af414d6411672b07f04256ae13441ca57a3a84016f076ae169a29ee21d06ca"></a>

## Next pages — azure.not_managed.node_list.interface_list.ipv6_auto_config.router / f0f1d00a4705 / 5

- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-005.md#canonical-5b609900e751bb4098290597017b601b32cd7cd14d0e22c1900cb89f2442299d)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-005.md#canonical-0f9cf03213bb34e70ea118fe43171363612f1ff1b2f74f79b202fc9ee79a8150)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-005.md#canonical-3e65690606802c0fe1497637bd9e45a97cfd10496916a30646149db6f6c1f31d)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-5b609900e751bb4098290597017b601b32cd7cd14d0e22c1900cb89f2442299d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cb72dd82bb2b02aaccbe6e98a974dac6fec1c0b81f475ba60543b64a3da5a950"></a>

## azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / 53bf1a2f8004 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [azure](resources--securemesh_site_v2--reference--group-004.md#canonical-d6b794d7e7e25165fa55b5be6623d8b4c48cd27cb72b40d9005227c301983678)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0579a21e4252f6dae8e786094b3ef46ed712e25dc0e83d91b29344662b25ec19)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-3f3f4bbb5b7bd2586cd71c68e3c3d54a6cfa6d08a40994c2f866af5461db8d70)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-005.md#canonical-3e65690606802c0fe1497637bd9e45a97cfd10496916a30646149db6f6c1f31d)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-005.md#canonical-d9d8d177f0b7dd20508380b4cc4396190c509f55901ad0a24378d73584854b09)
- azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-14ace497bb2e75dfa10b82ab137644c8205c271a0b603b8359ffb23786e2eb1d"></a>

Type: `"object"`. single nested block, Optional.

IPV6DnsConfig.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("configured_list",
    "local_dns")}
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
  "x-ves-oneof-field-dns_choice": "[\"configured_list\",\"local_dns\"]"
}
```

Terraform syntax:

```terraform
dns_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-0498b2b50493261f349dda192ab110a2d11236d40bc399ae51f9effebe1e3081"></a>

## Direct properties — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / 53bf1a2f8004 / 3

- [configured_list](resources--securemesh_site_v2--reference--group-005.md#canonical-8af432d9941cd6fca090aca2b33deac37be90b3316e39eb94f6032c514945d4f): complete subsection reference.

- [local_dns](resources--securemesh_site_v2--reference--group-005.md#canonical-ffa24ffca4360e8e7f11946afa60b21f2f7ccbada7f67167a2d6af6adea992c1): complete subsection reference.

<a id="canonical-4d8cc83d6f8ba1865d498c5c1e0d4868e1c3943273014330f87ce0410793659f"></a>

## Next pages — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / 53bf1a2f8004 / 4

- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list](resources--securemesh_site_v2--reference--group-005.md#canonical-8af432d9941cd6fca090aca2b33deac37be90b3316e39eb94f6032c514945d4f)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-005.md#canonical-ffa24ffca4360e8e7f11946afa60b21f2f7ccbada7f67167a2d6af6adea992c1)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-005.md#canonical-d9d8d177f0b7dd20508380b4cc4396190c509f55901ad0a24378d73584854b09)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-8af432d9941cd6fca090aca2b33deac37be90b3316e39eb94f6032c514945d4f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ef8545cab6063baa12e25cc4612e4b1b53a32fd839ada84459b67d3fc4c4424"></a>

## azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.co / 5f2568796c2f / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [azure](resources--securemesh_site_v2--reference--group-004.md#canonical-d6b794d7e7e25165fa55b5be6623d8b4c48cd27cb72b40d9005227c301983678)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0579a21e4252f6dae8e786094b3ef46ed712e25dc0e83d91b29344662b25ec19)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-3f3f4bbb5b7bd2586cd71c68e3c3d54a6cfa6d08a40994c2f866af5461db8d70)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-005.md#canonical-3e65690606802c0fe1497637bd9e45a97cfd10496916a30646149db6f6c1f31d)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-005.md#canonical-d9d8d177f0b7dd20508380b4cc4396190c509f55901ad0a24378d73584854b09)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-005.md#canonical-5b609900e751bb4098290597017b601b32cd7cd14d0e22c1900cb89f2442299d)
- azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-aebfd9f8b4830e92483265dbacc34e7c0e553a40656ba28a7c27cc0555bdb603"></a>

Type: `"object"`. single nested block, Optional.

IPV6DnsList.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_list")}
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
configured_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-157b2d6ccb195b2a30c5988597587f052218194e49f30d137e152c1cc94f879f"></a>

## Direct properties — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.co / 5f2568796c2f / 3

<a id="canonical-6e69a1c3564a295cf76673270d62db3923538ca1514fc667b4d9e5b5cecba1e4"></a>

<a id="canonical-d3f59bbb9435ff2dcefd10ba09442d2c28d9a27daaeaea85570ca3c04b4a27a3"></a>

## dns_list property — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.co / 5f2568796c2f / 4

Type: `["list", "string"]`. Optional.

List of IPv6 Addresses acting as DNS servers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-28906fc9da13c4d4d3ed9373e077dea25df396965b717b1f22c7eec9f1e37542"></a>

## Next pages — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.co / 5f2568796c2f / 5

- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-005.md#canonical-5b609900e751bb4098290597017b601b32cd7cd14d0e22c1900cb89f2442299d)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-ffa24ffca4360e8e7f11946afa60b21f2f7ccbada7f67167a2d6af6adea992c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e945f58cac51ab593009ea9f6f2f6460a26a352c6ae9ec6f68d09897e4da621"></a>

## azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.lo / f81df98641d9 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [azure](resources--securemesh_site_v2--reference--group-004.md#canonical-d6b794d7e7e25165fa55b5be6623d8b4c48cd27cb72b40d9005227c301983678)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0579a21e4252f6dae8e786094b3ef46ed712e25dc0e83d91b29344662b25ec19)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-3f3f4bbb5b7bd2586cd71c68e3c3d54a6cfa6d08a40994c2f866af5461db8d70)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-005.md#canonical-3e65690606802c0fe1497637bd9e45a97cfd10496916a30646149db6f6c1f31d)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-005.md#canonical-d9d8d177f0b7dd20508380b4cc4396190c509f55901ad0a24378d73584854b09)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-005.md#canonical-5b609900e751bb4098290597017b601b32cd7cd14d0e22c1900cb89f2442299d)
- azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-5d124ba7fc534e345f439fc11707abb7c6b695f782fe6dacec747affe8e69ae3"></a>

Type: `"object"`. single nested block, Optional.

IPV6LocalDnsAddress.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("configured_address",
    "first_address"),
  validators.ConflictingObjectAttributes("configured_address",
    "last_address"),
  validators.ConflictingObjectAttributes("first_address",
    "last_address")}
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
  "x-ves-oneof-field-local_dns_choice": "[\"configured_address\",\"first_address\",\"last_address\"]"
}
```

Terraform syntax:

```terraform
local_dns {
  # Configure direct properties listed below.
}
```

<a id="canonical-b3ad298ab1e8ddce223452c282a76c228d139f5285f566ce66ffc99baf2f6d13"></a>

## Direct properties — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.lo / f81df98641d9 / 3

<a id="canonical-204f8bb90b8ab70cf048663c405ab6c7ae81638b496fbb82876da7fa11df0cdf"></a>

<a id="canonical-085da904b1ebed973350afbf63b70d26845ef188b8d1bbeebdd4ec14ae307bdd"></a>

## configured_address property — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.lo / f81df98641d9 / 4

Type: `"string"`. Optional.

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Upstream description:

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

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

- [first_address](resources--securemesh_site_v2--reference--group-005.md#canonical-b28a5d6ae3d5ca7fbf0975775bc15301d2fd2727a6171c392ce0c5a3e95f93a3): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-005.md#canonical-8c1bb90800860b2bac4586d007c7081245b4eb3f544dc5fc04b34242f2969540): complete subsection reference.

<a id="canonical-b2f0d1245cc0ad1a21d794b0ed2a1a0a89b641872cb97325b67d107576655f54"></a>

## Next pages — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.lo / f81df98641d9 / 5

- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address](resources--securemesh_site_v2--reference--group-005.md#canonical-b28a5d6ae3d5ca7fbf0975775bc15301d2fd2727a6171c392ce0c5a3e95f93a3)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address](resources--securemesh_site_v2--reference--group-005.md#canonical-8c1bb90800860b2bac4586d007c7081245b4eb3f544dc5fc04b34242f2969540)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-005.md#canonical-5b609900e751bb4098290597017b601b32cd7cd14d0e22c1900cb89f2442299d)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-b28a5d6ae3d5ca7fbf0975775bc15301d2fd2727a6171c392ce0c5a3e95f93a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330266495d921f8a91160445e3ef07972bc14702c95a14335acb2d1a4099faf"></a>

## azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.lo / 248cce4c27bc / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [azure](resources--securemesh_site_v2--reference--group-004.md#canonical-d6b794d7e7e25165fa55b5be6623d8b4c48cd27cb72b40d9005227c301983678)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0579a21e4252f6dae8e786094b3ef46ed712e25dc0e83d91b29344662b25ec19)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-3f3f4bbb5b7bd2586cd71c68e3c3d54a6cfa6d08a40994c2f866af5461db8d70)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-005.md#canonical-3e65690606802c0fe1497637bd9e45a97cfd10496916a30646149db6f6c1f31d)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-005.md#canonical-d9d8d177f0b7dd20508380b4cc4396190c509f55901ad0a24378d73584854b09)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-005.md#canonical-5b609900e751bb4098290597017b601b32cd7cd14d0e22c1900cb89f2442299d)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-005.md#canonical-ffa24ffca4360e8e7f11946afa60b21f2f7ccbada7f67167a2d6af6adea992c1)
- azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-a5f0496ee1291d03a4e60a1b29f42d9cd705e3704342c2c71c97634002dafbe1"></a>

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
first_address = {}
```

<a id="canonical-15b880e6658082cb6dbaf8a65d52f10b1071fd7ab7b6731de666e595ae15feb6"></a>

## Direct properties — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.lo / 248cce4c27bc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f1ee8f3aa3f72ed5d677d03d5c511b0bd34a312782e04884c100cbda72153968"></a>

## Next pages — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.lo / 248cce4c27bc / 4

- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-005.md#canonical-ffa24ffca4360e8e7f11946afa60b21f2f7ccbada7f67167a2d6af6adea992c1)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-8c1bb90800860b2bac4586d007c7081245b4eb3f544dc5fc04b34242f2969540"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-87bdf6ddf50cdb4f589d7ed70e0a5087f82729b9a646710fe4d5242baa9a9297"></a>

## azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.lo / 7da9de394f01 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [azure](resources--securemesh_site_v2--reference--group-004.md#canonical-d6b794d7e7e25165fa55b5be6623d8b4c48cd27cb72b40d9005227c301983678)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0579a21e4252f6dae8e786094b3ef46ed712e25dc0e83d91b29344662b25ec19)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-3f3f4bbb5b7bd2586cd71c68e3c3d54a6cfa6d08a40994c2f866af5461db8d70)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-005.md#canonical-3e65690606802c0fe1497637bd9e45a97cfd10496916a30646149db6f6c1f31d)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-005.md#canonical-d9d8d177f0b7dd20508380b4cc4396190c509f55901ad0a24378d73584854b09)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-005.md#canonical-5b609900e751bb4098290597017b601b32cd7cd14d0e22c1900cb89f2442299d)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-005.md#canonical-ffa24ffca4360e8e7f11946afa60b21f2f7ccbada7f67167a2d6af6adea992c1)
- azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-f936c1a6a675eb1fe026d3beef9f1068b8822ed01f1cd1c8234c72eb7e176927"></a>

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
last_address = {}
```

<a id="canonical-e46a8928dcae438d02959dc921a271fd6807cfa85a843390d1537c7fc790d362"></a>

## Direct properties — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.lo / 7da9de394f01 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-14a7ed6c4a227495ea6cd6a14ef0bbf7cadef0bd74c8b2de9b7b470458a4824f"></a>

## Next pages — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.lo / 7da9de394f01 / 4

- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-005.md#canonical-ffa24ffca4360e8e7f11946afa60b21f2f7ccbada7f67167a2d6af6adea992c1)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-0f9cf03213bb34e70ea118fe43171363612f1ff1b2f74f79b202fc9ee79a8150"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d9c7a1fc0136a9b6550cbeefd2986c48f78e1c39aad0a67744ddf2af12a7014e"></a>

## azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / f1ba8c621464 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [azure](resources--securemesh_site_v2--reference--group-004.md#canonical-d6b794d7e7e25165fa55b5be6623d8b4c48cd27cb72b40d9005227c301983678)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0579a21e4252f6dae8e786094b3ef46ed712e25dc0e83d91b29344662b25ec19)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-3f3f4bbb5b7bd2586cd71c68e3c3d54a6cfa6d08a40994c2f866af5461db8d70)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-005.md#canonical-3e65690606802c0fe1497637bd9e45a97cfd10496916a30646149db6f6c1f31d)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-005.md#canonical-d9d8d177f0b7dd20508380b4cc4396190c509f55901ad0a24378d73584854b09)
- azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-873b6f23fa13ddfb55736181ed3da95c638726b5ebc75d1caaeb9f0a91463d86"></a>

Type: `"object"`. single nested block, Optional.

DHCPIPV6 Stateful Server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dhcp_networks"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "automatic_from_start"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "interface_ip_map"),
  validators.ConflictingObjectAttributes("automatic_from_start",
    "interface_ip_map")}
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
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

Terraform syntax:

```terraform
stateful {
  # Configure direct properties listed below.
}
```

<a id="canonical-d10d7c7ac20313d7f98de335dbd8e67f2ba58c2dd82403d75c9d389ff83a33fc"></a>

## Direct properties — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / f1ba8c621464 / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-005.md#canonical-0cb87d1b99c84e9d73090efcaed2ce18874adaa11fbbda86dd6492c495b8dca7): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-005.md#canonical-115653272e91b5baedc6cf74cb5813eecc7eeca782073684edfe3f1ade13063a): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-005.md#canonical-5e0b954477b9bbbe7801ec2addcd8f21fef61aab326686f981732f15ffad742f): complete subsection reference.

<a id="canonical-d18e3ed7d2029201200f0bd7f6d9699e5ebe9b2f223df2033163d280876ca213"></a>

<a id="canonical-2edcfbf404dd8581ecbac4f5d85971bf39bdedd017d2df104d4a64033eb42ae9"></a>

## fixed_ip_map property — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / f1ba8c621464 / 4

Type: `["map", "string"]`. Optional.

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

Upstream description:

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

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
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  }
}
```

- [interface_ip_map](resources--securemesh_site_v2--reference--group-005.md#canonical-3e0d77daa50034f58440b28b6dae0b082305e6ca51ee57581a9b5c208108633a): complete subsection reference.

<a id="canonical-b026dd5cce0312f49f83a06a84ba873385b5b12b60cd916f869966eaf2055e0e"></a>

## Next pages — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / f1ba8c621464 / 5

- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end](resources--securemesh_site_v2--reference--group-005.md#canonical-0cb87d1b99c84e9d73090efcaed2ce18874adaa11fbbda86dd6492c495b8dca7)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start](resources--securemesh_site_v2--reference--group-005.md#canonical-115653272e91b5baedc6cf74cb5813eecc7eeca782073684edfe3f1ade13063a)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-005.md#canonical-5e0b954477b9bbbe7801ec2addcd8f21fef61aab326686f981732f15ffad742f)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map](resources--securemesh_site_v2--reference--group-005.md#canonical-3e0d77daa50034f58440b28b6dae0b082305e6ca51ee57581a9b5c208108633a)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-005.md#canonical-d9d8d177f0b7dd20508380b4cc4396190c509f55901ad0a24378d73584854b09)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-0cb87d1b99c84e9d73090efcaed2ce18874adaa11fbbda86dd6492c495b8dca7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-97de9b92e82bc4aeaa5dcf79e6f4baab60128289e23f5fc695a330a97047731f"></a>

## azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.auto / 62632b83b49c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [azure](resources--securemesh_site_v2--reference--group-004.md#canonical-d6b794d7e7e25165fa55b5be6623d8b4c48cd27cb72b40d9005227c301983678)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0579a21e4252f6dae8e786094b3ef46ed712e25dc0e83d91b29344662b25ec19)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-3f3f4bbb5b7bd2586cd71c68e3c3d54a6cfa6d08a40994c2f866af5461db8d70)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-005.md#canonical-3e65690606802c0fe1497637bd9e45a97cfd10496916a30646149db6f6c1f31d)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-005.md#canonical-d9d8d177f0b7dd20508380b4cc4396190c509f55901ad0a24378d73584854b09)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-005.md#canonical-0f9cf03213bb34e70ea118fe43171363612f1ff1b2f74f79b202fc9ee79a8150)
- azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-b2d9686c31442a8ac015b3d2972e9d01df9fd59155b70414ee456a1da0131b3d"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
automatic_from_end = {}
```

<a id="canonical-4336fe913a981dbcbd8617c6462da8e2423f8bb92c7e19e5249cec574de061e5"></a>

## Direct properties — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.auto / 62632b83b49c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-76d6e44b55779c049b04097325f5df80e20ed3b854532ca9bfdae8d3b00af113"></a>

## Next pages — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.auto / 62632b83b49c / 4

- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-005.md#canonical-0f9cf03213bb34e70ea118fe43171363612f1ff1b2f74f79b202fc9ee79a8150)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-115653272e91b5baedc6cf74cb5813eecc7eeca782073684edfe3f1ade13063a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e6a206abc353a09bfacaf4eb438beb6526f8ebce722c6c9f5c2bb5e58b858174"></a>

## azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.auto / d20b7741c9d5 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [azure](resources--securemesh_site_v2--reference--group-004.md#canonical-d6b794d7e7e25165fa55b5be6623d8b4c48cd27cb72b40d9005227c301983678)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0579a21e4252f6dae8e786094b3ef46ed712e25dc0e83d91b29344662b25ec19)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-3f3f4bbb5b7bd2586cd71c68e3c3d54a6cfa6d08a40994c2f866af5461db8d70)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-005.md#canonical-3e65690606802c0fe1497637bd9e45a97cfd10496916a30646149db6f6c1f31d)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-005.md#canonical-d9d8d177f0b7dd20508380b4cc4396190c509f55901ad0a24378d73584854b09)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-005.md#canonical-0f9cf03213bb34e70ea118fe43171363612f1ff1b2f74f79b202fc9ee79a8150)
- azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-355c1d57edf6aef7a8d0556a2d53b847aa367018a5670ad51a9cf30f9eae78e9"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
automatic_from_start = {}
```

<a id="canonical-9b7cf6ef7bbd89ec349650d209f4ad2a626968dbdb6a39787eaada1bfbf84f20"></a>

## Direct properties — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.auto / d20b7741c9d5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-43f042dee7cea476a383b3ddc8d5847599d97a1cfcd3eb9081181191e74a91d2"></a>

## Next pages — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.auto / d20b7741c9d5 / 4

- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-005.md#canonical-0f9cf03213bb34e70ea118fe43171363612f1ff1b2f74f79b202fc9ee79a8150)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-5e0b954477b9bbbe7801ec2addcd8f21fef61aab326686f981732f15ffad742f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7cc69a5682e1233ea6e334d58f3c154f8f1c2236443f4d013fdb91692a893f66"></a>

## azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp / 518b559431d3 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [azure](resources--securemesh_site_v2--reference--group-004.md#canonical-d6b794d7e7e25165fa55b5be6623d8b4c48cd27cb72b40d9005227c301983678)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0579a21e4252f6dae8e786094b3ef46ed712e25dc0e83d91b29344662b25ec19)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-3f3f4bbb5b7bd2586cd71c68e3c3d54a6cfa6d08a40994c2f866af5461db8d70)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-005.md#canonical-3e65690606802c0fe1497637bd9e45a97cfd10496916a30646149db6f6c1f31d)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-005.md#canonical-d9d8d177f0b7dd20508380b4cc4396190c509f55901ad0a24378d73584854b09)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-005.md#canonical-0f9cf03213bb34e70ea118fe43171363612f1ff1b2f74f79b202fc9ee79a8150)
- azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-5eadf4096b95bbf13cfecbaa73952d372d2d4945c8f951e819232ec666d34efd"></a>

Type: `"object"`. list nested block, Optional.

List of networks from which DHCP server can allocate IP addresses.

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

Terraform syntax:

```terraform
dhcp_networks {
  # Configure direct properties listed below.
}
```

<a id="canonical-ab059a81396214e7ee5b26eebe8b37abd76ad1ea5ebe8772864074c454d2539c"></a>

## Direct properties — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp / 518b559431d3 / 3

<a id="canonical-3200178b8cc643916e197b2deb867c5201fcb255c760ec968b7b6dda09d87800"></a>

<a id="canonical-6cfef2527808201b19caddec31228710901970016fe61af98c0d3744bb8cf9b9"></a>

## network_prefix property — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp / 518b559431d3 / 4

Type: `"string"`. Optional.

Exclusive with \[\] Network Prefix to be used for IPv6 address auto configuration.

Upstream description:

Exclusive with \[\] Network Prefix to be used for IPv6 address auto configuration.

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
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

<a id="canonical-6d96fdbf72c032784444ab53fa13c534903082b539562e52a3abe2380dd0c13f"></a>

<a id="canonical-5a4702e5ceb92cef4147cea81fcbd12afca4a1f9be667514cd080c68e6c8a46e"></a>

## pool_settings property — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp / 518b559431d3 / 5

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"),
}
```

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

- [pools](resources--securemesh_site_v2--reference--group-005.md#canonical-56eb8a66355fe752c61952255a53b4182a4cd4023805b9b8761869dc661801b2): complete subsection reference.

<a id="canonical-e4702d93f4fb099e95dccfb6d5266afbd2aeeb4ceae75e0bc78c64f67b0d0d21"></a>

## Next pages — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp / 518b559431d3 / 6

- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools](resources--securemesh_site_v2--reference--group-005.md#canonical-56eb8a66355fe752c61952255a53b4182a4cd4023805b9b8761869dc661801b2)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-005.md#canonical-0f9cf03213bb34e70ea118fe43171363612f1ff1b2f74f79b202fc9ee79a8150)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-56eb8a66355fe752c61952255a53b4182a4cd4023805b9b8761869dc661801b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2938962a52d9ee74b52b1508443fad06f653b04ed082e838828ba76edfb3b4a8"></a>

## azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp / 837072e1e15c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [azure](resources--securemesh_site_v2--reference--group-004.md#canonical-d6b794d7e7e25165fa55b5be6623d8b4c48cd27cb72b40d9005227c301983678)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0579a21e4252f6dae8e786094b3ef46ed712e25dc0e83d91b29344662b25ec19)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-3f3f4bbb5b7bd2586cd71c68e3c3d54a6cfa6d08a40994c2f866af5461db8d70)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-005.md#canonical-3e65690606802c0fe1497637bd9e45a97cfd10496916a30646149db6f6c1f31d)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-005.md#canonical-d9d8d177f0b7dd20508380b4cc4396190c509f55901ad0a24378d73584854b09)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-005.md#canonical-0f9cf03213bb34e70ea118fe43171363612f1ff1b2f74f79b202fc9ee79a8150)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-005.md#canonical-5e0b954477b9bbbe7801ec2addcd8f21fef61aab326686f981732f15ffad742f)
- azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-55d1deec6a5cadecb17723fff88381232b2e400a64a4e19bd586f04b2507e543"></a>

Type: `"object"`. list nested block, Optional.

List of non overlapping IP address ranges.

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
    "minItems": 1,
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

Terraform syntax:

```terraform
pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-9ae1b3450e65ac58f6bc5533125c41e541c84aaa60ca0e0c962415e209bd72f4"></a>

## Direct properties — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp / 837072e1e15c / 3

<a id="canonical-104da9aca75a816aed9ed342711df605eb71c09f7c7075b831af03db2bdd263f"></a>

<a id="canonical-e70044b47099e183c6e74229ddae59f0ebc78c66fec1e7128b84f16876d12cd0"></a>

## end_ip property — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp / 837072e1e15c / 4

Type: `"string"`. Optional.

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

Upstream description:

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

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

<a id="canonical-9f4327097ffeb6b7c4661959bc11f4f461b3095c413449e2c2ea7c5b1429fe33"></a>

<a id="canonical-5791f2197939ae02385dd6402febb47714623c6f346e803b763678c2908e56cc"></a>

## start_ip property — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp / 837072e1e15c / 5

Type: `"string"`. Optional.

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

Upstream description:

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

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

<a id="canonical-6c41fdb8b7911924cc937c995d67bebc01a1a816ae29750d31fd484127555af9"></a>

## Next pages — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp / 837072e1e15c / 6

- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-005.md#canonical-5e0b954477b9bbbe7801ec2addcd8f21fef61aab326686f981732f15ffad742f)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-3e0d77daa50034f58440b28b6dae0b082305e6ca51ee57581a9b5c208108633a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8eabd7bb18a58ae0a16b121500879184c6ea358a3deb1fcc1d6719f6ee137903"></a>

## azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.inte / 787799da3421 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [azure](resources--securemesh_site_v2--reference--group-004.md#canonical-d6b794d7e7e25165fa55b5be6623d8b4c48cd27cb72b40d9005227c301983678)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0579a21e4252f6dae8e786094b3ef46ed712e25dc0e83d91b29344662b25ec19)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-3f3f4bbb5b7bd2586cd71c68e3c3d54a6cfa6d08a40994c2f866af5461db8d70)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-005.md#canonical-3e65690606802c0fe1497637bd9e45a97cfd10496916a30646149db6f6c1f31d)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-005.md#canonical-d9d8d177f0b7dd20508380b4cc4396190c509f55901ad0a24378d73584854b09)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-005.md#canonical-0f9cf03213bb34e70ea118fe43171363612f1ff1b2f74f79b202fc9ee79a8150)
- azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-6b4b1ae685c4c03a6f26d5d1675c99efc5a9c81f9187c7d4093c9483d44b7783"></a>

Type: `"object"`. single nested block, Optional.

Map of Interface IPv6 assignments per node.

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
interface_ip_map {
  # Configure direct properties listed below.
}
```

<a id="canonical-cf0285a3597afdd789291daab45e36c815d646c871b21dbc2691eda0325b95f4"></a>

## Direct properties — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.inte / 787799da3421 / 3

<a id="canonical-fc7150197f1289b7a4b583799bb40b07d0ddcd0e3a97d21182d77eabf8a26b40"></a>

<a id="canonical-b593ed3decca50380a73911f5ca4db3bce8c2cc9deb8e248b38d8241d9114a0e"></a>

## interface_ip_map property — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.inte / 787799da3421 / 4

Type: `["map", "string"]`. Optional.

Site:Node to IPv6 Mapping. Map of Site:Node to IPv6 address.

Upstream description:

Map of Site:Node to IPv6 address.

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
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  }
}
```

<a id="canonical-20eb661e79d7f0eef49aeaf7fad355773a5f11cb17804a94501404b2249691ab"></a>

## Next pages — azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.inte / 787799da3421 / 5

- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-005.md#canonical-0f9cf03213bb34e70ea118fe43171363612f1ff1b2f74f79b202fc9ee79a8150)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-06fbe0a0d25735b47e3cef7ab779eacaeb8789620c48a02a7719402c358faae3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd827ef3fa2f5680f3fbe19f9acff0b9a9e4e8e6c4ab5767ed7998649691c71a"></a>

## azure.not_managed.node_list.interface_list.monitor — azure.not_managed.node_list.interface_list.monitor / 84b4cae226fb / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [azure](resources--securemesh_site_v2--reference--group-004.md#canonical-d6b794d7e7e25165fa55b5be6623d8b4c48cd27cb72b40d9005227c301983678)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0579a21e4252f6dae8e786094b3ef46ed712e25dc0e83d91b29344662b25ec19)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-3f3f4bbb5b7bd2586cd71c68e3c3d54a6cfa6d08a40994c2f866af5461db8d70)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- azure.not_managed.node_list.interface_list.monitor

<a id="canonical-e4ac9b5c7cc8f111f551bca0bcb731aabbf46200ac3a79f1fe51eff3e7ee6d57"></a>

Type: `["object", {}]`. Optional.

Link Quality Monitoring configuration for a network interface.

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
monitor = {}
```

<a id="canonical-9688fd703de709004fff61591cedefc3bb360b4e6386788187d55fd01ac18b83"></a>

## Direct properties — azure.not_managed.node_list.interface_list.monitor / 84b4cae226fb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6c171f79e78064b19f8dae94decb99fe0770f0eccb55dbe08e856bf5589f41ea"></a>

## Next pages — azure.not_managed.node_list.interface_list.monitor / 84b4cae226fb / 4

- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-e6fb49e7f03d55f01158b4f4a51bff46b21e4f0c2f7e8f3a13a124f0f9d29866"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-840e1981e3b3c09a3bae99f64a7eef13ff86382cfda50704dbea32121a8a94e0"></a>

## azure.not_managed.node_list.interface_list.monitor_disabled — azure.not_managed.node_list.interface_list.monitor_disabled / e536fa6329c8 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [azure](resources--securemesh_site_v2--reference--group-004.md#canonical-d6b794d7e7e25165fa55b5be6623d8b4c48cd27cb72b40d9005227c301983678)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0579a21e4252f6dae8e786094b3ef46ed712e25dc0e83d91b29344662b25ec19)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-3f3f4bbb5b7bd2586cd71c68e3c3d54a6cfa6d08a40994c2f866af5461db8d70)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- azure.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-ad6c5618f05ae304d3f765dedfe01fe22ee5f3969aef81e4d1c2817bc182a29d"></a>

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
monitor_disabled = {}
```

<a id="canonical-bd5b1ffebdcab7d022b99b16e57070fee754541764ffdce3e4d7abd455de7c12"></a>

## Direct properties — azure.not_managed.node_list.interface_list.monitor_disabled / e536fa6329c8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-61e29a8e33f80c9031ae2e0014b374b213ae69e464f5bbf23d12aa5157806d01"></a>

## Next pages — azure.not_managed.node_list.interface_list.monitor_disabled / e536fa6329c8 / 4

- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-236ed3d90d43799075c1819f3f1074ed513703db60954f2008ad86c4b5b97dc0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3bdf1b6fc25d20006cd85135e4813a283da9a2a40ec216fa884e110523f9b419"></a>

## azure.not_managed.node_list.interface_list.network_option — azure.not_managed.node_list.interface_list.network_option / f41753f35163 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [azure](resources--securemesh_site_v2--reference--group-004.md#canonical-d6b794d7e7e25165fa55b5be6623d8b4c48cd27cb72b40d9005227c301983678)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0579a21e4252f6dae8e786094b3ef46ed712e25dc0e83d91b29344662b25ec19)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-3f3f4bbb5b7bd2586cd71c68e3c3d54a6cfa6d08a40994c2f866af5461db8d70)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- azure.not_managed.node_list.interface_list.network_option

<a id="canonical-a96481d31036c1a4af791b3fe92529c9654e3b3b2745f864ab63a64473a51530"></a>

Type: `"object"`. single nested block, Optional.

Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are
local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs,
Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is
optional.

Upstream description:

Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are
local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs,
Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is
optional. Global VRFs are configured via Networking &gt; Segments. A site can have multiple Network
Segments (global VRFs).

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
  "x-ves-oneof-field-network_choice": "[\"segment_network\",\"site_local_inside_network\",\"site_local_network\"]"
}
```

Terraform syntax:

```terraform
network_option {
  # Configure direct properties listed below.
}
```

<a id="canonical-568dd4be10a5c077cb5762fb92586b0e63137b94544c3270296e97888a21c866"></a>

## Direct properties — azure.not_managed.node_list.interface_list.network_option / f41753f35163 / 3

- [site_local_inside_network](resources--securemesh_site_v2--reference--group-005.md#canonical-87097345af00fc5d0511e10dcb1746447f62c4bc6aca191214d7d7336a30435a): complete subsection reference.

- [site_local_network](resources--securemesh_site_v2--reference--group-005.md#canonical-1cdda4a3146347139f956df00d0ae1dd0dcd649a1385083ca97a7af48d973bc3): complete subsection reference.

<a id="canonical-37573086a251ed818d10e395588e57808b365b24aa10d09d0615931149ad22e1"></a>

## Next pages — azure.not_managed.node_list.interface_list.network_option / f41753f35163 / 4

- [azure.not_managed.node_list.interface_list.network_option.site_local_inside_network](resources--securemesh_site_v2--reference--group-005.md#canonical-87097345af00fc5d0511e10dcb1746447f62c4bc6aca191214d7d7336a30435a)
- [azure.not_managed.node_list.interface_list.network_option.site_local_network](resources--securemesh_site_v2--reference--group-005.md#canonical-1cdda4a3146347139f956df00d0ae1dd0dcd649a1385083ca97a7af48d973bc3)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-87097345af00fc5d0511e10dcb1746447f62c4bc6aca191214d7d7336a30435a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f8413c307a4f5241b9be6335bdffb1e1b5dfa1379465b57923ddf6fbd53ab00"></a>

## azure.not_managed.node_list.interface_list.network_option.site_local_inside_network — azure.not_managed.node_list.interface_list.network_option.site_local_inside_netw / 0b5c044cf25d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [azure](resources--securemesh_site_v2--reference--group-004.md#canonical-d6b794d7e7e25165fa55b5be6623d8b4c48cd27cb72b40d9005227c301983678)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0579a21e4252f6dae8e786094b3ef46ed712e25dc0e83d91b29344662b25ec19)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-3f3f4bbb5b7bd2586cd71c68e3c3d54a6cfa6d08a40994c2f866af5461db8d70)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- [azure.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-005.md#canonical-236ed3d90d43799075c1819f3f1074ed513703db60954f2008ad86c4b5b97dc0)
- azure.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-e93f777be914ed5772d413aed73814b0f90936d0e6c68483a3f8a0209830784a"></a>

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

<a id="canonical-dc8cb1e74702562e6dde2880860e145d8521deeed30686b28497cb8397ba00c4"></a>

## Direct properties — azure.not_managed.node_list.interface_list.network_option.site_local_inside_netw / 0b5c044cf25d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-272743394c1d44aa59d86e486353e20807d9f56fd3c376d8b544d5382fbfc06b"></a>

## Next pages — azure.not_managed.node_list.interface_list.network_option.site_local_inside_netw / 0b5c044cf25d / 4

- [azure.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-005.md#canonical-236ed3d90d43799075c1819f3f1074ed513703db60954f2008ad86c4b5b97dc0)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-1cdda4a3146347139f956df00d0ae1dd0dcd649a1385083ca97a7af48d973bc3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8aa78368f5170180c5d77e1d3fa1f6547e584d105ec62b51ad7995fff27106e0"></a>

## azure.not_managed.node_list.interface_list.network_option.site_local_network — azure.not_managed.node_list.interface_list.network_option.site_local_network / bed912094f8e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [azure](resources--securemesh_site_v2--reference--group-004.md#canonical-d6b794d7e7e25165fa55b5be6623d8b4c48cd27cb72b40d9005227c301983678)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0579a21e4252f6dae8e786094b3ef46ed712e25dc0e83d91b29344662b25ec19)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-3f3f4bbb5b7bd2586cd71c68e3c3d54a6cfa6d08a40994c2f866af5461db8d70)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- [azure.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-005.md#canonical-236ed3d90d43799075c1819f3f1074ed513703db60954f2008ad86c4b5b97dc0)
- azure.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-cb3f629a3b63361d0581065159cb720688f307787f4a5cca18f3bdc993353ddd"></a>

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

<a id="canonical-15d273acb567c9bdb45c5d184a32d98a1b3bb559564e129e26cc1901d2f924de"></a>

## Direct properties — azure.not_managed.node_list.interface_list.network_option.site_local_network / bed912094f8e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2308befff573223294269c40cbec5ba4c341e90974d138d87cfa373d75d47e3f"></a>

## Next pages — azure.not_managed.node_list.interface_list.network_option.site_local_network / bed912094f8e / 4

- [azure.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-005.md#canonical-236ed3d90d43799075c1819f3f1074ed513703db60954f2008ad86c4b5b97dc0)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-f102eeafcae07b7617e9161d5842a46176a671f002cd7b757cd176ab0bcffa6c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa68976ffb799626b1e28721355e98416960a279f90e9c0350a0d796857515e1"></a>

## azure.not_managed.node_list.interface_list.no_ipv4_address — azure.not_managed.node_list.interface_list.no_ipv4_address / e6c87262b56a / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [azure](resources--securemesh_site_v2--reference--group-004.md#canonical-d6b794d7e7e25165fa55b5be6623d8b4c48cd27cb72b40d9005227c301983678)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0579a21e4252f6dae8e786094b3ef46ed712e25dc0e83d91b29344662b25ec19)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-3f3f4bbb5b7bd2586cd71c68e3c3d54a6cfa6d08a40994c2f866af5461db8d70)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- azure.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-3bb5d1542c1287a3bfcb9c8d9a533c4f53e24679b27f7c43ab7ff5e9c05d7fbf"></a>

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
no_ipv4_address = {}
```

<a id="canonical-0ec3772cd38cf247ae2fc581eb3a57c49b3f75313a0d3306ed87e6263e112793"></a>

## Direct properties — azure.not_managed.node_list.interface_list.no_ipv4_address / e6c87262b56a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-864afb82bbbdad1dbcf47e8703f02ebeaebb78f2aac9de7282d3307fa52d60a9"></a>

## Next pages — azure.not_managed.node_list.interface_list.no_ipv4_address / e6c87262b56a / 4

- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-0697c3260784935c10bc90afaeee43ca02436863965cc80446df6a540ebb4961"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-591846b94900b375059531bda211b41cc55ed19c6ed0bb185985129efa71831a"></a>

## azure.not_managed.node_list.interface_list.no_ipv6_address — azure.not_managed.node_list.interface_list.no_ipv6_address / d204d5c951be / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [azure](resources--securemesh_site_v2--reference--group-004.md#canonical-d6b794d7e7e25165fa55b5be6623d8b4c48cd27cb72b40d9005227c301983678)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0579a21e4252f6dae8e786094b3ef46ed712e25dc0e83d91b29344662b25ec19)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-3f3f4bbb5b7bd2586cd71c68e3c3d54a6cfa6d08a40994c2f866af5461db8d70)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- azure.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-5abc7be1d58f8977a3eac1b10ee4488ee6bfe656065960b0a9a71472adebe075"></a>

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
no_ipv6_address = {}
```

<a id="canonical-4e47b64c78573a122923392cf496394f6a0d3406bd05f8b490388050a1326fb3"></a>

## Direct properties — azure.not_managed.node_list.interface_list.no_ipv6_address / d204d5c951be / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9339eebab6ac5dd9cc275f489875ce78d8d054e7c16a11e66d53e99290644eaf"></a>

## Next pages — azure.not_managed.node_list.interface_list.no_ipv6_address / d204d5c951be / 4

- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-5749eedbb0fba7cf4eb7dbdbe02c6109110703b42ebe31eac9d0b024d70ca044"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a6a8db106686758e2dc15f90c226bc1a75973fdfcb44aa9d870e84d30e30aa0"></a>

## azure.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled — azure.not_managed.node_list.interface_list.site_to_site_connectivity_interface_d / 6d6ffd8054aa / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [azure](resources--securemesh_site_v2--reference--group-004.md#canonical-d6b794d7e7e25165fa55b5be6623d8b4c48cd27cb72b40d9005227c301983678)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0579a21e4252f6dae8e786094b3ef46ed712e25dc0e83d91b29344662b25ec19)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-3f3f4bbb5b7bd2586cd71c68e3c3d54a6cfa6d08a40994c2f866af5461db8d70)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- azure.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-50f24b888aa8d836fdc994dc48d6ebd5510ff3069bd8bf41e035ccba83d1c94c"></a>

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
site_to_site_connectivity_interface_disabled = {}
```

<a id="canonical-bfcef83553a7a685b7763df5f791db22aa446bb6b568199d4faeea67a466235c"></a>

## Direct properties — azure.not_managed.node_list.interface_list.site_to_site_connectivity_interface_d / 6d6ffd8054aa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2df4492fc948f047d3a4b64e58439ca4db3f5b70d0a83ae408316108cc3c0ef1"></a>

## Next pages — azure.not_managed.node_list.interface_list.site_to_site_connectivity_interface_d / 6d6ffd8054aa / 4

- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-75fcb7e41eb4a2c391289fba5feeb23273161c2f72b1ee2bee6b949e93eee3f5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-378e1b02dcfa40553cf57b5e7211cffbdfb091808480f5a305d93fa5f7d8761e"></a>

## azure.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled — azure.not_managed.node_list.interface_list.site_to_site_connectivity_interface_e / cb9dd65ae607 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [azure](resources--securemesh_site_v2--reference--group-004.md#canonical-d6b794d7e7e25165fa55b5be6623d8b4c48cd27cb72b40d9005227c301983678)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0579a21e4252f6dae8e786094b3ef46ed712e25dc0e83d91b29344662b25ec19)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-3f3f4bbb5b7bd2586cd71c68e3c3d54a6cfa6d08a40994c2f866af5461db8d70)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- azure.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-699f82957036466d7b4a86063fefaf2c02534ff77aaeed96f8836f704c664909"></a>

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
site_to_site_connectivity_interface_enabled = {}
```

<a id="canonical-056efe80a4310bc62febd91de904d2a0e03bcb9d33f5d50ca462bcf4e5fed42c"></a>

## Direct properties — azure.not_managed.node_list.interface_list.site_to_site_connectivity_interface_e / cb9dd65ae607 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-85555c1eb4676e282cf09cf0c5fa7d105b50bede13ab089e3f358a4a088c2966"></a>

## Next pages — azure.not_managed.node_list.interface_list.site_to_site_connectivity_interface_e / cb9dd65ae607 / 4

- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-94d35cae0c12416bd7beac4ef8d7cf7ccfee2a0b86d7c67624e891d7df626399"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3d72a5d645f5fe5aa64d54fabbd46c24d18f577d514b035f6c2685fead94714d"></a>

## azure.not_managed.node_list.interface_list.static_ip — azure.not_managed.node_list.interface_list.static_ip / f993480ca73b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [azure](resources--securemesh_site_v2--reference--group-004.md#canonical-d6b794d7e7e25165fa55b5be6623d8b4c48cd27cb72b40d9005227c301983678)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0579a21e4252f6dae8e786094b3ef46ed712e25dc0e83d91b29344662b25ec19)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-3f3f4bbb5b7bd2586cd71c68e3c3d54a6cfa6d08a40994c2f866af5461db8d70)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- azure.not_managed.node_list.interface_list.static_ip

<a id="canonical-668cf2602623a185a83478bef9d3e506266e00a6af0938c479b756a60989213c"></a>

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
static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-2bd3f03ba83de575f0d8e84f8b24b4110d1f29abb29352ea7d66d1d9630f316f"></a>

## Direct properties — azure.not_managed.node_list.interface_list.static_ip / f993480ca73b / 3

<a id="canonical-dbeb8ca1739a65d3f9c946ce90b52395efbce7ff1fbf165a586e2faee63fa9b6"></a>

<a id="canonical-e5bb89b0a60be036a39851faeb24b8053116ed3d8c35b32a902b7068e2bbc6bc"></a>

## default_gw property — azure.not_managed.node_list.interface_list.static_ip / f993480ca73b / 4

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

<a id="canonical-7aa72de41c02287708b1412ade3c1b52dfb27dbcbd9ab2fcf1bf4c3764edcca2"></a>

<a id="canonical-8316c879784e2cb8996c44f5017f3a647a1bc1548022bfb476bc121cf07d5a5b"></a>

## dns_server property — azure.not_managed.node_list.interface_list.static_ip / f993480ca73b / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-cd9071a82045770d3816e5b8bf84cad54933ef73e8cc5da91b0b2ee2ed041cde"></a>

<a id="canonical-6150c546bc2c4ca9eafc96fac4db42702bb6d4fa7d3e719966101407e8676099"></a>

## ip_address property — azure.not_managed.node_list.interface_list.static_ip / f993480ca73b / 6

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

<a id="canonical-35e347d3ad87a4a17b6ebfccaa8c7baea05d11f7a69c0227e6a090e9fc6af09f"></a>

## Next pages — azure.not_managed.node_list.interface_list.static_ip / f993480ca73b / 7

- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-a5e034c1d51afdbee265b5b3e1598312b5384c5f7f2f54bdf17037f79b0b665b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9a0ef4b317b57d17e66c4e1573924befc9706449a52fabd6216627647615e476"></a>

## azure.not_managed.node_list.interface_list.static_ipv6_address — azure.not_managed.node_list.interface_list.static_ipv6_address / 6664d945138b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [azure](resources--securemesh_site_v2--reference--group-004.md#canonical-d6b794d7e7e25165fa55b5be6623d8b4c48cd27cb72b40d9005227c301983678)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0579a21e4252f6dae8e786094b3ef46ed712e25dc0e83d91b29344662b25ec19)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-3f3f4bbb5b7bd2586cd71c68e3c3d54a6cfa6d08a40994c2f866af5461db8d70)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- azure.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-cef6b102dfdbb8625e858e618d5c94ae8d3cc95d29fb9a98f21d37a56cb1b6a4"></a>

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

<a id="canonical-b42bf39372c21532f02ed7ba1fde59fa18ce321b7b32017ee21f4c3534b668bf"></a>

## Direct properties — azure.not_managed.node_list.interface_list.static_ipv6_address / 6664d945138b / 3

- [cluster_static_ip](resources--securemesh_site_v2--reference--group-005.md#canonical-8e74e5f6a6da50c1a14e0be283ecafef40bb1644d61ecc11c72fa079c9b92cbe): complete subsection reference.

- [node_static_ip](resources--securemesh_site_v2--reference--group-005.md#canonical-c36d962adad6567f7442f6f1aac4459c2a55a5c035a6c183d9f0c5134f4b7f5d): complete subsection reference.

<a id="canonical-1a7987588c46833e8c97d7d555946bf00d102f20b4ce600ad5b8ce4e0782a6fe"></a>

## Next pages — azure.not_managed.node_list.interface_list.static_ipv6_address / 6664d945138b / 4

- [azure.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip](resources--securemesh_site_v2--reference--group-005.md#canonical-8e74e5f6a6da50c1a14e0be283ecafef40bb1644d61ecc11c72fa079c9b92cbe)
- [azure.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip](resources--securemesh_site_v2--reference--group-005.md#canonical-c36d962adad6567f7442f6f1aac4459c2a55a5c035a6c183d9f0c5134f4b7f5d)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-8e74e5f6a6da50c1a14e0be283ecafef40bb1644d61ecc11c72fa079c9b92cbe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9fd4db48b6744016f5ea69a6aa4a87e9365707bfd0d33731825d40fa6a49323b"></a>

## azure.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip — azure.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip / 101c8a3760be / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [azure](resources--securemesh_site_v2--reference--group-004.md#canonical-d6b794d7e7e25165fa55b5be6623d8b4c48cd27cb72b40d9005227c301983678)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0579a21e4252f6dae8e786094b3ef46ed712e25dc0e83d91b29344662b25ec19)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-3f3f4bbb5b7bd2586cd71c68e3c3d54a6cfa6d08a40994c2f866af5461db8d70)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- [azure.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-005.md#canonical-a5e034c1d51afdbee265b5b3e1598312b5384c5f7f2f54bdf17037f79b0b665b)
- azure.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-11f1ea3e965e13d70dfd04f214a52bd4a90155332b925e513f1925a8c3101d5b"></a>

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

<a id="canonical-ddd7c15fb7ee7b4213a65a95e9e27f8c88872eaac790dc1b7061d6acf341a77b"></a>

## Direct properties — azure.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip / 101c8a3760be / 3

<a id="canonical-87f5171cc64136212970bbc46c440812e51afe4f3476c31b78970b9b7428376e"></a>

<a id="canonical-90a8d8e177ff6a4021ac4d2ca3a5f29b252c43843198ff1ecafaf5bf0da84e5d"></a>

## interface_ip_map property — azure.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip / 101c8a3760be / 4

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

<a id="canonical-c704d10bb846a853cb77924fa70c3ad1148f8752251b80fc54766eab2e94945f"></a>

## Next pages — azure.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip / 101c8a3760be / 5

- [azure.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-005.md#canonical-a5e034c1d51afdbee265b5b3e1598312b5384c5f7f2f54bdf17037f79b0b665b)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-c36d962adad6567f7442f6f1aac4459c2a55a5c035a6c183d9f0c5134f4b7f5d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-953d00c74ee9b6248540d794a43a511d3bd19c03c6d67de06937b0b1ad1e3915"></a>

## azure.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip — azure.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 6ccd2d51dde0 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [azure](resources--securemesh_site_v2--reference--group-004.md#canonical-d6b794d7e7e25165fa55b5be6623d8b4c48cd27cb72b40d9005227c301983678)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0579a21e4252f6dae8e786094b3ef46ed712e25dc0e83d91b29344662b25ec19)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-3f3f4bbb5b7bd2586cd71c68e3c3d54a6cfa6d08a40994c2f866af5461db8d70)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- [azure.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-005.md#canonical-a5e034c1d51afdbee265b5b3e1598312b5384c5f7f2f54bdf17037f79b0b665b)
- azure.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-a3b4544a03a2194ff0157423e8f0b88c7b8f555366b0b741337240726b39bc28"></a>

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

<a id="canonical-96dd5e3f781371cea1088363a5c342e3648c6ffaee3a45ed9060c1dfe1ea5a88"></a>

## Direct properties — azure.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 6ccd2d51dde0 / 3

<a id="canonical-ee15e712dc91a6d799b7103188dc61ca243ebd27b1e6cef5c4ede0ac3c64a767"></a>

<a id="canonical-e856879d1c01d51817fb1e20fd0c9a2004108b5b5169eb1fe7866ac635ce0488"></a>

## default_gw property — azure.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 6ccd2d51dde0 / 4

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

<a id="canonical-672c77b60239af001d9a030d7c68e8ca67c8dabd7f3a2fbba5af1b5744853d48"></a>

<a id="canonical-dc176c2be91f7c6383e1351d8a9e3c477b98946042f076083bc5ab7a61f25ca5"></a>

## dns_server property — azure.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 6ccd2d51dde0 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-4331db20a733e2d2ae362243afa1257e03bf7cd87552f8a5dae095354c38e842"></a>

<a id="canonical-4531919f9d4cbdc2a3847aae5f8ae960c0ed971f9a9b2f549e8ef2025692fcfb"></a>

## ip_address property — azure.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 6ccd2d51dde0 / 6

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

<a id="canonical-9d8d3b6c94d161e04f0773706d855a89266e1f8fc0b14a94edec0ccd8223c2c6"></a>

## Next pages — azure.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 6ccd2d51dde0 / 7

- [azure.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-005.md#canonical-a5e034c1d51afdbee265b5b3e1598312b5384c5f7f2f54bdf17037f79b0b665b)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-3d9a98d0961fce7147ce874b304bdf7f98fb285ac81e4936157e5242dca6775b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ede16f6d712d0f6f10e0a9f1d5365c981656bceae7a8b5974a212761f4a9384"></a>

## azure.not_managed.node_list.interface_list.vlan_interface — azure.not_managed.node_list.interface_list.vlan_interface / 9520a0d36055 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [azure](resources--securemesh_site_v2--reference--group-004.md#canonical-d6b794d7e7e25165fa55b5be6623d8b4c48cd27cb72b40d9005227c301983678)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0579a21e4252f6dae8e786094b3ef46ed712e25dc0e83d91b29344662b25ec19)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-3f3f4bbb5b7bd2586cd71c68e3c3d54a6cfa6d08a40994c2f866af5461db8d70)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- azure.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-31b7fdc5d0b762ca20fd3d91ba08406c218bc1fd546b2fee325cf230c624862b"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for vlan interface.

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
vlan_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-aa77f3af922d80ffedc44893e60611a106bb353ada22a31a9f3b2b9301721138"></a>

## Direct properties — azure.not_managed.node_list.interface_list.vlan_interface / 9520a0d36055 / 3

<a id="canonical-8369102c18b55042da7e71e86f77f101f96e59d3eea4fd898d72137e1c162460"></a>

<a id="canonical-cbebf9a7212997f67d9c61971ddb8534e061d2b0b5d7a480a993e0f370fc70c9"></a>

## device property — azure.not_managed.node_list.interface_list.vlan_interface / 9520a0d36055 / 4

Type: `"string"`. Optional.

Select a parent interface from the dropdown.

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

<a id="canonical-08cc9aa08b9c17e8ee2d86fca411d4299395a7466a9a421197ce56d40a024829"></a>

<a id="canonical-9e33510b4561f14b16c96d61e8307c81ba9085c5a955c67bc69b358eb2c80737"></a>

## vlan_id property — azure.not_managed.node_list.interface_list.vlan_interface / 9520a0d36055 / 5

Type: `"number"`. Optional.

Configure the VLAN tag for this interface.

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```

<a id="canonical-5648c376d435f010f98e9bc9e694f1f3e68ff7702045b9ef5c4e0f176e40c9d4"></a>

## Next pages — azure.not_managed.node_list.interface_list.vlan_interface / 9520a0d36055 / 6

- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-80914e17925d76ddd91332aa9870b55841fb37e132ded1afe2ffe95b5e337d5e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-a15df73f71a4c0d9ee81c0cd83baade5f542ec9981ed827a060270cd4c587973"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-87c1677028a38c2f517599c716a69b4c0b47440ae9504a4d9c448d64db7d729f"></a>

## baremetal — baremetal / 0f07f9451b77 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- baremetal

<a id="canonical-ea32864777ea097f63e0153dbf69a39354f5dc07260c59b5c8fa96ba030f3cc9"></a>

Type: `"object"`. single nested block, Optional.

Baremetal Provider Type. Baremetal Provider Type.

Upstream description:

Baremetal Provider Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-orchestration_choice": "[\"not_managed\"]"
}
```

Terraform syntax:

```terraform
baremetal {
  # Configure direct properties listed below.
}
```

<a id="canonical-d9be6351f9d16b0fd74ee6ad5d89c8ec342a0b251748dca693eba1f55da69b7b"></a>

## Direct properties — baremetal / 0f07f9451b77 / 3

- [not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-fc7f76385435a915d0f870d69949f59a405049b13fa3b76417d5a0325bb03d4e): complete subsection reference.

<a id="canonical-08515eb4e86d8320deb43e33649cb8f82cb67b003a431ddfc6bd8153beb5936a"></a>

## Next pages — baremetal / 0f07f9451b77 / 4

- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-fc7f76385435a915d0f870d69949f59a405049b13fa3b76417d5a0325bb03d4e)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-fc7f76385435a915d0f870d69949f59a405049b13fa3b76417d5a0325bb03d4e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e9a0421f1ef8003d978bfb9b6fa5f603e26682db6e445e8b2288f1a5cf5c7409"></a>

## baremetal.not_managed — baremetal.not_managed / 600843493257 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-a15df73f71a4c0d9ee81c0cd83baade5f542ec9981ed827a060270cd4c587973)
- baremetal.not_managed

<a id="canonical-7b3c54b2518ff4f0f7c67c9ee0185276a9bd5b0323226f7507d4aadbb61ba7ae"></a>

Type: `"object"`. single nested block, Optional.

Section will show nodes associated with this site.

Upstream description:

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

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
not_managed {
  # Configure direct properties listed below.
}
```

<a id="canonical-62585bbde0551c8e1d2ec078de1db671a46529ad6f3e257381760092645d7619"></a>

## Direct properties — baremetal.not_managed / 600843493257 / 3

- [node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-c8d11ab5e7018fcb305dfbf1c30f2dc552163272ae517fbfb8251fdb18ead345): complete subsection reference.

<a id="canonical-ec494940d0664cfb2b87323c7fd5014fdb26311e05f99ae5e6848c7e20c74011"></a>

## Next pages — baremetal.not_managed / 600843493257 / 4

- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-c8d11ab5e7018fcb305dfbf1c30f2dc552163272ae517fbfb8251fdb18ead345)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-a15df73f71a4c0d9ee81c0cd83baade5f542ec9981ed827a060270cd4c587973)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-c8d11ab5e7018fcb305dfbf1c30f2dc552163272ae517fbfb8251fdb18ead345"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cfd78acccca9228f53add712b61073dca16d5cca6e0591848b1ef8411ec8de6e"></a>

## baremetal.not_managed.node_list — baremetal.not_managed.node_list / 5cfca7f9f873 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-a15df73f71a4c0d9ee81c0cd83baade5f542ec9981ed827a060270cd4c587973)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-fc7f76385435a915d0f870d69949f59a405049b13fa3b76417d5a0325bb03d4e)
- baremetal.not_managed.node_list

<a id="canonical-7b5e21134322e3033590e4ccf8f80f966b273e283b11e28835f32c3cd1095529"></a>

Type: `"object"`. list nested block, Optional.

Section will show nodes associated with this site.

Upstream description:

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
node_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-19024e6338e316b42839606899046cb6b6a69aa980b8cf881e9c2cf07aa38e0b"></a>

## Direct properties — baremetal.not_managed.node_list / 5cfca7f9f873 / 3

<a id="canonical-bf99ec2fd9b990ca238bbcc866d6b237ceea7c17d98f380c9489536801b49f10"></a>

<a id="canonical-8c0f01975ab0cd8944f9100a7b4ec1fe60e9efaf1876ba73b185551044f8aaa5"></a>

## hostname property — baremetal.not_managed.node_list / 5cfca7f9f873 / 4

Type: `"string"`. Optional.

Hostname. Hostname for this Node.

Upstream description:

Hostname for this Node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
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
    "format": "fqdn",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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

- [interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-6b6a8d6a5fc50a9f5a10b15b452c5bc52586309b69b52f8abec65cb20341f8c1): complete subsection reference.

<a id="canonical-a4ffe141488ac3a61822c3c5227df3fc413d1c7d9a58ab1ce10268aee1292413"></a>

<a id="canonical-00cdbb38888e5ad75b0a0a332c710f450dac933dc3899aa70addd04326b6c2de"></a>

## public_ip property — baremetal.not_managed.node_list / 5cfca7f9f873 / 5

Type: `"string"`. Optional.

Public IP. Public IP for this Node.

Upstream description:

Public IP for this Node.

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

<a id="canonical-2a47413a15bde1a2ac74f79623c6b438cc79dec9feac6657ddd56cb8c344bdec"></a>

<a id="canonical-26ab1054c69c98404cc6a66fa7d3fae578b46b55160966776a29cb145754a3d1"></a>

## type property — baremetal.not_managed.node_list / 5cfca7f9f873 / 6

Type: `"string"`. Optional.

\[Enum: Control|Worker\] Type for this Node, can be Control or Worker. Possible values are
\`Control\`, \`Worker\`.

Upstream description:

Type for this Node, can be Control or Worker.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("Control",
    "Worker"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "Control",
    "Worker"
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
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  }
}
```

<a id="canonical-fed242c044443c76e638a296f5f47af1aff1376d1d479dcac37824bdb277cb78"></a>

## Next pages — baremetal.not_managed.node_list / 5cfca7f9f873 / 7

- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-6b6a8d6a5fc50a9f5a10b15b452c5bc52586309b69b52f8abec65cb20341f8c1)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-fc7f76385435a915d0f870d69949f59a405049b13fa3b76417d5a0325bb03d4e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-6b6a8d6a5fc50a9f5a10b15b452c5bc52586309b69b52f8abec65cb20341f8c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae564d8abc4d2df86c530936f977d6618373c8a88773a317ce18540cf871e519"></a>

## baremetal.not_managed.node_list.interface_list — baremetal.not_managed.node_list.interface_list / 1d895621f472 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-a15df73f71a4c0d9ee81c0cd83baade5f542ec9981ed827a060270cd4c587973)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-fc7f76385435a915d0f870d69949f59a405049b13fa3b76417d5a0325bb03d4e)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-c8d11ab5e7018fcb305dfbf1c30f2dc552163272ae517fbfb8251fdb18ead345)
- baremetal.not_managed.node_list.interface_list

<a id="canonical-d64acff67676e9e7c958bd925e4239d9fe062a735850c1d72ff1a553617a614f"></a>

Type: `"object"`. list nested block, Optional.

Manage interfaces belonging to this node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("bond_interface",
    "ethernet_interface"),
  validators.ConflictingListObjectAttributes("bond_interface",
    "vlan_interface"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "dhcp_server"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "no_ipv4_address"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "static_ip"),
  validators.ConflictingListObjectAttributes("dhcp_server",
    "no_ipv4_address"),
  validators.ConflictingListObjectAttributes("dhcp_server",
    "static_ip"),
  validators.ConflictingListObjectAttributes("ethernet_interface",
    "vlan_interface"),
  validators.ConflictingListObjectAttributes("ipv6_auto_config",
    "no_ipv6_address"),
  validators.ConflictingListObjectAttributes("ipv6_auto_config",
    "static_ipv6_address"),
  validators.ConflictingListObjectAttributes("monitor",
    "monitor_disabled"),
  validators.ConflictingListObjectAttributes("no_ipv4_address",
    "static_ip"),
  validators.ConflictingListObjectAttributes("no_ipv6_address",
    "static_ipv6_address"),
  validators.ConflictingListObjectAttributes("site_to_site_connectivity_interface_disabled",
    "site_to_site_connectivity_interface_enabled")}
```

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
interface_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-544791190b35ab1f52e010334df73df8656154861e7c7a05454781d218e2d73c"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list / 1d895621f472 / 3

- [bond_interface](resources--securemesh_site_v2--reference--group-005.md#canonical-f93f35af6eb3bcd20545b533e190d87c6f79d1db343003521366bc32004287eb): complete subsection reference.

<a id="canonical-7637357359c74ebecf34a37f07da00c80de1a34885e4f9ee1f688af1a7e0c531"></a>

<a id="canonical-359b85f05a409e8effdbb3fe6fba0975e65f82316e74908196ac3c995a990c34"></a>

## description_spec property — baremetal.not_managed.node_list.interface_list / 1d895621f472 / 4

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [dhcp_client](resources--securemesh_site_v2--reference--group-005.md#canonical-01601f5ec093a1c5b74c2f3298be3b8ec23d9fc4d3a3d8dfe222a2cb131e2c84): complete subsection reference.

- [dhcp_server](resources--securemesh_site_v2--reference--group-005.md#canonical-ebce07839e1c7296072446a21a569ccc40fc34c8d68ac4ef5915c4fb77ffc428): complete subsection reference.

- [ethernet_interface](resources--securemesh_site_v2--reference--group-006.md#canonical-b5577c7d0343a5e9d90715db56e88f7f2e8dc99abe21e99dc87910b9af99205b): complete subsection reference.

- [ipv6_auto_config](resources--securemesh_site_v2--reference--group-006.md#canonical-cdf949d413591037a576db25a13031db92865e13d838ce360551be22209138f1): complete subsection reference.

<a id="canonical-5a97ef0aeab89b25ffd1f1554f2044260f4f800c0198ac1610be2640032f7637"></a>

<a id="canonical-87febdc1e8adeedd5a189388d48af34124582c706a5632e0703b3b482daf3391"></a>

## is_management property — baremetal.not_managed.node_list.interface_list / 1d895621f472 / 5

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-c88f5d0940cd4d0fb3f7b7510aa3e5ac0e4669ab3f080993faab350b258952de"></a>

<a id="canonical-c776b697d880eca1f8b5ff3911d905adae255794aecde1dd0f13e383a3d06223"></a>

## is_primary property — baremetal.not_managed.node_list.interface_list / 1d895621f472 / 6

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-ba65643e7e02e908d76b638272da176aaff94436d5cfa60a3eb6094399f2fd5a"></a>

<a id="canonical-0ed9a476911537bc717a7b712855e58f0bfef51d597971379a2b882615d9c204"></a>

## labels property — baremetal.not_managed.node_list.interface_list / 1d895621f472 / 7

Type: `["map", "string"]`. Optional.

Add Labels for this Interface, these labels can be used in firewall policy.

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
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [monitor](resources--securemesh_site_v2--reference--group-006.md#canonical-41c64e7bfa0ac0de51b7bddbfd115d147e2634f2514a589e869b409e519ccc4b): complete subsection reference.

- [monitor_disabled](resources--securemesh_site_v2--reference--group-006.md#canonical-65319d075c62465398393491dba0f5d9942f20d0720ffb29ae9f76f9ef672d3b): complete subsection reference.

<a id="canonical-6565c3cbab3e3608150bbf7cb608bbc8470f3a199a82597fc5decbf808527f8f"></a>

<a id="canonical-bfaf385f864bf3befac897930518e9f37d19f24de9072add338b459c9f108813"></a>

## mtu property — baremetal.not_managed.node_list.interface_list / 1d895621f472 / 8

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 512, Maximum: 8000},
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
    "maximum": 8000,
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
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  }
}
```

<a id="canonical-5f1fc538e7add2dc63c0057662b21025c9afa5c752c01bad0f1d2a4f71c24049"></a>

<a id="canonical-b8c83c5b79cfe8f74fe83d2c44c2747758b6c4fe0402c92ccdc94e125f7f1ebf"></a>

## name property — baremetal.not_managed.node_list.interface_list / 1d895621f472 / 9

Type: `"string"`. Optional.

Interface Name. Name of this Interface.

Upstream description:

Name of this Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

- [network_option](resources--securemesh_site_v2--reference--group-006.md#canonical-b3404c572f5eeae7f2cfefdd8c2f78562396d6780008ed33a8bc63ca63116c0e): complete subsection reference.

- [no_ipv4_address](resources--securemesh_site_v2--reference--group-006.md#canonical-a1f7076106cbba2625c8ec35664d15b0f51c7a80f10956479b1e780f6674bc83): complete subsection reference.

- [no_ipv6_address](resources--securemesh_site_v2--reference--group-006.md#canonical-075f8d1c361c021601e74fef63ab8800890f567c46d98ae16cb41ff3a6e5106f): complete subsection reference.

<a id="canonical-e217db022a53656eac686d7b56036c5270f3e6775ef6f98f3ded775386d1d431"></a>

<a id="canonical-2df854aa598a03a4c72bc3667ec781051857417284ffcf18407bd0627c0f37b0"></a>

## priority property — baremetal.not_managed.node_list.interface_list / 1d895621f472 / 10

Type: `"number"`. Optional.

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Upstream description:

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

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

- [site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-006.md#canonical-8c598b78ef45a993038fc72731d55848985125adefd9820115f7303eebb10ecc): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-006.md#canonical-6e855368eb1f7a75c3c76fbe797069ebe4f4c53eb30e8cf78e726031075c7edd): complete subsection reference.

- [static_ip](resources--securemesh_site_v2--reference--group-006.md#canonical-9e54ebe6a2696da75846a7b743ce03b84aaa0c0467162d322c2bed2cafc54e2d): complete subsection reference.

- [static_ipv6_address](resources--securemesh_site_v2--reference--group-006.md#canonical-87eb00fe70ef1bfbe01fcde85d942da80923d07fb3112e48947769423112ba87): complete subsection reference.

- [vlan_interface](resources--securemesh_site_v2--reference--group-006.md#canonical-5eb617386088fb763324add0b3b808b39034b52361638825f0d9d70a36bcee93): complete subsection reference.

<a id="canonical-7b6b0a9ef9cbb3122c371092ec04993c38ee9ca43400056870492324f424cf00"></a>

## Next pages — baremetal.not_managed.node_list.interface_list / 1d895621f472 / 11

- [baremetal.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-005.md#canonical-f93f35af6eb3bcd20545b533e190d87c6f79d1db343003521366bc32004287eb)
- [baremetal.not_managed.node_list.interface_list.dhcp_client](resources--securemesh_site_v2--reference--group-005.md#canonical-01601f5ec093a1c5b74c2f3298be3b8ec23d9fc4d3a3d8dfe222a2cb131e2c84)
- [baremetal.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-005.md#canonical-ebce07839e1c7296072446a21a569ccc40fc34c8d68ac4ef5915c4fb77ffc428)
- [baremetal.not_managed.node_list.interface_list.ethernet_interface](resources--securemesh_site_v2--reference--group-006.md#canonical-b5577c7d0343a5e9d90715db56e88f7f2e8dc99abe21e99dc87910b9af99205b)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-006.md#canonical-cdf949d413591037a576db25a13031db92865e13d838ce360551be22209138f1)
- [baremetal.not_managed.node_list.interface_list.monitor](resources--securemesh_site_v2--reference--group-006.md#canonical-41c64e7bfa0ac0de51b7bddbfd115d147e2634f2514a589e869b409e519ccc4b)
- [baremetal.not_managed.node_list.interface_list.monitor_disabled](resources--securemesh_site_v2--reference--group-006.md#canonical-65319d075c62465398393491dba0f5d9942f20d0720ffb29ae9f76f9ef672d3b)
- [baremetal.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-006.md#canonical-b3404c572f5eeae7f2cfefdd8c2f78562396d6780008ed33a8bc63ca63116c0e)
- [baremetal.not_managed.node_list.interface_list.no_ipv4_address](resources--securemesh_site_v2--reference--group-006.md#canonical-a1f7076106cbba2625c8ec35664d15b0f51c7a80f10956479b1e780f6674bc83)
- [baremetal.not_managed.node_list.interface_list.no_ipv6_address](resources--securemesh_site_v2--reference--group-006.md#canonical-075f8d1c361c021601e74fef63ab8800890f567c46d98ae16cb41ff3a6e5106f)
- [baremetal.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-006.md#canonical-8c598b78ef45a993038fc72731d55848985125adefd9820115f7303eebb10ecc)
- [baremetal.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-006.md#canonical-6e855368eb1f7a75c3c76fbe797069ebe4f4c53eb30e8cf78e726031075c7edd)
- [baremetal.not_managed.node_list.interface_list.static_ip](resources--securemesh_site_v2--reference--group-006.md#canonical-9e54ebe6a2696da75846a7b743ce03b84aaa0c0467162d322c2bed2cafc54e2d)
- [baremetal.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-006.md#canonical-87eb00fe70ef1bfbe01fcde85d942da80923d07fb3112e48947769423112ba87)
- [baremetal.not_managed.node_list.interface_list.vlan_interface](resources--securemesh_site_v2--reference--group-006.md#canonical-5eb617386088fb763324add0b3b808b39034b52361638825f0d9d70a36bcee93)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-c8d11ab5e7018fcb305dfbf1c30f2dc552163272ae517fbfb8251fdb18ead345)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-f93f35af6eb3bcd20545b533e190d87c6f79d1db343003521366bc32004287eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-042bfc1d180b412b911ea295f3b9113922d194fe68d12281879e41c1d1128325"></a>

## baremetal.not_managed.node_list.interface_list.bond_interface — baremetal.not_managed.node_list.interface_list.bond_interface / 5f17796bb9b3 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-a15df73f71a4c0d9ee81c0cd83baade5f542ec9981ed827a060270cd4c587973)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-fc7f76385435a915d0f870d69949f59a405049b13fa3b76417d5a0325bb03d4e)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-c8d11ab5e7018fcb305dfbf1c30f2dc552163272ae517fbfb8251fdb18ead345)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-6b6a8d6a5fc50a9f5a10b15b452c5bc52586309b69b52f8abec65cb20341f8c1)
- baremetal.not_managed.node_list.interface_list.bond_interface

<a id="canonical-12e24f293a2409a32b91e417a5b52e567918317a96b4ff91ba48406fb4da2cbe"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bond interface.

Upstream description:

Bond devices configuration for fleet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("devices",
    "link_polling_interval",
    "link_up_delay",
    "name"),
  validators.ConflictingObjectAttributes("active_backup",
    "lacp")}
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
  "x-ves-oneof-field-lacp_choice": "[\"active_backup\",\"lacp\"]"
}
```

Terraform syntax:

```terraform
bond_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-99da5d1082a39f6bc77d0dfb2b31f45d33cf20f2e9c3d25947c80c4039e667f5"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.bond_interface / 5f17796bb9b3 / 3

- [active_backup](resources--securemesh_site_v2--reference--group-005.md#canonical-fe154f009216e201824b2a5d4bcb1d7af8828ca1c6139c6a6e1523db3175922f): complete subsection reference.

<a id="canonical-a49761fbb232ad84351fa8eb73a1239d8e2dffc0abe66874e18a819032d7b8a0"></a>

<a id="canonical-d4c5024dcaf904908f9cf55edefd30b5eaad2b684820af3fa54d12ee59153137"></a>

## devices property — baremetal.not_managed.node_list.interface_list.bond_interface / 5f17796bb9b3 / 4

Type: `["list", "string"]`. Optional.

Ethernet devices that will make up this bond.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 8),
}
```

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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [lacp](resources--securemesh_site_v2--reference--group-005.md#canonical-55ae5899b3c3831cc2a346cf8a39944cbb540de5ed3d928926e7868b986010e1): complete subsection reference.

<a id="canonical-bcb25ef2eb65cc5e484fd55104c1a3372b3cbad069dcd47a8f06c7ad7e159dbd"></a>

<a id="canonical-78deb8abd2a72b1c3eb56d116bade879ea29e883a0218cc90488569c96d76616"></a>

## link_polling_interval property — baremetal.not_managed.node_list.interface_list.bond_interface / 5f17796bb9b3 / 5

Type: `"number"`. Optional.

Link Polling Interval. Link polling interval in milliseconds.

Upstream description:

Link polling interval in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(500, 5000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 500
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  }
}
```

<a id="canonical-402da5282e3c9b42a5151bbc2bb2c1f70ae1a39103398cd1e10e33956be6f158"></a>

<a id="canonical-fc08341d00089aa2a3c8c1323ef5253862b36309fd21906da060e251038cb818"></a>

## link_up_delay property — baremetal.not_managed.node_list.interface_list.bond_interface / 5f17796bb9b3 / 6

Type: `"number"`. Optional.

Milliseconds wait before link is declared up.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 1000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1000"
  }
}
```

<a id="canonical-e53de28f23b1bfa2756ad592554a752ee1b83bfa17fefb5a943ef0427b9eaec9"></a>

<a id="canonical-460fbc1bfa9e78a2513e75025604dff3661681017def04649770dde8d9105a1a"></a>

## name property — baremetal.not_managed.node_list.interface_list.bond_interface / 5f17796bb9b3 / 7

Type: `"string"`. Optional.

Bond Device Name. Name for the Bond. Ex 'bond0'

Upstream description:

Name for the Bond. Ex 'bond0'

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
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-f482167c17191f1d229753e279766170ea6a565d41a8c48b4607da9db71cc72c"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.bond_interface / 5f17796bb9b3 / 8

- [baremetal.not_managed.node_list.interface_list.bond_interface.active_backup](resources--securemesh_site_v2--reference--group-005.md#canonical-fe154f009216e201824b2a5d4bcb1d7af8828ca1c6139c6a6e1523db3175922f)
- [baremetal.not_managed.node_list.interface_list.bond_interface.lacp](resources--securemesh_site_v2--reference--group-005.md#canonical-55ae5899b3c3831cc2a346cf8a39944cbb540de5ed3d928926e7868b986010e1)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-6b6a8d6a5fc50a9f5a10b15b452c5bc52586309b69b52f8abec65cb20341f8c1)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-fe154f009216e201824b2a5d4bcb1d7af8828ca1c6139c6a6e1523db3175922f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c8be752b44c5f80de93ebd8f7b09e922493130c8fbd9364b8f3a3e9dd550916"></a>

## baremetal.not_managed.node_list.interface_list.bond_interface.active_backup — baremetal.not_managed.node_list.interface_list.bond_interface.active_backup / 134530e0bb4e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-a15df73f71a4c0d9ee81c0cd83baade5f542ec9981ed827a060270cd4c587973)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-fc7f76385435a915d0f870d69949f59a405049b13fa3b76417d5a0325bb03d4e)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-c8d11ab5e7018fcb305dfbf1c30f2dc552163272ae517fbfb8251fdb18ead345)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-6b6a8d6a5fc50a9f5a10b15b452c5bc52586309b69b52f8abec65cb20341f8c1)
- [baremetal.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-005.md#canonical-f93f35af6eb3bcd20545b533e190d87c6f79d1db343003521366bc32004287eb)
- baremetal.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-6704dfa0e747acc01b64aa1907b103f115a3c8627a93556a8f08f234ee7de14f"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for active backup.

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
active_backup = {}
```

<a id="canonical-b6b4f20334ff16fa7c566846993b2d5623ac78b7d72f800fe546201390f13f68"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.bond_interface.active_backup / 134530e0bb4e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8e0734f4d1c67e98e4ab4f1fef0a50b90df52c932787b1c7daba1335089a66a2"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.bond_interface.active_backup / 134530e0bb4e / 4

- [baremetal.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-005.md#canonical-f93f35af6eb3bcd20545b533e190d87c6f79d1db343003521366bc32004287eb)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-55ae5899b3c3831cc2a346cf8a39944cbb540de5ed3d928926e7868b986010e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c83296637cd1882266f0758864c1ce14b3f3736b8bebc3f60720430f89d3aa1b"></a>

## baremetal.not_managed.node_list.interface_list.bond_interface.lacp — baremetal.not_managed.node_list.interface_list.bond_interface.lacp / 91f191215d27 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-a15df73f71a4c0d9ee81c0cd83baade5f542ec9981ed827a060270cd4c587973)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-fc7f76385435a915d0f870d69949f59a405049b13fa3b76417d5a0325bb03d4e)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-c8d11ab5e7018fcb305dfbf1c30f2dc552163272ae517fbfb8251fdb18ead345)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-6b6a8d6a5fc50a9f5a10b15b452c5bc52586309b69b52f8abec65cb20341f8c1)
- [baremetal.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-005.md#canonical-f93f35af6eb3bcd20545b533e190d87c6f79d1db343003521366bc32004287eb)
- baremetal.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-a59a2d463039cc775f6b066d0a781f1d22d47aa1262d5b696a47bcf58fa98005"></a>

Type: `"object"`. single nested block, Optional.

LACP parameters. LACP parameters for the bond device.

Upstream description:

LACP parameters for the bond device.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rate")}
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
lacp {
  # Configure direct properties listed below.
}
```

<a id="canonical-ffcd481ddf83e2b5193a26bb6a19047da1ef595b0353c27c6205e4bbb71531b8"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.bond_interface.lacp / 91f191215d27 / 3

<a id="canonical-b504d7d072b48dd3ec12cd9ff7fa826eca8f97166efbfcebb5a3709a3ed76d15"></a>

<a id="canonical-efc81e65974f5e0abeb01463aab644122535891ddf6bda7b4c50e7515ee9ddc8"></a>

## rate property — baremetal.not_managed.node_list.interface_list.bond_interface.lacp / 91f191215d27 / 4

Type: `"number"`. Optional.

Interval in seconds to transmit LACP packets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 30),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```

<a id="canonical-d2a9c8416c152e6da202338beefbbb35a31305e828f4bcabce8961fd0db7d572"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.bond_interface.lacp / 91f191215d27 / 5

- [baremetal.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-005.md#canonical-f93f35af6eb3bcd20545b533e190d87c6f79d1db343003521366bc32004287eb)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-01601f5ec093a1c5b74c2f3298be3b8ec23d9fc4d3a3d8dfe222a2cb131e2c84"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41e42e9f5deba545674a3b7d8c2ee36ebde5e7134549d999df18fa50b3ac6910"></a>

## baremetal.not_managed.node_list.interface_list.dhcp_client — baremetal.not_managed.node_list.interface_list.dhcp_client / 6d2ff54fac5a / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-a15df73f71a4c0d9ee81c0cd83baade5f542ec9981ed827a060270cd4c587973)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-fc7f76385435a915d0f870d69949f59a405049b13fa3b76417d5a0325bb03d4e)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-c8d11ab5e7018fcb305dfbf1c30f2dc552163272ae517fbfb8251fdb18ead345)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-6b6a8d6a5fc50a9f5a10b15b452c5bc52586309b69b52f8abec65cb20341f8c1)
- baremetal.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-3e65a8c00f83f4606cf3ed0cc7e61b2edf043cb3eed1ead285f06fccdfe590b8"></a>

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
dhcp_client = {}
```

<a id="canonical-bf321bca1efcf3ec5e0ba0dac27bd43a54c0e19bbc2b76d6bcd35b99d84ee6bb"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.dhcp_client / 6d2ff54fac5a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6857a6546098dfe596f2e3c4ae81b2702ec59d6f5afb777136f818db6e027457"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.dhcp_client / 6d2ff54fac5a / 4

- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-6b6a8d6a5fc50a9f5a10b15b452c5bc52586309b69b52f8abec65cb20341f8c1)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-ebce07839e1c7296072446a21a569ccc40fc34c8d68ac4ef5915c4fb77ffc428"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-685a904c6ccd10099c86025fedc3412a6df9f18467e5a66e050f58e58fc61d7c"></a>

## baremetal.not_managed.node_list.interface_list.dhcp_server — baremetal.not_managed.node_list.interface_list.dhcp_server / a510a38e2cac / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-a15df73f71a4c0d9ee81c0cd83baade5f542ec9981ed827a060270cd4c587973)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-fc7f76385435a915d0f870d69949f59a405049b13fa3b76417d5a0325bb03d4e)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-c8d11ab5e7018fcb305dfbf1c30f2dc552163272ae517fbfb8251fdb18ead345)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-6b6a8d6a5fc50a9f5a10b15b452c5bc52586309b69b52f8abec65cb20341f8c1)
- baremetal.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-50a39dbd82c038d59dedfd87bf1e997f099fd29d4f2de4ecb2e90723a7f6cc1e"></a>

Type: `"object"`. single nested block, Optional.

DHCPServerParametersType.

Upstream description:

DHCP server configuration for this interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dhcp_networks"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "automatic_from_start"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "interface_ip_map"),
  validators.ConflictingObjectAttributes("automatic_from_start",
    "interface_ip_map")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

Terraform syntax:

```terraform
dhcp_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-0331b5dc3fe33ff7c00a51f725a0285bff5fb8738904902dc0ae9971c82627f4"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.dhcp_server / a510a38e2cac / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-005.md#canonical-72758156cf174b1eacdd2bf6f8f729618749992f54784dc200101ed00eac64bb): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-005.md#canonical-a499c8ddadcd6914da5b2e2804041c808f1abff09e3956001320cb65f52f298f): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-005.md#canonical-bad7a4771e30e995c659f27e5996f725af675ecad86daf9c73b174893a34c200): complete subsection reference.

<a id="canonical-65dafa6bc6d941fd2a7966db991b149186f5fcf2c18edbd363c6857d7d183731"></a>

<a id="canonical-c2e2b3c3e46a845b72cbfaef451161ee19cf194658477d2dc4da3386a6c0a208"></a>

## dhcp_option82_tag property — baremetal.not_managed.node_list.interface_list.dhcp_server / a510a38e2cac / 4

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="canonical-b40eb455eb727f3ebf5322f5c23c45777b5fa903a49a24e0c1de1fbc180162e5"></a>

<a id="canonical-a4afbfe9252c11ff3ee08a07c821e97beef1f2f6aa3d57b23974144319204d13"></a>

## fixed_ip_map property — baremetal.not_managed.node_list.interface_list.dhcp_server / a510a38e2cac / 5

Type: `["map", "string"]`. Optional.

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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-006.md#canonical-46fdfecb66b470d2672f78a3709159161d23c3b6eea985eb1f28602bdd3ebeab): complete subsection reference.

<a id="canonical-97f50beb2a9c0ebe2a0ed1dc8e180cae86e5683ea0d41a0df8a14041fe22eeb7"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.dhcp_server / a510a38e2cac / 6

- [baremetal.not_managed.node_list.interface_list.dhcp_server.automatic_from_end](resources--securemesh_site_v2--reference--group-005.md#canonical-72758156cf174b1eacdd2bf6f8f729618749992f54784dc200101ed00eac64bb)
- [baremetal.not_managed.node_list.interface_list.dhcp_server.automatic_from_start](resources--securemesh_site_v2--reference--group-005.md#canonical-a499c8ddadcd6914da5b2e2804041c808f1abff09e3956001320cb65f52f298f)
- [baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-005.md#canonical-bad7a4771e30e995c659f27e5996f725af675ecad86daf9c73b174893a34c200)
- [baremetal.not_managed.node_list.interface_list.dhcp_server.interface_ip_map](resources--securemesh_site_v2--reference--group-006.md#canonical-46fdfecb66b470d2672f78a3709159161d23c3b6eea985eb1f28602bdd3ebeab)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-6b6a8d6a5fc50a9f5a10b15b452c5bc52586309b69b52f8abec65cb20341f8c1)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-72758156cf174b1eacdd2bf6f8f729618749992f54784dc200101ed00eac64bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a5093ba2bc099945e3863fb28e831077230f9c4b3c06a41bedc9142409818c0"></a>

## baremetal.not_managed.node_list.interface_list.dhcp_server.automatic_from_end — baremetal.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / f6c7c4381209 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-a15df73f71a4c0d9ee81c0cd83baade5f542ec9981ed827a060270cd4c587973)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-fc7f76385435a915d0f870d69949f59a405049b13fa3b76417d5a0325bb03d4e)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-c8d11ab5e7018fcb305dfbf1c30f2dc552163272ae517fbfb8251fdb18ead345)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-6b6a8d6a5fc50a9f5a10b15b452c5bc52586309b69b52f8abec65cb20341f8c1)
- [baremetal.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-005.md#canonical-ebce07839e1c7296072446a21a569ccc40fc34c8d68ac4ef5915c4fb77ffc428)
- baremetal.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-d73fcf692dd4f0e5af2c18a1b1ae77fc9c21ee8b1a6d96587aefdc3f97ea40af"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
automatic_from_end = {}
```

<a id="canonical-66877b0321375b55277287b23156e6988a7c60e8ae712e10c592663e02c16364"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / f6c7c4381209 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d1318851f4b434ebafa9f44e5ca9e6aa5ab09817905757b322aa854e492f559d"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / f6c7c4381209 / 4

- [baremetal.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-005.md#canonical-ebce07839e1c7296072446a21a569ccc40fc34c8d68ac4ef5915c4fb77ffc428)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-a499c8ddadcd6914da5b2e2804041c808f1abff09e3956001320cb65f52f298f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6eaf0c407502f83e68664c25a829ec93cccd7cab45f09932da844e456542af91"></a>

## baremetal.not_managed.node_list.interface_list.dhcp_server.automatic_from_start — baremetal.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / e368d4856ba9 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-a15df73f71a4c0d9ee81c0cd83baade5f542ec9981ed827a060270cd4c587973)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-fc7f76385435a915d0f870d69949f59a405049b13fa3b76417d5a0325bb03d4e)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-c8d11ab5e7018fcb305dfbf1c30f2dc552163272ae517fbfb8251fdb18ead345)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-6b6a8d6a5fc50a9f5a10b15b452c5bc52586309b69b52f8abec65cb20341f8c1)
- [baremetal.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-005.md#canonical-ebce07839e1c7296072446a21a569ccc40fc34c8d68ac4ef5915c4fb77ffc428)
- baremetal.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-e3bcdabdadde29cc6f84cab4a33b478e090d93934c3a795a636c3015adfcd787"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
automatic_from_start = {}
```

<a id="canonical-be2a23086ee3e1815a4700eee35a8e760638051ec7141dc8dbfe0040a115e9b4"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / e368d4856ba9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bfa56d25fc51f73f540c0d096904e1b7f791a0f259628b67a48a6b0f6c17de72"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / e368d4856ba9 / 4

- [baremetal.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-005.md#canonical-ebce07839e1c7296072446a21a569ccc40fc34c8d68ac4ef5915c4fb77ffc428)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-bad7a4771e30e995c659f27e5996f725af675ecad86daf9c73b174893a34c200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-847326c8a9061423498a4eb6ca2f5740760570f75116352aec8197dd2e952e38"></a>

## baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks — baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 9aca42f2ec49 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-a15df73f71a4c0d9ee81c0cd83baade5f542ec9981ed827a060270cd4c587973)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-fc7f76385435a915d0f870d69949f59a405049b13fa3b76417d5a0325bb03d4e)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-c8d11ab5e7018fcb305dfbf1c30f2dc552163272ae517fbfb8251fdb18ead345)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-6b6a8d6a5fc50a9f5a10b15b452c5bc52586309b69b52f8abec65cb20341f8c1)
- [baremetal.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-005.md#canonical-ebce07839e1c7296072446a21a569ccc40fc34c8d68ac4ef5915c4fb77ffc428)
- baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-280901656d07b6cae4f4c2c3223819afdf0bd5d42c16aabc837ae172f669737f"></a>

Type: `"object"`. list nested block, Optional.

List of networks from which DHCP Server can allocate IPv4 Addresses.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("dgw_address",
    "first_address"),
  validators.ConflictingListObjectAttributes("dgw_address",
    "last_address"),
  validators.ConflictingListObjectAttributes("dns_address",
    "same_as_dgw"),
  validators.ConflictingListObjectAttributes("first_address",
    "last_address")}
```

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

Terraform syntax:

```terraform
dhcp_networks {
  # Configure direct properties listed below.
}
```

<a id="canonical-eaaf9129114d6996e36101c50e0eb813208528e9d520e122c8c889ad601e8aa4"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 9aca42f2ec49 / 3

<a id="canonical-4b7bc65b2ef585efa1b85f58eaf4ef8f27d8b23d76f459b9526bd0194c13ffd7"></a>

<a id="canonical-f37f7e3d4d01deda4e224871aa036b44db162cad46e2a6324fdf442af1fe04d2"></a>

## dgw_address property — baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 9aca42f2ec49 / 4

Type: `"string"`. Optional.

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Upstream description:

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

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

<a id="canonical-7b1d3864a3eb965f76e4cfd8974424aa1f159084a2354be3db8d9b26db1c775d"></a>

<a id="canonical-0ba44c1afeef9c58131dd89698f5086068da54e434bcb314385f320b5065ee02"></a>

## dns_address property — baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 9aca42f2ec49 / 5

Type: `"string"`. Optional.

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Upstream description:

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

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

- [first_address](resources--securemesh_site_v2--reference--group-005.md#canonical-2200ede4093330777611d6566c959564de2bc0cef6bd83f0d5b9087dc3efc99d): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-006.md#canonical-13a2246dd866be48c0260b181d10979303852c712ebde46a134398bc47611735): complete subsection reference.

<a id="canonical-7a1cd2ee7c9393e52bf52f169296808daff0182aa3fcc99f5e03e269cbac0cdf"></a>

<a id="canonical-2a2ed28233d8d926e8dd31a9d7182296413c020fb5fbdf914b3f1505ded9ef0f"></a>

## network_prefix property — baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 9aca42f2ec49 / 6

Type: `"string"`. Optional.

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

<a id="canonical-b022de57b264f9a852d27571f1cd5d7e4ea7722e088edb71f600c4544a12ff95"></a>

<a id="canonical-0506ce8e2b6ffec28d94a08310edca785aa587550a6096483d496d0f209bd2a6"></a>

## pool_settings property — baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 9aca42f2ec49 / 7

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"),
}
```

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

- [pools](resources--securemesh_site_v2--reference--group-006.md#canonical-9c5cffa5822f9f9c18f84d04d8f500113cd2db2896dde876f4e573c90b7acb26): complete subsection reference.

- [same_as_dgw](resources--securemesh_site_v2--reference--group-006.md#canonical-06e907cdb5c158788077adc9f3d34055e09b125acccc7b60c6626b7b0350e60a): complete subsection reference.

<a id="canonical-1f57065338321a2feb6bf9b4c667d7075be70a1e22705dab7560fa228447dfa0"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 9aca42f2ec49 / 8

- [baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address](resources--securemesh_site_v2--reference--group-005.md#canonical-2200ede4093330777611d6566c959564de2bc0cef6bd83f0d5b9087dc3efc99d)
- [baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address](resources--securemesh_site_v2--reference--group-006.md#canonical-13a2246dd866be48c0260b181d10979303852c712ebde46a134398bc47611735)
- [baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools](resources--securemesh_site_v2--reference--group-006.md#canonical-9c5cffa5822f9f9c18f84d04d8f500113cd2db2896dde876f4e573c90b7acb26)
- [baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw](resources--securemesh_site_v2--reference--group-006.md#canonical-06e907cdb5c158788077adc9f3d34055e09b125acccc7b60c6626b7b0350e60a)
- [baremetal.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-005.md#canonical-ebce07839e1c7296072446a21a569ccc40fc34c8d68ac4ef5915c4fb77ffc428)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-2200ede4093330777611d6566c959564de2bc0cef6bd83f0d5b9087dc3efc99d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-128740eaafa37277020c54b6777316922a5db02a9cde665316b0498c06af4c26"></a>

## baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address — baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_a / f141df7165d8 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-a15df73f71a4c0d9ee81c0cd83baade5f542ec9981ed827a060270cd4c587973)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-fc7f76385435a915d0f870d69949f59a405049b13fa3b76417d5a0325bb03d4e)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-c8d11ab5e7018fcb305dfbf1c30f2dc552163272ae517fbfb8251fdb18ead345)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-6b6a8d6a5fc50a9f5a10b15b452c5bc52586309b69b52f8abec65cb20341f8c1)
- [baremetal.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-005.md#canonical-ebce07839e1c7296072446a21a569ccc40fc34c8d68ac4ef5915c4fb77ffc428)
- [baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-005.md#canonical-bad7a4771e30e995c659f27e5996f725af675ecad86daf9c73b174893a34c200)
- baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-bae5f81d072b5410c39d6b1b6b8f2a249a69ecfddba3fa6110205b293e4fe72d"></a>

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
first_address = {}
```

<a id="canonical-39729e8df26aab1985cdaee54fa2a2ec170a7e84f11a29765fb6d571f85e222d"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_a / f141df7165d8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0e1357c1d5259ec7d830752575dd84e2d56a2acc68e42511af80903d9569cb21"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_a / f141df7165d8 / 4

- [baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-005.md#canonical-bad7a4771e30e995c659f27e5996f725af675ecad86daf9c73b174893a34c200)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
