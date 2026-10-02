---
page_title: "routes.simple_route.advanced_options.specific_hash_policy"
subcategory: "Load Balancing"
description: "List of hash policy rules."
xcsh_docs: {"aliases": ["routes simple route advanced options specific hash policy"], "body_bytes": 2025, "body_sha256": "sha256:e12d871a49b6611288b7d5dbac81c040f2e72c56daa4f641ff34a814f0e9665a", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options", "path": "documentation/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/specific_hash_policy/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0102213113001301-2132232011232310-2301301201331013-0033113222132023-2230311110230131-2212002213230233-0131303133313113-3102312200120103", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-025.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "simple_route", "advanced_options", "specific_hash_policy"], "schema_version": 1, "sections": [{"aliases": ["hash policy"], "anchor": "section", "description": "Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated individually and the combined result is used to route the request.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "specific_hash_policy", "hash_policy"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/specific_hash_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of hash policy rules.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.simple_route.advanced_options.specific_hash_policy

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/)
- [routes.simple_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/)
- [routes.simple_route.advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/)
- routes.simple_route.advanced_options.specific_hash_policy

<a id="section"></a>

Type: `"single"`. Computed.

Policy configuration for this feature.

Upstream description:

List of hash policy rules.

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

- [hash_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/specific_hash_policy/hash_policy/): complete subsection reference.

## Next pages

- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/specific_hash_policy/hash_policy/)
- [routes.simple_route.advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
