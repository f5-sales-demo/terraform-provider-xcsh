---
page_title: "network_policy_content.manual_routing_list.manual_routing"
subcategory: ""
description: "network_policy_content.manual_routing_list.manual_routing for xcsh_bot_network_policy."
xcsh_docs: {"aliases": [], "body_bytes": 2976, "body_sha256": "sha256:6a2ae178a6d1a591d16e31e291caab5f7485526a28f77e7bcb95c78ef624af8c", "canonical_id": "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list:manual_routing", "child_ids": ["xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list:manual_routing:http", "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list:manual_routing:https", "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list:manual_routing:protocol_http", "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list:manual_routing:protocol_https"], "collection_id": "xcsh-docs:data-sources:bot_network_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list:manual_routing", "parent_id": "xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content:manual_routing_list", "path": "docs/guides/data-sources--bot_network_policy--properties--network_policy_content--manual_routing_list--manual_routing.md", "provider_name": "bot_network_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["network_policy_content", "manual_routing_list", "manual_routing"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_network_policy/properties/network_policy_content/manual_routing_list/manual_routing/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "network_policy_content.manual_routing_list.manual_routing for xcsh_bot_network_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# network_policy_content.manual_routing_list.manual_routing

Breadcrumbs:

- [xcsh_bot_network_policy](../data-sources/bot_network_policy.md)
- [Property reference](data-sources--bot_network_policy--reference.md)
- [network_policy_content](data-sources--bot_network_policy--properties--network_policy_content.md)
- [network_policy_content.manual_routing_list](data-sources--bot_network_policy--properties--network_policy_content--manual_routing_list.md)
- network_policy_content.manual_routing_list.manual_routing

<a id="section"></a>

Type: `"list"`. Computed.

Manual Routing. Routing or forwarding configuration

## Direct properties

<a id="schema-network_policy_content--manual_routing_list--manual_routing--domain_name"></a>

### domain_name property

Type: `"string"`. Computed.

Inbound FQDN. Inbound FQDN value.

- [http](data-sources--bot_network_policy--properties--network_policy_content--manual_routing_list--manual_routing--http.md): complete subsection reference.

- [https](data-sources--bot_network_policy--properties--network_policy_content--manual_routing_list--manual_routing--https.md): complete subsection reference.

<a id="schema-network_policy_content--manual_routing_list--manual_routing--outbound_domain_name"></a>

### outbound_domain_name property

Type: `"string"`. Computed.

Outbound FQDN / IP. Outbound FQDN or IP value.

<a id="schema-network_policy_content--manual_routing_list--manual_routing--port"></a>

### port property

Type: `"number"`. Computed.

Outbound Port. Outbound Port value.

- [protocol_http](data-sources--bot_network_policy--properties--network_policy_content--manual_routing_list--manual_routing--protocol_http.md): complete subsection reference.

- [protocol_https](data-sources--bot_network_policy--properties--network_policy_content--manual_routing_list--manual_routing--protocol_https.md): complete subsection reference.

## Next pages

- [network_policy_content.manual_routing_list.manual_routing.http](data-sources--bot_network_policy--properties--network_policy_content--manual_routing_list--manual_routing--http.md)
- [network_policy_content.manual_routing_list.manual_routing.https](data-sources--bot_network_policy--properties--network_policy_content--manual_routing_list--manual_routing--https.md)
- [network_policy_content.manual_routing_list.manual_routing.protocol_http](data-sources--bot_network_policy--properties--network_policy_content--manual_routing_list--manual_routing--protocol_http.md)
- [network_policy_content.manual_routing_list.manual_routing.protocol_https](data-sources--bot_network_policy--properties--network_policy_content--manual_routing_list--manual_routing--protocol_https.md)
- [network_policy_content.manual_routing_list](data-sources--bot_network_policy--properties--network_policy_content--manual_routing_list.md)
- [xcsh_bot_network_policy](../data-sources/bot_network_policy.md)
