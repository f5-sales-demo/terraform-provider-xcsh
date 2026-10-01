---
page_title: "api_specification.validation_custom_list.open_api_validation_rules.validation_mode"
subcategory: "Load Balancing"
description: "api_specification.validation_custom_list.open_api_validation_rules.validation_mode for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 4268, "body_sha256": "sha256:a8ff66d67f1d8c0f9025ae0098a14f417d9b4cf0516a2cbf59cc2a780fff9022", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active", "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:skip_response_validation", "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:skip_validation", "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:validation_mode_active"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules", "path": "docs/guides/resources--cdn_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification.validation_custom_list.open_api_validation_rules.validation_mode for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_custom_list.open_api_validation_rules.validation_mode

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [api_specification](resources--cdn_loadbalancer--properties--api_specification.md)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--properties--api_specification--validation_custom_list.md)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules.md)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger).

Upstream description:

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("response_validation_mode_active",
    "skip_response_validation"),
  validators.ConflictingObjectAttributes("skip_validation",
    "validation_mode_active")}
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
  "x-ves-oneof-field-response_validation_mode_choice": "[\"response_validation_mode_active\",\"skip_response_validation\"]",
  "x-ves-oneof-field-validation_mode_choice": "[\"skip_validation\",\"validation_mode_active\"]"
}
```

Terraform syntax:

```terraform
validation_mode {
  # Configure direct properties listed below.
}
```

## Direct properties

- [response_validation_mode_active](resources--cdn_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--response_validation_mode_active.md): complete subsection reference.

- [skip_response_validation](resources--cdn_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--skip_response_validation.md): complete subsection reference.

- [skip_validation](resources--cdn_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--skip_validation.md): complete subsection reference.

- [validation_mode_active](resources--cdn_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--validation_mode_active.md): complete subsection reference.

## Next pages

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](resources--cdn_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--response_validation_mode_active.md)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation](resources--cdn_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--skip_response_validation.md)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_validation](resources--cdn_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--skip_validation.md)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](resources--cdn_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--validation_mode_active.md)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
