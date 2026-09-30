---
page_title: "api_specification.validation_all_spec_endpoints"
subcategory: "Load Balancing"
description: "api_specification.validation_all_spec_endpoints for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1959, "body_sha256": "sha256:f5eb5ef64025241ffbd2d914d45aa1c702f40944eb9e5f07a8f05b7cf86192af", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode", "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings", "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification", "path": "docs/guides/data-sources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification.validation_all_spec_endpoints for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# api_specification.validation_all_spec_endpoints

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [api_specification](data-sources--http_loadbalancer--properties--api_specification.md)
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

- [fall_through_mode](data-sources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--fall_through_mode.md): complete subsection reference.

- [settings](data-sources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--settings.md): complete subsection reference.

- [validation_mode](data-sources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode.md): complete subsection reference.

## Next pages

- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--fall_through_mode.md)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--settings.md)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode.md)
- [api_specification](data-sources--http_loadbalancer--properties--api_specification.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
