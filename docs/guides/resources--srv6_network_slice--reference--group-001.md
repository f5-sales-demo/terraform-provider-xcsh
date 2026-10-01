---
page_title: "xcsh_srv6_network_slice reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_srv6_network_slice reference."
---

# xcsh_srv6_network_slice reference

<a id="canonical-a9921f7005bdd8167184c5a6a56fca49fff6fc9a386f4d22e032d049c63fc870"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb27bc8e27e65e5d74ed79788a33017120d8b1cc7a32df9313857adce6489719"></a>

## Property reference — Property reference / 6ee38da4b4c6 / 2

Breadcrumbs:

- [xcsh_srv6_network_slice](../resources/srv6_network_slice.md#canonical-431435cc9f00e58094c1dd963b8aa9cec4251ea8dcbf3313b046ca1921402d3e)
- Property reference

<a id="canonical-1ec81b01faae18a817fe06124ca53bfc6bed36051c41fea66636ca191c7f316b"></a>

## Direct properties — Property reference / 6ee38da4b4c6 / 3

<a id="canonical-eb6f1cb8a849ab0ca6377699ae9494d9461351ec9e6eb787fd08a6efa557d10a"></a>

<a id="canonical-28e679949c94b8a15cba95baf78471de18df70bd3ef2961e71fcbe86fdb08dbd"></a>

## annotations property — Property reference / 6ee38da4b4c6 / 4

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

<a id="canonical-ba7f4e4ee691cd21843eefa2e212014b9aace4f5796d10d8ed39447c38b24b46"></a>

<a id="canonical-a5475f8119c77491b84ef28d0c7c3bdc804f73c677f3a150aae513e99a31ef7f"></a>

## connect_to_access_networks property — Property reference / 6ee38da4b4c6 / 5

Type: `"bool"`. Optional, Computed.

Connect all SRv6 Virtual Networks in this slice to their corresponding access networks by importing
route targets specified in the virtual network.

Upstream description:

Connect all SRv6 Virtual Networks in this slice to their corresponding access networks by importing
route targets specified in the virtual network.

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

<a id="canonical-3f123c26a55a6f32b9675a1aa011d6afca801e466fb032be64d898653f5947d1"></a>

<a id="canonical-c92dbfdffc63d744e381286ca61004ce43b25f20d207f12d239e4d26f7eed0c5"></a>

## connect_to_enterprise_networks property — Property reference / 6ee38da4b4c6 / 6

Type: `"bool"`. Optional, Computed.

Connect all SRv6 Virtual Networks in this slice to their corresponding enterprise networks by
importing route targets specified in the virtual network.

Upstream description:

Connect all SRv6 Virtual Networks in this slice to their corresponding enterprise networks by
importing route targets specified in the virtual network.

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

<a id="canonical-4b20dd12970071b6d1e9a09c6ccba976c9db5059ed461c4ff1f7ed64387742a4"></a>

<a id="canonical-b3f60d09135bed38360c729d3a64b8453260a5d4e13959741f1bba7314110811"></a>

## connect_to_internet property — Property reference / 6ee38da4b4c6 / 7

Type: `"bool"`. Optional, Computed.

Connect all SRv6 Virtual Networks in this slice to the Internet by importing route targets specified
in the virtual network.

Upstream description:

Connect all SRv6 Virtual Networks in this slice to the Internet by importing route targets specified
in the virtual network.

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

<a id="canonical-7398a463cc259f8b890bdecc5dd7d081706b10356d11af0a7d9782d883c4697a"></a>

<a id="canonical-68f760a68724c15b18abbaef1c9aa41856a25e7f7ac6889e5babb714bda2da39"></a>

## description property — Property reference / 6ee38da4b4c6 / 8

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

<a id="canonical-bce9e2ff0e9be7376248de4d4923aab72811670bbc7182c62cb920b576953e5a"></a>

<a id="canonical-cfb39006dd776425dc35d478332f5c499b62e01f0ec69c00c5f9c208dfb2db54"></a>

## disable property — Property reference / 6ee38da4b4c6 / 9

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

<a id="canonical-ef261ac59a2b35908cf0921515fe3c2d0b59918f4ef1fbac92ff8b0e48611746"></a>

<a id="canonical-eb7eb7bd5d3e48c6ab6e8af6cd1251fb616ef13e99f5e1160c34462598622b96"></a>

## id property — Property reference / 6ee38da4b4c6 / 10

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0391adbfef8c5a6ae08f4be13252b8402e1c25157a9e16e0379df8fab741b348"></a>

<a id="canonical-2bb48152d870d9c3ef29351fc069edec90cf876f98ae5ffd721860a19425b5ad"></a>

## labels property — Property reference / 6ee38da4b4c6 / 11

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

<a id="canonical-f19432212dec036539b172bc70f61cbffec4b4c5441a9130d666ec6eb8009ad7"></a>

<a id="canonical-ac43df64be13fd78fbb9f0eaf1adee5d48407f96eec9c438a93b6fff3d383d60"></a>

## name property — Property reference / 6ee38da4b4c6 / 12

Type: `"string"`. Required.

Name of the Srv6 Network Slice. Must be unique within the namespace.

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

<a id="canonical-f92c150e05bba48ff94081a1e1b81001683745a0c828520d3362d2a2dbb73645"></a>

<a id="canonical-ec6d5faa4d9b88159d09ae95ada02810e4cc0c6dbaec2c77c34af0dc1f9f154f"></a>

## namespace property — Property reference / 6ee38da4b4c6 / 13

Type: `"string"`. Optional, Computed.

Namespace for the Srv6 Network Slice. The F5 XC API restricts this resource to the system namespace;
it defaults to that value and may be omitted.

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

<a id="canonical-a34f3258c940279ce5501e7d8d38c7ae447b92efa469aa58799b95df8dfd72d3"></a>

<a id="canonical-30995fa416358563ae9bd84b46d30c228a6c0cd5f2fa83a8cadfddd1e1a8bded"></a>

## sid_prefixes property — Property reference / 6ee38da4b4c6 / 14

Type: `["list", "string"]`. Required.

SID Locator from the prefix is allocated automatically for each node in each site.

Upstream description:

A SID Locator from the prefix is allocated automatically for each node in each site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.max_ip_prefix_length": "32",
    "ves.io.schema.rules.repeated.items.string.min_ip_prefix_length": "32",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.max_ip_prefix_length": "32",
    "ves.io.schema.rules.repeated.items.string.min_ip_prefix_length": "32",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [timeouts](resources--srv6_network_slice--reference--group-001.md#canonical-bb84fcc743229e13d3b165a9c3c89d3253afa395a8d97cfc70dee42e9bbffb97): complete subsection reference.

<a id="canonical-484084493b6365e47a1cb6ad0b7313db7e2016d826e026dd414a0d9c433c8644"></a>

## All schema paths — Property reference / 6ee38da4b4c6 / 15

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--srv6_network_slice--reference--group-001.md#canonical-eb6f1cb8a849ab0ca6377699ae9494d9461351ec9e6eb787fd08a6efa557d10a) |
| `connect_to_access_networks` | [connect_to_access_networks](resources--srv6_network_slice--reference--group-001.md#canonical-ba7f4e4ee691cd21843eefa2e212014b9aace4f5796d10d8ed39447c38b24b46) |
| `connect_to_enterprise_networks` | [connect_to_enterprise_networks](resources--srv6_network_slice--reference--group-001.md#canonical-3f123c26a55a6f32b9675a1aa011d6afca801e466fb032be64d898653f5947d1) |
| `connect_to_internet` | [connect_to_internet](resources--srv6_network_slice--reference--group-001.md#canonical-4b20dd12970071b6d1e9a09c6ccba976c9db5059ed461c4ff1f7ed64387742a4) |
| `description` | [description](resources--srv6_network_slice--reference--group-001.md#canonical-7398a463cc259f8b890bdecc5dd7d081706b10356d11af0a7d9782d883c4697a) |
| `disable` | [disable](resources--srv6_network_slice--reference--group-001.md#canonical-bce9e2ff0e9be7376248de4d4923aab72811670bbc7182c62cb920b576953e5a) |
| `id` | [id](resources--srv6_network_slice--reference--group-001.md#canonical-ef261ac59a2b35908cf0921515fe3c2d0b59918f4ef1fbac92ff8b0e48611746) |
| `labels` | [labels](resources--srv6_network_slice--reference--group-001.md#canonical-0391adbfef8c5a6ae08f4be13252b8402e1c25157a9e16e0379df8fab741b348) |
| `name` | [name](resources--srv6_network_slice--reference--group-001.md#canonical-f19432212dec036539b172bc70f61cbffec4b4c5441a9130d666ec6eb8009ad7) |
| `namespace` | [namespace](resources--srv6_network_slice--reference--group-001.md#canonical-f92c150e05bba48ff94081a1e1b81001683745a0c828520d3362d2a2dbb73645) |
| `sid_prefixes` | [sid_prefixes](resources--srv6_network_slice--reference--group-001.md#canonical-a34f3258c940279ce5501e7d8d38c7ae447b92efa469aa58799b95df8dfd72d3) |
| `timeouts` | [timeouts](resources--srv6_network_slice--reference--group-001.md#canonical-f786a81597c2917399b8c0d061a73b59c35252ebb6b9139e5561f5a690b6b302) |
| `timeouts.create` | [timeouts.create](resources--srv6_network_slice--reference--group-001.md#canonical-021764de23a09409f280a829bbd5ece424305c2acfde3b58c876c9cc5380ae37) |
| `timeouts.delete` | [timeouts.delete](resources--srv6_network_slice--reference--group-001.md#canonical-94e032e1ba504de546a75dab73275e575d9cefdd0dbde236306c0884ef77ca75) |
| `timeouts.read` | [timeouts.read](resources--srv6_network_slice--reference--group-001.md#canonical-51a22d5f9ce0b766578ca6ffdad450efec6bcdb259efb4fa46308e3c10b95c12) |
| `timeouts.update` | [timeouts.update](resources--srv6_network_slice--reference--group-001.md#canonical-8ba8c1d10128a1a1bc3c8d55651ac2ddb04e6cd6cf72cbe99e7306c0a64a7e4e) |

<a id="canonical-304a469990668cd66dbbfeac8bdf34dd84efe18e3fa7cfc0c96eb051d7358787"></a>

## Next pages — Property reference / 6ee38da4b4c6 / 16

- [timeouts](resources--srv6_network_slice--reference--group-001.md#canonical-bb84fcc743229e13d3b165a9c3c89d3253afa395a8d97cfc70dee42e9bbffb97)
- [xcsh_srv6_network_slice](../resources/srv6_network_slice.md#canonical-431435cc9f00e58094c1dd963b8aa9cec4251ea8dcbf3313b046ca1921402d3e)

<a id="canonical-bb84fcc743229e13d3b165a9c3c89d3253afa395a8d97cfc70dee42e9bbffb97"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7661bc4f401d5c662a73021849f197d6890aa88b75196a96ef564ddb8f4c6926"></a>

## timeouts — timeouts / eff01abfec1d / 2

Breadcrumbs:

- [xcsh_srv6_network_slice](../resources/srv6_network_slice.md#canonical-431435cc9f00e58094c1dd963b8aa9cec4251ea8dcbf3313b046ca1921402d3e)
- [Property reference](resources--srv6_network_slice--reference--group-001.md#canonical-a9921f7005bdd8167184c5a6a56fca49fff6fc9a386f4d22e032d049c63fc870)
- timeouts

<a id="canonical-f786a81597c2917399b8c0d061a73b59c35252ebb6b9139e5561f5a690b6b302"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-21fabf6b6cef739120eb1526d4688740b03f878de067b08d609d4cf81919f9b3"></a>

## Direct properties — timeouts / eff01abfec1d / 3

<a id="canonical-021764de23a09409f280a829bbd5ece424305c2acfde3b58c876c9cc5380ae37"></a>

<a id="canonical-05d27588b3de53f564dc7058f274764755b17030a07f0164200131eaf6a3300f"></a>

## create property — timeouts / eff01abfec1d / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-94e032e1ba504de546a75dab73275e575d9cefdd0dbde236306c0884ef77ca75"></a>

<a id="canonical-74e6793b9f9c64379e8c331a6af756cd294027849385e4ec38fe6f1ad339e7f0"></a>

## delete property — timeouts / eff01abfec1d / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-51a22d5f9ce0b766578ca6ffdad450efec6bcdb259efb4fa46308e3c10b95c12"></a>

<a id="canonical-6c3708a937f08e0fb941531b5ed6daeadc2d8ecf866c03cea01510623fbc4eea"></a>

## read property — timeouts / eff01abfec1d / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-8ba8c1d10128a1a1bc3c8d55651ac2ddb04e6cd6cf72cbe99e7306c0a64a7e4e"></a>

<a id="canonical-b07517a97b491cae757b2f8a9317c3b0e5daf4df4bcbc2ce7e662567adc3ae94"></a>

## update property — timeouts / eff01abfec1d / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-7afaf3176aa1d21b20111421e9b3fdbc6cea1785da88d41a08823a517aa6ee89"></a>

## Next pages — timeouts / eff01abfec1d / 8

- [Property reference](resources--srv6_network_slice--reference--group-001.md#canonical-a9921f7005bdd8167184c5a6a56fca49fff6fc9a386f4d22e032d049c63fc870)
- [xcsh_srv6_network_slice](../resources/srv6_network_slice.md#canonical-431435cc9f00e58094c1dd963b8aa9cec4251ea8dcbf3313b046ca1921402d3e)
