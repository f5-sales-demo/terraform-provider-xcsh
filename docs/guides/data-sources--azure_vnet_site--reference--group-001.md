---
page_title: "xcsh_azure_vnet_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_azure_vnet_site reference."
---

# xcsh_azure_vnet_site reference

<a id="canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf129cadece97320d596eebcbf9336a943b586cead7ad7990741b322002db9c2"></a>

## Property reference — Property reference / 2181621047d8 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- Property reference

<a id="canonical-33f89c29cf78cd790fabd6ec18befd456a49569af0f6994a94040edd5d55fa66"></a>

## Direct properties — Property reference / 2181621047d8 / 3

<a id="canonical-dab7241d40d9c5f2d7775736fed85dbee8dd369f8d5904c0b61350de51a49908"></a>

<a id="canonical-be5689a8f6f0d3df2abc2a28a061eaf221b7dccefdfa42a9f81ee0cdead45e00"></a>

## address property — Property reference / 2181621047d8 / 4

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

- [admin_password](data-sources--azure_vnet_site--reference--group-003.md#canonical-1a89f0e7e6b9bcfe5177d50035410a8c6e2c9fa52534c2a498b911ab328e9f9e): complete subsection reference.

<a id="canonical-127257e5489363f188ed2392a950f786be6e2ddcce1aa0b77ab3231542015fcd"></a>

<a id="canonical-f34c5a99bdc9e4962d415eed2c3c2f6693daa1e2067359bdac9627985c18ea87"></a>

## alternate_region property — Property reference / 2181621047d8 / 5

Type: `"string"`. Computed.

\[OneOf: alternate\_region, azure\_region\] Exclusive with \[azure\_region\] Name of the Azure
region which does not support availability zones.

Upstream description:

Exclusive with \[azure\_region\] Name of the Azure region which does not support availability zones.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

OneOf alternatives in this subsection:

- [alternate_region](data-sources--azure_vnet_site--reference--group-001.md#canonical-127257e5489363f188ed2392a950f786be6e2ddcce1aa0b77ab3231542015fcd)
- [azure_region](data-sources--azure_vnet_site--reference--group-001.md#canonical-3133b9405001a497e9eca1ae6b88fb071f611616ec6c482bd18a6e1764200953)

Select alternatives according to the provider validators above.

<a id="canonical-7e418b1e8ff8301b737f3cef46faca8f86ebd37026c7ccf4d512776ddafd06c9"></a>

<a id="canonical-87dc5285ee7176a6b0a8994cad819c25ce5e0ffee048db9cbd5355699961b63c"></a>

## annotations property — Property reference / 2181621047d8 / 6

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

- [azure_cred](data-sources--azure_vnet_site--reference--group-003.md#canonical-b21bed6e8628c7216a487fd4a5e375f59520eaba89cfd1018754ce10830eebdb): complete subsection reference.

<a id="canonical-3133b9405001a497e9eca1ae6b88fb071f611616ec6c482bd18a6e1764200953"></a>

<a id="canonical-eec5759fa06cb8502af2daa16f5546f9e97d7ad32c5b221a22cc36f781bef49b"></a>

## azure_region property — Property reference / 2181621047d8 / 7

Type: `"string"`. Computed.

Exclusive with \[alternate\_region\] Name of the Azure region which supports availability zones.

Upstream description:

Exclusive with \[alternate\_region\] Name of the Azure region which supports availability zones.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [block_all_services](data-sources--azure_vnet_site--reference--group-003.md#canonical-09b38e222f95d396a200ee40392cfbfd45374863bfdfa2d7b3bd8970497cb810): complete subsection reference.

- [blocked_services](data-sources--azure_vnet_site--reference--group-003.md#canonical-67325104c94d07874d8339b4c13ff70d60796dada4353dc8e17f7a394b4fef06): complete subsection reference.

- [coordinates](data-sources--azure_vnet_site--reference--group-003.md#canonical-a9317ff5fda9cec68c5760219234cab9edcfec3147cff76b51fab7e3a955e7d7): complete subsection reference.

- [custom_dns](data-sources--azure_vnet_site--reference--group-003.md#canonical-e2e793591316e3081ad55c1aefe5796e694efe21a578d575d8fd1ff0e8fbf03c): complete subsection reference.

- [default_blocked_services](data-sources--azure_vnet_site--reference--group-003.md#canonical-7d24bf85a14ffb122c47f20beb266f6d8de90059cc37f27b7da319c335333c96): complete subsection reference.

<a id="canonical-7ef107904742a5616aae869f9683434943f626a1ec12f7e40974547f8d6b0d38"></a>

<a id="canonical-171b84c61442d35fb9c37bc542165838bdf12e20ae2cb9cd88e99355eac84108"></a>

## description property — Property reference / 2181621047d8 / 8

Type: `"string"`. Computed.

Description of the AzureVNETSite.

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

- [disable_encryption](data-sources--azure_vnet_site--reference--group-003.md#canonical-548427ec503920ebc5bd16b9ed510eaafbd0bb6630fa283762152162e7c10917): complete subsection reference.

<a id="canonical-a5ffc41b58388cc04a09183820dbb2935ff6df6e1876702b15e42db7576d3c51"></a>

<a id="canonical-f7298bb29af8b749af830d1bef86531923872611873469c25277f339cfb9881d"></a>

## disk_size property — Property reference / 2181621047d8 / 9

Type: `"number"`. Computed.

Disk size to be used for this instance in GiB. 80 is 80 GiB. Server applies default when omitted.

Upstream description:

Disk size to be used for this instance in GiB. 80 is 80 GiB.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```

- [enable_encryption](data-sources--azure_vnet_site--reference--group-003.md#canonical-d99609a79b5132232d0bf4fc94b99b930299e9135bc8cad67fb8cf7c3df65297): complete subsection reference.

<a id="canonical-3b9c44831d5c9898a2fb11d89f2cb28df1f27be9c511d6212fb5e2f9fdb63fd2"></a>

<a id="canonical-295e85c93e644b740a997d8c0fd4f91f1a4c324d1c87e8cf848c1ff2f60d7acd"></a>

## id property — Property reference / 2181621047d8 / 10

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16): complete subsection reference.

- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5): complete subsection reference.

- [ingress_gw](data-sources--azure_vnet_site--reference--group-007.md#canonical-dd08dca1dd832961f9cb82aa1953d8dd272fdd8d597cb1b22d06d4ce23d8a901): complete subsection reference.

- [ingress_gw_ar](data-sources--azure_vnet_site--reference--group-007.md#canonical-9c0501d24be91d8ead0aff7cd66024d0a1b0773dd2a304d184993697820fbe3a): complete subsection reference.

- [kubernetes_upgrade_drain](data-sources--azure_vnet_site--reference--group-008.md#canonical-9d91b4d3269c468f37e5b1c26a807d4e08c4d28f8020f29879ec6f29f6fd447d): complete subsection reference.

<a id="canonical-09375c4ddd02b63249077dc69f4bd0adbd6036c3b08d8adb678a95a541ac3b7f"></a>

<a id="canonical-ab1b998759e9c24a3879b3ded21281dd35e572bb56ed61c4d5459022a7858746"></a>

## labels property — Property reference / 2181621047d8 / 11

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

- [log_receiver](data-sources--azure_vnet_site--reference--group-008.md#canonical-b99b8524e1a013a6a72d3a9f7661fed1c7864e36e3eea957a9dd57c77fbba74a): complete subsection reference.

- [logs_streaming_disabled](data-sources--azure_vnet_site--reference--group-008.md#canonical-50ddbe15a0b4f8cc869258011a291f265fb6150a389032ac6a5688aa0c3150cc): complete subsection reference.

<a id="canonical-c997d7228e8d665ff1164308c3e433bba07c33b728bde325bd71283e75cb26dc"></a>

<a id="canonical-4e1829eba87a932b81fc3e3962e8e05dc26bb7c246d151e37229e70274f04026"></a>

## machine_type property — Property reference / 2181621047d8 / 12

Type: `"string"`. Computed.

Select Instance size based on performance needed. The default setting for Accelerated Networking is
enabled, thus make sure you select a Virtual Machine that supports accelerated networking or disable
the setting under, Select Ingress Gateway or Ingress/Egress Gateway &gt; advanced OPTIONS.

Upstream description:

Select Instance size based on performance needed. The default setting for Accelerated Networking is
enabled, thus make sure you select a Virtual Machine that supports accelerated networking or disable
the setting under, Select Ingress Gateway or Ingress/Egress Gateway &gt; advanced OPTIONS.

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

<a id="canonical-22e5c50a466294bb79c6d712f766f184652bae948e8d6eb636ecc1e6bacb6eb5"></a>

<a id="canonical-e69781dc45e5ae084a21ca0af0308f9cd0814d1db5cbf814c329885c8a207ceb"></a>

## name property — Property reference / 2181621047d8 / 13

Type: `"string"`. Required.

Name of the AzureVNETSite.

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

<a id="canonical-cb2a20a9e6f997d9461277b3cb8601c6f3f48b9a48eeda1ccc54665dda4add8e"></a>

<a id="canonical-ef27ab9da276ed9a617f1aaaa63b8dff2c2d456484422a2f059895fb1d0ca9a6"></a>

## namespace property — Property reference / 2181621047d8 / 14

Type: `"string"`. Optional, Computed.

Namespace where the AzureVNETSite exists.

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

- [no_worker_nodes](data-sources--azure_vnet_site--reference--group-008.md#canonical-3dde6c9991cabec13fc09da9b97b5e4b5c4b8c57f696f860679ba7ba45452227): complete subsection reference.

<a id="canonical-7ece710b97abfc4379f02edc9ed52bc729ac4a083599346d19bb2b0bcaa76c78"></a>

<a id="canonical-f66b18c55aff3f6c7acfadf21a58f3718f12bea800f92c8f8176622e689f6e82"></a>

## nodes_per_az property — Property reference / 2181621047d8 / 15

Type: `"number"`. Computed.

Exclusive with \[no\_worker\_nodes total\_nodes\] Desired Worker Nodes Per AZ. Max limit is up to
21.

Upstream description:

Exclusive with \[no\_worker\_nodes total\_nodes\] Desired Worker Nodes Per AZ. Max limit is up to
21.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 21,
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
    "ves.io.schema.rules.uint32.lte": "21"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "21"
  }
}
```

- [offline_survivability_mode](data-sources--azure_vnet_site--reference--group-008.md#canonical-505af7c60d9977df112766d44d329998bfd8676f72709f37807c8a30d3beb8dd): complete subsection reference.

- [os](data-sources--azure_vnet_site--reference--group-008.md#canonical-ad91ef1a8c575ad8a704453f901a847cca12446e9b798a946f001a6b78ae9f43): complete subsection reference.

<a id="canonical-c50574a78c32a51c24535b998c3c4d6ffd2395e07c8c556cedf6171788722029"></a>

<a id="canonical-1eb6fc3f46740170a24c2ba3a1a9a4ed72880b968dbeb004e81fd12ec1bfc315"></a>

## resource_group property — Property reference / 2181621047d8 / 16

Type: `"string"`. Computed.

Azure resource group for resources that will be created.

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

<a id="canonical-d29cb8a584b3ff3e79f38247edabbe600240b26125c21eaa6305ee6994d36a2a"></a>

<a id="canonical-94b4197d8b6e6f31363c5ca4feb85fbf861a322be0a199f5f59c7816e103223a"></a>

## ssh_key property — Property reference / 2181621047d8 / 17

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

- [sw](data-sources--azure_vnet_site--reference--group-008.md#canonical-d758a0c910bf13db8ddd772fbbacefc572ecd33de2a09929b0f08904f1ea8116): complete subsection reference.

<a id="canonical-73c8b565edd1311f04cb2045dfb020d60e02d67739e86df738fd23f13f275ce2"></a>

<a id="canonical-04d748ea47c8e27436e427d1782f028891d7926b3bfaed6e16b34df4a2a81296"></a>

## tags property — Property reference / 2181621047d8 / 18

Type: `["map", "string"]`. Computed.

Azure Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in Azure console. Defaults to \`map\[\]\`. Server applies
default when omitted.

Upstream description:

Azure Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in Azure console.

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

<a id="canonical-6744264082a13fd5eb9e4a545a0887d9a23b73fefa85f23f734800095ba47b17"></a>

<a id="canonical-1ea48262918625a3ae26d3f0b3b3a5a781153591f52e95308ece4d282fd169ec"></a>

## total_nodes property — Property reference / 2181621047d8 / 19

Type: `"number"`. Computed.

Exclusive with \[no\_worker\_nodes nodes\_per\_az\] Total number of worker nodes to be deployed
across all AZ's used in the Site.

Upstream description:

Exclusive with \[no\_worker\_nodes nodes\_per\_az\] Total number of worker nodes to be deployed
across all AZ's used in the Site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 61,
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
    "ves.io.schema.rules.uint32.lte": "61"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "61"
  }
}
```

- [vnet](data-sources--azure_vnet_site--reference--group-008.md#canonical-65c989ffc21754741fafe57c771fc79f8a91db922463bf8a0b699799d04e4a5a): complete subsection reference.

- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0be67d2d7b5bbccba8043c7cc827083a7accdc2484c337bdd8b596fe11dde700): complete subsection reference.

- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197): complete subsection reference.

- [waf_signatures](data-sources--azure_vnet_site--reference--group-010.md#canonical-16645207c364e82299693024e80142d2a5ca8e82de875ea58d822f036f12778d): complete subsection reference.
