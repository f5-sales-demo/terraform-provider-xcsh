---
page_title: "network_policy_content.manual_routing_list.manual_routing"
subcategory: ""
description: "Manual Routing. Routing or forwarding configuration"
xcsh_docs: {"aliases": ["network policy content manual routing list manual routing"], "body_bytes": 3666, "body_sha256": "sha256:e1c2000b9caa5058707b956c0bf7830d1d4a770eaf36d670df7995e95e9c2b1b", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list:manual_routing:http", "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list:manual_routing:https", "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list:manual_routing:protocol_http", "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list:manual_routing:protocol_https"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_network_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list:manual_routing", "parent_id": "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list", "path": "documentation/data-sources/bot_network_policy/properties/network_policy_content/manual_routing_list/manual_routing/index.md", "product": "distributed-cloud", "provider_name": "bot_network_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0012001030002311-0213321022003013-2032333303110130-3323201210303213-2320111032001200-3021213331320032-1323032221030132-1001232100102311", "registry_path": "docs/guides/data-sources--bot_network_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["network_policy_content", "manual_routing_list", "manual_routing"], "schema_version": 1, "sections": [{"aliases": ["domain name"], "anchor": "schema-network_policy_content--manual_routing_list--manual_routing--domain_name", "description": "Inbound FQDN. Inbound FQDN value.", "document_id": "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list:manual_routing", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["network_policy_content", "manual_routing_list", "manual_routing", "domain_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["http"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list:manual_routing:http", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["network_policy_content", "manual_routing_list", "manual_routing", "http"], "syntax": "attribute", "type": "object"}, {"aliases": ["https"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list:manual_routing:https", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["network_policy_content", "manual_routing_list", "manual_routing", "https"], "syntax": "attribute", "type": "object"}, {"aliases": ["outbound domain name"], "anchor": "schema-network_policy_content--manual_routing_list--manual_routing--outbound_domain_name", "description": "Outbound FQDN / IP. Outbound FQDN or IP value.", "document_id": "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list:manual_routing", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["network_policy_content", "manual_routing_list", "manual_routing", "outbound_domain_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["port"], "anchor": "schema-network_policy_content--manual_routing_list--manual_routing--port", "description": "Outbound Port. Outbound Port value.", "document_id": "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list:manual_routing", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["network_policy_content", "manual_routing_list", "manual_routing", "port"], "syntax": "attribute", "type": "number"}, {"aliases": ["protocol http"], "anchor": "section", "description": "Configuration parameter for protocol http.", "document_id": "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list:manual_routing:protocol_http", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["network_policy_content", "manual_routing_list", "manual_routing", "protocol_http"], "syntax": "attribute", "type": "object"}, {"aliases": ["protocol https"], "anchor": "section", "description": "Configuration parameter for protocol https.", "document_id": "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list:manual_routing:protocol_https", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["network_policy_content", "manual_routing_list", "manual_routing", "protocol_https"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_network_policy/properties/network_policy_content/manual_routing_list/manual_routing/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manual Routing. Routing or forwarding configuration", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [network_policy_content.manual_routing_list.manual_routing.http](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/network_policy_content/manual_routing_list/manual_routing/http/)
- [network_policy_content.manual_routing_list.manual_routing.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/network_policy_content/manual_routing_list/manual_routing/https/)
- [network_policy_content.manual_routing_list.manual_routing.protocol_http](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/network_policy_content/manual_routing_list/manual_routing/protocol_http/)
- [network_policy_content.manual_routing_list.manual_routing.protocol_https](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/network_policy_content/manual_routing_list/manual_routing/protocol_https/)
- [network_policy_content.manual_routing_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/network_policy_content/manual_routing_list/)
- [xcsh_bot_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/)
