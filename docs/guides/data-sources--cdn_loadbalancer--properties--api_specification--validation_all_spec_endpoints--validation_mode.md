---
page_title: "api_specification.validation_all_spec_endpoints.validation_mode"
subcategory: "Load Balancing"
description: "api_specification.validation_all_spec_endpoints.validation_mode for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3307, "body_sha256": "sha256:edc8deb9a1677314f27fcab5f635ba52af9a33fde506b8053de873f125d7dc41", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:response_validation_mode_active", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:skip_response_validation", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:skip_validation", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:validation_mode_active"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints", "validation_mode"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification.validation_all_spec_endpoints.validation_mode for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# api_specification.validation_all_spec_endpoints.validation_mode

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [api_specification](data-sources--cdn_loadbalancer--properties--api_specification.md)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints.md)
- api_specification.validation_all_spec_endpoints.validation_mode

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

- [response_validation_mode_active](data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode--response_validation_mode_active.md): complete subsection reference.

- [skip_response_validation](data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode--skip_response_validation.md): complete subsection reference.

- [skip_validation](data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode--skip_validation.md): complete subsection reference.

- [validation_mode_active](data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode--validation_mode_active.md): complete subsection reference.

## Next pages

- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode--response_validation_mode_active.md)
- [api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation](data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode--skip_response_validation.md)
- [api_specification.validation_all_spec_endpoints.validation_mode.skip_validation](data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode--skip_validation.md)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode--validation_mode_active.md)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
