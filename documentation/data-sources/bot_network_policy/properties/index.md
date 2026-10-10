---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_bot_network_policy."
xcsh_docs: {"aliases": ["bot network policy"], "body_bytes": 6643, "body_sha256": "sha256:3802a2663ed24196082d4edc6902559f3bde2e87f8e82e8c0c7dd1fed5353da8", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_network_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_network_policy:reference", "parent_id": "xcsh-docs:data-sources:bot_network_policy:fundamentals", "path": "documentation/data-sources/bot_network_policy/properties/index.md", "product": "distributed-cloud", "provider_name": "bot_network_policy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1000220123222103-1311203111200122-1310233113223010-2303233320132230-2023011120122200-0020202331302112-2201330132003322-1123131221011230", "registry_path": "docs/guides/data-sources--bot_network_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations.", "document_id": "xcsh-docs:data-sources:bot_network_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Description.", "document_id": "xcsh-docs:data-sources:bot_network_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier.", "document_id": "xcsh-docs:data-sources:bot_network_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Labels.", "document_id": "xcsh-docs:data-sources:bot_network_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["latest version"], "anchor": "schema-latest_version", "description": "Version. Version number or identifier", "document_id": "xcsh-docs:data-sources:bot_network_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["latest_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["name"], "anchor": "schema-name", "description": "Name of the BotNetworkPolicy to look up.", "document_id": "xcsh-docs:data-sources:bot_network_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace of the BotNetworkPolicy.", "document_id": "xcsh-docs:data-sources:bot_network_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["network policy content"], "anchor": "section", "description": "Configuration parameter for network policy content.", "document_id": "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["network_policy_content"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_network_policy/properties/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Property reference for xcsh_bot_network_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_bot_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels.

<a id="schema-latest_version"></a>

### latest_version property

Type: `"string"`. Computed.

Version. Version number or identifier

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the BotNetworkPolicy to look up.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace of the BotNetworkPolicy.

- [network_policy_content](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/network_policy_content/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/#schema-description) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/#schema-labels) |
| `latest_version` | [latest_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/#schema-latest_version) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/#schema-namespace) |
| `network_policy_content` | [network_policy_content](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/network_policy_content/#section) |
| `network_policy_content.manual_routing_list` | [network_policy_content.manual_routing_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/network_policy_content/manual_routing_list/#section) |
| `network_policy_content.manual_routing_list.manual_routing` | [network_policy_content.manual_routing_list.manual_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/network_policy_content/manual_routing_list/manual_routing/#section) |
| `network_policy_content.manual_routing_list.manual_routing.domain_name` | [network_policy_content.manual_routing_list.manual_routing.domain_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/network_policy_content/manual_routing_list/manual_routing/#schema-network_policy_content--manual_routing_list--manual_routing--domain_name) |
| `network_policy_content.manual_routing_list.manual_routing.http` | [network_policy_content.manual_routing_list.manual_routing.http](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/network_policy_content/manual_routing_list/manual_routing/http/#section) |
| `network_policy_content.manual_routing_list.manual_routing.https` | [network_policy_content.manual_routing_list.manual_routing.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/network_policy_content/manual_routing_list/manual_routing/https/#section) |
| `network_policy_content.manual_routing_list.manual_routing.outbound_domain_name` | [network_policy_content.manual_routing_list.manual_routing.outbound_domain_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/network_policy_content/manual_routing_list/manual_routing/#schema-network_policy_content--manual_routing_list--manual_routing--outbound_domain_name) |
| `network_policy_content.manual_routing_list.manual_routing.port` | [network_policy_content.manual_routing_list.manual_routing.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/network_policy_content/manual_routing_list/manual_routing/#schema-network_policy_content--manual_routing_list--manual_routing--port) |
| `network_policy_content.manual_routing_list.manual_routing.protocol_http` | [network_policy_content.manual_routing_list.manual_routing.protocol_http](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/network_policy_content/manual_routing_list/manual_routing/protocol_http/#section) |
| `network_policy_content.manual_routing_list.manual_routing.protocol_https` | [network_policy_content.manual_routing_list.manual_routing.protocol_https](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/network_policy_content/manual_routing_list/manual_routing/protocol_https/#section) |
| `network_policy_content.upstream_routing_list` | [network_policy_content.upstream_routing_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/network_policy_content/upstream_routing_list/#section) |
| `network_policy_content.upstream_routing_list.upstream_routing` | [network_policy_content.upstream_routing_list.upstream_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/network_policy_content/upstream_routing_list/upstream_routing/#section) |
| `network_policy_content.upstream_routing_list.upstream_routing.domain_name` | [network_policy_content.upstream_routing_list.upstream_routing.domain_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/network_policy_content/upstream_routing_list/upstream_routing/#schema-network_policy_content--upstream_routing_list--upstream_routing--domain_name) |
