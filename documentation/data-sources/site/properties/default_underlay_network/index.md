---
page_title: "default_underlay_network"
subcategory: "Infrastructure"
description: "Optional, virtual network to be used as underlay for different overlay protocols (SRv6, IP-in-IP tunnels for DC Cluster Group) Default is site-local-outside network."
xcsh_docs: {"aliases": ["default underlay network"], "body_bytes": 1552, "body_sha256": "sha256:a805ee8b3bd4347be6fa4ae39b0149b64d60b8ed9f370062daf946139492ca54", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:site:properties:default_underlay_network:site_local_inside", "xcsh-docs:data-sources:site:properties:default_underlay_network:site_local_outside"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site:properties:default_underlay_network", "parent_id": "xcsh-docs:data-sources:site:reference", "path": "documentation/data-sources/site/properties/default_underlay_network/index.md", "product": "distributed-cloud", "provider_name": "site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2331011302122111-1002222231122333-2011102021021020-1112333113122130-1022201111202013-3113000222232220-1111222213122020-3212033033033311", "registry_path": "docs/guides/data-sources--site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_underlay_network"], "schema_version": 1, "sections": [{"aliases": ["default underlay network site local inside"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:site:properties:default_underlay_network:site_local_inside", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_underlay_network", "site_local_inside"], "syntax": "attribute", "type": "object"}, {"aliases": ["default underlay network site local outside"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:site:properties:default_underlay_network:site_local_outside", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_underlay_network", "site_local_outside"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site/properties/default_underlay_network/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Optional, virtual network to be used as underlay for different overlay protocols (SRv6, IP-in-IP tunnels for DC Cluster Group) Default is site-local-outside network.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": [], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_underlay_network

Breadcrumbs:

- [xcsh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/)
- default_underlay_network

<a id="section"></a>

Type: `"single"`. Computed.

Optional, virtual network to be used as underlay for different overlay protocols (SRv6, IP-in-IP
tunnels for DC Cluster Group) Default is site-local-outside network.

## Direct properties

- [site_local_inside](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/default_underlay_network/site_local_inside/): complete subsection reference.

- [site_local_outside](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/default_underlay_network/site_local_outside/): complete subsection reference.

## Next pages

- [default_underlay_network.site_local_inside](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/default_underlay_network/site_local_inside/)
- [default_underlay_network.site_local_outside](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/default_underlay_network/site_local_outside/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/)
- [xcsh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/)
