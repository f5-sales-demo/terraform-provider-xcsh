---
page_title: "api_specification"
subcategory: "Load Balancing"
description: "api_specification for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2533, "body_sha256": "sha256:b9f2d85e894ede5f793a32d88c25bd238813843bd4a36108748ce00d4d03ca85", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:api_definition", "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints", "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list", "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_disabled"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "docs/guides/data-sources--http_loadbalancer--properties--api_specification.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_specification/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- api_specification

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: api\_specification, disable\_api\_definition; Default: disable\_api\_definition\] Settings
for API specification (API definition, OpenAPI validation, etc.).

Upstream description:

Settings for API specification (API definition, OpenAPI validation, etc.)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-validation_target_choice": "[\"validation_all_spec_endpoints\",\"validation_custom_list\",\"validation_disabled\"]"
}
```

OneOf alternatives in this subsection:

- [api_specification](data-sources--http_loadbalancer--properties--api_specification.md#section)
- [disable_api_definition](data-sources--http_loadbalancer--properties--disable_api_definition.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [api_definition](data-sources--http_loadbalancer--properties--api_specification--api_definition.md): complete subsection reference.

- [validation_all_spec_endpoints](data-sources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints.md): complete subsection reference.

- [validation_custom_list](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list.md): complete subsection reference.

- [validation_disabled](data-sources--http_loadbalancer--properties--api_specification--validation_disabled.md): complete subsection reference.

## Next pages

- [api_specification.api_definition](data-sources--http_loadbalancer--properties--api_specification--api_definition.md)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints.md)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list.md)
- [api_specification.validation_disabled](data-sources--http_loadbalancer--properties--api_specification--validation_disabled.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
