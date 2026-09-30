---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_addon_service_activation_status."
xcsh_docs: {"aliases": [], "body_bytes": 1892, "body_sha256": "sha256:b97cbb1b96751a48f7560eeb21b93e297c5c8bf5854ad2e3a4bd02a3f8d9c840", "canonical_id": "xcsh-docs:data-sources:addon_service_activation_status:reference", "child_ids": [], "collection_id": "xcsh-docs:data-sources:addon_service_activation_status:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:addon_service_activation_status:reference", "parent_id": "xcsh-docs:data-sources:addon_service_activation_status:fundamentals", "path": "docs/guides/data-sources--addon_service_activation_status--reference.md", "provider_name": "addon_service_activation_status", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/addon_service_activation_status/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_addon_service_activation_status.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_addon_service_activation_status](../data-sources/addon_service_activation_status.md)
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
| `addon_service` | [addon_service](data-sources--addon_service_activation_status--reference.md#schema-addon_service) |
| `can_activate` | [can_activate](data-sources--addon_service_activation_status--reference.md#schema-can_activate) |
| `id` | [id](data-sources--addon_service_activation_status--reference.md#schema-id) |
| `message` | [message](data-sources--addon_service_activation_status--reference.md#schema-message) |
| `state` | [state](data-sources--addon_service_activation_status--reference.md#schema-state) |

## Next pages

- [xcsh_addon_service_activation_status](../data-sources/addon_service_activation_status.md)
