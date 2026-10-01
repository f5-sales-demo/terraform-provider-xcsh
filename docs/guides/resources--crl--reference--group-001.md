---
page_title: "xcsh_crl reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_crl reference."
---

# xcsh_crl reference

<a id="canonical-0a1bcf7add556b43d14c3820e55e0a6110324a87da60186d84fa20c107de9d30"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d700477b7c8d63f3db9ae1e660ff0dadaf0da83987126c736f4fdf22a223dc6f"></a>

## Property reference — Property reference / 07ab472888cd / 2

Breadcrumbs:

- [xcsh_crl](../resources/crl.md#canonical-96a63453b75e823f71de872f51de2a086a64ac241ba947a0a9c61f4a01cb3404)
- Property reference

<a id="canonical-9de1fd54645934d57356e3f236267dbf4f277de00670f184a006b15a5b19a39a"></a>

## Direct properties — Property reference / 07ab472888cd / 3

<a id="canonical-ab9a84170dc7f9c7e799b86f95fe2c38f7318be109069eb130b1bc788564e5c6"></a>

<a id="canonical-131dc6650cc1a69df59986625066cb75e9b3e33c8ea4932ab134abab0f9c4481"></a>

## annotations property — Property reference / 07ab472888cd / 4

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

<a id="canonical-b96b107bc88ef463999915852d1977c0cf29a37db0654eca8265b13068e09c7f"></a>

<a id="canonical-026d6fe7899862b30f63a5ce603d5f876e35e251f65d9fc153b57d260ae0f650"></a>

## description property — Property reference / 07ab472888cd / 5

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

<a id="canonical-f5b57a90a0ef47775cf54b94cc740728b1b642f25864e4e000d9d3c0fb7702a3"></a>

<a id="canonical-1ae928315efbb9ac4ff5ff5d43cfca7127c1daef3acfb2f364fb6dbabbb17e79"></a>

## disable property — Property reference / 07ab472888cd / 6

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

- [http_access](resources--crl--reference--group-001.md#canonical-f9300ba2540c2e95d4b977e41822f343904828ea111ab54d3df1befee814c370): complete subsection reference.

<a id="canonical-6f5b97a11a22fc5d3778746110327ee89edfd7fd88fc0c69188cc047845c1573"></a>

<a id="canonical-c1d759de93fd83d60372e287166eb6087e5d8ba177545d24e32ce6b78961d8be"></a>

## id property — Property reference / 07ab472888cd / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2e915ec5fa998b1830752f8c08c5431070c0dc580241b5120852c40c89cb295c"></a>

<a id="canonical-a2f33d448e8226678ef962c4e58b314dd2a240d0349b7dfd79782a1096f58629"></a>

## labels property — Property reference / 07ab472888cd / 8

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

<a id="canonical-c69a9d673fa661916777b77e09a443f295dc3ff4455501aefbfed9c60a206ca2"></a>

<a id="canonical-05af966f56c32a67dee1928df5507d7587737197d6fbb2550df66b9af880f85f"></a>

## name property — Property reference / 07ab472888cd / 9

Type: `"string"`. Required.

Name of the CRL. Must be unique within the namespace.

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

<a id="canonical-91805e0fa3155f700256071c05ac21af7361c0bc59a2db7841e427344d09fdc4"></a>

<a id="canonical-66f33b5cd4d0b224d545565c61d17a6918381db7a97752e6b6b2656e48ff8509"></a>

## namespace property — Property reference / 07ab472888cd / 10

Type: `"string"`. Required.

Namespace where the CRL is created.

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

<a id="canonical-e60920ec12ef8d3c95518addd0ea96c47d66bdff72005e31da340a0731480572"></a>

<a id="canonical-45e8b241fd81c4e7e6d4d297e0917920bc222511cd36c00ef55c4c73d68ff4fc"></a>

## refresh_interval property — Property reference / 07ab472888cd / 11

Type: `"number"`. Required.

CRL Refresh interval. CRL refresh interval, in hours.

Upstream description:

CRL refresh interval, in hours.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(6, 168),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 168,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 6
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "6",
    "ves.io.schema.rules.uint32.lte": "168"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "6",
    "ves.io.schema.rules.uint32.lte": "168"
  }
}
```

<a id="canonical-d4f0e63d8b140d822f4f9080c581c36ed6409b602fe7b2fe6bb034912dda015b"></a>

<a id="canonical-5feb32b195cea22bd7e0cd10d1578a99c23b481f972e61ca3e1ec66c8bf30bed"></a>

## server_address property — Property reference / 07ab472888cd / 12

Type: `"string"`. Required.

CRL Server address. CRL server address or hostname.

Upstream description:

CRL server address or hostname.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
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
    "ves.io.schema.rules.string.hostname_or_ip": "true",
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname_or_ip": "true",
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

<a id="canonical-f09b1153b7d7c8d8ffaab65d04c46f9c61fec9d23d1803f79fa87f3f2f86b2e6"></a>

<a id="canonical-253b7359ae27aaa66612c23690791e54836574db84cd4c8a2c8e731c58781c86"></a>

## server_port property — Property reference / 07ab472888cd / 13

Type: `"number"`. Required.

CRL Server Port. Set CRL Server port number.

Upstream description:

Set CRL Server port number.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-8649d1f39e70bf157276421bfac3b43453dcdfd331e35893f9ee18393c800d58"></a>

<a id="canonical-81b0d412bd31ab4d919dfbffb3b24a0cd72c3d3b5a75a4c9af5620fe8b4c7ef2"></a>

## timeout property — Property reference / 07ab472888cd / 14

Type: `"number"`. Required.

CRL download timeout. CRL download wait time, in seconds.

Upstream description:

CRL download wait time, in seconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 180),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 180,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "180"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "180"
  }
}
```

