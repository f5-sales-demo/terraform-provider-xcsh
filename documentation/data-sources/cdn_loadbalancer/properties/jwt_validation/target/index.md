---
page_title: "jwt_validation.target"
subcategory: "Load Balancing"
description: "Define endpoints for which JWT token validation will be performed."
xcsh_docs: {"aliases": ["jwt validation target"], "body_bytes": 2247, "body_sha256": "sha256:0b43ee7c22b460fd0981b8d296bcc857dc0a4dc1c9e4eafc3d749be2daacfabc", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:target:all_endpoint", "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:target:api_groups", "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:target:base_paths"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:target", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation", "path": "documentation/data-sources/cdn_loadbalancer/properties/jwt_validation/target/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2003033103100212-0221322000032122-1033301212022110-3311221200010103-3112220320320320-1031002021233000-3210122333012323-1113020133332202", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-011.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["jwt_validation", "target"], "schema_version": 1, "sections": [{"aliases": ["jwt validation target all endpoint"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:target:all_endpoint", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["jwt_validation", "target", "all_endpoint"], "syntax": "attribute", "type": "object"}, {"aliases": ["jwt validation target api groups"], "anchor": "section", "description": "API Groups.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:target:api_groups", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["jwt_validation", "target", "api_groups"], "syntax": "attribute", "type": "object"}, {"aliases": ["jwt validation target base paths"], "anchor": "section", "description": "Base Paths.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:target:base_paths", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["jwt_validation", "target", "base_paths"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/jwt_validation/target/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Define endpoints for which JWT token validation will be performed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation.target

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [jwt_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/jwt_validation/)
- jwt_validation.target

<a id="section"></a>

Type: `"single"`. Computed.

Define endpoints for which JWT token validation will be performed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-target": "[\"all_endpoint\",\"api_groups\",\"base_paths\"]"
}
```

## Direct properties

- [all_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/jwt_validation/target/all_endpoint/): complete subsection reference.

- [api_groups](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/jwt_validation/target/api_groups/): complete subsection reference.

- [base_paths](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/jwt_validation/target/base_paths/): complete subsection reference.

## Next pages

- [jwt_validation.target.all_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/jwt_validation/target/all_endpoint/)
- [jwt_validation.target.api_groups](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/jwt_validation/target/api_groups/)
- [jwt_validation.target.base_paths](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/jwt_validation/target/base_paths/)
- [jwt_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/jwt_validation/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
