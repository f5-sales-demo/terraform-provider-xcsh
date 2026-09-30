---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_bot_network_policy."
xcsh_docs: {"aliases": [], "body_bytes": 5617, "body_sha256": "sha256:7f470bb7eff0573bb4314a0c556c8d3f7fa05bbf24dcf6b56e7a125ee102c534", "canonical_id": "xcsh-docs:data-sources:bot_network_policy:reference", "child_ids": ["xcsh-docs:data-sources:bot_network_policy:properties:network_policy_content"], "collection_id": "xcsh-docs:data-sources:bot_network_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_network_policy:reference", "parent_id": "xcsh-docs:data-sources:bot_network_policy:fundamentals", "path": "docs/guides/data-sources--bot_network_policy--reference.md", "provider_name": "bot_network_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_network_policy/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_bot_network_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_bot_network_policy](../data-sources/bot_network_policy.md)
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

- [network_policy_content](data-sources--bot_network_policy--properties--network_policy_content.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--bot_network_policy--reference.md#schema-annotations) |
| `description` | [description](data-sources--bot_network_policy--reference.md#schema-description) |
| `id` | [id](data-sources--bot_network_policy--reference.md#schema-id) |
| `labels` | [labels](data-sources--bot_network_policy--reference.md#schema-labels) |
| `latest_version` | [latest_version](data-sources--bot_network_policy--reference.md#schema-latest_version) |
| `name` | [name](data-sources--bot_network_policy--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--bot_network_policy--reference.md#schema-namespace) |
| `network_policy_content` | [network_policy_content](data-sources--bot_network_policy--properties--network_policy_content.md#section) |
| `network_policy_content.manual_routing_list` | [network_policy_content.manual_routing_list](data-sources--bot_network_policy--properties--network_policy_content--manual_routing_list.md#section) |
| `network_policy_content.manual_routing_list.manual_routing` | [network_policy_content.manual_routing_list.manual_routing](data-sources--bot_network_policy--properties--network_policy_content--manual_routing_list--manual_routing.md#section) |
| `network_policy_content.manual_routing_list.manual_routing.domain_name` | [network_policy_content.manual_routing_list.manual_routing.domain_name](data-sources--bot_network_policy--properties--network_policy_content--manual_routing_list--manual_routing.md#schema-network_policy_content--manual_routing_list--manual_routing--domain_name) |
| `network_policy_content.manual_routing_list.manual_routing.http` | [network_policy_content.manual_routing_list.manual_routing.http](data-sources--bot_network_policy--properties--network_policy_content--manual_routing_list--manual_routing--http.md#section) |
| `network_policy_content.manual_routing_list.manual_routing.https` | [network_policy_content.manual_routing_list.manual_routing.https](data-sources--bot_network_policy--properties--network_policy_content--manual_routing_list--manual_routing--https.md#section) |
| `network_policy_content.manual_routing_list.manual_routing.outbound_domain_name` | [network_policy_content.manual_routing_list.manual_routing.outbound_domain_name](data-sources--bot_network_policy--properties--network_policy_content--manual_routing_list--manual_routing.md#schema-network_policy_content--manual_routing_list--manual_routing--outbound_domain_name) |
| `network_policy_content.manual_routing_list.manual_routing.port` | [network_policy_content.manual_routing_list.manual_routing.port](data-sources--bot_network_policy--properties--network_policy_content--manual_routing_list--manual_routing.md#schema-network_policy_content--manual_routing_list--manual_routing--port) |
| `network_policy_content.manual_routing_list.manual_routing.protocol_http` | [network_policy_content.manual_routing_list.manual_routing.protocol_http](data-sources--bot_network_policy--properties--network_policy_content--manual_routing_list--manual_routing--protocol_http.md#section) |
| `network_policy_content.manual_routing_list.manual_routing.protocol_https` | [network_policy_content.manual_routing_list.manual_routing.protocol_https](data-sources--bot_network_policy--properties--network_policy_content--manual_routing_list--manual_routing--protocol_https.md#section) |
| `network_policy_content.upstream_routing_list` | [network_policy_content.upstream_routing_list](data-sources--bot_network_policy--properties--network_policy_content--upstream_routing_list.md#section) |
| `network_policy_content.upstream_routing_list.upstream_routing` | [network_policy_content.upstream_routing_list.upstream_routing](data-sources--bot_network_policy--properties--network_policy_content--upstream_routing_list--upstream_routing.md#section) |
| `network_policy_content.upstream_routing_list.upstream_routing.domain_name` | [network_policy_content.upstream_routing_list.upstream_routing.domain_name](data-sources--bot_network_policy--properties--network_policy_content--upstream_routing_list--upstream_routing.md#schema-network_policy_content--upstream_routing_list--upstream_routing--domain_name) |

## Next pages

- [network_policy_content](data-sources--bot_network_policy--properties--network_policy_content.md)
- [xcsh_bot_network_policy](../data-sources/bot_network_policy.md)
