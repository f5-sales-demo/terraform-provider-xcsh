---
page_title: "api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom"
subcategory: "Load Balancing"
description: "api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2097, "body_sha256": "sha256:7c398fe13d0450292bf70e64839dba375befddb47cbdb6923416f070211ab506", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom:open_api_validation_rules"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode", "path": "docs/guides/resources--cdn_loadbalancer--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "fall_through_mode", "fall_through_mode_custom"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_custom/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [api_specification](resources--cdn_loadbalancer--properties--api_specification.md)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--properties--api_specification--validation_custom_list.md)
- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--properties--api_specification--validation_custom_list--fall_through_mode.md)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom

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

- [open_api_validation_rules](resources--cdn_loadbalancer--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules.md): complete subsection reference.

## Next pages

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules.md)
- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--properties--api_specification--validation_custom_list--fall_through_mode.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
