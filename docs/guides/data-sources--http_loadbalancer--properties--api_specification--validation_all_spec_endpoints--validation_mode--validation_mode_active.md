---
page_title: "api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active"
subcategory: "Load Balancing"
description: "api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 4693, "body_sha256": "sha256:22fb601a06e42ec89580a00c02e1a3a94eb2be1d7f4cf061177025c6fade6ef4", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:validation_mode_active", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:validation_mode_active:enforcement_block", "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:validation_mode_active:enforcement_report"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:validation_mode_active", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode", "path": "docs/guides/data-sources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode--validation_mode_active.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints", "validation_mode", "validation_mode_active"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/validation_mode_active/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [api_specification](data-sources--http_loadbalancer--properties--api_specification.md)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints.md)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode.md)
- api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active

<a id="section"></a>

Type: `"single"`. Computed.

Enable OpenAPI validation and explicitly select enforcement\_report to allow and log invalid
traffic, or enforcement\_block to reject invalid requests with HTTP 403.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-validation_enforcement_type": "[\"enforcement_block\",\"enforcement_report\"]"
}
```

## Direct properties

- [enforcement_block](data-sources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode--validation_mode_active--enforcement_block.md): complete subsection reference.

- [enforcement_report](data-sources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode--validation_mode_active--enforcement_report.md): complete subsection reference.

<a id="schema-api_specification--validation_all_spec_endpoints--validation_mode--validation_mode_active--request_validation_properties"></a>

### request_validation_properties property

Type: `["list", "string"]`. Computed.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the request to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

Upstream description:

List of properties of the request to validate according to the OpenAPI specification file (a.k.a.
Swagger)

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.enum.not_in": "[7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block](data-sources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode--validation_mode_active--enforcement_block.md)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_report](data-sources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode--validation_mode_active--enforcement_report.md)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
