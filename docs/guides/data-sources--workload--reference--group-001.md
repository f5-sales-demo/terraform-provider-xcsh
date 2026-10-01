---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ec8a4ff55e7e7a83b52fa2e26486e2779a36b0cfbd317ba919bdc95ccd8b8817"></a>

## Property reference — Property reference / dacf77bc523e / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- Property reference

<a id="canonical-cf35b06ea1ea0d1465306ee6ad2ab1312730ffac2aeea938074c2a497c5eecc6"></a>

## Direct properties — Property reference / dacf77bc523e / 3

<a id="canonical-fd479db49ef7bedcfa1ecd726b0bbbf8d948e4f2d7faccc51633c8159340913a"></a>

<a id="canonical-d7fdf77fc159627591fe2f15df16792e8537d77e86dd3334c10936b6a7964954"></a>

## annotations property — Property reference / dacf77bc523e / 4

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

<a id="canonical-c5eee35f5bb0610e55891b3d299ef5ed199a97c529fedd098d0106a77c69c46e"></a>

<a id="canonical-cf1f36fdd77e6642e333b2ade9f3eb3577ae28d8bd8501d2370eeae4777b33e8"></a>

## description property — Property reference / dacf77bc523e / 5

Type: `"string"`. Computed.

Description of the Workload.

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

<a id="canonical-50dd48a811ec3ebb519eb826260f24605f1518e0277f61e83360d96b3bae552e"></a>

<a id="canonical-9e758e9ec9a438ceea3237190a5b96e4d17c1eb9b92b6c36f1ba8c966afe104d"></a>

## id property — Property reference / dacf77bc523e / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [job](data-sources--workload--reference--group-004.md#canonical-7350f554fe4d86d88468d1c9f06e5a182c8c9bdbd89f5cca6ed4489c26853e5d): complete subsection reference.

<a id="canonical-c4581ae32ce86d824a68ab3f551918441553427090886d4388fdbd1633507da9"></a>

<a id="canonical-e55a25540662fde1eecb527fec60bc5cfc5beb7da8ce32ea834d3ef3d4ee3673"></a>

## labels property — Property reference / dacf77bc523e / 7

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

<a id="canonical-56b09ee03bbed25ec29e62f230ffc73add3a13a4c5bca39814efc3db2f74541f"></a>

<a id="canonical-35baab57680bb86f7f6d57f1de7cfd25d6a84d8e8d57e60d84d4f2a0adf8fa70"></a>

## name property — Property reference / dacf77bc523e / 8

Type: `"string"`. Required.

Name of the Workload.

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

<a id="canonical-ccd4e84f0f6beb90a965ab48f51dad4ade0c624c930563228d8dbd19b5314a08"></a>

<a id="canonical-e1759c20a908a17d5d194ff07889c90b952c869646de1b7602cf3e1ac33529b3"></a>

## namespace property — Property reference / dacf77bc523e / 9

Type: `"string"`. Required.

Namespace where the Workload exists.

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

- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025): complete subsection reference.

- [simple_service](data-sources--workload--reference--group-016.md#canonical-5ff01a9f0b1ff1fcfc1fa389f29f4c0fb638d848ef363c6d333d0ba02b54fa04): complete subsection reference.

- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497): complete subsection reference.
