---
page_title: "jwt_validation.action"
subcategory: "Load Balancing"
description: "Action"
xcsh_docs: {"aliases": ["jwt validation action"], "body_bytes": 1775, "body_sha256": "sha256:f70e31ae0c7521a9426d7eae116f71b154f28d6ad49fb0d1c084dc0602c62fb7", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:action:block", "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:action:report"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:action", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation", "path": "documentation/data-sources/cdn_loadbalancer/properties/jwt_validation/action/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2121132033013110-0203011222103223-3211020302222303-0223111212332002-1211112230203201-2000211102021000-2110302303321020-0031111212212330", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-011.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["jwt_validation", "action"], "schema_version": 1, "sections": [{"aliases": ["block"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:action:block", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["jwt_validation", "action", "block"], "syntax": "attribute", "type": "object"}, {"aliases": ["report"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:jwt_validation:action:report", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["jwt_validation", "action", "report"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/jwt_validation/action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Action", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation.action

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [jwt_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/jwt_validation/)
- jwt_validation.action

<a id="section"></a>

Type: `"single"`. Computed.

Action

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_choice": "[\"block\",\"report\"]"
}
```

## Direct properties

- [block](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/jwt_validation/action/block/): complete subsection reference.

- [report](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/jwt_validation/action/report/): complete subsection reference.

## Next pages

- [jwt_validation.action.block](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/jwt_validation/action/block/)
- [jwt_validation.action.report](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/jwt_validation/action/report/)
- [jwt_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/jwt_validation/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
