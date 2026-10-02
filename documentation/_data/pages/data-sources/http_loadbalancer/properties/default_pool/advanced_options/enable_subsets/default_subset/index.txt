---
page_title: "default_pool.advanced_options.enable_subsets.default_subset"
subcategory: "Load Balancing"
description: "Default Subset definition."
xcsh_docs: {"aliases": ["default pool advanced options enable subsets default subset"], "body_bytes": 2116, "body_sha256": "sha256:f6a42f47351e7b66178c4e9cfa705715140f09471b961beeb38b65830585ea93", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:default_subset:default_subset"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:default_subset", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets", "path": "documentation/data-sources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/default_subset/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3221030202032311-3023220331221201-0010011333222003-3130103002003322-1120311323130221-2112030023023112-0331303202211111-0202213322011212", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-015.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_pool", "advanced_options", "enable_subsets", "default_subset"], "schema_version": 1, "sections": [{"aliases": ["default subset"], "anchor": "section", "description": "List of key-value pairs that define default subset. Which gets used when route specifies no metadata or no subset matching the metadata exists.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:default_subset:default_subset", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "advanced_options", "enable_subsets", "default_subset", "default_subset"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/default_subset/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Default Subset definition.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.advanced_options.enable_subsets.default_subset

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/)
- [default_pool.advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/advanced_options/)
- [default_pool.advanced_options.enable_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/)
- default_pool.advanced_options.enable_subsets.default_subset

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for default subset.

Upstream description:

Default Subset definition.

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

## Direct properties

- [default_subset](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/default_subset/default_subset/): complete subsection reference.

## Next pages

- [default_pool.advanced_options.enable_subsets.default_subset.default_subset](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/default_subset/default_subset/)
- [default_pool.advanced_options.enable_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
