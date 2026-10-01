---
page_title: "cloudfront.protected_endpoints"
subcategory: ""
description: "cloudfront.protected_endpoints for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 8200, "body_sha256": "sha256:27e7985c7235a552f0c2d797fcddbab48a3eadc6b4b5d987193891de8004d617", "canonical_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:any_domain", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:domain", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:metadata", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:mobile_client", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:undefined_flow_label", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_client", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client"], "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront", "path": "docs/guides/data-sources--protected_application--properties--cloudfront--protected_endpoints.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudfront", "protected_endpoints"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudfront/protected_endpoints/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudfront.protected_endpoints for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.protected_endpoints

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md)
- [Property reference](data-sources--protected_application--reference.md)
- [cloudfront](data-sources--protected_application--properties--cloudfront.md)
- cloudfront.protected_endpoints

<a id="section"></a>

Type: `"list"`. Computed.

List of protected endpoints (max 128 items).

Upstream description:

List of protected endpoints (max 128 items)

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [any_domain](data-sources--protected_application--properties--cloudfront--protected_endpoints--any_domain.md): complete subsection reference.

- [domain](data-sources--protected_application--properties--cloudfront--protected_endpoints--domain.md): complete subsection reference.

- [flow_label](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label.md): complete subsection reference.

<a id="schema-cloudfront--protected_endpoints--http_methods"></a>

### http_methods property

Type: `["list", "string"]`. Computed.

\[Enum:
METHOD\_ANY|METHOD\_GET|METHOD\_POST|METHOD\_PUT|METHOD\_PATCH|METHOD\_DELETE|METHOD\_GET\_DOCUMENT\]
HTTP Methods. List of HTTP methods. Possible values are \`METHOD\_ANY\`, \`METHOD\_GET\`,
\`METHOD\_POST\`, \`METHOD\_PUT\`, \`METHOD\_PATCH\`, \`METHOD\_DELETE\`, \`METHOD\_GET\_DOCUMENT\`.
Defaults to \`METHOD\_ANY\`.

Upstream description:

List of HTTP methods.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[1,3,4]",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[1,3,4]",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [metadata](data-sources--protected_application--properties--cloudfront--protected_endpoints--metadata.md): complete subsection reference.

- [mobile_client](data-sources--protected_application--properties--cloudfront--protected_endpoints--mobile_client.md): complete subsection reference.

<a id="schema-cloudfront--protected_endpoints--path"></a>

### path property

Type: `"string"`. Computed.

Accepts wildcards \* to match multiple characters or ? To match a single character.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
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
    "pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,999}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,999}$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,999}$"
  }
}
```

<a id="schema-cloudfront--protected_endpoints--query"></a>

### query property

Type: `"string"`. Computed.

Enter a regular expression to match your query parameters of interest.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

- [undefined_flow_label](data-sources--protected_application--properties--cloudfront--protected_endpoints--undefined_flow_label.md): complete subsection reference.

- [web_client](data-sources--protected_application--properties--cloudfront--protected_endpoints--web_client.md): complete subsection reference.

- [web_mobile_client](data-sources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client.md): complete subsection reference.

## Next pages

- [cloudfront.protected_endpoints.any_domain](data-sources--protected_application--properties--cloudfront--protected_endpoints--any_domain.md)
- [cloudfront.protected_endpoints.domain](data-sources--protected_application--properties--cloudfront--protected_endpoints--domain.md)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label.md)
- [cloudfront.protected_endpoints.metadata](data-sources--protected_application--properties--cloudfront--protected_endpoints--metadata.md)
- [cloudfront.protected_endpoints.mobile_client](data-sources--protected_application--properties--cloudfront--protected_endpoints--mobile_client.md)
- [cloudfront.protected_endpoints.undefined_flow_label](data-sources--protected_application--properties--cloudfront--protected_endpoints--undefined_flow_label.md)
- [cloudfront.protected_endpoints.web_client](data-sources--protected_application--properties--cloudfront--protected_endpoints--web_client.md)
- [cloudfront.protected_endpoints.web_mobile_client](data-sources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client.md)
- [cloudfront](data-sources--protected_application--properties--cloudfront.md)
- [xcsh_protected_application](../data-sources/protected_application.md)
