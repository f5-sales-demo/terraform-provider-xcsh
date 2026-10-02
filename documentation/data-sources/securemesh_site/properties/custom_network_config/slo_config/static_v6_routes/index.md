---
page_title: "custom_network_config.slo_config.static_v6_routes"
subcategory: ""
description: "List of IPv6 static routes."
xcsh_docs: {"aliases": ["custom network config slo config static v6 routes"], "body_bytes": 1851, "body_sha256": "sha256:4812f913bd8a58ff505b2176aa7e7d50f73429f0b8e6d1b0cf2ec5dc849ee9e0", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:slo_config:static_v6_routes:static_routes"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:slo_config:static_v6_routes", "parent_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:slo_config", "path": "documentation/data-sources/securemesh_site/properties/custom_network_config/slo_config/static_v6_routes/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1002121031202323-2322130011221312-3320211122103210-1230301100013321-3102122303301211-0220020310031221-1011131300231103-3112323020303331", "registry_path": "docs/guides/data-sources--securemesh_site--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "slo_config", "static_v6_routes"], "schema_version": 1, "sections": [{"aliases": ["static routes"], "anchor": "section", "description": "List of IPv6 static routes.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:slo_config:static_v6_routes:static_routes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["custom_network_config", "slo_config", "static_v6_routes", "static_routes"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site/properties/custom_network_config/slo_config/static_v6_routes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of IPv6 static routes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.slo_config.static_v6_routes

Breadcrumbs:

- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/)
- [custom_network_config.slo_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/slo_config/)
- custom_network_config.slo_config.static_v6_routes

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for static v6 routes.

Upstream description:

List of IPv6 static routes.

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

- [static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/slo_config/static_v6_routes/static_routes/): complete subsection reference.

## Next pages

- [custom_network_config.slo_config.static_v6_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/slo_config/static_v6_routes/static_routes/)
- [custom_network_config.slo_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/slo_config/)
- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
