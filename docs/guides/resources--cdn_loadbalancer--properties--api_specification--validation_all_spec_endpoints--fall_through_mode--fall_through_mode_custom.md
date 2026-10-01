---
page_title: "api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom"
subcategory: "Load Balancing"
description: "api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2273, "body_sha256": "sha256:4dd62f5d8537186bd24681612dffb063416f2cd250eb80ef287a632bd83ce665", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode:fall_through_mode_custom", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode:fall_through_mode_custom:open_api_validation_rules"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode:fall_through_mode_custom", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode", "path": "docs/guides/resources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints", "fall_through_mode", "fall_through_mode_custom"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/fall_through_mode/fall_through_mode_custom/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [api_specification](resources--cdn_loadbalancer--properties--api_specification.md)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints.md)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--fall_through_mode.md)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for fall through mode custom.

Upstream description:

Define the fall through settings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("open_api_validation_rules")}
```

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
fall_through_mode_custom {
  # Configure direct properties listed below.
}
```

## Direct properties

- [open_api_validation_rules](resources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom--open_api_validation_rules.md): complete subsection reference.

## Next pages

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom--open_api_validation_rules.md)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--fall_through_mode.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
