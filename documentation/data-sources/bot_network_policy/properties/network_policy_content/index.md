---
page_title: "network_policy_content"
subcategory: ""
description: "Configuration parameter for network policy content."
xcsh_docs: {"aliases": ["network policy content"], "body_bytes": 938, "body_sha256": "sha256:ba34e5d2246a58b290b9d196af07cc2118ce7363f941eb93fbc5a9780caee203", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list", "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:upstream_routing_list"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_network_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content", "parent_id": "xcsh-docs:data-sources:bot_network_policy:reference", "path": "documentation/data-sources/bot_network_policy/properties/network_policy_content/index.md", "product": "distributed-cloud", "provider_name": "bot_network_policy", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-2021023211330003-0222133223333021-1300011020020233-3222322203000300-2120001002232331-1200013000103331-1222032222212133-2112220110020020", "registry_path": "docs/guides/data-sources--bot_network_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["network_policy_content"], "schema_version": 1, "sections": [{"aliases": ["network policy content manual routing list"], "anchor": "section", "description": "Manual Routings. The list of manual routing.", "document_id": "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["network_policy_content", "manual_routing_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["network policy content upstream routing list"], "anchor": "section", "description": "Upstream Routings. Upstream DNS Routings.", "document_id": "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:upstream_routing_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["network_policy_content", "upstream_routing_list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_network_policy/properties/network_policy_content/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Configuration parameter for network policy content.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": [], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# network_policy_content

Breadcrumbs:

- [xcsh_bot_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/)
- network_policy_content

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for network policy content.

## Direct properties

- [manual_routing_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/network_policy_content/manual_routing_list/): complete subsection reference.

- [upstream_routing_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/network_policy_content/upstream_routing_list/): complete subsection reference.
