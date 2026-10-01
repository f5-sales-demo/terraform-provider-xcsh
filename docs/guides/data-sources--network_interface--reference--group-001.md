---
page_title: "xcsh_network_interface reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_interface reference."
---

# xcsh_network_interface reference

<a id="canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae9b5b665537083935b95765b5ffb3ac4dd175f5f579c91319ac70c608d5cba8"></a>

## Property reference — Property reference / 6fd94a001801 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- Property reference

<a id="canonical-2db64f0c4e8ace851c35efa02893047332f50817ddd7cf00f77587439ab94b41"></a>

## Direct properties — Property reference / 6fd94a001801 / 3

<a id="canonical-e53130067d630a31f54bdd106ed9935eebcc2dc4a9deadceca1afece306b0561"></a>

<a id="canonical-06acebde55178acc4a06c1a0ee53c09eac832b82cf434e8cdc6f211e67392182"></a>

## annotations property — Property reference / 6fd94a001801 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

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

- [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-95403d00f6e6ec9eec91753b5f56180465405be1b63e1e5adc64309fbe9777b9): complete subsection reference.

- [dedicated_management_interface](data-sources--network_interface--reference--group-001.md#canonical-0d50dbb76d469213d5008157d7c1d5c05b12a55c0bb2bc4aa2b2a050d06fcc72): complete subsection reference.

<a id="canonical-fc751d8200522fb8bf5c162223cbf5cd9f5102c84c4c909b15f2e2e8560d44ff"></a>

<a id="canonical-3b4abc60943342eb9cfb2414a6a1ad47b0e201b8041e5c80b717b1bdeb581753"></a>

## description property — Property reference / 6fd94a001801 / 5

Type: `"string"`. Computed.

Description of the NetworkInterface.

Upstream description:

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

- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6): complete subsection reference.

<a id="canonical-749249d40058881f4bff9049202b8703bc6c52c57d0e07f5212746bed49675ee"></a>

<a id="canonical-a979a64c57b3536011b0f553e80b6ee073bdb19d4a29fdb19c3aae5790d1c981"></a>

## id property — Property reference / 6fd94a001801 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-f0e7384252877ef59d4201bcadf2f7b03b6b71f0065242f00b7fbffe288bf0c4"></a>

<a id="canonical-f51223a82f6728b5956b4eced39e094cf7cd3a69eb66410d959d1fdec7a3b013"></a>

## labels property — Property reference / 6fd94a001801 / 7

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

- [layer2_interface](data-sources--network_interface--reference--group-002.md#canonical-23fb46783886bf9cc3aab8b71636b36b489f59cff700e25eee93f42b2c95b70c): complete subsection reference.

<a id="canonical-14eb3f33e7743e975fb13997dbf4e99ece8abe77c68ceacb3539dea23511a268"></a>

<a id="canonical-4cc965ec7d6635330ffcf0f1b301a645c36632fdab18c8ce29fa61bfe687fd07"></a>

## name property — Property reference / 6fd94a001801 / 8

Type: `"string"`. Required.

Name of the NetworkInterface.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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

<a id="canonical-0d80bd4603e6d6448a8471a95be0bb5e12a55e8645edf25afa15c7ee4bf7f01a"></a>

<a id="canonical-b0249e6358db2206b3f046c5ddf2366ba1c6e032d9f58b0b2a32a4cd1913b08a"></a>

## namespace property — Property reference / 6fd94a001801 / 9

Type: `"string"`. Required.

Namespace where the NetworkInterface exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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

- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-4c4a8d0d91a091335e2c8a7c47c83f1fbdc5b32befca2a5bf820c8b278cccbb7): complete subsection reference.

<a id="canonical-acf582de76d4eab036ec909baddf0b4e55ce47af4c31f10846212e6b768281e9"></a>

