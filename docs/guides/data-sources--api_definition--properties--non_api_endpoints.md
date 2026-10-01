---
page_title: "non_api_endpoints"
subcategory: "API Management"
description: "non_api_endpoints for xcsh_api_definition."
xcsh_docs: {"aliases": [], "body_bytes": 3799, "body_sha256": "sha256:fce3aff1547007219599b8b1aea87b684ed34baccf0b8c283b37a98bb058751b", "canonical_id": "xcsh-docs:data-sources:api_definition:properties:non_api_endpoints", "child_ids": [], "collection_id": "xcsh-docs:data-sources:api_definition:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_definition:properties:non_api_endpoints", "parent_id": "xcsh-docs:data-sources:api_definition:reference", "path": "docs/guides/data-sources--api_definition--properties--non_api_endpoints.md", "provider_name": "api_definition", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["non_api_endpoints"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_definition/properties/non_api_endpoints/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "non_api_endpoints for xcsh_api_definition.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_definitionCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# non_api_endpoints

Breadcrumbs:

- [xcsh_api_definition](../data-sources/api_definition.md)
- [Property reference](data-sources--api_definition--reference.md)
- non_api_endpoints

<a id="section"></a>

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

## Direct properties

<a id="schema-non_api_endpoints--method"></a>

### method property

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

<a id="schema-non_api_endpoints--path"></a>

### path property

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

## Next pages

- [Property reference](data-sources--api_definition--reference.md)
- [xcsh_api_definition](../data-sources/api_definition.md)
