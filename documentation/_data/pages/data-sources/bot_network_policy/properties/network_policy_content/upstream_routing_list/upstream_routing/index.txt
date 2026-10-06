---
page_title: "network_policy_content.upstream_routing_list.upstream_routing"
subcategory: ""
description: "Upstream Routing. Routing or forwarding configuration"
xcsh_docs: {"aliases": ["network policy content upstream routing list upstream routing"], "body_bytes": 1129, "body_sha256": "sha256:4808ff04ca5cf0779c27ef944458eb2d0e9f20d906cd52b165ee687daf6f57a7", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_network_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:upstream_routing_list:upstream_routing", "parent_id": "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:upstream_routing_list", "path": "documentation/data-sources/bot_network_policy/properties/network_policy_content/upstream_routing_list/upstream_routing/index.md", "product": "distributed-cloud", "provider_name": "bot_network_policy", "provider_schema_digest": "sha256:057968f86e4ef0ae0087dd4d6097131e285b60998d97655ee1a02c00315f6b0f", "provider_type": "data-sources", "registry_anchor": "canonical-0201101112023233-3021323213231313-3020033032203120-2332122023222211-0022202220303212-3101311322222020-1131320210010031-3131322221110030", "registry_path": "docs/guides/data-sources--bot_network_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["network_policy_content", "upstream_routing_list", "upstream_routing"], "schema_version": 1, "sections": [{"aliases": ["network policy content upstream routing list upstream routing domain name"], "anchor": "schema-network_policy_content--upstream_routing_list--upstream_routing--domain_name", "description": "FQDN. Domain Name.", "document_id": "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:upstream_routing_list:upstream_routing", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["network_policy_content", "upstream_routing_list", "upstream_routing", "domain_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_network_policy/properties/network_policy_content/upstream_routing_list/upstream_routing/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Upstream Routing. Routing or forwarding configuration", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": [], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# network_policy_content.upstream_routing_list.upstream_routing

Breadcrumbs:

- [xcsh_bot_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/)
- [network_policy_content](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/network_policy_content/)
- [network_policy_content.upstream_routing_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/network_policy_content/upstream_routing_list/)
- network_policy_content.upstream_routing_list.upstream_routing

<a id="section"></a>

Type: `"list"`. Computed.

Upstream Routing. Routing or forwarding configuration

## Direct properties

<a id="schema-network_policy_content--upstream_routing_list--upstream_routing--domain_name"></a>

### domain_name property

Type: `"string"`. Computed.

FQDN. Domain Name.
