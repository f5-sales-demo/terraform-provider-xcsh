---
page_title: "xcsh_aws_vpc_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_aws_vpc_site reference."
---

# xcsh_aws_vpc_site reference

<a id="canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34502a8d20e518cc186be915d97491158d3a29356be31475ac438d585bbef0a2"></a>

## Property reference — Property reference / 4602ed3d6274 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- Property reference

<a id="canonical-88fff3c67d08a9689bd91fb5b830fa93c541383b83e482906c5ba6edcd1aba8b"></a>

## Direct properties — Property reference / 4602ed3d6274 / 3

<a id="canonical-65ec0992a7f6786adcc851ea247fcf1af370315bbe859f84fcd0e5cb06ba7863"></a>

<a id="canonical-f413eb521c04c4d2b098afda1aef399de42c6ef3814a504430bd502bfb1629bd"></a>

## address property — Property reference / 4602ed3d6274 / 4

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

- [admin_password](data-sources--aws_vpc_site--reference--group-001.md#canonical-cf61adb96e3255226dfecf4b2896b0db41c666d77e2b0b5f6245e0277b70c590): complete subsection reference.

<a id="canonical-febab45cb2d0bc602067af855b3fff967939a632079b17403740b5062e3b2bb7"></a>

<a id="canonical-f702b02585dc1fadaab23b5b13a80668eaa8b6cae11fd81511930bc2148ccbc7"></a>

## annotations property — Property reference / 4602ed3d6274 / 5

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

- [aws_cred](data-sources--aws_vpc_site--reference--group-001.md#canonical-49afbd2ef73399bad085c85ad245ccdfff8e955e7d434c32061687434d272168): complete subsection reference.

<a id="canonical-410e95a596ffe176c3aba2ea3ba78c0478061b7afc3ccf44acc1bf465e168f78"></a>

<a id="canonical-6a0fb79a2bfa6f191d5ddbcf1bbdfa5066932a1c4f708ac685d037d52051c6ea"></a>

## aws_region property — Property reference / 4602ed3d6274 / 6

Type: `"string"`. Computed.

AWS Region. Name for AWS Region.

Upstream description:

Name for AWS Region.

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

- [block_all_services](data-sources--aws_vpc_site--reference--group-001.md#canonical-b76dad76c53b4b19f8e8a154258132c8821bf9eb73ef0a9b94c8f4bd894f6ab7): complete subsection reference.

- [blocked_services](data-sources--aws_vpc_site--reference--group-001.md#canonical-d032a2669236aad9149c48eacbb9b9406c7baf6a30a2cba5c0b1773f6c359df0): complete subsection reference.

- [coordinates](data-sources--aws_vpc_site--reference--group-001.md#canonical-8e3fff956ccd777564613ff46454763434a51cac9b780741e9eb4b2807b63726): complete subsection reference.

- [custom_dns](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ecf2d6efcd15f9d35d6e047898741be3bc651f590662021e0080d0845258cf0): complete subsection reference.

- [custom_security_group](data-sources--aws_vpc_site--reference--group-001.md#canonical-b0731941313bd2cd269ea4438ac865999a040ab6e6f79209ee26f8f25364d2b3): complete subsection reference.

- [default_blocked_services](data-sources--aws_vpc_site--reference--group-002.md#canonical-a07d6b9e92df65b11b26bcede2ff5ea151145741fb6404881ba827c9c45df362): complete subsection reference.

<a id="canonical-e8f34009973d62ac906e4c852594e8c3cea5e0c3494bc2fbf1426aa8e389ecda"></a>

<a id="canonical-c83d6500e3f3f249140c7d1899c1f47f6665eb600e2e383c54eb19f3a1654657"></a>

## description property — Property reference / 4602ed3d6274 / 7

Type: `"string"`. Computed.

Description of the AWSVPCSite.

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

- [direct_connect_disabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-99c0899d6bf0f8198853a25a8ad9829c746898d5f8959e6a601b46e61cfcbcf1): complete subsection reference.

- [direct_connect_enabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-0573a5c999b2ef934474e06f6951d3bd7be0ac2c426b6cad83da0f5a329e4a6e): complete subsection reference.

- [disable_encryption](data-sources--aws_vpc_site--reference--group-002.md#canonical-7516c1587d0df53ad876cae0477b2c38385dec4dd103145bda5084516dbca326): complete subsection reference.

- [disable_internet_vip](data-sources--aws_vpc_site--reference--group-002.md#canonical-1fa273c1c1a47fed8481715b68dfeeea56376ac2fd09470b7846e01af544bae7): complete subsection reference.

<a id="canonical-4abc8ada4ead50df1a43f0bc76df1ef2f897332e2346d316c205f36bd0939cc7"></a>

<a id="canonical-681f6b1ca580a4f18ba8cd5ce3655a6567e44998c7912cf4ff6e0d728761bb86"></a>

## disk_size property — Property reference / 4602ed3d6274 / 8

Type: `"number"`. Computed.

Disk size to be used for this instance in GiB. 80 is 80 GiB.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2048,
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
    "ves.io.schema.rules.uint32.lte": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "2048"
  }
}
```

- [egress_gateway_default](data-sources--aws_vpc_site--reference--group-002.md#canonical-27eda6087dad60ebfab2bd9d183d63d4fd3d5ff5c8054c9ab718e31ae3b124e5): complete subsection reference.

- [egress_nat_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e4ad9435e0b0e6716c7f67d167fce5206201d39fc76b01abb1877385ec5606b): complete subsection reference.

- [egress_virtual_private_gateway](data-sources--aws_vpc_site--reference--group-002.md#canonical-408490482dd6c239b209790d7fbc93eab5b8299b6bc34c608b552cf8779a4ac7): complete subsection reference.

- [enable_encryption](data-sources--aws_vpc_site--reference--group-002.md#canonical-ebce9f17e54eded7f5f4b5f98bed9bd8bc031751d987e072803886c4bfa51dfb): complete subsection reference.

- [enable_internet_vip](data-sources--aws_vpc_site--reference--group-002.md#canonical-b2e9032440c417e11bfaa49c40aaf20079c79a6264edb2d63f11954983deeeca): complete subsection reference.

- [f5_orchestrated_routing](data-sources--aws_vpc_site--reference--group-002.md#canonical-cc60cce8b152be015a54c11c552724934297d0d6a25ba937a96666f5a1e6ba16): complete subsection reference.

- [f5xc_security_group](data-sources--aws_vpc_site--reference--group-002.md#canonical-e52d2230b230a613685d90c39f2ec99ce98ccd1da9751e604aed0656f9695db2): complete subsection reference.

<a id="canonical-1bd3dd1b6831e3700e0ca2f22eb4bdb7b80a28c2a4471f7cf8b5879da49a2835"></a>

<a id="canonical-0aac8d35f9bf68f01bdb8a2604f0c2266a64da75de1af7c1b211d7d4f9740321"></a>

## id property — Property reference / 4602ed3d6274 / 9

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d): complete subsection reference.

- [ingress_gw](data-sources--aws_vpc_site--reference--group-003.md#canonical-729ff7fed866124cf88a6351d29b823249a97f34b315308350e2b160cf54482b): complete subsection reference.

<a id="canonical-30b1865391e8cd43376af723820a8512dd3048141366b41895246b2f94af110b"></a>

<a id="canonical-31e611574595513e0f35df14da55d3ec6add80e7e2fc68c00b296564e04c567e"></a>

## instance_type property — Property reference / 4602ed3d6274 / 10

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

- [kubernetes_upgrade_drain](data-sources--aws_vpc_site--reference--group-004.md#canonical-59427859ce0ecf818307189237b8e8aa6b84229ce4372bc1a87c0e5503bccea9): complete subsection reference.

<a id="canonical-f5e4d4d3f2be7824afac2b85208e11f4c1d54ac3cce487cd89fe54bb24e48d8b"></a>

<a id="canonical-79776b8d1c922a565dc0ca46cb44ec187c927067217b0048532884847110cc5d"></a>

## labels property — Property reference / 4602ed3d6274 / 11

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

- [log_receiver](data-sources--aws_vpc_site--reference--group-004.md#canonical-1b88ba6da7c0147d0dad13276774e0bac834e226743ad35d599608626827317c): complete subsection reference.

- [logs_streaming_disabled](data-sources--aws_vpc_site--reference--group-004.md#canonical-2b8539e757fd3d9ffb4d3a36d093e0bd00f97a60629b501630154cfb65f30443): complete subsection reference.

- [manual_routing](data-sources--aws_vpc_site--reference--group-004.md#canonical-b195dd72eff5ecd49a6772d9d3512c0bb0648fb63e7b4b1b6248732bfe89f956): complete subsection reference.

<a id="canonical-aaa49b161a2dfdb50ed5475aac9f22ffe33eabb40c240f01d77b1552afdf78ae"></a>

<a id="canonical-76903ed307f67b608a589ac0f7e41ae36905193cdc9d5171df7236f7c9850b1e"></a>

## name property — Property reference / 4602ed3d6274 / 12

Type: `"string"`. Required.

Name of the AWSVPCSite.

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

<a id="canonical-b9dbe80fcc28c1f0875f8cc6377685f17641d72b3d8e99e8e08c0b2b69055a66"></a>

<a id="canonical-b2a17400d1e5962156661fb355f3b8f7aa2dd2836e809ba7ad15ae233e9bf534"></a>

## namespace property — Property reference / 4602ed3d6274 / 13

Type: `"string"`. Required.

Namespace where the AWSVPCSite exists.

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

- [no_worker_nodes](data-sources--aws_vpc_site--reference--group-004.md#canonical-f61dc2f597dcfadab5dcd644ce7cafde336974edd69d7e6dc56ac3e6cfeb7a80): complete subsection reference.

<a id="canonical-2498bcfc13ac5c6ab1e2a88a6814fb1e8d2d48a53f9718e39099cfaee0f674f0"></a>

<a id="canonical-6d175f69073dfb11c9514aa69d208c460c9b906a43f6df1207f1d0cfb1a3b181"></a>

## nodes_per_az property — Property reference / 4602ed3d6274 / 14

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

- [offline_survivability_mode](data-sources--aws_vpc_site--reference--group-004.md#canonical-1159cfbac09c60f8602b2b8e1fbb30d2be51e0e94e95e31a3745fbf56d1abd44): complete subsection reference.

- [os](data-sources--aws_vpc_site--reference--group-004.md#canonical-66baedec83b86131b792251eedb3fe856dec7e0caa69a3a1f165865f86da46d2): complete subsection reference.

- [private_connectivity](data-sources--aws_vpc_site--reference--group-004.md#canonical-ea5f3bd2603cd639f551392ffb29a7d414cf78ca1f13f34511f60a7f0b3e2fef): complete subsection reference.

<a id="canonical-4ca32422f176713adf96dbe54b95d7bdb8da7c9d452d666b89fcc351adc5b651"></a>

<a id="canonical-c06d9345b22315c694698821d56bbc9fb1f88e349eaa823f06320fdad0e20db7"></a>

## ssh_key property — Property reference / 4602ed3d6274 / 15

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

- [sw](data-sources--aws_vpc_site--reference--group-004.md#canonical-7ad1e1c86a364ee2ccd7fd4e86d325460994f69655dcf789db6b6d631c261de9): complete subsection reference.

<a id="canonical-c2dab72df5e4c9bb4bc60b9150c8bf00920b51240faaacdb6001108faa149ef8"></a>

<a id="canonical-efb3bf562ec222c8443e0b5bd8ad2e53965317baf82906a57097a81c2835f5dd"></a>

## tags property — Property reference / 4602ed3d6274 / 16

Type: `["map", "string"]`. Computed.

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

Upstream description:

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

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

<a id="canonical-94e8332871a3bfac446cd90bbb0ac94cba7e61e3feedc6425a9e52946d6c27e6"></a>

<a id="canonical-8c518c0d658bc40e12c74719c1b4127406d2a56ba9afc28bcfd9fc42edb55e42"></a>

## total_nodes property — Property reference / 4602ed3d6274 / 17

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

- [voltstack_cluster](data-sources--aws_vpc_site--reference--group-004.md#canonical-ce03987dc4fa70508f8e432b0fc8b0c65d9dacdc10695d71f9c4f1c61d03dbb9): complete subsection reference.

- [vpc](data-sources--aws_vpc_site--reference--group-005.md#canonical-c482fd7fdef218c974f1bc3d19e9ec98b067330464c5d118dc6b9acff264f21e): complete subsection reference.

- [waf_signatures](data-sources--aws_vpc_site--reference--group-005.md#canonical-1aaf17a2664f507686e37894dc171e25232558868c23cf23a2ff535374e932d3): complete subsection reference.

<a id="canonical-f59e2ed0eccf09230f08f24c42e1fadb520f2be22d5c72cb10fb6cb0b8bf0013"></a>

## All schema paths — Property reference / 4602ed3d6274 / 18

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address` | [address](data-sources--aws_vpc_site--reference--group-001.md#canonical-65ec0992a7f6786adcc851ea247fcf1af370315bbe859f84fcd0e5cb06ba7863) |
| `admin_password` | [admin_password](data-sources--aws_vpc_site--reference--group-001.md#canonical-fc291d0108076a236e1c4781fef57ae3ed27cfb4a03b4241230663d36b5f7070) |
| `admin_password.blindfold_secret_info` | [admin_password.blindfold_secret_info](data-sources--aws_vpc_site--reference--group-001.md#canonical-6f2722bc04c9be396235bfc73ed21a8fbee69f4df771009c264e41cd10a9803f) |
| `admin_password.blindfold_secret_info.decryption_provider` | [admin_password.blindfold_secret_info.decryption_provider](data-sources--aws_vpc_site--reference--group-001.md#canonical-8b1afbb930f07e4ae9bbf7517fca2873f1cbd20d39d8ef3f6ac1e27671946de7) |
| `admin_password.blindfold_secret_info.location` | [admin_password.blindfold_secret_info.location](data-sources--aws_vpc_site--reference--group-001.md#canonical-6c4f93fb40d6c4c1c5efc20e96f317a7d106a05c5800f3f607d0f6f4a05401fc) |
| `admin_password.blindfold_secret_info.store_provider` | [admin_password.blindfold_secret_info.store_provider](data-sources--aws_vpc_site--reference--group-001.md#canonical-36a8c7aff4c6e4151b4b4f16b6cb955b9a9bedf73bd830133d3420018e903c9f) |
| `admin_password.clear_secret_info` | [admin_password.clear_secret_info](data-sources--aws_vpc_site--reference--group-001.md#canonical-b31c8c1a3c838451dfd07df727c71be98e9d3d07a1b292e1391dea347ca3ff16) |
| `admin_password.clear_secret_info.provider_ref` | [admin_password.clear_secret_info.provider_ref](data-sources--aws_vpc_site--reference--group-001.md#canonical-351a09faadb7a301ed78f1a5740de6e28535622a3376f18e713b8708cc28c4b8) |
| `admin_password.clear_secret_info.url` | [admin_password.clear_secret_info.url](data-sources--aws_vpc_site--reference--group-001.md#canonical-36faf7acd05ac8b2e941c025cb47416043cbc1a3002be6ab7e33e1a71914117d) |
| `annotations` | [annotations](data-sources--aws_vpc_site--reference--group-001.md#canonical-febab45cb2d0bc602067af855b3fff967939a632079b17403740b5062e3b2bb7) |
| `aws_cred` | [aws_cred](data-sources--aws_vpc_site--reference--group-001.md#canonical-ed248ef5f9d51d868c83d7bb88f448c8f42a8a365b0d83fe0912db8c605d6f02) |
| `aws_cred.name` | [aws_cred.name](data-sources--aws_vpc_site--reference--group-001.md#canonical-f1ecc446525df5e6c845872fe1df093a8e9cafdddac455ba932eb9fd053a1372) |
| `aws_cred.namespace` | [aws_cred.namespace](data-sources--aws_vpc_site--reference--group-001.md#canonical-20e54c316d5c977054946199d295b626eab79ca2d4ec1db192560bfe4bc0d747) |
| `aws_cred.tenant` | [aws_cred.tenant](data-sources--aws_vpc_site--reference--group-001.md#canonical-a5911cfd95ccaf5f2c5fb0c2a46d4d11c2875537bfed6a06a8b5fb5d1445d758) |
| `aws_region` | [aws_region](data-sources--aws_vpc_site--reference--group-001.md#canonical-410e95a596ffe176c3aba2ea3ba78c0478061b7afc3ccf44acc1bf465e168f78) |
| `block_all_services` | [block_all_services](data-sources--aws_vpc_site--reference--group-001.md#canonical-dcb5608889f17fb4fd10b86025288ca1e52df507032af5ece883d86825b98250) |
| `blocked_services` | [blocked_services](data-sources--aws_vpc_site--reference--group-001.md#canonical-01d4814845c77eee790efcbacb2e6ceec072086ae3d672dbd6b704e4e01a6719) |
| `blocked_services.blocked_service` | [blocked_services.blocked_service](data-sources--aws_vpc_site--reference--group-001.md#canonical-d47764819710995a28c89be323fa87e2860f0ee7fc63de18d09c535ff7beb538) |
| `blocked_services.blocked_service.dns` | [blocked_services.blocked_service.dns](data-sources--aws_vpc_site--reference--group-001.md#canonical-ab259de40d2b3ff98a1bbdbba3b6f34a04faa2430342191906b1ff148a7bc13b) |
| `blocked_services.blocked_service.network_type` | [blocked_services.blocked_service.network_type](data-sources--aws_vpc_site--reference--group-001.md#canonical-29fb40b6587280108d94cac7c630ae5b929fe868abb24a32956371a8153e4e59) |
| `blocked_services.blocked_service.ssh` | [blocked_services.blocked_service.ssh](data-sources--aws_vpc_site--reference--group-001.md#canonical-e1280c0a66514e315b03a7ce9405d02931e96fc028ae065ac7f9f1256e351c24) |
| `blocked_services.blocked_service.web_user_interface` | [blocked_services.blocked_service.web_user_interface](data-sources--aws_vpc_site--reference--group-001.md#canonical-07c5b97e2ec618e1e71b0dea5fbd7ed6c8eb4e44cd19d972f8aa88805c2625e3) |
| `coordinates` | [coordinates](data-sources--aws_vpc_site--reference--group-001.md#canonical-38e253c1a0624adbe8b38bca06c8759cb3a395663ebf5f924e7e85d11a37d8eb) |
| `coordinates.latitude` | [coordinates.latitude](data-sources--aws_vpc_site--reference--group-001.md#canonical-7097d53ce0e870c19d8d94e0833b4126074279edfbfa242f2e03619cea445d1d) |
| `coordinates.longitude` | [coordinates.longitude](data-sources--aws_vpc_site--reference--group-001.md#canonical-f46901776730bc93c25ac7c2b202f9331ff37c62f8b4a6697b40afa941c41228) |
| `custom_dns` | [custom_dns](data-sources--aws_vpc_site--reference--group-001.md#canonical-a609cb474ffa654f25860da877b1f74674b8078fd224b09743477084bf59f6a1) |
| `custom_dns.inside_nameserver` | [custom_dns.inside_nameserver](data-sources--aws_vpc_site--reference--group-001.md#canonical-7bc6948ad76621b21d8c40dfce9cd921f637d196430606567f0e889b2d24fbb7) |
| `custom_dns.outside_nameserver` | [custom_dns.outside_nameserver](data-sources--aws_vpc_site--reference--group-001.md#canonical-40171d8b490253947e2254b85e89c9e4b49fe527ddff98b201f4ddf16ff3cce1) |
| `custom_security_group` | [custom_security_group](data-sources--aws_vpc_site--reference--group-002.md#canonical-59152abbfa2de683070f08747663621c3f154c22740e88e3a0bb7516f1688fff) |
| `custom_security_group.inside_security_group_id` | [custom_security_group.inside_security_group_id](data-sources--aws_vpc_site--reference--group-002.md#canonical-8f7b6827edc8b41b7f4312ff6d8a93f73c92393387a5a3a2600fb6dfe2f0c935) |
| `custom_security_group.outside_security_group_id` | [custom_security_group.outside_security_group_id](data-sources--aws_vpc_site--reference--group-002.md#canonical-37101993285dbd6d18accf978294df316fc9a4ef28c0a893a1b8bd4c41881b1f) |
| `default_blocked_services` | [default_blocked_services](data-sources--aws_vpc_site--reference--group-002.md#canonical-16f59aaf40db507d24180fa9a4990a0a0b3bee4380e31b5e5a4d2a3449015f84) |
| `description` | [description](data-sources--aws_vpc_site--reference--group-001.md#canonical-e8f34009973d62ac906e4c852594e8c3cea5e0c3494bc2fbf1426aa8e389ecda) |
| `direct_connect_disabled` | [direct_connect_disabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-2c617222c25217dc0b439d73cd7069474b4675af810537f936c2dee162515414) |
| `direct_connect_enabled` | [direct_connect_enabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-43a92841faa32b6ecee06bc534a421c302ba7c38832fd15815af04b5652ba3a1) |
| `direct_connect_enabled.auto_asn` | [direct_connect_enabled.auto_asn](data-sources--aws_vpc_site--reference--group-002.md#canonical-ff95c5d513e9ca3e0fb5766ea15926d414fe78ab80495923e5ccd79f3356a6fd) |
| `direct_connect_enabled.custom_asn` | [direct_connect_enabled.custom_asn](data-sources--aws_vpc_site--reference--group-002.md#canonical-87700e9942b8e0681a414f5166aee5acc77e8dfa8e3bd6617a75aefcd9f531e4) |
| `direct_connect_enabled.hosted_vifs` | [direct_connect_enabled.hosted_vifs](data-sources--aws_vpc_site--reference--group-002.md#canonical-3b1467446beb1aafda887fbcb3b9b7db8696e1c626eda9d6b13bc9545676b141) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect` | [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect](data-sources--aws_vpc_site--reference--group-002.md#canonical-b88c0dd12095ea5de077c90ba05052b633b08316d6af82dea8776bb13601acc1) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect.cloudlink_network_name` | [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect.cloudlink_network_name](data-sources--aws_vpc_site--reference--group-002.md#canonical-2c95dac838294e91c99cc067b2fc25896b7fd94e58de7af3ca34ac757bcd069d) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_internet` | [direct_connect_enabled.hosted_vifs.site_registration_over_internet](data-sources--aws_vpc_site--reference--group-002.md#canonical-a67fd1b69bd6ad2b80d539705d1314f70e8af744ba51152e949df6d87dc0cc1f) |
| `direct_connect_enabled.hosted_vifs.vif_list` | [direct_connect_enabled.hosted_vifs.vif_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-81791c20f393c64c06743821d0728623b0a6b48a5c96fa59bbbe1891d94a3e60) |
| `direct_connect_enabled.hosted_vifs.vif_list.other_region` | [direct_connect_enabled.hosted_vifs.vif_list.other_region](data-sources--aws_vpc_site--reference--group-002.md#canonical-d2e104ad110ded5dcf30906543c60724c8bf46fa0f2371a4aa354b52e090ce48) |
| `direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region` | [direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region](data-sources--aws_vpc_site--reference--group-002.md#canonical-f6f548d5eae93de74aa7230aec1453293535a0b11b9d22cad1b300fed13d6845) |
| `direct_connect_enabled.hosted_vifs.vif_list.vif_id` | [direct_connect_enabled.hosted_vifs.vif_list.vif_id](data-sources--aws_vpc_site--reference--group-002.md#canonical-1d3c3434295ac0408805dda5cf71279e7ecfa4f525331620fedef430b6f5c417) |
| `direct_connect_enabled.standard_vifs` | [direct_connect_enabled.standard_vifs](data-sources--aws_vpc_site--reference--group-002.md#canonical-116d3f8e44034f7159af45880438dff56f4804ce0730e96df3cc8d27a00d3505) |
| `disable_encryption` | [disable_encryption](data-sources--aws_vpc_site--reference--group-002.md#canonical-fffbfd32e36933134466c77289daef811a43e2eb3261a682831aa3e2b1651c73) |
| `disable_internet_vip` | [disable_internet_vip](data-sources--aws_vpc_site--reference--group-002.md#canonical-0933b0e5d9c81b8e2cb92c5d815ae1885ec94bf6d2396c60dde2742e087b6959) |
| `disk_size` | [disk_size](data-sources--aws_vpc_site--reference--group-001.md#canonical-4abc8ada4ead50df1a43f0bc76df1ef2f897332e2346d316c205f36bd0939cc7) |
| `egress_gateway_default` | [egress_gateway_default](data-sources--aws_vpc_site--reference--group-002.md#canonical-04ae9d37f36abc850a51c14aa5065a12dc1299d51fe56cb03192ea4be5d6aa03) |
| `egress_nat_gw` | [egress_nat_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-f9b981f1221b3bc14a00e436edcdf6531a5a0bbd875174e5cb016c74e7d2b225) |
| `egress_nat_gw.nat_gw_id` | [egress_nat_gw.nat_gw_id](data-sources--aws_vpc_site--reference--group-002.md#canonical-b8ade35052c84311900517fa685e3ebefef41ba0f97387416eeee83e0bd65f40) |
| `egress_virtual_private_gateway` | [egress_virtual_private_gateway](data-sources--aws_vpc_site--reference--group-002.md#canonical-7c275335231a1534b98345464271e54a30ed72381f8e15caeafbd207f57dc8ba) |
| `egress_virtual_private_gateway.vgw_id` | [egress_virtual_private_gateway.vgw_id](data-sources--aws_vpc_site--reference--group-002.md#canonical-51625879836dd753cde6495349ce898e3ff0c2d349fea5cf14cde078a4b9bafa) |
| `enable_encryption` | [enable_encryption](data-sources--aws_vpc_site--reference--group-002.md#canonical-f12821ee886d6be62e13aefc0b88f56f687993fdb3c8bed20d616976e32424de) |
| `enable_encryption.kms_key_id` | [enable_encryption.kms_key_id](data-sources--aws_vpc_site--reference--group-002.md#canonical-6eb0bc33796fdbf1639a3313764a8d10659acee07a0f48f34702fd6d7db4910d) |
| `enable_internet_vip` | [enable_internet_vip](data-sources--aws_vpc_site--reference--group-002.md#canonical-e9ec1ae00695b0d7b92a7a826fe89a1ad5b50947d7f25fd0a123ff27f6403d60) |
| `f5_orchestrated_routing` | [f5_orchestrated_routing](data-sources--aws_vpc_site--reference--group-002.md#canonical-727e6614fe7e42957a96f7e90ec05f73e047b974ef839990292469dc87ac9f85) |
| `f5xc_security_group` | [f5xc_security_group](data-sources--aws_vpc_site--reference--group-002.md#canonical-d60c801bbd7a4ffe2c7b01eafc17214e0bf8bc4c935d5ae36e2fb9d0ab2a3c8b) |
| `id` | [id](data-sources--aws_vpc_site--reference--group-001.md#canonical-1bd3dd1b6831e3700e0ca2f22eb4bdb7b80a28c2a4471f7cf8b5879da49a2835) |
| `ingress_egress_gw` | [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-8558d68e84e4446e22acc7dc87c078b052dee2c7a631eb6fee6fc62d3325f80f) |
| `ingress_egress_gw.active_enhanced_firewall_policies` | [ingress_egress_gw.active_enhanced_firewall_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-f4a1e550939e4a7e68aa4d4525621da7f58e6895becc49ca3a9c77b9d1659709) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-68b7057b9a6bb79c33ae9bb4c24ff914eba5c83ff2e2e840837ddb6e3988020e) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.name](data-sources--aws_vpc_site--reference--group-002.md#canonical-6b4f2cb8f51347c8934aedc1eab6d17ad733b646888657f0f093103cc1af7dc8) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](data-sources--aws_vpc_site--reference--group-002.md#canonical-a46575a67db96ad25431e57272a00f44da963d63a037a66112ad01bfaf716c6d) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](data-sources--aws_vpc_site--reference--group-002.md#canonical-87869e5a70428e895af6b0717689464bb7c86f1cbd5c383deba33d15dd0fa158) |
| `ingress_egress_gw.active_forward_proxy_policies` | [ingress_egress_gw.active_forward_proxy_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-7a327cb15ea1d59564ab2d30fcbcad843c0335543cdf05698e5bb6b023d34c0c) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-0a9b38ce7db69f72b267503f7caa9c87b0a1247961bab9a01ce1451aa9b6936e) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.name` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.name](data-sources--aws_vpc_site--reference--group-002.md#canonical-4304c534b11e2749b35b7fde65e0bd2a41ad158ef3f61ebac4d2249f6cc4fd6f) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.namespace` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.namespace](data-sources--aws_vpc_site--reference--group-002.md#canonical-5a0469d922fde628b649c41cf3111a11693c169b3bec50056d5cb223124fa36d) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.tenant` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.tenant](data-sources--aws_vpc_site--reference--group-002.md#canonical-ad2afd32399442e87ed8997d76a820b8b325b21bd290493d2032ebe6e2eeb4bb) |
| `ingress_egress_gw.active_network_policies` | [ingress_egress_gw.active_network_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-4dfbaaf99a1a0c711279300c25f85e0f1cfb5d4bf72ce76703836d4a2cfa5c37) |
| `ingress_egress_gw.active_network_policies.network_policies` | [ingress_egress_gw.active_network_policies.network_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-d0e224b3ed476cbae5fb9e78b12bf64e9b3c006200c282e289bd97e85941caed) |
| `ingress_egress_gw.active_network_policies.network_policies.name` | [ingress_egress_gw.active_network_policies.network_policies.name](data-sources--aws_vpc_site--reference--group-002.md#canonical-4412e995e1f6725edaba613a077d38d71d63ae8e1f39ec7379cc485133a05b2d) |
| `ingress_egress_gw.active_network_policies.network_policies.namespace` | [ingress_egress_gw.active_network_policies.network_policies.namespace](data-sources--aws_vpc_site--reference--group-002.md#canonical-49db6808a6039be1e7a5297762ff1f4f37e5059a0ab845b3a89d3506f1e7244d) |
| `ingress_egress_gw.active_network_policies.network_policies.tenant` | [ingress_egress_gw.active_network_policies.network_policies.tenant](data-sources--aws_vpc_site--reference--group-002.md#canonical-5cf83fa1435bb3be4bcac570c3d86bf52397d0ebb8b7830e06113ed6f3c16168) |
| `ingress_egress_gw.allowed_vip_port` | [ingress_egress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-e2dcf82a43c2e49d2780d22ed25ab13392b2e45615b73a8821c1554c2bd2113a) |
| `ingress_egress_gw.allowed_vip_port.custom_ports` | [ingress_egress_gw.allowed_vip_port.custom_ports](data-sources--aws_vpc_site--reference--group-002.md#canonical-722d2f4f028de81b574b8726052cc8abc397d4c03ebab8184748b1b4f6d80e54) |
| `ingress_egress_gw.allowed_vip_port.custom_ports.port_ranges` | [ingress_egress_gw.allowed_vip_port.custom_ports.port_ranges](data-sources--aws_vpc_site--reference--group-002.md#canonical-6b99cf6545bbb3694726c3f812992218832dfd688e9b939915629779b183cae0) |
| `ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port` | [ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-1df93eb95ccec2fe23bf8a9377649883a18109ca40fce8ca0ea8ee2dd0cb6b0e) |
| `ingress_egress_gw.allowed_vip_port.use_http_https_port` | [ingress_egress_gw.allowed_vip_port.use_http_https_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-7c121dbda2f667aff7aea2a3186ae1c29a48e35d13fa08524fb97126d108a7b0) |
| `ingress_egress_gw.allowed_vip_port.use_http_port` | [ingress_egress_gw.allowed_vip_port.use_http_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-fb12bd0262dd722998e25b5d97c1e116ff3d975c075c1599a10b5fcfdb978ab6) |
| `ingress_egress_gw.allowed_vip_port.use_https_port` | [ingress_egress_gw.allowed_vip_port.use_https_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-a81b3919d58ca353f1bc3db302bd493fb62bca4feb8ac4bcacd0010efa1aeffd) |
| `ingress_egress_gw.allowed_vip_port_sli` | [ingress_egress_gw.allowed_vip_port_sli](data-sources--aws_vpc_site--reference--group-002.md#canonical-0fd5a127f65155ac00096004f1e41df56f316f015d2127e37a064cc58ee4a751) |
| `ingress_egress_gw.allowed_vip_port_sli.custom_ports` | [ingress_egress_gw.allowed_vip_port_sli.custom_ports](data-sources--aws_vpc_site--reference--group-002.md#canonical-9d6e538eba4e1c2e9682861721c83d8d59aea3dba0f374eab40ed9cb5aa72164) |
| `ingress_egress_gw.allowed_vip_port_sli.custom_ports.port_ranges` | [ingress_egress_gw.allowed_vip_port_sli.custom_ports.port_ranges](data-sources--aws_vpc_site--reference--group-002.md#canonical-fa059e6d18e3cf35f4af3f554fcce780ccba3709ebef79fd4f89f3e771cdb2dd) |
| `ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port` | [ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-657cf327e956a3828fa46f52d023471f38e4bdd8ccf77a409c90126b50efe3c9) |
| `ingress_egress_gw.allowed_vip_port_sli.use_http_https_port` | [ingress_egress_gw.allowed_vip_port_sli.use_http_https_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-457f833b30b3e5ca566e873daa37fb3053ef3948a269ba1e6cc047be2d2993a6) |
| `ingress_egress_gw.allowed_vip_port_sli.use_http_port` | [ingress_egress_gw.allowed_vip_port_sli.use_http_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-e3002714a1b4a16ff09ee650ef9797a5e0e574e91a44fb61a4ed52da052eaf4b) |
| `ingress_egress_gw.allowed_vip_port_sli.use_https_port` | [ingress_egress_gw.allowed_vip_port_sli.use_https_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-0ae522f689ecbdfa3eaa2598e3e37c99701fa357a585d2446132bcf43a3f0034) |
| `ingress_egress_gw.aws_certified_hw` | [ingress_egress_gw.aws_certified_hw](data-sources--aws_vpc_site--reference--group-002.md#canonical-e2a0e257d95d06fdb4e2629f5a8d7dffac5741f4b00f7b656595406b912df3c9) |
| `ingress_egress_gw.az_nodes` | [ingress_egress_gw.az_nodes](data-sources--aws_vpc_site--reference--group-002.md#canonical-29e3b9fb11728347a51210058cae464686194a7e0d8d2c5549178c53ca52bb9c) |
| `ingress_egress_gw.az_nodes.aws_az_name` | [ingress_egress_gw.az_nodes.aws_az_name](data-sources--aws_vpc_site--reference--group-002.md#canonical-f2c111696289bdbf3919366325edccd101bfb9f2b4aba7d41fd540e64e27a70b) |
| `ingress_egress_gw.az_nodes.inside_subnet` | [ingress_egress_gw.az_nodes.inside_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-315137bda29437a98ee9aa6e7e81a6ee6629cfec48c7d9ce03d66f994cf80dbb) |
| `ingress_egress_gw.az_nodes.inside_subnet.existing_subnet_id` | [ingress_egress_gw.az_nodes.inside_subnet.existing_subnet_id](data-sources--aws_vpc_site--reference--group-002.md#canonical-e6f687636e7ea38071f275fca358391265eb2a8bf9ebba5b97e3e2b279b9fc99) |
| `ingress_egress_gw.az_nodes.inside_subnet.subnet_param` | [ingress_egress_gw.az_nodes.inside_subnet.subnet_param](data-sources--aws_vpc_site--reference--group-002.md#canonical-cef5a8343b4dafff8e645341a11c546cc3330e3309c9b41ab97365e65c356a55) |
| `ingress_egress_gw.az_nodes.inside_subnet.subnet_param.ipv4` | [ingress_egress_gw.az_nodes.inside_subnet.subnet_param.ipv4](data-sources--aws_vpc_site--reference--group-002.md#canonical-24db48ad6b0f43791d99ce284d071e7fc088562788774a781f08d49a14127536) |
| `ingress_egress_gw.az_nodes.outside_subnet` | [ingress_egress_gw.az_nodes.outside_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-b91a6b60febe15b641e8cab8db705ce10cc2f4b2402c20028974603e5ef4241e) |
| `ingress_egress_gw.az_nodes.outside_subnet.existing_subnet_id` | [ingress_egress_gw.az_nodes.outside_subnet.existing_subnet_id](data-sources--aws_vpc_site--reference--group-002.md#canonical-244f4a27a6c57fe9a10d7a0548fa102790517e5eef2fed4af04e30420d5743b7) |
| `ingress_egress_gw.az_nodes.outside_subnet.subnet_param` | [ingress_egress_gw.az_nodes.outside_subnet.subnet_param](data-sources--aws_vpc_site--reference--group-002.md#canonical-45b00ef1713fbc4f33be6b17dc45ff62e8959edf1408653d8b2ec2b95c363975) |
| `ingress_egress_gw.az_nodes.outside_subnet.subnet_param.ipv4` | [ingress_egress_gw.az_nodes.outside_subnet.subnet_param.ipv4](data-sources--aws_vpc_site--reference--group-002.md#canonical-6f325b172d23a7c4ac8474bdc4464a6cffb9c9837d78ad3855e6856a58885d07) |
| `ingress_egress_gw.az_nodes.reserved_inside_subnet` | [ingress_egress_gw.az_nodes.reserved_inside_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-6b4170c9d6f64469d3ec880948163c65dca73a4aaf31dc60d98989a8faeb6cfa) |
| `ingress_egress_gw.az_nodes.workload_subnet` | [ingress_egress_gw.az_nodes.workload_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-3864657aed666dd9ab4d27391d0d5d3255c6364f4c8ea2264625a2de016861fd) |
| `ingress_egress_gw.az_nodes.workload_subnet.existing_subnet_id` | [ingress_egress_gw.az_nodes.workload_subnet.existing_subnet_id](data-sources--aws_vpc_site--reference--group-002.md#canonical-e261a82eb30f2345d70187b4a99a010a3ad6f11981a5fb5b6106e7d58c1e9174) |
| `ingress_egress_gw.az_nodes.workload_subnet.subnet_param` | [ingress_egress_gw.az_nodes.workload_subnet.subnet_param](data-sources--aws_vpc_site--reference--group-002.md#canonical-bc883d794512cc03f2afee012b9966e7773109b0982c7ea23d5c22cb7951f44c) |
| `ingress_egress_gw.az_nodes.workload_subnet.subnet_param.ipv4` | [ingress_egress_gw.az_nodes.workload_subnet.subnet_param.ipv4](data-sources--aws_vpc_site--reference--group-002.md#canonical-79ecacdbb7b8f7c853cbcdcb7249ee86708b3b868d1c87c83cef9639fe5a0cc8) |
| `ingress_egress_gw.dc_cluster_group_inside_vn` | [ingress_egress_gw.dc_cluster_group_inside_vn](data-sources--aws_vpc_site--reference--group-002.md#canonical-72f5dcfd839e3b2032383615ed0623c2754f259abe9e242836ecc7961b45de4f) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.name` | [ingress_egress_gw.dc_cluster_group_inside_vn.name](data-sources--aws_vpc_site--reference--group-002.md#canonical-61816814407d2fa22920ff5cd934c2e2e7c17be50ba3c5788d49be2282efd091) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.namespace` | [ingress_egress_gw.dc_cluster_group_inside_vn.namespace](data-sources--aws_vpc_site--reference--group-002.md#canonical-91148a4776aa94d49c5fb4ee86aae40e54cc88516813848ecdec4f82b7330000) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.tenant` | [ingress_egress_gw.dc_cluster_group_inside_vn.tenant](data-sources--aws_vpc_site--reference--group-002.md#canonical-f95cc387036a8a44d45e915be0bebea1aaf9b557c054cffbd1e221811c5350fc) |
| `ingress_egress_gw.dc_cluster_group_outside_vn` | [ingress_egress_gw.dc_cluster_group_outside_vn](data-sources--aws_vpc_site--reference--group-002.md#canonical-e2c05b684d8a16ff6994951ad9fd9460824854e559144b95e178c8cdaa52cb7b) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.name` | [ingress_egress_gw.dc_cluster_group_outside_vn.name](data-sources--aws_vpc_site--reference--group-002.md#canonical-2716edb4b0f594c6bfcd75e54921f68e91af7b2bb3d2dfdd0a75002d849d17fa) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.namespace` | [ingress_egress_gw.dc_cluster_group_outside_vn.namespace](data-sources--aws_vpc_site--reference--group-002.md#canonical-5f180e8ba8e7021176f4ae358b14337be3a403acf3a55ede767e20689a8e561a) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.tenant` | [ingress_egress_gw.dc_cluster_group_outside_vn.tenant](data-sources--aws_vpc_site--reference--group-002.md#canonical-025eb969eb8dd591670569d1f6ad98e8f8d790c4d898177eb280ee94f530bb47) |
| `ingress_egress_gw.forward_proxy_allow_all` | [ingress_egress_gw.forward_proxy_allow_all](data-sources--aws_vpc_site--reference--group-002.md#canonical-4bf09b6012f00f3ba136d25aae4021f8b2db0ce973bb476260d15dee2dd25b54) |
| `ingress_egress_gw.global_network_list` | [ingress_egress_gw.global_network_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-5521a81083b62d98e7da46e24e813c488ec85e8c56c22ead1966d72608c3248a) |
| `ingress_egress_gw.global_network_list.global_network_connections` | [ingress_egress_gw.global_network_list.global_network_connections](data-sources--aws_vpc_site--reference--group-002.md#canonical-05862f8af1d646fd3e5ffe76f006f3531ffae38f22c128a04754c3096dd901e4) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](data-sources--aws_vpc_site--reference--group-002.md#canonical-4188507b8b4d44c76aad65e36eef3b4a412835924b1338aa61e5d9a97d144da7) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn](data-sources--aws_vpc_site--reference--group-002.md#canonical-29af3ba871a0a2b53d6a6084c6550214fe0812b9a3d979eacb9de7dd94b42273) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](data-sources--aws_vpc_site--reference--group-002.md#canonical-8f5ddb808ab7d5239ecac9a808d9aef37eb560a9f8f6a957ef8966502fb41ae2) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](data-sources--aws_vpc_site--reference--group-002.md#canonical-3a4ed092f9dd647fec552a41a8d339840f26626bf9e4a09d3f5fbc8816c78d81) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](data-sources--aws_vpc_site--reference--group-002.md#canonical-c969f95f27e91b4dc3218382eda1ec082c16cfc289de2910460c1947ae14b2a4) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](data-sources--aws_vpc_site--reference--group-002.md#canonical-5e74501abe51c29e1f93ed0d4ff47a676becb8fb388e5b5294da9d8b5f4da4a1) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn](data-sources--aws_vpc_site--reference--group-002.md#canonical-f7e17e9978f36d12ee61b1e91b3394693dd8d184303876d1e59d51de1c5be722) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](data-sources--aws_vpc_site--reference--group-002.md#canonical-e5bfaf3ae0bbce04df3b1ad232e79b6bbcb5bd274741e6e2b9154a2b15ffa2a5) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](data-sources--aws_vpc_site--reference--group-002.md#canonical-4ac3962137596f256ef1287c4c2a19d6cd588fe1e587bcfc168d4c37c24e146f) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](data-sources--aws_vpc_site--reference--group-002.md#canonical-10e1f1ed059fc725f0ccbdcd190b3fd478e7ce597e8510386b4cbc9219d8822e) |
| `ingress_egress_gw.inside_static_routes` | [ingress_egress_gw.inside_static_routes](data-sources--aws_vpc_site--reference--group-002.md#canonical-3386420ad603ea21d4462c53d89bbadb38be73cb65e6ff506fb4b1b47b50aafc) |
| `ingress_egress_gw.inside_static_routes.static_route_list` | [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-ebccf0a40a4275148d96809d79c748a69b1d6b27150583f6b12ce12e30ec4288) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-002.md#canonical-622c74b78190ca224666f357756c6e9b8e10b54fc9818b786013c8a0cc6865dc) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.attrs` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.attrs](data-sources--aws_vpc_site--reference--group-002.md#canonical-e2fb42ce1ae32f79ebba007e358e6abbab237a9e47818bd6c009a94daffcf25d) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels](data-sources--aws_vpc_site--reference--group-003.md#canonical-45f8502bee27cc148dd76466ad664957c10fae8e3b5874d11ba1838995b917c4) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-003.md#canonical-a6072767f1ea1ff2a9371785109cbdbe8fd5ce8acba040948c7f0de30647fac4) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--aws_vpc_site--reference--group-003.md#canonical-f85f3353fd2bf1329dd0e15f6fa5b089ca6cd4f79fe7010827a4e8d1512037ef) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](data-sources--aws_vpc_site--reference--group-003.md#canonical-85d276780c7b02db6f2d7ce60a0c64193e95781533c0cc39ff228b10c64f892e) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](data-sources--aws_vpc_site--reference--group-003.md#canonical-2b11c0c5ec7913865cf99bb356206cb86bfa9633abffaec74a6f6ba7f466e012) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](data-sources--aws_vpc_site--reference--group-003.md#canonical-ec90cb9d10ddf66ca3b4b1c82b0486510708f233f371372694a0a52d01294883) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](data-sources--aws_vpc_site--reference--group-003.md#canonical-3a2c7737396afcaefbdf5f1140cf22618d9ed8cb28cd4f463ab698fbfd250275) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](data-sources--aws_vpc_site--reference--group-003.md#canonical-331b825d6b9608d9d283aee07bcd12f32efdc6e081219dd0e53e595c9c0e9613) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-003.md#canonical-43de7dbfe32825a6368e8e7bbd9623c819f0f9f63244b501b2bc3d3745dd256b) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_vpc_site--reference--group-003.md#canonical-aecac3638167e9181057c3329c98bb58352ffa329a8d856a1c624ec765cc02ed) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--aws_vpc_site--reference--group-003.md#canonical-9af5ed20354997e49bd0843b19e93afaf11f9c5674204ce079e54ae2462c2ac9) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](data-sources--aws_vpc_site--reference--group-003.md#canonical-2fcfeac5b9697a9e152e08a88ac4c55f2b6225ee6fa632b7ba3bcea42d49f053) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--aws_vpc_site--reference--group-003.md#canonical-83f8f5b85f1dad0ecdf5115196a37510993ed2b98a6b6e8d9ea93b623eebec36) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](data-sources--aws_vpc_site--reference--group-003.md#canonical-15368698337b55c4b7e11c0b1344b5ba124ca0904e3e97f81b5f449be5da04eb) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--aws_vpc_site--reference--group-003.md#canonical-f6d6dc292dc536a9174393434f863dd2b66522e700f6a83676ce24e9a9475ab8) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](data-sources--aws_vpc_site--reference--group-003.md#canonical-14dafdc8d78fdef20995c1858739fcf60c14c6afc86aa772cdf5b36cf93022f2) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--aws_vpc_site--reference--group-003.md#canonical-b7c63d23a017dcd8ec326c322f730d6e8cd80dce7df33502a1fbe27587864db8) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](data-sources--aws_vpc_site--reference--group-003.md#canonical-9c8d1e9a702758d977ee1a46d02c2a86c88305e894b8b253a0893c4aa9637ba5) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.type` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.type](data-sources--aws_vpc_site--reference--group-003.md#canonical-e5cfc643ceed789fca34925d3a27b1fee46fc9864c2326387630b44b9ee3be5d) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_vpc_site--reference--group-003.md#canonical-9a824d932bf7460d9988d9cb1a8ec932ef2efce5eaaa04d5681eb380fdc685f1) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--aws_vpc_site--reference--group-003.md#canonical-365773838d15310b0720bc038493e74937b71f9b7405450aac2e9bd2b3b42c4c) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](data-sources--aws_vpc_site--reference--group-003.md#canonical-446913654939ba2af8cbc45d19c5c05a902d2bf5092aa0c0604109d5ecf5e684) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](data-sources--aws_vpc_site--reference--group-003.md#canonical-c1ae7b178eedfc78461aaf8adbf19dc061f247c3e77a43316dc69f24e4abf000) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--aws_vpc_site--reference--group-003.md#canonical-2e8cf085e81101e3faddde3bf365d67fc46e3a22c310e2b6b989c87852828bc3) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](data-sources--aws_vpc_site--reference--group-003.md#canonical-ca4203367f786afeb82d775baaaf31e07fcdc392a950993e1b3e809a06ea7c76) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](data-sources--aws_vpc_site--reference--group-003.md#canonical-d5b7d9f207efa71f0a8393a707e38c11c30a9eced0d5a9a548edd5e988288667) |
| `ingress_egress_gw.inside_static_routes.static_route_list.simple_static_route` | [ingress_egress_gw.inside_static_routes.static_route_list.simple_static_route](data-sources--aws_vpc_site--reference--group-002.md#canonical-fdca5986a767f77eb2a988c5e69ee5a69cf99062c3eb547a651889071596e41b) |
| `ingress_egress_gw.no_dc_cluster_group` | [ingress_egress_gw.no_dc_cluster_group](data-sources--aws_vpc_site--reference--group-003.md#canonical-a59363132017160c2a6633e55dbf6012a0f368350a3438bacb2a2a2510de12ad) |
| `ingress_egress_gw.no_forward_proxy` | [ingress_egress_gw.no_forward_proxy](data-sources--aws_vpc_site--reference--group-003.md#canonical-30b3bf93b9a7fede73a4b64778be17bd935724ddad8e3d32449b37916e06ff8e) |
| `ingress_egress_gw.no_global_network` | [ingress_egress_gw.no_global_network](data-sources--aws_vpc_site--reference--group-003.md#canonical-754f55387b53244197fba45b30d6ed2f48dfb32601d6c56464eaeecd0a1ddf8b) |
| `ingress_egress_gw.no_inside_static_routes` | [ingress_egress_gw.no_inside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-2ed59c877d217d5b475b41780e42b0221e53bc07a8472e562b44d1b9dd452675) |
| `ingress_egress_gw.no_network_policy` | [ingress_egress_gw.no_network_policy](data-sources--aws_vpc_site--reference--group-003.md#canonical-bf5478e247b15eefec5ad1619c536a20798a6ee633bcaa3e6801274f0abde0d5) |
| `ingress_egress_gw.no_outside_static_routes` | [ingress_egress_gw.no_outside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-140cbd33de2094965eeee39edaea4a346803f1dcdf64d8f9e0206b1d9cb5cf42) |
| `ingress_egress_gw.outside_static_routes` | [ingress_egress_gw.outside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-7f09be5b6b290e9a636c03f913df453b4b1a57e6e0e902957d29979a7bb974a3) |
| `ingress_egress_gw.outside_static_routes.static_route_list` | [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-ed37c67968eb0543ad0ac73290509e558221ea1c71cfa0334237074aa8469628) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-0d9d65521f70375ff782b5fe12e4a2ee9747151f4c0a86ec4dcc924e97a798a7) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.attrs` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.attrs](data-sources--aws_vpc_site--reference--group-003.md#canonical-d8a8c1ef033607d9606b8cf0381d3fe5bae4f8c413400267ad866f19721bff8c) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels](data-sources--aws_vpc_site--reference--group-003.md#canonical-8ca4b8ce95e2ef2449b5078b471d7c865da5ad159fa84698b41b7b8a1978f791) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-003.md#canonical-1d06e0f691171cf405fed9da26eb3c96d2f3036b110c214f7811b03d9b32e0dc) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--aws_vpc_site--reference--group-003.md#canonical-c125a88ab09bf436be29f5591135dc86a49f7d76c5bb8a30e79cde0af1b543b3) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](data-sources--aws_vpc_site--reference--group-003.md#canonical-e4cfb23eda8da627c8160f0eab63c08325d31ac341553b331b0725a330e4f9c1) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](data-sources--aws_vpc_site--reference--group-003.md#canonical-41108069854b5ac42fb628c935bb23fdf0158b4b7d56589c4a9c79c05c3b0336) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](data-sources--aws_vpc_site--reference--group-003.md#canonical-e8ecc1f93a13b7c4d4fd6fd64494a1525c008601146c9fc756fb5a0ba305e43c) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](data-sources--aws_vpc_site--reference--group-003.md#canonical-bccd782ae1b4f00758a156e0a6111107e24861927ff5fd84d5146de641d7975d) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](data-sources--aws_vpc_site--reference--group-003.md#canonical-8ee18ff21419bf4499084bf825a35bc6af2e5a7a7455467705e6e06402aa0879) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-003.md#canonical-d72217af01449d6862d72aa7ebd5ba2b09139bad23b591ac0c8ec088166af3ab) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_vpc_site--reference--group-003.md#canonical-377bd8dd19b81d7477bb192d851fa8031658b221123e4f694d361ec93348c1fe) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--aws_vpc_site--reference--group-003.md#canonical-bbdc86d1196da539bff26a648bf855625195af1b60af61d5671a317919b4d8ac) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](data-sources--aws_vpc_site--reference--group-003.md#canonical-aaaa5cb078382b42d170b972b03cbee3a80c5f27072ed562824535d2b482a9f4) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--aws_vpc_site--reference--group-003.md#canonical-d3d141692807d2f9a43a103e50c448595cc9ff55b69a7f6cd265defa929a265a) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](data-sources--aws_vpc_site--reference--group-003.md#canonical-dbd489939edb1969af84e73a3e3c1f0b5c5d5bef9159e37d7a9a20ca60cf7549) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--aws_vpc_site--reference--group-003.md#canonical-5e25003778457a2f134e83e165c816e17b989050fc2819e3be9c4ff8bc8a9f30) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](data-sources--aws_vpc_site--reference--group-003.md#canonical-08efbd95755a5e3e394ff2d06f1381185127281d03e9165cd347dfde23e68247) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--aws_vpc_site--reference--group-003.md#canonical-950c03546dc9250f925cf87c23f0338b25a10dd34b8fff0df8f1c1a77fdbc58f) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](data-sources--aws_vpc_site--reference--group-003.md#canonical-95d19a754adf5ded6434167660dad58e1f3ae527bc39de3e82669b1a1560ee94) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.type](data-sources--aws_vpc_site--reference--group-003.md#canonical-9e95b4a4bb1ebf12cb990ae640ed555574f99755757b6f39cc5ba667488b71f3) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_vpc_site--reference--group-003.md#canonical-57f8f8351946d453a277b4e13c63d8ee031e2d95eb585d9266f97380ec5b98d3) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--aws_vpc_site--reference--group-003.md#canonical-ac66323f713b95d8cdb8bbd40501e7640aec82e2d1574decb8f41eb7ded6732f) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](data-sources--aws_vpc_site--reference--group-003.md#canonical-3918196308effc8091a4a88c37ba0eb1053a7d1c57dc25ca00e1cb38f767ecf3) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](data-sources--aws_vpc_site--reference--group-003.md#canonical-806a48b18e8551a65c6c61e4a4853d8abffc8713baeedf8712226d021915e5ce) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--aws_vpc_site--reference--group-003.md#canonical-eb5b727f9ec9b7a2d183cf77f62928836430d8836c90667855c427efc01de0dd) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](data-sources--aws_vpc_site--reference--group-003.md#canonical-7c18324e7ef83850712e51854881632847e0a76bb8981d783991b9d7eec671fd) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](data-sources--aws_vpc_site--reference--group-003.md#canonical-94fe530a982ab5ef7f03a9611c162b528f6b59ab93087e6d4e588574f70b1dd7) |
| `ingress_egress_gw.outside_static_routes.static_route_list.simple_static_route` | [ingress_egress_gw.outside_static_routes.static_route_list.simple_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-90faa0c97745e9de80a6ff8cdbd7f95e95f2b385b84be305d0f961cd3d1c2a15) |
| `ingress_egress_gw.performance_enhancement_mode` | [ingress_egress_gw.performance_enhancement_mode](data-sources--aws_vpc_site--reference--group-003.md#canonical-402aa54751300aa07c4b6ce3bc5fcee67da7903f513daa549ca5fa858df8a2e2) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--aws_vpc_site--reference--group-003.md#canonical-88b33916531b65ca9610c864c015e2ee7a66f0586fffd91ac00151142e839b67) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--aws_vpc_site--reference--group-003.md#canonical-9dbce07a68dd4912563d241e735b4575e216e8e13cee52f23b80b4e3351b471f) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--aws_vpc_site--reference--group-003.md#canonical-87ab1d6c86a40f1b9d3d7fc68ef74e6cdfbd82cf5cbd026acc0c59febf1a0308) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--aws_vpc_site--reference--group-003.md#canonical-827e73b343c52b3e77c0318a3bfd155b80a37821f2c386dab45e53a350c55fb7) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--aws_vpc_site--reference--group-003.md#canonical-c283e30c078468d3b2a52b5c7d26f85e60ca5c97f8a8a8befdcd53313805a482) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--aws_vpc_site--reference--group-003.md#canonical-d6d43b0338e5e7491406c890520db1077ca9554b0fbf31a3a4b326c7c42c95d9) |
| `ingress_egress_gw.sm_connection_public_ip` | [ingress_egress_gw.sm_connection_public_ip](data-sources--aws_vpc_site--reference--group-003.md#canonical-f82fd0f8963bdb974715a68a1f2d818226d64bcab330558bbc1595bf615f5535) |
| `ingress_egress_gw.sm_connection_pvt_ip` | [ingress_egress_gw.sm_connection_pvt_ip](data-sources--aws_vpc_site--reference--group-003.md#canonical-18ff9de87bfe5331e02abc3431961536687d1fac391a9fd16c67d34d1367d97f) |
| `ingress_gw` | [ingress_gw](data-sources--aws_vpc_site--reference--group-003.md#canonical-2587d0e59e93e14e4342d3ebfcc4f857e71f684d5db3c70a77657e26d368a922) |
| `ingress_gw.allowed_vip_port` | [ingress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-003.md#canonical-a287a215ee27bb50399eb59fcdc6e56d681a502b78189ff406ea6caa22fbabad) |
| `ingress_gw.allowed_vip_port.custom_ports` | [ingress_gw.allowed_vip_port.custom_ports](data-sources--aws_vpc_site--reference--group-003.md#canonical-62443e7e4887416ea3d5509b6d3534e3fd8d2a9b676f50416bcdc7f646919167) |
| `ingress_gw.allowed_vip_port.custom_ports.port_ranges` | [ingress_gw.allowed_vip_port.custom_ports.port_ranges](data-sources--aws_vpc_site--reference--group-003.md#canonical-7565f932fc6846d831fa4d42adcd88a325306f9cd0a7e604f0e287e4668d6b33) |
| `ingress_gw.allowed_vip_port.disable_allowed_vip_port` | [ingress_gw.allowed_vip_port.disable_allowed_vip_port](data-sources--aws_vpc_site--reference--group-003.md#canonical-247e119991e51f208b9abfdb9745c174af79956d83245025d267bb1248f1cc5e) |
| `ingress_gw.allowed_vip_port.use_http_https_port` | [ingress_gw.allowed_vip_port.use_http_https_port](data-sources--aws_vpc_site--reference--group-003.md#canonical-531624817125e6938835eb736f8213fcd0f336d554aa25510d9657876259f906) |
| `ingress_gw.allowed_vip_port.use_http_port` | [ingress_gw.allowed_vip_port.use_http_port](data-sources--aws_vpc_site--reference--group-003.md#canonical-7b5a2c703c55eafa837bac7015d90d8ecf38f9295ecd62517cff8d8b12843858) |
| `ingress_gw.allowed_vip_port.use_https_port` | [ingress_gw.allowed_vip_port.use_https_port](data-sources--aws_vpc_site--reference--group-003.md#canonical-3cff10e7098fec0af06aeae27b764b452510ed0e6148ad98897b894da18b4dcf) |
| `ingress_gw.aws_certified_hw` | [ingress_gw.aws_certified_hw](data-sources--aws_vpc_site--reference--group-003.md#canonical-3e727f28e74e1b0c9a9f14a2a653d50c36f3ec13ea26e7d601fa435bda7ece01) |
| `ingress_gw.az_nodes` | [ingress_gw.az_nodes](data-sources--aws_vpc_site--reference--group-003.md#canonical-3222ef47d68890ac071fda538470f3817f809ee4c33af77526590fea74ad6423) |
| `ingress_gw.az_nodes.aws_az_name` | [ingress_gw.az_nodes.aws_az_name](data-sources--aws_vpc_site--reference--group-003.md#canonical-9312b4c7b3fa543795aeab5a7b6a78bb10da642f6afd3288fd137d5f728468e3) |
| `ingress_gw.az_nodes.local_subnet` | [ingress_gw.az_nodes.local_subnet](data-sources--aws_vpc_site--reference--group-003.md#canonical-e5589b3bf89bc336f6179edc96476f70c767a7d269eba5340109a699d4f23049) |
| `ingress_gw.az_nodes.local_subnet.existing_subnet_id` | [ingress_gw.az_nodes.local_subnet.existing_subnet_id](data-sources--aws_vpc_site--reference--group-003.md#canonical-dcdb9ba03c086e042a69415d0a0c72a88959214a98f14c879649d74cd47e33f4) |
| `ingress_gw.az_nodes.local_subnet.subnet_param` | [ingress_gw.az_nodes.local_subnet.subnet_param](data-sources--aws_vpc_site--reference--group-003.md#canonical-5b02213ea47c94bba2b7d6e5dab8910724d1d3c5d7a41cc2a3d2eeaf377653ad) |
| `ingress_gw.az_nodes.local_subnet.subnet_param.ipv4` | [ingress_gw.az_nodes.local_subnet.subnet_param.ipv4](data-sources--aws_vpc_site--reference--group-003.md#canonical-3ac1d4b65639a939524eedefc567de351882fa8183e6905000e8d8081acd0060) |
| `ingress_gw.performance_enhancement_mode` | [ingress_gw.performance_enhancement_mode](data-sources--aws_vpc_site--reference--group-003.md#canonical-c6e2e83305340a68d0500e87322b379dd1d4fbc0ffef419b166d83264d25bfb3) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--aws_vpc_site--reference--group-003.md#canonical-02db3744fcb01af65a678d5dbdca1f3b45fc66b918600ace4482bf04e1593ba4) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--aws_vpc_site--reference--group-004.md#canonical-d6b91178a47dcf3ed4311283007e6fc307c50bb6a5ab421e489d53b8dfc3a98c) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--aws_vpc_site--reference--group-004.md#canonical-1463263be5776f6b1c0543420e136eb78cc524ba7345bc7d48347ea594510763) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--aws_vpc_site--reference--group-004.md#canonical-d02a547b2f9d7d91dc1c0644733f4e5afa6a1ef9291333c5a042fc10a7569e2d) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--aws_vpc_site--reference--group-004.md#canonical-9ca0bc569209aa5544d84c72a426205c8618cbb0ff9264f9bc96ca0e8479848e) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--aws_vpc_site--reference--group-004.md#canonical-c2e634966612daebb6f36af1b2284ae24519bc7d097de5aa09f624a5b1539d61) |
| `instance_type` | [instance_type](data-sources--aws_vpc_site--reference--group-001.md#canonical-30b1865391e8cd43376af723820a8512dd3048141366b41895246b2f94af110b) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](data-sources--aws_vpc_site--reference--group-004.md#canonical-9c50ac265de96864502ca23191a28283559b3c437bb8c2a537185fd9b6d47a05) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](data-sources--aws_vpc_site--reference--group-004.md#canonical-97274c053034e1f322944bed45344a07f98b46e38f83cc583aa227f7bd0c1a13) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--aws_vpc_site--reference--group-004.md#canonical-421bd84ca02db29fc6736309e8efeabfdaaa6de11535573f62b473b440c6a28c) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](data-sources--aws_vpc_site--reference--group-004.md#canonical-1f93bf156ed4d6426d136881f27fb2a9fc7642f8062153e9f9392aa39cb9ef71) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](data-sources--aws_vpc_site--reference--group-004.md#canonical-24cf19cabefbf601345c32ff4edb4d262012cb3abbf64cd1ae3c78c205604532) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](data-sources--aws_vpc_site--reference--group-004.md#canonical-5432b0b8ccadb27e0d4a88c6ced9bc67e07f51aba5b9b14c2df26e0ca75d6ed8) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](data-sources--aws_vpc_site--reference--group-004.md#canonical-3ca9b3a5db9ada30c077eeb37592f73dee848a94cb1b440e2134ab3da52e4834) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](data-sources--aws_vpc_site--reference--group-004.md#canonical-fb7c49e9cf2b0af7ca95448d36af7a1387c0e96b64920a54b969dc863948fe91) |
| `labels` | [labels](data-sources--aws_vpc_site--reference--group-001.md#canonical-f5e4d4d3f2be7824afac2b85208e11f4c1d54ac3cce487cd89fe54bb24e48d8b) |
| `log_receiver` | [log_receiver](data-sources--aws_vpc_site--reference--group-004.md#canonical-0af6513c861a6efcb8e25c1921b0f4a0f148a674ffe1a9772e26577629ed6c3e) |
| `log_receiver.name` | [log_receiver.name](data-sources--aws_vpc_site--reference--group-004.md#canonical-10aa77d02a171c51d356d41c5d46d38e48fe679223af3a91d85b191e8903a70d) |
| `log_receiver.namespace` | [log_receiver.namespace](data-sources--aws_vpc_site--reference--group-004.md#canonical-cc8325b0578b9c2d7d34f6608180254714971a624c1586bdf9937ffab0c71368) |
| `log_receiver.tenant` | [log_receiver.tenant](data-sources--aws_vpc_site--reference--group-004.md#canonical-1514f1282276521c289a5fe12a8bf02dee3ea6313e6198c67aab74c7c76963c7) |
| `logs_streaming_disabled` | [logs_streaming_disabled](data-sources--aws_vpc_site--reference--group-004.md#canonical-dbbbc558b5523d34a19a80516af91e1c498573c058b4441b735ce1209cce768c) |
| `manual_routing` | [manual_routing](data-sources--aws_vpc_site--reference--group-004.md#canonical-006673fa3223785abfa75471531e49830b0bab9c65f232e240dc7fd6200a09ee) |
| `name` | [name](data-sources--aws_vpc_site--reference--group-001.md#canonical-aaa49b161a2dfdb50ed5475aac9f22ffe33eabb40c240f01d77b1552afdf78ae) |
| `namespace` | [namespace](data-sources--aws_vpc_site--reference--group-001.md#canonical-b9dbe80fcc28c1f0875f8cc6377685f17641d72b3d8e99e8e08c0b2b69055a66) |
| `no_worker_nodes` | [no_worker_nodes](data-sources--aws_vpc_site--reference--group-004.md#canonical-5a5e94ddac27d92beb1b9c93d475d6012b1a063e66b97a9ec94c3dc8e1b0b48a) |
| `nodes_per_az` | [nodes_per_az](data-sources--aws_vpc_site--reference--group-001.md#canonical-2498bcfc13ac5c6ab1e2a88a6814fb1e8d2d48a53f9718e39099cfaee0f674f0) |
| `offline_survivability_mode` | [offline_survivability_mode](data-sources--aws_vpc_site--reference--group-004.md#canonical-fb04d29228bd93376e31cca82bb673d09b36a2229f88ce74bb49f2052f7179f4) |
| `offline_survivability_mode.enable_offline_survivability_mode` | [offline_survivability_mode.enable_offline_survivability_mode](data-sources--aws_vpc_site--reference--group-004.md#canonical-99c29492ef2aded351cb70d8e2cdeb27495c18d29a87272d510249ba46aabdd8) |
| `offline_survivability_mode.no_offline_survivability_mode` | [offline_survivability_mode.no_offline_survivability_mode](data-sources--aws_vpc_site--reference--group-004.md#canonical-78407d3abd8136c1d77012ad23020e2e9a483b04d3576bf294b33fed78011e19) |
| `os` | [os](data-sources--aws_vpc_site--reference--group-004.md#canonical-9f89af2f15b507dc909aa3057d31efea2be06fb0b592ebd6e7120d575d9d9242) |
| `os.default_os_version` | [os.default_os_version](data-sources--aws_vpc_site--reference--group-004.md#canonical-72b9280313b352d990bd08ecd3a741e2a730f60656b62a945d5d1f6c7a96e52b) |
| `os.operating_system_version` | [os.operating_system_version](data-sources--aws_vpc_site--reference--group-004.md#canonical-603001d640554b6f13841e01645051479e861c9ae800c44fd039e971423f15a1) |
| `private_connectivity` | [private_connectivity](data-sources--aws_vpc_site--reference--group-004.md#canonical-24ba5ea35ad5164775cda02456f1802f5c22784c2358d37c6fa4df6862c43fde) |
| `private_connectivity.cloud_link` | [private_connectivity.cloud_link](data-sources--aws_vpc_site--reference--group-004.md#canonical-a61abecd91ece4996d4bee03e2549fd4c7e3a8cf4ae74dfc379ec0c9e1b6706c) |
| `private_connectivity.cloud_link.name` | [private_connectivity.cloud_link.name](data-sources--aws_vpc_site--reference--group-004.md#canonical-f75bd3e13a2f20d685e27837092a0c83743d99b31a0db3236db289192150b7b1) |
| `private_connectivity.cloud_link.namespace` | [private_connectivity.cloud_link.namespace](data-sources--aws_vpc_site--reference--group-004.md#canonical-e497e4bddeff65d642b45baddcf86caa603dac8e4ff3ff07f0f5a6b8cf7dc98e) |
| `private_connectivity.cloud_link.tenant` | [private_connectivity.cloud_link.tenant](data-sources--aws_vpc_site--reference--group-004.md#canonical-1221a50283d4c81430cf453c072f9872b9cc9a9be53b202d2d99f193733caaf1) |
| `private_connectivity.inside` | [private_connectivity.inside](data-sources--aws_vpc_site--reference--group-004.md#canonical-bed2713adda1d33dd167fb0a27d5e8255d3c22d1472fddf992b95e6073bd1b4d) |
| `private_connectivity.outside` | [private_connectivity.outside](data-sources--aws_vpc_site--reference--group-004.md#canonical-20305a6e704741f3f5af1497907d5917b3710d1d68e6e7839a8f7ed46bd30893) |
| `ssh_key` | [ssh_key](data-sources--aws_vpc_site--reference--group-001.md#canonical-4ca32422f176713adf96dbe54b95d7bdb8da7c9d452d666b89fcc351adc5b651) |
| `sw` | [sw](data-sources--aws_vpc_site--reference--group-004.md#canonical-eb70b02c9f47c955d99df02ffc90e97515834033783e7587a1e7c1bd134c446e) |
| `sw.default_sw_version` | [sw.default_sw_version](data-sources--aws_vpc_site--reference--group-004.md#canonical-6bbd307f0d410c789a2ac652cb8b6ad0aa178e834dc182fc4d0cbfd3e6d9d628) |
| `sw.volterra_software_version` | [sw.volterra_software_version](data-sources--aws_vpc_site--reference--group-004.md#canonical-140e52676b247c6378a84fe06aad27810442ae2078c00b06d408cce70dcb32da) |
| `tags` | [tags](data-sources--aws_vpc_site--reference--group-001.md#canonical-c2dab72df5e4c9bb4bc60b9150c8bf00920b51240faaacdb6001108faa149ef8) |
| `total_nodes` | [total_nodes](data-sources--aws_vpc_site--reference--group-001.md#canonical-94e8332871a3bfac446cd90bbb0ac94cba7e61e3feedc6425a9e52946d6c27e6) |
| `voltstack_cluster` | [voltstack_cluster](data-sources--aws_vpc_site--reference--group-004.md#canonical-86b971c090b17b2b68826ac670d5250c7ea2fcc25680b525a5e4fba7db9673a0) |
| `voltstack_cluster.active_enhanced_firewall_policies` | [voltstack_cluster.active_enhanced_firewall_policies](data-sources--aws_vpc_site--reference--group-004.md#canonical-9d880735d9dc0034f2fa9556f7db46e974cc985b8347f25f650d3feb5555677b) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--aws_vpc_site--reference--group-004.md#canonical-294b3c48de69d46d42774ed8f6d7cd6cac8e3d41e3c85b761da0066894ce17a6) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.name](data-sources--aws_vpc_site--reference--group-004.md#canonical-ab8b6f4a7481be59dd63487f2b9d10072872f12e3d454c9ee4dd80a03c98a3e5) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](data-sources--aws_vpc_site--reference--group-004.md#canonical-94f28586c663db407642d41da86fb901abe5a372fe1c0d1511d460cc93b24254) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](data-sources--aws_vpc_site--reference--group-004.md#canonical-e000c5be7bbf8eedaa13235222a2634d32fc5c3b337bdaa0123c38d170e7073b) |
| `voltstack_cluster.active_forward_proxy_policies` | [voltstack_cluster.active_forward_proxy_policies](data-sources--aws_vpc_site--reference--group-004.md#canonical-5d36e2f6a239bffc0634a0e3e96e34eed6d1f0cf5c7a773bf68674443d2128a9) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies](data-sources--aws_vpc_site--reference--group-004.md#canonical-0e65b930b55c2db88903094dda36021519aac2a1109ec1c8ecf116a0fe5ead28) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.name` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.name](data-sources--aws_vpc_site--reference--group-004.md#canonical-0b5c049ec0f7066c59cf9f75594826f4a492d645cc1583ec982d1aa5f3fba00b) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.namespace` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.namespace](data-sources--aws_vpc_site--reference--group-004.md#canonical-4bbdac41dc73385e2e68b82b177107b60fd7dd3c36b7dfff43bb7145e5a5e683) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.tenant` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.tenant](data-sources--aws_vpc_site--reference--group-004.md#canonical-e4acb512e0fb58c79ae6d49676f7a3f8f93fe72e57c81dfeb83ec3e425270aeb) |
| `voltstack_cluster.active_network_policies` | [voltstack_cluster.active_network_policies](data-sources--aws_vpc_site--reference--group-004.md#canonical-7db2d462178f99f1db53b8c85a64b5db512c36a88b81e58b58dd7f0de6ea750b) |
| `voltstack_cluster.active_network_policies.network_policies` | [voltstack_cluster.active_network_policies.network_policies](data-sources--aws_vpc_site--reference--group-004.md#canonical-db8aa9ddb461a6bd39c2f6e51e54adbb6af0f5f063fd1b5cbc7257275b615cc7) |
| `voltstack_cluster.active_network_policies.network_policies.name` | [voltstack_cluster.active_network_policies.network_policies.name](data-sources--aws_vpc_site--reference--group-004.md#canonical-46f0d6fcc98ec2bd97f223e6f88be6b06acc455418dcd925f06b903a23c29c89) |
| `voltstack_cluster.active_network_policies.network_policies.namespace` | [voltstack_cluster.active_network_policies.network_policies.namespace](data-sources--aws_vpc_site--reference--group-004.md#canonical-9f38790c0600ac388715c8453f288d946e426a33c3611a7d9b338c3a462a71be) |
| `voltstack_cluster.active_network_policies.network_policies.tenant` | [voltstack_cluster.active_network_policies.network_policies.tenant](data-sources--aws_vpc_site--reference--group-004.md#canonical-5cf536f864fbab4d679cfffa729066464d4bc14b8eb2bfbe7dc8165df1c9e066) |
| `voltstack_cluster.allowed_vip_port` | [voltstack_cluster.allowed_vip_port](data-sources--aws_vpc_site--reference--group-004.md#canonical-b345e41cc73e90255816e96ceb3636f1a0f3386614699aa9b30916380d027947) |
| `voltstack_cluster.allowed_vip_port.custom_ports` | [voltstack_cluster.allowed_vip_port.custom_ports](data-sources--aws_vpc_site--reference--group-004.md#canonical-2d5c9267925f2c5578f4e7f72acd003a6e0398d134f70dde0d3e3d0181b43b9b) |
| `voltstack_cluster.allowed_vip_port.custom_ports.port_ranges` | [voltstack_cluster.allowed_vip_port.custom_ports.port_ranges](data-sources--aws_vpc_site--reference--group-004.md#canonical-7672f71f7c19fa41091c3f2ff4c2e304680cf80ca0a2fd74d5d8181e982b1836) |
| `voltstack_cluster.allowed_vip_port.disable_allowed_vip_port` | [voltstack_cluster.allowed_vip_port.disable_allowed_vip_port](data-sources--aws_vpc_site--reference--group-004.md#canonical-3ae2fb1160b9ff780ace2133a94c387a2e353aca33730c6c3f7ba9dce582aeae) |
| `voltstack_cluster.allowed_vip_port.use_http_https_port` | [voltstack_cluster.allowed_vip_port.use_http_https_port](data-sources--aws_vpc_site--reference--group-004.md#canonical-0e9a2af7cfb77eed7a2fabebbacc40bb6d02559f09c1754c1e5a949af7eeecf2) |
| `voltstack_cluster.allowed_vip_port.use_http_port` | [voltstack_cluster.allowed_vip_port.use_http_port](data-sources--aws_vpc_site--reference--group-004.md#canonical-0e30f91de4145c0c4459a928fe059fc32b8857db50c2cf85644918d8e3929355) |
| `voltstack_cluster.allowed_vip_port.use_https_port` | [voltstack_cluster.allowed_vip_port.use_https_port](data-sources--aws_vpc_site--reference--group-004.md#canonical-e30b9519660e936c3d7cf927d43e19153e48e63850d76ee5102005f0a746236e) |
| `voltstack_cluster.aws_certified_hw` | [voltstack_cluster.aws_certified_hw](data-sources--aws_vpc_site--reference--group-004.md#canonical-6111ae71eeac1bbc8ca45bb2af4d6ff9e467d6ef793c9973cd40bb39cb16c0f8) |
| `voltstack_cluster.az_nodes` | [voltstack_cluster.az_nodes](data-sources--aws_vpc_site--reference--group-004.md#canonical-553c243feb84a1d33089e6070ade46bc4a7c4102641f11b1b1973ff03bf950e3) |
| `voltstack_cluster.az_nodes.aws_az_name` | [voltstack_cluster.az_nodes.aws_az_name](data-sources--aws_vpc_site--reference--group-004.md#canonical-22e0e119365e544f7438292106417ccae7c3943dc5892b226ce8e4c044aa34a7) |
| `voltstack_cluster.az_nodes.local_subnet` | [voltstack_cluster.az_nodes.local_subnet](data-sources--aws_vpc_site--reference--group-004.md#canonical-4d46fb6ae59b027c3f34b22c4fc2c9875e280842aad65505dfc9c7695d33ed28) |
| `voltstack_cluster.az_nodes.local_subnet.existing_subnet_id` | [voltstack_cluster.az_nodes.local_subnet.existing_subnet_id](data-sources--aws_vpc_site--reference--group-004.md#canonical-76bdcf2da3afcd60a1eb752d92c12275dba39f0ad6457e60093ed5576217973c) |
| `voltstack_cluster.az_nodes.local_subnet.subnet_param` | [voltstack_cluster.az_nodes.local_subnet.subnet_param](data-sources--aws_vpc_site--reference--group-004.md#canonical-8f90166ea48a45cafd649d58259422c3acfa60aba635a37328e2cb9bac059cc2) |
| `voltstack_cluster.az_nodes.local_subnet.subnet_param.ipv4` | [voltstack_cluster.az_nodes.local_subnet.subnet_param.ipv4](data-sources--aws_vpc_site--reference--group-004.md#canonical-97400e3b05dd4e4d49b030efe2712dc8ace26397542e9f446e0027761e2c8096) |
| `voltstack_cluster.dc_cluster_group` | [voltstack_cluster.dc_cluster_group](data-sources--aws_vpc_site--reference--group-004.md#canonical-0d454d973a1b7bcb70c11b86640085f175247932c4ba0c40a5bbd698335b1768) |
| `voltstack_cluster.dc_cluster_group.name` | [voltstack_cluster.dc_cluster_group.name](data-sources--aws_vpc_site--reference--group-004.md#canonical-e6e769be1ef5a5d58d756eea7162af55da67b5083b091266dae96b9fdef45271) |
| `voltstack_cluster.dc_cluster_group.namespace` | [voltstack_cluster.dc_cluster_group.namespace](data-sources--aws_vpc_site--reference--group-004.md#canonical-23a20dfe1069668a4e3d00c5fd6947cb6a48b04003e86bb55964056841989114) |
| `voltstack_cluster.dc_cluster_group.tenant` | [voltstack_cluster.dc_cluster_group.tenant](data-sources--aws_vpc_site--reference--group-004.md#canonical-266510b4ba0adcea041c06f226762445d069001d01edf6f9edeeb6407e218b3a) |
| `voltstack_cluster.default_storage` | [voltstack_cluster.default_storage](data-sources--aws_vpc_site--reference--group-004.md#canonical-dd5058cc223841600823f412f3529dde4df235b8def331b8b87893ca898082e9) |
| `voltstack_cluster.forward_proxy_allow_all` | [voltstack_cluster.forward_proxy_allow_all](data-sources--aws_vpc_site--reference--group-004.md#canonical-1afe6bcc2808e116665ff9e7fa5efa8ed855e635d1778afbe3c969743b6ab143) |
| `voltstack_cluster.global_network_list` | [voltstack_cluster.global_network_list](data-sources--aws_vpc_site--reference--group-004.md#canonical-d53af095db74d005a13a99896b570aac44dca83ce63354e6f3bc991e950f938b) |
| `voltstack_cluster.global_network_list.global_network_connections` | [voltstack_cluster.global_network_list.global_network_connections](data-sources--aws_vpc_site--reference--group-004.md#canonical-22205f1faca9675b065fbc3ec693921618b024146b61c1168e62bb4fd27c3fe4) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](data-sources--aws_vpc_site--reference--group-004.md#canonical-2e89c1f74ed5c7ae628a1b5e4ca5ffec04551ddeffe61748fc9dbdac597f6862) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn](data-sources--aws_vpc_site--reference--group-004.md#canonical-5c75599c5d8e109cd48e12c832feab814091713293b28cf6a7daebe8a03d7d87) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](data-sources--aws_vpc_site--reference--group-004.md#canonical-b3042df81e7cbc20d5be88f0928d78051ab46c538c111ad8359d59aebbbcebc7) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](data-sources--aws_vpc_site--reference--group-004.md#canonical-f9b05b94fb350fea554da88c19d2d81c20bc5f65bfdca32e1830022f676f5ff5) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](data-sources--aws_vpc_site--reference--group-004.md#canonical-dc5ffda649709e856dfb520fb54eaef216cf2cdca7f7221e5e8261ebc03b86b3) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](data-sources--aws_vpc_site--reference--group-004.md#canonical-bcab050717115a6ef30446fb19164bce495cbcbcd30c96c888ada1de6a31bb0f) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn](data-sources--aws_vpc_site--reference--group-004.md#canonical-442e51644f8011015bcaf8f173aa933810c147c1cea5759ed3d1ab851a1860e2) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](data-sources--aws_vpc_site--reference--group-004.md#canonical-d3978a06e391308324afe2720b10877fbeb95c06759230d9a919c4702206987a) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](data-sources--aws_vpc_site--reference--group-004.md#canonical-5d9c78d2917ff7dc1bd5574c4c288a44809fd851126cec7d82c47c5ec1824727) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](data-sources--aws_vpc_site--reference--group-004.md#canonical-509dca68e584fcf29dbe2500b99ff3cad2ed56ba6c4233311edaa2c3ac076e09) |
| `voltstack_cluster.k8s_cluster` | [voltstack_cluster.k8s_cluster](data-sources--aws_vpc_site--reference--group-004.md#canonical-8881987e86516b6d117f893c21ba8002c3e29f083747f685defbbc359f06ee24) |
| `voltstack_cluster.k8s_cluster.name` | [voltstack_cluster.k8s_cluster.name](data-sources--aws_vpc_site--reference--group-004.md#canonical-c2cdc31b6e6e67de1084762b2aa3455d0ae49b3ff89070a6c996404c4b5c8d1d) |
| `voltstack_cluster.k8s_cluster.namespace` | [voltstack_cluster.k8s_cluster.namespace](data-sources--aws_vpc_site--reference--group-004.md#canonical-95359cca9abb96d9b803b747014fde05f82d54b60cff557875de1da68abd47b5) |
| `voltstack_cluster.k8s_cluster.tenant` | [voltstack_cluster.k8s_cluster.tenant](data-sources--aws_vpc_site--reference--group-004.md#canonical-58e7f71f7aefdd73b20deecba8ce3180c28110d913e72516665fd3193a3a3358) |
| `voltstack_cluster.no_dc_cluster_group` | [voltstack_cluster.no_dc_cluster_group](data-sources--aws_vpc_site--reference--group-004.md#canonical-94c9fc8ac4c55cbf86b2e847f68f6da15940fdb9ffacd654b795064cffda6f23) |
| `voltstack_cluster.no_forward_proxy` | [voltstack_cluster.no_forward_proxy](data-sources--aws_vpc_site--reference--group-004.md#canonical-02f4b704a49609b610e67ffd09c934b3ceb138528595a92cf3268c7f5dadb224) |
| `voltstack_cluster.no_global_network` | [voltstack_cluster.no_global_network](data-sources--aws_vpc_site--reference--group-004.md#canonical-bc9228325590f3454b036f044acb37105b22c8ed644e0ec43c225cad863b244c) |
| `voltstack_cluster.no_k8s_cluster` | [voltstack_cluster.no_k8s_cluster](data-sources--aws_vpc_site--reference--group-004.md#canonical-9f0e7c06011867ddf86c559cd867c4a3e346ad073c2ca7920cccb6abcf2fa3f6) |
| `voltstack_cluster.no_network_policy` | [voltstack_cluster.no_network_policy](data-sources--aws_vpc_site--reference--group-004.md#canonical-0f976db41e78b374215688b30017a8f0874d1586c089869551732046cfcccbde) |
| `voltstack_cluster.no_outside_static_routes` | [voltstack_cluster.no_outside_static_routes](data-sources--aws_vpc_site--reference--group-004.md#canonical-9a2156e19f8ef745760cf4d57d2c0ef65764d7e5dbba6bbf35e7a55ffebaac47) |
| `voltstack_cluster.outside_static_routes` | [voltstack_cluster.outside_static_routes](data-sources--aws_vpc_site--reference--group-004.md#canonical-75bb10b45513b42f7c6771a3e8959ee8ee77c912a044b7b82df8b6d2fcf00b1f) |
| `voltstack_cluster.outside_static_routes.static_route_list` | [voltstack_cluster.outside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-004.md#canonical-223651eefacc56d519a2963ead28b8f24143ea2da85546207338b48705f154bc) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-004.md#canonical-6573a520cb4a4346a7eeb16668e28cab65d762639a2c207f62b2bf81cad02489) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.attrs` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.attrs](data-sources--aws_vpc_site--reference--group-004.md#canonical-1af0ebab950f3e1b694bfc504c78e4c2baeb76175898073b493e0e5e0e09e6a5) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels](data-sources--aws_vpc_site--reference--group-004.md#canonical-cb41131ee3579d8f98b9992c6c854fd55579d486822dc76ddea8545e430fd90d) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-004.md#canonical-29021a9505cdf906668898bbda0252a740442e16763b3ab88b8cc6129d62db46) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--aws_vpc_site--reference--group-005.md#canonical-a91516dc853dfb59f011abbb2367553a8e4d44cc4be0994c7b1cea88fefab9a0) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](data-sources--aws_vpc_site--reference--group-005.md#canonical-fc918de0222c5145b7cad5978a0ed7014036f808ef780c6aa6100d6ffd9e55b6) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](data-sources--aws_vpc_site--reference--group-005.md#canonical-4f6d01539b122512445f4f9b45925e6541933fcc229e66c5de6475711e47df00) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](data-sources--aws_vpc_site--reference--group-005.md#canonical-2f7b15c4eb82208326e65d5b1e0c8df50ff069e1006195c146b325335177925c) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](data-sources--aws_vpc_site--reference--group-005.md#canonical-8075fd652d3794df0a83d7dbbf588be74b62bb1bfed4ed1865b50bf232cd740a) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](data-sources--aws_vpc_site--reference--group-005.md#canonical-c802f3840b787b89c22f0addb7d3978859fd6f7ce9b51d86d055c8ef874fff39) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-005.md#canonical-cc37e2af2764962d50e2e85a48e6f850eecdd6450a458a999de463e98f9ce40c) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_vpc_site--reference--group-005.md#canonical-cb49e1d2b69bfdf214b7a47814d28f75c8136b2e6fc3ebe9ceafaa22d2f8b86f) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--aws_vpc_site--reference--group-005.md#canonical-8881e421a8ce6024aae66fab3adf4177007b8e1b0b35df9b6ab26650a9210bf7) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](data-sources--aws_vpc_site--reference--group-005.md#canonical-bdb3a34e171eeaa0b516b5e7a9b2564d751c4074638155b3f3215b31877ca1fb) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--aws_vpc_site--reference--group-005.md#canonical-57af923f27067f0bef1b643a163d04ee79608fa8a3bc8d4c8509a174f8e2f5a1) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](data-sources--aws_vpc_site--reference--group-005.md#canonical-b36ff329768a100c6f9d5b671a74428e2dac6539f556bace573a88de4b7a31b5) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--aws_vpc_site--reference--group-005.md#canonical-77229533f077d8a4a4cafb86f8cb9b1a35b48b7add5f43d9841d55d4f4168702) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](data-sources--aws_vpc_site--reference--group-005.md#canonical-98d9814efe7991b926193901edfdbc731cf7c4e7f61fd954b2ccb343640beffd) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--aws_vpc_site--reference--group-005.md#canonical-b1f65d5acb898a25f831b1f33a63aa9fe2fd7b2dff9b255ea266908760b22cc1) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](data-sources--aws_vpc_site--reference--group-005.md#canonical-30cf87d4b52c20b594145d1aa157623ef4482ea5f664925dcea5d62c5f3ebee9) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.type](data-sources--aws_vpc_site--reference--group-005.md#canonical-3e0f615565540c7168beee69f2deb0ea6d4ca4c46ce60809e6ffa33fd73641fb) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_vpc_site--reference--group-005.md#canonical-672f39e85c9fdf072317aca3f40a7abef1d9a27e00a02597e5946ef64497cecf) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--aws_vpc_site--reference--group-005.md#canonical-ac12620670a209f5166ce200322c16e9d6af5738c30787951e114480daa626da) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](data-sources--aws_vpc_site--reference--group-005.md#canonical-6a6b98882adfedc67c0d956ceb460b6ff35b82abcc9a33815193b48169e31f2a) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](data-sources--aws_vpc_site--reference--group-005.md#canonical-10a741faf449e4bffbc02b7db374b4b90a50df0c400822fe22efbb20fb2400a6) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--aws_vpc_site--reference--group-005.md#canonical-0d5580d4e1c1dab90da025b675ea5d352e9ec2636a9fdc42af9a5049d258ee88) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](data-sources--aws_vpc_site--reference--group-005.md#canonical-ab020fcbae4852c93dae29a8cf749311c64789a31e097ef86b04f38016bcd3d0) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](data-sources--aws_vpc_site--reference--group-005.md#canonical-b01f550f45ff3592bdf84b544e81d42e180512887ee7c26a822f32020b556d5f) |
| `voltstack_cluster.outside_static_routes.static_route_list.simple_static_route` | [voltstack_cluster.outside_static_routes.static_route_list.simple_static_route](data-sources--aws_vpc_site--reference--group-004.md#canonical-3ec0907e2079f7dbc0e8bc4c82c118015029d1e78ce631a5ae7e292d9d8b17d7) |
| `voltstack_cluster.sm_connection_public_ip` | [voltstack_cluster.sm_connection_public_ip](data-sources--aws_vpc_site--reference--group-005.md#canonical-97857809b6e4aaf698e2af018369b2cb71e290a736c573d0c3b92a741ccdeed3) |
| `voltstack_cluster.sm_connection_pvt_ip` | [voltstack_cluster.sm_connection_pvt_ip](data-sources--aws_vpc_site--reference--group-005.md#canonical-8e5ad07413ed37e6218a8390a4a915685ab48b3c5af4c1d81f61f0bdba77c931) |
| `voltstack_cluster.storage_class_list` | [voltstack_cluster.storage_class_list](data-sources--aws_vpc_site--reference--group-005.md#canonical-071b825154ef5ff42e2b2eb01aa5669579d44b9b0354af443f14769848cac3cc) |
| `voltstack_cluster.storage_class_list.storage_classes` | [voltstack_cluster.storage_class_list.storage_classes](data-sources--aws_vpc_site--reference--group-005.md#canonical-9ba2c7dfc1e6ee5bd9723e300341f6d64c75d5c34c6bc205b73f350376d8f886) |
| `voltstack_cluster.storage_class_list.storage_classes.default_storage_class` | [voltstack_cluster.storage_class_list.storage_classes.default_storage_class](data-sources--aws_vpc_site--reference--group-005.md#canonical-a613fc40c316ce6e283381a182e20a2f7e9ad37f17c7d8e96651b0eea4bf291a) |
| `voltstack_cluster.storage_class_list.storage_classes.storage_class_name` | [voltstack_cluster.storage_class_list.storage_classes.storage_class_name](data-sources--aws_vpc_site--reference--group-005.md#canonical-ddd0b3b50a6a7e3d5f6fa456d3498f38ebe4498b4b0ada470bc78fb76d5b7592) |
| `vpc` | [vpc](data-sources--aws_vpc_site--reference--group-005.md#canonical-0168a69df836d59bf8bcdf842afa7944a40e90f9cf9b59b64b7063b32141cd71) |
| `vpc.new_vpc` | [vpc.new_vpc](data-sources--aws_vpc_site--reference--group-005.md#canonical-458f41240440241bd926f9145cdba59c9db071450fe3d574027217ded5a730d6) |
| `vpc.new_vpc.autogenerate` | [vpc.new_vpc.autogenerate](data-sources--aws_vpc_site--reference--group-005.md#canonical-4f17d001a7ddfddd0fabf2c4e978587f46e60dd4508e17ac3074562e3b4a1261) |
| `vpc.new_vpc.name_tag` | [vpc.new_vpc.name_tag](data-sources--aws_vpc_site--reference--group-005.md#canonical-e8e2e033a02c013a17cf6ae162f84181e2412faef59152860dc8c7dcb988de8f) |
| `vpc.new_vpc.primary_ipv4` | [vpc.new_vpc.primary_ipv4](data-sources--aws_vpc_site--reference--group-005.md#canonical-f5456b797aa4a6aa34786e22aaade28ff289bf344f7c97e30ee0ba0d1c6f076c) |
| `vpc.vpc_id` | [vpc.vpc_id](data-sources--aws_vpc_site--reference--group-005.md#canonical-d26f494a21e204a94afae9a3e9787c1b1c5e341d8ce7dace2e3f5375f3922d69) |
| `waf_signatures` | [waf_signatures](data-sources--aws_vpc_site--reference--group-005.md#canonical-8ebffefa7527917c1526c33cdd0d25fdd82a056dfb700c3a99363af0eccbad85) |
| `waf_signatures.automatic` | [waf_signatures.automatic](data-sources--aws_vpc_site--reference--group-005.md#canonical-bc736858814028dbf28a63691e3fc0e4aede256d82f72ca851df9ab0b715f8c6) |
| `waf_signatures.manual` | [waf_signatures.manual](data-sources--aws_vpc_site--reference--group-005.md#canonical-923c39f5de1c0e50537b66f5688143106a161e710bad00e4d82b7a3087d0b0d6) |

<a id="canonical-1ef0ae60d595fcfa4a04d3c2edc7682055f059dcdbcf34f54fd3dd492a05b622"></a>

## Next pages — Property reference / 4602ed3d6274 / 19

- [admin_password](data-sources--aws_vpc_site--reference--group-001.md#canonical-cf61adb96e3255226dfecf4b2896b0db41c666d77e2b0b5f6245e0277b70c590)
- [aws_cred](data-sources--aws_vpc_site--reference--group-001.md#canonical-49afbd2ef73399bad085c85ad245ccdfff8e955e7d434c32061687434d272168)
- [block_all_services](data-sources--aws_vpc_site--reference--group-001.md#canonical-b76dad76c53b4b19f8e8a154258132c8821bf9eb73ef0a9b94c8f4bd894f6ab7)
- [blocked_services](data-sources--aws_vpc_site--reference--group-001.md#canonical-d032a2669236aad9149c48eacbb9b9406c7baf6a30a2cba5c0b1773f6c359df0)
- [coordinates](data-sources--aws_vpc_site--reference--group-001.md#canonical-8e3fff956ccd777564613ff46454763434a51cac9b780741e9eb4b2807b63726)
- [custom_dns](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ecf2d6efcd15f9d35d6e047898741be3bc651f590662021e0080d0845258cf0)
- [custom_security_group](data-sources--aws_vpc_site--reference--group-001.md#canonical-b0731941313bd2cd269ea4438ac865999a040ab6e6f79209ee26f8f25364d2b3)
- [default_blocked_services](data-sources--aws_vpc_site--reference--group-002.md#canonical-a07d6b9e92df65b11b26bcede2ff5ea151145741fb6404881ba827c9c45df362)
- [direct_connect_disabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-99c0899d6bf0f8198853a25a8ad9829c746898d5f8959e6a601b46e61cfcbcf1)
- [direct_connect_enabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-0573a5c999b2ef934474e06f6951d3bd7be0ac2c426b6cad83da0f5a329e4a6e)
- [disable_encryption](data-sources--aws_vpc_site--reference--group-002.md#canonical-7516c1587d0df53ad876cae0477b2c38385dec4dd103145bda5084516dbca326)
- [disable_internet_vip](data-sources--aws_vpc_site--reference--group-002.md#canonical-1fa273c1c1a47fed8481715b68dfeeea56376ac2fd09470b7846e01af544bae7)
- [egress_gateway_default](data-sources--aws_vpc_site--reference--group-002.md#canonical-27eda6087dad60ebfab2bd9d183d63d4fd3d5ff5c8054c9ab718e31ae3b124e5)
- [egress_nat_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e4ad9435e0b0e6716c7f67d167fce5206201d39fc76b01abb1877385ec5606b)
- [egress_virtual_private_gateway](data-sources--aws_vpc_site--reference--group-002.md#canonical-408490482dd6c239b209790d7fbc93eab5b8299b6bc34c608b552cf8779a4ac7)
- [enable_encryption](data-sources--aws_vpc_site--reference--group-002.md#canonical-ebce9f17e54eded7f5f4b5f98bed9bd8bc031751d987e072803886c4bfa51dfb)
- [enable_internet_vip](data-sources--aws_vpc_site--reference--group-002.md#canonical-b2e9032440c417e11bfaa49c40aaf20079c79a6264edb2d63f11954983deeeca)
- [f5_orchestrated_routing](data-sources--aws_vpc_site--reference--group-002.md#canonical-cc60cce8b152be015a54c11c552724934297d0d6a25ba937a96666f5a1e6ba16)
- [f5xc_security_group](data-sources--aws_vpc_site--reference--group-002.md#canonical-e52d2230b230a613685d90c39f2ec99ce98ccd1da9751e604aed0656f9695db2)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-9e9cfba6f9af5174603e76117bd78bdda89bd50417a2f9a90f00ee2a6e665f9d)
- [ingress_gw](data-sources--aws_vpc_site--reference--group-003.md#canonical-729ff7fed866124cf88a6351d29b823249a97f34b315308350e2b160cf54482b)
- [kubernetes_upgrade_drain](data-sources--aws_vpc_site--reference--group-004.md#canonical-59427859ce0ecf818307189237b8e8aa6b84229ce4372bc1a87c0e5503bccea9)
- [log_receiver](data-sources--aws_vpc_site--reference--group-004.md#canonical-1b88ba6da7c0147d0dad13276774e0bac834e226743ad35d599608626827317c)
- [logs_streaming_disabled](data-sources--aws_vpc_site--reference--group-004.md#canonical-2b8539e757fd3d9ffb4d3a36d093e0bd00f97a60629b501630154cfb65f30443)
- [manual_routing](data-sources--aws_vpc_site--reference--group-004.md#canonical-b195dd72eff5ecd49a6772d9d3512c0bb0648fb63e7b4b1b6248732bfe89f956)
- [no_worker_nodes](data-sources--aws_vpc_site--reference--group-004.md#canonical-f61dc2f597dcfadab5dcd644ce7cafde336974edd69d7e6dc56ac3e6cfeb7a80)
- [offline_survivability_mode](data-sources--aws_vpc_site--reference--group-004.md#canonical-1159cfbac09c60f8602b2b8e1fbb30d2be51e0e94e95e31a3745fbf56d1abd44)
- [os](data-sources--aws_vpc_site--reference--group-004.md#canonical-66baedec83b86131b792251eedb3fe856dec7e0caa69a3a1f165865f86da46d2)
- [private_connectivity](data-sources--aws_vpc_site--reference--group-004.md#canonical-ea5f3bd2603cd639f551392ffb29a7d414cf78ca1f13f34511f60a7f0b3e2fef)
- [sw](data-sources--aws_vpc_site--reference--group-004.md#canonical-7ad1e1c86a364ee2ccd7fd4e86d325460994f69655dcf789db6b6d631c261de9)
- [voltstack_cluster](data-sources--aws_vpc_site--reference--group-004.md#canonical-ce03987dc4fa70508f8e432b0fc8b0c65d9dacdc10695d71f9c4f1c61d03dbb9)
- [vpc](data-sources--aws_vpc_site--reference--group-005.md#canonical-c482fd7fdef218c974f1bc3d19e9ec98b067330464c5d118dc6b9acff264f21e)
- [waf_signatures](data-sources--aws_vpc_site--reference--group-005.md#canonical-1aaf17a2664f507686e37894dc171e25232558868c23cf23a2ff535374e932d3)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-cf61adb96e3255226dfecf4b2896b0db41c666d77e2b0b5f6245e0277b70c590"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-855087556f44e41dbc2b3a82e71ef5f86c040fcd5bd068e31828aa44c6226f04"></a>

## admin_password — admin_password / 5b3e414320f6 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- admin_password

<a id="canonical-fc291d0108076a236e1c4781fef57ae3ed27cfb4a03b4241230663d36b5f7070"></a>

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

<a id="canonical-4d90145ae352b7ea026276f59cddf3678c66ee68ccc87c3c234b0b886521fb2e"></a>

## Direct properties — admin_password / 5b3e414320f6 / 3

- [blindfold_secret_info](data-sources--aws_vpc_site--reference--group-001.md#canonical-f73b462f8d72535f2dc3cfe54c9e4041085fc4369bc739c48348fa120436267f): complete subsection reference.

- [clear_secret_info](data-sources--aws_vpc_site--reference--group-001.md#canonical-7db7ba0419d99e4efaf839652df741e3e8f03ae6a0ec1c5286a3d5ab5ac83f63): complete subsection reference.

<a id="canonical-053c6fd88dc45d2cfa5b670898fa7ed2d80cbd0f36406469b15f059c9b1f141c"></a>

## Next pages — admin_password / 5b3e414320f6 / 4

- [admin_password.blindfold_secret_info](data-sources--aws_vpc_site--reference--group-001.md#canonical-f73b462f8d72535f2dc3cfe54c9e4041085fc4369bc739c48348fa120436267f)
- [admin_password.clear_secret_info](data-sources--aws_vpc_site--reference--group-001.md#canonical-7db7ba0419d99e4efaf839652df741e3e8f03ae6a0ec1c5286a3d5ab5ac83f63)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-f73b462f8d72535f2dc3cfe54c9e4041085fc4369bc739c48348fa120436267f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-414fbc8c1a01adc01e97890a6ae77a0a26de107ff0c85eceb92517482e608e61"></a>

## admin_password.blindfold_secret_info — admin_password.blindfold_secret_info / 5c2f8d6e52a4 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [admin_password](data-sources--aws_vpc_site--reference--group-001.md#canonical-cf61adb96e3255226dfecf4b2896b0db41c666d77e2b0b5f6245e0277b70c590)
- admin_password.blindfold_secret_info

<a id="canonical-6f2722bc04c9be396235bfc73ed21a8fbee69f4df771009c264e41cd10a9803f"></a>

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

<a id="canonical-1df3bcf6f635aac1d5ba22c914a5d3b12a9dbed7753a48a7d059f3ed727cf399"></a>

## Direct properties — admin_password.blindfold_secret_info / 5c2f8d6e52a4 / 3

<a id="canonical-8b1afbb930f07e4ae9bbf7517fca2873f1cbd20d39d8ef3f6ac1e27671946de7"></a>

<a id="canonical-ba0ec92d0f0301517f74023b80de3fecf12bc443dddd10c62184af111edbfee9"></a>

## decryption_provider property — admin_password.blindfold_secret_info / 5c2f8d6e52a4 / 4

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

<a id="canonical-6c4f93fb40d6c4c1c5efc20e96f317a7d106a05c5800f3f607d0f6f4a05401fc"></a>

<a id="canonical-7ef5d9e0b76b7aa6f1d7800f2d65001d212bb048282eca88a646296e69c0e006"></a>

## location property — admin_password.blindfold_secret_info / 5c2f8d6e52a4 / 5

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

<a id="canonical-36a8c7aff4c6e4151b4b4f16b6cb955b9a9bedf73bd830133d3420018e903c9f"></a>

<a id="canonical-605f10f20e4090596ae9e3fbbaf6c6baf66aeca1088dab7ec86c5a3fb7a7ad48"></a>

## store_provider property — admin_password.blindfold_secret_info / 5c2f8d6e52a4 / 6

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

<a id="canonical-523f8a173716744a115341630ef769cc8080e10b4cadfb886a6ac3e8f7f7f5e1"></a>

## Next pages — admin_password.blindfold_secret_info / 5c2f8d6e52a4 / 7

- [admin_password](data-sources--aws_vpc_site--reference--group-001.md#canonical-cf61adb96e3255226dfecf4b2896b0db41c666d77e2b0b5f6245e0277b70c590)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-7db7ba0419d99e4efaf839652df741e3e8f03ae6a0ec1c5286a3d5ab5ac83f63"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9a61123130964cd4e8620249334afc20efa09b43f3d88747ceb4be36379cdb6a"></a>

## admin_password.clear_secret_info — admin_password.clear_secret_info / 2ad8c93c3b41 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [admin_password](data-sources--aws_vpc_site--reference--group-001.md#canonical-cf61adb96e3255226dfecf4b2896b0db41c666d77e2b0b5f6245e0277b70c590)
- admin_password.clear_secret_info

<a id="canonical-b31c8c1a3c838451dfd07df727c71be98e9d3d07a1b292e1391dea347ca3ff16"></a>

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

<a id="canonical-7e1b45e0a4a96ad86bcfc996f7bf6eb0de222e497b456e264df49811a4d2878e"></a>

## Direct properties — admin_password.clear_secret_info / 2ad8c93c3b41 / 3

<a id="canonical-351a09faadb7a301ed78f1a5740de6e28535622a3376f18e713b8708cc28c4b8"></a>

<a id="canonical-e438450338259b4b98c52c42514527d16756b5f2b5a57d34e170dd2ad08513f1"></a>

## provider_ref property — admin_password.clear_secret_info / 2ad8c93c3b41 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-36faf7acd05ac8b2e941c025cb47416043cbc1a3002be6ab7e33e1a71914117d"></a>

<a id="canonical-69fce74cabecbdabdbf885855cf419624a7bcc1415d4f9d0f5fa15f904385629"></a>

## url property — admin_password.clear_secret_info / 2ad8c93c3b41 / 5

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

<a id="canonical-159daf7a09a2cbd59cda921988968ac9517a0b461adbf812508b0b88702c9eb5"></a>

## Next pages — admin_password.clear_secret_info / 2ad8c93c3b41 / 6

- [admin_password](data-sources--aws_vpc_site--reference--group-001.md#canonical-cf61adb96e3255226dfecf4b2896b0db41c666d77e2b0b5f6245e0277b70c590)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-49afbd2ef73399bad085c85ad245ccdfff8e955e7d434c32061687434d272168"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c14f9de40d89179adf9ef5f40c071352f8dbb040159f09b6b1e29894a6ce2926"></a>

## aws_cred — aws_cred / 477ccc0e3ea2 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- aws_cred

<a id="canonical-ed248ef5f9d51d868c83d7bb88f448c8f42a8a365b0d83fe0912db8c605d6f02"></a>

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

<a id="canonical-f2f74dcf3cdb35f51ea6dda4ff6970eb451af2aaba6f56037746bbcea89a9d07"></a>

## Direct properties — aws_cred / 477ccc0e3ea2 / 3

<a id="canonical-f1ecc446525df5e6c845872fe1df093a8e9cafdddac455ba932eb9fd053a1372"></a>

<a id="canonical-22ec7674a7e0d7b0eecb1c81017e18b986909bbbceb06073734ae6d08e7e6a86"></a>

## name property — aws_cred / 477ccc0e3ea2 / 4

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

<a id="canonical-20e54c316d5c977054946199d295b626eab79ca2d4ec1db192560bfe4bc0d747"></a>

<a id="canonical-fa920d49611d31505fa73936cb0923b5f57cf8d5bf5542020186e6b56bed6bec"></a>

## namespace property — aws_cred / 477ccc0e3ea2 / 5

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

<a id="canonical-a5911cfd95ccaf5f2c5fb0c2a46d4d11c2875537bfed6a06a8b5fb5d1445d758"></a>

<a id="canonical-5d1b4f6962ddac9686e9411c7ce9f8c16a6ec97f88865d853005ba716f3e7523"></a>

## tenant property — aws_cred / 477ccc0e3ea2 / 6

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

<a id="canonical-260626cfd127c9fc292768d9b026226bca73711a420222a19f4a311de6a99f86"></a>

## Next pages — aws_cred / 477ccc0e3ea2 / 7

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-b76dad76c53b4b19f8e8a154258132c8821bf9eb73ef0a9b94c8f4bd894f6ab7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-96d9d3f2fa9c39eea26d6c314ebad3866875207b6ba84763dd5dd7b0e7794860"></a>

## block_all_services — block_all_services / 77d2ba778a6e / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- block_all_services

<a id="canonical-dcb5608889f17fb4fd10b86025288ca1e52df507032af5ece883d86825b98250"></a>

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

- [block_all_services](data-sources--aws_vpc_site--reference--group-001.md#canonical-dcb5608889f17fb4fd10b86025288ca1e52df507032af5ece883d86825b98250)
- [blocked_services](data-sources--aws_vpc_site--reference--group-001.md#canonical-01d4814845c77eee790efcbacb2e6ceec072086ae3d672dbd6b704e4e01a6719)
- [default_blocked_services](data-sources--aws_vpc_site--reference--group-002.md#canonical-16f59aaf40db507d24180fa9a4990a0a0b3bee4380e31b5e5a4d2a3449015f84)

Select alternatives according to the provider validators above.

<a id="canonical-b260cff1aa54915a17616a90451c6281ff7715869389e0b4efc125011f4035d5"></a>

## Direct properties — block_all_services / 77d2ba778a6e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-800afb48f4db26b57e11c50aca162a107980c38717b5ecaa8d57bc637b538137"></a>

## Next pages — block_all_services / 77d2ba778a6e / 4

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-d032a2669236aad9149c48eacbb9b9406c7baf6a30a2cba5c0b1773f6c359df0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e61973cd78219718b07f6bfc37269e4c0f4888eb5c98dc38d87985e165bfea3"></a>

## blocked_services — blocked_services / 0165abbadc64 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- blocked_services

<a id="canonical-01d4814845c77eee790efcbacb2e6ceec072086ae3d672dbd6b704e4e01a6719"></a>

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

<a id="canonical-0c32b35197ef67bf33a22546368bf1b79a89412dfa0434ae3489ece7e6e687ad"></a>

## Direct properties — blocked_services / 0165abbadc64 / 3

- [blocked_service](data-sources--aws_vpc_site--reference--group-001.md#canonical-d020118ea77b066c7c14355c46ca97382140e1d63f1018592518815248ea1b39): complete subsection reference.

<a id="canonical-adeb70b582ddb73125c262cea7ca93c0b879afefb1bd6f009fef164294ae62c4"></a>

## Next pages — blocked_services / 0165abbadc64 / 4

- [blocked_services.blocked_service](data-sources--aws_vpc_site--reference--group-001.md#canonical-d020118ea77b066c7c14355c46ca97382140e1d63f1018592518815248ea1b39)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-d020118ea77b066c7c14355c46ca97382140e1d63f1018592518815248ea1b39"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-96626ee2bd5373c83a9db418dbafae6be8696b9bf0cbf472b13c32218c6ed380"></a>

## blocked_services.blocked_service — blocked_services.blocked_service / 5529396b3b77 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [blocked_services](data-sources--aws_vpc_site--reference--group-001.md#canonical-d032a2669236aad9149c48eacbb9b9406c7baf6a30a2cba5c0b1773f6c359df0)
- blocked_services.blocked_service

<a id="canonical-d47764819710995a28c89be323fa87e2860f0ee7fc63de18d09c535ff7beb538"></a>

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

<a id="canonical-d84dc931496a2a0f3b7b1f6ae6019a48d94fd2634ed1e4f07f8585256e282fd8"></a>

## Direct properties — blocked_services.blocked_service / 5529396b3b77 / 3

- [dns](data-sources--aws_vpc_site--reference--group-001.md#canonical-cd7979239cf954b1dd2de236b0cc2a30a9e30370e435ce0d3d7c50837b61d8d0): complete subsection reference.

<a id="canonical-29fb40b6587280108d94cac7c630ae5b929fe868abb24a32956371a8153e4e59"></a>

<a id="canonical-42c5788424416b7778fed5424f3369c8184b9fcf2cb2fbedd24b53249aa9d037"></a>

## network_type property — blocked_services.blocked_service / 5529396b3b77 / 4

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

- [ssh](data-sources--aws_vpc_site--reference--group-001.md#canonical-bab59949b36476b5da96ed942194d670fe655e546750fed3a61eeb6441cca9da): complete subsection reference.

- [web_user_interface](data-sources--aws_vpc_site--reference--group-001.md#canonical-269b70e841657c7669aaa6148dde7ef4d1557db33374f7ffc9a66c53fe7ecba3): complete subsection reference.

<a id="canonical-7acf1de3f57777ccd968424f4d6bdc4be82f97b13ef896e5b80e82b9042072ec"></a>

## Next pages — blocked_services.blocked_service / 5529396b3b77 / 5

- [blocked_services.blocked_service.dns](data-sources--aws_vpc_site--reference--group-001.md#canonical-cd7979239cf954b1dd2de236b0cc2a30a9e30370e435ce0d3d7c50837b61d8d0)
- [blocked_services.blocked_service.ssh](data-sources--aws_vpc_site--reference--group-001.md#canonical-bab59949b36476b5da96ed942194d670fe655e546750fed3a61eeb6441cca9da)
- [blocked_services.blocked_service.web_user_interface](data-sources--aws_vpc_site--reference--group-001.md#canonical-269b70e841657c7669aaa6148dde7ef4d1557db33374f7ffc9a66c53fe7ecba3)
- [blocked_services](data-sources--aws_vpc_site--reference--group-001.md#canonical-d032a2669236aad9149c48eacbb9b9406c7baf6a30a2cba5c0b1773f6c359df0)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-cd7979239cf954b1dd2de236b0cc2a30a9e30370e435ce0d3d7c50837b61d8d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-589656bbade5b7d009e4223807a6e125404da0f0109cb80fb0b0152df692bdcd"></a>

## blocked_services.blocked_service.dns — blocked_services.blocked_service.dns / 9a6daa17aec6 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [blocked_services](data-sources--aws_vpc_site--reference--group-001.md#canonical-d032a2669236aad9149c48eacbb9b9406c7baf6a30a2cba5c0b1773f6c359df0)
- [blocked_services.blocked_service](data-sources--aws_vpc_site--reference--group-001.md#canonical-d020118ea77b066c7c14355c46ca97382140e1d63f1018592518815248ea1b39)
- blocked_services.blocked_service.dns

<a id="canonical-ab259de40d2b3ff98a1bbdbba3b6f34a04faa2430342191906b1ff148a7bc13b"></a>

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

<a id="canonical-e4707e531850f29ffc3d23cd94ce1898827e9ba191f80034614d380acef4fc78"></a>

## Direct properties — blocked_services.blocked_service.dns / 9a6daa17aec6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5b3205ece4821bf58fc543b60c4ee0c5e4ab83082a2f07e614c1e32ce94fd3f0"></a>

## Next pages — blocked_services.blocked_service.dns / 9a6daa17aec6 / 4

- [blocked_services.blocked_service](data-sources--aws_vpc_site--reference--group-001.md#canonical-d020118ea77b066c7c14355c46ca97382140e1d63f1018592518815248ea1b39)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-bab59949b36476b5da96ed942194d670fe655e546750fed3a61eeb6441cca9da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-abeae77b4b1aa51f33fb8c71d7eeb809e010c07d2f6e66479d6ceeb2eff3e5d0"></a>

## blocked_services.blocked_service.ssh — blocked_services.blocked_service.ssh / fc31b7090935 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [blocked_services](data-sources--aws_vpc_site--reference--group-001.md#canonical-d032a2669236aad9149c48eacbb9b9406c7baf6a30a2cba5c0b1773f6c359df0)
- [blocked_services.blocked_service](data-sources--aws_vpc_site--reference--group-001.md#canonical-d020118ea77b066c7c14355c46ca97382140e1d63f1018592518815248ea1b39)
- blocked_services.blocked_service.ssh

<a id="canonical-e1280c0a66514e315b03a7ce9405d02931e96fc028ae065ac7f9f1256e351c24"></a>

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

<a id="canonical-339ef40701fd3a75f6e524bbccd5e1fe59b5e2e96e91461e3744235259bb705b"></a>

## Direct properties — blocked_services.blocked_service.ssh / fc31b7090935 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-da524f7ce1d675eeb7eefea14414004a2732d549f42a4d96dff19d61be768454"></a>

## Next pages — blocked_services.blocked_service.ssh / fc31b7090935 / 4

- [blocked_services.blocked_service](data-sources--aws_vpc_site--reference--group-001.md#canonical-d020118ea77b066c7c14355c46ca97382140e1d63f1018592518815248ea1b39)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-269b70e841657c7669aaa6148dde7ef4d1557db33374f7ffc9a66c53fe7ecba3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6724a74f9c1de5054b23d6a1b6a8e04a3afab2d66714ddadcb739659cde6e65c"></a>

## blocked_services.blocked_service.web_user_interface — blocked_services.blocked_service.web_user_interface / c459c97db09a / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [blocked_services](data-sources--aws_vpc_site--reference--group-001.md#canonical-d032a2669236aad9149c48eacbb9b9406c7baf6a30a2cba5c0b1773f6c359df0)
- [blocked_services.blocked_service](data-sources--aws_vpc_site--reference--group-001.md#canonical-d020118ea77b066c7c14355c46ca97382140e1d63f1018592518815248ea1b39)
- blocked_services.blocked_service.web_user_interface

<a id="canonical-07c5b97e2ec618e1e71b0dea5fbd7ed6c8eb4e44cd19d972f8aa88805c2625e3"></a>

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

<a id="canonical-e70e929a9de5dbdb84a93d86e54bb001a37cb6a6819453e412e883e4de3edc25"></a>

## Direct properties — blocked_services.blocked_service.web_user_interface / c459c97db09a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-40a89ee2a21db6ce6034f522880c993256ee123029d7277e6fb06f731590c377"></a>

## Next pages — blocked_services.blocked_service.web_user_interface / c459c97db09a / 4

- [blocked_services.blocked_service](data-sources--aws_vpc_site--reference--group-001.md#canonical-d020118ea77b066c7c14355c46ca97382140e1d63f1018592518815248ea1b39)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-8e3fff956ccd777564613ff46454763434a51cac9b780741e9eb4b2807b63726"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72124fd7b89256fe83841f2c37751d8b56b038de4dc1ab9328e02f47a48ce135"></a>

## coordinates — coordinates / f3278c3a4833 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- coordinates

<a id="canonical-38e253c1a0624adbe8b38bca06c8759cb3a395663ebf5f924e7e85d11a37d8eb"></a>

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

<a id="canonical-416fc1b497abad44d72f9999e3d0078565b93f7369ab5abaf1490c7964e6e1c3"></a>

## Direct properties — coordinates / f3278c3a4833 / 3

<a id="canonical-7097d53ce0e870c19d8d94e0833b4126074279edfbfa242f2e03619cea445d1d"></a>

<a id="canonical-9667c4c109a74bb2c031b57071aa7b7d2de3f5d4933fcfe01b508df89c8bf8e7"></a>

## latitude property — coordinates / f3278c3a4833 / 4

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

<a id="canonical-f46901776730bc93c25ac7c2b202f9331ff37c62f8b4a6697b40afa941c41228"></a>

<a id="canonical-f98173482b1d8f9ab081c6fae71d8ec3d1781631e8774025e48c28e38588ce1e"></a>

## longitude property — coordinates / f3278c3a4833 / 5

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

<a id="canonical-e552740f35c6937db5de0d24da0b8fb2c7792ead82c514909638ad12b94a4c46"></a>

## Next pages — coordinates / f3278c3a4833 / 6

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-0ecf2d6efcd15f9d35d6e047898741be3bc651f590662021e0080d0845258cf0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-86100f5e4a9584bb07f227f61acd37b50ec2808615b1352f4b7dea9c16011922"></a>

## custom_dns — custom_dns / 7efe654abec2 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- custom_dns

<a id="canonical-a609cb474ffa654f25860da877b1f74674b8078fd224b09743477084bf59f6a1"></a>

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

<a id="canonical-72c3042b6ef2cd495ea4041e620af790d67da5597100bf4964b5c297fb7b447e"></a>

## Direct properties — custom_dns / 7efe654abec2 / 3

<a id="canonical-7bc6948ad76621b21d8c40dfce9cd921f637d196430606567f0e889b2d24fbb7"></a>

<a id="canonical-d16eaca37959e83a63b66b9418ac57632537c4bddf1f5f05e903fce8e0f085d0"></a>

## inside_nameserver property — custom_dns / 7efe654abec2 / 4

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

<a id="canonical-40171d8b490253947e2254b85e89c9e4b49fe527ddff98b201f4ddf16ff3cce1"></a>

<a id="canonical-45497ff044502927c30202450b9b837ac56eb2859f978094dfc15da80fe1b770"></a>

## outside_nameserver property — custom_dns / 7efe654abec2 / 5

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

<a id="canonical-2c3df09381ec806f71a634aab390b3a078bc0941d1abfd158d01f1c675192729"></a>

## Next pages — custom_dns / 7efe654abec2 / 6

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-b0731941313bd2cd269ea4438ac865999a040ab6e6f79209ee26f8f25364d2b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
