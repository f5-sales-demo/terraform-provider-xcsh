---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e34a3025bd4454da4be7432e500b6fb0b71b056a59bb1ba7e32f124a7e60bb54"></a>

## Property reference — Property reference / bfd8907c5b43 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- Property reference

<a id="canonical-6f60fdef316c6c3392f25406f819db8611d72a0db38587be6bd535d37a09e75f"></a>

## Direct properties — Property reference / bfd8907c5b43 / 3

- [active_enhanced_firewall_policies](data-sources--securemesh_site_v2--reference--group-003.md#canonical-7da68363f1cd3be848e53c2426044d5fd935738a020a6e6dcb6af64f5fdccf88): complete subsection reference.

- [active_forward_proxy_policies](data-sources--securemesh_site_v2--reference--group-003.md#canonical-b6f00e1f35a1f8fbbcb4ed83a820af216ea35febbccf12baa419613d6f833353): complete subsection reference.

- [admin_user_credentials](data-sources--securemesh_site_v2--reference--group-003.md#canonical-534f99f467219e2c292db350e82de3c57db5bb821272c88aa542520837b4d39d): complete subsection reference.

<a id="canonical-b46e8223db8052d4c273c77fdcc9fd0635a49718bdbbfe44438a475daa462167"></a>

<a id="canonical-098dceb1913abbac86f88cc55687444c7dd0f17c9b7f60208020d20e461ea833"></a>

## annotations property — Property reference / bfd8907c5b43 / 4

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

- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-62d5e21dcfa4f3b673848d97041cca00221aead04682eaa3e75a721f89c53efd): complete subsection reference.

- [azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0869f169673e4654df8f600e75a61b00bf722c049e0984768f1cd9a6da4c3cd9): complete subsection reference.

- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5be45692217f4e1ad758a54b0a12fa9e69b659f93e60defa41c8dc59cbe511e3): complete subsection reference.

- [block_all_services](data-sources--securemesh_site_v2--reference--group-006.md#canonical-0a7cf87c580938dbb83972241837a48690174f115478c3f0483ef78da1c8bc51): complete subsection reference.

- [blocked_services](data-sources--securemesh_site_v2--reference--group-006.md#canonical-bdc07c6fdb88576d02aecaba3b780f257c3969a706dc14bee0ae1b77043879c0): complete subsection reference.

- [custom_proxy](data-sources--securemesh_site_v2--reference--group-006.md#canonical-df3a000486b8f3fa28eaabd732100163b19fcea69dffea5684c9b1e55e56db10): complete subsection reference.

- [custom_proxy_bypass](data-sources--securemesh_site_v2--reference--group-006.md#canonical-4d3f4e8a317a63047b869979f7e14de5e4ebbc8faaf48da0a26ca6d4d07d1d08): complete subsection reference.

- [dc_cluster_group_sli](data-sources--securemesh_site_v2--reference--group-006.md#canonical-3dfd8a56f9719e72c10f919b04b4ac0a5237eb8192c28cc4f9b3b4c05e8731c0): complete subsection reference.

- [dc_cluster_group_slo](data-sources--securemesh_site_v2--reference--group-006.md#canonical-1cd8beb30390eb9dff0352e7eaba35691cbae29677af2490983826c0247f7293): complete subsection reference.

<a id="canonical-66ff7aada31c6e01a8548c8609e5ebee86f260970784c49253835ea557d27648"></a>

<a id="canonical-ff418465c92c136131fe18476c01a2b0c6980929951f60352cba9cb78b9d4b01"></a>

## description property — Property reference / bfd8907c5b43 / 5

Type: `"string"`. Computed.

Description of the SecuremeshSiteV2.

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

- [disable_advanced_delivery](data-sources--securemesh_site_v2--reference--group-006.md#canonical-672ec91c4e0d8791a168094e3a55b37bcc49ad0407eb2b1d7e760924d8087c7a): complete subsection reference.

- [disable_ha](data-sources--securemesh_site_v2--reference--group-006.md#canonical-879e23ae61314061c8979fa698c980b6ebc7ce92552213fe62e4693c7cacf58a): complete subsection reference.

- [disable_log_anonymization](data-sources--securemesh_site_v2--reference--group-006.md#canonical-806e429bdc10c108ea7fb0bc7c4a9bf6abcbfb6acc1bca8b6147ff9d43804711): complete subsection reference.

- [disable_management_network](data-sources--securemesh_site_v2--reference--group-006.md#canonical-ec6e84c00f01885007916a5339b1bd2e894e8f79e2f84c9b5074b10559010d0f): complete subsection reference.

- [disable_url_categorization](data-sources--securemesh_site_v2--reference--group-006.md#canonical-7390627391f5fd1c7117c27a9d1df6c5aa028d6e39dd04493b74000fd77f471f): complete subsection reference.

- [dns_ntp_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-fc91c77d554a75aed311fab57b0cf31e82dbecdb9a2297a2cc41bcff7ab2c2a6): complete subsection reference.

- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb): complete subsection reference.

- [enable_advanced_delivery](data-sources--securemesh_site_v2--reference--group-008.md#canonical-d4b6f8536582fa3d6ba3ebdfac41e7a754efe97f90c964c4ae33de0f7deaa51a): complete subsection reference.

- [enable_ha](data-sources--securemesh_site_v2--reference--group-008.md#canonical-110bfe58069a4c03d94d1db4fd495fb1a79d2b474c0ddaad5f2973457b7b5b61): complete subsection reference.

- [enable_log_anonymization](data-sources--securemesh_site_v2--reference--group-008.md#canonical-bb396062a93950d63ade3ee0d4aafdb1b0d1538c84b2018f6d5de4923a32ff21): complete subsection reference.

- [enable_management_network](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5091645bfaa723812d5a8f2b667f8230bc7d521bc763a6858a94a5360e286f55): complete subsection reference.

- [enable_url_categorization](data-sources--securemesh_site_v2--reference--group-008.md#canonical-e2126e998b0f941149b8a149f21f7314dedb2d1712b4b282bc4e91af2320a305): complete subsection reference.

- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383): complete subsection reference.

- [f5_proxy](data-sources--securemesh_site_v2--reference--group-009.md#canonical-b047108d9014c60a1bf0fd07a01ddf5a301070b39565754db0ebec377758331a): complete subsection reference.

- [gcp](data-sources--securemesh_site_v2--reference--group-009.md#canonical-3229300f6c35c2676b1f294409ebc51b012bc5e57eaeba576a588d62f9bbc8f9): complete subsection reference.

<a id="canonical-7b00a383f6eb4123a3960a1bd63912eeb9195343ec90615be7a009258204cfac"></a>

<a id="canonical-698a7e66f07abb929275175ea52dfc302f5f59ef8884104e2856f69e22bb7faf"></a>

## id property — Property reference / bfd8907c5b43 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-671e616ce176e8a5200edfede43edb42df057387de83e93fc1735eb82d31614e): complete subsection reference.

<a id="canonical-245020cb2d64c880d0333b94f78b8572f9a47209005f2306d67647d0f68392c3"></a>

<a id="canonical-2abffd8b2c7267889150f74153dbd2aa4fa8558687c259a94df77494c5195878"></a>

## labels property — Property reference / bfd8907c5b43 / 7

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

- [load_balancing](data-sources--securemesh_site_v2--reference--group-011.md#canonical-19ad5681f96657661c8c290f68e21059592b727992303e34d6da3f181f75a994): complete subsection reference.

- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25): complete subsection reference.

- [log_receiver_with_net](data-sources--securemesh_site_v2--reference--group-011.md#canonical-cb521801b7b2b15af35a5400de6226f064cc01542e4f27303d5c4a4313dd55be): complete subsection reference.

- [logs_streaming_disabled](data-sources--securemesh_site_v2--reference--group-011.md#canonical-f6fc226cb76785de226517bf48defedf6623821c38740083e7a7933885135110): complete subsection reference.

<a id="canonical-c3dc0f8ec14134ca5edf85a2e71244044baada024a4727e47739713f9f7f812a"></a>

<a id="canonical-92190ef565b493ffc724f54afd278c290e7b10526f9d0d695bdbebf28c5824bc"></a>

## name property — Property reference / bfd8907c5b43 / 8

Type: `"string"`. Required.

Name of the SecuremeshSiteV2.

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

<a id="canonical-b6263fe8257f251e42f2bd146acc1381eb995a1105e0d0f6bd0968014f816028"></a>

<a id="canonical-2c43314e67a88620630b9aa112bd2960df4e20209b5f5b0f699db45540d3c0c1"></a>

## namespace property — Property reference / bfd8907c5b43 / 9

Type: `"string"`. Optional, Computed.

Namespace where the SecuremeshSiteV2 exists.

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

- [no_forward_proxy](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2ffe5fefb62142719b26f3c6c44bede4eddcfc1859f2c8d964484226f2f01664): complete subsection reference.

- [no_network_policy](data-sources--securemesh_site_v2--reference--group-012.md#canonical-b0cbebc306e4bb01028db7f9f1b73fcc9dd6db70f925d6b1550c52c7f20a2ed3): complete subsection reference.

- [no_proxy_bypass](data-sources--securemesh_site_v2--reference--group-012.md#canonical-f85fc5f62c07163ac28f9a21b912d039437da07636240823d618f05c7f426f44): complete subsection reference.

- [no_s2s_connectivity_sli](data-sources--securemesh_site_v2--reference--group-012.md#canonical-7b55f26b6ebc1a044f4722479264cd29a212320cb07e2d2798b1044085b7f504): complete subsection reference.

- [no_s2s_connectivity_slo](data-sources--securemesh_site_v2--reference--group-012.md#canonical-88f8c795251a869f39da9f92a0fd9d2ed5adf9a0bd50a712df4df0e9adb657df): complete subsection reference.

- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-afaa4ccccb7523424e4df02b4f1513c21cf529a0d8c2065d006bb7faff828456): complete subsection reference.

- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5): complete subsection reference.

- [offline_survivability_mode](data-sources--securemesh_site_v2--reference--group-014.md#canonical-4981acf5ef64e52376db78ddf0a263627c8fb719f3f2a4bf4f1e72f878daef2e): complete subsection reference.

- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-c02cb902cb11cbb24d08944bbb190d4b4f598ac07935807d790ba553815e160d): complete subsection reference.

- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063): complete subsection reference.

- [performance_enhancement_mode](data-sources--securemesh_site_v2--reference--group-016.md#canonical-64f2ae85fa48c44bbb83615acc4fbb8fadc3ac1244737deb3cd3e8c54019a396): complete subsection reference.

- [private_adn](data-sources--securemesh_site_v2--reference--group-016.md#canonical-62dc9acc1977be7eb68ba5ade693dc3233e3ac4159fd0957f9d16196dff3b2ea): complete subsection reference.

- [re_select](data-sources--securemesh_site_v2--reference--group-016.md#canonical-995c3fca9a40cf2905d5a96884867c0a1ee30217eab9d5abee4111572773623e): complete subsection reference.

- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-2c730e24052cdf5bed8649e2387ffe17dc259ad692e7eedd4a9c2f8d3fd9cbf2): complete subsection reference.

- [site_mesh_group_on_slo](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1cc3e86d92efd4ec57d162359e35a8358882499a7f7471702a4568e46798b04d): complete subsection reference.

<a id="canonical-f452c162b916cf41940b1cee0c25f5a46e803708d8304243a58162f98a5a952d"></a>

<a id="canonical-c1d20f5c0755f2a04699978129eab0dce86822055133cafe7bf5177ef6e65308"></a>

## tunnel_dead_timeout property — Property reference / bfd8907c5b43 / 10

Type: `"number"`. Computed.

Time interval, in millisec, within which any IPsec / SSL connection from the site going down is
detected. When not set (== 0), a default value of 10000 msec will be used.

Upstream description:

Time interval, in millisec, within which any IPsec / SSL connection from the site going down is
detected. When not set (== 0), a default value of 10000 msec will be used.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 180000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "180000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "180000"
  }
}
```

<a id="canonical-5e86678a4aec930e50e4bd6f86dae5da581c1715da5ce342733d52acbec4336b"></a>

<a id="canonical-0c0a123bd017f53b3c6f8296d07420a2a85599a6fedf2229cccd0acdc9484e04"></a>

## tunnel_type property — Property reference / bfd8907c5b43 / 11

Type: `"string"`. Computed.

\[Enum:
SITE\_TO\_SITE\_TUNNEL\_IPSEC\_OR\_SSL|SITE\_TO\_SITE\_TUNNEL\_IPSEC|SITE\_TO\_SITE\_TUNNEL\_SSL\]
Tunnel encapsulation to be used between sites Tunnel can operate in both IPsec and SSL, with IPsec
being preferred over SSL. Tunnel is of type IPsec Tunnel is of type SSL. Possible values are
\`SITE\_TO\_SITE\_TUNNEL\_IPSEC\_OR\_SSL\`, \`SITE\_TO\_SITE\_TUNNEL\_IPSEC\`,
\`SITE\_TO\_SITE\_TUNNEL\_SSL\`. Defaults to \`SITE\_TO\_SITE\_TUNNEL\_IPSEC\_OR\_SSL\`.

Upstream description:

Tunnel encapsulation to be used between sites

Tunnel can operate in both IPsec and SSL, with IPsec being preferred over SSL. Tunnel is of type
IPsec Tunnel is of type SSL.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_TO_SITE_TUNNEL_IPSEC_OR_SSL",
  "enum": [
    "SITE_TO_SITE_TUNNEL_IPSEC_OR_SSL",
    "SITE_TO_SITE_TUNNEL_IPSEC",
    "SITE_TO_SITE_TUNNEL_SSL"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [upgrade_settings](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0c41f12660315276db3be349044c62d8a28c6e5717bd4e1297c55801131822dd): complete subsection reference.

- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee): complete subsection reference.
