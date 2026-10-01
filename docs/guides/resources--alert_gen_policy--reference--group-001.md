---
page_title: "xcsh_alert_gen_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_gen_policy reference."
---

# xcsh_alert_gen_policy reference

<a id="canonical-17ec4655f437510f18a5ab3c4e0b2710c50f8dd33e3adde8ef55402086569933"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-302072d3b1cfaa6237ef9fcdbae298544d5c237e09c254dc4a1136dc54e52aca"></a>

## Property reference — Property reference / 31e9a48e1094 / 2

Breadcrumbs:

- [xcsh_alert_gen_policy](../resources/alert_gen_policy.md#canonical-a1aae63c73737fe28891b8dddcd40970e1ba1214a48d75c4ae655469cca81541)
- Property reference

<a id="canonical-2c74ea971802a82d2479fd13f6227b1f0fa70419572d320bbdbad7fc76b6bd0e"></a>

## Direct properties — Property reference / 31e9a48e1094 / 3

<a id="canonical-0232ad96b0f85482590377498678cde5aa6643c5b7ef8c1e9eea0ef6413843e8"></a>

<a id="canonical-bba755a72ca99b008b41fded44c26edea5da7ea3eff7c88e1472d0105cf50514"></a>

## alert_status property — Property reference / 31e9a48e1094 / 4

Type: `"string"`. Optional, Computed.

\[Enum: ALERT\_ACTIVE|ALERT\_INACTIVE\] Alert Status. List of alert statuses Active Inactive.
Possible values are \`ALERT\_ACTIVE\`, \`ALERT\_INACTIVE\`. Defaults to \`ALERT\_ACTIVE\`.

Upstream description:

List of alert statuses

Active Inactive.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ALERT_ACTIVE",
    "ALERT_INACTIVE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ALERT_ACTIVE",
  "enum": [
    "ALERT_ACTIVE",
    "ALERT_INACTIVE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-e4f95209c5ff77fffcb7378b0c351b8d220099b84294ad02d5b0b94642e59ffe"></a>

<a id="canonical-5ac98c3beb96d4d7aa27bf5aec2957f6b4ca1189c000ee15106da760a2a71ac1"></a>

## annotations property — Property reference / 31e9a48e1094 / 5

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

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

<a id="canonical-177b59eb7de1c61d293432522f84378b413a0da8ce91e3f3317ad1d663631ec0"></a>

<a id="canonical-1889518c5a48b1eeab7ff85c959d4df7575b500d88997502b5fe079220a14bd0"></a>

## description property — Property reference / 31e9a48e1094 / 6

Type: `"string"`. Optional.

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

- [details](resources--alert_gen_policy--reference--group-001.md#canonical-389c81d944eafb66b0fb8705d9f33c8401bdeeacfbbe52b4d7f1a679917223b8): complete subsection reference.

<a id="canonical-dc7a124376add9513f8580d6b0e60e563e8b2494d43c6b054e7bf3eafc4e01fb"></a>

<a id="canonical-1c6862a570a66cad1047ce46a5dfa7b284620ae236e6d0a72be477d6689a2ae2"></a>

## disable property — Property reference / 31e9a48e1094 / 7

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

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

<a id="canonical-c83510f1ddf486a199aca726f9dce44455af5b36e75a673115c4dfd7d36706ea"></a>

<a id="canonical-5958b13e4f883c05f0d78a023eec0928873dadbe611cb3fa962ddd62a49f0833"></a>

## id property — Property reference / 31e9a48e1094 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0447859da328b30210871a8f29f4c01af16f0f1a2a5e13bd35d0051b67abeeb3"></a>

<a id="canonical-2fdd8ded8703b715c1beaf93af46b5e75e3079abf36cee132e3a3ea6ed95403b"></a>

## labels property — Property reference / 31e9a48e1094 / 9

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

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

<a id="canonical-aa1ff59cdaa8e9d2eb9c476ba1e2f328346382891241cbda8fb9e47d189ece8e"></a>

<a id="canonical-e829c92b60a8e9245fce67ff83a3c2378f3706a120a5f8603e3ac8895d53a8d5"></a>

## name property — Property reference / 31e9a48e1094 / 10

Type: `"string"`. Required.

Name of the Alert Gen Policy. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

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

<a id="canonical-8bd330f281cb10f7d90332671cd47e97eba7d4b8de2731c6014df312ad76d8d9"></a>

<a id="canonical-77f14cca446f5a3b1ef1dd2ac7a58f1e6167b08a4c7139081bd7dd48251098a0"></a>

## namespace property — Property reference / 31e9a48e1094 / 11

Type: `"string"`. Required.

Namespace where the Alert Gen Policy is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

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

- [timeouts](resources--alert_gen_policy--reference--group-001.md#canonical-ea139050cc03b9f150b8bdaac01befa332938c9cb5f31af2007f481149b74d20): complete subsection reference.

<a id="canonical-ae8832ec7678fd20ff6e5900c588111446331192305e7486d20147e5d8411ab1"></a>

## All schema paths — Property reference / 31e9a48e1094 / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `alert_status` | [alert_status](resources--alert_gen_policy--reference--group-001.md#canonical-0232ad96b0f85482590377498678cde5aa6643c5b7ef8c1e9eea0ef6413843e8) |
| `annotations` | [annotations](resources--alert_gen_policy--reference--group-001.md#canonical-e4f95209c5ff77fffcb7378b0c351b8d220099b84294ad02d5b0b94642e59ffe) |
| `description` | [description](resources--alert_gen_policy--reference--group-001.md#canonical-177b59eb7de1c61d293432522f84378b413a0da8ce91e3f3317ad1d663631ec0) |
| `details` | [details](resources--alert_gen_policy--reference--group-001.md#canonical-86bebd9bd3d837b43d5f85106e540c661bfb832c42ea4e1cd2cb07c4955b0b2c) |
| `details.alert_message` | [details.alert_message](resources--alert_gen_policy--reference--group-001.md#canonical-10ce7d675151aadc50c9cfc03e4053d2823962444fd1b9007684dfcd530f6ce1) |
| `details.alert_message_details` | [details.alert_message_details](resources--alert_gen_policy--reference--group-001.md#canonical-5bcc7aabd5dfe4ca523a4b333fde3311dea4adb6fcd51beafc77872edac5dd5c) |
| `details.alert_name` | [details.alert_name](resources--alert_gen_policy--reference--group-001.md#canonical-5785876cd54ab35b2fe00d1aa787856066c2ba0367d43577619e61cb6e0f34c0) |
| `details.severity` | [details.severity](resources--alert_gen_policy--reference--group-001.md#canonical-1638eec9c9cd2e935de0ae1af837b70f6943d7573a4bff996c8a9034a7cba10d) |
| `disable` | [disable](resources--alert_gen_policy--reference--group-001.md#canonical-dc7a124376add9513f8580d6b0e60e563e8b2494d43c6b054e7bf3eafc4e01fb) |
| `id` | [id](resources--alert_gen_policy--reference--group-001.md#canonical-c83510f1ddf486a199aca726f9dce44455af5b36e75a673115c4dfd7d36706ea) |
| `labels` | [labels](resources--alert_gen_policy--reference--group-001.md#canonical-0447859da328b30210871a8f29f4c01af16f0f1a2a5e13bd35d0051b67abeeb3) |
| `name` | [name](resources--alert_gen_policy--reference--group-001.md#canonical-aa1ff59cdaa8e9d2eb9c476ba1e2f328346382891241cbda8fb9e47d189ece8e) |
| `namespace` | [namespace](resources--alert_gen_policy--reference--group-001.md#canonical-8bd330f281cb10f7d90332671cd47e97eba7d4b8de2731c6014df312ad76d8d9) |
| `timeouts` | [timeouts](resources--alert_gen_policy--reference--group-001.md#canonical-008d519c83651387fb98041e8bf2352f7f8c5d55c30bb2d9064583a492b07d71) |
| `timeouts.create` | [timeouts.create](resources--alert_gen_policy--reference--group-001.md#canonical-85b7e520138cc08cb62b33e56c729c051dbaa5caeca865fb4bec9a2f039ded27) |
| `timeouts.delete` | [timeouts.delete](resources--alert_gen_policy--reference--group-001.md#canonical-1aec209e5419c714c82f29c51541a488db9f1583abd93e71d3605a55c676372e) |
| `timeouts.read` | [timeouts.read](resources--alert_gen_policy--reference--group-001.md#canonical-019f82c8f8e1a93d11887b79946152bef03b47851279411ddb34022c97fc0ea7) |
| `timeouts.update` | [timeouts.update](resources--alert_gen_policy--reference--group-001.md#canonical-0db22a1a7bc297675c48f31c12bf5394ddebea28ecdd8d5f5cdbf85ff009b8e4) |

<a id="canonical-8a33b586affb3cfa021049e6388328836a92d93093273f8e59c21b52445a3ada"></a>

## Next pages — Property reference / 31e9a48e1094 / 13

- [details](resources--alert_gen_policy--reference--group-001.md#canonical-389c81d944eafb66b0fb8705d9f33c8401bdeeacfbbe52b4d7f1a679917223b8)
- [timeouts](resources--alert_gen_policy--reference--group-001.md#canonical-ea139050cc03b9f150b8bdaac01befa332938c9cb5f31af2007f481149b74d20)
- [xcsh_alert_gen_policy](../resources/alert_gen_policy.md#canonical-a1aae63c73737fe28891b8dddcd40970e1ba1214a48d75c4ae655469cca81541)

<a id="canonical-389c81d944eafb66b0fb8705d9f33c8401bdeeacfbbe52b4d7f1a679917223b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4cc20cd9af353dd7e31d86435213bf8da611c608a2d23f5c46a5e6023bf77e15"></a>

## details — details / 393b0aeb5ab8 / 2

Breadcrumbs:

- [xcsh_alert_gen_policy](../resources/alert_gen_policy.md#canonical-a1aae63c73737fe28891b8dddcd40970e1ba1214a48d75c4ae655469cca81541)
- [Property reference](resources--alert_gen_policy--reference--group-001.md#canonical-17ec4655f437510f18a5ab3c4e0b2710c50f8dd33e3adde8ef55402086569933)
- details

<a id="canonical-86bebd9bd3d837b43d5f85106e540c661bfb832c42ea4e1cd2cb07c4955b0b2c"></a>

Type: `"object"`. single nested block, Optional.

Notification Details. Notification Details.

Upstream description:

Notification Details.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("alert_message",
    "alert_message_details",
    "alert_name")}
```

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

Terraform syntax:

```terraform
details {
  # Configure direct properties listed below.
}
```

<a id="canonical-b9305d720a2637cd3566254a162be2d3265e3ef2c57890e13da79c6acf20a71b"></a>

## Direct properties — details / 393b0aeb5ab8 / 3

<a id="canonical-10ce7d675151aadc50c9cfc03e4053d2823962444fd1b9007684dfcd530f6ce1"></a>

<a id="canonical-866b0f7bb86b629c5702e229eacc72d5bd37ce48eece89d50ff3686916d7f154"></a>

## alert_message property — details / 393b0aeb5ab8 / 4

Type: `"string"`. Optional.

Alert Message. Alert Message.

Upstream description:

Alert Message.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512"
  }
}
```

<a id="canonical-5bcc7aabd5dfe4ca523a4b333fde3311dea4adb6fcd51beafc77872edac5dd5c"></a>

<a id="canonical-ea2a441e2dd24d8cdbf91940ff63d0a4de1e553372e6a8ca174760ef151254e3"></a>

## alert_message_details property — details / 393b0aeb5ab8 / 5

Type: `"string"`. Optional.

Alert Message Details. Detailed message of the alert.

Upstream description:

Detailed message of the alert.

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
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-5785876cd54ab35b2fe00d1aa787856066c2ba0367d43577619e61cb6e0f34c0"></a>

<a id="canonical-dd810f7ee9e3c5bc7d0a4a67d12ac35a8afa65a81e79cd4292c51ba7734d350c"></a>

## alert_name property — details / 393b0aeb5ab8 / 6

Type: `"string"`. Optional.

Alert Name. Alert Name.

Upstream description:

Alert Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 16,
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
    "ves.io.schema.rules.string.max_len": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "16"
  }
}
```

