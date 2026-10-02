---
page_title: "api_specification.validation_custom_list.fall_through_mode"
subcategory: "Load Balancing"
description: "Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a. Swagger) or doesn't have a specific rule in custom rules)"
xcsh_docs: {"aliases": ["api specification validation custom list fall through mode"], "body_bytes": 2829, "body_sha256": "sha256:600bc9f1b8e0f8d41ab5c4f90cf2853618e943c23f07888faf264bd9e04bf55c", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_allow", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list", "path": "documentation/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0123133020003332-0303201321022202-2103000012220020-1123010020233030-1210313311213211-2331133333300023-2202123013033303-3121111020322233", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-006.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "fall_through_mode"], "schema_version": 1, "sections": [{"aliases": ["fall through mode allow"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_allow", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "fall_through_mode", "fall_through_mode_allow"], "syntax": "attribute", "type": "object"}, {"aliases": ["fall through mode custom"], "anchor": "section", "description": "Define the fall through settings.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "fall_through_mode", "fall_through_mode_custom"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a. Swagger) or doesn't have a specific rule in custom rules)", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_custom_list.fall_through_mode

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/)
- [api_specification.validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/)
- api_specification.validation_custom_list.fall_through_mode

<a id="section"></a>

Type: `"single"`. Computed.

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules).

Upstream description:

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules)

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

## Direct properties

- [fall_through_mode_allow](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_allow/): complete subsection reference.

- [fall_through_mode_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_custom/): complete subsection reference.

## Next pages

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_allow](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_allow/)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_custom/)
- [api_specification.validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
