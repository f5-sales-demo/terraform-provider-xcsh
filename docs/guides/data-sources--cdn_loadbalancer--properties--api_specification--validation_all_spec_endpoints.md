---
page_title: "api_specification.validation_all_spec_endpoints"
subcategory: "Load Balancing"
description: "api_specification.validation_all_spec_endpoints for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2045, "body_sha256": "sha256:86858f74e9305df4b8c8d7a7113b13903ed675f21a0860fcfff8394039235822", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification.validation_all_spec_endpoints for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_all_spec_endpoints

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [api_specification](data-sources--cdn_loadbalancer--properties--api_specification.md)
- api_specification.validation_all_spec_endpoints

<a id="section"></a>

Type: `"single"`. Computed.

API Inventory. Settings for API Inventory validation.

Upstream description:

Settings for API Inventory validation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-oversized_body_choice": "[]"
}
```

## Direct properties

- [fall_through_mode](data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--fall_through_mode.md): complete subsection reference.

- [settings](data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--settings.md): complete subsection reference.

- [validation_mode](data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode.md): complete subsection reference.

## Next pages

- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--fall_through_mode.md)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--settings.md)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode.md)
- [api_specification](data-sources--cdn_loadbalancer--properties--api_specification.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
