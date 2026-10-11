---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_addon_service."
xcsh_docs: {"aliases": ["addon service"], "body_bytes": 3143, "body_sha256": "sha256:c6473bf6f4f4802929b7256b5ab103326db5ea170c87dc8929c2ef842b127e3c", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:addon_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:addon_service:reference", "parent_id": "xcsh-docs:data-sources:addon_service:fundamentals", "path": "documentation/data-sources/addon_service/properties/index.md", "product": "distributed-cloud", "provider_name": "addon_service", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-0101311200030200-2033233113311222-3123301302321122-2232020112101123-3323000031131011-2103311130212112-2210113200221131-3330200122001111", "registry_path": "docs/guides/data-sources--addon_service--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["activation type"], "anchor": "schema-activation_type", "description": "How the addon service is activated. Possible values: `self` (user can activate directly), `partial` (requires partial SRE management), `managed` (requires full manual intervention).", "document_id": "xcsh-docs:data-sources:addon_service:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["activation_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["addon service group display name"], "anchor": "schema-addon_service_group_display_name", "description": "Display name of the addon service group.", "document_id": "xcsh-docs:data-sources:addon_service:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["addon_service_group_display_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["addon service group name"], "anchor": "schema-addon_service_group_name", "description": "Name of the addon service group this service belongs to.", "document_id": "xcsh-docs:data-sources:addon_service:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["addon_service_group_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["display name"], "anchor": "schema-display_name", "description": "Human-readable display name of the addon service.", "document_id": "xcsh-docs:data-sources:addon_service:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["display_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the data source.", "document_id": "xcsh-docs:data-sources:addon_service:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["name"], "anchor": "schema-name", "description": "Name of the addon service (e.g., `bot_defense`, `client_side_defense`).", "document_id": "xcsh-docs:data-sources:addon_service:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace where the addon service is defined. Usually `shared`.", "document_id": "xcsh-docs:data-sources:addon_service:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["tier"], "anchor": "schema-tier", "description": "Subscription tier required for this addon service. Possible values: `NO_TIER`, `BASIC`, `STANDARD`, `ADVANCED`, `PREMIUM`.", "document_id": "xcsh-docs:data-sources:addon_service:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tier"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/addon_service/properties/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Property reference for xcsh_addon_service.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_addon_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service/)
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
| `activation_type` | [activation_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service/properties/#schema-activation_type) |
| `addon_service_group_display_name` | [addon_service_group_display_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service/properties/#schema-addon_service_group_display_name) |
| `addon_service_group_name` | [addon_service_group_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service/properties/#schema-addon_service_group_name) |
| `display_name` | [display_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service/properties/#schema-display_name) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service/properties/#schema-id) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service/properties/#schema-namespace) |
| `tier` | [tier](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service/properties/#schema-tier) |
