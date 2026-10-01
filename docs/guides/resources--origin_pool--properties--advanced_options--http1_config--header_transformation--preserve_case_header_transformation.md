---
page_title: "advanced_options.http1_config.header_transformation.preserve_case_header_transformation"
subcategory: "Load Balancing"
description: "advanced_options.http1_config.header_transformation.preserve_case_header_transformation for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1432, "body_sha256": "sha256:0d22e1fe22893d0fd41b7ed4408e40ac7cd70f67ca181ecec1ab9b38458ab63a", "canonical_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:http1_config:header_transformation:preserve_case_header_transformation", "child_ids": [], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:advanced_options:http1_config:header_transformation:preserve_case_header_transformation", "parent_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:http1_config:header_transformation", "path": "docs/guides/resources--origin_pool--properties--advanced_options--http1_config--header_transformation--preserve_case_header_transformation.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advanced_options", "http1_config", "header_transformation", "preserve_case_header_transformation"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/advanced_options/http1_config/header_transformation/preserve_case_header_transformation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advanced_options.http1_config.header_transformation.preserve_case_header_transformation for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_options.http1_config.header_transformation.preserve_case_header_transformation

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md)
- [Property reference](resources--origin_pool--reference.md)
- [advanced_options](resources--origin_pool--properties--advanced_options.md)
- [advanced_options.http1_config](resources--origin_pool--properties--advanced_options--http1_config.md)
- [advanced_options.http1_config.header_transformation](resources--origin_pool--properties--advanced_options--http1_config--header_transformation.md)
- advanced_options.http1_config.header_transformation.preserve_case_header_transformation

<a id="section"></a>

Type: `["object", {}]`. Optional.

Preserve HTTP header-name case when upstream case must remain unchanged.

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
preserve_case_header_transformation = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [advanced_options.http1_config.header_transformation](resources--origin_pool--properties--advanced_options--http1_config--header_transformation.md)
- [xcsh_origin_pool](../resources/origin_pool.md)
