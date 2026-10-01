---
page_title: "xcsh_cloud_elastic_ip reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_elastic_ip reference."
---

# xcsh_cloud_elastic_ip reference

<a id="canonical-ef15d26eed8ba06029d76ff217be5d95cee1296437e2bcdfc5add660013e8dcb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72878f855bff32dffc280c1dbd9075935530e1d73b82a5bafa9ab81079746da0"></a>

## Property reference — Property reference / c7ae0427900f / 2

Breadcrumbs:

- [xcsh_cloud_elastic_ip](../data-sources/cloud_elastic_ip.md#canonical-2b2f30e89bfb277ce22595cc89aff6edb2de7c5d6f5d72c5d631bef5fa2ab292)
- Property reference

<a id="canonical-4bd7e13d3094f92e62a1e3e78b137c0efec042f5633597880b5971ee700e9a5f"></a>

## Direct properties — Property reference / c7ae0427900f / 3

<a id="canonical-725681ef06a155f41be86f382772bb5859cec2f016b990f987ddd42d50718f26"></a>

<a id="canonical-289d518f87359955d780ee1a79da2c6287092ec0d911f2e2012e73b1bf782113"></a>

## annotations property — Property reference / c7ae0427900f / 4

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

<a id="canonical-0033c3f597eda74540e0dd15db7b554db4aee923aa300d926a99e1b0808d7608"></a>

<a id="canonical-e2ae13bbcf30a5027b00a9200658289b5fa1ca35f6a34f6d251dbcdf0fc97058"></a>

## description property — Property reference / c7ae0427900f / 5

Type: `"string"`. Computed.

Description of the CloudElasticIP.

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

<a id="canonical-8ab23ac29f34d719e0f67564ae058009972453d81d364e9b98944bf1cfe0a9e4"></a>

<a id="canonical-f371aa2dbdd968af723699a5f3db9191ff1982d78300913e5686a668f88766d2"></a>

## id property — Property reference / c7ae0427900f / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-b26682dcce7e1069e63de7b05ff9d95b92bc2d9958fd119349c4d80169370eca"></a>

<a id="canonical-c343a0608dad0dcd6c042d6e2b988afeaec52bcb8bbda7f698b296a9a00985fd"></a>

## item_count property — Property reference / c7ae0427900f / 7

Type: `"number"`. Computed.

Number of Elastic Ips / Public Ips associated with this object per Node.

<a id="canonical-a006e7b483fa4adf3f3f39dea8f698088cd45987749bbfb9ddba3500bab28bdc"></a>

<a id="canonical-e61a9dee513a3f5e03a770aa031f7f2c426e89c15211cfef54890efd61540704"></a>

## labels property — Property reference / c7ae0427900f / 8

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

<a id="canonical-fab9d0b177e6bf1b12715037d5542f712546e43b1edb024fcf0671d81f053448"></a>

<a id="canonical-c48a49cadcd043de6d27dd6f7bd2f85679254b10831dd49cd014de0f56a0ccc6"></a>

## name property — Property reference / c7ae0427900f / 9

Type: `"string"`. Required.

Name of the CloudElasticIP.

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

<a id="canonical-76a51a394e0e415039b42699683643f9ad24f3bae4c3db5f8b8d2b8bbe3dd133"></a>

<a id="canonical-a2330f837246473cc75a11762700b19baf0ce05df1b9eba2cfcbd732809ce0f9"></a>

## namespace property — Property reference / c7ae0427900f / 10

Type: `"string"`. Required.

Namespace where the CloudElasticIP exists.

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

- [site_ref](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-cdd28936b4f4026b6c29bb0cf7632bd08291939f574e7c2bf62b1b7c37e060ff): complete subsection reference.

<a id="canonical-c6c1ecd02046a2a406a86274a0ceee108e524a421fa9d75ab29089db8bf6036c"></a>

## All schema paths — Property reference / c7ae0427900f / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-725681ef06a155f41be86f382772bb5859cec2f016b990f987ddd42d50718f26) |
| `description` | [description](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-0033c3f597eda74540e0dd15db7b554db4aee923aa300d926a99e1b0808d7608) |
| `id` | [id](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-8ab23ac29f34d719e0f67564ae058009972453d81d364e9b98944bf1cfe0a9e4) |
| `item_count` | [item_count](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-b26682dcce7e1069e63de7b05ff9d95b92bc2d9958fd119349c4d80169370eca) |
| `labels` | [labels](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-a006e7b483fa4adf3f3f39dea8f698088cd45987749bbfb9ddba3500bab28bdc) |
| `name` | [name](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-fab9d0b177e6bf1b12715037d5542f712546e43b1edb024fcf0671d81f053448) |
| `namespace` | [namespace](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-76a51a394e0e415039b42699683643f9ad24f3bae4c3db5f8b8d2b8bbe3dd133) |
| `site_ref` | [site_ref](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-9f847d2e4a6873ab02d547c72d9363ba1c358c8d2f3a5d572a7e4f0b4f9158cf) |
| `site_ref.kind` | [site_ref.kind](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-8386a31f5a93e0b0e727a410516070dd92782daf632d74c063695c36dde425ff) |
| `site_ref.name` | [site_ref.name](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-c5d794015d6206759febe26ead065e5e6f8cb1e5296c0935a7d43fca2e353f3e) |
| `site_ref.namespace` | [site_ref.namespace](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-922c5e8c6938b0fc6b9f0e934762f35189b1031b34cc60c9d15241dc335cc11d) |
| `site_ref.tenant` | [site_ref.tenant](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-4370551dd9aaced04d8c9e864b941e311c926a934db9d4a9ad8dfd41ee09c7e5) |
| `site_ref.uid` | [site_ref.uid](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-97b454d3eb06a050430d1acbe25012554902b6756fabe4e806dc209e21dd6d9e) |

<a id="canonical-e4fcca04b0f8be8e7da754a0545a47438d9384fd3b088dc7fd1a35c378d63d0b"></a>

## Next pages — Property reference / c7ae0427900f / 12

- [site_ref](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-cdd28936b4f4026b6c29bb0cf7632bd08291939f574e7c2bf62b1b7c37e060ff)
- [xcsh_cloud_elastic_ip](../data-sources/cloud_elastic_ip.md#canonical-2b2f30e89bfb277ce22595cc89aff6edb2de7c5d6f5d72c5d631bef5fa2ab292)

<a id="canonical-cdd28936b4f4026b6c29bb0cf7632bd08291939f574e7c2bf62b1b7c37e060ff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0fc98720ab33f463fc7a36ce3f6b25c58ae153ba8e0b5efbb3a10a6417f0f6a8"></a>

## site_ref — site_ref / 2dee5a2452a4 / 2

Breadcrumbs:

- [xcsh_cloud_elastic_ip](../data-sources/cloud_elastic_ip.md#canonical-2b2f30e89bfb277ce22595cc89aff6edb2de7c5d6f5d72c5d631bef5fa2ab292)
- [Property reference](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-ef15d26eed8ba06029d76ff217be5d95cee1296437e2bcdfc5add660013e8dcb)
- site_ref

<a id="canonical-9f847d2e4a6873ab02d547c72d9363ba1c358c8d2f3a5d572a7e4f0b4f9158cf"></a>

Type: `"list"`. Computed.

Site to which this cloud elastic IP object is attached.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-4b48c135ca002929393b9e1b9509a39ad20b68da15696c8fae80b5e1bf7b45c4"></a>

## Direct properties — site_ref / 2dee5a2452a4 / 3

<a id="canonical-8386a31f5a93e0b0e727a410516070dd92782daf632d74c063695c36dde425ff"></a>

<a id="canonical-4e37ce34280a6ebf51e2ca71b0f23eacd3053a49045e01f0fd69385a91ba5d44"></a>

## kind property — site_ref / 2dee5a2452a4 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-c5d794015d6206759febe26ead065e5e6f8cb1e5296c0935a7d43fca2e353f3e"></a>

<a id="canonical-c511eb7b48ac3bf76fce1083b290b2945ca920102d73e550fa0361d6a224038d"></a>

## name property — site_ref / 2dee5a2452a4 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-922c5e8c6938b0fc6b9f0e934762f35189b1031b34cc60c9d15241dc335cc11d"></a>

<a id="canonical-c1824f758ba33d8904ef98581d0fa2b6b403ba85cdd6afac282d103652b5d82b"></a>

## namespace property — site_ref / 2dee5a2452a4 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-4370551dd9aaced04d8c9e864b941e311c926a934db9d4a9ad8dfd41ee09c7e5"></a>

<a id="canonical-64424da369532f4a361c161c5014b7de8f22b73e54a7fb35c58d33db6533ac39"></a>

## tenant property — site_ref / 2dee5a2452a4 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-97b454d3eb06a050430d1acbe25012554902b6756fabe4e806dc209e21dd6d9e"></a>

<a id="canonical-12878c1a2d5ba1ac24041fdcf4080139d37743e9f0b94a24e649cd863b85d53c"></a>

## uid property — site_ref / 2dee5a2452a4 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-03cf1d692d5c84088df1f42162124d6110099066e45ec3a71ad339a7d12d0f9e"></a>

## Next pages — site_ref / 2dee5a2452a4 / 9

- [Property reference](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-ef15d26eed8ba06029d76ff217be5d95cee1296437e2bcdfc5add660013e8dcb)
- [xcsh_cloud_elastic_ip](../data-sources/cloud_elastic_ip.md#canonical-2b2f30e89bfb277ce22595cc89aff6edb2de7c5d6f5d72c5d631bef5fa2ab292)