- [timeouts](resources--crl--reference--group-001.md#canonical-283125fe51ce42d36a691b088faaae078cfb08869081cef252d73f3a3aee6702): complete subsection reference.

<a id="canonical-fcc321ae67d9464e35a856642a0a15accfa91354702415abd999d11373d92867"></a>

## All schema paths — Property reference / 07ab472888cd / 15

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--crl--reference--group-001.md#canonical-ab9a84170dc7f9c7e799b86f95fe2c38f7318be109069eb130b1bc788564e5c6) |
| `description` | [description](resources--crl--reference--group-001.md#canonical-b96b107bc88ef463999915852d1977c0cf29a37db0654eca8265b13068e09c7f) |
| `disable` | [disable](resources--crl--reference--group-001.md#canonical-f5b57a90a0ef47775cf54b94cc740728b1b642f25864e4e000d9d3c0fb7702a3) |
| `http_access` | [http_access](resources--crl--reference--group-001.md#canonical-d8bca8e51a61e5905230f9043af9fb1ceb1529b18c69757c625bb9786d7f5af2) |
| `http_access.path` | [http_access.path](resources--crl--reference--group-001.md#canonical-fa30ad852670ea57378ebc951ec7ed7eb3ffeca57ae4fb652c77558b5d8bd7f0) |
| `id` | [id](resources--crl--reference--group-001.md#canonical-6f5b97a11a22fc5d3778746110327ee89edfd7fd88fc0c69188cc047845c1573) |
| `labels` | [labels](resources--crl--reference--group-001.md#canonical-2e915ec5fa998b1830752f8c08c5431070c0dc580241b5120852c40c89cb295c) |
| `name` | [name](resources--crl--reference--group-001.md#canonical-c69a9d673fa661916777b77e09a443f295dc3ff4455501aefbfed9c60a206ca2) |
| `namespace` | [namespace](resources--crl--reference--group-001.md#canonical-91805e0fa3155f700256071c05ac21af7361c0bc59a2db7841e427344d09fdc4) |
| `refresh_interval` | [refresh_interval](resources--crl--reference--group-001.md#canonical-e60920ec12ef8d3c95518addd0ea96c47d66bdff72005e31da340a0731480572) |
| `server_address` | [server_address](resources--crl--reference--group-001.md#canonical-d4f0e63d8b140d822f4f9080c581c36ed6409b602fe7b2fe6bb034912dda015b) |
| `server_port` | [server_port](resources--crl--reference--group-001.md#canonical-f09b1153b7d7c8d8ffaab65d04c46f9c61fec9d23d1803f79fa87f3f2f86b2e6) |
| `timeout` | [timeout](resources--crl--reference--group-001.md#canonical-8649d1f39e70bf157276421bfac3b43453dcdfd331e35893f9ee18393c800d58) |
| `timeouts` | [timeouts](resources--crl--reference--group-001.md#canonical-a13c1342230193cfd62bc36416eac6739fbd806a4eca7fe7ddffb503876a030d) |
| `timeouts.create` | [timeouts.create](resources--crl--reference--group-001.md#canonical-189e7c054814f8a1128b12a0ad640caf6082294ee727caae82821f269004e3f4) |
| `timeouts.delete` | [timeouts.delete](resources--crl--reference--group-001.md#canonical-e789fe87218c43d25bbc088cfcae83ca94658bd4f5544baca309c1a70f342cea) |
| `timeouts.read` | [timeouts.read](resources--crl--reference--group-001.md#canonical-e800cc2050f6cf8410c7a24915aace709caa247e8633f213638c7458ac2c9054) |
| `timeouts.update` | [timeouts.update](resources--crl--reference--group-001.md#canonical-dc36d6636e08d79c937d29aa32beac2fbfd5da515aa716a231b0a43035af6c2a) |

<a id="canonical-ca82fd3f502f217dd6777d6615f1f4a3265ce2e8422754f9e3317dec248c6113"></a>

## Next pages — Property reference / 07ab472888cd / 16

- [http_access](resources--crl--reference--group-001.md#canonical-f9300ba2540c2e95d4b977e41822f343904828ea111ab54d3df1befee814c370)
- [timeouts](resources--crl--reference--group-001.md#canonical-283125fe51ce42d36a691b088faaae078cfb08869081cef252d73f3a3aee6702)
- [xcsh_crl](../resources/crl.md#canonical-96a63453b75e823f71de872f51de2a086a64ac241ba947a0a9c61f4a01cb3404)

<a id="canonical-f9300ba2540c2e95d4b977e41822f343904828ea111ab54d3df1befee814c370"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63b260d47b75e9afd402396a55541d966a168b64bd150ae3870ca63fc518c268"></a>

## http_access — http_access / 569f3d031597 / 2

Breadcrumbs:

- [xcsh_crl](../resources/crl.md#canonical-96a63453b75e823f71de872f51de2a086a64ac241ba947a0a9c61f4a01cb3404)
- [Property reference](resources--crl--reference--group-001.md#canonical-0a1bcf7add556b43d14c3820e55e0a6110324a87da60186d84fa20c107de9d30)
- http_access

<a id="canonical-d8bca8e51a61e5905230f9043af9fb1ceb1529b18c69757c625bb9786d7f5af2"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http access.

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
http_access {
  # Configure direct properties listed below.
}
```

<a id="canonical-6fb7fc12306f158061bff2de2503692776130edd1624119aa08ba5d646949204"></a>

## Direct properties — http_access / 569f3d031597 / 3

<a id="canonical-fa30ad852670ea57378ebc951ec7ed7eb3ffeca57ae4fb652c77558b5d8bd7f0"></a>

<a id="canonical-7060cdfef85a2b44004c4ee5739ab18b0a1735fb86f16e7d4570b7848649feaa"></a>

## path property — http_access / 569f3d031597 / 4

Type: `"string"`. Optional.

CRL File path. CRL file location.

Upstream description:

CRL file location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-d407c04cfe1c578fc6904d383546498f308f439ce7039d8cb4164dd5222496d2"></a>

## Next pages — http_access / 569f3d031597 / 5

- [Property reference](resources--crl--reference--group-001.md#canonical-0a1bcf7add556b43d14c3820e55e0a6110324a87da60186d84fa20c107de9d30)
- [xcsh_crl](../resources/crl.md#canonical-96a63453b75e823f71de872f51de2a086a64ac241ba947a0a9c61f4a01cb3404)

<a id="canonical-283125fe51ce42d36a691b088faaae078cfb08869081cef252d73f3a3aee6702"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0ce06625223a4af70b1989dda59f390c7d0ae13268876b8cad84db6fcb7a5ad4"></a>

## timeouts — timeouts / 1ac4e2cf1f23 / 2

Breadcrumbs:

- [xcsh_crl](../resources/crl.md#canonical-96a63453b75e823f71de872f51de2a086a64ac241ba947a0a9c61f4a01cb3404)
- [Property reference](resources--crl--reference--group-001.md#canonical-0a1bcf7add556b43d14c3820e55e0a6110324a87da60186d84fa20c107de9d30)
- timeouts

<a id="canonical-a13c1342230193cfd62bc36416eac6739fbd806a4eca7fe7ddffb503876a030d"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-cc081b1a9a2a509c4977293e0341c74d754380d227bfd82b2ebd7eeedf582a5d"></a>

## Direct properties — timeouts / 1ac4e2cf1f23 / 3

<a id="canonical-189e7c054814f8a1128b12a0ad640caf6082294ee727caae82821f269004e3f4"></a>

<a id="canonical-bee5c30ec1f6cce766f1c9faffde214c2a75ea71638cfb5257c5563f0dc693a0"></a>

## create property — timeouts / 1ac4e2cf1f23 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-e789fe87218c43d25bbc088cfcae83ca94658bd4f5544baca309c1a70f342cea"></a>

<a id="canonical-2d0806f2ffb8b51d9112a70a850d0e98ce8197744454c776195f602971e63b78"></a>

## delete property — timeouts / 1ac4e2cf1f23 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-e800cc2050f6cf8410c7a24915aace709caa247e8633f213638c7458ac2c9054"></a>

<a id="canonical-45201ccc0abe229d3e92fe6148a84c9b0d6f8aeaff57dbe22e581c3d58d1603c"></a>

## read property — timeouts / 1ac4e2cf1f23 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-dc36d6636e08d79c937d29aa32beac2fbfd5da515aa716a231b0a43035af6c2a"></a>

<a id="canonical-062d520a9a9d61076e58ddc2b6b8b4f0004013a251741db1ee743c4cc6d8c9fa"></a>

## update property — timeouts / 1ac4e2cf1f23 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-fec9414045d0314fbe2adc336d3400b889c4fee9416c605f7f24a6597091d839"></a>

## Next pages — timeouts / 1ac4e2cf1f23 / 8

- [Property reference](resources--crl--reference--group-001.md#canonical-0a1bcf7add556b43d14c3820e55e0a6110324a87da60186d84fa20c107de9d30)
- [xcsh_crl](../resources/crl.md#canonical-96a63453b75e823f71de872f51de2a086a64ac241ba947a0a9c61f4a01cb3404)
