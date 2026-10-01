---
page_title: "api_specification.validation_custom_list.fall_through_mode"
subcategory: "Load Balancing"
description: "api_specification.validation_custom_list.fall_through_mode for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2343, "body_sha256": "sha256:4273a011704bf47d9f0f30414b8eda717cf5101731fef4247c17ad4eeef61723", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_allow", "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list", "path": "docs/guides/data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--fall_through_mode.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "fall_through_mode"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification.validation_custom_list.fall_through_mode for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_custom_list.fall_through_mode

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [api_specification](data-sources--http_loadbalancer--properties--api_specification.md)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list.md)
- api_specification.validation_custom_list.fall_through_mode

<a id="section"></a>

Type: `"single"`. Computed.

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules).

Upstream description:

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-fall_through_mode_choice": "[\"fall_through_mode_allow\",\"fall_through_mode_custom\"]"
}
```

## Direct properties

- [fall_through_mode_allow](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_allow.md): complete subsection reference.

- [fall_through_mode_custom](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom.md): complete subsection reference.

## Next pages

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_allow](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_allow.md)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom.md)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
