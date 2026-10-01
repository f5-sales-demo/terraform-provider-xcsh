---
page_title: "api_specification.validation_all_spec_endpoints"
subcategory: "Load Balancing"
description: "api_specification.validation_all_spec_endpoints for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2163, "body_sha256": "sha256:761eb4b9ce6f19736ce072ed922e6392efd3f91b63a966200c430e0b91f359a9", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode", "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings", "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification", "path": "docs/guides/resources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification.validation_all_spec_endpoints for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_all_spec_endpoints

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [api_specification](resources--http_loadbalancer--properties--api_specification.md)
- api_specification.validation_all_spec_endpoints

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
validation_all_spec_endpoints {
  # Configure direct properties listed below.
}
```

## Direct properties

- [fall_through_mode](resources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--fall_through_mode.md): complete subsection reference.

- [settings](resources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--settings.md): complete subsection reference.

- [validation_mode](resources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode.md): complete subsection reference.

## Next pages

- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--fall_through_mode.md)
- [api_specification.validation_all_spec_endpoints.settings](resources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--settings.md)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode.md)
- [api_specification](resources--http_loadbalancer--properties--api_specification.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
