---
page_title: "api_specification"
subcategory: "Load Balancing"
description: "api_specification for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2999, "body_sha256": "sha256:1fcf3f2299b2f3e27748e95d22d8698c1775bed3999eefd159b1452eee9842dc", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:api_definition", "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints", "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list", "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_disabled"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "docs/guides/resources--cdn_loadbalancer--properties--api_specification.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/api_specification/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
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

- [api_specification](resources--cdn_loadbalancer--properties--api_specification.md#section)
- [disable_api_definition](resources--cdn_loadbalancer--properties--disable_api_definition.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
api_specification {
  # Configure direct properties listed below.
}
```

## Direct properties

- [api_definition](resources--cdn_loadbalancer--properties--api_specification--api_definition.md): complete subsection reference.

- [validation_all_spec_endpoints](resources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints.md): complete subsection reference.

- [validation_custom_list](resources--cdn_loadbalancer--properties--api_specification--validation_custom_list.md): complete subsection reference.

- [validation_disabled](resources--cdn_loadbalancer--properties--api_specification--validation_disabled.md): complete subsection reference.

## Next pages

- [api_specification.api_definition](resources--cdn_loadbalancer--properties--api_specification--api_definition.md)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints.md)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--properties--api_specification--validation_custom_list.md)
- [api_specification.validation_disabled](resources--cdn_loadbalancer--properties--api_specification--validation_disabled.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
