---
page_title: "jwt_validation.target"
subcategory: "Load Balancing"
description: "Define endpoints for which JWT token validation will be performed."
xcsh_docs: {"aliases": ["jwt validation target"], "body_bytes": 2260, "body_sha256": "sha256:c77fb35c63b78da305e336d9bdf4634f59bc591212730f0f357cfb912b2f4682", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:target:all_endpoint", "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:target:api_groups", "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:target:base_paths"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:target", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation", "path": "documentation/data-sources/http_loadbalancer/properties/jwt_validation/target/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3320100111233100-2113302231232203-3303311321213110-0121302310123021-1231113212032011-3001021101020021-2023122021203001-2010102033133321", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-020.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["jwt_validation", "target"], "schema_version": 1, "sections": [{"aliases": ["all endpoint"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:target:all_endpoint", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["jwt_validation", "target", "all_endpoint"], "syntax": "attribute", "type": "object"}, {"aliases": ["api groups"], "anchor": "section", "description": "API Groups.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:target:api_groups", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["jwt_validation", "target", "api_groups"], "syntax": "attribute", "type": "object"}, {"aliases": ["base paths"], "anchor": "section", "description": "Base Paths.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:target:base_paths", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["jwt_validation", "target", "base_paths"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/jwt_validation/target/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Define endpoints for which JWT token validation will be performed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation.target

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [jwt_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/)
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

- [all_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/target/all_endpoint/): complete subsection reference.

- [api_groups](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/target/api_groups/): complete subsection reference.

- [base_paths](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/target/base_paths/): complete subsection reference.

## Next pages

- [jwt_validation.target.all_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/target/all_endpoint/)
- [jwt_validation.target.api_groups](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/target/api_groups/)
- [jwt_validation.target.base_paths](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/target/base_paths/)
- [jwt_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
