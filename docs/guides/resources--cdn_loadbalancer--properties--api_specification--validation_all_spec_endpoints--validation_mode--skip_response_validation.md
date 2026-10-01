---
page_title: "api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation"
subcategory: "Load Balancing"
description: "api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1582, "body_sha256": "sha256:f16c9f154ba15b97a866c8a966caa7fc9b6b74e6b879f3748ce9dec3ade7aa24", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:skip_response_validation", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:skip_response_validation", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode", "path": "docs/guides/resources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode--skip_response_validation.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints", "validation_mode", "skip_response_validation"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/skip_response_validation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [api_specification](resources--cdn_loadbalancer--properties--api_specification.md)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints.md)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode.md)
- api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

Terraform syntax:

```terraform
skip_response_validation = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
