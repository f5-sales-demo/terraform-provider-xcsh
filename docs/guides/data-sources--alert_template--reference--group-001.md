---
page_title: "xcsh_alert_template reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_template reference."
---

# xcsh_alert_template reference

<a id="canonical-e6acd3bb08b31083263c203b17dd9b602b533a79339a12f5c11f80d073103560"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-05defc633c9f513ffd97de4bbf03238c1706e3b5e4cc8f8994b6d06a12733fbe"></a>

## Property reference — Property reference / 8db8649fb9ae / 2

Breadcrumbs:

- [xcsh_alert_template](../data-sources/alert_template.md#canonical-43748ee092eb77be7c417e05e7d416bc1aead9dc6699db467e2027d12e27cf88)
- Property reference

<a id="canonical-47d3d05628877ce25531c8b61c3f513122b9edf7868c90f11cadf8b5a0c397cd"></a>

## Direct properties — Property reference / 8db8649fb9ae / 3

<a id="canonical-916217ab560d290f41b113d139b5436e3de697fe24a49ae5b0dd9a54f4d0cb74"></a>

<a id="canonical-7afff9363ec7f032fb12ad94454be511620a5504e00f28768574633bc1beb01c"></a>

## alert_message property — Property reference / 8db8649fb9ae / 4

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

<a id="canonical-f822516bb37323f42d2b4639358548995603b131225f47e1dd73bb09fe11bacc"></a>

<a id="canonical-88cf530c92001f2bea2c08a957da12e9b09b8f0c020934acae22823e95e0c182"></a>

## alert_message_details property — Property reference / 8db8649fb9ae / 5

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-b21e12b2d575750081b1ee0129a78f4db3bbcc5b11a3ca3e4945eeae322f98b5"></a>

<a id="canonical-648132ae1f4b58d3bb3de9393f590800aef85046119b14f029a8d10b18c8df7a"></a>

## alert_name property — Property reference / 8db8649fb9ae / 6

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

<a id="canonical-82d5866f05c4b776f04b1a05a5b7359a6df41342454caf795ff5f3dc6bbf2145"></a>

<a id="canonical-bd673342d3d4da4a0d883f37b0e7f7442a8ea16342ad6e18214b1e5d552c3894"></a>

## annotations property — Property reference / 8db8649fb9ae / 7

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

<a id="canonical-09a147436fbb81d0722a927d03dc063c3a5adebd1b615f39a717c91aca5f9828"></a>

<a id="canonical-7371e7581bed44002c833786e0d8e63351dfa0015e14fd13ceb44889d823d9c9"></a>

## description property — Property reference / 8db8649fb9ae / 8

Type: `"string"`. Computed.

Description of the AlertTemplate.

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

<a id="canonical-fd34a725de39b6462110cd647577133b5b4264cd0cb5b83e4f843355c46dee59"></a>

<a id="canonical-a59ac3a606483bcd602ecd501abbb8f687ffbf1cd11dcafbc0bcc9c380c6e955"></a>

## id property — Property reference / 8db8649fb9ae / 9

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-36daf31c9e81c3d93ce4034780d961b0060fdf01b58871261235f65ab6b8a98a"></a>

<a id="canonical-7d9fb2d06bf3e798e335c2aa8e0456f3b9c3019a434168703e4a9b0e70352131"></a>

## labels property — Property reference / 8db8649fb9ae / 10

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

<a id="canonical-ca9277418ffa9c2f4360bea5b430d0eb8c0bf8492483bee981a31f8f5755f3d0"></a>

<a id="canonical-7c18de2d41b5ad71de883a7a7324d49ac84938330094283c8c7f7591d912b0ae"></a>

## name property — Property reference / 8db8649fb9ae / 11

Type: `"string"`. Required.

Name of the AlertTemplate.

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

<a id="canonical-b93894502759173024c5551abbf7be56ac74c90cdf1c0fab19a9202af1f65bb9"></a>

<a id="canonical-0bbbddc83d422e166edafa490d3f387a06bd27c0a3abb461913a77a67f7aa428"></a>

## namespace property — Property reference / 8db8649fb9ae / 12

Type: `"string"`. Required.

Namespace where the AlertTemplate exists.

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

<a id="canonical-c3acb95ded43007ac27d040c29bcb14a9e5315eb5bb1a87c30316dd0d998e8ad"></a>

<a id="canonical-9cf1d5056f6b99cc0f89fe7091ea249f33f581c5442e82a84907f3fb7708d213"></a>

## severity property — Property reference / 8db8649fb9ae / 13

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

<a id="canonical-6511931d8e30e587b4e871f648e69b66cbb9a3625334e7c692527abe6e10f246"></a>

## All schema paths — Property reference / 8db8649fb9ae / 14

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `alert_message` | [alert_message](data-sources--alert_template--reference--group-001.md#canonical-916217ab560d290f41b113d139b5436e3de697fe24a49ae5b0dd9a54f4d0cb74) |
| `alert_message_details` | [alert_message_details](data-sources--alert_template--reference--group-001.md#canonical-f822516bb37323f42d2b4639358548995603b131225f47e1dd73bb09fe11bacc) |
| `alert_name` | [alert_name](data-sources--alert_template--reference--group-001.md#canonical-b21e12b2d575750081b1ee0129a78f4db3bbcc5b11a3ca3e4945eeae322f98b5) |
| `annotations` | [annotations](data-sources--alert_template--reference--group-001.md#canonical-82d5866f05c4b776f04b1a05a5b7359a6df41342454caf795ff5f3dc6bbf2145) |
| `description` | [description](data-sources--alert_template--reference--group-001.md#canonical-09a147436fbb81d0722a927d03dc063c3a5adebd1b615f39a717c91aca5f9828) |
| `id` | [id](data-sources--alert_template--reference--group-001.md#canonical-fd34a725de39b6462110cd647577133b5b4264cd0cb5b83e4f843355c46dee59) |
| `labels` | [labels](data-sources--alert_template--reference--group-001.md#canonical-36daf31c9e81c3d93ce4034780d961b0060fdf01b58871261235f65ab6b8a98a) |
| `name` | [name](data-sources--alert_template--reference--group-001.md#canonical-ca9277418ffa9c2f4360bea5b430d0eb8c0bf8492483bee981a31f8f5755f3d0) |
| `namespace` | [namespace](data-sources--alert_template--reference--group-001.md#canonical-b93894502759173024c5551abbf7be56ac74c90cdf1c0fab19a9202af1f65bb9) |
| `severity` | [severity](data-sources--alert_template--reference--group-001.md#canonical-c3acb95ded43007ac27d040c29bcb14a9e5315eb5bb1a87c30316dd0d998e8ad) |

<a id="canonical-c90bff15b7394f0b443ab658d45431e77ffd531d3a43a21d4e34b2c3c7973b44"></a>

## Next pages — Property reference / 8db8649fb9ae / 15

- [xcsh_alert_template](../data-sources/alert_template.md#canonical-43748ee092eb77be7c417e05e7d416bc1aead9dc6699db467e2027d12e27cf88)
