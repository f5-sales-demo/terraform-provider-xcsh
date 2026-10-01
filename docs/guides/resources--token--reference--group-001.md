---
page_title: "xcsh_token reference"
subcategory: "Identity"
description: "Complete grouped canonical reference for xcsh_token reference."
---

# xcsh_token reference

<a id="canonical-05a33cba933cd1a2d79fe19b098f45465ae3f47fe73462661d6025e7ca495e14"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8a43bcd04d2dabca2f44e5a1bf10fd2ec2f91c3d9167956c7cc2a53eda305a3"></a>

## Property reference — Property reference / 196bdbc436be / 2

Breadcrumbs:

- [xcsh_token](../resources/token.md#canonical-4f14709ee0bf8c6a1f43d547a150dbc9ce3bfe820f4d53be2e91fa13ae4cb3c8)
- Property reference

<a id="canonical-9027b0e8433d21a8be71e0093d122d33f93a2f25c445a7e5a6fd0163ce611b8c"></a>

## Direct properties — Property reference / 196bdbc436be / 3

<a id="canonical-808ca8251c6a7dd688578db9f702b2db07b44aa12d8cee444a39068566e32030"></a>

<a id="canonical-790271d566224dae1823831076dcb4f102ce219abf1bf28505e628735764b1fb"></a>

## annotations property — Property reference / 196bdbc436be / 4

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

<a id="canonical-9890defe61ee40d217cb57e365df390718184fa05e2ea54fbba474bd46e3d719"></a>

<a id="canonical-7031db9b1f45d4b59a4cb26fb15c52d36026aa2baaa59c530896af2c1dd4bc67"></a>

## content property — Property reference / 196bdbc436be / 5

Type: `"string"`. Computed, Sensitive.

Server-issued JWT registration credential.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-sensitive": true
}
```

<a id="canonical-1be80b23259cb0a1cb4752ccc4768e50746cf823e8dc3cb276a96c4e2440d3c9"></a>

<a id="canonical-9d7c97740078d770f9ced9cd0adfbd5e2063d1c64a50d800488737a2a9e1a3df"></a>

## description property — Property reference / 196bdbc436be / 6

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

<a id="canonical-a056c1ed34e73af4a9fc33e0ed67d094db14755ee1ed8958e4225d1b729a9bb4"></a>

<a id="canonical-bb5abae2692dafd8153a328bf7347fb3f384d049c4d404cdc3f1bef58b62c822"></a>

## disable property — Property reference / 196bdbc436be / 7

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

<a id="canonical-c4a7224519677fd81c906f2050414b40d3183a74c156d667ffcb0167e1ec4ca8"></a>

<a id="canonical-dd2cc0d02f9d28673189ba8a620eb51ac730a9a84111ad85677386cb82b2242a"></a>

## id property — Property reference / 196bdbc436be / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-77df38d8cce88e030be5f775ff3f30c51ea8d6cde4302d8ce5740b49ecd4482d"></a>

<a id="canonical-4480748b0dd11919d474448d80b432e8ffa2ca6d8ac5f8bae833036b6e10d18b"></a>

## labels property — Property reference / 196bdbc436be / 9

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

<a id="canonical-bdef3f8577b31d31d1afe63f0628df0464dd38a1918b9e2bf21d8c0490696aef"></a>

<a id="canonical-3edc156eb5289ab479e850629a6ee8d3aa414cec2f447158119cb0980c8b980b"></a>

## name property — Property reference / 196bdbc436be / 10

Type: `"string"`. Required.

Name of the Token. Must be unique within the namespace.

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

<a id="canonical-548213cec6e771943453cd975c5e1514d3dc7ec1bbd914e13d0c8c971a48e073"></a>

<a id="canonical-34d9f1f39cfc8927775811de85b356b7e08a7b09e6e6d79f6b435acbc84ceabf"></a>

## namespace property — Property reference / 196bdbc436be / 11

Type: `"string"`. Optional, Computed.

Namespace for the Token. The F5 XC API restricts this resource to the system namespace; it defaults
to that value and may be omitted.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
Validators: []validator.String{
  validators.NamespaceValidator(),
  stringvalidator.OneOf("system"),
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

<a id="canonical-3082394fb79255718c374150c56ef80c4c5cc7a954041b302c5eaa05c717a633"></a>

<a id="canonical-2684afa5b29d52d1463777318ff46a9077832244006807b1862d777bd083f6ca"></a>

## site_name property — Property reference / 196bdbc436be / 12

Type: `"string"`. Optional, Computed.

Secure Mesh Site v2 name bound into a JWT token.

- [timeouts](resources--token--reference--group-001.md#canonical-d9b7631b3fe2f4454950053f0b6a4b35bf306548f5832979230bf24974ea5462): complete subsection reference.

<a id="canonical-adb6842d4acf9141b92f81a7320d894a531cb9c0868a2a38a5f50beb8b4d7e66"></a>

<a id="canonical-6d26248ba2877d73fd44459c7b00ede261a00db776628715e89f7d71d36ff275"></a>

## type property — Property reference / 196bdbc436be / 13

Type: `"number"`. Optional, Computed.

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

<a id="canonical-30e9f30d9d0518321f9b921eea9f0e3bc472a6084f4293955516007ed91b366a"></a>

<a id="canonical-00252fd01a846d169390e353d4e90e85a3e196e924ecaeb43626a89d371af9a7"></a>

## uid property — Property reference / 196bdbc436be / 14

Type: `"string"`. Computed, Sensitive.

Effective sensitive CE registration credential. NORMAL tokens use \`system\_metadata.uid\`; JWT
tokens use \`spec.content\`. This value is stored in plain text in the Terraform state file; ensure
your state file is properly secured.

<a id="canonical-b194cfa91ff4f5ec4c49db53290a0e8faa1b8b1bee8b4e1989a0d43bbebdd946"></a>

## All schema paths — Property reference / 196bdbc436be / 15

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--token--reference--group-001.md#canonical-808ca8251c6a7dd688578db9f702b2db07b44aa12d8cee444a39068566e32030) |
| `content` | [content](resources--token--reference--group-001.md#canonical-9890defe61ee40d217cb57e365df390718184fa05e2ea54fbba474bd46e3d719) |
| `description` | [description](resources--token--reference--group-001.md#canonical-1be80b23259cb0a1cb4752ccc4768e50746cf823e8dc3cb276a96c4e2440d3c9) |
| `disable` | [disable](resources--token--reference--group-001.md#canonical-a056c1ed34e73af4a9fc33e0ed67d094db14755ee1ed8958e4225d1b729a9bb4) |
| `id` | [id](resources--token--reference--group-001.md#canonical-c4a7224519677fd81c906f2050414b40d3183a74c156d667ffcb0167e1ec4ca8) |
| `labels` | [labels](resources--token--reference--group-001.md#canonical-77df38d8cce88e030be5f775ff3f30c51ea8d6cde4302d8ce5740b49ecd4482d) |
| `name` | [name](resources--token--reference--group-001.md#canonical-bdef3f8577b31d31d1afe63f0628df0464dd38a1918b9e2bf21d8c0490696aef) |
| `namespace` | [namespace](resources--token--reference--group-001.md#canonical-548213cec6e771943453cd975c5e1514d3dc7ec1bbd914e13d0c8c971a48e073) |
| `site_name` | [site_name](resources--token--reference--group-001.md#canonical-3082394fb79255718c374150c56ef80c4c5cc7a954041b302c5eaa05c717a633) |
| `timeouts` | [timeouts](resources--token--reference--group-001.md#canonical-86eafaac4fe076391f14350a839e70ef8adb435ca00fa7086a4b7c6c737bb1ba) |
| `timeouts.create` | [timeouts.create](resources--token--reference--group-001.md#canonical-2f6a2257ee4d0db409ab036585ecdb0d1acfc9366199cad63f623bb0514fc93d) |
| `timeouts.delete` | [timeouts.delete](resources--token--reference--group-001.md#canonical-a5b22a7219f7f22a6af53d4ef6bf228eaa830a03fa9184829df0c61f2535624d) |
| `timeouts.read` | [timeouts.read](resources--token--reference--group-001.md#canonical-99b4e9fd82f1869e7b5ec39b64d736d1c398e4e6325afae25d469e3743404eed) |
| `timeouts.update` | [timeouts.update](resources--token--reference--group-001.md#canonical-b672713feb531e8be0cc16caa4c6bc0ecbde70036d7a4464dd872e277a7c75f9) |
| `type` | [type](resources--token--reference--group-001.md#canonical-adb6842d4acf9141b92f81a7320d894a531cb9c0868a2a38a5f50beb8b4d7e66) |
| `uid` | [uid](resources--token--reference--group-001.md#canonical-30e9f30d9d0518321f9b921eea9f0e3bc472a6084f4293955516007ed91b366a) |

<a id="canonical-38a406c81610b2397b8f86ce8a839b35f2d94c92112c7a22237e25c7c097192b"></a>

## Next pages — Property reference / 196bdbc436be / 16

- [timeouts](resources--token--reference--group-001.md#canonical-d9b7631b3fe2f4454950053f0b6a4b35bf306548f5832979230bf24974ea5462)
- [xcsh_token](../resources/token.md#canonical-4f14709ee0bf8c6a1f43d547a150dbc9ce3bfe820f4d53be2e91fa13ae4cb3c8)

<a id="canonical-d9b7631b3fe2f4454950053f0b6a4b35bf306548f5832979230bf24974ea5462"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-64c7488a02032aef338a13d2d5b72b15ca269684a9e0868a0258fe529f72d1f7"></a>

## timeouts — timeouts / 141b49b32f53 / 2

Breadcrumbs:

- [xcsh_token](../resources/token.md#canonical-4f14709ee0bf8c6a1f43d547a150dbc9ce3bfe820f4d53be2e91fa13ae4cb3c8)
- [Property reference](resources--token--reference--group-001.md#canonical-05a33cba933cd1a2d79fe19b098f45465ae3f47fe73462661d6025e7ca495e14)
- timeouts

<a id="canonical-86eafaac4fe076391f14350a839e70ef8adb435ca00fa7086a4b7c6c737bb1ba"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-00bb4414f7aee478f1992a808adfc117884879715f3d5de266f25895e992369f"></a>

## Direct properties — timeouts / 141b49b32f53 / 3

<a id="canonical-2f6a2257ee4d0db409ab036585ecdb0d1acfc9366199cad63f623bb0514fc93d"></a>

<a id="canonical-f3f2e11be32af399e715ad77ab5dbf4890df29f491821587503c8bf7b98f6194"></a>

## create property — timeouts / 141b49b32f53 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-a5b22a7219f7f22a6af53d4ef6bf228eaa830a03fa9184829df0c61f2535624d"></a>

<a id="canonical-e71bbd105ed9010522ab4dd26995e569b4af1d543ae799a0a65b31ff0d9b8be1"></a>

## delete property — timeouts / 141b49b32f53 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-99b4e9fd82f1869e7b5ec39b64d736d1c398e4e6325afae25d469e3743404eed"></a>

<a id="canonical-6a39d128f1bf2171a8cadadccd41f3b30d6fd1d2e6a6fb49c0131f87899ec694"></a>

## read property — timeouts / 141b49b32f53 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-b672713feb531e8be0cc16caa4c6bc0ecbde70036d7a4464dd872e277a7c75f9"></a>

<a id="canonical-73bb821116d786beeb253c8b00640f0477322c030c6abb15f8c10584d51579b2"></a>

## update property — timeouts / 141b49b32f53 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-d50a9d3a1bdf61d30fd8a89b071436d3a3eb2a36d7e6a3df57f009056300a2f7"></a>

## Next pages — timeouts / 141b49b32f53 / 8

- [Property reference](resources--token--reference--group-001.md#canonical-05a33cba933cd1a2d79fe19b098f45465ae3f47fe73462661d6025e7ca495e14)
- [xcsh_token](../resources/token.md#canonical-4f14709ee0bf8c6a1f43d547a150dbc9ce3bfe820f4d53be2e91fa13ae4cb3c8)
