---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-4166ee2d047636770fb58c562d2e17d3ea81d07a02c9d7b736b88cfeb54a0416"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / 284523cbc1d3 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-017.md#canonical-1ca26003927c6657202334c8a528e4b985397ff13caad46c5efcc976b7f7333f)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-017.md#canonical-c5c7b3fe8db397038c37d95ab016e54688593303819651f1927e249b240ebe08)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-eaba80cd4719640b28135094dbf18c11c90262f4a0ad44c88b46660b428667da"></a>

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

<a id="canonical-172424836803f151ebcb585767363f701ef32a3f898e4e0a253573980628dc11"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / 284523cbc1d3 / 3

- [configured_list](resources--securemesh_site_v2--reference--group-018.md#canonical-ffd2ced7394b3e59c8dbd682223e2a9647f78dc76a1725378519c697e03ee56b): complete subsection reference.

- [local_dns](resources--securemesh_site_v2--reference--group-018.md#canonical-ed9c677711a4773cf4a7656dad29b46105cf18e7092ba73981908cfc8aceca34): complete subsection reference.

<a id="canonical-3ff76618603d1a8d437ea56c7d9d3eb692f338d52ddef58db604136a62897556"></a>

## Next pages — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / 284523cbc1d3 / 4

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list](resources--securemesh_site_v2--reference--group-018.md#canonical-ffd2ced7394b3e59c8dbd682223e2a9647f78dc76a1725378519c697e03ee56b)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-018.md#canonical-ed9c677711a4773cf4a7656dad29b46105cf18e7092ba73981908cfc8aceca34)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-017.md#canonical-c5c7b3fe8db397038c37d95ab016e54688593303819651f1927e249b240ebe08)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-ffd2ced7394b3e59c8dbd682223e2a9647f78dc76a1725378519c697e03ee56b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3bd7eef947e98286565255ff7f021763761b6efd99a7c67b9d51097e94b5086a"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.c / fc2293848b2f / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-017.md#canonical-1ca26003927c6657202334c8a528e4b985397ff13caad46c5efcc976b7f7333f)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-017.md#canonical-c5c7b3fe8db397038c37d95ab016e54688593303819651f1927e249b240ebe08)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-017.md#canonical-3bc5c5da15314133ff03418fbec7fb825fb63d7b5ba1b72a8c28fceb10bb9739)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-ef97b487e0e7c58da478ba64c5f580000bd63f1a923cbc28fecb72c897a540e5"></a>

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

<a id="canonical-e760d9f2351a4e271a4e13c13875083cb24bfb872c90177be396ab1a4345c3aa"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.c / fc2293848b2f / 3

<a id="canonical-6ebe5a846f4ed9a6ad06959b7bb9887011f5b5eb181bb8097681711f0d93470f"></a>

<a id="canonical-d6086f5acd1f54acbbadd63ccd536ac6234763a688868519345914aa266f509a"></a>

## dns_list property — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.c / fc2293848b2f / 4

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

<a id="canonical-9c9260bedba707cb4aea44337751c28b49b8121e0085bb077d0a6b6a5fec5600"></a>

## Next pages — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.c / fc2293848b2f / 5

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-017.md#canonical-3bc5c5da15314133ff03418fbec7fb825fb63d7b5ba1b72a8c28fceb10bb9739)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-ed9c677711a4773cf4a7656dad29b46105cf18e7092ba73981908cfc8aceca34"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c6443d9d64800b5b4caf6582c35f55f589a3d9e0dd433273ca505b85a74dd99"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.l / 8555eb874504 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-017.md#canonical-1ca26003927c6657202334c8a528e4b985397ff13caad46c5efcc976b7f7333f)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-017.md#canonical-c5c7b3fe8db397038c37d95ab016e54688593303819651f1927e249b240ebe08)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-017.md#canonical-3bc5c5da15314133ff03418fbec7fb825fb63d7b5ba1b72a8c28fceb10bb9739)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-be93b3115de11dbf56a19177fd0875305810e3f4c2d39a1362ef51f3d941e927"></a>

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

<a id="canonical-f1d21e958ab2d98c0b3dc204772acb3fd108a4a4bb84e8e58028550758a009e2"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.l / 8555eb874504 / 3

<a id="canonical-b5b57e58067cd9bca1bc7e8b4f448298b254e79ad79f3264eb94c8becfacbe93"></a>

<a id="canonical-8551523a4192239b71bbfca79a734bcd15227ae66504c65abd151d854b93cba1"></a>

## configured_address property — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.l / 8555eb874504 / 4

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

- [first_address](resources--securemesh_site_v2--reference--group-018.md#canonical-bd703ca45c39fb8d263e3e8eb8321df7a3c605e95c16eeaefa46dfdc38e1a7c5): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-018.md#canonical-42c9cc7055fb7d2c073504d28f368e3c32300cd8dacb268e188fac119c59079d): complete subsection reference.

<a id="canonical-dd67c696f591b49671e6ced174421edf6a840b46077de553a5ca951a772fad77"></a>

## Next pages — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.l / 8555eb874504 / 5

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address](resources--securemesh_site_v2--reference--group-018.md#canonical-bd703ca45c39fb8d263e3e8eb8321df7a3c605e95c16eeaefa46dfdc38e1a7c5)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address](resources--securemesh_site_v2--reference--group-018.md#canonical-42c9cc7055fb7d2c073504d28f368e3c32300cd8dacb268e188fac119c59079d)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-017.md#canonical-3bc5c5da15314133ff03418fbec7fb825fb63d7b5ba1b72a8c28fceb10bb9739)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-bd703ca45c39fb8d263e3e8eb8321df7a3c605e95c16eeaefa46dfdc38e1a7c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1bc51cc20218f70a4dac0c9b23919b48eadeef4fd088995664e7504eb44046f2"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.l / 41c0337d6dea / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-017.md#canonical-1ca26003927c6657202334c8a528e4b985397ff13caad46c5efcc976b7f7333f)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-017.md#canonical-c5c7b3fe8db397038c37d95ab016e54688593303819651f1927e249b240ebe08)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-017.md#canonical-3bc5c5da15314133ff03418fbec7fb825fb63d7b5ba1b72a8c28fceb10bb9739)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-018.md#canonical-ed9c677711a4773cf4a7656dad29b46105cf18e7092ba73981908cfc8aceca34)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-67558273553174023a0fd352df656d02ffb39f89438b5580c744d6db13f076df"></a>

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

<a id="canonical-eeb2d5f69922b2eb674a37ab905a8ce53548c473a8f8f0cdcf7419a49d236af5"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.l / 41c0337d6dea / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b598d46296b46ef9a474a9c31c9fcca9380482298f5f4ec671ae3f92c778e43d"></a>

## Next pages — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.l / 41c0337d6dea / 4

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-018.md#canonical-ed9c677711a4773cf4a7656dad29b46105cf18e7092ba73981908cfc8aceca34)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-42c9cc7055fb7d2c073504d28f368e3c32300cd8dacb268e188fac119c59079d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1fe1217126651399ced67d460490d3be0587987bc1ee8738c0bdd2eb1d63ec98"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.l / 4bc6a4de5227 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-017.md#canonical-1ca26003927c6657202334c8a528e4b985397ff13caad46c5efcc976b7f7333f)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-017.md#canonical-c5c7b3fe8db397038c37d95ab016e54688593303819651f1927e249b240ebe08)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-017.md#canonical-3bc5c5da15314133ff03418fbec7fb825fb63d7b5ba1b72a8c28fceb10bb9739)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-018.md#canonical-ed9c677711a4773cf4a7656dad29b46105cf18e7092ba73981908cfc8aceca34)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-5f08473d669c84a9d5e6d867d97c484f54db08c1f4370ce1288acc0d7e7fc650"></a>

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

<a id="canonical-31ccc07b30100ce001bead5b0f3c817f958731633df2f6a99385bf02e7a37cb2"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.l / 4bc6a4de5227 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cc6587fc3c96f5eeec21016ea6a8fe9f3536511651b082a5177c04fcc4346f74"></a>

## Next pages — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.l / 4bc6a4de5227 / 4

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-018.md#canonical-ed9c677711a4773cf4a7656dad29b46105cf18e7092ba73981908cfc8aceca34)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-7c7d839a74b1dbce9e5eeb5e2072c6578f923d2e53c5ca1cbb463fce794b6d6e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95d96df33a94883718090182efaaacca17902f72abd123ba90e0b5eb2a37fa28"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / d840feba3da5 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-017.md#canonical-1ca26003927c6657202334c8a528e4b985397ff13caad46c5efcc976b7f7333f)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-017.md#canonical-c5c7b3fe8db397038c37d95ab016e54688593303819651f1927e249b240ebe08)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-bba01aa4b8e0758e723a0c2e2416722cd6e8f06e5bc76a2b49b74297a72699c5"></a>

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

<a id="canonical-f56e1df2c94ef6c95f73f492fac598b48a3598e4d41408143fc10ba7635a05d2"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / d840feba3da5 / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-018.md#canonical-0c2f2e7ec780751c1301ce67ee1d633dc06a7db40b3e683a7cf6638b03db956a): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-018.md#canonical-bb94f038c96a76aafbf8e732ef87d7e18efec6bf03a1b9f2aa58514634f3acc8): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-018.md#canonical-7e1bc02f3b605df9e846a09616952d66b1a434e1df937feccc864ee6df8b9392): complete subsection reference.

<a id="canonical-91485dab549470e1f82ae9fdb9e995017b9193fd40625fac3535e49ba790fc00"></a>

<a id="canonical-3ba6970dc6c303797495ba26c8443fb0baae9348684aa93df11349d71c35eba6"></a>

## fixed_ip_map property — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / d840feba3da5 / 4

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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-018.md#canonical-9e09269cba9b5aff93f4fa4f4f75853877e9483bb3166b9f1462493e0d51a152): complete subsection reference.

<a id="canonical-cf590537c2529988f47d915b4161048ac2c342b74ad430cf853ef51798de059d"></a>

## Next pages — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / d840feba3da5 / 5

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end](resources--securemesh_site_v2--reference--group-018.md#canonical-0c2f2e7ec780751c1301ce67ee1d633dc06a7db40b3e683a7cf6638b03db956a)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start](resources--securemesh_site_v2--reference--group-018.md#canonical-bb94f038c96a76aafbf8e732ef87d7e18efec6bf03a1b9f2aa58514634f3acc8)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-018.md#canonical-7e1bc02f3b605df9e846a09616952d66b1a434e1df937feccc864ee6df8b9392)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map](resources--securemesh_site_v2--reference--group-018.md#canonical-9e09269cba9b5aff93f4fa4f4f75853877e9483bb3166b9f1462493e0d51a152)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-017.md#canonical-c5c7b3fe8db397038c37d95ab016e54688593303819651f1927e249b240ebe08)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-0c2f2e7ec780751c1301ce67ee1d633dc06a7db40b3e683a7cf6638b03db956a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9cadb6e81e8bfcd9f6c8526710bcb63e8fc81047685db74761d489df364896a2"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.aut / 7d4ce2ea9340 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-017.md#canonical-1ca26003927c6657202334c8a528e4b985397ff13caad46c5efcc976b7f7333f)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-017.md#canonical-c5c7b3fe8db397038c37d95ab016e54688593303819651f1927e249b240ebe08)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-018.md#canonical-7c7d839a74b1dbce9e5eeb5e2072c6578f923d2e53c5ca1cbb463fce794b6d6e)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-577db1a0f887b2cddba6b815d294870a087523232337c9a28b097b3644ba7bb2"></a>

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

<a id="canonical-0e60e906e2b0d98effb19897b2a64315ee56cda502cba186a5501abfa34e69cc"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.aut / 7d4ce2ea9340 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b975c4e64bccfa7c86c6eb56ae5e2625e23ea341cbb188c8fa85f072b3201f67"></a>

## Next pages — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.aut / 7d4ce2ea9340 / 4

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-018.md#canonical-7c7d839a74b1dbce9e5eeb5e2072c6578f923d2e53c5ca1cbb463fce794b6d6e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-bb94f038c96a76aafbf8e732ef87d7e18efec6bf03a1b9f2aa58514634f3acc8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b8a02f732e94a193d65f7c27a46a6e0762bac82fe6eb5610a859e2b6313bacc"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.aut / c52aec8e5fc7 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-017.md#canonical-1ca26003927c6657202334c8a528e4b985397ff13caad46c5efcc976b7f7333f)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-017.md#canonical-c5c7b3fe8db397038c37d95ab016e54688593303819651f1927e249b240ebe08)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-018.md#canonical-7c7d839a74b1dbce9e5eeb5e2072c6578f923d2e53c5ca1cbb463fce794b6d6e)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-18c00472e27c25ce5e142589b67abdd0b250f930a373d003f73c2944b19be7b2"></a>

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

<a id="canonical-d393f1f8a7bbc4a12e1af8fce5d3f4edeb5e34d8a79056a7828834797b25c3a3"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.aut / c52aec8e5fc7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-900247928ccd19877bf556732d225bdf1944647787edcea808718f212f24f612"></a>

## Next pages — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.aut / c52aec8e5fc7 / 4

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-018.md#canonical-7c7d839a74b1dbce9e5eeb5e2072c6578f923d2e53c5ca1cbb463fce794b6d6e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-7e1bc02f3b605df9e846a09616952d66b1a434e1df937feccc864ee6df8b9392"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-362bb84af560b4370c35d077f0d857d48096c8429e8f9402de9e6718ec7ae2f0"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhc / da721c029d86 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-017.md#canonical-1ca26003927c6657202334c8a528e4b985397ff13caad46c5efcc976b7f7333f)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-017.md#canonical-c5c7b3fe8db397038c37d95ab016e54688593303819651f1927e249b240ebe08)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-018.md#canonical-7c7d839a74b1dbce9e5eeb5e2072c6578f923d2e53c5ca1cbb463fce794b6d6e)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-1754f5f385471abaf97ee8c23c8033edd3388010fbf107c52b1d062ffae26e3f"></a>

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

<a id="canonical-8064ffb9bcce18ad0a1343dc6d49d1e04b8053ccd1e393229df5a46fd60ed3ef"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhc / da721c029d86 / 3

<a id="canonical-ca9960af6fa1e51c767bf6f3aa018e46f1e3f9b370ee1a930451853a44969ba8"></a>

<a id="canonical-02150bc79d3ee09292a7bf2515beac927076fcdea117a5c7d210aba2a91fe8f1"></a>

## network_prefix property — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhc / da721c029d86 / 4

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

<a id="canonical-517689430439692d31e99f914c8c5fc6c6c3ff9df4a7a2cda7e0f91f1e007859"></a>

<a id="canonical-e57836d56515b54d6d77ede2b323485c5690e0b76ea4ae874e9c0d7e8fe9ba7c"></a>

## pool_settings property — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhc / da721c029d86 / 5

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

- [pools](resources--securemesh_site_v2--reference--group-018.md#canonical-334d1651f6f85979144f7408f746304f48a54a503e9745fece534e4ff2264c1a): complete subsection reference.

<a id="canonical-823fa149cb439cc2e588e5e1a89dae526e16971d288bf1df5773639874801116"></a>

## Next pages — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhc / da721c029d86 / 6

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools](resources--securemesh_site_v2--reference--group-018.md#canonical-334d1651f6f85979144f7408f746304f48a54a503e9745fece534e4ff2264c1a)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-018.md#canonical-7c7d839a74b1dbce9e5eeb5e2072c6578f923d2e53c5ca1cbb463fce794b6d6e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-334d1651f6f85979144f7408f746304f48a54a503e9745fece534e4ff2264c1a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc618717b13935a12f3d33ed206d5159ee9f993244791d4eba6c381586f84074"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhc / 637a25da55ed / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-017.md#canonical-1ca26003927c6657202334c8a528e4b985397ff13caad46c5efcc976b7f7333f)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-017.md#canonical-c5c7b3fe8db397038c37d95ab016e54688593303819651f1927e249b240ebe08)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-018.md#canonical-7c7d839a74b1dbce9e5eeb5e2072c6578f923d2e53c5ca1cbb463fce794b6d6e)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-018.md#canonical-7e1bc02f3b605df9e846a09616952d66b1a434e1df937feccc864ee6df8b9392)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-6049e9ac0072ef92e5c7b4a4248fcf3640e2b68ab1922f84321347531ea81334"></a>

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

<a id="canonical-af143749b1ccec0062b1306e7da362814fdb26af37b67ac04f3e3cd386d3b390"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhc / 637a25da55ed / 3

<a id="canonical-99ba3da97de46cded0b31aeba03a0ce50b225183c107495fdd963d8cc5d72b5f"></a>

<a id="canonical-a2abe566540e996a3b15fb47e38beb17aa26b52d2c264fc88bf388f1cf9e3690"></a>

## end_ip property — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhc / 637a25da55ed / 4

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

<a id="canonical-352b5b568409898d304d303ae98c290157daa87dfbd93b560fca7b51cd0dd848"></a>

<a id="canonical-ad2b7669163c7d87ca31329774b28d84dc8046a531e170703573c5f0d00178f9"></a>

## start_ip property — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhc / 637a25da55ed / 5

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

<a id="canonical-f074a92d772429158148b3e3ba785cb579992c89f4198196e366402271eaff1f"></a>

## Next pages — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhc / 637a25da55ed / 6

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-018.md#canonical-7e1bc02f3b605df9e846a09616952d66b1a434e1df937feccc864ee6df8b9392)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-9e09269cba9b5aff93f4fa4f4f75853877e9483bb3166b9f1462493e0d51a152"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4d83a68f5a3f1bc007b92e9d42fba3d62e8cb0b7f3d47f8bc056b5fd0557890a"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.int / c452d7b42764 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-017.md#canonical-1ca26003927c6657202334c8a528e4b985397ff13caad46c5efcc976b7f7333f)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-017.md#canonical-c5c7b3fe8db397038c37d95ab016e54688593303819651f1927e249b240ebe08)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-018.md#canonical-7c7d839a74b1dbce9e5eeb5e2072c6578f923d2e53c5ca1cbb463fce794b6d6e)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-d5a6b1d09b60e69637a0e6b0ed6d2fafdc73ea47af20a24638282976c42d7503"></a>

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

<a id="canonical-3f67becd7d1e9dc271d5227cdee36539e495c2b135d7b27343d05a47bc8c33af"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.int / c452d7b42764 / 3

<a id="canonical-97742a575eb58bd9ebd04446565db5511985945e7f9045ee85c6af36b5f38ccf"></a>

<a id="canonical-455efbd116f6f5f5be5f7fe6a01a356a780004e65cbf2647a85ba2fc85075d9a"></a>

## interface_ip_map property — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.int / c452d7b42764 / 4

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

<a id="canonical-7aa69241c2b50fc3de60998d73ab141bd0e2efba90fad0c137f4c1e7f5ba2e85"></a>

## Next pages — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.int / c452d7b42764 / 5

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-018.md#canonical-7c7d839a74b1dbce9e5eeb5e2072c6578f923d2e53c5ca1cbb463fce794b6d6e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-9b0b544b6e519fd2f2fd04ceb1d404d8c1fa9f1e1e5f095d06e30599522e0ebc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2505072d61d215003a49b098a6e27f9b6eab2065515dae6c6cf6e36ea619b142"></a>

## vmware.not_managed.node_list.interface_list.monitor — vmware.not_managed.node_list.interface_list.monitor / 2aa1c552ad6e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- vmware.not_managed.node_list.interface_list.monitor

<a id="canonical-56e0f54d2a2cf09a1ab0d8c40c9c48ceb380e7cc33fd6415c244215743fa5b7d"></a>

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

<a id="canonical-7a1f23dde2d4961956ac24da00a756779854f9bc0355333ca9480fd8931a300a"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.monitor / 2aa1c552ad6e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bc33c6c09423a85976e64a23be5ee0d38546525368a831060c2f24a502453431"></a>

## Next pages — vmware.not_managed.node_list.interface_list.monitor / 2aa1c552ad6e / 4

- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-9f32a06af7c948d6bb7e5d21fa67f7084d73a977512950f8ffa1bbb8cb23579c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bc88802c9ec8c6c6fe5d8e633260335a17581a305c5ba295c022a04f408502da"></a>

## vmware.not_managed.node_list.interface_list.monitor_disabled — vmware.not_managed.node_list.interface_list.monitor_disabled / 9027ffd9f87c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- vmware.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-501a7c0209108cbd13a022b3a3594a2ec4bb29631a0346ee47386b3a538c59be"></a>

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

<a id="canonical-3e6891a8bad39c89cfe2dc6f3bdd23a3d38850cce44cad36e4e49684792aeb50"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.monitor_disabled / 9027ffd9f87c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d835f7b6b467a1af118c661560e59c4aab2583bdef3eaa75e57fd9d61c4fc454"></a>

## Next pages — vmware.not_managed.node_list.interface_list.monitor_disabled / 9027ffd9f87c / 4

- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-5ebfba2a52831d712df58de767013f1f5a79056918e8170689ab47792a430ebf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-75ac9e1b1307f6de29989fa2b1b86dea9d48cf31176f86ec4adbd6cd6fd9fe72"></a>

## vmware.not_managed.node_list.interface_list.network_option — vmware.not_managed.node_list.interface_list.network_option / 6b1381aa02b5 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- vmware.not_managed.node_list.interface_list.network_option

<a id="canonical-00504c1341d002ffeca88b4c1b30adaf0af07ba9bc03b151115ca9c5021da19b"></a>

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

<a id="canonical-64d355f2f48b2f23cfe430a3055ef7f8433fdff71f81bfb3c24e0b14cde6c166"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.network_option / 6b1381aa02b5 / 3

- [site_local_inside_network](resources--securemesh_site_v2--reference--group-018.md#canonical-4ccecf0fef28899e603a7583ac8047043032cd08bbd07dc296d7c311aceae158): complete subsection reference.

- [site_local_network](resources--securemesh_site_v2--reference--group-018.md#canonical-059c2db0fa1ce5314827f119f5f83662ee1c3b5138b63221eee34e0b8135991e): complete subsection reference.

<a id="canonical-34070ddf3766f3d39142e0e26abdfb792de7bc3bde7be4194f96dcc5abedade0"></a>

## Next pages — vmware.not_managed.node_list.interface_list.network_option / 6b1381aa02b5 / 4

- [vmware.not_managed.node_list.interface_list.network_option.site_local_inside_network](resources--securemesh_site_v2--reference--group-018.md#canonical-4ccecf0fef28899e603a7583ac8047043032cd08bbd07dc296d7c311aceae158)
- [vmware.not_managed.node_list.interface_list.network_option.site_local_network](resources--securemesh_site_v2--reference--group-018.md#canonical-059c2db0fa1ce5314827f119f5f83662ee1c3b5138b63221eee34e0b8135991e)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-4ccecf0fef28899e603a7583ac8047043032cd08bbd07dc296d7c311aceae158"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b56b9f4a03d02f14ef8eb95935ea1e8b7a66d53fe11a07f9058cf159b83bb7f6"></a>

## vmware.not_managed.node_list.interface_list.network_option.site_local_inside_network — vmware.not_managed.node_list.interface_list.network_option.site_local_inside_net / 00284ff264b4 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [vmware.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-018.md#canonical-5ebfba2a52831d712df58de767013f1f5a79056918e8170689ab47792a430ebf)
- vmware.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-09fa4886fb37244efa17b4322939626cc696c3c9082846acec30f56234f873fe"></a>

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

<a id="canonical-2790ff320b9519e4a0edc49e16e5856dc792ba60c6a6bbead1ae2fe2cc112550"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.network_option.site_local_inside_net / 00284ff264b4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d1e825dedf392d845dd47553f32802cbb5ced801943a03d67e3510abcf557a53"></a>

## Next pages — vmware.not_managed.node_list.interface_list.network_option.site_local_inside_net / 00284ff264b4 / 4

- [vmware.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-018.md#canonical-5ebfba2a52831d712df58de767013f1f5a79056918e8170689ab47792a430ebf)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-059c2db0fa1ce5314827f119f5f83662ee1c3b5138b63221eee34e0b8135991e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e1361052ad96435315d7cf5bc3343426277fe2d6ab265778399500d44aa358c"></a>

## vmware.not_managed.node_list.interface_list.network_option.site_local_network — vmware.not_managed.node_list.interface_list.network_option.site_local_network / 0b7fc5d6ebbd / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [vmware.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-018.md#canonical-5ebfba2a52831d712df58de767013f1f5a79056918e8170689ab47792a430ebf)
- vmware.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-16fce005a5e12369aa328d023596da44e5a388838e99e4943986eecb289c0e37"></a>

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

<a id="canonical-7a699b5552fd8264f25931c06bc44e9cc36a5113a9091bf2e446a747cb14af14"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.network_option.site_local_network / 0b7fc5d6ebbd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6b54ae90975236ebc4219fe99dd75d4f1202a278f73faf4a16b3d38a8d159e1e"></a>

## Next pages — vmware.not_managed.node_list.interface_list.network_option.site_local_network / 0b7fc5d6ebbd / 4

- [vmware.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-018.md#canonical-5ebfba2a52831d712df58de767013f1f5a79056918e8170689ab47792a430ebf)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-40a0319a571432be435379ef7835965d96a611c1ab4e5ed201e4e13825a3a7d6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78a18345bfe86c5f78f06c14f5eaa90e11f5932efd16092f7fd01a212a251348"></a>

## vmware.not_managed.node_list.interface_list.no_ipv4_address — vmware.not_managed.node_list.interface_list.no_ipv4_address / bfe27e25f8d1 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- vmware.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-e8fadf46a9d030c47b552d439d5f6978a28beac4c9c819ac8da9b9bef0043511"></a>

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

<a id="canonical-4747e99a9cf66a35b8beaceec6a77ecb0be2b49e93dffaedd019726ce774eb45"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.no_ipv4_address / bfe27e25f8d1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-705cfcecfa3e117afdffed113c991ceea3babb7fb1a8b51c174dafe238edad5d"></a>

## Next pages — vmware.not_managed.node_list.interface_list.no_ipv4_address / bfe27e25f8d1 / 4

- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-28b9acf8bfb7a6aa850449e054ed55aa389f96f6e827218c26e7a0f14894fd36"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1bf98da4fc749b3efeea9e41f747498325d97efa3159f1898fdf010bac5b8080"></a>

## vmware.not_managed.node_list.interface_list.no_ipv6_address — vmware.not_managed.node_list.interface_list.no_ipv6_address / 26b7eed16520 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- vmware.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-33a66cd2ce18e22d9fb9f40bc3680f62447be3a341c03b830a5fd9d91b342dd5"></a>

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

<a id="canonical-49626848a71e500acb53eca3a1ecb2314dbfafaf2af1c14f6cfe38c5a79a1baa"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.no_ipv6_address / 26b7eed16520 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-96e2e9f85ae1fe9bc8cb452e911ff44e857f52965906609780a666e216074189"></a>

## Next pages — vmware.not_managed.node_list.interface_list.no_ipv6_address / 26b7eed16520 / 4

- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-ba07765ef645996ba29d90fea43545c54e0b60bee7de81bfbfc90b950ce32925"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72562ff60c61f1631bd63422d10a8dabc825f68ba6bcc51bf07b41ea98e11fa1"></a>

## vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled — vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_ / fc0a97a07ad0 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-7ae5b16fc6d70d55d257674dd7dcd3cf4e2ac10ab215f1186ab9d4b50be5acaf"></a>

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

<a id="canonical-cce4bf6b019c24743c8882dfd3d38e420ebedd08575379098a56290895b7c8dd"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_ / fc0a97a07ad0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7ca4be49b9a8cee0125aad6a634a764e332feea79104bb210eda8e30783fab4e"></a>

## Next pages — vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_ / fc0a97a07ad0 / 4

- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-85a9dab358e3ed5eac5cc0be0abf30d385589cd8cd221b8cd36a13a57b4da4a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a8465a61102cced5f8d44c2e81ea31c5345a4b13b39865eb634b7fcb37841862"></a>

## vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled — vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_ / 7b6011427430 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-218db329f62c86d9d2f4de07acc993d3d93513d1a40e6738e6c48687115df3cf"></a>

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

<a id="canonical-02444a8d18e83fa6eb39108df4fbb312e7adb77c548a573913a31ebe615a308c"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_ / 7b6011427430 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cc7eca00fedc670c14c24fbc14fed3973c1184b42bca003ca13c511d90d1ade3"></a>

## Next pages — vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_ / 7b6011427430 / 4

- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-f2d28f83ef14c14c7dfd32288a76cbab391c55e2ec73d753f0208ae3dd28ab80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c73144173fc66ad4fd9dc1fe63b78b126750ec947c636d9560af69a51c85fa5"></a>

## vmware.not_managed.node_list.interface_list.static_ip — vmware.not_managed.node_list.interface_list.static_ip / 375616c55767 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- vmware.not_managed.node_list.interface_list.static_ip

<a id="canonical-39ec63d41c082e170ba4980e8ac01f209ec46cc59863cafb3f08e4e0f8570be7"></a>

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

<a id="canonical-019a9df74ffb84074fb9570c896587365347574bdc73e948f08b07090c4684ca"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.static_ip / 375616c55767 / 3

<a id="canonical-ced0e95f19ca37a51977e0ffbdd1f22ac3622f4462f1a8e2d3f5e1d3b16bf6b3"></a>

<a id="canonical-00dc6c2505c52e330c7c1fd17b5c73689194551ba7a83871d5ef3b5965ee9616"></a>

## default_gw property — vmware.not_managed.node_list.interface_list.static_ip / 375616c55767 / 4

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

<a id="canonical-191c6d19b0c6de2538a6ac469b3f568e4c01ecbf01ff5fa450d3a157401703ab"></a>

<a id="canonical-b460e4b99595e2e4f3245e9a3426a0fa5a62a42eb2c029728a899d7613dcfd95"></a>

## dns_server property — vmware.not_managed.node_list.interface_list.static_ip / 375616c55767 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-37ab538c5ffa3e843f6b9080f7b69ef87a452c2b1d52147a4e38a3722e206c03"></a>

<a id="canonical-4ee4654c2b757cc6702050d236dece012e09841a74a45e9693e1adc409ad3e3d"></a>

## ip_address property — vmware.not_managed.node_list.interface_list.static_ip / 375616c55767 / 6

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

<a id="canonical-ee7a1de786e3ba523f1e325e1849823fb57a8d3804d13fa38e6ed9f4af712bf1"></a>

## Next pages — vmware.not_managed.node_list.interface_list.static_ip / 375616c55767 / 7

- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-4de96dd703c5f5036299256506315b1ffb4e0bd083b71d7fa30914ebd2b15d00"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-08a93eb24c6e2c4f919479a2be0e9b9c211080fd437bde815d50ff3c9bbf21fe"></a>

## vmware.not_managed.node_list.interface_list.static_ipv6_address — vmware.not_managed.node_list.interface_list.static_ipv6_address / 1e2b1994389e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- vmware.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-641a1a045adde40ee202e5e0247f9e08157b2f4c108aee9d219f4907b2704c8d"></a>

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

<a id="canonical-6dfcc054c6b24cab9747868c2b121d9e61738d2983d31c976f23db696e49f77d"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.static_ipv6_address / 1e2b1994389e / 3

- [cluster_static_ip](resources--securemesh_site_v2--reference--group-018.md#canonical-570b131ebfeefe21e49ab8bce1cf5a890f5ee41887aac81353750e4c45fb196e): complete subsection reference.

- [node_static_ip](resources--securemesh_site_v2--reference--group-018.md#canonical-7ace96446594d254c7159e8d265fd8fbbdf6827551bf51479d0bf5dedeb50c4d): complete subsection reference.

<a id="canonical-897b2ef7389699b4a4db3ad18f1d4828a7f76d983eaad18c2e903588b98afd23"></a>

## Next pages — vmware.not_managed.node_list.interface_list.static_ipv6_address / 1e2b1994389e / 4

- [vmware.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip](resources--securemesh_site_v2--reference--group-018.md#canonical-570b131ebfeefe21e49ab8bce1cf5a890f5ee41887aac81353750e4c45fb196e)
- [vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip](resources--securemesh_site_v2--reference--group-018.md#canonical-7ace96446594d254c7159e8d265fd8fbbdf6827551bf51479d0bf5dedeb50c4d)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-570b131ebfeefe21e49ab8bce1cf5a890f5ee41887aac81353750e4c45fb196e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-faa0d25897a75f856139ff7f55ff209bded5594b91bca4bb5dfd1361fef46f86"></a>

## vmware.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip — vmware.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_i / b7956391fc43 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [vmware.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-018.md#canonical-4de96dd703c5f5036299256506315b1ffb4e0bd083b71d7fa30914ebd2b15d00)
- vmware.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-5920c4d9051b2845155190027139a608059676aec273c43b63f38dfcbce4cd16"></a>

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

<a id="canonical-d41c6d0d150db88225ecbcb12c8053beed406df22b00d25bddfa78f16489799f"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_i / b7956391fc43 / 3

<a id="canonical-dfd543b0d522296bd8d9dd44f15bfc40f44855f40fd425b991a79b651d2393cb"></a>

<a id="canonical-8d1e503964749bc54b086496ac9e6a156c250246a8d12344af228191a16b51ff"></a>

## interface_ip_map property — vmware.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_i / b7956391fc43 / 4

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

<a id="canonical-ee3b7c9489f99de3a3d7acc0ea91ab529e641c54bcdfc7c220d49211d7176848"></a>

## Next pages — vmware.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_i / b7956391fc43 / 5

- [vmware.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-018.md#canonical-4de96dd703c5f5036299256506315b1ffb4e0bd083b71d7fa30914ebd2b15d00)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-7ace96446594d254c7159e8d265fd8fbbdf6827551bf51479d0bf5dedeb50c4d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-71733260b03cd60781c5681bc92a29909c2e9a08fca5f8ed7c70f91960c52df5"></a>

## vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip — vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / f23c5f750af6 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [vmware.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-018.md#canonical-4de96dd703c5f5036299256506315b1ffb4e0bd083b71d7fa30914ebd2b15d00)
- vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-81b8a9e826e090e8639b5b8efdff22f8c7dbdaa42cc54c68084be086137fde70"></a>

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

<a id="canonical-da974a515d56121b1ff207eaeea567310e551de0f5d1f383f07fda68d9933aab"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / f23c5f750af6 / 3

<a id="canonical-0c8f5d8901f2c44b625d1087409c267e28a34107828a07b0fcd202fceb7805e6"></a>

<a id="canonical-79136d4a91b09b73ccda2abccb011fc64aa5c2a4fd99d3d8392ddd733169d123"></a>

## default_gw property — vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / f23c5f750af6 / 4

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

<a id="canonical-f7ca676bd602f72f2eae8a5a10c6bc6dc9728c6e792587fd32190002a713819b"></a>

<a id="canonical-25eea43d91961892fd60264ef245d3b3cee7abb835f62da6ef944442d8bf8e12"></a>

## dns_server property — vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / f23c5f750af6 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-310ecc5a42c5c36d672fb314650b16b93217d362d69d21c2892ec94c717669e4"></a>

<a id="canonical-a726661c603fc7a700a507db82bda3aaa6a2cb5f3c7c0c5adf79ffc9c61fdc3c"></a>

## ip_address property — vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / f23c5f750af6 / 6

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

<a id="canonical-0db2750df4bfc6b3ddabfb7b6f66a19037c7b1490f23a52e6ece619877b3aa7b"></a>

## Next pages — vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / f23c5f750af6 / 7

- [vmware.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-018.md#canonical-4de96dd703c5f5036299256506315b1ffb4e0bd083b71d7fa30914ebd2b15d00)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-1a243e1484d9cde82529035943a86f47d148392a171c0b3e9bb0e156e7d96c89"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d04d64643ddff1c969f9526428370f454b83be49586a3403af1afdbb0b25ca0e"></a>

## vmware.not_managed.node_list.interface_list.vlan_interface — vmware.not_managed.node_list.interface_list.vlan_interface / fb6c3430655d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- vmware.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-0fa9600777fcff6179a7906870624e3752f7758cbcc5ae70dd9d9b8d9bf84b2b"></a>

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

<a id="canonical-24c2e58c492c73b218724a1d36ab31ad2cacd0c00c02232e88301b50265e4bf0"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.vlan_interface / fb6c3430655d / 3

<a id="canonical-b3c86517bfc593abb86af01173070a9c77ec15565413c6e59c680498ae29af69"></a>

<a id="canonical-530b9512527c1bcbd4a2de344bc33abcd501ddbddce40910ba4f025b2d23d5ea"></a>

## device property — vmware.not_managed.node_list.interface_list.vlan_interface / fb6c3430655d / 4

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

<a id="canonical-056f5b5479e868add3f1ce221349628ee5357af0b69e70696afc0fbd3ea6904c"></a>

<a id="canonical-c9e514e2eb2e88c5f368828c0e0e36d449fc97ffa4ed6da00d9cf34bf1682403"></a>

## vlan_id property — vmware.not_managed.node_list.interface_list.vlan_interface / fb6c3430655d / 5

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

<a id="canonical-3e2fc57ac39dfd17d8cf161801ae82cb14d5bde10db272d07a463b57fc3f3b3a"></a>

## Next pages — vmware.not_managed.node_list.interface_list.vlan_interface / fb6c3430655d / 6

- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
