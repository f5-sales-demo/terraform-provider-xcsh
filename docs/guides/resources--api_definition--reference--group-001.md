---
page_title: "xcsh_api_definition reference"
subcategory: "API Management"
description: "Complete grouped canonical reference for xcsh_api_definition reference."
---

# xcsh_api_definition reference

<a id="canonical-471f59a89019910686397873c505f66f07862a821758d35780a19c9fb9b7aaa5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-80e7115d24d04f8780236b0c5ae54723f498467923f393c1c84efd0316b5a6cc"></a>

## Property reference — Property reference / 07e8bbe3bdf9 / 2

Breadcrumbs:

- [xcsh_api_definition](../resources/api_definition.md#canonical-1a9d8086edc81a3081bdb5af5cf08e65d779e0999f32054324107309653381b0)
- Property reference

<a id="canonical-45549a0bea156f2ab55ee0565d210e5dec9e5b1f2b8b9faace9d3a75fdb4e77a"></a>

## Direct properties — Property reference / 07e8bbe3bdf9 / 3

<a id="canonical-d2be01f35541d386903d48792141117153e268c5bba7f11a3723cefa55bb29f7"></a>

<a id="canonical-4e234ddbb6dcee023fec9aba86fb3d066b14a030c8efe7e56ceb74347c0be8ea"></a>

## annotations property — Property reference / 07e8bbe3bdf9 / 4

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

- [api_inventory_exclusion_list](resources--api_definition--reference--group-001.md#canonical-3a57ece012b36f3dacb7d4ba150ccc27d33f5417b6aea1ea3e7ff0fd6c1562d4): complete subsection reference.

- [api_inventory_inclusion_list](resources--api_definition--reference--group-001.md#canonical-3b8f0d3223b459a1dac5fb5260f38b53c7268295943d87813a8193b271a95451): complete subsection reference.

<a id="canonical-5689a30114093b22cd5aaa3c178bf503e8a1141997b9c00b3b49eea0cecfe477"></a>

<a id="canonical-8abdb22929ed44fd8fc7288b8aded4f938e997dbacaad06343d3b46d52097687"></a>

## description property — Property reference / 07e8bbe3bdf9 / 5

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

<a id="canonical-47028baf9ece142104aeed7a52b5fd2bcf0057592efd5003f5d0f1dfa91ed58d"></a>

<a id="canonical-dbdc9d172596c56d30e54d1ac730cbe87a5c619870da95df7b5c7c298542bb81"></a>

## disable property — Property reference / 07e8bbe3bdf9 / 6

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

<a id="canonical-17ae0e5c494ff8e1cd5d3814cf348dc3eacc0c9a085bfc33fd2d2a1e0f1698f0"></a>

<a id="canonical-95996ca785b0d34457d8830cf7dc3d6304f904a5c3f99e439485834baa717521"></a>

## id property — Property reference / 07e8bbe3bdf9 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-39134cc8fb1b8ebc1fae1a048d2f4abbecbb6ab33bbedaf65e25cf9ace318325"></a>

<a id="canonical-18bcf07764013b7fff394a45a0ff20a61ed23ca92a632d83615aa8ac5ec27c3b"></a>

## labels property — Property reference / 07e8bbe3bdf9 / 8

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

- [mixed_schema_origin](resources--api_definition--reference--group-001.md#canonical-8721b1f4be6e670a0bfc41696c5c2ab15c40a5bfb940f67eba8e179201d7e893): complete subsection reference.

<a id="canonical-30626fe53a675500fa13ba60e8d666dd77a6c6f5f77d201bcf79c743209483ec"></a>

<a id="canonical-7d6999c87a199c68e630ee70bb6acaee1880291a7c760d4e91452058b1958473"></a>

## name property — Property reference / 07e8bbe3bdf9 / 9

Type: `"string"`. Required.

Name of the API Definition. Must be unique within the namespace.

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

<a id="canonical-d3fb3f11f5c255055b07833bb3e53950c7d2fa83b13da2c903dd7f24abbcf452"></a>

<a id="canonical-2e24024703e3981758e56de7791e9cdf7521049590e16e0f99c37203bc0b8ef1"></a>

## namespace property — Property reference / 07e8bbe3bdf9 / 10

Type: `"string"`. Required.

Namespace where the API Definition is created.

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

- [non_api_endpoints](resources--api_definition--reference--group-001.md#canonical-c75787ca587824fe7a6604667e683b01de5ea22d0adbdaf19cf00c03770c717b): complete subsection reference.

- [strict_schema_origin](resources--api_definition--reference--group-001.md#canonical-b4c878f72d4018d76366a0f1b5059004007be54a51c900dd311f2fa34f7bd817): complete subsection reference.

<a id="canonical-dde53f29813ad52ee19cf7f65242d2835843056c07c2502f6f027ad7009b9c21"></a>

<a id="canonical-e596977bcf3593d0fce05b4dfc35af2ca5b9f381788e79ba8ed53821ed2e907b"></a>

## swagger_specs property — Property reference / 07e8bbe3bdf9 / 11

Type: `["list", "string"]`. Optional, Computed.

URLs of versioned OpenAPI files uploaded through Web App &amp; API Protection &gt; Files &gt;
Swagger Files. The 512-byte item limit is a URL-length limit; inline string:/// OpenAPI content is
rejected and does not create an API-definition object. Defaults to \`\[\]\`. Server applies default
when omitted.

Upstream description:

URLs of versioned OpenAPI files uploaded through Web App &amp; API Protection &gt; Files &gt;
Swagger Files. The 512-byte item limit is a URL-length limit; inline string:/// OpenAPI content is
rejected and does not create an API-definition object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "512",
    "ves.io.schema.rules.repeated.items.string.pattern": "/api/object_store/namespaces/([a-z]([-a-z0-9]*[a-z0-9])?)/stored_objects/swagger/([a-z]([-a-z0-9]*[a-z0-9])?)/(v|V)[0-9]+(-[0-9]{2}){3}$",
    "ves.io.schema.rules.repeated.max_items": "20",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [timeouts](resources--api_definition--reference--group-001.md#canonical-063f390185e91f4a6a78e9b784adbfe9a9b2ca5513f3cd6081f93c79122eae36): complete subsection reference.

<a id="canonical-2c14c247f46c09dfd86c45daf84150883ac4593f179f63957e70b59ac4989bcf"></a>

## All schema paths — Property reference / 07e8bbe3bdf9 / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--api_definition--reference--group-001.md#canonical-d2be01f35541d386903d48792141117153e268c5bba7f11a3723cefa55bb29f7) |
| `api_inventory_exclusion_list` | [api_inventory_exclusion_list](resources--api_definition--reference--group-001.md#canonical-5e3fe69a7617d3e5dfb3cb2a2d09b47a4d932299592f568bf18c82edaabcb80a) |
| `api_inventory_exclusion_list.method` | [api_inventory_exclusion_list.method](resources--api_definition--reference--group-001.md#canonical-2fa6d992fda2eccdacc1aa98fb80dc091b4f52cf548e5d1c4756fbbc0da5754a) |
| `api_inventory_exclusion_list.path` | [api_inventory_exclusion_list.path](resources--api_definition--reference--group-001.md#canonical-d79c278363c31987c85ec7d976baabfb2aee19f14dfa9957b294e1c4b33dd770) |
| `api_inventory_inclusion_list` | [api_inventory_inclusion_list](resources--api_definition--reference--group-001.md#canonical-aa7a97be354035a25b5746643132dc5688b704da4277dc15334657d35de2c65d) |
| `api_inventory_inclusion_list.method` | [api_inventory_inclusion_list.method](resources--api_definition--reference--group-001.md#canonical-7fd3be5d79ae100ff29d764a16a28b4e47057cf4db2bd3eef0958cb12cbdde7e) |
| `api_inventory_inclusion_list.path` | [api_inventory_inclusion_list.path](resources--api_definition--reference--group-001.md#canonical-e6a0eb1f2188410f35deb85f1d0cbf3e1b65a884316e8a744dd1212441d7acb3) |
| `description` | [description](resources--api_definition--reference--group-001.md#canonical-5689a30114093b22cd5aaa3c178bf503e8a1141997b9c00b3b49eea0cecfe477) |
| `disable` | [disable](resources--api_definition--reference--group-001.md#canonical-47028baf9ece142104aeed7a52b5fd2bcf0057592efd5003f5d0f1dfa91ed58d) |
| `id` | [id](resources--api_definition--reference--group-001.md#canonical-17ae0e5c494ff8e1cd5d3814cf348dc3eacc0c9a085bfc33fd2d2a1e0f1698f0) |
| `labels` | [labels](resources--api_definition--reference--group-001.md#canonical-39134cc8fb1b8ebc1fae1a048d2f4abbecbb6ab33bbedaf65e25cf9ace318325) |
| `mixed_schema_origin` | [mixed_schema_origin](resources--api_definition--reference--group-001.md#canonical-4fdbe3ed03157e7bda5a84ad5dfa5c27ec49ee986cf3d44cfab0950f798ba07b) |
| `name` | [name](resources--api_definition--reference--group-001.md#canonical-30626fe53a675500fa13ba60e8d666dd77a6c6f5f77d201bcf79c743209483ec) |
| `namespace` | [namespace](resources--api_definition--reference--group-001.md#canonical-d3fb3f11f5c255055b07833bb3e53950c7d2fa83b13da2c903dd7f24abbcf452) |
| `non_api_endpoints` | [non_api_endpoints](resources--api_definition--reference--group-001.md#canonical-7d034a603ec28ef8ba804d31b0f610bec685a8e06085325a8f3b6028fb9bc7f3) |
| `non_api_endpoints.method` | [non_api_endpoints.method](resources--api_definition--reference--group-001.md#canonical-d3f8a0fdda7cbac1dac40054f302f24e0ce33c24a96626bedc77000ba3cefc29) |
| `non_api_endpoints.path` | [non_api_endpoints.path](resources--api_definition--reference--group-001.md#canonical-9b6937b8ba43ddbd8566130207327b688e31dd3f4191f736937e0f11150cf1fb) |
| `strict_schema_origin` | [strict_schema_origin](resources--api_definition--reference--group-001.md#canonical-2540e3d8eb64629f5ae049ec109296aa07878fdbce5d1febed82e3b3370f4569) |
| `swagger_specs` | [swagger_specs](resources--api_definition--reference--group-001.md#canonical-dde53f29813ad52ee19cf7f65242d2835843056c07c2502f6f027ad7009b9c21) |
| `timeouts` | [timeouts](resources--api_definition--reference--group-001.md#canonical-0f44854c0db47a7abb79b28a2d3ae2520a5bf7682d220bf679a874d0e0ec9bca) |
| `timeouts.create` | [timeouts.create](resources--api_definition--reference--group-001.md#canonical-1caf5ec53e2726de925942101ac4d330850047855183550646bf44aa790ad784) |
| `timeouts.delete` | [timeouts.delete](resources--api_definition--reference--group-001.md#canonical-2098f8d9bc82ef8a1b22c94a056e6c9362b4accf97036d526b5f2518f54dcfba) |
| `timeouts.read` | [timeouts.read](resources--api_definition--reference--group-001.md#canonical-3db1ac2a24cdaca1449609b1075bbc18806e0c850f9a57c6004e69d5d5cfbd64) |
| `timeouts.update` | [timeouts.update](resources--api_definition--reference--group-001.md#canonical-d18f926d39417f815c165ec231abe9ed436313e568ef56a11f35a7917152a67e) |

<a id="canonical-82cd5665368b6e3a7491ad7e19feea7696c4750c0e7c414efc8f9b6164d3244e"></a>

## Next pages — Property reference / 07e8bbe3bdf9 / 13

- [api_inventory_exclusion_list](resources--api_definition--reference--group-001.md#canonical-3a57ece012b36f3dacb7d4ba150ccc27d33f5417b6aea1ea3e7ff0fd6c1562d4)
- [api_inventory_inclusion_list](resources--api_definition--reference--group-001.md#canonical-3b8f0d3223b459a1dac5fb5260f38b53c7268295943d87813a8193b271a95451)
- [mixed_schema_origin](resources--api_definition--reference--group-001.md#canonical-8721b1f4be6e670a0bfc41696c5c2ab15c40a5bfb940f67eba8e179201d7e893)
- [non_api_endpoints](resources--api_definition--reference--group-001.md#canonical-c75787ca587824fe7a6604667e683b01de5ea22d0adbdaf19cf00c03770c717b)
- [strict_schema_origin](resources--api_definition--reference--group-001.md#canonical-b4c878f72d4018d76366a0f1b5059004007be54a51c900dd311f2fa34f7bd817)
- [timeouts](resources--api_definition--reference--group-001.md#canonical-063f390185e91f4a6a78e9b784adbfe9a9b2ca5513f3cd6081f93c79122eae36)
- [xcsh_api_definition](../resources/api_definition.md#canonical-1a9d8086edc81a3081bdb5af5cf08e65d779e0999f32054324107309653381b0)

<a id="canonical-3a57ece012b36f3dacb7d4ba150ccc27d33f5417b6aea1ea3e7ff0fd6c1562d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-09e55a9430025d943ceaa3b312636614da230b6a854cf7743b397e5c6a32d801"></a>

## api_inventory_exclusion_list — api_inventory_exclusion_list / 444e03491f36 / 2

Breadcrumbs:

- [xcsh_api_definition](../resources/api_definition.md#canonical-1a9d8086edc81a3081bdb5af5cf08e65d779e0999f32054324107309653381b0)
- [Property reference](resources--api_definition--reference--group-001.md#canonical-471f59a89019910686397873c505f66f07862a821758d35780a19c9fb9b7aaa5)
- api_inventory_exclusion_list

<a id="canonical-5e3fe69a7617d3e5dfb3cb2a2d09b47a4d932299592f568bf18c82edaabcb80a"></a>

Type: `"object"`. list nested block, Optional.

List of API Endpoints excluded from the API Inventory. Defaults to \`\[\]\`. Server applies default
when omitted.

Upstream description:

List of API Endpoints excluded from the API Inventory.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("path")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5000,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "5000",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "5000",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
api_inventory_exclusion_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-8ae06b241b21287c6b2c63afb011111cf8232562b85f0e460a2176cd6ce2e98b"></a>

## Direct properties — api_inventory_exclusion_list / 444e03491f36 / 3

<a id="canonical-2fa6d992fda2eccdacc1aa98fb80dc091b4f52cf548e5d1c4756fbbc0da5754a"></a>

<a id="canonical-2a0d98c7c0a5d050538671a6dbd30d24ec72f45b8dd38745cd249862b22cce77"></a>

## method property — api_inventory_exclusion_list / 444e03491f36 / 4

Type: `"string"`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-d79c278363c31987c85ec7d976baabfb2aee19f14dfa9957b294e1c4b33dd770"></a>

<a id="canonical-841315e6715485f8cb5f25a86330a2a04d918853c52a51b07b3b0b30bfdfe0c7"></a>

## path property — api_inventory_exclusion_list / 444e03491f36 / 5

Type: `"string"`. Optional.

Endpoint path, as specified in OpenAPI, including parameters. The path should comply with RFC 3986
and may have parameters according to OpenAPI specification.

Upstream description:

An endpoint path, as specified in OpenAPI, including parameters. The path should comply with RFC
3986 and may have parameters according to OpenAPI specification.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

<a id="canonical-2b9f95ca36f62b48b5b766147b95500a8bec28c1e3babb9fb2d3f6d28cf4ea44"></a>

## Next pages — api_inventory_exclusion_list / 444e03491f36 / 6

- [Property reference](resources--api_definition--reference--group-001.md#canonical-471f59a89019910686397873c505f66f07862a821758d35780a19c9fb9b7aaa5)
- [xcsh_api_definition](../resources/api_definition.md#canonical-1a9d8086edc81a3081bdb5af5cf08e65d779e0999f32054324107309653381b0)

<a id="canonical-3b8f0d3223b459a1dac5fb5260f38b53c7268295943d87813a8193b271a95451"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-811470415e76d71d2da27c0a0b07ec21c82e567e592e3e845bf367edf9719125"></a>

## api_inventory_inclusion_list — api_inventory_inclusion_list / ec5f18f86526 / 2

Breadcrumbs:

- [xcsh_api_definition](../resources/api_definition.md#canonical-1a9d8086edc81a3081bdb5af5cf08e65d779e0999f32054324107309653381b0)
- [Property reference](resources--api_definition--reference--group-001.md#canonical-471f59a89019910686397873c505f66f07862a821758d35780a19c9fb9b7aaa5)
- api_inventory_inclusion_list

<a id="canonical-aa7a97be354035a25b5746643132dc5688b704da4277dc15334657d35de2c65d"></a>

Type: `"object"`. list nested block, Optional.

List of API Endpoints included in the API Inventory. Typically, discovered API endpoints are added
to the API Inventory using this list. Defaults to \`\[\]\`. Server applies default when omitted.

Upstream description:

List of API Endpoints included in the API Inventory. Typically, discovered API endpoints are added
to the API Inventory using this list.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("path")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5000,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "5000",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "5000",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
api_inventory_inclusion_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-a5c147cb4e68793fff6b50ae1ba8d2a52efbcf975cffdb054ed1cb6a339b91c7"></a>

## Direct properties — api_inventory_inclusion_list / ec5f18f86526 / 3

<a id="canonical-7fd3be5d79ae100ff29d764a16a28b4e47057cf4db2bd3eef0958cb12cbdde7e"></a>

<a id="canonical-9e6afec506509311e99770cbcf29817c0d23ba437e159d1d715e8203d917f3f4"></a>

## method property — api_inventory_inclusion_list / ec5f18f86526 / 4

Type: `"string"`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-e6a0eb1f2188410f35deb85f1d0cbf3e1b65a884316e8a744dd1212441d7acb3"></a>

<a id="canonical-3a6b9a8a470538b1360dfcd2332b49cb3f6f3ca294853ade1a5a8f69c81360a7"></a>

## path property — api_inventory_inclusion_list / ec5f18f86526 / 5

Type: `"string"`. Optional.

Endpoint path, as specified in OpenAPI, including parameters. The path should comply with RFC 3986
and may have parameters according to OpenAPI specification.

Upstream description:

An endpoint path, as specified in OpenAPI, including parameters. The path should comply with RFC
3986 and may have parameters according to OpenAPI specification.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

<a id="canonical-f039c87f92e49456991da09e32154ad8e18774258f5632be3720e7a1e51f99f4"></a>

## Next pages — api_inventory_inclusion_list / ec5f18f86526 / 6

- [Property reference](resources--api_definition--reference--group-001.md#canonical-471f59a89019910686397873c505f66f07862a821758d35780a19c9fb9b7aaa5)
- [xcsh_api_definition](../resources/api_definition.md#canonical-1a9d8086edc81a3081bdb5af5cf08e65d779e0999f32054324107309653381b0)

<a id="canonical-8721b1f4be6e670a0bfc41696c5c2ab15c40a5bfb940f67eba8e179201d7e893"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e8656b44e60c9484510a1593146ddea9b0c73703796f6c0786614a6d27b3bea"></a>

## mixed_schema_origin — mixed_schema_origin / 929f1ad1b431 / 2

Breadcrumbs:

- [xcsh_api_definition](../resources/api_definition.md#canonical-1a9d8086edc81a3081bdb5af5cf08e65d779e0999f32054324107309653381b0)
- [Property reference](resources--api_definition--reference--group-001.md#canonical-471f59a89019910686397873c505f66f07862a821758d35780a19c9fb9b7aaa5)
- mixed_schema_origin

<a id="canonical-4fdbe3ed03157e7bda5a84ad5dfa5c27ec49ee986cf3d44cfab0950f798ba07b"></a>

Type: `["object", {}]`. Optional.

\[OneOf: mixed\_schema\_origin, strict\_schema\_origin\] Configuration parameter for mixed schema
origin.

Upstream description:

This can be used for messages where no values are needed.

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

OneOf alternatives in this subsection:

- [mixed_schema_origin](resources--api_definition--reference--group-001.md#canonical-4fdbe3ed03157e7bda5a84ad5dfa5c27ec49ee986cf3d44cfab0950f798ba07b)
- [strict_schema_origin](resources--api_definition--reference--group-001.md#canonical-2540e3d8eb64629f5ae049ec109296aa07878fdbce5d1febed82e3b3370f4569)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
mixed_schema_origin = {}
```

<a id="canonical-217a7571067447759d0f57dc1d29b0480ac9acd46314cd3682c6aaae6aafa4cc"></a>

## Direct properties — mixed_schema_origin / 929f1ad1b431 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6a928efc94e92ee9386ebddcb7a1ea18edbf7abbd14c90676d87a99d2ec44b96"></a>

## Next pages — mixed_schema_origin / 929f1ad1b431 / 4

- [Property reference](resources--api_definition--reference--group-001.md#canonical-471f59a89019910686397873c505f66f07862a821758d35780a19c9fb9b7aaa5)
- [xcsh_api_definition](../resources/api_definition.md#canonical-1a9d8086edc81a3081bdb5af5cf08e65d779e0999f32054324107309653381b0)

<a id="canonical-c75787ca587824fe7a6604667e683b01de5ea22d0adbdaf19cf00c03770c717b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2dba195694c83e3d412e2ed78b8a5a4dd0aace5ad0ecb4550f6e5056b4c54a7d"></a>

## non_api_endpoints — non_api_endpoints / 3b0abc7c616a / 2

Breadcrumbs:

- [xcsh_api_definition](../resources/api_definition.md#canonical-1a9d8086edc81a3081bdb5af5cf08e65d779e0999f32054324107309653381b0)
- [Property reference](resources--api_definition--reference--group-001.md#canonical-471f59a89019910686397873c505f66f07862a821758d35780a19c9fb9b7aaa5)
- non_api_endpoints

<a id="canonical-7d034a603ec28ef8ba804d31b0f610bec685a8e06085325a8f3b6028fb9bc7f3"></a>

Type: `"object"`. list nested block, Optional.

API Discovery Exclusion List. List of Non-API Endpoints. Defaults to \`\[\]\`. Server applies
default when omitted.

Upstream description:

List of Non-API Endpoints.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("path")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5000,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "5000",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "5000",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
non_api_endpoints {
  # Configure direct properties listed below.
}
```

<a id="canonical-7de3603456e5c89566ffb9f660c0708082e96f33933cc66f42af67ec66ce3943"></a>

## Direct properties — non_api_endpoints / 3b0abc7c616a / 3

<a id="canonical-d3f8a0fdda7cbac1dac40054f302f24e0ce33c24a96626bedc77000ba3cefc29"></a>

<a id="canonical-df1ce8df62e2d7f263263dd681da81a40d76dc4fad169b35f47878b1427eb3f5"></a>

## method property — non_api_endpoints / 3b0abc7c616a / 4

Type: `"string"`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-9b6937b8ba43ddbd8566130207327b688e31dd3f4191f736937e0f11150cf1fb"></a>

<a id="canonical-62758950252f25865de54c6cf166a17b4cca53afc825af7ebb26ced1421fe938"></a>

## path property — non_api_endpoints / 3b0abc7c616a / 5

Type: `"string"`. Optional.

Endpoint path, as specified in OpenAPI, including parameters. The path should comply with RFC 3986
and may have parameters according to OpenAPI specification.

Upstream description:

An endpoint path, as specified in OpenAPI, including parameters. The path should comply with RFC
3986 and may have parameters according to OpenAPI specification.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

<a id="canonical-8123f8e2f1595783528e76c957d7429c4de1be5322bc07f0407b49e50a09a8a6"></a>

## Next pages — non_api_endpoints / 3b0abc7c616a / 6

- [Property reference](resources--api_definition--reference--group-001.md#canonical-471f59a89019910686397873c505f66f07862a821758d35780a19c9fb9b7aaa5)
- [xcsh_api_definition](../resources/api_definition.md#canonical-1a9d8086edc81a3081bdb5af5cf08e65d779e0999f32054324107309653381b0)

<a id="canonical-b4c878f72d4018d76366a0f1b5059004007be54a51c900dd311f2fa34f7bd817"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-92e8817f646fe30670c3c913ff01b21e2a52c69f077d4def7ba0a8ff61348888"></a>

## strict_schema_origin — strict_schema_origin / abb1079a650d / 2

Breadcrumbs:

- [xcsh_api_definition](../resources/api_definition.md#canonical-1a9d8086edc81a3081bdb5af5cf08e65d779e0999f32054324107309653381b0)
- [Property reference](resources--api_definition--reference--group-001.md#canonical-471f59a89019910686397873c505f66f07862a821758d35780a19c9fb9b7aaa5)
- strict_schema_origin

<a id="canonical-2540e3d8eb64629f5ae049ec109296aa07878fdbce5d1febed82e3b3370f4569"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for strict schema origin. Defaults to \`map\[\]\`. Server applies default
when omitted.

Upstream description:

This can be used for messages where no values are needed.

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
strict_schema_origin = {}
```

<a id="canonical-241aa843192f8cee0d34eb002662ba52545f450e9540a898c2436c3fd824217d"></a>

## Direct properties — strict_schema_origin / abb1079a650d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e9a9c8d049a2137410d7be95e6dc9c56d9c227f1253fb6116b46e52fcd857c69"></a>

## Next pages — strict_schema_origin / abb1079a650d / 4

- [Property reference](resources--api_definition--reference--group-001.md#canonical-471f59a89019910686397873c505f66f07862a821758d35780a19c9fb9b7aaa5)
- [xcsh_api_definition](../resources/api_definition.md#canonical-1a9d8086edc81a3081bdb5af5cf08e65d779e0999f32054324107309653381b0)

<a id="canonical-063f390185e91f4a6a78e9b784adbfe9a9b2ca5513f3cd6081f93c79122eae36"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-96b6a1c9da325718d3ab771b48e3ab5e617eb6951b64f756d58d8c89e8391b80"></a>

## timeouts — timeouts / 672518abbfc2 / 2

Breadcrumbs:

- [xcsh_api_definition](../resources/api_definition.md#canonical-1a9d8086edc81a3081bdb5af5cf08e65d779e0999f32054324107309653381b0)
- [Property reference](resources--api_definition--reference--group-001.md#canonical-471f59a89019910686397873c505f66f07862a821758d35780a19c9fb9b7aaa5)
- timeouts

<a id="canonical-0f44854c0db47a7abb79b28a2d3ae2520a5bf7682d220bf679a874d0e0ec9bca"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-9f68eae1685e56876ed4ae5a3d24aa8f56153baadd9de1810c619ce6cef1d21d"></a>

## Direct properties — timeouts / 672518abbfc2 / 3

<a id="canonical-1caf5ec53e2726de925942101ac4d330850047855183550646bf44aa790ad784"></a>

<a id="canonical-f572f93491dee1eb5de42b3b1e9d23a78af8c538e0aa7adbb2b837c0e5a0181c"></a>

## create property — timeouts / 672518abbfc2 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2098f8d9bc82ef8a1b22c94a056e6c9362b4accf97036d526b5f2518f54dcfba"></a>

<a id="canonical-d4298e014589e37835dd147585dbdd286b347464e250b532d9436a38ccdccdae"></a>

## delete property — timeouts / 672518abbfc2 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3db1ac2a24cdaca1449609b1075bbc18806e0c850f9a57c6004e69d5d5cfbd64"></a>

<a id="canonical-2998827ba5ce42cdb14ed3f939636394a6019eb0d577d7310223b23aaa9002d5"></a>

## read property — timeouts / 672518abbfc2 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-d18f926d39417f815c165ec231abe9ed436313e568ef56a11f35a7917152a67e"></a>

<a id="canonical-0fdacce4adbf9501595b00bc92111160743def7a3481e7ca33f236d4acfe6d6e"></a>

## update property — timeouts / 672518abbfc2 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-93f268f338089a36c67f824965c1475284ca086ca45c02b75fd503e4d00e1b79"></a>

## Next pages — timeouts / 672518abbfc2 / 8

- [Property reference](resources--api_definition--reference--group-001.md#canonical-471f59a89019910686397873c505f66f07862a821758d35780a19c9fb9b7aaa5)
- [xcsh_api_definition](../resources/api_definition.md#canonical-1a9d8086edc81a3081bdb5af5cf08e65d779e0999f32054324107309653381b0)
