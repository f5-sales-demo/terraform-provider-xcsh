---
page_title: "default_underlay_network"
subcategory: "Infrastructure"
description: "Optional, virtual network to be used as underlay for different overlay protocols (SRv6, IP-in-IP tunnels for DC Cluster Group) Default is site-local-outside network."
xcsh_docs: {"aliases": ["default underlay network"], "body_bytes": 980, "body_sha256": "sha256:298461e5b8edaee855da2cdecbf7baa12da2498e529136db1c1e42774157fc98", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:site:properties:default_underlay_network:site_local_inside", "xcsh-docs:data-sources:site:properties:default_underlay_network:site_local_outside"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site:properties:default_underlay_network", "parent_id": "xcsh-docs:data-sources:site:reference", "path": "documentation/data-sources/site/properties/default_underlay_network/index.md", "product": "distributed-cloud", "provider_name": "site", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "data-sources", "registry_anchor": "canonical-2331011302122111-1002222231122333-2011102021021020-1112333113122130-1022201111202013-3113000222232220-1111222213122020-3212033033033311", "registry_path": "docs/guides/data-sources--site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_underlay_network"], "schema_version": 1, "sections": [{"aliases": ["default underlay network site local inside"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:site:properties:default_underlay_network:site_local_inside", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_underlay_network", "site_local_inside"], "syntax": "attribute", "type": "object"}, {"aliases": ["default underlay network site local outside"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:site:properties:default_underlay_network:site_local_outside", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_underlay_network", "site_local_outside"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site/properties/default_underlay_network/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Optional, virtual network to be used as underlay for different overlay protocols (SRv6, IP-in-IP tunnels for DC Cluster Group) Default is site-local-outside network.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": [], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
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
