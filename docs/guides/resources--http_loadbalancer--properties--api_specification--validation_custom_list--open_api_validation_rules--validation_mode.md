---
page_title: "api_specification.validation_custom_list.open_api_validation_rules.validation_mode"
subcategory: "Load Balancing"
description: "api_specification.validation_custom_list.open_api_validation_rules.validation_mode for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 4186, "body_sha256": "sha256:537c55964149a78a403563d49cfa55d10f7aba6a5c218dc40d2d41a4ffe2a95a", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active", "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:skip_response_validation", "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:skip_validation", "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:validation_mode_active"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules", "path": "docs/guides/resources--http_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification.validation_custom_list.open_api_validation_rules.validation_mode for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# api_specification.validation_custom_list.open_api_validation_rules.validation_mode

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [api_specification](resources--http_loadbalancer--properties--api_specification.md)
- [api_specification.validation_custom_list](resources--http_loadbalancer--properties--api_specification--validation_custom_list.md)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules.md)
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

- [response_validation_mode_active](resources--http_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--response_validation_mode_active.md): complete subsection reference.

- [skip_response_validation](resources--http_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--skip_response_validation.md): complete subsection reference.

- [skip_validation](resources--http_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--skip_validation.md): complete subsection reference.

- [validation_mode_active](resources--http_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--validation_mode_active.md): complete subsection reference.

## Next pages

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](resources--http_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--response_validation_mode_active.md)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation](resources--http_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--skip_response_validation.md)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_validation](resources--http_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--skip_validation.md)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](resources--http_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--validation_mode_active.md)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
