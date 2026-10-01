---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-3a9bb6234cd405101365f7e1700c549bbfa1f3fcb01ad1ce12a5a833da8701be"></a>

## vmware.not_managed.node_list.interface_list — vmware.not_managed.node_list.interface_list / 1fd7728e8782 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- vmware.not_managed.node_list.interface_list

<a id="canonical-b387fe228fa40651113e049e2228b247e0b5d793b4181a29f87c89a29c4a1635"></a>

Type: `"list"`. Computed.

Manage interfaces belonging to this node.

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

<a id="canonical-553b88c04c72412b95020d5a720860f03ab9146604a596226e55ce9c87108d4a"></a>

## Direct properties — vmware.not_managed.node_list.interface_list / 1fd7728e8782 / 3

- [bond_interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-d7bfdee621d5eca93bd783d969229a48f6299a1d74b3f667f66a661bdfee2133): complete subsection reference.

<a id="canonical-9739b6efab1f1f253593d97b222ebc16b10f790d92e5047123c3d8fc8e624c12"></a>

<a id="canonical-20487659413e9fb91c3afa580f5c983b094f58ee1c7b285298adb0798fbf6d49"></a>

## description_spec property — vmware.not_managed.node_list.interface_list / 1fd7728e8782 / 4

Type: `"string"`. Computed.

Interface Description. Description for this Interface.