## All schema paths — Property reference / 6fd94a001801 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--network_interface--reference--group-001.md#canonical-e53130067d630a31f54bdd106ed9935eebcc2dc4a9deadceca1afece306b0561) |
| `dedicated_interface` | [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-b775a2292fc7b58232035ca5d6c6705af10630f7da746f8ce9eee881fcc0167c) |
| `dedicated_interface.cluster` | [dedicated_interface.cluster](data-sources--network_interface--reference--group-001.md#canonical-18bdafd10a866c0d01882573832000d0491665b1c8c4bc5382d8da137360355f) |
| `dedicated_interface.device` | [dedicated_interface.device](data-sources--network_interface--reference--group-001.md#canonical-1570e62522b5bb40f4406c6291d903656da8932eaea92e4d5251ee871d4ddffc) |
| `dedicated_interface.is_primary` | [dedicated_interface.is_primary](data-sources--network_interface--reference--group-001.md#canonical-2285b48e8622dcf92ed3d02a5acc3628de03c7339a0a9edf197e6488ee27ad8b) |
| `dedicated_interface.monitor` | [dedicated_interface.monitor](data-sources--network_interface--reference--group-001.md#canonical-41cb2fbd437a67a7d79b4c30fb6fe8a62a78e518885fc0855d8826feda1c4e88) |
| `dedicated_interface.monitor_disabled` | [dedicated_interface.monitor_disabled](data-sources--network_interface--reference--group-001.md#canonical-22e311a06f46d006645f9f76af7be36ae14cc65d4f70f2bf454926fc37b3ab5d) |
| `dedicated_interface.mtu` | [dedicated_interface.mtu](data-sources--network_interface--reference--group-001.md#canonical-b90c44ec0a78b5f12b0dc003afce048b4498b349ed749e0690258cb0d18f8914) |
| `dedicated_interface.node` | [dedicated_interface.node](data-sources--network_interface--reference--group-001.md#canonical-79c0aebd42a508357ebca3f9d0aa4febd3c0920b2f9ccb35653c7fbff3554a38) |
| `dedicated_interface.not_primary` | [dedicated_interface.not_primary](data-sources--network_interface--reference--group-001.md#canonical-54c870762554406cec20c1762de3d4acc074ab2a6baf7caf1b19bef13bb83100) |
| `dedicated_interface.priority` | [dedicated_interface.priority](data-sources--network_interface--reference--group-001.md#canonical-84ec0e14e17241d66209cbe744fd01d6e94874d8eee0be0b31319aafa53977d8) |
| `dedicated_management_interface` | [dedicated_management_interface](data-sources--network_interface--reference--group-001.md#canonical-d17ecc49710b63ca0b55e9db9eb444523afad5db4dcff17034553c581b87ff19) |
| `dedicated_management_interface.cluster` | [dedicated_management_interface.cluster](data-sources--network_interface--reference--group-001.md#canonical-33b197e238e944da8ab05edb4edf6f21cac78c29dc785cf2dee105c7c2d9696e) |
| `dedicated_management_interface.device` | [dedicated_management_interface.device](data-sources--network_interface--reference--group-001.md#canonical-85bbbe63de852f01c219002005b11108d3e281b9073a114d9587fe1c6090d514) |
| `dedicated_management_interface.mtu` | [dedicated_management_interface.mtu](data-sources--network_interface--reference--group-001.md#canonical-60e071a911ff8140320da7f0de36994f2f246dfc2bf677123d2f02bf72094f25) |
| `dedicated_management_interface.node` | [dedicated_management_interface.node](data-sources--network_interface--reference--group-001.md#canonical-1c3f6fd643be3fee5425e541154398616324b9882537ab84f99c832d998d7456) |
| `description` | [description](data-sources--network_interface--reference--group-001.md#canonical-fc751d8200522fb8bf5c162223cbf5cd9f5102c84c4c909b15f2e2e8560d44ff) |
| `ethernet_interface` | [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-ed977f370a7a88690700eecaf0aa97a386c9f0713b4ba438353c228c92662644) |
| `ethernet_interface.cluster` | [ethernet_interface.cluster](data-sources--network_interface--reference--group-001.md#canonical-201fea49c35ca3602d9b075febafab6838f60d7cfd4b181029ec593ad694ee22) |
| `ethernet_interface.device` | [ethernet_interface.device](data-sources--network_interface--reference--group-001.md#canonical-515937a1d7ff2e6ec78a8f95c89397c9ff70f94f3881861687170c15eff22e42) |
| `ethernet_interface.dhcp_client` | [ethernet_interface.dhcp_client](data-sources--network_interface--reference--group-001.md#canonical-ba3b1c0472e207ddf21e7d9c971a4e358d3138e277dcb3bdcd382f2b58b8f5c2) |
| `ethernet_interface.dhcp_server` | [ethernet_interface.dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-210b4c2ee1490b37475d97f74261dbe4a0443e1e376b29c3ac077ba47a18ac9d) |
| `ethernet_interface.dhcp_server.automatic_from_end` | [ethernet_interface.dhcp_server.automatic_from_end](data-sources--network_interface--reference--group-001.md#canonical-f870a0623ad6f1a2e5ab9cec792b7fcb8f40af81bdc68fa9deb3b2c9c1d63295) |
| `ethernet_interface.dhcp_server.automatic_from_start` | [ethernet_interface.dhcp_server.automatic_from_start](data-sources--network_interface--reference--group-001.md#canonical-f7ee4e104f435a975774740af7a599d2eba5692221bd3c9d2716fc63f5a872ae) |
| `ethernet_interface.dhcp_server.dhcp_networks` | [ethernet_interface.dhcp_server.dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-b23646f331a3fac94d2281e9410566c6ae09adba9a4a0b8afdaf4316f8e76649) |
| `ethernet_interface.dhcp_server.dhcp_networks.dgw_address` | [ethernet_interface.dhcp_server.dhcp_networks.dgw_address](data-sources--network_interface--reference--group-001.md#canonical-3c3f11a9e90ea394084c24faae70a3f76dc97228d218a6aff29d91826b697511) |
| `ethernet_interface.dhcp_server.dhcp_networks.dns_address` | [ethernet_interface.dhcp_server.dhcp_networks.dns_address](data-sources--network_interface--reference--group-001.md#canonical-4991ef0c385c11b86862dc52128f6266ca15e003becd3a3df1ddee5956f40889) |
| `ethernet_interface.dhcp_server.dhcp_networks.first_address` | [ethernet_interface.dhcp_server.dhcp_networks.first_address](data-sources--network_interface--reference--group-001.md#canonical-20789423749358083c255998675923223c280cffd71f654eef3a09a130ae32d1) |
| `ethernet_interface.dhcp_server.dhcp_networks.last_address` | [ethernet_interface.dhcp_server.dhcp_networks.last_address](data-sources--network_interface--reference--group-001.md#canonical-2c6446688cb3d6cee12518a2d5985d18037a1deeef3842a5e667670cb541c2c0) |
| `ethernet_interface.dhcp_server.dhcp_networks.network_prefix` | [ethernet_interface.dhcp_server.dhcp_networks.network_prefix](data-sources--network_interface--reference--group-001.md#canonical-81e4132f1a0fa34120636ae018794f07f1bda1aa44c8fb5890963c00bb0bc144) |
| `ethernet_interface.dhcp_server.dhcp_networks.pool_settings` | [ethernet_interface.dhcp_server.dhcp_networks.pool_settings](data-sources--network_interface--reference--group-001.md#canonical-c0f86457340f34560b8ee841d1bb3fb0e1584a0b084d897da9b9f917f24bf26c) |
| `ethernet_interface.dhcp_server.dhcp_networks.pools` | [ethernet_interface.dhcp_server.dhcp_networks.pools](data-sources--network_interface--reference--group-001.md#canonical-850f6806611f42df587b791194b83a05e4ac988523043a06187d6cf672c20941) |
| `ethernet_interface.dhcp_server.dhcp_networks.pools.end_ip` | [ethernet_interface.dhcp_server.dhcp_networks.pools.end_ip](data-sources--network_interface--reference--group-001.md#canonical-904673383562ab7085d12ffa1ea64f113f8379b9e3861ebdd05b6c6f84810c41) |
| `ethernet_interface.dhcp_server.dhcp_networks.pools.exclude` | [ethernet_interface.dhcp_server.dhcp_networks.pools.exclude](data-sources--network_interface--reference--group-001.md#canonical-252dfe698fb21347a42f99af9f6bb1402e55f8665fb5a694882f8c8d594c799e) |
| `ethernet_interface.dhcp_server.dhcp_networks.pools.start_ip` | [ethernet_interface.dhcp_server.dhcp_networks.pools.start_ip](data-sources--network_interface--reference--group-001.md#canonical-bcac4a5e8c1b7093d3394f0c606c6a6a14cd60aab9f28df04d3b727ddd2704d3) |
| `ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw` | [ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw](data-sources--network_interface--reference--group-001.md#canonical-006145690a9f64a2487279ce5d889a3ec199f23d4abcd530213dcbe213e00bd1) |
| `ethernet_interface.dhcp_server.dhcp_option82_tag` | [ethernet_interface.dhcp_server.dhcp_option82_tag](data-sources--network_interface--reference--group-001.md#canonical-a0a138aaa34162ee8b1096fe5adba1537346ed4b9fe6cc662123b10bba60b6e8) |
| `ethernet_interface.dhcp_server.fixed_ip_map` | [ethernet_interface.dhcp_server.fixed_ip_map](data-sources--network_interface--reference--group-001.md#canonical-033991a4fa71b049d0be2af4f224520d769c4cb091816a62d95c2e86fdd963c6) |
| `ethernet_interface.dhcp_server.interface_ip_map` | [ethernet_interface.dhcp_server.interface_ip_map](data-sources--network_interface--reference--group-001.md#canonical-f6891f1d0aad9a25c52fc2448db5308b903b19c5cb46da65200b06c314b48f93) |
| `ethernet_interface.dhcp_server.interface_ip_map.interface_ip_map` | [ethernet_interface.dhcp_server.interface_ip_map.interface_ip_map](data-sources--network_interface--reference--group-001.md#canonical-0835d94dc9f73c10105cad94a1d049f25ed1557565b74d2dba4b0893b22fb6b4) |
| `ethernet_interface.ipv6_auto_config` | [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-32224430e9c6613018c1114103f99d733cea2983351d2724f994036cfbcb7887) |
| `ethernet_interface.ipv6_auto_config.host` | [ethernet_interface.ipv6_auto_config.host](data-sources--network_interface--reference--group-001.md#canonical-fda544d3e19b3fac3f12b92213c9e9a400014e95bb43741194d936ddd5f3fef2) |
| `ethernet_interface.ipv6_auto_config.router` | [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-e13ef8cbfa67809aa7dd7755779911e3b1d103dbc6926042520f6eb2059c9150) |
| `ethernet_interface.ipv6_auto_config.router.dns_config` | [ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--network_interface--reference--group-001.md#canonical-f09a9cc868d06d2c2e828a276d3fbe459b52888341535d2923b96548ec0d0d9c) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.configured_list` | [ethernet_interface.ipv6_auto_config.router.dns_config.configured_list](data-sources--network_interface--reference--group-001.md#canonical-a502f4c063385d62c6c218e43756fc1d77903d8459f0f345afecd53981f01921) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.configured_list.dns_list` | [ethernet_interface.ipv6_auto_config.router.dns_config.configured_list.dns_list](data-sources--network_interface--reference--group-001.md#canonical-5412181c39d943907635707762bc2af4d4760693ce9b5d69af88098eef400a08) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.local_dns` | [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](data-sources--network_interface--reference--group-001.md#canonical-a34ac4372c642d785b3e0db93d8196631b3bf100cc91001c0ebafc2ae99a6e31) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.configured_address` | [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.configured_address](data-sources--network_interface--reference--group-001.md#canonical-f5f518197ab61febee5110baa0a89abaab204527e1984c9733707ba3f83b27e8) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address` | [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address](data-sources--network_interface--reference--group-001.md#canonical-c40793311ab5385a2c6299c22d68458eab8fff1e5e480f6a4bb168ecda8716f3) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address` | [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address](data-sources--network_interface--reference--group-001.md#canonical-1f9e39e86c18841990b8c123986681a19fdd85cd70aceb34dd40ffa16a9f2660) |
| `ethernet_interface.ipv6_auto_config.router.network_prefix` | [ethernet_interface.ipv6_auto_config.router.network_prefix](data-sources--network_interface--reference--group-001.md#canonical-0a05f6edec7fb66ad7cda39be06835d4ad3e36eb08e940bcb0d544e319cc6b76) |
| `ethernet_interface.ipv6_auto_config.router.stateful` | [ethernet_interface.ipv6_auto_config.router.stateful](data-sources--network_interface--reference--group-001.md#canonical-2591ffd1665b732807013fd93c9f86ca06f17ee8e938ae66da9dd159971482fd) |
| `ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end` | [ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end](data-sources--network_interface--reference--group-001.md#canonical-e6bf166e9c6f517ed0350a9ab9c075ce19306bf1c9457170ba67d54344a142b4) |
| `ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start` | [ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start](data-sources--network_interface--reference--group-001.md#canonical-dfad90a1bdd79c109d170691ae329491e679ec52962da2c2e5d2674cf756feb5) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-255f756ecdaa570a62428a25b228d26ebcb5e7ef0e235c642dd01f6e499eb9c0) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix](data-sources--network_interface--reference--group-001.md#canonical-a8e0d9ecec0e26a1ebebc8d8cbeaf1a0850e901a7cfafaceb9bb2a07933a9dea) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings](data-sources--network_interface--reference--group-001.md#canonical-597be91254c52a1729afcfbc635fd7d710f548d767890229025478134b2e2fc1) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools](data-sources--network_interface--reference--group-001.md#canonical-0fa37a54c99f3ed748a08f97d52493cf9478a35fa25e2a2a406d5c4b0aad81f2) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip](data-sources--network_interface--reference--group-001.md#canonical-5896178f188c5e81b9402c74f4f6f95842b42068aa6f7ab43a3ed2d21b7ed826) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip](data-sources--network_interface--reference--group-001.md#canonical-e535cf166cbfd3d08eb25adbf6af51e5bb048aab0bad70c4ac1f218e76597344) |
| `ethernet_interface.ipv6_auto_config.router.stateful.fixed_ip_map` | [ethernet_interface.ipv6_auto_config.router.stateful.fixed_ip_map](data-sources--network_interface--reference--group-001.md#canonical-2291391ad5516c2e29b15c7d18a96cdf2e5f6fc00d17c94a563afe7defda4938) |
| `ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map` | [ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map](data-sources--network_interface--reference--group-001.md#canonical-145942f17d945d03a0cb74817243d3d542162823f9c68eb309d2a70f904b5204) |
| `ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` | [ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map](data-sources--network_interface--reference--group-001.md#canonical-88eac661e4c606f93ea3f7fe5216657ecf2074e3946f49a81a7ca2f3c1ccca81) |
| `ethernet_interface.is_primary` | [ethernet_interface.is_primary](data-sources--network_interface--reference--group-001.md#canonical-272504013cf4717e71c2473a90ca18d8bf7556be3911889e94adf5e4ea7fc48e) |
| `ethernet_interface.monitor` | [ethernet_interface.monitor](data-sources--network_interface--reference--group-001.md#canonical-4fde18526237441ca005ff4c402e750236a0f624b16ba704d600f1ddd48029c5) |
| `ethernet_interface.monitor_disabled` | [ethernet_interface.monitor_disabled](data-sources--network_interface--reference--group-001.md#canonical-6f5ba37cc5702ea92aaf9f6b246bb9e8b7217b30d6c2ff406f5fae20d2744f5c) |
| `ethernet_interface.mtu` | [ethernet_interface.mtu](data-sources--network_interface--reference--group-001.md#canonical-ca2651856c24606f1d5de70742025141929261bb489dd34f4090ac5f4eeab763) |
| `ethernet_interface.no_ipv6_address` | [ethernet_interface.no_ipv6_address](data-sources--network_interface--reference--group-001.md#canonical-f4aa795ad9870d874d3a77605489ac4c0736e7284aeda008217a1f7cdcdb5b42) |
| `ethernet_interface.node` | [ethernet_interface.node](data-sources--network_interface--reference--group-001.md#canonical-3226a284b5cdfee4636cc54430f7d4e154054cbd4f69fd0ac9d349dae5504bbf) |
| `ethernet_interface.not_primary` | [ethernet_interface.not_primary](data-sources--network_interface--reference--group-001.md#canonical-ad4ed6be239a874e3e0b92500208d19e743225083a42389f010a40a5a24c213c) |
| `ethernet_interface.priority` | [ethernet_interface.priority](data-sources--network_interface--reference--group-001.md#canonical-7001aed971c33b5e31055b05464905d07d805a07fd8facfd759248dea27ab0f4) |
| `ethernet_interface.site_local_inside_network` | [ethernet_interface.site_local_inside_network](data-sources--network_interface--reference--group-001.md#canonical-d6f24b2e1d87f515fb70e214d3358fdf3f54ca4d10e4a6f78e82c5b61371a5fb) |
| `ethernet_interface.site_local_network` | [ethernet_interface.site_local_network](data-sources--network_interface--reference--group-001.md#canonical-63e80dfff901fb24c8718cd4493d86e2dd790510027ecb54870af1b82bbbbb54) |
| `ethernet_interface.static_ip` | [ethernet_interface.static_ip](data-sources--network_interface--reference--group-001.md#canonical-f07c55637c9b1a25b38b01ad551f29bd820199ecbe30902bb07460d61422d73b) |
| `ethernet_interface.static_ip.cluster_static_ip` | [ethernet_interface.static_ip.cluster_static_ip](data-sources--network_interface--reference--group-001.md#canonical-4d721af642105b494b7581adf6dc86c347732d26954cb293f27cee0066eb0a78) |
| `ethernet_interface.static_ip.cluster_static_ip.interface_ip_map` | [ethernet_interface.static_ip.cluster_static_ip.interface_ip_map](data-sources--network_interface--reference--group-001.md#canonical-c42e1280024e09675d4781e52cc2db305557fd8f02882de66e4d059b389cf53b) |
| `ethernet_interface.static_ip.node_static_ip` | [ethernet_interface.static_ip.node_static_ip](data-sources--network_interface--reference--group-001.md#canonical-e850cf6c4e0a238bbf12185ce3bd3d735ee4634e1cd60116f505287322196270) |
| `ethernet_interface.static_ip.node_static_ip.default_gw` | [ethernet_interface.static_ip.node_static_ip.default_gw](data-sources--network_interface--reference--group-001.md#canonical-d428695040f29a7e848ca0b1566669c05e8e929bf348a5ca3c22c074ecdb25f3) |
| `ethernet_interface.static_ip.node_static_ip.dns_server` | [ethernet_interface.static_ip.node_static_ip.dns_server](data-sources--network_interface--reference--group-001.md#canonical-5c569cd502d3fc0b7d745215efbfb40af475dc3d8c7372fef8a9e53bd4306d07) |
| `ethernet_interface.static_ip.node_static_ip.ip_address` | [ethernet_interface.static_ip.node_static_ip.ip_address](data-sources--network_interface--reference--group-001.md#canonical-935e486a053f3989926320ac9d25cf35b2e2fe99a39bbe8d633d8ed3b0740c21) |
| `ethernet_interface.static_ipv6_address` | [ethernet_interface.static_ipv6_address](data-sources--network_interface--reference--group-002.md#canonical-86190bd308c892284953123125e23154cf85a8118f24d845f5aec209bb3290dd) |
| `ethernet_interface.static_ipv6_address.cluster_static_ip` | [ethernet_interface.static_ipv6_address.cluster_static_ip](data-sources--network_interface--reference--group-002.md#canonical-9d10e91331d1012afdd8e980d8eeb35801a48aa81a393d21d7f2aefaa314245d) |
| `ethernet_interface.static_ipv6_address.cluster_static_ip.interface_ip_map` | [ethernet_interface.static_ipv6_address.cluster_static_ip.interface_ip_map](data-sources--network_interface--reference--group-002.md#canonical-a5308dbc5f6a7372182239dfeb7b4f362b6a5e1edc28cd2dfc4d765c23afae68) |
| `ethernet_interface.static_ipv6_address.node_static_ip` | [ethernet_interface.static_ipv6_address.node_static_ip](data-sources--network_interface--reference--group-002.md#canonical-4e094e28389305c09b9ee16dd232f1089081aff3263b393d662c4ba4159518c8) |
| `ethernet_interface.static_ipv6_address.node_static_ip.default_gw` | [ethernet_interface.static_ipv6_address.node_static_ip.default_gw](data-sources--network_interface--reference--group-002.md#canonical-809520a2997efbe19d87d992581ed570bee25918b8715d4492ddb8a9e20a5ef2) |
| `ethernet_interface.static_ipv6_address.node_static_ip.dns_server` | [ethernet_interface.static_ipv6_address.node_static_ip.dns_server](data-sources--network_interface--reference--group-002.md#canonical-83ae3ccc28db1dc1749925a132188110d0b69120b67cef72df705e68e7a26c5f) |
| `ethernet_interface.static_ipv6_address.node_static_ip.ip_address` | [ethernet_interface.static_ipv6_address.node_static_ip.ip_address](data-sources--network_interface--reference--group-002.md#canonical-231c2ad72141a37118c291802237a8d4328303ad83f01537d7faab39b9a481e1) |
| `ethernet_interface.storage_network` | [ethernet_interface.storage_network](data-sources--network_interface--reference--group-002.md#canonical-aaa72a8a6c3e64516aefd2801c0e422a127e9315fdc0a5fc047c86fbf46edef4) |
| `ethernet_interface.untagged` | [ethernet_interface.untagged](data-sources--network_interface--reference--group-002.md#canonical-fe8ae3fbcba88305ec8356fdc919a3be1d1283017e35bbc7d8ff2c8c5151c112) |
| `ethernet_interface.vlan_id` | [ethernet_interface.vlan_id](data-sources--network_interface--reference--group-001.md#canonical-c72d0e2d5a111107a6f35cb3ce8f6a45841211243c03d5c6d6ad7c75f3ca6ed1) |
| `id` | [id](data-sources--network_interface--reference--group-001.md#canonical-749249d40058881f4bff9049202b8703bc6c52c57d0e07f5212746bed49675ee) |
| `labels` | [labels](data-sources--network_interface--reference--group-001.md#canonical-f0e7384252877ef59d4201bcadf2f7b03b6b71f0065242f00b7fbffe288bf0c4) |
| `layer2_interface` | [layer2_interface](data-sources--network_interface--reference--group-002.md#canonical-e25c4b23e2e1ec1252a64394f70c4aac2e7034a01010527a43b3b66f0f420ceb) |
| `layer2_interface.l2sriov_interface` | [layer2_interface.l2sriov_interface](data-sources--network_interface--reference--group-002.md#canonical-de8091034db81fd6f72be817732751028b91672188945e42bf9d4f12a32d5886) |
| `layer2_interface.l2sriov_interface.device` | [layer2_interface.l2sriov_interface.device](data-sources--network_interface--reference--group-002.md#canonical-563f58745d2099c61019eb401c116dfe0d497b6df9172f5cfcb1a5faf8eb3123) |
| `layer2_interface.l2sriov_interface.untagged` | [layer2_interface.l2sriov_interface.untagged](data-sources--network_interface--reference--group-002.md#canonical-85b06d296fc171d55fb3d610af76da9119fa7696e7398c4c526b35b20206d3ea) |
| `layer2_interface.l2sriov_interface.vlan_id` | [layer2_interface.l2sriov_interface.vlan_id](data-sources--network_interface--reference--group-002.md#canonical-7cff783164227e29be43971bd10b5bc399a9f605a966bc527b117df227fd5b90) |
| `layer2_interface.l2vlan_interface` | [layer2_interface.l2vlan_interface](data-sources--network_interface--reference--group-002.md#canonical-2e684f5ff27b6ff85205c9274e35ea3fc52289b5fd480adf4968919f7e3cd424) |
| `layer2_interface.l2vlan_interface.device` | [layer2_interface.l2vlan_interface.device](data-sources--network_interface--reference--group-002.md#canonical-172f0c411c33895503d64a66e56d1132a4684337e357d3fac2821f6a6c6fb940) |
| `layer2_interface.l2vlan_interface.vlan_id` | [layer2_interface.l2vlan_interface.vlan_id](data-sources--network_interface--reference--group-002.md#canonical-0a29bba28098c7b2206fff78e66f9e082c978fd2b99393d67e990b0ad16eb13d) |
| `layer2_interface.l2vlan_slo_interface` | [layer2_interface.l2vlan_slo_interface](data-sources--network_interface--reference--group-002.md#canonical-a76b9c4cb3f24fecfc400619b2fa619033228f6f8ed2047ac0e387fd3350225e) |
| `layer2_interface.l2vlan_slo_interface.vlan_id` | [layer2_interface.l2vlan_slo_interface.vlan_id](data-sources--network_interface--reference--group-002.md#canonical-975c0ff9b818e5b3dfb430cc3daeb9944a1aa8f0a86e2c212ae2a80a3ec85d55) |
| `name` | [name](data-sources--network_interface--reference--group-001.md#canonical-14eb3f33e7743e975fb13997dbf4e99ece8abe77c68ceacb3539dea23511a268) |
| `namespace` | [namespace](data-sources--network_interface--reference--group-001.md#canonical-0d80bd4603e6d6448a8471a95be0bb5e12a55e8645edf25afa15c7ee4bf7f01a) |
| `tunnel_interface` | [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-2269d485ee9d6abbab5bf7aa51105026e13fedb756d7e5e7f41e68c320cf0cef) |
| `tunnel_interface.mtu` | [tunnel_interface.mtu](data-sources--network_interface--reference--group-002.md#canonical-b16af4d8d658af6a2d6ce94ed4d95227a38d37ef20053ddb91cc794880697d2e) |
| `tunnel_interface.node` | [tunnel_interface.node](data-sources--network_interface--reference--group-002.md#canonical-affaa248268d9aa661fe2672faf0f9d1ab03a30897a5e02bb9c08cb6767f77e1) |
| `tunnel_interface.priority` | [tunnel_interface.priority](data-sources--network_interface--reference--group-002.md#canonical-9a463d81bb6733b17adbaca9fad7dbbf2ce260e621600eb998569bfd370eafef) |
| `tunnel_interface.site_local_inside_network` | [tunnel_interface.site_local_inside_network](data-sources--network_interface--reference--group-002.md#canonical-ddfce1eb4985dad3d94241abc3dbd44edb756ed2ea830471afab9bce2b48151b) |
| `tunnel_interface.site_local_network` | [tunnel_interface.site_local_network](data-sources--network_interface--reference--group-002.md#canonical-40da0852fce20e65a6795890ccf3851f55b9ea511f4eebe016f215074d076519) |
| `tunnel_interface.static_ip` | [tunnel_interface.static_ip](data-sources--network_interface--reference--group-002.md#canonical-d759e4dee3392c091389e994fa5b71251ac2d86bf489e5f5b1b782eebe1bfddc) |
| `tunnel_interface.static_ip.cluster_static_ip` | [tunnel_interface.static_ip.cluster_static_ip](data-sources--network_interface--reference--group-002.md#canonical-3951090c5704fe4b24b3845979b1af9bc587e36046f2e9660a54a3d7c9858fa4) |
| `tunnel_interface.static_ip.cluster_static_ip.interface_ip_map` | [tunnel_interface.static_ip.cluster_static_ip.interface_ip_map](data-sources--network_interface--reference--group-002.md#canonical-d17252b54bc679bb9c9fd0710f654221d0f17d369b1972f405a9d0bd5875239f) |
| `tunnel_interface.static_ip.node_static_ip` | [tunnel_interface.static_ip.node_static_ip](data-sources--network_interface--reference--group-002.md#canonical-75fba8069b09750e46990355a84447034150dad44ad0c2c2d3aac0d9a244bc1f) |
| `tunnel_interface.static_ip.node_static_ip.default_gw` | [tunnel_interface.static_ip.node_static_ip.default_gw](data-sources--network_interface--reference--group-002.md#canonical-0c78a0677c44b526d506638e3630ef7ceeeea99eda1bdcd6ddc6572209f368c9) |
| `tunnel_interface.static_ip.node_static_ip.dns_server` | [tunnel_interface.static_ip.node_static_ip.dns_server](data-sources--network_interface--reference--group-002.md#canonical-b365993c73b1fcce6c37f66fcba25f9334059b0ef3324ef83306a3198dde137b) |
| `tunnel_interface.static_ip.node_static_ip.ip_address` | [tunnel_interface.static_ip.node_static_ip.ip_address](data-sources--network_interface--reference--group-002.md#canonical-6b22ad6f23f80e4a5a4fb31288968511e7034e5ed3d7940915a1037cd63c8418) |
| `tunnel_interface.tunnel` | [tunnel_interface.tunnel](data-sources--network_interface--reference--group-002.md#canonical-37d22a39d0a701213a15d6b37e005801363a06fb154abbc1c5d88afa0a13f6af) |
| `tunnel_interface.tunnel.name` | [tunnel_interface.tunnel.name](data-sources--network_interface--reference--group-002.md#canonical-ad6cb919c8091cd74ee36c3b3661456de78020eb9a8a38dc8fa8cd87843d5055) |
| `tunnel_interface.tunnel.namespace` | [tunnel_interface.tunnel.namespace](data-sources--network_interface--reference--group-002.md#canonical-9ca364562be93c9d662267a70d9bdb137089324fd7b7590d98fa74430fcf48fc) |
| `tunnel_interface.tunnel.tenant` | [tunnel_interface.tunnel.tenant](data-sources--network_interface--reference--group-002.md#canonical-c5a4034dbd7582f802d4e530cb65507fd8d1368bd2745ef20e3a393799187e7d) |

<a id="canonical-4f6afa79717f8c575a663c47f0dde8da937610e97559f8ff60c70f298e90b4e3"></a>

## Next pages — Property reference / 6fd94a001801 / 11

- [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-95403d00f6e6ec9eec91753b5f56180465405be1b63e1e5adc64309fbe9777b9)
- [dedicated_management_interface](data-sources--network_interface--reference--group-001.md#canonical-0d50dbb76d469213d5008157d7c1d5c05b12a55c0bb2bc4aa2b2a050d06fcc72)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [layer2_interface](data-sources--network_interface--reference--group-002.md#canonical-23fb46783886bf9cc3aab8b71636b36b489f59cff700e25eee93f42b2c95b70c)
- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-4c4a8d0d91a091335e2c8a7c47c83f1fbdc5b32befca2a5bf820c8b278cccbb7)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-95403d00f6e6ec9eec91753b5f56180465405be1b63e1e5adc64309fbe9777b9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fefdbf2a0c29563b32dce40951c435272ae4aa0b69aaa513687b63f44ca038e3"></a>

## dedicated_interface — dedicated_interface / 32fc4c3dacd7 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- dedicated_interface

<a id="canonical-b775a2292fc7b58232035ca5d6c6705af10630f7da746f8ce9eee881fcc0167c"></a>

Type: `"single"`. Computed.

\[OneOf: dedicated\_interface, dedicated\_management\_interface, ethernet\_interface,
layer2\_interface, tunnel\_interface\] Configuration parameter for dedicated interface.

Upstream description:

Dedicated Interface Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-monitoring_choice": "[\"monitor\",\"monitor_disabled\"]",
  "x-ves-oneof-field-node_choice": "[\"cluster\",\"node\"]",
  "x-ves-oneof-field-primary_choice": "[\"is_primary\",\"not_primary\"]"
}
```

OneOf alternatives in this subsection:

- [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-b775a2292fc7b58232035ca5d6c6705af10630f7da746f8ce9eee881fcc0167c)
- [dedicated_management_interface](data-sources--network_interface--reference--group-001.md#canonical-d17ecc49710b63ca0b55e9db9eb444523afad5db4dcff17034553c581b87ff19)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-ed977f370a7a88690700eecaf0aa97a386c9f0713b4ba438353c228c92662644)
- [layer2_interface](data-sources--network_interface--reference--group-002.md#canonical-e25c4b23e2e1ec1252a64394f70c4aac2e7034a01010527a43b3b66f0f420ceb)
- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-2269d485ee9d6abbab5bf7aa51105026e13fedb756d7e5e7f41e68c320cf0cef)

Select alternatives according to the provider validators above.

<a id="canonical-0f03868d275740350ab90bf4366d1e86949997a285ef35049e0aa4c230574767"></a>

## Direct properties — dedicated_interface / 32fc4c3dacd7 / 3

- [cluster](data-sources--network_interface--reference--group-001.md#canonical-657bc5fee634f1922a4697fe35c988e78e7b04b9ea1ce73e642e2fa1502ce2ca): complete subsection reference.

<a id="canonical-1570e62522b5bb40f4406c6291d903656da8932eaea92e4d5251ee871d4ddffc"></a>

<a id="canonical-d881399f638912fed6aaaf21c1aca01fbd5de7e78c88ebac3a9a83036199ce16"></a>

## device property — dedicated_interface / 32fc4c3dacd7 / 4

Type: `"string"`. Computed.

Name of the device for which interface is configured. Use wwan0 for 4G/LTE.

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

- [is_primary](data-sources--network_interface--reference--group-001.md#canonical-0a746f59d172683aa5889967d708ed984a696cec6b505666fb41c990bc515bfd): complete subsection reference.

- [monitor](data-sources--network_interface--reference--group-001.md#canonical-71c0f20bbc91b46bdfb8309950bd371050f7ef14235b0028b93e3dc5b465b639): complete subsection reference.

- [monitor_disabled](data-sources--network_interface--reference--group-001.md#canonical-0cdf26626f6d5b7efe841d6e4aa006459c24a36ef906aa945c10668f8f003d30): complete subsection reference.

<a id="canonical-b90c44ec0a78b5f12b0dc003afce048b4498b349ed749e0690258cb0d18f8914"></a>

<a id="canonical-7922b9b80f0026bdbdf38276e467ef2397a08190db269fd6d39cd00ef698fcfc"></a>

## mtu property — dedicated_interface / 32fc4c3dacd7 / 5

Type: `"number"`. Computed.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

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

<a id="canonical-79c0aebd42a508357ebca3f9d0aa4febd3c0920b2f9ccb35653c7fbff3554a38"></a>

<a id="canonical-824c980385ba76d1177f43f233c163876a060a915dede8b677d77988e347c62c"></a>

## node property — dedicated_interface / 32fc4c3dacd7 / 6

Type: `"string"`. Computed.

Exclusive with \[cluster\] Configuration will apply to a device on the given node of the site.

Upstream description:

Exclusive with \[cluster\] Configuration will apply to a device on the given node of the site.

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

- [not_primary](data-sources--network_interface--reference--group-001.md#canonical-91bbb0e3e31347a1ad3574d599190ec67291289381dd13f99fe26184fc37d605): complete subsection reference.

<a id="canonical-84ec0e14e17241d66209cbe744fd01d6e94874d8eee0be0b31319aafa53977d8"></a>

<a id="canonical-85c3c3b478d974450424f2488dd1ef028353e187629ba5c9e2541e7cd64b1910"></a>

## priority property — dedicated_interface / 32fc4c3dacd7 / 7

Type: `"number"`. Computed.

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Upstream description:

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

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

<a id="canonical-fd197e58c3b2d55a83a6837e7c441f9a6a95eed5809b9ffbca4ff683acb002d0"></a>

## Next pages — dedicated_interface / 32fc4c3dacd7 / 8

- [dedicated_interface.cluster](data-sources--network_interface--reference--group-001.md#canonical-657bc5fee634f1922a4697fe35c988e78e7b04b9ea1ce73e642e2fa1502ce2ca)
- [dedicated_interface.is_primary](data-sources--network_interface--reference--group-001.md#canonical-0a746f59d172683aa5889967d708ed984a696cec6b505666fb41c990bc515bfd)
- [dedicated_interface.monitor](data-sources--network_interface--reference--group-001.md#canonical-71c0f20bbc91b46bdfb8309950bd371050f7ef14235b0028b93e3dc5b465b639)
- [dedicated_interface.monitor_disabled](data-sources--network_interface--reference--group-001.md#canonical-0cdf26626f6d5b7efe841d6e4aa006459c24a36ef906aa945c10668f8f003d30)
- [dedicated_interface.not_primary](data-sources--network_interface--reference--group-001.md#canonical-91bbb0e3e31347a1ad3574d599190ec67291289381dd13f99fe26184fc37d605)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-657bc5fee634f1922a4697fe35c988e78e7b04b9ea1ce73e642e2fa1502ce2ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4e7032ecc2bf52645594347d456bcb6c3c005bfe33359c4ae8a5fddbd167135e"></a>

## dedicated_interface.cluster — dedicated_interface.cluster / fc9465c32cc9 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-95403d00f6e6ec9eec91753b5f56180465405be1b63e1e5adc64309fbe9777b9)
- dedicated_interface.cluster

<a id="canonical-18bdafd10a866c0d01882573832000d0491665b1c8c4bc5382d8da137360355f"></a>

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

<a id="canonical-809577a2f3cbbb4196b7aa8d408f788687d3c5abb2c2ec8b5fce6b040ca597c6"></a>

## Direct properties — dedicated_interface.cluster / fc9465c32cc9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-58abca79eea194cc75493dad84e215187c8915fdae91c2265d8767b7bd5b4491"></a>

## Next pages — dedicated_interface.cluster / fc9465c32cc9 / 4

- [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-95403d00f6e6ec9eec91753b5f56180465405be1b63e1e5adc64309fbe9777b9)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-0a746f59d172683aa5889967d708ed984a696cec6b505666fb41c990bc515bfd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dfcb56faf2c2a0c9391c6a99885e0627fc6b6753f09c603349d122d54e78930e"></a>

## dedicated_interface.is_primary — dedicated_interface.is_primary / 0e1aa328d997 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-95403d00f6e6ec9eec91753b5f56180465405be1b63e1e5adc64309fbe9777b9)
- dedicated_interface.is_primary

<a id="canonical-2285b48e8622dcf92ed3d02a5acc3628de03c7339a0a9edf197e6488ee27ad8b"></a>

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

<a id="canonical-f3faba1b4863eb53547042c4643981646505e43ef02e1ad2a66f9fc56d7c1a9b"></a>

## Direct properties — dedicated_interface.is_primary / 0e1aa328d997 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6a97119fb50db8f8519e3703c000983060a12617a0331fe55351ba76154601a4"></a>

## Next pages — dedicated_interface.is_primary / 0e1aa328d997 / 4

- [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-95403d00f6e6ec9eec91753b5f56180465405be1b63e1e5adc64309fbe9777b9)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-71c0f20bbc91b46bdfb8309950bd371050f7ef14235b0028b93e3dc5b465b639"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9dd1a87bccc1851fb3812e0de6b860fdcd72c2c9cc7d4e597e19ebce6c4ca764"></a>

## dedicated_interface.monitor — dedicated_interface.monitor / 237c5c6c34cb / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-95403d00f6e6ec9eec91753b5f56180465405be1b63e1e5adc64309fbe9777b9)
- dedicated_interface.monitor

<a id="canonical-41cb2fbd437a67a7d79b4c30fb6fe8a62a78e518885fc0855d8826feda1c4e88"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-5251f93b82f7e0f94eec72ce324574c4d905a0ec33842ea6c9942f54cc677fc2"></a>

## Direct properties — dedicated_interface.monitor / 237c5c6c34cb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ee2782c764de07e863c3646043ce44b758fe3b5560f24a2695c60ef69609375e"></a>

## Next pages — dedicated_interface.monitor / 237c5c6c34cb / 4

- [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-95403d00f6e6ec9eec91753b5f56180465405be1b63e1e5adc64309fbe9777b9)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-0cdf26626f6d5b7efe841d6e4aa006459c24a36ef906aa945c10668f8f003d30"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-62876926bb3966c1a1457139e99f0d9aa5aba7dc13d130c5bdc7ad766d5da975"></a>

## dedicated_interface.monitor_disabled — dedicated_interface.monitor_disabled / a10e778bf809 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-95403d00f6e6ec9eec91753b5f56180465405be1b63e1e5adc64309fbe9777b9)
- dedicated_interface.monitor_disabled

<a id="canonical-22e311a06f46d006645f9f76af7be36ae14cc65d4f70f2bf454926fc37b3ab5d"></a>

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

<a id="canonical-92c0d224ffb33089edfacb1d04a27f7df01f6270f21252608a01755540a4fa49"></a>

## Direct properties — dedicated_interface.monitor_disabled / a10e778bf809 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e0b52cd6ad22bfc94a86ab37c47417ee4050654b68960928d6381ded9b05b012"></a>

## Next pages — dedicated_interface.monitor_disabled / a10e778bf809 / 4

- [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-95403d00f6e6ec9eec91753b5f56180465405be1b63e1e5adc64309fbe9777b9)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-91bbb0e3e31347a1ad3574d599190ec67291289381dd13f99fe26184fc37d605"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-476dda569530bf7b90f4052be13d128d9943e905b6571775c065482fbcca0fed"></a>

## dedicated_interface.not_primary — dedicated_interface.not_primary / c49557007471 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-95403d00f6e6ec9eec91753b5f56180465405be1b63e1e5adc64309fbe9777b9)
- dedicated_interface.not_primary

<a id="canonical-54c870762554406cec20c1762de3d4acc074ab2a6baf7caf1b19bef13bb83100"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for not primary.

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

<a id="canonical-75655d2bcd01332776357a89480af5ca4ab3782e79900d0b3e66572e9acdc12d"></a>

## Direct properties — dedicated_interface.not_primary / c49557007471 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-285a5f59a2543e93cc0d4d288097bd41b9483b28004426352f7d1c091101584e"></a>

## Next pages — dedicated_interface.not_primary / c49557007471 / 4

- [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-95403d00f6e6ec9eec91753b5f56180465405be1b63e1e5adc64309fbe9777b9)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-0d50dbb76d469213d5008157d7c1d5c05b12a55c0bb2bc4aa2b2a050d06fcc72"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-42b5047d0f602fd5c9ff814643506646850986e32c920ee49671c191f43e5970"></a>

## dedicated_management_interface — dedicated_management_interface / 9a26bf4ee4a0 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- dedicated_management_interface

<a id="canonical-d17ecc49710b63ca0b55e9db9eb444523afad5db4dcff17034553c581b87ff19"></a>

Type: `"single"`. Computed.

Configuration parameter for dedicated management interface.

Upstream description:

Dedicated Interface Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-node_choice": "[\"cluster\",\"node\"]"
}
```

<a id="canonical-9394b929c729be21af5b067e2802446b3a0ead79f8003ecca05d0d60827d1772"></a>

## Direct properties — dedicated_management_interface / 9a26bf4ee4a0 / 3

- [cluster](data-sources--network_interface--reference--group-001.md#canonical-59306e2906ccccc3e35988b9c3a40ddfca5a039bf6f7fa8ce6844bace980f08f): complete subsection reference.

<a id="canonical-85bbbe63de852f01c219002005b11108d3e281b9073a114d9587fe1c6090d514"></a>

<a id="canonical-38f0be0eb63e4c2ec0e3052a7ace9edeaa53c779e0d734cd5752f889e12320b0"></a>

## device property — dedicated_management_interface / 9a26bf4ee4a0 / 4

Type: `"string"`. Computed.

Name of the device for which interface is configured.

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

<a id="canonical-60e071a911ff8140320da7f0de36994f2f246dfc2bf677123d2f02bf72094f25"></a>

<a id="canonical-a76ef76f1cbfddddbf1328ba90718f3302fce9111f1bcf4d39e38704cb94e052"></a>

## mtu property — dedicated_management_interface / 9a26bf4ee4a0 / 5

Type: `"number"`. Computed.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

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

<a id="canonical-1c3f6fd643be3fee5425e541154398616324b9882537ab84f99c832d998d7456"></a>

<a id="canonical-0040af7a1c2755988ba7fe8bc445fa430935be50c2f76d0d0c3e654ee3fbab49"></a>

## node property — dedicated_management_interface / 9a26bf4ee4a0 / 6

Type: `"string"`. Computed.

Exclusive with \[cluster\] Configuration will apply to a device on the given node of the site.

Upstream description:

Exclusive with \[cluster\] Configuration will apply to a device on the given node of the site.

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

<a id="canonical-085fc50bc507aceb0fe3364fc7faecbe0f0e97305dffb6ed8aacd512abfc87ac"></a>

## Next pages — dedicated_management_interface / 9a26bf4ee4a0 / 7

- [dedicated_management_interface.cluster](data-sources--network_interface--reference--group-001.md#canonical-59306e2906ccccc3e35988b9c3a40ddfca5a039bf6f7fa8ce6844bace980f08f)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-59306e2906ccccc3e35988b9c3a40ddfca5a039bf6f7fa8ce6844bace980f08f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5dff7e6e1fde03c34227aadc82e1201bf48b544f3808dde93d0a57b6dd4018bf"></a>

## dedicated_management_interface.cluster — dedicated_management_interface.cluster / 41a3f6214848 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [dedicated_management_interface](data-sources--network_interface--reference--group-001.md#canonical-0d50dbb76d469213d5008157d7c1d5c05b12a55c0bb2bc4aa2b2a050d06fcc72)
- dedicated_management_interface.cluster

<a id="canonical-33b197e238e944da8ab05edb4edf6f21cac78c29dc785cf2dee105c7c2d9696e"></a>

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

<a id="canonical-dd6bf5960f2fd02fe70ffd57186f71b65bb1701d9aabec95f58bf3477a7ee416"></a>

## Direct properties — dedicated_management_interface.cluster / 41a3f6214848 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6b11086d2bb545d712e40e5a0bfbebca418b60c0f8021eda20ba5b1bbdc9ff7f"></a>

## Next pages — dedicated_management_interface.cluster / 41a3f6214848 / 4

- [dedicated_management_interface](data-sources--network_interface--reference--group-001.md#canonical-0d50dbb76d469213d5008157d7c1d5c05b12a55c0bb2bc4aa2b2a050d06fcc72)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-69bbe1bbefceb8f621308fe53ac22f0179faef94824c5adfc9f1f15ff994d637"></a>

## ethernet_interface — ethernet_interface / d4ca96a5eef2 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- ethernet_interface

<a id="canonical-ed977f370a7a88690700eecaf0aa97a386c9f0713b4ba438353c228c92662644"></a>

Type: `"single"`. Computed.

Configuration parameter for ethernet interface.

Upstream description:

Ethernet Interface Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-address_choice": "[\"dhcp_client\",\"dhcp_server\",\"static_ip\"]",
  "x-ves-oneof-field-ipv6_address_choice": "[\"ipv6_auto_config\",\"no_ipv6_address\",\"static_ipv6_address\"]",
  "x-ves-oneof-field-monitoring_choice": "[\"monitor\",\"monitor_disabled\"]",
  "x-ves-oneof-field-network_choice": "[\"segment_network\",\"site_local_inside_network\",\"site_local_network\",\"storage_network\"]",
  "x-ves-oneof-field-node_choice": "[\"cluster\",\"node\"]",
  "x-ves-oneof-field-primary_choice": "[\"is_primary\",\"not_primary\"]",
  "x-ves-oneof-field-vlan_choice": "[\"untagged\",\"vlan_id\"]"
}
```

<a id="canonical-a5db739fdd0d529c603f582f88af35eecbc38d2070b5b9b258d09916c2d6bd6a"></a>

## Direct properties — ethernet_interface / d4ca96a5eef2 / 3

- [cluster](data-sources--network_interface--reference--group-001.md#canonical-40d2ec4a08f043190a49f519bb6b45850a439b964993b597a51aef00d90bf9ce): complete subsection reference.

<a id="canonical-515937a1d7ff2e6ec78a8f95c89397c9ff70f94f3881861687170c15eff22e42"></a>

<a id="canonical-00dca3928a64f61891a9c27a14e45133e5a7ba0b4f7bf3e4abb649d7319e47db"></a>

## device property — ethernet_interface / d4ca96a5eef2 / 4

Type: `"string"`. Computed.

Interface configuration for the ethernet device.

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

- [dhcp_client](data-sources--network_interface--reference--group-001.md#canonical-958f4ff1237e1d1af2d528e90f3c89fdf9f69ccd4ae2d4e90c954055d2b8ee2b): complete subsection reference.

- [dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-a1c03784815f7c28f4c5371b2b908f57ea6660b142df69b6bb36c2aa689d900f): complete subsection reference.

- [ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-7b12eb5b085cdc770291316c0876e0a8a6e9a5e76abde116ce768c70707c4b41): complete subsection reference.

- [is_primary](data-sources--network_interface--reference--group-001.md#canonical-1de8a146ae46e90e592cf7bc92684f0a33166b790c639ec602fab32bbd7a33e1): complete subsection reference.

- [monitor](data-sources--network_interface--reference--group-001.md#canonical-7e59e10a823eaead070dcef8c323e667a853846698f1bd876926d367a56f9521): complete subsection reference.

- [monitor_disabled](data-sources--network_interface--reference--group-001.md#canonical-d9f070c79019cfaa5f37f86a89e79b9582c43664ba57efac2914032d8bbe3397): complete subsection reference.

<a id="canonical-ca2651856c24606f1d5de70742025141929261bb489dd34f4090ac5f4eeab763"></a>

<a id="canonical-a1f94d45df309ca991be6f6631c1785a6cba6d3fa3da389a20461c41a62d7fa4"></a>

## mtu property — ethernet_interface / d4ca96a5eef2 / 5

Type: `"number"`. Computed.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

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

- [no_ipv6_address](data-sources--network_interface--reference--group-001.md#canonical-573b9e41b481f04ee78288ebc0a173626ca7680662d4131a3974064055dfb395): complete subsection reference.

<a id="canonical-3226a284b5cdfee4636cc54430f7d4e154054cbd4f69fd0ac9d349dae5504bbf"></a>

<a id="canonical-2b8df98709d8ce92683a77bdac285a70bcec2d0138c58c40d49ee19c6b0ec8d5"></a>

## node property — ethernet_interface / d4ca96a5eef2 / 6

Type: `"string"`. Computed.

Exclusive with \[cluster\] Configuration will apply to a device on the given node.

Upstream description:

Exclusive with \[cluster\] Configuration will apply to a device on the given node.

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

- [not_primary](data-sources--network_interface--reference--group-001.md#canonical-f6b72b300a33a5d0a16dd78cd0144638f4b9024d74cd353a4d1ea773dc4b2dfc): complete subsection reference.

<a id="canonical-7001aed971c33b5e31055b05464905d07d805a07fd8facfd759248dea27ab0f4"></a>

<a id="canonical-6d2a9bc8e3f40d93d5bc8fca967f64ace2993f7c47708a2aadc658e24e26fa63"></a>

## priority property — ethernet_interface / d4ca96a5eef2 / 7

Type: `"number"`. Computed.

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Upstream description:

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

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

- [site_local_inside_network](data-sources--network_interface--reference--group-001.md#canonical-26db361e57025a942717ab5ae7b3ed631fc466b9d2ab1b08aece2bfbf6740c04): complete subsection reference.

- [site_local_network](data-sources--network_interface--reference--group-001.md#canonical-de90c1d6529c71ee3906c0de7866ef0b3558698a9733c41930f7f064cf3dc849): complete subsection reference.

- [static_ip](data-sources--network_interface--reference--group-001.md#canonical-f6a7a273d85d7cfa1c8614e9eecb3589eb3848255af051b6a00adcd274fd76ed): complete subsection reference.

- [static_ipv6_address](data-sources--network_interface--reference--group-002.md#canonical-9a76ad147e80b2c51535b56cd403c93b9e85bbc9aa57014fc100c9bb299ff1f4): complete subsection reference.

- [storage_network](data-sources--network_interface--reference--group-002.md#canonical-3f3e41399e956ac76c0e4c72fc7f2a6417f8c1cc7de2f968c11e0059a3d8411a): complete subsection reference.

- [untagged](data-sources--network_interface--reference--group-002.md#canonical-b3402af0a5d18205df42bc79bbeb04d66cf0db5fec4c9f1f4bea5d4601a46b76): complete subsection reference.

<a id="canonical-c72d0e2d5a111107a6f35cb3ce8f6a45841211243c03d5c6d6ad7c75f3ca6ed1"></a>

<a id="canonical-cb823ef9762c95997e0b47762dfafc96e695080cfad58f786bf539abc5269ee1"></a>

## vlan_id property — ethernet_interface / d4ca96a5eef2 / 8

Type: `"number"`. Computed.

Exclusive with \[untagged\] Configure a VLAN tagged ethernet interface.

Upstream description:

Exclusive with \[untagged\] Configure a VLAN tagged ethernet interface.

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

<a id="canonical-3598e1b808607e745c0809da04a3196324008d93c500ad3e3c2ec2221c4a55c7"></a>

## Next pages — ethernet_interface / d4ca96a5eef2 / 9

- [ethernet_interface.cluster](data-sources--network_interface--reference--group-001.md#canonical-40d2ec4a08f043190a49f519bb6b45850a439b964993b597a51aef00d90bf9ce)
- [ethernet_interface.dhcp_client](data-sources--network_interface--reference--group-001.md#canonical-958f4ff1237e1d1af2d528e90f3c89fdf9f69ccd4ae2d4e90c954055d2b8ee2b)
- [ethernet_interface.dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-a1c03784815f7c28f4c5371b2b908f57ea6660b142df69b6bb36c2aa689d900f)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-7b12eb5b085cdc770291316c0876e0a8a6e9a5e76abde116ce768c70707c4b41)
- [ethernet_interface.is_primary](data-sources--network_interface--reference--group-001.md#canonical-1de8a146ae46e90e592cf7bc92684f0a33166b790c639ec602fab32bbd7a33e1)
- [ethernet_interface.monitor](data-sources--network_interface--reference--group-001.md#canonical-7e59e10a823eaead070dcef8c323e667a853846698f1bd876926d367a56f9521)
- [ethernet_interface.monitor_disabled](data-sources--network_interface--reference--group-001.md#canonical-d9f070c79019cfaa5f37f86a89e79b9582c43664ba57efac2914032d8bbe3397)
- [ethernet_interface.no_ipv6_address](data-sources--network_interface--reference--group-001.md#canonical-573b9e41b481f04ee78288ebc0a173626ca7680662d4131a3974064055dfb395)
- [ethernet_interface.not_primary](data-sources--network_interface--reference--group-001.md#canonical-f6b72b300a33a5d0a16dd78cd0144638f4b9024d74cd353a4d1ea773dc4b2dfc)
- [ethernet_interface.site_local_inside_network](data-sources--network_interface--reference--group-001.md#canonical-26db361e57025a942717ab5ae7b3ed631fc466b9d2ab1b08aece2bfbf6740c04)
- [ethernet_interface.site_local_network](data-sources--network_interface--reference--group-001.md#canonical-de90c1d6529c71ee3906c0de7866ef0b3558698a9733c41930f7f064cf3dc849)
- [ethernet_interface.static_ip](data-sources--network_interface--reference--group-001.md#canonical-f6a7a273d85d7cfa1c8614e9eecb3589eb3848255af051b6a00adcd274fd76ed)
- [ethernet_interface.static_ipv6_address](data-sources--network_interface--reference--group-002.md#canonical-9a76ad147e80b2c51535b56cd403c93b9e85bbc9aa57014fc100c9bb299ff1f4)
- [ethernet_interface.storage_network](data-sources--network_interface--reference--group-002.md#canonical-3f3e41399e956ac76c0e4c72fc7f2a6417f8c1cc7de2f968c11e0059a3d8411a)
- [ethernet_interface.untagged](data-sources--network_interface--reference--group-002.md#canonical-b3402af0a5d18205df42bc79bbeb04d66cf0db5fec4c9f1f4bea5d4601a46b76)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-40d2ec4a08f043190a49f519bb6b45850a439b964993b597a51aef00d90bf9ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-92f1b19314c376ceea44ccc23fed2b5310ec12ecec10ad42d8dbef33c224f9e2"></a>

## ethernet_interface.cluster — ethernet_interface.cluster / 4d96e9882140 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- ethernet_interface.cluster

<a id="canonical-201fea49c35ca3602d9b075febafab6838f60d7cfd4b181029ec593ad694ee22"></a>

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

<a id="canonical-9217c9a900f61037dd83f7a517407b5ffeba1c537aab4d05e5cb03fd23a0d9e1"></a>

## Direct properties — ethernet_interface.cluster / 4d96e9882140 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1c713f061036d57c36ee4ab3bb423d56f7fec02da5a9da42f999d9c3dc1c1168"></a>

## Next pages — ethernet_interface.cluster / 4d96e9882140 / 4

- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-958f4ff1237e1d1af2d528e90f3c89fdf9f69ccd4ae2d4e90c954055d2b8ee2b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-970b9df061dae4eaac16776df40491670b272ebb21569f9a60f67d823d731c3a"></a>

## ethernet_interface.dhcp_client — ethernet_interface.dhcp_client / 5399d914ee13 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- ethernet_interface.dhcp_client

<a id="canonical-ba3b1c0472e207ddf21e7d9c971a4e358d3138e277dcb3bdcd382f2b58b8f5c2"></a>

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

<a id="canonical-022e5cfbc4af7f8d757748927c21dda6ac0a7d2338646893838f15f3906d80b1"></a>

## Direct properties — ethernet_interface.dhcp_client / 5399d914ee13 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-672741a55b181d67014c8b9f9dbe10f210faf988dbea0f63a3748c7a27729ff8"></a>

## Next pages — ethernet_interface.dhcp_client / 5399d914ee13 / 4

- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-a1c03784815f7c28f4c5371b2b908f57ea6660b142df69b6bb36c2aa689d900f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cb80aa0cbe235659597aa3103476408880127d3be6c94703809a6f015e91a9ad"></a>

## ethernet_interface.dhcp_server — ethernet_interface.dhcp_server / 648e57d20186 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- ethernet_interface.dhcp_server

<a id="canonical-210b4c2ee1490b37475d97f74261dbe4a0443e1e376b29c3ac077ba47a18ac9d"></a>

Type: `"single"`. Computed.

Configuration parameter for dhcp server.

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

<a id="canonical-2631763a5a9e255626ba8367e2a9bdf336cb5cdaea2a54a686b958e9282f005a"></a>

## Direct properties — ethernet_interface.dhcp_server / 648e57d20186 / 3

- [automatic_from_end](data-sources--network_interface--reference--group-001.md#canonical-1aa1629054811f74d840e23dc0ccf897ffe12bf82c66e709e0e0e40b5e324a35): complete subsection reference.

- [automatic_from_start](data-sources--network_interface--reference--group-001.md#canonical-6779fd15c646e2886cf3975f29a670a69420d602c36c56e647ea63ea943adede): complete subsection reference.

- [dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-02a1ddae55b02a2bbdc16d52d9e0954f902f591b7f3c000560127eb8190bbe99): complete subsection reference.

<a id="canonical-a0a138aaa34162ee8b1096fe5adba1537346ed4b9fe6cc662123b10bba60b6e8"></a>

<a id="canonical-b582d148902a506f583e326958bd7c6caac223f6ec11bde70193c3b397d6557a"></a>

## dhcp_option82_tag property — ethernet_interface.dhcp_server / 648e57d20186 / 4

Type: `"string"`. Computed.

DHCP option 82 tag.

<a id="canonical-033991a4fa71b049d0be2af4f224520d769c4cb091816a62d95c2e86fdd963c6"></a>

<a id="canonical-179adb377c1c7f695d8cd50f04a40fed63472fc0c27a4d3f8cf0c8dad9a1d8a3"></a>

## fixed_ip_map property — ethernet_interface.dhcp_server / 648e57d20186 / 5

Type: `["map", "string"]`. Computed.

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

- [interface_ip_map](data-sources--network_interface--reference--group-001.md#canonical-123736dfe43ed355fe07480e3266061f2d54291c09b89d1531b96027c34939ae): complete subsection reference.

<a id="canonical-8829081e8f666608db0b59b36076a9642ca250764d58333c51ef55de42cf3d07"></a>

## Next pages — ethernet_interface.dhcp_server / 648e57d20186 / 6

- [ethernet_interface.dhcp_server.automatic_from_end](data-sources--network_interface--reference--group-001.md#canonical-1aa1629054811f74d840e23dc0ccf897ffe12bf82c66e709e0e0e40b5e324a35)
- [ethernet_interface.dhcp_server.automatic_from_start](data-sources--network_interface--reference--group-001.md#canonical-6779fd15c646e2886cf3975f29a670a69420d602c36c56e647ea63ea943adede)
- [ethernet_interface.dhcp_server.dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-02a1ddae55b02a2bbdc16d52d9e0954f902f591b7f3c000560127eb8190bbe99)
- [ethernet_interface.dhcp_server.interface_ip_map](data-sources--network_interface--reference--group-001.md#canonical-123736dfe43ed355fe07480e3266061f2d54291c09b89d1531b96027c34939ae)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-1aa1629054811f74d840e23dc0ccf897ffe12bf82c66e709e0e0e40b5e324a35"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-94a9ded205469a9ddc0fa138de354d28473aafb47118dc8e050a6c0e0cd87485"></a>

## ethernet_interface.dhcp_server.automatic_from_end — ethernet_interface.dhcp_server.automatic_from_end / d59c07682ac3 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [ethernet_interface.dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-a1c03784815f7c28f4c5371b2b908f57ea6660b142df69b6bb36c2aa689d900f)
- ethernet_interface.dhcp_server.automatic_from_end

<a id="canonical-f870a0623ad6f1a2e5ab9cec792b7fcb8f40af81bdc68fa9deb3b2c9c1d63295"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-77f3829db434649a2d68a2ebd724654dcb9a42b4775071e2bf46cfabd0b13d9e"></a>

## Direct properties — ethernet_interface.dhcp_server.automatic_from_end / d59c07682ac3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9ce0b4062c494b4bfd13363de0244280fd53723826c18841b34083747acd0854"></a>

## Next pages — ethernet_interface.dhcp_server.automatic_from_end / d59c07682ac3 / 4

- [ethernet_interface.dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-a1c03784815f7c28f4c5371b2b908f57ea6660b142df69b6bb36c2aa689d900f)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-6779fd15c646e2886cf3975f29a670a69420d602c36c56e647ea63ea943adede"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63e23c88d8a19a87089773a24c62115c16d1c324806fea291078e94619125f7d"></a>

## ethernet_interface.dhcp_server.automatic_from_start — ethernet_interface.dhcp_server.automatic_from_start / 18ed1b6bf589 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [ethernet_interface.dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-a1c03784815f7c28f4c5371b2b908f57ea6660b142df69b6bb36c2aa689d900f)
- ethernet_interface.dhcp_server.automatic_from_start

<a id="canonical-f7ee4e104f435a975774740af7a599d2eba5692221bd3c9d2716fc63f5a872ae"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-f8711a7b23efd58fd740969ff05d1474b850c05f94d9a9f8d53428994af8ea97"></a>

## Direct properties — ethernet_interface.dhcp_server.automatic_from_start / 18ed1b6bf589 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4faee2a339db2fff5951d38dfad2b31a8aefdec9a8436fab1c14bd2445e16bf3"></a>

## Next pages — ethernet_interface.dhcp_server.automatic_from_start / 18ed1b6bf589 / 4

- [ethernet_interface.dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-a1c03784815f7c28f4c5371b2b908f57ea6660b142df69b6bb36c2aa689d900f)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-02a1ddae55b02a2bbdc16d52d9e0954f902f591b7f3c000560127eb8190bbe99"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca491f388f755b9e75671d5d0876daf2a212d5b9a910d2e5e8e3ba472c56001f"></a>

## ethernet_interface.dhcp_server.dhcp_networks — ethernet_interface.dhcp_server.dhcp_networks / bf97488c064d / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [ethernet_interface.dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-a1c03784815f7c28f4c5371b2b908f57ea6660b142df69b6bb36c2aa689d900f)
- ethernet_interface.dhcp_server.dhcp_networks

<a id="canonical-b23646f331a3fac94d2281e9410566c6ae09adba9a4a0b8afdaf4316f8e76649"></a>

Type: `"list"`. Computed.

List of networks from which DHCP Server can allocate IPv4 Addresses.

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

<a id="canonical-d35d0d394b362a3f98cfe6884fb356df102417c9c0dc1669b3e4dd195ec4b6f5"></a>

## Direct properties — ethernet_interface.dhcp_server.dhcp_networks / bf97488c064d / 3

<a id="canonical-3c3f11a9e90ea394084c24faae70a3f76dc97228d218a6aff29d91826b697511"></a>

<a id="canonical-6331ee63f75ef81be170b1f44ce5e48ccbec908d13f2282d94b4d5fea80dd0cb"></a>

## dgw_address property — ethernet_interface.dhcp_server.dhcp_networks / bf97488c064d / 4

Type: `"string"`. Computed.

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Upstream description:

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

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

<a id="canonical-4991ef0c385c11b86862dc52128f6266ca15e003becd3a3df1ddee5956f40889"></a>

<a id="canonical-d5e8e4909ad86b4191c29570508474a8447e302fc68c4e5469b0f4d28dabd0ae"></a>

## dns_address property — ethernet_interface.dhcp_server.dhcp_networks / bf97488c064d / 5

Type: `"string"`. Computed.

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Upstream description:

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

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

- [first_address](data-sources--network_interface--reference--group-001.md#canonical-dbd1172db8458be3a80f3c9167c8da2bdb17dc731db1d899f8a0663d2d3761e9): complete subsection reference.

- [last_address](data-sources--network_interface--reference--group-001.md#canonical-8332d8ec40e2f98b2042eb93016110f3cc2eb8604bb30a5e7565c5bae80fcd0a): complete subsection reference.

<a id="canonical-81e4132f1a0fa34120636ae018794f07f1bda1aa44c8fb5890963c00bb0bc144"></a>

<a id="canonical-1f059d4569fc2918875f5b8d3eb923474a4f0e985491d44aaf996541b8127d69"></a>

## network_prefix property — ethernet_interface.dhcp_server.dhcp_networks / bf97488c064d / 6

Type: `"string"`. Computed.

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

<a id="canonical-c0f86457340f34560b8ee841d1bb3fb0e1584a0b084d897da9b9f917f24bf26c"></a>

<a id="canonical-cab78d4cdfcdf1e7063b75f16531b21204d77a46e724cb73d0a4d634e9703573"></a>

## pool_settings property — ethernet_interface.dhcp_server.dhcp_networks / bf97488c064d / 7

Type: `"string"`. Computed.

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

- [pools](data-sources--network_interface--reference--group-001.md#canonical-72c61d4278c7a5de804138a7a4cf2e3dbbbe444101f891f53b60fd0730d7c87e): complete subsection reference.

- [same_as_dgw](data-sources--network_interface--reference--group-001.md#canonical-363e8ae78d7b0c38942992fb57a30aad94ee41a561cc248fb6859c01ce37185d): complete subsection reference.

<a id="canonical-07debffc4c444dc6d7bf5fb31cf73f772948cf44ee0d14b1ddc7cdb2c13d52ed"></a>

## Next pages — ethernet_interface.dhcp_server.dhcp_networks / bf97488c064d / 8

- [ethernet_interface.dhcp_server.dhcp_networks.first_address](data-sources--network_interface--reference--group-001.md#canonical-dbd1172db8458be3a80f3c9167c8da2bdb17dc731db1d899f8a0663d2d3761e9)
- [ethernet_interface.dhcp_server.dhcp_networks.last_address](data-sources--network_interface--reference--group-001.md#canonical-8332d8ec40e2f98b2042eb93016110f3cc2eb8604bb30a5e7565c5bae80fcd0a)
- [ethernet_interface.dhcp_server.dhcp_networks.pools](data-sources--network_interface--reference--group-001.md#canonical-72c61d4278c7a5de804138a7a4cf2e3dbbbe444101f891f53b60fd0730d7c87e)
- [ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw](data-sources--network_interface--reference--group-001.md#canonical-363e8ae78d7b0c38942992fb57a30aad94ee41a561cc248fb6859c01ce37185d)
- [ethernet_interface.dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-a1c03784815f7c28f4c5371b2b908f57ea6660b142df69b6bb36c2aa689d900f)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-dbd1172db8458be3a80f3c9167c8da2bdb17dc731db1d899f8a0663d2d3761e9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-efdbba9551098e223932910879256be5881258e8cc5300e93e110f2ca0758b8b"></a>

## ethernet_interface.dhcp_server.dhcp_networks.first_address — ethernet_interface.dhcp_server.dhcp_networks.first_address / bc5a7881b4af / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [ethernet_interface.dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-a1c03784815f7c28f4c5371b2b908f57ea6660b142df69b6bb36c2aa689d900f)
- [ethernet_interface.dhcp_server.dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-02a1ddae55b02a2bbdc16d52d9e0954f902f591b7f3c000560127eb8190bbe99)
- ethernet_interface.dhcp_server.dhcp_networks.first_address

<a id="canonical-20789423749358083c255998675923223c280cffd71f654eef3a09a130ae32d1"></a>

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

<a id="canonical-b57107297fb2ed6f1e2793fb0c7c2d99a129445c8b2dc57d331a0f33028aa4bd"></a>

## Direct properties — ethernet_interface.dhcp_server.dhcp_networks.first_address / bc5a7881b4af / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2483d6cad73f5a0e450ca0248557fb7c0451997fea6da8778e6b671f486a3d2e"></a>

## Next pages — ethernet_interface.dhcp_server.dhcp_networks.first_address / bc5a7881b4af / 4

- [ethernet_interface.dhcp_server.dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-02a1ddae55b02a2bbdc16d52d9e0954f902f591b7f3c000560127eb8190bbe99)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-8332d8ec40e2f98b2042eb93016110f3cc2eb8604bb30a5e7565c5bae80fcd0a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d0d827bbd9653a3f87223eb527e55a5989c18958c149fc1a314fed0fe7437b9c"></a>

## ethernet_interface.dhcp_server.dhcp_networks.last_address — ethernet_interface.dhcp_server.dhcp_networks.last_address / 2bd4b87e6007 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [ethernet_interface.dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-a1c03784815f7c28f4c5371b2b908f57ea6660b142df69b6bb36c2aa689d900f)
- [ethernet_interface.dhcp_server.dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-02a1ddae55b02a2bbdc16d52d9e0954f902f591b7f3c000560127eb8190bbe99)
- ethernet_interface.dhcp_server.dhcp_networks.last_address

<a id="canonical-2c6446688cb3d6cee12518a2d5985d18037a1deeef3842a5e667670cb541c2c0"></a>

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

<a id="canonical-1d572b70c3540940bbcdd74d09f5fff2608b4f03de3eb3131449e80e3f8e3db4"></a>

## Direct properties — ethernet_interface.dhcp_server.dhcp_networks.last_address / 2bd4b87e6007 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-99f9b6fa4e860d236edcff94e0a24a5610aefe193aef32feaf2b618487d53df1"></a>

## Next pages — ethernet_interface.dhcp_server.dhcp_networks.last_address / 2bd4b87e6007 / 4

- [ethernet_interface.dhcp_server.dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-02a1ddae55b02a2bbdc16d52d9e0954f902f591b7f3c000560127eb8190bbe99)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-72c61d4278c7a5de804138a7a4cf2e3dbbbe444101f891f53b60fd0730d7c87e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a00594b487d77b61249cf3a023a2789ed746ad75681c28ac900b4378df30f992"></a>

## ethernet_interface.dhcp_server.dhcp_networks.pools — ethernet_interface.dhcp_server.dhcp_networks.pools / 19f0d83db2ea / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [ethernet_interface.dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-a1c03784815f7c28f4c5371b2b908f57ea6660b142df69b6bb36c2aa689d900f)
- [ethernet_interface.dhcp_server.dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-02a1ddae55b02a2bbdc16d52d9e0954f902f591b7f3c000560127eb8190bbe99)
- ethernet_interface.dhcp_server.dhcp_networks.pools

<a id="canonical-850f6806611f42df587b791194b83a05e4ac988523043a06187d6cf672c20941"></a>

Type: `"list"`. Computed.

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

<a id="canonical-b17c855049ae2f924b067cfaeedd2b228e754f99add3d40c915e53141df7af05"></a>

## Direct properties — ethernet_interface.dhcp_server.dhcp_networks.pools / 19f0d83db2ea / 3

<a id="canonical-904673383562ab7085d12ffa1ea64f113f8379b9e3861ebdd05b6c6f84810c41"></a>

<a id="canonical-69dca0ec912bd978aa4916fcf3bb97e838ffdabace58cd00a5cc3ca2c20feda8"></a>

## end_ip property — ethernet_interface.dhcp_server.dhcp_networks.pools / 19f0d83db2ea / 4

Type: `"string"`. Computed.

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

Upstream description:

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

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

<a id="canonical-252dfe698fb21347a42f99af9f6bb1402e55f8665fb5a694882f8c8d594c799e"></a>

<a id="canonical-3e3655f23a7c5c0c85d66ca42ab73e9a7f513b9d298321a3d9c34e3ec95faf1d"></a>

## exclude property — ethernet_interface.dhcp_server.dhcp_networks.pools / 19f0d83db2ea / 5

Type: `"bool"`. Computed.

Exclude this address range from DHCP allocation.

<a id="canonical-bcac4a5e8c1b7093d3394f0c606c6a6a14cd60aab9f28df04d3b727ddd2704d3"></a>

<a id="canonical-64ff678a352f82b185df13f4798b670b613cb56079045602df6b6e03a2b8baf0"></a>

## start_ip property — ethernet_interface.dhcp_server.dhcp_networks.pools / 19f0d83db2ea / 6

Type: `"string"`. Computed.

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

Upstream description:

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

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

<a id="canonical-1422fb4da8fdfcc90cb9210ab7904c729fc0a9f0b39e2beabd21820395d8a1d3"></a>

## Next pages — ethernet_interface.dhcp_server.dhcp_networks.pools / 19f0d83db2ea / 7

- [ethernet_interface.dhcp_server.dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-02a1ddae55b02a2bbdc16d52d9e0954f902f591b7f3c000560127eb8190bbe99)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-363e8ae78d7b0c38942992fb57a30aad94ee41a561cc248fb6859c01ce37185d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e9ed0f2deb07c651ff19cd681b02fe503d5847d2f383ef946647891dbe390a36"></a>

## ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw — ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw / f4d39352f591 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [ethernet_interface.dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-a1c03784815f7c28f4c5371b2b908f57ea6660b142df69b6bb36c2aa689d900f)
- [ethernet_interface.dhcp_server.dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-02a1ddae55b02a2bbdc16d52d9e0954f902f591b7f3c000560127eb8190bbe99)
- ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-006145690a9f64a2487279ce5d889a3ec199f23d4abcd530213dcbe213e00bd1"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for same as dgw.

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

<a id="canonical-532382c6811ca0af3647b15656cdeb50f69f62f253a024b41cbb12c8c05c79ca"></a>

## Direct properties — ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw / f4d39352f591 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5b77c8e84f4cf676e714a11f0c3202b72079cfef4bbb9ebd98ae3f0a981c59df"></a>

## Next pages — ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw / f4d39352f591 / 4

- [ethernet_interface.dhcp_server.dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-02a1ddae55b02a2bbdc16d52d9e0954f902f591b7f3c000560127eb8190bbe99)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-123736dfe43ed355fe07480e3266061f2d54291c09b89d1531b96027c34939ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fe085fa0b3ecebc1d6c9abdde044e3dfc4646ba42a5e7656572b63a7c52e996b"></a>

## ethernet_interface.dhcp_server.interface_ip_map — ethernet_interface.dhcp_server.interface_ip_map / 2c3e60e8720c / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [ethernet_interface.dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-a1c03784815f7c28f4c5371b2b908f57ea6660b142df69b6bb36c2aa689d900f)
- ethernet_interface.dhcp_server.interface_ip_map

<a id="canonical-f6891f1d0aad9a25c52fc2448db5308b903b19c5cb46da65200b06c314b48f93"></a>

Type: `"single"`. Computed.

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

<a id="canonical-f3d360e52ea244266f4cdc7f6c004e943f4f4573e511af42aaebb175820f7c35"></a>

## Direct properties — ethernet_interface.dhcp_server.interface_ip_map / 2c3e60e8720c / 3

<a id="canonical-0835d94dc9f73c10105cad94a1d049f25ed1557565b74d2dba4b0893b22fb6b4"></a>

<a id="canonical-99414cc92c72c83ced6ed477dea65901449c6eeeb4b24859a425773ca1f08d15"></a>

## interface_ip_map property — ethernet_interface.dhcp_server.interface_ip_map / 2c3e60e8720c / 4

Type: `["map", "string"]`. Computed.

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

<a id="canonical-2c37a88753ff82935fdd188d4e51f73ac337e17e97e4a37e15738795043b4174"></a>

## Next pages — ethernet_interface.dhcp_server.interface_ip_map / 2c3e60e8720c / 5

- [ethernet_interface.dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-a1c03784815f7c28f4c5371b2b908f57ea6660b142df69b6bb36c2aa689d900f)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-7b12eb5b085cdc770291316c0876e0a8a6e9a5e76abde116ce768c70707c4b41"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-36db9a81c2ca6ebe82c94356e9111a762a81be3ebd05b555bd08cb253c00612d"></a>

## ethernet_interface.ipv6_auto_config — ethernet_interface.ipv6_auto_config / 24a3bd228686 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- ethernet_interface.ipv6_auto_config

<a id="canonical-32224430e9c6613018c1114103f99d733cea2983351d2724f994036cfbcb7887"></a>

Type: `"single"`. Computed.

IPV6AutoConfigType.

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

<a id="canonical-c8b336b394110fbf7ea2f63a9698f824efbd0708c0af312f799f4f784bad008b"></a>

## Direct properties — ethernet_interface.ipv6_auto_config / 24a3bd228686 / 3

- [host](data-sources--network_interface--reference--group-001.md#canonical-ddab4295a6b515a0e96bcf81400f4a47c35d22e50cf40075d64770b7e1cb47dc): complete subsection reference.

- [router](data-sources--network_interface--reference--group-001.md#canonical-09992a12de87aecff602762fe48959ff2c00c76dfeae4f4028a8b2ccf71518a1): complete subsection reference.

<a id="canonical-642ca539b2d9d65ec2cd46bc74ae3a38321cda8645fe6eacd5eac3e42cc3d8bd"></a>

## Next pages — ethernet_interface.ipv6_auto_config / 24a3bd228686 / 4

- [ethernet_interface.ipv6_auto_config.host](data-sources--network_interface--reference--group-001.md#canonical-ddab4295a6b515a0e96bcf81400f4a47c35d22e50cf40075d64770b7e1cb47dc)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-09992a12de87aecff602762fe48959ff2c00c76dfeae4f4028a8b2ccf71518a1)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-ddab4295a6b515a0e96bcf81400f4a47c35d22e50cf40075d64770b7e1cb47dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a15513eea58c83861b36ce8a79f81b38f86ebd6fd92fa35e58180fe18014eafe"></a>

## ethernet_interface.ipv6_auto_config.host — ethernet_interface.ipv6_auto_config.host / 911d33094de8 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-7b12eb5b085cdc770291316c0876e0a8a6e9a5e76abde116ce768c70707c4b41)
- ethernet_interface.ipv6_auto_config.host

<a id="canonical-fda544d3e19b3fac3f12b92213c9e9a400014e95bb43741194d936ddd5f3fef2"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0efeb695ea07aa4ff626c3fcc2ea41e14d7b6d654c3082964d20b8354d802a95"></a>

## Direct properties — ethernet_interface.ipv6_auto_config.host / 911d33094de8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7959b79bceee9317c6d484d4922d548418a1200d3ce8d27004e61e3ac5edb913"></a>

## Next pages — ethernet_interface.ipv6_auto_config.host / 911d33094de8 / 4

- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-7b12eb5b085cdc770291316c0876e0a8a6e9a5e76abde116ce768c70707c4b41)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-09992a12de87aecff602762fe48959ff2c00c76dfeae4f4028a8b2ccf71518a1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a2004f5cee4d374a79c24a0b64bc6e64caa61d6a65c07f2a2b81d48075fe72a9"></a>

## ethernet_interface.ipv6_auto_config.router — ethernet_interface.ipv6_auto_config.router / 5b49189f3f6a / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-7b12eb5b085cdc770291316c0876e0a8a6e9a5e76abde116ce768c70707c4b41)
- ethernet_interface.ipv6_auto_config.router

<a id="canonical-e13ef8cbfa67809aa7dd7755779911e3b1d103dbc6926042520f6eb2059c9150"></a>

Type: `"single"`. Computed.

IPV6AutoConfigRouterType.

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

<a id="canonical-11c2fce272a3cc7caebf492285ef06e05bf362a11707224d50211b65cd847dfc"></a>

## Direct properties — ethernet_interface.ipv6_auto_config.router / 5b49189f3f6a / 3

- [dns_config](data-sources--network_interface--reference--group-001.md#canonical-5bc25a772987541c959c266e81783f689c4e20795918d7a2188ca6cd0f70eccf): complete subsection reference.

<a id="canonical-0a05f6edec7fb66ad7cda39be06835d4ad3e36eb08e940bcb0d544e319cc6b76"></a>

<a id="canonical-d274f47f3f9a126879dd538caefb6d847d1f0382e2f3f916808b7e0ab6f95f57"></a>

## network_prefix property — ethernet_interface.ipv6_auto_config.router / 5b49189f3f6a / 4

Type: `"string"`. Computed.

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Upstream description:

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

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

- [stateful](data-sources--network_interface--reference--group-001.md#canonical-94c2be9d9397047ff932739152da71df44728a298cfb39a7dec7a79d2135b808): complete subsection reference.

<a id="canonical-5610eee46f0483b8d32e4bf4ab49566cc07f545dedb1e75619bcc0b390173655"></a>

## Next pages — ethernet_interface.ipv6_auto_config.router / 5b49189f3f6a / 5

- [ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--network_interface--reference--group-001.md#canonical-5bc25a772987541c959c266e81783f689c4e20795918d7a2188ca6cd0f70eccf)
- [ethernet_interface.ipv6_auto_config.router.stateful](data-sources--network_interface--reference--group-001.md#canonical-94c2be9d9397047ff932739152da71df44728a298cfb39a7dec7a79d2135b808)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-7b12eb5b085cdc770291316c0876e0a8a6e9a5e76abde116ce768c70707c4b41)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-5bc25a772987541c959c266e81783f689c4e20795918d7a2188ca6cd0f70eccf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-806dd290b475bfc5eb2ede0d6130c87c1555e644c49246e0a4ae62fc844fd3f8"></a>

## ethernet_interface.ipv6_auto_config.router.dns_config — ethernet_interface.ipv6_auto_config.router.dns_config / a238e8fefd63 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-7b12eb5b085cdc770291316c0876e0a8a6e9a5e76abde116ce768c70707c4b41)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-09992a12de87aecff602762fe48959ff2c00c76dfeae4f4028a8b2ccf71518a1)
- ethernet_interface.ipv6_auto_config.router.dns_config

<a id="canonical-f09a9cc868d06d2c2e828a276d3fbe459b52888341535d2923b96548ec0d0d9c"></a>

Type: `"single"`. Computed.

IPV6DnsConfig.

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

<a id="canonical-6626e82e7b8052fb116fa8d737ca09f1bebc487bd99149159b9a0b2767a9ae93"></a>

## Direct properties — ethernet_interface.ipv6_auto_config.router.dns_config / a238e8fefd63 / 3

- [configured_list](data-sources--network_interface--reference--group-001.md#canonical-883986ec67fb012d86c711c1958ba6733e02395273ae3e221c4fcd8c1e6c24df): complete subsection reference.

- [local_dns](data-sources--network_interface--reference--group-001.md#canonical-c5e027aa6f2475ecc6bcdf6f9832fdb680a9c16c4d299bc86c58889e42d9ecbe): complete subsection reference.

<a id="canonical-d0df7401e0619e4063ddaf7a68afcf5816343c8a166826024211955f898bf7cc"></a>

## Next pages — ethernet_interface.ipv6_auto_config.router.dns_config / a238e8fefd63 / 4

- [ethernet_interface.ipv6_auto_config.router.dns_config.configured_list](data-sources--network_interface--reference--group-001.md#canonical-883986ec67fb012d86c711c1958ba6733e02395273ae3e221c4fcd8c1e6c24df)
- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](data-sources--network_interface--reference--group-001.md#canonical-c5e027aa6f2475ecc6bcdf6f9832fdb680a9c16c4d299bc86c58889e42d9ecbe)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-09992a12de87aecff602762fe48959ff2c00c76dfeae4f4028a8b2ccf71518a1)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-883986ec67fb012d86c711c1958ba6733e02395273ae3e221c4fcd8c1e6c24df"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-14a99895af5502049d4ac8f375940c47fd63947210c38ac846e2b07e7e1b1292"></a>

## ethernet_interface.ipv6_auto_config.router.dns_config.configured_list — ethernet_interface.ipv6_auto_config.router.dns_config.configured_list / befb7ce1418d / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-7b12eb5b085cdc770291316c0876e0a8a6e9a5e76abde116ce768c70707c4b41)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-09992a12de87aecff602762fe48959ff2c00c76dfeae4f4028a8b2ccf71518a1)
- [ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--network_interface--reference--group-001.md#canonical-5bc25a772987541c959c266e81783f689c4e20795918d7a2188ca6cd0f70eccf)
- ethernet_interface.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-a502f4c063385d62c6c218e43756fc1d77903d8459f0f345afecd53981f01921"></a>

Type: `"single"`. Computed.

IPV6DnsList.

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

<a id="canonical-77d5431738888cfc5daeef78d26bcf0d3e636b725072bf1e5f19f72967ccb33c"></a>

## Direct properties — ethernet_interface.ipv6_auto_config.router.dns_config.configured_list / befb7ce1418d / 3

<a id="canonical-5412181c39d943907635707762bc2af4d4760693ce9b5d69af88098eef400a08"></a>

<a id="canonical-7dcb41ef2089fe8ef60127f8cbf50f686b3f71348e903003bb408648ac17c4b0"></a>

## dns_list property — ethernet_interface.ipv6_auto_config.router.dns_config.configured_list / befb7ce1418d / 4

Type: `["list", "string"]`. Computed.

List of IPv6 Addresses acting as DNS servers.

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

<a id="canonical-77d64d218a483afee0ff3974b9d25c5439d1be676a66929250ab498f5324f4cc"></a>

## Next pages — ethernet_interface.ipv6_auto_config.router.dns_config.configured_list / befb7ce1418d / 5

- [ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--network_interface--reference--group-001.md#canonical-5bc25a772987541c959c266e81783f689c4e20795918d7a2188ca6cd0f70eccf)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-c5e027aa6f2475ecc6bcdf6f9832fdb680a9c16c4d299bc86c58889e42d9ecbe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-99274c2d47ead85d588b38519345219384cd55e937aad4fe7f962477cc0c0339"></a>

## ethernet_interface.ipv6_auto_config.router.dns_config.local_dns — ethernet_interface.ipv6_auto_config.router.dns_config.local_dns / 7f37aa983879 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-7b12eb5b085cdc770291316c0876e0a8a6e9a5e76abde116ce768c70707c4b41)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-09992a12de87aecff602762fe48959ff2c00c76dfeae4f4028a8b2ccf71518a1)
- [ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--network_interface--reference--group-001.md#canonical-5bc25a772987541c959c266e81783f689c4e20795918d7a2188ca6cd0f70eccf)
- ethernet_interface.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-a34ac4372c642d785b3e0db93d8196631b3bf100cc91001c0ebafc2ae99a6e31"></a>

Type: `"single"`. Computed.

IPV6LocalDnsAddress.

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

<a id="canonical-04b9cbab871e72ee9488082ef738b84665175d9c4bd276b51f63fa46198b828e"></a>

## Direct properties — ethernet_interface.ipv6_auto_config.router.dns_config.local_dns / 7f37aa983879 / 3

<a id="canonical-f5f518197ab61febee5110baa0a89abaab204527e1984c9733707ba3f83b27e8"></a>

<a id="canonical-2d8304ea8d3cbeba60632085bf723b79da171f90237b3399a4505af0c5ca82cf"></a>

## configured_address property — ethernet_interface.ipv6_auto_config.router.dns_config.local_dns / 7f37aa983879 / 4

Type: `"string"`. Computed.

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Upstream description:

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

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

- [first_address](data-sources--network_interface--reference--group-001.md#canonical-b0c368817ca8cd010e6e038f84b0bd8e9cf6c323cc12f3f4244e0940572951d6): complete subsection reference.

- [last_address](data-sources--network_interface--reference--group-001.md#canonical-3873c0e4f3f4947940f08acc043fe0c21ab05d08e8c897ab8c68694b0136f47e): complete subsection reference.

<a id="canonical-8e2e547d7e1da9033edfa31c310df6a0ec556975f20acef2c4cf6bc2f9cb9641"></a>

## Next pages — ethernet_interface.ipv6_auto_config.router.dns_config.local_dns / 7f37aa983879 / 5

- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address](data-sources--network_interface--reference--group-001.md#canonical-b0c368817ca8cd010e6e038f84b0bd8e9cf6c323cc12f3f4244e0940572951d6)
- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address](data-sources--network_interface--reference--group-001.md#canonical-3873c0e4f3f4947940f08acc043fe0c21ab05d08e8c897ab8c68694b0136f47e)
- [ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--network_interface--reference--group-001.md#canonical-5bc25a772987541c959c266e81783f689c4e20795918d7a2188ca6cd0f70eccf)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-b0c368817ca8cd010e6e038f84b0bd8e9cf6c323cc12f3f4244e0940572951d6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4cb75b1caf2490af61300a8b9a71d657e0162aedf61b9cd84263ad37dc59449e"></a>

## ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address — ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address / cabd42d32c50 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-7b12eb5b085cdc770291316c0876e0a8a6e9a5e76abde116ce768c70707c4b41)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-09992a12de87aecff602762fe48959ff2c00c76dfeae4f4028a8b2ccf71518a1)
- [ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--network_interface--reference--group-001.md#canonical-5bc25a772987541c959c266e81783f689c4e20795918d7a2188ca6cd0f70eccf)
- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](data-sources--network_interface--reference--group-001.md#canonical-c5e027aa6f2475ecc6bcdf6f9832fdb680a9c16c4d299bc86c58889e42d9ecbe)
- ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-c40793311ab5385a2c6299c22d68458eab8fff1e5e480f6a4bb168ecda8716f3"></a>

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

<a id="canonical-a115005de4142257890caf05d86e6a771911b3702d198bd3f1a954c2d58fc351"></a>

## Direct properties — ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address / cabd42d32c50 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-facf0ff2f588597e51cdb0dfdc1365dc8bd625fa15479d287e6ecdc1f9a58630"></a>

## Next pages — ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address / cabd42d32c50 / 4

- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](data-sources--network_interface--reference--group-001.md#canonical-c5e027aa6f2475ecc6bcdf6f9832fdb680a9c16c4d299bc86c58889e42d9ecbe)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-3873c0e4f3f4947940f08acc043fe0c21ab05d08e8c897ab8c68694b0136f47e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-92bfd0efa5058d13efb6224d8afdfb831959109a5ac9ac44b7821a30a039e86b"></a>

## ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address — ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address / 1ac168d23a1c / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-7b12eb5b085cdc770291316c0876e0a8a6e9a5e76abde116ce768c70707c4b41)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-09992a12de87aecff602762fe48959ff2c00c76dfeae4f4028a8b2ccf71518a1)
- [ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--network_interface--reference--group-001.md#canonical-5bc25a772987541c959c266e81783f689c4e20795918d7a2188ca6cd0f70eccf)
- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](data-sources--network_interface--reference--group-001.md#canonical-c5e027aa6f2475ecc6bcdf6f9832fdb680a9c16c4d299bc86c58889e42d9ecbe)
- ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-1f9e39e86c18841990b8c123986681a19fdd85cd70aceb34dd40ffa16a9f2660"></a>

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

<a id="canonical-40bb1f7c4e67a2b1f2b4423ca85971a37fc2a2041d8b99c3cadca86313c7393b"></a>

## Direct properties — ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address / 1ac168d23a1c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-39be144776ccd8bada96707acabd796e92120d713a9a0bae9690b9851cc22658"></a>

## Next pages — ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address / 1ac168d23a1c / 4

- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](data-sources--network_interface--reference--group-001.md#canonical-c5e027aa6f2475ecc6bcdf6f9832fdb680a9c16c4d299bc86c58889e42d9ecbe)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-94c2be9d9397047ff932739152da71df44728a298cfb39a7dec7a79d2135b808"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ec5826fa28492b239d34caab955da480ba48e1455798333a0978e9733d7149b"></a>

## ethernet_interface.ipv6_auto_config.router.stateful — ethernet_interface.ipv6_auto_config.router.stateful / b9a50a3a810a / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-7b12eb5b085cdc770291316c0876e0a8a6e9a5e76abde116ce768c70707c4b41)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-09992a12de87aecff602762fe48959ff2c00c76dfeae4f4028a8b2ccf71518a1)
- ethernet_interface.ipv6_auto_config.router.stateful

<a id="canonical-2591ffd1665b732807013fd93c9f86ca06f17ee8e938ae66da9dd159971482fd"></a>

Type: `"single"`. Computed.

DHCPIPV6 Stateful Server.

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

<a id="canonical-0f492a792d5ec2f8aff6bd4721059fdd30aa1ac2e1436c9a6860655f9316500d"></a>

## Direct properties — ethernet_interface.ipv6_auto_config.router.stateful / b9a50a3a810a / 3

- [automatic_from_end](data-sources--network_interface--reference--group-001.md#canonical-0018fb209beb4e2dcbfb875a975b31ecd92950bad1d28a27f7c99ce068d4332b): complete subsection reference.

- [automatic_from_start](data-sources--network_interface--reference--group-001.md#canonical-843df77fa0215d54ef81e7677b5956a1db81da4cc6ca4502ef799c853b631d6f): complete subsection reference.

- [dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-ed9950af074c25dd6092613d02766dde8b8d27c45f14d6f46004497e3adcfdb5): complete subsection reference.

<a id="canonical-2291391ad5516c2e29b15c7d18a96cdf2e5f6fc00d17c94a563afe7defda4938"></a>

<a id="canonical-844ab98f5d36d975feaf19ea8e1918288dad988f334322c166c637f95a130f26"></a>

## fixed_ip_map property — ethernet_interface.ipv6_auto_config.router.stateful / b9a50a3a810a / 4

Type: `["map", "string"]`. Computed.

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

- [interface_ip_map](data-sources--network_interface--reference--group-001.md#canonical-3741aa5f16c1be2ed456288f8f665107be477a82202867592e38d43e66936aa0): complete subsection reference.

<a id="canonical-45eb70f752e2c0296917bd8c2789ce481305fc2997814df35a70f4c69d34ed5a"></a>

## Next pages — ethernet_interface.ipv6_auto_config.router.stateful / b9a50a3a810a / 5

- [ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end](data-sources--network_interface--reference--group-001.md#canonical-0018fb209beb4e2dcbfb875a975b31ecd92950bad1d28a27f7c99ce068d4332b)
- [ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start](data-sources--network_interface--reference--group-001.md#canonical-843df77fa0215d54ef81e7677b5956a1db81da4cc6ca4502ef799c853b631d6f)
- [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-ed9950af074c25dd6092613d02766dde8b8d27c45f14d6f46004497e3adcfdb5)
- [ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map](data-sources--network_interface--reference--group-001.md#canonical-3741aa5f16c1be2ed456288f8f665107be477a82202867592e38d43e66936aa0)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-09992a12de87aecff602762fe48959ff2c00c76dfeae4f4028a8b2ccf71518a1)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-0018fb209beb4e2dcbfb875a975b31ecd92950bad1d28a27f7c99ce068d4332b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a568c10588cf02b2846ceea0fd3b5ca76f3357f30516ad70b2c95c376f54ce4b"></a>

## ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end — ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end / dbd8edde4028 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-7b12eb5b085cdc770291316c0876e0a8a6e9a5e76abde116ce768c70707c4b41)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-09992a12de87aecff602762fe48959ff2c00c76dfeae4f4028a8b2ccf71518a1)
- [ethernet_interface.ipv6_auto_config.router.stateful](data-sources--network_interface--reference--group-001.md#canonical-94c2be9d9397047ff932739152da71df44728a298cfb39a7dec7a79d2135b808)
- ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-e6bf166e9c6f517ed0350a9ab9c075ce19306bf1c9457170ba67d54344a142b4"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-f9f3a915b75604e5877fabaf503f4b4a1572ecec98c6b7e4fcd2552cb2a832cb"></a>

## Direct properties — ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end / dbd8edde4028 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-542b6028f89ad493611272f8aa02630c380d92553ffeba0914d2f871479f233f"></a>

## Next pages — ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end / dbd8edde4028 / 4

- [ethernet_interface.ipv6_auto_config.router.stateful](data-sources--network_interface--reference--group-001.md#canonical-94c2be9d9397047ff932739152da71df44728a298cfb39a7dec7a79d2135b808)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-843df77fa0215d54ef81e7677b5956a1db81da4cc6ca4502ef799c853b631d6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5b3f5a06bc11256d2fe6fb1e9bb3a26d7582c04ff774b49b52d4190cb061220"></a>

## ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start — ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start / 4123aa9793f8 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-7b12eb5b085cdc770291316c0876e0a8a6e9a5e76abde116ce768c70707c4b41)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-09992a12de87aecff602762fe48959ff2c00c76dfeae4f4028a8b2ccf71518a1)
- [ethernet_interface.ipv6_auto_config.router.stateful](data-sources--network_interface--reference--group-001.md#canonical-94c2be9d9397047ff932739152da71df44728a298cfb39a7dec7a79d2135b808)
- ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-dfad90a1bdd79c109d170691ae329491e679ec52962da2c2e5d2674cf756feb5"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-919b54caf9b6dd41cf9fb5ada615b637d52b7183b21e68ccecf84da50f60fe4c"></a>

## Direct properties — ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start / 4123aa9793f8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-12095cc8e765b6c9515540a9433e110d2a06ce5e38fc1aedabe3368d501a43b1"></a>

## Next pages — ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start / 4123aa9793f8 / 4

- [ethernet_interface.ipv6_auto_config.router.stateful](data-sources--network_interface--reference--group-001.md#canonical-94c2be9d9397047ff932739152da71df44728a298cfb39a7dec7a79d2135b808)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-ed9950af074c25dd6092613d02766dde8b8d27c45f14d6f46004497e3adcfdb5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-51542904d4d3195c0166a157e8d02c33f78e03b226e01df8b8bffa7cfa8140c2"></a>

## ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks — ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks / e3736264da00 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-7b12eb5b085cdc770291316c0876e0a8a6e9a5e76abde116ce768c70707c4b41)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-09992a12de87aecff602762fe48959ff2c00c76dfeae4f4028a8b2ccf71518a1)
- [ethernet_interface.ipv6_auto_config.router.stateful](data-sources--network_interface--reference--group-001.md#canonical-94c2be9d9397047ff932739152da71df44728a298cfb39a7dec7a79d2135b808)
- ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-255f756ecdaa570a62428a25b228d26ebcb5e7ef0e235c642dd01f6e499eb9c0"></a>

Type: `"list"`. Computed.

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

<a id="canonical-fb3ddb0aacac15cc9ae2e24f6d1476850612182aed364dcfaa9b57aaae9490d7"></a>

## Direct properties — ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks / e3736264da00 / 3

<a id="canonical-a8e0d9ecec0e26a1ebebc8d8cbeaf1a0850e901a7cfafaceb9bb2a07933a9dea"></a>

<a id="canonical-e5a204a59e7e91177702dc2be3864598ce06d10e5478061268728f8e7fd3fd36"></a>

## network_prefix property — ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks / e3736264da00 / 4

Type: `"string"`. Computed.

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

<a id="canonical-597be91254c52a1729afcfbc635fd7d710f548d767890229025478134b2e2fc1"></a>

<a id="canonical-afdefc3c42db8201f8f14db0ee92eeca583b82ab7939234d194d8684bdc90faf"></a>

## pool_settings property — ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks / e3736264da00 / 5

Type: `"string"`. Computed.

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

- [pools](data-sources--network_interface--reference--group-001.md#canonical-9a3ac4443042321b6c7c9679d47470f70a8db19c11b84b216acde1cb1bcda72e): complete subsection reference.

<a id="canonical-3f15653d9cd088ee1d7db3e0f2f6dd0674d342684e699da9c0894e879a5c6eb0"></a>

## Next pages — ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks / e3736264da00 / 6

- [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools](data-sources--network_interface--reference--group-001.md#canonical-9a3ac4443042321b6c7c9679d47470f70a8db19c11b84b216acde1cb1bcda72e)
- [ethernet_interface.ipv6_auto_config.router.stateful](data-sources--network_interface--reference--group-001.md#canonical-94c2be9d9397047ff932739152da71df44728a298cfb39a7dec7a79d2135b808)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-9a3ac4443042321b6c7c9679d47470f70a8db19c11b84b216acde1cb1bcda72e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-035d0fd8832975edc6b01f7b2852434ea9675b06499fffff63210143fb10b7da"></a>

## ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools — ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools / 2a4a94acd735 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-7b12eb5b085cdc770291316c0876e0a8a6e9a5e76abde116ce768c70707c4b41)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-09992a12de87aecff602762fe48959ff2c00c76dfeae4f4028a8b2ccf71518a1)
- [ethernet_interface.ipv6_auto_config.router.stateful](data-sources--network_interface--reference--group-001.md#canonical-94c2be9d9397047ff932739152da71df44728a298cfb39a7dec7a79d2135b808)
- [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-ed9950af074c25dd6092613d02766dde8b8d27c45f14d6f46004497e3adcfdb5)
- ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-0fa37a54c99f3ed748a08f97d52493cf9478a35fa25e2a2a406d5c4b0aad81f2"></a>

Type: `"list"`. Computed.

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

<a id="canonical-3c0a43974e096bb4fa04d776b5ec0afcd40b1d44f57b0a962889977cea322bdc"></a>

## Direct properties — ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools / 2a4a94acd735 / 3

<a id="canonical-5896178f188c5e81b9402c74f4f6f95842b42068aa6f7ab43a3ed2d21b7ed826"></a>

<a id="canonical-763bdbe7bb79887f7f32321b172b59d407ec6fb1add069ebdeb1cb9ff453d7bc"></a>

## end_ip property — ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools / 2a4a94acd735 / 4

Type: `"string"`. Computed.

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

Upstream description:

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

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

<a id="canonical-e535cf166cbfd3d08eb25adbf6af51e5bb048aab0bad70c4ac1f218e76597344"></a>

<a id="canonical-678ee9efbf0a37ddad3c4d2786a6b29ab3020701adf566af235680819fb1981c"></a>

## start_ip property — ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools / 2a4a94acd735 / 5

Type: `"string"`. Computed.

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

Upstream description:

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

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

<a id="canonical-94d4fae79d0e173e6f04892ef1437135fcd851ad6bde2edb1a9fe675966355fb"></a>

## Next pages — ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools / 2a4a94acd735 / 6

- [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-ed9950af074c25dd6092613d02766dde8b8d27c45f14d6f46004497e3adcfdb5)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-3741aa5f16c1be2ed456288f8f665107be477a82202867592e38d43e66936aa0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d28b3590e68b480b65ba468f9af5cdb047f955afd1fc6d9d80db400e60508ee1"></a>

## ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map — ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map / 4f7d7b3e252a / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-7b12eb5b085cdc770291316c0876e0a8a6e9a5e76abde116ce768c70707c4b41)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-09992a12de87aecff602762fe48959ff2c00c76dfeae4f4028a8b2ccf71518a1)
- [ethernet_interface.ipv6_auto_config.router.stateful](data-sources--network_interface--reference--group-001.md#canonical-94c2be9d9397047ff932739152da71df44728a298cfb39a7dec7a79d2135b808)
- ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-145942f17d945d03a0cb74817243d3d542162823f9c68eb309d2a70f904b5204"></a>

Type: `"single"`. Computed.

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

<a id="canonical-99fc33d12aa606afa25bc476ad46ff63d297db9a6e42442b9d9730dd15a0e048"></a>

## Direct properties — ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map / 4f7d7b3e252a / 3

<a id="canonical-88eac661e4c606f93ea3f7fe5216657ecf2074e3946f49a81a7ca2f3c1ccca81"></a>

<a id="canonical-ab2afd0bdbe7c8fc02680ad4b21ce04b55686f083993f6bbdd4fa93a41353b32"></a>

## interface_ip_map property — ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map / 4f7d7b3e252a / 4

Type: `["map", "string"]`. Computed.

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

<a id="canonical-f1cd89480708a4be3e85cc20afe578298c86cdb05a274ad2cab5c38424461426"></a>

## Next pages — ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map / 4f7d7b3e252a / 5

- [ethernet_interface.ipv6_auto_config.router.stateful](data-sources--network_interface--reference--group-001.md#canonical-94c2be9d9397047ff932739152da71df44728a298cfb39a7dec7a79d2135b808)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-1de8a146ae46e90e592cf7bc92684f0a33166b790c639ec602fab32bbd7a33e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26a7e20d89fcfd5e2f5b425521b4b3601573a367520d81dbba0754b6bb973135"></a>

## ethernet_interface.is_primary — ethernet_interface.is_primary / fa253ec58c24 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- ethernet_interface.is_primary

<a id="canonical-272504013cf4717e71c2473a90ca18d8bf7556be3911889e94adf5e4ea7fc48e"></a>

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

<a id="canonical-e082d9f8942c9d243e21612157084dce0782d8d34486205c34012962002d2ce4"></a>

## Direct properties — ethernet_interface.is_primary / fa253ec58c24 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-12b873138479c9af7c579ae27864f4b89b0809c13d90f5b6854fa33b4e962ef0"></a>

## Next pages — ethernet_interface.is_primary / fa253ec58c24 / 4

- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-7e59e10a823eaead070dcef8c323e667a853846698f1bd876926d367a56f9521"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4cae228d9d19979e86f4e1db5f781798a3f85a1942423fdcaca2e1e86dacf8b2"></a>

## ethernet_interface.monitor — ethernet_interface.monitor / c87b58b65ef9 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- ethernet_interface.monitor

<a id="canonical-4fde18526237441ca005ff4c402e750236a0f624b16ba704d600f1ddd48029c5"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-ef53a9c15c3f86928b3bbaee0636ba13631d01566795a19b60d69bcc25b72772"></a>

## Direct properties — ethernet_interface.monitor / c87b58b65ef9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-35de61bac1e18be72c9a7fb155a90a0ca53bd11687a6c7371f9254787aa7d131"></a>

## Next pages — ethernet_interface.monitor / c87b58b65ef9 / 4

- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-d9f070c79019cfaa5f37f86a89e79b9582c43664ba57efac2914032d8bbe3397"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f9ef12365681ff28de82968676983b84f0995a6a6cd410935e170d3533fbd6b6"></a>

## ethernet_interface.monitor_disabled — ethernet_interface.monitor_disabled / 9c6d818f792a / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- ethernet_interface.monitor_disabled

<a id="canonical-6f5ba37cc5702ea92aaf9f6b246bb9e8b7217b30d6c2ff406f5fae20d2744f5c"></a>

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

<a id="canonical-268147c60f1407a10ac4171a16770759e86f0afeca0f7ee3d111972b4d01f812"></a>

## Direct properties — ethernet_interface.monitor_disabled / 9c6d818f792a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1d414faa3af96e901e39c7499241503078106166f884ba7c4b1124aab9c73088"></a>

## Next pages — ethernet_interface.monitor_disabled / 9c6d818f792a / 4

- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-573b9e41b481f04ee78288ebc0a173626ca7680662d4131a3974064055dfb395"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be0e89a54b3f715e363473f279de4ff933795a5bd507b09ef527b3c2fa9f8153"></a>

## ethernet_interface.no_ipv6_address — ethernet_interface.no_ipv6_address / cbe7fd02ddcc / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- ethernet_interface.no_ipv6_address

<a id="canonical-f4aa795ad9870d874d3a77605489ac4c0736e7284aeda008217a1f7cdcdb5b42"></a>

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

<a id="canonical-cd3503ac79d28d07c82951ebcb78c1be5e3898d0277ee6dca1ba840098aacca3"></a>

## Direct properties — ethernet_interface.no_ipv6_address / cbe7fd02ddcc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-caf40b462ec16dca39424c0a7ab6d465b6efbb88f3780d121da5f15c6d73d8db"></a>

## Next pages — ethernet_interface.no_ipv6_address / cbe7fd02ddcc / 4

- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-f6b72b300a33a5d0a16dd78cd0144638f4b9024d74cd353a4d1ea773dc4b2dfc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d2b52d75840b77b5751ff93dc45e1e0301968d3c2c90112b04c626943e83380"></a>

## ethernet_interface.not_primary — ethernet_interface.not_primary / 8a49b4f2c390 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- ethernet_interface.not_primary

<a id="canonical-ad4ed6be239a874e3e0b92500208d19e743225083a42389f010a40a5a24c213c"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for not primary.

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

<a id="canonical-f2014fe3e647191fa4695b2d2ed51bd551ae2e84d5483a544e38499838c25d81"></a>

## Direct properties — ethernet_interface.not_primary / 8a49b4f2c390 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c78a92eb11a465e5975e2ccb23906c91737de3801a7ad0421538a0e57f39683e"></a>

## Next pages — ethernet_interface.not_primary / 8a49b4f2c390 / 4

- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-26db361e57025a942717ab5ae7b3ed631fc466b9d2ab1b08aece2bfbf6740c04"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5eb6debe758c8ef98bab124bc2f0cb98da723ecd5c388035179dd53df7330627"></a>

## ethernet_interface.site_local_inside_network — ethernet_interface.site_local_inside_network / 4bad2562d9a4 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- ethernet_interface.site_local_inside_network

<a id="canonical-d6f24b2e1d87f515fb70e214d3358fdf3f54ca4d10e4a6f78e82c5b61371a5fb"></a>

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

<a id="canonical-e59942ba0a311700d799e3e059c2563201c6555a303c5bc7a016536151f23701"></a>

## Direct properties — ethernet_interface.site_local_inside_network / 4bad2562d9a4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5442a9a5b116aae50cc10edcb80518de9c81bca306e039eb9cb6c461032c80a9"></a>

## Next pages — ethernet_interface.site_local_inside_network / 4bad2562d9a4 / 4

- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-de90c1d6529c71ee3906c0de7866ef0b3558698a9733c41930f7f064cf3dc849"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-929a1a2b12561b25a59a1dcee652d45fb779f2be4fece0473c46f68015a85843"></a>

## ethernet_interface.site_local_network — ethernet_interface.site_local_network / 3306732e8fd7 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- ethernet_interface.site_local_network

<a id="canonical-63e80dfff901fb24c8718cd4493d86e2dd790510027ecb54870af1b82bbbbb54"></a>

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

<a id="canonical-da15c50ef51dddb8fe3b9d4ef86c11e8619e98c23cc65208af4edcba904f2bf5"></a>

## Direct properties — ethernet_interface.site_local_network / 3306732e8fd7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cc688ae96f787e489c5dfde9320fd340a752cac7b3522618c516c7cdad697c42"></a>

## Next pages — ethernet_interface.site_local_network / 3306732e8fd7 / 4

- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-f6a7a273d85d7cfa1c8614e9eecb3589eb3848255af051b6a00adcd274fd76ed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ed2f16824537d1f3627e42736502544c809a290a7773010750fd3c021c28931"></a>

## ethernet_interface.static_ip — ethernet_interface.static_ip / 9eeffd0737ec / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- ethernet_interface.static_ip

<a id="canonical-f07c55637c9b1a25b38b01ad551f29bd820199ecbe30902bb07460d61422d73b"></a>

Type: `"single"`. Computed.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

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

<a id="canonical-499b46e3534efe0b6e3ed9ad7b38cd5db67474a4cdcd1591f3c2c87f1f8a747f"></a>

## Direct properties — ethernet_interface.static_ip / 9eeffd0737ec / 3

- [cluster_static_ip](data-sources--network_interface--reference--group-001.md#canonical-1f632d2cd614a214a6c70404e0271028222007b00bbdf1d693358731af8be487): complete subsection reference.

- [node_static_ip](data-sources--network_interface--reference--group-001.md#canonical-ec35aff45c1cb3b1f317d9a1adfb88edfb1e4a66f2008ab82455910db6d62352): complete subsection reference.

<a id="canonical-14ed836cbc9d5578f2b6d8932d48c1bea7af7f1827a31a4e454401ac2b5e8ff0"></a>

## Next pages — ethernet_interface.static_ip / 9eeffd0737ec / 4

- [ethernet_interface.static_ip.cluster_static_ip](data-sources--network_interface--reference--group-001.md#canonical-1f632d2cd614a214a6c70404e0271028222007b00bbdf1d693358731af8be487)
- [ethernet_interface.static_ip.node_static_ip](data-sources--network_interface--reference--group-001.md#canonical-ec35aff45c1cb3b1f317d9a1adfb88edfb1e4a66f2008ab82455910db6d62352)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-1f632d2cd614a214a6c70404e0271028222007b00bbdf1d693358731af8be487"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cba2b1b02bf04fdba7a41e74b9371c8f3c68eada56fdc4dcd91475a1d5068940"></a>

## ethernet_interface.static_ip.cluster_static_ip — ethernet_interface.static_ip.cluster_static_ip / 706842c4ca6a / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [ethernet_interface.static_ip](data-sources--network_interface--reference--group-001.md#canonical-f6a7a273d85d7cfa1c8614e9eecb3589eb3848255af051b6a00adcd274fd76ed)
- ethernet_interface.static_ip.cluster_static_ip

<a id="canonical-4d721af642105b494b7581adf6dc86c347732d26954cb293f27cee0066eb0a78"></a>

Type: `"single"`. Computed.

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

<a id="canonical-7bfd6c783ce3ef60885114879cd7e93418623dd8e6c38cf4c27fc4766b1d85a2"></a>

## Direct properties — ethernet_interface.static_ip.cluster_static_ip / 706842c4ca6a / 3

<a id="canonical-c42e1280024e09675d4781e52cc2db305557fd8f02882de66e4d059b389cf53b"></a>

<a id="canonical-d9ef80352f2c2e5cdc3ce00aa3960852057b3d1d536bb8e8ee951e009bf36f0b"></a>

## interface_ip_map property — ethernet_interface.static_ip.cluster_static_ip / 706842c4ca6a / 4

Type: `["map", "string"]`. Computed.

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

<a id="canonical-16a3c7cff8e3b20f6e8ef0cf00bcb3a3526a041f85119007b865f03574c89af4"></a>

## Next pages — ethernet_interface.static_ip.cluster_static_ip / 706842c4ca6a / 5

- [ethernet_interface.static_ip](data-sources--network_interface--reference--group-001.md#canonical-f6a7a273d85d7cfa1c8614e9eecb3589eb3848255af051b6a00adcd274fd76ed)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-ec35aff45c1cb3b1f317d9a1adfb88edfb1e4a66f2008ab82455910db6d62352"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e37017c90261dfe0954475c4479bb90eb98694a79c7e3235e6d83b7e7841fd1f"></a>

## ethernet_interface.static_ip.node_static_ip — ethernet_interface.static_ip.node_static_ip / 380b2fecce0b / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [ethernet_interface.static_ip](data-sources--network_interface--reference--group-001.md#canonical-f6a7a273d85d7cfa1c8614e9eecb3589eb3848255af051b6a00adcd274fd76ed)
- ethernet_interface.static_ip.node_static_ip

<a id="canonical-e850cf6c4e0a238bbf12185ce3bd3d735ee4634e1cd60116f505287322196270"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for a node.

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

<a id="canonical-16ef570f063778df64fdd725be578889d593c3b0eafc5a7ab4264c6920629e4b"></a>

## Direct properties — ethernet_interface.static_ip.node_static_ip / 380b2fecce0b / 3

<a id="canonical-d428695040f29a7e848ca0b1566669c05e8e929bf348a5ca3c22c074ecdb25f3"></a>

<a id="canonical-f2a040fd45a11e0a06351c20bdcb984ee4e0b4c46d13e2bdbf7598150580f115"></a>

## default_gw property — ethernet_interface.static_ip.node_static_ip / 380b2fecce0b / 4

Type: `"string"`. Computed.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

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

<a id="canonical-5c569cd502d3fc0b7d745215efbfb40af475dc3d8c7372fef8a9e53bd4306d07"></a>

<a id="canonical-06464e438c65fa01a3529323e2baf19da768899150c0a2b609cf88dff76d57ed"></a>

## dns_server property — ethernet_interface.static_ip.node_static_ip / 380b2fecce0b / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-935e486a053f3989926320ac9d25cf35b2e2fe99a39bbe8d633d8ed3b0740c21"></a>

<a id="canonical-fa2f55f1a029e5e982f77068d88af46384c7bf79c56489710b01f2812f0cd83a"></a>

## ip_address property — ethernet_interface.static_ip.node_static_ip / 380b2fecce0b / 6

Type: `"string"`. Computed.

IP address of the interface and prefix length.

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
