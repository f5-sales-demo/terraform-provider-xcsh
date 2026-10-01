---
page_title: "api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block"
subcategory: "Load Balancing"
description: "api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1934, "body_sha256": "sha256:2e25ebbca21c5f1df7579d9b9b3db7134ecf07d6b57e86689f84557b9f575cb3", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:response_validation_mode_active:enforcement_block", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:response_validation_mode_active:enforcement_block", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:response_validation_mode_active", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode--response_validation_mode_active--enforcement_block.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints", "validation_mode", "response_validation_mode_active", "enforcement_block"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/response_validation_mode_active/enforcement_block/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [api_specification](data-sources--cdn_loadbalancer--properties--api_specification.md)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints.md)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode.md)
- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode--response_validation_mode_active.md)
- api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block

<a id="section"></a>

Type: `["object", {}]`. Computed.

Blocking validation: reject traffic that violates the selected OpenAPI validation properties.
Invalid requests are returned as HTTP 403.

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

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode--response_validation_mode_active.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
