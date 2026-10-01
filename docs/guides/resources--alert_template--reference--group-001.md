---
page_title: "xcsh_alert_template reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_template reference."
---

# xcsh_alert_template reference

<a id="canonical-5fdfbe7749f50797d48aa600091b2ae0aae8cf488d15e19a3a5448fc91baf7ef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020bdd6d437e28b81095b0041061eed10a4babce7f484ea215232f9bb0cbd5a"></a>

## Property reference — Property reference / b719b8f8b03e / 2

Breadcrumbs:

- [xcsh_alert_template](../resources/alert_template.md#canonical-97fd669adea834d391f4ccd360fb87c16cdfc5afd855377f3e5665ed96c729d7)
- Property reference

<a id="canonical-5cb7230deef3ecaf98f42d859b16a032c0240667f9a39779055cfe738a98dfc7"></a>

## Direct properties — Property reference / b719b8f8b03e / 3

<a id="canonical-af31f398467b4cc820ca3296ca87ee85e5b2c5ebfd5c9388058e6543cb795d83"></a>

<a id="canonical-d807609d1b7a4cef218f0c94426563a7e22301919c055d8b9fb637e3d84fd4d2"></a>

## alert_message property — Property reference / b719b8f8b03e / 4

Type: `"string"`. Required.

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

<a id="canonical-23235837ea156f914ababd6e0a965539278478082167eda810c32f12e7a2e99d"></a>

<a id="canonical-3b06cf8c316eba15e67622ce85ccc58dd8c10a3045ff877de238ec70dac6e298"></a>

## alert_message_details property — Property reference / b719b8f8b03e / 5

Type: `"string"`. Required.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-4b3655db6a9c24748155dbe835fa4ca0fb744284148e56f0d0920138c3bc2474"></a>

<a id="canonical-f94f749e6ddca7a3cffe0e0a3543fb323410828e56ea5cbe6aafa55be401b943"></a>

## alert_name property — Property reference / b719b8f8b03e / 6

Type: `"string"`. Required.

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

<a id="canonical-ffde4dfca210c4f1ed24c3a06b92d79958b9657cac03a92163caada70fd61aed"></a>

<a id="canonical-a7613955ed4fd43699579b1e7ceec2974cdd7e8d62669f2e46c75f631c7b4483"></a>

## annotations property — Property reference / b719b8f8b03e / 7

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

<a id="canonical-c14f8c53bc8d5de1669d4bb46733c2ee4fed28dd46e3f700777861c5699c448c"></a>

<a id="canonical-5fb154e9a6786aab4d05407d87a06ff56d28436c76f6f9d8d86edf6683958018"></a>

## description property — Property reference / b719b8f8b03e / 8

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

<a id="canonical-407542d9259934ddf119dc33f98be853eb9755aa577285b3469a2f1ca9f5c430"></a>

<a id="canonical-b0d750b4b1802fbd0a34bc3fd6bb713c7a4055f3eb4be31f660e7aa65cccdfb2"></a>

## disable property — Property reference / b719b8f8b03e / 9

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

<a id="canonical-ec66de99fee4d7c0a0c4fc6e4a9414088f3f6ccc40267aec20cb17439b871a85"></a>

<a id="canonical-60e55aaa3074c0defc3c728f64ee8ff49cc6396824227e80f1d3092475236522"></a>

## id property — Property reference / b719b8f8b03e / 10

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-d75519610a6d2d2380dc3724b4edc031957d72084e54f1c1219c9697fade9ef3"></a>

<a id="canonical-c9c8ff1608bb2e48d2d653cc25e63bc0c24de922aa9768f126114061a41bfccd"></a>

## labels property — Property reference / b719b8f8b03e / 11

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

<a id="canonical-e9c227609ddc696d20e9cec0fde31142e595a257de9aad1ba52ab1926367166a"></a>

<a id="canonical-a302659b23812b4ad1fdb89d18c194365b5b4c8d734ade6bc7db4fb938b3f9d8"></a>

## name property — Property reference / b719b8f8b03e / 12

Type: `"string"`. Required.

Name of the Alert Template. Must be unique within the namespace.

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

<a id="canonical-7b08dbacb1242c5d4e76a42d7b64a69ff910240a867f57a293c410df3cae7aa8"></a>

<a id="canonical-97ef3f07f44c02fbc15b43366c30d79e3f938aa7f168e350585bda93deb2d1cc"></a>

## namespace property — Property reference / b719b8f8b03e / 13

Type: `"string"`. Required.

Namespace where the Alert Template is created.

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

<a id="canonical-b45cf69e30c3a3074bfbe15b7cda1f5d57eea332892b88d3dd636ac65625f27d"></a>

<a id="canonical-168f92b0fef728da924928c4a72efebbdc9dc2967ef0d91c6101e273c3422c6c"></a>

## severity property — Property reference / b719b8f8b03e / 14

Type: `"string"`. Optional, Computed.

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

- [timeouts](resources--alert_template--reference--group-001.md#canonical-b8ecfdad2d678acc1adbe892acd6f12391ea85818aa2f003ca83b6eca9ec49ae): complete subsection reference.

<a id="canonical-e250d8bf54b9bd988bf42e4e3899e3205ec792eafd770960fbbd138b114a942f"></a>

## All schema paths — Property reference / b719b8f8b03e / 15

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `alert_message` | [alert_message](resources--alert_template--reference--group-001.md#canonical-af31f398467b4cc820ca3296ca87ee85e5b2c5ebfd5c9388058e6543cb795d83) |
| `alert_message_details` | [alert_message_details](resources--alert_template--reference--group-001.md#canonical-23235837ea156f914ababd6e0a965539278478082167eda810c32f12e7a2e99d) |
| `alert_name` | [alert_name](resources--alert_template--reference--group-001.md#canonical-4b3655db6a9c24748155dbe835fa4ca0fb744284148e56f0d0920138c3bc2474) |
| `annotations` | [annotations](resources--alert_template--reference--group-001.md#canonical-ffde4dfca210c4f1ed24c3a06b92d79958b9657cac03a92163caada70fd61aed) |
| `description` | [description](resources--alert_template--reference--group-001.md#canonical-c14f8c53bc8d5de1669d4bb46733c2ee4fed28dd46e3f700777861c5699c448c) |
| `disable` | [disable](resources--alert_template--reference--group-001.md#canonical-407542d9259934ddf119dc33f98be853eb9755aa577285b3469a2f1ca9f5c430) |
| `id` | [id](resources--alert_template--reference--group-001.md#canonical-ec66de99fee4d7c0a0c4fc6e4a9414088f3f6ccc40267aec20cb17439b871a85) |
| `labels` | [labels](resources--alert_template--reference--group-001.md#canonical-d75519610a6d2d2380dc3724b4edc031957d72084e54f1c1219c9697fade9ef3) |
| `name` | [name](resources--alert_template--reference--group-001.md#canonical-e9c227609ddc696d20e9cec0fde31142e595a257de9aad1ba52ab1926367166a) |
| `namespace` | [namespace](resources--alert_template--reference--group-001.md#canonical-7b08dbacb1242c5d4e76a42d7b64a69ff910240a867f57a293c410df3cae7aa8) |
| `severity` | [severity](resources--alert_template--reference--group-001.md#canonical-b45cf69e30c3a3074bfbe15b7cda1f5d57eea332892b88d3dd636ac65625f27d) |
| `timeouts` | [timeouts](resources--alert_template--reference--group-001.md#canonical-8141ce10e0bdb00b0cf37a7f12e6164ba612ac150ebece62af497e70a7421b49) |
| `timeouts.create` | [timeouts.create](resources--alert_template--reference--group-001.md#canonical-2271dc90d027ccc5e688804a9485afce06b63266478700942b712feeaa212330) |
| `timeouts.delete` | [timeouts.delete](resources--alert_template--reference--group-001.md#canonical-35166c2682aba646e3d51ebb99bdb71e76801b3c4b21cc65df3ebaec285fb133) |
| `timeouts.read` | [timeouts.read](resources--alert_template--reference--group-001.md#canonical-5e004fcc1648ecec7cd1083bfff452fd1f28e957302739debaea712b70fe4828) |
| `timeouts.update` | [timeouts.update](resources--alert_template--reference--group-001.md#canonical-2bff74d0b938a0ca05f60f8cf75ab87ff3675cbb7497ba51b279e71035968ca7) |

<a id="canonical-1969be0c70b1e28fecda9e9eec35d4333301303236fe257f7f432cb54f9e7537"></a>

## Next pages — Property reference / b719b8f8b03e / 16

- [timeouts](resources--alert_template--reference--group-001.md#canonical-b8ecfdad2d678acc1adbe892acd6f12391ea85818aa2f003ca83b6eca9ec49ae)
- [xcsh_alert_template](../resources/alert_template.md#canonical-97fd669adea834d391f4ccd360fb87c16cdfc5afd855377f3e5665ed96c729d7)

<a id="canonical-b8ecfdad2d678acc1adbe892acd6f12391ea85818aa2f003ca83b6eca9ec49ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3731a90b78c6e7ea6fb248c04c92de299d643cc38a84313037105495627a3853"></a>

## timeouts — timeouts / 603676af4368 / 2

Breadcrumbs:

- [xcsh_alert_template](../resources/alert_template.md#canonical-97fd669adea834d391f4ccd360fb87c16cdfc5afd855377f3e5665ed96c729d7)
- [Property reference](resources--alert_template--reference--group-001.md#canonical-5fdfbe7749f50797d48aa600091b2ae0aae8cf488d15e19a3a5448fc91baf7ef)
- timeouts

<a id="canonical-8141ce10e0bdb00b0cf37a7f12e6164ba612ac150ebece62af497e70a7421b49"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-819eed4000bf2ab3aca270decd27af25468751baee3930f17ba0d292ea093bd8"></a>

## Direct properties — timeouts / 603676af4368 / 3

<a id="canonical-2271dc90d027ccc5e688804a9485afce06b63266478700942b712feeaa212330"></a>

<a id="canonical-adab292125dd911bbbbd5bc4c45e37d7efe3a20f5b048124709c417341589abc"></a>

## create property — timeouts / 603676af4368 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-35166c2682aba646e3d51ebb99bdb71e76801b3c4b21cc65df3ebaec285fb133"></a>

<a id="canonical-080f8d4f3050d70892dce04576df796e8596bbedbd090581fbb85d02dd2ed196"></a>

## delete property — timeouts / 603676af4368 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-5e004fcc1648ecec7cd1083bfff452fd1f28e957302739debaea712b70fe4828"></a>

<a id="canonical-10ee255891083c65985e5272c52450f31fd1a9ac63f28838b4df7879264ea29f"></a>

## read property — timeouts / 603676af4368 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2bff74d0b938a0ca05f60f8cf75ab87ff3675cbb7497ba51b279e71035968ca7"></a>

<a id="canonical-940b96833be4b744705b8fcea8c3767b4ddac02ce68843f3edb1e490d7ecd1e1"></a>

## update property — timeouts / 603676af4368 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-6aff81ff4665abb335edc3275b5df69a1aeeaf87ebd02eae425418006c36f8d8"></a>

## Next pages — timeouts / 603676af4368 / 8

- [Property reference](resources--alert_template--reference--group-001.md#canonical-5fdfbe7749f50797d48aa600091b2ae0aae8cf488d15e19a3a5448fc91baf7ef)
- [xcsh_alert_template](../resources/alert_template.md#canonical-97fd669adea834d391f4ccd360fb87c16cdfc5afd855377f3e5665ed96c729d7)
