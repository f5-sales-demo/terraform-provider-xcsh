---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_bot_allowlist_policy."
xcsh_docs: {"aliases": [], "body_bytes": 5332, "body_sha256": "sha256:aa352d1f8fd6887263e830ddc3ddfca28efadd290e6c59057f15d8cdc0eaa4ab", "canonical_id": "xcsh-docs:data-sources:bot_allowlist_policy:reference", "child_ids": ["xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content"], "collection_id": "xcsh-docs:data-sources:bot_allowlist_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_allowlist_policy:reference", "parent_id": "xcsh-docs:data-sources:bot_allowlist_policy:fundamentals", "path": "docs/guides/data-sources--bot_allowlist_policy--reference.md", "provider_name": "bot_allowlist_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_allowlist_policy/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_bot_allowlist_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_bot_allowlist_policy](../data-sources/bot_allowlist_policy.md)
- Property reference

## Direct properties

- [allowlist_policy_content](data-sources--bot_allowlist_policy--properties--allowlist_policy_content.md): complete subsection reference.

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

Name of the BotAllowlistPolicy to look up.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace of the BotAllowlistPolicy.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allowlist_policy_content` | [allowlist_policy_content](data-sources--bot_allowlist_policy--properties--allowlist_policy_content.md#section) |
| `allowlist_policy_content.ip_allowlist` | [allowlist_policy_content.ip_allowlist](data-sources--bot_allowlist_policy--properties--allowlist_policy_content--ip_allowlist.md#section) |
| `allowlist_policy_content.ip_allowlist.ip_detail` | [allowlist_policy_content.ip_allowlist.ip_detail](data-sources--bot_allowlist_policy--properties--allowlist_policy_content--ip_allowlist--ip_detail.md#section) |
| `allowlist_policy_content.ip_allowlist.ip_detail.ip_description` | [allowlist_policy_content.ip_allowlist.ip_detail.ip_description](data-sources--bot_allowlist_policy--properties--allowlist_policy_content--ip_allowlist--ip_detail.md#schema-allowlist_policy_content--ip_allowlist--ip_detail--ip_description) |
| `allowlist_policy_content.ip_allowlist.ip_detail.ip_value` | [allowlist_policy_content.ip_allowlist.ip_detail.ip_value](data-sources--bot_allowlist_policy--properties--allowlist_policy_content--ip_allowlist--ip_detail.md#schema-allowlist_policy_content--ip_allowlist--ip_detail--ip_value) |
| `allowlist_policy_content.ip_allowlist.ip_prefix_detail` | [allowlist_policy_content.ip_allowlist.ip_prefix_detail](data-sources--bot_allowlist_policy--properties--allowlist_policy_content--ip_allowlist--ip_prefix_detail.md#section) |
| `allowlist_policy_content.ip_allowlist.ip_prefix_detail.ip_description` | [allowlist_policy_content.ip_allowlist.ip_prefix_detail.ip_description](data-sources--bot_allowlist_policy--properties--allowlist_policy_content--ip_allowlist--ip_prefix_detail.md#schema-allowlist_policy_content--ip_allowlist--ip_prefix_detail--ip_description) |
| `allowlist_policy_content.ip_allowlist.ip_prefix_detail.ip_value` | [allowlist_policy_content.ip_allowlist.ip_prefix_detail.ip_value](data-sources--bot_allowlist_policy--properties--allowlist_policy_content--ip_allowlist--ip_prefix_detail.md#schema-allowlist_policy_content--ip_allowlist--ip_prefix_detail--ip_value) |
| `allowlist_policy_content.ip_range_allowlist` | [allowlist_policy_content.ip_range_allowlist](data-sources--bot_allowlist_policy--properties--allowlist_policy_content--ip_range_allowlist.md#section) |
| `allowlist_policy_content.ip_range_allowlist.end_with` | [allowlist_policy_content.ip_range_allowlist.end_with](data-sources--bot_allowlist_policy--properties--allowlist_policy_content--ip_range_allowlist.md#schema-allowlist_policy_content--ip_range_allowlist--end_with) |
| `allowlist_policy_content.ip_range_allowlist.ip_description` | [allowlist_policy_content.ip_range_allowlist.ip_description](data-sources--bot_allowlist_policy--properties--allowlist_policy_content--ip_range_allowlist.md#schema-allowlist_policy_content--ip_range_allowlist--ip_description) |
| `allowlist_policy_content.ip_range_allowlist.start_with` | [allowlist_policy_content.ip_range_allowlist.start_with](data-sources--bot_allowlist_policy--properties--allowlist_policy_content--ip_range_allowlist.md#schema-allowlist_policy_content--ip_range_allowlist--start_with) |
| `annotations` | [annotations](data-sources--bot_allowlist_policy--reference.md#schema-annotations) |
| `description` | [description](data-sources--bot_allowlist_policy--reference.md#schema-description) |
| `id` | [id](data-sources--bot_allowlist_policy--reference.md#schema-id) |
| `labels` | [labels](data-sources--bot_allowlist_policy--reference.md#schema-labels) |
| `latest_version` | [latest_version](data-sources--bot_allowlist_policy--reference.md#schema-latest_version) |
| `name` | [name](data-sources--bot_allowlist_policy--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--bot_allowlist_policy--reference.md#schema-namespace) |

## Next pages

- [allowlist_policy_content](data-sources--bot_allowlist_policy--properties--allowlist_policy_content.md)
- [xcsh_bot_allowlist_policy](../data-sources/bot_allowlist_policy.md)
