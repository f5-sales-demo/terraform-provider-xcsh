---
page_title: "default_underlay_network"
subcategory: "Infrastructure"
description: "Optional, virtual network to be used as underlay for different overlay protocols (SRv6, IP-in-IP tunnels for DC Cluster Group) Default is site-local-outside network."
xcsh_docs: {"aliases": ["default underlay network"], "body_bytes": 980, "body_sha256": "sha256:298461e5b8edaee855da2cdecbf7baa12da2498e529136db1c1e42774157fc98", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:site:properties:default_underlay_network:site_local_inside", "xcsh-docs:data-sources:site:properties:default_underlay_network:site_local_outside"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site:properties:default_underlay_network", "parent_id": "xcsh-docs:data-sources:site:reference", "path": "documentation/data-sources/site/properties/default_underlay_network/index.md", "product": "distributed-cloud", "provider_name": "site", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2331011302122111-1002222231122333-2011102021021020-1112333113122130-1022201111202013-3113000222232220-1111222213122020-3212033033033311", "registry_path": "docs/guides/data-sources--site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_underlay_network"], "schema_version": 1, "sections": [{"aliases": ["default underlay network site local inside"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:site:properties:default_underlay_network:site_local_inside", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_underlay_network", "site_local_inside"], "syntax": "attribute", "type": "object"}, {"aliases": ["default underlay network site local outside"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:site:properties:default_underlay_network:site_local_outside", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_underlay_network", "site_local_outside"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site/properties/default_underlay_network/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Optional, virtual network to be used as underlay for different overlay protocols (SRv6, IP-in-IP tunnels for DC Cluster Group) Default is site-local-outside network.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
