---
page_title: "api_specification.validation_custom_list"
subcategory: "Load Balancing"
description: "api_specification.validation_custom_list for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2138, "body_sha256": "sha256:946098831fe7a1d9edbaa58cb40df6cc5435f9b4cce5e7992a919e31e11d552d", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:settings"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--api_specification--validation_custom_list.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification", "validation_custom_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification.validation_custom_list for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# api_specification.validation_custom_list

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [api_specification](data-sources--cdn_loadbalancer--properties--api_specification.md)
- api_specification.validation_custom_list

<a id="section"></a>

Type: `"single"`. Computed.

Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other
API-endpoint not listed will act according to 'Fall Through Mode'.

Upstream description:

Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other
API-endpoint not listed will act according to "Fall Through Mode".

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

- [fall_through_mode](data-sources--cdn_loadbalancer--properties--api_specification--validation_custom_list--fall_through_mode.md): complete subsection reference.

- [open_api_validation_rules](data-sources--cdn_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules.md): complete subsection reference.

- [settings](data-sources--cdn_loadbalancer--properties--api_specification--validation_custom_list--settings.md): complete subsection reference.

## Next pages

- [api_specification.validation_custom_list.fall_through_mode](data-sources--cdn_loadbalancer--properties--api_specification--validation_custom_list--fall_through_mode.md)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--cdn_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules.md)
- [api_specification.validation_custom_list.settings](data-sources--cdn_loadbalancer--properties--api_specification--validation_custom_list--settings.md)
- [api_specification](data-sources--cdn_loadbalancer--properties--api_specification.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
