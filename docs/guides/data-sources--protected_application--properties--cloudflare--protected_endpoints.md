---
page_title: "cloudflare.protected_endpoints"
subcategory: ""
description: "cloudflare.protected_endpoints for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 6494, "body_sha256": "sha256:b1a49e51e080c439a33ad62e1b42e2740b59796e0b1a4231e625cbb9faaad834", "canonical_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:any_domain", "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:domain", "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:metadata", "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:mobile_client", "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:path", "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_client", "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client"], "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare", "path": "docs/guides/data-sources--protected_application--properties--cloudflare--protected_endpoints.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudflare", "protected_endpoints"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudflare/protected_endpoints/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudflare.protected_endpoints for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudflare.protected_endpoints

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md)
- [Property reference](data-sources--protected_application--reference.md)
- [cloudflare](data-sources--protected_application--properties--cloudflare.md)
- cloudflare.protected_endpoints

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

- [any_domain](data-sources--protected_application--properties--cloudflare--protected_endpoints--any_domain.md): complete subsection reference.

- [domain](data-sources--protected_application--properties--cloudflare--protected_endpoints--domain.md): complete subsection reference.

<a id="schema-cloudflare--protected_endpoints--http_methods"></a>

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

- [metadata](data-sources--protected_application--properties--cloudflare--protected_endpoints--metadata.md): complete subsection reference.

- [mobile_client](data-sources--protected_application--properties--cloudflare--protected_endpoints--mobile_client.md): complete subsection reference.

- [path](data-sources--protected_application--properties--cloudflare--protected_endpoints--path.md): complete subsection reference.

<a id="schema-cloudflare--protected_endpoints--query"></a>

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

- [web_client](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_client.md): complete subsection reference.

- [web_mobile_client](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client.md): complete subsection reference.

## Next pages

- [cloudflare.protected_endpoints.any_domain](data-sources--protected_application--properties--cloudflare--protected_endpoints--any_domain.md)
- [cloudflare.protected_endpoints.domain](data-sources--protected_application--properties--cloudflare--protected_endpoints--domain.md)
- [cloudflare.protected_endpoints.metadata](data-sources--protected_application--properties--cloudflare--protected_endpoints--metadata.md)
- [cloudflare.protected_endpoints.mobile_client](data-sources--protected_application--properties--cloudflare--protected_endpoints--mobile_client.md)
- [cloudflare.protected_endpoints.path](data-sources--protected_application--properties--cloudflare--protected_endpoints--path.md)
- [cloudflare.protected_endpoints.web_client](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_client.md)
- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client.md)
- [cloudflare](data-sources--protected_application--properties--cloudflare.md)
- [xcsh_protected_application](../data-sources/protected_application.md)
