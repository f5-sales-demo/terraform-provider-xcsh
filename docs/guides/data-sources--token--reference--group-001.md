---
page_title: "xcsh_token reference"
subcategory: "Identity"
description: "Complete grouped canonical reference for xcsh_token reference."
---

# xcsh_token reference

<a id="canonical-722430f94f35b0d2e9c17930c6fab80d7093f146d88e1f310120119163a9f05a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2bf515f701649c420ee94ffd132e57e6b1d9def369a03fa8847c33e5cd6d1bc3"></a>

## Property reference — Property reference / d5f34feda478 / 2

Breadcrumbs:

- [xcsh_token](../data-sources/token.md#canonical-807087d4571bee22ab642d9c81935b55fbfa12fbb563a3245dda111f3a3e04a2)
- Property reference

<a id="canonical-4a2842142e6b6dce692e04f13d570642bb51cc5e29c4ef4a90b56bff4177662c"></a>

## Direct properties — Property reference / d5f34feda478 / 3

<a id="canonical-31d73498ccf7ccb4236b58e45d6a232b437608cbf3ffb8ed0c9863ae540e42f1"></a>

<a id="canonical-a7bbe12456a4073547818359c9c4afde86871b8368bde213f5c28acf06f85c67"></a>

## annotations property — Property reference / d5f34feda478 / 4

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

<a id="canonical-895f76d970af5234489f95c9c173e3ec1c5ecc07a0296d83c4f7658afde16c96"></a>

<a id="canonical-25ac9ff38027f3b66bc1b1a94041c80cc37ae3683393c08f5fd08f7db7f77be4"></a>

## content property — Property reference / d5f34feda478 / 5

Type: `"string"`. Computed, Sensitive.

Server-issued JWT registration credential.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-sensitive": true
}
```

<a id="canonical-7e5f3803d5618199271b33ef0558232a2036016d8a4886ee55ebb6005573ad49"></a>

<a id="canonical-264e543f594bf9e09e74ba340c859359a14e8b8daac9504cd158f4d2537b6cd9"></a>

## description property — Property reference / d5f34feda478 / 6

Type: `"string"`. Computed.

Description of the Token.

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

<a id="canonical-8f9eacbc8700f758e69e0eafdb6128e94e6697e85df07cea5d2b7fe444828453"></a>

<a id="canonical-0c3f16e8b5f8aedad0bf666eda432d5bb367c4d5ce1b08d3d72654c80ebdddee"></a>

## id property — Property reference / d5f34feda478 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-d369921a404466468c6195463c2597c13bf790f272f685e9d4fb347ca4f78b38"></a>

<a id="canonical-a7614dbcb24d1be657be032fbe306d268e6b026b76f7716197a5ecd8926316ff"></a>

## labels property — Property reference / d5f34feda478 / 8

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

<a id="canonical-b5fb8fa57f33bdbf69c0c8defe0acc17b23fe50e1414b1ae63e82c4da6ce5f23"></a>

<a id="canonical-34152d485f27c0b3e5402a34bf623edfb74ccc88047a4db242c5cb20e403bed3"></a>

## name property — Property reference / d5f34feda478 / 9

Type: `"string"`. Required.

Name of the Token.

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

<a id="canonical-d60a70054be83a4469ba08b26f84d93becfa5d201a9cf13711c890933e7e9b34"></a>

<a id="canonical-660cecca74fb1059156cc71e38d6eec9bc5a32e8777732f2bb56f81e4866bf83"></a>

## namespace property — Property reference / d5f34feda478 / 10

Type: `"string"`. Optional, Computed.

Namespace where the Token exists.

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

<a id="canonical-db9c31e9a919a299b569e2cea6ad9d3fc4b992406ecf8a989d481cfc0be197b7"></a>

<a id="canonical-a49ce1d21200b9775876cc7b8c3698d99b3076bdb1284ace20ee997acd21b57f"></a>

## site_name property — Property reference / d5f34feda478 / 11

Type: `"string"`. Computed.

Secure Mesh Site v2 name bound into a JWT token.

<a id="canonical-e2daf0a26bf473c96ad424680cc430c53bc06431c989c436206f37ca2e87f466"></a>

<a id="canonical-c163158b5758a66c753ff914dac9b49a7e6c6bf43828cc368a1ae731980cb989"></a>

## type property — Property reference / d5f34feda478 / 12

Type: `"number"`. Computed.

\[Enum: 0|1\] Token type, where 0 is NORMAL and 1 is JWT. Possible values are \`0\`, \`1\`.

Upstream description:

Token type, where 0 is NORMAL and 1 is JWT.

Receipt-pinned upstream constraints:

```json
{
  "default": 0,
  "enum": [
    0,
    1
  ]
}
```

<a id="canonical-91db5af3ae23ab15f01587d149d26c7524394b5c36e09140876fd77d50b8abea"></a>

<a id="canonical-ae53c8ec5ac237ddaad91055f2fa64ecdb98cdd2eea187b2e6b7ca2730dd182d"></a>

## uid property — Property reference / d5f34feda478 / 13

Type: `"string"`. Computed, Sensitive.

Effective sensitive CE registration credential. NORMAL tokens use \`system\_metadata.uid\`; JWT
tokens use \`spec.content\`. This value is stored in plain text in the Terraform state file; ensure
your state file is properly secured.

<a id="canonical-c2b3acc5037689c681f58eec344c099e90f6f4d336428717a5621f3c185fde1d"></a>

## All schema paths — Property reference / d5f34feda478 / 14

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--token--reference--group-001.md#canonical-31d73498ccf7ccb4236b58e45d6a232b437608cbf3ffb8ed0c9863ae540e42f1) |
| `content` | [content](data-sources--token--reference--group-001.md#canonical-895f76d970af5234489f95c9c173e3ec1c5ecc07a0296d83c4f7658afde16c96) |
| `description` | [description](data-sources--token--reference--group-001.md#canonical-7e5f3803d5618199271b33ef0558232a2036016d8a4886ee55ebb6005573ad49) |
| `id` | [id](data-sources--token--reference--group-001.md#canonical-8f9eacbc8700f758e69e0eafdb6128e94e6697e85df07cea5d2b7fe444828453) |
| `labels` | [labels](data-sources--token--reference--group-001.md#canonical-d369921a404466468c6195463c2597c13bf790f272f685e9d4fb347ca4f78b38) |
| `name` | [name](data-sources--token--reference--group-001.md#canonical-b5fb8fa57f33bdbf69c0c8defe0acc17b23fe50e1414b1ae63e82c4da6ce5f23) |
| `namespace` | [namespace](data-sources--token--reference--group-001.md#canonical-d60a70054be83a4469ba08b26f84d93becfa5d201a9cf13711c890933e7e9b34) |
| `site_name` | [site_name](data-sources--token--reference--group-001.md#canonical-db9c31e9a919a299b569e2cea6ad9d3fc4b992406ecf8a989d481cfc0be197b7) |
| `type` | [type](data-sources--token--reference--group-001.md#canonical-e2daf0a26bf473c96ad424680cc430c53bc06431c989c436206f37ca2e87f466) |
| `uid` | [uid](data-sources--token--reference--group-001.md#canonical-91db5af3ae23ab15f01587d149d26c7524394b5c36e09140876fd77d50b8abea) |

<a id="canonical-38a566fdc5eb9bc3987c9611a8ee60bfd35af103d41aebd8949330c44c727407"></a>

## Next pages — Property reference / d5f34feda478 / 15

- [xcsh_token](../data-sources/token.md#canonical-807087d4571bee22ab642d9c81935b55fbfa12fbb563a3245dda111f3a3e04a2)
