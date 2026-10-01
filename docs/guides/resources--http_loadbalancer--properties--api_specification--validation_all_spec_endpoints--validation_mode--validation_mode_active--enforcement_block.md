---
page_title: "api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block"
subcategory: "Load Balancing"
description: "api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1926, "body_sha256": "sha256:7f20cbe1e2402b6d21e3bf9525b9c858934a64850667ceb62f25ac86b4cb4a23", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:validation_mode_active:enforcement_block", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:validation_mode_active:enforcement_block", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:validation_mode_active", "path": "docs/guides/resources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode--validation_mode_active--enforcement_block.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints", "validation_mode", "validation_mode_active", "enforcement_block"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/validation_mode_active/enforcement_block/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [api_specification](resources--http_loadbalancer--properties--api_specification.md)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints.md)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode.md)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](resources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode--validation_mode_active.md)
- api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block

<a id="section"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
enforcement_block = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](resources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--validation_mode--validation_mode_active.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
