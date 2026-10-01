---
page_title: "api_specification.validation_custom_list.settings.oversized_body_fail_validation"
subcategory: "Load Balancing"
description: "api_specification.validation_custom_list.settings.oversized_body_fail_validation for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1511, "body_sha256": "sha256:fbcabff83faf0867aa0bd39bec25a1c4f0333b49dd354196cb23a9a6c184d514", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:oversized_body_fail_validation", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:oversized_body_fail_validation", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings", "path": "docs/guides/resources--http_loadbalancer--properties--api_specification--validation_custom_list--settings--oversized_body_fail_validation.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "settings", "oversized_body_fail_validation"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_specification/validation_custom_list/settings/oversized_body_fail_validation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification.validation_custom_list.settings.oversized_body_fail_validation for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_custom_list.settings.oversized_body_fail_validation

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [api_specification](resources--http_loadbalancer--properties--api_specification.md)
- [api_specification.validation_custom_list](resources--http_loadbalancer--properties--api_specification--validation_custom_list.md)
- [api_specification.validation_custom_list.settings](resources--http_loadbalancer--properties--api_specification--validation_custom_list--settings.md)
- api_specification.validation_custom_list.settings.oversized_body_fail_validation

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
oversized_body_fail_validation = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [api_specification.validation_custom_list.settings](resources--http_loadbalancer--properties--api_specification--validation_custom_list--settings.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
