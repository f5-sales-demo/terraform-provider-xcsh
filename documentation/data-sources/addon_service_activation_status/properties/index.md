---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_addon_service_activation_status."
xcsh_docs: {"aliases": ["addon service activation status"], "body_bytes": 2358, "body_sha256": "sha256:d49c886d0aeafaa697a41ec4fca32c35deb1798e0cb6c169d5cc8f73a8f47b3e", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:addon_service_activation_status:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:addon_service_activation_status:reference", "parent_id": "xcsh-docs:data-sources:addon_service_activation_status:fundamentals", "path": "documentation/data-sources/addon_service_activation_status/properties/index.md", "product": "distributed-cloud", "provider_name": "addon_service_activation_status", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3310012331220020-2113213123322210-1321301222103133-2120010110020003-1233221302030322-0221010010110101-3101032013202133-0120012300012123", "registry_path": "docs/guides/data-sources--addon_service_activation_status--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["addon service"], "anchor": "schema-addon_service", "description": "Name of the addon service to check (e.g., `bot_defense`, `client_side_defense`).", "document_id": "xcsh-docs:data-sources:addon_service_activation_status:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["addon_service"], "syntax": "attribute", "type": "string"}, {"aliases": ["can activate"], "anchor": "schema-can_activate", "description": "Whether the addon service can be activated. True if state is `AS_NONE` (not yet subscribed) or `AS_SUBSCRIBED` (already active).", "document_id": "xcsh-docs:data-sources:addon_service_activation_status:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["can_activate"], "syntax": "attribute", "type": "bool"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the data source.", "document_id": "xcsh-docs:data-sources:addon_service_activation_status:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["message"], "anchor": "schema-message", "description": "Human-readable message describing the current activation status.", "document_id": "xcsh-docs:data-sources:addon_service_activation_status:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["message"], "syntax": "attribute", "type": "string"}, {"aliases": ["state"], "anchor": "schema-state", "description": "Current state of the addon service subscription. Possible values: `AS_NONE`, `AS_PENDING`, `AS_SUBSCRIBED`, `AS_ERROR`.", "document_id": "xcsh-docs:data-sources:addon_service_activation_status:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["state"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/addon_service_activation_status/properties/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Property reference for xcsh_addon_service_activation_status.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": [], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_addon_service_activation_status](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service_activation_status/)
- Property reference

## Direct properties

<a id="schema-addon_service"></a>

### addon_service property

Type: `"string"`. Required.

Name of the addon service to check (e.g., \`bot\_defense\`, \`client\_side\_defense\`).

<a id="schema-can_activate"></a>

### can_activate property

Type: `"bool"`. Computed.

Whether the addon service can be activated. True if state is \`AS\_NONE\` (not yet subscribed) or
\`AS\_SUBSCRIBED\` (already active).

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the data source.

<a id="schema-message"></a>

### message property

Type: `"string"`. Computed.

Human-readable message describing the current activation status.

<a id="schema-state"></a>

### state property

Type: `"string"`. Computed.

Current state of the addon service subscription. Possible values: \`AS\_NONE\`, \`AS\_PENDING\`,
\`AS\_SUBSCRIBED\`, \`AS\_ERROR\`.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `addon_service` | [addon_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service_activation_status/properties/#schema-addon_service) |
| `can_activate` | [can_activate](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service_activation_status/properties/#schema-can_activate) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service_activation_status/properties/#schema-id) |
| `message` | [message](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service_activation_status/properties/#schema-message) |
| `state` | [state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service_activation_status/properties/#schema-state) |

## Next pages

- [xcsh_addon_service_activation_status](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service_activation_status/)
