---
page_title: "xcsh_gcp_vpc_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_gcp_vpc_site reference."
---

# xcsh_gcp_vpc_site reference

<a id="canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7dcd738114935b0ce059fd8b787e394dc77db4356313808f461f714b0a5996ce"></a>

## Property reference — Property reference / b4c55688522f / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- Property reference

<a id="canonical-e7512f8b743b81460d95f82896b251cc7eb7ff38892a20aba1cafa1ebac1646c"></a>

## Direct properties — Property reference / b4c55688522f / 3

<a id="canonical-b85f1d5502c006fc3bcc6602e8c3b3101feeb0e81a65671870f0cde4963b610e"></a>

<a id="canonical-eacbc2cb9994251ec1befe4d473d6c9e5f67363f3cda13bf8f797bcddbff2bc0"></a>

## address property — Property reference / b4c55688522f / 4

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

- [admin_password](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5a14f5870b3035dae459cdb64efb6ed9926a55eb75a13deb3148f3381938d385): complete subsection reference.

<a id="canonical-5b6bdb08bcd3e53f9542148cb807dcee7b9ad1b20a98f7aa0af3b5fde2a63a47"></a>

<a id="canonical-e8db1f1e174c5a4ba6faf223388ee127b3d422e643df6714e07bdcaf0f12b7d3"></a>

## annotations property — Property reference / b4c55688522f / 5

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

- [block_all_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1743bcd73450bc3dd30cbb6bf994b8c8c48c540ffbb2872e321b0ff1a581a1c6): complete subsection reference.

- [blocked_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-167e8ab000edb075a7e395fcf0e7b911b05439d69ce2b78ec25a808ae1153bf1): complete subsection reference.

