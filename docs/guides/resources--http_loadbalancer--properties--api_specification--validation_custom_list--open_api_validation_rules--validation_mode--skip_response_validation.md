---
page_title: "api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation"
subcategory: "Load Balancing"
description: "api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1879, "body_sha256": "sha256:a98f0275e4dc538c59d22f842f59d48d0897ca612bda64222ab1df493f0a2a7c", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:skip_response_validation", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:skip_response_validation", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode", "path": "docs/guides/resources--http_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--skip_response_validation.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode", "skip_response_validation"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/skip_response_validation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [api_specification](resources--http_loadbalancer--properties--api_specification.md)
- [api_specification.validation_custom_list](resources--http_loadbalancer--properties--api_specification--validation_custom_list.md)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules.md)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--http_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode.md)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation

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

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--http_loadbalancer--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
