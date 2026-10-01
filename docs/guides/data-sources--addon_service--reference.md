---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_addon_service."
xcsh_docs: {"aliases": [], "body_bytes": 2741, "body_sha256": "sha256:592561543dd84d2deacb540ad15b7af84212499fcd80f92c848a6bd994880959", "canonical_id": "xcsh-docs:data-sources:addon_service:reference", "child_ids": [], "collection_id": "xcsh-docs:data-sources:addon_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:addon_service:reference", "parent_id": "xcsh-docs:data-sources:addon_service:fundamentals", "path": "docs/guides/data-sources--addon_service--reference.md", "provider_name": "addon_service", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/addon_service/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_addon_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_addon_service](../data-sources/addon_service.md)
- Property reference

## Direct properties

<a id="schema-activation_type"></a>

### activation_type property

Type: `"string"`. Computed.

How the addon service is activated. Possible values: \`self\` (user can activate directly),
\`partial\` (requires partial SRE management), \`managed\` (requires full manual intervention).

<a id="schema-addon_service_group_display_name"></a>

### addon_service_group_display_name property

Type: `"string"`. Computed.

Display name of the addon service group.

<a id="schema-addon_service_group_name"></a>

### addon_service_group_name property

Type: `"string"`. Computed.

Name of the addon service group this service belongs to.

<a id="schema-display_name"></a>

### display_name property

Type: `"string"`. Computed.

Human-readable display name of the addon service.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the data source.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the addon service (e.g., \`bot\_defense\`, \`client\_side\_defense\`).

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Optional, Computed.

Namespace where the addon service is defined. Usually \`shared\`.

<a id="schema-tier"></a>

### tier property

Type: `"string"`. Computed.

Subscription tier required for this addon service. Possible values: \`NO\_TIER\`, \`BASIC\`,
\`STANDARD\`, \`ADVANCED\`, \`PREMIUM\`.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `activation_type` | [activation_type](data-sources--addon_service--reference.md#schema-activation_type) |
| `addon_service_group_display_name` | [addon_service_group_display_name](data-sources--addon_service--reference.md#schema-addon_service_group_display_name) |
| `addon_service_group_name` | [addon_service_group_name](data-sources--addon_service--reference.md#schema-addon_service_group_name) |
| `display_name` | [display_name](data-sources--addon_service--reference.md#schema-display_name) |
| `id` | [id](data-sources--addon_service--reference.md#schema-id) |
| `name` | [name](data-sources--addon_service--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--addon_service--reference.md#schema-namespace) |
| `tier` | [tier](data-sources--addon_service--reference.md#schema-tier) |

## Next pages

- [xcsh_addon_service](../data-sources/addon_service.md)
