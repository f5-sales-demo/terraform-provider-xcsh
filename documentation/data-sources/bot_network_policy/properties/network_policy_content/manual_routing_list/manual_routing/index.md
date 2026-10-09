---
page_title: "network_policy_content.manual_routing_list.manual_routing"
subcategory: ""
description: "Manual Routing. Routing or forwarding configuration"
xcsh_docs: {"aliases": ["network policy content manual routing list manual routing"], "body_bytes": 2377, "body_sha256": "sha256:0df7f4b4ebe3bde4397124e1dfc3122e5110852ec5c6681bd8066911e46fd3ea", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list:manual_routing:http", "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list:manual_routing:https", "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list:manual_routing:protocol_http", "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list:manual_routing:protocol_https"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_network_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list:manual_routing", "parent_id": "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list", "path": "documentation/data-sources/bot_network_policy/properties/network_policy_content/manual_routing_list/manual_routing/index.md", "product": "distributed-cloud", "provider_name": "bot_network_policy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0012001030002311-0213321022003013-2032333303110130-3323201210303213-2320111032001200-3021213331320032-1323032221030132-1001232100102311", "registry_path": "docs/guides/data-sources--bot_network_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["network_policy_content", "manual_routing_list", "manual_routing"], "schema_version": 1, "sections": [{"aliases": ["network policy content manual routing list manual routing domain name"], "anchor": "schema-network_policy_content--manual_routing_list--manual_routing--domain_name", "description": "Inbound FQDN. Inbound FQDN value.", "document_id": "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list:manual_routing", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["network_policy_content", "manual_routing_list", "manual_routing", "domain_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["network policy content manual routing list manual routing http"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list:manual_routing:http", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["network_policy_content", "manual_routing_list", "manual_routing", "http"], "syntax": "attribute", "type": "object"}, {"aliases": ["network policy content manual routing list manual routing https"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list:manual_routing:https", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["network_policy_content", "manual_routing_list", "manual_routing", "https"], "syntax": "attribute", "type": "object"}, {"aliases": ["network policy content manual routing list manual routing outbound domain name"], "anchor": "schema-network_policy_content--manual_routing_list--manual_routing--outbound_domain_name", "description": "Outbound FQDN / IP. Outbound FQDN or IP value.", "document_id": "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list:manual_routing", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["network_policy_content", "manual_routing_list", "manual_routing", "outbound_domain_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["network policy content manual routing list manual routing port"], "anchor": "schema-network_policy_content--manual_routing_list--manual_routing--port", "description": "Outbound Port. Outbound Port value.", "document_id": "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list:manual_routing", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["network_policy_content", "manual_routing_list", "manual_routing", "port"], "syntax": "attribute", "type": "number"}, {"aliases": ["network policy content manual routing list manual routing protocol http"], "anchor": "section", "description": "Configuration parameter for protocol http.", "document_id": "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list:manual_routing:protocol_http", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["network_policy_content", "manual_routing_list", "manual_routing", "protocol_http"], "syntax": "attribute", "type": "object"}, {"aliases": ["network policy content manual routing list manual routing protocol https"], "anchor": "section", "description": "Configuration parameter for protocol https.", "document_id": "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list:manual_routing:protocol_https", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["network_policy_content", "manual_routing_list", "manual_routing", "protocol_https"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_network_policy/properties/network_policy_content/manual_routing_list/manual_routing/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Manual Routing. Routing or forwarding configuration", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# network_policy_content.manual_routing_list.manual_routing

Breadcrumbs:

- [xcsh_bot_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/)
- [network_policy_content](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/network_policy_content/)
- [network_policy_content.manual_routing_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/network_policy_content/manual_routing_list/)
- network_policy_content.manual_routing_list.manual_routing

<a id="section"></a>

Type: `"list"`. Computed.

Manual Routing. Routing or forwarding configuration

## Direct properties

<a id="schema-network_policy_content--manual_routing_list--manual_routing--domain_name"></a>

### domain_name property

Type: `"string"`. Computed.

Inbound FQDN. Inbound FQDN value.

- [http](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/network_policy_content/manual_routing_list/manual_routing/http/): complete subsection reference.

- [https](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/network_policy_content/manual_routing_list/manual_routing/https/): complete subsection reference.

<a id="schema-network_policy_content--manual_routing_list--manual_routing--outbound_domain_name"></a>

### outbound_domain_name property

Type: `"string"`. Computed.

Outbound FQDN / IP. Outbound FQDN or IP value.

<a id="schema-network_policy_content--manual_routing_list--manual_routing--port"></a>

### port property

Type: `"number"`. Computed.

Outbound Port. Outbound Port value.

- [protocol_http](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/network_policy_content/manual_routing_list/manual_routing/protocol_http/): complete subsection reference.

- [protocol_https](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/network_policy_content/manual_routing_list/manual_routing/protocol_https/): complete subsection reference.
