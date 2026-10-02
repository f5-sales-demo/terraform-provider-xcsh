---
page_title: "cdn_loadbalancer"
subcategory: ""
description: "Set the scope of the API Group to a specific CDN Loadbalancer."
xcsh_docs: {"aliases": ["cdn loadbalancer"], "body_bytes": 1393, "body_sha256": "sha256:51fd84b0c561443fbf861801a1d455579cc62d5c3c57813e5a0e96ee2ad79329", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:resources:app_api_group:properties:cdn_loadbalancer:cdn_loadbalancer"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:app_api_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_api_group:properties:cdn_loadbalancer", "parent_id": "xcsh-docs:resources:app_api_group:reference", "path": "documentation/resources/app_api_group/properties/cdn_loadbalancer/index.md", "product": "distributed-cloud", "provider_name": "app_api_group", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2122310110220121-1311210010012320-2130100120201203-2012322123033101-1200100321200231-2131013122300131-2032300202032303-2103323121312333", "registry_path": "docs/guides/resources--app_api_group--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cdn_loadbalancer"], "schema_version": 1, "sections": [{"aliases": ["cdn loadbalancer"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:app_api_group:properties:cdn_loadbalancer:cdn_loadbalancer", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-cdn_loadbalancer--cdn_loadbalancer--name", "enforcement": "provider-schema", "group": "cdn_loadbalancer.cdn_loadbalancer:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:app_api_group:properties:cdn_loadbalancer:cdn_loadbalancer", "type": "requires"}], "schema_path": ["cdn_loadbalancer", "cdn_loadbalancer"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_api_group/properties/cdn_loadbalancer/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Set the scope of the API Group to a specific CDN Loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_api_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cdn_loadbalancer

Breadcrumbs:

- [xcsh_app_api_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/)
- cdn_loadbalancer

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Set the scope of the API Group to a specific CDN Loadbalancer.

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
cdn_loadbalancer {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/cdn_loadbalancer/cdn_loadbalancer/): complete subsection reference.

## Next pages

- [cdn_loadbalancer.cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/cdn_loadbalancer/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/)
- [xcsh_app_api_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/)
