---
page_title: "api_specification.validation_custom_list"
subcategory: "Load Balancing"
description: "api_specification.validation_custom_list for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2512, "body_sha256": "sha256:fed37e6b9adffb3bb7b23dbc6031e6aa1d23d41e8a6b54afdc69b74730b8f19f", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode", "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules", "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification", "path": "docs/guides/resources--http_loadbalancer--properties--api_specification--validation_custom_list.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification", "validation_custom_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_specification/validation_custom_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification.validation_custom_list for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_custom_list

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [api_specification](resources--http_loadbalancer--properties--api_specification.md)
- api_specification.validation_custom_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other
API-endpoint not listed will act according to 'Fall Through Mode'.

Upstream description:

Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other
API-endpoint not listed will act according to "Fall Through Mode".

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
  },
  "x-ves-oneof-field-oversized_body_choice": "[]"
}
```

Terraform syntax:

```terraform
validation_custom_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [fall_through_mode](resources--http_loadbalancer--properties--api_specification--validation_custom_list--fall_through_mode.md): complete subsection reference.

- [open_api_validation_rules](resources--http_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules.md): complete subsection reference.

- [settings](resources--http_loadbalancer--properties--api_specification--validation_custom_list--settings.md): complete subsection reference.

## Next pages

- [api_specification.validation_custom_list.fall_through_mode](resources--http_loadbalancer--properties--api_specification--validation_custom_list--fall_through_mode.md)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules.md)
- [api_specification.validation_custom_list.settings](resources--http_loadbalancer--properties--api_specification--validation_custom_list--settings.md)
- [api_specification](resources--http_loadbalancer--properties--api_specification.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
