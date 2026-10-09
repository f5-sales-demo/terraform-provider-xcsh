---
page_title: "network_policy_content.upstream_routing_list.upstream_routing"
subcategory: ""
description: "Upstream Routing. Routing or forwarding configuration"
xcsh_docs: {"aliases": ["network policy content upstream routing list upstream routing"], "body_bytes": 1129, "body_sha256": "sha256:4808ff04ca5cf0779c27ef944458eb2d0e9f20d906cd52b165ee687daf6f57a7", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_network_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:upstream_routing_list:upstream_routing", "parent_id": "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:upstream_routing_list", "path": "documentation/data-sources/bot_network_policy/properties/network_policy_content/upstream_routing_list/upstream_routing/index.md", "product": "distributed-cloud", "provider_name": "bot_network_policy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0201101112023233-3021323213231313-3020033032203120-2332122023222211-0022202220303212-3101311322222020-1131320210010031-3131322221110030", "registry_path": "docs/guides/data-sources--bot_network_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["network_policy_content", "upstream_routing_list", "upstream_routing"], "schema_version": 1, "sections": [{"aliases": ["network policy content upstream routing list upstream routing domain name"], "anchor": "schema-network_policy_content--upstream_routing_list--upstream_routing--domain_name", "description": "FQDN. Domain Name.", "document_id": "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:upstream_routing_list:upstream_routing", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["network_policy_content", "upstream_routing_list", "upstream_routing", "domain_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_network_policy/properties/network_policy_content/upstream_routing_list/upstream_routing/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Upstream Routing. Routing or forwarding configuration", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
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
