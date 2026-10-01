---
page_title: "xcsh_api_definition reference"
subcategory: "API Management"
description: "Complete grouped canonical reference for xcsh_api_definition reference."
---

# xcsh_api_definition reference

<a id="canonical-fad0050df3f4cd1ee16f15782b363b846ffd70034b8a18de354e7a072ec26d5b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7855e0e3541b00f4e05e9ec84e4f35bf52de30c67cc668bbbd1c9a946da8ea28"></a>

## Property reference — Property reference / 533de0104972 / 2

Breadcrumbs:

- [xcsh_api_definition](../data-sources/api_definition.md#canonical-1df393999e0c5facf84460c20dfc66c998ea2afc9c81bd04c6d3cf585cad53ab)
- Property reference

<a id="canonical-392b61e226d9a58725cb304ab34f6ab749970d3188e8c8eec4de5cae57d9917d"></a>

## Direct properties — Property reference / 533de0104972 / 3

<a id="canonical-a2877144ef18a48bad3fbdab5f358b9f041deccf3daa3cd10a42b75991c1e620"></a>

<a id="canonical-98cb5e2362d9294a60e71fc9ac0929584118a914ddac1e61d2449a9bfcff9fc1"></a>

## annotations property — Property reference / 533de0104972 / 4

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

- [api_inventory_exclusion_list](data-sources--api_definition--reference--group-001.md#canonical-2e076e32eb0f9baac97a955dba020f38ebfd3d65c8af0fb5c4736c2ac2b1f19c): complete subsection reference.

- [api_inventory_inclusion_list](data-sources--api_definition--reference--group-001.md#canonical-198d8acda02ecd2dd684e9389fc7d3a6b105de4e4510ce3873679235610d522c): complete subsection reference.

<a id="canonical-759fa549790769fcbb51bd7506a9c7aebf14a37e6f370792f261e1e97ee242f4"></a>

<a id="canonical-8ebebb5aa480a2d37d4e2a95c8c71b54713eaa59830f4f1aee7005d873cd330d"></a>

## description property — Property reference / 533de0104972 / 5

Type: `"string"`. Computed.

Description of the APIDefinition.

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

<a id="canonical-ee4d22dac33b484684b7b926cfcb0251983a52f08f3fc7d207b1588d0223a6a2"></a>

<a id="canonical-1fe4853b6419b4b45d82b31b03537e2c431d7b811cc1a6cc035d3f13d1411b4b"></a>

## id property — Property reference / 533de0104972 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-991ff040fadacc440f5c60678943d13f071d5ff98cfdbf7c1a4280e258c7da7f"></a>

<a id="canonical-f0f41b7fc72b0eab6440d1f6e56abf0e7e500b04dcee6cc341de60e4bd662373"></a>

## labels property — Property reference / 533de0104972 / 7

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

- [mixed_schema_origin](data-sources--api_definition--reference--group-001.md#canonical-57db1b0d2c8304075cf9831004863a1441d8563f5cb9e8f8e869a761bcebc086): complete subsection reference.

<a id="canonical-5469b694f67f8b38e063ef9b87a5f80c1c56709aef64ea33e024a7f98a6c6b0a"></a>

<a id="canonical-65008405f3b059331f84db78d51fade99a25a8829d3ff2cbb2acde106f0555f4"></a>

## name property — Property reference / 533de0104972 / 8

Type: `"string"`. Required.

Name of the APIDefinition.

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

<a id="canonical-6edaf113889090c70298e2c7b45fc1978497805efdffb7b899bdc1b4965881b3"></a>

<a id="canonical-4811d632e32061870d13c9cbb6c344ca062eb984db13ea95c095a7233a450a17"></a>

## namespace property — Property reference / 533de0104972 / 9

Type: `"string"`. Required.

Namespace where the APIDefinition exists.

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

- [non_api_endpoints](data-sources--api_definition--reference--group-001.md#canonical-cb5c898f0a4624e8c07fe12db587917411cb88689982594ff5967457fb3e6ea0): complete subsection reference.

- [strict_schema_origin](data-sources--api_definition--reference--group-001.md#canonical-733352b38edeadf56977b1434d4b45e4b728f46901c205ea889ccc66a4ebe4df): complete subsection reference.

<a id="canonical-64c467600eecf22d457e9719d4e9c6eb8bd94340b6bd61d1a004636899cd6a08"></a>

<a id="canonical-0aa14d81a20bdfe98f42fba5409e6b5dc1ea513f0990e9bbb29e8c5bff411097"></a>

## swagger_specs property — Property reference / 533de0104972 / 10

Type: `["list", "string"]`. Computed.

URLs of versioned OpenAPI files uploaded through Web App &amp; API Protection &gt; Files &gt;
Swagger Files. The 512-byte item limit is a URL-length limit; inline string:/// OpenAPI content is
rejected and does not create an API-definition object. Defaults to \`\[\]\`. Server applies default
when omitted.

Upstream description:

URLs of versioned OpenAPI files uploaded through Web App &amp; API Protection &gt; Files &gt;
Swagger Files. The 512-byte item limit is a URL-length limit; inline string:/// OpenAPI content is
rejected and does not create an API-definition object.

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

<a id="canonical-e48119146d7f470cc5df9e2a23a57804852315cf083dc754eeb0710e6c3a9709"></a>

## All schema paths — Property reference / 533de0104972 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--api_definition--reference--group-001.md#canonical-a2877144ef18a48bad3fbdab5f358b9f041deccf3daa3cd10a42b75991c1e620) |
| `api_inventory_exclusion_list` | [api_inventory_exclusion_list](data-sources--api_definition--reference--group-001.md#canonical-d6664d9af71f1a2387368e3a60463237ecdd83a71337335e2b394b7e2746ba2d) |
| `api_inventory_exclusion_list.method` | [api_inventory_exclusion_list.method](data-sources--api_definition--reference--group-001.md#canonical-7d83d2cd3e604dc52c3524ecb53b24f87775de15c6a528fbf21dc02e55f30df4) |
| `api_inventory_exclusion_list.path` | [api_inventory_exclusion_list.path](data-sources--api_definition--reference--group-001.md#canonical-f8e239436a2c4076b9a83d62b712cd090da9abacfd1d6d6a0ea51839dcea1e1d) |
| `api_inventory_inclusion_list` | [api_inventory_inclusion_list](data-sources--api_definition--reference--group-001.md#canonical-b8460344e765e02bb2dd1a10ba69d9b077cc44a10b572641bb56760418c890d8) |
| `api_inventory_inclusion_list.method` | [api_inventory_inclusion_list.method](data-sources--api_definition--reference--group-001.md#canonical-903b4ba77fa4e2ab13883b9c84f9643f6cfaad94aa026d572d2cb5f10f71fc72) |
| `api_inventory_inclusion_list.path` | [api_inventory_inclusion_list.path](data-sources--api_definition--reference--group-001.md#canonical-32cd26bcb4ce88fdf886ed2f8cb7ff9448a0838c5f71246759720d99c2acf372) |
| `description` | [description](data-sources--api_definition--reference--group-001.md#canonical-759fa549790769fcbb51bd7506a9c7aebf14a37e6f370792f261e1e97ee242f4) |
| `id` | [id](data-sources--api_definition--reference--group-001.md#canonical-ee4d22dac33b484684b7b926cfcb0251983a52f08f3fc7d207b1588d0223a6a2) |
| `labels` | [labels](data-sources--api_definition--reference--group-001.md#canonical-991ff040fadacc440f5c60678943d13f071d5ff98cfdbf7c1a4280e258c7da7f) |
| `mixed_schema_origin` | [mixed_schema_origin](data-sources--api_definition--reference--group-001.md#canonical-f5983f439bcdb6ef100e1978e95ee7b19e45ed428ec3bd5d3642f957553e09a3) |
| `name` | [name](data-sources--api_definition--reference--group-001.md#canonical-5469b694f67f8b38e063ef9b87a5f80c1c56709aef64ea33e024a7f98a6c6b0a) |
| `namespace` | [namespace](data-sources--api_definition--reference--group-001.md#canonical-6edaf113889090c70298e2c7b45fc1978497805efdffb7b899bdc1b4965881b3) |
| `non_api_endpoints` | [non_api_endpoints](data-sources--api_definition--reference--group-001.md#canonical-434dc96ecc040d4e1226b4bdd21b6b13637284db42044ceb7313755ce7c2ee03) |
| `non_api_endpoints.method` | [non_api_endpoints.method](data-sources--api_definition--reference--group-001.md#canonical-b9bc6250d0ec73d2c6ed9ab962770e61248eb2992b59ea3536b200bd247eb1f7) |
| `non_api_endpoints.path` | [non_api_endpoints.path](data-sources--api_definition--reference--group-001.md#canonical-02bdb465d18b5e302941bc61d1425c1280fad698e5ef21fa4e9250ed8ca179fd) |
| `strict_schema_origin` | [strict_schema_origin](data-sources--api_definition--reference--group-001.md#canonical-150ea7f4e9f4c27dc877b5ee4591a95f9934d99e8fe230c46b5e295b74c7155d) |
| `swagger_specs` | [swagger_specs](data-sources--api_definition--reference--group-001.md#canonical-64c467600eecf22d457e9719d4e9c6eb8bd94340b6bd61d1a004636899cd6a08) |

<a id="canonical-2aff48777c6307b57414d9e3605119b7d0b8c5b2be6a07f090bfeaa22dd03ad0"></a>

## Next pages — Property reference / 533de0104972 / 12

- [api_inventory_exclusion_list](data-sources--api_definition--reference--group-001.md#canonical-2e076e32eb0f9baac97a955dba020f38ebfd3d65c8af0fb5c4736c2ac2b1f19c)
- [api_inventory_inclusion_list](data-sources--api_definition--reference--group-001.md#canonical-198d8acda02ecd2dd684e9389fc7d3a6b105de4e4510ce3873679235610d522c)
- [mixed_schema_origin](data-sources--api_definition--reference--group-001.md#canonical-57db1b0d2c8304075cf9831004863a1441d8563f5cb9e8f8e869a761bcebc086)
- [non_api_endpoints](data-sources--api_definition--reference--group-001.md#canonical-cb5c898f0a4624e8c07fe12db587917411cb88689982594ff5967457fb3e6ea0)
- [strict_schema_origin](data-sources--api_definition--reference--group-001.md#canonical-733352b38edeadf56977b1434d4b45e4b728f46901c205ea889ccc66a4ebe4df)
- [xcsh_api_definition](../data-sources/api_definition.md#canonical-1df393999e0c5facf84460c20dfc66c998ea2afc9c81bd04c6d3cf585cad53ab)

<a id="canonical-2e076e32eb0f9baac97a955dba020f38ebfd3d65c8af0fb5c4736c2ac2b1f19c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2a50e3a033fc54e9c21ad847461592dc85b1cf7a94eff5916863c6a9f973345c"></a>

## api_inventory_exclusion_list — api_inventory_exclusion_list / 12503ecbfe22 / 2

Breadcrumbs:

- [xcsh_api_definition](../data-sources/api_definition.md#canonical-1df393999e0c5facf84460c20dfc66c998ea2afc9c81bd04c6d3cf585cad53ab)
- [Property reference](data-sources--api_definition--reference--group-001.md#canonical-fad0050df3f4cd1ee16f15782b363b846ffd70034b8a18de354e7a072ec26d5b)
- api_inventory_exclusion_list

<a id="canonical-d6664d9af71f1a2387368e3a60463237ecdd83a71337335e2b394b7e2746ba2d"></a>

Type: `"list"`. Computed.

List of API Endpoints excluded from the API Inventory. Defaults to \`\[\]\`. Server applies default
when omitted.

Upstream description:

List of API Endpoints excluded from the API Inventory.

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

<a id="canonical-05a2a304cc7984e2271d7fce1aac5bd40c7882b0e490b050dca838dda30e6ba1"></a>

## Direct properties — api_inventory_exclusion_list / 12503ecbfe22 / 3

<a id="canonical-7d83d2cd3e604dc52c3524ecb53b24f87775de15c6a528fbf21dc02e55f30df4"></a>

<a id="canonical-03bbf5fdbd0e6cd8ac1198f94b24a2201f0a70998b8fb950e422ead5a88f61c3"></a>

## method property — api_inventory_exclusion_list / 12503ecbfe22 / 4

Type: `"string"`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

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

<a id="canonical-f8e239436a2c4076b9a83d62b712cd090da9abacfd1d6d6a0ea51839dcea1e1d"></a>

<a id="canonical-97c600d57338ff669f2db0879b0f6141bd071d2fd022d817a12bf68055f31eca"></a>

## path property — api_inventory_exclusion_list / 12503ecbfe22 / 5

Type: `"string"`. Computed.

Endpoint path, as specified in OpenAPI, including parameters. The path should comply with RFC 3986
and may have parameters according to OpenAPI specification.

Upstream description:

An endpoint path, as specified in OpenAPI, including parameters. The path should comply with RFC
3986 and may have parameters according to OpenAPI specification.

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

<a id="canonical-5ede358aa9b900e5a78132dbac7115cb1e526ad881290f30a679e3c6374477ec"></a>

## Next pages — api_inventory_exclusion_list / 12503ecbfe22 / 6

- [Property reference](data-sources--api_definition--reference--group-001.md#canonical-fad0050df3f4cd1ee16f15782b363b846ffd70034b8a18de354e7a072ec26d5b)
- [xcsh_api_definition](../data-sources/api_definition.md#canonical-1df393999e0c5facf84460c20dfc66c998ea2afc9c81bd04c6d3cf585cad53ab)

<a id="canonical-198d8acda02ecd2dd684e9389fc7d3a6b105de4e4510ce3873679235610d522c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a43693cf9f849779af49a04be614172fc9490ae3d195e1c5ba15618628d19158"></a>

## api_inventory_inclusion_list — api_inventory_inclusion_list / 402228b43297 / 2

Breadcrumbs:

- [xcsh_api_definition](../data-sources/api_definition.md#canonical-1df393999e0c5facf84460c20dfc66c998ea2afc9c81bd04c6d3cf585cad53ab)
- [Property reference](data-sources--api_definition--reference--group-001.md#canonical-fad0050df3f4cd1ee16f15782b363b846ffd70034b8a18de354e7a072ec26d5b)
- api_inventory_inclusion_list

<a id="canonical-b8460344e765e02bb2dd1a10ba69d9b077cc44a10b572641bb56760418c890d8"></a>

Type: `"list"`. Computed.

List of API Endpoints included in the API Inventory. Typically, discovered API endpoints are added
to the API Inventory using this list. Defaults to \`\[\]\`. Server applies default when omitted.

Upstream description:

List of API Endpoints included in the API Inventory. Typically, discovered API endpoints are added
to the API Inventory using this list.

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

<a id="canonical-fbf47bb75eee8edbfa677e7f9752bc3af3916fa601a7495bf2a6d0129d2c961d"></a>

## Direct properties — api_inventory_inclusion_list / 402228b43297 / 3

<a id="canonical-903b4ba77fa4e2ab13883b9c84f9643f6cfaad94aa026d572d2cb5f10f71fc72"></a>

<a id="canonical-540b9a0d1bb20d367b965af7737a15f23ab30385d38ef4aaafbbc54989438b1e"></a>

## method property — api_inventory_inclusion_list / 402228b43297 / 4

Type: `"string"`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

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

<a id="canonical-32cd26bcb4ce88fdf886ed2f8cb7ff9448a0838c5f71246759720d99c2acf372"></a>

<a id="canonical-d0388fdbb9c89757ff666b87f4ce68c87539b8bfdf48c7640900fce4a25626ca"></a>

## path property — api_inventory_inclusion_list / 402228b43297 / 5

Type: `"string"`. Computed.

Endpoint path, as specified in OpenAPI, including parameters. The path should comply with RFC 3986
and may have parameters according to OpenAPI specification.

Upstream description:

An endpoint path, as specified in OpenAPI, including parameters. The path should comply with RFC
3986 and may have parameters according to OpenAPI specification.

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

<a id="canonical-9a81b306956326a54a0a1a0461d4f9f5dbaa4e33b4aad2fd3d50084ce2572245"></a>

## Next pages — api_inventory_inclusion_list / 402228b43297 / 6

- [Property reference](data-sources--api_definition--reference--group-001.md#canonical-fad0050df3f4cd1ee16f15782b363b846ffd70034b8a18de354e7a072ec26d5b)
- [xcsh_api_definition](../data-sources/api_definition.md#canonical-1df393999e0c5facf84460c20dfc66c998ea2afc9c81bd04c6d3cf585cad53ab)

<a id="canonical-57db1b0d2c8304075cf9831004863a1441d8563f5cb9e8f8e869a761bcebc086"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4f8088062c9a7ab3d64a41d88455300971dcd51a672d37c10e039434dd5deb93"></a>

## mixed_schema_origin — mixed_schema_origin / 51092bedb660 / 2

Breadcrumbs:

- [xcsh_api_definition](../data-sources/api_definition.md#canonical-1df393999e0c5facf84460c20dfc66c998ea2afc9c81bd04c6d3cf585cad53ab)
- [Property reference](data-sources--api_definition--reference--group-001.md#canonical-fad0050df3f4cd1ee16f15782b363b846ffd70034b8a18de354e7a072ec26d5b)
- mixed_schema_origin

<a id="canonical-f5983f439bcdb6ef100e1978e95ee7b19e45ed428ec3bd5d3642f957553e09a3"></a>

Type: `["object", {}]`. Computed.

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

- [mixed_schema_origin](data-sources--api_definition--reference--group-001.md#canonical-f5983f439bcdb6ef100e1978e95ee7b19e45ed428ec3bd5d3642f957553e09a3)
- [strict_schema_origin](data-sources--api_definition--reference--group-001.md#canonical-150ea7f4e9f4c27dc877b5ee4591a95f9934d99e8fe230c46b5e295b74c7155d)

Select alternatives according to the provider validators above.

<a id="canonical-0540d846db9f9b52a131b62fc4bec8ffc63628de9a1dbe021f6196f387eccb94"></a>

## Direct properties — mixed_schema_origin / 51092bedb660 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-89704d3463698251308461264cfc4dfee6dae314a211b88f39341f6fa377a061"></a>

## Next pages — mixed_schema_origin / 51092bedb660 / 4

- [Property reference](data-sources--api_definition--reference--group-001.md#canonical-fad0050df3f4cd1ee16f15782b363b846ffd70034b8a18de354e7a072ec26d5b)
- [xcsh_api_definition](../data-sources/api_definition.md#canonical-1df393999e0c5facf84460c20dfc66c998ea2afc9c81bd04c6d3cf585cad53ab)

<a id="canonical-cb5c898f0a4624e8c07fe12db587917411cb88689982594ff5967457fb3e6ea0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9b571aee653772db6de75c8d3b4547a41d9225b864a39387af04ae528a8227b0"></a>

## non_api_endpoints — non_api_endpoints / be9a1f4deba3 / 2

Breadcrumbs:

- [xcsh_api_definition](../data-sources/api_definition.md#canonical-1df393999e0c5facf84460c20dfc66c998ea2afc9c81bd04c6d3cf585cad53ab)
- [Property reference](data-sources--api_definition--reference--group-001.md#canonical-fad0050df3f4cd1ee16f15782b363b846ffd70034b8a18de354e7a072ec26d5b)
- non_api_endpoints

<a id="canonical-434dc96ecc040d4e1226b4bdd21b6b13637284db42044ceb7313755ce7c2ee03"></a>

Type: `"list"`. Computed.

API Discovery Exclusion List. List of Non-API Endpoints. Defaults to \`\[\]\`. Server applies
default when omitted.

Upstream description:

List of Non-API Endpoints.

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

<a id="canonical-1bd3ac5986630b4cc8aba3524bb9b5cd9865c826f73735fc361246a172466561"></a>

## Direct properties — non_api_endpoints / be9a1f4deba3 / 3

<a id="canonical-b9bc6250d0ec73d2c6ed9ab962770e61248eb2992b59ea3536b200bd247eb1f7"></a>

<a id="canonical-31ee14d40e8d76237fe1938245d9aa5364ddd079cad79501fe5b3d823f2668af"></a>

## method property — non_api_endpoints / be9a1f4deba3 / 4

Type: `"string"`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

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

<a id="canonical-02bdb465d18b5e302941bc61d1425c1280fad698e5ef21fa4e9250ed8ca179fd"></a>

<a id="canonical-7b5bd21b34e9682a233fbea63e89f9c1526bb50ef82ae648afc24ca260574c0f"></a>

## path property — non_api_endpoints / be9a1f4deba3 / 5

Type: `"string"`. Computed.

Endpoint path, as specified in OpenAPI, including parameters. The path should comply with RFC 3986
and may have parameters according to OpenAPI specification.

Upstream description:

An endpoint path, as specified in OpenAPI, including parameters. The path should comply with RFC
3986 and may have parameters according to OpenAPI specification.

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

<a id="canonical-7a3e158cd1f5ee4dc104c7c99606feaff7382c8fdd4c4aa53626df847d7241eb"></a>

## Next pages — non_api_endpoints / be9a1f4deba3 / 6

- [Property reference](data-sources--api_definition--reference--group-001.md#canonical-fad0050df3f4cd1ee16f15782b363b846ffd70034b8a18de354e7a072ec26d5b)
- [xcsh_api_definition](../data-sources/api_definition.md#canonical-1df393999e0c5facf84460c20dfc66c998ea2afc9c81bd04c6d3cf585cad53ab)

<a id="canonical-733352b38edeadf56977b1434d4b45e4b728f46901c205ea889ccc66a4ebe4df"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-57fff8537211088cb53a12d8b591aa06bd87fa59a8918cee3d255336651fd5e4"></a>

## strict_schema_origin — strict_schema_origin / f7954589ef6b / 2

Breadcrumbs:

- [xcsh_api_definition](../data-sources/api_definition.md#canonical-1df393999e0c5facf84460c20dfc66c998ea2afc9c81bd04c6d3cf585cad53ab)
- [Property reference](data-sources--api_definition--reference--group-001.md#canonical-fad0050df3f4cd1ee16f15782b363b846ffd70034b8a18de354e7a072ec26d5b)
- strict_schema_origin

<a id="canonical-150ea7f4e9f4c27dc877b5ee4591a95f9934d99e8fe230c46b5e295b74c7155d"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-493d01e0d7d7c4f3a71cd738d7aefae6de71ccc342eebe46aae575771df23b21"></a>

## Direct properties — strict_schema_origin / f7954589ef6b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f79b1c72deb93c08cda99ad9c9ff81f05dbcab67e65d5a73d8883f10046c46bc"></a>

## Next pages — strict_schema_origin / f7954589ef6b / 4

- [Property reference](data-sources--api_definition--reference--group-001.md#canonical-fad0050df3f4cd1ee16f15782b363b846ffd70034b8a18de354e7a072ec26d5b)
- [xcsh_api_definition](../data-sources/api_definition.md#canonical-1df393999e0c5facf84460c20dfc66c998ea2afc9c81bd04c6d3cf585cad53ab)
