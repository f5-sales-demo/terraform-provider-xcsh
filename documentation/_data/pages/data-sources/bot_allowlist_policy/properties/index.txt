---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_bot_allowlist_policy."
xcsh_docs: {"aliases": ["bot allowlist policy"], "body_bytes": 6502, "body_sha256": "sha256:1f07d46aef4cb7ea32c86a6d4b000c756494220878b32165b669fafe7b9ae701", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_allowlist_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_allowlist_policy:reference", "parent_id": "xcsh-docs:data-sources:bot_allowlist_policy:fundamentals", "path": "documentation/data-sources/bot_allowlist_policy/properties/index.md", "product": "distributed-cloud", "provider_name": "bot_allowlist_policy", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-1011231013012133-2121222203110231-0131022322212213-0313203222020020-1033321330133221-3030213123203133-3221132122003221-3323000023122122", "registry_path": "docs/guides/data-sources--bot_allowlist_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["allowlist policy content"], "anchor": "section", "description": "IP Allowlist. Allowlist Policy Content.", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:properties:allowlist_policy_content", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["allowlist_policy_content"], "syntax": "attribute", "type": "object"}, {"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations.", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Description.", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier.", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Labels.", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["latest version"], "anchor": "schema-latest_version", "description": "Version. Version number or identifier", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["latest_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["name"], "anchor": "schema-name", "description": "Name of the BotAllowlistPolicy to look up.", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace of the BotAllowlistPolicy.", "document_id": "xcsh-docs:data-sources:bot_allowlist_policy:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_allowlist_policy/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_bot_allowlist_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [allowlist_policy_content](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/allowlist_policy_content/)
- [xcsh_bot_allowlist_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/)
