---
page_title: "xcsh_fleet reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fleet reference."
---

# xcsh_fleet reference

<a id="canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e8e2e21fbfa14434de03b7e6423939dbfd4e3f764ac13cd6356d106ba4cb0414"></a>

## Property reference — Property reference / 7bbc950fdd9a / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- Property reference

<a id="canonical-7f413022c00dd044fe921e7b91a3f48aa541b8c4cbad36baa4f482229b1dc076"></a>

## Direct properties — Property reference / 7bbc950fdd9a / 3

- [allow_all_usb](data-sources--fleet--reference--group-001.md#canonical-48ce6fd2697b643d6cc6cf353a4861a2b25abf8b5d7f5aa320a99bb9769006d4): complete subsection reference.

<a id="canonical-cf16000a81490c28770f2f0368d84639f0d85f15fb1f63e4f90f2ca6275e64a5"></a>

<a id="canonical-7799d29d21e43409d13007768e288d32a3b3f88c1dc8616bd74c878a7a3fbd7d"></a>

## annotations property — Property reference / 7bbc950fdd9a / 4

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

- [blocked_services](data-sources--fleet--reference--group-001.md#canonical-831c4a8e2d5878bda105bfdecc47cf272091756efb6d21963401330750a9d5cd): complete subsection reference.

- [bond_device_list](data-sources--fleet--reference--group-002.md#canonical-65f237df10336d5db82d54cfe505146d90f609d8dfa0a04b5eb4d57217f22527): complete subsection reference.

- [dc_cluster_group](data-sources--fleet--reference--group-002.md#canonical-96db9f326526edd9bd9627855e7bea8fb3a49a23022b5ea37530be7c9fcee0ec): complete subsection reference.

- [dc_cluster_group_inside](data-sources--fleet--reference--group-002.md#canonical-459f9fc81a9ee6affc8dbdde520e801034d53cf449beca648f36288a2604774b): complete subsection reference.

- [default_config](data-sources--fleet--reference--group-002.md#canonical-6eb9aaec30d2e2cf90fa069c25811e7e00d5fa9a2c65b5052118698a69c96ae6): complete subsection reference.

- [default_sriov_interface](data-sources--fleet--reference--group-002.md#canonical-4f04b3ccf051a2819245908412155e72fb5342fe961059b254e3a05ed2e1bb22): complete subsection reference.

- [default_storage_class](data-sources--fleet--reference--group-002.md#canonical-9cca355dd9da1ba042afa381e3892933c464269ad176bfa370ab5c76359d7b89): complete subsection reference.

- [deny_all_usb](data-sources--fleet--reference--group-002.md#canonical-0a7984d4c53cce0a99415a554ce8e1658f8ac5ca16ccc67781060da9f8cdb325): complete subsection reference.

<a id="canonical-c0626524b8043f57fefb454f69545047cace523572f0d9f8272b2e4f6ba6b696"></a>

<a id="canonical-96beb38f7af21e612f7a66ae5c96151217c7b01c904a971dfe541a984c1d3f7a"></a>

## description property — Property reference / 7bbc950fdd9a / 5

Type: `"string"`. Computed.

Description of the Fleet.

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

- [device_list](data-sources--fleet--reference--group-002.md#canonical-3a034fa01f0f5be5a0c60a43005276bc329a2ca951d98dbcace4874749947c7b): complete subsection reference.

- [disable_gpu](data-sources--fleet--reference--group-002.md#canonical-dcd5dc425e3d3f5f62fa1be75e9fa17e45e1b23561e3fec274ac1b08e24fd1a1): complete subsection reference.

- [disable_log_anonymization](data-sources--fleet--reference--group-002.md#canonical-c4ac51a9f536518ad770e67b50bd98210a3ba77fa9688997d2d66fba2c9aec6e): complete subsection reference.

- [disable_vm](data-sources--fleet--reference--group-002.md#canonical-bc405b9f6c12852ac10ce65f43cec97b5028acc1848f49dbd329ab63a6fe5802): complete subsection reference.

<a id="canonical-8b9b7928f94aaf06499d26a9f8d0f745ca217e610f93afe1e3fe6b7f9794cc8b"></a>

<a id="canonical-aa5386b3140f0a9bf4f36de8e220ebd1e71a6ab81745c61a6e84115921a97eda"></a>

## enable_default_fleet_config_download property — Property reference / 7bbc950fdd9a / 6

Type: `"bool"`. Computed.

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

- [enable_gpu](data-sources--fleet--reference--group-002.md#canonical-343a0f4bbda728a37fc1135323420c28596c80d44352a60959ea457ef3585178): complete subsection reference.

- [enable_log_anonymization](data-sources--fleet--reference--group-002.md#canonical-e334ff3a795b467a7ba1700a237645e0f7c87496fd88a98d73de6044b3715e29): complete subsection reference.

- [enable_vgpu](data-sources--fleet--reference--group-002.md#canonical-4d479c596ef0edbbc7521d13615c04b2300386ef89701ddef222e8ee08ec76bb): complete subsection reference.

- [enable_vm](data-sources--fleet--reference--group-002.md#canonical-5753b2c2e22083ce920a71407a58f3aba1074413677b360bbd170dcd7c606c40): complete subsection reference.

<a id="canonical-50dd159ae81de36abeb174a1c959027028103857d3600ecbca83fad26eaa3c36"></a>

<a id="canonical-c36e4cc5421afb05744a1875ac3b91de8ca607021de7f0a2ef8667d7d26e2657"></a>

## fleet_label property — Property reference / 7bbc950fdd9a / 7

Type: `"string"`. Computed.

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

<a id="canonical-9217444b25c9876443ab838f9097d85fd0e994ca03032cd8df6fc191894dbce9"></a>

<a id="canonical-884a5902050c1613fdcdded36435e41e6190df5595e8a7b38aea3a53a0e7301c"></a>

## id property — Property reference / 7bbc950fdd9a / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

- [inside_virtual_network](data-sources--fleet--reference--group-002.md#canonical-d43e2c692c15b906ace1b7b6dc086ef0308ca562dc0c1a37c9dc2952860969e4): complete subsection reference.

- [interface_list](data-sources--fleet--reference--group-002.md#canonical-bd78cb7abb7bf49babdb2d1f8412fa641d087aaea7e289e18bb4c2765a388299): complete subsection reference.

- [kubernetes_upgrade_drain](data-sources--fleet--reference--group-002.md#canonical-928e79ae9910b9b44010c3a7162b05584deee20b397abdd43f95a1c3e6ffdfa2): complete subsection reference.

<a id="canonical-df171e753d49182f70894f585eb81d8330c0f9ad9171ae2ae224ba3f70299b5f"></a>

<a id="canonical-8fef4661e9d1e8c98199b6db53f3d08892b57fa320a5299020e807ba298a7c0d"></a>

## labels property — Property reference / 7bbc950fdd9a / 9

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

- [log_receiver](data-sources--fleet--reference--group-002.md#canonical-61c444a8907092b4fee4bb68bf308a340f1dc48a3d55f8e8a78cb08a3785585e): complete subsection reference.

- [logs_streaming_disabled](data-sources--fleet--reference--group-002.md#canonical-89d2be778f765995174cabbf84408245d7744773b9bfbf012061a0edee7a6b4f): complete subsection reference.

<a id="canonical-76cee58c7194c207a92250240ce82b492f887e9f4fa8878b1f8b14ad3b84578f"></a>

<a id="canonical-d0e95ecb450a2be0bd7bb21ab1fd43be4058ff2e7e1e7cfc0a5625a8e1a0550c"></a>

## name property — Property reference / 7bbc950fdd9a / 10

Type: `"string"`. Required.

Name of the Fleet.

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

<a id="canonical-e23ffc85f6c00fee828e425309b3cf30f41ecb2e1463682e7a52bdc00c968f17"></a>

<a id="canonical-ccc1bea4f02747da7dd66b098bc40339767f93ccc0aa918ac81139f74d44a071"></a>

## namespace property — Property reference / 7bbc950fdd9a / 11

Type: `"string"`. Required.

Namespace where the Fleet exists.

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

- [network_connectors](data-sources--fleet--reference--group-002.md#canonical-db94ada9e04151546142e0446c43da3207c69f67a8be3b5e19b2101a5e11749b): complete subsection reference.

- [network_firewall](data-sources--fleet--reference--group-002.md#canonical-f98a897e6f410bf88c46126874860b12450d8d5d9cbacb244a50731d8089ebff): complete subsection reference.

- [no_bond_devices](data-sources--fleet--reference--group-002.md#canonical-edc12fb79e78cdd39a24a98d8a96c9da8c00a3b5d4ca24930245f5d909319a5a): complete subsection reference.

- [no_dc_cluster_group](data-sources--fleet--reference--group-002.md#canonical-29d0f16c609febb8dd2f0364cd7c673dde515b6724e3e901dae8ad0e73cb9105): complete subsection reference.

- [no_storage_device](data-sources--fleet--reference--group-002.md#canonical-6e5eed396a70e6fbfffcd17a91c1e00b247560f80beb44a2adc77ceee989f19f): complete subsection reference.

- [no_storage_interfaces](data-sources--fleet--reference--group-002.md#canonical-71b24f1ab3f8e13620efb2e10f7d1d07129cf3ad57f9db0c697c30f074c23710): complete subsection reference.

- [no_storage_static_routes](data-sources--fleet--reference--group-002.md#canonical-fc4925032f82914f326f0d8f65812a7152ea81f11666e910ae18af87eca39eb3): complete subsection reference.

<a id="canonical-0e7f5c28d69a95da8165a24472d443c0b071d1e4ab3db5bfb2aacaa8571b4402"></a>

<a id="canonical-48cdcbd4d835d54f031fb78894000ff61d3f9762c5e9787fd8325bad8f7057a6"></a>

## operating_system_version property — Property reference / 7bbc950fdd9a / 12

Type: `"string"`. Computed.

Desired Operating System version that is applied to all sites that are member of the fleet. Current
Operating System version can be overridden via site config.

Upstream description:

Desired Operating System version that is applied to all sites that are member of the fleet. Current
Operating System version can be overridden via site config.

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

- [outside_virtual_network](data-sources--fleet--reference--group-002.md#canonical-fa0cf1c7719192675c358757e047a074ea7122e0eecf767d9fe5eed505c48826): complete subsection reference.

- [performance_enhancement_mode](data-sources--fleet--reference--group-002.md#canonical-df949e88be9e74f148f4cc6b648395bc2c2966c58242f5e5278b65420dc4f3a5): complete subsection reference.

- [sriov_interfaces](data-sources--fleet--reference--group-002.md#canonical-5e5180d53b4f741b97d268311fbb857c0ca039ef57d462526616beaed68774c4): complete subsection reference.

- [storage_class_list](data-sources--fleet--reference--group-002.md#canonical-8b52e5aaf9f369d1a18bdcd5a9eb805473cd4387d010bcc639d0f4e2ceac9dd9): complete subsection reference.

- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349): complete subsection reference.

- [storage_interface_list](data-sources--fleet--reference--group-004.md#canonical-7e877450b08eec03e31927e555048ebdc52fb1f21283e08c037fbbe29d5cd3e0): complete subsection reference.

- [storage_static_routes](data-sources--fleet--reference--group-004.md#canonical-9a825351348828d7439fc399c86956cc54202bb679948c2a6cb2e089de057870): complete subsection reference.

- [usb_policy](data-sources--fleet--reference--group-004.md#canonical-9ec18332905de0a2ace95545b134c28e3348331c9500cac4450e177bac5b0c17): complete subsection reference.

<a id="canonical-44044520d8e1ed53c4d629dc70427b698332e06e710fe153b43c59c4e3c5b65a"></a>

<a id="canonical-131c60d61db4d79f27ac959f2ff450ec9f43f52c71cbe13f4755de65746da2b9"></a>

## volterra_software_version property — Property reference / 7bbc950fdd9a / 13

Type: `"string"`. Computed.

F5XC software version is human readable string matching released set of version components. The
given software version is applied to all sites that are member of the fleet. Current software
installed can be overridden via site config.

Upstream description:

F5XC software version is human readable string matching released set of version components. The
given software version is applied to all sites that are member of the fleet. Current software
installed can be overridden via site config.

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

<a id="canonical-7d908bf3dfffe6b71c90b7d9fd2aeb2cd887bdac90035d03342d10e99496ae5b"></a>

## All schema paths — Property reference / 7bbc950fdd9a / 14

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allow_all_usb` | [allow_all_usb](data-sources--fleet--reference--group-001.md#canonical-894771c74fed127f34d6ee0a9d326e25f4cf9e34109cb73a14c868b69bb9dc3c) |
| `annotations` | [annotations](data-sources--fleet--reference--group-001.md#canonical-cf16000a81490c28770f2f0368d84639f0d85f15fb1f63e4f90f2ca6275e64a5) |
| `blocked_services` | [blocked_services](data-sources--fleet--reference--group-001.md#canonical-9c9620bdeab71579df50c1df124d492d42087ebc3300d493286618bbb688c865) |
| `blocked_services.dns` | [blocked_services.dns](data-sources--fleet--reference--group-001.md#canonical-352e08899acf3e9a2b389b7308292eee56fb665089439e565f7f183fc2a57a9e) |
| `blocked_services.network_type` | [blocked_services.network_type](data-sources--fleet--reference--group-001.md#canonical-98d7b640f7c272a6f60cf6f2871d6c31d46d60753eafb9fd2d904f0143d7bbb4) |
| `blocked_services.ssh` | [blocked_services.ssh](data-sources--fleet--reference--group-002.md#canonical-b113ce17129aa054b609b7eedafb7ddde4339ca736d5df95e2c6375768c7bcfc) |
| `blocked_services.web_user_interface` | [blocked_services.web_user_interface](data-sources--fleet--reference--group-002.md#canonical-6ec070c26386fefe2b266ec62e8438e27614ed9aba3884328d74b3cf62020cc2) |
| `bond_device_list` | [bond_device_list](data-sources--fleet--reference--group-002.md#canonical-95f10afdf0729a2edb8d1ad7ad5cf812fb62b76b14e27bcaeb0485cd561f61b1) |
| `bond_device_list.bond_devices` | [bond_device_list.bond_devices](data-sources--fleet--reference--group-002.md#canonical-4e9011520cfe2ce50cd09d07dfcfe1c6f6fa8b2d2df4371def69b8c82cae1292) |
| `bond_device_list.bond_devices.active_backup` | [bond_device_list.bond_devices.active_backup](data-sources--fleet--reference--group-002.md#canonical-331823a1ff94f4a2196258fc1ad29963c022488d00a0568ceea91babd2d853b0) |
| `bond_device_list.bond_devices.devices` | [bond_device_list.bond_devices.devices](data-sources--fleet--reference--group-002.md#canonical-039f535ae5585e5702cdc1522d2f9e55b81a4191360c10df07bc426c37020eea) |
| `bond_device_list.bond_devices.lacp` | [bond_device_list.bond_devices.lacp](data-sources--fleet--reference--group-002.md#canonical-dd120c82bc05b923d79cfbf2a6eff615a58fe24ebf6c42b9a1da6f06ae6a748d) |
| `bond_device_list.bond_devices.lacp.rate` | [bond_device_list.bond_devices.lacp.rate](data-sources--fleet--reference--group-002.md#canonical-7b2d3c990ef608f6be0cd2b43c62e8a0e15ab89376963776c571c10840c898b2) |
| `bond_device_list.bond_devices.link_polling_interval` | [bond_device_list.bond_devices.link_polling_interval](data-sources--fleet--reference--group-002.md#canonical-71044f2452c7173e6e62899f06c59ef6e7218dc55a1b7563fe8598281a9520dd) |
| `bond_device_list.bond_devices.link_up_delay` | [bond_device_list.bond_devices.link_up_delay](data-sources--fleet--reference--group-002.md#canonical-5af99059f44569fa8637012554869a7ccd2cd1254b7187b58032036315adf644) |
| `bond_device_list.bond_devices.name` | [bond_device_list.bond_devices.name](data-sources--fleet--reference--group-002.md#canonical-adaff2482ed781c42bb7bdee2d9d668b43b9cac66759174ce48b03a06c58229d) |
| `dc_cluster_group` | [dc_cluster_group](data-sources--fleet--reference--group-002.md#canonical-7b48dc52b22d9ff2c33f7c41e0aaefd6e93cae30a75b3596d543dad7ca1d0944) |
| `dc_cluster_group.name` | [dc_cluster_group.name](data-sources--fleet--reference--group-002.md#canonical-d182dea80c51019a5a3322cf909ea8fac7515bef0c1fd9a187a38256a506f522) |
| `dc_cluster_group.namespace` | [dc_cluster_group.namespace](data-sources--fleet--reference--group-002.md#canonical-550ab3d0c40b6bd75115fbf6f5eae3286cfed345d342a322e9b28c391f6b5d59) |
| `dc_cluster_group.tenant` | [dc_cluster_group.tenant](data-sources--fleet--reference--group-002.md#canonical-e0fb9d1157089bcbe879beea9451c217e9969169457bde4d661d8257cffcfaf5) |
| `dc_cluster_group_inside` | [dc_cluster_group_inside](data-sources--fleet--reference--group-002.md#canonical-387378186aeddb5699eb3c4a7a4d576c66dfbad91e85759c2c799c87cb6c8bc5) |
| `dc_cluster_group_inside.name` | [dc_cluster_group_inside.name](data-sources--fleet--reference--group-002.md#canonical-e84b7e9dad644b889dbb5b2c696009b4c1eb19b090a822a6cab23748c96117cb) |
| `dc_cluster_group_inside.namespace` | [dc_cluster_group_inside.namespace](data-sources--fleet--reference--group-002.md#canonical-14f91a70d391315c222582429a9fdbab01da1264813924f52ab507c5cee314cc) |
| `dc_cluster_group_inside.tenant` | [dc_cluster_group_inside.tenant](data-sources--fleet--reference--group-002.md#canonical-1dbf1d404bb3f37a889a7fb79d1d9c0076aa0ad8b09f89d734c10b88204b71b1) |
| `default_config` | [default_config](data-sources--fleet--reference--group-002.md#canonical-db4590aa0b595beb88d1f0f29029700788ec22fc164f1e0277cd33b844257cd7) |
| `default_sriov_interface` | [default_sriov_interface](data-sources--fleet--reference--group-002.md#canonical-7d7255411a7ba57e08c8e014a5fe557de87b487c3c5d5507a360750e1115a7ef) |
| `default_storage_class` | [default_storage_class](data-sources--fleet--reference--group-002.md#canonical-6cdcad478f70fbb3d7ee1aa3d4f66e089295b3a973514faeaa9660673adb8639) |
| `deny_all_usb` | [deny_all_usb](data-sources--fleet--reference--group-002.md#canonical-3b86c27fe792b4846aef4b2443130cf2d32d8cfbfd2558b862ab0512f1e84b2e) |
| `description` | [description](data-sources--fleet--reference--group-001.md#canonical-c0626524b8043f57fefb454f69545047cace523572f0d9f8272b2e4f6ba6b696) |
| `device_list` | [device_list](data-sources--fleet--reference--group-002.md#canonical-85c248d666a65fa0e72ab3ac00277a5492212f5d1c072ddce591bd24611c26c7) |
| `device_list.devices` | [device_list.devices](data-sources--fleet--reference--group-002.md#canonical-a2394979b04d6df94b2caec9319d727e0a07cc539072df30f6ea242508e755b6) |
| `device_list.devices.name` | [device_list.devices.name](data-sources--fleet--reference--group-002.md#canonical-e06cd5e04105eaa4259c6d16261b6704b52fe618334cadeb2d649b9a7ddfe204) |
| `device_list.devices.network_device` | [device_list.devices.network_device](data-sources--fleet--reference--group-002.md#canonical-38c220ce32b8ccace367a51bda33e41c6a6cd7190765ce67926374c2528f4916) |
| `device_list.devices.network_device.interface` | [device_list.devices.network_device.interface](data-sources--fleet--reference--group-002.md#canonical-9657f8296b11acf7c70083918b3341422d19561e10044bb75c293f0c028f9cd9) |
| `device_list.devices.network_device.interface.kind` | [device_list.devices.network_device.interface.kind](data-sources--fleet--reference--group-002.md#canonical-73e7387bbf3f8b1965a3c257cb4699872128f898853ff9a0d9a2c7858e79e32c) |
| `device_list.devices.network_device.interface.name` | [device_list.devices.network_device.interface.name](data-sources--fleet--reference--group-002.md#canonical-b3ca19f1d971f5ae7ad7f4ef85cfb06396877494722310492825151d9ba6b2d9) |
| `device_list.devices.network_device.interface.namespace` | [device_list.devices.network_device.interface.namespace](data-sources--fleet--reference--group-002.md#canonical-408b35f2ebbf7443b3e0895734eba01cfeab4f73a8e20c16a20d29fe36c33043) |
| `device_list.devices.network_device.interface.tenant` | [device_list.devices.network_device.interface.tenant](data-sources--fleet--reference--group-002.md#canonical-347eaaf397d8cdd73f14334d7df5ddb73e59b2c23f6ad91699f42d3a705b1881) |
| `device_list.devices.network_device.interface.uid` | [device_list.devices.network_device.interface.uid](data-sources--fleet--reference--group-002.md#canonical-dc822b32b432ab66c4f7a6187c54aced32c065d60ea2c9566b700f0e4f075560) |
| `device_list.devices.network_device.use` | [device_list.devices.network_device.use](data-sources--fleet--reference--group-002.md#canonical-0c2b10e7a4186c0e06115de0d7f3fd71f9b2ad999969c79645f247efd842f677) |
| `device_list.devices.owner` | [device_list.devices.owner](data-sources--fleet--reference--group-002.md#canonical-7bbe533eeb4c156e14417f68a71e22f24ef2ee975bf847b776fe226767241f08) |
| `disable_gpu` | [disable_gpu](data-sources--fleet--reference--group-002.md#canonical-393756cd006dd757163c2f00cccc333b1d6d760e212fad8bb8ffc46f4d4c0441) |
| `disable_log_anonymization` | [disable_log_anonymization](data-sources--fleet--reference--group-002.md#canonical-913cef78e0cea3a4ccd6bab49defebc426f5280b5f39eea92a19e63be7264686) |
| `disable_vm` | [disable_vm](data-sources--fleet--reference--group-002.md#canonical-862540541fba4ea77bae1e6557d170c35418d1c127193d2cc31221de28a96894) |
| `enable_default_fleet_config_download` | [enable_default_fleet_config_download](data-sources--fleet--reference--group-001.md#canonical-8b9b7928f94aaf06499d26a9f8d0f745ca217e610f93afe1e3fe6b7f9794cc8b) |
| `enable_gpu` | [enable_gpu](data-sources--fleet--reference--group-002.md#canonical-5be198495b980d5e822da4754033850b82f036e0c01e61548e5bff613f17b879) |
| `enable_log_anonymization` | [enable_log_anonymization](data-sources--fleet--reference--group-002.md#canonical-79e4eb09a23c52e9316b620926c72524361bf3a1d6213f70d381d8eefa9f3813) |
| `enable_vgpu` | [enable_vgpu](data-sources--fleet--reference--group-002.md#canonical-b0ccd661711791bd57ec0908b512e25ec87fdb9b01e6fa7d4bcf01efd5ea8c54) |
| `enable_vgpu.feature_type` | [enable_vgpu.feature_type](data-sources--fleet--reference--group-002.md#canonical-2d5b049935bea68a6bc1c57f9e9dd8e3e386a39e976071ffb7f9442756f7ae6d) |
| `enable_vgpu.server_address` | [enable_vgpu.server_address](data-sources--fleet--reference--group-002.md#canonical-c296508a4fc0a20a3731c605a143ebfc3ff1072a988570b74de6dea6d29f1ad5) |
| `enable_vgpu.server_port` | [enable_vgpu.server_port](data-sources--fleet--reference--group-002.md#canonical-24edcc406faa62587045230608a75227a3fab1d8c9c19f16d790a4810c27d0fd) |
| `enable_vm` | [enable_vm](data-sources--fleet--reference--group-002.md#canonical-912d79f3f6bd8cfeecae4a848be82971496d2629996fde612215034b5b04690f) |
| `fleet_label` | [fleet_label](data-sources--fleet--reference--group-001.md#canonical-50dd159ae81de36abeb174a1c959027028103857d3600ecbca83fad26eaa3c36) |
| `id` | [id](data-sources--fleet--reference--group-001.md#canonical-9217444b25c9876443ab838f9097d85fd0e994ca03032cd8df6fc191894dbce9) |
| `inside_virtual_network` | [inside_virtual_network](data-sources--fleet--reference--group-002.md#canonical-31e7d83fae62e45a0fbca4d77c619fcb9ecd5f737807e27bd2328ae37fcd306b) |
| `inside_virtual_network.kind` | [inside_virtual_network.kind](data-sources--fleet--reference--group-002.md#canonical-e4e3043a98c2443ab0f656bed1c16cbb04a0a7c0762ba0ad5903cdb6eb483a53) |
| `inside_virtual_network.name` | [inside_virtual_network.name](data-sources--fleet--reference--group-002.md#canonical-10947bd758df6ae2c3f5d43fd2b626c18f016f99978828044d82523bec3f9881) |
| `inside_virtual_network.namespace` | [inside_virtual_network.namespace](data-sources--fleet--reference--group-002.md#canonical-b4de7cc81539405dd5bbb49ba6c5c71361debcd7907c375fedaf692e465b73e4) |
| `inside_virtual_network.tenant` | [inside_virtual_network.tenant](data-sources--fleet--reference--group-002.md#canonical-4234d9dd427693c550ddfe3beddb38f022d985e1674454f081028cd9302df6c8) |
| `inside_virtual_network.uid` | [inside_virtual_network.uid](data-sources--fleet--reference--group-002.md#canonical-e7befca8dd0ebfc8b042d43e20da624cda5bd38544bf26f4fdf68c76eb37a4e2) |
| `interface_list` | [interface_list](data-sources--fleet--reference--group-002.md#canonical-594053051684997c2e6c782707798cf5250b26403b0c7c1a169d902f57f4480c) |
| `interface_list.interfaces` | [interface_list.interfaces](data-sources--fleet--reference--group-002.md#canonical-2f1fc8ac938e9b2118c188dc2a5def135b1133048e382f48665b111b0b44b4a3) |
| `interface_list.interfaces.name` | [interface_list.interfaces.name](data-sources--fleet--reference--group-002.md#canonical-80cd733deba4533fac1477658e64b8bf440f33752279634772501349aeda8ae3) |
| `interface_list.interfaces.namespace` | [interface_list.interfaces.namespace](data-sources--fleet--reference--group-002.md#canonical-13203b3f2bfba7149de8e5acd8a77fbffdec76418b17391289f8dd09fd92b758) |
| `interface_list.interfaces.tenant` | [interface_list.interfaces.tenant](data-sources--fleet--reference--group-002.md#canonical-b1e30de317c5dee0024047577140a37b2a1f55e296516ded513aa8bdf3dbc13c) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](data-sources--fleet--reference--group-002.md#canonical-0fa93d6271d5b788989a566713bacafcc9c4102c6f9fa3684d8c315ae3b26c77) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](data-sources--fleet--reference--group-002.md#canonical-f6385c702cf8d87bb8016f77dcc8488c8447fb11b013072b1415b60d0d40c96c) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--fleet--reference--group-002.md#canonical-a58f5c88002b868b3c4d45801068e682d05665ade4c79e35388bf25aa4b7febe) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](data-sources--fleet--reference--group-002.md#canonical-a08760b14ffd25671a41b1fd0e8e57ae71249ed6ca20dd9be42041b7850464a1) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](data-sources--fleet--reference--group-002.md#canonical-36108684040d0d42e1b99835f391c764de1afc058389fc6e6e118f2b5c4fc3be) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](data-sources--fleet--reference--group-002.md#canonical-5f13addbfc546f2ecf9a748282c97894963b599c8b50f48704008507fac3dc75) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](data-sources--fleet--reference--group-002.md#canonical-67c7123deab5a94c1a55eb918c4dbbcf9ccc97ff910a4a08ca00654b73e61f1c) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](data-sources--fleet--reference--group-002.md#canonical-204c60eedea860abc40f2f424e7af62e72f45f50ad78c0dc7b7e200a6f0a150f) |
| `labels` | [labels](data-sources--fleet--reference--group-001.md#canonical-df171e753d49182f70894f585eb81d8330c0f9ad9171ae2ae224ba3f70299b5f) |
| `log_receiver` | [log_receiver](data-sources--fleet--reference--group-002.md#canonical-f67eefa0dd1bcbe9f2e542656fa20dcdebda24a773e224ef3a7aa78dbd216dc4) |
| `log_receiver.name` | [log_receiver.name](data-sources--fleet--reference--group-002.md#canonical-4eaa5f2ac519209a282f363b347f1b179cace5fc425dcd969eac3dc039dcec70) |
| `log_receiver.namespace` | [log_receiver.namespace](data-sources--fleet--reference--group-002.md#canonical-17f501f9077c7d7f0162b7d86d3acecf5047a88b1e120fbe18675395ffd14711) |
| `log_receiver.tenant` | [log_receiver.tenant](data-sources--fleet--reference--group-002.md#canonical-1fd2ad1937133b0080eb2df432fc2b8d704cd249e440d72767be84882b4b35d1) |
| `logs_streaming_disabled` | [logs_streaming_disabled](data-sources--fleet--reference--group-002.md#canonical-148d514c786d7ba9cc8505a2a3db6fa5ba8fd042df08017eb7cb7fd9337eccad) |
| `name` | [name](data-sources--fleet--reference--group-001.md#canonical-76cee58c7194c207a92250240ce82b492f887e9f4fa8878b1f8b14ad3b84578f) |
| `namespace` | [namespace](data-sources--fleet--reference--group-001.md#canonical-e23ffc85f6c00fee828e425309b3cf30f41ecb2e1463682e7a52bdc00c968f17) |
| `network_connectors` | [network_connectors](data-sources--fleet--reference--group-002.md#canonical-e322aa9c968bf594c7ab07a8e0f8471c4486ef776b3f3ea49f0eb8437580ccd2) |
| `network_connectors.kind` | [network_connectors.kind](data-sources--fleet--reference--group-002.md#canonical-4036ba08a83c42e88d14e53ab2df16a25d78c0f11ea0212396db094db20e5880) |
| `network_connectors.name` | [network_connectors.name](data-sources--fleet--reference--group-002.md#canonical-44a773db397c304808f1d9f52611d4766276c8cb0b76b6e6d83b8e4f3f43e5e7) |
| `network_connectors.namespace` | [network_connectors.namespace](data-sources--fleet--reference--group-002.md#canonical-90db1b6b61977593a129d7ecd6b353aabec70ba0d5fe06367401c685855e15d2) |
| `network_connectors.tenant` | [network_connectors.tenant](data-sources--fleet--reference--group-002.md#canonical-5486cbe191c942eb83012a39c437d1028aa9bb9696ee1d68d0df9e908e0b591d) |
| `network_connectors.uid` | [network_connectors.uid](data-sources--fleet--reference--group-002.md#canonical-2a479dde038c21186d2c52d7b9e1f24354c660a6fbcfc00fd774603540cba7d6) |
| `network_firewall` | [network_firewall](data-sources--fleet--reference--group-002.md#canonical-3b08b7d0e75da2e6a445046ebc3805bd0e2eb7b6263852503213dcf1db93a19e) |
| `network_firewall.kind` | [network_firewall.kind](data-sources--fleet--reference--group-002.md#canonical-9ed9962d833a0a49dacf25ba2be7c31e9d05d7fc434f8a4894bf5a8b53ecb19d) |
| `network_firewall.name` | [network_firewall.name](data-sources--fleet--reference--group-002.md#canonical-d1e7843fb1cdcbe4c0d2cb90a2fee915c0b6f62fd1e866dd8ef5b02a657db8c1) |
| `network_firewall.namespace` | [network_firewall.namespace](data-sources--fleet--reference--group-002.md#canonical-33889ad7f1aa1f635290b0ba2b6707e8af1c82edb79123bba9036b3d6a4b67cf) |
| `network_firewall.tenant` | [network_firewall.tenant](data-sources--fleet--reference--group-002.md#canonical-66586e060d2745b968cc482e3ef3d9a9f96465d937e81529cbcba69a4f457a03) |
| `network_firewall.uid` | [network_firewall.uid](data-sources--fleet--reference--group-002.md#canonical-bfa97610430e3fe9d0fdf0029b28f1ea34f3f89d5d0589075fede92cba419f02) |
| `no_bond_devices` | [no_bond_devices](data-sources--fleet--reference--group-002.md#canonical-cbb77bea9b49473d1d275bf12e90429d9044118a6cbeb6e66ddb179594e19856) |
| `no_dc_cluster_group` | [no_dc_cluster_group](data-sources--fleet--reference--group-002.md#canonical-502ef3f94d885d9be80379bbe1e385f831565b7cc46d9da0b589e2ba71b1592a) |
| `no_storage_device` | [no_storage_device](data-sources--fleet--reference--group-002.md#canonical-26283880a4a76e904bb41755a8eb196ee8a833a13fb4b4a659d0880977356606) |
| `no_storage_interfaces` | [no_storage_interfaces](data-sources--fleet--reference--group-002.md#canonical-3afaa627c2c2f35365d345a07a31bb2887c1e6d5cdfd77dbb3e4c307fb1f6c2d) |
| `no_storage_static_routes` | [no_storage_static_routes](data-sources--fleet--reference--group-002.md#canonical-552be8dea02b1461e1607cd17c3ee16871243d7c0bb412091f94e775dcf3a1ed) |
| `operating_system_version` | [operating_system_version](data-sources--fleet--reference--group-001.md#canonical-0e7f5c28d69a95da8165a24472d443c0b071d1e4ab3db5bfb2aacaa8571b4402) |
| `outside_virtual_network` | [outside_virtual_network](data-sources--fleet--reference--group-002.md#canonical-50fa72c7f19798f19eb27a0d646d0dc3483d1cd2e7f384cb57167b3e820bced2) |
| `outside_virtual_network.kind` | [outside_virtual_network.kind](data-sources--fleet--reference--group-002.md#canonical-4efcc1cb8bc411325c6b55afb31e0322c20b19fb5df128a6eb7351b946558c07) |
| `outside_virtual_network.name` | [outside_virtual_network.name](data-sources--fleet--reference--group-002.md#canonical-9f86050bbc9c08bf01471bc9e724f4d6a01b9304d64911d30060f31d9c9c9b6a) |
| `outside_virtual_network.namespace` | [outside_virtual_network.namespace](data-sources--fleet--reference--group-002.md#canonical-ebed47af51a49765853f3848b5a5f230358fec0891c6536af1cbbe3b3d3a501b) |
| `outside_virtual_network.tenant` | [outside_virtual_network.tenant](data-sources--fleet--reference--group-002.md#canonical-e3f87fc9c8c6874fb87dee60b02023e416dc957b08be2fecf8e3ddfc924da6ea) |
| `outside_virtual_network.uid` | [outside_virtual_network.uid](data-sources--fleet--reference--group-002.md#canonical-ef63dfb797ad7fa6bf2f1b41a62c6c40da6e63c166dfe3247d2fc99ffc7ecace) |
| `performance_enhancement_mode` | [performance_enhancement_mode](data-sources--fleet--reference--group-002.md#canonical-9555ce75b046317de8d4660d695c6362b6dc9c0c9a4b7d9642d09dbba1871ac8) |
| `performance_enhancement_mode.perf_mode_l3_enhanced` | [performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--fleet--reference--group-002.md#canonical-646fe04b053b30994f39e1897c284f45f07024505a73d4de9b19fd2f4de35be0) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--fleet--reference--group-002.md#canonical-9c3704a7afdc646167496993a168fc0c4fad3138aa50327df13c24d61f687d0c) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--fleet--reference--group-002.md#canonical-f30cdfb356906d5e17dfa091116eecf2118d4faa78b3ba381609eaec43ebbcd1) |
| `performance_enhancement_mode.perf_mode_l7_enhanced` | [performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--fleet--reference--group-002.md#canonical-0dcca5c40a9f3fe735bac2cbaf2577547c3ba09d5f4a999f74b261e97311d627) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--fleet--reference--group-002.md#canonical-1c520543b208e45297d238967083d4fbaa88c2bb8f9c91c3f5b474ad1efb3162) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--fleet--reference--group-002.md#canonical-6ca49af97c12a91da54c1959bd714c84ddf606b3973e516e4a691b2cde225161) |
| `sriov_interfaces` | [sriov_interfaces](data-sources--fleet--reference--group-002.md#canonical-9bffb1f0147d28386e92cce8c50f6ea0ca43a02635671e42a84e3818ab038a24) |
| `sriov_interfaces.sriov_interface` | [sriov_interfaces.sriov_interface](data-sources--fleet--reference--group-002.md#canonical-16a907d88c47380697ff3f5bd62744c056ed9a9a52649b842fddf44d59d9a43f) |
| `sriov_interfaces.sriov_interface.interface_name` | [sriov_interfaces.sriov_interface.interface_name](data-sources--fleet--reference--group-002.md#canonical-34a482eb59eddd82f55bba64dc7db949d6fc516437a86867f306e59ce9e2eddc) |
| `sriov_interfaces.sriov_interface.number_of_vfio_vfs` | [sriov_interfaces.sriov_interface.number_of_vfio_vfs](data-sources--fleet--reference--group-002.md#canonical-9502bc0afd9a1eb3eedf331772545902d391cb75aeed87b2d8a2fc06b482bfb7) |
| `sriov_interfaces.sriov_interface.number_of_vfs` | [sriov_interfaces.sriov_interface.number_of_vfs](data-sources--fleet--reference--group-002.md#canonical-4b4b204dc369ed1e98f0ffbedb68fdfc30d16fc0a1f229f933d4740af6ceb218) |
| `storage_class_list` | [storage_class_list](data-sources--fleet--reference--group-002.md#canonical-46c93d8afd7a809458a19a0b00d530da1508df6540080255d13292b61c7076cb) |
| `storage_class_list.storage_classes` | [storage_class_list.storage_classes](data-sources--fleet--reference--group-002.md#canonical-5bc5d8130621166c2f6eb123c5dcc53d4e0ebc5df992e0811f10163c06975c9f) |
| `storage_class_list.storage_classes.advanced_storage_parameters` | [storage_class_list.storage_classes.advanced_storage_parameters](data-sources--fleet--reference--group-002.md#canonical-0e60fa0baba1bea8f29010dc76a3e408f17ab54f49e576c734e55944a8a35b77) |
| `storage_class_list.storage_classes.allow_volume_expansion` | [storage_class_list.storage_classes.allow_volume_expansion](data-sources--fleet--reference--group-002.md#canonical-9d59284bed479759cf27d499dd15ceb3ec1bc3db80b62ff36dc88047a453d408) |
| `storage_class_list.storage_classes.custom_storage` | [storage_class_list.storage_classes.custom_storage](data-sources--fleet--reference--group-002.md#canonical-453a51acc3f1b44ea9f8323d7e4884d22638d7ffb59586bcf26c20dbbad792f7) |
| `storage_class_list.storage_classes.custom_storage.yaml` | [storage_class_list.storage_classes.custom_storage.yaml](data-sources--fleet--reference--group-002.md#canonical-ebde82326fbbe94d50454e3d16f935012fea0d5c580b69b47f7311daf88f9f4c) |
| `storage_class_list.storage_classes.default_storage_class` | [storage_class_list.storage_classes.default_storage_class](data-sources--fleet--reference--group-002.md#canonical-325cd836a8c97a11c145898356eda0709e38588eab0a1cfe622bcd85dea853b6) |
| `storage_class_list.storage_classes.description_spec` | [storage_class_list.storage_classes.description_spec](data-sources--fleet--reference--group-002.md#canonical-d98f9e34f721fe47907c2ae4622e5d3da01c7e1e8cb712541df3de30d8ecc328) |
| `storage_class_list.storage_classes.hpe_storage` | [storage_class_list.storage_classes.hpe_storage](data-sources--fleet--reference--group-002.md#canonical-4fe26e387bb474daa5158e6392a3d5ef9d788276b645e42b872a88de9b41c658) |
| `storage_class_list.storage_classes.hpe_storage.allow_mutations` | [storage_class_list.storage_classes.hpe_storage.allow_mutations](data-sources--fleet--reference--group-002.md#canonical-4a71069f0e7b2a9ed891051e32850f4db635c605a092be3811b40e82054e5ca4) |
| `storage_class_list.storage_classes.hpe_storage.allow_overrides` | [storage_class_list.storage_classes.hpe_storage.allow_overrides](data-sources--fleet--reference--group-002.md#canonical-14c3285a1c585f2897d84079ded522db99459abc92e0d7057fb5ac79c65a0f3a) |
| `storage_class_list.storage_classes.hpe_storage.dedupe_enabled` | [storage_class_list.storage_classes.hpe_storage.dedupe_enabled](data-sources--fleet--reference--group-002.md#canonical-7414aec908160d20864da33323ad7570256fd5b25f21f825e90da8e6ac6b4c6b) |
| `storage_class_list.storage_classes.hpe_storage.description_spec` | [storage_class_list.storage_classes.hpe_storage.description_spec](data-sources--fleet--reference--group-002.md#canonical-e2c0e2d423dc6031f91c22c3310e01dc3f387be692e9ccb015927b5dde7924d4) |
| `storage_class_list.storage_classes.hpe_storage.destroy_on_delete` | [storage_class_list.storage_classes.hpe_storage.destroy_on_delete](data-sources--fleet--reference--group-002.md#canonical-6cc3cdd14ca7c2b0aa5f5f9d3e2f5adf0c2b42269d8bcfd911350585ea3f5208) |
| `storage_class_list.storage_classes.hpe_storage.encrypted` | [storage_class_list.storage_classes.hpe_storage.encrypted](data-sources--fleet--reference--group-002.md#canonical-1980172e913b385bcfb4ce0b81dd6e2a894e348f7ea41b0a7995414f93f0a675) |
| `storage_class_list.storage_classes.hpe_storage.folder` | [storage_class_list.storage_classes.hpe_storage.folder](data-sources--fleet--reference--group-002.md#canonical-80e67dd21ef32f3dd45bb559e3baf14610e5708c8ab9e02ced28acca3dfab286) |
| `storage_class_list.storage_classes.hpe_storage.limit_iops` | [storage_class_list.storage_classes.hpe_storage.limit_iops](data-sources--fleet--reference--group-002.md#canonical-1d724b6e45fbeae9053517d38537bfba9fe652391ed943db389e6d678fb0ceb4) |
| `storage_class_list.storage_classes.hpe_storage.limit_mbps` | [storage_class_list.storage_classes.hpe_storage.limit_mbps](data-sources--fleet--reference--group-002.md#canonical-ff0c8ef21d3462fd175a30d9da24d0e5f57b669753e3e2c1c6fb0386c24c1fbc) |
| `storage_class_list.storage_classes.hpe_storage.performance_policy` | [storage_class_list.storage_classes.hpe_storage.performance_policy](data-sources--fleet--reference--group-002.md#canonical-508aa2727338c1367d23460a18e9b120e7b1e84490be463ad268fcaffa2f4bee) |
| `storage_class_list.storage_classes.hpe_storage.pool` | [storage_class_list.storage_classes.hpe_storage.pool](data-sources--fleet--reference--group-002.md#canonical-8b251bf71ea73698b0dddfd759ed6d21af18aa985b8003b8231a9e4c9bae5356) |
| `storage_class_list.storage_classes.hpe_storage.protection_template` | [storage_class_list.storage_classes.hpe_storage.protection_template](data-sources--fleet--reference--group-002.md#canonical-aeb08a56dd47c9b8eedb008a1d862d65f68a10632bba7dfc5dbc8443f9d65865) |
| `storage_class_list.storage_classes.hpe_storage.secret_name` | [storage_class_list.storage_classes.hpe_storage.secret_name](data-sources--fleet--reference--group-002.md#canonical-62a1f6be24beb21fbc61b9a3868b68b564b4700e540d92475b5f5a39eb4bcaca) |
| `storage_class_list.storage_classes.hpe_storage.secret_namespace` | [storage_class_list.storage_classes.hpe_storage.secret_namespace](data-sources--fleet--reference--group-002.md#canonical-5890c74f22ce05fbe2c4d0e629ea2ad0bd89a57d3b3f46e4f1fbb2729d8c91c0) |
| `storage_class_list.storage_classes.hpe_storage.sync_on_detach` | [storage_class_list.storage_classes.hpe_storage.sync_on_detach](data-sources--fleet--reference--group-002.md#canonical-437e74a480edb70c1b051a361eb0ee067f7147fc521fb5e1b04f609c1559aa63) |
| `storage_class_list.storage_classes.hpe_storage.thick` | [storage_class_list.storage_classes.hpe_storage.thick](data-sources--fleet--reference--group-002.md#canonical-b8e97cd93baf54068e183ca3e2860dcc4d4a5060cc286d0454bdd981439fee46) |
| `storage_class_list.storage_classes.netapp_trident` | [storage_class_list.storage_classes.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-447a3f89adeab2e6898f8f651bf284d86e16e169ef6d616069d5b026e8c4ff16) |
| `storage_class_list.storage_classes.netapp_trident.selector` | [storage_class_list.storage_classes.netapp_trident.selector](data-sources--fleet--reference--group-002.md#canonical-64399c4d2fe6162c60a95d870c33e43ae61afab04e55cadd9cacda4d45f89b99) |
| `storage_class_list.storage_classes.netapp_trident.storage_pools` | [storage_class_list.storage_classes.netapp_trident.storage_pools](data-sources--fleet--reference--group-002.md#canonical-13c23a052393427becef8ea56c3aff2684f5f7b79c6edf21785c46b4db9b1928) |
| `storage_class_list.storage_classes.pure_service_orchestrator` | [storage_class_list.storage_classes.pure_service_orchestrator](data-sources--fleet--reference--group-002.md#canonical-26f987882bad549d35c70fed33a5527df57a6c7ad8d46361377897d1925b5dd0) |
| `storage_class_list.storage_classes.pure_service_orchestrator.backend` | [storage_class_list.storage_classes.pure_service_orchestrator.backend](data-sources--fleet--reference--group-002.md#canonical-702ba51e35b782138466b03a05ba1e124bfd407872e756f004e8a55fa1d2cf7e) |
| `storage_class_list.storage_classes.pure_service_orchestrator.bandwidth_limit` | [storage_class_list.storage_classes.pure_service_orchestrator.bandwidth_limit](data-sources--fleet--reference--group-002.md#canonical-f071e50e056b7405996c4c3a78f3145a1002473575818c9f4c9d5534b306a7c9) |
| `storage_class_list.storage_classes.pure_service_orchestrator.iops_limit` | [storage_class_list.storage_classes.pure_service_orchestrator.iops_limit](data-sources--fleet--reference--group-002.md#canonical-f3b3f300f4cc539adb712fe459887cad40e156d509c0f08d7d37803002efebc5) |
| `storage_class_list.storage_classes.reclaim_policy` | [storage_class_list.storage_classes.reclaim_policy](data-sources--fleet--reference--group-002.md#canonical-3ebd0790f795236256874f120aa80ab42e929e52670873fa1db250733bceaa75) |
| `storage_class_list.storage_classes.storage_class_name` | [storage_class_list.storage_classes.storage_class_name](data-sources--fleet--reference--group-002.md#canonical-1724a146b1c44f098430aef5e36817b0e5d0ca4f6649726f3273aee81873c808) |
| `storage_class_list.storage_classes.storage_device` | [storage_class_list.storage_classes.storage_device](data-sources--fleet--reference--group-002.md#canonical-2a7a4c6a41b175eb959f84783ff1b29ea05ba7a926c58de359604daf3eb66896) |
| `storage_device_list` | [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-11dc42425e26dc557fb3360db00c44c3009be00c8b173b3ddc632cc974c46928) |
| `storage_device_list.storage_devices` | [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-be78c01eddf6a7c854e0bdac20a940b797642c933c7d06d0904ead83425cd93c) |
| `storage_device_list.storage_devices.advanced_advanced_parameters` | [storage_device_list.storage_devices.advanced_advanced_parameters](data-sources--fleet--reference--group-002.md#canonical-136d78d7f127c5992942c90c42db08b357bc3c39e76da70e3f4c034d3f4f6244) |
| `storage_device_list.storage_devices.custom_storage` | [storage_device_list.storage_devices.custom_storage](data-sources--fleet--reference--group-002.md#canonical-8ffd540d3a86cdfd931c61d77f3d3da45b142f2129332712d8e8078ff0cac256) |
| `storage_device_list.storage_devices.hpe_storage` | [storage_device_list.storage_devices.hpe_storage](data-sources--fleet--reference--group-002.md#canonical-665a1e0eddcea9646f74171054074bd2ae373c0fb94b7a9116140f4af9e12ee9) |
| `storage_device_list.storage_devices.hpe_storage.api_server_port` | [storage_device_list.storage_devices.hpe_storage.api_server_port](data-sources--fleet--reference--group-002.md#canonical-9738af5bd97b61f0909e6c0a601dc5fc0dcab53aac9e885b336acda86147b567) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](data-sources--fleet--reference--group-002.md#canonical-c2c2f805769dc4102edd68f0ba1637dc25ec7217344c1e4a560033a187affeae) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info](data-sources--fleet--reference--group-002.md#canonical-a0484192c4598e05b1227bd08135e76d2d3f68a031ef9201123467d316cfd056) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.decryption_provider](data-sources--fleet--reference--group-002.md#canonical-ace761165b5f483139b8fb69e400b980a0913c32e2cddc58866a2dba13358a24) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.location` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.location](data-sources--fleet--reference--group-002.md#canonical-f52472070673758273d592ab6725c556c6d21bff9ff388981ed7018fbed02045) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.store_provider](data-sources--fleet--reference--group-003.md#canonical-3ceaaaffc3a5589f5ab28b5fc288d48e17d17043b01c9f9e693f4d7a30a63f8c) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-e5f7b5fb7c4021273e957bf96bb32343623e9280b73d822288513f3bcd639949) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info.provider_ref](data-sources--fleet--reference--group-003.md#canonical-24b4c39e2b8e10e2af9f7ab2aa2c5f764b3905e534bdbd4b6e9951040fcf3e63) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info.url` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info.url](data-sources--fleet--reference--group-003.md#canonical-b165b50d97945b80d5f239411bc706bc03d0659614544af3df82e7b28cba047b) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_user` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_user](data-sources--fleet--reference--group-002.md#canonical-1382b0283d293d59db10c3f948ad0efdee7ae58cf17348be1c82739a6048ab2f) |
| `storage_device_list.storage_devices.hpe_storage.password` | [storage_device_list.storage_devices.hpe_storage.password](data-sources--fleet--reference--group-003.md#canonical-cad629c71596e54508f4cb8798ec94088984fb05c95d654089234ca3316e179f) |
| `storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info` | [storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-f820ed53e7beede9fc6e2f654db83187ce593f146d0abc90bb4ddaf26b2c5270) |
| `storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.decryption_provider](data-sources--fleet--reference--group-003.md#canonical-59985ff0fe95da5e98e0c2732ce4b2e0ec4747667d3ad5864b2b16cce07906ba) |
| `storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.location` | [storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.location](data-sources--fleet--reference--group-003.md#canonical-9a453404c56e5cd9a07b3bb44e6e59fde1cf60b8fc796fb801e314f967d4a520) |
| `storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.store_provider](data-sources--fleet--reference--group-003.md#canonical-f2143b2ebd2cbcd4e9eb5942307181786de7c41c638d7905df64684675fb2cdc) |
| `storage_device_list.storage_devices.hpe_storage.password.clear_secret_info` | [storage_device_list.storage_devices.hpe_storage.password.clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-c0d9938be472df5f0f9d3cf105edacd7c3246c008d9cb1ee385de63d159f70fb) |
| `storage_device_list.storage_devices.hpe_storage.password.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.hpe_storage.password.clear_secret_info.provider_ref](data-sources--fleet--reference--group-003.md#canonical-48c86ebfd987ebde01ed831dc28efefa90537652501276fccd14cb192d681518) |
| `storage_device_list.storage_devices.hpe_storage.password.clear_secret_info.url` | [storage_device_list.storage_devices.hpe_storage.password.clear_secret_info.url](data-sources--fleet--reference--group-003.md#canonical-debdbbad88e9110853d98163faf3ca227de7ae8641bb2e90a6d256fc6fb80445) |
| `storage_device_list.storage_devices.hpe_storage.storage_server_ip_address` | [storage_device_list.storage_devices.hpe_storage.storage_server_ip_address](data-sources--fleet--reference--group-002.md#canonical-bcd95d3d2bcb97454bb48dc65f0bbc145a476f60eb9224588932f0c79e047dd6) |
| `storage_device_list.storage_devices.hpe_storage.storage_server_name` | [storage_device_list.storage_devices.hpe_storage.storage_server_name](data-sources--fleet--reference--group-002.md#canonical-6808956d0daffd006039082eb0f5225df1752051154625c31fa8f5dc51e8844f) |
| `storage_device_list.storage_devices.hpe_storage.username` | [storage_device_list.storage_devices.hpe_storage.username](data-sources--fleet--reference--group-002.md#canonical-2434c0998d817c128e66c7bc1665537940e94ace7f805bb0905f38215846f759) |
| `storage_device_list.storage_devices.netapp_trident` | [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-e6e8ddc2004f3efc215209c333d6fc8e8f2e810252541a93741245b250b7f870) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-ef381087c940b70f900fad61a3d17841cf31289ef9d4306bdcf6c03bb654a105) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs](data-sources--fleet--reference--group-003.md#canonical-0a25d9d71924c1de296affa4818381fd2665cc19fe2f63efa8e49123c71e4ad2) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs.prefixes` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs.prefixes](data-sources--fleet--reference--group-003.md#canonical-21ad16ea62fe8663eeccb1599a1cf5c3893843021f19d1e47628352e1fdd4414) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_policy](data-sources--fleet--reference--group-003.md#canonical-f527f14b77ca36542d945a79a2e1820cfa67c53dc08dada4765b200071fc3dca) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.backend_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.backend_name](data-sources--fleet--reference--group-003.md#canonical-911fa38195bdd95a9aba30c8f45575b87b43e56eb05ad9d1a54e95326edc8627) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_certificate` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_certificate](data-sources--fleet--reference--group-003.md#canonical-93bac4abfb33ab6069e2a5fa8b0caed91b318223aff59c94b4485d2527275d53) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](data-sources--fleet--reference--group-003.md#canonical-d13f1fd894cae72721e639fa5bdf622eb66187416776dbd8eaa5941723f8ac02) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-eeabe534359e2251315326bdcf44d24f7d45e9bc5a13a04cbe6eb6349905b416) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.decryption_provider](data-sources--fleet--reference--group-003.md#canonical-5717eb03112786b73e15f11200f900b0c7d4b6b64577b058c942f858780d0e89) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.location` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.location](data-sources--fleet--reference--group-003.md#canonical-310a9f14f214c38d9edd472293e1b66506f86a322956a6e0aa7739ddce30830a) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.store_provider](data-sources--fleet--reference--group-003.md#canonical-d01ad4fbf8d961fc1ca68a6dff4ef65387c1b309ad185855a32ffbb8bde6f2c8) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-14687681dc4f6116b4d5eaec68310dfb2bc15640595caa16705035d9897d5bff) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info.provider_ref](data-sources--fleet--reference--group-003.md#canonical-981b5a9608de85209a71b3299e811cc3b19a21e2a3ff47364d4943e6042e33b5) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info.url` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info.url](data-sources--fleet--reference--group-003.md#canonical-6218a7defd2f125e5cfbfbc9edf67ccb8a8744367602deb25f954e2a8f4eded5) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.data_lif_dns_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.data_lif_dns_name](data-sources--fleet--reference--group-003.md#canonical-e1d23caf74e577d7e8b2c7417b7d17a951ad83b3eeadee065fd458d88925d819) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.data_lif_ip` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.data_lif_ip](data-sources--fleet--reference--group-003.md#canonical-7036ea1b996295da10771061ef8e08ea9ec0c37a54466825960735a4e1550964) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.labels` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.labels](data-sources--fleet--reference--group-003.md#canonical-d6bc4c1df5b6412cf922ee98c241fbfc0471a36f9efcfe8823fd3fd114f0a580) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.limit_aggregate_usage` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.limit_aggregate_usage](data-sources--fleet--reference--group-003.md#canonical-60d936e6875ed9434e95eb40abb902c974820549c5b0eb506bde201d715d051d) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.limit_volume_size` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.limit_volume_size](data-sources--fleet--reference--group-003.md#canonical-82e9531e364225903fc89115fe4f46d3b887286f0939d40ea6f7b5989910c7db) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.management_lif_dns_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.management_lif_dns_name](data-sources--fleet--reference--group-003.md#canonical-7d37f2939cc5eeb50a16ad4b078f6acc5fcad138a6062e69603aaa0986afbf8b) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.management_lif_ip` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.management_lif_ip](data-sources--fleet--reference--group-003.md#canonical-4584e498ff56a9ab973227f7f4d74985fb87f8d313e1f042c23630f94cfdd1af) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.nfs_mount_options` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.nfs_mount_options](data-sources--fleet--reference--group-003.md#canonical-2f17a5b660b55226f79cad24638da083f29eeff3fc7b6f3a0dd29db9eb84f735) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](data-sources--fleet--reference--group-003.md#canonical-1e8bec6765a10855ec41055588703959a40e68394e862b3e2a872d3f4a851cbf) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-600a40f39c54548c26d9e47b85f44b9e11c194ab7deaac6540af0256c789486e) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.decryption_provider](data-sources--fleet--reference--group-003.md#canonical-105f021b10507171d89f554adf5ee42745d42671e045c39cfb5c0774c61148f4) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.location` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.location](data-sources--fleet--reference--group-003.md#canonical-24337c788254d47c3b811667c08e6927445cba6231e736de09821a82f2cce9b3) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.store_provider](data-sources--fleet--reference--group-003.md#canonical-ab504f20508b9662dc1bf7ffab6736ce8ec04d2d4ac19524fac51d81a1e47a32) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-15e7a0cb8345478648d6ae970bbdde9b4f0cf3d8aad8e37f087917c44c5b8e4f) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info.provider_ref](data-sources--fleet--reference--group-003.md#canonical-ce96d4fa5a7209bc404a5104ab069c08516a887a3df2388afddac098f0d4bca9) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info.url` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info.url](data-sources--fleet--reference--group-003.md#canonical-a308fb2f9e3f7a31998b15ea9ebd57eaae55c61c8ce7415af931885139f7478a) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.region` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.region](data-sources--fleet--reference--group-003.md#canonical-339f390d266a67e561fa7921218edc14e3b751a492f37e8060720eb75c56bec5) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](data-sources--fleet--reference--group-003.md#canonical-4ad8d17956f34c2269cbec737aa7bf8fd88c8059d8e352e48da99b51c39ffcf7) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.labels` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.labels](data-sources--fleet--reference--group-003.md#canonical-113a506c1c74514995bae3f0a6e8f4af39efd3b48725fa57b94a4709093db7cc) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](data-sources--fleet--reference--group-003.md#canonical-aa08f95aba9f659666dd1807733ff7d353c6f958d035d08fd0de553fb056d43d) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.adaptive_qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.adaptive_qos_policy](data-sources--fleet--reference--group-003.md#canonical-19cb805289c1898ceeba55c8cb604f5f9a5b990e24507b3cfb9a27bbac13460a) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.encryption` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.encryption](data-sources--fleet--reference--group-003.md#canonical-5e7f7e8ac1e27c45fe6af98cd40049eb9f33736df751b4edd922306be869508f) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.export_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.export_policy](data-sources--fleet--reference--group-003.md#canonical-d097c9d12f85acf11c2805f5f31b86810195c61c1bd6b072965fa46640c5f08f) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos](data-sources--fleet--reference--group-003.md#canonical-3efa8b4c3be828276d53655519bc93ab54c0e513f3b02e8e3cb0a6414682d27f) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.qos_policy](data-sources--fleet--reference--group-003.md#canonical-5e180fd7f97f946a626f0f08341be06240d0368b2540cd5744a437610722600b) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.security_style` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.security_style](data-sources--fleet--reference--group-003.md#canonical-88768ea87a1410abb74135ac5a711628de703806e1bba21454b55fce1e26b23e) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_dir` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_dir](data-sources--fleet--reference--group-003.md#canonical-bb52305686ed274431c897d416577621afad1bb1f98c05cd115afb724f2fcc31) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_policy](data-sources--fleet--reference--group-003.md#canonical-736bbc8fddfccbb9b70adfddf5741a316452d8e8711ac6e9ad02269a33382b3d) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_reserve](data-sources--fleet--reference--group-003.md#canonical-7675d6ce88606f8bd2ff8c1738a4c21c8655651388dd91354dd0ad1532ae950b) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.space_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.space_reserve](data-sources--fleet--reference--group-003.md#canonical-dd1f61b66983ae1893bcacaa3c860df73add2846fd89058f62b89c4244a4fbe5) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.split_on_clone` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.split_on_clone](data-sources--fleet--reference--group-003.md#canonical-aab4d628fb2c01a5b453e519d47233cad72fcd342c6464136bda5ef2dda32889) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.tiering_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.tiering_policy](data-sources--fleet--reference--group-003.md#canonical-83ec3364fd51b1733fac0038db7f7c22d298191ebc4a8f0b243b05c1459ea88f) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.unix_permissions` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.unix_permissions](data-sources--fleet--reference--group-003.md#canonical-bddd4bd67fd9f364be8cfb4d98785e5e1cb6ab69849093667f9a78eae4aadd23) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.zone` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.zone](data-sources--fleet--reference--group-003.md#canonical-a1350b100919214eed81f0513fb757410709e44548c31dd818023e7409e0c402) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage_driver_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage_driver_name](data-sources--fleet--reference--group-003.md#canonical-cfc67397bced749e408ced4ebb888bd8448ea88a894db62690efdcfb0984591e) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage_prefix` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage_prefix](data-sources--fleet--reference--group-003.md#canonical-c85ea91b14827785ab95c1ba86691f34d7fde9c084b22b179e0a83b5f5695930) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.svm` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.svm](data-sources--fleet--reference--group-003.md#canonical-0750b6275ca4edfd98816c1f1fab015212bfedf5c5ce68549fb2adc25103ca79) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.trusted_ca_certificate` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.trusted_ca_certificate](data-sources--fleet--reference--group-003.md#canonical-8bcfe3e0ec88da8fe5827c189f1606478370335f67ee291695f089f28036bd49) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.username` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.username](data-sources--fleet--reference--group-003.md#canonical-f0772c0af24cf87eaacfa85bc420fa2a70fb96b5cc5848e38daab1e29e38bcd8) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults](data-sources--fleet--reference--group-003.md#canonical-d8a1addaf905223e9e84b887410f9051bfcf8df7d297edf7cde44d6f9abc5ae4) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.adaptive_qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.adaptive_qos_policy](data-sources--fleet--reference--group-003.md#canonical-ab3440b2d1bfaf73f1018b9d94b6c3c15e705fa238077be194c2485fe0690783) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.encryption` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.encryption](data-sources--fleet--reference--group-003.md#canonical-d49a1e703f063d918795d43435bd60b88616076b5b05642643856e034f6c8504) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.export_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.export_policy](data-sources--fleet--reference--group-003.md#canonical-c0a5b8eaf0b867652f0a7276bf543aa7acadbb9a624055ce06c7125558750269) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos](data-sources--fleet--reference--group-003.md#canonical-8904ac0de240ab7430156ac0c7ccb462d87c58e86ed81e06511091386b0237e2) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.qos_policy](data-sources--fleet--reference--group-003.md#canonical-a7c15e2a83715be3200a9cdb918704ae8ad7fe6f91f5e478fc2b4d4b2f68ef33) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.security_style` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.security_style](data-sources--fleet--reference--group-003.md#canonical-7e1693a5d69b1b310e5d92d1d288ac96b75261ef1a9e93f4aaa2d1ee516fe1cb) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_dir` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_dir](data-sources--fleet--reference--group-003.md#canonical-7fdcc5b01b03c7b0f729cd6f95feb0dc583fc27bb19bbe3aa2cf9c56e19c380f) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_policy](data-sources--fleet--reference--group-003.md#canonical-86a4e989d58efd9bec3ebfcd819aab285dbf9af0581a2e240004ba4a6fdc8533) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_reserve](data-sources--fleet--reference--group-003.md#canonical-e45d85dc720e3a1271eb72e063932e2c40f6aa5722aad25a195231d08cb2202a) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.space_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.space_reserve](data-sources--fleet--reference--group-003.md#canonical-0cf5abc1c6c4ccce2274edeaf7e95ee14d7ba774c67b2ed389b4cb89974bbefa) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.split_on_clone` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.split_on_clone](data-sources--fleet--reference--group-003.md#canonical-7af7d8fecb51f970dd986445abebe678d633956af9c01151527058e7cd60784d) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.tiering_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.tiering_policy](data-sources--fleet--reference--group-003.md#canonical-3b154d1a7dc42a5bc85cc66903f1d9ba034d27cd3f2268c2bf858d7be8b255c2) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.unix_permissions` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.unix_permissions](data-sources--fleet--reference--group-003.md#canonical-a391ae65bea0d117eaeb410f49ea43c31f67e586cc567e58e64ce181f6b7647b) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-36582d3705b383c029ae82b01b7cf1043dd509a781f5786cdefd7082350dcc31) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_certificate` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_certificate](data-sources--fleet--reference--group-003.md#canonical-534bbdba5ee2585110861b0fc3181cf1592d2129875ce64c9a49262a75575e82) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](data-sources--fleet--reference--group-003.md#canonical-cd0f2624564cff86baf21360ae7943ab0f825455d0ae41bd2fac8abe53bb056d) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-e0bef78a9578e784b38d7175d1a6f8079067a7e72433512882a9331957eda40d) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.decryption_provider](data-sources--fleet--reference--group-003.md#canonical-542d5d966d25b3835494cc096cb9c6ba7d783038a4769869c67ef0a8e38d912b) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.location` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.location](data-sources--fleet--reference--group-003.md#canonical-03e50b74846cf4aedefd784076afdc184fd77ea973609ac37049ea24d1061d19) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.store_provider](data-sources--fleet--reference--group-003.md#canonical-610002603def2cb443245b30941bf1e0b057e5975e39506e2eb69e5c1bfcfc0d) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-7b0a6b69cd83c552613ae4501ecb025fe1d576c1bda8cde921374dbd5db2aea1) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info.provider_ref](data-sources--fleet--reference--group-003.md#canonical-93deb9b356a3e653c37cf32cccc9e0f8d9416d8d4fb07d3f1bc42296df7a5cd0) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info.url` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info.url](data-sources--fleet--reference--group-003.md#canonical-2e8be5c557bb37c46147289dd7358767dfde8eee787c82ada1b74e895926ab09) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.data_lif_dns_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.data_lif_dns_name](data-sources--fleet--reference--group-003.md#canonical-6b16d38df887113866eea136d0495ce315734658671d99e1ded159238a9f1874) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.data_lif_ip` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.data_lif_ip](data-sources--fleet--reference--group-003.md#canonical-67ca8ca4a71ef18e0fa297193c67ff7d32870961561479e2ccd4cedd16993832) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.igroup_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.igroup_name](data-sources--fleet--reference--group-003.md#canonical-84edcb7f4bcb89159bf34777ff15b3cb5f7cd840b98f1c49932ad5247933d943) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.labels` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.labels](data-sources--fleet--reference--group-003.md#canonical-2a0d90b35e5ecbfa91b77c71879d9f3f2d0f05a33291828cc84e5e7a543be643) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.limit_aggregate_usage` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.limit_aggregate_usage](data-sources--fleet--reference--group-003.md#canonical-ea75a01e99b30cfc5a8d559cbdb89ffdda4315fb00c9d729aa0aa9704ace6cad) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.limit_volume_size` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.limit_volume_size](data-sources--fleet--reference--group-003.md#canonical-1f6e0c4618b156220c8272389e6449ae6c9255fa35fbb34777e127513654d49a) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.management_lif_dns_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.management_lif_dns_name](data-sources--fleet--reference--group-003.md#canonical-f4bc4e2f90ffadcbe84cbfee964dbdfc4b2b1a42af28225ae0444e8fc8841ecd) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.management_lif_ip` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.management_lif_ip](data-sources--fleet--reference--group-003.md#canonical-40a13a5939bcdd2ae5bb31f34838d79145f3212dce4cb379e2c4f52413c57d0e) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap](data-sources--fleet--reference--group-003.md#canonical-34356a201b9f4dafe60d6d7e8c1c7a5d7209b338c919a7a4f17f3b3a471f3d1a) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](data-sources--fleet--reference--group-003.md#canonical-34b8d03a88fe8fafd20299f7fc341adb078cd349428a3aa32903a856ab2a51e0) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-45384795643df51965e2e8d4d18598b912b7065518d3a724a40023074ba98ab1) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.decryption_provider](data-sources--fleet--reference--group-003.md#canonical-c94e9e9d67a64be00c9ed5c0b51c0f999dd420c593944a28aad7451055d72c52) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.location` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.location](data-sources--fleet--reference--group-003.md#canonical-2aab6dcfdfe900466d79d0a247274935c95b54a98ffc4d61581fbb41a7dbe25b) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.store_provider](data-sources--fleet--reference--group-003.md#canonical-7567847b8b721a836d5c73dda0605f6dbdf546b21028e6f6080fedc2209c337f) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-ed109466bc8c79ee3f4fca86762e4b25a920fc861965da9647e85843b1976793) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info.provider_ref](data-sources--fleet--reference--group-003.md#canonical-f2d4159a0b89b3e6ed4a3b0d6e39279ae43282073b4ddcd15de3f7c46d678df9) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info.url` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info.url](data-sources--fleet--reference--group-003.md#canonical-42ed4f0373d3ee88762cd50e48b14af826d98d4da07ab9c698a0f4900b44d472) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.region` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.region](data-sources--fleet--reference--group-003.md#canonical-882844eb8971292bdfecd9b9c7f8ef8907b2390544c9339ea8df1f82c0091f77) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](data-sources--fleet--reference--group-003.md#canonical-34db7506229fb88cc65890a3b1d6237493a5e5010b7250952413de616090e5d4) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.labels` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.labels](data-sources--fleet--reference--group-003.md#canonical-1ea2aeb98e05db4366b2d7adf2b6806ff7eea0016574d782b3dd60ae9b2fb80d) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults](data-sources--fleet--reference--group-003.md#canonical-a34ddf84aace1ffcfbfe51bf78874c16d87e56cd9a3f56688169775cf5ce4334) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.adaptive_qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.adaptive_qos_policy](data-sources--fleet--reference--group-003.md#canonical-87e12be1f055c85f05af80600490ed0ab90150a7ae2847f936f719225e2c433f) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.encryption` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.encryption](data-sources--fleet--reference--group-003.md#canonical-fc03284f58ec6c78f23f815de85cee0ab46c78cb92102cced736d6a78e68a523) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.export_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.export_policy](data-sources--fleet--reference--group-003.md#canonical-9122c6b2e07d29a185a11a08b0f2f9bba37b5108230a48b382a97d0e7c6d4dca) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos](data-sources--fleet--reference--group-003.md#canonical-9ede7a09eebe38c255849beee0839423ff27675cde2c591c212deac82788f264) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.qos_policy](data-sources--fleet--reference--group-003.md#canonical-c4b1025e2058853902f304c8b7af0eb685e033633e1b890dceab3d70f15a50dc) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.security_style` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.security_style](data-sources--fleet--reference--group-003.md#canonical-2902d2087e75c3b6bcd5d2286d4b827656d208f63882e771393e18ace9849a06) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_dir` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_dir](data-sources--fleet--reference--group-003.md#canonical-5bb85f9c0584ea4025aebad3588a71473d8f1116ab1492f9a72ee8ab7c97b070) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_policy](data-sources--fleet--reference--group-003.md#canonical-529526fe56b0edb077be696d667dd8083f5d8c7b5a6aa78da423decb05dabae1) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_reserve](data-sources--fleet--reference--group-003.md#canonical-82c0cfb9cfd375038247c5cdb5074396830e125d908e31903392815ba4943451) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.space_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.space_reserve](data-sources--fleet--reference--group-003.md#canonical-c22e3cada904e2f4a5aa2b4bc8390f2fcc0e1f0bd17322cba291cd74cc03e34a) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.split_on_clone` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.split_on_clone](data-sources--fleet--reference--group-003.md#canonical-b8dfd772563349242c46f26ba544931fd13fd978853ddb96f2a58f21c1366617) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.tiering_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.tiering_policy](data-sources--fleet--reference--group-003.md#canonical-4ccf1e7332551c3b49a505028c54d4613af75ef88ddc5f40f28c9966b178eb42) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.unix_permissions` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.unix_permissions](data-sources--fleet--reference--group-003.md#canonical-86c2dd40c983a1efb7d8f2e59c0df79e865b6164224adac35b9dcc3e6255c5c1) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.zone` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.zone](data-sources--fleet--reference--group-003.md#canonical-c46b6c57c8d36c22cf9f1441208acfd573ffbcac7e02ff49ea9eda2a702b67c4) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage_driver_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage_driver_name](data-sources--fleet--reference--group-003.md#canonical-25ea3ca1b06728e1dfa5464c0f68dccc8d34daed42909959ce82961bc1d1bc51) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage_prefix` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage_prefix](data-sources--fleet--reference--group-003.md#canonical-3a699bcfc1367320c2317487d4b23b142d02990c79ab9a37e10d697a697f1e9b) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.svm` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.svm](data-sources--fleet--reference--group-003.md#canonical-1f819dacd7165fa952cc6144afa7e500b725243ca79061a5b7258ebd95794dfc) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.trusted_ca_certificate` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.trusted_ca_certificate](data-sources--fleet--reference--group-003.md#canonical-f65301a556358bda53f5e9095f15576e66505a7cf5be85f619b92738e076a1c4) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--fleet--reference--group-003.md#canonical-ca3a2b8dbc39e102a86a9168bbd20eb8bba941a3f223ab49cb16bd55a3b44714) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](data-sources--fleet--reference--group-003.md#canonical-596b557f8442c9feb215ee25e818b91644d047841d50d8114851ba3d89731cd3) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-64d88a9a3548696f090cd92c1da293a357460e7996d9932616c1be891c9abe9b) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.decryption_provider](data-sources--fleet--reference--group-003.md#canonical-58a1a3b2ceb4a361877cb24a08247787861f85b0f4b933bfb913c1e09d6eefaa) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.location` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.location](data-sources--fleet--reference--group-003.md#canonical-71ab2e324f4e1c4407380397bf78908a94d69c499b73aeb0ff79ef6e84642a2d) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.store_provider](data-sources--fleet--reference--group-003.md#canonical-df964a94735ceef8c964ba63a6c2b6fb3bd4bde768da4ac03893d1ccae81ca2c) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-ab2262bddd3fc16fee31169e7f31e6890d7d830f8687150906694245c330ac8d) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info.provider_ref](data-sources--fleet--reference--group-003.md#canonical-d70c595b3c18e08dd7c4c0e5ec73f88a0c6cc4889b3759afd6c26d511d0587ff) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info.url` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info.url](data-sources--fleet--reference--group-003.md#canonical-291e976e774bd6adba46199f38f4471cc2f1e3c5dc2b59ad6059b6f93499b499) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](data-sources--fleet--reference--group-003.md#canonical-953a2c3e224fb40d2b23c87f4535683e3f762c631fc423c1d2601794dadacd8f) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-b5437f38b1fb602a1e1881a8f658e9303597f88f5b28838739e2cb1f1f908dcd) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.decryption_provider](data-sources--fleet--reference--group-003.md#canonical-a5e1c5745790f991cc48ec151d723c2dd59323e9259d69ef1b052af8d78f3c34) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.location` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.location](data-sources--fleet--reference--group-003.md#canonical-46964e747974f14a30e97b6fe4f4f2c581a53f1c67290ba548e74f55ceb46de4) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.store_provider](data-sources--fleet--reference--group-003.md#canonical-f477c5217c9709720629ce5f859c758e83fe5d3b5af235ecea8312cc4729717e) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-4500af7aa2fedb643cdf601ef719a122ef9bb004eb7c563c7aabee907db04a36) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info.provider_ref](data-sources--fleet--reference--group-003.md#canonical-dc901a3d2f9dc55d0661eef73597e33850a0f19848cd7ba6f5bde040b8895392) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info.url` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info.url](data-sources--fleet--reference--group-003.md#canonical-8752a92467c1160daf47536314d0db90753211af1bf1c13cf50f95d1ef51335d) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_username` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_username](data-sources--fleet--reference--group-003.md#canonical-9deb5ac36903f1e0c73541bae7b39e94ed702d5200f70744fdaa34fc50903671) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_username` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_username](data-sources--fleet--reference--group-003.md#canonical-dfbceb276a530b231b94fdcb5889f461c94e27773784a7a0816a01b0350eeca1) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.username` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.username](data-sources--fleet--reference--group-003.md#canonical-8f0f52ad909f908fc7a4f915f440318f81182121448377f816cf3acf721bf6c7) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults](data-sources--fleet--reference--group-003.md#canonical-7e60ccd045312c0d783d450da947be77416982e25accf6a985e24594f1ca6bf8) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.adaptive_qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.adaptive_qos_policy](data-sources--fleet--reference--group-003.md#canonical-2bdaca4dab9daacce16edbd2659641c00046dbcf5a5af4c9257933cbccf1b6dd) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.encryption` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.encryption](data-sources--fleet--reference--group-003.md#canonical-c5d08422c13ec4b9c47ed41c98d796a5810f9fde0c71fc4b8bd17c6c5bd5ce2d) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.export_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.export_policy](data-sources--fleet--reference--group-003.md#canonical-7ddc801abea721291801ceaf9612c29a2abd7d96fd6c1b904fa810181d08e1ab) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos](data-sources--fleet--reference--group-004.md#canonical-5355cc8b078c656808ad8b942c5f86b834bb9ee9f22973415bba5df4d7e8a7e2) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.qos_policy](data-sources--fleet--reference--group-004.md#canonical-415cb181dd03b47e86268fc60369bdfd339ff8988d8f705578a4e29434b594e6) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.security_style` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.security_style](data-sources--fleet--reference--group-004.md#canonical-bbdd048deb2a6729fb6f68361010e3704988d7028f630cb0507e1991c5440f94) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_dir` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_dir](data-sources--fleet--reference--group-004.md#canonical-3384a803c595d4917e3e8c0cfdacd641a5d4370bfebda2eacc84f445b8708de4) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_policy](data-sources--fleet--reference--group-004.md#canonical-bc0b53607bc8e0d0e90ecf01edfee40527f2ae75cc500b642b143f27f920bcd3) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_reserve](data-sources--fleet--reference--group-004.md#canonical-048cbbda15c8ec27702c3d4807aa2d51b6eada8a63169becd83b050051b0035c) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.space_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.space_reserve](data-sources--fleet--reference--group-004.md#canonical-ceb422608f5537b1bdc1300ae21e9346f094e3c2b3c182abdd7cbd1501a5a7ae) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.split_on_clone` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.split_on_clone](data-sources--fleet--reference--group-004.md#canonical-1eb7996cc2ab180869a1beab1705dc6897d9a2b5a783804b2c1df28ae574f082) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.tiering_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.tiering_policy](data-sources--fleet--reference--group-004.md#canonical-824567af9e108a784eb1ae716dd84f207fa1d1b77423b9bec1fe4145db352881) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.unix_permissions` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.unix_permissions](data-sources--fleet--reference--group-004.md#canonical-ff3a2df2eed6ced21c0a768f42e462eeef4b9a986c7438f42857c18fcf35e29c) |
| `storage_device_list.storage_devices.pure_service_orchestrator` | [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-004.md#canonical-d38095c997988289cf8de0b08cf443f354f53ceea912861181847059421dc4ef) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-004.md#canonical-687fec49a8cb71f4aeb5657dfd9c37ed02b1f7bb311ed2e1fbe2023302676fde) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](data-sources--fleet--reference--group-004.md#canonical-3284ab3c06179a7c18d094c630ce5a4e02a611e0325f49b611efc888ab3a55fc) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_fs_opt` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_fs_opt](data-sources--fleet--reference--group-004.md#canonical-9dac598a14ee04a2d24da3e768b5ca5400e2f5002e1db9c40da43266059511c5) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_fs_type` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_fs_type](data-sources--fleet--reference--group-004.md#canonical-5c59c07ec5a30f88cd3fc19dd5a706d21561aa61c4b38ae1f4c8e550eb8eef35) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_mount_opts` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_mount_opts](data-sources--fleet--reference--group-004.md#canonical-95fa16b45525950148b1d2fe15cc170779fa4144c526dff7bc0a3d2aaf395992) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.disable_preempt_attachments` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.disable_preempt_attachments](data-sources--fleet--reference--group-004.md#canonical-7624e30f3ceb2c156d729efb63335d3574e25a322182d469f17e4b58c663d092) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](data-sources--fleet--reference--group-004.md#canonical-d35fb1c9876d40710e26f710093a151eb58387e6e0f1976f04928a316ad95199) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](data-sources--fleet--reference--group-004.md#canonical-36a0310de2813554b7b691fa2b3d920b9ad429794001ae147d49c7a8f231f91a) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info](data-sources--fleet--reference--group-004.md#canonical-33fbd312e7b10c6c9273b2106b9ab258e79b42a41fe9335b6c9a4bee52f20b40) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.decryption_provider](data-sources--fleet--reference--group-004.md#canonical-719e9285ada694e2c05a09c6ad02bc474998f2002f79a493d60dc3b04d8cd057) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.location` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.location](data-sources--fleet--reference--group-004.md#canonical-580ae659087fcd926f09b28baa406364ff1801a1b64faeba81047cf138379efa) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.store_provider](data-sources--fleet--reference--group-004.md#canonical-56b7ad511c7eccf862b5fa411dbcdb0dd6d0e5b13a85ae2ad6726c6b890684d0) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info](data-sources--fleet--reference--group-004.md#canonical-091fa3736fa811a166c64c919e44f7f38cbce6bd626b069cc108aead5d038afb) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info.provider_ref](data-sources--fleet--reference--group-004.md#canonical-2f25d7064329fde6bb9cc17d13a2959e91c502d230c232f406fab7728d873e44) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info.url` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info.url](data-sources--fleet--reference--group-004.md#canonical-ed2b77e6b4b70b81d38a8d30d41c3d3b5aeaec1e80322e96361df71c4a1f2600) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.labels` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.labels](data-sources--fleet--reference--group-004.md#canonical-56ab454d8b7b176589494ed39e25ecec05e5e5a33ea80a7637c7d3f5cdff2e8a) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.mgmt_dns_name` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.mgmt_dns_name](data-sources--fleet--reference--group-004.md#canonical-2006b920c64d338941cbce93b45be6a52764a0631ddfbf2c4b776ed85c0b8e14) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.mgmt_ip` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.mgmt_ip](data-sources--fleet--reference--group-004.md#canonical-42aa26a44a83d9dd3cf62f0d915dfb2ec26372214a3a33d9e4132d088670fcb9) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.iscsi_login_timeout` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.iscsi_login_timeout](data-sources--fleet--reference--group-004.md#canonical-d6459b8775e1d1d267ee5936c3f03e5f4b44125f4ed99ebaa899b672b331e7f6) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.san_type` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.san_type](data-sources--fleet--reference--group-004.md#canonical-3269b3d030eaa6941d9ee22a10333d7090ec08a33144bc63d6e315e740c713e7) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--fleet--reference--group-004.md#canonical-f19b0240ca575bfeeaf4a88ac53bf3b1719f8fe6fd6bc074f5b8d3a9774193ba) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.enable_snapshot_directory` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.enable_snapshot_directory](data-sources--fleet--reference--group-004.md#canonical-0b56a2189ce5f2004d6ceecbb2ca85d9924a37a162cb89bc48690b050c6dab70) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.export_rules` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.export_rules](data-sources--fleet--reference--group-004.md#canonical-2f235e49ffbe68631487ad1f6ebdfb694775ecbf3f9f03c21182387b2d6d9144) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](data-sources--fleet--reference--group-004.md#canonical-d7ea3d96b84d2b0ad13cbe369b67739ab86c9fc0a93d5510b3d2284ab9b3f649) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](data-sources--fleet--reference--group-004.md#canonical-093f313fa650cf94494741e2dc1413296c44e3c884b4bac97dc0acf782896481) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info](data-sources--fleet--reference--group-004.md#canonical-567c98ef18ef11d9b4e3b5e667e8286e6191ff3acd29bf16e0e6ec127a89c179) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.decryption_provider](data-sources--fleet--reference--group-004.md#canonical-e064ec906b3c94bc262c085e834ba6901909125b1975a92b80050754c8700cfc) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.location` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.location](data-sources--fleet--reference--group-004.md#canonical-8f2acd56fe05ecd6058694cd664ee2614ddeb97a2dd678e744529f27a77582bd) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.store_provider](data-sources--fleet--reference--group-004.md#canonical-ff8c17c267495d0772400beb93a78ab961e8573d5e3a8f678696042e0decf00e) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info](data-sources--fleet--reference--group-004.md#canonical-c3a0a9353de9281ab907ff79b160451a09b47cb3ecec1dc6e486901ca1359e67) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info.provider_ref](data-sources--fleet--reference--group-004.md#canonical-19b22d9cab256c1f5238d9963ae97425413256a185e4affe83d8af52fc2551e6) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info.url` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info.url](data-sources--fleet--reference--group-004.md#canonical-2ca48308f55f7da51103df9feb3f399275cc275b95decfb21ed8eb57a34b9a52) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.labels` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.labels](data-sources--fleet--reference--group-004.md#canonical-bf22fa6b63b77bf10530266e79306b032677a3a2e365e9f0ff0b883a15f9af97) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.mgmt_dns_name` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.mgmt_dns_name](data-sources--fleet--reference--group-004.md#canonical-1ab94586f0fa385bc4e01a0a9edbe38ae355e88f1724a091ac9dbc90d474b011) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.mgmt_ip` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.mgmt_ip](data-sources--fleet--reference--group-004.md#canonical-d17dbeb8d64681b610c778ea3766dbeb5e6cebdf9c36f8e418c6cfc2c72b15ce) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.nfs_endpoint_dns_name` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.nfs_endpoint_dns_name](data-sources--fleet--reference--group-004.md#canonical-dd0c819538ae1d52b71b3065ba04fa63f5f8d7cbc4a64cc788a154d8a928888e) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.nfs_endpoint_ip` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.nfs_endpoint_ip](data-sources--fleet--reference--group-004.md#canonical-6e07797af8f7c3c89962ba6e1fd91e3e8f3b333f0bd511c9301b4f6d4b3e8ade) |
| `storage_device_list.storage_devices.pure_service_orchestrator.cluster_id` | [storage_device_list.storage_devices.pure_service_orchestrator.cluster_id](data-sources--fleet--reference--group-004.md#canonical-913621344da69e5a163d523a77c4e9536f76bffe22213e1622f981a0a88a3b7e) |
| `storage_device_list.storage_devices.pure_service_orchestrator.enable_storage_topology` | [storage_device_list.storage_devices.pure_service_orchestrator.enable_storage_topology](data-sources--fleet--reference--group-004.md#canonical-90d43d1541a8f6b2a5489ebf6b406e902fb06a54c2f8ff4bf954290cfc24746a) |
| `storage_device_list.storage_devices.pure_service_orchestrator.enable_strict_topology` | [storage_device_list.storage_devices.pure_service_orchestrator.enable_strict_topology](data-sources--fleet--reference--group-004.md#canonical-7b1cd8fae06b517874d5fafdaed8ed765ac0ae68db5ce20d9a3555247327c477) |
| `storage_device_list.storage_devices.storage_device` | [storage_device_list.storage_devices.storage_device](data-sources--fleet--reference--group-002.md#canonical-5fae568a747d4bfb10957d04dccf47a99c493812806845c54c10691193361288) |
| `storage_interface_list` | [storage_interface_list](data-sources--fleet--reference--group-004.md#canonical-99c7687783cc9ee0578f5e98464c24ec87a14e0328f5d92b391257b3b8c422fa) |
| `storage_interface_list.interfaces` | [storage_interface_list.interfaces](data-sources--fleet--reference--group-004.md#canonical-5a8479f667e67e56bf7050430333044db6f8f077c3dceb0f62f0fc7595cb7e77) |
| `storage_interface_list.interfaces.name` | [storage_interface_list.interfaces.name](data-sources--fleet--reference--group-004.md#canonical-26f92560f8380e6362fb595221fb4e019be9e3e235dc85d2adcb3a99656fc74d) |
| `storage_interface_list.interfaces.namespace` | [storage_interface_list.interfaces.namespace](data-sources--fleet--reference--group-004.md#canonical-f122121084a18cd4107fe5a5babf7fa1b43157308cd9d707a20915973241cfa2) |
| `storage_interface_list.interfaces.tenant` | [storage_interface_list.interfaces.tenant](data-sources--fleet--reference--group-004.md#canonical-d1deb8111ed638377a6df50aec6c3c23b69f8ac81b1e4ee4eff232d0eaaaca19) |
| `storage_static_routes` | [storage_static_routes](data-sources--fleet--reference--group-004.md#canonical-c6f6fd551b57976de73b52a5c0b44295413498063b2a0254fffa4e58fc42fbe4) |
| `storage_static_routes.storage_routes` | [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-026769669d10bb75b33fe023235fd8fad766a3522511608e3c2de4355712c1b3) |
| `storage_static_routes.storage_routes.attrs` | [storage_static_routes.storage_routes.attrs](data-sources--fleet--reference--group-004.md#canonical-f45e34b928f7726b645d7a89341657f3155379440edd44b9c41099e2bf561461) |
| `storage_static_routes.storage_routes.labels` | [storage_static_routes.storage_routes.labels](data-sources--fleet--reference--group-004.md#canonical-d32e86d67d4d43b15082e447e1ced503c7dbe312f2dbacc36e92975626ec447f) |
| `storage_static_routes.storage_routes.nexthop` | [storage_static_routes.storage_routes.nexthop](data-sources--fleet--reference--group-004.md#canonical-7e01f9894bb00873031eca7328ca27bbee2f7361be7847fbf16a4537d163cc1e) |
| `storage_static_routes.storage_routes.nexthop.interface` | [storage_static_routes.storage_routes.nexthop.interface](data-sources--fleet--reference--group-004.md#canonical-51e0c0e6936e9fbf04081f776ec1315b6bb4a90f8432b05fe37fcc16560cce85) |
| `storage_static_routes.storage_routes.nexthop.interface.kind` | [storage_static_routes.storage_routes.nexthop.interface.kind](data-sources--fleet--reference--group-004.md#canonical-29f39c9ed37a27c29bca04986139e5f0acfc60bd4f3c483a099334fc8546022c) |
| `storage_static_routes.storage_routes.nexthop.interface.name` | [storage_static_routes.storage_routes.nexthop.interface.name](data-sources--fleet--reference--group-004.md#canonical-a62c66b8f9e5328f61113d953de77b1b46eb06bb325b940802714eec86299e17) |
| `storage_static_routes.storage_routes.nexthop.interface.namespace` | [storage_static_routes.storage_routes.nexthop.interface.namespace](data-sources--fleet--reference--group-004.md#canonical-4cd3cba1af042ab39348ad15755d2a7c0df2444eebc0c4a89fa1d51fa53bb6a0) |
| `storage_static_routes.storage_routes.nexthop.interface.tenant` | [storage_static_routes.storage_routes.nexthop.interface.tenant](data-sources--fleet--reference--group-004.md#canonical-493278edb190a819f80ad05f91990cc7215ee7380ba10ae85f0a703a8a7b3ca2) |
| `storage_static_routes.storage_routes.nexthop.interface.uid` | [storage_static_routes.storage_routes.nexthop.interface.uid](data-sources--fleet--reference--group-004.md#canonical-41d6d21d47076e3407907cdd01c1bbfdcb2d2f856b453aac4448817991ec5619) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address` | [storage_static_routes.storage_routes.nexthop.nexthop_address](data-sources--fleet--reference--group-004.md#canonical-3b99b7b056f6e19ee89fbb49c3354b312e568d4adb31f50496d866a2d94fd809) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack` | [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack](data-sources--fleet--reference--group-004.md#canonical-5382f0c2697d7807a3384b2271f040d45cbd65b3f53f53496b2bd849a47f3bb6) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4` | [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4](data-sources--fleet--reference--group-004.md#canonical-c5d1ae39bd21cb2fbaf4784eb2d615decc78b5c44e0c49bbe4dd7f2832e43acb) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4.addr` | [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4.addr](data-sources--fleet--reference--group-004.md#canonical-ccdaa0bea55bf44a3a6876a9d23a27861810178c20dfd6cd202efd51b278e2e4) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6` | [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6](data-sources--fleet--reference--group-004.md#canonical-9498970e6251a0128afbe7ef495c744f03d05a1430016d9c5d1a6aed84a3d917) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6.addr` | [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6.addr](data-sources--fleet--reference--group-004.md#canonical-0df9ab1bb848e6a383dca33d58ec7f368b2e55fd5943e1a0c0b1094d45ca4a46) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4` | [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4](data-sources--fleet--reference--group-004.md#canonical-14f790ac84fb675d9c7bb54a25211b767476a97f9d95a7ed4c11797355c80c9d) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4.addr` | [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4.addr](data-sources--fleet--reference--group-004.md#canonical-c7bee82edd4f1914e033d27a24ef39fd218878946607d146f0115563ddadc637) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6` | [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6](data-sources--fleet--reference--group-004.md#canonical-a8725eb8320870a87a27715813e4239549470b44728e2c48ea8f3da48dfdbb26) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6.addr` | [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6.addr](data-sources--fleet--reference--group-004.md#canonical-b07d1272621c8b441793939283df1b25b8954b28e2184ac1a83d5133ef74e3ae) |
| `storage_static_routes.storage_routes.nexthop.type` | [storage_static_routes.storage_routes.nexthop.type](data-sources--fleet--reference--group-004.md#canonical-5c87ee0ea571887a7706c54967563abbccbeafe9e7c7809342c9983d95a6467a) |
| `storage_static_routes.storage_routes.subnets` | [storage_static_routes.storage_routes.subnets](data-sources--fleet--reference--group-004.md#canonical-19a2db8fa1618e9888b690dd56d93f654f2f1cfb036ced150387f9ef57ccfa6c) |
| `storage_static_routes.storage_routes.subnets.ipv4` | [storage_static_routes.storage_routes.subnets.ipv4](data-sources--fleet--reference--group-004.md#canonical-d525b1d8572106bf3de79f855c2c8da3c9dcd65c18bbb674bb429388e7696c6f) |
| `storage_static_routes.storage_routes.subnets.ipv4.plen` | [storage_static_routes.storage_routes.subnets.ipv4.plen](data-sources--fleet--reference--group-004.md#canonical-4486e4efe62a81e1426cc45a320684bb12eb4e7f4919d7efbcc518597121ac6a) |
| `storage_static_routes.storage_routes.subnets.ipv4.prefix` | [storage_static_routes.storage_routes.subnets.ipv4.prefix](data-sources--fleet--reference--group-004.md#canonical-fc68b29fc989e4ec222272507c5775276abf6f05892fe6422e62c4ebea8d8561) |
| `storage_static_routes.storage_routes.subnets.ipv6` | [storage_static_routes.storage_routes.subnets.ipv6](data-sources--fleet--reference--group-004.md#canonical-7cfad6b7e82a1de6fe312e8fc95d18f7837384a938ea20e4e31d74c4ff3d5207) |
| `storage_static_routes.storage_routes.subnets.ipv6.plen` | [storage_static_routes.storage_routes.subnets.ipv6.plen](data-sources--fleet--reference--group-004.md#canonical-ae4e4dad694f726a1b41d0706098ce3741b4a08d382a1a149d424841778ff100) |
| `storage_static_routes.storage_routes.subnets.ipv6.prefix` | [storage_static_routes.storage_routes.subnets.ipv6.prefix](data-sources--fleet--reference--group-004.md#canonical-56e36ca3092453fe0e1ba96e8698f039e23708530e0ac3864a82dfca89d05869) |
| `usb_policy` | [usb_policy](data-sources--fleet--reference--group-004.md#canonical-0d5e453c786da8ccd200294830086acbd0edd330738a8ba1b231459c978d547c) |
| `usb_policy.name` | [usb_policy.name](data-sources--fleet--reference--group-004.md#canonical-31a448a51e44954976abbb84a6b8195ae5f90b907b6241c5c8a8ec70de589841) |
| `usb_policy.namespace` | [usb_policy.namespace](data-sources--fleet--reference--group-004.md#canonical-03a3a8cd97bc12ca874ac3730d1f42b91a3bf1254411d5593ffe3c10b810bc7a) |
| `usb_policy.tenant` | [usb_policy.tenant](data-sources--fleet--reference--group-004.md#canonical-854f34c61d4f3a8b7de4f4840cb265af31481c91d72f819b3aaf9c3e6e8e02c9) |
| `volterra_software_version` | [volterra_software_version](data-sources--fleet--reference--group-001.md#canonical-44044520d8e1ed53c4d629dc70427b698332e06e710fe153b43c59c4e3c5b65a) |

<a id="canonical-0348efd053cbc1368b238881d2c2a1a3bb21865efcca24dc0ed4d8048b0bbcaf"></a>

## Next pages — Property reference / 7bbc950fdd9a / 15

- [allow_all_usb](data-sources--fleet--reference--group-001.md#canonical-48ce6fd2697b643d6cc6cf353a4861a2b25abf8b5d7f5aa320a99bb9769006d4)
- [blocked_services](data-sources--fleet--reference--group-001.md#canonical-831c4a8e2d5878bda105bfdecc47cf272091756efb6d21963401330750a9d5cd)
- [bond_device_list](data-sources--fleet--reference--group-002.md#canonical-65f237df10336d5db82d54cfe505146d90f609d8dfa0a04b5eb4d57217f22527)
- [dc_cluster_group](data-sources--fleet--reference--group-002.md#canonical-96db9f326526edd9bd9627855e7bea8fb3a49a23022b5ea37530be7c9fcee0ec)
- [dc_cluster_group_inside](data-sources--fleet--reference--group-002.md#canonical-459f9fc81a9ee6affc8dbdde520e801034d53cf449beca648f36288a2604774b)
- [default_config](data-sources--fleet--reference--group-002.md#canonical-6eb9aaec30d2e2cf90fa069c25811e7e00d5fa9a2c65b5052118698a69c96ae6)
- [default_sriov_interface](data-sources--fleet--reference--group-002.md#canonical-4f04b3ccf051a2819245908412155e72fb5342fe961059b254e3a05ed2e1bb22)
- [default_storage_class](data-sources--fleet--reference--group-002.md#canonical-9cca355dd9da1ba042afa381e3892933c464269ad176bfa370ab5c76359d7b89)
- [deny_all_usb](data-sources--fleet--reference--group-002.md#canonical-0a7984d4c53cce0a99415a554ce8e1658f8ac5ca16ccc67781060da9f8cdb325)
- [device_list](data-sources--fleet--reference--group-002.md#canonical-3a034fa01f0f5be5a0c60a43005276bc329a2ca951d98dbcace4874749947c7b)
- [disable_gpu](data-sources--fleet--reference--group-002.md#canonical-dcd5dc425e3d3f5f62fa1be75e9fa17e45e1b23561e3fec274ac1b08e24fd1a1)
- [disable_log_anonymization](data-sources--fleet--reference--group-002.md#canonical-c4ac51a9f536518ad770e67b50bd98210a3ba77fa9688997d2d66fba2c9aec6e)
- [disable_vm](data-sources--fleet--reference--group-002.md#canonical-bc405b9f6c12852ac10ce65f43cec97b5028acc1848f49dbd329ab63a6fe5802)
- [enable_gpu](data-sources--fleet--reference--group-002.md#canonical-343a0f4bbda728a37fc1135323420c28596c80d44352a60959ea457ef3585178)
- [enable_log_anonymization](data-sources--fleet--reference--group-002.md#canonical-e334ff3a795b467a7ba1700a237645e0f7c87496fd88a98d73de6044b3715e29)
- [enable_vgpu](data-sources--fleet--reference--group-002.md#canonical-4d479c596ef0edbbc7521d13615c04b2300386ef89701ddef222e8ee08ec76bb)
- [enable_vm](data-sources--fleet--reference--group-002.md#canonical-5753b2c2e22083ce920a71407a58f3aba1074413677b360bbd170dcd7c606c40)
- [inside_virtual_network](data-sources--fleet--reference--group-002.md#canonical-d43e2c692c15b906ace1b7b6dc086ef0308ca562dc0c1a37c9dc2952860969e4)
- [interface_list](data-sources--fleet--reference--group-002.md#canonical-bd78cb7abb7bf49babdb2d1f8412fa641d087aaea7e289e18bb4c2765a388299)
- [kubernetes_upgrade_drain](data-sources--fleet--reference--group-002.md#canonical-928e79ae9910b9b44010c3a7162b05584deee20b397abdd43f95a1c3e6ffdfa2)
- [log_receiver](data-sources--fleet--reference--group-002.md#canonical-61c444a8907092b4fee4bb68bf308a340f1dc48a3d55f8e8a78cb08a3785585e)
- [logs_streaming_disabled](data-sources--fleet--reference--group-002.md#canonical-89d2be778f765995174cabbf84408245d7744773b9bfbf012061a0edee7a6b4f)
- [network_connectors](data-sources--fleet--reference--group-002.md#canonical-db94ada9e04151546142e0446c43da3207c69f67a8be3b5e19b2101a5e11749b)
- [network_firewall](data-sources--fleet--reference--group-002.md#canonical-f98a897e6f410bf88c46126874860b12450d8d5d9cbacb244a50731d8089ebff)
- [no_bond_devices](data-sources--fleet--reference--group-002.md#canonical-edc12fb79e78cdd39a24a98d8a96c9da8c00a3b5d4ca24930245f5d909319a5a)
- [no_dc_cluster_group](data-sources--fleet--reference--group-002.md#canonical-29d0f16c609febb8dd2f0364cd7c673dde515b6724e3e901dae8ad0e73cb9105)
- [no_storage_device](data-sources--fleet--reference--group-002.md#canonical-6e5eed396a70e6fbfffcd17a91c1e00b247560f80beb44a2adc77ceee989f19f)
- [no_storage_interfaces](data-sources--fleet--reference--group-002.md#canonical-71b24f1ab3f8e13620efb2e10f7d1d07129cf3ad57f9db0c697c30f074c23710)
- [no_storage_static_routes](data-sources--fleet--reference--group-002.md#canonical-fc4925032f82914f326f0d8f65812a7152ea81f11666e910ae18af87eca39eb3)
- [outside_virtual_network](data-sources--fleet--reference--group-002.md#canonical-fa0cf1c7719192675c358757e047a074ea7122e0eecf767d9fe5eed505c48826)
- [performance_enhancement_mode](data-sources--fleet--reference--group-002.md#canonical-df949e88be9e74f148f4cc6b648395bc2c2966c58242f5e5278b65420dc4f3a5)
- [sriov_interfaces](data-sources--fleet--reference--group-002.md#canonical-5e5180d53b4f741b97d268311fbb857c0ca039ef57d462526616beaed68774c4)
- [storage_class_list](data-sources--fleet--reference--group-002.md#canonical-8b52e5aaf9f369d1a18bdcd5a9eb805473cd4387d010bcc639d0f4e2ceac9dd9)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-a360c3e38eec02c9655c8befc4e7cdb1edb0ff76dea91b874f16a0cd7d8b8349)
- [storage_interface_list](data-sources--fleet--reference--group-004.md#canonical-7e877450b08eec03e31927e555048ebdc52fb1f21283e08c037fbbe29d5cd3e0)
- [storage_static_routes](data-sources--fleet--reference--group-004.md#canonical-9a825351348828d7439fc399c86956cc54202bb679948c2a6cb2e089de057870)
- [usb_policy](data-sources--fleet--reference--group-004.md#canonical-9ec18332905de0a2ace95545b134c28e3348331c9500cac4450e177bac5b0c17)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-48ce6fd2697b643d6cc6cf353a4861a2b25abf8b5d7f5aa320a99bb9769006d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7f6d5c7d3a8a1d9f65ab6f22ab9f7013caefb32a003c6205064cd772cb1039e0"></a>

## allow_all_usb — allow_all_usb / 82f67aadfe14 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- allow_all_usb

<a id="canonical-894771c74fed127f34d6ee0a9d326e25f4cf9e34109cb73a14c868b69bb9dc3c"></a>

Type: `["object", {}]`. Computed.

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

- [allow_all_usb](data-sources--fleet--reference--group-001.md#canonical-894771c74fed127f34d6ee0a9d326e25f4cf9e34109cb73a14c868b69bb9dc3c)
- [deny_all_usb](data-sources--fleet--reference--group-002.md#canonical-3b86c27fe792b4846aef4b2443130cf2d32d8cfbfd2558b862ab0512f1e84b2e)
- [usb_policy](data-sources--fleet--reference--group-004.md#canonical-0d5e453c786da8ccd200294830086acbd0edd330738a8ba1b231459c978d547c)

Select alternatives according to the provider validators above.

<a id="canonical-612c15405b0815766721f19700a3b3725fdb51a6fda6092c35feedbc837919a7"></a>

## Direct properties — allow_all_usb / 82f67aadfe14 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-eded80cdd16dcf48545079c3df25da79938962ab259e98a8b4072054351edd19"></a>

## Next pages — allow_all_usb / 82f67aadfe14 / 4

- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-831c4a8e2d5878bda105bfdecc47cf272091756efb6d21963401330750a9d5cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-786a50e56f210cbacdb1013fbc6928edbad3588f423f3565500a8d69e721732d"></a>

## blocked_services — blocked_services / 31dfb9b07048 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- blocked_services

<a id="canonical-9c9620bdeab71579df50c1df124d492d42087ebc3300d493286618bbb688c865"></a>

Type: `"list"`. Computed.

Disable node local services on this site.

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

<a id="canonical-b93e748d1250901622a66206529c21ccc908ea771fc9fe8bea5d399bab5879c8"></a>

## Direct properties — blocked_services / 31dfb9b07048 / 3

- [dns](data-sources--fleet--reference--group-001.md#canonical-8877b1e00a33900053581642afe2bc37049427470801abd6e930df90415c98d7): complete subsection reference.

<a id="canonical-98d7b640f7c272a6f60cf6f2871d6c31d46d60753eafb9fd2d904f0143d7bbb4"></a>

<a id="canonical-afc9e762cfe4e21ad5d025de0234b0b5eb1c659f5d04c031b100b26b1b5b6ef7"></a>

## network_type property — blocked_services / 31dfb9b07048 / 4

Type: `"string"`. Computed.

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

- [ssh](data-sources--fleet--reference--group-002.md#canonical-04e77ddb6eeb4d8fe2c1fcb413919f6c97e8c1f436472078f59492d98f3a0b67): complete subsection reference.

- [web_user_interface](data-sources--fleet--reference--group-002.md#canonical-860d031d46fc5c15ec0ac0ae29310b2e0d899ba2060585f5ebfb148a46435441): complete subsection reference.

<a id="canonical-eb3c8d7b0c756e1d2b2ad8f20607d903459d214902cae51c2c1a0e356729232e"></a>

## Next pages — blocked_services / 31dfb9b07048 / 5

- [blocked_services.dns](data-sources--fleet--reference--group-001.md#canonical-8877b1e00a33900053581642afe2bc37049427470801abd6e930df90415c98d7)
- [blocked_services.ssh](data-sources--fleet--reference--group-002.md#canonical-04e77ddb6eeb4d8fe2c1fcb413919f6c97e8c1f436472078f59492d98f3a0b67)
- [blocked_services.web_user_interface](data-sources--fleet--reference--group-002.md#canonical-860d031d46fc5c15ec0ac0ae29310b2e0d899ba2060585f5ebfb148a46435441)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-8877b1e00a33900053581642afe2bc37049427470801abd6e930df90415c98d7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce32d8b9626fa328387ae9d1b5990f0a029fe7de07ab3b8a84cbacd55fcf9b30"></a>

## blocked_services.dns — blocked_services.dns / 38324c52fd93 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [blocked_services](data-sources--fleet--reference--group-001.md#canonical-831c4a8e2d5878bda105bfdecc47cf272091756efb6d21963401330750a9d5cd)
- blocked_services.dns

<a id="canonical-352e08899acf3e9a2b389b7308292eee56fb665089439e565f7f183fc2a57a9e"></a>

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

<a id="canonical-36bdd1f4cb1caadeda52177b0caa73fc556f1b43d3c0818d0f1b76b5b014c067"></a>

## Direct properties — blocked_services.dns / 38324c52fd93 / 3

This is an empty object or choice marker. It has no direct properties.
