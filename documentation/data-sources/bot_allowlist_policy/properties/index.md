---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_bot_allowlist_policy."
xcsh_docs: {"aliases": ["bot allowlist policy"], "body_bytes": 6207, "body_sha256": "sha256:5260dac7b982db56fd05dae0ac63a71d87e1587ab2cb6af8d012cb9e8fe6a3d0", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_allowlist_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_allowlist_policy:reference", "parent_id": "xcsh-docs:data-sources:bot_allowlist_policy:fundamentals", "path": "documentation/data-sources/bot_allowlist_policy/properties/index.md", "product": "distributed-cloud", "provider_name": "bot_allowlist_policy", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-1011231013012133-2121222203110231-0131022322212213-0313203222020020-1033321330133221-3030213123203133-3221132122003221-3323000023122122", "registry_path": "docs/guides/data-sources--bot_allowlist_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["allowlist policy content"], "anchor": "section", "description": "IP Allowlist. Allowlist Policy Content.", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["allowlist_policy_content"], "syntax": "attribute", "type": "object"}, {"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations.", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Description.", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier.", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Labels.", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["latest version"], "anchor": "schema-latest_version", "description": "Version. Version number or identifier", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["latest_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["name"], "anchor": "schema-name", "description": "Name of the BotAllowlistPolicy to look up.", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace of the BotAllowlistPolicy.", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_allowlist_policy/properties/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Property reference for xcsh_bot_allowlist_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": [], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_bot_allowlist_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/)
- Property reference

## Direct properties

- [allowlist_policy_content](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/): complete subsection reference.

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
| `allowlist_policy_content` | [allowlist_policy_content](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/#section) |
| `allowlist_policy_content.ip_allowlist` | [allowlist_policy_content.ip_allowlist](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/ip_allowlist/#section) |
| `allowlist_policy_content.ip_allowlist.ip_detail` | [allowlist_policy_content.ip_allowlist.ip_detail](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/ip_allowlist/ip_detail/#section) |
| `allowlist_policy_content.ip_allowlist.ip_detail.ip_description` | [allowlist_policy_content.ip_allowlist.ip_detail.ip_description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/ip_allowlist/ip_detail/#schema-allowlist_policy_content--ip_allowlist--ip_detail--ip_description) |
| `allowlist_policy_content.ip_allowlist.ip_detail.ip_value` | [allowlist_policy_content.ip_allowlist.ip_detail.ip_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/ip_allowlist/ip_detail/#schema-allowlist_policy_content--ip_allowlist--ip_detail--ip_value) |
| `allowlist_policy_content.ip_allowlist.ip_prefix_detail` | [allowlist_policy_content.ip_allowlist.ip_prefix_detail](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/ip_allowlist/ip_prefix_detail/#section) |
| `allowlist_policy_content.ip_allowlist.ip_prefix_detail.ip_description` | [allowlist_policy_content.ip_allowlist.ip_prefix_detail.ip_description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/ip_allowlist/ip_prefix_detail/#schema-allowlist_policy_content--ip_allowlist--ip_prefix_detail--ip_description) |
| `allowlist_policy_content.ip_allowlist.ip_prefix_detail.ip_value` | [allowlist_policy_content.ip_allowlist.ip_prefix_detail.ip_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/ip_allowlist/ip_prefix_detail/#schema-allowlist_policy_content--ip_allowlist--ip_prefix_detail--ip_value) |
| `allowlist_policy_content.ip_range_allowlist` | [allowlist_policy_content.ip_range_allowlist](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/ip_range_allowlist/#section) |
| `allowlist_policy_content.ip_range_allowlist.end_with` | [allowlist_policy_content.ip_range_allowlist.end_with](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/ip_range_allowlist/#schema-allowlist_policy_content--ip_range_allowlist--end_with) |
| `allowlist_policy_content.ip_range_allowlist.ip_description` | [allowlist_policy_content.ip_range_allowlist.ip_description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/ip_range_allowlist/#schema-allowlist_policy_content--ip_range_allowlist--ip_description) |
| `allowlist_policy_content.ip_range_allowlist.start_with` | [allowlist_policy_content.ip_range_allowlist.start_with](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/ip_range_allowlist/#schema-allowlist_policy_content--ip_range_allowlist--start_with) |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/#schema-description) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/#schema-labels) |
| `latest_version` | [latest_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/#schema-latest_version) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/#schema-namespace) |