<a id="canonical-1638eec9c9cd2e935de0ae1af837b70f6943d7573a4bff996c8a9034a7cba10d"></a>

<a id="canonical-5b341f228a98f5a4209950761369f117797f767399ddc6da53e864fa10b89535"></a>

## severity property — details / 393b0aeb5ab8 / 7

Type: `"string"`. Optional.

\[Enum: MINOR|MAJOR|CRITICAL\] List of alert severities Minor Major Critical. Possible values are
\`MINOR\`, \`MAJOR\`, \`CRITICAL\`. Defaults to \`MINOR\`.

Upstream description:

List of alert severities

Minor Major Critical.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("MINOR",
    "MAJOR",
    "CRITICAL"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "MINOR",
  "enum": [
    "MINOR",
    "MAJOR",
    "CRITICAL"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-61c5c0df8d1dc3709c7ec3e8a7c46547e9d21cb9e855766d6e7c60593db0d844"></a>

## Next pages — details / 393b0aeb5ab8 / 8

- [Property reference](resources--alert_gen_policy--reference--group-001.md#canonical-17ec4655f437510f18a5ab3c4e0b2710c50f8dd33e3adde8ef55402086569933)
- [xcsh_alert_gen_policy](../resources/alert_gen_policy.md#canonical-a1aae63c73737fe28891b8dddcd40970e1ba1214a48d75c4ae655469cca81541)

<a id="canonical-ea139050cc03b9f150b8bdaac01befa332938c9cb5f31af2007f481149b74d20"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e568dd761ac49ade4a5e7e379a83af74c30f41fbd1a94dc05737fdfc529d4ea5"></a>

## timeouts — timeouts / 9206181daf4d / 2

Breadcrumbs:

- [xcsh_alert_gen_policy](../resources/alert_gen_policy.md#canonical-a1aae63c73737fe28891b8dddcd40970e1ba1214a48d75c4ae655469cca81541)
- [Property reference](resources--alert_gen_policy--reference--group-001.md#canonical-17ec4655f437510f18a5ab3c4e0b2710c50f8dd33e3adde8ef55402086569933)
- timeouts

<a id="canonical-008d519c83651387fb98041e8bf2352f7f8c5d55c30bb2d9064583a492b07d71"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-e3d52bfbf206e204ded82b47fb65a8e30c44ba100a5fd6d159a0e90949596fe3"></a>

## Direct properties — timeouts / 9206181daf4d / 3

<a id="canonical-85b7e520138cc08cb62b33e56c729c051dbaa5caeca865fb4bec9a2f039ded27"></a>

<a id="canonical-7710d50527865c39ea061f7799b583979c89654e288baee23f6ef0155d996adc"></a>

## create property — timeouts / 9206181daf4d / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1aec209e5419c714c82f29c51541a488db9f1583abd93e71d3605a55c676372e"></a>

<a id="canonical-71e08044606b95a6f5a871d75b912f1e78ebd51c7dc9dab26cd94ef12a5dcf72"></a>

## delete property — timeouts / 9206181daf4d / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-019f82c8f8e1a93d11887b79946152bef03b47851279411ddb34022c97fc0ea7"></a>

<a id="canonical-1dd6c68d3ca0e00acff11013cc8a1887b620ec686967ea4e1b5aefdfb554b9be"></a>

## read property — timeouts / 9206181daf4d / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0db22a1a7bc297675c48f31c12bf5394ddebea28ecdd8d5f5cdbf85ff009b8e4"></a>

<a id="canonical-fc92a8b38051cb55eb48451fa27627ee4b53ad332e7b0a8acf0360ce0d4a6d31"></a>

## update property — timeouts / 9206181daf4d / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-66209e2789379edc96909e0f8765a588ea7ce696d723ea0ea8ce6a90373021e3"></a>

## Next pages — timeouts / 9206181daf4d / 8

- [Property reference](resources--alert_gen_policy--reference--group-001.md#canonical-17ec4655f437510f18a5ab3c4e0b2710c50f8dd33e3adde8ef55402086569933)
- [xcsh_alert_gen_policy](../resources/alert_gen_policy.md#canonical-a1aae63c73737fe28891b8dddcd40970e1ba1214a48d75c4ae655469cca81541)
