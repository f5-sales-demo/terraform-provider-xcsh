---
page_title: "api_specification.validation_custom_list.open_api_validation_rules.validation_mode"
subcategory: "Load Balancing"
description: "api_specification.validation_custom_list.open_api_validation_rules.validation_mode for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3794, "body_sha256": "sha256:9842dba9029a757210a447c529b3702a4a56b17a7e6f562574dffc4793503bb4", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:skip_response_validation", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:skip_validation", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:validation_mode_active"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification.validation_custom_list.open_api_validation_rules.validation_mode for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# api_specification.validation_custom_list.open_api_validation_rules.validation_mode

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [api_specification](data-sources--cdn_loadbalancer--properties--api_specification.md)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--properties--api_specification--validation_custom_list.md)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--cdn_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules.md)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode

<a id="section"></a>

Type: `"single"`. Computed.

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger).

Upstream description:

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-response_validation_mode_choice": "[\"response_validation_mode_active\",\"skip_response_validation\"]",
  "x-ves-oneof-field-validation_mode_choice": "[\"skip_validation\",\"validation_mode_active\"]"
}
```

## Direct properties

- [response_validation_mode_active](data-sources--cdn_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--response_validation_mode_active.md): complete subsection reference.

- [skip_response_validation](data-sources--cdn_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--skip_response_validation.md): complete subsection reference.

- [skip_validation](data-sources--cdn_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--skip_validation.md): complete subsection reference.

- [validation_mode_active](data-sources--cdn_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--validation_mode_active.md): complete subsection reference.

## Next pages

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](data-sources--cdn_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--response_validation_mode_active.md)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation](data-sources--cdn_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--skip_response_validation.md)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_validation](data-sources--cdn_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--skip_validation.md)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](data-sources--cdn_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--validation_mode_active.md)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--cdn_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