- [dhcp_client](data-sources--securemesh_site_v2--reference--group-017.md#canonical-820048c87747341ee92a305eb48331c1ad7170c3802f93011c1b29a15d1fb513): complete subsection reference.

- [dhcp_server](data-sources--securemesh_site_v2--reference--group-017.md#canonical-90db80840d6492b757411db0899cb8591bf2cbe35d58e7264e798815fd1cb85c): complete subsection reference.

- [ethernet_interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-9d944837d4c7972a526e05cea9804ab6f5f114eb7ef9e7670a296540362f28ef): complete subsection reference.

- [ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-017.md#canonical-77129aec4ff1f628e53657b8ee2ad737e637dabc8b2c251ad5c258f92e52a9fa): complete subsection reference.

<a id="canonical-9986d09e4d83101a80022f9b547aa737c88136d88a9e47cefefde272e63f2b59"></a>

<a id="canonical-e8d02daa377d49a780bfdce103d395399875a5d2c890f878152068a4bdae7363"></a>

## is_management property — vmware.not_managed.node_list.interface_list / 1fd7728e8782 / 5

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-b5139ec647fdb584ff968b57b56c3da623433471b563935b42048f59c3890c14"></a>

<a id="canonical-1865bfdc372fcba04ac49cf9be367c7734f11d1aee4a940e8c78a05d7250c4f4"></a>

## is_primary property — vmware.not_managed.node_list.interface_list / 1fd7728e8782 / 6

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-438788153610446f3b4620dde41845ce03ac4e1e297fc20f14f5b583e9561861"></a>

<a id="canonical-031885dfdf348f919768445ea2841d6ea393c31b16b2203cb2c400f7faa280e9"></a>

## labels property — vmware.not_managed.node_list.interface_list / 1fd7728e8782 / 7

Type: `["map", "string"]`. Computed.

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

- [monitor](data-sources--securemesh_site_v2--reference--group-017.md#canonical-c0d532ea954a4fd60833369da7f4607c9a2d91db72c587e740164bfc42f33f20): complete subsection reference.

- [monitor_disabled](data-sources--securemesh_site_v2--reference--group-017.md#canonical-d5a180851c2a785e2cf71829e969187cac0b0f828c10ea3c977c3ca059329a21): complete subsection reference.

<a id="canonical-299645ef098e69be0be2125494184960e5c28126fe61a748033a4277f2345b5f"></a>

<a id="canonical-c5d6d953aec3867573d0307454de141c733c474bfae2a5b985b06349aba4ab0b"></a>

## mtu property — vmware.not_managed.node_list.interface_list / 1fd7728e8782 / 8

Type: `"number"`. Computed.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

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

<a id="canonical-828cb9e6a958314cfbef835db5fcfec8668a607fd63725614221b7d20c534368"></a>

<a id="canonical-30f5aebd58bb159dcb6d8df2afc9828b598b8f676af7022dc313b48308192cd3"></a>

## name property — vmware.not_managed.node_list.interface_list / 1fd7728e8782 / 9

Type: `"string"`. Computed.

Interface Name. Name of this Interface.

Upstream description:

Name of this Interface.

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

- [network_option](data-sources--securemesh_site_v2--reference--group-017.md#canonical-f872def9ca690acf4a7892b9854268c20f37c626269035aca2e242e8fa89f7b8): complete subsection reference.

- [no_ipv4_address](data-sources--securemesh_site_v2--reference--group-017.md#canonical-75c45b1993190ce2ad19a20d8e5b01f216e6775bce3004c8506756ec8fceb0c6): complete subsection reference.

- [no_ipv6_address](data-sources--securemesh_site_v2--reference--group-017.md#canonical-a9b19516002b4330fd2f2374b11caf3cea679396fc69d21c7a0e7abceb217584): complete subsection reference.

<a id="canonical-284402a2c624c9ebef934635d9f798b8cfc7fb38a265ce0b929369f35ec5cf19"></a>

<a id="canonical-c1fdcf35a70399aa86e6b15c31b0890ba6076f087db0b0e77f102c639f18abbc"></a>

## priority property — vmware.not_managed.node_list.interface_list / 1fd7728e8782 / 10

Type: `"number"`. Computed.

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Upstream description:

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

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

- [site_to_site_connectivity_interface_disabled](data-sources--securemesh_site_v2--reference--group-017.md#canonical-9995a1d8d38c789c34d29c5ff06da51cb98796fb20dcaac83b30028c3ad6c6d0): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](data-sources--securemesh_site_v2--reference--group-017.md#canonical-8c2dee918558431a357fa0cc5fc17d639304b7e01b5d4dee3b08c42a2d4b1283): complete subsection reference.

- [static_ip](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3a1d19b4aaf04e2319159d92244412ad68de95e84de0a9b3380ca724ff0c0314): complete subsection reference.

- [static_ipv6_address](data-sources--securemesh_site_v2--reference--group-017.md#canonical-119877391dc95ffeb676073bd50f4e8f0fb10a1efaaeab1e0564044c3e425117): complete subsection reference.

- [vlan_interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2c336614881f32108bf45b8d32c9338cb9eb1b6eab663f99a48f91edcef0c177): complete subsection reference.

<a id="canonical-4637ef1aa40bfaf38530064e5b3a07100d11751a5c37d3b698ddc13759fbd408"></a>

## Next pages — vmware.not_managed.node_list.interface_list / 1fd7728e8782 / 11

- [vmware.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-d7bfdee621d5eca93bd783d969229a48f6299a1d74b3f667f66a661bdfee2133)
- [vmware.not_managed.node_list.interface_list.dhcp_client](data-sources--securemesh_site_v2--reference--group-017.md#canonical-820048c87747341ee92a305eb48331c1ad7170c3802f93011c1b29a15d1fb513)
- [vmware.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-017.md#canonical-90db80840d6492b757411db0899cb8591bf2cbe35d58e7264e798815fd1cb85c)
- [vmware.not_managed.node_list.interface_list.ethernet_interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-9d944837d4c7972a526e05cea9804ab6f5f114eb7ef9e7670a296540362f28ef)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-017.md#canonical-77129aec4ff1f628e53657b8ee2ad737e637dabc8b2c251ad5c258f92e52a9fa)
- [vmware.not_managed.node_list.interface_list.monitor](data-sources--securemesh_site_v2--reference--group-017.md#canonical-c0d532ea954a4fd60833369da7f4607c9a2d91db72c587e740164bfc42f33f20)
- [vmware.not_managed.node_list.interface_list.monitor_disabled](data-sources--securemesh_site_v2--reference--group-017.md#canonical-d5a180851c2a785e2cf71829e969187cac0b0f828c10ea3c977c3ca059329a21)
- [vmware.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-017.md#canonical-f872def9ca690acf4a7892b9854268c20f37c626269035aca2e242e8fa89f7b8)
- [vmware.not_managed.node_list.interface_list.no_ipv4_address](data-sources--securemesh_site_v2--reference--group-017.md#canonical-75c45b1993190ce2ad19a20d8e5b01f216e6775bce3004c8506756ec8fceb0c6)
- [vmware.not_managed.node_list.interface_list.no_ipv6_address](data-sources--securemesh_site_v2--reference--group-017.md#canonical-a9b19516002b4330fd2f2374b11caf3cea679396fc69d21c7a0e7abceb217584)
- [vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled](data-sources--securemesh_site_v2--reference--group-017.md#canonical-9995a1d8d38c789c34d29c5ff06da51cb98796fb20dcaac83b30028c3ad6c6d0)
- [vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled](data-sources--securemesh_site_v2--reference--group-017.md#canonical-8c2dee918558431a357fa0cc5fc17d639304b7e01b5d4dee3b08c42a2d4b1283)
- [vmware.not_managed.node_list.interface_list.static_ip](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3a1d19b4aaf04e2319159d92244412ad68de95e84de0a9b3380ca724ff0c0314)
- [vmware.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-017.md#canonical-119877391dc95ffeb676073bd50f4e8f0fb10a1efaaeab1e0564044c3e425117)
- [vmware.not_managed.node_list.interface_list.vlan_interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2c336614881f32108bf45b8d32c9338cb9eb1b6eab663f99a48f91edcef0c177)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-d7bfdee621d5eca93bd783d969229a48f6299a1d74b3f667f66a661bdfee2133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0c51a107c49662213b831541e4ccbded63d4d75a9fa4df94db53173b18bd87b5"></a>

## vmware.not_managed.node_list.interface_list.bond_interface — vmware.not_managed.node_list.interface_list.bond_interface / 0e2430aa3991 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- vmware.not_managed.node_list.interface_list.bond_interface

<a id="canonical-8608ac064aee30d1ac7124d7e36f25a3f7c5782e494fd516aa012a8a808ce4ac"></a>

Type: `"single"`. Computed.

Configuration parameter for bond interface.

Upstream description:

Bond devices configuration for fleet.

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

<a id="canonical-0b1f1bd5cbeab329ef978c4443d5dab38b5026b9b2563d2e4416a6ffc261df9a"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.bond_interface / 0e2430aa3991 / 3

- [active_backup](data-sources--securemesh_site_v2--reference--group-017.md#canonical-8df02ece9b575b8458cf28b5d929b3a243f2eacdbc4f06826aa2a973321067ae): complete subsection reference.

<a id="canonical-b9c182446b55fdab634dde196f1b0819108f98051b766c4141259b72f25eeb5d"></a>

<a id="canonical-6b23979f5df85babb411cf8e4fdd4bd84750e0a6292e177669f52c3f6bc15895"></a>

## devices property — vmware.not_managed.node_list.interface_list.bond_interface / 0e2430aa3991 / 4

Type: `["list", "string"]`. Computed.

Ethernet devices that will make up this bond.

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

- [lacp](data-sources--securemesh_site_v2--reference--group-017.md#canonical-9c638cfaf9e22db98977f617beec79611eb0ba052ea49ed44417692ec630babf): complete subsection reference.

<a id="canonical-f5d8b9f3bbc0625430e73cdd64876bfd2edb7b08ad95bfc58e0032d7329ef597"></a>

<a id="canonical-cdf21dd811fbe7243ecab3921812c96279d96e5f42ee60afdb45cfec2f427168"></a>

## link_polling_interval property — vmware.not_managed.node_list.interface_list.bond_interface / 0e2430aa3991 / 5

Type: `"number"`. Computed.

Link Polling Interval. Link polling interval in milliseconds.

Upstream description:

Link polling interval in milliseconds.

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

<a id="canonical-86a82e3fca2e99a6e8d585688b66e33299914781865cd3573f09b423bd54cfad"></a>

<a id="canonical-0d212e4af1ff1e4afaa8a45c4a726d5afd2be9dc9241ea1e5eee0f50a7b00feb"></a>

## link_up_delay property — vmware.not_managed.node_list.interface_list.bond_interface / 0e2430aa3991 / 6

Type: `"number"`. Computed.

Milliseconds wait before link is declared up.

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

<a id="canonical-1c0a5f6e8dcc844c8a7bed62147d814c7808f465fb9171143facc4e2e0345a69"></a>

<a id="canonical-f0584ed711f3e431847b134896cf5774d2a8575baa6545e2e5c9f0e2d8ca6099"></a>

## name property — vmware.not_managed.node_list.interface_list.bond_interface / 0e2430aa3991 / 7

Type: `"string"`. Computed.

Bond Device Name. Name for the Bond. Ex 'bond0'

Upstream description:

Name for the Bond. Ex 'bond0'

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

<a id="canonical-3c20e135285cb23f433fe81b2895fdee8f4782b40cb94edd38ffa0f70a9b1b14"></a>

## Next pages — vmware.not_managed.node_list.interface_list.bond_interface / 0e2430aa3991 / 8

- [vmware.not_managed.node_list.interface_list.bond_interface.active_backup](data-sources--securemesh_site_v2--reference--group-017.md#canonical-8df02ece9b575b8458cf28b5d929b3a243f2eacdbc4f06826aa2a973321067ae)
- [vmware.not_managed.node_list.interface_list.bond_interface.lacp](data-sources--securemesh_site_v2--reference--group-017.md#canonical-9c638cfaf9e22db98977f617beec79611eb0ba052ea49ed44417692ec630babf)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-8df02ece9b575b8458cf28b5d929b3a243f2eacdbc4f06826aa2a973321067ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6074ad03da5a07ede3f987086ac71521c0d061e6643d146f41c27c30381372d0"></a>

## vmware.not_managed.node_list.interface_list.bond_interface.active_backup — vmware.not_managed.node_list.interface_list.bond_interface.active_backup / 797363bd6e7b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [vmware.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-d7bfdee621d5eca93bd783d969229a48f6299a1d74b3f667f66a661bdfee2133)
- vmware.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-132ef815bf97661fe0555c15015722493cce39ce751bcb6d8914f897344c802b"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2e2e61d48f6cd941d09dae0e0e7ec205d9f2dca76b1580bab365b4fede97c883"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.bond_interface.active_backup / 797363bd6e7b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fdacbae6cbaebe38301317101776bb652af56ec1e6ad43d5c32f79c7ece7e9dc"></a>

## Next pages — vmware.not_managed.node_list.interface_list.bond_interface.active_backup / 797363bd6e7b / 4

- [vmware.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-d7bfdee621d5eca93bd783d969229a48f6299a1d74b3f667f66a661bdfee2133)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-9c638cfaf9e22db98977f617beec79611eb0ba052ea49ed44417692ec630babf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b43cff767c9d2b15e9d2c46d3c94b4ddbf849a0b118e9b2cf7f083cddb2b1bc"></a>

## vmware.not_managed.node_list.interface_list.bond_interface.lacp — vmware.not_managed.node_list.interface_list.bond_interface.lacp / 2ce1faeea125 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [vmware.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-d7bfdee621d5eca93bd783d969229a48f6299a1d74b3f667f66a661bdfee2133)
- vmware.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-57466d45eec3838e72f0fdc9997023776332574a5f12e6fe46354437461295c3"></a>

Type: `"single"`. Computed.

LACP parameters. LACP parameters for the bond device.

Upstream description:

LACP parameters for the bond device.

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

<a id="canonical-074aedd71a0b7c0b82360977ac9c7891eaa224d5307a1af118cddda7528edce9"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.bond_interface.lacp / 2ce1faeea125 / 3

<a id="canonical-60b326b1ae61bed5cfae8dbf7df35d2cd39563102fa20701100a2f00583ebf75"></a>

<a id="canonical-a986ab7c8f06b9c95c94d2111e21f85eda2f90640cfa8c9b00a846b7714f1972"></a>

## rate property — vmware.not_managed.node_list.interface_list.bond_interface.lacp / 2ce1faeea125 / 4

Type: `"number"`. Computed.

Interval in seconds to transmit LACP packets.

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

<a id="canonical-47a964cdc486fef8e269869ab8a0e13c28696ac7641b983603a966bb8484ad70"></a>

## Next pages — vmware.not_managed.node_list.interface_list.bond_interface.lacp / 2ce1faeea125 / 5

- [vmware.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-d7bfdee621d5eca93bd783d969229a48f6299a1d74b3f667f66a661bdfee2133)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-820048c87747341ee92a305eb48331c1ad7170c3802f93011c1b29a15d1fb513"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f551bcc6f6b28ed146d130c61bdee91dfc32fc44ccc82bdfb3702a0f02967334"></a>

## vmware.not_managed.node_list.interface_list.dhcp_client — vmware.not_managed.node_list.interface_list.dhcp_client / 87ad0d781e7a / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- vmware.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-2e7b1245dd2c00bb514470a59c2fda899806424f35be5c86c3edf2dac19fa6ce"></a>

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

<a id="canonical-612205a7aa8939396ef804e2a70dbb712b7e7fc1d2eef0e229bb47e425ac7c67"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.dhcp_client / 87ad0d781e7a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3fee3eb29b9a3d08c8196d3f03975fd357b458462a80c9a36e3f4e236b472e4d"></a>

## Next pages — vmware.not_managed.node_list.interface_list.dhcp_client / 87ad0d781e7a / 4

- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-90db80840d6492b757411db0899cb8591bf2cbe35d58e7264e798815fd1cb85c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cdc8ca3734dfc99cd36b9f859c8d269c749afca45764b3f00459b302cc6e6f55"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server — vmware.not_managed.node_list.interface_list.dhcp_server / ff2dae5b60d5 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- vmware.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-df92c78c316ba25ccc8f6304dce29c11e4578acd215a6f4de57765620a1ba04f"></a>

Type: `"single"`. Computed.

DHCPServerParametersType.

Upstream description:

DHCP server configuration for this interface.

Receipt-pinned upstream constraints:

```json
{
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

<a id="canonical-e1130af9f1273b408995ea78d729e5a832015475fe765e486d69af25af14a9ec"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.dhcp_server / ff2dae5b60d5 / 3

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3851e06e665df7b44e50ba9ee30ecbce5588e45e6fdc0f4591eaa9fa890ce08d): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-017.md#canonical-73458f45584034ff429c4450eca753bc36c433b1c7b888f304a00e8b9e04c138): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-017.md#canonical-f6f81ca3508cc12c4b2531615e4c52b7e56a9ec2652b5f07837f1259d472e6d0): complete subsection reference.

<a id="canonical-65fd3b6ef8c56731c4411f2da458b14c89ad89026647df75ff3cca519f40a5c8"></a>

<a id="canonical-aba49f49d0f765c84a95fee1e5ff9deded541778f41df6a5b6096f1002b53075"></a>

## dhcp_option82_tag property — vmware.not_managed.node_list.interface_list.dhcp_server / ff2dae5b60d5 / 4

Type: `"string"`. Computed.

DHCP option 82 tag.

<a id="canonical-43d4a99c8ebc32c05357565f156d747e2d368e3341c2571a534233472760ac34"></a>

<a id="canonical-c78686d847bd0236b9dbd48f70fdd2ef5e5392e45c17e19e63ea2664ab453212"></a>

## fixed_ip_map property — vmware.not_managed.node_list.interface_list.dhcp_server / ff2dae5b60d5 / 5

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

- [interface_ip_map](data-sources--securemesh_site_v2--reference--group-017.md#canonical-8935cb98e0a00760412188310b5f7cf70cfd6afe15ca69909f5aaf3b752c94fd): complete subsection reference.

<a id="canonical-f8aa524789234a7050081a4fe413263ab0ae1dbc2b9559f7f3443ec5251ea8c6"></a>

## Next pages — vmware.not_managed.node_list.interface_list.dhcp_server / ff2dae5b60d5 / 6

- [vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_end](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3851e06e665df7b44e50ba9ee30ecbce5588e45e6fdc0f4591eaa9fa890ce08d)
- [vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_start](data-sources--securemesh_site_v2--reference--group-017.md#canonical-73458f45584034ff429c4450eca753bc36c433b1c7b888f304a00e8b9e04c138)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-017.md#canonical-f6f81ca3508cc12c4b2531615e4c52b7e56a9ec2652b5f07837f1259d472e6d0)
- [vmware.not_managed.node_list.interface_list.dhcp_server.interface_ip_map](data-sources--securemesh_site_v2--reference--group-017.md#canonical-8935cb98e0a00760412188310b5f7cf70cfd6afe15ca69909f5aaf3b752c94fd)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-3851e06e665df7b44e50ba9ee30ecbce5588e45e6fdc0f4591eaa9fa890ce08d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1147e2a43d8aa4a493988a4858430d4b09ce6091d5c92148580d176114315a51"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_end — vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / 23d2a80077c5 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [vmware.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-017.md#canonical-90db80840d6492b757411db0899cb8591bf2cbe35d58e7264e798815fd1cb85c)
- vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-b9f1d09dca8dc97bf0f92ab509d79195d390ad21d41f02c885adf977c07e5852"></a>

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

<a id="canonical-d5a916ab954bec218065275c4ac310b397912d868dbb42cf726fe339121afa74"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / 23d2a80077c5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d8fc4723c0637b41583f31a8f06294b893c82ebcc435307dfb57b5439eecdfb3"></a>

## Next pages — vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / 23d2a80077c5 / 4

- [vmware.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-017.md#canonical-90db80840d6492b757411db0899cb8591bf2cbe35d58e7264e798815fd1cb85c)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-73458f45584034ff429c4450eca753bc36c433b1c7b888f304a00e8b9e04c138"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-81f83350d9106b800f58be812d8a5ec5a75f5ac8a6b928841178d8cd8e580c73"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_start — vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / 5bfda3263049 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [vmware.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-017.md#canonical-90db80840d6492b757411db0899cb8591bf2cbe35d58e7264e798815fd1cb85c)
- vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-9661f754ad90c2ee4b2b84865abc0ceaa1f823239a9bd6ab15fd0fe71cd59ca0"></a>

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

<a id="canonical-21f9a5ce4e9426834788411b3c2c460f70c1bdf0e78f7e9c0ea710773911d131"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / 5bfda3263049 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-945508733766705e6d862bd59809b04a6383c446cdab9dfc37ac5a949538d653"></a>

## Next pages — vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / 5bfda3263049 / 4

- [vmware.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-017.md#canonical-90db80840d6492b757411db0899cb8591bf2cbe35d58e7264e798815fd1cb85c)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-f6f81ca3508cc12c4b2531615e4c52b7e56a9ec2652b5f07837f1259d472e6d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-049bd6df2640f87df9f648c21aed3426072dfd79eb1171b40cc8baebff410f59"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 7ce9fb4516f0 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [vmware.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-017.md#canonical-90db80840d6492b757411db0899cb8591bf2cbe35d58e7264e798815fd1cb85c)
- vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-7659d1ae6d97d8e01c0758709f282b8916e52de6b49194c0183e7ce53a634f0a"></a>

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

<a id="canonical-688b056cb62d3644d5bef5aee1b52e3d1b507ea7537fd6695c93cc499335c528"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 7ce9fb4516f0 / 3

<a id="canonical-64c1558a241426c5fd58425d0acdff1ac5aef984ab3fbeb52e934b8e2c12bb19"></a>

<a id="canonical-f7a0a4567afef121ad1e87b66027ee7e9a1025363b307f464ea15ceda9a98d59"></a>

## dgw_address property — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 7ce9fb4516f0 / 4

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

<a id="canonical-40b76b674a04fac66273908dc4ce21672dafa54b732b88ee89013dc354fc31f8"></a>

<a id="canonical-35e879f6221a867757ea4331fc97f01b6c98b47c78f6225df4d5ef65c781cc82"></a>

## dns_address property — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 7ce9fb4516f0 / 5

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

- [first_address](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1147bbf4ad0e3903cb4307335c68d7e6b983dc0189aa9b40ec063f185a869513): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-017.md#canonical-7b5399d7a3edea50da928d4c501254fd0480d6b93f55fc854be0ca5f95bf990b): complete subsection reference.

<a id="canonical-0160ed6a311c7a736b0f5ce27a8a796855d54552bd9ab5f6b85cc530b69fe1c0"></a>

<a id="canonical-342ccda09f7cb891cbb4efd02965b1ba686a99c65247384baa5f6ec8afd054ed"></a>

## network_prefix property — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 7ce9fb4516f0 / 6

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

<a id="canonical-26b14c3648e54b40bade3105d51da81c6d977dfd2d9521b16ab3e948b29bd06a"></a>

<a id="canonical-950e583a2474cfffeb871fb4f9819b1c4e2a39a2614ac3de4cdb10ee5e1a92df"></a>

## pool_settings property — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 7ce9fb4516f0 / 7

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

- [pools](data-sources--securemesh_site_v2--reference--group-017.md#canonical-527d612738332c97a6442b44b60f54000ffd3ed2be95491f99aa8ae754515b1f): complete subsection reference.

- [same_as_dgw](data-sources--securemesh_site_v2--reference--group-017.md#canonical-ef590b3ce218d0b5b2c1d744b8d05cc880e3264e6ac85bf809c140cd9a4fe4cd): complete subsection reference.

<a id="canonical-ebdc5511df2db9726b5a560c7557178f592ffb22d1e3f0f5e52a775b1e905f5d"></a>

## Next pages — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 7ce9fb4516f0 / 8

- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1147bbf4ad0e3903cb4307335c68d7e6b983dc0189aa9b40ec063f185a869513)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address](data-sources--securemesh_site_v2--reference--group-017.md#canonical-7b5399d7a3edea50da928d4c501254fd0480d6b93f55fc854be0ca5f95bf990b)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools](data-sources--securemesh_site_v2--reference--group-017.md#canonical-527d612738332c97a6442b44b60f54000ffd3ed2be95491f99aa8ae754515b1f)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw](data-sources--securemesh_site_v2--reference--group-017.md#canonical-ef590b3ce218d0b5b2c1d744b8d05cc880e3264e6ac85bf809c140cd9a4fe4cd)
- [vmware.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-017.md#canonical-90db80840d6492b757411db0899cb8591bf2cbe35d58e7264e798815fd1cb85c)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-1147bbf4ad0e3903cb4307335c68d7e6b983dc0189aa9b40ec063f185a869513"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-60ba6bdeb4a825595fbf7fd69d0a7dbefe01ebed2cbc0ed748508f872bd0de34"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_addr / cc7ed5a8d004 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [vmware.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-017.md#canonical-90db80840d6492b757411db0899cb8591bf2cbe35d58e7264e798815fd1cb85c)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-017.md#canonical-f6f81ca3508cc12c4b2531615e4c52b7e56a9ec2652b5f07837f1259d472e6d0)
- vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-62f69fd6940cb9055bb4c55304f87e5a38bac96caf911379b2b31783de9dd39f"></a>

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

<a id="canonical-27e92efd581b456b166989bebc161c9a964607b17622910b0ec477ba82d7a94c"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_addr / cc7ed5a8d004 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4805bcf7ca4d68171e77e7d7bb852757410567de08047be7610e2572a38c4cc5"></a>

## Next pages — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_addr / cc7ed5a8d004 / 4

- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-017.md#canonical-f6f81ca3508cc12c4b2531615e4c52b7e56a9ec2652b5f07837f1259d472e6d0)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-7b5399d7a3edea50da928d4c501254fd0480d6b93f55fc854be0ca5f95bf990b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e817768dd72b43a1b1bd12a366afb2023d11894fb61c292fb9ffd6fb0d530553"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_addre / 84ff52d9819d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [vmware.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-017.md#canonical-90db80840d6492b757411db0899cb8591bf2cbe35d58e7264e798815fd1cb85c)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-017.md#canonical-f6f81ca3508cc12c4b2531615e4c52b7e56a9ec2652b5f07837f1259d472e6d0)
- vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-04e733f53b717df8edd9109cdd7402c159e3908764d0c15655d116ce3778489e"></a>

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

<a id="canonical-d57b09fb8536271492e2be406a0ebf5358b61578968e6b297283280ab79feb86"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_addre / 84ff52d9819d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5210ed18fc29d19cd89667c9ed1c09c17f9906597fbba63b4567e3616843556d"></a>

## Next pages — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_addre / 84ff52d9819d / 4

- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-017.md#canonical-f6f81ca3508cc12c4b2531615e4c52b7e56a9ec2652b5f07837f1259d472e6d0)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-527d612738332c97a6442b44b60f54000ffd3ed2be95491f99aa8ae754515b1f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-776b6e1c17a5b612d508df2d232073ae78ff833ceefb15266edbb0c585927a68"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 967fb79a934b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [vmware.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-017.md#canonical-90db80840d6492b757411db0899cb8591bf2cbe35d58e7264e798815fd1cb85c)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-017.md#canonical-f6f81ca3508cc12c4b2531615e4c52b7e56a9ec2652b5f07837f1259d472e6d0)
- vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-8edf06061d60b9d5a57459923dd047b88f2eb0eb451ef3889b7710cd73c5e2b0"></a>

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

<a id="canonical-d684c7d49693839b2b991dc68425ddfef6397736a84ec077ff269fe6e9a68d90"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 967fb79a934b / 3

<a id="canonical-e63a364fd93ef4caf7eca85ebd2712f5633324ea3de6c1037f41e9e74dfcdb75"></a>

<a id="canonical-4e141876be7ea258c3d0e1a59dce87b145a44e7900508dad296dd10bf60afd9b"></a>

## end_ip property — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 967fb79a934b / 4

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

<a id="canonical-bb7a7ee357260dcc0376e5b3941fb52f6f9fb4fe9434280f04691c08801e8195"></a>

<a id="canonical-9e8950260bad2bf56f00ab04a595860e77344a43da480e28ddd8c1b122a4830d"></a>

## exclude property — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 967fb79a934b / 5

Type: `"bool"`. Computed.

Exclude this address range from DHCP allocation.

<a id="canonical-79b1a57fa3155cf8a5cf12005c9f688d8f6f31353ca2080d9720fc7c92e3e6c9"></a>

<a id="canonical-e7799b95e444b7bdc58ad13cb728eaad92a673353ccc54414f8ae650a8518f9f"></a>

## start_ip property — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 967fb79a934b / 6

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

<a id="canonical-5b72a685ab0be644eb23f97fdf04f62aa8094473ef503d83c3be08fb8bb41e9a"></a>

## Next pages — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 967fb79a934b / 7

- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-017.md#canonical-f6f81ca3508cc12c4b2531615e4c52b7e56a9ec2652b5f07837f1259d472e6d0)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-ef590b3ce218d0b5b2c1d744b8d05cc880e3264e6ac85bf809c140cd9a4fe4cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-926afad285b85696764964b8f249cdb04afa7d6a4bc9302a7111664019de7331"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dg / 8f2cee990cc5 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [vmware.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-017.md#canonical-90db80840d6492b757411db0899cb8591bf2cbe35d58e7264e798815fd1cb85c)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-017.md#canonical-f6f81ca3508cc12c4b2531615e4c52b7e56a9ec2652b5f07837f1259d472e6d0)
- vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-8baa53aaa1993c07ce9475014c6d382ec4b5e80b311fe376cb274ebd5ff7ebc5"></a>

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

<a id="canonical-b03c2c6eb9e1d56f0167f634fb49d5c3954c2b470eac13aa998aeacbe166338d"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dg / 8f2cee990cc5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d189f5346ae5b7167eb6fba0a961ce79b2293d8b213fe12ec6236804b29f2d55"></a>

## Next pages — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dg / 8f2cee990cc5 / 4

- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-017.md#canonical-f6f81ca3508cc12c4b2531615e4c52b7e56a9ec2652b5f07837f1259d472e6d0)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-8935cb98e0a00760412188310b5f7cf70cfd6afe15ca69909f5aaf3b752c94fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f6e6d75e8a2729e8a9b619ee6561fda64fd9aaa97b2687bb0b14333d8a0d26f"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.interface_ip_map — vmware.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 32369df4c104 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [vmware.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-017.md#canonical-90db80840d6492b757411db0899cb8591bf2cbe35d58e7264e798815fd1cb85c)
- vmware.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-aa875f643207e40cb1e054aa7db2b2b5e9d447e68f0fe8e9acca8e3c958ae3dd"></a>

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

<a id="canonical-0f4032d8fbcc0e5238b74115a6c71d85a4d8d2f8e055011de320fadd96da902e"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 32369df4c104 / 3

<a id="canonical-ff1244b40efc93c0ca39f4db4f6c01094839b8eb8487ed4ca771baf1cdc44355"></a>

<a id="canonical-efed05f520173b04eab2e6eecf3983d67b3ca9c86ff488562d23f1d63c033509"></a>

## interface_ip_map property — vmware.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 32369df4c104 / 4

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

<a id="canonical-6eeb1330899f24db5590aedec93a7c871fdfa038fd976b199b9a6a23f5b28eda"></a>

## Next pages — vmware.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 32369df4c104 / 5

- [vmware.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-017.md#canonical-90db80840d6492b757411db0899cb8591bf2cbe35d58e7264e798815fd1cb85c)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-9d944837d4c7972a526e05cea9804ab6f5f114eb7ef9e7670a296540362f28ef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d4a5db8bfa26dd3ddf53a6950d06e306ed9668f813203dda18bde75aeadc0c3a"></a>

## vmware.not_managed.node_list.interface_list.ethernet_interface — vmware.not_managed.node_list.interface_list.ethernet_interface / 20c5b0e7f76c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- vmware.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-25aa968e22d6e34101f44009f77bfe00728043ff4fe0be08e4469f2c40139ef1"></a>

Type: `"single"`. Computed.

Configuration parameter for ethernet interface.

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

<a id="canonical-c888e400280534e7abd87cda13586b2a272ae7b5fbd0d79422065e5557f88a25"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.ethernet_interface / 20c5b0e7f76c / 3

<a id="canonical-02d3d01e54f7eb1f18b90196eb2051e0f1ec6d1bac9460588ca633d7ee98ccb1"></a>

<a id="canonical-a62eb9f67d6d989dacb150827f48026f297519741bbca1d5c5a1071ce9dcd82b"></a>

## device property — vmware.not_managed.node_list.interface_list.ethernet_interface / 20c5b0e7f76c / 4

Type: `"string"`. Computed.

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

Upstream description:

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

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

<a id="canonical-a7b0e12680861e6c28e9bb83950ad4aec8976861bf92048262572d7ac7a61ca1"></a>

<a id="canonical-25e1cf7caafdbca36bb38a54be591010c05b28d2da1b0c0af12691b0f6676ac9"></a>

## mac property — vmware.not_managed.node_list.interface_list.ethernet_interface / 20c5b0e7f76c / 5

Type: `"string"`. Computed.

MAC Address. Configuration parameter for mac

Upstream description:

Configuration parameter for mac

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

<a id="canonical-af7c03d0c19f75e33e8e0c9e919fea6bbd3892d20607f97cfa420e79527031a1"></a>

## Next pages — vmware.not_managed.node_list.interface_list.ethernet_interface / 20c5b0e7f76c / 6

- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-77129aec4ff1f628e53657b8ee2ad737e637dabc8b2c251ad5c258f92e52a9fa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad589e43fcf86c67d73618cfcaa5c614dd1ed3ac0b0387ac35eb5d71e9a74ded"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config — vmware.not_managed.node_list.interface_list.ipv6_auto_config / c5ae580b35b4 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-4d57221498c8be40f9c9ffa21a3f6748697b048b703c1d50b34907e33ee75117"></a>

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

<a id="canonical-d2f77cb2876442236b4b90bd1223ba2d4e9f08b88ddd788d7a24a093186a93b7"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.ipv6_auto_config / c5ae580b35b4 / 3

- [host](data-sources--securemesh_site_v2--reference--group-017.md#canonical-5e1f65854a02d5ecbf4de109571e82427744bf185fb97075f71c5ae61a4b5a8f): complete subsection reference.

- [router](data-sources--securemesh_site_v2--reference--group-017.md#canonical-07acf2c5d11f4fb50860146a4a96efc2994626ae6b632b8000751a1f268fea34): complete subsection reference.

<a id="canonical-d6423036b6355e26e330cc23a9d571ad1e2ed81c86dcf8a00fe1bff3de926d43"></a>

## Next pages — vmware.not_managed.node_list.interface_list.ipv6_auto_config / c5ae580b35b4 / 4

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.host](data-sources--securemesh_site_v2--reference--group-017.md#canonical-5e1f65854a02d5ecbf4de109571e82427744bf185fb97075f71c5ae61a4b5a8f)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-017.md#canonical-07acf2c5d11f4fb50860146a4a96efc2994626ae6b632b8000751a1f268fea34)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-5e1f65854a02d5ecbf4de109571e82427744bf185fb97075f71c5ae61a4b5a8f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd76d18a0a2e8fcde007a993558559c16327ea77c743b7263099fd841e37bf17"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.host — vmware.not_managed.node_list.interface_list.ipv6_auto_config.host / d62d0140e93a / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-017.md#canonical-77129aec4ff1f628e53657b8ee2ad737e637dabc8b2c251ad5c258f92e52a9fa)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-ccc2c5c642c911df3de758b794c8d27ffafb04531d958c8cc5188a4fad9d570e"></a>

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

<a id="canonical-f11e3d26e73fa13e0b35dc3fdcbd6dc4a787928456c4762bff5502a435b09071"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.ipv6_auto_config.host / d62d0140e93a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-608a680160d7bb0b708da664f3a19fb2d19c4282ee408951c18b804b2d3e7cf4"></a>

## Next pages — vmware.not_managed.node_list.interface_list.ipv6_auto_config.host / d62d0140e93a / 4

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-017.md#canonical-77129aec4ff1f628e53657b8ee2ad737e637dabc8b2c251ad5c258f92e52a9fa)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-07acf2c5d11f4fb50860146a4a96efc2994626ae6b632b8000751a1f268fea34"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-44c011f0e7fb314f33af50af8cfbf0dff983cd3e06406e68ebf656d7e2927ac3"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router / ef676c921418 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-017.md#canonical-77129aec4ff1f628e53657b8ee2ad737e637dabc8b2c251ad5c258f92e52a9fa)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-89e5d6b564233c30bfb24e937b6f1fbd478d45a21e3422bc6a7718ee4476d249"></a>

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

<a id="canonical-b067142a70583fae9fe15296931d0e2c34db23913c473545ed3e4535c88c0485"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router / ef676c921418 / 3

- [dns_config](data-sources--securemesh_site_v2--reference--group-017.md#canonical-c5e5230b713709a027bf159ae0615941fe33d81e9c461b4b8230ffac34ecead2): complete subsection reference.

<a id="canonical-3327e8b08b686d2cdbdb6f0f9a39b046f92681e23c708e25772d7e1a6da24338"></a>

<a id="canonical-d24010b144b0fecb5a71fe83b200fb062083d8ae6ec1aaaa296db8f2516402f6"></a>

## network_prefix property — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router / ef676c921418 / 4

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

- [stateful](data-sources--securemesh_site_v2--reference--group-017.md#canonical-b37eb6693f16c3a2cfd2faa396ec234a31796bde068c6994125bee4df61e956a): complete subsection reference.

<a id="canonical-1b45bae8a4f50a3934587d8ff6791de6755745c15352c7330d90c35c776b6674"></a>

## Next pages — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router / ef676c921418 / 5

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-017.md#canonical-c5e5230b713709a027bf159ae0615941fe33d81e9c461b4b8230ffac34ecead2)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-017.md#canonical-b37eb6693f16c3a2cfd2faa396ec234a31796bde068c6994125bee4df61e956a)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-017.md#canonical-77129aec4ff1f628e53657b8ee2ad737e637dabc8b2c251ad5c258f92e52a9fa)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-c5e5230b713709a027bf159ae0615941fe33d81e9c461b4b8230ffac34ecead2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c84e2061f8ce370cec1e39d74759804b4e0f0dada1a450ef8215abfe4f9910b"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / d583ad980a49 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-017.md#canonical-77129aec4ff1f628e53657b8ee2ad737e637dabc8b2c251ad5c258f92e52a9fa)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-017.md#canonical-07acf2c5d11f4fb50860146a4a96efc2994626ae6b632b8000751a1f268fea34)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-702199146779e9d8796c09e2dd245c1c3e5518f36bb82b81b3b262319585c3fa"></a>

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

<a id="canonical-fe8310fd3d3420ce24d75ff8ca3e0cfafaec7e759f5d58dbc14930c895bb740d"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / d583ad980a49 / 3

- [configured_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-c8b1616b84dc1a5852c6610542ff901d82b6422fa2f555e4839a5d8bacd52fb2): complete subsection reference.

- [local_dns](data-sources--securemesh_site_v2--reference--group-017.md#canonical-ef171900030298d12447ed5e18f318aaed61443e6e84db6abf764d16b87887ff): complete subsection reference.

<a id="canonical-03a6775326e8abcd85f7ebb606e73d16450cc9c669d7cc7578725dea600dcc0c"></a>

## Next pages — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / d583ad980a49 / 4

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-c8b1616b84dc1a5852c6610542ff901d82b6422fa2f555e4839a5d8bacd52fb2)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-017.md#canonical-ef171900030298d12447ed5e18f318aaed61443e6e84db6abf764d16b87887ff)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-017.md#canonical-07acf2c5d11f4fb50860146a4a96efc2994626ae6b632b8000751a1f268fea34)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-c8b1616b84dc1a5852c6610542ff901d82b6422fa2f555e4839a5d8bacd52fb2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac8e8d9b8a812ef0de7c4dd1ceeecfbc8c22c3bfa1ad211e9c90783e835d275f"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.c / 9eb436ac72ce / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-017.md#canonical-77129aec4ff1f628e53657b8ee2ad737e637dabc8b2c251ad5c258f92e52a9fa)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-017.md#canonical-07acf2c5d11f4fb50860146a4a96efc2994626ae6b632b8000751a1f268fea34)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-017.md#canonical-c5e5230b713709a027bf159ae0615941fe33d81e9c461b4b8230ffac34ecead2)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-730638e9c01ef6e358efe397c4e43c389434da842d5766e3db0319e937855288"></a>

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

<a id="canonical-4136c0c0ffa9973776595f248e3e783a40aa9198799658363dbedec7daffd483"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.c / 9eb436ac72ce / 3

<a id="canonical-e2729252d4ee82984106b5e6200c8b3c6a57a4e4223b4670d6662afdf4bff8ea"></a>

<a id="canonical-bd7c1fedb3e3744f71b695e4e96328cc62755caea683ef64c36b2309d8d8a3de"></a>

## dns_list property — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.c / 9eb436ac72ce / 4

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

<a id="canonical-71ca6094c8664b9d5c842c024761d1961e961a8e1a850a782f03a4a3386b3d77"></a>

## Next pages — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.c / 9eb436ac72ce / 5

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-017.md#canonical-c5e5230b713709a027bf159ae0615941fe33d81e9c461b4b8230ffac34ecead2)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-ef171900030298d12447ed5e18f318aaed61443e6e84db6abf764d16b87887ff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-719488b17af4417694c8133e000c505117eeb6420aeb297572ab3b72a18d874a"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.l / 5eb7d2bc93da / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-017.md#canonical-77129aec4ff1f628e53657b8ee2ad737e637dabc8b2c251ad5c258f92e52a9fa)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-017.md#canonical-07acf2c5d11f4fb50860146a4a96efc2994626ae6b632b8000751a1f268fea34)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-017.md#canonical-c5e5230b713709a027bf159ae0615941fe33d81e9c461b4b8230ffac34ecead2)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-317235efc82e2d3f035cebecc3c72031ef89590a3860246c9ff5bc4a930e2410"></a>

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

<a id="canonical-6ae9ecffd1bc5dca4235ae93d29932228fe5f672251b4273083109c926145a19"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.l / 5eb7d2bc93da / 3

<a id="canonical-0dca6ef911aa4aa592dce803466808ead853f706b75091274129b7fe4144a04f"></a>

<a id="canonical-acfbda16ec229f2a7695ac6f1bbcdb07e3693b4f0a3afe13fd3c529b8eab9010"></a>

## configured_address property — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.l / 5eb7d2bc93da / 4

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

- [first_address](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0ca64d3897ab93b207bf87de2ca08167fe1733a7a342956b1143f72f459ed954): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-017.md#canonical-09418ea08cc0982221cb8215bf766726be05ec5912a6ff6bc411a7e54f70c52c): complete subsection reference.

<a id="canonical-e80301d711ae5d85e4a5abf4e10f94c172180c75fe2d812ab23842a0406a7c1a"></a>

## Next pages — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.l / 5eb7d2bc93da / 5

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0ca64d3897ab93b207bf87de2ca08167fe1733a7a342956b1143f72f459ed954)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address](data-sources--securemesh_site_v2--reference--group-017.md#canonical-09418ea08cc0982221cb8215bf766726be05ec5912a6ff6bc411a7e54f70c52c)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-017.md#canonical-c5e5230b713709a027bf159ae0615941fe33d81e9c461b4b8230ffac34ecead2)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-0ca64d3897ab93b207bf87de2ca08167fe1733a7a342956b1143f72f459ed954"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f234de55796d04afd9721b3aaf037db7f18a04b15a0cce5ec71b45ed5505f897"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.l / 4fc8ab7a51a8 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-017.md#canonical-77129aec4ff1f628e53657b8ee2ad737e637dabc8b2c251ad5c258f92e52a9fa)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-017.md#canonical-07acf2c5d11f4fb50860146a4a96efc2994626ae6b632b8000751a1f268fea34)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-017.md#canonical-c5e5230b713709a027bf159ae0615941fe33d81e9c461b4b8230ffac34ecead2)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-017.md#canonical-ef171900030298d12447ed5e18f318aaed61443e6e84db6abf764d16b87887ff)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-fccfe221afb3ad063611e4874ec6f5d6ac59f7ac4d8e1b4a055d7a3a34514524"></a>

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

<a id="canonical-63406d3bcb461d59260265a8c36385bf68c0ac63a40c95521081e345843abb44"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.l / 4fc8ab7a51a8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-237b490e810f65045604fda7d03bfdcb773b97f49417b7e7f84ac14ddf9ba8ac"></a>

## Next pages — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.l / 4fc8ab7a51a8 / 4

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-017.md#canonical-ef171900030298d12447ed5e18f318aaed61443e6e84db6abf764d16b87887ff)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-09418ea08cc0982221cb8215bf766726be05ec5912a6ff6bc411a7e54f70c52c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a2b2aad3dce22652943cd1aeb9b374262010ae9c7835659f0799faa2c72e26fc"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.l / 484ede26d9e5 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-017.md#canonical-77129aec4ff1f628e53657b8ee2ad737e637dabc8b2c251ad5c258f92e52a9fa)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-017.md#canonical-07acf2c5d11f4fb50860146a4a96efc2994626ae6b632b8000751a1f268fea34)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-017.md#canonical-c5e5230b713709a027bf159ae0615941fe33d81e9c461b4b8230ffac34ecead2)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-017.md#canonical-ef171900030298d12447ed5e18f318aaed61443e6e84db6abf764d16b87887ff)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-abb1dfcb019ea8ea6f05ed0ac9d78375a9ccac3162bf2dc1d149c78f6949a999"></a>

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

<a id="canonical-eb39521ad3016b03e28fca8ddbd0a875ce739c21c207db51115c58f2c8a33e7f"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.l / 484ede26d9e5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0419b885b1a9bd532ca2372a03fec1a0dbc9047614c50afae27cb066c89da53e"></a>

## Next pages — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.l / 484ede26d9e5 / 4

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-017.md#canonical-ef171900030298d12447ed5e18f318aaed61443e6e84db6abf764d16b87887ff)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-b37eb6693f16c3a2cfd2faa396ec234a31796bde068c6994125bee4df61e956a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a1e76e03f93c487f88cca2f7b5f4f546d69abf03fdd66e0d46d141fb994bcb3f"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 9d7daf9d15d9 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-017.md#canonical-77129aec4ff1f628e53657b8ee2ad737e637dabc8b2c251ad5c258f92e52a9fa)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-017.md#canonical-07acf2c5d11f4fb50860146a4a96efc2994626ae6b632b8000751a1f268fea34)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-854c0c31d5626e4e6ca3f5d980bfcd99102799ac3345a99ccd09683a6e9204b1"></a>

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

<a id="canonical-51af99ad67d73005391873fed2951ac698289cd3375f665f234ec2d7c54b1c58"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 9d7daf9d15d9 / 3

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-017.md#canonical-ee959231c7c461eec8c2b42ebc0fdb0386f484a838b709e678cd9bc6ddf7b237): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-017.md#canonical-7242213f56c65caf8ea14b484d0ac1f8eef2791eae679e4dece7bf655c3bdb7c): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1f309a615fa3d7c612d27db8f6c4d73070e761d175590e8afe9a6c71d36418b1): complete subsection reference.

<a id="canonical-e34ea3a786b2aa7865176b535942997d8fe8dcce1fac33de2343310723175307"></a>

<a id="canonical-3ae974cac825066fcca33bbf772d55813f7383f311b57be2774952aefbf22475"></a>

## fixed_ip_map property — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 9d7daf9d15d9 / 4

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

- [interface_ip_map](data-sources--securemesh_site_v2--reference--group-017.md#canonical-c6aa10188639ea4e465341457d908393487db96780b26ffaadcb4afda8a6574b): complete subsection reference.

<a id="canonical-130a97fec0ab59bef197d81627aafb6cb43cbadf9555c2c757fc27695169f0b7"></a>

## Next pages — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 9d7daf9d15d9 / 5

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end](data-sources--securemesh_site_v2--reference--group-017.md#canonical-ee959231c7c461eec8c2b42ebc0fdb0386f484a838b709e678cd9bc6ddf7b237)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start](data-sources--securemesh_site_v2--reference--group-017.md#canonical-7242213f56c65caf8ea14b484d0ac1f8eef2791eae679e4dece7bf655c3bdb7c)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1f309a615fa3d7c612d27db8f6c4d73070e761d175590e8afe9a6c71d36418b1)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map](data-sources--securemesh_site_v2--reference--group-017.md#canonical-c6aa10188639ea4e465341457d908393487db96780b26ffaadcb4afda8a6574b)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-017.md#canonical-07acf2c5d11f4fb50860146a4a96efc2994626ae6b632b8000751a1f268fea34)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-ee959231c7c461eec8c2b42ebc0fdb0386f484a838b709e678cd9bc6ddf7b237"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c0841b60fb8d01145d03664504a184ffe2d0f3763b35274d87ef0e1239f1865"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.aut / 18c7492b47e3 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-017.md#canonical-77129aec4ff1f628e53657b8ee2ad737e637dabc8b2c251ad5c258f92e52a9fa)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-017.md#canonical-07acf2c5d11f4fb50860146a4a96efc2994626ae6b632b8000751a1f268fea34)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-017.md#canonical-b37eb6693f16c3a2cfd2faa396ec234a31796bde068c6994125bee4df61e956a)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-e7ba93f4d08d1fea8b71d4ef369617c3167881a36954e683df19b26900b7ecac"></a>

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

<a id="canonical-5a4e3e8083bd9ea8ab00fdf38f3b5351c185b0edead22a35d701b092bb93fd11"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.aut / 18c7492b47e3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c42c7860aceecfecac4b213c053bb37e181214f83b31f96f8f124648858afefb"></a>

## Next pages — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.aut / 18c7492b47e3 / 4

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-017.md#canonical-b37eb6693f16c3a2cfd2faa396ec234a31796bde068c6994125bee4df61e956a)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-7242213f56c65caf8ea14b484d0ac1f8eef2791eae679e4dece7bf655c3bdb7c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1df9638d729568a72c108bbab8803e23b6b9eaa29ded5abcc15524e78a04b28e"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.aut / 96e3c82464a1 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-017.md#canonical-77129aec4ff1f628e53657b8ee2ad737e637dabc8b2c251ad5c258f92e52a9fa)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-017.md#canonical-07acf2c5d11f4fb50860146a4a96efc2994626ae6b632b8000751a1f268fea34)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-017.md#canonical-b37eb6693f16c3a2cfd2faa396ec234a31796bde068c6994125bee4df61e956a)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-250b63a2879d8fe84d9c0e414fc34da4b38b32fee9543784331ea687a2c17bd9"></a>

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

<a id="canonical-ccdfd32eda9ecfce48d3b8c42bfac7d4174ce0930eff6a21b5b980af2d6d52be"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.aut / 96e3c82464a1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7111863f2c8f64fcf921d21ac6563f1a4886a8749cab80a054d0b27944ddbdd0"></a>

## Next pages — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.aut / 96e3c82464a1 / 4

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-017.md#canonical-b37eb6693f16c3a2cfd2faa396ec234a31796bde068c6994125bee4df61e956a)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-1f309a615fa3d7c612d27db8f6c4d73070e761d175590e8afe9a6c71d36418b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ada99f1fb025283a4c79663c1facd9d5b6e9ede092a5ff2a101702d8ec51d6b7"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhc / bc58d3ddb360 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-017.md#canonical-77129aec4ff1f628e53657b8ee2ad737e637dabc8b2c251ad5c258f92e52a9fa)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-017.md#canonical-07acf2c5d11f4fb50860146a4a96efc2994626ae6b632b8000751a1f268fea34)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-017.md#canonical-b37eb6693f16c3a2cfd2faa396ec234a31796bde068c6994125bee4df61e956a)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-7470a3b64aa579c0817e7a905e552c5cb038db64388dd5ec927a4dbca51db84d"></a>

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

<a id="canonical-f89b0aa27ffdb25c021b8b1eb2140fb11326269a4622f3315e00652cf094feb8"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhc / bc58d3ddb360 / 3

<a id="canonical-3eed26e597415fbfa1466bed8da26f4ffe1f7a39f8aa2b52c366698b8096e0ec"></a>

<a id="canonical-f74045064916a034522f1fc3e2df9495b8d6696f45801cf82d284c9ecf80e27d"></a>

## network_prefix property — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhc / bc58d3ddb360 / 4

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

<a id="canonical-cb4c8508ca973c57c08032b6864b54a6d927c22cdbc832bf07bdcef484b60420"></a>

<a id="canonical-d858246c63044580ea2bb0ac0b90ddac18938cb26ba0a2c05d0c5f6de1972694"></a>

## pool_settings property — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhc / bc58d3ddb360 / 5

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

- [pools](data-sources--securemesh_site_v2--reference--group-017.md#canonical-d968065e87b4fda9fc5be4f29232ac21f11f8ea9e0e70b3b2c7282fc82f3d71a): complete subsection reference.

<a id="canonical-da8a09a511dedfddac46efdad2e2e690762ca33d134fd24329cc6534ae1e87d0"></a>

## Next pages — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhc / bc58d3ddb360 / 6

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools](data-sources--securemesh_site_v2--reference--group-017.md#canonical-d968065e87b4fda9fc5be4f29232ac21f11f8ea9e0e70b3b2c7282fc82f3d71a)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-017.md#canonical-b37eb6693f16c3a2cfd2faa396ec234a31796bde068c6994125bee4df61e956a)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-d968065e87b4fda9fc5be4f29232ac21f11f8ea9e0e70b3b2c7282fc82f3d71a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-18d0f76dca0d54f241efa1511540efa743a6743e8bbe238005e952c8e65f25a2"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhc / 9dc2881f9825 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-017.md#canonical-77129aec4ff1f628e53657b8ee2ad737e637dabc8b2c251ad5c258f92e52a9fa)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-017.md#canonical-07acf2c5d11f4fb50860146a4a96efc2994626ae6b632b8000751a1f268fea34)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-017.md#canonical-b37eb6693f16c3a2cfd2faa396ec234a31796bde068c6994125bee4df61e956a)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1f309a615fa3d7c612d27db8f6c4d73070e761d175590e8afe9a6c71d36418b1)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-13db9dcb93cd20de316867ec553fcd56c20cca199d5730aca06525d6b219b002"></a>

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

<a id="canonical-a2684dc9a945d1fd187ece6b0a8fedd00080b6554675c756ed7c02994daf8fa3"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhc / 9dc2881f9825 / 3

<a id="canonical-08e131d217df4182869aa7570c8f9793c4361b0886261668b695540ebddc8dab"></a>

<a id="canonical-94d2e630c4e422b29e7b0d2451443dc925820ab5e8f22d431906ff7b232225f6"></a>

## end_ip property — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhc / 9dc2881f9825 / 4

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

<a id="canonical-1cf4460f1408ed35563b679d4722b51e1bd734704f98d5ff2f917257b1fac083"></a>

<a id="canonical-d90d5efb4c84e2afc9e3a8945416c8bb67a3c053e25c9054860bd3666291c6f0"></a>

## start_ip property — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhc / 9dc2881f9825 / 5

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

<a id="canonical-a566ac3cf9bffbdbe8d9949c31e07cec240535bcefcf8fbf8967b05cee7e26aa"></a>

## Next pages — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhc / 9dc2881f9825 / 6

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1f309a615fa3d7c612d27db8f6c4d73070e761d175590e8afe9a6c71d36418b1)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-c6aa10188639ea4e465341457d908393487db96780b26ffaadcb4afda8a6574b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-10c61c82557ffe7075fb3e0d69057d23da82abc6452c382413806b502a3baf6e"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.int / a1448439af79 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-017.md#canonical-77129aec4ff1f628e53657b8ee2ad737e637dabc8b2c251ad5c258f92e52a9fa)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-017.md#canonical-07acf2c5d11f4fb50860146a4a96efc2994626ae6b632b8000751a1f268fea34)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-017.md#canonical-b37eb6693f16c3a2cfd2faa396ec234a31796bde068c6994125bee4df61e956a)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-95eec95b97fe1caf3ee33db810e1daf569d460b3d483fe28812bf5bf8b17aa87"></a>

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

<a id="canonical-ea2c318edc9581be41b02c65b3f8c3a81c323a5e851534d2ce0e8f3524a57040"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.int / a1448439af79 / 3

<a id="canonical-89c699bf5276ea24acb117d5841c34549f1f0358da45e72d488af01a3d7cdc47"></a>

<a id="canonical-acbb401546986b15384209fd030b47e9a97b1b92a3901a9cf1b1302fd154cae4"></a>

## interface_ip_map property — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.int / a1448439af79 / 4

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

<a id="canonical-807a0412186240d401df4ad20170263b20200ee5dee5b3f7c762fdc797ae9dac"></a>

## Next pages — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.int / a1448439af79 / 5

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-017.md#canonical-b37eb6693f16c3a2cfd2faa396ec234a31796bde068c6994125bee4df61e956a)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-c0d532ea954a4fd60833369da7f4607c9a2d91db72c587e740164bfc42f33f20"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9239c5d4e8302fdcccc4d2236c306f3f40ce0c197cf7b4bd44a57ca92d009596"></a>

## vmware.not_managed.node_list.interface_list.monitor — vmware.not_managed.node_list.interface_list.monitor / 17656b1e1ece / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- vmware.not_managed.node_list.interface_list.monitor

<a id="canonical-4929686c0e4668c8ab9d4a8e671c3af5ac278be813720e2bc4661debec74e64a"></a>

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

<a id="canonical-78fe8008b9761f2ca360e11e3b4045ce638f63bc4df851b7f322e00b15e58ad9"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.monitor / 17656b1e1ece / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-88c7f52fe3aaf6fc2f723cec06331abd3276baced10156b4783427e0e0bfd930"></a>

## Next pages — vmware.not_managed.node_list.interface_list.monitor / 17656b1e1ece / 4

- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-d5a180851c2a785e2cf71829e969187cac0b0f828c10ea3c977c3ca059329a21"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-88743a67737a0d0405d83f7ddfa36baa2a5d75bcbf8f264d919d5cc662c6ebf0"></a>

## vmware.not_managed.node_list.interface_list.monitor_disabled — vmware.not_managed.node_list.interface_list.monitor_disabled / a075ecf3297b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- vmware.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-cddab7e2fdd8a37dea1f777e1b16906503c5a2dc9db1cde8cbfc2479cdfca624"></a>

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

<a id="canonical-8a631d94ef672c5d9a0a04388ac9a2185127571d84dae0adb168459bfb00074c"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.monitor_disabled / a075ecf3297b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f4019efc1223509ad53fe4f6d7963266aaffd50d894d5cebe7e75023bb462351"></a>

## Next pages — vmware.not_managed.node_list.interface_list.monitor_disabled / a075ecf3297b / 4

- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-f872def9ca690acf4a7892b9854268c20f37c626269035aca2e242e8fa89f7b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d090fe093bbce799d5dfb10509381a934f00e9614619670f9e100961a23013a1"></a>

## vmware.not_managed.node_list.interface_list.network_option — vmware.not_managed.node_list.interface_list.network_option / f427dde0a9a2 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- vmware.not_managed.node_list.interface_list.network_option

<a id="canonical-00ccded25439523b3999dd1c3a45251127f8c0e5ef71859c599f1757e7647a2c"></a>

Type: `"single"`. Computed.

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

<a id="canonical-a947eca727e151ca61076c1215be3379afd8553caf6f822f46f62da3ac716397"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.network_option / f427dde0a9a2 / 3

- [site_local_inside_network](data-sources--securemesh_site_v2--reference--group-017.md#canonical-8d56c90d01f26d713783d061e8efbbd0ca04a8a8bb424cfd01e7957aec01e384): complete subsection reference.

- [site_local_network](data-sources--securemesh_site_v2--reference--group-017.md#canonical-016f4054bc3f0e4c0b0f080c00ebc916a5ed0119a682ee615215754386316fb2): complete subsection reference.

<a id="canonical-0368c3e607d8c951770942c993c2613c1763a4f26ec17e3dafe9ced8979c51e4"></a>

## Next pages — vmware.not_managed.node_list.interface_list.network_option / f427dde0a9a2 / 4

- [vmware.not_managed.node_list.interface_list.network_option.site_local_inside_network](data-sources--securemesh_site_v2--reference--group-017.md#canonical-8d56c90d01f26d713783d061e8efbbd0ca04a8a8bb424cfd01e7957aec01e384)
- [vmware.not_managed.node_list.interface_list.network_option.site_local_network](data-sources--securemesh_site_v2--reference--group-017.md#canonical-016f4054bc3f0e4c0b0f080c00ebc916a5ed0119a682ee615215754386316fb2)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-8d56c90d01f26d713783d061e8efbbd0ca04a8a8bb424cfd01e7957aec01e384"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6a93af678b89a3cab87a0ac2c93bca8db645ee8da4b4066146072009b84c6c0"></a>

## vmware.not_managed.node_list.interface_list.network_option.site_local_inside_network — vmware.not_managed.node_list.interface_list.network_option.site_local_inside_net / 3ffb45a6fe7d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [vmware.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-017.md#canonical-f872def9ca690acf4a7892b9854268c20f37c626269035aca2e242e8fa89f7b8)
- vmware.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-5b205ffacade9074d892d1bbcfc0407a6b206a3dd85dd68cb5a7ef3b3cc04456"></a>

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

<a id="canonical-d4d18537b1c66687812344cb3d394fdda826769408d534e45089ffbcf71afab3"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.network_option.site_local_inside_net / 3ffb45a6fe7d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-66f964d6bcd544caad0e2ac68d385e5cc4fd1375b72bad940773c26559e9e529"></a>

## Next pages — vmware.not_managed.node_list.interface_list.network_option.site_local_inside_net / 3ffb45a6fe7d / 4

- [vmware.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-017.md#canonical-f872def9ca690acf4a7892b9854268c20f37c626269035aca2e242e8fa89f7b8)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-016f4054bc3f0e4c0b0f080c00ebc916a5ed0119a682ee615215754386316fb2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-57fc350456254c34d5481f2f7de634f2af98b214f92b512979969620deea9784"></a>

## vmware.not_managed.node_list.interface_list.network_option.site_local_network — vmware.not_managed.node_list.interface_list.network_option.site_local_network / db844354e68e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [vmware.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-017.md#canonical-f872def9ca690acf4a7892b9854268c20f37c626269035aca2e242e8fa89f7b8)
- vmware.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-0ba7d3995bcd0616dacf9efcfc9dac65289ad4965c496c1f3edf0e53be3683a6"></a>

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

<a id="canonical-d6c0108e4b31755e3a03010cb2e96bb8c1ff7e85f7e2417c13f717dcb45a9aae"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.network_option.site_local_network / db844354e68e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-785cb366d987e206fdcfc4ba2c7070f193a961c21b83468df84aff72ba6dd65f"></a>

## Next pages — vmware.not_managed.node_list.interface_list.network_option.site_local_network / db844354e68e / 4

- [vmware.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-017.md#canonical-f872def9ca690acf4a7892b9854268c20f37c626269035aca2e242e8fa89f7b8)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-75c45b1993190ce2ad19a20d8e5b01f216e6775bce3004c8506756ec8fceb0c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df958c9b810361cc5ff8527239901d19f8d28060c7aeeaed76a7e43db4dc81e8"></a>

## vmware.not_managed.node_list.interface_list.no_ipv4_address — vmware.not_managed.node_list.interface_list.no_ipv4_address / 764446fbec97 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- vmware.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-3376fbda50f92414aedcf0da3a7c2f8681e53cfe99fb81dd300fbddc8adfb916"></a>

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

<a id="canonical-cdd69fa578fc8482a65b5313615d2e0df01f251afb07638390d00c9d250ccd87"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.no_ipv4_address / 764446fbec97 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b8c7ac9567f244409f7ace9db05d80ab9ae64b9444fbf16dc837da3eba4e6c28"></a>

## Next pages — vmware.not_managed.node_list.interface_list.no_ipv4_address / 764446fbec97 / 4

- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-a9b19516002b4330fd2f2374b11caf3cea679396fc69d21c7a0e7abceb217584"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0fad89cc8a4ab2956fabc16d8a87f1991ac0cef86674b46f750743382974e09"></a>

## vmware.not_managed.node_list.interface_list.no_ipv6_address — vmware.not_managed.node_list.interface_list.no_ipv6_address / 289874a6d5ba / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- vmware.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-db422eea15f24536d83f5eb9a694713baa6c1a8c6908927fc573088dd15b5d30"></a>

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

<a id="canonical-201e24fb216467f20a9629e98120749445f81e0d8c1b40ec05314a3cf4a5c515"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.no_ipv6_address / 289874a6d5ba / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3e4006b06da74f06316a8eb78c9861ec9552531e4d47ec53dd8432e23d6d3f4f"></a>

## Next pages — vmware.not_managed.node_list.interface_list.no_ipv6_address / 289874a6d5ba / 4

- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-9995a1d8d38c789c34d29c5ff06da51cb98796fb20dcaac83b30028c3ad6c6d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3ca205eeda7e2ea6deeee22c63d2a44c611cc18faf3e572469217c95fb06b861"></a>

## vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled — vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_ / 9868b5770ccd / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-b435a4b1ad216de9e2a81e2f4956500511a9167220942da0a081897d0f6dfa0f"></a>

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

<a id="canonical-547b4032fddb9c8c6aec47f33821a5b7014d1ddfa8903f03415b824cccc26f43"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_ / 9868b5770ccd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-df7893368bb7b1932561ba3ecc0f51ce3ffc9bdf5067f0c60322352fc4b71c03"></a>

## Next pages — vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_ / 9868b5770ccd / 4

- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-8c2dee918558431a357fa0cc5fc17d639304b7e01b5d4dee3b08c42a2d4b1283"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0d03ce26c5d9cfaba8257fbb7a15635e1e34b5a784d8e86eb9a7bc645198d245"></a>

## vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled — vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_ / 2f9a230d2b62 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-c4dfcb2fed7360c21ab5eff364f2ebf964ee7c1f95c18743e1f544475fb6f9c9"></a>

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

<a id="canonical-0e727871c9ddfb8ea82f2666ec1a05accf58def2f9360f2c3b79157386ee968f"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_ / 2f9a230d2b62 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-32cf96b600be0fa304096813c535e57cfa6c7892b8cc0a6cc570bd5a8deed09a"></a>

## Next pages — vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_ / 2f9a230d2b62 / 4

- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-3a1d19b4aaf04e2319159d92244412ad68de95e84de0a9b3380ca724ff0c0314"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e7d32c17c286b32aa42793da548527ccb4d5ef7bcabade97c4dd6e92e00c9f6"></a>

## vmware.not_managed.node_list.interface_list.static_ip — vmware.not_managed.node_list.interface_list.static_ip / 0f7021d41f01 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- vmware.not_managed.node_list.interface_list.static_ip

<a id="canonical-402b59d805312ce41145deb427df2138a9f95ae27fac8a432c615cff0aeef020"></a>

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

<a id="canonical-966c346524b2567bf0e5f3e8ef41b770bc2e85240ae34af2125404baf6bfb99f"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.static_ip / 0f7021d41f01 / 3

<a id="canonical-6b4d6fd1a2c3e5cddc5d53b46b5a95174250b4f7e0c2c3fcc17a4ee537803570"></a>

<a id="canonical-3ab38e5cc20de28a3f24f891358f876145f57a5203843565ba08aaf8bea0bf29"></a>

## default_gw property — vmware.not_managed.node_list.interface_list.static_ip / 0f7021d41f01 / 4

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

<a id="canonical-7d9fe9852cd771b7d5c74936ab1ca580e47cc153c8d5d0587ff75e5fd35ee488"></a>

<a id="canonical-a5947d05038403ac480931b789819e04615b62a8ef5fae46256205cbde280c22"></a>

## dns_server property — vmware.not_managed.node_list.interface_list.static_ip / 0f7021d41f01 / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-c53ed9c4e9dfb7042029c5a4219c29987f8f588fac7c569f5624f22769ad0d60"></a>

<a id="canonical-cdae082939882864b8f617a4b2d9088467577ec11cd253348bfe98db196fe042"></a>

## ip_address property — vmware.not_managed.node_list.interface_list.static_ip / 0f7021d41f01 / 6

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

<a id="canonical-8b1c156372a0629f9c2f864dee17b7642b048596da698ecad90bf49b1abd0277"></a>

## Next pages — vmware.not_managed.node_list.interface_list.static_ip / 0f7021d41f01 / 7

- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-119877391dc95ffeb676073bd50f4e8f0fb10a1efaaeab1e0564044c3e425117"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5fd41d62cfeeb473d8549858d89682acd408d2ce9d9111f0d5593f587123165c"></a>

## vmware.not_managed.node_list.interface_list.static_ipv6_address — vmware.not_managed.node_list.interface_list.static_ipv6_address / 6fb2ba275d9e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- vmware.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-2a6c957baab3983dbca9355ec4bc8200c32291d7a892a80a64efa29b0b779c09"></a>

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

<a id="canonical-7ebb4cfa9d15c940ccdeb36ecedf3b482e1269ab7d8e51b74fb0fc25854312bb"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.static_ipv6_address / 6fb2ba275d9e / 3

- [cluster_static_ip](data-sources--securemesh_site_v2--reference--group-017.md#canonical-7e5c5e106d4842c0e42aac1a34f32e4f51d499d467a1c83044274b8dbe7e7cb6): complete subsection reference.

- [node_static_ip](data-sources--securemesh_site_v2--reference--group-017.md#canonical-eb0425f17208340b7f39cf8eb77883b8be5232b23ff01ac48926bf3c5f5330d5): complete subsection reference.

<a id="canonical-68c1ccf1f568de2523a7ee513c340ddd73cb6f768aafb94ee4c84bb8fce1f841"></a>

## Next pages — vmware.not_managed.node_list.interface_list.static_ipv6_address / 6fb2ba275d9e / 4

- [vmware.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip](data-sources--securemesh_site_v2--reference--group-017.md#canonical-7e5c5e106d4842c0e42aac1a34f32e4f51d499d467a1c83044274b8dbe7e7cb6)
- [vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip](data-sources--securemesh_site_v2--reference--group-017.md#canonical-eb0425f17208340b7f39cf8eb77883b8be5232b23ff01ac48926bf3c5f5330d5)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-7e5c5e106d4842c0e42aac1a34f32e4f51d499d467a1c83044274b8dbe7e7cb6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a4d35e9034f393755d688588bd48a55e05fa306dcc3e57b5f128bddc98eb983d"></a>

## vmware.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip — vmware.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_i / 32bd669c21a5 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [vmware.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-017.md#canonical-119877391dc95ffeb676073bd50f4e8f0fb10a1efaaeab1e0564044c3e425117)
- vmware.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-95937acfd28a456ff4f5fef5a34711525009a0029a84ff2e06cba25fea5102f5"></a>

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

<a id="canonical-4080641ca9cfaca39cffd10ae82976aefb5a21af6015824f5d3509acbf971971"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_i / 32bd669c21a5 / 3

<a id="canonical-697668aae71f1b3aadf7f12a9f890a511718cf80fa1ae430fbefa9a1e98b79d6"></a>

<a id="canonical-96d70a52b6ccb23c57aeb147d7fd586d29320703625d291649f969ba3a6234fd"></a>

## interface_ip_map property — vmware.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_i / 32bd669c21a5 / 4

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

<a id="canonical-ddab9dcce0d72a050fbb6bf666c9ff6a132598510af631fd65bd5ec635fa7159"></a>

## Next pages — vmware.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_i / 32bd669c21a5 / 5

- [vmware.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-017.md#canonical-119877391dc95ffeb676073bd50f4e8f0fb10a1efaaeab1e0564044c3e425117)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-eb0425f17208340b7f39cf8eb77883b8be5232b23ff01ac48926bf3c5f5330d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cffa3290aecc52a62d195829197d51ebc493e46a4b0e823874f1152420ee3f45"></a>

## vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip — vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 39e2305f28d1 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [vmware.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-017.md#canonical-119877391dc95ffeb676073bd50f4e8f0fb10a1efaaeab1e0564044c3e425117)
- vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-f5a64cca541e1b2f2df27f233d60f60d07fbb446aff757ec221cdd0c12ab19b1"></a>

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

<a id="canonical-cc803b76100f012c6f3e4c4e314c1eb4277e37111c65f34b6bd99d632a8c420a"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 39e2305f28d1 / 3

<a id="canonical-c6f63cbfefe040ae4a5a2f7c3876d02b00347db501476103f736064d86be423e"></a>

<a id="canonical-6c29cca1509608df802cc8c9c4f84a4c0f0e537311868d3abc1fe0a0b49fb77b"></a>

## default_gw property — vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 39e2305f28d1 / 4

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

<a id="canonical-6d4f8599a1294610cc259912a50de3f960140a95445b40e6df4820db1a12dc46"></a>

<a id="canonical-102c115fba2c4e1a19373c619d59df5f2bfc2d3138e7d06652fa3b7aeb915f65"></a>

## dns_server property — vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 39e2305f28d1 / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-80f793c0ec586cc48c0461e0cdf924dc6280db0086f147891d80d952554853b2"></a>

<a id="canonical-bdc7a1cf122f743b97bbfa78f3bd8eb3619634c94245e41f03d14aa1ab42fc97"></a>

## ip_address property — vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 39e2305f28d1 / 6

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

<a id="canonical-441ad7b027735c4d62c28c93c042fe8e6201133c594c1988795facd4fa7bc89c"></a>

## Next pages — vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 39e2305f28d1 / 7

- [vmware.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-017.md#canonical-119877391dc95ffeb676073bd50f4e8f0fb10a1efaaeab1e0564044c3e425117)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-2c336614881f32108bf45b8d32c9338cb9eb1b6eab663f99a48f91edcef0c177"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e00f6dc6e56ae50ed67cf6568e61c9ebaf99e434626f7993eaf1c63686d5d35d"></a>

## vmware.not_managed.node_list.interface_list.vlan_interface — vmware.not_managed.node_list.interface_list.vlan_interface / 2bd4076f6357 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- vmware.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-5a7c5c10e4dd40515b41f1df1f205d69d8b673d3b5023b58a329d421892ba6e9"></a>

Type: `"single"`. Computed.

Configuration parameter for vlan interface.

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

<a id="canonical-9860970146cacc5dcded7a998b311aa60a3f37e2e805f0f417293f64442620eb"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.vlan_interface / 2bd4076f6357 / 3

<a id="canonical-c7448ef13d2dc0660f26f929e3ac4602a5684b67621ddfb3673821bf5d181bb6"></a>

<a id="canonical-0a81eb364dba276efbdaa91fbf0fe5c025bd67f7c9815ef227c1dcc28896bfb4"></a>

## device property — vmware.not_managed.node_list.interface_list.vlan_interface / 2bd4076f6357 / 4

Type: `"string"`. Computed.

Select a parent interface from the dropdown.

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

<a id="canonical-33d5f4e18afa7f860b6ba3ba41714e5984e2ab64a756178e67bd06a106f07ade"></a>

<a id="canonical-cb9445d636225054ac50ee58fdd6a93bdde38753d2496e9ef4c89f01625e9da1"></a>

## vlan_id property — vmware.not_managed.node_list.interface_list.vlan_interface / 2bd4076f6357 / 5

Type: `"number"`. Computed.

Configure the VLAN tag for this interface.

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

<a id="canonical-c244c504d536f078713c58d8ce10a796fa7bbf84ea7d796ab25a40b0c1bae81a"></a>

## Next pages — vmware.not_managed.node_list.interface_list.vlan_interface / 2bd4076f6357 / 6

- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
