---
page_title: "api_specification.validation_all_spec_endpoints.fall_through_mode"
subcategory: "Load Balancing"
description: "api_specification.validation_all_spec_endpoints.fall_through_mode for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2708, "body_sha256": "sha256:135c4de0d6fe5454755f377bc86f157ccfbd078e30162408474cf3d117965ac9", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode:fall_through_mode_allow", "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode:fall_through_mode_custom"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints", "path": "docs/guides/resources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--fall_through_mode.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints", "fall_through_mode"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/fall_through_mode/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification.validation_all_spec_endpoints.fall_through_mode for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_all_spec_endpoints.fall_through_mode

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [api_specification](resources--cdn_loadbalancer--properties--api_specification.md)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints.md)
- api_specification.validation_all_spec_endpoints.fall_through_mode

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules).

Upstream description:

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("fall_through_mode_allow",
    "fall_through_mode_custom")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-fall_through_mode_choice": "[\"fall_through_mode_allow\",\"fall_through_mode_custom\"]"
}
```

Terraform syntax:

```terraform
fall_through_mode {
  # Configure direct properties listed below.
}
```

## Direct properties

- [fall_through_mode_allow](resources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_allow.md): complete subsection reference.

- [fall_through_mode_custom](resources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom.md): complete subsection reference.

## Next pages

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow](resources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_allow.md)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom.md)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
