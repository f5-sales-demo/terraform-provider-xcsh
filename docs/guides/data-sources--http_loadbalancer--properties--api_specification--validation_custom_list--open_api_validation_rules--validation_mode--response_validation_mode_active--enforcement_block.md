---
page_title: "api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_block"
subcategory: "Load Balancing"
description: "api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_block for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2175, "body_sha256": "sha256:df7acb12745db46278e618c8e76d8be7bb4919b1dedd08811965b83c991d7488", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active:enforcement_block", "child_ids": [], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active:enforcement_block", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active", "path": "docs/guides/data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--response_validation_mode_active--enforcement_block.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode", "response_validation_mode_active", "enforcement_block"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/response_validation_mode_active/enforcement_block/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_block for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_block

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [api_specification](data-sources--http_loadbalancer--properties--api_specification.md)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list.md)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules.md)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode.md)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--response_validation_mode_active.md)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_block

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

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--response_validation_mode_active.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
