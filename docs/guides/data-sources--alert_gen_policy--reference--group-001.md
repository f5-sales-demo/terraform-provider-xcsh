---
page_title: "xcsh_alert_gen_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_gen_policy reference."
---

# xcsh_alert_gen_policy reference

<a id="canonical-2e8fb0c00c43506e5e88f5573a46c7fc3c5d2904f33ee950e03580f94f30d516"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fde4aed4f3fc186a2e533ee7e3a579f043bca42291dd34a829c189c8c6805978"></a>

## Property reference — Property reference / d22d01ea01a6 / 2

Breadcrumbs:

- [xcsh_alert_gen_policy](../data-sources/alert_gen_policy.md#canonical-66af5b4db37114bd99ef2798c10e6b2faf4f99657029d35c516b10fefc2f1aa8)
- Property reference

<a id="canonical-d4502c92ba424716be6a77268848b7d01c093745a1e7659dda08edb92f4cb0bf"></a>

## Direct properties — Property reference / d22d01ea01a6 / 3

<a id="canonical-61e30b7edb754b94a5bbf1a1e7f2a0cf00a04c430b07fe6e8c9f87c774d099a3"></a>

<a id="canonical-898a856cec50bea29011e57872094e396a21a21df88072010308684d5f93b3e3"></a>

## alert_status property — Property reference / d22d01ea01a6 / 4

Type: `"string"`. Computed.

\[Enum: ALERT\_ACTIVE|ALERT\_INACTIVE\] Alert Status. List of alert statuses Active Inactive.
Possible values are \`ALERT\_ACTIVE\`, \`ALERT\_INACTIVE\`. Defaults to \`ALERT\_ACTIVE\`.

Upstream description:

List of alert statuses

Active Inactive.

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

<a id="canonical-51b0b72926d246687d476f7c9213c03637f06d2fd1581d92d33e35c450a5ec54"></a>

<a id="canonical-e4c5d46cc37f3fd82d769625915062be72a105157704878a3c6c14a048f8806e"></a>

## annotations property — Property reference / d22d01ea01a6 / 5

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

<a id="canonical-340dbb3e2873ffcd79c74e2ebf7d0e7520cc7f965eccf4266813163c5e63596e"></a>

<a id="canonical-07d532023f19aea401e306206bcf1d4620647b3ade1491418de6a669235410a7"></a>

## description property — Property reference / d22d01ea01a6 / 6

Type: `"string"`. Computed.

Description of the AlertGenPolicy.

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

- [details](data-sources--alert_gen_policy--reference--group-001.md#canonical-be0b58d4454d3e08847c1ab79670db07e543174bc4d995705451160f3bab3394): complete subsection reference.

<a id="canonical-d0a9294ab17b61ca66e0e56e1b08ae7210bc619f336333aaea3537ccfe17fa51"></a>

<a id="canonical-8ff7a805f358969f0aaa946631f27287ca6175a20eedc87f09b369f146381439"></a>

## id property — Property reference / d22d01ea01a6 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-15fd5b7494bc4e5c1a1027be1e0324d226e106d194b4f3ce787921e9e70c4086"></a>

<a id="canonical-e4064b2f27b1a7b1ba61bd6302b53575dabb1a8aa6a2c49b8b67a07f10fefe52"></a>

## labels property — Property reference / d22d01ea01a6 / 8

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

<a id="canonical-202652cc92a88d52c708ff2a8e1868a5af00068297b692f9586fa161ab2bf536"></a>

<a id="canonical-c5430baf4bd6618b33f66643cc9fd9361bfe9ef40a34758966d1ae695ad73bcc"></a>

## name property — Property reference / d22d01ea01a6 / 9

Type: `"string"`. Required.

Name of the AlertGenPolicy.

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

<a id="canonical-8acdff6a93ac7342a0367dda9ac9fa3e8507b265660250256848f14f2941e069"></a>

<a id="canonical-0af5511bbd0b4fdaa728202d9aeaa4790aeef8ac690d6415eab69deba106dbb6"></a>

## namespace property — Property reference / d22d01ea01a6 / 10

Type: `"string"`. Required.

Namespace where the AlertGenPolicy exists.

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

<a id="canonical-8e9f7cf876eddaebd64c6a7a11c20523a7e6ccadf0ce2302a409d2ce77465a40"></a>

## All schema paths — Property reference / d22d01ea01a6 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `alert_status` | [alert_status](data-sources--alert_gen_policy--reference--group-001.md#canonical-61e30b7edb754b94a5bbf1a1e7f2a0cf00a04c430b07fe6e8c9f87c774d099a3) |
| `annotations` | [annotations](data-sources--alert_gen_policy--reference--group-001.md#canonical-51b0b72926d246687d476f7c9213c03637f06d2fd1581d92d33e35c450a5ec54) |
| `description` | [description](data-sources--alert_gen_policy--reference--group-001.md#canonical-340dbb3e2873ffcd79c74e2ebf7d0e7520cc7f965eccf4266813163c5e63596e) |
| `details` | [details](data-sources--alert_gen_policy--reference--group-001.md#canonical-a04cfa8ca4fbae14acdef375e045e8987bd010bdac950f0afc85fdcee40b3e18) |
| `details.alert_message` | [details.alert_message](data-sources--alert_gen_policy--reference--group-001.md#canonical-c198971ef12eaad034fc2812a40be93f385fc9b8c0ef02e04e247c290374f0a5) |
| `details.alert_message_details` | [details.alert_message_details](data-sources--alert_gen_policy--reference--group-001.md#canonical-cf1d289c7a6a120b1e5bd76af6c5227eb10b5a19543b3718367c6e4c4c68abe2) |
| `details.alert_name` | [details.alert_name](data-sources--alert_gen_policy--reference--group-001.md#canonical-818894ced3fee4961b8bd5bddd45aa304da9943fbb2d5fa6b0f1b8429bfcf587) |
| `details.severity` | [details.severity](data-sources--alert_gen_policy--reference--group-001.md#canonical-a49be3d8027fda37749b3bd0c368112be8d1aecd28ffef77ffe7a35ba42c8f20) |
| `id` | [id](data-sources--alert_gen_policy--reference--group-001.md#canonical-d0a9294ab17b61ca66e0e56e1b08ae7210bc619f336333aaea3537ccfe17fa51) |
| `labels` | [labels](data-sources--alert_gen_policy--reference--group-001.md#canonical-15fd5b7494bc4e5c1a1027be1e0324d226e106d194b4f3ce787921e9e70c4086) |
| `name` | [name](data-sources--alert_gen_policy--reference--group-001.md#canonical-202652cc92a88d52c708ff2a8e1868a5af00068297b692f9586fa161ab2bf536) |
| `namespace` | [namespace](data-sources--alert_gen_policy--reference--group-001.md#canonical-8acdff6a93ac7342a0367dda9ac9fa3e8507b265660250256848f14f2941e069) |

<a id="canonical-e6f7ca1e04b3fa932e9e616bf78971b4c759087bb7b8a696f61922a5f6990eca"></a>

## Next pages — Property reference / d22d01ea01a6 / 12

- [details](data-sources--alert_gen_policy--reference--group-001.md#canonical-be0b58d4454d3e08847c1ab79670db07e543174bc4d995705451160f3bab3394)
- [xcsh_alert_gen_policy](../data-sources/alert_gen_policy.md#canonical-66af5b4db37114bd99ef2798c10e6b2faf4f99657029d35c516b10fefc2f1aa8)

<a id="canonical-be0b58d4454d3e08847c1ab79670db07e543174bc4d995705451160f3bab3394"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-564be600fef2883d212d29c61e98a1eb150a7091cadc1404fef8bf3ceceea576"></a>

## details — details / d740210ac56c / 2

Breadcrumbs:

- [xcsh_alert_gen_policy](../data-sources/alert_gen_policy.md#canonical-66af5b4db37114bd99ef2798c10e6b2faf4f99657029d35c516b10fefc2f1aa8)
- [Property reference](data-sources--alert_gen_policy--reference--group-001.md#canonical-2e8fb0c00c43506e5e88f5573a46c7fc3c5d2904f33ee950e03580f94f30d516)
- details

<a id="canonical-a04cfa8ca4fbae14acdef375e045e8987bd010bdac950f0afc85fdcee40b3e18"></a>

Type: `"single"`. Computed.

Notification Details. Notification Details.

Upstream description:

Notification Details.

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

<a id="canonical-074195d3cc1b5bf74addb93d9afe914905f03fe6538c5b0d905568df3ecdbd4b"></a>

## Direct properties — details / d740210ac56c / 3

<a id="canonical-c198971ef12eaad034fc2812a40be93f385fc9b8c0ef02e04e247c290374f0a5"></a>

<a id="canonical-52994b8bff77337dde5ab781e81f3197bcd91cb8238a0d4464fbb7f14f2cb779"></a>

## alert_message property — details / d740210ac56c / 4

Type: `"string"`. Computed.

Alert Message. Alert Message.

Upstream description:

Alert Message.

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

<a id="canonical-cf1d289c7a6a120b1e5bd76af6c5227eb10b5a19543b3718367c6e4c4c68abe2"></a>

<a id="canonical-7ba7a5f96d885688cf9c007e4d716a6b94323be0f4e76f074a4748847bcbd296"></a>

## alert_message_details property — details / d740210ac56c / 5

Type: `"string"`. Computed.

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

<a id="canonical-818894ced3fee4961b8bd5bddd45aa304da9943fbb2d5fa6b0f1b8429bfcf587"></a>

<a id="canonical-0fcede34c7b097674e09601ccc73c74161e530930b493a29da1d5bf196305474"></a>

## alert_name property — details / d740210ac56c / 6

Type: `"string"`. Computed.

Alert Name. Alert Name.

Upstream description:

Alert Name.

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

<a id="canonical-a49be3d8027fda37749b3bd0c368112be8d1aecd28ffef77ffe7a35ba42c8f20"></a>

<a id="canonical-f769c534f28fdbdccfaf506b66db184dce2dffb02bf86f492e6003ced3f88212"></a>

## severity property — details / d740210ac56c / 7

Type: `"string"`. Computed.

\[Enum: MINOR|MAJOR|CRITICAL\] List of alert severities Minor Major Critical. Possible values are
\`MINOR\`, \`MAJOR\`, \`CRITICAL\`. Defaults to \`MINOR\`.

Upstream description:

List of alert severities

Minor Major Critical.

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

<a id="canonical-339948c335004eb4b9f2c83f4eca96443be2652dc33d9690bf8429965e1a33c2"></a>

## Next pages — details / d740210ac56c / 8

- [Property reference](data-sources--alert_gen_policy--reference--group-001.md#canonical-2e8fb0c00c43506e5e88f5573a46c7fc3c5d2904f33ee950e03580f94f30d516)
- [xcsh_alert_gen_policy](../data-sources/alert_gen_policy.md#canonical-66af5b4db37114bd99ef2798c10e6b2faf4f99657029d35c516b10fefc2f1aa8)