- [cloud_credentials](data-sources--gcp_vpc_site--reference--group-001.md#canonical-95e08173bb2b66246ba892a158913465bb1a3e1f3e73312b25bb569b581e7f01): complete subsection reference.

- [coordinates](data-sources--gcp_vpc_site--reference--group-001.md#canonical-3a09fef7692d7d176ee26a7f5e0dfa507c1426702ad0dee764975478e642e824): complete subsection reference.

- [custom_dns](data-sources--gcp_vpc_site--reference--group-001.md#canonical-664e363d354ecd5b80118961ededf199212a347905ced87272319db78cdae8c8): complete subsection reference.

- [default_blocked_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-39ab1e4957001d2198ac5bcd38c773bd46dd67c9cc4a1d35fd047a526a137431): complete subsection reference.

<a id="canonical-e3d4bce2f2fc911af6c6767aacbfbfd4512eed8db212ff60ed969e0c508d8515"></a>

<a id="canonical-2ab2c151c16fa91d3b3ab9743ac76a9e172638e76206b880d644b0a1cf0807d6"></a>

## description property — Property reference / b4c55688522f / 6

Type: `"string"`. Computed.

Description of the GCPVPCSite.

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

- [disable_encryption](data-sources--gcp_vpc_site--reference--group-001.md#canonical-d0da01367080452c3a1b1fd1766dfd2520c59e731ace4092876fedfada95b872): complete subsection reference.

<a id="canonical-4c9ef34221aa70b906e201b03cb2bb7ce5f29748ac4a6427825e3e96742c3f7e"></a>

<a id="canonical-8ca50bd430fa18b70e16cf2902ba714d9e44e481c9be87cecc000798de78afd3"></a>

## disk_size property — Property reference / b4c55688522f / 7

Type: `"number"`. Computed.

Disk size to be used for this instance in GiB. 80 is 80 GiB.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64000,
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
    "ves.io.schema.rules.uint32.lte": "64000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "64000"
  }
}
```

- [enable_encryption](data-sources--gcp_vpc_site--reference--group-001.md#canonical-acd4c983a00cc6d63f7585e07639826ad2c3037b5203a366d2ac239abb58430f): complete subsection reference.

<a id="canonical-8e76137af5a74e6b8761ae3d0047e8f0345f03d99bd9cc2fe43347ff28ce0a69"></a>

<a id="canonical-87178e433182c879e979fb9e29660cef29f706d9e7df1bbdd4f3593a5c068b3d"></a>

## gcp_labels property — Property reference / b4c55688522f / 8

Type: `["map", "string"]`. Computed.

GCP Label is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in GCP console.

Upstream description:

GCP Label is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in GCP console.

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
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  }
}
```

<a id="canonical-3fa109a82c117f073f7571ba0930034d9f6d2ba3de2e3bd99d1e34b1323d02e5"></a>

<a id="canonical-91bafe9278214fdf64a66c712bef1d656085056e75d56495c4875feaa826a904"></a>

## gcp_region property — Property reference / b4c55688522f / 9

Type: `"string"`. Computed.

GCP Region. Name for GCP Region.

Upstream description:

Name for GCP Region.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
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

<a id="canonical-fd6ff81733cb9fb325a89accb2224fb0b5921b2231c1d0ed67d1f6f06eda8634"></a>

<a id="canonical-6dfcf7714bd7f383aa01dffb34285debb11fc71a29984c60df1633857b3437ef"></a>

## id property — Property reference / b4c55688522f / 10

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5bdfc213a88a3b5597f3d2bf1f9c923933afd2c4787d3747babb76e0bca0414d): complete subsection reference.

- [ingress_gw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-63c61e3d6c8d338e0f96c9d653534784b5ade2085751330ca6e32dd4d52485ea): complete subsection reference.

<a id="canonical-a7a106cad60291550dfe5e7cc079185b62bf0431163a4586080a277ab46143ab"></a>

<a id="canonical-3e9825598d841ffc5f0a6f29c297388f8007afe0b27c69cf514edeb1dba6e359"></a>

## instance_type property — Property reference / b4c55688522f / 11

Type: `"string"`. Computed.

Select Instance size based on performance needed.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
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

- [kubernetes_upgrade_drain](data-sources--gcp_vpc_site--reference--group-003.md#canonical-4a5d7773bbd1b78fbff21ae8236013e3de266b7fd703d86ac213753afc5d6638): complete subsection reference.

<a id="canonical-0c5cb37017a00519b24c360fe76f0c2bceabc0fefb24a09754eb06a572b83f9e"></a>

<a id="canonical-756ee9e19c833fcaadc0fe834fe414b20fe88ef36e4ee7c9eec69ab8752e5a94"></a>

## labels property — Property reference / b4c55688522f / 12

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

- [log_receiver](data-sources--gcp_vpc_site--reference--group-003.md#canonical-c62c46141ea3c05a0ddfe9ee1c3e3f5236aacb95d035b7d2379729dc9adaad66): complete subsection reference.

- [logs_streaming_disabled](data-sources--gcp_vpc_site--reference--group-003.md#canonical-a22274d6f2b4e4aa01fd0485ae5fec471b2949c4b413dae12f16e8d498728f73): complete subsection reference.

<a id="canonical-a5a69745334a5f4c62a72e1e2c525fddba3292bf7df727d2488d6edd02b27e9e"></a>

<a id="canonical-b6be871bbbcdaadf061f2aa83b8837d09c3821790e88c7b4870ec1ed2e36a736"></a>

## name property — Property reference / b4c55688522f / 13

Type: `"string"`. Required.

Name of the GCPVPCSite.

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

<a id="canonical-6c7008256eb62217b2725dfa7cf24321655387da00ff2a7c6fceb3f5cf961aea"></a>

<a id="canonical-cde74303f8ae57c570cb1bea80c4e5feefe591d89e6d34dd9855377fb48190b1"></a>

## namespace property — Property reference / b4c55688522f / 14

Type: `"string"`. Required.

Namespace where the GCPVPCSite exists.

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

- [offline_survivability_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-5407bdb8e6094e8c680518737f50555632401ffb1ef121d722e03c9a5d6c4536): complete subsection reference.

- [os](data-sources--gcp_vpc_site--reference--group-003.md#canonical-b663f6aa7a1c6c99cd1a7670e652cfc36d3df002f9b7a3b57049db2b9312d25e): complete subsection reference.

- [private_connect_disabled](data-sources--gcp_vpc_site--reference--group-003.md#canonical-a78509fd22231425759d40c893af69315775e0ffd8a5bc1c8ec9250f2acb496b): complete subsection reference.

- [private_connectivity](data-sources--gcp_vpc_site--reference--group-003.md#canonical-630f9c052b3a52854049dda1da13fd44492ebaa15d266601ef878f649412091a): complete subsection reference.

<a id="canonical-5d8399261041a098b5277e427f4424231c39ff0567b62011ea1cddce05231782"></a>

<a id="canonical-146bea7035e1a8c9cf0dfe7f0e24678833c834f3f53b28d367fbc8b0be254039"></a>

## ssh_key property — Property reference / b4c55688522f / 15

Type: `"string"`. Computed.

Public SSH key. Public SSH key for accessing the site.

Upstream description:

Public SSH key for accessing the site.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [sw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-34f2626bf35417bf94a0af238c2fe36603d4617d24af28174e43297681a6aa08): complete subsection reference.

- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-003.md#canonical-a7dfede66c6f88c6d8649f57e4468a404b29ff1dcbfc30fbb50f7fbfa3cd54a8): complete subsection reference.

- [waf_signatures](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0618f370940a8adff38a4bf27dd14fd5cd43884d9c0c76971bf3a53a4537d31a): complete subsection reference.

<a id="canonical-76f24fd71cb1f9ec5fc1260140be0cb714b7d82bd887fa7224e405201118c2ef"></a>

## All schema paths — Property reference / b4c55688522f / 16

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address` | [address](data-sources--gcp_vpc_site--reference--group-001.md#canonical-b85f1d5502c006fc3bcc6602e8c3b3101feeb0e81a65671870f0cde4963b610e) |
| `admin_password` | [admin_password](data-sources--gcp_vpc_site--reference--group-001.md#canonical-cfec52c430194d1ffa6620690099aa3b1d74e704ebba7d2267d263f15ca50f12) |
| `admin_password.blindfold_secret_info` | [admin_password.blindfold_secret_info](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1b8df325c33c1f572ec001fbec798770c46c8613993954135e78ed5d3578664b) |
| `admin_password.blindfold_secret_info.decryption_provider` | [admin_password.blindfold_secret_info.decryption_provider](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5fe5ec1db47ccf36d96ccb9724f2ff50fc9252b8d84d7c342283ab2bb1d6c00e) |
| `admin_password.blindfold_secret_info.location` | [admin_password.blindfold_secret_info.location](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1a05edc385bf0ee4d01c3e86aebdfee23f1e884727f7ddefe6b5d013b069d1ed) |
| `admin_password.blindfold_secret_info.store_provider` | [admin_password.blindfold_secret_info.store_provider](data-sources--gcp_vpc_site--reference--group-001.md#canonical-dd7de4b0da578506454d16c90a57c44e4f116e202564afcfe88f92d2ff5dc29e) |
| `admin_password.clear_secret_info` | [admin_password.clear_secret_info](data-sources--gcp_vpc_site--reference--group-001.md#canonical-ecdd36e50f734f26c3dea6d6d8a2553e7d4baef4da41a26e8a2cb17b97cfe6f5) |
| `admin_password.clear_secret_info.provider_ref` | [admin_password.clear_secret_info.provider_ref](data-sources--gcp_vpc_site--reference--group-001.md#canonical-b33eef66db46c5b6618cf3d2c78be03442c7264cb4c8c55b70f0088286cf5825) |
| `admin_password.clear_secret_info.url` | [admin_password.clear_secret_info.url](data-sources--gcp_vpc_site--reference--group-001.md#canonical-75aa40bf66007a9f7b80545a0ae08ae4c7cef206ebb5c181d0a15c409328db1d) |
| `annotations` | [annotations](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5b6bdb08bcd3e53f9542148cb807dcee7b9ad1b20a98f7aa0af3b5fde2a63a47) |
| `block_all_services` | [block_all_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-b97b0749d84ed8431ebb47d4fe461f3875a3dadb7b8f37398bc246b852ce4397) |
| `blocked_services` | [blocked_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-b05c226b6e8dc0e4c9a6e98613eafc9ea661945336f3154dae6264a33214ecca) |
| `blocked_services.blocked_service` | [blocked_services.blocked_service](data-sources--gcp_vpc_site--reference--group-001.md#canonical-a8f66547b5f643c159d718cd607cdd8a068a9aba2d054558f81cd572c3f1f4d6) |
| `blocked_services.blocked_service.dns` | [blocked_services.blocked_service.dns](data-sources--gcp_vpc_site--reference--group-001.md#canonical-92fd2350efab52c10b664154e5b6f6eb5dc3e8810da3b5518b386e7a09c61338) |
| `blocked_services.blocked_service.network_type` | [blocked_services.blocked_service.network_type](data-sources--gcp_vpc_site--reference--group-001.md#canonical-e97986549006948ea0d633d153706bd0c2bd131487f4bc86f6a1de3494ca1b9a) |
| `blocked_services.blocked_service.ssh` | [blocked_services.blocked_service.ssh](data-sources--gcp_vpc_site--reference--group-001.md#canonical-c5ac66f15a51e35566af2c5fa7edaf244803ec08f8a80d6c2983ce2c794e70bf) |
| `blocked_services.blocked_service.web_user_interface` | [blocked_services.blocked_service.web_user_interface](data-sources--gcp_vpc_site--reference--group-001.md#canonical-213f54b831d2132d0a23e680d372a10623efe2f49ac2897f0d6dafcae0e041ce) |
| `cloud_credentials` | [cloud_credentials](data-sources--gcp_vpc_site--reference--group-001.md#canonical-4d973002981b2f9fc5c86a7a880f1dd02a48f2c6327ceef9603d02bf2897d638) |
| `cloud_credentials.name` | [cloud_credentials.name](data-sources--gcp_vpc_site--reference--group-001.md#canonical-447b17a081109fed6397aec6979270bae14792ba61a5afca767a62cebfc2c579) |
| `cloud_credentials.namespace` | [cloud_credentials.namespace](data-sources--gcp_vpc_site--reference--group-001.md#canonical-47dd0304547e0bb825b699a3e65220b38434ae5083db4813bf7d38b0f1f6698a) |
| `cloud_credentials.tenant` | [cloud_credentials.tenant](data-sources--gcp_vpc_site--reference--group-001.md#canonical-e4c89512032e31fddca57183d17a761bdfd0c3276178355ffcc1c55bad92d9fa) |
| `coordinates` | [coordinates](data-sources--gcp_vpc_site--reference--group-001.md#canonical-28b4e170b6113670af0a8a8c36e8256374f9d4a7ae5844dce632545554d02773) |
| `coordinates.latitude` | [coordinates.latitude](data-sources--gcp_vpc_site--reference--group-001.md#canonical-38b3f37b1f6c0bf864b01f3e287c581fb466d4c6e215ad00d6460fdf2b93e8c7) |
| `coordinates.longitude` | [coordinates.longitude](data-sources--gcp_vpc_site--reference--group-001.md#canonical-54c0a70d00576ad1c638b64c3f320937bd1b440085f045862191bd09ffae1e2f) |
| `custom_dns` | [custom_dns](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0fa6367a2f6fb11a43df2a1de90fb4ba172f25bc5bdd63d6d5689cec45f80bac) |
| `custom_dns.inside_nameserver` | [custom_dns.inside_nameserver](data-sources--gcp_vpc_site--reference--group-001.md#canonical-a9e5ebbd1e605c9329e2c34396e00d4d090496fe61c79cd06c27560f73200ae1) |
| `custom_dns.outside_nameserver` | [custom_dns.outside_nameserver](data-sources--gcp_vpc_site--reference--group-001.md#canonical-86a9c8f0a2fb45cf15842fcefdc7d0a0e62cce606795304e276bb30b92f3c6e3) |
| `default_blocked_services` | [default_blocked_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-b197e4112f2b9172d1f211f632e6913450da6eee35c2c1008adb3e70c6ab6b7a) |
| `description` | [description](data-sources--gcp_vpc_site--reference--group-001.md#canonical-e3d4bce2f2fc911af6c6767aacbfbfd4512eed8db212ff60ed969e0c508d8515) |
| `disable_encryption` | [disable_encryption](data-sources--gcp_vpc_site--reference--group-001.md#canonical-4f58efe14f9f6273695bd62fbc1e53bff00356d47d9d071ed6166e7a66528810) |
| `disk_size` | [disk_size](data-sources--gcp_vpc_site--reference--group-001.md#canonical-4c9ef34221aa70b906e201b03cb2bb7ce5f29748ac4a6427825e3e96742c3f7e) |
| `enable_encryption` | [enable_encryption](data-sources--gcp_vpc_site--reference--group-001.md#canonical-419df70af1c1534230d6063fc16381b8c7979a951e7f2b7461019bd1627483ef) |
| `enable_encryption.kms_key_resource_id` | [enable_encryption.kms_key_resource_id](data-sources--gcp_vpc_site--reference--group-001.md#canonical-99043520f92d62fbbafe369c5a02b07faddfd459b586cf7111ddd2933a2478f7) |
| `enable_encryption.kms_key_ring_id` | [enable_encryption.kms_key_ring_id](data-sources--gcp_vpc_site--reference--group-001.md#canonical-44f69e75cffa229e69201b9c81ae9e9475cf003e0dcbe91b1d5af5945b8874ff) |
| `gcp_labels` | [gcp_labels](data-sources--gcp_vpc_site--reference--group-001.md#canonical-8e76137af5a74e6b8761ae3d0047e8f0345f03d99bd9cc2fe43347ff28ce0a69) |
| `gcp_region` | [gcp_region](data-sources--gcp_vpc_site--reference--group-001.md#canonical-3fa109a82c117f073f7571ba0930034d9f6d2ba3de2e3bd99d1e34b1323d02e5) |
| `id` | [id](data-sources--gcp_vpc_site--reference--group-001.md#canonical-fd6ff81733cb9fb325a89accb2224fb0b5921b2231c1d0ed67d1f6f06eda8634) |
| `ingress_egress_gw` | [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-371abe320a34379892077e9edd66cc32a98bd0311ff10c7da30dc8a2d3898231) |
| `ingress_egress_gw.active_enhanced_firewall_policies` | [ingress_egress_gw.active_enhanced_firewall_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-05a03994966a32ee8aa9c7b81fbf0ec21fa55f42d28d73b366c6750f88832777) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-e21fd247a492f42298a6afd78d35d8f016aa6d9cdf1d3704b47906d822c961fb) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.name](data-sources--gcp_vpc_site--reference--group-002.md#canonical-86342c248dcf07948bc47ed681ea61381d6e329dd318546fa67bb852ea9bdf0e) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](data-sources--gcp_vpc_site--reference--group-002.md#canonical-90591bd37bb949f1d8f8d5d19659b42abca2739f8396f2a8ee467f03174fca08) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](data-sources--gcp_vpc_site--reference--group-002.md#canonical-a834d3db9a9ffdf29f9a502148bc8e5bb0fa1c10d45dacd0387ced148b3eb5a2) |
| `ingress_egress_gw.active_forward_proxy_policies` | [ingress_egress_gw.active_forward_proxy_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-695079c894bbd6c096c9a4e636809aaffa2a05792859392000fa545f0c19060b) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-b45bc352d8ca781cd1d3517ff19ee983c7ae152b17490331c0592537b3f82e8a) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.name` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.name](data-sources--gcp_vpc_site--reference--group-002.md#canonical-de48f7fc82bd5c0829593d6c83af95f03d9fa2853d3067dbce386600ecdb6ea5) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.namespace` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.namespace](data-sources--gcp_vpc_site--reference--group-002.md#canonical-94403ec8e661a0993c488309b39738744da2fbfd0d7f727206ecfb55d8728140) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.tenant` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.tenant](data-sources--gcp_vpc_site--reference--group-002.md#canonical-a377c3f8309cbe6b06d94f7159b4d143cd645ee3797eccd0bd05427d384d93fc) |
| `ingress_egress_gw.active_network_policies` | [ingress_egress_gw.active_network_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-ad550dfbb72183a60f769d7c1b4a1465b72cf3aa83fd5fa56bb93405d429a67a) |
| `ingress_egress_gw.active_network_policies.network_policies` | [ingress_egress_gw.active_network_policies.network_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-67082a3e1d68f19d456ceca3940c8bfc4e83bdc5f85c708b5b304c2a24d8e78f) |
| `ingress_egress_gw.active_network_policies.network_policies.name` | [ingress_egress_gw.active_network_policies.network_policies.name](data-sources--gcp_vpc_site--reference--group-002.md#canonical-ebb5dc688acbe752bab5eb942f33d6c80afcf7cd93e82052d3f1ab2b8310035e) |
| `ingress_egress_gw.active_network_policies.network_policies.namespace` | [ingress_egress_gw.active_network_policies.network_policies.namespace](data-sources--gcp_vpc_site--reference--group-002.md#canonical-b92870c991ef97a7c1105c86a5dbc6080110293423eb9948ff91c24c27d35d9e) |
| `ingress_egress_gw.active_network_policies.network_policies.tenant` | [ingress_egress_gw.active_network_policies.network_policies.tenant](data-sources--gcp_vpc_site--reference--group-002.md#canonical-4aa1df16e37a4d4c16990a2f5d560c7747b137f285036bb61f0bc71f5287ba45) |
| `ingress_egress_gw.dc_cluster_group_inside_vn` | [ingress_egress_gw.dc_cluster_group_inside_vn](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1a924c5234dca1bc77e14a93e9367829381bb2ce3e8a4cfa248a7142c4bac920) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.name` | [ingress_egress_gw.dc_cluster_group_inside_vn.name](data-sources--gcp_vpc_site--reference--group-002.md#canonical-c52e2ece37a1a6f49bc86a90974c2b4ee9c7a9ab27f126ce7736e4238a2d2690) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.namespace` | [ingress_egress_gw.dc_cluster_group_inside_vn.namespace](data-sources--gcp_vpc_site--reference--group-002.md#canonical-528523e6ebac663e77681aa893e904e2276c39ff3d165facf07c008852f9deff) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.tenant` | [ingress_egress_gw.dc_cluster_group_inside_vn.tenant](data-sources--gcp_vpc_site--reference--group-002.md#canonical-bfa15e3b7f4782d4000f0d2712af9e1b9e21d8c79e29ac5bbfc0353e7a731290) |
| `ingress_egress_gw.dc_cluster_group_outside_vn` | [ingress_egress_gw.dc_cluster_group_outside_vn](data-sources--gcp_vpc_site--reference--group-002.md#canonical-8e321be65fd21a8b50d3bb935f74aebe5f06667fa6a68e55a2d30a370a8d1f95) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.name` | [ingress_egress_gw.dc_cluster_group_outside_vn.name](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2d7951f19297b7d69e4a330949714d90592e32ffc9b2bd6f461e92b4cc9243a4) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.namespace` | [ingress_egress_gw.dc_cluster_group_outside_vn.namespace](data-sources--gcp_vpc_site--reference--group-002.md#canonical-ab981348fbdd21b6083bef7e9e762692eed5605445bb69e2dc267a566401e8e2) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.tenant` | [ingress_egress_gw.dc_cluster_group_outside_vn.tenant](data-sources--gcp_vpc_site--reference--group-002.md#canonical-7e7058607a5320fb11b75ec4fd5524c7f2d1ec567deb4f8a5cdc29210b056380) |
| `ingress_egress_gw.forward_proxy_allow_all` | [ingress_egress_gw.forward_proxy_allow_all](data-sources--gcp_vpc_site--reference--group-002.md#canonical-46ac9fc7cf1e1531ef6b3a73d486fa5153d07fdbfb0201cdaad68f656b930e1e) |
| `ingress_egress_gw.gcp_certified_hw` | [ingress_egress_gw.gcp_certified_hw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-275f2fc528afae81c7cc474ec21f6846f4433a733f5fc0fd481139d80dcc1dcb) |
| `ingress_egress_gw.gcp_zone_names` | [ingress_egress_gw.gcp_zone_names](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1334c4be72ad9d685fdddad85b7d421c3f92d6429e21625f8b92d1898602242b) |
| `ingress_egress_gw.global_network_list` | [ingress_egress_gw.global_network_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1e34f1280a3660785e0a7e879f23b074e71a3aa8df3b64f16be695a2746d621e) |
| `ingress_egress_gw.global_network_list.global_network_connections` | [ingress_egress_gw.global_network_list.global_network_connections](data-sources--gcp_vpc_site--reference--group-002.md#canonical-ea69a499b17b889d4569dcdb05922e63500d41e554bd5e1150b689354d495103) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](data-sources--gcp_vpc_site--reference--group-002.md#canonical-bd0dc4bfb1ebd0e1169c4fafe374c0d5e4ba4c75145961cd2128a5a009ab0c69) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn](data-sources--gcp_vpc_site--reference--group-002.md#canonical-f24d788c8f0f1022f296e6cbb8339f3f35125eba0d55ed7aa7f10d248028a995) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](data-sources--gcp_vpc_site--reference--group-002.md#canonical-10cb41e8451dc8e2f8c2cd2a827592fc8c411759b42efc56cc0058181081623f) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](data-sources--gcp_vpc_site--reference--group-002.md#canonical-a2d61c0c189a88493783cb8367fa61d39c9736ba7c9b0462eb2ded4428221be4) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](data-sources--gcp_vpc_site--reference--group-002.md#canonical-00317e50e47910f1b376ccda2359f28eed541de4daa9e540d5068682d47a17cd) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](data-sources--gcp_vpc_site--reference--group-002.md#canonical-79c39b69e191c2259c8026f3766f56f28170d67a0798cf6ff54fda960eb7441b) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn](data-sources--gcp_vpc_site--reference--group-002.md#canonical-772c6457054cee60eb2aa4be010d069b4f1bd72e18fa9f5b1fecd0db121f8a75) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](data-sources--gcp_vpc_site--reference--group-002.md#canonical-9ee42a0d162c55151f1b783f050ea034cb42a1bcc1c7e2c8477a24995224547c) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](data-sources--gcp_vpc_site--reference--group-002.md#canonical-f7b3c8e147ad819282494130d1bdb4bd603a22b96cd63e4729ad4c7164288e94) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](data-sources--gcp_vpc_site--reference--group-002.md#canonical-5f3d48c9286cf0026862a05bcd405cb2b2b558d67065852bfcb03414b9a1ce6a) |
| `ingress_egress_gw.inside_network` | [ingress_egress_gw.inside_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1095eb74f4381e5b47493a55e9a210bccd81b177eab98cac730cd02e29237ecd) |
| `ingress_egress_gw.inside_network.existing_network` | [ingress_egress_gw.inside_network.existing_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-68653692d3671744ed2c5b11f377396e979f0cb64c41686ebf994994dc92b923) |
| `ingress_egress_gw.inside_network.existing_network.name` | [ingress_egress_gw.inside_network.existing_network.name](data-sources--gcp_vpc_site--reference--group-002.md#canonical-508fc0f2639d4a9253025f1ece25e3ec760a6eedb7a912f3994afae93db0c9dc) |
| `ingress_egress_gw.inside_network.new_network` | [ingress_egress_gw.inside_network.new_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-a4c15ca1215f2a5e1e0025bd75dd4552bb7a7350f4b4266e4f363d398a4a3664) |
| `ingress_egress_gw.inside_network.new_network.name` | [ingress_egress_gw.inside_network.new_network.name](data-sources--gcp_vpc_site--reference--group-002.md#canonical-4a0b386c12e9036dc5ed4966b27cc6e3f61ff3feda9e8dec91fb94084caa24a2) |
| `ingress_egress_gw.inside_network.new_network_autogenerate` | [ingress_egress_gw.inside_network.new_network_autogenerate](data-sources--gcp_vpc_site--reference--group-002.md#canonical-a9474624d3111b8d738d98365d5aa3f421f1aa98d8b203bddf09a7adc0f67f28) |
| `ingress_egress_gw.inside_static_routes` | [ingress_egress_gw.inside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3cc81644002ebf0b75e9e942d4f2b7c3e4c9b8641db116604f03de9d8ff0f397) |
| `ingress_egress_gw.inside_static_routes.static_route_list` | [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-5f526675a65834d200d9559925d1e472e3ba8791cc8d581a0551865899931964) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-050d4f240974661669494c8eff01f5dccb3f3b0b8fda2682a92e959b5a7251bd) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.attrs` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.attrs](data-sources--gcp_vpc_site--reference--group-002.md#canonical-59946752ce1b15bc22e6f0acaece02c9b522640f5590e2c0f2765d8ea5d0bc30) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels](data-sources--gcp_vpc_site--reference--group-002.md#canonical-15ec07705e35d9d2b3f2149442d2b4fd9858014d0b35ea121c798ef9a1fc3123) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-002.md#canonical-a43f7130957498f5dfd7bda305b2578dd8fa70ca326bf0cdad8fbc19050de83b) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--gcp_vpc_site--reference--group-002.md#canonical-910a363a14ed8c22131ea8c21dc7aab36e5dcd5bae0e2279c8adf7f590b57e48) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](data-sources--gcp_vpc_site--reference--group-002.md#canonical-4bdf8a89d33792fc15b3dadedec0c102c140b70f7903916cf92f2c1ed8ac39d0) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](data-sources--gcp_vpc_site--reference--group-002.md#canonical-8868d99f49132191a4658172ac07734cb8e9912f67d29a0c794f0dd518717d54) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1dabe761431a8e81c8500281effa66ef517dc1471fbfc9d24366cb6c4878bf4d) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](data-sources--gcp_vpc_site--reference--group-002.md#canonical-5ee468ff3be71dc6081b15aa1d6d2ebb06dbaa6c6c035bc6494cc21dcfd437b5) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](data-sources--gcp_vpc_site--reference--group-002.md#canonical-54664304f118e9ffbbec3944ea486610585635ec5dd9439d5d2ba60d36d91f28) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-002.md#canonical-91cb23b20bfec404512cdf2ed963792ddb4de0d3bbd0801a5da683aa6039a51a) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--gcp_vpc_site--reference--group-002.md#canonical-4f160074ff1ac8f5c611ed19f3a9db05f0a193a44f75c713dc069983aa08ea71) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--gcp_vpc_site--reference--group-002.md#canonical-f3698bcefa1e3f1b3bc7630ada04a3eecde193123ba2632af16e82c84edba57f) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](data-sources--gcp_vpc_site--reference--group-002.md#canonical-54291a2acfd9228bae798fb2fbf37d553a4af6714e967367f961805077ce4043) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--gcp_vpc_site--reference--group-002.md#canonical-f93147a9bbe02754b0c13abc2bca0815d8b8824874b71f030ba76c82e3f5e1ba) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1b78f94f929e144a4f4497320b07779226c5b2fb9540ded7070410b25e0e1f47) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--gcp_vpc_site--reference--group-002.md#canonical-9af6efd871227dcec0d80b805179e832e81bfaf15a6435f39a684295d76a32d7) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](data-sources--gcp_vpc_site--reference--group-002.md#canonical-b912a6e7df06a628174187c6a21209dc7010dd9a67082becce5d6ff02e736778) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--gcp_vpc_site--reference--group-002.md#canonical-332d9623dfc2975a668d30f66d68e24664d555085e5b771be0a37628f0316a1e) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](data-sources--gcp_vpc_site--reference--group-002.md#canonical-f82134761559ae1674d67d286758a788ae1b8bd1cc92c46128493b64e0451943) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.type` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.type](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3f2c36e49ff6d1236056a28a0a20fdbd533ee41272b3531c54847fc7ecd92bd5) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--gcp_vpc_site--reference--group-002.md#canonical-9c68edb0aec394c24945a10ac9f177503c164d9a70048d85377803b4c7c8113c) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3ce8d84e63cb24dfd9646cd20ab2630f90fcc5399dc5b1fb862b68a6a25f958a) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](data-sources--gcp_vpc_site--reference--group-002.md#canonical-388c92d57ceaea7544d8b35a26bf1e70fb08990cc5f2a69f08038c0fdbb01ecf) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](data-sources--gcp_vpc_site--reference--group-002.md#canonical-a1babbec8a27eb2de16fb4515688368c4b91760b1e3c15b81262fd5057c5d43b) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--gcp_vpc_site--reference--group-002.md#canonical-a7bbbc9331fbbb6531b213084400b16df162d13eca14e20dce69cc89f2dda0be) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](data-sources--gcp_vpc_site--reference--group-002.md#canonical-86e8e329cb53565c93f8888912678694405b312adff97bd59f54d6d4d3c910a4) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](data-sources--gcp_vpc_site--reference--group-002.md#canonical-5fe78c8733b567110b5427f4c4bfc7289ca5c0ba475119266c43a8cca9516f85) |
| `ingress_egress_gw.inside_static_routes.static_route_list.simple_static_route` | [ingress_egress_gw.inside_static_routes.static_route_list.simple_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-417248f611466b45355441e8a84728eeaf494b28ae25f8febfb690748fa552a5) |
| `ingress_egress_gw.inside_subnet` | [ingress_egress_gw.inside_subnet](data-sources--gcp_vpc_site--reference--group-002.md#canonical-edf3cff42f3a00883878eaeea6692e69291f3d06287009d8023c92e8380a7e0f) |
| `ingress_egress_gw.inside_subnet.existing_subnet` | [ingress_egress_gw.inside_subnet.existing_subnet](data-sources--gcp_vpc_site--reference--group-002.md#canonical-b1129518208d91f6a9766dcdc64c99f6546c6a1405e3db2e1b51e3fdf7711281) |
| `ingress_egress_gw.inside_subnet.existing_subnet.subnet_name` | [ingress_egress_gw.inside_subnet.existing_subnet.subnet_name](data-sources--gcp_vpc_site--reference--group-002.md#canonical-4a2f22c68eb1b4c4866e323d5b29e32f3746e84c612e1e125d37f4393ce8fee2) |
| `ingress_egress_gw.inside_subnet.new_subnet` | [ingress_egress_gw.inside_subnet.new_subnet](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1ef9f3ba39e15136ba507e45b9df8e70bb5d17ccac6cce32d5e54e7c4434d6be) |
| `ingress_egress_gw.inside_subnet.new_subnet.primary_ipv4` | [ingress_egress_gw.inside_subnet.new_subnet.primary_ipv4](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1fdb77363b45a85d9359ace84a08c86182eec1f8d6eee49aa7417dcf601faf44) |
| `ingress_egress_gw.inside_subnet.new_subnet.subnet_name` | [ingress_egress_gw.inside_subnet.new_subnet.subnet_name](data-sources--gcp_vpc_site--reference--group-002.md#canonical-a86fb737bf9acf78894e3abd533c65b516a1b42297d361f946ba0087ea3399a0) |
| `ingress_egress_gw.no_dc_cluster_group` | [ingress_egress_gw.no_dc_cluster_group](data-sources--gcp_vpc_site--reference--group-002.md#canonical-e7505cdef5ba384fb3ac127911ad7f6b8108948552c17478334803d8b5e9e723) |
| `ingress_egress_gw.no_forward_proxy` | [ingress_egress_gw.no_forward_proxy](data-sources--gcp_vpc_site--reference--group-002.md#canonical-4c57c00d474b52b6639c6a59c4822a3603614af89b02b9f98497c50e44266734) |
| `ingress_egress_gw.no_global_network` | [ingress_egress_gw.no_global_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-e36b964c3d923758b486d24f172b1c4cc4b5b6bdfb64dfaa3cc7237eb89837e7) |
| `ingress_egress_gw.no_inside_static_routes` | [ingress_egress_gw.no_inside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-e9f823778476916c6cb62176221d84d27f29761807e360a5d203797280f506cb) |
| `ingress_egress_gw.no_network_policy` | [ingress_egress_gw.no_network_policy](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1e6a245d8879b514970c06d2aba373a6347f730f10d77c54037ff051dfc59cf2) |
| `ingress_egress_gw.no_outside_static_routes` | [ingress_egress_gw.no_outside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3f8d4e3c5e6345d4fe1b03e0673ce591a5e1b6aa6cc5cfe5ac534ac76744b3b0) |
| `ingress_egress_gw.node_number` | [ingress_egress_gw.node_number](data-sources--gcp_vpc_site--reference--group-001.md#canonical-2fac4443ae821e9585781dad2d8094f27954b69957df0d55ea606dc94c5b1b68) |
| `ingress_egress_gw.outside_network` | [ingress_egress_gw.outside_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-378b14164134bb2dff46a192e27e5ca4f37f8ffe166bb4bd5246fe0f191fcd76) |
| `ingress_egress_gw.outside_network.existing_network` | [ingress_egress_gw.outside_network.existing_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-94294a61b82bc3f9d4c1437e181d4a2960fb8890113790186515eaa33460aae2) |
| `ingress_egress_gw.outside_network.existing_network.name` | [ingress_egress_gw.outside_network.existing_network.name](data-sources--gcp_vpc_site--reference--group-002.md#canonical-5b8aafc0098fad4ed2c884bb4c540a03c4cad9f50ea44639ac77e11c216a9616) |
| `ingress_egress_gw.outside_network.new_network` | [ingress_egress_gw.outside_network.new_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-d4c8e92bdff247cda25b933181122fe9e408fab4fdcbccc1530e7106f6ad8fd6) |
| `ingress_egress_gw.outside_network.new_network.name` | [ingress_egress_gw.outside_network.new_network.name](data-sources--gcp_vpc_site--reference--group-002.md#canonical-75958271c7a2c55b23bbe9e406d41701c7390f8b8c2347ddab1798d7beaea358) |
| `ingress_egress_gw.outside_network.new_network_autogenerate` | [ingress_egress_gw.outside_network.new_network_autogenerate](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0d5675c71310cb85030e2e53de10c15e2d1a9e163dc283a99891648c3f3173bd) |
| `ingress_egress_gw.outside_static_routes` | [ingress_egress_gw.outside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3b2eeceb9996429df03cad504e77ececc160ea2279d99216624ae35536026151) |
| `ingress_egress_gw.outside_static_routes.static_route_list` | [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-4e201a63e54caf395571852da5c60749406459f5ea1ad203d37510347afeb8ac) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-93837faa5d2f3ed9aade6e38b35fd70d7034eceb96985767cdc5bed2a047d7c4) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.attrs` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.attrs](data-sources--gcp_vpc_site--reference--group-002.md#canonical-5725fa9e2fabc6602a2ea74533467b2e96e381800ee34ebfa2e828ce8008ec4e) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels](data-sources--gcp_vpc_site--reference--group-002.md#canonical-a35656c172d451a332cd127fc70d76952834df45bfbcbaa3d5fa640c2cc1cad7) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-002.md#canonical-f6326f9a12642d58313240a8409e6af66728ad6b443cfddb5a4c631520a8dccd) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--gcp_vpc_site--reference--group-003.md#canonical-cc4443798a349f1e548b8c703e5bde762fa33032f232fdd59876298c2a345373) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](data-sources--gcp_vpc_site--reference--group-003.md#canonical-a098bb8421eae256fdbe283076069ef8ee6df65a0abee7ba5a26cf91c2e34c21) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](data-sources--gcp_vpc_site--reference--group-003.md#canonical-24aaa55969e8714d13c0f9be262ea01c4f29bbaf2c1dd2c9b5017f134ff6e42b) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1357929b2f1a0a98c8b67066d7846497e8659247a94eeb044f404089f15e6b57) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](data-sources--gcp_vpc_site--reference--group-003.md#canonical-dfa26edc92420f21404a4f51896695d0d1e6afcb6b70ed1a60a2ddc898036390) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](data-sources--gcp_vpc_site--reference--group-003.md#canonical-f0d25a826ca744a7fc279b7e6ea9d70cf4b0b51a24f33aa0ec7c5acaab43f942) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-003.md#canonical-98f63b7663675cc67c4095d6324b0b7a1bc7f08f359c905271e26b237cc929e0) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--gcp_vpc_site--reference--group-003.md#canonical-532479bb3b78e08dee2fd35f54bd4bc3486fcacc034bd8f5cf7933d04d6653f9) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--gcp_vpc_site--reference--group-003.md#canonical-e595ac628b1b066ed2fb2d02a1b2583258a6995392bbdd6ea5753c90cb4195b5) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](data-sources--gcp_vpc_site--reference--group-003.md#canonical-f1660ecaf116c77ac2163d30b07470880720bba037a729c386240de8124dd1e0) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--gcp_vpc_site--reference--group-003.md#canonical-20c91dd3c0a30d5bca734fe8119f34af51a03583cdeec41e1806b35c72d7e84c) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](data-sources--gcp_vpc_site--reference--group-003.md#canonical-5d49670ae98a5110c82b751219d3b85d876e82171c088053a6d0955fae9b3ad5) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--gcp_vpc_site--reference--group-003.md#canonical-400e187c21a9f3f8ca85da9e39bf433115d897c373b4d62b5060283087ea8f82) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1f3c45ba1d2c8ec8febeb9f0413ee126b2e969395caf87f80723bf49c75aebd4) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--gcp_vpc_site--reference--group-003.md#canonical-482dfa21a591c0f15d9bcedd5fe8799b119464862fcca4ef7639fbe51b0ee417) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](data-sources--gcp_vpc_site--reference--group-003.md#canonical-368a1944df5047a9f961020b4d56bca54434c762a130cacf5fd21ce5d72e411c) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.type](data-sources--gcp_vpc_site--reference--group-002.md#canonical-b55a2a59414b74165fc8d4fce5b95d1aa584408706d6ddaa54c711c4bec6d073) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--gcp_vpc_site--reference--group-003.md#canonical-4b7abb8186e32a0e625819ae39950addc60585b04a250492dd5a1239de3d5a21) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--gcp_vpc_site--reference--group-003.md#canonical-6408d232945f1227213dcc82c1e6dceb9394b830f7597d72ed81ab58fd58b89a) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](data-sources--gcp_vpc_site--reference--group-003.md#canonical-2b36ee5f56beba20d8e2731a89b40677e18ae8c79b67f9c9f9bb193d04a9d696) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](data-sources--gcp_vpc_site--reference--group-003.md#canonical-a5c51127f2d6a58c9522bb14379f2e825f5b2951b84ff49482a14ab1f5b6ef5a) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0f3c5531a7991a5129f7691bbef199e59eff8db9a7e6ff24069659fd67d710c1) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](data-sources--gcp_vpc_site--reference--group-003.md#canonical-b816fbcbc551dd01bc15e87094c9739570f695db9045116b1f89134ffa290ab1) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](data-sources--gcp_vpc_site--reference--group-003.md#canonical-fa7ed146e7908f04ce5eb09648c789f9baf18fdaf49a4d95c28951803df814da) |
| `ingress_egress_gw.outside_static_routes.static_route_list.simple_static_route` | [ingress_egress_gw.outside_static_routes.static_route_list.simple_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-be3319687ed05d04383ab02399da3e3347c1829512ca4af852831a8990d72c1f) |
| `ingress_egress_gw.outside_subnet` | [ingress_egress_gw.outside_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-f8cd88576dd543f6e6c4635eec228efab99d1fe1e00133dd6d9ac53cc015f164) |
| `ingress_egress_gw.outside_subnet.existing_subnet` | [ingress_egress_gw.outside_subnet.existing_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-2e590eb4c0b9a99d53d427f12f43a990fdfaf2c317e74ef2a4f9bb5ec2fe439a) |
| `ingress_egress_gw.outside_subnet.existing_subnet.subnet_name` | [ingress_egress_gw.outside_subnet.existing_subnet.subnet_name](data-sources--gcp_vpc_site--reference--group-003.md#canonical-8575e574ac51d637ba5eafceb2953770f71c69f51f4c000f202ee8335950ae0c) |
| `ingress_egress_gw.outside_subnet.new_subnet` | [ingress_egress_gw.outside_subnet.new_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-63c30f2a5257aa53764f388468b82d885bdef7b154dbfee063bdae593ad3b120) |
| `ingress_egress_gw.outside_subnet.new_subnet.primary_ipv4` | [ingress_egress_gw.outside_subnet.new_subnet.primary_ipv4](data-sources--gcp_vpc_site--reference--group-003.md#canonical-51f8c2d20891b9172f8cec6ebb67d09a744570d48d806fe6a77d2bc6de65c8b0) |
| `ingress_egress_gw.outside_subnet.new_subnet.subnet_name` | [ingress_egress_gw.outside_subnet.new_subnet.subnet_name](data-sources--gcp_vpc_site--reference--group-003.md#canonical-00ee8bcfd545c7c7fb1016886ab0018cf89875a1ad9b32e58a9a46e27646e689) |
| `ingress_egress_gw.performance_enhancement_mode` | [ingress_egress_gw.performance_enhancement_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-e2dd85b81d5ea2aa318efca5076e763ee383a0d30b324d42cfc738f5a9dc90f4) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-385cd85551365076c4afd174713bbe375dd44f887c52c7fa9c809e4da4d2bdfa) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--gcp_vpc_site--reference--group-003.md#canonical-96cfe6d2ff2be8a0192c64f01b122efabd421d9d02d17635aefafdacd7edfdaf) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--gcp_vpc_site--reference--group-003.md#canonical-e1626ec32003f9c77147036a69112cb126d027da9357a9e86f34c1c70b953002) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0df98a05d13cf6e16b6b1b290efb96e2a8c20ee6e83ab51ce5827fdee57f70e6) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--gcp_vpc_site--reference--group-003.md#canonical-793e6bbdc9a8f4dbfb53752d393b830adf9f637da98f63f57c2f9912a77b7dd8) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1471b51cee8f9bd043a7157daa97c12b14e8cb9e32a1e37c6320e7123a9f2aef) |
| `ingress_egress_gw.sm_connection_public_ip` | [ingress_egress_gw.sm_connection_public_ip](data-sources--gcp_vpc_site--reference--group-003.md#canonical-495f375688a6c5b31a9672585e65b2262cf4989e12c71c3f1891ed570e5fe520) |
| `ingress_egress_gw.sm_connection_pvt_ip` | [ingress_egress_gw.sm_connection_pvt_ip](data-sources--gcp_vpc_site--reference--group-003.md#canonical-c7e967a8c5731a941213c6e7b2b5406eab34c033fec9d9d3ecc151c328aeed15) |
| `ingress_gw` | [ingress_gw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-aabb0996e5fa997bf2ff3a50da2ecedf568fefa97b3dfc132843c1eefe4bc0dc) |
| `ingress_gw.gcp_certified_hw` | [ingress_gw.gcp_certified_hw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-4ca378f5cc102c5d5eb3e2769f111e397f76449a2e4f07aff46258668df58cd4) |
| `ingress_gw.gcp_zone_names` | [ingress_gw.gcp_zone_names](data-sources--gcp_vpc_site--reference--group-003.md#canonical-18b44f50ba201e42d866cf1ac629e82305911431b2017fc995aa3ed6019ce17a) |
| `ingress_gw.local_network` | [ingress_gw.local_network](data-sources--gcp_vpc_site--reference--group-003.md#canonical-b025cdbf474731cd9ba166b29cd99553722a0d7260bbe247c24e52a655f5929f) |
| `ingress_gw.local_network.existing_network` | [ingress_gw.local_network.existing_network](data-sources--gcp_vpc_site--reference--group-003.md#canonical-4ab8c1fd4888c773bcf535b66259a1868d2d1bb72fad4420a514fb997fd2444a) |
| `ingress_gw.local_network.existing_network.name` | [ingress_gw.local_network.existing_network.name](data-sources--gcp_vpc_site--reference--group-003.md#canonical-534b9d23c7eb7db66df8d672533ab22911c9d19f7c236f28363d9cf2126c5b37) |
| `ingress_gw.local_network.new_network` | [ingress_gw.local_network.new_network](data-sources--gcp_vpc_site--reference--group-003.md#canonical-4b7e1c8124d8ec4fa29192d47a25f4248266b7be3d24a2f59255c64d5ebaa241) |
| `ingress_gw.local_network.new_network.name` | [ingress_gw.local_network.new_network.name](data-sources--gcp_vpc_site--reference--group-003.md#canonical-c5f131d442b44b489f95f2ee24e92c52233849511846b8f3f1388cc645af546a) |
| `ingress_gw.local_network.new_network_autogenerate` | [ingress_gw.local_network.new_network_autogenerate](data-sources--gcp_vpc_site--reference--group-003.md#canonical-564b9cc8c81b6e2693d91afe2c7ce4f4a0aef152b0d28cf1c35326ba0e8cd97a) |
| `ingress_gw.local_subnet` | [ingress_gw.local_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-69fe675f863b2c9706dc2dc1e81a0c909db1bfceb5fb4771bbf3b23635a4022d) |
| `ingress_gw.local_subnet.existing_subnet` | [ingress_gw.local_subnet.existing_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-652bf7196b5762a019a27c826fda740cb24210d768ce384b4147511380e331b9) |
| `ingress_gw.local_subnet.existing_subnet.subnet_name` | [ingress_gw.local_subnet.existing_subnet.subnet_name](data-sources--gcp_vpc_site--reference--group-003.md#canonical-c242c72be6af2f8fa94148a4ce554cbd8177ac694218e0f854cbf945572d7325) |
| `ingress_gw.local_subnet.new_subnet` | [ingress_gw.local_subnet.new_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1193621c01047293f2c50ad9eb6485619935f0423054050fc805f37cf9416faf) |
| `ingress_gw.local_subnet.new_subnet.primary_ipv4` | [ingress_gw.local_subnet.new_subnet.primary_ipv4](data-sources--gcp_vpc_site--reference--group-003.md#canonical-da43d4368b225a9b835f1155f5e2ce0f251f1baa11faa130262bb86de4e4c86d) |
| `ingress_gw.local_subnet.new_subnet.subnet_name` | [ingress_gw.local_subnet.new_subnet.subnet_name](data-sources--gcp_vpc_site--reference--group-003.md#canonical-8f7daf3074b1c812b866a26850c233a6fe2d137192dd37161cd0cd3d30b65b8a) |
| `ingress_gw.node_number` | [ingress_gw.node_number](data-sources--gcp_vpc_site--reference--group-003.md#canonical-c90b4e62337c696962c069c38f1a68cd5521dda9b2fd100b0e5fefd11b8c05f5) |
| `ingress_gw.performance_enhancement_mode` | [ingress_gw.performance_enhancement_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-4784c4b989e0feec8b060049ea862eaf9d3383da8ec72eee744e1578082e6a91) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-4066cdbc71961c565ee75a558ed7189ed2af27656ba89da06a089ca0d4cd067b) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--gcp_vpc_site--reference--group-003.md#canonical-c50a382c4cf8027d1cf5060ccb70c52f361f0e33e078d83d32dd390347e44ca1) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--gcp_vpc_site--reference--group-003.md#canonical-924dd473405c42b6fa8415572ad42742f032dd5ab40802a34f5422c3a1815ed1) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3783daa7c50d7b9be1a08bb5fb6426a6b32781d544732f9fc6d48d98dd3cc0dc) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--gcp_vpc_site--reference--group-003.md#canonical-c4482003092e40fc90ba0d777308ea875138be775fa8034e9bd965895d12cd61) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--gcp_vpc_site--reference--group-003.md#canonical-adc78edef91a1e59df0b398c03a709fe80ffaf17fc7d78bb3a3689e8a31fec4c) |
| `instance_type` | [instance_type](data-sources--gcp_vpc_site--reference--group-001.md#canonical-a7a106cad60291550dfe5e7cc079185b62bf0431163a4586080a277ab46143ab) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0268e772a08896abca785d965a1fd0d7d6a1eb88536dcbe6988097065a964f34) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](data-sources--gcp_vpc_site--reference--group-003.md#canonical-f159937ac70275d80158bfe5807991339f3ae51fc4a495674bcd50beb14f7bab) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--gcp_vpc_site--reference--group-003.md#canonical-342fe8ee401299b451d973348f86b8b1ec97ab448830d5ff62d302029b356a85) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-5d4a035171a30a1c5dfc3d58dd7ec5d5c66ac73cdb6fe89f55edf7f1c0f0be8b) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](data-sources--gcp_vpc_site--reference--group-003.md#canonical-9a9461b1a3fb747c8ec94a179e6a2c9d3958edf9d39b9d25767bd2d1168f6a81) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](data-sources--gcp_vpc_site--reference--group-003.md#canonical-96e7b67e87ab5ea93261a30ca2d2a4a4eff9d87449c9e56e0a53bfcbf64eab63) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](data-sources--gcp_vpc_site--reference--group-003.md#canonical-40d830e91e3e58daf4b3b66d90363bcfabf4bcd28441f235b1017c29915d6eb6) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-ee928bd26fbe59cf2300eeff39543e8469959772a8c5a889d19ec7c2f4426283) |
| `labels` | [labels](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0c5cb37017a00519b24c360fe76f0c2bceabc0fefb24a09754eb06a572b83f9e) |
| `log_receiver` | [log_receiver](data-sources--gcp_vpc_site--reference--group-003.md#canonical-fdee6744d558d4e71fe3e34712018c2f7c35dcc1f9574e2c207e17cd9ef7943b) |
| `log_receiver.name` | [log_receiver.name](data-sources--gcp_vpc_site--reference--group-003.md#canonical-ec1c479950b739d137fe786a1fa3e6f6a9336e624d2723574e5d7efad1369c70) |
| `log_receiver.namespace` | [log_receiver.namespace](data-sources--gcp_vpc_site--reference--group-003.md#canonical-6960d7beb0e2596a51c671730654f8d789faac138d9d0bc913fe06e79c23623e) |
| `log_receiver.tenant` | [log_receiver.tenant](data-sources--gcp_vpc_site--reference--group-003.md#canonical-d116bc4137e040463ae11fb35d3251c01b7a13fe9b0caa1e15d7fda06488a27a) |
| `logs_streaming_disabled` | [logs_streaming_disabled](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1548d29124796eb96edf7e8c04c533c717c5f95b4cbf3898d0d80d3ba00c0cc1) |
| `name` | [name](data-sources--gcp_vpc_site--reference--group-001.md#canonical-a5a69745334a5f4c62a72e1e2c525fddba3292bf7df727d2488d6edd02b27e9e) |
| `namespace` | [namespace](data-sources--gcp_vpc_site--reference--group-001.md#canonical-6c7008256eb62217b2725dfa7cf24321655387da00ff2a7c6fceb3f5cf961aea) |
| `offline_survivability_mode` | [offline_survivability_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-6de2af680f24fe7199319dfc3a733cc421376ba82a48b7fa0473c5758b4ea04d) |
| `offline_survivability_mode.enable_offline_survivability_mode` | [offline_survivability_mode.enable_offline_survivability_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-efb7a9580421794aa919fe6bb3975a662c68b6267a222abcec69762734158b70) |
| `offline_survivability_mode.no_offline_survivability_mode` | [offline_survivability_mode.no_offline_survivability_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-e9781e3460c28bd8d66fac4c8af31c886de8825d6bcbe2f96b9173b505c6949d) |
| `os` | [os](data-sources--gcp_vpc_site--reference--group-003.md#canonical-022c88b9af45074655f9cbda2b60a4e804068f95b48ae7f759da2a1bb98cfd1e) |
| `os.default_os_version` | [os.default_os_version](data-sources--gcp_vpc_site--reference--group-003.md#canonical-4448d6b432dec4648caeafd7fb34b2b265aba59658bf1a840984ad29efe4e601) |
| `os.operating_system_version` | [os.operating_system_version](data-sources--gcp_vpc_site--reference--group-003.md#canonical-73dc8ae8ceb28857e0a535c823d66b28ae7fba5a464e78ca932c37c0e3a05b1a) |
| `private_connect_disabled` | [private_connect_disabled](data-sources--gcp_vpc_site--reference--group-003.md#canonical-ed8be2ab70fa1176e184fddc1a39b746916aae75e68558a2ec04589ad34ee73f) |
| `private_connectivity` | [private_connectivity](data-sources--gcp_vpc_site--reference--group-003.md#canonical-d14b31de9b35e994d69a640b2e408bbd04ccafcf3ce9143cf4858355c9ae3f9f) |
| `private_connectivity.cloud_link` | [private_connectivity.cloud_link](data-sources--gcp_vpc_site--reference--group-003.md#canonical-6b1adc8c43e25f15d237b16aeafecf46e49104839df0d9f760ee13048c56be42) |
| `private_connectivity.cloud_link.name` | [private_connectivity.cloud_link.name](data-sources--gcp_vpc_site--reference--group-003.md#canonical-09ec659c7daa34d7722eb8467c2396b139751456e5754b8ec18f069b5d9b9851) |
| `private_connectivity.cloud_link.namespace` | [private_connectivity.cloud_link.namespace](data-sources--gcp_vpc_site--reference--group-003.md#canonical-e8fbd459fbaad54b4c21c5de821b0a6525831c7abb60c3ba28e2db289fc05e61) |
| `private_connectivity.cloud_link.tenant` | [private_connectivity.cloud_link.tenant](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1d802fe6c857c31fb3e3d4f6f2f62dc387ea260e362af02621cea6399b7f6fa1) |
| `private_connectivity.inside` | [private_connectivity.inside](data-sources--gcp_vpc_site--reference--group-003.md#canonical-c1901975a1a05b8599634952b603530ba72302db7fe8f7c3306f47b66664e78c) |
| `private_connectivity.outside` | [private_connectivity.outside](data-sources--gcp_vpc_site--reference--group-003.md#canonical-e038781147b5823e39e40b0db7dd6817328b1ede39d224923b0a8f3b23e49b86) |
| `ssh_key` | [ssh_key](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5d8399261041a098b5277e427f4424231c39ff0567b62011ea1cddce05231782) |
| `sw` | [sw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-2d11c7b1ecae586c5d9a973c05d5a34a18633255a86052b4454efe063ceb9f65) |
| `sw.default_sw_version` | [sw.default_sw_version](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1ebcfee274f7425599b968996267fae3813db4272088c7cbd192cf0f3f9b5d32) |
| `sw.volterra_software_version` | [sw.volterra_software_version](data-sources--gcp_vpc_site--reference--group-003.md#canonical-a3b6eb8b63e4cd0cd39084c8fa2a6f17ab89a426ad486386231507414a836a82) |
| `voltstack_cluster` | [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-003.md#canonical-237e3ceae77303258b3ff82422e1aa07d05f60ce403476754ef00c324a4d28d1) |
| `voltstack_cluster.active_enhanced_firewall_policies` | [voltstack_cluster.active_enhanced_firewall_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-71d1e790d858fac38911f3a66e0df589b2419c1398e07ae895d31dd41c983726) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-bffa9723a4e692ac767f684e0684d6c2ebdf0e2f1a60ff785af97ec879a9d445) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.name](data-sources--gcp_vpc_site--reference--group-004.md#canonical-5d7d6ffabdcbb88feae95daa7ca98c0b74ade4cd32cbc557e29394dbbc1985fa) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](data-sources--gcp_vpc_site--reference--group-004.md#canonical-cb8b3716803709c418f39ede5022b3b6f6d0afaba48b2ebb9045105f4af781c6) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](data-sources--gcp_vpc_site--reference--group-004.md#canonical-bd27545a1e5585d61eaed41bec74872995957555a863ff986cb3b2adf58d6dd8) |
| `voltstack_cluster.active_forward_proxy_policies` | [voltstack_cluster.active_forward_proxy_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-409bcfc1bf91ef9800349c6a2ffec1285057b927bee87f471dae1127f32dba0e) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-aa8233eb19f3b5e19437c609dadf633331b4f17d490c4ea0521f0da459f71906) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.name` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.name](data-sources--gcp_vpc_site--reference--group-004.md#canonical-d9fda8ab41a221cbb74f1d48b7ec2818f9c41413f63bd49c5e71789f1c4053cd) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.namespace` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.namespace](data-sources--gcp_vpc_site--reference--group-004.md#canonical-e7257b6665347375cab1b0fca9d7b55c1b9d403a998128807b5fdb39f5e2ab9c) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.tenant` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.tenant](data-sources--gcp_vpc_site--reference--group-004.md#canonical-06d113de096352d5a3d74e863ea577b103ee94457269bc136be9f3f2e6e3d4b7) |
| `voltstack_cluster.active_network_policies` | [voltstack_cluster.active_network_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-42255d6dceb316748439d5f0eb7a7226aaf52baefd6a49be58ae32abded8dbb0) |
| `voltstack_cluster.active_network_policies.network_policies` | [voltstack_cluster.active_network_policies.network_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-42e94722c1cd8ae5d6e48fd31037513824e8dcc5dd7f988a243f88ac2e58d253) |
| `voltstack_cluster.active_network_policies.network_policies.name` | [voltstack_cluster.active_network_policies.network_policies.name](data-sources--gcp_vpc_site--reference--group-004.md#canonical-6773d1302ddb84160d5229ba99518346f74698f4fed1ddfac6f69ea3debe6245) |
| `voltstack_cluster.active_network_policies.network_policies.namespace` | [voltstack_cluster.active_network_policies.network_policies.namespace](data-sources--gcp_vpc_site--reference--group-004.md#canonical-5216b6ac6a153786ca6cfa2a6cd2c92d2a8a08b5e10477a0ebe97b27b28cafb4) |
| `voltstack_cluster.active_network_policies.network_policies.tenant` | [voltstack_cluster.active_network_policies.network_policies.tenant](data-sources--gcp_vpc_site--reference--group-004.md#canonical-b85b460c213ea2e705a6c6141cd0c6669ed8758ae6c4e9c0aabc1f5b690c5d2f) |
| `voltstack_cluster.dc_cluster_group` | [voltstack_cluster.dc_cluster_group](data-sources--gcp_vpc_site--reference--group-004.md#canonical-8418173cee498a6c0c962d4b2f88c336cb91f29b0c7f01b1dcf6f214a57a434a) |
| `voltstack_cluster.dc_cluster_group.name` | [voltstack_cluster.dc_cluster_group.name](data-sources--gcp_vpc_site--reference--group-004.md#canonical-53eba679b91c4ac387d239c9db5243996070d22630f72c3b2fb4f14aecede78f) |
| `voltstack_cluster.dc_cluster_group.namespace` | [voltstack_cluster.dc_cluster_group.namespace](data-sources--gcp_vpc_site--reference--group-004.md#canonical-b08401a48d073162b6214c4ba94abc62b90f4c5fc233d14f8c43665a6050ea6c) |
| `voltstack_cluster.dc_cluster_group.tenant` | [voltstack_cluster.dc_cluster_group.tenant](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0f7dee6ca366a982ae8eba36d22bd533579690d925312a9b7fd2266fd1b3cc55) |
| `voltstack_cluster.default_storage` | [voltstack_cluster.default_storage](data-sources--gcp_vpc_site--reference--group-004.md#canonical-32d09ffb7908e93c60725040fdb07ac7b4fbb4b336aa15a81942e391b7f4b3c0) |
| `voltstack_cluster.forward_proxy_allow_all` | [voltstack_cluster.forward_proxy_allow_all](data-sources--gcp_vpc_site--reference--group-004.md#canonical-f0b1fbed0617293d56779102cc1966b87664b2f5a97a655e52db6e4b51cb5c90) |
| `voltstack_cluster.gcp_certified_hw` | [voltstack_cluster.gcp_certified_hw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-fcb47aaef9dec15c0ee7ef6d678d2a101157b556ca6f197bc2d83a2d01559c87) |
| `voltstack_cluster.gcp_zone_names` | [voltstack_cluster.gcp_zone_names](data-sources--gcp_vpc_site--reference--group-003.md#canonical-126757094d19ebc76dc55516f6711b7a76e35af16a3ab8f722ed64ecc88f6c49) |
| `voltstack_cluster.global_network_list` | [voltstack_cluster.global_network_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-bb948fd1bd46e853cdb62be9118bf53bf5fec9b3ec3817794527086d6807bc19) |
| `voltstack_cluster.global_network_list.global_network_connections` | [voltstack_cluster.global_network_list.global_network_connections](data-sources--gcp_vpc_site--reference--group-004.md#canonical-bb260699566d9f1d87ba824fb7090a465208069d2dce318c491b29b95aa82487) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](data-sources--gcp_vpc_site--reference--group-004.md#canonical-c5992c2efea95d6a7d3dd1d3e0039d3ac7569486b915992848f65d78f586c490) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn](data-sources--gcp_vpc_site--reference--group-004.md#canonical-9b6da37224b6741edac36bb35555ea018e25657f110154effde8706524e3606e) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](data-sources--gcp_vpc_site--reference--group-004.md#canonical-17eb64729954fe4cb0851a3a80134cc08f7ab93adf32e1be49f9477db53205c2) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](data-sources--gcp_vpc_site--reference--group-004.md#canonical-bde0a295ed81890a60dd9273ce5c7cf785d690d82330215961d5d84df0d2b3cf) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](data-sources--gcp_vpc_site--reference--group-004.md#canonical-603e46bf76a1084d1e7b985d60430b60206d4c7514ec2f7a2d88b6e28592ce03) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](data-sources--gcp_vpc_site--reference--group-004.md#canonical-bf1a9f6f78f61ab98a70f51b27892fd29304f9df24a8b7cd1e42fc2f3a643253) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn](data-sources--gcp_vpc_site--reference--group-004.md#canonical-9292922b6c80b3c5ba99c24809d42ef2958f01afb2fe46fc1086105eb2bc7b37) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](data-sources--gcp_vpc_site--reference--group-004.md#canonical-a8a64611f9e66fc17cbaa36d67ecb71deb50e89e8207985a84d40f21d0cc679c) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](data-sources--gcp_vpc_site--reference--group-004.md#canonical-d1687523a80920238df5b0f028d180c3751f39d1741ceeeed40c8ac449523555) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](data-sources--gcp_vpc_site--reference--group-004.md#canonical-066c5f1a9320c418e58866db99dcb3999c2b660ff3594d66bab2a1357c764176) |
| `voltstack_cluster.k8s_cluster` | [voltstack_cluster.k8s_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-c6ff6096fef5e7be601ace773095cf66aedab5763847292897beac312abdb7f6) |
| `voltstack_cluster.k8s_cluster.name` | [voltstack_cluster.k8s_cluster.name](data-sources--gcp_vpc_site--reference--group-004.md#canonical-6d05afe89c0013f8a2405ccc8effb71904bba823c7bc8e07f31e2aa0e356451b) |
| `voltstack_cluster.k8s_cluster.namespace` | [voltstack_cluster.k8s_cluster.namespace](data-sources--gcp_vpc_site--reference--group-004.md#canonical-ba04609b24a4d5dfc2a97af5fbec642980117c1da093965afcc470090b26461f) |
| `voltstack_cluster.k8s_cluster.tenant` | [voltstack_cluster.k8s_cluster.tenant](data-sources--gcp_vpc_site--reference--group-004.md#canonical-14ec8994054ecf2009f6576edd630bb7711806b04bcd56b9346c936cbe96304c) |
| `voltstack_cluster.no_dc_cluster_group` | [voltstack_cluster.no_dc_cluster_group](data-sources--gcp_vpc_site--reference--group-004.md#canonical-bd55948660557d981fac93e599b4564adbdc13c080c3acd1f676cb60213534a4) |
| `voltstack_cluster.no_forward_proxy` | [voltstack_cluster.no_forward_proxy](data-sources--gcp_vpc_site--reference--group-004.md#canonical-21cfe78a22285ecf0ae5551f58fb4bd3a47175dd676661403071fbd1e86257b3) |
| `voltstack_cluster.no_global_network` | [voltstack_cluster.no_global_network](data-sources--gcp_vpc_site--reference--group-004.md#canonical-47c593da3f38dac84324e871d18fb3d597ff1db67bf24ee3cd88242af235a4ff) |
| `voltstack_cluster.no_k8s_cluster` | [voltstack_cluster.no_k8s_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-e4f454c5141b9aae5c5a9674c572dc0e811259b07870d94cd97266279226ce79) |
| `voltstack_cluster.no_network_policy` | [voltstack_cluster.no_network_policy](data-sources--gcp_vpc_site--reference--group-004.md#canonical-953fef9187bd2259a74e81db15617f670091b386b2a4e586e6033c55e4b39997) |
| `voltstack_cluster.no_outside_static_routes` | [voltstack_cluster.no_outside_static_routes](data-sources--gcp_vpc_site--reference--group-004.md#canonical-a487ca88b1506ee22cb2c83197f7778b3a2c9e58cde173855dec5b2f559b7315) |
| `voltstack_cluster.node_number` | [voltstack_cluster.node_number](data-sources--gcp_vpc_site--reference--group-003.md#canonical-d70324685a22392dd31d1ba2ad760eb91ee8da1e465ed73b3ebd834168d4d687) |
| `voltstack_cluster.outside_static_routes` | [voltstack_cluster.outside_static_routes](data-sources--gcp_vpc_site--reference--group-004.md#canonical-c986a471ee76ade54afe6fdcb7087863e0952690f5bb5ba67fa2c26600c394e3) |
| `voltstack_cluster.outside_static_routes.static_route_list` | [voltstack_cluster.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1e8fa133c81d849259ecfa1ebe0548a867ed2ced7977839c4be4d07704acfb25) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-004.md#canonical-ba2d89acdd32557a5dc1cd7e0d2efa1ceb5ca2d38a8ce6cbdb09183c317bca9c) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.attrs` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.attrs](data-sources--gcp_vpc_site--reference--group-004.md#canonical-40497016d72b5fa1c09d83514811c7fae0fe838a3ad6665dc154410d6a674ba6) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels](data-sources--gcp_vpc_site--reference--group-004.md#canonical-c5b4dbd922ead2105e2a4c084da269d2b760d7dd84c3f2977d3b479ae1cb354e) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-004.md#canonical-6292a78de288880a5d112b09387d130a768989582c023f4ac508ea4c7c20bc03) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--gcp_vpc_site--reference--group-004.md#canonical-c6caf5849f161c8ebc8820b3f65b84e3504f8b4da61367bdbebad61fdf68d371) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](data-sources--gcp_vpc_site--reference--group-004.md#canonical-27d8324e701c5d8dad38e343f71f710b6b451e3c79ac3f70d312dc150275657b) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](data-sources--gcp_vpc_site--reference--group-004.md#canonical-58f992abb1b94f413c6098bffe996044881dc469916ade54479b920e9000de9d) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](data-sources--gcp_vpc_site--reference--group-004.md#canonical-965adf13840a02343667b93083507802b19a45e5b808d70d18e917af0a7611d1) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](data-sources--gcp_vpc_site--reference--group-004.md#canonical-4cd78c1abce66f4258fa87acb5282146cad0c3395cb1a0295d1c5ac48eb7f9ff) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](data-sources--gcp_vpc_site--reference--group-004.md#canonical-c824b535e2c130c74a77f2aceb68c3fd621e27056e70713702e2af45366ba3ad) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-004.md#canonical-759d1b70cf2e5d040f7b9c4a3a83713cd440447db769b5c5f6f78796e6c38331) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0de2231757646d68648aacb74f8d481495cc045fb0dda826a9b337ab8656ea09) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--gcp_vpc_site--reference--group-004.md#canonical-4181fe6f50d415bd8602f8f8ad5283645e8470ef70ebf8718d9549d3b3f9db37) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1c3975a086b6ca7d878e7e1c6bcca2781fa60a83b78e2a24a5ba1c8f4e86334e) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--gcp_vpc_site--reference--group-004.md#canonical-fea1211c351343f00d9398281b6219c6dd92d92353e886bb831337278d5141a8) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](data-sources--gcp_vpc_site--reference--group-004.md#canonical-ee0712011e7fee49e755a37308284e418fe0cf81e0027c78cdce4a128bce2800) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0ca1e09d7d0b0e08967471dd83efd9815cbff103cea03ae672a98f9253343d69) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](data-sources--gcp_vpc_site--reference--group-004.md#canonical-19af68a368bbfe54fdc90b30ab471113661b6b06f9e881058e81667f8bbd0f07) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--gcp_vpc_site--reference--group-004.md#canonical-06e66031145223c2d0162e120c4cd6248c9059fdaecf495402d2d0ab671b0895) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](data-sources--gcp_vpc_site--reference--group-004.md#canonical-9f0d42bdbc8b1ff1d562318763604badc85926ef3d69545b6745c604565d2a55) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.type](data-sources--gcp_vpc_site--reference--group-004.md#canonical-e6801fb35c6e22040999098c045db36a261fe4c5e9adfc3da8291f2b1075310a) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--gcp_vpc_site--reference--group-004.md#canonical-f137532331ca4af31b3445c6ba9719623a81e8b17b4f5b3618ac8f3dac86a63f) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--gcp_vpc_site--reference--group-004.md#canonical-c38fbbd76a5ebbc025a221bc6efa615234e4176d622014b5fbf39d6ed3255835) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](data-sources--gcp_vpc_site--reference--group-004.md#canonical-f59687a562a74e45772be0a53e2627b575933b57b5600968d59df7411ff67bfb) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](data-sources--gcp_vpc_site--reference--group-004.md#canonical-7b734309fb06ce4d87b9b97c81365eafcf66b7c31671d1af7ba89ca74ccaabf8) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--gcp_vpc_site--reference--group-004.md#canonical-b934a5ec0d87181aa30dca916eedceef237c0210ecfdaae581761de596b53570) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](data-sources--gcp_vpc_site--reference--group-004.md#canonical-f3b99c41da25c5fdbf441a375598b0765b4c7dd236426306403175baee7e62c3) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](data-sources--gcp_vpc_site--reference--group-004.md#canonical-e44e49bf9c6a4a219865e42df6de59e62c48de7e042d79a23c62604e9d0b8755) |
| `voltstack_cluster.outside_static_routes.static_route_list.simple_static_route` | [voltstack_cluster.outside_static_routes.static_route_list.simple_static_route](data-sources--gcp_vpc_site--reference--group-004.md#canonical-a09d383a48f957f78ea8e25937cf06c6fa6acf16ecd3b216587d7e0e0579e623) |
| `voltstack_cluster.site_local_network` | [voltstack_cluster.site_local_network](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0d75a74e2a9fa87f73e6810603a053e9b11118d2d00cb71e93cb6fe4332c62c8) |
| `voltstack_cluster.site_local_network.existing_network` | [voltstack_cluster.site_local_network.existing_network](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2cbe9fb0660f5fa9384b203f8aad30c87147cfbea245920aa3910d934391af88) |
| `voltstack_cluster.site_local_network.existing_network.name` | [voltstack_cluster.site_local_network.existing_network.name](data-sources--gcp_vpc_site--reference--group-004.md#canonical-b8e755082d59bf7f76b4fc1544f4b68a8787e00b75d9f1df9c2bf45702ea6986) |
| `voltstack_cluster.site_local_network.new_network` | [voltstack_cluster.site_local_network.new_network](data-sources--gcp_vpc_site--reference--group-004.md#canonical-a18c42a0b56bf8c1f19efba0a22fa696ffa322ca27b3c2e8807b1dc31d431439) |
| `voltstack_cluster.site_local_network.new_network.name` | [voltstack_cluster.site_local_network.new_network.name](data-sources--gcp_vpc_site--reference--group-004.md#canonical-33f0ce13c9088d9c248cd2a6701addaf55e80a56752a32355d8533f2a99d31e9) |
| `voltstack_cluster.site_local_network.new_network_autogenerate` | [voltstack_cluster.site_local_network.new_network_autogenerate](data-sources--gcp_vpc_site--reference--group-004.md#canonical-8af6e840c804ca22b33c5a5cc14dcabc3a2468212b475d321883e4f15a62d942) |
| `voltstack_cluster.site_local_subnet` | [voltstack_cluster.site_local_subnet](data-sources--gcp_vpc_site--reference--group-004.md#canonical-950a96446963dc41f6ba30fd145dc82aa607860dede00000f5aff3f72eb9e0eb) |
| `voltstack_cluster.site_local_subnet.existing_subnet` | [voltstack_cluster.site_local_subnet.existing_subnet](data-sources--gcp_vpc_site--reference--group-004.md#canonical-40f8cc9f383f273c3b397b2522fb90cd17ed33716cf5df3d237f130142cea6a5) |
| `voltstack_cluster.site_local_subnet.existing_subnet.subnet_name` | [voltstack_cluster.site_local_subnet.existing_subnet.subnet_name](data-sources--gcp_vpc_site--reference--group-004.md#canonical-722bd6c0d3aff8991fac96447376a873d1f82e2790ca178d6737b323aa70bb15) |
| `voltstack_cluster.site_local_subnet.new_subnet` | [voltstack_cluster.site_local_subnet.new_subnet](data-sources--gcp_vpc_site--reference--group-004.md#canonical-9771b0a43ae32b497fa38df9d1fb586a034cd1fda10be661ca6dcc00a8771b7f) |
| `voltstack_cluster.site_local_subnet.new_subnet.primary_ipv4` | [voltstack_cluster.site_local_subnet.new_subnet.primary_ipv4](data-sources--gcp_vpc_site--reference--group-004.md#canonical-e8c4564d365ffda15a890e02fd8b83f3564ad44c469ebf709c9098bfd4377441) |
| `voltstack_cluster.site_local_subnet.new_subnet.subnet_name` | [voltstack_cluster.site_local_subnet.new_subnet.subnet_name](data-sources--gcp_vpc_site--reference--group-004.md#canonical-4226b83ec56ce2eb083551b8479b785e0fa1e0d22a3ba182ae5b75691befdaf5) |
| `voltstack_cluster.sm_connection_public_ip` | [voltstack_cluster.sm_connection_public_ip](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1ff04572c4387bcc8066168ad72465d566bb50a0adee6242989e9d5d91f78c24) |
| `voltstack_cluster.sm_connection_pvt_ip` | [voltstack_cluster.sm_connection_pvt_ip](data-sources--gcp_vpc_site--reference--group-004.md#canonical-85b5935ab640aae1c6976ba50de01864c7112a63a87a41454c13a78e2b6374bf) |
| `voltstack_cluster.storage_class_list` | [voltstack_cluster.storage_class_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-9100b1735650a07faa29af34dc0aa2b37fc18416f426a7a28d6975a23c811c61) |
| `voltstack_cluster.storage_class_list.storage_classes` | [voltstack_cluster.storage_class_list.storage_classes](data-sources--gcp_vpc_site--reference--group-004.md#canonical-35cc5b1ff6da1144ab206ab037a4863b75dbcd97ec436b8cf7348d9bb739e89f) |
| `voltstack_cluster.storage_class_list.storage_classes.default_storage_class` | [voltstack_cluster.storage_class_list.storage_classes.default_storage_class](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2a46c857191708ee7cd315ac9a05b1553bb66239758def03f9d38bbd46a2adbc) |
| `voltstack_cluster.storage_class_list.storage_classes.storage_class_name` | [voltstack_cluster.storage_class_list.storage_classes.storage_class_name](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3254100e4adff776f98597d2abe498aa6295438802f11eeee64cd5408c317b1f) |
| `waf_signatures` | [waf_signatures](data-sources--gcp_vpc_site--reference--group-004.md#canonical-7e82e8e9097c7d7e32c8dd20ad82317e982094a178ac708c722b7477cbfc0eb3) |
| `waf_signatures.automatic` | [waf_signatures.automatic](data-sources--gcp_vpc_site--reference--group-004.md#canonical-5fdb73e1600c6075c808512d51d56c6f7c2e75ddba2db82b964e1d2663918a1a) |
| `waf_signatures.manual` | [waf_signatures.manual](data-sources--gcp_vpc_site--reference--group-004.md#canonical-84c2fddc4b125b35c705e8fa458087b23283cd159b83b352f1c63977e12946a6) |

<a id="canonical-266fe342975d6d65df39d728a179f43e2f6eba804fbbf27f04174dff574293fa"></a>

## Next pages — Property reference / b4c55688522f / 17

- [admin_password](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5a14f5870b3035dae459cdb64efb6ed9926a55eb75a13deb3148f3381938d385)
- [block_all_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1743bcd73450bc3dd30cbb6bf994b8c8c48c540ffbb2872e321b0ff1a581a1c6)
- [blocked_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-167e8ab000edb075a7e395fcf0e7b911b05439d69ce2b78ec25a808ae1153bf1)
- [cloud_credentials](data-sources--gcp_vpc_site--reference--group-001.md#canonical-95e08173bb2b66246ba892a158913465bb1a3e1f3e73312b25bb569b581e7f01)
- [coordinates](data-sources--gcp_vpc_site--reference--group-001.md#canonical-3a09fef7692d7d176ee26a7f5e0dfa507c1426702ad0dee764975478e642e824)
- [custom_dns](data-sources--gcp_vpc_site--reference--group-001.md#canonical-664e363d354ecd5b80118961ededf199212a347905ced87272319db78cdae8c8)
- [default_blocked_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-39ab1e4957001d2198ac5bcd38c773bd46dd67c9cc4a1d35fd047a526a137431)
- [disable_encryption](data-sources--gcp_vpc_site--reference--group-001.md#canonical-d0da01367080452c3a1b1fd1766dfd2520c59e731ace4092876fedfada95b872)
- [enable_encryption](data-sources--gcp_vpc_site--reference--group-001.md#canonical-acd4c983a00cc6d63f7585e07639826ad2c3037b5203a366d2ac239abb58430f)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5bdfc213a88a3b5597f3d2bf1f9c923933afd2c4787d3747babb76e0bca0414d)
- [ingress_gw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-63c61e3d6c8d338e0f96c9d653534784b5ade2085751330ca6e32dd4d52485ea)
- [kubernetes_upgrade_drain](data-sources--gcp_vpc_site--reference--group-003.md#canonical-4a5d7773bbd1b78fbff21ae8236013e3de266b7fd703d86ac213753afc5d6638)
- [log_receiver](data-sources--gcp_vpc_site--reference--group-003.md#canonical-c62c46141ea3c05a0ddfe9ee1c3e3f5236aacb95d035b7d2379729dc9adaad66)
- [logs_streaming_disabled](data-sources--gcp_vpc_site--reference--group-003.md#canonical-a22274d6f2b4e4aa01fd0485ae5fec471b2949c4b413dae12f16e8d498728f73)
- [offline_survivability_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-5407bdb8e6094e8c680518737f50555632401ffb1ef121d722e03c9a5d6c4536)
- [os](data-sources--gcp_vpc_site--reference--group-003.md#canonical-b663f6aa7a1c6c99cd1a7670e652cfc36d3df002f9b7a3b57049db2b9312d25e)
- [private_connect_disabled](data-sources--gcp_vpc_site--reference--group-003.md#canonical-a78509fd22231425759d40c893af69315775e0ffd8a5bc1c8ec9250f2acb496b)
- [private_connectivity](data-sources--gcp_vpc_site--reference--group-003.md#canonical-630f9c052b3a52854049dda1da13fd44492ebaa15d266601ef878f649412091a)
- [sw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-34f2626bf35417bf94a0af238c2fe36603d4617d24af28174e43297681a6aa08)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-003.md#canonical-a7dfede66c6f88c6d8649f57e4468a404b29ff1dcbfc30fbb50f7fbfa3cd54a8)
- [waf_signatures](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0618f370940a8adff38a4bf27dd14fd5cd43884d9c0c76971bf3a53a4537d31a)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-5a14f5870b3035dae459cdb64efb6ed9926a55eb75a13deb3148f3381938d385"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c98e75642b6be517f28cef127e79e99cbb0cce2eb299a5bc9b2fd0bf60351fb"></a>

## admin_password — admin_password / 78abccabb70a / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- admin_password

<a id="canonical-cfec52c430194d1ffa6620690099aa3b1d74e704ebba7d2267d263f15ca50f12"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-af15b8f3e9702206f929ab8775559ca20a26c3c0d380017492e1d74486829534"></a>

## Direct properties — admin_password / 78abccabb70a / 3

- [blindfold_secret_info](data-sources--gcp_vpc_site--reference--group-001.md#canonical-57fb3a5564b8b3edb941dfde1e969edea5f3d881c43690b7658b868ff155cff7): complete subsection reference.

- [clear_secret_info](data-sources--gcp_vpc_site--reference--group-001.md#canonical-d09cae00d78081c0df2f1c3939dff2af14fd0e257b5e1266f48b722e49f86d92): complete subsection reference.

<a id="canonical-17d4af2da8c0a931857e34285e8f9b892e00416aaf9d3826a1fb4559d3b928d2"></a>

## Next pages — admin_password / 78abccabb70a / 4

- [admin_password.blindfold_secret_info](data-sources--gcp_vpc_site--reference--group-001.md#canonical-57fb3a5564b8b3edb941dfde1e969edea5f3d881c43690b7658b868ff155cff7)
- [admin_password.clear_secret_info](data-sources--gcp_vpc_site--reference--group-001.md#canonical-d09cae00d78081c0df2f1c3939dff2af14fd0e257b5e1266f48b722e49f86d92)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-57fb3a5564b8b3edb941dfde1e969edea5f3d881c43690b7658b868ff155cff7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2c1e5ea68312145dbc3709a88b5f9cf831d9aa34f8984e40b4136a7a223ff39"></a>

## admin_password.blindfold_secret_info — admin_password.blindfold_secret_info / 4e9116ee143e / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [admin_password](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5a14f5870b3035dae459cdb64efb6ed9926a55eb75a13deb3148f3381938d385)
- admin_password.blindfold_secret_info

<a id="canonical-1b8df325c33c1f572ec001fbec798770c46c8613993954135e78ed5d3578664b"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-41cceae193b1ccb829c2165252c9383a7f7d4ad409c2af6a04ab9efc9b82934d"></a>

## Direct properties — admin_password.blindfold_secret_info / 4e9116ee143e / 3

<a id="canonical-5fe5ec1db47ccf36d96ccb9724f2ff50fc9252b8d84d7c342283ab2bb1d6c00e"></a>

<a id="canonical-173d0340c4fe2419dcb13a226463a3d0dc29660190d5b3f9a7a33b5afe0e2bd2"></a>

## decryption_provider property — admin_password.blindfold_secret_info / 4e9116ee143e / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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
  }
}
```

<a id="canonical-1a05edc385bf0ee4d01c3e86aebdfee23f1e884727f7ddefe6b5d013b069d1ed"></a>

<a id="canonical-5c4e31c62fa77fa5c5063a1d56718b19c57a1c8c351d1c206eb8707a0b402121"></a>

## location property — admin_password.blindfold_secret_info / 4e9116ee143e / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-dd7de4b0da578506454d16c90a57c44e4f116e202564afcfe88f92d2ff5dc29e"></a>

<a id="canonical-9d883ccad6c69425cf4f2e646861477c7286630327bb132879db766878a03e6e"></a>

## store_provider property — admin_password.blindfold_secret_info / 4e9116ee143e / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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
  }
}
```

<a id="canonical-5b820a221dab2d696d631f6339c975c44b61c3556af41353e8278eff821d7d7f"></a>

## Next pages — admin_password.blindfold_secret_info / 4e9116ee143e / 7

- [admin_password](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5a14f5870b3035dae459cdb64efb6ed9926a55eb75a13deb3148f3381938d385)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-d09cae00d78081c0df2f1c3939dff2af14fd0e257b5e1266f48b722e49f86d92"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd411631f8293417a4a0a64e5c929ea3285df8371476b16806f9eedf06e80814"></a>

## admin_password.clear_secret_info — admin_password.clear_secret_info / a29fcc79e01f / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [admin_password](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5a14f5870b3035dae459cdb64efb6ed9926a55eb75a13deb3148f3381938d385)
- admin_password.clear_secret_info

<a id="canonical-ecdd36e50f734f26c3dea6d6d8a2553e7d4baef4da41a26e8a2cb17b97cfe6f5"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-a34c82b5b1ebfe4047e9f1c7908246fe2e9172ffff20c12591b1fc296a837fda"></a>

## Direct properties — admin_password.clear_secret_info / a29fcc79e01f / 3

<a id="canonical-b33eef66db46c5b6618cf3d2c78be03442c7264cb4c8c55b70f0088286cf5825"></a>

<a id="canonical-3af66b0cc3ae161d3e61db3a67620c0860cf5888c281729510c84a40389d9d38"></a>

## provider_ref property — admin_password.clear_secret_info / a29fcc79e01f / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-75aa40bf66007a9f7b80545a0ae08ae4c7cef206ebb5c181d0a15c409328db1d"></a>

<a id="canonical-8fef85cf981d5a900ea5fd99f16f67212cac74d60416d4b95955b0e1619b9127"></a>

## url property — admin_password.clear_secret_info / a29fcc79e01f / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-bb635a64af7e43197d6fa9c210cf3c82bb42bde18a226df32f4c157bd32fcd5a"></a>

## Next pages — admin_password.clear_secret_info / a29fcc79e01f / 6

- [admin_password](data-sources--gcp_vpc_site--reference--group-001.md#canonical-5a14f5870b3035dae459cdb64efb6ed9926a55eb75a13deb3148f3381938d385)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-1743bcd73450bc3dd30cbb6bf994b8c8c48c540ffbb2872e321b0ff1a581a1c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4c0d23abd520b8113100d33fae7229a47302477ce8ab23c4e2503c2476d2ce58"></a>

## block_all_services — block_all_services / d0ae3f966b34 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- block_all_services

<a id="canonical-b97b0749d84ed8431ebb47d4fe461f3875a3dadb7b8f37398bc246b852ce4397"></a>

Type: `["object", {}]`. Computed.

\[OneOf: block\_all\_services, blocked\_services, default\_blocked\_services; Default:
default\_blocked\_services\] Enable this option

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

- [block_all_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-b97b0749d84ed8431ebb47d4fe461f3875a3dadb7b8f37398bc246b852ce4397)
- [blocked_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-b05c226b6e8dc0e4c9a6e98613eafc9ea661945336f3154dae6264a33214ecca)
- [default_blocked_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-b197e4112f2b9172d1f211f632e6913450da6eee35c2c1008adb3e70c6ab6b7a)

Select alternatives according to the provider validators above.

<a id="canonical-79f795f84e0072193e34b466e8bf6a5df6305032437e4894150d25450c5b6263"></a>

## Direct properties — block_all_services / d0ae3f966b34 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dd4f035a5914c7f525f57d8a1899090333cfea99e6a6b972422f1218998498f6"></a>

## Next pages — block_all_services / d0ae3f966b34 / 4

- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-167e8ab000edb075a7e395fcf0e7b911b05439d69ce2b78ec25a808ae1153bf1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4384569dfc5dbce2b2ee4c058290f8dc61abf6673e28f478191e5ec992a46ff9"></a>

## blocked_services — blocked_services / c592b297c9db / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- blocked_services

<a id="canonical-b05c226b6e8dc0e4c9a6e98613eafc9ea661945336f3154dae6264a33214ecca"></a>

Type: `"single"`. Computed.

Disable node local services on this site.

Upstream description:

Disable node local services on this site. Note: The chosen services will GET disabled on all nodes
in the site.

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

<a id="canonical-b7e64542b5f625d20f59edf0179d5552ed52f746b06b7813905e38c8f850bd99"></a>

## Direct properties — blocked_services / c592b297c9db / 3

- [blocked_service](data-sources--gcp_vpc_site--reference--group-001.md#canonical-a2286abe290b2e0a63df5e9e79f29ca24bc24d90082a007b7c2c51c09d75f119): complete subsection reference.

<a id="canonical-67152eef7eaaef56508c40b7132e05130f2719004ec6f499e45dc363803884b7"></a>

## Next pages — blocked_services / c592b297c9db / 4

- [blocked_services.blocked_service](data-sources--gcp_vpc_site--reference--group-001.md#canonical-a2286abe290b2e0a63df5e9e79f29ca24bc24d90082a007b7c2c51c09d75f119)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-a2286abe290b2e0a63df5e9e79f29ca24bc24d90082a007b7c2c51c09d75f119"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e77e9ca9b8a9f8e09d91f88b865c7c6dd4a1ed659d8c21a93a1e1d1a8e7e0bc"></a>

## blocked_services.blocked_service — blocked_services.blocked_service / 87ff71b9db42 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [blocked_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-167e8ab000edb075a7e395fcf0e7b911b05439d69ce2b78ec25a808ae1153bf1)
- blocked_services.blocked_service

<a id="canonical-a8f66547b5f643c159d718cd607cdd8a068a9aba2d054558f81cd572c3f1f4d6"></a>

Type: `"list"`. Computed.

Disable Node Local Services. Blocking or denial configuration

Upstream description:

Blocking or denial configuration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3f4302f4de06b541526422f3b680d64569d6037420a021195fc0a49f9aec3e67"></a>

## Direct properties — blocked_services.blocked_service / 87ff71b9db42 / 3

- [dns](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0408fe91cfda85902adef4bd2430c4226c60f9e9a90dabdc3894ee6e34da529a): complete subsection reference.

<a id="canonical-e97986549006948ea0d633d153706bd0c2bd131487f4bc86f6a1de3494ca1b9a"></a>

<a id="canonical-6482ba58ca597a96788fcaac41cf6f3d64c194c942e94aee77a85338fc9f8557"></a>

## network_type property — blocked_services.blocked_service / 87ff71b9db42 / 4

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

- [ssh](data-sources--gcp_vpc_site--reference--group-001.md#canonical-7c656ac476a78863bf8d9e23017bdf49b118630a7a7a34869ce8df354ca6db11): complete subsection reference.

- [web_user_interface](data-sources--gcp_vpc_site--reference--group-001.md#canonical-475c0c61d85d4847077e24f84d4a31258c04ec2007ac3ea590c62366aa8f27ec): complete subsection reference.

<a id="canonical-16b96ec150d039b4dd6409d915e98eee73e746122b1406233eb09f9b8673912a"></a>

## Next pages — blocked_services.blocked_service / 87ff71b9db42 / 5

- [blocked_services.blocked_service.dns](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0408fe91cfda85902adef4bd2430c4226c60f9e9a90dabdc3894ee6e34da529a)
- [blocked_services.blocked_service.ssh](data-sources--gcp_vpc_site--reference--group-001.md#canonical-7c656ac476a78863bf8d9e23017bdf49b118630a7a7a34869ce8df354ca6db11)
- [blocked_services.blocked_service.web_user_interface](data-sources--gcp_vpc_site--reference--group-001.md#canonical-475c0c61d85d4847077e24f84d4a31258c04ec2007ac3ea590c62366aa8f27ec)
- [blocked_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-167e8ab000edb075a7e395fcf0e7b911b05439d69ce2b78ec25a808ae1153bf1)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-0408fe91cfda85902adef4bd2430c4226c60f9e9a90dabdc3894ee6e34da529a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3595d8dd0f6a17e02ae49b95dc7a2f3fe9ad4d3b5ddb09a19395020f4bcf3ea6"></a>

## blocked_services.blocked_service.dns — blocked_services.blocked_service.dns / bdbf6e347828 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [blocked_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-167e8ab000edb075a7e395fcf0e7b911b05439d69ce2b78ec25a808ae1153bf1)
- [blocked_services.blocked_service](data-sources--gcp_vpc_site--reference--group-001.md#canonical-a2286abe290b2e0a63df5e9e79f29ca24bc24d90082a007b7c2c51c09d75f119)
- blocked_services.blocked_service.dns

<a id="canonical-92fd2350efab52c10b664154e5b6f6eb5dc3e8810da3b5518b386e7a09c61338"></a>

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

<a id="canonical-296ca4de476c892574baf44ea5a2ed89b7aa20e7d9a7cac6e6f21d328823749c"></a>

## Direct properties — blocked_services.blocked_service.dns / bdbf6e347828 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fd442ab4910982bd7ded54b21f378471e61f40c5a36373ebb2f931b9b04888cd"></a>

## Next pages — blocked_services.blocked_service.dns / bdbf6e347828 / 4

- [blocked_services.blocked_service](data-sources--gcp_vpc_site--reference--group-001.md#canonical-a2286abe290b2e0a63df5e9e79f29ca24bc24d90082a007b7c2c51c09d75f119)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-7c656ac476a78863bf8d9e23017bdf49b118630a7a7a34869ce8df354ca6db11"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5291809cafef17ec620918b6832ffbbe9e8a79a958d5576c0389ece4b643ab09"></a>

## blocked_services.blocked_service.ssh — blocked_services.blocked_service.ssh / 7094e4fff99c / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [blocked_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-167e8ab000edb075a7e395fcf0e7b911b05439d69ce2b78ec25a808ae1153bf1)
- [blocked_services.blocked_service](data-sources--gcp_vpc_site--reference--group-001.md#canonical-a2286abe290b2e0a63df5e9e79f29ca24bc24d90082a007b7c2c51c09d75f119)
- blocked_services.blocked_service.ssh

<a id="canonical-c5ac66f15a51e35566af2c5fa7edaf244803ec08f8a80d6c2983ce2c794e70bf"></a>

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

<a id="canonical-5cc05bf99ccfd1a98a02ca4185435529b2f71a260f497bc39cc1f01eeb96fefc"></a>

## Direct properties — blocked_services.blocked_service.ssh / 7094e4fff99c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a7c6e6eabce34b39784dafdb1369aa37bbac4c428196a9e179375f5ea108b905"></a>

## Next pages — blocked_services.blocked_service.ssh / 7094e4fff99c / 4

- [blocked_services.blocked_service](data-sources--gcp_vpc_site--reference--group-001.md#canonical-a2286abe290b2e0a63df5e9e79f29ca24bc24d90082a007b7c2c51c09d75f119)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-475c0c61d85d4847077e24f84d4a31258c04ec2007ac3ea590c62366aa8f27ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-88869003ee3f004c5be669bb1108f95bfef9d82077312a0661b66783408facfe"></a>

## blocked_services.blocked_service.web_user_interface — blocked_services.blocked_service.web_user_interface / 9ec35a247323 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [blocked_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-167e8ab000edb075a7e395fcf0e7b911b05439d69ce2b78ec25a808ae1153bf1)
- [blocked_services.blocked_service](data-sources--gcp_vpc_site--reference--group-001.md#canonical-a2286abe290b2e0a63df5e9e79f29ca24bc24d90082a007b7c2c51c09d75f119)
- blocked_services.blocked_service.web_user_interface

<a id="canonical-213f54b831d2132d0a23e680d372a10623efe2f49ac2897f0d6dafcae0e041ce"></a>

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

<a id="canonical-df68f5807fa87f6b360e2dbeaa3f3b5a07e86c18c0f013af5c6542d2c6dd5273"></a>

## Direct properties — blocked_services.blocked_service.web_user_interface / 9ec35a247323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c5549836b08c2f1ac43630607d30ca91fd0865fc1c32b31b4a6497273bf6558e"></a>

## Next pages — blocked_services.blocked_service.web_user_interface / 9ec35a247323 / 4

- [blocked_services.blocked_service](data-sources--gcp_vpc_site--reference--group-001.md#canonical-a2286abe290b2e0a63df5e9e79f29ca24bc24d90082a007b7c2c51c09d75f119)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-95e08173bb2b66246ba892a158913465bb1a3e1f3e73312b25bb569b581e7f01"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3ecade2d445c73abe6115ac9d1206e8464fb362f0c2611d9cf1790deaea386ac"></a>

## cloud_credentials — cloud_credentials / 5ec73c2e2ebc / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- cloud_credentials

<a id="canonical-4d973002981b2f9fc5c86a7a880f1dd02a48f2c6327ceef9603d02bf2897d638"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-6b2074ec91bcf004c0f4a2b7f0d7b387b273a6c2be33baf7e706f01a264efab3"></a>

## Direct properties — cloud_credentials / 5ec73c2e2ebc / 3

<a id="canonical-447b17a081109fed6397aec6979270bae14792ba61a5afca767a62cebfc2c579"></a>

<a id="canonical-3e3014f01329b2fe11d68a0d3265267edbaa6fac83957cae9c36ade477bb3955"></a>

## name property — cloud_credentials / 5ec73c2e2ebc / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-47dd0304547e0bb825b699a3e65220b38434ae5083db4813bf7d38b0f1f6698a"></a>

<a id="canonical-2b5ff7177e0499b87f364160cb3c9a1192178bfd5cec2ef3776d3714a9fa589c"></a>

## namespace property — cloud_credentials / 5ec73c2e2ebc / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-e4c89512032e31fddca57183d17a761bdfd0c3276178355ffcc1c55bad92d9fa"></a>

<a id="canonical-ea71b70f7064c39c272887cfb0533e369721ebc40dc9dcea45a70b85cd2a5d83"></a>

## tenant property — cloud_credentials / 5ec73c2e2ebc / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-b3963e5160f739c7c0543ab4f43717a08e6855b67f5dad30db36562d49e18f3c"></a>

## Next pages — cloud_credentials / 5ec73c2e2ebc / 7

- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-3a09fef7692d7d176ee26a7f5e0dfa507c1426702ad0dee764975478e642e824"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a82d96e1665c1d12b4183755bc8f2785182d90512546a97a196940a36ea86452"></a>

## coordinates — coordinates / 0be3ce9db06f / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- coordinates

<a id="canonical-28b4e170b6113670af0a8a8c36e8256374f9d4a7ae5844dce632545554d02773"></a>

Type: `"single"`. Computed.

Coordinates of the site which provides the site physical location.

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

<a id="canonical-c88cf863c370b63142c1fcd67d9a0830d527f2ada54c14b7644cfc89206b1ed3"></a>

## Direct properties — coordinates / 0be3ce9db06f / 3

<a id="canonical-38b3f37b1f6c0bf864b01f3e287c581fb466d4c6e215ad00d6460fdf2b93e8c7"></a>

<a id="canonical-b8a361aacbfb588454dd39000a992216ab0fd675707637bad796de82cf271376"></a>

## latitude property — coordinates / 0be3ce9db06f / 4

Type: `"number"`. Computed.

Latitude. Latitude of the site location.

Upstream description:

Latitude of the site location.

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
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0"
  }
}
```

<a id="canonical-54c0a70d00576ad1c638b64c3f320937bd1b440085f045862191bd09ffae1e2f"></a>

<a id="canonical-61f456b020e9bd2ce75eac8373e1f716bbd4339d701e6e2312b3c30899a1e360"></a>

## longitude property — coordinates / 0be3ce9db06f / 5

Type: `"number"`. Computed.

Longitude. Longitude of site location.

Upstream description:

Longitude of site location.

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
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0"
  }
}
```

<a id="canonical-f71b55f7097b9d3232f0d55b8f394eb670eb302ef6f6e36caf9548515b4dc51e"></a>

## Next pages — coordinates / 0be3ce9db06f / 6

- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-664e363d354ecd5b80118961ededf199212a347905ced87272319db78cdae8c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c845b3c7d38f1b630b26f88fd7c78c149dd440c39d3fe64583c5ece67fc43a55"></a>

## custom_dns — custom_dns / a280b889b00c / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- custom_dns

<a id="canonical-0fa6367a2f6fb11a43df2a1de90fb4ba172f25bc5bdd63d6d5689cec45f80bac"></a>

Type: `"single"`. Computed.

Custom DNS is the configured for specify CE site.

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

<a id="canonical-74c1e0d16caf6c3a6fea3e88ff6e93fe3636fd702dbc4d546ff33a24303d8295"></a>

## Direct properties — custom_dns / a280b889b00c / 3

<a id="canonical-a9e5ebbd1e605c9329e2c34396e00d4d090496fe61c79cd06c27560f73200ae1"></a>

<a id="canonical-7bde376a7d1d85a46fb724328f2397e368db1cb7bf65415cd337789089c0efc6"></a>

## inside_nameserver property — custom_dns / a280b889b00c / 4

Type: `"string"`. Computed.

Optional DNS server IP to be used for name resolution in inside network.

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

<a id="canonical-86a9c8f0a2fb45cf15842fcefdc7d0a0e62cce606795304e276bb30b92f3c6e3"></a>

<a id="canonical-6d6d2e6e9b2b4ac0c7af3f1e845bbf1bcb8501bf91509231f7edb7ea999de00c"></a>

## outside_nameserver property — custom_dns / a280b889b00c / 5

Type: `"string"`. Computed.

Optional DNS server IP to be used for name resolution in outside network.

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

<a id="canonical-de16befc644e3712314a75f0a7b5a4096bc187b3e111f91cdf8ea6f3563ecfbb"></a>

## Next pages — custom_dns / a280b889b00c / 6

- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-39ab1e4957001d2198ac5bcd38c773bd46dd67c9cc4a1d35fd047a526a137431"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a64c13deb6dd33072be0f76458f32638bd2248d21e84cca7fd4d1447da580ad"></a>

## default_blocked_services — default_blocked_services / 5231f90214bc / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- default_blocked_services

<a id="canonical-b197e4112f2b9172d1f211f632e6913450da6eee35c2c1008adb3e70c6ab6b7a"></a>

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

<a id="canonical-a7d2e19beb246164caf61e2057f84e8cb4dff80d9cf1a6427c0e8f4c55cd718f"></a>

## Direct properties — default_blocked_services / 5231f90214bc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d5aa7c0241ea15c3abecb29f194515ffce72cc78f5f90ee4cf2678d72a035a99"></a>

## Next pages — default_blocked_services / 5231f90214bc / 4

- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-d0da01367080452c3a1b1fd1766dfd2520c59e731ace4092876fedfada95b872"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c6fd85b3d64c72919690cae7ad7e65f0bf32c70d1490f33319b7576d01b3a739"></a>

## disable_encryption — disable_encryption / 21549d6d8449 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- disable_encryption

<a id="canonical-4f58efe14f9f6273695bd62fbc1e53bff00356d47d9d071ed6166e7a66528810"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_encryption, enable\_encryption; Default: disable\_encryption\] Configuration
parameter for disable encryption.

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

- [disable_encryption](data-sources--gcp_vpc_site--reference--group-001.md#canonical-4f58efe14f9f6273695bd62fbc1e53bff00356d47d9d071ed6166e7a66528810)
- [enable_encryption](data-sources--gcp_vpc_site--reference--group-001.md#canonical-419df70af1c1534230d6063fc16381b8c7979a951e7f2b7461019bd1627483ef)

Select alternatives according to the provider validators above.

<a id="canonical-42c0e4c426b13a55007be136e5852239b010247fbb305c4e5922810b0b6acc48"></a>

## Direct properties — disable_encryption / 21549d6d8449 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-146e054d59631918625dd93e4ef54de59178a8261ff8dd75a802badb1c90088c"></a>

## Next pages — disable_encryption / 21549d6d8449 / 4

- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-acd4c983a00cc6d63f7585e07639826ad2c3037b5203a366d2ac239abb58430f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c8e5eb6ee526ba8c245e865101e1f9649cb61a6f4e50e4971936c59c81d32cf5"></a>

## enable_encryption — enable_encryption / 014bbeb08af5 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- enable_encryption

<a id="canonical-419df70af1c1534230d6063fc16381b8c7979a951e7f2b7461019bd1627483ef"></a>

Type: `"single"`. Computed.

Configuration parameter for enable encryption.

Upstream description:

Information related to disk encryption.

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

<a id="canonical-b5d8c541d0e11caf27bc575d04e269ac68c2144dd3edd53dff17ffd3660c6a59"></a>

## Direct properties — enable_encryption / 014bbeb08af5 / 3

<a id="canonical-99043520f92d62fbbafe369c5a02b07faddfd459b586cf7111ddd2933a2478f7"></a>

<a id="canonical-ba9316f05d9b0088e22e28d2a4faabbd276e5fc06eb3dd7f18863d76d16a8cba"></a>

## kms_key_resource_id property — enable_encryption / 014bbeb08af5 / 4

Type: `"string"`. Computed.

GCP KMS Key to be used to encrypt the disk attached to the VM.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-44f69e75cffa229e69201b9c81ae9e9475cf003e0dcbe91b1d5af5945b8874ff"></a>

<a id="canonical-8c76323b9b60678e46c57e86ac7c7b38989581295dc14000a65fdd5185fa8a31"></a>

## kms_key_ring_id property — enable_encryption / 014bbeb08af5 / 5

Type: `"string"`. Computed.

Key ring in which the CMK to be used to encrypt is present.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-0155375dcec18e4aaf9a460ce27cc6280a07370209dd45b29b4cdef0eca9d92d"></a>

## Next pages — enable_encryption / 014bbeb08af5 / 6

- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)

<a id="canonical-5bdfc213a88a3b5597f3d2bf1f9c923933afd2c4787d3747babb76e0bca0414d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a2c1c36d7317eb688503603a8764a364d2bc4e8dfce624926394a0e7dae991a1"></a>

## ingress_egress_gw — ingress_egress_gw / d0da8695e028 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-5ea9fe516e4a08529559279a042e97dbbf9f03724114ad43d94e8dc493e60ce2)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0bb43411449195b0334decb33d03b877c3d64f4ac0db69b30c8b848a6e1d74e1)
- ingress_egress_gw

<a id="canonical-371abe320a34379892077e9edd66cc32a98bd0311ff10c7da30dc8a2d3898231"></a>

Type: `"single"`. Computed.

\[OneOf: ingress\_egress\_gw, ingress\_gw, voltstack\_cluster\] Configuration parameter for ingress
egress gw.

Upstream description:

Two interface GCP ingress/egress site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group_inside_vn\",\"dc_cluster_group_outside_vn\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-inside_static_route_choice": "[\"inside_static_routes\",\"no_inside_static_routes\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]"
}
```

OneOf alternatives in this subsection:

- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-371abe320a34379892077e9edd66cc32a98bd0311ff10c7da30dc8a2d3898231)
- [ingress_gw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-aabb0996e5fa997bf2ff3a50da2ecedf568fefa97b3dfc132843c1eefe4bc0dc)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-003.md#canonical-237e3ceae77303258b3ff82422e1aa07d05f60ce403476754ef00c324a4d28d1)

Select alternatives according to the provider validators above.

<a id="canonical-a31218ec567770096eb6c80b4836d1818a38766db9b0d301838d9fe3f0032351"></a>

## Direct properties — ingress_egress_gw / d0da8695e028 / 3

- [active_enhanced_firewall_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-efc2a86054f64a35677656e7fc4511da231db396e8ec6036cb32d0adccbd1b20): complete subsection reference.

- [active_forward_proxy_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-e6193631fbded0892f775713582f87e4f4f92a0afc9506e8f6bd69fdcc233fb3): complete subsection reference.

- [active_network_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-ead75b4735c4c4d1ccb3f2d4272d4cdb4437fa10dc14414b53d9e00328df5ab5): complete subsection reference.

- [dc_cluster_group_inside_vn](data-sources--gcp_vpc_site--reference--group-002.md#canonical-8d296cd3a4b0a3dd01ad6e193c79922aa508f5e0085aa50b8ebb49da5bef1c41): complete subsection reference.

- [dc_cluster_group_outside_vn](data-sources--gcp_vpc_site--reference--group-002.md#canonical-b389184d68f963f28b88c5d20c75047db5ac5a10d81c341fb052e0764dabcc8a): complete subsection reference.

- [forward_proxy_allow_all](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0e7a7fa48e7f4a61d9d39f303af18f67da980e12340944bd25f56abe0387fd6e): complete subsection reference.

<a id="canonical-275f2fc528afae81c7cc474ec21f6846f4433a733f5fc0fd481139d80dcc1dcb"></a>

<a id="canonical-fbc74150917f33f225b4cc15961b3bba94e6a0c8740fe729a96cbe51a8a35878"></a>

## gcp_certified_hw property — ingress_egress_gw / d0da8695e028 / 4

Type: `"string"`. Computed.

\[Enum: gcp-byol-multi-nic-voltmesh\] GCP Certified Hardware. Name for GCP certified hardware. The
only possible value is \`gcp-byol-multi-nic-voltmesh\`.

Upstream description:

Name for GCP certified hardware.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "gcp-byol-multi-nic-voltmesh"
  ],
  "maxLength": 64,
  "x-f5xc-constraints": {
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"gcp-byol-multi-nic-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"gcp-byol-multi-nic-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-1334c4be72ad9d685fdddad85b7d421c3f92d6429e21625f8b92d1898602242b"></a>

<a id="canonical-22fd5778b3cc12feb65a803f767e8a214c21badb85f6451fc924d42093fcc2b5"></a>

## gcp_zone_names property — ingress_egress_gw / d0da8695e028 / 5

Type: `["list", "string"]`. Computed.

X-required List of zones when instances will be created, needs to match with region selected.

Upstream description:

X-required List of zones when instances will be created, needs to match with region selected.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [global_network_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-5cd988d29a53be3052bd3c168a710c189428b1761bef49ee16b5b771d1dc8fa9): complete subsection reference.

- [inside_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-5313e1755a61b9423761a49166595bbb829f161d8d2d4738be2005c5918dd40b): complete subsection reference.

- [inside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-e0df476e2452247689da449856486f73cbe5fe0ff99df824d89bedf66faeff32): complete subsection reference.

- [inside_subnet](data-sources--gcp_vpc_site--reference--group-002.md#canonical-e30eac2c9637444ff3739a6aedd40d8f7c6d96386ab2f6a39ea912ec4be9fc95): complete subsection reference.

- [no_dc_cluster_group](data-sources--gcp_vpc_site--reference--group-002.md#canonical-223b4c6ee2e40200248995849832753ffee9c86c9dc87018fda2c5396ee2d86d): complete subsection reference.

- [no_forward_proxy](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0ec5a58c2c48917a7da61be2b47d1dc1f04eb405913c388dd7af42e80bf96886): complete subsection reference.

- [no_global_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-68d5da7018c8d7bd5cf21710075ae517e46d160ab6af6bf1673a62ff41882716): complete subsection reference.

- [no_inside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-fb46fba0591415559af77dfe530544dae32a7e7bf32ae45df8de3b9501866c55): complete subsection reference.

- [no_network_policy](data-sources--gcp_vpc_site--reference--group-002.md#canonical-b9b388cb8608d6e8258f2c383d10cbc61754ac095dbebb0a235c994b398ddf2b): complete subsection reference.

- [no_outside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-996ab0f8550bc922577f6cd9fd87e03a64433e0dc6f34e2495971689dcd2059f): complete subsection reference.

<a id="canonical-2fac4443ae821e9585781dad2d8094f27954b69957df0d55ea606dc94c5b1b68"></a>

<a id="canonical-9db0a06bdf566e662db3f8eb7f8de0d80c24dcf16612bc01ec2e550057c4b5ba"></a>

## node_number property — ingress_egress_gw / d0da8695e028 / 6

Type: `"number"`. Computed.

Number of main nodes to create, either 1 or 3.

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
    "ves.io.schema.rules.uint32.in": "[1,3]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.in": "[1,3]"
  }
}
```

- [outside_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-14e07dadda9dcb215aa04b56160b05f8413817df23514c6b36bd611e75ce8dfe): complete subsection reference.

- [outside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-aa9a31421a4023e843ca9d13ca73facf79a8c7062588f2573cd865c49a32c160): complete subsection reference.

- [outside_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-53753b916e13f7c3bd7e73fca5e5c65bc9fe70864890e28c9ebefd14a67900c4): complete subsection reference.

- [performance_enhancement_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-ac213ade1ab480150edaf039160afd20cf3644175777c503f38aad3679d02688): complete subsection reference.

- [sm_connection_public_ip](data-sources--gcp_vpc_site--reference--group-003.md#canonical-a77b1430d70c66150cc8dd053e3b84d88666df8becf3e78c28282070d1905aaf): complete subsection reference.

- [sm_connection_pvt_ip](data-sources--gcp_vpc_site--reference--group-003.md#canonical-6c2d5083214b37795bbea7501148b070fa5c5a5a70763d4879ba4cf47cdad721): complete subsection reference.
