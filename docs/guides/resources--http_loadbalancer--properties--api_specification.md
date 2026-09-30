---
page_title: "api_specification"
subcategory: "Load Balancing"
description: "api_specification for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2916, "body_sha256": "sha256:bf1c86cf4b764bc49632b151004385a152c683d10d42244ff4531de2cee79700", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_specification:api_definition", "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints", "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list", "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_disabled"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "docs/guides/resources--http_loadbalancer--properties--api_specification.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_specification/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# api_specification

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- api_specification

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: api\_specification, disable\_api\_definition; Default: disable\_api\_definition\] Settings
for API specification (API definition, OpenAPI validation, etc.).

Upstream description:

Settings for API specification (API definition, OpenAPI validation, etc.)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("validation_all_spec_endpoints",
    "validation_custom_list"),
  validators.ConflictingObjectAttributes("validation_all_spec_endpoints",
    "validation_disabled"),
  validators.ConflictingObjectAttributes("validation_custom_list",
    "validation_disabled")}
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
  "x-ves-oneof-field-validation_target_choice": "[\"validation_all_spec_endpoints\",\"validation_custom_list\",\"validation_disabled\"]"
}
```

OneOf alternatives in this subsection:

- [api_specification](resources--http_loadbalancer--properties--api_specification.md#section)
- [disable_api_definition](resources--http_loadbalancer--properties--disable_api_definition.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
api_specification {
  # Configure direct properties listed below.
}
```

## Direct properties

- [api_definition](resources--http_loadbalancer--properties--api_specification--api_definition.md): complete subsection reference.

- [validation_all_spec_endpoints](resources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints.md): complete subsection reference.

- [validation_custom_list](resources--http_loadbalancer--properties--api_specification--validation_custom_list.md): complete subsection reference.

- [validation_disabled](resources--http_loadbalancer--properties--api_specification--validation_disabled.md): complete subsection reference.

## Next pages

- [api_specification.api_definition](resources--http_loadbalancer--properties--api_specification--api_definition.md)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints.md)
- [api_specification.validation_custom_list](resources--http_loadbalancer--properties--api_specification--validation_custom_list.md)
- [api_specification.validation_disabled](resources--http_loadbalancer--properties--api_specification--validation_disabled.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
