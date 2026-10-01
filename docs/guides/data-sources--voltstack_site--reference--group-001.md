---
page_title: "xcsh_voltstack_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_voltstack_site reference."
---

# xcsh_voltstack_site reference

<a id="canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a1832e2ebd30e6bf644ef483cfea81ce7edd0477601a538a56455282a4b586b"></a>

## Property reference — Property reference / dca415f1cc0e / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- Property reference

<a id="canonical-32b77d208a683d50d8761c7885ad380c4bead08ebc7edcbbac6c103ddd6b13ad"></a>

## Direct properties — Property reference / dca415f1cc0e / 3

<a id="canonical-b6560301ef60cbf56f397a35e76d447d395550c8be6abf6124df8511ca535439"></a>

<a id="canonical-aac56ef396c725f29daa52ad62b4a3c5c255bf6029268705166c630e0629ef3b"></a>

## address property — Property reference / dca415f1cc0e / 4

Type: `"string"`. Computed.

Site's geographical address that can be used to determine its latitude and longitude.

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

- [allow_all_usb](data-sources--voltstack_site--reference--group-003.md#canonical-15ca2f85166bf6f2513a872b2eb8f5a33ee38dbd1020dbba22f5ab2581ceca41): complete subsection reference.

<a id="canonical-43481634809129f115d05774252762b12f7f93bb264b9852a25cdbf7546e01d4"></a>

<a id="canonical-3f8762b19d04db1bfab9836488e50f677fddc010fc54e16112371314fa5b8f99"></a>

## annotations property — Property reference / dca415f1cc0e / 5

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

- [blocked_services](data-sources--voltstack_site--reference--group-003.md#canonical-783b4dfdca6ef33d971d2ce952c8737ad6191aa6ba84e38b40ba6613aaa25a7d): complete subsection reference.

- [bond_device_list](data-sources--voltstack_site--reference--group-003.md#canonical-a903f3c7d122d62463f14d95ae5f76c866ffebca5bd9cc4ada1a20d304c9c149): complete subsection reference.

- [coordinates](data-sources--voltstack_site--reference--group-003.md#canonical-18431d3ef94a93baf7da4e2d6c03203828618e5a3ab8bdabda22b8107b8ac6b3): complete subsection reference.

- [custom_dns](data-sources--voltstack_site--reference--group-003.md#canonical-d6e518bfc3dd33924cce2c65ca1c6336785f5dc60b930f93f30e14165b2215e1): complete subsection reference.

- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061): complete subsection reference.

- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-7af1528a6af6a1f523e053cb7663436e0917d93055eb3ad078fbe31b54766fc8): complete subsection reference.

- [default_blocked_services](data-sources--voltstack_site--reference--group-008.md#canonical-e7461357edce11f12d0bf91666f7de842d4a7a9a7de54f1d105fa715d6d0dd34): complete subsection reference.

- [default_network_config](data-sources--voltstack_site--reference--group-008.md#canonical-9db08e9b7350ea2f3b245d9c8ae5ec65a693a262eadb83b937738ffb97a9e5fc): complete subsection reference.

- [default_sriov_interface](data-sources--voltstack_site--reference--group-008.md#canonical-0c5a585d2f3a3f2fc6968a6451277b8da248e4486792f317bc8a8a20635d8762): complete subsection reference.

- [default_storage_config](data-sources--voltstack_site--reference--group-008.md#canonical-093d4cd30332a4e713564955e87e1d8cd326edd84d20a2e9d2add38a9f8ad72d): complete subsection reference.

- [deny_all_usb](data-sources--voltstack_site--reference--group-008.md#canonical-122c1c62a1d45d8e3d23caa84a6df93f5500a77608120fdc365587e23d261fd4): complete subsection reference.

<a id="canonical-a48330c8e1551f0af069bc7e82646b17dc49043f833cd99cd9e70bdb646ddf7a"></a>

<a id="canonical-37e0d4a6d301bd930bdb54cd65700ad115e5e39d617cc8103a51f89a5261a594"></a>

## description property — Property reference / dca415f1cc0e / 6

Type: `"string"`. Computed.

Description of the VoltstackSite.

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

- [disable_gpu](data-sources--voltstack_site--reference--group-008.md#canonical-877efc4bc245d7bf1e70e3ad612b9045a1d3b4f1e6d7aa79f6398a1d2e6c30c1): complete subsection reference.

- [disable_vm](data-sources--voltstack_site--reference--group-008.md#canonical-4f021921de86fc6066d7cc47cc7c0989f8efbd6c658897a7c5021c1939c60c98): complete subsection reference.

- [enable_gpu](data-sources--voltstack_site--reference--group-008.md#canonical-f12c3308ffd3dedef1d42f9a057b93a343b9f7da682291ee8a15b4586c3c913a): complete subsection reference.

- [enable_vgpu](data-sources--voltstack_site--reference--group-008.md#canonical-5df00ec180f956bffc6b4be796878e17c59d72e865c90a0a8d1ca05d429253fb): complete subsection reference.

- [enable_vm](data-sources--voltstack_site--reference--group-008.md#canonical-983e8b21e1ec1033b0103e4c7a1f367075194d4096aa8f67badbffca8bf31712): complete subsection reference.

<a id="canonical-180469419e0d533cd49d764f9358b5479e1eec261a8194cd36a42dd55b7ff7c9"></a>

<a id="canonical-464138968aeda770a5f76a27e35d5eba887c3415be4a763b4980048f73600741"></a>

## id property — Property reference / dca415f1cc0e / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [k8s_cluster](data-sources--voltstack_site--reference--group-008.md#canonical-5eb2f7e3ddad0c8098df76d0e8d3eb412d9bf62bd1de42414277d6007534e26c): complete subsection reference.

- [kubernetes_upgrade_drain](data-sources--voltstack_site--reference--group-008.md#canonical-f877c8559c460480be99ad74a5c57b8fdc9f3785845575f8fae1e50846e3995c): complete subsection reference.

<a id="canonical-a45a3b7b2f677fd2131a997b92b90a3e2f4be421889da3acbb859e70ecff302d"></a>

<a id="canonical-84b6623aa63ee20d97bbce76d6f2cd3bfea8d37c8d220e9c826a242ae63c3f85"></a>

## labels property — Property reference / dca415f1cc0e / 8

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

- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280): complete subsection reference.

- [log_receiver](data-sources--voltstack_site--reference--group-009.md#canonical-50c701645a62b74a0ef3a5b0527f69b38d4685d66f5c40ed2a742cc015b587b7): complete subsection reference.

- [logs_streaming_disabled](data-sources--voltstack_site--reference--group-009.md#canonical-50a19fa81e3567c071c56c45969bd9631a50031ab64e99f68a3981b65f12af91): complete subsection reference.

- [master_node_configuration](data-sources--voltstack_site--reference--group-009.md#canonical-43ac5fc6c50ea5ab410ae17a760b8be1308b94745f377e19781f4caa5374b05b): complete subsection reference.

<a id="canonical-5b6050092de29fd94d7f6c19cdca64a3eea3fd3682d01eb6e7ef675a06dc35b8"></a>

<a id="canonical-73d86929739df0b37fbd4434e661a885f9fbbaca2458e7443bf2bee8da8911fa"></a>

## name property — Property reference / dca415f1cc0e / 9

Type: `"string"`. Required.

Name of the VoltstackSite.

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

<a id="canonical-d407318e3a7ae5876e89e1fc413795a44f4da8bfc9f3ff516b9439f4bc659ea4"></a>

<a id="canonical-2dc15065b62040868566f59631a79b8a534d2b64bbcd6d515620a9078bb416de"></a>

## namespace property — Property reference / dca415f1cc0e / 10

Type: `"string"`. Required.

Namespace where the VoltstackSite exists.

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

- [no_bond_devices](data-sources--voltstack_site--reference--group-009.md#canonical-1821b16b75027c2ba838b2ee0ece1cbc1bc9f397cf5a263538b839a6f53f64e4): complete subsection reference.

- [no_k8s_cluster](data-sources--voltstack_site--reference--group-009.md#canonical-95766e90509223e430f27a346b3dadade1ae3ee4386baa2a652f8b2a741a573c): complete subsection reference.

- [no_local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-79cf117dd5b3f3194bd236124ef6a31556e3c1805f4c2ab48064434c1a7b2bea): complete subsection reference.

- [offline_survivability_mode](data-sources--voltstack_site--reference--group-009.md#canonical-c68631f4a2c69397e59b3b8bc1b369582b6bc95af9c269ac5b32da9a345ae28a): complete subsection reference.

- [os](data-sources--voltstack_site--reference--group-009.md#canonical-d04a7772f57dd114a68678b0130f6613df5ef337a7444cc2a60f96dd1befc6cb): complete subsection reference.

- [sriov_interfaces](data-sources--voltstack_site--reference--group-009.md#canonical-539f39e53540d3d5f2b1c575b3bbfd381fd22d7eb5aa117bd8a5660f05d81ad0): complete subsection reference.

- [sw](data-sources--voltstack_site--reference--group-009.md#canonical-dfb31ba96b2d5a71625a57f3a22d7d04c58f53983362297d97cbd10e96c73ef6): complete subsection reference.

- [usb_policy](data-sources--voltstack_site--reference--group-009.md#canonical-6b3bcfe958943ca6de6f5324892b713b60459ce829b7164c23fba924583017d1): complete subsection reference.

<a id="canonical-97a182eb47a80c6fb8830eeb3c8b45ea55e385ccbbc0ca24079a730f112e31f4"></a>

<a id="canonical-0efde9b59cc890c040b7bf84ca08d8a7b47704efce3c24dfbd8e8b935069347e"></a>

## volterra_certified_hw property — Property reference / dca415f1cc0e / 11

Type: `"string"`. Computed.

Name for generic server certified hardware to form this App Stack site.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

- [waf_signatures](data-sources--voltstack_site--reference--group-009.md#canonical-5d176cdf4bc826837f7e0219c9acc47d6a8b24b94f26d46b2b52202c2d0463e7): complete subsection reference.

<a id="canonical-ebf97a06d0e557eada351eeb9b7ebd7ecf403c0c75363de9b8c94b8a31bef33d"></a>

<a id="canonical-d2bfb4ec91a8201d8ed0753022d6b3b64a415f662516412dba769cb0206767a1"></a>

## worker_nodes property — Property reference / dca415f1cc0e / 12

Type: `["list", "string"]`. Computed.

Worker Nodes. Names of worker nodes.

Upstream description:

Names of worker nodes.

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
